package main

type animal struct {
	name         string
	hasScales    bool
	hasTail      bool
	hasLegs      bool
	livesInWater bool
}

func fish() animal {
	return animal{
		name:         "fish",
		hasScales:    true,
		hasTail:      true,
		hasLegs:      false,
		livesInWater: true,
	}
}

func lizard() animal {
	return animal{
		name:         "lizard",
		hasScales:    true,
		hasTail:      true,
		hasLegs:      true,
		livesInWater: false,
	}
}

func crab() animal {
	return animal{
		name:         "crab",
		hasScales:    false,
		hasTail:      false,
		hasLegs:      true,
		livesInWater: true,
	}
}
