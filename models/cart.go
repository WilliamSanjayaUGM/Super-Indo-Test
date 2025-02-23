package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserId     string    `json:"userId"`
	TotalPrice int       `json:"totalPrice"`
	TotalItem  int       `json:"totalItem"`
	ListItem   []Product `json:"listProduct"`
}

func AddToCart(userId string, productId string) *Cart {
	var userCart Cart
	var product Product
	DB.Where("user_id = ?", userId).First(&userCart)
	DB.Where("id = ?", productId).First(&product)

	DB.Where("user_id = ?", userId).First(&userCart)

	if len(userCart.ListItem) == 0 {
		userCart.UserId = userId
		userCart.TotalPrice = product.Price
		userCart.TotalItem = 1
		userCart.ListItem = append(userCart.ListItem, product)
	} else {
		userCart.TotalPrice = userCart.TotalPrice + product.Price
		userCart.TotalItem = userCart.TotalItem + 1
		userCart.ListItem = append(userCart.ListItem, product)
	}
	if DB.Model(&userCart).Where("user_id = ?", userId).Updates(&userCart).RowsAffected == 0 {
		DB.Create(&userCart)
	}
	return &userCart
}

// func DeletCartByProductId(userId string, productId string) *Cart{

// }
