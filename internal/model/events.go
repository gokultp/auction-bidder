package model

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"
)

var (
	EventStatusScheduled = "scheduled"
	EventStatusQueued    = "queued"
	EventStatusCompleted = "completed"
	EventStatusFailed    = "failed"
)

// Event defines the Event model
type Event struct {
	gorm.Model
	Time   *time.Time
	Data   *string
	Status *string `gorm:"default:'scheduled'"`
	Type   *string
}

func (u *Event) Create(ctx context.Context) error {
	db := ctx.Value("db").(*gorm.DB)
	return db.Create(u).Error
}

func GetEventsForNextBatch(ctx context.Context) ([]Event, error) {
	var Events []Event
	db := ctx.Value("db").(*gorm.DB)
	now := time.Now().Add(5 * time.Minute)
	if err := db.
		Where("time <= ? and status =?", now, EventStatusScheduled).
		Find(&Events).Error; err != nil {
		return nil, err
	}
	return Events, nil
}

func GetEventByID(ctx context.Context, id uint) (*Event, error) {
	var event Event
	db := ctx.Value("db").(*gorm.DB)
	if err := db.First(&event, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &event, nil
}

func (e *Event) Update(ctx context.Context) error {
	db := ctx.Value("db").(*gorm.DB)
	return db.Save(e).Error
}
