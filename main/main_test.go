package main

import (
	"forbizbe/main/svmn/Gin"
	"forbizbe/main/svmn/Mysql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	err_mysql := Mysql.ConnDB()
	assert.NoErrorf(t, err_mysql, "Mysql is not setup!\n%s\n", err_mysql)

	r := Gin.SetupRouter()
	assert.NotNil(t, r, "Gin is not setup!")
}
