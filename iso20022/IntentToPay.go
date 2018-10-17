package iso20022

// Specifies the details of an intention to pay based on purchase orders or commercial invoice.
type IntentToPay1 struct {

	// The intention to pay is based on a purchase order.
	ByPurchaseOrder *ReportLine3 `xml:"ByPurchsOrdr"`

	// The intention to pay is based on a commercial invoice.
	ByCommercialInvoice *ReportLine4 `xml:"ByComrclInvc"`

	// Date at which the payment would be effected.
	ExpectedPaymentDate *ISODate `xml:"XpctdPmtDt"`

	// Specifies the beneficiary's account information.
	SettlementTerms *SettlementTerms2 `xml:"SttlmTerms,omitempty"`
}

func (i *IntentToPay1) AddByPurchaseOrder() *ReportLine3 {
	i.ByPurchaseOrder = new(ReportLine3)
	return i.ByPurchaseOrder
}

func (i *IntentToPay1) AddByCommercialInvoice() *ReportLine4 {
	i.ByCommercialInvoice = new(ReportLine4)
	return i.ByCommercialInvoice
}

func (i *IntentToPay1) SetExpectedPaymentDate(value string) {
	i.ExpectedPaymentDate = (*ISODate)(&value)
}

func (i *IntentToPay1) AddSettlementTerms() *SettlementTerms2 {
	i.SettlementTerms = new(SettlementTerms2)
	return i.SettlementTerms
}

// Specifies the details of an intention to pay based on purchase orders or commercial invoice.
type IntentToPay2 struct {

	// Specifies if breakdown is by purchase order or commercial invoice.
	Breakdown *BreakDown1Choice `xml:"Brkdwn"`

	// Date at which the payment would be effected.
	ExpectedPaymentDate *ISODate `xml:"XpctdPmtDt"`

	// Specifies the beneficiary's account information.
	SettlementTerms *SettlementTerms3 `xml:"SttlmTerms,omitempty"`
}

func (i *IntentToPay2) AddBreakdown() *BreakDown1Choice {
	i.Breakdown = new(BreakDown1Choice)
	return i.Breakdown
}

func (i *IntentToPay2) SetExpectedPaymentDate(value string) {
	i.ExpectedPaymentDate = (*ISODate)(&value)
}

func (i *IntentToPay2) AddSettlementTerms() *SettlementTerms3 {
	i.SettlementTerms = new(SettlementTerms3)
	return i.SettlementTerms
}
