package iso20022

// Specifies reasons for the accepted status.
type AcceptedStatusReason1 struct {

	// Specifies the reason why the instruction or instruction cancellation has been accepted.
	ReasonCode *AcceptedReason1Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcceptedStatusReason1) AddReasonCode() *AcceptedReason1Choice {
	a.ReasonCode = new(AcceptedReason1Choice)
	return a.ReasonCode
}

func (a *AcceptedStatusReason1) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies reasons for the accepted status.
type AcceptedStatusReason10 struct {

	// Specifies the reason why the instruction or instruction cancellation has been accepted.
	ReasonCode *AcceptedReason11Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcceptedStatusReason10) AddReasonCode() *AcceptedReason11Choice {
	a.ReasonCode = new(AcceptedReason11Choice)
	return a.ReasonCode
}

func (a *AcceptedStatusReason10) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Specifies reasons for the accepted status.
type AcceptedStatusReason11 struct {

	// Specifies the reason why the instruction has been accepted.
	ReasonCode *AcceptedReason12Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcceptedStatusReason11) AddReasonCode() *AcceptedReason12Choice {
	a.ReasonCode = new(AcceptedReason12Choice)
	return a.ReasonCode
}

func (a *AcceptedStatusReason11) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Specifies reasons for the accepted status.
type AcceptedStatusReason3 struct {

	// Specifies the reason why the instruction or instruction cancellation has been accepted.
	ReasonCode *AcceptedReason3Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcceptedStatusReason3) AddReasonCode() *AcceptedReason3Choice {
	a.ReasonCode = new(AcceptedReason3Choice)
	return a.ReasonCode
}

func (a *AcceptedStatusReason3) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies reasons for the accepted status.
type AcceptedStatusReason4 struct {

	// Specifies the reason why the instruction has been accepted.
	ReasonCode *AcceptedReason4Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcceptedStatusReason4) AddReasonCode() *AcceptedReason4Choice {
	a.ReasonCode = new(AcceptedReason4Choice)
	return a.ReasonCode
}

func (a *AcceptedStatusReason4) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Reason for an accepted status.
type AcceptedStatusReason7 struct {

	// Reason for the accepted status.
	Reason *AcceptedReason8Choice `xml:"Rsn"`

	// Additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcceptedStatusReason7) AddReason() *AcceptedReason8Choice {
	a.Reason = new(AcceptedReason8Choice)
	return a.Reason
}

func (a *AcceptedStatusReason7) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies reasons for the accepted status.
type AcceptedStatusReason8 struct {

	// Specifies the reason why the instruction has been accepted.
	ReasonCode *AcceptedReason9Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcceptedStatusReason8) AddReasonCode() *AcceptedReason9Choice {
	a.ReasonCode = new(AcceptedReason9Choice)
	return a.ReasonCode
}

func (a *AcceptedStatusReason8) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies reasons for the accepted status.
type AcceptedStatusReason9 struct {

	// Specifies the reason why the instruction or instruction cancellation has been accepted.
	ReasonCode *AcceptedReason10Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcceptedStatusReason9) AddReasonCode() *AcceptedReason10Choice {
	a.ReasonCode = new(AcceptedReason10Choice)
	return a.ReasonCode
}

func (a *AcceptedStatusReason9) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max210Text)(&value)
}
