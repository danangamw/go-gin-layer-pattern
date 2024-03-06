package entity

import "time"

type Post struct {
	ID         int
	UserID     int
	Tweet      string  // default value == ""
	PictureUrl *string // default value == nil
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
