package solution

import (
	"errors"
	"math"
	"sort"
	"strings"
	"unicode"
)

func sortedKeys(m map[int]string) []int {
	keys := make([]int, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	return keys
}

func decimalToRoman(number int64) string {
	if math.Abs(float64(number)) > 3999 {
		return "error: Impossible to get a roman number (absolute value > 3999)"
	}
	var out = make([]string, 0)
	abs := int64(math.Abs(float64(number)))
	keys := sortedKeys(RomanLiterals)
	for _, value := range keys {
		n := abs / int64(value)
		var i int64
		for i = 0; i < n; i++ {
			out = append(out, RomanLiterals[value])
		}
		abs -= n * int64(value)
	}
	if number < 0 {
		out = append([]string{"-"}, out...)
	}
	var res string
	if len(out) != 0 {
		res = strings.Join(out, "")
	} else {
		res = "Z"
	}
	return res
}

type Calculator struct {
	arr []byte
}

func sumSeq(seq []int64) (int64, error) {
	var res int64
	for i := range seq {
		value, err := Add64(res, seq[i])
		if err == nil {
			res = value
		} else {
			return 0, err
		}
	}
	return res, nil
}

func (calc *Calculator) Init(expr string) {
	calc.arr = []byte(expr)
}

func (calc *Calculator) helper(stream *[]byte) (int64, error) {
	if len(*stream) == 0 {
		return 0, errors.New("error: Line without expression")
	}
	var stack = make([]int64, 0)
	sign := '+'
	var num int64 = 0
	var err error
	for len(*stream) > 0 {
		c := rune((*stream)[0])
		*stream = (*stream)[1:]
		if unicode.IsDigit(c) {
			num = num * 10 + int64(c - 48)
		}
		if c == '(' {
			num, err = calc.helper(stream)
			if err != nil {
				return 0, err
			}
		}
		if len(*stream) == 0 || (c == '+' || c == '-' || c == '*' || c == '/' || c == ')') {
			if sign == '+' {
				stack = append(stack, num)
			} else if sign == '-' {
				stack = append(stack, -num)
			} else if sign == '*' {
				product, e := Mult64(stack[len(stack) - 1], num)
				if e == nil {
					stack[len(stack) - 1] = product
				} else{
					return 0, e
				}
				// stack[len(stack) - 1] = stack[len(stack) - 1] * num
			} else if sign == '/' {
				div, e := Div64(stack[len(stack) - 1], num)
				if e == nil {
					stack[len(stack) - 1] = div
				} else {
					return 0, e
				}
			}
			sign = c
			num = 0
			//sign, num = c, 0
			if sign == ')' {
				break
			}
		}
	}
	return sumSeq(stack)
}

func (calc *Calculator) Calculate() (string, error) {
	arr := make([]string, 0)
	for i := range calc.arr {
		arr = append(arr, string(calc.arr[i]))
	}
	if !validate(&arr) {
		return "", errors.New("error: Not valid expression")
	}
	transformStream, err := transform(calc.arr)
	if err {
		return "", errors.New("error: Not valid expression")
	}
	for _, c := range transformStream {
		arr = append(arr, string(c))
	}
	calc.arr = []byte(transformStream)
	value, e := calc.helper(&calc.arr)
	if e == nil {
		return decimalToRoman(value), nil
	} else {
		return decimalToRoman(0), e
	}
}