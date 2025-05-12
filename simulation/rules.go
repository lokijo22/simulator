package simulation

// Rule stores a function and variables to pass into it
type Rule struct {
	ApplyFunc func(state *SimulationState, params ...interface{})
	Params    []interface{}
}

// call the function with the stored variables
func (rule *Rule) Apply(state *SimulationState) {
	if rule.ApplyFunc != nil {
		rule.ApplyFunc(state, rule.Params...)
	}
}

// Adjust the Variables stored in Params
