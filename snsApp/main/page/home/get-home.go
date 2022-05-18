package home

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func homeUserList(c *gin.Context) {
	db, _ := sql.Open("mysql", "root:docker@tcp(localhost:3306)/swe-training")
	defer db.Close()

	rows, err := db.Query("SELECT sno, name from test_table")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var sno int
	var name string

	for rows.Next() {
		_ = rows.Scan(&sno, &name)
		fmt.Println(sno, name)
	}
	c.JSON(200, gin.H{
		"message": "OK",
	})
}

func HomeManager(r *gin.Engine) {
	r.GET("/post/home", homeUserList)
}
