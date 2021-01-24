package age

import (
	"errors"
	"time"
)

func yearsSince(year int) (int, error) {
	if year < 0 {
		return 0, errors.New("can't be a negative year")
	}

	currentYear := time.Now().Year()

	if year > currentYear {
		return 0, errors.New("can't be a year in the future")
	}

	return currentYear - year, nil
}
