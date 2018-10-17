package iso20022

// Status advising on the rejection of the instruction or cancellation request.
type AdditionalStatus1 struct {

	// Reason advising the rejection of the instruction.
	Reason *InstructionRejectionStatus1Choice `xml:"Rsn"`

	// Additional information about the reason.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (a *AdditionalStatus1) AddReason() *InstructionRejectionStatus1Choice {
	a.Reason = new(InstructionRejectionStatus1Choice)
	return a.Reason
}

func (a *AdditionalStatus1) SetAdditionalInformation(value string) {
	a.AdditionalInformation = (*Max350Text)(&value)
}

// Status advising on the rejection of the instruction or cancellation request.
type AdditionalStatus2 struct {

	// Reason advising the rejection of the instruction cancellation request.
	Reason *CancellationRejectionStatus1Choice `xml:"Rsn"`

	// Additional information about the reason.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (a *AdditionalStatus2) AddReason() *CancellationRejectionStatus1Choice {
	a.Reason = new(CancellationRejectionStatus1Choice)
	return a.Reason
}

func (a *AdditionalStatus2) SetAdditionalInformation(value string) {
	a.AdditionalInformation = (*Max350Text)(&value)
}
