package iso20022

// Information on the type of tax.
type TaxType struct {

	// Description of the tax that is being paid, including specific representation required by taxing authority.
	CategoryDescription *Max35Text `xml:"CtgyDesc,omitempty"`

	// Rate used to calculate the tax.
	Rate *PercentageRate `xml:"Rate,omitempty"`

	// Amount of money on which the tax is based.
	TaxableBaseAmount *CurrencyAndAmount `xml:"TaxblBaseAmt,omitempty"`

	// Amount of money resulting from the calculation of the tax.
	Amount *CurrencyAndAmount `xml:"Amt,omitempty"`
}

func (t *TaxType) SetCategoryDescription(value string) {
	t.CategoryDescription = (*Max35Text)(&value)
}

func (t *TaxType) SetRate(value string) {
	t.Rate = (*PercentageRate)(&value)
}

func (t *TaxType) SetTaxableBaseAmount(value, currency string) {
	t.TaxableBaseAmount = NewCurrencyAndAmount(value, currency)
}

func (t *TaxType) SetAmount(value, currency string) {
	t.Amount = NewCurrencyAndAmount(value, currency)
}

// Specification of the tax type.
type TaxType1 struct {

	// Structured format.
	Structured *TaxType7Code `xml:"Strd"`

	// Additional information about the type of tax.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (t *TaxType1) SetStructured(value string) {
	t.Structured = (*TaxType7Code)(&value)
}

func (t *TaxType1) SetAdditionalInformation(value string) {
	t.AdditionalInformation = (*Max350Text)(&value)
}

// Specification of the tax type.
type TaxType2 struct {

	// Structured format.
	Structured *TaxType5Code `xml:"Strd"`

	// Additional information about the type of tax.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (t *TaxType2) SetStructured(value string) {
	t.Structured = (*TaxType5Code)(&value)
}

func (t *TaxType2) SetAdditionalInformation(value string) {
	t.AdditionalInformation = (*Max350Text)(&value)
}

// Specification of the tax type.
type TaxType3 struct {

	// Structured format.
	Structured *TaxType6Code `xml:"Strd"`

	// Additional information about the type of tax.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (t *TaxType3) SetStructured(value string) {
	t.Structured = (*TaxType6Code)(&value)
}

func (t *TaxType3) SetAdditionalInformation(value string) {
	t.AdditionalInformation = (*Max350Text)(&value)
}
