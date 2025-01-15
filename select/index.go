package main

import (
    "fmt"
    "time"
)

func main() {

    // Tạo hai kênh để truyền thông điệp.
    c1 := make(chan string)
    c2 := make(chan string)

    // Goroutine 1 gửi thông điệp "one" vào kênh c1 sau 1 giây.
    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()

    // Goroutine 2 gửi thông điệp "two" vào kênh c2 sau 2 giây.
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()

    // Sử dụng select để nhận thông điệp từ kênh c1 và c2.
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            // In ra "received one" khi nhận được thông điệp từ c1.
            fmt.Println("received", msg1)  
        case msg2 := <-c2:
            // In ra "received two" khi nhận được thông điệp từ c2.
            fmt.Println("received", msg2)  
        }
    }
}
