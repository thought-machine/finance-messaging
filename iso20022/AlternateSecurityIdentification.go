package iso20022

// Proprietary or domestic identification scheme that uniquely identifies a security.
type AlternateSecurityIdentification1 struct {

	// Unique and unambiguous identifier of a security.
	Identification *Max35Text `xml:"Id"`

	// Country of the proprietary identification scheme.
	DomesticIdentificationSource *CountryCode `xml:"DmstIdSrc"`

	// Entity that issues the proprietary identification.
	ProprietaryIdentificationSource *Max35Text `xml:"PrtryIdSrc"`
}

func (a *AlternateSecurityIdentification1) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AlternateSecurityIdentification1) SetDomesticIdentificationSource(value string) {
	a.DomesticIdentificationSource = (*CountryCode)(&value)
}

func (a *AlternateSecurityIdentification1) SetProprietaryIdentificationSource(value string) {
	a.ProprietaryIdentificationSource = (*Max35Text)(&value)
}

// Proprietary or domestic identification scheme that uniquely identifies a security.
type AlternateSecurityIdentification2 struct {

	// Identifies the type of financial instrument identifier type.
	Type *Max35Text `xml:"Tp"`

	// Unique and unambiguous identifier of a security.
	Identification *Max35Text `xml:"Id"`
}

func (a *AlternateSecurityIdentification2) SetType(value string) {
	a.Type = (*Max35Text)(&value)
}

func (a *AlternateSecurityIdentification2) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

// Proprietary or domestic identification scheme that uniquely identifies a security.
type AlternateSecurityIdentification3 struct {

	// Unique and unambiguous identifier of a security.
	Identification *Max70Text `xml:"Id"`

	// Country of the proprietary identification scheme.
	DomesticIdentificationSource *CountryCode `xml:"DmstIdSrc"`

	// Entity that issues the proprietary identification.
	ProprietaryIdentificationSource *Max35Text `xml:"PrtryIdSrc"`
}

func (a *AlternateSecurityIdentification3) SetIdentification(value string) {
	a.Identification = (*Max70Text)(&value)
}

func (a *AlternateSecurityIdentification3) SetDomesticIdentificationSource(value string) {
	a.DomesticIdentificationSource = (*CountryCode)(&value)
}

func (a *AlternateSecurityIdentification3) SetProprietaryIdentificationSource(value string) {
	a.ProprietaryIdentificationSource = (*Max35Text)(&value)
}

// Proprietary or domestic identification scheme that uniquely identifies a security.
type AlternateSecurityIdentification7 struct {

	// Unique and unambiguous identifier of a security.
	Identification *Max35Text `xml:"Id"`

	// Source of the identification, that is, domestic (national) or proprietary.
	IdentificationSource *IdentificationSource1Choice `xml:"IdSrc"`
}

func (a *AlternateSecurityIdentification7) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AlternateSecurityIdentification7) AddIdentificationSource() *IdentificationSource1Choice {
	a.IdentificationSource = new(IdentificationSource1Choice)
	return a.IdentificationSource
}
