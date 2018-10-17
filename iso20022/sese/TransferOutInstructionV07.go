package sese

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document00100107 struct {
	XMLName xml.Name                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.07 Document"`
	Message *TransferOutInstructionV07 `xml:"TrfOutInstr"`
}

func (d *Document00100107) AddMessage() *TransferOutInstructionV07 {
	d.Message = new(TransferOutInstructionV07)
	return d.Message
}

// Scope
// An instructing party, for example, an investment manager or its authorised representative, sends the TransferOutInstruction message to the executing party, for example, a transfer agent, to instruct the delivery of a financial instrument, free of payment, on a given date from a specified party.
// This message may also be used to instruct the delivery of a financial instrument, free of payment, to another of the instructing parties own accounts or to a third party.
// Usage
// The TransferOutInstruction message is used to instruct the withdrawal of a financial instrument from one account and deliver it to either another account or to a third party.
type TransferOutInstructionV07 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference6 `xml:"PoolRef,omitempty"`

	// Reference of the linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference6 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference6 `xml:"RltdRef,omitempty"`

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *iso20022.Max35Text `xml:"MstrRef,omitempty"`

	// Requested date at which the instructing party places the transfer instruction.
	RequestedTransferDate *iso20022.DateFormat1Choice `xml:"ReqdTrfDt,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails []*iso20022.Transfer30 `xml:"TrfDtls"`

	// Information related to the account from which the financial instrument is to be withdrawn.
	AccountDetails *iso20022.InvestmentAccount54 `xml:"AcctDtls"`

	// Information related to the receiving side of the transfer.
	SettlementDetails *iso20022.ReceiveInformation16 `xml:"SttlmDtls,omitempty"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *iso20022.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation4 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferOutInstructionV07) AddMessageIdentification() *iso20022.MessageIdentification1 {
	t.MessageIdentification = new(iso20022.MessageIdentification1)
	return t.MessageIdentification
}

func (t *TransferOutInstructionV07) AddPoolReference() *iso20022.AdditionalReference6 {
	t.PoolReference = new(iso20022.AdditionalReference6)
	return t.PoolReference
}

func (t *TransferOutInstructionV07) AddPreviousReference() *iso20022.AdditionalReference6 {
	t.PreviousReference = new(iso20022.AdditionalReference6)
	return t.PreviousReference
}

func (t *TransferOutInstructionV07) AddRelatedReference() *iso20022.AdditionalReference6 {
	t.RelatedReference = new(iso20022.AdditionalReference6)
	return t.RelatedReference
}

func (t *TransferOutInstructionV07) SetMasterReference(value string) {
	t.MasterReference = (*iso20022.Max35Text)(&value)
}

func (t *TransferOutInstructionV07) AddRequestedTransferDate() *iso20022.DateFormat1Choice {
	t.RequestedTransferDate = new(iso20022.DateFormat1Choice)
	return t.RequestedTransferDate
}

func (t *TransferOutInstructionV07) AddTransferDetails() *iso20022.Transfer30 {
	newValue := new(iso20022.Transfer30)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferOutInstructionV07) AddAccountDetails() *iso20022.InvestmentAccount54 {
	t.AccountDetails = new(iso20022.InvestmentAccount54)
	return t.AccountDetails
}

func (t *TransferOutInstructionV07) AddSettlementDetails() *iso20022.ReceiveInformation16 {
	t.SettlementDetails = new(iso20022.ReceiveInformation16)
	return t.SettlementDetails
}

func (t *TransferOutInstructionV07) AddMarketPracticeVersion() *iso20022.MarketPracticeVersion1 {
	t.MarketPracticeVersion = new(iso20022.MarketPracticeVersion1)
	return t.MarketPracticeVersion
}

func (t *TransferOutInstructionV07) AddCopyDetails() *iso20022.CopyInformation4 {
	t.CopyDetails = new(iso20022.CopyInformation4)
	return t.CopyDetails
}

func (t *TransferOutInstructionV07) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
