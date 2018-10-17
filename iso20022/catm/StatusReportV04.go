package catm

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document00100104 struct {
	XMLName xml.Name         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.04 Document"`
	Message *StatusReportV04 `xml:"StsRpt"`
}

func (d *Document00100104) AddMessage() *StatusReportV04 {
	d.Message = new(StatusReportV04)
	return d.Message
}

// Informs the master terminal manager (MTM) or the terminal manager (TM) about the status of the acceptor system including the identification of the POI, its components and their installed versions.
type StatusReportV04 struct {

	// Set of characteristics related to the transfer of the status report.
	Header *iso20022.Header14 `xml:"Hdr"`

	// Status of the point of interaction (POI), its components and their installed versions.
	StatusReport *iso20022.StatusReport4 `xml:"StsRpt"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *iso20022.ContentInformationType12 `xml:"SctyTrlr"`
}

func (s *StatusReportV04) AddHeader() *iso20022.Header14 {
	s.Header = new(iso20022.Header14)
	return s.Header
}

func (s *StatusReportV04) AddStatusReport() *iso20022.StatusReport4 {
	s.StatusReport = new(iso20022.StatusReport4)
	return s.StatusReport
}

func (s *StatusReportV04) AddSecurityTrailer() *iso20022.ContentInformationType12 {
	s.SecurityTrailer = new(iso20022.ContentInformationType12)
	return s.SecurityTrailer
}
