package iso20022

// Card payment transactions from one or several data set of transactions.
type CardPaymentBatchTransfer1 struct {

	// Totals of transactions of all the data sets.
	TransactionTotals []*TransactionTotals2 `xml:"TxTtls,omitempty"`

	// Card payment transactions from one data set of transactions.
	DataSet []*CardPaymentDataSet4 `xml:"DataSet,omitempty"`
}

func (c *CardPaymentBatchTransfer1) AddTransactionTotals() *TransactionTotals2 {
	newValue := new(TransactionTotals2)
	c.TransactionTotals = append(c.TransactionTotals, newValue)
	return newValue
}

func (c *CardPaymentBatchTransfer1) AddDataSet() *CardPaymentDataSet4 {
	newValue := new(CardPaymentDataSet4)
	c.DataSet = append(c.DataSet, newValue)
	return newValue
}

// Card payment transactions from one or several data set of transactions.
type CardPaymentBatchTransfer2 struct {

	// Totals of transactions of all the data sets.
	TransactionTotals []*TransactionTotals2 `xml:"TxTtls,omitempty"`

	// Card payment transactions from one data set of transactions.
	DataSet []*CardPaymentDataSet7 `xml:"DataSet,omitempty"`
}

func (c *CardPaymentBatchTransfer2) AddTransactionTotals() *TransactionTotals2 {
	newValue := new(TransactionTotals2)
	c.TransactionTotals = append(c.TransactionTotals, newValue)
	return newValue
}

func (c *CardPaymentBatchTransfer2) AddDataSet() *CardPaymentDataSet7 {
	newValue := new(CardPaymentDataSet7)
	c.DataSet = append(c.DataSet, newValue)
	return newValue
}

// Card payment transactions from one or several data set of transactions.
type CardPaymentBatchTransfer3 struct {

	// Totals of transactions of all the data sets.
	TransactionTotals []*TransactionTotals3 `xml:"TxTtls,omitempty"`

	// Card payment transactions from one data set of transactions.
	DataSet []*CardPaymentDataSet10 `xml:"DataSet,omitempty"`
}

func (c *CardPaymentBatchTransfer3) AddTransactionTotals() *TransactionTotals3 {
	newValue := new(TransactionTotals3)
	c.TransactionTotals = append(c.TransactionTotals, newValue)
	return newValue
}

func (c *CardPaymentBatchTransfer3) AddDataSet() *CardPaymentDataSet10 {
	newValue := new(CardPaymentDataSet10)
	c.DataSet = append(c.DataSet, newValue)
	return newValue
}

// Card payment transactions from one or several data set of transactions.
type CardPaymentBatchTransfer4 struct {

	// Totals of transactions of all the data sets.
	TransactionTotals []*TransactionTotals7 `xml:"TxTtls,omitempty"`

	// Card payment transactions from one data set of transactions.
	DataSet []*CardPaymentDataSet13 `xml:"DataSet,omitempty"`
}

func (c *CardPaymentBatchTransfer4) AddTransactionTotals() *TransactionTotals7 {
	newValue := new(TransactionTotals7)
	c.TransactionTotals = append(c.TransactionTotals, newValue)
	return newValue
}

func (c *CardPaymentBatchTransfer4) AddDataSet() *CardPaymentDataSet13 {
	newValue := new(CardPaymentDataSet13)
	c.DataSet = append(c.DataSet, newValue)
	return newValue
}
