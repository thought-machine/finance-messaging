package iso20022

// Status is partially settled.
type PartiallySettledStatus1 struct {

	// Reason for the partially settled status.
	Reason *SettledStatusReason1Code `xml:"Rsn"`

	// Reason for the partially settled status.
	ExtendedReason *Extended350Code `xml:"XtndedRsn"`

	// Proprietary identification of the reason for the partially settled status in the report.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`

	// Additional information about the partially settled status reason.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (p *PartiallySettledStatus1) SetReason(value string) {
	p.Reason = (*SettledStatusReason1Code)(&value)
}

func (p *PartiallySettledStatus1) SetExtendedReason(value string) {
	p.ExtendedReason = (*Extended350Code)(&value)
}

func (p *PartiallySettledStatus1) AddDataSourceScheme() *GenericIdentification1 {
	p.DataSourceScheme = new(GenericIdentification1)
	return p.DataSourceScheme
}

func (p *PartiallySettledStatus1) SetAdditionalInformation(value string) {
	p.AdditionalInformation = (*Max350Text)(&value)
}

// Reason for a partially settled status.
type PartiallySettledStatus10 struct {

	// Reason for the partially settled status.
	Reason *PartiallySettled21Choice `xml:"Rsn"`

	// Additional information about the partially settled reason.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (p *PartiallySettledStatus10) AddReason() *PartiallySettled21Choice {
	p.Reason = new(PartiallySettled21Choice)
	return p.Reason
}

func (p *PartiallySettledStatus10) SetAdditionalInformation(value string) {
	p.AdditionalInformation = (*Max350Text)(&value)
}
