package seev

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document02500101 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.025.001.01 Document"`
	Message *AgentCAStandingInstructionRequestV01 `xml:"AgtCAStgInstrReq"`
}

func (d *Document02500101) AddMessage() *AgentCAStandingInstructionRequestV01 {
	d.Message = new(AgentCAStandingInstructionRequestV01)
	return d.Message
}

// Scope
// This message is sent by a CSD to the issuer (or its agent) to provide the issuer (or its agent) with the CSD's client details for the distribution of the proceeds of a corporate action event:
// - Gross or net payments.
// - Delivery details for securities that have to be delivered outside of the CSD.
// - Delivery details for cash amounts that have to be delivered outside of the CSD.
// Usage
// This message is used to request the issuer (or its agent) to put a standing instruction in place for proceeds.
// The amendment of a standing instruction is done through a cancel-and-replace mechanism. First a standing instruction cancellation request is sent followed by a new standing instruction request.
type AgentCAStandingInstructionRequestV01 struct {

	// Identification assigned by the Sender to unambiguously identify the request.
	Identification *iso20022.DocumentIdentification8 `xml:"Id"`

	// General information about the standing instruction.
	StandingInstructionGeneralInformation *iso20022.CorporateActionStandingInstructionGeneralInformation1 `xml:"StgInstrGnlInf"`

	// Provides detailed information about the standing instruction.
	StandingInstructionDetails *iso20022.CorporateActionStandingInstruction1 `xml:"StgInstrDtls"`

	// Contact responsible for the transaction identified in the message.
	ContactDetails *iso20022.ContactPerson1 `xml:"CtctDtls,omitempty"`
}

func (a *AgentCAStandingInstructionRequestV01) AddIdentification() *iso20022.DocumentIdentification8 {
	a.Identification = new(iso20022.DocumentIdentification8)
	return a.Identification
}

func (a *AgentCAStandingInstructionRequestV01) AddStandingInstructionGeneralInformation() *iso20022.CorporateActionStandingInstructionGeneralInformation1 {
	a.StandingInstructionGeneralInformation = new(iso20022.CorporateActionStandingInstructionGeneralInformation1)
	return a.StandingInstructionGeneralInformation
}

func (a *AgentCAStandingInstructionRequestV01) AddStandingInstructionDetails() *iso20022.CorporateActionStandingInstruction1 {
	a.StandingInstructionDetails = new(iso20022.CorporateActionStandingInstruction1)
	return a.StandingInstructionDetails
}

func (a *AgentCAStandingInstructionRequestV01) AddContactDetails() *iso20022.ContactPerson1 {
	a.ContactDetails = new(iso20022.ContactPerson1)
	return a.ContactDetails
}
