package httpEngine

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Run(Port string) {
	engine := gin.Default()
	v1 := engine.Group("v1/discounts")

	discountController := NewDiscountController()
	{
		v1.POST("/discounts", discountController.NewDiscount)
		v1.GET("/win/user", discountController.WinningUser)

	}

	engine.Use(gin.Recovery())

	fmt.Println(engine.Run(fmt.Sprintf(":%s", Port)))
}
