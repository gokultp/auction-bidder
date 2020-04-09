package model

import (
	"context"

	"github.com/jinzhu/gorm"
)

// Bid defines the Bid model
type Bid struct {
	gorm.Model
	Price     *uint
	UserID    *uint
	AuctionID *uint
}

func (u *Bid) Create(ctx context.Context) error {
	db := ctx.Value("db").(*gorm.DB)
	return db.Create(u).Error
}

func (u *Bid) Update(ctx context.Context) error {
	db := ctx.Value("db").(*gorm.DB)
	return db.Update(u).Error
}

func GetBidByID(ctx context.Context, id uint, auctionID uint) (*Bid, error) {
	var bid Bid
	db := ctx.Value("db").(*gorm.DB)
	if err := db.Where(Bid{AuctionID: &auctionID}).First(&bid, id).Error; err != nil {
		return nil, err
	}
	return &bid, nil
}

func GetBidsByAuction(ctx context.Context, auctionID, limit, page uint) ([]Bid, error) {
	var bids []Bid
	offset := (page - 1) * limit
	db := ctx.Value("db").(*gorm.DB)
	if err := db.Where(Bid{AuctionID: &auctionID}).
		Offset(offset).
		Limit(limit).
		Find(&bids).Error; err != nil {
		return nil, err
	}
	return bids, nil
}
