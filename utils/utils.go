package middleware

import (
	"net/mail"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetDurationInMilliSeconds(start time.Time) float64 {

	end := time.Now()
	duration := end.Sub(start)
	DurationInMillis := float64(duration) / float64(time.Millisecond)

	return DurationInMillis
}

func ValidateEmail(email string) (bool, error) {
	if _, err := mail.ParseAddress(email); err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func GenerateToken(email string) (string, error) {

	//creating a token
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}).SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func VerifyToken(c *gin.Context) (string, error) {
	token := c.Query("token")
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return "", err
	}
	return claims["email"].(string), nil
}
