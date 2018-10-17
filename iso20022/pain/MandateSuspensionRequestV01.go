package pain

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document01800101 struct {
	XMLName xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.01 Document"`
	Message *MandateSuspensionRequestV01 `xml:"MndtSspnsnReq"`
}

func (d *Document01800101) AddMessage() *MandateSuspensionRequestV01 {
	d.Message = new(MandateSuspensionRequestV01)
	return d.Message
}

// Scope
// The MandateSuspensionRequest message is sent by the initiator of the request to its agent. The initiator can either be the debtor, debtor agent, creditor or creditor agent.
// A MandateSuspensionRequest message is used to request the suspension of an existing mandate until the suspension is lifted.
// Usage
// The MandateSuspensionRequest message can contain one or more suspension requests.
// The messages can be exchanged between creditor and creditor agent or debtor and debtor agent and between creditor agent and debtor agent.
// The MandateSuspensionRequest message can be used in domestic and cross-border scenarios.
//
type MandateSuspensionRequestV01 struct {

	// Set of characteristics to identify the message and parties playing a role in the mandate suspension request, but which are not part of the mandate.
	GroupHeader *iso20022.GroupHeader47 `xml:"GrpHdr"`

	// Set of elements used to provide information on the suspension request of the mandate.
	UnderlyingSuspensionDetails []*iso20022.MandateSuspension1 `xml:"UndrlygSspnsnDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MandateSuspensionRequestV01) AddGroupHeader() *iso20022.GroupHeader47 {
	m.GroupHeader = new(iso20022.GroupHeader47)
	return m.GroupHeader
}

func (m *MandateSuspensionRequestV01) AddUnderlyingSuspensionDetails() *iso20022.MandateSuspension1 {
	newValue := new(iso20022.MandateSuspension1)
	m.UnderlyingSuspensionDetails = append(m.UnderlyingSuspensionDetails, newValue)
	return newValue
}

func (m *MandateSuspensionRequestV01) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}
