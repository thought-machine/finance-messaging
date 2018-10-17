package iso20022

// Specifies a ratio: amount price per financial instrument quantity.
type AmountPricePerFinancialInstrumentQuantity1 struct {

	// Type of amount price.
	AmountPriceType *AmountPriceType1FormatChoice `xml:"AmtPricTp"`

	// Value of the price.
	PriceValue *ActiveCurrencyAnd13DecimalAmount `xml:"PricVal"`

	// Quantity of financial instrument.
	FinancialInstrumentQuantity *UnitOrFaceAmount1Choice `xml:"FinInstrmQty"`
}

func (a *AmountPricePerFinancialInstrumentQuantity1) AddAmountPriceType() *AmountPriceType1FormatChoice {
	a.AmountPriceType = new(AmountPriceType1FormatChoice)
	return a.AmountPriceType
}

func (a *AmountPricePerFinancialInstrumentQuantity1) SetPriceValue(value, currency string) {
	a.PriceValue = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (a *AmountPricePerFinancialInstrumentQuantity1) AddFinancialInstrumentQuantity() *UnitOrFaceAmount1Choice {
	a.FinancialInstrumentQuantity = new(UnitOrFaceAmount1Choice)
	return a.FinancialInstrumentQuantity
}

// Specifies a ratio: amount price per financial instrument quantity.
type AmountPricePerFinancialInstrumentQuantity3 struct {

	// Type of amount price.
	AmountPriceType *AmountPriceType1Code `xml:"AmtPricTp"`

	// Value of the price.
	PriceValue *ActiveCurrencyAnd13DecimalAmount `xml:"PricVal"`

	// Quantity of financial instrument.
	FinancialInstrumentQuantity *FinancialInstrumentQuantity1Choice `xml:"FinInstrmQty"`
}

func (a *AmountPricePerFinancialInstrumentQuantity3) SetAmountPriceType(value string) {
	a.AmountPriceType = (*AmountPriceType1Code)(&value)
}

func (a *AmountPricePerFinancialInstrumentQuantity3) SetPriceValue(value, currency string) {
	a.PriceValue = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (a *AmountPricePerFinancialInstrumentQuantity3) AddFinancialInstrumentQuantity() *FinancialInstrumentQuantity1Choice {
	a.FinancialInstrumentQuantity = new(FinancialInstrumentQuantity1Choice)
	return a.FinancialInstrumentQuantity
}

// Specifies a ratio: amount price per financial instrument quantity.
type AmountPricePerFinancialInstrumentQuantity6 struct {

	// Type of amount price.
	AmountPriceType *AmountPriceType1Code `xml:"AmtPricTp"`

	// Value of the price.
	PriceValue *ActiveCurrencyAnd13DecimalAmount `xml:"PricVal"`

	// Quantity of financial instrument.
	FinancialInstrumentQuantity *FinancialInstrumentQuantity1Choice `xml:"FinInstrmQty"`
}

func (a *AmountPricePerFinancialInstrumentQuantity6) SetAmountPriceType(value string) {
	a.AmountPriceType = (*AmountPriceType1Code)(&value)
}

func (a *AmountPricePerFinancialInstrumentQuantity6) SetPriceValue(value, currency string) {
	a.PriceValue = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (a *AmountPricePerFinancialInstrumentQuantity6) AddFinancialInstrumentQuantity() *FinancialInstrumentQuantity1Choice {
	a.FinancialInstrumentQuantity = new(FinancialInstrumentQuantity1Choice)
	return a.FinancialInstrumentQuantity
}

// Specifies a ratio: amount price per financial instrument quantity.
type AmountPricePerFinancialInstrumentQuantity7 struct {

	// Type of amount price.
	AmountPriceType *AmountPriceType1Code `xml:"AmtPricTp"`

	// Value of the price.
	PriceValue *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"PricVal"`

	// Quantity of financial instrument.
	FinancialInstrumentQuantity *FinancialInstrumentQuantity15Choice `xml:"FinInstrmQty"`
}

func (a *AmountPricePerFinancialInstrumentQuantity7) SetAmountPriceType(value string) {
	a.AmountPriceType = (*AmountPriceType1Code)(&value)
}

func (a *AmountPricePerFinancialInstrumentQuantity7) SetPriceValue(value, currency string) {
	a.PriceValue = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (a *AmountPricePerFinancialInstrumentQuantity7) AddFinancialInstrumentQuantity() *FinancialInstrumentQuantity15Choice {
	a.FinancialInstrumentQuantity = new(FinancialInstrumentQuantity15Choice)
	return a.FinancialInstrumentQuantity
}
