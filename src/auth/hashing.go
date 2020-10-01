package auth

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func getPwd(password string) []byte {
	return []byte(password)
}

func hashAndSalt(pwd []byte) string {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Fatalln(err)
	}
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}
	return true
}

func CheckUser(username string, password string) bool {
	user := getUser(username)
	if user.username == "none" {
		return false
	}
	pwd := getPwd(password)
	pwdResult := comparePasswords(user.password, pwd)
	if pwdResult {
		return true
	}
	return false

}
