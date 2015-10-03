package byteformat
import (
	"fmt"
	"errors"
	"strings"
)

var precision uint8 = 2
var printfFormat string = "%.2f"

var TooPrecise = errors.New("Precision must not exceed 5")
var UnknownUnit = errors.New("Unknown unit")

const (
	BYTE = float64(1)
	KILOBYTE = 1024 * BYTE
	MEGABYTE = 1024 * KILOBYTE
	GIGABYTE = 1024 * MEGABYTE
	TERABYTE = 1024 * GIGABYTE
	PETABYTE = 1024 * TERABYTE
	EXABYTE = 1024 * PETABYTE
)

func SetPrecision(precision uint8) error {
	if precision > 5 {
		return TooPrecise
	}
	precision = precision
	printfFormat = fmt.Sprintf("%%.%df", precision)
	return nil
}

func HumanizeSize(size uint64) string {
	val := float64(size)
	var unit byte

	switch {
	case val >= EXABYTE:
		unit = 'E'
		val /= EXABYTE
	case val >= PETABYTE:
		unit = 'P'
		val /= PETABYTE
	case val >= TERABYTE:
		unit = 'T'
		val /= TERABYTE
	case val >= GIGABYTE:
		unit = 'G'
		val /= GIGABYTE
	case val >= MEGABYTE:
		unit = 'M'
		val /= MEGABYTE
	case val >= KILOBYTE:
		unit = 'K'
		val /= KILOBYTE
	case val >= BYTE:
		unit = 'b'
	case val == 0:
		return "0"
	}

	ret := fmt.Sprintf(printfFormat, val)
	ret = strings.TrimRight(ret, "0")
	ret = strings.TrimRight(ret, ".")
	return fmt.Sprintf("%s%c", ret, unit)
}

func HumanizeBytes(size uint64) string {
	var runes []rune
	for l := 0; size > 0; l++ {
		digit := size % 10
		size = (size - digit) / 10

		if (l > 2) {
			runes = append(runes, ',')
			l = 0
		}

		runes = append(runes, rune(digit + '0'))
	}

	for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
        runes[i], runes[j] = runes[j], runes[i]
    }

	return string(runes)
}

func FromString(input string) (uint64, error) {
	var base float64
	var unit byte
	var ret int

	if ret, err := fmt.Sscanf(input, "%f%c", &base, &unit); ret != 2 && ret != 1 {
		fmt.Println(err)
		return 0, err
	}

	if ret == 1 {
		unit = 'b'
	}

	switch unit {
	case 'b', 'B':
		base *= BYTE
	case 'k', 'K':
		base *= KILOBYTE
	case 'm', 'M':
		base *= MEGABYTE
	case 'g', 'G':
		base *= GIGABYTE
	case 't', 'T':
		base *= TERABYTE
	case 'p', 'P':
		base *= PETABYTE
	case 'e', 'E':
		base *= EXABYTE
	default:
		return 0, UnknownUnit
	}

	return uint64(base), nil
}
