package main

import "fmt"

func main(){
	
	// Tạo mảng gòm 5 giá trị số nguyên a[5]
	var a[5]int
	fmt.Println("emp:", a)

	//Khởi tạo giá trị thứ 4 trong mảng = 100
	a[4] = 100
	fmt.Println("set:", a)
	//in ra giá trị 4 trong mảng theo vị trí a[4]
	fmt.Println("get", a[4])

	//In ra độ dài của mảng dùng len
	fmt.Println("length:", len(a))

	//Khởi tạo mảng b có 5 giá trị nguyên và được gán sẵn
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	//khởi tạo mảng b khi đã khai báo 5 giá trị bên trên
	b = [...]int{1,2,3,4,5}
	fmt.Println("dcl:", b)

	//Khởi tạo mảng b, cho giá trị 200 là giá trị thứ 3 và sẽ chèn giá trị sau
	b = [...]int{100, 3:200, 500}
	fmt.Println("idx:", b)

	var twoD [2][3]int
	for i := 0; i<2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i +j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1,2,3},
		{3,4,5},
	}
	fmt.Println("2d:", twoD)
}