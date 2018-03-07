package sevenseg

import "testing"

func TestDisplayTwoDigits(t *testing.T) {

	display := new(Display)
	display.segment[0] = new(sevenSegMock)
	display.segment[1] = new(sevenSegMock)
	display.display("00")

	if display.segment[0].(*sevenSegMock).number != "0" {
		t.Errorf("Error: Not displaying 0 in first segment")
	}

	if display.segment[1].(*sevenSegMock).number != "0" {
		t.Errorf("Error: Not displaying 0 in first segment")
	}

}
