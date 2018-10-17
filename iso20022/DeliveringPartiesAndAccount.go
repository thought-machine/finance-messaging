package iso20022

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type DeliveringPartiesAndAccount1 struct {

	// Party that sells goods or services, or a financial instrument.
	DelivererDetails *InvestmentAccount11 `xml:"DlvrrDtls"`

	// Party that acts on behalf of the seller of securities when the seller does not have a direct relationship with the delivering agent.
	DeliverersCustodianDetails *PartyIdentificationAndAccount2 `xml:"DlvrrsCtdnDtls,omitempty"`

	// Party that the deliverer's custodian uses to effect the delivery of a security, when the deliverer's custodian does not have a direct relationship with the delivering agent.
	DeliverersIntermediaryDetails *PartyIdentificationAndAccount2 `xml:"DlvrrsIntrmyDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, eg, central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount2 `xml:"DlvrgAgtDtls"`

	// Identifies the securities settlement system to be used.
	SecuritiesSettlementSystem *Max35Text `xml:"SctiesSttlmSys,omitempty"`

	// Place where settlement of the securities takes place.
	PlaceOfSettlementDetails *PartyIdentificationAndAccount2 `xml:"PlcOfSttlmDtls"`
}

func (d *DeliveringPartiesAndAccount1) AddDelivererDetails() *InvestmentAccount11 {
	d.DelivererDetails = new(InvestmentAccount11)
	return d.DelivererDetails
}

func (d *DeliveringPartiesAndAccount1) AddDeliverersCustodianDetails() *PartyIdentificationAndAccount2 {
	d.DeliverersCustodianDetails = new(PartyIdentificationAndAccount2)
	return d.DeliverersCustodianDetails
}

func (d *DeliveringPartiesAndAccount1) AddDeliverersIntermediaryDetails() *PartyIdentificationAndAccount2 {
	d.DeliverersIntermediaryDetails = new(PartyIdentificationAndAccount2)
	return d.DeliverersIntermediaryDetails
}

func (d *DeliveringPartiesAndAccount1) AddDeliveringAgentDetails() *PartyIdentificationAndAccount2 {
	d.DeliveringAgentDetails = new(PartyIdentificationAndAccount2)
	return d.DeliveringAgentDetails
}

func (d *DeliveringPartiesAndAccount1) SetSecuritiesSettlementSystem(value string) {
	d.SecuritiesSettlementSystem = (*Max35Text)(&value)
}

func (d *DeliveringPartiesAndAccount1) AddPlaceOfSettlementDetails() *PartyIdentificationAndAccount2 {
	d.PlaceOfSettlementDetails = new(PartyIdentificationAndAccount2)
	return d.PlaceOfSettlementDetails
}

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type DeliveringPartiesAndAccount10 struct {

	// First party in the settlement chain. In a plain vanilla settlement, it is the Central Securities Depository where the counterparty requests to receive the financial instrument or from where the counterparty delivers the financial instruments.
	Depository *PartyIdentification34Choice `xml:"Dpstry"`

	// Party that, in a settlement chain interacts with the depository.
	Party1 *PartyIdentificationAndAccount102 `xml:"Pty1"`

	// Party that, in a settlement chain interacts with the party 1.
	Party2 *PartyIdentificationAndAccount77 `xml:"Pty2,omitempty"`
}

func (d *DeliveringPartiesAndAccount10) AddDepository() *PartyIdentification34Choice {
	d.Depository = new(PartyIdentification34Choice)
	return d.Depository
}

func (d *DeliveringPartiesAndAccount10) AddParty1() *PartyIdentificationAndAccount102 {
	d.Party1 = new(PartyIdentificationAndAccount102)
	return d.Party1
}

func (d *DeliveringPartiesAndAccount10) AddParty2() *PartyIdentificationAndAccount77 {
	d.Party2 = new(PartyIdentificationAndAccount77)
	return d.Party2
}

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type DeliveringPartiesAndAccount11 struct {

	// First party in the settlement chain. In a plain vanilla settlement, it is the Central Securities Depository where the counterparty requests to receive the financial instrument or from where the counterparty delivers the financial instruments.
	Depository *PartyIdentification34Choice `xml:"Dpstry"`

	// Party that, in a settlement chain interacts with the depository.
	Party1 *PartyIdentificationAndAccount102 `xml:"Pty1"`

	// Party that, in a settlement chain interacts with the party 1.
	Party2 *PartyIdentificationAndAccount102 `xml:"Pty2,omitempty"`

	// Identifies the securities settlement system to be used.
	SecuritiesSettlementSystem *Max35Text `xml:"SctiesSttlmSys,omitempty"`
}

func (d *DeliveringPartiesAndAccount11) AddDepository() *PartyIdentification34Choice {
	d.Depository = new(PartyIdentification34Choice)
	return d.Depository
}

func (d *DeliveringPartiesAndAccount11) AddParty1() *PartyIdentificationAndAccount102 {
	d.Party1 = new(PartyIdentificationAndAccount102)
	return d.Party1
}

func (d *DeliveringPartiesAndAccount11) AddParty2() *PartyIdentificationAndAccount102 {
	d.Party2 = new(PartyIdentificationAndAccount102)
	return d.Party2
}

func (d *DeliveringPartiesAndAccount11) SetSecuritiesSettlementSystem(value string) {
	d.SecuritiesSettlementSystem = (*Max35Text)(&value)
}

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type DeliveringPartiesAndAccount13 struct {

	// Party that sells goods or services, or a financial instrument.
	DelivererDetails *InvestmentAccount55 `xml:"DlvrrDtls,omitempty"`

	// Party that acts on behalf of the seller of securities when the seller does not have a direct relationship with the delivering agent.
	DeliverersCustodianDetails *PartyIdentificationAndAccount124 `xml:"DlvrrsCtdnDtls,omitempty"`

	// Party that the deliverer's custodian uses to effect the delivery of a security, when the deliverer's custodian does not have a direct relationship with the delivering agent.
	DeliverersIntermediary1Details *PartyIdentificationAndAccount124 `xml:"DlvrrsIntrmy1Dtls,omitempty"`

	// Party that interacts with the deliverer's intermediary.
	DeliverersIntermediary2Details *PartyIdentificationAndAccount124 `xml:"DlvrrsIntrmy2Dtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount124 `xml:"DlvrgAgtDtls"`

	// Identifies the securities settlement system to be used.
	SecuritiesSettlementSystem *Max35Text `xml:"SctiesSttlmSys,omitempty"`

	// Place where settlement of the securities takes place.
	PlaceOfSettlementDetails *PartyIdentification97 `xml:"PlcOfSttlmDtls,omitempty"`
}

func (d *DeliveringPartiesAndAccount13) AddDelivererDetails() *InvestmentAccount55 {
	d.DelivererDetails = new(InvestmentAccount55)
	return d.DelivererDetails
}

func (d *DeliveringPartiesAndAccount13) AddDeliverersCustodianDetails() *PartyIdentificationAndAccount124 {
	d.DeliverersCustodianDetails = new(PartyIdentificationAndAccount124)
	return d.DeliverersCustodianDetails
}

func (d *DeliveringPartiesAndAccount13) AddDeliverersIntermediary1Details() *PartyIdentificationAndAccount124 {
	d.DeliverersIntermediary1Details = new(PartyIdentificationAndAccount124)
	return d.DeliverersIntermediary1Details
}

func (d *DeliveringPartiesAndAccount13) AddDeliverersIntermediary2Details() *PartyIdentificationAndAccount124 {
	d.DeliverersIntermediary2Details = new(PartyIdentificationAndAccount124)
	return d.DeliverersIntermediary2Details
}

func (d *DeliveringPartiesAndAccount13) AddDeliveringAgentDetails() *PartyIdentificationAndAccount124 {
	d.DeliveringAgentDetails = new(PartyIdentificationAndAccount124)
	return d.DeliveringAgentDetails
}

func (d *DeliveringPartiesAndAccount13) SetSecuritiesSettlementSystem(value string) {
	d.SecuritiesSettlementSystem = (*Max35Text)(&value)
}

func (d *DeliveringPartiesAndAccount13) AddPlaceOfSettlementDetails() *PartyIdentification97 {
	d.PlaceOfSettlementDetails = new(PartyIdentification97)
	return d.PlaceOfSettlementDetails
}

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type DeliveringPartiesAndAccount14 struct {

	// Party that acts on behalf of the seller of securities when the seller does not have a direct relationship with the delivering agent.
	DeliverersCustodianDetails *PartyIdentificationAndAccount124 `xml:"DlvrrsCtdnDtls,omitempty"`

	// Party that the deliverer's custodian uses to effect the delivery of a security, when the deliverer's custodian does not have a direct relationship with the delivering agent.
	DeliverersIntermediary1Details *PartyIdentificationAndAccount124 `xml:"DlvrrsIntrmy1Dtls,omitempty"`

	// Party that interacts with the deliverer's intermediary.
	DeliverersIntermediary2Details *PartyIdentificationAndAccount124 `xml:"DlvrrsIntrmy2Dtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount124 `xml:"DlvrgAgtDtls"`

	// Identifies the securities settlement system to be used.
	SecuritiesSettlementSystem *Max35Text `xml:"SctiesSttlmSys,omitempty"`

	// Place where settlement of the securities takes place.
	PlaceOfSettlementDetails *PartyIdentification97 `xml:"PlcOfSttlmDtls,omitempty"`
}

func (d *DeliveringPartiesAndAccount14) AddDeliverersCustodianDetails() *PartyIdentificationAndAccount124 {
	d.DeliverersCustodianDetails = new(PartyIdentificationAndAccount124)
	return d.DeliverersCustodianDetails
}

func (d *DeliveringPartiesAndAccount14) AddDeliverersIntermediary1Details() *PartyIdentificationAndAccount124 {
	d.DeliverersIntermediary1Details = new(PartyIdentificationAndAccount124)
	return d.DeliverersIntermediary1Details
}

func (d *DeliveringPartiesAndAccount14) AddDeliverersIntermediary2Details() *PartyIdentificationAndAccount124 {
	d.DeliverersIntermediary2Details = new(PartyIdentificationAndAccount124)
	return d.DeliverersIntermediary2Details
}

func (d *DeliveringPartiesAndAccount14) AddDeliveringAgentDetails() *PartyIdentificationAndAccount124 {
	d.DeliveringAgentDetails = new(PartyIdentificationAndAccount124)
	return d.DeliveringAgentDetails
}

func (d *DeliveringPartiesAndAccount14) SetSecuritiesSettlementSystem(value string) {
	d.SecuritiesSettlementSystem = (*Max35Text)(&value)
}

func (d *DeliveringPartiesAndAccount14) AddPlaceOfSettlementDetails() *PartyIdentification97 {
	d.PlaceOfSettlementDetails = new(PartyIdentification97)
	return d.PlaceOfSettlementDetails
}

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type DeliveringPartiesAndAccount15 struct {

	// First party in the settlement chain. In a plain vanilla settlement, it is the Central Securities Depository where the counterparty requests to receive the financial instrument or from where the counterparty delivers the financial instruments.
	Depository *PartyIdentification102Choice `xml:"Dpstry"`

	// Party that, in a settlement chain interacts with the depository.
	Party1 *PartyIdentificationAndAccount126 `xml:"Pty1"`

	// Party that, in a settlement chain interacts with the party 1.
	Party2 *PartyIdentificationAndAccount127 `xml:"Pty2,omitempty"`
}

func (d *DeliveringPartiesAndAccount15) AddDepository() *PartyIdentification102Choice {
	d.Depository = new(PartyIdentification102Choice)
	return d.Depository
}

func (d *DeliveringPartiesAndAccount15) AddParty1() *PartyIdentificationAndAccount126 {
	d.Party1 = new(PartyIdentificationAndAccount126)
	return d.Party1
}

func (d *DeliveringPartiesAndAccount15) AddParty2() *PartyIdentificationAndAccount127 {
	d.Party2 = new(PartyIdentificationAndAccount127)
	return d.Party2
}

// Parameters applied to the settlement of a security transfer.
type DeliveringPartiesAndAccount16 struct {

	// Party that acts on behalf of the seller of securities when the seller does not have a direct relationship with the delivering agent.
	DeliverersCustodianDetails *PartyIdentificationAndAccount147 `xml:"DlvrrsCtdnDtls,omitempty"`

	// Party that the deliverer's custodian uses to effect the delivery of a security, when the deliverer's custodian does not have a direct relationship with the delivering agent.
	DeliverersIntermediary1Details *PartyIdentificationAndAccount147 `xml:"DlvrrsIntrmy1Dtls,omitempty"`

	// Party that interacts with the deliverer's intermediary 1.
	DeliverersIntermediary2Details *PartyIdentificationAndAccount147 `xml:"DlvrrsIntrmy2Dtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount147 `xml:"DlvrgAgtDtls"`
}

func (d *DeliveringPartiesAndAccount16) AddDeliverersCustodianDetails() *PartyIdentificationAndAccount147 {
	d.DeliverersCustodianDetails = new(PartyIdentificationAndAccount147)
	return d.DeliverersCustodianDetails
}

func (d *DeliveringPartiesAndAccount16) AddDeliverersIntermediary1Details() *PartyIdentificationAndAccount147 {
	d.DeliverersIntermediary1Details = new(PartyIdentificationAndAccount147)
	return d.DeliverersIntermediary1Details
}

func (d *DeliveringPartiesAndAccount16) AddDeliverersIntermediary2Details() *PartyIdentificationAndAccount147 {
	d.DeliverersIntermediary2Details = new(PartyIdentificationAndAccount147)
	return d.DeliverersIntermediary2Details
}

func (d *DeliveringPartiesAndAccount16) AddDeliveringAgentDetails() *PartyIdentificationAndAccount147 {
	d.DeliveringAgentDetails = new(PartyIdentificationAndAccount147)
	return d.DeliveringAgentDetails
}

// Parameters applied to the settlement of a security transfer.
type DeliveringPartiesAndAccount3 struct {

	// Party that acts on behalf of the seller of securities when the seller does not have a direct relationship with the delivering agent.
	DeliverersCustodianDetails *PartyIdentificationAndAccount3 `xml:"DlvrrsCtdnDtls,omitempty"`

	// Party that the deliverer's custodian uses to effect the delivery of a security, when the deliverer's custodian does not have a direct relationship with the delivering agent.
	DeliverersIntermediaryDetails *PartyIdentificationAndAccount3 `xml:"DlvrrsIntrmyDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, eg, central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount3 `xml:"DlvrgAgtDtls"`
}

func (d *DeliveringPartiesAndAccount3) AddDeliverersCustodianDetails() *PartyIdentificationAndAccount3 {
	d.DeliverersCustodianDetails = new(PartyIdentificationAndAccount3)
	return d.DeliverersCustodianDetails
}

func (d *DeliveringPartiesAndAccount3) AddDeliverersIntermediaryDetails() *PartyIdentificationAndAccount3 {
	d.DeliverersIntermediaryDetails = new(PartyIdentificationAndAccount3)
	return d.DeliverersIntermediaryDetails
}

func (d *DeliveringPartiesAndAccount3) AddDeliveringAgentDetails() *PartyIdentificationAndAccount3 {
	d.DeliveringAgentDetails = new(PartyIdentificationAndAccount3)
	return d.DeliveringAgentDetails
}

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type DeliveringPartiesAndAccount4 struct {

	// Party that sells goods or services, or a financial instrument.
	DelivererDetails *InvestmentAccount24 `xml:"DlvrrDtls,omitempty"`

	// Party that acts on behalf of the seller of securities when the seller does not have a direct relationship with the delivering agent.
	DeliverersCustodianDetails *PartyIdentificationAndAccount5 `xml:"DlvrrsCtdnDtls,omitempty"`

	// Party that the deliverer's custodian uses to effect the delivery of a security, when the deliverer's custodian does not have a direct relationship with the delivering agent.
	DeliverersIntermediaryDetails *PartyIdentificationAndAccount5 `xml:"DlvrrsIntrmyDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, eg, central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount4 `xml:"DlvrgAgtDtls"`

	// Identifies the securities settlement system to be used.
	SecuritiesSettlementSystem *Max35Text `xml:"SctiesSttlmSys,omitempty"`

	// Place where settlement of the securities takes place.
	PlaceOfSettlementDetails *PartyIdentification21 `xml:"PlcOfSttlmDtls"`
}

func (d *DeliveringPartiesAndAccount4) AddDelivererDetails() *InvestmentAccount24 {
	d.DelivererDetails = new(InvestmentAccount24)
	return d.DelivererDetails
}

func (d *DeliveringPartiesAndAccount4) AddDeliverersCustodianDetails() *PartyIdentificationAndAccount5 {
	d.DeliverersCustodianDetails = new(PartyIdentificationAndAccount5)
	return d.DeliverersCustodianDetails
}

func (d *DeliveringPartiesAndAccount4) AddDeliverersIntermediaryDetails() *PartyIdentificationAndAccount5 {
	d.DeliverersIntermediaryDetails = new(PartyIdentificationAndAccount5)
	return d.DeliverersIntermediaryDetails
}

func (d *DeliveringPartiesAndAccount4) AddDeliveringAgentDetails() *PartyIdentificationAndAccount4 {
	d.DeliveringAgentDetails = new(PartyIdentificationAndAccount4)
	return d.DeliveringAgentDetails
}

func (d *DeliveringPartiesAndAccount4) SetSecuritiesSettlementSystem(value string) {
	d.SecuritiesSettlementSystem = (*Max35Text)(&value)
}

func (d *DeliveringPartiesAndAccount4) AddPlaceOfSettlementDetails() *PartyIdentification21 {
	d.PlaceOfSettlementDetails = new(PartyIdentification21)
	return d.PlaceOfSettlementDetails
}

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type DeliveringPartiesAndAccount7 struct {

	// First party in the settlement chain. In a plain vanilla settlement, it is the Central Securities Depository where the counterparty requests to receive the financial instrument or from where the counterparty delivers the financial instruments.
	Depository *PartyIdentification34Choice `xml:"Dpstry"`

	// Party that, in a settlement chain interacts with the depository.
	Party1 *PartyIdentificationAndAccount32 `xml:"Pty1"`

	// Party that, in a settlement chain interacts with the party 1.
	Party2 *PartyIdentificationAndAccount32 `xml:"Pty2,omitempty"`

	// Identifies the securities settlement system to be used.
	SecuritiesSettlementSystem *Max35Text `xml:"SctiesSttlmSys,omitempty"`
}

func (d *DeliveringPartiesAndAccount7) AddDepository() *PartyIdentification34Choice {
	d.Depository = new(PartyIdentification34Choice)
	return d.Depository
}

func (d *DeliveringPartiesAndAccount7) AddParty1() *PartyIdentificationAndAccount32 {
	d.Party1 = new(PartyIdentificationAndAccount32)
	return d.Party1
}

func (d *DeliveringPartiesAndAccount7) AddParty2() *PartyIdentificationAndAccount32 {
	d.Party2 = new(PartyIdentificationAndAccount32)
	return d.Party2
}

func (d *DeliveringPartiesAndAccount7) SetSecuritiesSettlementSystem(value string) {
	d.SecuritiesSettlementSystem = (*Max35Text)(&value)
}

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type DeliveringPartiesAndAccount8 struct {

	// Party that sells goods or services, or a financial instrument.
	DelivererDetails *InvestmentAccount24 `xml:"DlvrrDtls,omitempty"`

	// Party that acts on behalf of the seller of securities when the seller does not have a direct relationship with the delivering agent.
	DeliverersCustodianDetails *PartyIdentificationAndAccount5 `xml:"DlvrrsCtdnDtls,omitempty"`

	// Party that the deliverer's custodian uses to effect the delivery of a security, when the deliverer's custodian does not have a direct relationship with the delivering agent.
	DeliverersIntermediaryDetails *PartyIdentificationAndAccount5 `xml:"DlvrrsIntrmyDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, eg, central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount4 `xml:"DlvrgAgtDtls"`

	// Identifies the securities settlement system to be used.
	SecuritiesSettlementSystem *Max35Text `xml:"SctiesSttlmSys,omitempty"`

	// Place where settlement of the securities takes place.
	PlaceOfSettlementDetails *PartyIdentification21 `xml:"PlcOfSttlmDtls,omitempty"`
}

func (d *DeliveringPartiesAndAccount8) AddDelivererDetails() *InvestmentAccount24 {
	d.DelivererDetails = new(InvestmentAccount24)
	return d.DelivererDetails
}

func (d *DeliveringPartiesAndAccount8) AddDeliverersCustodianDetails() *PartyIdentificationAndAccount5 {
	d.DeliverersCustodianDetails = new(PartyIdentificationAndAccount5)
	return d.DeliverersCustodianDetails
}

func (d *DeliveringPartiesAndAccount8) AddDeliverersIntermediaryDetails() *PartyIdentificationAndAccount5 {
	d.DeliverersIntermediaryDetails = new(PartyIdentificationAndAccount5)
	return d.DeliverersIntermediaryDetails
}

func (d *DeliveringPartiesAndAccount8) AddDeliveringAgentDetails() *PartyIdentificationAndAccount4 {
	d.DeliveringAgentDetails = new(PartyIdentificationAndAccount4)
	return d.DeliveringAgentDetails
}

func (d *DeliveringPartiesAndAccount8) SetSecuritiesSettlementSystem(value string) {
	d.SecuritiesSettlementSystem = (*Max35Text)(&value)
}

func (d *DeliveringPartiesAndAccount8) AddPlaceOfSettlementDetails() *PartyIdentification21 {
	d.PlaceOfSettlementDetails = new(PartyIdentification21)
	return d.PlaceOfSettlementDetails
}

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type DeliveringPartiesAndAccount9 struct {

	// Party that sells goods or services, or a financial instrument.
	DelivererDetails *InvestmentAccount41 `xml:"DlvrrDtls,omitempty"`

	// Party that acts on behalf of the seller of securities when the seller does not have a direct relationship with the delivering agent.
	DeliverersCustodianDetails *PartyIdentificationAndAccount5 `xml:"DlvrrsCtdnDtls,omitempty"`

	// Party that the deliverer's custodian uses to effect the delivery of a security, when the deliverer's custodian does not have a direct relationship with the delivering agent.
	DeliverersIntermediaryDetails *PartyIdentificationAndAccount5 `xml:"DlvrrsIntrmyDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, eg, central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount4 `xml:"DlvrgAgtDtls"`

	// Identifies the securities settlement system to be used.
	SecuritiesSettlementSystem *Max35Text `xml:"SctiesSttlmSys,omitempty"`

	// Place where settlement of the securities takes place.
	PlaceOfSettlementDetails *PartyIdentification21 `xml:"PlcOfSttlmDtls,omitempty"`
}

func (d *DeliveringPartiesAndAccount9) AddDelivererDetails() *InvestmentAccount41 {
	d.DelivererDetails = new(InvestmentAccount41)
	return d.DelivererDetails
}

func (d *DeliveringPartiesAndAccount9) AddDeliverersCustodianDetails() *PartyIdentificationAndAccount5 {
	d.DeliverersCustodianDetails = new(PartyIdentificationAndAccount5)
	return d.DeliverersCustodianDetails
}

func (d *DeliveringPartiesAndAccount9) AddDeliverersIntermediaryDetails() *PartyIdentificationAndAccount5 {
	d.DeliverersIntermediaryDetails = new(PartyIdentificationAndAccount5)
	return d.DeliverersIntermediaryDetails
}

func (d *DeliveringPartiesAndAccount9) AddDeliveringAgentDetails() *PartyIdentificationAndAccount4 {
	d.DeliveringAgentDetails = new(PartyIdentificationAndAccount4)
	return d.DeliveringAgentDetails
}

func (d *DeliveringPartiesAndAccount9) SetSecuritiesSettlementSystem(value string) {
	d.SecuritiesSettlementSystem = (*Max35Text)(&value)
}

func (d *DeliveringPartiesAndAccount9) AddPlaceOfSettlementDetails() *PartyIdentification21 {
	d.PlaceOfSettlementDetails = new(PartyIdentification21)
	return d.PlaceOfSettlementDetails
}
