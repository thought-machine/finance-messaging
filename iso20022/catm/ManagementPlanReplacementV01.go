package catm

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document00200101 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.002.001.01 Document"`
	Message *ManagementPlanReplacementV01 `xml:"MgmtPlanRplcmnt"`
}

func (d *Document00200101) AddMessage() *ManagementPlanReplacementV01 {
	d.Message = new(ManagementPlanReplacementV01)
	return d.Message
}

// Scope
// The ManagementPlanReplacement message is sent by the master terminal manager or delegated terminal manager to the acceptor system to replace the TMS action list of the POI system.
// Usage
// The ManagementPlanReplacement message may embed the information required by the acceptor system for the planning of the TMS actions to be performed by the POI including the trigger, time conditions and TMS addresses.
type ManagementPlanReplacementV01 struct {

	// Set of characteristics related to the transfer of the management plan.
	Header *iso20022.Header4 `xml:"Hdr"`

	// Sequence of terminal maintenance actions to be performed by a point of interaction (POI).
	ManagementPlan *iso20022.ManagementPlan1 `xml:"MgmtPlan"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *iso20022.ContentInformationType1 `xml:"SctyTrlr"`
}

func (m *ManagementPlanReplacementV01) AddHeader() *iso20022.Header4 {
	m.Header = new(iso20022.Header4)
	return m.Header
}

func (m *ManagementPlanReplacementV01) AddManagementPlan() *iso20022.ManagementPlan1 {
	m.ManagementPlan = new(iso20022.ManagementPlan1)
	return m.ManagementPlan
}

func (m *ManagementPlanReplacementV01) AddSecurityTrailer() *iso20022.ContentInformationType1 {
	m.SecurityTrailer = new(iso20022.ContentInformationType1)
	return m.SecurityTrailer
}
