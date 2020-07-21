package domain

import (
	"time"
)

// DiscountInfo struct method for holding discount info
type DiscountInfo struct {
	Id                              string      `json:"discount_id"`
	Number                       string      `json:"number"`
	PhoneNumber                       string      `json:"phone_number"`
	CreatedAt                       time.Time `json:"created_at"`

}

