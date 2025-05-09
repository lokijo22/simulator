package main

import (
	"fmt"
	"simulator/environment"
	"simulator/particles"
)

func main() {
	space := environment.Space2_f64{}
	var quantity int = 10

	for i := 0; i < quantity; i++ {
		particle := particles.Particle2D_64f{}
		particle.Randomize()

		space.AddParticle(particle)
		fmt.Printf("Particle -  Position:%f \n", particle.Position.Data)
	}

	for i := 0; i < space.Size; i++ {
		x, y := space.GetPoint(i)
		fmt.Printf("Item %d: %f %f\n", i, x, y)
	}
}
