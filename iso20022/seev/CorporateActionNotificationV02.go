package seev

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document03100102 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.031.001.02 Document"`
	Message *CorporateActionNotificationV02 `xml:"CorpActnNtfctn"`
}

func (d *Document03100102) AddMessage() *CorporateActionNotificationV02 {
	d.Message = new(CorporateActionNotificationV02)
	return d.Message
}

// Scope
// An account servicer sends the CorporateActionNotification message to an account owner or its designated agent to notify details of a corporate action event and optionally account information, eligible balance and entitlements. It may also include possible elections or choices available to the account owner.
// The account servicer can initially send the CorporateActionNotification message as a preliminary advice, subsequently replaced by another CorporateActionNotification message with complete or confirmed information.
// It may also be sent to an account owner or its designated agent, to remind of event details and/or of missing or incomplete instructions for a corporate action event.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate),
// using the relevant elements in the business application header (BAH).
// ISO 15022 - 20022 COEXISTENCE
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type CorporateActionNotificationV02 struct {

	// General information about the event notification type, status and contents.
	NotificationGeneralInformation *iso20022.CorporateActionNotification2 `xml:"NtfctnGnlInf"`

	// Identification of a previously sent notification document.
	PreviousNotificationIdentification *iso20022.DocumentIdentification15 `xml:"PrvsNtfctnId,omitempty"`

	// Identification of a related instruction document.
	InstructionIdentification *iso20022.DocumentIdentification9 `xml:"InstrId,omitempty"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*iso20022.DocumentIdentification13 `xml:"OthrDocId,omitempty"`

	// Identification of an other corporate action event that needs to be closely linked to the processing of the event notified in this document.
	EventsLinkage []*iso20022.CorporateActionEventReference1 `xml:"EvtsLkg,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionGeneralInformation22 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account, owner and account balance.
	AccountDetails *iso20022.AccountIdentification12Choice `xml:"AcctDtls"`

	// Provides details on rights credited to the account as for instance trading period, expiry date, renounceability.
	IntermediateSecurity *iso20022.FinancialInstrumentAttributes17 `xml:"IntrmdtScty,omitempty"`

	// Information about the corporate action event.
	CorporateActionDetails *iso20022.CorporateAction5 `xml:"CorpActnDtls,omitempty"`

	// Information about the corporate action option.
	CorporateActionOptionDetails []*iso20022.CorporateActionOption19 `xml:"CorpActnOptnDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *iso20022.CorporateActionNarrative11 `xml:"AddtlInf,omitempty"`

	// Party appointed to administer the event on behalf of the issuer company/offeror. The party may be contacted for more information about the event.
	IssuerAgent []*iso20022.PartyIdentification47Choice `xml:"IssrAgt,omitempty"`

	// Agent (principal or fiscal paying agent) appointed to execute the payment for the corporate action event on behalf of the issuer company/offeror.
	PayingAgent []*iso20022.PartyIdentification47Choice `xml:"PngAgt,omitempty"`

	// Sub-agent appointed to execute the payment for the corporate action event on behalf of the issuer company/offeror.
	SubPayingAgent []*iso20022.PartyIdentification47Choice `xml:"SubPngAgt,omitempty"`

	// Party/agent responsible for maintaining the register of a security.
	Registrar *iso20022.PartyIdentification47Choice `xml:"Regar,omitempty"`

	// A broker-dealer responsible for reselling to new investors securities (usually bonds) that have been tendered for purchase by their owner.
	ResellingAgent []*iso20022.PartyIdentification47Choice `xml:"RsellngAgt,omitempty"`

	// A trust company, bank or similar financial institution assigned by an issuer to accept presentations of instruments, usually bonds, for transfer and or exchange.
	PhysicalSecuritiesAgent *iso20022.PartyIdentification47Choice `xml:"PhysSctiesAgt,omitempty"`

	// A trust company, bank or similar financial institution who acts on behalf of an out of town agent or event agent where securities can be delivered in person.
	DropAgent *iso20022.PartyIdentification47Choice `xml:"DrpAgt,omitempty"`

	// A trust company, bank or similar financial institution assigned by an issuer to maintain records of investors and account balances and transactions for the consent of a material change.
	SolicitationAgent []*iso20022.PartyIdentification47Choice `xml:"SlctnAgt,omitempty"`

	// A trust company, bank or similar financial institution assigned by an Issuer to provide information and copies of the offering documentation.
	InformationAgent *iso20022.PartyIdentification47Choice `xml:"InfAgt,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionNotificationV02) AddNotificationGeneralInformation() *iso20022.CorporateActionNotification2 {
	c.NotificationGeneralInformation = new(iso20022.CorporateActionNotification2)
	return c.NotificationGeneralInformation
}

func (c *CorporateActionNotificationV02) AddPreviousNotificationIdentification() *iso20022.DocumentIdentification15 {
	c.PreviousNotificationIdentification = new(iso20022.DocumentIdentification15)
	return c.PreviousNotificationIdentification
}

func (c *CorporateActionNotificationV02) AddInstructionIdentification() *iso20022.DocumentIdentification9 {
	c.InstructionIdentification = new(iso20022.DocumentIdentification9)
	return c.InstructionIdentification
}

func (c *CorporateActionNotificationV02) AddOtherDocumentIdentification() *iso20022.DocumentIdentification13 {
	newValue := new(iso20022.DocumentIdentification13)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionNotificationV02) AddEventsLinkage() *iso20022.CorporateActionEventReference1 {
	newValue := new(iso20022.CorporateActionEventReference1)
	c.EventsLinkage = append(c.EventsLinkage, newValue)
	return newValue
}

func (c *CorporateActionNotificationV02) AddCorporateActionGeneralInformation() *iso20022.CorporateActionGeneralInformation22 {
	c.CorporateActionGeneralInformation = new(iso20022.CorporateActionGeneralInformation22)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionNotificationV02) AddAccountDetails() *iso20022.AccountIdentification12Choice {
	c.AccountDetails = new(iso20022.AccountIdentification12Choice)
	return c.AccountDetails
}

func (c *CorporateActionNotificationV02) AddIntermediateSecurity() *iso20022.FinancialInstrumentAttributes17 {
	c.IntermediateSecurity = new(iso20022.FinancialInstrumentAttributes17)
	return c.IntermediateSecurity
}

func (c *CorporateActionNotificationV02) AddCorporateActionDetails() *iso20022.CorporateAction5 {
	c.CorporateActionDetails = new(iso20022.CorporateAction5)
	return c.CorporateActionDetails
}

func (c *CorporateActionNotificationV02) AddCorporateActionOptionDetails() *iso20022.CorporateActionOption19 {
	newValue := new(iso20022.CorporateActionOption19)
	c.CorporateActionOptionDetails = append(c.CorporateActionOptionDetails, newValue)
	return newValue
}

func (c *CorporateActionNotificationV02) AddAdditionalInformation() *iso20022.CorporateActionNarrative11 {
	c.AdditionalInformation = new(iso20022.CorporateActionNarrative11)
	return c.AdditionalInformation
}

func (c *CorporateActionNotificationV02) AddIssuerAgent() *iso20022.PartyIdentification47Choice {
	newValue := new(iso20022.PartyIdentification47Choice)
	c.IssuerAgent = append(c.IssuerAgent, newValue)
	return newValue
}

func (c *CorporateActionNotificationV02) AddPayingAgent() *iso20022.PartyIdentification47Choice {
	newValue := new(iso20022.PartyIdentification47Choice)
	c.PayingAgent = append(c.PayingAgent, newValue)
	return newValue
}

func (c *CorporateActionNotificationV02) AddSubPayingAgent() *iso20022.PartyIdentification47Choice {
	newValue := new(iso20022.PartyIdentification47Choice)
	c.SubPayingAgent = append(c.SubPayingAgent, newValue)
	return newValue
}

func (c *CorporateActionNotificationV02) AddRegistrar() *iso20022.PartyIdentification47Choice {
	c.Registrar = new(iso20022.PartyIdentification47Choice)
	return c.Registrar
}

func (c *CorporateActionNotificationV02) AddResellingAgent() *iso20022.PartyIdentification47Choice {
	newValue := new(iso20022.PartyIdentification47Choice)
	c.ResellingAgent = append(c.ResellingAgent, newValue)
	return newValue
}

func (c *CorporateActionNotificationV02) AddPhysicalSecuritiesAgent() *iso20022.PartyIdentification47Choice {
	c.PhysicalSecuritiesAgent = new(iso20022.PartyIdentification47Choice)
	return c.PhysicalSecuritiesAgent
}

func (c *CorporateActionNotificationV02) AddDropAgent() *iso20022.PartyIdentification47Choice {
	c.DropAgent = new(iso20022.PartyIdentification47Choice)
	return c.DropAgent
}

func (c *CorporateActionNotificationV02) AddSolicitationAgent() *iso20022.PartyIdentification47Choice {
	newValue := new(iso20022.PartyIdentification47Choice)
	c.SolicitationAgent = append(c.SolicitationAgent, newValue)
	return newValue
}

func (c *CorporateActionNotificationV02) AddInformationAgent() *iso20022.PartyIdentification47Choice {
	c.InformationAgent = new(iso20022.PartyIdentification47Choice)
	return c.InformationAgent
}

func (c *CorporateActionNotificationV02) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
