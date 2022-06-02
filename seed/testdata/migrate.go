package testdata

import (
	"fmt"
	. "forbizbe/main/model"
	. "forbizbe/main/svmn/Mysql"
)

func TableMigrate() {
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Post{})
	fmt.Println("Table Migrated!!")
}
