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

	digit := sevenseg.NewSevenSeg(5,6,13,19,26,16,20,21, 12)
	for {
		for i := 9; i > 0 ; i-- {
			digit.Display(strconv.Itoa(i))
			time.Sleep(time.Second )
			if i == 0 {
				i = 9
			}
		}
	}
}
