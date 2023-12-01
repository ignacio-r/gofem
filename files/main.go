package main

import (
	"fmt"
	"os"

	"fem.com/go/files/fileutils"
)

func main() {

	rootPath, _ := os.Getwd()
	filepath := rootPath + "/data/text.txt"
	c, err := fileutils.ReadTextFile(rootPath + "/data/text.txt")
	if err == nil {
		fmt.Println(c)

		newContent := fmt.Sprintf("Original: %v\n Double the Original: %v%v", c, c, c)
		fileutils.WriteToFile(filepath+"output.txt", newContent)
	} else {
		fmt.Printf("ERROR PANIC %v", err)
	}
}
