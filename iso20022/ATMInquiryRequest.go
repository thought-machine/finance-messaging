package iso20022

// Information related to the request of an inquiry from an ATM.
type ATMInquiryRequest1 struct {

	// Environment in which the inquiry is performed.
	Environment *ATMEnvironment4 `xml:"Envt"`

	// Context in which the inquiry is performed.
	Context *ATMContext5 `xml:"Cntxt"`

	// Inquiry information for the transaction.
	Transaction *ATMTransaction6 `xml:"Tx"`
}

func (a *ATMInquiryRequest1) AddEnvironment() *ATMEnvironment4 {
	a.Environment = new(ATMEnvironment4)
	return a.Environment
}

func (a *ATMInquiryRequest1) AddContext() *ATMContext5 {
	a.Context = new(ATMContext5)
	return a.Context
}

func (a *ATMInquiryRequest1) AddTransaction() *ATMTransaction6 {
	a.Transaction = new(ATMTransaction6)
	return a.Transaction
}

// Information related to the request of an inquiry from an ATM.
type ATMInquiryRequest2 struct {

	// Environment in which the inquiry is performed.
	Environment *ATMEnvironment14 `xml:"Envt"`

	// Context in which the inquiry is performed.
	Context *ATMContext14 `xml:"Cntxt"`

	// Inquiry information for the transaction.
	Transaction *ATMTransaction29 `xml:"Tx"`
}

func (a *ATMInquiryRequest2) AddEnvironment() *ATMEnvironment14 {
	a.Environment = new(ATMEnvironment14)
	return a.Environment
}

func (a *ATMInquiryRequest2) AddContext() *ATMContext14 {
	a.Context = new(ATMContext14)
	return a.Context
}

func (a *ATMInquiryRequest2) AddTransaction() *ATMTransaction29 {
	a.Transaction = new(ATMTransaction29)
	return a.Transaction
}
