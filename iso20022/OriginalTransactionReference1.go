package iso20022

// Set of key elements of the original transaction being referred to.
type OriginalTransactionReference1 struct {

	// Amount of money moved between the instructing agent and the instructed agent.
	InterbankSettlementAmount *CurrencyAndAmount `xml:"IntrBkSttlmAmt,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	Amount *AmountType2Choice `xml:"Amt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Date at which the initiating party requests that the clearing agent to process the payment. If payment by cheque, the date when the cheque must be generated by the bank.
	//
	// Usage: This is the date on which the debtor's account(s) is (are) to be debited.
	RequestedExecutionDate *ISODate `xml:"ReqdExctnDt,omitempty"`

	// Date at which the creditor requests the amount of money to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt,omitempty"`

	// Credit party that signs the direct debit mandate.
	CreditorSchemeIdentification *PartyIdentification8 `xml:"CdtrSchmeId,omitempty"`

	// Specifies the details on how the settlement of the original transaction(s) between the instructing agent and the instructed agent was completed.
	SettlementInformation *SettlementInformation3 `xml:"SttlmInf,omitempty"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation6 `xml:"PmtTpInf,omitempty"`

	// Specifies the transfer method that will be used by the instructing agent to transfer the funds to the creditor.
	PaymentMethod *PaymentMethod4Code `xml:"PmtMtd,omitempty"`

	// Set of elements used to provide further details related to a direct debit mandate signed between the creditor and the debtor.
	//
	// Usage: Mandate related information is to be used only when the direct debit relates to a mandate signed between the debtor and the creditor.
	MandateRelatedInformation *MandateRelatedInformation1 `xml:"MndtRltdInf,omitempty"`

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RemittanceInformation *RemittanceInformation1 `xml:"RmtInf,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification8 `xml:"UltmtDbtr,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification8 `xml:"Dbtr,omitempty"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount7 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification3 `xml:"DbtrAgt,omitempty"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount7 `xml:"DbtrAgtAcct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification3 `xml:"CdtrAgt,omitempty"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CreditorAgentAccount *CashAccount7 `xml:"CdtrAgtAcct,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification8 `xml:"Cdtr,omitempty"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount7 `xml:"CdtrAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification8 `xml:"UltmtCdtr,omitempty"`
}

func (o *OriginalTransactionReference1) SetInterbankSettlementAmount(value, currency string) {
	o.InterbankSettlementAmount = NewCurrencyAndAmount(value, currency)
}

func (o *OriginalTransactionReference1) AddAmount() *AmountType2Choice {
	o.Amount = new(AmountType2Choice)
	return o.Amount
}

func (o *OriginalTransactionReference1) SetInterbankSettlementDate(value string) {
	o.InterbankSettlementDate = (*ISODate)(&value)
}

func (o *OriginalTransactionReference1) SetRequestedExecutionDate(value string) {
	o.RequestedExecutionDate = (*ISODate)(&value)
}

func (o *OriginalTransactionReference1) SetRequestedCollectionDate(value string) {
	o.RequestedCollectionDate = (*ISODate)(&value)
}

func (o *OriginalTransactionReference1) AddCreditorSchemeIdentification() *PartyIdentification8 {
	o.CreditorSchemeIdentification = new(PartyIdentification8)
	return o.CreditorSchemeIdentification
}

func (o *OriginalTransactionReference1) AddSettlementInformation() *SettlementInformation3 {
	o.SettlementInformation = new(SettlementInformation3)
	return o.SettlementInformation
}

func (o *OriginalTransactionReference1) AddPaymentTypeInformation() *PaymentTypeInformation6 {
	o.PaymentTypeInformation = new(PaymentTypeInformation6)
	return o.PaymentTypeInformation
}

func (o *OriginalTransactionReference1) SetPaymentMethod(value string) {
	o.PaymentMethod = (*PaymentMethod4Code)(&value)
}

func (o *OriginalTransactionReference1) AddMandateRelatedInformation() *MandateRelatedInformation1 {
	o.MandateRelatedInformation = new(MandateRelatedInformation1)
	return o.MandateRelatedInformation
}

func (o *OriginalTransactionReference1) AddRemittanceInformation() *RemittanceInformation1 {
	o.RemittanceInformation = new(RemittanceInformation1)
	return o.RemittanceInformation
}

func (o *OriginalTransactionReference1) AddUltimateDebtor() *PartyIdentification8 {
	o.UltimateDebtor = new(PartyIdentification8)
	return o.UltimateDebtor
}

func (o *OriginalTransactionReference1) AddDebtor() *PartyIdentification8 {
	o.Debtor = new(PartyIdentification8)
	return o.Debtor
}

func (o *OriginalTransactionReference1) AddDebtorAccount() *CashAccount7 {
	o.DebtorAccount = new(CashAccount7)
	return o.DebtorAccount
}

func (o *OriginalTransactionReference1) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification3 {
	o.DebtorAgent = new(BranchAndFinancialInstitutionIdentification3)
	return o.DebtorAgent
}

func (o *OriginalTransactionReference1) AddDebtorAgentAccount() *CashAccount7 {
	o.DebtorAgentAccount = new(CashAccount7)
	return o.DebtorAgentAccount
}

func (o *OriginalTransactionReference1) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification3 {
	o.CreditorAgent = new(BranchAndFinancialInstitutionIdentification3)
	return o.CreditorAgent
}

func (o *OriginalTransactionReference1) AddCreditorAgentAccount() *CashAccount7 {
	o.CreditorAgentAccount = new(CashAccount7)
	return o.CreditorAgentAccount
}

func (o *OriginalTransactionReference1) AddCreditor() *PartyIdentification8 {
	o.Creditor = new(PartyIdentification8)
	return o.Creditor
}

func (o *OriginalTransactionReference1) AddCreditorAccount() *CashAccount7 {
	o.CreditorAccount = new(CashAccount7)
	return o.CreditorAccount
}

func (o *OriginalTransactionReference1) AddUltimateCreditor() *PartyIdentification8 {
	o.UltimateCreditor = new(PartyIdentification8)
	return o.UltimateCreditor
}