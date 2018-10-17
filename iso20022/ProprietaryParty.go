package iso20022

// Set of elements to identify a proprietary party.
type ProprietaryParty1 struct {

	// Identifies the type of proprietary party reported.
	Type *Max35Text `xml:"Tp"`

	// Proprietary party.
	Party *PartyIdentification8 `xml:"Pty"`
}

func (p *ProprietaryParty1) SetType(value string) {
	p.Type = (*Max35Text)(&value)
}

func (p *ProprietaryParty1) AddParty() *PartyIdentification8 {
	p.Party = new(PartyIdentification8)
	return p.Party
}

// Set of elements used to identify a proprietary party.
type ProprietaryParty2 struct {

	// Specifies the type of proprietary party.
	Type *Max35Text `xml:"Tp"`

	// Proprietary party.
	Party *PartyIdentification32 `xml:"Pty"`
}

func (p *ProprietaryParty2) SetType(value string) {
	p.Type = (*Max35Text)(&value)
}

func (p *ProprietaryParty2) AddParty() *PartyIdentification32 {
	p.Party = new(PartyIdentification32)
	return p.Party
}

// Identifies a proprietary party.
type ProprietaryParty3 struct {

	// Specifies the type of proprietary party.
	Type *Max35Text `xml:"Tp"`

	// Proprietary party.
	Party *PartyIdentification43 `xml:"Pty"`
}

func (p *ProprietaryParty3) SetType(value string) {
	p.Type = (*Max35Text)(&value)
}

func (p *ProprietaryParty3) AddParty() *PartyIdentification43 {
	p.Party = new(PartyIdentification43)
	return p.Party
}
