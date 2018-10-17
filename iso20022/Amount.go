package iso20022

// Margin amount payable by one party to the other party.
type Amount1 struct {

	// Undisputed amount of the margin call request.
	AgreedAmount *ActiveCurrencyAndAmount `xml:"AgrdAmt"`

	// Unique identifier for the margin call request.
	MarginCallRequestIdentification *Max35Text `xml:"MrgnCallReqId"`

	// Provides additional information related to the margin call amount that has been agreed.
	AdditionalInformation *Max210Text `xml:"AddtlInf,omitempty"`
}

func (a *Amount1) SetAgreedAmount(value, currency string) {
	a.AgreedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (a *Amount1) SetMarginCallRequestIdentification(value string) {
	a.MarginCallRequestIdentification = (*Max35Text)(&value)
}

func (a *Amount1) SetAdditionalInformation(value string) {
	a.AdditionalInformation = (*Max210Text)(&value)
}

// Provides the amount in the reporting currency and optionally in the original currency.
type Amount2 struct {

	// Amount expressed in the original currency.
	OriginalCurrencyAmount *ActiveCurrencyAndAmount `xml:"OrgnlCcyAmt,omitempty"`

	// Amount expressed in the reporting currency.
	ReportingAmount *ImpliedCurrencyAndAmount `xml:"RptgAmt"`
}

func (a *Amount2) SetOriginalCurrencyAmount(value, currency string) {
	a.OriginalCurrencyAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (a *Amount2) SetReportingAmount(value, currency string) {
	a.ReportingAmount = NewImpliedCurrencyAndAmount(value, currency)
}
