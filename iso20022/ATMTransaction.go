package iso20022

// Withdrawal transaction for which an authorisation is requested.
type ATMTransaction1 struct {

	// True if cash has to be dispensed by the ATM for the transaction.
	CashDispensed *TrueFalseIndicator `xml:"CshDspnsd,omitempty"`

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the reconciliation period assigned by the ATM.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Unprotected account information.
	AccountData *CardAccount3 `xml:"AcctData,omitempty"`

	// Encryption of account information.
	ProtectedAccountData *ContentInformationType10 `xml:"PrtctdAcctData,omitempty"`

	// Amount to be authorised by the issuer.
	TotalRequestedAmount *AmountAndCurrency1 `xml:"TtlReqdAmt,omitempty"`

	// Amounts of the withdrawal transaction.
	DetailedRequestedAmount *DetailedAmount12 `xml:"DtldReqdAmt,omitempty"`

	// Currency conversion accepted by the customer, either to convert the amount to dispense in the base currency of the ATM, or to convert the total requested amount in the currency of the customer (so called dynamic currency conversion).
	CurrencyConversion *CurrencyConversion4 `xml:"CcyConvs,omitempty"`

	// Media mix algorithm, the identification of the algorithm is an agreement between the ATM and the ATM manager.
	SelectedMixType *Max35Text `xml:"SelctdMixTp,omitempty"`

	// Media mix selected.
	SelectedMix []*ATMMediaMix1 `xml:"SelctdMix,omitempty"`

	// True if a receipt has be requested by the customer.
	RequestedReceipt *TrueFalseIndicator `xml:"ReqdRct,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (a *ATMTransaction1) SetCashDispensed(value string) {
	a.CashDispensed = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction1) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction1) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction1) AddAccountData() *CardAccount3 {
	a.AccountData = new(CardAccount3)
	return a.AccountData
}

func (a *ATMTransaction1) AddProtectedAccountData() *ContentInformationType10 {
	a.ProtectedAccountData = new(ContentInformationType10)
	return a.ProtectedAccountData
}

func (a *ATMTransaction1) AddTotalRequestedAmount() *AmountAndCurrency1 {
	a.TotalRequestedAmount = new(AmountAndCurrency1)
	return a.TotalRequestedAmount
}

func (a *ATMTransaction1) AddDetailedRequestedAmount() *DetailedAmount12 {
	a.DetailedRequestedAmount = new(DetailedAmount12)
	return a.DetailedRequestedAmount
}

func (a *ATMTransaction1) AddCurrencyConversion() *CurrencyConversion4 {
	a.CurrencyConversion = new(CurrencyConversion4)
	return a.CurrencyConversion
}

func (a *ATMTransaction1) SetSelectedMixType(value string) {
	a.SelectedMixType = (*Max35Text)(&value)
}

func (a *ATMTransaction1) AddSelectedMix() *ATMMediaMix1 {
	newValue := new(ATMMediaMix1)
	a.SelectedMix = append(a.SelectedMix, newValue)
	return newValue
}

func (a *ATMTransaction1) SetRequestedReceipt(value string) {
	a.RequestedReceipt = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction1) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}

// Response to the PIN management transaction. request.
type ATMTransaction10 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the reconciliation period assigned by the ATM.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// True if a completion advice has to be sent after the end of the transaction.
	CompletionRequired *TrueFalseIndicator `xml:"CmpltnReqrd,omitempty"`

	// Result of the PIN service.
	TransactionResponse *ResponseType3 `xml:"TxRspn"`

	// Sequence of actions to be performed by the ATM to complete the transaction.
	Action []*Action5 `xml:"Actn,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`

	// Maintenance command to perform on the ATM.
	Command []*ATMCommand1 `xml:"Cmd,omitempty"`
}

func (a *ATMTransaction10) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction10) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction10) SetCompletionRequired(value string) {
	a.CompletionRequired = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction10) AddTransactionResponse() *ResponseType3 {
	a.TransactionResponse = new(ResponseType3)
	return a.TransactionResponse
}

func (a *ATMTransaction10) AddAction() *Action5 {
	newValue := new(Action5)
	a.Action = append(a.Action, newValue)
	return newValue
}

func (a *ATMTransaction10) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}

func (a *ATMTransaction10) AddCommand() *ATMCommand1 {
	newValue := new(ATMCommand1)
	a.Command = append(a.Command, newValue)
	return newValue
}

// Information about the reconciliation request.
type ATMTransaction11 struct {

	// Type of logical or physical operation on the ATM for which the counters are computed.
	TypeOfOperation *ATMOperation1Code `xml:"TpOfOpr,omitempty"`

	// Identification of the reconciliation transaction.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the reconciliation period.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId"`

	// Current totals of the ATM.
	ATMTotals []*ATMTotals1 `xml:"ATMTtls,omitempty"`

	// Information on the cassette of the ATM.
	Cassette []*ATMCassette1 `xml:"Csstt,omitempty"`

	// Transaction counters that are set to zero after a reconciliation with counter reinitialisation command.
	TransactionTotals []*ATMTotals3 `xml:"TxTtls,omitempty"`

	// Total number of retained cards.
	RetainedCard *Number `xml:"RtndCard,omitempty"`

	// Additional information about reconciliation.
	AdditionalTransactionInformation *Max140Text `xml:"AddtlTxInf,omitempty"`
}

func (a *ATMTransaction11) SetTypeOfOperation(value string) {
	a.TypeOfOperation = (*ATMOperation1Code)(&value)
}

func (a *ATMTransaction11) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction11) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction11) AddATMTotals() *ATMTotals1 {
	newValue := new(ATMTotals1)
	a.ATMTotals = append(a.ATMTotals, newValue)
	return newValue
}

func (a *ATMTransaction11) AddCassette() *ATMCassette1 {
	newValue := new(ATMCassette1)
	a.Cassette = append(a.Cassette, newValue)
	return newValue
}

func (a *ATMTransaction11) AddTransactionTotals() *ATMTotals3 {
	newValue := new(ATMTotals3)
	a.TransactionTotals = append(a.TransactionTotals, newValue)
	return newValue
}

func (a *ATMTransaction11) SetRetainedCard(value string) {
	a.RetainedCard = (*Number)(&value)
}

func (a *ATMTransaction11) SetAdditionalTransactionInformation(value string) {
	a.AdditionalTransactionInformation = (*Max140Text)(&value)
}

// Information about the reconciliation response.
type ATMTransaction12 struct {

	// Type of logical or physical operation on the ATM for which the counters are computed.
	TypeOfOperation *ATMOperation1Code `xml:"TpOfOpr,omitempty"`

	// Identification of the reconciliation transaction.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the reconciliation period assigned by the ATM.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId"`

	// Result of the reconciliation.
	TransactionResponse *ResponseType3 `xml:"TxRspn"`

	// Maintenance command to perform on the ATM.
	Command []*ATMCommand1 `xml:"Cmd,omitempty"`
}

func (a *ATMTransaction12) SetTypeOfOperation(value string) {
	a.TypeOfOperation = (*ATMOperation1Code)(&value)
}

func (a *ATMTransaction12) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction12) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction12) AddTransactionResponse() *ResponseType3 {
	a.TransactionResponse = new(ResponseType3)
	return a.TransactionResponse
}

func (a *ATMTransaction12) AddCommand() *ATMCommand1 {
	newValue := new(ATMCommand1)
	a.Command = append(a.Command, newValue)
	return newValue
}

// Withdrawal transaction for which an authorisation is requested.
type ATMTransaction13 struct {

	// True if cash has to be dispensed by the ATM for the transaction.
	CashDispensed *TrueFalseIndicator `xml:"CshDspnsd,omitempty"`

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the reconciliation period assigned by the ATM.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Unprotected account information.
	AccountData *CardAccount7 `xml:"AcctData,omitempty"`

	// Encryption of account information.
	ProtectedAccountData *ContentInformationType10 `xml:"PrtctdAcctData,omitempty"`

	// Amount to be authorised by the issuer.
	TotalRequestedAmount *AmountAndCurrency1 `xml:"TtlReqdAmt,omitempty"`

	// Amounts of the withdrawal transaction.
	DetailedRequestedAmount *DetailedAmount12 `xml:"DtldReqdAmt,omitempty"`

	// Currency conversion accepted by the customer, either to convert the amount to dispense in the base currency of the ATM, or to convert the total requested amount in the currency of the customer (so called dynamic currency conversion).
	CurrencyConversionResult *CurrencyConversion9 `xml:"CcyConvsRslt,omitempty"`

	// Media mix algorithm, the identification of the algorithm is an agreement between the ATM and the ATM manager.
	SelectedMixType *Max35Text `xml:"SelctdMixTp,omitempty"`

	// Media mix selected.
	SelectedMix []*ATMMediaMix1 `xml:"SelctdMix,omitempty"`

	// True if a receipt has be requested by the customer.
	RequestedReceipt *TrueFalseIndicator `xml:"ReqdRct,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (a *ATMTransaction13) SetCashDispensed(value string) {
	a.CashDispensed = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction13) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction13) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction13) AddAccountData() *CardAccount7 {
	a.AccountData = new(CardAccount7)
	return a.AccountData
}

func (a *ATMTransaction13) AddProtectedAccountData() *ContentInformationType10 {
	a.ProtectedAccountData = new(ContentInformationType10)
	return a.ProtectedAccountData
}

func (a *ATMTransaction13) AddTotalRequestedAmount() *AmountAndCurrency1 {
	a.TotalRequestedAmount = new(AmountAndCurrency1)
	return a.TotalRequestedAmount
}

func (a *ATMTransaction13) AddDetailedRequestedAmount() *DetailedAmount12 {
	a.DetailedRequestedAmount = new(DetailedAmount12)
	return a.DetailedRequestedAmount
}

func (a *ATMTransaction13) AddCurrencyConversionResult() *CurrencyConversion9 {
	a.CurrencyConversionResult = new(CurrencyConversion9)
	return a.CurrencyConversionResult
}

func (a *ATMTransaction13) SetSelectedMixType(value string) {
	a.SelectedMixType = (*Max35Text)(&value)
}

func (a *ATMTransaction13) AddSelectedMix() *ATMMediaMix1 {
	newValue := new(ATMMediaMix1)
	a.SelectedMix = append(a.SelectedMix, newValue)
	return newValue
}

func (a *ATMTransaction13) SetRequestedReceipt(value string) {
	a.RequestedReceipt = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction13) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}

// Response to the withdrawal transaction request.
type ATMTransaction14 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the reconciliation period assigned by the ATM.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// True if a withdrawal completion advice has to be sent after the end of the transaction.
	CompletionRequired *TrueFalseIndicator `xml:"CmpltnReqrd,omitempty"`

	// Unprotected account information.
	AccountData *CardAccount8 `xml:"AcctData,omitempty"`

	// Encryption of account information.
	ProtectedAccountData *ContentInformationType10 `xml:"PrtctdAcctData,omitempty"`

	// Total authorised amount.
	TotalAuthorisedAmount *AmountAndCurrency1 `xml:"TtlAuthrsdAmt"`

	// Total requested amount.
	TotalRequestedAmount *ImpliedCurrencyAndAmount `xml:"TtlReqdAmt,omitempty"`

	// Detail of the requested amounts for the withdrawal transaction.
	DetailedRequestedAmount *DetailedAmount12 `xml:"DtldReqdAmt,omitempty"`

	// Currency conversion accepted by the customer, either to convert the amount to dispense in the base currency of the ATM, or to convert the total requested amount in the currency of the customer (so called dynamic currency conversion).
	CurrencyConversionEligibility *CurrencyConversion9 `xml:"CcyConvsElgblty,omitempty"`

	// Additional charge (for instance tax or fee).
	AdditionalCharge []*DetailedAmount13 `xml:"AddtlChrg,omitempty"`

	// Limit of amounts for the customer.
	Limits *ATMTransactionAmounts6 `xml:"Lmts,omitempty"`

	// Media mix algorithm requested by the ATM Host, the identification of the algorithm is an agreement between the ATM and the ATM manager.
	MixType *Max35Text `xml:"MixTp,omitempty"`

	// Media mix selected requested by the ATM Host.
	Mix []*ATMMediaMix1 `xml:"Mix,omitempty"`

	// Outcome of the withdrawal authorisation.
	AuthorisationResult *AuthorisationResult13 `xml:"AuthstnRslt"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`

	// Maintenance command to perform on the ATM.
	Command []*ATMCommand7 `xml:"Cmd,omitempty"`
}

func (a *ATMTransaction14) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction14) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction14) SetCompletionRequired(value string) {
	a.CompletionRequired = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction14) AddAccountData() *CardAccount8 {
	a.AccountData = new(CardAccount8)
	return a.AccountData
}

func (a *ATMTransaction14) AddProtectedAccountData() *ContentInformationType10 {
	a.ProtectedAccountData = new(ContentInformationType10)
	return a.ProtectedAccountData
}

func (a *ATMTransaction14) AddTotalAuthorisedAmount() *AmountAndCurrency1 {
	a.TotalAuthorisedAmount = new(AmountAndCurrency1)
	return a.TotalAuthorisedAmount
}

func (a *ATMTransaction14) SetTotalRequestedAmount(value, currency string) {
	a.TotalRequestedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransaction14) AddDetailedRequestedAmount() *DetailedAmount12 {
	a.DetailedRequestedAmount = new(DetailedAmount12)
	return a.DetailedRequestedAmount
}

func (a *ATMTransaction14) AddCurrencyConversionEligibility() *CurrencyConversion9 {
	a.CurrencyConversionEligibility = new(CurrencyConversion9)
	return a.CurrencyConversionEligibility
}

func (a *ATMTransaction14) AddAdditionalCharge() *DetailedAmount13 {
	newValue := new(DetailedAmount13)
	a.AdditionalCharge = append(a.AdditionalCharge, newValue)
	return newValue
}

func (a *ATMTransaction14) AddLimits() *ATMTransactionAmounts6 {
	a.Limits = new(ATMTransactionAmounts6)
	return a.Limits
}

func (a *ATMTransaction14) SetMixType(value string) {
	a.MixType = (*Max35Text)(&value)
}

func (a *ATMTransaction14) AddMix() *ATMMediaMix1 {
	newValue := new(ATMMediaMix1)
	a.Mix = append(a.Mix, newValue)
	return newValue
}

func (a *ATMTransaction14) AddAuthorisationResult() *AuthorisationResult13 {
	a.AuthorisationResult = new(AuthorisationResult13)
	return a.AuthorisationResult
}

func (a *ATMTransaction14) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}

func (a *ATMTransaction14) AddCommand() *ATMCommand7 {
	newValue := new(ATMCommand7)
	a.Command = append(a.Command, newValue)
	return newValue
}

// Deposit transaction for which the service is requested.
type ATMTransaction15 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the reconciliation period assigned by the ATM.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Unprotected account information.
	AccountData *CardAccount9 `xml:"AcctData,omitempty"`

	// Encryption of account information.
	ProtectedAccountData *ContentInformationType10 `xml:"PrtctdAcctData,omitempty"`

	// Total amount of the deposit transaction.
	TotalAmount *AmountAndCurrency1 `xml:"TtlAmt,omitempty"`

	// Amounts of the deposit transaction.
	DetailedRequestedAmount []*DetailedAmount16 `xml:"DtldReqdAmt,omitempty"`

	// Deposited media put in the safe.
	DepositedMedia []*ATMDepositedMedia1 `xml:"DpstdMdia,omitempty"`

	// True if a receipt has be requested by the customer.
	RequestedReceipt *TrueFalseIndicator `xml:"ReqdRct,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (a *ATMTransaction15) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction15) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction15) AddAccountData() *CardAccount9 {
	a.AccountData = new(CardAccount9)
	return a.AccountData
}

func (a *ATMTransaction15) AddProtectedAccountData() *ContentInformationType10 {
	a.ProtectedAccountData = new(ContentInformationType10)
	return a.ProtectedAccountData
}

func (a *ATMTransaction15) AddTotalAmount() *AmountAndCurrency1 {
	a.TotalAmount = new(AmountAndCurrency1)
	return a.TotalAmount
}

func (a *ATMTransaction15) AddDetailedRequestedAmount() *DetailedAmount16 {
	newValue := new(DetailedAmount16)
	a.DetailedRequestedAmount = append(a.DetailedRequestedAmount, newValue)
	return newValue
}

func (a *ATMTransaction15) AddDepositedMedia() *ATMDepositedMedia1 {
	newValue := new(ATMDepositedMedia1)
	a.DepositedMedia = append(a.DepositedMedia, newValue)
	return newValue
}

func (a *ATMTransaction15) SetRequestedReceipt(value string) {
	a.RequestedReceipt = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction15) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}

// Response to the deposit request.
type ATMTransaction16 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the reconciliation period assigned by the ATM.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// True if a completion advice has to be sent after the end of the transaction.
	CompletionRequired *TrueFalseIndicator `xml:"CmpltnReqrd,omitempty"`

	// Unprotected account information.
	AccountData *CardAccount10 `xml:"AcctData,omitempty"`

	// Encryption of account information.
	ProtectedAccountData *ContentInformationType10 `xml:"PrtctdAcctData,omitempty"`

	// Total authorised amount of the deposit transaction.
	TotalAuthorisedAmount *AmountAndCurrency1 `xml:"TtlAuthrsdAmt"`

	// Total requested amount.
	TotalRequestedAmount *ImpliedCurrencyAndAmount `xml:"TtlReqdAmt,omitempty"`

	// Detail of the requested amounts for the deposit transaction.
	DetailedRequestedAmount []*DetailedAmount16 `xml:"DtldReqdAmt,omitempty"`

	// Additional charge (for instance tax or fee).
	AdditionalCharge []*DetailedAmount13 `xml:"AddtlChrg,omitempty"`

	// Outcome of the deposit authorisation.
	AuthorisationResult *AuthorisationResult13 `xml:"AuthstnRslt,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`

	// Maintenance command to perform on the ATM.
	Command []*ATMCommand7 `xml:"Cmd,omitempty"`
}

func (a *ATMTransaction16) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction16) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction16) SetCompletionRequired(value string) {
	a.CompletionRequired = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction16) AddAccountData() *CardAccount10 {
	a.AccountData = new(CardAccount10)
	return a.AccountData
}

func (a *ATMTransaction16) AddProtectedAccountData() *ContentInformationType10 {
	a.ProtectedAccountData = new(ContentInformationType10)
	return a.ProtectedAccountData
}

func (a *ATMTransaction16) AddTotalAuthorisedAmount() *AmountAndCurrency1 {
	a.TotalAuthorisedAmount = new(AmountAndCurrency1)
	return a.TotalAuthorisedAmount
}

func (a *ATMTransaction16) SetTotalRequestedAmount(value, currency string) {
	a.TotalRequestedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransaction16) AddDetailedRequestedAmount() *DetailedAmount16 {
	newValue := new(DetailedAmount16)
	a.DetailedRequestedAmount = append(a.DetailedRequestedAmount, newValue)
	return newValue
}

func (a *ATMTransaction16) AddAdditionalCharge() *DetailedAmount13 {
	newValue := new(DetailedAmount13)
	a.AdditionalCharge = append(a.AdditionalCharge, newValue)
	return newValue
}

func (a *ATMTransaction16) AddAuthorisationResult() *AuthorisationResult13 {
	a.AuthorisationResult = new(AuthorisationResult13)
	return a.AuthorisationResult
}

func (a *ATMTransaction16) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}

func (a *ATMTransaction16) AddCommand() *ATMCommand7 {
	newValue := new(ATMCommand7)
	a.Command = append(a.Command, newValue)
	return newValue
}

// Withdrawal transaction for which the completion is sent.
type ATMTransaction17 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Outcome of the financial transaction for the customer.
	TransactionStatus *ATMTransactionStatus1Code `xml:"TxSts"`

	// Multi bundle dispense.
	MultiBundle *TrueFalseIndicator `xml:"MultiBndl,omitempty"`

	// Amount per bundle in the currency of the total presented amount, in the order of the presentation.
	BundlePresentedAmount []*ImpliedCurrencyAndAmount `xml:"BndlPresntdAmt,omitempty"`

	// Status of the amount presented to the customer in the last bundle.
	PresentedAmountStatus *ATMTransactionStatus2Code `xml:"PresntdAmtSts"`

	// Incident occurring during the processing of the transaction.
	Incident []*FailureReason7Code `xml:"Incdnt,omitempty"`

	// Explanation of the incident.
	IncidentDetail []*Max70Text `xml:"IncdntDtl,omitempty"`

	// Identification of the reconciliation period assigned by the ATM.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Unprotected account information.
	AccountData *CardAccount11 `xml:"AcctData,omitempty"`

	// Encryption of account information.
	ProtectedAccountData *ContentInformationType10 `xml:"PrtctdAcctData,omitempty"`

	// Total amount presented to the customer.
	TotalPresentedAmount *AmountAndCurrency1 `xml:"TtlPresntdAmt"`

	// Total authorised amount.
	TotalAuthorisedAmount *ImpliedCurrencyAndAmount `xml:"TtlAuthrsdAmt,omitempty"`

	// Total requested amount.
	TotalRequestedAmount *ImpliedCurrencyAndAmount `xml:"TtlReqdAmt,omitempty"`

	// Detail of the requested amounts for the withdrawal transaction.
	DetailedRequestedAmount *DetailedAmount12 `xml:"DtldReqdAmt,omitempty"`

	// Currency conversion accepted by the customer, either to convert the amount to dispense in the base currency of the ATM, or to convert the total requested amount in the currency of the customer (so called dynamic currency conversion).
	CurrencyConversionResult *CurrencyConversion9 `xml:"CcyConvsRslt,omitempty"`

	// Additional charge (for instance tax or fee).
	AdditionalCharge []*DetailedAmount13 `xml:"AddtlChrg,omitempty"`

	// True if the customer has requested a receipt.
	RequestedReceipt *TrueFalseIndicator `xml:"ReqdRct,omitempty"`

	// True if a receipt has been printed and presented to the customer.
	ReceiptPrinted *TrueFalseIndicator `xml:"RctPrtd,omitempty"`

	// Explicit consent expressed by a customer on a card-related service proposed by an acquirer or an issuer or any agent acting on behalf of one of them.
	CustomerConsent *TrueFalseIndicator `xml:"CstmrCnsnt,omitempty"`

	// Limit of amounts for the customer.
	Limits *ATMTransactionAmounts6 `xml:"Lmts,omitempty"`

	// Outcome of the withdrawal authorisation.
	AuthorisationResult *AuthorisationResult13 `xml:"AuthstnRslt,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`

	// Current totals of the ATM.
	ATMTotals []*ATMTotals1 `xml:"ATMTtls,omitempty"`

	// Information on the cassettes of the ATM.
	Cassette []*ATMCassette2 `xml:"Csstt,omitempty"`
}

func (a *ATMTransaction17) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction17) SetTransactionStatus(value string) {
	a.TransactionStatus = (*ATMTransactionStatus1Code)(&value)
}

func (a *ATMTransaction17) SetMultiBundle(value string) {
	a.MultiBundle = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction17) AddBundlePresentedAmount(value, currency string) {
	a.BundlePresentedAmount = append(a.BundlePresentedAmount, NewImpliedCurrencyAndAmount(value, currency))
}

func (a *ATMTransaction17) SetPresentedAmountStatus(value string) {
	a.PresentedAmountStatus = (*ATMTransactionStatus2Code)(&value)
}

func (a *ATMTransaction17) AddIncident(value string) {
	a.Incident = append(a.Incident, (*FailureReason7Code)(&value))
}

func (a *ATMTransaction17) AddIncidentDetail(value string) {
	a.IncidentDetail = append(a.IncidentDetail, (*Max70Text)(&value))
}

func (a *ATMTransaction17) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction17) AddAccountData() *CardAccount11 {
	a.AccountData = new(CardAccount11)
	return a.AccountData
}

func (a *ATMTransaction17) AddProtectedAccountData() *ContentInformationType10 {
	a.ProtectedAccountData = new(ContentInformationType10)
	return a.ProtectedAccountData
}

func (a *ATMTransaction17) AddTotalPresentedAmount() *AmountAndCurrency1 {
	a.TotalPresentedAmount = new(AmountAndCurrency1)
	return a.TotalPresentedAmount
}

func (a *ATMTransaction17) SetTotalAuthorisedAmount(value, currency string) {
	a.TotalAuthorisedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransaction17) SetTotalRequestedAmount(value, currency string) {
	a.TotalRequestedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransaction17) AddDetailedRequestedAmount() *DetailedAmount12 {
	a.DetailedRequestedAmount = new(DetailedAmount12)
	return a.DetailedRequestedAmount
}

func (a *ATMTransaction17) AddCurrencyConversionResult() *CurrencyConversion9 {
	a.CurrencyConversionResult = new(CurrencyConversion9)
	return a.CurrencyConversionResult
}

func (a *ATMTransaction17) AddAdditionalCharge() *DetailedAmount13 {
	newValue := new(DetailedAmount13)
	a.AdditionalCharge = append(a.AdditionalCharge, newValue)
	return newValue
}

func (a *ATMTransaction17) SetRequestedReceipt(value string) {
	a.RequestedReceipt = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction17) SetReceiptPrinted(value string) {
	a.ReceiptPrinted = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction17) SetCustomerConsent(value string) {
	a.CustomerConsent = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction17) AddLimits() *ATMTransactionAmounts6 {
	a.Limits = new(ATMTransactionAmounts6)
	return a.Limits
}

func (a *ATMTransaction17) AddAuthorisationResult() *AuthorisationResult13 {
	a.AuthorisationResult = new(AuthorisationResult13)
	return a.AuthorisationResult
}

func (a *ATMTransaction17) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}

func (a *ATMTransaction17) AddATMTotals() *ATMTotals1 {
	newValue := new(ATMTotals1)
	a.ATMTotals = append(a.ATMTotals, newValue)
	return newValue
}

func (a *ATMTransaction17) AddCassette() *ATMCassette2 {
	newValue := new(ATMCassette2)
	a.Cassette = append(a.Cassette, newValue)
	return newValue
}

// Acknowledgement of a completion advice.
type ATMTransaction18 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Response to the advice.
	Response *Response4Code `xml:"Rspn"`

	// Detail of the response.
	ResponseReason *ResultDetail4Code `xml:"RspnRsn,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`

	// Maintenance command to perform on the ATM.
	Command []*ATMCommand7 `xml:"Cmd,omitempty"`
}

func (a *ATMTransaction18) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction18) SetResponse(value string) {
	a.Response = (*Response4Code)(&value)
}

func (a *ATMTransaction18) SetResponseReason(value string) {
	a.ResponseReason = (*ResultDetail4Code)(&value)
}

func (a *ATMTransaction18) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}

func (a *ATMTransaction18) AddCommand() *ATMCommand7 {
	newValue := new(ATMCommand7)
	a.Command = append(a.Command, newValue)
	return newValue
}

// Response to the deposit request.
type ATMTransaction19 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Outcome of the financial transaction for the customer.
	TransactionStatus *ATMTransactionStatus1Code `xml:"TxSts"`

	// Identification of the reconciliation period assigned by the ATM.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Incident occurring during the processing of the transaction.
	Incident []*FailureReason7Code `xml:"Incdnt,omitempty"`

	// Explanation of the incident.
	IncidentDetail []*Max70Text `xml:"IncdntDtl,omitempty"`

	// Unprotected account information.
	AccountData *CardAccount8 `xml:"AcctData,omitempty"`

	// Encryption of account information.
	ProtectedAccountData *ContentInformationType10 `xml:"PrtctdAcctData,omitempty"`

	// Total amount deposed by the customer.
	TotalDepositedAmount *AmountAndCurrency1 `xml:"TtlDpstdAmt"`

	// Total authorised amount of the deposit transaction.
	TotalAuthorisedAmount *ImpliedCurrencyAndAmount `xml:"TtlAuthrsdAmt"`

	// Total requested amount.
	TotalRequestedAmount *ImpliedCurrencyAndAmount `xml:"TtlReqdAmt,omitempty"`

	// Amounts of the deposit transaction.
	DetailedRequestedAmount *DetailedAmount16 `xml:"DtldReqdAmt,omitempty"`

	// Additional charge (for instance tax or fee).
	AdditionalCharge []*DetailedAmount13 `xml:"AddtlChrg,omitempty"`

	// True if the customer has requested a receipt.
	RequestedReceipt *TrueFalseIndicator `xml:"ReqdRct,omitempty"`

	// True if a receipt has been printed and presented to the customer.
	ReceiptPrinted *TrueFalseIndicator `xml:"RctPrtd,omitempty"`

	// Outcome of the deposit authorisation.
	AuthorisationResult *AuthorisationResult13 `xml:"AuthstnRslt,omitempty"`

	// Deposited media put in the safe.
	DepositedMedia []*ATMDepositedMedia1 `xml:"DpstdMdia,omitempty"`

	// Media unit not put in the safe. These deposits have to be reconciliated.
	ToBeReconciledMediaCounters []*ATMDepositedMedia3 `xml:"ToBeRcncldMdiaCntrs,omitempty"`

	// Current totals of the ATM.
	ATMTotals []*ATMTotals1 `xml:"ATMTtls,omitempty"`

	// Information on the cassettes of the ATM.
	Cassette []*ATMCassette2 `xml:"Csstt,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (a *ATMTransaction19) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction19) SetTransactionStatus(value string) {
	a.TransactionStatus = (*ATMTransactionStatus1Code)(&value)
}

func (a *ATMTransaction19) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction19) AddIncident(value string) {
	a.Incident = append(a.Incident, (*FailureReason7Code)(&value))
}

func (a *ATMTransaction19) AddIncidentDetail(value string) {
	a.IncidentDetail = append(a.IncidentDetail, (*Max70Text)(&value))
}

func (a *ATMTransaction19) AddAccountData() *CardAccount8 {
	a.AccountData = new(CardAccount8)
	return a.AccountData
}

func (a *ATMTransaction19) AddProtectedAccountData() *ContentInformationType10 {
	a.ProtectedAccountData = new(ContentInformationType10)
	return a.ProtectedAccountData
}

func (a *ATMTransaction19) AddTotalDepositedAmount() *AmountAndCurrency1 {
	a.TotalDepositedAmount = new(AmountAndCurrency1)
	return a.TotalDepositedAmount
}

func (a *ATMTransaction19) SetTotalAuthorisedAmount(value, currency string) {
	a.TotalAuthorisedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransaction19) SetTotalRequestedAmount(value, currency string) {
	a.TotalRequestedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransaction19) AddDetailedRequestedAmount() *DetailedAmount16 {
	a.DetailedRequestedAmount = new(DetailedAmount16)
	return a.DetailedRequestedAmount
}

func (a *ATMTransaction19) AddAdditionalCharge() *DetailedAmount13 {
	newValue := new(DetailedAmount13)
	a.AdditionalCharge = append(a.AdditionalCharge, newValue)
	return newValue
}

func (a *ATMTransaction19) SetRequestedReceipt(value string) {
	a.RequestedReceipt = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction19) SetReceiptPrinted(value string) {
	a.ReceiptPrinted = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction19) AddAuthorisationResult() *AuthorisationResult13 {
	a.AuthorisationResult = new(AuthorisationResult13)
	return a.AuthorisationResult
}

func (a *ATMTransaction19) AddDepositedMedia() *ATMDepositedMedia1 {
	newValue := new(ATMDepositedMedia1)
	a.DepositedMedia = append(a.DepositedMedia, newValue)
	return newValue
}

func (a *ATMTransaction19) AddToBeReconciledMediaCounters() *ATMDepositedMedia3 {
	newValue := new(ATMDepositedMedia3)
	a.ToBeReconciledMediaCounters = append(a.ToBeReconciledMediaCounters, newValue)
	return newValue
}

func (a *ATMTransaction19) AddATMTotals() *ATMTotals1 {
	newValue := new(ATMTotals1)
	a.ATMTotals = append(a.ATMTotals, newValue)
	return newValue
}

func (a *ATMTransaction19) AddCassette() *ATMCassette2 {
	newValue := new(ATMCassette2)
	a.Cassette = append(a.Cassette, newValue)
	return newValue
}

func (a *ATMTransaction19) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}

// Response to the withdrawal transaction request.
type ATMTransaction2 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the reconciliation period assigned by the ATM.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// True if a withdrawal completion advice has to be sent after the end of the transaction.
	CompletionRequired *TrueFalseIndicator `xml:"CmpltnReqrd,omitempty"`

	// Unprotected account information.
	AccountData *CardAccount4 `xml:"AcctData,omitempty"`

	// Encryption of account information.
	ProtectedAccountData *ContentInformationType10 `xml:"PrtctdAcctData,omitempty"`

	// Total authorised amount.
	TotalAuthorisedAmount *AmountAndCurrency1 `xml:"TtlAuthrsdAmt"`

	// Total requested amount.
	TotalRequestedAmount *ImpliedCurrencyAndAmount `xml:"TtlReqdAmt,omitempty"`

	// Detail of the requested amounts for the withdrawal transaction.
	DetailedRequestedAmount *DetailedAmount12 `xml:"DtldReqdAmt,omitempty"`

	// Currency conversion accepted by the customer, either to convert the amount to dispense in the base currency of the ATM, or to convert the total requested amount in the currency of the customer (so called dynamic currency conversion).
	CurrencyConversion *CurrencyConversion4 `xml:"CcyConvs,omitempty"`

	// Additional charge (for instance tax or fee).
	AdditionalCharge []*DetailedAmount13 `xml:"AddtlChrg,omitempty"`

	// Limit of amounts for the customer.
	Limits *ATMTransactionAmounts2 `xml:"Lmts,omitempty"`

	// Media mix algorithm requested by the ATM Host, the identification of the algorithm is an agreement between the ATM and the ATM manager.
	MixType *Max35Text `xml:"MixTp,omitempty"`

	// Media mix selected requested by the ATM Host.
	Mix []*ATMMediaMix1 `xml:"Mix,omitempty"`

	// Outcome of the withdrawal authorisation.
	AuthorisationResult *AuthorisationResult9 `xml:"AuthstnRslt"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`

	// Maintenance command to perform on the ATM.
	Command []*ATMCommand1 `xml:"Cmd,omitempty"`
}

func (a *ATMTransaction2) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction2) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction2) SetCompletionRequired(value string) {
	a.CompletionRequired = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction2) AddAccountData() *CardAccount4 {
	a.AccountData = new(CardAccount4)
	return a.AccountData
}

func (a *ATMTransaction2) AddProtectedAccountData() *ContentInformationType10 {
	a.ProtectedAccountData = new(ContentInformationType10)
	return a.ProtectedAccountData
}

func (a *ATMTransaction2) AddTotalAuthorisedAmount() *AmountAndCurrency1 {
	a.TotalAuthorisedAmount = new(AmountAndCurrency1)
	return a.TotalAuthorisedAmount
}

func (a *ATMTransaction2) SetTotalRequestedAmount(value, currency string) {
	a.TotalRequestedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransaction2) AddDetailedRequestedAmount() *DetailedAmount12 {
	a.DetailedRequestedAmount = new(DetailedAmount12)
	return a.DetailedRequestedAmount
}

func (a *ATMTransaction2) AddCurrencyConversion() *CurrencyConversion4 {
	a.CurrencyConversion = new(CurrencyConversion4)
	return a.CurrencyConversion
}

func (a *ATMTransaction2) AddAdditionalCharge() *DetailedAmount13 {
	newValue := new(DetailedAmount13)
	a.AdditionalCharge = append(a.AdditionalCharge, newValue)
	return newValue
}

func (a *ATMTransaction2) AddLimits() *ATMTransactionAmounts2 {
	a.Limits = new(ATMTransactionAmounts2)
	return a.Limits
}

func (a *ATMTransaction2) SetMixType(value string) {
	a.MixType = (*Max35Text)(&value)
}

func (a *ATMTransaction2) AddMix() *ATMMediaMix1 {
	newValue := new(ATMMediaMix1)
	a.Mix = append(a.Mix, newValue)
	return newValue
}

func (a *ATMTransaction2) AddAuthorisationResult() *AuthorisationResult9 {
	a.AuthorisationResult = new(AuthorisationResult9)
	return a.AuthorisationResult
}

func (a *ATMTransaction2) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}

func (a *ATMTransaction2) AddCommand() *ATMCommand1 {
	newValue := new(ATMCommand1)
	a.Command = append(a.Command, newValue)
	return newValue
}

// Withdrawal transaction for which the completion is sent.
type ATMTransaction20 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Outcome of the financial transaction for the customer.
	TransactionStatus *ATMTransactionStatus1Code `xml:"TxSts"`

	// Incident occurring during the processing of the transaction.
	Incident []*FailureReason7Code `xml:"Incdnt,omitempty"`

	// Explanation of the incident.
	IncidentDetail []*Max70Text `xml:"IncdntDtl,omitempty"`

	// Identification of the reconciliation period assigned by the ATM.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// True if the customer has requested a receipt.
	RequestedReceipt *TrueFalseIndicator `xml:"ReqdRct,omitempty"`

	// True if a receipt has been printed and presented to the customer.
	ReceiptPrinted *TrueFalseIndicator `xml:"RctPrtd,omitempty"`

	// Explicit consent expressed by a customer on a card-related service proposed by an acquirer or an issuer or any agent acting on behalf of one of them.
	CustomerConsent *TrueFalseIndicator `xml:"CstmrCnsnt,omitempty"`

	// Outcome of the withdrawal authorisation.
	AuthorisationResult *AuthorisationResult13 `xml:"AuthstnRslt,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (a *ATMTransaction20) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction20) SetTransactionStatus(value string) {
	a.TransactionStatus = (*ATMTransactionStatus1Code)(&value)
}

func (a *ATMTransaction20) AddIncident(value string) {
	a.Incident = append(a.Incident, (*FailureReason7Code)(&value))
}

func (a *ATMTransaction20) AddIncidentDetail(value string) {
	a.IncidentDetail = append(a.IncidentDetail, (*Max70Text)(&value))
}

func (a *ATMTransaction20) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction20) SetRequestedReceipt(value string) {
	a.RequestedReceipt = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction20) SetReceiptPrinted(value string) {
	a.ReceiptPrinted = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction20) SetCustomerConsent(value string) {
	a.CustomerConsent = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction20) AddAuthorisationResult() *AuthorisationResult13 {
	a.AuthorisationResult = new(AuthorisationResult13)
	return a.AuthorisationResult
}

func (a *ATMTransaction20) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}

// Inquiry information for the transaction.
type ATMTransaction21 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Result of the inquiry service.
	TransactionResponse *ResponseType7 `xml:"TxRspn"`

	// Sequence of actions to be performed by the ATM to complete the transaction.
	Action []*Action7 `xml:"Actn,omitempty"`

	// Profile of the customer with the allowed services and restrictions.
	CustomerServiceProfile *ATMCustomerProfile5 `xml:"CstmrSvcPrfl,omitempty"`

	// Dynamic currency conversion result.
	CurrencyConversion *CurrencyConversion10 `xml:"CcyConvs,omitempty"`

	// Account information associated to the customer.
	AccountInformation []*CardAccount12 `xml:"AcctInf,omitempty"`

	// Statement information of an account.
	AccountStatementData []*ATMAccountStatement1 `xml:"AcctStmtData,omitempty"`

	// Exchange rate and calculated amount to be presented to the customer when the dispense currency or the deposit currency (target currency) is different to account currency (source currency).
	CurrencyExchange *CurrencyConversion5 `xml:"CcyXchg,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`

	// Maintenance command to perform on the ATM.
	Command []*ATMCommand7 `xml:"Cmd,omitempty"`
}

func (a *ATMTransaction21) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction21) AddTransactionResponse() *ResponseType7 {
	a.TransactionResponse = new(ResponseType7)
	return a.TransactionResponse
}

func (a *ATMTransaction21) AddAction() *Action7 {
	newValue := new(Action7)
	a.Action = append(a.Action, newValue)
	return newValue
}

func (a *ATMTransaction21) AddCustomerServiceProfile() *ATMCustomerProfile5 {
	a.CustomerServiceProfile = new(ATMCustomerProfile5)
	return a.CustomerServiceProfile
}

func (a *ATMTransaction21) AddCurrencyConversion() *CurrencyConversion10 {
	a.CurrencyConversion = new(CurrencyConversion10)
	return a.CurrencyConversion
}

func (a *ATMTransaction21) AddAccountInformation() *CardAccount12 {
	newValue := new(CardAccount12)
	a.AccountInformation = append(a.AccountInformation, newValue)
	return newValue
}

func (a *ATMTransaction21) AddAccountStatementData() *ATMAccountStatement1 {
	newValue := new(ATMAccountStatement1)
	a.AccountStatementData = append(a.AccountStatementData, newValue)
	return newValue
}

func (a *ATMTransaction21) AddCurrencyExchange() *CurrencyConversion5 {
	a.CurrencyExchange = new(CurrencyConversion5)
	return a.CurrencyExchange
}

func (a *ATMTransaction21) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}

func (a *ATMTransaction21) AddCommand() *ATMCommand7 {
	newValue := new(ATMCommand7)
	a.Command = append(a.Command, newValue)
	return newValue
}

// Response to the PIN management transaction. request.
type ATMTransaction22 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the reconciliation period assigned by the ATM.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// True if a completion advice has to be sent after the end of the transaction.
	CompletionRequired *TrueFalseIndicator `xml:"CmpltnReqrd,omitempty"`

	// Result of the PIN service.
	TransactionResponse *ResponseType7 `xml:"TxRspn"`

	// Sequence of actions to be performed by the ATM to complete the transaction.
	Action []*Action7 `xml:"Actn,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`

	// Maintenance command to perform on the ATM.
	Command []*ATMCommand7 `xml:"Cmd,omitempty"`
}

func (a *ATMTransaction22) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction22) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction22) SetCompletionRequired(value string) {
	a.CompletionRequired = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction22) AddTransactionResponse() *ResponseType7 {
	a.TransactionResponse = new(ResponseType7)
	return a.TransactionResponse
}

func (a *ATMTransaction22) AddAction() *Action7 {
	newValue := new(Action7)
	a.Action = append(a.Action, newValue)
	return newValue
}

func (a *ATMTransaction22) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}

func (a *ATMTransaction22) AddCommand() *ATMCommand7 {
	newValue := new(ATMCommand7)
	a.Command = append(a.Command, newValue)
	return newValue
}

// Transfer information for the transaction.
type ATMTransaction23 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the reconciliation period assigned by the ATM.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Description of the transfer for the creditor.
	CreditorLabel *Max35Text `xml:"CdtrLabl,omitempty"`

	// Description of the transfer for the debtor.
	DebtorLabel *Max35Text `xml:"DbtrLabl,omitempty"`

	// Reference of the payment.
	PaymentReference *Max35Text `xml:"PmtRef,omitempty"`

	// Information about the source account of the transfer.
	AccountFrom *CardAccount7 `xml:"AcctFr,omitempty"`

	// Encryption of the source account information.
	ProtectedAccountFrom *ContentInformationType10 `xml:"PrtctdAcctFr,omitempty"`

	// Information about the destination account of the transfer.
	AccountTo []*CardAccount7 `xml:"AcctTo,omitempty"`

	// Encryption of the destination account information.
	ProtectedAccountTo *ContentInformationType10 `xml:"PrtctdAcctTo,omitempty"`

	// Amount of the transaction to be authorised.
	TotalRequestedAmount *ImpliedCurrencyAndAmount `xml:"TtlReqdAmt,omitempty"`

	// Details of the transfer transaction amounts.
	DetailedRequestedAmount *DetailedAmount17 `xml:"DtldReqdAmt,omitempty"`

	// Requested date of the execution of the transfer.
	RequestedExecutionDate *ISODate `xml:"ReqdExctnDt,omitempty"`

	// Identifies the instant transfer program.
	InstantTransferProgram *Max35Text `xml:"InstntTrfPrgm,omitempty"`

	// Information for reccurring transfer or standing orders.
	RecurringTransfer *RecurringTransaction3 `xml:"RcrngTrf,omitempty"`

	// True if the customer has requested a receipt.
	RequestedReceipt *TrueFalseIndicator `xml:"ReqdRct,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (a *ATMTransaction23) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction23) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction23) SetCreditorLabel(value string) {
	a.CreditorLabel = (*Max35Text)(&value)
}

func (a *ATMTransaction23) SetDebtorLabel(value string) {
	a.DebtorLabel = (*Max35Text)(&value)
}

func (a *ATMTransaction23) SetPaymentReference(value string) {
	a.PaymentReference = (*Max35Text)(&value)
}

func (a *ATMTransaction23) AddAccountFrom() *CardAccount7 {
	a.AccountFrom = new(CardAccount7)
	return a.AccountFrom
}

func (a *ATMTransaction23) AddProtectedAccountFrom() *ContentInformationType10 {
	a.ProtectedAccountFrom = new(ContentInformationType10)
	return a.ProtectedAccountFrom
}

func (a *ATMTransaction23) AddAccountTo() *CardAccount7 {
	newValue := new(CardAccount7)
	a.AccountTo = append(a.AccountTo, newValue)
	return newValue
}

func (a *ATMTransaction23) AddProtectedAccountTo() *ContentInformationType10 {
	a.ProtectedAccountTo = new(ContentInformationType10)
	return a.ProtectedAccountTo
}

func (a *ATMTransaction23) SetTotalRequestedAmount(value, currency string) {
	a.TotalRequestedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransaction23) AddDetailedRequestedAmount() *DetailedAmount17 {
	a.DetailedRequestedAmount = new(DetailedAmount17)
	return a.DetailedRequestedAmount
}

func (a *ATMTransaction23) SetRequestedExecutionDate(value string) {
	a.RequestedExecutionDate = (*ISODate)(&value)
}

func (a *ATMTransaction23) SetInstantTransferProgram(value string) {
	a.InstantTransferProgram = (*Max35Text)(&value)
}

func (a *ATMTransaction23) AddRecurringTransfer() *RecurringTransaction3 {
	a.RecurringTransfer = new(RecurringTransaction3)
	return a.RecurringTransfer
}

func (a *ATMTransaction23) SetRequestedReceipt(value string) {
	a.RequestedReceipt = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction23) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}

// Transfer information for the transaction.
type ATMTransaction24 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the reconciliation period assigned by the ATM.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Description of the transfer for the creditor.
	CreditorLabel *Max35Text `xml:"CdtrLabl,omitempty"`

	// Description of the transfer for the debtor.
	DebtorLabel *Max35Text `xml:"DbtrLabl,omitempty"`

	// Identifier of the approved transfer transaction for the bank.
	TransferIdentifier *Max70Text `xml:"TrfIdr,omitempty"`

	// Reference of the payment.
	PaymentReference *Max35Text `xml:"PmtRef,omitempty"`

	// Result of the fund transfer request.
	TransactionResponse *ResponseType7 `xml:"TxRspn"`

	// Sequence of actions to be performed by the ATM to complete the transaction.
	Action []*Action7 `xml:"Actn,omitempty"`

	// Information about the source account of the transfer.
	AccountFrom *CardAccount13 `xml:"AcctFr,omitempty"`

	// Encryption of the source account information.
	ProtectedAccountFrom *ContentInformationType10 `xml:"PrtctdAcctFr,omitempty"`

	// Information about the destination account of the transfer.
	AccountTo []*CardAccount13 `xml:"AcctTo,omitempty"`

	// Encryption of the destination account information.
	ProtectedAccountTo *ContentInformationType10 `xml:"PrtctdAcctTo,omitempty"`

	// Total authorised amount.
	TotalAuthorisedAmount *AmountAndCurrency1 `xml:"TtlAuthrsdAmt"`

	// Total requested amount.
	TotalRequestedAmount *ImpliedCurrencyAndAmount `xml:"TtlReqdAmt,omitempty"`

	// Details of the transfer transaction amounts.
	DetailedRequestedAmount *DetailedAmount17 `xml:"DtldReqdAmt,omitempty"`

	// Additional charge (for instance tax or fee).
	AdditionalCharge []*DetailedAmount18 `xml:"AddtlChrg,omitempty"`

	// Limit of amounts for the customer.
	Limits *ATMTransactionAmounts6 `xml:"Lmts,omitempty"`

	// Requested date of the execution of the transfer.
	RequestedExecutionDate *ISODate `xml:"ReqdExctnDt,omitempty"`

	// Proposed date of the execution of the transfer.
	ProposedExecutionDate *ISODate `xml:"PropsdExctnDt,omitempty"`

	// Identifies the instant transfer program.
	InstantTransferProgram *Max35Text `xml:"InstntTrfPrgm,omitempty"`

	// Information for reccurring transfer or standing orders.
	RecurringTransfer *RecurringTransaction3 `xml:"RcrngTrf,omitempty"`

	// Outcome of the transfer authorisation.
	AuthorisationResult *AuthorisationResult13 `xml:"AuthstnRslt,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`

	// Maintenance command to perform on the ATM.
	Command []*ATMCommand7 `xml:"Cmd,omitempty"`
}

func (a *ATMTransaction24) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction24) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction24) SetCreditorLabel(value string) {
	a.CreditorLabel = (*Max35Text)(&value)
}

func (a *ATMTransaction24) SetDebtorLabel(value string) {
	a.DebtorLabel = (*Max35Text)(&value)
}

func (a *ATMTransaction24) SetTransferIdentifier(value string) {
	a.TransferIdentifier = (*Max70Text)(&value)
}

func (a *ATMTransaction24) SetPaymentReference(value string) {
	a.PaymentReference = (*Max35Text)(&value)
}

func (a *ATMTransaction24) AddTransactionResponse() *ResponseType7 {
	a.TransactionResponse = new(ResponseType7)
	return a.TransactionResponse
}

func (a *ATMTransaction24) AddAction() *Action7 {
	newValue := new(Action7)
	a.Action = append(a.Action, newValue)
	return newValue
}

func (a *ATMTransaction24) AddAccountFrom() *CardAccount13 {
	a.AccountFrom = new(CardAccount13)
	return a.AccountFrom
}

func (a *ATMTransaction24) AddProtectedAccountFrom() *ContentInformationType10 {
	a.ProtectedAccountFrom = new(ContentInformationType10)
	return a.ProtectedAccountFrom
}

func (a *ATMTransaction24) AddAccountTo() *CardAccount13 {
	newValue := new(CardAccount13)
	a.AccountTo = append(a.AccountTo, newValue)
	return newValue
}

func (a *ATMTransaction24) AddProtectedAccountTo() *ContentInformationType10 {
	a.ProtectedAccountTo = new(ContentInformationType10)
	return a.ProtectedAccountTo
}

func (a *ATMTransaction24) AddTotalAuthorisedAmount() *AmountAndCurrency1 {
	a.TotalAuthorisedAmount = new(AmountAndCurrency1)
	return a.TotalAuthorisedAmount
}

func (a *ATMTransaction24) SetTotalRequestedAmount(value, currency string) {
	a.TotalRequestedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransaction24) AddDetailedRequestedAmount() *DetailedAmount17 {
	a.DetailedRequestedAmount = new(DetailedAmount17)
	return a.DetailedRequestedAmount
}

func (a *ATMTransaction24) AddAdditionalCharge() *DetailedAmount18 {
	newValue := new(DetailedAmount18)
	a.AdditionalCharge = append(a.AdditionalCharge, newValue)
	return newValue
}

func (a *ATMTransaction24) AddLimits() *ATMTransactionAmounts6 {
	a.Limits = new(ATMTransactionAmounts6)
	return a.Limits
}

func (a *ATMTransaction24) SetRequestedExecutionDate(value string) {
	a.RequestedExecutionDate = (*ISODate)(&value)
}

func (a *ATMTransaction24) SetProposedExecutionDate(value string) {
	a.ProposedExecutionDate = (*ISODate)(&value)
}

func (a *ATMTransaction24) SetInstantTransferProgram(value string) {
	a.InstantTransferProgram = (*Max35Text)(&value)
}

func (a *ATMTransaction24) AddRecurringTransfer() *RecurringTransaction3 {
	a.RecurringTransfer = new(RecurringTransaction3)
	return a.RecurringTransfer
}

func (a *ATMTransaction24) AddAuthorisationResult() *AuthorisationResult13 {
	a.AuthorisationResult = new(AuthorisationResult13)
	return a.AuthorisationResult
}

func (a *ATMTransaction24) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}

func (a *ATMTransaction24) AddCommand() *ATMCommand7 {
	newValue := new(ATMCommand7)
	a.Command = append(a.Command, newValue)
	return newValue
}

// Information about the reconciliation request.
type ATMTransaction25 struct {

	// Type of logical or physical operation on the ATM for which the counters are computed.
	TypeOfOperation *ATMOperation1Code `xml:"TpOfOpr,omitempty"`

	// Identification of the reconciliation transaction.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the reconciliation period.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId"`

	// Current totals of the ATM.
	ATMTotals []*ATMTotals1 `xml:"ATMTtls,omitempty"`

	// Information on the cassette of the ATM.
	Cassette []*ATMCassette2 `xml:"Csstt,omitempty"`

	// Transaction counters that are set to zero after a reconciliation with counter reinitialisation command.
	TransactionTotals []*ATMTotals3 `xml:"TxTtls,omitempty"`

	// Total number of retained cards.
	RetainedCard *Number `xml:"RtndCard,omitempty"`

	// Additional information about reconciliation.
	AdditionalTransactionInformation *Max140Text `xml:"AddtlTxInf,omitempty"`
}

func (a *ATMTransaction25) SetTypeOfOperation(value string) {
	a.TypeOfOperation = (*ATMOperation1Code)(&value)
}

func (a *ATMTransaction25) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction25) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction25) AddATMTotals() *ATMTotals1 {
	newValue := new(ATMTotals1)
	a.ATMTotals = append(a.ATMTotals, newValue)
	return newValue
}

func (a *ATMTransaction25) AddCassette() *ATMCassette2 {
	newValue := new(ATMCassette2)
	a.Cassette = append(a.Cassette, newValue)
	return newValue
}

func (a *ATMTransaction25) AddTransactionTotals() *ATMTotals3 {
	newValue := new(ATMTotals3)
	a.TransactionTotals = append(a.TransactionTotals, newValue)
	return newValue
}

func (a *ATMTransaction25) SetRetainedCard(value string) {
	a.RetainedCard = (*Number)(&value)
}

func (a *ATMTransaction25) SetAdditionalTransactionInformation(value string) {
	a.AdditionalTransactionInformation = (*Max140Text)(&value)
}

// Information about the reconciliation response.
type ATMTransaction26 struct {

	// Type of logical or physical operation on the ATM for which the counters are computed.
	TypeOfOperation *ATMOperation1Code `xml:"TpOfOpr,omitempty"`

	// Identification of the reconciliation transaction.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the reconciliation period assigned by the ATM.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId"`

	// Result of the reconciliation.
	TransactionResponse *ResponseType7 `xml:"TxRspn"`

	// Current totals of the ATM.
	ATMTotals []*ATMTotals1 `xml:"ATMTtls,omitempty"`

	// Information on the cassettes of the ATM.
	Cassette []*ATMCassette2 `xml:"Csstt,omitempty"`

	// Maintenance command to perform on the ATM.
	Command []*ATMCommand7 `xml:"Cmd,omitempty"`
}

func (a *ATMTransaction26) SetTypeOfOperation(value string) {
	a.TypeOfOperation = (*ATMOperation1Code)(&value)
}

func (a *ATMTransaction26) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction26) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction26) AddTransactionResponse() *ResponseType7 {
	a.TransactionResponse = new(ResponseType7)
	return a.TransactionResponse
}

func (a *ATMTransaction26) AddATMTotals() *ATMTotals1 {
	newValue := new(ATMTotals1)
	a.ATMTotals = append(a.ATMTotals, newValue)
	return newValue
}

func (a *ATMTransaction26) AddCassette() *ATMCassette2 {
	newValue := new(ATMCassette2)
	a.Cassette = append(a.Cassette, newValue)
	return newValue
}

func (a *ATMTransaction26) AddCommand() *ATMCommand7 {
	newValue := new(ATMCommand7)
	a.Command = append(a.Command, newValue)
	return newValue
}

// Transaction for which the exception is sent.
type ATMTransaction27 struct {

	// Identification of the transaction.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId,omitempty"`

	// Identification of the reconciliation period.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Exception occurring outside the service.
	Exception []*FailureReason8Code `xml:"Xcptn"`

	// Explanation of the exception.
	ExceptionDetail []*Max70Text `xml:"XcptnDtl,omitempty"`

	// Balance of the captured card or epurse if available.
	ElectronicPurseBalance *CurrencyAndAmount `xml:"ElctrncPrsBal,omitempty"`
}

func (a *ATMTransaction27) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction27) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction27) AddException(value string) {
	a.Exception = append(a.Exception, (*FailureReason8Code)(&value))
}

func (a *ATMTransaction27) AddExceptionDetail(value string) {
	a.ExceptionDetail = append(a.ExceptionDetail, (*Max70Text)(&value))
}

func (a *ATMTransaction27) SetElectronicPurseBalance(value, currency string) {
	a.ElectronicPurseBalance = NewCurrencyAndAmount(value, currency)
}

// Acknowledgement of the exception advice.
type ATMTransaction28 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId,omitempty"`

	// Response to the advice.
	Response *Response2Code `xml:"Rspn"`

	// Maintenance command to perform on the ATM.
	Command []*ATMCommand7 `xml:"Cmd,omitempty"`
}

func (a *ATMTransaction28) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction28) SetResponse(value string) {
	a.Response = (*Response2Code)(&value)
}

func (a *ATMTransaction28) AddCommand() *ATMCommand7 {
	newValue := new(ATMCommand7)
	a.Command = append(a.Command, newValue)
	return newValue
}

// Inquiry information for the transaction.
type ATMTransaction29 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Unprotected account information.
	AccountData *CardAccount7 `xml:"AcctData,omitempty"`

	// Encryption of account information.
	ProtectedAccountData *ContentInformationType10 `xml:"PrtctdAcctData,omitempty"`

	// Amount to be authorised by the issuer.
	TotalRequestedAmount *AmountAndCurrency1 `xml:"TtlReqdAmt,omitempty"`

	// Amounts of the withdrawal transaction.
	DetailedRequestedAmount *DetailedAmount12 `xml:"DtldReqdAmt,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (a *ATMTransaction29) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction29) AddAccountData() *CardAccount7 {
	a.AccountData = new(CardAccount7)
	return a.AccountData
}

func (a *ATMTransaction29) AddProtectedAccountData() *ContentInformationType10 {
	a.ProtectedAccountData = new(ContentInformationType10)
	return a.ProtectedAccountData
}

func (a *ATMTransaction29) AddTotalRequestedAmount() *AmountAndCurrency1 {
	a.TotalRequestedAmount = new(AmountAndCurrency1)
	return a.TotalRequestedAmount
}

func (a *ATMTransaction29) AddDetailedRequestedAmount() *DetailedAmount12 {
	a.DetailedRequestedAmount = new(DetailedAmount12)
	return a.DetailedRequestedAmount
}

func (a *ATMTransaction29) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}

// Withdrawal transaction for which the completion is sent.
type ATMTransaction3 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Outcome of the financial transaction for the customer.
	TransactionStatus *ATMTransactionStatus1Code `xml:"TxSts"`

	// Multi bundle dispense.
	MultiBundle *TrueFalseIndicator `xml:"MultiBndl,omitempty"`

	// Amount per bundle in the currency of the total presented amount, in the order of the presentation.
	BundlePresentedAmount []*ImpliedCurrencyAndAmount `xml:"BndlPresntdAmt,omitempty"`

	// Status of the amount presented to the customer in the last bundle.
	PresentedAmountStatus *ATMTransactionStatus2Code `xml:"PresntdAmtSts"`

	// Incident occurring during the processing of the transaction.
	Incident []*FailureReason4Code `xml:"Incdnt,omitempty"`

	// Explanation of the incident.
	IncidentDetail []*Max70Text `xml:"IncdntDtl,omitempty"`

	// Identification of the reconciliation period assigned by the ATM.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Unprotected account information.
	AccountData *CardAccount5 `xml:"AcctData,omitempty"`

	// Encryption of account information.
	ProtectedAccountData *ContentInformationType10 `xml:"PrtctdAcctData,omitempty"`

	// Total amount presented to the customer.
	TotalPresentedAmount *AmountAndCurrency1 `xml:"TtlPresntdAmt"`

	// Total authorised amount.
	TotalAuthorisedAmount *ImpliedCurrencyAndAmount `xml:"TtlAuthrsdAmt,omitempty"`

	// Total requested amount.
	TotalRequestedAmount *ImpliedCurrencyAndAmount `xml:"TtlReqdAmt,omitempty"`

	// Detail of the requested amounts for the withdrawal transaction.
	DetailedRequestedAmount *DetailedAmount12 `xml:"DtldReqdAmt,omitempty"`

	// Currency conversion accepted by the customer, either to convert the amount to dispense in the base currency of the ATM, or to convert the total requested amount in the currency of the customer (so called dynamic currency conversion).
	CurrencyConversion *CurrencyConversion4 `xml:"CcyConvs,omitempty"`

	// Additional charge (for instance tax or fee).
	AdditionalCharge []*DetailedAmount13 `xml:"AddtlChrg,omitempty"`

	// True if the customer has requested a receipt.
	RequestedReceipt *TrueFalseIndicator `xml:"ReqdRct,omitempty"`

	// True if a receipt has been printed and presented to the customer.
	ReceiptPrinted *TrueFalseIndicator `xml:"RctPrtd,omitempty"`

	// True when the card was captured by the ATM.
	CapturedCard *TrueFalseIndicator `xml:"CaptrdCard,omitempty"`

	// Limit of amounts for the customer.
	Limits *ATMTransactionAmounts2 `xml:"Lmts,omitempty"`

	// Outcome of the withdrawal authorisation.
	AuthorisationResult *AuthorisationResult9 `xml:"AuthstnRslt,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`

	// Current totals of the ATM.
	ATMTotals []*ATMTotals1 `xml:"ATMTtls,omitempty"`

	// Information on the cassettes of the ATM.
	Cassette []*ATMCassette1 `xml:"Csstt,omitempty"`
}

func (a *ATMTransaction3) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction3) SetTransactionStatus(value string) {
	a.TransactionStatus = (*ATMTransactionStatus1Code)(&value)
}

func (a *ATMTransaction3) SetMultiBundle(value string) {
	a.MultiBundle = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction3) AddBundlePresentedAmount(value, currency string) {
	a.BundlePresentedAmount = append(a.BundlePresentedAmount, NewImpliedCurrencyAndAmount(value, currency))
}

func (a *ATMTransaction3) SetPresentedAmountStatus(value string) {
	a.PresentedAmountStatus = (*ATMTransactionStatus2Code)(&value)
}

func (a *ATMTransaction3) AddIncident(value string) {
	a.Incident = append(a.Incident, (*FailureReason4Code)(&value))
}

func (a *ATMTransaction3) AddIncidentDetail(value string) {
	a.IncidentDetail = append(a.IncidentDetail, (*Max70Text)(&value))
}

func (a *ATMTransaction3) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction3) AddAccountData() *CardAccount5 {
	a.AccountData = new(CardAccount5)
	return a.AccountData
}

func (a *ATMTransaction3) AddProtectedAccountData() *ContentInformationType10 {
	a.ProtectedAccountData = new(ContentInformationType10)
	return a.ProtectedAccountData
}

func (a *ATMTransaction3) AddTotalPresentedAmount() *AmountAndCurrency1 {
	a.TotalPresentedAmount = new(AmountAndCurrency1)
	return a.TotalPresentedAmount
}

func (a *ATMTransaction3) SetTotalAuthorisedAmount(value, currency string) {
	a.TotalAuthorisedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransaction3) SetTotalRequestedAmount(value, currency string) {
	a.TotalRequestedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransaction3) AddDetailedRequestedAmount() *DetailedAmount12 {
	a.DetailedRequestedAmount = new(DetailedAmount12)
	return a.DetailedRequestedAmount
}

func (a *ATMTransaction3) AddCurrencyConversion() *CurrencyConversion4 {
	a.CurrencyConversion = new(CurrencyConversion4)
	return a.CurrencyConversion
}

func (a *ATMTransaction3) AddAdditionalCharge() *DetailedAmount13 {
	newValue := new(DetailedAmount13)
	a.AdditionalCharge = append(a.AdditionalCharge, newValue)
	return newValue
}

func (a *ATMTransaction3) SetRequestedReceipt(value string) {
	a.RequestedReceipt = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction3) SetReceiptPrinted(value string) {
	a.ReceiptPrinted = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction3) SetCapturedCard(value string) {
	a.CapturedCard = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction3) AddLimits() *ATMTransactionAmounts2 {
	a.Limits = new(ATMTransactionAmounts2)
	return a.Limits
}

func (a *ATMTransaction3) AddAuthorisationResult() *AuthorisationResult9 {
	a.AuthorisationResult = new(AuthorisationResult9)
	return a.AuthorisationResult
}

func (a *ATMTransaction3) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}

func (a *ATMTransaction3) AddATMTotals() *ATMTotals1 {
	newValue := new(ATMTotals1)
	a.ATMTotals = append(a.ATMTotals, newValue)
	return newValue
}

func (a *ATMTransaction3) AddCassette() *ATMCassette1 {
	newValue := new(ATMCassette1)
	a.Cassette = append(a.Cassette, newValue)
	return newValue
}

// Acknowledgement of a completion advice.
type ATMTransaction4 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Response to the withdrawal advice.
	Response *Response4Code `xml:"Rspn"`

	// Detail of the response.
	ResponseReason *ResultDetail2Code `xml:"RspnRsn,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`

	// Maintenance command to perform on the ATM.
	Command []*ATMCommand1 `xml:"Cmd,omitempty"`
}

func (a *ATMTransaction4) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction4) SetResponse(value string) {
	a.Response = (*Response4Code)(&value)
}

func (a *ATMTransaction4) SetResponseReason(value string) {
	a.ResponseReason = (*ResultDetail2Code)(&value)
}

func (a *ATMTransaction4) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}

func (a *ATMTransaction4) AddCommand() *ATMCommand1 {
	newValue := new(ATMCommand1)
	a.Command = append(a.Command, newValue)
	return newValue
}

// Withdrawal transaction for which the completion is sent.
type ATMTransaction5 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Outcome of the financial transaction for the customer.
	TransactionStatus *ATMTransactionStatus1Code `xml:"TxSts"`

	// Incident occurring during the processing of the transaction.
	Incident []*FailureReason4Code `xml:"Incdnt,omitempty"`

	// Explanation of the incident.
	IncidentDetail []*Max70Text `xml:"IncdntDtl,omitempty"`

	// Identification of the reconciliation period assigned by the ATM.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// True if the customer has requested a receipt.
	RequestedReceipt *TrueFalseIndicator `xml:"ReqdRct,omitempty"`

	// True if a receipt has been printed and presented to the customer.
	ReceiptPrinted *TrueFalseIndicator `xml:"RctPrtd,omitempty"`

	// True when the card was captured by the ATM.
	CapturedCard *TrueFalseIndicator `xml:"CaptrdCard,omitempty"`

	// Outcome of the withdrawal authorisation.
	AuthorisationResult *AuthorisationResult9 `xml:"AuthstnRslt,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (a *ATMTransaction5) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction5) SetTransactionStatus(value string) {
	a.TransactionStatus = (*ATMTransactionStatus1Code)(&value)
}

func (a *ATMTransaction5) AddIncident(value string) {
	a.Incident = append(a.Incident, (*FailureReason4Code)(&value))
}

func (a *ATMTransaction5) AddIncidentDetail(value string) {
	a.IncidentDetail = append(a.IncidentDetail, (*Max70Text)(&value))
}

func (a *ATMTransaction5) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction5) SetRequestedReceipt(value string) {
	a.RequestedReceipt = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction5) SetReceiptPrinted(value string) {
	a.ReceiptPrinted = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction5) SetCapturedCard(value string) {
	a.CapturedCard = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction5) AddAuthorisationResult() *AuthorisationResult9 {
	a.AuthorisationResult = new(AuthorisationResult9)
	return a.AuthorisationResult
}

func (a *ATMTransaction5) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}

// Inquiry information for the transaction.
type ATMTransaction6 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Unprotected account information.
	AccountData *CardAccount3 `xml:"AcctData,omitempty"`

	// Encryption of account information.
	ProtectedAccountData *ContentInformationType10 `xml:"PrtctdAcctData,omitempty"`

	// Amount to be authorised by the issuer.
	TotalRequestedAmount *AmountAndCurrency1 `xml:"TtlReqdAmt,omitempty"`

	// Amounts of the withdrawal transaction.
	DetailedRequestedAmount *DetailedAmount12 `xml:"DtldReqdAmt,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (a *ATMTransaction6) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction6) AddAccountData() *CardAccount3 {
	a.AccountData = new(CardAccount3)
	return a.AccountData
}

func (a *ATMTransaction6) AddProtectedAccountData() *ContentInformationType10 {
	a.ProtectedAccountData = new(ContentInformationType10)
	return a.ProtectedAccountData
}

func (a *ATMTransaction6) AddTotalRequestedAmount() *AmountAndCurrency1 {
	a.TotalRequestedAmount = new(AmountAndCurrency1)
	return a.TotalRequestedAmount
}

func (a *ATMTransaction6) AddDetailedRequestedAmount() *DetailedAmount12 {
	a.DetailedRequestedAmount = new(DetailedAmount12)
	return a.DetailedRequestedAmount
}

func (a *ATMTransaction6) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}

// Inquiry information for the transaction.
type ATMTransaction7 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Result of the inquiry service.
	TransactionResponse *ResponseType3 `xml:"TxRspn"`

	// Sequence of actions to be performed by the ATM to complete the transaction.
	Action []*Action5 `xml:"Actn,omitempty"`

	// Profile of the customer with the allowed services and restrictions.
	CustomerServiceProfile *ATMCustomerProfile3 `xml:"CstmrSvcPrfl,omitempty"`

	// Dynamic currency conversion result.
	CurrencyConversion *CurrencyConversion3 `xml:"CcyConvs,omitempty"`

	// Account information associated to the customer.
	AccountInformation []*CardAccount6 `xml:"AcctInf,omitempty"`

	// Statement information of an account.
	AccountStatementData []*ATMAccountStatement1 `xml:"AcctStmtData,omitempty"`

	// Exchange rate and calculated amount to be presented to the customer when the dispense currency or the deposit currency (target currency) is different to account currency (source currency).
	CurrencyExchange *CurrencyConversion5 `xml:"CcyXchg,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`

	// Maintenance command to perform on the ATM.
	Command []*ATMCommand1 `xml:"Cmd,omitempty"`
}

func (a *ATMTransaction7) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction7) AddTransactionResponse() *ResponseType3 {
	a.TransactionResponse = new(ResponseType3)
	return a.TransactionResponse
}

func (a *ATMTransaction7) AddAction() *Action5 {
	newValue := new(Action5)
	a.Action = append(a.Action, newValue)
	return newValue
}

func (a *ATMTransaction7) AddCustomerServiceProfile() *ATMCustomerProfile3 {
	a.CustomerServiceProfile = new(ATMCustomerProfile3)
	return a.CustomerServiceProfile
}

func (a *ATMTransaction7) AddCurrencyConversion() *CurrencyConversion3 {
	a.CurrencyConversion = new(CurrencyConversion3)
	return a.CurrencyConversion
}

func (a *ATMTransaction7) AddAccountInformation() *CardAccount6 {
	newValue := new(CardAccount6)
	a.AccountInformation = append(a.AccountInformation, newValue)
	return newValue
}

func (a *ATMTransaction7) AddAccountStatementData() *ATMAccountStatement1 {
	newValue := new(ATMAccountStatement1)
	a.AccountStatementData = append(a.AccountStatementData, newValue)
	return newValue
}

func (a *ATMTransaction7) AddCurrencyExchange() *CurrencyConversion5 {
	a.CurrencyExchange = new(CurrencyConversion5)
	return a.CurrencyExchange
}

func (a *ATMTransaction7) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}

func (a *ATMTransaction7) AddCommand() *ATMCommand1 {
	newValue := new(ATMCommand1)
	a.Command = append(a.Command, newValue)
	return newValue
}

// Preferred withdrawal transaction chosen by the the customer.
type ATMTransaction8 struct {

	// Amount to dispense.
	Amount *ImpliedCurrencyAndAmount `xml:"Amt,omitempty"`

	// Currency.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// True if a receipt has to be printed by the ATM for the customer.
	ReceiptFlag *TrueFalseIndicator `xml:"RctFlg,omitempty"`

	// True if a balance has to be printed by the ATM on the customer receipt.
	BalancePrintFlag *TrueFalseIndicator `xml:"BalPrtFlg,omitempty"`

	// Media mix algorithm, the identification of the algorithm is an agreement between the ATM and the ATM manager.
	MixType *Max35Text `xml:"MixTp,omitempty"`

	// Media mix to select.
	Mix []*ATMMediaMix2 `xml:"Mix,omitempty"`
}

func (a *ATMTransaction8) SetAmount(value, currency string) {
	a.Amount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransaction8) SetCurrency(value string) {
	a.Currency = (*ActiveCurrencyCode)(&value)
}

func (a *ATMTransaction8) SetReceiptFlag(value string) {
	a.ReceiptFlag = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction8) SetBalancePrintFlag(value string) {
	a.BalancePrintFlag = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction8) SetMixType(value string) {
	a.MixType = (*Max35Text)(&value)
}

func (a *ATMTransaction8) AddMix() *ATMMediaMix2 {
	newValue := new(ATMMediaMix2)
	a.Mix = append(a.Mix, newValue)
	return newValue
}

// Transaction for which the service is requested.
type ATMTransaction9 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the reconciliation period assigned by the ATM manager.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Encrypted personal identification number (PIN) and related information.
	CardholderNewPIN *OnLinePIN5 `xml:"CrdhldrNewPIN,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (a *ATMTransaction9) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction9) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction9) AddCardholderNewPIN() *OnLinePIN5 {
	a.CardholderNewPIN = new(OnLinePIN5)
	return a.CardholderNewPIN
}

func (a *ATMTransaction9) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}
