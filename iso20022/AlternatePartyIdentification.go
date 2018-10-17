package iso20022

// Alternate identification for a party using an identification type, a country code and a text field.
type AlternatePartyIdentification2 struct {

	// Specifies the type of alternate identification of the party identified.
	IdentificationType *IdentificationType4Choice `xml:"IdTp"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	Country *CountryCode `xml:"Ctry"`

	// Alternate identification for a party.
	AlternateIdentification *Max35Text `xml:"AltrnId"`
}

func (a *AlternatePartyIdentification2) AddIdentificationType() *IdentificationType4Choice {
	a.IdentificationType = new(IdentificationType4Choice)
	return a.IdentificationType
}

func (a *AlternatePartyIdentification2) SetCountry(value string) {
	a.Country = (*CountryCode)(&value)
}

func (a *AlternatePartyIdentification2) SetAlternateIdentification(value string) {
	a.AlternateIdentification = (*Max35Text)(&value)
}

// Alternate identification for a party using an identification type, a country code and a text field.
type AlternatePartyIdentification4 struct {

	// Specifies the type of alternate identification of the party identified.
	IdentificationType *IdentificationType6Choice `xml:"IdTp"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	Country *CountryCode `xml:"Ctry"`

	// Alternate identification for a party.
	AlternateIdentification *Max35Text `xml:"AltrnId"`
}

func (a *AlternatePartyIdentification4) AddIdentificationType() *IdentificationType6Choice {
	a.IdentificationType = new(IdentificationType6Choice)
	return a.IdentificationType
}

func (a *AlternatePartyIdentification4) SetCountry(value string) {
	a.Country = (*CountryCode)(&value)
}

func (a *AlternatePartyIdentification4) SetAlternateIdentification(value string) {
	a.AlternateIdentification = (*Max35Text)(&value)
}

// Alternate identification for a party using an identification type, a country code and a text field.
type AlternatePartyIdentification5 struct {

	// Specifies the type of alternate identification of the party identified.
	IdentificationType *IdentificationType40Choice `xml:"IdTp"`

	// Nation with its own government, occupying a particular territory.
	Country *CountryCode `xml:"Ctry"`

	// Alternate identification for a party.
	AlternateIdentification *Max35Text `xml:"AltrnId"`
}

func (a *AlternatePartyIdentification5) AddIdentificationType() *IdentificationType40Choice {
	a.IdentificationType = new(IdentificationType40Choice)
	return a.IdentificationType
}

func (a *AlternatePartyIdentification5) SetCountry(value string) {
	a.Country = (*CountryCode)(&value)
}

func (a *AlternatePartyIdentification5) SetAlternateIdentification(value string) {
	a.AlternateIdentification = (*Max35Text)(&value)
}

// Alternate identification for a party using an id type, a country code and an text field.
type AlternatePartyIdentification6 struct {

	// Specifies the type of alternate identification which can be used to give an alternate identification of the party identified.
	TypeOfIdentification *IdentificationType41Choice `xml:"TpOfId"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	Country *CountryCode `xml:"Ctry"`

	// Alternate identification for a party.
	AlternateIdentification *Max35Text `xml:"AltrnId"`
}

func (a *AlternatePartyIdentification6) AddTypeOfIdentification() *IdentificationType41Choice {
	a.TypeOfIdentification = new(IdentificationType41Choice)
	return a.TypeOfIdentification
}

func (a *AlternatePartyIdentification6) SetCountry(value string) {
	a.Country = (*CountryCode)(&value)
}

func (a *AlternatePartyIdentification6) SetAlternateIdentification(value string) {
	a.AlternateIdentification = (*Max35Text)(&value)
}

// Alternate identification for a party using an identification type, a country code and a text field.
type AlternatePartyIdentification7 struct {

	// Specifies the type of alternate identification of the party identified.
	IdentificationType *IdentificationType42Choice `xml:"IdTp"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	Country *CountryCode `xml:"Ctry"`

	// Alternate identification for a party.
	AlternateIdentification *Max35Text `xml:"AltrnId"`
}

func (a *AlternatePartyIdentification7) AddIdentificationType() *IdentificationType42Choice {
	a.IdentificationType = new(IdentificationType42Choice)
	return a.IdentificationType
}

func (a *AlternatePartyIdentification7) SetCountry(value string) {
	a.Country = (*CountryCode)(&value)
}

func (a *AlternatePartyIdentification7) SetAlternateIdentification(value string) {
	a.AlternateIdentification = (*Max35Text)(&value)
}

// Alternate identification for a party using an identification type, a country code and a text field.
type AlternatePartyIdentification8 struct {

	// Specifies the type of alternate identification of the party identified.
	IdentificationType *IdentificationType43Choice `xml:"IdTp"`

	// Nation with its own government, occupying a particular territory.
	Country *CountryCode `xml:"Ctry"`

	// Alternate identification for a party.
	AlternateIdentification *Max35Text `xml:"AltrnId"`
}

func (a *AlternatePartyIdentification8) AddIdentificationType() *IdentificationType43Choice {
	a.IdentificationType = new(IdentificationType43Choice)
	return a.IdentificationType
}

func (a *AlternatePartyIdentification8) SetCountry(value string) {
	a.Country = (*CountryCode)(&value)
}

func (a *AlternatePartyIdentification8) SetAlternateIdentification(value string) {
	a.AlternateIdentification = (*Max35Text)(&value)
}

// Alternate identification for a party using an identification type, a country code and a text field.
type AlternatePartyIdentification9 struct {

	// Specifies the type of alternate identification of the party identified.
	IdentificationType *IdentificationType44Choice `xml:"IdTp"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	Country *CountryCode `xml:"Ctry"`

	// Alternate identification for a party.
	AlternateIdentification *RestrictedFINXMax30Text `xml:"AltrnId"`
}

func (a *AlternatePartyIdentification9) AddIdentificationType() *IdentificationType44Choice {
	a.IdentificationType = new(IdentificationType44Choice)
	return a.IdentificationType
}

func (a *AlternatePartyIdentification9) SetCountry(value string) {
	a.Country = (*CountryCode)(&value)
}

func (a *AlternatePartyIdentification9) SetAlternateIdentification(value string) {
	a.AlternateIdentification = (*RestrictedFINXMax30Text)(&value)
}
