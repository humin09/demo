package main

import (
	"fmt"

	hello "github.com/humin09/helloworld"
)

// Hello is func of hello
func Hello() {
	fmt.Println("demo hello begin")
	hello.Hello()
}

//World is func of world
func World() {
	fmt.Println("demo world begin")
	hello.World()
}
func main() {
	Hello()
	World()
}
