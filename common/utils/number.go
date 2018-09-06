package utils

import (
	"math"
)

//places为小数位，0<=roundOn<=1 :进位决定值
//fmt.Println(utils.Round(3.444444,1,2))  //3.44
//fmt.Println(utils.Round(-3.444444,1,2))  //-3.44
//fmt.Println(utils.Round(-3.444444,0.4,2))  //-3.45
//fmt.Println(utils.Round(3.0,1,4))  //3
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
