package iso20022

// The status of an instruction, advice or request.
type DeniedReason1 struct {

	// Specifies the reason why the request has a denied status.
	Code *DeniedReason3Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (d *DeniedReason1) AddCode() *DeniedReason3Choice {
	d.Code = new(DeniedReason3Choice)
	return d.Code
}

func (d *DeniedReason1) SetAdditionalReasonInformation(value string) {
	d.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason why the request or instruction was denied.
type DeniedReason10 struct {

	// Specifies the reason why the request has a denied status.
	Code *DeniedReason15Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (d *DeniedReason10) AddCode() *DeniedReason15Choice {
	d.Code = new(DeniedReason15Choice)
	return d.Code
}

func (d *DeniedReason10) SetAdditionalReasonInformation(value string) {
	d.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason why the request or instruction was denied.
type DeniedReason11 struct {

	// Specifies the reason why the request has a denied status.
	Code *DeniedReason16Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (d *DeniedReason11) AddCode() *DeniedReason16Choice {
	d.Code = new(DeniedReason16Choice)
	return d.Code
}

func (d *DeniedReason11) SetAdditionalReasonInformation(value string) {
	d.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason why the request or instruction was denied.
type DeniedReason12 struct {

	// Specifies the reason why the request has a denied status.
	Code *DeniedReason17Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (d *DeniedReason12) AddCode() *DeniedReason17Choice {
	d.Code = new(DeniedReason17Choice)
	return d.Code
}

func (d *DeniedReason12) SetAdditionalReasonInformation(value string) {
	d.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason why the request or instruction was denied.
type DeniedReason13 struct {

	// Specifies the reason why the request has a denied status.
	Code *DeniedReason18Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (d *DeniedReason13) AddCode() *DeniedReason18Choice {
	d.Code = new(DeniedReason18Choice)
	return d.Code
}

func (d *DeniedReason13) SetAdditionalReasonInformation(value string) {
	d.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Specifies the reason why the request or instruction was denied.
type DeniedReason16 struct {

	// Specifies the reason why the request has a denied status.
	Code *DeniedReason21Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (d *DeniedReason16) AddCode() *DeniedReason21Choice {
	d.Code = new(DeniedReason21Choice)
	return d.Code
}

func (d *DeniedReason16) SetAdditionalReasonInformation(value string) {
	d.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// The status of an instruction, advice or request.
type DeniedReason17 struct {

	// Specifies the reason why the request has a denied status.
	Code *DeniedReason24Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (d *DeniedReason17) AddCode() *DeniedReason24Choice {
	d.Code = new(DeniedReason24Choice)
	return d.Code
}

func (d *DeniedReason17) SetAdditionalReasonInformation(value string) {
	d.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// The status of an instruction, advice or request.
type DeniedReason2 struct {

	// Specifies the reason why the request has a denied status.
	Code *DeniedReason2Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (d *DeniedReason2) AddCode() *DeniedReason2Choice {
	d.Code = new(DeniedReason2Choice)
	return d.Code
}

func (d *DeniedReason2) SetAdditionalReasonInformation(value string) {
	d.AdditionalReasonInformation = (*Max210Text)(&value)
}

// The status of an instruction, advice or request.
type DeniedReason5 struct {

	// Specifies the reason why the request has a denied status.
	Code *DeniedReason7Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (d *DeniedReason5) AddCode() *DeniedReason7Choice {
	d.Code = new(DeniedReason7Choice)
	return d.Code
}

func (d *DeniedReason5) SetAdditionalReasonInformation(value string) {
	d.AdditionalReasonInformation = (*Max210Text)(&value)
}
