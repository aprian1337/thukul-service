package middlewares

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"time"
)

type jwtCustomClaims struct {
	UserId uint `json:"user_id"`
	jwt.StandardClaims
}

func GenerateTokenJWT(id uint) (string, error) {
	claims := jwtCustomClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 72).Unix(),
		},
	}
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	viper.SetConfigFile(`config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	newJwt, errJwt := newToken.SignedString([]byte(viper.GetString("signed_string")))
	if errJwt != nil {
		return "", errJwt
	}
	return newJwt, nil
}

func GetClaimsUserId(c echo.Context) (int, error) {
	log.Printf("CONTEXT %+v", c)
	user := c.Get("user")
	if user != nil {
		userJwt := user.(*jwt.Token)
		if userJwt.Valid {
			claims := userJwt.Claims.(jwt.MapClaims)
			userId := claims["user_id"].(float64)
			return int(userId), nil
		}
	}
	return 0, errors.New("Failed Create JWT")
}
