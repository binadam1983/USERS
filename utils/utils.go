package middleware

import (
	"net/http"
	"net/mail"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"
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
		"exp":   time.Now().Add(3 * time.Minute).Unix(),
	}).SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func extractToken(c *gin.Context) string {

	token, err := c.Cookie("token")

	if err != nil || token == "" || len(strings.Split(token, ".")) != 3 {
		token = c.GetHeader("Authorization")
		if (strings.Split(token, " "))[0] == "Bearer" {
			token = strings.Split(token, " ")[1]
			log.Info("Bearer - ", token)
			return token
		}
	}
	return token
}

func VerifyToken(c *gin.Context) (string, error) {

	token := extractToken(c)
	if token == "" {
		log.Info("token is empty")
		return "", http.ErrNoCookie
	}
	log.Info("token is not empty...")
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		} else {
			return []byte(os.Getenv("SECRET_KEY")), nil
		}
	})
	if err != nil {
		return "", err
	}
	return claims["email"].(string), nil
}
