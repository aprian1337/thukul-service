package helpers

import (
	"strconv"
)

func StringToUint(str string) (uint, error) {
	convInt, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return uint(convInt), nil
}

func FloatToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}
