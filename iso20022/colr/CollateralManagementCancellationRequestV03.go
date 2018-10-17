package colr

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document00500103 struct {
	XMLName xml.Name                                    `xml:"urn:iso:std:iso:20022:tech:xsd:colr.005.001.03 Document"`
	Message *CollateralManagementCancellationRequestV03 `xml:"CollMgmtCxlReq"`
}

func (d *Document00500103) AddMessage() *CollateralManagementCancellationRequestV03 {
	d.Message = new(CollateralManagementCancellationRequestV03)
	return d.Message
}

// Scope
// This CollateralManagementCancellationRequest message is sent by:
// - the collateral taker or its collateral manager to the collateral giver or its collateral manager,
// - the collateral giver or its collateral manager to the collateral taker or its collateral manager
// This message is used to request the cancellation of a previously sent MarginCallRequest message, MarginCallResponse message, CollateralProposal message, CollateralProposalResponse message, MarginCallDisputeNotification message or a CollateralSubstitutionRequest message.
//
// The message definition is intended for use with the ISO20022 Business Application Header.
//
// Usage
// The CollateralManagementCancellationRequest message is used to request the cancellation of a collateral message. When requesting the cancellation of a message there must be a cancellation reason specified.
// When the CollateralManagementCancellationRequest message is used to cancel a collateral message the reference of the original message must be specified. The rejection or acceptance of a CollateralManagementCancellationRequest message is made using a CollateralManagementCancellationStatus message.
type CollateralManagementCancellationRequestV03 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *iso20022.Max35Text `xml:"TxId"`

	// Reference to the message advised to be cancelled.
	Reference *iso20022.Reference2Choice `xml:"Ref"`

	// Provides information like the identification of the party or parties associated with the collateral agreement, the exposure type and the valuation date.
	Obligation *iso20022.Obligation3 `xml:"Oblgtn"`

	// It is used to detail the reason for the cancellation of a previously sent request.
	CancellationReason *iso20022.CollateralCancellationReason1 `xml:"CxlRsn"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CollateralManagementCancellationRequestV03) SetTransactionIdentification(value string) {
	c.TransactionIdentification = (*iso20022.Max35Text)(&value)
}

func (c *CollateralManagementCancellationRequestV03) AddReference() *iso20022.Reference2Choice {
	c.Reference = new(iso20022.Reference2Choice)
	return c.Reference
}

func (c *CollateralManagementCancellationRequestV03) AddObligation() *iso20022.Obligation3 {
	c.Obligation = new(iso20022.Obligation3)
	return c.Obligation
}

func (c *CollateralManagementCancellationRequestV03) AddCancellationReason() *iso20022.CollateralCancellationReason1 {
	c.CancellationReason = new(iso20022.CollateralCancellationReason1)
	return c.CancellationReason
}

func (c *CollateralManagementCancellationRequestV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
