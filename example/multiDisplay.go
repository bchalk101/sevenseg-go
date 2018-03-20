package main

import (
	"github.com/stianeikeland/go-rpio"
	"log"
	//"github.com/bchalk101/sevenseg-go"
)

func main() {

	err := rpio.Open()
	if err != nil {
		log.Fatal("Could not open Raspberry Pi GPIO")
	}
	defer rpio.Close()

<<<<<<< Updated upstream
	//digit := sevenseg.NewDisplay(5,6,13,19,26,16,20,21, 12, 7)
	//for {
		//digit.Print("77")
=======

	digit := sevenseg.NewDisplay(5,6,13,19,26,16,20,21, 12)
	for {
		digit.Print("77")
>>>>>>> Stashed changes

		//for i := 99; i > 0 ; i-- {
		//	digit.Print(strconv.Itoa(i))
		//	time.Sleep(time.Second )
		//}
	//}
}

