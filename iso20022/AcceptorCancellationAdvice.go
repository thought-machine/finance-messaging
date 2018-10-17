package iso20022

// Cancellation transaction between an acceptor and an acquirer.
type AcceptorCancellationAdvice1 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment2 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext2 `xml:"Cntxt,omitempty"`

	// Cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction7 `xml:"Tx"`
}

func (a *AcceptorCancellationAdvice1) AddEnvironment() *CardPaymentEnvironment2 {
	a.Environment = new(CardPaymentEnvironment2)
	return a.Environment
}

func (a *AcceptorCancellationAdvice1) AddContext() *CardPaymentContext2 {
	a.Context = new(CardPaymentContext2)
	return a.Context
}

func (a *AcceptorCancellationAdvice1) AddTransaction() *CardPaymentTransaction7 {
	a.Transaction = new(CardPaymentTransaction7)
	return a.Transaction
}

// Cancellation transaction between an acceptor and an acquirer.
type AcceptorCancellationAdvice2 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment18 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext2 `xml:"Cntxt,omitempty"`

	// Cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction16 `xml:"Tx"`
}

func (a *AcceptorCancellationAdvice2) AddEnvironment() *CardPaymentEnvironment18 {
	a.Environment = new(CardPaymentEnvironment18)
	return a.Environment
}

func (a *AcceptorCancellationAdvice2) AddContext() *CardPaymentContext2 {
	a.Context = new(CardPaymentContext2)
	return a.Context
}

func (a *AcceptorCancellationAdvice2) AddTransaction() *CardPaymentTransaction16 {
	a.Transaction = new(CardPaymentTransaction16)
	return a.Transaction
}

// Cancellation transaction between an acceptor and an acquirer.
type AcceptorCancellationAdvice3 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment24 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext2 `xml:"Cntxt,omitempty"`

	// Cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction28 `xml:"Tx"`
}

func (a *AcceptorCancellationAdvice3) AddEnvironment() *CardPaymentEnvironment24 {
	a.Environment = new(CardPaymentEnvironment24)
	return a.Environment
}

func (a *AcceptorCancellationAdvice3) AddContext() *CardPaymentContext2 {
	a.Context = new(CardPaymentContext2)
	return a.Context
}

func (a *AcceptorCancellationAdvice3) AddTransaction() *CardPaymentTransaction28 {
	a.Transaction = new(CardPaymentTransaction28)
	return a.Transaction
}

// Cancellation transaction between an acceptor and an acquirer.
type AcceptorCancellationAdvice4 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment36 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext11 `xml:"Cntxt,omitempty"`

	// Cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction44 `xml:"Tx"`
}

func (a *AcceptorCancellationAdvice4) AddEnvironment() *CardPaymentEnvironment36 {
	a.Environment = new(CardPaymentEnvironment36)
	return a.Environment
}

func (a *AcceptorCancellationAdvice4) AddContext() *CardPaymentContext11 {
	a.Context = new(CardPaymentContext11)
	return a.Context
}

func (a *AcceptorCancellationAdvice4) AddTransaction() *CardPaymentTransaction44 {
	a.Transaction = new(CardPaymentTransaction44)
	return a.Transaction
}

// Cancellation transaction between an acceptor and an acquirer.
type AcceptorCancellationAdvice5 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment49 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext17 `xml:"Cntxt,omitempty"`

	// Cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction59 `xml:"Tx"`
}

func (a *AcceptorCancellationAdvice5) AddEnvironment() *CardPaymentEnvironment49 {
	a.Environment = new(CardPaymentEnvironment49)
	return a.Environment
}

func (a *AcceptorCancellationAdvice5) AddContext() *CardPaymentContext17 {
	a.Context = new(CardPaymentContext17)
	return a.Context
}

func (a *AcceptorCancellationAdvice5) AddTransaction() *CardPaymentTransaction59 {
	a.Transaction = new(CardPaymentTransaction59)
	return a.Transaction
}
