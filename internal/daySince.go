package internal

import (
	"slices"
	"time"
)

type DaysSince struct {
	Title string
	Date  string
	Days  int
}

func CalculateDaysSince(jsonFilePath string) ([]DaysSince, error) {

	days, err := ParseDaysJSON(jsonFilePath)
	if err != nil {
		return nil, err
	}

	today := time.Now().In(loc)
	cmp := func(d1, d2 Day) int {
		return d1.Date.Compare(d2.Date.Time)
	}

	slices.SortFunc(days, cmp)

	dayDuration := 24 * time.Hour
	layout := "January 02, 2006"

	var daysSince []DaysSince

	for _, day := range days {
		delta := int(today.Sub(day.Date.Time) / dayDuration)
		daysSince = append(
			daysSince,
			DaysSince{Title: day.Title, Date: day.Date.Format(layout), Days: delta},
		)

	}
	return daysSince, nil
}
