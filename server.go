package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func initMysql() *sql.DB {
	db, err := sql.Open("mysql", "qwexodn:rlaxodn123!@tcp(db:3308)/TEST_DB")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "hi")
	})

	r.Run(":80")
}
