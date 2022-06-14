package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().Format("02.01.2006 15:04")
	fmt.Printf("Hello: The curently time is %s", t)
	fmt.Println("Code run correctly")
}
