package iso20022

// Indicates whether the collateral is proprietarily owned or client owned.
type CollateralOwnership1 struct {

	// Indicates that the collateral is owned by the clearing member or not.
	Proprietary *YesNoIndicator `xml:"Prtry"`

	// Indicates that the client owns the collateral.
	ClientName *PartyIdentification33Choice `xml:"ClntNm,omitempty"`
}

func (c *CollateralOwnership1) SetProprietary(value string) {
	c.Proprietary = (*YesNoIndicator)(&value)
}

func (c *CollateralOwnership1) AddClientName() *PartyIdentification33Choice {
	c.ClientName = new(PartyIdentification33Choice)
	return c.ClientName
}

// Indicates whether the collateral is proprietarily owned or client owned.
type CollateralOwnership2 struct {

	// Indicates that the collateral is owned by the clearing member or not.
	Proprietary *YesNoIndicator `xml:"Prtry"`

	// Indicates that the client owns the collateral.
	ClientName *PartyIdentification100Choice `xml:"ClntNm,omitempty"`
}

func (c *CollateralOwnership2) SetProprietary(value string) {
	c.Proprietary = (*YesNoIndicator)(&value)
}

func (c *CollateralOwnership2) AddClientName() *PartyIdentification100Choice {
	c.ClientName = new(PartyIdentification100Choice)
	return c.ClientName
}
