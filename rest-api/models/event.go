package models

import (
	"time"
)

type Event struct {
	Id          int64     `json:"-"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	Date        time.Time `json:"date" binding:"required"`
	UserId      int64     `json:"-"`
}
