package iso20022

// Provides details information on the original transactions, to which the status report message refers.
type OriginalPaymentInstruction1 struct {

	// Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId"`

	// Number of individual transactions contained in the original payment information group.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original payment information group, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Specifies the status of the payment information group.
	PaymentInformationStatus *TransactionGroupStatus3Code `xml:"PmtInfSts,omitempty"`

	// Provides detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Detailed information on the number of transactions for each identical transaction status.
	NumberOfTransactionsPerStatus []*NumberOfTransactionsPerStatus3 `xml:"NbOfTxsPerSts,omitempty"`

	// Provides information on the original transactions to which the status report message refers.
	TransactionInformationAndStatus []*PaymentTransaction32 `xml:"TxInfAndSts,omitempty"`
}

func (o *OriginalPaymentInstruction1) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction1) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInstruction1) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInstruction1) SetPaymentInformationStatus(value string) {
	o.PaymentInformationStatus = (*TransactionGroupStatus3Code)(&value)
}

func (o *OriginalPaymentInstruction1) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	o.StatusReasonInformation = append(o.StatusReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction1) AddNumberOfTransactionsPerStatus() *NumberOfTransactionsPerStatus3 {
	newValue := new(NumberOfTransactionsPerStatus3)
	o.NumberOfTransactionsPerStatus = append(o.NumberOfTransactionsPerStatus, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction1) AddTransactionInformationAndStatus() *PaymentTransaction32 {
	newValue := new(PaymentTransaction32)
	o.TransactionInformationAndStatus = append(o.TransactionInformationAndStatus, newValue)
	return newValue
}

// Provides the details on the reference and status of the original transactions, included in the original instruction, to which the cancellation request message applies.
type OriginalPaymentInstruction10 struct {

	// Unique identification, as assigned by the original assigner, to unambiguously identify the original payment information cancellation request.
	OriginalPaymentInformationCancellationIdentification *Max35Text `xml:"OrgnlPmtInfCxlId,omitempty"`

	// Identifies the resolved case.
	ResolvedCase *Case3 `xml:"RslvdCase,omitempty"`

	// Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId"`

	// Provides information on the original message.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Number of individual transactions contained in the original payment information group.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original payment information group, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Specifies the status of a cancellation request, related to a payment information group.
	PaymentInformationCancellationStatus *GroupCancellationStatus1Code `xml:"PmtInfCxlSts,omitempty"`

	// Provides detailed information on the cancellation status reason.
	CancellationStatusReasonInformation []*CancellationStatusReason2 `xml:"CxlStsRsnInf,omitempty"`

	// Detailed information on the number of transactions for each identical cancellation status.
	NumberOfTransactionsPerCancellationStatus []*NumberOfCancellationsPerStatus1 `xml:"NbOfTxsPerCxlSts,omitempty"`

	// Provides information on the original transactions to which the cancellation request message refers.
	TransactionInformationAndStatus []*PaymentTransaction54 `xml:"TxInfAndSts,omitempty"`
}

func (o *OriginalPaymentInstruction10) SetOriginalPaymentInformationCancellationIdentification(value string) {
	o.OriginalPaymentInformationCancellationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction10) AddResolvedCase() *Case3 {
	o.ResolvedCase = new(Case3)
	return o.ResolvedCase
}

func (o *OriginalPaymentInstruction10) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction10) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	o.OriginalGroupInformation = new(OriginalGroupInformation3)
	return o.OriginalGroupInformation
}

func (o *OriginalPaymentInstruction10) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInstruction10) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInstruction10) SetPaymentInformationCancellationStatus(value string) {
	o.PaymentInformationCancellationStatus = (*GroupCancellationStatus1Code)(&value)
}

func (o *OriginalPaymentInstruction10) AddCancellationStatusReasonInformation() *CancellationStatusReason2 {
	newValue := new(CancellationStatusReason2)
	o.CancellationStatusReasonInformation = append(o.CancellationStatusReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction10) AddNumberOfTransactionsPerCancellationStatus() *NumberOfCancellationsPerStatus1 {
	newValue := new(NumberOfCancellationsPerStatus1)
	o.NumberOfTransactionsPerCancellationStatus = append(o.NumberOfTransactionsPerCancellationStatus, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction10) AddTransactionInformationAndStatus() *PaymentTransaction54 {
	newValue := new(PaymentTransaction54)
	o.TransactionInformationAndStatus = append(o.TransactionInformationAndStatus, newValue)
	return newValue
}

// Provides details on the original transactions, to which the status report message refers.
type OriginalPaymentInstruction11 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reversed payment information group.
	// Usage: The instructing party is the party sending the reversal message and not the party that sent the original instruction that is being reversed.
	ReversalPaymentInformationIdentification *Max35Text `xml:"RvslPmtInfId,omitempty"`

	// Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId"`

	// Number of individual transactions contained in the original payment information group.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original payment information group, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Indicates whether or not the reversal applies to the complete original payment information group or to individual transactions within that group.
	PaymentInformationReversal *TrueFalseIndicator `xml:"PmtInfRvsl,omitempty"`

	// Provides detailed information on the reversal reason.
	ReversalReasonInformation []*PaymentReversalReason7 `xml:"RvslRsnInf,omitempty"`

	// Provides information on the original transactions to which the reversal message refers.
	TransactionInformation []*PaymentTransaction56 `xml:"TxInf,omitempty"`
}

func (o *OriginalPaymentInstruction11) SetReversalPaymentInformationIdentification(value string) {
	o.ReversalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction11) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction11) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInstruction11) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInstruction11) SetBatchBooking(value string) {
	o.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (o *OriginalPaymentInstruction11) SetPaymentInformationReversal(value string) {
	o.PaymentInformationReversal = (*TrueFalseIndicator)(&value)
}

func (o *OriginalPaymentInstruction11) AddReversalReasonInformation() *PaymentReversalReason7 {
	newValue := new(PaymentReversalReason7)
	o.ReversalReasonInformation = append(o.ReversalReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction11) AddTransactionInformation() *PaymentTransaction56 {
	newValue := new(PaymentTransaction56)
	o.TransactionInformation = append(o.TransactionInformation, newValue)
	return newValue
}

// Provides details information on the original transactions, to which the status report message refers.
type OriginalPaymentInstruction12 struct {

	// Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId"`

	// Number of individual transactions contained in the original payment information group.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original payment information group, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Specifies the status of the payment information group.
	PaymentInformationStatus *TransactionGroupStatus3Code `xml:"PmtInfSts,omitempty"`

	// Provides detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Detailed information on the number of transactions for each identical transaction status.
	NumberOfTransactionsPerStatus []*NumberOfTransactionsPerStatus3 `xml:"NbOfTxsPerSts,omitempty"`

	// Provides information on the original transactions to which the status report message refers.
	TransactionInformationAndStatus []*PaymentTransaction57 `xml:"TxInfAndSts,omitempty"`
}

func (o *OriginalPaymentInstruction12) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction12) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInstruction12) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInstruction12) SetPaymentInformationStatus(value string) {
	o.PaymentInformationStatus = (*TransactionGroupStatus3Code)(&value)
}

func (o *OriginalPaymentInstruction12) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	o.StatusReasonInformation = append(o.StatusReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction12) AddNumberOfTransactionsPerStatus() *NumberOfTransactionsPerStatus3 {
	newValue := new(NumberOfTransactionsPerStatus3)
	o.NumberOfTransactionsPerStatus = append(o.NumberOfTransactionsPerStatus, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction12) AddTransactionInformationAndStatus() *PaymentTransaction57 {
	newValue := new(PaymentTransaction57)
	o.TransactionInformationAndStatus = append(o.TransactionInformationAndStatus, newValue)
	return newValue
}

// Provides details on the reference and status of the original transactions, included in the original instruction, to which the cancellation request message applies.
type OriginalPaymentInstruction13 struct {

	// Unique identification, as assigned by the assigner, to unambiguously identify the cancellation request.
	//
	// Usage: The cancellation request identification can be used for reconciliation or to link tasks relating to the cancellation request.
	PaymentCancellationIdentification *Max35Text `xml:"PmtCxlId,omitempty"`

	// Identifies the case.
	Case *Case3 `xml:"Case,omitempty"`

	// Unique and unambiguous identifier of the original payment information block, as assigned by the original sending party.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId"`

	// Information concerning the original group of transactions, to which the message refers.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Number of individual transactions contained in the cancellation payment information group.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs,omitempty"`

	// Total of all individual amounts included in the cancellation payment information group, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Indicates whether or not the cancellation applies to a whole group of transactions or to individual transactions within the original group.
	PaymentInformationCancellation *GroupCancellationIndicator `xml:"PmtInfCxl,omitempty"`

	// Detailed information on the cancellation reason.
	CancellationReasonInformation []*PaymentCancellationReason2 `xml:"CxlRsnInf,omitempty"`

	// Information concerning the original transactions, to which the cancellation request message refers.
	TransactionInformation []*PaymentTransaction58 `xml:"TxInf,omitempty"`
}

func (o *OriginalPaymentInstruction13) SetPaymentCancellationIdentification(value string) {
	o.PaymentCancellationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction13) AddCase() *Case3 {
	o.Case = new(Case3)
	return o.Case
}

func (o *OriginalPaymentInstruction13) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction13) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	o.OriginalGroupInformation = new(OriginalGroupInformation3)
	return o.OriginalGroupInformation
}

func (o *OriginalPaymentInstruction13) SetNumberOfTransactions(value string) {
	o.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInstruction13) SetControlSum(value string) {
	o.ControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInstruction13) SetPaymentInformationCancellation(value string) {
	o.PaymentInformationCancellation = (*GroupCancellationIndicator)(&value)
}

func (o *OriginalPaymentInstruction13) AddCancellationReasonInformation() *PaymentCancellationReason2 {
	newValue := new(PaymentCancellationReason2)
	o.CancellationReasonInformation = append(o.CancellationReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction13) AddTransactionInformation() *PaymentTransaction58 {
	newValue := new(PaymentTransaction58)
	o.TransactionInformation = append(o.TransactionInformation, newValue)
	return newValue
}

// Provides details on the original transactions, to which the status report message refers.
type OriginalPaymentInstruction14 struct {

	// Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId"`

	// Number of individual transactions contained in the original payment information group.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original payment information group, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Specifies the status of the payment information group.
	PaymentInformationStatus *TransactionGroupStatus3Code `xml:"PmtInfSts,omitempty"`

	// Provides detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Detailed information on the number of transactions for each identical transaction status.
	NumberOfTransactionsPerStatus []*NumberOfTransactionsPerStatus3 `xml:"NbOfTxsPerSts,omitempty"`

	// Provides information on the original transactions to which the status report message refers.
	TransactionInformationAndStatus []*PaymentTransaction59 `xml:"TxInfAndSts,omitempty"`
}

func (o *OriginalPaymentInstruction14) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction14) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInstruction14) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInstruction14) SetPaymentInformationStatus(value string) {
	o.PaymentInformationStatus = (*TransactionGroupStatus3Code)(&value)
}

func (o *OriginalPaymentInstruction14) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	o.StatusReasonInformation = append(o.StatusReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction14) AddNumberOfTransactionsPerStatus() *NumberOfTransactionsPerStatus3 {
	newValue := new(NumberOfTransactionsPerStatus3)
	o.NumberOfTransactionsPerStatus = append(o.NumberOfTransactionsPerStatus, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction14) AddTransactionInformationAndStatus() *PaymentTransaction59 {
	newValue := new(PaymentTransaction59)
	o.TransactionInformationAndStatus = append(o.TransactionInformationAndStatus, newValue)
	return newValue
}

// Provides details on the reference and status of the original transactions, included in the original instruction, to which the cancellation request message applies.
type OriginalPaymentInstruction15 struct {

	// Unique identification, as assigned by the assigner, to unambiguously identify the cancellation request.
	//
	// Usage: The cancellation request identification can be used for reconciliation or to link tasks relating to the cancellation request.
	PaymentCancellationIdentification *Max35Text `xml:"PmtCxlId,omitempty"`

	// Identifies the case.
	Case *Case3 `xml:"Case,omitempty"`

	// Unique and unambiguous identifier of the original payment information block, as assigned by the original sending party.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId"`

	// Information concerning the original group of transactions, to which the message refers.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Number of individual transactions contained in the cancellation payment information group.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs,omitempty"`

	// Total of all individual amounts included in the cancellation payment information group, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Indicates whether or not the cancellation applies to a whole group of transactions or to individual transactions within the original group.
	PaymentInformationCancellation *GroupCancellationIndicator `xml:"PmtInfCxl,omitempty"`

	// Detailed information on the cancellation reason.
	CancellationReasonInformation []*PaymentCancellationReason2 `xml:"CxlRsnInf,omitempty"`

	// Information concerning the original transactions, to which the cancellation request message refers.
	TransactionInformation []*PaymentTransaction61 `xml:"TxInf,omitempty"`
}

func (o *OriginalPaymentInstruction15) SetPaymentCancellationIdentification(value string) {
	o.PaymentCancellationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction15) AddCase() *Case3 {
	o.Case = new(Case3)
	return o.Case
}

func (o *OriginalPaymentInstruction15) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction15) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	o.OriginalGroupInformation = new(OriginalGroupInformation3)
	return o.OriginalGroupInformation
}

func (o *OriginalPaymentInstruction15) SetNumberOfTransactions(value string) {
	o.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInstruction15) SetControlSum(value string) {
	o.ControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInstruction15) SetPaymentInformationCancellation(value string) {
	o.PaymentInformationCancellation = (*GroupCancellationIndicator)(&value)
}

func (o *OriginalPaymentInstruction15) AddCancellationReasonInformation() *PaymentCancellationReason2 {
	newValue := new(PaymentCancellationReason2)
	o.CancellationReasonInformation = append(o.CancellationReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction15) AddTransactionInformation() *PaymentTransaction61 {
	newValue := new(PaymentTransaction61)
	o.TransactionInformation = append(o.TransactionInformation, newValue)
	return newValue
}

// Provides details on the original transactions, to which the status report message refers.
type OriginalPaymentInstruction16 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reversed payment information group.
	// Usage: The instructing party is the party sending the reversal message and not the party that sent the original instruction that is being reversed.
	ReversalPaymentInformationIdentification *Max35Text `xml:"RvslPmtInfId,omitempty"`

	// Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId"`

	// Number of individual transactions contained in the original payment information group.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original payment information group, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Indicates whether or not the reversal applies to the complete original payment information group or to individual transactions within that group.
	PaymentInformationReversal *TrueFalseIndicator `xml:"PmtInfRvsl,omitempty"`

	// Provides detailed information on the reversal reason.
	ReversalReasonInformation []*PaymentReversalReason7 `xml:"RvslRsnInf,omitempty"`

	// Provides information on the original transactions to which the reversal message refers.
	TransactionInformation []*PaymentTransaction64 `xml:"TxInf,omitempty"`
}

func (o *OriginalPaymentInstruction16) SetReversalPaymentInformationIdentification(value string) {
	o.ReversalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction16) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction16) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInstruction16) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInstruction16) SetBatchBooking(value string) {
	o.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (o *OriginalPaymentInstruction16) SetPaymentInformationReversal(value string) {
	o.PaymentInformationReversal = (*TrueFalseIndicator)(&value)
}

func (o *OriginalPaymentInstruction16) AddReversalReasonInformation() *PaymentReversalReason7 {
	newValue := new(PaymentReversalReason7)
	o.ReversalReasonInformation = append(o.ReversalReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction16) AddTransactionInformation() *PaymentTransaction64 {
	newValue := new(PaymentTransaction64)
	o.TransactionInformation = append(o.TransactionInformation, newValue)
	return newValue
}

// Provides the details on the reference and status of the original transactions, included in the original instruction, to which the cancellation request message applies.
type OriginalPaymentInstruction17 struct {

	// Unique identification, as assigned by the original assigner, to unambiguously identify the original payment information cancellation request.
	OriginalPaymentInformationCancellationIdentification *Max35Text `xml:"OrgnlPmtInfCxlId,omitempty"`

	// Identifies the resolved case.
	ResolvedCase *Case3 `xml:"RslvdCase,omitempty"`

	// Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId"`

	// Provides information on the original message.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Number of individual transactions contained in the original payment information group.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original payment information group, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Specifies the status of a cancellation request, related to a payment information group.
	PaymentInformationCancellationStatus *GroupCancellationStatus1Code `xml:"PmtInfCxlSts,omitempty"`

	// Provides detailed information on the cancellation status reason.
	CancellationStatusReasonInformation []*CancellationStatusReason2 `xml:"CxlStsRsnInf,omitempty"`

	// Detailed information on the number of transactions for each identical cancellation status.
	NumberOfTransactionsPerCancellationStatus []*NumberOfCancellationsPerStatus1 `xml:"NbOfTxsPerCxlSts,omitempty"`

	// Provides information on the original transactions to which the cancellation request message refers.
	TransactionInformationAndStatus []*PaymentTransaction66 `xml:"TxInfAndSts,omitempty"`
}

func (o *OriginalPaymentInstruction17) SetOriginalPaymentInformationCancellationIdentification(value string) {
	o.OriginalPaymentInformationCancellationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction17) AddResolvedCase() *Case3 {
	o.ResolvedCase = new(Case3)
	return o.ResolvedCase
}

func (o *OriginalPaymentInstruction17) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction17) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	o.OriginalGroupInformation = new(OriginalGroupInformation3)
	return o.OriginalGroupInformation
}

func (o *OriginalPaymentInstruction17) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInstruction17) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInstruction17) SetPaymentInformationCancellationStatus(value string) {
	o.PaymentInformationCancellationStatus = (*GroupCancellationStatus1Code)(&value)
}

func (o *OriginalPaymentInstruction17) AddCancellationStatusReasonInformation() *CancellationStatusReason2 {
	newValue := new(CancellationStatusReason2)
	o.CancellationStatusReasonInformation = append(o.CancellationStatusReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction17) AddNumberOfTransactionsPerCancellationStatus() *NumberOfCancellationsPerStatus1 {
	newValue := new(NumberOfCancellationsPerStatus1)
	o.NumberOfTransactionsPerCancellationStatus = append(o.NumberOfTransactionsPerCancellationStatus, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction17) AddTransactionInformationAndStatus() *PaymentTransaction66 {
	newValue := new(PaymentTransaction66)
	o.TransactionInformationAndStatus = append(o.TransactionInformationAndStatus, newValue)
	return newValue
}

// Provides details information on the original transactions, to which the status report message refers.
type OriginalPaymentInstruction18 struct {

	// Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId"`

	// Number of individual transactions contained in the original payment information group.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original payment information group, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Specifies the status of the payment information group.
	PaymentInformationStatus *TransactionGroupStatus3Code `xml:"PmtInfSts,omitempty"`

	// Provides detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Detailed information on the number of transactions for each identical transaction status.
	NumberOfTransactionsPerStatus []*NumberOfTransactionsPerStatus3 `xml:"NbOfTxsPerSts,omitempty"`

	// Provides information on the original transactions to which the status report message refers.
	TransactionInformationAndStatus []*PaymentTransaction68 `xml:"TxInfAndSts,omitempty"`
}

func (o *OriginalPaymentInstruction18) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction18) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInstruction18) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInstruction18) SetPaymentInformationStatus(value string) {
	o.PaymentInformationStatus = (*TransactionGroupStatus3Code)(&value)
}

func (o *OriginalPaymentInstruction18) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	o.StatusReasonInformation = append(o.StatusReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction18) AddNumberOfTransactionsPerStatus() *NumberOfTransactionsPerStatus3 {
	newValue := new(NumberOfTransactionsPerStatus3)
	o.NumberOfTransactionsPerStatus = append(o.NumberOfTransactionsPerStatus, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction18) AddTransactionInformationAndStatus() *PaymentTransaction68 {
	newValue := new(PaymentTransaction68)
	o.TransactionInformationAndStatus = append(o.TransactionInformationAndStatus, newValue)
	return newValue
}

// Provides details on the original transactions, to which the status report message refers.
type OriginalPaymentInstruction19 struct {

	// Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId"`

	// Number of individual transactions contained in the original payment information group.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original payment information group, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Specifies the status of the payment information group.
	PaymentInformationStatus *TransactionGroupStatus3Code `xml:"PmtInfSts,omitempty"`

	// Provides detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Detailed information on the number of transactions for each identical transaction status.
	NumberOfTransactionsPerStatus []*NumberOfTransactionsPerStatus3 `xml:"NbOfTxsPerSts,omitempty"`

	// Provides information on the original transactions to which the status report message refers.
	TransactionInformationAndStatus []*PaymentTransaction69 `xml:"TxInfAndSts,omitempty"`
}

func (o *OriginalPaymentInstruction19) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction19) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInstruction19) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInstruction19) SetPaymentInformationStatus(value string) {
	o.PaymentInformationStatus = (*TransactionGroupStatus3Code)(&value)
}

func (o *OriginalPaymentInstruction19) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	o.StatusReasonInformation = append(o.StatusReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction19) AddNumberOfTransactionsPerStatus() *NumberOfTransactionsPerStatus3 {
	newValue := new(NumberOfTransactionsPerStatus3)
	o.NumberOfTransactionsPerStatus = append(o.NumberOfTransactionsPerStatus, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction19) AddTransactionInformationAndStatus() *PaymentTransaction69 {
	newValue := new(PaymentTransaction69)
	o.TransactionInformationAndStatus = append(o.TransactionInformationAndStatus, newValue)
	return newValue
}

// Provides details on the original transactions, to which the status report message refers.
type OriginalPaymentInstruction2 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reversed payment information group.
	// Usage: The instructing party is the party sending the reversal message and not the party that sent the original instruction that is being reversed.
	ReversalPaymentInformationIdentification *Max35Text `xml:"RvslPmtInfId,omitempty"`

	// Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId"`

	// Number of individual transactions contained in the original payment information group.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original payment information group, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Indicates whether or not the reversal applies to the complete original payment information group or to individual transactions within that group.
	PaymentInformationReversal *TrueFalseIndicator `xml:"PmtInfRvsl,omitempty"`

	// Provides detailed information on the reversal reason.
	ReversalReasonInformation []*PaymentReversalReason7 `xml:"RvslRsnInf,omitempty"`

	// Provides information on the original transactions to which the reversal message refers.
	TransactionInformation []*PaymentTransaction35 `xml:"TxInf,omitempty"`
}

func (o *OriginalPaymentInstruction2) SetReversalPaymentInformationIdentification(value string) {
	o.ReversalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction2) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction2) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInstruction2) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInstruction2) SetBatchBooking(value string) {
	o.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (o *OriginalPaymentInstruction2) SetPaymentInformationReversal(value string) {
	o.PaymentInformationReversal = (*TrueFalseIndicator)(&value)
}

func (o *OriginalPaymentInstruction2) AddReversalReasonInformation() *PaymentReversalReason7 {
	newValue := new(PaymentReversalReason7)
	o.ReversalReasonInformation = append(o.ReversalReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction2) AddTransactionInformation() *PaymentTransaction35 {
	newValue := new(PaymentTransaction35)
	o.TransactionInformation = append(o.TransactionInformation, newValue)
	return newValue
}

// Provides details on the reference and status of the original transactions, included in the original instruction, to which the cancellation request message applies.
type OriginalPaymentInstruction20 struct {

	// Unique identification, as assigned by the assigner, to unambiguously identify the cancellation request.
	//
	// Usage: The cancellation request identification can be used for reconciliation or to link tasks relating to the cancellation request.
	PaymentCancellationIdentification *Max35Text `xml:"PmtCxlId,omitempty"`

	// Identifies the case.
	Case *Case3 `xml:"Case,omitempty"`

	// Unique and unambiguous identifier of the original payment information block, as assigned by the original sending party.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId"`

	// Information concerning the original group of transactions, to which the message refers.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Number of individual transactions contained in the cancellation payment information group.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs,omitempty"`

	// Total of all individual amounts included in the cancellation payment information group, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Indicates whether or not the cancellation applies to a whole group of transactions or to individual transactions within the original group.
	PaymentInformationCancellation *GroupCancellationIndicator `xml:"PmtInfCxl,omitempty"`

	// Detailed information on the cancellation reason.
	CancellationReasonInformation []*PaymentCancellationReason3 `xml:"CxlRsnInf,omitempty"`

	// Information concerning the original transactions, to which the cancellation request message refers.
	TransactionInformation []*PaymentTransaction74 `xml:"TxInf,omitempty"`
}

func (o *OriginalPaymentInstruction20) SetPaymentCancellationIdentification(value string) {
	o.PaymentCancellationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction20) AddCase() *Case3 {
	o.Case = new(Case3)
	return o.Case
}

func (o *OriginalPaymentInstruction20) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction20) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	o.OriginalGroupInformation = new(OriginalGroupInformation3)
	return o.OriginalGroupInformation
}

func (o *OriginalPaymentInstruction20) SetNumberOfTransactions(value string) {
	o.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInstruction20) SetControlSum(value string) {
	o.ControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInstruction20) SetPaymentInformationCancellation(value string) {
	o.PaymentInformationCancellation = (*GroupCancellationIndicator)(&value)
}

func (o *OriginalPaymentInstruction20) AddCancellationReasonInformation() *PaymentCancellationReason3 {
	newValue := new(PaymentCancellationReason3)
	o.CancellationReasonInformation = append(o.CancellationReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction20) AddTransactionInformation() *PaymentTransaction74 {
	newValue := new(PaymentTransaction74)
	o.TransactionInformation = append(o.TransactionInformation, newValue)
	return newValue
}

// Provides details on the original transactions, to which the status report message refers.
type OriginalPaymentInstruction21 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reversed payment information group.
	// Usage: The instructing party is the party sending the reversal message and not the party that sent the original instruction that is being reversed.
	ReversalPaymentInformationIdentification *Max35Text `xml:"RvslPmtInfId,omitempty"`

	// Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId"`

	// Number of individual transactions contained in the original payment information group.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original payment information group, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Indicates whether or not the reversal applies to the complete original payment information group or to individual transactions within that group.
	PaymentInformationReversal *TrueFalseIndicator `xml:"PmtInfRvsl,omitempty"`

	// Provides detailed information on the reversal reason.
	ReversalReasonInformation []*PaymentReversalReason7 `xml:"RvslRsnInf,omitempty"`

	// Provides information on the original transactions to which the reversal message refers.
	TransactionInformation []*PaymentTransaction77 `xml:"TxInf,omitempty"`
}

func (o *OriginalPaymentInstruction21) SetReversalPaymentInformationIdentification(value string) {
	o.ReversalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction21) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction21) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInstruction21) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInstruction21) SetBatchBooking(value string) {
	o.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (o *OriginalPaymentInstruction21) SetPaymentInformationReversal(value string) {
	o.PaymentInformationReversal = (*TrueFalseIndicator)(&value)
}

func (o *OriginalPaymentInstruction21) AddReversalReasonInformation() *PaymentReversalReason7 {
	newValue := new(PaymentReversalReason7)
	o.ReversalReasonInformation = append(o.ReversalReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction21) AddTransactionInformation() *PaymentTransaction77 {
	newValue := new(PaymentTransaction77)
	o.TransactionInformation = append(o.TransactionInformation, newValue)
	return newValue
}

// Provides the details on the reference and status of the original transactions, included in the original instruction, to which the cancellation request message applies.
type OriginalPaymentInstruction22 struct {

	// Unique identification, as assigned by the original assigner, to unambiguously identify the original payment information cancellation request.
	OriginalPaymentInformationCancellationIdentification *Max35Text `xml:"OrgnlPmtInfCxlId,omitempty"`

	// Identifies the resolved case.
	ResolvedCase *Case3 `xml:"RslvdCase,omitempty"`

	// Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId"`

	// Provides information on the original message.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Number of individual transactions contained in the original payment information group.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original payment information group, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Specifies the status of a cancellation request, related to a payment information group.
	PaymentInformationCancellationStatus *GroupCancellationStatus1Code `xml:"PmtInfCxlSts,omitempty"`

	// Provides detailed information on the cancellation status reason.
	CancellationStatusReasonInformation []*CancellationStatusReason2 `xml:"CxlStsRsnInf,omitempty"`

	// Detailed information on the number of transactions for each identical cancellation status.
	NumberOfTransactionsPerCancellationStatus []*NumberOfCancellationsPerStatus1 `xml:"NbOfTxsPerCxlSts,omitempty"`

	// Provides information on the original transactions to which the cancellation request message refers.
	TransactionInformationAndStatus []*PaymentTransaction78 `xml:"TxInfAndSts,omitempty"`
}

func (o *OriginalPaymentInstruction22) SetOriginalPaymentInformationCancellationIdentification(value string) {
	o.OriginalPaymentInformationCancellationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction22) AddResolvedCase() *Case3 {
	o.ResolvedCase = new(Case3)
	return o.ResolvedCase
}

func (o *OriginalPaymentInstruction22) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction22) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	o.OriginalGroupInformation = new(OriginalGroupInformation3)
	return o.OriginalGroupInformation
}

func (o *OriginalPaymentInstruction22) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInstruction22) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInstruction22) SetPaymentInformationCancellationStatus(value string) {
	o.PaymentInformationCancellationStatus = (*GroupCancellationStatus1Code)(&value)
}

func (o *OriginalPaymentInstruction22) AddCancellationStatusReasonInformation() *CancellationStatusReason2 {
	newValue := new(CancellationStatusReason2)
	o.CancellationStatusReasonInformation = append(o.CancellationStatusReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction22) AddNumberOfTransactionsPerCancellationStatus() *NumberOfCancellationsPerStatus1 {
	newValue := new(NumberOfCancellationsPerStatus1)
	o.NumberOfTransactionsPerCancellationStatus = append(o.NumberOfTransactionsPerCancellationStatus, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction22) AddTransactionInformationAndStatus() *PaymentTransaction78 {
	newValue := new(PaymentTransaction78)
	o.TransactionInformationAndStatus = append(o.TransactionInformationAndStatus, newValue)
	return newValue
}

// Provides details information on the original transactions, to which the status report message refers.
type OriginalPaymentInstruction23 struct {

	// Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId"`

	// Number of individual transactions contained in the original payment information group.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original payment information group, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Specifies the status of the payment information group.
	PaymentInformationStatus *ExternalPaymentGroupStatus1Code `xml:"PmtInfSts,omitempty"`

	// Provides detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Detailed information on the number of transactions for each identical transaction status.
	NumberOfTransactionsPerStatus []*NumberOfTransactionsPerStatus5 `xml:"NbOfTxsPerSts,omitempty"`

	// Provides information on the original transactions to which the status report message refers.
	TransactionInformationAndStatus []*PaymentTransaction82 `xml:"TxInfAndSts,omitempty"`
}

func (o *OriginalPaymentInstruction23) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction23) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInstruction23) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInstruction23) SetPaymentInformationStatus(value string) {
	o.PaymentInformationStatus = (*ExternalPaymentGroupStatus1Code)(&value)
}

func (o *OriginalPaymentInstruction23) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	o.StatusReasonInformation = append(o.StatusReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction23) AddNumberOfTransactionsPerStatus() *NumberOfTransactionsPerStatus5 {
	newValue := new(NumberOfTransactionsPerStatus5)
	o.NumberOfTransactionsPerStatus = append(o.NumberOfTransactionsPerStatus, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction23) AddTransactionInformationAndStatus() *PaymentTransaction82 {
	newValue := new(PaymentTransaction82)
	o.TransactionInformationAndStatus = append(o.TransactionInformationAndStatus, newValue)
	return newValue
}

// Provides details on the original transactions, to which the status report message refers.
type OriginalPaymentInstruction24 struct {

	// Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId"`

	// Number of individual transactions contained in the original payment information group.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original payment information group, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Specifies the status of the payment information group.
	PaymentInformationStatus *ExternalPaymentGroupStatus1Code `xml:"PmtInfSts,omitempty"`

	// Provides detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Detailed information on the number of transactions for each identical transaction status.
	NumberOfTransactionsPerStatus []*NumberOfTransactionsPerStatus5 `xml:"NbOfTxsPerSts,omitempty"`

	// Provides information on the original transactions to which the status report message refers.
	TransactionInformationAndStatus []*PaymentTransaction83 `xml:"TxInfAndSts,omitempty"`
}

func (o *OriginalPaymentInstruction24) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction24) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInstruction24) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInstruction24) SetPaymentInformationStatus(value string) {
	o.PaymentInformationStatus = (*ExternalPaymentGroupStatus1Code)(&value)
}

func (o *OriginalPaymentInstruction24) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	o.StatusReasonInformation = append(o.StatusReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction24) AddNumberOfTransactionsPerStatus() *NumberOfTransactionsPerStatus5 {
	newValue := new(NumberOfTransactionsPerStatus5)
	o.NumberOfTransactionsPerStatus = append(o.NumberOfTransactionsPerStatus, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction24) AddTransactionInformationAndStatus() *PaymentTransaction83 {
	newValue := new(PaymentTransaction83)
	o.TransactionInformationAndStatus = append(o.TransactionInformationAndStatus, newValue)
	return newValue
}

// Provides the details on the reference and status of the original transactions, included in the original instruction, to which the cancellation request message applies.
type OriginalPaymentInstruction3 struct {

	// Unique identification, as assigned by the original assigner, to unambiguously identify the original payment information cancellation request.
	OriginalPaymentInformationCancellationIdentification *Max35Text `xml:"OrgnlPmtInfCxlId,omitempty"`

	// Identifies the resolved case.
	ResolvedCase *Case3 `xml:"RslvdCase,omitempty"`

	// Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId"`

	// Provides information on the original message.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Number of individual transactions contained in the original payment information group.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original payment information group, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Specifies the status of a cancellation request, related to a payment information group.
	PaymentInformationCancellationStatus *GroupCancellationStatus1Code `xml:"PmtInfCxlSts,omitempty"`

	// Provides detailed information on the cancellation status reason.
	CancellationStatusReasonInformation []*CancellationStatusReason2 `xml:"CxlStsRsnInf,omitempty"`

	// Detailed information on the number of transactions for each identical cancellation status.
	NumberOfTransactionsPerCancellationStatus []*NumberOfCancellationsPerStatus1 `xml:"NbOfTxsPerCxlSts,omitempty"`

	// Provides information on the original transactions to which the cancellation request message refers.
	TransactionInformationAndStatus []*PaymentTransaction39 `xml:"TxInfAndSts,omitempty"`
}

func (o *OriginalPaymentInstruction3) SetOriginalPaymentInformationCancellationIdentification(value string) {
	o.OriginalPaymentInformationCancellationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction3) AddResolvedCase() *Case3 {
	o.ResolvedCase = new(Case3)
	return o.ResolvedCase
}

func (o *OriginalPaymentInstruction3) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction3) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	o.OriginalGroupInformation = new(OriginalGroupInformation3)
	return o.OriginalGroupInformation
}

func (o *OriginalPaymentInstruction3) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInstruction3) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInstruction3) SetPaymentInformationCancellationStatus(value string) {
	o.PaymentInformationCancellationStatus = (*GroupCancellationStatus1Code)(&value)
}

func (o *OriginalPaymentInstruction3) AddCancellationStatusReasonInformation() *CancellationStatusReason2 {
	newValue := new(CancellationStatusReason2)
	o.CancellationStatusReasonInformation = append(o.CancellationStatusReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction3) AddNumberOfTransactionsPerCancellationStatus() *NumberOfCancellationsPerStatus1 {
	newValue := new(NumberOfCancellationsPerStatus1)
	o.NumberOfTransactionsPerCancellationStatus = append(o.NumberOfTransactionsPerCancellationStatus, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction3) AddTransactionInformationAndStatus() *PaymentTransaction39 {
	newValue := new(PaymentTransaction39)
	o.TransactionInformationAndStatus = append(o.TransactionInformationAndStatus, newValue)
	return newValue
}

// Provides details on the reference and status of the original transactions, included in the original instruction, to which the cancellation request message applies.
type OriginalPaymentInstruction4 struct {

	// Unique identification, as assigned by the assigner, to unambiguously identify the cancellation request.
	//
	// Usage: The cancellation request identification can be used for reconciliation or to link tasks relating to the cancellation request.
	PaymentCancellationIdentification *Max35Text `xml:"PmtCxlId,omitempty"`

	// Identifies the case.
	Case *Case3 `xml:"Case,omitempty"`

	// Unique and unambiguous identifier of the original payment information block, as assigned by the original sending party.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId"`

	// Information concerning the original group of transactions, to which the message refers.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Number of individual transactions contained in the cancellation payment information group.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs,omitempty"`

	// Total of all individual amounts included in the cancellation payment information group, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Indicates whether or not the cancellation applies to a whole group of transactions or to individual transactions within the original group.
	PaymentInformationCancellation *GroupCancellationIndicator `xml:"PmtInfCxl,omitempty"`

	// Detailed information on the cancellation reason.
	CancellationReasonInformation []*PaymentCancellationReason2 `xml:"CxlRsnInf,omitempty"`

	// Information concerning the original transactions, to which the cancellation request message refers.
	TransactionInformation []*PaymentTransaction37 `xml:"TxInf,omitempty"`
}

func (o *OriginalPaymentInstruction4) SetPaymentCancellationIdentification(value string) {
	o.PaymentCancellationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction4) AddCase() *Case3 {
	o.Case = new(Case3)
	return o.Case
}

func (o *OriginalPaymentInstruction4) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction4) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	o.OriginalGroupInformation = new(OriginalGroupInformation3)
	return o.OriginalGroupInformation
}

func (o *OriginalPaymentInstruction4) SetNumberOfTransactions(value string) {
	o.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInstruction4) SetControlSum(value string) {
	o.ControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInstruction4) SetPaymentInformationCancellation(value string) {
	o.PaymentInformationCancellation = (*GroupCancellationIndicator)(&value)
}

func (o *OriginalPaymentInstruction4) AddCancellationReasonInformation() *PaymentCancellationReason2 {
	newValue := new(PaymentCancellationReason2)
	o.CancellationReasonInformation = append(o.CancellationReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction4) AddTransactionInformation() *PaymentTransaction37 {
	newValue := new(PaymentTransaction37)
	o.TransactionInformation = append(o.TransactionInformation, newValue)
	return newValue
}

// Provides details on the original transactions, to which the status report message refers.
type OriginalPaymentInstruction5 struct {

	// Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId"`

	// Number of individual transactions contained in the original payment information group.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original payment information group, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Specifies the status of the payment information group.
	PaymentInformationStatus *TransactionGroupStatus3Code `xml:"PmtInfSts,omitempty"`

	// Provides detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Detailed information on the number of transactions for each identical transaction status.
	NumberOfTransactionsPerStatus []*NumberOfTransactionsPerStatus3 `xml:"NbOfTxsPerSts,omitempty"`

	// Provides information on the original transactions to which the status report message refers.
	TransactionInformationAndStatus []*PaymentTransaction41 `xml:"TxInfAndSts,omitempty"`
}

func (o *OriginalPaymentInstruction5) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction5) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInstruction5) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInstruction5) SetPaymentInformationStatus(value string) {
	o.PaymentInformationStatus = (*TransactionGroupStatus3Code)(&value)
}

func (o *OriginalPaymentInstruction5) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	o.StatusReasonInformation = append(o.StatusReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction5) AddNumberOfTransactionsPerStatus() *NumberOfTransactionsPerStatus3 {
	newValue := new(NumberOfTransactionsPerStatus3)
	o.NumberOfTransactionsPerStatus = append(o.NumberOfTransactionsPerStatus, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction5) AddTransactionInformationAndStatus() *PaymentTransaction41 {
	newValue := new(PaymentTransaction41)
	o.TransactionInformationAndStatus = append(o.TransactionInformationAndStatus, newValue)
	return newValue
}

// Provides details information on the original transactions, to which the status report message refers.
type OriginalPaymentInstruction6 struct {

	// Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId"`

	// Number of individual transactions contained in the original payment information group.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original payment information group, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Specifies the status of the payment information group.
	PaymentInformationStatus *TransactionGroupStatus3Code `xml:"PmtInfSts,omitempty"`

	// Provides detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Detailed information on the number of transactions for each identical transaction status.
	NumberOfTransactionsPerStatus []*NumberOfTransactionsPerStatus3 `xml:"NbOfTxsPerSts,omitempty"`

	// Provides information on the original transactions to which the status report message refers.
	TransactionInformationAndStatus []*PaymentTransaction46 `xml:"TxInfAndSts,omitempty"`
}

func (o *OriginalPaymentInstruction6) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction6) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInstruction6) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInstruction6) SetPaymentInformationStatus(value string) {
	o.PaymentInformationStatus = (*TransactionGroupStatus3Code)(&value)
}

func (o *OriginalPaymentInstruction6) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	o.StatusReasonInformation = append(o.StatusReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction6) AddNumberOfTransactionsPerStatus() *NumberOfTransactionsPerStatus3 {
	newValue := new(NumberOfTransactionsPerStatus3)
	o.NumberOfTransactionsPerStatus = append(o.NumberOfTransactionsPerStatus, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction6) AddTransactionInformationAndStatus() *PaymentTransaction46 {
	newValue := new(PaymentTransaction46)
	o.TransactionInformationAndStatus = append(o.TransactionInformationAndStatus, newValue)
	return newValue
}

// Provides details on the original transactions, to which the status report message refers.
type OriginalPaymentInstruction7 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reversed payment information group.
	// Usage: The instructing party is the party sending the reversal message and not the party that sent the original instruction that is being reversed.
	ReversalPaymentInformationIdentification *Max35Text `xml:"RvslPmtInfId,omitempty"`

	// Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId"`

	// Number of individual transactions contained in the original payment information group.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original payment information group, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Indicates whether or not the reversal applies to the complete original payment information group or to individual transactions within that group.
	PaymentInformationReversal *TrueFalseIndicator `xml:"PmtInfRvsl,omitempty"`

	// Provides detailed information on the reversal reason.
	ReversalReasonInformation []*PaymentReversalReason7 `xml:"RvslRsnInf,omitempty"`

	// Provides information on the original transactions to which the reversal message refers.
	TransactionInformation []*PaymentTransaction42 `xml:"TxInf,omitempty"`
}

func (o *OriginalPaymentInstruction7) SetReversalPaymentInformationIdentification(value string) {
	o.ReversalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction7) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction7) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInstruction7) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInstruction7) SetBatchBooking(value string) {
	o.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (o *OriginalPaymentInstruction7) SetPaymentInformationReversal(value string) {
	o.PaymentInformationReversal = (*TrueFalseIndicator)(&value)
}

func (o *OriginalPaymentInstruction7) AddReversalReasonInformation() *PaymentReversalReason7 {
	newValue := new(PaymentReversalReason7)
	o.ReversalReasonInformation = append(o.ReversalReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction7) AddTransactionInformation() *PaymentTransaction42 {
	newValue := new(PaymentTransaction42)
	o.TransactionInformation = append(o.TransactionInformation, newValue)
	return newValue
}

// Provides details on the reference and status of the original transactions, included in the original instruction, to which the cancellation request message applies.
type OriginalPaymentInstruction8 struct {

	// Unique identification, as assigned by the assigner, to unambiguously identify the cancellation request.
	//
	// Usage: The cancellation request identification can be used for reconciliation or to link tasks relating to the cancellation request.
	PaymentCancellationIdentification *Max35Text `xml:"PmtCxlId,omitempty"`

	// Identifies the case.
	Case *Case3 `xml:"Case,omitempty"`

	// Unique and unambiguous identifier of the original payment information block, as assigned by the original sending party.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId"`

	// Information concerning the original group of transactions, to which the message refers.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Number of individual transactions contained in the cancellation payment information group.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs,omitempty"`

	// Total of all individual amounts included in the cancellation payment information group, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Indicates whether or not the cancellation applies to a whole group of transactions or to individual transactions within the original group.
	PaymentInformationCancellation *GroupCancellationIndicator `xml:"PmtInfCxl,omitempty"`

	// Detailed information on the cancellation reason.
	CancellationReasonInformation []*PaymentCancellationReason2 `xml:"CxlRsnInf,omitempty"`

	// Information concerning the original transactions, to which the cancellation request message refers.
	TransactionInformation []*PaymentTransaction47 `xml:"TxInf,omitempty"`
}

func (o *OriginalPaymentInstruction8) SetPaymentCancellationIdentification(value string) {
	o.PaymentCancellationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction8) AddCase() *Case3 {
	o.Case = new(Case3)
	return o.Case
}

func (o *OriginalPaymentInstruction8) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction8) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	o.OriginalGroupInformation = new(OriginalGroupInformation3)
	return o.OriginalGroupInformation
}

func (o *OriginalPaymentInstruction8) SetNumberOfTransactions(value string) {
	o.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInstruction8) SetControlSum(value string) {
	o.ControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInstruction8) SetPaymentInformationCancellation(value string) {
	o.PaymentInformationCancellation = (*GroupCancellationIndicator)(&value)
}

func (o *OriginalPaymentInstruction8) AddCancellationReasonInformation() *PaymentCancellationReason2 {
	newValue := new(PaymentCancellationReason2)
	o.CancellationReasonInformation = append(o.CancellationReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction8) AddTransactionInformation() *PaymentTransaction47 {
	newValue := new(PaymentTransaction47)
	o.TransactionInformation = append(o.TransactionInformation, newValue)
	return newValue
}

// Provides details on the original transactions, to which the status report message refers.
type OriginalPaymentInstruction9 struct {

	// Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId"`

	// Number of individual transactions contained in the original payment information group.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original payment information group, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Specifies the status of the payment information group.
	PaymentInformationStatus *TransactionGroupStatus3Code `xml:"PmtInfSts,omitempty"`

	// Provides detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Detailed information on the number of transactions for each identical transaction status.
	NumberOfTransactionsPerStatus []*NumberOfTransactionsPerStatus3 `xml:"NbOfTxsPerSts,omitempty"`

	// Provides information on the original transactions to which the status report message refers.
	TransactionInformationAndStatus []*PaymentTransaction49 `xml:"TxInfAndSts,omitempty"`
}

func (o *OriginalPaymentInstruction9) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction9) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInstruction9) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInstruction9) SetPaymentInformationStatus(value string) {
	o.PaymentInformationStatus = (*TransactionGroupStatus3Code)(&value)
}

func (o *OriginalPaymentInstruction9) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	o.StatusReasonInformation = append(o.StatusReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction9) AddNumberOfTransactionsPerStatus() *NumberOfTransactionsPerStatus3 {
	newValue := new(NumberOfTransactionsPerStatus3)
	o.NumberOfTransactionsPerStatus = append(o.NumberOfTransactionsPerStatus, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction9) AddTransactionInformationAndStatus() *PaymentTransaction49 {
	newValue := new(PaymentTransaction49)
	o.TransactionInformationAndStatus = append(o.TransactionInformationAndStatus, newValue)
	return newValue
}
