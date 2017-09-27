package Problem0166

import (
	"fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	minus := ""
	if numerator*denominator < 0 {
		minus = "-"
	}

	// 确保 n 和 d 是非负数
	n, d := abs(numerator), abs(denominator)

	// n / d 的小数部分
	ds := ""

	if n < d {
		rec := make(map[int]int)
		idx := 0
		digit := 0
		for {
			if i, ok := rec[n]; ok {
				return fmt.Sprintf("0.%s(%s)", ds[:i], ds[i:])
			}

			idx += digit
			rec[n] = idx

			digit = 0
			for n < d {
				n *= 10
				digit++
			}

			format := fmt.Sprintf("%%0%dd", digit)
			ds += fmt.Sprintf(format, n/d)
			n %= d

			if n == 0 {
				return "0." + ds
			}
		}
	}

	ds = fractionToDecimal(n%d, d)
	return minus + strconv.Itoa(n/d) + ds[1:]
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
