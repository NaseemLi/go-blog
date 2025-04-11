package main

import (
	"fmt"
	"goblog/utils/pwd"
)

func main() {
	p, _ := pwd.GenerateFromPassword("123456")
	fmt.Println(p)
	fmt.Println(pwd.CompareHashAndPassword(p, "123456"))
}
