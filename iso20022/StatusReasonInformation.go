package iso20022

// Further information on the status reason of the transaction.
type StatusReasonInformation1 struct {

	// Party issuing the status.
	StatusOriginator *PartyIdentification8 `xml:"StsOrgtr,omitempty"`

	// Specifies the reason for the status report.
	StatusReason *StatusReason1Choice `xml:"StsRsn,omitempty"`

	// Further details on the status reason.
	//
	// Usage: Additional information can be used for several purposes, e.g. to report repaired information, or to further detail the status reason.
	AdditionalStatusReasonInformation []*Max105Text `xml:"AddtlStsRsnInf,omitempty"`
}

func (s *StatusReasonInformation1) AddStatusOriginator() *PartyIdentification8 {
	s.StatusOriginator = new(PartyIdentification8)
	return s.StatusOriginator
}

func (s *StatusReasonInformation1) AddStatusReason() *StatusReason1Choice {
	s.StatusReason = new(StatusReason1Choice)
	return s.StatusReason
}

func (s *StatusReasonInformation1) AddAdditionalStatusReasonInformation(value string) {
	s.AdditionalStatusReasonInformation = append(s.AdditionalStatusReasonInformation, (*Max105Text)(&value))
}

// Set of elements used to provide information on the status reason of the transaction.
type StatusReasonInformation8 struct {

	// Party that issues the status.
	Originator *PartyIdentification32 `xml:"Orgtr,omitempty"`

	// Specifies the reason for the status report.
	Reason *StatusReason6Choice `xml:"Rsn,omitempty"`

	// Further details on the status reason.
	//
	// Usage: Additional information can be used for several purposes such as the reporting of repaired information.
	AdditionalInformation []*Max105Text `xml:"AddtlInf,omitempty"`
}

func (s *StatusReasonInformation8) AddOriginator() *PartyIdentification32 {
	s.Originator = new(PartyIdentification32)
	return s.Originator
}

func (s *StatusReasonInformation8) AddReason() *StatusReason6Choice {
	s.Reason = new(StatusReason6Choice)
	return s.Reason
}

func (s *StatusReasonInformation8) AddAdditionalInformation(value string) {
	s.AdditionalInformation = append(s.AdditionalInformation, (*Max105Text)(&value))
}

// Set of elements used to provide information on the status reason of the transaction.
type StatusReasonInformation9 struct {

	// Party that issues the status.
	Originator *PartyIdentification43 `xml:"Orgtr,omitempty"`

	// Specifies the reason for the status report.
	Reason *StatusReason6Choice `xml:"Rsn,omitempty"`

	// Further details on the status reason.
	//
	// Usage: Additional information can be used for several purposes such as the reporting of repaired information.
	AdditionalInformation []*Max105Text `xml:"AddtlInf,omitempty"`
}

func (s *StatusReasonInformation9) AddOriginator() *PartyIdentification43 {
	s.Originator = new(PartyIdentification43)
	return s.Originator
}

func (s *StatusReasonInformation9) AddReason() *StatusReason6Choice {
	s.Reason = new(StatusReason6Choice)
	return s.Reason
}

func (s *StatusReasonInformation9) AddAdditionalInformation(value string) {
	s.AdditionalInformation = append(s.AdditionalInformation, (*Max105Text)(&value))
}
