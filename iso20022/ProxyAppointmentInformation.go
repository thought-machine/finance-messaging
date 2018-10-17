package iso20022

// Indicates how to register a proxy.
type ProxyAppointmentInformation1 struct {

	// Indicates how to register a proxy.
	RegistrationMethod *Max350Text `xml:"RegnMtd,omitempty"`

	// Date by which the information on proxy assignment must be received by the intermediary.
	Deadline *DateFormat2Choice `xml:"Ddln,omitempty"`

	// Date by which the information on proxy assignment must be received by the intermediary (STP mode).
	STPDeadline *DateFormat2Choice `xml:"STPDdln,omitempty"`

	// Date by which the information on proxy assignment must be received by the issuer.
	MarketDeadline *DateFormat2Choice `xml:"MktDdln,omitempty"`

	// Specifies the proxy persons which are authorised by the issuer.
	AuthorisedProxy []*Proxy1 `xml:"AuthrsdPrxy,omitempty"`
}

func (p *ProxyAppointmentInformation1) SetRegistrationMethod(value string) {
	p.RegistrationMethod = (*Max350Text)(&value)
}

func (p *ProxyAppointmentInformation1) AddDeadline() *DateFormat2Choice {
	p.Deadline = new(DateFormat2Choice)
	return p.Deadline
}

func (p *ProxyAppointmentInformation1) AddSTPDeadline() *DateFormat2Choice {
	p.STPDeadline = new(DateFormat2Choice)
	return p.STPDeadline
}

func (p *ProxyAppointmentInformation1) AddMarketDeadline() *DateFormat2Choice {
	p.MarketDeadline = new(DateFormat2Choice)
	return p.MarketDeadline
}

func (p *ProxyAppointmentInformation1) AddAuthorisedProxy() *Proxy1 {
	newValue := new(Proxy1)
	p.AuthorisedProxy = append(p.AuthorisedProxy, newValue)
	return newValue
}

// Indicates how to register a proxy.
type ProxyAppointmentInformation2 struct {

	// Indicates how to register a proxy.
	RegistrationMethod *Max350Text `xml:"RegnMtd,omitempty"`

	// Date by which the information on proxy assignment must be received by the intermediary.
	Deadline *DateFormat2Choice `xml:"Ddln,omitempty"`

	// Date by which the information on proxy assignment must be received by the intermediary (STP mode).
	STPDeadline *DateFormat2Choice `xml:"STPDdln,omitempty"`

	// Date by which the information on proxy assignment must be received by the issuer.
	MarketDeadline *DateFormat2Choice `xml:"MktDdln,omitempty"`

	// Specifies the proxy persons which are authorised by the issuer.
	AuthorisedProxy []*Proxy3 `xml:"AuthrsdPrxy,omitempty"`
}

func (p *ProxyAppointmentInformation2) SetRegistrationMethod(value string) {
	p.RegistrationMethod = (*Max350Text)(&value)
}

func (p *ProxyAppointmentInformation2) AddDeadline() *DateFormat2Choice {
	p.Deadline = new(DateFormat2Choice)
	return p.Deadline
}

func (p *ProxyAppointmentInformation2) AddSTPDeadline() *DateFormat2Choice {
	p.STPDeadline = new(DateFormat2Choice)
	return p.STPDeadline
}

func (p *ProxyAppointmentInformation2) AddMarketDeadline() *DateFormat2Choice {
	p.MarketDeadline = new(DateFormat2Choice)
	return p.MarketDeadline
}

func (p *ProxyAppointmentInformation2) AddAuthorisedProxy() *Proxy3 {
	newValue := new(Proxy3)
	p.AuthorisedProxy = append(p.AuthorisedProxy, newValue)
	return newValue
}

// Indicates how a proxy is registered.
type ProxyAppointmentInformation3 struct {

	// Specifies how to register the proxy.
	RegistrationMethod *Max350Text `xml:"RegnMtd,omitempty"`

	// Date by which the information on the proxy assignment must be received by the intermediary.
	Deadline *DateFormat29Choice `xml:"Ddln,omitempty"`

	// Date by which the information on the proxy assignment must be received by the intermediary (STP mode).
	STPDeadline *DateFormat29Choice `xml:"STPDdln,omitempty"`

	// Date by which the information on the proxy assignment must be received by the issuer.
	MarketDeadline *DateFormat29Choice `xml:"MktDdln,omitempty"`

	// Specifies the proxy person that is authorised by the issuer.
	AuthorisedProxy []*Proxy5 `xml:"AuthrsdPrxy,omitempty"`
}

func (p *ProxyAppointmentInformation3) SetRegistrationMethod(value string) {
	p.RegistrationMethod = (*Max350Text)(&value)
}

func (p *ProxyAppointmentInformation3) AddDeadline() *DateFormat29Choice {
	p.Deadline = new(DateFormat29Choice)
	return p.Deadline
}

func (p *ProxyAppointmentInformation3) AddSTPDeadline() *DateFormat29Choice {
	p.STPDeadline = new(DateFormat29Choice)
	return p.STPDeadline
}

func (p *ProxyAppointmentInformation3) AddMarketDeadline() *DateFormat29Choice {
	p.MarketDeadline = new(DateFormat29Choice)
	return p.MarketDeadline
}

func (p *ProxyAppointmentInformation3) AddAuthorisedProxy() *Proxy5 {
	newValue := new(Proxy5)
	p.AuthorisedProxy = append(p.AuthorisedProxy, newValue)
	return newValue
}
