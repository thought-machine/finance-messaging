package iso20022

// Specifies the status and optionally the sub status.
type StatusAndSubStatus1 struct {

	// Status expressed as a code.
	StatusCode *Status13Choice `xml:"StsCd"`

	// Sub status expressed as a code.
	SubStatusCode *Exact4AlphaNumericText `xml:"SubStsCd,omitempty"`
}

func (s *StatusAndSubStatus1) AddStatusCode() *Status13Choice {
	s.StatusCode = new(Status13Choice)
	return s.StatusCode
}

func (s *StatusAndSubStatus1) SetSubStatusCode(value string) {
	s.SubStatusCode = (*Exact4AlphaNumericText)(&value)
}

// Specifies the status and optionally the sub status.
type StatusAndSubStatus2 struct {

	// Status expressed as a code.
	StatusCode *Status27Choice `xml:"StsCd"`

	// Sub status expressed as a code.
	SubStatusCode *Exact4AlphaNumericText `xml:"SubStsCd,omitempty"`
}

func (s *StatusAndSubStatus2) AddStatusCode() *Status27Choice {
	s.StatusCode = new(Status27Choice)
	return s.StatusCode
}

func (s *StatusAndSubStatus2) SetSubStatusCode(value string) {
	s.SubStatusCode = (*Exact4AlphaNumericText)(&value)
}
