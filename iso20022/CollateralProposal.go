package iso20022

// Provides details about the proposal for the variation margin and optionaly the segregated independent amount.
type CollateralProposal4 struct {

	// Provides details about the proposal for the variation margin.
	VariationMargin *CollateralMovement5 `xml:"VartnMrgn"`

	// Provides details about the proposal for the segregated independent amount.
	SegregatedIndependentAmount *CollateralMovement5 `xml:"SgrtdIndpdntAmt,omitempty"`
}

func (c *CollateralProposal4) AddVariationMargin() *CollateralMovement5 {
	c.VariationMargin = new(CollateralMovement5)
	return c.VariationMargin
}

func (c *CollateralProposal4) AddSegregatedIndependentAmount() *CollateralMovement5 {
	c.SegregatedIndependentAmount = new(CollateralMovement5)
	return c.SegregatedIndependentAmount
}

// Provides details about the proposal for the variation margin and optionaly the segregated independent amount.
type CollateralProposal5 struct {

	// Provides details about the proposal for the variation margin.
	VariationMargin *CollateralMovement7 `xml:"VartnMrgn"`

	// Provides details about the proposal for the segregated independent amount.
	SegregatedIndependentAmount *CollateralMovement7 `xml:"SgrtdIndpdntAmt,omitempty"`
}

func (c *CollateralProposal5) AddVariationMargin() *CollateralMovement7 {
	c.VariationMargin = new(CollateralMovement7)
	return c.VariationMargin
}

func (c *CollateralProposal5) AddSegregatedIndependentAmount() *CollateralMovement7 {
	c.SegregatedIndependentAmount = new(CollateralMovement7)
	return c.SegregatedIndependentAmount
}

// Provides details about the proposal for the variation margin and optionally the segregated independent amount.
type CollateralProposal6 struct {

	// Provides details about the proposal for the variation margin.
	VariationMargin *CollateralMovement10 `xml:"VartnMrgn"`

	// Provides details about the proposal for the segregated independent amount.
	SegregatedIndependentAmount *CollateralMovement10 `xml:"SgrtdIndpdntAmt,omitempty"`
}

func (c *CollateralProposal6) AddVariationMargin() *CollateralMovement10 {
	c.VariationMargin = new(CollateralMovement10)
	return c.VariationMargin
}

func (c *CollateralProposal6) AddSegregatedIndependentAmount() *CollateralMovement10 {
	c.SegregatedIndependentAmount = new(CollateralMovement10)
	return c.SegregatedIndependentAmount
}
