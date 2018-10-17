package iso20022

// Information about the status of a transfer instruction and its reason.
type TransferStatusAndReason struct {

	// Business reference of the transfer instruction.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Instruction status and the reason for the status.
	Status *TransferInstructionStatus `xml:"Sts"`

	// Status of the transfer instruction is pending settlement.
	PendingSettlement *PendingSettlementStatusChoice `xml:"PdgSttlm"`

	// Status of the transfer instruction is unmatched.
	Unmatched *TransferUnmatchedStatus `xml:"Umtchd"`

	// Status is in repair.
	InRepair *InRepairStatus2Choice `xml:"InRpr"`

	// Status of the transfer instructed is rejected.
	Rejected *RejectedStatus3Choice `xml:"Rjctd"`

	// Party that initiates the status.
	StatusInitiator *PartyIdentification1Choice `xml:"StsInitr,omitempty"`
}

func (t *TransferStatusAndReason) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *TransferStatusAndReason) AddStatus() *TransferInstructionStatus {
	t.Status = new(TransferInstructionStatus)
	return t.Status
}

func (t *TransferStatusAndReason) AddPendingSettlement() *PendingSettlementStatusChoice {
	t.PendingSettlement = new(PendingSettlementStatusChoice)
	return t.PendingSettlement
}

func (t *TransferStatusAndReason) AddUnmatched() *TransferUnmatchedStatus {
	t.Unmatched = new(TransferUnmatchedStatus)
	return t.Unmatched
}

func (t *TransferStatusAndReason) AddInRepair() *InRepairStatus2Choice {
	t.InRepair = new(InRepairStatus2Choice)
	return t.InRepair
}

func (t *TransferStatusAndReason) AddRejected() *RejectedStatus3Choice {
	t.Rejected = new(RejectedStatus3Choice)
	return t.Rejected
}

func (t *TransferStatusAndReason) AddStatusInitiator() *PartyIdentification1Choice {
	t.StatusInitiator = new(PartyIdentification1Choice)
	return t.StatusInitiator
}

// Information about the status of a transfer instruction and its reason.
type TransferStatusAndReason2 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identification of a transfer, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Status of the transfer is accepted, sent to next party, matched, already executed, or settled.
	Status *TransferInstructionStatus2 `xml:"Sts"`

	// Status of the transfer is pending settlement.
	PendingSettlement *PendingSettlementStatus2 `xml:"PdgSttlm"`

	// Status of the transfer is unmatched.
	Unmatched *TransferUnmatchedStatus2 `xml:"Umtchd"`

	// Status of the transfer is in repair.
	InRepair *InRepairStatus3 `xml:"InRpr"`

	// Status of the transfer is rejected.
	Rejected *RejectedStatus8Choice `xml:"Rjctd"`

	// Status of the transfer is failed settlement, ie, settlement in the International Central Securities Depository (ICSD) or Central Securities Depository (CSD) failed.
	FailedSettlement *FailedSettlementStatus1 `xml:"FaildSttlm"`

	// Status of the transfer is cancelled.
	Cancelled *CancelledStatus3 `xml:"Canc"`

	// Status of the transfer is reversed.
	Reversed *ReversedStatus1 `xml:"Rvsd"`

	// Status of the transfer is cancellation pending.
	CancellationPending *CancellationPendingStatus1 `xml:"CxlPdg"`

	// Date and time at which the transfer was executed.
	TradeDate *ISODate `xml:"TradDt,omitempty"`

	// Date on which the document, for example, the application form, was sent.
	SendOutDate *ISODate `xml:"SndOutDt,omitempty"`

	// Party that initiates the status.
	StatusInitiator *PartyIdentification2Choice `xml:"StsInitr,omitempty"`
}

func (t *TransferStatusAndReason2) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *TransferStatusAndReason2) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *TransferStatusAndReason2) SetClientReference(value string) {
	t.ClientReference = (*Max35Text)(&value)
}

func (t *TransferStatusAndReason2) SetCancellationReference(value string) {
	t.CancellationReference = (*Max35Text)(&value)
}

func (t *TransferStatusAndReason2) AddStatus() *TransferInstructionStatus2 {
	t.Status = new(TransferInstructionStatus2)
	return t.Status
}

func (t *TransferStatusAndReason2) AddPendingSettlement() *PendingSettlementStatus2 {
	t.PendingSettlement = new(PendingSettlementStatus2)
	return t.PendingSettlement
}

func (t *TransferStatusAndReason2) AddUnmatched() *TransferUnmatchedStatus2 {
	t.Unmatched = new(TransferUnmatchedStatus2)
	return t.Unmatched
}

func (t *TransferStatusAndReason2) AddInRepair() *InRepairStatus3 {
	t.InRepair = new(InRepairStatus3)
	return t.InRepair
}

func (t *TransferStatusAndReason2) AddRejected() *RejectedStatus8Choice {
	t.Rejected = new(RejectedStatus8Choice)
	return t.Rejected
}

func (t *TransferStatusAndReason2) AddFailedSettlement() *FailedSettlementStatus1 {
	t.FailedSettlement = new(FailedSettlementStatus1)
	return t.FailedSettlement
}

func (t *TransferStatusAndReason2) AddCancelled() *CancelledStatus3 {
	t.Cancelled = new(CancelledStatus3)
	return t.Cancelled
}

func (t *TransferStatusAndReason2) AddReversed() *ReversedStatus1 {
	t.Reversed = new(ReversedStatus1)
	return t.Reversed
}

func (t *TransferStatusAndReason2) AddCancellationPending() *CancellationPendingStatus1 {
	t.CancellationPending = new(CancellationPendingStatus1)
	return t.CancellationPending
}

func (t *TransferStatusAndReason2) SetTradeDate(value string) {
	t.TradeDate = (*ISODate)(&value)
}

func (t *TransferStatusAndReason2) SetSendOutDate(value string) {
	t.SendOutDate = (*ISODate)(&value)
}

func (t *TransferStatusAndReason2) AddStatusInitiator() *PartyIdentification2Choice {
	t.StatusInitiator = new(PartyIdentification2Choice)
	return t.StatusInitiator
}

// Information about the status of a transfer instruction and its reason.
type TransferStatusAndReason3 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identification of a transfer, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Status of the transfer instruction.
	TransferStatus *TransferStatus1Choice `xml:"TrfSts"`

	// Date and time at which the transfer was executed.
	TradeDate *ISODate `xml:"TradDt,omitempty"`

	// Date on which the document, for example, the application form, was sent.
	SendOutDate *ISODate `xml:"SndOutDt,omitempty"`

	// Party that initiates the status.
	StatusInitiator *PartyIdentification2Choice `xml:"StsInitr,omitempty"`
}

func (t *TransferStatusAndReason3) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *TransferStatusAndReason3) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *TransferStatusAndReason3) SetClientReference(value string) {
	t.ClientReference = (*Max35Text)(&value)
}

func (t *TransferStatusAndReason3) SetCancellationReference(value string) {
	t.CancellationReference = (*Max35Text)(&value)
}

func (t *TransferStatusAndReason3) AddTransferStatus() *TransferStatus1Choice {
	t.TransferStatus = new(TransferStatus1Choice)
	return t.TransferStatus
}

func (t *TransferStatusAndReason3) SetTradeDate(value string) {
	t.TradeDate = (*ISODate)(&value)
}

func (t *TransferStatusAndReason3) SetSendOutDate(value string) {
	t.SendOutDate = (*ISODate)(&value)
}

func (t *TransferStatusAndReason3) AddStatusInitiator() *PartyIdentification2Choice {
	t.StatusInitiator = new(PartyIdentification2Choice)
	return t.StatusInitiator
}

// Information about the status of a transfer instruction and its reason.
type TransferStatusAndReason4 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identification of a transfer, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *AdditionalReference7 `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Status of the transfer instruction.
	TransferStatus *TransferStatus2Choice `xml:"TrfSts"`

	// Date and time at which the transfer was executed.
	TradeDate *ISODate `xml:"TradDt,omitempty"`

	// Date on which the document, for example, the application form, was sent.
	SendOutDate *ISODate `xml:"SndOutDt,omitempty"`

	// Party that initiates the status.
	StatusInitiator *PartyIdentification70Choice `xml:"StsInitr,omitempty"`
}

func (t *TransferStatusAndReason4) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *TransferStatusAndReason4) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *TransferStatusAndReason4) AddClientReference() *AdditionalReference7 {
	t.ClientReference = new(AdditionalReference7)
	return t.ClientReference
}

func (t *TransferStatusAndReason4) SetCancellationReference(value string) {
	t.CancellationReference = (*Max35Text)(&value)
}

func (t *TransferStatusAndReason4) AddTransferStatus() *TransferStatus2Choice {
	t.TransferStatus = new(TransferStatus2Choice)
	return t.TransferStatus
}

func (t *TransferStatusAndReason4) SetTradeDate(value string) {
	t.TradeDate = (*ISODate)(&value)
}

func (t *TransferStatusAndReason4) SetSendOutDate(value string) {
	t.SendOutDate = (*ISODate)(&value)
}

func (t *TransferStatusAndReason4) AddStatusInitiator() *PartyIdentification70Choice {
	t.StatusInitiator = new(PartyIdentification70Choice)
	return t.StatusInitiator
}
