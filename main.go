package main

import (
	"github.com/gin-gonic/gin"
	"github.com/giriaditya/test-gin/app"
	"github.com/giriaditya/test-gin/controllers"
	"github.com/giriaditya/test-gin/service"
)

func main() {
	db := app.ConnectDatabase()
	bookService := service.NewBookService(db)
	bookController := controllers.NewBookController(bookService)

	// Init Gin
	r := gin.Default()

	r.POST("/books", bookController.Create)
	r.PATCH("/books/:id", bookController.Update)
	r.DELETE("/books/:id", bookController.Delete)
	r.GET("/books/", bookController.FetchAll)
	r.Run(":80")
}
