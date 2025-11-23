package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	
	userPassword,err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil{
		return "", err
	}

	return string(userPassword), nil
}

func ValidatePassword(hashedPassword, plainPassword []byte) error{
	
	err:= bcrypt.CompareHashAndPassword(hashedPassword,plainPassword)
	if err != nil{
		return err
	}

	return nil
}

