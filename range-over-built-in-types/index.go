package main

import "fmt"

func main() {

	//Khởi tạo biến chuỗi 3 giá trị
    nums := []int{2, 3, 4}
	//Khai báo biến sum =0
    sum := 0
	//Tạo vòng lặp tính tổng 
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    for k := range kvs {
        fmt.Println("key:", k)
    }

    for i, c := range "go" {
        fmt.Println(i, c)
    }
}