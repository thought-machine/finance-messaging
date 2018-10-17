package seev

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document04000103 struct {
	XMLName xml.Name                                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.03 Document"`
	Message *CorporateActionInstructionCancellationRequestV03 `xml:"CorpActnInstrCxlReq"`
}

func (d *Document04000103) AddMessage() *CorporateActionInstructionCancellationRequestV03 {
	d.Message = new(CorporateActionInstructionCancellationRequestV03)
	return d.Message
}

// Scope
// An account owner sends the CorporateActionInstructionCancellationRequest message to an account servicer to request cancellation of a previously sent corporate action election instruction.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate),
// using the relevant elements in the business application header (BAH).
// ISO 15022 - 20022 COEXISTENCE
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type CorporateActionInstructionCancellationRequestV03 struct {

	// When used in a corporate action instruction, indicates that the current instruction is replacing a previous one that was cancelled earlier. When used in a corporate action instruction cancellation request, indicates that cancelled instruction will be replaced by a new corporate action instruction to be sent later.
	ChangeInstructionIndicator *iso20022.YesNoIndicator `xml:"ChngInstrInd,omitempty"`

	// Identification of a previously sent instruction document.
	InstructionIdentification *iso20022.DocumentIdentification15 `xml:"InstrId"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionGeneralInformation33 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account and the account owner.
	AccountDetails *iso20022.AccountIdentification17 `xml:"AcctDtls"`

	// Information about the corporate action option.
	CorporateActionInstruction *iso20022.CorporateActionOption42 `xml:"CorpActnInstr"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionInstructionCancellationRequestV03) SetChangeInstructionIndicator(value string) {
	c.ChangeInstructionIndicator = (*iso20022.YesNoIndicator)(&value)
}

func (c *CorporateActionInstructionCancellationRequestV03) AddInstructionIdentification() *iso20022.DocumentIdentification15 {
	c.InstructionIdentification = new(iso20022.DocumentIdentification15)
	return c.InstructionIdentification
}

func (c *CorporateActionInstructionCancellationRequestV03) AddCorporateActionGeneralInformation() *iso20022.CorporateActionGeneralInformation33 {
	c.CorporateActionGeneralInformation = new(iso20022.CorporateActionGeneralInformation33)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionInstructionCancellationRequestV03) AddAccountDetails() *iso20022.AccountIdentification17 {
	c.AccountDetails = new(iso20022.AccountIdentification17)
	return c.AccountDetails
}

func (c *CorporateActionInstructionCancellationRequestV03) AddCorporateActionInstruction() *iso20022.CorporateActionOption42 {
	c.CorporateActionInstruction = new(iso20022.CorporateActionOption42)
	return c.CorporateActionInstruction
}

func (c *CorporateActionInstructionCancellationRequestV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
