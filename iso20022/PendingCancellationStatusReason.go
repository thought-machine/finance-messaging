package iso20022

// Specifies reasons for the pending cancellation status.
type PendingCancellationStatusReason1 struct {

	// Specifies the reason why the cancellation request is pending.
	ReasonCode *PendingCancellationReason1Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingCancellationStatusReason1) AddReasonCode() *PendingCancellationReason1Choice {
	p.ReasonCode = new(PendingCancellationReason1Choice)
	return p.ReasonCode
}

func (p *PendingCancellationStatusReason1) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies reasons for the pending cancellation status.
type PendingCancellationStatusReason3 struct {

	// Specifies the reason why the cancellation request is pending.
	ReasonCode *PendingCancellationReason1Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingCancellationStatusReason3) AddReasonCode() *PendingCancellationReason1Choice {
	p.ReasonCode = new(PendingCancellationReason1Choice)
	return p.ReasonCode
}

func (p *PendingCancellationStatusReason3) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies reasons for the pending cancellation status.
type PendingCancellationStatusReason5 struct {

	// Specifies the reason why the cancellation request is pending.
	ReasonCode *PendingCancellationReason3Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingCancellationStatusReason5) AddReasonCode() *PendingCancellationReason3Choice {
	p.ReasonCode = new(PendingCancellationReason3Choice)
	return p.ReasonCode
}

func (p *PendingCancellationStatusReason5) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies reasons for the pending cancellation status.
type PendingCancellationStatusReason6 struct {

	// Specifies the reason why the cancellation request is pending.
	ReasonCode *PendingCancellationReason4Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingCancellationStatusReason6) AddReasonCode() *PendingCancellationReason4Choice {
	p.ReasonCode = new(PendingCancellationReason4Choice)
	return p.ReasonCode
}

func (p *PendingCancellationStatusReason6) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Specifies reasons for the pending cancellation status.
type PendingCancellationStatusReason7 struct {

	// Specifies the reason why the cancellation request is pending.
	ReasonCode *PendingCancellationReason5Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingCancellationStatusReason7) AddReasonCode() *PendingCancellationReason5Choice {
	p.ReasonCode = new(PendingCancellationReason5Choice)
	return p.ReasonCode
}

func (p *PendingCancellationStatusReason7) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies reasons for the pending cancellation status.
type PendingCancellationStatusReason8 struct {

	// Specifies the reason why the cancellation request is pending.
	ReasonCode *PendingCancellationReason6Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingCancellationStatusReason8) AddReasonCode() *PendingCancellationReason6Choice {
	p.ReasonCode = new(PendingCancellationReason6Choice)
	return p.ReasonCode
}

func (p *PendingCancellationStatusReason8) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
