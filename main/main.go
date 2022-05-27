package main

import (
	"forbizbe/main/svmn/Gin"
	"forbizbe/main/svmn/Mysql"
)

func main() {
	Mysql.ConnDB()
	r := Gin.SetupRouter()
	r.Run(":8080")
}
