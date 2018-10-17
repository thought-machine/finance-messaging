package iso20022

// Characteristics that apply to the credit side of the payment transactions included in the direct debit initiation.
type PaymentInstruction10 struct {

	// Unique identification, as assigned by a sending party, to unambiguously identify the payment information group within the message.
	PaymentInformationIdentification *Max35Text `xml:"PmtInfId"`

	// Specifies the means of payment that will be used to move the amount of money.
	PaymentMethod *PaymentMethod2Code `xml:"PmtMtd"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the payment information group.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs,omitempty"`

	// Total of all individual amounts included in the group, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation24 `xml:"PmtTpInf,omitempty"`

	// Date and time at which the creditor requests that the amount of money is to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent in the payment chain.
	CreditorAgentAccount *CashAccount24 `xml:"CdtrAgtAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Account used to process charges associated with a transaction.
	//
	// Usage: Charges account should be used when charges have to be booked to an account different from the account identified in debtor's account.
	ChargesAccount *CashAccount24 `xml:"ChrgsAcct,omitempty"`

	// Agent that services a charges account.
	//
	// Usage: Charges account agent should only be used when the charges account agent is different from the creditor agent.
	ChargesAccountAgent *BranchAndFinancialInstitutionIdentification5 `xml:"ChrgsAcctAgt,omitempty"`

	// Credit party that signs the mandate.
	CreditorSchemeIdentification *PartyIdentification43 `xml:"CdtrSchmeId,omitempty"`

	// Provides information on the individual transaction(s) included in the message.
	DirectDebitTransactionInformation []*DirectDebitTransactionInformation13 `xml:"DrctDbtTxInf"`
}

func (p *PaymentInstruction10) SetPaymentInformationIdentification(value string) {
	p.PaymentInformationIdentification = (*Max35Text)(&value)
}

func (p *PaymentInstruction10) SetPaymentMethod(value string) {
	p.PaymentMethod = (*PaymentMethod2Code)(&value)
}

func (p *PaymentInstruction10) SetBatchBooking(value string) {
	p.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (p *PaymentInstruction10) SetNumberOfTransactions(value string) {
	p.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (p *PaymentInstruction10) SetControlSum(value string) {
	p.ControlSum = (*DecimalNumber)(&value)
}

func (p *PaymentInstruction10) AddPaymentTypeInformation() *PaymentTypeInformation24 {
	p.PaymentTypeInformation = new(PaymentTypeInformation24)
	return p.PaymentTypeInformation
}

func (p *PaymentInstruction10) SetRequestedCollectionDate(value string) {
	p.RequestedCollectionDate = (*ISODate)(&value)
}

func (p *PaymentInstruction10) AddCreditor() *PartyIdentification43 {
	p.Creditor = new(PartyIdentification43)
	return p.Creditor
}

func (p *PaymentInstruction10) AddCreditorAccount() *CashAccount24 {
	p.CreditorAccount = new(CashAccount24)
	return p.CreditorAccount
}

func (p *PaymentInstruction10) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.CreditorAgent
}

func (p *PaymentInstruction10) AddCreditorAgentAccount() *CashAccount24 {
	p.CreditorAgentAccount = new(CashAccount24)
	return p.CreditorAgentAccount
}

func (p *PaymentInstruction10) AddUltimateCreditor() *PartyIdentification43 {
	p.UltimateCreditor = new(PartyIdentification43)
	return p.UltimateCreditor
}

func (p *PaymentInstruction10) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentInstruction10) AddChargesAccount() *CashAccount24 {
	p.ChargesAccount = new(CashAccount24)
	return p.ChargesAccount
}

func (p *PaymentInstruction10) AddChargesAccountAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.ChargesAccountAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.ChargesAccountAgent
}

func (p *PaymentInstruction10) AddCreditorSchemeIdentification() *PartyIdentification43 {
	p.CreditorSchemeIdentification = new(PartyIdentification43)
	return p.CreditorSchemeIdentification
}

func (p *PaymentInstruction10) AddDirectDebitTransactionInformation() *DirectDebitTransactionInformation13 {
	newValue := new(DirectDebitTransactionInformation13)
	p.DirectDebitTransactionInformation = append(p.DirectDebitTransactionInformation, newValue)
	return newValue
}

// Instruction to pay an amount of money to an ultimate beneficiary, on behalf of an originator. This instruction may have to be forwarded several times to complete the settlement chain.
type PaymentInstruction11 struct {

	// Reference assigned by a sending party to unambiguously identify the payment information block within the message.
	PaymentInformationIdentification *Max35Text `xml:"PmtInfId,omitempty"`

	// Specifies the means of payment that will be used to move the amount of money.
	PaymentMethod *PaymentMethod7Code `xml:"PmtMtd"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation19 `xml:"PmtTpInf,omitempty"`

	// Date at which the initiating party requests the clearing agent to process the payment. If payment by cheque, the date when the cheque must be generated by the bank.
	//
	// Usage: This is the date on which the debtor's account(s) is (are) to be debited.
	RequestedExecutionDate *ISODate `xml:"ReqdExctnDt"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr"`

	// Account used to process charges associated with a transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Payment processes required to transfer cash from the debtor to the creditor.
	CreditTransferTransaction []*CreditTransferTransaction10 `xml:"CdtTrfTx"`
}

func (p *PaymentInstruction11) SetPaymentInformationIdentification(value string) {
	p.PaymentInformationIdentification = (*Max35Text)(&value)
}

func (p *PaymentInstruction11) SetPaymentMethod(value string) {
	p.PaymentMethod = (*PaymentMethod7Code)(&value)
}

func (p *PaymentInstruction11) AddPaymentTypeInformation() *PaymentTypeInformation19 {
	p.PaymentTypeInformation = new(PaymentTypeInformation19)
	return p.PaymentTypeInformation
}

func (p *PaymentInstruction11) SetRequestedExecutionDate(value string) {
	p.RequestedExecutionDate = (*ISODate)(&value)
}

func (p *PaymentInstruction11) AddDebtor() *PartyIdentification43 {
	p.Debtor = new(PartyIdentification43)
	return p.Debtor
}

func (p *PaymentInstruction11) AddDebtorAccount() *CashAccount24 {
	p.DebtorAccount = new(CashAccount24)
	return p.DebtorAccount
}

func (p *PaymentInstruction11) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.DebtorAgent
}

func (p *PaymentInstruction11) AddUltimateDebtor() *PartyIdentification43 {
	p.UltimateDebtor = new(PartyIdentification43)
	return p.UltimateDebtor
}

func (p *PaymentInstruction11) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentInstruction11) AddCreditTransferTransaction() *CreditTransferTransaction10 {
	newValue := new(CreditTransferTransaction10)
	p.CreditTransferTransaction = append(p.CreditTransferTransaction, newValue)
	return newValue
}

// Characteristics that apply to the credit side of the payment transactions included in the direct debit initiation.
type PaymentInstruction15 struct {

	// Unique identification, as assigned by a sending party, to unambiguously identify the payment information group within the message.
	PaymentInformationIdentification *Max35Text `xml:"PmtInfId"`

	// Specifies the means of payment that will be used to move the amount of money.
	PaymentMethod *PaymentMethod2Code `xml:"PmtMtd"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the payment information group.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs,omitempty"`

	// Total of all individual amounts included in the group, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation24 `xml:"PmtTpInf,omitempty"`

	// Date and time at which the creditor requests that the amount of money is to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent in the payment chain.
	CreditorAgentAccount *CashAccount24 `xml:"CdtrAgtAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Account used to process charges associated with a transaction.
	//
	// Usage: Charges account should be used when charges have to be booked to an account different from the account identified in debtor's account.
	ChargesAccount *CashAccount24 `xml:"ChrgsAcct,omitempty"`

	// Agent that services a charges account.
	//
	// Usage: Charges account agent should only be used when the charges account agent is different from the creditor agent.
	ChargesAccountAgent *BranchAndFinancialInstitutionIdentification5 `xml:"ChrgsAcctAgt,omitempty"`

	// Credit party that signs the mandate.
	CreditorSchemeIdentification *PartyIdentification43 `xml:"CdtrSchmeId,omitempty"`

	// Provides information on the individual transaction(s) included in the message.
	DirectDebitTransactionInformation []*DirectDebitTransactionInformation18 `xml:"DrctDbtTxInf"`
}

func (p *PaymentInstruction15) SetPaymentInformationIdentification(value string) {
	p.PaymentInformationIdentification = (*Max35Text)(&value)
}

func (p *PaymentInstruction15) SetPaymentMethod(value string) {
	p.PaymentMethod = (*PaymentMethod2Code)(&value)
}

func (p *PaymentInstruction15) SetBatchBooking(value string) {
	p.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (p *PaymentInstruction15) SetNumberOfTransactions(value string) {
	p.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (p *PaymentInstruction15) SetControlSum(value string) {
	p.ControlSum = (*DecimalNumber)(&value)
}

func (p *PaymentInstruction15) AddPaymentTypeInformation() *PaymentTypeInformation24 {
	p.PaymentTypeInformation = new(PaymentTypeInformation24)
	return p.PaymentTypeInformation
}

func (p *PaymentInstruction15) SetRequestedCollectionDate(value string) {
	p.RequestedCollectionDate = (*ISODate)(&value)
}

func (p *PaymentInstruction15) AddCreditor() *PartyIdentification43 {
	p.Creditor = new(PartyIdentification43)
	return p.Creditor
}

func (p *PaymentInstruction15) AddCreditorAccount() *CashAccount24 {
	p.CreditorAccount = new(CashAccount24)
	return p.CreditorAccount
}

func (p *PaymentInstruction15) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.CreditorAgent
}

func (p *PaymentInstruction15) AddCreditorAgentAccount() *CashAccount24 {
	p.CreditorAgentAccount = new(CashAccount24)
	return p.CreditorAgentAccount
}

func (p *PaymentInstruction15) AddUltimateCreditor() *PartyIdentification43 {
	p.UltimateCreditor = new(PartyIdentification43)
	return p.UltimateCreditor
}

func (p *PaymentInstruction15) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentInstruction15) AddChargesAccount() *CashAccount24 {
	p.ChargesAccount = new(CashAccount24)
	return p.ChargesAccount
}

func (p *PaymentInstruction15) AddChargesAccountAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.ChargesAccountAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.ChargesAccountAgent
}

func (p *PaymentInstruction15) AddCreditorSchemeIdentification() *PartyIdentification43 {
	p.CreditorSchemeIdentification = new(PartyIdentification43)
	return p.CreditorSchemeIdentification
}

func (p *PaymentInstruction15) AddDirectDebitTransactionInformation() *DirectDebitTransactionInformation18 {
	newValue := new(DirectDebitTransactionInformation18)
	p.DirectDebitTransactionInformation = append(p.DirectDebitTransactionInformation, newValue)
	return newValue
}

// Characteristics that apply to the debit side of the payment transactions included in the credit transfer initiation.
type PaymentInstruction16 struct {

	// Unique identification, as assigned by a sending party, to unambiguously identify the payment information group within the message.
	PaymentInformationIdentification *Max35Text `xml:"PmtInfId"`

	// Specifies the means of payment that will be used to move the amount of money.
	PaymentMethod *PaymentMethod3Code `xml:"PmtMtd"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the payment information group.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs,omitempty"`

	// Total of all individual amounts included in the group, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation19 `xml:"PmtTpInf,omitempty"`

	// Date at which the initiating party requests the clearing agent to process the payment.
	// Usage: This is the date on which the debtor's account is to be debited. If payment by cheque, the date when the cheque must be generated by the bank.
	RequestedExecutionDate *ISODate `xml:"ReqdExctnDt"`

	// Date used for the correction of the value date of a cash pool movement that has been posted with a different value date.
	PoolingAdjustmentDate *ISODate `xml:"PoolgAdjstmntDt,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount24 `xml:"DbtrAgtAcct,omitempty"`

	// Further information related to the processing of the payment instruction, that may need to be acted upon by the debtor agent, depending on agreement between debtor and the debtor agent.
	//
	// Usage: when present, then the instructions for the debtor agent apply for all credit transfer transaction information occurrences, present in the payment information.
	InstructionForDebtorAgent *Max140Text `xml:"InstrForDbtrAgt,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Account used to process charges associated with a transaction.
	//
	// Usage: Charges account should be used when charges have to be booked to an account different from the account identified in debtor's account.
	ChargesAccount *CashAccount24 `xml:"ChrgsAcct,omitempty"`

	// Agent that services a charges account.
	//
	// Usage: Charges account agent should only be used when the charges account agent is different from the debtor agent.
	ChargesAccountAgent *BranchAndFinancialInstitutionIdentification5 `xml:"ChrgsAcctAgt,omitempty"`

	// Provides information on the individual transaction(s) included in the message.
	CreditTransferTransactionInformation []*CreditTransferTransaction20 `xml:"CdtTrfTxInf"`
}

func (p *PaymentInstruction16) SetPaymentInformationIdentification(value string) {
	p.PaymentInformationIdentification = (*Max35Text)(&value)
}

func (p *PaymentInstruction16) SetPaymentMethod(value string) {
	p.PaymentMethod = (*PaymentMethod3Code)(&value)
}

func (p *PaymentInstruction16) SetBatchBooking(value string) {
	p.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (p *PaymentInstruction16) SetNumberOfTransactions(value string) {
	p.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (p *PaymentInstruction16) SetControlSum(value string) {
	p.ControlSum = (*DecimalNumber)(&value)
}

func (p *PaymentInstruction16) AddPaymentTypeInformation() *PaymentTypeInformation19 {
	p.PaymentTypeInformation = new(PaymentTypeInformation19)
	return p.PaymentTypeInformation
}

func (p *PaymentInstruction16) SetRequestedExecutionDate(value string) {
	p.RequestedExecutionDate = (*ISODate)(&value)
}

func (p *PaymentInstruction16) SetPoolingAdjustmentDate(value string) {
	p.PoolingAdjustmentDate = (*ISODate)(&value)
}

func (p *PaymentInstruction16) AddDebtor() *PartyIdentification43 {
	p.Debtor = new(PartyIdentification43)
	return p.Debtor
}

func (p *PaymentInstruction16) AddDebtorAccount() *CashAccount24 {
	p.DebtorAccount = new(CashAccount24)
	return p.DebtorAccount
}

func (p *PaymentInstruction16) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.DebtorAgent
}

func (p *PaymentInstruction16) AddDebtorAgentAccount() *CashAccount24 {
	p.DebtorAgentAccount = new(CashAccount24)
	return p.DebtorAgentAccount
}

func (p *PaymentInstruction16) SetInstructionForDebtorAgent(value string) {
	p.InstructionForDebtorAgent = (*Max140Text)(&value)
}

func (p *PaymentInstruction16) AddUltimateDebtor() *PartyIdentification43 {
	p.UltimateDebtor = new(PartyIdentification43)
	return p.UltimateDebtor
}

func (p *PaymentInstruction16) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentInstruction16) AddChargesAccount() *CashAccount24 {
	p.ChargesAccount = new(CashAccount24)
	return p.ChargesAccount
}

func (p *PaymentInstruction16) AddChargesAccountAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.ChargesAccountAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.ChargesAccountAgent
}

func (p *PaymentInstruction16) AddCreditTransferTransactionInformation() *CreditTransferTransaction20 {
	newValue := new(CreditTransferTransaction20)
	p.CreditTransferTransactionInformation = append(p.CreditTransferTransactionInformation, newValue)
	return newValue
}

// Instruction to pay an amount of money to an ultimate beneficiary, on behalf of an originator. This instruction may have to be forwarded several times to complete the settlement chain.
type PaymentInstruction17 struct {

	// Reference assigned by a sending party to unambiguously identify the payment information block within the message.
	PaymentInformationIdentification *Max35Text `xml:"PmtInfId,omitempty"`

	// Specifies the means of payment that will be used to move the amount of money.
	PaymentMethod *PaymentMethod7Code `xml:"PmtMtd"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation19 `xml:"PmtTpInf,omitempty"`

	// Date at which the initiating party requests the clearing agent to process the payment. If payment by cheque, the date when the cheque must be generated by the bank.
	//
	// Usage: This is the date on which the debtor's account(s) is (are) to be debited.
	RequestedExecutionDate *ISODate `xml:"ReqdExctnDt"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr"`

	// Account used to process charges associated with a transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Payment processes required to transfer cash from the debtor to the creditor.
	CreditTransferTransaction []*CreditTransferTransaction21 `xml:"CdtTrfTx"`
}

func (p *PaymentInstruction17) SetPaymentInformationIdentification(value string) {
	p.PaymentInformationIdentification = (*Max35Text)(&value)
}

func (p *PaymentInstruction17) SetPaymentMethod(value string) {
	p.PaymentMethod = (*PaymentMethod7Code)(&value)
}

func (p *PaymentInstruction17) AddPaymentTypeInformation() *PaymentTypeInformation19 {
	p.PaymentTypeInformation = new(PaymentTypeInformation19)
	return p.PaymentTypeInformation
}

func (p *PaymentInstruction17) SetRequestedExecutionDate(value string) {
	p.RequestedExecutionDate = (*ISODate)(&value)
}

func (p *PaymentInstruction17) AddDebtor() *PartyIdentification43 {
	p.Debtor = new(PartyIdentification43)
	return p.Debtor
}

func (p *PaymentInstruction17) AddDebtorAccount() *CashAccount24 {
	p.DebtorAccount = new(CashAccount24)
	return p.DebtorAccount
}

func (p *PaymentInstruction17) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.DebtorAgent
}

func (p *PaymentInstruction17) AddUltimateDebtor() *PartyIdentification43 {
	p.UltimateDebtor = new(PartyIdentification43)
	return p.UltimateDebtor
}

func (p *PaymentInstruction17) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentInstruction17) AddCreditTransferTransaction() *CreditTransferTransaction21 {
	newValue := new(CreditTransferTransaction21)
	p.CreditTransferTransaction = append(p.CreditTransferTransaction, newValue)
	return newValue
}

// Characteristics that apply to the credit side of the payment transactions included in the direct debit initiation.
type PaymentInstruction18 struct {

	// Unique identification, as assigned by a sending party, to unambiguously identify the payment information group within the message.
	PaymentInformationIdentification *Max35Text `xml:"PmtInfId"`

	// Specifies the means of payment that will be used to move the amount of money.
	PaymentMethod *PaymentMethod2Code `xml:"PmtMtd"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the payment information group.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs,omitempty"`

	// Total of all individual amounts included in the group, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation24 `xml:"PmtTpInf,omitempty"`

	// Date and time at which the creditor requests that the amount of money is to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent in the payment chain.
	CreditorAgentAccount *CashAccount24 `xml:"CdtrAgtAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Account used to process charges associated with a transaction.
	//
	// Usage: Charges account should be used when charges have to be booked to an account different from the account identified in debtor's account.
	ChargesAccount *CashAccount24 `xml:"ChrgsAcct,omitempty"`

	// Agent that services a charges account.
	//
	// Usage: Charges account agent should only be used when the charges account agent is different from the creditor agent.
	ChargesAccountAgent *BranchAndFinancialInstitutionIdentification5 `xml:"ChrgsAcctAgt,omitempty"`

	// Credit party that signs the mandate.
	CreditorSchemeIdentification *PartyIdentification43 `xml:"CdtrSchmeId,omitempty"`

	// Provides information on the individual transaction(s) included in the message.
	DirectDebitTransactionInformation []*DirectDebitTransactionInformation19 `xml:"DrctDbtTxInf"`
}

func (p *PaymentInstruction18) SetPaymentInformationIdentification(value string) {
	p.PaymentInformationIdentification = (*Max35Text)(&value)
}

func (p *PaymentInstruction18) SetPaymentMethod(value string) {
	p.PaymentMethod = (*PaymentMethod2Code)(&value)
}

func (p *PaymentInstruction18) SetBatchBooking(value string) {
	p.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (p *PaymentInstruction18) SetNumberOfTransactions(value string) {
	p.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (p *PaymentInstruction18) SetControlSum(value string) {
	p.ControlSum = (*DecimalNumber)(&value)
}

func (p *PaymentInstruction18) AddPaymentTypeInformation() *PaymentTypeInformation24 {
	p.PaymentTypeInformation = new(PaymentTypeInformation24)
	return p.PaymentTypeInformation
}

func (p *PaymentInstruction18) SetRequestedCollectionDate(value string) {
	p.RequestedCollectionDate = (*ISODate)(&value)
}

func (p *PaymentInstruction18) AddCreditor() *PartyIdentification43 {
	p.Creditor = new(PartyIdentification43)
	return p.Creditor
}

func (p *PaymentInstruction18) AddCreditorAccount() *CashAccount24 {
	p.CreditorAccount = new(CashAccount24)
	return p.CreditorAccount
}

func (p *PaymentInstruction18) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.CreditorAgent
}

func (p *PaymentInstruction18) AddCreditorAgentAccount() *CashAccount24 {
	p.CreditorAgentAccount = new(CashAccount24)
	return p.CreditorAgentAccount
}

func (p *PaymentInstruction18) AddUltimateCreditor() *PartyIdentification43 {
	p.UltimateCreditor = new(PartyIdentification43)
	return p.UltimateCreditor
}

func (p *PaymentInstruction18) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentInstruction18) AddChargesAccount() *CashAccount24 {
	p.ChargesAccount = new(CashAccount24)
	return p.ChargesAccount
}

func (p *PaymentInstruction18) AddChargesAccountAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.ChargesAccountAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.ChargesAccountAgent
}

func (p *PaymentInstruction18) AddCreditorSchemeIdentification() *PartyIdentification43 {
	p.CreditorSchemeIdentification = new(PartyIdentification43)
	return p.CreditorSchemeIdentification
}

func (p *PaymentInstruction18) AddDirectDebitTransactionInformation() *DirectDebitTransactionInformation19 {
	newValue := new(DirectDebitTransactionInformation19)
	p.DirectDebitTransactionInformation = append(p.DirectDebitTransactionInformation, newValue)
	return newValue
}

// Instruction to pay an amount of money to an ultimate beneficiary, on behalf of an originator. This instruction may have to be forwarded several times to complete the settlement chain.
type PaymentInstruction19 struct {

	// Reference assigned by a sending party to unambiguously identify the payment information block within the message.
	PaymentInformationIdentification *Max35Text `xml:"PmtInfId,omitempty"`

	// Specifies the means of payment that will be used to move the amount of money.
	PaymentMethod *PaymentMethod7Code `xml:"PmtMtd"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation19 `xml:"PmtTpInf,omitempty"`

	// Date at which the initiating party requests the clearing agent to process the payment. If payment by cheque, the date when the cheque must be generated by the bank.
	//
	// Usage: This is the date on which the debtor's account(s) is (are) to be debited.
	RequestedExecutionDate *ISODate `xml:"ReqdExctnDt"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr"`

	// Account used to process charges associated with a transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Payment processes required to transfer cash from the debtor to the creditor.
	CreditTransferTransaction []*CreditTransferTransaction22 `xml:"CdtTrfTx"`
}

func (p *PaymentInstruction19) SetPaymentInformationIdentification(value string) {
	p.PaymentInformationIdentification = (*Max35Text)(&value)
}

func (p *PaymentInstruction19) SetPaymentMethod(value string) {
	p.PaymentMethod = (*PaymentMethod7Code)(&value)
}

func (p *PaymentInstruction19) AddPaymentTypeInformation() *PaymentTypeInformation19 {
	p.PaymentTypeInformation = new(PaymentTypeInformation19)
	return p.PaymentTypeInformation
}

func (p *PaymentInstruction19) SetRequestedExecutionDate(value string) {
	p.RequestedExecutionDate = (*ISODate)(&value)
}

func (p *PaymentInstruction19) AddDebtor() *PartyIdentification43 {
	p.Debtor = new(PartyIdentification43)
	return p.Debtor
}

func (p *PaymentInstruction19) AddDebtorAccount() *CashAccount24 {
	p.DebtorAccount = new(CashAccount24)
	return p.DebtorAccount
}

func (p *PaymentInstruction19) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.DebtorAgent
}

func (p *PaymentInstruction19) AddUltimateDebtor() *PartyIdentification43 {
	p.UltimateDebtor = new(PartyIdentification43)
	return p.UltimateDebtor
}

func (p *PaymentInstruction19) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentInstruction19) AddCreditTransferTransaction() *CreditTransferTransaction22 {
	newValue := new(CreditTransferTransaction22)
	p.CreditTransferTransaction = append(p.CreditTransferTransaction, newValue)
	return newValue
}

// Characteristics that apply to the debit side of the payment transactions included in the credit transfer initiation.
type PaymentInstruction20 struct {

	// Unique identification, as assigned by a sending party, to unambiguously identify the payment information group within the message.
	PaymentInformationIdentification *Max35Text `xml:"PmtInfId"`

	// Specifies the means of payment that will be used to move the amount of money.
	PaymentMethod *PaymentMethod3Code `xml:"PmtMtd"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the payment information group.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs,omitempty"`

	// Total of all individual amounts included in the group, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation19 `xml:"PmtTpInf,omitempty"`

	// Date at which the initiating party requests the clearing agent to process the payment.
	// Usage: This is the date on which the debtor's account is to be debited. If payment by cheque, the date when the cheque must be generated by the bank.
	RequestedExecutionDate *ISODate `xml:"ReqdExctnDt"`

	// Date used for the correction of the value date of a cash pool movement that has been posted with a different value date.
	PoolingAdjustmentDate *ISODate `xml:"PoolgAdjstmntDt,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount24 `xml:"DbtrAgtAcct,omitempty"`

	// Further information related to the processing of the payment instruction, that may need to be acted upon by the debtor agent, depending on agreement between debtor and the debtor agent.
	//
	// Usage: when present, then the instructions for the debtor agent apply for all credit transfer transaction information occurrences, present in the payment information.
	InstructionForDebtorAgent *Max140Text `xml:"InstrForDbtrAgt,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Account used to process charges associated with a transaction.
	//
	// Usage: Charges account should be used when charges have to be booked to an account different from the account identified in debtor's account.
	ChargesAccount *CashAccount24 `xml:"ChrgsAcct,omitempty"`

	// Agent that services a charges account.
	//
	// Usage: Charges account agent should only be used when the charges account agent is different from the debtor agent.
	ChargesAccountAgent *BranchAndFinancialInstitutionIdentification5 `xml:"ChrgsAcctAgt,omitempty"`

	// Provides information on the individual transaction(s) included in the message.
	CreditTransferTransactionInformation []*CreditTransferTransaction26 `xml:"CdtTrfTxInf"`
}

func (p *PaymentInstruction20) SetPaymentInformationIdentification(value string) {
	p.PaymentInformationIdentification = (*Max35Text)(&value)
}

func (p *PaymentInstruction20) SetPaymentMethod(value string) {
	p.PaymentMethod = (*PaymentMethod3Code)(&value)
}

func (p *PaymentInstruction20) SetBatchBooking(value string) {
	p.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (p *PaymentInstruction20) SetNumberOfTransactions(value string) {
	p.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (p *PaymentInstruction20) SetControlSum(value string) {
	p.ControlSum = (*DecimalNumber)(&value)
}

func (p *PaymentInstruction20) AddPaymentTypeInformation() *PaymentTypeInformation19 {
	p.PaymentTypeInformation = new(PaymentTypeInformation19)
	return p.PaymentTypeInformation
}

func (p *PaymentInstruction20) SetRequestedExecutionDate(value string) {
	p.RequestedExecutionDate = (*ISODate)(&value)
}

func (p *PaymentInstruction20) SetPoolingAdjustmentDate(value string) {
	p.PoolingAdjustmentDate = (*ISODate)(&value)
}

func (p *PaymentInstruction20) AddDebtor() *PartyIdentification43 {
	p.Debtor = new(PartyIdentification43)
	return p.Debtor
}

func (p *PaymentInstruction20) AddDebtorAccount() *CashAccount24 {
	p.DebtorAccount = new(CashAccount24)
	return p.DebtorAccount
}

func (p *PaymentInstruction20) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.DebtorAgent
}

func (p *PaymentInstruction20) AddDebtorAgentAccount() *CashAccount24 {
	p.DebtorAgentAccount = new(CashAccount24)
	return p.DebtorAgentAccount
}

func (p *PaymentInstruction20) SetInstructionForDebtorAgent(value string) {
	p.InstructionForDebtorAgent = (*Max140Text)(&value)
}

func (p *PaymentInstruction20) AddUltimateDebtor() *PartyIdentification43 {
	p.UltimateDebtor = new(PartyIdentification43)
	return p.UltimateDebtor
}

func (p *PaymentInstruction20) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentInstruction20) AddChargesAccount() *CashAccount24 {
	p.ChargesAccount = new(CashAccount24)
	return p.ChargesAccount
}

func (p *PaymentInstruction20) AddChargesAccountAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.ChargesAccountAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.ChargesAccountAgent
}

func (p *PaymentInstruction20) AddCreditTransferTransactionInformation() *CreditTransferTransaction26 {
	newValue := new(CreditTransferTransaction26)
	p.CreditTransferTransactionInformation = append(p.CreditTransferTransactionInformation, newValue)
	return newValue
}

// Characteristics that apply to the credit side of the payment transactions included in the direct debit initiation.
type PaymentInstruction21 struct {

	// Unique identification, as assigned by a sending party, to unambiguously identify the payment information group within the message.
	PaymentInformationIdentification *Max35Text `xml:"PmtInfId"`

	// Specifies the means of payment that will be used to move the amount of money.
	PaymentMethod *PaymentMethod2Code `xml:"PmtMtd"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the payment information group.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs,omitempty"`

	// Total of all individual amounts included in the group, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation24 `xml:"PmtTpInf,omitempty"`

	// Date and time at which the creditor requests that the amount of money is to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent in the payment chain.
	CreditorAgentAccount *CashAccount24 `xml:"CdtrAgtAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Account used to process charges associated with a transaction.
	//
	// Usage: Charges account should be used when charges have to be booked to an account different from the account identified in debtor's account.
	ChargesAccount *CashAccount24 `xml:"ChrgsAcct,omitempty"`

	// Agent that services a charges account.
	//
	// Usage: Charges account agent should only be used when the charges account agent is different from the creditor agent.
	ChargesAccountAgent *BranchAndFinancialInstitutionIdentification5 `xml:"ChrgsAcctAgt,omitempty"`

	// Credit party that signs the mandate.
	CreditorSchemeIdentification *PartyIdentification43 `xml:"CdtrSchmeId,omitempty"`

	// Provides information on the individual transaction(s) included in the message.
	DirectDebitTransactionInformation []*DirectDebitTransactionInformation22 `xml:"DrctDbtTxInf"`
}

func (p *PaymentInstruction21) SetPaymentInformationIdentification(value string) {
	p.PaymentInformationIdentification = (*Max35Text)(&value)
}

func (p *PaymentInstruction21) SetPaymentMethod(value string) {
	p.PaymentMethod = (*PaymentMethod2Code)(&value)
}

func (p *PaymentInstruction21) SetBatchBooking(value string) {
	p.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (p *PaymentInstruction21) SetNumberOfTransactions(value string) {
	p.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (p *PaymentInstruction21) SetControlSum(value string) {
	p.ControlSum = (*DecimalNumber)(&value)
}

func (p *PaymentInstruction21) AddPaymentTypeInformation() *PaymentTypeInformation24 {
	p.PaymentTypeInformation = new(PaymentTypeInformation24)
	return p.PaymentTypeInformation
}

func (p *PaymentInstruction21) SetRequestedCollectionDate(value string) {
	p.RequestedCollectionDate = (*ISODate)(&value)
}

func (p *PaymentInstruction21) AddCreditor() *PartyIdentification43 {
	p.Creditor = new(PartyIdentification43)
	return p.Creditor
}

func (p *PaymentInstruction21) AddCreditorAccount() *CashAccount24 {
	p.CreditorAccount = new(CashAccount24)
	return p.CreditorAccount
}

func (p *PaymentInstruction21) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.CreditorAgent
}

func (p *PaymentInstruction21) AddCreditorAgentAccount() *CashAccount24 {
	p.CreditorAgentAccount = new(CashAccount24)
	return p.CreditorAgentAccount
}

func (p *PaymentInstruction21) AddUltimateCreditor() *PartyIdentification43 {
	p.UltimateCreditor = new(PartyIdentification43)
	return p.UltimateCreditor
}

func (p *PaymentInstruction21) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentInstruction21) AddChargesAccount() *CashAccount24 {
	p.ChargesAccount = new(CashAccount24)
	return p.ChargesAccount
}

func (p *PaymentInstruction21) AddChargesAccountAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.ChargesAccountAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.ChargesAccountAgent
}

func (p *PaymentInstruction21) AddCreditorSchemeIdentification() *PartyIdentification43 {
	p.CreditorSchemeIdentification = new(PartyIdentification43)
	return p.CreditorSchemeIdentification
}

func (p *PaymentInstruction21) AddDirectDebitTransactionInformation() *DirectDebitTransactionInformation22 {
	newValue := new(DirectDebitTransactionInformation22)
	p.DirectDebitTransactionInformation = append(p.DirectDebitTransactionInformation, newValue)
	return newValue
}

// Characteristics that apply to the debit side of the payment transactions included in the credit transfer initiation.
type PaymentInstruction22 struct {

	// Unique identification, as assigned by a sending party, to unambiguously identify the payment information group within the message.
	PaymentInformationIdentification *Max35Text `xml:"PmtInfId"`

	// Specifies the means of payment that will be used to move the amount of money.
	PaymentMethod *PaymentMethod3Code `xml:"PmtMtd"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the payment information group.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs,omitempty"`

	// Total of all individual amounts included in the group, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation19 `xml:"PmtTpInf,omitempty"`

	// Date at which the initiating party requests the clearing agent to process the payment.
	// Usage: This is the date on which the debtor's account is to be debited. If payment by cheque, the date when the cheque must be generated by the bank.
	RequestedExecutionDate *DateAndDateTimeChoice `xml:"ReqdExctnDt"`

	// Date used for the correction of the value date of a cash pool movement that has been posted with a different value date.
	PoolingAdjustmentDate *ISODate `xml:"PoolgAdjstmntDt,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount24 `xml:"DbtrAgtAcct,omitempty"`

	// Further information related to the processing of the payment instruction, that may need to be acted upon by the debtor agent, depending on agreement between debtor and the debtor agent.
	//
	// Usage: when present, then the instructions for the debtor agent apply for all credit transfer transaction information occurrences, present in the payment information.
	InstructionForDebtorAgent *Max140Text `xml:"InstrForDbtrAgt,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Account used to process charges associated with a transaction.
	//
	// Usage: Charges account should be used when charges have to be booked to an account different from the account identified in debtor's account.
	ChargesAccount *CashAccount24 `xml:"ChrgsAcct,omitempty"`

	// Agent that services a charges account.
	//
	// Usage: Charges account agent should only be used when the charges account agent is different from the debtor agent.
	ChargesAccountAgent *BranchAndFinancialInstitutionIdentification5 `xml:"ChrgsAcctAgt,omitempty"`

	// Provides information on the individual transaction(s) included in the message.
	CreditTransferTransactionInformation []*CreditTransferTransaction26 `xml:"CdtTrfTxInf"`
}

func (p *PaymentInstruction22) SetPaymentInformationIdentification(value string) {
	p.PaymentInformationIdentification = (*Max35Text)(&value)
}

func (p *PaymentInstruction22) SetPaymentMethod(value string) {
	p.PaymentMethod = (*PaymentMethod3Code)(&value)
}

func (p *PaymentInstruction22) SetBatchBooking(value string) {
	p.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (p *PaymentInstruction22) SetNumberOfTransactions(value string) {
	p.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (p *PaymentInstruction22) SetControlSum(value string) {
	p.ControlSum = (*DecimalNumber)(&value)
}

func (p *PaymentInstruction22) AddPaymentTypeInformation() *PaymentTypeInformation19 {
	p.PaymentTypeInformation = new(PaymentTypeInformation19)
	return p.PaymentTypeInformation
}

func (p *PaymentInstruction22) AddRequestedExecutionDate() *DateAndDateTimeChoice {
	p.RequestedExecutionDate = new(DateAndDateTimeChoice)
	return p.RequestedExecutionDate
}

func (p *PaymentInstruction22) SetPoolingAdjustmentDate(value string) {
	p.PoolingAdjustmentDate = (*ISODate)(&value)
}

func (p *PaymentInstruction22) AddDebtor() *PartyIdentification43 {
	p.Debtor = new(PartyIdentification43)
	return p.Debtor
}

func (p *PaymentInstruction22) AddDebtorAccount() *CashAccount24 {
	p.DebtorAccount = new(CashAccount24)
	return p.DebtorAccount
}

func (p *PaymentInstruction22) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.DebtorAgent
}

func (p *PaymentInstruction22) AddDebtorAgentAccount() *CashAccount24 {
	p.DebtorAgentAccount = new(CashAccount24)
	return p.DebtorAgentAccount
}

func (p *PaymentInstruction22) SetInstructionForDebtorAgent(value string) {
	p.InstructionForDebtorAgent = (*Max140Text)(&value)
}

func (p *PaymentInstruction22) AddUltimateDebtor() *PartyIdentification43 {
	p.UltimateDebtor = new(PartyIdentification43)
	return p.UltimateDebtor
}

func (p *PaymentInstruction22) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentInstruction22) AddChargesAccount() *CashAccount24 {
	p.ChargesAccount = new(CashAccount24)
	return p.ChargesAccount
}

func (p *PaymentInstruction22) AddChargesAccountAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.ChargesAccountAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.ChargesAccountAgent
}

func (p *PaymentInstruction22) AddCreditTransferTransactionInformation() *CreditTransferTransaction26 {
	newValue := new(CreditTransferTransaction26)
	p.CreditTransferTransactionInformation = append(p.CreditTransferTransactionInformation, newValue)
	return newValue
}

// Instruction to pay an amount of money to an ultimate beneficiary, on behalf of an originator. This instruction may have to be forwarded several times to complete the settlement chain.
type PaymentInstruction23 struct {

	// Reference assigned by a sending party to unambiguously identify the payment information block within the message.
	PaymentInformationIdentification *Max35Text `xml:"PmtInfId,omitempty"`

	// Specifies the means of payment that will be used to move the amount of money.
	PaymentMethod *PaymentMethod7Code `xml:"PmtMtd"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation19 `xml:"PmtTpInf,omitempty"`

	// Date at which the initiating party requests the clearing agent to process the payment. If payment by cheque, the date when the cheque must be generated by the bank.
	//
	// Usage: This is the date on which the debtor's account(s) is (are) to be debited.
	RequestedExecutionDate *DateAndDateTimeChoice `xml:"ReqdExctnDt"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr"`

	// Account used to process charges associated with a transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Payment processes required to transfer cash from the debtor to the creditor.
	CreditTransferTransaction []*CreditTransferTransaction22 `xml:"CdtTrfTx"`
}

func (p *PaymentInstruction23) SetPaymentInformationIdentification(value string) {
	p.PaymentInformationIdentification = (*Max35Text)(&value)
}

func (p *PaymentInstruction23) SetPaymentMethod(value string) {
	p.PaymentMethod = (*PaymentMethod7Code)(&value)
}

func (p *PaymentInstruction23) AddPaymentTypeInformation() *PaymentTypeInformation19 {
	p.PaymentTypeInformation = new(PaymentTypeInformation19)
	return p.PaymentTypeInformation
}

func (p *PaymentInstruction23) AddRequestedExecutionDate() *DateAndDateTimeChoice {
	p.RequestedExecutionDate = new(DateAndDateTimeChoice)
	return p.RequestedExecutionDate
}

func (p *PaymentInstruction23) AddDebtor() *PartyIdentification43 {
	p.Debtor = new(PartyIdentification43)
	return p.Debtor
}

func (p *PaymentInstruction23) AddDebtorAccount() *CashAccount24 {
	p.DebtorAccount = new(CashAccount24)
	return p.DebtorAccount
}

func (p *PaymentInstruction23) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.DebtorAgent
}

func (p *PaymentInstruction23) AddUltimateDebtor() *PartyIdentification43 {
	p.UltimateDebtor = new(PartyIdentification43)
	return p.UltimateDebtor
}

func (p *PaymentInstruction23) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentInstruction23) AddCreditTransferTransaction() *CreditTransferTransaction22 {
	newValue := new(CreditTransferTransaction22)
	p.CreditTransferTransaction = append(p.CreditTransferTransaction, newValue)
	return newValue
}

// Instruction to pay an amount of money to an ultimate beneficiary, on behalf of an originator. This instruction may have to be forwarded several times to complete the settlement chain.
type PaymentInstruction5 struct {

	// Reference assigned by a sending party to unambiguously identify the payment information block within the message.
	PaymentInformationIdentification *Max35Text `xml:"PmtInfId,omitempty"`

	// Specifies the means of payment that will be used to move the amount of money.
	PaymentMethod *PaymentMethod7Code `xml:"PmtMtd"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation19 `xml:"PmtTpInf,omitempty"`

	// Date at which the initiating party requests the clearing agent to process the payment. If payment by cheque, the date when the cheque must be generated by the bank.
	//
	// Usage: This is the date on which the debtor's account(s) is (are) to be debited.
	RequestedExecutionDate *ISODate `xml:"ReqdExctnDt"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr"`

	// Account used to process charges associated with a transaction.
	DebtorAccount *CashAccount7 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Payment processes required to transfer cash from the debtor to the creditor.
	CreditTransferTransaction []*CreditTransferTransactionInformation14 `xml:"CdtTrfTx"`
}

func (p *PaymentInstruction5) SetPaymentInformationIdentification(value string) {
	p.PaymentInformationIdentification = (*Max35Text)(&value)
}

func (p *PaymentInstruction5) SetPaymentMethod(value string) {
	p.PaymentMethod = (*PaymentMethod7Code)(&value)
}

func (p *PaymentInstruction5) AddPaymentTypeInformation() *PaymentTypeInformation19 {
	p.PaymentTypeInformation = new(PaymentTypeInformation19)
	return p.PaymentTypeInformation
}

func (p *PaymentInstruction5) SetRequestedExecutionDate(value string) {
	p.RequestedExecutionDate = (*ISODate)(&value)
}

func (p *PaymentInstruction5) AddDebtor() *PartyIdentification43 {
	p.Debtor = new(PartyIdentification43)
	return p.Debtor
}

func (p *PaymentInstruction5) AddDebtorAccount() *CashAccount7 {
	p.DebtorAccount = new(CashAccount7)
	return p.DebtorAccount
}

func (p *PaymentInstruction5) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.DebtorAgent
}

func (p *PaymentInstruction5) AddUltimateDebtor() *PartyIdentification43 {
	p.UltimateDebtor = new(PartyIdentification43)
	return p.UltimateDebtor
}

func (p *PaymentInstruction5) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentInstruction5) AddCreditTransferTransaction() *CreditTransferTransactionInformation14 {
	newValue := new(CreditTransferTransactionInformation14)
	p.CreditTransferTransaction = append(p.CreditTransferTransaction, newValue)
	return newValue
}

// Characteristics that apply to the debit side of the payment transactions included in the credit transfer initiation.
type PaymentInstruction6 struct {

	// Unique identification, as assigned by a sending party, to unambiguously identify the payment information group within the message.
	PaymentInformationIdentification *Max35Text `xml:"PmtInfId"`

	// Specifies the means of payment that will be used to move the amount of money.
	PaymentMethod *PaymentMethod3Code `xml:"PmtMtd"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the payment information group.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs,omitempty"`

	// Total of all individual amounts included in the group, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation19 `xml:"PmtTpInf,omitempty"`

	// Date at which the initiating party requests the clearing agent to process the payment.
	// Usage: This is the date on which the debtor's account is to be debited. If payment by cheque, the date when the cheque must be generated by the bank.
	RequestedExecutionDate *ISODate `xml:"ReqdExctnDt"`

	// Date used for the correction of the value date of a cash pool movement that has been posted with a different value date.
	PoolingAdjustmentDate *ISODate `xml:"PoolgAdjstmntDt,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount24 `xml:"DbtrAgtAcct,omitempty"`

	// Further information related to the processing of the payment instruction, that may need to be acted upon by the debtor agent, depending on agreement between debtor and the debtor agent.
	//
	// Usage: when present, then the instructions for the debtor agent apply for all credit transfer transaction information occurrences, present in the payment information.
	InstructionForDebtorAgent *Max140Text `xml:"InstrForDbtrAgt,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Account used to process charges associated with a transaction.
	//
	// Usage: Charges account should be used when charges have to be booked to an account different from the account identified in debtor's account.
	ChargesAccount *CashAccount24 `xml:"ChrgsAcct,omitempty"`

	// Agent that services a charges account.
	//
	// Usage: Charges account agent should only be used when the charges account agent is different from the debtor agent.
	ChargesAccountAgent *BranchAndFinancialInstitutionIdentification5 `xml:"ChrgsAcctAgt,omitempty"`

	// Provides information on the individual transaction(s) included in the message.
	CreditTransferTransactionInformation []*CreditTransferTransaction1 `xml:"CdtTrfTxInf"`
}

func (p *PaymentInstruction6) SetPaymentInformationIdentification(value string) {
	p.PaymentInformationIdentification = (*Max35Text)(&value)
}

func (p *PaymentInstruction6) SetPaymentMethod(value string) {
	p.PaymentMethod = (*PaymentMethod3Code)(&value)
}

func (p *PaymentInstruction6) SetBatchBooking(value string) {
	p.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (p *PaymentInstruction6) SetNumberOfTransactions(value string) {
	p.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (p *PaymentInstruction6) SetControlSum(value string) {
	p.ControlSum = (*DecimalNumber)(&value)
}

func (p *PaymentInstruction6) AddPaymentTypeInformation() *PaymentTypeInformation19 {
	p.PaymentTypeInformation = new(PaymentTypeInformation19)
	return p.PaymentTypeInformation
}

func (p *PaymentInstruction6) SetRequestedExecutionDate(value string) {
	p.RequestedExecutionDate = (*ISODate)(&value)
}

func (p *PaymentInstruction6) SetPoolingAdjustmentDate(value string) {
	p.PoolingAdjustmentDate = (*ISODate)(&value)
}

func (p *PaymentInstruction6) AddDebtor() *PartyIdentification43 {
	p.Debtor = new(PartyIdentification43)
	return p.Debtor
}

func (p *PaymentInstruction6) AddDebtorAccount() *CashAccount24 {
	p.DebtorAccount = new(CashAccount24)
	return p.DebtorAccount
}

func (p *PaymentInstruction6) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.DebtorAgent
}

func (p *PaymentInstruction6) AddDebtorAgentAccount() *CashAccount24 {
	p.DebtorAgentAccount = new(CashAccount24)
	return p.DebtorAgentAccount
}

func (p *PaymentInstruction6) SetInstructionForDebtorAgent(value string) {
	p.InstructionForDebtorAgent = (*Max140Text)(&value)
}

func (p *PaymentInstruction6) AddUltimateDebtor() *PartyIdentification43 {
	p.UltimateDebtor = new(PartyIdentification43)
	return p.UltimateDebtor
}

func (p *PaymentInstruction6) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentInstruction6) AddChargesAccount() *CashAccount24 {
	p.ChargesAccount = new(CashAccount24)
	return p.ChargesAccount
}

func (p *PaymentInstruction6) AddChargesAccountAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.ChargesAccountAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.ChargesAccountAgent
}

func (p *PaymentInstruction6) AddCreditTransferTransactionInformation() *CreditTransferTransaction1 {
	newValue := new(CreditTransferTransaction1)
	p.CreditTransferTransactionInformation = append(p.CreditTransferTransactionInformation, newValue)
	return newValue
}

// Characteristics that apply to the credit side of the payment transactions included in the direct debit initiation.
type PaymentInstruction7 struct {

	// Unique identification, as assigned by a sending party, to unambiguously identify the payment information group within the message.
	PaymentInformationIdentification *Max35Text `xml:"PmtInfId"`

	// Specifies the means of payment that will be used to move the amount of money.
	PaymentMethod *PaymentMethod2Code `xml:"PmtMtd"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the payment information group.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs,omitempty"`

	// Total of all individual amounts included in the group, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation24 `xml:"PmtTpInf,omitempty"`

	// Date and time at which the creditor requests that the amount of money is to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent in the payment chain.
	CreditorAgentAccount *CashAccount24 `xml:"CdtrAgtAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Account used to process charges associated with a transaction.
	//
	// Usage: Charges account should be used when charges have to be booked to an account different from the account identified in debtor's account.
	ChargesAccount *CashAccount24 `xml:"ChrgsAcct,omitempty"`

	// Agent that services a charges account.
	//
	// Usage: Charges account agent should only be used when the charges account agent is different from the creditor agent.
	ChargesAccountAgent *BranchAndFinancialInstitutionIdentification5 `xml:"ChrgsAcctAgt,omitempty"`

	// Credit party that signs the mandate.
	CreditorSchemeIdentification *PartyIdentification43 `xml:"CdtrSchmeId,omitempty"`

	// Provides information on the individual transaction(s) included in the message.
	DirectDebitTransactionInformation []*DirectDebitTransactionInformation11 `xml:"DrctDbtTxInf"`
}

func (p *PaymentInstruction7) SetPaymentInformationIdentification(value string) {
	p.PaymentInformationIdentification = (*Max35Text)(&value)
}

func (p *PaymentInstruction7) SetPaymentMethod(value string) {
	p.PaymentMethod = (*PaymentMethod2Code)(&value)
}

func (p *PaymentInstruction7) SetBatchBooking(value string) {
	p.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (p *PaymentInstruction7) SetNumberOfTransactions(value string) {
	p.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (p *PaymentInstruction7) SetControlSum(value string) {
	p.ControlSum = (*DecimalNumber)(&value)
}

func (p *PaymentInstruction7) AddPaymentTypeInformation() *PaymentTypeInformation24 {
	p.PaymentTypeInformation = new(PaymentTypeInformation24)
	return p.PaymentTypeInformation
}

func (p *PaymentInstruction7) SetRequestedCollectionDate(value string) {
	p.RequestedCollectionDate = (*ISODate)(&value)
}

func (p *PaymentInstruction7) AddCreditor() *PartyIdentification43 {
	p.Creditor = new(PartyIdentification43)
	return p.Creditor
}

func (p *PaymentInstruction7) AddCreditorAccount() *CashAccount24 {
	p.CreditorAccount = new(CashAccount24)
	return p.CreditorAccount
}

func (p *PaymentInstruction7) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.CreditorAgent
}

func (p *PaymentInstruction7) AddCreditorAgentAccount() *CashAccount24 {
	p.CreditorAgentAccount = new(CashAccount24)
	return p.CreditorAgentAccount
}

func (p *PaymentInstruction7) AddUltimateCreditor() *PartyIdentification43 {
	p.UltimateCreditor = new(PartyIdentification43)
	return p.UltimateCreditor
}

func (p *PaymentInstruction7) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentInstruction7) AddChargesAccount() *CashAccount24 {
	p.ChargesAccount = new(CashAccount24)
	return p.ChargesAccount
}

func (p *PaymentInstruction7) AddChargesAccountAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.ChargesAccountAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.ChargesAccountAgent
}

func (p *PaymentInstruction7) AddCreditorSchemeIdentification() *PartyIdentification43 {
	p.CreditorSchemeIdentification = new(PartyIdentification43)
	return p.CreditorSchemeIdentification
}

func (p *PaymentInstruction7) AddDirectDebitTransactionInformation() *DirectDebitTransactionInformation11 {
	newValue := new(DirectDebitTransactionInformation11)
	p.DirectDebitTransactionInformation = append(p.DirectDebitTransactionInformation, newValue)
	return newValue
}

// Instruction to pay an amount of money to an ultimate beneficiary, on behalf of an originator. This instruction may have to be forwarded several times to complete the settlement chain.
type PaymentInstruction8 struct {

	// Reference assigned by a sending party to unambiguously identify the payment information block within the message.
	PaymentInformationIdentification *Max35Text `xml:"PmtInfId,omitempty"`

	// Specifies the means of payment that will be used to move the amount of money.
	PaymentMethod *PaymentMethod7Code `xml:"PmtMtd"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation19 `xml:"PmtTpInf,omitempty"`

	// Date at which the initiating party requests the clearing agent to process the payment. If payment by cheque, the date when the cheque must be generated by the bank.
	//
	// Usage: This is the date on which the debtor's account(s) is (are) to be debited.
	RequestedExecutionDate *ISODate `xml:"ReqdExctnDt"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr"`

	// Account used to process charges associated with a transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Payment processes required to transfer cash from the debtor to the creditor.
	CreditTransferTransaction []*CreditTransferTransaction5 `xml:"CdtTrfTx"`
}

func (p *PaymentInstruction8) SetPaymentInformationIdentification(value string) {
	p.PaymentInformationIdentification = (*Max35Text)(&value)
}

func (p *PaymentInstruction8) SetPaymentMethod(value string) {
	p.PaymentMethod = (*PaymentMethod7Code)(&value)
}

func (p *PaymentInstruction8) AddPaymentTypeInformation() *PaymentTypeInformation19 {
	p.PaymentTypeInformation = new(PaymentTypeInformation19)
	return p.PaymentTypeInformation
}

func (p *PaymentInstruction8) SetRequestedExecutionDate(value string) {
	p.RequestedExecutionDate = (*ISODate)(&value)
}

func (p *PaymentInstruction8) AddDebtor() *PartyIdentification43 {
	p.Debtor = new(PartyIdentification43)
	return p.Debtor
}

func (p *PaymentInstruction8) AddDebtorAccount() *CashAccount24 {
	p.DebtorAccount = new(CashAccount24)
	return p.DebtorAccount
}

func (p *PaymentInstruction8) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.DebtorAgent
}

func (p *PaymentInstruction8) AddUltimateDebtor() *PartyIdentification43 {
	p.UltimateDebtor = new(PartyIdentification43)
	return p.UltimateDebtor
}

func (p *PaymentInstruction8) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentInstruction8) AddCreditTransferTransaction() *CreditTransferTransaction5 {
	newValue := new(CreditTransferTransaction5)
	p.CreditTransferTransaction = append(p.CreditTransferTransaction, newValue)
	return newValue
}

// Characteristics that apply to the debit side of the payment transactions included in the credit transfer initiation.
type PaymentInstruction9 struct {

	// Unique identification, as assigned by a sending party, to unambiguously identify the payment information group within the message.
	PaymentInformationIdentification *Max35Text `xml:"PmtInfId"`

	// Specifies the means of payment that will be used to move the amount of money.
	PaymentMethod *PaymentMethod3Code `xml:"PmtMtd"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the payment information group.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs,omitempty"`

	// Total of all individual amounts included in the group, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation19 `xml:"PmtTpInf,omitempty"`

	// Date at which the initiating party requests the clearing agent to process the payment.
	// Usage: This is the date on which the debtor's account is to be debited. If payment by cheque, the date when the cheque must be generated by the bank.
	RequestedExecutionDate *ISODate `xml:"ReqdExctnDt"`

	// Date used for the correction of the value date of a cash pool movement that has been posted with a different value date.
	PoolingAdjustmentDate *ISODate `xml:"PoolgAdjstmntDt,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount24 `xml:"DbtrAgtAcct,omitempty"`

	// Further information related to the processing of the payment instruction, that may need to be acted upon by the debtor agent, depending on agreement between debtor and the debtor agent.
	//
	// Usage: when present, then the instructions for the debtor agent apply for all credit transfer transaction information occurrences, present in the payment information.
	InstructionForDebtorAgent *Max140Text `xml:"InstrForDbtrAgt,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Account used to process charges associated with a transaction.
	//
	// Usage: Charges account should be used when charges have to be booked to an account different from the account identified in debtor's account.
	ChargesAccount *CashAccount24 `xml:"ChrgsAcct,omitempty"`

	// Agent that services a charges account.
	//
	// Usage: Charges account agent should only be used when the charges account agent is different from the debtor agent.
	ChargesAccountAgent *BranchAndFinancialInstitutionIdentification5 `xml:"ChrgsAcctAgt,omitempty"`

	// Provides information on the individual transaction(s) included in the message.
	CreditTransferTransactionInformation []*CreditTransferTransaction6 `xml:"CdtTrfTxInf"`
}

func (p *PaymentInstruction9) SetPaymentInformationIdentification(value string) {
	p.PaymentInformationIdentification = (*Max35Text)(&value)
}

func (p *PaymentInstruction9) SetPaymentMethod(value string) {
	p.PaymentMethod = (*PaymentMethod3Code)(&value)
}

func (p *PaymentInstruction9) SetBatchBooking(value string) {
	p.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (p *PaymentInstruction9) SetNumberOfTransactions(value string) {
	p.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (p *PaymentInstruction9) SetControlSum(value string) {
	p.ControlSum = (*DecimalNumber)(&value)
}

func (p *PaymentInstruction9) AddPaymentTypeInformation() *PaymentTypeInformation19 {
	p.PaymentTypeInformation = new(PaymentTypeInformation19)
	return p.PaymentTypeInformation
}

func (p *PaymentInstruction9) SetRequestedExecutionDate(value string) {
	p.RequestedExecutionDate = (*ISODate)(&value)
}

func (p *PaymentInstruction9) SetPoolingAdjustmentDate(value string) {
	p.PoolingAdjustmentDate = (*ISODate)(&value)
}

func (p *PaymentInstruction9) AddDebtor() *PartyIdentification43 {
	p.Debtor = new(PartyIdentification43)
	return p.Debtor
}

func (p *PaymentInstruction9) AddDebtorAccount() *CashAccount24 {
	p.DebtorAccount = new(CashAccount24)
	return p.DebtorAccount
}

func (p *PaymentInstruction9) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.DebtorAgent
}

func (p *PaymentInstruction9) AddDebtorAgentAccount() *CashAccount24 {
	p.DebtorAgentAccount = new(CashAccount24)
	return p.DebtorAgentAccount
}

func (p *PaymentInstruction9) SetInstructionForDebtorAgent(value string) {
	p.InstructionForDebtorAgent = (*Max140Text)(&value)
}

func (p *PaymentInstruction9) AddUltimateDebtor() *PartyIdentification43 {
	p.UltimateDebtor = new(PartyIdentification43)
	return p.UltimateDebtor
}

func (p *PaymentInstruction9) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentInstruction9) AddChargesAccount() *CashAccount24 {
	p.ChargesAccount = new(CashAccount24)
	return p.ChargesAccount
}

func (p *PaymentInstruction9) AddChargesAccountAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.ChargesAccountAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.ChargesAccountAgent
}

func (p *PaymentInstruction9) AddCreditTransferTransactionInformation() *CreditTransferTransaction6 {
	newValue := new(CreditTransferTransaction6)
	p.CreditTransferTransactionInformation = append(p.CreditTransferTransactionInformation, newValue)
	return newValue
}
