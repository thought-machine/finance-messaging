package iso20022

// Authorisation request from an acceptor.
type AcceptorAuthorisationRequest1 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment1 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext1 `xml:"Cntxt"`

	// Card payment transaction for which the authorisation is requested.
	Transaction *CardPaymentTransaction1 `xml:"Tx"`
}

func (a *AcceptorAuthorisationRequest1) AddEnvironment() *CardPaymentEnvironment1 {
	a.Environment = new(CardPaymentEnvironment1)
	return a.Environment
}

func (a *AcceptorAuthorisationRequest1) AddContext() *CardPaymentContext1 {
	a.Context = new(CardPaymentContext1)
	return a.Context
}

func (a *AcceptorAuthorisationRequest1) AddTransaction() *CardPaymentTransaction1 {
	a.Transaction = new(CardPaymentTransaction1)
	return a.Transaction
}

// Authorisation request from an acceptor.
type AcceptorAuthorisationRequest2 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment9 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext1 `xml:"Cntxt"`

	// Card payment transaction for which the authorisation is requested.
	Transaction *CardPaymentTransaction11 `xml:"Tx"`
}

func (a *AcceptorAuthorisationRequest2) AddEnvironment() *CardPaymentEnvironment9 {
	a.Environment = new(CardPaymentEnvironment9)
	return a.Environment
}

func (a *AcceptorAuthorisationRequest2) AddContext() *CardPaymentContext1 {
	a.Context = new(CardPaymentContext1)
	return a.Context
}

func (a *AcceptorAuthorisationRequest2) AddTransaction() *CardPaymentTransaction11 {
	a.Transaction = new(CardPaymentTransaction11)
	return a.Transaction
}

// Authorisation request from an acceptor.
type AcceptorAuthorisationRequest3 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment20 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext5 `xml:"Cntxt"`

	// Card payment transaction for which the authorisation is requested.
	Transaction *CardPaymentTransaction22 `xml:"Tx"`
}

func (a *AcceptorAuthorisationRequest3) AddEnvironment() *CardPaymentEnvironment20 {
	a.Environment = new(CardPaymentEnvironment20)
	return a.Environment
}

func (a *AcceptorAuthorisationRequest3) AddContext() *CardPaymentContext5 {
	a.Context = new(CardPaymentContext5)
	return a.Context
}

func (a *AcceptorAuthorisationRequest3) AddTransaction() *CardPaymentTransaction22 {
	a.Transaction = new(CardPaymentTransaction22)
	return a.Transaction
}

// Authorisation request from an acceptor.
type AcceptorAuthorisationRequest4 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment32 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext8 `xml:"Cntxt"`

	// Card payment transaction for which the authorisation is requested.
	Transaction *CardPaymentTransaction36 `xml:"Tx"`

	// Additional information incorporated as an extension to the message.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AcceptorAuthorisationRequest4) AddEnvironment() *CardPaymentEnvironment32 {
	a.Environment = new(CardPaymentEnvironment32)
	return a.Environment
}

func (a *AcceptorAuthorisationRequest4) AddContext() *CardPaymentContext8 {
	a.Context = new(CardPaymentContext8)
	return a.Context
}

func (a *AcceptorAuthorisationRequest4) AddTransaction() *CardPaymentTransaction36 {
	a.Transaction = new(CardPaymentTransaction36)
	return a.Transaction
}

func (a *AcceptorAuthorisationRequest4) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}

// Authorisation request from an acceptor.
type AcceptorAuthorisationRequest5 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment45 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext14 `xml:"Cntxt"`

	// Card payment transaction for which the authorisation is requested.
	Transaction *CardPaymentTransaction51 `xml:"Tx"`

	// Additional information incorporated as an extension to the message.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AcceptorAuthorisationRequest5) AddEnvironment() *CardPaymentEnvironment45 {
	a.Environment = new(CardPaymentEnvironment45)
	return a.Environment
}

func (a *AcceptorAuthorisationRequest5) AddContext() *CardPaymentContext14 {
	a.Context = new(CardPaymentContext14)
	return a.Context
}

func (a *AcceptorAuthorisationRequest5) AddTransaction() *CardPaymentTransaction51 {
	a.Transaction = new(CardPaymentTransaction51)
	return a.Transaction
}

func (a *AcceptorAuthorisationRequest5) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}
