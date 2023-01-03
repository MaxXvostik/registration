package main

import (
	"example.com/htmlRegistration/autorizetion"
	"example.com/htmlRegistration/database"
	"example.com/htmlRegistration/handler"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {

	database.GetDB()

	router := gin.New()
	router.LoadHTMLGlob("html/*")

	store := cookie.NewStore([]byte("ratatata"))
	router.Use(sessions.Sessions("sess", store))

	router.GET("/", handler.Redirect)
	router.GET("/index.html", handler.Index)

	router.GET("/login.html", handler.Login)
	router.POST("/login.html", autorizetion.Validation)

	router.GET("/signup.html", handler.Signup)
	router.POST("/signup.html", autorizetion.Registration)

	router.GET("/my.html", handler.My)

	router.Run("127.00.01:8080")
}
