// Package routes contains the API routes for the application and server setup.
package routes

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/okyws/dashboard-backend/adapter/handler"
	"github.com/okyws/dashboard-backend/adapter/repository"
	"github.com/okyws/dashboard-backend/config"
	"github.com/okyws/dashboard-backend/constants"
	"github.com/okyws/dashboard-backend/domain"
	"github.com/okyws/dashboard-backend/middleware"
	"github.com/okyws/dashboard-backend/services"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// RegisterRoutes registers all API routes
func RegisterRoutes(router *gin.Engine, db *gorm.DB, redisClient *redis.Client) {
	userRepo := repository.NewUserRepositoryAdapter(db)
	customerRepo := repository.NewCustomerRepositoryAdapter(db)
	bankInfoRepo := repository.NewBankAccountRepositoryAdapter(db)
	transactionRepo := repository.NewTransactionRepositoryAdapter(db)
	authRepo := repository.NewAuthRepositoryRedis(redisClient)

	userService := services.NewUserService(userRepo)
	customerService := services.NewCustomerService(customerRepo, userRepo)
	accountValidator := services.NewAccountValidator(userRepo, bankInfoRepo)
	bankInfoService := services.NewBankAccountService(userRepo, bankInfoRepo, accountValidator)
	transactionValidator := services.NewTransactionValidator(transactionRepo, bankInfoRepo)
	transactionService := services.NewTransactionService(db, transactionRepo, bankInfoRepo, transactionValidator)
	authService := services.NewAuthService(authRepo, userRepo)

	userHandler := handler.NewUserHandler(userService)
	customerHandler := handler.NewCustomerHandler(customerService)
	bankInfoHandler := handler.NewBankInfoHandler(bankInfoService)
	transactionHandler := handler.NewTransactionHandler(transactionService)
	authHandler := handler.NewAuthHandler(authService)

	apiRoutes := router.Group("/api/v1")
	userRoutes := apiRoutes.Group("/users", middleware.AuthMiddleware())

	userRoutes.GET("/", middleware.CheckRoleMiddleware("admin"), userHandler.HandleGetAllUsers)
	userRoutes.GET("/:id", middleware.CheckRoleMiddleware("admin"), userHandler.HandleGetUserByID)
	userRoutes.POST("/add", middleware.CheckRoleMiddleware("admin"), userHandler.HandleCreateUser)
	userRoutes.GET("/by-username/:username", middleware.CheckRoleMiddleware("admin"), userHandler.HandleGetUserByUsername)
	userRoutes.PUT("/:id/update", middleware.CheckRoleMiddleware("admin"), userHandler.HandleUpdateUser)
	userRoutes.DELETE("/:id/delete", middleware.CheckRoleMiddleware("admin"), userHandler.HandleDeleteUser)

	customerRoutes := apiRoutes.Group("/customers", middleware.AuthMiddleware())

	customerRoutes.GET("/", middleware.CheckRoleMiddleware("admin"), customerHandler.HandleGetAllCustomers)
	customerRoutes.POST("/add", middleware.CheckRoleMiddleware("admin"), customerHandler.HandleCreateCustomer)
	customerRoutes.GET("/:id", middleware.CheckRoleMiddleware("admin"), customerHandler.HandleGetCustomerByID)
	customerRoutes.GET("/by-user-id/:user_id", middleware.CheckRoleMiddleware("admin"), customerHandler.HandleGetCustomerByUserID)
	customerRoutes.PUT("/:id/update", middleware.CheckRoleMiddleware("admin"), customerHandler.HandleUpdateCustomer)
	customerRoutes.DELETE("/:id/delete", middleware.CheckRoleMiddleware("admin"), customerHandler.HandleDeleteCustomer)

	bankInfoRoutes := apiRoutes.Group("/bank-accounts", middleware.AuthMiddleware())

	bankInfoRoutes.GET("/", middleware.CheckRoleMiddleware("admin"), bankInfoHandler.HandleGetAllBankAccounts)
	bankInfoRoutes.POST("/add", middleware.CheckRoleMiddleware("admin"), bankInfoHandler.HandleCreateBankInfo)
	bankInfoRoutes.GET("/:id", middleware.CheckRoleMiddleware("admin"), bankInfoHandler.HandleGetBankInfoByID)
	bankInfoRoutes.GET("/by-user-id/:user_id", middleware.CheckRoleMiddleware("user"), bankInfoHandler.HandleGetBankInfoByUserID)
	bankInfoRoutes.DELETE("/:id/delete", middleware.CheckRoleMiddleware("user", "admin"), bankInfoHandler.HandleDeleteBankInfo)

	transactionRoutes := apiRoutes.Group("/transactions", middleware.AuthMiddleware())

	transactionRoutes.GET("/", middleware.CheckRoleMiddleware("admin"), transactionHandler.HandleGetAllTransactions)
	transactionRoutes.POST("/add", middleware.CheckRoleMiddleware("user"), transactionHandler.HandleTransactionProcess)
	transactionRoutes.GET("/by-account-id/:account_id", middleware.CheckRoleMiddleware("user"), transactionHandler.HandleGetAllTransactionsByAccountID)
	transactionRoutes.GET("/:id", middleware.CheckRoleMiddleware("admin"), transactionHandler.HandleGetTransactionByID)

	authRoutes := apiRoutes.Group("/auth")
	authRoutes.POST("/login", authHandler.Login)

	log.Info().Msg("Successfully configured routes with database " + db.Name())
}

// SetupRouter initializes the Gin router
func SetupRouter() (*gin.Engine, *gorm.DB, *redis.Client) {
	router := gin.Default()
	router.Use(middleware.ZerologMiddleware())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	db, redisClient, err := config.GetDatabaseConnection()
	if err != nil {
		log.Fatal().Err(err).Str("error", err.Error()).Msg(constants.MsgDBConnectFail)
	}

	// Register API routes
	RegisterRoutes(router, db, redisClient)

	return router, db, redisClient
}

// RunServer starts the Gin server
func RunServer() {
	router, db, redisClient := SetupRouter()

	configuration, err := domain.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Str("error", err.Error()).Msg(constants.MsgConfigLoadFail)
	}

	server := &http.Server{
		Addr:         ":" + configuration.ServerPort,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// initialize channels
	errChan := make(chan error, 1)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Info().Msg("Starting server on port " + configuration.ServerPort)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- err // Send error to channel
		}
	}()

	select {
	case <-stop:
		log.Info().Msg(constants.MsgServerShutdown)
	case err := <-errChan:
		log.Error().Err(err).Str("error", err.Error()).Msg(constants.MsgServerError)
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Error().Err(err).Str("error", err.Error()).Msg(constants.MsgServerShutdownErr)
	}

	// Cleanup resources
	defer cancel()
	defer config.CloseDatabase(db)
	defer config.CloseRedisClient(redisClient)

	log.Info().Msg(constants.MsgServerGraceful)
}
