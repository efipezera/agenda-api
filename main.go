package main

import (
	"github.com/fplaraujo/agenda/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	contactRouter := r.Group("api/contact")
	{
		contactRouter.GET("/", controller.FindContacts)
		contactRouter.POST("/", controller.CreateContact)
		contactRouter.GET("/:id", controller.FindContactByID)
		contactRouter.PUT("/:id", controller.UpdateContact)
	}

	r.Run()
}
