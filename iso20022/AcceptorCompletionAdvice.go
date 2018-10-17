package iso20022

// Notification to the acquirer of the completion of the card payment at the acceptor.
type AcceptorCompletionAdvice1 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment2 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext2 `xml:"Cntxt,omitempty"`

	// Card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction3 `xml:"Tx"`
}

func (a *AcceptorCompletionAdvice1) AddEnvironment() *CardPaymentEnvironment2 {
	a.Environment = new(CardPaymentEnvironment2)
	return a.Environment
}

func (a *AcceptorCompletionAdvice1) AddContext() *CardPaymentContext2 {
	a.Context = new(CardPaymentContext2)
	return a.Context
}

func (a *AcceptorCompletionAdvice1) AddTransaction() *CardPaymentTransaction3 {
	a.Transaction = new(CardPaymentTransaction3)
	return a.Transaction
}

// Notification to the acquirer of the completion of the card payment at the acceptor.
type AcceptorCompletionAdvice2 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment10 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext2 `xml:"Cntxt,omitempty"`

	// Card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction13 `xml:"Tx"`
}

func (a *AcceptorCompletionAdvice2) AddEnvironment() *CardPaymentEnvironment10 {
	a.Environment = new(CardPaymentEnvironment10)
	return a.Environment
}

func (a *AcceptorCompletionAdvice2) AddContext() *CardPaymentContext2 {
	a.Context = new(CardPaymentContext2)
	return a.Context
}

func (a *AcceptorCompletionAdvice2) AddTransaction() *CardPaymentTransaction13 {
	a.Transaction = new(CardPaymentTransaction13)
	return a.Transaction
}

// Notification to the acquirer of the completion of the card payment at the acceptor.
type AcceptorCompletionAdvice3 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment22 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext6 `xml:"Cntxt,omitempty"`

	// Card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction25 `xml:"Tx"`
}

func (a *AcceptorCompletionAdvice3) AddEnvironment() *CardPaymentEnvironment22 {
	a.Environment = new(CardPaymentEnvironment22)
	return a.Environment
}

func (a *AcceptorCompletionAdvice3) AddContext() *CardPaymentContext6 {
	a.Context = new(CardPaymentContext6)
	return a.Context
}

func (a *AcceptorCompletionAdvice3) AddTransaction() *CardPaymentTransaction25 {
	a.Transaction = new(CardPaymentTransaction25)
	return a.Transaction
}

// Notification to the acquirer of the completion of the card payment at the acceptor.
type AcceptorCompletionAdvice4 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment34 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext9 `xml:"Cntxt,omitempty"`

	// Card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction40 `xml:"Tx"`

	// Additional information incorporated as an extension to the message.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AcceptorCompletionAdvice4) AddEnvironment() *CardPaymentEnvironment34 {
	a.Environment = new(CardPaymentEnvironment34)
	return a.Environment
}

func (a *AcceptorCompletionAdvice4) AddContext() *CardPaymentContext9 {
	a.Context = new(CardPaymentContext9)
	return a.Context
}

func (a *AcceptorCompletionAdvice4) AddTransaction() *CardPaymentTransaction40 {
	a.Transaction = new(CardPaymentTransaction40)
	return a.Transaction
}

func (a *AcceptorCompletionAdvice4) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}

// Notification to the acquirer of the completion of the card payment at the acceptor.
type AcceptorCompletionAdvice5 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment47 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext15 `xml:"Cntxt,omitempty"`

	// Card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction55 `xml:"Tx"`

	// Additional information incorporated as an extension to the message.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AcceptorCompletionAdvice5) AddEnvironment() *CardPaymentEnvironment47 {
	a.Environment = new(CardPaymentEnvironment47)
	return a.Environment
}

func (a *AcceptorCompletionAdvice5) AddContext() *CardPaymentContext15 {
	a.Context = new(CardPaymentContext15)
	return a.Context
}

func (a *AcceptorCompletionAdvice5) AddTransaction() *CardPaymentTransaction55 {
	a.Transaction = new(CardPaymentTransaction55)
	return a.Transaction
}

func (a *AcceptorCompletionAdvice5) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}
