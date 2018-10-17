package tsmt

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document00200103 struct {
	XMLName xml.Name           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.002.001.03 Document"`
	Message *ActivityReportV03 `xml:"ActvtyRpt"`
}

func (d *Document00200103) AddMessage() *ActivityReportV03 {
	d.Message = new(ActivityReportV03)
	return d.Message
}

// Scope
// The ActivityReport message is sent by the matching application to the requester of an activity report.
// This message is used to report on all transactions for which an activity has taken place within a given time span.
// Usage
// The ActivityReport message can be sent
// - at a pre-determined time every 24 hours. The message reports on all transactions that the requester is involved in and for which an activity has taken place within the last 24 hours.
// - on demand in response to an ActivityReportRequest message. The message reports on all transactions that the requester is involved in and for which an activity has taken place within a time span specified by the requester in the ActivityReportRequest message.
type ActivityReportV03 struct {

	// Identifies the report.
	ReportIdentification *iso20022.MessageIdentification1 `xml:"RptId"`

	// Reference to the previous message requesting the report.
	RelatedMessageReference *iso20022.MessageIdentification1 `xml:"RltdMsgRef,omitempty"`

	// Describes the events that occurred for one transaction.
	Report []*iso20022.ActivityReportItems2 `xml:"Rpt,omitempty"`
}

func (a *ActivityReportV03) AddReportIdentification() *iso20022.MessageIdentification1 {
	a.ReportIdentification = new(iso20022.MessageIdentification1)
	return a.ReportIdentification
}

func (a *ActivityReportV03) AddRelatedMessageReference() *iso20022.MessageIdentification1 {
	a.RelatedMessageReference = new(iso20022.MessageIdentification1)
	return a.RelatedMessageReference
}

func (a *ActivityReportV03) AddReport() *iso20022.ActivityReportItems2 {
	newValue := new(iso20022.ActivityReportItems2)
	a.Report = append(a.Report, newValue)
	return newValue
}
