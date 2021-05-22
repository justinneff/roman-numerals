package convert

type ArabicToRomanFactor struct {
	Base     int
	Numerals []string
}

var factors = []ArabicToRomanFactor{
	{1000, []string{"", "M", "MM", "MMM"}},
	{100, []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}},
	{10, []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}},
	{1, []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}},
}

func ToRoman(input int) string {
	result := ""
	remainder := input

	for _, fact := range factors {
		num := remainder / fact.Base
		result += fact.Numerals[num]
		remainder -= fact.Base * num
	}

	return result
}
