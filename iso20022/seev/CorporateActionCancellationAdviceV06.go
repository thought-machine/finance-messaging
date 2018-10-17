package seev

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document03900106 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.06 Document"`
	Message *CorporateActionCancellationAdviceV06 `xml:"CorpActnCxlAdvc"`
}

func (d *Document03900106) AddMessage() *CorporateActionCancellationAdviceV06 {
	d.Message = new(CorporateActionCancellationAdviceV06)
	return d.Message
}

// Scope
// An account servicer sends the CorporateActionCancellationAdvice message to an account owner or its designated agent to cancel a previously announced corporate action event in case of error from the account servicer or in case of withdrawal by the issuer.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate),
// using the relevant elements in the business application header (BAH).
type CorporateActionCancellationAdviceV06 struct {

	// General information about the event cancellation status and cancellation reason.
	CancellationAdviceGeneralInformation *iso20022.CorporateActionCancellation3 `xml:"CxlAdvcGnlInf"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionGeneralInformation86 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account and the account owner.
	AccountsDetails *iso20022.AccountIdentification29Choice `xml:"AcctsDtls"`

	// Information about the corporate action event.
	CorporateActionDetails *iso20022.CorporateAction34 `xml:"CorpActnDtls,omitempty"`

	// Party appointed to administer the event on behalf of the issuer company/offeror. The party may be contacted for more information about the event.
	IssuerAgent []*iso20022.PartyIdentification71Choice `xml:"IssrAgt,omitempty"`

	// Agent (principal or fiscal paying agent) appointed to execute the payment for the corporate action event on behalf of the issuer company/offeror.
	PayingAgent []*iso20022.PartyIdentification71Choice `xml:"PngAgt,omitempty"`

	// Sub-agent appointed to execute the payment for the corporate action event on behalf of the issuer company/offeror.
	SubPayingAgent []*iso20022.PartyIdentification71Choice `xml:"SubPngAgt,omitempty"`

	// Party/agent responsible for maintaining the register of a security.
	Registrar *iso20022.PartyIdentification71Choice `xml:"Regar,omitempty"`

	// A broker-dealer responsible for reselling to new investors securities (usually bonds) that have been tendered for purchase by their owner.
	ResellingAgent []*iso20022.PartyIdentification71Choice `xml:"RsellngAgt,omitempty"`

	// A trust company, bank or similar financial institution assigned by an issuer to accept presentations of instruments, usually bonds, for transfer and or exchange.
	PhysicalSecuritiesAgent *iso20022.PartyIdentification71Choice `xml:"PhysSctiesAgt,omitempty"`

	// A trust company, bank or similar financial institution who acts on behalf of an out of town agent or event agent where securities can be delivered in person.
	DropAgent *iso20022.PartyIdentification71Choice `xml:"DrpAgt,omitempty"`

	// A trust company, bank or similar financial institution assigned by an issuer to maintain records of investors and account balances and transactions for the consent of a material change.
	SolicitationAgent []*iso20022.PartyIdentification71Choice `xml:"SlctnAgt,omitempty"`

	// A trust company, bank or similar financial institution assigned by an Issuer to provide information and copies of the offering documentation.
	InformationAgent *iso20022.PartyIdentification71Choice `xml:"InfAgt,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionCancellationAdviceV06) AddCancellationAdviceGeneralInformation() *iso20022.CorporateActionCancellation3 {
	c.CancellationAdviceGeneralInformation = new(iso20022.CorporateActionCancellation3)
	return c.CancellationAdviceGeneralInformation
}

func (c *CorporateActionCancellationAdviceV06) AddCorporateActionGeneralInformation() *iso20022.CorporateActionGeneralInformation86 {
	c.CorporateActionGeneralInformation = new(iso20022.CorporateActionGeneralInformation86)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionCancellationAdviceV06) AddAccountsDetails() *iso20022.AccountIdentification29Choice {
	c.AccountsDetails = new(iso20022.AccountIdentification29Choice)
	return c.AccountsDetails
}

func (c *CorporateActionCancellationAdviceV06) AddCorporateActionDetails() *iso20022.CorporateAction34 {
	c.CorporateActionDetails = new(iso20022.CorporateAction34)
	return c.CorporateActionDetails
}

func (c *CorporateActionCancellationAdviceV06) AddIssuerAgent() *iso20022.PartyIdentification71Choice {
	newValue := new(iso20022.PartyIdentification71Choice)
	c.IssuerAgent = append(c.IssuerAgent, newValue)
	return newValue
}

func (c *CorporateActionCancellationAdviceV06) AddPayingAgent() *iso20022.PartyIdentification71Choice {
	newValue := new(iso20022.PartyIdentification71Choice)
	c.PayingAgent = append(c.PayingAgent, newValue)
	return newValue
}

func (c *CorporateActionCancellationAdviceV06) AddSubPayingAgent() *iso20022.PartyIdentification71Choice {
	newValue := new(iso20022.PartyIdentification71Choice)
	c.SubPayingAgent = append(c.SubPayingAgent, newValue)
	return newValue
}

func (c *CorporateActionCancellationAdviceV06) AddRegistrar() *iso20022.PartyIdentification71Choice {
	c.Registrar = new(iso20022.PartyIdentification71Choice)
	return c.Registrar
}

func (c *CorporateActionCancellationAdviceV06) AddResellingAgent() *iso20022.PartyIdentification71Choice {
	newValue := new(iso20022.PartyIdentification71Choice)
	c.ResellingAgent = append(c.ResellingAgent, newValue)
	return newValue
}

func (c *CorporateActionCancellationAdviceV06) AddPhysicalSecuritiesAgent() *iso20022.PartyIdentification71Choice {
	c.PhysicalSecuritiesAgent = new(iso20022.PartyIdentification71Choice)
	return c.PhysicalSecuritiesAgent
}

func (c *CorporateActionCancellationAdviceV06) AddDropAgent() *iso20022.PartyIdentification71Choice {
	c.DropAgent = new(iso20022.PartyIdentification71Choice)
	return c.DropAgent
}

func (c *CorporateActionCancellationAdviceV06) AddSolicitationAgent() *iso20022.PartyIdentification71Choice {
	newValue := new(iso20022.PartyIdentification71Choice)
	c.SolicitationAgent = append(c.SolicitationAgent, newValue)
	return newValue
}

func (c *CorporateActionCancellationAdviceV06) AddInformationAgent() *iso20022.PartyIdentification71Choice {
	c.InformationAgent = new(iso20022.PartyIdentification71Choice)
	return c.InformationAgent
}

func (c *CorporateActionCancellationAdviceV06) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
