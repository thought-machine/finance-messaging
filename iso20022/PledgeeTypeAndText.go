package iso20022

// Identification of the entity to which the financial instruments are pledged expressed as a code and a narrative description.
type PledgeeTypeAndText1 struct {

	// Additional information about the entity to which the financial instruments are pledged.
	Identification *Max35Text `xml:"Id,omitempty"`

	// Entity to which the financial instruments are pledged expressed as a code.
	PledgeeType *PledgeeType1Code `xml:"PldgeeTp"`
}

func (p *PledgeeTypeAndText1) SetIdentification(value string) {
	p.Identification = (*Max35Text)(&value)
}

func (p *PledgeeTypeAndText1) SetPledgeeType(value string) {
	p.PledgeeType = (*PledgeeType1Code)(&value)
}

// Identification of the entity to which the financial instruments are pledged expressed as a code and a narrative description.
type PledgeeTypeAndText2 struct {

	// Additional information about the entity to which the financial instruments are pledged.
	Identification *RestrictedFINMax30Text `xml:"Id,omitempty"`

	// Entity to which the financial instruments are pledged expressed as a code.
	PledgeeType *PledgeeType1Code `xml:"PldgeeTp"`
}

func (p *PledgeeTypeAndText2) SetIdentification(value string) {
	p.Identification = (*RestrictedFINMax30Text)(&value)
}

func (p *PledgeeTypeAndText2) SetPledgeeType(value string) {
	p.PledgeeType = (*PledgeeType1Code)(&value)
}
