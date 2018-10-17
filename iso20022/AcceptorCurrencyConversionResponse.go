package iso20022

// Currencey conversion outcome from the service provider.
type AcceptorCurrencyConversionResponse1 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment21 `xml:"Envt"`

	// Currency conversion of a card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction23 `xml:"Tx"`

	// Result of the currency conversion.
	TransactionResponse *Response1Code `xml:"TxRspn"`

	// Details of the currency conversion.
	CurrencyConversion *CurrencyConversion1 `xml:"CcyConvs"`
}

func (a *AcceptorCurrencyConversionResponse1) AddEnvironment() *CardPaymentEnvironment21 {
	a.Environment = new(CardPaymentEnvironment21)
	return a.Environment
}

func (a *AcceptorCurrencyConversionResponse1) AddTransaction() *CardPaymentTransaction23 {
	a.Transaction = new(CardPaymentTransaction23)
	return a.Transaction
}

func (a *AcceptorCurrencyConversionResponse1) SetTransactionResponse(value string) {
	a.TransactionResponse = (*Response1Code)(&value)
}

func (a *AcceptorCurrencyConversionResponse1) AddCurrencyConversion() *CurrencyConversion1 {
	a.CurrencyConversion = new(CurrencyConversion1)
	return a.CurrencyConversion
}

// Currency conversion outcome from the service provider.
type AcceptorCurrencyConversionResponse2 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment33 `xml:"Envt"`

	// Currency conversion of a card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction38 `xml:"Tx"`

	// Details of the currency conversion.
	CurrencyConversion *CurrencyConversion3 `xml:"CcyConvs"`
}

func (a *AcceptorCurrencyConversionResponse2) AddEnvironment() *CardPaymentEnvironment33 {
	a.Environment = new(CardPaymentEnvironment33)
	return a.Environment
}

func (a *AcceptorCurrencyConversionResponse2) AddTransaction() *CardPaymentTransaction38 {
	a.Transaction = new(CardPaymentTransaction38)
	return a.Transaction
}

func (a *AcceptorCurrencyConversionResponse2) AddCurrencyConversion() *CurrencyConversion3 {
	a.CurrencyConversion = new(CurrencyConversion3)
	return a.CurrencyConversion
}

// Currency conversion outcome from the service provider.
type AcceptorCurrencyConversionResponse3 struct {

	// Environment of the transaction.
	Environment *CardPaymentEnvironment46 `xml:"Envt"`

	// Currency conversion of a card payment transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction53 `xml:"Tx"`

	// Details of the currency conversion.
	CurrencyConversionResult *CurrencyConversion7 `xml:"CcyConvsRslt"`
}

func (a *AcceptorCurrencyConversionResponse3) AddEnvironment() *CardPaymentEnvironment46 {
	a.Environment = new(CardPaymentEnvironment46)
	return a.Environment
}

func (a *AcceptorCurrencyConversionResponse3) AddTransaction() *CardPaymentTransaction53 {
	a.Transaction = new(CardPaymentTransaction53)
	return a.Transaction
}

func (a *AcceptorCurrencyConversionResponse3) AddCurrencyConversionResult() *CurrencyConversion7 {
	a.CurrencyConversionResult = new(CurrencyConversion7)
	return a.CurrencyConversionResult
}
