package solution

import (
	"strings"
	"unicode/utf8"
)

func InOperator(a byte, list []string) bool {
	for _, b := range list {
		if b == string(rune(a)) {
			return true
		}
	}
	return false
}

func handleZero(s *string, idx, brackets, status int) (int, int, int) {
	if ((*s)[idx] == '-' || (*s)[idx] == '+' || (*s)[idx] == '*' || (*s)[idx] == '/') && status == 0 && idx != 0 {
		status = 2
	} else if (*s)[idx] == '(' {
		brackets += 1
	} else if v := InOperator((*s)[idx], Letters); !v && (*s)[idx] != '-' {
		status = 3
	} else if v = InOperator((*s)[idx], Letters); v && (*s)[idx] != '-' {
		for idx < utf8.RuneCountInString(*s) - 1 && InOperator((*s)[idx + 1], Letters) {
			idx += 1
		}
		status = 1
	}
	return idx, brackets, status
}

func handleFirst(s *string, idx, brackets, status int) (int, int, int) {
	if (*s)[idx] == ')' {
		brackets -= 1
		if brackets < 0 {
			status = 3
		}
	} else {
		if v := InOperator((*s)[idx], Ops); v {
			status = 0
		} else {
			status = 3
		}
	}
	return idx, brackets, status
}

func handleSecond(s *string, idx, brackets int) bool {
	return (InOperator((*s)[idx - 1], Letters) || (*s)[idx - 1] == ')') && brackets == 0
}

func handler(line *string) bool {
	idx := 0
	bracketsCnt := 0
	state := 0
	flag := false
	for state != 3 {
		if state == 0 {
			if idx < utf8.RuneCountInString(*line) {
				idx, bracketsCnt, state = handleZero(line, idx, bracketsCnt, state)
				if state != 3 {
					idx += 1
				}
			}
		} else if state == 1 {
			if idx < utf8.RuneCountInString(*line) {
				idx, bracketsCnt, state = handleFirst(line, idx, bracketsCnt, state)
				if state != 3 {
					idx += 1
				}
			} else {
				state = 2
			}
		} else if state == 2 {
			flag = handleSecond(line, idx, bracketsCnt)
			state = 3
		}
	}
	return flag
}

func validate(stream *[]string) bool {
	line := strings.Join(strings.Split(strings.Join(*stream, ""), " "), "")
	return handler(&line)
}

