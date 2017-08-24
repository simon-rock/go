package moretypes

import (
	"fmt"
	"math"
	"strings"
)

// 1 point test
func point_test() {
	i, j := 42, 2701
	// k := true
	p := &i
	tobool := false
	const f = "%T(%v)\n"
	fmt.Printf(f, tobool, tobool)

	fmt.Printf(f, p, *p)
	*p = 21
	fmt.Printf(f, p, *p)
	p = &j
	fmt.Printf(f, p, *p)
	//p = &k // error, type of p is int*
}

// 2 3 4 5 struct
type vertex struct {
	x int
	y int
}

func structtest() {
	fmt.Println(vertex{1, 2})

	v := vertex{1, 2}
	v.x = 10 // like c,
	fmt.Println(v)

	p := v
	p.x = 20
	fmt.Println(p)

	const f = "%T(%v)\n"
	v1 := vertex{1, 2}
	v2 := vertex{x: 2}
	v3 := vertex{}
	p2 := &vertex{1, 2}
	fmt.Println(v1, v2, v3, p2)
	fmt.Printf(f, v2, v2)
	fmt.Printf(f, p2, *p2) // type of p2 is *moretypes.vertex
}

// 6 数组
func arraytest() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}

// 7 slice
// 8 slice can inlcude another slice
// 9 process slice
// 10 create slice
// 11 empty slice
// 12 append slice, like vector
// 13 for to map slice or map
// 14 for _, v :=range a{}
// 15 练习 slice， 跳过
func slicetest() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%T  s[%d] == %d\n", s, i, s[i])
	}

	game := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	game[0][0] = "x"
	game[2][2] = "o"
	game[2][0] = "x"
	game[1][0] = "o"
	game[0][2] = "x"
	printboard(game)

	// 9 // 下标和python 一样，左闭右开
	fmt.Println("s[1:4] ==", s[1:4])
	fmt.Println("s[:3]", s[:3])
	fmt.Println("s[4:] ==", s[4:])
	// 10 create slice
	a := make([]int, 5)
	a[3] = 10
	printslice("a", a)
	b := make([]int, 0, 5)
	printslice("b", b)
	c := a[:2]
	printslice("c", c) // [0 0]
	d := c[2:5]
	printslice("d", d) // [0 10 0]
	// 11
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil { // slice is nil
		fmt.Println("nil")
	}

	// 12
	printslice("a", a)
	// append works on nil slices.
	a = append(a, 0)
	printslice("a", a)
	// the slice grows as needed.
	a = append(a, 1)
	printslice("a", a)
	// we can add more than one element at a time.
	a = append(a, 2, 3, 4)
	printslice("a", a)

	// 13
	for i, v := range a {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	// 14
	for i := range a {
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n")
	for _, v := range a {
		fmt.Printf("%d ", v)
	}
}

// 16 map
// 17
// 18
// 19 change map
// 20 a example named wordcount
type mapvertex struct {
	lat, long float64
}

var m map[string]mapvertex // map<string, mapvertex> m;
func maptest() {
	// 16
	fmt.Println("")
	m = make(map[string]mapvertex)
	m["bell labs"] = mapvertex{40.9, -74.3}
	fmt.Println(m["bell labs"])
	// 17
	var m2 = map[string]mapvertex{
		"bell labs": mapvertex{40.5, -74.3},
		"google":    mapvertex{20.5, -122.3},
	}
	fmt.Println(m2)
	// 18
	var m3 = map[string]mapvertex{
		"bell labs": {40.5, -74.3},
		"google":    {20.5, -122.3},
	}
	fmt.Println(m3)

	// 19
	m4 := make(map[string]int)
	m4["Answer"] = 42                       // insert a item
	fmt.Println("The value:", m4["Answer"]) // get item form map
	m4["Answer"] = 48
	fmt.Println("The value:", m4["Answer"])
	delete(m4, "Answer") // delete a item from map
	fmt.Println("The value:", m4["Answer"])
	v, ok := m4["Answer"] // check a item if exist
	fmt.Println("The value:", v, "Present?", ok)
}
func printslice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
func printboard(s [][]string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s\n", strings.Join(s[i], " ")) // string.Join 将一个数组合并
	}
}

// 21 make func as param or return
// 22 闭包 , 感觉类似于函数对象， 即函数带有状态
// 23 练习 斐波那契
func computer(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int { // 此处为匿名函数声明
		sum += x
		return sum
	}
}
func functest() {
	hyhot := func(x, y float64) float64 { // 此为匿名函数声明
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hyhot(5, 12))
	fmt.Println(computer(hyhot))
	fmt.Println(computer(math.Pow))

	// 22
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}

// Out is
func Out() {
	// 1
	point_test()
	// 2
	structtest()
	// 6
	arraytest()
	// 7
	slicetest()
	// 16
	maptest()
	// 21
	functest()
}
