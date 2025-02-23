package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserId      string        `json:"userId"`
	TotalPrice  int           `json:"totalPrice"`
	TotalItem   int           `json:"totalItem"`
	ListProduct []ProductCart `json:"listProduct"`
}

type ProductCart struct {
	gorm.Model
	ProductId string `json:"productId"`
	Price     int    `json:"price"`
	UserId    string `json:"userId"`
}

func AddToCart(userId string, productId string) *Cart {
	var userCart Cart
	var product Product
	var productCart ProductCart
	DB.Where("user_id = ?", userId).First(&userCart)
	DB.Where("id = ?", productId).First(&product)

	productCart.UserId = userId
	productCart.ProductId = product.Id
	productCart.Price = product.Price

	if len(userCart.ListProduct) == 0 {
		userCart.UserId = userId
		userCart.TotalPrice = product.Price
		userCart.TotalItem = 1
		userCart.ListProduct = append(userCart.ListProduct, productCart)
	} else {
		userCart.TotalPrice = userCart.TotalPrice + product.Price
		userCart.TotalItem = userCart.TotalItem + 1
		userCart.ListProduct = append(userCart.ListProduct, productCart)
	}
	if DB.Model(&userCart).Where("user_id = ?", userId).Updates(&userCart).RowsAffected == 0 {
		DB.Create(&userCart)
	}
	return &userCart
}

func DeletCartByProductId(userId string, productId string) *Cart {
	var userCart Cart
	var productCart ProductCart
	var singleProductCart ProductCart
	DB.Where("user_id = ?", userId).First(&userCart)
	DB.Where("user_id = ?", userId).Where("product_id", productId).First(&singleProductCart)
	productPrice := singleProductCart.Price

	if len(userCart.ListProduct) == 1 {
		DB.Where("user_id = ?", userId).Delete(&userCart)
	} else {
		DB.Where("user_id = ?", userId).Where("product_id", productId).Limit(1).Delete(&productCart)
		userCart.TotalPrice = userCart.TotalPrice - productPrice
		userCart.TotalItem = userCart.TotalItem - 1
		userCart.ListProduct = append(userCart.ListProduct, productCart)
		DB.Where("user_id = ?", userId).Updates(&userCart)
	}
	return &userCart
}
