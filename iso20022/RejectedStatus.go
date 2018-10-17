package iso20022

// Reason for a rejected status.
type RejectedStatus10 struct {

	// Reason for the rejected status.
	Reason *RejectedReason21Choice `xml:"Rsn,omitempty"`

	// Additional information about the rejected reason.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (r *RejectedStatus10) AddReason() *RejectedReason21Choice {
	r.Reason = new(RejectedReason21Choice)
	return r.Reason
}

func (r *RejectedStatus10) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*Max350Text)(&value)
}

// Status is rejected.
type RejectedStatus3 struct {

	// Reason for a rejected status in the report.
	Reason []*RejectedStatusReason6 `xml:"Rsn"`

	// Proprietary identification of a reason for a rejected status in the report.
	DataSourceScheme []*GenericIdentification1 `xml:"DataSrcSchme"`
}

func (r *RejectedStatus3) AddReason() *RejectedStatusReason6 {
	newValue := new(RejectedStatusReason6)
	r.Reason = append(r.Reason, newValue)
	return newValue
}

func (r *RejectedStatus3) AddDataSourceScheme() *GenericIdentification1 {
	newValue := new(GenericIdentification1)
	r.DataSourceScheme = append(r.DataSourceScheme, newValue)
	return newValue
}

// Status is rejected.
type RejectedStatus4 struct {

	// Reason for a rejected status.
	Reason *RejectedStatusReason4 `xml:"Rsn"`

	// Proprietary identification for a reason of a rejected status in the report.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`
}

func (r *RejectedStatus4) AddReason() *RejectedStatusReason4 {
	r.Reason = new(RejectedStatusReason4)
	return r.Reason
}

func (r *RejectedStatus4) AddDataSourceScheme() *GenericIdentification1 {
	r.DataSourceScheme = new(GenericIdentification1)
	return r.DataSourceScheme
}

// Status is rejected.
type RejectedStatus5 struct {

	// Reason for a rejected status.
	Reason *RejectedStatusReason6Code `xml:"Rsn"`

	// Reason for a rejected status.
	ExtendedReason *Extended350Code `xml:"XtndedRsn"`
}

func (r *RejectedStatus5) SetReason(value string) {
	r.Reason = (*RejectedStatusReason6Code)(&value)
}

func (r *RejectedStatus5) SetExtendedReason(value string) {
	r.ExtendedReason = (*Extended350Code)(&value)
}

// Status is rejected.
type RejectedStatus6 struct {

	// Reason for the rejected status.
	Reason *RejectedStatusReason7Code `xml:"Rsn"`

	// Reason for the rejected status.
	ExtendedReason *Extended350Code `xml:"XtndedRsn"`

	// Proprietary identification of the reason for the rejected status.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`

	// Additional information about the rejected status reason.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (r *RejectedStatus6) SetReason(value string) {
	r.Reason = (*RejectedStatusReason7Code)(&value)
}

func (r *RejectedStatus6) SetExtendedReason(value string) {
	r.ExtendedReason = (*Extended350Code)(&value)
}

func (r *RejectedStatus6) AddDataSourceScheme() *GenericIdentification1 {
	r.DataSourceScheme = new(GenericIdentification1)
	return r.DataSourceScheme
}

func (r *RejectedStatus6) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*Max350Text)(&value)
}

// Status is rejected.
type RejectedStatus7 struct {

	// Reason for the rejected status.
	Reason *RejectedStatusReason8Code `xml:"Rsn"`

	// Reason for the rejected status.
	ExtendedReason *Extended350Code `xml:"XtndedRsn"`

	// Proprietary identification of the reason for the rejcted status.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`
}

func (r *RejectedStatus7) SetReason(value string) {
	r.Reason = (*RejectedStatusReason8Code)(&value)
}

func (r *RejectedStatus7) SetExtendedReason(value string) {
	r.ExtendedReason = (*Extended350Code)(&value)
}

func (r *RejectedStatus7) AddDataSourceScheme() *GenericIdentification1 {
	r.DataSourceScheme = new(GenericIdentification1)
	return r.DataSourceScheme
}

// Reason for a rejected status.
type RejectedStatus9 struct {

	// Reason for the rejected status.
	Reason *RejectedReason20Choice `xml:"Rsn,omitempty"`

	// Additional information about the rejected reason.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (r *RejectedStatus9) AddReason() *RejectedReason20Choice {
	r.Reason = new(RejectedReason20Choice)
	return r.Reason
}

func (r *RejectedStatus9) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*Max350Text)(&value)
}
