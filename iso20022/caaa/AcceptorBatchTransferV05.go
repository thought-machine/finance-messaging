package caaa

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document01100105 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.011.001.05 Document"`
	Message *AcceptorBatchTransferV05 `xml:"AccptrBtchTrf"`
}

func (d *Document01100105) AddMessage() *AcceptorBatchTransferV05 {
	d.Message = new(AcceptorBatchTransferV05)
	return d.Message
}

// The AcceptorBatchTransfer is sent by an acceptor (or its agent) to transfer the  financial data of a collection of transactions to the acquirer (or its agent).
type AcceptorBatchTransferV05 struct {

	// Batch capture message management information.
	Header *iso20022.Header25 `xml:"Hdr"`

	// Card payment transactions from one or several data set of transactions.
	BatchTransfer *iso20022.CardPaymentBatchTransfer4 `xml:"BtchTrf"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *iso20022.ContentInformationType12 `xml:"SctyTrlr,omitempty"`
}

func (a *AcceptorBatchTransferV05) AddHeader() *iso20022.Header25 {
	a.Header = new(iso20022.Header25)
	return a.Header
}

func (a *AcceptorBatchTransferV05) AddBatchTransfer() *iso20022.CardPaymentBatchTransfer4 {
	a.BatchTransfer = new(iso20022.CardPaymentBatchTransfer4)
	return a.BatchTransfer
}

func (a *AcceptorBatchTransferV05) AddSecurityTrailer() *iso20022.ContentInformationType12 {
	a.SecurityTrailer = new(iso20022.ContentInformationType12)
	return a.SecurityTrailer
}
