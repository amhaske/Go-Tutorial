package main

import "fmt"

func main() {
	i := 0

	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	for i := 1; i < 5; i++ {
		fmt.Println("i is now", i)
	}

}
