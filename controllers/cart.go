package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
)

var (
	ErrCantFindProduct    = errors.New("can't find the product")
	ErrCantDecodeProduct  = errors.New("can't find the product")
	ErrUserIdIsNotValid   = errors.New("this user is not valid")
	ErrCantUpdateUser     = errors.New("cannot add this product to the cart")
	ErrCantRemoveItemCart = errors.New("cannot remove this item from the cart")
	ErrCantGetItem        = errors.New("was unable to get the item from the cart")
	ErrCantBuyCartItem    = errors.New("cannot update the purchase")
)

func AddToCart() gin.HandlerFunc {

}

func RemoveFromCart() gin.HandlerFunc {

}

func GetItemFromCart() gin.HandlerFunc {

}

func BuyItemFromCart() gin.HandlerFunc {

}

func Instantiate() gin.HandlerFunc {

}
