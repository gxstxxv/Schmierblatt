package app

import (
	"bufio"
	"fmt"
	"os"
)

var path = ""

func addFile() {

}

func readDir() []string {

	entries, err := os.ReadDir(path + "/assets/")
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
	}

	var files []string 

	for _, entry := range entries {
		files = append(files, entry.Name())
	}

	return files

}

func readDesc() []string {

	files := readDir()
	descs := []string{}

	desc_length := FilemenuWidth - 6

	for _, file := range files {

		file, err := os.Open(path+ "/assets/" + file)
		if err != nil {
			fmt.Printf("Error reading descs: %v\n", err)
		}

		scanner := bufio.NewScanner(file)
		
		var desc string

		if scanner.Scan() {
			desc = scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading file: %v\n", err)
		}

		file.Close()

		if len(desc) > desc_length {
			desc = desc[:desc_length-3]
			desc += "..."
		}

		descs = append(descs, desc)
		
	}
	
	return descs

}

func readFile(file string) string {

	data, err := os.ReadFile(path + "/assets/" + file)

	if err != nil {
		fmt.Printf("Error while reading file: %v", err)
	}

	return string(data)

}

func writeFile(data string, file_name string) {

	file, err := os.OpenFile(path+"/assets/" + file_name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

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
