package iso20022

// Status is in repair.
type InRepairStatus1 struct {

	// Indicates that there is no reason available or to report.
	NoReason *NoReasonCode `xml:"NoRsn"`

	// Reason of an in repair status in the report.
	Reason *InRepairStatusReason1 `xml:"Rsn"`

	// Proprietary identification for a reason of a specific status in the report.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`
}

func (i *InRepairStatus1) SetNoReason(value string) {
	i.NoReason = (*NoReasonCode)(&value)
}

func (i *InRepairStatus1) AddReason() *InRepairStatusReason1 {
	i.Reason = new(InRepairStatusReason1)
	return i.Reason
}

func (i *InRepairStatus1) AddDataSourceScheme() *GenericIdentification1 {
	i.DataSourceScheme = new(GenericIdentification1)
	return i.DataSourceScheme
}

// Status is in repair.
type InRepairStatus2 struct {

	// Reason for the in-repair status.
	ReasonDetails *InRepairStatusReason3 `xml:"RsnDtls"`

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`
}

func (i *InRepairStatus2) AddReasonDetails() *InRepairStatusReason3 {
	i.ReasonDetails = new(InRepairStatusReason3)
	return i.ReasonDetails
}

func (i *InRepairStatus2) SetNoSpecifiedReason(value string) {
	i.NoSpecifiedReason = (*NoReasonCode)(&value)
}

// Reason for the in repair status.
type InRepairStatus3 struct {

	// Reason for the in-repair status.
	Reason *Max350Text `xml:"Rsn"`

	// Proprietary identification of the reason for the in-repair status.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`
}

func (i *InRepairStatus3) SetReason(value string) {
	i.Reason = (*Max350Text)(&value)
}

func (i *InRepairStatus3) AddDataSourceScheme() *GenericIdentification1 {
	i.DataSourceScheme = new(GenericIdentification1)
	return i.DataSourceScheme
}

func (i *InRepairStatus3) SetNoSpecifiedReason(value string) {
	i.NoSpecifiedReason = (*NoReasonCode)(&value)
}
