package iso20022

// Entity involved in an activity.
type PartyAndSignature1 struct {

	// Entity involved in an activity.
	Party *PartyIdentification41 `xml:"Pty"`

	// Signature of a party.
	Signature *ProprietaryData3 `xml:"Sgntr"`
}

func (p *PartyAndSignature1) AddParty() *PartyIdentification41 {
	p.Party = new(PartyIdentification41)
	return p.Party
}

func (p *PartyAndSignature1) AddSignature() *ProprietaryData3 {
	p.Signature = new(ProprietaryData3)
	return p.Signature
}

// Entity involved in an activity.
type PartyAndSignature2 struct {

	// Entity involved in an activity.
	Party *PartyIdentification43 `xml:"Pty"`

	// Signature of a party.
	Signature *ProprietaryData3 `xml:"Sgntr"`
}

func (p *PartyAndSignature2) AddParty() *PartyIdentification43 {
	p.Party = new(PartyIdentification43)
	return p.Party
}

func (p *PartyAndSignature2) AddSignature() *ProprietaryData3 {
	p.Signature = new(ProprietaryData3)
	return p.Signature
}
