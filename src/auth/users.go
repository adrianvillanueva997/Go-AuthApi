package auth

import (
	"Auth_Api/src/db"
	"log"
)

type User struct {
	username string
	password string
}

func getUser(username string) User {
	cursor, err := db.SetConnection()
	if err != nil {
		log.Fatalln(err)
	}
	var user User
	results, err := cursor.Query("SELECT name, password from ApiAuth.App where name = ?", username)
	if err != nil {
		log.Fatalln(err)
	}
	if results.Next() {
		err := results.Scan(&user.username, &user.password)
		if err != nil {
			log.Fatalln(err)
		}
		return user
	}
	user.username = "none"
	return user
}
