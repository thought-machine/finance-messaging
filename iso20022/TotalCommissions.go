package iso20022

// Total amount of commissions related to a specific order.
type TotalCommissions2 struct {

	// Total value of the commissions for a specific order.
	TotalAmountOfCommissions *ActiveCurrencyAnd13DecimalAmount `xml:"TtlAmtOfComssns,omitempty"`

	// Information related to a specific commission.
	CommissionDetails []*Commission6 `xml:"ComssnDtls"`
}

func (t *TotalCommissions2) SetTotalAmountOfCommissions(value, currency string) {
	t.TotalAmountOfCommissions = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (t *TotalCommissions2) AddCommissionDetails() *Commission6 {
	newValue := new(Commission6)
	t.CommissionDetails = append(t.CommissionDetails, newValue)
	return newValue
}

// Total amount of commissions related to a specific order.
type TotalCommissions3 struct {

	// Total value of the commissions for a specific order.
	TotalAmountOfCommissions *ActiveCurrencyAnd13DecimalAmount `xml:"TtlAmtOfComssns,omitempty"`

	// Information related to a specific commission.
	CommissionDetails []*Commission10 `xml:"ComssnDtls"`
}

func (t *TotalCommissions3) SetTotalAmountOfCommissions(value, currency string) {
	t.TotalAmountOfCommissions = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (t *TotalCommissions3) AddCommissionDetails() *Commission10 {
	newValue := new(Commission10)
	t.CommissionDetails = append(t.CommissionDetails, newValue)
	return newValue
}
