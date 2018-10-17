package iso20022

// Scope of the modification to be applied on an identified set of information.
type ModificationScope1 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress *PostalAddress3 `xml:"PstlAdr"`
}

func (m *ModificationScope1) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope1) AddPostalAddress() *PostalAddress3 {
	m.PostalAddress = new(PostalAddress3)
	return m.PostalAddress
}

// Scope of the modification to be applied on an identified set of information.
type ModificationScope10 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Identification of information which is part of a service level agreement.
	ServiceLevelAgreement *DocumentToSend1 `xml:"SvcLvlAgrmt"`
}

func (m *ModificationScope10) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope10) AddServiceLevelAgreement() *DocumentToSend1 {
	m.ServiceLevelAgreement = new(DocumentToSend1)
	return m.ServiceLevelAgreement
}

// Scope of the modification to be applied on an identified set of information.
type ModificationScope11 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification2Code `xml:"ModScpIndctn"`

	// Detailed information about the party profile information.
	InvestorProfileValidation *PartyProfileInformation1 `xml:"InvstrPrflVldtn"`
}

func (m *ModificationScope11) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification2Code)(&value)
}

func (m *ModificationScope11) AddInvestorProfileValidation() *PartyProfileInformation1 {
	m.InvestorProfileValidation = new(PartyProfileInformation1)
	return m.InvestorProfileValidation
}

// Scope of the modification to be applied on an identified set of information.
type ModificationScope12 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification2Code `xml:"ModScpIndctn"`

	// Information related to the account to be modified.
	FundDetails *FinancialInstrument10 `xml:"FndDtls"`
}

func (m *ModificationScope12) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification2Code)(&value)
}

func (m *ModificationScope12) AddFundDetails() *FinancialInstrument10 {
	m.FundDetails = new(FinancialInstrument10)
	return m.FundDetails
}

// Scope of the modification to be applied on an identified set of information.
type ModificationScope13 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification2Code `xml:"ModScpIndctn"`

	// Information related to the account to be modified.
	FundDetails *FinancialInstrument29 `xml:"FndDtls"`
}

func (m *ModificationScope13) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification2Code)(&value)
}

func (m *ModificationScope13) AddFundDetails() *FinancialInstrument29 {
	m.FundDetails = new(FinancialInstrument29)
	return m.FundDetails
}

// Scope of the modification to be applied on an identified set of information.
type ModificationScope14 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification2Code `xml:"ModScpIndctn"`

	// Detailed information about the party profile information.
	InvestorProfileValidation *PartyProfileInformation2 `xml:"InvstrPrflVldtn"`
}

func (m *ModificationScope14) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification2Code)(&value)
}

func (m *ModificationScope14) AddInvestorProfileValidation() *PartyProfileInformation2 {
	m.InvestorProfileValidation = new(PartyProfileInformation2)
	return m.InvestorProfileValidation
}

// Scope of the modification to be applied on an identified set of information.
type ModificationScope16 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Plan that allows investors to schedule periodical investments or divestments, according to pre-defined criteria.
	InvestmentPlan *InvestmentPlan7 `xml:"InvstmtPlan"`
}

func (m *ModificationScope16) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope16) AddInvestmentPlan() *InvestmentPlan7 {
	m.InvestmentPlan = new(InvestmentPlan7)
	return m.InvestmentPlan
}

// Specifies the type of modification to be applied on a data set.
type ModificationScope17 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification2Code `xml:"ModScpIndctn"`

	// Alternative identification, for example, national registration identification number, passport number, or an account number used to further identify the beneficial owner, for example, a Central Provident Fund (CFP) account as required for Singapore.
	OtherIdentification *GenericIdentification44 `xml:"OthrId"`
}

func (m *ModificationScope17) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification2Code)(&value)
}

func (m *ModificationScope17) AddOtherIdentification() *GenericIdentification44 {
	m.OtherIdentification = new(GenericIdentification44)
	return m.OtherIdentification
}

// Scope of the modification to be applied on an identified set of information.
type ModificationScope18 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Plan that allows investors to schedule periodical investments or divestments, according to pre-defined criteria.
	InvestmentPlan *InvestmentPlan8 `xml:"InvstmtPlan"`
}

func (m *ModificationScope18) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope18) AddInvestmentPlan() *InvestmentPlan8 {
	m.InvestmentPlan = new(InvestmentPlan8)
	return m.InvestmentPlan
}

// Scope of the modification to be applied on an identified set of information.
type ModificationScope19 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification2Code `xml:"ModScpIndctn"`

	// Detailed information about the party profile information.
	InvestorProfileValidation *PartyProfileInformation3 `xml:"InvstrPrflVldtn"`
}

func (m *ModificationScope19) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification2Code)(&value)
}

func (m *ModificationScope19) AddInvestorProfileValidation() *PartyProfileInformation3 {
	m.InvestorProfileValidation = new(PartyProfileInformation3)
	return m.InvestorProfileValidation
}

// Scope of the modification to be applied on an identified set of information.
type ModificationScope2 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification2Code `xml:"ModScpIndctn"`

	// Alternative identification, for example, national registration identification number, passport number, or an account number used to further identify the beneficial owner, for example, a Central Provident Fund (CFP) account as required for Singapore.
	OtherIdentification *GenericIdentification9 `xml:"OthrId"`
}

func (m *ModificationScope2) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification2Code)(&value)
}

func (m *ModificationScope2) AddOtherIdentification() *GenericIdentification9 {
	m.OtherIdentification = new(GenericIdentification9)
	return m.OtherIdentification
}

// Scope of the modification to be applied on an identified set of information.
type ModificationScope20 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Information related to intermediaries.
	Intermediary *Intermediary24 `xml:"Intrmy"`
}

func (m *ModificationScope20) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope20) AddIntermediary() *Intermediary24 {
	m.Intermediary = new(Intermediary24)
	return m.Intermediary
}

// Scope of the modification to be applied on an identified set of information.
type ModificationScope21 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Information about the investment account ownership with respect to new issue allocation for a hedge fund.
	IssueAllocation *NewIssueAllocation2 `xml:"IsseAllcn"`
}

func (m *ModificationScope21) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope21) AddIssueAllocation() *NewIssueAllocation2 {
	m.IssueAllocation = new(NewIssueAllocation2)
	return m.IssueAllocation
}

// Scope of the modification to be applied on an identified set of information.
type ModificationScope22 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Identification of information which is part of a service level agreement.
	ServiceLevelAgreement *DocumentToSend2 `xml:"SvcLvlAgrmt"`
}

func (m *ModificationScope22) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope22) AddServiceLevelAgreement() *DocumentToSend2 {
	m.ServiceLevelAgreement = new(DocumentToSend2)
	return m.ServiceLevelAgreement
}

// Specifies the type of modification to be applied on a data set.
type ModificationScope23 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification2Code `xml:"ModScpIndctn"`

	// Alternative identification, for example, national registration identification number, passport number, or an account number used to further identify the beneficial owner, for example, a Central Provident Fund (CFP) account as required for Singapore.
	OtherIdentification *GenericIdentification55 `xml:"OthrId"`
}

func (m *ModificationScope23) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification2Code)(&value)
}

func (m *ModificationScope23) AddOtherIdentification() *GenericIdentification55 {
	m.OtherIdentification = new(GenericIdentification55)
	return m.OtherIdentification
}

// Scope of the modification to be applied on an identified set of information.
type ModificationScope25 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Plan that allows investors to schedule periodical investments or divestments, according to pre-defined criteria.
	InvestmentPlan *InvestmentPlan11 `xml:"InvstmtPlan"`
}

func (m *ModificationScope25) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope25) AddInvestmentPlan() *InvestmentPlan11 {
	m.InvestmentPlan = new(InvestmentPlan11)
	return m.InvestmentPlan
}

// Scope of the modification to be applied on an identified set of information.
type ModificationScope26 struct {

	// Specifies the type of modification to be applied.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Intermediary or other party related to the management of the account. In some markets, when this intermediary is a party acting on behalf of the investor for which it has opened an account at, for example, a central securities depository or international central securities depository, this party is known by the investor as the 'account controller'.
	Intermediary *Intermediary36 `xml:"Intrmy"`
}

func (m *ModificationScope26) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope26) AddIntermediary() *Intermediary36 {
	m.Intermediary = new(Intermediary36)
	return m.Intermediary
}

// Scope of the modification to be applied on an identified set of information.
type ModificationScope27 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification2Code `xml:"ModScpIndctn"`

	// Detailed information about the party profile information.
	InvestorProfileValidation *PartyProfileInformation5 `xml:"InvstrPrflVldtn"`
}

func (m *ModificationScope27) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification2Code)(&value)
}

func (m *ModificationScope27) AddInvestorProfileValidation() *PartyProfileInformation5 {
	m.InvestorProfileValidation = new(PartyProfileInformation5)
	return m.InvestorProfileValidation
}

// Scope of the modification to be applied on an identified set of information.
type ModificationScope28 struct {

	// Specifies the type of modification to be applied.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Plan that allows investors to schedule periodical investments or divestments, according to pre-defined criteria.
	InvestmentPlan *InvestmentPlan13 `xml:"InvstmtPlan"`
}

func (m *ModificationScope28) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope28) AddInvestmentPlan() *InvestmentPlan13 {
	m.InvestmentPlan = new(InvestmentPlan13)
	return m.InvestmentPlan
}

// Scope of the modification to be applied on an identified set of information.
type ModificationScope29 struct {

	// Specifies the type of modification to be applied.
	ModificationScopeIndication *DataModification2Code `xml:"ModScpIndctn"`

	// Detailed information about the investment fund or security associated to the account.
	FinancialInstrumentDetails *FinancialInstrument51 `xml:"FinInstrmDtls"`
}

func (m *ModificationScope29) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification2Code)(&value)
}

func (m *ModificationScope29) AddFinancialInstrumentDetails() *FinancialInstrument51 {
	m.FinancialInstrumentDetails = new(FinancialInstrument51)
	return m.FinancialInstrumentDetails
}

// Scope of the modification to be applied on an identified set of information.
type ModificationScope3 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification2Code `xml:"ModScpIndctn"`

	// Information about the nationality and the legal status (minor or major) of a person.
	Citizenship *CitizenshipInformation `xml:"Ctznsh"`
}

func (m *ModificationScope3) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification2Code)(&value)
}

func (m *ModificationScope3) AddCitizenship() *CitizenshipInformation {
	m.Citizenship = new(CitizenshipInformation)
	return m.Citizenship
}

// Scope of the modification to be applied on an identified set of information.
type ModificationScope30 struct {

	// Specifies the type of modification to be applied.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Additional information concerning limitations and restrictions on the account.
	AdditionalInformation []*AccountRestrictions1 `xml:"AddtlInf"`
}

func (m *ModificationScope30) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope30) AddAdditionalInformation() *AccountRestrictions1 {
	newValue := new(AccountRestrictions1)
	m.AdditionalInformation = append(m.AdditionalInformation, newValue)
	return newValue
}

// Scope of the modification to be applied on an identified set of information.
type ModificationScope31 struct {

	// Specifies the type of modification to be applied.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Identification of information which is part of a service level agreement.
	ServiceLevelAgreement *DocumentToSend3 `xml:"SvcLvlAgrmt"`
}

func (m *ModificationScope31) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope31) AddServiceLevelAgreement() *DocumentToSend3 {
	m.ServiceLevelAgreement = new(DocumentToSend3)
	return m.ServiceLevelAgreement
}

// Scope of the modification to be applied on an identified set of information.
type ModificationScope32 struct {

	// Specifies the type of modification to be applied.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Information used to determine fees and types of operations that can be carried out on the account.
	InvestorProfile *InvestorProfile1 `xml:"InvstrPrfl"`
}

func (m *ModificationScope32) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope32) AddInvestorProfile() *InvestorProfile1 {
	m.InvestorProfile = new(InvestorProfile1)
	return m.InvestorProfile
}

// Scope of the modification to be applied on an identified set of information.
type ModificationScope33 struct {

	// Specifies the type of modification to be applied.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Referral information.
	Placement *ReferredAgent2 `xml:"Plcmnt"`
}

func (m *ModificationScope33) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope33) AddPlacement() *ReferredAgent2 {
	m.Placement = new(ReferredAgent2)
	return m.Placement
}

// Scope of the modification to be applied on an identified set of information.
type ModificationScope34 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Address of the organisation.
	PostalAddress *PostalAddress21 `xml:"PstlAdr"`
}

func (m *ModificationScope34) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope34) AddPostalAddress() *PostalAddress21 {
	m.PostalAddress = new(PostalAddress21)
	return m.PostalAddress
}

// Scope of the modification to be applied on an identified set of information.
type ModificationScope35 struct {

	// Type of modification to be applied.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Additional information such as remarks or notes that must be conveyed about the party and or  limitations and restrictions.
	AdditionalInformation []*AdditiononalInformation12 `xml:"AddtlInf"`
}

func (m *ModificationScope35) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope35) AddAdditionalInformation() *AdditiononalInformation12 {
	newValue := new(AdditiononalInformation12)
	m.AdditionalInformation = append(m.AdditionalInformation, newValue)
	return newValue
}

// Scope of the modification to be applied on an identified set of information.
type ModificationScope36 struct {

	// Type of modification to be applied.
	ModificationScopeIndication *DataModification2Code `xml:"ModScpIndctn"`

	// Detailed information about the investment fund or security associated to the account.
	FinancialInstrumentDetails *FinancialInstrument56 `xml:"FinInstrmDtls"`
}

func (m *ModificationScope36) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification2Code)(&value)
}

func (m *ModificationScope36) AddFinancialInstrumentDetails() *FinancialInstrument56 {
	m.FinancialInstrumentDetails = new(FinancialInstrument56)
	return m.FinancialInstrumentDetails
}

// Scope of the modification to be applied on an identified set of information.
type ModificationScope37 struct {

	// Type of modification to be applied.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Plan that allows investors to schedule periodical investments or divestments, according to pre-defined criteria.
	InvestmentPlan *InvestmentPlan15 `xml:"InvstmtPlan"`
}

func (m *ModificationScope37) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope37) AddInvestmentPlan() *InvestmentPlan15 {
	m.InvestmentPlan = new(InvestmentPlan15)
	return m.InvestmentPlan
}

// Scope of the modification to be applied on an identified set of information.
type ModificationScope38 struct {

	// Type of modification to be applied.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Intermediary or other party related to the management of the account.
	Intermediary *Intermediary36 `xml:"Intrmy"`
}

func (m *ModificationScope38) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope38) AddIntermediary() *Intermediary36 {
	m.Intermediary = new(Intermediary36)
	return m.Intermediary
}

// Scope of the modification to be applied on an identified set of information.
type ModificationScope39 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification2Code `xml:"ModScpIndctn"`

	// Nationality and legal status (minor or major).
	Citizenship *CitizenshipInformation2 `xml:"Ctznsh"`
}

func (m *ModificationScope39) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification2Code)(&value)
}

func (m *ModificationScope39) AddCitizenship() *CitizenshipInformation2 {
	m.Citizenship = new(CitizenshipInformation2)
	return m.Citizenship
}

// Scope of the modification to be applied on an identified set of information.
type ModificationScope7 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Information related to intermediaries.
	Intermediary *Intermediary13 `xml:"Intrmy"`
}

func (m *ModificationScope7) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope7) AddIntermediary() *Intermediary13 {
	m.Intermediary = new(Intermediary13)
	return m.Intermediary
}

// Scope of the modification to be applied on an identified set of information.
type ModificationScope8 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Plan that allows investors to schedule periodical investments or divestments, according to pre-defined criteria.
	InvestmentPlan *InvestmentPlan5 `xml:"InvstmtPlan"`
}

func (m *ModificationScope8) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope8) AddInvestmentPlan() *InvestmentPlan5 {
	m.InvestmentPlan = new(InvestmentPlan5)
	return m.InvestmentPlan
}

// Scope of the modification to be applied on an identified set of information.
type ModificationScope9 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Information about the investment account ownership with respect to new issue allocation for a hedge fund.
	IssueAllocation *NewIssueAllocation1 `xml:"IsseAllcn"`
}

func (m *ModificationScope9) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (m *ModificationScope9) AddIssueAllocation() *NewIssueAllocation1 {
	m.IssueAllocation = new(NewIssueAllocation1)
	return m.IssueAllocation
}
