package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strings"
)

func main() {
    // เชื่อมต่อไปยัง server ที่ localhost port 8080
    conn, err := net.Dial("tcp", "localhost:5000")
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    // รับข้อมูล username จากผู้ใช้
    fmt.Print("Enter username: ")
    username := readInput()
    // รับข้อมูล password จากผู้ใช้
    fmt.Print("Enter password: ")
    password := readInput()

    // ส่ง username และ password ไปยัง server
    fmt.Fprintln(conn, username)
    fmt.Fprintln(conn, password)

    // รอคำตอบจาก server
    response, err := bufio.NewReader(conn).ReadString('\n')
    if err != nil {
        panic(err)
    }

    // แสดงคำตอบจาก server
    fmt.Print("Server response: " + response)
}

// ฟังก์ชันในการอ่านข้อมูลจาก command line
func readInput() string {
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    return strings.TrimSpace(input)
}
