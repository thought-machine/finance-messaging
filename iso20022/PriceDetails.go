package iso20022

// Provides information about the prices related to a corporate action option.
type PriceDetails10 struct {

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount or a number of points above an index, for example, reinvestment price, strike price and exercise price.
	GenericCashPricePaidPerProduct *PriceFormat23Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat20Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (p *PriceDetails10) AddGenericCashPricePaidPerProduct() *PriceFormat23Choice {
	p.GenericCashPricePaidPerProduct = new(PriceFormat23Choice)
	return p.GenericCashPricePaidPerProduct
}

func (p *PriceDetails10) AddGenericCashPriceReceivedPerProduct() *PriceFormat20Choice {
	p.GenericCashPriceReceivedPerProduct = new(PriceFormat20Choice)
	return p.GenericCashPriceReceivedPerProduct
}

// Provides information about the prices related to a corporate action option.
type PriceDetails11 struct {

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount or a number of points above an index, for example, reinvestment price, strike price and exercise price.
	GenericCashPricePaidPerProduct *PriceFormat6Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat7Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (p *PriceDetails11) AddGenericCashPricePaidPerProduct() *PriceFormat6Choice {
	p.GenericCashPricePaidPerProduct = new(PriceFormat6Choice)
	return p.GenericCashPricePaidPerProduct
}

func (p *PriceDetails11) AddGenericCashPriceReceivedPerProduct() *PriceFormat7Choice {
	p.GenericCashPriceReceivedPerProduct = new(PriceFormat7Choice)
	return p.GenericCashPriceReceivedPerProduct
}

// Provides information about the prices related to a corporate action option.
type PriceDetails14 struct {

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount or a number of points above an index, for example, reinvestment price, strike price and exercise price.
	GenericCashPricePaidPerProduct *PriceFormat6Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat34Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (p *PriceDetails14) AddGenericCashPricePaidPerProduct() *PriceFormat6Choice {
	p.GenericCashPricePaidPerProduct = new(PriceFormat6Choice)
	return p.GenericCashPricePaidPerProduct
}

func (p *PriceDetails14) AddGenericCashPriceReceivedPerProduct() *PriceFormat34Choice {
	p.GenericCashPriceReceivedPerProduct = new(PriceFormat34Choice)
	return p.GenericCashPriceReceivedPerProduct
}

// Provides information about the prices related to a corporate action option.
type PriceDetails15 struct {

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount or a number of points above an index, for example, reinvestment price, strike price and exercise price.
	GenericCashPricePaidPerProduct *PriceFormat23Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat32Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (p *PriceDetails15) AddGenericCashPricePaidPerProduct() *PriceFormat23Choice {
	p.GenericCashPricePaidPerProduct = new(PriceFormat23Choice)
	return p.GenericCashPricePaidPerProduct
}

func (p *PriceDetails15) AddGenericCashPriceReceivedPerProduct() *PriceFormat32Choice {
	p.GenericCashPriceReceivedPerProduct = new(PriceFormat32Choice)
	return p.GenericCashPriceReceivedPerProduct
}

// Provides information about the prices related to a corporate action option.
type PriceDetails2 struct {

	// 1. Price at which security will be purchased/sold if warrant is exercised, either as an actual amount or a percentage.
	// 2. Price at which a bond is converted to underlying security either as an actual amount or a percentage.
	// 3. Strike price of an option, represented either as an actual amount, a percentage or a number of points above an index.
	ExercisePrice *PriceFormat6Choice `xml:"ExrcPric,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount, for example, reinvestment price.
	GenericCashPricePaidPerProduct *PriceFormat5Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat7Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (p *PriceDetails2) AddExercisePrice() *PriceFormat6Choice {
	p.ExercisePrice = new(PriceFormat6Choice)
	return p.ExercisePrice
}

func (p *PriceDetails2) AddGenericCashPricePaidPerProduct() *PriceFormat5Choice {
	p.GenericCashPricePaidPerProduct = new(PriceFormat5Choice)
	return p.GenericCashPricePaidPerProduct
}

func (p *PriceDetails2) AddGenericCashPriceReceivedPerProduct() *PriceFormat7Choice {
	p.GenericCashPriceReceivedPerProduct = new(PriceFormat7Choice)
	return p.GenericCashPriceReceivedPerProduct
}

// Provides information about the prices related to a corporate action option.
type PriceDetails22 struct {

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount or a number of points above an index, for example, reinvestment price, strike price and exercise price.
	GenericCashPricePaidPerProduct *PriceFormat44Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat47Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (p *PriceDetails22) AddGenericCashPricePaidPerProduct() *PriceFormat44Choice {
	p.GenericCashPricePaidPerProduct = new(PriceFormat44Choice)
	return p.GenericCashPricePaidPerProduct
}

func (p *PriceDetails22) AddGenericCashPriceReceivedPerProduct() *PriceFormat47Choice {
	p.GenericCashPriceReceivedPerProduct = new(PriceFormat47Choice)
	return p.GenericCashPriceReceivedPerProduct
}

// Provides information about the prices related to a corporate action option.
type PriceDetails23 struct {

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount or a number of points above an index, for example, reinvestment price, strike price and exercise price.
	GenericCashPricePaidPerProduct *PriceFormat51Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat48Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (p *PriceDetails23) AddGenericCashPricePaidPerProduct() *PriceFormat51Choice {
	p.GenericCashPricePaidPerProduct = new(PriceFormat51Choice)
	return p.GenericCashPricePaidPerProduct
}

func (p *PriceDetails23) AddGenericCashPriceReceivedPerProduct() *PriceFormat48Choice {
	p.GenericCashPriceReceivedPerProduct = new(PriceFormat48Choice)
	return p.GenericCashPriceReceivedPerProduct
}

// Provides information about the prices related to a corporate action option.
type PriceDetails24 struct {

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount or a number of points above an index, for example, reinvestment price, strike price and exercise price.
	GenericCashPricePaidPerProduct *PriceFormat55Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat56Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (p *PriceDetails24) AddGenericCashPricePaidPerProduct() *PriceFormat55Choice {
	p.GenericCashPricePaidPerProduct = new(PriceFormat55Choice)
	return p.GenericCashPricePaidPerProduct
}

func (p *PriceDetails24) AddGenericCashPriceReceivedPerProduct() *PriceFormat56Choice {
	p.GenericCashPriceReceivedPerProduct = new(PriceFormat56Choice)
	return p.GenericCashPriceReceivedPerProduct
}

// Provides information about the prices related to a corporate action option.
type PriceDetails25 struct {

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount or a number of points above an index, for example, reinvestment price, strike price and exercise price.
	GenericCashPricePaidPerProduct *PriceFormat59Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat60Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (p *PriceDetails25) AddGenericCashPricePaidPerProduct() *PriceFormat59Choice {
	p.GenericCashPricePaidPerProduct = new(PriceFormat59Choice)
	return p.GenericCashPricePaidPerProduct
}

func (p *PriceDetails25) AddGenericCashPriceReceivedPerProduct() *PriceFormat60Choice {
	p.GenericCashPriceReceivedPerProduct = new(PriceFormat60Choice)
	return p.GenericCashPriceReceivedPerProduct
}

// Provides information about the prices related to a corporate action option.
type PriceDetails3 struct {

	// 1. Price at which security will be purchased/sold if warrant is exercised, either as an actual amount or a percentage.
	// 2. Price at which a bond is converted to underlying security either as an actual amount or a percentage.
	// 3. Strike price of an option, represented either as an actual amount, a percentage or a number of points above an index.
	ExercisePrice *PriceFormat23Choice `xml:"ExrcPric,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount, for example, reinvestment price.
	GenericCashPricePaidPerProduct *PriceFormat19Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat22Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (p *PriceDetails3) AddExercisePrice() *PriceFormat23Choice {
	p.ExercisePrice = new(PriceFormat23Choice)
	return p.ExercisePrice
}

func (p *PriceDetails3) AddGenericCashPricePaidPerProduct() *PriceFormat19Choice {
	p.GenericCashPricePaidPerProduct = new(PriceFormat19Choice)
	return p.GenericCashPricePaidPerProduct
}

func (p *PriceDetails3) AddGenericCashPriceReceivedPerProduct() *PriceFormat22Choice {
	p.GenericCashPriceReceivedPerProduct = new(PriceFormat22Choice)
	return p.GenericCashPriceReceivedPerProduct
}

// Provides information about the prices related to a corporate action option.
type PriceDetails6 struct {

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount, for example, reinvestment price.
	GenericCashPricePaidPerProduct *PriceFormat23Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat20Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (p *PriceDetails6) AddGenericCashPricePaidPerProduct() *PriceFormat23Choice {
	p.GenericCashPricePaidPerProduct = new(PriceFormat23Choice)
	return p.GenericCashPricePaidPerProduct
}

func (p *PriceDetails6) AddGenericCashPriceReceivedPerProduct() *PriceFormat20Choice {
	p.GenericCashPriceReceivedPerProduct = new(PriceFormat20Choice)
	return p.GenericCashPriceReceivedPerProduct
}

// Provides information about the prices related to a corporate action option.
type PriceDetails7 struct {

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount, for example, reinvestment price.
	GenericCashPricePaidPerProduct *PriceFormat6Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat7Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (p *PriceDetails7) AddGenericCashPricePaidPerProduct() *PriceFormat6Choice {
	p.GenericCashPricePaidPerProduct = new(PriceFormat6Choice)
	return p.GenericCashPricePaidPerProduct
}

func (p *PriceDetails7) AddGenericCashPriceReceivedPerProduct() *PriceFormat7Choice {
	p.GenericCashPriceReceivedPerProduct = new(PriceFormat7Choice)
	return p.GenericCashPriceReceivedPerProduct
}
