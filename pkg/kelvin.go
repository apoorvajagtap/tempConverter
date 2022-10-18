package pkg

import "fmt"

func ToKelvin(temperature float64, unit string) {
	var res float64

	if unit == "k" || unit == "kelvin" {
		res = temperature
	}

	if unit == "c" || unit == "celcius" {
		res = temperature + 273.15
	}

	if unit == "f" || unit == "farenhiet" {
		res = (temperature + 459.67) * 5 / 9
	}

	fmt.Printf("%v %s = %v Kelvin", temperature, unit, res)
}
