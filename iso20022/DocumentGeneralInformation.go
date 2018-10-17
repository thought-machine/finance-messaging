package iso20022

// General information that unambiguously identifies a document, such as identification number and issue date time.
type DocumentGeneralInformation1 struct {

	// Specifies the type of the document, for example commercial invoice.
	DocumentType *DocumentType4Code `xml:"DocTp"`

	// Unique identifier of the document.
	DocumentNumber *Max35Text `xml:"DocNb"`

	// Specifies the identification sequence number for a specific couple sender/receiver.
	SenderReceiverSequenceIdentification *Max140Text `xml:"SndrRcvrSeqId,omitempty"`

	// Issue date of the document.
	IssueDate *ISODate `xml:"IsseDt"`

	// URL (Uniform Resource Locator) related to the document.
	URL *Max256Text `xml:"URL,omitempty"`
}

func (d *DocumentGeneralInformation1) SetDocumentType(value string) {
	d.DocumentType = (*DocumentType4Code)(&value)
}

func (d *DocumentGeneralInformation1) SetDocumentNumber(value string) {
	d.DocumentNumber = (*Max35Text)(&value)
}

func (d *DocumentGeneralInformation1) SetSenderReceiverSequenceIdentification(value string) {
	d.SenderReceiverSequenceIdentification = (*Max140Text)(&value)
}

func (d *DocumentGeneralInformation1) SetIssueDate(value string) {
	d.IssueDate = (*ISODate)(&value)
}

func (d *DocumentGeneralInformation1) SetURL(value string) {
	d.URL = (*Max256Text)(&value)
}

// General information that unambiguously identifies a document, such as identification number and issue date time.
type DocumentGeneralInformation2 struct {

	// Specifies the type of the document, for example commercial invoice.
	DocumentType *ExternalDocumentType1Code `xml:"DocTp"`

	// Unique identifier of the document.
	DocumentNumber *Max35Text `xml:"DocNb"`

	// Specifies the identification sequence number for a specific couple sender/receiver.
	SenderReceiverSequenceIdentification *Max140Text `xml:"SndrRcvrSeqId,omitempty"`

	// Issue date of the document.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// URL (Uniform Resource Locator) where the document can be found
	URL *Max256Text `xml:"URL,omitempty"`

	// Attached binary file for this document.
	AttachedBinaryFile []*BinaryFile1 `xml:"AttchdBinryFile,omitempty"`
}

func (d *DocumentGeneralInformation2) SetDocumentType(value string) {
	d.DocumentType = (*ExternalDocumentType1Code)(&value)
}

func (d *DocumentGeneralInformation2) SetDocumentNumber(value string) {
	d.DocumentNumber = (*Max35Text)(&value)
}

func (d *DocumentGeneralInformation2) SetSenderReceiverSequenceIdentification(value string) {
	d.SenderReceiverSequenceIdentification = (*Max140Text)(&value)
}

func (d *DocumentGeneralInformation2) SetIssueDate(value string) {
	d.IssueDate = (*ISODate)(&value)
}

func (d *DocumentGeneralInformation2) SetURL(value string) {
	d.URL = (*Max256Text)(&value)
}

func (d *DocumentGeneralInformation2) AddAttachedBinaryFile() *BinaryFile1 {
	newValue := new(BinaryFile1)
	d.AttachedBinaryFile = append(d.AttachedBinaryFile, newValue)
	return newValue
}

// General information that unambiguously identifies a document, such as identification number and issue date time.
type DocumentGeneralInformation3 struct {

	// Specifies the type of the document, for example commercial invoice.
	DocumentType *ExternalDocumentType1Code `xml:"DocTp"`

	// Unique identifier of the document.
	DocumentNumber *Max35Text `xml:"DocNb"`

	// Specifies the identification sequence number for a specific couple sender/receiver.
	SenderReceiverSequenceIdentification *Max140Text `xml:"SndrRcvrSeqId,omitempty"`

	// Issue date of the document.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// URL (Uniform Resource Locator) where the document can be found
	URL *Max256Text `xml:"URL,omitempty"`

	// Full Signature Structure without Signature itself:
	// Hash  + Certificate.
	LinkFileHash *SignatureEnvelopeReference `xml:"LkFileHash,omitempty"`

	// Attached binary file for this document.
	AttachedBinaryFile *BinaryFile1 `xml:"AttchdBinryFile"`
}

func (d *DocumentGeneralInformation3) SetDocumentType(value string) {
	d.DocumentType = (*ExternalDocumentType1Code)(&value)
}

func (d *DocumentGeneralInformation3) SetDocumentNumber(value string) {
	d.DocumentNumber = (*Max35Text)(&value)
}

func (d *DocumentGeneralInformation3) SetSenderReceiverSequenceIdentification(value string) {
	d.SenderReceiverSequenceIdentification = (*Max140Text)(&value)
}

func (d *DocumentGeneralInformation3) SetIssueDate(value string) {
	d.IssueDate = (*ISODate)(&value)
}

func (d *DocumentGeneralInformation3) SetURL(value string) {
	d.URL = (*Max256Text)(&value)
}

func (d *DocumentGeneralInformation3) AddLinkFileHash() *SignatureEnvelopeReference {
	d.LinkFileHash = new(SignatureEnvelopeReference)
	return d.LinkFileHash
}

func (d *DocumentGeneralInformation3) AddAttachedBinaryFile() *BinaryFile1 {
	d.AttachedBinaryFile = new(BinaryFile1)
	return d.AttachedBinaryFile
}
