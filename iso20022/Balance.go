package iso20022

// Net position of a segregated holding, in a single security, within the overall position held in a securities account. A securities balance is calculated from the sum of securities' receipts minus the sum of securities' deliveries.
type Balance1 struct {

	// Indication that the position is short or long.
	ShortLongIndicator *ShortLong1Code `xml:"ShrtLngInd"`

	// Total quantity of financial instruments of the balance.
	Quantity *BalanceQuantity4Choice `xml:"Qty"`
}

func (b *Balance1) SetShortLongIndicator(value string) {
	b.ShortLongIndicator = (*ShortLong1Code)(&value)
}

func (b *Balance1) AddQuantity() *BalanceQuantity4Choice {
	b.Quantity = new(BalanceQuantity4Choice)
	return b.Quantity
}

// Net position of a segregated holding, in a single security, within the overall position held in a securities account. A securities balance is calculated from the sum of securities' receipts minus the sum of securities' deliveries.
type Balance10 struct {

	// Indication that the position is short or long.
	ShortLongIndicator *ShortLong1Code `xml:"ShrtLngInd"`

	// Total quantity of financial instruments of the balance.
	Quantity *BalanceQuantity10Choice `xml:"Qty"`
}

func (b *Balance10) SetShortLongIndicator(value string) {
	b.ShortLongIndicator = (*ShortLong1Code)(&value)
}

func (b *Balance10) AddQuantity() *BalanceQuantity10Choice {
	b.Quantity = new(BalanceQuantity10Choice)
	return b.Quantity
}

// Net position of a segregated holding, in a single security, within the overall position held in a securities account. A securities balance is calculated from the sum of securities' receipts minus the sum of securities' deliveries.
type Balance11 struct {

	// Indication that the position is short or long.
	ShortLongIndicator *ShortLong1Code `xml:"ShrtLngInd,omitempty"`

	// Total quantity of financial instruments of the balance.
	Quantity *BalanceQuantity10Choice `xml:"Qty"`
}

func (b *Balance11) SetShortLongIndicator(value string) {
	b.ShortLongIndicator = (*ShortLong1Code)(&value)
}

func (b *Balance11) AddQuantity() *BalanceQuantity10Choice {
	b.Quantity = new(BalanceQuantity10Choice)
	return b.Quantity
}

// Net position of a segregated holding, in a single security, within the overall position held in a securities account. A securities balance is calculated from the sum of securities' receipts minus the sum of securities' deliveries.
type Balance12 struct {

	// Indication that the position is short or long.
	ShortLongIndicator *ShortLong1Code `xml:"ShrtLngInd,omitempty"`

	// Total quantity of financial instruments of the balance.
	Quantity *BalanceQuantity12Choice `xml:"Qty"`
}

func (b *Balance12) SetShortLongIndicator(value string) {
	b.ShortLongIndicator = (*ShortLong1Code)(&value)
}

func (b *Balance12) AddQuantity() *BalanceQuantity12Choice {
	b.Quantity = new(BalanceQuantity12Choice)
	return b.Quantity
}

// Net position of a segregated holding, in a single security, within the overall position held in a securities account. A securities balance is calculated from the sum of securities' receipts minus the sum of securities' deliveries.
type Balance13 struct {

	// Indication that the position is short or long.
	ShortLongIndicator *ShortLong1Code `xml:"ShrtLngInd,omitempty"`

	// Total quantity of financial instruments of the balance.
	Quantity *SubBalanceQuantity7Choice `xml:"Qty"`
}

func (b *Balance13) SetShortLongIndicator(value string) {
	b.ShortLongIndicator = (*ShortLong1Code)(&value)
}

func (b *Balance13) AddQuantity() *SubBalanceQuantity7Choice {
	b.Quantity = new(SubBalanceQuantity7Choice)
	return b.Quantity
}

// Net position of a segregated holding, in a single security, within the overall position held in a securities account. A securities balance is calculated from the sum of securities' receipts minus the sum of securities' deliveries.
type Balance4 struct {

	// Indication that the position is short or long.
	ShortLongIndicator *ShortLong1Code `xml:"ShrtLngInd,omitempty"`

	// Total quantity of financial instruments of the balance.
	Quantity *BalanceQuantity4Choice `xml:"Qty"`
}

func (b *Balance4) SetShortLongIndicator(value string) {
	b.ShortLongIndicator = (*ShortLong1Code)(&value)
}

func (b *Balance4) AddQuantity() *BalanceQuantity4Choice {
	b.Quantity = new(BalanceQuantity4Choice)
	return b.Quantity
}

// Net position of a segregated holding, in a single security, within the overall position held in a securities account. A securities balance is calculated from the sum of securities' receipts minus the sum of securities' deliveries.
type Balance6 struct {

	// Indication that the position is short or long.
	ShortLongIndicator *ShortLong1Code `xml:"ShrtLngInd"`

	// Total quantity of financial instruments of the balance.
	Quantity *BalanceQuantity9Choice `xml:"Qty"`
}

func (b *Balance6) SetShortLongIndicator(value string) {
	b.ShortLongIndicator = (*ShortLong1Code)(&value)
}

func (b *Balance6) AddQuantity() *BalanceQuantity9Choice {
	b.Quantity = new(BalanceQuantity9Choice)
	return b.Quantity
}

// Net position of a segregated holding, in a single security, within the overall position held in a securities account. A securities balance is calculated from the sum of securities' receipts minus the sum of securities' deliveries.
type Balance7 struct {

	// Indication that the position is short or long.
	ShortLongIndicator *ShortLong1Code `xml:"ShrtLngInd,omitempty"`

	// Total quantity of financial instruments of the balance.
	Quantity *BalanceQuantity9Choice `xml:"Qty"`
}

func (b *Balance7) SetShortLongIndicator(value string) {
	b.ShortLongIndicator = (*ShortLong1Code)(&value)
}

func (b *Balance7) AddQuantity() *BalanceQuantity9Choice {
	b.Quantity = new(BalanceQuantity9Choice)
	return b.Quantity
}

// Net position of a segregated holding, in a single security, within the overall position held in a securities account. A securities balance is calculated from the sum of securities' receipts minus the sum of securities' deliveries.
type Balance8 struct {

	// Indication that the position is short or long.
	ShortLongIndicator *ShortLong1Code `xml:"ShrtLngInd,omitempty"`

	// Total quantity of financial instruments of the balance.
	Quantity *BalanceQuantity8Choice `xml:"Qty"`
}

func (b *Balance8) SetShortLongIndicator(value string) {
	b.ShortLongIndicator = (*ShortLong1Code)(&value)
}

func (b *Balance8) AddQuantity() *BalanceQuantity8Choice {
	b.Quantity = new(BalanceQuantity8Choice)
	return b.Quantity
}

// Net position of a segregated holding, in a single security, within the overall position held in a securities account. A securities balance is calculated from the sum of securities' receipts minus the sum of securities' deliveries.
type Balance9 struct {

	// Indication that the position is short or long.
	ShortLongIndicator *ShortLong1Code `xml:"ShrtLngInd,omitempty"`

	// Total quantity of financial instruments of the balance.
	Quantity *SubBalanceQuantity6Choice `xml:"Qty"`
}

func (b *Balance9) SetShortLongIndicator(value string) {
	b.ShortLongIndicator = (*ShortLong1Code)(&value)
}

func (b *Balance9) AddQuantity() *SubBalanceQuantity6Choice {
	b.Quantity = new(SubBalanceQuantity6Choice)
	return b.Quantity
}
