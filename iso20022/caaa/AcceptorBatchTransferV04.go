package caaa

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document01100104 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.011.001.04 Document"`
	Message *AcceptorBatchTransferV04 `xml:"AccptrBtchTrf"`
}

func (d *Document01100104) AddMessage() *AcceptorBatchTransferV04 {
	d.Message = new(AcceptorBatchTransferV04)
	return d.Message
}

// The AcceptorBatchTransfer is sent by an acceptor (or its agent) to transfer the  financial data of a collection of transactions to the acquirer (or its agent).
type AcceptorBatchTransferV04 struct {

	// Batch capture message management information.
	Header *iso20022.Header12 `xml:"Hdr"`

	// Card payment transactions from one or several data set of transactions.
	BatchTransfer *iso20022.CardPaymentBatchTransfer3 `xml:"BtchTrf"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *iso20022.ContentInformationType12 `xml:"SctyTrlr"`
}

func (a *AcceptorBatchTransferV04) AddHeader() *iso20022.Header12 {
	a.Header = new(iso20022.Header12)
	return a.Header
}

func (a *AcceptorBatchTransferV04) AddBatchTransfer() *iso20022.CardPaymentBatchTransfer3 {
	a.BatchTransfer = new(iso20022.CardPaymentBatchTransfer3)
	return a.BatchTransfer
}

func (a *AcceptorBatchTransferV04) AddSecurityTrailer() *iso20022.ContentInformationType12 {
	a.SecurityTrailer = new(iso20022.ContentInformationType12)
	return a.SecurityTrailer
}
