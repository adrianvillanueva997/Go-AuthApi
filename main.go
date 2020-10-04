package main

import (
	"Auth_Api/src/auth"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
	router := gin.Default()
	router.GET("/api/auth", func(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password")
		if (username != "") && (password != "") {
			authCheck := auth.CheckUser(username, password)
			if authCheck {
				c.JSON(200, gin.H{"message": "Ok"})
			} else {
				c.JSON(401, gin.H{"message": "No"})
			}
		} else {
			c.JSON(401, gin.H{"message": "No"})
		}
	})
	log.Println("Server running!")
	err = router.Run("0.0.0.0:3000")
	if err != nil {
		log.Fatalln(err)
	}
}
