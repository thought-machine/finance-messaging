package iso20022

// Specifies an amount and a rate status.
type AmountAndRateStatus1 struct {

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus1Code `xml:"RateSts"`
}

func (a *AmountAndRateStatus1) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (a *AmountAndRateStatus1) SetRateStatus(value string) {
	a.RateStatus = (*RateStatus1Code)(&value)
}

// Specifies an amount and a rate status.
type AmountAndRateStatus2 struct {

	// Value expressed as an amount.
	Amount *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus1Code `xml:"RateSts"`
}

func (a *AmountAndRateStatus2) SetAmount(value, currency string) {
	a.Amount = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (a *AmountAndRateStatus2) SetRateStatus(value string) {
	a.RateStatus = (*RateStatus1Code)(&value)
}
