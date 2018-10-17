package iso20022

// Status of a transfer cancellation instruction and the reason for the status.
type CancellationStatusAndReason struct {

	// Status of the transfer cancellation instruction.
	Status *TransferCancellationStatus `xml:"Sts"`

	// Status of transfer cancellation is rejected.
	Rejected *TransferCancellationRejectedStatus1Choice `xml:"Rjctd"`

	// Status of the transfer cancellation is complete. The cancellation instruction has been accepted and processed, the cancellation is complete.
	Complete *TransferCancellationCompleteStatusChoice `xml:"Cmplt"`

	// Party that initiates the status.
	StatusInitiator *PartyIdentification1Choice `xml:"StsInitr,omitempty"`
}

func (c *CancellationStatusAndReason) AddStatus() *TransferCancellationStatus {
	c.Status = new(TransferCancellationStatus)
	return c.Status
}

func (c *CancellationStatusAndReason) AddRejected() *TransferCancellationRejectedStatus1Choice {
	c.Rejected = new(TransferCancellationRejectedStatus1Choice)
	return c.Rejected
}

func (c *CancellationStatusAndReason) AddComplete() *TransferCancellationCompleteStatusChoice {
	c.Complete = new(TransferCancellationCompleteStatusChoice)
	return c.Complete
}

func (c *CancellationStatusAndReason) AddStatusInitiator() *PartyIdentification1Choice {
	c.StatusInitiator = new(PartyIdentification1Choice)
	return c.StatusInitiator
}

// Status of a transfer cancellation instruction and the reason for the status.
type CancellationStatusAndReason2 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identification of a transfer, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Status of the transfer cancellation is accepted or sent to next party.
	Status *TransferCancellationStatus2 `xml:"Sts"`

	// Status of the transfer cancellation is rejected.
	Rejected *TransferCancellationRejectedStatus1 `xml:"Rjctd"`

	// Status of the transfer cancellation is complete. The cancellation instruction has been accepted and processed, the cancellation is complete.
	Complete *TransferCancellationCompleteStatusAndReason1 `xml:"Cmplt"`

	// Status of the transfer cancellation is pending.
	Pending *TransferCancellationPendingStatus1 `xml:"Pdg"`

	// Party that initiates the status.
	StatusInitiator *PartyIdentification2Choice `xml:"StsInitr,omitempty"`
}

func (c *CancellationStatusAndReason2) SetMasterReference(value string) {
	c.MasterReference = (*Max35Text)(&value)
}

func (c *CancellationStatusAndReason2) SetTransferReference(value string) {
	c.TransferReference = (*Max35Text)(&value)
}

func (c *CancellationStatusAndReason2) SetClientReference(value string) {
	c.ClientReference = (*Max35Text)(&value)
}

func (c *CancellationStatusAndReason2) SetCancellationReference(value string) {
	c.CancellationReference = (*Max35Text)(&value)
}

func (c *CancellationStatusAndReason2) AddStatus() *TransferCancellationStatus2 {
	c.Status = new(TransferCancellationStatus2)
	return c.Status
}

func (c *CancellationStatusAndReason2) AddRejected() *TransferCancellationRejectedStatus1 {
	c.Rejected = new(TransferCancellationRejectedStatus1)
	return c.Rejected
}

func (c *CancellationStatusAndReason2) AddComplete() *TransferCancellationCompleteStatusAndReason1 {
	c.Complete = new(TransferCancellationCompleteStatusAndReason1)
	return c.Complete
}

func (c *CancellationStatusAndReason2) AddPending() *TransferCancellationPendingStatus1 {
	c.Pending = new(TransferCancellationPendingStatus1)
	return c.Pending
}

func (c *CancellationStatusAndReason2) AddStatusInitiator() *PartyIdentification2Choice {
	c.StatusInitiator = new(PartyIdentification2Choice)
	return c.StatusInitiator
}

// Status of a transfer cancellation instruction and the reason for the status.
type CancellationStatusAndReason3 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identification of a transfer, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *AdditionalReference7 `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Status of the transfer cancellation.
	Status *Status21Choice `xml:"Sts"`

	// Party that initiates the status.
	StatusInitiator *PartyIdentification70Choice `xml:"StsInitr,omitempty"`
}

func (c *CancellationStatusAndReason3) SetMasterReference(value string) {
	c.MasterReference = (*Max35Text)(&value)
}

func (c *CancellationStatusAndReason3) SetTransferReference(value string) {
	c.TransferReference = (*Max35Text)(&value)
}

func (c *CancellationStatusAndReason3) AddClientReference() *AdditionalReference7 {
	c.ClientReference = new(AdditionalReference7)
	return c.ClientReference
}

func (c *CancellationStatusAndReason3) SetCancellationReference(value string) {
	c.CancellationReference = (*Max35Text)(&value)
}

func (c *CancellationStatusAndReason3) AddStatus() *Status21Choice {
	c.Status = new(Status21Choice)
	return c.Status
}

func (c *CancellationStatusAndReason3) AddStatusInitiator() *PartyIdentification70Choice {
	c.StatusInitiator = new(PartyIdentification70Choice)
	return c.StatusInitiator
}
