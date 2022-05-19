package home

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type user struct {
	ID        uint      `gorm:"primaryKey; autoIncrement"`
	Name      string    `gorm:"not null; size:150"`
	Email     string    `gorm:"not null; unique; size:300"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func homeUserList(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK",
	})

	dns := "docker:docker@tcp(auwellnessforbiz-cms-db-container:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&user{})

}

// func homeUserList(c *gin.Context) {
// 	db, _ := sql.Open("mysql", "docker:docker@tcp(auwellnessforbiz-cms-db-container:3306)/test")
// 	defer db.Close()

// 	rows, err := db.Query("SELECT sno, name from test_table")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer rows.Close()

// 	var sno int
// 	var name string

// 	for rows.Next() {
// 		_ = rows.Scan(&sno, &name)
// 		fmt.Println(sno, name)
// 	}

// 	c.JSON(200, gin.H{
// 		"message": "OK",
// 	})
// }

func HomeManager(r *gin.Engine) {
	r.GET("/post/home", homeUserList)
	r.GET("/post/home/seed", seed)
}

func seed(c *gin.Context) {
	dns := "docker:docker@tcp(auwellnessforbiz-cms-db-container:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	tmp := []user{
		{ID: 1, Name: "ABC", Email: "ta11@ss.com", Password: "1234", CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
		{ID: 2, Name: "ABC", Email: "ta21@ss.com", Password: "1234", CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
		{ID: 3, Name: "ABC", Email: "ta31@ss.com", Password: "1234", CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
		{ID: 4, Name: "ABC", Email: "ta41@ss.com", Password: "1234", CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
		{ID: 5, Name: "ABC", Email: "ta51@ss.com", Password: "1234", CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
		{ID: 6, Name: "ABC", Email: "ta61@ss.com", Password: "1234", CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
		{ID: 7, Name: "ABC", Email: "ta71@ss.com", Password: "1234", CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
		{ID: 8, Name: "ABC", Email: "ta81@ss.com", Password: "1234", CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
		{ID: 9, Name: "ABC", Email: "ta91@ss.com", Password: "1234", CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
		{ID: 10, Name: "ABC", Email: "ta110@ss.com", Password: "1234", CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
	}
	result := db.Create(&tmp)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println(result.RowsAffected)
}
