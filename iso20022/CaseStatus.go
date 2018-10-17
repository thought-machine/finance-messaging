package iso20022

// Defines the status of an investigation case.
type CaseStatus struct {

	// Date and time of the status.
	DateTime *ISODateTime `xml:"DtTm"`

	// Status of the case.
	CaseStatus *CaseStatus1Code `xml:"CaseSts"`

	// Status of the investigation.
	InvestigationStatus *InvestigationExecutionConfirmation1Code `xml:"InvstgtnSts,omitempty"`

	// Free text justification of the status.
	Reason *Max140Text `xml:"Rsn,omitempty"`
}

func (c *CaseStatus) SetDateTime(value string) {
	c.DateTime = (*ISODateTime)(&value)
}

func (c *CaseStatus) SetCaseStatus(value string) {
	c.CaseStatus = (*CaseStatus1Code)(&value)
}

func (c *CaseStatus) SetInvestigationStatus(value string) {
	c.InvestigationStatus = (*InvestigationExecutionConfirmation1Code)(&value)
}

func (c *CaseStatus) SetReason(value string) {
	c.Reason = (*Max140Text)(&value)
}

// Defines the status of an investigation case.
type CaseStatus2 struct {

	// Date and time of the status.
	DateTime *ISODateTime `xml:"DtTm"`

	// Status of the case.
	CaseStatus *CaseStatus2Code `xml:"CaseSts"`

	// Free text justification of the status.
	Reason *Max140Text `xml:"Rsn,omitempty"`
}

func (c *CaseStatus2) SetDateTime(value string) {
	c.DateTime = (*ISODateTime)(&value)
}

func (c *CaseStatus2) SetCaseStatus(value string) {
	c.CaseStatus = (*CaseStatus2Code)(&value)
}

func (c *CaseStatus2) SetReason(value string) {
	c.Reason = (*Max140Text)(&value)
}
