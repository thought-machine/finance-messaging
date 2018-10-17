package iso20022

// Set of elements used to provide information on the original transactions, to which the status report message refers.
type OriginalPaymentInformation1 struct {

	// Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId"`

	// Number of individual transactions contained in the original payment information group.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original payment information group, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Specifies the status of the payment information group.
	PaymentInformationStatus *TransactionGroupStatus3Code `xml:"PmtInfSts,omitempty"`

	// Set of elements used to provide detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation8 `xml:"StsRsnInf,omitempty"`

	// Detailed information on the number of transactions for each identical transaction status.
	NumberOfTransactionsPerStatus []*NumberOfTransactionsPerStatus3 `xml:"NbOfTxsPerSts,omitempty"`

	// Set of elements used to provide information on the original transactions to which the status report message refers.
	TransactionInformationAndStatus []*PaymentTransactionInformation25 `xml:"TxInfAndSts,omitempty"`
}

func (o *OriginalPaymentInformation1) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInformation1) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInformation1) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInformation1) SetPaymentInformationStatus(value string) {
	o.PaymentInformationStatus = (*TransactionGroupStatus3Code)(&value)
}

func (o *OriginalPaymentInformation1) AddStatusReasonInformation() *StatusReasonInformation8 {
	newValue := new(StatusReasonInformation8)
	o.StatusReasonInformation = append(o.StatusReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInformation1) AddNumberOfTransactionsPerStatus() *NumberOfTransactionsPerStatus3 {
	newValue := new(NumberOfTransactionsPerStatus3)
	o.NumberOfTransactionsPerStatus = append(o.NumberOfTransactionsPerStatus, newValue)
	return newValue
}

func (o *OriginalPaymentInformation1) AddTransactionInformationAndStatus() *PaymentTransactionInformation25 {
	newValue := new(PaymentTransactionInformation25)
	o.TransactionInformationAndStatus = append(o.TransactionInformationAndStatus, newValue)
	return newValue
}

// Set of elements used to provide information on the original transactions, to which the status report message refers.
type OriginalPaymentInformation2 struct {

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

	// Set of elements used to provide detailed information on the reversal reason.
	ReversalReasonInformation []*ReversalReasonInformation6 `xml:"RvslRsnInf,omitempty"`

	// Set of elements used to provide information on the original transactions to which the reversal message refers.
	TransactionInformation []*PaymentTransactionInformation28 `xml:"TxInf,omitempty"`
}

func (o *OriginalPaymentInformation2) SetReversalPaymentInformationIdentification(value string) {
	o.ReversalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInformation2) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInformation2) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInformation2) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInformation2) SetBatchBooking(value string) {
	o.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (o *OriginalPaymentInformation2) SetPaymentInformationReversal(value string) {
	o.PaymentInformationReversal = (*TrueFalseIndicator)(&value)
}

func (o *OriginalPaymentInformation2) AddReversalReasonInformation() *ReversalReasonInformation6 {
	newValue := new(ReversalReasonInformation6)
	o.ReversalReasonInformation = append(o.ReversalReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInformation2) AddTransactionInformation() *PaymentTransactionInformation28 {
	newValue := new(PaymentTransactionInformation28)
	o.TransactionInformation = append(o.TransactionInformation, newValue)
	return newValue
}

// Set of elements used to provide reference and status information on the original transactions, included in the original instruction, to which the cancellation request message applies.
type OriginalPaymentInformation3 struct {

	// Unique identification, as assigned by the original assigner, to unambiguously identify the original payment information cancellation request.
	OriginalPaymentInformationCancellationIdentification *Max35Text `xml:"OrgnlPmtInfCxlId,omitempty"`

	// Identifies the resolved case.
	ResolvedCase *Case2 `xml:"RslvdCase,omitempty"`

	// Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId"`

	// Set of elements used to provide information on the original messsage.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Number of individual transactions contained in the original payment information group.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original payment information group, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Specifies the status of a cancellation request, related to a payment information group.
	PaymentInformationCancellationStatus *GroupCancellationStatus1Code `xml:"PmtInfCxlSts,omitempty"`

	// Set of elements used to provide detailed information on the cancellation status reason.
	CancellationStatusReasonInformation []*CancellationStatusReasonInformation1 `xml:"CxlStsRsnInf,omitempty"`

	// Detailed information on the number of transactions for each identical cancellation status.
	NumberOfTransactionsPerCancellationStatus []*NumberOfCancellationsPerStatus1 `xml:"NbOfTxsPerCxlSts,omitempty"`

	// Set of elements used to provide information on the original transactions to which the cancellation request message refers.
	TransactionInformationAndStatus []*PaymentTransactionInformation32 `xml:"TxInfAndSts,omitempty"`
}

func (o *OriginalPaymentInformation3) SetOriginalPaymentInformationCancellationIdentification(value string) {
	o.OriginalPaymentInformationCancellationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInformation3) AddResolvedCase() *Case2 {
	o.ResolvedCase = new(Case2)
	return o.ResolvedCase
}

func (o *OriginalPaymentInformation3) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInformation3) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	o.OriginalGroupInformation = new(OriginalGroupInformation3)
	return o.OriginalGroupInformation
}

func (o *OriginalPaymentInformation3) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInformation3) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInformation3) SetPaymentInformationCancellationStatus(value string) {
	o.PaymentInformationCancellationStatus = (*GroupCancellationStatus1Code)(&value)
}

func (o *OriginalPaymentInformation3) AddCancellationStatusReasonInformation() *CancellationStatusReasonInformation1 {
	newValue := new(CancellationStatusReasonInformation1)
	o.CancellationStatusReasonInformation = append(o.CancellationStatusReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInformation3) AddNumberOfTransactionsPerCancellationStatus() *NumberOfCancellationsPerStatus1 {
	newValue := new(NumberOfCancellationsPerStatus1)
	o.NumberOfTransactionsPerCancellationStatus = append(o.NumberOfTransactionsPerCancellationStatus, newValue)
	return newValue
}

func (o *OriginalPaymentInformation3) AddTransactionInformationAndStatus() *PaymentTransactionInformation32 {
	newValue := new(PaymentTransactionInformation32)
	o.TransactionInformationAndStatus = append(o.TransactionInformationAndStatus, newValue)
	return newValue
}

// Set of elements used to provide reference and status information on the original transactions, included in the original instruction, to which the cancellation request message applies.
type OriginalPaymentInformation4 struct {

	// Unique identification, as assigned by the assigner, to unambiguously identify the cancellation request.
	//
	// Usage: The cancellation request identification can be used for reconciliation or to link tasks relating to the cancellation request.
	PaymentCancellationIdentification *Max35Text `xml:"PmtCxlId,omitempty"`

	// Identifies the case.
	Case *Case2 `xml:"Case,omitempty"`

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
	CancellationReasonInformation []*CancellationReasonInformation3 `xml:"CxlRsnInf,omitempty"`

	// Information concerning the original transactions, to which the cancellation request message refers.
	TransactionInformation []*PaymentTransactionInformation30 `xml:"TxInf,omitempty"`
}

func (o *OriginalPaymentInformation4) SetPaymentCancellationIdentification(value string) {
	o.PaymentCancellationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInformation4) AddCase() *Case2 {
	o.Case = new(Case2)
	return o.Case
}

func (o *OriginalPaymentInformation4) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInformation4) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	o.OriginalGroupInformation = new(OriginalGroupInformation3)
	return o.OriginalGroupInformation
}

func (o *OriginalPaymentInformation4) SetNumberOfTransactions(value string) {
	o.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInformation4) SetControlSum(value string) {
	o.ControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInformation4) SetPaymentInformationCancellation(value string) {
	o.PaymentInformationCancellation = (*GroupCancellationIndicator)(&value)
}

func (o *OriginalPaymentInformation4) AddCancellationReasonInformation() *CancellationReasonInformation3 {
	newValue := new(CancellationReasonInformation3)
	o.CancellationReasonInformation = append(o.CancellationReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInformation4) AddTransactionInformation() *PaymentTransactionInformation30 {
	newValue := new(PaymentTransactionInformation30)
	o.TransactionInformation = append(o.TransactionInformation, newValue)
	return newValue
}

// Set of elements used to provide information on the original transactions, to which the status report message refers.
type OriginalPaymentInformation5 struct {

	// Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId"`

	// Number of individual transactions contained in the original payment information group.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original payment information group, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Specifies the status of the payment information group.
	PaymentInformationStatus *TransactionGroupStatus3Code `xml:"PmtInfSts,omitempty"`

	// Set of elements used to provide detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Detailed information on the number of transactions for each identical transaction status.
	NumberOfTransactionsPerStatus []*NumberOfTransactionsPerStatus3 `xml:"NbOfTxsPerSts,omitempty"`

	// Set of elements used to provide information on the original transactions to which the status report message refers.
	TransactionInformationAndStatus []*PaymentTransactionInformation34 `xml:"TxInfAndSts,omitempty"`
}

func (o *OriginalPaymentInformation5) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInformation5) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInformation5) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInformation5) SetPaymentInformationStatus(value string) {
	o.PaymentInformationStatus = (*TransactionGroupStatus3Code)(&value)
}

func (o *OriginalPaymentInformation5) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	o.StatusReasonInformation = append(o.StatusReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInformation5) AddNumberOfTransactionsPerStatus() *NumberOfTransactionsPerStatus3 {
	newValue := new(NumberOfTransactionsPerStatus3)
	o.NumberOfTransactionsPerStatus = append(o.NumberOfTransactionsPerStatus, newValue)
	return newValue
}

func (o *OriginalPaymentInformation5) AddTransactionInformationAndStatus() *PaymentTransactionInformation34 {
	newValue := new(PaymentTransactionInformation34)
	o.TransactionInformationAndStatus = append(o.TransactionInformationAndStatus, newValue)
	return newValue
}

// Provides information on the original payment transaction, to which the remittance message applies.
type OriginalPaymentInformation6 struct {

	// Identifies the underlying transaction.
	References *TransactionReferences4 `xml:"Refs"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation19 `xml:"PmtTpInf,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	Amount *AmountType3Choice `xml:"Amt,omitempty"`

	// Provides details of the currency exchange rate and contract.
	ExchangeRateInformation *ExchangeRate1 `xml:"XchgRateInf,omitempty"`

	// Date at which the initiating party requests the clearing agent to process the payment.
	// Usage: This is the date on which the debtor's account is to be debited. If payment by cheque, the date when the cheque must be generated by the bank.
	RequestedExecutionDate *ISODate `xml:"ReqdExctnDt,omitempty"`

	// Date and time at which the creditor requests that the amount of money is to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr,omitempty"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr,omitempty"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt,omitempty"`
}

func (o *OriginalPaymentInformation6) AddReferences() *TransactionReferences4 {
	o.References = new(TransactionReferences4)
	return o.References
}

func (o *OriginalPaymentInformation6) AddPaymentTypeInformation() *PaymentTypeInformation19 {
	o.PaymentTypeInformation = new(PaymentTypeInformation19)
	return o.PaymentTypeInformation
}

func (o *OriginalPaymentInformation6) AddAmount() *AmountType3Choice {
	o.Amount = new(AmountType3Choice)
	return o.Amount
}

func (o *OriginalPaymentInformation6) AddExchangeRateInformation() *ExchangeRate1 {
	o.ExchangeRateInformation = new(ExchangeRate1)
	return o.ExchangeRateInformation
}

func (o *OriginalPaymentInformation6) SetRequestedExecutionDate(value string) {
	o.RequestedExecutionDate = (*ISODate)(&value)
}

func (o *OriginalPaymentInformation6) SetRequestedCollectionDate(value string) {
	o.RequestedCollectionDate = (*ISODate)(&value)
}

func (o *OriginalPaymentInformation6) AddDebtor() *PartyIdentification43 {
	o.Debtor = new(PartyIdentification43)
	return o.Debtor
}

func (o *OriginalPaymentInformation6) AddDebtorAccount() *CashAccount24 {
	o.DebtorAccount = new(CashAccount24)
	return o.DebtorAccount
}

func (o *OriginalPaymentInformation6) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	o.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return o.DebtorAgent
}

func (o *OriginalPaymentInformation6) AddCreditor() *PartyIdentification43 {
	o.Creditor = new(PartyIdentification43)
	return o.Creditor
}

func (o *OriginalPaymentInformation6) AddCreditorAccount() *CashAccount24 {
	o.CreditorAccount = new(CashAccount24)
	return o.CreditorAccount
}

func (o *OriginalPaymentInformation6) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	o.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return o.CreditorAgent
}

// Provides information on the original payment transaction, to which the remittance message applies.
type OriginalPaymentInformation7 struct {

	// Identifies the underlying transaction.
	References *TransactionReferences4 `xml:"Refs"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation19 `xml:"PmtTpInf,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	Amount *AmountType3Choice `xml:"Amt,omitempty"`

	// Provides details of the currency exchange rate and contract.
	ExchangeRateInformation *ExchangeRate1 `xml:"XchgRateInf,omitempty"`

	// Date at which the initiating party requests the clearing agent to process the payment.
	// Usage: This is the date on which the debtor's account is to be debited. If payment by cheque, the date when the cheque must be generated by the bank.
	RequestedExecutionDate *DateAndDateTimeChoice `xml:"ReqdExctnDt,omitempty"`

	// Date and time at which the creditor requests that the amount of money is to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr,omitempty"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr,omitempty"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt,omitempty"`
}

func (o *OriginalPaymentInformation7) AddReferences() *TransactionReferences4 {
	o.References = new(TransactionReferences4)
	return o.References
}

func (o *OriginalPaymentInformation7) AddPaymentTypeInformation() *PaymentTypeInformation19 {
	o.PaymentTypeInformation = new(PaymentTypeInformation19)
	return o.PaymentTypeInformation
}

func (o *OriginalPaymentInformation7) AddAmount() *AmountType3Choice {
	o.Amount = new(AmountType3Choice)
	return o.Amount
}

func (o *OriginalPaymentInformation7) AddExchangeRateInformation() *ExchangeRate1 {
	o.ExchangeRateInformation = new(ExchangeRate1)
	return o.ExchangeRateInformation
}

func (o *OriginalPaymentInformation7) AddRequestedExecutionDate() *DateAndDateTimeChoice {
	o.RequestedExecutionDate = new(DateAndDateTimeChoice)
	return o.RequestedExecutionDate
}

func (o *OriginalPaymentInformation7) SetRequestedCollectionDate(value string) {
	o.RequestedCollectionDate = (*ISODate)(&value)
}

func (o *OriginalPaymentInformation7) AddDebtor() *PartyIdentification43 {
	o.Debtor = new(PartyIdentification43)
	return o.Debtor
}

func (o *OriginalPaymentInformation7) AddDebtorAccount() *CashAccount24 {
	o.DebtorAccount = new(CashAccount24)
	return o.DebtorAccount
}

func (o *OriginalPaymentInformation7) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	o.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return o.DebtorAgent
}

func (o *OriginalPaymentInformation7) AddCreditor() *PartyIdentification43 {
	o.Creditor = new(PartyIdentification43)
	return o.Creditor
}

func (o *OriginalPaymentInformation7) AddCreditorAccount() *CashAccount24 {
	o.CreditorAccount = new(CashAccount24)
	return o.CreditorAccount
}

func (o *OriginalPaymentInformation7) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	o.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return o.CreditorAgent
}
