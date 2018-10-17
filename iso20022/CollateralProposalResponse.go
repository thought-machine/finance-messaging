package iso20022

// Provides the collateral proposal response for the variation margin and optionaly the segregated independent amount.
type CollateralProposalResponse2 struct {

	// Provides the collateral proposal response for the variation margin.
	VariationMargin *CollateralProposalResponseType2 `xml:"VartnMrgn"`

	// Provides the collateral proposal response for the segregated independent amount.
	SegregatedIndependentAmount *CollateralProposalResponseType2 `xml:"SgrtdIndpdntAmt,omitempty"`
}

func (c *CollateralProposalResponse2) AddVariationMargin() *CollateralProposalResponseType2 {
	c.VariationMargin = new(CollateralProposalResponseType2)
	return c.VariationMargin
}

func (c *CollateralProposalResponse2) AddSegregatedIndependentAmount() *CollateralProposalResponseType2 {
	c.SegregatedIndependentAmount = new(CollateralProposalResponseType2)
	return c.SegregatedIndependentAmount
}

// Provides the collateral proposal response for the variation margin and optionally the segregated independent amount.
type CollateralProposalResponse3 struct {

	// Provides the collateral proposal response for the variation margin.
	VariationMargin *CollateralProposalResponseType3 `xml:"VartnMrgn"`

	// Provides the collateral proposal response for the segregated independent amount.
	SegregatedIndependentAmount *CollateralProposalResponseType3 `xml:"SgrtdIndpdntAmt,omitempty"`
}

func (c *CollateralProposalResponse3) AddVariationMargin() *CollateralProposalResponseType3 {
	c.VariationMargin = new(CollateralProposalResponseType3)
	return c.VariationMargin
}

func (c *CollateralProposalResponse3) AddSegregatedIndependentAmount() *CollateralProposalResponseType3 {
	c.SegregatedIndependentAmount = new(CollateralProposalResponseType3)
	return c.SegregatedIndependentAmount
}
