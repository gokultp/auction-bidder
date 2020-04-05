package contract

import "time"

type User struct {
	ID        uint       `json:"id"`
	FirstName *string    `json:"first_name"`
	LastName  *string    `json:"last_name"`
	Email     *string    `json:"email"`
	Token     *string    `json:"token"`
	IsAdmin   *bool      `json:"is_admin"`
	IsActive  *bool      `json:"is_active"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type UserResponse struct {
	Meta *Metadata `json:"metadata"`
	Data *User     `json:"data"`
}

type MultiUserResponse struct {
	Meta *Metadata `json:"metadata"`
	Data []User    `json:"data"`
}
