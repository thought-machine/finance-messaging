package iso20022

// Information related to the request of a PIN management from an ATM.
type ATMPINManagementRequest1 struct {

	// Environment in which the transaction is performed.
	Environment *ATMEnvironment1 `xml:"Envt"`

	// Context in which the transaction is performed.
	Context *ATMContext7 `xml:"Cntxt"`

	// Transaction for which the service is requested.
	Transaction *ATMTransaction9 `xml:"Tx"`
}

func (a *ATMPINManagementRequest1) AddEnvironment() *ATMEnvironment1 {
	a.Environment = new(ATMEnvironment1)
	return a.Environment
}

func (a *ATMPINManagementRequest1) AddContext() *ATMContext7 {
	a.Context = new(ATMContext7)
	return a.Context
}

func (a *ATMPINManagementRequest1) AddTransaction() *ATMTransaction9 {
	a.Transaction = new(ATMTransaction9)
	return a.Transaction
}

// Information related to the request of a PIN management from an ATM.
type ATMPINManagementRequest2 struct {

	// Environment in which the transaction is performed.
	Environment *ATMEnvironment11 `xml:"Envt"`

	// Context in which the transaction is performed.
	Context *ATMContext16 `xml:"Cntxt"`

	// Transaction for which the service is requested.
	Transaction *ATMTransaction9 `xml:"Tx"`
}

func (a *ATMPINManagementRequest2) AddEnvironment() *ATMEnvironment11 {
	a.Environment = new(ATMEnvironment11)
	return a.Environment
}

func (a *ATMPINManagementRequest2) AddContext() *ATMContext16 {
	a.Context = new(ATMContext16)
	return a.Context
}

func (a *ATMPINManagementRequest2) AddTransaction() *ATMTransaction9 {
	a.Transaction = new(ATMTransaction9)
	return a.Transaction
}
