package iso20022

// Specifies the type of instruction requested by the submitter by means of a code.
type InstructionType1 struct {

	// Specifies whether the baseline has to be pushed to the other party or simply lodged.
	Type *InstructionType1Code `xml:"Tp"`
}

func (i *InstructionType1) SetType(value string) {
	i.Type = (*InstructionType1Code)(&value)
}

// Specifies the type of instruction requested by the submitter by means of a code.
type InstructionType3 struct {

	// Specifies whether the data set has to be matched or pre-matched.
	Type *InstructionType3Code `xml:"Tp"`
}

func (i *InstructionType3) SetType(value string) {
	i.Type = (*InstructionType3Code)(&value)
}
