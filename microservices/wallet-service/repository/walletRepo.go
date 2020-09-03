package repository

import (
	"Online-Shopping-Microservices/microservices/wallet-service/constant"
	"Online-Shopping-Microservices/microservices/wallet-service/domain"
	"context"
	"errors"
	"github.com/RezaOptic/go-utils/logger"
	"gopkg.in/mgo.v2/bson"
)

// WalletRepoInterface
type WalletRepoInterface interface {
	AddCredit(credit domain.WalletInfo) (*domain.WalletInfo, error)
	GetWalletInfo(PhoneNumber string) (*domain.WalletInfo, error)
}

// WalletRepo
type WalletRepo struct {
	Context context.Context
	Self    WalletRepoInterface
}

func NewWalletRepo(ctx context.Context) WalletRepoInterface {
	return WalletRepo{
		Context: ctx,
	}
}

// PopAuthOtpFromQueue method for listen on auth otp event queue and return new data
func (q WalletRepo) AddCredit(wallet domain.WalletInfo) (*domain.WalletInfo, error) {
	err := DBS.MongoDB.DB(constant.DBMongoName).C(constant.WalletDocument).Insert(wallet)
	if err != nil {
		logger.ZSLogger.Errorf("error on insert wallet with error :%s", err)
		return nil, errors.New(constant.InsertWalletError)
	}

	return &wallet, nil
}

// get wallet info from mongodb
func (q WalletRepo) GetWalletInfo(PhoneNumber string) (*domain.WalletInfo, error) {
	var walletInfo domain.WalletInfo

	err := DBS.MongoDB.DB(constant.DBMongoName).C(constant.WalletDocument).Find(bson.M{"phone_number": PhoneNumber}).One(&walletInfo)
	if err != nil {
		logger.ZSLogger.Errorf("error on insert wallet with error :%s", err)
		return nil, errors.New(constant.InsertWalletError)
	}

	return &walletInfo, nil
}
