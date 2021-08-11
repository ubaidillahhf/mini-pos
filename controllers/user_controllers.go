package controllers

import (
	"net/http"
	"pos/lib/database"
	"pos/middlewares"
	"pos/models/users"
	"pos/validations"
	"strconv"

	"github.com/labstack/echo/v4"
)

func RegisterControllers(c echo.Context) error {

	var usersCreate users.UserCreate
	c.Bind(&usersCreate)

	// Validasi Field
	errorValidate := validations.Validate(usersCreate)
	if errorValidate != nil {
		return errorValidate
	}

	userDB, err := database.RegisterUser(usersCreate)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, BaseResponse(
		http.StatusCreated,
		"Success Register User",
		userDB,
	))
}

func LoginControllers(c echo.Context) error {

	userLogin := users.UserLogin{}
	c.Bind(&userLogin)

	userDB, e := database.LoginUser(userLogin)
	if e != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Get Data",
			nil,
		))
	}

	token, _ := middlewares.GenerateTokenJWT(userDB.Id)

	var userResponse = users.UserResponse{
		Id:    userDB.Id,
		Name:  userDB.Name,
		Email: userDB.Name,
		Token: token,
	}
	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Get Data",
		userResponse,
	))

}

func DetailUserControllers(c echo.Context) error {
	// userId := middlewares.GetUserIdFromJWT(c)
	userId := 1

	userDB, e := database.GetUserDetail(userId)
	paramsUserId := c.Param("userId")
	convertToInt, _ := strconv.Atoi(paramsUserId)

	if convertToInt != userDB.Id {
		return c.JSON(http.StatusBadRequest, BaseResponse(
			http.StatusBadRequest,
			"Bad Request Url Params",
			nil,
		))
	}
	if e != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Get Data",
			nil,
		))
	}
	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Get Data UserId",
		userDB,
	))
}

func GetUserControllers(c echo.Context) error {

	var userData []users.User
	var err error
	userData, err = database.GetDataUserAll()

	if err != nil {
		return c.JSON(http.StatusOK, BaseResponse(
			http.StatusInternalServerError,
			"Failed Get Data",
			userData,
		))
	}

	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Get Data Users",
		userData,
	))
}

func EditUserControllers(c echo.Context) error {

	userId := middlewares.GetUserIdFromJWT(c)
	var userEditData users.UserEdit
	c.Bind(&userEditData)

	// Validasi Field
	errorValidate := validations.Validate(userEditData)
	if errorValidate != nil {
		return errorValidate
	}

	confirmedUser, _ := database.CheckHashPassword(userEditData.ConfirmPassword, userId)
	if !confirmedUser {
		return c.JSON(http.StatusUnauthorized, "Password Konfirmasi Salah")
	}

	userEdit, err := database.EditUser(userEditData, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Edit Data User",
		userEdit,
	))
}

func DeleteUserControllers(c echo.Context) error {
	userId := middlewares.GetUserIdFromJWT(c)

	_, e := database.DeleteUser(userId)

	if e != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Get Data",
			nil,
		))
	}
	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Delete User",
		nil,
	))
}
