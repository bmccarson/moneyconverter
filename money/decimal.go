package money

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Decimal can represent a floating-point number with a fixed precision.
// example: 1.52 = 152 * 10^(-2) will be stored as {152, 2}
type Decimal struct {
	//subunits is the amount of subunits. Multiply it by the precision to get the real value
	subunits int64
	// Number of "subunits" in a unit, expressed as a power of 10
	precision byte
}

const (
	// ErrInvalidDecimal is returned if the decimal is malformed.
	ErrInvalidDecimal = Error("unable to convert the decimal")

	// ErrTooLarge is returned if the quantity is too large - this would cause floating point percision errors
	ErrTooLarge = Error("quantity over 10^12 is too large")
)

// ParseDecimal converts a string into a Decimal representation.
// It assumes there is up to one decimal separator, and that the separator is '.'
func ParseDecimal(value string) (Decimal, error) {
	intPart, fracPart, _ := strings.Cut(value, ".")

	// maxDecimal is the number of digits in a thousand billion.
	const maxDecimal = 12

	if len(intPart) > maxDecimal {
		return Decimal{}, ErrTooLarge
	}

	subunits, err := strconv.ParseInt(intPart+fracPart, 10, 64)
	if err != nil {
		return Decimal{}, fmt.Errorf("%w: %s", ErrInvalidDecimal, err.Error())
	}

	precision := byte(len(fracPart))

	dec := Decimal{subunits: subunits, precision: precision}

	dec.simplify()

	return dec, nil
}

// simplfy checks if the last didget after a decimal is zero and removes it.
func (d *Decimal) simplify() {
	for d.subunits%10 == 0 && d.precision > 0 {
		d.precision--
		d.subunits /= 10
	}
}

// pow10 is a quick implementation to raise 10 to a given power.
// It's optimised for small powers, and slow for unusuall high powers.
func pow10(power int) int {
	switch power {
	case 0:
		return 1
	case 1:
		return 10
	case 2:
		return 100
	case 3:
		return 1000
	default:
		return int(math.Pow(10, float64(power)))
	}
}
