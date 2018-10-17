package colr

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document01400104 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:colr.014.001.04 Document"`
	Message *InterestPaymentResponseV04 `xml:"IntrstPmtRspn"`
}

func (d *Document01400104) AddMessage() *InterestPaymentResponseV04 {
	d.Message = new(InterestPaymentResponseV04)
	return d.Message
}

// Scope
// The InterestPaymentResponse message is sent by either;
// - the collateral taker or its collateral manager to the collateral giver or its collateral manager, or
// - the collateral giver or its collateral manager to the collateral taker or its collateral manager
// This is a response to the InterestPaymentRequest message and the amount of interest requested or advised can be accepted or rejected.
//
// The message definition is intended for use with the ISO20022 Business Application Header.
//
// Usage
// The InterestPaymentResponse message is sent in response to the InterestPaymentRequest in order to accept or reject the amount of interest requested or advised. A rejection reason and information can be provide if the InterestPaymentRequest is being rejected.
type InterestPaymentResponseV04 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *iso20022.Max35Text `xml:"TxId"`

	// Provides information like the identification of the party or parties associated with the collateral agreement, the exposure type and the valuation date.
	Obligation *iso20022.Obligation4 `xml:"Oblgtn"`

	// Agreement details for the over the counter market.
	Agreement *iso20022.Agreement4 `xml:"Agrmt"`

	// Provides details on the interest amount due to party A.
	InterestDueToA *iso20022.InterestAmount2 `xml:"IntrstDueToA,omitempty"`

	// Provides details on the interest amount due to party B.
	InterestDueToB *iso20022.InterestAmount2 `xml:"IntrstDueToB,omitempty"`

	// Provides details on the response to the interest payment request.
	InterestResponse *iso20022.InterestResponse1 `xml:"IntrstRspn"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *InterestPaymentResponseV04) SetTransactionIdentification(value string) {
	i.TransactionIdentification = (*iso20022.Max35Text)(&value)
}

func (i *InterestPaymentResponseV04) AddObligation() *iso20022.Obligation4 {
	i.Obligation = new(iso20022.Obligation4)
	return i.Obligation
}

func (i *InterestPaymentResponseV04) AddAgreement() *iso20022.Agreement4 {
	i.Agreement = new(iso20022.Agreement4)
	return i.Agreement
}

func (i *InterestPaymentResponseV04) AddInterestDueToA() *iso20022.InterestAmount2 {
	i.InterestDueToA = new(iso20022.InterestAmount2)
	return i.InterestDueToA
}

func (i *InterestPaymentResponseV04) AddInterestDueToB() *iso20022.InterestAmount2 {
	i.InterestDueToB = new(iso20022.InterestAmount2)
	return i.InterestDueToB
}

func (i *InterestPaymentResponseV04) AddInterestResponse() *iso20022.InterestResponse1 {
	i.InterestResponse = new(iso20022.InterestResponse1)
	return i.InterestResponse
}

func (i *InterestPaymentResponseV04) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}
