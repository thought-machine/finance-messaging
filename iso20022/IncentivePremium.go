package iso20022

// Cash premium made available if the securities holder consents or participates to an event.
type IncentivePremium2 struct {

	// Description of the premium.
	Description *Max350Text `xml:"Desc,omitempty"`

	// Cash premium paid per security, per vote or per attendee.
	Amount *PriceRateOrAmountChoice `xml:"Amt"`

	// Number of securities giving right to a premium.
	PerSecurity *Number `xml:"PerScty"`

	// Number of votes giving right to a premium.
	PerVote *Number `xml:"PerVote"`

	// Indicates that the premium is given per attendee.
	PerAttendee *YesNoIndicator `xml:"PerAttndee"`

	// Date/time for the payment of the premium.
	PaymentDate *DateFormat3Choice `xml:"PmtDt,omitempty"`
}

func (i *IncentivePremium2) SetDescription(value string) {
	i.Description = (*Max350Text)(&value)
}

func (i *IncentivePremium2) AddAmount() *PriceRateOrAmountChoice {
	i.Amount = new(PriceRateOrAmountChoice)
	return i.Amount
}

func (i *IncentivePremium2) SetPerSecurity(value string) {
	i.PerSecurity = (*Number)(&value)
}

func (i *IncentivePremium2) SetPerVote(value string) {
	i.PerVote = (*Number)(&value)
}

func (i *IncentivePremium2) SetPerAttendee(value string) {
	i.PerAttendee = (*YesNoIndicator)(&value)
}

func (i *IncentivePremium2) AddPaymentDate() *DateFormat3Choice {
	i.PaymentDate = new(DateFormat3Choice)
	return i.PaymentDate
}

// Cash premium made available if the securities holder consents or participates to an event.
type IncentivePremium3 struct {

	// Description of the premium.
	Description *Max350Text `xml:"Desc,omitempty"`

	// Cash premium paid per security, per vote or per attendee.
	Amount *PriceRateOrAmountChoice `xml:"Amt"`

	// Type of incentive premium.
	Type *IncentivePremiumType1Choice `xml:"Tp"`

	// Date/time for the payment of the premium.
	PaymentDate *DateFormat3Choice `xml:"PmtDt,omitempty"`
}

func (i *IncentivePremium3) SetDescription(value string) {
	i.Description = (*Max350Text)(&value)
}

func (i *IncentivePremium3) AddAmount() *PriceRateOrAmountChoice {
	i.Amount = new(PriceRateOrAmountChoice)
	return i.Amount
}

func (i *IncentivePremium3) AddType() *IncentivePremiumType1Choice {
	i.Type = new(IncentivePremiumType1Choice)
	return i.Type
}

func (i *IncentivePremium3) AddPaymentDate() *DateFormat3Choice {
	i.PaymentDate = new(DateFormat3Choice)
	return i.PaymentDate
}
