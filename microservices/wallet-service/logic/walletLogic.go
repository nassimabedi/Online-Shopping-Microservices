package logic

import (
	"Online-Shopping-Microservices/microservices/wallet-service/domain"
	"Online-Shopping-Microservices/microservices/wallet-service/repository"
	"context"
	"time"
)

// CreateWalletInterface
type CreateWalletInterface interface {
	CreateNewCredit(PhoneNumber string, credit int) (*domain.WalletInfo, error)
	WalletInfo(PhoneNumber string) (*domain.WalletInfo, error)
}

// CreateWalletLogic struct
type CreateWalletLogic struct {
	Context          context.Context
	Self             CreateWalletInterface
	PublishLogic     PublishInterface
	CreateWalletRepo repository.WalletRepoInterface
	CreateQueueRepo  repository.QueueRepoInterface
}

// NewCreateWallet
func NewCreateWallet(ctx context.Context) CreateWalletInterface {
	x := &CreateWalletLogic{Context: ctx}
	x.Self = x
	return x
}

//create new credit to customer wallet
func (c *CreateWalletLogic) CreateNewCredit(PhoneNumber string, credit int) (*domain.WalletInfo, error) {

	if c.CreateWalletRepo == nil {
		c.CreateWalletRepo = repository.NewWalletRepo(c.Context)
	}

	var Wallet domain.WalletInfo
	Wallet.PhoneNumber = PhoneNumber
	Wallet.Credit = credit
	Wallet.CreatedAt = time.Now()

	result, err := c.CreateWalletRepo.AddCredit(Wallet)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//Get customer Wallet info by phone number
func (c *CreateWalletLogic) WalletInfo(PhoneNumber string) (*domain.WalletInfo, error) {

	if c.CreateWalletRepo == nil {
		c.CreateWalletRepo = repository.NewWalletRepo(c.Context)
	}

	result, err := c.CreateWalletRepo.GetWalletInfo(PhoneNumber)
	if err != nil {
		return nil, err
	}

	return result, nil
}
