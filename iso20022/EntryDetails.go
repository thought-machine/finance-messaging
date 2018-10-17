package iso20022

// Set of elements used to identify the underlying transaction(s) and/or batched entries.
type EntryDetails1 struct {

	// Set of elements used to provide details on batched transactions.
	Batch *BatchInformation2 `xml:"Btch,omitempty"`

	// Set of elements used to provide information on the underlying transaction(s).
	TransactionDetails []*EntryTransaction2 `xml:"TxDtls,omitempty"`
}

func (e *EntryDetails1) AddBatch() *BatchInformation2 {
	e.Batch = new(BatchInformation2)
	return e.Batch
}

func (e *EntryDetails1) AddTransactionDetails() *EntryTransaction2 {
	newValue := new(EntryTransaction2)
	e.TransactionDetails = append(e.TransactionDetails, newValue)
	return newValue
}

// Identifies the underlying transaction(s) and/or batched entries.
type EntryDetails2 struct {

	// Provides details on batched transactions.
	Batch *BatchInformation2 `xml:"Btch,omitempty"`

	// Provides information on the underlying transaction(s).
	TransactionDetails []*EntryTransaction3 `xml:"TxDtls,omitempty"`
}

func (e *EntryDetails2) AddBatch() *BatchInformation2 {
	e.Batch = new(BatchInformation2)
	return e.Batch
}

func (e *EntryDetails2) AddTransactionDetails() *EntryTransaction3 {
	newValue := new(EntryTransaction3)
	e.TransactionDetails = append(e.TransactionDetails, newValue)
	return newValue
}

// Identifies the underlying transaction(s) and/or batched entries.
type EntryDetails3 struct {

	// Provides details on batched transactions.
	Batch *BatchInformation2 `xml:"Btch,omitempty"`

	// Provides information on the underlying transaction(s).
	TransactionDetails []*EntryTransaction4 `xml:"TxDtls,omitempty"`
}

func (e *EntryDetails3) AddBatch() *BatchInformation2 {
	e.Batch = new(BatchInformation2)
	return e.Batch
}

func (e *EntryDetails3) AddTransactionDetails() *EntryTransaction4 {
	newValue := new(EntryTransaction4)
	e.TransactionDetails = append(e.TransactionDetails, newValue)
	return newValue
}

// Identifies the underlying transaction(s) and/or batched entries.
type EntryDetails6 struct {

	// Provides details on batched transactions.
	Batch *BatchInformation2 `xml:"Btch,omitempty"`

	// Provides information on the underlying transaction(s).
	TransactionDetails []*EntryTransaction7 `xml:"TxDtls,omitempty"`
}

func (e *EntryDetails6) AddBatch() *BatchInformation2 {
	e.Batch = new(BatchInformation2)
	return e.Batch
}

func (e *EntryDetails6) AddTransactionDetails() *EntryTransaction7 {
	newValue := new(EntryTransaction7)
	e.TransactionDetails = append(e.TransactionDetails, newValue)
	return newValue
}

// Identifies the underlying transaction(s) and/or batched entries.
type EntryDetails7 struct {

	// Provides details on batched transactions.
	Batch *BatchInformation2 `xml:"Btch,omitempty"`

	// Provides information on the underlying transaction(s).
	TransactionDetails []*EntryTransaction8 `xml:"TxDtls,omitempty"`
}

func (e *EntryDetails7) AddBatch() *BatchInformation2 {
	e.Batch = new(BatchInformation2)
	return e.Batch
}

func (e *EntryDetails7) AddTransactionDetails() *EntryTransaction8 {
	newValue := new(EntryTransaction8)
	e.TransactionDetails = append(e.TransactionDetails, newValue)
	return newValue
}
