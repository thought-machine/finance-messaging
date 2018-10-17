package iso20022

// Status of the transactions sent in a previous batch of card payment transactions.
type CardPaymentBatchTransferResponse1 struct {

	// Totals of transactions of all the data sets.
	TransactionTotals *TransactionTotals2 `xml:"TxTtls"`

	// Information related to the previously sent set of transaction.
	DataSet []*CardPaymentDataSet5 `xml:"DataSet,omitempty"`
}

func (c *CardPaymentBatchTransferResponse1) AddTransactionTotals() *TransactionTotals2 {
	c.TransactionTotals = new(TransactionTotals2)
	return c.TransactionTotals
}

func (c *CardPaymentBatchTransferResponse1) AddDataSet() *CardPaymentDataSet5 {
	newValue := new(CardPaymentDataSet5)
	c.DataSet = append(c.DataSet, newValue)
	return newValue
}

// Status of the transactions sent in a previous batch of card payment transactions.
type CardPaymentBatchTransferResponse2 struct {

	// Totals of transactions of all the data sets.
	TransactionTotals []*TransactionTotals2 `xml:"TxTtls,omitempty"`

	// Information related to the previously sent set of transaction.
	DataSet []*CardPaymentDataSet9 `xml:"DataSet,omitempty"`
}

func (c *CardPaymentBatchTransferResponse2) AddTransactionTotals() *TransactionTotals2 {
	newValue := new(TransactionTotals2)
	c.TransactionTotals = append(c.TransactionTotals, newValue)
	return newValue
}

func (c *CardPaymentBatchTransferResponse2) AddDataSet() *CardPaymentDataSet9 {
	newValue := new(CardPaymentDataSet9)
	c.DataSet = append(c.DataSet, newValue)
	return newValue
}

// Status of the transactions sent in a previous batch of card payment transactions.
type CardPaymentBatchTransferResponse3 struct {

	// Totals of transactions of all the data sets.
	TransactionTotals []*TransactionTotals3 `xml:"TxTtls,omitempty"`

	// Information related to the previously sent set of transaction.
	DataSet []*CardPaymentDataSet12 `xml:"DataSet,omitempty"`
}

func (c *CardPaymentBatchTransferResponse3) AddTransactionTotals() *TransactionTotals3 {
	newValue := new(TransactionTotals3)
	c.TransactionTotals = append(c.TransactionTotals, newValue)
	return newValue
}

func (c *CardPaymentBatchTransferResponse3) AddDataSet() *CardPaymentDataSet12 {
	newValue := new(CardPaymentDataSet12)
	c.DataSet = append(c.DataSet, newValue)
	return newValue
}

// Status of the transactions sent in a previous batch of card payment transactions.
type CardPaymentBatchTransferResponse4 struct {

	// Totals of transactions of all the data sets.
	TransactionTotals []*TransactionTotals7 `xml:"TxTtls,omitempty"`

	// Information related to the previously sent set of transaction.
	DataSet []*CardPaymentDataSet14 `xml:"DataSet,omitempty"`
}

func (c *CardPaymentBatchTransferResponse4) AddTransactionTotals() *TransactionTotals7 {
	newValue := new(TransactionTotals7)
	c.TransactionTotals = append(c.TransactionTotals, newValue)
	return newValue
}

func (c *CardPaymentBatchTransferResponse4) AddDataSet() *CardPaymentDataSet14 {
	newValue := new(CardPaymentDataSet14)
	c.DataSet = append(c.DataSet, newValue)
	return newValue
}
