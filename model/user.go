package model

import (
	"log"

	utils "github.com/users/utils"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       uint   `json:"id" form:"id"`
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required, gte=8, lte=20"`
}

func (u *User) SaveUser() error {
	db := DB_Connection()
	defer db.Close()

	//hashing the password before saving it to the database

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//inserting the user into the database
	result, err := db.Prepare("INSERT INTO USERS(email, password) VALUES(?,?)")
	if err != nil {
		return err
	}

	_, err = result.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) AuthenticateUser() (string, error) {

	db := DB_Connection()
	defer db.Close()

	//checking if the user exists in the database

	matchedRow, err := db.Query("SELECT password FROM users WHERE email = ?", u.Email)
	if err != nil {
		log.Println("Account does not exist. Please register")
		return "", err
	}

	for matchedRow.Next() {
		var hashedPassword string
		err = matchedRow.Scan(&hashedPassword)
		if err != nil {
			return "", err
		}

		//comparing the password with the hashed password in the database
		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(u.Password))
		if err != nil {
			if err == bcrypt.ErrMismatchedHashAndPassword {
				log.Println("Incorrect password")
				return "", err
			} else {
				return "", err
			}
		}
	}

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
	var id uint
	db := DB_Connection()
	defer db.Close()

	result, err := db.Query("select * from users")
	if err != nil {
		return nil, err
	}
	for result.Next() {
		_ = result.Scan(&id, &email, &password)
		users = append(users, email)
	}
	return users, nil
}
