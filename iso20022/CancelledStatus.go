package iso20022

// Status is cancelled.
type CancelledStatus1 struct {

	// Indicates that there is no reason available or to report.
	NoReason *NoReasonCode `xml:"NoRsn"`

	// Reason for a cancelled status in the report.
	Reason *CancelledStatusReason1 `xml:"Rsn"`

	// Proprietary identification of a reason for a cancelled status in the report.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`
}

func (c *CancelledStatus1) SetNoReason(value string) {
	c.NoReason = (*NoReasonCode)(&value)
}

func (c *CancelledStatus1) AddReason() *CancelledStatusReason1 {
	c.Reason = new(CancelledStatusReason1)
	return c.Reason
}

func (c *CancelledStatus1) AddDataSourceScheme() *GenericIdentification1 {
	c.DataSourceScheme = new(GenericIdentification1)
	return c.DataSourceScheme
}

// Status is cancelled.
type CancelledStatus2 struct {

	// Reason for the cancelled status.
	Reason *CancelledStatusReason2Code `xml:"Rsn"`

	// Reason for the cancelled status.
	ExtendedReason *Extended350Code `xml:"XtndedRsn"`

	// Proprietary identification of the reason for the cancelled status.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`
}

func (c *CancelledStatus2) SetReason(value string) {
	c.Reason = (*CancelledStatusReason2Code)(&value)
}

func (c *CancelledStatus2) SetExtendedReason(value string) {
	c.ExtendedReason = (*Extended350Code)(&value)
}

func (c *CancelledStatus2) AddDataSourceScheme() *GenericIdentification1 {
	c.DataSourceScheme = new(GenericIdentification1)
	return c.DataSourceScheme
}

func (c *CancelledStatus2) SetNoSpecifiedReason(value string) {
	c.NoSpecifiedReason = (*NoReasonCode)(&value)
}

// Reason for the cancelled status.
type CancelledStatus3 struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the cancelled status.
	Reason *CancelledStatusReason3Code `xml:"Rsn"`

	// Reason for the cancelled status.
	ExtendedReason *Extended350Code `xml:"XtndedRsn"`

	// Proprietary identification of the reason for the cancelled status.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`
}

func (c *CancelledStatus3) SetNoSpecifiedReason(value string) {
	c.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (c *CancelledStatus3) SetReason(value string) {
	c.Reason = (*CancelledStatusReason3Code)(&value)
}

func (c *CancelledStatus3) SetExtendedReason(value string) {
	c.ExtendedReason = (*Extended350Code)(&value)
}

func (c *CancelledStatus3) AddDataSourceScheme() *GenericIdentification1 {
	c.DataSourceScheme = new(GenericIdentification1)
	return c.DataSourceScheme
}
