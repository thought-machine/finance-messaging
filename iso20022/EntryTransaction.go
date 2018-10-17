package iso20022

// Posting to an account that results in an increase or decrease to a balance.
type EntryTransaction1 struct {

	// Set of elements providing the identification of the underlying transaction.
	References *TransactionReferences1 `xml:"Refs,omitempty"`

	// Set of elements providing details information on the original amount.
	//
	// Usage: This component (on transaction level) should be used in case booking is for a single transaction and the original amount is different from the entry amount.  It can also be used in case individual original amounts are provided in case of a batch or aggregate booking.
	AmountDetails *AmountAndCurrencyExchange2 `xml:"AmtDtls,omitempty"`

	// Set of elements used to indicate when the booked funds will become available, ie can be accessed and start generating interest.
	//
	// Usage : this type of info is eg used in US, and is linked to particular instruments, such as cheques.
	// Example : When a cheque is deposited, it will be booked on the deposit day, but the funds will only be accessible as of the indicated availability day (according to national banking regulations).
	Availability []*CashBalanceAvailability1 `xml:"Avlbty,omitempty"`

	// Set of elements to fully identify the type of underlying transaction resulting in an entry.
	BankTransactionCode *BankTransactionCodeStructure1 `xml:"BkTxCd,omitempty"`

	// Provides information on the charges included in the entry amount.
	//
	// Usage : This component (on transaction level) can be used in case the booking is for a single transaction, and charges are included in the entry amount.  It can also be used in case individual charge amounts are applied to individual transactions in case of a batch or aggregate amount booking.
	Charges []*ChargesInformation3 `xml:"Chrgs,omitempty"`

	// Set of elements providing details on the interest amount included in the entry amount.
	//
	// Usage : This component (on transaction level) can be used in case the booking is for a single transaction, and interest amount is included in the entry amount.  It can also be used in case individual interest amounts are applied to individual transactions in case of a batch or aggregate amount booking.
	Interest []*TransactionInterest1 `xml:"Intrst,omitempty"`

	// Set of elements identifying the parties related to the underlying transaction.
	RelatedParties *TransactionParty1 `xml:"RltdPties,omitempty"`

	// Set of elements identifying the agents related to the underlying transaction.
	RelatedAgents *TransactionAgents1 `xml:"RltdAgts,omitempty"`

	// Underlying reason for the payment transaction, eg, a charity payment, or a commercial agreement between the creditor and the debtor.
	//
	// Usage: purpose is used by the end-customers, ie originating party, initiating party, debtor, creditor, final party, to provide information concerning the nature of the payment transaction. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose1Choice `xml:"Purp,omitempty"`

	// Information related to the handling of the remittance information by any of the agents in the transaction processing chain.
	RelatedRemittanceInformation []*RemittanceLocation1 `xml:"RltdRmtInf,omitempty"`

	// Information that enables the matching, ie, reconciliation, of a payment with the items that the payment is intended to settle, eg, commercial invoices in an account receivable system.
	RemittanceInformation *RemittanceInformation1 `xml:"RmtInf,omitempty"`

	// Set of elements identifying the dates related to the underlying transactions.
	RelatedDates *TransactionDates1 `xml:"RltdDts,omitempty"`

	// Set of elements identifying the price information related to the underlying transaction.
	RelatedPrice *TransactionPrice1Choice `xml:"RltdPric,omitempty"`

	// Identifies related quantities (eg of securities) in the underlying transaction.
	RelatedQuantities []*TransactionQuantities1Choice `xml:"RltdQties,omitempty"`

	// Identification of a security, as assigned under a formal or proprietary identification scheme.
	FinancialInstrumentIdentification *SecurityIdentification4Choice `xml:"FinInstrmId,omitempty"`

	// Amount of money due to the government or tax authority, according to various pre-defined parameters such as thresholds or income.
	Tax *TaxInformation2 `xml:"Tax,omitempty"`

	// Set of elements specifying the return information.
	ReturnInformation *ReturnReasonInformation5 `xml:"RtrInf,omitempty"`

	// Set of elements identifying the underlying corporate action.
	CorporateAction *CorporateAction1 `xml:"CorpActn,omitempty"`

	// Safekeeping or investment account. A safekeeping account is an account on which a securities entry is made. An investment account is an account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
	SafekeepingAccount *CashAccount7 `xml:"SfkpgAcct,omitempty"`

	// Further details on the transaction details.
	AdditionalTransactionInformation *Max500Text `xml:"AddtlTxInf,omitempty"`
}

func (e *EntryTransaction1) AddReferences() *TransactionReferences1 {
	e.References = new(TransactionReferences1)
	return e.References
}

func (e *EntryTransaction1) AddAmountDetails() *AmountAndCurrencyExchange2 {
	e.AmountDetails = new(AmountAndCurrencyExchange2)
	return e.AmountDetails
}

func (e *EntryTransaction1) AddAvailability() *CashBalanceAvailability1 {
	newValue := new(CashBalanceAvailability1)
	e.Availability = append(e.Availability, newValue)
	return newValue
}

func (e *EntryTransaction1) AddBankTransactionCode() *BankTransactionCodeStructure1 {
	e.BankTransactionCode = new(BankTransactionCodeStructure1)
	return e.BankTransactionCode
}

func (e *EntryTransaction1) AddCharges() *ChargesInformation3 {
	newValue := new(ChargesInformation3)
	e.Charges = append(e.Charges, newValue)
	return newValue
}

func (e *EntryTransaction1) AddInterest() *TransactionInterest1 {
	newValue := new(TransactionInterest1)
	e.Interest = append(e.Interest, newValue)
	return newValue
}

func (e *EntryTransaction1) AddRelatedParties() *TransactionParty1 {
	e.RelatedParties = new(TransactionParty1)
	return e.RelatedParties
}

func (e *EntryTransaction1) AddRelatedAgents() *TransactionAgents1 {
	e.RelatedAgents = new(TransactionAgents1)
	return e.RelatedAgents
}

func (e *EntryTransaction1) AddPurpose() *Purpose1Choice {
	e.Purpose = new(Purpose1Choice)
	return e.Purpose
}

func (e *EntryTransaction1) AddRelatedRemittanceInformation() *RemittanceLocation1 {
	newValue := new(RemittanceLocation1)
	e.RelatedRemittanceInformation = append(e.RelatedRemittanceInformation, newValue)
	return newValue
}

func (e *EntryTransaction1) AddRemittanceInformation() *RemittanceInformation1 {
	e.RemittanceInformation = new(RemittanceInformation1)
	return e.RemittanceInformation
}

func (e *EntryTransaction1) AddRelatedDates() *TransactionDates1 {
	e.RelatedDates = new(TransactionDates1)
	return e.RelatedDates
}

func (e *EntryTransaction1) AddRelatedPrice() *TransactionPrice1Choice {
	e.RelatedPrice = new(TransactionPrice1Choice)
	return e.RelatedPrice
}

func (e *EntryTransaction1) AddRelatedQuantities() *TransactionQuantities1Choice {
	newValue := new(TransactionQuantities1Choice)
	e.RelatedQuantities = append(e.RelatedQuantities, newValue)
	return newValue
}

func (e *EntryTransaction1) AddFinancialInstrumentIdentification() *SecurityIdentification4Choice {
	e.FinancialInstrumentIdentification = new(SecurityIdentification4Choice)
	return e.FinancialInstrumentIdentification
}

func (e *EntryTransaction1) AddTax() *TaxInformation2 {
	e.Tax = new(TaxInformation2)
	return e.Tax
}

func (e *EntryTransaction1) AddReturnInformation() *ReturnReasonInformation5 {
	e.ReturnInformation = new(ReturnReasonInformation5)
	return e.ReturnInformation
}

func (e *EntryTransaction1) AddCorporateAction() *CorporateAction1 {
	e.CorporateAction = new(CorporateAction1)
	return e.CorporateAction
}

func (e *EntryTransaction1) AddSafekeepingAccount() *CashAccount7 {
	e.SafekeepingAccount = new(CashAccount7)
	return e.SafekeepingAccount
}

func (e *EntryTransaction1) SetAdditionalTransactionInformation(value string) {
	e.AdditionalTransactionInformation = (*Max500Text)(&value)
}

// Set of elements used to identify the underlying transaction.
type EntryTransaction2 struct {

	// Set of elements used to provide the identification of the underlying transaction.
	References *TransactionReferences2 `xml:"Refs,omitempty"`

	// Set of elements providing detailed information on the original amount.
	//
	// Usage: This component (on transaction level) should be used in case booking is for a single transaction and the original amount is different from the entry amount. It can also be used in case individual original amounts are provided in case of a batch or aggregate booking.
	AmountDetails *AmountAndCurrencyExchange3 `xml:"AmtDtls,omitempty"`

	// Set of elements used to indicate when the booked amount of money will become available, that is can be accessed and starts generating interest.
	//
	// Usage: This type of information is used in the US and is linked to particular instruments such as cheques.
	// Example: When a cheque is deposited, it will be booked on the deposit day, but the amount of money will only be accessible as of the indicated availability day (according to national banking regulations).
	Availability []*CashBalanceAvailability2 `xml:"Avlbty,omitempty"`

	// Set of elements used to fully identify the type of underlying transaction resulting in an entry.
	BankTransactionCode *BankTransactionCodeStructure4 `xml:"BkTxCd,omitempty"`

	// Provides information on the charges included in the entry amount.
	//
	// Usage: This component (on transaction level) can be used in case the booking is for a single transaction, and charges are included in the entry amount. It can also be used in case individual charge amounts are applied to individual transactions in case of a batch or aggregate amount booking.
	Charges []*ChargesInformation6 `xml:"Chrgs,omitempty"`

	// Set of elements used to provide details of the interest amount included in the entry amount.
	//
	// Usage: This component (on transaction level) can be used if the booking is for a single transaction, and interest amount is included in the entry amount.  It can also be used if individual interest amounts are applied to individual transactions in the case of a batch or aggregate amount booking.
	Interest []*TransactionInterest2 `xml:"Intrst,omitempty"`

	// Set of elements used to identify the parties related to the underlying transaction.
	RelatedParties *TransactionParty2 `xml:"RltdPties,omitempty"`

	// Set of elements used to identify the agents related to the underlying transaction.
	RelatedAgents *TransactionAgents2 `xml:"RltdAgts,omitempty"`

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`

	// Set of elements used to provide information related to the handling of the remittance information by any of the agents in the transaction processing chain.
	RelatedRemittanceInformation []*RemittanceLocation2 `xml:"RltdRmtInf,omitempty"`

	// Structured information that enables the matching, ie, reconciliation, of a payment with the items that the payment is intended to settle, such as commercial invoices in an account receivable system.
	RemittanceInformation *RemittanceInformation5 `xml:"RmtInf,omitempty"`

	// Set of elements used to identify the dates related to the underlying transactions.
	RelatedDates *TransactionDates2 `xml:"RltdDts,omitempty"`

	// Set of elements used to identify the price information related to the underlying transaction.
	RelatedPrice *TransactionPrice2Choice `xml:"RltdPric,omitempty"`

	// Set of elements used to identify the related quantities, such as securities, in the underlying transaction.
	RelatedQuantities []*TransactionQuantities1Choice `xml:"RltdQties,omitempty"`

	// Identification of a security, as assigned under a formal or proprietary identification scheme.
	FinancialInstrumentIdentification *SecurityIdentification4Choice `xml:"FinInstrmId,omitempty"`

	// Set of elements used to provide details on the tax.
	Tax *TaxInformation3 `xml:"Tax,omitempty"`

	// Set of elements used to provide the return information.
	ReturnInformation *ReturnReasonInformation10 `xml:"RtrInf,omitempty"`

	// Set of elements used to identify the underlying corporate action.
	CorporateAction *CorporateAction1 `xml:"CorpActn,omitempty"`

	// Safekeeping or investment account. A safekeeping account is an account on which a securities entry is made. An investment account is an account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
	SafekeepingAccount *CashAccount16 `xml:"SfkpgAcct,omitempty"`

	// Further details of the transaction.
	AdditionalTransactionInformation *Max500Text `xml:"AddtlTxInf,omitempty"`
}

func (e *EntryTransaction2) AddReferences() *TransactionReferences2 {
	e.References = new(TransactionReferences2)
	return e.References
}

func (e *EntryTransaction2) AddAmountDetails() *AmountAndCurrencyExchange3 {
	e.AmountDetails = new(AmountAndCurrencyExchange3)
	return e.AmountDetails
}

func (e *EntryTransaction2) AddAvailability() *CashBalanceAvailability2 {
	newValue := new(CashBalanceAvailability2)
	e.Availability = append(e.Availability, newValue)
	return newValue
}

func (e *EntryTransaction2) AddBankTransactionCode() *BankTransactionCodeStructure4 {
	e.BankTransactionCode = new(BankTransactionCodeStructure4)
	return e.BankTransactionCode
}

func (e *EntryTransaction2) AddCharges() *ChargesInformation6 {
	newValue := new(ChargesInformation6)
	e.Charges = append(e.Charges, newValue)
	return newValue
}

func (e *EntryTransaction2) AddInterest() *TransactionInterest2 {
	newValue := new(TransactionInterest2)
	e.Interest = append(e.Interest, newValue)
	return newValue
}

func (e *EntryTransaction2) AddRelatedParties() *TransactionParty2 {
	e.RelatedParties = new(TransactionParty2)
	return e.RelatedParties
}

func (e *EntryTransaction2) AddRelatedAgents() *TransactionAgents2 {
	e.RelatedAgents = new(TransactionAgents2)
	return e.RelatedAgents
}

func (e *EntryTransaction2) AddPurpose() *Purpose2Choice {
	e.Purpose = new(Purpose2Choice)
	return e.Purpose
}

func (e *EntryTransaction2) AddRelatedRemittanceInformation() *RemittanceLocation2 {
	newValue := new(RemittanceLocation2)
	e.RelatedRemittanceInformation = append(e.RelatedRemittanceInformation, newValue)
	return newValue
}

func (e *EntryTransaction2) AddRemittanceInformation() *RemittanceInformation5 {
	e.RemittanceInformation = new(RemittanceInformation5)
	return e.RemittanceInformation
}

func (e *EntryTransaction2) AddRelatedDates() *TransactionDates2 {
	e.RelatedDates = new(TransactionDates2)
	return e.RelatedDates
}

func (e *EntryTransaction2) AddRelatedPrice() *TransactionPrice2Choice {
	e.RelatedPrice = new(TransactionPrice2Choice)
	return e.RelatedPrice
}

func (e *EntryTransaction2) AddRelatedQuantities() *TransactionQuantities1Choice {
	newValue := new(TransactionQuantities1Choice)
	e.RelatedQuantities = append(e.RelatedQuantities, newValue)
	return newValue
}

func (e *EntryTransaction2) AddFinancialInstrumentIdentification() *SecurityIdentification4Choice {
	e.FinancialInstrumentIdentification = new(SecurityIdentification4Choice)
	return e.FinancialInstrumentIdentification
}

func (e *EntryTransaction2) AddTax() *TaxInformation3 {
	e.Tax = new(TaxInformation3)
	return e.Tax
}

func (e *EntryTransaction2) AddReturnInformation() *ReturnReasonInformation10 {
	e.ReturnInformation = new(ReturnReasonInformation10)
	return e.ReturnInformation
}

func (e *EntryTransaction2) AddCorporateAction() *CorporateAction1 {
	e.CorporateAction = new(CorporateAction1)
	return e.CorporateAction
}

func (e *EntryTransaction2) AddSafekeepingAccount() *CashAccount16 {
	e.SafekeepingAccount = new(CashAccount16)
	return e.SafekeepingAccount
}

func (e *EntryTransaction2) SetAdditionalTransactionInformation(value string) {
	e.AdditionalTransactionInformation = (*Max500Text)(&value)
}

// Identifies the underlying transaction.
type EntryTransaction3 struct {

	// Provides the identification of the underlying transaction.
	References *TransactionReferences3 `xml:"Refs,omitempty"`

	// Amount of money in the cash transaction.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether the transaction is a credit or a debit transaction.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Provides detailed information on the original amount.
	//
	// Usage: This component (on transaction level) should be used in case booking is for a single transaction and the original amount is different from the entry amount. It can also be used in case individual original amounts are provided in case of a batch or aggregate booking.
	AmountDetails *AmountAndCurrencyExchange3 `xml:"AmtDtls,omitempty"`

	// Indicates when the booked amount of money will become available, that is can be accessed and starts generating interest.
	//
	// Usage: This type of information is used in the US and is linked to particular instruments such as cheques.
	// Example: When a cheque is deposited, it will be booked on the deposit day, but the amount of money will only be accessible as of the indicated availability day (according to national banking regulations).
	Availability []*CashBalanceAvailability2 `xml:"Avlbty,omitempty"`

	// Set of elements used to fully identify the type of underlying transaction resulting in an entry.
	BankTransactionCode *BankTransactionCodeStructure4 `xml:"BkTxCd,omitempty"`

	// Provides information on the charges included in the entry amount.
	//
	// Usage: This component (on transaction level) can be used in case the booking is for a single transaction, and charges are included in the entry amount. It can also be used in case individual charge amounts are applied to individual transactions in case of a batch or aggregate amount booking.
	Charges []*Charges3 `xml:"Chrgs,omitempty"`

	// Provides details of the interest amount included in the entry amount.
	//
	// Usage: This component (on transaction level) can be used if the booking is for a single transaction, and interest amount is included in the entry amount.  It can also be used if individual interest amounts are applied to individual transactions in the case of a batch or aggregate amount booking.
	Interest []*TransactionInterest3 `xml:"Intrst,omitempty"`

	// Set of elements used to identify the parties related to the underlying transaction.
	RelatedParties *TransactionParties3 `xml:"RltdPties,omitempty"`

	// Set of elements used to identify the agents related to the underlying transaction.
	RelatedAgents *TransactionAgents3 `xml:"RltdAgts,omitempty"`

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`

	// Provides information related to the handling of the remittance information by any of the agents in the transaction processing chain.
	RelatedRemittanceInformation []*RemittanceLocation2 `xml:"RltdRmtInf,omitempty"`

	// Structured information that enables the matching, that is reconciliation, of a payment with the items that the payment is intended to settle, such as commercial invoices in an account receivable system.
	RemittanceInformation *RemittanceInformation7 `xml:"RmtInf,omitempty"`

	// Set of elements used to identify the dates related to the underlying transactions.
	RelatedDates *TransactionDates2 `xml:"RltdDts,omitempty"`

	// Set of elements used to identify the price information related to the underlying transaction.
	RelatedPrice *TransactionPrice3Choice `xml:"RltdPric,omitempty"`

	// Set of elements used to identify the related quantities, such as securities, in the underlying transaction.
	RelatedQuantities []*TransactionQuantities2Choice `xml:"RltdQties,omitempty"`

	// Identification of a security, as assigned under a formal or proprietary identification scheme.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId,omitempty"`

	// Provides details on the tax.
	Tax *TaxInformation3 `xml:"Tax,omitempty"`

	// Provides the return information.
	ReturnInformation *PaymentReturnReason2 `xml:"RtrInf,omitempty"`

	// Set of elements used to identify the underlying corporate action.
	CorporateAction *CorporateAction9 `xml:"CorpActn,omitempty"`

	// Safekeeping or investment account. A safekeeping account is an account on which a securities entry is made. An investment account is an account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct,omitempty"`

	// Provides the details of a cash deposit for an amount of money in cash notes and/or coins.
	CashDeposit []*CashDeposit1 `xml:"CshDpst,omitempty"`

	// Provides the data related to the card (number, scheme), terminal (number, identification) and transactional data used to uniquely identify a card transaction.
	CardTransaction *CardTransaction1 `xml:"CardTx,omitempty"`

	// Further details of the transaction.
	AdditionalTransactionInformation *Max500Text `xml:"AddtlTxInf,omitempty"`
}

func (e *EntryTransaction3) AddReferences() *TransactionReferences3 {
	e.References = new(TransactionReferences3)
	return e.References
}

func (e *EntryTransaction3) SetAmount(value, currency string) {
	e.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (e *EntryTransaction3) SetCreditDebitIndicator(value string) {
	e.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (e *EntryTransaction3) AddAmountDetails() *AmountAndCurrencyExchange3 {
	e.AmountDetails = new(AmountAndCurrencyExchange3)
	return e.AmountDetails
}

func (e *EntryTransaction3) AddAvailability() *CashBalanceAvailability2 {
	newValue := new(CashBalanceAvailability2)
	e.Availability = append(e.Availability, newValue)
	return newValue
}

func (e *EntryTransaction3) AddBankTransactionCode() *BankTransactionCodeStructure4 {
	e.BankTransactionCode = new(BankTransactionCodeStructure4)
	return e.BankTransactionCode
}

func (e *EntryTransaction3) AddCharges() *Charges3 {
	newValue := new(Charges3)
	e.Charges = append(e.Charges, newValue)
	return newValue
}

func (e *EntryTransaction3) AddInterest() *TransactionInterest3 {
	newValue := new(TransactionInterest3)
	e.Interest = append(e.Interest, newValue)
	return newValue
}

func (e *EntryTransaction3) AddRelatedParties() *TransactionParties3 {
	e.RelatedParties = new(TransactionParties3)
	return e.RelatedParties
}

func (e *EntryTransaction3) AddRelatedAgents() *TransactionAgents3 {
	e.RelatedAgents = new(TransactionAgents3)
	return e.RelatedAgents
}

func (e *EntryTransaction3) AddPurpose() *Purpose2Choice {
	e.Purpose = new(Purpose2Choice)
	return e.Purpose
}

func (e *EntryTransaction3) AddRelatedRemittanceInformation() *RemittanceLocation2 {
	newValue := new(RemittanceLocation2)
	e.RelatedRemittanceInformation = append(e.RelatedRemittanceInformation, newValue)
	return newValue
}

func (e *EntryTransaction3) AddRemittanceInformation() *RemittanceInformation7 {
	e.RemittanceInformation = new(RemittanceInformation7)
	return e.RemittanceInformation
}

func (e *EntryTransaction3) AddRelatedDates() *TransactionDates2 {
	e.RelatedDates = new(TransactionDates2)
	return e.RelatedDates
}

func (e *EntryTransaction3) AddRelatedPrice() *TransactionPrice3Choice {
	e.RelatedPrice = new(TransactionPrice3Choice)
	return e.RelatedPrice
}

func (e *EntryTransaction3) AddRelatedQuantities() *TransactionQuantities2Choice {
	newValue := new(TransactionQuantities2Choice)
	e.RelatedQuantities = append(e.RelatedQuantities, newValue)
	return newValue
}

func (e *EntryTransaction3) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	e.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return e.FinancialInstrumentIdentification
}

func (e *EntryTransaction3) AddTax() *TaxInformation3 {
	e.Tax = new(TaxInformation3)
	return e.Tax
}

func (e *EntryTransaction3) AddReturnInformation() *PaymentReturnReason2 {
	e.ReturnInformation = new(PaymentReturnReason2)
	return e.ReturnInformation
}

func (e *EntryTransaction3) AddCorporateAction() *CorporateAction9 {
	e.CorporateAction = new(CorporateAction9)
	return e.CorporateAction
}

func (e *EntryTransaction3) AddSafekeepingAccount() *SecuritiesAccount13 {
	e.SafekeepingAccount = new(SecuritiesAccount13)
	return e.SafekeepingAccount
}

func (e *EntryTransaction3) AddCashDeposit() *CashDeposit1 {
	newValue := new(CashDeposit1)
	e.CashDeposit = append(e.CashDeposit, newValue)
	return newValue
}

func (e *EntryTransaction3) AddCardTransaction() *CardTransaction1 {
	e.CardTransaction = new(CardTransaction1)
	return e.CardTransaction
}

func (e *EntryTransaction3) SetAdditionalTransactionInformation(value string) {
	e.AdditionalTransactionInformation = (*Max500Text)(&value)
}

// Identifies the underlying transaction.
type EntryTransaction4 struct {

	// Provides the identification of the underlying transaction.
	References *TransactionReferences3 `xml:"Refs,omitempty"`

	// Amount of money in the cash transaction.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether the transaction is a credit or a debit transaction.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Provides detailed information on the original amount.
	//
	// Usage: This component (on transaction level) should be used in case booking is for a single transaction and the original amount is different from the entry amount. It can also be used in case individual original amounts are provided in case of a batch or aggregate booking.
	AmountDetails *AmountAndCurrencyExchange3 `xml:"AmtDtls,omitempty"`

	// Indicates when the booked amount of money will become available, that is can be accessed and starts generating interest.
	//
	// Usage: This type of information is used in the US and is linked to particular instruments such as cheques.
	// Example: When a cheque is deposited, it will be booked on the deposit day, but the amount of money will only be accessible as of the indicated availability day (according to national banking regulations).
	Availability []*CashBalanceAvailability2 `xml:"Avlbty,omitempty"`

	// Set of elements used to fully identify the type of underlying transaction resulting in an entry.
	BankTransactionCode *BankTransactionCodeStructure4 `xml:"BkTxCd,omitempty"`

	// Provides information on the charges, pre-advised or included in the entry amount.
	//
	// Usage: This component (on transaction level) can be used in case the booking is for a single transaction, and charges are included in the entry amount. It can also be used in case individual charge amounts are applied to individual transactions in case of a batch or aggregate amount booking.
	Charges *Charges4 `xml:"Chrgs,omitempty"`

	// Provides details of the interest amount included in the entry amount.
	//
	// Usage: This component (on transaction level) can be used if the booking is for a single transaction, and interest amount is included in the entry amount.  It can also be used if individual interest amounts are applied to individual transactions in the case of a batch or aggregate amount booking.
	Interest *TransactionInterest3 `xml:"Intrst,omitempty"`

	// Set of elements used to identify the parties related to the underlying transaction.
	RelatedParties *TransactionParties3 `xml:"RltdPties,omitempty"`

	// Set of elements used to identify the agents related to the underlying transaction.
	RelatedAgents *TransactionAgents3 `xml:"RltdAgts,omitempty"`

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`

	// Provides information related to the handling of the remittance information by any of the agents in the transaction processing chain.
	RelatedRemittanceInformation []*RemittanceLocation2 `xml:"RltdRmtInf,omitempty"`

	// Structured information that enables the matching, that is reconciliation, of a payment with the items that the payment is intended to settle, such as commercial invoices in an account receivable system.
	RemittanceInformation *RemittanceInformation7 `xml:"RmtInf,omitempty"`

	// Set of elements used to identify the dates related to the underlying transactions.
	RelatedDates *TransactionDates2 `xml:"RltdDts,omitempty"`

	// Set of elements used to identify the price information related to the underlying transaction.
	RelatedPrice *TransactionPrice3Choice `xml:"RltdPric,omitempty"`

	// Set of elements used to identify the related quantities, such as securities, in the underlying transaction.
	RelatedQuantities []*TransactionQuantities2Choice `xml:"RltdQties,omitempty"`

	// Identification of a security, as assigned under a formal or proprietary identification scheme.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId,omitempty"`

	// Provides details on the tax.
	Tax *TaxInformation3 `xml:"Tax,omitempty"`

	// Provides the return information.
	ReturnInformation *PaymentReturnReason2 `xml:"RtrInf,omitempty"`

	// Set of elements used to identify the underlying corporate action.
	CorporateAction *CorporateAction9 `xml:"CorpActn,omitempty"`

	// Safekeeping or investment account. A safekeeping account is an account on which a securities entry is made. An investment account is an account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct,omitempty"`

	// Provides the details of a cash deposit for an amount of money in cash notes and/or coins.
	CashDeposit []*CashDeposit1 `xml:"CshDpst,omitempty"`

	// Provides the data related to the card (number, scheme), terminal (number, identification) and transactional data used to uniquely identify a card transaction.
	CardTransaction *CardTransaction1 `xml:"CardTx,omitempty"`

	// Further details of the transaction.
	AdditionalTransactionInformation *Max500Text `xml:"AddtlTxInf,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (e *EntryTransaction4) AddReferences() *TransactionReferences3 {
	e.References = new(TransactionReferences3)
	return e.References
}

func (e *EntryTransaction4) SetAmount(value, currency string) {
	e.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (e *EntryTransaction4) SetCreditDebitIndicator(value string) {
	e.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (e *EntryTransaction4) AddAmountDetails() *AmountAndCurrencyExchange3 {
	e.AmountDetails = new(AmountAndCurrencyExchange3)
	return e.AmountDetails
}

func (e *EntryTransaction4) AddAvailability() *CashBalanceAvailability2 {
	newValue := new(CashBalanceAvailability2)
	e.Availability = append(e.Availability, newValue)
	return newValue
}

func (e *EntryTransaction4) AddBankTransactionCode() *BankTransactionCodeStructure4 {
	e.BankTransactionCode = new(BankTransactionCodeStructure4)
	return e.BankTransactionCode
}

func (e *EntryTransaction4) AddCharges() *Charges4 {
	e.Charges = new(Charges4)
	return e.Charges
}

func (e *EntryTransaction4) AddInterest() *TransactionInterest3 {
	e.Interest = new(TransactionInterest3)
	return e.Interest
}

func (e *EntryTransaction4) AddRelatedParties() *TransactionParties3 {
	e.RelatedParties = new(TransactionParties3)
	return e.RelatedParties
}

func (e *EntryTransaction4) AddRelatedAgents() *TransactionAgents3 {
	e.RelatedAgents = new(TransactionAgents3)
	return e.RelatedAgents
}

func (e *EntryTransaction4) AddPurpose() *Purpose2Choice {
	e.Purpose = new(Purpose2Choice)
	return e.Purpose
}

func (e *EntryTransaction4) AddRelatedRemittanceInformation() *RemittanceLocation2 {
	newValue := new(RemittanceLocation2)
	e.RelatedRemittanceInformation = append(e.RelatedRemittanceInformation, newValue)
	return newValue
}

func (e *EntryTransaction4) AddRemittanceInformation() *RemittanceInformation7 {
	e.RemittanceInformation = new(RemittanceInformation7)
	return e.RemittanceInformation
}

func (e *EntryTransaction4) AddRelatedDates() *TransactionDates2 {
	e.RelatedDates = new(TransactionDates2)
	return e.RelatedDates
}

func (e *EntryTransaction4) AddRelatedPrice() *TransactionPrice3Choice {
	e.RelatedPrice = new(TransactionPrice3Choice)
	return e.RelatedPrice
}

func (e *EntryTransaction4) AddRelatedQuantities() *TransactionQuantities2Choice {
	newValue := new(TransactionQuantities2Choice)
	e.RelatedQuantities = append(e.RelatedQuantities, newValue)
	return newValue
}

func (e *EntryTransaction4) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	e.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return e.FinancialInstrumentIdentification
}

func (e *EntryTransaction4) AddTax() *TaxInformation3 {
	e.Tax = new(TaxInformation3)
	return e.Tax
}

func (e *EntryTransaction4) AddReturnInformation() *PaymentReturnReason2 {
	e.ReturnInformation = new(PaymentReturnReason2)
	return e.ReturnInformation
}

func (e *EntryTransaction4) AddCorporateAction() *CorporateAction9 {
	e.CorporateAction = new(CorporateAction9)
	return e.CorporateAction
}

func (e *EntryTransaction4) AddSafekeepingAccount() *SecuritiesAccount13 {
	e.SafekeepingAccount = new(SecuritiesAccount13)
	return e.SafekeepingAccount
}

func (e *EntryTransaction4) AddCashDeposit() *CashDeposit1 {
	newValue := new(CashDeposit1)
	e.CashDeposit = append(e.CashDeposit, newValue)
	return newValue
}

func (e *EntryTransaction4) AddCardTransaction() *CardTransaction1 {
	e.CardTransaction = new(CardTransaction1)
	return e.CardTransaction
}

func (e *EntryTransaction4) SetAdditionalTransactionInformation(value string) {
	e.AdditionalTransactionInformation = (*Max500Text)(&value)
}

func (e *EntryTransaction4) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	e.SupplementaryData = append(e.SupplementaryData, newValue)
	return newValue
}

// Identifies the underlying transaction.
type EntryTransaction7 struct {

	// Provides the identification of the underlying transaction.
	References *TransactionReferences3 `xml:"Refs,omitempty"`

	// Amount of money in the cash transaction.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether the transaction is a credit or a debit transaction.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Provides detailed information on the original amount.
	//
	// Usage: This component (on transaction level) should be used in case booking is for a single transaction and the original amount is different from the entry amount. It can also be used in case individual original amounts are provided in case of a batch or aggregate booking.
	AmountDetails *AmountAndCurrencyExchange3 `xml:"AmtDtls,omitempty"`

	// Indicates when the booked amount of money will become available, that is can be accessed and starts generating interest.
	//
	// Usage: This type of information is used in the US and is linked to particular instruments such as cheques.
	// Example: When a cheque is deposited, it will be booked on the deposit day, but the amount of money will only be accessible as of the indicated availability day (according to national banking regulations).
	Availability []*CashBalanceAvailability2 `xml:"Avlbty,omitempty"`

	// Set of elements used to fully identify the type of underlying transaction resulting in an entry.
	BankTransactionCode *BankTransactionCodeStructure4 `xml:"BkTxCd,omitempty"`

	// Provides information on the charges, pre-advised or included in the entry amount.
	//
	// Usage: This component (on transaction level) can be used in case the booking is for a single transaction, and charges are included in the entry amount. It can also be used in case individual charge amounts are applied to individual transactions in case of a batch or aggregate amount booking.
	Charges *Charges4 `xml:"Chrgs,omitempty"`

	// Provides details of the interest amount included in the entry amount.
	//
	// Usage: This component (on transaction level) can be used if the booking is for a single transaction, and interest amount is included in the entry amount.  It can also be used if individual interest amounts are applied to individual transactions in the case of a batch or aggregate amount booking.
	Interest *TransactionInterest3 `xml:"Intrst,omitempty"`

	// Set of elements used to identify the parties related to the underlying transaction.
	RelatedParties *TransactionParties3 `xml:"RltdPties,omitempty"`

	// Set of elements used to identify the agents related to the underlying transaction.
	RelatedAgents *TransactionAgents3 `xml:"RltdAgts,omitempty"`

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`

	// Provides information related to the handling of the remittance information by any of the agents in the transaction processing chain.
	RelatedRemittanceInformation []*RemittanceLocation4 `xml:"RltdRmtInf,omitempty"`

	// Structured information that enables the matching, that is reconciliation, of a payment with the items that the payment is intended to settle, such as commercial invoices in an account receivable system.
	RemittanceInformation *RemittanceInformation10 `xml:"RmtInf,omitempty"`

	// Set of elements used to identify the dates related to the underlying transactions.
	RelatedDates *TransactionDates2 `xml:"RltdDts,omitempty"`

	// Set of elements used to identify the price information related to the underlying transaction.
	RelatedPrice *TransactionPrice3Choice `xml:"RltdPric,omitempty"`

	// Set of elements used to identify the related quantities, such as securities, in the underlying transaction.
	RelatedQuantities []*TransactionQuantities2Choice `xml:"RltdQties,omitempty"`

	// Identification of a security, as assigned under a formal or proprietary identification scheme.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId,omitempty"`

	// Provides details on the tax.
	Tax *TaxInformation3 `xml:"Tax,omitempty"`

	// Provides the return information.
	ReturnInformation *PaymentReturnReason2 `xml:"RtrInf,omitempty"`

	// Set of elements used to identify the underlying corporate action.
	CorporateAction *CorporateAction9 `xml:"CorpActn,omitempty"`

	// Safekeeping or investment account. A safekeeping account is an account on which a securities entry is made. An investment account is an account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct,omitempty"`

	// Provides the details of a cash deposit for an amount of money in cash notes and/or coins.
	CashDeposit []*CashDeposit1 `xml:"CshDpst,omitempty"`

	// Provides the data related to the card (number, scheme), terminal (number, identification) and transactional data used to uniquely identify a card transaction.
	CardTransaction *CardTransaction2 `xml:"CardTx,omitempty"`

	// Further details of the transaction.
	AdditionalTransactionInformation *Max500Text `xml:"AddtlTxInf,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (e *EntryTransaction7) AddReferences() *TransactionReferences3 {
	e.References = new(TransactionReferences3)
	return e.References
}

func (e *EntryTransaction7) SetAmount(value, currency string) {
	e.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (e *EntryTransaction7) SetCreditDebitIndicator(value string) {
	e.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (e *EntryTransaction7) AddAmountDetails() *AmountAndCurrencyExchange3 {
	e.AmountDetails = new(AmountAndCurrencyExchange3)
	return e.AmountDetails
}

func (e *EntryTransaction7) AddAvailability() *CashBalanceAvailability2 {
	newValue := new(CashBalanceAvailability2)
	e.Availability = append(e.Availability, newValue)
	return newValue
}

func (e *EntryTransaction7) AddBankTransactionCode() *BankTransactionCodeStructure4 {
	e.BankTransactionCode = new(BankTransactionCodeStructure4)
	return e.BankTransactionCode
}

func (e *EntryTransaction7) AddCharges() *Charges4 {
	e.Charges = new(Charges4)
	return e.Charges
}

func (e *EntryTransaction7) AddInterest() *TransactionInterest3 {
	e.Interest = new(TransactionInterest3)
	return e.Interest
}

func (e *EntryTransaction7) AddRelatedParties() *TransactionParties3 {
	e.RelatedParties = new(TransactionParties3)
	return e.RelatedParties
}

func (e *EntryTransaction7) AddRelatedAgents() *TransactionAgents3 {
	e.RelatedAgents = new(TransactionAgents3)
	return e.RelatedAgents
}

func (e *EntryTransaction7) AddPurpose() *Purpose2Choice {
	e.Purpose = new(Purpose2Choice)
	return e.Purpose
}

func (e *EntryTransaction7) AddRelatedRemittanceInformation() *RemittanceLocation4 {
	newValue := new(RemittanceLocation4)
	e.RelatedRemittanceInformation = append(e.RelatedRemittanceInformation, newValue)
	return newValue
}

func (e *EntryTransaction7) AddRemittanceInformation() *RemittanceInformation10 {
	e.RemittanceInformation = new(RemittanceInformation10)
	return e.RemittanceInformation
}

func (e *EntryTransaction7) AddRelatedDates() *TransactionDates2 {
	e.RelatedDates = new(TransactionDates2)
	return e.RelatedDates
}

func (e *EntryTransaction7) AddRelatedPrice() *TransactionPrice3Choice {
	e.RelatedPrice = new(TransactionPrice3Choice)
	return e.RelatedPrice
}

func (e *EntryTransaction7) AddRelatedQuantities() *TransactionQuantities2Choice {
	newValue := new(TransactionQuantities2Choice)
	e.RelatedQuantities = append(e.RelatedQuantities, newValue)
	return newValue
}

func (e *EntryTransaction7) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	e.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return e.FinancialInstrumentIdentification
}

func (e *EntryTransaction7) AddTax() *TaxInformation3 {
	e.Tax = new(TaxInformation3)
	return e.Tax
}

func (e *EntryTransaction7) AddReturnInformation() *PaymentReturnReason2 {
	e.ReturnInformation = new(PaymentReturnReason2)
	return e.ReturnInformation
}

func (e *EntryTransaction7) AddCorporateAction() *CorporateAction9 {
	e.CorporateAction = new(CorporateAction9)
	return e.CorporateAction
}

func (e *EntryTransaction7) AddSafekeepingAccount() *SecuritiesAccount13 {
	e.SafekeepingAccount = new(SecuritiesAccount13)
	return e.SafekeepingAccount
}

func (e *EntryTransaction7) AddCashDeposit() *CashDeposit1 {
	newValue := new(CashDeposit1)
	e.CashDeposit = append(e.CashDeposit, newValue)
	return newValue
}

func (e *EntryTransaction7) AddCardTransaction() *CardTransaction2 {
	e.CardTransaction = new(CardTransaction2)
	return e.CardTransaction
}

func (e *EntryTransaction7) SetAdditionalTransactionInformation(value string) {
	e.AdditionalTransactionInformation = (*Max500Text)(&value)
}

func (e *EntryTransaction7) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	e.SupplementaryData = append(e.SupplementaryData, newValue)
	return newValue
}

// Identifies the underlying transaction.
type EntryTransaction8 struct {

	// Provides the identification of the underlying transaction.
	References *TransactionReferences3 `xml:"Refs,omitempty"`

	// Amount of money in the cash transaction.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether the transaction is a credit or a debit transaction.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Provides detailed information on the original amount.
	//
	// Usage: This component (on transaction level) should be used in case booking is for a single transaction and the original amount is different from the entry amount. It can also be used in case individual original amounts are provided in case of a batch or aggregate booking.
	AmountDetails *AmountAndCurrencyExchange3 `xml:"AmtDtls,omitempty"`

	// Indicates when the booked amount of money will become available, that is can be accessed and starts generating interest.
	//
	// Usage: This type of information is used in the US and is linked to particular instruments such as cheques.
	// Example: When a cheque is deposited, it will be booked on the deposit day, but the amount of money will only be accessible as of the indicated availability day (according to national banking regulations).
	Availability []*CashAvailability1 `xml:"Avlbty,omitempty"`

	// Set of elements used to fully identify the type of underlying transaction resulting in an entry.
	BankTransactionCode *BankTransactionCodeStructure4 `xml:"BkTxCd,omitempty"`

	// Provides information on the charges, pre-advised or included in the entry amount.
	//
	// Usage: This component (on transaction level) can be used in case the booking is for a single transaction, and charges are included in the entry amount. It can also be used in case individual charge amounts are applied to individual transactions in case of a batch or aggregate amount booking.
	Charges *Charges4 `xml:"Chrgs,omitempty"`

	// Provides details of the interest amount included in the entry amount.
	//
	// Usage: This component (on transaction level) can be used if the booking is for a single transaction, and interest amount is included in the entry amount.  It can also be used if individual interest amounts are applied to individual transactions in the case of a batch or aggregate amount booking.
	Interest *TransactionInterest3 `xml:"Intrst,omitempty"`

	// Set of elements used to identify the parties related to the underlying transaction.
	RelatedParties *TransactionParties3 `xml:"RltdPties,omitempty"`

	// Set of elements used to identify the agents related to the underlying transaction.
	RelatedAgents *TransactionAgents3 `xml:"RltdAgts,omitempty"`

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`

	// Provides information related to the handling of the remittance information by any of the agents in the transaction processing chain.
	RelatedRemittanceInformation []*RemittanceLocation4 `xml:"RltdRmtInf,omitempty"`

	// Structured information that enables the matching, that is reconciliation, of a payment with the items that the payment is intended to settle, such as commercial invoices in an account receivable system.
	RemittanceInformation *RemittanceInformation11 `xml:"RmtInf,omitempty"`

	// Set of elements used to identify the dates related to the underlying transactions.
	RelatedDates *TransactionDates2 `xml:"RltdDts,omitempty"`

	// Set of elements used to identify the price information related to the underlying transaction.
	RelatedPrice *TransactionPrice3Choice `xml:"RltdPric,omitempty"`

	// Set of elements used to identify the related quantities, such as securities, in the underlying transaction.
	RelatedQuantities []*TransactionQuantities2Choice `xml:"RltdQties,omitempty"`

	// Identification of a security, as assigned under a formal or proprietary identification scheme.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId,omitempty"`

	// Provides details on the tax.
	Tax *TaxInformation3 `xml:"Tax,omitempty"`

	// Provides the return information.
	ReturnInformation *PaymentReturnReason2 `xml:"RtrInf,omitempty"`

	// Set of elements used to identify the underlying corporate action.
	CorporateAction *CorporateAction9 `xml:"CorpActn,omitempty"`

	// Safekeeping or investment account. A safekeeping account is an account on which a securities entry is made. An investment account is an account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
	SafekeepingAccount *SecuritiesAccount19 `xml:"SfkpgAcct,omitempty"`

	// Provides the details of a cash deposit for an amount of money in cash notes and/or coins.
	CashDeposit []*CashDeposit1 `xml:"CshDpst,omitempty"`

	// Provides the data related to the card (number, scheme), terminal (number, identification) and transactional data used to uniquely identify a card transaction.
	CardTransaction *CardTransaction2 `xml:"CardTx,omitempty"`

	// Further details of the transaction.
	AdditionalTransactionInformation *Max500Text `xml:"AddtlTxInf,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (e *EntryTransaction8) AddReferences() *TransactionReferences3 {
	e.References = new(TransactionReferences3)
	return e.References
}

func (e *EntryTransaction8) SetAmount(value, currency string) {
	e.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (e *EntryTransaction8) SetCreditDebitIndicator(value string) {
	e.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (e *EntryTransaction8) AddAmountDetails() *AmountAndCurrencyExchange3 {
	e.AmountDetails = new(AmountAndCurrencyExchange3)
	return e.AmountDetails
}

func (e *EntryTransaction8) AddAvailability() *CashAvailability1 {
	newValue := new(CashAvailability1)
	e.Availability = append(e.Availability, newValue)
	return newValue
}

func (e *EntryTransaction8) AddBankTransactionCode() *BankTransactionCodeStructure4 {
	e.BankTransactionCode = new(BankTransactionCodeStructure4)
	return e.BankTransactionCode
}

func (e *EntryTransaction8) AddCharges() *Charges4 {
	e.Charges = new(Charges4)
	return e.Charges
}

func (e *EntryTransaction8) AddInterest() *TransactionInterest3 {
	e.Interest = new(TransactionInterest3)
	return e.Interest
}

func (e *EntryTransaction8) AddRelatedParties() *TransactionParties3 {
	e.RelatedParties = new(TransactionParties3)
	return e.RelatedParties
}

func (e *EntryTransaction8) AddRelatedAgents() *TransactionAgents3 {
	e.RelatedAgents = new(TransactionAgents3)
	return e.RelatedAgents
}

func (e *EntryTransaction8) AddPurpose() *Purpose2Choice {
	e.Purpose = new(Purpose2Choice)
	return e.Purpose
}

func (e *EntryTransaction8) AddRelatedRemittanceInformation() *RemittanceLocation4 {
	newValue := new(RemittanceLocation4)
	e.RelatedRemittanceInformation = append(e.RelatedRemittanceInformation, newValue)
	return newValue
}

func (e *EntryTransaction8) AddRemittanceInformation() *RemittanceInformation11 {
	e.RemittanceInformation = new(RemittanceInformation11)
	return e.RemittanceInformation
}

func (e *EntryTransaction8) AddRelatedDates() *TransactionDates2 {
	e.RelatedDates = new(TransactionDates2)
	return e.RelatedDates
}

func (e *EntryTransaction8) AddRelatedPrice() *TransactionPrice3Choice {
	e.RelatedPrice = new(TransactionPrice3Choice)
	return e.RelatedPrice
}

func (e *EntryTransaction8) AddRelatedQuantities() *TransactionQuantities2Choice {
	newValue := new(TransactionQuantities2Choice)
	e.RelatedQuantities = append(e.RelatedQuantities, newValue)
	return newValue
}

func (e *EntryTransaction8) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	e.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return e.FinancialInstrumentIdentification
}

func (e *EntryTransaction8) AddTax() *TaxInformation3 {
	e.Tax = new(TaxInformation3)
	return e.Tax
}

func (e *EntryTransaction8) AddReturnInformation() *PaymentReturnReason2 {
	e.ReturnInformation = new(PaymentReturnReason2)
	return e.ReturnInformation
}

func (e *EntryTransaction8) AddCorporateAction() *CorporateAction9 {
	e.CorporateAction = new(CorporateAction9)
	return e.CorporateAction
}

func (e *EntryTransaction8) AddSafekeepingAccount() *SecuritiesAccount19 {
	e.SafekeepingAccount = new(SecuritiesAccount19)
	return e.SafekeepingAccount
}

func (e *EntryTransaction8) AddCashDeposit() *CashDeposit1 {
	newValue := new(CashDeposit1)
	e.CashDeposit = append(e.CashDeposit, newValue)
	return newValue
}

func (e *EntryTransaction8) AddCardTransaction() *CardTransaction2 {
	e.CardTransaction = new(CardTransaction2)
	return e.CardTransaction
}

func (e *EntryTransaction8) SetAdditionalTransactionInformation(value string) {
	e.AdditionalTransactionInformation = (*Max500Text)(&value)
}

func (e *EntryTransaction8) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	e.SupplementaryData = append(e.SupplementaryData, newValue)
	return newValue
}
