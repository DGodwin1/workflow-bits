package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(GenerateFood([]string{"beans", "rice"}))
}

func GenerateHello(name string) string {
	return fmt.Sprintf("hello %s", name)
}

func GenerateFood(ingredients []string) string{
	return strings.Join(ingredients, "+")
}