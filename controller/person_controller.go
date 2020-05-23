package controller

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/matimartini/api-numerologia/service"
)

type PersonController struct {
	Context *gin.Context
}

func (c PersonController) GetPerson() {

	fullName := c.Context.Query("name")
	date := c.Context.Query("date")
	c.validateParam(fullName, date)

	peroson := &service.Person{}
	peroson.NewPerson(fullName, date)

	c.Context.JSON(200, gin.H{
		"data": peroson,
	})

}

func (c PersonController) validateParam(name string, date string) {
	if _, err := strconv.Atoi(name); err == nil {
		c.Context.JSON(400, gin.H{"message ": "Name parsing error"})
	}

	if _, err := time.Parse("2006-01-02", date); err != nil {
		c.Context.JSON(400, gin.H{"message ": "Error entering date of birth"})
	}
}
