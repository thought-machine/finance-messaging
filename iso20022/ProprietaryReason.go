package iso20022

// Proprietary identification of the reason related to a status.
type ProprietaryReason1 struct {

	// Proprietary identification of the reason related to a status.
	Reason *GenericIdentification20 `xml:"Rsn,omitempty"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *ProprietaryReason1) AddReason() *GenericIdentification20 {
	p.Reason = new(GenericIdentification20)
	return p.Reason
}

func (p *ProprietaryReason1) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Proprietary identification of the reason related to a status.
type ProprietaryReason4 struct {

	// Proprietary identification of the reason related to a status.
	Reason *GenericIdentification30 `xml:"Rsn,omitempty"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *ProprietaryReason4) AddReason() *GenericIdentification30 {
	p.Reason = new(GenericIdentification30)
	return p.Reason
}

func (p *ProprietaryReason4) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Proprietary identification of the reason related to a status.
type ProprietaryReason5 struct {

	// Proprietary identification of the reason related to a status.
	Reason *GenericIdentification47 `xml:"Rsn,omitempty"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *ProprietaryReason5) AddReason() *GenericIdentification47 {
	p.Reason = new(GenericIdentification47)
	return p.Reason
}

func (p *ProprietaryReason5) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
