package trea

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document00700102 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:trea.007.001.02 Document"`
	Message *NonDeliverableForwardNotificationV02 `xml:"NDFNtfctnV02"`
}

func (d *Document00700102) AddMessage() *NonDeliverableForwardNotificationV02 {
	d.Message = new(NonDeliverableForwardNotificationV02)
	return d.Message
}

// Scope
// The NonDeliverableForwardNotification message is sent by a central system to a participant to provide details of a non deliverable forward trade.
// Usage
// The notification is sent by a central settlement system to the two trading parties after it has received create, amend or cancel messages from both. The message may also contain information on the settlement of the trade.
type NonDeliverableForwardNotificationV02 struct {

	// Specifies the trading side of the non deliverable trade which is reported.
	TradingSideIdentification *iso20022.TradePartyIdentification3 `xml:"TradgSdId"`

	// Specifies the counterparty of the non deliverable trade which is reported.
	CounterpartySideIdentification *iso20022.TradePartyIdentification3 `xml:"CtrPtySdId"`

	// Set of data specified for the opening of a non deliverable trade.
	OpeningData *iso20022.OpeningData2 `xml:"OpngData,omitempty"`

	// Set of data specified for the valuation of a non deliverable trade.
	ValuationData *iso20022.ClosingData2 `xml:"ValtnData,omitempty"`

	// Provides information on the status of a trade in a central system.
	TradeInformationAndStatus *iso20022.TradeStatus1 `xml:"TradInfAndSts"`

	// Provides information on the settlement of a trade.
	SettlementData *iso20022.SettlementData2 `xml:"SttlmData,omitempty"`
}

func (n *NonDeliverableForwardNotificationV02) AddTradingSideIdentification() *iso20022.TradePartyIdentification3 {
	n.TradingSideIdentification = new(iso20022.TradePartyIdentification3)
	return n.TradingSideIdentification
}

func (n *NonDeliverableForwardNotificationV02) AddCounterpartySideIdentification() *iso20022.TradePartyIdentification3 {
	n.CounterpartySideIdentification = new(iso20022.TradePartyIdentification3)
	return n.CounterpartySideIdentification
}

func (n *NonDeliverableForwardNotificationV02) AddOpeningData() *iso20022.OpeningData2 {
	n.OpeningData = new(iso20022.OpeningData2)
	return n.OpeningData
}

func (n *NonDeliverableForwardNotificationV02) AddValuationData() *iso20022.ClosingData2 {
	n.ValuationData = new(iso20022.ClosingData2)
	return n.ValuationData
}

func (n *NonDeliverableForwardNotificationV02) AddTradeInformationAndStatus() *iso20022.TradeStatus1 {
	n.TradeInformationAndStatus = new(iso20022.TradeStatus1)
	return n.TradeInformationAndStatus
}

func (n *NonDeliverableForwardNotificationV02) AddSettlementData() *iso20022.SettlementData2 {
	n.SettlementData = new(iso20022.SettlementData2)
	return n.SettlementData
}
