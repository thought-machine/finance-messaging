package iso20022

// Indicates how the amount of the investment plan is split amongst the funds.
type Repartition1 struct {

	// Percentage of amount invested in a funds.
	Percentage *PercentageRate `xml:"Pctg"`

	// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
	FinancialInstrument *FinancialInstrument10 `xml:"FinInstrm"`
}

func (r *Repartition1) SetPercentage(value string) {
	r.Percentage = (*PercentageRate)(&value)
}

func (r *Repartition1) AddFinancialInstrument() *FinancialInstrument10 {
	r.FinancialInstrument = new(FinancialInstrument10)
	return r.FinancialInstrument
}

// Indicates how the amount of the investment plan is split amongst the funds.
type Repartition2 struct {

	// Percentage of amount invested in a funds.
	Percentage *PercentageRate `xml:"Pctg"`

	// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
	FinancialInstrument *FinancialInstrument29 `xml:"FinInstrm"`

	// When a fund has multiple currencies within same ISIN, this indicates the currency of the savings or withdrawal plan.
	CurrencyOfPlan *CurrencyCode `xml:"CcyOfPlan,omitempty"`
}

func (r *Repartition2) SetPercentage(value string) {
	r.Percentage = (*PercentageRate)(&value)
}

func (r *Repartition2) AddFinancialInstrument() *FinancialInstrument29 {
	r.FinancialInstrument = new(FinancialInstrument29)
	return r.FinancialInstrument
}

func (r *Repartition2) SetCurrencyOfPlan(value string) {
	r.CurrencyOfPlan = (*CurrencyCode)(&value)
}

// Indicates how the amount of the investment plan is split amongst the funds.
type Repartition3 struct {

	// Amount, units or percentage of financial instrument invested or withdrawn.
	Quantity *UnitsOrAmountOrPercentage1Choice `xml:"Qty"`

	// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
	FinancialInstrument *FinancialInstrument29 `xml:"FinInstrm"`

	// When a fund has multiple currencies within same ISIN, this indicates the currency of the savings or withdrawal plan.
	CurrencyOfPlan *CurrencyCode `xml:"CcyOfPlan,omitempty"`
}

func (r *Repartition3) AddQuantity() *UnitsOrAmountOrPercentage1Choice {
	r.Quantity = new(UnitsOrAmountOrPercentage1Choice)
	return r.Quantity
}

func (r *Repartition3) AddFinancialInstrument() *FinancialInstrument29 {
	r.FinancialInstrument = new(FinancialInstrument29)
	return r.FinancialInstrument
}

func (r *Repartition3) SetCurrencyOfPlan(value string) {
	r.CurrencyOfPlan = (*CurrencyCode)(&value)
}

// Indicates how the amount of the investment plan is split amongst the funds.
type Repartition4 struct {

	// Amount, units or percentage of financial instrument invested or withdrawn.
	Quantity *UnitsOrAmountOrPercentage1Choice `xml:"Qty"`

	// Detailed information about the security or investment fund.
	FinancialInstrument *FinancialInstrument51 `xml:"FinInstrm"`

	// When a fund has multiple currencies within same ISIN, this indicates the currency of the savings or withdrawal plan.
	CurrencyOfPlan *ActiveOrHistoricCurrencyCode `xml:"CcyOfPlan,omitempty"`
}

func (r *Repartition4) AddQuantity() *UnitsOrAmountOrPercentage1Choice {
	r.Quantity = new(UnitsOrAmountOrPercentage1Choice)
	return r.Quantity
}

func (r *Repartition4) AddFinancialInstrument() *FinancialInstrument51 {
	r.FinancialInstrument = new(FinancialInstrument51)
	return r.FinancialInstrument
}

func (r *Repartition4) SetCurrencyOfPlan(value string) {
	r.CurrencyOfPlan = (*ActiveOrHistoricCurrencyCode)(&value)
}

// Indicates how the amount of the investment plan is split amongst the funds.
type Repartition5 struct {

	// Amount, units or percentage of financial instrument invested or withdrawn.
	Quantity *UnitsOrAmountOrPercentage1Choice `xml:"Qty"`

	// Detailed information about the security or investment fund.
	FinancialInstrument *FinancialInstrument56 `xml:"FinInstrm"`

	// When a fund has multiple currencies within same ISIN, this indicates the currency of the savings or withdrawal plan.
	CurrencyOfPlan *ActiveOrHistoricCurrencyCode `xml:"CcyOfPlan,omitempty"`
}

func (r *Repartition5) AddQuantity() *UnitsOrAmountOrPercentage1Choice {
	r.Quantity = new(UnitsOrAmountOrPercentage1Choice)
	return r.Quantity
}

func (r *Repartition5) AddFinancialInstrument() *FinancialInstrument56 {
	r.FinancialInstrument = new(FinancialInstrument56)
	return r.FinancialInstrument
}

func (r *Repartition5) SetCurrencyOfPlan(value string) {
	r.CurrencyOfPlan = (*ActiveOrHistoricCurrencyCode)(&value)
}
