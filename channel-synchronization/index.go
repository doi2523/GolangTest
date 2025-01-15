//Đồng bộ hóa kênh
package main

import (
	"fmt"
	"time"
)

// worker là hàm sẽ được chạy dưới dạng goroutine.
func worker(done chan bool) {
	fmt.Print("working...")    
	time.Sleep(time.Second)       
	fmt.Println("done")              

	// Gửi tín hiệu vào kênh done khi công việc hoàn thành.
	done <- true                       
}

func main() {
	// Tạo một kênh done với buffer kích thước 1.
	done := make(chan bool, 1)        
	// Chạy hàm worker trong một goroutine.
	go worker(done)                   

	// Chờ nhận tín hiệu từ kênh done khi worker hoàn thành công việc.
	<- done                           
}
