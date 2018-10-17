package iso20022

// Detailed invoice data.
type CardPaymentInvoice1 struct {

	// General data relevant to the main body of the invoice such as date of issue, currency code and identification number.
	InvoiceHeader *InvoiceHeader1 `xml:"InvcHdr"`

	// Contractual details related to the agreement between parties.
	TradeAgreement *TradeAgreement6 `xml:"TradAgrmt"`

	// Supply chain shipping arrangements for delivery of invoiced products and/or services.
	TradeDelivery *TradeDelivery1 `xml:"TradDlvry"`

	// Unit of information showing the related provision of products and/or services and monetary summations reported as a discrete line items.
	LineItem []*LineItem10 `xml:"LineItm,omitempty"`
}

func (c *CardPaymentInvoice1) AddInvoiceHeader() *InvoiceHeader1 {
	c.InvoiceHeader = new(InvoiceHeader1)
	return c.InvoiceHeader
}

func (c *CardPaymentInvoice1) AddTradeAgreement() *TradeAgreement6 {
	c.TradeAgreement = new(TradeAgreement6)
	return c.TradeAgreement
}

func (c *CardPaymentInvoice1) AddTradeDelivery() *TradeDelivery1 {
	c.TradeDelivery = new(TradeDelivery1)
	return c.TradeDelivery
}

func (c *CardPaymentInvoice1) AddLineItem() *LineItem10 {
	newValue := new(LineItem10)
	c.LineItem = append(c.LineItem, newValue)
	return newValue
}

// Detailed invoice data.
type CardPaymentInvoice2 struct {

	// General data relevant to the main body of the invoice such as date of issue, currency code and identification number.
	InvoiceHeader *InvoiceHeader2 `xml:"InvcHdr"`

	// Contractual details related to the agreement between parties.
	TradeAgreement *TradeAgreement13 `xml:"TradAgrmt"`

	// Supply chain shipping arrangements for delivery of invoiced products and/or services.
	TradeDelivery *TradeDelivery2 `xml:"TradDlvry"`

	// Unit of information showing the related provision of products and/or services and monetary summations reported as a discrete line items.
	LineItem []*LineItem16 `xml:"LineItm,omitempty"`
}

func (c *CardPaymentInvoice2) AddInvoiceHeader() *InvoiceHeader2 {
	c.InvoiceHeader = new(InvoiceHeader2)
	return c.InvoiceHeader
}

func (c *CardPaymentInvoice2) AddTradeAgreement() *TradeAgreement13 {
	c.TradeAgreement = new(TradeAgreement13)
	return c.TradeAgreement
}

func (c *CardPaymentInvoice2) AddTradeDelivery() *TradeDelivery2 {
	c.TradeDelivery = new(TradeDelivery2)
	return c.TradeDelivery
}

func (c *CardPaymentInvoice2) AddLineItem() *LineItem16 {
	newValue := new(LineItem16)
	c.LineItem = append(c.LineItem, newValue)
	return newValue
}
