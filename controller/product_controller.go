package controller

import (
	"go-api/model"
	"go-api/usecase"
	"go-api/validator"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUsecase usecase.ProductUsecase
}

type CreateProductRequest struct {
	ProductName  string   `json:"product_name" binding:"required"`
	Price        float64  `json:"price" binding:"required,gt=0"`
}


func NewProductController(usecase usecase.ProductUsecase) *productController {
	return &productController{
		productUsecase: usecase,
	}
}

func (p *productController) GetProducts(c *gin.Context) {
	products, err := p.productUsecase.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func (p *productController) CreateProduct(c *gin.Context) {
	var req CreateProductRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		validator.ValidateRequest(err, c)
		return
	}
	
	product := model.Product{
		Name:        req.ProductName,
		Price:       req.Price,
	}

	insertedProduct, err := p.productUsecase.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": insertedProduct})
}

func (p *productController) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        panic(err)
    }
	
	ok, err := p.productUsecase.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": ok, "message": "Product deleted successfully"})
}

func (p *productController) UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        panic(err)
    }
	
	var req usecase.UpdateProductRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		validator.ValidateRequest(err, c)
		return
	}

	ok, err := p.productUsecase.UpdateProduct(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": ok, "message": "Product updated successfully"})
}
