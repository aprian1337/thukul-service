package middlewares

import (
	"context"
	"github.com/labstack/echo/v4"
)

func HttpHeaderMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Request().WithContext(context.WithValue(c.Request().Context(), "tokentest", "zzzzzz"))
		c.QueryParams().Add("token", "ZZZZZ")
		c.QueryParams().Add("tokentest", "ZZZZZ")
		return next(c)
	}
}


//func IsUserId(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		user := c.Get("user").(*jwt.Token)
//		userId := c.Param("userId")
//		convUserId, err := helpers.StringToUint(userId)
//		if err != nil {
//			return echo.ErrBadRequest
//		}
//		claims := user.Claims.(*JWTCustomClaims)
//		claimUserId := claims.ID
//		if claimUserId != convUserId {
//			return echo.ErrUnauthorized
//		}
//		return next(c)
//	}
//}