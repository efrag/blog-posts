package age

import (
	"errors"
	"time"
)

func monthsSince(year, month int) (int, error) {
	if year < 0 {
		return 0, errors.New("can't be a negative year")
	}

	if month < 1 || month > 12 {
		return 0, errors.New("invalid month provided")
	}

	currentYear := time.Now().Year()
	if year > currentYear {
		return 0, errors.New("can't be a year in the future")
	}

	currentMonth := int(time.Now().Month())
	if year == currentYear && month > currentMonth {
		return 0, errors.New("can't be a month in the future")
	}

	return (currentYear-year)*12 + currentMonth - month, nil
}
