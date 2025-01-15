//Hướng kênh
package main

import "fmt"

// Hàm ping gửi thông điệp msg vào kênh pings.
func ping(pings chan<- string, msg string) {
	pings <- msg  // Gửi thông điệp vào kênh pings.
}

// Hàm pong nhận thông điệp từ kênh pings và gửi nó vào kênh pongs.
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings  // Nhận thông điệp từ kênh pings.
	pongs <- msg     // Gửi thông điệp vào kênh pongs.
}

func main() {
	// Tạo hai kênh pings và pongs, mỗi kênh có buffer dung lượng 1.
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	// Gửi thông điệp "passed message" vào kênh pings.
	ping(pings, "passed message")

	// Nhận thông điệp từ kênh pings và gửi nó vào kênh pongs.
	pong(pings, pongs)

	// In thông điệp nhận được từ kênh pongs.
	fmt.Println(<-pongs)
}
