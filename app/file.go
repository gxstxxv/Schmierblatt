package app

import (
	"fmt"
	"os"
)

var path = ""

func readFile() string {

	data, err := os.ReadFile(path + "/assets/schmierblatt.txt")

	if err != nil {
		fmt.Printf("Error while reading file: %v", err)
	}

	return string(data)

}

func writeFile(data string) {

	file, err := os.OpenFile(path+"/assets/schmierblatt.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Printf("Error while opening file: %v", err)
	}

	if err := file.Truncate(0); err != nil {
		fmt.Printf("Error while truncate: %v", err)
	}

	if _, err := file.WriteString(data); err != nil {
		fmt.Printf("Error while writing file: %v", err)
	}

	if err := file.Close(); err != nil {
		fmt.Printf("Error while closing file: %v", err)
	}

}
