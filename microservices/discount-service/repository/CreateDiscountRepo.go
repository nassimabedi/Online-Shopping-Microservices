package repository

import (
	"Online-Shopping-Microservices/microservices/discount-service/constant"
	"Online-Shopping-Microservices/microservices/discount-service/domain"
	"context"
	"errors"
	"github.com/RezaOptic/go-utils/logger"
)

// CreateDiscountInterface
type CreateDiscountInterface interface {
	CreateDiscount(Discount domain.DiscountInfo) (*domain.DiscountInfo, error)
	InsertRedisDiscount(Discount domain.DiscountInfo) (*domain.DiscountInfo, error)
	GetWinningUsers() ([]string, error)
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
		logger.ZSLogger.Errorf("error on insert discount with error :%s", err)
		return nil, errors.New(constant.InsertDiscountError)
	}

	return &Discount, nil

}

//push phone number to queue
func (c *CreateDiscountRepo) InsertRedisDiscount(Discount domain.DiscountInfo) (*domain.DiscountInfo, error) {
	_, err := DBS.Redis.RPush(constant.UsersKey, Discount.PhoneNumber).Result()

	if err != nil {
		logger.ZSLogger.Errorf("error on insert discount with error in redis :%s", err)
		return nil, errors.New(constant.InsertDiscountError)
	}

	return &Discount, nil

}

//insert into hash for get real time user
func (q *CreateDiscountRepo) GetWinningUsers() ([]string, error) {
	var result []string

	result, err := DBS.Redis.LRange(constant.RealTimeKey, constant.WinningUserStart, constant.WinningUserEnd).Result()

	if err != nil {
		if err == redis.Nil {
			logger.ZSLogger.Errorf("error on push  to hash for winning users with error :%s", err)
			return result, nil
		}
		return result, err
	}
	return result, nil
}
