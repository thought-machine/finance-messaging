package reda

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document06000101 struct {
	XMLName xml.Name                                    `xml:"urn:iso:std:iso:20022:tech:xsd:reda.060.001.01 Document"`
	Message *NettingCutOffReferenceDataUpdateRequestV01 `xml:"NetgCutOffRefDataUpdReq"`
}

func (d *Document06000101) AddMessage() *NettingCutOffReferenceDataUpdateRequestV01 {
	d.Message = new(NettingCutOffReferenceDataUpdateRequestV01)
	return d.Message
}

// The Netting Cut Off Reference Data Update Request message is sent to a central system by a participant to request an update to a participant's netting cut off held at the central system.
type NettingCutOffReferenceDataUpdateRequestV01 struct {

	// Specifies the meta data for the netting cut off update request.
	RequestData *iso20022.RequestData1 `xml:"ReqData"`

	// Specifies the information regarding the changes to the netting cut off.
	NettingCutOffRequest []*iso20022.NettingCutOff1 `xml:"NetgCutOffReq"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (n *NettingCutOffReferenceDataUpdateRequestV01) AddRequestData() *iso20022.RequestData1 {
	n.RequestData = new(iso20022.RequestData1)
	return n.RequestData
}

func (n *NettingCutOffReferenceDataUpdateRequestV01) AddNettingCutOffRequest() *iso20022.NettingCutOff1 {
	newValue := new(iso20022.NettingCutOff1)
	n.NettingCutOffRequest = append(n.NettingCutOffRequest, newValue)
	return newValue
}

func (n *NettingCutOffReferenceDataUpdateRequestV01) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	n.SupplementaryData = append(n.SupplementaryData, newValue)
	return newValue
}
