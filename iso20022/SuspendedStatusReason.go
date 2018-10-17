package iso20022

// Reason for a suspended status.
type SuspendedStatusReason1 struct {

	// Reason for a suspended status in structured form.
	Structured []*SuspendedStatusReason2Code `xml:"Strd"`

	// Reason for a suspended status in free format text.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (s *SuspendedStatusReason1) AddStructured(value string) {
	s.Structured = append(s.Structured, (*SuspendedStatusReason2Code)(&value))
}

func (s *SuspendedStatusReason1) SetAdditionalInformation(value string) {
	s.AdditionalInformation = (*Max350Text)(&value)
}

// Identification of the reason for the suspended status.
type SuspendedStatusReason2 struct {

	// Reason for the suspended status.
	Reason *SuspendedStatusReason3Code `xml:"Rsn"`

	// Reason for the suspended status.
	ExtendedReason *Extended350Code `xml:"XtndedRsn"`

	// Proprietary identification of the reason for the suspended status.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`

	// Additional information about the suspended status reason.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (s *SuspendedStatusReason2) SetReason(value string) {
	s.Reason = (*SuspendedStatusReason3Code)(&value)
}

func (s *SuspendedStatusReason2) SetExtendedReason(value string) {
	s.ExtendedReason = (*Extended350Code)(&value)
}

func (s *SuspendedStatusReason2) AddDataSourceScheme() *GenericIdentification1 {
	s.DataSourceScheme = new(GenericIdentification1)
	return s.DataSourceScheme
}

func (s *SuspendedStatusReason2) SetAdditionalInformation(value string) {
	s.AdditionalInformation = (*Max350Text)(&value)
}

// Reason for a suspended status.
type SuspendedStatusReason4 struct {

	// Reason for the conditionally accepted status expressed as a code.
	Reason *SuspendedStatusReason5Choice `xml:"Rsn"`

	// Additional information about the suspended reason.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (s *SuspendedStatusReason4) AddReason() *SuspendedStatusReason5Choice {
	s.Reason = new(SuspendedStatusReason5Choice)
	return s.Reason
}

func (s *SuspendedStatusReason4) SetAdditionalInformation(value string) {
	s.AdditionalInformation = (*Max350Text)(&value)
}
