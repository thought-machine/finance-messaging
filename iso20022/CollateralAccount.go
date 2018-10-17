package iso20022

// The Collateral Account provides additional information on the Collateral Account of the Party delivering the collateral.
type CollateralAccount1 struct {

	// Unique identification of the Collateral Account.
	Identification *Max35Text `xml:"Id"`

	// Indicates the Type of Collateral Account.
	Type *CollateralAccountIdentificationType1Choice `xml:"Tp,omitempty"`

	// Description of the account.
	Name *Max70Text `xml:"Nm,omitempty"`
}

func (c *CollateralAccount1) SetIdentification(value string) {
	c.Identification = (*Max35Text)(&value)
}

func (c *CollateralAccount1) AddType() *CollateralAccountIdentificationType1Choice {
	c.Type = new(CollateralAccountIdentificationType1Choice)
	return c.Type
}

func (c *CollateralAccount1) SetName(value string) {
	c.Name = (*Max70Text)(&value)
}

// Provides additional information on the collateral account of the party delivering/receiving the collateral.
type CollateralAccount2 struct {

	// Unique identification of the collateral account.
	Identification *Max35Text `xml:"Id"`

	// Indicates the type of collateral account.
	Type *CollateralAccountIdentificationType2Choice `xml:"Tp,omitempty"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max70Text `xml:"Nm,omitempty"`
}

func (c *CollateralAccount2) SetIdentification(value string) {
	c.Identification = (*Max35Text)(&value)
}

func (c *CollateralAccount2) AddType() *CollateralAccountIdentificationType2Choice {
	c.Type = new(CollateralAccountIdentificationType2Choice)
	return c.Type
}

func (c *CollateralAccount2) SetName(value string) {
	c.Name = (*Max70Text)(&value)
}

// Provides additional information on the collateral account of the party delivering/receiving the collateral.
type CollateralAccount3 struct {

	// Unique identification of the collateral account.
	Identification *Max35Text `xml:"Id"`

	// Indicates the type of collateral account.
	Type *CollateralAccountIdentificationType3Choice `xml:"Tp,omitempty"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max70Text `xml:"Nm,omitempty"`
}

func (c *CollateralAccount3) SetIdentification(value string) {
	c.Identification = (*Max35Text)(&value)
}

func (c *CollateralAccount3) AddType() *CollateralAccountIdentificationType3Choice {
	c.Type = new(CollateralAccountIdentificationType3Choice)
	return c.Type
}

func (c *CollateralAccount3) SetName(value string) {
	c.Name = (*Max70Text)(&value)
}
