package main

import (
	"github.com/stianeikeland/go-rpio"
	"log"
	"github.com/bchalk101/sevenseg-go"
	"time"
	"fmt"
)

func main() {

	err := rpio.Open()
	if err != nil {
		log.Fatal("Could not open Raspberry Pi GPIO")
	}
	defer rpio.Close()

	digit := sevenseg.NewSevenSeg(5,6,13,19,26,16,20,21, 12)
	for {
		fmt.Printf("Zero")
		digit.Display(sevenseg.ZERO)
		time.Sleep(time.Second * 2)
		fmt.Printf("One")
		digit.Display(sevenseg.ONE)
		time.Sleep(time.Second * 2)
		fmt.Printf("Two")
		digit.Display(sevenseg.TWO)
		time.Sleep(time.Second * 2)
		fmt.Printf("Three")
		digit.Display(sevenseg.THREE)
		time.Sleep(time.Second * 2)
		fmt.Printf("4")
		digit.Display(sevenseg.FOUR)
		time.Sleep(time.Second * 2)
		fmt.Printf("5")
		digit.Display(sevenseg.FIVE)
		time.Sleep(time.Second * 2)
		fmt.Printf("6")
		digit.Display(sevenseg.SIX)
		time.Sleep(time.Second * 2)
		fmt.Printf("7")
		digit.Display(sevenseg.SEVEN)
		time.Sleep(time.Second * 2)
		fmt.Printf("8")
		digit.Display(sevenseg.EIGHT)
		time.Sleep(time.Second * 2)
		fmt.Printf("9")
		digit.Display(sevenseg.NINE)
		time.Sleep(time.Second * 2)
	}
}
