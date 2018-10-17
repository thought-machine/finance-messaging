package iso20022

// Specifies the reason for cancelling a meeting.
type MeetingCancellationReason1 struct {

	// Specifies the reason for cancelling a meeting in coded form.
	Code *MeetingCancellationReason2Code `xml:"Cd"`

	// Specifies the reason for cancelling a meeting in free text form.
	ExtendedCode *Extended350Code `xml:"XtndedCd"`

	// Provides more information on the reason for cancelling a meeting in free format form.
	CancellationReason *Max140Text `xml:"CxlRsn,omitempty"`
}

func (m *MeetingCancellationReason1) SetCode(value string) {
	m.Code = (*MeetingCancellationReason2Code)(&value)
}

func (m *MeetingCancellationReason1) SetExtendedCode(value string) {
	m.ExtendedCode = (*Extended350Code)(&value)
}

func (m *MeetingCancellationReason1) SetCancellationReason(value string) {
	m.CancellationReason = (*Max140Text)(&value)
}

// Specifies the reason for cancelling a meeting.
type MeetingCancellationReason2 struct {

	// Reason for cancelling a meeting.
	CancellationReasonCode *MeetingCancellationReason1Choice `xml:"CxlRsnCd,omitempty"`

	// Provides more information on the reason for cancelling a meeting in free format form.
	CancellationReason *Max140Text `xml:"CxlRsn,omitempty"`
}

func (m *MeetingCancellationReason2) AddCancellationReasonCode() *MeetingCancellationReason1Choice {
	m.CancellationReasonCode = new(MeetingCancellationReason1Choice)
	return m.CancellationReasonCode
}

func (m *MeetingCancellationReason2) SetCancellationReason(value string) {
	m.CancellationReason = (*Max140Text)(&value)
}
