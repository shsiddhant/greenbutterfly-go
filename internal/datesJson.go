package internal

import (
	"encoding/json"
	"fmt"
	"os"
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

type Ordering string

func ParseDaysJSON(
	jsonFilePath string,
) ([]Day, error) {
	data, err := os.ReadFile(jsonFilePath)
	if err != nil {
		return nil, err
	}
	var days []Day
	err = json.Unmarshal(data, &days)
	if err != nil {
		return nil, err
	}
	return days, nil
}
