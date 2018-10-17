package iso20022

// Defines the elements used to calculate the collateral margin call for the variation margin and optionally the segregated independent amount.
type Margin1 struct {

	// Elements used to calculate the collateral margin call for the variation margin.
	VariationMargin *VariationMargin1 `xml:"VartnMrgn"`

	// Elements used to calculate the collateral margin call for the segregated independent amount.
	SegregatedIndependentAmountMargin *SegregatedIndependentAmountMargin1 `xml:"SgrtdIndpdntAmtMrgn,omitempty"`
}

func (m *Margin1) AddVariationMargin() *VariationMargin1 {
	m.VariationMargin = new(VariationMargin1)
	return m.VariationMargin
}

func (m *Margin1) AddSegregatedIndependentAmountMargin() *SegregatedIndependentAmountMargin1 {
	m.SegregatedIndependentAmountMargin = new(SegregatedIndependentAmountMargin1)
	return m.SegregatedIndependentAmountMargin
}

// Provides details on the calculation of the margin.
type Margin3 struct {

	// Margin required for absorbing future market price fluctuations (market risks) occurring between the default of a member and close-out of unsettled securities positions by the central counterparty.
	InitialMargin *Amount2 `xml:"InitlMrgn,omitempty"`

	// Provides details on the calculation of the variation margin.
	VariationMargin []*VariationMargin3 `xml:"VartnMrgn,omitempty"`

	// Provides details on the margin type and amount.
	OtherMargin []*Margin4 `xml:"OthrMrgn,omitempty"`
}

func (m *Margin3) AddInitialMargin() *Amount2 {
	m.InitialMargin = new(Amount2)
	return m.InitialMargin
}

func (m *Margin3) AddVariationMargin() *VariationMargin3 {
	newValue := new(VariationMargin3)
	m.VariationMargin = append(m.VariationMargin, newValue)
	return newValue
}

func (m *Margin3) AddOtherMargin() *Margin4 {
	newValue := new(Margin4)
	m.OtherMargin = append(m.OtherMargin, newValue)
	return newValue
}

// Provides details on the type of margin amounts.
type Margin4 struct {

	// Specifies the type of margin that is calculated.
	Type *MarginType1Choice `xml:"Tp"`

	// Provides the margin amount in the reporting currency and optionally in the original currency.
	Amount *Amount2 `xml:"Amt"`

	// Specifies whether the margin type position is short or long, that is, whether the balance is a negative or positive balance.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`
}

func (m *Margin4) AddType() *MarginType1Choice {
	m.Type = new(MarginType1Choice)
	return m.Type
}

func (m *Margin4) AddAmount() *Amount2 {
	m.Amount = new(Amount2)
	return m.Amount
}

func (m *Margin4) SetCreditDebitIndicator(value string) {
	m.CreditDebitIndicator = (*CreditDebitCode)(&value)
}
