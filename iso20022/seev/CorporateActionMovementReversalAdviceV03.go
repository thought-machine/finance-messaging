package seev

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document03700103 struct {
	XMLName xml.Name                                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.03 Document"`
	Message *CorporateActionMovementReversalAdviceV03 `xml:"CorpActnMvmntRvslAdvc"`
}

func (d *Document03700103) AddMessage() *CorporateActionMovementReversalAdviceV03 {
	d.Message = new(CorporateActionMovementReversalAdviceV03)
	return d.Message
}

// Scope
// An account servicer sends the CorporateActionMovementReversalAdvice message to an account owner or its designated agent to reverse previously confirmed posting of securities or cash.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate),
// using the relevant elements in the business application header (BAH).
// ISO 15022 - 20022 COEXISTENCE
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type CorporateActionMovementReversalAdviceV03 struct {

	// Identification of a previously sent movement confirmation document.
	MovementConfirmationIdentification *iso20022.DocumentIdentification15 `xml:"MvmntConfId"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*iso20022.DocumentIdentification13 `xml:"OthrDocId,omitempty"`

	// Identification of an other corporate action event that needs to be closely linked to the processing of the event notified in this document.
	EventsLinkage []*iso20022.CorporateActionEventReference1 `xml:"EvtsLkg,omitempty"`

	// Reason for the reversal.
	ReversalReason *iso20022.CorporateActionReversalReason1 `xml:"RvslRsn,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionGeneralInformation39 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account, owner and account balance.
	AccountDetails *iso20022.AccountAndBalance4 `xml:"AcctDtls"`

	// Information about the corporate action option.
	CorporateActionConfirmationDetails *iso20022.CorporateActionOption39 `xml:"CorpActnConfDtls"`

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

func (c *CorporateActionMovementReversalAdviceV03) AddMovementConfirmationIdentification() *iso20022.DocumentIdentification15 {
	c.MovementConfirmationIdentification = new(iso20022.DocumentIdentification15)
	return c.MovementConfirmationIdentification
}

func (c *CorporateActionMovementReversalAdviceV03) AddOtherDocumentIdentification() *iso20022.DocumentIdentification13 {
	newValue := new(iso20022.DocumentIdentification13)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionMovementReversalAdviceV03) AddEventsLinkage() *iso20022.CorporateActionEventReference1 {
	newValue := new(iso20022.CorporateActionEventReference1)
	c.EventsLinkage = append(c.EventsLinkage, newValue)
	return newValue
}

func (c *CorporateActionMovementReversalAdviceV03) AddReversalReason() *iso20022.CorporateActionReversalReason1 {
	c.ReversalReason = new(iso20022.CorporateActionReversalReason1)
	return c.ReversalReason
}

func (c *CorporateActionMovementReversalAdviceV03) AddCorporateActionGeneralInformation() *iso20022.CorporateActionGeneralInformation39 {
	c.CorporateActionGeneralInformation = new(iso20022.CorporateActionGeneralInformation39)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionMovementReversalAdviceV03) AddAccountDetails() *iso20022.AccountAndBalance4 {
	c.AccountDetails = new(iso20022.AccountAndBalance4)
	return c.AccountDetails
}

func (c *CorporateActionMovementReversalAdviceV03) AddCorporateActionConfirmationDetails() *iso20022.CorporateActionOption39 {
	c.CorporateActionConfirmationDetails = new(iso20022.CorporateActionOption39)
	return c.CorporateActionConfirmationDetails
}

func (c *CorporateActionMovementReversalAdviceV03) AddAdditionalInformation() *iso20022.CorporateActionNarrative4 {
	c.AdditionalInformation = new(iso20022.CorporateActionNarrative4)
	return c.AdditionalInformation
}

func (c *CorporateActionMovementReversalAdviceV03) AddIssuerAgent() *iso20022.PartyIdentification46Choice {
	newValue := new(iso20022.PartyIdentification46Choice)
	c.IssuerAgent = append(c.IssuerAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementReversalAdviceV03) AddPayingAgent() *iso20022.PartyIdentification46Choice {
	newValue := new(iso20022.PartyIdentification46Choice)
	c.PayingAgent = append(c.PayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementReversalAdviceV03) AddSubPayingAgent() *iso20022.PartyIdentification46Choice {
	newValue := new(iso20022.PartyIdentification46Choice)
	c.SubPayingAgent = append(c.SubPayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementReversalAdviceV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
