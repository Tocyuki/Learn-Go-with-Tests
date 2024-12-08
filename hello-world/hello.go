// https://andmorefine.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world

package main

import "fmt"

func Hello() string {
	return "Hello, world"
}
func main() {
	fmt.Println(Hello())
}
