package iso20022

// Set of elements used to provide information concerning the identification data that is requested to be verified.
type IdentificationVerification1 struct {

	// Unique identification, as assigned by a sending party, to unambiguously identify the party and account identification information group within the message.
	Identification *Max35Text `xml:"Id"`

	// Party and/or account identification information for which verification is requested.
	PartyAndAccountIdentification *IdentificationInformation1 `xml:"PtyAndAcctId"`
}

func (i *IdentificationVerification1) SetIdentification(value string) {
	i.Identification = (*Max35Text)(&value)
}

func (i *IdentificationVerification1) AddPartyAndAccountIdentification() *IdentificationInformation1 {
	i.PartyAndAccountIdentification = new(IdentificationInformation1)
	return i.PartyAndAccountIdentification
}

// Provides the details of the identification data that is requested to be verified.
type IdentificationVerification2 struct {

	// Unique identification, as assigned by a sending party, to unambiguously identify the party and account identification information group within the message.
	Identification *Max35Text `xml:"Id"`

	// Party and/or account identification information for which verification is requested.
	PartyAndAccountIdentification *IdentificationInformation2 `xml:"PtyAndAcctId"`
}

func (i *IdentificationVerification2) SetIdentification(value string) {
	i.Identification = (*Max35Text)(&value)
}

func (i *IdentificationVerification2) AddPartyAndAccountIdentification() *IdentificationInformation2 {
	i.PartyAndAccountIdentification = new(IdentificationInformation2)
	return i.PartyAndAccountIdentification
}
