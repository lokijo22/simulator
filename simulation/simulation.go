package simulation

import (
	"simulator/environment"
)

type Simulation_f64 struct {
	Rules       []Rule
	Environment environment.Environment_f64
	History     []environment.Environment_f64
}

type SimulationState struct {
	Environment environment.Environment_f64
}
