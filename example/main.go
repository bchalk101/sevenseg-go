package main

import (
	"github.com/stianeikeland/go-rpio"
	"log"
	"github.com/bchalk101/sevenseg-go"
	"time"
)

func main() {

	err := rpio.Open()
	if err != nil {
		log.Fatal("Could not open Raspberry Pi GPIO")
	}
	defer rpio.Close()

	digit := sevenseg.NewSevenSeg(5,6,13,19,26,16,20,21)
	for {
		digit.Display(sevenseg.ZERO)
		time.Sleep(time.Second)
		digit.Display(sevenseg.ONE)
		time.Sleep(time.Second)
		digit.Display(sevenseg.TWO)
		time.Sleep(time.Second)
		digit.Display(sevenseg.THREE)
		time.Sleep(time.Second)
		digit.Display(sevenseg.FOUR)
		time.Sleep(time.Second)
		digit.Display(sevenseg.FIVE)
		time.Sleep(time.Second)
		digit.Display(sevenseg.SIX)
		time.Sleep(time.Second)
		digit.Display(sevenseg.SEVEN)
		time.Sleep(time.Second)
		digit.Display(sevenseg.EIGHT)
		time.Sleep(time.Second)
		digit.Display(sevenseg.NINE)
		time.Sleep(time.Second)
	}
}
