package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/robustrade/wallet-transfer-assignment/internal/domain"
	"github.com/robustrade/wallet-transfer-assignment/internal/service"
)

type TransferHandler struct {
	service *service.TransferService
}

func NewTransferHandler(svc *service.TransferService) *TransferHandler {
	return &TransferHandler{service: svc}
}

func (h *TransferHandler) Create(c *gin.Context) {
	var req CreateTransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	result, err := h.service.CreateTransfer(c.Request.Context(), service.CreateTransferInput{
		IdempotencyKey: req.IdempotencyKey,
		FromWalletID:   req.FromWalletID,
		ToWalletID:     req.ToWalletID,
		Amount:         req.Amount,
	})

	if err != nil {
		writeError(c, err)
		return
	}

	status := http.StatusCreated
	switch {
	case result.Replayed:
		status = http.StatusOK
	case result.Status == domain.TransferFailed:
		status = http.StatusUnprocessableEntity
	}
	c.JSON(status, toTransferResponse(result))
}

func (h *TransferHandler) Get(c *gin.Context) {
	result, err := h.service.GetTransfer(c.Request.Context(), c.Param("id"))
	if err != nil {
		writeError(c, err)
		return
	}
	c.JSON(http.StatusOK, toTransferResponse(result))
}
