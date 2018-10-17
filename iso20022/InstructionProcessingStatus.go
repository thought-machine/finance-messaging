package iso20022

// Status advising on the processing of the instruction.
type InstructionProcessingStatus1 struct {

	// Status on the processing of the instructions.
	Status *Status3Code `xml:"Sts"`

	// Additional information about the status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (i *InstructionProcessingStatus1) SetStatus(value string) {
	i.Status = (*Status3Code)(&value)
}

func (i *InstructionProcessingStatus1) SetAdditionalInformation(value string) {
	i.AdditionalInformation = (*Max350Text)(&value)
}

// Status advising on the processing of the instruction.
type InstructionProcessingStatus3 struct {

	// Status on the processing of the instructions.
	Status *Status7Code `xml:"Sts"`

	// Additional information about the status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (i *InstructionProcessingStatus3) SetStatus(value string) {
	i.Status = (*Status7Code)(&value)
}

func (i *InstructionProcessingStatus3) SetAdditionalInformation(value string) {
	i.AdditionalInformation = (*Max350Text)(&value)
}
