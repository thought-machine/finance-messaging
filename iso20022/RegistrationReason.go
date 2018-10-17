package iso20022

// Reason of registration.
type RegistrationReason1 struct {

	// Specifies the reaoson of the holding status.
	Code *Registration3Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalInformation *Max210Text `xml:"AddtlInf,omitempty"`
}

func (r *RegistrationReason1) AddCode() *Registration3Choice {
	r.Code = new(Registration3Choice)
	return r.Code
}

func (r *RegistrationReason1) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*Max210Text)(&value)
}

// Reason of registration.
type RegistrationReason3 struct {

	// Specifies the reaoson of the holding status.
	Code *Registration7Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalInformation *Max210Text `xml:"AddtlInf,omitempty"`
}

func (r *RegistrationReason3) AddCode() *Registration7Choice {
	r.Code = new(Registration7Choice)
	return r.Code
}

func (r *RegistrationReason3) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*Max210Text)(&value)
}

// Reason of registration.
type RegistrationReason5 struct {

	// Specifies the reaoson of the holding status.
	Code *Registration10Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalInformation *Max210Text `xml:"AddtlInf,omitempty"`
}

func (r *RegistrationReason5) AddCode() *Registration10Choice {
	r.Code = new(Registration10Choice)
	return r.Code
}

func (r *RegistrationReason5) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*Max210Text)(&value)
}

// Reason of registration.
type RegistrationReason6 struct {

	// Specifies the reaoson of the holding status.
	Code *Registration12Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalInformation *RestrictedFINXMax210Text `xml:"AddtlInf,omitempty"`
}

func (r *RegistrationReason6) AddCode() *Registration12Choice {
	r.Code = new(Registration12Choice)
	return r.Code
}

func (r *RegistrationReason6) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*RestrictedFINXMax210Text)(&value)
}
