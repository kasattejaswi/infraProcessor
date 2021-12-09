package main

// Importing packages
import "os"

// Function to write data to the file
func WriteToFile(filePath string, data string) {
	// Open the file and passing constants
	// O_APPEND - Appends to the file
	// O_WRONLY - Opens file in write only mode
	// O_CREATE - Creates file if not exists
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)

	// In case of errors, terminate the execution
	if err != nil {
		panic(err)
	}

	// Close the file once passed data is written
	defer f.Close()

	// Writing data to the file
	_, err = f.WriteString(data)

	// In case of errors, terminate the execution
	if err != nil {
		panic(err)
	}
}
