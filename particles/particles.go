package particles

import (
	"math/rand/v2"
	"simulator/vectors"
	"unsafe"
)

type Particle_32f struct {
	Dims     int8
	Position vectors.Vector_32f
	Velocity vectors.Vector_32f
	Mass     float64
}

type Particle2D_32f struct {
	Position vectors.Vector2_32f
	Velocity vectors.Vector2_32f
	Mass     float64
}

type Particle3D_32f struct {
	Position vectors.Vector3_32f
	Velocity vectors.Vector3_32f
	Mass     float64
}

type Particle_64f struct {
	Dims     int8
	Position vectors.Vector_64f
	Velocity vectors.Vector_64f
	Mass     float64
}

type Particle2D_64f struct {
	Position vectors.Vector2_64f
	Velocity vectors.Vector2_64f
	Mass     float64
}

type Particle3D_64f struct {
	Position vectors.Vector3_64f
	Velocity vectors.Vector3_64f
	Mass     float64
}

func (particle Particle_64f) Sizeof() uintptr {
	size := unsafe.Sizeof(particle) + unsafe.Sizeof(particle.Position) + unsafe.Sizeof(particle.Velocity) + unsafe.Sizeof(particle.Position.Data) + unsafe.Sizeof(particle.Velocity.Data)
	return size
}

func (particle Particle2D_64f) Sizeof() uintptr {
	size := unsafe.Sizeof(particle) + unsafe.Sizeof(particle.Position) + unsafe.Sizeof(particle.Velocity) + unsafe.Sizeof(particle.Position.Data) + unsafe.Sizeof(particle.Velocity.Data)
	return size
}

func (particle Particle3D_64f) Sizeof() uintptr {
	size := unsafe.Sizeof(particle) + unsafe.Sizeof(particle.Position) + unsafe.Sizeof(particle.Velocity) + unsafe.Sizeof(particle.Position.Data) + unsafe.Sizeof(particle.Velocity.Data)
	return size
}

func (particle Particle_32f) Sizeof() uintptr {
	size := unsafe.Sizeof(particle) + unsafe.Sizeof(particle.Position) + unsafe.Sizeof(particle.Velocity) + unsafe.Sizeof(particle.Position.Data) + unsafe.Sizeof(particle.Velocity.Data)
	return size
}

func (particle Particle2D_32f) Sizeof() uintptr {
	size := unsafe.Sizeof(particle) + unsafe.Sizeof(particle.Position) + unsafe.Sizeof(particle.Velocity) + unsafe.Sizeof(particle.Position.Data) + unsafe.Sizeof(particle.Velocity.Data)
	return size
}

func (particle Particle3D_32f) Sizeof() uintptr {
	size := unsafe.Sizeof(particle) + unsafe.Sizeof(particle.Position) + unsafe.Sizeof(particle.Velocity) + unsafe.Sizeof(particle.Position.Data) + unsafe.Sizeof(particle.Velocity.Data)
	return size
}

func (particle *Particle_32f) Randomize() {
	particle.Position = vectors.Vector_32f{Data: make([]float32, particle.Dims)}
	particle.Velocity = vectors.Vector_32f{Data: make([]float32, particle.Dims)}

	for i := int8(0); i < particle.Dims; i++ {
		particle.Position.Data[i] = rand.Float32()
		particle.Velocity.Data[i] = rand.Float32()
	}

	particle.Mass = rand.Float64()
}

func (particle *Particle2D_32f) Randomize() {
	particle.Position = vectors.Vector2_32f{Data: [2]float32{rand.Float32(), rand.Float32()}}
	particle.Velocity = vectors.Vector2_32f{Data: [2]float32{rand.Float32(), rand.Float32()}}
	particle.Mass = rand.Float64()
}

func (particle *Particle3D_32f) Randomize() {
	particle.Position = vectors.Vector3_32f{Data: [3]float32{rand.Float32(), rand.Float32(), rand.Float32()}}
	particle.Velocity = vectors.Vector3_32f{Data: [3]float32{rand.Float32(), rand.Float32(), rand.Float32()}}
	particle.Mass = rand.Float64()
}

func (particle *Particle_64f) Randomize() {
	particle.Position = vectors.Vector_64f{Data: make([]float64, particle.Dims)}
	particle.Velocity = vectors.Vector_64f{Data: make([]float64, particle.Dims)}

	for i := int8(0); i < particle.Dims; i++ {
		particle.Position.Data[i] = rand.Float64()
		particle.Velocity.Data[i] = rand.Float64()
	}

	particle.Mass = rand.Float64()
}

func (particle *Particle2D_64f) Randomize() {
	particle.Position = vectors.Vector2_64f{Data: [2]float64{rand.Float64(), rand.Float64()}}
	particle.Velocity = vectors.Vector2_64f{Data: [2]float64{rand.Float64(), rand.Float64()}}
	particle.Mass = rand.Float64()
}

func (particle *Particle3D_64f) Randomize() {
	particle.Position = vectors.Vector3_64f{Data: [3]float64{rand.Float64(), rand.Float64(), rand.Float64()}}
	particle.Velocity = vectors.Vector3_64f{Data: [3]float64{rand.Float64(), rand.Float64(), rand.Float64()}}
	particle.Mass = rand.Float64()
}
