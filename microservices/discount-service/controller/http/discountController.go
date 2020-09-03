package httpEngine

import (
	"Online-Shopping-Microservices/microservices/discount-service/domain"
	"Online-Shopping-Microservices/microservices/discount-service/logic"
	"github.com/RezaOptic/go-utils/errorsHandler"
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
		errorsHandler.GinErrorResponseHandler(c, err)
		return
	}

	discountNumber := discountInfo.Number
	phoneNumber := discountInfo.PhoneNumber

	result, err := logic.NewCreateDiscount(c).CreateNewDiscount(phoneNumber, discountNumber)
	if err != nil {
		errorsHandler.GinErrorResponseHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated,
		result,
	)
	return
}
