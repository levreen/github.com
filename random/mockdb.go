type MockDB struct {
	// FakeFetch is used to provide unique test case results
	FakeFetch func(string) (int, error)
}

func (m *MockDB) Fetch(k string) (int, error) {
	if m.FakeFetch != nil {
		return m.FakeFetch(k)
	}
	return 0, nil
}