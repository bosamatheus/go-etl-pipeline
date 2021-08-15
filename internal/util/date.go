package util

import "time"

const layoutISO = "2006-01-02"

// FormatDateISO returns the date in ISO format (YYYY-MM-DD).
func FormatDateISO(t time.Time) string {
	return t.Format(layoutISO)
}
