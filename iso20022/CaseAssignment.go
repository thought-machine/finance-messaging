package iso20022

// Represents the assignment of a case to a party. Assignment is a step in the overall process of managing a case.
type CaseAssignment struct {

	// Identification of an assignment within a case.
	Identification *Max35Text `xml:"Id"`

	// Party that assigns the case to another party. This is also the sender of the message.
	Assigner *AnyBICIdentifier `xml:"Assgnr"`

	// Party that the case is assigned to. This is also the receiver of the message.
	Assignee *AnyBICIdentifier `xml:"Assgne"`

	// Date and time at which the assignment was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`
}

func (c *CaseAssignment) SetIdentification(value string) {
	c.Identification = (*Max35Text)(&value)
}

func (c *CaseAssignment) SetAssigner(value string) {
	c.Assigner = (*AnyBICIdentifier)(&value)
}

func (c *CaseAssignment) SetAssignee(value string) {
	c.Assignee = (*AnyBICIdentifier)(&value)
}

func (c *CaseAssignment) SetCreationDateTime(value string) {
	c.CreationDateTime = (*ISODateTime)(&value)
}

// Set of elements used to represent the assignment of a case to a party.
type CaseAssignment2 struct {

	// Uniquely identifies the case assignment.
	Identification *Max35Text `xml:"Id"`

	// Party who assigns the case.
	// Usage: This is also the sender of the message.
	Assigner *Party7Choice `xml:"Assgnr"`

	// Party to which the case is assigned.
	// Usage: This is also the receiver of the message.
	Assignee *Party7Choice `xml:"Assgne"`

	// Date and time at which the assignment was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`
}

func (c *CaseAssignment2) SetIdentification(value string) {
	c.Identification = (*Max35Text)(&value)
}

func (c *CaseAssignment2) AddAssigner() *Party7Choice {
	c.Assigner = new(Party7Choice)
	return c.Assigner
}

func (c *CaseAssignment2) AddAssignee() *Party7Choice {
	c.Assignee = new(Party7Choice)
	return c.Assignee
}

func (c *CaseAssignment2) SetCreationDateTime(value string) {
	c.CreationDateTime = (*ISODateTime)(&value)
}

// Represents the assignment of a case to a party.
type CaseAssignment3 struct {

	// Uniquely identifies the case assignment.
	Identification *Max35Text `xml:"Id"`

	// Party who assigns the case.
	// Usage: This is also the sender of the message.
	Assigner *Party12Choice `xml:"Assgnr"`

	// Party to which the case is assigned.
	// Usage: This is also the receiver of the message.
	Assignee *Party12Choice `xml:"Assgne"`

	// Date and time at which the assignment was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`
}

func (c *CaseAssignment3) SetIdentification(value string) {
	c.Identification = (*Max35Text)(&value)
}

func (c *CaseAssignment3) AddAssigner() *Party12Choice {
	c.Assigner = new(Party12Choice)
	return c.Assigner
}

func (c *CaseAssignment3) AddAssignee() *Party12Choice {
	c.Assignee = new(Party12Choice)
	return c.Assignee
}

func (c *CaseAssignment3) SetCreationDateTime(value string) {
	c.CreationDateTime = (*ISODateTime)(&value)
}
