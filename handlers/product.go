package handlers

import (
	"fmt"
	"net/http"
	"product/models"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var products []models.Product
		db.Find(&products)
		c.JSON(http.StatusOK, products)
	}
}

func CreateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		groups, exist := c.Get("userGroups")
		if !exist {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user is not authorized"})
			return
		}
		availableGroups, ok := groups.([]string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user is not authorized"})
			return
		}
		isWriterGroupAvailable := false
		for _, g := range availableGroups {
			if strings.Contains(g, "goappwriter") {
				fmt.Println(g)
				isWriterGroupAvailable = true
				break
			}
		}
		if !isWriterGroupAvailable {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user is not authorized"})
			return
		}
		var product models.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&product)
		c.JSON(http.StatusCreated, product)
	}
}

func UpdateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product models.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Updates(&product)
		c.JSON(http.StatusAccepted, product)
	}
}

func DeleteProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product models.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Delete(&product)
		c.JSON(http.StatusAccepted, product)
	}
}
