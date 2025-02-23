package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type Product struct {
	gorm.Model
	Id          string `json:"id"`
	Productname string `json:"productName"`
	BrandName   string `json:"brandName"`
	Price       int    `json:"price"`
	CategoryId  string `json:"categoryId"`
	Status      string `json:"status"`
}

var Categories = []Category{
	{Id: "1", Name: "Fresh Food", Description: "", Status: "Y"},
	{Id: "2", Name: "Process & Packaged Food", Description: "", Status: "Y"},
	{Id: "3", Name: "Household Goods", Description: "", Status: "Y"},
	{Id: "4", Name: "Clothing & Personal Equipment", Description: "", Status: "Y"},
	{Id: "5", Name: "Electronic & Large House Equipment", Description: "", Status: "Y"},
	{Id: "6", Name: "Baby & Children Equipment", Description: "", Status: "Y"},
	{Id: "7", Name: "Frozen Food", Description: "", Status: "Y"},
	{Id: "8", Name: "Stationary & Office Supplies", Description: "", Status: "Y"},
	{Id: "9", Name: "Beauty & Personal Care", Description: "", Status: "Y"},
	{Id: "10", Name: "Medicines", Description: "", Status: "Y"},
	{Id: "11", Name: "Pet Supplies", Description: "", Status: "Y"},
	{Id: "12", Name: "Alcohol & Cigarettes", Description: "", Status: "Y"},
}

var Products = []Product{
	{Id: "1", Productname: "Fresh Tomat", BrandName: "Super Indo", Price: 2000, CategoryId: "1", Status: "Y"},
	{Id: "2", Productname: "Fresh Fish", BrandName: "Super Indo", Price: 300, CategoryId: "1", Status: "Y"},
	{Id: "3", Productname: "Fresh Potato", BrandName: "Super Indo", Price: 2000, CategoryId: "1", Status: "Y"},
	{Id: "4", Productname: "Sawi", BrandName: "Super Indo", Price: 1000, CategoryId: "1", Status: "Y"},
	{Id: "5", Productname: "Daging Maling", BrandName: "Super Indo", Price: 10000, CategoryId: "2", Status: "Y"},
	{Id: "6", Productname: "Crackers", BrandName: "Nissin", Price: 20000, CategoryId: "2", Status: "Y"},
	{Id: "7", Productname: "Sarden Kaleng", BrandName: "Del Monte", Price: 25000, CategoryId: "2", Status: "Y"},
	{Id: "8", Productname: "Sapu", BrandName: "Super Indo", Price: 50000, CategoryId: "3", Status: "Y"},
	{Id: "9", Productname: "Pel", BrandName: "Super Indo", Price: 50000, CategoryId: "3", Status: "Y"},
	{Id: "10", Productname: "Pengki", BrandName: "Super Indo", Price: 50000, CategoryId: "3", Status: "Y"},
	{Id: "11", Productname: "Torent Air", BrandName: "Super Indo", Price: 100000, CategoryId: "3", Status: "Y"},
	{Id: "12", Productname: "Kemeja", BrandName: "H&M", Price: 200000, CategoryId: "4", Status: "Y"},
	{Id: "13", Productname: "Jaket", BrandName: "H&M", Price: 300000, CategoryId: "4", Status: "Y"},
	{Id: "14", Productname: "Kabel Roll", BrandName: "Super Indo", Price: 400000, CategoryId: "5", Status: "Y"},
	{Id: "15", Productname: "Terminal T", BrandName: "Super Indo", Price: 600000, CategoryId: "5", Status: "Y"},
}

func CategoriesAll() *[]Category {
	var listCategory []Category
	DB.Find(&listCategory)
	return &listCategory
}

func ProductsAll(categoryId string) *[]Product {
	var listProducts []Product
	DB.Where("category_id = ?", categoryId).Find(&listProducts)
	return &listProducts
}

func ProductById(productId string) *Product {
	var product Product
	DB.Where("id = ?", productId).First(&product)
	return &product
}
