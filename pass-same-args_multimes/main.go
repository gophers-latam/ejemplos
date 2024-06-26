package main

import "fmt"

func main() {
	name := "Gophers"
	fmt.Printf(
		"El nombre de esta comunidad de %[1]s es: %[1]s LATAM\n",
		name,
	)

	name, version, year := "Go", 1, 2012
	fmt.Printf(
		"La version %[2]d.0 de %[1]s fue lanzada en %[3]d\n",
		name,
		version,
		year,
	)
}

// pasar mismo argumento varias veces para formar string
