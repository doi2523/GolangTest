package main

import (
	"fmt"
	"maps"
)

func main() {

	m := make(map[string]int)

	//Tạo map giá trị tương ứng
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	//Truy cập giá trị k3 không tồn tại = 0
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	//Độ dài của map
	fmt.Println("len:", len(m))

	//Lệnh xóa một giá trị khỏi map
	delete(m, "k2")
    fmt.Println("map:", m)

	//Lệnh xóa toàn bộ
	clear(m)
    fmt.Println("map:", m)

	//Kiểm tra khóa m["k2"] có tồn tại trong map ko?
	_, prs := m["k2"]
    fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)

	n2 := map[string]int{"foo": 3, "bar": 2}

	//So sánh 2 map xem có bằng hay không ở đây Equal
    if maps.Equal(n, n2) {
        fmt.Println("n == n2")
    } else {
		fmt.Println("n != n2")
	}
}