package fxtr

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document03000104 struct {
	XMLName xml.Name                                       `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.030.001.04 Document"`
	Message *ForeignExchangeTradeBulkStatusNotificationV04 `xml:"FXTradBlkStsNtfctn"`
}

func (d *Document03000104) AddMessage() *ForeignExchangeTradeBulkStatusNotificationV04 {
	d.Message = new(ForeignExchangeTradeBulkStatusNotificationV04)
	return d.Message
}

// Scope
// The ForeignExchangeTradeBulkStatusNotification message is sent by a central system to the participant to provide notification of the current status of one or more foreign exchange trades.
type ForeignExchangeTradeBulkStatusNotificationV04 struct {

	// Information on the status of the trade in the central system.
	//
	StatusDetails *iso20022.TradeData12 `xml:"StsDtls"`

	// Identifies one or more trades for which the status notification is sent.
	//
	TradeData []*iso20022.TradeData11 `xml:"TradData"`

	// Page number of the message (within the status report) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the report.
	MessagePagination *iso20022.Pagination `xml:"MsgPgntn,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (f *ForeignExchangeTradeBulkStatusNotificationV04) AddStatusDetails() *iso20022.TradeData12 {
	f.StatusDetails = new(iso20022.TradeData12)
	return f.StatusDetails
}

func (f *ForeignExchangeTradeBulkStatusNotificationV04) AddTradeData() *iso20022.TradeData11 {
	newValue := new(iso20022.TradeData11)
	f.TradeData = append(f.TradeData, newValue)
	return newValue
}

func (f *ForeignExchangeTradeBulkStatusNotificationV04) AddMessagePagination() *iso20022.Pagination {
	f.MessagePagination = new(iso20022.Pagination)
	return f.MessagePagination
}

func (f *ForeignExchangeTradeBulkStatusNotificationV04) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	f.SupplementaryData = append(f.SupplementaryData, newValue)
	return newValue
}
