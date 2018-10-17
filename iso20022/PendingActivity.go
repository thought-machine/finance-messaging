package iso20022

// Specifies the event that require an action from one of the parties to the trade transaction.
type PendingActivity1 struct {

	// Code which specifies the next course of action that the receiver of the message must take.
	Type *Action1Code `xml:"Tp"`

	// Further information on the course of action that the receiver of the message must take.
	Description *Max140Text `xml:"Desc,omitempty"`
}

func (p *PendingActivity1) SetType(value string) {
	p.Type = (*Action1Code)(&value)
}

func (p *PendingActivity1) SetDescription(value string) {
	p.Description = (*Max140Text)(&value)
}

// Specifies the event that require an action from one of the parties to the trade transaction.
type PendingActivity2 struct {

	// Code which specifies the next course of action that the receiver of the message must take.
	Type *Action2Code `xml:"Tp"`

	// Further information on the course of action that the receiver of the message must take.
	Description *Max140Text `xml:"Desc,omitempty"`
}

func (p *PendingActivity2) SetType(value string) {
	p.Type = (*Action2Code)(&value)
}

func (p *PendingActivity2) SetDescription(value string) {
	p.Description = (*Max140Text)(&value)
}
