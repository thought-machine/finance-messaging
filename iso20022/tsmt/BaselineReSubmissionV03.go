package tsmt

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document01200103 struct {
	XMLName xml.Name                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.012.001.03 Document"`
	Message *BaselineReSubmissionV03 `xml:"BaselnReSubmissn"`
}

func (d *Document01200103) AddMessage() *BaselineReSubmissionV03 {
	d.Message = new(BaselineReSubmissionV03)
	return d.Message
}

// Scope
// The BaselineReSubmission message is sent by either the counterparty or the initiator of a transaction (baseline) to the matching application.
// This message is used by the counterparty to respond on the registration of a push-through transaction in the matching application or by the initiator or counterparty to re-send earlier mis-matched baseline information.
// Usage
// The BaselineReSubmission message can be sent by the counterparty of a transaction to the matching application in response to a FullPushThroughReport message received from the matching application conveying the details of an InitialBaselineSubmission message. The objective of the BaselineReSubmission message sent in the outlined scenario is to achieve a successful match of two baseline initiation messages in order to establish a transaction in the matching application.
// or
// The BaselineReSubmission message can be sent by the initiator of a transaction to the matching application in response to a BaselineMatchReport message indicating mis-matches. The objective of the BaselineReSubmission message sent in the outlined scenario is to correct an InitialBaselineSubmission or BaselineReSubmission message submitted earlier in order to achieve the establishment of a transaction in the matching application.
// or
// The BaselineReSubmission message can be sent by the counterparty of a transaction to the matching application in response to a BaselineMatchReport message indicating mis-matches. The objective of the BaselineReSubmission message sent in the outlined scenario is to correct a BaselineReSubmission message submitted earlier in order to achieve the establishment of a transaction in the matching application.
type BaselineReSubmissionV03 struct {

	// Identifies the submitted information
	SubmissionIdentification *iso20022.MessageIdentification1 `xml:"SubmissnId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *iso20022.SimpleIdentificationInformation `xml:"TxId"`

	// Reference to the transaction for the requesting financial institution.
	SubmitterTransactionReference *iso20022.SimpleIdentificationInformation `xml:"SubmitrTxRef,omitempty"`

	// Specifies the commercial details of the underlying transaction.
	Baseline *iso20022.Baseline3 `xml:"Baseln"`

	// Person to be contacted in the organisation of the buyer.
	BuyerContactPerson []*iso20022.ContactIdentification1 `xml:"BuyrCtctPrsn,omitempty"`

	// Person to be contacted in the organisation of the seller.
	SellerContactPerson []*iso20022.ContactIdentification1 `xml:"SellrCtctPrsn,omitempty"`

	// Person to be contacted in the buyer's bank.
	BuyerBankContactPerson []*iso20022.ContactIdentification1 `xml:"BuyrBkCtctPrsn"`

	// Person to be contacted in the seller's bank.
	SellerBankContactPerson []*iso20022.ContactIdentification1 `xml:"SellrBkCtctPrsn"`

	// Person to be contacted in another bank than the seller or buyer's bank.
	OtherBankContactPerson []*iso20022.ContactIdentification3 `xml:"OthrBkCtctPrsn,omitempty"`
}

func (b *BaselineReSubmissionV03) AddSubmissionIdentification() *iso20022.MessageIdentification1 {
	b.SubmissionIdentification = new(iso20022.MessageIdentification1)
	return b.SubmissionIdentification
}

func (b *BaselineReSubmissionV03) AddTransactionIdentification() *iso20022.SimpleIdentificationInformation {
	b.TransactionIdentification = new(iso20022.SimpleIdentificationInformation)
	return b.TransactionIdentification
}

func (b *BaselineReSubmissionV03) AddSubmitterTransactionReference() *iso20022.SimpleIdentificationInformation {
	b.SubmitterTransactionReference = new(iso20022.SimpleIdentificationInformation)
	return b.SubmitterTransactionReference
}

func (b *BaselineReSubmissionV03) AddBaseline() *iso20022.Baseline3 {
	b.Baseline = new(iso20022.Baseline3)
	return b.Baseline
}

func (b *BaselineReSubmissionV03) AddBuyerContactPerson() *iso20022.ContactIdentification1 {
	newValue := new(iso20022.ContactIdentification1)
	b.BuyerContactPerson = append(b.BuyerContactPerson, newValue)
	return newValue
}

func (b *BaselineReSubmissionV03) AddSellerContactPerson() *iso20022.ContactIdentification1 {
	newValue := new(iso20022.ContactIdentification1)
	b.SellerContactPerson = append(b.SellerContactPerson, newValue)
	return newValue
}

func (b *BaselineReSubmissionV03) AddBuyerBankContactPerson() *iso20022.ContactIdentification1 {
	newValue := new(iso20022.ContactIdentification1)
	b.BuyerBankContactPerson = append(b.BuyerBankContactPerson, newValue)
	return newValue
}

func (b *BaselineReSubmissionV03) AddSellerBankContactPerson() *iso20022.ContactIdentification1 {
	newValue := new(iso20022.ContactIdentification1)
	b.SellerBankContactPerson = append(b.SellerBankContactPerson, newValue)
	return newValue
}

func (b *BaselineReSubmissionV03) AddOtherBankContactPerson() *iso20022.ContactIdentification3 {
	newValue := new(iso20022.ContactIdentification3)
	b.OtherBankContactPerson = append(b.OtherBankContactPerson, newValue)
	return newValue
}
