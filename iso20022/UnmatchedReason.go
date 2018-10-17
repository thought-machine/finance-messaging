package iso20022

// The status of an instruction, advice or request.
type UnmatchedReason1 struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason1Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (u *UnmatchedReason1) AddCode() *UnmatchedReason1Choice {
	u.Code = new(UnmatchedReason1Choice)
	return u.Code
}

func (u *UnmatchedReason1) SetAdditionalReasonInformation(value string) {
	u.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Status of an instruction, advice or request.
type UnmatchedReason11 struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason14Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (u *UnmatchedReason11) AddCode() *UnmatchedReason14Choice {
	u.Code = new(UnmatchedReason14Choice)
	return u.Code
}

func (u *UnmatchedReason11) SetAdditionalReasonInformation(value string) {
	u.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Status of an instruction, advice or request.
type UnmatchedReason12 struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason15Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (u *UnmatchedReason12) AddCode() *UnmatchedReason15Choice {
	u.Code = new(UnmatchedReason15Choice)
	return u.Code
}

func (u *UnmatchedReason12) SetAdditionalReasonInformation(value string) {
	u.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Status of an instruction, advice or request.
type UnmatchedReason15 struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason21Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (u *UnmatchedReason15) AddCode() *UnmatchedReason21Choice {
	u.Code = new(UnmatchedReason21Choice)
	return u.Code
}

func (u *UnmatchedReason15) SetAdditionalReasonInformation(value string) {
	u.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason why the instruction has an unmatched status.
type UnmatchedReason16 struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason23Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (u *UnmatchedReason16) AddCode() *UnmatchedReason23Choice {
	u.Code = new(UnmatchedReason23Choice)
	return u.Code
}

func (u *UnmatchedReason16) SetAdditionalReasonInformation(value string) {
	u.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Status of an instruction, advice or request.
type UnmatchedReason17 struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason24Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (u *UnmatchedReason17) AddCode() *UnmatchedReason24Choice {
	u.Code = new(UnmatchedReason24Choice)
	return u.Code
}

func (u *UnmatchedReason17) SetAdditionalReasonInformation(value string) {
	u.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Status of an instruction, advice or request.
type UnmatchedReason18 struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason25Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (u *UnmatchedReason18) AddCode() *UnmatchedReason25Choice {
	u.Code = new(UnmatchedReason25Choice)
	return u.Code
}

func (u *UnmatchedReason18) SetAdditionalReasonInformation(value string) {
	u.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Specifies the reason why the instruction has an unmatched status.
type UnmatchedReason19 struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason26Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (u *UnmatchedReason19) AddCode() *UnmatchedReason26Choice {
	u.Code = new(UnmatchedReason26Choice)
	return u.Code
}

func (u *UnmatchedReason19) SetAdditionalReasonInformation(value string) {
	u.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// The status of an instruction, advice or request.
type UnmatchedReason2 struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason3Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (u *UnmatchedReason2) AddCode() *UnmatchedReason3Choice {
	u.Code = new(UnmatchedReason3Choice)
	return u.Code
}

func (u *UnmatchedReason2) SetAdditionalReasonInformation(value string) {
	u.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Status of an instruction, advice or request.
type UnmatchedReason20 struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason27Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (u *UnmatchedReason20) AddCode() *UnmatchedReason27Choice {
	u.Code = new(UnmatchedReason27Choice)
	return u.Code
}

func (u *UnmatchedReason20) SetAdditionalReasonInformation(value string) {
	u.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// The status of an instruction, advice or request.
type UnmatchedReason5 struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason7Choice `xml:"Cd"`

	// Provides additional information about the reason in narrative form.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (u *UnmatchedReason5) AddCode() *UnmatchedReason7Choice {
	u.Code = new(UnmatchedReason7Choice)
	return u.Code
}

func (u *UnmatchedReason5) SetAdditionalReasonInformation(value string) {
	u.AdditionalReasonInformation = (*Max210Text)(&value)
}

// The status of an instruction, advice or request.
type UnmatchedReason6 struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason9Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (u *UnmatchedReason6) AddCode() *UnmatchedReason9Choice {
	u.Code = new(UnmatchedReason9Choice)
	return u.Code
}

func (u *UnmatchedReason6) SetAdditionalReasonInformation(value string) {
	u.AdditionalReasonInformation = (*Max210Text)(&value)
}

// The status of an instruction, advice or request.
type UnmatchedReason7 struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason10Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (u *UnmatchedReason7) AddCode() *UnmatchedReason10Choice {
	u.Code = new(UnmatchedReason10Choice)
	return u.Code
}

func (u *UnmatchedReason7) SetAdditionalReasonInformation(value string) {
	u.AdditionalReasonInformation = (*Max210Text)(&value)
}
