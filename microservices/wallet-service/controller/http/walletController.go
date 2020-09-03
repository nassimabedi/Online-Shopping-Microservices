package httpEngine

import (
	"Online-Shopping-Microservices/microservices/wallet-service/logic"
	"github.com/RezaOptic/go-utils/errorsHandler"
	"github.com/gin-gonic/gin"
	"net/http"
)

// WalletControllerInterface
type WalletControllerInterface interface {
	WalletInfo(c *gin.Context)
	WinningUser(c *gin.Context)
}

//WalletControllerStruct
type WalletControllerStruct struct {
	Self WalletControllerInterface
}

// NewWalletController
func NewWalletController() WalletControllerInterface {
	x := &WalletControllerStruct{}
	x.Self = x
	return x
}

// get wallet info from logic
func (d *WalletControllerStruct) WalletInfo(c *gin.Context) {
	phoneNumber := c.Param("phone_number")

	result, err := logic.NewCreateWallet(c).WalletInfo(phoneNumber)
	if err != nil {
		errorsHandler.GinErrorResponseHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated,
		result,
	)
	return
}

//get winning users
func (d *WalletControllerStruct) WinningUser(c *gin.Context) {
	result, err := logic.NewCreateWallet(c).WinningUser()
	if err != nil {
		errorsHandler.GinErrorResponseHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated,
		result,
	)
	return

}
