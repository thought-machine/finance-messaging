package camt

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document05900103 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.03 Document"`
	Message *NotificationToReceiveStatusReportV03 `xml:"NtfctnToRcvStsRpt"`
}

func (d *Document05900103) AddMessage() *NotificationToReceiveStatusReportV03 {
	d.Message = new(NotificationToReceiveStatusReportV03)
	return d.Message
}

// Scope
// The NotificationToReceiveStatusReport message is sent by an account servicing institution to an account owner or to a party acting on the account owner's behalf. It is used to notify the account owner about the status of one or more expected payments that were advised in a previous NotificationToReceive message.
// Usage
// The NotificationToReceiveStatusReport message is sent in response to a NotificationToReceive message and can be used in either a direct or a relay scenario. It is used to advise the account owner of receipt of the funds as expected. It is also used to notify the account owner of non-receipt of funds or of discrepancies in the funds received.
type NotificationToReceiveStatusReportV03 struct {

	// Set of elements used to provide further details on the message.
	GroupHeader *iso20022.GroupHeader60 `xml:"GrpHdr"`

	// Set of elements used to identify the original notification and to provide the status.
	OriginalNotificationAndStatus *iso20022.OriginalNotification5 `xml:"OrgnlNtfctnAndSts"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (n *NotificationToReceiveStatusReportV03) AddGroupHeader() *iso20022.GroupHeader60 {
	n.GroupHeader = new(iso20022.GroupHeader60)
	return n.GroupHeader
}

func (n *NotificationToReceiveStatusReportV03) AddOriginalNotificationAndStatus() *iso20022.OriginalNotification5 {
	n.OriginalNotificationAndStatus = new(iso20022.OriginalNotification5)
	return n.OriginalNotificationAndStatus
}

func (n *NotificationToReceiveStatusReportV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	n.SupplementaryData = append(n.SupplementaryData, newValue)
	return newValue
}
