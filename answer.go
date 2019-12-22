package main

import (
	"math/rand"
)

type answer struct {
	chosen animal
	all    []animal
}

// newAnswer returns an answer
func newAnswer() answer {
	var a answer
	a.all = []animal{
		fish(),
		lizard(),
		crab(),
	}
	// generate a random integer based on the length of the slice
	randomInteger := rand.Intn(len(a.all) - 1)
	// assign the chosen animal
	a.chosen = a.all[randomInteger]
	return a
}

func (a answer) name() string {
	return a.chosen.name
}

func (a answer) hasScales() string {
	if a.chosen.hasScales == true {
		return "Yes, this animal has scales."
	}
	return "No, this animal does not have scales."
}

func (a answer) hasTail() string {
	if a.chosen.hasTail == true {
		return "Yes, this animal does have a tail."
	}
	return "No, this animal does not have a tail."
}

func (a answer) hasLegs() string {
	if a.chosen.hasLegs == true {
		return "Yes, this animal does have legs."
	}
	return "No, this animal does not have legs."
}

func (a answer) livesInWater() string {
	if a.chosen.livesInWater == true {
		return "Yes, this animal lives in the water."
	}
	return "No, this animal does not live in the water."
}
