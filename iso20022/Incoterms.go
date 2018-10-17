package iso20022

// Specifies the applicable Incoterm and associated location.
type Incoterms1 struct {

	// Specifies the applicable Incoterm by means of a code.
	Code *Incoterms1Code `xml:"Cd"`

	// Specifies Incoterm not present in code list.
	Other *Max35Text `xml:"Othr"`

	// Location where the Incoterms are actioned.
	Location *Max35Text `xml:"Lctn,omitempty"`
}

func (i *Incoterms1) SetCode(value string) {
	i.Code = (*Incoterms1Code)(&value)
}

func (i *Incoterms1) SetOther(value string) {
	i.Other = (*Max35Text)(&value)
}

func (i *Incoterms1) SetLocation(value string) {
	i.Location = (*Max35Text)(&value)
}

// Specifies the applicable Incoterm and associated location.
type Incoterms2 struct {

	// Specifies the applicable Incoterm by means of a code.
	Code *Incoterms1Code `xml:"Cd"`

	// Specifies Incoterm not present in code list.
	Other *Max35Text `xml:"Othr"`

	// Location where the Incoterms are actioned.
	Location *Max35Text `xml:"Lctn"`
}

func (i *Incoterms2) SetCode(value string) {
	i.Code = (*Incoterms1Code)(&value)
}

func (i *Incoterms2) SetOther(value string) {
	i.Other = (*Max35Text)(&value)
}

func (i *Incoterms2) SetLocation(value string) {
	i.Location = (*Max35Text)(&value)
}

// Specifies the applicable Incoterm and associated location.
type Incoterms3 struct {

	// Specifies the Incoterms.
	IncotermsCode *Incoterms4Choice `xml:"IncotrmsCd"`

	// Location where the Incoterms are actioned. Reference UN/ECE Recommendation 16 - LOCODE - Code for Trade and Transport Locations (www.unece.org/cefact/recommendations).
	Location *Max35Text `xml:"Lctn,omitempty"`
}

func (i *Incoterms3) AddIncotermsCode() *Incoterms4Choice {
	i.IncotermsCode = new(Incoterms4Choice)
	return i.IncotermsCode
}

func (i *Incoterms3) SetLocation(value string) {
	i.Location = (*Max35Text)(&value)
}

// Specifies the applicable Incoterm and associated location.
type Incoterms4 struct {

	// Specifies the Incoterms.
	IncotermsCode *Incoterms4Choice `xml:"IncotrmsCd"`

	// Location where the Incoterms are actioned. Reference UN/ECE Recommendation 16 - LOCODE - Code for Trade and Transport Locations (www.unece.org/cefact/recommendations).
	Location *Max70Text `xml:"Lctn,omitempty"`
}

func (i *Incoterms4) AddIncotermsCode() *Incoterms4Choice {
	i.IncotermsCode = new(Incoterms4Choice)
	return i.IncotermsCode
}

func (i *Incoterms4) SetLocation(value string) {
	i.Location = (*Max70Text)(&value)
}
