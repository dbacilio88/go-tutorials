package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	rang()
}

func rang() {
	for i := 0; i < 20; i++ {
		fmt.Println(i)
	}
}
