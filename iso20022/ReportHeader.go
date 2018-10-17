package iso20022

// Specifies generic information about an investigation report.
type ReportHeader struct {

	// Identification of the report.
	Identification *Max35Text `xml:"Id"`

	// Party reporting the status of the case.
	From *AnyBICIdentifier `xml:"Fr"`

	// Party to which the status of the case is reported.
	To *AnyBICIdentifier `xml:"To"`

	// Creation date and time of the report generation.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`
}

func (r *ReportHeader) SetIdentification(value string) {
	r.Identification = (*Max35Text)(&value)
}

func (r *ReportHeader) SetFrom(value string) {
	r.From = (*AnyBICIdentifier)(&value)
}

func (r *ReportHeader) SetTo(value string) {
	r.To = (*AnyBICIdentifier)(&value)
}

func (r *ReportHeader) SetCreationDateTime(value string) {
	r.CreationDateTime = (*ISODateTime)(&value)
}

// Specifies generic information about an investigation report.
type ReportHeader2 struct {

	// Point to point reference as assigned by the case assigner to unambiguously identify the case status report.
	Identification *Max35Text `xml:"Id"`

	// Party reporting the status of the investigation case.
	From *Party7Choice `xml:"Fr"`

	// Party to which the status of the case is reported.
	To *Party7Choice `xml:"To"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`
}

func (r *ReportHeader2) SetIdentification(value string) {
	r.Identification = (*Max35Text)(&value)
}

func (r *ReportHeader2) AddFrom() *Party7Choice {
	r.From = new(Party7Choice)
	return r.From
}

func (r *ReportHeader2) AddTo() *Party7Choice {
	r.To = new(Party7Choice)
	return r.To
}

func (r *ReportHeader2) SetCreationDateTime(value string) {
	r.CreationDateTime = (*ISODateTime)(&value)
}

// Provides header details on the report.
type ReportHeader3 struct {

	// Identification of a report billing statement.
	ReportIdentification *Max35Text `xml:"RptId"`

	// Provides details on the page number of the message.
	//
	// Usage: The pagination of the message is only allowed when agreed between the parties.
	MessagePagination *Pagination `xml:"MsgPgntn,omitempty"`
}

func (r *ReportHeader3) SetReportIdentification(value string) {
	r.ReportIdentification = (*Max35Text)(&value)
}

func (r *ReportHeader3) AddMessagePagination() *Pagination {
	r.MessagePagination = new(Pagination)
	return r.MessagePagination
}

// Specifies generic information about an investigation report.
type ReportHeader4 struct {

	// Point to point reference as assigned by the case assigner to unambiguously identify the case status report.
	Identification *Max35Text `xml:"Id"`

	// Party reporting the status of the investigation case.
	From *Party12Choice `xml:"Fr"`

	// Party to which the status of the case is reported.
	To *Party12Choice `xml:"To"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`
}

func (r *ReportHeader4) SetIdentification(value string) {
	r.Identification = (*Max35Text)(&value)
}

func (r *ReportHeader4) AddFrom() *Party12Choice {
	r.From = new(Party12Choice)
	return r.From
}

func (r *ReportHeader4) AddTo() *Party12Choice {
	r.To = new(Party12Choice)
	return r.To
}

func (r *ReportHeader4) SetCreationDateTime(value string) {
	r.CreationDateTime = (*ISODateTime)(&value)
}
