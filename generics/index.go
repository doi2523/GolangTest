package main

import "fmt"

// Hàm SlicesIndex tìm chỉ số của giá trị v trong slice s, nếu không tìm thấy trả về -1.
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
    for i := range s {
        if v == s[i] {
            return i // Trả về chỉ số nếu tìm thấy
        }
    }
    return -1 // Trả về -1 nếu không tìm thấy
}

// Kiểu element đại diện cho phần tử trong danh sách liên kết với giá trị và con trỏ đến phần tử tiếp theo.
type element[T any] struct {
    next *element[T] // Trỏ đến phần tử tiếp theo
    val  T           // Giá trị của phần tử
}

// Kiểu List đại diện cho danh sách liên kết, bao gồm con trỏ đến phần tử đầu và cuối.
type List[T any] struct {
    head, tail *element[T] // Con trỏ đến phần tử đầu và cuối
}

// Phương thức Push thêm phần tử v vào cuối danh sách.
func (lst *List[T]) Push(v T) {
    if lst.tail == nil {
        lst.head = &element[T]{val: v} // Tạo phần tử đầu tiên
        lst.tail = lst.head            // Cập nhật tail
    } else {
        lst.tail.next = &element[T]{val: v} // Thêm phần tử vào cuối
        lst.tail = lst.tail.next            // Cập nhật tail
    }
}

// Phương thức AllElements trả về slice chứa tất cả các phần tử trong danh sách.
func (lst *List[T]) AllElements() []T {
    var elems []T
    for e := lst.head; e != nil; e = e.next {
        elems = append(elems, e.val) // Thêm giá trị phần tử vào slice
    }
    return elems
}

func main() {
    // Sử dụng hàm SlicesIndex để tìm chỉ số phần tử "zoo" trong slice s.
    var s = []string{"foo", "bar","cat", "zoo"}
    fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

    // Tạo danh sách liên kết và thêm các phần tử vào danh sách.
    lst := List[int]{}
    lst.Push(10)
    lst.Push(13)
    lst.Push(23)
    lst.Push(45)
    fmt.Println("list:", lst.AllElements()) // In tất cả các phần tử trong danh sách
}
