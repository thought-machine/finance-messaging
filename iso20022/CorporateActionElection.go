package iso20022

// Provides information about an amended election instruction.
type CorporateActionElection1 struct {

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption1FormatChoice `xml:"OptnTp"`

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb"`

	// Quantity of the securities that was instructed in the original election advice.
	OriginalInstructedQuantity *UnitOrFaceAmount1Choice `xml:"OrgnlInstdQty"`

	// Remaining instructed securities quantity after the amendment of the election.
	RemainingQuantity *UnitOrFaceAmount1Choice `xml:"RmngQty"`
}

func (c *CorporateActionElection1) AddOptionType() *CorporateActionOption1FormatChoice {
	c.OptionType = new(CorporateActionOption1FormatChoice)
	return c.OptionType
}

func (c *CorporateActionElection1) SetOptionNumber(value string) {
	c.OptionNumber = (*Exact3NumericText)(&value)
}

func (c *CorporateActionElection1) AddOriginalInstructedQuantity() *UnitOrFaceAmount1Choice {
	c.OriginalInstructedQuantity = new(UnitOrFaceAmount1Choice)
	return c.OriginalInstructedQuantity
}

func (c *CorporateActionElection1) AddRemainingQuantity() *UnitOrFaceAmount1Choice {
	c.RemainingQuantity = new(UnitOrFaceAmount1Choice)
	return c.RemainingQuantity
}

// Provides information about a corporate action election amendment request.
type CorporateActionElection2 struct {

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption1FormatChoice `xml:"OptnTp"`

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb"`

	// New instructed securities quantity after the amendment.
	NewInstructedQuantity *UnitOrFaceAmount1Choice `xml:"NewInstdQty"`

	// The reason for the amendment request.
	Reason *Max350Text `xml:"Rsn,omitempty"`
}

func (c *CorporateActionElection2) AddOptionType() *CorporateActionOption1FormatChoice {
	c.OptionType = new(CorporateActionOption1FormatChoice)
	return c.OptionType
}

func (c *CorporateActionElection2) SetOptionNumber(value string) {
	c.OptionNumber = (*Exact3NumericText)(&value)
}

func (c *CorporateActionElection2) AddNewInstructedQuantity() *UnitOrFaceAmount1Choice {
	c.NewInstructedQuantity = new(UnitOrFaceAmount1Choice)
	return c.NewInstructedQuantity
}

func (c *CorporateActionElection2) SetReason(value string) {
	c.Reason = (*Max350Text)(&value)
}

// Provides information about a corporate action election.
type CorporateActionElection3 struct {

	// Provides information about the account.
	AccountDetails *SecuritiesAccount7 `xml:"AcctDtls,omitempty"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption1FormatChoice `xml:"OptnTp,omitempty"`

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb,omitempty"`

	// Quantity of underlying securities to which this instruction applies.
	InstructedUnderlyingSecuritiesQuantity *UnitOrFaceAmount1Choice `xml:"InstdUndrlygSctiesQty,omitempty"`

	// Quantity of the benefits that the account owner wants to receive, eg, as a result of dividend reinvestment.
	InstructedSecuritiesQuantityToReceive *UnitOrFaceAmount1Choice `xml:"InstdSctiesQtyToRcv,omitempty"`

	// Rate proposed in a remarketing of variable rate notes.
	ProposedRate *PercentageRate `xml:"PropsdRate,omitempty"`

	// Provides information about the cash movement resulting from the election instruction.
	CashMovementDetails []*CorporateActionCashMovements2 `xml:"CshMvmntDtls,omitempty"`

	// Provides information about the securities movement resulting from the election instruction.
	SecuritiesMovementDetails []*CorporateActionSecuritiesMovement2 `xml:"SctiesMvmntDtls,omitempty"`
}

func (c *CorporateActionElection3) AddAccountDetails() *SecuritiesAccount7 {
	c.AccountDetails = new(SecuritiesAccount7)
	return c.AccountDetails
}

func (c *CorporateActionElection3) AddOptionType() *CorporateActionOption1FormatChoice {
	c.OptionType = new(CorporateActionOption1FormatChoice)
	return c.OptionType
}

func (c *CorporateActionElection3) SetOptionNumber(value string) {
	c.OptionNumber = (*Exact3NumericText)(&value)
}

func (c *CorporateActionElection3) AddInstructedUnderlyingSecuritiesQuantity() *UnitOrFaceAmount1Choice {
	c.InstructedUnderlyingSecuritiesQuantity = new(UnitOrFaceAmount1Choice)
	return c.InstructedUnderlyingSecuritiesQuantity
}

func (c *CorporateActionElection3) AddInstructedSecuritiesQuantityToReceive() *UnitOrFaceAmount1Choice {
	c.InstructedSecuritiesQuantityToReceive = new(UnitOrFaceAmount1Choice)
	return c.InstructedSecuritiesQuantityToReceive
}

func (c *CorporateActionElection3) SetProposedRate(value string) {
	c.ProposedRate = (*PercentageRate)(&value)
}

func (c *CorporateActionElection3) AddCashMovementDetails() *CorporateActionCashMovements2 {
	newValue := new(CorporateActionCashMovements2)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionElection3) AddSecuritiesMovementDetails() *CorporateActionSecuritiesMovement2 {
	newValue := new(CorporateActionSecuritiesMovement2)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}
