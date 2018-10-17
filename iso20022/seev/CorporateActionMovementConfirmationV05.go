package seev

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document03600105 struct {
	XMLName xml.Name                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.036.001.05 Document"`
	Message *CorporateActionMovementConfirmationV05 `xml:"CorpActnMvmntConf"`
}

func (d *Document03600105) AddMessage() *CorporateActionMovementConfirmationV05 {
	d.Message = new(CorporateActionMovementConfirmationV05)
	return d.Message
}

// Scope
// An account servicer sends the CorporateActionMovementConfirmation message to an account owner or its designated agent to confirm posting of securities or cash as a result of a corporate action event.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate),
// using the relevant elements in the business application header (BAH).
// ISO 15022 - 20022 COEXISTENCE
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type CorporateActionMovementConfirmationV05 struct {

	// Identification of a previously sent notification document.
	NotificationIdentification *iso20022.DocumentIdentification15 `xml:"NtfctnId,omitempty"`

	// Identification of a previously sent movement preliminary advice document.
	MovementPreliminaryAdviceIdentification *iso20022.DocumentIdentification15 `xml:"MvmntPrlimryAdvcId,omitempty"`

	// Identification of a related instruction document.
	InstructionIdentification *iso20022.DocumentIdentification9 `xml:"InstrId,omitempty"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*iso20022.DocumentIdentification13 `xml:"OthrDocId,omitempty"`

	// Identification of an other corporate action event that needs to be closely  linked to the processing of the event notified in this document.
	EventsLinkage []*iso20022.CorporateActionEventReference1 `xml:"EvtsLkg,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionGeneralInformation50 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account, owner and account balance.
	AccountDetails *iso20022.AccountAndBalance26 `xml:"AcctDtls"`

	// Information about the corporate action event.
	CorporateActionDetails *iso20022.CorporateAction14 `xml:"CorpActnDtls,omitempty"`

	// Information about the corporate action option.
	CorporateActionConfirmationDetails *iso20022.CorporateActionOption102 `xml:"CorpActnConfDtls"`

	// Provides additional information.
	AdditionalInformation *iso20022.CorporateActionNarrative4 `xml:"AddtlInf,omitempty"`

	// Party appointed to administer the event on behalf of the issuer company/offeror. The party may be contacted for more information about the event.
	IssuerAgent []*iso20022.PartyIdentification46Choice `xml:"IssrAgt,omitempty"`

	// Agent (principal or fiscal paying agent) appointed to execute the payment for the corporate action event on behalf of the issuer company/offeror.
	PayingAgent []*iso20022.PartyIdentification46Choice `xml:"PngAgt,omitempty"`

	// Sub-agent appointed to execute the payment for the corporate action event on behalf of the issuer company/offeror.
	SubPayingAgent []*iso20022.PartyIdentification46Choice `xml:"SubPngAgt,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionMovementConfirmationV05) AddNotificationIdentification() *iso20022.DocumentIdentification15 {
	c.NotificationIdentification = new(iso20022.DocumentIdentification15)
	return c.NotificationIdentification
}

func (c *CorporateActionMovementConfirmationV05) AddMovementPreliminaryAdviceIdentification() *iso20022.DocumentIdentification15 {
	c.MovementPreliminaryAdviceIdentification = new(iso20022.DocumentIdentification15)
	return c.MovementPreliminaryAdviceIdentification
}

func (c *CorporateActionMovementConfirmationV05) AddInstructionIdentification() *iso20022.DocumentIdentification9 {
	c.InstructionIdentification = new(iso20022.DocumentIdentification9)
	return c.InstructionIdentification
}

func (c *CorporateActionMovementConfirmationV05) AddOtherDocumentIdentification() *iso20022.DocumentIdentification13 {
	newValue := new(iso20022.DocumentIdentification13)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionMovementConfirmationV05) AddEventsLinkage() *iso20022.CorporateActionEventReference1 {
	newValue := new(iso20022.CorporateActionEventReference1)
	c.EventsLinkage = append(c.EventsLinkage, newValue)
	return newValue
}

func (c *CorporateActionMovementConfirmationV05) AddCorporateActionGeneralInformation() *iso20022.CorporateActionGeneralInformation50 {
	c.CorporateActionGeneralInformation = new(iso20022.CorporateActionGeneralInformation50)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionMovementConfirmationV05) AddAccountDetails() *iso20022.AccountAndBalance26 {
	c.AccountDetails = new(iso20022.AccountAndBalance26)
	return c.AccountDetails
}

func (c *CorporateActionMovementConfirmationV05) AddCorporateActionDetails() *iso20022.CorporateAction14 {
	c.CorporateActionDetails = new(iso20022.CorporateAction14)
	return c.CorporateActionDetails
}

func (c *CorporateActionMovementConfirmationV05) AddCorporateActionConfirmationDetails() *iso20022.CorporateActionOption102 {
	c.CorporateActionConfirmationDetails = new(iso20022.CorporateActionOption102)
	return c.CorporateActionConfirmationDetails
}

func (c *CorporateActionMovementConfirmationV05) AddAdditionalInformation() *iso20022.CorporateActionNarrative4 {
	c.AdditionalInformation = new(iso20022.CorporateActionNarrative4)
	return c.AdditionalInformation
}

func (c *CorporateActionMovementConfirmationV05) AddIssuerAgent() *iso20022.PartyIdentification46Choice {
	newValue := new(iso20022.PartyIdentification46Choice)
	c.IssuerAgent = append(c.IssuerAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementConfirmationV05) AddPayingAgent() *iso20022.PartyIdentification46Choice {
	newValue := new(iso20022.PartyIdentification46Choice)
	c.PayingAgent = append(c.PayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementConfirmationV05) AddSubPayingAgent() *iso20022.PartyIdentification46Choice {
	newValue := new(iso20022.PartyIdentification46Choice)
	c.SubPayingAgent = append(c.SubPayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementConfirmationV05) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
