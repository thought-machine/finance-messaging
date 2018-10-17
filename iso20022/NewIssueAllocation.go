package iso20022

// Information about the investment account ownership with respect to new issue allocation for a hedge fund.
type NewIssueAllocation1 struct {

	// Indicates whether the investor is eligible to participate in the profits and losses from a new issue.
	Restricted *YesNoIndicator `xml:"Rstrctd"`

	// Reason for exemption.
	ExemptPersonReason *Max350Text `xml:"XmptPrsnRsn,omitempty"`

	// Conditions applicable when the investor is covered by the "de minimis" exemption.
	DeMinimusApplicable *DeMinimusApplicable1 `xml:"DeMnmsAplbl,omitempty"`

	// Conditions applicable when the investor is not covered by the "de minimis" exemption.
	DeMinimusNotApplicable *DeMinimusNotApplicable1 `xml:"DeMnmsNotAplbl,omitempty"`
}

func (n *NewIssueAllocation1) SetRestricted(value string) {
	n.Restricted = (*YesNoIndicator)(&value)
}

func (n *NewIssueAllocation1) SetExemptPersonReason(value string) {
	n.ExemptPersonReason = (*Max350Text)(&value)
}

func (n *NewIssueAllocation1) AddDeMinimusApplicable() *DeMinimusApplicable1 {
	n.DeMinimusApplicable = new(DeMinimusApplicable1)
	return n.DeMinimusApplicable
}

func (n *NewIssueAllocation1) AddDeMinimusNotApplicable() *DeMinimusNotApplicable1 {
	n.DeMinimusNotApplicable = new(DeMinimusNotApplicable1)
	return n.DeMinimusNotApplicable
}

// Information about the investment account ownership with respect to new issue allocation for a hedge fund.
type NewIssueAllocation2 struct {

	// Indicates whether the investor is eligible to participate in the profits and losses from a new issue.
	Restricted *YesNoIndicator `xml:"Rstrctd"`

	// Reason for exemption.
	ExemptPersonReason *Max350Text `xml:"XmptPrsnRsn,omitempty"`

	// Conditions applicable when the investor is covered by the "de minimis" exemption.
	DeMinimus *DeMinimus1Choice `xml:"DeMnms,omitempty"`
}

func (n *NewIssueAllocation2) SetRestricted(value string) {
	n.Restricted = (*YesNoIndicator)(&value)
}

func (n *NewIssueAllocation2) SetExemptPersonReason(value string) {
	n.ExemptPersonReason = (*Max350Text)(&value)
}

func (n *NewIssueAllocation2) AddDeMinimus() *DeMinimus1Choice {
	n.DeMinimus = new(DeMinimus1Choice)
	return n.DeMinimus
}
