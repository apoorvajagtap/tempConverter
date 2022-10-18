package pkg

import "fmt"

func ToFarenhiet(temperature float64, unit string) {
	var res float64

	if unit == "k" || unit == "kelvin" {
		res = (temperature-273.15)*9/5 + 32
	}

	if unit == "c" || unit == "celcius" {
		res = temperature*(9/5) + 32
	}

	if unit == "f" || unit == "farenhiet" {
		res = temperature
	}

	fmt.Printf("%v %s = %v Farenhiet", temperature, unit, res)
}
