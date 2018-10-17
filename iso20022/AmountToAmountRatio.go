package iso20022

// Ratio expressed as a quotient of amounts.
type AmountToAmountRatio1 struct {

	// Numerator of the quotient of amounts.
	Amount1 *ActiveCurrencyAndAmount `xml:"Amt1"`

	// Denominator of the quotient of amounts
	Amount2 *ActiveCurrencyAndAmount `xml:"Amt2"`
}

func (a *AmountToAmountRatio1) SetAmount1(value, currency string) {
	a.Amount1 = NewActiveCurrencyAndAmount(value, currency)
}

func (a *AmountToAmountRatio1) SetAmount2(value, currency string) {
	a.Amount2 = NewActiveCurrencyAndAmount(value, currency)
}

// Ratio expressed as a quotient of amounts.
type AmountToAmountRatio2 struct {

	// Numerator of the quotient of amounts.
	Amount1 *ActiveCurrencyAnd13DecimalAmount `xml:"Amt1"`

	// Denominator of the quotient of amounts
	Amount2 *ActiveCurrencyAnd13DecimalAmount `xml:"Amt2"`
}

func (a *AmountToAmountRatio2) SetAmount1(value, currency string) {
	a.Amount1 = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (a *AmountToAmountRatio2) SetAmount2(value, currency string) {
	a.Amount2 = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

// Ratio expressed as a quotient of amounts.
type AmountToAmountRatio3 struct {

	// Numerator of the quotient of amounts.
	Amount1 *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt1"`

	// Denominator of the quotient of amounts
	Amount2 *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt2"`
}

func (a *AmountToAmountRatio3) SetAmount1(value, currency string) {
	a.Amount1 = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (a *AmountToAmountRatio3) SetAmount2(value, currency string) {
	a.Amount2 = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}
