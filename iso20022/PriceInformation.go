package iso20022

// Amount of money for which goods or services are offered, sold, or bought.
type PriceInformation1 struct {

	// Value of the price, eg, as a currency and value.
	Value *PriceRateOrAmountOrUnknownChoice `xml:"Val"`

	// Type of value in which the price is expressed.
	ValueType *PriceValueType2Code `xml:"ValTp,omitempty"`

	// Type and information about a price.
	Type *TypeOfPrice5Code `xml:"Tp"`

	// Place from which the price was obtained.
	SourceOfPrice *PriceSourceFormatChoice `xml:"SrcOfPric,omitempty"`

	// Date on which the price is obtained. With an investment fund, this is as stated in the prospectus.
	QuotationDate *DateAndDateTimeChoice `xml:"QtnDt,omitempty"`

	// Indicates whether the price is expressed as a yield. The absence of Yielded means it is not applicable.
	Yielded *YesNoIndicator `xml:"Yldd,omitempty"`
}

func (p *PriceInformation1) AddValue() *PriceRateOrAmountOrUnknownChoice {
	p.Value = new(PriceRateOrAmountOrUnknownChoice)
	return p.Value
}

func (p *PriceInformation1) SetValueType(value string) {
	p.ValueType = (*PriceValueType2Code)(&value)
}

func (p *PriceInformation1) SetType(value string) {
	p.Type = (*TypeOfPrice5Code)(&value)
}

func (p *PriceInformation1) AddSourceOfPrice() *PriceSourceFormatChoice {
	p.SourceOfPrice = new(PriceSourceFormatChoice)
	return p.SourceOfPrice
}

func (p *PriceInformation1) AddQuotationDate() *DateAndDateTimeChoice {
	p.QuotationDate = new(DateAndDateTimeChoice)
	return p.QuotationDate
}

func (p *PriceInformation1) SetYielded(value string) {
	p.Yielded = (*YesNoIndicator)(&value)
}

// Amount of money for which goods or services are offered, sold, or bought.
type PriceInformation10 struct {

	// Current value of the price, eg, as a currency and value.
	CurrentPrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"CurPric"`

	// Specifies the type of price and information about the price.
	Type *TypeOfPrice27Choice `xml:"Tp"`

	// Previous value of the price, eg, as a currency and value.
	PreviousPrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"PrvsPric,omitempty"`

	// Difference or change between the previous price value and the current price value.
	AmountOfChange *PriceValueAndRate4 `xml:"AmtOfChng,omitempty"`
}

func (p *PriceInformation10) SetCurrentPrice(value, currency string) {
	p.CurrentPrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (p *PriceInformation10) AddType() *TypeOfPrice27Choice {
	p.Type = new(TypeOfPrice27Choice)
	return p.Type
}

func (p *PriceInformation10) SetPreviousPrice(value, currency string) {
	p.PreviousPrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (p *PriceInformation10) AddAmountOfChange() *PriceValueAndRate4 {
	p.AmountOfChange = new(PriceValueAndRate4)
	return p.AmountOfChange
}

// Amount of money for which goods or services are offered, sold, or bought.
type PriceInformation11 struct {

	// Value of the price, for instance, as a currency and value.
	Value *Price4 `xml:"Val"`

	// Date on which the price is obtained. With an investment fund, this is as stated in the prospectus.
	QuotationDate *DateAndDateTime1Choice `xml:"QtnDt,omitempty"`

	// Period during which the price of a security is determined (For  outturn securities).
	PriceCalculationPeriod *DateTimePeriodChoice `xml:"PricClctnPrd,omitempty"`

	// Place from which the price was obtained.
	SourceOfPrice *MarketIdentification79 `xml:"SrcOfPric,omitempty"`
}

func (p *PriceInformation11) AddValue() *Price4 {
	p.Value = new(Price4)
	return p.Value
}

func (p *PriceInformation11) AddQuotationDate() *DateAndDateTime1Choice {
	p.QuotationDate = new(DateAndDateTime1Choice)
	return p.QuotationDate
}

func (p *PriceInformation11) AddPriceCalculationPeriod() *DateTimePeriodChoice {
	p.PriceCalculationPeriod = new(DateTimePeriodChoice)
	return p.PriceCalculationPeriod
}

func (p *PriceInformation11) AddSourceOfPrice() *MarketIdentification79 {
	p.SourceOfPrice = new(MarketIdentification79)
	return p.SourceOfPrice
}

// Amount of money for which goods or services are offered, sold, or bought.
type PriceInformation12 struct {

	// Specifies the type of price and information about the price.
	Type *TypeOfPrice28Choice `xml:"Tp"`

	// Value of the price, eg, as a currency and value.
	Value *PriceRateOrAmountOrUnknownChoice `xml:"Val"`

	// Type of value in which the price is expressed.
	ValueType *YieldedOrValueType1Choice `xml:"ValTp"`

	// Place from which the price was obtained.
	SourceOfPrice *MarketIdentification89 `xml:"SrcOfPric,omitempty"`

	// Date on which the price is obtained. With an investment fund, this is as stated in the prospectus.
	QuotationDate *DateAndDateTimeChoice `xml:"QtnDt,omitempty"`
}

func (p *PriceInformation12) AddType() *TypeOfPrice28Choice {
	p.Type = new(TypeOfPrice28Choice)
	return p.Type
}

func (p *PriceInformation12) AddValue() *PriceRateOrAmountOrUnknownChoice {
	p.Value = new(PriceRateOrAmountOrUnknownChoice)
	return p.Value
}

func (p *PriceInformation12) AddValueType() *YieldedOrValueType1Choice {
	p.ValueType = new(YieldedOrValueType1Choice)
	return p.ValueType
}

func (p *PriceInformation12) AddSourceOfPrice() *MarketIdentification89 {
	p.SourceOfPrice = new(MarketIdentification89)
	return p.SourceOfPrice
}

func (p *PriceInformation12) AddQuotationDate() *DateAndDateTimeChoice {
	p.QuotationDate = new(DateAndDateTimeChoice)
	return p.QuotationDate
}

// Amount of money for which goods or services are offered, sold, or bought.
type PriceInformation13 struct {

	// Specifies the type of price and information about the price.
	Type *TypeOfPrice30Choice `xml:"Tp"`

	// Type of value in which the price is expressed.
	ValueType *YieldedOrValueType1Choice `xml:"ValTp"`

	// Value of the price, eg, as a currency and value.
	Value *PriceRateOrAmountOrUnknownChoice `xml:"Val"`

	// Place from which the price was obtained.
	SourceOfPrice *MarketIdentification89 `xml:"SrcOfPric,omitempty"`

	// Date on which the price is obtained. With an investment fund, this is as stated in the prospectus.
	QuotationDate *DateAndDateTimeChoice `xml:"QtnDt,omitempty"`
}

func (p *PriceInformation13) AddType() *TypeOfPrice30Choice {
	p.Type = new(TypeOfPrice30Choice)
	return p.Type
}

func (p *PriceInformation13) AddValueType() *YieldedOrValueType1Choice {
	p.ValueType = new(YieldedOrValueType1Choice)
	return p.ValueType
}

func (p *PriceInformation13) AddValue() *PriceRateOrAmountOrUnknownChoice {
	p.Value = new(PriceRateOrAmountOrUnknownChoice)
	return p.Value
}

func (p *PriceInformation13) AddSourceOfPrice() *MarketIdentification89 {
	p.SourceOfPrice = new(MarketIdentification89)
	return p.SourceOfPrice
}

func (p *PriceInformation13) AddQuotationDate() *DateAndDateTimeChoice {
	p.QuotationDate = new(DateAndDateTimeChoice)
	return p.QuotationDate
}

// Amount of money for which goods or services are offered, sold, or bought.
type PriceInformation14 struct {

	// Specifies the type of price and information about the price.
	Type *TypeOfPrice33Choice `xml:"Tp"`

	// Value of the price, eg, as a currency and value.
	Value *PriceRateOrAmountOrUnknown1Choice `xml:"Val"`

	// Type of value in which the price is expressed.
	ValueType *YieldedOrValueType1Choice `xml:"ValTp"`

	// Place from which the price was obtained.
	SourceOfPrice *MarketIdentification91 `xml:"SrcOfPric,omitempty"`

	// Date on which the price is obtained. With an investment fund, this is as stated in the prospectus.
	QuotationDate *DateAndDateTimeChoice `xml:"QtnDt,omitempty"`
}

func (p *PriceInformation14) AddType() *TypeOfPrice33Choice {
	p.Type = new(TypeOfPrice33Choice)
	return p.Type
}

func (p *PriceInformation14) AddValue() *PriceRateOrAmountOrUnknown1Choice {
	p.Value = new(PriceRateOrAmountOrUnknown1Choice)
	return p.Value
}

func (p *PriceInformation14) AddValueType() *YieldedOrValueType1Choice {
	p.ValueType = new(YieldedOrValueType1Choice)
	return p.ValueType
}

func (p *PriceInformation14) AddSourceOfPrice() *MarketIdentification91 {
	p.SourceOfPrice = new(MarketIdentification91)
	return p.SourceOfPrice
}

func (p *PriceInformation14) AddQuotationDate() *DateAndDateTimeChoice {
	p.QuotationDate = new(DateAndDateTimeChoice)
	return p.QuotationDate
}

// Amount of money for which goods or services are offered, sold, or bought.
type PriceInformation16 struct {

	// Specifies the type of price and information about the price.
	Type *TypeOfPrice45Choice `xml:"Tp"`

	// Type of value in which the price is expressed.
	ValueType *YieldedOrValueType1Choice `xml:"ValTp"`

	// Value of the price, eg, as a currency and value.
	Value *PriceRateOrAmountOrUnknown1Choice `xml:"Val"`

	// Place from which the price was obtained.
	SourceOfPrice *MarketIdentification91 `xml:"SrcOfPric,omitempty"`

	// Date on which the price is obtained. With an investment fund, this is as stated in the prospectus.
	QuotationDate *DateAndDateTimeChoice `xml:"QtnDt,omitempty"`
}

func (p *PriceInformation16) AddType() *TypeOfPrice45Choice {
	p.Type = new(TypeOfPrice45Choice)
	return p.Type
}

func (p *PriceInformation16) AddValueType() *YieldedOrValueType1Choice {
	p.ValueType = new(YieldedOrValueType1Choice)
	return p.ValueType
}

func (p *PriceInformation16) AddValue() *PriceRateOrAmountOrUnknown1Choice {
	p.Value = new(PriceRateOrAmountOrUnknown1Choice)
	return p.Value
}

func (p *PriceInformation16) AddSourceOfPrice() *MarketIdentification91 {
	p.SourceOfPrice = new(MarketIdentification91)
	return p.SourceOfPrice
}

func (p *PriceInformation16) AddQuotationDate() *DateAndDateTimeChoice {
	p.QuotationDate = new(DateAndDateTimeChoice)
	return p.QuotationDate
}

// Amount of money for which goods or services are offered, sold, or bought.
type PriceInformation2 struct {

	// Value of the price, eg, as a currency and value.
	Value *PriceRateOrAmountOrUnknownChoice `xml:"Val"`

	// Type of value in which the price is expressed.
	ValueType *PriceValueType2Code `xml:"ValTp,omitempty"`

	// Type and information about a price.
	Type *TypeOfPrice11Code `xml:"Tp"`

	// Type and information about a price.
	ExtendedType *Extended350Code `xml:"XtndedTp"`

	// Place from which the price was obtained.
	SourceOfPrice *PriceSourceFormatChoice `xml:"SrcOfPric,omitempty"`

	// Date on which the price is obtained. With an investment fund, this is as stated in the prospectus.
	QuotationDate *DateAndDateTimeChoice `xml:"QtnDt,omitempty"`

	// Indicates whether the price is expressed as a yield. The absence of Yielded means it is not applicable.
	Yielded *YesNoIndicator `xml:"Yldd,omitempty"`
}

func (p *PriceInformation2) AddValue() *PriceRateOrAmountOrUnknownChoice {
	p.Value = new(PriceRateOrAmountOrUnknownChoice)
	return p.Value
}

func (p *PriceInformation2) SetValueType(value string) {
	p.ValueType = (*PriceValueType2Code)(&value)
}

func (p *PriceInformation2) SetType(value string) {
	p.Type = (*TypeOfPrice11Code)(&value)
}

func (p *PriceInformation2) SetExtendedType(value string) {
	p.ExtendedType = (*Extended350Code)(&value)
}

func (p *PriceInformation2) AddSourceOfPrice() *PriceSourceFormatChoice {
	p.SourceOfPrice = new(PriceSourceFormatChoice)
	return p.SourceOfPrice
}

func (p *PriceInformation2) AddQuotationDate() *DateAndDateTimeChoice {
	p.QuotationDate = new(DateAndDateTimeChoice)
	return p.QuotationDate
}

func (p *PriceInformation2) SetYielded(value string) {
	p.Yielded = (*YesNoIndicator)(&value)
}

// Amount of money for which goods or services are offered, sold, or bought.
type PriceInformation5 struct {

	// Specifies the type of price and information about the price.
	Type *TypeOfPrice4Choice `xml:"Tp"`

	// Value of the price, eg, as a currency and value.
	Value *PriceRateOrAmountOrUnknownChoice `xml:"Val"`

	// Type of value in which the price is expressed.
	ValueType *YieldedOrValueType1Choice `xml:"ValTp"`

	// Place from which the price was obtained.
	SourceOfPrice *MarketIdentification6 `xml:"SrcOfPric,omitempty"`

	// Date on which the price is obtained. With an investment fund, this is as stated in the prospectus.
	QuotationDate *DateAndDateTimeChoice `xml:"QtnDt,omitempty"`
}

func (p *PriceInformation5) AddType() *TypeOfPrice4Choice {
	p.Type = new(TypeOfPrice4Choice)
	return p.Type
}

func (p *PriceInformation5) AddValue() *PriceRateOrAmountOrUnknownChoice {
	p.Value = new(PriceRateOrAmountOrUnknownChoice)
	return p.Value
}

func (p *PriceInformation5) AddValueType() *YieldedOrValueType1Choice {
	p.ValueType = new(YieldedOrValueType1Choice)
	return p.ValueType
}

func (p *PriceInformation5) AddSourceOfPrice() *MarketIdentification6 {
	p.SourceOfPrice = new(MarketIdentification6)
	return p.SourceOfPrice
}

func (p *PriceInformation5) AddQuotationDate() *DateAndDateTimeChoice {
	p.QuotationDate = new(DateAndDateTimeChoice)
	return p.QuotationDate
}

// Amount of money for which goods or services are offered, sold, or bought.
type PriceInformation6 struct {

	// Specifies the type of price and information about the price.
	Type *TypeOfPrice6Choice `xml:"Tp"`

	// Type of value in which the price is expressed.
	ValueType *YieldedOrValueType1Choice `xml:"ValTp"`

	// Value of the price, eg, as a currency and value.
	Value *PriceRateOrAmountOrUnknownChoice `xml:"Val"`

	// Place from which the price was obtained.
	SourceOfPrice *MarketIdentification6 `xml:"SrcOfPric,omitempty"`

	// Date on which the price is obtained. With an investment fund, this is as stated in the prospectus.
	QuotationDate *DateAndDateTimeChoice `xml:"QtnDt,omitempty"`
}

func (p *PriceInformation6) AddType() *TypeOfPrice6Choice {
	p.Type = new(TypeOfPrice6Choice)
	return p.Type
}

func (p *PriceInformation6) AddValueType() *YieldedOrValueType1Choice {
	p.ValueType = new(YieldedOrValueType1Choice)
	return p.ValueType
}

func (p *PriceInformation6) AddValue() *PriceRateOrAmountOrUnknownChoice {
	p.Value = new(PriceRateOrAmountOrUnknownChoice)
	return p.Value
}

func (p *PriceInformation6) AddSourceOfPrice() *MarketIdentification6 {
	p.SourceOfPrice = new(MarketIdentification6)
	return p.SourceOfPrice
}

func (p *PriceInformation6) AddQuotationDate() *DateAndDateTimeChoice {
	p.QuotationDate = new(DateAndDateTimeChoice)
	return p.QuotationDate
}

// Amount of money for which goods or services are offered, sold, or bought.
type PriceInformation9 struct {

	// Value of the price, for instance, as a currency and value.
	Value *Price4 `xml:"Val"`

	// Date on which the price is obtained. With an investment fund, this is as stated in the prospectus.
	QuotationDate *DateAndDateTime1Choice `xml:"QtnDt,omitempty"`

	// Period during which the price of a security is determined (For  outturn securities).
	PriceCalculationPeriod *DateTimePeriodChoice `xml:"PricClctnPrd,omitempty"`

	// Place from which the price was obtained.
	SourceOfPrice *MarketIdentification77 `xml:"SrcOfPric,omitempty"`
}

func (p *PriceInformation9) AddValue() *Price4 {
	p.Value = new(Price4)
	return p.Value
}

func (p *PriceInformation9) AddQuotationDate() *DateAndDateTime1Choice {
	p.QuotationDate = new(DateAndDateTime1Choice)
	return p.QuotationDate
}

func (p *PriceInformation9) AddPriceCalculationPeriod() *DateTimePeriodChoice {
	p.PriceCalculationPeriod = new(DateTimePeriodChoice)
	return p.PriceCalculationPeriod
}

func (p *PriceInformation9) AddSourceOfPrice() *MarketIdentification77 {
	p.SourceOfPrice = new(MarketIdentification77)
	return p.SourceOfPrice
}
