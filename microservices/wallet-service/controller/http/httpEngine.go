package httpEngine

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Run(Port string) {
	engine := gin.Default()
	v1 := engine.Group("v1/wallet")

	walletController := NewWalletController()
	{
		v1.GET("/info/:phone_number", walletController.WalletInfo)
		v1.GET("/win/user", walletController.WinningUser)

	}

	engine.Use(gin.Recovery())

	fmt.Println(engine.Run(fmt.Sprintf(":%s", Port)))
}
