package iso20022

// Modification reasons.
type ModificationReason2 struct {

	// Specifies the reason why the transaction is modified.
	Code *ModificationReason2Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (m *ModificationReason2) AddCode() *ModificationReason2Choice {
	m.Code = new(ModificationReason2Choice)
	return m.Code
}

func (m *ModificationReason2) SetAdditionalReasonInformation(value string) {
	m.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Modification reasons.
type ModificationReason4 struct {

	// Specifies the reason why the transaction is modified.
	Code *ModificationReason4Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (m *ModificationReason4) AddCode() *ModificationReason4Choice {
	m.Code = new(ModificationReason4Choice)
	return m.Code
}

func (m *ModificationReason4) SetAdditionalReasonInformation(value string) {
	m.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Modification reasons.
type ModificationReason5 struct {

	// Specifies the reason why the transaction is modified.
	Code *ModificationReason5Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (m *ModificationReason5) AddCode() *ModificationReason5Choice {
	m.Code = new(ModificationReason5Choice)
	return m.Code
}

func (m *ModificationReason5) SetAdditionalReasonInformation(value string) {
	m.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
