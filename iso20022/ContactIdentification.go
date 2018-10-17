package iso20022

// Information needed to contact a physical person.
type ContactIdentification1 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max35Text `xml:"Nm"`

	// Specifies the terms used to formally address a person.
	NamePrefix *NamePrefix1Code `xml:"NmPrfx,omitempty"`

	// First name of a person.
	GivenName *Max35Text `xml:"GvnNm,omitempty"`

	// Role of the party in the activity.
	Role *Max35Text `xml:"Role,omitempty"`

	// Collection of information that identifies a phone number, as defined by telecom services.
	PhoneNumber *PhoneNumber `xml:"PhneNb,omitempty"`

	// Collection of information that identifies a FAX number, as defined by telecom services.
	FaxNumber *PhoneNumber `xml:"FaxNb,omitempty"`

	// Address for electronic mail (e-mail).
	EmailAddress *Max256Text `xml:"EmailAdr,omitempty"`
}

func (c *ContactIdentification1) SetName(value string) {
	c.Name = (*Max35Text)(&value)
}

func (c *ContactIdentification1) SetNamePrefix(value string) {
	c.NamePrefix = (*NamePrefix1Code)(&value)
}

func (c *ContactIdentification1) SetGivenName(value string) {
	c.GivenName = (*Max35Text)(&value)
}

func (c *ContactIdentification1) SetRole(value string) {
	c.Role = (*Max35Text)(&value)
}

func (c *ContactIdentification1) SetPhoneNumber(value string) {
	c.PhoneNumber = (*PhoneNumber)(&value)
}

func (c *ContactIdentification1) SetFaxNumber(value string) {
	c.FaxNumber = (*PhoneNumber)(&value)
}

func (c *ContactIdentification1) SetEmailAddress(value string) {
	c.EmailAddress = (*Max256Text)(&value)
}

// Information needed to contact a physical person.
type ContactIdentification2 struct {

	// Specifies the terms used to formally address a person.
	NamePrefix *NamePrefix1Code `xml:"NmPrfx,omitempty"`

	// First name of a person.
	GivenName *Max35Text `xml:"GvnNm,omitempty"`

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max35Text `xml:"Nm"`

	// Collection of information that identifies a phone number, as defined by telecom services.
	PhoneNumber *PhoneNumber `xml:"PhneNb,omitempty"`

	// Collection of information that identifies a mobile phone number, as defined by telecom services.
	MobileNumber *PhoneNumber `xml:"MobNb,omitempty"`

	// Collection of information that identifies a FAX number, as defined by telecom services.
	FaxNumber *PhoneNumber `xml:"FaxNb,omitempty"`

	// Address for electronic mail (e-mail).
	EmailAddress *Max256Text `xml:"EmailAdr,omitempty"`
}

func (c *ContactIdentification2) SetNamePrefix(value string) {
	c.NamePrefix = (*NamePrefix1Code)(&value)
}

func (c *ContactIdentification2) SetGivenName(value string) {
	c.GivenName = (*Max35Text)(&value)
}

func (c *ContactIdentification2) SetName(value string) {
	c.Name = (*Max35Text)(&value)
}

func (c *ContactIdentification2) SetPhoneNumber(value string) {
	c.PhoneNumber = (*PhoneNumber)(&value)
}

func (c *ContactIdentification2) SetMobileNumber(value string) {
	c.MobileNumber = (*PhoneNumber)(&value)
}

func (c *ContactIdentification2) SetFaxNumber(value string) {
	c.FaxNumber = (*PhoneNumber)(&value)
}

func (c *ContactIdentification2) SetEmailAddress(value string) {
	c.EmailAddress = (*Max256Text)(&value)
}

// Information needed to contact a physical person.
type ContactIdentification3 struct {

	// Business Identifier Code to identify the financial institution that the contact person belongs to.
	BIC *BICIdentifier `xml:"BIC"`

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max35Text `xml:"Nm"`

	// Specifies the terms used to formally address a person.
	NamePrefix *NamePrefix1Code `xml:"NmPrfx,omitempty"`

	// First name of a person.
	GivenName *Max35Text `xml:"GvnNm,omitempty"`

	// Role of the party in the activity.
	Role *Max35Text `xml:"Role,omitempty"`

	// Collection of information that identifies a phone number, as defined by telecom services.
	PhoneNumber *PhoneNumber `xml:"PhneNb,omitempty"`

	// Collection of information that identifies a FAX number, as defined by telecom services.
	FaxNumber *PhoneNumber `xml:"FaxNb,omitempty"`

	// Address for electronic mail (e-mail).
	EmailAddress *Max256Text `xml:"EmailAdr,omitempty"`
}

func (c *ContactIdentification3) SetBIC(value string) {
	c.BIC = (*BICIdentifier)(&value)
}

func (c *ContactIdentification3) SetName(value string) {
	c.Name = (*Max35Text)(&value)
}

func (c *ContactIdentification3) SetNamePrefix(value string) {
	c.NamePrefix = (*NamePrefix1Code)(&value)
}

func (c *ContactIdentification3) SetGivenName(value string) {
	c.GivenName = (*Max35Text)(&value)
}

func (c *ContactIdentification3) SetRole(value string) {
	c.Role = (*Max35Text)(&value)
}

func (c *ContactIdentification3) SetPhoneNumber(value string) {
	c.PhoneNumber = (*PhoneNumber)(&value)
}

func (c *ContactIdentification3) SetFaxNumber(value string) {
	c.FaxNumber = (*PhoneNumber)(&value)
}

func (c *ContactIdentification3) SetEmailAddress(value string) {
	c.EmailAddress = (*Max256Text)(&value)
}

// Information needed to contact a physical person.
type ContactIdentification4 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm"`

	// Specifies the terms used to formally address a person, eg, Ms.
	NamePrefix *NamePrefix1Code `xml:"NmPrfx,omitempty"`

	// First name of a person.
	GivenName *Max350Text `xml:"GvnNm,omitempty"`

	// Role of the party in the activity.
	Role *Max35Text `xml:"Role,omitempty"`

	// Collection of information that identifies a phone number, as defined by telecom services.
	PhoneNumber *PhoneNumber `xml:"PhneNb,omitempty"`

	// Collection of information that identifies a FAX number, as defined by telecom services.
	FaxNumber *PhoneNumber `xml:"FaxNb,omitempty"`

	// Address for electronic mail (e-mail).
	EmailAddress *Max256Text `xml:"EmailAdr,omitempty"`
}

func (c *ContactIdentification4) SetName(value string) {
	c.Name = (*Max350Text)(&value)
}

func (c *ContactIdentification4) SetNamePrefix(value string) {
	c.NamePrefix = (*NamePrefix1Code)(&value)
}

func (c *ContactIdentification4) SetGivenName(value string) {
	c.GivenName = (*Max350Text)(&value)
}

func (c *ContactIdentification4) SetRole(value string) {
	c.Role = (*Max35Text)(&value)
}

func (c *ContactIdentification4) SetPhoneNumber(value string) {
	c.PhoneNumber = (*PhoneNumber)(&value)
}

func (c *ContactIdentification4) SetFaxNumber(value string) {
	c.FaxNumber = (*PhoneNumber)(&value)
}

func (c *ContactIdentification4) SetEmailAddress(value string) {
	c.EmailAddress = (*Max256Text)(&value)
}
