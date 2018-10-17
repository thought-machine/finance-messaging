package iso20022

// Specifies corporate action quantities.
type CorporateActionQuantity1 struct {

	// Minimum quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblQty,omitempty"`

	// Minimum multiple quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableMultipleQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblMltplQty,omitempty"`

	// Maximum number of securities the offeror is requesting to complete the event.
	MaximumQuantity *FinancialInstrumentQuantity2Choice `xml:"MaxQty,omitempty"`

	// Minimum quantity of securities the offeror/issuer will purchase or redeem under the terms of the event. This can be a number or the term "any and all".
	MinimumQuantitySought *FinancialInstrumentQuantity2Choice `xml:"MinQtySght,omitempty"`

	// Quantity of equity that makes up the new board lot.
	NewBoardLotQuantity *FinancialInstrumentQuantity1Choice `xml:"NewBrdLotQty,omitempty"`

	// New denomination of the equity following, for example, an increase or decrease in nominal value.
	NewDenominationQuantity *FinancialInstrumentQuantity1Choice `xml:"NewDnmtnQty,omitempty"`

	// Minimum integral amount of securities that each account owner must have remaining after the called amounts are applied.
	BaseDenomination *FinancialInstrumentQuantity1Choice `xml:"BaseDnmtn,omitempty"`

	// Amount used when the called amount is not met by running the lottery with the base denomination.
	IncrementalDenomination *FinancialInstrumentQuantity1Choice `xml:"IncrmtlDnmtn,omitempty"`
}

func (c *CorporateActionQuantity1) AddMinimumExercisableQuantity() *FinancialInstrumentQuantity1Choice {
	c.MinimumExercisableQuantity = new(FinancialInstrumentQuantity1Choice)
	return c.MinimumExercisableQuantity
}

func (c *CorporateActionQuantity1) AddMinimumExercisableMultipleQuantity() *FinancialInstrumentQuantity1Choice {
	c.MinimumExercisableMultipleQuantity = new(FinancialInstrumentQuantity1Choice)
	return c.MinimumExercisableMultipleQuantity
}

func (c *CorporateActionQuantity1) AddMaximumQuantity() *FinancialInstrumentQuantity2Choice {
	c.MaximumQuantity = new(FinancialInstrumentQuantity2Choice)
	return c.MaximumQuantity
}

func (c *CorporateActionQuantity1) AddMinimumQuantitySought() *FinancialInstrumentQuantity2Choice {
	c.MinimumQuantitySought = new(FinancialInstrumentQuantity2Choice)
	return c.MinimumQuantitySought
}

func (c *CorporateActionQuantity1) AddNewBoardLotQuantity() *FinancialInstrumentQuantity1Choice {
	c.NewBoardLotQuantity = new(FinancialInstrumentQuantity1Choice)
	return c.NewBoardLotQuantity
}

func (c *CorporateActionQuantity1) AddNewDenominationQuantity() *FinancialInstrumentQuantity1Choice {
	c.NewDenominationQuantity = new(FinancialInstrumentQuantity1Choice)
	return c.NewDenominationQuantity
}

func (c *CorporateActionQuantity1) AddBaseDenomination() *FinancialInstrumentQuantity1Choice {
	c.BaseDenomination = new(FinancialInstrumentQuantity1Choice)
	return c.BaseDenomination
}

func (c *CorporateActionQuantity1) AddIncrementalDenomination() *FinancialInstrumentQuantity1Choice {
	c.IncrementalDenomination = new(FinancialInstrumentQuantity1Choice)
	return c.IncrementalDenomination
}

// Specifies corporate action quantities.
type CorporateActionQuantity3 struct {

	// Minimum quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblQty,omitempty"`

	// Minimum multiple quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableMultipleQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblMltplQty,omitempty"`

	// Maximum number of securities the offeror is requesting to complete the event.
	MaximumQuantity *FinancialInstrumentQuantity16Choice `xml:"MaxQty,omitempty"`

	// Minimum quantity of securities the offeror/issuer will purchase or redeem under the terms of the event. This can be a number or the term "any and all".
	MinimumQuantitySought *FinancialInstrumentQuantity16Choice `xml:"MinQtySght,omitempty"`

	// Quantity of equity that makes up the new board lot.
	NewBoardLotQuantity *FinancialInstrumentQuantity1Choice `xml:"NewBrdLotQty,omitempty"`

	// New denomination of the equity following, for example, an increase or decrease in nominal value.
	NewDenominationQuantity *FinancialInstrumentQuantity1Choice `xml:"NewDnmtnQty,omitempty"`

	// Minimum integral amount of securities that each account owner must have remaining after the called amounts are applied.
	BaseDenomination *FinancialInstrumentQuantity1Choice `xml:"BaseDnmtn,omitempty"`

	// Amount used when the called amount is not met by running the lottery with the base denomination.
	IncrementalDenomination *FinancialInstrumentQuantity1Choice `xml:"IncrmtlDnmtn,omitempty"`
}

func (c *CorporateActionQuantity3) AddMinimumExercisableQuantity() *FinancialInstrumentQuantity1Choice {
	c.MinimumExercisableQuantity = new(FinancialInstrumentQuantity1Choice)
	return c.MinimumExercisableQuantity
}

func (c *CorporateActionQuantity3) AddMinimumExercisableMultipleQuantity() *FinancialInstrumentQuantity1Choice {
	c.MinimumExercisableMultipleQuantity = new(FinancialInstrumentQuantity1Choice)
	return c.MinimumExercisableMultipleQuantity
}

func (c *CorporateActionQuantity3) AddMaximumQuantity() *FinancialInstrumentQuantity16Choice {
	c.MaximumQuantity = new(FinancialInstrumentQuantity16Choice)
	return c.MaximumQuantity
}

func (c *CorporateActionQuantity3) AddMinimumQuantitySought() *FinancialInstrumentQuantity16Choice {
	c.MinimumQuantitySought = new(FinancialInstrumentQuantity16Choice)
	return c.MinimumQuantitySought
}

func (c *CorporateActionQuantity3) AddNewBoardLotQuantity() *FinancialInstrumentQuantity1Choice {
	c.NewBoardLotQuantity = new(FinancialInstrumentQuantity1Choice)
	return c.NewBoardLotQuantity
}

func (c *CorporateActionQuantity3) AddNewDenominationQuantity() *FinancialInstrumentQuantity1Choice {
	c.NewDenominationQuantity = new(FinancialInstrumentQuantity1Choice)
	return c.NewDenominationQuantity
}

func (c *CorporateActionQuantity3) AddBaseDenomination() *FinancialInstrumentQuantity1Choice {
	c.BaseDenomination = new(FinancialInstrumentQuantity1Choice)
	return c.BaseDenomination
}

func (c *CorporateActionQuantity3) AddIncrementalDenomination() *FinancialInstrumentQuantity1Choice {
	c.IncrementalDenomination = new(FinancialInstrumentQuantity1Choice)
	return c.IncrementalDenomination
}

// Specifies corporate action quantities.
type CorporateActionQuantity5 struct {

	// Minimum quantity (or lot) of financial instrument that may be exercised or tendered.
	MinimumExercisableQuantity *FinancialInstrumentQuantity19Choice `xml:"MinExrcblQty,omitempty"`

	// Minimum multiple quantity (or lot) of financial instrument that  may be exercised or tendered.
	MinimumExercisableMultipleQuantity *FinancialInstrumentQuantity20Choice `xml:"MinExrcblMltplQty,omitempty"`

	// The maximum number of securities the offeror/issuer is ready to purchase or redeem. This can be a number or the term "any and all".
	MaximumQuantity *FinancialInstrumentQuantity19Choice `xml:"MaxQty,omitempty"`

	// Minimum quantity of securities the offeror/issuer is ready to purchase or redeem under the terms of the event. This can be a number or the term "any and all".
	MinimumQuantitySought *FinancialInstrumentQuantity19Choice `xml:"MinQtySght,omitempty"`

	// Quantity of equity that makes up the new board lot.
	NewBoardLotQuantity *FinancialInstrumentQuantity20Choice `xml:"NewBrdLotQty,omitempty"`

	// New denomination of the equity following, for example, an increase or decrease in nominal value.
	NewDenominationQuantity *FinancialInstrumentQuantity20Choice `xml:"NewDnmtnQty,omitempty"`

	// Minimum integral amount of securities that each account owner must have remaining after the called amounts are applied.
	BaseDenomination *FinancialInstrumentQuantity20Choice `xml:"BaseDnmtn,omitempty"`

	// Amount used when the called amount is not met by running the lottery with the base denomination.
	IncrementalDenomination *FinancialInstrumentQuantity20Choice `xml:"IncrmtlDnmtn,omitempty"`
}

func (c *CorporateActionQuantity5) AddMinimumExercisableQuantity() *FinancialInstrumentQuantity19Choice {
	c.MinimumExercisableQuantity = new(FinancialInstrumentQuantity19Choice)
	return c.MinimumExercisableQuantity
}

func (c *CorporateActionQuantity5) AddMinimumExercisableMultipleQuantity() *FinancialInstrumentQuantity20Choice {
	c.MinimumExercisableMultipleQuantity = new(FinancialInstrumentQuantity20Choice)
	return c.MinimumExercisableMultipleQuantity
}

func (c *CorporateActionQuantity5) AddMaximumQuantity() *FinancialInstrumentQuantity19Choice {
	c.MaximumQuantity = new(FinancialInstrumentQuantity19Choice)
	return c.MaximumQuantity
}

func (c *CorporateActionQuantity5) AddMinimumQuantitySought() *FinancialInstrumentQuantity19Choice {
	c.MinimumQuantitySought = new(FinancialInstrumentQuantity19Choice)
	return c.MinimumQuantitySought
}

func (c *CorporateActionQuantity5) AddNewBoardLotQuantity() *FinancialInstrumentQuantity20Choice {
	c.NewBoardLotQuantity = new(FinancialInstrumentQuantity20Choice)
	return c.NewBoardLotQuantity
}

func (c *CorporateActionQuantity5) AddNewDenominationQuantity() *FinancialInstrumentQuantity20Choice {
	c.NewDenominationQuantity = new(FinancialInstrumentQuantity20Choice)
	return c.NewDenominationQuantity
}

func (c *CorporateActionQuantity5) AddBaseDenomination() *FinancialInstrumentQuantity20Choice {
	c.BaseDenomination = new(FinancialInstrumentQuantity20Choice)
	return c.BaseDenomination
}

func (c *CorporateActionQuantity5) AddIncrementalDenomination() *FinancialInstrumentQuantity20Choice {
	c.IncrementalDenomination = new(FinancialInstrumentQuantity20Choice)
	return c.IncrementalDenomination
}

// Specifies corporate action quantities.
type CorporateActionQuantity7 struct {

	// The maximum number of securities the offeror/issuer is ready to purchase or redeem. This can be a number or the term "any and all".
	MaximumQuantity *FinancialInstrumentQuantity19Choice `xml:"MaxQty,omitempty"`

	// Minimum quantity of securities the offeror/issuer is ready to purchase or redeem under the terms of the event. This can be a number or the term "any and all".
	MinimumQuantitySought *FinancialInstrumentQuantity19Choice `xml:"MinQtySght,omitempty"`

	// Quantity of equity that makes up the new board lot.
	NewBoardLotQuantity *FinancialInstrumentQuantity20Choice `xml:"NewBrdLotQty,omitempty"`

	// New denomination of the equity following, for example, an increase or decrease in nominal value.
	NewDenominationQuantity *FinancialInstrumentQuantity20Choice `xml:"NewDnmtnQty,omitempty"`

	// Minimum integral amount of securities that each account owner must have remaining after the called amounts are applied.
	BaseDenomination *FinancialInstrumentQuantity20Choice `xml:"BaseDnmtn,omitempty"`

	// Amount used when the called amount is not met by running the lottery with the base denomination.
	IncrementalDenomination *FinancialInstrumentQuantity20Choice `xml:"IncrmtlDnmtn,omitempty"`
}

func (c *CorporateActionQuantity7) AddMaximumQuantity() *FinancialInstrumentQuantity19Choice {
	c.MaximumQuantity = new(FinancialInstrumentQuantity19Choice)
	return c.MaximumQuantity
}

func (c *CorporateActionQuantity7) AddMinimumQuantitySought() *FinancialInstrumentQuantity19Choice {
	c.MinimumQuantitySought = new(FinancialInstrumentQuantity19Choice)
	return c.MinimumQuantitySought
}

func (c *CorporateActionQuantity7) AddNewBoardLotQuantity() *FinancialInstrumentQuantity20Choice {
	c.NewBoardLotQuantity = new(FinancialInstrumentQuantity20Choice)
	return c.NewBoardLotQuantity
}

func (c *CorporateActionQuantity7) AddNewDenominationQuantity() *FinancialInstrumentQuantity20Choice {
	c.NewDenominationQuantity = new(FinancialInstrumentQuantity20Choice)
	return c.NewDenominationQuantity
}

func (c *CorporateActionQuantity7) AddBaseDenomination() *FinancialInstrumentQuantity20Choice {
	c.BaseDenomination = new(FinancialInstrumentQuantity20Choice)
	return c.BaseDenomination
}

func (c *CorporateActionQuantity7) AddIncrementalDenomination() *FinancialInstrumentQuantity20Choice {
	c.IncrementalDenomination = new(FinancialInstrumentQuantity20Choice)
	return c.IncrementalDenomination
}

// Specifies corporate action quantities.
type CorporateActionQuantity8 struct {

	// The maximum number of securities the offeror/issuer is ready to purchase or redeem. This can be a number or the term "any and all".
	MaximumQuantity *FinancialInstrumentQuantity21Choice `xml:"MaxQty,omitempty"`

	// Minimum quantity of securities the offeror/issuer is ready to purchase or redeem under the terms of the event. This can be a number or the term "any and all".
	MinimumQuantitySought *FinancialInstrumentQuantity21Choice `xml:"MinQtySght,omitempty"`

	// Quantity of equity that makes up the new board lot.
	NewBoardLotQuantity *FinancialInstrumentQuantity22Choice `xml:"NewBrdLotQty,omitempty"`

	// New denomination of the equity following, for example, an increase or decrease in nominal value.
	NewDenominationQuantity *FinancialInstrumentQuantity22Choice `xml:"NewDnmtnQty,omitempty"`

	// Minimum integral amount of securities that each account owner must have remaining after the called amounts are applied.
	BaseDenomination *FinancialInstrumentQuantity22Choice `xml:"BaseDnmtn,omitempty"`

	// Amount used when the called amount is not met by running the lottery with the base denomination.
	IncrementalDenomination *FinancialInstrumentQuantity22Choice `xml:"IncrmtlDnmtn,omitempty"`
}

func (c *CorporateActionQuantity8) AddMaximumQuantity() *FinancialInstrumentQuantity21Choice {
	c.MaximumQuantity = new(FinancialInstrumentQuantity21Choice)
	return c.MaximumQuantity
}

func (c *CorporateActionQuantity8) AddMinimumQuantitySought() *FinancialInstrumentQuantity21Choice {
	c.MinimumQuantitySought = new(FinancialInstrumentQuantity21Choice)
	return c.MinimumQuantitySought
}

func (c *CorporateActionQuantity8) AddNewBoardLotQuantity() *FinancialInstrumentQuantity22Choice {
	c.NewBoardLotQuantity = new(FinancialInstrumentQuantity22Choice)
	return c.NewBoardLotQuantity
}

func (c *CorporateActionQuantity8) AddNewDenominationQuantity() *FinancialInstrumentQuantity22Choice {
	c.NewDenominationQuantity = new(FinancialInstrumentQuantity22Choice)
	return c.NewDenominationQuantity
}

func (c *CorporateActionQuantity8) AddBaseDenomination() *FinancialInstrumentQuantity22Choice {
	c.BaseDenomination = new(FinancialInstrumentQuantity22Choice)
	return c.BaseDenomination
}

func (c *CorporateActionQuantity8) AddIncrementalDenomination() *FinancialInstrumentQuantity22Choice {
	c.IncrementalDenomination = new(FinancialInstrumentQuantity22Choice)
	return c.IncrementalDenomination
}
