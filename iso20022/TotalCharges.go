package iso20022

// Total amount of charges.
type TotalCharges2 struct {

	// Total value of the charges for a specific order.
	TotalAmountOfCharges *ActiveCurrencyAnd13DecimalAmount `xml:"TtlAmtOfChrgs,omitempty"`

	// Information related to a specific charge.
	ChargeDetails []*Charge10 `xml:"ChrgDtls"`
}

func (t *TotalCharges2) SetTotalAmountOfCharges(value, currency string) {
	t.TotalAmountOfCharges = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (t *TotalCharges2) AddChargeDetails() *Charge10 {
	newValue := new(Charge10)
	t.ChargeDetails = append(t.ChargeDetails, newValue)
	return newValue
}

// Total amount of charges.
type TotalCharges3 struct {

	// Total value of the charges for a specific order.
	TotalAmountOfCharges *ActiveCurrencyAnd13DecimalAmount `xml:"TtlAmtOfChrgs,omitempty"`

	// Information related to a specific charge.
	ChargeDetails []*Charge18 `xml:"ChrgDtls"`
}

func (t *TotalCharges3) SetTotalAmountOfCharges(value, currency string) {
	t.TotalAmountOfCharges = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (t *TotalCharges3) AddChargeDetails() *Charge18 {
	newValue := new(Charge18)
	t.ChargeDetails = append(t.ChargeDetails, newValue)
	return newValue
}
