package iso20022

// Cancellation advice response from the acquirer.
type AcceptorCancellationAdviceResponse1 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment3 `xml:"Envt"`

	// Cancellation transaction from an acceptor to the acquirer.
	Transaction *CardPaymentTransactionAdviceResponse2 `xml:"Tx"`

	// Instructions for contacting the terminal management host.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`
}

func (a *AcceptorCancellationAdviceResponse1) AddEnvironment() *CardPaymentEnvironment3 {
	a.Environment = new(CardPaymentEnvironment3)
	return a.Environment
}

func (a *AcceptorCancellationAdviceResponse1) AddTransaction() *CardPaymentTransactionAdviceResponse2 {
	a.Transaction = new(CardPaymentTransactionAdviceResponse2)
	return a.Transaction
}

func (a *AcceptorCancellationAdviceResponse1) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}

// Cancellation advice response from the acquirer.
type AcceptorCancellationAdviceResponse2 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment11 `xml:"Envt"`

	// Cancellation transaction from an acceptor to the acquirer.
	Transaction *CardPaymentTransactionAdviceResponse3 `xml:"Tx"`

	// Instructions for contacting the terminal management host.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`
}

func (a *AcceptorCancellationAdviceResponse2) AddEnvironment() *CardPaymentEnvironment11 {
	a.Environment = new(CardPaymentEnvironment11)
	return a.Environment
}

func (a *AcceptorCancellationAdviceResponse2) AddTransaction() *CardPaymentTransactionAdviceResponse3 {
	a.Transaction = new(CardPaymentTransactionAdviceResponse3)
	return a.Transaction
}

func (a *AcceptorCancellationAdviceResponse2) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}

// Cancellation advice response from the acquirer.
type AcceptorCancellationAdviceResponse3 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment21 `xml:"Envt"`

	// Cancellation transaction from an acceptor to the acquirer.
	Transaction *CardPaymentTransactionAdviceResponse4 `xml:"Tx"`

	// Instructions for contacting the terminal management host.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`
}

func (a *AcceptorCancellationAdviceResponse3) AddEnvironment() *CardPaymentEnvironment21 {
	a.Environment = new(CardPaymentEnvironment21)
	return a.Environment
}

func (a *AcceptorCancellationAdviceResponse3) AddTransaction() *CardPaymentTransactionAdviceResponse4 {
	a.Transaction = new(CardPaymentTransactionAdviceResponse4)
	return a.Transaction
}

func (a *AcceptorCancellationAdviceResponse3) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}

// Cancellation advice response from the acquirer.
type AcceptorCancellationAdviceResponse4 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment33 `xml:"Envt"`

	// Cancellation transaction from an acceptor to the acquirer.
	Transaction *CardPaymentTransactionAdviceResponse4 `xml:"Tx"`

	// Instructions for contacting the terminal management host.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`
}

func (a *AcceptorCancellationAdviceResponse4) AddEnvironment() *CardPaymentEnvironment33 {
	a.Environment = new(CardPaymentEnvironment33)
	return a.Environment
}

func (a *AcceptorCancellationAdviceResponse4) AddTransaction() *CardPaymentTransactionAdviceResponse4 {
	a.Transaction = new(CardPaymentTransactionAdviceResponse4)
	return a.Transaction
}

func (a *AcceptorCancellationAdviceResponse4) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}

// Cancellation advice response from the acquirer.
type AcceptorCancellationAdviceResponse5 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment46 `xml:"Envt"`

	// Cancellation transaction from an acceptor to the acquirer.
	Transaction *CardPaymentTransactionAdviceResponse6 `xml:"Tx"`

	// Instructions for contacting the terminal management host.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`
}

func (a *AcceptorCancellationAdviceResponse5) AddEnvironment() *CardPaymentEnvironment46 {
	a.Environment = new(CardPaymentEnvironment46)
	return a.Environment
}

func (a *AcceptorCancellationAdviceResponse5) AddTransaction() *CardPaymentTransactionAdviceResponse6 {
	a.Transaction = new(CardPaymentTransactionAdviceResponse6)
	return a.Transaction
}

func (a *AcceptorCancellationAdviceResponse5) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}
