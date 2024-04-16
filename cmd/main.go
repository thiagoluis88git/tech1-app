package main

import (
	"fmt"
	"net/http"
	"thiagoluis88git/tech1/internal/adapters/driven/entities"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=database user=fastfood password=fastfood1234 dbname=fastfood_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("could not open database: %v", err.Error()))
	}

	db.AutoMigrate(
		&entities.Customer{},
		&entities.Order{},
		&entities.OrderProduct{},
		&entities.PaymentOutbox{},
		&entities.Product{},
		&entities.ProducImage{},
		&entities.ProductCombo{},
	)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3210", r)
}
