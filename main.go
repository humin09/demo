package main

import (
	"fmt"

	"github.com/humin09/demo/example"
	"github.com/humin09/helloworld/hello"
	"rsc.io/quote"
)

// Hello is func of hello
func Hello() {
	fmt.Println("demo hello begin")
	hello.Hello()
	quote.Hello()
	i := example.Add(1, 2)
	fmt.Printf("%d", i)
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
