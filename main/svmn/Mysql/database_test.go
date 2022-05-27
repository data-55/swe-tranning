package Mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnDB(t *testing.T) {
	file, err_opneFile := openFile()
	assert.NoError(t, err_opneFile, "File not opened")

	infoDB, err_infoDB := infoDB(file)
	assert.NoError(t, err_infoDB, "Can not take DB information")

	_, err_openDB := openDB(infoDB)
	assert.NoError(t, err_openDB, "Database not opened")
}
