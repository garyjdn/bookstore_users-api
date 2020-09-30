package users

import (
	"github.com/garyjdn/bookstore_users-api/domain/users"
	"github.com/garyjdn/bookstore_users-api/services"
	"github.com/garyjdn/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.BadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	res, err := services.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, res)
}

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		err := errors.BadRequestError("invalid user id, user id should be number")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func FindUser(c *gin.Context) {}
