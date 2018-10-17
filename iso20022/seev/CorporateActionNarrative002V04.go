package seev

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document03800204 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.002.04 Document"`
	Message *CorporateActionNarrative002V04 `xml:"CorpActnNrrtv"`
}

func (d *Document03800204) AddMessage() *CorporateActionNarrative002V04 {
	d.Message = new(CorporateActionNarrative002V04)
	return d.Message
}

// Scope
// The CorporateActionNarrative message is sent between an account servicer and an account owner or its designated agent to cater for tax reclaims, restrictions, documentation requirements. This message is bi-directional.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate),
// using the relevant elements in the business application header (BAH).
type CorporateActionNarrative002V04 struct {

	// General information about the safekeeping account and the account owner.
	AccountDetails *iso20022.AccountIdentification37Choice `xml:"AcctDtls,omitempty"`

	// Provides information about the securitised right for entitlement.
	UnderlyingSecurity *iso20022.SecurityIdentification20 `xml:"UndrlygScty,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionGeneralInformation102 `xml:"CorpActnGnlInf"`

	// Provides additional information.
	AdditionalInformation *iso20022.UpdatedAdditionalInformation10 `xml:"AddtlInf"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionNarrative002V04) AddAccountDetails() *iso20022.AccountIdentification37Choice {
	c.AccountDetails = new(iso20022.AccountIdentification37Choice)
	return c.AccountDetails
}

func (c *CorporateActionNarrative002V04) AddUnderlyingSecurity() *iso20022.SecurityIdentification20 {
	c.UnderlyingSecurity = new(iso20022.SecurityIdentification20)
	return c.UnderlyingSecurity
}

func (c *CorporateActionNarrative002V04) AddCorporateActionGeneralInformation() *iso20022.CorporateActionGeneralInformation102 {
	c.CorporateActionGeneralInformation = new(iso20022.CorporateActionGeneralInformation102)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionNarrative002V04) AddAdditionalInformation() *iso20022.UpdatedAdditionalInformation10 {
	c.AdditionalInformation = new(iso20022.UpdatedAdditionalInformation10)
	return c.AdditionalInformation
}

func (c *CorporateActionNarrative002V04) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
