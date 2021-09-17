package funcs

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_createFile(t *testing.T) {
	tables := []struct {
		fullUrlFile string
		fileName    string
	}{
		{"https://www.google.com/", "google"},
		{"https://www.google.com/", ""},
		{"https://www.google.com/qwe", ""},
		{"https://www.google.com/qwe", "nf"},
	}

	for _, table := range tables {
		filePath, err := createFile(table.fullUrlFile, table.fileName)

		assert.NoError(t, err)
		assert.Equal(t, filePath, filePath)
		_, err = os.Stat(filePath)
		assert.NoError(t, err)
		err = os.RemoveAll(strings.Split(filePath, "/")[0])
		assert.NoError(t, err)
	}
}

func Test_getResponse(t *testing.T) {
	urls := []string{
		"https://golang.org/",
		"https://www.google.ru/",
		"https://www.google.com/",
		"https://yandex.ru/",
		"https://yandex.com/",
	}

	for _, url := range urls {
		res, err := getResponse(url)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	}
}
