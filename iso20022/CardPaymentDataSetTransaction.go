package iso20022

// Transaction to capture in the batch.
type CardPaymentDataSetTransaction1 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Identification of partners involved in the exchange from the merchant to the issuer, with the corresponding timestamp of their exchanges.
	Traceability []*Traceability1 `xml:"Tracblt,omitempty"`

	// Data related to the environment of the transaction in a transaction captured in batch.
	Environment *CardPaymentEnvironment6 `xml:"Envt"`

	// Data related to the context of the transaction.
	Context *CardPaymentContext3 `xml:"Cntxt,omitempty"`

	// Transaction information to be captured.
	Transaction *CardPaymentTransaction4 `xml:"Tx"`
}

func (c *CardPaymentDataSetTransaction1) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSetTransaction1) AddTraceability() *Traceability1 {
	newValue := new(Traceability1)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSetTransaction1) AddEnvironment() *CardPaymentEnvironment6 {
	c.Environment = new(CardPaymentEnvironment6)
	return c.Environment
}

func (c *CardPaymentDataSetTransaction1) AddContext() *CardPaymentContext3 {
	c.Context = new(CardPaymentContext3)
	return c.Context
}

func (c *CardPaymentDataSetTransaction1) AddTransaction() *CardPaymentTransaction4 {
	c.Transaction = new(CardPaymentTransaction4)
	return c.Transaction
}

// Completed card payment transaction to be captured in batch.
type CardPaymentDataSetTransaction10 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Identification of partners involved in the exchange from the merchant to the issuer, with the corresponding timestamp of their exchanges.
	Traceability []*Traceability2 `xml:"Tracblt,omitempty"`

	// Data related to the environment of the transaction in a transaction captured in batch.
	Environment *CardPaymentEnvironment40 `xml:"Envt"`

	// Data related to the context of the transaction.
	Context *CardPaymentContext12 `xml:"Cntxt,omitempty"`

	// Transaction information to be captured.
	Transaction *CardPaymentTransaction45 `xml:"Tx"`
}

func (c *CardPaymentDataSetTransaction10) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSetTransaction10) AddTraceability() *Traceability2 {
	newValue := new(Traceability2)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSetTransaction10) AddEnvironment() *CardPaymentEnvironment40 {
	c.Environment = new(CardPaymentEnvironment40)
	return c.Environment
}

func (c *CardPaymentDataSetTransaction10) AddContext() *CardPaymentContext12 {
	c.Context = new(CardPaymentContext12)
	return c.Context
}

func (c *CardPaymentDataSetTransaction10) AddTransaction() *CardPaymentTransaction45 {
	c.Transaction = new(CardPaymentTransaction45)
	return c.Transaction
}

// Cancelled card payment transaction to be captured in a batch.
type CardPaymentDataSetTransaction11 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Identification of partners involved in the exchange from the merchant to the issuer, with the corresponding timestamp of their exchanges.
	Traceability []*Traceability2 `xml:"Tracblt,omitempty"`

	// Data related to the environment of the cancelled transaction captured in batch.
	Environment *CardPaymentEnvironment40 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext13 `xml:"Cntxt,omitempty"`

	// Card payment cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction46 `xml:"Tx"`
}

func (c *CardPaymentDataSetTransaction11) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSetTransaction11) AddTraceability() *Traceability2 {
	newValue := new(Traceability2)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSetTransaction11) AddEnvironment() *CardPaymentEnvironment40 {
	c.Environment = new(CardPaymentEnvironment40)
	return c.Environment
}

func (c *CardPaymentDataSetTransaction11) AddContext() *CardPaymentContext13 {
	c.Context = new(CardPaymentContext13)
	return c.Context
}

func (c *CardPaymentDataSetTransaction11) AddTransaction() *CardPaymentTransaction46 {
	c.Transaction = new(CardPaymentTransaction46)
	return c.Transaction
}

// Card payment transaction to be authorised in a batch.
type CardPaymentDataSetTransaction12 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Identification of partners involved in the exchange from the merchant to the issuer, with the corresponding timestamp of their exchanges.
	Traceability []*Traceability2 `xml:"Tracblt,omitempty"`

	// Data related to the environment of the card payment transaction to authorise.
	Environment *CardPaymentEnvironment32 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext12 `xml:"Cntxt,omitempty"`

	// Card payment transaction to authorise.
	Transaction *CardPaymentTransaction47 `xml:"Tx"`
}

func (c *CardPaymentDataSetTransaction12) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSetTransaction12) AddTraceability() *Traceability2 {
	newValue := new(Traceability2)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSetTransaction12) AddEnvironment() *CardPaymentEnvironment32 {
	c.Environment = new(CardPaymentEnvironment32)
	return c.Environment
}

func (c *CardPaymentDataSetTransaction12) AddContext() *CardPaymentContext12 {
	c.Context = new(CardPaymentContext12)
	return c.Context
}

func (c *CardPaymentDataSetTransaction12) AddTransaction() *CardPaymentTransaction47 {
	c.Transaction = new(CardPaymentTransaction47)
	return c.Transaction
}

// Card payment transaction including an authorisation response.
type CardPaymentDataSetTransaction13 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Identification of partners involved in the exchange from the merchant to the issuer, with the corresponding timestamp of their exchanges.
	Traceability []*Traceability2 `xml:"Tracblt,omitempty"`

	// Data related to the environment of the card payment transaction.
	Environment *CardPaymentEnvironment41 `xml:"Envt"`

	// Card payment transaction authorisation result.
	Transaction *CardPaymentTransaction48 `xml:"Tx"`

	// Response to the authorisation request from the acquirer.
	TransactionResponse *CardPaymentTransaction39 `xml:"TxRspn"`
}

func (c *CardPaymentDataSetTransaction13) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSetTransaction13) AddTraceability() *Traceability2 {
	newValue := new(Traceability2)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSetTransaction13) AddEnvironment() *CardPaymentEnvironment41 {
	c.Environment = new(CardPaymentEnvironment41)
	return c.Environment
}

func (c *CardPaymentDataSetTransaction13) AddTransaction() *CardPaymentTransaction48 {
	c.Transaction = new(CardPaymentTransaction48)
	return c.Transaction
}

func (c *CardPaymentDataSetTransaction13) AddTransactionResponse() *CardPaymentTransaction39 {
	c.TransactionResponse = new(CardPaymentTransaction39)
	return c.TransactionResponse
}

// Completed card payment transaction to be captured in batch.
type CardPaymentDataSetTransaction14 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Identification of partners involved in the exchange from the merchant to the issuer, with the corresponding timestamp of their exchanges.
	Traceability []*Traceability5 `xml:"Tracblt,omitempty"`

	// Data related to the environment of the transaction in a transaction captured in batch.
	Environment *CardPaymentEnvironment52 `xml:"Envt"`

	// Data related to the context of the transaction.
	Context *CardPaymentContext18 `xml:"Cntxt,omitempty"`

	// Transaction information to be captured.
	Transaction *CardPaymentTransaction60 `xml:"Tx"`
}

func (c *CardPaymentDataSetTransaction14) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSetTransaction14) AddTraceability() *Traceability5 {
	newValue := new(Traceability5)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSetTransaction14) AddEnvironment() *CardPaymentEnvironment52 {
	c.Environment = new(CardPaymentEnvironment52)
	return c.Environment
}

func (c *CardPaymentDataSetTransaction14) AddContext() *CardPaymentContext18 {
	c.Context = new(CardPaymentContext18)
	return c.Context
}

func (c *CardPaymentDataSetTransaction14) AddTransaction() *CardPaymentTransaction60 {
	c.Transaction = new(CardPaymentTransaction60)
	return c.Transaction
}

// Cancelled card payment transaction to be captured in a batch.
type CardPaymentDataSetTransaction15 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Identification of partners involved in the exchange from the merchant to the issuer, with the corresponding timestamp of their exchanges.
	Traceability []*Traceability5 `xml:"Tracblt,omitempty"`

	// Data related to the environment of the cancelled transaction captured in batch.
	Environment *CardPaymentEnvironment52 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext19 `xml:"Cntxt,omitempty"`

	// Card payment cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction61 `xml:"Tx"`
}

func (c *CardPaymentDataSetTransaction15) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSetTransaction15) AddTraceability() *Traceability5 {
	newValue := new(Traceability5)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSetTransaction15) AddEnvironment() *CardPaymentEnvironment52 {
	c.Environment = new(CardPaymentEnvironment52)
	return c.Environment
}

func (c *CardPaymentDataSetTransaction15) AddContext() *CardPaymentContext19 {
	c.Context = new(CardPaymentContext19)
	return c.Context
}

func (c *CardPaymentDataSetTransaction15) AddTransaction() *CardPaymentTransaction61 {
	c.Transaction = new(CardPaymentTransaction61)
	return c.Transaction
}

// Card payment transaction to be authorised in a batch.
type CardPaymentDataSetTransaction16 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Identification of partners involved in the exchange from the merchant to the issuer, with the corresponding timestamp of their exchanges.
	Traceability []*Traceability5 `xml:"Tracblt,omitempty"`

	// Data related to the environment of the card payment transaction to authorise.
	Environment *CardPaymentEnvironment53 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext18 `xml:"Cntxt,omitempty"`

	// Card payment transaction to authorise.
	Transaction *CardPaymentTransaction62 `xml:"Tx"`
}

func (c *CardPaymentDataSetTransaction16) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSetTransaction16) AddTraceability() *Traceability5 {
	newValue := new(Traceability5)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSetTransaction16) AddEnvironment() *CardPaymentEnvironment53 {
	c.Environment = new(CardPaymentEnvironment53)
	return c.Environment
}

func (c *CardPaymentDataSetTransaction16) AddContext() *CardPaymentContext18 {
	c.Context = new(CardPaymentContext18)
	return c.Context
}

func (c *CardPaymentDataSetTransaction16) AddTransaction() *CardPaymentTransaction62 {
	c.Transaction = new(CardPaymentTransaction62)
	return c.Transaction
}

// Card payment transaction including an authorisation response.
type CardPaymentDataSetTransaction17 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Identification of partners involved in the exchange from the merchant to the issuer, with the corresponding timestamp of their exchanges.
	Traceability []*Traceability5 `xml:"Tracblt,omitempty"`

	// Data related to the environment of the card payment transaction.
	Environment *CardPaymentEnvironment54 `xml:"Envt"`

	// Card payment transaction authorisation result.
	Transaction *CardPaymentTransaction63 `xml:"Tx"`

	// Response to the authorisation request from the acquirer.
	TransactionResponse *CardPaymentTransaction54 `xml:"TxRspn"`
}

func (c *CardPaymentDataSetTransaction17) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSetTransaction17) AddTraceability() *Traceability5 {
	newValue := new(Traceability5)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSetTransaction17) AddEnvironment() *CardPaymentEnvironment54 {
	c.Environment = new(CardPaymentEnvironment54)
	return c.Environment
}

func (c *CardPaymentDataSetTransaction17) AddTransaction() *CardPaymentTransaction63 {
	c.Transaction = new(CardPaymentTransaction63)
	return c.Transaction
}

func (c *CardPaymentDataSetTransaction17) AddTransactionResponse() *CardPaymentTransaction54 {
	c.TransactionResponse = new(CardPaymentTransaction54)
	return c.TransactionResponse
}

// Completed card payment transaction to be captured in batch.
type CardPaymentDataSetTransaction2 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Identification of partners involved in the exchange from the merchant to the issuer, with the corresponding timestamp of their exchanges.
	Traceability []*Traceability1 `xml:"Tracblt,omitempty"`

	// Data related to the environment of the transaction in a transaction captured in batch.
	Environment *CardPaymentEnvironment14 `xml:"Envt"`

	// Data related to the context of the transaction.
	Context *CardPaymentContext4 `xml:"Cntxt,omitempty"`

	// Transaction information to be captured.
	Transaction *CardPaymentTransaction14 `xml:"Tx"`
}

func (c *CardPaymentDataSetTransaction2) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSetTransaction2) AddTraceability() *Traceability1 {
	newValue := new(Traceability1)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSetTransaction2) AddEnvironment() *CardPaymentEnvironment14 {
	c.Environment = new(CardPaymentEnvironment14)
	return c.Environment
}

func (c *CardPaymentDataSetTransaction2) AddContext() *CardPaymentContext4 {
	c.Context = new(CardPaymentContext4)
	return c.Context
}

func (c *CardPaymentDataSetTransaction2) AddTransaction() *CardPaymentTransaction14 {
	c.Transaction = new(CardPaymentTransaction14)
	return c.Transaction
}

// Cancelled card payment transaction to be captured in a batch.
type CardPaymentDataSetTransaction3 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Identification of partners involved in the exchange from the merchant to the issuer, with the corresponding timestamp of their exchanges.
	Traceability []*Traceability1 `xml:"Tracblt,omitempty"`

	// Data related to the environment of the cancelled transaction captured in batch.
	Environment *CardPaymentEnvironment14 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext4 `xml:"Cntxt,omitempty"`

	// Card payment cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction20 `xml:"Tx"`
}

func (c *CardPaymentDataSetTransaction3) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSetTransaction3) AddTraceability() *Traceability1 {
	newValue := new(Traceability1)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSetTransaction3) AddEnvironment() *CardPaymentEnvironment14 {
	c.Environment = new(CardPaymentEnvironment14)
	return c.Environment
}

func (c *CardPaymentDataSetTransaction3) AddContext() *CardPaymentContext4 {
	c.Context = new(CardPaymentContext4)
	return c.Context
}

func (c *CardPaymentDataSetTransaction3) AddTransaction() *CardPaymentTransaction20 {
	c.Transaction = new(CardPaymentTransaction20)
	return c.Transaction
}

// Card payment transaction to be authorised in a batch.
type CardPaymentDataSetTransaction4 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Identification of partners involved in the exchange from the merchant to the issuer, with the corresponding timestamp of their exchanges.
	Traceability []*Traceability1 `xml:"Tracblt,omitempty"`

	// Data related to the environment of the card payment transaction to authorise.
	Environment *CardPaymentEnvironment14 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext4 `xml:"Cntxt,omitempty"`

	// Card payment transaction to authorise.
	Transaction *CardPaymentTransaction19 `xml:"Tx"`
}

func (c *CardPaymentDataSetTransaction4) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSetTransaction4) AddTraceability() *Traceability1 {
	newValue := new(Traceability1)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSetTransaction4) AddEnvironment() *CardPaymentEnvironment14 {
	c.Environment = new(CardPaymentEnvironment14)
	return c.Environment
}

func (c *CardPaymentDataSetTransaction4) AddContext() *CardPaymentContext4 {
	c.Context = new(CardPaymentContext4)
	return c.Context
}

func (c *CardPaymentDataSetTransaction4) AddTransaction() *CardPaymentTransaction19 {
	c.Transaction = new(CardPaymentTransaction19)
	return c.Transaction
}

// Card payment transaction including an authorisation response.
type CardPaymentDataSetTransaction5 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Identification of partners involved in the exchange from the merchant to the issuer, with the corresponding timestamp of their exchanges.
	Traceability []*Traceability1 `xml:"Tracblt,omitempty"`

	// Data related to the environment of the card payment transaction.
	Environment *CardPaymentEnvironment16 `xml:"Envt"`

	// Card payment transaction authorisation result.
	Transaction *CardPaymentTransaction12 `xml:"Tx"`

	// Response to the authorisation request from the acquirer.
	TransactionResponse *CardPaymentTransaction18 `xml:"TxRspn"`
}

func (c *CardPaymentDataSetTransaction5) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSetTransaction5) AddTraceability() *Traceability1 {
	newValue := new(Traceability1)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSetTransaction5) AddEnvironment() *CardPaymentEnvironment16 {
	c.Environment = new(CardPaymentEnvironment16)
	return c.Environment
}

func (c *CardPaymentDataSetTransaction5) AddTransaction() *CardPaymentTransaction12 {
	c.Transaction = new(CardPaymentTransaction12)
	return c.Transaction
}

func (c *CardPaymentDataSetTransaction5) AddTransactionResponse() *CardPaymentTransaction18 {
	c.TransactionResponse = new(CardPaymentTransaction18)
	return c.TransactionResponse
}

// Completed card payment transaction to be captured in batch.
type CardPaymentDataSetTransaction6 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Identification of partners involved in the exchange from the merchant to the issuer, with the corresponding timestamp of their exchanges.
	Traceability []*Traceability1 `xml:"Tracblt,omitempty"`

	// Data related to the environment of the transaction in a transaction captured in batch.
	Environment *CardPaymentEnvironment27 `xml:"Envt"`

	// Data related to the context of the transaction.
	Context *CardPaymentContext7 `xml:"Cntxt,omitempty"`

	// Transaction information to be captured.
	Transaction *CardPaymentTransaction29 `xml:"Tx"`
}

func (c *CardPaymentDataSetTransaction6) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSetTransaction6) AddTraceability() *Traceability1 {
	newValue := new(Traceability1)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSetTransaction6) AddEnvironment() *CardPaymentEnvironment27 {
	c.Environment = new(CardPaymentEnvironment27)
	return c.Environment
}

func (c *CardPaymentDataSetTransaction6) AddContext() *CardPaymentContext7 {
	c.Context = new(CardPaymentContext7)
	return c.Context
}

func (c *CardPaymentDataSetTransaction6) AddTransaction() *CardPaymentTransaction29 {
	c.Transaction = new(CardPaymentTransaction29)
	return c.Transaction
}

// Cancelled card payment transaction to be captured in a batch.
type CardPaymentDataSetTransaction7 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Identification of partners involved in the exchange from the merchant to the issuer, with the corresponding timestamp of their exchanges.
	Traceability []*Traceability1 `xml:"Tracblt,omitempty"`

	// Data related to the environment of the cancelled transaction captured in batch.
	Environment *CardPaymentEnvironment27 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext4 `xml:"Cntxt,omitempty"`

	// Card payment cancellation transaction between an acceptor and an acquirer.
	Transaction *CardPaymentTransaction30 `xml:"Tx"`
}

func (c *CardPaymentDataSetTransaction7) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSetTransaction7) AddTraceability() *Traceability1 {
	newValue := new(Traceability1)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSetTransaction7) AddEnvironment() *CardPaymentEnvironment27 {
	c.Environment = new(CardPaymentEnvironment27)
	return c.Environment
}

func (c *CardPaymentDataSetTransaction7) AddContext() *CardPaymentContext4 {
	c.Context = new(CardPaymentContext4)
	return c.Context
}

func (c *CardPaymentDataSetTransaction7) AddTransaction() *CardPaymentTransaction30 {
	c.Transaction = new(CardPaymentTransaction30)
	return c.Transaction
}

// Card payment transaction to be authorised in a batch.
type CardPaymentDataSetTransaction8 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Identification of partners involved in the exchange from the merchant to the issuer, with the corresponding timestamp of their exchanges.
	Traceability []*Traceability1 `xml:"Tracblt,omitempty"`

	// Data related to the environment of the card payment transaction to authorise.
	Environment *CardPaymentEnvironment27 `xml:"Envt"`

	// Context in which the transaction is performed (payment and sale).
	Context *CardPaymentContext7 `xml:"Cntxt,omitempty"`

	// Card payment transaction to authorise.
	Transaction *CardPaymentTransaction31 `xml:"Tx"`
}

func (c *CardPaymentDataSetTransaction8) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSetTransaction8) AddTraceability() *Traceability1 {
	newValue := new(Traceability1)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSetTransaction8) AddEnvironment() *CardPaymentEnvironment27 {
	c.Environment = new(CardPaymentEnvironment27)
	return c.Environment
}

func (c *CardPaymentDataSetTransaction8) AddContext() *CardPaymentContext7 {
	c.Context = new(CardPaymentContext7)
	return c.Context
}

func (c *CardPaymentDataSetTransaction8) AddTransaction() *CardPaymentTransaction31 {
	c.Transaction = new(CardPaymentTransaction31)
	return c.Transaction
}

// Card payment transaction including an authorisation response.
type CardPaymentDataSetTransaction9 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Identification of partners involved in the exchange from the merchant to the issuer, with the corresponding timestamp of their exchanges.
	Traceability []*Traceability1 `xml:"Tracblt,omitempty"`

	// Data related to the environment of the card payment transaction.
	Environment *CardPaymentEnvironment28 `xml:"Envt"`

	// Card payment transaction authorisation result.
	Transaction *CardPaymentTransaction32 `xml:"Tx"`

	// Response to the authorisation request from the acquirer.
	TransactionResponse *CardPaymentTransaction24 `xml:"TxRspn"`
}

func (c *CardPaymentDataSetTransaction9) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSetTransaction9) AddTraceability() *Traceability1 {
	newValue := new(Traceability1)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSetTransaction9) AddEnvironment() *CardPaymentEnvironment28 {
	c.Environment = new(CardPaymentEnvironment28)
	return c.Environment
}

func (c *CardPaymentDataSetTransaction9) AddTransaction() *CardPaymentTransaction32 {
	c.Transaction = new(CardPaymentTransaction32)
	return c.Transaction
}

func (c *CardPaymentDataSetTransaction9) AddTransactionResponse() *CardPaymentTransaction24 {
	c.TransactionResponse = new(CardPaymentTransaction24)
	return c.TransactionResponse
}
