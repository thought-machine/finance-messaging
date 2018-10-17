package iso20022

// Specifies reasons for the rejected status.
type RejectedStatusReason10 struct {

	// Specifies the reason why the instruction or instruction cancellation has been rejected.
	ReasonCode *RejectedReason5Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectedStatusReason10) AddReasonCode() *RejectedReason5Choice {
	r.ReasonCode = new(RejectedReason5Choice)
	return r.ReasonCode
}

func (r *RejectedStatusReason10) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Reason for a rejected status.
type RejectedStatusReason12 struct {

	// Reason for the rejected status.
	Reason *RejectedReason8Choice `xml:"Rsn"`

	// Additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectedStatusReason12) AddReason() *RejectedReason8Choice {
	r.Reason = new(RejectedReason8Choice)
	return r.Reason
}

func (r *RejectedStatusReason12) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies reasons for the rejected status.
type RejectedStatusReason13 struct {

	// Specifies the reason why the instruction or instruction cancellation has been rejected.
	ReasonCode *RejectedReason9Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectedStatusReason13) AddReasonCode() *RejectedReason9Choice {
	r.ReasonCode = new(RejectedReason9Choice)
	return r.ReasonCode
}

func (r *RejectedStatusReason13) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies reasons for the rejected status.
type RejectedStatusReason14 struct {

	// Specifies the reason why the instruction or instruction cancellation has been rejected.
	ReasonCode *RejectedReason10Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectedStatusReason14) AddReasonCode() *RejectedReason10Choice {
	r.ReasonCode = new(RejectedReason10Choice)
	return r.ReasonCode
}

func (r *RejectedStatusReason14) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies reasons for the rejected status.
type RejectedStatusReason17 struct {

	// Specifies the reason why the instruction or instruction cancellation has been rejected.
	ReasonCode *RejectedReason13Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectedStatusReason17) AddReasonCode() *RejectedReason13Choice {
	r.ReasonCode = new(RejectedReason13Choice)
	return r.ReasonCode
}

func (r *RejectedStatusReason17) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies reasons for the rejected status.
type RejectedStatusReason18 struct {

	// Specifies the reason why the instruction or instruction cancellation has been rejected.
	ReasonCode *RejectedReason14Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectedStatusReason18) AddReasonCode() *RejectedReason14Choice {
	r.ReasonCode = new(RejectedReason14Choice)
	return r.ReasonCode
}

func (r *RejectedStatusReason18) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies reasons for the rejected status.
type RejectedStatusReason19 struct {

	// Specifies the reason why the instruction or instruction cancellation has been rejected.
	ReasonCode *RejectedReason18Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectedStatusReason19) AddReasonCode() *RejectedReason18Choice {
	r.ReasonCode = new(RejectedReason18Choice)
	return r.ReasonCode
}

func (r *RejectedStatusReason19) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Specifies reasons for the rejected status.
type RejectedStatusReason20 struct {

	// Specifies the reason why the instruction or instruction cancellation has been rejected.
	ReasonCode *RejectedReason19Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectedStatusReason20) AddReasonCode() *RejectedReason19Choice {
	r.ReasonCode = new(RejectedReason19Choice)
	return r.ReasonCode
}

func (r *RejectedStatusReason20) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Specifies reasons for the rejected status.
type RejectedStatusReason21 struct {

	// Specifies the reason why the instruction or instruction cancellation has been rejected.
	ReasonCode *RejectedReason22Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectedStatusReason21) AddReasonCode() *RejectedReason22Choice {
	r.ReasonCode = new(RejectedReason22Choice)
	return r.ReasonCode
}

func (r *RejectedStatusReason21) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies reasons for the rejected status.
type RejectedStatusReason22 struct {

	// Specifies the reason why the instruction or instruction cancellation has been rejected.
	ReasonCode *RejectedReason23Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectedStatusReason22) AddReasonCode() *RejectedReason23Choice {
	r.ReasonCode = new(RejectedReason23Choice)
	return r.ReasonCode
}

func (r *RejectedStatusReason22) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Reason for a rejected status.
type RejectedStatusReason4 struct {

	// Reason for a rejected status in structured form.
	Structured *RejectedStatusReason4Code `xml:"Strd"`

	// Reason for a rejected status in free format text.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (r *RejectedStatusReason4) SetStructured(value string) {
	r.Structured = (*RejectedStatusReason4Code)(&value)
}

func (r *RejectedStatusReason4) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*Max350Text)(&value)
}

// Reason for a rejected status.
type RejectedStatusReason5 struct {

	// Reason for a rejected status in structured form.
	Structured *TransferRejectedStatusReason1Code `xml:"Strd"`

	// Additional information about the reason for the rejected status in textual form.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (r *RejectedStatusReason5) SetStructured(value string) {
	r.Structured = (*TransferRejectedStatusReason1Code)(&value)
}

func (r *RejectedStatusReason5) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*Max350Text)(&value)
}

// Reason for a rejected status.
type RejectedStatusReason6 struct {

	// Reason for a rejected status in structured form.
	Structured *RejectedStatusReason5Code `xml:"Strd"`

	// Reason for a rejected status in free format text.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (r *RejectedStatusReason6) SetStructured(value string) {
	r.Structured = (*RejectedStatusReason5Code)(&value)
}

func (r *RejectedStatusReason6) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*Max350Text)(&value)
}

// Specifies reasons for the rejected status.
type RejectedStatusReason8 struct {

	// Specifies the reason why the instruction or instruction cancellation has been rejected.
	ReasonCode *RejectedReason1Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectedStatusReason8) AddReasonCode() *RejectedReason1Choice {
	r.ReasonCode = new(RejectedReason1Choice)
	return r.ReasonCode
}

func (r *RejectedStatusReason8) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
