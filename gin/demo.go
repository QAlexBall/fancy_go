package main

import (
	"gin_demo/models"
	"gin_demo/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	db := utils.Connect_DB("chris", "chris", "demo")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		var err = db.Insert(&models.User{
			Name:   "chris",
			Emails: []string{"chris@chris"},
		})
		if err != nil {
			panic(err)
		}
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/list_users", func(c *gin.Context) {
		var users []models.User
		var err = db.Model(&users).Select()
		if err != nil {
			panic(err)
		}
		c.JSON(200, gin.H{
			"message": users,
		})
	})
	r.Run(":8081")
}
