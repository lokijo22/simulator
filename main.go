package main

import (
	"fmt"
	"simulator/particles"
)

func main() {
	// Generate Randomized Particle
	particle := particles.Particle2D{}
	particle.Randomize()

	fmt.Printf("Particle\nPosition: %f\nVelocity: %f\nMass: %f", particle.Position.Data, particle.Velocity.Data, particle.Mass)
}
