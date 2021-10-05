package helpers

import (
	"aprian1337/thukul-service/helpers/constants"
	"time"
)

func DateTimeToString(date time.Time) string {
	return date.Format(constants.BirthdayFormat)
}

func DateStringToTime(date string) time.Time {
	conv, err := time.Parse(constants.BirthdayFormat, date)
	if err != nil {
		panic(err)
	}
	return conv
}
