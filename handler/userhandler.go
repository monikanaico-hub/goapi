package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/monikanaico-hub/goapi/database"
	"github.com/monikanaico-hub/goapi/models"
	"github.com/monikanaico-hub/goapi/views"
)

func CreateUserHandler(c *gin.Context) {
	var userdata models.User

	// App level validation
	bindErr := c.BindJSON(&userdata)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(bindErr))
		return
	}

	// Inserting data
	id, insertErr := views.CreateUser(userdata)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Something wrong on our server"))
		database.Dberror(insertErr)
	} else {
		userdata.Id = id
		c.JSON(http.StatusCreated, userdata)
	}
}

func GetallUserHandler(c *gin.Context) {

	// fetch data
	userreslt, selectErr := views.GetallUser()
	if selectErr != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Something wrong on our server"))
		database.Dberror(selectErr)
	} else {
		c.JSON(http.StatusOK, userreslt)
	}
}
