package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var spaceline string
	var velocityRange = 14
	var ticketRange = 14
	const distance = 62100000 // Растояние в км
	var tripType string

	fmt.Printf("\n%-16v %v %v %6v\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Println("======================================")

	for ticket := 0; ticket < 10; ticket++ {

		velocity := rand.Intn(velocityRange) + 16                           // Скорость км/с
		tripDuration := distance / (velocity * 60 * 60 * 24)                // Продолжительность полёта в днях
		ticketPrice := (velocity-16)*100/velocityRange*ticketRange/100 + 36 // Цена в миллионах

		ttype := rand.Intn(2)

		if ttype == 0 {
			tripType = "One-way"
		} else {
			tripType = "Round-trip"
			ticketPrice *= 2
		}

		switch rand.Intn(3) {
		case 0:
			spaceline = "Space Adventures"
		case 1:
			spaceline = "SpaceX"
		case 2:
			spaceline = "Virgin Galactic"
		}

		fmt.Printf("%-16v %4v %-10v %2v %v\n", spaceline, tripDuration, tripType, "$", ticketPrice)
		//fmt.Printf(spaceline, " - ", tripDuration, " - ", velocity, " - ", ticketPrice, " - ", tripType, "\n")
	}

}
