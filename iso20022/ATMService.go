package iso20022

// Withdrawal service provided by the ATM inside the session.
type ATMService1 struct {

	// Unique identification of the withdrawal service provided by the ATM inside the session.
	ServiceReference *Max35Text `xml:"SvcRef,omitempty"`

	// Codification of the type of service for the ATM.
	ATMServiceCode *Max35Text `xml:"ATMSvcCd,omitempty"`

	// Describes the type of transaction selected by the customer.
	ServiceType *ATMServiceType1Code `xml:"SvcTp"`
}

func (a *ATMService1) SetServiceReference(value string) {
	a.ServiceReference = (*Max35Text)(&value)
}

func (a *ATMService1) SetATMServiceCode(value string) {
	a.ATMServiceCode = (*Max35Text)(&value)
}

func (a *ATMService1) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType1Code)(&value)
}

// Withdrawal service provided by the ATM inside the session.
type ATMService10 struct {

	// Unique identification of the withdrawal service provided by the ATM inside the session.
	ServiceReference *Max35Text `xml:"SvcRef,omitempty"`

	// Codification of the type of service for the ATM.
	ATMServiceCode *Max35Text `xml:"ATMSvcCd,omitempty"`

	// Codification of the type of service for the host.
	HostServiceCode *Max35Text `xml:"HstSvcCd,omitempty"`

	// Describes the type of withdrawal selected by the customer.
	ServiceType *ATMServiceType1Code `xml:"SvcTp"`

	// Identification of the variant of the service.
	ServiceVariantIdentification []*Max35Text `xml:"SvcVarntId,omitempty"`
}

func (a *ATMService10) SetServiceReference(value string) {
	a.ServiceReference = (*Max35Text)(&value)
}

func (a *ATMService10) SetATMServiceCode(value string) {
	a.ATMServiceCode = (*Max35Text)(&value)
}

func (a *ATMService10) SetHostServiceCode(value string) {
	a.HostServiceCode = (*Max35Text)(&value)
}

func (a *ATMService10) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType1Code)(&value)
}

func (a *ATMService10) AddServiceVariantIdentification(value string) {
	a.ServiceVariantIdentification = append(a.ServiceVariantIdentification, (*Max35Text)(&value))
}

// Deposit service provided by the ATM inside the session.
type ATMService11 struct {

	// Unique identification of the deposit service provided by the ATM inside the session.
	ServiceReference *Max35Text `xml:"SvcRef,omitempty"`

	// Codification of the type of service for the ATM.
	ATMServiceCode *Max35Text `xml:"ATMSvcCd,omitempty"`

	// Describes the type of deposit service selected by the customer.
	ServiceType *ATMServiceType6Code `xml:"SvcTp"`

	// Identification of the variant of the service.
	ServiceVariantIdentification []*Max35Text `xml:"SvcVarntId,omitempty"`

	// True if deposit with cash back transaction.
	CashBack *TrueFalseIndicator `xml:"CshBck,omitempty"`

	// True if the deposit transaction is split in multiple accounts.
	MultiAccount *TrueFalseIndicator `xml:"MultiAcct,omitempty"`
}

func (a *ATMService11) SetServiceReference(value string) {
	a.ServiceReference = (*Max35Text)(&value)
}

func (a *ATMService11) SetATMServiceCode(value string) {
	a.ATMServiceCode = (*Max35Text)(&value)
}

func (a *ATMService11) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType6Code)(&value)
}

func (a *ATMService11) AddServiceVariantIdentification(value string) {
	a.ServiceVariantIdentification = append(a.ServiceVariantIdentification, (*Max35Text)(&value))
}

func (a *ATMService11) SetCashBack(value string) {
	a.CashBack = (*TrueFalseIndicator)(&value)
}

func (a *ATMService11) SetMultiAccount(value string) {
	a.MultiAccount = (*TrueFalseIndicator)(&value)
}

// Deposit service provided by the ATM inside the session.
type ATMService12 struct {

	// Unique identification of the deposit service provided by the ATM inside the session.
	ServiceReference *Max35Text `xml:"SvcRef,omitempty"`

	// Codification of the type of service for the ATM.
	ATMServiceCode *Max35Text `xml:"ATMSvcCd,omitempty"`

	// Codification of the type of service for the host.
	HostServiceCode *Max35Text `xml:"HstSvcCd,omitempty"`

	// Describes the type of deposit service selected by the customer.
	ServiceType *ATMServiceType6Code `xml:"SvcTp"`

	// Identification of the variant of the service.
	ServiceVariantIdentification []*Max35Text `xml:"SvcVarntId,omitempty"`

	// True if deposit with cash back transaction.
	CashBack *TrueFalseIndicator `xml:"CshBck,omitempty"`

	// True if the deposit transaction is split in multiple accounts.
	MultiAccount *TrueFalseIndicator `xml:"MultiAcct,omitempty"`
}

func (a *ATMService12) SetServiceReference(value string) {
	a.ServiceReference = (*Max35Text)(&value)
}

func (a *ATMService12) SetATMServiceCode(value string) {
	a.ATMServiceCode = (*Max35Text)(&value)
}

func (a *ATMService12) SetHostServiceCode(value string) {
	a.HostServiceCode = (*Max35Text)(&value)
}

func (a *ATMService12) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType6Code)(&value)
}

func (a *ATMService12) AddServiceVariantIdentification(value string) {
	a.ServiceVariantIdentification = append(a.ServiceVariantIdentification, (*Max35Text)(&value))
}

func (a *ATMService12) SetCashBack(value string) {
	a.CashBack = (*TrueFalseIndicator)(&value)
}

func (a *ATMService12) SetMultiAccount(value string) {
	a.MultiAccount = (*TrueFalseIndicator)(&value)
}

// Deposit service provided by the ATM inside the session.
type ATMService13 struct {

	// Unique identification of the deposit service provided by the ATM inside the session.
	ServiceReference *Max35Text `xml:"SvcRef,omitempty"`

	// Codification of the type of service for the ATM.
	ATMServiceCode *Max35Text `xml:"ATMSvcCd,omitempty"`

	// Codification of the type of service for the host.
	HostServiceCode *Max35Text `xml:"HstSvcCd,omitempty"`

	// Describes the type of deposit service selected by the customer.
	ServiceType *ATMServiceType6Code `xml:"SvcTp"`

	// Identification of the variant of the service.
	ServiceVariantIdentification []*Max35Text `xml:"SvcVarntId,omitempty"`

	// True if deposit with cash back transaction.
	CashBack *TrueFalseIndicator `xml:"CshBck,omitempty"`

	// True if the deposit transaction is split in multiple accounts.
	MultiAccount *TrueFalseIndicator `xml:"MultiAcct,omitempty"`

	// True if this is not the final deposit.
	PartialDeposit *TrueFalseIndicator `xml:"PrtlDpst,omitempty"`
}

func (a *ATMService13) SetServiceReference(value string) {
	a.ServiceReference = (*Max35Text)(&value)
}

func (a *ATMService13) SetATMServiceCode(value string) {
	a.ATMServiceCode = (*Max35Text)(&value)
}

func (a *ATMService13) SetHostServiceCode(value string) {
	a.HostServiceCode = (*Max35Text)(&value)
}

func (a *ATMService13) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType6Code)(&value)
}

func (a *ATMService13) AddServiceVariantIdentification(value string) {
	a.ServiceVariantIdentification = append(a.ServiceVariantIdentification, (*Max35Text)(&value))
}

func (a *ATMService13) SetCashBack(value string) {
	a.CashBack = (*TrueFalseIndicator)(&value)
}

func (a *ATMService13) SetMultiAccount(value string) {
	a.MultiAccount = (*TrueFalseIndicator)(&value)
}

func (a *ATMService13) SetPartialDeposit(value string) {
	a.PartialDeposit = (*TrueFalseIndicator)(&value)
}

// Withdrawal service provided by the ATM inside the session.
type ATMService14 struct {

	// Unique identification of the service provided by the ATM inside the session.
	ServiceReference *Max35Text `xml:"SvcRef,omitempty"`

	// Codification of the type of service for the ATM.
	ATMServiceCode *Max35Text `xml:"ATMSvcCd,omitempty"`

	// Codification of the type of service for the host.
	HostServiceCode *Max35Text `xml:"HstSvcCd,omitempty"`

	// Describes the type of transaction selected by the customer.
	ServiceType *ATMServiceType7Code `xml:"SvcTp"`

	// Identification of the variant of the service.
	ServiceVariantIdentification []*Max35Text `xml:"SvcVarntId,omitempty"`
}

func (a *ATMService14) SetServiceReference(value string) {
	a.ServiceReference = (*Max35Text)(&value)
}

func (a *ATMService14) SetATMServiceCode(value string) {
	a.ATMServiceCode = (*Max35Text)(&value)
}

func (a *ATMService14) SetHostServiceCode(value string) {
	a.HostServiceCode = (*Max35Text)(&value)
}

func (a *ATMService14) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType7Code)(&value)
}

func (a *ATMService14) AddServiceVariantIdentification(value string) {
	a.ServiceVariantIdentification = append(a.ServiceVariantIdentification, (*Max35Text)(&value))
}

// Inquiry service provided by the ATM inside the session.
type ATMService15 struct {

	// Unique identification of the withdrawal service provided by the ATM inside the session.
	ServiceReference *Max35Text `xml:"SvcRef,omitempty"`

	// Codification of the type of service for the ATM.
	ATMServiceCode *Max35Text `xml:"ATMSvcCd,omitempty"`

	// Describes the type of inquiry selected by the customer or the ATM.
	ServiceType *ATMServiceType3Code `xml:"SvcTp"`

	// Identification of the variant of the service.
	ServiceVariantIdentification []*Max35Text `xml:"SvcVarntId,omitempty"`
}

func (a *ATMService15) SetServiceReference(value string) {
	a.ServiceReference = (*Max35Text)(&value)
}

func (a *ATMService15) SetATMServiceCode(value string) {
	a.ATMServiceCode = (*Max35Text)(&value)
}

func (a *ATMService15) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType3Code)(&value)
}

func (a *ATMService15) AddServiceVariantIdentification(value string) {
	a.ServiceVariantIdentification = append(a.ServiceVariantIdentification, (*Max35Text)(&value))
}

// Service provided by the ATM inside the session.
type ATMService16 struct {

	// Unique identification of the withdrawal service provided by the ATM inside the session.
	ServiceReference *Max35Text `xml:"SvcRef,omitempty"`

	// Codification of the type of service for the ATM.
	ATMServiceCode *Max35Text `xml:"ATMSvcCd,omitempty"`

	// Codification of the type of service for the host.
	HostServiceCode *Max35Text `xml:"HstSvcCd,omitempty"`

	// Describes the type of withdrawal selected by the customer.
	ServiceType *ATMServiceType3Code `xml:"SvcTp"`

	// Identification of the variant of the service.
	ServiceVariantIdentification []*Max35Text `xml:"SvcVarntId,omitempty"`
}

func (a *ATMService16) SetServiceReference(value string) {
	a.ServiceReference = (*Max35Text)(&value)
}

func (a *ATMService16) SetATMServiceCode(value string) {
	a.ATMServiceCode = (*Max35Text)(&value)
}

func (a *ATMService16) SetHostServiceCode(value string) {
	a.HostServiceCode = (*Max35Text)(&value)
}

func (a *ATMService16) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType3Code)(&value)
}

func (a *ATMService16) AddServiceVariantIdentification(value string) {
	a.ServiceVariantIdentification = append(a.ServiceVariantIdentification, (*Max35Text)(&value))
}

// Services allowed for the customer's profile.
type ATMService17 struct {

	// Describes the type of inquiry selected by the customer or the ATM.
	ServiceType *ATMServiceType8Code `xml:"SvcTp"`

	// Variant of the service.
	ServiceVariant []*ATMService18 `xml:"SvcVarnt,omitempty"`

	// Limits of amounts.
	Limits []*ATMTransactionAmounts6 `xml:"Lmts,omitempty"`

	// Preferred withdrawal transaction chosen by the customer.
	PreferredWithdrawal *ATMTransaction8 `xml:"PrefrdWdrwl,omitempty"`
}

func (a *ATMService17) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType8Code)(&value)
}

func (a *ATMService17) AddServiceVariant() *ATMService18 {
	newValue := new(ATMService18)
	a.ServiceVariant = append(a.ServiceVariant, newValue)
	return newValue
}

func (a *ATMService17) AddLimits() *ATMTransactionAmounts6 {
	newValue := new(ATMTransactionAmounts6)
	a.Limits = append(a.Limits, newValue)
	return newValue
}

func (a *ATMService17) AddPreferredWithdrawal() *ATMTransaction8 {
	a.PreferredWithdrawal = new(ATMTransaction8)
	return a.PreferredWithdrawal
}

// Service provided by the ATM inside the session.
type ATMService18 struct {

	// Identification of the service variant.
	Identification *Max35Text `xml:"Id"`

	// Label of the service variant.
	Label *Max35Text `xml:"Labl,omitempty"`
}

func (a *ATMService18) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *ATMService18) SetLabel(value string) {
	a.Label = (*Max35Text)(&value)
}

// Service allowed on the account.
type ATMService19 struct {

	// Describes the type of inquiry selected by the customer or the ATM.
	ServiceType *ATMServiceType8Code `xml:"SvcTp"`

	// Variant of the service.
	ServiceVariant []*ATMService18 `xml:"SvcVarnt,omitempty"`

	// Limits of amounts.
	Limits []*ATMTransactionAmounts6 `xml:"Lmts,omitempty"`
}

func (a *ATMService19) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType8Code)(&value)
}

func (a *ATMService19) AddServiceVariant() *ATMService18 {
	newValue := new(ATMService18)
	a.ServiceVariant = append(a.ServiceVariant, newValue)
	return newValue
}

func (a *ATMService19) AddLimits() *ATMTransactionAmounts6 {
	newValue := new(ATMTransactionAmounts6)
	a.Limits = append(a.Limits, newValue)
	return newValue
}

// Withdrawal service provided by the ATM inside the session.
type ATMService2 struct {

	// Unique identification of the withdrawal service provided by the ATM inside the session.
	ServiceReference *Max35Text `xml:"SvcRef,omitempty"`

	// Codification of the type of service for the ATM.
	ATMServiceCode *Max35Text `xml:"ATMSvcCd,omitempty"`

	// Codification of the type of service for the host.
	HostServiceCode *Max35Text `xml:"HstSvcCd,omitempty"`

	// Describes the type of withdrawal selected by the customer.
	ServiceType *ATMServiceType1Code `xml:"SvcTp"`
}

func (a *ATMService2) SetServiceReference(value string) {
	a.ServiceReference = (*Max35Text)(&value)
}

func (a *ATMService2) SetATMServiceCode(value string) {
	a.ATMServiceCode = (*Max35Text)(&value)
}

func (a *ATMService2) SetHostServiceCode(value string) {
	a.HostServiceCode = (*Max35Text)(&value)
}

func (a *ATMService2) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType1Code)(&value)
}

// Context in which the transaction is performed.
type ATMService20 struct {

	// Unique identification of the withdrawal service provided by the ATM inside the session.
	ServiceReference *Max35Text `xml:"SvcRef,omitempty"`

	// Codification of the type of service for the ATM.
	ATMServiceCode *Max35Text `xml:"ATMSvcCd,omitempty"`

	// Describes the type of inquiry selected by the customer or the ATM.
	ServiceType *ATMServiceType5Code `xml:"SvcTp"`

	// Identification of the variant of the service.
	ServiceVariantIdentification []*Max35Text `xml:"SvcVarntId,omitempty"`
}

func (a *ATMService20) SetServiceReference(value string) {
	a.ServiceReference = (*Max35Text)(&value)
}

func (a *ATMService20) SetATMServiceCode(value string) {
	a.ATMServiceCode = (*Max35Text)(&value)
}

func (a *ATMService20) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType5Code)(&value)
}

func (a *ATMService20) AddServiceVariantIdentification(value string) {
	a.ServiceVariantIdentification = append(a.ServiceVariantIdentification, (*Max35Text)(&value))
}

// Service provided by the ATM inside the session.
type ATMService21 struct {

	// Unique identification of the service provided by the ATM inside the session.
	ServiceReference *Max35Text `xml:"SvcRef,omitempty"`

	// Codification of the type of service for the ATM.
	ATMServiceCode *Max35Text `xml:"ATMSvcCd,omitempty"`

	// Codification of the type of service for the host.
	HostServiceCode *Max35Text `xml:"HstSvcCd,omitempty"`

	// Describes the type of withdrawal selected by the customer.
	ServiceType *ATMServiceType5Code `xml:"SvcTp"`

	// Identification of the variant of the service.
	ServiceVariantIdentification []*Max35Text `xml:"SvcVarntId,omitempty"`
}

func (a *ATMService21) SetServiceReference(value string) {
	a.ServiceReference = (*Max35Text)(&value)
}

func (a *ATMService21) SetATMServiceCode(value string) {
	a.ATMServiceCode = (*Max35Text)(&value)
}

func (a *ATMService21) SetHostServiceCode(value string) {
	a.HostServiceCode = (*Max35Text)(&value)
}

func (a *ATMService21) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType5Code)(&value)
}

func (a *ATMService21) AddServiceVariantIdentification(value string) {
	a.ServiceVariantIdentification = append(a.ServiceVariantIdentification, (*Max35Text)(&value))
}

// Inquiry service provided by the ATM inside the session.
type ATMService22 struct {

	// Unique identification of the customer session in which the transfer is performed.
	ServiceReference *Max35Text `xml:"SvcRef,omitempty"`

	// Codification of the type of service for the ATM.
	ATMServiceCode *Max35Text `xml:"ATMSvcCd,omitempty"`

	// Describes the type of fund transfer selected by the customer or the ATM.
	ServiceType *ATMServiceType9Code `xml:"SvcTp"`

	// Identification of the variant of the service.
	ServiceVariantIdentification []*Max35Text `xml:"SvcVarntId,omitempty"`
}

func (a *ATMService22) SetServiceReference(value string) {
	a.ServiceReference = (*Max35Text)(&value)
}

func (a *ATMService22) SetATMServiceCode(value string) {
	a.ATMServiceCode = (*Max35Text)(&value)
}

func (a *ATMService22) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType9Code)(&value)
}

func (a *ATMService22) AddServiceVariantIdentification(value string) {
	a.ServiceVariantIdentification = append(a.ServiceVariantIdentification, (*Max35Text)(&value))
}

// Inquiry service provided by the ATM inside the session.
type ATMService23 struct {

	// Unique identification of the customer session in which the transfer is performed.
	ServiceReference *Max35Text `xml:"SvcRef,omitempty"`

	// Codification of the type of service for the ATM.
	ATMServiceCode *Max35Text `xml:"ATMSvcCd,omitempty"`

	// Codification of the type of service for the host.
	HostServiceCode *Max35Text `xml:"HstSvcCd,omitempty"`

	// Describes the type of fund transfer selected by the customer or the ATM.
	ServiceType *ATMServiceType9Code `xml:"SvcTp"`

	// Identification of the variant of the service.
	ServiceVariantIdentification []*Max35Text `xml:"SvcVarntId,omitempty"`
}

func (a *ATMService23) SetServiceReference(value string) {
	a.ServiceReference = (*Max35Text)(&value)
}

func (a *ATMService23) SetATMServiceCode(value string) {
	a.ATMServiceCode = (*Max35Text)(&value)
}

func (a *ATMService23) SetHostServiceCode(value string) {
	a.HostServiceCode = (*Max35Text)(&value)
}

func (a *ATMService23) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType9Code)(&value)
}

func (a *ATMService23) AddServiceVariantIdentification(value string) {
	a.ServiceVariantIdentification = append(a.ServiceVariantIdentification, (*Max35Text)(&value))
}

// Service provided by the ATM inside the session.
type ATMService24 struct {

	// Unique identification of the service provided by the ATM inside the session.
	ServiceReference *Max35Text `xml:"SvcRef,omitempty"`

	// Codification of the type of service for the ATM.
	ATMServiceCode *Max35Text `xml:"ATMSvcCd,omitempty"`

	// Codification of the type of service for the host.
	HostServiceCode *Max35Text `xml:"HstSvcCd,omitempty"`

	// Describes the type of service selected by the customer.
	ServiceType *ATMServiceType10Code `xml:"SvcTp"`

	// Identification of the variant of the service.
	ServiceVariantIdentification []*Max35Text `xml:"SvcVarntId,omitempty"`
}

func (a *ATMService24) SetServiceReference(value string) {
	a.ServiceReference = (*Max35Text)(&value)
}

func (a *ATMService24) SetATMServiceCode(value string) {
	a.ATMServiceCode = (*Max35Text)(&value)
}

func (a *ATMService24) SetHostServiceCode(value string) {
	a.HostServiceCode = (*Max35Text)(&value)
}

func (a *ATMService24) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType10Code)(&value)
}

func (a *ATMService24) AddServiceVariantIdentification(value string) {
	a.ServiceVariantIdentification = append(a.ServiceVariantIdentification, (*Max35Text)(&value))
}

// Withdrawal service provided by the ATM inside the session.
type ATMService3 struct {

	// Unique identification of the service provided by the ATM inside the session.
	ServiceReference *Max35Text `xml:"SvcRef,omitempty"`

	// Codification of the type of service for the ATM.
	ATMServiceCode *Max35Text `xml:"ATMSvcCd,omitempty"`

	// Codification of the type of service for the host.
	HostServiceCode *Max35Text `xml:"HstSvcCd,omitempty"`

	// Describes the type of transaction selected by the customer.
	ServiceType *ATMServiceType2Code `xml:"SvcTp"`
}

func (a *ATMService3) SetServiceReference(value string) {
	a.ServiceReference = (*Max35Text)(&value)
}

func (a *ATMService3) SetATMServiceCode(value string) {
	a.ATMServiceCode = (*Max35Text)(&value)
}

func (a *ATMService3) SetHostServiceCode(value string) {
	a.HostServiceCode = (*Max35Text)(&value)
}

func (a *ATMService3) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType2Code)(&value)
}

// Service provided by the ATM inside the session.
type ATMService4 struct {

	// Unique identification of the service provided by the ATM inside the session.
	ServiceReference *Max35Text `xml:"SvcRef,omitempty"`

	// Codification of the type of service for the ATM.
	ATMServiceCode *Max35Text `xml:"ATMSvcCd,omitempty"`

	// Codification of the type of service for the host.
	HostServiceCode *Max35Text `xml:"HstSvcCd,omitempty"`

	// Describes the type of withdrawal selected by the customer.
	ServiceType *ATMServiceType5Code `xml:"SvcTp"`
}

func (a *ATMService4) SetServiceReference(value string) {
	a.ServiceReference = (*Max35Text)(&value)
}

func (a *ATMService4) SetATMServiceCode(value string) {
	a.ATMServiceCode = (*Max35Text)(&value)
}

func (a *ATMService4) SetHostServiceCode(value string) {
	a.HostServiceCode = (*Max35Text)(&value)
}

func (a *ATMService4) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType5Code)(&value)
}

// Inquiry service provided by the ATM inside the session.
type ATMService5 struct {

	// Unique identification of the withdrawal service provided by the ATM inside the session.
	ServiceReference *Max35Text `xml:"SvcRef,omitempty"`

	// Codification of the type of service for the ATM.
	ATMServiceCode *Max35Text `xml:"ATMSvcCd,omitempty"`

	// Describes the type of inquiry selected by the customer or the ATM.
	ServiceType *ATMServiceType3Code `xml:"SvcTp"`
}

func (a *ATMService5) SetServiceReference(value string) {
	a.ServiceReference = (*Max35Text)(&value)
}

func (a *ATMService5) SetATMServiceCode(value string) {
	a.ATMServiceCode = (*Max35Text)(&value)
}

func (a *ATMService5) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType3Code)(&value)
}

// Service provided by the ATM inside the session.
type ATMService6 struct {

	// Unique identification of the withdrawal service provided by the ATM inside the session.
	ServiceReference *Max35Text `xml:"SvcRef,omitempty"`

	// Codification of the type of service for the ATM.
	ATMServiceCode *Max35Text `xml:"ATMSvcCd,omitempty"`

	// Codification of the type of service for the host.
	HostServiceCode *Max35Text `xml:"HstSvcCd,omitempty"`

	// Describes the type of withdrawal selected by the customer.
	ServiceType *ATMServiceType3Code `xml:"SvcTp"`
}

func (a *ATMService6) SetServiceReference(value string) {
	a.ServiceReference = (*Max35Text)(&value)
}

func (a *ATMService6) SetATMServiceCode(value string) {
	a.ATMServiceCode = (*Max35Text)(&value)
}

func (a *ATMService6) SetHostServiceCode(value string) {
	a.HostServiceCode = (*Max35Text)(&value)
}

func (a *ATMService6) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType3Code)(&value)
}

// Services allowed for the customer's profile.
type ATMService7 struct {

	// Describes the type of inquiry selected by the customer or the ATM.
	ServiceType *ATMServiceType4Code `xml:"SvcTp"`

	// Limits of amounts.
	Limits []*ATMTransactionAmounts3 `xml:"Lmts,omitempty"`

	// Preferred withdrawal transaction chosen by the customer.
	PreferredWithdrawal *ATMTransaction8 `xml:"PrefrdWdrwl,omitempty"`
}

func (a *ATMService7) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType4Code)(&value)
}

func (a *ATMService7) AddLimits() *ATMTransactionAmounts3 {
	newValue := new(ATMTransactionAmounts3)
	a.Limits = append(a.Limits, newValue)
	return newValue
}

func (a *ATMService7) AddPreferredWithdrawal() *ATMTransaction8 {
	a.PreferredWithdrawal = new(ATMTransaction8)
	return a.PreferredWithdrawal
}

// Context in which the transaction is performed.
type ATMService8 struct {

	// Unique identification of the withdrawal service provided by the ATM inside the session.
	ServiceReference *Max35Text `xml:"SvcRef,omitempty"`

	// Codification of the type of service for the ATM.
	ATMServiceCode *Max35Text `xml:"ATMSvcCd,omitempty"`

	// Describes the type of inquiry selected by the customer or the ATM.
	ServiceType *ATMServiceType5Code `xml:"SvcTp"`
}

func (a *ATMService8) SetServiceReference(value string) {
	a.ServiceReference = (*Max35Text)(&value)
}

func (a *ATMService8) SetATMServiceCode(value string) {
	a.ATMServiceCode = (*Max35Text)(&value)
}

func (a *ATMService8) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType5Code)(&value)
}

// Withdrawal service provided by the ATM inside the session.
type ATMService9 struct {

	// Unique identification of the withdrawal service provided by the ATM inside the session.
	ServiceReference *Max35Text `xml:"SvcRef,omitempty"`

	// Codification of the type of service for the ATM.
	ATMServiceCode *Max35Text `xml:"ATMSvcCd,omitempty"`

	// Describes the type of transaction selected by the customer.
	ServiceType *ATMServiceType1Code `xml:"SvcTp"`

	// Identification of the variant of the service.
	ServiceVariantIdentification []*Max35Text `xml:"SvcVarntId,omitempty"`
}

func (a *ATMService9) SetServiceReference(value string) {
	a.ServiceReference = (*Max35Text)(&value)
}

func (a *ATMService9) SetATMServiceCode(value string) {
	a.ATMServiceCode = (*Max35Text)(&value)
}

func (a *ATMService9) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType1Code)(&value)
}

func (a *ATMService9) AddServiceVariantIdentification(value string) {
	a.ServiceVariantIdentification = append(a.ServiceVariantIdentification, (*Max35Text)(&value))
}
