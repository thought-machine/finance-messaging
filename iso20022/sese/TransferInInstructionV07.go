package sese

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document00500107 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.07 Document"`
	Message *TransferInInstructionV07 `xml:"TrfInInstr"`
}

func (d *Document00500107) AddMessage() *TransferInInstructionV07 {
	d.Message = new(TransferInInstructionV07)
	return d.Message
}

// Scope
// An instructing party, for example, an investment manager or its authorised representative, sends the TransferInInstruction message to the executing party, for example, a transfer agent, to instruct the receipt of a financial instrument, free of payment, on a given date from a specified party.
// This message may also be used to instruct the receipt of a financial instrument, free of payment, from another of the instructing parties own accounts or from a third party.
// This message may also be used as an advice and request, that is, the message is used to inform the receiver to expect an unsolicited transfer in confirmation and to request account information for the transfer.
// Usage
// The TransferInInstruction message is used to instruct the receipt of a financial instrument from another account, either owned by the instructing party or by a third party.
type TransferInInstructionV07 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference6 `xml:"PoolRef,omitempty"`

	// Reference of the linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference6 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference6 `xml:"RltdRef,omitempty"`

	// Function of the transfer in, that is, whether the transfer in message is used as an instruction or an advice and request for information. The absence of Function indicates the message is an instruction.
	Function *iso20022.TransferInFunction1Code `xml:"Fctn,omitempty"`

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *iso20022.Max35Text `xml:"MstrRef,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails []*iso20022.Transfer32 `xml:"TrfDtls"`

	// Information related to the account into which the financial instrument is to be received.
	AccountDetails *iso20022.InvestmentAccount56 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *iso20022.DeliverInformation16 `xml:"SttlmDtls,omitempty"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *iso20022.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation4 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferInInstructionV07) AddMessageIdentification() *iso20022.MessageIdentification1 {
	t.MessageIdentification = new(iso20022.MessageIdentification1)
	return t.MessageIdentification
}

func (t *TransferInInstructionV07) AddPoolReference() *iso20022.AdditionalReference6 {
	t.PoolReference = new(iso20022.AdditionalReference6)
	return t.PoolReference
}

func (t *TransferInInstructionV07) AddPreviousReference() *iso20022.AdditionalReference6 {
	t.PreviousReference = new(iso20022.AdditionalReference6)
	return t.PreviousReference
}

func (t *TransferInInstructionV07) AddRelatedReference() *iso20022.AdditionalReference6 {
	t.RelatedReference = new(iso20022.AdditionalReference6)
	return t.RelatedReference
}

func (t *TransferInInstructionV07) SetFunction(value string) {
	t.Function = (*iso20022.TransferInFunction1Code)(&value)
}

func (t *TransferInInstructionV07) SetMasterReference(value string) {
	t.MasterReference = (*iso20022.Max35Text)(&value)
}

func (t *TransferInInstructionV07) AddTransferDetails() *iso20022.Transfer32 {
	newValue := new(iso20022.Transfer32)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferInInstructionV07) AddAccountDetails() *iso20022.InvestmentAccount56 {
	t.AccountDetails = new(iso20022.InvestmentAccount56)
	return t.AccountDetails
}

func (t *TransferInInstructionV07) AddSettlementDetails() *iso20022.DeliverInformation16 {
	t.SettlementDetails = new(iso20022.DeliverInformation16)
	return t.SettlementDetails
}

func (t *TransferInInstructionV07) AddMarketPracticeVersion() *iso20022.MarketPracticeVersion1 {
	t.MarketPracticeVersion = new(iso20022.MarketPracticeVersion1)
	return t.MarketPracticeVersion
}

func (t *TransferInInstructionV07) AddCopyDetails() *iso20022.CopyInformation4 {
	t.CopyDetails = new(iso20022.CopyInformation4)
	return t.CopyDetails
}

func (t *TransferInInstructionV07) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
