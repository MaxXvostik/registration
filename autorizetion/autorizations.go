package autorizetion

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignupForm struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginForm struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func Validation(c *gin.Context) {
	var form LoginForm
	if err := c.ShouldBind(&form); err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"msg": err.Error(),
		})
		return
	}
	if form.Email != "my@mail.ru" || form.Password != "123" {
		c.HTML(http.StatusUnauthorized, "error.html", gin.H{
			"msg": "unathorized",
		})
		return
	}
	c.Redirect(http.StatusFound, "/index.html")
}

func Registration(c *gin.Context) {
	var form SignupForm
	if err := c.ShouldBind(&form); err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"msg": err.Error(),
		})
		return
	}
	println(form.Email, form.Name, form.Password)
	c.Redirect(http.StatusFound, "/login.html")
}
