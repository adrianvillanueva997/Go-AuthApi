package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // It is needed to run MySQL
	"log"
	"os"
)

func SetConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", ""+os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+
		"@tcp("+os.Getenv("SERVER_IP")+")/"+os.Getenv("DATABASE)+"))
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return db, nil
}
