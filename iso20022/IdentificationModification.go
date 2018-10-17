package iso20022

// Set of elements used to provide information concerning the identification data that is advised to be modified.
type IdentificationModification1 struct {

	// Unique identification, as assigned by a sending party, to unambigiously identify the party and account identification information group within the message.
	Identification *Max35Text `xml:"Id"`

	// Provides party and/or account identification information as given in the original message.
	OriginalPartyAndAccountIdentification *IdentificationInformation1 `xml:"OrgnlPtyAndAcctId,omitempty"`

	// Provides updated party and/or account identification information.
	UpdatedPartyAndAccountIdentification *IdentificationInformation1 `xml:"UpdtdPtyAndAcctId"`

	// Additional information, in free text form, to complement the modification information.
	AdditionalInformation *Max140Text `xml:"AddtlInf,omitempty"`
}

func (i *IdentificationModification1) SetIdentification(value string) {
	i.Identification = (*Max35Text)(&value)
}

func (i *IdentificationModification1) AddOriginalPartyAndAccountIdentification() *IdentificationInformation1 {
	i.OriginalPartyAndAccountIdentification = new(IdentificationInformation1)
	return i.OriginalPartyAndAccountIdentification
}

func (i *IdentificationModification1) AddUpdatedPartyAndAccountIdentification() *IdentificationInformation1 {
	i.UpdatedPartyAndAccountIdentification = new(IdentificationInformation1)
	return i.UpdatedPartyAndAccountIdentification
}

func (i *IdentificationModification1) SetAdditionalInformation(value string) {
	i.AdditionalInformation = (*Max140Text)(&value)
}

// Provides the details of the identification data that is advised to be modified.
type IdentificationModification2 struct {

	// Unique identification, as assigned by a sending party, to unambiguously identify the party and account identification information group within the message.
	Identification *Max35Text `xml:"Id"`

	// Provides party and/or account identification information as given in the original message.
	OriginalPartyAndAccountIdentification *IdentificationInformation2 `xml:"OrgnlPtyAndAcctId,omitempty"`

	// Provides updated party and/or account identification information.
	UpdatedPartyAndAccountIdentification *IdentificationInformation2 `xml:"UpdtdPtyAndAcctId"`

	// Additional information, in free text form, to complement the modification information.
	AdditionalInformation *Max140Text `xml:"AddtlInf,omitempty"`
}

func (i *IdentificationModification2) SetIdentification(value string) {
	i.Identification = (*Max35Text)(&value)
}

func (i *IdentificationModification2) AddOriginalPartyAndAccountIdentification() *IdentificationInformation2 {
	i.OriginalPartyAndAccountIdentification = new(IdentificationInformation2)
	return i.OriginalPartyAndAccountIdentification
}

func (i *IdentificationModification2) AddUpdatedPartyAndAccountIdentification() *IdentificationInformation2 {
	i.UpdatedPartyAndAccountIdentification = new(IdentificationInformation2)
	return i.UpdatedPartyAndAccountIdentification
}

func (i *IdentificationModification2) SetAdditionalInformation(value string) {
	i.AdditionalInformation = (*Max140Text)(&value)
}
