package model

import "golang.org/x/crypto/bcrypt"

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
