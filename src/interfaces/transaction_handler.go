package interfaces

import (
	"ezustore/src/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct{ usecase domain.TransactionUsecase }

func NewTransactionHandler(rg *gin.RouterGroup, uc domain.TransactionUsecase) {
	h := &TransactionHandler{usecase: uc}
	rg.POST("/transactions", h.Create)
	rg.GET("/transactions", h.GetAll)
	rg.GET("/transactions/:id", h.GetByID)
}

func (h *TransactionHandler) Create(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var req struct {
		Items []domain.CreateTransactionItemInput `json:"items"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.usecase.Create(userID, req.Items); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Transaksi berhasil"})
}

func (h *TransactionHandler) GetAll(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	txs, err := h.usecase.GetByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, txs)
}

func (h *TransactionHandler) GetByID(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	id, _ := strconv.Atoi(c.Param("id"))
	trx, err := h.usecase.GetByID(userID, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, trx)
}
