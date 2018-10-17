package iso20022

// Identifies a document by a unique identification and a version together with the identification of the submitter of the document.
type DocumentIdentification1 struct {

	// Identification of a set of data.
	Identification *Max35Text `xml:"Id"`

	// Unambiguous identification of the version of a set of data. Example: Version 1.
	Version *Number `xml:"Vrsn"`

	// Uniquely identifies the financial institution which has submitted the set of data by using a BIC.
	Submitter *BICIdentification1 `xml:"Submitr"`
}

func (d *DocumentIdentification1) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *DocumentIdentification1) SetVersion(value string) {
	d.Version = (*Number)(&value)
}

func (d *DocumentIdentification1) AddSubmitter() *BICIdentification1 {
	d.Submitter = new(BICIdentification1)
	return d.Submitter
}

// Identifies a document by a unique identification and a version together with the submitter of the document.
// Also specifies the type of document and an index for easy referencing.
type DocumentIdentification10 struct {

	// Identification of a set of data.
	Identification *Max35Text `xml:"Id"`

	// Unambiguous identification of the version of a set of data. Example: Version 1.
	Version *Number `xml:"Vrsn"`

	// Identifies the type of data set.
	Type *DataSetType2Code `xml:"Tp"`

	// Uniquely identifies the financial institution which has submitted the set of data by using a BIC.
	Submitter *BICIdentification1 `xml:"Submitr"`

	// Index assigned to the document, to allow easy referencing.
	DocumentIndex *Max3NumericText `xml:"DocIndx"`
}

func (d *DocumentIdentification10) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *DocumentIdentification10) SetVersion(value string) {
	d.Version = (*Number)(&value)
}

func (d *DocumentIdentification10) SetType(value string) {
	d.Type = (*DataSetType2Code)(&value)
}

func (d *DocumentIdentification10) AddSubmitter() *BICIdentification1 {
	d.Submitter = new(BICIdentification1)
	return d.Submitter
}

func (d *DocumentIdentification10) SetDocumentIndex(value string) {
	d.DocumentIndex = (*Max3NumericText)(&value)
}

// Identification and creation date of a document.
type DocumentIdentification11 struct {

	// Unique identifier of the document (message) assigned by the sender of the document.
	Identification *Max35Text `xml:"Id"`

	// Date and time at which the document (message) was created by the sender.
	CreationDateTime *DateAndDateTimeChoice `xml:"CreDtTm,omitempty"`

	// Specifies if this document is a copy, a duplicate, or a duplicate of a copy.
	CopyDuplicate *CopyDuplicate1Code `xml:"CpyDplct,omitempty"`
}

func (d *DocumentIdentification11) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *DocumentIdentification11) AddCreationDateTime() *DateAndDateTimeChoice {
	d.CreationDateTime = new(DateAndDateTimeChoice)
	return d.CreationDateTime
}

func (d *DocumentIdentification11) SetCopyDuplicate(value string) {
	d.CopyDuplicate = (*CopyDuplicate1Code)(&value)
}

// Provides information about identification of the document.
type DocumentIdentification12 struct {

	// Unique identifier of the document (message) assigned by the sender of the document.
	Identification *Max35Text `xml:"Id"`

	// Date and time at which the document (message) was created by the sender.
	CreationDateTime *DateAndDateTimeChoice `xml:"CreDtTm,omitempty"`

	// Specifies if this document is a copy, a duplicate, or a duplicate of a copy.
	CopyDuplicate *CopyDuplicate1Code `xml:"CpyDplct,omitempty"`

	// When used in a corporate action instruction, indicates that the current instruction is replacing a previous one that was cancelled earlier. When used in a corporate action instruction cancellation request, indicates that cancelled instruction will be replaced by a new corporate action instruction to be sent later.
	ChangeInstructionIndicator *YesNoIndicator `xml:"ChngInstrInd,omitempty"`
}

func (d *DocumentIdentification12) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *DocumentIdentification12) AddCreationDateTime() *DateAndDateTimeChoice {
	d.CreationDateTime = new(DateAndDateTimeChoice)
	return d.CreationDateTime
}

func (d *DocumentIdentification12) SetCopyDuplicate(value string) {
	d.CopyDuplicate = (*CopyDuplicate1Code)(&value)
}

func (d *DocumentIdentification12) SetChangeInstructionIndicator(value string) {
	d.ChangeInstructionIndicator = (*YesNoIndicator)(&value)
}

// Identification of a document as well as the document number and type of link.
type DocumentIdentification13 struct {

	// Unique identifier of the document (message) assigned either by the account servicer or the account owner.
	Identification *DocumentIdentification1Choice `xml:"Id"`

	// Identification of the type of document.
	DocumentNumber *DocumentNumber1Choice `xml:"DocNb,omitempty"`

	// Specifies when this document is to be processed relative to an other referred document.
	LinkageType *ProcessingPosition1Choice `xml:"LkgTp,omitempty"`
}

func (d *DocumentIdentification13) AddIdentification() *DocumentIdentification1Choice {
	d.Identification = new(DocumentIdentification1Choice)
	return d.Identification
}

func (d *DocumentIdentification13) AddDocumentNumber() *DocumentNumber1Choice {
	d.DocumentNumber = new(DocumentNumber1Choice)
	return d.DocumentNumber
}

func (d *DocumentIdentification13) AddLinkageType() *ProcessingPosition1Choice {
	d.LinkageType = new(ProcessingPosition1Choice)
	return d.LinkageType
}

// Identification of a document as well as the document number.
type DocumentIdentification14 struct {

	// Unique identifier of the document (message) assigned either by the account servicer or the account owner.
	Identification *DocumentIdentification1Choice `xml:"Id"`

	// Identification of the type of document.
	DocumentNumber *DocumentNumber1Choice `xml:"DocNb,omitempty"`
}

func (d *DocumentIdentification14) AddIdentification() *DocumentIdentification1Choice {
	d.Identification = new(DocumentIdentification1Choice)
	return d.Identification
}

func (d *DocumentIdentification14) AddDocumentNumber() *DocumentNumber1Choice {
	d.DocumentNumber = new(DocumentNumber1Choice)
	return d.DocumentNumber
}

// Identification of a document and type of link.
type DocumentIdentification15 struct {

	// Identifies the document.
	Identification *Max35Text `xml:"Id"`

	// Specifies when this document is to be processed relative to an other referred document.
	LinkageType *ProcessingPosition1Choice `xml:"LkgTp,omitempty"`
}

func (d *DocumentIdentification15) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *DocumentIdentification15) AddLinkageType() *ProcessingPosition1Choice {
	d.LinkageType = new(ProcessingPosition1Choice)
	return d.LinkageType
}

// Identifies a document by a unique identification.
type DocumentIdentification17 struct {

	// Identifies the document.
	Identification *RestrictedFINXMax16Text `xml:"Id"`
}

func (d *DocumentIdentification17) SetIdentification(value string) {
	d.Identification = (*RestrictedFINXMax16Text)(&value)
}

// Identifies a document by a unique identification and a date of issue.
type DocumentIdentification22 struct {

	// Identifies the document.
	Identification *Max35Text `xml:"Id"`

	// Date of issuance of the document.
	DateOfIssue *ISODate `xml:"DtOfIsse,omitempty"`
}

func (d *DocumentIdentification22) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *DocumentIdentification22) SetDateOfIssue(value string) {
	d.DateOfIssue = (*ISODate)(&value)
}

// Identifies a document by a unique identification and a date of issue.
type DocumentIdentification23 struct {

	// Identifies the document.
	Identification *Max35Text `xml:"Id"`

	// Date of issuance of the document.
	DateOfIssue *ISODate `xml:"DtOfIsse,omitempty"`

	// Identification of buyer order line item.
	OrderLineIdentification *Max35Text `xml:"OrdrLineId,omitempty"`
}

func (d *DocumentIdentification23) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *DocumentIdentification23) SetDateOfIssue(value string) {
	d.DateOfIssue = (*ISODate)(&value)
}

func (d *DocumentIdentification23) SetOrderLineIdentification(value string) {
	d.OrderLineIdentification = (*Max35Text)(&value)
}

// Identification of the message number and the query identification.
type DocumentIdentification24 struct {

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber1Choice `xml:"MsgNb,omitempty"`

	// Reference to the query identification.
	Reference *Identification1 `xml:"Ref"`
}

func (d *DocumentIdentification24) AddMessageNumber() *DocumentNumber1Choice {
	d.MessageNumber = new(DocumentNumber1Choice)
	return d.MessageNumber
}

func (d *DocumentIdentification24) AddReference() *Identification1 {
	d.Reference = new(Identification1)
	return d.Reference
}

// Identifies a document by a unique identification and a date of issue.
type DocumentIdentification28 struct {

	// Identifies the document.
	Identification *Max35Text `xml:"Id,omitempty"`

	// Date of issuance of the document.
	DateOfIssue *ISODate `xml:"DtOfIsse"`
}

func (d *DocumentIdentification28) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *DocumentIdentification28) SetDateOfIssue(value string) {
	d.DateOfIssue = (*ISODate)(&value)
}

// Identifies a document by a unique identification and a date of issue.
type DocumentIdentification29 struct {

	// Identifies the document.
	Identification *Max35Text `xml:"Id"`

	// Date of issuance of the document.
	DateOfIssue *ISODate `xml:"DtOfIsse"`
}

func (d *DocumentIdentification29) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *DocumentIdentification29) SetDateOfIssue(value string) {
	d.DateOfIssue = (*ISODate)(&value)
}

// Identifies a document by a unique identification and a version.
type DocumentIdentification3 struct {

	// Identification of a set of data.
	Identification *Max35Text `xml:"Id"`

	// Unambiguous identification of the version of a set of data. Example: Version 1.
	Version *Number `xml:"Vrsn"`
}

func (d *DocumentIdentification3) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *DocumentIdentification3) SetVersion(value string) {
	d.Version = (*Number)(&value)
}

// Identification of the message number and the query identification.
type DocumentIdentification30 struct {

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber5Choice `xml:"MsgNb,omitempty"`

	// Reference to the query identification.
	Reference *Identification14 `xml:"Ref"`
}

func (d *DocumentIdentification30) AddMessageNumber() *DocumentNumber5Choice {
	d.MessageNumber = new(DocumentNumber5Choice)
	return d.MessageNumber
}

func (d *DocumentIdentification30) AddReference() *Identification14 {
	d.Reference = new(Identification14)
	return d.Reference
}

// Identification of a document and type of link.
type DocumentIdentification31 struct {

	// Identifies the document.
	Identification *Max35Text `xml:"Id"`

	// Specifies when this document is to be processed relative to an other referred document.
	LinkageType *ProcessingPosition7Choice `xml:"LkgTp,omitempty"`
}

func (d *DocumentIdentification31) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *DocumentIdentification31) AddLinkageType() *ProcessingPosition7Choice {
	d.LinkageType = new(ProcessingPosition7Choice)
	return d.LinkageType
}

// Identification of a document as well as the document number and type of link.
type DocumentIdentification32 struct {

	// Unique identifier of the document (message) assigned either by the account servicer or the account owner.
	Identification *DocumentIdentification3Choice `xml:"Id"`

	// Identification of the type of document.
	DocumentNumber *DocumentNumber5Choice `xml:"DocNb,omitempty"`

	// Specifies when this document is to be processed relative to an other referred document.
	LinkageType *ProcessingPosition7Choice `xml:"LkgTp,omitempty"`
}

func (d *DocumentIdentification32) AddIdentification() *DocumentIdentification3Choice {
	d.Identification = new(DocumentIdentification3Choice)
	return d.Identification
}

func (d *DocumentIdentification32) AddDocumentNumber() *DocumentNumber5Choice {
	d.DocumentNumber = new(DocumentNumber5Choice)
	return d.DocumentNumber
}

func (d *DocumentIdentification32) AddLinkageType() *ProcessingPosition7Choice {
	d.LinkageType = new(ProcessingPosition7Choice)
	return d.LinkageType
}

// Identification of a document as well as the document number.
type DocumentIdentification33 struct {

	// Unique identifier of the document (message) assigned either by the account servicer or the account owner.
	Identification *DocumentIdentification3Choice `xml:"Id"`

	// Identification of the type of document.
	DocumentNumber *DocumentNumber5Choice `xml:"DocNb,omitempty"`
}

func (d *DocumentIdentification33) AddIdentification() *DocumentIdentification3Choice {
	d.Identification = new(DocumentIdentification3Choice)
	return d.Identification
}

func (d *DocumentIdentification33) AddDocumentNumber() *DocumentNumber5Choice {
	d.DocumentNumber = new(DocumentNumber5Choice)
	return d.DocumentNumber
}

// Identification of a document as well as the document number.
type DocumentIdentification34 struct {

	// Unique identifier of the document (message) assigned either by the account servicer or the account owner.
	Identification *DocumentIdentification4Choice `xml:"Id"`

	// Identification of the type of document.
	DocumentNumber *DocumentNumber6Choice `xml:"DocNb,omitempty"`
}

func (d *DocumentIdentification34) AddIdentification() *DocumentIdentification4Choice {
	d.Identification = new(DocumentIdentification4Choice)
	return d.Identification
}

func (d *DocumentIdentification34) AddDocumentNumber() *DocumentNumber6Choice {
	d.DocumentNumber = new(DocumentNumber6Choice)
	return d.DocumentNumber
}

// Identification of a document and type of link.
type DocumentIdentification37 struct {

	// Identifies the document.
	Identification *RestrictedFINXMax16Text `xml:"Id"`

	// Specifies when this document is to be processed relative to an other referred document.
	LinkageType *ProcessingPosition10Choice `xml:"LkgTp,omitempty"`
}

func (d *DocumentIdentification37) SetIdentification(value string) {
	d.Identification = (*RestrictedFINXMax16Text)(&value)
}

func (d *DocumentIdentification37) AddLinkageType() *ProcessingPosition10Choice {
	d.LinkageType = new(ProcessingPosition10Choice)
	return d.LinkageType
}

// Identification of a document as well as the document number and type of link.
type DocumentIdentification38 struct {

	// Unique identifier of the document (message) assigned either by the account servicer or the account owner.
	Identification *DocumentIdentification4Choice `xml:"Id"`

	// Identification of the type of document.
	DocumentNumber *DocumentNumber6Choice `xml:"DocNb,omitempty"`

	// Specifies when this document is to be processed relative to an other referred document.
	LinkageType *ProcessingPosition10Choice `xml:"LkgTp,omitempty"`
}

func (d *DocumentIdentification38) AddIdentification() *DocumentIdentification4Choice {
	d.Identification = new(DocumentIdentification4Choice)
	return d.Identification
}

func (d *DocumentIdentification38) AddDocumentNumber() *DocumentNumber6Choice {
	d.DocumentNumber = new(DocumentNumber6Choice)
	return d.DocumentNumber
}

func (d *DocumentIdentification38) AddLinkageType() *ProcessingPosition10Choice {
	d.LinkageType = new(ProcessingPosition10Choice)
	return d.LinkageType
}

// Identifies a document by a unique identification and a version together with the submitter of the document.
// Also specifies an index for easy referencing.
type DocumentIdentification4 struct {

	// Identification of a set of data.
	Identification *Max35Text `xml:"Id"`

	// Unambiguous identification of the version of a set of data. Example: Version 1.0.
	Version *Number `xml:"Vrsn"`

	// Uniquely identifies the financial institution which has submitted the set of data by using a BIC.
	Submitter *BICIdentification1 `xml:"Submitr"`

	// Index assigned to the document, to allow easy referencing.
	DocumentIndex *Max3NumericText `xml:"DocIndx"`
}

func (d *DocumentIdentification4) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *DocumentIdentification4) SetVersion(value string) {
	d.Version = (*Number)(&value)
}

func (d *DocumentIdentification4) AddSubmitter() *BICIdentification1 {
	d.Submitter = new(BICIdentification1)
	return d.Submitter
}

func (d *DocumentIdentification4) SetDocumentIndex(value string) {
	d.DocumentIndex = (*Max3NumericText)(&value)
}

// Identification of the message number and the query identification.
type DocumentIdentification48 struct {

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber6Choice `xml:"MsgNb,omitempty"`

	// Reference to the query identification.
	Reference *Identification16 `xml:"Ref"`
}

func (d *DocumentIdentification48) AddMessageNumber() *DocumentNumber6Choice {
	d.MessageNumber = new(DocumentNumber6Choice)
	return d.MessageNumber
}

func (d *DocumentIdentification48) AddReference() *Identification16 {
	d.Reference = new(Identification16)
	return d.Reference
}

// Identification of a document and type of link.
type DocumentIdentification49 struct {

	// Identifies the document.
	Identification *RestrictedFINXMax16Text `xml:"Id"`

	// Specifies when this document is to be processed relative to an other referred document.
	LinkageType *ProcessingPosition22Choice `xml:"LkgTp,omitempty"`
}

func (d *DocumentIdentification49) SetIdentification(value string) {
	d.Identification = (*RestrictedFINXMax16Text)(&value)
}

func (d *DocumentIdentification49) AddLinkageType() *ProcessingPosition22Choice {
	d.LinkageType = new(ProcessingPosition22Choice)
	return d.LinkageType
}

// Identifies a document by a unique identification and its issuer.
type DocumentIdentification5 struct {

	// Identification of a set of data.
	Identification *Max35Text `xml:"Id"`

	// Uniquely identifies the financial institution which has issued the identification of the set of data by using a BIC.
	IdentificationIssuer *BICIdentification1 `xml:"IdIssr"`
}

func (d *DocumentIdentification5) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *DocumentIdentification5) AddIdentificationIssuer() *BICIdentification1 {
	d.IdentificationIssuer = new(BICIdentification1)
	return d.IdentificationIssuer
}

// Identifies a document by a unique identification and a version.
// Also provides reference to a baseline amendment number.
type DocumentIdentification6 struct {

	// Identification of a set of data.
	Identification *Max35Text `xml:"Id"`

	// Unambiguous identification of the version of a set of data. Example: Version 1.
	Version *Number `xml:"Vrsn"`

	// Number that is assigned sequentially by the TSU to a baseline amendment.
	AmendmentSequenceNumber *Max3NumericText `xml:"AmdmntSeqNb,omitempty"`
}

func (d *DocumentIdentification6) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *DocumentIdentification6) SetVersion(value string) {
	d.Version = (*Number)(&value)
}

func (d *DocumentIdentification6) SetAmendmentSequenceNumber(value string) {
	d.AmendmentSequenceNumber = (*Max3NumericText)(&value)
}

// Identifies a document by a unique identification and a date of issue.
type DocumentIdentification7 struct {

	// Identifies the document.
	Identification *Max35Text `xml:"Id"`

	// Date of issuance of the document.
	DateOfIssue *ISODate `xml:"DtOfIsse"`
}

func (d *DocumentIdentification7) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *DocumentIdentification7) SetDateOfIssue(value string) {
	d.DateOfIssue = (*ISODate)(&value)
}

// Identifies the document by providing a unique identification and optionally the date/time of the creation of the document.
type DocumentIdentification8 struct {

	// Unique identification of the document.
	Identification *Max35Text `xml:"Id"`

	// Date/time of the creation of the document.
	CreationDateTime *ISODateTime `xml:"CreDtTm,omitempty"`
}

func (d *DocumentIdentification8) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *DocumentIdentification8) SetCreationDateTime(value string) {
	d.CreationDateTime = (*ISODateTime)(&value)
}

// Identifies a document by a unique identification.
type DocumentIdentification9 struct {

	// Identifies the document.
	Identification *Max35Text `xml:"Id"`
}

func (d *DocumentIdentification9) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}
