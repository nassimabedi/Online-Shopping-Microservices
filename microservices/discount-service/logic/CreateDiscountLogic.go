package logic

import (
	"Online-Shopping-Microservices/microservices/discount-service/domain"
	"Online-Shopping-Microservices/microservices/discount-service/repository"
	"context"
	"time"
)

// CreateDiscountInterface
type CreateDiscountInterface interface {
	CreateNewDiscount(PhoneNumber string, DiscountNumber string) (*domain.DiscountInfo, error)
}

// CreateDiscountLogic struct
type CreateDiscountLogic struct {
	Context            context.Context
	Self               CreateDiscountInterface
	PublishLogic       PublishInterface
	CreateDiscountRepo repository.CreateDiscountInterface
}

// NewCreateDiscount
func NewCreateDiscount(ctx context.Context) CreateDiscountInterface {
	x := &CreateDiscountLogic{Context: ctx}
	x.Self = x
	return x
}

//create new discount
func (c *CreateDiscountLogic) CreateNewDiscount(PhoneNumber string, DiscountNumber string) (*domain.DiscountInfo, error) {

	if c.CreateDiscountRepo == nil {
		c.CreateDiscountRepo = repository.NewCreateDiscount(c.Context)
	}

	var Discount domain.DiscountInfo

	Discount.Number = DiscountNumber
	Discount.PhoneNumber = PhoneNumber
	Discount.CreatedAt = time.Now()

	//TODO: check key exit

	//insert in redis
	_, err := c.CreateDiscountRepo.InsertRedisDiscount(Discount)
	if err != nil {
		return nil, err
	}

	//insert in mongodb
	result, err := c.CreateDiscountRepo.CreateDiscount(Discount)
	if err != nil {
		return nil, err
	}
	return result, nil
}
