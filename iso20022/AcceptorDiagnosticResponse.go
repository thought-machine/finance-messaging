package iso20022

// Diagnostic response from the acquirer.
type AcceptorDiagnosticResponse1 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment8 `xml:"Envt"`

	// Instructions for contacting the terminal management host.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`
}

func (a *AcceptorDiagnosticResponse1) AddEnvironment() *CardPaymentEnvironment8 {
	a.Environment = new(CardPaymentEnvironment8)
	return a.Environment
}

func (a *AcceptorDiagnosticResponse1) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}

// Diagnostic response from the acquirer.
type AcceptorDiagnosticResponse2 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment17 `xml:"Envt"`

	// Instructions for contacting the terminal management host.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`
}

func (a *AcceptorDiagnosticResponse2) AddEnvironment() *CardPaymentEnvironment17 {
	a.Environment = new(CardPaymentEnvironment17)
	return a.Environment
}

func (a *AcceptorDiagnosticResponse2) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}

// Diagnostic response from the acquirer.
type AcceptorDiagnosticResponse3 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment31 `xml:"Envt"`

	// Instructions for contacting the terminal management host.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`
}

func (a *AcceptorDiagnosticResponse3) AddEnvironment() *CardPaymentEnvironment31 {
	a.Environment = new(CardPaymentEnvironment31)
	return a.Environment
}

func (a *AcceptorDiagnosticResponse3) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}

// Diagnostic response from the acquirer.
type AcceptorDiagnosticResponse4 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment43 `xml:"Envt"`

	// Instructions for contacting the terminal management host.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`
}

func (a *AcceptorDiagnosticResponse4) AddEnvironment() *CardPaymentEnvironment43 {
	a.Environment = new(CardPaymentEnvironment43)
	return a.Environment
}

func (a *AcceptorDiagnosticResponse4) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}
