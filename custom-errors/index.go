package main

import (
	"errors"
	"fmt"
)

// Định nghĩa kiểu lỗi tùy chỉnh `argError`
type argError struct {
	arg     int    // Tham số lỗi
	message string // Thông báo lỗi
}

// Phương thức `Error` cho phép `argError` trở thành lỗi (implement interface error)
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

// Hàm `f` trả về một lỗi nếu tham số là 42, nếu không trả về kết quả cộng thêm 3
func f(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"} // Trả về lỗi tùy chỉnh khi arg = 42
	}
	return arg + 3, nil // Trả về giá trị +3 nếu không có lỗi
}

func main() {
	// Gọi hàm f với tham số 43, không có lỗi
	_, err := f(43)

	// Khai báo biến `ae` là con trỏ của loại `argError`
	var ae *argError

	// Kiểm tra xem lỗi có phải là lỗi kiểu `argError` không
	if errors.As(err, &ae) {
		// Nếu là lỗi kiểu `argError`, in thông tin lỗi
		fmt.Println(ae.arg)      // In tham số lỗi
		fmt.Println(ae.message)  // In thông báo lỗi
	} else {
		// Nếu không phải lỗi kiểu `argError`, in thông báo
		fmt.Println("err doesn't match argError")
	}
}
