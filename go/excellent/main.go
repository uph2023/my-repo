package main

import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("Hell, %s!", name)
}

func main() {
	fmt.Println(Greet("World"))
}
