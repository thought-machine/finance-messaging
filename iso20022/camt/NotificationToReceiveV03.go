package camt

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document05700103 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.057.001.03 Document"`
	Message *NotificationToReceiveV03 `xml:"NtfctnToRcv"`
}

func (d *Document05700103) AddMessage() *NotificationToReceiveV03 {
	d.Message = new(NotificationToReceiveV03)
	return d.Message
}

// Scope
// The NotificationToReceive message is sent by an account owner or by a party acting on the account owner's behalf to one of the account owner's account servicing institutions. It is an advance notice that the account servicing institution will receive funds to be credited to the account of the account owner.
// Usage
// The NotificationToReceive message is used to advise the account servicing institution of funds that the account owner expects to have credited to its account. The message can be used in either a direct or a relay scenario.
type NotificationToReceiveV03 struct {

	// Set of elements used to provide further details on the message.
	GroupHeader *iso20022.GroupHeader59 `xml:"GrpHdr"`

	// Set of elements used to provide further details on the account notification.
	Notification *iso20022.AccountNotification6 `xml:"Ntfctn"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (n *NotificationToReceiveV03) AddGroupHeader() *iso20022.GroupHeader59 {
	n.GroupHeader = new(iso20022.GroupHeader59)
	return n.GroupHeader
}

func (n *NotificationToReceiveV03) AddNotification() *iso20022.AccountNotification6 {
	n.Notification = new(iso20022.AccountNotification6)
	return n.Notification
}

func (n *NotificationToReceiveV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	n.SupplementaryData = append(n.SupplementaryData, newValue)
	return newValue
}
