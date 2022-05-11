package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type jwtMiddleware struct {
	key string
}

type JWTService interface {
	JwtMiddleware() echo.MiddlewareFunc
}

var (
	jwtSignedMethod = jwt.SigningMethodHS256
)

func NewJwtService(secretKey string) JWTService {
	return &jwtMiddleware{
		key: secretKey,
	}
}

// JwtMiddleware implements JWTService
func (s *jwtMiddleware) JwtMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			fmt.Println("masuk middleware")

			signature := strings.Split(c.Request().Header.Get("Authorization"), " ")
			if len(signature) < 2 {
				return c.JSON(http.StatusForbidden, "Invalid token")
			}

			if signature[0] != "Bearer" {
				return c.JSON(http.StatusForbidden, "Invalid token")
			}

			claim := jwt.MapClaims{}

			token, _ := jwt.ParseWithClaims(signature[1], claim, func(t *jwt.Token) (interface{}, error) {
				return []byte("my_screet_key"), nil
			})

			method, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok || method != jwtSignedMethod {
				return c.JSON(http.StatusForbidden, "Invalid token")
			}

			c.Set("payload", fmt.Sprintf("%v:%s", claim["ID"], claim["Username"].(string)))

			return next(c)
		}
	}
}
