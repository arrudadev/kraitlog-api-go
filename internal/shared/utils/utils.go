package utils

import "time"

func FormatDateTimeUTC(dateTime time.Time) string {
	if dateTime.IsZero() {
		return ""
	}

	return dateTime.Format("2006-01-02T15:04:05Z07:00")
}
