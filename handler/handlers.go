package handler

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Redirect(c *gin.Context) {
	c.Redirect(http.StatusFound, "index.html")
}

func Index(c *gin.Context) {
	sessions := sessions.Default(c)
	uid, _ := sessions.Get("uid").(int64)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"uid": uid,
	})
}

func Login(c *gin.Context) {

	sessions := sessions.Default(c)
	uid, _ := sessions.Get("uid").(int64)

	//if !ok || uid == 0 {
	//	c.Redirect(302, "signup.html")
	//	}

	c.HTML(http.StatusOK, "login.html", gin.H{
		"uid": uid,
	})
}

func My(c *gin.Context) {
	sessions := sessions.Default(c)
	uid, _ := sessions.Get("uid").(int64)
	c.HTML(http.StatusOK, "my.html", gin.H{
		"uid": uid,
	})
}

func Signup(c *gin.Context) {
	sessions := sessions.Default(c)
	uid, _ := sessions.Get("uid").(int64)
	c.HTML(http.StatusOK, "signup.html", gin.H{
		"uid": uid,
	})
}
