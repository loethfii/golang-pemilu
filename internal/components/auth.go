package components

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"luthfi/pemilu/domain"
	"luthfi/pemilu/internal/config"
	"net/http"
	"time"
)

type JwtClaims struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(userData domain.User) string {
	cnf := config.Get()
	
	claims := &JwtClaims{
		ID:       userData.ID,
		Username: userData.Username,
		Role:     userData.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
		},
	}
	
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	t, err := tokenClaim.SignedString([]byte(cnf.SecretKey.SecretKey))
	if err != nil {
		panic(err)
	}
	
	return t
}

func TokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	cnf := config.Get()
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Need Token")
		}
		
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(cnf.SecretKey.SecretKey), nil
		})
		
		if err != nil || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Token")
		}
		
		return next(c)
	}
}
