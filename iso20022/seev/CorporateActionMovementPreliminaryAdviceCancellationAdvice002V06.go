package seev

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document04400206 struct {
	XMLName xml.Name                                                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.002.06 Document"`
	Message *CorporateActionMovementPreliminaryAdviceCancellationAdvice002V06 `xml:"CorpActnMvmntPrlimryAdvcCxlAdvc"`
}

func (d *Document04400206) AddMessage() *CorporateActionMovementPreliminaryAdviceCancellationAdvice002V06 {
	d.Message = new(CorporateActionMovementPreliminaryAdviceCancellationAdvice002V06)
	return d.Message
}

// Scope
// An account servicer sends the CorporateActionMovementPreliminaryAdviceCancellationAdvice message to an account owner or its designated agent to cancel a previously announced CorporateActionMovementPreliminaryAdvice.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate),
// using the relevant elements in the business application header (BAH).
type CorporateActionMovementPreliminaryAdviceCancellationAdvice002V06 struct {

	// Identification of a previously sent movement preliminary advice document.
	MovementPreliminaryAdviceIdentification *iso20022.DocumentIdentification37 `xml:"MvmntPrlimryAdvcId"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionGeneralInformation99 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account and the account owner.
	AccountDetails *iso20022.AccountIdentification34Choice `xml:"AcctDtls"`

	// Information about the corporate action event.
	CorporateActionDetails *iso20022.CorporateAction35 `xml:"CorpActnDtls,omitempty"`

	// Party appointed to administer the event on behalf of the issuer company/offeror. The party may be contacted for more information about the event.
	IssuerAgent []*iso20022.PartyIdentification104Choice `xml:"IssrAgt,omitempty"`

	// Agent (principal or fiscal paying agent) appointed to execute the payment for the corporate action event on behalf of the issuer company/offeror.
	PayingAgent []*iso20022.PartyIdentification104Choice `xml:"PngAgt,omitempty"`

	// Sub-agent appointed to execute the payment for the corporate action event on behalf of the issuer company/offeror.
	SubPayingAgent []*iso20022.PartyIdentification104Choice `xml:"SubPngAgt,omitempty"`

	// Party/agent responsible for maintaining the register of a security.
	Registrar *iso20022.PartyIdentification104Choice `xml:"Regar,omitempty"`

	// A broker-dealer responsible for reselling to new investors securities (usually bonds) that have been tendered for purchase by their owner.
	ResellingAgent []*iso20022.PartyIdentification104Choice `xml:"RsellngAgt,omitempty"`

	// A trust company, bank or similar financial institution assigned by an issuer to accept presentations of instruments, usually bonds, for transfer and or exchange.
	PhysicalSecuritiesAgent *iso20022.PartyIdentification104Choice `xml:"PhysSctiesAgt,omitempty"`

	// A trust company, bank or similar financial institution who acts on behalf of an out of town agent or event agent where securities can be delivered in person.
	DropAgent *iso20022.PartyIdentification104Choice `xml:"DrpAgt,omitempty"`

	// A trust company, bank or similar financial institution assigned by an issuer to maintain records of investors and account balances and transactions for the consent of a material change.
	SolicitationAgent []*iso20022.PartyIdentification104Choice `xml:"SlctnAgt,omitempty"`

	// A trust company, bank or similar financial institution assigned by an Issuer to provide information and copies of the offering documentation.
	InformationAgent *iso20022.PartyIdentification104Choice `xml:"InfAgt,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdvice002V06) AddMovementPreliminaryAdviceIdentification() *iso20022.DocumentIdentification37 {
	c.MovementPreliminaryAdviceIdentification = new(iso20022.DocumentIdentification37)
	return c.MovementPreliminaryAdviceIdentification
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdvice002V06) AddCorporateActionGeneralInformation() *iso20022.CorporateActionGeneralInformation99 {
	c.CorporateActionGeneralInformation = new(iso20022.CorporateActionGeneralInformation99)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdvice002V06) AddAccountDetails() *iso20022.AccountIdentification34Choice {
	c.AccountDetails = new(iso20022.AccountIdentification34Choice)
	return c.AccountDetails
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdvice002V06) AddCorporateActionDetails() *iso20022.CorporateAction35 {
	c.CorporateActionDetails = new(iso20022.CorporateAction35)
	return c.CorporateActionDetails
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdvice002V06) AddIssuerAgent() *iso20022.PartyIdentification104Choice {
	newValue := new(iso20022.PartyIdentification104Choice)
	c.IssuerAgent = append(c.IssuerAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdvice002V06) AddPayingAgent() *iso20022.PartyIdentification104Choice {
	newValue := new(iso20022.PartyIdentification104Choice)
	c.PayingAgent = append(c.PayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdvice002V06) AddSubPayingAgent() *iso20022.PartyIdentification104Choice {
	newValue := new(iso20022.PartyIdentification104Choice)
	c.SubPayingAgent = append(c.SubPayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdvice002V06) AddRegistrar() *iso20022.PartyIdentification104Choice {
	c.Registrar = new(iso20022.PartyIdentification104Choice)
	return c.Registrar
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdvice002V06) AddResellingAgent() *iso20022.PartyIdentification104Choice {
	newValue := new(iso20022.PartyIdentification104Choice)
	c.ResellingAgent = append(c.ResellingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdvice002V06) AddPhysicalSecuritiesAgent() *iso20022.PartyIdentification104Choice {
	c.PhysicalSecuritiesAgent = new(iso20022.PartyIdentification104Choice)
	return c.PhysicalSecuritiesAgent
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdvice002V06) AddDropAgent() *iso20022.PartyIdentification104Choice {
	c.DropAgent = new(iso20022.PartyIdentification104Choice)
	return c.DropAgent
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdvice002V06) AddSolicitationAgent() *iso20022.PartyIdentification104Choice {
	newValue := new(iso20022.PartyIdentification104Choice)
	c.SolicitationAgent = append(c.SolicitationAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdvice002V06) AddInformationAgent() *iso20022.PartyIdentification104Choice {
	c.InformationAgent = new(iso20022.PartyIdentification104Choice)
	return c.InformationAgent
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdvice002V06) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
