package main

import "fmt"

// Định nghĩa kiểu cấu trúc base với một trường num kiểu int.
type base struct {
    num int
}

// Phương thức describe() cho kiểu base trả về một chuỗi mô tả giá trị của trường num.
func (b base) describe() string {
    return fmt.Sprintf("base with num=%v", b.num)
}

// Định nghĩa kiểu cấu trúc container, bao gồm một trường base và một trường str kiểu string.
type container struct {
    base   // Embed base vào trong container (lấy lại tất cả các trường và phương thức của base).
    str string
}

func main() {

    // Tạo một đối tượng container co với num = 1 và str = "some name".
    co := container{
        base: base{
            num: 1,  // Gán giá trị cho trường num của base.
        },
        str: "some name", // Gán giá trị cho trường str của container.
    }

    // In ra giá trị của các trường num và str trong container co.
    fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

    // Truy cập trực tiếp vào trường num của base thông qua container co.
    fmt.Println("also num:", co.base.num)

    // Gọi phương thức describe() của base thông qua đối tượng container co.
    fmt.Println("describe:", co.describe())

    // Định nghĩa một interface describer với phương thức describe().
    type describer interface {
        describe() string
    }

    // Tạo một biến d kiểu describer và gán cho nó đối tượng co (container).
    var d describer = co

    // Gọi phương thức describe() thông qua interface describer.
    fmt.Println("describer:", d.describe())
}
