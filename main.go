package main

import (
	"fmt"
	"simulator/particles"
	"unsafe"
)

func main() {
	// Generate Randomized Particle
	particle := particles.Particle2D{}
	particle.Randomize()

	size := unsafe.Sizeof(particle) + unsafe.Sizeof(particle.Position) + unsafe.Sizeof(particle.Velocity) + unsafe.Sizeof(particle.Position.Data) + unsafe.Sizeof(particle.Velocity.Data)

	fmt.Printf("Particle (%d bytes)\nPosition: %f\nVelocity: %f\nMass: %f", size, particle.Position.Data, particle.Velocity.Data, particle.Mass)
}
