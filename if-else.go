package main

import "fmt"

func main() {

	flag := true

	if flag {
		fmt.Println("flag is true.")
		flag = false
	}

	if flag {
		fmt.Println("flag is true.")
	} else {
		fmt.Println("flag is false.")
	}

}
