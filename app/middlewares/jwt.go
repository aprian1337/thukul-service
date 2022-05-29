package middlewares

import (
	businesses "aprian1337/thukul-service/business"
	"aprian1337/thukul-service/deliveries"
	"aprian1337/thukul-service/helpers"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

type JWTCustomClaims struct {
	ID      uint `json:"id"`
	IsAdmin int  `json:"is_admin"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiresDuration int
}

func (jwtConf *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JWTCustomClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
		SuccessHandler: func(context echo.Context) {
			//userId := context.Get("userId")
		},
		ErrorHandlerWithContext: func(e error, c echo.Context) error {
			return deliveries.NewErrorResponse(c, http.StatusUnauthorized, businesses.ErrInvalidTokenCredential)
		},
	}
}

func (jwtConf *ConfigJWT) GenerateTokenJWT(id uint, isAdmin int) (string, error) {
	claims := JWTCustomClaims{
		id,
		isAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConf.ExpiresDuration))).Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := t.SignedString([]byte(jwtConf.SecretJWT))

	return token, nil
}

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*JWTCustomClaims)
		isAdmin := claims.IsAdmin
		if isAdmin == 0 {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

func IsUserId(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		userId := c.Param("userId")
		convUserId, err := helpers.StringToUint(userId)
		if err != nil {
			return echo.ErrBadRequest
		}
		claims := user.Claims.(*JWTCustomClaims)
		claimUserId := claims.ID
		if claimUserId != convUserId {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

func JwtValidate(token string) (*jwt.Token, error){
	return jwt.ParseWithClaims(token, &JWTCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there's a problem with the signing method")
		}
		viper.SetConfigFile(`config.json`)
		secret := viper.GetString(`jwt.secret`)
		return []byte(secret), nil
	})
}

func GetClaimsUserId(token string) (uint, error){
	validate, err := JwtValidate(token)
	if err != nil {
		return 0, err
	}
	jwtClaims := validate.Claims.(*JWTCustomClaims).ID
	return jwtClaims, nil
}

func GetClaimsAdminId(token string) (uint, error){
	validate, err := JwtValidate(token)
	if err != nil {
		return 0, err
	}
	if validate.Claims.(*JWTCustomClaims).IsAdmin == 0 {
		return 0, fmt.Errorf("you're not have authorized for this action")
	}
	jwtClaims := validate.Claims.(*JWTCustomClaims).ID
	return jwtClaims, nil
}
