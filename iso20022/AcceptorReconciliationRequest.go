package iso20022

// Reconciliation request from an acceptor.
type AcceptorReconciliationRequest1 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment7 `xml:"Envt"`

	// Reconciliation transaction between an acceptor an acquirer.
	Transaction *TransactionReconciliation1 `xml:"Tx"`
}

func (a *AcceptorReconciliationRequest1) AddEnvironment() *CardPaymentEnvironment7 {
	a.Environment = new(CardPaymentEnvironment7)
	return a.Environment
}

func (a *AcceptorReconciliationRequest1) AddTransaction() *TransactionReconciliation1 {
	a.Transaction = new(TransactionReconciliation1)
	return a.Transaction
}

// Reconciliation request from an acceptor.
type AcceptorReconciliationRequest2 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment15 `xml:"Envt"`

	// Reconciliation transaction between an acceptor an acquirer.
	Transaction *TransactionReconciliation2 `xml:"Tx"`
}

func (a *AcceptorReconciliationRequest2) AddEnvironment() *CardPaymentEnvironment15 {
	a.Environment = new(CardPaymentEnvironment15)
	return a.Environment
}

func (a *AcceptorReconciliationRequest2) AddTransaction() *TransactionReconciliation2 {
	a.Transaction = new(TransactionReconciliation2)
	return a.Transaction
}

// Reconciliation request from an acceptor.
type AcceptorReconciliationRequest3 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment25 `xml:"Envt"`

	// Reconciliation transaction between an acceptor an acquirer.
	Transaction *TransactionReconciliation2 `xml:"Tx"`
}

func (a *AcceptorReconciliationRequest3) AddEnvironment() *CardPaymentEnvironment25 {
	a.Environment = new(CardPaymentEnvironment25)
	return a.Environment
}

func (a *AcceptorReconciliationRequest3) AddTransaction() *TransactionReconciliation2 {
	a.Transaction = new(TransactionReconciliation2)
	return a.Transaction
}

// Reconciliation request from an acceptor.
type AcceptorReconciliationRequest4 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment37 `xml:"Envt"`

	// Reconciliation transaction between an acceptor an acquirer.
	Transaction *TransactionReconciliation3 `xml:"Tx"`
}

func (a *AcceptorReconciliationRequest4) AddEnvironment() *CardPaymentEnvironment37 {
	a.Environment = new(CardPaymentEnvironment37)
	return a.Environment
}

func (a *AcceptorReconciliationRequest4) AddTransaction() *TransactionReconciliation3 {
	a.Transaction = new(TransactionReconciliation3)
	return a.Transaction
}

// Reconciliation request from an acceptor.
type AcceptorReconciliationRequest5 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment50 `xml:"Envt"`

	// Reconciliation transaction between an acceptor an acquirer.
	Transaction *TransactionReconciliation4 `xml:"Tx"`
}

func (a *AcceptorReconciliationRequest5) AddEnvironment() *CardPaymentEnvironment50 {
	a.Environment = new(CardPaymentEnvironment50)
	return a.Environment
}

func (a *AcceptorReconciliationRequest5) AddTransaction() *TransactionReconciliation4 {
	a.Transaction = new(TransactionReconciliation4)
	return a.Transaction
}
