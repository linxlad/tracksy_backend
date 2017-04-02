package models

import "time"

type EarlyAccess struct {
	Email string `gorethink:"email"`
	CreatedAt time.Time `gorethink:"createdAt"`
}