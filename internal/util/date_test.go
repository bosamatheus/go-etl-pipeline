package util

import (
	"testing"
	"time"
)

func TestFormatDateISO(t *testing.T) {
	tests := []struct {
		name string
		t    time.Time
		want string
	}{
		{"Christmas day", time.Date(2021, time.December, 25, 0, 0, 0, 0, time.UTC), "2021-12-25"},
		{"Leap year", time.Date(2024, time.February, 29, 10, 30, 12, 123, time.UTC), "2024-02-29"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatDateISO(tt.t); got != tt.want {
				t.Errorf("FormatDateISO() = %v, want %v", got, tt.want)
			}
		})
	}
}
