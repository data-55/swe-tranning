package testdata

import (
	"fmt"
	. "forbizbe/main/model"
	. "forbizbe/main/svmn/Mysql"
	"math/rand"
	"strconv"
	"time"
)

func InputData() {
	// Input users
	user := []User{
		{ID: 1, Name: "1_name", Email: "1_email@test.com", Password: "1234"},
		{ID: 2, Name: "2_name", Email: "2_email@test.com", Password: "1234"},
		{ID: 3, Name: "3_name", Email: "3_email@test.com", Password: "1234"},
		{ID: 4, Name: "4_name", Email: "4_email@test.com", Password: "1234"},
		{ID: 5, Name: "5_name", Email: "5_email@test.com", Password: "1234"},
		{ID: 6, Name: "6_name", Email: "6_email@test.com", Password: "1234"},
		{ID: 7, Name: "7_name", Email: "7_email@test.com", Password: "1234"},
		{ID: 8, Name: "8_name", Email: "8_email@test.com", Password: "1234"},
		{ID: 9, Name: "9_name", Email: "9_email@test.com", Password: "1234"},
		{ID: 10, Name: "10_name", Email: "10_email@test.com", Password: "1234"},
	}
	DB.Create(&user)

	// Input user_follows
	follow1 := []User{{ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}, {ID: 6}}
	DB.Find(&user, 1).Association("Follows").Append(follow1)

	follow2 := []User{{ID: 3}, {ID: 4}, {ID: 5}, {ID: 6}, {ID: 7}}
	DB.Find(&user, 2).Association("Follows").Append(follow2)

	// Input posts
	var cnt uint = 0
	var maxCnt int64

	DB.Find(&user).Count(&maxCnt)

	rand.Seed(100)
	for i := int64(1); i < maxCnt+1; i++ {
		for j := int64(0); j < i; j++ {
			cnt++
			time := time.Now().AddDate(0, rand.Intn(20), 0)
			DB.Find(&user, i).Association("Posts").Append(&Post{
				ID:        cnt,
				Comment:   strconv.FormatInt(i, 10) + "comment",
				UpdatedAt: &time,
				CreatedAt: &time,
			})
		}
	}

	fmt.Println("Data Inputed!")
}
