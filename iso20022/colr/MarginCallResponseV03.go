package colr

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document00400103 struct {
	XMLName xml.Name               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.004.001.03 Document"`
	Message *MarginCallResponseV03 `xml:"MrgnCallRspn"`
}

func (d *Document00400103) AddMessage() *MarginCallResponseV03 {
	d.Message = new(MarginCallResponseV03)
	return d.Message
}

// Scope
// This message is sent by the collateral giver or its collateral manager to the collateral taker or its collateral manager or vice versa. This is a response to the MarginCallRequest message. The margin call can be accepted, fully disputed or partially disputed.
//
// The message definition is intended for use with the ISO20022 Business Application Header.
//
// Usage
// When sent by the collateral giver the MarginCallResponse message is used to:
// - fully accept the MarginCallRequest
// - or partially accept the MarginCallRequest.
//
// When sent by the collateral taker the MarginCallResponse message is used to:
// - fully accept the recall of collateral
// - or partially accept the recall of collateral.
type MarginCallResponseV03 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *iso20022.Max35Text `xml:"TxId"`

	// Provides information like the identification of the party or parties associated with the collateral agreement, the exposure type and the valuation date.
	Obligation *iso20022.Obligation3 `xml:"Oblgtn"`

	// Agreement details for the over the counter market.
	Agreement *iso20022.Agreement2 `xml:"Agrmt,omitempty"`

	// Provides details about the margin calculation that would be due to party A.
	MarginDetailsDueToA *iso20022.MarginCall1 `xml:"MrgnDtlsDueToA,omitempty"`

	// Provides details about the margin calculation that would be due to party B.
	MarginDetailsDueToB *iso20022.MarginCall1 `xml:"MrgnDtlsDueToB,omitempty"`

	// Provides details about the agreed amount that would be due to party A.
	AgreedAmountDueToA *iso20022.AgreedAmount1Choice `xml:"AgrdAmtDueToA,omitempty"`

	// Provides details about the agreed amount that would be due to party B.
	AgreedAmountDueToB *iso20022.AgreedAmount1Choice `xml:"AgrdAmtDueToB,omitempty"`

	// Provides response details about the margin call.
	ResponseDetails *iso20022.Response1 `xml:"RspnDtls,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MarginCallResponseV03) SetTransactionIdentification(value string) {
	m.TransactionIdentification = (*iso20022.Max35Text)(&value)
}

func (m *MarginCallResponseV03) AddObligation() *iso20022.Obligation3 {
	m.Obligation = new(iso20022.Obligation3)
	return m.Obligation
}

func (m *MarginCallResponseV03) AddAgreement() *iso20022.Agreement2 {
	m.Agreement = new(iso20022.Agreement2)
	return m.Agreement
}

func (m *MarginCallResponseV03) AddMarginDetailsDueToA() *iso20022.MarginCall1 {
	m.MarginDetailsDueToA = new(iso20022.MarginCall1)
	return m.MarginDetailsDueToA
}

func (m *MarginCallResponseV03) AddMarginDetailsDueToB() *iso20022.MarginCall1 {
	m.MarginDetailsDueToB = new(iso20022.MarginCall1)
	return m.MarginDetailsDueToB
}

func (m *MarginCallResponseV03) AddAgreedAmountDueToA() *iso20022.AgreedAmount1Choice {
	m.AgreedAmountDueToA = new(iso20022.AgreedAmount1Choice)
	return m.AgreedAmountDueToA
}

func (m *MarginCallResponseV03) AddAgreedAmountDueToB() *iso20022.AgreedAmount1Choice {
	m.AgreedAmountDueToB = new(iso20022.AgreedAmount1Choice)
	return m.AgreedAmountDueToB
}

func (m *MarginCallResponseV03) AddResponseDetails() *iso20022.Response1 {
	m.ResponseDetails = new(iso20022.Response1)
	return m.ResponseDetails
}

func (m *MarginCallResponseV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}
