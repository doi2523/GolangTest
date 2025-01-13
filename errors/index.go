package main

import (
	"errors"
	"fmt"
)

// Hàm f nhận một tham số int, nếu giá trị là 42 trả về lỗi, 
// nếu không trả về giá trị cộng thêm 3 và không có lỗi.
func f(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Can't work with 42") // Trả về lỗi khi arg = 42
	}
	return arg + 3, nil // Trả về kết quả +3 nếu không có lỗi
}

// Định nghĩa các lỗi tùy chỉnh.
var ErrOutOfTea = fmt.Errorf("no more tea available") // Lỗi hết trà
var ErrPower = fmt.Errorf("can't boil water")        // Lỗi không đun được nước

// Hàm makeTea nhận một tham số int để mô phỏng quá trình pha trà.
// Nếu tham số là 2, trả về lỗi hết trà. Nếu là 4, trả về lỗi không thể đun nước.
func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea // Nếu arg = 2, trả về lỗi hết trà
	} else if arg == 4 {
		// Nếu arg = 4, trả về lỗi không thể đun nước, bọc lỗi ErrPower
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil // Trả về nil nếu không có lỗi
}

func main() {
	// Vòng lặp kiểm tra hàm f với các giá trị 7 và 42.
	for _, i := range []int{7, 42} {
		if r, e := f(i); e != nil {
			// Nếu có lỗi, in ra lỗi
			fmt.Println("f failed:", e)
		} else {
			// Nếu không có lỗi, in ra kết quả
			fmt.Println("f worked:", r)
		}
	}

	// Vòng lặp kiểm tra hàm makeTea với các giá trị từ 0 đến 4.
	for i := 0; i < 5; i++ {
		if err := makeTea(i); err != nil {
			// Kiểm tra lỗi trả về từ makeTea
			if errors.Is(err, ErrOutOfTea) {
				// Nếu lỗi là ErrOutOfTea, in ra thông báo hết trà
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				// Nếu lỗi là ErrPower, in ra thông báo mất điện
				fmt.Println("Now it is dark.")
			} else {
				// Nếu là lỗi khác, in ra lỗi không xác định
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		}
		// Nếu không có lỗi, in ra trà đã sẵn sàng
		fmt.Println("Tea is ready!")
	}
}
