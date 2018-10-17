package camt

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document05800102 struct {
	XMLName xml.Name                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.058.001.02 Document"`
	Message *NotificationToReceiveCancellationAdviceV02 `xml:"NtfctnToRcvCxlAdvc"`
}

func (d *Document05800102) AddMessage() *NotificationToReceiveCancellationAdviceV02 {
	d.Message = new(NotificationToReceiveCancellationAdviceV02)
	return d.Message
}

// Scope
// The NotificationToReceiveCancellationAdvice message is sent by an account owner or by a party acting on the account owner's behalf to one of the account owner's account servicing institutions. It is used to advise the account servicing institution about the cancellation of one or more notifications in a previous NotificationToReceive message.
// Usage
// The NotificationToReceiveCancellationAdvice message is used to advise the account servicing institution that the funds are no longer expected. The message can be used in either a direct or a relay scenario.
type NotificationToReceiveCancellationAdviceV02 struct {

	// Set of elements used to provide further details on the message.
	GroupHeader *iso20022.GroupHeader43 `xml:"GrpHdr"`

	// Set of elements used to identify the original notification, to which the cancellation advice refers.
	OriginalNotification *iso20022.OriginalNotification4 `xml:"OrgnlNtfctn"`
}

func (n *NotificationToReceiveCancellationAdviceV02) AddGroupHeader() *iso20022.GroupHeader43 {
	n.GroupHeader = new(iso20022.GroupHeader43)
	return n.GroupHeader
}

func (n *NotificationToReceiveCancellationAdviceV02) AddOriginalNotification() *iso20022.OriginalNotification4 {
	n.OriginalNotification = new(iso20022.OriginalNotification4)
	return n.OriginalNotification
}
