package handler

import (
	"database/sql"
	"errors"
	"golang-online-shop/model"
	"log"

	"github.com/gin-gonic/gin"
)

func ListProducts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: get from database
		products, err := model.SelectProduct(db)
		if err != nil {
			log.Printf("An error occurred while retrieving the product: %v", err)
			c.JSON(500, gin.H{"error": "There is an error"})
			return
		}
		// TODO: give response
		c.JSON(200, products)

	}

}

func GetProducts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: read id from url
		id := c.Param(":id")

		// TODO: get from database with id
		product, err := model.SelectProductByID(db, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				log.Printf("An error occurred while retrieving the product: %v", err)
				c.JSON(404, gin.H{"error": "Product not found"})
				return

			}

			log.Printf("An error occurred while retrieving the product: %v", err)
			c.JSON(500, gin.H{"error": "There is an error"})
			return

		}

		// TODO: give response
		c.JSON(200, product)
	}

}
