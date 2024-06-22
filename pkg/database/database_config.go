package database

import (
	"fmt"

	"github.com/thiagoluis88git/tech1/internal/adapters/driven/entities"
	"github.com/thiagoluis88git/tech1/pkg/environment"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConfigDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		environment.GetDBHost(),
		environment.GetDBUser(),
		environment.GetDBPassword(),
		environment.GetDBName(),
		environment.GetDBPort(),
	)
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
		&entities.ComboProduct{},
		&entities.OrderTicketNumber{},
	)

	return db
}
