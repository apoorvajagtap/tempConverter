package pkg

import (
	"fmt"
)

func ToCelcius(temperature float64, unit string) {
	var res float64

	if unit == "k" || unit == "kelvin" {
		res = temperature - 273.15
	}

	if unit == "c" || unit == "celcius" {
		res = temperature
	}

	if unit == "f" || unit == "farenhiet" {
		res = (((temperature - 32) * 5) / 9)
	}

	fmt.Printf("%v %s = %v Celcius", temperature, unit, res)
}
