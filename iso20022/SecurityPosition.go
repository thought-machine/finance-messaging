package iso20022

// Identifies the securities for which the meeting is organised.
type SecurityPosition5 struct {

	// Security held in an account on which the balance is calculated.
	Identification *SecurityIdentification3 `xml:"Id"`

	// Amount of securities that are eligible for the vote.
	Position []*EligiblePosition2 `xml:"Pos,omitempty"`
}

func (s *SecurityPosition5) AddIdentification() *SecurityIdentification3 {
	s.Identification = new(SecurityIdentification3)
	return s.Identification
}

func (s *SecurityPosition5) AddPosition() *EligiblePosition2 {
	newValue := new(EligiblePosition2)
	s.Position = append(s.Position, newValue)
	return newValue
}

// Identifies the securities for which the meeting is organised.
type SecurityPosition6 struct {

	// Identification of a security, as assigned under a formal or proprietary identification scheme.
	Identification *SecurityIdentification11 `xml:"Id"`

	// Amount of securities that are eligible for the vote.
	Position []*EligiblePosition3 `xml:"Pos,omitempty"`
}

func (s *SecurityPosition6) AddIdentification() *SecurityIdentification11 {
	s.Identification = new(SecurityIdentification11)
	return s.Identification
}

func (s *SecurityPosition6) AddPosition() *EligiblePosition3 {
	newValue := new(EligiblePosition3)
	s.Position = append(s.Position, newValue)
	return newValue
}

// Identifies the securities for which the meeting is organised.
type SecurityPosition7 struct {

	// Identification of a security, as assigned under a formal or proprietary identification scheme.
	Identification *SecurityIdentification11 `xml:"Id"`

	// Amount of securities that are eligible for the vote.
	Position []*EligiblePosition4 `xml:"Pos,omitempty"`
}

func (s *SecurityPosition7) AddIdentification() *SecurityIdentification11 {
	s.Identification = new(SecurityIdentification11)
	return s.Identification
}

func (s *SecurityPosition7) AddPosition() *EligiblePosition4 {
	newValue := new(EligiblePosition4)
	s.Position = append(s.Position, newValue)
	return newValue
}

// Identifies the securities for which the meeting is organised.
type SecurityPosition8 struct {

	// Identification of the security, as assigned under a formal or proprietary identification scheme.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Amount of securities that are eligible for the vote.
	Position []*EligiblePosition5 `xml:"Pos,omitempty"`
}

func (s *SecurityPosition8) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecurityPosition8) AddPosition() *EligiblePosition5 {
	newValue := new(EligiblePosition5)
	s.Position = append(s.Position, newValue)
	return newValue
}

// Identifies the securities for which the meeting is organised.
type SecurityPosition9 struct {

	// Identification of the security, as assigned under a formal or proprietary identification scheme.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Amount of securities that are eligible for the vote.
	Position []*EligiblePosition6 `xml:"Pos,omitempty"`
}

func (s *SecurityPosition9) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecurityPosition9) AddPosition() *EligiblePosition6 {
	newValue := new(EligiblePosition6)
	s.Position = append(s.Position, newValue)
	return newValue
}
