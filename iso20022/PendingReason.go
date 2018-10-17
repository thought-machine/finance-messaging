package iso20022

// The status of an instruction, advice or request.
type PendingReason1 struct {

	// Specifies the reason why the instruction has a pending status.
	Code *PendingReason1Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingReason1) AddCode() *PendingReason1Choice {
	p.Code = new(PendingReason1Choice)
	return p.Code
}

func (p *PendingReason1) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason why the instruction or request has a pending status.
type PendingReason14 struct {

	// Specifies the reason why the instruction has a pending status.
	Code *PendingReason26Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingReason14) AddCode() *PendingReason26Choice {
	p.Code = new(PendingReason26Choice)
	return p.Code
}

func (p *PendingReason14) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason why the instruction or request has a pending status.
type PendingReason15 struct {

	// Specifies the reason why the instruction has a pending status.
	Code *PendingReason27Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingReason15) AddCode() *PendingReason27Choice {
	p.Code = new(PendingReason27Choice)
	return p.Code
}

func (p *PendingReason15) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason why the instruction or request has a pending status.
type PendingReason16 struct {

	// Specifies the reason why a cancellation request sent for the related instruction is pending.
	Code *PendingReason28Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingReason16) AddCode() *PendingReason28Choice {
	p.Code = new(PendingReason28Choice)
	return p.Code
}

func (p *PendingReason16) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason why the instruction or request has a pending status.
type PendingReason17 struct {

	// Specifies the reason why a cancellation request sent for the related instruction is pending.
	Code *PendingReason30Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingReason17) AddCode() *PendingReason30Choice {
	p.Code = new(PendingReason30Choice)
	return p.Code
}

func (p *PendingReason17) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason why the instruction or request has a pending status.
type PendingReason18 struct {

	// Specifies the reason why the instruction has a pending status.
	Code *PendingReason31Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingReason18) AddCode() *PendingReason31Choice {
	p.Code = new(PendingReason31Choice)
	return p.Code
}

func (p *PendingReason18) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason why the instruction or request has a pending status.
type PendingReason19 struct {

	// Specifies the reason why the instruction has a pending status.
	Code *PendingReason36Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingReason19) AddCode() *PendingReason36Choice {
	p.Code = new(PendingReason36Choice)
	return p.Code
}

func (p *PendingReason19) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Specifies the reason of the PendingStatus.
type PendingReason2 struct {

	// Specifies the reason why a cancellation request sent for the related instruction is pending.
	Code *PendingReason2Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingReason2) AddCode() *PendingReason2Choice {
	p.Code = new(PendingReason2Choice)
	return p.Code
}

func (p *PendingReason2) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason why the instruction or request has a pending status.
type PendingReason20 struct {

	// Specifies the reason why a cancellation request sent for the related instruction is pending.
	Code *PendingReason37Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingReason20) AddCode() *PendingReason37Choice {
	p.Code = new(PendingReason37Choice)
	return p.Code
}

func (p *PendingReason20) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Specifies the reason why the instruction or request has a pending status.
type PendingReason21 struct {

	// Specifies the reason why the instruction has a pending status.
	Code *PendingReason38Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingReason21) AddCode() *PendingReason38Choice {
	p.Code = new(PendingReason38Choice)
	return p.Code
}

func (p *PendingReason21) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Specifies the reason why the instruction or request has a pending status.
type PendingReason24 struct {

	// Specifies the reason why the instruction has a pending status.
	Code *PendingReason41Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingReason24) AddCode() *PendingReason41Choice {
	p.Code = new(PendingReason41Choice)
	return p.Code
}

func (p *PendingReason24) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Specifies the reason why the instruction or request has a pending status.
type PendingReason25 struct {

	// Specifies the reason why a cancellation request sent for the related instruction is pending.
	Code *PendingReason42Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingReason25) AddCode() *PendingReason42Choice {
	p.Code = new(PendingReason42Choice)
	return p.Code
}

func (p *PendingReason25) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// The status of an instruction, advice or request.
type PendingReason5 struct {

	// Specifies the reason why the instruction has a pending status.
	Code *PendingReason13Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingReason5) AddCode() *PendingReason13Choice {
	p.Code = new(PendingReason13Choice)
	return p.Code
}

func (p *PendingReason5) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason of the PendingStatus.
type PendingReason7 struct {

	// Specifies the reason why a cancellation request sent for the related instruction is pending.
	Code *PendingReason15Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingReason7) AddCode() *PendingReason15Choice {
	p.Code = new(PendingReason15Choice)
	return p.Code
}

func (p *PendingReason7) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}

// The status of an instruction, advice or request.
type PendingReason8 struct {

	// Specifies the reason why the instruction has a pending status.
	Code *PendingReason16Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingReason8) AddCode() *PendingReason16Choice {
	p.Code = new(PendingReason16Choice)
	return p.Code
}

func (p *PendingReason8) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}
