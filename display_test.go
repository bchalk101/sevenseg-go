package sevenseg

import "testing"

func TestDisplayTwoDigits(t *testing.T) {

	display := new(Display)
	display.segment = new(sevenSegMock)
	display.segmentActivePin[0] = new(mockRaspberryPiPinImpl)
	display.segmentActivePin[1] = new(mockRaspberryPiPinImpl)
	display.display("00")

	if display.segment.(*sevenSegMock).number != "0" {
		t.Errorf("Error: Not displaying 0 in first segment")
	}

}
