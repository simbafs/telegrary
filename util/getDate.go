package util

import (
	"strconv"
	"time"
	log "github.com/sirupsen/logrus"
)

// GetDate returns the current date in the format YYYY MM DD
func GetDate(raw []string) (string, string, string) {
	// convert date from string to int
	var date []int
	for _, v := range raw {
		i, err := strconv.Atoi(v)
		if err != nil {
			break
		}
		date = append(date, i)
	}
	log.Debugln(date)

	year, month, day := time.Now().Date()
	switch len(date) {
	case 3:
		year, month, day = date[0], time.Month(date[1]), date[2]
	case 2:
		month, day = time.Month(date[0]), date[1]
	case 1:
		day = date[0]
	}

	return strconv.Itoa(year), AddZero(int(month)), AddZero(day)
}

