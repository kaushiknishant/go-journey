package main

import (
	"fmt"
	"os"
	"strings"
)

// taking Args[0] also 
func main(){
	fmt.Println(strings.Join(os.Args[0:]," "))
}