package iso20022

// Cancellation response from the acquirer.
type AcceptorCancellationResponse1 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment3 `xml:"Envt"`

	// Cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction6 `xml:"Tx"`

	// Response from the acquirer to the cancellation transaction.
	TransactionResponse *CardPaymentTransaction10 `xml:"TxRspn"`
}

func (a *AcceptorCancellationResponse1) AddEnvironment() *CardPaymentEnvironment3 {
	a.Environment = new(CardPaymentEnvironment3)
	return a.Environment
}

func (a *AcceptorCancellationResponse1) AddTransaction() *CardPaymentTransaction6 {
	a.Transaction = new(CardPaymentTransaction6)
	return a.Transaction
}

func (a *AcceptorCancellationResponse1) AddTransactionResponse() *CardPaymentTransaction10 {
	a.TransactionResponse = new(CardPaymentTransaction10)
	return a.TransactionResponse
}

// Cancellation response from the acquirer.
type AcceptorCancellationResponse2 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment11 `xml:"Envt"`

	// Cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction6 `xml:"Tx"`

	// Response from the acquirer to the cancellation transaction.
	TransactionResponse *CardPaymentTransaction10 `xml:"TxRspn"`
}

func (a *AcceptorCancellationResponse2) AddEnvironment() *CardPaymentEnvironment11 {
	a.Environment = new(CardPaymentEnvironment11)
	return a.Environment
}

func (a *AcceptorCancellationResponse2) AddTransaction() *CardPaymentTransaction6 {
	a.Transaction = new(CardPaymentTransaction6)
	return a.Transaction
}

func (a *AcceptorCancellationResponse2) AddTransactionResponse() *CardPaymentTransaction10 {
	a.TransactionResponse = new(CardPaymentTransaction10)
	return a.TransactionResponse
}

// Cancellation response from the acquirer.
type AcceptorCancellationResponse3 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment21 `xml:"Envt"`

	// Cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction35 `xml:"Tx"`

	// Response from the acquirer to the cancellation transaction.
	TransactionResponse *CardPaymentTransaction27 `xml:"TxRspn"`
}

func (a *AcceptorCancellationResponse3) AddEnvironment() *CardPaymentEnvironment21 {
	a.Environment = new(CardPaymentEnvironment21)
	return a.Environment
}

func (a *AcceptorCancellationResponse3) AddTransaction() *CardPaymentTransaction35 {
	a.Transaction = new(CardPaymentTransaction35)
	return a.Transaction
}

func (a *AcceptorCancellationResponse3) AddTransactionResponse() *CardPaymentTransaction27 {
	a.TransactionResponse = new(CardPaymentTransaction27)
	return a.TransactionResponse
}

// Cancellation response from the acquirer.
type AcceptorCancellationResponse4 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment33 `xml:"Envt"`

	// Cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction42 `xml:"Tx"`

	// Response from the acquirer to the cancellation transaction.
	TransactionResponse *CardPaymentTransaction43 `xml:"TxRspn"`
}

func (a *AcceptorCancellationResponse4) AddEnvironment() *CardPaymentEnvironment33 {
	a.Environment = new(CardPaymentEnvironment33)
	return a.Environment
}

func (a *AcceptorCancellationResponse4) AddTransaction() *CardPaymentTransaction42 {
	a.Transaction = new(CardPaymentTransaction42)
	return a.Transaction
}

func (a *AcceptorCancellationResponse4) AddTransactionResponse() *CardPaymentTransaction43 {
	a.TransactionResponse = new(CardPaymentTransaction43)
	return a.TransactionResponse
}

// Cancellation response from the acquirer.
type AcceptorCancellationResponse5 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment46 `xml:"Envt"`

	// Cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction57 `xml:"Tx"`

	// Response from the acquirer to the cancellation transaction.
	TransactionResponse *CardPaymentTransaction58 `xml:"TxRspn"`
}

func (a *AcceptorCancellationResponse5) AddEnvironment() *CardPaymentEnvironment46 {
	a.Environment = new(CardPaymentEnvironment46)
	return a.Environment
}

func (a *AcceptorCancellationResponse5) AddTransaction() *CardPaymentTransaction57 {
	a.Transaction = new(CardPaymentTransaction57)
	return a.Transaction
}

func (a *AcceptorCancellationResponse5) AddTransactionResponse() *CardPaymentTransaction58 {
	a.TransactionResponse = new(CardPaymentTransaction58)
	return a.TransactionResponse
}
