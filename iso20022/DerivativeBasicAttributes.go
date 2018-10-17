package iso20022

// Details of the derivative contract not included in the general financial instrument attributes component.
type DerivativeBasicAttributes1 struct {

	// Amount underlying a financial derivatives contract necessary for calculating payments or receipts.
	NotionalCurrencyAndAmount *ActiveOrHistoricCurrencyAndAmount `xml:"NtnlCcyAndAmt"`

	// Indicates whether the given derivative price includes a prorated accrued interest component.
	InterestIncludedInPrice *YesNoIndicator `xml:"IntrstInclInPric,omitempty"`
}

func (d *DerivativeBasicAttributes1) SetNotionalCurrencyAndAmount(value, currency string) {
	d.NotionalCurrencyAndAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (d *DerivativeBasicAttributes1) SetInterestIncludedInPrice(value string) {
	d.InterestIncludedInPrice = (*YesNoIndicator)(&value)
}

// Details of the derivative contract not included in the general financial instrument attributes component.
type DerivativeBasicAttributes2 struct {

	// Amount underlying a financial derivatives contract necessary for calculating payments or receipts.
	NotionalCurrencyAndAmount *RestrictedFINActiveOrHistoricCurrencyAndAmount `xml:"NtnlCcyAndAmt"`

	// Indicates whether the given derivative price includes a prorated accrued interest component.
	InterestIncludedInPrice *YesNoIndicator `xml:"IntrstInclInPric,omitempty"`
}

func (d *DerivativeBasicAttributes2) SetNotionalCurrencyAndAmount(value, currency string) {
	d.NotionalCurrencyAndAmount = NewRestrictedFINActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (d *DerivativeBasicAttributes2) SetInterestIncludedInPrice(value string) {
	d.InterestIncludedInPrice = (*YesNoIndicator)(&value)
}
