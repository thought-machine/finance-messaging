package iso20022

// Information related to the currency conversion request.
type AcceptorCurrencyConversionRequest1 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment30 `xml:"Envt"`

	// Card payment transaction for which the currency conversion is requested.
	Transaction *CardPaymentTransaction34 `xml:"Tx"`
}

func (a *AcceptorCurrencyConversionRequest1) AddEnvironment() *CardPaymentEnvironment30 {
	a.Environment = new(CardPaymentEnvironment30)
	return a.Environment
}

func (a *AcceptorCurrencyConversionRequest1) AddTransaction() *CardPaymentTransaction34 {
	a.Transaction = new(CardPaymentTransaction34)
	return a.Transaction
}

// Information related to the currency conversion request.
type AcceptorCurrencyConversionRequest2 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment44 `xml:"Envt"`

	// Card payment transaction for which the currency conversion is requested.
	Transaction *CardPaymentTransaction49 `xml:"Tx"`
}

func (a *AcceptorCurrencyConversionRequest2) AddEnvironment() *CardPaymentEnvironment44 {
	a.Environment = new(CardPaymentEnvironment44)
	return a.Environment
}

func (a *AcceptorCurrencyConversionRequest2) AddTransaction() *CardPaymentTransaction49 {
	a.Transaction = new(CardPaymentTransaction49)
	return a.Transaction
}

// Information related to the currency conversion request.
type AcceptorCurrencyConversionRequest3 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment56 `xml:"Envt"`

	// Card payment transaction for which the currency conversion is requested.
	Transaction *CardPaymentTransaction64 `xml:"Tx"`
}

func (a *AcceptorCurrencyConversionRequest3) AddEnvironment() *CardPaymentEnvironment56 {
	a.Environment = new(CardPaymentEnvironment56)
	return a.Environment
}

func (a *AcceptorCurrencyConversionRequest3) AddTransaction() *CardPaymentTransaction64 {
	a.Transaction = new(CardPaymentTransaction64)
	return a.Transaction
}
