package autorizetion

import (
	"net/http"

	"example.com/htmlRegistration/database"
	"example.com/htmlRegistration/handler"
	"example.com/htmlRegistration/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var Test int = 5

func Validation(c *gin.Context) {
	var form model.LoginForm
	if err := c.ShouldBind(&form); err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"msg": err.Error(),
		})
		return
	}

	row := database.DB.QueryRow("SELECT id,name, email,passhash FROM user WHERE email=?", form.Email)
	u := new(model.User)
	err := row.Scan(&u.Id, &u.Name, &u.Email, &u.Passhash)
	if err != nil {
		logrus.Error(err)
	}

	validation := bcrypt.CompareHashAndPassword([]byte(u.Passhash), []byte(form.Password))
	if validation != nil {
		c.HTML(http.StatusUnauthorized, "error.html", gin.H{
			"msg": "unathorized",
		})
		return
	}
	sessions := sessions.Default(c)
	sessions.Set("uid", u.Id)
	sessions.Save()

	c.Redirect(http.StatusFound, "my.html")
}

func Registration(c *gin.Context) {
	var form model.SignupForm
	if err := c.ShouldBind(&form); err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"msg": err.Error(),
		})
		return
	}

	println(form.Email, form.Name, form.Password)

	//Хеширование пароля
	hash, _ := handler.HashPassword([]byte(form.Password), 13)
	println(hash)

	//Внесение в базу данных
	sqlStatement := "INSERT INTO user (name, email,password, passhash) VALUES(?,?,?,?)"
	_, err := database.DB.Exec(sqlStatement, form.Name, form.Email, form.Password, string(hash))
	if err != nil {
		logrus.Error(err)
	}

	c.Redirect(http.StatusFound, "login.html")
}
