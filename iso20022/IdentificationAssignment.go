package iso20022

// Set of elements that identify the identification assignment.
type IdentificationAssignment1 struct {

	// Point to point reference, as assigned by the assigner, and sent to the next party in the chain to unambiguously identify the message.
	//
	// Usage: The assigner has to make sure that MessageIdentification is unique per assignee for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the identification assignment was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Party that created the identification assignment.
	Creator *Party7Choice `xml:"Cretr,omitempty"`

	// Party that assigns the identification assignment to another party. This is also the sender of the message.
	Assigner *Party7Choice `xml:"Assgnr"`

	// Party that the identification assignment is assigned to. This is also the receiver of the message.
	Assignee *Party7Choice `xml:"Assgne"`
}

func (i *IdentificationAssignment1) SetMessageIdentification(value string) {
	i.MessageIdentification = (*Max35Text)(&value)
}

func (i *IdentificationAssignment1) SetCreationDateTime(value string) {
	i.CreationDateTime = (*ISODateTime)(&value)
}

func (i *IdentificationAssignment1) AddCreator() *Party7Choice {
	i.Creator = new(Party7Choice)
	return i.Creator
}

func (i *IdentificationAssignment1) AddAssigner() *Party7Choice {
	i.Assigner = new(Party7Choice)
	return i.Assigner
}

func (i *IdentificationAssignment1) AddAssignee() *Party7Choice {
	i.Assignee = new(Party7Choice)
	return i.Assignee
}

// Provides the details of the identification assignment.
type IdentificationAssignment2 struct {

	// Point to point reference, as assigned by the assigner, and sent to the next party in the chain to unambiguously identify the message.
	//
	// Usage: The assigner has to make sure that MessageIdentification is unique per assignee for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the identification assignment was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Party that created the identification assignment.
	Creator *Party12Choice `xml:"Cretr,omitempty"`

	// Identifies the first agent in the identification chain, following the payment initiating party.
	FirstAgent *BranchAndFinancialInstitutionIdentification5 `xml:"FrstAgt,omitempty"`

	// Party that assigns the identification assignment to another party. This is also the sender of the message.
	Assigner *Party12Choice `xml:"Assgnr"`

	// Party that the identification assignment is assigned to. This is also the receiver of the message.
	Assignee *Party12Choice `xml:"Assgne"`
}

func (i *IdentificationAssignment2) SetMessageIdentification(value string) {
	i.MessageIdentification = (*Max35Text)(&value)
}

func (i *IdentificationAssignment2) SetCreationDateTime(value string) {
	i.CreationDateTime = (*ISODateTime)(&value)
}

func (i *IdentificationAssignment2) AddCreator() *Party12Choice {
	i.Creator = new(Party12Choice)
	return i.Creator
}

func (i *IdentificationAssignment2) AddFirstAgent() *BranchAndFinancialInstitutionIdentification5 {
	i.FirstAgent = new(BranchAndFinancialInstitutionIdentification5)
	return i.FirstAgent
}

func (i *IdentificationAssignment2) AddAssigner() *Party12Choice {
	i.Assigner = new(Party12Choice)
	return i.Assigner
}

func (i *IdentificationAssignment2) AddAssignee() *Party12Choice {
	i.Assignee = new(Party12Choice)
	return i.Assignee
}
