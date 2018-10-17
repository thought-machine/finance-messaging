package iso20022

// Instruction status and the reason for the status.
type TransferInstructionStatus struct {

	// Status of the transfer instruction.
	Status *TransferStatus1Code `xml:"Sts"`

	// Additional information about the status in textual form.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (t *TransferInstructionStatus) SetStatus(value string) {
	t.Status = (*TransferStatus1Code)(&value)
}

func (t *TransferInstructionStatus) SetAdditionalInformation(value string) {
	t.AdditionalInformation = (*Max350Text)(&value)
}

// Instruction status.
type TransferInstructionStatus2 struct {

	// Status of the transfer is accepted, sent to next party, matched, already executed, or settled.
	Status *TransferStatus2Code `xml:"Sts"`

	// Reason for the status.
	Reason *Max350Text `xml:"Rsn,omitempty"`
}

func (t *TransferInstructionStatus2) SetStatus(value string) {
	t.Status = (*TransferStatus2Code)(&value)
}

func (t *TransferInstructionStatus2) SetReason(value string) {
	t.Reason = (*Max350Text)(&value)
}

// Instruction status.
type TransferInstructionStatus3 struct {

	// Status code.
	Status *TransferStatus3Code `xml:"Sts"`

	// Reason for the status.
	Reason *Max350Text `xml:"Rsn,omitempty"`
}

func (t *TransferInstructionStatus3) SetStatus(value string) {
	t.Status = (*TransferStatus3Code)(&value)
}

func (t *TransferInstructionStatus3) SetReason(value string) {
	t.Reason = (*Max350Text)(&value)
}

// Instruction status.
type TransferInstructionStatus4 struct {

	// Status code.
	Status *TransferStatus4Code `xml:"Sts"`

	// Reason for the status.
	Reason *Max350Text `xml:"Rsn,omitempty"`
}

func (t *TransferInstructionStatus4) SetStatus(value string) {
	t.Status = (*TransferStatus4Code)(&value)
}

func (t *TransferInstructionStatus4) SetReason(value string) {
	t.Reason = (*Max350Text)(&value)
}
