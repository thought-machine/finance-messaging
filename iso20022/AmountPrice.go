package iso20022

// Price expressed as an amount.
type AmountPrice1 struct {

	// Type of amount price.
	AmountPriceType *AmountPriceType1FormatChoice `xml:"AmtPricTp"`

	// Value of the price.
	PriceValue *ActiveCurrencyAnd13DecimalAmount `xml:"PricVal"`
}

func (a *AmountPrice1) AddAmountPriceType() *AmountPriceType1FormatChoice {
	a.AmountPriceType = new(AmountPriceType1FormatChoice)
	return a.AmountPriceType
}

func (a *AmountPrice1) SetPriceValue(value, currency string) {
	a.PriceValue = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

// Price expressed as an actual amount.
type AmountPrice2 struct {

	// Type of amount price.
	AmountPriceType *AmountPriceType2Code `xml:"AmtPricTp"`

	// Value of the price.
	PriceValue *ActiveCurrencyAnd13DecimalAmount `xml:"PricVal"`
}

func (a *AmountPrice2) SetAmountPriceType(value string) {
	a.AmountPriceType = (*AmountPriceType2Code)(&value)
}

func (a *AmountPrice2) SetPriceValue(value, currency string) {
	a.PriceValue = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

// Price expressed as an amount.
type AmountPrice3 struct {

	// Type of amount price.
	AmountPriceType *AmountPriceType1Code `xml:"AmtPricTp"`

	// Value of the price.
	PriceValue *ActiveCurrencyAnd13DecimalAmount `xml:"PricVal"`
}

func (a *AmountPrice3) SetAmountPriceType(value string) {
	a.AmountPriceType = (*AmountPriceType1Code)(&value)
}

func (a *AmountPrice3) SetPriceValue(value, currency string) {
	a.PriceValue = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

// Price expressed as an actual amount.
type AmountPrice4 struct {

	// Type of amount price.
	AmountPriceType *AmountPriceType2Code `xml:"AmtPricTp"`

	// Value of the price.
	PriceValue *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"PricVal"`
}

func (a *AmountPrice4) SetAmountPriceType(value string) {
	a.AmountPriceType = (*AmountPriceType2Code)(&value)
}

func (a *AmountPrice4) SetPriceValue(value, currency string) {
	a.PriceValue = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

// Price expressed as an amount.
type AmountPrice5 struct {

	// Type of amount price.
	AmountPriceType *AmountPriceType1Code `xml:"AmtPricTp"`

	// Value of the price.
	PriceValue *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"PricVal"`
}

func (a *AmountPrice5) SetAmountPriceType(value string) {
	a.AmountPriceType = (*AmountPriceType1Code)(&value)
}

func (a *AmountPrice5) SetPriceValue(value, currency string) {
	a.PriceValue = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}
