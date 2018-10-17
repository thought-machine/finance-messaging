package iso20022

// Determine the type of document and the type of communication method to be used to notify a Party.
type DocumentToSend1 struct {

	// Type of document.
	Type *Max140Text `xml:"Tp"`

	// Party that should receive the document.
	Recipient *PartyIdentification2Choice `xml:"Rcpt"`

	// Communication method to be used.
	MethodOfTransmission *CommunicationMethod1Code `xml:"MtdOfTrnsmssn"`

	// Communication means used to send information.
	ExtendedMethodOfTransmission *Extended350Code `xml:"XtndedMtdOfTrnsmssn"`
}

func (d *DocumentToSend1) SetType(value string) {
	d.Type = (*Max140Text)(&value)
}

func (d *DocumentToSend1) AddRecipient() *PartyIdentification2Choice {
	d.Recipient = new(PartyIdentification2Choice)
	return d.Recipient
}

func (d *DocumentToSend1) SetMethodOfTransmission(value string) {
	d.MethodOfTransmission = (*CommunicationMethod1Code)(&value)
}

func (d *DocumentToSend1) SetExtendedMethodOfTransmission(value string) {
	d.ExtendedMethodOfTransmission = (*Extended350Code)(&value)
}

// Determine the type of document and the type of communication method to be used to notify a Party.
type DocumentToSend2 struct {

	// Type of document.
	Type *Max140Text `xml:"Tp"`

	// Party that should receive the document.
	Recipient *PartyIdentification2Choice `xml:"Rcpt"`

	// Communication method to be used.
	MethodOfTransmission *CommunicationMethod3Choice `xml:"MtdOfTrnsmssn"`
}

func (d *DocumentToSend2) SetType(value string) {
	d.Type = (*Max140Text)(&value)
}

func (d *DocumentToSend2) AddRecipient() *PartyIdentification2Choice {
	d.Recipient = new(PartyIdentification2Choice)
	return d.Recipient
}

func (d *DocumentToSend2) AddMethodOfTransmission() *CommunicationMethod3Choice {
	d.MethodOfTransmission = new(CommunicationMethod3Choice)
	return d.MethodOfTransmission
}

// Type of document and the type of communication method to be used to notify a party.
type DocumentToSend3 struct {

	// Type of document.
	Type *Max140Text `xml:"Tp"`

	// Party that should receive the document.
	Recipient *PartyIdentification70Choice `xml:"Rcpt"`

	// Communication method to be used.
	MethodOfTransmission *CommunicationMethod3Choice `xml:"MtdOfTrnsmssn"`
}

func (d *DocumentToSend3) SetType(value string) {
	d.Type = (*Max140Text)(&value)
}

func (d *DocumentToSend3) AddRecipient() *PartyIdentification70Choice {
	d.Recipient = new(PartyIdentification70Choice)
	return d.Recipient
}

func (d *DocumentToSend3) AddMethodOfTransmission() *CommunicationMethod3Choice {
	d.MethodOfTransmission = new(CommunicationMethod3Choice)
	return d.MethodOfTransmission
}
