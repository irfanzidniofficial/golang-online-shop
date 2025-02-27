package main

import (
	"database/sql"
	"fmt"
	"golang-online-shop/handler"
	"golang-online-shop/middleware"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	db, err := sql.Open("pgx", os.Getenv("DB_URI"))

	if err != nil {
		fmt.Printf("Failed to connected to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		fmt.Printf("Failed to verify database: %v\n", err)
		os.Exit(1)
	}

	if _, err := migrate(db); err != nil {
		fmt.Printf("Failed to migrate database: %v\n", err)
		os.Exit(1)
	}

	r := gin.Default()

	r.GET("/api/v1/products", handler.ListProducts(db))
	r.GET("/api/v1/products/:id", handler.GetProducts(db))
	r.POST("/api/v1/checkout", handler.ChekoutOrder(db))

	r.POST("/api/v1/orders/:id/confirm")
	r.GET("/api/v1/orders/:id")

	r.POST("/admin/products", middleware.AdminOnly(), handler.CreateProduct(db))
	r.PUT("/admin/products/:id", middleware.AdminOnly(), handler.UpdateProduct(db))
	r.DELETE("/admin/products/:id", middleware.AdminOnly(), handler.DeleteProduct(db))

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	if err = server.ListenAndServe(); err != nil {
		fmt.Printf("failed to start the server: %v\n", err)
		os.Exit(1)

	}

}
