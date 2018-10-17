package iso20022

// Specifies the opening and valuation conditions for the non deliverable forward
type NonDeliverableForwardConditions1 struct {

	// Specifies whether the instruction is an NDF opening or fixing.
	OpeningIndicator *YesNoIndicator `xml:"OpngInd"`

	// Specifies either the conditions for an NDF oepning or an NDF fixing confirmation.
	OpeningFixingConditions *NDFOpeningFixing1Choice `xml:"OpngFxgConds"`
}

func (n *NonDeliverableForwardConditions1) SetOpeningIndicator(value string) {
	n.OpeningIndicator = (*YesNoIndicator)(&value)
}

func (n *NonDeliverableForwardConditions1) AddOpeningFixingConditions() *NDFOpeningFixing1Choice {
	n.OpeningFixingConditions = new(NDFOpeningFixing1Choice)
	return n.OpeningFixingConditions
}

// Specifies the opening and valuation conditions for the non deliverable forward
type NonDeliverableForwardConditions2 struct {

	// Provides the opening information associated with an NDF trade.
	OpeningConditions *OpeningConditions1 `xml:"OpngConds"`

	// Provides the additional information for an NDF as supplied on a fixing instruction.
	FixingConditions *FixingConditions1 `xml:"FxgConds,omitempty"`
}

func (n *NonDeliverableForwardConditions2) AddOpeningConditions() *OpeningConditions1 {
	n.OpeningConditions = new(OpeningConditions1)
	return n.OpeningConditions
}

func (n *NonDeliverableForwardConditions2) AddFixingConditions() *FixingConditions1 {
	n.FixingConditions = new(FixingConditions1)
	return n.FixingConditions
}
