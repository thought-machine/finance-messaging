package iso20022

// Specifies reasons for the pending status.
type PendingStatusReason1 struct {

	// Specifies the reason why the instruction's processing is pending.
	ReasonCode *PendingReason6Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingStatusReason1) AddReasonCode() *PendingReason6Choice {
	p.ReasonCode = new(PendingReason6Choice)
	return p.ReasonCode
}

func (p *PendingStatusReason1) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies reasons for the pending status.
type PendingStatusReason10 struct {

	// Specifies the reason why the instruction's processing is pending.
	ReasonCode *PendingReason33Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingStatusReason10) AddReasonCode() *PendingReason33Choice {
	p.ReasonCode = new(PendingReason33Choice)
	return p.ReasonCode
}

func (p *PendingStatusReason10) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies reasons for the pending status.
type PendingStatusReason11 struct {

	// Specifies the reason why the instruction's processing is pending.
	ReasonCode *PendingReason34Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingStatusReason11) AddReasonCode() *PendingReason34Choice {
	p.ReasonCode = new(PendingReason34Choice)
	return p.ReasonCode
}

func (p *PendingStatusReason11) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Specifies reasons for the pending status.
type PendingStatusReason12 struct {

	// Specifies the reason why the instruction's processing is pending.
	ReasonCode *PendingReason35Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingStatusReason12) AddReasonCode() *PendingReason35Choice {
	p.ReasonCode = new(PendingReason35Choice)
	return p.ReasonCode
}

func (p *PendingStatusReason12) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Specifies reasons for the pending status.
type PendingStatusReason13 struct {

	// Specifies the reason why the instruction's processing is pending.
	ReasonCode *PendingReason48Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingStatusReason13) AddReasonCode() *PendingReason48Choice {
	p.ReasonCode = new(PendingReason48Choice)
	return p.ReasonCode
}

func (p *PendingStatusReason13) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Reason for a pending status.
type PendingStatusReason14 struct {

	// Reason for the pending account status.
	Code *PendingStatusReason2Choice `xml:"Cd"`

	// Additional information about the reason for the pending account status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (p *PendingStatusReason14) AddCode() *PendingStatusReason2Choice {
	p.Code = new(PendingStatusReason2Choice)
	return p.Code
}

func (p *PendingStatusReason14) SetAdditionalInformation(value string) {
	p.AdditionalInformation = (*Max350Text)(&value)
}

// Specifies reasons for the pending status.
type PendingStatusReason15 struct {

	// Specifies the reason why the instruction's processing is pending.
	ReasonCode *PendingReason49Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingStatusReason15) AddReasonCode() *PendingReason49Choice {
	p.ReasonCode = new(PendingReason49Choice)
	return p.ReasonCode
}

func (p *PendingStatusReason15) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Specifies reasons for the pending status.
type PendingStatusReason2 struct {

	// Specifies the reason why the instruction's processing is pending.
	ReasonCode *PendingReason4Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingStatusReason2) AddReasonCode() *PendingReason4Choice {
	p.ReasonCode = new(PendingReason4Choice)
	return p.ReasonCode
}

func (p *PendingStatusReason2) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies reasons for the pending status.
type PendingStatusReason5 struct {

	// Specifies the reason why the instruction's processing is pending.
	ReasonCode *PendingReason22Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingStatusReason5) AddReasonCode() *PendingReason22Choice {
	p.ReasonCode = new(PendingReason22Choice)
	return p.ReasonCode
}

func (p *PendingStatusReason5) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies reasons for the pending status.
type PendingStatusReason7 struct {

	// Specifies the reason why the instruction's processing is pending.
	ReasonCode *PendingReason24Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingStatusReason7) AddReasonCode() *PendingReason24Choice {
	p.ReasonCode = new(PendingReason24Choice)
	return p.ReasonCode
}

func (p *PendingStatusReason7) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies reasons for the pending status.
type PendingStatusReason9 struct {

	// Specifies the reason why the instruction's processing is pending.
	ReasonCode *PendingReason32Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingStatusReason9) AddReasonCode() *PendingReason32Choice {
	p.ReasonCode = new(PendingReason32Choice)
	return p.ReasonCode
}

func (p *PendingStatusReason9) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}
