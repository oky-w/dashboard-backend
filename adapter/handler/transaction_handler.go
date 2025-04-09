package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/okyws/dashboard-backend/constants"
	"github.com/okyws/dashboard-backend/domain"
	"github.com/okyws/dashboard-backend/dto"
	"github.com/okyws/dashboard-backend/services"
	"github.com/okyws/dashboard-backend/utils"
	"gorm.io/gorm"
)

// TransactionHandler is the HTTP handler for the transaction service
type TransactionHandler struct {
	TransactionService *services.TransactionService
}

// NewTransactionHandler creates a new transaction handler
func NewTransactionHandler(transactionService *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		TransactionService: transactionService,
	}
}

// HandleTransactionProcess implements the HTTP handler for processing a transaction
func (h *TransactionHandler) HandleTransactionProcess(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		utils.ErrorResponse(c, http.StatusMethodNotAllowed, constants.MsgNotAllowed)
		return
	}

	var request dto.TransactionCreateDTO

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.TransactionService.ProcessTransaction(request.FromAccountNumber, request.ToAccountNumber, request.TransactionType, request.Amount)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err == gorm.ErrRecordNotFound {
		utils.ErrorResponse(c, http.StatusUnprocessableEntity, constants.MsgUnprocessable)
		return
	}

	utils.ResponseJSON(c, nil, http.StatusOK, "Balance transferred successfully")
}

// HandleGetAllTransactions implements the HTTP handler for getting all transactions
func (h *TransactionHandler) HandleGetAllTransactions(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		utils.ErrorResponse(c, http.StatusMethodNotAllowed, constants.MsgNotAllowed)
		return
	}

	limit, offset := utils.GetPaginationParams(c)

	transactions, err := h.TransactionService.GetAllTransactions(limit, offset)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	transactionDTOs := make([]dto.TransactionDTO, len(transactions))
	for i, transaction := range transactions {
		transactionDTOs[i] = *domain.MapTransactionToDTO(&transaction)
	}

	utils.ResponseJSON(c, transactionDTOs, http.StatusOK, "Transactions retrieved successfully")
}

// HandleGetAllTransactionsByAccountID implements the HTTP handler for getting all transactions by account ID
func (h *TransactionHandler) HandleGetAllTransactionsByAccountID(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		utils.ErrorResponse(c, http.StatusMethodNotAllowed, constants.MsgNotAllowed)
		return
	}

	userID := c.Param("account_id")

	transactions, err := h.TransactionService.GetTransactionByAccountID(userID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if transactions == nil {
		utils.ErrorResponse(c, http.StatusNotFound, constants.MsgNotFound)
		return
	}

	transactionDTOs := make([]dto.TransactionDTO, len(transactions))
	for i, transaction := range transactions {
		transactionDTOs[i] = *domain.MapTransactionToDTO(&transaction)
	}

	utils.ResponseJSON(c, transactionDTOs, http.StatusOK, "Transactions retrieved successfully")
}

// HandleGetTransactionByID implements the HTTP handler for getting a transaction by ID
func (h *TransactionHandler) HandleGetTransactionByID(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		utils.ErrorResponse(c, http.StatusMethodNotAllowed, constants.MsgNotAllowed)
		return
	}

	id := c.Param("id")

	transaction, err := h.TransactionService.GetTransactionByID(id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if transaction == nil {
		utils.ErrorResponse(c, http.StatusNotFound, constants.MsgNotFound)
		return
	}

	transactionDTO := *domain.MapTransactionToDTO(transaction)

	utils.ResponseJSON(c, transactionDTO, http.StatusOK, "Transaction fetched successfully")
}
