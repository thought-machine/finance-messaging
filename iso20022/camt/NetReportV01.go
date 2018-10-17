package camt

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document08800101 struct {
	XMLName xml.Name      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.088.001.01 Document"`
	Message *NetReportV01 `xml:"NetRpt"`
}

func (d *Document08800101) AddMessage() *NetReportV01 {
	d.Message = new(NetReportV01)
	return d.Message
}

// The Net Report message is sent to a participant by a central system to provide details of the of the bi-lateral payment obligations, calculated by the central system per currency.
type NetReportV01 struct {

	// Specifies the meta data associated with the net report.
	NetReportData *iso20022.NetReportData1 `xml:"NetRptData"`

	// Describes the participant receiving the net report.
	NetServiceParticipantIdentification *iso20022.PartyIdentification73Choice `xml:"NetSvcPtcptId"`

	// Describes the counterparty participant involved in (all of) the obligations of the report.
	NetServiceCounterpartyIdentification *iso20022.PartyIdentification73Choice `xml:"NetSvcCtrPtyId,omitempty"`

	// Provides the amount, direct parties or netting groups involved in the obligation.
	NetObligation []*iso20022.NetObligation1 `xml:"NetOblgtn"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (n *NetReportV01) AddNetReportData() *iso20022.NetReportData1 {
	n.NetReportData = new(iso20022.NetReportData1)
	return n.NetReportData
}

func (n *NetReportV01) AddNetServiceParticipantIdentification() *iso20022.PartyIdentification73Choice {
	n.NetServiceParticipantIdentification = new(iso20022.PartyIdentification73Choice)
	return n.NetServiceParticipantIdentification
}

func (n *NetReportV01) AddNetServiceCounterpartyIdentification() *iso20022.PartyIdentification73Choice {
	n.NetServiceCounterpartyIdentification = new(iso20022.PartyIdentification73Choice)
	return n.NetServiceCounterpartyIdentification
}

func (n *NetReportV01) AddNetObligation() *iso20022.NetObligation1 {
	newValue := new(iso20022.NetObligation1)
	n.NetObligation = append(n.NetObligation, newValue)
	return newValue
}

func (n *NetReportV01) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	n.SupplementaryData = append(n.SupplementaryData, newValue)
	return newValue
}
