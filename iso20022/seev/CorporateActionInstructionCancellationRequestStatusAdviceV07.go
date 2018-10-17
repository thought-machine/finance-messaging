package seev

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document04100107 struct {
	XMLName xml.Name                                                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.07 Document"`
	Message *CorporateActionInstructionCancellationRequestStatusAdviceV07 `xml:"CorpActnInstrCxlReqStsAdvc"`
}

func (d *Document04100107) AddMessage() *CorporateActionInstructionCancellationRequestStatusAdviceV07 {
	d.Message = new(CorporateActionInstructionCancellationRequestStatusAdviceV07)
	return d.Message
}

// Scope
// An account servicer sends the CorporateActionInstructionCancellationRequestStatusAdvice message to an account owner or its designated agent to report status of a previously received CorporateActionInstructionCancellationRequest message sent by the account owner. This will include the acknowledgement/rejection of a request to cancel an outstanding instruction.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate),
// using the relevant elements in the business application header (BAH).
type CorporateActionInstructionCancellationRequestStatusAdviceV07 struct {

	// Identification of a related instruction cancellation request document.
	InstructionCancellationRequestIdentification *iso20022.DocumentIdentification9 `xml:"InstrCxlReqId,omitempty"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*iso20022.DocumentIdentification33 `xml:"OthrDocId,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionGeneralInformation109 `xml:"CorpActnGnlInf"`

	// Provides information about the processing status of the instruction cancellation request.
	InstructionCancellationRequestStatus []*iso20022.InstructionCancellationRequestStatus11Choice `xml:"InstrCxlReqSts"`

	// Information about the corporate action option.
	CorporateActionInstruction *iso20022.CorporateActionOption116 `xml:"CorpActnInstr,omitempty"`

	// Provides additional information.
	AdditionalInformation *iso20022.CorporateActionNarrative10 `xml:"AddtlInf,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionInstructionCancellationRequestStatusAdviceV07) AddInstructionCancellationRequestIdentification() *iso20022.DocumentIdentification9 {
	c.InstructionCancellationRequestIdentification = new(iso20022.DocumentIdentification9)
	return c.InstructionCancellationRequestIdentification
}

func (c *CorporateActionInstructionCancellationRequestStatusAdviceV07) AddOtherDocumentIdentification() *iso20022.DocumentIdentification33 {
	newValue := new(iso20022.DocumentIdentification33)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionInstructionCancellationRequestStatusAdviceV07) AddCorporateActionGeneralInformation() *iso20022.CorporateActionGeneralInformation109 {
	c.CorporateActionGeneralInformation = new(iso20022.CorporateActionGeneralInformation109)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionInstructionCancellationRequestStatusAdviceV07) AddInstructionCancellationRequestStatus() *iso20022.InstructionCancellationRequestStatus11Choice {
	newValue := new(iso20022.InstructionCancellationRequestStatus11Choice)
	c.InstructionCancellationRequestStatus = append(c.InstructionCancellationRequestStatus, newValue)
	return newValue
}

func (c *CorporateActionInstructionCancellationRequestStatusAdviceV07) AddCorporateActionInstruction() *iso20022.CorporateActionOption116 {
	c.CorporateActionInstruction = new(iso20022.CorporateActionOption116)
	return c.CorporateActionInstruction
}

func (c *CorporateActionInstructionCancellationRequestStatusAdviceV07) AddAdditionalInformation() *iso20022.CorporateActionNarrative10 {
	c.AdditionalInformation = new(iso20022.CorporateActionNarrative10)
	return c.AdditionalInformation
}

func (c *CorporateActionInstructionCancellationRequestStatusAdviceV07) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
