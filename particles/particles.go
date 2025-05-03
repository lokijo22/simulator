package particles

import (
	"math/rand/v2"
	"simulator/vectors"
)

type Particle2D struct {
	Position vectors.Vector2
	Velocity vectors.Vector2
	Mass     float64
}

func (particle *Particle2D) Randomize() {
	particle.Position = vectors.Vector2{Data: [2]float64{rand.Float64(), rand.Float64()}}
	particle.Velocity = vectors.Vector2{Data: [2]float64{rand.Float64(), rand.Float64()}}
	particle.Mass = rand.Float64()
}
