package model

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"
)

// Auction defines the Auction model
type Auction struct {
	gorm.Model
	Name          *string
	Description   *string
	StartAt       *time.Time
	EndAt         *time.Time
	StartPrice    *uint
	AuctionWinner *uint
}

// Init will create table auto index and create custom indexed if needed
func (u *Auction) Init(db *gorm.DB) {
	db.AutoMigrate(u)
}

func (u *Auction) Create(ctx context.Context) error {
	db := ctx.Value("db").(*gorm.DB)
	return db.Create(u).Error
}

func (u *Auction) Update(ctx context.Context) error {
	db := ctx.Value("db").(*gorm.DB)
	return db.Update(u).Error
}

func GetAuctionByID(ctx context.Context, id uint) (*Auction, error) {
	var auction Auction
	db := ctx.Value("db").(*gorm.DB)
	if err := db.First(&auction, id).Error; err != nil {
		return nil, err
	}
	return &auction, nil
}

func GetAuctions(ctx context.Context, limit, offset uint) ([]Auction, error) {
	var auctions []Auction
	db := ctx.Value("db").(*gorm.DB)
	if err := db.Offset(offset).Limit(limit).Find(&auctions).Error; err != nil {
		return nil, err
	}
	return auctions, nil
}
