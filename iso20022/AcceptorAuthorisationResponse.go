package iso20022

// Authorisation response from the acquirer.
type AcceptorAuthorisationResponse1 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment3 `xml:"Envt"`

	// Authorisation of a card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction2 `xml:"Tx"`

	// Authorisation response from the acquirer.
	// Authorisation of a card payment transaction between an acceptor and an acquirer.
	TransactionResponse *CardPaymentTransaction9 `xml:"TxRspn"`
}

func (a *AcceptorAuthorisationResponse1) AddEnvironment() *CardPaymentEnvironment3 {
	a.Environment = new(CardPaymentEnvironment3)
	return a.Environment
}

func (a *AcceptorAuthorisationResponse1) AddTransaction() *CardPaymentTransaction2 {
	a.Transaction = new(CardPaymentTransaction2)
	return a.Transaction
}

func (a *AcceptorAuthorisationResponse1) AddTransactionResponse() *CardPaymentTransaction9 {
	a.TransactionResponse = new(CardPaymentTransaction9)
	return a.TransactionResponse
}

// Authorisation response from the acquirer.
type AcceptorAuthorisationResponse2 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment11 `xml:"Envt"`

	// Authorisation of a card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction2 `xml:"Tx"`

	// Authorisation response from the acquirer.
	// Authorisation of a card payment transaction between an acceptor and an acquirer.
	TransactionResponse *CardPaymentTransaction18 `xml:"TxRspn"`
}

func (a *AcceptorAuthorisationResponse2) AddEnvironment() *CardPaymentEnvironment11 {
	a.Environment = new(CardPaymentEnvironment11)
	return a.Environment
}

func (a *AcceptorAuthorisationResponse2) AddTransaction() *CardPaymentTransaction2 {
	a.Transaction = new(CardPaymentTransaction2)
	return a.Transaction
}

func (a *AcceptorAuthorisationResponse2) AddTransactionResponse() *CardPaymentTransaction18 {
	a.TransactionResponse = new(CardPaymentTransaction18)
	return a.TransactionResponse
}

// Authorisation response from the acquirer.
type AcceptorAuthorisationResponse3 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment21 `xml:"Envt"`

	// Authorisation of a card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction23 `xml:"Tx"`

	// Authorisation response from the acquirer.
	// Authorisation of a card payment transaction between an acceptor and an acquirer.
	TransactionResponse *CardPaymentTransaction24 `xml:"TxRspn"`
}

func (a *AcceptorAuthorisationResponse3) AddEnvironment() *CardPaymentEnvironment21 {
	a.Environment = new(CardPaymentEnvironment21)
	return a.Environment
}

func (a *AcceptorAuthorisationResponse3) AddTransaction() *CardPaymentTransaction23 {
	a.Transaction = new(CardPaymentTransaction23)
	return a.Transaction
}

func (a *AcceptorAuthorisationResponse3) AddTransactionResponse() *CardPaymentTransaction24 {
	a.TransactionResponse = new(CardPaymentTransaction24)
	return a.TransactionResponse
}

// Authorisation response from the acquirer.
type AcceptorAuthorisationResponse4 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment33 `xml:"Envt"`

	// Authorisation of a card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction38 `xml:"Tx"`

	// Authorisation response from the acquirer.
	// Authorisation of a card payment transaction between an acceptor and an acquirer.
	TransactionResponse *CardPaymentTransaction39 `xml:"TxRspn"`

	// Additional information incorporated as an extension to the message.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AcceptorAuthorisationResponse4) AddEnvironment() *CardPaymentEnvironment33 {
	a.Environment = new(CardPaymentEnvironment33)
	return a.Environment
}

func (a *AcceptorAuthorisationResponse4) AddTransaction() *CardPaymentTransaction38 {
	a.Transaction = new(CardPaymentTransaction38)
	return a.Transaction
}

func (a *AcceptorAuthorisationResponse4) AddTransactionResponse() *CardPaymentTransaction39 {
	a.TransactionResponse = new(CardPaymentTransaction39)
	return a.TransactionResponse
}

func (a *AcceptorAuthorisationResponse4) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}

// Authorisation response from the acquirer.
type AcceptorAuthorisationResponse5 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment46 `xml:"Envt"`

	// Authorisation of a card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction53 `xml:"Tx"`

	// Authorisation response from the acquirer.
	// Authorisation of a card payment transaction between an acceptor and an acquirer.
	TransactionResponse *CardPaymentTransaction54 `xml:"TxRspn"`

	// Additional information incorporated as an extension to the message.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AcceptorAuthorisationResponse5) AddEnvironment() *CardPaymentEnvironment46 {
	a.Environment = new(CardPaymentEnvironment46)
	return a.Environment
}

func (a *AcceptorAuthorisationResponse5) AddTransaction() *CardPaymentTransaction53 {
	a.Transaction = new(CardPaymentTransaction53)
	return a.Transaction
}

func (a *AcceptorAuthorisationResponse5) AddTransactionResponse() *CardPaymentTransaction54 {
	a.TransactionResponse = new(CardPaymentTransaction54)
	return a.TransactionResponse
}

func (a *AcceptorAuthorisationResponse5) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}
