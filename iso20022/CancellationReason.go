package iso20022

// The status of an instruction, advice or request.
type CancellationReason1 struct {

	// Specifies the reason why the instruction is cancelled.
	Code *CancellationReason5Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancellationReason1) AddCode() *CancellationReason5Choice {
	c.Code = new(CancellationReason5Choice)
	return c.Code
}

func (c *CancellationReason1) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason why the instruction or request is cancelled.
type CancellationReason10 struct {

	// Specifies the reason why the instruction is cancelled.
	Code *CancellationReason21Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancellationReason10) AddCode() *CancellationReason21Choice {
	c.Code = new(CancellationReason21Choice)
	return c.Code
}

func (c *CancellationReason10) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason why the instruction or request is cancelled.
type CancellationReason11 struct {

	// Specifies the reason why the instruction is cancelled.
	Code *CancellationReason22Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancellationReason11) AddCode() *CancellationReason22Choice {
	c.Code = new(CancellationReason22Choice)
	return c.Code
}

func (c *CancellationReason11) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason why the instruction or request is cancelled.
type CancellationReason12 struct {

	// Specifies the reason why the instruction is cancelled.
	Code *CancellationReason23Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancellationReason12) AddCode() *CancellationReason23Choice {
	c.Code = new(CancellationReason23Choice)
	return c.Code
}

func (c *CancellationReason12) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason why the instruction or request is cancelled.
type CancellationReason13 struct {

	// Specifies the reason why the instruction is cancelled.
	Code *CancellationReason19Choice `xml:"Cd"`

	// Provides the corporate action event identification of the event that triggered the cancellation.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`
}

func (c *CancellationReason13) AddCode() *CancellationReason19Choice {
	c.Code = new(CancellationReason19Choice)
	return c.Code
}

func (c *CancellationReason13) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

// Specifies the reason why the instruction or request is cancelled.
type CancellationReason14 struct {

	// Specifies the reason why the instruction is cancelled.
	Code *CancellationReason24Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancellationReason14) AddCode() *CancellationReason24Choice {
	c.Code = new(CancellationReason24Choice)
	return c.Code
}

func (c *CancellationReason14) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Specifies the reason why the instruction or request is cancelled.
type CancellationReason15 struct {

	// Specifies the reason why the instruction is cancelled.
	Code *CancellationReason25Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancellationReason15) AddCode() *CancellationReason25Choice {
	c.Code = new(CancellationReason25Choice)
	return c.Code
}

func (c *CancellationReason15) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Specifies the reason why the instruction or request is cancelled.
type CancellationReason16 struct {

	// Specifies the reason why the instruction is cancelled.
	Code *CancellationReason26Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancellationReason16) AddCode() *CancellationReason26Choice {
	c.Code = new(CancellationReason26Choice)
	return c.Code
}

func (c *CancellationReason16) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Specifies the reason why the instruction or request is cancelled.
type CancellationReason18 struct {

	// Specifies the reason why the instruction is cancelled.
	Code *CancellationReason28Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancellationReason18) AddCode() *CancellationReason28Choice {
	c.Code = new(CancellationReason28Choice)
	return c.Code
}

func (c *CancellationReason18) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Specifies the reason why the instruction or request is cancelled.
type CancellationReason19 struct {

	// Specifies the reason why the instruction is cancelled.
	Code *CancellationReason24Choice `xml:"Cd"`

	// Provides the corporate action event identification of the event that triggered the cancellation.
	CorporateActionEventIdentification *RestrictedFINMax16Text `xml:"CorpActnEvtId,omitempty"`
}

func (c *CancellationReason19) AddCode() *CancellationReason24Choice {
	c.Code = new(CancellationReason24Choice)
	return c.Code
}

func (c *CancellationReason19) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*RestrictedFINMax16Text)(&value)
}

// The status of an instruction, advice or request.
type CancellationReason2 struct {

	// Specifies the reason why the instruction is cancelled.
	Code *CancellationReason3Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancellationReason2) AddCode() *CancellationReason3Choice {
	c.Code = new(CancellationReason3Choice)
	return c.Code
}

func (c *CancellationReason2) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max210Text)(&value)
}

// The status of an instruction, advice or request.
type CancellationReason5 struct {

	// Specifies the reason why the instruction is cancelled.
	Code *CancellationReason12Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancellationReason5) AddCode() *CancellationReason12Choice {
	c.Code = new(CancellationReason12Choice)
	return c.Code
}

func (c *CancellationReason5) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max210Text)(&value)
}

// The status of an instruction, advice or request.
type CancellationReason7 struct {

	// Specifies the reason why the instruction is cancelled.
	Code *CancellationReason17Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancellationReason7) AddCode() *CancellationReason17Choice {
	c.Code = new(CancellationReason17Choice)
	return c.Code
}

func (c *CancellationReason7) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason why the instruction or request is cancelled.
type CancellationReason9 struct {

	// Specifies the reason why the instruction is cancelled.
	Code *CancellationReason19Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancellationReason9) AddCode() *CancellationReason19Choice {
	c.Code = new(CancellationReason19Choice)
	return c.Code
}

func (c *CancellationReason9) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max210Text)(&value)
}
