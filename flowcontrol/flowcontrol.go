package flowcontrol

import "fmt"
import "math"
import "runtime"
import "time"

// 1  for flowcontrol, 2
// 3 while
// 4
func for_test() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum++
	}
	fmt.Println(sum, " ")
	sum2 := 1 // ":=" create a var , "=" is assign!!!
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2, " ")
	// 3 while
	sum3 := 1
	for sum3 < 1000 {
		sum3 += sum3
	}
	fmt.Println(sum3, " ")
	// 4
	//for {} // return process took too long, 死循环
}

// 5 if test
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// 6 if_test, 7
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { // if 前可以执行简单的语句
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

// 8 求开放练习
// 9 switch, 10
// 11 if then else
func switch_test() {
	fmt.Print("go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.") // if the os is mac, will break, don't print Linux
	case "linux":
		fmt.Println("Linux.")
	default:
		// windows freebsd openbsd plan9
		fmt.Printf("%s.", os)
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("good morning")
	case t.Hour() < 17:
		fmt.Println("good afternoon")
	default:
		fmt.Println("good evening")
	}
}

// 12 defer
// 13 defer stack
func defet_test() {
	defer fmt.Println("world")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("hello")

}

// Out is test
func Out() {
	for_test()
	// 5
	fmt.Println(sqrt(2), sqrt(-4))

	// 6
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
	// 9
	switch_test()
	// 12
	defet_test()
}
