//Phạm vi trên các Trình lặp lại - Iterator
package main

import (
	"fmt"
	"iter"
	"slices"
)

// Kiểu element đại diện cho phần tử trong danh sách liên kết
type element[T any] struct {
	next *element[T] // Trỏ đến phần tử tiếp theo
	val  T           // Giá trị của phần tử
}

// Kiểu List đại diện cho danh sách liên kết
type List[T any] struct {
	head, tail *element[T] // Đầu và cuối danh sách
}

// Phương thức Push thêm phần tử vào cuối danh sách
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// Phương thức All trả về iterator để duyệt qua tất cả phần tử trong danh sách
func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

// Hàm genFib tạo chuỗi số Fibonacci dưới dạng iterator
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1
		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	// Tạo danh sách liên kết và thêm các phần tử
	lst := List[int]{}
	lst.Push(10)
	lst.Push(23)
	lst.Push(45)
	
	// Duyệt và in các phần tử trong danh sách
	for e := range lst.All() {
		fmt.Println(e)
	}
	
	// Thu thập các phần tử của danh sách vào slice
	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	// In các số Fibonacci nhỏ hơn 10
	for n := range genFib() {
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}
