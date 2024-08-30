package api

import (
	"strconv"
	"strings"
)

func ConvertFractionToDecimal(fraction string) float64 {
	var result float64
	if strings.Contains(fraction, "/") {
		parts := strings.Split(fraction, "/")

		num, _ := strconv.Atoi(parts[0])
		dem, _ := strconv.Atoi(parts[1])

		result = float64(num) / float64(dem)
	} else {
		result, _ = strconv.ParseFloat(fraction, 64)
	}

	return result

}
