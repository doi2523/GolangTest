// Bộ đệm kênh
package main

import "fmt"

func main() {
	// Tạo một kênh có bộ đệm, có thể chứa tối đa 2 giá trị kiểu string
	messages := make(chan string, 2)

	// Gửi hai thông điệp vào kênh mà không cần phải có người nhận ngay lập tức
	messages <- "buffered"
	messages <- "channel"

	// Nhận và in các thông điệp từ kênh
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
