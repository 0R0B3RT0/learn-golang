package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var expected = []string{"/tmp/file1", "/tmp/file2"}

func TestCreateFile(t *testing.T) {
	files := CreateFile("/tmp", "file1", "file2")

	assert.EqualValues(t, expected, files, "they should be equals")

	for _, file := range files {
		assert.FileExists(t, file, "they should be exists")
	}
}
