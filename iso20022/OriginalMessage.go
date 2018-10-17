package iso20022

// Identification of a message previously sent.
type OriginalMessage1 struct {

	// XML schema-instance namespace, for example "tsin.008.001.01"
	MessageDefinitionIdentifier *Max35Text `xml:"MsgDefIdr"`

	// Message sender specified in the original message.
	//
	From *Party9Choice `xml:"Fr"`

	// Message recipient specified in the original message.
	To *Party9Choice `xml:"To"`

	// Message identification specified in the original message.
	BusinessMessageIdentifier *Max35Text `xml:"BizMsgIdr"`

	// Message creation date and time specified in the original message.
	CreationDate *ISONormalisedDateTime `xml:"CreDt"`

	// Indicates whether the message is a copy, a duplicate or a copy of a duplicate of a previously sent ISO 20022 message.
	CopyDuplicate *CopyDuplicate1Code `xml:"CpyDplct,omitempty"`
}

func (o *OriginalMessage1) SetMessageDefinitionIdentifier(value string) {
	o.MessageDefinitionIdentifier = (*Max35Text)(&value)
}

func (o *OriginalMessage1) AddFrom() *Party9Choice {
	o.From = new(Party9Choice)
	return o.From
}

func (o *OriginalMessage1) AddTo() *Party9Choice {
	o.To = new(Party9Choice)
	return o.To
}

func (o *OriginalMessage1) SetBusinessMessageIdentifier(value string) {
	o.BusinessMessageIdentifier = (*Max35Text)(&value)
}

func (o *OriginalMessage1) SetCreationDate(value string) {
	o.CreationDate = (*ISONormalisedDateTime)(&value)
}

func (o *OriginalMessage1) SetCopyDuplicate(value string) {
	o.CopyDuplicate = (*CopyDuplicate1Code)(&value)
}

// Unique and unambiguous identification of the original message references.
type OriginalMessage2 struct {

	// Original message sender used to identify the message.
	OriginalSender *Party28Choice `xml:"OrgnlSndr,omitempty"`

	// Point to point reference assigned by the original instructing party to unambiguously identify the original group of individual transactions.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Specifies the original message name identifier to which the message refers, such as pacs.003.001.01 or MT103.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Original date and time at which the message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Specifies the identification of original package of instructions, entries or records.
	OriginalPackageIdentification *Max35Text `xml:"OrgnlPackgId,omitempty"`

	// Specifies the identification of original entry, instruction or record within the package.
	OriginalRecordIdentification *Max35Text `xml:"OrgnlRcrdId"`
}

func (o *OriginalMessage2) AddOriginalSender() *Party28Choice {
	o.OriginalSender = new(Party28Choice)
	return o.OriginalSender
}

func (o *OriginalMessage2) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalMessage2) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalMessage2) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalMessage2) SetOriginalPackageIdentification(value string) {
	o.OriginalPackageIdentification = (*Max35Text)(&value)
}

func (o *OriginalMessage2) SetOriginalRecordIdentification(value string) {
	o.OriginalRecordIdentification = (*Max35Text)(&value)
}

// Unique and unambiguous identification of the original message references.
type OriginalMessage3 struct {

	// Original message sender used to identify the message.
	OriginalSender *Party28Choice `xml:"OrgnlSndr,omitempty"`

	// Point to point reference assigned by the original instructing party to unambiguously identify the original group of individual transactions.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Specifies the original message name identifier to which the message refers, such as pacs.003.001.01 or MT103.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Original date and time at which the message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`
}

func (o *OriginalMessage3) AddOriginalSender() *Party28Choice {
	o.OriginalSender = new(Party28Choice)
	return o.OriginalSender
}

func (o *OriginalMessage3) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalMessage3) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalMessage3) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}
