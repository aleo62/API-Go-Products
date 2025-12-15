package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	// DB Connection
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Repositories layer
	ProductRepository := repository.NewProductRepository(dbConnection)

	// Usecases layer
	ProductUsecase := usecase.NewProductUsecase(*ProductRepository)

	// Controllers layer
	ProductController := controller.NewProductController(*ProductUsecase)


	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/products", ProductController.CreateProduct)
	server.DELETE("/products/:id", ProductController.DeleteProduct)

	server.Run(":8080")

}