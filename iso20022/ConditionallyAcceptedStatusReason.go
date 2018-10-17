package iso20022

// Reason for a conditionally accepted status.
type ConditionallyAcceptedStatusReason1 struct {

	// Reason for a conditionally accepted status in structured form.
	Structured []*ConditionallyAcceptedStatusReason1Code `xml:"Strd"`

	// Reason for a conditionally accepted status in free format text.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *ConditionallyAcceptedStatusReason1) AddStructured(value string) {
	c.Structured = append(c.Structured, (*ConditionallyAcceptedStatusReason1Code)(&value))
}

func (c *ConditionallyAcceptedStatusReason1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}

// Identification of the reason for the conditionally accepted status.
type ConditionallyAcceptedStatusReason2 struct {

	// Reason for the conditionally accepted status.
	Reason *ConditionallyAcceptedStatusReason2Code `xml:"Rsn"`

	// Reason for the conditionally accepted status.
	ExtendedReason *Extended350Code `xml:"XtndedRsn"`

	// Proprietary identification of the reason for the conditionally accepted status.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`

	// Additional information about the conditionally accepted status reason.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *ConditionallyAcceptedStatusReason2) SetReason(value string) {
	c.Reason = (*ConditionallyAcceptedStatusReason2Code)(&value)
}

func (c *ConditionallyAcceptedStatusReason2) SetExtendedReason(value string) {
	c.ExtendedReason = (*Extended350Code)(&value)
}

func (c *ConditionallyAcceptedStatusReason2) AddDataSourceScheme() *GenericIdentification1 {
	c.DataSourceScheme = new(GenericIdentification1)
	return c.DataSourceScheme
}

func (c *ConditionallyAcceptedStatusReason2) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}

// Reason for a conditionally accepted status.
type ConditionallyAcceptedStatusReason3 struct {

	// Reason for the conditionally accepted status expressed as a code.
	Reason *ConditionallyAcceptedStatusReason3Choice `xml:"Rsn"`

	// Additional information about the conditionally accepted reason.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *ConditionallyAcceptedStatusReason3) AddReason() *ConditionallyAcceptedStatusReason3Choice {
	c.Reason = new(ConditionallyAcceptedStatusReason3Choice)
	return c.Reason
}

func (c *ConditionallyAcceptedStatusReason3) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
