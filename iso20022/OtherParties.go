package iso20022

// Other parties information.
type OtherParties10 struct {

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification37Choice `xml:"Invstr,omitempty"`

	// Identification of the stock exchange to which transaction reporting will be done.
	StockExchange *PartyIdentification49Choice `xml:"StockXchg,omitempty"`

	// Institution to which a trade must be reported.
	TradeRegulator *PartyIdentification49Choice `xml:"TradRgltr,omitempty"`
}

func (o *OtherParties10) AddInvestor() *PartyIdentification37Choice {
	o.Investor = new(PartyIdentification37Choice)
	return o.Investor
}

func (o *OtherParties10) AddStockExchange() *PartyIdentification49Choice {
	o.StockExchange = new(PartyIdentification49Choice)
	return o.StockExchange
}

func (o *OtherParties10) AddTradeRegulator() *PartyIdentification49Choice {
	o.TradeRegulator = new(PartyIdentification49Choice)
	return o.TradeRegulator
}

// Other parties information.
type OtherParties11 struct {

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification37Choice `xml:"Invstr,omitempty"`

	// Foreign Financial Institution which has been authorised by local authorities to act as account management institution in the country.
	QualifiedForeignIntermediary *PartyIdentification43Choice `xml:"QlfdFrgnIntrmy,omitempty"`

	// Identification of the stock exchange to which transaction reporting will be done.
	StockExchange *PartyIdentification43Choice `xml:"StockXchg,omitempty"`

	// Institution to which a trade must be reported.
	TradeRegulator *PartyIdentification43Choice `xml:"TradRgltr,omitempty"`

	// Party responsible for the administration of a tri-party collateral transaction including collateral allocation, marking to market and substitution of collateral.
	TripartyAgent *PartyIdentification43Choice `xml:"TrptyAgt,omitempty"`
}

func (o *OtherParties11) AddInvestor() *PartyIdentification37Choice {
	o.Investor = new(PartyIdentification37Choice)
	return o.Investor
}

func (o *OtherParties11) AddQualifiedForeignIntermediary() *PartyIdentification43Choice {
	o.QualifiedForeignIntermediary = new(PartyIdentification43Choice)
	return o.QualifiedForeignIntermediary
}

func (o *OtherParties11) AddStockExchange() *PartyIdentification43Choice {
	o.StockExchange = new(PartyIdentification43Choice)
	return o.StockExchange
}

func (o *OtherParties11) AddTradeRegulator() *PartyIdentification43Choice {
	o.TradeRegulator = new(PartyIdentification43Choice)
	return o.TradeRegulator
}

func (o *OtherParties11) AddTripartyAgent() *PartyIdentification43Choice {
	o.TripartyAgent = new(PartyIdentification43Choice)
	return o.TripartyAgent
}

// Other parties information.
type OtherParties12 struct {

	// Party, either an individual or organisation, whose assets are being invested.
	Investor []*PartyIdentificationAndAccount46 `xml:"Invstr,omitempty"`

	// Foreign Financial Institution which has been authorised by local authorities to act as account management institution in the country.
	QualifiedForeignIntermediary *PartyIdentificationAndAccount47 `xml:"QlfdFrgnIntrmy,omitempty"`

	// Identification of the stock exchange to which transaction reporting will be done.
	StockExchange *PartyIdentificationAndAccount47 `xml:"StockXchg,omitempty"`

	// Institution to which a trade must be reported.
	TradeRegulator *PartyIdentificationAndAccount47 `xml:"TradRgltr,omitempty"`

	// Party responsible for the administration of a tri-party collateral transaction including collateral allocation, marking to market and substitution of collateral.
	TripartyAgent *PartyIdentificationAndAccount47 `xml:"TrptyAgt,omitempty"`

	// Party that identifies a broker when required (for example, authorised broker, prime broker, etc).
	Broker *PartyIdentificationAndAccount41 `xml:"Brkr,omitempty"`
}

func (o *OtherParties12) AddInvestor() *PartyIdentificationAndAccount46 {
	newValue := new(PartyIdentificationAndAccount46)
	o.Investor = append(o.Investor, newValue)
	return newValue
}

func (o *OtherParties12) AddQualifiedForeignIntermediary() *PartyIdentificationAndAccount47 {
	o.QualifiedForeignIntermediary = new(PartyIdentificationAndAccount47)
	return o.QualifiedForeignIntermediary
}

func (o *OtherParties12) AddStockExchange() *PartyIdentificationAndAccount47 {
	o.StockExchange = new(PartyIdentificationAndAccount47)
	return o.StockExchange
}

func (o *OtherParties12) AddTradeRegulator() *PartyIdentificationAndAccount47 {
	o.TradeRegulator = new(PartyIdentificationAndAccount47)
	return o.TradeRegulator
}

func (o *OtherParties12) AddTripartyAgent() *PartyIdentificationAndAccount47 {
	o.TripartyAgent = new(PartyIdentificationAndAccount47)
	return o.TripartyAgent
}

func (o *OtherParties12) AddBroker() *PartyIdentificationAndAccount41 {
	o.Broker = new(PartyIdentificationAndAccount41)
	return o.Broker
}

// Provides details about business parties involved in the transaction.
type OtherParties18 struct {

	// Party that identifies the underlying investor.
	Investor []*PartyIdentificationAndAccount79 `xml:"Invstr,omitempty"`

	// Party that identifies the stock exchange.
	StockExchange *PartyIdentificationAndAccount87 `xml:"StockXchg,omitempty"`

	// Party that identifies the trade regulator.
	TradeRegulator *PartyIdentificationAndAccount87 `xml:"TradRgltr,omitempty"`

	// Party that handles tri-party transactions.
	TripartyAgent *PartyIdentificationAndAccount83 `xml:"TrptyAgt,omitempty"`

	// Foreign Financial Institution which has been authorised by local authorities to act as account management institution in the country.
	QualifiedForeignIntermediary *PartyIdentificationAndAccount77 `xml:"QlfdFrgnIntrmy,omitempty"`
}

func (o *OtherParties18) AddInvestor() *PartyIdentificationAndAccount79 {
	newValue := new(PartyIdentificationAndAccount79)
	o.Investor = append(o.Investor, newValue)
	return newValue
}

func (o *OtherParties18) AddStockExchange() *PartyIdentificationAndAccount87 {
	o.StockExchange = new(PartyIdentificationAndAccount87)
	return o.StockExchange
}

func (o *OtherParties18) AddTradeRegulator() *PartyIdentificationAndAccount87 {
	o.TradeRegulator = new(PartyIdentificationAndAccount87)
	return o.TradeRegulator
}

func (o *OtherParties18) AddTripartyAgent() *PartyIdentificationAndAccount83 {
	o.TripartyAgent = new(PartyIdentificationAndAccount83)
	return o.TripartyAgent
}

func (o *OtherParties18) AddQualifiedForeignIntermediary() *PartyIdentificationAndAccount77 {
	o.QualifiedForeignIntermediary = new(PartyIdentificationAndAccount77)
	return o.QualifiedForeignIntermediary
}

// Other parties information.
type OtherParties19 struct {

	// Party, either an individual or organisation, whose assets are being invested.
	Investor []*PartyIdentificationAndAccount81 `xml:"Invstr,omitempty"`

	// Foreign Financial Institution which has been authorised by local authorities to act as account management institution in the country.
	QualifiedForeignIntermediary *PartyIdentificationAndAccount41 `xml:"QlfdFrgnIntrmy,omitempty"`

	// Identification of the stock exchange to which transaction reporting will be done.
	StockExchange *PartyIdentificationAndAccount86 `xml:"StockXchg,omitempty"`

	// Institution to which a trade must be reported.
	TradeRegulator *PartyIdentificationAndAccount86 `xml:"TradRgltr,omitempty"`

	// Party responsible for the administration of a tri-party collateral transaction including collateral allocation, marking to market and substitution of collateral.
	TripartyAgent *PartyIdentificationAndAccount41 `xml:"TrptyAgt,omitempty"`

	// Party that identifies a broker when required (for example, authorised broker, prime broker, etc).
	Broker *PartyIdentificationAndAccount41 `xml:"Brkr,omitempty"`
}

func (o *OtherParties19) AddInvestor() *PartyIdentificationAndAccount81 {
	newValue := new(PartyIdentificationAndAccount81)
	o.Investor = append(o.Investor, newValue)
	return newValue
}

func (o *OtherParties19) AddQualifiedForeignIntermediary() *PartyIdentificationAndAccount41 {
	o.QualifiedForeignIntermediary = new(PartyIdentificationAndAccount41)
	return o.QualifiedForeignIntermediary
}

func (o *OtherParties19) AddStockExchange() *PartyIdentificationAndAccount86 {
	o.StockExchange = new(PartyIdentificationAndAccount86)
	return o.StockExchange
}

func (o *OtherParties19) AddTradeRegulator() *PartyIdentificationAndAccount86 {
	o.TradeRegulator = new(PartyIdentificationAndAccount86)
	return o.TradeRegulator
}

func (o *OtherParties19) AddTripartyAgent() *PartyIdentificationAndAccount41 {
	o.TripartyAgent = new(PartyIdentificationAndAccount41)
	return o.TripartyAgent
}

func (o *OtherParties19) AddBroker() *PartyIdentificationAndAccount41 {
	o.Broker = new(PartyIdentificationAndAccount41)
	return o.Broker
}

// Other parties information.
type OtherParties2 struct {

	// Party, either an individual or organisation, whose assets are being invested.
	Investor []*PartyIdentificationAndAccount19 `xml:"Invstr,omitempty"`

	// Foreign Financial Institution which has been authorised by local authorities to act as account management institution in the country.
	QualifiedForeignIntermediary *PartyIdentificationAndAccount21 `xml:"QlfdFrgnIntrmy,omitempty"`

	// Identification of the stock exchange to which transaction reporting will be done.
	StockExchange *PartyIdentificationAndAccount21 `xml:"StockXchg,omitempty"`

	// Institution to which a trade must be reported.
	TradeRegulator *PartyIdentificationAndAccount21 `xml:"TradRgltr,omitempty"`

	// Party responsible for the administration of a tri-party collateral transaction including collateral allocation, marking to market and substitution of collateral.
	TripartyAgent *PartyIdentificationAndAccount21 `xml:"TrptyAgt,omitempty"`
}

func (o *OtherParties2) AddInvestor() *PartyIdentificationAndAccount19 {
	newValue := new(PartyIdentificationAndAccount19)
	o.Investor = append(o.Investor, newValue)
	return newValue
}

func (o *OtherParties2) AddQualifiedForeignIntermediary() *PartyIdentificationAndAccount21 {
	o.QualifiedForeignIntermediary = new(PartyIdentificationAndAccount21)
	return o.QualifiedForeignIntermediary
}

func (o *OtherParties2) AddStockExchange() *PartyIdentificationAndAccount21 {
	o.StockExchange = new(PartyIdentificationAndAccount21)
	return o.StockExchange
}

func (o *OtherParties2) AddTradeRegulator() *PartyIdentificationAndAccount21 {
	o.TradeRegulator = new(PartyIdentificationAndAccount21)
	return o.TradeRegulator
}

func (o *OtherParties2) AddTripartyAgent() *PartyIdentificationAndAccount21 {
	o.TripartyAgent = new(PartyIdentificationAndAccount21)
	return o.TripartyAgent
}

// Other parties information.
type OtherParties26 struct {

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification99 `xml:"Invstr,omitempty"`

	// Identification of the stock exchange to which transaction reporting will be done.
	StockExchange *PartyIdentification100 `xml:"StockXchg,omitempty"`

	// Institution to which a trade must be reported.
	TradeRegulator *PartyIdentification100 `xml:"TradRgltr,omitempty"`
}

func (o *OtherParties26) AddInvestor() *PartyIdentification99 {
	o.Investor = new(PartyIdentification99)
	return o.Investor
}

func (o *OtherParties26) AddStockExchange() *PartyIdentification100 {
	o.StockExchange = new(PartyIdentification100)
	return o.StockExchange
}

func (o *OtherParties26) AddTradeRegulator() *PartyIdentification100 {
	o.TradeRegulator = new(PartyIdentification100)
	return o.TradeRegulator
}

// Other parties information.
type OtherParties27 struct {

	// Party, either an individual or organisation, whose assets are being invested.
	Investor []*PartyIdentificationAndAccount108 `xml:"Invstr,omitempty"`

	// Foreign Financial Institution which has been authorised by local authorities to act as account management institution in the country.
	QualifiedForeignIntermediary *PartyIdentificationAndAccount107 `xml:"QlfdFrgnIntrmy,omitempty"`

	// Identification of the stock exchange to which transaction reporting will be done.
	StockExchange *PartyIdentificationAndAccount109 `xml:"StockXchg,omitempty"`

	// Institution to which a trade must be reported.
	TradeRegulator *PartyIdentificationAndAccount109 `xml:"TradRgltr,omitempty"`

	// Party responsible for the administration of a tri-party collateral transaction including collateral allocation, marking to market and substitution of collateral.
	TripartyAgent *PartyIdentificationAndAccount107 `xml:"TrptyAgt,omitempty"`

	// Party that identifies a broker when required (for example, authorised broker, prime broker, etc).
	Broker *PartyIdentificationAndAccount107 `xml:"Brkr,omitempty"`
}

func (o *OtherParties27) AddInvestor() *PartyIdentificationAndAccount108 {
	newValue := new(PartyIdentificationAndAccount108)
	o.Investor = append(o.Investor, newValue)
	return newValue
}

func (o *OtherParties27) AddQualifiedForeignIntermediary() *PartyIdentificationAndAccount107 {
	o.QualifiedForeignIntermediary = new(PartyIdentificationAndAccount107)
	return o.QualifiedForeignIntermediary
}

func (o *OtherParties27) AddStockExchange() *PartyIdentificationAndAccount109 {
	o.StockExchange = new(PartyIdentificationAndAccount109)
	return o.StockExchange
}

func (o *OtherParties27) AddTradeRegulator() *PartyIdentificationAndAccount109 {
	o.TradeRegulator = new(PartyIdentificationAndAccount109)
	return o.TradeRegulator
}

func (o *OtherParties27) AddTripartyAgent() *PartyIdentificationAndAccount107 {
	o.TripartyAgent = new(PartyIdentificationAndAccount107)
	return o.TripartyAgent
}

func (o *OtherParties27) AddBroker() *PartyIdentificationAndAccount107 {
	o.Broker = new(PartyIdentificationAndAccount107)
	return o.Broker
}

// Other parties information.
type OtherParties28 struct {

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification99 `xml:"Invstr,omitempty"`

	// Foreign Financial Institution which has been authorised by local authorities to act as account management institution in the country.
	QualifiedForeignIntermediary *PartyIdentification100 `xml:"QlfdFrgnIntrmy,omitempty"`

	// Identification of the stock exchange to which transaction reporting will be done.
	StockExchange *PartyIdentification100 `xml:"StockXchg,omitempty"`

	// Institution to which a trade must be reported.
	TradeRegulator *PartyIdentification100 `xml:"TradRgltr,omitempty"`

	// Party responsible for the administration of a tri-party collateral transaction including collateral allocation, marking to market and substitution of collateral.
	TripartyAgent *PartyIdentification100 `xml:"TrptyAgt,omitempty"`
}

func (o *OtherParties28) AddInvestor() *PartyIdentification99 {
	o.Investor = new(PartyIdentification99)
	return o.Investor
}

func (o *OtherParties28) AddQualifiedForeignIntermediary() *PartyIdentification100 {
	o.QualifiedForeignIntermediary = new(PartyIdentification100)
	return o.QualifiedForeignIntermediary
}

func (o *OtherParties28) AddStockExchange() *PartyIdentification100 {
	o.StockExchange = new(PartyIdentification100)
	return o.StockExchange
}

func (o *OtherParties28) AddTradeRegulator() *PartyIdentification100 {
	o.TradeRegulator = new(PartyIdentification100)
	return o.TradeRegulator
}

func (o *OtherParties28) AddTripartyAgent() *PartyIdentification100 {
	o.TripartyAgent = new(PartyIdentification100)
	return o.TripartyAgent
}

// Other parties information.
type OtherParties29 struct {

	// Party, either an individual or organisation, whose assets are being invested.
	Investor []*PartyIdentificationAndAccount135 `xml:"Invstr,omitempty"`

	// Foreign Financial Institution which has been authorised by local authorities to act as account management institution in the country.
	QualifiedForeignIntermediary *PartyIdentificationAndAccount136 `xml:"QlfdFrgnIntrmy,omitempty"`

	// Identification of the stock exchange to which transaction reporting will be done.
	StockExchange *PartyIdentificationAndAccount137 `xml:"StockXchg,omitempty"`

	// Institution to which a trade must be reported.
	TradeRegulator *PartyIdentificationAndAccount137 `xml:"TradRgltr,omitempty"`

	// Party responsible for the administration of a tri-party collateral transaction including collateral allocation, marking to market and substitution of collateral.
	TripartyAgent *PartyIdentificationAndAccount136 `xml:"TrptyAgt,omitempty"`

	// Party that identifies a broker when required (for example, authorised broker, prime broker, etc).
	Broker *PartyIdentificationAndAccount136 `xml:"Brkr,omitempty"`
}

func (o *OtherParties29) AddInvestor() *PartyIdentificationAndAccount135 {
	newValue := new(PartyIdentificationAndAccount135)
	o.Investor = append(o.Investor, newValue)
	return newValue
}

func (o *OtherParties29) AddQualifiedForeignIntermediary() *PartyIdentificationAndAccount136 {
	o.QualifiedForeignIntermediary = new(PartyIdentificationAndAccount136)
	return o.QualifiedForeignIntermediary
}

func (o *OtherParties29) AddStockExchange() *PartyIdentificationAndAccount137 {
	o.StockExchange = new(PartyIdentificationAndAccount137)
	return o.StockExchange
}

func (o *OtherParties29) AddTradeRegulator() *PartyIdentificationAndAccount137 {
	o.TradeRegulator = new(PartyIdentificationAndAccount137)
	return o.TradeRegulator
}

func (o *OtherParties29) AddTripartyAgent() *PartyIdentificationAndAccount136 {
	o.TripartyAgent = new(PartyIdentificationAndAccount136)
	return o.TripartyAgent
}

func (o *OtherParties29) AddBroker() *PartyIdentificationAndAccount136 {
	o.Broker = new(PartyIdentificationAndAccount136)
	return o.Broker
}

// Other parties information.
type OtherParties3 struct {

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification14Choice `xml:"Invstr,omitempty"`

	// Foreign Financial Institution which has been authorised by local authorities to act as account management institution in the country.
	QualifiedForeignIntermediary *PartyIdentification10Choice `xml:"QlfdFrgnIntrmy,omitempty"`

	// Identification of the stock exchange to which transaction reporting will be done.
	StockExchange *PartyIdentification10Choice `xml:"StockXchg,omitempty"`

	// Institution to which a trade must be reported.
	TradeRegulator *PartyIdentification10Choice `xml:"TradRgltr,omitempty"`

	// Party responsible for the administration of a tri-party collateral transaction including collateral allocation, marking to market and substitution of collateral.
	TripartyAgent *PartyIdentification10Choice `xml:"TrptyAgt,omitempty"`
}

func (o *OtherParties3) AddInvestor() *PartyIdentification14Choice {
	o.Investor = new(PartyIdentification14Choice)
	return o.Investor
}

func (o *OtherParties3) AddQualifiedForeignIntermediary() *PartyIdentification10Choice {
	o.QualifiedForeignIntermediary = new(PartyIdentification10Choice)
	return o.QualifiedForeignIntermediary
}

func (o *OtherParties3) AddStockExchange() *PartyIdentification10Choice {
	o.StockExchange = new(PartyIdentification10Choice)
	return o.StockExchange
}

func (o *OtherParties3) AddTradeRegulator() *PartyIdentification10Choice {
	o.TradeRegulator = new(PartyIdentification10Choice)
	return o.TradeRegulator
}

func (o *OtherParties3) AddTripartyAgent() *PartyIdentification10Choice {
	o.TripartyAgent = new(PartyIdentification10Choice)
	return o.TripartyAgent
}

// Other parties information.
type OtherParties30 struct {

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification110 `xml:"Invstr,omitempty"`

	// Identification of the stock exchange to which transaction reporting will be done.
	StockExchange *PartyIdentification111 `xml:"StockXchg,omitempty"`

	// Institution to which a trade must be reported.
	TradeRegulator *PartyIdentification111 `xml:"TradRgltr,omitempty"`
}

func (o *OtherParties30) AddInvestor() *PartyIdentification110 {
	o.Investor = new(PartyIdentification110)
	return o.Investor
}

func (o *OtherParties30) AddStockExchange() *PartyIdentification111 {
	o.StockExchange = new(PartyIdentification111)
	return o.StockExchange
}

func (o *OtherParties30) AddTradeRegulator() *PartyIdentification111 {
	o.TradeRegulator = new(PartyIdentification111)
	return o.TradeRegulator
}

// Other parties information.
type OtherParties31 struct {

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification110 `xml:"Invstr,omitempty"`

	// Foreign Financial Institution which has been authorised by local authorities to act as account management institution in the country.
	QualifiedForeignIntermediary *PartyIdentification111 `xml:"QlfdFrgnIntrmy,omitempty"`

	// Identification of the stock exchange to which transaction reporting will be done.
	StockExchange *PartyIdentification111 `xml:"StockXchg,omitempty"`

	// Institution to which a trade must be reported.
	TradeRegulator *PartyIdentification111 `xml:"TradRgltr,omitempty"`

	// Party responsible for the administration of a tri-party collateral transaction including collateral allocation, marking to market and substitution of collateral.
	TripartyAgent *PartyIdentification111 `xml:"TrptyAgt,omitempty"`
}

func (o *OtherParties31) AddInvestor() *PartyIdentification110 {
	o.Investor = new(PartyIdentification110)
	return o.Investor
}

func (o *OtherParties31) AddQualifiedForeignIntermediary() *PartyIdentification111 {
	o.QualifiedForeignIntermediary = new(PartyIdentification111)
	return o.QualifiedForeignIntermediary
}

func (o *OtherParties31) AddStockExchange() *PartyIdentification111 {
	o.StockExchange = new(PartyIdentification111)
	return o.StockExchange
}

func (o *OtherParties31) AddTradeRegulator() *PartyIdentification111 {
	o.TradeRegulator = new(PartyIdentification111)
	return o.TradeRegulator
}

func (o *OtherParties31) AddTripartyAgent() *PartyIdentification111 {
	o.TripartyAgent = new(PartyIdentification111)
	return o.TripartyAgent
}

// Other parties information.
type OtherParties4 struct {

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification14Choice `xml:"Invstr,omitempty"`

	// Identification of the stock exchange to which transaction reporting will be done.
	StockExchange *PartyIdentification10Choice `xml:"StockXchg,omitempty"`

	// Institution to which a trade must be reported.
	TradeRegulator *PartyIdentification10Choice `xml:"TradRgltr,omitempty"`
}

func (o *OtherParties4) AddInvestor() *PartyIdentification14Choice {
	o.Investor = new(PartyIdentification14Choice)
	return o.Investor
}

func (o *OtherParties4) AddStockExchange() *PartyIdentification10Choice {
	o.StockExchange = new(PartyIdentification10Choice)
	return o.StockExchange
}

func (o *OtherParties4) AddTradeRegulator() *PartyIdentification10Choice {
	o.TradeRegulator = new(PartyIdentification10Choice)
	return o.TradeRegulator
}

// Other parties information.
type OtherParties8 struct {

	// Party, either an individual or organisation, whose assets are being invested.
	Investor []*PartyIdentificationAndAccount40 `xml:"Invstr,omitempty"`

	// Foreign Financial Institution which has been authorised by local authorities to act as account management institution in the country.
	QualifiedForeignIntermediary *PartyIdentificationAndAccount41 `xml:"QlfdFrgnIntrmy,omitempty"`

	// Identification of the stock exchange to which transaction reporting will be done.
	StockExchange *PartyIdentificationAndAccount41 `xml:"StockXchg,omitempty"`

	// Institution to which a trade must be reported.
	TradeRegulator *PartyIdentificationAndAccount41 `xml:"TradRgltr,omitempty"`

	// Party responsible for the administration of a tri-party collateral transaction including collateral allocation, marking to market and substitution of collateral.
	TripartyAgent *PartyIdentificationAndAccount41 `xml:"TrptyAgt,omitempty"`

	// Party that identifies a broker when required (for example, authorised broker, prime broker, etc).
	Broker *PartyIdentificationAndAccount41 `xml:"Brkr,omitempty"`
}

func (o *OtherParties8) AddInvestor() *PartyIdentificationAndAccount40 {
	newValue := new(PartyIdentificationAndAccount40)
	o.Investor = append(o.Investor, newValue)
	return newValue
}

func (o *OtherParties8) AddQualifiedForeignIntermediary() *PartyIdentificationAndAccount41 {
	o.QualifiedForeignIntermediary = new(PartyIdentificationAndAccount41)
	return o.QualifiedForeignIntermediary
}

func (o *OtherParties8) AddStockExchange() *PartyIdentificationAndAccount41 {
	o.StockExchange = new(PartyIdentificationAndAccount41)
	return o.StockExchange
}

func (o *OtherParties8) AddTradeRegulator() *PartyIdentificationAndAccount41 {
	o.TradeRegulator = new(PartyIdentificationAndAccount41)
	return o.TradeRegulator
}

func (o *OtherParties8) AddTripartyAgent() *PartyIdentificationAndAccount41 {
	o.TripartyAgent = new(PartyIdentificationAndAccount41)
	return o.TripartyAgent
}

func (o *OtherParties8) AddBroker() *PartyIdentificationAndAccount41 {
	o.Broker = new(PartyIdentificationAndAccount41)
	return o.Broker
}

// Other parties information.
type OtherParties9 struct {

	// Party, either an individual or organisation, whose assets are being invested.
	Investor []*PartyIdentificationAndAccount36 `xml:"Invstr,omitempty"`

	// Foreign Financial Institution which has been authorised by local authorities to act as account management institution in the country.
	QualifiedForeignIntermediary *PartyIdentificationAndAccount37 `xml:"QlfdFrgnIntrmy,omitempty"`

	// Identification of the stock exchange to which transaction reporting will be done.
	StockExchange *PartyIdentificationAndAccount37 `xml:"StockXchg,omitempty"`

	// Institution to which a trade must be reported.
	TradeRegulator *PartyIdentificationAndAccount37 `xml:"TradRgltr,omitempty"`

	// Party responsible for the administration of a tri-party collateral transaction including collateral allocation, marking to market and substitution of collateral.
	TripartyAgent *PartyIdentificationAndAccount37 `xml:"TrptyAgt,omitempty"`
}

func (o *OtherParties9) AddInvestor() *PartyIdentificationAndAccount36 {
	newValue := new(PartyIdentificationAndAccount36)
	o.Investor = append(o.Investor, newValue)
	return newValue
}

func (o *OtherParties9) AddQualifiedForeignIntermediary() *PartyIdentificationAndAccount37 {
	o.QualifiedForeignIntermediary = new(PartyIdentificationAndAccount37)
	return o.QualifiedForeignIntermediary
}

func (o *OtherParties9) AddStockExchange() *PartyIdentificationAndAccount37 {
	o.StockExchange = new(PartyIdentificationAndAccount37)
	return o.StockExchange
}

func (o *OtherParties9) AddTradeRegulator() *PartyIdentificationAndAccount37 {
	o.TradeRegulator = new(PartyIdentificationAndAccount37)
	return o.TradeRegulator
}

func (o *OtherParties9) AddTripartyAgent() *PartyIdentificationAndAccount37 {
	o.TripartyAgent = new(PartyIdentificationAndAccount37)
	return o.TripartyAgent
}
