package particles

import (
	"math/rand/v2"
	"simulator/vectors"
)

type Particle struct {
	Dims     int8
	Position vectors.Vector
	Velocity vectors.Vector
	Mass     float64
}

type Particle2D struct {
	Position vectors.Vector2
	Velocity vectors.Vector2
	Mass     float64
}

type Particle3D struct {
	Position vectors.Vector3
	Velocity vectors.Vector3
	Mass     float64
}

func (particle *Particle) Randomize() {
	particle.Position = vectors.Vector{Data: make([]float64, particle.Dims)}
	particle.Velocity = vectors.Vector{Data: make([]float64, particle.Dims)}

	for i := int8(0); i < particle.Dims; i++ {
		particle.Position.Data[i] = rand.Float64()
		particle.Velocity.Data[i] = rand.Float64()
	}

	particle.Mass = rand.Float64()
}

func (particle *Particle2D) Randomize() {
	particle.Position = vectors.Vector2{Data: [2]float64{rand.Float64(), rand.Float64()}}
	particle.Velocity = vectors.Vector2{Data: [2]float64{rand.Float64(), rand.Float64()}}
	particle.Mass = rand.Float64()
}

func (particle *Particle3D) Randomize() {
	particle.Position = vectors.Vector3{Data: [3]float64{rand.Float64(), rand.Float64(), rand.Float64()}}
	particle.Velocity = vectors.Vector3{Data: [3]float64{rand.Float64(), rand.Float64(), rand.Float64()}}
	particle.Mass = rand.Float64()
}
