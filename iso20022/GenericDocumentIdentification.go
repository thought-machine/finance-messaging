package iso20022

// Generic identification scheme for a document.
type GenericDocumentIdentification1 struct {

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber1Choice `xml:"MsgNb,omitempty"`

	// Identification of the document.
	Identification *Max35Text `xml:"Id"`
}

func (g *GenericDocumentIdentification1) AddMessageNumber() *DocumentNumber1Choice {
	g.MessageNumber = new(DocumentNumber1Choice)
	return g.MessageNumber
}

func (g *GenericDocumentIdentification1) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

// Generic identification scheme for a document.
type GenericDocumentIdentification4 struct {

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber5Choice `xml:"MsgNb,omitempty"`

	// Identification of the document.
	Identification *Max35Text `xml:"Id"`
}

func (g *GenericDocumentIdentification4) AddMessageNumber() *DocumentNumber5Choice {
	g.MessageNumber = new(DocumentNumber5Choice)
	return g.MessageNumber
}

func (g *GenericDocumentIdentification4) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

// Generic identification scheme for a document.
type GenericDocumentIdentification5 struct {

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber6Choice `xml:"MsgNb,omitempty"`

	// Identification of the document.
	Identification *RestrictedFINXMax16Text `xml:"Id"`
}

func (g *GenericDocumentIdentification5) AddMessageNumber() *DocumentNumber6Choice {
	g.MessageNumber = new(DocumentNumber6Choice)
	return g.MessageNumber
}

func (g *GenericDocumentIdentification5) SetIdentification(value string) {
	g.Identification = (*RestrictedFINXMax16Text)(&value)
}

// Generic identification scheme for a document.
type GenericDocumentIdentification6 struct {

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber16Choice `xml:"MsgNb,omitempty"`

	// Identification of the document.
	Identification *RestrictedFINXMax16Text `xml:"Id"`
}

func (g *GenericDocumentIdentification6) AddMessageNumber() *DocumentNumber16Choice {
	g.MessageNumber = new(DocumentNumber16Choice)
	return g.MessageNumber
}

func (g *GenericDocumentIdentification6) SetIdentification(value string) {
	g.Identification = (*RestrictedFINXMax16Text)(&value)
}
