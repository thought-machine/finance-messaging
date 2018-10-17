package iso20022

// Acquirer configuration parameters for a host.
type AcquirerHostConfiguration1 struct {

	// Identification of a host.
	HostIdentification *Max35Text `xml:"HstId"`

	// Types of message to sent to this host.
	MessageToSend []*MessageFunction3Code `xml:"MsgToSnd"`
}

func (a *AcquirerHostConfiguration1) SetHostIdentification(value string) {
	a.HostIdentification = (*Max35Text)(&value)
}

func (a *AcquirerHostConfiguration1) AddMessageToSend(value string) {
	a.MessageToSend = append(a.MessageToSend, (*MessageFunction3Code)(&value))
}

// Acquirer configuration parameters for a host.
type AcquirerHostConfiguration2 struct {

	// Identification of a host.
	HostIdentification *Max35Text `xml:"HstId"`

	// Types of message to sent to this host.
	MessageToSend []*MessageFunction3Code `xml:"MsgToSnd,omitempty"`
}

func (a *AcquirerHostConfiguration2) SetHostIdentification(value string) {
	a.HostIdentification = (*Max35Text)(&value)
}

func (a *AcquirerHostConfiguration2) AddMessageToSend(value string) {
	a.MessageToSend = append(a.MessageToSend, (*MessageFunction3Code)(&value))
}

// Acquirer configuration parameters for a host.
type AcquirerHostConfiguration3 struct {

	// Identification of a host.
	HostIdentification *Max35Text `xml:"HstId"`

	// Types of message to sent to this host.
	MessageToSend []*MessageFunction5Code `xml:"MsgToSnd,omitempty"`
}

func (a *AcquirerHostConfiguration3) SetHostIdentification(value string) {
	a.HostIdentification = (*Max35Text)(&value)
}

func (a *AcquirerHostConfiguration3) AddMessageToSend(value string) {
	a.MessageToSend = append(a.MessageToSend, (*MessageFunction5Code)(&value))
}
