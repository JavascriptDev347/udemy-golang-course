package models

import "time"

// User represents a user in the system
type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
	Profile        Profile   `json:"profile"`
}

// Profile represents a user's profile information
type Profile struct {
	UserID  int       `json:"user_id"`
	Avatar  string    `json:"avatar"`
	Created time.Time `json:"created"`
}
