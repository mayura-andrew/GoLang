package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type name struct {
    fname string
    lname string
}

func read() {
    var filename string
    fmt.Print("Enter the name of the file: ")
    fmt.Scan(&filename)

    file, err := os.Open(filename)

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    var names []name

    for scanner.Scan() {
        line := scanner.Text()
        fields := strings.Fields(line)

        if len(fields) != 2 {
            fmt.Println("Invalid input:", line)
            continue
        }

        n := name{fname: fields[0], lname: fields[1]}
        names = append(names, n)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error:", err)
        return
    }

    for _, n := range names {
        fmt.Println(n.fname, n.lname)
    }
}
