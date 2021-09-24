package middlewares

import (
	"aprian1337/thukul-service/deliveries"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

type JWTCustomClaims struct {
	ID uint `json:"id"`
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
		ErrorHandlerWithContext: func(e error, c echo.Context) error {
			return deliveries.NewErrorResponse(c, http.StatusForbidden, e)
		},
	}
}

func (jwtConf *ConfigJWT) GenerateTokenJWT(id uint) (string, error) {
	claims := JWTCustomClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConf.ExpiresDuration))).Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := t.SignedString([]byte(jwtConf.SecretJWT))

	return token, nil
}

//
//func GetClaimsUserId(c echo.Context) (int, error) {
//	log.Printf("CONTEXT %+v", c)
//	user := c.Get("user")
//	if user != nil {
//		userJwt := user.(*jwt.Token)
//		if userJwt.Valid {
//			claims := userJwt.Claims.(jwt.MapClaims)
//			userId := claims["user_id"].(float64)
//			return int(userId), nil
//		}
//	}
//	return 0, errors.New("Failed Create JWT")
//}
