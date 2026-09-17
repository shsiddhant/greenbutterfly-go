package internal

import (
	"slices"
	"time"
)

type Anniversary struct {
	Date      string
	DaysUntil int
	Number    int
	Title     string
}

func CalculateAnniversaries(jsonFilePath string, nextYear bool) ([]Anniversary, error) {

	days, err := ParseDaysJSON(jsonFilePath)
	if err != nil {
		return nil, err
	}

	now := time.Now().In(loc)
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	dayDuration := 24 * time.Hour
	layout := "January 02, 2006"

	var anniversaries []Anniversary

	for _, day := range days {
		year, date, month := today.Year(), day.Date.Day(), day.Date.Month()

		anniversaryPassed := today.Month() > month ||
			(today.Month() == month && today.Day() > date)

		if anniversaryPassed {
			if !nextYear {
				continue
			}
			year++
		}

		nextAnniversaryDate := time.Date(year, month, date, 0, 0, 0, 0, loc)
		number := nextAnniversaryDate.Year() - day.Date.Year()
		delta := int(nextAnniversaryDate.Sub(today) / dayDuration)
		anniversaries = append(
			anniversaries,
			Anniversary{
				Date:      nextAnniversaryDate.Format(layout),
				DaysUntil: delta,
				Number:    number,
				Title:     day.Title,
			},
		)
	}
	slices.SortFunc(anniversaries, func(a1, a2 Anniversary) int {
		return a1.DaysUntil - a2.DaysUntil
	})
	return anniversaries, nil
}
