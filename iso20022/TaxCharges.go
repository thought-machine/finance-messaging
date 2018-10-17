package iso20022

// Amount of money due to the government or tax authority, according to various pre-defined parameters such as thresholds or income.
type TaxCharges1 struct {

	// Reference used to identify the nature of tax levied, such as Value Added Tax (VAT).
	Identification *Max35Text `xml:"Id,omitempty"`

	// Rate used to calculate the tax.
	Rate *PercentageRate `xml:"Rate,omitempty"`

	// Amount of money resulting from the calculation of the tax.
	Amount *CurrencyAndAmount `xml:"Amt,omitempty"`
}

func (t *TaxCharges1) SetIdentification(value string) {
	t.Identification = (*Max35Text)(&value)
}

func (t *TaxCharges1) SetRate(value string) {
	t.Rate = (*PercentageRate)(&value)
}

func (t *TaxCharges1) SetAmount(value, currency string) {
	t.Amount = NewCurrencyAndAmount(value, currency)
}

// Amount of money due to the government or tax authority, according to various pre-defined parameters such as thresholds or income.
type TaxCharges2 struct {

	// Unique reference to unambiguously identify the nature of the tax levied, such as Value Added Tax (VAT)
	Identification *Max35Text `xml:"Id,omitempty"`

	// Rate used to calculate the tax.
	Rate *PercentageRate `xml:"Rate,omitempty"`

	// Amount of money resulting from the calculation of the tax.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`
}

func (t *TaxCharges2) SetIdentification(value string) {
	t.Identification = (*Max35Text)(&value)
}

func (t *TaxCharges2) SetRate(value string) {
	t.Rate = (*PercentageRate)(&value)
}

func (t *TaxCharges2) SetAmount(value, currency string) {
	t.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}
