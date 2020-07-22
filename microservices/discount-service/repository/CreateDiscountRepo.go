package repository

import (
	"arvan.ir/app-services/discount-service/constant"
	"arvan.ir/app-services/discount-service/domain"
	"context"
	"errors"
	"fmt"
)

// CreateDiscountInterface
type CreateDiscountInterface interface {
	CreateDiscount(Discount domain.DiscountInfo) (*domain.DiscountInfo, error)
	InsertRedisDiscount(Discount domain.DiscountInfo) (*domain.DiscountInfo, error)
}

// CreateDiscountRepo struct
type CreateDiscountRepo struct {
	Context context.Context
	Self    CreateDiscountInterface
}

// NewCreateDiscount
func NewCreateDiscount(ctx context.Context) CreateDiscountInterface {
	x := &CreateDiscountRepo{Context: ctx}
	x.Self = x
	return x
}

// CreateDiscount method for insert new Discount info
func (c *CreateDiscountRepo) CreateDiscount(Discount domain.DiscountInfo) (*domain.DiscountInfo, error) {

	err := DBS.MongoDB.DB(constant.DBMongoName).C(constant.DiscountDocument).Insert(Discount)
	if err != nil {
		fmt.Printf("error on insert discount with error :%s", err)
		return nil, errors.New(constant.InsertDiscountError)
	}

	return &Discount, nil

}

//push phone number to queue
func (c *CreateDiscountRepo) InsertRedisDiscount(Discount domain.DiscountInfo) (*domain.DiscountInfo, error) {
	_, err := DBS.Redis.RPush(constant.UsersKey, Discount.PhoneNumber).Result()

	if err != nil {
		fmt.Printf("error on insert discount with error in redis :%v \n", err)
		return nil, errors.New(constant.InsertDiscountError)
	}

	return &Discount, nil

}
