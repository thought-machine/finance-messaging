package iso20022

// The status of an instruction, advice or request.
type FailingReason1 struct {

	// Specifies the reason why the instruction has a failing status.
	Code *FailingReason1Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (f *FailingReason1) AddCode() *FailingReason1Choice {
	f.Code = new(FailingReason1Choice)
	return f.Code
}

func (f *FailingReason1) SetAdditionalReasonInformation(value string) {
	f.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason why the instruction or request has a failing settlement status.
type FailingReason10 struct {

	// Specifies the reason why the instruction has a failing status.
	Code *FailingReason11Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (f *FailingReason10) AddCode() *FailingReason11Choice {
	f.Code = new(FailingReason11Choice)
	return f.Code
}

func (f *FailingReason10) SetAdditionalReasonInformation(value string) {
	f.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// The status of an instruction, advice or request.
type FailingReason3 struct {

	// Specifies the reason why the instruction has a failing status.
	Code *FailingReason3Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (f *FailingReason3) AddCode() *FailingReason3Choice {
	f.Code = new(FailingReason3Choice)
	return f.Code
}

func (f *FailingReason3) SetAdditionalReasonInformation(value string) {
	f.AdditionalReasonInformation = (*Max210Text)(&value)
}

// The status of an instruction, advice or request.
type FailingReason4 struct {

	// Specifies the reason why the instruction has a failing status.
	Code *FailingReason4Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (f *FailingReason4) AddCode() *FailingReason4Choice {
	f.Code = new(FailingReason4Choice)
	return f.Code
}

func (f *FailingReason4) SetAdditionalReasonInformation(value string) {
	f.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason why the instruction or request has a failing settlement status.
type FailingReason7 struct {

	// Specifies the reason why the instruction has a failing status.
	Code *FailingReason7Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (f *FailingReason7) AddCode() *FailingReason7Choice {
	f.Code = new(FailingReason7Choice)
	return f.Code
}

func (f *FailingReason7) SetAdditionalReasonInformation(value string) {
	f.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason why the instruction or request has a failing settlement status.
type FailingReason8 struct {

	// Specifies the reason why the instruction has a failing status.
	Code *FailingReason8Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (f *FailingReason8) AddCode() *FailingReason8Choice {
	f.Code = new(FailingReason8Choice)
	return f.Code
}

func (f *FailingReason8) SetAdditionalReasonInformation(value string) {
	f.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason why the instruction or request has a failing settlement status.
type FailingReason9 struct {

	// Specifies the reason why the instruction has a failing status.
	Code *FailingReason10Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (f *FailingReason9) AddCode() *FailingReason10Choice {
	f.Code = new(FailingReason10Choice)
	return f.Code
}

func (f *FailingReason9) SetAdditionalReasonInformation(value string) {
	f.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
