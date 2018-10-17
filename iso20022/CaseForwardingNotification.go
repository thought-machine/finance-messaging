package iso20022

// Status of a case resulting from a case assignment.
type CaseForwardingNotification struct {

	// Justification for the forward action.
	Justification *CaseForwardingNotification1Code `xml:"Justfn"`
}

func (c *CaseForwardingNotification) SetJustification(value string) {
	c.Justification = (*CaseForwardingNotification1Code)(&value)
}

// Status of a case resulting from a case assignment.
type CaseForwardingNotification3 struct {

	// Justification for the forward action.
	Justification *CaseForwardingNotification3Code `xml:"Justfn"`
}

func (c *CaseForwardingNotification3) SetJustification(value string) {
	c.Justification = (*CaseForwardingNotification3Code)(&value)
}
