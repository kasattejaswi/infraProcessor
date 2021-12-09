package main

import (
	"errors"
	"os"
	"testing"
)

func check(e error, t *testing.T) {
	if e != nil {
		t.Error(e)
	}
}

func TestWriteToFile(t *testing.T) {
	fileName := "test12ijwr8wj339ju9"
	// Create file if does not exists
	WriteToFile(fileName, "Hello World")
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		t.Error("File is not created though creation function is called")
	}
	// Check data of file
	dat, err := os.ReadFile(fileName)
	check(err, t)
	if string(dat) != "Hello World" {
		t.Error("File created but data is not written")
	}
	// Appending data to existing file
	WriteToFile(fileName, "Hello World")
	// Checking if data is appended or not
	dat, err = os.ReadFile(fileName)
	check(err, t)
	if string(dat) != "Hello WorldHello World" {
		t.Error("Data is not appended to file")
	}
	// Deleting temp file after testing is completed
	e := os.Remove(fileName)
	if e != nil {
		panic(e)
	}
}
