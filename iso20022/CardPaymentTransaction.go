package iso20022

// Data associated with the transaction during the authorisation.
type CardPaymentTransaction1 struct {

	// Flag indicating whether the transaction data must be captured or not in addition to the message process.
	TransactionCapture *TrueFalseIndicator `xml:"TxCaptr"`

	// Type of transaction being undertaken for the main service.
	TransactionType *CardPaymentServiceType1Code `xml:"TxTp"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType2Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd"`

	// Identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction8 `xml:"OrgnlTx,omitempty"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails1 `xml:"TxDtls"`

	// Additional information related to the transaction.
	AdditionalTransactionData *Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction1) SetTransactionCapture(value string) {
	c.TransactionCapture = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction1) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType1Code)(&value)
}

func (c *CardPaymentTransaction1) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType2Code)(&value))
}

func (c *CardPaymentTransaction1) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CardPaymentTransaction1) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction1) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction1) AddOriginalTransaction() *CardPaymentTransaction8 {
	c.OriginalTransaction = new(CardPaymentTransaction8)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction1) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction1) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction1) AddTransactionDetails() *CardPaymentTransactionDetails1 {
	c.TransactionDetails = new(CardPaymentTransactionDetails1)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction1) SetAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = (*Max70Text)(&value)
}

// Cancellation response from the acquirer.
type CardPaymentTransaction10 struct {

	// Outcome of the authorisation, and actions to perform.
	AuthorisationResult *AuthorisationResult1 `xml:"AuthstnRslt"`

	// Set of actions to be performed by the POI (Point Of Interaction) system.
	Action []*Action1 `xml:"Actn,omitempty"`
}

func (c *CardPaymentTransaction10) AddAuthorisationResult() *AuthorisationResult1 {
	c.AuthorisationResult = new(AuthorisationResult1)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction10) AddAction() *Action1 {
	newValue := new(Action1)
	c.Action = append(c.Action, newValue)
	return newValue
}

// Data associated with the transaction during the authorisation.
type CardPaymentTransaction11 struct {

	// Flag indicating whether the transaction data must be captured or not in addition to the message process.
	TransactionCapture *TrueFalseIndicator `xml:"TxCaptr"`

	// Type of transaction being undertaken for the main service.
	TransactionType *CardPaymentServiceType4Code `xml:"TxTp"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType2Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd"`

	// Identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction17 `xml:"OrgnlTx,omitempty"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails1 `xml:"TxDtls"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction11) SetTransactionCapture(value string) {
	c.TransactionCapture = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction11) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType4Code)(&value)
}

func (c *CardPaymentTransaction11) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType2Code)(&value))
}

func (c *CardPaymentTransaction11) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CardPaymentTransaction11) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction11) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction11) AddOriginalTransaction() *CardPaymentTransaction17 {
	c.OriginalTransaction = new(CardPaymentTransaction17)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction11) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction11) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction11) AddTransactionDetails() *CardPaymentTransactionDetails1 {
	c.TransactionDetails = new(CardPaymentTransactionDetails1)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction11) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}

// Transaction information in the authorisation response.
type CardPaymentTransaction12 struct {

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max35Text `xml:"IntrchngData,omitempty"`

	// Detail of the transaction transported.
	TransactionDetails *CardPaymentTransactionDetails11 `xml:"TxDtls"`
}

func (c *CardPaymentTransaction12) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction12) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction12) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction12) SetInterchangeData(value string) {
	c.InterchangeData = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction12) AddTransactionDetails() *CardPaymentTransactionDetails11 {
	c.TransactionDetails = new(CardPaymentTransactionDetails11)
	return c.TransactionDetails
}

// Transaction information in the completion advice message.
type CardPaymentTransaction13 struct {

	// Flag indicating whether the transaction data must be captured or not in addition to the message process.
	TransactionCapture *TrueFalseIndicator `xml:"TxCaptr,omitempty"`

	// Type of transaction being undertaken for the main service.
	TransactionType *CardPaymentServiceType4Code `xml:"TxTp"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType2Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction17 `xml:"OrgnlTx,omitempty"`

	// Outcome of the transaction at the acceptor.
	TransactionSuccess *TrueFalseIndicator `xml:"TxSucss"`

	// Notify that a previous transaction has to be reversed if this original transaction has been appoved by the acquirer.
	Reversal *TrueFalseIndicator `xml:"Rvsl,omitempty"`

	// Indicate that the acceptor has forced the transaction in spite of the authorisation result (online or offline), or incident to complete the transaction.
	MerchantOverride *TrueFalseIndicator `xml:"MrchntOvrrd,omitempty"`

	// List of incidents during the transaction.
	FailureReason []*FailureReason2Code `xml:"FailrRsn,omitempty"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max35Text `xml:"IntrchngData,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails3 `xml:"TxDtls"`

	// Outcome of the authorisation request.
	AuthorisationResult *AuthorisationResult2 `xml:"AuthstnRslt,omitempty"`

	// Result of the verifications performed by the issuer to deliver or decline the authorisation.
	TransactionVerificationResult *TransactionVerificationResult2 `xml:"TxVrfctnRslt,omitempty"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction13) SetTransactionCapture(value string) {
	c.TransactionCapture = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction13) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType4Code)(&value)
}

func (c *CardPaymentTransaction13) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType2Code)(&value))
}

func (c *CardPaymentTransaction13) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CardPaymentTransaction13) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction13) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction13) AddOriginalTransaction() *CardPaymentTransaction17 {
	c.OriginalTransaction = new(CardPaymentTransaction17)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction13) SetTransactionSuccess(value string) {
	c.TransactionSuccess = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction13) SetReversal(value string) {
	c.Reversal = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction13) SetMerchantOverride(value string) {
	c.MerchantOverride = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction13) AddFailureReason(value string) {
	c.FailureReason = append(c.FailureReason, (*FailureReason2Code)(&value))
}

func (c *CardPaymentTransaction13) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction13) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction13) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction13) SetInterchangeData(value string) {
	c.InterchangeData = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction13) AddTransactionDetails() *CardPaymentTransactionDetails3 {
	c.TransactionDetails = new(CardPaymentTransactionDetails3)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction13) AddAuthorisationResult() *AuthorisationResult2 {
	c.AuthorisationResult = new(AuthorisationResult2)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction13) AddTransactionVerificationResult() *TransactionVerificationResult2 {
	c.TransactionVerificationResult = new(TransactionVerificationResult2)
	return c.TransactionVerificationResult
}

func (c *CardPaymentTransaction13) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}

// Transaction information in the batch capture.
type CardPaymentTransaction14 struct {

	// Type of transaction being undertaken for the main service.
	TransactionType *CardPaymentServiceType4Code `xml:"TxTp,omitempty"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType2Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction17 `xml:"OrgnlTx,omitempty"`

	// Outcome of the transaction at the acceptor.
	TransactionSuccess *TrueFalseIndicator `xml:"TxSucss"`

	// Notify that a previous transaction has to be reversed if this original transaction has been approved by the acquirer.
	Reversal *TrueFalseIndicator `xml:"Rvsl,omitempty"`

	// Indicate that the acceptor has forced the transaction in spite of the authorisation result (online or offline), or incident to complete the transaction.
	MerchantOverride *TrueFalseIndicator `xml:"MrchntOvrrd,omitempty"`

	// List of incidents during the transaction.
	FailureReason []*FailureReason2Code `xml:"FailrRsn,omitempty"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max35Text `xml:"IntrchngData,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails4 `xml:"TxDtls"`

	// Outcome of the authorisation request, and actions to perform.
	AuthorisationResult *AuthorisationResult1 `xml:"AuthstnRslt,omitempty"`

	// Result of the verifications performed by the issuer to deliver or decline the authorisation.
	TransactionVerificationResult *TransactionVerificationResult2 `xml:"TxVrfctnRslt,omitempty"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction14) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType4Code)(&value)
}

func (c *CardPaymentTransaction14) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType2Code)(&value))
}

func (c *CardPaymentTransaction14) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CardPaymentTransaction14) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction14) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction14) AddOriginalTransaction() *CardPaymentTransaction17 {
	c.OriginalTransaction = new(CardPaymentTransaction17)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction14) SetTransactionSuccess(value string) {
	c.TransactionSuccess = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction14) SetReversal(value string) {
	c.Reversal = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction14) SetMerchantOverride(value string) {
	c.MerchantOverride = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction14) AddFailureReason(value string) {
	c.FailureReason = append(c.FailureReason, (*FailureReason2Code)(&value))
}

func (c *CardPaymentTransaction14) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction14) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction14) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction14) SetInterchangeData(value string) {
	c.InterchangeData = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction14) AddTransactionDetails() *CardPaymentTransactionDetails4 {
	c.TransactionDetails = new(CardPaymentTransactionDetails4)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction14) AddAuthorisationResult() *AuthorisationResult1 {
	c.AuthorisationResult = new(AuthorisationResult1)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction14) AddTransactionVerificationResult() *TransactionVerificationResult2 {
	c.TransactionVerificationResult = new(TransactionVerificationResult2)
	return c.TransactionVerificationResult
}

func (c *CardPaymentTransaction14) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}

// Transaction information in the cancellation request.
type CardPaymentTransaction15 struct {

	// Flag indicating whether the transaction data must be captured or not in addition to the message process.
	TransactionCapture *TrueFalseIndicator `xml:"TxCaptr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction17 `xml:"OrgnlTx"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails5 `xml:"TxDtls"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction15) SetTransactionCapture(value string) {
	c.TransactionCapture = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction15) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction15) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction15) AddOriginalTransaction() *CardPaymentTransaction17 {
	c.OriginalTransaction = new(CardPaymentTransaction17)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction15) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction15) AddTransactionDetails() *CardPaymentTransactionDetails5 {
	c.TransactionDetails = new(CardPaymentTransactionDetails5)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction15) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}

// Transaction information in the cancellation advice.
type CardPaymentTransaction16 struct {

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction17 `xml:"OrgnlTx,omitempty"`

	// Outcome of the transaction at the acceptor.
	TransactionSuccess *TrueFalseIndicator `xml:"TxSucss"`

	// Notify that a previous transaction has to be reversed if this original transaction has been approved by the acquirer.
	Reversal *TrueFalseIndicator `xml:"Rvsl,omitempty"`

	// List of incidents during the transaction.
	FailureReason []*FailureReason2Code `xml:"FailrRsn,omitempty"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max35Text `xml:"IntrchngData,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails7 `xml:"TxDtls"`

	// Outcome of the authorisation.
	AuthorisationResult *AuthorisationResult2 `xml:"AuthstnRslt,omitempty"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction16) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction16) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction16) AddOriginalTransaction() *CardPaymentTransaction17 {
	c.OriginalTransaction = new(CardPaymentTransaction17)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction16) SetTransactionSuccess(value string) {
	c.TransactionSuccess = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction16) SetReversal(value string) {
	c.Reversal = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction16) AddFailureReason(value string) {
	c.FailureReason = append(c.FailureReason, (*FailureReason2Code)(&value))
}

func (c *CardPaymentTransaction16) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction16) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction16) SetInterchangeData(value string) {
	c.InterchangeData = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction16) AddTransactionDetails() *CardPaymentTransactionDetails7 {
	c.TransactionDetails = new(CardPaymentTransactionDetails7)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction16) AddAuthorisationResult() *AuthorisationResult2 {
	c.AuthorisationResult = new(AuthorisationResult2)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction16) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}

// Identification of the original transaction.
type CardPaymentTransaction17 struct {

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the POI (Point Of Interaction) performing the transaction.
	POIIdentification *GenericIdentification32 `xml:"POIId,omitempty"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Type of transaction being undertaken for the main service.
	TransactionType *CardPaymentServiceType4Code `xml:"TxTp"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType2Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Result of the original transaction.
	TransactionResult *CardPaymentTransactionResult1 `xml:"TxRslt,omitempty"`
}

func (c *CardPaymentTransaction17) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction17) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

func (c *CardPaymentTransaction17) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction17) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction17) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType4Code)(&value)
}

func (c *CardPaymentTransaction17) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType2Code)(&value))
}

func (c *CardPaymentTransaction17) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CardPaymentTransaction17) AddTransactionResult() *CardPaymentTransactionResult1 {
	c.TransactionResult = new(CardPaymentTransactionResult1)
	return c.TransactionResult
}

// Authorisation response from the acquirer.
type CardPaymentTransaction18 struct {

	// Outcome of the authorisation, and actions to perform.
	AuthorisationResult *AuthorisationResult1 `xml:"AuthstnRslt"`

	// Result of the verifications performed by the issuer to deliver or decline the authorisation.
	TransactionVerificationResult *TransactionVerificationResult2 `xml:"TxVrfctnRslt,omitempty"`

	// Balance of the account, related to the payment.
	Balance *ImpliedCurrencyAndAmount `xml:"Bal,omitempty"`

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Set of actions to be performed by the POI (Point Of Interaction) system.
	Action []*Action2 `xml:"Actn,omitempty"`
}

func (c *CardPaymentTransaction18) AddAuthorisationResult() *AuthorisationResult1 {
	c.AuthorisationResult = new(AuthorisationResult1)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction18) AddTransactionVerificationResult() *TransactionVerificationResult2 {
	c.TransactionVerificationResult = new(TransactionVerificationResult2)
	return c.TransactionVerificationResult
}

func (c *CardPaymentTransaction18) SetBalance(value, currency string) {
	c.Balance = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransaction18) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransaction18) AddAction() *Action2 {
	newValue := new(Action2)
	c.Action = append(c.Action, newValue)
	return newValue
}

// Data associated with the transaction to authorise.
type CardPaymentTransaction19 struct {

	// Flag indicating whether the transaction data must be captured or not in addition to the message process.
	TransactionCapture *TrueFalseIndicator `xml:"TxCaptr"`

	// Type of transaction being undertaken for the main service.
	TransactionType *CardPaymentServiceType4Code `xml:"TxTp,omitempty"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType2Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd,omitempty"`

	// Identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction17 `xml:"OrgnlTx,omitempty"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails10 `xml:"TxDtls"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction19) SetTransactionCapture(value string) {
	c.TransactionCapture = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction19) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType4Code)(&value)
}

func (c *CardPaymentTransaction19) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType2Code)(&value))
}

func (c *CardPaymentTransaction19) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CardPaymentTransaction19) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction19) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction19) AddOriginalTransaction() *CardPaymentTransaction17 {
	c.OriginalTransaction = new(CardPaymentTransaction17)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction19) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction19) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction19) AddTransactionDetails() *CardPaymentTransactionDetails10 {
	c.TransactionDetails = new(CardPaymentTransactionDetails10)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction19) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}

// Transaction information in the authorisation response.
type CardPaymentTransaction2 struct {

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max35Text `xml:"IntrchngData,omitempty"`

	// Detail of the transaction transported.
	TransactionDetails *CardPaymentTransactionDetails2 `xml:"TxDtls"`
}

func (c *CardPaymentTransaction2) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction2) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction2) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction2) SetInterchangeData(value string) {
	c.InterchangeData = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction2) AddTransactionDetails() *CardPaymentTransactionDetails2 {
	c.TransactionDetails = new(CardPaymentTransactionDetails2)
	return c.TransactionDetails
}

// Transaction information in the cancellation advice.
type CardPaymentTransaction20 struct {

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction17 `xml:"OrgnlTx,omitempty"`

	// Outcome of the transaction at the acceptor.
	TransactionSuccess *TrueFalseIndicator `xml:"TxSucss"`

	// Notify that a previous transaction has to be reversed if this original transaction has been approved by the acquirer.
	Reversal *TrueFalseIndicator `xml:"Rvsl,omitempty"`

	// List of incidents during the transaction.
	FailureReason []*FailureReason2Code `xml:"FailrRsn,omitempty"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max35Text `xml:"IntrchngData,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails9 `xml:"TxDtls"`

	// Outcome of the authorisation.
	AuthorisationResult *AuthorisationResult3 `xml:"AuthstnRslt,omitempty"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction20) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction20) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction20) AddOriginalTransaction() *CardPaymentTransaction17 {
	c.OriginalTransaction = new(CardPaymentTransaction17)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction20) SetTransactionSuccess(value string) {
	c.TransactionSuccess = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction20) SetReversal(value string) {
	c.Reversal = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction20) AddFailureReason(value string) {
	c.FailureReason = append(c.FailureReason, (*FailureReason2Code)(&value))
}

func (c *CardPaymentTransaction20) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction20) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction20) SetInterchangeData(value string) {
	c.InterchangeData = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction20) AddTransactionDetails() *CardPaymentTransactionDetails9 {
	c.TransactionDetails = new(CardPaymentTransactionDetails9)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction20) AddAuthorisationResult() *AuthorisationResult3 {
	c.AuthorisationResult = new(AuthorisationResult3)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction20) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}

// Identification of the original transaction.
type CardPaymentTransaction21 struct {

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the POI (Point Of Interaction) performing the transaction.
	POIIdentification *GenericIdentification32 `xml:"POIId,omitempty"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Type of transaction being undertaken for the main service.
	TransactionType *CardPaymentServiceType5Code `xml:"TxTp"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType6Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Result of the original transaction.
	TransactionResult *CardPaymentTransactionResult1 `xml:"TxRslt,omitempty"`
}

func (c *CardPaymentTransaction21) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction21) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction21) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

func (c *CardPaymentTransaction21) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction21) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction21) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType5Code)(&value)
}

func (c *CardPaymentTransaction21) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType6Code)(&value))
}

func (c *CardPaymentTransaction21) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CardPaymentTransaction21) AddTransactionResult() *CardPaymentTransactionResult1 {
	c.TransactionResult = new(CardPaymentTransactionResult1)
	return c.TransactionResult
}

// Data associated with the transaction during the authorisation.
type CardPaymentTransaction22 struct {

	// Flag indicating whether the transaction data must be captured or not in addition to the message process.
	TransactionCapture *TrueFalseIndicator `xml:"TxCaptr"`

	// Type of transaction being undertaken for the main service.
	TransactionType *CardPaymentServiceType5Code `xml:"TxTp"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType6Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd"`

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction21 `xml:"OrgnlTx,omitempty"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails12 `xml:"TxDtls"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction22) SetTransactionCapture(value string) {
	c.TransactionCapture = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction22) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType5Code)(&value)
}

func (c *CardPaymentTransaction22) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType6Code)(&value))
}

func (c *CardPaymentTransaction22) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CardPaymentTransaction22) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction22) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction22) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction22) AddOriginalTransaction() *CardPaymentTransaction21 {
	c.OriginalTransaction = new(CardPaymentTransaction21)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction22) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction22) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction22) AddTransactionDetails() *CardPaymentTransactionDetails12 {
	c.TransactionDetails = new(CardPaymentTransactionDetails12)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction22) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}

// Transaction information in the authorisation response.
type CardPaymentTransaction23 struct {

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max35Text `xml:"IntrchngData,omitempty"`

	// Detail of the transaction transported.
	TransactionDetails *CardPaymentTransactionDetails13 `xml:"TxDtls"`
}

func (c *CardPaymentTransaction23) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction23) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction23) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction23) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction23) SetInterchangeData(value string) {
	c.InterchangeData = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction23) AddTransactionDetails() *CardPaymentTransactionDetails13 {
	c.TransactionDetails = new(CardPaymentTransactionDetails13)
	return c.TransactionDetails
}

// Authorisation response from the acquirer.
type CardPaymentTransaction24 struct {

	// Outcome of the authorisation, and actions to perform.
	AuthorisationResult *AuthorisationResult1 `xml:"AuthstnRslt"`

	// Result of the verifications performed by the issuer to deliver or decline the authorisation.
	TransactionVerificationResult *TransactionVerificationResult2 `xml:"TxVrfctnRslt,omitempty"`

	// Balance of the account, related to the payment.
	Balance *AmountAndDirection41 `xml:"Bal,omitempty"`

	// Set of actions to be performed by the POI (Point Of Interaction) system.
	Action []*Action3 `xml:"Actn,omitempty"`

	// Conversion between the currency of a card acceptor and the currency of a card issuer, provided by a dedicated service provider. The currency conversion has to be accepted by the cardholder.
	CurrencyConversion *CurrencyConversion1 `xml:"CcyConvs,omitempty"`
}

func (c *CardPaymentTransaction24) AddAuthorisationResult() *AuthorisationResult1 {
	c.AuthorisationResult = new(AuthorisationResult1)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction24) AddTransactionVerificationResult() *TransactionVerificationResult2 {
	c.TransactionVerificationResult = new(TransactionVerificationResult2)
	return c.TransactionVerificationResult
}

func (c *CardPaymentTransaction24) AddBalance() *AmountAndDirection41 {
	c.Balance = new(AmountAndDirection41)
	return c.Balance
}

func (c *CardPaymentTransaction24) AddAction() *Action3 {
	newValue := new(Action3)
	c.Action = append(c.Action, newValue)
	return newValue
}

func (c *CardPaymentTransaction24) AddCurrencyConversion() *CurrencyConversion1 {
	c.CurrencyConversion = new(CurrencyConversion1)
	return c.CurrencyConversion
}

// Transaction information in the completion advice message.
type CardPaymentTransaction25 struct {

	// Flag indicating whether the transaction data must be captured or not in addition to the message process.
	TransactionCapture *TrueFalseIndicator `xml:"TxCaptr,omitempty"`

	// Type of transaction being undertaken for the main service.
	TransactionType *CardPaymentServiceType5Code `xml:"TxTp"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType6Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd"`

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction21 `xml:"OrgnlTx,omitempty"`

	// Outcome of the transaction at the acceptor.
	TransactionSuccess *TrueFalseIndicator `xml:"TxSucss"`

	// Notify that a previous transaction has to be reversed if this original transaction has been appoved by the acquirer.
	Reversal *TrueFalseIndicator `xml:"Rvsl,omitempty"`

	// Indicate that the acceptor has forced the transaction in spite of the authorisation result (online or offline), or incident to complete the transaction.
	MerchantOverride *TrueFalseIndicator `xml:"MrchntOvrrd,omitempty"`

	// List of incidents during the transaction.
	FailureReason []*FailureReason3Code `xml:"FailrRsn,omitempty"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max35Text `xml:"IntrchngData,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails14 `xml:"TxDtls"`

	// Outcome of the authorisation request.
	AuthorisationResult *AuthorisationResult2 `xml:"AuthstnRslt,omitempty"`

	// Result of the verifications performed by the issuer to deliver or decline the authorisation.
	TransactionVerificationResult *TransactionVerificationResult2 `xml:"TxVrfctnRslt,omitempty"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction25) SetTransactionCapture(value string) {
	c.TransactionCapture = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction25) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType5Code)(&value)
}

func (c *CardPaymentTransaction25) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType6Code)(&value))
}

func (c *CardPaymentTransaction25) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CardPaymentTransaction25) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction25) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction25) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction25) AddOriginalTransaction() *CardPaymentTransaction21 {
	c.OriginalTransaction = new(CardPaymentTransaction21)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction25) SetTransactionSuccess(value string) {
	c.TransactionSuccess = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction25) SetReversal(value string) {
	c.Reversal = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction25) SetMerchantOverride(value string) {
	c.MerchantOverride = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction25) AddFailureReason(value string) {
	c.FailureReason = append(c.FailureReason, (*FailureReason3Code)(&value))
}

func (c *CardPaymentTransaction25) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction25) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction25) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction25) SetInterchangeData(value string) {
	c.InterchangeData = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction25) AddTransactionDetails() *CardPaymentTransactionDetails14 {
	c.TransactionDetails = new(CardPaymentTransactionDetails14)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction25) AddAuthorisationResult() *AuthorisationResult2 {
	c.AuthorisationResult = new(AuthorisationResult2)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction25) AddTransactionVerificationResult() *TransactionVerificationResult2 {
	c.TransactionVerificationResult = new(TransactionVerificationResult2)
	return c.TransactionVerificationResult
}

func (c *CardPaymentTransaction25) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}

// Transaction information in the cancellation request.
type CardPaymentTransaction26 struct {

	// Flag indicating whether the transaction data must be captured or not in addition to the message process.
	TransactionCapture *TrueFalseIndicator `xml:"TxCaptr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd"`

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction21 `xml:"OrgnlTx"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails5 `xml:"TxDtls"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction26) SetTransactionCapture(value string) {
	c.TransactionCapture = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction26) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction26) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction26) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction26) AddOriginalTransaction() *CardPaymentTransaction21 {
	c.OriginalTransaction = new(CardPaymentTransaction21)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction26) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction26) AddTransactionDetails() *CardPaymentTransactionDetails5 {
	c.TransactionDetails = new(CardPaymentTransactionDetails5)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction26) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}

// Cancellation response from the acquirer.
type CardPaymentTransaction27 struct {

	// Outcome of the authorisation, and actions to perform.
	AuthorisationResult *AuthorisationResult3 `xml:"AuthstnRslt"`

	// Set of actions to be performed by the POI (Point Of Interaction) system.
	Action []*Action1 `xml:"Actn,omitempty"`
}

func (c *CardPaymentTransaction27) AddAuthorisationResult() *AuthorisationResult3 {
	c.AuthorisationResult = new(AuthorisationResult3)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction27) AddAction() *Action1 {
	newValue := new(Action1)
	c.Action = append(c.Action, newValue)
	return newValue
}

// Transaction information in the cancellation advice.
type CardPaymentTransaction28 struct {

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd"`

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction21 `xml:"OrgnlTx,omitempty"`

	// Outcome of the transaction at the acceptor.
	TransactionSuccess *TrueFalseIndicator `xml:"TxSucss"`

	// Notify that a previous transaction has to be reversed if this original transaction has been approved by the acquirer.
	Reversal *TrueFalseIndicator `xml:"Rvsl,omitempty"`

	// List of incidents during the transaction.
	FailureReason []*FailureReason3Code `xml:"FailrRsn,omitempty"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max35Text `xml:"IntrchngData,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails7 `xml:"TxDtls"`

	// Outcome of the authorisation.
	AuthorisationResult *AuthorisationResult2 `xml:"AuthstnRslt,omitempty"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction28) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction28) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction28) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction28) AddOriginalTransaction() *CardPaymentTransaction21 {
	c.OriginalTransaction = new(CardPaymentTransaction21)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction28) SetTransactionSuccess(value string) {
	c.TransactionSuccess = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction28) SetReversal(value string) {
	c.Reversal = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction28) AddFailureReason(value string) {
	c.FailureReason = append(c.FailureReason, (*FailureReason3Code)(&value))
}

func (c *CardPaymentTransaction28) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction28) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction28) SetInterchangeData(value string) {
	c.InterchangeData = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction28) AddTransactionDetails() *CardPaymentTransactionDetails7 {
	c.TransactionDetails = new(CardPaymentTransactionDetails7)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction28) AddAuthorisationResult() *AuthorisationResult2 {
	c.AuthorisationResult = new(AuthorisationResult2)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction28) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}

// Transaction information in the batch capture.
type CardPaymentTransaction29 struct {

	// Type of transaction being undertaken for the main service.
	TransactionType *CardPaymentServiceType5Code `xml:"TxTp,omitempty"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType6Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd,omitempty"`

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction21 `xml:"OrgnlTx,omitempty"`

	// Outcome of the transaction at the acceptor.
	TransactionSuccess *TrueFalseIndicator `xml:"TxSucss"`

	// Notify that a previous transaction has to be reversed if this original transaction has been approved by the acquirer.
	Reversal *TrueFalseIndicator `xml:"Rvsl,omitempty"`

	// Indicate that the acceptor has forced the transaction in spite of the authorisation result (online or offline), or incident to complete the transaction.
	MerchantOverride *TrueFalseIndicator `xml:"MrchntOvrrd,omitempty"`

	// List of incidents during the transaction.
	FailureReason []*FailureReason3Code `xml:"FailrRsn,omitempty"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max35Text `xml:"IntrchngData,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails15 `xml:"TxDtls"`

	// Outcome of the authorisation request, and actions to perform.
	AuthorisationResult *AuthorisationResult1 `xml:"AuthstnRslt,omitempty"`

	// Result of the verifications performed by the issuer to deliver or decline the authorisation.
	TransactionVerificationResult *TransactionVerificationResult2 `xml:"TxVrfctnRslt,omitempty"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction29) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType5Code)(&value)
}

func (c *CardPaymentTransaction29) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType6Code)(&value))
}

func (c *CardPaymentTransaction29) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CardPaymentTransaction29) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction29) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction29) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction29) AddOriginalTransaction() *CardPaymentTransaction21 {
	c.OriginalTransaction = new(CardPaymentTransaction21)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction29) SetTransactionSuccess(value string) {
	c.TransactionSuccess = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction29) SetReversal(value string) {
	c.Reversal = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction29) SetMerchantOverride(value string) {
	c.MerchantOverride = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction29) AddFailureReason(value string) {
	c.FailureReason = append(c.FailureReason, (*FailureReason3Code)(&value))
}

func (c *CardPaymentTransaction29) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction29) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction29) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction29) SetInterchangeData(value string) {
	c.InterchangeData = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction29) AddTransactionDetails() *CardPaymentTransactionDetails15 {
	c.TransactionDetails = new(CardPaymentTransactionDetails15)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction29) AddAuthorisationResult() *AuthorisationResult1 {
	c.AuthorisationResult = new(AuthorisationResult1)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction29) AddTransactionVerificationResult() *TransactionVerificationResult2 {
	c.TransactionVerificationResult = new(TransactionVerificationResult2)
	return c.TransactionVerificationResult
}

func (c *CardPaymentTransaction29) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}

// Transaction information in the completion advice message.
type CardPaymentTransaction3 struct {

	// Flag indicating whether the transaction data must be captured or not in addition to the message process.
	TransactionCapture *TrueFalseIndicator `xml:"TxCaptr,omitempty"`

	// Type of transaction being undertaken for the main service.
	TransactionType *CardPaymentServiceType1Code `xml:"TxTp,omitempty"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType2Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction8 `xml:"OrgnlTx,omitempty"`

	// Outcome of the transaction at the acceptor.
	TransactionSuccess *TrueFalseIndicator `xml:"TxSucss"`

	// Notify that a previous transaction has to be reversed if this original transaction has been appoved by the acquirer.
	Reversal *TrueFalseIndicator `xml:"Rvsl,omitempty"`

	// Indicate that the acceptor has forced the transaction in spite of the authorisation result (online or offline), or incident to complete the transaction.
	MerchantOverride *TrueFalseIndicator `xml:"MrchntOvrrd,omitempty"`

	// List of incidents during the transaction.
	FailureReason []*FailureReason1Code `xml:"FailrRsn,omitempty"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max35Text `xml:"IntrchngData,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails3 `xml:"TxDtls"`

	// Outcome of the authorisation request.
	AuthorisationResult *AuthorisationResult2 `xml:"AuthstnRslt,omitempty"`

	// Result of the verifications performed by the issuer to deliver or decline the authorisation.
	TransactionVerificationResult *TransactionVerificationResult1 `xml:"TxVrfctnRslt,omitempty"`

	// Additional information related to the transaction.
	AdditionalTransactionData *Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction3) SetTransactionCapture(value string) {
	c.TransactionCapture = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction3) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType1Code)(&value)
}

func (c *CardPaymentTransaction3) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType2Code)(&value))
}

func (c *CardPaymentTransaction3) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CardPaymentTransaction3) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction3) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction3) AddOriginalTransaction() *CardPaymentTransaction8 {
	c.OriginalTransaction = new(CardPaymentTransaction8)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction3) SetTransactionSuccess(value string) {
	c.TransactionSuccess = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction3) SetReversal(value string) {
	c.Reversal = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction3) SetMerchantOverride(value string) {
	c.MerchantOverride = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction3) AddFailureReason(value string) {
	c.FailureReason = append(c.FailureReason, (*FailureReason1Code)(&value))
}

func (c *CardPaymentTransaction3) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction3) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction3) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction3) SetInterchangeData(value string) {
	c.InterchangeData = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction3) AddTransactionDetails() *CardPaymentTransactionDetails3 {
	c.TransactionDetails = new(CardPaymentTransactionDetails3)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction3) AddAuthorisationResult() *AuthorisationResult2 {
	c.AuthorisationResult = new(AuthorisationResult2)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction3) AddTransactionVerificationResult() *TransactionVerificationResult1 {
	c.TransactionVerificationResult = new(TransactionVerificationResult1)
	return c.TransactionVerificationResult
}

func (c *CardPaymentTransaction3) SetAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = (*Max70Text)(&value)
}

// Transaction information in the cancellation advice.
type CardPaymentTransaction30 struct {

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd,omitempty"`

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction21 `xml:"OrgnlTx,omitempty"`

	// Outcome of the transaction at the acceptor.
	TransactionSuccess *TrueFalseIndicator `xml:"TxSucss"`

	// Notify that a previous transaction has to be reversed if this original transaction has been approved by the acquirer.
	Reversal *TrueFalseIndicator `xml:"Rvsl,omitempty"`

	// List of incidents during the transaction.
	FailureReason []*FailureReason3Code `xml:"FailrRsn,omitempty"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max35Text `xml:"IntrchngData,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails9 `xml:"TxDtls"`

	// Outcome of the authorisation.
	AuthorisationResult *AuthorisationResult3 `xml:"AuthstnRslt,omitempty"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction30) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction30) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction30) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction30) AddOriginalTransaction() *CardPaymentTransaction21 {
	c.OriginalTransaction = new(CardPaymentTransaction21)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction30) SetTransactionSuccess(value string) {
	c.TransactionSuccess = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction30) SetReversal(value string) {
	c.Reversal = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction30) AddFailureReason(value string) {
	c.FailureReason = append(c.FailureReason, (*FailureReason3Code)(&value))
}

func (c *CardPaymentTransaction30) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction30) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction30) SetInterchangeData(value string) {
	c.InterchangeData = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction30) AddTransactionDetails() *CardPaymentTransactionDetails9 {
	c.TransactionDetails = new(CardPaymentTransactionDetails9)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction30) AddAuthorisationResult() *AuthorisationResult3 {
	c.AuthorisationResult = new(AuthorisationResult3)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction30) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}

// Data associated with the transaction to authorise.
type CardPaymentTransaction31 struct {

	// Flag indicating whether the transaction data must be captured or not in addition to the message process.
	TransactionCapture *TrueFalseIndicator `xml:"TxCaptr"`

	// Type of transaction being undertaken for the main service.
	TransactionType *CardPaymentServiceType5Code `xml:"TxTp,omitempty"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType6Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd,omitempty"`

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction21 `xml:"OrgnlTx,omitempty"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails16 `xml:"TxDtls"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction31) SetTransactionCapture(value string) {
	c.TransactionCapture = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction31) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType5Code)(&value)
}

func (c *CardPaymentTransaction31) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType6Code)(&value))
}

func (c *CardPaymentTransaction31) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CardPaymentTransaction31) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction31) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction31) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction31) AddOriginalTransaction() *CardPaymentTransaction21 {
	c.OriginalTransaction = new(CardPaymentTransaction21)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction31) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction31) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction31) AddTransactionDetails() *CardPaymentTransactionDetails16 {
	c.TransactionDetails = new(CardPaymentTransactionDetails16)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction31) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}

// Transaction information in the authorisation response.
type CardPaymentTransaction32 struct {

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max35Text `xml:"IntrchngData,omitempty"`

	// Detail of the transaction transported.
	TransactionDetails *CardPaymentTransactionDetails17 `xml:"TxDtls"`
}

func (c *CardPaymentTransaction32) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction32) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction32) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction32) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction32) SetInterchangeData(value string) {
	c.InterchangeData = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction32) AddTransactionDetails() *CardPaymentTransactionDetails17 {
	c.TransactionDetails = new(CardPaymentTransactionDetails17)
	return c.TransactionDetails
}

// Identification of the original transaction.
type CardPaymentTransaction33 struct {

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the POI (Point Of Interaction) performing the transaction.
	POIIdentification *GenericIdentification32 `xml:"POIId"`

	// Link to a previous currency conversion.
	CurrencyConversion *CurrencyConversion1 `xml:"CcyConvs"`
}

func (c *CardPaymentTransaction33) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction33) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction33) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

func (c *CardPaymentTransaction33) AddCurrencyConversion() *CurrencyConversion1 {
	c.CurrencyConversion = new(CurrencyConversion1)
	return c.CurrencyConversion
}

// Data associated with the transaction for a potential currency conversion.
type CardPaymentTransaction34 struct {

	// Flag indicating whether the transaction data must be captured or not in addition to the message process.
	TransactionCapture *TrueFalseIndicator `xml:"TxCaptr"`

	// Type of transaction being undertaken for the main service.
	TransactionType *CardPaymentServiceType5Code `xml:"TxTp"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType6Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd"`

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction33 `xml:"OrgnlTx,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails18 `xml:"TxDtls"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction34) SetTransactionCapture(value string) {
	c.TransactionCapture = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction34) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType5Code)(&value)
}

func (c *CardPaymentTransaction34) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType6Code)(&value))
}

func (c *CardPaymentTransaction34) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CardPaymentTransaction34) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction34) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction34) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction34) AddOriginalTransaction() *CardPaymentTransaction33 {
	c.OriginalTransaction = new(CardPaymentTransaction33)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction34) AddTransactionDetails() *CardPaymentTransactionDetails18 {
	c.TransactionDetails = new(CardPaymentTransactionDetails18)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction34) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}

// Transaction information in the cancellation response.
type CardPaymentTransaction35 struct {

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max35Text `xml:"IntrchngData,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails6 `xml:"TxDtls"`
}

func (c *CardPaymentTransaction35) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction35) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction35) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction35) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction35) SetInterchangeData(value string) {
	c.InterchangeData = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction35) AddTransactionDetails() *CardPaymentTransactionDetails6 {
	c.TransactionDetails = new(CardPaymentTransactionDetails6)
	return c.TransactionDetails
}

// Data associated with the transaction during the authorisation.
type CardPaymentTransaction36 struct {

	// Flag indicating whether the transaction data must be captured or not in addition to the message process.
	TransactionCapture *TrueFalseIndicator `xml:"TxCaptr"`

	// Type of transaction being undertaken for the main service.
	TransactionType *CardPaymentServiceType5Code `xml:"TxTp"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType6Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd"`

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction37 `xml:"OrgnlTx,omitempty"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails19 `xml:"TxDtls"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction36) SetTransactionCapture(value string) {
	c.TransactionCapture = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction36) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType5Code)(&value)
}

func (c *CardPaymentTransaction36) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType6Code)(&value))
}

func (c *CardPaymentTransaction36) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CardPaymentTransaction36) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction36) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction36) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction36) AddOriginalTransaction() *CardPaymentTransaction37 {
	c.OriginalTransaction = new(CardPaymentTransaction37)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction36) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction36) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction36) AddTransactionDetails() *CardPaymentTransactionDetails19 {
	c.TransactionDetails = new(CardPaymentTransactionDetails19)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction36) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}

// Identification of the original transaction.
type CardPaymentTransaction37 struct {

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the POI (Point Of Interaction) performing the transaction.
	POIIdentification *GenericIdentification32 `xml:"POIId,omitempty"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Type of transaction being undertaken for the main service.
	TransactionType *CardPaymentServiceType5Code `xml:"TxTp"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType6Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Result of the original transaction.
	TransactionResult *CardPaymentTransactionResult2 `xml:"TxRslt,omitempty"`
}

func (c *CardPaymentTransaction37) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction37) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction37) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

func (c *CardPaymentTransaction37) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction37) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction37) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType5Code)(&value)
}

func (c *CardPaymentTransaction37) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType6Code)(&value))
}

func (c *CardPaymentTransaction37) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CardPaymentTransaction37) AddTransactionResult() *CardPaymentTransactionResult2 {
	c.TransactionResult = new(CardPaymentTransactionResult2)
	return c.TransactionResult
}

// Transaction information in the authorisation response.
type CardPaymentTransaction38 struct {

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max140Text `xml:"IntrchngData,omitempty"`

	// Detail of the transaction transported.
	TransactionDetails *CardPaymentTransactionDetails20 `xml:"TxDtls"`
}

func (c *CardPaymentTransaction38) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction38) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction38) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction38) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction38) SetInterchangeData(value string) {
	c.InterchangeData = (*Max140Text)(&value)
}

func (c *CardPaymentTransaction38) AddTransactionDetails() *CardPaymentTransactionDetails20 {
	c.TransactionDetails = new(CardPaymentTransactionDetails20)
	return c.TransactionDetails
}

// Authorisation response from the acquirer.
type CardPaymentTransaction39 struct {

	// Outcome of the authorisation, and actions to perform.
	AuthorisationResult *AuthorisationResult4 `xml:"AuthstnRslt"`

	// Result of the verifications performed by the issuer to deliver or decline the authorisation.
	TransactionVerificationResult []*TransactionVerificationResult3 `xml:"TxVrfctnRslt,omitempty"`

	// Product code for which the authorisation was declined.
	DeclinedProductCode []*Max70Text `xml:"DclndPdctCd,omitempty"`

	// Balance of the account, related to the payment.
	Balance *AmountAndDirection41 `xml:"Bal,omitempty"`

	// Encrypted balance of the account.
	ProtectedBalance *ContentInformationType10 `xml:"PrtctdBal,omitempty"`

	// Set of actions to be performed by the POI (Point Of Interaction) system.
	Action []*Action3 `xml:"Actn,omitempty"`

	// Conversion between the currency of a card acceptor and the currency of a card issuer, provided by a dedicated service provider. The currency conversion has to be accepted by the cardholder.
	CurrencyConversion *CurrencyConversion3 `xml:"CcyConvs,omitempty"`
}

func (c *CardPaymentTransaction39) AddAuthorisationResult() *AuthorisationResult4 {
	c.AuthorisationResult = new(AuthorisationResult4)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction39) AddTransactionVerificationResult() *TransactionVerificationResult3 {
	newValue := new(TransactionVerificationResult3)
	c.TransactionVerificationResult = append(c.TransactionVerificationResult, newValue)
	return newValue
}

func (c *CardPaymentTransaction39) AddDeclinedProductCode(value string) {
	c.DeclinedProductCode = append(c.DeclinedProductCode, (*Max70Text)(&value))
}

func (c *CardPaymentTransaction39) AddBalance() *AmountAndDirection41 {
	c.Balance = new(AmountAndDirection41)
	return c.Balance
}

func (c *CardPaymentTransaction39) AddProtectedBalance() *ContentInformationType10 {
	c.ProtectedBalance = new(ContentInformationType10)
	return c.ProtectedBalance
}

func (c *CardPaymentTransaction39) AddAction() *Action3 {
	newValue := new(Action3)
	c.Action = append(c.Action, newValue)
	return newValue
}

func (c *CardPaymentTransaction39) AddCurrencyConversion() *CurrencyConversion3 {
	c.CurrencyConversion = new(CurrencyConversion3)
	return c.CurrencyConversion
}

// Transaction information in the batch capture.
type CardPaymentTransaction4 struct {

	// Type of transaction being undertaken for the main service.
	TransactionType *CardPaymentServiceType1Code `xml:"TxTp,omitempty"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType2Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction8 `xml:"OrgnlTx,omitempty"`

	// Outcome of the transaction at the acceptor.
	TransactionSuccess *TrueFalseIndicator `xml:"TxSucss"`

	// Notify that a previous transaction has to be reversed if this original transaction has been approved by the acquirer.
	Reversal *TrueFalseIndicator `xml:"Rvsl,omitempty"`

	// Indicate that the acceptor has forced the transaction in spite of the authorisation result (online or offline), or incident to complete the transaction.
	MerchantOverride *TrueFalseIndicator `xml:"MrchntOvrrd,omitempty"`

	// List of incidents during the transaction.
	FailureReason []*FailureReason1Code `xml:"FailrRsn,omitempty"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max35Text `xml:"IntrchngData,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails4 `xml:"TxDtls"`

	// Outcome of the authorisation request, and actions to perform.
	AuthorisationResult *AuthorisationResult1 `xml:"AuthstnRslt,omitempty"`

	// Result of the verifications performed by the issuer to deliver or decline the authorisation.
	TransactionVerificationResult *TransactionVerificationResult1 `xml:"TxVrfctnRslt,omitempty"`

	// Additional information related to the transaction.
	AdditionalTransactionData *Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction4) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType1Code)(&value)
}

func (c *CardPaymentTransaction4) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType2Code)(&value))
}

func (c *CardPaymentTransaction4) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CardPaymentTransaction4) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction4) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction4) AddOriginalTransaction() *CardPaymentTransaction8 {
	c.OriginalTransaction = new(CardPaymentTransaction8)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction4) SetTransactionSuccess(value string) {
	c.TransactionSuccess = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction4) SetReversal(value string) {
	c.Reversal = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction4) SetMerchantOverride(value string) {
	c.MerchantOverride = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction4) AddFailureReason(value string) {
	c.FailureReason = append(c.FailureReason, (*FailureReason1Code)(&value))
}

func (c *CardPaymentTransaction4) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction4) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction4) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction4) SetInterchangeData(value string) {
	c.InterchangeData = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction4) AddTransactionDetails() *CardPaymentTransactionDetails4 {
	c.TransactionDetails = new(CardPaymentTransactionDetails4)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction4) AddAuthorisationResult() *AuthorisationResult1 {
	c.AuthorisationResult = new(AuthorisationResult1)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction4) AddTransactionVerificationResult() *TransactionVerificationResult1 {
	c.TransactionVerificationResult = new(TransactionVerificationResult1)
	return c.TransactionVerificationResult
}

func (c *CardPaymentTransaction4) SetAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = (*Max70Text)(&value)
}

// Transaction information in the completion advice message.
type CardPaymentTransaction40 struct {

	// Flag indicating whether the transaction data must be captured or not in addition to the message process.
	TransactionCapture *TrueFalseIndicator `xml:"TxCaptr,omitempty"`

	// Type of transaction being undertaken for the main service.
	TransactionType *CardPaymentServiceType5Code `xml:"TxTp"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType6Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd"`

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction37 `xml:"OrgnlTx,omitempty"`

	// Outcome of the transaction at the acceptor.
	TransactionSuccess *TrueFalseIndicator `xml:"TxSucss"`

	// Notify that a previous transaction has to be reversed if this original transaction has been approved by the acquirer.
	Reversal *TrueFalseIndicator `xml:"Rvsl,omitempty"`

	// Indicate that the acceptor has forced the transaction in spite of the authorisation result (online or offline), or incident to complete the transaction.
	MerchantOverride *TrueFalseIndicator `xml:"MrchntOvrrd,omitempty"`

	// List of incidents during the transaction.
	FailureReason []*FailureReason3Code `xml:"FailrRsn,omitempty"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max140Text `xml:"IntrchngData,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails21 `xml:"TxDtls"`

	// Outcome of the authorisation request.
	AuthorisationResult *AuthorisationResult5 `xml:"AuthstnRslt,omitempty"`

	// Result of the performed verifications for the transaction.
	TransactionVerificationResult []*TransactionVerificationResult3 `xml:"TxVrfctnRslt,omitempty"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction40) SetTransactionCapture(value string) {
	c.TransactionCapture = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction40) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType5Code)(&value)
}

func (c *CardPaymentTransaction40) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType6Code)(&value))
}

func (c *CardPaymentTransaction40) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CardPaymentTransaction40) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction40) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction40) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction40) AddOriginalTransaction() *CardPaymentTransaction37 {
	c.OriginalTransaction = new(CardPaymentTransaction37)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction40) SetTransactionSuccess(value string) {
	c.TransactionSuccess = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction40) SetReversal(value string) {
	c.Reversal = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction40) SetMerchantOverride(value string) {
	c.MerchantOverride = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction40) AddFailureReason(value string) {
	c.FailureReason = append(c.FailureReason, (*FailureReason3Code)(&value))
}

func (c *CardPaymentTransaction40) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction40) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction40) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction40) SetInterchangeData(value string) {
	c.InterchangeData = (*Max140Text)(&value)
}

func (c *CardPaymentTransaction40) AddTransactionDetails() *CardPaymentTransactionDetails21 {
	c.TransactionDetails = new(CardPaymentTransactionDetails21)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction40) AddAuthorisationResult() *AuthorisationResult5 {
	c.AuthorisationResult = new(AuthorisationResult5)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction40) AddTransactionVerificationResult() *TransactionVerificationResult3 {
	newValue := new(TransactionVerificationResult3)
	c.TransactionVerificationResult = append(c.TransactionVerificationResult, newValue)
	return newValue
}

func (c *CardPaymentTransaction40) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}

// Transaction information in the cancellation request.
type CardPaymentTransaction41 struct {

	// Flag indicating whether the transaction data must be captured or not in addition to the message process.
	TransactionCapture *TrueFalseIndicator `xml:"TxCaptr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd"`

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction37 `xml:"OrgnlTx"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails5 `xml:"TxDtls"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction41) SetTransactionCapture(value string) {
	c.TransactionCapture = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction41) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction41) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction41) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction41) AddOriginalTransaction() *CardPaymentTransaction37 {
	c.OriginalTransaction = new(CardPaymentTransaction37)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction41) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction41) AddTransactionDetails() *CardPaymentTransactionDetails5 {
	c.TransactionDetails = new(CardPaymentTransactionDetails5)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction41) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}

// Transaction information in the cancellation response.
type CardPaymentTransaction42 struct {

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max140Text `xml:"IntrchngData,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails22 `xml:"TxDtls"`
}

func (c *CardPaymentTransaction42) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction42) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction42) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction42) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction42) SetInterchangeData(value string) {
	c.InterchangeData = (*Max140Text)(&value)
}

func (c *CardPaymentTransaction42) AddTransactionDetails() *CardPaymentTransactionDetails22 {
	c.TransactionDetails = new(CardPaymentTransactionDetails22)
	return c.TransactionDetails
}

// Cancellation response from the acquirer.
type CardPaymentTransaction43 struct {

	// Outcome of the authorisation, and actions to perform.
	AuthorisationResult *AuthorisationResult6 `xml:"AuthstnRslt"`

	// Set of actions to be performed by the POI (Point Of Interaction) system.
	Action []*Action3 `xml:"Actn,omitempty"`
}

func (c *CardPaymentTransaction43) AddAuthorisationResult() *AuthorisationResult6 {
	c.AuthorisationResult = new(AuthorisationResult6)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction43) AddAction() *Action3 {
	newValue := new(Action3)
	c.Action = append(c.Action, newValue)
	return newValue
}

// Transaction information in the cancellation advice.
type CardPaymentTransaction44 struct {

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd"`

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction37 `xml:"OrgnlTx,omitempty"`

	// Outcome of the transaction at the acceptor.
	TransactionSuccess *TrueFalseIndicator `xml:"TxSucss"`

	// Notify that a previous transaction has to be reversed if this original transaction has been approved by the acquirer.
	Reversal *TrueFalseIndicator `xml:"Rvsl,omitempty"`

	// List of incidents during the transaction.
	FailureReason []*FailureReason3Code `xml:"FailrRsn,omitempty"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max140Text `xml:"IntrchngData,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails7 `xml:"TxDtls"`

	// Outcome of the authorisation.
	AuthorisationResult *AuthorisationResult5 `xml:"AuthstnRslt,omitempty"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction44) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction44) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction44) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction44) AddOriginalTransaction() *CardPaymentTransaction37 {
	c.OriginalTransaction = new(CardPaymentTransaction37)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction44) SetTransactionSuccess(value string) {
	c.TransactionSuccess = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction44) SetReversal(value string) {
	c.Reversal = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction44) AddFailureReason(value string) {
	c.FailureReason = append(c.FailureReason, (*FailureReason3Code)(&value))
}

func (c *CardPaymentTransaction44) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction44) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction44) SetInterchangeData(value string) {
	c.InterchangeData = (*Max140Text)(&value)
}

func (c *CardPaymentTransaction44) AddTransactionDetails() *CardPaymentTransactionDetails7 {
	c.TransactionDetails = new(CardPaymentTransactionDetails7)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction44) AddAuthorisationResult() *AuthorisationResult5 {
	c.AuthorisationResult = new(AuthorisationResult5)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction44) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}

// Transaction information in the batch capture.
type CardPaymentTransaction45 struct {

	// Type of transaction being undertaken for the main service.
	TransactionType *CardPaymentServiceType5Code `xml:"TxTp,omitempty"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType6Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd,omitempty"`

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction37 `xml:"OrgnlTx,omitempty"`

	// Outcome of the transaction at the acceptor.
	TransactionSuccess *TrueFalseIndicator `xml:"TxSucss"`

	// Notify that a previous transaction has to be reversed if this original transaction has been approved by the acquirer.
	Reversal *TrueFalseIndicator `xml:"Rvsl,omitempty"`

	// Indicate that the acceptor has forced the transaction in spite of the authorisation result (online or offline), or incident to complete the transaction.
	MerchantOverride *TrueFalseIndicator `xml:"MrchntOvrrd,omitempty"`

	// List of incidents during the transaction.
	FailureReason []*FailureReason3Code `xml:"FailrRsn,omitempty"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max140Text `xml:"IntrchngData,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails23 `xml:"TxDtls"`

	// Outcome of the authorisation request, and actions to perform.
	AuthorisationResult *AuthorisationResult4 `xml:"AuthstnRslt,omitempty"`

	// Result of performed verifications for the transaction.
	TransactionVerificationResult []*TransactionVerificationResult3 `xml:"TxVrfctnRslt,omitempty"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction45) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType5Code)(&value)
}

func (c *CardPaymentTransaction45) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType6Code)(&value))
}

func (c *CardPaymentTransaction45) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CardPaymentTransaction45) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction45) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction45) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction45) AddOriginalTransaction() *CardPaymentTransaction37 {
	c.OriginalTransaction = new(CardPaymentTransaction37)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction45) SetTransactionSuccess(value string) {
	c.TransactionSuccess = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction45) SetReversal(value string) {
	c.Reversal = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction45) SetMerchantOverride(value string) {
	c.MerchantOverride = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction45) AddFailureReason(value string) {
	c.FailureReason = append(c.FailureReason, (*FailureReason3Code)(&value))
}

func (c *CardPaymentTransaction45) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction45) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction45) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction45) SetInterchangeData(value string) {
	c.InterchangeData = (*Max140Text)(&value)
}

func (c *CardPaymentTransaction45) AddTransactionDetails() *CardPaymentTransactionDetails23 {
	c.TransactionDetails = new(CardPaymentTransactionDetails23)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction45) AddAuthorisationResult() *AuthorisationResult4 {
	c.AuthorisationResult = new(AuthorisationResult4)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction45) AddTransactionVerificationResult() *TransactionVerificationResult3 {
	newValue := new(TransactionVerificationResult3)
	c.TransactionVerificationResult = append(c.TransactionVerificationResult, newValue)
	return newValue
}

func (c *CardPaymentTransaction45) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}

// Transaction information in the cancellation advice.
type CardPaymentTransaction46 struct {

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd,omitempty"`

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction37 `xml:"OrgnlTx,omitempty"`

	// Outcome of the transaction at the acceptor.
	TransactionSuccess *TrueFalseIndicator `xml:"TxSucss"`

	// Notify that a previous transaction has to be reversed if this original transaction has been approved by the acquirer.
	Reversal *TrueFalseIndicator `xml:"Rvsl,omitempty"`

	// List of incidents during the transaction.
	FailureReason []*FailureReason3Code `xml:"FailrRsn,omitempty"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max140Text `xml:"IntrchngData,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails9 `xml:"TxDtls"`

	// Outcome of the authorisation.
	AuthorisationResult *AuthorisationResult6 `xml:"AuthstnRslt,omitempty"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction46) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction46) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction46) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction46) AddOriginalTransaction() *CardPaymentTransaction37 {
	c.OriginalTransaction = new(CardPaymentTransaction37)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction46) SetTransactionSuccess(value string) {
	c.TransactionSuccess = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction46) SetReversal(value string) {
	c.Reversal = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction46) AddFailureReason(value string) {
	c.FailureReason = append(c.FailureReason, (*FailureReason3Code)(&value))
}

func (c *CardPaymentTransaction46) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction46) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction46) SetInterchangeData(value string) {
	c.InterchangeData = (*Max140Text)(&value)
}

func (c *CardPaymentTransaction46) AddTransactionDetails() *CardPaymentTransactionDetails9 {
	c.TransactionDetails = new(CardPaymentTransactionDetails9)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction46) AddAuthorisationResult() *AuthorisationResult6 {
	c.AuthorisationResult = new(AuthorisationResult6)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction46) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}

// Data associated with the transaction to authorise.
type CardPaymentTransaction47 struct {

	// Flag indicating whether the transaction data must be captured or not in addition to the message process.
	TransactionCapture *TrueFalseIndicator `xml:"TxCaptr"`

	// Type of transaction being undertaken for the main service.
	TransactionType *CardPaymentServiceType5Code `xml:"TxTp,omitempty"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType6Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd,omitempty"`

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction37 `xml:"OrgnlTx,omitempty"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails24 `xml:"TxDtls"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction47) SetTransactionCapture(value string) {
	c.TransactionCapture = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction47) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType5Code)(&value)
}

func (c *CardPaymentTransaction47) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType6Code)(&value))
}

func (c *CardPaymentTransaction47) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CardPaymentTransaction47) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction47) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction47) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction47) AddOriginalTransaction() *CardPaymentTransaction37 {
	c.OriginalTransaction = new(CardPaymentTransaction37)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction47) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction47) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction47) AddTransactionDetails() *CardPaymentTransactionDetails24 {
	c.TransactionDetails = new(CardPaymentTransactionDetails24)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction47) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}

// Transaction information in the authorisation response.
type CardPaymentTransaction48 struct {

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max140Text `xml:"IntrchngData,omitempty"`

	// Detail of the transaction transported.
	TransactionDetails *CardPaymentTransactionDetails25 `xml:"TxDtls"`
}

func (c *CardPaymentTransaction48) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction48) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction48) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction48) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction48) SetInterchangeData(value string) {
	c.InterchangeData = (*Max140Text)(&value)
}

func (c *CardPaymentTransaction48) AddTransactionDetails() *CardPaymentTransactionDetails25 {
	c.TransactionDetails = new(CardPaymentTransactionDetails25)
	return c.TransactionDetails
}

// Data associated with the transaction for a potential currency conversion.
type CardPaymentTransaction49 struct {

	// Flag indicating whether the transaction data must be captured or not in addition to the message process.
	TransactionCapture *TrueFalseIndicator `xml:"TxCaptr"`

	// Type of transaction being undertaken for the main service.
	TransactionType *CardPaymentServiceType5Code `xml:"TxTp"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType6Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd"`

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction50 `xml:"OrgnlTx,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails26 `xml:"TxDtls"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction49) SetTransactionCapture(value string) {
	c.TransactionCapture = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction49) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType5Code)(&value)
}

func (c *CardPaymentTransaction49) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType6Code)(&value))
}

func (c *CardPaymentTransaction49) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CardPaymentTransaction49) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction49) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction49) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction49) AddOriginalTransaction() *CardPaymentTransaction50 {
	c.OriginalTransaction = new(CardPaymentTransaction50)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction49) AddTransactionDetails() *CardPaymentTransactionDetails26 {
	c.TransactionDetails = new(CardPaymentTransactionDetails26)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction49) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}

// Transaction information in the cancellation request.
type CardPaymentTransaction5 struct {

	// Flag indicating whether the transaction data must be captured or not in addition to the message process.
	TransactionCapture *TrueFalseIndicator `xml:"TxCaptr"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction8 `xml:"OrgnlTx"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails5 `xml:"TxDtls"`

	// Additional information related to the transaction.
	AdditionalTransactionData *Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction5) SetTransactionCapture(value string) {
	c.TransactionCapture = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction5) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction5) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction5) AddOriginalTransaction() *CardPaymentTransaction8 {
	c.OriginalTransaction = new(CardPaymentTransaction8)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction5) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction5) AddTransactionDetails() *CardPaymentTransactionDetails5 {
	c.TransactionDetails = new(CardPaymentTransactionDetails5)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction5) SetAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = (*Max70Text)(&value)
}

// Identification of the original transaction.
type CardPaymentTransaction50 struct {

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the POI (Point Of Interaction) performing the transaction.
	POIIdentification *GenericIdentification32 `xml:"POIId"`

	// Link to a previous currency conversion.
	CurrencyConversion *CurrencyConversion3 `xml:"CcyConvs"`
}

func (c *CardPaymentTransaction50) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction50) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction50) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

func (c *CardPaymentTransaction50) AddCurrencyConversion() *CurrencyConversion3 {
	c.CurrencyConversion = new(CurrencyConversion3)
	return c.CurrencyConversion
}

// Data associated with the transaction during the authorisation.
type CardPaymentTransaction51 struct {

	// Flag indicating whether the transaction data must be captured or not in addition to the message process.
	TransactionCapture *TrueFalseIndicator `xml:"TxCaptr"`

	// Type of transaction being undertaken for the main service.
	TransactionType *CardPaymentServiceType5Code `xml:"TxTp"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType9Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd"`

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction52 `xml:"OrgnlTx,omitempty"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails27 `xml:"TxDtls"`

	// Merchant information that must be returned unchanged in the response.
	MerchantReferenceData *Max70Text `xml:"MrchntRefData,omitempty"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction51) SetTransactionCapture(value string) {
	c.TransactionCapture = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction51) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType5Code)(&value)
}

func (c *CardPaymentTransaction51) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType9Code)(&value))
}

func (c *CardPaymentTransaction51) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CardPaymentTransaction51) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction51) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction51) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction51) AddOriginalTransaction() *CardPaymentTransaction52 {
	c.OriginalTransaction = new(CardPaymentTransaction52)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction51) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction51) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction51) AddTransactionDetails() *CardPaymentTransactionDetails27 {
	c.TransactionDetails = new(CardPaymentTransactionDetails27)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction51) SetMerchantReferenceData(value string) {
	c.MerchantReferenceData = (*Max70Text)(&value)
}

func (c *CardPaymentTransaction51) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}

// Identification of the original transaction.
type CardPaymentTransaction52 struct {

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the POI (Point Of Interaction) performing the transaction.
	POIIdentification *GenericIdentification32 `xml:"POIId,omitempty"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Type of transaction being undertaken for the main service.
	TransactionType *CardPaymentServiceType5Code `xml:"TxTp"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType9Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Entry mode of the card information.
	CardDataEntryMode *CardDataReading5Code `xml:"CardDataNtryMd,omitempty"`

	// Result of the original transaction.
	TransactionResult *CardPaymentTransactionResult3 `xml:"TxRslt,omitempty"`
}

func (c *CardPaymentTransaction52) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction52) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction52) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

func (c *CardPaymentTransaction52) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction52) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction52) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType5Code)(&value)
}

func (c *CardPaymentTransaction52) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType9Code)(&value))
}

func (c *CardPaymentTransaction52) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CardPaymentTransaction52) SetCardDataEntryMode(value string) {
	c.CardDataEntryMode = (*CardDataReading5Code)(&value)
}

func (c *CardPaymentTransaction52) AddTransactionResult() *CardPaymentTransactionResult3 {
	c.TransactionResult = new(CardPaymentTransactionResult3)
	return c.TransactionResult
}

// Transaction information in the authorisation response.
type CardPaymentTransaction53 struct {

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max140Text `xml:"IntrchngData,omitempty"`

	// Detail of the transaction transported.
	TransactionDetails *CardPaymentTransactionDetails28 `xml:"TxDtls"`

	// Merchant related information provided in the request.
	MerchantReferenceData *Max70Text `xml:"MrchntRefData,omitempty"`
}

func (c *CardPaymentTransaction53) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction53) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction53) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction53) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction53) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction53) SetInterchangeData(value string) {
	c.InterchangeData = (*Max140Text)(&value)
}

func (c *CardPaymentTransaction53) AddTransactionDetails() *CardPaymentTransactionDetails28 {
	c.TransactionDetails = new(CardPaymentTransactionDetails28)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction53) SetMerchantReferenceData(value string) {
	c.MerchantReferenceData = (*Max70Text)(&value)
}

// Authorisation response from the acquirer.
type CardPaymentTransaction54 struct {

	// Outcome of the authorisation, and actions to perform.
	AuthorisationResult *AuthorisationResult10 `xml:"AuthstnRslt"`

	// Result of the verifications performed by the issuer to deliver or decline the authorisation.
	TransactionVerificationResult []*TransactionVerificationResult4 `xml:"TxVrfctnRslt,omitempty"`

	// Product code which are allowed by the payment card.
	AllowedProductCode []*Product4 `xml:"AllwdPdctCd,omitempty"`

	// Product code not allowed by the payment card.
	NotAllowedProductCode []*Product4 `xml:"NotAllwdPdctCd,omitempty"`

	// Products that may be added to the purchase after the authorisation.
	AdditionalAvailableProduct []*Product5 `xml:"AddtlAvlblPdct,omitempty"`

	// Balance of the account, related to the payment.
	Balance *AmountAndDirection41 `xml:"Bal,omitempty"`

	// Encrypted balance of the account.
	ProtectedBalance *ContentInformationType10 `xml:"PrtctdBal,omitempty"`

	// Set of actions to be performed by the POI (Point Of Interaction) system.
	Action []*Action6 `xml:"Actn,omitempty"`

	// Conversion between the currency of a card acceptor and the currency of a card issuer, provided by a dedicated service provider. The currency conversion has to be proposed to the cardholder.
	CurrencyConversionEligibility *CurrencyConversion6 `xml:"CcyConvsElgblty,omitempty"`
}

func (c *CardPaymentTransaction54) AddAuthorisationResult() *AuthorisationResult10 {
	c.AuthorisationResult = new(AuthorisationResult10)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction54) AddTransactionVerificationResult() *TransactionVerificationResult4 {
	newValue := new(TransactionVerificationResult4)
	c.TransactionVerificationResult = append(c.TransactionVerificationResult, newValue)
	return newValue
}

func (c *CardPaymentTransaction54) AddAllowedProductCode() *Product4 {
	newValue := new(Product4)
	c.AllowedProductCode = append(c.AllowedProductCode, newValue)
	return newValue
}

func (c *CardPaymentTransaction54) AddNotAllowedProductCode() *Product4 {
	newValue := new(Product4)
	c.NotAllowedProductCode = append(c.NotAllowedProductCode, newValue)
	return newValue
}

func (c *CardPaymentTransaction54) AddAdditionalAvailableProduct() *Product5 {
	newValue := new(Product5)
	c.AdditionalAvailableProduct = append(c.AdditionalAvailableProduct, newValue)
	return newValue
}

func (c *CardPaymentTransaction54) AddBalance() *AmountAndDirection41 {
	c.Balance = new(AmountAndDirection41)
	return c.Balance
}

func (c *CardPaymentTransaction54) AddProtectedBalance() *ContentInformationType10 {
	c.ProtectedBalance = new(ContentInformationType10)
	return c.ProtectedBalance
}

func (c *CardPaymentTransaction54) AddAction() *Action6 {
	newValue := new(Action6)
	c.Action = append(c.Action, newValue)
	return newValue
}

func (c *CardPaymentTransaction54) AddCurrencyConversionEligibility() *CurrencyConversion6 {
	c.CurrencyConversionEligibility = new(CurrencyConversion6)
	return c.CurrencyConversionEligibility
}

// Transaction information in the completion advice message.
type CardPaymentTransaction55 struct {

	// Flag indicating whether the transaction data must be captured or not in addition to the message process.
	TransactionCapture *TrueFalseIndicator `xml:"TxCaptr,omitempty"`

	// Type of transaction being undertaken for the main service.
	TransactionType *CardPaymentServiceType5Code `xml:"TxTp"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType9Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd"`

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction52 `xml:"OrgnlTx,omitempty"`

	// Outcome of the transaction at the acceptor.
	TransactionSuccess *TrueFalseIndicator `xml:"TxSucss"`

	// Notify that a previous transaction has to be reversed if this original transaction has been approved by the acquirer.
	Reversal *TrueFalseIndicator `xml:"Rvsl,omitempty"`

	// Indicate that the acceptor has forced the transaction in spite of the authorisation result (online or offline), or incident to complete the transaction.
	MerchantOverride *TrueFalseIndicator `xml:"MrchntOvrrd,omitempty"`

	// List of incidents during the transaction.
	FailureReason []*FailureReason3Code `xml:"FailrRsn,omitempty"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max140Text `xml:"IntrchngData,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails29 `xml:"TxDtls"`

	// Outcome of the authorisation request.
	AuthorisationResult *AuthorisationResult11 `xml:"AuthstnRslt,omitempty"`

	// Result of the performed verifications for the transaction.
	TransactionVerificationResult []*TransactionVerificationResult4 `xml:"TxVrfctnRslt,omitempty"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction55) SetTransactionCapture(value string) {
	c.TransactionCapture = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction55) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType5Code)(&value)
}

func (c *CardPaymentTransaction55) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType9Code)(&value))
}

func (c *CardPaymentTransaction55) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CardPaymentTransaction55) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction55) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction55) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction55) AddOriginalTransaction() *CardPaymentTransaction52 {
	c.OriginalTransaction = new(CardPaymentTransaction52)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction55) SetTransactionSuccess(value string) {
	c.TransactionSuccess = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction55) SetReversal(value string) {
	c.Reversal = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction55) SetMerchantOverride(value string) {
	c.MerchantOverride = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction55) AddFailureReason(value string) {
	c.FailureReason = append(c.FailureReason, (*FailureReason3Code)(&value))
}

func (c *CardPaymentTransaction55) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction55) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction55) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction55) SetInterchangeData(value string) {
	c.InterchangeData = (*Max140Text)(&value)
}

func (c *CardPaymentTransaction55) AddTransactionDetails() *CardPaymentTransactionDetails29 {
	c.TransactionDetails = new(CardPaymentTransactionDetails29)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction55) AddAuthorisationResult() *AuthorisationResult11 {
	c.AuthorisationResult = new(AuthorisationResult11)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction55) AddTransactionVerificationResult() *TransactionVerificationResult4 {
	newValue := new(TransactionVerificationResult4)
	c.TransactionVerificationResult = append(c.TransactionVerificationResult, newValue)
	return newValue
}

func (c *CardPaymentTransaction55) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}

// Transaction information in the cancellation request.
type CardPaymentTransaction56 struct {

	// Flag indicating whether the transaction data must be captured or not in addition to the message process.
	TransactionCapture *TrueFalseIndicator `xml:"TxCaptr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd"`

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction52 `xml:"OrgnlTx"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails34 `xml:"TxDtls"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction56) SetTransactionCapture(value string) {
	c.TransactionCapture = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction56) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction56) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction56) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction56) AddOriginalTransaction() *CardPaymentTransaction52 {
	c.OriginalTransaction = new(CardPaymentTransaction52)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction56) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction56) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction56) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction56) AddTransactionDetails() *CardPaymentTransactionDetails34 {
	c.TransactionDetails = new(CardPaymentTransactionDetails34)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction56) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}

// Transaction information in the cancellation response.
type CardPaymentTransaction57 struct {

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max140Text `xml:"IntrchngData,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails35 `xml:"TxDtls"`
}

func (c *CardPaymentTransaction57) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction57) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction57) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction57) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction57) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction57) SetInterchangeData(value string) {
	c.InterchangeData = (*Max140Text)(&value)
}

func (c *CardPaymentTransaction57) AddTransactionDetails() *CardPaymentTransactionDetails35 {
	c.TransactionDetails = new(CardPaymentTransactionDetails35)
	return c.TransactionDetails
}

// Cancellation response from the acquirer.
type CardPaymentTransaction58 struct {

	// Outcome of the authorisation, and actions to perform.
	AuthorisationResult *AuthorisationResult12 `xml:"AuthstnRslt"`

	// Set of actions to be performed by the POI (Point Of Interaction) system.
	Action []*Action6 `xml:"Actn,omitempty"`
}

func (c *CardPaymentTransaction58) AddAuthorisationResult() *AuthorisationResult12 {
	c.AuthorisationResult = new(AuthorisationResult12)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction58) AddAction() *Action6 {
	newValue := new(Action6)
	c.Action = append(c.Action, newValue)
	return newValue
}

// Transaction information in the cancellation advice.
type CardPaymentTransaction59 struct {

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd"`

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction52 `xml:"OrgnlTx,omitempty"`

	// Outcome of the transaction at the acceptor.
	TransactionSuccess *TrueFalseIndicator `xml:"TxSucss"`

	// Notify that a previous transaction has to be reversed if this original transaction has been approved by the acquirer.
	Reversal *TrueFalseIndicator `xml:"Rvsl,omitempty"`

	// List of incidents during the transaction.
	FailureReason []*FailureReason3Code `xml:"FailrRsn,omitempty"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max140Text `xml:"IntrchngData,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails34 `xml:"TxDtls"`

	// Outcome of the authorisation.
	AuthorisationResult *AuthorisationResult11 `xml:"AuthstnRslt,omitempty"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction59) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction59) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction59) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction59) AddOriginalTransaction() *CardPaymentTransaction52 {
	c.OriginalTransaction = new(CardPaymentTransaction52)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction59) SetTransactionSuccess(value string) {
	c.TransactionSuccess = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction59) SetReversal(value string) {
	c.Reversal = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction59) AddFailureReason(value string) {
	c.FailureReason = append(c.FailureReason, (*FailureReason3Code)(&value))
}

func (c *CardPaymentTransaction59) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction59) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction59) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction59) SetInterchangeData(value string) {
	c.InterchangeData = (*Max140Text)(&value)
}

func (c *CardPaymentTransaction59) AddTransactionDetails() *CardPaymentTransactionDetails34 {
	c.TransactionDetails = new(CardPaymentTransactionDetails34)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction59) AddAuthorisationResult() *AuthorisationResult11 {
	c.AuthorisationResult = new(AuthorisationResult11)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction59) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}

// Transaction information in the cancellation response.
type CardPaymentTransaction6 struct {

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max35Text `xml:"IntrchngData,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails6 `xml:"TxDtls"`
}

func (c *CardPaymentTransaction6) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction6) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction6) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction6) SetInterchangeData(value string) {
	c.InterchangeData = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction6) AddTransactionDetails() *CardPaymentTransactionDetails6 {
	c.TransactionDetails = new(CardPaymentTransactionDetails6)
	return c.TransactionDetails
}

// Transaction information in the batch capture.
type CardPaymentTransaction60 struct {

	// Type of transaction being undertaken for the main service.
	TransactionType *CardPaymentServiceType5Code `xml:"TxTp,omitempty"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType9Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd,omitempty"`

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction52 `xml:"OrgnlTx,omitempty"`

	// Outcome of the transaction at the acceptor.
	TransactionSuccess *TrueFalseIndicator `xml:"TxSucss"`

	// Notify that a previous transaction has to be reversed if this original transaction has been approved by the acquirer.
	Reversal *TrueFalseIndicator `xml:"Rvsl,omitempty"`

	// Indicate that the acceptor has forced the transaction in spite of the authorisation result (online or offline), or incident to complete the transaction.
	MerchantOverride *TrueFalseIndicator `xml:"MrchntOvrrd,omitempty"`

	// List of incidents during the transaction.
	FailureReason []*FailureReason3Code `xml:"FailrRsn,omitempty"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max140Text `xml:"IntrchngData,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails30 `xml:"TxDtls"`

	// Outcome of the authorisation request, and actions to perform.
	AuthorisationResult *AuthorisationResult10 `xml:"AuthstnRslt,omitempty"`

	// Result of performed verifications for the transaction.
	TransactionVerificationResult []*TransactionVerificationResult4 `xml:"TxVrfctnRslt,omitempty"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction60) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType5Code)(&value)
}

func (c *CardPaymentTransaction60) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType9Code)(&value))
}

func (c *CardPaymentTransaction60) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CardPaymentTransaction60) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction60) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction60) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction60) AddOriginalTransaction() *CardPaymentTransaction52 {
	c.OriginalTransaction = new(CardPaymentTransaction52)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction60) SetTransactionSuccess(value string) {
	c.TransactionSuccess = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction60) SetReversal(value string) {
	c.Reversal = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction60) SetMerchantOverride(value string) {
	c.MerchantOverride = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction60) AddFailureReason(value string) {
	c.FailureReason = append(c.FailureReason, (*FailureReason3Code)(&value))
}

func (c *CardPaymentTransaction60) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction60) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction60) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction60) SetInterchangeData(value string) {
	c.InterchangeData = (*Max140Text)(&value)
}

func (c *CardPaymentTransaction60) AddTransactionDetails() *CardPaymentTransactionDetails30 {
	c.TransactionDetails = new(CardPaymentTransactionDetails30)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction60) AddAuthorisationResult() *AuthorisationResult10 {
	c.AuthorisationResult = new(AuthorisationResult10)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction60) AddTransactionVerificationResult() *TransactionVerificationResult4 {
	newValue := new(TransactionVerificationResult4)
	c.TransactionVerificationResult = append(c.TransactionVerificationResult, newValue)
	return newValue
}

func (c *CardPaymentTransaction60) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}

// Transaction information in the cancellation advice.
type CardPaymentTransaction61 struct {

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd,omitempty"`

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction52 `xml:"OrgnlTx,omitempty"`

	// Outcome of the transaction at the acceptor.
	TransactionSuccess *TrueFalseIndicator `xml:"TxSucss"`

	// Notify that a previous transaction has to be reversed if this original transaction has been approved by the acquirer.
	Reversal *TrueFalseIndicator `xml:"Rvsl,omitempty"`

	// List of incidents during the transaction.
	FailureReason []*FailureReason3Code `xml:"FailrRsn,omitempty"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max140Text `xml:"IntrchngData,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails34 `xml:"TxDtls"`

	// Outcome of the authorisation.
	AuthorisationResult *AuthorisationResult12 `xml:"AuthstnRslt,omitempty"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction61) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction61) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction61) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction61) AddOriginalTransaction() *CardPaymentTransaction52 {
	c.OriginalTransaction = new(CardPaymentTransaction52)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction61) SetTransactionSuccess(value string) {
	c.TransactionSuccess = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction61) SetReversal(value string) {
	c.Reversal = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction61) AddFailureReason(value string) {
	c.FailureReason = append(c.FailureReason, (*FailureReason3Code)(&value))
}

func (c *CardPaymentTransaction61) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction61) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction61) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction61) SetInterchangeData(value string) {
	c.InterchangeData = (*Max140Text)(&value)
}

func (c *CardPaymentTransaction61) AddTransactionDetails() *CardPaymentTransactionDetails34 {
	c.TransactionDetails = new(CardPaymentTransactionDetails34)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction61) AddAuthorisationResult() *AuthorisationResult12 {
	c.AuthorisationResult = new(AuthorisationResult12)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction61) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}

// Data associated with the transaction to authorise.
type CardPaymentTransaction62 struct {

	// Flag indicating whether the transaction data must be captured or not in addition to the message process.
	TransactionCapture *TrueFalseIndicator `xml:"TxCaptr"`

	// Type of transaction being undertaken for the main service.
	TransactionType *CardPaymentServiceType5Code `xml:"TxTp,omitempty"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType9Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd,omitempty"`

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction52 `xml:"OrgnlTx,omitempty"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails31 `xml:"TxDtls"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction62) SetTransactionCapture(value string) {
	c.TransactionCapture = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction62) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType5Code)(&value)
}

func (c *CardPaymentTransaction62) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType9Code)(&value))
}

func (c *CardPaymentTransaction62) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CardPaymentTransaction62) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction62) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction62) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction62) AddOriginalTransaction() *CardPaymentTransaction52 {
	c.OriginalTransaction = new(CardPaymentTransaction52)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction62) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction62) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction62) AddTransactionDetails() *CardPaymentTransactionDetails31 {
	c.TransactionDetails = new(CardPaymentTransactionDetails31)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction62) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}

// Transaction information in the authorisation response.
type CardPaymentTransaction63 struct {

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max140Text `xml:"IntrchngData,omitempty"`

	// Detail of the transaction transported.
	TransactionDetails *CardPaymentTransactionDetails32 `xml:"TxDtls"`
}

func (c *CardPaymentTransaction63) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction63) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction63) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction63) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction63) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction63) SetInterchangeData(value string) {
	c.InterchangeData = (*Max140Text)(&value)
}

func (c *CardPaymentTransaction63) AddTransactionDetails() *CardPaymentTransactionDetails32 {
	c.TransactionDetails = new(CardPaymentTransactionDetails32)
	return c.TransactionDetails
}

// Data associated with the transaction for a potential currency conversion.
type CardPaymentTransaction64 struct {

	// Flag indicating whether the transaction data must be captured or not in addition to the message process.
	TransactionCapture *TrueFalseIndicator `xml:"TxCaptr,omitempty"`

	// Type of transaction being undertaken for the main service.
	TransactionType *CardPaymentServiceType5Code `xml:"TxTp"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType9Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd"`

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction65 `xml:"OrgnlTx,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails33 `xml:"TxDtls"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction64) SetTransactionCapture(value string) {
	c.TransactionCapture = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction64) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType5Code)(&value)
}

func (c *CardPaymentTransaction64) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType9Code)(&value))
}

func (c *CardPaymentTransaction64) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CardPaymentTransaction64) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction64) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction64) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction64) AddOriginalTransaction() *CardPaymentTransaction65 {
	c.OriginalTransaction = new(CardPaymentTransaction65)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction64) AddTransactionDetails() *CardPaymentTransactionDetails33 {
	c.TransactionDetails = new(CardPaymentTransactionDetails33)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction64) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}

// Identification of the original transaction.
type CardPaymentTransaction65 struct {

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the POI (Point Of Interaction) performing the transaction.
	POIIdentification *GenericIdentification32 `xml:"POIId"`

	// Link to a previous currency conversion.
	CurrencyConversion *CurrencyConversion7 `xml:"CcyConvs"`
}

func (c *CardPaymentTransaction65) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction65) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction65) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

func (c *CardPaymentTransaction65) AddCurrencyConversion() *CurrencyConversion7 {
	c.CurrencyConversion = new(CurrencyConversion7)
	return c.CurrencyConversion
}

// Transaction information in the cancellation advice.
type CardPaymentTransaction7 struct {

	// Flag indicating whether the transaction data must be captured or not in addition to the message process.
	TransactionCapture *TrueFalseIndicator `xml:"TxCaptr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction8 `xml:"OrgnlTx,omitempty"`

	// Outcome of the transaction at the acceptor.
	TransactionSuccess *TrueFalseIndicator `xml:"TxSucss"`

	// Notify that a previous transaction has to be reversed if this original transaction has been approved by the acquirer.
	Reversal *TrueFalseIndicator `xml:"Rvsl,omitempty"`

	// List of incidents during the transaction.
	FailureReason []*FailureReason1Code `xml:"FailrRsn,omitempty"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max35Text `xml:"IntrchngData,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails7 `xml:"TxDtls"`

	// Outcome of the authorisation.
	AuthorisationResult *AuthorisationResult2 `xml:"AuthstnRslt,omitempty"`

	// Additional information related to the transaction.
	AdditionalTransactionData *Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction7) SetTransactionCapture(value string) {
	c.TransactionCapture = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction7) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction7) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction7) AddOriginalTransaction() *CardPaymentTransaction8 {
	c.OriginalTransaction = new(CardPaymentTransaction8)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction7) SetTransactionSuccess(value string) {
	c.TransactionSuccess = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction7) SetReversal(value string) {
	c.Reversal = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction7) AddFailureReason(value string) {
	c.FailureReason = append(c.FailureReason, (*FailureReason1Code)(&value))
}

func (c *CardPaymentTransaction7) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction7) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction7) SetInterchangeData(value string) {
	c.InterchangeData = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction7) AddTransactionDetails() *CardPaymentTransactionDetails7 {
	c.TransactionDetails = new(CardPaymentTransactionDetails7)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction7) AddAuthorisationResult() *AuthorisationResult2 {
	c.AuthorisationResult = new(AuthorisationResult2)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction7) SetAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = (*Max70Text)(&value)
}

// Identification of the original transaction.
type CardPaymentTransaction8 struct {

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the POI (Point Of Interaction) performing the transaction.
	POIIdentification *GenericIdentification32 `xml:"POIId,omitempty"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Type of transaction being undertaken for the main service.
	TransactionType *CardPaymentServiceType1Code `xml:"TxTp"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType2Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Result of the original transaction.
	TransactionResult *CardPaymentTransactionResult1 `xml:"TxRslt,omitempty"`
}

func (c *CardPaymentTransaction8) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction8) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

func (c *CardPaymentTransaction8) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction8) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction8) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType1Code)(&value)
}

func (c *CardPaymentTransaction8) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType2Code)(&value))
}

func (c *CardPaymentTransaction8) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CardPaymentTransaction8) AddTransactionResult() *CardPaymentTransactionResult1 {
	c.TransactionResult = new(CardPaymentTransactionResult1)
	return c.TransactionResult
}

// Authorisation response from the acquirer.
type CardPaymentTransaction9 struct {

	// Outcome of the authorisation, and actions to perform.
	AuthorisationResult *AuthorisationResult1 `xml:"AuthstnRslt"`

	// Result of the verifications performed by the issuer to deliver or decline the authorisation.
	TransactionVerificationResult *TransactionVerificationResult1 `xml:"TxVrfctnRslt,omitempty"`

	// Balance of the account, related to the payment.
	Balance *ImpliedCurrencyAndAmount `xml:"Bal,omitempty"`

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Set of actions to be performed by the POI (Point Of Interaction) system.
	Action []*Action1 `xml:"Actn,omitempty"`
}

func (c *CardPaymentTransaction9) AddAuthorisationResult() *AuthorisationResult1 {
	c.AuthorisationResult = new(AuthorisationResult1)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction9) AddTransactionVerificationResult() *TransactionVerificationResult1 {
	c.TransactionVerificationResult = new(TransactionVerificationResult1)
	return c.TransactionVerificationResult
}

func (c *CardPaymentTransaction9) SetBalance(value, currency string) {
	c.Balance = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransaction9) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransaction9) AddAction() *Action1 {
	newValue := new(Action1)
	c.Action = append(c.Action, newValue)
	return newValue
}
