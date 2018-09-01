package main

import (
	"fmt"

	"github.com/lakshanwd/muve-go/go-crud/repo"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/lakshanwd/go-crud/handler"
)

func main() {
	fmt.Println("setup database connections")
	repo.SetupRepo()
	defer repo.CloseRepo()

	fmt.Println("handling roures")
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.GET("/book", handler.BookGetHandler)
	router.POST("/book", handler.BookPostHandler)
	router.PUT("/book", handler.BookPutHandler)
	router.DELETE("/book/:id", handler.BookDeleteHandler)

	router.GET("/student", handler.StudentGetHandler)
	router.POST("/student", handler.StudentPostHandler)
	router.PUT("/student", handler.StudentPutHandler)
	router.DELETE("/student/:id", handler.StudentDeleteHandler)

	router.Run(":8080")
}
