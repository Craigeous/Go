package main

import "fmt"

// this is a comment

func main() {
	fmt.Println("Hello World")
	fmt.Println("1 + 1 =", 1.0+1.0)
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")

	fmt.Print("Enter a temp to convert from Fahrenheit to Celsius: ")
	var far float64
	fmt.Scanf("%f\n", &far)
	var cel float64
	cel = ((far - 32) * 5 / 9)
	fmt.Println(far, " degrees Fahrenheit is ", cel, " degrees Celsius.")

	fmt.Print("Enter the measurement to convert from Feet to Meters: ")
	var ft float64
	fmt.Scanf("%f\n", &ft)
	var mtr float64
	mtr = (ft * 0.3048)
	fmt.Println(ft, "ft is ", mtr, "m")
}
