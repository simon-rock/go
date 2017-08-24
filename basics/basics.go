package basics

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

// 4
//func add(a int, b int) int { // 函数类型 和 变量类型都在后面， 与c++相反
func add(a, b int) int { // 相同的类型，可以使用省略的方式
	return a + b
}

// 6
func swap(x, y string) (string, string) { // 可以返回多个值
	return y, x
}

// 7
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // 没有参数的 return 语句返回各个返回变量的当前值。这种用法被称作“裸”返回。 same to return x,y
}

// 8, 9, 10
var c, python, java bool
var k, j int = 1, 2

func var_test() {
	var i int
	var cpp, cxx = true, false // 初始化赋值，不声明变量类型
	ruby := 33                 // 		函数外的每个语句都必须以关键字开始（ var 、 func 、等等）， := 结构不能使用在函数外, 函数内可以替代var 定义
	fmt.Println(i, c, python, java, k, j, cpp, cxx, ruby)
}

// 11
var (
	tobool bool       = false
	word   string     = "go world"
	maxint uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)

// bool
// string
// int  int8  int16  int32  int64
//uint uint8 uint16 uint32 uint64 uintptr
// int uint uintptr is 32bit in x86, is 64bit in x64 !!!
// byte // uint8 的别名
//rune // int32 的别名 // 代表一个Unicode码
//float32 float64
//complex64 complex128
)

// 12
// 变量初始化会赋空值 数字型-> 0, bool false, string ""
func basic_type() {
	const f = "%T(%v)\n"
	fmt.Printf(f, tobool, tobool)
	fmt.Printf(f, maxint, maxint)
	fmt.Printf(f, z, z)
}

// 13 类型转换, 必须显示转换
// 14 赋值时改变类型, 右值推导类型
// 15 常量使用const
func cast_test() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Print("\n")
	fmt.Println(x, y, z)

	v := 42 // change me! // 42.5 or 0.867 + 0.5i
	fmt.Printf("v is of type %T\n", v)

	const work = "haha" // 常量, 不能使用:= 赋值!!
}

func Out() {

	// 2
	rand.Seed(time.Now().Unix())
	fmt.Println("My favorite number is", rand.Intn(100))

	// 3
	fmt.Println(math.Pi) // 首字母大写的名称是被导出的， 所以math.pi 是不会被导出的

	// 4
	fmt.Println(add(42, 13))

	// 6
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	// 7
	fmt.Println(split(17))
	// 8
	var_test()
	// 11
	basic_type()
	// 13
	cast_test()
}
