package iso20022

// Provides the summation of the call amounts for the variation margin and optionaly the segregated independent amount.
type MarginCallResult2 struct {

	// Provides the summation of the call amounts for the variation margin amount only.
	VariationMarginResult *Result1 `xml:"VartnMrgnRslt"`

	// Provides the summation of the call amounts for the segregated independent amount.
	SegregatedIndependentAmount *Result1 `xml:"SgrtdIndpdntAmt,omitempty"`
}

func (m *MarginCallResult2) AddVariationMarginResult() *Result1 {
	m.VariationMarginResult = new(Result1)
	return m.VariationMarginResult
}

func (m *MarginCallResult2) AddSegregatedIndependentAmount() *Result1 {
	m.SegregatedIndependentAmount = new(Result1)
	return m.SegregatedIndependentAmount
}

// Provides the summation of the call amounts per margin type and optionaly the default fund amount (only for CCP).
type MarginCallResult3 struct {

	// Specifies the total amount required by the clearing member to participate to the default fund.
	DefaultFundAmount *ActiveCurrencyAndAmount `xml:"DfltFndAmt,omitempty"`

	// Provides the summation of the call amounts for the variation margin and the segregated independent amount or the segregated independent amount only or the total margin call amount only.
	MarginCallResult *MarginCallResult2Choice `xml:"MrgnCallRslt"`
}

func (m *MarginCallResult3) SetDefaultFundAmount(value, currency string) {
	m.DefaultFundAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (m *MarginCallResult3) AddMarginCallResult() *MarginCallResult2Choice {
	m.MarginCallResult = new(MarginCallResult2Choice)
	return m.MarginCallResult
}
