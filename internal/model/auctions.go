package model

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"
)

var (
	AuctionStatusActive = "active"
	AuctionStatusClosed = "closed"
)

// Auction defines the Auction model
type Auction struct {
	gorm.Model
	Name          *string
	Description   *string
	StartAt       *time.Time
	EndAt         *time.Time
	Status        *string
	StartPrice    *uint
	AuctionWinner *uint
	CreatedBy     *uint
}

func (u *Auction) Create(ctx context.Context) error {
	db := ctx.Value("db").(*gorm.DB)
	return db.Create(u).Error
}

func (u *Auction) Update(ctx context.Context) error {
	db := ctx.Value("db").(*gorm.DB)
	return db.Save(u).Error
}

func GetAuctionByID(ctx context.Context, id uint) (*Auction, error) {
	var auction Auction
	db := ctx.Value("db").(*gorm.DB)
	if err := db.First(&auction, id).Error; err != nil {
		return nil, err
	}
	return &auction, nil
}

func GetAuctions(ctx context.Context, limit, page uint) ([]Auction, error) {
	var auctions []Auction
	offset := (page - 1) * limit
	db := ctx.Value("db").(*gorm.DB)
	if err := db.Offset(offset).Limit(limit).Find(&auctions).Error; err != nil {
		return nil, err
	}
	return auctions, nil
}
