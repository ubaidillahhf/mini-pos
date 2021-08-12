package database

import (
	"pos/configs"
	"pos/helpers"
	"pos/models/users"
)

func RegisterUser(usersCreate users.UserCreate) (users.User, error) {
	hash, _ := helpers.HashPassword(usersCreate.Password)
	var userDB users.User

	userDB.Name = usersCreate.Name
	userDB.Address = usersCreate.Address
	userDB.Phone = usersCreate.Phone
	userDB.Email = usersCreate.Email
	userDB.Password = hash
	userDB.Remark = usersCreate.Remark

	err := configs.DB.Create(&userDB).Error
	if err != nil {
		return userDB, err
	}
	return userDB, nil
}

func GetDataUserAll() (dataResult []users.User, err error) {
	err = configs.DB.Select("id", "name", "address", "phone", "email", "remark", "created_at", "updated_at", "deleted_at").Find(&dataResult).Error
	if err != nil {
		return nil, err
	}
	return
}

func LoginUser(userLogin users.UserLogin) (users.User, error) {
	var userDB users.User

	err := configs.DB.Where("email = ?", userLogin.Email).First(&userDB).Error
	checkHash, _ := helpers.CheckPasswordHash(userLogin.Password, userDB.Password)

	if err != nil && !checkHash {
		return userDB, err
	}
	return userDB, nil
}

func GetUserDetail(userId int) (users.User, error) {
	var userDB users.User
	err := configs.DB.First(&userDB, userId).Error

	if err != nil {
		return userDB, err
	}
	return userDB, nil
}

func CheckHashPassword(confirmPassword string, userId int) (verified bool, err error) {
	var userDB users.User

	err = configs.DB.Where("id = ?", userId).First(&userDB).Error
	confirmedUser, _ := helpers.CheckPasswordHash(confirmPassword, userDB.Password)

	if !confirmedUser {
		return false, err
	}
	return true, err
}

func EditUser(userEdit users.UserEdit, userId int) (users.User, error) {
	hash, _ := helpers.HashPassword(userEdit.NewPassword)
	var userDB users.User
	err := configs.DB.First(&userDB, userId).Error

	userDB.Name = userEdit.Name
	userDB.Address = userEdit.Address
	userDB.Phone = userEdit.Phone
	userDB.Email = userEdit.Email
	userDB.Password = hash
	userDB.Remark = userEdit.Remark

	configs.DB.Save(&userDB)

	if err != nil {
		return userDB, err
	}
	return userDB, nil
}

func DeleteUser(userId int) (users.User, error) {
	var userDB users.User
	err := configs.DB.Where("id = ?", userId).Delete(&userDB).Error

	if err != nil {
		return userDB, err
	}
	return userDB, nil
}
