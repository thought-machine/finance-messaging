package trea

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document01300101 struct {
	XMLName xml.Name                   `xml:"urn:iso:std:iso:20022:tech:xsd:trea.013.001.01 Document"`
	Message *WithdrawalNotificationV01 `xml:"WdrwlNtfctnV01"`
}

func (d *Document01300101) AddMessage() *WithdrawalNotificationV01 {
	d.Message = new(WithdrawalNotificationV01)
	return d.Message
}

// Scope
// The WithdrawalNotification message is sent by a central system to notify the withdrawal of a trade which was previously notified to the receiver as an alleged trade.
// Usage
// The message is used to confirm the cancellation of a previously notified trade.
type WithdrawalNotificationV01 struct {

	// Reference assigned by the central matching system which is notifying the deletion of a previously reported trade.
	MatchingSystemUniqueReference *iso20022.MessageReference `xml:"MtchgSysUnqRef"`
}

func (w *WithdrawalNotificationV01) AddMatchingSystemUniqueReference() *iso20022.MessageReference {
	w.MatchingSystemUniqueReference = new(iso20022.MessageReference)
	return w.MatchingSystemUniqueReference
}
