package db

import (
	"github.com/gokultp/auction-bidder/internal/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=postgres port=5432 user=postgresadmin dbname=postgresdb password=admin123 sslmode=disable")
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Bid{})
	db.AutoMigrate(&model.Auction{})
	db.AutoMigrate(&model.Event{})

	return db, nil
}
