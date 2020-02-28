package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "adf字"
	fmt.Println(len([]rune(str)))
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9\\p{Han}]+$", str); ok {
		fmt.Println("true")
		return
	}

	fmt.Println("false")
}
