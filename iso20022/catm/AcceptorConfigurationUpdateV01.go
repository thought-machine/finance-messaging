package catm

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document00300101 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.003.001.01 Document"`
	Message *AcceptorConfigurationUpdateV01 `xml:"AccptrCfgtnUpd"`
}

func (d *Document00300101) AddMessage() *AcceptorConfigurationUpdateV01 {
	d.Message = new(AcceptorConfigurationUpdateV01)
	return d.Message
}

// Scope
// The AcceptorConfigurationUpdate message is sent by the master terminal manager or delegated terminal manager to the acceptor system for the update of acquirer parameters, merchant parameters, vendor parameters or cryptographic keys of the acquirer.
// Usage
// The AcceptorConfigurationUpdate message may embed the information required by the acceptor system for the configuration of:
// - the application parameters necessary for software applications processed by the POI system,
// - the acquirer protocol parameters for the message content and message exchange behaviour of the acquirer protocol supported by the POI system,
// - the host communication parameters to define the addresses of the connected acquirer hosts, and
// - the merchant parameters needed for the retailer protocol settings of the POI system.
type AcceptorConfigurationUpdateV01 struct {

	// Set of characteristics related to the transfer of the acceptor parameters.
	Header *iso20022.Header4 `xml:"Hdr"`

	// Acceptor configuration to be downloaded from the terminal management system.
	AcceptorConfiguration *iso20022.AcceptorConfiguration1 `xml:"AccptrCfgtn"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *iso20022.ContentInformationType1 `xml:"SctyTrlr"`
}

func (a *AcceptorConfigurationUpdateV01) AddHeader() *iso20022.Header4 {
	a.Header = new(iso20022.Header4)
	return a.Header
}

func (a *AcceptorConfigurationUpdateV01) AddAcceptorConfiguration() *iso20022.AcceptorConfiguration1 {
	a.AcceptorConfiguration = new(iso20022.AcceptorConfiguration1)
	return a.AcceptorConfiguration
}

func (a *AcceptorConfigurationUpdateV01) AddSecurityTrailer() *iso20022.ContentInformationType1 {
	a.SecurityTrailer = new(iso20022.ContentInformationType1)
	return a.SecurityTrailer
}
