package iso20022

// Reason for a cancelled status.
type CancelledStatusReason1 struct {

	// Reason for a cancelled status in free format text.
	AdditionalInformation *Max350Text `xml:"AddtlInf"`
}

func (c *CancelledStatusReason1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}

// Specifies reasons for the cancelled status.
type CancelledStatusReason11 struct {

	// Specifies the reason why the instruction or instruction cancellation has been cancelled.
	ReasonCode *CancelledReason8Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancelledStatusReason11) AddReasonCode() *CancelledReason8Choice {
	c.ReasonCode = new(CancelledReason8Choice)
	return c.ReasonCode
}

func (c *CancelledStatusReason11) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies reasons for the cancelled status.
type CancelledStatusReason12 struct {

	// Specifies the reason why the instruction or instruction cancellation has been cancelled.
	ReasonCode *CancelledReason9Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancelledStatusReason12) AddReasonCode() *CancelledReason9Choice {
	c.ReasonCode = new(CancelledReason9Choice)
	return c.ReasonCode
}

func (c *CancelledStatusReason12) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies reasons for the cancelled status.
type CancelledStatusReason13 struct {

	// Specifies the reason why the instruction or instruction cancellation has been cancelled.
	ReasonCode *CancelledReason10Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancelledStatusReason13) AddReasonCode() *CancelledReason10Choice {
	c.ReasonCode = new(CancelledReason10Choice)
	return c.ReasonCode
}

func (c *CancelledStatusReason13) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Specifies reasons for the cancelled status.
type CancelledStatusReason14 struct {

	// Specifies the reason why the instruction or instruction cancellation has been cancelled.
	ReasonCode *CancelledReason11Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancelledStatusReason14) AddReasonCode() *CancelledReason11Choice {
	c.ReasonCode = new(CancelledReason11Choice)
	return c.ReasonCode
}

func (c *CancelledStatusReason14) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Reason for a cancelled status.
type CancelledStatusReason16 struct {

	// Reason for the cancelled status.
	Reason *CancelledReason12Choice `xml:"Rsn,omitempty"`

	// Additional information about the cancelled status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *CancelledStatusReason16) AddReason() *CancelledReason12Choice {
	c.Reason = new(CancelledReason12Choice)
	return c.Reason
}

func (c *CancelledStatusReason16) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}

// Specifies reasons for the cancelled status.
type CancelledStatusReason4 struct {

	// Specifies the reason why the instruction or instruction cancellation has been cancelled.
	ReasonCode *CancelledReason1Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancelledStatusReason4) AddReasonCode() *CancelledReason1Choice {
	c.ReasonCode = new(CancelledReason1Choice)
	return c.ReasonCode
}

func (c *CancelledStatusReason4) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies reasons for the cancelled status.
type CancelledStatusReason6 struct {

	// Specifies the reason why the instruction or instruction cancellation has been cancelled.
	ReasonCode *CancelledReason3Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancelledStatusReason6) AddReasonCode() *CancelledReason3Choice {
	c.ReasonCode = new(CancelledReason3Choice)
	return c.ReasonCode
}

func (c *CancelledStatusReason6) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies reasons for the cancelled status.
type CancelledStatusReason8 struct {

	// Specifies the reason why the instruction or instruction cancellation has been cancelled.
	ReasonCode *CancelledReason5Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancelledStatusReason8) AddReasonCode() *CancelledReason5Choice {
	c.ReasonCode = new(CancelledReason5Choice)
	return c.ReasonCode
}

func (c *CancelledStatusReason8) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max210Text)(&value)
}
