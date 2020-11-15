package main

import (
	"fmt"
	"time"
)

func main()  {
	start := time.Now()
	exec()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func exec() {


}