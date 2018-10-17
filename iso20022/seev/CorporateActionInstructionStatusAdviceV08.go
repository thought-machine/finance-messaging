package seev

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document03400108 struct {
	XMLName xml.Name                                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.08 Document"`
	Message *CorporateActionInstructionStatusAdviceV08 `xml:"CorpActnInstrStsAdvc"`
}

func (d *Document03400108) AddMessage() *CorporateActionInstructionStatusAdviceV08 {
	d.Message = new(CorporateActionInstructionStatusAdviceV08)
	return d.Message
}

// Scope
// An account servicer sends the CorporateActionInstructionStatusAdvice message to an account owner or its designated agent, to report status of a received corporate action election instruction.
// This message is used to advise the status, or a change in status, of a corporate action-related transaction previously instructed by, or executed on behalf of, the account owner. This will include the acknowledgement/rejection of a corporate action instruction.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate),
// using the relevant elements in the business application header (BAH).
type CorporateActionInstructionStatusAdviceV08 struct {

	// Identification of a related instruction document.
	InstructionIdentification *iso20022.DocumentIdentification9 `xml:"InstrId,omitempty"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*iso20022.DocumentIdentification33 `xml:"OthrDocId,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionGeneralInformation109 `xml:"CorpActnGnlInf"`

	// Provides information about the processing status of the instruction.
	InstructionProcessingStatus []*iso20022.InstructionProcessingStatus29Choice `xml:"InstrPrcgSts"`

	// Information about the corporate action instruction.
	CorporateActionInstruction *iso20022.CorporateActionOption116 `xml:"CorpActnInstr,omitempty"`

	// Provides additional information.
	AdditionalInformation *iso20022.CorporateActionNarrative10 `xml:"AddtlInf,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionInstructionStatusAdviceV08) AddInstructionIdentification() *iso20022.DocumentIdentification9 {
	c.InstructionIdentification = new(iso20022.DocumentIdentification9)
	return c.InstructionIdentification
}

func (c *CorporateActionInstructionStatusAdviceV08) AddOtherDocumentIdentification() *iso20022.DocumentIdentification33 {
	newValue := new(iso20022.DocumentIdentification33)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionInstructionStatusAdviceV08) AddCorporateActionGeneralInformation() *iso20022.CorporateActionGeneralInformation109 {
	c.CorporateActionGeneralInformation = new(iso20022.CorporateActionGeneralInformation109)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionInstructionStatusAdviceV08) AddInstructionProcessingStatus() *iso20022.InstructionProcessingStatus29Choice {
	newValue := new(iso20022.InstructionProcessingStatus29Choice)
	c.InstructionProcessingStatus = append(c.InstructionProcessingStatus, newValue)
	return newValue
}

func (c *CorporateActionInstructionStatusAdviceV08) AddCorporateActionInstruction() *iso20022.CorporateActionOption116 {
	c.CorporateActionInstruction = new(iso20022.CorporateActionOption116)
	return c.CorporateActionInstruction
}

func (c *CorporateActionInstructionStatusAdviceV08) AddAdditionalInformation() *iso20022.CorporateActionNarrative10 {
	c.AdditionalInformation = new(iso20022.CorporateActionNarrative10)
	return c.AdditionalInformation
}

func (c *CorporateActionInstructionStatusAdviceV08) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
