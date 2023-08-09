package main

import (
	"fmt"
	"os"
	"time"
)


func main(){
	fmt.Println("Measuring time in Go")
	start := time.Now()
	var s, sep string 
	for i:= 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	timeElapsed := time.Since(start)
	fmt.Println("The program took", timeElapsed)
}