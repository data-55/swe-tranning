package main

import (
	"flag"
	"fmt"
	"forbizbe/main/svmn/Mysql"
	"forbizbe/seed/testdata"
)

func main() {
	Mysql.ConnDB()

	flag.Parse()
	mode := flag.Arg(0)

	switch mode {
	case "1":
		testdata.TableMigrate()
	case "2":
		testdata.InputData()
	case "3":
		testdata.DeleteData()
	case "4":
		testdata.DeleteData()
		testdata.TableMigrate()
		testdata.InputData()
	default:
		fmt.Println(`
			Please, Choose a mode.
			1 - Table Migrate
			2 - Input Data
			3 - Delete Data
			4 - All
		`)
	}
}
