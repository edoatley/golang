package main

import "fmt"

func main() {
	a := 1
	fmt.Printf(HelloWorld("Ed"))

	if a >= 1 {
		fmt.Println("a >= 1")
	}
}

func HelloWorld(userName string) string {
	return fmt.Sprintf("Hi, %s ", userName)
}
