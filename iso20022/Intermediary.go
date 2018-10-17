package iso20022

// Organised structure that is set up for a particular purpose, eg, a business, government body, department, charity, or financial institution.
type Intermediary1 struct {

	// Unique and unambiguous identifier for an organisation that is allocated by an institution, eg, Dun & Bradstreet Identification.
	Identification *PartyIdentification1Choice `xml:"Id"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account1 `xml:"Acct,omitempty"`

	// Set of  functions performed by an intermediary in a given situation.
	Role *Max35Text `xml:"Role,omitempty"`
}

func (i *Intermediary1) AddIdentification() *PartyIdentification1Choice {
	i.Identification = new(PartyIdentification1Choice)
	return i.Identification
}

func (i *Intermediary1) AddAccount() *Account1 {
	i.Account = new(Account1)
	return i.Account
}

func (i *Intermediary1) SetRole(value string) {
	i.Role = (*Max35Text)(&value)
}

// Party that provides services to investors relating to financial products.
type Intermediary10 struct {

	// Unique and unambiguous identifier of the intermediary.
	Identification *PartyIdentification2Choice `xml:"Id"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account7 `xml:"Acct,omitempty"`

	// Function performed by the intermediary.
	Role *InvestmentFundRole2Code `xml:"Role,omitempty"`

	// Function performed by the intermediary.
	ExtendedRole *Extended350Code `xml:"XtndedRole,omitempty"`

	// Contact person and contact information.
	ContactPerson *ContactIdentification2 `xml:"CtctPrsn,omitempty"`
}

func (i *Intermediary10) AddIdentification() *PartyIdentification2Choice {
	i.Identification = new(PartyIdentification2Choice)
	return i.Identification
}

func (i *Intermediary10) AddAccount() *Account7 {
	i.Account = new(Account7)
	return i.Account
}

func (i *Intermediary10) SetRole(value string) {
	i.Role = (*InvestmentFundRole2Code)(&value)
}

func (i *Intermediary10) SetExtendedRole(value string) {
	i.ExtendedRole = (*Extended350Code)(&value)
}

func (i *Intermediary10) AddContactPerson() *ContactIdentification2 {
	i.ContactPerson = new(ContactIdentification2)
	return i.ContactPerson
}

// Party that provides services to investors relating to financial products.
type Intermediary11 struct {

	// Unique and unambiguous identifier of the intermediary.
	Identification *PartyIdentification2Choice `xml:"Id"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account7 `xml:"Acct,omitempty"`

	// Function performed by the intermediary.
	Role *InvestmentFundRole2Code `xml:"Role,omitempty"`

	// Function performed by the intermediary.
	ExtendedRole *Extended350Code `xml:"XtndedRole,omitempty"`
}

func (i *Intermediary11) AddIdentification() *PartyIdentification2Choice {
	i.Identification = new(PartyIdentification2Choice)
	return i.Identification
}

func (i *Intermediary11) AddAccount() *Account7 {
	i.Account = new(Account7)
	return i.Account
}

func (i *Intermediary11) SetRole(value string) {
	i.Role = (*InvestmentFundRole2Code)(&value)
}

func (i *Intermediary11) SetExtendedRole(value string) {
	i.ExtendedRole = (*Extended350Code)(&value)
}

// Party that provides services relating to financial products to investors, eg, advice on products and placement of orders for the investment fund.
type Intermediary12 struct {

	// Unique and unambiguous identifier for an organisation that is allocated by an institution, eg, Dun & Bradstreet Identification.
	Identification *PartyIdentification4Choice `xml:"Id"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account2 `xml:"Acct,omitempty"`

	// Non-enforcement of the right to all or part of a commission by the party entitled to the commission.
	WaivedTrailerCommissionIndicator *YesNoIndicator `xml:"WvdTrlrComssnInd,omitempty"`

	// The role or function performed by an intermediary in a given situation.
	Role *InvestmentFundRole3Code `xml:"Role,omitempty"`

	// The role or function performed by an intermediary in a given situation.
	ExtendedRole *Extended350Code `xml:"XtndedRole,omitempty"`

	// Communication device number or electronic address used for communication.
	PrimaryCommunicationAddress *CommunicationAddress3 `xml:"PmryComAdr,omitempty"`

	// Communication device number or electronic address used for communication.
	SecondaryCommunicationAddress *CommunicationAddress3 `xml:"ScndryComAdr,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	NameAndAddress *NameAndAddress4 `xml:"NmAndAdr,omitempty"`
}

func (i *Intermediary12) AddIdentification() *PartyIdentification4Choice {
	i.Identification = new(PartyIdentification4Choice)
	return i.Identification
}

func (i *Intermediary12) AddAccount() *Account2 {
	i.Account = new(Account2)
	return i.Account
}

func (i *Intermediary12) SetWaivedTrailerCommissionIndicator(value string) {
	i.WaivedTrailerCommissionIndicator = (*YesNoIndicator)(&value)
}

func (i *Intermediary12) SetRole(value string) {
	i.Role = (*InvestmentFundRole3Code)(&value)
}

func (i *Intermediary12) SetExtendedRole(value string) {
	i.ExtendedRole = (*Extended350Code)(&value)
}

func (i *Intermediary12) AddPrimaryCommunicationAddress() *CommunicationAddress3 {
	i.PrimaryCommunicationAddress = new(CommunicationAddress3)
	return i.PrimaryCommunicationAddress
}

func (i *Intermediary12) AddSecondaryCommunicationAddress() *CommunicationAddress3 {
	i.SecondaryCommunicationAddress = new(CommunicationAddress3)
	return i.SecondaryCommunicationAddress
}

func (i *Intermediary12) AddNameAndAddress() *NameAndAddress4 {
	i.NameAndAddress = new(NameAndAddress4)
	return i.NameAndAddress
}

// Party that provides services relating to financial products to investors, eg, advice on products and placement of orders for the investment fund.
type Intermediary13 struct {

	// Unique and unambiguous identifier for an organisation that is allocated by an institution, eg, Dun & Bradstreet Identification.
	Identification *PartyIdentification4Choice `xml:"Id"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account2 `xml:"Acct,omitempty"`

	// Non-enforcement of the right to all or part of a commission by the party entitled to the commission.
	WaivedTrailerCommissionIndicator *YesNoIndicator `xml:"WvdTrlrComssnInd,omitempty"`

	// Role or function performed by an intermediary in a given situation.
	Role *InvestmentFundRole3Code `xml:"Role,omitempty"`

	// Role or function performed by an intermediary in a given situation.
	ExtendedRole *Extended350Code `xml:"XtndedRole,omitempty"`

	// Communication device number or electronic address used for communication.
	PrimaryCommunicationAddress *CommunicationAddress3 `xml:"PmryComAdr,omitempty"`

	// Communication device number or electronic address used for communication.
	SecondaryCommunicationAddress *CommunicationAddress3 `xml:"ScndryComAdr,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	NameAndAddress *NameAndAddress4 `xml:"NmAndAdr,omitempty"`
}

func (i *Intermediary13) AddIdentification() *PartyIdentification4Choice {
	i.Identification = new(PartyIdentification4Choice)
	return i.Identification
}

func (i *Intermediary13) AddAccount() *Account2 {
	i.Account = new(Account2)
	return i.Account
}

func (i *Intermediary13) SetWaivedTrailerCommissionIndicator(value string) {
	i.WaivedTrailerCommissionIndicator = (*YesNoIndicator)(&value)
}

func (i *Intermediary13) SetRole(value string) {
	i.Role = (*InvestmentFundRole3Code)(&value)
}

func (i *Intermediary13) SetExtendedRole(value string) {
	i.ExtendedRole = (*Extended350Code)(&value)
}

func (i *Intermediary13) AddPrimaryCommunicationAddress() *CommunicationAddress3 {
	i.PrimaryCommunicationAddress = new(CommunicationAddress3)
	return i.PrimaryCommunicationAddress
}

func (i *Intermediary13) AddSecondaryCommunicationAddress() *CommunicationAddress3 {
	i.SecondaryCommunicationAddress = new(CommunicationAddress3)
	return i.SecondaryCommunicationAddress
}

func (i *Intermediary13) AddNameAndAddress() *NameAndAddress4 {
	i.NameAndAddress = new(NameAndAddress4)
	return i.NameAndAddress
}

// Party that provides services to investors relating to financial products (Investment Funds).
type Intermediary2 struct {

	// Unique and unambiguous identifier of the intermediary.
	Identification *PartyIdentification10Choice `xml:"Id"`

	// Function performed by the intermediary (investment funds).
	Role *Role2Choice `xml:"Role,omitempty"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account9 `xml:"Acct,omitempty"`
}

func (i *Intermediary2) AddIdentification() *PartyIdentification10Choice {
	i.Identification = new(PartyIdentification10Choice)
	return i.Identification
}

func (i *Intermediary2) AddRole() *Role2Choice {
	i.Role = new(Role2Choice)
	return i.Role
}

func (i *Intermediary2) AddAccount() *Account9 {
	i.Account = new(Account9)
	return i.Account
}

// Party that provides services to investors relating to financial products (Investment Funds).
type Intermediary21 struct {

	// Unique and unambiguous identifier of the intermediary.
	Identification *PartyIdentification49Choice `xml:"Id"`

	// Function performed by the intermediary (investment funds).
	Role *Role2Choice `xml:"Role,omitempty"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account11 `xml:"Acct,omitempty"`
}

func (i *Intermediary21) AddIdentification() *PartyIdentification49Choice {
	i.Identification = new(PartyIdentification49Choice)
	return i.Identification
}

func (i *Intermediary21) AddRole() *Role2Choice {
	i.Role = new(Role2Choice)
	return i.Role
}

func (i *Intermediary21) AddAccount() *Account11 {
	i.Account = new(Account11)
	return i.Account
}

// Party that provides services to investors relating to financial products (Investment Funds).
type Intermediary23 struct {

	// Unique and unambiguous identifier of the intermediary.
	Identification *PartyIdentification49Choice `xml:"Id"`

	// Function performed by the intermediary (investment funds).
	Role *Role2Choice `xml:"Role,omitempty"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account13 `xml:"Acct,omitempty"`
}

func (i *Intermediary23) AddIdentification() *PartyIdentification49Choice {
	i.Identification = new(PartyIdentification49Choice)
	return i.Identification
}

func (i *Intermediary23) AddRole() *Role2Choice {
	i.Role = new(Role2Choice)
	return i.Role
}

func (i *Intermediary23) AddAccount() *Account13 {
	i.Account = new(Account13)
	return i.Account
}

// Party that provides services to investors relating to financial products.
type Intermediary24 struct {

	// Unique and unambiguous identifier for an organisation that is allocated by an institution, eg, Dun & Bradstreet Identification.
	Identification *PartyIdentification4Choice `xml:"Id"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account2 `xml:"Acct,omitempty"`

	// Non-enforcement of the right to all or part of a commission by the party entitled to the commission.
	WaivedTrailerCommissionIndicator *YesNoIndicator `xml:"WvdTrlrComssnInd,omitempty"`

	// Role or function performed by the intermediary.
	Role *PartyRole2Choice `xml:"Role,omitempty"`

	// Communication device number or electronic address used for communication.
	PrimaryCommunicationAddress *CommunicationAddress3 `xml:"PmryComAdr,omitempty"`

	// Communication device number or electronic address used for communication.
	SecondaryCommunicationAddress *CommunicationAddress3 `xml:"ScndryComAdr,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	NameAndAddress *NameAndAddress4 `xml:"NmAndAdr,omitempty"`
}

func (i *Intermediary24) AddIdentification() *PartyIdentification4Choice {
	i.Identification = new(PartyIdentification4Choice)
	return i.Identification
}

func (i *Intermediary24) AddAccount() *Account2 {
	i.Account = new(Account2)
	return i.Account
}

func (i *Intermediary24) SetWaivedTrailerCommissionIndicator(value string) {
	i.WaivedTrailerCommissionIndicator = (*YesNoIndicator)(&value)
}

func (i *Intermediary24) AddRole() *PartyRole2Choice {
	i.Role = new(PartyRole2Choice)
	return i.Role
}

func (i *Intermediary24) AddPrimaryCommunicationAddress() *CommunicationAddress3 {
	i.PrimaryCommunicationAddress = new(CommunicationAddress3)
	return i.PrimaryCommunicationAddress
}

func (i *Intermediary24) AddSecondaryCommunicationAddress() *CommunicationAddress3 {
	i.SecondaryCommunicationAddress = new(CommunicationAddress3)
	return i.SecondaryCommunicationAddress
}

func (i *Intermediary24) AddNameAndAddress() *NameAndAddress4 {
	i.NameAndAddress = new(NameAndAddress4)
	return i.NameAndAddress
}

// Party that provides services to investors relating to financial products.
type Intermediary25 struct {

	// Unique and unambiguous identifier of the intermediary.
	Identification *PartyIdentification2Choice `xml:"Id"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account14 `xml:"Acct,omitempty"`

	// Function performed by the intermediary.
	Role *Role4Choice `xml:"Role,omitempty"`
}

func (i *Intermediary25) AddIdentification() *PartyIdentification2Choice {
	i.Identification = new(PartyIdentification2Choice)
	return i.Identification
}

func (i *Intermediary25) AddAccount() *Account14 {
	i.Account = new(Account14)
	return i.Account
}

func (i *Intermediary25) AddRole() *Role4Choice {
	i.Role = new(Role4Choice)
	return i.Role
}

// Party that provides services to investors relating to financial products.
type Intermediary26 struct {

	// Unique and unambiguous identifier of the intermediary.
	Identification *PartyIdentification2Choice `xml:"Id"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account14 `xml:"Acct,omitempty"`

	// Function performed by the intermediary.
	Role *Role4Choice `xml:"Role,omitempty"`

	// Contact person and contact information.
	ContactPerson *ContactIdentification2 `xml:"CtctPrsn,omitempty"`
}

func (i *Intermediary26) AddIdentification() *PartyIdentification2Choice {
	i.Identification = new(PartyIdentification2Choice)
	return i.Identification
}

func (i *Intermediary26) AddAccount() *Account14 {
	i.Account = new(Account14)
	return i.Account
}

func (i *Intermediary26) AddRole() *Role4Choice {
	i.Role = new(Role4Choice)
	return i.Role
}

func (i *Intermediary26) AddContactPerson() *ContactIdentification2 {
	i.ContactPerson = new(ContactIdentification2)
	return i.ContactPerson
}

// Party that provides services to investors relating to financial products.
type Intermediary27 struct {

	// Unique and unambiguous identifier of the intermediary.
	Identification *PartyIdentification2Choice `xml:"Id"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account7 `xml:"Acct,omitempty"`

	// Function performed by the intermediary.
	Role *Role4Choice `xml:"Role,omitempty"`
}

func (i *Intermediary27) AddIdentification() *PartyIdentification2Choice {
	i.Identification = new(PartyIdentification2Choice)
	return i.Identification
}

func (i *Intermediary27) AddAccount() *Account7 {
	i.Account = new(Account7)
	return i.Account
}

func (i *Intermediary27) AddRole() *Role4Choice {
	i.Role = new(Role4Choice)
	return i.Role
}

// Party that provides services to investors relating to financial products.
type Intermediary29 struct {

	// Unique and unambiguous identifier of the intermediary.
	Identification *PartyIdentification100 `xml:"Id"`

	// Function performed by the intermediary.
	Role *Role5Choice `xml:"Role"`

	// Counterparties eligibility as defined by article 24 of the EU MiFID Directive applicable to transactions executed by investment firms for eligible counterparties.
	OrderOriginatorEligibility *OrderOriginatorEligibility1Code `xml:"OrdrOrgtrElgblty,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *Intermediary29) AddIdentification() *PartyIdentification100 {
	i.Identification = new(PartyIdentification100)
	return i.Identification
}

func (i *Intermediary29) AddRole() *Role5Choice {
	i.Role = new(Role5Choice)
	return i.Role
}

func (i *Intermediary29) SetOrderOriginatorEligibility(value string) {
	i.OrderOriginatorEligibility = (*OrderOriginatorEligibility1Code)(&value)
}

func (i *Intermediary29) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}

// Party that provides services to investors relating to financial products (Investment Funds).
type Intermediary32 struct {

	// Unique and unambiguous identifier of the intermediary.
	Identification *PartyIdentification100 `xml:"Id"`

	// Function performed by the intermediary (investment funds).
	Role *Role6Choice `xml:"Role,omitempty"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account18 `xml:"Acct,omitempty"`
}

func (i *Intermediary32) AddIdentification() *PartyIdentification100 {
	i.Identification = new(PartyIdentification100)
	return i.Identification
}

func (i *Intermediary32) AddRole() *Role6Choice {
	i.Role = new(Role6Choice)
	return i.Role
}

func (i *Intermediary32) AddAccount() *Account18 {
	i.Account = new(Account18)
	return i.Account
}

// Identification of a party and its role.
type Intermediary33 struct {

	// Unique and unambiguous identifier for an organisation that is allocated by an institution.
	Identification *PartyIdentification70Choice `xml:"Id"`

	// Identification of the party with a Legal Entity Identifier. This is a code allocated to a party as described in ISO 17442 "Financial Services - Legal Entity Identifier (LEI)".
	LegalEntityIdentifier *LEIIdentifier `xml:"LglNttyIdr,omitempty"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account20 `xml:"Acct,omitempty"`
}

func (i *Intermediary33) AddIdentification() *PartyIdentification70Choice {
	i.Identification = new(PartyIdentification70Choice)
	return i.Identification
}

func (i *Intermediary33) SetLegalEntityIdentifier(value string) {
	i.LegalEntityIdentifier = (*LEIIdentifier)(&value)
}

func (i *Intermediary33) AddAccount() *Account20 {
	i.Account = new(Account20)
	return i.Account
}

// Party that provides services to investors relating to financial products.
type Intermediary34 struct {

	// Unique and unambiguous identifier of the intermediary.
	Identification *PartyIdentification70Choice `xml:"Id"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account20 `xml:"Acct,omitempty"`

	// Function performed by the intermediary.
	Role *Role4Choice `xml:"Role,omitempty"`

	// Contact person and contact information.
	ContactPerson *ContactIdentification2 `xml:"CtctPrsn,omitempty"`
}

func (i *Intermediary34) AddIdentification() *PartyIdentification70Choice {
	i.Identification = new(PartyIdentification70Choice)
	return i.Identification
}

func (i *Intermediary34) AddAccount() *Account20 {
	i.Account = new(Account20)
	return i.Account
}

func (i *Intermediary34) AddRole() *Role4Choice {
	i.Role = new(Role4Choice)
	return i.Role
}

func (i *Intermediary34) AddContactPerson() *ContactIdentification2 {
	i.ContactPerson = new(ContactIdentification2)
	return i.ContactPerson
}

// Party that provides services to investors relating to financial products.
type Intermediary35 struct {

	// Unique and unambiguous identifier of the intermediary.
	Identification *PartyIdentification70Choice `xml:"Id"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account20 `xml:"Acct,omitempty"`

	// Function performed by the intermediary.
	Role *Role4Choice `xml:"Role,omitempty"`
}

func (i *Intermediary35) AddIdentification() *PartyIdentification70Choice {
	i.Identification = new(PartyIdentification70Choice)
	return i.Identification
}

func (i *Intermediary35) AddAccount() *Account20 {
	i.Account = new(Account20)
	return i.Account
}

func (i *Intermediary35) AddRole() *Role4Choice {
	i.Role = new(Role4Choice)
	return i.Role
}

// Identification of a party and its role.
type Intermediary36 struct {

	// Unique and unambiguous identifier for an organisation that is allocated by an institution.
	Identification *PartyIdentification72Choice `xml:"Id"`

	// Identification of the organisation with a Legal Entity Identifier. This is a code allocated to a party as described in ISO 17442 "Financial Services - Legal Entity Identifier (LEI)".
	LegalEntityIdentifier *LEIIdentifier `xml:"LglNttyIdr,omitempty"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account20 `xml:"Acct,omitempty"`

	// Non-enforcement of the right to all or part of a commission by the party entitled to the commission.
	WaivedTrailerCommissionIndicator *YesNoIndicator `xml:"WvdTrlrComssnInd,omitempty"`

	// Role or function performed by the intermediary.
	Role *PartyRole2Choice `xml:"Role,omitempty"`

	// Communication device number or electronic address used for communication.
	PrimaryCommunicationAddress []*CommunicationAddress6 `xml:"PmryComAdr,omitempty"`

	// Communication device number or electronic address used for communication.
	SecondaryCommunicationAddress []*CommunicationAddress6 `xml:"ScndryComAdr,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	NameAndAddress *NameAndAddress4 `xml:"NmAndAdr,omitempty"`
}

func (i *Intermediary36) AddIdentification() *PartyIdentification72Choice {
	i.Identification = new(PartyIdentification72Choice)
	return i.Identification
}

func (i *Intermediary36) SetLegalEntityIdentifier(value string) {
	i.LegalEntityIdentifier = (*LEIIdentifier)(&value)
}

func (i *Intermediary36) AddAccount() *Account20 {
	i.Account = new(Account20)
	return i.Account
}

func (i *Intermediary36) SetWaivedTrailerCommissionIndicator(value string) {
	i.WaivedTrailerCommissionIndicator = (*YesNoIndicator)(&value)
}

func (i *Intermediary36) AddRole() *PartyRole2Choice {
	i.Role = new(PartyRole2Choice)
	return i.Role
}

func (i *Intermediary36) AddPrimaryCommunicationAddress() *CommunicationAddress6 {
	newValue := new(CommunicationAddress6)
	i.PrimaryCommunicationAddress = append(i.PrimaryCommunicationAddress, newValue)
	return newValue
}

func (i *Intermediary36) AddSecondaryCommunicationAddress() *CommunicationAddress6 {
	newValue := new(CommunicationAddress6)
	i.SecondaryCommunicationAddress = append(i.SecondaryCommunicationAddress, newValue)
	return newValue
}

func (i *Intermediary36) AddNameAndAddress() *NameAndAddress4 {
	i.NameAndAddress = new(NameAndAddress4)
	return i.NameAndAddress
}

// Party that provides services to investors relating to financial products (Investment Funds).
type Intermediary37 struct {

	// Unique and unambiguous identifier of the intermediary.
	Identification *PartyIdentification111 `xml:"Id"`

	// Function performed by the intermediary (investment funds).
	Role *Role7Choice `xml:"Role,omitempty"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account21 `xml:"Acct,omitempty"`
}

func (i *Intermediary37) AddIdentification() *PartyIdentification111 {
	i.Identification = new(PartyIdentification111)
	return i.Identification
}

func (i *Intermediary37) AddRole() *Role7Choice {
	i.Role = new(Role7Choice)
	return i.Role
}

func (i *Intermediary37) AddAccount() *Account21 {
	i.Account = new(Account21)
	return i.Account
}

// Party that provides services to investors relating to financial products.
type Intermediary39 struct {

	// Unique and unambiguous identifier of the intermediary.
	Identification *PartyIdentification113 `xml:"Id"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account22 `xml:"Acct,omitempty"`

	// Counterparties eligibility as defined by article 24 of the EU MiFID Directive applicable to transactions executed by investment firms for eligible counterparties.
	OrderOriginatorEligibility *OrderOriginatorEligibility1Code `xml:"OrdrOrgtrElgblty,omitempty"`

	// Capacity of the party executing an order.
	TradingPartyCapacity *TradingCapacity8Code `xml:"TradgPtyCpcty,omitempty"`

	// Function performed by the intermediary.
	Role *InvestmentFundRole2Choice `xml:"Role,omitempty"`
}

func (i *Intermediary39) AddIdentification() *PartyIdentification113 {
	i.Identification = new(PartyIdentification113)
	return i.Identification
}

func (i *Intermediary39) AddAccount() *Account22 {
	i.Account = new(Account22)
	return i.Account
}

func (i *Intermediary39) SetOrderOriginatorEligibility(value string) {
	i.OrderOriginatorEligibility = (*OrderOriginatorEligibility1Code)(&value)
}

func (i *Intermediary39) SetTradingPartyCapacity(value string) {
	i.TradingPartyCapacity = (*TradingCapacity8Code)(&value)
}

func (i *Intermediary39) AddRole() *InvestmentFundRole2Choice {
	i.Role = new(InvestmentFundRole2Choice)
	return i.Role
}

// Party that provides services to investors relating to financial products.
type Intermediary4 struct {

	// Unique and unambiguous identifier of the intermediary.
	Identification *PartyIdentification2Choice `xml:"Id"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account2 `xml:"Acct,omitempty"`

	// Function performed by the intermediary.
	Role *Max35Text `xml:"Role,omitempty"`
}

func (i *Intermediary4) AddIdentification() *PartyIdentification2Choice {
	i.Identification = new(PartyIdentification2Choice)
	return i.Identification
}

func (i *Intermediary4) AddAccount() *Account2 {
	i.Account = new(Account2)
	return i.Account
}

func (i *Intermediary4) SetRole(value string) {
	i.Role = (*Max35Text)(&value)
}

// Party that provides services to investors relating to financial products.
type Intermediary40 struct {

	// Unique and unambiguous identifier of the intermediary.
	Identification *PartyIdentification113 `xml:"Id"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account22 `xml:"Acct,omitempty"`

	// Counterparties eligibility as defined by article 24 of the EU MiFID Directive applicable to transactions executed by investment firms for eligible counterparties.
	OrderOriginatorEligibility *OrderOriginatorEligibility1Code `xml:"OrdrOrgtrElgblty,omitempty"`

	// Function performed by the intermediary.
	Role *InvestmentFundRole2Choice `xml:"Role,omitempty"`
}

func (i *Intermediary40) AddIdentification() *PartyIdentification113 {
	i.Identification = new(PartyIdentification113)
	return i.Identification
}

func (i *Intermediary40) AddAccount() *Account22 {
	i.Account = new(Account22)
	return i.Account
}

func (i *Intermediary40) SetOrderOriginatorEligibility(value string) {
	i.OrderOriginatorEligibility = (*OrderOriginatorEligibility1Code)(&value)
}

func (i *Intermediary40) AddRole() *InvestmentFundRole2Choice {
	i.Role = new(InvestmentFundRole2Choice)
	return i.Role
}

// Party that provides services relating to financial products to investors, eg, advice on products and placement of orders for the investment fund.
type Intermediary7 struct {

	// Unique and unambiguous identifier for an organisation that is allocated by an institution, eg, Dun & Bradstreet Identification.
	Identification *PartyIdentification2Choice `xml:"Id"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account2 `xml:"Acct,omitempty"`
}

func (i *Intermediary7) AddIdentification() *PartyIdentification2Choice {
	i.Identification = new(PartyIdentification2Choice)
	return i.Identification
}

func (i *Intermediary7) AddAccount() *Account2 {
	i.Account = new(Account2)
	return i.Account
}

// Party that provides services to investors relating to financial products.
type Intermediary8 struct {

	// Unique and unambiguous identifier of the intermediary.
	Identification *PartyIdentification2Choice `xml:"Id"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account7 `xml:"Acct,omitempty"`

	// Counterparties eligibility as defined by article 24 of the EU MiFID Directive applicable to transactions executed by investment firms for eligible counterparties.
	OrderOriginatorEligibility *OrderOriginatorEligibility1Code `xml:"OrdrOrgtrElgblty,omitempty"`

	// Function performed by the intermediary.
	Role *InvestmentFundRole2Code `xml:"Role,omitempty"`

	// Function performed by the intermediary.
	ExtendedRole *Extended350Code `xml:"XtndedRole,omitempty"`
}

func (i *Intermediary8) AddIdentification() *PartyIdentification2Choice {
	i.Identification = new(PartyIdentification2Choice)
	return i.Identification
}

func (i *Intermediary8) AddAccount() *Account7 {
	i.Account = new(Account7)
	return i.Account
}

func (i *Intermediary8) SetOrderOriginatorEligibility(value string) {
	i.OrderOriginatorEligibility = (*OrderOriginatorEligibility1Code)(&value)
}

func (i *Intermediary8) SetRole(value string) {
	i.Role = (*InvestmentFundRole2Code)(&value)
}

func (i *Intermediary8) SetExtendedRole(value string) {
	i.ExtendedRole = (*Extended350Code)(&value)
}

// Party that provides services to investors relating to financial products.
type Intermediary9 struct {

	// Unique and unambiguous identifier of the intermediary.
	Identification *PartyIdentification2Choice `xml:"Id"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account7 `xml:"Acct,omitempty"`

	// Counterparties eligibility as defined by article 24 of the EU MiFID Directive applicable to transactions executed by investment firms for eligible counterparties.
	OrderOriginatorEligibility *OrderOriginatorEligibility1Code `xml:"OrdrOrgtrElgblty,omitempty"`

	// Capacity of the party executing an order.
	TradingPartyCapacity *TradingCapacity2Code `xml:"TradgPtyCpcty,omitempty"`

	// Function performed by the intermediary.
	Role *InvestmentFundRole2Code `xml:"Role,omitempty"`

	// Function performed by the intermediary.
	ExtendedRole *Extended350Code `xml:"XtndedRole,omitempty"`
}

func (i *Intermediary9) AddIdentification() *PartyIdentification2Choice {
	i.Identification = new(PartyIdentification2Choice)
	return i.Identification
}

func (i *Intermediary9) AddAccount() *Account7 {
	i.Account = new(Account7)
	return i.Account
}

func (i *Intermediary9) SetOrderOriginatorEligibility(value string) {
	i.OrderOriginatorEligibility = (*OrderOriginatorEligibility1Code)(&value)
}

func (i *Intermediary9) SetTradingPartyCapacity(value string) {
	i.TradingPartyCapacity = (*TradingCapacity2Code)(&value)
}

func (i *Intermediary9) SetRole(value string) {
	i.Role = (*InvestmentFundRole2Code)(&value)
}

func (i *Intermediary9) SetExtendedRole(value string) {
	i.ExtendedRole = (*Extended350Code)(&value)
}
