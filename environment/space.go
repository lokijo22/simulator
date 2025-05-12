package environment

import "simulator/particles"

// Space classes deal with coordinate logic
type Space2_f32 struct {
	Points []float32
	Size   int
}

type Space2_f64 struct {
	Points []float64
	Size   int
}

// Space which accepts ints
// TODO
// type Space2_int32 struct {
// 	x     []int32
// 	y     []int32
// 	scale int32
// }

// type Space2_int64 struct {
// 	x     []int64
// 	y     []int64
// 	scale int64
// }

// Non-Cartesian Spaces
// TODO

// Input

func (space *Space2_f64) AddParticle(particle particles.Particle2D_64f) {
	vector := particle.Position.Data
	space.Points = append(space.Points, vector[0])
	space.Points = append(space.Points, vector[1])

	space.Size++
}

func (space *Space2_f32) AddParticle(particle particles.Particle2D_32f) {
	vector := particle.Position.Data
	space.Points = append(space.Points, vector[0])
	space.Points = append(space.Points, vector[1])

	space.Size++
}

// Retrive point values
func (space *Space2_f64) GetPoint(index int) (float64, float64) {
	return space.Points[index*2], space.Points[index*2+1]
}

func (space *Space2_f32) GetPoint(index int) (float32, float32) {
	return space.Points[index*2], space.Points[index*2+1]
}
