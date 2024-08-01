package handler

import (
	"database/sql"
	"errors"
	"golang-online-shop/model"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		id := c.Param("id")

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

func CreateProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product model.Product
		if err := c.Bind(&product); err != nil {
			log.Printf("An error occurred while reading the request body: %v\n", err)
			c.JSON(400, gin.H{"error": "Product data not valid"})
			return
		}
		product.ID = uuid.New().String()

		if err := model.InsertProduct(db, product); err != nil {
			log.Printf("An error occurred on create product: %v\n", err)
			c.JSON(500, gin.H{"error": "An error occurred on the server"})
			return
		}

		c.JSON(201, product)

	}
}

func UpdateProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		var productReq model.Product
		if err := c.Bind(&productReq); err != nil {
			log.Printf("An error occurred while reading the request body: %v\n", err)
			c.JSON(400, gin.H{"error": "Product data not valid"})
			return
		}

		product, err := model.SelectProductByID(db, id)
		if err != nil {
			log.Printf("An error occurred while retrieving the product: %v\n", err)
			c.JSON(400, gin.H{"error": "Product data not valid"})
			return
		}
		// if err != nil {
		// 	if errors.Is(err, sql.ErrNoRows) {
		// 		log.Printf("An error occurred while retrieving the product: %v\n", err)
		// 		c.JSON(404, gin.H{"error": "Product not found"})
		// 		return
		// 	}
		// 	log.Printf("An error occurred while retrieving the product: %v\n", err)
		// 	c.JSON(500, gin.H{"error": "There is an error"})
		// 	return
		// }

		if productReq.Name != "" {
			product.Name = productReq.Name
		}
		if productReq.Price != 0 {
			product.Price = productReq.Price
		}

		if err := model.UpdateProduct(db, product); err != nil {
			log.Printf("An error occurred while update the product: %v\n", err)
			c.JSON(500, gin.H{"error": "Product data not valid"})
			return

		}

		c.JSON(201, product)
	}
}

func DeleteProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {}
}
