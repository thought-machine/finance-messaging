package colr

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document00600104 struct {
	XMLName xml.Name                                   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.006.001.04 Document"`
	Message *CollateralManagementCancellationStatusV04 `xml:"CollMgmtCxlSts"`
}

func (d *Document00600104) AddMessage() *CollateralManagementCancellationStatusV04 {
	d.Message = new(CollateralManagementCancellationStatusV04)
	return d.Message
}

// Scope
// The CollateralManagementCancellationStatus message is sent by:
// - the collateral taker or its collateral manager to the collateral giver or its collateral manager, or
// - the collateral giver or its collateral manager to the collateral taker or its collateral manager.
// This message is used to provide the status of accepted or rejected for the CollateralManagementCancellationRequest message previously received.
//
// The message definition is intended for use with the ISO20022 Business Application Header.
//
// Usage
// The CollateralManagementCancellationStatus message can be sent to provide a status for a previously received CollateralManagementCancellationRequest message.
type CollateralManagementCancellationStatusV04 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *iso20022.Max35Text `xml:"TxId"`

	// Provides reference to the previous received message.
	Reference *iso20022.Reference16 `xml:"Ref"`

	// Provides information like the identification of the party or parties associated with the collateral agreement, the exposure type and the valuation date.
	Obligation *iso20022.Obligation4 `xml:"Oblgtn"`

	// Provides status details of the collateral cancellation message.
	CancellationStatus *iso20022.CollateralCancellationStatus1 `xml:"CxlSts"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CollateralManagementCancellationStatusV04) SetTransactionIdentification(value string) {
	c.TransactionIdentification = (*iso20022.Max35Text)(&value)
}

func (c *CollateralManagementCancellationStatusV04) AddReference() *iso20022.Reference16 {
	c.Reference = new(iso20022.Reference16)
	return c.Reference
}

func (c *CollateralManagementCancellationStatusV04) AddObligation() *iso20022.Obligation4 {
	c.Obligation = new(iso20022.Obligation4)
	return c.Obligation
}

func (c *CollateralManagementCancellationStatusV04) AddCancellationStatus() *iso20022.CollateralCancellationStatus1 {
	c.CancellationStatus = new(iso20022.CollateralCancellationStatus1)
	return c.CancellationStatus
}

func (c *CollateralManagementCancellationStatusV04) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
