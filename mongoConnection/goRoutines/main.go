package main

import "fmt"

func main() {

	greet("hello")
	greet("world")

}

func greet(s string) {
	fmt.Println(s)
}