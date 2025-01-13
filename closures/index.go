package main

import "fmt"

// Tạo hàm intSeq trả về một hàm con (closure).
// Hàm intSeq không trả về một giá trị cố định, mà trả về một hàm khác.
// Hàm con này có thể truy cập và thay đổi biến i được khai báo trong hàm intSeq.
func intSeq() func() int {
    // Khai báo biến i và gán giá trị ban đầu là 0
    i := 0

    // Trả về một hàm ẩn danh (anonymous function) có thể thay đổi giá trị của i
    return func() int {
        i++ // Tăng giá trị của i lên 1 mỗi lần hàm con được gọi
        return i // Trả về giá trị của i sau khi tăng
    }
}

func main() {

    // Gọi intSeq để nhận được một hàm con (closure) và gán vào biến nextInt
    nextInt := intSeq()

    // Mỗi lần gọi nextInt() sẽ nhận giá trị của biến i, sau đó i sẽ được tăng lên 1.
    // Kết quả sẽ là 1, 2, 3 cho lần gọi tiếp theo
    fmt.Println(nextInt()) // In ra: 1
    fmt.Println(nextInt()) // In ra: 2
    fmt.Println(nextInt()) // In ra: 3

    // Gọi lại intSeq để tạo ra một closure mới với biến i được khởi tạo lại từ 0
    newInts := intSeq()

    // nextInts sẽ trả về giá trị đầu tiên của closure mới (i = 1)
    fmt.Println(newInts()) // In ra: 1

    // nextInt vẫn giữ giá trị của biến i là 4 (tiếp tục từ lần gọi cuối cùng), vì nó là một closure riêng biệt
    fmt.Println(nextInt()) // In ra: 4
}
