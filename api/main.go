package main

import (
	"context"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
)

type AppState struct {
	Firebase *firebase.App
}

func main() {
	router := gin.Default()
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	add_routes(router)

	router.Use(func(ctx *gin.Context) {
		ctx.Set("state", app)
		ctx.Next()
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

type HelloForm struct {
	Name string `form:"name"`
}

func hello_form(c *gin.Context) {
	var form HelloForm
	c.Bind(&form)
	c.JSON(200, gin.H{
		"message": "Hello " + form.Name,
	})
}

func add_routes(router *gin.Engine) {
	router.GET("/ping", pong)
	router.POST("/hello", hello_form)
}
