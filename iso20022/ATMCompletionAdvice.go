package iso20022

// Information related to the completion of an operation on the ATM.
type ATMCompletionAdvice1 struct {

	// Environment of the transaction.
	Environment *ATMEnvironment3 `xml:"Envt"`

	// Context in which the transaction is performed.
	Context *ATMContext3 `xml:"Cntxt"`

	// Transaction for which the completion is sent.
	Transaction *ATMTransaction5 `xml:"Tx"`
}

func (a *ATMCompletionAdvice1) AddEnvironment() *ATMEnvironment3 {
	a.Environment = new(ATMEnvironment3)
	return a.Environment
}

func (a *ATMCompletionAdvice1) AddContext() *ATMContext3 {
	a.Context = new(ATMContext3)
	return a.Context
}

func (a *ATMCompletionAdvice1) AddTransaction() *ATMTransaction5 {
	a.Transaction = new(ATMTransaction5)
	return a.Transaction
}

// Information related to the completion of an operation on the ATM.
type ATMCompletionAdvice2 struct {

	// Environment of the transaction.
	Environment *ATMEnvironment13 `xml:"Envt"`

	// Context in which the transaction is performed.
	Context *ATMContext13 `xml:"Cntxt"`

	// Transaction for which the completion is sent.
	Transaction *ATMTransaction20 `xml:"Tx"`
}

func (a *ATMCompletionAdvice2) AddEnvironment() *ATMEnvironment13 {
	a.Environment = new(ATMEnvironment13)
	return a.Environment
}

func (a *ATMCompletionAdvice2) AddContext() *ATMContext13 {
	a.Context = new(ATMContext13)
	return a.Context
}

func (a *ATMCompletionAdvice2) AddTransaction() *ATMTransaction20 {
	a.Transaction = new(ATMTransaction20)
	return a.Transaction
}
