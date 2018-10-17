package iso20022

// Specifies the reason why the corporate action reversal occurs.
type CorporateActionReversalReason1 struct {

	// Specifies the reason for the reversal.
	Reason *CorporateActionReversalReason1Choice `xml:"Rsn"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max256Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CorporateActionReversalReason1) AddReason() *CorporateActionReversalReason1Choice {
	c.Reason = new(CorporateActionReversalReason1Choice)
	return c.Reason
}

func (c *CorporateActionReversalReason1) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max256Text)(&value)
}

// Specifies the reason why the corporate action reversal occurs.
type CorporateActionReversalReason3 struct {

	// Specifies the reason for the reversal.
	Reason *CorporateActionReversalReason3Choice `xml:"Rsn"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max256Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CorporateActionReversalReason3) AddReason() *CorporateActionReversalReason3Choice {
	c.Reason = new(CorporateActionReversalReason3Choice)
	return c.Reason
}

func (c *CorporateActionReversalReason3) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max256Text)(&value)
}

// Specifies the reason why the corporate action reversal occurs.
type CorporateActionReversalReason4 struct {

	// Specifies the reason for the reversal.
	Reason *CorporateActionReversalReason4Choice `xml:"Rsn"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *RestrictedFINXMax256Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CorporateActionReversalReason4) AddReason() *CorporateActionReversalReason4Choice {
	c.Reason = new(CorporateActionReversalReason4Choice)
	return c.Reason
}

func (c *CorporateActionReversalReason4) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*RestrictedFINXMax256Text)(&value)
}
