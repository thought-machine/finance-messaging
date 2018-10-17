package iso20022

// Characteristics of the ownership of an investment account.
type InvestmentAccountOwnershipInformation10 struct {

	// Information about the organisation or individual person.
	Party *Party23Choice `xml:"Pty"`

	// Status of an identity check to prevent money laundering. This includes the counter-terrorism check.
	MoneyLaunderingCheck *MoneyLaunderingCheck1Choice `xml:"MnyLndrgChck,omitempty"`

	// Information to support Know Your Customer processes.
	InvestorProfileValidation []*PartyProfileInformation4 `xml:"InvstrPrflVldtn,omitempty"`

	// Percentage of ownership or of beneficial ownership of the shares/units in the account. All subsequent subscriptions and or redemptions will be allocated using the same percentage.
	OwnershipBeneficiaryRate *PercentageRate `xml:"OwnrshBnfcryRate,omitempty"`

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	ClientIdentification *Max35Text `xml:"ClntId,omitempty"`

	// Indicates whether an owner of an investment account may benefit from a fiscal exemption or amnesty for instance for declaring overseas investments.
	FiscalExemption *YesNoIndicator `xml:"FsclXmptn,omitempty"`

	// Indicates whether the signature of the account owner is required to authorise transactions on the account.
	SignatoryRightIndicator *YesNoIndicator `xml:"SgntryRghtInd,omitempty"`

	// Details about the MiFID classification of the account owner.
	MiFIDClassification *MiFIDClassification1 `xml:"MiFIDClssfctn,omitempty"`

	// Specifies how information is sent to the account holder.
	InformationDistribution *InformationDistribution1Code `xml:"InfDstrbtn,omitempty"`

	// Type of Foreign Account Tax Compliance Act (FATCA) form submitted by the investor.
	FATCAFormType []*FATCAForm1Choice `xml:"FATCAFormTp,omitempty"`

	// Foreign Account Tax Compliance Act (FATCA) status of the investor.
	FATCAStatus []*FATCAStatus1 `xml:"FATCASts,omitempty"`
}

func (i *InvestmentAccountOwnershipInformation10) AddParty() *Party23Choice {
	i.Party = new(Party23Choice)
	return i.Party
}

func (i *InvestmentAccountOwnershipInformation10) AddMoneyLaunderingCheck() *MoneyLaunderingCheck1Choice {
	i.MoneyLaunderingCheck = new(MoneyLaunderingCheck1Choice)
	return i.MoneyLaunderingCheck
}

func (i *InvestmentAccountOwnershipInformation10) AddInvestorProfileValidation() *PartyProfileInformation4 {
	newValue := new(PartyProfileInformation4)
	i.InvestorProfileValidation = append(i.InvestorProfileValidation, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation10) SetOwnershipBeneficiaryRate(value string) {
	i.OwnershipBeneficiaryRate = (*PercentageRate)(&value)
}

func (i *InvestmentAccountOwnershipInformation10) SetClientIdentification(value string) {
	i.ClientIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccountOwnershipInformation10) SetFiscalExemption(value string) {
	i.FiscalExemption = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountOwnershipInformation10) SetSignatoryRightIndicator(value string) {
	i.SignatoryRightIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountOwnershipInformation10) AddMiFIDClassification() *MiFIDClassification1 {
	i.MiFIDClassification = new(MiFIDClassification1)
	return i.MiFIDClassification
}

func (i *InvestmentAccountOwnershipInformation10) SetInformationDistribution(value string) {
	i.InformationDistribution = (*InformationDistribution1Code)(&value)
}

func (i *InvestmentAccountOwnershipInformation10) AddFATCAFormType() *FATCAForm1Choice {
	newValue := new(FATCAForm1Choice)
	i.FATCAFormType = append(i.FATCAFormType, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation10) AddFATCAStatus() *FATCAStatus1 {
	newValue := new(FATCAStatus1)
	i.FATCAStatus = append(i.FATCAStatus, newValue)
	return newValue
}

// Characteristics of the ownership of an investment account.
type InvestmentAccountOwnershipInformation11 struct {

	// Information about the organisation or individual person.
	Party *Party24Choice `xml:"Pty"`

	// Status of an identity check to prevent money laundering. This includes the counter-terrorism check.
	MoneyLaunderingCheck *MoneyLaunderingCheck1Choice `xml:"MnyLndrgChck,omitempty"`

	// Percentage of ownership or beneficiary ownership of the shares/units in the account. All subsequent subscriptions and or redemptions will be allocated using the same percentage.
	OwnershipBeneficiaryRate *PercentageRate `xml:"OwnrshBnfcryRate,omitempty"`

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	ClientIdentification *Max35Text `xml:"ClntId,omitempty"`

	// Indicates whether an owner of an investment account may benefit from a fiscal exemption or amnesty for instance for declaring overseas investments.
	FiscalExemption *YesNoIndicator `xml:"FsclXmptn,omitempty"`

	// Indicates whether the account owner signature is required to authorise transactions on the account.
	SignatoryRightIndicator *YesNoIndicator `xml:"SgntryRghtInd,omitempty"`

	// Information related to the party profile to be inserted or deleted.
	ModifiedInvestorProfileValidation []*ModificationScope19 `xml:"ModfdInvstrPrflVldtn,omitempty"`

	// Details about the MiFID classification of the account owner.
	MiFIDClassification *MiFIDClassification1 `xml:"MiFIDClssfctn,omitempty"`

	// Specifies how information is sent to the account holder.
	InformationDistribution *InformationDistribution1Code `xml:"InfDstrbtn,omitempty"`

	// Type of Foreign Account Tax Compliance Act (FATCA) form submitted by the investor.
	FATCAFormType []*FATCAForm1Choice `xml:"FATCAFormTp,omitempty"`

	// Foreign Account Tax Compliance Act (FATCA) status of the investor.
	FATCAStatus []*FATCAStatus1 `xml:"FATCASts,omitempty"`
}

func (i *InvestmentAccountOwnershipInformation11) AddParty() *Party24Choice {
	i.Party = new(Party24Choice)
	return i.Party
}

func (i *InvestmentAccountOwnershipInformation11) AddMoneyLaunderingCheck() *MoneyLaunderingCheck1Choice {
	i.MoneyLaunderingCheck = new(MoneyLaunderingCheck1Choice)
	return i.MoneyLaunderingCheck
}

func (i *InvestmentAccountOwnershipInformation11) SetOwnershipBeneficiaryRate(value string) {
	i.OwnershipBeneficiaryRate = (*PercentageRate)(&value)
}

func (i *InvestmentAccountOwnershipInformation11) SetClientIdentification(value string) {
	i.ClientIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccountOwnershipInformation11) SetFiscalExemption(value string) {
	i.FiscalExemption = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountOwnershipInformation11) SetSignatoryRightIndicator(value string) {
	i.SignatoryRightIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountOwnershipInformation11) AddModifiedInvestorProfileValidation() *ModificationScope19 {
	newValue := new(ModificationScope19)
	i.ModifiedInvestorProfileValidation = append(i.ModifiedInvestorProfileValidation, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation11) AddMiFIDClassification() *MiFIDClassification1 {
	i.MiFIDClassification = new(MiFIDClassification1)
	return i.MiFIDClassification
}

func (i *InvestmentAccountOwnershipInformation11) SetInformationDistribution(value string) {
	i.InformationDistribution = (*InformationDistribution1Code)(&value)
}

func (i *InvestmentAccountOwnershipInformation11) AddFATCAFormType() *FATCAForm1Choice {
	newValue := new(FATCAForm1Choice)
	i.FATCAFormType = append(i.FATCAFormType, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation11) AddFATCAStatus() *FATCAStatus1 {
	newValue := new(FATCAStatus1)
	i.FATCAStatus = append(i.FATCAStatus, newValue)
	return newValue
}

// Characteristics of the ownership of a securities account.
type InvestmentAccountOwnershipInformation12 struct {

	// Information about the organisation or individual person.
	Party *Party31Choice `xml:"Pty"`

	// Status of an identity check to prevent money laundering. This includes the counter-terrorism check.
	MoneyLaunderingCheck *MoneyLaunderingCheck1Choice `xml:"MnyLndrgChck,omitempty"`

	// Information to support Know Your Customer processes.
	InvestorProfileValidation []*PartyProfileInformation5 `xml:"InvstrPrflVldtn,omitempty"`

	// Percentage of ownership or of beneficial ownership of the shares/units in the account. All subsequent subscriptions or purchases and or redemptions or sells  will be allocated using the same percentage.
	OwnershipBeneficiaryRate *PercentageRate `xml:"OwnrshBnfcryRate,omitempty"`

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	ClientIdentification *Max35Text `xml:"ClntId,omitempty"`

	// Indicates whether an owner of the account may benefit from a fiscal exemption or amnesty, for example, when declaring overseas investments.
	FiscalExemption *YesNoIndicator `xml:"FsclXmptn,omitempty"`

	// Indicates whether the signature of the account owner is required to authorise transactions on the account.
	SignatoryRightIndicator *YesNoIndicator `xml:"SgntryRghtInd,omitempty"`

	// Details about the MiFID classification of the account owner.
	MiFIDClassification *MiFIDClassification1 `xml:"MiFIDClssfctn,omitempty"`

	// Type of information that must be provided to the account holder.
	Notification []*Notification2 `xml:"Ntfctn,omitempty"`

	// Type of Foreign Account Tax Compliance Act (FATCA) form submitted by the investor or account owner.
	FATCAFormType []*FATCAForm1Choice `xml:"FATCAFormTp,omitempty"`

	// Foreign Account Tax Compliance Act (FATCA) status of the investor or account owner.
	FATCAStatus []*FATCAStatus2 `xml:"FATCASts,omitempty"`

	// Alternative identification, for example, national registration identification number, passport number, tax identification number. This may be an account number used to further identify the beneficial owner, for example, a Central Provident Fund (CFP) account as required for Singapore.
	OtherIdentification []*GenericIdentification82 `xml:"OthrId,omitempty"`

	// Tax advantage specific to the account party.
	TaxExemption *TaxExemptionReason2Choice `xml:"TaxXmptn,omitempty"`

	// Details for the reporting of tax, for example, the country of taxation.
	TaxReporting []*TaxReporting1 `xml:"TaxRptg,omitempty"`

	// Language in which the organisation or person communicates.
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Method used for postal mailing.
	MailType *MailType1Choice `xml:"MailTp,omitempty"`

	// Country and residential status of the organisation or individual person.
	CountryAndResidentialStatus *CountryAndResidentialStatusType2 `xml:"CtryAndResdtlSts,omitempty"`

	// Annual wealth of the individual person or share capital value of the legal entity and date on which the annual wealth of the individual person was registered or declared or the date the stock value of the organisation was registered.
	MonetaryWealth *DateAndAmount1 `xml:"MntryWlth,omitempty"`

	// Amount of total assets minus liabilities of the individual person or the amount of the difference between assets and liabilities plus rights over obligations (net equity) of the organisation and the date on which the equity value was registered.
	EquityValue *DateAndAmount1 `xml:"EqtyVal,omitempty"`

	// Resource or value owned or used by a third-party company and the date on which the working capital amount was registered.
	WorkingCapital *DateAndAmount1 `xml:"WorkgCptl,omitempty"`

	// Account owner's connection with the trading party or broker.
	CompanyLink *CompanyLink1Choice `xml:"CpnyLk,omitempty"`

	// Reference to be specified when a letter (for example, an order confirmation) is sent by an automated mailing system.
	ElectronicMailingServiceReference *Max350Text `xml:"ElctrncMlngSvcRef,omitempty"`

	// Communication device number or electronic address used for communication.
	PrimaryCommunicationAddress []*CommunicationAddress6 `xml:"PmryComAdr,omitempty"`

	// Communication device number or electronic address used for communication.
	SecondaryCommunicationAddress []*CommunicationAddress6 `xml:"ScndryComAdr,omitempty"`

	// Additional regulatory information about the investor or account owner that is required in some markets to support anti-money laundering laws.
	AdditionalRegulatoryInformation *RegulatoryInformation1 `xml:"AddtlRgltryInf,omitempty"`

	// Specifies if the account party is regarded as domestic or non-domestic for reporting purposes.
	AccountingStatus *AccountingStatus1Choice `xml:"AcctgSts,omitempty"`

	// Additional information concerning limitations and restrictions on the account party.
	AdditionalInformation []*AccountRestrictions1 `xml:"AddtlInf,omitempty"`
}

func (i *InvestmentAccountOwnershipInformation12) AddParty() *Party31Choice {
	i.Party = new(Party31Choice)
	return i.Party
}

func (i *InvestmentAccountOwnershipInformation12) AddMoneyLaunderingCheck() *MoneyLaunderingCheck1Choice {
	i.MoneyLaunderingCheck = new(MoneyLaunderingCheck1Choice)
	return i.MoneyLaunderingCheck
}

func (i *InvestmentAccountOwnershipInformation12) AddInvestorProfileValidation() *PartyProfileInformation5 {
	newValue := new(PartyProfileInformation5)
	i.InvestorProfileValidation = append(i.InvestorProfileValidation, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation12) SetOwnershipBeneficiaryRate(value string) {
	i.OwnershipBeneficiaryRate = (*PercentageRate)(&value)
}

func (i *InvestmentAccountOwnershipInformation12) SetClientIdentification(value string) {
	i.ClientIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccountOwnershipInformation12) SetFiscalExemption(value string) {
	i.FiscalExemption = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountOwnershipInformation12) SetSignatoryRightIndicator(value string) {
	i.SignatoryRightIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountOwnershipInformation12) AddMiFIDClassification() *MiFIDClassification1 {
	i.MiFIDClassification = new(MiFIDClassification1)
	return i.MiFIDClassification
}

func (i *InvestmentAccountOwnershipInformation12) AddNotification() *Notification2 {
	newValue := new(Notification2)
	i.Notification = append(i.Notification, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation12) AddFATCAFormType() *FATCAForm1Choice {
	newValue := new(FATCAForm1Choice)
	i.FATCAFormType = append(i.FATCAFormType, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation12) AddFATCAStatus() *FATCAStatus2 {
	newValue := new(FATCAStatus2)
	i.FATCAStatus = append(i.FATCAStatus, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation12) AddOtherIdentification() *GenericIdentification82 {
	newValue := new(GenericIdentification82)
	i.OtherIdentification = append(i.OtherIdentification, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation12) AddTaxExemption() *TaxExemptionReason2Choice {
	i.TaxExemption = new(TaxExemptionReason2Choice)
	return i.TaxExemption
}

func (i *InvestmentAccountOwnershipInformation12) AddTaxReporting() *TaxReporting1 {
	newValue := new(TaxReporting1)
	i.TaxReporting = append(i.TaxReporting, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation12) SetLanguage(value string) {
	i.Language = (*LanguageCode)(&value)
}

func (i *InvestmentAccountOwnershipInformation12) AddMailType() *MailType1Choice {
	i.MailType = new(MailType1Choice)
	return i.MailType
}

func (i *InvestmentAccountOwnershipInformation12) AddCountryAndResidentialStatus() *CountryAndResidentialStatusType2 {
	i.CountryAndResidentialStatus = new(CountryAndResidentialStatusType2)
	return i.CountryAndResidentialStatus
}

func (i *InvestmentAccountOwnershipInformation12) AddMonetaryWealth() *DateAndAmount1 {
	i.MonetaryWealth = new(DateAndAmount1)
	return i.MonetaryWealth
}

func (i *InvestmentAccountOwnershipInformation12) AddEquityValue() *DateAndAmount1 {
	i.EquityValue = new(DateAndAmount1)
	return i.EquityValue
}

func (i *InvestmentAccountOwnershipInformation12) AddWorkingCapital() *DateAndAmount1 {
	i.WorkingCapital = new(DateAndAmount1)
	return i.WorkingCapital
}

func (i *InvestmentAccountOwnershipInformation12) AddCompanyLink() *CompanyLink1Choice {
	i.CompanyLink = new(CompanyLink1Choice)
	return i.CompanyLink
}

func (i *InvestmentAccountOwnershipInformation12) SetElectronicMailingServiceReference(value string) {
	i.ElectronicMailingServiceReference = (*Max350Text)(&value)
}

func (i *InvestmentAccountOwnershipInformation12) AddPrimaryCommunicationAddress() *CommunicationAddress6 {
	newValue := new(CommunicationAddress6)
	i.PrimaryCommunicationAddress = append(i.PrimaryCommunicationAddress, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation12) AddSecondaryCommunicationAddress() *CommunicationAddress6 {
	newValue := new(CommunicationAddress6)
	i.SecondaryCommunicationAddress = append(i.SecondaryCommunicationAddress, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation12) AddAdditionalRegulatoryInformation() *RegulatoryInformation1 {
	i.AdditionalRegulatoryInformation = new(RegulatoryInformation1)
	return i.AdditionalRegulatoryInformation
}

func (i *InvestmentAccountOwnershipInformation12) AddAccountingStatus() *AccountingStatus1Choice {
	i.AccountingStatus = new(AccountingStatus1Choice)
	return i.AccountingStatus
}

func (i *InvestmentAccountOwnershipInformation12) AddAdditionalInformation() *AccountRestrictions1 {
	newValue := new(AccountRestrictions1)
	i.AdditionalInformation = append(i.AdditionalInformation, newValue)
	return newValue
}

// Characteristics of the ownership of an investment account.
type InvestmentAccountOwnershipInformation13 struct {

	// Information about the organisation or individual person.
	Party *Party30Choice `xml:"Pty"`

	// Status of an identity check to prevent money laundering. This includes the counter-terrorism check.
	MoneyLaunderingCheck *MoneyLaunderingCheck1Choice `xml:"MnyLndrgChck,omitempty"`

	// Information related to the party profile to be inserted or deleted.
	ModifiedInvestorProfileValidation []*ModificationScope27 `xml:"ModfdInvstrPrflVldtn,omitempty"`

	// Percentage of ownership or of beneficial ownership of the shares/units in the account. All subsequent subscriptions or purchases and or redemptions or sells  will be allocated using the same percentage.
	OwnershipBeneficiaryRate *PercentageRate `xml:"OwnrshBnfcryRate,omitempty"`

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	ClientIdentification *Max35Text `xml:"ClntId,omitempty"`

	// Indicates whether an owner of the account may benefit from a fiscal exemption or amnesty, for example, when declaring overseas investments.
	FiscalExemption *YesNoIndicator `xml:"FsclXmptn,omitempty"`

	// Indicates whether the signature of the account owner is required to authorise transactions on the account.
	SignatoryRightIndicator *YesNoIndicator `xml:"SgntryRghtInd,omitempty"`

	// Details about the MiFID classification of the account owner.
	MiFIDClassification *MiFIDClassification1 `xml:"MiFIDClssfctn,omitempty"`

	// Type of information that must be provided to the account holder.
	Notification []*Notification2 `xml:"Ntfctn,omitempty"`

	// Type of Foreign Account Tax Compliance Act (FATCA) form submitted by the investor or account owner.
	FATCAFormType []*FATCAForm1Choice `xml:"FATCAFormTp,omitempty"`

	// Foreign Account Tax Compliance Act (FATCA) status of the investor or account owner.
	FATCAStatus []*FATCAStatus2 `xml:"FATCASts,omitempty"`

	// Alternative identification, for example, national registration identification number, passport number, tax identification number. This may be an account number used to further identify the beneficial owner, for example, a Central Provident Fund (CFP) account as required for Singapore.
	OtherIdentification []*GenericIdentification82 `xml:"OthrId,omitempty"`

	// Tax advantage specific to the account party.
	TaxExemption *TaxExemptionReason2Choice `xml:"TaxXmptn,omitempty"`

	// Details for the reporting of tax, for example, the country of taxation.
	TaxReporting []*TaxReporting1 `xml:"TaxRptg,omitempty"`

	// Language in which the organisation or person communicates.
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Method used for postal mailing.
	MailType *MailType1Choice `xml:"MailTp,omitempty"`

	// Country and residential status of the organisation or individual person.
	CountryAndResidentialStatus *CountryAndResidentialStatusType2 `xml:"CtryAndResdtlSts,omitempty"`

	// Annual wealth of the individual person or share capital value of the legal entity and date on which the annual wealth of the individual person was registered or declared or the date the  stock value of the organisation was registered.
	MonetaryWealth *DateAndAmount1 `xml:"MntryWlth,omitempty"`

	// Amount of total assets minus liabilities of the individual person or the amount of the difference between assets and liabilities plus rights over obligations (net equity) of the organisation and the date on which the equity value was registered.
	EquityValue *DateAndAmount1 `xml:"EqtyVal,omitempty"`

	// Resource or value owned or used by a third-party company and the date on which the working capital amount was registered.
	WorkingCapital *DateAndAmount1 `xml:"WorkgCptl,omitempty"`

	// Account owner's connection with the trading party or broker.
	CompanyLink *CompanyLink1Choice `xml:"CpnyLk,omitempty"`

	// Reference to be specified when a letter (for example, an order confirmation) is sent by an automated mailing system.
	ElectronicMailingServiceReference *Max350Text `xml:"ElctrncMlngSvcRef,omitempty"`

	// Communication device number or electronic address used for communication.
	PrimaryCommunicationAddress []*CommunicationAddress6 `xml:"PmryComAdr,omitempty"`

	// Communication device number or electronic address used for communication.
	SecondaryCommunicationAddress []*CommunicationAddress6 `xml:"ScndryComAdr,omitempty"`

	// Additional regulatory information about the investor or account owner that is required in some markets to support anti-money laundering laws.
	AdditionalRegulatoryInformation *RegulatoryInformation1 `xml:"AddtlRgltryInf,omitempty"`

	// Specifies if the account party is regarded as domestic or non-domestic for reporting purposes.
	AccountingStatus *AccountingStatus1Choice `xml:"AcctgSts,omitempty"`

	// Additional information concerning limitations and restrictions on the account party.
	AdditionalInformation []*AccountRestrictions1 `xml:"AddtlInf,omitempty"`
}

func (i *InvestmentAccountOwnershipInformation13) AddParty() *Party30Choice {
	i.Party = new(Party30Choice)
	return i.Party
}

func (i *InvestmentAccountOwnershipInformation13) AddMoneyLaunderingCheck() *MoneyLaunderingCheck1Choice {
	i.MoneyLaunderingCheck = new(MoneyLaunderingCheck1Choice)
	return i.MoneyLaunderingCheck
}

func (i *InvestmentAccountOwnershipInformation13) AddModifiedInvestorProfileValidation() *ModificationScope27 {
	newValue := new(ModificationScope27)
	i.ModifiedInvestorProfileValidation = append(i.ModifiedInvestorProfileValidation, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation13) SetOwnershipBeneficiaryRate(value string) {
	i.OwnershipBeneficiaryRate = (*PercentageRate)(&value)
}

func (i *InvestmentAccountOwnershipInformation13) SetClientIdentification(value string) {
	i.ClientIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccountOwnershipInformation13) SetFiscalExemption(value string) {
	i.FiscalExemption = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountOwnershipInformation13) SetSignatoryRightIndicator(value string) {
	i.SignatoryRightIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountOwnershipInformation13) AddMiFIDClassification() *MiFIDClassification1 {
	i.MiFIDClassification = new(MiFIDClassification1)
	return i.MiFIDClassification
}

func (i *InvestmentAccountOwnershipInformation13) AddNotification() *Notification2 {
	newValue := new(Notification2)
	i.Notification = append(i.Notification, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation13) AddFATCAFormType() *FATCAForm1Choice {
	newValue := new(FATCAForm1Choice)
	i.FATCAFormType = append(i.FATCAFormType, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation13) AddFATCAStatus() *FATCAStatus2 {
	newValue := new(FATCAStatus2)
	i.FATCAStatus = append(i.FATCAStatus, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation13) AddOtherIdentification() *GenericIdentification82 {
	newValue := new(GenericIdentification82)
	i.OtherIdentification = append(i.OtherIdentification, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation13) AddTaxExemption() *TaxExemptionReason2Choice {
	i.TaxExemption = new(TaxExemptionReason2Choice)
	return i.TaxExemption
}

func (i *InvestmentAccountOwnershipInformation13) AddTaxReporting() *TaxReporting1 {
	newValue := new(TaxReporting1)
	i.TaxReporting = append(i.TaxReporting, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation13) SetLanguage(value string) {
	i.Language = (*LanguageCode)(&value)
}

func (i *InvestmentAccountOwnershipInformation13) AddMailType() *MailType1Choice {
	i.MailType = new(MailType1Choice)
	return i.MailType
}

func (i *InvestmentAccountOwnershipInformation13) AddCountryAndResidentialStatus() *CountryAndResidentialStatusType2 {
	i.CountryAndResidentialStatus = new(CountryAndResidentialStatusType2)
	return i.CountryAndResidentialStatus
}

func (i *InvestmentAccountOwnershipInformation13) AddMonetaryWealth() *DateAndAmount1 {
	i.MonetaryWealth = new(DateAndAmount1)
	return i.MonetaryWealth
}

func (i *InvestmentAccountOwnershipInformation13) AddEquityValue() *DateAndAmount1 {
	i.EquityValue = new(DateAndAmount1)
	return i.EquityValue
}

func (i *InvestmentAccountOwnershipInformation13) AddWorkingCapital() *DateAndAmount1 {
	i.WorkingCapital = new(DateAndAmount1)
	return i.WorkingCapital
}

func (i *InvestmentAccountOwnershipInformation13) AddCompanyLink() *CompanyLink1Choice {
	i.CompanyLink = new(CompanyLink1Choice)
	return i.CompanyLink
}

func (i *InvestmentAccountOwnershipInformation13) SetElectronicMailingServiceReference(value string) {
	i.ElectronicMailingServiceReference = (*Max350Text)(&value)
}

func (i *InvestmentAccountOwnershipInformation13) AddPrimaryCommunicationAddress() *CommunicationAddress6 {
	newValue := new(CommunicationAddress6)
	i.PrimaryCommunicationAddress = append(i.PrimaryCommunicationAddress, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation13) AddSecondaryCommunicationAddress() *CommunicationAddress6 {
	newValue := new(CommunicationAddress6)
	i.SecondaryCommunicationAddress = append(i.SecondaryCommunicationAddress, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation13) AddAdditionalRegulatoryInformation() *RegulatoryInformation1 {
	i.AdditionalRegulatoryInformation = new(RegulatoryInformation1)
	return i.AdditionalRegulatoryInformation
}

func (i *InvestmentAccountOwnershipInformation13) AddAccountingStatus() *AccountingStatus1Choice {
	i.AccountingStatus = new(AccountingStatus1Choice)
	return i.AccountingStatus
}

func (i *InvestmentAccountOwnershipInformation13) AddAdditionalInformation() *AccountRestrictions1 {
	newValue := new(AccountRestrictions1)
	i.AdditionalInformation = append(i.AdditionalInformation, newValue)
	return newValue
}

// Characteristics of the ownership of a securities account.
type InvestmentAccountOwnershipInformation14 struct {

	// Information about the organisation or individual person.
	Party *Party32Choice `xml:"Pty"`

	// Status of an identity check to prevent money laundering. This includes the counter-terrorism check.
	MoneyLaunderingCheck *MoneyLaunderingCheck1Choice `xml:"MnyLndrgChck,omitempty"`

	// Information to support Know Your Customer (KYC) processes.
	InvestorProfileValidation []*PartyProfileInformation5 `xml:"InvstrPrflVldtn,omitempty"`

	// Percentage of ownership or of beneficial ownership of the shares/units in the account. All subsequent subscriptions or purchases and or redemptions or sells  will be allocated using the same percentage.
	OwnershipBeneficiaryRate *OwnershipBeneficiaryRate1 `xml:"OwnrshBnfcryRate,omitempty"`

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	ClientIdentification *Max35Text `xml:"ClntId,omitempty"`

	// Indicates whether an owner of the account may benefit from a fiscal exemption or amnesty, for example, when declaring overseas investments.
	FiscalExemption *YesNoIndicator `xml:"FsclXmptn,omitempty"`

	// Indicates whether the signature of the account owner is required to authorise transactions on the account.
	SignatoryRightIndicator *YesNoIndicator `xml:"SgntryRghtInd,omitempty"`

	// Details about the MiFID classification of the account owner.
	MiFIDClassification *MiFIDClassification1 `xml:"MiFIDClssfctn,omitempty"`

	// Type of information that must be provided to the account holder.
	Notification []*Notification2 `xml:"Ntfctn,omitempty"`

	// Type of Foreign Account Tax Compliance Act (FATCA) form submitted by the investor or account owner.
	FATCAFormType []*FATCAForm1Choice `xml:"FATCAFormTp,omitempty"`

	// Foreign Account Tax Compliance Act (FATCA) status of the investor or account owner.
	FATCAStatus []*FATCAStatus2 `xml:"FATCASts,omitempty"`

	// Date provided by the account owner to inform the account servicer of the date on which the holdings must be reported before the account is subsequently closed.
	FATCAReportingDate *ISODate `xml:"FATCARptgDt,omitempty"`

	// Type of Common Reporting Standard (CRS) form submitted by the investor or account owner.
	CRSFormType []*CRSForm1Choice `xml:"CRSFormTp,omitempty"`

	// Common Reporting Standard (CRS) status of the investor or account owner.
	CRSStatus []*CRSStatus4 `xml:"CRSSts,omitempty"`

	// Date provided by the account owner to inform the account servicer of the date on which the holdings must be reported before the account is subsequently closed.
	CRSReportingDate *ISODate `xml:"CRSRptgDt,omitempty"`

	// Alternative identification, for example, national registration identification number, passport number, tax identification number. This may be an account number used to further identify the beneficial owner, for example, a Central Provident Fund (CFP) account as required for Singapore.
	OtherIdentification []*GenericIdentification82 `xml:"OthrId,omitempty"`

	// Tax advantage specific to the account party.
	TaxExemption *TaxExemptionReason2Choice `xml:"TaxXmptn,omitempty"`

	// Details for the reporting of tax, for example, the country of taxation.
	TaxReporting []*TaxReporting2 `xml:"TaxRptg,omitempty"`

	// Language in which the organisation or person communicates.
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Method used for postal mailing.
	MailType *MailType1Choice `xml:"MailTp,omitempty"`

	// Country and residential status of the organisation or individual person.
	CountryAndResidentialStatus *CountryAndResidentialStatusType2 `xml:"CtryAndResdtlSts,omitempty"`

	// Annual wealth of the individual person or share capital value of the legal entity and date on which the annual wealth of the individual person was registered or declared or the date the stock value of the organisation was registered.
	MonetaryWealth *DateAndAmount1 `xml:"MntryWlth,omitempty"`

	// Amount of total assets minus liabilities of the individual person or the amount of the difference between assets and liabilities plus rights over obligations (net equity) of the organisation and the date on which the equity value was registered.
	EquityValue *DateAndAmount1 `xml:"EqtyVal,omitempty"`

	// Resource or value owned or used by a third-party company and the date on which the working capital amount was registered.
	WorkingCapital *DateAndAmount1 `xml:"WorkgCptl,omitempty"`

	// Account owner's connection with the trading party or broker.
	CompanyLink *CompanyLink1Choice `xml:"CpnyLk,omitempty"`

	// Reference to be specified when a letter (for example, an order confirmation) is sent by an automated mailing system.
	ElectronicMailingServiceReference *Max350Text `xml:"ElctrncMlngSvcRef,omitempty"`

	// Communication device number or electronic address used for communication.
	PrimaryCommunicationAddress []*CommunicationAddress6 `xml:"PmryComAdr,omitempty"`

	// Communication device number or electronic address used for communication.
	SecondaryCommunicationAddress []*CommunicationAddress6 `xml:"ScndryComAdr,omitempty"`

	// Additional regulatory information about the investor or account owner that is required in some markets to support anti-money laundering laws.
	AdditionalRegulatoryInformation *RegulatoryInformation1 `xml:"AddtlRgltryInf,omitempty"`

	// Specifies if the account party is regarded as domestic or non-domestic for reporting purposes.
	AccountingStatus *AccountingStatus1Choice `xml:"AcctgSts,omitempty"`

	// Additional information such as remarks or notes that must be conveyed about the party and or  limitations and restrictions.
	AdditionalInformation []*AdditiononalInformation12 `xml:"AddtlInf,omitempty"`

	// Party is the controlling person.
	// (For an Entity that is a legal person, the term “Controlling Persons” means the natural person(s) who exercises control over the Entity. “Control” over an Entity is generally exercised by the natural person(s) who ultimately has a controlling ownership interest in the Entity. A “control ownership interest” depends on the ownership structure of the legal person and is usually identified on the basis of a threshold applying a risk-based approach (e.g. any person(s) owning more than a certain percentage of the legal person, such as 25%). Where no natural person(s) exercises control through ownership interests, the Controlling Person(s) of the Entity will be the natural person(s) who exercises control of the Entity through other means. Where no natural person(s) is identified as exercising control of the Entity, the Controlling Person(s) of the Entity will be the natural person(s) who holds the position of senior managing official.)
	ControllingParty *YesNoIndicator `xml:"CtrlgPty,omitempty"`
}

func (i *InvestmentAccountOwnershipInformation14) AddParty() *Party32Choice {
	i.Party = new(Party32Choice)
	return i.Party
}

func (i *InvestmentAccountOwnershipInformation14) AddMoneyLaunderingCheck() *MoneyLaunderingCheck1Choice {
	i.MoneyLaunderingCheck = new(MoneyLaunderingCheck1Choice)
	return i.MoneyLaunderingCheck
}

func (i *InvestmentAccountOwnershipInformation14) AddInvestorProfileValidation() *PartyProfileInformation5 {
	newValue := new(PartyProfileInformation5)
	i.InvestorProfileValidation = append(i.InvestorProfileValidation, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation14) AddOwnershipBeneficiaryRate() *OwnershipBeneficiaryRate1 {
	i.OwnershipBeneficiaryRate = new(OwnershipBeneficiaryRate1)
	return i.OwnershipBeneficiaryRate
}

func (i *InvestmentAccountOwnershipInformation14) SetClientIdentification(value string) {
	i.ClientIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccountOwnershipInformation14) SetFiscalExemption(value string) {
	i.FiscalExemption = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountOwnershipInformation14) SetSignatoryRightIndicator(value string) {
	i.SignatoryRightIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountOwnershipInformation14) AddMiFIDClassification() *MiFIDClassification1 {
	i.MiFIDClassification = new(MiFIDClassification1)
	return i.MiFIDClassification
}

func (i *InvestmentAccountOwnershipInformation14) AddNotification() *Notification2 {
	newValue := new(Notification2)
	i.Notification = append(i.Notification, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation14) AddFATCAFormType() *FATCAForm1Choice {
	newValue := new(FATCAForm1Choice)
	i.FATCAFormType = append(i.FATCAFormType, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation14) AddFATCAStatus() *FATCAStatus2 {
	newValue := new(FATCAStatus2)
	i.FATCAStatus = append(i.FATCAStatus, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation14) SetFATCAReportingDate(value string) {
	i.FATCAReportingDate = (*ISODate)(&value)
}

func (i *InvestmentAccountOwnershipInformation14) AddCRSFormType() *CRSForm1Choice {
	newValue := new(CRSForm1Choice)
	i.CRSFormType = append(i.CRSFormType, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation14) AddCRSStatus() *CRSStatus4 {
	newValue := new(CRSStatus4)
	i.CRSStatus = append(i.CRSStatus, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation14) SetCRSReportingDate(value string) {
	i.CRSReportingDate = (*ISODate)(&value)
}

func (i *InvestmentAccountOwnershipInformation14) AddOtherIdentification() *GenericIdentification82 {
	newValue := new(GenericIdentification82)
	i.OtherIdentification = append(i.OtherIdentification, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation14) AddTaxExemption() *TaxExemptionReason2Choice {
	i.TaxExemption = new(TaxExemptionReason2Choice)
	return i.TaxExemption
}

func (i *InvestmentAccountOwnershipInformation14) AddTaxReporting() *TaxReporting2 {
	newValue := new(TaxReporting2)
	i.TaxReporting = append(i.TaxReporting, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation14) SetLanguage(value string) {
	i.Language = (*LanguageCode)(&value)
}

func (i *InvestmentAccountOwnershipInformation14) AddMailType() *MailType1Choice {
	i.MailType = new(MailType1Choice)
	return i.MailType
}

func (i *InvestmentAccountOwnershipInformation14) AddCountryAndResidentialStatus() *CountryAndResidentialStatusType2 {
	i.CountryAndResidentialStatus = new(CountryAndResidentialStatusType2)
	return i.CountryAndResidentialStatus
}

func (i *InvestmentAccountOwnershipInformation14) AddMonetaryWealth() *DateAndAmount1 {
	i.MonetaryWealth = new(DateAndAmount1)
	return i.MonetaryWealth
}

func (i *InvestmentAccountOwnershipInformation14) AddEquityValue() *DateAndAmount1 {
	i.EquityValue = new(DateAndAmount1)
	return i.EquityValue
}

func (i *InvestmentAccountOwnershipInformation14) AddWorkingCapital() *DateAndAmount1 {
	i.WorkingCapital = new(DateAndAmount1)
	return i.WorkingCapital
}

func (i *InvestmentAccountOwnershipInformation14) AddCompanyLink() *CompanyLink1Choice {
	i.CompanyLink = new(CompanyLink1Choice)
	return i.CompanyLink
}

func (i *InvestmentAccountOwnershipInformation14) SetElectronicMailingServiceReference(value string) {
	i.ElectronicMailingServiceReference = (*Max350Text)(&value)
}

func (i *InvestmentAccountOwnershipInformation14) AddPrimaryCommunicationAddress() *CommunicationAddress6 {
	newValue := new(CommunicationAddress6)
	i.PrimaryCommunicationAddress = append(i.PrimaryCommunicationAddress, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation14) AddSecondaryCommunicationAddress() *CommunicationAddress6 {
	newValue := new(CommunicationAddress6)
	i.SecondaryCommunicationAddress = append(i.SecondaryCommunicationAddress, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation14) AddAdditionalRegulatoryInformation() *RegulatoryInformation1 {
	i.AdditionalRegulatoryInformation = new(RegulatoryInformation1)
	return i.AdditionalRegulatoryInformation
}

func (i *InvestmentAccountOwnershipInformation14) AddAccountingStatus() *AccountingStatus1Choice {
	i.AccountingStatus = new(AccountingStatus1Choice)
	return i.AccountingStatus
}

func (i *InvestmentAccountOwnershipInformation14) AddAdditionalInformation() *AdditiononalInformation12 {
	newValue := new(AdditiononalInformation12)
	i.AdditionalInformation = append(i.AdditionalInformation, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation14) SetControllingParty(value string) {
	i.ControllingParty = (*YesNoIndicator)(&value)
}

// Characteristics of the ownership of an investment account.
type InvestmentAccountOwnershipInformation15 struct {

	// Information about the organisation or individual person.
	Party *Party33Choice `xml:"Pty"`

	// Status of an identity check to prevent money laundering. This includes the counter-terrorism check.
	MoneyLaunderingCheck *MoneyLaunderingCheck1Choice `xml:"MnyLndrgChck,omitempty"`

	// Information related to the party profile to be inserted or deleted.
	ModifiedInvestorProfileValidation []*ModificationScope27 `xml:"ModfdInvstrPrflVldtn,omitempty"`

	// Percentage of ownership or of beneficial ownership of the shares/units in the account. All subsequent subscriptions or purchases and or redemptions or sells  will be allocated using the same percentage.
	OwnershipBeneficiaryRate *OwnershipBeneficiaryRate1 `xml:"OwnrshBnfcryRate,omitempty"`

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	ClientIdentification *Max35Text `xml:"ClntId,omitempty"`

	// Indicates whether an owner of the account may benefit from a fiscal exemption or amnesty, for example, when declaring overseas investments.
	FiscalExemption *YesNoIndicator `xml:"FsclXmptn,omitempty"`

	// Indicates whether the signature of the account owner is required to authorise transactions on the account.
	SignatoryRightIndicator *YesNoIndicator `xml:"SgntryRghtInd,omitempty"`

	// Details about the MiFID classification of the account owner.
	MiFIDClassification *MiFIDClassification1 `xml:"MiFIDClssfctn,omitempty"`

	// Type of information that must be provided to the account holder.
	Notification []*Notification2 `xml:"Ntfctn,omitempty"`

	// Type of Foreign Account Tax Compliance Act (FATCA) form submitted by the investor or account owner.
	FATCAFormType []*FATCAForm1Choice `xml:"FATCAFormTp,omitempty"`

	// Foreign Account Tax Compliance Act (FATCA) status of the investor or account owner.
	FATCAStatus []*FATCAStatus2 `xml:"FATCASts,omitempty"`

	// Date provided by the account owner to inform the account servicer of the date on which the holdings must be reported before the account is subsequently closed.
	FATCAReportingDate *ISODate `xml:"FATCARptgDt,omitempty"`

	// Type of Common Reporting Standard (CRS) form submitted by the investor or account owner.
	CRSFormType []*CRSForm1Choice `xml:"CRSFormTp,omitempty"`

	// Common Reporting Standard (CRS) status of the investor or account owner.
	CRSStatus []*CRSStatus4 `xml:"CRSSts,omitempty"`

	// Date provided by the account owner to inform the account servicer of the date on which the holdings must be reported before the account is subsequently closed.
	CRSReportingDate *ISODate `xml:"CRSRptgDt,omitempty"`

	// Alternative identification, for example, national registration identification number, passport number, tax identification number. This may be an account number used to further identify the beneficial owner, for example, a Central Provident Fund (CFP) account as required for Singapore.
	OtherIdentification []*GenericIdentification82 `xml:"OthrId,omitempty"`

	// Tax advantage specific to the account party.
	TaxExemption *TaxExemptionReason2Choice `xml:"TaxXmptn,omitempty"`

	// Details for the reporting of tax, for example, the country of taxation.
	TaxReporting []*TaxReporting2 `xml:"TaxRptg,omitempty"`

	// Language in which the organisation or person communicates.
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Method used for postal mailing.
	MailType *MailType1Choice `xml:"MailTp,omitempty"`

	// Country and residential status of the organisation or individual person.
	CountryAndResidentialStatus *CountryAndResidentialStatusType2 `xml:"CtryAndResdtlSts,omitempty"`

	// Annual wealth of the individual person or share capital value of the legal entity and date on which the annual wealth of the individual person was registered or declared or the date the  stock value of the organisation was registered.
	MonetaryWealth *DateAndAmount1 `xml:"MntryWlth,omitempty"`

	// Amount of total assets minus liabilities of the individual person or the amount of the difference between assets and liabilities plus rights over obligations (net equity) of the organisation and the date on which the equity value was registered.
	EquityValue *DateAndAmount1 `xml:"EqtyVal,omitempty"`

	// Resource or value owned or used by a third-party company and the date on which the working capital amount was registered.
	WorkingCapital *DateAndAmount1 `xml:"WorkgCptl,omitempty"`

	// Account owner's connection with the trading party or broker.
	CompanyLink *CompanyLink1Choice `xml:"CpnyLk,omitempty"`

	// Reference to be specified when a letter (for example, an order confirmation) is sent by an automated mailing system.
	ElectronicMailingServiceReference *Max350Text `xml:"ElctrncMlngSvcRef,omitempty"`

	// Communication device number or electronic address used for communication.
	PrimaryCommunicationAddress []*CommunicationAddress6 `xml:"PmryComAdr,omitempty"`

	// Communication device number or electronic address used for communication.
	SecondaryCommunicationAddress []*CommunicationAddress6 `xml:"ScndryComAdr,omitempty"`

	// Additional regulatory information about the investor or account owner that is required in some markets to support anti-money laundering laws.
	AdditionalRegulatoryInformation *RegulatoryInformation1 `xml:"AddtlRgltryInf,omitempty"`

	// Specifies if the account party is regarded as domestic or non-domestic for reporting purposes.
	AccountingStatus *AccountingStatus1Choice `xml:"AcctgSts,omitempty"`

	// Additional information such as remarks or notes that must be conveyed about the party and or  limitations and restrictions.
	AdditionalInformation []*AdditiononalInformation12 `xml:"AddtlInf,omitempty"`

	// Party is the controlling person.
	// (For an Entity that is a legal person, the term “Controlling Persons” means the natural person(s) who exercises control over the Entity. “Control” over an Entity is generally exercised by the natural person(s) who ultimately has a controlling ownership interest in the Entity. A “control ownership interest” depends on the ownership structure of the legal person and is usually identified on the basis of a threshold applying a risk-based approach (e.g. any person(s) owning more than a certain percentage of the legal person, such as 25%). Where no natural person(s) exercises control through ownership interests, the Controlling Person(s) of the Entity will be the natural person(s) who exercises control of the Entity through other means. Where no natural person(s) is identified as exercising control of the Entity, the Controlling Person(s) of the Entity will be the natural person(s) who holds the position of senior managing official.)
	ControllingParty *YesNoIndicator `xml:"CtrlgPty,omitempty"`
}

func (i *InvestmentAccountOwnershipInformation15) AddParty() *Party33Choice {
	i.Party = new(Party33Choice)
	return i.Party
}

func (i *InvestmentAccountOwnershipInformation15) AddMoneyLaunderingCheck() *MoneyLaunderingCheck1Choice {
	i.MoneyLaunderingCheck = new(MoneyLaunderingCheck1Choice)
	return i.MoneyLaunderingCheck
}

func (i *InvestmentAccountOwnershipInformation15) AddModifiedInvestorProfileValidation() *ModificationScope27 {
	newValue := new(ModificationScope27)
	i.ModifiedInvestorProfileValidation = append(i.ModifiedInvestorProfileValidation, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation15) AddOwnershipBeneficiaryRate() *OwnershipBeneficiaryRate1 {
	i.OwnershipBeneficiaryRate = new(OwnershipBeneficiaryRate1)
	return i.OwnershipBeneficiaryRate
}

func (i *InvestmentAccountOwnershipInformation15) SetClientIdentification(value string) {
	i.ClientIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccountOwnershipInformation15) SetFiscalExemption(value string) {
	i.FiscalExemption = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountOwnershipInformation15) SetSignatoryRightIndicator(value string) {
	i.SignatoryRightIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountOwnershipInformation15) AddMiFIDClassification() *MiFIDClassification1 {
	i.MiFIDClassification = new(MiFIDClassification1)
	return i.MiFIDClassification
}

func (i *InvestmentAccountOwnershipInformation15) AddNotification() *Notification2 {
	newValue := new(Notification2)
	i.Notification = append(i.Notification, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation15) AddFATCAFormType() *FATCAForm1Choice {
	newValue := new(FATCAForm1Choice)
	i.FATCAFormType = append(i.FATCAFormType, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation15) AddFATCAStatus() *FATCAStatus2 {
	newValue := new(FATCAStatus2)
	i.FATCAStatus = append(i.FATCAStatus, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation15) SetFATCAReportingDate(value string) {
	i.FATCAReportingDate = (*ISODate)(&value)
}

func (i *InvestmentAccountOwnershipInformation15) AddCRSFormType() *CRSForm1Choice {
	newValue := new(CRSForm1Choice)
	i.CRSFormType = append(i.CRSFormType, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation15) AddCRSStatus() *CRSStatus4 {
	newValue := new(CRSStatus4)
	i.CRSStatus = append(i.CRSStatus, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation15) SetCRSReportingDate(value string) {
	i.CRSReportingDate = (*ISODate)(&value)
}

func (i *InvestmentAccountOwnershipInformation15) AddOtherIdentification() *GenericIdentification82 {
	newValue := new(GenericIdentification82)
	i.OtherIdentification = append(i.OtherIdentification, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation15) AddTaxExemption() *TaxExemptionReason2Choice {
	i.TaxExemption = new(TaxExemptionReason2Choice)
	return i.TaxExemption
}

func (i *InvestmentAccountOwnershipInformation15) AddTaxReporting() *TaxReporting2 {
	newValue := new(TaxReporting2)
	i.TaxReporting = append(i.TaxReporting, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation15) SetLanguage(value string) {
	i.Language = (*LanguageCode)(&value)
}

func (i *InvestmentAccountOwnershipInformation15) AddMailType() *MailType1Choice {
	i.MailType = new(MailType1Choice)
	return i.MailType
}

func (i *InvestmentAccountOwnershipInformation15) AddCountryAndResidentialStatus() *CountryAndResidentialStatusType2 {
	i.CountryAndResidentialStatus = new(CountryAndResidentialStatusType2)
	return i.CountryAndResidentialStatus
}

func (i *InvestmentAccountOwnershipInformation15) AddMonetaryWealth() *DateAndAmount1 {
	i.MonetaryWealth = new(DateAndAmount1)
	return i.MonetaryWealth
}

func (i *InvestmentAccountOwnershipInformation15) AddEquityValue() *DateAndAmount1 {
	i.EquityValue = new(DateAndAmount1)
	return i.EquityValue
}

func (i *InvestmentAccountOwnershipInformation15) AddWorkingCapital() *DateAndAmount1 {
	i.WorkingCapital = new(DateAndAmount1)
	return i.WorkingCapital
}

func (i *InvestmentAccountOwnershipInformation15) AddCompanyLink() *CompanyLink1Choice {
	i.CompanyLink = new(CompanyLink1Choice)
	return i.CompanyLink
}

func (i *InvestmentAccountOwnershipInformation15) SetElectronicMailingServiceReference(value string) {
	i.ElectronicMailingServiceReference = (*Max350Text)(&value)
}

func (i *InvestmentAccountOwnershipInformation15) AddPrimaryCommunicationAddress() *CommunicationAddress6 {
	newValue := new(CommunicationAddress6)
	i.PrimaryCommunicationAddress = append(i.PrimaryCommunicationAddress, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation15) AddSecondaryCommunicationAddress() *CommunicationAddress6 {
	newValue := new(CommunicationAddress6)
	i.SecondaryCommunicationAddress = append(i.SecondaryCommunicationAddress, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation15) AddAdditionalRegulatoryInformation() *RegulatoryInformation1 {
	i.AdditionalRegulatoryInformation = new(RegulatoryInformation1)
	return i.AdditionalRegulatoryInformation
}

func (i *InvestmentAccountOwnershipInformation15) AddAccountingStatus() *AccountingStatus1Choice {
	i.AccountingStatus = new(AccountingStatus1Choice)
	return i.AccountingStatus
}

func (i *InvestmentAccountOwnershipInformation15) AddAdditionalInformation() *AdditiononalInformation12 {
	newValue := new(AdditiononalInformation12)
	i.AdditionalInformation = append(i.AdditionalInformation, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation15) SetControllingParty(value string) {
	i.ControllingParty = (*YesNoIndicator)(&value)
}

// Characteristics of the ownership of an investment account.
type InvestmentAccountOwnershipInformation4 struct {

	// Organised structure that is set up for a particular purpose, eg, a business, government body, department, charity, or financial institution.
	Organisation *Organisation3 `xml:"Org"`

	// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
	IndividualPerson *IndividualPerson11 `xml:"IndvPrsn"`

	// Status of an identity check to prevent money laundering. This includes the counter-terrorism check.
	MoneyLaunderingCheck *MoneyLaunderingCheck1Code `xml:"MnyLndrgChck,omitempty"`

	// Status of an identity check to prevent money laundering. This includes the counter-terrorism check.
	ExtendedMoneyLaunderingCheck *Extended350Code `xml:"XtndedMnyLndrgChck,omitempty"`

	// Percentage of ownership or beneficiary ownership of the shares/units in the account. All subsequent subscriptions and or redemptions will be allocated using the same percentage.
	OwnershipBeneficiaryRate *PercentageRate `xml:"OwnrshBnfcryRate,omitempty"`

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	ClientIdentification *Max35Text `xml:"ClntId,omitempty"`

	// Indicates whether an owner of an investment account may benefit from a fiscal exemption or amnesty for instance for declaring overseas investments.
	FiscalExemption *YesNoIndicator `xml:"FsclXmptn,omitempty"`

	// Indicates whether the account owner signature is required to authorise transactions on the account.
	SignatoryRightIndicator *YesNoIndicator `xml:"SgntryRghtInd,omitempty"`

	// Information related to the party profile to be inserted or deleted.
	ModifiedInvestorProfileValidation []*ModificationScope11 `xml:"ModfdInvstrPrflVldtn,omitempty"`
}

func (i *InvestmentAccountOwnershipInformation4) AddOrganisation() *Organisation3 {
	i.Organisation = new(Organisation3)
	return i.Organisation
}

func (i *InvestmentAccountOwnershipInformation4) AddIndividualPerson() *IndividualPerson11 {
	i.IndividualPerson = new(IndividualPerson11)
	return i.IndividualPerson
}

func (i *InvestmentAccountOwnershipInformation4) SetMoneyLaunderingCheck(value string) {
	i.MoneyLaunderingCheck = (*MoneyLaunderingCheck1Code)(&value)
}

func (i *InvestmentAccountOwnershipInformation4) SetExtendedMoneyLaunderingCheck(value string) {
	i.ExtendedMoneyLaunderingCheck = (*Extended350Code)(&value)
}

func (i *InvestmentAccountOwnershipInformation4) SetOwnershipBeneficiaryRate(value string) {
	i.OwnershipBeneficiaryRate = (*PercentageRate)(&value)
}

func (i *InvestmentAccountOwnershipInformation4) SetClientIdentification(value string) {
	i.ClientIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccountOwnershipInformation4) SetFiscalExemption(value string) {
	i.FiscalExemption = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountOwnershipInformation4) SetSignatoryRightIndicator(value string) {
	i.SignatoryRightIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountOwnershipInformation4) AddModifiedInvestorProfileValidation() *ModificationScope11 {
	newValue := new(ModificationScope11)
	i.ModifiedInvestorProfileValidation = append(i.ModifiedInvestorProfileValidation, newValue)
	return newValue
}

// Characteristics of the ownership of an investment account.
type InvestmentAccountOwnershipInformation5 struct {

	// Organised structure that is set up for a particular purpose, eg, a business, government body, department, charity, or financial institution.
	Organisation *Organisation2 `xml:"Org"`

	// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
	IndividualPerson *IndividualPerson10 `xml:"IndvPrsn"`

	// Status of an identity check to prevent money laundering. This includes the counter-terrorism check.
	MoneyLaunderingCheck *MoneyLaunderingCheck1Code `xml:"MnyLndrgChck,omitempty"`

	// Status of an identity check to prevent money laundering. This includes the counter-terrorism check.
	ExtendedMoneyLaunderingCheck *Extended350Code `xml:"XtndedMnyLndrgChck,omitempty"`

	// Information to support Know Your Customer processes.
	InvestorProfileValidation []*PartyProfileInformation1 `xml:"InvstrPrflVldtn,omitempty"`

	// Percentage of ownership or of beneficial ownership of the shares/units in the account. All subsequent subscriptions and or redemptions will be allocated using the same percentage.
	OwnershipBeneficiaryRate *PercentageRate `xml:"OwnrshBnfcryRate,omitempty"`

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	ClientIdentification *Max35Text `xml:"ClntId,omitempty"`

	// Indicates whether an owner of an investment account may benefit from a fiscal exemption or amnesty for instance for declaring overseas investments.
	FiscalExemption *YesNoIndicator `xml:"FsclXmptn,omitempty"`

	// Indicates whether the signature of the account owner is required to authorise transactions on the account.
	SignatoryRightIndicator *YesNoIndicator `xml:"SgntryRghtInd,omitempty"`
}

func (i *InvestmentAccountOwnershipInformation5) AddOrganisation() *Organisation2 {
	i.Organisation = new(Organisation2)
	return i.Organisation
}

func (i *InvestmentAccountOwnershipInformation5) AddIndividualPerson() *IndividualPerson10 {
	i.IndividualPerson = new(IndividualPerson10)
	return i.IndividualPerson
}

func (i *InvestmentAccountOwnershipInformation5) SetMoneyLaunderingCheck(value string) {
	i.MoneyLaunderingCheck = (*MoneyLaunderingCheck1Code)(&value)
}

func (i *InvestmentAccountOwnershipInformation5) SetExtendedMoneyLaunderingCheck(value string) {
	i.ExtendedMoneyLaunderingCheck = (*Extended350Code)(&value)
}

func (i *InvestmentAccountOwnershipInformation5) AddInvestorProfileValidation() *PartyProfileInformation1 {
	newValue := new(PartyProfileInformation1)
	i.InvestorProfileValidation = append(i.InvestorProfileValidation, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation5) SetOwnershipBeneficiaryRate(value string) {
	i.OwnershipBeneficiaryRate = (*PercentageRate)(&value)
}

func (i *InvestmentAccountOwnershipInformation5) SetClientIdentification(value string) {
	i.ClientIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccountOwnershipInformation5) SetFiscalExemption(value string) {
	i.FiscalExemption = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountOwnershipInformation5) SetSignatoryRightIndicator(value string) {
	i.SignatoryRightIndicator = (*YesNoIndicator)(&value)
}

// Characteristics of the ownership of an investment account.
type InvestmentAccountOwnershipInformation6 struct {

	// Information about the organisation or individual person.
	Party *Party14Choice `xml:"Pty"`

	// Status of an identity check to prevent money laundering. This includes the counter-terrorism check.
	MoneyLaunderingCheck *MoneyLaunderingCheck1Choice `xml:"MnyLndrgChck,omitempty"`

	// Information to support Know Your Customer processes.
	InvestorProfileValidation []*PartyProfileInformation2 `xml:"InvstrPrflVldtn,omitempty"`

	// Percentage of ownership or of beneficial ownership of the shares/units in the account. All subsequent subscriptions and or redemptions will be allocated using the same percentage.
	OwnershipBeneficiaryRate *PercentageRate `xml:"OwnrshBnfcryRate,omitempty"`

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	ClientIdentification *Max35Text `xml:"ClntId,omitempty"`

	// Indicates whether an owner of an investment account may benefit from a fiscal exemption or amnesty for instance for declaring overseas investments.
	FiscalExemption *YesNoIndicator `xml:"FsclXmptn,omitempty"`

	// Indicates whether the signature of the account owner is required to authorise transactions on the account.
	SignatoryRightIndicator *YesNoIndicator `xml:"SgntryRghtInd,omitempty"`

	// Details about the MiFID classification of the account owner.
	MiFIDClassification *MiFIDClassification1 `xml:"MiFIDClssfctn,omitempty"`

	// Specifies how information is sent to the account holder.
	InformationDistribution *InformationDistribution1Code `xml:"InfDstrbtn,omitempty"`
}

func (i *InvestmentAccountOwnershipInformation6) AddParty() *Party14Choice {
	i.Party = new(Party14Choice)
	return i.Party
}

func (i *InvestmentAccountOwnershipInformation6) AddMoneyLaunderingCheck() *MoneyLaunderingCheck1Choice {
	i.MoneyLaunderingCheck = new(MoneyLaunderingCheck1Choice)
	return i.MoneyLaunderingCheck
}

func (i *InvestmentAccountOwnershipInformation6) AddInvestorProfileValidation() *PartyProfileInformation2 {
	newValue := new(PartyProfileInformation2)
	i.InvestorProfileValidation = append(i.InvestorProfileValidation, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation6) SetOwnershipBeneficiaryRate(value string) {
	i.OwnershipBeneficiaryRate = (*PercentageRate)(&value)
}

func (i *InvestmentAccountOwnershipInformation6) SetClientIdentification(value string) {
	i.ClientIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccountOwnershipInformation6) SetFiscalExemption(value string) {
	i.FiscalExemption = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountOwnershipInformation6) SetSignatoryRightIndicator(value string) {
	i.SignatoryRightIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountOwnershipInformation6) AddMiFIDClassification() *MiFIDClassification1 {
	i.MiFIDClassification = new(MiFIDClassification1)
	return i.MiFIDClassification
}

func (i *InvestmentAccountOwnershipInformation6) SetInformationDistribution(value string) {
	i.InformationDistribution = (*InformationDistribution1Code)(&value)
}

// Characteristics of the ownership of an investment account.
type InvestmentAccountOwnershipInformation7 struct {

	// Information about the organisation or individual person.
	Party *Party14Choice `xml:"Pty"`

	// Status of an identity check to prevent money laundering. This includes the counter-terrorism check.
	MoneyLaunderingCheck *MoneyLaunderingCheck1Choice `xml:"MnyLndrgChck,omitempty"`

	// Percentage of ownership or beneficiary ownership of the shares/units in the account. All subsequent subscriptions and or redemptions will be allocated using the same percentage.
	OwnershipBeneficiaryRate *PercentageRate `xml:"OwnrshBnfcryRate,omitempty"`

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	ClientIdentification *Max35Text `xml:"ClntId,omitempty"`

	// Indicates whether an owner of an investment account may benefit from a fiscal exemption or amnesty for instance for declaring overseas investments.
	FiscalExemption *YesNoIndicator `xml:"FsclXmptn,omitempty"`

	// Indicates whether the account owner signature is required to authorise transactions on the account.
	SignatoryRightIndicator *YesNoIndicator `xml:"SgntryRghtInd,omitempty"`

	// Information related to the party profile to be inserted or deleted.
	ModifiedInvestorProfileValidation []*ModificationScope14 `xml:"ModfdInvstrPrflVldtn,omitempty"`

	// Details about the MiFID classification of the account owner.
	MiFIDClassification *MiFIDClassification1 `xml:"MiFIDClssfctn,omitempty"`

	// Specifies how information is sent to the account holder.
	InformationDistribution *InformationDistribution1Code `xml:"InfDstrbtn,omitempty"`
}

func (i *InvestmentAccountOwnershipInformation7) AddParty() *Party14Choice {
	i.Party = new(Party14Choice)
	return i.Party
}

func (i *InvestmentAccountOwnershipInformation7) AddMoneyLaunderingCheck() *MoneyLaunderingCheck1Choice {
	i.MoneyLaunderingCheck = new(MoneyLaunderingCheck1Choice)
	return i.MoneyLaunderingCheck
}

func (i *InvestmentAccountOwnershipInformation7) SetOwnershipBeneficiaryRate(value string) {
	i.OwnershipBeneficiaryRate = (*PercentageRate)(&value)
}

func (i *InvestmentAccountOwnershipInformation7) SetClientIdentification(value string) {
	i.ClientIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccountOwnershipInformation7) SetFiscalExemption(value string) {
	i.FiscalExemption = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountOwnershipInformation7) SetSignatoryRightIndicator(value string) {
	i.SignatoryRightIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountOwnershipInformation7) AddModifiedInvestorProfileValidation() *ModificationScope14 {
	newValue := new(ModificationScope14)
	i.ModifiedInvestorProfileValidation = append(i.ModifiedInvestorProfileValidation, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation7) AddMiFIDClassification() *MiFIDClassification1 {
	i.MiFIDClassification = new(MiFIDClassification1)
	return i.MiFIDClassification
}

func (i *InvestmentAccountOwnershipInformation7) SetInformationDistribution(value string) {
	i.InformationDistribution = (*InformationDistribution1Code)(&value)
}

// Characteristics of the ownership of an investment account.
type InvestmentAccountOwnershipInformation8 struct {

	// Information about the organisation or individual person.
	Party *Party15Choice `xml:"Pty"`

	// Status of an identity check to prevent money laundering. This includes the counter-terrorism check.
	MoneyLaunderingCheck *MoneyLaunderingCheck1Choice `xml:"MnyLndrgChck,omitempty"`

	// Percentage of ownership or beneficiary ownership of the shares/units in the account. All subsequent subscriptions and or redemptions will be allocated using the same percentage.
	OwnershipBeneficiaryRate *PercentageRate `xml:"OwnrshBnfcryRate,omitempty"`

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	ClientIdentification *Max35Text `xml:"ClntId,omitempty"`

	// Indicates whether an owner of an investment account may benefit from a fiscal exemption or amnesty for instance for declaring overseas investments.
	FiscalExemption *YesNoIndicator `xml:"FsclXmptn,omitempty"`

	// Indicates whether the account owner signature is required to authorise transactions on the account.
	SignatoryRightIndicator *YesNoIndicator `xml:"SgntryRghtInd,omitempty"`

	// Information related to the party profile to be inserted or deleted.
	ModifiedInvestorProfileValidation []*ModificationScope19 `xml:"ModfdInvstrPrflVldtn,omitempty"`

	// Details about the MiFID classification of the account owner.
	MiFIDClassification *MiFIDClassification1 `xml:"MiFIDClssfctn,omitempty"`

	// Specifies how information is sent to the account holder.
	InformationDistribution *InformationDistribution1Code `xml:"InfDstrbtn,omitempty"`
}

func (i *InvestmentAccountOwnershipInformation8) AddParty() *Party15Choice {
	i.Party = new(Party15Choice)
	return i.Party
}

func (i *InvestmentAccountOwnershipInformation8) AddMoneyLaunderingCheck() *MoneyLaunderingCheck1Choice {
	i.MoneyLaunderingCheck = new(MoneyLaunderingCheck1Choice)
	return i.MoneyLaunderingCheck
}

func (i *InvestmentAccountOwnershipInformation8) SetOwnershipBeneficiaryRate(value string) {
	i.OwnershipBeneficiaryRate = (*PercentageRate)(&value)
}

func (i *InvestmentAccountOwnershipInformation8) SetClientIdentification(value string) {
	i.ClientIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccountOwnershipInformation8) SetFiscalExemption(value string) {
	i.FiscalExemption = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountOwnershipInformation8) SetSignatoryRightIndicator(value string) {
	i.SignatoryRightIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountOwnershipInformation8) AddModifiedInvestorProfileValidation() *ModificationScope19 {
	newValue := new(ModificationScope19)
	i.ModifiedInvestorProfileValidation = append(i.ModifiedInvestorProfileValidation, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation8) AddMiFIDClassification() *MiFIDClassification1 {
	i.MiFIDClassification = new(MiFIDClassification1)
	return i.MiFIDClassification
}

func (i *InvestmentAccountOwnershipInformation8) SetInformationDistribution(value string) {
	i.InformationDistribution = (*InformationDistribution1Code)(&value)
}

// Characteristics of the ownership of an investment account.
type InvestmentAccountOwnershipInformation9 struct {

	// Information about the organisation or individual person.
	Party *Party16Choice `xml:"Pty"`

	// Status of an identity check to prevent money laundering. This includes the counter-terrorism check.
	MoneyLaunderingCheck *MoneyLaunderingCheck1Choice `xml:"MnyLndrgChck,omitempty"`

	// Information to support Know Your Customer processes.
	InvestorProfileValidation []*PartyProfileInformation3 `xml:"InvstrPrflVldtn,omitempty"`

	// Percentage of ownership or of beneficial ownership of the shares/units in the account. All subsequent subscriptions and or redemptions will be allocated using the same percentage.
	OwnershipBeneficiaryRate *PercentageRate `xml:"OwnrshBnfcryRate,omitempty"`

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	ClientIdentification *Max35Text `xml:"ClntId,omitempty"`

	// Indicates whether an owner of an investment account may benefit from a fiscal exemption or amnesty for instance for declaring overseas investments.
	FiscalExemption *YesNoIndicator `xml:"FsclXmptn,omitempty"`

	// Indicates whether the signature of the account owner is required to authorise transactions on the account.
	SignatoryRightIndicator *YesNoIndicator `xml:"SgntryRghtInd,omitempty"`

	// Details about the MiFID classification of the account owner.
	MiFIDClassification *MiFIDClassification1 `xml:"MiFIDClssfctn,omitempty"`

	// Specifies how information is sent to the account holder.
	InformationDistribution *InformationDistribution1Code `xml:"InfDstrbtn,omitempty"`
}

func (i *InvestmentAccountOwnershipInformation9) AddParty() *Party16Choice {
	i.Party = new(Party16Choice)
	return i.Party
}

func (i *InvestmentAccountOwnershipInformation9) AddMoneyLaunderingCheck() *MoneyLaunderingCheck1Choice {
	i.MoneyLaunderingCheck = new(MoneyLaunderingCheck1Choice)
	return i.MoneyLaunderingCheck
}

func (i *InvestmentAccountOwnershipInformation9) AddInvestorProfileValidation() *PartyProfileInformation3 {
	newValue := new(PartyProfileInformation3)
	i.InvestorProfileValidation = append(i.InvestorProfileValidation, newValue)
	return newValue
}

func (i *InvestmentAccountOwnershipInformation9) SetOwnershipBeneficiaryRate(value string) {
	i.OwnershipBeneficiaryRate = (*PercentageRate)(&value)
}

func (i *InvestmentAccountOwnershipInformation9) SetClientIdentification(value string) {
	i.ClientIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccountOwnershipInformation9) SetFiscalExemption(value string) {
	i.FiscalExemption = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountOwnershipInformation9) SetSignatoryRightIndicator(value string) {
	i.SignatoryRightIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountOwnershipInformation9) AddMiFIDClassification() *MiFIDClassification1 {
	i.MiFIDClassification = new(MiFIDClassification1)
	return i.MiFIDClassification
}

func (i *InvestmentAccountOwnershipInformation9) SetInformationDistribution(value string) {
	i.InformationDistribution = (*InformationDistribution1Code)(&value)
}
