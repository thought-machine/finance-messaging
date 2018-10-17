package iso20022

// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
type SubAccount1 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification1 `xml:"Id"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Specifies additional properties of the account.
	Characteristic *Max35Text `xml:"Chrtc,omitempty"`
}

func (s *SubAccount1) AddIdentification() *AccountIdentification1 {
	s.Identification = new(AccountIdentification1)
	return s.Identification
}

func (s *SubAccount1) SetName(value string) {
	s.Name = (*Max35Text)(&value)
}

func (s *SubAccount1) SetCharacteristic(value string) {
	s.Characteristic = (*Max35Text)(&value)
}

// The subaccount of the safekeeping account
type SubAccount2 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *Max35Text `xml:"Id"`
}

func (s *SubAccount2) SetIdentification(value string) {
	s.Identification = (*Max35Text)(&value)
}

// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
type SubAccount4 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification26 `xml:"Id"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Specifies additional properties of the account.
	Characteristic *Max35Text `xml:"Chrtc,omitempty"`
}

func (s *SubAccount4) AddIdentification() *AccountIdentification26 {
	s.Identification = new(AccountIdentification26)
	return s.Identification
}

func (s *SubAccount4) SetName(value string) {
	s.Name = (*Max35Text)(&value)
}

func (s *SubAccount4) SetCharacteristic(value string) {
	s.Characteristic = (*Max35Text)(&value)
}

// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
type SubAccount5 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *Max35Text `xml:"Id"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Specifies additional properties of the account.
	Characteristic *Max35Text `xml:"Chrtc,omitempty"`
}

func (s *SubAccount5) SetIdentification(value string) {
	s.Identification = (*Max35Text)(&value)
}

func (s *SubAccount5) SetName(value string) {
	s.Name = (*Max35Text)(&value)
}

func (s *SubAccount5) SetCharacteristic(value string) {
	s.Characteristic = (*Max35Text)(&value)
}

// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
type SubAccount6 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *Max35Text `xml:"Id"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Additional properties of the account.
	Characteristic *Max35Text `xml:"Chrtc,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	AccountDesignation *Max35Text `xml:"AcctDsgnt,omitempty"`
}

func (s *SubAccount6) SetIdentification(value string) {
	s.Identification = (*Max35Text)(&value)
}

func (s *SubAccount6) SetName(value string) {
	s.Name = (*Max35Text)(&value)
}

func (s *SubAccount6) SetCharacteristic(value string) {
	s.Characteristic = (*Max35Text)(&value)
}

func (s *SubAccount6) SetAccountDesignation(value string) {
	s.AccountDesignation = (*Max35Text)(&value)
}
