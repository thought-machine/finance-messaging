package iso20022

// Diagnostic request from an acceptor.
type AcceptorDiagnosticRequest1 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment8 `xml:"Envt"`
}

func (a *AcceptorDiagnosticRequest1) AddEnvironment() *CardPaymentEnvironment8 {
	a.Environment = new(CardPaymentEnvironment8)
	return a.Environment
}

// Diagnostic request from an acceptor.
type AcceptorDiagnosticRequest2 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment17 `xml:"Envt"`
}

func (a *AcceptorDiagnosticRequest2) AddEnvironment() *CardPaymentEnvironment17 {
	a.Environment = new(CardPaymentEnvironment17)
	return a.Environment
}

// Diagnostic request from an acceptor.
type AcceptorDiagnosticRequest3 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment29 `xml:"Envt"`
}

func (a *AcceptorDiagnosticRequest3) AddEnvironment() *CardPaymentEnvironment29 {
	a.Environment = new(CardPaymentEnvironment29)
	return a.Environment
}

// Diagnostic request from an acceptor.
type AcceptorDiagnosticRequest4 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment42 `xml:"Envt"`
}

func (a *AcceptorDiagnosticRequest4) AddEnvironment() *CardPaymentEnvironment42 {
	a.Environment = new(CardPaymentEnvironment42)
	return a.Environment
}

// Diagnostic request from an acceptor.
type AcceptorDiagnosticRequest5 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment55 `xml:"Envt"`
}

func (a *AcceptorDiagnosticRequest5) AddEnvironment() *CardPaymentEnvironment55 {
	a.Environment = new(CardPaymentEnvironment55)
	return a.Environment
}
