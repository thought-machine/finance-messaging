package seev

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document01700101 struct {
	XMLName xml.Name                                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.017.001.01 Document"`
	Message *AgentCAGlobalDistributionAuthorisationRequestV01 `xml:"AgtCAGblDstrbtnAuthstnReq"`
}

func (d *Document01700101) AddMessage() *AgentCAGlobalDistributionAuthorisationRequestV01 {
	d.Message = new(AgentCAGlobalDistributionAuthorisationRequestV01)
	return d.Message
}

// Scope
// This message is sent by a CSD to an issuer (or its agent) to request the authorisation to process the entitlement movements (cash and/or securities) calculated by the CSD for a given corporate action entire event, a given corporate action option and optionally a given resource.
// This message can also be sent to request the issuer (or its agent) to make available / deliver the relevant resources to the CSD.
// Usage
// This message is used to request the authorisation to process the entitlement movements calculated by the CSD for a given corporate action event and option. An Agent Corporate Action Global Distribution Authorisation Request message must be sent for each option and if several resources are associated to an option, an Agent Corporate Action Global Distribution Authorisation Request message can be sent for each resource.
// This message can also be used to pre-advise a global distribution authorisation request, in which case the value of the field pre-advice indicator must be set to yes.
type AgentCAGlobalDistributionAuthorisationRequestV01 struct {

	// Identification assigned by the Sender to unambiguously identify the request.
	Identification *iso20022.DocumentIdentification8 `xml:"Id"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *iso20022.CorporateActionInformation1 `xml:"CorpActnGnlInf"`

	// Provides detailed information about the global distribution.
	GlobalDistributionDetails *iso20022.GlobalDistributionRequest1 `xml:"GblDstrbtnDtls"`
}

func (a *AgentCAGlobalDistributionAuthorisationRequestV01) AddIdentification() *iso20022.DocumentIdentification8 {
	a.Identification = new(iso20022.DocumentIdentification8)
	return a.Identification
}

func (a *AgentCAGlobalDistributionAuthorisationRequestV01) AddCorporateActionGeneralInformation() *iso20022.CorporateActionInformation1 {
	a.CorporateActionGeneralInformation = new(iso20022.CorporateActionInformation1)
	return a.CorporateActionGeneralInformation
}

func (a *AgentCAGlobalDistributionAuthorisationRequestV01) AddGlobalDistributionDetails() *iso20022.GlobalDistributionRequest1 {
	a.GlobalDistributionDetails = new(iso20022.GlobalDistributionRequest1)
	return a.GlobalDistributionDetails
}
