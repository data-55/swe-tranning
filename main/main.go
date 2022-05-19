package main

import (
	"forbizbe/main/page/home"
	"forbizbe/main/svmn"
)

func main() {
	r := svmn.SetupRouter()
	home.HomeManager(r)
	r.Run(":8080")
}

// ↓↓↓↓↓This is a sample↓↓↓↓↓
// func main() {
// 	engine := gin.Default()
// 	engine.GET("/", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "hello world",
// 		})
// 	})
// 	engine.Run(":8080")
// }
