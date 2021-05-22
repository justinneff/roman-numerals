package convert

import (
	"fmt"
	"testing"
)

func TestToRoman(t *testing.T) {
	var tests = []struct {
		i    int
		want string
	}{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{7, "VII"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},
		{40, "XL"},
		{50, "L"},
		{60, "LX"},
		{90, "XC"},
		{100, "C"},
		{110, "CX"},
		{400, "CD"},
		{500, "D"},
		{600, "DC"},
		{900, "CM"},
		{1000, "M"},
		{1100, "MC"},
		{1912, "MCMXII"},
		{2021, "MMXXI"},
		{789, "DCCLXXXIX"},
		{1776, "MDCCLXXVI"},
	}

	for _, tt := range tests {
		testname := fmt.Sprint(tt.i)
		t.Run(testname, func(t *testing.T) {
			result := ToRoman(tt.i)
			if result != tt.want {
				t.Errorf("got %s, want %s", result, tt.want)
			}
		})
	}
}
