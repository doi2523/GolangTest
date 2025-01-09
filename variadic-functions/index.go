package main

import "fmt"

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0

	//Tạo vòng lặp tính tổng giá trị có trong nums với range
	for _, num := range nums{
		total += num
	}
	fmt.Println(total)
}
func main()  {
	sum(1,2)
	sum(1,2,3)
	sum(4,5,6)

	nums := []int{1,2,3,4}
	sum(nums...)
}