package users

import (
	"strconv"
	"net/http"
	"github.com/uuthman/bookstore_users-api/services"
	"github.com/uuthman/bookstore_users-api/services/utils/errors"
	"github.com/uuthman/bookstore_users-api/domain/users"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context){
	userID,userErr := strconv.ParseInt(c.Param("user_id"),10,64)

	if userErr != nil{
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status,err)
		return
	}

	result,saveErr := services.GetUser(userID)
	if saveErr != nil{
		c.JSON(saveErr.Status,saveErr)
		return
	}

	c.JSON(http.StatusCreated,result)
}

func CreateUser(c *gin.Context){
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil{
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status,restErr)
		return
	}

	result,saveErr := services.CreateUser(user)
	if saveErr != nil{
		c.JSON(saveErr.Status,saveErr)
		return
	}

	c.JSON(http.StatusCreated,result)
}

// func FindUser(c *gin.Context){
	
// }