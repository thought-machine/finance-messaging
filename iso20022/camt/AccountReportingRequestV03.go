package camt

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document06000103 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Document"`
	Message *AccountReportingRequestV03 `xml:"AcctRptgReq"`
}

func (d *Document06000103) AddMessage() *AccountReportingRequestV03 {
	d.Message = new(AccountReportingRequestV03)
	return d.Message
}

// Scope
// The AccountReportingRequest message is sent by the account owner, either directly or through a forwarding agent, to one of its account servicing institutions. It is used to ask the account servicing institution to send a report on the account owner's account in a BankToCustomerAccountReport (camt.052.001.03), a BankToCustomerStatement (camt.053.001.03) or a BankToCustomerDebitCreditNotification (camt.054.001.03).
// Usage
// The AccountReportingRequest message is used to advise the account servicing institution of funds that the account owner expects to have credited to its account. The message can be used in either a direct or a relay scenario.
type AccountReportingRequestV03 struct {

	// Set of elements used to provide further details on the message.
	GroupHeader *iso20022.GroupHeader59 `xml:"GrpHdr"`

	// Set of elements used to provide further details on the reporting request.
	ReportingRequest []*iso20022.ReportingRequest3 `xml:"RptgReq"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AccountReportingRequestV03) AddGroupHeader() *iso20022.GroupHeader59 {
	a.GroupHeader = new(iso20022.GroupHeader59)
	return a.GroupHeader
}

func (a *AccountReportingRequestV03) AddReportingRequest() *iso20022.ReportingRequest3 {
	newValue := new(iso20022.ReportingRequest3)
	a.ReportingRequest = append(a.ReportingRequest, newValue)
	return newValue
}

func (a *AccountReportingRequestV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}
