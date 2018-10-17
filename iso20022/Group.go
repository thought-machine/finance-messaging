package iso20022

// Group of parties with their related security certificate.
type Group1 struct {

	// Specifies the identification of the group.
	GroupIdentification *Max4AlphaNumericText `xml:"GrpId"`

	// Specifies a party and related certificate.
	Party []*PartyAndCertificate2 `xml:"Pty"`
}

func (g *Group1) SetGroupIdentification(value string) {
	g.GroupIdentification = (*Max4AlphaNumericText)(&value)
}

func (g *Group1) AddParty() *PartyAndCertificate2 {
	newValue := new(PartyAndCertificate2)
	g.Party = append(g.Party, newValue)
	return newValue
}

// Group of parties with their related security certificate.
type Group2 struct {

	// Specifies the type of change.
	ModificationCode *Modification1Code `xml:"ModCd,omitempty"`

	// Specifies the identification of the group.
	GroupIdentification *Max4AlphaNumericText `xml:"GrpId"`

	// Specifies a party and related certificate.
	Party []*PartyAndCertificate3 `xml:"Pty"`
}

func (g *Group2) SetModificationCode(value string) {
	g.ModificationCode = (*Modification1Code)(&value)
}

func (g *Group2) SetGroupIdentification(value string) {
	g.GroupIdentification = (*Max4AlphaNumericText)(&value)
}

func (g *Group2) AddParty() *PartyAndCertificate3 {
	newValue := new(PartyAndCertificate3)
	g.Party = append(g.Party, newValue)
	return newValue
}
