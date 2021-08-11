package middlewares

import (
	"pos/constants"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GenerateTokenJWT(userId int) (string, error) {
	// generate payload
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// header
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// token
	return token.SignedString([]byte(constants.JWT_SECRET))

}

func GetUserIdFromJWT(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return int(userId)
	}
	return 0
}
