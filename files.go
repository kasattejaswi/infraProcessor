package main

import "os"

func WriteToFile(filePath string, data string) {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString(data)
	if err != nil {
		panic(err)
	}
}
