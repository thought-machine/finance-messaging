package iso20022

// Details of breakdown of a quantity.
type QuantityBreakdown11 struct {

	// Identification, for tax purposes, of a lot of identical securities that are bought at a certain date and at a certain price.
	LotNumber *GenericIdentification37 `xml:"LotNb,omitempty"`

	// Quantity of financial instruments that is part of the lot described.
	LotQuantity *FinancialInstrumentQuantity1Choice `xml:"LotQty,omitempty"`

	// Date/time at which the lot was purchased.
	LotDateTime *DateAndDateTime1Choice `xml:"LotDtTm,omitempty"`

	// Price at which the lot was purchased.
	LotPrice *Price4 `xml:"LotPric,omitempty"`
}

func (q *QuantityBreakdown11) AddLotNumber() *GenericIdentification37 {
	q.LotNumber = new(GenericIdentification37)
	return q.LotNumber
}

func (q *QuantityBreakdown11) AddLotQuantity() *FinancialInstrumentQuantity1Choice {
	q.LotQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.LotQuantity
}

func (q *QuantityBreakdown11) AddLotDateTime() *DateAndDateTime1Choice {
	q.LotDateTime = new(DateAndDateTime1Choice)
	return q.LotDateTime
}

func (q *QuantityBreakdown11) AddLotPrice() *Price4 {
	q.LotPrice = new(Price4)
	return q.LotPrice
}

// Details of breakdown of a quantity.
type QuantityBreakdown12 struct {

	// Identification, for tax purposes, of a lot of identical securities that are bought at a certain date and at a certain price.
	LotNumber *GenericIdentification37 `xml:"LotNb,omitempty"`

	// Quantity of financial instruments that is part of the lot described.
	LotQuantity *FinancialInstrumentQuantity1Choice `xml:"LotQty,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification20 `xml:"SctiesSubBalTp,omitempty"`
}

func (q *QuantityBreakdown12) AddLotNumber() *GenericIdentification37 {
	q.LotNumber = new(GenericIdentification37)
	return q.LotNumber
}

func (q *QuantityBreakdown12) AddLotQuantity() *FinancialInstrumentQuantity1Choice {
	q.LotQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.LotQuantity
}

func (q *QuantityBreakdown12) AddSecuritiesSubBalanceType() *GenericIdentification20 {
	q.SecuritiesSubBalanceType = new(GenericIdentification20)
	return q.SecuritiesSubBalanceType
}

// Details of breakdown of a quantity.
type QuantityBreakdown13 struct {

	// Identification, for tax purposes, of a lot of identical securities that are bought at a certain date and at a certain price.
	LotNumber *GenericIdentification37 `xml:"LotNb,omitempty"`

	// Quantity of financial instruments that is part of the lot described.
	LotQuantity *FinancialInstrumentQuantity1Choice `xml:"LotQty,omitempty"`

	// Date/time at which the lot was purchased.
	LotDateTime *DateAndDateTimeChoice `xml:"LotDtTm,omitempty"`

	// Price at which the lot was purchased.
	LotPrice *Price2 `xml:"LotPric,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice3Choice `xml:"TpOfPric,omitempty"`
}

func (q *QuantityBreakdown13) AddLotNumber() *GenericIdentification37 {
	q.LotNumber = new(GenericIdentification37)
	return q.LotNumber
}

func (q *QuantityBreakdown13) AddLotQuantity() *FinancialInstrumentQuantity1Choice {
	q.LotQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.LotQuantity
}

func (q *QuantityBreakdown13) AddLotDateTime() *DateAndDateTimeChoice {
	q.LotDateTime = new(DateAndDateTimeChoice)
	return q.LotDateTime
}

func (q *QuantityBreakdown13) AddLotPrice() *Price2 {
	q.LotPrice = new(Price2)
	return q.LotPrice
}

func (q *QuantityBreakdown13) AddTypeOfPrice() *TypeOfPrice3Choice {
	q.TypeOfPrice = new(TypeOfPrice3Choice)
	return q.TypeOfPrice
}

// Details of breakdown of a quantity.
type QuantityBreakdown14 struct {

	// Identification, for tax purposes, of a lot of identical securities that are bought at a certain date and at a certain price.
	LotNumber *GenericIdentification37 `xml:"LotNb,omitempty"`

	// Quantity of financial instruments that is part of the lot described.
	LotQuantity *FinancialInstrumentQuantity1Choice `xml:"LotQty,omitempty"`

	// Date/time at which the lot was purchased.
	LotDateTime *DateAndDateTimeChoice `xml:"LotDtTm,omitempty"`

	// Price at which the lot was purchased.
	LotPrice *Price2 `xml:"LotPric,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice3Choice `xml:"TpOfPric,omitempty"`

	// Valuation amounts for the lot provided in the base currency of the account.
	AccountBaseCurrencyAmounts *BalanceAmounts2 `xml:"AcctBaseCcyAmts,omitempty"`

	// Valuation amounts for the lot provided in the currency of the financial instrument.
	InstrumentCurrencyAmounts *BalanceAmounts2 `xml:"InstrmCcyAmts,omitempty"`

	// Valuation amounts for the lot provided in another currency than the base currency of the account.
	AlternateReportingCurrencyAmounts *BalanceAmounts2 `xml:"AltrnRptgCcyAmts,omitempty"`
}

func (q *QuantityBreakdown14) AddLotNumber() *GenericIdentification37 {
	q.LotNumber = new(GenericIdentification37)
	return q.LotNumber
}

func (q *QuantityBreakdown14) AddLotQuantity() *FinancialInstrumentQuantity1Choice {
	q.LotQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.LotQuantity
}

func (q *QuantityBreakdown14) AddLotDateTime() *DateAndDateTimeChoice {
	q.LotDateTime = new(DateAndDateTimeChoice)
	return q.LotDateTime
}

func (q *QuantityBreakdown14) AddLotPrice() *Price2 {
	q.LotPrice = new(Price2)
	return q.LotPrice
}

func (q *QuantityBreakdown14) AddTypeOfPrice() *TypeOfPrice3Choice {
	q.TypeOfPrice = new(TypeOfPrice3Choice)
	return q.TypeOfPrice
}

func (q *QuantityBreakdown14) AddAccountBaseCurrencyAmounts() *BalanceAmounts2 {
	q.AccountBaseCurrencyAmounts = new(BalanceAmounts2)
	return q.AccountBaseCurrencyAmounts
}

func (q *QuantityBreakdown14) AddInstrumentCurrencyAmounts() *BalanceAmounts2 {
	q.InstrumentCurrencyAmounts = new(BalanceAmounts2)
	return q.InstrumentCurrencyAmounts
}

func (q *QuantityBreakdown14) AddAlternateReportingCurrencyAmounts() *BalanceAmounts2 {
	q.AlternateReportingCurrencyAmounts = new(BalanceAmounts2)
	return q.AlternateReportingCurrencyAmounts
}

// Details of breakdown of a quantity.
type QuantityBreakdown15 struct {

	// Identification, for tax purposes, of a lot of identical securities that are bought at a certain date and at a certain price.
	LotNumber *GenericIdentification37 `xml:"LotNb"`

	// Quantity of financial instruments that is part of the lot described.
	LotQuantity *FinancialInstrumentQuantity1Choice `xml:"LotQty,omitempty"`
}

func (q *QuantityBreakdown15) AddLotNumber() *GenericIdentification37 {
	q.LotNumber = new(GenericIdentification37)
	return q.LotNumber
}

func (q *QuantityBreakdown15) AddLotQuantity() *FinancialInstrumentQuantity1Choice {
	q.LotQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.LotQuantity
}

// Details of breakdown of a quantity.
type QuantityBreakdown16 struct {

	// Identification, for tax purposes, of a lot of identical securities that are bought at a certain date and at a certain price.
	LotNumber *GenericIdentification37 `xml:"LotNb,omitempty"`

	// Quantity of financial instruments that is part of the lot described.
	LotQuantity *FinancialInstrumentQuantity1Choice `xml:"LotQty,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification20 `xml:"SctiesSubBalTp,omitempty"`

	// Date/time at which the lot was purchased.
	LotDateTime *DateAndDateTimeChoice `xml:"LotDtTm,omitempty"`

	// Price at which the lot was purchased.
	LotPrice *Price2 `xml:"LotPric,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice3Choice `xml:"TpOfPric,omitempty"`
}

func (q *QuantityBreakdown16) AddLotNumber() *GenericIdentification37 {
	q.LotNumber = new(GenericIdentification37)
	return q.LotNumber
}

func (q *QuantityBreakdown16) AddLotQuantity() *FinancialInstrumentQuantity1Choice {
	q.LotQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.LotQuantity
}

func (q *QuantityBreakdown16) AddSecuritiesSubBalanceType() *GenericIdentification20 {
	q.SecuritiesSubBalanceType = new(GenericIdentification20)
	return q.SecuritiesSubBalanceType
}

func (q *QuantityBreakdown16) AddLotDateTime() *DateAndDateTimeChoice {
	q.LotDateTime = new(DateAndDateTimeChoice)
	return q.LotDateTime
}

func (q *QuantityBreakdown16) AddLotPrice() *Price2 {
	q.LotPrice = new(Price2)
	return q.LotPrice
}

func (q *QuantityBreakdown16) AddTypeOfPrice() *TypeOfPrice3Choice {
	q.TypeOfPrice = new(TypeOfPrice3Choice)
	return q.TypeOfPrice
}

// Details of breakdown of a quantity.
type QuantityBreakdown23 struct {

	// Identification, for tax purposes, of a lot of identical securities that are bought at a certain date and at a certain price.
	LotNumber *GenericIdentification37 `xml:"LotNb,omitempty"`

	// Quantity of financial instruments that is part of the lot described.
	LotQuantity *Balance4 `xml:"LotQty,omitempty"`

	// Date/time at which the lot was purchased.
	LotDateTime *DateAndDateTimeChoice `xml:"LotDtTm,omitempty"`

	// Price at which the lot was purchased.
	LotPrice *Price2 `xml:"LotPric,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice3Choice `xml:"TpOfPric,omitempty"`
}

func (q *QuantityBreakdown23) AddLotNumber() *GenericIdentification37 {
	q.LotNumber = new(GenericIdentification37)
	return q.LotNumber
}

func (q *QuantityBreakdown23) AddLotQuantity() *Balance4 {
	q.LotQuantity = new(Balance4)
	return q.LotQuantity
}

func (q *QuantityBreakdown23) AddLotDateTime() *DateAndDateTimeChoice {
	q.LotDateTime = new(DateAndDateTimeChoice)
	return q.LotDateTime
}

func (q *QuantityBreakdown23) AddLotPrice() *Price2 {
	q.LotPrice = new(Price2)
	return q.LotPrice
}

func (q *QuantityBreakdown23) AddTypeOfPrice() *TypeOfPrice3Choice {
	q.TypeOfPrice = new(TypeOfPrice3Choice)
	return q.TypeOfPrice
}

// Details of breakdown of a quantity.
type QuantityBreakdown24 struct {

	// Identification, for tax purposes, of a lot of identical securities that are bought at a certain date and at a certain price.
	LotNumber *GenericIdentification37 `xml:"LotNb,omitempty"`

	// Quantity of financial instruments that is part of the lot described.
	LotQuantity *Balance4 `xml:"LotQty,omitempty"`

	// Date/time at which the lot was purchased.
	LotDateTime *DateAndDateTimeChoice `xml:"LotDtTm,omitempty"`

	// Price at which the lot was purchased.
	LotPrice *Price2 `xml:"LotPric,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice3Choice `xml:"TpOfPric,omitempty"`

	// Valuation amounts for the lot provided in the base currency of the account.
	AccountBaseCurrencyAmounts *BalanceAmounts2 `xml:"AcctBaseCcyAmts,omitempty"`

	// Valuation amounts for the lot provided in the currency of the financial instrument.
	InstrumentCurrencyAmounts *BalanceAmounts2 `xml:"InstrmCcyAmts,omitempty"`

	// Valuation amounts for the lot provided in another currency than the base currency of the account.
	AlternateReportingCurrencyAmounts *BalanceAmounts2 `xml:"AltrnRptgCcyAmts,omitempty"`
}

func (q *QuantityBreakdown24) AddLotNumber() *GenericIdentification37 {
	q.LotNumber = new(GenericIdentification37)
	return q.LotNumber
}

func (q *QuantityBreakdown24) AddLotQuantity() *Balance4 {
	q.LotQuantity = new(Balance4)
	return q.LotQuantity
}

func (q *QuantityBreakdown24) AddLotDateTime() *DateAndDateTimeChoice {
	q.LotDateTime = new(DateAndDateTimeChoice)
	return q.LotDateTime
}

func (q *QuantityBreakdown24) AddLotPrice() *Price2 {
	q.LotPrice = new(Price2)
	return q.LotPrice
}

func (q *QuantityBreakdown24) AddTypeOfPrice() *TypeOfPrice3Choice {
	q.TypeOfPrice = new(TypeOfPrice3Choice)
	return q.TypeOfPrice
}

func (q *QuantityBreakdown24) AddAccountBaseCurrencyAmounts() *BalanceAmounts2 {
	q.AccountBaseCurrencyAmounts = new(BalanceAmounts2)
	return q.AccountBaseCurrencyAmounts
}

func (q *QuantityBreakdown24) AddInstrumentCurrencyAmounts() *BalanceAmounts2 {
	q.InstrumentCurrencyAmounts = new(BalanceAmounts2)
	return q.InstrumentCurrencyAmounts
}

func (q *QuantityBreakdown24) AddAlternateReportingCurrencyAmounts() *BalanceAmounts2 {
	q.AlternateReportingCurrencyAmounts = new(BalanceAmounts2)
	return q.AlternateReportingCurrencyAmounts
}

// Details of breakdown of a quantity.
type QuantityBreakdown27 struct {

	// Identification, for tax purposes, of a lot of identical securities that are bought at a certain date and at a certain price.
	LotNumber *GenericIdentification37 `xml:"LotNb,omitempty"`

	// Quantity of financial instruments that is part of the lot described.
	LotQuantity *Balance7 `xml:"LotQty,omitempty"`

	// Date/time at which the lot was purchased.
	LotDateTime *DateAndDateTimeChoice `xml:"LotDtTm,omitempty"`

	// Price at which the lot was purchased.
	LotPrice *Price2 `xml:"LotPric,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice29Choice `xml:"TpOfPric,omitempty"`
}

func (q *QuantityBreakdown27) AddLotNumber() *GenericIdentification37 {
	q.LotNumber = new(GenericIdentification37)
	return q.LotNumber
}

func (q *QuantityBreakdown27) AddLotQuantity() *Balance7 {
	q.LotQuantity = new(Balance7)
	return q.LotQuantity
}

func (q *QuantityBreakdown27) AddLotDateTime() *DateAndDateTimeChoice {
	q.LotDateTime = new(DateAndDateTimeChoice)
	return q.LotDateTime
}

func (q *QuantityBreakdown27) AddLotPrice() *Price2 {
	q.LotPrice = new(Price2)
	return q.LotPrice
}

func (q *QuantityBreakdown27) AddTypeOfPrice() *TypeOfPrice29Choice {
	q.TypeOfPrice = new(TypeOfPrice29Choice)
	return q.TypeOfPrice
}

// Details of breakdown of a quantity.
type QuantityBreakdown28 struct {

	// Identification, for tax purposes, of a lot of identical securities that are bought at a certain date and at a certain price.
	LotNumber *GenericIdentification37 `xml:"LotNb,omitempty"`

	// Quantity of financial instruments that is part of the lot described.
	LotQuantity *Balance7 `xml:"LotQty,omitempty"`

	// Date/time at which the lot was purchased.
	LotDateTime *DateAndDateTimeChoice `xml:"LotDtTm,omitempty"`

	// Price at which the lot was purchased.
	LotPrice *Price2 `xml:"LotPric,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice29Choice `xml:"TpOfPric,omitempty"`

	// Valuation amounts for the lot provided in the base currency of the account.
	AccountBaseCurrencyAmounts *BalanceAmounts2 `xml:"AcctBaseCcyAmts,omitempty"`

	// Valuation amounts for the lot provided in the currency of the financial instrument.
	InstrumentCurrencyAmounts *BalanceAmounts2 `xml:"InstrmCcyAmts,omitempty"`

	// Valuation amounts for the lot provided in another currency than the base currency of the account.
	AlternateReportingCurrencyAmounts *BalanceAmounts2 `xml:"AltrnRptgCcyAmts,omitempty"`
}

func (q *QuantityBreakdown28) AddLotNumber() *GenericIdentification37 {
	q.LotNumber = new(GenericIdentification37)
	return q.LotNumber
}

func (q *QuantityBreakdown28) AddLotQuantity() *Balance7 {
	q.LotQuantity = new(Balance7)
	return q.LotQuantity
}

func (q *QuantityBreakdown28) AddLotDateTime() *DateAndDateTimeChoice {
	q.LotDateTime = new(DateAndDateTimeChoice)
	return q.LotDateTime
}

func (q *QuantityBreakdown28) AddLotPrice() *Price2 {
	q.LotPrice = new(Price2)
	return q.LotPrice
}

func (q *QuantityBreakdown28) AddTypeOfPrice() *TypeOfPrice29Choice {
	q.TypeOfPrice = new(TypeOfPrice29Choice)
	return q.TypeOfPrice
}

func (q *QuantityBreakdown28) AddAccountBaseCurrencyAmounts() *BalanceAmounts2 {
	q.AccountBaseCurrencyAmounts = new(BalanceAmounts2)
	return q.AccountBaseCurrencyAmounts
}

func (q *QuantityBreakdown28) AddInstrumentCurrencyAmounts() *BalanceAmounts2 {
	q.InstrumentCurrencyAmounts = new(BalanceAmounts2)
	return q.InstrumentCurrencyAmounts
}

func (q *QuantityBreakdown28) AddAlternateReportingCurrencyAmounts() *BalanceAmounts2 {
	q.AlternateReportingCurrencyAmounts = new(BalanceAmounts2)
	return q.AlternateReportingCurrencyAmounts
}

// Details of breakdown of a quantity.
type QuantityBreakdown29 struct {

	// Identification, for tax purposes, of a lot of identical securities that are bought at a certain date and at a certain price.
	LotNumber *GenericIdentification37 `xml:"LotNb,omitempty"`

	// Quantity of financial instruments that is part of the lot described.
	LotQuantity *FinancialInstrumentQuantity1Choice `xml:"LotQty,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification30 `xml:"SctiesSubBalTp,omitempty"`

	// Date/time at which the lot was purchased.
	LotDateTime *DateAndDateTimeChoice `xml:"LotDtTm,omitempty"`

	// Price at which the lot was purchased.
	LotPrice *Price2 `xml:"LotPric,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice29Choice `xml:"TpOfPric,omitempty"`
}

func (q *QuantityBreakdown29) AddLotNumber() *GenericIdentification37 {
	q.LotNumber = new(GenericIdentification37)
	return q.LotNumber
}

func (q *QuantityBreakdown29) AddLotQuantity() *FinancialInstrumentQuantity1Choice {
	q.LotQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.LotQuantity
}

func (q *QuantityBreakdown29) AddSecuritiesSubBalanceType() *GenericIdentification30 {
	q.SecuritiesSubBalanceType = new(GenericIdentification30)
	return q.SecuritiesSubBalanceType
}

func (q *QuantityBreakdown29) AddLotDateTime() *DateAndDateTimeChoice {
	q.LotDateTime = new(DateAndDateTimeChoice)
	return q.LotDateTime
}

func (q *QuantityBreakdown29) AddLotPrice() *Price2 {
	q.LotPrice = new(Price2)
	return q.LotPrice
}

func (q *QuantityBreakdown29) AddTypeOfPrice() *TypeOfPrice29Choice {
	q.TypeOfPrice = new(TypeOfPrice29Choice)
	return q.TypeOfPrice
}

// Details of breakdown of a quantity.
type QuantityBreakdown3 struct {

	// Identification, for tax purposes, of a lot of identical securities that are bought at a certain date and at a certain price.
	LotNumber *Number2Choice `xml:"LotNb,omitempty"`

	// Quantity of financial instruments that is part of the lot described.
	LotQuantity *FinancialInstrumentQuantity1Choice `xml:"LotQty,omitempty"`

	// Date/time at which the lot was purchased.
	LotDateTime *DateAndDateTimeChoice `xml:"LotDtTm,omitempty"`

	// Price at which the lot was purchased.
	LotPrice *Price2 `xml:"LotPric,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice3Choice `xml:"TpOfPric,omitempty"`
}

func (q *QuantityBreakdown3) AddLotNumber() *Number2Choice {
	q.LotNumber = new(Number2Choice)
	return q.LotNumber
}

func (q *QuantityBreakdown3) AddLotQuantity() *FinancialInstrumentQuantity1Choice {
	q.LotQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.LotQuantity
}

func (q *QuantityBreakdown3) AddLotDateTime() *DateAndDateTimeChoice {
	q.LotDateTime = new(DateAndDateTimeChoice)
	return q.LotDateTime
}

func (q *QuantityBreakdown3) AddLotPrice() *Price2 {
	q.LotPrice = new(Price2)
	return q.LotPrice
}

func (q *QuantityBreakdown3) AddTypeOfPrice() *TypeOfPrice3Choice {
	q.TypeOfPrice = new(TypeOfPrice3Choice)
	return q.TypeOfPrice
}

// Details of breakdown of a quantity.
type QuantityBreakdown30 struct {

	// Identification, for tax purposes, of a lot of identical securities that are bought at a certain date and at a certain price.
	LotNumber *GenericIdentification37 `xml:"LotNb,omitempty"`

	// Quantity of financial instruments that is part of the lot described.
	LotQuantity *FinancialInstrumentQuantity1Choice `xml:"LotQty,omitempty"`

	// Date/time at which the lot was purchased.
	LotDateTime *DateAndDateTimeChoice `xml:"LotDtTm,omitempty"`

	// Price at which the lot was purchased.
	LotPrice *Price2 `xml:"LotPric,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice29Choice `xml:"TpOfPric,omitempty"`
}

func (q *QuantityBreakdown30) AddLotNumber() *GenericIdentification37 {
	q.LotNumber = new(GenericIdentification37)
	return q.LotNumber
}

func (q *QuantityBreakdown30) AddLotQuantity() *FinancialInstrumentQuantity1Choice {
	q.LotQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.LotQuantity
}

func (q *QuantityBreakdown30) AddLotDateTime() *DateAndDateTimeChoice {
	q.LotDateTime = new(DateAndDateTimeChoice)
	return q.LotDateTime
}

func (q *QuantityBreakdown30) AddLotPrice() *Price2 {
	q.LotPrice = new(Price2)
	return q.LotPrice
}

func (q *QuantityBreakdown30) AddTypeOfPrice() *TypeOfPrice29Choice {
	q.TypeOfPrice = new(TypeOfPrice29Choice)
	return q.TypeOfPrice
}

// Details of breakdown of a quantity.
type QuantityBreakdown31 struct {

	// Identification, for tax purposes, of a lot of identical securities that are bought at a certain date and at a certain price.
	LotNumber *GenericIdentification37 `xml:"LotNb"`

	// Quantity of financial instruments that is part of the lot described.
	LotQuantity *FinancialInstrumentQuantity1Choice `xml:"LotQty,omitempty"`
}

func (q *QuantityBreakdown31) AddLotNumber() *GenericIdentification37 {
	q.LotNumber = new(GenericIdentification37)
	return q.LotNumber
}

func (q *QuantityBreakdown31) AddLotQuantity() *FinancialInstrumentQuantity1Choice {
	q.LotQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.LotQuantity
}

// Details of breakdown of a quantity.
type QuantityBreakdown32 struct {

	// Identification, for tax purposes, of a lot of identical securities that are bought at a certain date and at a certain price.
	LotNumber *GenericIdentification37 `xml:"LotNb,omitempty"`

	// Quantity of financial instruments that is part of the lot described.
	LotQuantity *FinancialInstrumentQuantity1Choice `xml:"LotQty,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification30 `xml:"SctiesSubBalTp,omitempty"`
}

func (q *QuantityBreakdown32) AddLotNumber() *GenericIdentification37 {
	q.LotNumber = new(GenericIdentification37)
	return q.LotNumber
}

func (q *QuantityBreakdown32) AddLotQuantity() *FinancialInstrumentQuantity1Choice {
	q.LotQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.LotQuantity
}

func (q *QuantityBreakdown32) AddSecuritiesSubBalanceType() *GenericIdentification30 {
	q.SecuritiesSubBalanceType = new(GenericIdentification30)
	return q.SecuritiesSubBalanceType
}

// Details of breakdown of a quantity.
type QuantityBreakdown33 struct {

	// Identification, for tax purposes, of a lot of identical securities that are bought at a certain date and at a certain price.
	LotNumber *GenericIdentification39 `xml:"LotNb"`

	// Quantity of financial instruments that is part of the lot described.
	LotQuantity *FinancialInstrumentQuantity15Choice `xml:"LotQty,omitempty"`
}

func (q *QuantityBreakdown33) AddLotNumber() *GenericIdentification39 {
	q.LotNumber = new(GenericIdentification39)
	return q.LotNumber
}

func (q *QuantityBreakdown33) AddLotQuantity() *FinancialInstrumentQuantity15Choice {
	q.LotQuantity = new(FinancialInstrumentQuantity15Choice)
	return q.LotQuantity
}

// Details of breakdown of a quantity.
type QuantityBreakdown34 struct {

	// Identification, for tax purposes, of a lot of identical securities that are bought at a certain date and at a certain price.
	LotNumber *GenericIdentification39 `xml:"LotNb,omitempty"`

	// Quantity of financial instruments that is part of the lot described.
	LotQuantity *FinancialInstrumentQuantity1Choice `xml:"LotQty,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification47 `xml:"SctiesSubBalTp,omitempty"`
}

func (q *QuantityBreakdown34) AddLotNumber() *GenericIdentification39 {
	q.LotNumber = new(GenericIdentification39)
	return q.LotNumber
}

func (q *QuantityBreakdown34) AddLotQuantity() *FinancialInstrumentQuantity1Choice {
	q.LotQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.LotQuantity
}

func (q *QuantityBreakdown34) AddSecuritiesSubBalanceType() *GenericIdentification47 {
	q.SecuritiesSubBalanceType = new(GenericIdentification47)
	return q.SecuritiesSubBalanceType
}

// Details of breakdown of a quantity.
type QuantityBreakdown38 struct {

	// Identification, for tax purposes, of a lot of identical securities that are bought at a certain date and at a certain price.
	LotNumber *GenericIdentification39 `xml:"LotNb,omitempty"`

	// Quantity of financial instruments that is part of the lot described.
	LotQuantity *FinancialInstrumentQuantity15Choice `xml:"LotQty,omitempty"`

	// Date/time at which the lot was purchased.
	LotDateTime *DateAndDateTimeChoice `xml:"LotDtTm,omitempty"`

	// Price at which the lot was purchased.
	LotPrice *Price3 `xml:"LotPric,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice32Choice `xml:"TpOfPric,omitempty"`
}

func (q *QuantityBreakdown38) AddLotNumber() *GenericIdentification39 {
	q.LotNumber = new(GenericIdentification39)
	return q.LotNumber
}

func (q *QuantityBreakdown38) AddLotQuantity() *FinancialInstrumentQuantity15Choice {
	q.LotQuantity = new(FinancialInstrumentQuantity15Choice)
	return q.LotQuantity
}

func (q *QuantityBreakdown38) AddLotDateTime() *DateAndDateTimeChoice {
	q.LotDateTime = new(DateAndDateTimeChoice)
	return q.LotDateTime
}

func (q *QuantityBreakdown38) AddLotPrice() *Price3 {
	q.LotPrice = new(Price3)
	return q.LotPrice
}

func (q *QuantityBreakdown38) AddTypeOfPrice() *TypeOfPrice32Choice {
	q.TypeOfPrice = new(TypeOfPrice32Choice)
	return q.TypeOfPrice
}

// Details of breakdown of a quantity.
type QuantityBreakdown39 struct {

	// Identification, for tax purposes, of a lot of identical securities that are bought at a certain date and at a certain price.
	LotNumber *GenericIdentification39 `xml:"LotNb,omitempty"`

	// Quantity of financial instruments that is part of the lot described.
	LotQuantity *Balance11 `xml:"LotQty,omitempty"`

	// Date/time at which the lot was purchased.
	LotDateTime *DateAndDateTimeChoice `xml:"LotDtTm,omitempty"`

	// Price at which the lot was purchased.
	LotPrice *Price3 `xml:"LotPric,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice32Choice `xml:"TpOfPric,omitempty"`

	// Valuation amounts for the lot provided in the base currency of the account.
	AccountBaseCurrencyAmounts *BalanceAmounts6 `xml:"AcctBaseCcyAmts,omitempty"`

	// Valuation amounts for the lot provided in the currency of the financial instrument.
	InstrumentCurrencyAmounts *BalanceAmounts6 `xml:"InstrmCcyAmts,omitempty"`

	// Valuation amounts for the lot provided in another currency than the base currency of the account.
	AlternateReportingCurrencyAmounts *BalanceAmounts6 `xml:"AltrnRptgCcyAmts,omitempty"`
}

func (q *QuantityBreakdown39) AddLotNumber() *GenericIdentification39 {
	q.LotNumber = new(GenericIdentification39)
	return q.LotNumber
}

func (q *QuantityBreakdown39) AddLotQuantity() *Balance11 {
	q.LotQuantity = new(Balance11)
	return q.LotQuantity
}

func (q *QuantityBreakdown39) AddLotDateTime() *DateAndDateTimeChoice {
	q.LotDateTime = new(DateAndDateTimeChoice)
	return q.LotDateTime
}

func (q *QuantityBreakdown39) AddLotPrice() *Price3 {
	q.LotPrice = new(Price3)
	return q.LotPrice
}

func (q *QuantityBreakdown39) AddTypeOfPrice() *TypeOfPrice32Choice {
	q.TypeOfPrice = new(TypeOfPrice32Choice)
	return q.TypeOfPrice
}

func (q *QuantityBreakdown39) AddAccountBaseCurrencyAmounts() *BalanceAmounts6 {
	q.AccountBaseCurrencyAmounts = new(BalanceAmounts6)
	return q.AccountBaseCurrencyAmounts
}

func (q *QuantityBreakdown39) AddInstrumentCurrencyAmounts() *BalanceAmounts6 {
	q.InstrumentCurrencyAmounts = new(BalanceAmounts6)
	return q.InstrumentCurrencyAmounts
}

func (q *QuantityBreakdown39) AddAlternateReportingCurrencyAmounts() *BalanceAmounts6 {
	q.AlternateReportingCurrencyAmounts = new(BalanceAmounts6)
	return q.AlternateReportingCurrencyAmounts
}

// Details of breakdown of a quantity.
type QuantityBreakdown4 struct {

	// Identification, for tax purposes, of a lot of identical securities that are bought at a certain date and at a certain price.
	LotNumber *Number2Choice `xml:"LotNb,omitempty"`

	// Quantity of financial instruments that is part of the lot described.
	LotQuantity *FinancialInstrumentQuantity1Choice `xml:"LotQty,omitempty"`

	// Date/time at which the lot was purchased.
	LotDateTime *DateAndDateTimeChoice `xml:"LotDtTm,omitempty"`

	// Price at which the lot was purchased.
	LotPrice *Price2 `xml:"LotPric,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice3Choice `xml:"TpOfPric,omitempty"`

	// Valuation amounts for the lot provided in the base currency of the account.
	AccountBaseCurrencyAmounts *BalanceAmounts2 `xml:"AcctBaseCcyAmts,omitempty"`

	// Valuation amounts for the lot provided in the currency of the financial instrument.
	InstrumentCurrencyAmounts *BalanceAmounts2 `xml:"InstrmCcyAmts,omitempty"`

	// Valuation amounts for the lot provided in another currency than the base currency of the account.
	AlternateReportingCurrencyAmounts *BalanceAmounts2 `xml:"AltrnRptgCcyAmts,omitempty"`
}

func (q *QuantityBreakdown4) AddLotNumber() *Number2Choice {
	q.LotNumber = new(Number2Choice)
	return q.LotNumber
}

func (q *QuantityBreakdown4) AddLotQuantity() *FinancialInstrumentQuantity1Choice {
	q.LotQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.LotQuantity
}

func (q *QuantityBreakdown4) AddLotDateTime() *DateAndDateTimeChoice {
	q.LotDateTime = new(DateAndDateTimeChoice)
	return q.LotDateTime
}

func (q *QuantityBreakdown4) AddLotPrice() *Price2 {
	q.LotPrice = new(Price2)
	return q.LotPrice
}

func (q *QuantityBreakdown4) AddTypeOfPrice() *TypeOfPrice3Choice {
	q.TypeOfPrice = new(TypeOfPrice3Choice)
	return q.TypeOfPrice
}

func (q *QuantityBreakdown4) AddAccountBaseCurrencyAmounts() *BalanceAmounts2 {
	q.AccountBaseCurrencyAmounts = new(BalanceAmounts2)
	return q.AccountBaseCurrencyAmounts
}

func (q *QuantityBreakdown4) AddInstrumentCurrencyAmounts() *BalanceAmounts2 {
	q.InstrumentCurrencyAmounts = new(BalanceAmounts2)
	return q.InstrumentCurrencyAmounts
}

func (q *QuantityBreakdown4) AddAlternateReportingCurrencyAmounts() *BalanceAmounts2 {
	q.AlternateReportingCurrencyAmounts = new(BalanceAmounts2)
	return q.AlternateReportingCurrencyAmounts
}

// Details of breakdown of a quantity.
type QuantityBreakdown40 struct {

	// Identification, for tax purposes, of a lot of identical securities that are bought at a certain date and at a certain price.
	LotNumber *GenericIdentification39 `xml:"LotNb,omitempty"`

	// Quantity of financial instruments that is part of the lot described.
	LotQuantity *Balance11 `xml:"LotQty,omitempty"`

	// Date/time at which the lot was purchased.
	LotDateTime *DateAndDateTimeChoice `xml:"LotDtTm,omitempty"`

	// Price at which the lot was purchased.
	LotPrice *Price3 `xml:"LotPric,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice32Choice `xml:"TpOfPric,omitempty"`
}

func (q *QuantityBreakdown40) AddLotNumber() *GenericIdentification39 {
	q.LotNumber = new(GenericIdentification39)
	return q.LotNumber
}

func (q *QuantityBreakdown40) AddLotQuantity() *Balance11 {
	q.LotQuantity = new(Balance11)
	return q.LotQuantity
}

func (q *QuantityBreakdown40) AddLotDateTime() *DateAndDateTimeChoice {
	q.LotDateTime = new(DateAndDateTimeChoice)
	return q.LotDateTime
}

func (q *QuantityBreakdown40) AddLotPrice() *Price3 {
	q.LotPrice = new(Price3)
	return q.LotPrice
}

func (q *QuantityBreakdown40) AddTypeOfPrice() *TypeOfPrice32Choice {
	q.TypeOfPrice = new(TypeOfPrice32Choice)
	return q.TypeOfPrice
}

// Details of breakdown of a quantity.
type QuantityBreakdown44 struct {

	// Identification, for tax purposes, of a lot of identical securities that are bought at a certain date and at a certain price.
	LotNumber *GenericIdentification39 `xml:"LotNb,omitempty"`

	// Quantity of financial instruments that is part of the lot described.
	LotQuantity *FinancialInstrumentQuantity15Choice `xml:"LotQty,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification47 `xml:"SctiesSubBalTp,omitempty"`

	// Date/time at which the lot was purchased.
	LotDateTime *DateAndDateTimeChoice `xml:"LotDtTm,omitempty"`

	// Price at which the lot was purchased.
	LotPrice *Price3 `xml:"LotPric,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice32Choice `xml:"TpOfPric,omitempty"`
}

func (q *QuantityBreakdown44) AddLotNumber() *GenericIdentification39 {
	q.LotNumber = new(GenericIdentification39)
	return q.LotNumber
}

func (q *QuantityBreakdown44) AddLotQuantity() *FinancialInstrumentQuantity15Choice {
	q.LotQuantity = new(FinancialInstrumentQuantity15Choice)
	return q.LotQuantity
}

func (q *QuantityBreakdown44) AddSecuritiesSubBalanceType() *GenericIdentification47 {
	q.SecuritiesSubBalanceType = new(GenericIdentification47)
	return q.SecuritiesSubBalanceType
}

func (q *QuantityBreakdown44) AddLotDateTime() *DateAndDateTimeChoice {
	q.LotDateTime = new(DateAndDateTimeChoice)
	return q.LotDateTime
}

func (q *QuantityBreakdown44) AddLotPrice() *Price3 {
	q.LotPrice = new(Price3)
	return q.LotPrice
}

func (q *QuantityBreakdown44) AddTypeOfPrice() *TypeOfPrice32Choice {
	q.TypeOfPrice = new(TypeOfPrice32Choice)
	return q.TypeOfPrice
}

// Details of breakdown of a quantity.
type QuantityBreakdown5 struct {

	// Identification, for tax purposes, of a lot of identical securities that are bought at a certain date and at a certain price.
	LotNumber *Number2Choice `xml:"LotNb"`

	// Quantity of financial instruments that is part of the lot described.
	LotQuantity *FinancialInstrumentQuantity1Choice `xml:"LotQty,omitempty"`
}

func (q *QuantityBreakdown5) AddLotNumber() *Number2Choice {
	q.LotNumber = new(Number2Choice)
	return q.LotNumber
}

func (q *QuantityBreakdown5) AddLotQuantity() *FinancialInstrumentQuantity1Choice {
	q.LotQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.LotQuantity
}

// Details of breakdown of a quantity.
type QuantityBreakdown9 struct {

	// Identification, for tax purposes, of a lot of identical securities that are bought at a certain date and at a certain price.
	LotNumber *Number2Choice `xml:"LotNb,omitempty"`

	// Quantity of financial instruments that is part of the lot described.
	LotQuantity *FinancialInstrumentQuantity1Choice `xml:"LotQty,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification19 `xml:"SctiesSubBalTp,omitempty"`

	// Date/time at which the lot was purchased.
	LotDateTime *DateAndDateTimeChoice `xml:"LotDtTm,omitempty"`

	// Price at which the lot was purchased.
	LotPrice *Price2 `xml:"LotPric,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice3Choice `xml:"TpOfPric,omitempty"`
}

func (q *QuantityBreakdown9) AddLotNumber() *Number2Choice {
	q.LotNumber = new(Number2Choice)
	return q.LotNumber
}

func (q *QuantityBreakdown9) AddLotQuantity() *FinancialInstrumentQuantity1Choice {
	q.LotQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.LotQuantity
}

func (q *QuantityBreakdown9) AddSecuritiesSubBalanceType() *GenericIdentification19 {
	q.SecuritiesSubBalanceType = new(GenericIdentification19)
	return q.SecuritiesSubBalanceType
}

func (q *QuantityBreakdown9) AddLotDateTime() *DateAndDateTimeChoice {
	q.LotDateTime = new(DateAndDateTimeChoice)
	return q.LotDateTime
}

func (q *QuantityBreakdown9) AddLotPrice() *Price2 {
	q.LotPrice = new(Price2)
	return q.LotPrice
}

func (q *QuantityBreakdown9) AddTypeOfPrice() *TypeOfPrice3Choice {
	q.TypeOfPrice = new(TypeOfPrice3Choice)
	return q.TypeOfPrice
}
