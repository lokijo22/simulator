package main

import (
	"fmt"
	"simulator/particles"
)

func main() {
	// Generate Randomized Particle
	particle := particles.Particle2D_32f{}
	particle.Randomize()

	fmt.Printf("Particle (%d bytes)\nPosition: %f\nVelocity: %f\nMass: %f", particle.Sizeof(), particle.Position.Data, particle.Velocity.Data, particle.Mass)
}
