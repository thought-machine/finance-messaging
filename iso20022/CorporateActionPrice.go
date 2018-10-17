package iso20022

// Specifies prices.
type CorporateActionPrice1 struct {

	// 1. Price at which security will be purchased/sold if warrant is exercised, either as an actual amount or a percentage.
	// 2. Price at which a bond is converted to underlying security either as an actual amount or a percentage.
	// 3. Strike price of an option, represented either as an actual amount, a percentage or or a number of points above an index.
	ExercisePrice *PriceFormat4Choice `xml:"ExrcPric,omitempty"`

	// Initial issue price of a financial instrument.
	IssuePrice *PriceFormat2Choice `xml:"IssePric,omitempty"`

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat2Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Amount included in the dividend/NAV that is identified as gains directly or indirectly derived from interest payments within the scope of the EU Savings directive.
	TaxableIncomePerDividendShare *AmountPrice1 `xml:"TaxblIncmPerDvddShr,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, eg, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat1Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount, eg, reinvestment price.
	GenericCashPricePaidPerProduct *PriceFormat2Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Amount of money required per over-subscribed equity as defined by the issuer.
	OverSubscriptionDepositPrice *PriceFormat2Choice `xml:"OverSbcptDpstPric,omitempty"`
}

func (c *CorporateActionPrice1) AddExercisePrice() *PriceFormat4Choice {
	c.ExercisePrice = new(PriceFormat4Choice)
	return c.ExercisePrice
}

func (c *CorporateActionPrice1) AddIssuePrice() *PriceFormat2Choice {
	c.IssuePrice = new(PriceFormat2Choice)
	return c.IssuePrice
}

func (c *CorporateActionPrice1) AddCashInLieuOfSharePrice() *PriceFormat2Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat2Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice1) AddTaxableIncomePerDividendShare() *AmountPrice1 {
	c.TaxableIncomePerDividendShare = new(AmountPrice1)
	return c.TaxableIncomePerDividendShare
}

func (c *CorporateActionPrice1) AddGenericCashPriceReceivedPerProduct() *PriceFormat1Choice {
	c.GenericCashPriceReceivedPerProduct = new(PriceFormat1Choice)
	return c.GenericCashPriceReceivedPerProduct
}

func (c *CorporateActionPrice1) AddGenericCashPricePaidPerProduct() *PriceFormat2Choice {
	c.GenericCashPricePaidPerProduct = new(PriceFormat2Choice)
	return c.GenericCashPricePaidPerProduct
}

func (c *CorporateActionPrice1) AddOverSubscriptionDepositPrice() *PriceFormat2Choice {
	c.OverSubscriptionDepositPrice = new(PriceFormat2Choice)
	return c.OverSubscriptionDepositPrice
}

// Specifies prices related to a corporate action option.
type CorporateActionPrice10 struct {

	// Indicates whether the price is an indicative price or a market price.
	IndicativeOrMarketPrice *IndicativeOrMarketPrice1Choice `xml:"IndctvOrMktPric,omitempty"`

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat11Choice `xml:"CshInLieuOfShrPric,omitempty"`
}

func (c *CorporateActionPrice10) AddIndicativeOrMarketPrice() *IndicativeOrMarketPrice1Choice {
	c.IndicativeOrMarketPrice = new(IndicativeOrMarketPrice1Choice)
	return c.IndicativeOrMarketPrice
}

func (c *CorporateActionPrice10) AddCashInLieuOfSharePrice() *PriceFormat11Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat11Choice)
	return c.CashInLieuOfSharePrice
}

// Specifies prices related to a corporate action option.
type CorporateActionPrice16 struct {

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat19Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct []*PriceFormat20Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`

	// Amount of money required per over-subscribed equity as defined by the issuer.
	OverSubscriptionDepositPrice *PriceFormat19Choice `xml:"OverSbcptDpstPric,omitempty"`
}

func (c *CorporateActionPrice16) AddCashInLieuOfSharePrice() *PriceFormat19Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat19Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice16) AddGenericCashPriceReceivedPerProduct() *PriceFormat20Choice {
	newValue := new(PriceFormat20Choice)
	c.GenericCashPriceReceivedPerProduct = append(c.GenericCashPriceReceivedPerProduct, newValue)
	return newValue
}

func (c *CorporateActionPrice16) AddOverSubscriptionDepositPrice() *PriceFormat19Choice {
	c.OverSubscriptionDepositPrice = new(PriceFormat19Choice)
	return c.OverSubscriptionDepositPrice
}

// Specifies prices of a corporate action.
type CorporateActionPrice17 struct {

	// Maximum or cap price at which a holder can bid, for example, on a Dutch auction offer.
	MaximumPrice *PriceFormat19Choice `xml:"MaxPric,omitempty"`

	// Minimum or floor price at which a holder can bid, for example, on a Dutch auction offer.
	MinimumPrice *PriceFormat19Choice `xml:"MinPric,omitempty"`
}

func (c *CorporateActionPrice17) AddMaximumPrice() *PriceFormat19Choice {
	c.MaximumPrice = new(PriceFormat19Choice)
	return c.MaximumPrice
}

func (c *CorporateActionPrice17) AddMinimumPrice() *PriceFormat19Choice {
	c.MinimumPrice = new(PriceFormat19Choice)
	return c.MinimumPrice
}

// Specifies prices related to a corporate action option.
type CorporateActionPrice18 struct {

	// Indicates whether the price is an indicative price or a market price.
	IndicativeOrMarketPrice *IndicativeOrMarketPrice5Choice `xml:"IndctvOrMktPric,omitempty"`

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat19Choice `xml:"CshInLieuOfShrPric,omitempty"`
}

func (c *CorporateActionPrice18) AddIndicativeOrMarketPrice() *IndicativeOrMarketPrice5Choice {
	c.IndicativeOrMarketPrice = new(IndicativeOrMarketPrice5Choice)
	return c.IndicativeOrMarketPrice
}

func (c *CorporateActionPrice18) AddCashInLieuOfSharePrice() *PriceFormat19Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat19Choice)
	return c.CashInLieuOfSharePrice
}

// Specifies prices related to a corporate action option.
type CorporateActionPrice19 struct {

	// Indicates whether the price is an indicative price or a market price.
	IndicativeOrMarketPrice *IndicativeOrMarketPrice2Choice `xml:"IndctvOrMktPric,omitempty"`

	// 1. Price at which security will be purchased/sold if warrant is exercised, either as an actual amount or a percentage.
	// 2. Price at which a bond is converted to underlying security either as an actual amount or a percentage.
	// 3. Strike price of an option, represented either as an actual amount, a percentage or a number of points above an index.
	ExercisePrice *PriceFormat5Choice `xml:"ExrcPric,omitempty"`

	// Initial issue price of a financial instrument.
	IssuePrice *PriceFormat5Choice `xml:"IssePric,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat21Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount, for example, reinvestment price.
	GenericCashPricePaidPerProduct *PriceFormat5Choice `xml:"GncCshPricPdPerPdct,omitempty"`
}

func (c *CorporateActionPrice19) AddIndicativeOrMarketPrice() *IndicativeOrMarketPrice2Choice {
	c.IndicativeOrMarketPrice = new(IndicativeOrMarketPrice2Choice)
	return c.IndicativeOrMarketPrice
}

func (c *CorporateActionPrice19) AddExercisePrice() *PriceFormat5Choice {
	c.ExercisePrice = new(PriceFormat5Choice)
	return c.ExercisePrice
}

func (c *CorporateActionPrice19) AddIssuePrice() *PriceFormat5Choice {
	c.IssuePrice = new(PriceFormat5Choice)
	return c.IssuePrice
}

func (c *CorporateActionPrice19) AddGenericCashPriceReceivedPerProduct() *PriceFormat21Choice {
	c.GenericCashPriceReceivedPerProduct = new(PriceFormat21Choice)
	return c.GenericCashPriceReceivedPerProduct
}

func (c *CorporateActionPrice19) AddGenericCashPricePaidPerProduct() *PriceFormat5Choice {
	c.GenericCashPricePaidPerProduct = new(PriceFormat5Choice)
	return c.GenericCashPricePaidPerProduct
}

// Specifies prices.
type CorporateActionPrice2 struct {

	// Maximum or cap price at which a holder can bid, e.g. on a Dutch auction offer.
	MaximumPrice *PriceFormat3Choice `xml:"MaxPric,omitempty"`

	// Minimum or floor price at which a holder can bid, e.g. on a Dutch auction offer.
	MinimumPrice *PriceFormat3Choice `xml:"MinPric,omitempty"`
}

func (c *CorporateActionPrice2) AddMaximumPrice() *PriceFormat3Choice {
	c.MaximumPrice = new(PriceFormat3Choice)
	return c.MaximumPrice
}

func (c *CorporateActionPrice2) AddMinimumPrice() *PriceFormat3Choice {
	c.MinimumPrice = new(PriceFormat3Choice)
	return c.MinimumPrice
}

// Specifies prices.
type CorporateActionPrice21 struct {

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct []*PriceFormat7Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat5Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Amount of money required per over-subscribed equity as defined by the issuer.
	OverSubscriptionDepositPrice *PriceFormat5Choice `xml:"OverSbcptDpstPric,omitempty"`
}

func (c *CorporateActionPrice21) AddGenericCashPriceReceivedPerProduct() *PriceFormat7Choice {
	newValue := new(PriceFormat7Choice)
	c.GenericCashPriceReceivedPerProduct = append(c.GenericCashPriceReceivedPerProduct, newValue)
	return newValue
}

func (c *CorporateActionPrice21) AddCashInLieuOfSharePrice() *PriceFormat5Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat5Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice21) AddOverSubscriptionDepositPrice() *PriceFormat5Choice {
	c.OverSubscriptionDepositPrice = new(PriceFormat5Choice)
	return c.OverSubscriptionDepositPrice
}

// Specifies prices related to a corporate action option.
type CorporateActionPrice26 struct {

	// Indicates whether the price is an indicative price or a market price.
	IndicativeOrMarketPrice *IndicativeOrMarketPrice5Choice `xml:"IndctvOrMktPric,omitempty"`

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat19Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Cash value of resulting securities proceeds for tax calculation and/or reporting.
	CashValueForTax *PriceFormat29Choice `xml:"CshValForTax,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount, for example, reinvestment price.
	GenericCashPricePaidPerProduct *PriceFormat23Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat20Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (c *CorporateActionPrice26) AddIndicativeOrMarketPrice() *IndicativeOrMarketPrice5Choice {
	c.IndicativeOrMarketPrice = new(IndicativeOrMarketPrice5Choice)
	return c.IndicativeOrMarketPrice
}

func (c *CorporateActionPrice26) AddCashInLieuOfSharePrice() *PriceFormat19Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat19Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice26) AddCashValueForTax() *PriceFormat29Choice {
	c.CashValueForTax = new(PriceFormat29Choice)
	return c.CashValueForTax
}

func (c *CorporateActionPrice26) AddGenericCashPricePaidPerProduct() *PriceFormat23Choice {
	c.GenericCashPricePaidPerProduct = new(PriceFormat23Choice)
	return c.GenericCashPricePaidPerProduct
}

func (c *CorporateActionPrice26) AddGenericCashPriceReceivedPerProduct() *PriceFormat20Choice {
	c.GenericCashPriceReceivedPerProduct = new(PriceFormat20Choice)
	return c.GenericCashPriceReceivedPerProduct
}

// Specifies prices.
type CorporateActionPrice27 struct {

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat5Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Indicates whether the price is an indicative price or a market price.
	IndicativeOrMarketPrice *IndicativeOrMarketPrice2Choice `xml:"IndctvOrMktPric,omitempty"`

	// Cash value of resulting securities proceeds for tax calculation and/or reporting.
	CashValueForTax *AmountPrice2 `xml:"CshValForTax,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount, for example, reinvestment price.
	GenericCashPricePaidPerProduct *PriceFormat6Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat7Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (c *CorporateActionPrice27) AddCashInLieuOfSharePrice() *PriceFormat5Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat5Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice27) AddIndicativeOrMarketPrice() *IndicativeOrMarketPrice2Choice {
	c.IndicativeOrMarketPrice = new(IndicativeOrMarketPrice2Choice)
	return c.IndicativeOrMarketPrice
}

func (c *CorporateActionPrice27) AddCashValueForTax() *AmountPrice2 {
	c.CashValueForTax = new(AmountPrice2)
	return c.CashValueForTax
}

func (c *CorporateActionPrice27) AddGenericCashPricePaidPerProduct() *PriceFormat6Choice {
	c.GenericCashPricePaidPerProduct = new(PriceFormat6Choice)
	return c.GenericCashPricePaidPerProduct
}

func (c *CorporateActionPrice27) AddGenericCashPriceReceivedPerProduct() *PriceFormat7Choice {
	c.GenericCashPriceReceivedPerProduct = new(PriceFormat7Choice)
	return c.GenericCashPriceReceivedPerProduct
}

// Specifies prices related to a corporate action option.
type CorporateActionPrice28 struct {

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat19Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Amount of money required per over-subscribed equity as defined by the issuer.
	OverSubscriptionDepositPrice *PriceFormat19Choice `xml:"OverSbcptDpstPric,omitempty"`
}

func (c *CorporateActionPrice28) AddCashInLieuOfSharePrice() *PriceFormat19Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat19Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice28) AddOverSubscriptionDepositPrice() *PriceFormat19Choice {
	c.OverSubscriptionDepositPrice = new(PriceFormat19Choice)
	return c.OverSubscriptionDepositPrice
}

// Specifies prices related to a corporate action option.
type CorporateActionPrice29 struct {

	// Indicates whether the price is an indicative price or a market price.
	IndicativeOrMarketPrice *IndicativeOrMarketPrice2Choice `xml:"IndctvOrMktPric,omitempty"`

	// Initial issue price of a financial instrument.
	IssuePrice *PriceFormat5Choice `xml:"IssePric,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat21Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount, for example, reinvestment price.
	GenericCashPricePaidPerProduct *PriceFormat5Choice `xml:"GncCshPricPdPerPdct,omitempty"`
}

func (c *CorporateActionPrice29) AddIndicativeOrMarketPrice() *IndicativeOrMarketPrice2Choice {
	c.IndicativeOrMarketPrice = new(IndicativeOrMarketPrice2Choice)
	return c.IndicativeOrMarketPrice
}

func (c *CorporateActionPrice29) AddIssuePrice() *PriceFormat5Choice {
	c.IssuePrice = new(PriceFormat5Choice)
	return c.IssuePrice
}

func (c *CorporateActionPrice29) AddGenericCashPriceReceivedPerProduct() *PriceFormat21Choice {
	c.GenericCashPriceReceivedPerProduct = new(PriceFormat21Choice)
	return c.GenericCashPriceReceivedPerProduct
}

func (c *CorporateActionPrice29) AddGenericCashPricePaidPerProduct() *PriceFormat5Choice {
	c.GenericCashPricePaidPerProduct = new(PriceFormat5Choice)
	return c.GenericCashPricePaidPerProduct
}

// Specifies prices of a corporate action.
type CorporateActionPrice3 struct {

	// Maximum or cap price at which a holder can bid, for example, on a Dutch auction offer.
	MaximumPrice *PriceFormat11Choice `xml:"MaxPric,omitempty"`

	// Minimum or floor price at which a holder can bid, for example, on a Dutch auction offer.
	MinimumPrice *PriceFormat11Choice `xml:"MinPric,omitempty"`
}

func (c *CorporateActionPrice3) AddMaximumPrice() *PriceFormat11Choice {
	c.MaximumPrice = new(PriceFormat11Choice)
	return c.MaximumPrice
}

func (c *CorporateActionPrice3) AddMinimumPrice() *PriceFormat11Choice {
	c.MinimumPrice = new(PriceFormat11Choice)
	return c.MinimumPrice
}

// Specifies prices.
type CorporateActionPrice30 struct {

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat5Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Amount of money required per over-subscribed equity as defined by the issuer.
	OverSubscriptionDepositPrice *PriceFormat5Choice `xml:"OverSbcptDpstPric,omitempty"`
}

func (c *CorporateActionPrice30) AddCashInLieuOfSharePrice() *PriceFormat5Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat5Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice30) AddOverSubscriptionDepositPrice() *PriceFormat5Choice {
	c.OverSubscriptionDepositPrice = new(PriceFormat5Choice)
	return c.OverSubscriptionDepositPrice
}

// Specifies prices related to a corporate action option.
type CorporateActionPrice31 struct {

	// Indicates whether the price is an indicative price or a market price.
	IndicativeOrMarketPrice *IndicativeOrMarketPrice5Choice `xml:"IndctvOrMktPric,omitempty"`

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat19Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Cash value of resulting securities proceeds for tax calculation and/or reporting.
	CashValueForTax *PriceFormat29Choice `xml:"CshValForTax,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount, for example, reinvestment price.
	GenericCashPricePaidPerProduct *PriceFormat23Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat20Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (c *CorporateActionPrice31) AddIndicativeOrMarketPrice() *IndicativeOrMarketPrice5Choice {
	c.IndicativeOrMarketPrice = new(IndicativeOrMarketPrice5Choice)
	return c.IndicativeOrMarketPrice
}

func (c *CorporateActionPrice31) AddCashInLieuOfSharePrice() *PriceFormat19Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat19Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice31) AddCashValueForTax() *PriceFormat29Choice {
	c.CashValueForTax = new(PriceFormat29Choice)
	return c.CashValueForTax
}

func (c *CorporateActionPrice31) AddGenericCashPricePaidPerProduct() *PriceFormat23Choice {
	c.GenericCashPricePaidPerProduct = new(PriceFormat23Choice)
	return c.GenericCashPricePaidPerProduct
}

func (c *CorporateActionPrice31) AddGenericCashPriceReceivedPerProduct() *PriceFormat20Choice {
	c.GenericCashPriceReceivedPerProduct = new(PriceFormat20Choice)
	return c.GenericCashPriceReceivedPerProduct
}

// Specifies prices related to a corporate action option.
type CorporateActionPrice38 struct {

	// Indicates whether the price is an indicative price or a market price.
	IndicativeOrMarketPrice *IndicativeOrMarketPrice5Choice `xml:"IndctvOrMktPric,omitempty"`

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat19Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Cash value of resulting securities proceeds for tax calculation and/or reporting.
	CashValueForTax *PriceFormat29Choice `xml:"CshValForTax,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount or a number of points above an index, for example, reinvestment price, strike price and exercise price.
	GenericCashPricePaidPerProduct *PriceFormat23Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat20Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (c *CorporateActionPrice38) AddIndicativeOrMarketPrice() *IndicativeOrMarketPrice5Choice {
	c.IndicativeOrMarketPrice = new(IndicativeOrMarketPrice5Choice)
	return c.IndicativeOrMarketPrice
}

func (c *CorporateActionPrice38) AddCashInLieuOfSharePrice() *PriceFormat19Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat19Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice38) AddCashValueForTax() *PriceFormat29Choice {
	c.CashValueForTax = new(PriceFormat29Choice)
	return c.CashValueForTax
}

func (c *CorporateActionPrice38) AddGenericCashPricePaidPerProduct() *PriceFormat23Choice {
	c.GenericCashPricePaidPerProduct = new(PriceFormat23Choice)
	return c.GenericCashPricePaidPerProduct
}

func (c *CorporateActionPrice38) AddGenericCashPriceReceivedPerProduct() *PriceFormat20Choice {
	c.GenericCashPriceReceivedPerProduct = new(PriceFormat20Choice)
	return c.GenericCashPriceReceivedPerProduct
}

// Specifies prices.
type CorporateActionPrice39 struct {

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat5Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Indicates whether the price is an indicative price or a market price.
	IndicativeOrMarketPrice *IndicativeOrMarketPrice2Choice `xml:"IndctvOrMktPric,omitempty"`

	// Cash value of resulting securities proceeds for tax calculation and/or reporting.
	CashValueForTax *AmountPrice2 `xml:"CshValForTax,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount or a number of points above an index, for example, reinvestment price, strike price and exercise price.
	GenericCashPricePaidPerProduct *PriceFormat6Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat7Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (c *CorporateActionPrice39) AddCashInLieuOfSharePrice() *PriceFormat5Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat5Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice39) AddIndicativeOrMarketPrice() *IndicativeOrMarketPrice2Choice {
	c.IndicativeOrMarketPrice = new(IndicativeOrMarketPrice2Choice)
	return c.IndicativeOrMarketPrice
}

func (c *CorporateActionPrice39) AddCashValueForTax() *AmountPrice2 {
	c.CashValueForTax = new(AmountPrice2)
	return c.CashValueForTax
}

func (c *CorporateActionPrice39) AddGenericCashPricePaidPerProduct() *PriceFormat6Choice {
	c.GenericCashPricePaidPerProduct = new(PriceFormat6Choice)
	return c.GenericCashPricePaidPerProduct
}

func (c *CorporateActionPrice39) AddGenericCashPriceReceivedPerProduct() *PriceFormat7Choice {
	c.GenericCashPriceReceivedPerProduct = new(PriceFormat7Choice)
	return c.GenericCashPriceReceivedPerProduct
}

// Specifies prices.
type CorporateActionPrice4 struct {

	// Estimated price, eg, for valuation purposes.
	IndicativePrice *PriceFormat2Choice `xml:"IndctvPric,omitempty"`

	// Last reported/known price of a financial instrument in a market.
	MarketPrice *PriceFormat2Choice `xml:"MktPric,omitempty"`
}

func (c *CorporateActionPrice4) AddIndicativePrice() *PriceFormat2Choice {
	c.IndicativePrice = new(PriceFormat2Choice)
	return c.IndicativePrice
}

func (c *CorporateActionPrice4) AddMarketPrice() *PriceFormat2Choice {
	c.MarketPrice = new(PriceFormat2Choice)
	return c.MarketPrice
}

// Specifies prices of a corporate action.
type CorporateActionPrice42 struct {

	// Maximum or cap price at which a holder can bid, for example, on a Dutch auction offer.
	MaximumPrice *PriceFormat23Choice `xml:"MaxPric,omitempty"`

	// Minimum or floor price at which a holder can bid, for example, on a Dutch auction offer.
	MinimumPrice *PriceFormat23Choice `xml:"MinPric,omitempty"`
}

func (c *CorporateActionPrice42) AddMaximumPrice() *PriceFormat23Choice {
	c.MaximumPrice = new(PriceFormat23Choice)
	return c.MaximumPrice
}

func (c *CorporateActionPrice42) AddMinimumPrice() *PriceFormat23Choice {
	c.MinimumPrice = new(PriceFormat23Choice)
	return c.MinimumPrice
}

// Specifies prices related to a corporate action option.
type CorporateActionPrice43 struct {

	// Indicates whether the price is an indicative price or a market price.
	IndicativeOrMarketPrice *IndicativeOrMarketPrice5Choice `xml:"IndctvOrMktPric,omitempty"`

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat19Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Cash value of resulting securities proceeds for tax calculation and/or reporting.
	CashValueForTax *PriceFormat29Choice `xml:"CshValForTax,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount or a number of points above an index, for example, reinvestment price, strike price and exercise price.
	GenericCashPricePaidPerProduct *PriceFormat23Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat32Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (c *CorporateActionPrice43) AddIndicativeOrMarketPrice() *IndicativeOrMarketPrice5Choice {
	c.IndicativeOrMarketPrice = new(IndicativeOrMarketPrice5Choice)
	return c.IndicativeOrMarketPrice
}

func (c *CorporateActionPrice43) AddCashInLieuOfSharePrice() *PriceFormat19Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat19Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice43) AddCashValueForTax() *PriceFormat29Choice {
	c.CashValueForTax = new(PriceFormat29Choice)
	return c.CashValueForTax
}

func (c *CorporateActionPrice43) AddGenericCashPricePaidPerProduct() *PriceFormat23Choice {
	c.GenericCashPricePaidPerProduct = new(PriceFormat23Choice)
	return c.GenericCashPricePaidPerProduct
}

func (c *CorporateActionPrice43) AddGenericCashPriceReceivedPerProduct() *PriceFormat32Choice {
	c.GenericCashPriceReceivedPerProduct = new(PriceFormat32Choice)
	return c.GenericCashPriceReceivedPerProduct
}

// Specifies prices related to a corporate action option.
type CorporateActionPrice44 struct {

	// Indicates whether the price is an indicative price or a market price.
	IndicativeOrMarketPrice *IndicativeOrMarketPrice2Choice `xml:"IndctvOrMktPric,omitempty"`

	// Initial issue price of a financial instrument.
	IssuePrice *PriceFormat5Choice `xml:"IssePric,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat33Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount, for example, reinvestment price.
	GenericCashPricePaidPerProduct *PriceFormat5Choice `xml:"GncCshPricPdPerPdct,omitempty"`
}

func (c *CorporateActionPrice44) AddIndicativeOrMarketPrice() *IndicativeOrMarketPrice2Choice {
	c.IndicativeOrMarketPrice = new(IndicativeOrMarketPrice2Choice)
	return c.IndicativeOrMarketPrice
}

func (c *CorporateActionPrice44) AddIssuePrice() *PriceFormat5Choice {
	c.IssuePrice = new(PriceFormat5Choice)
	return c.IssuePrice
}

func (c *CorporateActionPrice44) AddGenericCashPriceReceivedPerProduct() *PriceFormat33Choice {
	c.GenericCashPriceReceivedPerProduct = new(PriceFormat33Choice)
	return c.GenericCashPriceReceivedPerProduct
}

func (c *CorporateActionPrice44) AddGenericCashPricePaidPerProduct() *PriceFormat5Choice {
	c.GenericCashPricePaidPerProduct = new(PriceFormat5Choice)
	return c.GenericCashPricePaidPerProduct
}

// Specifies prices.
type CorporateActionPrice45 struct {

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat5Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Indicates whether the price is an indicative price or a market price.
	IndicativeOrMarketPrice *IndicativeOrMarketPrice2Choice `xml:"IndctvOrMktPric,omitempty"`

	// Cash value of resulting securities proceeds for tax calculation and/or reporting.
	CashValueForTax *AmountPrice2 `xml:"CshValForTax,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount or a number of points above an index, for example, reinvestment price, strike price and exercise price.
	GenericCashPricePaidPerProduct *PriceFormat6Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat34Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (c *CorporateActionPrice45) AddCashInLieuOfSharePrice() *PriceFormat5Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat5Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice45) AddIndicativeOrMarketPrice() *IndicativeOrMarketPrice2Choice {
	c.IndicativeOrMarketPrice = new(IndicativeOrMarketPrice2Choice)
	return c.IndicativeOrMarketPrice
}

func (c *CorporateActionPrice45) AddCashValueForTax() *AmountPrice2 {
	c.CashValueForTax = new(AmountPrice2)
	return c.CashValueForTax
}

func (c *CorporateActionPrice45) AddGenericCashPricePaidPerProduct() *PriceFormat6Choice {
	c.GenericCashPricePaidPerProduct = new(PriceFormat6Choice)
	return c.GenericCashPricePaidPerProduct
}

func (c *CorporateActionPrice45) AddGenericCashPriceReceivedPerProduct() *PriceFormat34Choice {
	c.GenericCashPriceReceivedPerProduct = new(PriceFormat34Choice)
	return c.GenericCashPriceReceivedPerProduct
}

// Specifies prices.
type CorporateActionPrice5 struct {

	// 1. Price at which security will be purchased/sold if warrant is exercised, either as an actual amount or a percentage.
	// 2. Price at which a bond is converted to underlying security either as an actual amount or a percentage.
	// 3. Strike price of an option, represented either as an actual amount, a percentage or a number of points above an index.
	ExercisePrice *PriceFormat6Choice `xml:"ExrcPric,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct []*PriceFormat7Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount, for example, reinvestment price.
	GenericCashPricePaidPerProduct *PriceFormat5Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Amount included in the dividend/NAV that is identified as gains directly or indirectly derived from interest payments within the scope of the EU Savings directive.
	TaxableIncomePerDividendShare *AmountPrice3 `xml:"TaxblIncmPerDvddShr,omitempty"`

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat5Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Amount of money required per over-subscribed equity as defined by the issuer.
	OverSubscriptionDepositPrice *PriceFormat5Choice `xml:"OverSbcptDpstPric,omitempty"`
}

func (c *CorporateActionPrice5) AddExercisePrice() *PriceFormat6Choice {
	c.ExercisePrice = new(PriceFormat6Choice)
	return c.ExercisePrice
}

func (c *CorporateActionPrice5) AddGenericCashPriceReceivedPerProduct() *PriceFormat7Choice {
	newValue := new(PriceFormat7Choice)
	c.GenericCashPriceReceivedPerProduct = append(c.GenericCashPriceReceivedPerProduct, newValue)
	return newValue
}

func (c *CorporateActionPrice5) AddGenericCashPricePaidPerProduct() *PriceFormat5Choice {
	c.GenericCashPricePaidPerProduct = new(PriceFormat5Choice)
	return c.GenericCashPricePaidPerProduct
}

func (c *CorporateActionPrice5) AddTaxableIncomePerDividendShare() *AmountPrice3 {
	c.TaxableIncomePerDividendShare = new(AmountPrice3)
	return c.TaxableIncomePerDividendShare
}

func (c *CorporateActionPrice5) AddCashInLieuOfSharePrice() *PriceFormat5Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat5Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice5) AddOverSubscriptionDepositPrice() *PriceFormat5Choice {
	c.OverSubscriptionDepositPrice = new(PriceFormat5Choice)
	return c.OverSubscriptionDepositPrice
}

// Specifies prices related to a corporate action option.
type CorporateActionPrice56 struct {

	// Indicates whether the price is an indicative price or a market price.
	IndicativeOrMarketPrice *IndicativeOrMarketPrice7Choice `xml:"IndctvOrMktPric,omitempty"`

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat45Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Cash value of resulting securities proceeds for tax calculation and/or reporting.
	CashValueForTax *PriceFormat46Choice `xml:"CshValForTax,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount or a number of points above an index, for example, reinvestment price, strike price and exercise price.
	GenericCashPricePaidPerProduct *PriceFormat44Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat47Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (c *CorporateActionPrice56) AddIndicativeOrMarketPrice() *IndicativeOrMarketPrice7Choice {
	c.IndicativeOrMarketPrice = new(IndicativeOrMarketPrice7Choice)
	return c.IndicativeOrMarketPrice
}

func (c *CorporateActionPrice56) AddCashInLieuOfSharePrice() *PriceFormat45Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat45Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice56) AddCashValueForTax() *PriceFormat46Choice {
	c.CashValueForTax = new(PriceFormat46Choice)
	return c.CashValueForTax
}

func (c *CorporateActionPrice56) AddGenericCashPricePaidPerProduct() *PriceFormat44Choice {
	c.GenericCashPricePaidPerProduct = new(PriceFormat44Choice)
	return c.GenericCashPricePaidPerProduct
}

func (c *CorporateActionPrice56) AddGenericCashPriceReceivedPerProduct() *PriceFormat47Choice {
	c.GenericCashPriceReceivedPerProduct = new(PriceFormat47Choice)
	return c.GenericCashPriceReceivedPerProduct
}

// Specifies prices of a corporate action.
type CorporateActionPrice57 struct {

	// Maximum or cap price at which a holder can bid, for example, on a Dutch auction offer.
	MaximumPrice *PriceFormat44Choice `xml:"MaxPric,omitempty"`

	// Minimum or floor price at which a holder can bid, for example, on a Dutch auction offer.
	MinimumPrice *PriceFormat44Choice `xml:"MinPric,omitempty"`
}

func (c *CorporateActionPrice57) AddMaximumPrice() *PriceFormat44Choice {
	c.MaximumPrice = new(PriceFormat44Choice)
	return c.MaximumPrice
}

func (c *CorporateActionPrice57) AddMinimumPrice() *PriceFormat44Choice {
	c.MinimumPrice = new(PriceFormat44Choice)
	return c.MinimumPrice
}

// Specifies prices related to a corporate action option.
type CorporateActionPrice58 struct {

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat45Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Amount of money required per over-subscribed equity as defined by the issuer.
	OverSubscriptionDepositPrice *PriceFormat45Choice `xml:"OverSbcptDpstPric,omitempty"`
}

func (c *CorporateActionPrice58) AddCashInLieuOfSharePrice() *PriceFormat45Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat45Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice58) AddOverSubscriptionDepositPrice() *PriceFormat45Choice {
	c.OverSubscriptionDepositPrice = new(PriceFormat45Choice)
	return c.OverSubscriptionDepositPrice
}

// Specifies prices.
type CorporateActionPrice59 struct {

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat50Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Indicates whether the price is an indicative price or a market price.
	IndicativeOrMarketPrice *IndicativeOrMarketPrice8Choice `xml:"IndctvOrMktPric,omitempty"`

	// Cash value of resulting securities proceeds for tax calculation and/or reporting.
	CashValueForTax *AmountPrice2 `xml:"CshValForTax,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount or a number of points above an index, for example, reinvestment price, strike price and exercise price.
	GenericCashPricePaidPerProduct *PriceFormat51Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat48Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (c *CorporateActionPrice59) AddCashInLieuOfSharePrice() *PriceFormat50Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat50Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice59) AddIndicativeOrMarketPrice() *IndicativeOrMarketPrice8Choice {
	c.IndicativeOrMarketPrice = new(IndicativeOrMarketPrice8Choice)
	return c.IndicativeOrMarketPrice
}

func (c *CorporateActionPrice59) AddCashValueForTax() *AmountPrice2 {
	c.CashValueForTax = new(AmountPrice2)
	return c.CashValueForTax
}

func (c *CorporateActionPrice59) AddGenericCashPricePaidPerProduct() *PriceFormat51Choice {
	c.GenericCashPricePaidPerProduct = new(PriceFormat51Choice)
	return c.GenericCashPricePaidPerProduct
}

func (c *CorporateActionPrice59) AddGenericCashPriceReceivedPerProduct() *PriceFormat48Choice {
	c.GenericCashPriceReceivedPerProduct = new(PriceFormat48Choice)
	return c.GenericCashPriceReceivedPerProduct
}

// Specifies prices related to a corporate action option.
type CorporateActionPrice6 struct {

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat11Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// 1. Price at which security will be purchased/sold if warrant is exercised, either as an actual amount or a percentage.
	// 2. Price at which a bond is converted to underlying security either as an actual amount or a percentage.
	// 3. Strike price of an option, represented either as an actual amount, a percentage or a number of points above an index.
	ExercisePrice *PriceFormat8Choice `xml:"ExrcPric,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct []*PriceFormat9Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`

	// Amount of money required per over-subscribed equity as defined by the issuer.
	OverSubscriptionDepositPrice *PriceFormat11Choice `xml:"OverSbcptDpstPric,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount, for example, reinvestment price.
	GenericCashPricePaidPerProduct *PriceFormat11Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Amount included in the dividend/NAV that is identified as gains directly or indirectly derived from interest payments within the scope of the EU Savings directive.
	TaxableIncomePerDividendShare *AmountPrice3 `xml:"TaxblIncmPerDvddShr,omitempty"`
}

func (c *CorporateActionPrice6) AddCashInLieuOfSharePrice() *PriceFormat11Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat11Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice6) AddExercisePrice() *PriceFormat8Choice {
	c.ExercisePrice = new(PriceFormat8Choice)
	return c.ExercisePrice
}

func (c *CorporateActionPrice6) AddGenericCashPriceReceivedPerProduct() *PriceFormat9Choice {
	newValue := new(PriceFormat9Choice)
	c.GenericCashPriceReceivedPerProduct = append(c.GenericCashPriceReceivedPerProduct, newValue)
	return newValue
}

func (c *CorporateActionPrice6) AddOverSubscriptionDepositPrice() *PriceFormat11Choice {
	c.OverSubscriptionDepositPrice = new(PriceFormat11Choice)
	return c.OverSubscriptionDepositPrice
}

func (c *CorporateActionPrice6) AddGenericCashPricePaidPerProduct() *PriceFormat11Choice {
	c.GenericCashPricePaidPerProduct = new(PriceFormat11Choice)
	return c.GenericCashPricePaidPerProduct
}

func (c *CorporateActionPrice6) AddTaxableIncomePerDividendShare() *AmountPrice3 {
	c.TaxableIncomePerDividendShare = new(AmountPrice3)
	return c.TaxableIncomePerDividendShare
}

// Specifies prices related to a corporate action option.
type CorporateActionPrice60 struct {

	// Indicates whether the price is an indicative price or a market price.
	IndicativeOrMarketPrice *IndicativeOrMarketPrice8Choice `xml:"IndctvOrMktPric,omitempty"`

	// Initial issue price of a financial instrument.
	IssuePrice *PriceFormat50Choice `xml:"IssePric,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat49Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount, for example, reinvestment price.
	GenericCashPricePaidPerProduct *PriceFormat50Choice `xml:"GncCshPricPdPerPdct,omitempty"`
}

func (c *CorporateActionPrice60) AddIndicativeOrMarketPrice() *IndicativeOrMarketPrice8Choice {
	c.IndicativeOrMarketPrice = new(IndicativeOrMarketPrice8Choice)
	return c.IndicativeOrMarketPrice
}

func (c *CorporateActionPrice60) AddIssuePrice() *PriceFormat50Choice {
	c.IssuePrice = new(PriceFormat50Choice)
	return c.IssuePrice
}

func (c *CorporateActionPrice60) AddGenericCashPriceReceivedPerProduct() *PriceFormat49Choice {
	c.GenericCashPriceReceivedPerProduct = new(PriceFormat49Choice)
	return c.GenericCashPriceReceivedPerProduct
}

func (c *CorporateActionPrice60) AddGenericCashPricePaidPerProduct() *PriceFormat50Choice {
	c.GenericCashPricePaidPerProduct = new(PriceFormat50Choice)
	return c.GenericCashPricePaidPerProduct
}

// Specifies prices.
type CorporateActionPrice61 struct {

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat50Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Amount of money required per over-subscribed equity as defined by the issuer.
	OverSubscriptionDepositPrice *PriceFormat50Choice `xml:"OverSbcptDpstPric,omitempty"`
}

func (c *CorporateActionPrice61) AddCashInLieuOfSharePrice() *PriceFormat50Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat50Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice61) AddOverSubscriptionDepositPrice() *PriceFormat50Choice {
	c.OverSubscriptionDepositPrice = new(PriceFormat50Choice)
	return c.OverSubscriptionDepositPrice
}

// Specifies prices related to a corporate action option.
type CorporateActionPrice62 struct {

	// Indicates whether the price is an indicative price or a market price.
	IndicativeOrMarketPrice *IndicativeOrMarketPrice9Choice `xml:"IndctvOrMktPric,omitempty"`

	// Initial issue price of a financial instrument.
	IssuePrice *PriceFormat52Choice `xml:"IssePric,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat53Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount, for example, reinvestment price.
	GenericCashPricePaidPerProduct *PriceFormat52Choice `xml:"GncCshPricPdPerPdct,omitempty"`
}

func (c *CorporateActionPrice62) AddIndicativeOrMarketPrice() *IndicativeOrMarketPrice9Choice {
	c.IndicativeOrMarketPrice = new(IndicativeOrMarketPrice9Choice)
	return c.IndicativeOrMarketPrice
}

func (c *CorporateActionPrice62) AddIssuePrice() *PriceFormat52Choice {
	c.IssuePrice = new(PriceFormat52Choice)
	return c.IssuePrice
}

func (c *CorporateActionPrice62) AddGenericCashPriceReceivedPerProduct() *PriceFormat53Choice {
	c.GenericCashPriceReceivedPerProduct = new(PriceFormat53Choice)
	return c.GenericCashPriceReceivedPerProduct
}

func (c *CorporateActionPrice62) AddGenericCashPricePaidPerProduct() *PriceFormat52Choice {
	c.GenericCashPricePaidPerProduct = new(PriceFormat52Choice)
	return c.GenericCashPricePaidPerProduct
}

// Specifies prices.
type CorporateActionPrice63 struct {

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat52Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Amount of money required per over-subscribed equity as defined by the issuer.
	OverSubscriptionDepositPrice *PriceFormat52Choice `xml:"OverSbcptDpstPric,omitempty"`
}

func (c *CorporateActionPrice63) AddCashInLieuOfSharePrice() *PriceFormat52Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat52Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice63) AddOverSubscriptionDepositPrice() *PriceFormat52Choice {
	c.OverSubscriptionDepositPrice = new(PriceFormat52Choice)
	return c.OverSubscriptionDepositPrice
}

// Specifies prices.
type CorporateActionPrice64 struct {

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat52Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Indicates whether the price is an indicative price or a market price.
	IndicativeOrMarketPrice *IndicativeOrMarketPrice9Choice `xml:"IndctvOrMktPric,omitempty"`

	// Cash value of resulting securities proceeds for tax calculation and/or reporting.
	CashValueForTax *AmountPrice4 `xml:"CshValForTax,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount or a number of points above an index, for example, reinvestment price, strike price and exercise price.
	GenericCashPricePaidPerProduct *PriceFormat55Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat56Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (c *CorporateActionPrice64) AddCashInLieuOfSharePrice() *PriceFormat52Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat52Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice64) AddIndicativeOrMarketPrice() *IndicativeOrMarketPrice9Choice {
	c.IndicativeOrMarketPrice = new(IndicativeOrMarketPrice9Choice)
	return c.IndicativeOrMarketPrice
}

func (c *CorporateActionPrice64) AddCashValueForTax() *AmountPrice4 {
	c.CashValueForTax = new(AmountPrice4)
	return c.CashValueForTax
}

func (c *CorporateActionPrice64) AddGenericCashPricePaidPerProduct() *PriceFormat55Choice {
	c.GenericCashPricePaidPerProduct = new(PriceFormat55Choice)
	return c.GenericCashPricePaidPerProduct
}

func (c *CorporateActionPrice64) AddGenericCashPriceReceivedPerProduct() *PriceFormat56Choice {
	c.GenericCashPriceReceivedPerProduct = new(PriceFormat56Choice)
	return c.GenericCashPriceReceivedPerProduct
}

// Specifies prices related to a corporate action option.
type CorporateActionPrice65 struct {

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat57Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Amount of money required per over-subscribed equity as defined by the issuer.
	OverSubscriptionDepositPrice *PriceFormat57Choice `xml:"OverSbcptDpstPric,omitempty"`
}

func (c *CorporateActionPrice65) AddCashInLieuOfSharePrice() *PriceFormat57Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat57Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice65) AddOverSubscriptionDepositPrice() *PriceFormat57Choice {
	c.OverSubscriptionDepositPrice = new(PriceFormat57Choice)
	return c.OverSubscriptionDepositPrice
}

// Specifies prices related to a corporate action option.
type CorporateActionPrice66 struct {

	// Indicates whether the price is an indicative price or a market price.
	IndicativeOrMarketPrice *IndicativeOrMarketPrice11Choice `xml:"IndctvOrMktPric,omitempty"`

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat57Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Cash value of resulting securities proceeds for tax calculation and/or reporting.
	CashValueForTax *PriceFormat58Choice `xml:"CshValForTax,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount or a number of points above an index, for example, reinvestment price, strike price and exercise price.
	GenericCashPricePaidPerProduct *PriceFormat59Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat60Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (c *CorporateActionPrice66) AddIndicativeOrMarketPrice() *IndicativeOrMarketPrice11Choice {
	c.IndicativeOrMarketPrice = new(IndicativeOrMarketPrice11Choice)
	return c.IndicativeOrMarketPrice
}

func (c *CorporateActionPrice66) AddCashInLieuOfSharePrice() *PriceFormat57Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat57Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice66) AddCashValueForTax() *PriceFormat58Choice {
	c.CashValueForTax = new(PriceFormat58Choice)
	return c.CashValueForTax
}

func (c *CorporateActionPrice66) AddGenericCashPricePaidPerProduct() *PriceFormat59Choice {
	c.GenericCashPricePaidPerProduct = new(PriceFormat59Choice)
	return c.GenericCashPricePaidPerProduct
}

func (c *CorporateActionPrice66) AddGenericCashPriceReceivedPerProduct() *PriceFormat60Choice {
	c.GenericCashPriceReceivedPerProduct = new(PriceFormat60Choice)
	return c.GenericCashPriceReceivedPerProduct
}

// Specifies prices of a corporate action.
type CorporateActionPrice67 struct {

	// Maximum or cap price at which a holder can bid, for example, on a Dutch auction offer.
	MaximumPrice *PriceFormat59Choice `xml:"MaxPric,omitempty"`

	// Minimum or floor price at which a holder can bid, for example, on a Dutch auction offer.
	MinimumPrice *PriceFormat59Choice `xml:"MinPric,omitempty"`
}

func (c *CorporateActionPrice67) AddMaximumPrice() *PriceFormat59Choice {
	c.MaximumPrice = new(PriceFormat59Choice)
	return c.MaximumPrice
}

func (c *CorporateActionPrice67) AddMinimumPrice() *PriceFormat59Choice {
	c.MinimumPrice = new(PriceFormat59Choice)
	return c.MinimumPrice
}

// Specifies prices related to a corporate action option.
type CorporateActionPrice8 struct {

	// Indicates whether the price is an indicative price or a market price.
	IndicativeOrMarketPrice *IndicativeOrMarketPrice2Choice `xml:"IndctvOrMktPric,omitempty"`

	// 1. Price at which security will be purchased/sold if warrant is exercised, either as an actual amount or a percentage.
	// 2. Price at which a bond is converted to underlying security either as an actual amount or a percentage.
	// 3. Strike price of an option, represented either as an actual amount, a percentage or a number of points above an index.
	ExercisePrice *PriceFormat5Choice `xml:"ExrcPric,omitempty"`

	// Initial issue price of a financial instrument.
	IssuePrice *PriceFormat5Choice `xml:"IssePric,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat9Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount, for example, reinvestment price.
	GenericCashPricePaidPerProduct *PriceFormat5Choice `xml:"GncCshPricPdPerPdct,omitempty"`
}

func (c *CorporateActionPrice8) AddIndicativeOrMarketPrice() *IndicativeOrMarketPrice2Choice {
	c.IndicativeOrMarketPrice = new(IndicativeOrMarketPrice2Choice)
	return c.IndicativeOrMarketPrice
}

func (c *CorporateActionPrice8) AddExercisePrice() *PriceFormat5Choice {
	c.ExercisePrice = new(PriceFormat5Choice)
	return c.ExercisePrice
}

func (c *CorporateActionPrice8) AddIssuePrice() *PriceFormat5Choice {
	c.IssuePrice = new(PriceFormat5Choice)
	return c.IssuePrice
}

func (c *CorporateActionPrice8) AddGenericCashPriceReceivedPerProduct() *PriceFormat9Choice {
	c.GenericCashPriceReceivedPerProduct = new(PriceFormat9Choice)
	return c.GenericCashPriceReceivedPerProduct
}

func (c *CorporateActionPrice8) AddGenericCashPricePaidPerProduct() *PriceFormat5Choice {
	c.GenericCashPricePaidPerProduct = new(PriceFormat5Choice)
	return c.GenericCashPricePaidPerProduct
}

// Specifies prices.
type CorporateActionPrice9 struct {

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat5Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Indicates whether the price is an indicative price or a market price.
	IndicativeOrMarketPrice *IndicativeOrMarketPrice2Choice `xml:"IndctvOrMktPric,omitempty"`
}

func (c *CorporateActionPrice9) AddCashInLieuOfSharePrice() *PriceFormat5Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat5Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice9) AddIndicativeOrMarketPrice() *IndicativeOrMarketPrice2Choice {
	c.IndicativeOrMarketPrice = new(IndicativeOrMarketPrice2Choice)
	return c.IndicativeOrMarketPrice
}
