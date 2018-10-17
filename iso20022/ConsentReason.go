package iso20022

// Specifies the underlying reason for the status of an object.
type ConsentReason2 struct {

	// Reason provided for the status.
	Code *ConsentOrRejectionReason2Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *ConsentReason2) AddCode() *ConsentOrRejectionReason2Choice {
	c.Code = new(ConsentOrRejectionReason2Choice)
	return c.Code
}

func (c *ConsentReason2) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the underlying reason for the status of an object.
type ConsentReason4 struct {

	// Reason provided for the status.
	Code *ConsentOrRejectionReason4Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *ConsentReason4) AddCode() *ConsentOrRejectionReason4Choice {
	c.Code = new(ConsentOrRejectionReason4Choice)
	return c.Code
}

func (c *ConsentReason4) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the underlying reason for the status of an object.
type ConsentReason5 struct {

	// Reason provided for the status.
	Code *ConsentOrRejectionReason5Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *ConsentReason5) AddCode() *ConsentOrRejectionReason5Choice {
	c.Code = new(ConsentOrRejectionReason5Choice)
	return c.Code
}

func (c *ConsentReason5) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
