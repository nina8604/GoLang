package main

import "fmt"

func main() {
	defer foo()
	defer bar()
	karaoke()

}

func foo() {
	fmt.Println("Hello from foo")
}
func bar() {
	fmt.Println("Hello from bar")
}
func karaoke() {
	fmt.Println("Hello from karaoke")
}
