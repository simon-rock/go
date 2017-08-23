package welcome

import (
	"fmt"
	"time"
)

func Out() {
	fmt.Print("welcome to the palyground!\n")
	fmt.Print("time is ", time.Now(), "\n")
}
