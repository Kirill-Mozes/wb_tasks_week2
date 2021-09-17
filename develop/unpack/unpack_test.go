package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnpackingString(t *testing.T) {
	tables := []struct {
		s   string
		res string
		err error
	}{
		{"a4bc2d5e", "aaaabccddddde", nil},
		{"abcd", "abcd", nil},
		{"45", "", errors.New("invalid string")},
		{"", "", nil},
		{`qwe\4\5`, "qwe45", nil},
		{`qwe\45`, "qwe44444", nil},
		{`qwe\\5`, `qwe\\\\\`, nil},
		{`q\413`, "q4444444444444", nil},
		{"b2c12", "bbcccccccccccc", nil},
		{"b2c12", "bbcccccccccccc", nil},
		{"ф0", "", errors.New("invalid string")},
	}
	for _, table := range tables {
		totalS, totalErr := Unpacking(table.s)

		assert.Equal(t, totalS, table.res)
		assert.Equal(t, totalErr, table.err)
	}
}
