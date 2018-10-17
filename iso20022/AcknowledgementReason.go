package iso20022

// The status of an instruction, advice or request.
type AcknowledgementReason1 struct {

	// Reason provided for the status.
	Code *AcknowledgementReason1Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcknowledgementReason1) AddCode() *AcknowledgementReason1Choice {
	a.Code = new(AcknowledgementReason1Choice)
	return a.Code
}

func (a *AcknowledgementReason1) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies additional information about the processed instruction.
type AcknowledgementReason10 struct {

	// Specifies additional information about the processed instruction.
	Code *AcknowledgementReason13Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcknowledgementReason10) AddCode() *AcknowledgementReason13Choice {
	a.Code = new(AcknowledgementReason13Choice)
	return a.Code
}

func (a *AcknowledgementReason10) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies additional information about the processed instruction.
type AcknowledgementReason11 struct {

	// Choice of format for the acknowledgement reason.
	Code *AcknowledgementReason14Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcknowledgementReason11) AddCode() *AcknowledgementReason14Choice {
	a.Code = new(AcknowledgementReason14Choice)
	return a.Code
}

func (a *AcknowledgementReason11) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies additional information about the processed instruction.
type AcknowledgementReason12 struct {

	// Reason provided for the status.
	Code *AcknowledgementReason15Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcknowledgementReason12) AddCode() *AcknowledgementReason15Choice {
	a.Code = new(AcknowledgementReason15Choice)
	return a.Code
}

func (a *AcknowledgementReason12) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies additional information about the processed instruction.
type AcknowledgementReason13 struct {

	// Reason provided for the status.
	Code *AcknowledgementReason16Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcknowledgementReason13) AddCode() *AcknowledgementReason16Choice {
	a.Code = new(AcknowledgementReason16Choice)
	return a.Code
}

func (a *AcknowledgementReason13) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Specifies additional information about the processed instruction.
type AcknowledgementReason15 struct {

	// Specifies additional information about the processed instruction.
	Code *AcknowledgementReason18Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcknowledgementReason15) AddCode() *AcknowledgementReason18Choice {
	a.Code = new(AcknowledgementReason18Choice)
	return a.Code
}

func (a *AcknowledgementReason15) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Specifies additional information about the processed instruction.
type AcknowledgementReason18 struct {

	// Choice of format for the acknowledgement reason.
	Code *AcknowledgementReason21Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcknowledgementReason18) AddCode() *AcknowledgementReason21Choice {
	a.Code = new(AcknowledgementReason21Choice)
	return a.Code
}

func (a *AcknowledgementReason18) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Specifies additional information about the processed instruction.
type AcknowledgementReason19 struct {

	// Reason provided for the status.
	Code *AcknowledgementReason22Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcknowledgementReason19) AddCode() *AcknowledgementReason22Choice {
	a.Code = new(AcknowledgementReason22Choice)
	return a.Code
}

func (a *AcknowledgementReason19) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// The status of an instruction, advice or request.
type AcknowledgementReason2 struct {

	// Specifies additional information about the processed instruction.
	Code *AcknowledgementReason3Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcknowledgementReason2) AddCode() *AcknowledgementReason3Choice {
	a.Code = new(AcknowledgementReason3Choice)
	return a.Code
}

func (a *AcknowledgementReason2) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max210Text)(&value)
}

// The status of an instruction, advice or request.
type AcknowledgementReason3 struct {

	// Reason provided for the status.
	Code *AcknowledgementReason4Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcknowledgementReason3) AddCode() *AcknowledgementReason4Choice {
	a.Code = new(AcknowledgementReason4Choice)
	return a.Code
}

func (a *AcknowledgementReason3) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max210Text)(&value)
}

// The status of an instruction, advice or request.
type AcknowledgementReason7 struct {

	// Choice of format for the acknowledgement reason.
	Code *AcknowledgementReason9Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcknowledgementReason7) AddCode() *AcknowledgementReason9Choice {
	a.Code = new(AcknowledgementReason9Choice)
	return a.Code
}

func (a *AcknowledgementReason7) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies additional information about the processed instruction.
type AcknowledgementReason9 struct {

	// Reason provided for the status.
	Code *AcknowledgementReason12Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcknowledgementReason9) AddCode() *AcknowledgementReason12Choice {
	a.Code = new(AcknowledgementReason12Choice)
	return a.Code
}

func (a *AcknowledgementReason9) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max210Text)(&value)
}
