package handler

import (
	"database/sql"
	"golang-online-shop/model"
	"log"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func ChekoutOrder(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: get data order from request
		var checkoutOrder model.Checkout
		if err := c.BindJSON(&checkoutOrder); err != nil {
			log.Printf("An error occurred while read request body: %v\n", err)
			c.JSON(400, gin.H{"error": "Product data not valid"})
			return
		}

		ids := []string{}
		orderQty := make(map[string]int32)
		for _, o := range checkoutOrder.Products {
			ids = append(ids, o.ID)
			orderQty[o.ID] = o.Quantity
		}

		// TODO: get product from database
		products, err := model.SelectProductIn(db, ids)
		if err != nil {
			log.Printf("An error occurred while retrieving products: %v\n", err)
			c.JSON(500, gin.H{"error": "an error occurred on the server"})
			return
		}

		// TODO: create password
		passcode := generatePasscode(5)

		// TODO: hash pasword
		hashcode, err := bcrypt.GenerateFromPassword([]byte(passcode), 10)

		if err != nil {
			log.Printf("An error occurred while create hash: %v\n", err)
			c.JSON(500, gin.H{"error": "an error occurred on the server"})
			return
		}
		hashcodeString := string(hashcode)
		// TODO: create order & detail

		order := model.Order{
			ID:         uuid.New().String(),
			Email:      checkoutOrder.Email,
			Address:    checkoutOrder.Address,
			Passcode:   &hashcodeString,
			GrandTotal: 0,
		}
		details := []model.OrderDetail{}

		for _, p := range products {
			total := p.Price * int64(orderQty[p.ID])

			detail := model.OrderDetail{
				ID:        uuid.New().String(),
				OrderID:   order.ID,
				ProductID: p.ID,
				Quantity:  orderQty[p.ID],
				Price:     p.Price,
				Total:     total,
			}
			details = append(details, detail)

			order.GrandTotal += total
		}

		model.CreateOrder(db, order, details)

		orderWithDetail := model.OrderWithDetail{
			Order:   order,
			Details: details,
		}

		orderWithDetail.Order.Passcode = &passcode

		c.JSON(200, orderWithDetail)

	}
}

func generatePasscode(length int) string {
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))

	code := make([]byte, length)

	for i := range code {
		code[i] = charset[randomGenerator.Intn(len(charset))]
	}
	return string(code)
}

func ConfirmOrder(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: get id from param
		id := c.Param("id")

		// TODO: red request body
		var confirmReq model.Confirm
		if err := c.BindJSON(&confirmReq); err != nil {
			log.Printf("An error occurred read request body: %v\n", err)
			c.JSON(400, gin.H{"error": "an error occurred on the server"})
			return

		}

		// TODO: get order from database
		order, err := model.SelectOrderByID(db, id)
		if err != nil {
			log.Printf("An error occurred read product order: %v\n", err)
			c.JSON(400, gin.H{"error": "Product order not valid"})
			return

		}

		if order.Passcode == nil {
			log.Println("Passcode not valid")
			c.JSON(400, gin.H{"error": "Product order not valid"})
			return

		}

		// TODO: match the order password
		if err = bcrypt.CompareHashAndPassword([]byte(*order.Passcode), []byte(confirmReq.Passcode)); err != nil {
			log.Println("An error occurred while matching the password")
			c.JSON(401, gin.H{"error": "not permitted to access orders"})
			return

		}

		// TODO: make sure the order has not been paid
		if order.PaidAt != nil {
			log.Println("order has been paid")
			c.JSON(400, gin.H{"error": "order has been paid"})
			return

		}

		// TODO: match the payment amount
		if order.GrandTotal != confirmReq.Amount {
			log.Printf("the price amount does not match: %v", confirmReq.Amount)
			c.JSON(400, gin.H{"error": "payment amount does not match"})
			return

		}
		// TODO: update the order status information
		current := time.Now()
		if err = model.UpdateOrderByID(db, id, confirmReq, current); err != nil {
			log.Printf("An error occurred while updating the order data: %v\n", err)
			c.JSON(400, gin.H{"error": "Product order not valid"})
			return

		}

		order.Passcode = nil

		order.PaidAt = &current
		order.PaidBank = &confirmReq.Bank
		order.PaidAccountNumber = &confirmReq.AccountNumber

		c.JSON(200, order)

	}
}

func GetOrder(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
