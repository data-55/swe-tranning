package testdata

import (
	"fmt"
	. "forbizbe/main/model"
	. "forbizbe/main/svmn/Mysql"
)

func DeleteData() {
	DB.Where("True").Delete(&User{})
	fmt.Println("Data Deleted!")
}
