package basics

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// 4
//func add(a int, b int) int { // 函数类型 和 变量类型都在后面， 与c++相反
func add(a, b int) int { // 相同的类型，可以使用省略的方式
	return a + b
}

// 6
func swap(x, y string) (string, string) {
	return y, x
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
}
