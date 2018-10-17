package iso20022

// Cancellation request from an acceptor.
type AcceptorCancellationRequest1 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment4 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext1 `xml:"Cntxt"`

	// Cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction5 `xml:"Tx"`
}

func (a *AcceptorCancellationRequest1) AddEnvironment() *CardPaymentEnvironment4 {
	a.Environment = new(CardPaymentEnvironment4)
	return a.Environment
}

func (a *AcceptorCancellationRequest1) AddContext() *CardPaymentContext1 {
	a.Context = new(CardPaymentContext1)
	return a.Context
}

func (a *AcceptorCancellationRequest1) AddTransaction() *CardPaymentTransaction5 {
	a.Transaction = new(CardPaymentTransaction5)
	return a.Transaction
}

// Cancellation request from an acceptor.
type AcceptorCancellationRequest2 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment12 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext1 `xml:"Cntxt"`

	// Cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction15 `xml:"Tx"`
}

func (a *AcceptorCancellationRequest2) AddEnvironment() *CardPaymentEnvironment12 {
	a.Environment = new(CardPaymentEnvironment12)
	return a.Environment
}

func (a *AcceptorCancellationRequest2) AddContext() *CardPaymentContext1 {
	a.Context = new(CardPaymentContext1)
	return a.Context
}

func (a *AcceptorCancellationRequest2) AddTransaction() *CardPaymentTransaction15 {
	a.Transaction = new(CardPaymentTransaction15)
	return a.Transaction
}

// Cancellation request from an acceptor.
type AcceptorCancellationRequest3 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment23 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext1 `xml:"Cntxt"`

	// Cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction26 `xml:"Tx"`
}

func (a *AcceptorCancellationRequest3) AddEnvironment() *CardPaymentEnvironment23 {
	a.Environment = new(CardPaymentEnvironment23)
	return a.Environment
}

func (a *AcceptorCancellationRequest3) AddContext() *CardPaymentContext1 {
	a.Context = new(CardPaymentContext1)
	return a.Context
}

func (a *AcceptorCancellationRequest3) AddTransaction() *CardPaymentTransaction26 {
	a.Transaction = new(CardPaymentTransaction26)
	return a.Transaction
}

// Cancellation request from an acceptor.
type AcceptorCancellationRequest4 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment35 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext10 `xml:"Cntxt"`

	// Cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction41 `xml:"Tx"`
}

func (a *AcceptorCancellationRequest4) AddEnvironment() *CardPaymentEnvironment35 {
	a.Environment = new(CardPaymentEnvironment35)
	return a.Environment
}

func (a *AcceptorCancellationRequest4) AddContext() *CardPaymentContext10 {
	a.Context = new(CardPaymentContext10)
	return a.Context
}

func (a *AcceptorCancellationRequest4) AddTransaction() *CardPaymentTransaction41 {
	a.Transaction = new(CardPaymentTransaction41)
	return a.Transaction
}

// Cancellation request from an acceptor.
type AcceptorCancellationRequest5 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment48 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext16 `xml:"Cntxt"`

	// Cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction56 `xml:"Tx"`
}

func (a *AcceptorCancellationRequest5) AddEnvironment() *CardPaymentEnvironment48 {
	a.Environment = new(CardPaymentEnvironment48)
	return a.Environment
}

func (a *AcceptorCancellationRequest5) AddContext() *CardPaymentContext16 {
	a.Context = new(CardPaymentContext16)
	return a.Context
}

func (a *AcceptorCancellationRequest5) AddTransaction() *CardPaymentTransaction56 {
	a.Transaction = new(CardPaymentTransaction56)
	return a.Transaction
}
