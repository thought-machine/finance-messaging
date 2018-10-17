package iso20022

// Unique and unambiguous identifier for a message, as assigned by the Sender. This unique identifier can be used for cross-referencing purposes in subsequent messages.
type MessageIdentification struct {

	// String of characters that uniquely identifies a message.
	Identification *Max35Text `xml:"Id"`
}

func (m *MessageIdentification) SetIdentification(value string) {
	m.Identification = (*Max35Text)(&value)
}

// Identifies a message by a unique identifier and the date and time when the message was created by the sender.
type MessageIdentification1 struct {

	// Identification of the message.
	Identification *Max35Text `xml:"Id"`

	// Date of creation of the message.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`
}

func (m *MessageIdentification1) SetIdentification(value string) {
	m.Identification = (*Max35Text)(&value)
}

func (m *MessageIdentification1) SetCreationDateTime(value string) {
	m.CreationDateTime = (*ISODateTime)(&value)
}

// Set of elements providing the identification of a message.
type MessageIdentification2 struct {

	// Specifies the message name identifier of the message that will be used to provide additional details.
	MessageNameIdentification *Max35Text `xml:"MsgNmId,omitempty"`

	// Specifies the identification of the message that will be used to provide additional details.
	MessageIdentification *Max35Text `xml:"MsgId,omitempty"`
}

func (m *MessageIdentification2) SetMessageNameIdentification(value string) {
	m.MessageNameIdentification = (*Max35Text)(&value)
}

func (m *MessageIdentification2) SetMessageIdentification(value string) {
	m.MessageIdentification = (*Max35Text)(&value)
}

// Set of element to identify a message by a unique identifier and the date and time when the message was created by the sender.
type MessageIdentification4 struct {

	// Point to point reference, as assigned by the assigner, and sent to the next party in the chain to unambiguously identify the message.
	//
	// Usage: The assigner has to make sure that MessageIdentification is unique per assignee for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm,omitempty"`
}

func (m *MessageIdentification4) SetMessageIdentification(value string) {
	m.MessageIdentification = (*Max35Text)(&value)
}

func (m *MessageIdentification4) SetCreationDateTime(value string) {
	m.CreationDateTime = (*ISODateTime)(&value)
}

// Identifies a message by a unique identifier and the date and time when the message was created by the sender.
type MessageIdentification5 struct {

	// Point to point reference, as assigned by the assigner, and sent to the next party in the chain to unambiguously identify the message.
	//
	// Usage: The assigner has to make sure that MessageIdentification is unique per assignee for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm,omitempty"`

	// Identifies the first agent in the identification chain, following the payment initiating party.
	FirstAgent *BranchAndFinancialInstitutionIdentification5 `xml:"FrstAgt,omitempty"`
}

func (m *MessageIdentification5) SetMessageIdentification(value string) {
	m.MessageIdentification = (*Max35Text)(&value)
}

func (m *MessageIdentification5) SetCreationDateTime(value string) {
	m.CreationDateTime = (*ISODateTime)(&value)
}

func (m *MessageIdentification5) AddFirstAgent() *BranchAndFinancialInstitutionIdentification5 {
	m.FirstAgent = new(BranchAndFinancialInstitutionIdentification5)
	return m.FirstAgent
}
