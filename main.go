package main

import (
	"log"
	"product/handlers"
	"product/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Initialize PostgreSQL with GORM
	dsn := "host=postgres-postgresql.db.svc user=postgres password=vse4wQKN2x dbname=goappdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	db.AutoMigrate(&models.Product{})

	// Gin router
	r := gin.Default()

	// Keycloak auth middleware
	r.Use(handlers.KeycloakAuthMiddleware())

	// CRUD routes
	r.GET("/products", handlers.ListProducts(db))
	r.POST("/products", handlers.CreateProduct(db))
	r.PUT("/products/:id", handlers.UpdateProduct(db))
	r.DELETE("/products/:id", handlers.DeleteProduct(db))

	// Start server
	r.Run(":8080")
}
