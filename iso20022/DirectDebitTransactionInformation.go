package iso20022

// Set of characteristics that apply to the the direct debit transaction(s).
type DirectDebitTransactionInformation1 struct {

	// Set of elements to reference a payment instruction.
	PaymentIdentification *PaymentIdentification1 `xml:"PmtId"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation2 `xml:"PmtTpInf,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	InstructedAmount *CurrencyAndAmount `xml:"InstdAmt"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Set of elements providing information specific to the direct debit mandate.
	DirectDebitTransaction *DirectDebitTransaction1 `xml:"DrctDbtTx,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification8 `xml:"UltmtCdtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification3 `xml:"DbtrAgt"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount7 `xml:"DbtrAgtAcct,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification8 `xml:"Dbtr"`

	// Identification of the account of the debtor to which a debit entry will be made to execute the transfer.
	DebtorAccount *CashAccount7 `xml:"DbtrAcct"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification8 `xml:"UltmtDbtr,omitempty"`

	// Further information, related to the processing of the payment instruction, that may need to be acted upon by the creditor agent, depending on agreement between creditor and the creditor agent.
	InstructionForCreditorAgent *Max140Text `xml:"InstrForCdtrAgt,omitempty"`

	// Underlying reason for the payment transaction, eg, a charity payment, or a commercial agreement between the creditor and the debtor.
	//
	// Usage: purpose is used by the end-customers, ie originating party, initiating party, debtor, creditor, final party, to provide information concerning the nature of the payment transaction. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose1Choice `xml:"Purp,omitempty"`

	// Information needed due to regulatory and statutory requirements.
	RegulatoryReporting []*RegulatoryReporting2 `xml:"RgltryRptg,omitempty"`

	// Amount of money due to the government or tax authority, according to various pre-defined parameters such as thresholds or income.
	Tax *TaxInformation2 `xml:"Tax,omitempty"`

	// Information related to the handling of the remittance information by any of the agents in the transaction processing chain.
	RelatedRemittanceInformation []*RemittanceLocation1 `xml:"RltdRmtInf,omitempty"`

	// Information that enables the matching, ie, reconciliation, of a payment with the items that the payment is intended to settle, eg, commercial invoices in an account receivable system.
	RemittanceInformation *RemittanceInformation1 `xml:"RmtInf,omitempty"`
}

func (d *DirectDebitTransactionInformation1) AddPaymentIdentification() *PaymentIdentification1 {
	d.PaymentIdentification = new(PaymentIdentification1)
	return d.PaymentIdentification
}

func (d *DirectDebitTransactionInformation1) AddPaymentTypeInformation() *PaymentTypeInformation2 {
	d.PaymentTypeInformation = new(PaymentTypeInformation2)
	return d.PaymentTypeInformation
}

func (d *DirectDebitTransactionInformation1) SetInstructedAmount(value, currency string) {
	d.InstructedAmount = NewCurrencyAndAmount(value, currency)
}

func (d *DirectDebitTransactionInformation1) SetChargeBearer(value string) {
	d.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (d *DirectDebitTransactionInformation1) AddDirectDebitTransaction() *DirectDebitTransaction1 {
	d.DirectDebitTransaction = new(DirectDebitTransaction1)
	return d.DirectDebitTransaction
}

func (d *DirectDebitTransactionInformation1) AddUltimateCreditor() *PartyIdentification8 {
	d.UltimateCreditor = new(PartyIdentification8)
	return d.UltimateCreditor
}

func (d *DirectDebitTransactionInformation1) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification3 {
	d.DebtorAgent = new(BranchAndFinancialInstitutionIdentification3)
	return d.DebtorAgent
}

func (d *DirectDebitTransactionInformation1) AddDebtorAgentAccount() *CashAccount7 {
	d.DebtorAgentAccount = new(CashAccount7)
	return d.DebtorAgentAccount
}

func (d *DirectDebitTransactionInformation1) AddDebtor() *PartyIdentification8 {
	d.Debtor = new(PartyIdentification8)
	return d.Debtor
}

func (d *DirectDebitTransactionInformation1) AddDebtorAccount() *CashAccount7 {
	d.DebtorAccount = new(CashAccount7)
	return d.DebtorAccount
}

func (d *DirectDebitTransactionInformation1) AddUltimateDebtor() *PartyIdentification8 {
	d.UltimateDebtor = new(PartyIdentification8)
	return d.UltimateDebtor
}

func (d *DirectDebitTransactionInformation1) SetInstructionForCreditorAgent(value string) {
	d.InstructionForCreditorAgent = (*Max140Text)(&value)
}

func (d *DirectDebitTransactionInformation1) AddPurpose() *Purpose1Choice {
	d.Purpose = new(Purpose1Choice)
	return d.Purpose
}

func (d *DirectDebitTransactionInformation1) AddRegulatoryReporting() *RegulatoryReporting2 {
	newValue := new(RegulatoryReporting2)
	d.RegulatoryReporting = append(d.RegulatoryReporting, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation1) AddTax() *TaxInformation2 {
	d.Tax = new(TaxInformation2)
	return d.Tax
}

func (d *DirectDebitTransactionInformation1) AddRelatedRemittanceInformation() *RemittanceLocation1 {
	newValue := new(RemittanceLocation1)
	d.RelatedRemittanceInformation = append(d.RelatedRemittanceInformation, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation1) AddRemittanceInformation() *RemittanceInformation1 {
	d.RemittanceInformation = new(RemittanceInformation1)
	return d.RemittanceInformation
}

// Set of elements used to provide information specific to the individual direct debit transaction(s) included in the message.
type DirectDebitTransactionInformation10 struct {

	// Set of elements used to reference a payment instruction.
	PaymentIdentification *PaymentIdentification3 `xml:"PmtId"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation22 `xml:"PmtTpInf,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent.
	InterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"IntrBkSttlmAmt"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	InstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt,omitempty"`

	// Factor used to convert an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr"`

	// Provides information on the charges to be paid by the charge bearer(s) related to the payment transaction.
	ChargesInformation []*ChargesInformation5 `xml:"ChrgsInf,omitempty"`

	// Date and time at which the creditor requests that the amount of money is to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt,omitempty"`

	// Set of elements providing information specific to the direct debit mandate.
	DirectDebitTransaction *DirectDebitTransaction6 `xml:"DrctDbtTx,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification32 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount16 `xml:"CdtrAcct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification4 `xml:"CdtrAgt"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CreditorAgentAccount *CashAccount16 `xml:"CdtrAgtAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification32 `xml:"UltmtCdtr,omitempty"`

	// Party that initiates the payment.
	// Usage: This can be either the creditor or a party that initiates the direct debit on behalf of the creditor.
	InitiatingParty *PartyIdentification32 `xml:"InitgPty,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstdAgt,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If more than one intermediary agent is present, then IntermediaryAgent1 identifies the agent between the DebtorAgent and the IntermediaryAgent2.
	IntermediaryAgent1 *BranchAndFinancialInstitutionIdentification4 `xml:"IntrmyAgt1,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 1 at its servicing agent in the payment chain.
	IntermediaryAgent1Account *CashAccount16 `xml:"IntrmyAgt1Acct,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If more than two intermediary agents are present, then IntermediaryAgent2 identifies the agent between the IntermediaryAgent1 and the IntermediaryAgent3.
	IntermediaryAgent2 *BranchAndFinancialInstitutionIdentification4 `xml:"IntrmyAgt2,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 2 at its servicing agent in the payment chain.
	IntermediaryAgent2Account *CashAccount16 `xml:"IntrmyAgt2Acct,omitempty"`

	// Agent between the debtor agent and creditor agent.
	//
	// Usage: If IntermediaryAgent3 is present, then it identifies the agent between the intermediary agent 2 and the debtor agent.
	IntermediaryAgent3 *BranchAndFinancialInstitutionIdentification4 `xml:"IntrmyAgt3,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 3 at its servicing agent in the payment chain.
	IntermediaryAgent3Account *CashAccount16 `xml:"IntrmyAgt3Acct,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification32 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount16 `xml:"DbtrAcct"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification4 `xml:"DbtrAgt"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount16 `xml:"DbtrAgtAcct,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification32 `xml:"UltmtDbtr,omitempty"`

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`

	// Information needed due to regulatory and statutory requirements.
	RegulatoryReporting []*RegulatoryReporting3 `xml:"RgltryRptg,omitempty"`

	// Set of elements used to provide information related to the handling of the remittance information by any of the agents in the transaction processing chain.
	RelatedRemittanceInformation []*RemittanceLocation2 `xml:"RltdRmtInf,omitempty"`

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RemittanceInformation *RemittanceInformation5 `xml:"RmtInf,omitempty"`
}

func (d *DirectDebitTransactionInformation10) AddPaymentIdentification() *PaymentIdentification3 {
	d.PaymentIdentification = new(PaymentIdentification3)
	return d.PaymentIdentification
}

func (d *DirectDebitTransactionInformation10) AddPaymentTypeInformation() *PaymentTypeInformation22 {
	d.PaymentTypeInformation = new(PaymentTypeInformation22)
	return d.PaymentTypeInformation
}

func (d *DirectDebitTransactionInformation10) SetInterbankSettlementAmount(value, currency string) {
	d.InterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DirectDebitTransactionInformation10) SetInterbankSettlementDate(value string) {
	d.InterbankSettlementDate = (*ISODate)(&value)
}

func (d *DirectDebitTransactionInformation10) SetInstructedAmount(value, currency string) {
	d.InstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (d *DirectDebitTransactionInformation10) SetExchangeRate(value string) {
	d.ExchangeRate = (*BaseOneRate)(&value)
}

func (d *DirectDebitTransactionInformation10) SetChargeBearer(value string) {
	d.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (d *DirectDebitTransactionInformation10) AddChargesInformation() *ChargesInformation5 {
	newValue := new(ChargesInformation5)
	d.ChargesInformation = append(d.ChargesInformation, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation10) SetRequestedCollectionDate(value string) {
	d.RequestedCollectionDate = (*ISODate)(&value)
}

func (d *DirectDebitTransactionInformation10) AddDirectDebitTransaction() *DirectDebitTransaction6 {
	d.DirectDebitTransaction = new(DirectDebitTransaction6)
	return d.DirectDebitTransaction
}

func (d *DirectDebitTransactionInformation10) AddCreditor() *PartyIdentification32 {
	d.Creditor = new(PartyIdentification32)
	return d.Creditor
}

func (d *DirectDebitTransactionInformation10) AddCreditorAccount() *CashAccount16 {
	d.CreditorAccount = new(CashAccount16)
	return d.CreditorAccount
}

func (d *DirectDebitTransactionInformation10) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification4 {
	d.CreditorAgent = new(BranchAndFinancialInstitutionIdentification4)
	return d.CreditorAgent
}

func (d *DirectDebitTransactionInformation10) AddCreditorAgentAccount() *CashAccount16 {
	d.CreditorAgentAccount = new(CashAccount16)
	return d.CreditorAgentAccount
}

func (d *DirectDebitTransactionInformation10) AddUltimateCreditor() *PartyIdentification32 {
	d.UltimateCreditor = new(PartyIdentification32)
	return d.UltimateCreditor
}

func (d *DirectDebitTransactionInformation10) AddInitiatingParty() *PartyIdentification32 {
	d.InitiatingParty = new(PartyIdentification32)
	return d.InitiatingParty
}

func (d *DirectDebitTransactionInformation10) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification4 {
	d.InstructingAgent = new(BranchAndFinancialInstitutionIdentification4)
	return d.InstructingAgent
}

func (d *DirectDebitTransactionInformation10) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification4 {
	d.InstructedAgent = new(BranchAndFinancialInstitutionIdentification4)
	return d.InstructedAgent
}

func (d *DirectDebitTransactionInformation10) AddIntermediaryAgent1() *BranchAndFinancialInstitutionIdentification4 {
	d.IntermediaryAgent1 = new(BranchAndFinancialInstitutionIdentification4)
	return d.IntermediaryAgent1
}

func (d *DirectDebitTransactionInformation10) AddIntermediaryAgent1Account() *CashAccount16 {
	d.IntermediaryAgent1Account = new(CashAccount16)
	return d.IntermediaryAgent1Account
}

func (d *DirectDebitTransactionInformation10) AddIntermediaryAgent2() *BranchAndFinancialInstitutionIdentification4 {
	d.IntermediaryAgent2 = new(BranchAndFinancialInstitutionIdentification4)
	return d.IntermediaryAgent2
}

func (d *DirectDebitTransactionInformation10) AddIntermediaryAgent2Account() *CashAccount16 {
	d.IntermediaryAgent2Account = new(CashAccount16)
	return d.IntermediaryAgent2Account
}

func (d *DirectDebitTransactionInformation10) AddIntermediaryAgent3() *BranchAndFinancialInstitutionIdentification4 {
	d.IntermediaryAgent3 = new(BranchAndFinancialInstitutionIdentification4)
	return d.IntermediaryAgent3
}

func (d *DirectDebitTransactionInformation10) AddIntermediaryAgent3Account() *CashAccount16 {
	d.IntermediaryAgent3Account = new(CashAccount16)
	return d.IntermediaryAgent3Account
}

func (d *DirectDebitTransactionInformation10) AddDebtor() *PartyIdentification32 {
	d.Debtor = new(PartyIdentification32)
	return d.Debtor
}

func (d *DirectDebitTransactionInformation10) AddDebtorAccount() *CashAccount16 {
	d.DebtorAccount = new(CashAccount16)
	return d.DebtorAccount
}

func (d *DirectDebitTransactionInformation10) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification4 {
	d.DebtorAgent = new(BranchAndFinancialInstitutionIdentification4)
	return d.DebtorAgent
}

func (d *DirectDebitTransactionInformation10) AddDebtorAgentAccount() *CashAccount16 {
	d.DebtorAgentAccount = new(CashAccount16)
	return d.DebtorAgentAccount
}

func (d *DirectDebitTransactionInformation10) AddUltimateDebtor() *PartyIdentification32 {
	d.UltimateDebtor = new(PartyIdentification32)
	return d.UltimateDebtor
}

func (d *DirectDebitTransactionInformation10) AddPurpose() *Purpose2Choice {
	d.Purpose = new(Purpose2Choice)
	return d.Purpose
}

func (d *DirectDebitTransactionInformation10) AddRegulatoryReporting() *RegulatoryReporting3 {
	newValue := new(RegulatoryReporting3)
	d.RegulatoryReporting = append(d.RegulatoryReporting, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation10) AddRelatedRemittanceInformation() *RemittanceLocation2 {
	newValue := new(RemittanceLocation2)
	d.RelatedRemittanceInformation = append(d.RelatedRemittanceInformation, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation10) AddRemittanceInformation() *RemittanceInformation5 {
	d.RemittanceInformation = new(RemittanceInformation5)
	return d.RemittanceInformation
}

// Provides further details specific to the individual direct debit transaction(s) included in the message.
type DirectDebitTransactionInformation11 struct {

	// Set of elements used to reference a payment instruction.
	PaymentIdentification *PaymentIdentification1 `xml:"PmtId"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation24 `xml:"PmtTpInf,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	InstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Provides information specific to the direct debit mandate.
	DirectDebitTransaction *DirectDebitTransaction7 `xml:"DrctDbtTx,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount24 `xml:"DbtrAgtAcct,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Further information, related to the processing of the payment instruction, that may need to be acted upon by the creditor agent, depending on agreement between creditor and the creditor agent.
	InstructionForCreditorAgent *Max140Text `xml:"InstrForCdtrAgt,omitempty"`

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`

	// Information needed due to regulatory and statutory requirements.
	RegulatoryReporting []*RegulatoryReporting3 `xml:"RgltryRptg,omitempty"`

	// Provides details on the tax.
	Tax *TaxInformation3 `xml:"Tax,omitempty"`

	// Provides information related to the handling of the remittance information by any of the agents in the transaction processing chain.
	RelatedRemittanceInformation []*RemittanceLocation2 `xml:"RltdRmtInf,omitempty"`

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RemittanceInformation *RemittanceInformation7 `xml:"RmtInf,omitempty"`
}

func (d *DirectDebitTransactionInformation11) AddPaymentIdentification() *PaymentIdentification1 {
	d.PaymentIdentification = new(PaymentIdentification1)
	return d.PaymentIdentification
}

func (d *DirectDebitTransactionInformation11) AddPaymentTypeInformation() *PaymentTypeInformation24 {
	d.PaymentTypeInformation = new(PaymentTypeInformation24)
	return d.PaymentTypeInformation
}

func (d *DirectDebitTransactionInformation11) SetInstructedAmount(value, currency string) {
	d.InstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (d *DirectDebitTransactionInformation11) SetChargeBearer(value string) {
	d.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (d *DirectDebitTransactionInformation11) AddDirectDebitTransaction() *DirectDebitTransaction7 {
	d.DirectDebitTransaction = new(DirectDebitTransaction7)
	return d.DirectDebitTransaction
}

func (d *DirectDebitTransactionInformation11) AddUltimateCreditor() *PartyIdentification43 {
	d.UltimateCreditor = new(PartyIdentification43)
	return d.UltimateCreditor
}

func (d *DirectDebitTransactionInformation11) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	d.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return d.DebtorAgent
}

func (d *DirectDebitTransactionInformation11) AddDebtorAgentAccount() *CashAccount24 {
	d.DebtorAgentAccount = new(CashAccount24)
	return d.DebtorAgentAccount
}

func (d *DirectDebitTransactionInformation11) AddDebtor() *PartyIdentification43 {
	d.Debtor = new(PartyIdentification43)
	return d.Debtor
}

func (d *DirectDebitTransactionInformation11) AddDebtorAccount() *CashAccount24 {
	d.DebtorAccount = new(CashAccount24)
	return d.DebtorAccount
}

func (d *DirectDebitTransactionInformation11) AddUltimateDebtor() *PartyIdentification43 {
	d.UltimateDebtor = new(PartyIdentification43)
	return d.UltimateDebtor
}

func (d *DirectDebitTransactionInformation11) SetInstructionForCreditorAgent(value string) {
	d.InstructionForCreditorAgent = (*Max140Text)(&value)
}

func (d *DirectDebitTransactionInformation11) AddPurpose() *Purpose2Choice {
	d.Purpose = new(Purpose2Choice)
	return d.Purpose
}

func (d *DirectDebitTransactionInformation11) AddRegulatoryReporting() *RegulatoryReporting3 {
	newValue := new(RegulatoryReporting3)
	d.RegulatoryReporting = append(d.RegulatoryReporting, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation11) AddTax() *TaxInformation3 {
	d.Tax = new(TaxInformation3)
	return d.Tax
}

func (d *DirectDebitTransactionInformation11) AddRelatedRemittanceInformation() *RemittanceLocation2 {
	newValue := new(RemittanceLocation2)
	d.RelatedRemittanceInformation = append(d.RelatedRemittanceInformation, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation11) AddRemittanceInformation() *RemittanceInformation7 {
	d.RemittanceInformation = new(RemittanceInformation7)
	return d.RemittanceInformation
}

// Provides further details specific to the individual direct debit transaction(s) included in the message.
type DirectDebitTransactionInformation12 struct {

	// Set of elements used to reference a payment instruction.
	PaymentIdentification *PaymentIdentification3 `xml:"PmtId"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation25 `xml:"PmtTpInf,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent.
	InterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"IntrBkSttlmAmt"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	InstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt,omitempty"`

	// Factor used to convert an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr"`

	// Provides information on the charges to be paid by the charge bearer(s) related to the payment transaction.
	ChargesInformation []*Charges2 `xml:"ChrgsInf,omitempty"`

	// Date and time at which the creditor requests that the amount of money is to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt,omitempty"`

	// Provides information specific to the direct debit mandate.
	DirectDebitTransaction *DirectDebitTransaction7 `xml:"DrctDbtTx,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CreditorAgentAccount *CashAccount24 `xml:"CdtrAgtAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Party that initiates the payment.
	// Usage: This can be either the creditor or a party that initiates the direct debit on behalf of the creditor.
	InitiatingParty *PartyIdentification43 `xml:"InitgPty,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If more than one intermediary agent is present, then IntermediaryAgent1 identifies the agent between the DebtorAgent and the IntermediaryAgent2.
	IntermediaryAgent1 *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt1,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 1 at its servicing agent in the payment chain.
	IntermediaryAgent1Account *CashAccount24 `xml:"IntrmyAgt1Acct,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If more than two intermediary agents are present, then IntermediaryAgent2 identifies the agent between the IntermediaryAgent1 and the IntermediaryAgent3.
	IntermediaryAgent2 *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt2,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 2 at its servicing agent in the payment chain.
	IntermediaryAgent2Account *CashAccount24 `xml:"IntrmyAgt2Acct,omitempty"`

	// Agent between the debtor agent and creditor agent.
	//
	// Usage: If IntermediaryAgent3 is present, then it identifies the agent between the intermediary agent 2 and the debtor agent.
	IntermediaryAgent3 *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt3,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 3 at its servicing agent in the payment chain.
	IntermediaryAgent3Account *CashAccount24 `xml:"IntrmyAgt3Acct,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount24 `xml:"DbtrAgtAcct,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`

	// Information needed due to regulatory and statutory requirements.
	RegulatoryReporting []*RegulatoryReporting3 `xml:"RgltryRptg,omitempty"`

	// Provides information related to the handling of the remittance information by any of the agents in the transaction processing chain.
	RelatedRemittanceInformation []*RemittanceLocation2 `xml:"RltdRmtInf,omitempty"`

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RemittanceInformation *RemittanceInformation7 `xml:"RmtInf,omitempty"`
}

func (d *DirectDebitTransactionInformation12) AddPaymentIdentification() *PaymentIdentification3 {
	d.PaymentIdentification = new(PaymentIdentification3)
	return d.PaymentIdentification
}

func (d *DirectDebitTransactionInformation12) AddPaymentTypeInformation() *PaymentTypeInformation25 {
	d.PaymentTypeInformation = new(PaymentTypeInformation25)
	return d.PaymentTypeInformation
}

func (d *DirectDebitTransactionInformation12) SetInterbankSettlementAmount(value, currency string) {
	d.InterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DirectDebitTransactionInformation12) SetInterbankSettlementDate(value string) {
	d.InterbankSettlementDate = (*ISODate)(&value)
}

func (d *DirectDebitTransactionInformation12) SetInstructedAmount(value, currency string) {
	d.InstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (d *DirectDebitTransactionInformation12) SetExchangeRate(value string) {
	d.ExchangeRate = (*BaseOneRate)(&value)
}

func (d *DirectDebitTransactionInformation12) SetChargeBearer(value string) {
	d.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (d *DirectDebitTransactionInformation12) AddChargesInformation() *Charges2 {
	newValue := new(Charges2)
	d.ChargesInformation = append(d.ChargesInformation, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation12) SetRequestedCollectionDate(value string) {
	d.RequestedCollectionDate = (*ISODate)(&value)
}

func (d *DirectDebitTransactionInformation12) AddDirectDebitTransaction() *DirectDebitTransaction7 {
	d.DirectDebitTransaction = new(DirectDebitTransaction7)
	return d.DirectDebitTransaction
}

func (d *DirectDebitTransactionInformation12) AddCreditor() *PartyIdentification43 {
	d.Creditor = new(PartyIdentification43)
	return d.Creditor
}

func (d *DirectDebitTransactionInformation12) AddCreditorAccount() *CashAccount24 {
	d.CreditorAccount = new(CashAccount24)
	return d.CreditorAccount
}

func (d *DirectDebitTransactionInformation12) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	d.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return d.CreditorAgent
}

func (d *DirectDebitTransactionInformation12) AddCreditorAgentAccount() *CashAccount24 {
	d.CreditorAgentAccount = new(CashAccount24)
	return d.CreditorAgentAccount
}

func (d *DirectDebitTransactionInformation12) AddUltimateCreditor() *PartyIdentification43 {
	d.UltimateCreditor = new(PartyIdentification43)
	return d.UltimateCreditor
}

func (d *DirectDebitTransactionInformation12) AddInitiatingParty() *PartyIdentification43 {
	d.InitiatingParty = new(PartyIdentification43)
	return d.InitiatingParty
}

func (d *DirectDebitTransactionInformation12) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	d.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return d.InstructingAgent
}

func (d *DirectDebitTransactionInformation12) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	d.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return d.InstructedAgent
}

func (d *DirectDebitTransactionInformation12) AddIntermediaryAgent1() *BranchAndFinancialInstitutionIdentification5 {
	d.IntermediaryAgent1 = new(BranchAndFinancialInstitutionIdentification5)
	return d.IntermediaryAgent1
}

func (d *DirectDebitTransactionInformation12) AddIntermediaryAgent1Account() *CashAccount24 {
	d.IntermediaryAgent1Account = new(CashAccount24)
	return d.IntermediaryAgent1Account
}

func (d *DirectDebitTransactionInformation12) AddIntermediaryAgent2() *BranchAndFinancialInstitutionIdentification5 {
	d.IntermediaryAgent2 = new(BranchAndFinancialInstitutionIdentification5)
	return d.IntermediaryAgent2
}

func (d *DirectDebitTransactionInformation12) AddIntermediaryAgent2Account() *CashAccount24 {
	d.IntermediaryAgent2Account = new(CashAccount24)
	return d.IntermediaryAgent2Account
}

func (d *DirectDebitTransactionInformation12) AddIntermediaryAgent3() *BranchAndFinancialInstitutionIdentification5 {
	d.IntermediaryAgent3 = new(BranchAndFinancialInstitutionIdentification5)
	return d.IntermediaryAgent3
}

func (d *DirectDebitTransactionInformation12) AddIntermediaryAgent3Account() *CashAccount24 {
	d.IntermediaryAgent3Account = new(CashAccount24)
	return d.IntermediaryAgent3Account
}

func (d *DirectDebitTransactionInformation12) AddDebtor() *PartyIdentification43 {
	d.Debtor = new(PartyIdentification43)
	return d.Debtor
}

func (d *DirectDebitTransactionInformation12) AddDebtorAccount() *CashAccount24 {
	d.DebtorAccount = new(CashAccount24)
	return d.DebtorAccount
}

func (d *DirectDebitTransactionInformation12) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	d.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return d.DebtorAgent
}

func (d *DirectDebitTransactionInformation12) AddDebtorAgentAccount() *CashAccount24 {
	d.DebtorAgentAccount = new(CashAccount24)
	return d.DebtorAgentAccount
}

func (d *DirectDebitTransactionInformation12) AddUltimateDebtor() *PartyIdentification43 {
	d.UltimateDebtor = new(PartyIdentification43)
	return d.UltimateDebtor
}

func (d *DirectDebitTransactionInformation12) AddPurpose() *Purpose2Choice {
	d.Purpose = new(Purpose2Choice)
	return d.Purpose
}

func (d *DirectDebitTransactionInformation12) AddRegulatoryReporting() *RegulatoryReporting3 {
	newValue := new(RegulatoryReporting3)
	d.RegulatoryReporting = append(d.RegulatoryReporting, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation12) AddRelatedRemittanceInformation() *RemittanceLocation2 {
	newValue := new(RemittanceLocation2)
	d.RelatedRemittanceInformation = append(d.RelatedRemittanceInformation, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation12) AddRemittanceInformation() *RemittanceInformation7 {
	d.RemittanceInformation = new(RemittanceInformation7)
	return d.RemittanceInformation
}

// Provides further details specific to the individual direct debit transaction(s) included in the message.
type DirectDebitTransactionInformation13 struct {

	// Set of elements used to reference a payment instruction.
	PaymentIdentification *PaymentIdentification1 `xml:"PmtId"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation24 `xml:"PmtTpInf,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	InstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Provides information specific to the direct debit mandate.
	DirectDebitTransaction *DirectDebitTransaction7 `xml:"DrctDbtTx,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount24 `xml:"DbtrAgtAcct,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Further information, related to the processing of the payment instruction, that may need to be acted upon by the creditor agent, depending on agreement between creditor and the creditor agent.
	InstructionForCreditorAgent *Max140Text `xml:"InstrForCdtrAgt,omitempty"`

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`

	// Information needed due to regulatory and statutory requirements.
	RegulatoryReporting []*RegulatoryReporting3 `xml:"RgltryRptg,omitempty"`

	// Provides details on the tax.
	Tax *TaxInformation3 `xml:"Tax,omitempty"`

	// Provides information related to the handling of the remittance information by any of the agents in the transaction processing chain.
	RelatedRemittanceInformation []*RemittanceLocation2 `xml:"RltdRmtInf,omitempty"`

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RemittanceInformation *RemittanceInformation7 `xml:"RmtInf,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (d *DirectDebitTransactionInformation13) AddPaymentIdentification() *PaymentIdentification1 {
	d.PaymentIdentification = new(PaymentIdentification1)
	return d.PaymentIdentification
}

func (d *DirectDebitTransactionInformation13) AddPaymentTypeInformation() *PaymentTypeInformation24 {
	d.PaymentTypeInformation = new(PaymentTypeInformation24)
	return d.PaymentTypeInformation
}

func (d *DirectDebitTransactionInformation13) SetInstructedAmount(value, currency string) {
	d.InstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (d *DirectDebitTransactionInformation13) SetChargeBearer(value string) {
	d.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (d *DirectDebitTransactionInformation13) AddDirectDebitTransaction() *DirectDebitTransaction7 {
	d.DirectDebitTransaction = new(DirectDebitTransaction7)
	return d.DirectDebitTransaction
}

func (d *DirectDebitTransactionInformation13) AddUltimateCreditor() *PartyIdentification43 {
	d.UltimateCreditor = new(PartyIdentification43)
	return d.UltimateCreditor
}

func (d *DirectDebitTransactionInformation13) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	d.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return d.DebtorAgent
}

func (d *DirectDebitTransactionInformation13) AddDebtorAgentAccount() *CashAccount24 {
	d.DebtorAgentAccount = new(CashAccount24)
	return d.DebtorAgentAccount
}

func (d *DirectDebitTransactionInformation13) AddDebtor() *PartyIdentification43 {
	d.Debtor = new(PartyIdentification43)
	return d.Debtor
}

func (d *DirectDebitTransactionInformation13) AddDebtorAccount() *CashAccount24 {
	d.DebtorAccount = new(CashAccount24)
	return d.DebtorAccount
}

func (d *DirectDebitTransactionInformation13) AddUltimateDebtor() *PartyIdentification43 {
	d.UltimateDebtor = new(PartyIdentification43)
	return d.UltimateDebtor
}

func (d *DirectDebitTransactionInformation13) SetInstructionForCreditorAgent(value string) {
	d.InstructionForCreditorAgent = (*Max140Text)(&value)
}

func (d *DirectDebitTransactionInformation13) AddPurpose() *Purpose2Choice {
	d.Purpose = new(Purpose2Choice)
	return d.Purpose
}

func (d *DirectDebitTransactionInformation13) AddRegulatoryReporting() *RegulatoryReporting3 {
	newValue := new(RegulatoryReporting3)
	d.RegulatoryReporting = append(d.RegulatoryReporting, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation13) AddTax() *TaxInformation3 {
	d.Tax = new(TaxInformation3)
	return d.Tax
}

func (d *DirectDebitTransactionInformation13) AddRelatedRemittanceInformation() *RemittanceLocation2 {
	newValue := new(RemittanceLocation2)
	d.RelatedRemittanceInformation = append(d.RelatedRemittanceInformation, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation13) AddRemittanceInformation() *RemittanceInformation7 {
	d.RemittanceInformation = new(RemittanceInformation7)
	return d.RemittanceInformation
}

func (d *DirectDebitTransactionInformation13) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	d.SupplementaryData = append(d.SupplementaryData, newValue)
	return newValue
}

// Provides further details specific to the individual direct debit transaction(s) included in the message.
type DirectDebitTransactionInformation14 struct {

	// Set of elements used to reference a payment instruction.
	PaymentIdentification *PaymentIdentification3 `xml:"PmtId"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation25 `xml:"PmtTpInf,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent.
	InterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"IntrBkSttlmAmt"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	InstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt,omitempty"`

	// Factor used to convert an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr"`

	// Provides information on the charges to be paid by the charge bearer(s) related to the payment transaction.
	ChargesInformation []*Charges2 `xml:"ChrgsInf,omitempty"`

	// Date and time at which the creditor requests that the amount of money is to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt,omitempty"`

	// Provides information specific to the direct debit mandate.
	DirectDebitTransaction *DirectDebitTransaction7 `xml:"DrctDbtTx,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CreditorAgentAccount *CashAccount24 `xml:"CdtrAgtAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Party that initiates the payment.
	// Usage: This can be either the creditor or a party that initiates the direct debit on behalf of the creditor.
	InitiatingParty *PartyIdentification43 `xml:"InitgPty,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If more than one intermediary agent is present, then IntermediaryAgent1 identifies the agent between the DebtorAgent and the IntermediaryAgent2.
	IntermediaryAgent1 *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt1,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 1 at its servicing agent in the payment chain.
	IntermediaryAgent1Account *CashAccount24 `xml:"IntrmyAgt1Acct,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If more than two intermediary agents are present, then IntermediaryAgent2 identifies the agent between the IntermediaryAgent1 and the IntermediaryAgent3.
	IntermediaryAgent2 *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt2,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 2 at its servicing agent in the payment chain.
	IntermediaryAgent2Account *CashAccount24 `xml:"IntrmyAgt2Acct,omitempty"`

	// Agent between the debtor agent and creditor agent.
	//
	// Usage: If IntermediaryAgent3 is present, then it identifies the agent between the intermediary agent 2 and the debtor agent.
	IntermediaryAgent3 *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt3,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 3 at its servicing agent in the payment chain.
	IntermediaryAgent3Account *CashAccount24 `xml:"IntrmyAgt3Acct,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount24 `xml:"DbtrAgtAcct,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`

	// Information needed due to regulatory and statutory requirements.
	RegulatoryReporting []*RegulatoryReporting3 `xml:"RgltryRptg,omitempty"`

	// Provides information related to the handling of the remittance information by any of the agents in the transaction processing chain.
	RelatedRemittanceInformation []*RemittanceLocation2 `xml:"RltdRmtInf,omitempty"`

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RemittanceInformation *RemittanceInformation7 `xml:"RmtInf,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (d *DirectDebitTransactionInformation14) AddPaymentIdentification() *PaymentIdentification3 {
	d.PaymentIdentification = new(PaymentIdentification3)
	return d.PaymentIdentification
}

func (d *DirectDebitTransactionInformation14) AddPaymentTypeInformation() *PaymentTypeInformation25 {
	d.PaymentTypeInformation = new(PaymentTypeInformation25)
	return d.PaymentTypeInformation
}

func (d *DirectDebitTransactionInformation14) SetInterbankSettlementAmount(value, currency string) {
	d.InterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DirectDebitTransactionInformation14) SetInterbankSettlementDate(value string) {
	d.InterbankSettlementDate = (*ISODate)(&value)
}

func (d *DirectDebitTransactionInformation14) SetInstructedAmount(value, currency string) {
	d.InstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (d *DirectDebitTransactionInformation14) SetExchangeRate(value string) {
	d.ExchangeRate = (*BaseOneRate)(&value)
}

func (d *DirectDebitTransactionInformation14) SetChargeBearer(value string) {
	d.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (d *DirectDebitTransactionInformation14) AddChargesInformation() *Charges2 {
	newValue := new(Charges2)
	d.ChargesInformation = append(d.ChargesInformation, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation14) SetRequestedCollectionDate(value string) {
	d.RequestedCollectionDate = (*ISODate)(&value)
}

func (d *DirectDebitTransactionInformation14) AddDirectDebitTransaction() *DirectDebitTransaction7 {
	d.DirectDebitTransaction = new(DirectDebitTransaction7)
	return d.DirectDebitTransaction
}

func (d *DirectDebitTransactionInformation14) AddCreditor() *PartyIdentification43 {
	d.Creditor = new(PartyIdentification43)
	return d.Creditor
}

func (d *DirectDebitTransactionInformation14) AddCreditorAccount() *CashAccount24 {
	d.CreditorAccount = new(CashAccount24)
	return d.CreditorAccount
}

func (d *DirectDebitTransactionInformation14) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	d.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return d.CreditorAgent
}

func (d *DirectDebitTransactionInformation14) AddCreditorAgentAccount() *CashAccount24 {
	d.CreditorAgentAccount = new(CashAccount24)
	return d.CreditorAgentAccount
}

func (d *DirectDebitTransactionInformation14) AddUltimateCreditor() *PartyIdentification43 {
	d.UltimateCreditor = new(PartyIdentification43)
	return d.UltimateCreditor
}

func (d *DirectDebitTransactionInformation14) AddInitiatingParty() *PartyIdentification43 {
	d.InitiatingParty = new(PartyIdentification43)
	return d.InitiatingParty
}

func (d *DirectDebitTransactionInformation14) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	d.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return d.InstructingAgent
}

func (d *DirectDebitTransactionInformation14) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	d.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return d.InstructedAgent
}

func (d *DirectDebitTransactionInformation14) AddIntermediaryAgent1() *BranchAndFinancialInstitutionIdentification5 {
	d.IntermediaryAgent1 = new(BranchAndFinancialInstitutionIdentification5)
	return d.IntermediaryAgent1
}

func (d *DirectDebitTransactionInformation14) AddIntermediaryAgent1Account() *CashAccount24 {
	d.IntermediaryAgent1Account = new(CashAccount24)
	return d.IntermediaryAgent1Account
}

func (d *DirectDebitTransactionInformation14) AddIntermediaryAgent2() *BranchAndFinancialInstitutionIdentification5 {
	d.IntermediaryAgent2 = new(BranchAndFinancialInstitutionIdentification5)
	return d.IntermediaryAgent2
}

func (d *DirectDebitTransactionInformation14) AddIntermediaryAgent2Account() *CashAccount24 {
	d.IntermediaryAgent2Account = new(CashAccount24)
	return d.IntermediaryAgent2Account
}

func (d *DirectDebitTransactionInformation14) AddIntermediaryAgent3() *BranchAndFinancialInstitutionIdentification5 {
	d.IntermediaryAgent3 = new(BranchAndFinancialInstitutionIdentification5)
	return d.IntermediaryAgent3
}

func (d *DirectDebitTransactionInformation14) AddIntermediaryAgent3Account() *CashAccount24 {
	d.IntermediaryAgent3Account = new(CashAccount24)
	return d.IntermediaryAgent3Account
}

func (d *DirectDebitTransactionInformation14) AddDebtor() *PartyIdentification43 {
	d.Debtor = new(PartyIdentification43)
	return d.Debtor
}

func (d *DirectDebitTransactionInformation14) AddDebtorAccount() *CashAccount24 {
	d.DebtorAccount = new(CashAccount24)
	return d.DebtorAccount
}

func (d *DirectDebitTransactionInformation14) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	d.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return d.DebtorAgent
}

func (d *DirectDebitTransactionInformation14) AddDebtorAgentAccount() *CashAccount24 {
	d.DebtorAgentAccount = new(CashAccount24)
	return d.DebtorAgentAccount
}

func (d *DirectDebitTransactionInformation14) AddUltimateDebtor() *PartyIdentification43 {
	d.UltimateDebtor = new(PartyIdentification43)
	return d.UltimateDebtor
}

func (d *DirectDebitTransactionInformation14) AddPurpose() *Purpose2Choice {
	d.Purpose = new(Purpose2Choice)
	return d.Purpose
}

func (d *DirectDebitTransactionInformation14) AddRegulatoryReporting() *RegulatoryReporting3 {
	newValue := new(RegulatoryReporting3)
	d.RegulatoryReporting = append(d.RegulatoryReporting, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation14) AddRelatedRemittanceInformation() *RemittanceLocation2 {
	newValue := new(RemittanceLocation2)
	d.RelatedRemittanceInformation = append(d.RelatedRemittanceInformation, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation14) AddRemittanceInformation() *RemittanceInformation7 {
	d.RemittanceInformation = new(RemittanceInformation7)
	return d.RemittanceInformation
}

func (d *DirectDebitTransactionInformation14) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	d.SupplementaryData = append(d.SupplementaryData, newValue)
	return newValue
}

// Provides further details specific to the individual direct debit transaction(s) included in the message.
type DirectDebitTransactionInformation15 struct {

	// References used for  a payment instruction.
	PaymentIdentification *PaymentIdentification3 `xml:"PmtId"`

	// Specifies the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation21 `xml:"PmtTpInf,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent.
	InterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"IntrBkSttlmAmt"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the settlement instruction.
	SettlementPriority *Priority3Code `xml:"SttlmPrty,omitempty"`

	// Provides information on the requested settlement time(s) of the payment instruction.
	SettlementTimeRequest *SettlementTimeRequest2 `xml:"SttlmTmReq,omitempty"`

	// Ultimate financial institution that owes an amount of money to the (ultimate) institutional creditor.
	UltimateDebtor *BranchAndFinancialInstitutionIdentification5 `xml:"UltmtDbtr,omitempty"`

	// Financial institution that owes an amount of money to the (ultimate) financial institutional creditor.
	Debtor *BranchAndFinancialInstitutionIdentification5 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount24 `xml:"DbtrAgtAcct,omitempty"`

	// Further information related to the processing of the payment instruction, that may need to be acted upon by the debtor agent, depending on agreement between debtor and the debtor agent.
	//
	InstructionForDebtorAgent *Max210Text `xml:"InstrForDbtrAgt,omitempty"`

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RemittanceInformation *RemittanceInformation2 `xml:"RmtInf,omitempty"`
}

func (d *DirectDebitTransactionInformation15) AddPaymentIdentification() *PaymentIdentification3 {
	d.PaymentIdentification = new(PaymentIdentification3)
	return d.PaymentIdentification
}

func (d *DirectDebitTransactionInformation15) AddPaymentTypeInformation() *PaymentTypeInformation21 {
	d.PaymentTypeInformation = new(PaymentTypeInformation21)
	return d.PaymentTypeInformation
}

func (d *DirectDebitTransactionInformation15) SetInterbankSettlementAmount(value, currency string) {
	d.InterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DirectDebitTransactionInformation15) SetInterbankSettlementDate(value string) {
	d.InterbankSettlementDate = (*ISODate)(&value)
}

func (d *DirectDebitTransactionInformation15) SetSettlementPriority(value string) {
	d.SettlementPriority = (*Priority3Code)(&value)
}

func (d *DirectDebitTransactionInformation15) AddSettlementTimeRequest() *SettlementTimeRequest2 {
	d.SettlementTimeRequest = new(SettlementTimeRequest2)
	return d.SettlementTimeRequest
}

func (d *DirectDebitTransactionInformation15) AddUltimateDebtor() *BranchAndFinancialInstitutionIdentification5 {
	d.UltimateDebtor = new(BranchAndFinancialInstitutionIdentification5)
	return d.UltimateDebtor
}

func (d *DirectDebitTransactionInformation15) AddDebtor() *BranchAndFinancialInstitutionIdentification5 {
	d.Debtor = new(BranchAndFinancialInstitutionIdentification5)
	return d.Debtor
}

func (d *DirectDebitTransactionInformation15) AddDebtorAccount() *CashAccount24 {
	d.DebtorAccount = new(CashAccount24)
	return d.DebtorAccount
}

func (d *DirectDebitTransactionInformation15) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	d.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return d.DebtorAgent
}

func (d *DirectDebitTransactionInformation15) AddDebtorAgentAccount() *CashAccount24 {
	d.DebtorAgentAccount = new(CashAccount24)
	return d.DebtorAgentAccount
}

func (d *DirectDebitTransactionInformation15) SetInstructionForDebtorAgent(value string) {
	d.InstructionForDebtorAgent = (*Max210Text)(&value)
}

func (d *DirectDebitTransactionInformation15) AddRemittanceInformation() *RemittanceInformation2 {
	d.RemittanceInformation = new(RemittanceInformation2)
	return d.RemittanceInformation
}

// Provides further details specific to the individual direct debit transaction(s) included in the message.
type DirectDebitTransactionInformation17 struct {

	// Set of elements used to reference a payment instruction.
	PaymentIdentification *PaymentIdentification3 `xml:"PmtId"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation25 `xml:"PmtTpInf,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent.
	InterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"IntrBkSttlmAmt"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the settlement instruction.
	SettlementPriority *Priority3Code `xml:"SttlmPrty,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	// Usage: This amount has to be transported unchanged through the transaction chain.
	InstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt,omitempty"`

	// Factor used to convert an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr"`

	// Provides information on the charges to be paid by the charge bearer(s) related to the payment transaction.
	ChargesInformation []*Charges2 `xml:"ChrgsInf,omitempty"`

	// Date and time at which the creditor requests that the amount of money is to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt,omitempty"`

	// Provides information specific to the direct debit mandate.
	DirectDebitTransaction *DirectDebitTransaction8 `xml:"DrctDbtTx,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CreditorAgentAccount *CashAccount24 `xml:"CdtrAgtAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Party that initiates the payment.
	// Usage: This can be either the creditor or a party that initiates the direct debit on behalf of the creditor.
	InitiatingParty *PartyIdentification43 `xml:"InitgPty,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If more than one intermediary agent is present, then IntermediaryAgent1 identifies the agent between the DebtorAgent and the IntermediaryAgent2.
	IntermediaryAgent1 *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt1,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 1 at its servicing agent in the payment chain.
	IntermediaryAgent1Account *CashAccount24 `xml:"IntrmyAgt1Acct,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If more than two intermediary agents are present, then IntermediaryAgent2 identifies the agent between the IntermediaryAgent1 and the IntermediaryAgent3.
	IntermediaryAgent2 *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt2,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 2 at its servicing agent in the payment chain.
	IntermediaryAgent2Account *CashAccount24 `xml:"IntrmyAgt2Acct,omitempty"`

	// Agent between the debtor agent and creditor agent.
	//
	// Usage: If IntermediaryAgent3 is present, then it identifies the agent between the intermediary agent 2 and the debtor agent.
	IntermediaryAgent3 *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt3,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 3 at its servicing agent in the payment chain.
	IntermediaryAgent3Account *CashAccount24 `xml:"IntrmyAgt3Acct,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount24 `xml:"DbtrAgtAcct,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`

	// Information needed due to regulatory and statutory requirements.
	RegulatoryReporting []*RegulatoryReporting3 `xml:"RgltryRptg,omitempty"`

	// Provides information related to the handling of the remittance information by any of the agents in the transaction processing chain.
	RelatedRemittanceInformation []*RemittanceLocation4 `xml:"RltdRmtInf,omitempty"`

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RemittanceInformation *RemittanceInformation10 `xml:"RmtInf,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (d *DirectDebitTransactionInformation17) AddPaymentIdentification() *PaymentIdentification3 {
	d.PaymentIdentification = new(PaymentIdentification3)
	return d.PaymentIdentification
}

func (d *DirectDebitTransactionInformation17) AddPaymentTypeInformation() *PaymentTypeInformation25 {
	d.PaymentTypeInformation = new(PaymentTypeInformation25)
	return d.PaymentTypeInformation
}

func (d *DirectDebitTransactionInformation17) SetInterbankSettlementAmount(value, currency string) {
	d.InterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DirectDebitTransactionInformation17) SetInterbankSettlementDate(value string) {
	d.InterbankSettlementDate = (*ISODate)(&value)
}

func (d *DirectDebitTransactionInformation17) SetSettlementPriority(value string) {
	d.SettlementPriority = (*Priority3Code)(&value)
}

func (d *DirectDebitTransactionInformation17) SetInstructedAmount(value, currency string) {
	d.InstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (d *DirectDebitTransactionInformation17) SetExchangeRate(value string) {
	d.ExchangeRate = (*BaseOneRate)(&value)
}

func (d *DirectDebitTransactionInformation17) SetChargeBearer(value string) {
	d.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (d *DirectDebitTransactionInformation17) AddChargesInformation() *Charges2 {
	newValue := new(Charges2)
	d.ChargesInformation = append(d.ChargesInformation, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation17) SetRequestedCollectionDate(value string) {
	d.RequestedCollectionDate = (*ISODate)(&value)
}

func (d *DirectDebitTransactionInformation17) AddDirectDebitTransaction() *DirectDebitTransaction8 {
	d.DirectDebitTransaction = new(DirectDebitTransaction8)
	return d.DirectDebitTransaction
}

func (d *DirectDebitTransactionInformation17) AddCreditor() *PartyIdentification43 {
	d.Creditor = new(PartyIdentification43)
	return d.Creditor
}

func (d *DirectDebitTransactionInformation17) AddCreditorAccount() *CashAccount24 {
	d.CreditorAccount = new(CashAccount24)
	return d.CreditorAccount
}

func (d *DirectDebitTransactionInformation17) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	d.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return d.CreditorAgent
}

func (d *DirectDebitTransactionInformation17) AddCreditorAgentAccount() *CashAccount24 {
	d.CreditorAgentAccount = new(CashAccount24)
	return d.CreditorAgentAccount
}

func (d *DirectDebitTransactionInformation17) AddUltimateCreditor() *PartyIdentification43 {
	d.UltimateCreditor = new(PartyIdentification43)
	return d.UltimateCreditor
}

func (d *DirectDebitTransactionInformation17) AddInitiatingParty() *PartyIdentification43 {
	d.InitiatingParty = new(PartyIdentification43)
	return d.InitiatingParty
}

func (d *DirectDebitTransactionInformation17) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	d.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return d.InstructingAgent
}

func (d *DirectDebitTransactionInformation17) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	d.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return d.InstructedAgent
}

func (d *DirectDebitTransactionInformation17) AddIntermediaryAgent1() *BranchAndFinancialInstitutionIdentification5 {
	d.IntermediaryAgent1 = new(BranchAndFinancialInstitutionIdentification5)
	return d.IntermediaryAgent1
}

func (d *DirectDebitTransactionInformation17) AddIntermediaryAgent1Account() *CashAccount24 {
	d.IntermediaryAgent1Account = new(CashAccount24)
	return d.IntermediaryAgent1Account
}

func (d *DirectDebitTransactionInformation17) AddIntermediaryAgent2() *BranchAndFinancialInstitutionIdentification5 {
	d.IntermediaryAgent2 = new(BranchAndFinancialInstitutionIdentification5)
	return d.IntermediaryAgent2
}

func (d *DirectDebitTransactionInformation17) AddIntermediaryAgent2Account() *CashAccount24 {
	d.IntermediaryAgent2Account = new(CashAccount24)
	return d.IntermediaryAgent2Account
}

func (d *DirectDebitTransactionInformation17) AddIntermediaryAgent3() *BranchAndFinancialInstitutionIdentification5 {
	d.IntermediaryAgent3 = new(BranchAndFinancialInstitutionIdentification5)
	return d.IntermediaryAgent3
}

func (d *DirectDebitTransactionInformation17) AddIntermediaryAgent3Account() *CashAccount24 {
	d.IntermediaryAgent3Account = new(CashAccount24)
	return d.IntermediaryAgent3Account
}

func (d *DirectDebitTransactionInformation17) AddDebtor() *PartyIdentification43 {
	d.Debtor = new(PartyIdentification43)
	return d.Debtor
}

func (d *DirectDebitTransactionInformation17) AddDebtorAccount() *CashAccount24 {
	d.DebtorAccount = new(CashAccount24)
	return d.DebtorAccount
}

func (d *DirectDebitTransactionInformation17) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	d.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return d.DebtorAgent
}

func (d *DirectDebitTransactionInformation17) AddDebtorAgentAccount() *CashAccount24 {
	d.DebtorAgentAccount = new(CashAccount24)
	return d.DebtorAgentAccount
}

func (d *DirectDebitTransactionInformation17) AddUltimateDebtor() *PartyIdentification43 {
	d.UltimateDebtor = new(PartyIdentification43)
	return d.UltimateDebtor
}

func (d *DirectDebitTransactionInformation17) AddPurpose() *Purpose2Choice {
	d.Purpose = new(Purpose2Choice)
	return d.Purpose
}

func (d *DirectDebitTransactionInformation17) AddRegulatoryReporting() *RegulatoryReporting3 {
	newValue := new(RegulatoryReporting3)
	d.RegulatoryReporting = append(d.RegulatoryReporting, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation17) AddRelatedRemittanceInformation() *RemittanceLocation4 {
	newValue := new(RemittanceLocation4)
	d.RelatedRemittanceInformation = append(d.RelatedRemittanceInformation, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation17) AddRemittanceInformation() *RemittanceInformation10 {
	d.RemittanceInformation = new(RemittanceInformation10)
	return d.RemittanceInformation
}

func (d *DirectDebitTransactionInformation17) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	d.SupplementaryData = append(d.SupplementaryData, newValue)
	return newValue
}

// Provides further details specific to the individual direct debit transaction(s) included in the message.
type DirectDebitTransactionInformation18 struct {

	// Set of elements used to reference a payment instruction.
	PaymentIdentification *PaymentIdentification1 `xml:"PmtId"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation24 `xml:"PmtTpInf,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	// Usage: This amount has to be transported unchanged through the transaction chain.
	InstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Provides information specific to the direct debit mandate.
	DirectDebitTransaction *DirectDebitTransaction8 `xml:"DrctDbtTx,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount24 `xml:"DbtrAgtAcct,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Further information, related to the processing of the payment instruction, that may need to be acted upon by the creditor agent, depending on agreement between creditor and the creditor agent.
	InstructionForCreditorAgent *Max140Text `xml:"InstrForCdtrAgt,omitempty"`

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`

	// Information needed due to regulatory and statutory requirements.
	RegulatoryReporting []*RegulatoryReporting3 `xml:"RgltryRptg,omitempty"`

	// Provides details on the tax.
	Tax *TaxInformation3 `xml:"Tax,omitempty"`

	// Provides information related to the handling of the remittance information by any of the agents in the transaction processing chain.
	RelatedRemittanceInformation []*RemittanceLocation4 `xml:"RltdRmtInf,omitempty"`

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RemittanceInformation *RemittanceInformation10 `xml:"RmtInf,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (d *DirectDebitTransactionInformation18) AddPaymentIdentification() *PaymentIdentification1 {
	d.PaymentIdentification = new(PaymentIdentification1)
	return d.PaymentIdentification
}

func (d *DirectDebitTransactionInformation18) AddPaymentTypeInformation() *PaymentTypeInformation24 {
	d.PaymentTypeInformation = new(PaymentTypeInformation24)
	return d.PaymentTypeInformation
}

func (d *DirectDebitTransactionInformation18) SetInstructedAmount(value, currency string) {
	d.InstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (d *DirectDebitTransactionInformation18) SetChargeBearer(value string) {
	d.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (d *DirectDebitTransactionInformation18) AddDirectDebitTransaction() *DirectDebitTransaction8 {
	d.DirectDebitTransaction = new(DirectDebitTransaction8)
	return d.DirectDebitTransaction
}

func (d *DirectDebitTransactionInformation18) AddUltimateCreditor() *PartyIdentification43 {
	d.UltimateCreditor = new(PartyIdentification43)
	return d.UltimateCreditor
}

func (d *DirectDebitTransactionInformation18) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	d.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return d.DebtorAgent
}

func (d *DirectDebitTransactionInformation18) AddDebtorAgentAccount() *CashAccount24 {
	d.DebtorAgentAccount = new(CashAccount24)
	return d.DebtorAgentAccount
}

func (d *DirectDebitTransactionInformation18) AddDebtor() *PartyIdentification43 {
	d.Debtor = new(PartyIdentification43)
	return d.Debtor
}

func (d *DirectDebitTransactionInformation18) AddDebtorAccount() *CashAccount24 {
	d.DebtorAccount = new(CashAccount24)
	return d.DebtorAccount
}

func (d *DirectDebitTransactionInformation18) AddUltimateDebtor() *PartyIdentification43 {
	d.UltimateDebtor = new(PartyIdentification43)
	return d.UltimateDebtor
}

func (d *DirectDebitTransactionInformation18) SetInstructionForCreditorAgent(value string) {
	d.InstructionForCreditorAgent = (*Max140Text)(&value)
}

func (d *DirectDebitTransactionInformation18) AddPurpose() *Purpose2Choice {
	d.Purpose = new(Purpose2Choice)
	return d.Purpose
}

func (d *DirectDebitTransactionInformation18) AddRegulatoryReporting() *RegulatoryReporting3 {
	newValue := new(RegulatoryReporting3)
	d.RegulatoryReporting = append(d.RegulatoryReporting, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation18) AddTax() *TaxInformation3 {
	d.Tax = new(TaxInformation3)
	return d.Tax
}

func (d *DirectDebitTransactionInformation18) AddRelatedRemittanceInformation() *RemittanceLocation4 {
	newValue := new(RemittanceLocation4)
	d.RelatedRemittanceInformation = append(d.RelatedRemittanceInformation, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation18) AddRemittanceInformation() *RemittanceInformation10 {
	d.RemittanceInformation = new(RemittanceInformation10)
	return d.RemittanceInformation
}

func (d *DirectDebitTransactionInformation18) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	d.SupplementaryData = append(d.SupplementaryData, newValue)
	return newValue
}

// Provides further details specific to the individual direct debit transaction(s) included in the message.
type DirectDebitTransactionInformation19 struct {

	// Set of elements used to reference a payment instruction.
	PaymentIdentification *PaymentIdentification1 `xml:"PmtId"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation24 `xml:"PmtTpInf,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	// Usage: This amount has to be transported unchanged through the transaction chain.
	InstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Provides information specific to the direct debit mandate.
	DirectDebitTransaction *DirectDebitTransaction8 `xml:"DrctDbtTx,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount24 `xml:"DbtrAgtAcct,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Further information, related to the processing of the payment instruction, that may need to be acted upon by the creditor agent, depending on agreement between creditor and the creditor agent.
	InstructionForCreditorAgent *Max140Text `xml:"InstrForCdtrAgt,omitempty"`

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`

	// Information needed due to regulatory and statutory requirements.
	RegulatoryReporting []*RegulatoryReporting3 `xml:"RgltryRptg,omitempty"`

	// Provides details on the tax.
	Tax *TaxInformation3 `xml:"Tax,omitempty"`

	// Provides information related to the handling of the remittance information by any of the agents in the transaction processing chain.
	RelatedRemittanceInformation []*RemittanceLocation4 `xml:"RltdRmtInf,omitempty"`

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RemittanceInformation *RemittanceInformation11 `xml:"RmtInf,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (d *DirectDebitTransactionInformation19) AddPaymentIdentification() *PaymentIdentification1 {
	d.PaymentIdentification = new(PaymentIdentification1)
	return d.PaymentIdentification
}

func (d *DirectDebitTransactionInformation19) AddPaymentTypeInformation() *PaymentTypeInformation24 {
	d.PaymentTypeInformation = new(PaymentTypeInformation24)
	return d.PaymentTypeInformation
}

func (d *DirectDebitTransactionInformation19) SetInstructedAmount(value, currency string) {
	d.InstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (d *DirectDebitTransactionInformation19) SetChargeBearer(value string) {
	d.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (d *DirectDebitTransactionInformation19) AddDirectDebitTransaction() *DirectDebitTransaction8 {
	d.DirectDebitTransaction = new(DirectDebitTransaction8)
	return d.DirectDebitTransaction
}

func (d *DirectDebitTransactionInformation19) AddUltimateCreditor() *PartyIdentification43 {
	d.UltimateCreditor = new(PartyIdentification43)
	return d.UltimateCreditor
}

func (d *DirectDebitTransactionInformation19) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	d.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return d.DebtorAgent
}

func (d *DirectDebitTransactionInformation19) AddDebtorAgentAccount() *CashAccount24 {
	d.DebtorAgentAccount = new(CashAccount24)
	return d.DebtorAgentAccount
}

func (d *DirectDebitTransactionInformation19) AddDebtor() *PartyIdentification43 {
	d.Debtor = new(PartyIdentification43)
	return d.Debtor
}

func (d *DirectDebitTransactionInformation19) AddDebtorAccount() *CashAccount24 {
	d.DebtorAccount = new(CashAccount24)
	return d.DebtorAccount
}

func (d *DirectDebitTransactionInformation19) AddUltimateDebtor() *PartyIdentification43 {
	d.UltimateDebtor = new(PartyIdentification43)
	return d.UltimateDebtor
}

func (d *DirectDebitTransactionInformation19) SetInstructionForCreditorAgent(value string) {
	d.InstructionForCreditorAgent = (*Max140Text)(&value)
}

func (d *DirectDebitTransactionInformation19) AddPurpose() *Purpose2Choice {
	d.Purpose = new(Purpose2Choice)
	return d.Purpose
}

func (d *DirectDebitTransactionInformation19) AddRegulatoryReporting() *RegulatoryReporting3 {
	newValue := new(RegulatoryReporting3)
	d.RegulatoryReporting = append(d.RegulatoryReporting, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation19) AddTax() *TaxInformation3 {
	d.Tax = new(TaxInformation3)
	return d.Tax
}

func (d *DirectDebitTransactionInformation19) AddRelatedRemittanceInformation() *RemittanceLocation4 {
	newValue := new(RemittanceLocation4)
	d.RelatedRemittanceInformation = append(d.RelatedRemittanceInformation, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation19) AddRemittanceInformation() *RemittanceInformation11 {
	d.RemittanceInformation = new(RemittanceInformation11)
	return d.RemittanceInformation
}

func (d *DirectDebitTransactionInformation19) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	d.SupplementaryData = append(d.SupplementaryData, newValue)
	return newValue
}

// Set of characteristics that apply to the the direct debit transaction(s).
type DirectDebitTransactionInformation2 struct {

	// Set of elements to reference a payment instruction.
	PaymentIdentification *PaymentIdentification2 `xml:"PmtId"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation4 `xml:"PmtTpInf,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent.
	InterbankSettlementAmount *CurrencyAndAmount `xml:"IntrBkSttlmAmt"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	InstructedAmount *CurrencyAndAmount `xml:"InstdAmt,omitempty"`

	// The factor used for conversion of an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr"`

	// Provides information on the charges related to the payment transaction.
	ChargesInformation []*ChargesInformation1 `xml:"ChrgsInf,omitempty"`

	// Date at which the creditor requests the amount of money to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt,omitempty"`

	// Set of elements providing information specific to the direct debit mandate.
	DirectDebitTransaction *DirectDebitTransaction1 `xml:"DrctDbtTx,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification8 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount7 `xml:"CdtrAcct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification3 `xml:"CdtrAgt"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CreditorAgentAccount *CashAccount7 `xml:"CdtrAgtAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification8 `xml:"UltmtCdtr,omitempty"`

	// Party that initiates the payment. In the payment context, this can either be the debtor (in a credit transfer), the creditor (in a direct debit), or a party that initiates the payment on behalf of the debtor or creditor.
	InitiatingParty *PartyIdentification8 `xml:"InitgPty,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstdAgt,omitempty"`

	// Agent between the debtor agent and creditor agent.
	//
	// Usage: If more than one intermediary agent is present, then IntermediaryAgent1 identifies the agent between the creditor agent and the intermediary agent 2.
	IntermediaryAgent1 *BranchAndFinancialInstitutionIdentification3 `xml:"IntrmyAgt1,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 1 at its servicing agent in the payment chain.
	IntermediaryAgent1Account *CashAccount7 `xml:"IntrmyAgt1Acct,omitempty"`

	// Agent between the debtor agent and creditor agent.
	//
	// Usage: If more than two intermediary agents are present, then IntermediaryAgent2 identifies the agent between the intermediary agent 1 and the intermediary agent 3.
	IntermediaryAgent2 *BranchAndFinancialInstitutionIdentification3 `xml:"IntrmyAgt2,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 2 at its servicing agent in the payment chain.
	IntermediaryAgent2Account *CashAccount7 `xml:"IntrmyAgt2Acct,omitempty"`

	// Agent between the debtor agent and creditor agent.
	//
	// Usage: If IntermediaryAgent3 is present, then it identifies the agent between the intermediary agent 2 and the debtor agent.
	IntermediaryAgent3 *BranchAndFinancialInstitutionIdentification3 `xml:"IntrmyAgt3,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 3 at its servicing agent in the payment chain.
	IntermediaryAgent3Account *CashAccount7 `xml:"IntrmyAgt3Acct,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification8 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount7 `xml:"DbtrAcct"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification3 `xml:"DbtrAgt"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount7 `xml:"DbtrAgtAcct,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification8 `xml:"UltmtDbtr,omitempty"`

	// Underlying reason for the payment transaction.
	//
	// Usage: Purpose is used by the end-customers, i.e. initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose1Choice `xml:"Purp,omitempty"`

	// Information needed due to regulatory and statutory requirements.
	RegulatoryReporting []*RegulatoryReporting2 `xml:"RgltryRptg,omitempty"`

	// Information related to the handling of the remittance information by any of the agents in the transaction processing chain.
	RelatedRemittanceInformation []*RemittanceLocation1 `xml:"RltdRmtInf,omitempty"`

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RemittanceInformation *RemittanceInformation1 `xml:"RmtInf,omitempty"`
}

func (d *DirectDebitTransactionInformation2) AddPaymentIdentification() *PaymentIdentification2 {
	d.PaymentIdentification = new(PaymentIdentification2)
	return d.PaymentIdentification
}

func (d *DirectDebitTransactionInformation2) AddPaymentTypeInformation() *PaymentTypeInformation4 {
	d.PaymentTypeInformation = new(PaymentTypeInformation4)
	return d.PaymentTypeInformation
}

func (d *DirectDebitTransactionInformation2) SetInterbankSettlementAmount(value, currency string) {
	d.InterbankSettlementAmount = NewCurrencyAndAmount(value, currency)
}

func (d *DirectDebitTransactionInformation2) SetInterbankSettlementDate(value string) {
	d.InterbankSettlementDate = (*ISODate)(&value)
}

func (d *DirectDebitTransactionInformation2) SetInstructedAmount(value, currency string) {
	d.InstructedAmount = NewCurrencyAndAmount(value, currency)
}

func (d *DirectDebitTransactionInformation2) SetExchangeRate(value string) {
	d.ExchangeRate = (*BaseOneRate)(&value)
}

func (d *DirectDebitTransactionInformation2) SetChargeBearer(value string) {
	d.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (d *DirectDebitTransactionInformation2) AddChargesInformation() *ChargesInformation1 {
	newValue := new(ChargesInformation1)
	d.ChargesInformation = append(d.ChargesInformation, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation2) SetRequestedCollectionDate(value string) {
	d.RequestedCollectionDate = (*ISODate)(&value)
}

func (d *DirectDebitTransactionInformation2) AddDirectDebitTransaction() *DirectDebitTransaction1 {
	d.DirectDebitTransaction = new(DirectDebitTransaction1)
	return d.DirectDebitTransaction
}

func (d *DirectDebitTransactionInformation2) AddCreditor() *PartyIdentification8 {
	d.Creditor = new(PartyIdentification8)
	return d.Creditor
}

func (d *DirectDebitTransactionInformation2) AddCreditorAccount() *CashAccount7 {
	d.CreditorAccount = new(CashAccount7)
	return d.CreditorAccount
}

func (d *DirectDebitTransactionInformation2) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification3 {
	d.CreditorAgent = new(BranchAndFinancialInstitutionIdentification3)
	return d.CreditorAgent
}

func (d *DirectDebitTransactionInformation2) AddCreditorAgentAccount() *CashAccount7 {
	d.CreditorAgentAccount = new(CashAccount7)
	return d.CreditorAgentAccount
}

func (d *DirectDebitTransactionInformation2) AddUltimateCreditor() *PartyIdentification8 {
	d.UltimateCreditor = new(PartyIdentification8)
	return d.UltimateCreditor
}

func (d *DirectDebitTransactionInformation2) AddInitiatingParty() *PartyIdentification8 {
	d.InitiatingParty = new(PartyIdentification8)
	return d.InitiatingParty
}

func (d *DirectDebitTransactionInformation2) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification3 {
	d.InstructingAgent = new(BranchAndFinancialInstitutionIdentification3)
	return d.InstructingAgent
}

func (d *DirectDebitTransactionInformation2) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification3 {
	d.InstructedAgent = new(BranchAndFinancialInstitutionIdentification3)
	return d.InstructedAgent
}

func (d *DirectDebitTransactionInformation2) AddIntermediaryAgent1() *BranchAndFinancialInstitutionIdentification3 {
	d.IntermediaryAgent1 = new(BranchAndFinancialInstitutionIdentification3)
	return d.IntermediaryAgent1
}

func (d *DirectDebitTransactionInformation2) AddIntermediaryAgent1Account() *CashAccount7 {
	d.IntermediaryAgent1Account = new(CashAccount7)
	return d.IntermediaryAgent1Account
}

func (d *DirectDebitTransactionInformation2) AddIntermediaryAgent2() *BranchAndFinancialInstitutionIdentification3 {
	d.IntermediaryAgent2 = new(BranchAndFinancialInstitutionIdentification3)
	return d.IntermediaryAgent2
}

func (d *DirectDebitTransactionInformation2) AddIntermediaryAgent2Account() *CashAccount7 {
	d.IntermediaryAgent2Account = new(CashAccount7)
	return d.IntermediaryAgent2Account
}

func (d *DirectDebitTransactionInformation2) AddIntermediaryAgent3() *BranchAndFinancialInstitutionIdentification3 {
	d.IntermediaryAgent3 = new(BranchAndFinancialInstitutionIdentification3)
	return d.IntermediaryAgent3
}

func (d *DirectDebitTransactionInformation2) AddIntermediaryAgent3Account() *CashAccount7 {
	d.IntermediaryAgent3Account = new(CashAccount7)
	return d.IntermediaryAgent3Account
}

func (d *DirectDebitTransactionInformation2) AddDebtor() *PartyIdentification8 {
	d.Debtor = new(PartyIdentification8)
	return d.Debtor
}

func (d *DirectDebitTransactionInformation2) AddDebtorAccount() *CashAccount7 {
	d.DebtorAccount = new(CashAccount7)
	return d.DebtorAccount
}

func (d *DirectDebitTransactionInformation2) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification3 {
	d.DebtorAgent = new(BranchAndFinancialInstitutionIdentification3)
	return d.DebtorAgent
}

func (d *DirectDebitTransactionInformation2) AddDebtorAgentAccount() *CashAccount7 {
	d.DebtorAgentAccount = new(CashAccount7)
	return d.DebtorAgentAccount
}

func (d *DirectDebitTransactionInformation2) AddUltimateDebtor() *PartyIdentification8 {
	d.UltimateDebtor = new(PartyIdentification8)
	return d.UltimateDebtor
}

func (d *DirectDebitTransactionInformation2) AddPurpose() *Purpose1Choice {
	d.Purpose = new(Purpose1Choice)
	return d.Purpose
}

func (d *DirectDebitTransactionInformation2) AddRegulatoryReporting() *RegulatoryReporting2 {
	newValue := new(RegulatoryReporting2)
	d.RegulatoryReporting = append(d.RegulatoryReporting, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation2) AddRelatedRemittanceInformation() *RemittanceLocation1 {
	newValue := new(RemittanceLocation1)
	d.RelatedRemittanceInformation = append(d.RelatedRemittanceInformation, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation2) AddRemittanceInformation() *RemittanceInformation1 {
	d.RemittanceInformation = new(RemittanceInformation1)
	return d.RemittanceInformation
}

// Provides further details specific to the individual direct debit transaction(s) included in the message.
type DirectDebitTransactionInformation20 struct {

	// Set of elements used to reference a payment instruction.
	PaymentIdentification *PaymentIdentification3 `xml:"PmtId"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation25 `xml:"PmtTpInf,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent.
	InterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"IntrBkSttlmAmt"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the settlement instruction.
	SettlementPriority *Priority3Code `xml:"SttlmPrty,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	// Usage: This amount has to be transported unchanged through the transaction chain.
	InstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt,omitempty"`

	// Factor used to convert an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr"`

	// Provides information on the charges to be paid by the charge bearer(s) related to the payment transaction.
	ChargesInformation []*Charges2 `xml:"ChrgsInf,omitempty"`

	// Date and time at which the creditor requests that the amount of money is to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt,omitempty"`

	// Provides information specific to the direct debit mandate.
	DirectDebitTransaction *DirectDebitTransaction8 `xml:"DrctDbtTx,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CreditorAgentAccount *CashAccount24 `xml:"CdtrAgtAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Party that initiates the payment.
	// Usage: This can be either the creditor or a party that initiates the direct debit on behalf of the creditor.
	InitiatingParty *PartyIdentification43 `xml:"InitgPty,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If more than one intermediary agent is present, then IntermediaryAgent1 identifies the agent between the DebtorAgent and the IntermediaryAgent2.
	IntermediaryAgent1 *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt1,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 1 at its servicing agent in the payment chain.
	IntermediaryAgent1Account *CashAccount24 `xml:"IntrmyAgt1Acct,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If more than two intermediary agents are present, then IntermediaryAgent2 identifies the agent between the IntermediaryAgent1 and the IntermediaryAgent3.
	IntermediaryAgent2 *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt2,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 2 at its servicing agent in the payment chain.
	IntermediaryAgent2Account *CashAccount24 `xml:"IntrmyAgt2Acct,omitempty"`

	// Agent between the debtor agent and creditor agent.
	//
	// Usage: If IntermediaryAgent3 is present, then it identifies the agent between the intermediary agent 2 and the debtor agent.
	IntermediaryAgent3 *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt3,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 3 at its servicing agent in the payment chain.
	IntermediaryAgent3Account *CashAccount24 `xml:"IntrmyAgt3Acct,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount24 `xml:"DbtrAgtAcct,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`

	// Information needed due to regulatory and statutory requirements.
	RegulatoryReporting []*RegulatoryReporting3 `xml:"RgltryRptg,omitempty"`

	// Provides information related to the handling of the remittance information by any of the agents in the transaction processing chain.
	RelatedRemittanceInformation []*RemittanceLocation4 `xml:"RltdRmtInf,omitempty"`

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RemittanceInformation *RemittanceInformation11 `xml:"RmtInf,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (d *DirectDebitTransactionInformation20) AddPaymentIdentification() *PaymentIdentification3 {
	d.PaymentIdentification = new(PaymentIdentification3)
	return d.PaymentIdentification
}

func (d *DirectDebitTransactionInformation20) AddPaymentTypeInformation() *PaymentTypeInformation25 {
	d.PaymentTypeInformation = new(PaymentTypeInformation25)
	return d.PaymentTypeInformation
}

func (d *DirectDebitTransactionInformation20) SetInterbankSettlementAmount(value, currency string) {
	d.InterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DirectDebitTransactionInformation20) SetInterbankSettlementDate(value string) {
	d.InterbankSettlementDate = (*ISODate)(&value)
}

func (d *DirectDebitTransactionInformation20) SetSettlementPriority(value string) {
	d.SettlementPriority = (*Priority3Code)(&value)
}

func (d *DirectDebitTransactionInformation20) SetInstructedAmount(value, currency string) {
	d.InstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (d *DirectDebitTransactionInformation20) SetExchangeRate(value string) {
	d.ExchangeRate = (*BaseOneRate)(&value)
}

func (d *DirectDebitTransactionInformation20) SetChargeBearer(value string) {
	d.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (d *DirectDebitTransactionInformation20) AddChargesInformation() *Charges2 {
	newValue := new(Charges2)
	d.ChargesInformation = append(d.ChargesInformation, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation20) SetRequestedCollectionDate(value string) {
	d.RequestedCollectionDate = (*ISODate)(&value)
}

func (d *DirectDebitTransactionInformation20) AddDirectDebitTransaction() *DirectDebitTransaction8 {
	d.DirectDebitTransaction = new(DirectDebitTransaction8)
	return d.DirectDebitTransaction
}

func (d *DirectDebitTransactionInformation20) AddCreditor() *PartyIdentification43 {
	d.Creditor = new(PartyIdentification43)
	return d.Creditor
}

func (d *DirectDebitTransactionInformation20) AddCreditorAccount() *CashAccount24 {
	d.CreditorAccount = new(CashAccount24)
	return d.CreditorAccount
}

func (d *DirectDebitTransactionInformation20) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	d.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return d.CreditorAgent
}

func (d *DirectDebitTransactionInformation20) AddCreditorAgentAccount() *CashAccount24 {
	d.CreditorAgentAccount = new(CashAccount24)
	return d.CreditorAgentAccount
}

func (d *DirectDebitTransactionInformation20) AddUltimateCreditor() *PartyIdentification43 {
	d.UltimateCreditor = new(PartyIdentification43)
	return d.UltimateCreditor
}

func (d *DirectDebitTransactionInformation20) AddInitiatingParty() *PartyIdentification43 {
	d.InitiatingParty = new(PartyIdentification43)
	return d.InitiatingParty
}

func (d *DirectDebitTransactionInformation20) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	d.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return d.InstructingAgent
}

func (d *DirectDebitTransactionInformation20) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	d.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return d.InstructedAgent
}

func (d *DirectDebitTransactionInformation20) AddIntermediaryAgent1() *BranchAndFinancialInstitutionIdentification5 {
	d.IntermediaryAgent1 = new(BranchAndFinancialInstitutionIdentification5)
	return d.IntermediaryAgent1
}

func (d *DirectDebitTransactionInformation20) AddIntermediaryAgent1Account() *CashAccount24 {
	d.IntermediaryAgent1Account = new(CashAccount24)
	return d.IntermediaryAgent1Account
}

func (d *DirectDebitTransactionInformation20) AddIntermediaryAgent2() *BranchAndFinancialInstitutionIdentification5 {
	d.IntermediaryAgent2 = new(BranchAndFinancialInstitutionIdentification5)
	return d.IntermediaryAgent2
}

func (d *DirectDebitTransactionInformation20) AddIntermediaryAgent2Account() *CashAccount24 {
	d.IntermediaryAgent2Account = new(CashAccount24)
	return d.IntermediaryAgent2Account
}

func (d *DirectDebitTransactionInformation20) AddIntermediaryAgent3() *BranchAndFinancialInstitutionIdentification5 {
	d.IntermediaryAgent3 = new(BranchAndFinancialInstitutionIdentification5)
	return d.IntermediaryAgent3
}

func (d *DirectDebitTransactionInformation20) AddIntermediaryAgent3Account() *CashAccount24 {
	d.IntermediaryAgent3Account = new(CashAccount24)
	return d.IntermediaryAgent3Account
}

func (d *DirectDebitTransactionInformation20) AddDebtor() *PartyIdentification43 {
	d.Debtor = new(PartyIdentification43)
	return d.Debtor
}

func (d *DirectDebitTransactionInformation20) AddDebtorAccount() *CashAccount24 {
	d.DebtorAccount = new(CashAccount24)
	return d.DebtorAccount
}

func (d *DirectDebitTransactionInformation20) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	d.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return d.DebtorAgent
}

func (d *DirectDebitTransactionInformation20) AddDebtorAgentAccount() *CashAccount24 {
	d.DebtorAgentAccount = new(CashAccount24)
	return d.DebtorAgentAccount
}

func (d *DirectDebitTransactionInformation20) AddUltimateDebtor() *PartyIdentification43 {
	d.UltimateDebtor = new(PartyIdentification43)
	return d.UltimateDebtor
}

func (d *DirectDebitTransactionInformation20) AddPurpose() *Purpose2Choice {
	d.Purpose = new(Purpose2Choice)
	return d.Purpose
}

func (d *DirectDebitTransactionInformation20) AddRegulatoryReporting() *RegulatoryReporting3 {
	newValue := new(RegulatoryReporting3)
	d.RegulatoryReporting = append(d.RegulatoryReporting, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation20) AddRelatedRemittanceInformation() *RemittanceLocation4 {
	newValue := new(RemittanceLocation4)
	d.RelatedRemittanceInformation = append(d.RelatedRemittanceInformation, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation20) AddRemittanceInformation() *RemittanceInformation11 {
	d.RemittanceInformation = new(RemittanceInformation11)
	return d.RemittanceInformation
}

func (d *DirectDebitTransactionInformation20) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	d.SupplementaryData = append(d.SupplementaryData, newValue)
	return newValue
}

// Provides further details specific to the individual direct debit transaction(s) included in the message.
type DirectDebitTransactionInformation21 struct {

	// Set of elements used to reference a payment instruction.
	PaymentIdentification *PaymentIdentification3 `xml:"PmtId"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation25 `xml:"PmtTpInf,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent.
	InterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"IntrBkSttlmAmt"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the settlement instruction.
	SettlementPriority *Priority3Code `xml:"SttlmPrty,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	// Usage: This amount has to be transported unchanged through the transaction chain.
	InstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt,omitempty"`

	// Factor used to convert an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr"`

	// Provides information on the charges to be paid by the charge bearer(s) related to the payment transaction.
	ChargesInformation []*Charges2 `xml:"ChrgsInf,omitempty"`

	// Date and time at which the creditor requests that the amount of money is to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt,omitempty"`

	// Provides information specific to the direct debit mandate.
	DirectDebitTransaction *DirectDebitTransaction9 `xml:"DrctDbtTx,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CreditorAgentAccount *CashAccount24 `xml:"CdtrAgtAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Party that initiates the payment.
	// Usage: This can be either the creditor or a party that initiates the direct debit on behalf of the creditor.
	InitiatingParty *PartyIdentification43 `xml:"InitgPty,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If more than one intermediary agent is present, then IntermediaryAgent1 identifies the agent between the DebtorAgent and the IntermediaryAgent2.
	IntermediaryAgent1 *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt1,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 1 at its servicing agent in the payment chain.
	IntermediaryAgent1Account *CashAccount24 `xml:"IntrmyAgt1Acct,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If more than two intermediary agents are present, then IntermediaryAgent2 identifies the agent between the IntermediaryAgent1 and the IntermediaryAgent3.
	IntermediaryAgent2 *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt2,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 2 at its servicing agent in the payment chain.
	IntermediaryAgent2Account *CashAccount24 `xml:"IntrmyAgt2Acct,omitempty"`

	// Agent between the debtor agent and creditor agent.
	//
	// Usage: If IntermediaryAgent3 is present, then it identifies the agent between the intermediary agent 2 and the debtor agent.
	IntermediaryAgent3 *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt3,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 3 at its servicing agent in the payment chain.
	IntermediaryAgent3Account *CashAccount24 `xml:"IntrmyAgt3Acct,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount24 `xml:"DbtrAgtAcct,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`

	// Information needed due to regulatory and statutory requirements.
	RegulatoryReporting []*RegulatoryReporting3 `xml:"RgltryRptg,omitempty"`

	// Provides information related to the handling of the remittance information by any of the agents in the transaction processing chain.
	RelatedRemittanceInformation []*RemittanceLocation4 `xml:"RltdRmtInf,omitempty"`

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RemittanceInformation *RemittanceInformation11 `xml:"RmtInf,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (d *DirectDebitTransactionInformation21) AddPaymentIdentification() *PaymentIdentification3 {
	d.PaymentIdentification = new(PaymentIdentification3)
	return d.PaymentIdentification
}

func (d *DirectDebitTransactionInformation21) AddPaymentTypeInformation() *PaymentTypeInformation25 {
	d.PaymentTypeInformation = new(PaymentTypeInformation25)
	return d.PaymentTypeInformation
}

func (d *DirectDebitTransactionInformation21) SetInterbankSettlementAmount(value, currency string) {
	d.InterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DirectDebitTransactionInformation21) SetInterbankSettlementDate(value string) {
	d.InterbankSettlementDate = (*ISODate)(&value)
}

func (d *DirectDebitTransactionInformation21) SetSettlementPriority(value string) {
	d.SettlementPriority = (*Priority3Code)(&value)
}

func (d *DirectDebitTransactionInformation21) SetInstructedAmount(value, currency string) {
	d.InstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (d *DirectDebitTransactionInformation21) SetExchangeRate(value string) {
	d.ExchangeRate = (*BaseOneRate)(&value)
}

func (d *DirectDebitTransactionInformation21) SetChargeBearer(value string) {
	d.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (d *DirectDebitTransactionInformation21) AddChargesInformation() *Charges2 {
	newValue := new(Charges2)
	d.ChargesInformation = append(d.ChargesInformation, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation21) SetRequestedCollectionDate(value string) {
	d.RequestedCollectionDate = (*ISODate)(&value)
}

func (d *DirectDebitTransactionInformation21) AddDirectDebitTransaction() *DirectDebitTransaction9 {
	d.DirectDebitTransaction = new(DirectDebitTransaction9)
	return d.DirectDebitTransaction
}

func (d *DirectDebitTransactionInformation21) AddCreditor() *PartyIdentification43 {
	d.Creditor = new(PartyIdentification43)
	return d.Creditor
}

func (d *DirectDebitTransactionInformation21) AddCreditorAccount() *CashAccount24 {
	d.CreditorAccount = new(CashAccount24)
	return d.CreditorAccount
}

func (d *DirectDebitTransactionInformation21) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	d.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return d.CreditorAgent
}

func (d *DirectDebitTransactionInformation21) AddCreditorAgentAccount() *CashAccount24 {
	d.CreditorAgentAccount = new(CashAccount24)
	return d.CreditorAgentAccount
}

func (d *DirectDebitTransactionInformation21) AddUltimateCreditor() *PartyIdentification43 {
	d.UltimateCreditor = new(PartyIdentification43)
	return d.UltimateCreditor
}

func (d *DirectDebitTransactionInformation21) AddInitiatingParty() *PartyIdentification43 {
	d.InitiatingParty = new(PartyIdentification43)
	return d.InitiatingParty
}

func (d *DirectDebitTransactionInformation21) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	d.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return d.InstructingAgent
}

func (d *DirectDebitTransactionInformation21) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	d.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return d.InstructedAgent
}

func (d *DirectDebitTransactionInformation21) AddIntermediaryAgent1() *BranchAndFinancialInstitutionIdentification5 {
	d.IntermediaryAgent1 = new(BranchAndFinancialInstitutionIdentification5)
	return d.IntermediaryAgent1
}

func (d *DirectDebitTransactionInformation21) AddIntermediaryAgent1Account() *CashAccount24 {
	d.IntermediaryAgent1Account = new(CashAccount24)
	return d.IntermediaryAgent1Account
}

func (d *DirectDebitTransactionInformation21) AddIntermediaryAgent2() *BranchAndFinancialInstitutionIdentification5 {
	d.IntermediaryAgent2 = new(BranchAndFinancialInstitutionIdentification5)
	return d.IntermediaryAgent2
}

func (d *DirectDebitTransactionInformation21) AddIntermediaryAgent2Account() *CashAccount24 {
	d.IntermediaryAgent2Account = new(CashAccount24)
	return d.IntermediaryAgent2Account
}

func (d *DirectDebitTransactionInformation21) AddIntermediaryAgent3() *BranchAndFinancialInstitutionIdentification5 {
	d.IntermediaryAgent3 = new(BranchAndFinancialInstitutionIdentification5)
	return d.IntermediaryAgent3
}

func (d *DirectDebitTransactionInformation21) AddIntermediaryAgent3Account() *CashAccount24 {
	d.IntermediaryAgent3Account = new(CashAccount24)
	return d.IntermediaryAgent3Account
}

func (d *DirectDebitTransactionInformation21) AddDebtor() *PartyIdentification43 {
	d.Debtor = new(PartyIdentification43)
	return d.Debtor
}

func (d *DirectDebitTransactionInformation21) AddDebtorAccount() *CashAccount24 {
	d.DebtorAccount = new(CashAccount24)
	return d.DebtorAccount
}

func (d *DirectDebitTransactionInformation21) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	d.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return d.DebtorAgent
}

func (d *DirectDebitTransactionInformation21) AddDebtorAgentAccount() *CashAccount24 {
	d.DebtorAgentAccount = new(CashAccount24)
	return d.DebtorAgentAccount
}

func (d *DirectDebitTransactionInformation21) AddUltimateDebtor() *PartyIdentification43 {
	d.UltimateDebtor = new(PartyIdentification43)
	return d.UltimateDebtor
}

func (d *DirectDebitTransactionInformation21) AddPurpose() *Purpose2Choice {
	d.Purpose = new(Purpose2Choice)
	return d.Purpose
}

func (d *DirectDebitTransactionInformation21) AddRegulatoryReporting() *RegulatoryReporting3 {
	newValue := new(RegulatoryReporting3)
	d.RegulatoryReporting = append(d.RegulatoryReporting, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation21) AddRelatedRemittanceInformation() *RemittanceLocation4 {
	newValue := new(RemittanceLocation4)
	d.RelatedRemittanceInformation = append(d.RelatedRemittanceInformation, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation21) AddRemittanceInformation() *RemittanceInformation11 {
	d.RemittanceInformation = new(RemittanceInformation11)
	return d.RemittanceInformation
}

func (d *DirectDebitTransactionInformation21) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	d.SupplementaryData = append(d.SupplementaryData, newValue)
	return newValue
}

// Provides further details specific to the individual direct debit transaction(s) included in the message.
type DirectDebitTransactionInformation22 struct {

	// Set of elements used to reference a payment instruction.
	PaymentIdentification *PaymentIdentification1 `xml:"PmtId"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation24 `xml:"PmtTpInf,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	// Usage: This amount has to be transported unchanged through the transaction chain.
	InstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Provides information specific to the direct debit mandate.
	DirectDebitTransaction *DirectDebitTransaction9 `xml:"DrctDbtTx,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount24 `xml:"DbtrAgtAcct,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Further information, related to the processing of the payment instruction, that may need to be acted upon by the creditor agent, depending on agreement between creditor and the creditor agent.
	InstructionForCreditorAgent *Max140Text `xml:"InstrForCdtrAgt,omitempty"`

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`

	// Information needed due to regulatory and statutory requirements.
	RegulatoryReporting []*RegulatoryReporting3 `xml:"RgltryRptg,omitempty"`

	// Provides details on the tax.
	Tax *TaxInformation3 `xml:"Tax,omitempty"`

	// Provides information related to the handling of the remittance information by any of the agents in the transaction processing chain.
	RelatedRemittanceInformation []*RemittanceLocation4 `xml:"RltdRmtInf,omitempty"`

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RemittanceInformation *RemittanceInformation11 `xml:"RmtInf,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (d *DirectDebitTransactionInformation22) AddPaymentIdentification() *PaymentIdentification1 {
	d.PaymentIdentification = new(PaymentIdentification1)
	return d.PaymentIdentification
}

func (d *DirectDebitTransactionInformation22) AddPaymentTypeInformation() *PaymentTypeInformation24 {
	d.PaymentTypeInformation = new(PaymentTypeInformation24)
	return d.PaymentTypeInformation
}

func (d *DirectDebitTransactionInformation22) SetInstructedAmount(value, currency string) {
	d.InstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (d *DirectDebitTransactionInformation22) SetChargeBearer(value string) {
	d.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (d *DirectDebitTransactionInformation22) AddDirectDebitTransaction() *DirectDebitTransaction9 {
	d.DirectDebitTransaction = new(DirectDebitTransaction9)
	return d.DirectDebitTransaction
}

func (d *DirectDebitTransactionInformation22) AddUltimateCreditor() *PartyIdentification43 {
	d.UltimateCreditor = new(PartyIdentification43)
	return d.UltimateCreditor
}

func (d *DirectDebitTransactionInformation22) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	d.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return d.DebtorAgent
}

func (d *DirectDebitTransactionInformation22) AddDebtorAgentAccount() *CashAccount24 {
	d.DebtorAgentAccount = new(CashAccount24)
	return d.DebtorAgentAccount
}

func (d *DirectDebitTransactionInformation22) AddDebtor() *PartyIdentification43 {
	d.Debtor = new(PartyIdentification43)
	return d.Debtor
}

func (d *DirectDebitTransactionInformation22) AddDebtorAccount() *CashAccount24 {
	d.DebtorAccount = new(CashAccount24)
	return d.DebtorAccount
}

func (d *DirectDebitTransactionInformation22) AddUltimateDebtor() *PartyIdentification43 {
	d.UltimateDebtor = new(PartyIdentification43)
	return d.UltimateDebtor
}

func (d *DirectDebitTransactionInformation22) SetInstructionForCreditorAgent(value string) {
	d.InstructionForCreditorAgent = (*Max140Text)(&value)
}

func (d *DirectDebitTransactionInformation22) AddPurpose() *Purpose2Choice {
	d.Purpose = new(Purpose2Choice)
	return d.Purpose
}

func (d *DirectDebitTransactionInformation22) AddRegulatoryReporting() *RegulatoryReporting3 {
	newValue := new(RegulatoryReporting3)
	d.RegulatoryReporting = append(d.RegulatoryReporting, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation22) AddTax() *TaxInformation3 {
	d.Tax = new(TaxInformation3)
	return d.Tax
}

func (d *DirectDebitTransactionInformation22) AddRelatedRemittanceInformation() *RemittanceLocation4 {
	newValue := new(RemittanceLocation4)
	d.RelatedRemittanceInformation = append(d.RelatedRemittanceInformation, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation22) AddRemittanceInformation() *RemittanceInformation11 {
	d.RemittanceInformation = new(RemittanceInformation11)
	return d.RemittanceInformation
}

func (d *DirectDebitTransactionInformation22) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	d.SupplementaryData = append(d.SupplementaryData, newValue)
	return newValue
}

// Set of elements used to provide information specific to the individual direct debit transaction(s) included in the message.
type DirectDebitTransactionInformation9 struct {

	// Set of elements used to reference a payment instruction.
	PaymentIdentification *PaymentIdentification1 `xml:"PmtId"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation20 `xml:"PmtTpInf,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	InstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Set of elements providing information specific to the direct debit mandate.
	DirectDebitTransaction *DirectDebitTransaction6 `xml:"DrctDbtTx,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification32 `xml:"UltmtCdtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification4 `xml:"DbtrAgt"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount16 `xml:"DbtrAgtAcct,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification32 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount16 `xml:"DbtrAcct"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification32 `xml:"UltmtDbtr,omitempty"`

	// Further information, related to the processing of the payment instruction, that may need to be acted upon by the creditor agent, depending on agreement between creditor and the creditor agent.
	InstructionForCreditorAgent *Max140Text `xml:"InstrForCdtrAgt,omitempty"`

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`

	// Information needed due to regulatory and statutory requirements.
	RegulatoryReporting []*RegulatoryReporting3 `xml:"RgltryRptg,omitempty"`

	// Set of elements used to provide details on the tax.
	Tax *TaxInformation3 `xml:"Tax,omitempty"`

	// Set of elements used to provide information related to the handling of the remittance information by any of the agents in the transaction processing chain.
	RelatedRemittanceInformation []*RemittanceLocation2 `xml:"RltdRmtInf,omitempty"`

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RemittanceInformation *RemittanceInformation5 `xml:"RmtInf,omitempty"`
}

func (d *DirectDebitTransactionInformation9) AddPaymentIdentification() *PaymentIdentification1 {
	d.PaymentIdentification = new(PaymentIdentification1)
	return d.PaymentIdentification
}

func (d *DirectDebitTransactionInformation9) AddPaymentTypeInformation() *PaymentTypeInformation20 {
	d.PaymentTypeInformation = new(PaymentTypeInformation20)
	return d.PaymentTypeInformation
}

func (d *DirectDebitTransactionInformation9) SetInstructedAmount(value, currency string) {
	d.InstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (d *DirectDebitTransactionInformation9) SetChargeBearer(value string) {
	d.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (d *DirectDebitTransactionInformation9) AddDirectDebitTransaction() *DirectDebitTransaction6 {
	d.DirectDebitTransaction = new(DirectDebitTransaction6)
	return d.DirectDebitTransaction
}

func (d *DirectDebitTransactionInformation9) AddUltimateCreditor() *PartyIdentification32 {
	d.UltimateCreditor = new(PartyIdentification32)
	return d.UltimateCreditor
}

func (d *DirectDebitTransactionInformation9) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification4 {
	d.DebtorAgent = new(BranchAndFinancialInstitutionIdentification4)
	return d.DebtorAgent
}

func (d *DirectDebitTransactionInformation9) AddDebtorAgentAccount() *CashAccount16 {
	d.DebtorAgentAccount = new(CashAccount16)
	return d.DebtorAgentAccount
}

func (d *DirectDebitTransactionInformation9) AddDebtor() *PartyIdentification32 {
	d.Debtor = new(PartyIdentification32)
	return d.Debtor
}

func (d *DirectDebitTransactionInformation9) AddDebtorAccount() *CashAccount16 {
	d.DebtorAccount = new(CashAccount16)
	return d.DebtorAccount
}

func (d *DirectDebitTransactionInformation9) AddUltimateDebtor() *PartyIdentification32 {
	d.UltimateDebtor = new(PartyIdentification32)
	return d.UltimateDebtor
}

func (d *DirectDebitTransactionInformation9) SetInstructionForCreditorAgent(value string) {
	d.InstructionForCreditorAgent = (*Max140Text)(&value)
}

func (d *DirectDebitTransactionInformation9) AddPurpose() *Purpose2Choice {
	d.Purpose = new(Purpose2Choice)
	return d.Purpose
}

func (d *DirectDebitTransactionInformation9) AddRegulatoryReporting() *RegulatoryReporting3 {
	newValue := new(RegulatoryReporting3)
	d.RegulatoryReporting = append(d.RegulatoryReporting, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation9) AddTax() *TaxInformation3 {
	d.Tax = new(TaxInformation3)
	return d.Tax
}

func (d *DirectDebitTransactionInformation9) AddRelatedRemittanceInformation() *RemittanceLocation2 {
	newValue := new(RemittanceLocation2)
	d.RelatedRemittanceInformation = append(d.RelatedRemittanceInformation, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation9) AddRemittanceInformation() *RemittanceInformation5 {
	d.RemittanceInformation = new(RemittanceInformation5)
	return d.RemittanceInformation
}
