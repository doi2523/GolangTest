//Giao diện
package main

import (
    "fmt"
    "math"
)

// Định nghĩa giao diện geometry với 2 phương thức: area() và perim().
type geometry interface {
    area() float64  // Tính diện tích
    perim() float64 // Tính chu vi
}

// Định nghĩa kiểu rect (hình chữ nhật) với 2 trường: width (chiều rộng) và height (chiều cao).
type rect struct {
    width, height float64
}

// Định nghĩa kiểu circle (hình tròn) với một trường: radius (bán kính).
type circle struct {
    radius float64
}

// Phương thức tính diện tích của hình chữ nhật.
func (r rect) area() float64 {
    return r.width * r.height
}

// Phương thức tính chu vi của hình chữ nhật.
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

// Phương thức tính diện tích của hình tròn.
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

// Phương thức tính chu vi của hình tròn.
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

// Hàm measure nhận vào đối tượng geometry và in ra diện tích và chu vi của nó.
func measure(g geometry) {
    fmt.Println(g)            // In ra thông tin đối tượng geometry.
    fmt.Println(g.area())     // In ra diện tích của đối tượng.
    fmt.Println(g.perim())    // In ra chu vi của đối tượng.
}

// Hàm detectCircle kiểm tra xem đối tượng geometry có phải là hình tròn không.
func detectCircle(g geometry) {
    // Sử dụng type assertion để kiểm tra nếu đối tượng là hình tròn.
    if c, ok := g.(circle); ok {
        fmt.Println("circle with radius", c.radius) // In ra bán kính nếu đối tượng là hình tròn.
    }
}

func main() {
    // Tạo đối tượng hình chữ nhật và hình tròn.
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    // Đo lường và in kết quả cho hình chữ nhật và hình tròn.
    measure(r)
    measure(c)

    // Kiểm tra xem đối tượng có phải là hình tròn không.
    detectCircle(r)
    detectCircle(c)
}
