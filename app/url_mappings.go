package app

import (
	"github.com/uuthman/bookstore_users-api/controllers/users"
	"github.com/uuthman/bookstore_users-api/controllers/ping"
)

func mapUrls(){
	router.GET("/ping",ping.Ping)
	router.POST("/users",users.CreateUser)
	router.GET("/users/:user_id",users.GetUser)
	// router.GET("/users/search",controllers.FindUser)
} 