package domain

import (
	"time"
)

// DiscountInfo struct method for holding discount info
type WalletInfo struct {
	Id                              string      `json:"discount_id"`
	PhoneNumber                     string      `json:"phone_number"`
	Credit                          int         `json:"credit"`
	CreatedAt                       time.Time `json:"created_at"`

}

