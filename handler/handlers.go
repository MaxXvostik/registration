package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Redirect(c *gin.Context) {
	c.Redirect(http.StatusFound, "/index.html")
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main page",
	})
}

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Main login",
	})
}

func My(c *gin.Context) {
	c.HTML(http.StatusOK, "my.html", gin.H{
		"title": "Main my",
	})
}

func Signup(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", gin.H{
		"title": "Main signup",
	})
}
