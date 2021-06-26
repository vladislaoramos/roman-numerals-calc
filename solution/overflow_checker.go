package solution

import (
	"errors"
	"math"
)

var ErrOverflow = errors.New("error: int64 overflow")

func Add64(left, right int64) (int64, error) {
	if right > 0 {
		if left > math.MaxInt64 - right {
			return 0, ErrOverflow
		}
	} else {
		if left < math.MinInt64 - right {
			return 0, ErrOverflow
		}
	}
	return left + right, nil
}

func Mult64(left, right int64) (int64, error) {
	result := left * right
	if left == 0 || right == 0 || left == 1 || right == 1 {
		return result, nil
	}
	if left == math.MaxInt64 || right == math.MaxInt64 {
		return 0, ErrOverflow
	}
	if result / right != left {
		return result, ErrOverflow
	}
	return result, nil
}

func Div64(left, right int64) (int64, error) {
	if right == 0 {
		return 0, errors.New("error: Cannot divide by zero")
	}
	return int64(math.Floor(float64(left) / float64(right))), nil
}
