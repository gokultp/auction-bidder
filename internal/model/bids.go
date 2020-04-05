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

// Init will create table auto index and create custom indexed if needed
func (u *Bid) Init(db *gorm.DB) {
	db.AutoMigrate(u)
}

func (u *Bid) Create(ctx context.Context) error {
	db := ctx.Value("db").(*gorm.DB)
	return db.Create(u).Error
}

func (u *Bid) Update(ctx context.Context) error {
	db := ctx.Value("db").(*gorm.DB)
	return db.Update(u).Error
}

func GetBidByID(ctx context.Context, id uint) (*Bid, error) {
	var Bid Bid
	db := ctx.Value("db").(*gorm.DB)
	if err := db.First(&Bid, id).Error; err != nil {
		return nil, err
	}
	return &Bid, nil
}

func GetBidsByAuction(ctx context.Context, auctionID, limit, offset uint) ([]Bid, error) {
	var Bids []Bid
	db := ctx.Value("db").(*gorm.DB)
	if err := db.Where(Bid{AuctionID: &auctionID}).
		Offset(offset).
		Limit(limit).
		Find(&Bids).Error; err != nil {
		return nil, err
	}
	return Bids, nil
}
