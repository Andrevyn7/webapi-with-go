package controllers

import (
	"github.com/Andrevyn7/webapi-with-go.git/database"
	"github.com/Andrevyn7/webapi-with-go.git/database/migrations"
	"github.com/Andrevyn7/webapi-with-go.git/models"
	"github.com/Andrevyn7/webapi-with-go.git/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	db := database.GetDatabase()

	migrations.RunMigrations(db)

	var p models.User

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	p.Password = services.SHA256Encoder(p.Password)

	err = db.Create(&p).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create book:" + err.Error(),
		})

		return
	}
	c.Status(204)
}
