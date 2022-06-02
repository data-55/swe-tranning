package Mysql

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"forbizbe/main/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type sweDB struct {
	Name     string
	User     string
	Password string
}

var DB *gorm.DB

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func openFile() (*os.File, error) {
	path := model.FILE_PATH
	file, err := os.Open(path)
	check(err)

	return file, err
}

func infoDB(file *os.File) (sweDB, error) {
	var infoDB sweDB

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case strings.Contains(line, "DB_DATABASE"):
			infoDB.Name = line[strings.Index(line, "=")+1:]
		case strings.Contains(line, "DB_USERNAME"):
			infoDB.User = line[strings.Index(line, "=")+1:]
		case strings.Contains(line, "DB_PASSWORD"):
			infoDB.Password = line[strings.Index(line, "=")+1:]
		}
	}
	err := scanner.Err()
	check(err)

	return infoDB, err
}

func openDB(infoDB sweDB) (*gorm.DB, error) {
	dns := fmt.Sprintf("%s:%s@tcp(auwellnessforbiz-cms-db-container:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", infoDB.User, infoDB.Password, infoDB.Name)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	check(err)

	return db, err
}

func ConnDB() error {
	file, err := openFile()
	defer file.Close()
	infoDB, err := infoDB(file)
	db, err := openDB(infoDB)

	DB = db
	fmt.Printf("Database Connection Success (DB:%s)\n", infoDB.Name)

	return err
}
