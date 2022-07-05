func TestIsOver9k (t *testing.T) {
	t.Run("TestOver9000", func(t *testing.T) {
		DB = &MockDB {
			FakeFetch: func(string) (int, error) {
				return 9001, nil
			},
		}
		if !isOver9000() {
			t.Errorf('Test did not return expected results')
		}
	})

	t.Run("TestUnder9000", func(t *testing.T) {
		DB = &MockDB {
			FakeFetch: func(string) (int, error) {
				return 8999, nil
			},
		}
		if isOver9000() {
			t.Errorf("Test did not return expected results")
		}
	})

	t.Run("TestDBError", func(t *testing.T) {
		DB = &MockDB{
			FakeFetch: func(string) (int, error) {
				return 0, fmt.Errorf("unable to connect to database")
			},
		}
		if isOver9000() {
			t.Errorf("Test did not return expected results")
		}
	})
}