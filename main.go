package main

import "fmt"

func Hello(name, lang string) string {
	prefix := "Hello"
	if name == "" {
		name = "world"
	}
	if lang == "Thai" {
		prefix = "สวัสดี"
	}
	return prefix + " " + name
}

func Add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(Hello("", ""))
	fmt.Println(Hello("", "Thai"))
	fmt.Println(Add(2, 5))
}
