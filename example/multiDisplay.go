package main

import (
	"github.com/stianeikeland/go-rpio"
	"log"
	"github.com/bchalk101/sevenseg-go"
	"time"
	"strconv"
)

func main() {

	err := rpio.Open()
	if err != nil {
		log.Fatal("Could not open Raspberry Pi GPIO")
	}
	defer rpio.Close()


	digit := sevenseg.NewDisplay(5,6,13,19,26,16,20,21, 12, 7)
	for {
		digit.Print("77")

		//for i := 99; i > 0 ; i-- {
		//	digit.Print(strconv.Itoa(i))
		//	time.Sleep(time.Second )
		//}
	}
}

