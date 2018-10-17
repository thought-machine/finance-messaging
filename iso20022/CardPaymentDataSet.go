package iso20022

// Set of transactions to capture, sharing common characteristics.
type CardPaymentDataSet1 struct {

	// Identification of the data set.
	DataSetIdentification *DataSetIdentification1 `xml:"DataSetId"`

	// Identification of partners involved in the data set building.
	Traceability []*Traceability1 `xml:"Tracblt,omitempty"`

	// Initiator of the data set.
	DataSetInitiator *GenericIdentification32 `xml:"DataSetInitr,omitempty"`

	// Transaction totals of the data set.
	TransactionTotals []*TransactionTotals1 `xml:"TxTtls"`

	// Data common to all transactions of the data set.
	CommonData *CommonData1 `xml:"CmonData,omitempty"`

	// Set of transaction to capture.
	TransactionToCapture []*CardPaymentDataSetTransaction1 `xml:"TxToCaptr,omitempty"`
}

func (c *CardPaymentDataSet1) AddDataSetIdentification() *DataSetIdentification1 {
	c.DataSetIdentification = new(DataSetIdentification1)
	return c.DataSetIdentification
}

func (c *CardPaymentDataSet1) AddTraceability() *Traceability1 {
	newValue := new(Traceability1)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSet1) AddDataSetInitiator() *GenericIdentification32 {
	c.DataSetInitiator = new(GenericIdentification32)
	return c.DataSetInitiator
}

func (c *CardPaymentDataSet1) AddTransactionTotals() *TransactionTotals1 {
	newValue := new(TransactionTotals1)
	c.TransactionTotals = append(c.TransactionTotals, newValue)
	return newValue
}

func (c *CardPaymentDataSet1) AddCommonData() *CommonData1 {
	c.CommonData = new(CommonData1)
	return c.CommonData
}

func (c *CardPaymentDataSet1) AddTransactionToCapture() *CardPaymentDataSetTransaction1 {
	newValue := new(CardPaymentDataSetTransaction1)
	c.TransactionToCapture = append(c.TransactionToCapture, newValue)
	return newValue
}

// Set of transactions to capture, sharing common characteristics.
type CardPaymentDataSet10 struct {

	// Identification of the data set.
	DataSetIdentification *DataSetIdentification1 `xml:"DataSetId"`

	// Identification of partners involved in the data set building.
	Traceability []*Traceability2 `xml:"Tracblt,omitempty"`

	// Initiator of the data set.
	DataSetInitiator *GenericIdentification53 `xml:"DataSetInitr,omitempty"`

	// Transaction totals of the data set.
	TransactionTotals []*TransactionTotals3 `xml:"TxTtls"`

	// Data common to all transactions of the data set.
	CommonData *CommonData4 `xml:"CmonData,omitempty"`

	// Set of transaction to Process.
	Transaction []*CardPaymentDataSetTransaction3Choice `xml:"Tx"`
}

func (c *CardPaymentDataSet10) AddDataSetIdentification() *DataSetIdentification1 {
	c.DataSetIdentification = new(DataSetIdentification1)
	return c.DataSetIdentification
}

func (c *CardPaymentDataSet10) AddTraceability() *Traceability2 {
	newValue := new(Traceability2)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSet10) AddDataSetInitiator() *GenericIdentification53 {
	c.DataSetInitiator = new(GenericIdentification53)
	return c.DataSetInitiator
}

func (c *CardPaymentDataSet10) AddTransactionTotals() *TransactionTotals3 {
	newValue := new(TransactionTotals3)
	c.TransactionTotals = append(c.TransactionTotals, newValue)
	return newValue
}

func (c *CardPaymentDataSet10) AddCommonData() *CommonData4 {
	c.CommonData = new(CommonData4)
	return c.CommonData
}

func (c *CardPaymentDataSet10) AddTransaction() *CardPaymentDataSetTransaction3Choice {
	newValue := new(CardPaymentDataSetTransaction3Choice)
	c.Transaction = append(c.Transaction, newValue)
	return newValue
}

// Transaction for whose batch capture has been rejected.
type CardPaymentDataSet11 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Response to the capture of the transaction.
	TransactionResponse *ResponseType1 `xml:"TxRspn"`

	// Data related to the environment of the transaction.
	Environment *CardPaymentEnvironment33 `xml:"Envt"`

	// Transaction that has been rejected.
	Transaction *CardPaymentTransactionAdviceResponse5 `xml:"Tx"`
}

func (c *CardPaymentDataSet11) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSet11) AddTransactionResponse() *ResponseType1 {
	c.TransactionResponse = new(ResponseType1)
	return c.TransactionResponse
}

func (c *CardPaymentDataSet11) AddEnvironment() *CardPaymentEnvironment33 {
	c.Environment = new(CardPaymentEnvironment33)
	return c.Environment
}

func (c *CardPaymentDataSet11) AddTransaction() *CardPaymentTransactionAdviceResponse5 {
	c.Transaction = new(CardPaymentTransactionAdviceResponse5)
	return c.Transaction
}

// Result of the captured set of transactions.
type CardPaymentDataSet12 struct {

	// Identification of the data set.
	DataSetIdentification *DataSetIdentification1 `xml:"DataSetId"`

	// Result of the data set capture.
	DataSetResult *ResponseType1 `xml:"DataSetRslt"`

	// Indicates if the data set must be removed from the POI (Point Of Interaction).
	RemoveDataSet *TrueFalseIndicator `xml:"RmvDataSet"`

	// Initiator of the data set.
	DataSetInitiator *GenericIdentification53 `xml:"DataSetInitr,omitempty"`

	// Transaction totals of the batch.
	TransactionTotals []*TransactionTotals3 `xml:"TxTtls"`

	// Transaction in the batch, whose capture has been rejected.
	RejectedTransaction []*CardPaymentDataSet11 `xml:"RjctdTx,omitempty"`
}

func (c *CardPaymentDataSet12) AddDataSetIdentification() *DataSetIdentification1 {
	c.DataSetIdentification = new(DataSetIdentification1)
	return c.DataSetIdentification
}

func (c *CardPaymentDataSet12) AddDataSetResult() *ResponseType1 {
	c.DataSetResult = new(ResponseType1)
	return c.DataSetResult
}

func (c *CardPaymentDataSet12) SetRemoveDataSet(value string) {
	c.RemoveDataSet = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentDataSet12) AddDataSetInitiator() *GenericIdentification53 {
	c.DataSetInitiator = new(GenericIdentification53)
	return c.DataSetInitiator
}

func (c *CardPaymentDataSet12) AddTransactionTotals() *TransactionTotals3 {
	newValue := new(TransactionTotals3)
	c.TransactionTotals = append(c.TransactionTotals, newValue)
	return newValue
}

func (c *CardPaymentDataSet12) AddRejectedTransaction() *CardPaymentDataSet11 {
	newValue := new(CardPaymentDataSet11)
	c.RejectedTransaction = append(c.RejectedTransaction, newValue)
	return newValue
}

// Set of transactions to capture, sharing common characteristics.
type CardPaymentDataSet13 struct {

	// Identification of the data set.
	DataSetIdentification *DataSetIdentification5 `xml:"DataSetId"`

	// Identification of partners involved in the data set building.
	Traceability []*Traceability5 `xml:"Tracblt,omitempty"`

	// Initiator of the data set.
	DataSetInitiator *GenericIdentification53 `xml:"DataSetInitr,omitempty"`

	// Transaction totals of the data set.
	TransactionTotals []*TransactionTotals7 `xml:"TxTtls"`

	// Data common to all transactions of the data set.
	CommonData *CommonData5 `xml:"CmonData,omitempty"`

	// Set of transaction to Process.
	Transaction []*CardPaymentDataSetTransaction4Choice `xml:"Tx"`
}

func (c *CardPaymentDataSet13) AddDataSetIdentification() *DataSetIdentification5 {
	c.DataSetIdentification = new(DataSetIdentification5)
	return c.DataSetIdentification
}

func (c *CardPaymentDataSet13) AddTraceability() *Traceability5 {
	newValue := new(Traceability5)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSet13) AddDataSetInitiator() *GenericIdentification53 {
	c.DataSetInitiator = new(GenericIdentification53)
	return c.DataSetInitiator
}

func (c *CardPaymentDataSet13) AddTransactionTotals() *TransactionTotals7 {
	newValue := new(TransactionTotals7)
	c.TransactionTotals = append(c.TransactionTotals, newValue)
	return newValue
}

func (c *CardPaymentDataSet13) AddCommonData() *CommonData5 {
	c.CommonData = new(CommonData5)
	return c.CommonData
}

func (c *CardPaymentDataSet13) AddTransaction() *CardPaymentDataSetTransaction4Choice {
	newValue := new(CardPaymentDataSetTransaction4Choice)
	c.Transaction = append(c.Transaction, newValue)
	return newValue
}

// Result of the captured set of transactions.
type CardPaymentDataSet14 struct {

	// Identification of the data set.
	DataSetIdentification *DataSetIdentification5 `xml:"DataSetId"`

	// Result of the data set capture.
	DataSetResult *ResponseType5 `xml:"DataSetRslt"`

	// Indicates if the data set must be removed from the POI (Point Of Interaction).
	RemoveDataSet *TrueFalseIndicator `xml:"RmvDataSet"`

	// Initiator of the data set.
	DataSetInitiator *GenericIdentification53 `xml:"DataSetInitr,omitempty"`

	// Transaction totals of the batch.
	TransactionTotals []*TransactionTotals7 `xml:"TxTtls"`

	// Transaction in the batch, whose capture has been rejected.
	RejectedTransaction []*CardPaymentDataSet15 `xml:"RjctdTx,omitempty"`
}

func (c *CardPaymentDataSet14) AddDataSetIdentification() *DataSetIdentification5 {
	c.DataSetIdentification = new(DataSetIdentification5)
	return c.DataSetIdentification
}

func (c *CardPaymentDataSet14) AddDataSetResult() *ResponseType5 {
	c.DataSetResult = new(ResponseType5)
	return c.DataSetResult
}

func (c *CardPaymentDataSet14) SetRemoveDataSet(value string) {
	c.RemoveDataSet = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentDataSet14) AddDataSetInitiator() *GenericIdentification53 {
	c.DataSetInitiator = new(GenericIdentification53)
	return c.DataSetInitiator
}

func (c *CardPaymentDataSet14) AddTransactionTotals() *TransactionTotals7 {
	newValue := new(TransactionTotals7)
	c.TransactionTotals = append(c.TransactionTotals, newValue)
	return newValue
}

func (c *CardPaymentDataSet14) AddRejectedTransaction() *CardPaymentDataSet15 {
	newValue := new(CardPaymentDataSet15)
	c.RejectedTransaction = append(c.RejectedTransaction, newValue)
	return newValue
}

// Transaction for whose batch capture has been rejected.
type CardPaymentDataSet15 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Response to the capture of the transaction.
	TransactionResponse *ResponseType1 `xml:"TxRspn"`

	// Data related to the environment of the transaction.
	Environment *CardPaymentEnvironment46 `xml:"Envt"`

	// Transaction that has been rejected.
	Transaction *CardPaymentTransactionAdviceResponse5 `xml:"Tx"`
}

func (c *CardPaymentDataSet15) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSet15) AddTransactionResponse() *ResponseType1 {
	c.TransactionResponse = new(ResponseType1)
	return c.TransactionResponse
}

func (c *CardPaymentDataSet15) AddEnvironment() *CardPaymentEnvironment46 {
	c.Environment = new(CardPaymentEnvironment46)
	return c.Environment
}

func (c *CardPaymentDataSet15) AddTransaction() *CardPaymentTransactionAdviceResponse5 {
	c.Transaction = new(CardPaymentTransactionAdviceResponse5)
	return c.Transaction
}

// Result of the captured set of transactions.
type CardPaymentDataSet2 struct {

	// Identification of the data set.
	DataSetIdentification *DataSetIdentification1 `xml:"DataSetId"`

	// Result of the data set capture.
	DataSetResult *ResponseType1 `xml:"DataSetRslt"`

	// Indicates if the data set must be removed from the POI (Point Of Interaction).
	RemoveDataSet *TrueFalseIndicator `xml:"RmvDataSet"`

	// Initiator of the data set.
	DataSetInitiator *GenericIdentification32 `xml:"DataSetInitr,omitempty"`

	// Transaction totals of the batch.
	TransactionTotals []*TransactionTotals1 `xml:"TxTtls"`

	// Transaction in the batch, whose capture has been rejected.
	RejectedTransaction []*CardPaymentDataSet3 `xml:"RjctdTx,omitempty"`
}

func (c *CardPaymentDataSet2) AddDataSetIdentification() *DataSetIdentification1 {
	c.DataSetIdentification = new(DataSetIdentification1)
	return c.DataSetIdentification
}

func (c *CardPaymentDataSet2) AddDataSetResult() *ResponseType1 {
	c.DataSetResult = new(ResponseType1)
	return c.DataSetResult
}

func (c *CardPaymentDataSet2) SetRemoveDataSet(value string) {
	c.RemoveDataSet = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentDataSet2) AddDataSetInitiator() *GenericIdentification32 {
	c.DataSetInitiator = new(GenericIdentification32)
	return c.DataSetInitiator
}

func (c *CardPaymentDataSet2) AddTransactionTotals() *TransactionTotals1 {
	newValue := new(TransactionTotals1)
	c.TransactionTotals = append(c.TransactionTotals, newValue)
	return newValue
}

func (c *CardPaymentDataSet2) AddRejectedTransaction() *CardPaymentDataSet3 {
	newValue := new(CardPaymentDataSet3)
	c.RejectedTransaction = append(c.RejectedTransaction, newValue)
	return newValue
}

// Transaction for whose batch capture has been rejected.
type CardPaymentDataSet3 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Response to the capture of the transaction.
	TransactionResponse *ResponseType1 `xml:"TxRspn"`

	// Data related to the environment of the transaction.
	Environment *CardPaymentEnvironment3 `xml:"Envt"`

	// Transaction that has been rejected.
	Transaction *CardPaymentTransactionAdviceResponse1 `xml:"Tx"`
}

func (c *CardPaymentDataSet3) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSet3) AddTransactionResponse() *ResponseType1 {
	c.TransactionResponse = new(ResponseType1)
	return c.TransactionResponse
}

func (c *CardPaymentDataSet3) AddEnvironment() *CardPaymentEnvironment3 {
	c.Environment = new(CardPaymentEnvironment3)
	return c.Environment
}

func (c *CardPaymentDataSet3) AddTransaction() *CardPaymentTransactionAdviceResponse1 {
	c.Transaction = new(CardPaymentTransactionAdviceResponse1)
	return c.Transaction
}

// Set of transactions to capture, sharing common characteristics.
type CardPaymentDataSet4 struct {

	// Identification of the data set.
	DataSetIdentification *DataSetIdentification1 `xml:"DataSetId"`

	// Identification of partners involved in the data set building.
	Traceability []*Traceability1 `xml:"Tracblt,omitempty"`

	// Initiator of the data set.
	DataSetInitiator *GenericIdentification32 `xml:"DataSetInitr,omitempty"`

	// Transaction totals of the data set.
	TransactionTotals []*TransactionTotals2 `xml:"TxTtls"`

	// Data common to all transactions of the data set.
	CommonData *CommonData2 `xml:"CmonData,omitempty"`

	// Set of transaction to Process.
	Transaction []*CardPaymentDataSetTransaction1Choice `xml:"Tx"`
}

func (c *CardPaymentDataSet4) AddDataSetIdentification() *DataSetIdentification1 {
	c.DataSetIdentification = new(DataSetIdentification1)
	return c.DataSetIdentification
}

func (c *CardPaymentDataSet4) AddTraceability() *Traceability1 {
	newValue := new(Traceability1)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSet4) AddDataSetInitiator() *GenericIdentification32 {
	c.DataSetInitiator = new(GenericIdentification32)
	return c.DataSetInitiator
}

func (c *CardPaymentDataSet4) AddTransactionTotals() *TransactionTotals2 {
	newValue := new(TransactionTotals2)
	c.TransactionTotals = append(c.TransactionTotals, newValue)
	return newValue
}

func (c *CardPaymentDataSet4) AddCommonData() *CommonData2 {
	c.CommonData = new(CommonData2)
	return c.CommonData
}

func (c *CardPaymentDataSet4) AddTransaction() *CardPaymentDataSetTransaction1Choice {
	newValue := new(CardPaymentDataSetTransaction1Choice)
	c.Transaction = append(c.Transaction, newValue)
	return newValue
}

// Result of the captured set of transactions.
type CardPaymentDataSet5 struct {

	// Identification of the data set.
	DataSetIdentification *DataSetIdentification1 `xml:"DataSetId"`

	// Result of the data set capture.
	DataSetResult *ResponseType1 `xml:"DataSetRslt"`

	// Indicates if the data set must be removed from the POI (Point Of Interaction).
	RemoveDataSet *TrueFalseIndicator `xml:"RmvDataSet"`

	// Initiator of the data set.
	DataSetInitiator *GenericIdentification32 `xml:"DataSetInitr,omitempty"`

	// Transaction totals of the batch.
	TransactionTotals []*TransactionTotals2 `xml:"TxTtls"`

	// Transaction in the batch, whose capture has been rejected.
	RejectedTransaction []*CardPaymentDataSet6 `xml:"RjctdTx,omitempty"`
}

func (c *CardPaymentDataSet5) AddDataSetIdentification() *DataSetIdentification1 {
	c.DataSetIdentification = new(DataSetIdentification1)
	return c.DataSetIdentification
}

func (c *CardPaymentDataSet5) AddDataSetResult() *ResponseType1 {
	c.DataSetResult = new(ResponseType1)
	return c.DataSetResult
}

func (c *CardPaymentDataSet5) SetRemoveDataSet(value string) {
	c.RemoveDataSet = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentDataSet5) AddDataSetInitiator() *GenericIdentification32 {
	c.DataSetInitiator = new(GenericIdentification32)
	return c.DataSetInitiator
}

func (c *CardPaymentDataSet5) AddTransactionTotals() *TransactionTotals2 {
	newValue := new(TransactionTotals2)
	c.TransactionTotals = append(c.TransactionTotals, newValue)
	return newValue
}

func (c *CardPaymentDataSet5) AddRejectedTransaction() *CardPaymentDataSet6 {
	newValue := new(CardPaymentDataSet6)
	c.RejectedTransaction = append(c.RejectedTransaction, newValue)
	return newValue
}

// Transaction for whose batch capture has been rejected.
type CardPaymentDataSet6 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Response to the capture of the transaction.
	TransactionResponse *ResponseType1 `xml:"TxRspn"`

	// Data related to the environment of the transaction.
	Environment *CardPaymentEnvironment11 `xml:"Envt"`

	// Transaction that has been rejected.
	Transaction *CardPaymentTransactionAdviceResponse1 `xml:"Tx"`
}

func (c *CardPaymentDataSet6) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSet6) AddTransactionResponse() *ResponseType1 {
	c.TransactionResponse = new(ResponseType1)
	return c.TransactionResponse
}

func (c *CardPaymentDataSet6) AddEnvironment() *CardPaymentEnvironment11 {
	c.Environment = new(CardPaymentEnvironment11)
	return c.Environment
}

func (c *CardPaymentDataSet6) AddTransaction() *CardPaymentTransactionAdviceResponse1 {
	c.Transaction = new(CardPaymentTransactionAdviceResponse1)
	return c.Transaction
}

// Set of transactions to capture, sharing common characteristics.
type CardPaymentDataSet7 struct {

	// Identification of the data set.
	DataSetIdentification *DataSetIdentification1 `xml:"DataSetId"`

	// Identification of partners involved in the data set building.
	Traceability []*Traceability1 `xml:"Tracblt,omitempty"`

	// Initiator of the data set.
	DataSetInitiator *GenericIdentification32 `xml:"DataSetInitr,omitempty"`

	// Transaction totals of the data set.
	TransactionTotals []*TransactionTotals2 `xml:"TxTtls"`

	// Data common to all transactions of the data set.
	CommonData *CommonData3 `xml:"CmonData,omitempty"`

	// Set of transaction to Process.
	Transaction []*CardPaymentDataSetTransaction2Choice `xml:"Tx"`
}

func (c *CardPaymentDataSet7) AddDataSetIdentification() *DataSetIdentification1 {
	c.DataSetIdentification = new(DataSetIdentification1)
	return c.DataSetIdentification
}

func (c *CardPaymentDataSet7) AddTraceability() *Traceability1 {
	newValue := new(Traceability1)
	c.Traceability = append(c.Traceability, newValue)
	return newValue
}

func (c *CardPaymentDataSet7) AddDataSetInitiator() *GenericIdentification32 {
	c.DataSetInitiator = new(GenericIdentification32)
	return c.DataSetInitiator
}

func (c *CardPaymentDataSet7) AddTransactionTotals() *TransactionTotals2 {
	newValue := new(TransactionTotals2)
	c.TransactionTotals = append(c.TransactionTotals, newValue)
	return newValue
}

func (c *CardPaymentDataSet7) AddCommonData() *CommonData3 {
	c.CommonData = new(CommonData3)
	return c.CommonData
}

func (c *CardPaymentDataSet7) AddTransaction() *CardPaymentDataSetTransaction2Choice {
	newValue := new(CardPaymentDataSetTransaction2Choice)
	c.Transaction = append(c.Transaction, newValue)
	return newValue
}

// Transaction for whose batch capture has been rejected.
type CardPaymentDataSet8 struct {

	// Sequential counter of the transaction.
	TransactionSequenceCounter *Max9NumericText `xml:"TxSeqCntr"`

	// Response to the capture of the transaction.
	TransactionResponse *ResponseType1 `xml:"TxRspn"`

	// Data related to the environment of the transaction.
	Environment *CardPaymentEnvironment21 `xml:"Envt"`

	// Transaction that has been rejected.
	Transaction *CardPaymentTransactionAdviceResponse5 `xml:"Tx"`
}

func (c *CardPaymentDataSet8) SetTransactionSequenceCounter(value string) {
	c.TransactionSequenceCounter = (*Max9NumericText)(&value)
}

func (c *CardPaymentDataSet8) AddTransactionResponse() *ResponseType1 {
	c.TransactionResponse = new(ResponseType1)
	return c.TransactionResponse
}

func (c *CardPaymentDataSet8) AddEnvironment() *CardPaymentEnvironment21 {
	c.Environment = new(CardPaymentEnvironment21)
	return c.Environment
}

func (c *CardPaymentDataSet8) AddTransaction() *CardPaymentTransactionAdviceResponse5 {
	c.Transaction = new(CardPaymentTransactionAdviceResponse5)
	return c.Transaction
}

// Result of the captured set of transactions.
type CardPaymentDataSet9 struct {

	// Identification of the data set.
	DataSetIdentification *DataSetIdentification1 `xml:"DataSetId"`

	// Result of the data set capture.
	DataSetResult *ResponseType1 `xml:"DataSetRslt"`

	// Indicates if the data set must be removed from the POI (Point Of Interaction).
	RemoveDataSet *TrueFalseIndicator `xml:"RmvDataSet"`

	// Initiator of the data set.
	DataSetInitiator *GenericIdentification32 `xml:"DataSetInitr,omitempty"`

	// Transaction totals of the batch.
	TransactionTotals []*TransactionTotals2 `xml:"TxTtls"`

	// Transaction in the batch, whose capture has been rejected.
	RejectedTransaction []*CardPaymentDataSet8 `xml:"RjctdTx,omitempty"`
}

func (c *CardPaymentDataSet9) AddDataSetIdentification() *DataSetIdentification1 {
	c.DataSetIdentification = new(DataSetIdentification1)
	return c.DataSetIdentification
}

func (c *CardPaymentDataSet9) AddDataSetResult() *ResponseType1 {
	c.DataSetResult = new(ResponseType1)
	return c.DataSetResult
}

func (c *CardPaymentDataSet9) SetRemoveDataSet(value string) {
	c.RemoveDataSet = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentDataSet9) AddDataSetInitiator() *GenericIdentification32 {
	c.DataSetInitiator = new(GenericIdentification32)
	return c.DataSetInitiator
}

func (c *CardPaymentDataSet9) AddTransactionTotals() *TransactionTotals2 {
	newValue := new(TransactionTotals2)
	c.TransactionTotals = append(c.TransactionTotals, newValue)
	return newValue
}

func (c *CardPaymentDataSet9) AddRejectedTransaction() *CardPaymentDataSet8 {
	newValue := new(CardPaymentDataSet8)
	c.RejectedTransaction = append(c.RejectedTransaction, newValue)
	return newValue
}
