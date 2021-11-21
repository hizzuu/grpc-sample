package domain

import "time"

type User struct {
	ID                int64
	Email             string
	Name              string
	EncryptedPassword string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}