package httpEngine

import (
	"arvan.ir/app-services/discount-service/domain"
	"arvan.ir/app-services/discount-service/logic"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DiscountControllerInterface
type DiscountControllerInterface interface {
	NewDiscount(c *gin.Context)
}

//
type DiscountControllerStruct struct {
	Self DiscountControllerInterface
}

// NewDiscountController
func NewDiscountController() DiscountControllerInterface {
	x := &DiscountControllerStruct{}
	x.Self = x
	return x
}

// Create discount from logic
func (d *DiscountControllerStruct) NewDiscount(c *gin.Context) {
	var discountInfo domain.DiscountInfo
	if err := c.Bind(&discountInfo); err != nil {
		fmt.Printf("%v \n", err)
		return
	}

	discountNumber := discountInfo.Number
	phoneNumber := discountInfo.PhoneNumber

	result, err := logic.NewCreateDiscount(c).CreateNewDiscount(phoneNumber, discountNumber)
	if err != nil {
		fmt.Printf("%v \n", err)
		return
	}

	c.JSON(http.StatusCreated,
		result,
	)
	return
}
