package iso20022

// Ratio expressed as amount per quantity.
type AmountAndQuantityRatio1 struct {

	// Cash amount.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Quantity expressed as number.
	Quantity *DecimalNumber `xml:"Qty"`
}

func (a *AmountAndQuantityRatio1) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndQuantityRatio1) SetQuantity(value string) {
	a.Quantity = (*DecimalNumber)(&value)
}

// Ratio expressed as amount per quantity.
type AmountAndQuantityRatio2 struct {

	// Cash amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Quantity expressed as number.
	Quantity *DecimalNumber `xml:"Qty"`
}

func (a *AmountAndQuantityRatio2) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (a *AmountAndQuantityRatio2) SetQuantity(value string) {
	a.Quantity = (*DecimalNumber)(&value)
}

// Ratio expressed as amount per quantity.
type AmountAndQuantityRatio4 struct {

	// Cash amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Quantity expressed as number.
	Quantity *DecimalNumber `xml:"Qty"`
}

func (a *AmountAndQuantityRatio4) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (a *AmountAndQuantityRatio4) SetQuantity(value string) {
	a.Quantity = (*DecimalNumber)(&value)
}

// Ratio expressed as amount per quantity.
type AmountAndQuantityRatio5 struct {

	// Cash amount.
	Amount *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Quantity expressed as number.
	Quantity *RestrictedFINDecimalNumber `xml:"Qty"`
}

func (a *AmountAndQuantityRatio5) SetAmount(value, currency string) {
	a.Amount = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (a *AmountAndQuantityRatio5) SetQuantity(value string) {
	a.Quantity = (*RestrictedFINDecimalNumber)(&value)
}
