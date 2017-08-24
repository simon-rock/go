package main

import (
	"fmt"

	"./basics"
	"./flowcontrol"
	"./methods"
	"./moretypes"
	"./welcome"
)

func main() {
	fmt.Print("-----------------\n")
	fmt.Print("-----welcome-----\n")
	fmt.Print("-----------------\n")
	welcome.Out()
	fmt.Print("-----------------\n")
	fmt.Print("-----basics------\n")
	fmt.Print("-----------------\n")
	basics.Out()
	fmt.Print("-----------------\n")
	fmt.Print("-----flowcontrol------\n")
	fmt.Print("-----------------\n")
	flowcontrol.Out()
	fmt.Print("-----------------\n")
	fmt.Print("-----moretypes------\n")
	fmt.Print("-----------------\n")
	moretypes.Out()
	fmt.Print("-----------------\n")
	fmt.Print("-----methods------\n")
	fmt.Print("-----------------\n")
	methods.Out()

}
