package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"time"
)

const offset = int((5*time.Hour + 30*time.Minute) / time.Second)

var loc = time.FixedZone("IST", offset)

const dateLayout = "2006-01-02"

type CustomDate struct {
	time.Time
}

func (date *CustomDate) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return fmt.Errorf("failed to unmarshal to string: %w", err)
	}
	dt, err := time.ParseInLocation(dateLayout, s, loc)
	if err != nil {
		return fmt.Errorf("failed to parse date: %w", err)
	}
	*date = CustomDate{dt}
	return nil
}

type Day struct {
	Title string     `json:"title"`
	Date  CustomDate `json:"date"`
}

type DaysSince struct {
	Title string
	Date  string
	Days  int
}

func CalculateDaysSince(jsonFilePath string) ([]DaysSince, error) {

	data, err := os.ReadFile(jsonFilePath)
	if err != nil {
		return nil, err
	}

	var days []Day
	err = json.Unmarshal(data, &days)
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
