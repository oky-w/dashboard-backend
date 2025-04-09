// Package handler contains the HTTP handlers for the bank information service
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

// BankInfoHandlerAdapter is the HTTP handler for the bank information service
type BankInfoHandlerAdapter struct {
	BankInfoService services.BankAccountService
}

// NewBankInfoHandler creates a new bank information handler
func NewBankInfoHandler(service *services.BankAccountService) *BankInfoHandlerAdapter {
	return &BankInfoHandlerAdapter{BankInfoService: *service}
}

// HandleCreateBankInfo creates a new bank information
func (h *BankInfoHandlerAdapter) HandleCreateBankInfo(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		utils.ErrorResponse(c, http.StatusMethodNotAllowed, constants.MsgNotAllowed)
		return
	}

	var req dto.BankAccountCreateDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	bankInfo := &domain.BankAccount{
		UserID:      req.UserID,
		AccountType: req.AccountType,
		Balance:     req.Balance,
	}

	bankInfo, err := h.BankInfoService.CreateBankAccount(bankInfo)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	bankInfoDTO := *domain.MapBankAccountToDTO(bankInfo)

	utils.ResponseJSON(c, bankInfoDTO, http.StatusOK, "Bank information created successfully")
}

// HandleGetAllBankAccounts returns all bank information
func (h *BankInfoHandlerAdapter) HandleGetAllBankAccounts(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		utils.ErrorResponse(c, http.StatusMethodNotAllowed, constants.MsgNotAllowed)
		return
	}

	limit, offset := utils.GetPaginationParams(c)

	bankInfos, err := h.BankInfoService.GetAllBankAccount(limit, offset)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	bankInfoDTOs := make([]dto.BankAccountDTO, len(bankInfos))
	for i := range bankInfos {
		bankInfoDTOs[i] = *domain.MapBankAccountToDTO(&bankInfos[i])
	}

	utils.ResponseJSON(c, bankInfoDTOs, http.StatusOK, "Bank information fetched successfully")
}

// HandleGetBankInfoByID returns the bank information by ID
func (h *BankInfoHandlerAdapter) HandleGetBankInfoByID(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		utils.ErrorResponse(c, http.StatusMethodNotAllowed, constants.MsgNotAllowed)
		return
	}

	id := c.Param("id")

	bankInfo, err := h.BankInfoService.GetBankAccountByID(id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if bankInfo == nil || err == gorm.ErrRecordNotFound {
		utils.ErrorResponse(c, http.StatusNotFound, constants.MsgNotFound)
		return
	}

	bankInfoDTO := *domain.MapBankAccountToDTO(bankInfo)

	utils.ResponseJSON(c, bankInfoDTO, http.StatusOK, "Bank information fetched successfully")
}

// HandleDeleteBankInfo deletes a bank information
func (h *BankInfoHandlerAdapter) HandleDeleteBankInfo(c *gin.Context) {
	if c.Request.Method != http.MethodDelete {
		utils.ErrorResponse(c, http.StatusMethodNotAllowed, constants.MsgNotAllowed)
		return
	}

	id := c.Param("id")

	err := h.BankInfoService.DeleteBankAccount(id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err == gorm.ErrRecordNotFound {
		utils.ErrorResponse(c, http.StatusNotFound, constants.MsgNotFound)
		return
	}

	utils.ResponseJSON(c, nil, http.StatusOK, "Bank information deleted successfully")
}

// HandleGetBankInfoByUserID returns the bank information for a user
func (h *BankInfoHandlerAdapter) HandleGetBankInfoByUserID(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		utils.ErrorResponse(c, http.StatusMethodNotAllowed, constants.MsgNotAllowed)
		return
	}

	userID := c.Param("user_id")

	bankInfos, err := h.BankInfoService.GetByUserID(userID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if bankInfos == nil || err == gorm.ErrRecordNotFound {
		utils.ErrorResponse(c, http.StatusNotFound, constants.MsgNotFound)
		return
	}

	bankInfoDTOs := make([]dto.BankAccountDTO, len(bankInfos))
	for i, bankInfo := range bankInfos {
		bankInfoDTOs[i] = *domain.MapBankAccountToDTO(&bankInfo)
	}

	utils.ResponseJSON(c, bankInfoDTOs, http.StatusOK, "Bank information fetched successfully")
}
