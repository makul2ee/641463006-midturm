package main

import (
    "bufio"
    "fmt"
    "net"
    "strings"
)

func main() {
    // แสดงข้อความเริ่มต้น server
    fmt.Println("Starting server...")

    // เปิดการเชื่อมต่อใน port 8080
    ln, err := net.Listen("tcp", ":5000")
    if err != nil {
        panic(err)
    }
    defer ln.Close()

    // ลูปรับการเชื่อมต่อจาก client
    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println(err)
            continue
        }

        // จัดการการเชื่อมต่อแต่ละครั้งใน goroutine แยกต่างหาก
        go handleConnection(conn)
    }
}

// ฟังก์ชันในการจัดการการเชื่อมต่อ
func handleConnection(conn net.Conn) {
    // ปิดการเชื่อมต่อเมื่อจบการทำงาน
    defer conn.Close()

    // อ่านข้อมูลจาก client
    reader := bufio.NewReader(conn)
    username, _ := reader.ReadString('\n')
    password, _ := reader.ReadString('\n')

    // ตัดช่องว่างออกจาก username และ password
    username = strings.TrimSpace(username)
    password = strings.TrimSpace(password)

    // ตรวจสอบข้อมูลว่าตรงกับที่กำหนดหรือไม่
    if username == "std1" && password == "p@ssw0rd" {
        conn.Write([]byte("Hello\n")) // ถ้าถูกต้อง ส่งคำตอบกลับไปว่า "Hello"
    } else {
        conn.Write([]byte("Invalid credentials\n")) // ถ้าไม่ถูกต้อง ส่งคำตอบกลับไปว่า "Invalid credentials"
    }
}
