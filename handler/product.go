package handler

import (
	"database/sql"
	"golang-online-shop/model"

	"github.com/gin-gonic/gin"
)

func ListProducts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: get from database
		products, err := model.SelectProduct(db)
		if err != nil {
			c.JSON(500, gin.H{"error": "There is an error"})
			return
		}
		// TODO: give response
		c.JSON(200, products)

	}

}

func GetProducts(c *gin.Context) {
	// TODO: read id from url

	// TODO: get from database with id

	// TODO: give response
	c.JSON(200, "OKE")
}
