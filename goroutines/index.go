//Goroutine là một luồng thực thi nhẹ.
package main

import (
	"fmt"
	"time"
)

// Hàm f nhận vào một chuỗi `from` và in nó cùng với chỉ số từ 0 đến 2
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// Gọi hàm f một cách trực tiếp (chạy trên luồng chính)
	f("direct")

	// Khởi tạo một goroutine để chạy hàm f song song với luồng chính
	go f("goroutine")

	// Khởi tạo một goroutine khác với một hàm ẩn danh (anonymous function)
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Đợi một chút để các goroutines hoàn thành
	time.Sleep(time.Second)

	// In ra thông báo "done" sau khi goroutines đã chạy xong
	fmt.Println("done")
}
