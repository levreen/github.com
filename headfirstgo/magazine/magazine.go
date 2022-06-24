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
	Address // Anonymous Struct field.
	Name    string
	Rate    float64
	Active  bool
}

// Employee exported type
type Employee struct {
	Name    string
	Salary  float64
	Address // Anonymous Struct field.
}
