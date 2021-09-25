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
