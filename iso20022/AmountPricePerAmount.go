package iso20022

// Specifies a ratio: Amount price per amount.
//
// Example:
// ISIN US629377AS17. Repurchase USD1087.17 cash for every USD1000 stock (NRG Energy Inc 8% Senior Notes 15/12/13).
type AmountPricePerAmount1 struct {

	// Type of amount price.
	AmountPriceType *AmountPriceType1FormatChoice `xml:"AmtPricTp"`

	// Value of the price.
	PriceValue *ActiveCurrencyAnd13DecimalAmount `xml:"PricVal"`

	// The amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`
}

func (a *AmountPricePerAmount1) AddAmountPriceType() *AmountPriceType1FormatChoice {
	a.AmountPriceType = new(AmountPriceType1FormatChoice)
	return a.AmountPriceType
}

func (a *AmountPricePerAmount1) SetPriceValue(value, currency string) {
	a.PriceValue = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (a *AmountPricePerAmount1) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

// Specifies a ratio: Amount price per amount.
// Example:
// ISIN US629377AS17. Repurchase USD1087.17 cash for every USD1000 stock (NRG Energy Inc 8% Senior Notes 15/12/13).
type AmountPricePerAmount2 struct {

	// Type of amount price.
	AmountPriceType *AmountPriceType1Code `xml:"AmtPricTp"`

	// Value of the price.
	PriceValue *ActiveCurrencyAnd13DecimalAmount `xml:"PricVal"`

	// The amount on which the price is based.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`
}

func (a *AmountPricePerAmount2) SetAmountPriceType(value string) {
	a.AmountPriceType = (*AmountPriceType1Code)(&value)
}

func (a *AmountPricePerAmount2) SetPriceValue(value, currency string) {
	a.PriceValue = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (a *AmountPricePerAmount2) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

// Specifies a ratio: Amount price per amount.
// Example:
// ISIN US629377AS17. Repurchase USD1087.17 cash for every USD1000 stock (NRG Energy Inc 8% Senior Notes 15/12/13).
type AmountPricePerAmount3 struct {

	// Type of amount price.
	AmountPriceType *AmountPriceType1Code `xml:"AmtPricTp"`

	// Value of the price.
	PriceValue *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"PricVal"`

	// The amount on which the price is based.
	Amount *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt"`
}

func (a *AmountPricePerAmount3) SetAmountPriceType(value string) {
	a.AmountPriceType = (*AmountPriceType1Code)(&value)
}

func (a *AmountPricePerAmount3) SetPriceValue(value, currency string) {
	a.PriceValue = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (a *AmountPricePerAmount3) SetAmount(value, currency string) {
	a.Amount = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}
