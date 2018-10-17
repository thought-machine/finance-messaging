package iso20022

// Unique and unambiguous way to identify an organisation.
type IssuerInformation1 struct {

	// Unique and unambiguous way to identify an organisation.
	Identification *PartyIdentification9Choice `xml:"Id"`

	// Address for the Universal Resource Locator (URL), eg, used over the www (HTTP) service.
	URLAddress *Max256Text `xml:"URLAdr,omitempty"`
}

func (i *IssuerInformation1) AddIdentification() *PartyIdentification9Choice {
	i.Identification = new(PartyIdentification9Choice)
	return i.Identification
}

func (i *IssuerInformation1) SetURLAddress(value string) {
	i.URLAddress = (*Max256Text)(&value)
}

// Unique and unambiguous way to identify an organisation.
type IssuerInformation2 struct {

	// Unique and unambiguous way to identify the organisation.
	Identification *PartyIdentification40Choice `xml:"Id"`

	// Address for the Universal Resource Locator (URL), for example, used over the www (HTTP) service.
	URLAddress *Max256Text `xml:"URLAdr,omitempty"`
}

func (i *IssuerInformation2) AddIdentification() *PartyIdentification40Choice {
	i.Identification = new(PartyIdentification40Choice)
	return i.Identification
}

func (i *IssuerInformation2) SetURLAddress(value string) {
	i.URLAddress = (*Max256Text)(&value)
}
