package iso20022

// Information about the securities movements.
type CorporateActionSecuritiesMovement1 struct {

	// Date and time of the posting.
	PostingDateTime *DateAndDateTimeChoice `xml:"PstngDtTm,omitempty"`

	// Identification of the financial instrument.
	SecurityIdentification *SecurityIdentification7 `xml:"SctyId"`

	// Quantity of securities that is posted.
	PostingQuantity *UnitOrFaceAmount1Choice `xml:"PstngQty"`

	// Provides information about the account which is debited/credited.
	AccountDetails []*SecuritiesAccount8 `xml:"AcctDtls"`
}

func (c *CorporateActionSecuritiesMovement1) AddPostingDateTime() *DateAndDateTimeChoice {
	c.PostingDateTime = new(DateAndDateTimeChoice)
	return c.PostingDateTime
}

func (c *CorporateActionSecuritiesMovement1) AddSecurityIdentification() *SecurityIdentification7 {
	c.SecurityIdentification = new(SecurityIdentification7)
	return c.SecurityIdentification
}

func (c *CorporateActionSecuritiesMovement1) AddPostingQuantity() *UnitOrFaceAmount1Choice {
	c.PostingQuantity = new(UnitOrFaceAmount1Choice)
	return c.PostingQuantity
}

func (c *CorporateActionSecuritiesMovement1) AddAccountDetails() *SecuritiesAccount8 {
	newValue := new(SecuritiesAccount8)
	c.AccountDetails = append(c.AccountDetails, newValue)
	return newValue
}

// Provides information about the securities movement resulting from the election instruction.
type CorporateActionSecuritiesMovement2 struct {

	// Identification of the financial instrument.
	SecurityIdentification *SecurityIdentification7 `xml:"SctyId"`

	// Date and time of the posting of the movement.
	PostingDateTime *DateAndDateTimeChoice `xml:"PstngDtTm,omitempty"`

	// Posting identification of the securities movement.
	PostingIdentification *Max35Text `xml:"PstngId,omitempty"`

	// Securities quantity posted as a result of the securities movement.
	PostingQuantity *UnitOrFaceAmount1Choice `xml:"PstngQty"`

	// Provides information about the account which is debited/credited as a result of the movement.
	AccountDetails []*SecuritiesAccount9 `xml:"AcctDtls"`
}

func (c *CorporateActionSecuritiesMovement2) AddSecurityIdentification() *SecurityIdentification7 {
	c.SecurityIdentification = new(SecurityIdentification7)
	return c.SecurityIdentification
}

func (c *CorporateActionSecuritiesMovement2) AddPostingDateTime() *DateAndDateTimeChoice {
	c.PostingDateTime = new(DateAndDateTimeChoice)
	return c.PostingDateTime
}

func (c *CorporateActionSecuritiesMovement2) SetPostingIdentification(value string) {
	c.PostingIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionSecuritiesMovement2) AddPostingQuantity() *UnitOrFaceAmount1Choice {
	c.PostingQuantity = new(UnitOrFaceAmount1Choice)
	return c.PostingQuantity
}

func (c *CorporateActionSecuritiesMovement2) AddAccountDetails() *SecuritiesAccount9 {
	newValue := new(SecuritiesAccount9)
	c.AccountDetails = append(c.AccountDetails, newValue)
	return newValue
}
