package iso20022

// Provides additional information regarding the party, eg, the contact unit or person responsible for the transaction identified in the message.
type PartyTextInformation1 struct {

	// Provides declaration details narrative relative to the party.
	DeclarationDetails *Max350Text `xml:"DclrtnDtls,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	PartyContactDetails *Max140Text `xml:"PtyCtctDtls,omitempty"`

	// Provides information required for the registration.
	RegistrationDetails *Max350Text `xml:"RegnDtls,omitempty"`
}

func (p *PartyTextInformation1) SetDeclarationDetails(value string) {
	p.DeclarationDetails = (*Max350Text)(&value)
}

func (p *PartyTextInformation1) SetPartyContactDetails(value string) {
	p.PartyContactDetails = (*Max140Text)(&value)
}

func (p *PartyTextInformation1) SetRegistrationDetails(value string) {
	p.RegistrationDetails = (*Max350Text)(&value)
}

// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
type PartyTextInformation2 struct {

	// Provides declaration details narrative relative to the party.
	DeclarationDetails *Max350Text `xml:"DclrtnDtls,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	PartyContactDetails *Max140Text `xml:"PtyCtctDtls,omitempty"`
}

func (p *PartyTextInformation2) SetDeclarationDetails(value string) {
	p.DeclarationDetails = (*Max350Text)(&value)
}

func (p *PartyTextInformation2) SetPartyContactDetails(value string) {
	p.PartyContactDetails = (*Max140Text)(&value)
}

// Provides additional information regarding the party, eg, the contact unit or person responsible for the transaction identified in the message.
type PartyTextInformation3 struct {

	// Provides declaration details narrative relative to the party.
	DeclarationDetails *RestrictedFINXMax350Text `xml:"DclrtnDtls,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	PartyContactDetails *RestrictedFINXMax140Text `xml:"PtyCtctDtls,omitempty"`

	// Provides information required for the registration.
	RegistrationDetails *RestrictedFINXMax350Text `xml:"RegnDtls,omitempty"`
}

func (p *PartyTextInformation3) SetDeclarationDetails(value string) {
	p.DeclarationDetails = (*RestrictedFINXMax350Text)(&value)
}

func (p *PartyTextInformation3) SetPartyContactDetails(value string) {
	p.PartyContactDetails = (*RestrictedFINXMax140Text)(&value)
}

func (p *PartyTextInformation3) SetRegistrationDetails(value string) {
	p.RegistrationDetails = (*RestrictedFINXMax350Text)(&value)
}

// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
type PartyTextInformation4 struct {

	// Provides declaration details narrative relative to the party.
	DeclarationDetails *RestrictedFINXMax350Text `xml:"DclrtnDtls,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	PartyContactDetails *RestrictedFINXMax140Text `xml:"PtyCtctDtls,omitempty"`
}

func (p *PartyTextInformation4) SetDeclarationDetails(value string) {
	p.DeclarationDetails = (*RestrictedFINXMax350Text)(&value)
}

func (p *PartyTextInformation4) SetPartyContactDetails(value string) {
	p.PartyContactDetails = (*RestrictedFINXMax140Text)(&value)
}

// Provides additional information regarding the party, eg, the contact unit or person responsible for the transaction identified in the message.
type PartyTextInformation5 struct {

	// Provides declaration details narrative relative to the party.
	DeclarationDetails *Max350Text `xml:"DclrtnDtls,omitempty"`

	// Provides additional information regarding the party, eg, the contact unit or person responsible for the transaction identified in the message.
	PartyContactDetails *Max140Text `xml:"PtyCtctDtls,omitempty"`
}

func (p *PartyTextInformation5) SetDeclarationDetails(value string) {
	p.DeclarationDetails = (*Max350Text)(&value)
}

func (p *PartyTextInformation5) SetPartyContactDetails(value string) {
	p.PartyContactDetails = (*Max140Text)(&value)
}
