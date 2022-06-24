package magazine

// Address exported type
type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

// Subscriber exported type
type Subscriber struct {
	HomeAddress Address
	Name        string
	Rate        float64
	Active      bool
}

// Employee exported type
type Employee struct {
	Name        string
	Salary      float64
	HomeAddress Address
}
