package model

import (
	log "github.com/sirupsen/logrus"
	utils "github.com/users/utils"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	//	Id       uint   `json:"id" form:"id"`
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required, gte=8, lte=20"`
}

func generateHash(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return (string(hashedPassword)), nil
}

func (u *User) SaveUser() error {
	db := DB_Connection()
	defer db.Close()

	//hashing the password before saving it to the database

	hashedPassword, err := generateHash(u.Password)
	if err != nil {
		log.Fatal("Hasing of password un-successful")
	}

	//inserting the user into the database
	result, err := db.Prepare("INSERT INTO USERS(email, password) VALUES(?,?)")
	if err != nil {
		return err
	}

	_, err = result.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) AuthenticateUser() (string, error) {

	db := DB_Connection()
	defer db.Close()

	//checking if the user exists in the database
	//matchedRow, err := db.Query("SELECT FROM users WHERE email=? and password=?", u.Email, hash)
	hash, err := generateHash(u.Password)
	if err != nil {
		log.Info(hash)
		log.Fatal("hashing of password un-successful")
	}

	var password string
	err = db.QueryRow("SELECT password FROM users WHERE email = ?", u.Email).Scan(&password)

	if err != nil || password == "" {
		log.Info("Incorrect credentials OR Account does not exist. if later, Please Register first")
		return "", err
	}

	//comparing the password with the hashed password in the database
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(u.Password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			log.Println("Incorrect password. Try Again...")
			return "", err
		}
	}
	//generating a token
	token, err := utils.GenerateToken(u.Email)
	if err != nil {
		log.Println("Error generating token")
		return "", err
	}
	return token, nil
}

func GetUsersFromDB() ([]string, error) {

	var users []string
	var email, password string
	db := DB_Connection()
	defer db.Close()

	result, err := db.Query("select * from users")
	if err != nil {
		return nil, err
	}
	for result.Next() {
		_ = result.Scan(&email, &password)
		users = append(users, email)
	}
	return users, nil
}
