package iso20022

// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
type Traceability1 struct {

	// Identification of a partner of a message exchange.
	RelayIdentification *GenericIdentification31 `xml:"RlayId"`

	// Date and time of incoming data exchange for relaying or processing.
	TraceDateTimeIn *ISODateTime `xml:"TracDtTmIn"`

	// Date and time of the outgoing exchange for relaying or processing.
	TraceDateTimeOut *ISODateTime `xml:"TracDtTmOut"`
}

func (t *Traceability1) AddRelayIdentification() *GenericIdentification31 {
	t.RelayIdentification = new(GenericIdentification31)
	return t.RelayIdentification
}

func (t *Traceability1) SetTraceDateTimeIn(value string) {
	t.TraceDateTimeIn = (*ISODateTime)(&value)
}

func (t *Traceability1) SetTraceDateTimeOut(value string) {
	t.TraceDateTimeOut = (*ISODateTime)(&value)
}

// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
type Traceability2 struct {

	// Identification of a partner of a message exchange.
	RelayIdentification *GenericIdentification76 `xml:"RlayId"`

	// Date and time of incoming data exchange for relaying or processing.
	TraceDateTimeIn *ISODateTime `xml:"TracDtTmIn"`

	// Date and time of the outgoing exchange for relaying or processing.
	TraceDateTimeOut *ISODateTime `xml:"TracDtTmOut"`
}

func (t *Traceability2) AddRelayIdentification() *GenericIdentification76 {
	t.RelayIdentification = new(GenericIdentification76)
	return t.RelayIdentification
}

func (t *Traceability2) SetTraceDateTimeIn(value string) {
	t.TraceDateTimeIn = (*ISODateTime)(&value)
}

func (t *Traceability2) SetTraceDateTimeOut(value string) {
	t.TraceDateTimeOut = (*ISODateTime)(&value)
}

// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
type Traceability3 struct {

	// Identification of a partner of a message exchange.
	RelayIdentification *GenericIdentification74 `xml:"RlayId"`

	// Date and time of incoming data exchange for relaying or processing.
	TraceDateTimeIn *ISODateTime `xml:"TracDtTmIn"`

	// Date and time of the outgoing exchange for relaying or processing.
	TraceDateTimeOut *ISODateTime `xml:"TracDtTmOut"`
}

func (t *Traceability3) AddRelayIdentification() *GenericIdentification74 {
	t.RelayIdentification = new(GenericIdentification74)
	return t.RelayIdentification
}

func (t *Traceability3) SetTraceDateTimeIn(value string) {
	t.TraceDateTimeIn = (*ISODateTime)(&value)
}

func (t *Traceability3) SetTraceDateTimeOut(value string) {
	t.TraceDateTimeOut = (*ISODateTime)(&value)
}

// Identification of partners involved in exchange from the ATM to the issuer, with the relative timestamp of their exchanges.
type Traceability4 struct {

	// Identification of a partner of a message exchange.
	RelayIdentification *GenericIdentification77 `xml:"RlayId"`

	// Identification of the relay node in the path, to enable identification of several hosts in parallel.
	SequenceNumber *Max35Text `xml:"SeqNb,omitempty"`

	// Date and time of incoming data exchange for relaying or processing.
	TraceDateTimeIn *ISODateTime `xml:"TracDtTmIn"`

	// Date and time of the outgoing exchange for relaying or processing.
	TraceDateTimeOut *ISODateTime `xml:"TracDtTmOut"`
}

func (t *Traceability4) AddRelayIdentification() *GenericIdentification77 {
	t.RelayIdentification = new(GenericIdentification77)
	return t.RelayIdentification
}

func (t *Traceability4) SetSequenceNumber(value string) {
	t.SequenceNumber = (*Max35Text)(&value)
}

func (t *Traceability4) SetTraceDateTimeIn(value string) {
	t.TraceDateTimeIn = (*ISODateTime)(&value)
}

func (t *Traceability4) SetTraceDateTimeOut(value string) {
	t.TraceDateTimeOut = (*ISODateTime)(&value)
}

// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
type Traceability5 struct {

	// Identification of a partner of a message exchange.
	RelayIdentification *GenericIdentification76 `xml:"RlayId"`

	// Name of the outgoing protocol used by the node.
	ProtocolName *Max35Text `xml:"PrtcolNm,omitempty"`

	// Version of the protocol.
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn,omitempty"`

	// Date and time of incoming data exchange for relaying or processing.
	TraceDateTimeIn *ISODateTime `xml:"TracDtTmIn"`

	// Date and time of the outgoing exchange for relaying or processing.
	TraceDateTimeOut *ISODateTime `xml:"TracDtTmOut"`
}

func (t *Traceability5) AddRelayIdentification() *GenericIdentification76 {
	t.RelayIdentification = new(GenericIdentification76)
	return t.RelayIdentification
}

func (t *Traceability5) SetProtocolName(value string) {
	t.ProtocolName = (*Max35Text)(&value)
}

func (t *Traceability5) SetProtocolVersion(value string) {
	t.ProtocolVersion = (*Max6Text)(&value)
}

func (t *Traceability5) SetTraceDateTimeIn(value string) {
	t.TraceDateTimeIn = (*ISODateTime)(&value)
}

func (t *Traceability5) SetTraceDateTimeOut(value string) {
	t.TraceDateTimeOut = (*ISODateTime)(&value)
}
