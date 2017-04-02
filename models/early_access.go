package models

import (
	"time"
)

type EarlyAccess struct {
	Email string `valid:"email" gorethink:"email"`
	CreatedAt time.Time `gorethink:"createdAt"`
}