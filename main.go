package main

import (
	"math/rand"
	"time"
)

func main() {
	// generate a random seed
	rand.Seed(time.Now().UnixNano())

	a := newAnswer()
	m := newMenu(a)
	m.run()

}
