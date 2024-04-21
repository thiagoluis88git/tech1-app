package database

import (
	"fmt"
	"thiagoluis88git/tech1/internal/adapters/driven/entities"
	"thiagoluis88git/tech1/pkg/environment"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConfigDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=fastfood password=fastfood1234 dbname=fastfood_db port=5432 sslmode=disable", *environment.DbHost)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("could not open database: %v", err.Error()))
	}

	db.AutoMigrate(
		&entities.Customer{},
		&entities.Order{},
		&entities.OrderProduct{},
		&entities.Payment{},
		&entities.Product{},
		&entities.ProductImage{},
		&entities.ProductCombo{},
	)

	return db
}
