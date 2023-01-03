package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//type User struct {
//	Id       int    `json:"id"`
//	Name     string `json:"name"`
//	Email    string `json:"email"`
//	Password string `json:"password"`
//}

const (
	host     = "localhost"
	user     = "root"
	password = "pass"
	dbname   = "htmlRegistration"
)

var DB *sql.DB

func Init() *sql.DB {
	psqlInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		user, password, host, dbname)

	db, err := sql.Open("mysql", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}

func GetDB() *sql.DB {
	if DB == nil {
		DB = Init()
		var sleep = time.Duration(1)
		for DB == nil {
			sleep *= 2
			fmt.Println("db не доступна")
			time.Sleep(sleep * time.Second)
			DB = Init()
		}
	}
	return DB
}
