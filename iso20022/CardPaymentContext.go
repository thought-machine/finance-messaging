package iso20022

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext1 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext1 `xml:"PmtCntxt"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext1 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext1) AddPaymentContext() *PaymentContext1 {
	c.PaymentContext = new(PaymentContext1)
	return c.PaymentContext
}

func (c *CardPaymentContext1) AddSaleContext() *SaleContext1 {
	c.SaleContext = new(SaleContext1)
	return c.SaleContext
}

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext10 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext10 `xml:"PmtCntxt"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext1 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext10) AddPaymentContext() *PaymentContext10 {
	c.PaymentContext = new(PaymentContext10)
	return c.PaymentContext
}

func (c *CardPaymentContext10) AddSaleContext() *SaleContext1 {
	c.SaleContext = new(SaleContext1)
	return c.SaleContext
}

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext11 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext11 `xml:"PmtCntxt"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext1 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext11) AddPaymentContext() *PaymentContext11 {
	c.PaymentContext = new(PaymentContext11)
	return c.PaymentContext
}

func (c *CardPaymentContext11) AddSaleContext() *SaleContext1 {
	c.SaleContext = new(SaleContext1)
	return c.SaleContext
}

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext12 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext12 `xml:"PmtCntxt,omitempty"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext1 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext12) AddPaymentContext() *PaymentContext12 {
	c.PaymentContext = new(PaymentContext12)
	return c.PaymentContext
}

func (c *CardPaymentContext12) AddSaleContext() *SaleContext1 {
	c.SaleContext = new(SaleContext1)
	return c.SaleContext
}

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext13 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext13 `xml:"PmtCntxt,omitempty"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext1 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext13) AddPaymentContext() *PaymentContext13 {
	c.PaymentContext = new(PaymentContext13)
	return c.PaymentContext
}

func (c *CardPaymentContext13) AddSaleContext() *SaleContext1 {
	c.SaleContext = new(SaleContext1)
	return c.SaleContext
}

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext14 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext14 `xml:"PmtCntxt"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext2 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext14) AddPaymentContext() *PaymentContext14 {
	c.PaymentContext = new(PaymentContext14)
	return c.PaymentContext
}

func (c *CardPaymentContext14) AddSaleContext() *SaleContext2 {
	c.SaleContext = new(SaleContext2)
	return c.SaleContext
}

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext15 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext15 `xml:"PmtCntxt"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext2 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext15) AddPaymentContext() *PaymentContext15 {
	c.PaymentContext = new(PaymentContext15)
	return c.PaymentContext
}

func (c *CardPaymentContext15) AddSaleContext() *SaleContext2 {
	c.SaleContext = new(SaleContext2)
	return c.SaleContext
}

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext16 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext16 `xml:"PmtCntxt"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext2 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext16) AddPaymentContext() *PaymentContext16 {
	c.PaymentContext = new(PaymentContext16)
	return c.PaymentContext
}

func (c *CardPaymentContext16) AddSaleContext() *SaleContext2 {
	c.SaleContext = new(SaleContext2)
	return c.SaleContext
}

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext17 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext17 `xml:"PmtCntxt"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext2 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext17) AddPaymentContext() *PaymentContext17 {
	c.PaymentContext = new(PaymentContext17)
	return c.PaymentContext
}

func (c *CardPaymentContext17) AddSaleContext() *SaleContext2 {
	c.SaleContext = new(SaleContext2)
	return c.SaleContext
}

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext18 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext18 `xml:"PmtCntxt,omitempty"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext2 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext18) AddPaymentContext() *PaymentContext18 {
	c.PaymentContext = new(PaymentContext18)
	return c.PaymentContext
}

func (c *CardPaymentContext18) AddSaleContext() *SaleContext2 {
	c.SaleContext = new(SaleContext2)
	return c.SaleContext
}

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext19 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext19 `xml:"PmtCntxt,omitempty"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext2 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext19) AddPaymentContext() *PaymentContext19 {
	c.PaymentContext = new(PaymentContext19)
	return c.PaymentContext
}

func (c *CardPaymentContext19) AddSaleContext() *SaleContext2 {
	c.SaleContext = new(SaleContext2)
	return c.SaleContext
}

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext2 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext2 `xml:"PmtCntxt"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext1 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext2) AddPaymentContext() *PaymentContext2 {
	c.PaymentContext = new(PaymentContext2)
	return c.PaymentContext
}

func (c *CardPaymentContext2) AddSaleContext() *SaleContext1 {
	c.SaleContext = new(SaleContext1)
	return c.SaleContext
}

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext3 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext1 `xml:"PmtCntxt,omitempty"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext1 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext3) AddPaymentContext() *PaymentContext1 {
	c.PaymentContext = new(PaymentContext1)
	return c.PaymentContext
}

func (c *CardPaymentContext3) AddSaleContext() *SaleContext1 {
	c.SaleContext = new(SaleContext1)
	return c.SaleContext
}

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext4 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext4 `xml:"PmtCntxt,omitempty"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext1 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext4) AddPaymentContext() *PaymentContext4 {
	c.PaymentContext = new(PaymentContext4)
	return c.PaymentContext
}

func (c *CardPaymentContext4) AddSaleContext() *SaleContext1 {
	c.SaleContext = new(SaleContext1)
	return c.SaleContext
}

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext5 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext5 `xml:"PmtCntxt"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext1 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext5) AddPaymentContext() *PaymentContext5 {
	c.PaymentContext = new(PaymentContext5)
	return c.PaymentContext
}

func (c *CardPaymentContext5) AddSaleContext() *SaleContext1 {
	c.SaleContext = new(SaleContext1)
	return c.SaleContext
}

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext6 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext6 `xml:"PmtCntxt"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext1 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext6) AddPaymentContext() *PaymentContext6 {
	c.PaymentContext = new(PaymentContext6)
	return c.PaymentContext
}

func (c *CardPaymentContext6) AddSaleContext() *SaleContext1 {
	c.SaleContext = new(SaleContext1)
	return c.SaleContext
}

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext7 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext7 `xml:"PmtCntxt,omitempty"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext1 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext7) AddPaymentContext() *PaymentContext7 {
	c.PaymentContext = new(PaymentContext7)
	return c.PaymentContext
}

func (c *CardPaymentContext7) AddSaleContext() *SaleContext1 {
	c.SaleContext = new(SaleContext1)
	return c.SaleContext
}

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext8 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext8 `xml:"PmtCntxt"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext1 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext8) AddPaymentContext() *PaymentContext8 {
	c.PaymentContext = new(PaymentContext8)
	return c.PaymentContext
}

func (c *CardPaymentContext8) AddSaleContext() *SaleContext1 {
	c.SaleContext = new(SaleContext1)
	return c.SaleContext
}

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext9 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext9 `xml:"PmtCntxt"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext1 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext9) AddPaymentContext() *PaymentContext9 {
	c.PaymentContext = new(PaymentContext9)
	return c.PaymentContext
}

func (c *CardPaymentContext9) AddSaleContext() *SaleContext1 {
	c.SaleContext = new(SaleContext1)
	return c.SaleContext
}
