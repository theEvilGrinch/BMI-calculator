package main

import (
	"fmt"
	"math"
)

func main() {
	const bmiPower float64 = 2
	var (
		userWeight     float64
		userHeight     float64
		isRepeatNeeded string
	)
	for {
		fmt.Println("\tBody Mass Index Calculator")
		fmt.Print("Enter your height in centimeters (e.g., 180): ")
		_, err := fmt.Scanln(&userHeight)
		if err != nil || userHeight <= 0 {
			fmt.Println("Error: Please enter a positive number greater than zero for your height in centimeters")
			continue
		}
		fmt.Print("Enter your weight in kilograms: ")
		_, err = fmt.Scanln(&userWeight)
		if err != nil || userWeight <= 0 {
			fmt.Println("Error: Please enter a positive number greater than zero for your weight in kilograms")
			continue
		}
		bmi := userWeight / math.Pow(userHeight/100, bmiPower)

		fmt.Printf("Your Body Mass Index: %.1f \n", bmi)
		switch {
		case bmi < 16:
			fmt.Println("(severe thinness)")
		case bmi < 18.5:
			fmt.Println("(underweight)")
		case bmi < 25:
			fmt.Println("(normal weight)")
		case bmi < 30:
			fmt.Println("(overweight)")
		case bmi < 35:
			fmt.Println("(obesity class I)")
		case bmi < 40:
			fmt.Println("(obesity class II)")
		default:
			fmt.Println("(obesity class III, morbid)")
		}
		fmt.Println("Repeat calculate? y - yes, n - no (close program)")
		_, err = fmt.Scanln(&isRepeatNeeded)
		if err != nil {
			fmt.Println("Failed to retrieve value.")
			continue
		}
		switch isRepeatNeeded {
		case "y":
			continue
		case "n":
			return
		default:
			fmt.Println("Incorrect input, y - repeat program, n - close program")
			continue
		}
	}
}
