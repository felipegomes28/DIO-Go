package main

import "fmt"

const ebuK = 373.0

func main() {
  
	k := ebuK
	c := k - 273.0

	fmt.Printf("A temperatura de ebulição da água em = %g graus Kelvin, temperatura de ebulição da água em = %g graus Celsius.", k, c)
}
