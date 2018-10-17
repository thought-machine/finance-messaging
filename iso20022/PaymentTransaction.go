package iso20022

// Payment processes required to transfer cash from the debtor to the creditor.
type PaymentTransaction13 struct {

	// Amount of money to be transferred between the debtor and creditor before bank transaction charges.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Date on which the first agent expects the cash to be available to the final agent.
	SettlementDate *ISODate `xml:"SttlmDt,omitempty"`

	// Choice between types of payment instrument, ie, cheque, credit transfer, direct debit, investment account or payment card.
	PaymentInstrument *PaymentInstrument6Choice `xml:"PmtInstrm,omitempty"`
}

func (p *PaymentTransaction13) SetSettlementAmount(value, currency string) {
	p.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction13) SetSettlementDate(value string) {
	p.SettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction13) AddPaymentInstrument() *PaymentInstrument6Choice {
	p.PaymentInstrument = new(PaymentInstrument6Choice)
	return p.PaymentInstrument
}

// Payment processes required to transfer cash from the debtor to the creditor.
type PaymentTransaction14 struct {

	// Amount of money to be transferred between the debtor and creditor before bank transaction charges.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Date on which the first agent expects the cash to be available to the final agent.
	SettlementDate *ISODate `xml:"SttlmDt,omitempty"`

	// Choice between cash-in or cash-out.
	CashInOrOutChoice *CashInOrOut3Choice `xml:"CshInOrOutChc,omitempty"`
}

func (p *PaymentTransaction14) SetSettlementAmount(value, currency string) {
	p.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction14) SetSettlementDate(value string) {
	p.SettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction14) AddCashInOrOutChoice() *CashInOrOut3Choice {
	p.CashInOrOutChoice = new(CashInOrOut3Choice)
	return p.CashInOrOutChoice
}

// Payment processes required to transfer cash from the debtor to the creditor.
type PaymentTransaction15 struct {

	// Amount of money to be transferred between the debtor and creditor before bank transaction charges.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Date on which the first agent expects the cash to be available to the final agent.
	SettlementDate *ISODate `xml:"SttlmDt,omitempty"`

	// Choice between types of payment instrument, ie, cheque, credit transfer or investment account.
	PaymentInstrument *PaymentInstrument7Choice `xml:"PmtInstrm,omitempty"`
}

func (p *PaymentTransaction15) SetSettlementAmount(value, currency string) {
	p.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction15) SetSettlementDate(value string) {
	p.SettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction15) AddPaymentInstrument() *PaymentInstrument7Choice {
	p.PaymentInstrument = new(PaymentInstrument7Choice)
	return p.PaymentInstrument
}

// Payment processes required to transfer cash from the debtor to the creditor.
type PaymentTransaction16 struct {

	// Amount of money to be transferred between the debtor and creditor before bank transaction charges.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Date on which the first agent expects the cash to be available to the final agent.
	SettlementDate *ISODate `xml:"SttlmDt,omitempty"`

	// Choice between types of payment instrument, ie, cheque, credit transfer, direct debit, investment account or payment card.
	PaymentInstrument *PaymentInstrument8Choice `xml:"PmtInstrm,omitempty"`
}

func (p *PaymentTransaction16) SetSettlementAmount(value, currency string) {
	p.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction16) SetSettlementDate(value string) {
	p.SettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction16) AddPaymentInstrument() *PaymentInstrument8Choice {
	p.PaymentInstrument = new(PaymentInstrument8Choice)
	return p.PaymentInstrument
}

// Payment processes required to transfer cash from the debtor to the creditor.
type PaymentTransaction17 struct {

	// Amount of money to be transferred between the debtor and creditor before bank transaction charges.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Date on which the first agent expects the cash to be available to the final agent.
	SettlementDate *ISODate `xml:"SttlmDt,omitempty"`

	// Choice between types of payment instrument, ie, cheque, credit transfer, direct debit, investment account or payment card.
	PaymentInstrument *PaymentInstrument9Choice `xml:"PmtInstrm,omitempty"`
}

func (p *PaymentTransaction17) SetSettlementAmount(value, currency string) {
	p.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction17) SetSettlementDate(value string) {
	p.SettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction17) AddPaymentInstrument() *PaymentInstrument9Choice {
	p.PaymentInstrument = new(PaymentInstrument9Choice)
	return p.PaymentInstrument
}

// Payment processes required to transfer cash from the debtor to the creditor.
type PaymentTransaction18 struct {

	// Amount of money to be transferred between the debtor and creditor before bank transaction charges.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Date on which the first agent expects the cash to be available to the final agent.
	SettlementDate *ISODate `xml:"SttlmDt,omitempty"`

	// Choice between types of payment instrument, ie, cheque, credit transfer or investment account.
	PaymentInstrument *PaymentInstrument7Choice `xml:"PmtInstrm"`
}

func (p *PaymentTransaction18) SetSettlementAmount(value, currency string) {
	p.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction18) SetSettlementDate(value string) {
	p.SettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction18) AddPaymentInstrument() *PaymentInstrument7Choice {
	p.PaymentInstrument = new(PaymentInstrument7Choice)
	return p.PaymentInstrument
}

// Payment processes required to transfer cash from the debtor to the creditor.
type PaymentTransaction19 struct {

	// Amount of money to be transferred between the debtor and creditor before bank transaction charges.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Date on which the first agent expects the cash to be available to the final agent.
	SettlementDate *ISODate `xml:"SttlmDt,omitempty"`

	// Choice between types of payment instrument, ie, credit transfer, cheque, payment card, investment cash account or direct debit.
	PaymentInstrument *PaymentInstrument10Choice `xml:"PmtInstrm,omitempty"`
}

func (p *PaymentTransaction19) SetSettlementAmount(value, currency string) {
	p.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction19) SetSettlementDate(value string) {
	p.SettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction19) AddPaymentInstrument() *PaymentInstrument10Choice {
	p.PaymentInstrument = new(PaymentInstrument10Choice)
	return p.PaymentInstrument
}

// Payment processes required to transfer cash from the debtor to the creditor.
type PaymentTransaction20 struct {

	// Amount of money to be transferred between the debtor and creditor before bank transaction charges.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Date on which the first agent expects the cash to be available to the final agent.
	SettlementDate *ISODate `xml:"SttlmDt,omitempty"`

	// Choice between cash-in or cash-out.
	CashInOrOut *CashInOrOut4Choice `xml:"CshInOrOut"`
}

func (p *PaymentTransaction20) SetSettlementAmount(value, currency string) {
	p.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction20) SetSettlementDate(value string) {
	p.SettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction20) AddCashInOrOut() *CashInOrOut4Choice {
	p.CashInOrOut = new(CashInOrOut4Choice)
	return p.CashInOrOut
}

// Payment processes required to transfer cash from the debtor to the creditor.
type PaymentTransaction21 struct {

	// Choice between types of payment instrument, ie, cheque, credit transfer or investment account.
	PaymentInstrument *PaymentInstrument11Choice `xml:"PmtInstrm,omitempty"`
}

func (p *PaymentTransaction21) AddPaymentInstrument() *PaymentInstrument11Choice {
	p.PaymentInstrument = new(PaymentInstrument11Choice)
	return p.PaymentInstrument
}

// Payment processes required to transfer cash from the debtor to the creditor.
type PaymentTransaction22 struct {

	// Choice between types of payment instrument, ie, cheque, credit transfer or investment account.
	PaymentInstrument *PaymentInstrument11Choice `xml:"PmtInstrm,omitempty"`
}

func (p *PaymentTransaction22) AddPaymentInstrument() *PaymentInstrument11Choice {
	p.PaymentInstrument = new(PaymentInstrument11Choice)
	return p.PaymentInstrument
}

// Payment processes required to transfer cash from the debtor to the creditor.
type PaymentTransaction23 struct {

	// Choice between types of payment instrument, ie, cheque, credit transfer, direct debit, investment account or payment card.
	PaymentInstrument *PaymentInstrument12Choice `xml:"PmtInstrm,omitempty"`
}

func (p *PaymentTransaction23) AddPaymentInstrument() *PaymentInstrument12Choice {
	p.PaymentInstrument = new(PaymentInstrument12Choice)
	return p.PaymentInstrument
}

// Payment processes required to transfer cash from the debtor to the creditor.
type PaymentTransaction24 struct {

	// Choice between types of payment instrument, ie, cheque, credit transfer, direct debit, investment account or payment card.
	PaymentInstrument *PaymentInstrument12Choice `xml:"PmtInstrm,omitempty"`
}

func (p *PaymentTransaction24) AddPaymentInstrument() *PaymentInstrument12Choice {
	p.PaymentInstrument = new(PaymentInstrument12Choice)
	return p.PaymentInstrument
}

// Payment processes required to transfer cash from the debtor to the creditor.
type PaymentTransaction25 struct {

	// Choice between cash-in or cash-out.
	CashInOrOut *CashInOrOut5Choice `xml:"CshInOrOut,omitempty"`
}

func (p *PaymentTransaction25) AddCashInOrOut() *CashInOrOut5Choice {
	p.CashInOrOut = new(CashInOrOut5Choice)
	return p.CashInOrOut
}

// Payment processes required to transfer cash from the debtor to the creditor.
type PaymentTransaction26 struct {

	// Choice between cash-in or cash-out.
	CashInOrOut *CashInOrOut6Choice `xml:"CshInOrOut,omitempty"`
}

func (p *PaymentTransaction26) AddCashInOrOut() *CashInOrOut6Choice {
	p.CashInOrOut = new(CashInOrOut6Choice)
	return p.CashInOrOut
}

// Provides further details on the original transactions, to which the status report message refers.
type PaymentTransaction32 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reported status.
	// Usage: The instructing party is the party sending the status message and not the party that sent the original instruction that is being reported on.
	StatusIdentification *Max35Text `xml:"StsId,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Specifies the status of a transaction, in a coded form.
	TransactionStatus *TransactionIndividualStatus3Code `xml:"TxSts,omitempty"`

	// Provides detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Provides information on the charges related to the processing of the rejection of the instruction.
	// Usage: This is passed on for information purposes only. Settlement of the charges will be done separately.
	ChargesInformation []*Charges2 `xml:"ChrgsInf,omitempty"`

	// Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.
	AcceptanceDateTime *ISODateTime `xml:"AccptncDtTm,omitempty"`

	// Unique reference, as assigned by the account servicing institution, to unambiguously identify the instruction.
	AccountServicerReference *Max35Text `xml:"AcctSvcrRef,omitempty"`

	// Unique reference, as assigned by a clearing system, to unambiguously identify the instruction.
	ClearingSystemReference *Max35Text `xml:"ClrSysRef,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference16 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransaction32) SetStatusIdentification(value string) {
	p.StatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction32) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction32) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction32) SetTransactionStatus(value string) {
	p.TransactionStatus = (*TransactionIndividualStatus3Code)(&value)
}

func (p *PaymentTransaction32) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	p.StatusReasonInformation = append(p.StatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction32) AddChargesInformation() *Charges2 {
	newValue := new(Charges2)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransaction32) SetAcceptanceDateTime(value string) {
	p.AcceptanceDateTime = (*ISODateTime)(&value)
}

func (p *PaymentTransaction32) SetAccountServicerReference(value string) {
	p.AccountServicerReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction32) SetClearingSystemReference(value string) {
	p.ClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction32) AddOriginalTransactionReference() *OriginalTransactionReference16 {
	p.OriginalTransactionReference = new(OriginalTransactionReference16)
	return p.OriginalTransactionReference
}

// Provides further details on the original transactions, to which the status report message refers.
type PaymentTransaction33 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reported status.
	// Usage: The instructing party is the party sending the status message and not the party that sent the original instruction that is being reported on.
	StatusIdentification *Max35Text `xml:"StsId,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Specifies the status of a transaction, in a coded form.
	TransactionStatus *TransactionIndividualStatus3Code `xml:"TxSts,omitempty"`

	// Provides detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Provides information on the charges related to the processing of the rejection of the instruction.
	//
	// Usage: This is passed on for information purposes only. Settlement of the charges will be done separately.
	ChargesInformation []*Charges2 `xml:"ChrgsInf,omitempty"`

	// Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.
	AcceptanceDateTime *ISODateTime `xml:"AccptncDtTm,omitempty"`

	// Unique reference, as assigned by the account servicing institution, to unambiguously identify the instruction.
	AccountServicerReference *Max35Text `xml:"AcctSvcrRef,omitempty"`

	// Unique reference, as assigned by a clearing system, to unambiguously identify the instruction.
	ClearingSystemReference *Max35Text `xml:"ClrSysRef,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructing agent is the party sending the status message and not the party that sent the original instruction that is being reported on.
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructed agent is the party receiving the status message and not the party that received the original instruction that is being reported on.
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference16 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransaction33) SetStatusIdentification(value string) {
	p.StatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction33) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction33) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction33) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction33) SetTransactionStatus(value string) {
	p.TransactionStatus = (*TransactionIndividualStatus3Code)(&value)
}

func (p *PaymentTransaction33) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	p.StatusReasonInformation = append(p.StatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction33) AddChargesInformation() *Charges2 {
	newValue := new(Charges2)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransaction33) SetAcceptanceDateTime(value string) {
	p.AcceptanceDateTime = (*ISODateTime)(&value)
}

func (p *PaymentTransaction33) SetAccountServicerReference(value string) {
	p.AccountServicerReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction33) SetClearingSystemReference(value string) {
	p.ClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction33) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructingAgent
}

func (p *PaymentTransaction33) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructedAgent
}

func (p *PaymentTransaction33) AddOriginalTransactionReference() *OriginalTransactionReference16 {
	p.OriginalTransactionReference = new(OriginalTransactionReference16)
	return p.OriginalTransactionReference
}

// Provides further details on the reference and status on the original transactions, included in the original instruction, to which the return message applies.
type PaymentTransaction34 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the returned transaction.
	// Usage: The instructing party is the party sending the return message and not the party that sent the original instruction that is being returned.
	ReturnIdentification *Max35Text `xml:"RtrId,omitempty"`

	// Provides information on the original message.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Unique reference, as assigned by the original clearing system, to unambiguously identify the original instruction.
	OriginalClearingSystemReference *Max35Text `xml:"OrgnlClrSysRef,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the instructing agent and the instructed agent.
	OriginalInterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent in the returned transaction.
	ReturnedInterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"RtrdIntrBkSttlmAmt"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	//
	// Usage: the InterbankSettlementDate is the interbank settlement date of the return message, and not of the original instruction.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Amount of money to be moved between the debtor and the creditor, before deduction of charges, in the returned transaction.
	ReturnedInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"RtrdInstdAmt,omitempty"`

	// Factor used to convert an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Amount of money asked or paid as compensation for the processing of the instruction.
	CompensationAmount *ActiveOrHistoricCurrencyAndAmount `xml:"CompstnAmt,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	//
	// Usage: The ChargeBearer applies to the return message, not to the original instruction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Provides information on the charges to be paid by the charge bearer(s) related to the processing of the return transaction.
	ChargesInformation []*Charges2 `xml:"ChrgsInf,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructing agent is the party sending the return message and not the party that sent the original instruction that is being returned.
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructed agent is the party receiving the return message and not the party that received the original instruction that is being returned.
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`

	// Provides detailed information on the return reason.
	ReturnReasonInformation []*PaymentReturnReason1 `xml:"RtrRsnInf,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference16 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransaction34) SetReturnIdentification(value string) {
	p.ReturnIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction34) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	p.OriginalGroupInformation = new(OriginalGroupInformation3)
	return p.OriginalGroupInformation
}

func (p *PaymentTransaction34) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction34) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction34) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction34) SetOriginalClearingSystemReference(value string) {
	p.OriginalClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction34) SetOriginalInterbankSettlementAmount(value, currency string) {
	p.OriginalInterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction34) SetReturnedInterbankSettlementAmount(value, currency string) {
	p.ReturnedInterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction34) SetInterbankSettlementDate(value string) {
	p.InterbankSettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction34) SetReturnedInstructedAmount(value, currency string) {
	p.ReturnedInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction34) SetExchangeRate(value string) {
	p.ExchangeRate = (*BaseOneRate)(&value)
}

func (p *PaymentTransaction34) SetCompensationAmount(value, currency string) {
	p.CompensationAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction34) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentTransaction34) AddChargesInformation() *Charges2 {
	newValue := new(Charges2)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransaction34) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructingAgent
}

func (p *PaymentTransaction34) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructedAgent
}

func (p *PaymentTransaction34) AddReturnReasonInformation() *PaymentReturnReason1 {
	newValue := new(PaymentReturnReason1)
	p.ReturnReasonInformation = append(p.ReturnReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction34) AddOriginalTransactionReference() *OriginalTransactionReference16 {
	p.OriginalTransactionReference = new(OriginalTransactionReference16)
	return p.OriginalTransactionReference
}

// Provides further details on the reference and status on the original transactions, included in the original instruction, to which the reversal message applies.
type PaymentTransaction35 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reversed transaction.
	// Usage: The instructing party is the party sending the reversal message and not the party that sent the original instruction that is being reversed.
	ReversalIdentification *Max35Text `xml:"RvslId,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the debtor and the creditor, before deduction of charges, expressed in the currency, as ordered by the original initiating party.
	OriginalInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt,omitempty"`

	// Amount of money to be moved between the debtor and the creditor, before deduction of charges, in the reversed transaction.
	ReversedInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"RvsdInstdAmt,omitempty"`

	// Specifies if the creditor and/or debtor will bear the charges associated with the processing of the payment transaction.
	//
	// Usage: The ChargeBearer applies to the reversal message, not to the original instruction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Provides detailed information on the reversal reason.
	ReversalReasonInformation []*PaymentReversalReason7 `xml:"RvslRsnInf,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference16 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransaction35) SetReversalIdentification(value string) {
	p.ReversalIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction35) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction35) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction35) SetOriginalInstructedAmount(value, currency string) {
	p.OriginalInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction35) SetReversedInstructedAmount(value, currency string) {
	p.ReversedInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction35) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentTransaction35) AddReversalReasonInformation() *PaymentReversalReason7 {
	newValue := new(PaymentReversalReason7)
	p.ReversalReasonInformation = append(p.ReversalReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction35) AddOriginalTransactionReference() *OriginalTransactionReference16 {
	p.OriginalTransactionReference = new(OriginalTransactionReference16)
	return p.OriginalTransactionReference
}

// Provides further details on the reference and status on the original transactions, included in the original instruction, to which the reversal message applies.
type PaymentTransaction36 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reversed transaction.
	// Usage: The instructing party is the party sending the reversal message and not the party that sent the original instruction that is being reversed.
	ReversalIdentification *Max35Text `xml:"RvslId,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Unique reference, as assigned by the original clearing system, to unambiguously identify the original instruction.
	OriginalClearingSystemReference *Max35Text `xml:"OrgnlClrSysRef,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the instructing agent and the instructed agent.
	OriginalInterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent in the reversed transaction.
	ReversedInterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"RvsdIntrBkSttlmAmt"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	//
	// Usage: The InterbankSettlementDate is the interbank settlement date of the reversal message, and not of the original instruction.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Amount of money to be moved between the debtor and the creditor, before deduction of charges, in the reversed transaction.
	ReversedInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"RvsdInstdAmt,omitempty"`

	// Factor used to convert an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Amount of money asked or paid as compensation for the processing of the instruction.
	CompensationAmount *ActiveOrHistoricCurrencyAndAmount `xml:"CompstnAmt,omitempty"`

	// Specifies if the creditor and/or debtor will bear the charges associated with the processing of the payment transaction.
	//
	// Usage: The ChargeBearer applies to the reversal message, not to the original instruction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Provides information on the charges to be paid by the charge bearer(s) related to the processing of the reversal transaction.
	ChargesInformation []*Charges2 `xml:"ChrgsInf,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructing agent is the party sending the reversal message and not the party that sent the original instruction that is being reversed.
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructed agent is the party receiving the reversal message and not the party that received the original instruction that is being reversed.
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`

	// Provides detailed information on the reversal reason.
	ReversalReasonInformation []*PaymentReversalReason7 `xml:"RvslRsnInf,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference16 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransaction36) SetReversalIdentification(value string) {
	p.ReversalIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction36) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction36) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction36) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction36) SetOriginalClearingSystemReference(value string) {
	p.OriginalClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction36) SetOriginalInterbankSettlementAmount(value, currency string) {
	p.OriginalInterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction36) SetReversedInterbankSettlementAmount(value, currency string) {
	p.ReversedInterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction36) SetInterbankSettlementDate(value string) {
	p.InterbankSettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction36) SetReversedInstructedAmount(value, currency string) {
	p.ReversedInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction36) SetExchangeRate(value string) {
	p.ExchangeRate = (*BaseOneRate)(&value)
}

func (p *PaymentTransaction36) SetCompensationAmount(value, currency string) {
	p.CompensationAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction36) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentTransaction36) AddChargesInformation() *Charges2 {
	newValue := new(Charges2)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransaction36) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructingAgent
}

func (p *PaymentTransaction36) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructedAgent
}

func (p *PaymentTransaction36) AddReversalReasonInformation() *PaymentReversalReason7 {
	newValue := new(PaymentReversalReason7)
	p.ReversalReasonInformation = append(p.ReversalReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction36) AddOriginalTransactionReference() *OriginalTransactionReference16 {
	p.OriginalTransactionReference = new(OriginalTransactionReference16)
	return p.OriginalTransactionReference
}

// Set of elements used to provide reference and status information on the original transactions, included in the original instruction, to which the cancellation request message applies.
type PaymentTransaction37 struct {

	// Unique identification, as assigned by the assigner, to unambiguously identify a cancellation request.
	//
	// Usage: The cancellation request identification can be used for reconciliation or to link tasks relating to the cancellation request.
	CancellationIdentification *Max35Text `xml:"CxlId,omitempty"`

	// Set of elements to uniquely and unambiguously identify an exception or an investigation workflow.
	Case *Case3 `xml:"Case,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the debtor and the creditor, before deduction of charges, expressed in the currency, as ordered by the original initiating party.
	OriginalInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt,omitempty"`

	// Date at which the initiating party originally requested the clearing agent to process the payment.
	OriginalRequestedExecutionDate *ISODate `xml:"OrgnlReqdExctnDt,omitempty"`

	// Date at which the creditor originally requested the collection of the amount of money from the debtor.
	OriginalRequestedCollectionDate *ISODate `xml:"OrgnlReqdColltnDt,omitempty"`

	// Provides detailed information on the cancellation reason.
	CancellationReasonInformation []*PaymentCancellationReason2 `xml:"CxlRsnInf,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference16 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransaction37) SetCancellationIdentification(value string) {
	p.CancellationIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction37) AddCase() *Case3 {
	p.Case = new(Case3)
	return p.Case
}

func (p *PaymentTransaction37) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction37) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction37) SetOriginalInstructedAmount(value, currency string) {
	p.OriginalInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction37) SetOriginalRequestedExecutionDate(value string) {
	p.OriginalRequestedExecutionDate = (*ISODate)(&value)
}

func (p *PaymentTransaction37) SetOriginalRequestedCollectionDate(value string) {
	p.OriginalRequestedCollectionDate = (*ISODate)(&value)
}

func (p *PaymentTransaction37) AddCancellationReasonInformation() *PaymentCancellationReason2 {
	newValue := new(PaymentCancellationReason2)
	p.CancellationReasonInformation = append(p.CancellationReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction37) AddOriginalTransactionReference() *OriginalTransactionReference16 {
	p.OriginalTransactionReference = new(OriginalTransactionReference16)
	return p.OriginalTransactionReference
}

// Provides reference and status information on the original transactions, included in the original instruction, to which the cancellation request message applies.
type PaymentTransaction38 struct {

	// Unique and unambiguous identifier of a cancellation request, as assigned by the assigner.
	//
	// Usage: The cancellation request identification can be used for reconciliation or to link tasks relating to the cancellation request.
	CancellationIdentification *Max35Text `xml:"CxlId,omitempty"`

	// Set of elements to uniquely and unambiguously identify an exception or an investigation workflow.
	Case *Case3 `xml:"Case,omitempty"`

	// Provides information on the original message.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Unique reference, as assigned by the original clearing system, to unambiguously identify the original instruction.
	OriginalClearingSystemReference *Max35Text `xml:"OrgnlClrSysRef,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the instructing agent and the instructed agent.
	OriginalInterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty"`

	// Date, as provided in the original transaction, on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	OriginalInterbankSettlementDate *ISODate `xml:"OrgnlIntrBkSttlmDt,omitempty"`

	// Party who assigns the case.
	// Usage: This is also the agent that instructs the next party in the chain to carry out the (set of) cancellation request(s).
	Assigner *BranchAndFinancialInstitutionIdentification5 `xml:"Assgnr,omitempty"`

	// Party to which the case is assigned.
	// Usage: This is also the agent that is instructed by the previous party in the chain to carry out the (set of) cancellation request(s).
	Assignee *BranchAndFinancialInstitutionIdentification5 `xml:"Assgne,omitempty"`

	// Provides detailed information on the cancellation reason.
	CancellationReasonInformation []*PaymentCancellationReason2 `xml:"CxlRsnInf,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference16 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransaction38) SetCancellationIdentification(value string) {
	p.CancellationIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction38) AddCase() *Case3 {
	p.Case = new(Case3)
	return p.Case
}

func (p *PaymentTransaction38) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	p.OriginalGroupInformation = new(OriginalGroupInformation3)
	return p.OriginalGroupInformation
}

func (p *PaymentTransaction38) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction38) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction38) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction38) SetOriginalClearingSystemReference(value string) {
	p.OriginalClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction38) SetOriginalInterbankSettlementAmount(value, currency string) {
	p.OriginalInterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction38) SetOriginalInterbankSettlementDate(value string) {
	p.OriginalInterbankSettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction38) AddAssigner() *BranchAndFinancialInstitutionIdentification5 {
	p.Assigner = new(BranchAndFinancialInstitutionIdentification5)
	return p.Assigner
}

func (p *PaymentTransaction38) AddAssignee() *BranchAndFinancialInstitutionIdentification5 {
	p.Assignee = new(BranchAndFinancialInstitutionIdentification5)
	return p.Assignee
}

func (p *PaymentTransaction38) AddCancellationReasonInformation() *PaymentCancellationReason2 {
	newValue := new(PaymentCancellationReason2)
	p.CancellationReasonInformation = append(p.CancellationReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction38) AddOriginalTransactionReference() *OriginalTransactionReference16 {
	p.OriginalTransactionReference = new(OriginalTransactionReference16)
	return p.OriginalTransactionReference
}

// Provides reference and status information on the original transactions, included in the original instruction, to which the cancellation request message applies.
type PaymentTransaction39 struct {

	// Unique and unambiguous identifier of a cancellation request status, as assigned by the assigner.
	//
	// Usage: The cancellation status identification can be used for reconciliation or to link tasks relating to the cancellation request.
	CancellationStatusIdentification *Max35Text `xml:"CxlStsId,omitempty"`

	// Identifies the resolved case.
	ResolvedCase *Case3 `xml:"RslvdCase,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Specifies the status of the transaction cancellation request.
	TransactionCancellationStatus *CancellationIndividualStatus1Code `xml:"TxCxlSts,omitempty"`

	// Provides detailed information on the cancellation status reason.
	CancellationStatusReasonInformation []*CancellationStatusReason2 `xml:"CxlStsRsnInf,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the debtor and the creditor, before deduction of charges, expressed in the currency, as ordered by the original initiating party.
	OriginalInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt,omitempty"`

	// Date at which the initiating party originally requested the clearing agent to process the payment.
	OriginalRequestedExecutionDate *ISODate `xml:"OrgnlReqdExctnDt,omitempty"`

	// Date at which the creditor originally requested the collection of the amount of money from the debtor.
	OriginalRequestedCollectionDate *ISODate `xml:"OrgnlReqdColltnDt,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference16 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransaction39) SetCancellationStatusIdentification(value string) {
	p.CancellationStatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction39) AddResolvedCase() *Case3 {
	p.ResolvedCase = new(Case3)
	return p.ResolvedCase
}

func (p *PaymentTransaction39) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction39) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction39) SetTransactionCancellationStatus(value string) {
	p.TransactionCancellationStatus = (*CancellationIndividualStatus1Code)(&value)
}

func (p *PaymentTransaction39) AddCancellationStatusReasonInformation() *CancellationStatusReason2 {
	newValue := new(CancellationStatusReason2)
	p.CancellationStatusReasonInformation = append(p.CancellationStatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction39) SetOriginalInstructedAmount(value, currency string) {
	p.OriginalInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction39) SetOriginalRequestedExecutionDate(value string) {
	p.OriginalRequestedExecutionDate = (*ISODate)(&value)
}

func (p *PaymentTransaction39) SetOriginalRequestedCollectionDate(value string) {
	p.OriginalRequestedCollectionDate = (*ISODate)(&value)
}

func (p *PaymentTransaction39) AddOriginalTransactionReference() *OriginalTransactionReference16 {
	p.OriginalTransactionReference = new(OriginalTransactionReference16)
	return p.OriginalTransactionReference
}

// Provides further details on the reference and status on the original transactions, included in the original instruction, to which the cancellation request message applies.
type PaymentTransaction40 struct {

	// Unique and unambiguous identifier of a cancellation request status, as assigned by the assigner.
	//
	// Usage: The cancellation status identification can be used for reconciliation or to link tasks relating to the cancellation request.
	CancellationStatusIdentification *Max35Text `xml:"CxlStsId,omitempty"`

	// Identifies the resolved case.
	ResolvedCase *Case3 `xml:"RslvdCase,omitempty"`

	// Provides information on the original message.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Unique reference, as assigned by the original clearing system, to unambiguously identify the original instruction.
	OriginalClearingSystemReference *Max35Text `xml:"OrgnlClrSysRef,omitempty"`

	// Specifies the status of the transaction cancellation request.
	TransactionCancellationStatus *CancellationIndividualStatus1Code `xml:"TxCxlSts,omitempty"`

	// Provides detailed information on the cancellation status reason.
	CancellationStatusReasonInformation []*CancellationStatusReason2 `xml:"CxlStsRsnInf,omitempty"`

	// Reference of a return or a reversal transaction that is initiated to fix the case under investigation as part of the resolution.
	ResolutionRelatedInformation *ResolutionInformation1 `xml:"RsltnRltdInf,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the instructing agent and the instructed agent.
	OriginalInterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty"`

	// Date, as provided in the original transaction, on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	OriginalInterbankSettlementDate *ISODate `xml:"OrgnlIntrBkSttlmDt,omitempty"`

	// Party who assigns the case.
	// Usage: This is also the agent that instructs the next party in the chain to carry out the (set of) cancellation request(s).
	Assigner *Party12Choice `xml:"Assgnr,omitempty"`

	// Party to which the case is assigned.
	// Usage: This is also the agent that is instructed by the previous party in the chain to carry out the (set of) cancellation request(s).
	Assignee *Party12Choice `xml:"Assgne,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference16 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransaction40) SetCancellationStatusIdentification(value string) {
	p.CancellationStatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction40) AddResolvedCase() *Case3 {
	p.ResolvedCase = new(Case3)
	return p.ResolvedCase
}

func (p *PaymentTransaction40) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	p.OriginalGroupInformation = new(OriginalGroupInformation3)
	return p.OriginalGroupInformation
}

func (p *PaymentTransaction40) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction40) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction40) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction40) SetOriginalClearingSystemReference(value string) {
	p.OriginalClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction40) SetTransactionCancellationStatus(value string) {
	p.TransactionCancellationStatus = (*CancellationIndividualStatus1Code)(&value)
}

func (p *PaymentTransaction40) AddCancellationStatusReasonInformation() *CancellationStatusReason2 {
	newValue := new(CancellationStatusReason2)
	p.CancellationStatusReasonInformation = append(p.CancellationStatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction40) AddResolutionRelatedInformation() *ResolutionInformation1 {
	p.ResolutionRelatedInformation = new(ResolutionInformation1)
	return p.ResolutionRelatedInformation
}

func (p *PaymentTransaction40) SetOriginalInterbankSettlementAmount(value, currency string) {
	p.OriginalInterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction40) SetOriginalInterbankSettlementDate(value string) {
	p.OriginalInterbankSettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction40) AddAssigner() *Party12Choice {
	p.Assigner = new(Party12Choice)
	return p.Assigner
}

func (p *PaymentTransaction40) AddAssignee() *Party12Choice {
	p.Assignee = new(Party12Choice)
	return p.Assignee
}

func (p *PaymentTransaction40) AddOriginalTransactionReference() *OriginalTransactionReference16 {
	p.OriginalTransactionReference = new(OriginalTransactionReference16)
	return p.OriginalTransactionReference
}

// Provides further details on the original transactions, to which the status report message refers.
type PaymentTransaction41 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reported status.
	// Usage: The instructing party is the party sending the status message and not the party that sent the original instruction that is being reported on.
	StatusIdentification *Max35Text `xml:"StsId,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Specifies the status of a transaction, in a coded form.
	TransactionStatus *TransactionIndividualStatus3Code `xml:"TxSts,omitempty"`

	// Provides detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Provides information on the charges related to the processing of the rejection of the instruction.
	// Usage: This is passed on for information purposes only. Settlement of the charges will be done separately.
	ChargesInformation []*Charges2 `xml:"ChrgsInf,omitempty"`

	// Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.
	AcceptanceDateTime *ISODateTime `xml:"AccptncDtTm,omitempty"`

	// Unique reference, as assigned by the account servicing institution, to unambiguously identify the instruction.
	AccountServicerReference *Max35Text `xml:"AcctSvcrRef,omitempty"`

	// Unique reference, as assigned by a clearing system, to unambiguously identify the instruction.
	ClearingSystemReference *Max35Text `xml:"ClrSysRef,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference17 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransaction41) SetStatusIdentification(value string) {
	p.StatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction41) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction41) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction41) SetTransactionStatus(value string) {
	p.TransactionStatus = (*TransactionIndividualStatus3Code)(&value)
}

func (p *PaymentTransaction41) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	p.StatusReasonInformation = append(p.StatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction41) AddChargesInformation() *Charges2 {
	newValue := new(Charges2)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransaction41) SetAcceptanceDateTime(value string) {
	p.AcceptanceDateTime = (*ISODateTime)(&value)
}

func (p *PaymentTransaction41) SetAccountServicerReference(value string) {
	p.AccountServicerReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction41) SetClearingSystemReference(value string) {
	p.ClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction41) AddOriginalTransactionReference() *OriginalTransactionReference17 {
	p.OriginalTransactionReference = new(OriginalTransactionReference17)
	return p.OriginalTransactionReference
}

// Provides further details on the reference and status on the original transactions, included in the original instruction, to which the reversal message applies.
type PaymentTransaction42 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reversed transaction.
	// Usage: The instructing party is the party sending the reversal message and not the party that sent the original instruction that is being reversed.
	ReversalIdentification *Max35Text `xml:"RvslId,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the debtor and the creditor, before deduction of charges, expressed in the currency, as ordered by the original initiating party.
	OriginalInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt,omitempty"`

	// Amount of money to be moved between the debtor and the creditor, before deduction of charges, in the reversed transaction.
	ReversedInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"RvsdInstdAmt,omitempty"`

	// Specifies if the creditor and/or debtor will bear the charges associated with the processing of the payment transaction.
	//
	// Usage: The ChargeBearer applies to the reversal message, not to the original instruction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Provides detailed information on the reversal reason.
	ReversalReasonInformation []*PaymentReversalReason7 `xml:"RvslRsnInf,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference16 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction42) SetReversalIdentification(value string) {
	p.ReversalIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction42) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction42) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction42) SetOriginalInstructedAmount(value, currency string) {
	p.OriginalInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction42) SetReversedInstructedAmount(value, currency string) {
	p.ReversedInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction42) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentTransaction42) AddReversalReasonInformation() *PaymentReversalReason7 {
	newValue := new(PaymentReversalReason7)
	p.ReversalReasonInformation = append(p.ReversalReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction42) AddOriginalTransactionReference() *OriginalTransactionReference16 {
	p.OriginalTransactionReference = new(OriginalTransactionReference16)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction42) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

// Provides further details on the original transactions, to which the status report message refers.
type PaymentTransaction43 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reported status.
	// Usage: The instructing party is the party sending the status message and not the party that sent the original instruction that is being reported on.
	StatusIdentification *Max35Text `xml:"StsId,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Specifies the status of a transaction, in a coded form.
	TransactionStatus *TransactionIndividualStatus3Code `xml:"TxSts,omitempty"`

	// Provides detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Provides information on the charges related to the processing of the rejection of the instruction.
	//
	// Usage: This is passed on for information purposes only. Settlement of the charges will be done separately.
	ChargesInformation []*Charges2 `xml:"ChrgsInf,omitempty"`

	// Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.
	AcceptanceDateTime *ISODateTime `xml:"AccptncDtTm,omitempty"`

	// Unique reference, as assigned by the account servicing institution, to unambiguously identify the instruction.
	AccountServicerReference *Max35Text `xml:"AcctSvcrRef,omitempty"`

	// Unique reference, as assigned by a clearing system, to unambiguously identify the instruction.
	ClearingSystemReference *Max35Text `xml:"ClrSysRef,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructing agent is the party sending the status message and not the party that sent the original instruction that is being reported on.
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructed agent is the party receiving the status message and not the party that received the original instruction that is being reported on.
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference16 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction43) SetStatusIdentification(value string) {
	p.StatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction43) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction43) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction43) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction43) SetTransactionStatus(value string) {
	p.TransactionStatus = (*TransactionIndividualStatus3Code)(&value)
}

func (p *PaymentTransaction43) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	p.StatusReasonInformation = append(p.StatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction43) AddChargesInformation() *Charges2 {
	newValue := new(Charges2)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransaction43) SetAcceptanceDateTime(value string) {
	p.AcceptanceDateTime = (*ISODateTime)(&value)
}

func (p *PaymentTransaction43) SetAccountServicerReference(value string) {
	p.AccountServicerReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction43) SetClearingSystemReference(value string) {
	p.ClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction43) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructingAgent
}

func (p *PaymentTransaction43) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructedAgent
}

func (p *PaymentTransaction43) AddOriginalTransactionReference() *OriginalTransactionReference16 {
	p.OriginalTransactionReference = new(OriginalTransactionReference16)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction43) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

// Provides further details on the reference and status on the original transactions, included in the original instruction, to which the return message applies.
type PaymentTransaction44 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the returned transaction.
	// Usage: The instructing party is the party sending the return message and not the party that sent the original instruction that is being returned.
	ReturnIdentification *Max35Text `xml:"RtrId,omitempty"`

	// Provides information on the original message.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Unique reference, as assigned by the original clearing system, to unambiguously identify the original instruction.
	OriginalClearingSystemReference *Max35Text `xml:"OrgnlClrSysRef,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the instructing agent and the instructed agent.
	OriginalInterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent in the returned transaction.
	ReturnedInterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"RtrdIntrBkSttlmAmt"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	//
	// Usage: the InterbankSettlementDate is the interbank settlement date of the return message, and not of the original instruction.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Amount of money to be moved between the debtor and the creditor, before deduction of charges, in the returned transaction.
	ReturnedInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"RtrdInstdAmt,omitempty"`

	// Factor used to convert an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Amount of money asked or paid as compensation for the processing of the instruction.
	CompensationAmount *ActiveOrHistoricCurrencyAndAmount `xml:"CompstnAmt,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	//
	// Usage: The ChargeBearer applies to the return message, not to the original instruction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Provides information on the charges to be paid by the charge bearer(s) related to the processing of the return transaction.
	ChargesInformation []*Charges2 `xml:"ChrgsInf,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructing agent is the party sending the return message and not the party that sent the original instruction that is being returned.
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructed agent is the party receiving the return message and not the party that received the original instruction that is being returned.
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`

	// Provides detailed information on the return reason.
	ReturnReasonInformation []*PaymentReturnReason1 `xml:"RtrRsnInf,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference16 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction44) SetReturnIdentification(value string) {
	p.ReturnIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction44) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	p.OriginalGroupInformation = new(OriginalGroupInformation3)
	return p.OriginalGroupInformation
}

func (p *PaymentTransaction44) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction44) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction44) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction44) SetOriginalClearingSystemReference(value string) {
	p.OriginalClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction44) SetOriginalInterbankSettlementAmount(value, currency string) {
	p.OriginalInterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction44) SetReturnedInterbankSettlementAmount(value, currency string) {
	p.ReturnedInterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction44) SetInterbankSettlementDate(value string) {
	p.InterbankSettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction44) SetReturnedInstructedAmount(value, currency string) {
	p.ReturnedInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction44) SetExchangeRate(value string) {
	p.ExchangeRate = (*BaseOneRate)(&value)
}

func (p *PaymentTransaction44) SetCompensationAmount(value, currency string) {
	p.CompensationAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction44) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentTransaction44) AddChargesInformation() *Charges2 {
	newValue := new(Charges2)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransaction44) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructingAgent
}

func (p *PaymentTransaction44) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructedAgent
}

func (p *PaymentTransaction44) AddReturnReasonInformation() *PaymentReturnReason1 {
	newValue := new(PaymentReturnReason1)
	p.ReturnReasonInformation = append(p.ReturnReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction44) AddOriginalTransactionReference() *OriginalTransactionReference16 {
	p.OriginalTransactionReference = new(OriginalTransactionReference16)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction44) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

// Provides further details on the reference and status on the original transactions, included in the original instruction, to which the reversal message applies.
type PaymentTransaction45 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reversed transaction.
	// Usage: The instructing party is the party sending the reversal message and not the party that sent the original instruction that is being reversed.
	ReversalIdentification *Max35Text `xml:"RvslId,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Unique reference, as assigned by the original clearing system, to unambiguously identify the original instruction.
	OriginalClearingSystemReference *Max35Text `xml:"OrgnlClrSysRef,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the instructing agent and the instructed agent.
	OriginalInterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent in the reversed transaction.
	ReversedInterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"RvsdIntrBkSttlmAmt"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	//
	// Usage: The InterbankSettlementDate is the interbank settlement date of the reversal message, and not of the original instruction.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Amount of money to be moved between the debtor and the creditor, before deduction of charges, in the reversed transaction.
	ReversedInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"RvsdInstdAmt,omitempty"`

	// Factor used to convert an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Amount of money asked or paid as compensation for the processing of the instruction.
	CompensationAmount *ActiveOrHistoricCurrencyAndAmount `xml:"CompstnAmt,omitempty"`

	// Specifies if the creditor and/or debtor will bear the charges associated with the processing of the payment transaction.
	//
	// Usage: The ChargeBearer applies to the reversal message, not to the original instruction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Provides information on the charges to be paid by the charge bearer(s) related to the processing of the reversal transaction.
	ChargesInformation []*Charges2 `xml:"ChrgsInf,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructing agent is the party sending the reversal message and not the party that sent the original instruction that is being reversed.
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructed agent is the party receiving the reversal message and not the party that received the original instruction that is being reversed.
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`

	// Provides detailed information on the reversal reason.
	ReversalReasonInformation []*PaymentReversalReason7 `xml:"RvslRsnInf,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference16 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction45) SetReversalIdentification(value string) {
	p.ReversalIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction45) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction45) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction45) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction45) SetOriginalClearingSystemReference(value string) {
	p.OriginalClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction45) SetOriginalInterbankSettlementAmount(value, currency string) {
	p.OriginalInterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction45) SetReversedInterbankSettlementAmount(value, currency string) {
	p.ReversedInterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction45) SetInterbankSettlementDate(value string) {
	p.InterbankSettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction45) SetReversedInstructedAmount(value, currency string) {
	p.ReversedInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction45) SetExchangeRate(value string) {
	p.ExchangeRate = (*BaseOneRate)(&value)
}

func (p *PaymentTransaction45) SetCompensationAmount(value, currency string) {
	p.CompensationAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction45) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentTransaction45) AddChargesInformation() *Charges2 {
	newValue := new(Charges2)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransaction45) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructingAgent
}

func (p *PaymentTransaction45) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructedAgent
}

func (p *PaymentTransaction45) AddReversalReasonInformation() *PaymentReversalReason7 {
	newValue := new(PaymentReversalReason7)
	p.ReversalReasonInformation = append(p.ReversalReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction45) AddOriginalTransactionReference() *OriginalTransactionReference16 {
	p.OriginalTransactionReference = new(OriginalTransactionReference16)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction45) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

// Provides further details on the original transactions, to which the status report message refers.
type PaymentTransaction46 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reported status.
	// Usage: The instructing party is the party sending the status message and not the party that sent the original instruction that is being reported on.
	StatusIdentification *Max35Text `xml:"StsId,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Specifies the status of a transaction, in a coded form.
	TransactionStatus *TransactionIndividualStatus3Code `xml:"TxSts,omitempty"`

	// Provides detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Provides information on the charges related to the processing of the rejection of the instruction.
	// Usage: This is passed on for information purposes only. Settlement of the charges will be done separately.
	ChargesInformation []*Charges2 `xml:"ChrgsInf,omitempty"`

	// Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.
	AcceptanceDateTime *ISODateTime `xml:"AccptncDtTm,omitempty"`

	// Unique reference, as assigned by the account servicing institution, to unambiguously identify the instruction.
	AccountServicerReference *Max35Text `xml:"AcctSvcrRef,omitempty"`

	// Unique reference, as assigned by a clearing system, to unambiguously identify the instruction.
	ClearingSystemReference *Max35Text `xml:"ClrSysRef,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference16 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction46) SetStatusIdentification(value string) {
	p.StatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction46) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction46) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction46) SetTransactionStatus(value string) {
	p.TransactionStatus = (*TransactionIndividualStatus3Code)(&value)
}

func (p *PaymentTransaction46) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	p.StatusReasonInformation = append(p.StatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction46) AddChargesInformation() *Charges2 {
	newValue := new(Charges2)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransaction46) SetAcceptanceDateTime(value string) {
	p.AcceptanceDateTime = (*ISODateTime)(&value)
}

func (p *PaymentTransaction46) SetAccountServicerReference(value string) {
	p.AccountServicerReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction46) SetClearingSystemReference(value string) {
	p.ClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction46) AddOriginalTransactionReference() *OriginalTransactionReference16 {
	p.OriginalTransactionReference = new(OriginalTransactionReference16)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction46) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

// Set of elements used to provide reference and status information on the original transactions, included in the original instruction, to which the cancellation request message applies.
type PaymentTransaction47 struct {

	// Unique identification, as assigned by the assigner, to unambiguously identify a cancellation request.
	//
	// Usage: The cancellation request identification can be used for reconciliation or to link tasks relating to the cancellation request.
	CancellationIdentification *Max35Text `xml:"CxlId,omitempty"`

	// Set of elements to uniquely and unambiguously identify an exception or an investigation workflow.
	Case *Case3 `xml:"Case,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the debtor and the creditor, before deduction of charges, expressed in the currency, as ordered by the original initiating party.
	OriginalInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt,omitempty"`

	// Date at which the initiating party originally requested the clearing agent to process the payment.
	OriginalRequestedExecutionDate *ISODate `xml:"OrgnlReqdExctnDt,omitempty"`

	// Date at which the creditor originally requested the collection of the amount of money from the debtor.
	OriginalRequestedCollectionDate *ISODate `xml:"OrgnlReqdColltnDt,omitempty"`

	// Provides detailed information on the cancellation reason.
	CancellationReasonInformation []*PaymentCancellationReason2 `xml:"CxlRsnInf,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference16 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction47) SetCancellationIdentification(value string) {
	p.CancellationIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction47) AddCase() *Case3 {
	p.Case = new(Case3)
	return p.Case
}

func (p *PaymentTransaction47) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction47) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction47) SetOriginalInstructedAmount(value, currency string) {
	p.OriginalInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction47) SetOriginalRequestedExecutionDate(value string) {
	p.OriginalRequestedExecutionDate = (*ISODate)(&value)
}

func (p *PaymentTransaction47) SetOriginalRequestedCollectionDate(value string) {
	p.OriginalRequestedCollectionDate = (*ISODate)(&value)
}

func (p *PaymentTransaction47) AddCancellationReasonInformation() *PaymentCancellationReason2 {
	newValue := new(PaymentCancellationReason2)
	p.CancellationReasonInformation = append(p.CancellationReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction47) AddOriginalTransactionReference() *OriginalTransactionReference16 {
	p.OriginalTransactionReference = new(OriginalTransactionReference16)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction47) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

// Provides reference and status information on the original transactions, included in the original instruction, to which the cancellation request message applies.
type PaymentTransaction48 struct {

	// Unique and unambiguous identifier of a cancellation request, as assigned by the assigner.
	//
	// Usage: The cancellation request identification can be used for reconciliation or to link tasks relating to the cancellation request.
	CancellationIdentification *Max35Text `xml:"CxlId,omitempty"`

	// Set of elements to uniquely and unambiguously identify an exception or an investigation workflow.
	Case *Case3 `xml:"Case,omitempty"`

	// Provides information on the original message.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Unique reference, as assigned by the original clearing system, to unambiguously identify the original instruction.
	OriginalClearingSystemReference *Max35Text `xml:"OrgnlClrSysRef,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the instructing agent and the instructed agent.
	OriginalInterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty"`

	// Date, as provided in the original transaction, on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	OriginalInterbankSettlementDate *ISODate `xml:"OrgnlIntrBkSttlmDt,omitempty"`

	// Party who assigns the case.
	// Usage: This is also the agent that instructs the next party in the chain to carry out the (set of) cancellation request(s).
	Assigner *BranchAndFinancialInstitutionIdentification5 `xml:"Assgnr,omitempty"`

	// Party to which the case is assigned.
	// Usage: This is also the agent that is instructed by the previous party in the chain to carry out the (set of) cancellation request(s).
	Assignee *BranchAndFinancialInstitutionIdentification5 `xml:"Assgne,omitempty"`

	// Provides detailed information on the cancellation reason.
	CancellationReasonInformation []*PaymentCancellationReason2 `xml:"CxlRsnInf,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference16 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction48) SetCancellationIdentification(value string) {
	p.CancellationIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction48) AddCase() *Case3 {
	p.Case = new(Case3)
	return p.Case
}

func (p *PaymentTransaction48) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	p.OriginalGroupInformation = new(OriginalGroupInformation3)
	return p.OriginalGroupInformation
}

func (p *PaymentTransaction48) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction48) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction48) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction48) SetOriginalClearingSystemReference(value string) {
	p.OriginalClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction48) SetOriginalInterbankSettlementAmount(value, currency string) {
	p.OriginalInterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction48) SetOriginalInterbankSettlementDate(value string) {
	p.OriginalInterbankSettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction48) AddAssigner() *BranchAndFinancialInstitutionIdentification5 {
	p.Assigner = new(BranchAndFinancialInstitutionIdentification5)
	return p.Assigner
}

func (p *PaymentTransaction48) AddAssignee() *BranchAndFinancialInstitutionIdentification5 {
	p.Assignee = new(BranchAndFinancialInstitutionIdentification5)
	return p.Assignee
}

func (p *PaymentTransaction48) AddCancellationReasonInformation() *PaymentCancellationReason2 {
	newValue := new(PaymentCancellationReason2)
	p.CancellationReasonInformation = append(p.CancellationReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction48) AddOriginalTransactionReference() *OriginalTransactionReference16 {
	p.OriginalTransactionReference = new(OriginalTransactionReference16)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction48) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

// Provides further details on the original transactions, to which the status report message refers.
type PaymentTransaction49 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reported status.
	// Usage: The instructing party is the party sending the status message and not the party that sent the original instruction that is being reported on.
	StatusIdentification *Max35Text `xml:"StsId,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Specifies the status of a transaction, in a coded form.
	TransactionStatus *TransactionIndividualStatus3Code `xml:"TxSts,omitempty"`

	// Provides detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Provides information on the charges related to the processing of the rejection of the instruction.
	// Usage: This is passed on for information purposes only. Settlement of the charges will be done separately.
	ChargesInformation []*Charges2 `xml:"ChrgsInf,omitempty"`

	// Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.
	AcceptanceDateTime *ISODateTime `xml:"AccptncDtTm,omitempty"`

	// Unique reference, as assigned by the account servicing institution, to unambiguously identify the instruction.
	AccountServicerReference *Max35Text `xml:"AcctSvcrRef,omitempty"`

	// Unique reference, as assigned by a clearing system, to unambiguously identify the instruction.
	ClearingSystemReference *Max35Text `xml:"ClrSysRef,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference19 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction49) SetStatusIdentification(value string) {
	p.StatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction49) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction49) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction49) SetTransactionStatus(value string) {
	p.TransactionStatus = (*TransactionIndividualStatus3Code)(&value)
}

func (p *PaymentTransaction49) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	p.StatusReasonInformation = append(p.StatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction49) AddChargesInformation() *Charges2 {
	newValue := new(Charges2)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransaction49) SetAcceptanceDateTime(value string) {
	p.AcceptanceDateTime = (*ISODateTime)(&value)
}

func (p *PaymentTransaction49) SetAccountServicerReference(value string) {
	p.AccountServicerReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction49) SetClearingSystemReference(value string) {
	p.ClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction49) AddOriginalTransactionReference() *OriginalTransactionReference19 {
	p.OriginalTransactionReference = new(OriginalTransactionReference19)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction49) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

// Provides further details on the reference and status on the original transactions, included in the original instruction, to which the return message applies.
type PaymentTransaction50 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the returned transaction.
	// Usage: The instructing party is the party sending the return message and not the party that sent the original instruction that is being returned.
	ReturnIdentification *Max35Text `xml:"RtrId,omitempty"`

	// Provides information on the original message.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Unique reference, as assigned by the original clearing system, to unambiguously identify the original instruction.
	OriginalClearingSystemReference *Max35Text `xml:"OrgnlClrSysRef,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the instructing agent and the instructed agent.
	OriginalInterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent in the returned transaction.
	ReturnedInterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"RtrdIntrBkSttlmAmt"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	//
	// Usage: the InterbankSettlementDate is the interbank settlement date of the return message, and not of the original instruction.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the settlement instruction.
	//
	//
	// Usage: the SettlementPriority is the settlement priority of the return message, and not of the original instruction.
	SettlementPriority *Priority3Code `xml:"SttlmPrty,omitempty"`

	// Amount of money to be moved between the debtor and the creditor, before deduction of charges, in the returned transaction.
	// Usage: This amount has to be transported unchanged through the transaction chain.
	ReturnedInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"RtrdInstdAmt,omitempty"`

	// Factor used to convert an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Amount of money asked or paid as compensation for the processing of the instruction.
	CompensationAmount *ActiveOrHistoricCurrencyAndAmount `xml:"CompstnAmt,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	//
	// Usage: The ChargeBearer applies to the return message, not to the original instruction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Provides information on the charges to be paid by the charge bearer(s) related to the processing of the return transaction.
	ChargesInformation []*Charges2 `xml:"ChrgsInf,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructing agent is the party sending the return message and not the party that sent the original instruction that is being returned.
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructed agent is the party receiving the return message and not the party that received the original instruction that is being returned.
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`

	// Provides detailed information on the return reason.
	ReturnReasonInformation []*PaymentReturnReason1 `xml:"RtrRsnInf,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference20 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction50) SetReturnIdentification(value string) {
	p.ReturnIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction50) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	p.OriginalGroupInformation = new(OriginalGroupInformation3)
	return p.OriginalGroupInformation
}

func (p *PaymentTransaction50) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction50) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction50) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction50) SetOriginalClearingSystemReference(value string) {
	p.OriginalClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction50) SetOriginalInterbankSettlementAmount(value, currency string) {
	p.OriginalInterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction50) SetReturnedInterbankSettlementAmount(value, currency string) {
	p.ReturnedInterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction50) SetInterbankSettlementDate(value string) {
	p.InterbankSettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction50) SetSettlementPriority(value string) {
	p.SettlementPriority = (*Priority3Code)(&value)
}

func (p *PaymentTransaction50) SetReturnedInstructedAmount(value, currency string) {
	p.ReturnedInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction50) SetExchangeRate(value string) {
	p.ExchangeRate = (*BaseOneRate)(&value)
}

func (p *PaymentTransaction50) SetCompensationAmount(value, currency string) {
	p.CompensationAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction50) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentTransaction50) AddChargesInformation() *Charges2 {
	newValue := new(Charges2)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransaction50) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructingAgent
}

func (p *PaymentTransaction50) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructedAgent
}

func (p *PaymentTransaction50) AddReturnReasonInformation() *PaymentReturnReason1 {
	newValue := new(PaymentReturnReason1)
	p.ReturnReasonInformation = append(p.ReturnReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction50) AddOriginalTransactionReference() *OriginalTransactionReference20 {
	p.OriginalTransactionReference = new(OriginalTransactionReference20)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction50) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

// Provides further details on the reference and status on the original transactions, included in the original instruction, to which the reversal message applies.
type PaymentTransaction51 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reversed transaction.
	// Usage: The instructing party is the party sending the reversal message and not the party that sent the original instruction that is being reversed.
	ReversalIdentification *Max35Text `xml:"RvslId,omitempty"`

	// Provides information on the original message.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Unique reference, as assigned by the original clearing system, to unambiguously identify the original instruction.
	OriginalClearingSystemReference *Max35Text `xml:"OrgnlClrSysRef,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the instructing agent and the instructed agent.
	OriginalInterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent in the reversed transaction.
	ReversedInterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"RvsdIntrBkSttlmAmt"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	//
	// Usage: The InterbankSettlementDate is the interbank settlement date of the reversal message, and not of the original instruction.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the settlement instruction.
	//
	//
	// Usage: the SettlementPriority is the settlement priority of the reversal message, and not of the original instruction.
	SettlementPriority *Priority3Code `xml:"SttlmPrty,omitempty"`

	// Amount of money to be moved between the debtor and the creditor, before deduction of charges, in the reversed transaction.
	// Usage: This amount has to be transported unchanged through the transaction chain.
	ReversedInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"RvsdInstdAmt,omitempty"`

	// Factor used to convert an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Amount of money asked or paid as compensation for the processing of the instruction.
	CompensationAmount *ActiveOrHistoricCurrencyAndAmount `xml:"CompstnAmt,omitempty"`

	// Specifies if the creditor and/or debtor will bear the charges associated with the processing of the payment transaction.
	//
	// Usage: The ChargeBearer applies to the reversal message, not to the original instruction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Provides information on the charges to be paid by the charge bearer(s) related to the processing of the reversal transaction.
	ChargesInformation []*Charges2 `xml:"ChrgsInf,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructing agent is the party sending the reversal message and not the party that sent the original instruction that is being reversed.
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructed agent is the party receiving the reversal message and not the party that received the original instruction that is being reversed.
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`

	// Provides detailed information on the reversal reason.
	ReversalReasonInformation []*PaymentReversalReason7 `xml:"RvslRsnInf,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference20 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction51) SetReversalIdentification(value string) {
	p.ReversalIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction51) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	p.OriginalGroupInformation = new(OriginalGroupInformation3)
	return p.OriginalGroupInformation
}

func (p *PaymentTransaction51) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction51) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction51) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction51) SetOriginalClearingSystemReference(value string) {
	p.OriginalClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction51) SetOriginalInterbankSettlementAmount(value, currency string) {
	p.OriginalInterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction51) SetReversedInterbankSettlementAmount(value, currency string) {
	p.ReversedInterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction51) SetInterbankSettlementDate(value string) {
	p.InterbankSettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction51) SetSettlementPriority(value string) {
	p.SettlementPriority = (*Priority3Code)(&value)
}

func (p *PaymentTransaction51) SetReversedInstructedAmount(value, currency string) {
	p.ReversedInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction51) SetExchangeRate(value string) {
	p.ExchangeRate = (*BaseOneRate)(&value)
}

func (p *PaymentTransaction51) SetCompensationAmount(value, currency string) {
	p.CompensationAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction51) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentTransaction51) AddChargesInformation() *Charges2 {
	newValue := new(Charges2)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransaction51) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructingAgent
}

func (p *PaymentTransaction51) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructedAgent
}

func (p *PaymentTransaction51) AddReversalReasonInformation() *PaymentReversalReason7 {
	newValue := new(PaymentReversalReason7)
	p.ReversalReasonInformation = append(p.ReversalReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction51) AddOriginalTransactionReference() *OriginalTransactionReference20 {
	p.OriginalTransactionReference = new(OriginalTransactionReference20)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction51) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

// Provides further details on the original transactions, to which the status report message refers.
type PaymentTransaction52 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reported status.
	// Usage: The instructing party is the party sending the status message and not the party that sent the original instruction that is being reported on.
	StatusIdentification *Max35Text `xml:"StsId,omitempty"`

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Specifies the status of a transaction, in a coded form.
	TransactionStatus *TransactionIndividualStatus3Code `xml:"TxSts,omitempty"`

	// Provides detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Provides information on the charges related to the processing of the rejection of the instruction.
	//
	// Usage: This is passed on for information purposes only. Settlement of the charges will be done separately.
	ChargesInformation []*Charges2 `xml:"ChrgsInf,omitempty"`

	// Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.
	AcceptanceDateTime *ISODateTime `xml:"AccptncDtTm,omitempty"`

	// Unique reference, as assigned by the account servicing institution, to unambiguously identify the instruction.
	AccountServicerReference *Max35Text `xml:"AcctSvcrRef,omitempty"`

	// Unique reference, as assigned by a clearing system, to unambiguously identify the instruction.
	ClearingSystemReference *Max35Text `xml:"ClrSysRef,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructing agent is the party sending the status message and not the party that sent the original instruction that is being reported on.
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructed agent is the party receiving the status message and not the party that received the original instruction that is being reported on.
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference20 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction52) SetStatusIdentification(value string) {
	p.StatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction52) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	p.OriginalGroupInformation = new(OriginalGroupInformation3)
	return p.OriginalGroupInformation
}

func (p *PaymentTransaction52) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction52) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction52) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction52) SetTransactionStatus(value string) {
	p.TransactionStatus = (*TransactionIndividualStatus3Code)(&value)
}

func (p *PaymentTransaction52) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	p.StatusReasonInformation = append(p.StatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction52) AddChargesInformation() *Charges2 {
	newValue := new(Charges2)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransaction52) SetAcceptanceDateTime(value string) {
	p.AcceptanceDateTime = (*ISODateTime)(&value)
}

func (p *PaymentTransaction52) SetAccountServicerReference(value string) {
	p.AccountServicerReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction52) SetClearingSystemReference(value string) {
	p.ClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction52) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructingAgent
}

func (p *PaymentTransaction52) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructedAgent
}

func (p *PaymentTransaction52) AddOriginalTransactionReference() *OriginalTransactionReference20 {
	p.OriginalTransactionReference = new(OriginalTransactionReference20)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction52) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

// Provides further details on the reference and status on the original transactions, included in the original instruction, to which the cancellation request message applies.
type PaymentTransaction53 struct {

	// Unique and unambiguous identifier of a cancellation request status, as assigned by the assigner.
	//
	// Usage: The cancellation status identification can be used for reconciliation or to link tasks relating to the cancellation request.
	CancellationStatusIdentification *Max35Text `xml:"CxlStsId,omitempty"`

	// Identifies the resolved case.
	ResolvedCase *Case3 `xml:"RslvdCase,omitempty"`

	// Provides information on the original message.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Unique reference, as assigned by the original clearing system, to unambiguously identify the original instruction.
	OriginalClearingSystemReference *Max35Text `xml:"OrgnlClrSysRef,omitempty"`

	// Specifies the status of the transaction cancellation request.
	TransactionCancellationStatus *CancellationIndividualStatus1Code `xml:"TxCxlSts,omitempty"`

	// Provides detailed information on the cancellation status reason.
	CancellationStatusReasonInformation []*CancellationStatusReason2 `xml:"CxlStsRsnInf,omitempty"`

	// Reference of a return or a reversal transaction that is initiated to fix the case under investigation as part of the resolution.
	ResolutionRelatedInformation *ResolutionInformation1 `xml:"RsltnRltdInf,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the instructing agent and the instructed agent.
	OriginalInterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty"`

	// Date, as provided in the original transaction, on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	OriginalInterbankSettlementDate *ISODate `xml:"OrgnlIntrBkSttlmDt,omitempty"`

	// Party who assigns the case.
	// Usage: This is also the agent that instructs the next party in the chain to carry out the (set of) cancellation request(s).
	Assigner *Party12Choice `xml:"Assgnr,omitempty"`

	// Party to which the case is assigned.
	// Usage: This is also the agent that is instructed by the previous party in the chain to carry out the (set of) cancellation request(s).
	Assignee *Party12Choice `xml:"Assgne,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference20 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransaction53) SetCancellationStatusIdentification(value string) {
	p.CancellationStatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction53) AddResolvedCase() *Case3 {
	p.ResolvedCase = new(Case3)
	return p.ResolvedCase
}

func (p *PaymentTransaction53) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	p.OriginalGroupInformation = new(OriginalGroupInformation3)
	return p.OriginalGroupInformation
}

func (p *PaymentTransaction53) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction53) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction53) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction53) SetOriginalClearingSystemReference(value string) {
	p.OriginalClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction53) SetTransactionCancellationStatus(value string) {
	p.TransactionCancellationStatus = (*CancellationIndividualStatus1Code)(&value)
}

func (p *PaymentTransaction53) AddCancellationStatusReasonInformation() *CancellationStatusReason2 {
	newValue := new(CancellationStatusReason2)
	p.CancellationStatusReasonInformation = append(p.CancellationStatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction53) AddResolutionRelatedInformation() *ResolutionInformation1 {
	p.ResolutionRelatedInformation = new(ResolutionInformation1)
	return p.ResolutionRelatedInformation
}

func (p *PaymentTransaction53) SetOriginalInterbankSettlementAmount(value, currency string) {
	p.OriginalInterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction53) SetOriginalInterbankSettlementDate(value string) {
	p.OriginalInterbankSettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction53) AddAssigner() *Party12Choice {
	p.Assigner = new(Party12Choice)
	return p.Assigner
}

func (p *PaymentTransaction53) AddAssignee() *Party12Choice {
	p.Assignee = new(Party12Choice)
	return p.Assignee
}

func (p *PaymentTransaction53) AddOriginalTransactionReference() *OriginalTransactionReference20 {
	p.OriginalTransactionReference = new(OriginalTransactionReference20)
	return p.OriginalTransactionReference
}

// Provides reference and status information on the original transactions, included in the original instruction, to which the cancellation request message applies.
type PaymentTransaction54 struct {

	// Unique and unambiguous identifier of a cancellation request status, as assigned by the assigner.
	//
	// Usage: The cancellation status identification can be used for reconciliation or to link tasks relating to the cancellation request.
	CancellationStatusIdentification *Max35Text `xml:"CxlStsId,omitempty"`

	// Identifies the resolved case.
	ResolvedCase *Case3 `xml:"RslvdCase,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Specifies the status of the transaction cancellation request.
	TransactionCancellationStatus *CancellationIndividualStatus1Code `xml:"TxCxlSts,omitempty"`

	// Provides detailed information on the cancellation status reason.
	CancellationStatusReasonInformation []*CancellationStatusReason2 `xml:"CxlStsRsnInf,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the debtor and the creditor, before deduction of charges, expressed in the currency, as ordered by the original initiating party.
	// Usage: This amount has to be transported unchanged through the transaction chain.
	OriginalInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt,omitempty"`

	// Date at which the initiating party originally requested the clearing agent to process the payment.
	OriginalRequestedExecutionDate *ISODate `xml:"OrgnlReqdExctnDt,omitempty"`

	// Date at which the creditor originally requested the collection of the amount of money from the debtor.
	OriginalRequestedCollectionDate *ISODate `xml:"OrgnlReqdColltnDt,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference20 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransaction54) SetCancellationStatusIdentification(value string) {
	p.CancellationStatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction54) AddResolvedCase() *Case3 {
	p.ResolvedCase = new(Case3)
	return p.ResolvedCase
}

func (p *PaymentTransaction54) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction54) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction54) SetTransactionCancellationStatus(value string) {
	p.TransactionCancellationStatus = (*CancellationIndividualStatus1Code)(&value)
}

func (p *PaymentTransaction54) AddCancellationStatusReasonInformation() *CancellationStatusReason2 {
	newValue := new(CancellationStatusReason2)
	p.CancellationStatusReasonInformation = append(p.CancellationStatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction54) SetOriginalInstructedAmount(value, currency string) {
	p.OriginalInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction54) SetOriginalRequestedExecutionDate(value string) {
	p.OriginalRequestedExecutionDate = (*ISODate)(&value)
}

func (p *PaymentTransaction54) SetOriginalRequestedCollectionDate(value string) {
	p.OriginalRequestedCollectionDate = (*ISODate)(&value)
}

func (p *PaymentTransaction54) AddOriginalTransactionReference() *OriginalTransactionReference20 {
	p.OriginalTransactionReference = new(OriginalTransactionReference20)
	return p.OriginalTransactionReference
}

// Provides reference and status information on the original transactions, included in the original instruction, to which the cancellation request message applies.
type PaymentTransaction55 struct {

	// Unique and unambiguous identifier of a cancellation request, as assigned by the assigner.
	//
	// Usage: The cancellation request identification can be used for reconciliation or to link tasks relating to the cancellation request.
	CancellationIdentification *Max35Text `xml:"CxlId,omitempty"`

	// Set of elements to uniquely and unambiguously identify an exception or an investigation workflow.
	Case *Case3 `xml:"Case,omitempty"`

	// Provides information on the original message.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Unique reference, as assigned by the original clearing system, to unambiguously identify the original instruction.
	OriginalClearingSystemReference *Max35Text `xml:"OrgnlClrSysRef,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the instructing agent and the instructed agent.
	OriginalInterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty"`

	// Date, as provided in the original transaction, on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	OriginalInterbankSettlementDate *ISODate `xml:"OrgnlIntrBkSttlmDt,omitempty"`

	// Party who assigns the case.
	// Usage: This is also the agent that instructs the next party in the chain to carry out the (set of) cancellation request(s).
	Assigner *BranchAndFinancialInstitutionIdentification5 `xml:"Assgnr,omitempty"`

	// Party to which the case is assigned.
	// Usage: This is also the agent that is instructed by the previous party in the chain to carry out the (set of) cancellation request(s).
	Assignee *BranchAndFinancialInstitutionIdentification5 `xml:"Assgne,omitempty"`

	// Provides detailed information on the cancellation reason.
	CancellationReasonInformation []*PaymentCancellationReason2 `xml:"CxlRsnInf,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference20 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction55) SetCancellationIdentification(value string) {
	p.CancellationIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction55) AddCase() *Case3 {
	p.Case = new(Case3)
	return p.Case
}

func (p *PaymentTransaction55) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	p.OriginalGroupInformation = new(OriginalGroupInformation3)
	return p.OriginalGroupInformation
}

func (p *PaymentTransaction55) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction55) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction55) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction55) SetOriginalClearingSystemReference(value string) {
	p.OriginalClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction55) SetOriginalInterbankSettlementAmount(value, currency string) {
	p.OriginalInterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction55) SetOriginalInterbankSettlementDate(value string) {
	p.OriginalInterbankSettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction55) AddAssigner() *BranchAndFinancialInstitutionIdentification5 {
	p.Assigner = new(BranchAndFinancialInstitutionIdentification5)
	return p.Assigner
}

func (p *PaymentTransaction55) AddAssignee() *BranchAndFinancialInstitutionIdentification5 {
	p.Assignee = new(BranchAndFinancialInstitutionIdentification5)
	return p.Assignee
}

func (p *PaymentTransaction55) AddCancellationReasonInformation() *PaymentCancellationReason2 {
	newValue := new(PaymentCancellationReason2)
	p.CancellationReasonInformation = append(p.CancellationReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction55) AddOriginalTransactionReference() *OriginalTransactionReference20 {
	p.OriginalTransactionReference = new(OriginalTransactionReference20)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction55) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

// Provides further details on the reference and status on the original transactions, included in the original instruction, to which the reversal message applies.
type PaymentTransaction56 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reversed transaction.
	// Usage: The instructing party is the party sending the reversal message and not the party that sent the original instruction that is being reversed.
	ReversalIdentification *Max35Text `xml:"RvslId,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the debtor and the creditor, before deduction of charges, expressed in the currency, as ordered by the original initiating party.
	// Usage: This amount has to be transported unchanged through the transaction chain.
	OriginalInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt,omitempty"`

	// Amount of money to be moved between the debtor and the creditor, before deduction of charges, in the reversed transaction.
	// Usage: This amount has to be transported unchanged through the transaction chain.
	ReversedInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"RvsdInstdAmt,omitempty"`

	// Specifies if the creditor and/or debtor will bear the charges associated with the processing of the payment transaction.
	//
	// Usage: The ChargeBearer applies to the reversal message, not to the original instruction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Provides detailed information on the reversal reason.
	ReversalReasonInformation []*PaymentReversalReason7 `xml:"RvslRsnInf,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference20 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction56) SetReversalIdentification(value string) {
	p.ReversalIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction56) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction56) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction56) SetOriginalInstructedAmount(value, currency string) {
	p.OriginalInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction56) SetReversedInstructedAmount(value, currency string) {
	p.ReversedInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction56) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentTransaction56) AddReversalReasonInformation() *PaymentReversalReason7 {
	newValue := new(PaymentReversalReason7)
	p.ReversalReasonInformation = append(p.ReversalReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction56) AddOriginalTransactionReference() *OriginalTransactionReference20 {
	p.OriginalTransactionReference = new(OriginalTransactionReference20)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction56) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

// Provides further details on the original transactions, to which the status report message refers.
type PaymentTransaction57 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reported status.
	// Usage: The instructing party is the party sending the status message and not the party that sent the original instruction that is being reported on.
	StatusIdentification *Max35Text `xml:"StsId,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Specifies the status of a transaction, in a coded form.
	TransactionStatus *TransactionIndividualStatus3Code `xml:"TxSts,omitempty"`

	// Provides detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Provides information on the charges related to the processing of the rejection of the instruction.
	// Usage: This is passed on for information purposes only. Settlement of the charges will be done separately.
	ChargesInformation []*Charges2 `xml:"ChrgsInf,omitempty"`

	// Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.
	AcceptanceDateTime *ISODateTime `xml:"AccptncDtTm,omitempty"`

	// Unique reference, as assigned by the account servicing institution, to unambiguously identify the instruction.
	AccountServicerReference *Max35Text `xml:"AcctSvcrRef,omitempty"`

	// Unique reference, as assigned by a clearing system, to unambiguously identify the instruction.
	ClearingSystemReference *Max35Text `xml:"ClrSysRef,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference20 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction57) SetStatusIdentification(value string) {
	p.StatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction57) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction57) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction57) SetTransactionStatus(value string) {
	p.TransactionStatus = (*TransactionIndividualStatus3Code)(&value)
}

func (p *PaymentTransaction57) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	p.StatusReasonInformation = append(p.StatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction57) AddChargesInformation() *Charges2 {
	newValue := new(Charges2)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransaction57) SetAcceptanceDateTime(value string) {
	p.AcceptanceDateTime = (*ISODateTime)(&value)
}

func (p *PaymentTransaction57) SetAccountServicerReference(value string) {
	p.AccountServicerReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction57) SetClearingSystemReference(value string) {
	p.ClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction57) AddOriginalTransactionReference() *OriginalTransactionReference20 {
	p.OriginalTransactionReference = new(OriginalTransactionReference20)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction57) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

// Set of elements used to provide reference and status information on the original transactions, included in the original instruction, to which the cancellation request message applies.
type PaymentTransaction58 struct {

	// Unique identification, as assigned by the assigner, to unambiguously identify a cancellation request.
	//
	// Usage: The cancellation request identification can be used for reconciliation or to link tasks relating to the cancellation request.
	CancellationIdentification *Max35Text `xml:"CxlId,omitempty"`

	// Set of elements to uniquely and unambiguously identify an exception or an investigation workflow.
	Case *Case3 `xml:"Case,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the debtor and the creditor, before deduction of charges, expressed in the currency, as ordered by the original initiating party.
	//
	// Usage: This amount has to be transported unchanged through the transaction chain.
	OriginalInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt,omitempty"`

	// Date at which the initiating party originally requested the clearing agent to process the payment.
	OriginalRequestedExecutionDate *ISODate `xml:"OrgnlReqdExctnDt,omitempty"`

	// Date at which the creditor originally requested the collection of the amount of money from the debtor.
	OriginalRequestedCollectionDate *ISODate `xml:"OrgnlReqdColltnDt,omitempty"`

	// Provides detailed information on the cancellation reason.
	CancellationReasonInformation []*PaymentCancellationReason2 `xml:"CxlRsnInf,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference20 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction58) SetCancellationIdentification(value string) {
	p.CancellationIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction58) AddCase() *Case3 {
	p.Case = new(Case3)
	return p.Case
}

func (p *PaymentTransaction58) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction58) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction58) SetOriginalInstructedAmount(value, currency string) {
	p.OriginalInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction58) SetOriginalRequestedExecutionDate(value string) {
	p.OriginalRequestedExecutionDate = (*ISODate)(&value)
}

func (p *PaymentTransaction58) SetOriginalRequestedCollectionDate(value string) {
	p.OriginalRequestedCollectionDate = (*ISODate)(&value)
}

func (p *PaymentTransaction58) AddCancellationReasonInformation() *PaymentCancellationReason2 {
	newValue := new(PaymentCancellationReason2)
	p.CancellationReasonInformation = append(p.CancellationReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction58) AddOriginalTransactionReference() *OriginalTransactionReference20 {
	p.OriginalTransactionReference = new(OriginalTransactionReference20)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction58) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

// Provides further details on the original transactions, to which the status report message refers.
type PaymentTransaction59 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reported status.
	// Usage: The instructing party is the party sending the status message and not the party that sent the original instruction that is being reported on.
	StatusIdentification *Max35Text `xml:"StsId,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Specifies the status of a transaction, in a coded form.
	TransactionStatus *TransactionIndividualStatus3Code `xml:"TxSts,omitempty"`

	// Provides detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Provides information on the charges related to the processing of the rejection of the instruction.
	// Usage: This is passed on for information purposes only. Settlement of the charges will be done separately.
	ChargesInformation []*Charges2 `xml:"ChrgsInf,omitempty"`

	// Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.
	AcceptanceDateTime *ISODateTime `xml:"AccptncDtTm,omitempty"`

	// Unique reference, as assigned by the account servicing institution, to unambiguously identify the instruction.
	AccountServicerReference *Max35Text `xml:"AcctSvcrRef,omitempty"`

	// Unique reference, as assigned by a clearing system, to unambiguously identify the instruction.
	ClearingSystemReference *Max35Text `xml:"ClrSysRef,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference21 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction59) SetStatusIdentification(value string) {
	p.StatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction59) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction59) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction59) SetTransactionStatus(value string) {
	p.TransactionStatus = (*TransactionIndividualStatus3Code)(&value)
}

func (p *PaymentTransaction59) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	p.StatusReasonInformation = append(p.StatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction59) AddChargesInformation() *Charges2 {
	newValue := new(Charges2)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransaction59) SetAcceptanceDateTime(value string) {
	p.AcceptanceDateTime = (*ISODateTime)(&value)
}

func (p *PaymentTransaction59) SetAccountServicerReference(value string) {
	p.AccountServicerReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction59) SetClearingSystemReference(value string) {
	p.ClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction59) AddOriginalTransactionReference() *OriginalTransactionReference21 {
	p.OriginalTransactionReference = new(OriginalTransactionReference21)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction59) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

// Provides further details on the reference and status on the original transactions, included in the original instruction, to which the reversal message applies.
type PaymentTransaction60 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reversed transaction.
	// Usage: The instructing party is the party sending the reversal message and not the party that sent the original instruction that is being reversed.
	ReversalIdentification *Max35Text `xml:"RvslId,omitempty"`

	// Provides information on the original message.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Unique reference, as assigned by the original clearing system, to unambiguously identify the original instruction.
	OriginalClearingSystemReference *Max35Text `xml:"OrgnlClrSysRef,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the instructing agent and the instructed agent.
	OriginalInterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent in the reversed transaction.
	ReversedInterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"RvsdIntrBkSttlmAmt"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	//
	// Usage: The InterbankSettlementDate is the interbank settlement date of the reversal message, and not of the original instruction.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the settlement instruction.
	//
	//
	// Usage: the SettlementPriority is the settlement priority of the reversal message, and not of the original instruction.
	SettlementPriority *Priority3Code `xml:"SttlmPrty,omitempty"`

	// Amount of money to be moved between the debtor and the creditor, before deduction of charges, in the reversed transaction.
	// Usage: This amount has to be transported unchanged through the transaction chain.
	ReversedInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"RvsdInstdAmt,omitempty"`

	// Factor used to convert an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Amount of money asked or paid as compensation for the processing of the instruction.
	CompensationAmount *ActiveOrHistoricCurrencyAndAmount `xml:"CompstnAmt,omitempty"`

	// Specifies if the creditor and/or debtor will bear the charges associated with the processing of the payment transaction.
	//
	// Usage: The ChargeBearer applies to the reversal message, not to the original instruction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Provides information on the charges to be paid by the charge bearer(s) related to the processing of the reversal transaction.
	ChargesInformation []*Charges2 `xml:"ChrgsInf,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructing agent is the party sending the reversal message and not the party that sent the original instruction that is being reversed.
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructed agent is the party receiving the reversal message and not the party that received the original instruction that is being reversed.
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`

	// Provides detailed information on the reversal reason.
	ReversalReasonInformation []*PaymentReversalReason7 `xml:"RvslRsnInf,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference22 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction60) SetReversalIdentification(value string) {
	p.ReversalIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction60) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	p.OriginalGroupInformation = new(OriginalGroupInformation3)
	return p.OriginalGroupInformation
}

func (p *PaymentTransaction60) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction60) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction60) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction60) SetOriginalClearingSystemReference(value string) {
	p.OriginalClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction60) SetOriginalInterbankSettlementAmount(value, currency string) {
	p.OriginalInterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction60) SetReversedInterbankSettlementAmount(value, currency string) {
	p.ReversedInterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction60) SetInterbankSettlementDate(value string) {
	p.InterbankSettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction60) SetSettlementPriority(value string) {
	p.SettlementPriority = (*Priority3Code)(&value)
}

func (p *PaymentTransaction60) SetReversedInstructedAmount(value, currency string) {
	p.ReversedInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction60) SetExchangeRate(value string) {
	p.ExchangeRate = (*BaseOneRate)(&value)
}

func (p *PaymentTransaction60) SetCompensationAmount(value, currency string) {
	p.CompensationAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction60) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentTransaction60) AddChargesInformation() *Charges2 {
	newValue := new(Charges2)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransaction60) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructingAgent
}

func (p *PaymentTransaction60) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructedAgent
}

func (p *PaymentTransaction60) AddReversalReasonInformation() *PaymentReversalReason7 {
	newValue := new(PaymentReversalReason7)
	p.ReversalReasonInformation = append(p.ReversalReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction60) AddOriginalTransactionReference() *OriginalTransactionReference22 {
	p.OriginalTransactionReference = new(OriginalTransactionReference22)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction60) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

// Set of elements used to provide reference and status information on the original transactions, included in the original instruction, to which the cancellation request message applies.
type PaymentTransaction61 struct {

	// Unique identification, as assigned by the assigner, to unambiguously identify a cancellation request.
	//
	// Usage: The cancellation request identification can be used for reconciliation or to link tasks relating to the cancellation request.
	CancellationIdentification *Max35Text `xml:"CxlId,omitempty"`

	// Set of elements to uniquely and unambiguously identify an exception or an investigation workflow.
	Case *Case3 `xml:"Case,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the debtor and the creditor, before deduction of charges, expressed in the currency, as ordered by the original initiating party.
	//
	// Usage: This amount has to be transported unchanged through the transaction chain.
	OriginalInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt,omitempty"`

	// Date at which the initiating party originally requested the clearing agent to process the payment.
	OriginalRequestedExecutionDate *ISODate `xml:"OrgnlReqdExctnDt,omitempty"`

	// Date at which the creditor originally requested the collection of the amount of money from the debtor.
	OriginalRequestedCollectionDate *ISODate `xml:"OrgnlReqdColltnDt,omitempty"`

	// Provides detailed information on the cancellation reason.
	CancellationReasonInformation []*PaymentCancellationReason2 `xml:"CxlRsnInf,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference22 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction61) SetCancellationIdentification(value string) {
	p.CancellationIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction61) AddCase() *Case3 {
	p.Case = new(Case3)
	return p.Case
}

func (p *PaymentTransaction61) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction61) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction61) SetOriginalInstructedAmount(value, currency string) {
	p.OriginalInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction61) SetOriginalRequestedExecutionDate(value string) {
	p.OriginalRequestedExecutionDate = (*ISODate)(&value)
}

func (p *PaymentTransaction61) SetOriginalRequestedCollectionDate(value string) {
	p.OriginalRequestedCollectionDate = (*ISODate)(&value)
}

func (p *PaymentTransaction61) AddCancellationReasonInformation() *PaymentCancellationReason2 {
	newValue := new(PaymentCancellationReason2)
	p.CancellationReasonInformation = append(p.CancellationReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction61) AddOriginalTransactionReference() *OriginalTransactionReference22 {
	p.OriginalTransactionReference = new(OriginalTransactionReference22)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction61) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

// Provides reference and status information on the original transactions, included in the original instruction, to which the cancellation request message applies.
type PaymentTransaction62 struct {

	// Unique and unambiguous identifier of a cancellation request, as assigned by the assigner.
	//
	// Usage: The cancellation request identification can be used for reconciliation or to link tasks relating to the cancellation request.
	CancellationIdentification *Max35Text `xml:"CxlId,omitempty"`

	// Set of elements to uniquely and unambiguously identify an exception or an investigation workflow.
	Case *Case3 `xml:"Case,omitempty"`

	// Provides information on the original message.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Unique reference, as assigned by the original clearing system, to unambiguously identify the original instruction.
	OriginalClearingSystemReference *Max35Text `xml:"OrgnlClrSysRef,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the instructing agent and the instructed agent.
	OriginalInterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty"`

	// Date, as provided in the original transaction, on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	OriginalInterbankSettlementDate *ISODate `xml:"OrgnlIntrBkSttlmDt,omitempty"`

	// Party who assigns the case.
	// Usage: This is also the agent that instructs the next party in the chain to carry out the (set of) cancellation request(s).
	Assigner *BranchAndFinancialInstitutionIdentification5 `xml:"Assgnr,omitempty"`

	// Party to which the case is assigned.
	// Usage: This is also the agent that is instructed by the previous party in the chain to carry out the (set of) cancellation request(s).
	Assignee *BranchAndFinancialInstitutionIdentification5 `xml:"Assgne,omitempty"`

	// Provides detailed information on the cancellation reason.
	CancellationReasonInformation []*PaymentCancellationReason2 `xml:"CxlRsnInf,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference22 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction62) SetCancellationIdentification(value string) {
	p.CancellationIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction62) AddCase() *Case3 {
	p.Case = new(Case3)
	return p.Case
}

func (p *PaymentTransaction62) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	p.OriginalGroupInformation = new(OriginalGroupInformation3)
	return p.OriginalGroupInformation
}

func (p *PaymentTransaction62) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction62) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction62) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction62) SetOriginalClearingSystemReference(value string) {
	p.OriginalClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction62) SetOriginalInterbankSettlementAmount(value, currency string) {
	p.OriginalInterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction62) SetOriginalInterbankSettlementDate(value string) {
	p.OriginalInterbankSettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction62) AddAssigner() *BranchAndFinancialInstitutionIdentification5 {
	p.Assigner = new(BranchAndFinancialInstitutionIdentification5)
	return p.Assigner
}

func (p *PaymentTransaction62) AddAssignee() *BranchAndFinancialInstitutionIdentification5 {
	p.Assignee = new(BranchAndFinancialInstitutionIdentification5)
	return p.Assignee
}

func (p *PaymentTransaction62) AddCancellationReasonInformation() *PaymentCancellationReason2 {
	newValue := new(PaymentCancellationReason2)
	p.CancellationReasonInformation = append(p.CancellationReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction62) AddOriginalTransactionReference() *OriginalTransactionReference22 {
	p.OriginalTransactionReference = new(OriginalTransactionReference22)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction62) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

// Provides further details on the original transactions, to which the status report message refers.
type PaymentTransaction63 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reported status.
	// Usage: The instructing party is the party sending the status message and not the party that sent the original instruction that is being reported on.
	StatusIdentification *Max35Text `xml:"StsId,omitempty"`

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Specifies the status of a transaction, in a coded form.
	TransactionStatus *TransactionIndividualStatus3Code `xml:"TxSts,omitempty"`

	// Provides detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Provides information on the charges related to the processing of the rejection of the instruction.
	//
	// Usage: This is passed on for information purposes only. Settlement of the charges will be done separately.
	ChargesInformation []*Charges2 `xml:"ChrgsInf,omitempty"`

	// Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.
	AcceptanceDateTime *ISODateTime `xml:"AccptncDtTm,omitempty"`

	// Unique reference, as assigned by the account servicing institution, to unambiguously identify the instruction.
	AccountServicerReference *Max35Text `xml:"AcctSvcrRef,omitempty"`

	// Unique reference, as assigned by a clearing system, to unambiguously identify the instruction.
	ClearingSystemReference *Max35Text `xml:"ClrSysRef,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructing agent is the party sending the status message and not the party that sent the original instruction that is being reported on.
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructed agent is the party receiving the status message and not the party that received the original instruction that is being reported on.
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference22 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction63) SetStatusIdentification(value string) {
	p.StatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction63) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	p.OriginalGroupInformation = new(OriginalGroupInformation3)
	return p.OriginalGroupInformation
}

func (p *PaymentTransaction63) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction63) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction63) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction63) SetTransactionStatus(value string) {
	p.TransactionStatus = (*TransactionIndividualStatus3Code)(&value)
}

func (p *PaymentTransaction63) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	p.StatusReasonInformation = append(p.StatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction63) AddChargesInformation() *Charges2 {
	newValue := new(Charges2)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransaction63) SetAcceptanceDateTime(value string) {
	p.AcceptanceDateTime = (*ISODateTime)(&value)
}

func (p *PaymentTransaction63) SetAccountServicerReference(value string) {
	p.AccountServicerReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction63) SetClearingSystemReference(value string) {
	p.ClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction63) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructingAgent
}

func (p *PaymentTransaction63) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructedAgent
}

func (p *PaymentTransaction63) AddOriginalTransactionReference() *OriginalTransactionReference22 {
	p.OriginalTransactionReference = new(OriginalTransactionReference22)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction63) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

// Provides further details on the reference and status on the original transactions, included in the original instruction, to which the reversal message applies.
type PaymentTransaction64 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reversed transaction.
	// Usage: The instructing party is the party sending the reversal message and not the party that sent the original instruction that is being reversed.
	ReversalIdentification *Max35Text `xml:"RvslId,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the debtor and the creditor, before deduction of charges, expressed in the currency, as ordered by the original initiating party.
	// Usage: This amount has to be transported unchanged through the transaction chain.
	OriginalInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt,omitempty"`

	// Amount of money to be moved between the debtor and the creditor, before deduction of charges, in the reversed transaction.
	// Usage: This amount has to be transported unchanged through the transaction chain.
	ReversedInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"RvsdInstdAmt,omitempty"`

	// Specifies if the creditor and/or debtor will bear the charges associated with the processing of the payment transaction.
	//
	// Usage: The ChargeBearer applies to the reversal message, not to the original instruction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Provides detailed information on the reversal reason.
	ReversalReasonInformation []*PaymentReversalReason7 `xml:"RvslRsnInf,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference22 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction64) SetReversalIdentification(value string) {
	p.ReversalIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction64) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction64) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction64) SetOriginalInstructedAmount(value, currency string) {
	p.OriginalInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction64) SetReversedInstructedAmount(value, currency string) {
	p.ReversedInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction64) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentTransaction64) AddReversalReasonInformation() *PaymentReversalReason7 {
	newValue := new(PaymentReversalReason7)
	p.ReversalReasonInformation = append(p.ReversalReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction64) AddOriginalTransactionReference() *OriginalTransactionReference22 {
	p.OriginalTransactionReference = new(OriginalTransactionReference22)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction64) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

// Provides further details on the reference and status on the original transactions, included in the original instruction, to which the return message applies.
type PaymentTransaction65 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the returned transaction.
	// Usage: The instructing party is the party sending the return message and not the party that sent the original instruction that is being returned.
	ReturnIdentification *Max35Text `xml:"RtrId,omitempty"`

	// Provides information on the original message.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Unique reference, as assigned by the original clearing system, to unambiguously identify the original instruction.
	OriginalClearingSystemReference *Max35Text `xml:"OrgnlClrSysRef,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the instructing agent and the instructed agent.
	OriginalInterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent in the returned transaction.
	ReturnedInterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"RtrdIntrBkSttlmAmt"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	//
	// Usage: the InterbankSettlementDate is the interbank settlement date of the return message, and not of the original instruction.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the settlement instruction.
	//
	//
	// Usage: the SettlementPriority is the settlement priority of the return message, and not of the original instruction.
	SettlementPriority *Priority3Code `xml:"SttlmPrty,omitempty"`

	// Amount of money to be moved between the debtor and the creditor, before deduction of charges, in the returned transaction.
	// Usage: This amount has to be transported unchanged through the transaction chain.
	ReturnedInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"RtrdInstdAmt,omitempty"`

	// Factor used to convert an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Amount of money asked or paid as compensation for the processing of the instruction.
	CompensationAmount *ActiveOrHistoricCurrencyAndAmount `xml:"CompstnAmt,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	//
	// Usage: The ChargeBearer applies to the return message, not to the original instruction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Provides information on the charges to be paid by the charge bearer(s) related to the processing of the return transaction.
	ChargesInformation []*Charges2 `xml:"ChrgsInf,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructing agent is the party sending the return message and not the party that sent the original instruction that is being returned.
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructed agent is the party receiving the return message and not the party that received the original instruction that is being returned.
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`

	// Provides detailed information on the return reason.
	ReturnReasonInformation []*PaymentReturnReason1 `xml:"RtrRsnInf,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference22 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction65) SetReturnIdentification(value string) {
	p.ReturnIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction65) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	p.OriginalGroupInformation = new(OriginalGroupInformation3)
	return p.OriginalGroupInformation
}

func (p *PaymentTransaction65) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction65) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction65) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction65) SetOriginalClearingSystemReference(value string) {
	p.OriginalClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction65) SetOriginalInterbankSettlementAmount(value, currency string) {
	p.OriginalInterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction65) SetReturnedInterbankSettlementAmount(value, currency string) {
	p.ReturnedInterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction65) SetInterbankSettlementDate(value string) {
	p.InterbankSettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction65) SetSettlementPriority(value string) {
	p.SettlementPriority = (*Priority3Code)(&value)
}

func (p *PaymentTransaction65) SetReturnedInstructedAmount(value, currency string) {
	p.ReturnedInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction65) SetExchangeRate(value string) {
	p.ExchangeRate = (*BaseOneRate)(&value)
}

func (p *PaymentTransaction65) SetCompensationAmount(value, currency string) {
	p.CompensationAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction65) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentTransaction65) AddChargesInformation() *Charges2 {
	newValue := new(Charges2)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransaction65) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructingAgent
}

func (p *PaymentTransaction65) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructedAgent
}

func (p *PaymentTransaction65) AddReturnReasonInformation() *PaymentReturnReason1 {
	newValue := new(PaymentReturnReason1)
	p.ReturnReasonInformation = append(p.ReturnReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction65) AddOriginalTransactionReference() *OriginalTransactionReference22 {
	p.OriginalTransactionReference = new(OriginalTransactionReference22)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction65) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

// Provides reference and status information on the original transactions, included in the original instruction, to which the cancellation request message applies.
type PaymentTransaction66 struct {

	// Unique and unambiguous identifier of a cancellation request status, as assigned by the assigner.
	//
	// Usage: The cancellation status identification can be used for reconciliation or to link tasks relating to the cancellation request.
	CancellationStatusIdentification *Max35Text `xml:"CxlStsId,omitempty"`

	// Identifies the resolved case.
	ResolvedCase *Case3 `xml:"RslvdCase,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Specifies the status of the transaction cancellation request.
	TransactionCancellationStatus *CancellationIndividualStatus1Code `xml:"TxCxlSts,omitempty"`

	// Provides detailed information on the cancellation status reason.
	CancellationStatusReasonInformation []*CancellationStatusReason2 `xml:"CxlStsRsnInf,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the debtor and the creditor, before deduction of charges, expressed in the currency, as ordered by the original initiating party.
	// Usage: This amount has to be transported unchanged through the transaction chain.
	OriginalInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt,omitempty"`

	// Date at which the initiating party originally requested the clearing agent to process the payment.
	OriginalRequestedExecutionDate *ISODate `xml:"OrgnlReqdExctnDt,omitempty"`

	// Date at which the creditor originally requested the collection of the amount of money from the debtor.
	OriginalRequestedCollectionDate *ISODate `xml:"OrgnlReqdColltnDt,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference22 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransaction66) SetCancellationStatusIdentification(value string) {
	p.CancellationStatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction66) AddResolvedCase() *Case3 {
	p.ResolvedCase = new(Case3)
	return p.ResolvedCase
}

func (p *PaymentTransaction66) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction66) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction66) SetTransactionCancellationStatus(value string) {
	p.TransactionCancellationStatus = (*CancellationIndividualStatus1Code)(&value)
}

func (p *PaymentTransaction66) AddCancellationStatusReasonInformation() *CancellationStatusReason2 {
	newValue := new(CancellationStatusReason2)
	p.CancellationStatusReasonInformation = append(p.CancellationStatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction66) SetOriginalInstructedAmount(value, currency string) {
	p.OriginalInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction66) SetOriginalRequestedExecutionDate(value string) {
	p.OriginalRequestedExecutionDate = (*ISODate)(&value)
}

func (p *PaymentTransaction66) SetOriginalRequestedCollectionDate(value string) {
	p.OriginalRequestedCollectionDate = (*ISODate)(&value)
}

func (p *PaymentTransaction66) AddOriginalTransactionReference() *OriginalTransactionReference22 {
	p.OriginalTransactionReference = new(OriginalTransactionReference22)
	return p.OriginalTransactionReference
}

// Provides further details on the reference and status on the original transactions, included in the original instruction, to which the cancellation request message applies.
type PaymentTransaction67 struct {

	// Unique and unambiguous identifier of a cancellation request status, as assigned by the assigner.
	//
	// Usage: The cancellation status identification can be used for reconciliation or to link tasks relating to the cancellation request.
	CancellationStatusIdentification *Max35Text `xml:"CxlStsId,omitempty"`

	// Identifies the resolved case.
	ResolvedCase *Case3 `xml:"RslvdCase,omitempty"`

	// Provides information on the original message.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Unique reference, as assigned by the original clearing system, to unambiguously identify the original instruction.
	OriginalClearingSystemReference *Max35Text `xml:"OrgnlClrSysRef,omitempty"`

	// Specifies the status of the transaction cancellation request.
	TransactionCancellationStatus *CancellationIndividualStatus1Code `xml:"TxCxlSts,omitempty"`

	// Provides detailed information on the cancellation status reason.
	CancellationStatusReasonInformation []*CancellationStatusReason2 `xml:"CxlStsRsnInf,omitempty"`

	// Reference of a return or a reversal transaction that is initiated to fix the case under investigation as part of the resolution.
	ResolutionRelatedInformation *ResolutionInformation1 `xml:"RsltnRltdInf,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the instructing agent and the instructed agent.
	OriginalInterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty"`

	// Date, as provided in the original transaction, on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	OriginalInterbankSettlementDate *ISODate `xml:"OrgnlIntrBkSttlmDt,omitempty"`

	// Party who assigns the case.
	// Usage: This is also the agent that instructs the next party in the chain to carry out the (set of) cancellation request(s).
	Assigner *Party12Choice `xml:"Assgnr,omitempty"`

	// Party to which the case is assigned.
	// Usage: This is also the agent that is instructed by the previous party in the chain to carry out the (set of) cancellation request(s).
	Assignee *Party12Choice `xml:"Assgne,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference22 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransaction67) SetCancellationStatusIdentification(value string) {
	p.CancellationStatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction67) AddResolvedCase() *Case3 {
	p.ResolvedCase = new(Case3)
	return p.ResolvedCase
}

func (p *PaymentTransaction67) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	p.OriginalGroupInformation = new(OriginalGroupInformation3)
	return p.OriginalGroupInformation
}

func (p *PaymentTransaction67) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction67) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction67) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction67) SetOriginalClearingSystemReference(value string) {
	p.OriginalClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction67) SetTransactionCancellationStatus(value string) {
	p.TransactionCancellationStatus = (*CancellationIndividualStatus1Code)(&value)
}

func (p *PaymentTransaction67) AddCancellationStatusReasonInformation() *CancellationStatusReason2 {
	newValue := new(CancellationStatusReason2)
	p.CancellationStatusReasonInformation = append(p.CancellationStatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction67) AddResolutionRelatedInformation() *ResolutionInformation1 {
	p.ResolutionRelatedInformation = new(ResolutionInformation1)
	return p.ResolutionRelatedInformation
}

func (p *PaymentTransaction67) SetOriginalInterbankSettlementAmount(value, currency string) {
	p.OriginalInterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction67) SetOriginalInterbankSettlementDate(value string) {
	p.OriginalInterbankSettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction67) AddAssigner() *Party12Choice {
	p.Assigner = new(Party12Choice)
	return p.Assigner
}

func (p *PaymentTransaction67) AddAssignee() *Party12Choice {
	p.Assignee = new(Party12Choice)
	return p.Assignee
}

func (p *PaymentTransaction67) AddOriginalTransactionReference() *OriginalTransactionReference22 {
	p.OriginalTransactionReference = new(OriginalTransactionReference22)
	return p.OriginalTransactionReference
}

// Provides further details on the original transactions, to which the status report message refers.
type PaymentTransaction68 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reported status.
	// Usage: The instructing party is the party sending the status message and not the party that sent the original instruction that is being reported on.
	StatusIdentification *Max35Text `xml:"StsId,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Specifies the status of a transaction, in a coded form.
	TransactionStatus *TransactionIndividualStatus3Code `xml:"TxSts,omitempty"`

	// Provides detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Provides information on the charges related to the processing of the rejection of the instruction.
	// Usage: This is passed on for information purposes only. Settlement of the charges will be done separately.
	ChargesInformation []*Charges2 `xml:"ChrgsInf,omitempty"`

	// Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.
	AcceptanceDateTime *ISODateTime `xml:"AccptncDtTm,omitempty"`

	// Unique reference, as assigned by the account servicing institution, to unambiguously identify the instruction.
	AccountServicerReference *Max35Text `xml:"AcctSvcrRef,omitempty"`

	// Unique reference, as assigned by a clearing system, to unambiguously identify the instruction.
	ClearingSystemReference *Max35Text `xml:"ClrSysRef,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference22 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction68) SetStatusIdentification(value string) {
	p.StatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction68) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction68) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction68) SetTransactionStatus(value string) {
	p.TransactionStatus = (*TransactionIndividualStatus3Code)(&value)
}

func (p *PaymentTransaction68) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	p.StatusReasonInformation = append(p.StatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction68) AddChargesInformation() *Charges2 {
	newValue := new(Charges2)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransaction68) SetAcceptanceDateTime(value string) {
	p.AcceptanceDateTime = (*ISODateTime)(&value)
}

func (p *PaymentTransaction68) SetAccountServicerReference(value string) {
	p.AccountServicerReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction68) SetClearingSystemReference(value string) {
	p.ClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction68) AddOriginalTransactionReference() *OriginalTransactionReference22 {
	p.OriginalTransactionReference = new(OriginalTransactionReference22)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction68) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

// Provides further details on the original transactions, to which the status report message refers.
type PaymentTransaction69 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reported status.
	// Usage: The instructing party is the party sending the status message and not the party that sent the original instruction that is being reported on.
	StatusIdentification *Max35Text `xml:"StsId,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Specifies the status of a transaction, in a coded form.
	TransactionStatus *TransactionIndividualStatus3Code `xml:"TxSts,omitempty"`

	// Provides detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Provides information on the charges related to the processing of the rejection of the instruction.
	// Usage: This is passed on for information purposes only. Settlement of the charges will be done separately.
	ChargesInformation []*Charges2 `xml:"ChrgsInf,omitempty"`

	// Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.
	AcceptanceDateTime *ISODateTime `xml:"AccptncDtTm,omitempty"`

	// Unique reference, as assigned by the account servicing institution, to unambiguously identify the instruction.
	AccountServicerReference *Max35Text `xml:"AcctSvcrRef,omitempty"`

	// Unique reference, as assigned by a clearing system, to unambiguously identify the instruction.
	ClearingSystemReference *Max35Text `xml:"ClrSysRef,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference23 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction69) SetStatusIdentification(value string) {
	p.StatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction69) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction69) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction69) SetTransactionStatus(value string) {
	p.TransactionStatus = (*TransactionIndividualStatus3Code)(&value)
}

func (p *PaymentTransaction69) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	p.StatusReasonInformation = append(p.StatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction69) AddChargesInformation() *Charges2 {
	newValue := new(Charges2)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransaction69) SetAcceptanceDateTime(value string) {
	p.AcceptanceDateTime = (*ISODateTime)(&value)
}

func (p *PaymentTransaction69) SetAccountServicerReference(value string) {
	p.AccountServicerReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction69) SetClearingSystemReference(value string) {
	p.ClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction69) AddOriginalTransactionReference() *OriginalTransactionReference23 {
	p.OriginalTransactionReference = new(OriginalTransactionReference23)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction69) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

// Payment processes required to transfer cash from the debtor to the creditor.
type PaymentTransaction70 struct {

	// Choice between types of payment instrument, for example, cheque, credit transfer, direct debit, investment account or payment card.
	PaymentInstrument *PaymentInstrument20Choice `xml:"PmtInstrm"`
}

func (p *PaymentTransaction70) AddPaymentInstrument() *PaymentInstrument20Choice {
	p.PaymentInstrument = new(PaymentInstrument20Choice)
	return p.PaymentInstrument
}

// Payment processes required to transfer cash from the debtor to the creditor.
type PaymentTransaction71 struct {

	// Choice between cash-in or cash-out.
	CashInOrOut *CashInOrOut7Choice `xml:"CshInOrOut"`
}

func (p *PaymentTransaction71) AddCashInOrOut() *CashInOrOut7Choice {
	p.CashInOrOut = new(CashInOrOut7Choice)
	return p.CashInOrOut
}

// Payment processes required to transfer cash from the debtor to the creditor.
type PaymentTransaction72 struct {

	// Choice between types of payment instrument, for example, cheque, credit transfer or investment account.
	PaymentInstrument *PaymentInstrument21Choice `xml:"PmtInstrm"`
}

func (p *PaymentTransaction72) AddPaymentInstrument() *PaymentInstrument21Choice {
	p.PaymentInstrument = new(PaymentInstrument21Choice)
	return p.PaymentInstrument
}

// Provides further details on the original transactions, to which the status report message refers.
type PaymentTransaction73 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the status request.
	// Usage: The instructing party is the party sending the request message and not the party that sent the original instruction that is being reported on.
	StatusRequestIdentification *Max35Text `xml:"StsReqId,omitempty"`

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.
	AcceptanceDateTime *ISODateTime `xml:"AccptncDtTm,omitempty"`

	// Unique reference, as assigned by a clearing system, to unambiguously identify the instruction.
	ClearingSystemReference *Max35Text `xml:"ClrSysRef,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	//
	//
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	//
	//
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference24 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction73) SetStatusRequestIdentification(value string) {
	p.StatusRequestIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction73) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	p.OriginalGroupInformation = new(OriginalGroupInformation3)
	return p.OriginalGroupInformation
}

func (p *PaymentTransaction73) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction73) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction73) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction73) SetAcceptanceDateTime(value string) {
	p.AcceptanceDateTime = (*ISODateTime)(&value)
}

func (p *PaymentTransaction73) SetClearingSystemReference(value string) {
	p.ClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction73) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructingAgent
}

func (p *PaymentTransaction73) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructedAgent
}

func (p *PaymentTransaction73) AddOriginalTransactionReference() *OriginalTransactionReference24 {
	p.OriginalTransactionReference = new(OriginalTransactionReference24)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction73) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

// Set of elements used to provide reference and status information on the original transactions, included in the original instruction, to which the cancellation request message applies.
type PaymentTransaction74 struct {

	// Unique identification, as assigned by the assigner, to unambiguously identify a cancellation request.
	//
	// Usage: The cancellation request identification can be used for reconciliation or to link tasks relating to the cancellation request.
	CancellationIdentification *Max35Text `xml:"CxlId,omitempty"`

	// Set of elements to uniquely and unambiguously identify an exception or an investigation workflow.
	Case *Case3 `xml:"Case,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the debtor and the creditor, before deduction of charges, expressed in the currency, as ordered by the original initiating party.
	//
	// Usage: This amount has to be transported unchanged through the transaction chain.
	OriginalInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt,omitempty"`

	// Date at which the initiating party originally requested the clearing agent to process the payment.
	OriginalRequestedExecutionDate *DateAndDateTimeChoice `xml:"OrgnlReqdExctnDt,omitempty"`

	// Date at which the creditor originally requested the collection of the amount of money from the debtor.
	OriginalRequestedCollectionDate *ISODate `xml:"OrgnlReqdColltnDt,omitempty"`

	// Provides detailed information on the cancellation reason.
	CancellationReasonInformation []*PaymentCancellationReason3 `xml:"CxlRsnInf,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference24 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction74) SetCancellationIdentification(value string) {
	p.CancellationIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction74) AddCase() *Case3 {
	p.Case = new(Case3)
	return p.Case
}

func (p *PaymentTransaction74) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction74) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction74) SetOriginalInstructedAmount(value, currency string) {
	p.OriginalInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction74) AddOriginalRequestedExecutionDate() *DateAndDateTimeChoice {
	p.OriginalRequestedExecutionDate = new(DateAndDateTimeChoice)
	return p.OriginalRequestedExecutionDate
}

func (p *PaymentTransaction74) SetOriginalRequestedCollectionDate(value string) {
	p.OriginalRequestedCollectionDate = (*ISODate)(&value)
}

func (p *PaymentTransaction74) AddCancellationReasonInformation() *PaymentCancellationReason3 {
	newValue := new(PaymentCancellationReason3)
	p.CancellationReasonInformation = append(p.CancellationReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction74) AddOriginalTransactionReference() *OriginalTransactionReference24 {
	p.OriginalTransactionReference = new(OriginalTransactionReference24)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction74) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

// Provides reference and status information on the original transactions, included in the original instruction, to which the cancellation request message applies.
type PaymentTransaction75 struct {

	// Unique and unambiguous identifier of a cancellation request, as assigned by the assigner.
	//
	// Usage: The cancellation request identification can be used for reconciliation or to link tasks relating to the cancellation request.
	CancellationIdentification *Max35Text `xml:"CxlId,omitempty"`

	// Set of elements to uniquely and unambiguously identify an exception or an investigation workflow.
	Case *Case3 `xml:"Case,omitempty"`

	// Provides information on the original message.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Unique reference, as assigned by the original clearing system, to unambiguously identify the original instruction.
	OriginalClearingSystemReference *Max35Text `xml:"OrgnlClrSysRef,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the instructing agent and the instructed agent.
	OriginalInterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty"`

	// Date, as provided in the original transaction, on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	OriginalInterbankSettlementDate *ISODate `xml:"OrgnlIntrBkSttlmDt,omitempty"`

	// Party who assigns the case.
	// Usage: This is also the agent that instructs the next party in the chain to carry out the (set of) cancellation request(s).
	Assigner *BranchAndFinancialInstitutionIdentification5 `xml:"Assgnr,omitempty"`

	// Party to which the case is assigned.
	// Usage: This is also the agent that is instructed by the previous party in the chain to carry out the (set of) cancellation request(s).
	Assignee *BranchAndFinancialInstitutionIdentification5 `xml:"Assgne,omitempty"`

	// Provides detailed information on the cancellation reason.
	CancellationReasonInformation []*PaymentCancellationReason3 `xml:"CxlRsnInf,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference24 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction75) SetCancellationIdentification(value string) {
	p.CancellationIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction75) AddCase() *Case3 {
	p.Case = new(Case3)
	return p.Case
}

func (p *PaymentTransaction75) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	p.OriginalGroupInformation = new(OriginalGroupInformation3)
	return p.OriginalGroupInformation
}

func (p *PaymentTransaction75) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction75) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction75) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction75) SetOriginalClearingSystemReference(value string) {
	p.OriginalClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction75) SetOriginalInterbankSettlementAmount(value, currency string) {
	p.OriginalInterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction75) SetOriginalInterbankSettlementDate(value string) {
	p.OriginalInterbankSettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction75) AddAssigner() *BranchAndFinancialInstitutionIdentification5 {
	p.Assigner = new(BranchAndFinancialInstitutionIdentification5)
	return p.Assigner
}

func (p *PaymentTransaction75) AddAssignee() *BranchAndFinancialInstitutionIdentification5 {
	p.Assignee = new(BranchAndFinancialInstitutionIdentification5)
	return p.Assignee
}

func (p *PaymentTransaction75) AddCancellationReasonInformation() *PaymentCancellationReason3 {
	newValue := new(PaymentCancellationReason3)
	p.CancellationReasonInformation = append(p.CancellationReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction75) AddOriginalTransactionReference() *OriginalTransactionReference24 {
	p.OriginalTransactionReference = new(OriginalTransactionReference24)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction75) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

// Provides further details on the reference and status on the original transactions, included in the original instruction, to which the return message applies.
type PaymentTransaction76 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the returned transaction.
	// Usage: The instructing party is the party sending the return message and not the party that sent the original instruction that is being returned.
	ReturnIdentification *Max35Text `xml:"RtrId,omitempty"`

	// Provides information on the original message.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Unique reference, as assigned by the original clearing system, to unambiguously identify the original instruction.
	OriginalClearingSystemReference *Max35Text `xml:"OrgnlClrSysRef,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the instructing agent and the instructed agent.
	OriginalInterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent in the returned transaction.
	ReturnedInterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"RtrdIntrBkSttlmAmt"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	//
	// Usage: the InterbankSettlementDate is the interbank settlement date of the return message, and not of the original instruction.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the settlement instruction.
	//
	//
	// Usage: the SettlementPriority is the settlement priority of the return message, and not of the original instruction.
	SettlementPriority *Priority3Code `xml:"SttlmPrty,omitempty"`

	// Amount of money to be moved between the debtor and the creditor, before deduction of charges, in the returned transaction.
	// Usage: This amount has to be transported unchanged through the transaction chain.
	ReturnedInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"RtrdInstdAmt,omitempty"`

	// Factor used to convert an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Amount of money asked or paid as compensation for the processing of the instruction.
	CompensationAmount *ActiveOrHistoricCurrencyAndAmount `xml:"CompstnAmt,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	//
	// Usage: The ChargeBearer applies to the return message, not to the original instruction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Provides information on the charges to be paid by the charge bearer(s) related to the processing of the return transaction.
	ChargesInformation []*Charges2 `xml:"ChrgsInf,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructing agent is the party sending the return message and not the party that sent the original instruction that is being returned.
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructed agent is the party receiving the return message and not the party that received the original instruction that is being returned.
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`

	// Provides detailed information on the return reason.
	ReturnReasonInformation []*PaymentReturnReason1 `xml:"RtrRsnInf,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference24 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction76) SetReturnIdentification(value string) {
	p.ReturnIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction76) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	p.OriginalGroupInformation = new(OriginalGroupInformation3)
	return p.OriginalGroupInformation
}

func (p *PaymentTransaction76) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction76) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction76) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction76) SetOriginalClearingSystemReference(value string) {
	p.OriginalClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction76) SetOriginalInterbankSettlementAmount(value, currency string) {
	p.OriginalInterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction76) SetReturnedInterbankSettlementAmount(value, currency string) {
	p.ReturnedInterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction76) SetInterbankSettlementDate(value string) {
	p.InterbankSettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction76) SetSettlementPriority(value string) {
	p.SettlementPriority = (*Priority3Code)(&value)
}

func (p *PaymentTransaction76) SetReturnedInstructedAmount(value, currency string) {
	p.ReturnedInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction76) SetExchangeRate(value string) {
	p.ExchangeRate = (*BaseOneRate)(&value)
}

func (p *PaymentTransaction76) SetCompensationAmount(value, currency string) {
	p.CompensationAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction76) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentTransaction76) AddChargesInformation() *Charges2 {
	newValue := new(Charges2)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransaction76) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructingAgent
}

func (p *PaymentTransaction76) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructedAgent
}

func (p *PaymentTransaction76) AddReturnReasonInformation() *PaymentReturnReason1 {
	newValue := new(PaymentReturnReason1)
	p.ReturnReasonInformation = append(p.ReturnReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction76) AddOriginalTransactionReference() *OriginalTransactionReference24 {
	p.OriginalTransactionReference = new(OriginalTransactionReference24)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction76) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

// Provides further details on the reference and status on the original transactions, included in the original instruction, to which the reversal message applies.
type PaymentTransaction77 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reversed transaction.
	// Usage: The instructing party is the party sending the reversal message and not the party that sent the original instruction that is being reversed.
	ReversalIdentification *Max35Text `xml:"RvslId,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the debtor and the creditor, before deduction of charges, expressed in the currency, as ordered by the original initiating party.
	// Usage: This amount has to be transported unchanged through the transaction chain.
	OriginalInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt,omitempty"`

	// Amount of money to be moved between the debtor and the creditor, before deduction of charges, in the reversed transaction.
	// Usage: This amount has to be transported unchanged through the transaction chain.
	ReversedInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"RvsdInstdAmt,omitempty"`

	// Specifies if the creditor and/or debtor will bear the charges associated with the processing of the payment transaction.
	//
	// Usage: The ChargeBearer applies to the reversal message, not to the original instruction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Provides detailed information on the reversal reason.
	ReversalReasonInformation []*PaymentReversalReason7 `xml:"RvslRsnInf,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference24 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction77) SetReversalIdentification(value string) {
	p.ReversalIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction77) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction77) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction77) SetOriginalInstructedAmount(value, currency string) {
	p.OriginalInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction77) SetReversedInstructedAmount(value, currency string) {
	p.ReversedInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction77) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentTransaction77) AddReversalReasonInformation() *PaymentReversalReason7 {
	newValue := new(PaymentReversalReason7)
	p.ReversalReasonInformation = append(p.ReversalReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction77) AddOriginalTransactionReference() *OriginalTransactionReference24 {
	p.OriginalTransactionReference = new(OriginalTransactionReference24)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction77) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

// Provides reference and status information on the original transactions, included in the original instruction, to which the cancellation request message applies.
type PaymentTransaction78 struct {

	// Unique and unambiguous identifier of a cancellation request status, as assigned by the assigner.
	//
	// Usage: The cancellation status identification can be used for reconciliation or to link tasks relating to the cancellation request.
	CancellationStatusIdentification *Max35Text `xml:"CxlStsId,omitempty"`

	// Identifies the resolved case.
	ResolvedCase *Case3 `xml:"RslvdCase,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Specifies the status of the transaction cancellation request.
	TransactionCancellationStatus *CancellationIndividualStatus1Code `xml:"TxCxlSts,omitempty"`

	// Provides detailed information on the cancellation status reason.
	CancellationStatusReasonInformation []*CancellationStatusReason2 `xml:"CxlStsRsnInf,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the debtor and the creditor, before deduction of charges, expressed in the currency, as ordered by the original initiating party.
	// Usage: This amount has to be transported unchanged through the transaction chain.
	OriginalInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt,omitempty"`

	// Date at which the initiating party originally requested the clearing agent to process the payment.
	OriginalRequestedExecutionDate *DateAndDateTimeChoice `xml:"OrgnlReqdExctnDt,omitempty"`

	// Date at which the creditor originally requested the collection of the amount of money from the debtor.
	OriginalRequestedCollectionDate *ISODate `xml:"OrgnlReqdColltnDt,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference24 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransaction78) SetCancellationStatusIdentification(value string) {
	p.CancellationStatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction78) AddResolvedCase() *Case3 {
	p.ResolvedCase = new(Case3)
	return p.ResolvedCase
}

func (p *PaymentTransaction78) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction78) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction78) SetTransactionCancellationStatus(value string) {
	p.TransactionCancellationStatus = (*CancellationIndividualStatus1Code)(&value)
}

func (p *PaymentTransaction78) AddCancellationStatusReasonInformation() *CancellationStatusReason2 {
	newValue := new(CancellationStatusReason2)
	p.CancellationStatusReasonInformation = append(p.CancellationStatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction78) SetOriginalInstructedAmount(value, currency string) {
	p.OriginalInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction78) AddOriginalRequestedExecutionDate() *DateAndDateTimeChoice {
	p.OriginalRequestedExecutionDate = new(DateAndDateTimeChoice)
	return p.OriginalRequestedExecutionDate
}

func (p *PaymentTransaction78) SetOriginalRequestedCollectionDate(value string) {
	p.OriginalRequestedCollectionDate = (*ISODate)(&value)
}

func (p *PaymentTransaction78) AddOriginalTransactionReference() *OriginalTransactionReference24 {
	p.OriginalTransactionReference = new(OriginalTransactionReference24)
	return p.OriginalTransactionReference
}

// Provides further details on the reference and status on the original transactions, included in the original instruction, to which the cancellation request message applies.
type PaymentTransaction79 struct {

	// Unique and unambiguous identifier of a cancellation request status, as assigned by the assigner.
	//
	// Usage: The cancellation status identification can be used for reconciliation or to link tasks relating to the cancellation request.
	CancellationStatusIdentification *Max35Text `xml:"CxlStsId,omitempty"`

	// Identifies the resolved case.
	ResolvedCase *Case3 `xml:"RslvdCase,omitempty"`

	// Provides information on the original message.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Unique reference, as assigned by the original clearing system, to unambiguously identify the original instruction.
	OriginalClearingSystemReference *Max35Text `xml:"OrgnlClrSysRef,omitempty"`

	// Specifies the status of the transaction cancellation request.
	TransactionCancellationStatus *CancellationIndividualStatus1Code `xml:"TxCxlSts,omitempty"`

	// Provides detailed information on the cancellation status reason.
	CancellationStatusReasonInformation []*CancellationStatusReason2 `xml:"CxlStsRsnInf,omitempty"`

	// Reference of a return or a reversal transaction that is initiated to fix the case under investigation as part of the resolution.
	ResolutionRelatedInformation *ResolutionInformation1 `xml:"RsltnRltdInf,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the instructing agent and the instructed agent.
	OriginalInterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty"`

	// Date, as provided in the original transaction, on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	OriginalInterbankSettlementDate *ISODate `xml:"OrgnlIntrBkSttlmDt,omitempty"`

	// Party who assigns the case.
	// Usage: This is also the agent that instructs the next party in the chain to carry out the (set of) cancellation request(s).
	Assigner *Party12Choice `xml:"Assgnr,omitempty"`

	// Party to which the case is assigned.
	// Usage: This is also the agent that is instructed by the previous party in the chain to carry out the (set of) cancellation request(s).
	Assignee *Party12Choice `xml:"Assgne,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference24 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransaction79) SetCancellationStatusIdentification(value string) {
	p.CancellationStatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction79) AddResolvedCase() *Case3 {
	p.ResolvedCase = new(Case3)
	return p.ResolvedCase
}

func (p *PaymentTransaction79) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	p.OriginalGroupInformation = new(OriginalGroupInformation3)
	return p.OriginalGroupInformation
}

func (p *PaymentTransaction79) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction79) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction79) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction79) SetOriginalClearingSystemReference(value string) {
	p.OriginalClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction79) SetTransactionCancellationStatus(value string) {
	p.TransactionCancellationStatus = (*CancellationIndividualStatus1Code)(&value)
}

func (p *PaymentTransaction79) AddCancellationStatusReasonInformation() *CancellationStatusReason2 {
	newValue := new(CancellationStatusReason2)
	p.CancellationStatusReasonInformation = append(p.CancellationStatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction79) AddResolutionRelatedInformation() *ResolutionInformation1 {
	p.ResolutionRelatedInformation = new(ResolutionInformation1)
	return p.ResolutionRelatedInformation
}

func (p *PaymentTransaction79) SetOriginalInterbankSettlementAmount(value, currency string) {
	p.OriginalInterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction79) SetOriginalInterbankSettlementDate(value string) {
	p.OriginalInterbankSettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction79) AddAssigner() *Party12Choice {
	p.Assigner = new(Party12Choice)
	return p.Assigner
}

func (p *PaymentTransaction79) AddAssignee() *Party12Choice {
	p.Assignee = new(Party12Choice)
	return p.Assignee
}

func (p *PaymentTransaction79) AddOriginalTransactionReference() *OriginalTransactionReference24 {
	p.OriginalTransactionReference = new(OriginalTransactionReference24)
	return p.OriginalTransactionReference
}

// Provides further details on the original transactions, to which the status report message refers.
type PaymentTransaction80 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reported status.
	// Usage: The instructing party is the party sending the status message and not the party that sent the original instruction that is being reported on.
	StatusIdentification *Max35Text `xml:"StsId,omitempty"`

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Specifies the status of a transaction, in a coded form.
	TransactionStatus *ExternalPaymentTransactionStatus1Code `xml:"TxSts,omitempty"`

	// Provides detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Provides information on the charges related to the processing of the rejection of the instruction.
	//
	// Usage: This is passed on for information purposes only. Settlement of the charges will be done separately.
	ChargesInformation []*Charges2 `xml:"ChrgsInf,omitempty"`

	// Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.
	AcceptanceDateTime *ISODateTime `xml:"AccptncDtTm,omitempty"`

	// Unique reference, as assigned by the account servicing institution, to unambiguously identify the instruction.
	AccountServicerReference *Max35Text `xml:"AcctSvcrRef,omitempty"`

	// Unique reference, as assigned by a clearing system, to unambiguously identify the instruction.
	ClearingSystemReference *Max35Text `xml:"ClrSysRef,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructing agent is the party sending the status message and not the party that sent the original instruction that is being reported on.
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructed agent is the party receiving the status message and not the party that received the original instruction that is being reported on.
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference24 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction80) SetStatusIdentification(value string) {
	p.StatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction80) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	p.OriginalGroupInformation = new(OriginalGroupInformation3)
	return p.OriginalGroupInformation
}

func (p *PaymentTransaction80) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction80) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction80) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction80) SetTransactionStatus(value string) {
	p.TransactionStatus = (*ExternalPaymentTransactionStatus1Code)(&value)
}

func (p *PaymentTransaction80) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	p.StatusReasonInformation = append(p.StatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction80) AddChargesInformation() *Charges2 {
	newValue := new(Charges2)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransaction80) SetAcceptanceDateTime(value string) {
	p.AcceptanceDateTime = (*ISODateTime)(&value)
}

func (p *PaymentTransaction80) SetAccountServicerReference(value string) {
	p.AccountServicerReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction80) SetClearingSystemReference(value string) {
	p.ClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction80) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructingAgent
}

func (p *PaymentTransaction80) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructedAgent
}

func (p *PaymentTransaction80) AddOriginalTransactionReference() *OriginalTransactionReference24 {
	p.OriginalTransactionReference = new(OriginalTransactionReference24)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction80) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

// Provides further details on the reference and status on the original transactions, included in the original instruction, to which the reversal message applies.
type PaymentTransaction81 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reversed transaction.
	// Usage: The instructing party is the party sending the reversal message and not the party that sent the original instruction that is being reversed.
	ReversalIdentification *Max35Text `xml:"RvslId,omitempty"`

	// Provides information on the original message.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Unique reference, as assigned by the original clearing system, to unambiguously identify the original instruction.
	OriginalClearingSystemReference *Max35Text `xml:"OrgnlClrSysRef,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the instructing agent and the instructed agent.
	OriginalInterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent in the reversed transaction.
	ReversedInterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"RvsdIntrBkSttlmAmt"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	//
	// Usage: The InterbankSettlementDate is the interbank settlement date of the reversal message, and not of the original instruction.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the settlement instruction.
	//
	//
	// Usage: the SettlementPriority is the settlement priority of the reversal message, and not of the original instruction.
	SettlementPriority *Priority3Code `xml:"SttlmPrty,omitempty"`

	// Amount of money to be moved between the debtor and the creditor, before deduction of charges, in the reversed transaction.
	// Usage: This amount has to be transported unchanged through the transaction chain.
	ReversedInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"RvsdInstdAmt,omitempty"`

	// Factor used to convert an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Amount of money asked or paid as compensation for the processing of the instruction.
	CompensationAmount *ActiveOrHistoricCurrencyAndAmount `xml:"CompstnAmt,omitempty"`

	// Specifies if the creditor and/or debtor will bear the charges associated with the processing of the payment transaction.
	//
	// Usage: The ChargeBearer applies to the reversal message, not to the original instruction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Provides information on the charges to be paid by the charge bearer(s) related to the processing of the reversal transaction.
	ChargesInformation []*Charges2 `xml:"ChrgsInf,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructing agent is the party sending the reversal message and not the party that sent the original instruction that is being reversed.
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructed agent is the party receiving the reversal message and not the party that received the original instruction that is being reversed.
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`

	// Provides detailed information on the reversal reason.
	ReversalReasonInformation []*PaymentReversalReason7 `xml:"RvslRsnInf,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference24 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction81) SetReversalIdentification(value string) {
	p.ReversalIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction81) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	p.OriginalGroupInformation = new(OriginalGroupInformation3)
	return p.OriginalGroupInformation
}

func (p *PaymentTransaction81) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction81) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction81) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction81) SetOriginalClearingSystemReference(value string) {
	p.OriginalClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction81) SetOriginalInterbankSettlementAmount(value, currency string) {
	p.OriginalInterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction81) SetReversedInterbankSettlementAmount(value, currency string) {
	p.ReversedInterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction81) SetInterbankSettlementDate(value string) {
	p.InterbankSettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction81) SetSettlementPriority(value string) {
	p.SettlementPriority = (*Priority3Code)(&value)
}

func (p *PaymentTransaction81) SetReversedInstructedAmount(value, currency string) {
	p.ReversedInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction81) SetExchangeRate(value string) {
	p.ExchangeRate = (*BaseOneRate)(&value)
}

func (p *PaymentTransaction81) SetCompensationAmount(value, currency string) {
	p.CompensationAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction81) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentTransaction81) AddChargesInformation() *Charges2 {
	newValue := new(Charges2)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransaction81) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructingAgent
}

func (p *PaymentTransaction81) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructedAgent
}

func (p *PaymentTransaction81) AddReversalReasonInformation() *PaymentReversalReason7 {
	newValue := new(PaymentReversalReason7)
	p.ReversalReasonInformation = append(p.ReversalReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction81) AddOriginalTransactionReference() *OriginalTransactionReference24 {
	p.OriginalTransactionReference = new(OriginalTransactionReference24)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction81) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

// Provides further details on the original transactions, to which the status report message refers.
type PaymentTransaction82 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reported status.
	// Usage: The instructing party is the party sending the status message and not the party that sent the original instruction that is being reported on.
	StatusIdentification *Max35Text `xml:"StsId,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Specifies the status of a transaction, in a coded form.
	TransactionStatus *ExternalPaymentTransactionStatus1Code `xml:"TxSts,omitempty"`

	// Provides detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Provides information on the charges related to the processing of the rejection of the instruction.
	// Usage: This is passed on for information purposes only. Settlement of the charges will be done separately.
	ChargesInformation []*Charges2 `xml:"ChrgsInf,omitempty"`

	// Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.
	AcceptanceDateTime *ISODateTime `xml:"AccptncDtTm,omitempty"`

	// Unique reference, as assigned by the account servicing institution, to unambiguously identify the instruction.
	AccountServicerReference *Max35Text `xml:"AcctSvcrRef,omitempty"`

	// Unique reference, as assigned by a clearing system, to unambiguously identify the instruction.
	ClearingSystemReference *Max35Text `xml:"ClrSysRef,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference24 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction82) SetStatusIdentification(value string) {
	p.StatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction82) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction82) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction82) SetTransactionStatus(value string) {
	p.TransactionStatus = (*ExternalPaymentTransactionStatus1Code)(&value)
}

func (p *PaymentTransaction82) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	p.StatusReasonInformation = append(p.StatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction82) AddChargesInformation() *Charges2 {
	newValue := new(Charges2)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransaction82) SetAcceptanceDateTime(value string) {
	p.AcceptanceDateTime = (*ISODateTime)(&value)
}

func (p *PaymentTransaction82) SetAccountServicerReference(value string) {
	p.AccountServicerReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction82) SetClearingSystemReference(value string) {
	p.ClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction82) AddOriginalTransactionReference() *OriginalTransactionReference24 {
	p.OriginalTransactionReference = new(OriginalTransactionReference24)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction82) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}

// Provides further details on the original transactions, to which the status report message refers.
type PaymentTransaction83 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reported status.
	// Usage: The instructing party is the party sending the status message and not the party that sent the original instruction that is being reported on.
	StatusIdentification *Max35Text `xml:"StsId,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Specifies the status of a transaction, in a coded form.
	TransactionStatus *ExternalPaymentTransactionStatus1Code `xml:"TxSts,omitempty"`

	// Provides detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Provides information on the charges related to the processing of the rejection of the instruction.
	// Usage: This is passed on for information purposes only. Settlement of the charges will be done separately.
	ChargesInformation []*Charges2 `xml:"ChrgsInf,omitempty"`

	// Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.
	AcceptanceDateTime *ISODateTime `xml:"AccptncDtTm,omitempty"`

	// Unique reference, as assigned by the account servicing institution, to unambiguously identify the instruction.
	AccountServicerReference *Max35Text `xml:"AcctSvcrRef,omitempty"`

	// Unique reference, as assigned by a clearing system, to unambiguously identify the instruction.
	ClearingSystemReference *Max35Text `xml:"ClrSysRef,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference26 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction83) SetStatusIdentification(value string) {
	p.StatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction83) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction83) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction83) SetTransactionStatus(value string) {
	p.TransactionStatus = (*ExternalPaymentTransactionStatus1Code)(&value)
}

func (p *PaymentTransaction83) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	p.StatusReasonInformation = append(p.StatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction83) AddChargesInformation() *Charges2 {
	newValue := new(Charges2)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransaction83) SetAcceptanceDateTime(value string) {
	p.AcceptanceDateTime = (*ISODateTime)(&value)
}

func (p *PaymentTransaction83) SetAccountServicerReference(value string) {
	p.AccountServicerReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction83) SetClearingSystemReference(value string) {
	p.ClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction83) AddOriginalTransactionReference() *OriginalTransactionReference26 {
	p.OriginalTransactionReference = new(OriginalTransactionReference26)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction83) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}
