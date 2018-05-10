package main

import "fmt"

type myType struct {
	a     int     //eine Zahl
	b     string  // ein String
	cdefg float64 // Abbildung einer Zahl
}

func main() {
	a :=
		"hallo"
	// Umbruch bei composite literals
	s := []string{
		"eins",
		"zwei",
	}
	// Umbruch bei Funktionsaufruf
	fmt.Println(
		a,
		s,
	)
}
