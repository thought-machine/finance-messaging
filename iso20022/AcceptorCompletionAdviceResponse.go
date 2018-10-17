package iso20022

// Acknowledgement by the acquirer, of the completion advice of the card payment at the acceptor.
type AcceptorCompletionAdviceResponse1 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment3 `xml:"Envt"`

	// Card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransactionAdviceResponse1 `xml:"Tx"`

	// Instructions for contacting the terminal management host.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`
}

func (a *AcceptorCompletionAdviceResponse1) AddEnvironment() *CardPaymentEnvironment3 {
	a.Environment = new(CardPaymentEnvironment3)
	return a.Environment
}

func (a *AcceptorCompletionAdviceResponse1) AddTransaction() *CardPaymentTransactionAdviceResponse1 {
	a.Transaction = new(CardPaymentTransactionAdviceResponse1)
	return a.Transaction
}

func (a *AcceptorCompletionAdviceResponse1) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}

// Acknowledgement by the acquirer, of the completion advice of the card payment at the acceptor.
type AcceptorCompletionAdviceResponse2 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment11 `xml:"Envt"`

	// Card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransactionAdviceResponse3 `xml:"Tx"`

	// Instructions for contacting the terminal management host.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`
}

func (a *AcceptorCompletionAdviceResponse2) AddEnvironment() *CardPaymentEnvironment11 {
	a.Environment = new(CardPaymentEnvironment11)
	return a.Environment
}

func (a *AcceptorCompletionAdviceResponse2) AddTransaction() *CardPaymentTransactionAdviceResponse3 {
	a.Transaction = new(CardPaymentTransactionAdviceResponse3)
	return a.Transaction
}

func (a *AcceptorCompletionAdviceResponse2) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}

// Acknowledgement by the acquirer, of the completion advice of the card payment at the acceptor.
type AcceptorCompletionAdviceResponse3 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment21 `xml:"Envt"`

	// Card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransactionAdviceResponse4 `xml:"Tx"`

	// Instructions for contacting the terminal management host.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`
}

func (a *AcceptorCompletionAdviceResponse3) AddEnvironment() *CardPaymentEnvironment21 {
	a.Environment = new(CardPaymentEnvironment21)
	return a.Environment
}

func (a *AcceptorCompletionAdviceResponse3) AddTransaction() *CardPaymentTransactionAdviceResponse4 {
	a.Transaction = new(CardPaymentTransactionAdviceResponse4)
	return a.Transaction
}

func (a *AcceptorCompletionAdviceResponse3) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}

// Acknowledgement by the acquirer, of the completion advice of the card payment at the acceptor.
type AcceptorCompletionAdviceResponse4 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment33 `xml:"Envt"`

	// Card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransactionAdviceResponse4 `xml:"Tx"`

	// Instructions for contacting the terminal management host.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`

	// Additional information incorporated as an extension to the message.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AcceptorCompletionAdviceResponse4) AddEnvironment() *CardPaymentEnvironment33 {
	a.Environment = new(CardPaymentEnvironment33)
	return a.Environment
}

func (a *AcceptorCompletionAdviceResponse4) AddTransaction() *CardPaymentTransactionAdviceResponse4 {
	a.Transaction = new(CardPaymentTransactionAdviceResponse4)
	return a.Transaction
}

func (a *AcceptorCompletionAdviceResponse4) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}

func (a *AcceptorCompletionAdviceResponse4) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}

// Acknowledgement by the acquirer, of the completion advice of the card payment at the acceptor.
type AcceptorCompletionAdviceResponse5 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment46 `xml:"Envt"`

	// Card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransactionAdviceResponse6 `xml:"Tx"`

	// Instructions for contacting the terminal management host.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`

	// Additional information incorporated as an extension to the message.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AcceptorCompletionAdviceResponse5) AddEnvironment() *CardPaymentEnvironment46 {
	a.Environment = new(CardPaymentEnvironment46)
	return a.Environment
}

func (a *AcceptorCompletionAdviceResponse5) AddTransaction() *CardPaymentTransactionAdviceResponse6 {
	a.Transaction = new(CardPaymentTransactionAdviceResponse6)
	return a.Transaction
}

func (a *AcceptorCompletionAdviceResponse5) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}

func (a *AcceptorCompletionAdviceResponse5) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}
