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

const (
	MinArabic = 1
	MaxArabic = 3999
)

func CanConvertToRoman(arabic int) bool {
	return arabic >= MinArabic && arabic <= MaxArabic
}

func ToRoman(arabic int) string {
	result := ""
	remainder := arabic

	for _, fact := range factors {
		num := remainder / fact.Base
		result += fact.Numerals[num]
		remainder -= fact.Base * num
	}

	return result
}
