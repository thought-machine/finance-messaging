package iso20022

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type SettlementParties10 struct {

	// First party in the settlement chain. In a plain vanilla settlement, it is the Central Securities Depository where the counterparty requests to receive the financial instrument or from where the counterparty delivers the financial instruments.
	Depository *PartyIdentification48 `xml:"Dpstry,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	Party1 *PartyIdentificationAndAccount35 `xml:"Pty1,omitempty"`

	// Party that, in a settlement chain interacts with the party 1.
	Party2 *PartyIdentificationAndAccount35 `xml:"Pty2,omitempty"`

	// Party that, in a settlement chain interacts with the party 2.
	Party3 *PartyIdentificationAndAccount35 `xml:"Pty3,omitempty"`

	// Party that, in a settlement chain interacts with the party 3.
	Party4 *PartyIdentificationAndAccount35 `xml:"Pty4,omitempty"`

	// Party that, in a settlement chain interacts with the party 4.
	Party5 *PartyIdentificationAndAccount35 `xml:"Pty5,omitempty"`
}

func (s *SettlementParties10) AddDepository() *PartyIdentification48 {
	s.Depository = new(PartyIdentification48)
	return s.Depository
}

func (s *SettlementParties10) AddParty1() *PartyIdentificationAndAccount35 {
	s.Party1 = new(PartyIdentificationAndAccount35)
	return s.Party1
}

func (s *SettlementParties10) AddParty2() *PartyIdentificationAndAccount35 {
	s.Party2 = new(PartyIdentificationAndAccount35)
	return s.Party2
}

func (s *SettlementParties10) AddParty3() *PartyIdentificationAndAccount35 {
	s.Party3 = new(PartyIdentificationAndAccount35)
	return s.Party3
}

func (s *SettlementParties10) AddParty4() *PartyIdentificationAndAccount35 {
	s.Party4 = new(PartyIdentificationAndAccount35)
	return s.Party4
}

func (s *SettlementParties10) AddParty5() *PartyIdentificationAndAccount35 {
	s.Party5 = new(PartyIdentificationAndAccount35)
	return s.Party5
}

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type SettlementParties11 struct {

	// First party in the settlement chain. In a plain vanilla settlement, it is the Central Securities Depository where the counterparty requests to receive the financial instrument or from where the counterparty delivers the financial instruments.
	Depository *PartyIdentification48 `xml:"Dpstry,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	Party1 *PartyIdentificationAndAccount42 `xml:"Pty1,omitempty"`

	// Party that, in a settlement chain interacts with the party 1.
	Party2 *PartyIdentificationAndAccount42 `xml:"Pty2,omitempty"`

	// Party that, in a settlement chain interacts with the party 2.
	Party3 *PartyIdentificationAndAccount42 `xml:"Pty3,omitempty"`

	// Party that, in a settlement chain interacts with the party 3.
	Party4 *PartyIdentificationAndAccount42 `xml:"Pty4,omitempty"`

	// Party that, in a settlement chain interacts with the party 4.
	Party5 *PartyIdentificationAndAccount42 `xml:"Pty5,omitempty"`
}

func (s *SettlementParties11) AddDepository() *PartyIdentification48 {
	s.Depository = new(PartyIdentification48)
	return s.Depository
}

func (s *SettlementParties11) AddParty1() *PartyIdentificationAndAccount42 {
	s.Party1 = new(PartyIdentificationAndAccount42)
	return s.Party1
}

func (s *SettlementParties11) AddParty2() *PartyIdentificationAndAccount42 {
	s.Party2 = new(PartyIdentificationAndAccount42)
	return s.Party2
}

func (s *SettlementParties11) AddParty3() *PartyIdentificationAndAccount42 {
	s.Party3 = new(PartyIdentificationAndAccount42)
	return s.Party3
}

func (s *SettlementParties11) AddParty4() *PartyIdentificationAndAccount42 {
	s.Party4 = new(PartyIdentificationAndAccount42)
	return s.Party4
}

func (s *SettlementParties11) AddParty5() *PartyIdentificationAndAccount42 {
	s.Party5 = new(PartyIdentificationAndAccount42)
	return s.Party5
}

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type SettlementParties12 struct {

	// First party in the settlement chain. In a plain vanilla settlement, it is the Central Securities Depository where the counterparty requests to receive the financial instrument or from where the counterparty delivers the financial instruments.
	Depository *PartyIdentification46 `xml:"Dpstry,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	Party1 *PartyIdentificationAndAccount43 `xml:"Pty1,omitempty"`

	// Party that, in a settlement chain interacts with the party 1.
	Party2 *PartyIdentificationAndAccount43 `xml:"Pty2,omitempty"`

	// Party that, in a settlement chain interacts with the party 2.
	Party3 *PartyIdentificationAndAccount43 `xml:"Pty3,omitempty"`

	// Party that, in a settlement chain interacts with the party 3.
	Party4 *PartyIdentificationAndAccount43 `xml:"Pty4,omitempty"`

	// Party that, in a settlement chain interacts with the party 4.
	Party5 *PartyIdentificationAndAccount43 `xml:"Pty5,omitempty"`
}

func (s *SettlementParties12) AddDepository() *PartyIdentification46 {
	s.Depository = new(PartyIdentification46)
	return s.Depository
}

func (s *SettlementParties12) AddParty1() *PartyIdentificationAndAccount43 {
	s.Party1 = new(PartyIdentificationAndAccount43)
	return s.Party1
}

func (s *SettlementParties12) AddParty2() *PartyIdentificationAndAccount43 {
	s.Party2 = new(PartyIdentificationAndAccount43)
	return s.Party2
}

func (s *SettlementParties12) AddParty3() *PartyIdentificationAndAccount43 {
	s.Party3 = new(PartyIdentificationAndAccount43)
	return s.Party3
}

func (s *SettlementParties12) AddParty4() *PartyIdentificationAndAccount43 {
	s.Party4 = new(PartyIdentificationAndAccount43)
	return s.Party4
}

func (s *SettlementParties12) AddParty5() *PartyIdentificationAndAccount43 {
	s.Party5 = new(PartyIdentificationAndAccount43)
	return s.Party5
}

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type SettlementParties13 struct {

	// First party in the settlement chain. In a plain vanilla settlement, it is the Central Securities Depository where the counterparty requests to receive the financial instrument or from where the counterparty delivers the financial instruments.
	Depository *PartyIdentification46 `xml:"Dpstry,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	Party1 *PartyIdentificationAndAccount44 `xml:"Pty1,omitempty"`

	// Party that, in a settlement chain interacts with the party 1.
	Party2 *PartyIdentificationAndAccount44 `xml:"Pty2,omitempty"`

	// Party that, in a settlement chain interacts with the party 2.
	Party3 *PartyIdentificationAndAccount44 `xml:"Pty3,omitempty"`

	// Party that, in a settlement chain interacts with the party 3.
	Party4 *PartyIdentificationAndAccount44 `xml:"Pty4,omitempty"`

	// Party that, in a settlement chain interacts with the party 4.
	Party5 *PartyIdentificationAndAccount44 `xml:"Pty5,omitempty"`
}

func (s *SettlementParties13) AddDepository() *PartyIdentification46 {
	s.Depository = new(PartyIdentification46)
	return s.Depository
}

func (s *SettlementParties13) AddParty1() *PartyIdentificationAndAccount44 {
	s.Party1 = new(PartyIdentificationAndAccount44)
	return s.Party1
}

func (s *SettlementParties13) AddParty2() *PartyIdentificationAndAccount44 {
	s.Party2 = new(PartyIdentificationAndAccount44)
	return s.Party2
}

func (s *SettlementParties13) AddParty3() *PartyIdentificationAndAccount44 {
	s.Party3 = new(PartyIdentificationAndAccount44)
	return s.Party3
}

func (s *SettlementParties13) AddParty4() *PartyIdentificationAndAccount44 {
	s.Party4 = new(PartyIdentificationAndAccount44)
	return s.Party4
}

func (s *SettlementParties13) AddParty5() *PartyIdentificationAndAccount44 {
	s.Party5 = new(PartyIdentificationAndAccount44)
	return s.Party5
}

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type SettlementParties14 struct {

	// First party in the settlement chain. In a plain vanilla settlement, it is the Central Securities Depository where the counterparty requests to receive the financial instrument or from where the counterparty delivers the financial instruments.
	Depository *PartyIdentification48 `xml:"Dpstry,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	Party1 *PartyIdentificationAndAccount45 `xml:"Pty1,omitempty"`

	// Party that, in a settlement chain interacts with the party 1.
	Party2 *PartyIdentificationAndAccount45 `xml:"Pty2,omitempty"`

	// Party that, in a settlement chain interacts with the party 2.
	Party3 *PartyIdentificationAndAccount45 `xml:"Pty3,omitempty"`

	// Party that, in a settlement chain interacts with the party 3.
	Party4 *PartyIdentificationAndAccount45 `xml:"Pty4,omitempty"`

	// Party that, in a settlement chain interacts with the party 4.
	Party5 *PartyIdentificationAndAccount45 `xml:"Pty5,omitempty"`
}

func (s *SettlementParties14) AddDepository() *PartyIdentification48 {
	s.Depository = new(PartyIdentification48)
	return s.Depository
}

func (s *SettlementParties14) AddParty1() *PartyIdentificationAndAccount45 {
	s.Party1 = new(PartyIdentificationAndAccount45)
	return s.Party1
}

func (s *SettlementParties14) AddParty2() *PartyIdentificationAndAccount45 {
	s.Party2 = new(PartyIdentificationAndAccount45)
	return s.Party2
}

func (s *SettlementParties14) AddParty3() *PartyIdentificationAndAccount45 {
	s.Party3 = new(PartyIdentificationAndAccount45)
	return s.Party3
}

func (s *SettlementParties14) AddParty4() *PartyIdentificationAndAccount45 {
	s.Party4 = new(PartyIdentificationAndAccount45)
	return s.Party4
}

func (s *SettlementParties14) AddParty5() *PartyIdentificationAndAccount45 {
	s.Party5 = new(PartyIdentificationAndAccount45)
	return s.Party5
}

// Specifies settlement parties (delivering/receiving).
type SettlementParties15 struct {

	// First receiving party in the settlement chain. In a plain vanilla settlement, it is the central securities depository where the receiving side of the transaction requests to receive the financial instrument.
	Depository *PartyIdentification47 `xml:"Dpstry"`

	// Party that interacts with the depository.
	Party1 *PartyIdentificationAndAccount51 `xml:"Pty1"`

	// Party that interacts with the party1.
	Party2 *PartyIdentificationAndAccount51 `xml:"Pty2,omitempty"`

	// Party that interacts with the party2.
	Party3 *PartyIdentificationAndAccount51 `xml:"Pty3,omitempty"`
}

func (s *SettlementParties15) AddDepository() *PartyIdentification47 {
	s.Depository = new(PartyIdentification47)
	return s.Depository
}

func (s *SettlementParties15) AddParty1() *PartyIdentificationAndAccount51 {
	s.Party1 = new(PartyIdentificationAndAccount51)
	return s.Party1
}

func (s *SettlementParties15) AddParty2() *PartyIdentificationAndAccount51 {
	s.Party2 = new(PartyIdentificationAndAccount51)
	return s.Party2
}

func (s *SettlementParties15) AddParty3() *PartyIdentificationAndAccount51 {
	s.Party3 = new(PartyIdentificationAndAccount51)
	return s.Party3
}

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type SettlementParties2 struct {

	// First party in the settlement chain. In a plain vanilla settlement, it is the Central Securities Depository where the counterparty requests to receive the financial instrument or from where the counterparty delivers the financial instruments.
	Depository *PartyIdentification36 `xml:"Dpstry,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	Party1 *PartyIdentificationAndAccount16 `xml:"Pty1,omitempty"`

	// Party that, in a settlement chain interacts with the party 1.
	Party2 *PartyIdentificationAndAccount16 `xml:"Pty2,omitempty"`

	// Party that, in a settlement chain interacts with the party 2.
	Party3 *PartyIdentificationAndAccount16 `xml:"Pty3,omitempty"`

	// Party that, in a settlement chain interacts with the party 3.
	Party4 *PartyIdentificationAndAccount16 `xml:"Pty4,omitempty"`

	// Party that, in a settlement chain interacts with the party 4.
	Party5 *PartyIdentificationAndAccount16 `xml:"Pty5,omitempty"`
}

func (s *SettlementParties2) AddDepository() *PartyIdentification36 {
	s.Depository = new(PartyIdentification36)
	return s.Depository
}

func (s *SettlementParties2) AddParty1() *PartyIdentificationAndAccount16 {
	s.Party1 = new(PartyIdentificationAndAccount16)
	return s.Party1
}

func (s *SettlementParties2) AddParty2() *PartyIdentificationAndAccount16 {
	s.Party2 = new(PartyIdentificationAndAccount16)
	return s.Party2
}

func (s *SettlementParties2) AddParty3() *PartyIdentificationAndAccount16 {
	s.Party3 = new(PartyIdentificationAndAccount16)
	return s.Party3
}

func (s *SettlementParties2) AddParty4() *PartyIdentificationAndAccount16 {
	s.Party4 = new(PartyIdentificationAndAccount16)
	return s.Party4
}

func (s *SettlementParties2) AddParty5() *PartyIdentificationAndAccount16 {
	s.Party5 = new(PartyIdentificationAndAccount16)
	return s.Party5
}

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type SettlementParties23 struct {

	// First party in the settlement chain. In a plain vanilla settlement, it is the Central Securities Depository where the counterparty requests to receive the financial instrument or from where the counterparty delivers the financial instruments.
	Depository *PartyIdentification55 `xml:"Dpstry,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	Party1 *PartyIdentificationAndAccount34 `xml:"Pty1,omitempty"`

	// Party that, in a settlement chain interacts with the party 1.
	Party2 *PartyIdentificationAndAccount34 `xml:"Pty2,omitempty"`

	// Party that, in a settlement chain interacts with the party 2.
	Party3 *PartyIdentificationAndAccount34 `xml:"Pty3,omitempty"`

	// Party that, in a settlement chain interacts with the party 3.
	Party4 *PartyIdentificationAndAccount34 `xml:"Pty4,omitempty"`

	// Party that, in a settlement chain interacts with the party 4.
	Party5 *PartyIdentificationAndAccount34 `xml:"Pty5,omitempty"`
}

func (s *SettlementParties23) AddDepository() *PartyIdentification55 {
	s.Depository = new(PartyIdentification55)
	return s.Depository
}

func (s *SettlementParties23) AddParty1() *PartyIdentificationAndAccount34 {
	s.Party1 = new(PartyIdentificationAndAccount34)
	return s.Party1
}

func (s *SettlementParties23) AddParty2() *PartyIdentificationAndAccount34 {
	s.Party2 = new(PartyIdentificationAndAccount34)
	return s.Party2
}

func (s *SettlementParties23) AddParty3() *PartyIdentificationAndAccount34 {
	s.Party3 = new(PartyIdentificationAndAccount34)
	return s.Party3
}

func (s *SettlementParties23) AddParty4() *PartyIdentificationAndAccount34 {
	s.Party4 = new(PartyIdentificationAndAccount34)
	return s.Party4
}

func (s *SettlementParties23) AddParty5() *PartyIdentificationAndAccount34 {
	s.Party5 = new(PartyIdentificationAndAccount34)
	return s.Party5
}

// Specifies settlement parties (delivering/receiving).
type SettlementParties24 struct {

	// First receiving party in the settlement chain. In a plain vanilla settlement, it is the central securities depository where the receiving side of the transaction requests to receive the financial instrument.
	Depository *PartyIdentification47 `xml:"Dpstry,omitempty"`

	// Party that interacts with the depository.
	Party1 *PartyIdentificationAndAccount51 `xml:"Pty1,omitempty"`

	// Party that interacts with the party1.
	Party2 *PartyIdentificationAndAccount51 `xml:"Pty2,omitempty"`

	// Party that interacts with the party2.
	Party3 *PartyIdentificationAndAccount51 `xml:"Pty3,omitempty"`
}

func (s *SettlementParties24) AddDepository() *PartyIdentification47 {
	s.Depository = new(PartyIdentification47)
	return s.Depository
}

func (s *SettlementParties24) AddParty1() *PartyIdentificationAndAccount51 {
	s.Party1 = new(PartyIdentificationAndAccount51)
	return s.Party1
}

func (s *SettlementParties24) AddParty2() *PartyIdentificationAndAccount51 {
	s.Party2 = new(PartyIdentificationAndAccount51)
	return s.Party2
}

func (s *SettlementParties24) AddParty3() *PartyIdentificationAndAccount51 {
	s.Party3 = new(PartyIdentificationAndAccount51)
	return s.Party3
}

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type SettlementParties26 struct {

	// Party that, in a settlement chain interacts with the party 1.
	Party2 *PartyIdentificationAndAccount42 `xml:"Pty2,omitempty"`

	// Party that, in a settlement chain interacts with the party 2.
	Party3 *PartyIdentificationAndAccount42 `xml:"Pty3,omitempty"`

	// Party that, in a settlement chain interacts with the party 3.
	Party4 *PartyIdentificationAndAccount42 `xml:"Pty4,omitempty"`

	// Party that, in a settlement chain interacts with the party 4.
	Party5 *PartyIdentificationAndAccount42 `xml:"Pty5,omitempty"`
}

func (s *SettlementParties26) AddParty2() *PartyIdentificationAndAccount42 {
	s.Party2 = new(PartyIdentificationAndAccount42)
	return s.Party2
}

func (s *SettlementParties26) AddParty3() *PartyIdentificationAndAccount42 {
	s.Party3 = new(PartyIdentificationAndAccount42)
	return s.Party3
}

func (s *SettlementParties26) AddParty4() *PartyIdentificationAndAccount42 {
	s.Party4 = new(PartyIdentificationAndAccount42)
	return s.Party4
}

func (s *SettlementParties26) AddParty5() *PartyIdentificationAndAccount42 {
	s.Party5 = new(PartyIdentificationAndAccount42)
	return s.Party5
}

// Identification of a settlement party by a choice between a BIC or a name and address or a party identification.
type SettlementParties29 struct {

	// Financial institution from which cash will be transferred.
	DeliveryAgent *PartyIdentification73Choice `xml:"DlvryAgt,omitempty"`

	// Party, within the settlement chain, between the delivery and receiving agents.
	Intermediary *PartyIdentification73Choice `xml:"Intrmy,omitempty"`

	// Financial institution where the payee will receive the funds.
	ReceivingAgent *PartyIdentification73Choice `xml:"RcvgAgt"`

	// Ultimate institution that will receive the funds when different from the trading or counterparty side.
	BeneficiaryInstitution *PartyIdentification73Choice `xml:"BnfcryInstn,omitempty"`
}

func (s *SettlementParties29) AddDeliveryAgent() *PartyIdentification73Choice {
	s.DeliveryAgent = new(PartyIdentification73Choice)
	return s.DeliveryAgent
}

func (s *SettlementParties29) AddIntermediary() *PartyIdentification73Choice {
	s.Intermediary = new(PartyIdentification73Choice)
	return s.Intermediary
}

func (s *SettlementParties29) AddReceivingAgent() *PartyIdentification73Choice {
	s.ReceivingAgent = new(PartyIdentification73Choice)
	return s.ReceivingAgent
}

func (s *SettlementParties29) AddBeneficiaryInstitution() *PartyIdentification73Choice {
	s.BeneficiaryInstitution = new(PartyIdentification73Choice)
	return s.BeneficiaryInstitution
}

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type SettlementParties32 struct {

	// First party in the settlement chain. In a plain vanilla settlement, it is the central securities depository where the counterparty requests to receive the financial instrument or from where the counterparty delivers the financial instruments.
	Depository *PartyIdentification63 `xml:"Dpstry"`

	// Party that, in a settlement chain, interacts with the depository. This may also be known as the “local agent”, “sub-custodian”, “receiving agent” or “delivering agent”.
	Party1 *PartyIdentificationAndAccount95 `xml:"Pty1,omitempty"`

	// Party that, in a settlement chain, interacts with party 1. This may also be known as the “investment manager” or “custodian”.
	Party2 *PartyIdentificationAndAccount95 `xml:"Pty2,omitempty"`

	// Party that, in a settlement chain, interacts with party 2.
	Party3 *PartyIdentificationAndAccount95 `xml:"Pty3,omitempty"`

	// Party that, in a settlement chain, interacts with party 3.
	Party4 *PartyIdentificationAndAccount95 `xml:"Pty4,omitempty"`

	// Party that, in a settlement chain, interacts with party 4.
	Party5 *PartyIdentificationAndAccount95 `xml:"Pty5,omitempty"`
}

func (s *SettlementParties32) AddDepository() *PartyIdentification63 {
	s.Depository = new(PartyIdentification63)
	return s.Depository
}

func (s *SettlementParties32) AddParty1() *PartyIdentificationAndAccount95 {
	s.Party1 = new(PartyIdentificationAndAccount95)
	return s.Party1
}

func (s *SettlementParties32) AddParty2() *PartyIdentificationAndAccount95 {
	s.Party2 = new(PartyIdentificationAndAccount95)
	return s.Party2
}

func (s *SettlementParties32) AddParty3() *PartyIdentificationAndAccount95 {
	s.Party3 = new(PartyIdentificationAndAccount95)
	return s.Party3
}

func (s *SettlementParties32) AddParty4() *PartyIdentificationAndAccount95 {
	s.Party4 = new(PartyIdentificationAndAccount95)
	return s.Party4
}

func (s *SettlementParties32) AddParty5() *PartyIdentificationAndAccount95 {
	s.Party5 = new(PartyIdentificationAndAccount95)
	return s.Party5
}

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type SettlementParties35 struct {

	// Parties through which settlement is to take place.
	StandingSettlementParties *SettlementParties32 `xml:"StgSttlmPties"`

	// Identifier needed for settlement purposes. This identifier could be, for example, an identifier that identifies an institution or agent at a CDS or ICSD (Depository Trust Clearing Corporation (DTC) Institution ID or DTC Agent ID). It could also be a local tax identification number or an ‘investor identification’, as mandated by local market practice.
	LocalMarketIdentification []*GenericIdentification49 `xml:"LclMktId,omitempty"`

	// Registration information required for settlement. For some markets, for example, Spain (Iberclear) registration details are mandatory and should be part of the SSI. In some cases, the name of the institution is different than what's provided in the BIC Directory. If this is the case, the name should be provided.
	RegistrationDetails *PartyIdentification99Choice `xml:"RegnDtls,omitempty"`
}

func (s *SettlementParties35) AddStandingSettlementParties() *SettlementParties32 {
	s.StandingSettlementParties = new(SettlementParties32)
	return s.StandingSettlementParties
}

func (s *SettlementParties35) AddLocalMarketIdentification() *GenericIdentification49 {
	newValue := new(GenericIdentification49)
	s.LocalMarketIdentification = append(s.LocalMarketIdentification, newValue)
	return newValue
}

func (s *SettlementParties35) AddRegistrationDetails() *PartyIdentification99Choice {
	s.RegistrationDetails = new(PartyIdentification99Choice)
	return s.RegistrationDetails
}

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type SettlementParties36 struct {

	// First party in the settlement chain. In a plain vanilla settlement, it is the Central Securities Depository where the counterparty requests to receive the financial instrument or from where the counterparty delivers the financial instruments.
	Depository *PartyIdentification75 `xml:"Dpstry,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	Party1 *PartyIdentificationAndAccount106 `xml:"Pty1,omitempty"`

	// Party that, in a settlement chain interacts with the party 1.
	Party2 *PartyIdentificationAndAccount106 `xml:"Pty2,omitempty"`

	// Party that, in a settlement chain interacts with the party 2.
	Party3 *PartyIdentificationAndAccount106 `xml:"Pty3,omitempty"`

	// Party that, in a settlement chain interacts with the party 3.
	Party4 *PartyIdentificationAndAccount106 `xml:"Pty4,omitempty"`

	// Party that, in a settlement chain interacts with the party 4.
	Party5 *PartyIdentificationAndAccount106 `xml:"Pty5,omitempty"`
}

func (s *SettlementParties36) AddDepository() *PartyIdentification75 {
	s.Depository = new(PartyIdentification75)
	return s.Depository
}

func (s *SettlementParties36) AddParty1() *PartyIdentificationAndAccount106 {
	s.Party1 = new(PartyIdentificationAndAccount106)
	return s.Party1
}

func (s *SettlementParties36) AddParty2() *PartyIdentificationAndAccount106 {
	s.Party2 = new(PartyIdentificationAndAccount106)
	return s.Party2
}

func (s *SettlementParties36) AddParty3() *PartyIdentificationAndAccount106 {
	s.Party3 = new(PartyIdentificationAndAccount106)
	return s.Party3
}

func (s *SettlementParties36) AddParty4() *PartyIdentificationAndAccount106 {
	s.Party4 = new(PartyIdentificationAndAccount106)
	return s.Party4
}

func (s *SettlementParties36) AddParty5() *PartyIdentificationAndAccount106 {
	s.Party5 = new(PartyIdentificationAndAccount106)
	return s.Party5
}

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type SettlementParties39 struct {

	// Party that, in a settlement chain interacts with the party 1.
	Party2 *PartyIdentificationAndAccount106 `xml:"Pty2,omitempty"`

	// Party that, in a settlement chain interacts with the party 2.
	Party3 *PartyIdentificationAndAccount106 `xml:"Pty3,omitempty"`

	// Party that, in a settlement chain interacts with the party 3.
	Party4 *PartyIdentificationAndAccount106 `xml:"Pty4,omitempty"`

	// Party that, in a settlement chain interacts with the party 4.
	Party5 *PartyIdentificationAndAccount106 `xml:"Pty5,omitempty"`
}

func (s *SettlementParties39) AddParty2() *PartyIdentificationAndAccount106 {
	s.Party2 = new(PartyIdentificationAndAccount106)
	return s.Party2
}

func (s *SettlementParties39) AddParty3() *PartyIdentificationAndAccount106 {
	s.Party3 = new(PartyIdentificationAndAccount106)
	return s.Party3
}

func (s *SettlementParties39) AddParty4() *PartyIdentificationAndAccount106 {
	s.Party4 = new(PartyIdentificationAndAccount106)
	return s.Party4
}

func (s *SettlementParties39) AddParty5() *PartyIdentificationAndAccount106 {
	s.Party5 = new(PartyIdentificationAndAccount106)
	return s.Party5
}

// Specifies settlement parties (delivering/receiving).
type SettlementParties4 struct {

	// First receiving party in the settlement chain. In a plain vanilla settlement, it is the Central Securities Depository where the receiving side of the transaction requests to receive the financial instrument.
	Depository *PartyIdentification35 `xml:"Dpstry"`

	// Party that interacts with the Depository.
	Party1 *PartyIdentificationAndAccount14 `xml:"Pty1"`

	// Party that interacts with the Party1.
	Party2 *PartyIdentificationAndAccount14 `xml:"Pty2,omitempty"`

	// Party that interacts with the Party2.
	Party3 *PartyIdentificationAndAccount14 `xml:"Pty3,omitempty"`
}

func (s *SettlementParties4) AddDepository() *PartyIdentification35 {
	s.Depository = new(PartyIdentification35)
	return s.Depository
}

func (s *SettlementParties4) AddParty1() *PartyIdentificationAndAccount14 {
	s.Party1 = new(PartyIdentificationAndAccount14)
	return s.Party1
}

func (s *SettlementParties4) AddParty2() *PartyIdentificationAndAccount14 {
	s.Party2 = new(PartyIdentificationAndAccount14)
	return s.Party2
}

func (s *SettlementParties4) AddParty3() *PartyIdentificationAndAccount14 {
	s.Party3 = new(PartyIdentificationAndAccount14)
	return s.Party3
}

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type SettlementParties40 struct {

	// First party in the settlement chain. In a plain vanilla settlement, it is the Central Securities Depository where the counterparty requests to receive the financial instrument or from where the counterparty delivers the financial instruments.
	Depository *PartyIdentification91 `xml:"Dpstry,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	Party1 *PartyIdentificationAndAccount117 `xml:"Pty1,omitempty"`

	// Party that, in a settlement chain interacts with the party 1.
	Party2 *PartyIdentificationAndAccount117 `xml:"Pty2,omitempty"`

	// Party that, in a settlement chain interacts with the party 2.
	Party3 *PartyIdentificationAndAccount117 `xml:"Pty3,omitempty"`

	// Party that, in a settlement chain interacts with the party 3.
	Party4 *PartyIdentificationAndAccount117 `xml:"Pty4,omitempty"`

	// Party that, in a settlement chain interacts with the party 4.
	Party5 *PartyIdentificationAndAccount117 `xml:"Pty5,omitempty"`
}

func (s *SettlementParties40) AddDepository() *PartyIdentification91 {
	s.Depository = new(PartyIdentification91)
	return s.Depository
}

func (s *SettlementParties40) AddParty1() *PartyIdentificationAndAccount117 {
	s.Party1 = new(PartyIdentificationAndAccount117)
	return s.Party1
}

func (s *SettlementParties40) AddParty2() *PartyIdentificationAndAccount117 {
	s.Party2 = new(PartyIdentificationAndAccount117)
	return s.Party2
}

func (s *SettlementParties40) AddParty3() *PartyIdentificationAndAccount117 {
	s.Party3 = new(PartyIdentificationAndAccount117)
	return s.Party3
}

func (s *SettlementParties40) AddParty4() *PartyIdentificationAndAccount117 {
	s.Party4 = new(PartyIdentificationAndAccount117)
	return s.Party4
}

func (s *SettlementParties40) AddParty5() *PartyIdentificationAndAccount117 {
	s.Party5 = new(PartyIdentificationAndAccount117)
	return s.Party5
}

// Specifies settlement parties (delivering/receiving).
type SettlementParties42 struct {

	// First receiving party in the settlement chain. In a plain vanilla settlement, it is the central securities depository where the receiving side of the transaction requests to receive the financial instrument.
	Depository *PartyIdentification92 `xml:"Dpstry,omitempty"`

	// Party that interacts with the depository.
	Party1 *PartyIdentificationAndAccount122 `xml:"Pty1,omitempty"`

	// Party that interacts with the party1.
	Party2 *PartyIdentificationAndAccount122 `xml:"Pty2,omitempty"`

	// Party that interacts with the party2.
	Party3 *PartyIdentificationAndAccount122 `xml:"Pty3,omitempty"`
}

func (s *SettlementParties42) AddDepository() *PartyIdentification92 {
	s.Depository = new(PartyIdentification92)
	return s.Depository
}

func (s *SettlementParties42) AddParty1() *PartyIdentificationAndAccount122 {
	s.Party1 = new(PartyIdentificationAndAccount122)
	return s.Party1
}

func (s *SettlementParties42) AddParty2() *PartyIdentificationAndAccount122 {
	s.Party2 = new(PartyIdentificationAndAccount122)
	return s.Party2
}

func (s *SettlementParties42) AddParty3() *PartyIdentificationAndAccount122 {
	s.Party3 = new(PartyIdentificationAndAccount122)
	return s.Party3
}

// Specifies settlement parties (delivering/receiving).
type SettlementParties43 struct {

	// First receiving party in the settlement chain. In a plain vanilla settlement, it is the central securities depository where the receiving side of the transaction requests to receive the financial instrument.
	Depository *PartyIdentification102 `xml:"Dpstry,omitempty"`

	// Party that interacts with the depository.
	Party1 *PartyIdentificationAndAccount128 `xml:"Pty1,omitempty"`

	// Party that interacts with the party1.
	Party2 *PartyIdentificationAndAccount128 `xml:"Pty2,omitempty"`

	// Party that interacts with the party2.
	Party3 *PartyIdentificationAndAccount128 `xml:"Pty3,omitempty"`
}

func (s *SettlementParties43) AddDepository() *PartyIdentification102 {
	s.Depository = new(PartyIdentification102)
	return s.Depository
}

func (s *SettlementParties43) AddParty1() *PartyIdentificationAndAccount128 {
	s.Party1 = new(PartyIdentificationAndAccount128)
	return s.Party1
}

func (s *SettlementParties43) AddParty2() *PartyIdentificationAndAccount128 {
	s.Party2 = new(PartyIdentificationAndAccount128)
	return s.Party2
}

func (s *SettlementParties43) AddParty3() *PartyIdentificationAndAccount128 {
	s.Party3 = new(PartyIdentificationAndAccount128)
	return s.Party3
}

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type SettlementParties44 struct {

	// First party in the settlement chain. In a plain vanilla settlement, it is the Central Securities Depository where the counterparty requests to receive the financial instrument or from where the counterparty delivers the financial instruments.
	Depository *PartyIdentification103 `xml:"Dpstry,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	Party1 *PartyIdentificationAndAccount131 `xml:"Pty1,omitempty"`

	// Party that, in a settlement chain interacts with the party 1.
	Party2 *PartyIdentificationAndAccount131 `xml:"Pty2,omitempty"`

	// Party that, in a settlement chain interacts with the party 2.
	Party3 *PartyIdentificationAndAccount131 `xml:"Pty3,omitempty"`

	// Party that, in a settlement chain interacts with the party 3.
	Party4 *PartyIdentificationAndAccount131 `xml:"Pty4,omitempty"`

	// Party that, in a settlement chain interacts with the party 4.
	Party5 *PartyIdentificationAndAccount131 `xml:"Pty5,omitempty"`
}

func (s *SettlementParties44) AddDepository() *PartyIdentification103 {
	s.Depository = new(PartyIdentification103)
	return s.Depository
}

func (s *SettlementParties44) AddParty1() *PartyIdentificationAndAccount131 {
	s.Party1 = new(PartyIdentificationAndAccount131)
	return s.Party1
}

func (s *SettlementParties44) AddParty2() *PartyIdentificationAndAccount131 {
	s.Party2 = new(PartyIdentificationAndAccount131)
	return s.Party2
}

func (s *SettlementParties44) AddParty3() *PartyIdentificationAndAccount131 {
	s.Party3 = new(PartyIdentificationAndAccount131)
	return s.Party3
}

func (s *SettlementParties44) AddParty4() *PartyIdentificationAndAccount131 {
	s.Party4 = new(PartyIdentificationAndAccount131)
	return s.Party4
}

func (s *SettlementParties44) AddParty5() *PartyIdentificationAndAccount131 {
	s.Party5 = new(PartyIdentificationAndAccount131)
	return s.Party5
}

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type SettlementParties49 struct {

	// First party in the settlement chain. In a plain vanilla settlement, it is the Central Securities Depository where the counterparty requests to receive the financial instrument or from where the counterparty delivers the financial instruments.
	Depository *PartyIdentification108 `xml:"Dpstry,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	Party1 *PartyIdentificationAndAccount146 `xml:"Pty1,omitempty"`

	// Party that, in a settlement chain interacts with the party 1.
	Party2 *PartyIdentificationAndAccount146 `xml:"Pty2,omitempty"`

	// Party that, in a settlement chain interacts with the party 2.
	Party3 *PartyIdentificationAndAccount146 `xml:"Pty3,omitempty"`

	// Party that, in a settlement chain interacts with the party 3.
	Party4 *PartyIdentificationAndAccount146 `xml:"Pty4,omitempty"`

	// Party that, in a settlement chain interacts with the party 4.
	Party5 *PartyIdentificationAndAccount146 `xml:"Pty5,omitempty"`
}

func (s *SettlementParties49) AddDepository() *PartyIdentification108 {
	s.Depository = new(PartyIdentification108)
	return s.Depository
}

func (s *SettlementParties49) AddParty1() *PartyIdentificationAndAccount146 {
	s.Party1 = new(PartyIdentificationAndAccount146)
	return s.Party1
}

func (s *SettlementParties49) AddParty2() *PartyIdentificationAndAccount146 {
	s.Party2 = new(PartyIdentificationAndAccount146)
	return s.Party2
}

func (s *SettlementParties49) AddParty3() *PartyIdentificationAndAccount146 {
	s.Party3 = new(PartyIdentificationAndAccount146)
	return s.Party3
}

func (s *SettlementParties49) AddParty4() *PartyIdentificationAndAccount146 {
	s.Party4 = new(PartyIdentificationAndAccount146)
	return s.Party4
}

func (s *SettlementParties49) AddParty5() *PartyIdentificationAndAccount146 {
	s.Party5 = new(PartyIdentificationAndAccount146)
	return s.Party5
}

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type SettlementParties5 struct {

	// First party in the settlement chain. In a plain vanilla settlement, it is the Central Securities Depository where the counterparty requests to receive the financial instrument or from where the counterparty delivers the financial instruments.
	Depository *PartyIdentification2 `xml:"Dpstry,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	Party1 *PartyIdentificationAndAccount1 `xml:"Pty1,omitempty"`

	// Party that, in a settlement chain interacts with the party 1.
	Party2 *PartyIdentificationAndAccount1 `xml:"Pty2,omitempty"`

	// Party that, in a settlement chain interacts with the party 2.
	Party3 *PartyIdentificationAndAccount1 `xml:"Pty3,omitempty"`

	// Party that, in a settlement chain interacts with the party 3.
	Party4 *PartyIdentificationAndAccount1 `xml:"Pty4,omitempty"`

	// Party that, in a settlement chain interacts with the party 4.
	Party5 *PartyIdentificationAndAccount1 `xml:"Pty5,omitempty"`
}

func (s *SettlementParties5) AddDepository() *PartyIdentification2 {
	s.Depository = new(PartyIdentification2)
	return s.Depository
}

func (s *SettlementParties5) AddParty1() *PartyIdentificationAndAccount1 {
	s.Party1 = new(PartyIdentificationAndAccount1)
	return s.Party1
}

func (s *SettlementParties5) AddParty2() *PartyIdentificationAndAccount1 {
	s.Party2 = new(PartyIdentificationAndAccount1)
	return s.Party2
}

func (s *SettlementParties5) AddParty3() *PartyIdentificationAndAccount1 {
	s.Party3 = new(PartyIdentificationAndAccount1)
	return s.Party3
}

func (s *SettlementParties5) AddParty4() *PartyIdentificationAndAccount1 {
	s.Party4 = new(PartyIdentificationAndAccount1)
	return s.Party4
}

func (s *SettlementParties5) AddParty5() *PartyIdentificationAndAccount1 {
	s.Party5 = new(PartyIdentificationAndAccount1)
	return s.Party5
}

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type SettlementParties58 struct {

	// Party that, in a settlement chain interacts with the party 1.
	Party2 *PartyIdentificationAndAccount131 `xml:"Pty2,omitempty"`

	// Party that, in a settlement chain interacts with the party 2.
	Party3 *PartyIdentificationAndAccount131 `xml:"Pty3,omitempty"`

	// Party that, in a settlement chain interacts with the party 3.
	Party4 *PartyIdentificationAndAccount131 `xml:"Pty4,omitempty"`

	// Party that, in a settlement chain interacts with the party 4.
	Party5 *PartyIdentificationAndAccount131 `xml:"Pty5,omitempty"`
}

func (s *SettlementParties58) AddParty2() *PartyIdentificationAndAccount131 {
	s.Party2 = new(PartyIdentificationAndAccount131)
	return s.Party2
}

func (s *SettlementParties58) AddParty3() *PartyIdentificationAndAccount131 {
	s.Party3 = new(PartyIdentificationAndAccount131)
	return s.Party3
}

func (s *SettlementParties58) AddParty4() *PartyIdentificationAndAccount131 {
	s.Party4 = new(PartyIdentificationAndAccount131)
	return s.Party4
}

func (s *SettlementParties58) AddParty5() *PartyIdentificationAndAccount131 {
	s.Party5 = new(PartyIdentificationAndAccount131)
	return s.Party5
}

// Specifies settlement parties (delivering/receiving).
type SettlementParties61 struct {

	// First receiving party in the settlement chain. In a plain vanilla settlement, it is the central securities depository where the receiving side of the transaction requests to receive the financial instrument.
	Depository *PartyIdentification120 `xml:"Dpstry,omitempty"`

	// Party that interacts with the depository.
	Party1 *PartyIdentificationAndAccount128 `xml:"Pty1,omitempty"`

	// Party that interacts with the party1.
	Party2 *PartyIdentificationAndAccount128 `xml:"Pty2,omitempty"`

	// Party that interacts with the party2.
	Party3 *PartyIdentificationAndAccount128 `xml:"Pty3,omitempty"`
}

func (s *SettlementParties61) AddDepository() *PartyIdentification120 {
	s.Depository = new(PartyIdentification120)
	return s.Depository
}

func (s *SettlementParties61) AddParty1() *PartyIdentificationAndAccount128 {
	s.Party1 = new(PartyIdentificationAndAccount128)
	return s.Party1
}

func (s *SettlementParties61) AddParty2() *PartyIdentificationAndAccount128 {
	s.Party2 = new(PartyIdentificationAndAccount128)
	return s.Party2
}

func (s *SettlementParties61) AddParty3() *PartyIdentificationAndAccount128 {
	s.Party3 = new(PartyIdentificationAndAccount128)
	return s.Party3
}
