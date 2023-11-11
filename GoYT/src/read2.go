package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// //Name is a struct of names
// type Name struct {
// 	fname []byte
// 	lname []byte
// }

// const prompt = "Please enter the name of the text file: "

// type fr int

// const (
// 	fname fr = iota
// 	lname
// )

// func main() {
// 	fmt.Print(prompt)
// 	scan := bufio.NewScanner(os.Stdin)

// 	valid := scan.Scan()
// 	if valid == false {
// 		fmt.Println("Error. Ending execution.")
// 		return
// 	}
// 	filename := scan.Text()
// 	names := make([]Name, 0)

// 	file, err := os.Open(filename)
// 	if err != nil {
// 		fmt.Println("Error: File not found. Please check the validity of your file name.")
// 		return
// 	}
// 	elt := make([]byte, 1, 1)
// 	name := newName()
// 	var fieldRead fr
// 	n, err := file.Read(elt)
// 	for n > 0 {
// 		if err != nil {
// 			fmt.Println("Error while reading file. Ending execution.")
// 			file.Close()
// 			return
// 		}
// 		if elt[0] == '\n' {
// 			names = append(names, name)
// 			name = newName()
// 			fieldRead = fname
// 		} else if elt[0] == ' ' {
// 			fieldRead = lname
// 		} else {
// 			if fieldRead == fname {
// 				name.fname = append(name.fname, elt[0])
// 			} else {
// 				name.lname = append(name.lname, elt[0])
// 			}
// 		}
// 		n, err = file.Read(elt)
// 	}
// 	names = append(names, name)
// 	fmt.Println("Names read:")
// 	for _, i := range names {
// 		fmt.Println(string(i.fname) + " " + string(i.lname))
// 	}
// 	file.Close()
// 	return
// }

// func newName() Name {
// 	n := Name{
// 		fname: make([]byte, 0, 20),
// 		lname: make([]byte, 0, 20),
// 	}
// 	return n
// }
