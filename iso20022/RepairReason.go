package iso20022

// The status of an instruction, advice or request.
type RepairReason1 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RepairReason1Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RepairReason1) AddCode() *RepairReason1Choice {
	r.Code = new(RepairReason1Choice)
	return r.Code
}

func (r *RepairReason1) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason why the instruction is in repair.
type RepairReason10 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RepairReason12Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RepairReason10) AddCode() *RepairReason12Choice {
	r.Code = new(RepairReason12Choice)
	return r.Code
}

func (r *RepairReason10) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason why the instruction is in repair.
type RepairReason11 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RepairReason13Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RepairReason11) AddCode() *RepairReason13Choice {
	r.Code = new(RepairReason13Choice)
	return r.Code
}

func (r *RepairReason11) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Specifies the reason why the instruction is in repair.
type RepairReason12 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RepairReason14Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RepairReason12) AddCode() *RepairReason14Choice {
	r.Code = new(RepairReason14Choice)
	return r.Code
}

func (r *RepairReason12) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Specifies the reason why the instruction is in repair.
type RepairReason13 struct {

	// Specifies the reason why the instruction/request has a repair status.
	Code *RepairReason14Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RepairReason13) AddCode() *RepairReason14Choice {
	r.Code = new(RepairReason14Choice)
	return r.Code
}

func (r *RepairReason13) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// The status of an instruction, advice or request.
type RepairReason3 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RepairReason2Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RepairReason3) AddCode() *RepairReason2Choice {
	r.Code = new(RepairReason2Choice)
	return r.Code
}

func (r *RepairReason3) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// The status of an instruction, advice or request.
type RepairReason5 struct {

	// Specifies the reason why the instruction is in repair.
	Code *RepairReason9Choice `xml:"Cd"`

	// Provides additional information about the reason in narrative form.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RepairReason5) AddCode() *RepairReason9Choice {
	r.Code = new(RepairReason9Choice)
	return r.Code
}

func (r *RepairReason5) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// The status of an instruction, advice or request.
type RepairReason6 struct {

	// Specifies the reason why the instruction/request has a repair status.
	Code *RepairReason1Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RepairReason6) AddCode() *RepairReason1Choice {
	r.Code = new(RepairReason1Choice)
	return r.Code
}

func (r *RepairReason6) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason why the instruction is in repair.
type RepairReason8 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RepairReason10Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RepairReason8) AddCode() *RepairReason10Choice {
	r.Code = new(RepairReason10Choice)
	return r.Code
}

func (r *RepairReason8) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason why the instruction is in repair.
type RepairReason9 struct {

	// Specifies the reason why the instruction/request has a repair status.
	Code *RepairReason10Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RepairReason9) AddCode() *RepairReason10Choice {
	r.Code = new(RepairReason10Choice)
	return r.Code
}

func (r *RepairReason9) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
