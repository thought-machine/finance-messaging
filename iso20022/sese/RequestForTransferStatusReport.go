package sese

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document00900101 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.009.001.01 Document"`
	Message *RequestForTransferStatusReport `xml:"sese.009.001.01"`
}

func (d *Document00900101) AddMessage() *RequestForTransferStatusReport {
	d.Message = new(RequestForTransferStatusReport)
	return d.Message
}

// Scope
// The RequestForTransferStatusReport is sent by an instructing party to the executing party.
// This message requests the status of a transfer instruction or the status of a transfer cancellation instruction.
// Usage
// The RequestForTransferStatusReport is sent by an instructing party to the executing party to request
// - the status of one or several transfer instructions or
// - the status of one or several transfer cancellation instructions.
type RequestForTransferStatusReport struct {

	// Information to identify the transfer for which the status is requested.
	//
	RequestDetails []*iso20022.MessageAndBusinessReference1 `xml:"ReqDtls"`
}

func (r *RequestForTransferStatusReport) AddRequestDetails() *iso20022.MessageAndBusinessReference1 {
	newValue := new(iso20022.MessageAndBusinessReference1)
	r.RequestDetails = append(r.RequestDetails, newValue)
	return newValue
}
