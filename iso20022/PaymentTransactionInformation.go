package iso20022

// Information concerning the original transactions, to which the status report message refers.
type PaymentTransactionInformation1 struct {

	// Unique identification as assigned by an instructing party for an instructed party to unambiguously identify the reported status.
	StatusIdentification *Max35Text `xml:"StsId,omitempty"`

	// Unique and unambiguous identifier of the original payment information block as assigned by the original sending party.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId,omitempty"`

	// Original unique instruction identification as assigned by an instructing party for an instructed party to unambiguously identify the original instruction.
	//
	// Usage: the original instruction identification is the original point to point reference used between the instructing party and the instructed party to refer to the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Original unique identification assigned by the initiating party to unambiguously identify the original transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Original identification of a transaction, as assigned by the first instructing agent and passed on, unchanged, throughout the entire interbank chain.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Specifies the status of a transaction, in a coded form.
	TransactionStatus *TransactionIndividualStatus1Code `xml:"TxSts,omitempty"`

	// Detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation1 `xml:"StsRsnInf,omitempty"`

	// Information on charges related to the processing of the rejection of the instruction.
	//
	// Usage: ChargesInformation is past on for information purposes only. Settlement of the charges will be done separately.
	ChargesInformation []*ChargesInformation1 `xml:"ChrgsInf,omitempty"`

	// Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent (debtor's agent in case of a credit transfer, creditor's agent in case of a direct debit). This means - amongst others - that the account servicing agent has received the payment order and has applied checks as eg, authorisation, availability of funds.
	AcceptanceDateTime *ISODateTime `xml:"AccptncDtTm,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstdAgt,omitempty"`

	// Set of key elements of the original transaction being referred to.
	OriginalTransactionReference *OriginalTransactionReference1 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransactionInformation1) SetStatusIdentification(value string) {
	p.StatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation1) SetOriginalPaymentInformationIdentification(value string) {
	p.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation1) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation1) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation1) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation1) SetTransactionStatus(value string) {
	p.TransactionStatus = (*TransactionIndividualStatus1Code)(&value)
}

func (p *PaymentTransactionInformation1) AddStatusReasonInformation() *StatusReasonInformation1 {
	newValue := new(StatusReasonInformation1)
	p.StatusReasonInformation = append(p.StatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation1) AddChargesInformation() *ChargesInformation1 {
	newValue := new(ChargesInformation1)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation1) SetAcceptanceDateTime(value string) {
	p.AcceptanceDateTime = (*ISODateTime)(&value)
}

func (p *PaymentTransactionInformation1) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification3 {
	p.InstructingAgent = new(BranchAndFinancialInstitutionIdentification3)
	return p.InstructingAgent
}

func (p *PaymentTransactionInformation1) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification3 {
	p.InstructedAgent = new(BranchAndFinancialInstitutionIdentification3)
	return p.InstructedAgent
}

func (p *PaymentTransactionInformation1) AddOriginalTransactionReference() *OriginalTransactionReference1 {
	p.OriginalTransactionReference = new(OriginalTransactionReference1)
	return p.OriginalTransactionReference
}

// Reference and status information concerning the original transactions, included in the original instruction, to which the return message applies.
type PaymentTransactionInformation2 struct {

	// Unique identification as assigned by an instructing party for an instructed party to unambiguously identify the returned transaction.
	ReturnIdentification *Max35Text `xml:"RtrId,omitempty"`

	// Information concerning the original group of transactions, to which the message refers.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Original unique instruction identification as assigned by an instructing party for an instructed party to unambiguously identify the original instruction.
	//
	// Usage: the original instruction identification is the original point to point reference used between the instructing party and the instructed party to refer to the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Original unique identification assigned by the initiating party to unambiguously identify the original transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Original identification of a transaction, as assigned by the first instructing agent and passed on, unchanged, throughout the entire interbank chain.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Original amount of money as moved between the instructing agent and the instructed agent in the original transaction.
	OriginalInterbankSettlementAmount *CurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty"`

	// Returned amount of money moved between the instructing agent and the instructed agent in the return transaction.
	ReturnedInterbankSettlementAmount *CurrencyAndAmount `xml:"RtrdIntrBkSttlmAmt"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	//
	// Usage: the InterbankSettlementDate is the interbank settlement date of the return message, and not of the original instruction.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Amount of money to be moved between the debtor and the creditor, before deduction of charges, in the returned transaction.
	ReturnedInstructedAmount *CurrencyAndAmount `xml:"RtrdInstdAmt,omitempty"`

	// The factor used for conversion of an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Amount of money asked or paid as compensation for the processing of the instruction.
	CompensationAmount *CurrencyAndAmount `xml:"CompstnAmt,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	//
	// Usage: The ChargeBearer applies to the return message, not to the original instruction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Transaction charges to be paid by the charge bearer for the return transaction.
	ChargesInformation []*ChargesInformation1 `xml:"ChrgsInf,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstdAgt,omitempty"`

	// Detailed information on the return reason.
	ReturnReasonInformation []*ReturnReasonInformation1 `xml:"RtrRsnInf,omitempty"`

	// Set of key elements of the original transaction being referred to.
	OriginalTransactionReference *OriginalTransactionReference1 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransactionInformation2) SetReturnIdentification(value string) {
	p.ReturnIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation2) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	p.OriginalGroupInformation = new(OriginalGroupInformation3)
	return p.OriginalGroupInformation
}

func (p *PaymentTransactionInformation2) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation2) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation2) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation2) SetOriginalInterbankSettlementAmount(value, currency string) {
	p.OriginalInterbankSettlementAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation2) SetReturnedInterbankSettlementAmount(value, currency string) {
	p.ReturnedInterbankSettlementAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation2) SetInterbankSettlementDate(value string) {
	p.InterbankSettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransactionInformation2) SetReturnedInstructedAmount(value, currency string) {
	p.ReturnedInstructedAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation2) SetExchangeRate(value string) {
	p.ExchangeRate = (*BaseOneRate)(&value)
}

func (p *PaymentTransactionInformation2) SetCompensationAmount(value, currency string) {
	p.CompensationAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation2) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentTransactionInformation2) AddChargesInformation() *ChargesInformation1 {
	newValue := new(ChargesInformation1)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation2) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification3 {
	p.InstructingAgent = new(BranchAndFinancialInstitutionIdentification3)
	return p.InstructingAgent
}

func (p *PaymentTransactionInformation2) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification3 {
	p.InstructedAgent = new(BranchAndFinancialInstitutionIdentification3)
	return p.InstructedAgent
}

func (p *PaymentTransactionInformation2) AddReturnReasonInformation() *ReturnReasonInformation1 {
	newValue := new(ReturnReasonInformation1)
	p.ReturnReasonInformation = append(p.ReturnReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation2) AddOriginalTransactionReference() *OriginalTransactionReference1 {
	p.OriginalTransactionReference = new(OriginalTransactionReference1)
	return p.OriginalTransactionReference
}

// Set of elements used to provide information on the original transactions, to which the status report message refers.
type PaymentTransactionInformation25 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reported status.
	// Usage: The instructing party is the party sending the status message and not the party that sent the original instruction that is being reported on.
	StatusIdentification *Max35Text `xml:"StsId,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Specifies the status of a transaction, in a coded form.
	TransactionStatus *TransactionIndividualStatus3Code `xml:"TxSts,omitempty"`

	// Set of elements used to provide detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation8 `xml:"StsRsnInf,omitempty"`

	// Provides information on the charges related to the processing of the rejection of the instruction.
	// Usage: This is passed on for information purposes only. Settlement of the charges will be done separately.
	ChargesInformation []*ChargesInformation5 `xml:"ChrgsInf,omitempty"`

	// Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.
	AcceptanceDateTime *ISODateTime `xml:"AccptncDtTm,omitempty"`

	// Unique reference, as assigned by the account servicing institution, to unambiguously identify the instruction.
	AccountServicerReference *Max35Text `xml:"AcctSvcrRef,omitempty"`

	// Unique reference, as assigned by a clearing system, to unambiguously identify the instruction.
	ClearingSystemReference *Max35Text `xml:"ClrSysRef,omitempty"`

	// Set of key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference13 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransactionInformation25) SetStatusIdentification(value string) {
	p.StatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation25) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation25) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation25) SetTransactionStatus(value string) {
	p.TransactionStatus = (*TransactionIndividualStatus3Code)(&value)
}

func (p *PaymentTransactionInformation25) AddStatusReasonInformation() *StatusReasonInformation8 {
	newValue := new(StatusReasonInformation8)
	p.StatusReasonInformation = append(p.StatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation25) AddChargesInformation() *ChargesInformation5 {
	newValue := new(ChargesInformation5)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation25) SetAcceptanceDateTime(value string) {
	p.AcceptanceDateTime = (*ISODateTime)(&value)
}

func (p *PaymentTransactionInformation25) SetAccountServicerReference(value string) {
	p.AccountServicerReference = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation25) SetClearingSystemReference(value string) {
	p.ClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation25) AddOriginalTransactionReference() *OriginalTransactionReference13 {
	p.OriginalTransactionReference = new(OriginalTransactionReference13)
	return p.OriginalTransactionReference
}

// Set of elements used to provide information on the original transactions, to which the status report message refers.
type PaymentTransactionInformation26 struct {

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

	// Set of elements used to provide detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation8 `xml:"StsRsnInf,omitempty"`

	// Provides information on the charges related to the processing of the rejection of the instruction.
	//
	// Usage: This is passed on for information purposes only. Settlement of the charges will be done separately.
	ChargesInformation []*ChargesInformation5 `xml:"ChrgsInf,omitempty"`

	// Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.
	AcceptanceDateTime *ISODateTime `xml:"AccptncDtTm,omitempty"`

	// Unique reference, as assigned by the account servicing institution, to unambiguously identify the instruction.
	AccountServicerReference *Max35Text `xml:"AcctSvcrRef,omitempty"`

	// Unique reference, as assigned by a clearing system, to unambiguously identify the instruction.
	ClearingSystemReference *Max35Text `xml:"ClrSysRef,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructing agent is the party sending the status message and not the party that sent the original instruction that is being reported on.
	InstructingAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructed agent is the party receiving the status message and not the party that received the original instruction that is being reported on.
	InstructedAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstdAgt,omitempty"`

	// Set of key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference13 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransactionInformation26) SetStatusIdentification(value string) {
	p.StatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation26) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation26) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation26) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation26) SetTransactionStatus(value string) {
	p.TransactionStatus = (*TransactionIndividualStatus3Code)(&value)
}

func (p *PaymentTransactionInformation26) AddStatusReasonInformation() *StatusReasonInformation8 {
	newValue := new(StatusReasonInformation8)
	p.StatusReasonInformation = append(p.StatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation26) AddChargesInformation() *ChargesInformation5 {
	newValue := new(ChargesInformation5)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation26) SetAcceptanceDateTime(value string) {
	p.AcceptanceDateTime = (*ISODateTime)(&value)
}

func (p *PaymentTransactionInformation26) SetAccountServicerReference(value string) {
	p.AccountServicerReference = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation26) SetClearingSystemReference(value string) {
	p.ClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation26) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification4 {
	p.InstructingAgent = new(BranchAndFinancialInstitutionIdentification4)
	return p.InstructingAgent
}

func (p *PaymentTransactionInformation26) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification4 {
	p.InstructedAgent = new(BranchAndFinancialInstitutionIdentification4)
	return p.InstructedAgent
}

func (p *PaymentTransactionInformation26) AddOriginalTransactionReference() *OriginalTransactionReference13 {
	p.OriginalTransactionReference = new(OriginalTransactionReference13)
	return p.OriginalTransactionReference
}

// Set of elements used to provide reference and status information on the original transactions, included in the original instruction, to which the return message applies.
type PaymentTransactionInformation27 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the returned transaction.
	// Usage: The instructing party is the party sending the return message and not the party that sent the original instruction that is being returned.
	ReturnIdentification *Max35Text `xml:"RtrId,omitempty"`

	// Set of elements used to provide information on the original messsage.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to umambiguously identify the transaction.
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
	ChargesInformation []*ChargesInformation5 `xml:"ChrgsInf,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructing agent is the party sending the return message and not the party that sent the original instruction that is being returned.
	InstructingAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructed agent is the party receiving the return message and not the party that received the original instruction that is being returned.
	InstructedAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstdAgt,omitempty"`

	// Set of elements used to provide detailed information on the return reason.
	ReturnReasonInformation []*ReturnReasonInformation9 `xml:"RtrRsnInf,omitempty"`

	// Set of key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference13 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransactionInformation27) SetReturnIdentification(value string) {
	p.ReturnIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation27) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	p.OriginalGroupInformation = new(OriginalGroupInformation3)
	return p.OriginalGroupInformation
}

func (p *PaymentTransactionInformation27) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation27) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation27) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation27) SetOriginalClearingSystemReference(value string) {
	p.OriginalClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation27) SetOriginalInterbankSettlementAmount(value, currency string) {
	p.OriginalInterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation27) SetReturnedInterbankSettlementAmount(value, currency string) {
	p.ReturnedInterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation27) SetInterbankSettlementDate(value string) {
	p.InterbankSettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransactionInformation27) SetReturnedInstructedAmount(value, currency string) {
	p.ReturnedInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation27) SetExchangeRate(value string) {
	p.ExchangeRate = (*BaseOneRate)(&value)
}

func (p *PaymentTransactionInformation27) SetCompensationAmount(value, currency string) {
	p.CompensationAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation27) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentTransactionInformation27) AddChargesInformation() *ChargesInformation5 {
	newValue := new(ChargesInformation5)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation27) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification4 {
	p.InstructingAgent = new(BranchAndFinancialInstitutionIdentification4)
	return p.InstructingAgent
}

func (p *PaymentTransactionInformation27) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification4 {
	p.InstructedAgent = new(BranchAndFinancialInstitutionIdentification4)
	return p.InstructedAgent
}

func (p *PaymentTransactionInformation27) AddReturnReasonInformation() *ReturnReasonInformation9 {
	newValue := new(ReturnReasonInformation9)
	p.ReturnReasonInformation = append(p.ReturnReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation27) AddOriginalTransactionReference() *OriginalTransactionReference13 {
	p.OriginalTransactionReference = new(OriginalTransactionReference13)
	return p.OriginalTransactionReference
}

// Set of elements used to provide reference and status information on the original transactions, included in the original instruction, to which the reversal message applies.
type PaymentTransactionInformation28 struct {

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

	// Set of elements used to provide detailed information on the reversal reason.
	ReversalReasonInformation []*ReversalReasonInformation6 `xml:"RvslRsnInf,omitempty"`

	// Set of key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference13 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransactionInformation28) SetReversalIdentification(value string) {
	p.ReversalIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation28) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation28) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation28) SetOriginalInstructedAmount(value, currency string) {
	p.OriginalInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation28) SetReversedInstructedAmount(value, currency string) {
	p.ReversedInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation28) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentTransactionInformation28) AddReversalReasonInformation() *ReversalReasonInformation6 {
	newValue := new(ReversalReasonInformation6)
	p.ReversalReasonInformation = append(p.ReversalReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation28) AddOriginalTransactionReference() *OriginalTransactionReference13 {
	p.OriginalTransactionReference = new(OriginalTransactionReference13)
	return p.OriginalTransactionReference
}

// Set of elements used to provide reference and status information on the original transactions, included in the original instruction, to which the reversal message applies.
type PaymentTransactionInformation29 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reversed transaction.
	// Usage: The instructing party is the party sending the reversal message and not the party that sent the original instruction that is being reversed.
	ReversalIdentification *Max35Text `xml:"RvslId,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to umambiguously identify the transaction.
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
	ChargesInformation []*ChargesInformation5 `xml:"ChrgsInf,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructing agent is the party sending the reversal message and not the party that sent the original instruction that is being reversed.
	InstructingAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructed agent is the party receiving the reversal message and not the party that received the original instruction that is being reversed.
	InstructedAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstdAgt,omitempty"`

	// Set of elements used to provide detailed information on the reversal reason.
	ReversalReasonInformation []*ReversalReasonInformation6 `xml:"RvslRsnInf,omitempty"`

	// Set of key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference13 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransactionInformation29) SetReversalIdentification(value string) {
	p.ReversalIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation29) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation29) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation29) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation29) SetOriginalClearingSystemReference(value string) {
	p.OriginalClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation29) SetOriginalInterbankSettlementAmount(value, currency string) {
	p.OriginalInterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation29) SetReversedInterbankSettlementAmount(value, currency string) {
	p.ReversedInterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation29) SetInterbankSettlementDate(value string) {
	p.InterbankSettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransactionInformation29) SetReversedInstructedAmount(value, currency string) {
	p.ReversedInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation29) SetExchangeRate(value string) {
	p.ExchangeRate = (*BaseOneRate)(&value)
}

func (p *PaymentTransactionInformation29) SetCompensationAmount(value, currency string) {
	p.CompensationAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation29) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentTransactionInformation29) AddChargesInformation() *ChargesInformation5 {
	newValue := new(ChargesInformation5)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation29) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification4 {
	p.InstructingAgent = new(BranchAndFinancialInstitutionIdentification4)
	return p.InstructingAgent
}

func (p *PaymentTransactionInformation29) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification4 {
	p.InstructedAgent = new(BranchAndFinancialInstitutionIdentification4)
	return p.InstructedAgent
}

func (p *PaymentTransactionInformation29) AddReversalReasonInformation() *ReversalReasonInformation6 {
	newValue := new(ReversalReasonInformation6)
	p.ReversalReasonInformation = append(p.ReversalReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation29) AddOriginalTransactionReference() *OriginalTransactionReference13 {
	p.OriginalTransactionReference = new(OriginalTransactionReference13)
	return p.OriginalTransactionReference
}

// Reference and status information concerning the original transactions, included in the original instruction, to which the cancellation request message applies.
type PaymentTransactionInformation3 struct {

	// Unique and unambiguous identifier of a cancellation request, assigned by an instructing party.
	//
	// Usage: the cancellation request identification can be used for reconciliation or to link tasks relating to the cancellation request transaction.
	CancellationIdentification *Max35Text `xml:"CxlId,omitempty"`

	// Unique and unambiguous identifier of the original payment information block as assigned by the original sending party.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId,omitempty"`

	// Original unique instruction identification as assigned by an instructing party for an instructed party to unambiguously identify the original instruction.
	//
	// Usage: the original instruction identification is the original point to point reference used between the instructing party and the instructed party to refer to the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Original unique identification assigned by the initiating party to unambiguously identify the original transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Original identification of a transaction, as assigned by the first instructing agent and passed on, unchanged, throughout the entire interbank chain.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Amount of money transferred between the instructing agent and the instructed agent in the original transaction.
	OriginalInterbankSettlementAmount *CurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty"`

	// Amount of money to be moved between the debtor and the creditor, before deduction of charges, expressed in the currency as ordered by the initiating party in the original transaction.
	OriginalInstructedAmount *CurrencyAndAmount `xml:"OrgnlInstdAmt,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstdAgt,omitempty"`

	// Detailed information on the cancellation reason.
	CancellationReasonInformation []*CancellationReasonInformation1 `xml:"CxlRsnInf,omitempty"`

	// Set of key elements of the original transaction being referred to.
	OriginalTransactionReference *OriginalTransactionReference1 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransactionInformation3) SetCancellationIdentification(value string) {
	p.CancellationIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation3) SetOriginalPaymentInformationIdentification(value string) {
	p.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation3) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation3) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation3) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation3) SetOriginalInterbankSettlementAmount(value, currency string) {
	p.OriginalInterbankSettlementAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation3) SetOriginalInstructedAmount(value, currency string) {
	p.OriginalInstructedAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation3) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification3 {
	p.InstructingAgent = new(BranchAndFinancialInstitutionIdentification3)
	return p.InstructingAgent
}

func (p *PaymentTransactionInformation3) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification3 {
	p.InstructedAgent = new(BranchAndFinancialInstitutionIdentification3)
	return p.InstructedAgent
}

func (p *PaymentTransactionInformation3) AddCancellationReasonInformation() *CancellationReasonInformation1 {
	newValue := new(CancellationReasonInformation1)
	p.CancellationReasonInformation = append(p.CancellationReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation3) AddOriginalTransactionReference() *OriginalTransactionReference1 {
	p.OriginalTransactionReference = new(OriginalTransactionReference1)
	return p.OriginalTransactionReference
}

// Set of elements used to provide reference and status information on the original transactions, included in the original instruction, to which the cancellation request message applies.
type PaymentTransactionInformation30 struct {

	// Unique identification, as assigned by the assigner, to unambiguously identify a cancellation request.
	//
	// Usage: The cancellation request identification can be used for reconciliation or to link tasks relating to the cancellation request.
	CancellationIdentification *Max35Text `xml:"CxlId,omitempty"`

	// Set of elements to uniquely and unambiguously identify an exception or an investigation workflow.
	Case *Case2 `xml:"Case,omitempty"`

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

	// Set of elements used to provide detailed information on the cancellation reason.
	CancellationReasonInformation []*CancellationReasonInformation3 `xml:"CxlRsnInf,omitempty"`

	// Set of key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference13 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransactionInformation30) SetCancellationIdentification(value string) {
	p.CancellationIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation30) AddCase() *Case2 {
	p.Case = new(Case2)
	return p.Case
}

func (p *PaymentTransactionInformation30) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation30) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation30) SetOriginalInstructedAmount(value, currency string) {
	p.OriginalInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation30) SetOriginalRequestedExecutionDate(value string) {
	p.OriginalRequestedExecutionDate = (*ISODate)(&value)
}

func (p *PaymentTransactionInformation30) SetOriginalRequestedCollectionDate(value string) {
	p.OriginalRequestedCollectionDate = (*ISODate)(&value)
}

func (p *PaymentTransactionInformation30) AddCancellationReasonInformation() *CancellationReasonInformation3 {
	newValue := new(CancellationReasonInformation3)
	p.CancellationReasonInformation = append(p.CancellationReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation30) AddOriginalTransactionReference() *OriginalTransactionReference13 {
	p.OriginalTransactionReference = new(OriginalTransactionReference13)
	return p.OriginalTransactionReference
}

// Set of elements used to provide reference and status information on the original transactions, included in the original instruction, to which the cancellation request message applies.
type PaymentTransactionInformation31 struct {

	// Unique and unambiguous identifier of a cancellation request, as assigned by the assigner.
	//
	// Usage: The cancellation request identification can be used for reconciliation or to link tasks relating to the cancellation request.
	CancellationIdentification *Max35Text `xml:"CxlId,omitempty"`

	// Set of elements to uniquely and unambiguously identify an exception or an investigation workflow.
	Case *Case2 `xml:"Case,omitempty"`

	// Set of elements used to provide information on the original messsage.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to umambiguously identify the transaction.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Unique reference, as assigned by the original clearing system, to unambiguously identify the original instruction.
	OriginalClearingSystemReference *Max35Text `xml:"OrgnlClrSysRef,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the instructing agent and the instructed agent.
	OriginalInterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty"`

	// Date, as provided in the original transaction, on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	OriginalInterbankSettlementDate *ISODate `xml:"OrgnlIntrBkSttlmDt,omitempty"`

	// Party who assigns the case.
	// Usage: This is also the agent that instructs the next party in the chain to carry out the (set of) cancellation request(s).
	Assigner *BranchAndFinancialInstitutionIdentification4 `xml:"Assgnr,omitempty"`

	// Party to which the case is assigned.
	// Usage: This is also the agent that is instructed by the previous party in the chain to carry out the (set of) cancellation request(s).
	Assignee *BranchAndFinancialInstitutionIdentification4 `xml:"Assgne,omitempty"`

	// Set of elements used to provide detailed information on the cancellation reason.
	CancellationReasonInformation []*CancellationReasonInformation3 `xml:"CxlRsnInf,omitempty"`

	// Set of key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference13 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransactionInformation31) SetCancellationIdentification(value string) {
	p.CancellationIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation31) AddCase() *Case2 {
	p.Case = new(Case2)
	return p.Case
}

func (p *PaymentTransactionInformation31) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	p.OriginalGroupInformation = new(OriginalGroupInformation3)
	return p.OriginalGroupInformation
}

func (p *PaymentTransactionInformation31) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation31) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation31) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation31) SetOriginalClearingSystemReference(value string) {
	p.OriginalClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation31) SetOriginalInterbankSettlementAmount(value, currency string) {
	p.OriginalInterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation31) SetOriginalInterbankSettlementDate(value string) {
	p.OriginalInterbankSettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransactionInformation31) AddAssigner() *BranchAndFinancialInstitutionIdentification4 {
	p.Assigner = new(BranchAndFinancialInstitutionIdentification4)
	return p.Assigner
}

func (p *PaymentTransactionInformation31) AddAssignee() *BranchAndFinancialInstitutionIdentification4 {
	p.Assignee = new(BranchAndFinancialInstitutionIdentification4)
	return p.Assignee
}

func (p *PaymentTransactionInformation31) AddCancellationReasonInformation() *CancellationReasonInformation3 {
	newValue := new(CancellationReasonInformation3)
	p.CancellationReasonInformation = append(p.CancellationReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation31) AddOriginalTransactionReference() *OriginalTransactionReference13 {
	p.OriginalTransactionReference = new(OriginalTransactionReference13)
	return p.OriginalTransactionReference
}

// Set of elements used to provide reference and status information on the original transactions, included in the original instruction, to which the cancellation request message applies.
type PaymentTransactionInformation32 struct {

	// Unique and unambiguous identifier of a cancellation request status, as assigned by the assigner.
	//
	// Usage: The cancellation status identification can be used for reconciliation or to link tasks relating to the cancellation request.
	CancellationStatusIdentification *Max35Text `xml:"CxlStsId,omitempty"`

	// Identifies the resolved case.
	ResolvedCase *Case2 `xml:"RslvdCase,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Specifies the status of the transaction cancellation request.
	TransactionCancellationStatus *CancellationIndividualStatus1Code `xml:"TxCxlSts,omitempty"`

	// Set of elements used to provide detailed information on the cancellation status reason.
	CancellationStatusReasonInformation []*CancellationStatusReasonInformation1 `xml:"CxlStsRsnInf,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the debtor and the creditor, before deduction of charges, expressed in the currency, as ordered by the original initiating party.
	OriginalInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt,omitempty"`

	// Date at which the initiating party originally requested the clearing agent to process the payment.
	OriginalRequestedExecutionDate *ISODate `xml:"OrgnlReqdExctnDt,omitempty"`

	// Date at which the creditor originally requested the collection of the amount of money from the debtor.
	OriginalRequestedCollectionDate *ISODate `xml:"OrgnlReqdColltnDt,omitempty"`

	// Set of key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference13 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransactionInformation32) SetCancellationStatusIdentification(value string) {
	p.CancellationStatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation32) AddResolvedCase() *Case2 {
	p.ResolvedCase = new(Case2)
	return p.ResolvedCase
}

func (p *PaymentTransactionInformation32) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation32) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation32) SetTransactionCancellationStatus(value string) {
	p.TransactionCancellationStatus = (*CancellationIndividualStatus1Code)(&value)
}

func (p *PaymentTransactionInformation32) AddCancellationStatusReasonInformation() *CancellationStatusReasonInformation1 {
	newValue := new(CancellationStatusReasonInformation1)
	p.CancellationStatusReasonInformation = append(p.CancellationStatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation32) SetOriginalInstructedAmount(value, currency string) {
	p.OriginalInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation32) SetOriginalRequestedExecutionDate(value string) {
	p.OriginalRequestedExecutionDate = (*ISODate)(&value)
}

func (p *PaymentTransactionInformation32) SetOriginalRequestedCollectionDate(value string) {
	p.OriginalRequestedCollectionDate = (*ISODate)(&value)
}

func (p *PaymentTransactionInformation32) AddOriginalTransactionReference() *OriginalTransactionReference13 {
	p.OriginalTransactionReference = new(OriginalTransactionReference13)
	return p.OriginalTransactionReference
}

// Set of elements used to provide reference and status information on the original transactions, included in the original instruction, to which the cancellation request message applies.
type PaymentTransactionInformation33 struct {

	// Unique and unambiguous identifier of a cancellation request status, as assigned by the assigner.
	//
	// Usage: The cancellation status identification can be used for reconciliation or to link tasks relating to the cancellation request.
	CancellationStatusIdentification *Max35Text `xml:"CxlStsId,omitempty"`

	// Identifies the resolved case.
	ResolvedCase *Case2 `xml:"RslvdCase,omitempty"`

	// Set of elements used to provide information on the original messsage.
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

	// Set of elements used to provide detailed information on the cancellation status reason.
	CancellationStatusReasonInformation []*CancellationStatusReasonInformation1 `xml:"CxlStsRsnInf,omitempty"`

	// Reference of a return or a reversal transaction that is initiated to fix the case under investigation as part of the resolution.
	ResolutionRelatedInformation *ResolutionInformation1 `xml:"RsltnRltdInf,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the instructing agent and the instructed agent.
	OriginalInterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty"`

	// Date, as provided in the original transaction, on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	OriginalInterbankSettlementDate *ISODate `xml:"OrgnlIntrBkSttlmDt,omitempty"`

	// Party who assigns the case.
	// Usage: This is also the agent that instructs the next party in the chain to carry out the (set of) cancellation request(s).
	Assigner *Party7Choice `xml:"Assgnr,omitempty"`

	// Party to which the case is assigned.
	// Usage: This is also the agent that is instructed by the previous party in the chain to carry out the (set of) cancellation request(s).
	Assignee *Party7Choice `xml:"Assgne,omitempty"`

	// Set of key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference13 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransactionInformation33) SetCancellationStatusIdentification(value string) {
	p.CancellationStatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation33) AddResolvedCase() *Case2 {
	p.ResolvedCase = new(Case2)
	return p.ResolvedCase
}

func (p *PaymentTransactionInformation33) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	p.OriginalGroupInformation = new(OriginalGroupInformation3)
	return p.OriginalGroupInformation
}

func (p *PaymentTransactionInformation33) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation33) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation33) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation33) SetOriginalClearingSystemReference(value string) {
	p.OriginalClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation33) SetTransactionCancellationStatus(value string) {
	p.TransactionCancellationStatus = (*CancellationIndividualStatus1Code)(&value)
}

func (p *PaymentTransactionInformation33) AddCancellationStatusReasonInformation() *CancellationStatusReasonInformation1 {
	newValue := new(CancellationStatusReasonInformation1)
	p.CancellationStatusReasonInformation = append(p.CancellationStatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation33) AddResolutionRelatedInformation() *ResolutionInformation1 {
	p.ResolutionRelatedInformation = new(ResolutionInformation1)
	return p.ResolutionRelatedInformation
}

func (p *PaymentTransactionInformation33) SetOriginalInterbankSettlementAmount(value, currency string) {
	p.OriginalInterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation33) SetOriginalInterbankSettlementDate(value string) {
	p.OriginalInterbankSettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransactionInformation33) AddAssigner() *Party7Choice {
	p.Assigner = new(Party7Choice)
	return p.Assigner
}

func (p *PaymentTransactionInformation33) AddAssignee() *Party7Choice {
	p.Assignee = new(Party7Choice)
	return p.Assignee
}

func (p *PaymentTransactionInformation33) AddOriginalTransactionReference() *OriginalTransactionReference13 {
	p.OriginalTransactionReference = new(OriginalTransactionReference13)
	return p.OriginalTransactionReference
}

// Set of elements used to provide information on the original transactions, to which the status report message refers.
type PaymentTransactionInformation34 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reported status.
	// Usage: The instructing party is the party sending the status message and not the party that sent the original instruction that is being reported on.
	StatusIdentification *Max35Text `xml:"StsId,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Specifies the status of a transaction, in a coded form.
	TransactionStatus *TransactionIndividualStatus3Code `xml:"TxSts,omitempty"`

	// Set of elements used to provide detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Provides information on the charges related to the processing of the rejection of the instruction.
	// Usage: This is passed on for information purposes only. Settlement of the charges will be done separately.
	ChargesInformation []*ChargesInformation7 `xml:"ChrgsInf,omitempty"`

	// Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.
	AcceptanceDateTime *ISODateTime `xml:"AccptncDtTm,omitempty"`

	// Unique reference, as assigned by the account servicing institution, to unambiguously identify the instruction.
	AccountServicerReference *Max35Text `xml:"AcctSvcrRef,omitempty"`

	// Unique reference, as assigned by a clearing system, to unambiguously identify the instruction.
	ClearingSystemReference *Max35Text `xml:"ClrSysRef,omitempty"`

	// Set of key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference15 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransactionInformation34) SetStatusIdentification(value string) {
	p.StatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation34) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation34) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation34) SetTransactionStatus(value string) {
	p.TransactionStatus = (*TransactionIndividualStatus3Code)(&value)
}

func (p *PaymentTransactionInformation34) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	p.StatusReasonInformation = append(p.StatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation34) AddChargesInformation() *ChargesInformation7 {
	newValue := new(ChargesInformation7)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation34) SetAcceptanceDateTime(value string) {
	p.AcceptanceDateTime = (*ISODateTime)(&value)
}

func (p *PaymentTransactionInformation34) SetAccountServicerReference(value string) {
	p.AccountServicerReference = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation34) SetClearingSystemReference(value string) {
	p.ClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation34) AddOriginalTransactionReference() *OriginalTransactionReference15 {
	p.OriginalTransactionReference = new(OriginalTransactionReference15)
	return p.OriginalTransactionReference
}

// Reference and status information concerning the original transactions, included in the original instruction, to which the reversal message applies.
type PaymentTransactionInformation4 struct {

	// Unique identification as assigned by an instructing party for an instructed party to unambiguously identify the reversed transaction.
	ReversalIdentification *Max35Text `xml:"RvslId,omitempty"`

	// Unique and unambiguous identifier of the original payment information block as assigned by the original sending party.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId,omitempty"`

	// Original unique instruction identification as assigned by an instructing party for an instructed party to unambiguously identify the original instruction.
	//
	// Usage: the original instruction identification is the original point to point reference used between the instructing party and the instructed party to refer to the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Original unique identification assigned by the initiating party to unambiguously identify the original transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Amount of money to be moved between the debtor and the creditor, before deduction of charges, expressed in the currency as ordered by the initiating party in the original transaction.
	OriginalInstructedAmount *CurrencyAndAmount `xml:"OrgnlInstdAmt,omitempty"`

	// Amount of money to be moved between the debtor and the creditor, before deduction of charges, in the reversed transaction.
	ReversedInstructedAmount *CurrencyAndAmount `xml:"RvsdInstdAmt,omitempty"`

	// Specifies if the creditor and/or debtor will bear the charges associated with the processing of the payment transaction.
	//
	// Usage: The ChargeBearer applies to the reversal message, not to the original instruction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Detailed information on the reversal reason.
	ReversalReasonInformation []*ReversalReasonInformation1 `xml:"RvslRsnInf,omitempty"`

	// Set of key elements of the original transaction being referred to.
	OriginalTransactionReference *OriginalTransactionReference1 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransactionInformation4) SetReversalIdentification(value string) {
	p.ReversalIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation4) SetOriginalPaymentInformationIdentification(value string) {
	p.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation4) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation4) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation4) SetOriginalInstructedAmount(value, currency string) {
	p.OriginalInstructedAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation4) SetReversedInstructedAmount(value, currency string) {
	p.ReversedInstructedAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation4) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentTransactionInformation4) AddReversalReasonInformation() *ReversalReasonInformation1 {
	newValue := new(ReversalReasonInformation1)
	p.ReversalReasonInformation = append(p.ReversalReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation4) AddOriginalTransactionReference() *OriginalTransactionReference1 {
	p.OriginalTransactionReference = new(OriginalTransactionReference1)
	return p.OriginalTransactionReference
}

// Reference and status information concerning the original transactions, included in the original instruction, to which the reversal message applies.
type PaymentTransactionInformation5 struct {

	// Unique identification as assigned by an instructing party for an instructed party to unambiguously identify the reversed transaction.
	ReversalIdentification *Max35Text `xml:"RvslId,omitempty"`

	// Unique and unambiguous identifier of the original payment information block as assigned by the original sending party.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId,omitempty"`

	// Original unique instruction identification as assigned by an instructing party for an instructed party to unambiguously identify the original instruction.
	//
	// Usage: the original instruction identification is the original point to point reference used between the instructing party and the instructed party to refer to the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Original unique identification assigned by the initiating party to unambiguously identify the original transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Original identification of a transaction, as assigned by the first instructing agent and passed on, unchanged, throughout the entire interbank chain.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Amount of money transferred between the instructing agent and the instructed agent in the original transaction.
	OriginalInterbankSettlementAmount *CurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent in the reversed transaction.
	ReversedInterbankSettlementAmount *CurrencyAndAmount `xml:"RvsdIntrBkSttlmAmt"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	//
	// Usage: the InterbankSettlementDate is the interbank settlement date of the reversal message, and not of the original instruction.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Amount of money to be moved between the debtor and the creditor, before deduction of charges, in the reversed transaction.
	ReversedInstructedAmount *CurrencyAndAmount `xml:"RvsdInstdAmt,omitempty"`

	// The factor used for conversion of an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Amount of money asked or paid as compensation for the processing of the instruction.
	CompensationAmount *CurrencyAndAmount `xml:"CompstnAmt,omitempty"`

	// Specifies if the creditor and/or debtor will bear the charges associated with the processing of the payment transaction.
	//
	// Usage: The ChargeBearer applies to the reversal message, not to the original instruction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Transaction charges to be paid by the charge bearer for the reversal transaction.
	ChargesInformation []*ChargesInformation1 `xml:"ChrgsInf,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstdAgt,omitempty"`

	// Detailed information on the reversal reason.
	ReversalReasonInformation []*ReversalReasonInformation1 `xml:"RvslRsnInf,omitempty"`

	// Set of key elements of the original transaction being referred to.
	OriginalTransactionReference *OriginalTransactionReference1 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransactionInformation5) SetReversalIdentification(value string) {
	p.ReversalIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation5) SetOriginalPaymentInformationIdentification(value string) {
	p.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation5) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation5) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation5) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation5) SetOriginalInterbankSettlementAmount(value, currency string) {
	p.OriginalInterbankSettlementAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation5) SetReversedInterbankSettlementAmount(value, currency string) {
	p.ReversedInterbankSettlementAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation5) SetInterbankSettlementDate(value string) {
	p.InterbankSettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransactionInformation5) SetReversedInstructedAmount(value, currency string) {
	p.ReversedInstructedAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation5) SetExchangeRate(value string) {
	p.ExchangeRate = (*BaseOneRate)(&value)
}

func (p *PaymentTransactionInformation5) SetCompensationAmount(value, currency string) {
	p.CompensationAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation5) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentTransactionInformation5) AddChargesInformation() *ChargesInformation1 {
	newValue := new(ChargesInformation1)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation5) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification3 {
	p.InstructingAgent = new(BranchAndFinancialInstitutionIdentification3)
	return p.InstructingAgent
}

func (p *PaymentTransactionInformation5) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification3 {
	p.InstructedAgent = new(BranchAndFinancialInstitutionIdentification3)
	return p.InstructedAgent
}

func (p *PaymentTransactionInformation5) AddReversalReasonInformation() *ReversalReasonInformation1 {
	newValue := new(ReversalReasonInformation1)
	p.ReversalReasonInformation = append(p.ReversalReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation5) AddOriginalTransactionReference() *OriginalTransactionReference1 {
	p.OriginalTransactionReference = new(OriginalTransactionReference1)
	return p.OriginalTransactionReference
}
