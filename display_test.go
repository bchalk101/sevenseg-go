package sevenseg

import (
	"testing"
	"time"
)

func TestDisplayTwoDigits(t *testing.T) {

	display := new(Display)
	display.segment = new(sevenSegMock)
	display.segmentActivePin[0] = new(mockRaspberryPiPinImpl)
	display.segmentActivePin[1] = new(mockRaspberryPiPinImpl)
	display.Print("99")

	if display.segment.(*sevenSegMock).number != "9" {
		t.Errorf("Error: Not displaying 9 in first segment")
	}

	if display.segmentActivePin[0].(*mockRaspberryPiPinImpl).PinWasOnceHigh {
		t.Errorf("Error: First Segment never turned on")
	}
	if display.segmentActivePin[1].(*mockRaspberryPiPinImpl).PinWasOnceHigh {
		t.Errorf("Error: Second segment never turned on")
	}

}

func TestDisplayTwoNumbersAfterEachOther(t *testing.T) {
	display := SetupNewTestDisplay()
	display.Print("99")

	display.Print("88")
	time.Sleep(time.Second)
	if display.segment.(*sevenSegMock).number != "8" {
		t.Errorf("Error: Not displaying 8 in first segment")
	}
}

func SetupNewTestDisplay() *Display {
	display := new(Display)
	display.segment = new(sevenSegMock)
	display.segmentActivePin[0] = new(mockRaspberryPiPinImpl)
	display.segmentActivePin[1] = new(mockRaspberryPiPinImpl)
	display.setupDisplayChannel()

	go display.display()
	return display
}
