package iso20022

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type ReceivingPartiesAndAccount1 struct {

	// Party that buys goods or services, or a financial instrument.
	ReceiverDetails *InvestmentAccount11 `xml:"RcvrDtls"`

	// Party that acts on behalf of the buyer of securities when the buyer does not have a direct relationship with the receiving agent.
	ReceiversCustodianDetails *PartyIdentificationAndAccount2 `xml:"RcvrsCtdnDtls,omitempty"`

	// Party that the Receiver's custodian uses to effect the receipt of a security, when the Receiver's custodian does not have a direct relationship with the Receiver agent.
	ReceiversIntermediaryDetails *PartyIdentificationAndAccount2 `xml:"RcvrsIntrmyDtls,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, eg, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount2 `xml:"RcvgAgtDtls"`

	// Identifies the securities settlement system to be used.
	SecuritiesSettlementSystem *Max35Text `xml:"SctiesSttlmSys,omitempty"`

	// Place where settlement of the securities takes place.
	PlaceOfSettlementDetails *PartyIdentificationAndAccount2 `xml:"PlcOfSttlmDtls"`
}

func (r *ReceivingPartiesAndAccount1) AddReceiverDetails() *InvestmentAccount11 {
	r.ReceiverDetails = new(InvestmentAccount11)
	return r.ReceiverDetails
}

func (r *ReceivingPartiesAndAccount1) AddReceiversCustodianDetails() *PartyIdentificationAndAccount2 {
	r.ReceiversCustodianDetails = new(PartyIdentificationAndAccount2)
	return r.ReceiversCustodianDetails
}

func (r *ReceivingPartiesAndAccount1) AddReceiversIntermediaryDetails() *PartyIdentificationAndAccount2 {
	r.ReceiversIntermediaryDetails = new(PartyIdentificationAndAccount2)
	return r.ReceiversIntermediaryDetails
}

func (r *ReceivingPartiesAndAccount1) AddReceivingAgentDetails() *PartyIdentificationAndAccount2 {
	r.ReceivingAgentDetails = new(PartyIdentificationAndAccount2)
	return r.ReceivingAgentDetails
}

func (r *ReceivingPartiesAndAccount1) SetSecuritiesSettlementSystem(value string) {
	r.SecuritiesSettlementSystem = (*Max35Text)(&value)
}

func (r *ReceivingPartiesAndAccount1) AddPlaceOfSettlementDetails() *PartyIdentificationAndAccount2 {
	r.PlaceOfSettlementDetails = new(PartyIdentificationAndAccount2)
	return r.PlaceOfSettlementDetails
}

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type ReceivingPartiesAndAccount10 struct {

	// First party in the settlement chain. In a plain vanilla settlement, it is the Central Securities Depository where the counterparty requests to receive the financial instrument or from where the counterparty delivers the financial instruments.
	Depository *PartyIdentification34Choice `xml:"Dpstry"`

	// Party that, in a settlement chain interacts with the depository.
	Party1 *PartyIdentificationAndAccount102 `xml:"Pty1"`

	// Party that, in a settlement chain interacts with the party 1.
	Party2 *PartyIdentificationAndAccount77 `xml:"Pty2,omitempty"`
}

func (r *ReceivingPartiesAndAccount10) AddDepository() *PartyIdentification34Choice {
	r.Depository = new(PartyIdentification34Choice)
	return r.Depository
}

func (r *ReceivingPartiesAndAccount10) AddParty1() *PartyIdentificationAndAccount102 {
	r.Party1 = new(PartyIdentificationAndAccount102)
	return r.Party1
}

func (r *ReceivingPartiesAndAccount10) AddParty2() *PartyIdentificationAndAccount77 {
	r.Party2 = new(PartyIdentificationAndAccount77)
	return r.Party2
}

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type ReceivingPartiesAndAccount11 struct {

	// First party in the settlement chain. In a plain vanilla settlement, it is the Central Securities Depository where the counterparty requests to receive the financial instrument or from where the counterparty delivers the financial instruments.
	Depository *PartyIdentification34Choice `xml:"Dpstry"`

	// Party that, in a settlement chain interacts with the depository.
	Party1 *PartyIdentificationAndAccount102 `xml:"Pty1"`

	// Party that, in a settlement chain interacts with the party 1.
	Party2 *PartyIdentificationAndAccount102 `xml:"Pty2,omitempty"`

	// Identifies the securities settlement system to be used.
	SecuritiesSettlementSystem *Max35Text `xml:"SctiesSttlmSys,omitempty"`
}

func (r *ReceivingPartiesAndAccount11) AddDepository() *PartyIdentification34Choice {
	r.Depository = new(PartyIdentification34Choice)
	return r.Depository
}

func (r *ReceivingPartiesAndAccount11) AddParty1() *PartyIdentificationAndAccount102 {
	r.Party1 = new(PartyIdentificationAndAccount102)
	return r.Party1
}

func (r *ReceivingPartiesAndAccount11) AddParty2() *PartyIdentificationAndAccount102 {
	r.Party2 = new(PartyIdentificationAndAccount102)
	return r.Party2
}

func (r *ReceivingPartiesAndAccount11) SetSecuritiesSettlementSystem(value string) {
	r.SecuritiesSettlementSystem = (*Max35Text)(&value)
}

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type ReceivingPartiesAndAccount13 struct {

	// Party that buys goods or services, or a financial instrument.
	ReceiverDetails *InvestmentAccount55 `xml:"RcvrDtls,omitempty"`

	// Party that acts on behalf of the buyer of securities when the buyer does not have a direct relationship with the receiving agent.
	ReceiversCustodianDetails *PartyIdentificationAndAccount124 `xml:"RcvrsCtdnDtls,omitempty"`

	// Party that the receiver's custodian uses to effect the receipt of a security, when the receiver's custodian does not have a direct relationship with the receiving agent.
	ReceiversIntermediary1Details *PartyIdentificationAndAccount124 `xml:"RcvrsIntrmy1Dtls,omitempty"`

	// Party that interacts with the receiver’s intermediary.
	ReceiversIntermediary2Details *PartyIdentificationAndAccount124 `xml:"RcvrsIntrmy2Dtls,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, for example, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount123 `xml:"RcvgAgtDtls"`

	// Identifies the securities settlement system to be used.
	SecuritiesSettlementSystem *Max35Text `xml:"SctiesSttlmSys,omitempty"`

	// Place where settlement of the securities takes place.
	PlaceOfSettlementDetails *PartyIdentification97 `xml:"PlcOfSttlmDtls,omitempty"`
}

func (r *ReceivingPartiesAndAccount13) AddReceiverDetails() *InvestmentAccount55 {
	r.ReceiverDetails = new(InvestmentAccount55)
	return r.ReceiverDetails
}

func (r *ReceivingPartiesAndAccount13) AddReceiversCustodianDetails() *PartyIdentificationAndAccount124 {
	r.ReceiversCustodianDetails = new(PartyIdentificationAndAccount124)
	return r.ReceiversCustodianDetails
}

func (r *ReceivingPartiesAndAccount13) AddReceiversIntermediary1Details() *PartyIdentificationAndAccount124 {
	r.ReceiversIntermediary1Details = new(PartyIdentificationAndAccount124)
	return r.ReceiversIntermediary1Details
}

func (r *ReceivingPartiesAndAccount13) AddReceiversIntermediary2Details() *PartyIdentificationAndAccount124 {
	r.ReceiversIntermediary2Details = new(PartyIdentificationAndAccount124)
	return r.ReceiversIntermediary2Details
}

func (r *ReceivingPartiesAndAccount13) AddReceivingAgentDetails() *PartyIdentificationAndAccount123 {
	r.ReceivingAgentDetails = new(PartyIdentificationAndAccount123)
	return r.ReceivingAgentDetails
}

func (r *ReceivingPartiesAndAccount13) SetSecuritiesSettlementSystem(value string) {
	r.SecuritiesSettlementSystem = (*Max35Text)(&value)
}

func (r *ReceivingPartiesAndAccount13) AddPlaceOfSettlementDetails() *PartyIdentification97 {
	r.PlaceOfSettlementDetails = new(PartyIdentification97)
	return r.PlaceOfSettlementDetails
}

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type ReceivingPartiesAndAccount14 struct {

	// Party that acts on behalf of the buyer of securities when the buyer does not have a direct relationship with the receiving agent.
	ReceiversCustodianDetails *PartyIdentificationAndAccount124 `xml:"RcvrsCtdnDtls,omitempty"`

	// Party that the receiver's custodian uses to effect the receipt of a security, when the receiver's custodian does not have a direct relationship with the receiving agent.
	ReceiversIntermediary1Details *PartyIdentificationAndAccount124 `xml:"RcvrsIntrmy1Dtls,omitempty"`

	// Party that interacts with the receiver’s intermediary.
	ReceiversIntermediary2Details *PartyIdentificationAndAccount124 `xml:"RcvrsIntrmy2Dtls,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, for example, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount123 `xml:"RcvgAgtDtls"`

	// Identifies the securities settlement system to be used.
	SecuritiesSettlementSystem *Max35Text `xml:"SctiesSttlmSys,omitempty"`

	// Place where settlement of the securities takes place.
	PlaceOfSettlementDetails *PartyIdentification97 `xml:"PlcOfSttlmDtls,omitempty"`
}

func (r *ReceivingPartiesAndAccount14) AddReceiversCustodianDetails() *PartyIdentificationAndAccount124 {
	r.ReceiversCustodianDetails = new(PartyIdentificationAndAccount124)
	return r.ReceiversCustodianDetails
}

func (r *ReceivingPartiesAndAccount14) AddReceiversIntermediary1Details() *PartyIdentificationAndAccount124 {
	r.ReceiversIntermediary1Details = new(PartyIdentificationAndAccount124)
	return r.ReceiversIntermediary1Details
}

func (r *ReceivingPartiesAndAccount14) AddReceiversIntermediary2Details() *PartyIdentificationAndAccount124 {
	r.ReceiversIntermediary2Details = new(PartyIdentificationAndAccount124)
	return r.ReceiversIntermediary2Details
}

func (r *ReceivingPartiesAndAccount14) AddReceivingAgentDetails() *PartyIdentificationAndAccount123 {
	r.ReceivingAgentDetails = new(PartyIdentificationAndAccount123)
	return r.ReceivingAgentDetails
}

func (r *ReceivingPartiesAndAccount14) SetSecuritiesSettlementSystem(value string) {
	r.SecuritiesSettlementSystem = (*Max35Text)(&value)
}

func (r *ReceivingPartiesAndAccount14) AddPlaceOfSettlementDetails() *PartyIdentification97 {
	r.PlaceOfSettlementDetails = new(PartyIdentification97)
	return r.PlaceOfSettlementDetails
}

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type ReceivingPartiesAndAccount15 struct {

	// First party in the settlement chain. In a plain vanilla settlement, it is the Central Securities Depository where the counterparty requests to receive the financial instrument or from where the counterparty delivers the financial instruments.
	Depository *PartyIdentification102Choice `xml:"Dpstry"`

	// Party that, in a settlement chain interacts with the depository.
	Party1 *PartyIdentificationAndAccount126 `xml:"Pty1"`

	// Party that, in a settlement chain interacts with the party 1.
	Party2 *PartyIdentificationAndAccount127 `xml:"Pty2,omitempty"`
}

func (r *ReceivingPartiesAndAccount15) AddDepository() *PartyIdentification102Choice {
	r.Depository = new(PartyIdentification102Choice)
	return r.Depository
}

func (r *ReceivingPartiesAndAccount15) AddParty1() *PartyIdentificationAndAccount126 {
	r.Party1 = new(PartyIdentificationAndAccount126)
	return r.Party1
}

func (r *ReceivingPartiesAndAccount15) AddParty2() *PartyIdentificationAndAccount127 {
	r.Party2 = new(PartyIdentificationAndAccount127)
	return r.Party2
}

// Parameters applied to the settlement of a security transfer.
type ReceivingPartiesAndAccount16 struct {

	// Party that acts on behalf of the buyer of securities when the buyer does not have a direct relationship with the receiving agent.
	ReceiversCustodianDetails *PartyIdentificationAndAccount147 `xml:"RcvrsCtdnDtls,omitempty"`

	// Party that the receiver's custodian uses to effect the receipt of a security, when the receiver's custodian does not have a direct relationship with the receiving agent.
	ReceiversIntermediary1Details *PartyIdentificationAndAccount147 `xml:"RcvrsIntrmy1Dtls,omitempty"`

	// Party that interacts with the receiver’s intermediary 1.
	ReceiversIntermediary2Details *PartyIdentificationAndAccount147 `xml:"RcvrsIntrmy2Dtls,omitempty"`

	// Party that receives securities from the delivering agent at the place of settlement, for example, central securities depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount147 `xml:"RcvgAgtDtls"`
}

func (r *ReceivingPartiesAndAccount16) AddReceiversCustodianDetails() *PartyIdentificationAndAccount147 {
	r.ReceiversCustodianDetails = new(PartyIdentificationAndAccount147)
	return r.ReceiversCustodianDetails
}

func (r *ReceivingPartiesAndAccount16) AddReceiversIntermediary1Details() *PartyIdentificationAndAccount147 {
	r.ReceiversIntermediary1Details = new(PartyIdentificationAndAccount147)
	return r.ReceiversIntermediary1Details
}

func (r *ReceivingPartiesAndAccount16) AddReceiversIntermediary2Details() *PartyIdentificationAndAccount147 {
	r.ReceiversIntermediary2Details = new(PartyIdentificationAndAccount147)
	return r.ReceiversIntermediary2Details
}

func (r *ReceivingPartiesAndAccount16) AddReceivingAgentDetails() *PartyIdentificationAndAccount147 {
	r.ReceivingAgentDetails = new(PartyIdentificationAndAccount147)
	return r.ReceivingAgentDetails
}

// Parameters applied to the settlement of a security transfer.
type ReceivingPartiesAndAccount3 struct {

	// Party that acts on behalf of the buyer of securities when the buyer does not have a direct relationship with the receiving agent.
	ReceiversCustodianDetails *PartyIdentificationAndAccount3 `xml:"RcvrsCtdnDtls,omitempty"`

	// Party that the Receiver's custodian uses to effect the receipt of a security, when the Receiver's custodian does not have a direct relationship with the Receiver agent.
	ReceiversIntermediaryDetails *PartyIdentificationAndAccount3 `xml:"RcvrsIntrmyDtls,omitempty"`

	// Party that receives securities from the delivering agent at the place of settlement, eg, central securities depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount3 `xml:"RcvgAgtDtls"`
}

func (r *ReceivingPartiesAndAccount3) AddReceiversCustodianDetails() *PartyIdentificationAndAccount3 {
	r.ReceiversCustodianDetails = new(PartyIdentificationAndAccount3)
	return r.ReceiversCustodianDetails
}

func (r *ReceivingPartiesAndAccount3) AddReceiversIntermediaryDetails() *PartyIdentificationAndAccount3 {
	r.ReceiversIntermediaryDetails = new(PartyIdentificationAndAccount3)
	return r.ReceiversIntermediaryDetails
}

func (r *ReceivingPartiesAndAccount3) AddReceivingAgentDetails() *PartyIdentificationAndAccount3 {
	r.ReceivingAgentDetails = new(PartyIdentificationAndAccount3)
	return r.ReceivingAgentDetails
}

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type ReceivingPartiesAndAccount4 struct {

	// Party that buys goods or services, or a financial instrument.
	ReceiverDetails *InvestmentAccount24 `xml:"RcvrDtls,omitempty"`

	// Party that acts on behalf of the buyer of securities when the buyer does not have a direct relationship with the receiving agent.
	ReceiversCustodianDetails *PartyIdentificationAndAccount5 `xml:"RcvrsCtdnDtls,omitempty"`

	// Party that the Receiver's custodian uses to effect the receipt of a security, when the Receiver's custodian does not have a direct relationship with the Receiver agent.
	ReceiversIntermediaryDetails *PartyIdentificationAndAccount5 `xml:"RcvrsIntrmyDtls,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, eg, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount4 `xml:"RcvgAgtDtls"`

	// Identifies the securities settlement system to be used.
	SecuritiesSettlementSystem *Max35Text `xml:"SctiesSttlmSys,omitempty"`

	// Place where settlement of the securities takes place.
	PlaceOfSettlementDetails *PartyIdentification21 `xml:"PlcOfSttlmDtls"`
}

func (r *ReceivingPartiesAndAccount4) AddReceiverDetails() *InvestmentAccount24 {
	r.ReceiverDetails = new(InvestmentAccount24)
	return r.ReceiverDetails
}

func (r *ReceivingPartiesAndAccount4) AddReceiversCustodianDetails() *PartyIdentificationAndAccount5 {
	r.ReceiversCustodianDetails = new(PartyIdentificationAndAccount5)
	return r.ReceiversCustodianDetails
}

func (r *ReceivingPartiesAndAccount4) AddReceiversIntermediaryDetails() *PartyIdentificationAndAccount5 {
	r.ReceiversIntermediaryDetails = new(PartyIdentificationAndAccount5)
	return r.ReceiversIntermediaryDetails
}

func (r *ReceivingPartiesAndAccount4) AddReceivingAgentDetails() *PartyIdentificationAndAccount4 {
	r.ReceivingAgentDetails = new(PartyIdentificationAndAccount4)
	return r.ReceivingAgentDetails
}

func (r *ReceivingPartiesAndAccount4) SetSecuritiesSettlementSystem(value string) {
	r.SecuritiesSettlementSystem = (*Max35Text)(&value)
}

func (r *ReceivingPartiesAndAccount4) AddPlaceOfSettlementDetails() *PartyIdentification21 {
	r.PlaceOfSettlementDetails = new(PartyIdentification21)
	return r.PlaceOfSettlementDetails
}

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type ReceivingPartiesAndAccount7 struct {

	// First party in the settlement chain. In a plain vanilla settlement, it is the Central Securities Depository where the counterparty requests to receive the financial instrument or from where the counterparty delivers the financial instruments.
	Depository *PartyIdentification34Choice `xml:"Dpstry"`

	// Party that, in a settlement chain interacts with the depository.
	Party1 *PartyIdentificationAndAccount32 `xml:"Pty1"`

	// Party that, in a settlement chain interacts with the party 1.
	Party2 *PartyIdentificationAndAccount32 `xml:"Pty2,omitempty"`

	// Identifies the securities settlement system to be used.
	SecuritiesSettlementSystem *Max35Text `xml:"SctiesSttlmSys,omitempty"`
}

func (r *ReceivingPartiesAndAccount7) AddDepository() *PartyIdentification34Choice {
	r.Depository = new(PartyIdentification34Choice)
	return r.Depository
}

func (r *ReceivingPartiesAndAccount7) AddParty1() *PartyIdentificationAndAccount32 {
	r.Party1 = new(PartyIdentificationAndAccount32)
	return r.Party1
}

func (r *ReceivingPartiesAndAccount7) AddParty2() *PartyIdentificationAndAccount32 {
	r.Party2 = new(PartyIdentificationAndAccount32)
	return r.Party2
}

func (r *ReceivingPartiesAndAccount7) SetSecuritiesSettlementSystem(value string) {
	r.SecuritiesSettlementSystem = (*Max35Text)(&value)
}

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type ReceivingPartiesAndAccount8 struct {

	// Party that buys goods or services, or a financial instrument.
	ReceiverDetails *InvestmentAccount24 `xml:"RcvrDtls,omitempty"`

	// Party that acts on behalf of the buyer of securities when the buyer does not have a direct relationship with the receiving agent.
	ReceiversCustodianDetails *PartyIdentificationAndAccount5 `xml:"RcvrsCtdnDtls,omitempty"`

	// Party that the Receiver's custodian uses to effect the receipt of a security, when the Receiver's custodian does not have a direct relationship with the Receiver agent.
	ReceiversIntermediaryDetails *PartyIdentificationAndAccount5 `xml:"RcvrsIntrmyDtls,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, eg, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount4 `xml:"RcvgAgtDtls"`

	// Identifies the securities settlement system to be used.
	SecuritiesSettlementSystem *Max35Text `xml:"SctiesSttlmSys,omitempty"`

	// Place where settlement of the securities takes place.
	PlaceOfSettlementDetails *PartyIdentification21 `xml:"PlcOfSttlmDtls,omitempty"`
}

func (r *ReceivingPartiesAndAccount8) AddReceiverDetails() *InvestmentAccount24 {
	r.ReceiverDetails = new(InvestmentAccount24)
	return r.ReceiverDetails
}

func (r *ReceivingPartiesAndAccount8) AddReceiversCustodianDetails() *PartyIdentificationAndAccount5 {
	r.ReceiversCustodianDetails = new(PartyIdentificationAndAccount5)
	return r.ReceiversCustodianDetails
}

func (r *ReceivingPartiesAndAccount8) AddReceiversIntermediaryDetails() *PartyIdentificationAndAccount5 {
	r.ReceiversIntermediaryDetails = new(PartyIdentificationAndAccount5)
	return r.ReceiversIntermediaryDetails
}

func (r *ReceivingPartiesAndAccount8) AddReceivingAgentDetails() *PartyIdentificationAndAccount4 {
	r.ReceivingAgentDetails = new(PartyIdentificationAndAccount4)
	return r.ReceivingAgentDetails
}

func (r *ReceivingPartiesAndAccount8) SetSecuritiesSettlementSystem(value string) {
	r.SecuritiesSettlementSystem = (*Max35Text)(&value)
}

func (r *ReceivingPartiesAndAccount8) AddPlaceOfSettlementDetails() *PartyIdentification21 {
	r.PlaceOfSettlementDetails = new(PartyIdentification21)
	return r.PlaceOfSettlementDetails
}

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type ReceivingPartiesAndAccount9 struct {

	// Party that buys goods or services, or a financial instrument.
	ReceiverDetails *InvestmentAccount41 `xml:"RcvrDtls,omitempty"`

	// Party that acts on behalf of the buyer of securities when the buyer does not have a direct relationship with the receiving agent.
	ReceiversCustodianDetails *PartyIdentificationAndAccount5 `xml:"RcvrsCtdnDtls,omitempty"`

	// Party that the Receiver's custodian uses to effect the receipt of a security, when the Receiver's custodian does not have a direct relationship with the Receiver agent.
	ReceiversIntermediaryDetails *PartyIdentificationAndAccount5 `xml:"RcvrsIntrmyDtls,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, eg, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount4 `xml:"RcvgAgtDtls"`

	// Identifies the securities settlement system to be used.
	SecuritiesSettlementSystem *Max35Text `xml:"SctiesSttlmSys,omitempty"`

	// Place where settlement of the securities takes place.
	PlaceOfSettlementDetails *PartyIdentification21 `xml:"PlcOfSttlmDtls,omitempty"`
}

func (r *ReceivingPartiesAndAccount9) AddReceiverDetails() *InvestmentAccount41 {
	r.ReceiverDetails = new(InvestmentAccount41)
	return r.ReceiverDetails
}

func (r *ReceivingPartiesAndAccount9) AddReceiversCustodianDetails() *PartyIdentificationAndAccount5 {
	r.ReceiversCustodianDetails = new(PartyIdentificationAndAccount5)
	return r.ReceiversCustodianDetails
}

func (r *ReceivingPartiesAndAccount9) AddReceiversIntermediaryDetails() *PartyIdentificationAndAccount5 {
	r.ReceiversIntermediaryDetails = new(PartyIdentificationAndAccount5)
	return r.ReceiversIntermediaryDetails
}

func (r *ReceivingPartiesAndAccount9) AddReceivingAgentDetails() *PartyIdentificationAndAccount4 {
	r.ReceivingAgentDetails = new(PartyIdentificationAndAccount4)
	return r.ReceivingAgentDetails
}

func (r *ReceivingPartiesAndAccount9) SetSecuritiesSettlementSystem(value string) {
	r.SecuritiesSettlementSystem = (*Max35Text)(&value)
}

func (r *ReceivingPartiesAndAccount9) AddPlaceOfSettlementDetails() *PartyIdentification21 {
	r.PlaceOfSettlementDetails = new(PartyIdentification21)
	return r.PlaceOfSettlementDetails
}
