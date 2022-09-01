package middleware

import (
	"net/mail"
	"time"
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
