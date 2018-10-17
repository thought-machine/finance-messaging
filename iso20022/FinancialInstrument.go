package iso20022

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument10 struct {

	// Identification of a security by an ISIN.
	Identification *SecurityIdentification3Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Additional information about a financial instrument to help identify the instrument.
	SupplementaryIdentification *Max35Text `xml:"SplmtryId,omitempty"`

	// Features of units offered by a fund. For example, a unit may have a specific load structure, eg, front end or back end, an income policy, eg, pay out or accumulate, or a trailer policy, eg, with or without. Fund classes are typically denoted by a single character, eg, 'Class A', 'Class 2'.
	ClassType *Max35Text `xml:"ClssTp,omitempty"`

	// Form, ie, ownership, of the security, eg, registered or bearer.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Income policy relating to a class type, ie, if income is paid out or retained in the fund.
	DistributionPolicy *DistributionPolicy1Code `xml:"DstrbtnPlcy,omitempty"`

	// Company specific description of a group of funds.
	ProductGroup *Max140Text `xml:"PdctGrp,omitempty"`
}

func (f *FinancialInstrument10) AddIdentification() *SecurityIdentification3Choice {
	f.Identification = new(SecurityIdentification3Choice)
	return f.Identification
}

func (f *FinancialInstrument10) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument10) SetSupplementaryIdentification(value string) {
	f.SupplementaryIdentification = (*Max35Text)(&value)
}

func (f *FinancialInstrument10) SetClassType(value string) {
	f.ClassType = (*Max35Text)(&value)
}

func (f *FinancialInstrument10) SetSecuritiesForm(value string) {
	f.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (f *FinancialInstrument10) SetDistributionPolicy(value string) {
	f.DistributionPolicy = (*DistributionPolicy1Code)(&value)
}

func (f *FinancialInstrument10) SetProductGroup(value string) {
	f.ProductGroup = (*Max140Text)(&value)
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument11 struct {

	// Unique and unambiguous identifier of a security, assigned under a formal or proprietary identification scheme.
	Identification *SecurityIdentification3Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Specifies whether the financial instrument is transferred as an asset or as cash.
	TransferType *TransferType1Code `xml:"TrfTp"`
}

func (f *FinancialInstrument11) AddIdentification() *SecurityIdentification3Choice {
	f.Identification = new(SecurityIdentification3Choice)
	return f.Identification
}

func (f *FinancialInstrument11) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument11) SetTransferType(value string) {
	f.TransferType = (*TransferType1Code)(&value)
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument12 struct {

	// Unique and unambiguous identifier of a security, assigned under a formal or proprietary identification scheme.
	Identification *SecurityIdentification3Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`
}

func (f *FinancialInstrument12) AddIdentification() *SecurityIdentification3Choice {
	f.Identification = new(SecurityIdentification3Choice)
	return f.Identification
}

func (f *FinancialInstrument12) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument13 struct {

	// Identification of a security by an ISIN.
	Identification *SecurityIdentification3Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Additional information about a financial instrument to help identify the instrument.
	SupplementaryIdentification *Max35Text `xml:"SplmtryId,omitempty"`

	// Features of units offered by a fund. For example, a unit may have a specific load structure, eg, front end or back end, an income policy, eg, pay out or accumulate, or a trailer policy, eg, with or without. Fund classes are typically denoted by a single character, eg, 'Class A', 'Class 2'.
	ClassType *Max35Text `xml:"ClssTp,omitempty"`

	// Form, ie, ownership, of the security, eg, registered or bearer.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Income policy relating to a class type, ie, if income is paid out or retained in the fund.
	DistributionPolicy *DistributionPolicy1Code `xml:"DstrbtnPlcy,omitempty"`
}

func (f *FinancialInstrument13) AddIdentification() *SecurityIdentification3Choice {
	f.Identification = new(SecurityIdentification3Choice)
	return f.Identification
}

func (f *FinancialInstrument13) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument13) SetSupplementaryIdentification(value string) {
	f.SupplementaryIdentification = (*Max35Text)(&value)
}

func (f *FinancialInstrument13) SetClassType(value string) {
	f.ClassType = (*Max35Text)(&value)
}

func (f *FinancialInstrument13) SetSecuritiesForm(value string) {
	f.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (f *FinancialInstrument13) SetDistributionPolicy(value string) {
	f.DistributionPolicy = (*DistributionPolicy1Code)(&value)
}

// Identifies the financial instrument.
type FinancialInstrument15 struct {

	// Identifies the financial instrument using a choice of either ISIN,  local code, or a description of the instrument. ISIN is the preferred format.
	Identification *SecurityIdentification6Choice `xml:"Id"`

	// Provides the ability to describe the instrument through a description and main characteristics.
	InstrumentDescription *SecurityInstrumentDescription2 `xml:"InstrmDesc,omitempty"`

	// Provides details of the underlying financial instrument for which the transaction report is being sent. If there is more than one underlying financial instrument then it is the dominant/ultimate instrument that should be identified here.
	UnderlyingInstrumentIdentification *SecurityIdentification6Choice `xml:"UndrlygInstrmId,omitempty"`
}

func (f *FinancialInstrument15) AddIdentification() *SecurityIdentification6Choice {
	f.Identification = new(SecurityIdentification6Choice)
	return f.Identification
}

func (f *FinancialInstrument15) AddInstrumentDescription() *SecurityInstrumentDescription2 {
	f.InstrumentDescription = new(SecurityInstrumentDescription2)
	return f.InstrumentDescription
}

func (f *FinancialInstrument15) AddUnderlyingInstrumentIdentification() *SecurityIdentification6Choice {
	f.UnderlyingInstrumentIdentification = new(SecurityIdentification6Choice)
	return f.UnderlyingInstrumentIdentification
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument17 struct {

	// Identification of a security by an ISIN.
	Identification *SecurityIdentification7 `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Additional information about a financial instrument to help identify the instrument.
	SupplementaryIdentification *Max35Text `xml:"SplmtryId,omitempty"`

	// Features of units offered by a fund. For example, a unit may have a specific load structure, eg, front end or back end, an income policy, eg, pay out or accumulate, or a trailer policy, eg, with or without. Fund classes are typically denoted by a single character, eg, 'Class A', 'Class 2'.
	ClassType *Max35Text `xml:"ClssTp,omitempty"`

	// Specifies the form, ie, ownership, of the security.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Income policy relating to a class type, ie, if income is paid out or retained in the fund.
	DistributionPolicy *DistributionPolicy1Code `xml:"DstrbtnPlcy,omitempty"`

	// Company specific description of a group of funds.
	ProductGroup *Max140Text `xml:"PdctGrp,omitempty"`
}

func (f *FinancialInstrument17) AddIdentification() *SecurityIdentification7 {
	f.Identification = new(SecurityIdentification7)
	return f.Identification
}

func (f *FinancialInstrument17) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument17) SetSupplementaryIdentification(value string) {
	f.SupplementaryIdentification = (*Max35Text)(&value)
}

func (f *FinancialInstrument17) SetClassType(value string) {
	f.ClassType = (*Max35Text)(&value)
}

func (f *FinancialInstrument17) SetSecuritiesForm(value string) {
	f.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (f *FinancialInstrument17) SetDistributionPolicy(value string) {
	f.DistributionPolicy = (*DistributionPolicy1Code)(&value)
}

func (f *FinancialInstrument17) SetProductGroup(value string) {
	f.ProductGroup = (*Max140Text)(&value)
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument20 struct {

	// Indicate whether or note it is possible to hold bearer units/shares in this class in certified form
	PhysicalBearerSecurities *YesNoIndicator `xml:"PhysBrScties"`

	// Indicate whether or not it is possible to hold bearer units/shares in paperless form
	DematerialisedBearerSecurities *YesNoIndicator `xml:"DmtrlsdBrScties"`

	// Indicate whether or not it is possible to hold registered units/shares in this class in paperless form
	PhysicalRegisteredSecurities *YesNoIndicator `xml:"PhysRegdScties"`

	// Indicate whether or not it is possible to hold registered units/shares in this class in paperless form
	DematerialisedRegisteredSecurities *YesNoIndicator `xml:"DmtrlsdRegdScties"`

	// Income policy relating to a class type, ie, if income is paid out or retained in the fund.
	DistributionPolicy *DistributionPolicy1Code `xml:"DstrbtnPlcy"`

	// Dividend policy of the fund, eg, cash, units.
	DividendPolicy *DividendPolicy1Code `xml:"DvddPlcy,omitempty"`

	// Frequency with which the income is allocated to investors.
	DividendFrequency *EventFrequency5Code `xml:"DvddFrqcy,omitempty"`

	// Frequency with which the reinvestment takes place,  This is the same or less than the dividend frequency,
	ReinvestmentFrequency *EventFrequency5Code `xml:"RinvstmtFrqcy,omitempty"`

	// Front end charge on subscription orders for this class can be applied.
	FrontEndLoad *YesNoIndicator `xml:"FrntEndLd"`

	// Exit charge (eg. CDSC) on redemption orders for this class can be applied.
	BackEndLoad *YesNoIndicator `xml:"BckEndLd"`

	// If a separate fee for switching between sub-funds of the same umbrella can be applied
	SwitchFee *YesNoIndicator `xml:"SwtchFee"`

	// Indicates whether the investment fund class is subject to the European Union Saving Directive.
	EUSavingsDirective *EUSavingsDirective1Code `xml:"EUSvgsDrctv"`
}

func (f *FinancialInstrument20) SetPhysicalBearerSecurities(value string) {
	f.PhysicalBearerSecurities = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrument20) SetDematerialisedBearerSecurities(value string) {
	f.DematerialisedBearerSecurities = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrument20) SetPhysicalRegisteredSecurities(value string) {
	f.PhysicalRegisteredSecurities = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrument20) SetDematerialisedRegisteredSecurities(value string) {
	f.DematerialisedRegisteredSecurities = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrument20) SetDistributionPolicy(value string) {
	f.DistributionPolicy = (*DistributionPolicy1Code)(&value)
}

func (f *FinancialInstrument20) SetDividendPolicy(value string) {
	f.DividendPolicy = (*DividendPolicy1Code)(&value)
}

func (f *FinancialInstrument20) SetDividendFrequency(value string) {
	f.DividendFrequency = (*EventFrequency5Code)(&value)
}

func (f *FinancialInstrument20) SetReinvestmentFrequency(value string) {
	f.ReinvestmentFrequency = (*EventFrequency5Code)(&value)
}

func (f *FinancialInstrument20) SetFrontEndLoad(value string) {
	f.FrontEndLoad = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrument20) SetBackEndLoad(value string) {
	f.BackEndLoad = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrument20) SetSwitchFee(value string) {
	f.SwitchFee = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrument20) SetEUSavingsDirective(value string) {
	f.EUSavingsDirective = (*EUSavingsDirective1Code)(&value)
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument21 struct {

	// Features of units offered by a fund. For example, a unit may have a specific load structure, eg, front end or back end, an income policy, eg, pay out or accumulate, or a trailer policy, eg, with or without. Fund classes are typically denoted by a single character, eg, 'Class A', 'Class 2'.
	ClassType *Max35Text `xml:"ClssTp,omitempty"`

	// Specifies the form, that is, ownership, of the security.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Income policy relating to a class type, that is, if income is paid out or retained in the fund.
	DistributionPolicy *DistributionPolicy1Code `xml:"DstrbtnPlcy,omitempty"`

	// Company specific description of a group of funds.
	ProductGroup *Max140Text `xml:"PdctGrp,omitempty"`

	// Name of the umbrella fund in which financial instrument is contained.
	UmbrellaName *Max35Text `xml:"UmbrllNm,omitempty"`

	// Currency of the investment fund class.
	BaseCurrency *ActiveCurrencyCode `xml:"BaseCcy,omitempty"`

	// Currency in which a security is issued or redenominated.
	DenominationCurrency *ActiveCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Indicates whether the fund has two prices.
	DualFundIndicator *YesNoIndicator `xml:"DualFndInd,omitempty"`

	// Country where the fund has legal domicile as reflected in the ISIN classification.
	CountryOfDomicile *CountryCode `xml:"CtryOfDmcl,omitempty"`

	// Countries where the fund is registered for distribution.
	RegisteredDistributionCountry []*CountryCode `xml:"RegdDstrbtnCtry,omitempty"`
}

func (f *FinancialInstrument21) SetClassType(value string) {
	f.ClassType = (*Max35Text)(&value)
}

func (f *FinancialInstrument21) SetSecuritiesForm(value string) {
	f.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (f *FinancialInstrument21) SetDistributionPolicy(value string) {
	f.DistributionPolicy = (*DistributionPolicy1Code)(&value)
}

func (f *FinancialInstrument21) SetProductGroup(value string) {
	f.ProductGroup = (*Max140Text)(&value)
}

func (f *FinancialInstrument21) SetUmbrellaName(value string) {
	f.UmbrellaName = (*Max35Text)(&value)
}

func (f *FinancialInstrument21) SetBaseCurrency(value string) {
	f.BaseCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *FinancialInstrument21) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *FinancialInstrument21) SetRequestedNAVCurrency(value string) {
	f.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrument21) SetDualFundIndicator(value string) {
	f.DualFundIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrument21) SetCountryOfDomicile(value string) {
	f.CountryOfDomicile = (*CountryCode)(&value)
}

func (f *FinancialInstrument21) AddRegisteredDistributionCountry(value string) {
	f.RegisteredDistributionCountry = append(f.RegisteredDistributionCountry, (*CountryCode)(&value))
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument22 struct {

	// Features of units offered by a fund. For example, a unit may have a specific load structure, eg, front end or back end, an income policy, eg, pay out or accumulate, or a trailer policy, eg, with or without. Fund classes are typically denoted by a single character, eg, 'Class A', 'Class 2'.
	ClassType *Max35Text `xml:"ClssTp,omitempty"`

	// Specifies the form, that is, ownership, of the security.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Income policy relating to a class type, that is, if income is paid out or retained in the fund.
	DistributionPolicy *DistributionPolicy1Code `xml:"DstrbtnPlcy,omitempty"`

	// Company specific description of a group of funds.
	ProductGroup *RestrictedFINXMax140Text `xml:"PdctGrp,omitempty"`

	// Name of the umbrella fund in which financial instrument is contained.
	UmbrellaName *Max35Text `xml:"UmbrllNm,omitempty"`

	// Currency of the investment fund class.
	BaseCurrency *ActiveCurrencyCode `xml:"BaseCcy,omitempty"`

	// Currency in which a security is issued or redenominated.
	DenominationCurrency *ActiveCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Indicates whether the fund has two prices.
	DualFundIndicator *YesNoIndicator `xml:"DualFndInd,omitempty"`

	// Country where the fund has legal domicile as reflected in the ISIN classification.
	CountryOfDomicile *CountryCode `xml:"CtryOfDmcl,omitempty"`

	// Countries where the fund is registered for distribution.
	RegisteredDistributionCountry []*CountryCode `xml:"RegdDstrbtnCtry,omitempty"`
}

func (f *FinancialInstrument22) SetClassType(value string) {
	f.ClassType = (*Max35Text)(&value)
}

func (f *FinancialInstrument22) SetSecuritiesForm(value string) {
	f.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (f *FinancialInstrument22) SetDistributionPolicy(value string) {
	f.DistributionPolicy = (*DistributionPolicy1Code)(&value)
}

func (f *FinancialInstrument22) SetProductGroup(value string) {
	f.ProductGroup = (*RestrictedFINXMax140Text)(&value)
}

func (f *FinancialInstrument22) SetUmbrellaName(value string) {
	f.UmbrellaName = (*Max35Text)(&value)
}

func (f *FinancialInstrument22) SetBaseCurrency(value string) {
	f.BaseCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *FinancialInstrument22) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *FinancialInstrument22) SetRequestedNAVCurrency(value string) {
	f.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrument22) SetDualFundIndicator(value string) {
	f.DualFundIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrument22) SetCountryOfDomicile(value string) {
	f.CountryOfDomicile = (*CountryCode)(&value)
}

func (f *FinancialInstrument22) AddRegisteredDistributionCountry(value string) {
	f.RegisteredDistributionCountry = append(f.RegisteredDistributionCountry, (*CountryCode)(&value))
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument23 struct {

	// Unique and unambiguous identifier of a security, assigned under a formal or proprietary identification scheme.
	Identification *SecurityIdentification3Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Specifies whether the financial instrument is transferred as an asset or as cash.
	TransferType *TransferType1Code `xml:"TrfTp"`

	// Specifies the quantity of assets to be transferred in units or in a percentage rate.
	Quantity *Quantity12Choice `xml:"Qty,omitempty"`

	// Average cost per share of a security, including all charges and commissions.
	AverageAcquisitionPrice *ActiveOrHistoricCurrencyAndAmount `xml:"AvrgAcqstnPric,omitempty"`

	// Net asset on balance sheet - total portfolio value minus or plus the unrealised gain or loss.
	TotalBookValue *ActiveOrHistoricCurrencyAndAmount `xml:"TtlBookVal,omitempty"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	TransfereeAccount *Account6 `xml:"TrfeeAcct,omitempty"`
}

func (f *FinancialInstrument23) AddIdentification() *SecurityIdentification3Choice {
	f.Identification = new(SecurityIdentification3Choice)
	return f.Identification
}

func (f *FinancialInstrument23) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument23) SetTransferType(value string) {
	f.TransferType = (*TransferType1Code)(&value)
}

func (f *FinancialInstrument23) AddQuantity() *Quantity12Choice {
	f.Quantity = new(Quantity12Choice)
	return f.Quantity
}

func (f *FinancialInstrument23) SetAverageAcquisitionPrice(value, currency string) {
	f.AverageAcquisitionPrice = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrument23) SetTotalBookValue(value, currency string) {
	f.TotalBookValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrument23) AddTransfereeAccount() *Account6 {
	f.TransfereeAccount = new(Account6)
	return f.TransfereeAccount
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument24 struct {

	// Unique and unambiguous identifier of a security, assigned under a formal or proprietary identification scheme.
	Identification *SecurityIdentification3Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Specifies whether the financial instrument is transferred as an asset or as cash.
	TransferType *TransferType1Code `xml:"TrfTp"`

	// Specifies the quantity of assets to be transferred in units or in a percentage rate.
	Quantity *Quantity12Choice `xml:"Qty,omitempty"`

	// Average cost per share of a security, including all charges and commissions.
	AverageAcquisitionPrice *ActiveCurrencyAndAmount `xml:"AvrgAcqstnPric,omitempty"`

	// Net asset on balance sheet - total portfolio value minus or plus the unrealised gain or loss.
	TotalBookValue *ActiveOrHistoricCurrencyAndAmount `xml:"TtlBookVal,omitempty"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	TransfereeAccount *Account6 `xml:"TrfeeAcct,omitempty"`
}

func (f *FinancialInstrument24) AddIdentification() *SecurityIdentification3Choice {
	f.Identification = new(SecurityIdentification3Choice)
	return f.Identification
}

func (f *FinancialInstrument24) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument24) SetTransferType(value string) {
	f.TransferType = (*TransferType1Code)(&value)
}

func (f *FinancialInstrument24) AddQuantity() *Quantity12Choice {
	f.Quantity = new(Quantity12Choice)
	return f.Quantity
}

func (f *FinancialInstrument24) SetAverageAcquisitionPrice(value, currency string) {
	f.AverageAcquisitionPrice = NewActiveCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrument24) SetTotalBookValue(value, currency string) {
	f.TotalBookValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrument24) AddTransfereeAccount() *Account6 {
	f.TransfereeAccount = new(Account6)
	return f.TransfereeAccount
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument25 struct {

	// Unique and unambiguous identifier of a security, assigned under a formal or proprietary identification scheme.
	Identification *SecurityIdentification3Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Specifies whether the financial instrument is transferred as an asset or as cash.
	TransferType *TransferType1Code `xml:"TrfTp"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	TransfereeAccount *Account6 `xml:"TrfeeAcct,omitempty"`
}

func (f *FinancialInstrument25) AddIdentification() *SecurityIdentification3Choice {
	f.Identification = new(SecurityIdentification3Choice)
	return f.Identification
}

func (f *FinancialInstrument25) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument25) SetTransferType(value string) {
	f.TransferType = (*TransferType1Code)(&value)
}

func (f *FinancialInstrument25) AddTransfereeAccount() *Account6 {
	f.TransfereeAccount = new(Account6)
	return f.TransfereeAccount
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument26 struct {

	// Unique and unambiguous identifier of a security, assigned under a formal or proprietary identification scheme.
	Identification *SecurityIdentification3Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Specifies the quantity of assets to be transferred in units or in a percentage rate.
	Quantity *Quantity12Choice `xml:"Qty,omitempty"`

	// Average cost per share of a security, including all charges and commissions.
	AverageAcquisitionPrice *ActiveOrHistoricCurrencyAndAmount `xml:"AvrgAcqstnPric,omitempty"`

	// Net asset on balance sheet - total portfolio value minus or plus the unrealised gain or loss.
	TotalBookValue *ActiveOrHistoricCurrencyAndAmount `xml:"TtlBookVal,omitempty"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	TransfereeAccount *Account6 `xml:"TrfeeAcct,omitempty"`
}

func (f *FinancialInstrument26) AddIdentification() *SecurityIdentification3Choice {
	f.Identification = new(SecurityIdentification3Choice)
	return f.Identification
}

func (f *FinancialInstrument26) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument26) AddQuantity() *Quantity12Choice {
	f.Quantity = new(Quantity12Choice)
	return f.Quantity
}

func (f *FinancialInstrument26) SetAverageAcquisitionPrice(value, currency string) {
	f.AverageAcquisitionPrice = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrument26) SetTotalBookValue(value, currency string) {
	f.TotalBookValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrument26) AddTransfereeAccount() *Account6 {
	f.TransfereeAccount = new(Account6)
	return f.TransfereeAccount
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument27 struct {

	// Unique and unambiguous identifier of a security, assigned under a formal or proprietary identification scheme.
	Identification *SecurityIdentification3Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	TransfereeAccount *Account6 `xml:"TrfeeAcct,omitempty"`
}

func (f *FinancialInstrument27) AddIdentification() *SecurityIdentification3Choice {
	f.Identification = new(SecurityIdentification3Choice)
	return f.Identification
}

func (f *FinancialInstrument27) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument27) AddTransfereeAccount() *Account6 {
	f.TransfereeAccount = new(Account6)
	return f.TransfereeAccount
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument29 struct {

	// Identification of a security by an ISIN.
	Identification *SecurityIdentification3Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Additional information about a financial instrument to help identify the instrument.
	SupplementaryIdentification *Max35Text `xml:"SplmtryId,omitempty"`

	// Features of units offered by a fund. For example, a unit may have a specific load structure, eg, front end or back end, an income policy, eg, pay out or accumulate, or a trailer policy, eg, with or without. Fund classes are typically denoted by a single character, eg, 'Class A', 'Class 2'.
	ClassType *Max35Text `xml:"ClssTp,omitempty"`

	// Form, ie, ownership, of the security, eg, registered or bearer.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Income policy relating to a class type, ie, if income is paid out or retained in the fund.
	DistributionPolicy *DistributionPolicy1Code `xml:"DstrbtnPlcy,omitempty"`

	// Company specific description of a group of funds.
	ProductGroup *Max140Text `xml:"PdctGrp,omitempty"`

	// When an account at fund level is blocked, this specifies details on how the holding is blocked.
	BlockedHoldingDetails *BlockedHoldingDetails1 `xml:"BlckdHldgDtls,omitempty"`
}

func (f *FinancialInstrument29) AddIdentification() *SecurityIdentification3Choice {
	f.Identification = new(SecurityIdentification3Choice)
	return f.Identification
}

func (f *FinancialInstrument29) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument29) SetSupplementaryIdentification(value string) {
	f.SupplementaryIdentification = (*Max35Text)(&value)
}

func (f *FinancialInstrument29) SetClassType(value string) {
	f.ClassType = (*Max35Text)(&value)
}

func (f *FinancialInstrument29) SetSecuritiesForm(value string) {
	f.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (f *FinancialInstrument29) SetDistributionPolicy(value string) {
	f.DistributionPolicy = (*DistributionPolicy1Code)(&value)
}

func (f *FinancialInstrument29) SetProductGroup(value string) {
	f.ProductGroup = (*Max140Text)(&value)
}

func (f *FinancialInstrument29) AddBlockedHoldingDetails() *BlockedHoldingDetails1 {
	f.BlockedHoldingDetails = new(BlockedHoldingDetails1)
	return f.BlockedHoldingDetails
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument3 struct {

	// Identification of a security by an ISIN.
	Identification *SecurityIdentification1Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Additional information about a financial instrument to help identify the instrument.
	SupplementaryIdentification *Max35Text `xml:"SplmtryId,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Features of units offered by a fund. For example, a unit may have a specific load structure, eg, front end or back end, an income policy, eg, pay out or accumulate, or a trailer policy, eg, with or without. Fund classes are typically denoted by a single character, eg, 'Class A', 'Class 2'.
	ClassType *Max35Text `xml:"ClssTp,omitempty"`

	// Form, ie, ownership, of the security, eg, registered or bearer.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Income policy relating to a class type, ie, if income is paid out or retained in the fund.
	DistributionPolicy *DistributionPolicy1Code `xml:"DstrbtnPlcy,omitempty"`
}

func (f *FinancialInstrument3) AddIdentification() *SecurityIdentification1Choice {
	f.Identification = new(SecurityIdentification1Choice)
	return f.Identification
}

func (f *FinancialInstrument3) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument3) SetSupplementaryIdentification(value string) {
	f.SupplementaryIdentification = (*Max35Text)(&value)
}

func (f *FinancialInstrument3) SetRequestedNAVCurrency(value string) {
	f.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrument3) SetClassType(value string) {
	f.ClassType = (*Max35Text)(&value)
}

func (f *FinancialInstrument3) SetSecuritiesForm(value string) {
	f.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (f *FinancialInstrument3) SetDistributionPolicy(value string) {
	f.DistributionPolicy = (*DistributionPolicy1Code)(&value)
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument30 struct {

	// Unique and unambiguous identifier of a security, assigned under a formal or proprietary identification scheme.
	Identification *SecurityIdentification3Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Specifies the quantity of assets to be transferred in units or in a percentage rate.
	Quantity *Quantity12Choice `xml:"Qty,omitempty"`

	// Average cost per share of a security, including all charges and commissions.
	AverageAcquisitionPrice *ActiveOrHistoricCurrencyAndAmount `xml:"AvrgAcqstnPric,omitempty"`

	// Net asset on balance sheet - total portfolio value minus or plus the unrealised gain or loss.
	TotalBookValue *ActiveOrHistoricCurrencyAndAmount `xml:"TtlBookVal,omitempty"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	TransfereeAccount *Account6 `xml:"TrfeeAcct,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount93 `xml:"DlvrgAgtDtls,omitempty"`
}

func (f *FinancialInstrument30) AddIdentification() *SecurityIdentification3Choice {
	f.Identification = new(SecurityIdentification3Choice)
	return f.Identification
}

func (f *FinancialInstrument30) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument30) AddQuantity() *Quantity12Choice {
	f.Quantity = new(Quantity12Choice)
	return f.Quantity
}

func (f *FinancialInstrument30) SetAverageAcquisitionPrice(value, currency string) {
	f.AverageAcquisitionPrice = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrument30) SetTotalBookValue(value, currency string) {
	f.TotalBookValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrument30) AddTransfereeAccount() *Account6 {
	f.TransfereeAccount = new(Account6)
	return f.TransfereeAccount
}

func (f *FinancialInstrument30) AddDeliveringAgentDetails() *PartyIdentificationAndAccount93 {
	f.DeliveringAgentDetails = new(PartyIdentificationAndAccount93)
	return f.DeliveringAgentDetails
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument31 struct {

	// Unique and unambiguous identifier of a security, assigned under a formal or proprietary identification scheme.
	Identification *SecurityIdentification3Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Specifies whether the financial instrument is transferred as an asset or as cash.
	TransferType *TransferType1Code `xml:"TrfTp"`

	// Specifies the quantity of assets to be transferred in units or in a percentage rate.
	Quantity *Quantity12Choice `xml:"Qty,omitempty"`

	// Average cost per share of a security, including all charges and commissions.
	AverageAcquisitionPrice *ActiveOrHistoricCurrencyAndAmount `xml:"AvrgAcqstnPric,omitempty"`

	// Net asset on balance sheet - total portfolio value minus or plus the unrealised gain or loss.
	TotalBookValue *ActiveOrHistoricCurrencyAndAmount `xml:"TtlBookVal,omitempty"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	TransfereeAccount *Account6 `xml:"TrfeeAcct,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, for example, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount93 `xml:"RcvgAgtDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount93 `xml:"DlvrgAgtDtls,omitempty"`
}

func (f *FinancialInstrument31) AddIdentification() *SecurityIdentification3Choice {
	f.Identification = new(SecurityIdentification3Choice)
	return f.Identification
}

func (f *FinancialInstrument31) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument31) SetTransferType(value string) {
	f.TransferType = (*TransferType1Code)(&value)
}

func (f *FinancialInstrument31) AddQuantity() *Quantity12Choice {
	f.Quantity = new(Quantity12Choice)
	return f.Quantity
}

func (f *FinancialInstrument31) SetAverageAcquisitionPrice(value, currency string) {
	f.AverageAcquisitionPrice = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrument31) SetTotalBookValue(value, currency string) {
	f.TotalBookValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrument31) AddTransfereeAccount() *Account6 {
	f.TransfereeAccount = new(Account6)
	return f.TransfereeAccount
}

func (f *FinancialInstrument31) AddReceivingAgentDetails() *PartyIdentificationAndAccount93 {
	f.ReceivingAgentDetails = new(PartyIdentificationAndAccount93)
	return f.ReceivingAgentDetails
}

func (f *FinancialInstrument31) AddDeliveringAgentDetails() *PartyIdentificationAndAccount93 {
	f.DeliveringAgentDetails = new(PartyIdentificationAndAccount93)
	return f.DeliveringAgentDetails
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument32 struct {

	// Unique and unambiguous identifier of a security, assigned under a formal or proprietary identification scheme.
	Identification *SecurityIdentification3Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Specifies whether the financial instrument is transferred as an asset or as cash.
	TransferType *TransferType1Code `xml:"TrfTp"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	TransfereeAccount *Account6 `xml:"TrfeeAcct,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, for example, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount93 `xml:"RcvgAgtDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount93 `xml:"DlvrgAgtDtls,omitempty"`
}

func (f *FinancialInstrument32) AddIdentification() *SecurityIdentification3Choice {
	f.Identification = new(SecurityIdentification3Choice)
	return f.Identification
}

func (f *FinancialInstrument32) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument32) SetTransferType(value string) {
	f.TransferType = (*TransferType1Code)(&value)
}

func (f *FinancialInstrument32) AddTransfereeAccount() *Account6 {
	f.TransfereeAccount = new(Account6)
	return f.TransfereeAccount
}

func (f *FinancialInstrument32) AddReceivingAgentDetails() *PartyIdentificationAndAccount93 {
	f.ReceivingAgentDetails = new(PartyIdentificationAndAccount93)
	return f.ReceivingAgentDetails
}

func (f *FinancialInstrument32) AddDeliveringAgentDetails() *PartyIdentificationAndAccount93 {
	f.DeliveringAgentDetails = new(PartyIdentificationAndAccount93)
	return f.DeliveringAgentDetails
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument33 struct {

	// Unique and unambiguous identifier of a security, assigned under a formal or proprietary identification scheme.
	Identification *SecurityIdentification3Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Specifies whether the financial instrument is transferred as an asset or as cash.
	TransferType *TransferType1Code `xml:"TrfTp"`

	// Specifies the quantity of assets to be transferred in units or in a percentage rate.
	Quantity *Quantity14Choice `xml:"Qty"`

	// Average cost per share of a security, including all charges and commissions.
	AverageAcquisitionPrice *ActiveCurrencyAndAmount `xml:"AvrgAcqstnPric,omitempty"`

	// Net asset on balance sheet - total portfolio value minus or plus the unrealised gain or loss.
	TotalBookValue *ActiveOrHistoricCurrencyAndAmount `xml:"TtlBookVal,omitempty"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	TransfereeAccount *Account6 `xml:"TrfeeAcct,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, for example, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount93 `xml:"RcvgAgtDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount93 `xml:"DlvrgAgtDtls,omitempty"`

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`
}

func (f *FinancialInstrument33) AddIdentification() *SecurityIdentification3Choice {
	f.Identification = new(SecurityIdentification3Choice)
	return f.Identification
}

func (f *FinancialInstrument33) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument33) SetTransferType(value string) {
	f.TransferType = (*TransferType1Code)(&value)
}

func (f *FinancialInstrument33) AddQuantity() *Quantity14Choice {
	f.Quantity = new(Quantity14Choice)
	return f.Quantity
}

func (f *FinancialInstrument33) SetAverageAcquisitionPrice(value, currency string) {
	f.AverageAcquisitionPrice = NewActiveCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrument33) SetTotalBookValue(value, currency string) {
	f.TotalBookValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrument33) AddTransfereeAccount() *Account6 {
	f.TransfereeAccount = new(Account6)
	return f.TransfereeAccount
}

func (f *FinancialInstrument33) AddReceivingAgentDetails() *PartyIdentificationAndAccount93 {
	f.ReceivingAgentDetails = new(PartyIdentificationAndAccount93)
	return f.ReceivingAgentDetails
}

func (f *FinancialInstrument33) AddDeliveringAgentDetails() *PartyIdentificationAndAccount93 {
	f.DeliveringAgentDetails = new(PartyIdentificationAndAccount93)
	return f.DeliveringAgentDetails
}

func (f *FinancialInstrument33) SetRequestedSettlementDate(value string) {
	f.RequestedSettlementDate = (*ISODate)(&value)
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument34 struct {

	// Unique and unambiguous identifier of a security, assigned under a formal or proprietary identification scheme.
	Identification *SecurityIdentification3Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Specifies whether the financial instrument is transferred as an asset or as cash.
	TransferType *TransferType1Code `xml:"TrfTp"`

	// Specifies the quantity of assets to be transferred in units or in a percentage rate.
	Quantity *Quantity12Choice `xml:"Qty,omitempty"`

	// Average cost per share of a security, including all charges and commissions.
	AverageAcquisitionPrice *ActiveOrHistoricCurrencyAndAmount `xml:"AvrgAcqstnPric,omitempty"`

	// Net asset on balance sheet - total portfolio value minus or plus the unrealised gain or loss.
	TotalBookValue *ActiveOrHistoricCurrencyAndAmount `xml:"TtlBookVal,omitempty"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	TransfereeAccount *Account16 `xml:"TrfeeAcct,omitempty"`

	// Sub-accounts that are grouped in a master or omnibus account.
	SubAccountDetails *SubAccount1 `xml:"SubAcctDtls,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, for example, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount93 `xml:"RcvgAgtDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount93 `xml:"DlvrgAgtDtls,omitempty"`
}

func (f *FinancialInstrument34) AddIdentification() *SecurityIdentification3Choice {
	f.Identification = new(SecurityIdentification3Choice)
	return f.Identification
}

func (f *FinancialInstrument34) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument34) SetTransferType(value string) {
	f.TransferType = (*TransferType1Code)(&value)
}

func (f *FinancialInstrument34) AddQuantity() *Quantity12Choice {
	f.Quantity = new(Quantity12Choice)
	return f.Quantity
}

func (f *FinancialInstrument34) SetAverageAcquisitionPrice(value, currency string) {
	f.AverageAcquisitionPrice = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrument34) SetTotalBookValue(value, currency string) {
	f.TotalBookValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrument34) AddTransfereeAccount() *Account16 {
	f.TransfereeAccount = new(Account16)
	return f.TransfereeAccount
}

func (f *FinancialInstrument34) AddSubAccountDetails() *SubAccount1 {
	f.SubAccountDetails = new(SubAccount1)
	return f.SubAccountDetails
}

func (f *FinancialInstrument34) AddReceivingAgentDetails() *PartyIdentificationAndAccount93 {
	f.ReceivingAgentDetails = new(PartyIdentificationAndAccount93)
	return f.ReceivingAgentDetails
}

func (f *FinancialInstrument34) AddDeliveringAgentDetails() *PartyIdentificationAndAccount93 {
	f.DeliveringAgentDetails = new(PartyIdentificationAndAccount93)
	return f.DeliveringAgentDetails
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument35 struct {

	// Unique and unambiguous identifier of a security, assigned under a formal or proprietary identification scheme.
	Identification *SecurityIdentification3Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Specifies whether the financial instrument is transferred as an asset or as cash.
	TransferType *TransferType1Code `xml:"TrfTp"`

	// Specifies the quantity of assets to be transferred in units or in a percentage rate.
	Quantity *Quantity14Choice `xml:"Qty"`

	// Average cost per share of a security, including all charges and commissions.
	AverageAcquisitionPrice *ActiveCurrencyAndAmount `xml:"AvrgAcqstnPric,omitempty"`

	// Net asset on balance sheet - total portfolio value minus or plus the unrealised gain or loss.
	TotalBookValue *ActiveOrHistoricCurrencyAndAmount `xml:"TtlBookVal,omitempty"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	TransfereeAccount *Account16 `xml:"TrfeeAcct,omitempty"`

	// Sub-accounts that are grouped in a master or omnibus account.
	SubAccountDetails *SubAccount1 `xml:"SubAcctDtls,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, for example, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount93 `xml:"RcvgAgtDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount93 `xml:"DlvrgAgtDtls,omitempty"`

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`
}

func (f *FinancialInstrument35) AddIdentification() *SecurityIdentification3Choice {
	f.Identification = new(SecurityIdentification3Choice)
	return f.Identification
}

func (f *FinancialInstrument35) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument35) SetTransferType(value string) {
	f.TransferType = (*TransferType1Code)(&value)
}

func (f *FinancialInstrument35) AddQuantity() *Quantity14Choice {
	f.Quantity = new(Quantity14Choice)
	return f.Quantity
}

func (f *FinancialInstrument35) SetAverageAcquisitionPrice(value, currency string) {
	f.AverageAcquisitionPrice = NewActiveCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrument35) SetTotalBookValue(value, currency string) {
	f.TotalBookValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrument35) AddTransfereeAccount() *Account16 {
	f.TransfereeAccount = new(Account16)
	return f.TransfereeAccount
}

func (f *FinancialInstrument35) AddSubAccountDetails() *SubAccount1 {
	f.SubAccountDetails = new(SubAccount1)
	return f.SubAccountDetails
}

func (f *FinancialInstrument35) AddReceivingAgentDetails() *PartyIdentificationAndAccount93 {
	f.ReceivingAgentDetails = new(PartyIdentificationAndAccount93)
	return f.ReceivingAgentDetails
}

func (f *FinancialInstrument35) AddDeliveringAgentDetails() *PartyIdentificationAndAccount93 {
	f.DeliveringAgentDetails = new(PartyIdentificationAndAccount93)
	return f.DeliveringAgentDetails
}

func (f *FinancialInstrument35) SetRequestedSettlementDate(value string) {
	f.RequestedSettlementDate = (*ISODate)(&value)
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument36 struct {

	// Unique and unambiguous identifier of a security, assigned under a formal or proprietary identification scheme.
	Identification *SecurityIdentification3Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	TransfereeAccount *Account16 `xml:"TrfeeAcct,omitempty"`

	// Sub-accounts that are grouped in a master or omnibus account.
	SubAccountDetails *SubAccount1 `xml:"SubAcctDtls,omitempty"`
}

func (f *FinancialInstrument36) AddIdentification() *SecurityIdentification3Choice {
	f.Identification = new(SecurityIdentification3Choice)
	return f.Identification
}

func (f *FinancialInstrument36) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument36) AddTransfereeAccount() *Account16 {
	f.TransfereeAccount = new(Account16)
	return f.TransfereeAccount
}

func (f *FinancialInstrument36) AddSubAccountDetails() *SubAccount1 {
	f.SubAccountDetails = new(SubAccount1)
	return f.SubAccountDetails
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument37 struct {

	// Unique and unambiguous identifier of a security, assigned under a formal or proprietary identification scheme.
	Identification *SecurityIdentification3Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Specifies the quantity of assets to be transferred in units or in a percentage rate.
	Quantity *Quantity12Choice `xml:"Qty,omitempty"`

	// Average cost per share of a security, including all charges and commissions.
	AverageAcquisitionPrice *ActiveOrHistoricCurrencyAndAmount `xml:"AvrgAcqstnPric,omitempty"`

	// Net asset on balance sheet - total portfolio value minus or plus the unrealised gain or loss.
	TotalBookValue *ActiveOrHistoricCurrencyAndAmount `xml:"TtlBookVal,omitempty"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	TransfereeAccount *Account16 `xml:"TrfeeAcct,omitempty"`

	// Sub-accounts that are grouped in a master or omnibus account.
	SubAccountDetails *SubAccount1 `xml:"SubAcctDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount93 `xml:"DlvrgAgtDtls,omitempty"`
}

func (f *FinancialInstrument37) AddIdentification() *SecurityIdentification3Choice {
	f.Identification = new(SecurityIdentification3Choice)
	return f.Identification
}

func (f *FinancialInstrument37) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument37) AddQuantity() *Quantity12Choice {
	f.Quantity = new(Quantity12Choice)
	return f.Quantity
}

func (f *FinancialInstrument37) SetAverageAcquisitionPrice(value, currency string) {
	f.AverageAcquisitionPrice = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrument37) SetTotalBookValue(value, currency string) {
	f.TotalBookValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrument37) AddTransfereeAccount() *Account16 {
	f.TransfereeAccount = new(Account16)
	return f.TransfereeAccount
}

func (f *FinancialInstrument37) AddSubAccountDetails() *SubAccount1 {
	f.SubAccountDetails = new(SubAccount1)
	return f.SubAccountDetails
}

func (f *FinancialInstrument37) AddDeliveringAgentDetails() *PartyIdentificationAndAccount93 {
	f.DeliveringAgentDetails = new(PartyIdentificationAndAccount93)
	return f.DeliveringAgentDetails
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument39 struct {

	// Unique and unambiguous identifier of a security, assigned under a formal or proprietary identification scheme.
	Identification *SecurityIdentification3Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Specifies whether the financial instrument is transferred as an asset or as cash.
	TransferType *TransferType1Code `xml:"TrfTp"`

	// Specifies the quantity of assets to be transferred in units or in a percentage rate.
	Quantity *Quantity12Choice `xml:"Qty,omitempty"`

	// Average cost per share of a security, including all charges and commissions.
	AverageAcquisitionPrice *ActiveOrHistoricCurrencyAndAmount `xml:"AvrgAcqstnPric,omitempty"`

	// Identifies the currency to be used to transfer the holdings. Some transfer agents register holdings grouped by currency in addition to using the ISIN for multi-currency fund shares.
	TransferCurrency *CurrencyCode `xml:"TrfCcy,omitempty"`

	// Net asset on balance sheet - total portfolio value minus or plus the unrealised gain or loss.
	TotalBookValue *ActiveOrHistoricCurrencyAndAmount `xml:"TtlBookVal,omitempty"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	TransfereeAccount *Account16 `xml:"TrfeeAcct,omitempty"`

	// Sub-accounts that are grouped in a master or omnibus account.
	SubAccountDetails *SubAccount1 `xml:"SubAcctDtls,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, for example, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount93 `xml:"RcvgAgtDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount93 `xml:"DlvrgAgtDtls,omitempty"`
}

func (f *FinancialInstrument39) AddIdentification() *SecurityIdentification3Choice {
	f.Identification = new(SecurityIdentification3Choice)
	return f.Identification
}

func (f *FinancialInstrument39) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument39) SetTransferType(value string) {
	f.TransferType = (*TransferType1Code)(&value)
}

func (f *FinancialInstrument39) AddQuantity() *Quantity12Choice {
	f.Quantity = new(Quantity12Choice)
	return f.Quantity
}

func (f *FinancialInstrument39) SetAverageAcquisitionPrice(value, currency string) {
	f.AverageAcquisitionPrice = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrument39) SetTransferCurrency(value string) {
	f.TransferCurrency = (*CurrencyCode)(&value)
}

func (f *FinancialInstrument39) SetTotalBookValue(value, currency string) {
	f.TotalBookValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrument39) AddTransfereeAccount() *Account16 {
	f.TransfereeAccount = new(Account16)
	return f.TransfereeAccount
}

func (f *FinancialInstrument39) AddSubAccountDetails() *SubAccount1 {
	f.SubAccountDetails = new(SubAccount1)
	return f.SubAccountDetails
}

func (f *FinancialInstrument39) AddReceivingAgentDetails() *PartyIdentificationAndAccount93 {
	f.ReceivingAgentDetails = new(PartyIdentificationAndAccount93)
	return f.ReceivingAgentDetails
}

func (f *FinancialInstrument39) AddDeliveringAgentDetails() *PartyIdentificationAndAccount93 {
	f.DeliveringAgentDetails = new(PartyIdentificationAndAccount93)
	return f.DeliveringAgentDetails
}

// Security that is a securitised right for entitlement, eg, equity or bond or a sub-set of an investment fund.
type FinancialInstrument4 struct {

	// Identification of a security, as assigned under a formal or proprietary identification scheme.
	Identification *SecurityIdentification1Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Additional information about a financial instrument to help identify the instrument.
	SupplementaryIdentification *Max35Text `xml:"SplmtryId,omitempty"`

	// Form, ie, ownership, of the security, eg, registered or bearer.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Features of units offered by a fund. For example, a unit may have a specific load structure, eg, front end or back end, an income policy, eg, pay out or accumulate, or a trailer policy, eg, with or without. Fund classes are typically denoted by a single character, eg, 'Class A', 'Class 2'.
	ClassType *Max35Text `xml:"ClssTp,omitempty"`

	// Income policy relating to a class type, ie, if income is paid out or retained in the fund.
	DistributionPolicy *DistributionPolicy1Code `xml:"DstrbtnPlcy,omitempty"`
}

func (f *FinancialInstrument4) AddIdentification() *SecurityIdentification1Choice {
	f.Identification = new(SecurityIdentification1Choice)
	return f.Identification
}

func (f *FinancialInstrument4) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument4) SetSupplementaryIdentification(value string) {
	f.SupplementaryIdentification = (*Max35Text)(&value)
}

func (f *FinancialInstrument4) SetSecuritiesForm(value string) {
	f.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (f *FinancialInstrument4) SetClassType(value string) {
	f.ClassType = (*Max35Text)(&value)
}

func (f *FinancialInstrument4) SetDistributionPolicy(value string) {
	f.DistributionPolicy = (*DistributionPolicy1Code)(&value)
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument40 struct {

	// Unique and unambiguous identifier of a security, assigned under a formal or proprietary identification scheme.
	Identification *SecurityIdentification3Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Specifies whether the financial instrument is transferred as an asset or as cash.
	TransferType *TransferType1Code `xml:"TrfTp"`

	// Specifies the quantity of assets to be transferred in units or in a percentage rate.
	Quantity *Quantity14Choice `xml:"Qty"`

	// Average cost per share of a security, including all charges and commissions.
	AverageAcquisitionPrice *ActiveCurrencyAndAmount `xml:"AvrgAcqstnPric,omitempty"`

	// Identifies the currency used to transfer the holdings. Some transfer agents register holdings grouped by currency in addition to using the ISIN for multi-currency fund shares.
	TransferCurrency *CurrencyCode `xml:"TrfCcy,omitempty"`

	// Net asset on balance sheet - total portfolio value minus or plus the unrealised gain or loss.
	TotalBookValue *ActiveOrHistoricCurrencyAndAmount `xml:"TtlBookVal,omitempty"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	TransfereeAccount *Account16 `xml:"TrfeeAcct,omitempty"`

	// Sub-accounts that are grouped in a master or omnibus account.
	SubAccountDetails *SubAccount1 `xml:"SubAcctDtls,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, for example, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount93 `xml:"RcvgAgtDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount93 `xml:"DlvrgAgtDtls,omitempty"`

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`
}

func (f *FinancialInstrument40) AddIdentification() *SecurityIdentification3Choice {
	f.Identification = new(SecurityIdentification3Choice)
	return f.Identification
}

func (f *FinancialInstrument40) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument40) SetTransferType(value string) {
	f.TransferType = (*TransferType1Code)(&value)
}

func (f *FinancialInstrument40) AddQuantity() *Quantity14Choice {
	f.Quantity = new(Quantity14Choice)
	return f.Quantity
}

func (f *FinancialInstrument40) SetAverageAcquisitionPrice(value, currency string) {
	f.AverageAcquisitionPrice = NewActiveCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrument40) SetTransferCurrency(value string) {
	f.TransferCurrency = (*CurrencyCode)(&value)
}

func (f *FinancialInstrument40) SetTotalBookValue(value, currency string) {
	f.TotalBookValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrument40) AddTransfereeAccount() *Account16 {
	f.TransfereeAccount = new(Account16)
	return f.TransfereeAccount
}

func (f *FinancialInstrument40) AddSubAccountDetails() *SubAccount1 {
	f.SubAccountDetails = new(SubAccount1)
	return f.SubAccountDetails
}

func (f *FinancialInstrument40) AddReceivingAgentDetails() *PartyIdentificationAndAccount93 {
	f.ReceivingAgentDetails = new(PartyIdentificationAndAccount93)
	return f.ReceivingAgentDetails
}

func (f *FinancialInstrument40) AddDeliveringAgentDetails() *PartyIdentificationAndAccount93 {
	f.DeliveringAgentDetails = new(PartyIdentificationAndAccount93)
	return f.DeliveringAgentDetails
}

func (f *FinancialInstrument40) SetRequestedSettlementDate(value string) {
	f.RequestedSettlementDate = (*ISODate)(&value)
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, for example, dividend option or valuation currency.
type FinancialInstrument45 struct {

	// Identification of a security by an ISIN.
	Identification *SecurityIdentification23Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Financial Instrument Short Name (FISN) expressed in conformance with the ISO 18774 standard.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`

	// Additional information about a financial instrument to help identify the instrument.
	SupplementaryIdentification *Max35Text `xml:"SplmtryId,omitempty"`

	// Features of units offered by a fund. For example, a unit may have a specific load structure, for example, front end or back end, an income policy, for example, pay out or accumulate, or a trailer policy, for example, with or without. Fund classes are typically denoted by a single character, for example, 'Class A', 'Class 2'.
	ClassType *Max35Text `xml:"ClssTp,omitempty"`

	// Form, that is, ownership, of the security, for example, registered or bearer.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Income policy relating to a class type, that is, if income is paid out or retained in the fund.
	DistributionPolicy *DistributionPolicy1Code `xml:"DstrbtnPlcy,omitempty"`

	// Company specific description of a group of funds.
	ProductGroup *Max140Text `xml:"PdctGrp,omitempty"`
}

func (f *FinancialInstrument45) AddIdentification() *SecurityIdentification23Choice {
	f.Identification = new(SecurityIdentification23Choice)
	return f.Identification
}

func (f *FinancialInstrument45) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument45) SetShortName(value string) {
	f.ShortName = (*Max35Text)(&value)
}

func (f *FinancialInstrument45) SetSupplementaryIdentification(value string) {
	f.SupplementaryIdentification = (*Max35Text)(&value)
}

func (f *FinancialInstrument45) SetClassType(value string) {
	f.ClassType = (*Max35Text)(&value)
}

func (f *FinancialInstrument45) SetSecuritiesForm(value string) {
	f.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (f *FinancialInstrument45) SetDistributionPolicy(value string) {
	f.DistributionPolicy = (*DistributionPolicy1Code)(&value)
}

func (f *FinancialInstrument45) SetProductGroup(value string) {
	f.ProductGroup = (*Max140Text)(&value)
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, for example, dividend option or valuation currency.
type FinancialInstrument46 struct {

	// Unique and unambiguous identifier of a security, assigned under a formal or proprietary identification scheme.
	Identification *SecurityIdentification23Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Financial Instrument Short Name (FISN) expressed in conformance with the ISO 18774 standard.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`

	// Specifies whether the financial instrument is transferred as an asset or as cash.
	TransferType *TransferType1Code `xml:"TrfTp"`

	// Specifies the quantity of assets to be transferred in units or in a percentage rate.
	Quantity *Quantity12Choice `xml:"Qty,omitempty"`

	// Average cost per share of a security, including all charges and commissions.
	AverageAcquisitionPrice *ActiveOrHistoricCurrencyAndAmount `xml:"AvrgAcqstnPric,omitempty"`

	// Identifies the currency to be used to transfer the holdings. Some transfer agents register holdings grouped by currency in addition to using the ISIN for multi-currency fund shares.
	TransferCurrency *ActiveOrHistoricCurrencyCode `xml:"TrfCcy,omitempty"`

	// Net asset on balance sheet - total portfolio value minus or plus the unrealised gain or loss.
	TotalBookValue *ActiveOrHistoricCurrencyAndAmount `xml:"TtlBookVal,omitempty"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	TransfereeAccount *Account19 `xml:"TrfeeAcct,omitempty"`

	// Sub-accounts that are grouped in a master or omnibus account.
	SubAccountDetails *SubAccount5 `xml:"SubAcctDtls,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesReceivingSideDetails *ReceivingPartiesAndAccount14 `xml:"SttlmPtiesRcvgSdDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount125 `xml:"DlvrgAgtDtls,omitempty"`
}

func (f *FinancialInstrument46) AddIdentification() *SecurityIdentification23Choice {
	f.Identification = new(SecurityIdentification23Choice)
	return f.Identification
}

func (f *FinancialInstrument46) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument46) SetShortName(value string) {
	f.ShortName = (*Max35Text)(&value)
}

func (f *FinancialInstrument46) SetTransferType(value string) {
	f.TransferType = (*TransferType1Code)(&value)
}

func (f *FinancialInstrument46) AddQuantity() *Quantity12Choice {
	f.Quantity = new(Quantity12Choice)
	return f.Quantity
}

func (f *FinancialInstrument46) SetAverageAcquisitionPrice(value, currency string) {
	f.AverageAcquisitionPrice = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrument46) SetTransferCurrency(value string) {
	f.TransferCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrument46) SetTotalBookValue(value, currency string) {
	f.TotalBookValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrument46) AddTransfereeAccount() *Account19 {
	f.TransfereeAccount = new(Account19)
	return f.TransfereeAccount
}

func (f *FinancialInstrument46) AddSubAccountDetails() *SubAccount5 {
	f.SubAccountDetails = new(SubAccount5)
	return f.SubAccountDetails
}

func (f *FinancialInstrument46) AddSettlementPartiesReceivingSideDetails() *ReceivingPartiesAndAccount14 {
	f.SettlementPartiesReceivingSideDetails = new(ReceivingPartiesAndAccount14)
	return f.SettlementPartiesReceivingSideDetails
}

func (f *FinancialInstrument46) AddDeliveringAgentDetails() *PartyIdentificationAndAccount125 {
	f.DeliveringAgentDetails = new(PartyIdentificationAndAccount125)
	return f.DeliveringAgentDetails
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, for example, dividend option or valuation currency.
type FinancialInstrument47 struct {

	// Unique and unambiguous identifier of a security, assigned under a formal or proprietary identification scheme.
	Identification *SecurityIdentification23Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Financial Instrument Short Name (FISN) expressed in conformance with the ISO 18774 standard.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`

	// Specifies the quantity of assets to be transferred in units or in a percentage rate.
	Quantity *Quantity12Choice `xml:"Qty,omitempty"`

	// Average cost per share of a security, including all charges and commissions.
	AverageAcquisitionPrice *ActiveOrHistoricCurrencyAndAmount `xml:"AvrgAcqstnPric,omitempty"`

	// Net asset on balance sheet - total portfolio value minus or plus the unrealised gain or loss.
	TotalBookValue *ActiveOrHistoricCurrencyAndAmount `xml:"TtlBookVal,omitempty"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	TransfereeAccount *Account19 `xml:"TrfeeAcct,omitempty"`

	// Sub-accounts that are grouped in a master or omnibus account.
	SubAccountDetails *SubAccount5 `xml:"SubAcctDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount125 `xml:"DlvrgAgtDtls,omitempty"`
}

func (f *FinancialInstrument47) AddIdentification() *SecurityIdentification23Choice {
	f.Identification = new(SecurityIdentification23Choice)
	return f.Identification
}

func (f *FinancialInstrument47) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument47) SetShortName(value string) {
	f.ShortName = (*Max35Text)(&value)
}

func (f *FinancialInstrument47) AddQuantity() *Quantity12Choice {
	f.Quantity = new(Quantity12Choice)
	return f.Quantity
}

func (f *FinancialInstrument47) SetAverageAcquisitionPrice(value, currency string) {
	f.AverageAcquisitionPrice = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrument47) SetTotalBookValue(value, currency string) {
	f.TotalBookValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrument47) AddTransfereeAccount() *Account19 {
	f.TransfereeAccount = new(Account19)
	return f.TransfereeAccount
}

func (f *FinancialInstrument47) AddSubAccountDetails() *SubAccount5 {
	f.SubAccountDetails = new(SubAccount5)
	return f.SubAccountDetails
}

func (f *FinancialInstrument47) AddDeliveringAgentDetails() *PartyIdentificationAndAccount125 {
	f.DeliveringAgentDetails = new(PartyIdentificationAndAccount125)
	return f.DeliveringAgentDetails
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, for example, dividend option or valuation currency.
type FinancialInstrument48 struct {

	// Unique and unambiguous identifier of a security, assigned under a formal or proprietary identification scheme.
	Identification *SecurityIdentification23Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Financial Instrument Short Name (FISN) expressed in conformance with the ISO 18774 standard.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`

	// Specifies whether the financial instrument is transferred as an asset or as cash.
	TransferType *TransferType1Code `xml:"TrfTp"`

	// Specifies the quantity of assets to be transferred in units or in a percentage rate.
	Quantity *Quantity14Choice `xml:"Qty"`

	// Average cost per share of a security, including all charges and commissions.
	AverageAcquisitionPrice *ActiveCurrencyAndAmount `xml:"AvrgAcqstnPric,omitempty"`

	// Identifies the currency used to transfer the holdings. Some transfer agents register holdings grouped by currency in addition to using the ISIN for multi-currency fund shares.
	TransferCurrency *ActiveOrHistoricCurrencyCode `xml:"TrfCcy,omitempty"`

	// Net asset on balance sheet - total portfolio value minus or plus the unrealised gain or loss.
	TotalBookValue *ActiveOrHistoricCurrencyAndAmount `xml:"TtlBookVal,omitempty"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	TransfereeAccount *Account19 `xml:"TrfeeAcct,omitempty"`

	// Sub-accounts that are grouped in a master or omnibus account.
	SubAccountDetails *SubAccount5 `xml:"SubAcctDtls,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesReceivingSideDetails *ReceivingPartiesAndAccount14 `xml:"SttlmPtiesRcvgSdDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount125 `xml:"DlvrgAgtDtls,omitempty"`

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`
}

func (f *FinancialInstrument48) AddIdentification() *SecurityIdentification23Choice {
	f.Identification = new(SecurityIdentification23Choice)
	return f.Identification
}

func (f *FinancialInstrument48) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument48) SetShortName(value string) {
	f.ShortName = (*Max35Text)(&value)
}

func (f *FinancialInstrument48) SetTransferType(value string) {
	f.TransferType = (*TransferType1Code)(&value)
}

func (f *FinancialInstrument48) AddQuantity() *Quantity14Choice {
	f.Quantity = new(Quantity14Choice)
	return f.Quantity
}

func (f *FinancialInstrument48) SetAverageAcquisitionPrice(value, currency string) {
	f.AverageAcquisitionPrice = NewActiveCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrument48) SetTransferCurrency(value string) {
	f.TransferCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrument48) SetTotalBookValue(value, currency string) {
	f.TotalBookValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrument48) AddTransfereeAccount() *Account19 {
	f.TransfereeAccount = new(Account19)
	return f.TransfereeAccount
}

func (f *FinancialInstrument48) AddSubAccountDetails() *SubAccount5 {
	f.SubAccountDetails = new(SubAccount5)
	return f.SubAccountDetails
}

func (f *FinancialInstrument48) AddSettlementPartiesReceivingSideDetails() *ReceivingPartiesAndAccount14 {
	f.SettlementPartiesReceivingSideDetails = new(ReceivingPartiesAndAccount14)
	return f.SettlementPartiesReceivingSideDetails
}

func (f *FinancialInstrument48) AddDeliveringAgentDetails() *PartyIdentificationAndAccount125 {
	f.DeliveringAgentDetails = new(PartyIdentificationAndAccount125)
	return f.DeliveringAgentDetails
}

func (f *FinancialInstrument48) SetRequestedSettlementDate(value string) {
	f.RequestedSettlementDate = (*ISODate)(&value)
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, for example, dividend option or valuation currency.
type FinancialInstrument49 struct {

	// Identification of a security by an ISIN.
	Identification *SecurityIdentification23Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Financial Instrument Short Name (FISN) expressed in conformance with the ISO 18774 standard.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`

	// Additional information about a financial instrument to help identify the instrument.
	SupplementaryIdentification *Max35Text `xml:"SplmtryId,omitempty"`

	// Features of units offered by a fund. For example, a unit may have a specific load structure, for example, front end or back end, an income policy, for example, pay out or accumulate, or a trailer policy, for example, with or without. Fund classes are typically denoted by a single character, for example, 'Class A', 'Class 2'.
	ClassType *Max35Text `xml:"ClssTp,omitempty"`

	// Form, that is, ownership, of the security, for example, registered or bearer.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Income policy relating to a class type, that is, if income is paid out or retained in the fund.
	DistributionPolicy *DistributionPolicy1Code `xml:"DstrbtnPlcy,omitempty"`
}

func (f *FinancialInstrument49) AddIdentification() *SecurityIdentification23Choice {
	f.Identification = new(SecurityIdentification23Choice)
	return f.Identification
}

func (f *FinancialInstrument49) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument49) SetShortName(value string) {
	f.ShortName = (*Max35Text)(&value)
}

func (f *FinancialInstrument49) SetSupplementaryIdentification(value string) {
	f.SupplementaryIdentification = (*Max35Text)(&value)
}

func (f *FinancialInstrument49) SetClassType(value string) {
	f.ClassType = (*Max35Text)(&value)
}

func (f *FinancialInstrument49) SetSecuritiesForm(value string) {
	f.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (f *FinancialInstrument49) SetDistributionPolicy(value string) {
	f.DistributionPolicy = (*DistributionPolicy1Code)(&value)
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument5 struct {

	// Identification of a security by an ISIN.
	Identification *SecurityIdentification1Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Additional information about a financial instrument to help identify the instrument.
	SupplementaryIdentification *Max35Text `xml:"SplmtryId,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Features of units offered by a fund. For example, a unit may have a specific load structure, eg, front end or back end, an income policy, eg, pay out or accumulate, or a trailer policy, eg, with or without. Fund classes are typically denoted by a single character, eg, 'Class A', 'Class 2'.
	ClassType *Max35Text `xml:"ClssTp,omitempty"`

	// Form, ie, ownership, of the security, eg, registered or bearer.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Income policy relating to a class type, ie, if income is paid out or retained in the fund.
	DistributionPolicy *DistributionPolicy1Code `xml:"DstrbtnPlcy,omitempty"`

	// Indicates whether the fund has two prices.
	DualFundIndicator *YesNoIndicator `xml:"DualFndInd"`
}

func (f *FinancialInstrument5) AddIdentification() *SecurityIdentification1Choice {
	f.Identification = new(SecurityIdentification1Choice)
	return f.Identification
}

func (f *FinancialInstrument5) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument5) SetSupplementaryIdentification(value string) {
	f.SupplementaryIdentification = (*Max35Text)(&value)
}

func (f *FinancialInstrument5) SetRequestedNAVCurrency(value string) {
	f.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrument5) SetClassType(value string) {
	f.ClassType = (*Max35Text)(&value)
}

func (f *FinancialInstrument5) SetSecuritiesForm(value string) {
	f.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (f *FinancialInstrument5) SetDistributionPolicy(value string) {
	f.DistributionPolicy = (*DistributionPolicy1Code)(&value)
}

func (f *FinancialInstrument5) SetDualFundIndicator(value string) {
	f.DualFundIndicator = (*YesNoIndicator)(&value)
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, for example, dividend option or valuation currency.
type FinancialInstrument50 struct {

	// Unique and unambiguous identifier of a security, assigned under a formal or proprietary identification scheme.
	Identification *SecurityIdentification23Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Financial Instrument Short Name (FISN) expressed in conformance with the ISO 18774 standard.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	TransfereeAccount *Account19 `xml:"TrfeeAcct,omitempty"`

	// Sub-accounts that are grouped in a master or omnibus account.
	SubAccountDetails *SubAccount5 `xml:"SubAcctDtls,omitempty"`
}

func (f *FinancialInstrument50) AddIdentification() *SecurityIdentification23Choice {
	f.Identification = new(SecurityIdentification23Choice)
	return f.Identification
}

func (f *FinancialInstrument50) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument50) SetShortName(value string) {
	f.ShortName = (*Max35Text)(&value)
}

func (f *FinancialInstrument50) AddTransfereeAccount() *Account19 {
	f.TransfereeAccount = new(Account19)
	return f.TransfereeAccount
}

func (f *FinancialInstrument50) AddSubAccountDetails() *SubAccount5 {
	f.SubAccountDetails = new(SubAccount5)
	return f.SubAccountDetails
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, for example, dividend option or valuation currency.
type FinancialInstrument51 struct {

	// Identification of the security by an ISIN.
	Identification *SecurityIdentification23Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Financial Instrument Short Name (FISN) expressed in conformance with the ISO 18774 standard.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`

	// Additional information about a financial instrument to help identify the instrument.
	SupplementaryIdentification *Max35Text `xml:"SplmtryId,omitempty"`

	// Features of units offered by a fund. For example, a unit may have a specific load structure, for example, front end or back end, an income policy,  for example, pay out or accumulate, or a trailer policy,  for example, with or without. Fund classes are typically denoted by a single character,  for example, 'Class A', 'Class 2'.
	ClassType *Max35Text `xml:"ClssTp,omitempty"`

	// Form of ownership, that is registered or bearer.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Income policy relating to a class type, that is, if income is paid out or retained in the fund.
	DistributionPolicy *DistributionPolicy1Code `xml:"DstrbtnPlcy,omitempty"`

	// Company specific description of a group of funds.
	ProductGroup *Max140Text `xml:"PdctGrp,omitempty"`

	// When an account at fund or security level is blocked, this specifies details on how the holding is blocked.
	BlockedHoldingDetails *BlockedHoldingDetails2 `xml:"BlckdHldgDtls,omitempty"`

	// Specifies whether the holdings in the account are eligible for pledging.
	Pledging *Eligible1Code `xml:"Pldgg,omitempty"`

	// Specifies whether the holdings in the account are used as collateral.
	Collateral *Collateral1Code `xml:"Coll,omitempty"`

	// Details of third party rights.
	ThirdPartyRights *ThirdPartyRights1 `xml:"ThrdPtyRghts,omitempty"`

	// Specifies if all the shares are owned exclusively by the fund company.
	FundOwnership *FundOwnership1Code `xml:"FndOwnrsh,omitempty"`

	// Specifies if the fund is intended for qualified investors.
	FundIntention *FundIntention1Code `xml:"FndIntntn,omitempty"`

	// Operational status of the fund.
	OperationalStatus *OperationalStatus1Code `xml:"OprlSts,omitempty"`
}

func (f *FinancialInstrument51) AddIdentification() *SecurityIdentification23Choice {
	f.Identification = new(SecurityIdentification23Choice)
	return f.Identification
}

func (f *FinancialInstrument51) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument51) SetShortName(value string) {
	f.ShortName = (*Max35Text)(&value)
}

func (f *FinancialInstrument51) SetSupplementaryIdentification(value string) {
	f.SupplementaryIdentification = (*Max35Text)(&value)
}

func (f *FinancialInstrument51) SetClassType(value string) {
	f.ClassType = (*Max35Text)(&value)
}

func (f *FinancialInstrument51) SetSecuritiesForm(value string) {
	f.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (f *FinancialInstrument51) SetDistributionPolicy(value string) {
	f.DistributionPolicy = (*DistributionPolicy1Code)(&value)
}

func (f *FinancialInstrument51) SetProductGroup(value string) {
	f.ProductGroup = (*Max140Text)(&value)
}

func (f *FinancialInstrument51) AddBlockedHoldingDetails() *BlockedHoldingDetails2 {
	f.BlockedHoldingDetails = new(BlockedHoldingDetails2)
	return f.BlockedHoldingDetails
}

func (f *FinancialInstrument51) SetPledging(value string) {
	f.Pledging = (*Eligible1Code)(&value)
}

func (f *FinancialInstrument51) SetCollateral(value string) {
	f.Collateral = (*Collateral1Code)(&value)
}

func (f *FinancialInstrument51) AddThirdPartyRights() *ThirdPartyRights1 {
	f.ThirdPartyRights = new(ThirdPartyRights1)
	return f.ThirdPartyRights
}

func (f *FinancialInstrument51) SetFundOwnership(value string) {
	f.FundOwnership = (*FundOwnership1Code)(&value)
}

func (f *FinancialInstrument51) SetFundIntention(value string) {
	f.FundIntention = (*FundIntention1Code)(&value)
}

func (f *FinancialInstrument51) SetOperationalStatus(value string) {
	f.OperationalStatus = (*OperationalStatus1Code)(&value)
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, for example, dividend option or valuation currency.
type FinancialInstrument55 struct {

	// Identification of the security by an ISIN.
	Identification *SecurityIdentification25Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Financial Instrument Short Name (FISN) expressed in conformance with the ISO 18774 standard.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`

	// Additional information about the financial instrument to help identify the instrument.
	SupplementaryIdentification *Max35Text `xml:"SplmtryId,omitempty"`

	// Features of units offered by the fund. For example, a unit may have a specific load structure, for example, front end or back end, an income policy, for example, pay out or accumulate, or a trailer policy, for example, with or without. Fund classes are typically denoted by a single character, for example, 'Class A', 'Class 2'.
	ClassType *Max35Text `xml:"ClssTp,omitempty"`

	// Form, that is, ownership, of the security, for example, registered or bearer.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Income policy relating to the class type, that is, if income is paid out or retained in the fund.
	DistributionPolicy *DistributionPolicy1Code `xml:"DstrbtnPlcy,omitempty"`

	// Company specific description of a group of funds.
	ProductGroup *Max140Text `xml:"PdctGrp,omitempty"`
}

func (f *FinancialInstrument55) AddIdentification() *SecurityIdentification25Choice {
	f.Identification = new(SecurityIdentification25Choice)
	return f.Identification
}

func (f *FinancialInstrument55) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument55) SetShortName(value string) {
	f.ShortName = (*Max35Text)(&value)
}

func (f *FinancialInstrument55) SetSupplementaryIdentification(value string) {
	f.SupplementaryIdentification = (*Max35Text)(&value)
}

func (f *FinancialInstrument55) SetClassType(value string) {
	f.ClassType = (*Max35Text)(&value)
}

func (f *FinancialInstrument55) SetSecuritiesForm(value string) {
	f.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (f *FinancialInstrument55) SetDistributionPolicy(value string) {
	f.DistributionPolicy = (*DistributionPolicy1Code)(&value)
}

func (f *FinancialInstrument55) SetProductGroup(value string) {
	f.ProductGroup = (*Max140Text)(&value)
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, for example, dividend option or valuation currency.
type FinancialInstrument56 struct {

	// Identification of the security by an ISIN.
	Identification *SecurityIdentification25Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Financial Instrument Short Name (FISN) expressed in conformance with the ISO 18774 standard.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`

	// Additional information about the financial instrument to help identify the instrument.
	SupplementaryIdentification *Max35Text `xml:"SplmtryId,omitempty"`

	// Features of units offered by the fund. For example, a unit may have a specific load structure, for example, front end or back end, an income policy,  for example, pay out or accumulate, or a trailer policy,  for example, with or without. Fund classes are typically denoted by a single character,  for example, 'Class A', 'Class 2'.
	ClassType *Max35Text `xml:"ClssTp,omitempty"`

	// Form of ownership, that is registered or bearer.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Income policy relating to the class type, that is, if income is paid out or retained in the fund.
	DistributionPolicy *DistributionPolicy1Code `xml:"DstrbtnPlcy,omitempty"`

	// Company specific description of a group of funds.
	ProductGroup *Max140Text `xml:"PdctGrp,omitempty"`

	// When an account at fund or security level is blocked, this specifies details on how the holding is blocked.
	BlockedHoldingDetails *BlockedHoldingDetails2 `xml:"BlckdHldgDtls,omitempty"`

	// Specifies whether the holdings in the account are eligible for pledging.
	Pledging *Eligible1Code `xml:"Pldgg,omitempty"`

	// Specifies whether the holdings in the account are used as collateral.
	Collateral *Collateral1Code `xml:"Coll,omitempty"`

	// Details of third party rights.
	ThirdPartyRights *ThirdPartyRights1 `xml:"ThrdPtyRghts,omitempty"`

	// Specifies if all the shares are owned exclusively by the fund company.
	FundOwnership *FundOwnership1Code `xml:"FndOwnrsh,omitempty"`

	// Specifies if the fund is intended for qualified investors.
	FundIntention *FundIntention1Code `xml:"FndIntntn,omitempty"`

	// Operational status of the fund.
	OperationalStatus *OperationalStatus1Code `xml:"OprlSts,omitempty"`
}

func (f *FinancialInstrument56) AddIdentification() *SecurityIdentification25Choice {
	f.Identification = new(SecurityIdentification25Choice)
	return f.Identification
}

func (f *FinancialInstrument56) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument56) SetShortName(value string) {
	f.ShortName = (*Max35Text)(&value)
}

func (f *FinancialInstrument56) SetSupplementaryIdentification(value string) {
	f.SupplementaryIdentification = (*Max35Text)(&value)
}

func (f *FinancialInstrument56) SetClassType(value string) {
	f.ClassType = (*Max35Text)(&value)
}

func (f *FinancialInstrument56) SetSecuritiesForm(value string) {
	f.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (f *FinancialInstrument56) SetDistributionPolicy(value string) {
	f.DistributionPolicy = (*DistributionPolicy1Code)(&value)
}

func (f *FinancialInstrument56) SetProductGroup(value string) {
	f.ProductGroup = (*Max140Text)(&value)
}

func (f *FinancialInstrument56) AddBlockedHoldingDetails() *BlockedHoldingDetails2 {
	f.BlockedHoldingDetails = new(BlockedHoldingDetails2)
	return f.BlockedHoldingDetails
}

func (f *FinancialInstrument56) SetPledging(value string) {
	f.Pledging = (*Eligible1Code)(&value)
}

func (f *FinancialInstrument56) SetCollateral(value string) {
	f.Collateral = (*Collateral1Code)(&value)
}

func (f *FinancialInstrument56) AddThirdPartyRights() *ThirdPartyRights1 {
	f.ThirdPartyRights = new(ThirdPartyRights1)
	return f.ThirdPartyRights
}

func (f *FinancialInstrument56) SetFundOwnership(value string) {
	f.FundOwnership = (*FundOwnership1Code)(&value)
}

func (f *FinancialInstrument56) SetFundIntention(value string) {
	f.FundIntention = (*FundIntention1Code)(&value)
}

func (f *FinancialInstrument56) SetOperationalStatus(value string) {
	f.OperationalStatus = (*OperationalStatus1Code)(&value)
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, for example,, dividend option or valuation currency.
type FinancialInstrument57 struct {

	// Identification of the security by an ISIN.
	Identification *SecurityIdentification25Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Financial Instrument Short Name (FISN) expressed in conformance with the ISO 18774 standard.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`

	// Additional information about the financial instrument to help identify the instrument.
	SupplementaryIdentification *Max35Text `xml:"SplmtryId,omitempty"`

	// Features of units offered by the fund. For example, a unit may have a specific load structure, for example, front end or back end, an income policy, for example, pay out or accumulate, or a trailer policy, for example, with or without. Fund classes are typically denoted by a single character, for example, 'Class A', 'Class 2'.
	ClassType *Max35Text `xml:"ClssTp,omitempty"`

	// Form, that is, ownership, of the security, for example, registered or bearer.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Income policy relating to a class type, that is, if income is paid out or retained in the fund.
	DistributionPolicy *DistributionPolicy1Code `xml:"DstrbtnPlcy,omitempty"`

	// Company specific description of a group of funds.
	ProductGroup *Max140Text `xml:"PdctGrp,omitempty"`

	// Choice of formats for the identification of a series.
	SeriesIdentification *Series1 `xml:"SrsId,omitempty"`
}

func (f *FinancialInstrument57) AddIdentification() *SecurityIdentification25Choice {
	f.Identification = new(SecurityIdentification25Choice)
	return f.Identification
}

func (f *FinancialInstrument57) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument57) SetShortName(value string) {
	f.ShortName = (*Max35Text)(&value)
}

func (f *FinancialInstrument57) SetSupplementaryIdentification(value string) {
	f.SupplementaryIdentification = (*Max35Text)(&value)
}

func (f *FinancialInstrument57) SetClassType(value string) {
	f.ClassType = (*Max35Text)(&value)
}

func (f *FinancialInstrument57) SetSecuritiesForm(value string) {
	f.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (f *FinancialInstrument57) SetDistributionPolicy(value string) {
	f.DistributionPolicy = (*DistributionPolicy1Code)(&value)
}

func (f *FinancialInstrument57) SetProductGroup(value string) {
	f.ProductGroup = (*Max140Text)(&value)
}

func (f *FinancialInstrument57) AddSeriesIdentification() *Series1 {
	f.SeriesIdentification = new(Series1)
	return f.SeriesIdentification
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument6 struct {

	// Identification of a security by an ISIN.
	Identification *SecurityIdentification1Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Additional information about a financial instrument to help identify the instrument.
	SupplementaryIdentification *Max35Text `xml:"SplmtryId,omitempty"`

	// Features of units offered by a fund. For example, a unit may have a specific load structure, eg, front end or back end, an income policy, eg, pay out or accumulate, or a trailer policy, eg, with or without. Fund classes are typically denoted by a single character, eg, 'Class A', 'Class 2'.
	ClassType *Max35Text `xml:"ClssTp,omitempty"`

	// Form, ie, ownership, of the security, eg, registered or bearer.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Income policy relating to a class type, ie, if income is paid out or retained in the fund.
	DistributionPolicy *DistributionPolicy1Code `xml:"DstrbtnPlcy,omitempty"`

	// Company specific description of a group of funds.
	ProductGroup *Max140Text `xml:"PdctGrp,omitempty"`
}

func (f *FinancialInstrument6) AddIdentification() *SecurityIdentification1Choice {
	f.Identification = new(SecurityIdentification1Choice)
	return f.Identification
}

func (f *FinancialInstrument6) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument6) SetSupplementaryIdentification(value string) {
	f.SupplementaryIdentification = (*Max35Text)(&value)
}

func (f *FinancialInstrument6) SetClassType(value string) {
	f.ClassType = (*Max35Text)(&value)
}

func (f *FinancialInstrument6) SetSecuritiesForm(value string) {
	f.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (f *FinancialInstrument6) SetDistributionPolicy(value string) {
	f.DistributionPolicy = (*DistributionPolicy1Code)(&value)
}

func (f *FinancialInstrument6) SetProductGroup(value string) {
	f.ProductGroup = (*Max140Text)(&value)
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument8 struct {

	// Identification of a security by an ISIN.
	Identification []*SecurityIdentification3Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Additional information about a financial instrument to help identify the instrument.
	SupplementaryIdentification *Max35Text `xml:"SplmtryId,omitempty"`

	// Currency in which a security is issued or redenominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Features of units offered by a fund. For example, a unit may have a specific load structure, eg, front end or back end, an income policy, eg, pay out or accumulate, or a trailer policy, eg, with or without. Fund classes are typically denoted by a single character, eg, 'Class A', 'Class 2'.
	ClassType *Max35Text `xml:"ClssTp,omitempty"`

	// Form, ie, ownership, of the security, eg, registered or bearer.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Income policy relating to a class type, ie, if income is paid out or retained in the fund.
	DistributionPolicy *DistributionPolicy1Code `xml:"DstrbtnPlcy,omitempty"`

	// Indicates whether the fund has two prices.
	DualFundIndicator *YesNoIndicator `xml:"DualFndInd"`
}

func (f *FinancialInstrument8) AddIdentification() *SecurityIdentification3Choice {
	newValue := new(SecurityIdentification3Choice)
	f.Identification = append(f.Identification, newValue)
	return newValue
}

func (f *FinancialInstrument8) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument8) SetSupplementaryIdentification(value string) {
	f.SupplementaryIdentification = (*Max35Text)(&value)
}

func (f *FinancialInstrument8) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrument8) SetClassType(value string) {
	f.ClassType = (*Max35Text)(&value)
}

func (f *FinancialInstrument8) SetSecuritiesForm(value string) {
	f.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (f *FinancialInstrument8) SetDistributionPolicy(value string) {
	f.DistributionPolicy = (*DistributionPolicy1Code)(&value)
}

func (f *FinancialInstrument8) SetDualFundIndicator(value string) {
	f.DualFundIndicator = (*YesNoIndicator)(&value)
}

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument9 struct {

	// Identification of a security by an ISIN.
	Identification *SecurityIdentification3Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Additional information about a financial instrument to help identify the instrument.
	SupplementaryIdentification *Max35Text `xml:"SplmtryId,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Features of units offered by a fund. For example, a unit may have a specific load structure, eg, front end or back end, an income policy, eg, pay out or accumulate, or a trailer policy, eg, with or without. Fund classes are typically denoted by a single character, eg, 'Class A', 'Class 2'.
	ClassType *Max35Text `xml:"ClssTp,omitempty"`

	// Form, ie, ownership, of the security, eg, registered or bearer.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Income policy relating to a class type, ie, if income is paid out or retained in the fund.
	DistributionPolicy *DistributionPolicy1Code `xml:"DstrbtnPlcy,omitempty"`

	// Indicates whether the fund has two prices.
	DualFundIndicator *YesNoIndicator `xml:"DualFndInd"`
}

func (f *FinancialInstrument9) AddIdentification() *SecurityIdentification3Choice {
	f.Identification = new(SecurityIdentification3Choice)
	return f.Identification
}

func (f *FinancialInstrument9) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument9) SetSupplementaryIdentification(value string) {
	f.SupplementaryIdentification = (*Max35Text)(&value)
}

func (f *FinancialInstrument9) SetRequestedNAVCurrency(value string) {
	f.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrument9) SetClassType(value string) {
	f.ClassType = (*Max35Text)(&value)
}

func (f *FinancialInstrument9) SetSecuritiesForm(value string) {
	f.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (f *FinancialInstrument9) SetDistributionPolicy(value string) {
	f.DistributionPolicy = (*DistributionPolicy1Code)(&value)
}

func (f *FinancialInstrument9) SetDualFundIndicator(value string) {
	f.DualFundIndicator = (*YesNoIndicator)(&value)
}
