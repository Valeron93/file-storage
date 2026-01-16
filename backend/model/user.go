package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID
	Username       string
	HashedPassword []byte
	CreatedAt      time.Time
}

type Session struct {
	ID        uuid.UUID
	User      *User
	CreatedAt time.Time
}
