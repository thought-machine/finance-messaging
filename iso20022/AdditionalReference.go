package iso20022

// References a related message or provides another reference, such as a pool reference, linking a set of messages. The party which issued the related reference may be the Sender of the referenced message or a party other than the Sender.
type AdditionalReference2 struct {

	// Business reference of a message assigned by the party issuing the message. This reference must be unique amongst all messages of the same name sent by the same party.
	Reference *Max35Text `xml:"Ref"`

	// Issuer of the reference.
	ReferenceIssuer *PartyIdentification1Choice `xml:"RefIssr,omitempty"`

	// Name of a message.
	MessageName *Max35Text `xml:"MsgNm,omitempty"`
}

func (a *AdditionalReference2) SetReference(value string) {
	a.Reference = (*Max35Text)(&value)
}

func (a *AdditionalReference2) AddReferenceIssuer() *PartyIdentification1Choice {
	a.ReferenceIssuer = new(PartyIdentification1Choice)
	return a.ReferenceIssuer
}

func (a *AdditionalReference2) SetMessageName(value string) {
	a.MessageName = (*Max35Text)(&value)
}

// References a related message or provides another reference, such as a pool reference, linking a set of messages. The party which issued the related reference may be the Sender of the referenced message or a party other than the Sender.
type AdditionalReference3 struct {

	// Business reference of a message assigned by the party issuing the message. This reference must be unique amongst all messages of the same name sent by the same party.
	Reference *Max35Text `xml:"Ref"`

	// Issuer of the reference.
	ReferenceIssuer *PartyIdentification2Choice `xml:"RefIssr,omitempty"`

	// Name of a message.
	MessageName *Max35Text `xml:"MsgNm,omitempty"`
}

func (a *AdditionalReference3) SetReference(value string) {
	a.Reference = (*Max35Text)(&value)
}

func (a *AdditionalReference3) AddReferenceIssuer() *PartyIdentification2Choice {
	a.ReferenceIssuer = new(PartyIdentification2Choice)
	return a.ReferenceIssuer
}

func (a *AdditionalReference3) SetMessageName(value string) {
	a.MessageName = (*Max35Text)(&value)
}

// Reference to a related message or transaction.
type AdditionalReference6 struct {

	// Message identification of a message. This reference was assigned by the party issuing the message.
	Reference *Max35Text `xml:"Ref"`

	// Issuer of the reference.
	ReferenceIssuer *PartyIdentification90Choice `xml:"RefIssr,omitempty"`

	// Name of the message.
	MessageName *Max35Text `xml:"MsgNm,omitempty"`
}

func (a *AdditionalReference6) SetReference(value string) {
	a.Reference = (*Max35Text)(&value)
}

func (a *AdditionalReference6) AddReferenceIssuer() *PartyIdentification90Choice {
	a.ReferenceIssuer = new(PartyIdentification90Choice)
	return a.ReferenceIssuer
}

func (a *AdditionalReference6) SetMessageName(value string) {
	a.MessageName = (*Max35Text)(&value)
}

// Reference to a related message or transaction.
type AdditionalReference7 struct {

	// Reference issued by a party to identify an instruction, transaction or a message.
	Reference *Max35Text `xml:"Ref"`

	// Party that issued the reference.
	ReferenceIssuer *PartyIdentification97Choice `xml:"RefIssr,omitempty"`

	// Name of the message.
	MessageName *Max35Text `xml:"MsgNm,omitempty"`
}

func (a *AdditionalReference7) SetReference(value string) {
	a.Reference = (*Max35Text)(&value)
}

func (a *AdditionalReference7) AddReferenceIssuer() *PartyIdentification97Choice {
	a.ReferenceIssuer = new(PartyIdentification97Choice)
	return a.ReferenceIssuer
}

func (a *AdditionalReference7) SetMessageName(value string) {
	a.MessageName = (*Max35Text)(&value)
}

// References a related message or provides another reference, such as a pool reference, linking a set of messages. The party which issued the related reference may be the Sender of the referenced message or a party other than the Sender.
type AdditionalReference8 struct {

	// Message identification of a message. This reference was assigned by the party issuing the message.
	Reference *Max35Text `xml:"Ref"`

	// Issuer of the reference.
	ReferenceIssuer *PartyIdentification113 `xml:"RefIssr,omitempty"`

	// Name of the message.
	MessageName *Max35Text `xml:"MsgNm,omitempty"`
}

func (a *AdditionalReference8) SetReference(value string) {
	a.Reference = (*Max35Text)(&value)
}

func (a *AdditionalReference8) AddReferenceIssuer() *PartyIdentification113 {
	a.ReferenceIssuer = new(PartyIdentification113)
	return a.ReferenceIssuer
}

func (a *AdditionalReference8) SetMessageName(value string) {
	a.MessageName = (*Max35Text)(&value)
}

// References a related message or provides another reference, such as a pool reference, linking a set of messages. The party which issued the related reference may be the Sender of the referenced message or a party other than the Sender.
type AdditionalReference9 struct {

	// Reference identifying a set of messages.
	Reference *Max35Text `xml:"Ref"`

	// Issuer of the reference.
	ReferenceIssuer *PartyIdentification113 `xml:"RefIssr,omitempty"`

	// Name of the message.
	MessageName *Max35Text `xml:"MsgNm,omitempty"`
}

func (a *AdditionalReference9) SetReference(value string) {
	a.Reference = (*Max35Text)(&value)
}

func (a *AdditionalReference9) AddReferenceIssuer() *PartyIdentification113 {
	a.ReferenceIssuer = new(PartyIdentification113)
	return a.ReferenceIssuer
}

func (a *AdditionalReference9) SetMessageName(value string) {
	a.MessageName = (*Max35Text)(&value)
}
