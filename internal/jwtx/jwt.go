package jwtx

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// The `type JWTCustomClaims struct` is defining a custom claims structure for JWT (JSON Web Token). It
// extends the `jwt.RegisteredClaims` structure and adds an additional field `ID` of type string. This
// custom claims structure will be used to store the user ID in the JWT token.
type JWTCustomClaims struct {
	ID string
	jwt.RegisteredClaims
}

// The `type JWTConfig struct` is defining a struct that represents the configuration for JWT (JSON Web
// Token) handling. It has two fields: `SecretKey` of type string, which represents the secret key used
// for signing and verifying JWT tokens, and `ExpiresDuration` of type int, which represents the
// duration in hours for which the JWT token will be valid.
type JWTConfig struct {
	SecretKey      string
	ExpiresDuraton int
}

// The `Init()` function is a method of the `JWTConfig` struct. It returns an `echojwt.Config` object
// that is used to configure JWT handling in an Echo web framework application.
func (jwtConfig *JWTConfig) Init() echojwt.Config {
	return echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JWTCustomClaims)
		},
		SigningKey: []byte(jwtConfig.SecretKey),
	}
}

// The `GenerateToken` function is a method of the `JWTConfig` struct. It takes a `userID` string as
// input and returns a JWT token string and an error.
func (jwtConfig *JWTConfig) GenerateToken(userID string) (string, error) {
	expire := jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConfig.ExpiresDuraton))))

	claims := &JWTCustomClaims{
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: expire,
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := rawToken.SignedString([]byte(jwtConfig.SecretKey))

	if err != nil {
		return "", err
	}

	return token, nil
}

// The GetUser function retrieves the JWT custom claims for a given context.
func GetUser(c echo.Context) (*JWTCustomClaims, error) {
	user := c.Get("user").(*jwt.Token)

	if user == nil {
		return nil, errors.New("invalid")
	}

	claims := user.Claims.(*JWTCustomClaims)

	return claims, nil
}

// The `func VerifyToken(next echo.HandlerFunc) echo.HandlerFunc` is a middleware function that is
// used to verify the JWT token in an Echo web framework application.
func VerifyToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userData, err := GetUser(c)

		isInvalid := userData == nil || err != nil

		if isInvalid {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "invalid token",
			})
		}

		return next(c)
	}
}
