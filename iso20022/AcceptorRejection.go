package iso20022

// Reject of an exchange.
type AcceptorRejection1 struct {

	// Reject reason of the request or the advice.
	RejectReason *RejectReason1Code `xml:"RjctRsn"`

	// Additional information related to the reject of the exchange.
	AdditionalInformation *Max500Text `xml:"AddtlInf,omitempty"`

	// Original request that caused the recipient party to reject it.
	MessageInError *Max5000Binary `xml:"MsgInErr,omitempty"`
}

func (a *AcceptorRejection1) SetRejectReason(value string) {
	a.RejectReason = (*RejectReason1Code)(&value)
}

func (a *AcceptorRejection1) SetAdditionalInformation(value string) {
	a.AdditionalInformation = (*Max500Text)(&value)
}

func (a *AcceptorRejection1) SetMessageInError(value string) {
	a.MessageInError = (*Max5000Binary)(&value)
}

// Reject of an exchange.
type AcceptorRejection2 struct {

	// Reject reason of the request or the advice.
	RejectReason *RejectReason1Code `xml:"RjctRsn"`

	// Additional information related to the reject of the exchange.
	AdditionalInformation *Max500Text `xml:"AddtlInf,omitempty"`

	// Original request that caused the recipient party to reject it.
	MessageInError *Max100KBinary `xml:"MsgInErr,omitempty"`
}

func (a *AcceptorRejection2) SetRejectReason(value string) {
	a.RejectReason = (*RejectReason1Code)(&value)
}

func (a *AcceptorRejection2) SetAdditionalInformation(value string) {
	a.AdditionalInformation = (*Max500Text)(&value)
}

func (a *AcceptorRejection2) SetMessageInError(value string) {
	a.MessageInError = (*Max100KBinary)(&value)
}

// Reject of an exchange.
type AcceptorRejection3 struct {

	// Reject reason of the request or the advice.
	RejectReason *RejectReason2Code `xml:"RjctRsn"`

	// Additional information related to the reject of the exchange.
	AdditionalInformation *Max500Text `xml:"AddtlInf,omitempty"`

	// Original request that caused the recipient party to reject it.
	MessageInError *Max100KBinary `xml:"MsgInErr,omitempty"`
}

func (a *AcceptorRejection3) SetRejectReason(value string) {
	a.RejectReason = (*RejectReason2Code)(&value)
}

func (a *AcceptorRejection3) SetAdditionalInformation(value string) {
	a.AdditionalInformation = (*Max500Text)(&value)
}

func (a *AcceptorRejection3) SetMessageInError(value string) {
	a.MessageInError = (*Max100KBinary)(&value)
}

// Reject of an exchange.
type AcceptorRejection4 struct {

	// Reject reason of the message.
	RejectReason *RejectReason1Code `xml:"RjctRsn"`

	// Detailed description of an error that caused the rejection for further analysis.
	ErrorReporting []*ErrorReporting1 `xml:"ErrRptg,omitempty"`

	// Original request that caused the party to reject it.
	MessageInError *Max100KBinary `xml:"MsgInErr,omitempty"`
}

func (a *AcceptorRejection4) SetRejectReason(value string) {
	a.RejectReason = (*RejectReason1Code)(&value)
}

func (a *AcceptorRejection4) AddErrorReporting() *ErrorReporting1 {
	newValue := new(ErrorReporting1)
	a.ErrorReporting = append(a.ErrorReporting, newValue)
	return newValue
}

func (a *AcceptorRejection4) SetMessageInError(value string) {
	a.MessageInError = (*Max100KBinary)(&value)
}
