package iso20022

// Additional information about a request (e.g. financing request).
type AdditionalInformation1 struct {

	// Specifies the type of additional information.
	InformationType *InformationType1Choice `xml:"InfTp"`

	// Contents of the additional information.
	InformationValue *Max350Text `xml:"InfVal"`
}

func (a *AdditionalInformation1) AddInformationType() *InformationType1Choice {
	a.InformationType = new(InformationType1Choice)
	return a.InformationType
}

func (a *AdditionalInformation1) SetInformationValue(value string) {
	a.InformationValue = (*Max350Text)(&value)
}

// Additional specific modification criteria.
type AdditionalInformation11 struct {

	// Identification of the transaction as known by the account owner. Will be used in a unilateral split to provide the executing party with the account owner identification of each split transaction.
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId,omitempty"`

	// Type of instrument involved in the transactions on which the modification request should apply.
	ClassificationType *ClassificationType32Choice `xml:"ClssfctnTp,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount19 `xml:"SfkpgAcct,omitempty"`

	// Identification of the financial instrument involved in the transactions on which the modification request should apply.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId,omitempty"`

	// Quantity of financial instrument concerned by the settlement condition modification request.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty,omitempty"`

	// Date/time when the request should take effect.
	EffectiveDate *DateAndDateTimeChoice `xml:"FctvDt,omitempty"`

	// Date/time when the request should cease to be in effect.
	ExpiryDate *DateAndDateTimeChoice `xml:"XpryDt,omitempty"`

	// Date/time of the release.
	CutOffDate *DateAndDateTimeChoice `xml:"CutOffDt,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification100 `xml:"Invstr,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	DeliveringParty1 *PartyIdentificationAndAccount117 `xml:"DlvrgPty1,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	ReceivingParty1 *PartyIdentificationAndAccount117 `xml:"RcvgPty1,omitempty"`

	// Provides details on the processing status of the transaction.
	ProcessingStatus *ProcessingStatus56Choice `xml:"PrcgSts,omitempty"`
}

func (a *AdditionalInformation11) SetAccountOwnerTransactionIdentification(value string) {
	a.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalInformation11) AddClassificationType() *ClassificationType32Choice {
	a.ClassificationType = new(ClassificationType32Choice)
	return a.ClassificationType
}

func (a *AdditionalInformation11) AddSafekeepingAccount() *SecuritiesAccount19 {
	a.SafekeepingAccount = new(SecuritiesAccount19)
	return a.SafekeepingAccount
}

func (a *AdditionalInformation11) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	a.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return a.FinancialInstrumentIdentification
}

func (a *AdditionalInformation11) AddQuantity() *FinancialInstrumentQuantity1Choice {
	a.Quantity = new(FinancialInstrumentQuantity1Choice)
	return a.Quantity
}

func (a *AdditionalInformation11) AddEffectiveDate() *DateAndDateTimeChoice {
	a.EffectiveDate = new(DateAndDateTimeChoice)
	return a.EffectiveDate
}

func (a *AdditionalInformation11) AddExpiryDate() *DateAndDateTimeChoice {
	a.ExpiryDate = new(DateAndDateTimeChoice)
	return a.ExpiryDate
}

func (a *AdditionalInformation11) AddCutOffDate() *DateAndDateTimeChoice {
	a.CutOffDate = new(DateAndDateTimeChoice)
	return a.CutOffDate
}

func (a *AdditionalInformation11) AddInvestor() *PartyIdentification100 {
	a.Investor = new(PartyIdentification100)
	return a.Investor
}

func (a *AdditionalInformation11) AddDeliveringParty1() *PartyIdentificationAndAccount117 {
	a.DeliveringParty1 = new(PartyIdentificationAndAccount117)
	return a.DeliveringParty1
}

func (a *AdditionalInformation11) AddReceivingParty1() *PartyIdentificationAndAccount117 {
	a.ReceivingParty1 = new(PartyIdentificationAndAccount117)
	return a.ReceivingParty1
}

func (a *AdditionalInformation11) AddProcessingStatus() *ProcessingStatus56Choice {
	a.ProcessingStatus = new(ProcessingStatus56Choice)
	return a.ProcessingStatus
}

// Additional specific modification criteria.
type AdditionalInformation12 struct {

	// Identification of the transaction as known by the account owner. Will be used in a unilateral split to provide the executing party with the account owner identification of each split transaction.
	AccountOwnerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctOwnrTxId,omitempty"`

	// Type of instrument involved in the transactions on which the modification request should apply.
	ClassificationType *ClassificationType33Choice `xml:"ClssfctnTp,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount30 `xml:"SfkpgAcct,omitempty"`

	// Identification of the financial instrument involved in the transactions on which the modification request should apply.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId,omitempty"`

	// Quantity of financial instrument concerned by the settlement condition modification request.
	Quantity *FinancialInstrumentQuantity15Choice `xml:"Qty,omitempty"`

	// Date/time when the request should take effect.
	EffectiveDate *DateAndDateTimeChoice `xml:"FctvDt,omitempty"`

	// Date/time when the request should cease to be in effect.
	ExpiryDate *DateAndDateTimeChoice `xml:"XpryDt,omitempty"`

	// Date/time of the release.
	CutOffDate *DateAndDateTimeChoice `xml:"CutOffDt,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification111 `xml:"Invstr,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	DeliveringParty1 *PartyIdentificationAndAccount146 `xml:"DlvrgPty1,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	ReceivingParty1 *PartyIdentificationAndAccount146 `xml:"RcvgPty1,omitempty"`

	// Provides details on the processing status of the transaction.
	ProcessingStatus *ProcessingStatus59Choice `xml:"PrcgSts,omitempty"`
}

func (a *AdditionalInformation12) SetAccountOwnerTransactionIdentification(value string) {
	a.AccountOwnerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (a *AdditionalInformation12) AddClassificationType() *ClassificationType33Choice {
	a.ClassificationType = new(ClassificationType33Choice)
	return a.ClassificationType
}

func (a *AdditionalInformation12) AddSafekeepingAccount() *SecuritiesAccount30 {
	a.SafekeepingAccount = new(SecuritiesAccount30)
	return a.SafekeepingAccount
}

func (a *AdditionalInformation12) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	a.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return a.FinancialInstrumentIdentification
}

func (a *AdditionalInformation12) AddQuantity() *FinancialInstrumentQuantity15Choice {
	a.Quantity = new(FinancialInstrumentQuantity15Choice)
	return a.Quantity
}

func (a *AdditionalInformation12) AddEffectiveDate() *DateAndDateTimeChoice {
	a.EffectiveDate = new(DateAndDateTimeChoice)
	return a.EffectiveDate
}

func (a *AdditionalInformation12) AddExpiryDate() *DateAndDateTimeChoice {
	a.ExpiryDate = new(DateAndDateTimeChoice)
	return a.ExpiryDate
}

func (a *AdditionalInformation12) AddCutOffDate() *DateAndDateTimeChoice {
	a.CutOffDate = new(DateAndDateTimeChoice)
	return a.CutOffDate
}

func (a *AdditionalInformation12) AddInvestor() *PartyIdentification111 {
	a.Investor = new(PartyIdentification111)
	return a.Investor
}

func (a *AdditionalInformation12) AddDeliveringParty1() *PartyIdentificationAndAccount146 {
	a.DeliveringParty1 = new(PartyIdentificationAndAccount146)
	return a.DeliveringParty1
}

func (a *AdditionalInformation12) AddReceivingParty1() *PartyIdentificationAndAccount146 {
	a.ReceivingParty1 = new(PartyIdentificationAndAccount146)
	return a.ReceivingParty1
}

func (a *AdditionalInformation12) AddProcessingStatus() *ProcessingStatus59Choice {
	a.ProcessingStatus = new(ProcessingStatus59Choice)
	return a.ProcessingStatus
}

// Additional specific modification criteria.
type AdditionalInformation13 struct {

	// Identification of the transaction as known by the account owner. Will be used in a unilateral split to provide the executing party with the account owner identification of each split transaction.
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId,omitempty"`

	// Type of instrument involved in the transactions on which the modification request should apply.
	ClassificationType *ClassificationType32Choice `xml:"ClssfctnTp,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount19 `xml:"SfkpgAcct,omitempty"`

	// Identification of the financial instrument involved in the transactions on which the modification request should apply.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId,omitempty"`

	// Quantity of financial instrument concerned by the settlement condition modification request.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty,omitempty"`

	// Date/time when the request should take effect.
	EffectiveDate *DateAndDateTimeChoice `xml:"FctvDt,omitempty"`

	// Date/time when the request should cease to be in effect.
	ExpiryDate *DateAndDateTimeChoice `xml:"XpryDt,omitempty"`

	// Date/time of the release.
	CutOffDate *DateAndDateTimeChoice `xml:"CutOffDt,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification100 `xml:"Invstr,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	DeliveringParty1 *PartyIdentificationAndAccount117 `xml:"DlvrgPty1,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	ReceivingParty1 *PartyIdentificationAndAccount117 `xml:"RcvgPty1,omitempty"`
}

func (a *AdditionalInformation13) SetAccountOwnerTransactionIdentification(value string) {
	a.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalInformation13) AddClassificationType() *ClassificationType32Choice {
	a.ClassificationType = new(ClassificationType32Choice)
	return a.ClassificationType
}

func (a *AdditionalInformation13) AddSafekeepingAccount() *SecuritiesAccount19 {
	a.SafekeepingAccount = new(SecuritiesAccount19)
	return a.SafekeepingAccount
}

func (a *AdditionalInformation13) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	a.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return a.FinancialInstrumentIdentification
}

func (a *AdditionalInformation13) AddQuantity() *FinancialInstrumentQuantity1Choice {
	a.Quantity = new(FinancialInstrumentQuantity1Choice)
	return a.Quantity
}

func (a *AdditionalInformation13) AddEffectiveDate() *DateAndDateTimeChoice {
	a.EffectiveDate = new(DateAndDateTimeChoice)
	return a.EffectiveDate
}

func (a *AdditionalInformation13) AddExpiryDate() *DateAndDateTimeChoice {
	a.ExpiryDate = new(DateAndDateTimeChoice)
	return a.ExpiryDate
}

func (a *AdditionalInformation13) AddCutOffDate() *DateAndDateTimeChoice {
	a.CutOffDate = new(DateAndDateTimeChoice)
	return a.CutOffDate
}

func (a *AdditionalInformation13) AddInvestor() *PartyIdentification100 {
	a.Investor = new(PartyIdentification100)
	return a.Investor
}

func (a *AdditionalInformation13) AddDeliveringParty1() *PartyIdentificationAndAccount117 {
	a.DeliveringParty1 = new(PartyIdentificationAndAccount117)
	return a.DeliveringParty1
}

func (a *AdditionalInformation13) AddReceivingParty1() *PartyIdentificationAndAccount117 {
	a.ReceivingParty1 = new(PartyIdentificationAndAccount117)
	return a.ReceivingParty1
}

// Additional specific modification criteria.
type AdditionalInformation14 struct {

	// Identification of the transaction as known by the account owner. Will be used in a unilateral split to provide the executing party with the account owner identification of each split transaction.
	AccountOwnerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctOwnrTxId,omitempty"`

	// Type of instrument involved in the transactions on which the modification request should apply.
	ClassificationType *ClassificationType33Choice `xml:"ClssfctnTp,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount30 `xml:"SfkpgAcct,omitempty"`

	// Identification of the financial instrument involved in the transactions on which the modification request should apply.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId,omitempty"`

	// Quantity of financial instrument concerned by the settlement condition modification request.
	Quantity *FinancialInstrumentQuantity15Choice `xml:"Qty,omitempty"`

	// Date/time when the request should take effect.
	EffectiveDate *DateAndDateTimeChoice `xml:"FctvDt,omitempty"`

	// Date/time when the request should cease to be in effect.
	ExpiryDate *DateAndDateTimeChoice `xml:"XpryDt,omitempty"`

	// Date/time of the release.
	CutOffDate *DateAndDateTimeChoice `xml:"CutOffDt,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification111 `xml:"Invstr,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	DeliveringParty1 *PartyIdentificationAndAccount146 `xml:"DlvrgPty1,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	ReceivingParty1 *PartyIdentificationAndAccount146 `xml:"RcvgPty1,omitempty"`
}

func (a *AdditionalInformation14) SetAccountOwnerTransactionIdentification(value string) {
	a.AccountOwnerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (a *AdditionalInformation14) AddClassificationType() *ClassificationType33Choice {
	a.ClassificationType = new(ClassificationType33Choice)
	return a.ClassificationType
}

func (a *AdditionalInformation14) AddSafekeepingAccount() *SecuritiesAccount30 {
	a.SafekeepingAccount = new(SecuritiesAccount30)
	return a.SafekeepingAccount
}

func (a *AdditionalInformation14) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	a.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return a.FinancialInstrumentIdentification
}

func (a *AdditionalInformation14) AddQuantity() *FinancialInstrumentQuantity15Choice {
	a.Quantity = new(FinancialInstrumentQuantity15Choice)
	return a.Quantity
}

func (a *AdditionalInformation14) AddEffectiveDate() *DateAndDateTimeChoice {
	a.EffectiveDate = new(DateAndDateTimeChoice)
	return a.EffectiveDate
}

func (a *AdditionalInformation14) AddExpiryDate() *DateAndDateTimeChoice {
	a.ExpiryDate = new(DateAndDateTimeChoice)
	return a.ExpiryDate
}

func (a *AdditionalInformation14) AddCutOffDate() *DateAndDateTimeChoice {
	a.CutOffDate = new(DateAndDateTimeChoice)
	return a.CutOffDate
}

func (a *AdditionalInformation14) AddInvestor() *PartyIdentification111 {
	a.Investor = new(PartyIdentification111)
	return a.Investor
}

func (a *AdditionalInformation14) AddDeliveringParty1() *PartyIdentificationAndAccount146 {
	a.DeliveringParty1 = new(PartyIdentificationAndAccount146)
	return a.DeliveringParty1
}

func (a *AdditionalInformation14) AddReceivingParty1() *PartyIdentificationAndAccount146 {
	a.ReceivingParty1 = new(PartyIdentificationAndAccount146)
	return a.ReceivingParty1
}

// Additional specific modification criteria.
type AdditionalInformation3 struct {

	// Identification of the transaction as known by the account owner. Will be used in a unilateral split to provide the executing party with the account owner identification of each split transaction.
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId,omitempty"`

	// Type of instrument involved in the transactions on which the modification request should apply.
	ClassificationType *ClassificationType1Choice `xml:"ClssfctnTp,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct,omitempty"`

	// Identification of the financial instrument involved in the transactions on which the modification request should apply.
	FinancialInstrumentIdentification *SecurityIdentification11 `xml:"FinInstrmId,omitempty"`

	// Quantity of financial instrument concerned by the settlement condition modification request.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty,omitempty"`

	// Date/time when the request should take effect.
	EffectiveDate *DateAndDateTimeChoice `xml:"FctvDt,omitempty"`

	// Date/time when the request should cease to be in effect.
	ExpiryDate *DateAndDateTimeChoice `xml:"XpryDt,omitempty"`

	// Date/time of the release.
	CutOffDate *DateAndDateTimeChoice `xml:"CutOffDt,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification10Choice `xml:"Invstr,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	DeliveringParty1 *PartyIdentificationAndAccount16 `xml:"DlvrgPty1,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	ReceivingParty1 *PartyIdentificationAndAccount16 `xml:"RcvgPty1,omitempty"`
}

func (a *AdditionalInformation3) SetAccountOwnerTransactionIdentification(value string) {
	a.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalInformation3) AddClassificationType() *ClassificationType1Choice {
	a.ClassificationType = new(ClassificationType1Choice)
	return a.ClassificationType
}

func (a *AdditionalInformation3) AddSafekeepingAccount() *SecuritiesAccount13 {
	a.SafekeepingAccount = new(SecuritiesAccount13)
	return a.SafekeepingAccount
}

func (a *AdditionalInformation3) AddFinancialInstrumentIdentification() *SecurityIdentification11 {
	a.FinancialInstrumentIdentification = new(SecurityIdentification11)
	return a.FinancialInstrumentIdentification
}

func (a *AdditionalInformation3) AddQuantity() *FinancialInstrumentQuantity1Choice {
	a.Quantity = new(FinancialInstrumentQuantity1Choice)
	return a.Quantity
}

func (a *AdditionalInformation3) AddEffectiveDate() *DateAndDateTimeChoice {
	a.EffectiveDate = new(DateAndDateTimeChoice)
	return a.EffectiveDate
}

func (a *AdditionalInformation3) AddExpiryDate() *DateAndDateTimeChoice {
	a.ExpiryDate = new(DateAndDateTimeChoice)
	return a.ExpiryDate
}

func (a *AdditionalInformation3) AddCutOffDate() *DateAndDateTimeChoice {
	a.CutOffDate = new(DateAndDateTimeChoice)
	return a.CutOffDate
}

func (a *AdditionalInformation3) AddInvestor() *PartyIdentification10Choice {
	a.Investor = new(PartyIdentification10Choice)
	return a.Investor
}

func (a *AdditionalInformation3) AddDeliveringParty1() *PartyIdentificationAndAccount16 {
	a.DeliveringParty1 = new(PartyIdentificationAndAccount16)
	return a.DeliveringParty1
}

func (a *AdditionalInformation3) AddReceivingParty1() *PartyIdentificationAndAccount16 {
	a.ReceivingParty1 = new(PartyIdentificationAndAccount16)
	return a.ReceivingParty1
}

// Contains additional information related to the message.
type AdditionalInformation5 struct {

	// Contains additional information related to the message.
	Information []*Max256Text `xml:"Inf"`
}

func (a *AdditionalInformation5) AddInformation(value string) {
	a.Information = append(a.Information, (*Max256Text)(&value))
}

// Additional information about a request.
type AdditionalInformation6 struct {

	// Specifies the type of additional information.
	InformationType *ExternalInformationType1Code `xml:"InfTp"`

	// Contents of the additional information.
	InformationValue *Max350Text `xml:"InfVal"`
}

func (a *AdditionalInformation6) SetInformationType(value string) {
	a.InformationType = (*ExternalInformationType1Code)(&value)
}

func (a *AdditionalInformation6) SetInformationValue(value string) {
	a.InformationValue = (*Max350Text)(&value)
}

// Additional specific modification criteria.
type AdditionalInformation7 struct {

	// Identification of the transaction as known by the account owner. Will be used in a unilateral split to provide the executing party with the account owner identification of each split transaction.
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId,omitempty"`

	// Type of instrument involved in the transactions on which the modification request should apply.
	ClassificationType *ClassificationType1Choice `xml:"ClssfctnTp,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct,omitempty"`

	// Identification of the financial instrument involved in the transactions on which the modification request should apply.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId,omitempty"`

	// Quantity of financial instrument concerned by the settlement condition modification request.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty,omitempty"`

	// Date/time when the request should take effect.
	EffectiveDate *DateAndDateTimeChoice `xml:"FctvDt,omitempty"`

	// Date/time when the request should cease to be in effect.
	ExpiryDate *DateAndDateTimeChoice `xml:"XpryDt,omitempty"`

	// Date/time of the release.
	CutOffDate *DateAndDateTimeChoice `xml:"CutOffDt,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification43Choice `xml:"Invstr,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	DeliveringParty1 *PartyIdentificationAndAccount43 `xml:"DlvrgPty1,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	ReceivingParty1 *PartyIdentificationAndAccount43 `xml:"RcvgPty1,omitempty"`
}

func (a *AdditionalInformation7) SetAccountOwnerTransactionIdentification(value string) {
	a.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalInformation7) AddClassificationType() *ClassificationType1Choice {
	a.ClassificationType = new(ClassificationType1Choice)
	return a.ClassificationType
}

func (a *AdditionalInformation7) AddSafekeepingAccount() *SecuritiesAccount13 {
	a.SafekeepingAccount = new(SecuritiesAccount13)
	return a.SafekeepingAccount
}

func (a *AdditionalInformation7) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	a.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return a.FinancialInstrumentIdentification
}

func (a *AdditionalInformation7) AddQuantity() *FinancialInstrumentQuantity1Choice {
	a.Quantity = new(FinancialInstrumentQuantity1Choice)
	return a.Quantity
}

func (a *AdditionalInformation7) AddEffectiveDate() *DateAndDateTimeChoice {
	a.EffectiveDate = new(DateAndDateTimeChoice)
	return a.EffectiveDate
}

func (a *AdditionalInformation7) AddExpiryDate() *DateAndDateTimeChoice {
	a.ExpiryDate = new(DateAndDateTimeChoice)
	return a.ExpiryDate
}

func (a *AdditionalInformation7) AddCutOffDate() *DateAndDateTimeChoice {
	a.CutOffDate = new(DateAndDateTimeChoice)
	return a.CutOffDate
}

func (a *AdditionalInformation7) AddInvestor() *PartyIdentification43Choice {
	a.Investor = new(PartyIdentification43Choice)
	return a.Investor
}

func (a *AdditionalInformation7) AddDeliveringParty1() *PartyIdentificationAndAccount43 {
	a.DeliveringParty1 = new(PartyIdentificationAndAccount43)
	return a.DeliveringParty1
}

func (a *AdditionalInformation7) AddReceivingParty1() *PartyIdentificationAndAccount43 {
	a.ReceivingParty1 = new(PartyIdentificationAndAccount43)
	return a.ReceivingParty1
}

// Additional specific modification criteria.
type AdditionalInformation9 struct {

	// Identification of the transaction as known by the account owner. Will be used in a unilateral split to provide the executing party with the account owner identification of each split transaction.
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId,omitempty"`

	// Type of instrument involved in the transactions on which the modification request should apply.
	ClassificationType *ClassificationType1Choice `xml:"ClssfctnTp,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct,omitempty"`

	// Identification of the financial instrument involved in the transactions on which the modification request should apply.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId,omitempty"`

	// Quantity of financial instrument concerned by the settlement condition modification request.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty,omitempty"`

	// Date/time when the request should take effect.
	EffectiveDate *DateAndDateTimeChoice `xml:"FctvDt,omitempty"`

	// Date/time when the request should cease to be in effect.
	ExpiryDate *DateAndDateTimeChoice `xml:"XpryDt,omitempty"`

	// Date/time of the release.
	CutOffDate *DateAndDateTimeChoice `xml:"CutOffDt,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification43Choice `xml:"Invstr,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	DeliveringParty1 *PartyIdentificationAndAccount43 `xml:"DlvrgPty1,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	ReceivingParty1 *PartyIdentificationAndAccount43 `xml:"RcvgPty1,omitempty"`

	// Provides details on the processing status of the transaction.
	ProcessingStatus *ProcessingStatus39Choice `xml:"PrcgSts,omitempty"`
}

func (a *AdditionalInformation9) SetAccountOwnerTransactionIdentification(value string) {
	a.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalInformation9) AddClassificationType() *ClassificationType1Choice {
	a.ClassificationType = new(ClassificationType1Choice)
	return a.ClassificationType
}

func (a *AdditionalInformation9) AddSafekeepingAccount() *SecuritiesAccount13 {
	a.SafekeepingAccount = new(SecuritiesAccount13)
	return a.SafekeepingAccount
}

func (a *AdditionalInformation9) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	a.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return a.FinancialInstrumentIdentification
}

func (a *AdditionalInformation9) AddQuantity() *FinancialInstrumentQuantity1Choice {
	a.Quantity = new(FinancialInstrumentQuantity1Choice)
	return a.Quantity
}

func (a *AdditionalInformation9) AddEffectiveDate() *DateAndDateTimeChoice {
	a.EffectiveDate = new(DateAndDateTimeChoice)
	return a.EffectiveDate
}

func (a *AdditionalInformation9) AddExpiryDate() *DateAndDateTimeChoice {
	a.ExpiryDate = new(DateAndDateTimeChoice)
	return a.ExpiryDate
}

func (a *AdditionalInformation9) AddCutOffDate() *DateAndDateTimeChoice {
	a.CutOffDate = new(DateAndDateTimeChoice)
	return a.CutOffDate
}

func (a *AdditionalInformation9) AddInvestor() *PartyIdentification43Choice {
	a.Investor = new(PartyIdentification43Choice)
	return a.Investor
}

func (a *AdditionalInformation9) AddDeliveringParty1() *PartyIdentificationAndAccount43 {
	a.DeliveringParty1 = new(PartyIdentificationAndAccount43)
	return a.DeliveringParty1
}

func (a *AdditionalInformation9) AddReceivingParty1() *PartyIdentificationAndAccount43 {
	a.ReceivingParty1 = new(PartyIdentificationAndAccount43)
	return a.ReceivingParty1
}

func (a *AdditionalInformation9) AddProcessingStatus() *ProcessingStatus39Choice {
	a.ProcessingStatus = new(ProcessingStatus39Choice)
	return a.ProcessingStatus
}
