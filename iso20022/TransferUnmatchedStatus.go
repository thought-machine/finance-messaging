package iso20022

// Status is unmatched.
type TransferUnmatchedStatus struct {

	// Reason for an unmatched status in the report.
	Reason *TransferUnmatchedStatusReason1 `xml:"Rsn"`

	// Proprietary identification for a reason of a specific status in the report.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`

	// Indicates that there is no reason available or to report.
	NoReason *NoReasonCode `xml:"NoRsn"`
}

func (t *TransferUnmatchedStatus) AddReason() *TransferUnmatchedStatusReason1 {
	t.Reason = new(TransferUnmatchedStatusReason1)
	return t.Reason
}

func (t *TransferUnmatchedStatus) AddDataSourceScheme() *GenericIdentification1 {
	t.DataSourceScheme = new(GenericIdentification1)
	return t.DataSourceScheme
}

func (t *TransferUnmatchedStatus) SetNoReason(value string) {
	t.NoReason = (*NoReasonCode)(&value)
}

// Reason for the unmatched status.
type TransferUnmatchedStatus2 struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the unmatched status.
	Reason *TransferUnmatchedReason2Code `xml:"Rsn"`

	// Reason for the unmatched status.
	ExtendedReason *Extended350Code `xml:"XtndedRsn"`

	// Proprietary identification of the reason for the unmatched status.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`
}

func (t *TransferUnmatchedStatus2) SetNoSpecifiedReason(value string) {
	t.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (t *TransferUnmatchedStatus2) SetReason(value string) {
	t.Reason = (*TransferUnmatchedReason2Code)(&value)
}

func (t *TransferUnmatchedStatus2) SetExtendedReason(value string) {
	t.ExtendedReason = (*Extended350Code)(&value)
}

func (t *TransferUnmatchedStatus2) AddDataSourceScheme() *GenericIdentification1 {
	t.DataSourceScheme = new(GenericIdentification1)
	return t.DataSourceScheme
}
