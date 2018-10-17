package auth

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document01100101 struct {
	XMLName xml.Name                                          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.011.001.01 Document"`
	Message *RegulatoryTransactionReportCancellationStatusV01 `xml:"RgltryTxRptCxlStsV01"`
}

func (d *Document01100101) AddMessage() *RegulatoryTransactionReportCancellationStatusV01 {
	d.Message = new(RegulatoryTransactionReportCancellationStatusV01)
	return d.Message
}

// Scope
// A regulator or an intermediary sends the RegulatoryTransactionReportCancellationStatus to a reporting institution to provide the status of a RegulatoryTransactionReportCancellationRequest previously sent by the reporting institution.
// Usage
// The message definition may be used to provide a status for the entire report or to provide a status at the level of individual transactions within the report. One of the following statuses can be reported:
// - Completed, or,
// - Pending, or,
// - Rejected.
// If the status is rejected, then reason for the rejection must be specified.
type RegulatoryTransactionReportCancellationStatusV01 struct {

	// Identification of the RegulatoryTransactionReportCancellationStatus document.
	Identification *iso20022.DocumentIdentification8 `xml:"Id"`

	// Identification of the firm that executed the transaction.
	ReportingInstitution *iso20022.PartyIdentification23Choice `xml:"RptgInstn"`

	// Provides the status of the entire RegulatoryTransactionReportCancellationRequest document that was previously sent by a reporting institution.
	//
	//
	ReportCancellationStatus *iso20022.ReportStatusAndReason2 `xml:"RptCxlSts"`

	// Provides the cancellation status of one or more transactions within a RegulatoryTransactionReportCancellationRequest that was previously sent by a reporting institution.
	IndividualTransactionCancellationStatus []*iso20022.TradeTransactionStatusAndReason2 `xml:"IndvTxCxlSts"`
}

func (r *RegulatoryTransactionReportCancellationStatusV01) AddIdentification() *iso20022.DocumentIdentification8 {
	r.Identification = new(iso20022.DocumentIdentification8)
	return r.Identification
}

func (r *RegulatoryTransactionReportCancellationStatusV01) AddReportingInstitution() *iso20022.PartyIdentification23Choice {
	r.ReportingInstitution = new(iso20022.PartyIdentification23Choice)
	return r.ReportingInstitution
}

func (r *RegulatoryTransactionReportCancellationStatusV01) AddReportCancellationStatus() *iso20022.ReportStatusAndReason2 {
	r.ReportCancellationStatus = new(iso20022.ReportStatusAndReason2)
	return r.ReportCancellationStatus
}

func (r *RegulatoryTransactionReportCancellationStatusV01) AddIndividualTransactionCancellationStatus() *iso20022.TradeTransactionStatusAndReason2 {
	newValue := new(iso20022.TradeTransactionStatusAndReason2)
	r.IndividualTransactionCancellationStatus = append(r.IndividualTransactionCancellationStatus, newValue)
	return newValue
}
