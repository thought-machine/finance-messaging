package camt

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document06300102 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.063.001.02 Document"`
	Message *PayInEventAcknowledgementV02 `xml:"PayInEvtAck"`
}

func (d *Document06300102) AddMessage() *PayInEventAcknowledgementV02 {
	d.Message = new(PayInEventAcknowledgementV02)
	return d.Message
}

// The PayInEventAcknowledgement message is sent by a participant of a central system to the central system to confirm a PayInSchedule or a PayInCall has been received.
//
type PayInEventAcknowledgementV02 struct {

	// Unique and unambiguous identifier for the message, as assigned by the sender.
	MessageIdentification *iso20022.Max35Text `xml:"MsgId"`

	// To indicate the requested CLS Settlement Session that the related trade is part of.
	SettlementSessionIdentifier *iso20022.Exact4AlphaNumericText `xml:"SttlmSsnIdr,omitempty"`

	// Details of the pay in schedule or pay in call being acknowledged .
	AcknowledgementDetails *iso20022.AcknowledgementDetails1Choice `xml:"AckDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PayInEventAcknowledgementV02) SetMessageIdentification(value string) {
	p.MessageIdentification = (*iso20022.Max35Text)(&value)
}

func (p *PayInEventAcknowledgementV02) SetSettlementSessionIdentifier(value string) {
	p.SettlementSessionIdentifier = (*iso20022.Exact4AlphaNumericText)(&value)
}

func (p *PayInEventAcknowledgementV02) AddAcknowledgementDetails() *iso20022.AcknowledgementDetails1Choice {
	p.AcknowledgementDetails = new(iso20022.AcknowledgementDetails1Choice)
	return p.AcknowledgementDetails
}

func (p *PayInEventAcknowledgementV02) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}
