package main

import (
	"example.com/htmlRegistration/autorizetion"
	"example.com/htmlRegistration/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.LoadHTMLGlob("html/*")

	router.GET("/", handler.Redirect)
	router.GET("/index.html", handler.Index)

	router.GET("/login.html", handler.Login)
	router.POST("/login.html", autorizetion.Validation)

	router.GET("/signup.html", handler.Signup)
	router.POST("/signup.html", autorizetion.Registration)

	router.GET("/my.html", handler.My)

	router.Run()
}
