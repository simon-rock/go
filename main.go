package main

import (
	"fmt"

	"./basics"
	"./welcome"
)

func main() {
	fmt.Print("--welcome--\n")
	welcome.Out()
	fmt.Print("--basics--\n")
	basics.Out()
}
