package iso20022

// Card payment completion advice response from the acquirer.
type CardPaymentTransactionAdviceResponse1 struct {

	// Unique identification of the transaction by the POI.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Result of a requested service.
	Response *Response1Code `xml:"Rspn"`
}

func (c *CardPaymentTransactionAdviceResponse1) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransactionAdviceResponse1) SetResponse(value string) {
	c.Response = (*Response1Code)(&value)
}

// Cancellation advice response from the acquirer.
type CardPaymentTransactionAdviceResponse2 struct {

	// Unique identification of the transaction by the POI.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`
}

func (c *CardPaymentTransactionAdviceResponse2) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

// Card payment completion advice response from the acquirer.
type CardPaymentTransactionAdviceResponse3 struct {

	// Unique identification of the transaction by the POI.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Result of a requested service.
	Response *Response1Code `xml:"Rspn"`
}

func (c *CardPaymentTransactionAdviceResponse3) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransactionAdviceResponse3) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransactionAdviceResponse3) SetResponse(value string) {
	c.Response = (*Response1Code)(&value)
}

// Card payment completion advice response from the acquirer.
type CardPaymentTransactionAdviceResponse4 struct {

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Result of a requested service.
	Response *Response1Code `xml:"Rspn"`
}

func (c *CardPaymentTransactionAdviceResponse4) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransactionAdviceResponse4) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransactionAdviceResponse4) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransactionAdviceResponse4) SetResponse(value string) {
	c.Response = (*Response1Code)(&value)
}

// Card payment completion advice response from the acquirer.
type CardPaymentTransactionAdviceResponse5 struct {

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Result of a requested service.
	Response *Response1Code `xml:"Rspn"`
}

func (c *CardPaymentTransactionAdviceResponse5) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransactionAdviceResponse5) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransactionAdviceResponse5) SetResponse(value string) {
	c.Response = (*Response1Code)(&value)
}

// Card payment completion advice response from the acquirer.
type CardPaymentTransactionAdviceResponse6 struct {

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Result of a requested service.
	Response *Response4Code `xml:"Rspn"`
}

func (c *CardPaymentTransactionAdviceResponse6) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransactionAdviceResponse6) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransactionAdviceResponse6) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransactionAdviceResponse6) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransactionAdviceResponse6) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransactionAdviceResponse6) SetResponse(value string) {
	c.Response = (*Response4Code)(&value)
}
