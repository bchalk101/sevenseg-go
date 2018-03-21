package sevenseg


type sevenSegMock struct {
	displayCalled bool
	number string
	previousNumbers []string
}

func (s *sevenSegMock) Display(item string) error {
	s.displayCalled = true
	s.number = item
	s.previousNumbers = append(s.previousNumbers, item)
	return nil
}
