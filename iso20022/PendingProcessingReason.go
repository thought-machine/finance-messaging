package iso20022

// The status of an instruction, advice or request.
type PendingProcessingReason1 struct {

	// Specifies the reason why the instruction has a pending processing status.
	Code *PendingProcessingReason1Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingProcessingReason1) AddCode() *PendingProcessingReason1Choice {
	p.Code = new(PendingProcessingReason1Choice)
	return p.Code
}

func (p *PendingProcessingReason1) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason why the instruction or request has a pending processing status.
type PendingProcessingReason10 struct {

	// Specifies the reason why the instruction has a pending processing status.
	Code *PendingProcessingReason12Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingProcessingReason10) AddCode() *PendingProcessingReason12Choice {
	p.Code = new(PendingProcessingReason12Choice)
	return p.Code
}

func (p *PendingProcessingReason10) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason why the instruction or request has a pending processing status.
type PendingProcessingReason11 struct {

	// Specifies the reason why the instruction has a pending processing status.
	Code *PendingProcessingReason13Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingProcessingReason11) AddCode() *PendingProcessingReason13Choice {
	p.Code = new(PendingProcessingReason13Choice)
	return p.Code
}

func (p *PendingProcessingReason11) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Specifies the reason why the instruction or request has a pending processing status.
type PendingProcessingReason12 struct {

	// Specifies the reason why the instruction has a pending processing status.
	Code *PendingProcessingReason14Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingProcessingReason12) AddCode() *PendingProcessingReason14Choice {
	p.Code = new(PendingProcessingReason14Choice)
	return p.Code
}

func (p *PendingProcessingReason12) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Specifies the reason why the instruction or request has a pending processing status.
type PendingProcessingReason13 struct {

	// Specifies the reason why the instruction has a pending processing status.
	Code *PendingProcessingReason15Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingProcessingReason13) AddCode() *PendingProcessingReason15Choice {
	p.Code = new(PendingProcessingReason15Choice)
	return p.Code
}

func (p *PendingProcessingReason13) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// The status of an instruction, advice or request.
type PendingProcessingReason3 struct {

	// Specifies the reason why the instruction has a pending processing status.
	Code *PendingProcessingReason3Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingProcessingReason3) AddCode() *PendingProcessingReason3Choice {
	p.Code = new(PendingProcessingReason3Choice)
	return p.Code
}

func (p *PendingProcessingReason3) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}

// The status of an instruction, advice or request.
type PendingProcessingReason5 struct {

	// Specifies the reason why the instruction has a pending processing status.
	Code *PendingProcessingReason5Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingProcessingReason5) AddCode() *PendingProcessingReason5Choice {
	p.Code = new(PendingProcessingReason5Choice)
	return p.Code
}

func (p *PendingProcessingReason5) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason why the instruction or request has a pending processing status.
type PendingProcessingReason8 struct {

	// Specifies the reason why the instruction has a pending processing status.
	Code *PendingProcessingReason10Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingProcessingReason8) AddCode() *PendingProcessingReason10Choice {
	p.Code = new(PendingProcessingReason10Choice)
	return p.Code
}

func (p *PendingProcessingReason8) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason why the instruction or request has a pending processing status.
type PendingProcessingReason9 struct {

	// Specifies the reason why the instruction has a pending processing status.
	Code *PendingProcessingReason11Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingProcessingReason9) AddCode() *PendingProcessingReason11Choice {
	p.Code = new(PendingProcessingReason11Choice)
	return p.Code
}

func (p *PendingProcessingReason9) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}
