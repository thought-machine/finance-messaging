package iso20022

// Reason for the rejection or repair status.
type RejectionOrRepairReason1 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionAndRepairReason1Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionOrRepairReason1) AddCode() *RejectionAndRepairReason1Choice {
	r.Code = new(RejectionAndRepairReason1Choice)
	return r.Code
}

func (r *RejectionOrRepairReason1) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Reason for the rejection or repair status.
type RejectionOrRepairReason10 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionAndRepairReason10Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionOrRepairReason10) AddCode() *RejectionAndRepairReason10Choice {
	r.Code = new(RejectionAndRepairReason10Choice)
	return r.Code
}

func (r *RejectionOrRepairReason10) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Reason for the rejection or repair status.
type RejectionOrRepairReason13 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionAndRepairReason13Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionOrRepairReason13) AddCode() *RejectionAndRepairReason13Choice {
	r.Code = new(RejectionAndRepairReason13Choice)
	return r.Code
}

func (r *RejectionOrRepairReason13) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Reason for the rejection or repair status.
type RejectionOrRepairReason14 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code []*RejectionAndRepairReason14Choice `xml:"Cd,omitempty"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionOrRepairReason14) AddCode() *RejectionAndRepairReason14Choice {
	newValue := new(RejectionAndRepairReason14Choice)
	r.Code = append(r.Code, newValue)
	return newValue
}

func (r *RejectionOrRepairReason14) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Reason for the rejection or repair status.
type RejectionOrRepairReason18 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionAndRepairReason18Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionOrRepairReason18) AddCode() *RejectionAndRepairReason18Choice {
	r.Code = new(RejectionAndRepairReason18Choice)
	return r.Code
}

func (r *RejectionOrRepairReason18) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Reason for the rejection or repair status.
type RejectionOrRepairReason2 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionAndRepairReason2Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionOrRepairReason2) AddCode() *RejectionAndRepairReason2Choice {
	r.Code = new(RejectionAndRepairReason2Choice)
	return r.Code
}

func (r *RejectionOrRepairReason2) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Reason for the rejection or repair status.
type RejectionOrRepairReason23 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code []*RejectionAndRepairReason23Choice `xml:"Cd,omitempty"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionOrRepairReason23) AddCode() *RejectionAndRepairReason23Choice {
	newValue := new(RejectionAndRepairReason23Choice)
	r.Code = append(r.Code, newValue)
	return newValue
}

func (r *RejectionOrRepairReason23) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Reason for the rejection or repair status.
type RejectionOrRepairReason24 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionAndRepairReason24Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionOrRepairReason24) AddCode() *RejectionAndRepairReason24Choice {
	r.Code = new(RejectionAndRepairReason24Choice)
	return r.Code
}

func (r *RejectionOrRepairReason24) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Reason for the rejection or repair status.
type RejectionOrRepairReason25 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionAndRepairReason25Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionOrRepairReason25) AddCode() *RejectionAndRepairReason25Choice {
	r.Code = new(RejectionAndRepairReason25Choice)
	return r.Code
}

func (r *RejectionOrRepairReason25) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Reason for the rejection or repair status.
type RejectionOrRepairReason26 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionAndRepairReason26Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionOrRepairReason26) AddCode() *RejectionAndRepairReason26Choice {
	r.Code = new(RejectionAndRepairReason26Choice)
	return r.Code
}

func (r *RejectionOrRepairReason26) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Reason for the rejection or repair status.
type RejectionOrRepairReason27 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code []*RejectionAndRepairReason27Choice `xml:"Cd,omitempty"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionOrRepairReason27) AddCode() *RejectionAndRepairReason27Choice {
	newValue := new(RejectionAndRepairReason27Choice)
	r.Code = append(r.Code, newValue)
	return newValue
}

func (r *RejectionOrRepairReason27) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Reason for the rejection or repair status.
type RejectionOrRepairReason28 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionAndRepairReason28Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionOrRepairReason28) AddCode() *RejectionAndRepairReason28Choice {
	r.Code = new(RejectionAndRepairReason28Choice)
	return r.Code
}

func (r *RejectionOrRepairReason28) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Reason for the rejection or repair status.
type RejectionOrRepairReason29 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionAndRepairReason29Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionOrRepairReason29) AddCode() *RejectionAndRepairReason29Choice {
	r.Code = new(RejectionAndRepairReason29Choice)
	return r.Code
}

func (r *RejectionOrRepairReason29) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Reason for the rejection or repair status.
type RejectionOrRepairReason3 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionAndRepairReason3Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionOrRepairReason3) AddCode() *RejectionAndRepairReason3Choice {
	r.Code = new(RejectionAndRepairReason3Choice)
	return r.Code
}

func (r *RejectionOrRepairReason3) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Reason for the rejection or repair status.
type RejectionOrRepairReason31 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionAndRepairReason31Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionOrRepairReason31) AddCode() *RejectionAndRepairReason31Choice {
	r.Code = new(RejectionAndRepairReason31Choice)
	return r.Code
}

func (r *RejectionOrRepairReason31) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Reason for the rejection or repair status.
type RejectionOrRepairReason4 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code []*RejectionAndRepairReason4Choice `xml:"Cd,omitempty"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionOrRepairReason4) AddCode() *RejectionAndRepairReason4Choice {
	newValue := new(RejectionAndRepairReason4Choice)
	r.Code = append(r.Code, newValue)
	return newValue
}

func (r *RejectionOrRepairReason4) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Reason for the rejection or repair status.
type RejectionOrRepairReason9 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionAndRepairReason9Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionOrRepairReason9) AddCode() *RejectionAndRepairReason9Choice {
	r.Code = new(RejectionAndRepairReason9Choice)
	return r.Code
}

func (r *RejectionOrRepairReason9) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
