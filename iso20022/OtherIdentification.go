package iso20022

// Other accepted financial instrument's identification than ISIN.
type OtherIdentification1 struct {

	// Identification of a security.
	Identification *Max35Text `xml:"Id"`

	// Identifies the suffix of the security identification.
	Suffix *Max16Text `xml:"Sfx,omitempty"`

	// Type of the identification.
	Type *IdentificationSource3Choice `xml:"Tp"`
}

func (o *OtherIdentification1) SetIdentification(value string) {
	o.Identification = (*Max35Text)(&value)
}

func (o *OtherIdentification1) SetSuffix(value string) {
	o.Suffix = (*Max16Text)(&value)
}

func (o *OtherIdentification1) AddType() *IdentificationSource3Choice {
	o.Type = new(IdentificationSource3Choice)
	return o.Type
}

// Other accepted financial instrument's identification than ISIN.
type OtherIdentification2 struct {

	// Identification of a security.
	Identification *RestrictedFINXMax31Text `xml:"Id"`

	// Identifies the suffix of the security identification.
	Suffix *Max16Text `xml:"Sfx,omitempty"`

	// Type of the identification.
	Type *IdentificationSource4Choice `xml:"Tp"`
}

func (o *OtherIdentification2) SetIdentification(value string) {
	o.Identification = (*RestrictedFINXMax31Text)(&value)
}

func (o *OtherIdentification2) SetSuffix(value string) {
	o.Suffix = (*Max16Text)(&value)
}

func (o *OtherIdentification2) AddType() *IdentificationSource4Choice {
	o.Type = new(IdentificationSource4Choice)
	return o.Type
}

// Other accepted financial instrument's identification than ISIN.
type OtherIdentification3 struct {

	// Identification of a security.
	Identification *RestrictedFINXMax31Text `xml:"Id"`

	// Identifies the suffix of the security identification.
	Suffix *Max16Text `xml:"Sfx,omitempty"`

	// Type of the identification.
	Type *IdentificationSource3Choice `xml:"Tp"`
}

func (o *OtherIdentification3) SetIdentification(value string) {
	o.Identification = (*RestrictedFINXMax31Text)(&value)
}

func (o *OtherIdentification3) SetSuffix(value string) {
	o.Suffix = (*Max16Text)(&value)
}

func (o *OtherIdentification3) AddType() *IdentificationSource3Choice {
	o.Type = new(IdentificationSource3Choice)
	return o.Type
}

// Identification of a financial instrument using an accepted format other than an ISIN.
type OtherIdentification4 struct {

	// Identification of the fund/sub fund.
	Identification *Max35Text `xml:"Id"`

	// Identification source
	Type *IdentificationSource5Choice `xml:"Tp"`
}

func (o *OtherIdentification4) SetIdentification(value string) {
	o.Identification = (*Max35Text)(&value)
}

func (o *OtherIdentification4) AddType() *IdentificationSource5Choice {
	o.Type = new(IdentificationSource5Choice)
	return o.Type
}
