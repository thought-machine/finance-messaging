package iso20022

// Status is rejected.
type ConfirmationRejectedStatus1 struct {

	// Reason for the rejected status.
	Reason *RejectedConfirmationStatusReason1Code `xml:"Rsn"`

	// Reason for the rejected status.
	ExtendedReason *Extended350Code `xml:"XtndedRsn"`

	// Proprietary identification of the reason for a rejected status.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`
}

func (c *ConfirmationRejectedStatus1) SetReason(value string) {
	c.Reason = (*RejectedConfirmationStatusReason1Code)(&value)
}

func (c *ConfirmationRejectedStatus1) SetExtendedReason(value string) {
	c.ExtendedReason = (*Extended350Code)(&value)
}

func (c *ConfirmationRejectedStatus1) AddDataSourceScheme() *GenericIdentification1 {
	c.DataSourceScheme = new(GenericIdentification1)
	return c.DataSourceScheme
}

// Rejection of a confirmation.
type ConfirmationRejectedStatus2 struct {

	// Reason for the rejected status.
	Reason *ConfirmationRejectedReason1Choice `xml:"Rsn,omitempty"`

	// Additional information about the rejected reason.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *ConfirmationRejectedStatus2) AddReason() *ConfirmationRejectedReason1Choice {
	c.Reason = new(ConfirmationRejectedReason1Choice)
	return c.Reason
}

func (c *ConfirmationRejectedStatus2) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
