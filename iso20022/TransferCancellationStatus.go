package iso20022

// Cancellation status.
type TransferCancellationStatus struct {

	// Status of the transfer cancellation instruction.
	Status *CancellationStatus1Code `xml:"Sts"`

	// Additional information about the status in textual form.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (t *TransferCancellationStatus) SetStatus(value string) {
	t.Status = (*CancellationStatus1Code)(&value)
}

func (t *TransferCancellationStatus) SetAdditionalInformation(value string) {
	t.AdditionalInformation = (*Max350Text)(&value)
}

// Transfer cancellation status is accepted or sent to next party.
type TransferCancellationStatus2 struct {

	// Status of the transfer cancellation is accepted or sent to next party.
	Status *CancellationStatus2Code `xml:"Sts"`

	// Reason for the status.
	Reason *Max350Text `xml:"Rsn,omitempty"`
}

func (t *TransferCancellationStatus2) SetStatus(value string) {
	t.Status = (*CancellationStatus2Code)(&value)
}

func (t *TransferCancellationStatus2) SetReason(value string) {
	t.Reason = (*Max350Text)(&value)
}
