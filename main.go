package main

import (
	"fmt"
	"simulator/environment"
	"simulator/particles"
	"simulator/user"
)

func main() {
	environment := environment.Environment_f64{}

	var quantity int = 10

	for i := 0; i < quantity; i++ {
		particle := particles.Particle2D_64f{}
		particle.Randomize()

		environment.Space.AddParticle(particle)
		fmt.Printf("Particle -  Position:%f \n", particle.Position.Data)
	}

	for i := 0; i < environment.Space.Size; i++ {
		x, y := environment.Space.GetPoint(i)
		fmt.Printf("Item %d: %f %f\n", i, x, y)
	}

	user.CommandLineInterface()
}
