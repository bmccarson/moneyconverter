package money

import (
	"testing"
)

func TestParseCurrency_Success(t *testing.T) {
	tt := map[string]struct {
		in       string
		expected Currency
	}{
		"majority EUR":   {in: "EUR", expected: Currency{code: "EUR", precision: 2}},
		"thousandth BHD": {in: "BHD", expected: Currency{code: "BHD", precision: 3}},
		"tenth VND":      {in: "VND", expected: Currency{code: "VND", precision: 1}},
		"integer IRR":    {in: "IRR", expected: Currency{code: "IRR", precision: 0}},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := ParseCurrency(tc.in)
			if err != nil {
				t.Errorf("expected no error, got %s", err.Error())
			}

			if got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}