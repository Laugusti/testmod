package main

import (
	"fmt"
	"log"
	"github.com/Laugusti/testmod/v2"
)

func main() {
	g, err := testmod.Hi("John", "es")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(g)
}
