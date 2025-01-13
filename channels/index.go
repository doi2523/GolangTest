// Channel là các đường ống kết nối các goroutine đồng thời
package main

import "fmt"

func main() {
    // Tạo một channel kiểu string
    messages := make(chan string)

    // Khởi tạo một goroutine để gửi một thông điệp vào channel
    go func() { 
        messages <- "ping" 
    }()

    // Nhận thông điệp từ channel và lưu vào biến msg
    msg := <-messages

    // In ra thông điệp nhận được
    fmt.Println(msg)
}
