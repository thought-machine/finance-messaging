package iso20022

// Specifies the type of the documents referred by the creditor.
type CreditorReferenceType1 struct {

	// Coded creditor reference type.
	Code *DocumentType3Code `xml:"Cd"`

	// Creditor reference type not avilable in a coded format.
	Proprietary *Max35Text `xml:"Prtry"`

	// Identification of the issuer of the credit reference type.
	Issuer *Max35Text `xml:"Issr,omitempty"`
}

func (c *CreditorReferenceType1) SetCode(value string) {
	c.Code = (*DocumentType3Code)(&value)
}

func (c *CreditorReferenceType1) SetProprietary(value string) {
	c.Proprietary = (*Max35Text)(&value)
}

func (c *CreditorReferenceType1) SetIssuer(value string) {
	c.Issuer = (*Max35Text)(&value)
}

// Specifies the type of creditor reference.
type CreditorReferenceType2 struct {

	// Coded or proprietary format creditor reference type.
	CodeOrProprietary *CreditorReferenceType1Choice `xml:"CdOrPrtry"`

	// Entity that assigns the credit reference type.
	Issuer *Max35Text `xml:"Issr,omitempty"`
}

func (c *CreditorReferenceType2) AddCodeOrProprietary() *CreditorReferenceType1Choice {
	c.CodeOrProprietary = new(CreditorReferenceType1Choice)
	return c.CodeOrProprietary
}

func (c *CreditorReferenceType2) SetIssuer(value string) {
	c.Issuer = (*Max35Text)(&value)
}
