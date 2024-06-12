package money

// Convert applies the change rate to convert an amount to a target currency
func Convert(amount Amount, to Currency) (Amount, error) {
	return Amount{}, nil
}

// applyExchangeRate returns a new Amount representing the input multiplied by the rate.
// The precision of the returned value is that of the target Currency.
// This function does not guarantee that the output amount is supported.
func applyExchangeRate(a Amount, target Currency, rate ExchangeRate) (Amount, error) {
	converted, err := multiply(a.quantity, rate)
	if err != nil {
		return Amount{}, err
	}

	switch {
	case converted.precision > target.precision:
		converted.subunits = converted.subunits / pow10(converted.precision-target.precision)
	case converted.precision < target.precision:
		converted.subunits = converted.subunits * pow10(target.precision-converted.precision)
	}

	converted.precision = target.precision

	return Amount{
		currency: target,
		quantity: converted,
	}, nil
}
