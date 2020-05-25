package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func InitMysql() *sql.DB {
	db, err := sql.Open("mysql", "root:123123@tcp(192.168.99.100:3308)/db_test")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func main() {

	r := gin.New()

	r.GET("/", func(c *gin.Context) {

		db := InitMysql()
		defer db.Close()

		rows, err := db.Query("SELECT * FROM users")

		if err != nil {
			log.Fatal(err)
		}

		type users struct {
			Pid  int
			Name string
		}

		var userss = map[string]interface{}{}

		for rows.Next() {
			user := users{}
			err := rows.Scan(&user.Pid, &user.Name)

			if err != nil {
				log.Fatal(err)
			}

			userss[fmt.Sprintf("%v", user.Pid)] = user
		}
		c.JSON(200, userss)
	})

	r.Run(":80")
}
