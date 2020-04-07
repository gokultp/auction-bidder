package model

import (
	"context"

	"github.com/jinzhu/gorm"
)

// User defines the user model
type User struct {
	gorm.Model
	FirstName *string
	LastName  *string
	Email     *string `gorm:"unique_index"`
	Token     *string
	IsAdmin   *bool
	IsActive  *bool
}

// Init will create table auto index and create custom indexed if needed
func (u *User) Init(db *gorm.DB) {
	db.AutoMigrate(u)
}

func (u *User) Create(ctx context.Context) error {
	db := ctx.Value("db").(*gorm.DB)
	return db.Create(u).Error
}

func (u *User) Update(ctx context.Context) error {
	db := ctx.Value("db").(*gorm.DB)
	return db.Save(u).Error
}

func GetUserByID(ctx context.Context, id uint) (*User, error) {
	var user User
	db := ctx.Value("db").(*gorm.DB)
	if err := db.First(&user, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	db := ctx.Value("db").(*gorm.DB)
	if err := db.Where(User{Email: &email}).
		First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func GetUsers(ctx context.Context, limit, offset uint) ([]User, error) {
	var users []User
	db := ctx.Value("db").(*gorm.DB)
	if err := db.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
