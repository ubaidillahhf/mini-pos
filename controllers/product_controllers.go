package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"pos/lib/database"
	"pos/models/products"
	"pos/validations"
	"strconv"

	"github.com/labstack/echo/v4"
)

func UploadImage(c echo.Context) (res string, err error) {

	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("image")
	if err != nil {
		return "error", err
	}
	src, err := file.Open()
	if err != nil {
		return "error", err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(file.Filename)
	if err != nil {
		return "error", err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return "error", err
	}

	return file.Filename, nil
}

/* PRODUCT */
func CreateProductControllers(c echo.Context) error {

	pathImage, _ := UploadImage(c)

	fmt.Println("destination =>", pathImage)

	var productCreate products.ProductPost
	merchant_id, _ := strconv.Atoi(c.FormValue("merchant_id"))

	productCreate.Name = c.FormValue("name")
	productCreate.MerchantId = merchant_id
	productCreate.Sku = c.FormValue("sku")
	productCreate.Remark = c.FormValue("remark")
	productCreate.Description = c.FormValue("description")
	productCreate.Image = pathImage

	// Validasi Field
	errorValidate := validations.Validate(productCreate)
	if errorValidate != nil {
		return errorValidate
	}

	createProductDB, err := database.CreateProduct(productCreate)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, BaseResponse(
		http.StatusCreated,
		"Success Create Product",
		createProductDB,
	))
}

func DetailProductControllers(c echo.Context) error {
	paramsProductId := c.Param("productId")
	productId, _ := strconv.Atoi(paramsProductId)

	categoryDB, e := database.GetProductDetail(productId)

	if e != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Get Data",
			nil,
		))
	}
	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Get Data by productId",
		categoryDB,
	))
}

func GetProductControllers(c echo.Context) error {

	var productData []products.Product
	var err error
	productData, err = database.GetProductAll()

	if err != nil {
		return c.JSON(http.StatusOK, BaseResponse(
			http.StatusInternalServerError,
			"Failed Get Data",
			productData,
		))
	}

	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Get Data Categories",
		productData,
	))
}

func EditProductControllers(c echo.Context) error {

	paramsProductId := c.Param("productId")
	categoryId, _ := strconv.Atoi(paramsProductId)
	var productEditDate products.ProductPost
	c.Bind(&productEditDate)

	// Validasi Field
	errorValidate := validations.Validate(productEditDate)
	if errorValidate != nil {
		return errorValidate
	}

	userEdit, err := database.EditProduct(productEditDate, categoryId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, BaseResponse(http.StatusOK, "Sukses Edit Produk", userEdit))
}

func DeleteProductControllers(c echo.Context) error {

	paramsProductId := c.Param("productId")
	productId, _ := strconv.Atoi(paramsProductId)
	_, e := database.DeleteProduct(productId)

	if e != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Get Data",
			nil,
		))
	}
	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Delete Category",
		nil,
	))
}
