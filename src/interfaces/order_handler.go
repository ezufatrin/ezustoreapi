package interfaces

import (
	"ezustore/src/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct{ usecase domain.OrderUsecase }

func NewOrderHandler(rg *gin.RouterGroup, uc domain.OrderUsecase) {
	h := &OrderHandler{usecase: uc}
	rg.POST("/orders", h.Checkout)
	rg.GET("/orders", h.GetHistory)
	rg.GET("/orders/:id", h.GetDetail)
}

func (h *OrderHandler) Checkout(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var req struct {
		Items []domain.OrderItemInput `json:"items"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.usecase.Checkout(userID, req.Items); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order created"})
}

func (h *OrderHandler) GetHistory(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	orders, err := h.usecase.GetHistory(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) GetDetail(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	id, _ := strconv.Atoi(c.Param("id"))
	order, err := h.usecase.GetDetail(userID, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}
