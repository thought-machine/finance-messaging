package iso20022

// Reconciliation response from the acquirer.
type AcceptorReconciliationResponse1 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment7 `xml:"Envt"`

	// Response from the acquirer to the reconciliation transaction.
	TransactionResponse *ResponseType1 `xml:"TxRspn"`

	// Reconciliation transaction between an acceptor an acquirer.
	Transaction *TransactionReconciliation1 `xml:"Tx"`
}

func (a *AcceptorReconciliationResponse1) AddEnvironment() *CardPaymentEnvironment7 {
	a.Environment = new(CardPaymentEnvironment7)
	return a.Environment
}

func (a *AcceptorReconciliationResponse1) AddTransactionResponse() *ResponseType1 {
	a.TransactionResponse = new(ResponseType1)
	return a.TransactionResponse
}

func (a *AcceptorReconciliationResponse1) AddTransaction() *TransactionReconciliation1 {
	a.Transaction = new(TransactionReconciliation1)
	return a.Transaction
}

// Reconciliation response from the acquirer.
type AcceptorReconciliationResponse2 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment19 `xml:"Envt"`

	// Response from the acquirer to the reconciliation transaction.
	TransactionResponse *ResponseType1 `xml:"TxRspn"`

	// Reconciliation transaction between an acceptor an acquirer.
	Transaction *TransactionReconciliation2 `xml:"Tx"`
}

func (a *AcceptorReconciliationResponse2) AddEnvironment() *CardPaymentEnvironment19 {
	a.Environment = new(CardPaymentEnvironment19)
	return a.Environment
}

func (a *AcceptorReconciliationResponse2) AddTransactionResponse() *ResponseType1 {
	a.TransactionResponse = new(ResponseType1)
	return a.TransactionResponse
}

func (a *AcceptorReconciliationResponse2) AddTransaction() *TransactionReconciliation2 {
	a.Transaction = new(TransactionReconciliation2)
	return a.Transaction
}

// Reconciliation response from the acquirer.
type AcceptorReconciliationResponse3 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment38 `xml:"Envt"`

	// Response from the acquirer to the reconciliation transaction.
	TransactionResponse *ResponseType1 `xml:"TxRspn"`

	// Reconciliation transaction between an acceptor an acquirer.
	Transaction *TransactionReconciliation3 `xml:"Tx"`
}

func (a *AcceptorReconciliationResponse3) AddEnvironment() *CardPaymentEnvironment38 {
	a.Environment = new(CardPaymentEnvironment38)
	return a.Environment
}

func (a *AcceptorReconciliationResponse3) AddTransactionResponse() *ResponseType1 {
	a.TransactionResponse = new(ResponseType1)
	return a.TransactionResponse
}

func (a *AcceptorReconciliationResponse3) AddTransaction() *TransactionReconciliation3 {
	a.Transaction = new(TransactionReconciliation3)
	return a.Transaction
}

// Reconciliation response from the acquirer.
type AcceptorReconciliationResponse4 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment38 `xml:"Envt"`

	// Response from the acquirer to the reconciliation transaction.
	TransactionResponse *ResponseType5 `xml:"TxRspn"`

	// Reconciliation transaction between an acceptor an acquirer.
	Transaction *TransactionReconciliation4 `xml:"Tx"`
}

func (a *AcceptorReconciliationResponse4) AddEnvironment() *CardPaymentEnvironment38 {
	a.Environment = new(CardPaymentEnvironment38)
	return a.Environment
}

func (a *AcceptorReconciliationResponse4) AddTransactionResponse() *ResponseType5 {
	a.TransactionResponse = new(ResponseType5)
	return a.TransactionResponse
}

func (a *AcceptorReconciliationResponse4) AddTransaction() *TransactionReconciliation4 {
	a.Transaction = new(TransactionReconciliation4)
	return a.Transaction
}
