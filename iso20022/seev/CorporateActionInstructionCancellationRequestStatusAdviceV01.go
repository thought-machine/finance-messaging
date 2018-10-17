package seev

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document04100101 struct {
	XMLName xml.Name                                                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.01 Document"`
	Message *CorporateActionInstructionCancellationRequestStatusAdviceV01 `xml:"CorpActnInstrCxlReqStsAdvc"`
}

func (d *Document04100101) AddMessage() *CorporateActionInstructionCancellationRequestStatusAdviceV01 {
	d.Message = new(CorporateActionInstructionCancellationRequestStatusAdviceV01)
	return d.Message
}

// Scope
// An account servicer sends the CorporateActionInstructionCancellationRequestStatusAdvice message to an account owner or its designated agent to report status of a previously received CorporateActionInstructionCancellationRequest message sent by the account owner.
// This will include the acknowledgement/rejection of a request to cancel an outstanding instruction.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate).
// ISO 15022 - 20022 COEXISTENCE
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type CorporateActionInstructionCancellationRequestStatusAdviceV01 struct {

	// Information that unambiguously identifies a CorporateActionInstructionCancellationRequestStatusAdvice message as know by the account servicer.
	Identification *iso20022.DocumentIdentification11 `xml:"Id"`

	// Identification of a related instruction cancellation request document.
	InstructionCancellationRequestIdentification *iso20022.DocumentIdentification9 `xml:"InstrCxlReqId,omitempty"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*iso20022.DocumentIdentification14 `xml:"OthrDocId,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionGeneralInformation9 `xml:"CorpActnGnlInf"`

	// Provides information about the processing status of the instruction cancellation request.
	InstructionCancellationRequestStatus []*iso20022.InstructionCancellationRequestStatus1Choice `xml:"InstrCxlReqSts"`

	// Information about the corporate action option.
	CorporateActionInstruction *iso20022.CorporateActionOption9 `xml:"CorpActnInstr,omitempty"`

	// Provides additional information.
	AdditionalInformation *iso20022.CorporateActionNarrative10 `xml:"AddtlInf,omitempty"`

	// Party that originated the message, if other than the sender.
	MessageOriginator *iso20022.PartyIdentification10Choice `xml:"MsgOrgtr,omitempty"`

	// Party that is the final destination of the message, if other than the receiver.
	MessageRecipient *iso20022.PartyIdentification10Choice `xml:"MsgRcpt,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*iso20022.Extension2 `xml:"Xtnsn,omitempty"`
}

func (c *CorporateActionInstructionCancellationRequestStatusAdviceV01) AddIdentification() *iso20022.DocumentIdentification11 {
	c.Identification = new(iso20022.DocumentIdentification11)
	return c.Identification
}

func (c *CorporateActionInstructionCancellationRequestStatusAdviceV01) AddInstructionCancellationRequestIdentification() *iso20022.DocumentIdentification9 {
	c.InstructionCancellationRequestIdentification = new(iso20022.DocumentIdentification9)
	return c.InstructionCancellationRequestIdentification
}

func (c *CorporateActionInstructionCancellationRequestStatusAdviceV01) AddOtherDocumentIdentification() *iso20022.DocumentIdentification14 {
	newValue := new(iso20022.DocumentIdentification14)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionInstructionCancellationRequestStatusAdviceV01) AddCorporateActionGeneralInformation() *iso20022.CorporateActionGeneralInformation9 {
	c.CorporateActionGeneralInformation = new(iso20022.CorporateActionGeneralInformation9)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionInstructionCancellationRequestStatusAdviceV01) AddInstructionCancellationRequestStatus() *iso20022.InstructionCancellationRequestStatus1Choice {
	newValue := new(iso20022.InstructionCancellationRequestStatus1Choice)
	c.InstructionCancellationRequestStatus = append(c.InstructionCancellationRequestStatus, newValue)
	return newValue
}

func (c *CorporateActionInstructionCancellationRequestStatusAdviceV01) AddCorporateActionInstruction() *iso20022.CorporateActionOption9 {
	c.CorporateActionInstruction = new(iso20022.CorporateActionOption9)
	return c.CorporateActionInstruction
}

func (c *CorporateActionInstructionCancellationRequestStatusAdviceV01) AddAdditionalInformation() *iso20022.CorporateActionNarrative10 {
	c.AdditionalInformation = new(iso20022.CorporateActionNarrative10)
	return c.AdditionalInformation
}

func (c *CorporateActionInstructionCancellationRequestStatusAdviceV01) AddMessageOriginator() *iso20022.PartyIdentification10Choice {
	c.MessageOriginator = new(iso20022.PartyIdentification10Choice)
	return c.MessageOriginator
}

func (c *CorporateActionInstructionCancellationRequestStatusAdviceV01) AddMessageRecipient() *iso20022.PartyIdentification10Choice {
	c.MessageRecipient = new(iso20022.PartyIdentification10Choice)
	return c.MessageRecipient
}

func (c *CorporateActionInstructionCancellationRequestStatusAdviceV01) AddExtension() *iso20022.Extension2 {
	newValue := new(iso20022.Extension2)
	c.Extension = append(c.Extension, newValue)
	return newValue
}
