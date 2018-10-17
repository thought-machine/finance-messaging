package seev

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document03500107 struct {
	XMLName xml.Name                                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.035.001.07 Document"`
	Message *CorporateActionMovementPreliminaryAdviceV07 `xml:"CorpActnMvmntPrlimryAdvc"`
}

func (d *Document03500107) AddMessage() *CorporateActionMovementPreliminaryAdviceV07 {
	d.Message = new(CorporateActionMovementPreliminaryAdviceV07)
	return d.Message
}

// Scope
// An account servicer sends the CorporateActionMovementPreliminaryAdvice message to an account owner or its designated agent to pre-advise upcoming posting or reversal of securities and/or cash postings.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate),
// using the relevant elements in the business application header (BAH).
type CorporateActionMovementPreliminaryAdviceV07 struct {

	// Page number of the message and  continuation indicator to indicate that the multi-parts preliminary advice is to continue or that the message is the last page of the multi-parts preliminary advice.
	Pagination *iso20022.Pagination `xml:"Pgntn,omitempty"`

	// General information about the movement preliminary advice document.
	MovementPreliminaryAdviceGeneralInformation *iso20022.CorporateActionPreliminaryAdviceType2 `xml:"MvmntPrlimryAdvcGnlInf"`

	// Identification of a previously sent movement preliminary advice document.
	PreviousMovementPreliminaryAdviceIdentification *iso20022.DocumentIdentification31 `xml:"PrvsMvmntPrlimryAdvcId,omitempty"`

	// Identification of a previously sent notification document.
	NotificationIdentification *iso20022.DocumentIdentification31 `xml:"NtfctnId,omitempty"`

	// Identification of a previously sent movement confirmation document.
	MovementConfirmationIdentification *iso20022.DocumentIdentification31 `xml:"MvmntConfId,omitempty"`

	// Identification of a related instruction document.
	InstructionIdentification *iso20022.DocumentIdentification9 `xml:"InstrId,omitempty"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*iso20022.DocumentIdentification32 `xml:"OthrDocId,omitempty"`

	// Identification of an other corporate action event that needs to be closely linked to the processing of the event notified in this document.
	EventsLinkage []*iso20022.CorporateActionEventReference3 `xml:"EvtsLkg,omitempty"`

	// Reason for the reversal.
	ReversalReason *iso20022.CorporateActionReversalReason3 `xml:"RvslRsn,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionGeneralInformation84 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account, owner and account balance.
	AccountDetails *iso20022.AccountIdentification32Choice `xml:"AcctDtls"`

	// Information about the corporate action event.
	CorporateActionDetails *iso20022.CorporateAction32 `xml:"CorpActnDtls,omitempty"`

	// Information about the corporate action option.
	CorporateActionMovementDetails []*iso20022.CorporateActionOption115 `xml:"CorpActnMvmntDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *iso20022.CorporateActionNarrative28 `xml:"AddtlInf,omitempty"`

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

func (c *CorporateActionMovementPreliminaryAdviceV07) AddPagination() *iso20022.Pagination {
	c.Pagination = new(iso20022.Pagination)
	return c.Pagination
}

func (c *CorporateActionMovementPreliminaryAdviceV07) AddMovementPreliminaryAdviceGeneralInformation() *iso20022.CorporateActionPreliminaryAdviceType2 {
	c.MovementPreliminaryAdviceGeneralInformation = new(iso20022.CorporateActionPreliminaryAdviceType2)
	return c.MovementPreliminaryAdviceGeneralInformation
}

func (c *CorporateActionMovementPreliminaryAdviceV07) AddPreviousMovementPreliminaryAdviceIdentification() *iso20022.DocumentIdentification31 {
	c.PreviousMovementPreliminaryAdviceIdentification = new(iso20022.DocumentIdentification31)
	return c.PreviousMovementPreliminaryAdviceIdentification
}

func (c *CorporateActionMovementPreliminaryAdviceV07) AddNotificationIdentification() *iso20022.DocumentIdentification31 {
	c.NotificationIdentification = new(iso20022.DocumentIdentification31)
	return c.NotificationIdentification
}

func (c *CorporateActionMovementPreliminaryAdviceV07) AddMovementConfirmationIdentification() *iso20022.DocumentIdentification31 {
	c.MovementConfirmationIdentification = new(iso20022.DocumentIdentification31)
	return c.MovementConfirmationIdentification
}

func (c *CorporateActionMovementPreliminaryAdviceV07) AddInstructionIdentification() *iso20022.DocumentIdentification9 {
	c.InstructionIdentification = new(iso20022.DocumentIdentification9)
	return c.InstructionIdentification
}

func (c *CorporateActionMovementPreliminaryAdviceV07) AddOtherDocumentIdentification() *iso20022.DocumentIdentification32 {
	newValue := new(iso20022.DocumentIdentification32)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV07) AddEventsLinkage() *iso20022.CorporateActionEventReference3 {
	newValue := new(iso20022.CorporateActionEventReference3)
	c.EventsLinkage = append(c.EventsLinkage, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV07) AddReversalReason() *iso20022.CorporateActionReversalReason3 {
	c.ReversalReason = new(iso20022.CorporateActionReversalReason3)
	return c.ReversalReason
}

func (c *CorporateActionMovementPreliminaryAdviceV07) AddCorporateActionGeneralInformation() *iso20022.CorporateActionGeneralInformation84 {
	c.CorporateActionGeneralInformation = new(iso20022.CorporateActionGeneralInformation84)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionMovementPreliminaryAdviceV07) AddAccountDetails() *iso20022.AccountIdentification32Choice {
	c.AccountDetails = new(iso20022.AccountIdentification32Choice)
	return c.AccountDetails
}

func (c *CorporateActionMovementPreliminaryAdviceV07) AddCorporateActionDetails() *iso20022.CorporateAction32 {
	c.CorporateActionDetails = new(iso20022.CorporateAction32)
	return c.CorporateActionDetails
}

func (c *CorporateActionMovementPreliminaryAdviceV07) AddCorporateActionMovementDetails() *iso20022.CorporateActionOption115 {
	newValue := new(iso20022.CorporateActionOption115)
	c.CorporateActionMovementDetails = append(c.CorporateActionMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV07) AddAdditionalInformation() *iso20022.CorporateActionNarrative28 {
	c.AdditionalInformation = new(iso20022.CorporateActionNarrative28)
	return c.AdditionalInformation
}

func (c *CorporateActionMovementPreliminaryAdviceV07) AddIssuerAgent() *iso20022.PartyIdentification71Choice {
	newValue := new(iso20022.PartyIdentification71Choice)
	c.IssuerAgent = append(c.IssuerAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV07) AddPayingAgent() *iso20022.PartyIdentification71Choice {
	newValue := new(iso20022.PartyIdentification71Choice)
	c.PayingAgent = append(c.PayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV07) AddSubPayingAgent() *iso20022.PartyIdentification71Choice {
	newValue := new(iso20022.PartyIdentification71Choice)
	c.SubPayingAgent = append(c.SubPayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV07) AddRegistrar() *iso20022.PartyIdentification71Choice {
	c.Registrar = new(iso20022.PartyIdentification71Choice)
	return c.Registrar
}

func (c *CorporateActionMovementPreliminaryAdviceV07) AddResellingAgent() *iso20022.PartyIdentification71Choice {
	newValue := new(iso20022.PartyIdentification71Choice)
	c.ResellingAgent = append(c.ResellingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV07) AddPhysicalSecuritiesAgent() *iso20022.PartyIdentification71Choice {
	c.PhysicalSecuritiesAgent = new(iso20022.PartyIdentification71Choice)
	return c.PhysicalSecuritiesAgent
}

func (c *CorporateActionMovementPreliminaryAdviceV07) AddDropAgent() *iso20022.PartyIdentification71Choice {
	c.DropAgent = new(iso20022.PartyIdentification71Choice)
	return c.DropAgent
}

func (c *CorporateActionMovementPreliminaryAdviceV07) AddSolicitationAgent() *iso20022.PartyIdentification71Choice {
	newValue := new(iso20022.PartyIdentification71Choice)
	c.SolicitationAgent = append(c.SolicitationAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV07) AddInformationAgent() *iso20022.PartyIdentification71Choice {
	c.InformationAgent = new(iso20022.PartyIdentification71Choice)
	return c.InformationAgent
}

func (c *CorporateActionMovementPreliminaryAdviceV07) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
