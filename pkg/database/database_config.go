package database

import (
	"fmt"

	"github.com/thiagoluis88git/tech1/internal/adapters/driven/model"
	"github.com/thiagoluis88git/tech1/pkg/environment"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConfigDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v",
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
		&model.Customer{},
		&model.Order{},
		&model.OrderProduct{},
		&model.Payment{},
		&model.Product{},
		&model.ProductImage{},
		&model.ComboProduct{},
		&model.OrderTicketNumber{},
	)

	return db
}
