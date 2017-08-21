package main

import (
	"testing"
	"time"
)

func TestParseUpdate(t *testing.T) {
	now := time.Now()
	tests := []struct {
		s    string
		want time.Time
	}{
		{"7/3 UP", time.Date(now.Year(), 7, 3, 0, 0, 0, 0, time.Local)},
		{"8/21 UP", time.Date(now.Year(), 8, 21, 0, 0, 0, 0, time.Local)},
	}
	for _, tt := range tests {
		update, err := parseUpdate(tt.s)
		if err != nil {
			t.Errorf("parseUpdate(%q) error: %v", tt.s, err)
		}
		if update.Equal(tt.want) {
			t.Errorf("parseUpdate(%q) => %q, want %q", tt.s, update, tt.want)
		}
	}
}
