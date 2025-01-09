package main

import "fmt"

//Tạo hàm mới intSeq bao một hàm khác 
//bên trong và đóng biến tạo thành 1 closure
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {

	nextInt := intSeq()

	//Mỗi lần in ra là một kq riêng biệt, giá trị này
	//sẽ được cập nhật nếu cta gọi nextInt
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
	fmt.Println(nextInt())
}