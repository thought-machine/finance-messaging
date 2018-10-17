package iso20022

// Further information on the cancellation reason of the transaction.
type CancellationReasonInformation1 struct {

	// Party issuing the cancellation request.
	CancellationOriginator *PartyIdentification8 `xml:"CxlOrgtr,omitempty"`

	// Specifies the reason for the cancellation.
	CancellationReason *CancellationReason1Choice `xml:"CxlRsn,omitempty"`

	// Further details on the cancellation request reason.
	//
	// Usage: Additional cancellation reason information can be used to further detail the cancellation request reason.
	AdditionalCancellationReasonInformation []*Max105Text `xml:"AddtlCxlRsnInf,omitempty"`
}

func (c *CancellationReasonInformation1) AddCancellationOriginator() *PartyIdentification8 {
	c.CancellationOriginator = new(PartyIdentification8)
	return c.CancellationOriginator
}

func (c *CancellationReasonInformation1) AddCancellationReason() *CancellationReason1Choice {
	c.CancellationReason = new(CancellationReason1Choice)
	return c.CancellationReason
}

func (c *CancellationReasonInformation1) AddAdditionalCancellationReasonInformation(value string) {
	c.AdditionalCancellationReasonInformation = append(c.AdditionalCancellationReasonInformation, (*Max105Text)(&value))
}

// Set of elements used to provide information on the reason of the mandate cancellation request.
type CancellationReasonInformation2 struct {

	// Party that issues the cancellation request.
	Originator *PartyIdentification32 `xml:"Orgtr,omitempty"`

	// Specifies the reason for the cancellation request.
	Reason *MandateReason1Choice `xml:"Rsn"`

	// Further details on the cancellation request reason.
	AdditionalInformation []*Max105Text `xml:"AddtlInf,omitempty"`
}

func (c *CancellationReasonInformation2) AddOriginator() *PartyIdentification32 {
	c.Originator = new(PartyIdentification32)
	return c.Originator
}

func (c *CancellationReasonInformation2) AddReason() *MandateReason1Choice {
	c.Reason = new(MandateReason1Choice)
	return c.Reason
}

func (c *CancellationReasonInformation2) AddAdditionalInformation(value string) {
	c.AdditionalInformation = append(c.AdditionalInformation, (*Max105Text)(&value))
}

// Set of elements used to provide information on the reason of the cancellation request.
type CancellationReasonInformation3 struct {

	// Party that issues the cancellation request.
	Originator *PartyIdentification32 `xml:"Orgtr,omitempty"`

	// Specifies the reason for the cancellation.
	Reason *CancellationReason2Choice `xml:"Rsn,omitempty"`

	// Further details on the cancellation request reason.
	AdditionalInformation []*Max105Text `xml:"AddtlInf,omitempty"`
}

func (c *CancellationReasonInformation3) AddOriginator() *PartyIdentification32 {
	c.Originator = new(PartyIdentification32)
	return c.Originator
}

func (c *CancellationReasonInformation3) AddReason() *CancellationReason2Choice {
	c.Reason = new(CancellationReason2Choice)
	return c.Reason
}

func (c *CancellationReasonInformation3) AddAdditionalInformation(value string) {
	c.AdditionalInformation = append(c.AdditionalInformation, (*Max105Text)(&value))
}
