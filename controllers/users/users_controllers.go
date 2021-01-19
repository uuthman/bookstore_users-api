package users

import (
	"strconv"
	"net/http"
	"github.com/uuthman/bookstore_users-api/services"
	"github.com/uuthman/bookstore_users-api/utils/errors"
	"github.com/uuthman/bookstore_users-api/domain/users"
	"github.com/gin-gonic/gin"
)



func getUserID(userIdParam string) (int64,*errors.RestErr){
	userID,userErr := strconv.ParseInt(userIdParam,10,64)

	if userErr != nil{
		return 0, errors.NewBadRequestError("user id should be a number")
	}

	return userID,nil
}

func Get(c *gin.Context){
	
	userID,err := getUserID(c.Param("user_id"))

	if err != nil{
		c.JSON(err.Status,err)
		return
	}

	result,saveErr := services.UsersService.GetUser(userID)
	if saveErr != nil{
		c.JSON(saveErr.Status,saveErr)
		return
	}

	c.JSON(http.StatusCreated,result.Marshall(c.GetHeader("X-Public") == "true"))
}

func Create(c *gin.Context){
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil{
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status,restErr)
		return
	}

	result,saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil{
		c.JSON(saveErr.Status,saveErr)
		return
	}

	c.JSON(http.StatusCreated,result.Marshall(c.GetHeader("X-Public") == "true"))
}

func Update(c *gin.Context){

	userID,err := getUserID(c.Param("user_id"))

	if err != nil{
		c.JSON(err.Status,err)
		return
	}
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil{
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status,restErr)
		return
	}

	user.ID = userID

	isPartial := c.Request.Method == http.MethodPatch

	result,err := services.UsersService.UpdateUser(isPartial,user)
	if err != nil{
		c.JSON(err.Status,err)
		return
	}

	c.JSON(http.StatusOK,result.Marshall(c.GetHeader("X-Public") == "true"))


}

func Delete(c *gin.Context){
	userID,err := getUserID(c.Param("user_id"))

	if err != nil{
		c.JSON(err.Status,err)
		return
	}


	if err := services.UsersService.DeleteUser(userID); err != nil{
		c.JSON(err.Status,err)
		return
	}

	c.JSON(http.StatusOK,map[string]string{"status":"deleted"})
}

func Search(c *gin.Context){
	status := c.Query("status")

	users,err :=services.UsersService.SearchUser(status)

	if err != nil{
		c.JSON(err.Status,err)
		return 
	}
	c.JSON(http.StatusOK,users.Marshall(c.GetHeader("X-Public") == "true"))
}