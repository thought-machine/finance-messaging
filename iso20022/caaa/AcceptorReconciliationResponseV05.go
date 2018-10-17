package caaa

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document01000105 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.010.001.05 Document"`
	Message *AcceptorReconciliationResponseV05 `xml:"AccptrRcncltnRspn"`
}

func (d *Document01000105) AddMessage() *AcceptorReconciliationResponseV05 {
	d.Message = new(AcceptorReconciliationResponseV05)
	return d.Message
}

// The AcceptorReconciliationResponse message is sent by the acquirer (or its agent) to an acceptor (or its agent), to ensure that the debits and credits performed by the acceptor matches the computed balances of the acquirer for the debits and credits performed during the same reconciliation period.
// If the acceptor or the acquirer notices a difference in totals, the discrepancy will be resolved by other means, outside the scope of the protocol.
type AcceptorReconciliationResponseV05 struct {

	// Reconciliation response message management information.
	Header *iso20022.Header30 `xml:"Hdr"`

	// Information related to the reconciliation response.
	ReconciliationResponse *iso20022.AcceptorReconciliationResponse4 `xml:"RcncltnRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *AcceptorReconciliationResponseV05) AddHeader() *iso20022.Header30 {
	a.Header = new(iso20022.Header30)
	return a.Header
}

func (a *AcceptorReconciliationResponseV05) AddReconciliationResponse() *iso20022.AcceptorReconciliationResponse4 {
	a.ReconciliationResponse = new(iso20022.AcceptorReconciliationResponse4)
	return a.ReconciliationResponse
}

func (a *AcceptorReconciliationResponseV05) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}
