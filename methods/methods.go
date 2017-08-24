package methods

import (
	"fmt"
	"image"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strings"
	"time"
)

// 1 定义结构体的方法
type vertex struct {
	x, y float64
}

func (v *vertex) abs() float64 { // 方法接收者 出现在 func 关键字和方法名之间的参数中
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// 2 可以为任意类型定义方法，不能对来自其他包的类型或基础类型定义方法！！！
type myfloat float64

func (f myfloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// 3 使用指针可以修改原始值
func (v *vertex) scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

// 4 接口 interface
type abser interface {
	abs() float64
}

func classtest() {
	v := &vertex{3, 4}
	fmt.Println(v.abs())
	// 2
	f := myfloat(-math.Sqrt2)
	fmt.Println(f.abs())
	// 3
	v2 := &vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v2, v2.abs())
	v2.scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v2, v2.abs())
	//4
	var a abser
	f4 := myfloat(-math.Sqrt2)
	v4 := vertex{3, 4}
	a = f4
	a = &v4
	fmt.Println(a.abs())
}

// 5 隐士接口
type Reader interface {
	Read(b []byte) (n int, err error)
}
type Writer interface {
	Write(b []byte) (n int, err error)
}
type ReadWriter interface {
	Reader
	Writer
}

func interfacetest() {
	var w Writer
	w = os.Stdout
	fmt.Fprintf(w, "hello world\n")
}

// 6
// 7 练习 Stringers
type Preson struct {
	Name string
	Age  int
}

func (p Preson) String() string { // fmt 就是调用的String 接口！！！！
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
func print_preson() {
	a := Preson{"terry", 15}
	b := Preson{"jack", 25}
	fmt.Println(a, b)
}

// 8 error 的内建接口
// 9 error 练习 skip
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func errortest() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

// 10 Readers
//io.Reader 接口有一个 Read 方法：
//func (T) Read(b []byte) (n int, err error)
//Read 用数据填充指定的字节 slice，并且返回填充的字节数和错误信息。 在遇到数据流结尾时，返回 io.EOF 错误。
//代码创建了一个 strings.Reader。 并且以每次 8 字节的速度读取它的输出
// 11 Readers 练习
// 12 rot13Reader 练习
func readerstest() {
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

// 13 web server
// 通过任何实现了 http.Handler 的值来响应 HTTP 请求：
// package http
// type Handler interface {
//    ServeHTTP(w ResponseWriter, r *Request)
// }

type Hello struct{}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func wbservertest() {
	var h Hello
	err := http.ListenAndServe("localhost:4000", h)
	if err != nil {
		log.Fatal(err)
	}
}

// 15 image
// package image
//type Image interface {
//    ColorModel() color.Model
//    Bounds() Rectangle
//    At(x, y int) color.Color
//}
// 16 image 练习 skip
func imagetest() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

// Out is
func Out() {
	classtest()
	// 5
	interfacetest()
	// 6
	print_preson()
	// 8
	errortest()
	// 10
	readerstest()
	// 13
	//wbservertest()
	// 15
	imagetest()
}
