package interfaces

import (
	"ezustore/src/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct{ usecase domain.ProductUsecase }

func NewProductHandler(r *gin.Engine, uc domain.ProductUsecase) {
	h := &ProductHandler{usecase: uc}
	grp := r.Group("/products")
	{
		grp.GET("/", h.GetAll)
		grp.GET("/:id", h.GetByID)
	}
}

func RegisterProductProtectedRoutes(rg *gin.RouterGroup, uc domain.ProductUsecase) {
	h := &ProductHandler{usecase: uc}
	rg.POST("/products", h.Create)
	rg.PUT("/products/:id", h.Update)
	rg.DELETE("/products/:id", h.Delete)
}

func (h *ProductHandler) Create(c *gin.Context) {
	var req struct {
		Name     string  `json:"name"`
		Price    float64 `json:"price"`
		Stock    int     `json:"stock"`
		Category string  `json:"category"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.usecase.Create(req.Name, req.Price, req.Stock, req.Category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Product created"})
}

func (h *ProductHandler) GetAll(c *gin.Context) {
	ps, err := h.usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ps)
}

func (h *ProductHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	p, err := h.usecase.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, p)
}

func (h *ProductHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req struct {
		Name     string  `json:"name"`
		Price    float64 `json:"price"`
		Stock    int     `json:"stock"`
		Category string  `json:"category"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.usecase.Update(uint(id), req.Name, req.Price, req.Stock, req.Category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product updated"})
}

func (h *ProductHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.usecase.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
