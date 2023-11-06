package main

import (
	"fmt"
	"os"
	"github.com/mayura-andrew/GoLangLearning/fileio"
)
func main(){
	rootPath, _ := os.Getwd()
	filePath := rootPath + "/data/text.txt"
	c, err := fileio.ReadTextFile(filePath)
	if err == nil {
		fmt.Println(c)

		newContent := fmt.Sprintf("Original: %v\n Double Original: %v%v", c, c, c)
		fileio.WriteToFile(filePath+".output.txt", newContent)
	} else {
		fmt.Printf("Errpr Occured! %v", err)
	}
}

