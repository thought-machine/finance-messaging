package iso20022

// Provides information about the securities account.
type SecuritiesAccount10 struct {

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Identification of the party that owns the account.
	AccountOwnerIdentification *PartyIdentification2Choice `xml:"AcctOwnrId,omitempty"`

	// Nationality of the account owner.
	AccountOwnerNationality *NationalityCode `xml:"AcctOwnrNtlty,omitempty"`

	// Idenfitication of the account.
	AccountIdentification *Max35Text `xml:"AcctId"`

	// Type of balance.
	BalanceType *SecuritiesBalanceType9FormatChoice `xml:"BalTp,omitempty"`

	// Specifies the form of the financial instrument.
	SecurityHoldingForm *FormOfSecurity1Code `xml:"SctyHldgForm,omitempty"`
}

func (s *SecuritiesAccount10) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesAccount10) AddAccountOwnerIdentification() *PartyIdentification2Choice {
	s.AccountOwnerIdentification = new(PartyIdentification2Choice)
	return s.AccountOwnerIdentification
}

func (s *SecuritiesAccount10) SetAccountOwnerNationality(value string) {
	s.AccountOwnerNationality = (*NationalityCode)(&value)
}

func (s *SecuritiesAccount10) SetAccountIdentification(value string) {
	s.AccountIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesAccount10) AddBalanceType() *SecuritiesBalanceType9FormatChoice {
	s.BalanceType = new(SecuritiesBalanceType9FormatChoice)
	return s.BalanceType
}

func (s *SecuritiesAccount10) SetSecurityHoldingForm(value string) {
	s.SecurityHoldingForm = (*FormOfSecurity1Code)(&value)
}

// Account to or from which a securities entry is made.
type SecuritiesAccount11 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *Max35Text `xml:"Id"`

	// Specifies the type of securities account.
	Type *PurposeCode1Choice `xml:"Tp,omitempty"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max70Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`
}

func (s *SecuritiesAccount11) SetIdentification(value string) {
	s.Identification = (*Max35Text)(&value)
}

func (s *SecuritiesAccount11) AddType() *PurposeCode1Choice {
	s.Type = new(PurposeCode1Choice)
	return s.Type
}

func (s *SecuritiesAccount11) SetName(value string) {
	s.Name = (*Max70Text)(&value)
}

func (s *SecuritiesAccount11) SetDesignation(value string) {
	s.Designation = (*Max35Text)(&value)
}

// Provides information about the securities account.
type SecuritiesAccount12 struct {

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Identification of the party that owns the account.
	AccountOwnerIdentification *PartyIdentification2Choice `xml:"AcctOwnrId,omitempty"`

	// Idenfitication of the account where financial instruments are maintained.
	AccountIdentification *Max35Text `xml:"AcctId"`

	// Type of balance.
	BalanceType *SecuritiesBalanceType6FormatChoice `xml:"BalTp,omitempty"`

	// Specifies the form of the financial instrument.
	SecurityHoldingForm *FormOfSecurity1Code `xml:"SctyHldgForm,omitempty"`
}

func (s *SecuritiesAccount12) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesAccount12) AddAccountOwnerIdentification() *PartyIdentification2Choice {
	s.AccountOwnerIdentification = new(PartyIdentification2Choice)
	return s.AccountOwnerIdentification
}

func (s *SecuritiesAccount12) SetAccountIdentification(value string) {
	s.AccountIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesAccount12) AddBalanceType() *SecuritiesBalanceType6FormatChoice {
	s.BalanceType = new(SecuritiesBalanceType6FormatChoice)
	return s.BalanceType
}

func (s *SecuritiesAccount12) SetSecurityHoldingForm(value string) {
	s.SecurityHoldingForm = (*FormOfSecurity1Code)(&value)
}

// Account to or from which a securities entry is made.
type SecuritiesAccount13 struct {

	// Unambiguous identification for the account between the account owner and the account servicer.”
	Identification *Max35Text `xml:"Id"`

	// Specifies the type of securities account.
	Type *GenericIdentification20 `xml:"Tp,omitempty"`

	// Description of the account.
	Name *Max70Text `xml:"Nm,omitempty"`
}

func (s *SecuritiesAccount13) SetIdentification(value string) {
	s.Identification = (*Max35Text)(&value)
}

func (s *SecuritiesAccount13) AddType() *GenericIdentification20 {
	s.Type = new(GenericIdentification20)
	return s.Type
}

func (s *SecuritiesAccount13) SetName(value string) {
	s.Name = (*Max70Text)(&value)
}

// Account to or from which a securities entry is made.
type SecuritiesAccount14 struct {

	// Unambiguous identification for the account between the account owner and the account servicer.”
	Identification *Max35Text `xml:"Id"`

	// Specifies the type of securities account.
	Type *PurposeCode2Choice `xml:"Tp,omitempty"`

	// Description of the account.
	Name *Max70Text `xml:"Nm,omitempty"`
}

func (s *SecuritiesAccount14) SetIdentification(value string) {
	s.Identification = (*Max35Text)(&value)
}

func (s *SecuritiesAccount14) AddType() *PurposeCode2Choice {
	s.Type = new(PurposeCode2Choice)
	return s.Type
}

func (s *SecuritiesAccount14) SetName(value string) {
	s.Name = (*Max70Text)(&value)
}

// Account to or from which a securities entry is made.
type SecuritiesAccount18 struct {

	// Unambiguous identification for the account between the account owner and the account servicer.
	Identification *Max35Text `xml:"Id"`

	// Specifies if the account is a House, a Client or a Liquidity Provider (Market Maker) account.
	Type *ClearingAccountType1Code `xml:"Tp"`

	// Description of the account.
	Name *Max70Text `xml:"Nm,omitempty"`
}

func (s *SecuritiesAccount18) SetIdentification(value string) {
	s.Identification = (*Max35Text)(&value)
}

func (s *SecuritiesAccount18) SetType(value string) {
	s.Type = (*ClearingAccountType1Code)(&value)
}

func (s *SecuritiesAccount18) SetName(value string) {
	s.Name = (*Max70Text)(&value)
}

// Account to or from which a securities entry is made.
type SecuritiesAccount19 struct {

	// Unambiguous identification for the account between the account owner and the account servicer.
	Identification *Max35Text `xml:"Id"`

	// Specifies the type of securities account.
	Type *GenericIdentification30 `xml:"Tp,omitempty"`

	// Description of the account.
	Name *Max70Text `xml:"Nm,omitempty"`
}

func (s *SecuritiesAccount19) SetIdentification(value string) {
	s.Identification = (*Max35Text)(&value)
}

func (s *SecuritiesAccount19) AddType() *GenericIdentification30 {
	s.Type = new(GenericIdentification30)
	return s.Type
}

func (s *SecuritiesAccount19) SetName(value string) {
	s.Name = (*Max70Text)(&value)
}

// Number assigned by a government agency to identify foreign nationals.
type SecuritiesAccount20 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *Max35Text `xml:"Id"`

	// Choice between a code and a data source scheme to identify the type of account.
	Type *ClearingAccountType1Code `xml:"Tp"`

	// .
	Name *Max70Text `xml:"Nm,omitempty"`
}

func (s *SecuritiesAccount20) SetIdentification(value string) {
	s.Identification = (*Max35Text)(&value)
}

func (s *SecuritiesAccount20) SetType(value string) {
	s.Type = (*ClearingAccountType1Code)(&value)
}

func (s *SecuritiesAccount20) SetName(value string) {
	s.Name = (*Max70Text)(&value)
}

// Unambiguous identification for the account between the account owner and the account servicer.
type SecuritiesAccount21 struct {

	// Account identification.
	Account *AccountIdentification5 `xml:"Acct"`

	// Sub-account identification.
	SubAccount *AccountIdentification5 `xml:"SubAcct,omitempty"`

	// Base currency for the account.
	BaseCurrency *ActiveOrHistoricCurrencyCode `xml:"BaseCcy,omitempty"`

	// Currency chosen for reporting purposes by the account owner in agreement with the account servicer.
	ReportingCurrency *ActiveOrHistoricCurrencyCode `xml:"RptgCcy,omitempty"`

	// Foreign exchange rate applied between the reporting and base currencies. It is assumed the valuation date is the same as the report date.
	ForeignExchangeRate *BaseOneRate `xml:"FXRate,omitempty"`
}

func (s *SecuritiesAccount21) AddAccount() *AccountIdentification5 {
	s.Account = new(AccountIdentification5)
	return s.Account
}

func (s *SecuritiesAccount21) AddSubAccount() *AccountIdentification5 {
	s.SubAccount = new(AccountIdentification5)
	return s.SubAccount
}

func (s *SecuritiesAccount21) SetBaseCurrency(value string) {
	s.BaseCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (s *SecuritiesAccount21) SetReportingCurrency(value string) {
	s.ReportingCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (s *SecuritiesAccount21) SetForeignExchangeRate(value string) {
	s.ForeignExchangeRate = (*BaseOneRate)(&value)
}

// Account to or from which a securities entry is made.
type SecuritiesAccount22 struct {

	// Unambiguous identification for the account between the account owner and the account servicer.
	Identification *Max35Text `xml:"Id"`

	// Type of securities account.
	Type *GenericIdentification30 `xml:"Tp,omitempty"`

	// Description of the account.
	Name *Max70Text `xml:"Nm,omitempty"`
}

func (s *SecuritiesAccount22) SetIdentification(value string) {
	s.Identification = (*Max35Text)(&value)
}

func (s *SecuritiesAccount22) AddType() *GenericIdentification30 {
	s.Type = new(GenericIdentification30)
	return s.Type
}

func (s *SecuritiesAccount22) SetName(value string) {
	s.Name = (*Max70Text)(&value)
}

// Account to or from which a securities entry is made.
type SecuritiesAccount24 struct {

	// Unambiguous identification for the account between the account owner and the account servicer.”
	Identification *Max35Text `xml:"Id"`

	// Specifies the type of securities account.
	Type *GenericIdentification30 `xml:"Tp,omitempty"`

	// Description of the account.
	Name *Max70Text `xml:"Nm,omitempty"`
}

func (s *SecuritiesAccount24) SetIdentification(value string) {
	s.Identification = (*Max35Text)(&value)
}

func (s *SecuritiesAccount24) AddType() *GenericIdentification30 {
	s.Type = new(GenericIdentification30)
	return s.Type
}

func (s *SecuritiesAccount24) SetName(value string) {
	s.Name = (*Max70Text)(&value)
}

// Account to or from which a securities entry is made.
type SecuritiesAccount25 struct {

	// Unambiguous identification for the account between the account owner and the account servicer.”
	Identification *Max35Text `xml:"Id"`

	// Specifies the type of securities account.
	Type *PurposeCode7Choice `xml:"Tp,omitempty"`

	// Description of the account.
	Name *Max70Text `xml:"Nm,omitempty"`
}

func (s *SecuritiesAccount25) SetIdentification(value string) {
	s.Identification = (*Max35Text)(&value)
}

func (s *SecuritiesAccount25) AddType() *PurposeCode7Choice {
	s.Type = new(PurposeCode7Choice)
	return s.Type
}

func (s *SecuritiesAccount25) SetName(value string) {
	s.Name = (*Max70Text)(&value)
}

// Account to or from which a securities entry is made.
type SecuritiesAccount26 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *Max35Text `xml:"Id"`

	// Specifies the type of securities account.
	Type *PurposeCode7Choice `xml:"Tp,omitempty"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max70Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`
}

func (s *SecuritiesAccount26) SetIdentification(value string) {
	s.Identification = (*Max35Text)(&value)
}

func (s *SecuritiesAccount26) AddType() *PurposeCode7Choice {
	s.Type = new(PurposeCode7Choice)
	return s.Type
}

func (s *SecuritiesAccount26) SetName(value string) {
	s.Name = (*Max70Text)(&value)
}

func (s *SecuritiesAccount26) SetDesignation(value string) {
	s.Designation = (*Max35Text)(&value)
}

// Account to or from which a securities entry is made.
type SecuritiesAccount27 struct {

	// Unambiguous identification for the account between the account owner and the account servicer.”
	Identification *RestrictedFINXMax35Text `xml:"Id"`

	// Specifies the type of securities account.
	Type *GenericIdentification47 `xml:"Tp,omitempty"`

	// Description of the account.
	Name *Max70Text `xml:"Nm,omitempty"`
}

func (s *SecuritiesAccount27) SetIdentification(value string) {
	s.Identification = (*RestrictedFINXMax35Text)(&value)
}

func (s *SecuritiesAccount27) AddType() *GenericIdentification47 {
	s.Type = new(GenericIdentification47)
	return s.Type
}

func (s *SecuritiesAccount27) SetName(value string) {
	s.Name = (*Max70Text)(&value)
}

// Number assigned by a government agency to identify foreign nationals.
type SecuritiesAccount3 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *Max35Text `xml:"Id"`

	// Choice between a code and a data source scheme to identify the type of account.
	Type *PurposeCode5Choice `xml:"Tp,omitempty"`

	// .
	Name *Max70Text `xml:"Nm,omitempty"`
}

func (s *SecuritiesAccount3) SetIdentification(value string) {
	s.Identification = (*Max35Text)(&value)
}

func (s *SecuritiesAccount3) AddType() *PurposeCode5Choice {
	s.Type = new(PurposeCode5Choice)
	return s.Type
}

func (s *SecuritiesAccount3) SetName(value string) {
	s.Name = (*Max70Text)(&value)
}

// Account to or from which a securities entry is made.
type SecuritiesAccount30 struct {

	// Unambiguous identification for the account between the account owner and the account servicer.
	Identification *RestrictedFINXMax35Text `xml:"Id"`

	// Specifies the type of securities account.
	Type *GenericIdentification47 `xml:"Tp,omitempty"`

	// Description of the account.
	Name *Max70Text `xml:"Nm,omitempty"`
}

func (s *SecuritiesAccount30) SetIdentification(value string) {
	s.Identification = (*RestrictedFINXMax35Text)(&value)
}

func (s *SecuritiesAccount30) AddType() *GenericIdentification47 {
	s.Type = new(GenericIdentification47)
	return s.Type
}

func (s *SecuritiesAccount30) SetName(value string) {
	s.Name = (*Max70Text)(&value)
}

// Account to or from which a securities entry is made.
type SecuritiesAccount33 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *RestrictedFINXMax35Text `xml:"Id"`

	// Specifies the type of securities account.
	Type *PurposeCode8Choice `xml:"Tp,omitempty"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max70Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`
}

func (s *SecuritiesAccount33) SetIdentification(value string) {
	s.Identification = (*RestrictedFINXMax35Text)(&value)
}

func (s *SecuritiesAccount33) AddType() *PurposeCode8Choice {
	s.Type = new(PurposeCode8Choice)
	return s.Type
}

func (s *SecuritiesAccount33) SetName(value string) {
	s.Name = (*Max70Text)(&value)
}

func (s *SecuritiesAccount33) SetDesignation(value string) {
	s.Designation = (*Max35Text)(&value)
}

// Account to or from which a securities entry is made.
type SecuritiesAccount34 struct {

	// Unambiguous identification for the account between the account owner and the account servicer.”
	Identification *RestrictedFINXMax35Text `xml:"Id"`

	// Specifies the type of securities account.
	Type *PurposeCode8Choice `xml:"Tp,omitempty"`

	// Description of the account.
	Name *Max70Text `xml:"Nm,omitempty"`
}

func (s *SecuritiesAccount34) SetIdentification(value string) {
	s.Identification = (*RestrictedFINXMax35Text)(&value)
}

func (s *SecuritiesAccount34) AddType() *PurposeCode8Choice {
	s.Type = new(PurposeCode8Choice)
	return s.Type
}

func (s *SecuritiesAccount34) SetName(value string) {
	s.Name = (*Max70Text)(&value)
}

// Account to or from which a securities entry is made.
type SecuritiesAccount4 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification1 `xml:"Id"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm"`

	// Additional information about a financial instrument to help identify the instrument.
	FinancialInstrumentSupplementaryIdentification *Max35Text `xml:"FinInstrmSplmtryId,omitempty"`

	// Identification of a security, as assigned under a formal or proprietary identification scheme.
	FinancialInstrumentIdentification *SecurityIdentification3Choice `xml:"FinInstrmId,omitempty"`

	// Name of the financial instrument in free format text.
	FinancialInstrumentName *Max350Text `xml:"FinInstrmNm,omitempty"`

	// Specifies the current state of an account, eg, enabled or deleted.
	Status *AccountStatus1Code `xml:"Sts"`
}

func (s *SecuritiesAccount4) AddIdentification() *AccountIdentification1 {
	s.Identification = new(AccountIdentification1)
	return s.Identification
}

func (s *SecuritiesAccount4) SetName(value string) {
	s.Name = (*Max35Text)(&value)
}

func (s *SecuritiesAccount4) SetFinancialInstrumentSupplementaryIdentification(value string) {
	s.FinancialInstrumentSupplementaryIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesAccount4) AddFinancialInstrumentIdentification() *SecurityIdentification3Choice {
	s.FinancialInstrumentIdentification = new(SecurityIdentification3Choice)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesAccount4) SetFinancialInstrumentName(value string) {
	s.FinancialInstrumentName = (*Max350Text)(&value)
}

func (s *SecuritiesAccount4) SetStatus(value string) {
	s.Status = (*AccountStatus1Code)(&value)
}

// Provides information about the securities account.
type SecuritiesAccount6 struct {

	// Identification of the financial instrument.
	SecurityIdentification *SecurityIdentification7 `xml:"SctyId"`

	// Idenfitication of the account where financial instruments are maintained.
	SecuritiesAccountIdentification *Max35Text `xml:"SctiesAcctId"`

	// Identification of the party that owns the account.
	AccountOwnerIdentification *PartyIdentification2Choice `xml:"AcctOwnrId,omitempty"`

	// Identification of the place of safekeeping.
	SafekeepingPlace *PartyIdentification2Choice `xml:"SfkpgPlc"`

	// Provides information required for the registration of the security.
	RegistrationDetails *Max350Text `xml:"RegnDtls,omitempty"`
}

func (s *SecuritiesAccount6) AddSecurityIdentification() *SecurityIdentification7 {
	s.SecurityIdentification = new(SecurityIdentification7)
	return s.SecurityIdentification
}

func (s *SecuritiesAccount6) SetSecuritiesAccountIdentification(value string) {
	s.SecuritiesAccountIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesAccount6) AddAccountOwnerIdentification() *PartyIdentification2Choice {
	s.AccountOwnerIdentification = new(PartyIdentification2Choice)
	return s.AccountOwnerIdentification
}

func (s *SecuritiesAccount6) AddSafekeepingPlace() *PartyIdentification2Choice {
	s.SafekeepingPlace = new(PartyIdentification2Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesAccount6) SetRegistrationDetails(value string) {
	s.RegistrationDetails = (*Max350Text)(&value)
}

// Provides information about the account identification and the account owner.
type SecuritiesAccount7 struct {

	// Identification of the party that owns the account.
	AccountOwnerIdentification *PartyIdentification2Choice `xml:"AcctOwnrId,omitempty"`

	// Idenfitication of the account where financial instruments are maintained.
	AccountIdentification *Max35Text `xml:"AcctId"`
}

func (s *SecuritiesAccount7) AddAccountOwnerIdentification() *PartyIdentification2Choice {
	s.AccountOwnerIdentification = new(PartyIdentification2Choice)
	return s.AccountOwnerIdentification
}

func (s *SecuritiesAccount7) SetAccountIdentification(value string) {
	s.AccountIdentification = (*Max35Text)(&value)
}

// Provides information about the securities account.
type SecuritiesAccount8 struct {

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Identification of the party that owns the account.
	AccountOwnerIdentification *PartyIdentification2Choice `xml:"AcctOwnrId,omitempty"`

	// Idenfitication of the account where financial instruments are maintained.
	AccountIdentification *Max35Text `xml:"AcctId"`

	// Type of balance.
	BalanceType *SecuritiesBalanceType10FormatChoice `xml:"BalTp,omitempty"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption1FormatChoice `xml:"OptnTp,omitempty"`

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb,omitempty"`

	// Specifies the form of the financial instrument.
	SecurityHoldingForm *FormOfSecurity1Code `xml:"SctyHldgForm,omitempty"`

	// Specifies if the stamp duty is applicable.
	StampDuty *StampDutyType1FormatChoice `xml:"StmpDty,omitempty"`
}

func (s *SecuritiesAccount8) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesAccount8) AddAccountOwnerIdentification() *PartyIdentification2Choice {
	s.AccountOwnerIdentification = new(PartyIdentification2Choice)
	return s.AccountOwnerIdentification
}

func (s *SecuritiesAccount8) SetAccountIdentification(value string) {
	s.AccountIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesAccount8) AddBalanceType() *SecuritiesBalanceType10FormatChoice {
	s.BalanceType = new(SecuritiesBalanceType10FormatChoice)
	return s.BalanceType
}

func (s *SecuritiesAccount8) AddOptionType() *CorporateActionOption1FormatChoice {
	s.OptionType = new(CorporateActionOption1FormatChoice)
	return s.OptionType
}

func (s *SecuritiesAccount8) SetOptionNumber(value string) {
	s.OptionNumber = (*Exact3NumericText)(&value)
}

func (s *SecuritiesAccount8) SetSecurityHoldingForm(value string) {
	s.SecurityHoldingForm = (*FormOfSecurity1Code)(&value)
}

func (s *SecuritiesAccount8) AddStampDuty() *StampDutyType1FormatChoice {
	s.StampDuty = new(StampDutyType1FormatChoice)
	return s.StampDuty
}

// Provides information about the securities account.
type SecuritiesAccount9 struct {

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Identification of the party that owns the account.
	AccountOwnerIdentification *PartyIdentification2Choice `xml:"AcctOwnrId,omitempty"`

	// Idenfitication of the account.
	AccountIdentification *Max35Text `xml:"AcctId"`

	// Type of balance.
	BalanceType *SecuritiesBalanceType10FormatChoice `xml:"BalTp,omitempty"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption1FormatChoice `xml:"OptnTp,omitempty"`

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb,omitempty"`

	// Specifies the form of the financial instrument.
	SecurityHoldingForm *FormOfSecurity1Code `xml:"SctyHldgForm,omitempty"`
}

func (s *SecuritiesAccount9) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesAccount9) AddAccountOwnerIdentification() *PartyIdentification2Choice {
	s.AccountOwnerIdentification = new(PartyIdentification2Choice)
	return s.AccountOwnerIdentification
}

func (s *SecuritiesAccount9) SetAccountIdentification(value string) {
	s.AccountIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesAccount9) AddBalanceType() *SecuritiesBalanceType10FormatChoice {
	s.BalanceType = new(SecuritiesBalanceType10FormatChoice)
	return s.BalanceType
}

func (s *SecuritiesAccount9) AddOptionType() *CorporateActionOption1FormatChoice {
	s.OptionType = new(CorporateActionOption1FormatChoice)
	return s.OptionType
}

func (s *SecuritiesAccount9) SetOptionNumber(value string) {
	s.OptionNumber = (*Exact3NumericText)(&value)
}

func (s *SecuritiesAccount9) SetSecurityHoldingForm(value string) {
	s.SecurityHoldingForm = (*FormOfSecurity1Code)(&value)
}
