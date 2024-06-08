package main

import (
	"fmt"

	"golang.org/x/text/language"
)

func main() {
	mx := language.Make("es-MX")
	en := language.Make("en")
	fmt.Println(mx.Region())
	fmt.Println(en.Region())
}
