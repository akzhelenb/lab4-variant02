package main

import (
	"fmt"

	"lab4-variant02/pkg/wallpapercalc"

	"github.com/fatih/color"
	"github.com/google/uuid"
)

func main() {
	length := 5.0
	width := 4.0
	height := 2.7

	rollWidth := 1.06
	rollLength := 10.0
	rollPrice := 1500.0

	perimeter, err := wallpapercalc.Perimeter(length, width)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	rolls, err := wallpapercalc.RollsNeeded(perimeter, height, rollWidth, rollLength)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = wallpapercalc.AddExtraRolls(&rolls, 1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	report, err := wallpapercalc.FormatWallpaperReport("Living Room", rolls, rollPrice)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	id := uuid.New()
	color.Cyan("Calculation ID: %s", id.String())

	fmt.Printf("Perimeter: %.2f\n", perimeter)
	fmt.Printf("Rolls: %d\n\n", rolls)

	color.Green(report)
}
