package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/monikanaico-hub/goapi/database"
	"github.com/monikanaico-hub/goapi/models"
)

func CreateUserHandler(c *gin.Context) {
	var userdata models.User

	// App level validation
	err := c.BindJSON(&userdata)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(err))
		return
	}

	// Inserting data
	id, err := database.CreateUser(userdata)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf(err.Error()))
	} else {
		userdata.Id = id
		c.JSON(http.StatusCreated, userdata)
	}
}

func GetallUserHandler(c *gin.Context) {

	// fetch data
	userreslt, err := database.GetallUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf(err.Error()))
	} else {
		c.JSON(http.StatusOK, userreslt)
	}
}

func GetSpecificHandler(c *gin.Context) {
	id := GetInt64IdFromReqContext(c)
	userList, _ := database.FindById(id)

	// Check if resource exist
	if userList.Id == 0 {
		c.JSON(http.StatusNotFound, "Not found")
	} else {
		c.JSON(http.StatusOK, userList)
	}
}

func UpdateUserHandler(c *gin.Context) {
	id := GetInt64IdFromReqContext(c)
	var userList models.User

	// App level validation
	err := c.BindJSON(&userList)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(err))
	}

	// Check if resource exist
	checkUserist, _ := database.FindById(id)
	if checkUserist.Id == 0 {
		c.JSON(http.StatusNotFound, "Not found")
		return
	}

	// Updating data
	userRslt, err := database.UpdateUser(checkUserist.Id, userList)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusCreated, userRslt)
	}
}

func DeleteUserHandler(c *gin.Context) {
	id := GetInt64IdFromReqContext(c)

	// Check if resource exist
	getUserList, _ := database.FindById(id)
	if getUserList.Id == 0 {
		c.JSON(http.StatusNotFound, "Not found")
		return
	}

	// Deleting data
	err := database.DeleteUser(getUserList)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, "Successful Deletion")
	}
}
