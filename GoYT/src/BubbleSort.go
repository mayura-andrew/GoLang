package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)



func BubbleSort(arr []int){
	n := len(arr)
	for i :=0; i<n-1; i++{
		for j :=0; j<n-1; j++ {
			if arr[j] > arr[j+1]{
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}


func Swap(arr []int, i int){
	arr[i], arr[i+1] = arr[i+1], arr[i]
}

func main(){
	fmt.Println("Enter a sequence of up to 10 integers, separated by spaces :")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	strArr := strings.Split(input, " ")
	var intArr []int
	for _, str := range strArr {
		num, err := strconv.Atoi(str)
		if err != nil {
			continue
		}
		intArr = append(intArr, num)
	}
	BubbleSort(intArr)
	fmt.Println(intArr)
}