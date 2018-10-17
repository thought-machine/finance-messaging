package iso20022

// Environment of the transaction.
type CardTransactionEnvironment1 struct {

	// Acquirer of the card transaction.
	Acquirer *Acquirer6 `xml:"Acqrr"`

	// Identification of the interconnected card scheme from which the request is coming.
	CardSchemeIdentification *Max35Text `xml:"CardSchmeId,omitempty"`

	// Acceptor performing the card transaction.
	Acceptor *Organisation18 `xml:"Accptr,omitempty"`

	// Payment terminal or ATM performing the transaction.
	Terminal *CardAcceptorTerminal1 `xml:"Termnl,omitempty"`

	// Card performing the transaction.
	Card *PaymentCard12 `xml:"Card"`

	// Container of tenders used by the customer to perform the payment.
	CustomerDevice *CustomerDevice1 `xml:"CstmrDvc,omitempty"`

	// Container of tenders used by the customer to perform the payment.
	Wallet *CustomerDevice1 `xml:"Wllt,omitempty"`

	// Payment token information.
	PaymentToken *CardPaymentToken4 `xml:"PmtTkn,omitempty"`

	// Cardholder involved in the card transaction.
	// It correspond partially to the ISO 8583:2003, field number 49-71.
	Cardholder *Cardholder9 `xml:"Crdhldr,omitempty"`

	// Protection of cardholder sensitive data by a digital envelope using a cryptographic key.
	ProtectedCardholderData *ContentInformationType10 `xml:"PrtctdCrdhldrData,omitempty"`
}

func (c *CardTransactionEnvironment1) AddAcquirer() *Acquirer6 {
	c.Acquirer = new(Acquirer6)
	return c.Acquirer
}

func (c *CardTransactionEnvironment1) SetCardSchemeIdentification(value string) {
	c.CardSchemeIdentification = (*Max35Text)(&value)
}

func (c *CardTransactionEnvironment1) AddAcceptor() *Organisation18 {
	c.Acceptor = new(Organisation18)
	return c.Acceptor
}

func (c *CardTransactionEnvironment1) AddTerminal() *CardAcceptorTerminal1 {
	c.Terminal = new(CardAcceptorTerminal1)
	return c.Terminal
}

func (c *CardTransactionEnvironment1) AddCard() *PaymentCard12 {
	c.Card = new(PaymentCard12)
	return c.Card
}

func (c *CardTransactionEnvironment1) AddCustomerDevice() *CustomerDevice1 {
	c.CustomerDevice = new(CustomerDevice1)
	return c.CustomerDevice
}

func (c *CardTransactionEnvironment1) AddWallet() *CustomerDevice1 {
	c.Wallet = new(CustomerDevice1)
	return c.Wallet
}

func (c *CardTransactionEnvironment1) AddPaymentToken() *CardPaymentToken4 {
	c.PaymentToken = new(CardPaymentToken4)
	return c.PaymentToken
}

func (c *CardTransactionEnvironment1) AddCardholder() *Cardholder9 {
	c.Cardholder = new(Cardholder9)
	return c.Cardholder
}

func (c *CardTransactionEnvironment1) AddProtectedCardholderData() *ContentInformationType10 {
	c.ProtectedCardholderData = new(ContentInformationType10)
	return c.ProtectedCardholderData
}

// Environment of the transaction.
type CardTransactionEnvironment2 struct {

	// Acquirer identification of the transaction.
	AcquirerIdentification *Max35Text `xml:"AcqrrId"`

	// Identification of the interconnected card scheme from which the response is coming.
	CardSchemeIdentification *Max35Text `xml:"CardSchmeId,omitempty"`

	// Identification of the card acceptor performing the transaction.
	AcceptorIdentification *Max35Text `xml:"AccptrId,omitempty"`

	// Identification of the card terminal performing the transaction.
	TerminalIdentification *Max35Text `xml:"TermnlId,omitempty"`

	// Card performing the transaction.
	Card *PaymentCard13 `xml:"Card"`

	// Payment token information.
	PaymentToken *CardPaymentToken2 `xml:"PmtTkn,omitempty"`

	// Postal address for delivery of goods or services.
	ShippingAddress *PostalAddress18 `xml:"ShppgAdr,omitempty"`
}

func (c *CardTransactionEnvironment2) SetAcquirerIdentification(value string) {
	c.AcquirerIdentification = (*Max35Text)(&value)
}

func (c *CardTransactionEnvironment2) SetCardSchemeIdentification(value string) {
	c.CardSchemeIdentification = (*Max35Text)(&value)
}

func (c *CardTransactionEnvironment2) SetAcceptorIdentification(value string) {
	c.AcceptorIdentification = (*Max35Text)(&value)
}

func (c *CardTransactionEnvironment2) SetTerminalIdentification(value string) {
	c.TerminalIdentification = (*Max35Text)(&value)
}

func (c *CardTransactionEnvironment2) AddCard() *PaymentCard13 {
	c.Card = new(PaymentCard13)
	return c.Card
}

func (c *CardTransactionEnvironment2) AddPaymentToken() *CardPaymentToken2 {
	c.PaymentToken = new(CardPaymentToken2)
	return c.PaymentToken
}

func (c *CardTransactionEnvironment2) AddShippingAddress() *PostalAddress18 {
	c.ShippingAddress = new(PostalAddress18)
	return c.ShippingAddress
}

// Environment of the transaction.
type CardTransactionEnvironment3 struct {

	// Acquirer of the card transaction.
	Acquirer *Acquirer6 `xml:"Acqrr"`

	// Identification of the interconnected card scheme from which the response is coming.
	CardSchemeIdentification *Max35Text `xml:"CardSchmeId,omitempty"`

	// Acceptor performing the card transaction.
	Acceptor *Organisation19 `xml:"Accptr"`

	// Identification of the card terminal performing the transaction.
	TerminalIdentification *GenericIdentification32 `xml:"TermnlId,omitempty"`

	// Card performing the transaction.
	Card *PaymentCard14 `xml:"Card"`

	// Container of tenders used by the customer to perform the payment.
	CustomerDevice *CustomerDevice1 `xml:"CstmrDvc,omitempty"`

	// Container of tenders used by the customer to perform the payment.
	Wallet *CustomerDevice1 `xml:"Wllt,omitempty"`

	// Payment token information.
	PaymentToken *CardPaymentToken4 `xml:"PmtTkn,omitempty"`
}

func (c *CardTransactionEnvironment3) AddAcquirer() *Acquirer6 {
	c.Acquirer = new(Acquirer6)
	return c.Acquirer
}

func (c *CardTransactionEnvironment3) SetCardSchemeIdentification(value string) {
	c.CardSchemeIdentification = (*Max35Text)(&value)
}

func (c *CardTransactionEnvironment3) AddAcceptor() *Organisation19 {
	c.Acceptor = new(Organisation19)
	return c.Acceptor
}

func (c *CardTransactionEnvironment3) AddTerminalIdentification() *GenericIdentification32 {
	c.TerminalIdentification = new(GenericIdentification32)
	return c.TerminalIdentification
}

func (c *CardTransactionEnvironment3) AddCard() *PaymentCard14 {
	c.Card = new(PaymentCard14)
	return c.Card
}

func (c *CardTransactionEnvironment3) AddCustomerDevice() *CustomerDevice1 {
	c.CustomerDevice = new(CustomerDevice1)
	return c.CustomerDevice
}

func (c *CardTransactionEnvironment3) AddWallet() *CustomerDevice1 {
	c.Wallet = new(CustomerDevice1)
	return c.Wallet
}

func (c *CardTransactionEnvironment3) AddPaymentToken() *CardPaymentToken4 {
	c.PaymentToken = new(CardPaymentToken4)
	return c.PaymentToken
}

// Environment of the transaction.
type CardTransactionEnvironment4 struct {

	// Acquirer identification of the card transaction.
	AcquirerIdentification *Max35Text `xml:"AcqrrId"`

	// Identification of the interconnected card scheme from which the request is coming.
	CardSchemeIdentification *Max35Text `xml:"CardSchmeId,omitempty"`

	// Identification of the card acceptor performing the transaction.
	AcceptorIdentification *Max35Text `xml:"AccptrId,omitempty"`

	// Identification of the card terminal performing the transaction.
	TerminalIdentification *Max35Text `xml:"TermnlId,omitempty"`

	// Card performing the transaction.
	Card *PaymentCard15 `xml:"Card"`

	// Payment token information.
	PaymentToken *CardPaymentToken2 `xml:"PmtTkn,omitempty"`
}

func (c *CardTransactionEnvironment4) SetAcquirerIdentification(value string) {
	c.AcquirerIdentification = (*Max35Text)(&value)
}

func (c *CardTransactionEnvironment4) SetCardSchemeIdentification(value string) {
	c.CardSchemeIdentification = (*Max35Text)(&value)
}

func (c *CardTransactionEnvironment4) SetAcceptorIdentification(value string) {
	c.AcceptorIdentification = (*Max35Text)(&value)
}

func (c *CardTransactionEnvironment4) SetTerminalIdentification(value string) {
	c.TerminalIdentification = (*Max35Text)(&value)
}

func (c *CardTransactionEnvironment4) AddCard() *PaymentCard15 {
	c.Card = new(PaymentCard15)
	return c.Card
}

func (c *CardTransactionEnvironment4) AddPaymentToken() *CardPaymentToken2 {
	c.PaymentToken = new(CardPaymentToken2)
	return c.PaymentToken
}

// Environment of the transaction.
type CardTransactionEnvironment5 struct {

	// Institution initiator of the reconciliation.
	// It corresponds to the ISO 8583 field number 94.
	SendingInstitution *GenericIdentification73 `xml:"SndgInstn"`

	// Institution destination of the reconciliation.
	// It corresponds to the ISO 8583 field number 93.
	ReceivingInstitution *GenericIdentification73 `xml:"RcvgInstn"`

	// Institution in charge of the settlement of the transaction.
	SettlementInstitution *GenericIdentification73 `xml:"SttlmInstn"`
}

func (c *CardTransactionEnvironment5) AddSendingInstitution() *GenericIdentification73 {
	c.SendingInstitution = new(GenericIdentification73)
	return c.SendingInstitution
}

func (c *CardTransactionEnvironment5) AddReceivingInstitution() *GenericIdentification73 {
	c.ReceivingInstitution = new(GenericIdentification73)
	return c.ReceivingInstitution
}

func (c *CardTransactionEnvironment5) AddSettlementInstitution() *GenericIdentification73 {
	c.SettlementInstitution = new(GenericIdentification73)
	return c.SettlementInstitution
}

// Environment of the transaction.
type CardTransactionEnvironment6 struct {

	// Institution initiator of the reconciliation (correspond to the ISO 8583 field 94).
	SendingInstitution *GenericIdentification73 `xml:"SndgInstn"`

	// Institution destination of the reconciliation (correspond to the ISO 8583 field 93).
	ReceivingInstitution *GenericIdentification73 `xml:"RcvgInstn"`
}

func (c *CardTransactionEnvironment6) AddSendingInstitution() *GenericIdentification73 {
	c.SendingInstitution = new(GenericIdentification73)
	return c.SendingInstitution
}

func (c *CardTransactionEnvironment6) AddReceivingInstitution() *GenericIdentification73 {
	c.ReceivingInstitution = new(GenericIdentification73)
	return c.ReceivingInstitution
}
