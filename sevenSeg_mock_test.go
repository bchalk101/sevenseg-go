package sevenseg


type sevenSegMock struct {
	displayCalled bool
	number string
}

func (s *sevenSegMock) Display(item string) error {
	s.displayCalled = true
	s.number = item
	return nil
}
