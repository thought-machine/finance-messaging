package iso20022

// Specifies the elements that identify a proxy appointed to represent a party authorised to vote at a shareholders meeting.
type Proxy1 struct {

	// Specifies the type of proxy.
	ProxyType []*ProxyType2Code `xml:"PrxyTp"`

	// Identifies an authorized proxy which has been assigned by the issuer of the meeting.
	PreassignedProxy *IndividualPerson14 `xml:"PrssgndPrxy,omitempty"`
}

func (p *Proxy1) AddProxyType(value string) {
	p.ProxyType = append(p.ProxyType, (*ProxyType2Code)(&value))
}

func (p *Proxy1) AddPreassignedProxy() *IndividualPerson14 {
	p.PreassignedProxy = new(IndividualPerson14)
	return p.PreassignedProxy
}

// Specifies the elements that identify a proxy appointed to represent a party authorised to vote at a shareholders meeting.
type Proxy2 struct {

	// Specifies the type of proxy.
	ProxyType *ProxyType2Code `xml:"PrxyTp"`

	// Person, other than the Chairman of the meeting, assigned by the security holder as proxy.
	PersonDetails *IndividualPerson13 `xml:"PrsnDtls,omitempty"`

	// Indicates the vote instruction for the resolutions which are announced via the meeting agenda in advance of the meeting.
	VoteInstructionForAgendaResolution *Vote1Choice `xml:"VoteInstrForAgndRsltn,omitempty"`
}

func (p *Proxy2) SetProxyType(value string) {
	p.ProxyType = (*ProxyType2Code)(&value)
}

func (p *Proxy2) AddPersonDetails() *IndividualPerson13 {
	p.PersonDetails = new(IndividualPerson13)
	return p.PersonDetails
}

func (p *Proxy2) AddVoteInstructionForAgendaResolution() *Vote1Choice {
	p.VoteInstructionForAgendaResolution = new(Vote1Choice)
	return p.VoteInstructionForAgendaResolution
}

// Specifies the elements that identify a proxy appointed to represent a party authorised to vote at a shareholders meeting.
type Proxy3 struct {

	// Specifies the type of proxy.
	ProxyType []*ProxyType2Code `xml:"PrxyTp"`

	// Identifies an authorized proxy which has been assigned by the issuer of the meeting.
	PreassignedProxy *IndividualPerson16 `xml:"PrssgndPrxy,omitempty"`
}

func (p *Proxy3) AddProxyType(value string) {
	p.ProxyType = append(p.ProxyType, (*ProxyType2Code)(&value))
}

func (p *Proxy3) AddPreassignedProxy() *IndividualPerson16 {
	p.PreassignedProxy = new(IndividualPerson16)
	return p.PreassignedProxy
}

// Specifies the elements that identify a proxy appointed to represent a party authorised to vote at a shareholders meeting.
type Proxy4 struct {

	// Specifies the type of proxy.
	ProxyType *ProxyType2Code `xml:"PrxyTp"`

	// Person, other than the Chairman of the meeting, assigned by the security holder as proxy.
	PersonDetails *IndividualPerson17 `xml:"PrsnDtls,omitempty"`

	// Indicates the vote instruction for the resolutions which are announced via the meeting agenda in advance of the meeting.
	VoteInstructionForAgendaResolution *Vote2Choice `xml:"VoteInstrForAgndRsltn,omitempty"`
}

func (p *Proxy4) SetProxyType(value string) {
	p.ProxyType = (*ProxyType2Code)(&value)
}

func (p *Proxy4) AddPersonDetails() *IndividualPerson17 {
	p.PersonDetails = new(IndividualPerson17)
	return p.PersonDetails
}

func (p *Proxy4) AddVoteInstructionForAgendaResolution() *Vote2Choice {
	p.VoteInstructionForAgendaResolution = new(Vote2Choice)
	return p.VoteInstructionForAgendaResolution
}

// Specifies the elements that identify a proxy appointed to represent a party authorised to vote at a shareholders meeting.
type Proxy5 struct {

	// Specifies the type of proxy.
	ProxyType []*ProxyType2Code `xml:"PrxyTp"`

	// Identifies an authorized proxy that has been assigned by the issuer of the meeting.
	PreassignedProxy *IndividualPerson25 `xml:"PrssgndPrxy,omitempty"`
}

func (p *Proxy5) AddProxyType(value string) {
	p.ProxyType = append(p.ProxyType, (*ProxyType2Code)(&value))
}

func (p *Proxy5) AddPreassignedProxy() *IndividualPerson25 {
	p.PreassignedProxy = new(IndividualPerson25)
	return p.PreassignedProxy
}

// Specifies the elements that identify a proxy appointed to represent a party authorised to vote at a shareholders meeting.
type Proxy6 struct {

	// Specifies the type of proxy.
	ProxyType *ProxyType2Code `xml:"PrxyTp"`

	// Person, other than the chairman of the meeting, assigned by the security holder as the proxy.
	PersonDetails *IndividualPerson26 `xml:"PrsnDtls,omitempty"`

	// Indicates the vote instruction for the resolutions that are announced via the meeting agenda in advance of the meeting.
	VoteInstructionForAgendaResolution *Vote3Choice `xml:"VoteInstrForAgndRsltn,omitempty"`
}

func (p *Proxy6) SetProxyType(value string) {
	p.ProxyType = (*ProxyType2Code)(&value)
}

func (p *Proxy6) AddPersonDetails() *IndividualPerson26 {
	p.PersonDetails = new(IndividualPerson26)
	return p.PersonDetails
}

func (p *Proxy6) AddVoteInstructionForAgendaResolution() *Vote3Choice {
	p.VoteInstructionForAgendaResolution = new(Vote3Choice)
	return p.VoteInstructionForAgendaResolution
}
