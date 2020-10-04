package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func getPwd(password string) []byte {
	return []byte(password)
}

/*
func hashAndSalt(pwd []byte) string {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Fatalln(err)
	}
	return string(hash)
}
*/
func comparePasswords(hashedPwd string, plainPwd []byte) error {
	byteHash := []byte(hashedPwd)
	return bcrypt.CompareHashAndPassword(byteHash, plainPwd)
}

func CheckUser(username string, password string) bool {
	user := getUser(username)
	if user.username == "none" {
		return false
	}
	pwd := getPwd(password)
	if comparePasswords(user.password, pwd) == nil {
		return true
	} else {
		return false
	}
}
