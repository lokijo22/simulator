package environment

// Evironments store data relevant to Simulation for each point
type Environment_f64 struct {
	Space      Space2_f64
	Data       []float64 // [Attribute1, Attribute2 ... AttributeN, Attribute1 ...]
	Attributes []string
}

func (environment *Environment_f64) AddAttribute(attributeName string) {
	// TODO: dynamically add empty data at the xth positions where x is the index of the attribute
	environment.Attributes = append(environment.Attributes, attributeName)
}

// func (environment Environment_f64) GetDataForPoint(index int) ([]float64)

// func (environment Environment_f64) GetAttributeData(attribute string) ([]float64)

// func (environment Environment_f64) HasAttribute(index int) bool

// func (environment Environment_f64) NumberOfPoints() int

// Gets the number of bytes each point uses
// func (environment Environment_f64) SizeOfPoint() int
