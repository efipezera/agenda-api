package controller

import (
	"fmt"
	"net/http"

	"github.com/fplaraujo/agenda/config"
	"github.com/fplaraujo/agenda/model"
	"github.com/gin-gonic/gin"
)

func CreateContact(c *gin.Context) {
	var input model.Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	config.SetupDatabaseConnection()
	defer config.CloseDatabaseConnection()

	contact := model.Contact{
		Name:        input.Name,
		PhoneNumber: input.PhoneNumber,
	}

	config.Database.Create(&contact)

	c.JSON(http.StatusCreated, gin.H{
		"created_contact": contact,
	})
}

func FindContacts(c *gin.Context) {
	config.SetupDatabaseConnection()
	defer config.CloseDatabaseConnection()

	var contacts []model.Contact
	config.Database.Find(&contacts)

	c.JSON(http.StatusOK, gin.H{
		"contacts": contacts,
	})
}

func FindContactByID(c *gin.Context) {
	config.SetupDatabaseConnection()
	defer config.CloseDatabaseConnection()

	contactByID := c.Param("id")
	var contact model.Contact
	config.Database.Table("contacts").Select("id, name, phone_number").Where("id = ?", contactByID).Scan(&contact)

	c.JSON(http.StatusOK, gin.H{
		"contact": contact,
	})
}

func UpdateContact(c *gin.Context) {
	var input model.Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	config.SetupDatabaseConnection()
	defer config.CloseDatabaseConnection()

	// updatedContact := model.Contact{
	// 	Name:        input.Name,
	// 	PhoneNumber: input.PhoneNumber,
	// }

	contactByID := c.Param("id")

	// config.Database.Table("contacts").Select("name, phone_number").Updates(updatedContact).Where("id = ?", contactByID).Scan(&updatedContact)
	config.Database.Table("contacts").Where("id = ?", contactByID).Save(&input)
	fmt.Println(input)
	c.JSON(http.StatusOK, gin.H{
		"updated": input,
	})
}
