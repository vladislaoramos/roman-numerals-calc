package solution

import (
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func isValidNumber(romanNumber string) (bool, error) {
	return regexp.MatchString(`^M{0,3}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$`, romanNumber)
}

func romanToDecimal(roman string) int64 {
	var out int64
	for i, c := range roman {
		if (i + 1) == utf8.RuneCountInString(roman) || RomanNumerals[string(c)] >= RomanNumerals[string(roman[i+1])] {
			out += RomanNumerals[string(c)]
		} else {
			out -= RomanNumerals[string(c)]
		}
	}
	return out
}

func transform(a []byte) (string, bool) {
	expr := strings.Join(strings.Split(string(a), " "), "")
	var out string
	var curRoman string
	var decimal int64
	flag := false
	index := 0
	for index < utf8.RuneCountInString(expr) {
		c := expr[index]
		if v := InOperator(c, Letters); !v {
			if !flag {
				out += string(c)
			} else {
				if v, _ = isValidNumber(curRoman); v || curRoman == "Z" {
					if curRoman != "Z" {
						decimal = romanToDecimal(curRoman)
					} else {
						decimal = 0
					}
					out += strconv.FormatInt(decimal, 10)
					out += string(c)
					curRoman = ""
					flag = false
				} else {
					return "error: Failed transformation", true
				}
			}
		} else {
			if flag {
				curRoman += string(c)
			} else {
				flag = true
				curRoman += string(c)
			}
			if index + 1 == utf8.RuneCountInString(expr) {
				if v, _ = isValidNumber(curRoman); v || curRoman == "Z" {
					if curRoman != "Z" {
						decimal = romanToDecimal(curRoman)
					} else {
						decimal = 0
					}
					out += strconv.FormatInt(decimal, 10)
				}
			}
		}
		index += 1
	}
	return out, false
}