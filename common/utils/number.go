package utils

import (
	"math"
)

func Round(val float64, roundOn float64, places int ) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div < 0 {
		div = -div
	}
	if div >= roundOn {
		if digit > 0 {
			round = math.Ceil(digit)
		}else{
			round = math.Floor(digit)
		}

	} else {
		if digit > 0 {
			round = math.Floor(digit)
		}else{
			round = math.Ceil(digit)
		}
	}
	newVal = round / pow
	return
}
