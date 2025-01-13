package main

import "fmt"

// Định nghĩa kiểu ServerState (kiểu int).
// Các giá trị trạng thái sẽ được định nghĩa thông qua hằng số iota.
type ServerState int

const (
	// Sử dụng iota để tạo các giá trị hằng số liên tiếp.
	// Trạng thái ban đầu là StateIdle = 0, sau đó các giá trị tiếp theo tự động tăng dần.
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

// Tạo một map ánh xạ các giá trị ServerState thành tên trạng thái dưới dạng chuỗi.
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

// Phương thức String giúp trả về tên trạng thái dưới dạng chuỗi (từ map stateName).
func (ss ServerState) String() string {
	return stateName[ss]
}

// Hàm transition mô phỏng sự chuyển tiếp giữa các trạng thái.
// Dựa trên trạng thái đầu vào, hàm này sẽ trả về trạng thái tiếp theo.
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		// Nếu trạng thái hiện tại là StateIdle, chuyển sang StateConnected.
		return StateConnected
	case StateConnected, StateRetrying:
		// Nếu trạng thái hiện tại là StateConnected hoặc StateRetrying, chuyển về StateIdle.
		return StateIdle
	case StateError:
		// Nếu trạng thái hiện tại là StateError, giữ nguyên trạng thái StateError.
		return StateError
	default:
		// Nếu gặp trạng thái không xác định, ném ra lỗi.
		panic(fmt.Errorf("unknown state: %s", s))
	}
}

func main() {
	// Bắt đầu với trạng thái StateIdle và chuyển sang trạng thái tiếp theo.
	ns := transition(StateIdle)
	// In ra trạng thái tiếp theo: connected
	fmt.Println(ns)

	// Tiếp tục chuyển trạng thái từ ns (StateConnected) và in ra trạng thái tiếp theo.
	ns2 := transition(ns)
	// In ra trạng thái tiếp theo: idle
	fmt.Println(ns2)
}
