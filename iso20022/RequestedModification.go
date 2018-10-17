package iso20022

// Contains the requested modifications.
type RequestedModification struct {

	// Reference relating to a linked payment instruction or agreement which is meaningful to both parties (eg, the content of field 21 in a cover instruction).
	RelatedReference *Max35Text `xml:"RltdRef,omitempty"`

	// SWIFT defined service level applies to the payment instruction.
	BankOperationCode *SWIFTServiceLevel2Code `xml:"BkOprCd,omitempty"`

	// Further information related to the processing of the payment instruction. The instruction can relate to a level of service between the bank and the customer, or give instructions to and for specific parties in the payment chain.
	InstructionCode *Instruction1Code `xml:"InstrCd,omitempty"`

	// Date and time the debtor requests the clearing agent to process the payment instruction.
	RequestedExecutionDate *ISODate `xml:"ReqdExctnDt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	ValueDate *ISODate `xml:"ValDt,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent.
	InterbankSettledAmount *CurrencyAndAmount `xml:"IntrBkSttldAmt,omitempty"`

	// Debtor or Ordering customer of the payment instruction.
	Debtor *PartyIdentification1 `xml:"Dbtr,omitempty"`

	// Account to or from which a cash entry is made.
	DebtorAccount *CashAccount3 `xml:"DbtrAcct,omitempty"`

	// Party that executes a cash transfer received via a clearing agent or on request of an agreement party.
	IntermediarySettlementAgent *BranchAndFinancialInstitutionIdentification `xml:"IntrmySttlmAgt,omitempty"`

	// Party that executes a cash transfer received via a clearing agent or on request of an agreement party.
	LastSettlementAgent *BranchAndFinancialInstitutionIdentification `xml:"LastSttlmAgt,omitempty"`

	// Specification of a pre-agreed offering between clearing agents, or the channel through which the payment instruction is to be processed. This payment scheme can point to a specific rulebook governing the rules of clearing and settlement between two parties.
	PaymentScheme *PaymentSchemeChoice `xml:"PmtSchme,omitempty"`

	// Account to or from which a cash entry is made.
	BeneficiaryInstitutionAccount *CashAccount3 `xml:"BnfcryInstnAcct,omitempty"`

	// Entity involved in an activity.
	Creditor *PartyIdentification1 `xml:"Cdtr,omitempty"`

	// Account to or from which a cash entry is made.
	CreditorAccount *CashAccount3 `xml:"CdtrAcct,omitempty"`

	// Structured information that enables the matching, ie, reconciliation, of a payment with the items that the payment is intended to settle, such as commercial invoices in an account receivable system.
	RemittanceInformation *RemittanceInformation3Choice `xml:"RmtInf,omitempty"`

	// Underlying reason for the payment transaction.
	Purpose *PurposeChoice `xml:"Purp,omitempty"`

	// Further information related to the processing of the payment instruction. The instruction can relate to a level of service between the bank and the customer, or give instructions to and for specific parties in the payment chain.
	InstructionForFinalAgent *InstructionForFinalAgent `xml:"InstrForFnlAgt,omitempty"`

	// Party(ies) liable for a charge associated with the transfer of cash.
	DetailsOfCharges *ChargeBearer1Code `xml:"DtlsOfChrgs,omitempty"`

	// Unformatted information from the sender to the receiver.
	SenderToReceiverInformation []*Max35Text `xml:"SndrToRcvrInf,omitempty"`
}

func (r *RequestedModification) SetRelatedReference(value string) {
	r.RelatedReference = (*Max35Text)(&value)
}

func (r *RequestedModification) SetBankOperationCode(value string) {
	r.BankOperationCode = (*SWIFTServiceLevel2Code)(&value)
}

func (r *RequestedModification) SetInstructionCode(value string) {
	r.InstructionCode = (*Instruction1Code)(&value)
}

func (r *RequestedModification) SetRequestedExecutionDate(value string) {
	r.RequestedExecutionDate = (*ISODate)(&value)
}

func (r *RequestedModification) SetValueDate(value string) {
	r.ValueDate = (*ISODate)(&value)
}

func (r *RequestedModification) SetInterbankSettledAmount(value, currency string) {
	r.InterbankSettledAmount = NewCurrencyAndAmount(value, currency)
}

func (r *RequestedModification) AddDebtor() *PartyIdentification1 {
	r.Debtor = new(PartyIdentification1)
	return r.Debtor
}

func (r *RequestedModification) AddDebtorAccount() *CashAccount3 {
	r.DebtorAccount = new(CashAccount3)
	return r.DebtorAccount
}

func (r *RequestedModification) AddIntermediarySettlementAgent() *BranchAndFinancialInstitutionIdentification {
	r.IntermediarySettlementAgent = new(BranchAndFinancialInstitutionIdentification)
	return r.IntermediarySettlementAgent
}

func (r *RequestedModification) AddLastSettlementAgent() *BranchAndFinancialInstitutionIdentification {
	r.LastSettlementAgent = new(BranchAndFinancialInstitutionIdentification)
	return r.LastSettlementAgent
}

func (r *RequestedModification) AddPaymentScheme() *PaymentSchemeChoice {
	r.PaymentScheme = new(PaymentSchemeChoice)
	return r.PaymentScheme
}

func (r *RequestedModification) AddBeneficiaryInstitutionAccount() *CashAccount3 {
	r.BeneficiaryInstitutionAccount = new(CashAccount3)
	return r.BeneficiaryInstitutionAccount
}

func (r *RequestedModification) AddCreditor() *PartyIdentification1 {
	r.Creditor = new(PartyIdentification1)
	return r.Creditor
}

func (r *RequestedModification) AddCreditorAccount() *CashAccount3 {
	r.CreditorAccount = new(CashAccount3)
	return r.CreditorAccount
}

func (r *RequestedModification) AddRemittanceInformation() *RemittanceInformation3Choice {
	r.RemittanceInformation = new(RemittanceInformation3Choice)
	return r.RemittanceInformation
}

func (r *RequestedModification) AddPurpose() *PurposeChoice {
	r.Purpose = new(PurposeChoice)
	return r.Purpose
}

func (r *RequestedModification) AddInstructionForFinalAgent() *InstructionForFinalAgent {
	r.InstructionForFinalAgent = new(InstructionForFinalAgent)
	return r.InstructionForFinalAgent
}

func (r *RequestedModification) SetDetailsOfCharges(value string) {
	r.DetailsOfCharges = (*ChargeBearer1Code)(&value)
}

func (r *RequestedModification) AddSenderToReceiverInformation(value string) {
	r.SenderToReceiverInformation = append(r.SenderToReceiverInformation, (*Max35Text)(&value))
}

// Set of elements used to provide information on the requested modifications of the underlying payment instruction.
type RequestedModification2 struct {

	// Unique identification as assigned by an instructing party for an instructed party to unambiguously identify the instruction.
	//
	// Usage: The instruction identification is a point to point reference that can be used between the instructing party and the instructed party to refer to the individual instruction. It can be included in several messages related to the instruction.
	InstructionIdentification *Max35Text `xml:"InstrId,omitempty"`

	// Unique identification, as assigned by the initiating party, to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	//
	// Usage: The end-to-end identification can be used for reconciliation or to link tasks relating to the transaction. It can be included in several messages related to the transaction.
	EndToEndIdentification *Max35Text `xml:"EndToEndId,omitempty"`

	// Unique identification, as assigned by the first instructing agent, to unambiguously identify the transaction that is passed on, unchanged, throughout the entire interbank chain.
	// Usage: The transaction identification can be used for reconciliation, tracking or to link tasks relating to the transaction on the interbank level. The instructing agent has to make sure that the transaction identification is unique for a pre-agreed period.
	TransactionIdentification *Max35Text `xml:"TxId,omitempty"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation22 `xml:"PmtTpInf,omitempty"`

	// Date at which the initiating party requests the clearing agent to process the payment.
	// Usage: This is the date on which the debtor's account is to be debited. If payment by cheque, the date when the cheque must be generated by the bank.
	RequestedExecutionDate *ISODate `xml:"ReqdExctnDt,omitempty"`

	// Date and time at which the creditor requests that the amount of money is to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	Amount *AmountType3Choice `xml:"Amt,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent.
	InterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"IntrBkSttlmAmt,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification32 `xml:"UltmtDbtr,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor. In the context of the payment model, the debtor is also the debit account owner.
	Debtor *PartyIdentification32 `xml:"Dbtr,omitempty"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount16 `xml:"DbtrAcct,omitempty"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount16 `xml:"DbtrAgtAcct,omitempty"`

	// Instruction between two clearing agents stipulating the cash transfer characteristics between the two parties.
	SettlementInformation *SettlementInformation15 `xml:"SttlmInf,omitempty"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CreditorAgentAccount *CashAccount16 `xml:"CdtrAgtAcct,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification32 `xml:"Cdtr,omitempty"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount16 `xml:"CdtrAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification32 `xml:"UltmtCdtr,omitempty"`

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`

	// Further information related to the processing of the payment instruction, provided by the initiating party, and intended for the debtor agent.
	InstructionForDebtorAgent *Max140Text `xml:"InstrForDbtrAgt,omitempty"`

	// Further information related to the processing of the payment instruction that may need to be acted upon by the next agent.
	//
	// Usage: The next agent may not be the creditor agent.
	// The instruction can relate to a level of service, can be an instruction that has to be executed by the agent, or can be information required by the next agent.
	InstructionForNextAgent []*InstructionForNextAgent1 `xml:"InstrForNxtAgt,omitempty"`

	// Further information related to the processing of the payment instruction, provided by the initiating party, and intended for the creditor agent.
	InstructionForCreditorAgent []*InstructionForCreditorAgent1 `xml:"InstrForCdtrAgt,omitempty"`

	// Structured information that enables the matching, ie, reconciliation, of a payment with the items that the payment is intended to settle, such as commercial invoices in an account receivable system.
	RemittanceInformation *RemittanceInformation5 `xml:"RmtInf,omitempty"`
}

func (r *RequestedModification2) SetInstructionIdentification(value string) {
	r.InstructionIdentification = (*Max35Text)(&value)
}

func (r *RequestedModification2) SetEndToEndIdentification(value string) {
	r.EndToEndIdentification = (*Max35Text)(&value)
}

func (r *RequestedModification2) SetTransactionIdentification(value string) {
	r.TransactionIdentification = (*Max35Text)(&value)
}

func (r *RequestedModification2) AddPaymentTypeInformation() *PaymentTypeInformation22 {
	r.PaymentTypeInformation = new(PaymentTypeInformation22)
	return r.PaymentTypeInformation
}

func (r *RequestedModification2) SetRequestedExecutionDate(value string) {
	r.RequestedExecutionDate = (*ISODate)(&value)
}

func (r *RequestedModification2) SetRequestedCollectionDate(value string) {
	r.RequestedCollectionDate = (*ISODate)(&value)
}

func (r *RequestedModification2) SetInterbankSettlementDate(value string) {
	r.InterbankSettlementDate = (*ISODate)(&value)
}

func (r *RequestedModification2) AddAmount() *AmountType3Choice {
	r.Amount = new(AmountType3Choice)
	return r.Amount
}

func (r *RequestedModification2) SetInterbankSettlementAmount(value, currency string) {
	r.InterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RequestedModification2) SetChargeBearer(value string) {
	r.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (r *RequestedModification2) AddUltimateDebtor() *PartyIdentification32 {
	r.UltimateDebtor = new(PartyIdentification32)
	return r.UltimateDebtor
}

func (r *RequestedModification2) AddDebtor() *PartyIdentification32 {
	r.Debtor = new(PartyIdentification32)
	return r.Debtor
}

func (r *RequestedModification2) AddDebtorAccount() *CashAccount16 {
	r.DebtorAccount = new(CashAccount16)
	return r.DebtorAccount
}

func (r *RequestedModification2) AddDebtorAgentAccount() *CashAccount16 {
	r.DebtorAgentAccount = new(CashAccount16)
	return r.DebtorAgentAccount
}

func (r *RequestedModification2) AddSettlementInformation() *SettlementInformation15 {
	r.SettlementInformation = new(SettlementInformation15)
	return r.SettlementInformation
}

func (r *RequestedModification2) AddCreditorAgentAccount() *CashAccount16 {
	r.CreditorAgentAccount = new(CashAccount16)
	return r.CreditorAgentAccount
}

func (r *RequestedModification2) AddCreditor() *PartyIdentification32 {
	r.Creditor = new(PartyIdentification32)
	return r.Creditor
}

func (r *RequestedModification2) AddCreditorAccount() *CashAccount16 {
	r.CreditorAccount = new(CashAccount16)
	return r.CreditorAccount
}

func (r *RequestedModification2) AddUltimateCreditor() *PartyIdentification32 {
	r.UltimateCreditor = new(PartyIdentification32)
	return r.UltimateCreditor
}

func (r *RequestedModification2) AddPurpose() *Purpose2Choice {
	r.Purpose = new(Purpose2Choice)
	return r.Purpose
}

func (r *RequestedModification2) SetInstructionForDebtorAgent(value string) {
	r.InstructionForDebtorAgent = (*Max140Text)(&value)
}

func (r *RequestedModification2) AddInstructionForNextAgent() *InstructionForNextAgent1 {
	newValue := new(InstructionForNextAgent1)
	r.InstructionForNextAgent = append(r.InstructionForNextAgent, newValue)
	return newValue
}

func (r *RequestedModification2) AddInstructionForCreditorAgent() *InstructionForCreditorAgent1 {
	newValue := new(InstructionForCreditorAgent1)
	r.InstructionForCreditorAgent = append(r.InstructionForCreditorAgent, newValue)
	return newValue
}

func (r *RequestedModification2) AddRemittanceInformation() *RemittanceInformation5 {
	r.RemittanceInformation = new(RemittanceInformation5)
	return r.RemittanceInformation
}

// Provide further details on the requested modifications of the underlying payment instruction.
type RequestedModification3 struct {

	// Unique identification as assigned by an instructing party for an instructed party to unambiguously identify the instruction.
	//
	// Usage: The instruction identification is a point to point reference that can be used between the instructing party and the instructed party to refer to the individual instruction. It can be included in several messages related to the instruction.
	InstructionIdentification *Max35Text `xml:"InstrId,omitempty"`

	// Unique identification, as assigned by the initiating party, to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	//
	// Usage: The end-to-end identification can be used for reconciliation or to link tasks relating to the transaction. It can be included in several messages related to the transaction.
	EndToEndIdentification *Max35Text `xml:"EndToEndId,omitempty"`

	// Unique identification, as assigned by the first instructing agent, to unambiguously identify the transaction that is passed on, unchanged, throughout the entire interbank chain.
	// Usage: The transaction identification can be used for reconciliation, tracking or to link tasks relating to the transaction on the interbank level. The instructing agent has to make sure that the transaction identification is unique for a pre-agreed period.
	TransactionIdentification *Max35Text `xml:"TxId,omitempty"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation25 `xml:"PmtTpInf,omitempty"`

	// Date at which the initiating party requests the clearing agent to process the payment.
	// Usage: This is the date on which the debtor's account is to be debited. If payment by cheque, the date when the cheque must be generated by the bank.
	RequestedExecutionDate *ISODate `xml:"ReqdExctnDt,omitempty"`

	// Date and time at which the creditor requests that the amount of money is to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	Amount *AmountType3Choice `xml:"Amt,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent.
	InterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"IntrBkSttlmAmt,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor. In the context of the payment model, the debtor is also the debit account owner.
	Debtor *PartyIdentification43 `xml:"Dbtr,omitempty"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct,omitempty"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount24 `xml:"DbtrAgtAcct,omitempty"`

	// Instruction between two clearing agents stipulating the cash transfer characteristics between the two parties.
	SettlementInformation *SettlementInstruction3 `xml:"SttlmInf,omitempty"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CreditorAgentAccount *CashAccount24 `xml:"CdtrAgtAcct,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr,omitempty"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`

	// Further information related to the processing of the payment instruction, provided by the initiating party, and intended for the debtor agent.
	InstructionForDebtorAgent *Max140Text `xml:"InstrForDbtrAgt,omitempty"`

	// Further information related to the processing of the payment instruction that may need to be acted upon by the next agent.
	//
	// Usage: The next agent may not be the creditor agent.
	// The instruction can relate to a level of service, can be an instruction that has to be executed by the agent, or can be information required by the next agent.
	InstructionForNextAgent []*InstructionForNextAgent1 `xml:"InstrForNxtAgt,omitempty"`

	// Further information related to the processing of the payment instruction, provided by the initiating party, and intended for the creditor agent.
	InstructionForCreditorAgent []*InstructionForCreditorAgent1 `xml:"InstrForCdtrAgt,omitempty"`

	// Structured information that enables the matching, that is reconciliation, of a payment with the items that the payment is intended to settle, such as commercial invoices in an account receivable system.
	RemittanceInformation *RemittanceInformation7 `xml:"RmtInf,omitempty"`
}

func (r *RequestedModification3) SetInstructionIdentification(value string) {
	r.InstructionIdentification = (*Max35Text)(&value)
}

func (r *RequestedModification3) SetEndToEndIdentification(value string) {
	r.EndToEndIdentification = (*Max35Text)(&value)
}

func (r *RequestedModification3) SetTransactionIdentification(value string) {
	r.TransactionIdentification = (*Max35Text)(&value)
}

func (r *RequestedModification3) AddPaymentTypeInformation() *PaymentTypeInformation25 {
	r.PaymentTypeInformation = new(PaymentTypeInformation25)
	return r.PaymentTypeInformation
}

func (r *RequestedModification3) SetRequestedExecutionDate(value string) {
	r.RequestedExecutionDate = (*ISODate)(&value)
}

func (r *RequestedModification3) SetRequestedCollectionDate(value string) {
	r.RequestedCollectionDate = (*ISODate)(&value)
}

func (r *RequestedModification3) SetInterbankSettlementDate(value string) {
	r.InterbankSettlementDate = (*ISODate)(&value)
}

func (r *RequestedModification3) AddAmount() *AmountType3Choice {
	r.Amount = new(AmountType3Choice)
	return r.Amount
}

func (r *RequestedModification3) SetInterbankSettlementAmount(value, currency string) {
	r.InterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RequestedModification3) SetChargeBearer(value string) {
	r.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (r *RequestedModification3) AddUltimateDebtor() *PartyIdentification43 {
	r.UltimateDebtor = new(PartyIdentification43)
	return r.UltimateDebtor
}

func (r *RequestedModification3) AddDebtor() *PartyIdentification43 {
	r.Debtor = new(PartyIdentification43)
	return r.Debtor
}

func (r *RequestedModification3) AddDebtorAccount() *CashAccount24 {
	r.DebtorAccount = new(CashAccount24)
	return r.DebtorAccount
}

func (r *RequestedModification3) AddDebtorAgentAccount() *CashAccount24 {
	r.DebtorAgentAccount = new(CashAccount24)
	return r.DebtorAgentAccount
}

func (r *RequestedModification3) AddSettlementInformation() *SettlementInstruction3 {
	r.SettlementInformation = new(SettlementInstruction3)
	return r.SettlementInformation
}

func (r *RequestedModification3) AddCreditorAgentAccount() *CashAccount24 {
	r.CreditorAgentAccount = new(CashAccount24)
	return r.CreditorAgentAccount
}

func (r *RequestedModification3) AddCreditor() *PartyIdentification43 {
	r.Creditor = new(PartyIdentification43)
	return r.Creditor
}

func (r *RequestedModification3) AddCreditorAccount() *CashAccount24 {
	r.CreditorAccount = new(CashAccount24)
	return r.CreditorAccount
}

func (r *RequestedModification3) AddUltimateCreditor() *PartyIdentification43 {
	r.UltimateCreditor = new(PartyIdentification43)
	return r.UltimateCreditor
}

func (r *RequestedModification3) AddPurpose() *Purpose2Choice {
	r.Purpose = new(Purpose2Choice)
	return r.Purpose
}

func (r *RequestedModification3) SetInstructionForDebtorAgent(value string) {
	r.InstructionForDebtorAgent = (*Max140Text)(&value)
}

func (r *RequestedModification3) AddInstructionForNextAgent() *InstructionForNextAgent1 {
	newValue := new(InstructionForNextAgent1)
	r.InstructionForNextAgent = append(r.InstructionForNextAgent, newValue)
	return newValue
}

func (r *RequestedModification3) AddInstructionForCreditorAgent() *InstructionForCreditorAgent1 {
	newValue := new(InstructionForCreditorAgent1)
	r.InstructionForCreditorAgent = append(r.InstructionForCreditorAgent, newValue)
	return newValue
}

func (r *RequestedModification3) AddRemittanceInformation() *RemittanceInformation7 {
	r.RemittanceInformation = new(RemittanceInformation7)
	return r.RemittanceInformation
}

// Provide further details on the requested modifications of the underlying payment instruction.
type RequestedModification4 struct {

	// Unique identification as assigned by an instructing party for an instructed party to unambiguously identify the instruction.
	//
	// Usage: The instruction identification is a point to point reference that can be used between the instructing party and the instructed party to refer to the individual instruction. It can be included in several messages related to the instruction.
	InstructionIdentification *Max35Text `xml:"InstrId,omitempty"`

	// Unique identification, as assigned by the initiating party, to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	//
	// Usage: The end-to-end identification can be used for reconciliation or to link tasks relating to the transaction. It can be included in several messages related to the transaction.
	EndToEndIdentification *Max35Text `xml:"EndToEndId,omitempty"`

	// Unique identification, as assigned by the first instructing agent, to unambiguously identify the transaction that is passed on, unchanged, throughout the entire interbank chain.
	// Usage: The transaction identification can be used for reconciliation, tracking or to link tasks relating to the transaction on the interbank level. The instructing agent has to make sure that the transaction identification is unique for a pre-agreed period.
	TransactionIdentification *Max35Text `xml:"TxId,omitempty"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation25 `xml:"PmtTpInf,omitempty"`

	// Date at which the initiating party requests the clearing agent to process the payment.
	// Usage: This is the date on which the debtor's account is to be debited. If payment by cheque, the date when the cheque must be generated by the bank.
	RequestedExecutionDate *ISODate `xml:"ReqdExctnDt,omitempty"`

	// Date and time at which the creditor requests that the amount of money is to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	Amount *AmountType4Choice `xml:"Amt,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent.
	InterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"IntrBkSttlmAmt,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor. In the context of the payment model, the debtor is also the debit account owner.
	Debtor *PartyIdentification43 `xml:"Dbtr,omitempty"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct,omitempty"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount24 `xml:"DbtrAgtAcct,omitempty"`

	// Instruction between two clearing agents stipulating the cash transfer characteristics between the two parties.
	SettlementInformation *SettlementInstruction3 `xml:"SttlmInf,omitempty"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CreditorAgentAccount *CashAccount24 `xml:"CdtrAgtAcct,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr,omitempty"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`

	// Further information related to the processing of the payment instruction, provided by the initiating party, and intended for the debtor agent.
	InstructionForDebtorAgent *Max140Text `xml:"InstrForDbtrAgt,omitempty"`

	// Further information related to the processing of the payment instruction that may need to be acted upon by the next agent.
	//
	// Usage: The next agent may not be the creditor agent.
	// The instruction can relate to a level of service, can be an instruction that has to be executed by the agent, or can be information required by the next agent.
	InstructionForNextAgent []*InstructionForNextAgent1 `xml:"InstrForNxtAgt,omitempty"`

	// Further information related to the processing of the payment instruction, provided by the initiating party, and intended for the creditor agent.
	InstructionForCreditorAgent []*InstructionForCreditorAgent1 `xml:"InstrForCdtrAgt,omitempty"`

	// Structured information that enables the matching, that is reconciliation, of a payment with the items that the payment is intended to settle, such as commercial invoices in an account receivable system.
	RemittanceInformation *RemittanceInformation10 `xml:"RmtInf,omitempty"`
}

func (r *RequestedModification4) SetInstructionIdentification(value string) {
	r.InstructionIdentification = (*Max35Text)(&value)
}

func (r *RequestedModification4) SetEndToEndIdentification(value string) {
	r.EndToEndIdentification = (*Max35Text)(&value)
}

func (r *RequestedModification4) SetTransactionIdentification(value string) {
	r.TransactionIdentification = (*Max35Text)(&value)
}

func (r *RequestedModification4) AddPaymentTypeInformation() *PaymentTypeInformation25 {
	r.PaymentTypeInformation = new(PaymentTypeInformation25)
	return r.PaymentTypeInformation
}

func (r *RequestedModification4) SetRequestedExecutionDate(value string) {
	r.RequestedExecutionDate = (*ISODate)(&value)
}

func (r *RequestedModification4) SetRequestedCollectionDate(value string) {
	r.RequestedCollectionDate = (*ISODate)(&value)
}

func (r *RequestedModification4) SetInterbankSettlementDate(value string) {
	r.InterbankSettlementDate = (*ISODate)(&value)
}

func (r *RequestedModification4) AddAmount() *AmountType4Choice {
	r.Amount = new(AmountType4Choice)
	return r.Amount
}

func (r *RequestedModification4) SetInterbankSettlementAmount(value, currency string) {
	r.InterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RequestedModification4) SetChargeBearer(value string) {
	r.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (r *RequestedModification4) AddUltimateDebtor() *PartyIdentification43 {
	r.UltimateDebtor = new(PartyIdentification43)
	return r.UltimateDebtor
}

func (r *RequestedModification4) AddDebtor() *PartyIdentification43 {
	r.Debtor = new(PartyIdentification43)
	return r.Debtor
}

func (r *RequestedModification4) AddDebtorAccount() *CashAccount24 {
	r.DebtorAccount = new(CashAccount24)
	return r.DebtorAccount
}

func (r *RequestedModification4) AddDebtorAgentAccount() *CashAccount24 {
	r.DebtorAgentAccount = new(CashAccount24)
	return r.DebtorAgentAccount
}

func (r *RequestedModification4) AddSettlementInformation() *SettlementInstruction3 {
	r.SettlementInformation = new(SettlementInstruction3)
	return r.SettlementInformation
}

func (r *RequestedModification4) AddCreditorAgentAccount() *CashAccount24 {
	r.CreditorAgentAccount = new(CashAccount24)
	return r.CreditorAgentAccount
}

func (r *RequestedModification4) AddCreditor() *PartyIdentification43 {
	r.Creditor = new(PartyIdentification43)
	return r.Creditor
}

func (r *RequestedModification4) AddCreditorAccount() *CashAccount24 {
	r.CreditorAccount = new(CashAccount24)
	return r.CreditorAccount
}

func (r *RequestedModification4) AddUltimateCreditor() *PartyIdentification43 {
	r.UltimateCreditor = new(PartyIdentification43)
	return r.UltimateCreditor
}

func (r *RequestedModification4) AddPurpose() *Purpose2Choice {
	r.Purpose = new(Purpose2Choice)
	return r.Purpose
}

func (r *RequestedModification4) SetInstructionForDebtorAgent(value string) {
	r.InstructionForDebtorAgent = (*Max140Text)(&value)
}

func (r *RequestedModification4) AddInstructionForNextAgent() *InstructionForNextAgent1 {
	newValue := new(InstructionForNextAgent1)
	r.InstructionForNextAgent = append(r.InstructionForNextAgent, newValue)
	return newValue
}

func (r *RequestedModification4) AddInstructionForCreditorAgent() *InstructionForCreditorAgent1 {
	newValue := new(InstructionForCreditorAgent1)
	r.InstructionForCreditorAgent = append(r.InstructionForCreditorAgent, newValue)
	return newValue
}

func (r *RequestedModification4) AddRemittanceInformation() *RemittanceInformation10 {
	r.RemittanceInformation = new(RemittanceInformation10)
	return r.RemittanceInformation
}

// Provide further details on the requested modifications of the underlying payment instruction.
type RequestedModification5 struct {

	// Unique identification as assigned by an instructing party for an instructed party to unambiguously identify the instruction.
	//
	// Usage: The instruction identification is a point to point reference that can be used between the instructing party and the instructed party to refer to the individual instruction. It can be included in several messages related to the instruction.
	InstructionIdentification *Max35Text `xml:"InstrId,omitempty"`

	// Unique identification, as assigned by the initiating party, to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	//
	// Usage: The end-to-end identification can be used for reconciliation or to link tasks relating to the transaction. It can be included in several messages related to the transaction.
	EndToEndIdentification *Max35Text `xml:"EndToEndId,omitempty"`

	// Unique identification, as assigned by the first instructing agent, to unambiguously identify the transaction that is passed on, unchanged, throughout the entire interbank chain.
	// Usage: The transaction identification can be used for reconciliation, tracking or to link tasks relating to the transaction on the interbank level. The instructing agent has to make sure that the transaction identification is unique for a pre-agreed period.
	TransactionIdentification *Max35Text `xml:"TxId,omitempty"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation25 `xml:"PmtTpInf,omitempty"`

	// Date at which the initiating party requests the clearing agent to process the payment.
	// Usage: This is the date on which the debtor's account is to be debited. If payment by cheque, the date when the cheque must be generated by the bank.
	RequestedExecutionDate *ISODate `xml:"ReqdExctnDt,omitempty"`

	// Date and time at which the creditor requests that the amount of money is to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	Amount *AmountType4Choice `xml:"Amt,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent.
	InterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"IntrBkSttlmAmt,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor. In the context of the payment model, the debtor is also the debit account owner.
	Debtor *PartyIdentification43 `xml:"Dbtr,omitempty"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct,omitempty"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount24 `xml:"DbtrAgtAcct,omitempty"`

	// Instruction between two clearing agents stipulating the cash transfer characteristics between the two parties.
	SettlementInformation *SettlementInstruction3 `xml:"SttlmInf,omitempty"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CreditorAgentAccount *CashAccount24 `xml:"CdtrAgtAcct,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr,omitempty"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`

	// Further information related to the processing of the payment instruction, provided by the initiating party, and intended for the debtor agent.
	InstructionForDebtorAgent *Max140Text `xml:"InstrForDbtrAgt,omitempty"`

	// Further information related to the processing of the payment instruction that may need to be acted upon by the next agent.
	//
	// Usage: The next agent may not be the creditor agent.
	// The instruction can relate to a level of service, can be an instruction that has to be executed by the agent, or can be information required by the next agent.
	InstructionForNextAgent []*InstructionForNextAgent1 `xml:"InstrForNxtAgt,omitempty"`

	// Further information related to the processing of the payment instruction, provided by the initiating party, and intended for the creditor agent.
	InstructionForCreditorAgent []*InstructionForCreditorAgent1 `xml:"InstrForCdtrAgt,omitempty"`

	// Structured information that enables the matching, that is reconciliation, of a payment with the items that the payment is intended to settle, such as commercial invoices in an account receivable system.
	RemittanceInformation *RemittanceInformation11 `xml:"RmtInf,omitempty"`
}

func (r *RequestedModification5) SetInstructionIdentification(value string) {
	r.InstructionIdentification = (*Max35Text)(&value)
}

func (r *RequestedModification5) SetEndToEndIdentification(value string) {
	r.EndToEndIdentification = (*Max35Text)(&value)
}

func (r *RequestedModification5) SetTransactionIdentification(value string) {
	r.TransactionIdentification = (*Max35Text)(&value)
}

func (r *RequestedModification5) AddPaymentTypeInformation() *PaymentTypeInformation25 {
	r.PaymentTypeInformation = new(PaymentTypeInformation25)
	return r.PaymentTypeInformation
}

func (r *RequestedModification5) SetRequestedExecutionDate(value string) {
	r.RequestedExecutionDate = (*ISODate)(&value)
}

func (r *RequestedModification5) SetRequestedCollectionDate(value string) {
	r.RequestedCollectionDate = (*ISODate)(&value)
}

func (r *RequestedModification5) SetInterbankSettlementDate(value string) {
	r.InterbankSettlementDate = (*ISODate)(&value)
}

func (r *RequestedModification5) AddAmount() *AmountType4Choice {
	r.Amount = new(AmountType4Choice)
	return r.Amount
}

func (r *RequestedModification5) SetInterbankSettlementAmount(value, currency string) {
	r.InterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RequestedModification5) SetChargeBearer(value string) {
	r.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (r *RequestedModification5) AddUltimateDebtor() *PartyIdentification43 {
	r.UltimateDebtor = new(PartyIdentification43)
	return r.UltimateDebtor
}

func (r *RequestedModification5) AddDebtor() *PartyIdentification43 {
	r.Debtor = new(PartyIdentification43)
	return r.Debtor
}

func (r *RequestedModification5) AddDebtorAccount() *CashAccount24 {
	r.DebtorAccount = new(CashAccount24)
	return r.DebtorAccount
}

func (r *RequestedModification5) AddDebtorAgentAccount() *CashAccount24 {
	r.DebtorAgentAccount = new(CashAccount24)
	return r.DebtorAgentAccount
}

func (r *RequestedModification5) AddSettlementInformation() *SettlementInstruction3 {
	r.SettlementInformation = new(SettlementInstruction3)
	return r.SettlementInformation
}

func (r *RequestedModification5) AddCreditorAgentAccount() *CashAccount24 {
	r.CreditorAgentAccount = new(CashAccount24)
	return r.CreditorAgentAccount
}

func (r *RequestedModification5) AddCreditor() *PartyIdentification43 {
	r.Creditor = new(PartyIdentification43)
	return r.Creditor
}

func (r *RequestedModification5) AddCreditorAccount() *CashAccount24 {
	r.CreditorAccount = new(CashAccount24)
	return r.CreditorAccount
}

func (r *RequestedModification5) AddUltimateCreditor() *PartyIdentification43 {
	r.UltimateCreditor = new(PartyIdentification43)
	return r.UltimateCreditor
}

func (r *RequestedModification5) AddPurpose() *Purpose2Choice {
	r.Purpose = new(Purpose2Choice)
	return r.Purpose
}

func (r *RequestedModification5) SetInstructionForDebtorAgent(value string) {
	r.InstructionForDebtorAgent = (*Max140Text)(&value)
}

func (r *RequestedModification5) AddInstructionForNextAgent() *InstructionForNextAgent1 {
	newValue := new(InstructionForNextAgent1)
	r.InstructionForNextAgent = append(r.InstructionForNextAgent, newValue)
	return newValue
}

func (r *RequestedModification5) AddInstructionForCreditorAgent() *InstructionForCreditorAgent1 {
	newValue := new(InstructionForCreditorAgent1)
	r.InstructionForCreditorAgent = append(r.InstructionForCreditorAgent, newValue)
	return newValue
}

func (r *RequestedModification5) AddRemittanceInformation() *RemittanceInformation11 {
	r.RemittanceInformation = new(RemittanceInformation11)
	return r.RemittanceInformation
}

// Provide further details on the requested modifications of the underlying payment instruction.
type RequestedModification6 struct {

	// Unique identification as assigned by an instructing party for an instructed party to unambiguously identify the instruction.
	//
	// Usage: The instruction identification is a point to point reference that can be used between the instructing party and the instructed party to refer to the individual instruction. It can be included in several messages related to the instruction.
	InstructionIdentification *Max35Text `xml:"InstrId,omitempty"`

	// Unique identification, as assigned by the initiating party, to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	//
	// Usage: The end-to-end identification can be used for reconciliation or to link tasks relating to the transaction. It can be included in several messages related to the transaction.
	EndToEndIdentification *Max35Text `xml:"EndToEndId,omitempty"`

	// Unique identification, as assigned by the first instructing agent, to unambiguously identify the transaction that is passed on, unchanged, throughout the entire interbank chain.
	// Usage: The transaction identification can be used for reconciliation, tracking or to link tasks relating to the transaction on the interbank level. The instructing agent has to make sure that the transaction identification is unique for a pre-agreed period.
	TransactionIdentification *Max35Text `xml:"TxId,omitempty"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation25 `xml:"PmtTpInf,omitempty"`

	// Date at which the initiating party requests the clearing agent to process the payment.
	// Usage: This is the date on which the debtor's account is to be debited. If payment by cheque, the date when the cheque must be generated by the bank.
	RequestedExecutionDate *DateAndDateTimeChoice `xml:"ReqdExctnDt,omitempty"`

	// Date and time at which the creditor requests that the amount of money is to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	Amount *AmountType4Choice `xml:"Amt,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent.
	InterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"IntrBkSttlmAmt,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor. In the context of the payment model, the debtor is also the debit account owner.
	Debtor *PartyIdentification43 `xml:"Dbtr,omitempty"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct,omitempty"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount24 `xml:"DbtrAgtAcct,omitempty"`

	// Instruction between two clearing agents stipulating the cash transfer characteristics between the two parties.
	SettlementInformation *SettlementInstruction3 `xml:"SttlmInf,omitempty"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CreditorAgentAccount *CashAccount24 `xml:"CdtrAgtAcct,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr,omitempty"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`

	// Further information related to the processing of the payment instruction, provided by the initiating party, and intended for the debtor agent.
	InstructionForDebtorAgent *Max140Text `xml:"InstrForDbtrAgt,omitempty"`

	// Further information related to the processing of the payment instruction that may need to be acted upon by the next agent.
	//
	// Usage: The next agent may not be the creditor agent.
	// The instruction can relate to a level of service, can be an instruction that has to be executed by the agent, or can be information required by the next agent.
	InstructionForNextAgent []*InstructionForNextAgent1 `xml:"InstrForNxtAgt,omitempty"`

	// Further information related to the processing of the payment instruction, provided by the initiating party, and intended for the creditor agent.
	InstructionForCreditorAgent []*InstructionForCreditorAgent1 `xml:"InstrForCdtrAgt,omitempty"`

	// Structured information that enables the matching, that is reconciliation, of a payment with the items that the payment is intended to settle, such as commercial invoices in an account receivable system.
	RemittanceInformation *RemittanceInformation11 `xml:"RmtInf,omitempty"`
}

func (r *RequestedModification6) SetInstructionIdentification(value string) {
	r.InstructionIdentification = (*Max35Text)(&value)
}

func (r *RequestedModification6) SetEndToEndIdentification(value string) {
	r.EndToEndIdentification = (*Max35Text)(&value)
}

func (r *RequestedModification6) SetTransactionIdentification(value string) {
	r.TransactionIdentification = (*Max35Text)(&value)
}

func (r *RequestedModification6) AddPaymentTypeInformation() *PaymentTypeInformation25 {
	r.PaymentTypeInformation = new(PaymentTypeInformation25)
	return r.PaymentTypeInformation
}

func (r *RequestedModification6) AddRequestedExecutionDate() *DateAndDateTimeChoice {
	r.RequestedExecutionDate = new(DateAndDateTimeChoice)
	return r.RequestedExecutionDate
}

func (r *RequestedModification6) SetRequestedCollectionDate(value string) {
	r.RequestedCollectionDate = (*ISODate)(&value)
}

func (r *RequestedModification6) SetInterbankSettlementDate(value string) {
	r.InterbankSettlementDate = (*ISODate)(&value)
}

func (r *RequestedModification6) AddAmount() *AmountType4Choice {
	r.Amount = new(AmountType4Choice)
	return r.Amount
}

func (r *RequestedModification6) SetInterbankSettlementAmount(value, currency string) {
	r.InterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RequestedModification6) SetChargeBearer(value string) {
	r.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (r *RequestedModification6) AddUltimateDebtor() *PartyIdentification43 {
	r.UltimateDebtor = new(PartyIdentification43)
	return r.UltimateDebtor
}

func (r *RequestedModification6) AddDebtor() *PartyIdentification43 {
	r.Debtor = new(PartyIdentification43)
	return r.Debtor
}

func (r *RequestedModification6) AddDebtorAccount() *CashAccount24 {
	r.DebtorAccount = new(CashAccount24)
	return r.DebtorAccount
}

func (r *RequestedModification6) AddDebtorAgentAccount() *CashAccount24 {
	r.DebtorAgentAccount = new(CashAccount24)
	return r.DebtorAgentAccount
}

func (r *RequestedModification6) AddSettlementInformation() *SettlementInstruction3 {
	r.SettlementInformation = new(SettlementInstruction3)
	return r.SettlementInformation
}

func (r *RequestedModification6) AddCreditorAgentAccount() *CashAccount24 {
	r.CreditorAgentAccount = new(CashAccount24)
	return r.CreditorAgentAccount
}

func (r *RequestedModification6) AddCreditor() *PartyIdentification43 {
	r.Creditor = new(PartyIdentification43)
	return r.Creditor
}

func (r *RequestedModification6) AddCreditorAccount() *CashAccount24 {
	r.CreditorAccount = new(CashAccount24)
	return r.CreditorAccount
}

func (r *RequestedModification6) AddUltimateCreditor() *PartyIdentification43 {
	r.UltimateCreditor = new(PartyIdentification43)
	return r.UltimateCreditor
}

func (r *RequestedModification6) AddPurpose() *Purpose2Choice {
	r.Purpose = new(Purpose2Choice)
	return r.Purpose
}

func (r *RequestedModification6) SetInstructionForDebtorAgent(value string) {
	r.InstructionForDebtorAgent = (*Max140Text)(&value)
}

func (r *RequestedModification6) AddInstructionForNextAgent() *InstructionForNextAgent1 {
	newValue := new(InstructionForNextAgent1)
	r.InstructionForNextAgent = append(r.InstructionForNextAgent, newValue)
	return newValue
}

func (r *RequestedModification6) AddInstructionForCreditorAgent() *InstructionForCreditorAgent1 {
	newValue := new(InstructionForCreditorAgent1)
	r.InstructionForCreditorAgent = append(r.InstructionForCreditorAgent, newValue)
	return newValue
}

func (r *RequestedModification6) AddRemittanceInformation() *RemittanceInformation11 {
	r.RemittanceInformation = new(RemittanceInformation11)
	return r.RemittanceInformation
}
