package database

import (
	"pos/configs"
	"pos/models/products"
)

func CreateProduct(productCreate products.ProductPost) (products.Product, error) {
	var productDB products.Product

	productDB.Name = productCreate.Name
	productDB.Image = productCreate.Image
	productDB.MerchantId = productCreate.MerchantId
	productDB.Sku = productCreate.Sku
	productDB.Remark = productCreate.Remark
	productDB.Description = productCreate.Description

	err := configs.DB.Create(&productDB).Error
	if err != nil {
		return productDB, err
	}
	return productDB, nil
}

func GetProductAll() (dataResult []products.Product, err error) {
	err = configs.DB.Find(&dataResult).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetProductDetail(productId int) (products.Product, error) {
	var productDB products.Product
	err := configs.DB.First(&productDB, productId).Error

	if err != nil {
		return productDB, err
	}
	return productDB, nil
}

func EditProduct(productEdit products.ProductPost, productId int) (products.Product, error) {
	var productDB products.Product
	err := configs.DB.First(&productDB, productId).Error

	productDB.Name = productEdit.Name
	productDB.MerchantId = productEdit.MerchantId
	productDB.Sku = productEdit.Sku
	productDB.Remark = productEdit.Remark
	productDB.Description = productEdit.Description

	configs.DB.Save(productDB)

	if err != nil {
		return productDB, err
	}

	return productDB, nil
}

func DeleteProduct(productId int) (products.Product, error) {
	var productDB products.Product
	err := configs.DB.Where("id = ?", productId).Delete(&productDB).Error

	if err != nil {
		return productDB, err
	}
	return productDB, nil
}
