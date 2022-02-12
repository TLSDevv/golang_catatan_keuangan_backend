package pkg

import "time"

func StringDateToDateTime(datetimestring string) (*time.Time, error) {
	datetime, err := time.Parse("2006-01-02 15:04:05", datetimestring)
	if err != nil {
		return nil, err
	}
	return &datetime, nil
}
