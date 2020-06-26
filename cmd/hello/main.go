package main

import "fmt"

func main() {
	message := SayHelloTo("world")
	fmt.Println(message)
}

func SayHelloTo(name string) string {
	return "Hello " + name + "!!!"
}
