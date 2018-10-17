package iso20022

// Information about a switch leg that is rejected or repaired.
type SwitchLegReferences1 struct {

	// Unique technical identifier for an instance of a leg within a switch.
	RedemptionLegIdentification *Max35Text `xml:"RedLegId"`

	// Unique technical identifier for an instance of a leg within a switch.
	SubscriptionLegIdentification *Max35Text `xml:"SbcptLegId"`

	// Additional information about the reason for the rejection of a leg.
	LegRejectionReason *Max350Text `xml:"LegRjctnRsn,omitempty"`

	// Elements from the original switch order that have been repaired so that the switch order can be accepted.
	RepairedConditions *RepairedConditions3 `xml:"RprdConds,omitempty"`

	// Account identification of the switch leg that is rejected or repaired.
	InvestmentAccountDetails *InvestmentAccount13 `xml:"InvstmtAcctDtls,omitempty"`

	// Financial instrument identification of the switch leg that is rejected or repaired.
	FinancialInstrumentDetails *FinancialInstrument10 `xml:"FinInstrmDtls,omitempty"`
}

func (s *SwitchLegReferences1) SetRedemptionLegIdentification(value string) {
	s.RedemptionLegIdentification = (*Max35Text)(&value)
}

func (s *SwitchLegReferences1) SetSubscriptionLegIdentification(value string) {
	s.SubscriptionLegIdentification = (*Max35Text)(&value)
}

func (s *SwitchLegReferences1) SetLegRejectionReason(value string) {
	s.LegRejectionReason = (*Max350Text)(&value)
}

func (s *SwitchLegReferences1) AddRepairedConditions() *RepairedConditions3 {
	s.RepairedConditions = new(RepairedConditions3)
	return s.RepairedConditions
}

func (s *SwitchLegReferences1) AddInvestmentAccountDetails() *InvestmentAccount13 {
	s.InvestmentAccountDetails = new(InvestmentAccount13)
	return s.InvestmentAccountDetails
}

func (s *SwitchLegReferences1) AddFinancialInstrumentDetails() *FinancialInstrument10 {
	s.FinancialInstrumentDetails = new(FinancialInstrument10)
	return s.FinancialInstrumentDetails
}

// Information about a switch leg that is rejected or repaired.
type SwitchLegReferences2 struct {

	// Identification of a switch leg.
	LegIdentification *LegIdentification1Choice `xml:"LegId"`

	// Additional information about the reason for the rejection of the leg.
	LegRejectionReason *Max350Text `xml:"LegRjctnRsn,omitempty"`

	// Elements from the original individual order that have been repaired so that the order can be accepted.
	RepairedFee []*Fee3 `xml:"RprdFee,omitempty"`

	// Account identification of the switch leg that is rejected or repaired.
	InvestmentAccountDetails *InvestmentAccount58 `xml:"InvstmtAcctDtls,omitempty"`

	// Financial instrument identification of the switch leg that is rejected or repaired.
	FinancialInstrumentDetails *FinancialInstrument57 `xml:"FinInstrmDtls,omitempty"`
}

func (s *SwitchLegReferences2) AddLegIdentification() *LegIdentification1Choice {
	s.LegIdentification = new(LegIdentification1Choice)
	return s.LegIdentification
}

func (s *SwitchLegReferences2) SetLegRejectionReason(value string) {
	s.LegRejectionReason = (*Max350Text)(&value)
}

func (s *SwitchLegReferences2) AddRepairedFee() *Fee3 {
	newValue := new(Fee3)
	s.RepairedFee = append(s.RepairedFee, newValue)
	return newValue
}

func (s *SwitchLegReferences2) AddInvestmentAccountDetails() *InvestmentAccount58 {
	s.InvestmentAccountDetails = new(InvestmentAccount58)
	return s.InvestmentAccountDetails
}

func (s *SwitchLegReferences2) AddFinancialInstrumentDetails() *FinancialInstrument57 {
	s.FinancialInstrumentDetails = new(FinancialInstrument57)
	return s.FinancialInstrumentDetails
}
