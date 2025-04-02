package main

import "fmt"

func main() {
	var boiling_water_kelvin float64 = 373.0
	var boiling_water_celsius float64

	boiling_water_celsius = boiling_water_kelvin - 273

	fmt.Printf("A temperatura de ebulição da água em celsius é %g. A temperatura de ebulição da água em Kelvin é %g.", boiling_water_celsius, boiling_water_kelvin)
}
