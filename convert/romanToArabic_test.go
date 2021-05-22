package convert

import (
	"fmt"
	"testing"
)

func TestCanConvertToArabic(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"I", true},
		{"V", true},
		{"X", true},
		{"L", true},
		{"C", true},
		{"D", true},
		{"M", true},
		{"XIV", true},
		{"XIVT", false},
		{"ERR", false},
	}

	for _, tt := range tests {
		testname := fmt.Sprint(tt.s)
		t.Run(testname, func(t *testing.T) {
			result := CanConvertToArabic(tt.s)
			if result != tt.want {
				t.Errorf("got %t, want %t", result, tt.want)
			}
		})
	}
}

func TestToArabic(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"I", 1},
		{"V", 5},
		{"X", 10},
		{"L", 50},
		{"C", 100},
		{"D", 500},
		{"M", 1000},
		{"IX", 9},
		{"XIV", 14},
		{"MDCCLXXVI", 1776},
		{"MCMXVIII", 1918},
	}

	for _, tt := range tests {
		testname := fmt.Sprint(tt.s)
		t.Run(testname, func(t *testing.T) {
			result := ToArabic(tt.s)
			if result != tt.want {
				t.Errorf("got %d, want %d", result, tt.want)
			}
		})
	}
}
