package library

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(pass),bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashPass), nil

}

func CheckPassword(bodypass, dbPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(bodypass), []byte(dbPass))
	
	if err != nil {
		return false
	}

	return true
}