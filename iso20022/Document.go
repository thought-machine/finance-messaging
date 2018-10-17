package iso20022

// Information about a document.
type Document10 struct {

	// Type of document.
	DocumentType *UndertakingDocumentType2Choice `xml:"DocTp"`

	// Channel through which the document should be presented.
	PresentationChannel *Channel1Choice `xml:"PresntnChanl,omitempty"`

	// Format of the document.
	DocumentFormat *DocumentFormat1Choice `xml:"DocFrmt,omitempty"`

	// Indication whether the document may be a copy of the original document.
	CopyIndicator *YesNoIndicator `xml:"CpyInd,omitempty"`

	// Indication whether the document must be signed.
	SignedIndicator *YesNoIndicator `xml:"SgndInd,omitempty"`
}

func (d *Document10) AddDocumentType() *UndertakingDocumentType2Choice {
	d.DocumentType = new(UndertakingDocumentType2Choice)
	return d.DocumentType
}

func (d *Document10) AddPresentationChannel() *Channel1Choice {
	d.PresentationChannel = new(Channel1Choice)
	return d.PresentationChannel
}

func (d *Document10) AddDocumentFormat() *DocumentFormat1Choice {
	d.DocumentFormat = new(DocumentFormat1Choice)
	return d.DocumentFormat
}

func (d *Document10) SetCopyIndicator(value string) {
	d.CopyIndicator = (*YesNoIndicator)(&value)
}

func (d *Document10) SetSignedIndicator(value string) {
	d.SignedIndicator = (*YesNoIndicator)(&value)
}

// Information about a document.
type Document11 struct {

	// Type of document.
	Type *PresentationDocumentFormat1Choice `xml:"Tp,omitempty"`

	// Wording for document.
	Wording *Max20000Text `xml:"Wrdg,omitempty"`

	// Details related to an electronic presentation.
	ElectronicDetails []*Presentation3 `xml:"ElctrncDtls,omitempty"`
}

func (d *Document11) AddType() *PresentationDocumentFormat1Choice {
	d.Type = new(PresentationDocumentFormat1Choice)
	return d.Type
}

func (d *Document11) SetWording(value string) {
	d.Wording = (*Max20000Text)(&value)
}

func (d *Document11) AddElectronicDetails() *Presentation3 {
	newValue := new(Presentation3)
	d.ElectronicDetails = append(d.ElectronicDetails, newValue)
	return newValue
}

// Information about a document.
type Document8 struct {

	// Type of document.
	Type *PresentationDocumentFormat1Choice `xml:"Tp"`

	// Wording for document.
	Wording *Max20000Text `xml:"Wrdg,omitempty"`

	// Details related to an electronic presentation.
	ElectronicDetails []*Presentation3 `xml:"ElctrncDtls,omitempty"`
}

func (d *Document8) AddType() *PresentationDocumentFormat1Choice {
	d.Type = new(PresentationDocumentFormat1Choice)
	return d.Type
}

func (d *Document8) SetWording(value string) {
	d.Wording = (*Max20000Text)(&value)
}

func (d *Document8) AddElectronicDetails() *Presentation3 {
	newValue := new(Presentation3)
	d.ElectronicDetails = append(d.ElectronicDetails, newValue)
	return newValue
}

// Information about a document.
type Document9 struct {

	// Type of document or template.
	Type *UndertakingDocumentType1Choice `xml:"Tp"`

	// Identification of the document or template.
	Identification *Max35Text `xml:"Id"`

	// Format of the document  or template, such as PDF, XML, XSLT.
	Format *DocumentFormat1Choice `xml:"Frmt,omitempty"`

	// Binary file representing the enclosed document or template, such as a PDF file, image file, XML file, MT message.
	Enclosure *Max2MBBinary `xml:"Nclsr"`

	// Digital signature of the enclosed binary file.
	DigitalSignature *PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`
}

func (d *Document9) AddType() *UndertakingDocumentType1Choice {
	d.Type = new(UndertakingDocumentType1Choice)
	return d.Type
}

func (d *Document9) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *Document9) AddFormat() *DocumentFormat1Choice {
	d.Format = new(DocumentFormat1Choice)
	return d.Format
}

func (d *Document9) SetEnclosure(value string) {
	d.Enclosure = (*Max2MBBinary)(&value)
}

func (d *Document9) AddDigitalSignature() *PartyAndSignature2 {
	d.DigitalSignature = new(PartyAndSignature2)
	return d.DigitalSignature
}
