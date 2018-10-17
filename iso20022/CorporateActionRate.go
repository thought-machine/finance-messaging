package iso20022

// Specifies rates.
type CorporateActionRate1 struct {

	// Annual rate of a financial instrument.
	Interest *RateAndAmountFormat1Choice `xml:"Intrst,omitempty"`

	// Index rate related to the interest rate of the forthcoming interest payment.
	RelatedIndex *RateFormat1Choice `xml:"RltdIndx,omitempty"`

	// Percentage of securities the offeror/issuer will purchase or redeem under the terms of the event. This can be a number or the term "any and all".
	PercentageSought *RateFormat1Choice `xml:"PctgSght,omitempty"`

	// Rate of discount for securities purchased through a reinvestment scheme as compared to the current market price of security.
	ReinvestmentDiscountToMarket *RateFormat1Choice `xml:"RinvstmtDscntToMkt,omitempty"`

	// Margin allowed over or under a given rate.
	Spread *RateFormat1Choice `xml:"Sprd,omitempty"`

	// Acceptable price increment used for submitting a bid.
	BidInterval *AmountAndRateFormat3Choice `xml:"BidIntrvl,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	Charges *RateAndAmountFormat1Choice `xml:"Chrgs,omitempty"`
}

func (c *CorporateActionRate1) AddInterest() *RateAndAmountFormat1Choice {
	c.Interest = new(RateAndAmountFormat1Choice)
	return c.Interest
}

func (c *CorporateActionRate1) AddRelatedIndex() *RateFormat1Choice {
	c.RelatedIndex = new(RateFormat1Choice)
	return c.RelatedIndex
}

func (c *CorporateActionRate1) AddPercentageSought() *RateFormat1Choice {
	c.PercentageSought = new(RateFormat1Choice)
	return c.PercentageSought
}

func (c *CorporateActionRate1) AddReinvestmentDiscountToMarket() *RateFormat1Choice {
	c.ReinvestmentDiscountToMarket = new(RateFormat1Choice)
	return c.ReinvestmentDiscountToMarket
}

func (c *CorporateActionRate1) AddSpread() *RateFormat1Choice {
	c.Spread = new(RateFormat1Choice)
	return c.Spread
}

func (c *CorporateActionRate1) AddBidInterval() *AmountAndRateFormat3Choice {
	c.BidInterval = new(AmountAndRateFormat3Choice)
	return c.BidInterval
}

func (c *CorporateActionRate1) AddCharges() *RateAndAmountFormat1Choice {
	c.Charges = new(RateAndAmountFormat1Choice)
	return c.Charges
}

// Specifies rates related to a corporate action option.
type CorporateActionRate15 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat14Choice `xml:"AddtlTax,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat5Choice `xml:"GrssDvddRate,omitempty"`

	// Public index rate applied to the amount paid to adjust it to inflation.
	IndexFactor *RateAndAmountFormat14Choice `xml:"IndxFctr,omitempty"`

	// The actual interest rate used for the payment of the interest for the specified interest period.
	// Usage guideline: It is used to provide the applicable rate for the current payment, after all calculations have been performed, that is, application of period and method of interest computation.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat5Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Maximum percentage of shares available through the over subscription privilege, usually a percentage of the basic subscription shares, for example, an account owner subscribing to 100 shares may over subscribe to a maximum of 50 additional shares when the over subscription maximum is 50 percent.
	MaximumAllowedOversubscriptionRate *RateFormat6Choice `xml:"MaxAllwdOvrsbcptRate,omitempty"`

	// Proportionate allocation used for the offer.
	ProrationRate *RateFormat6Choice `xml:"PrratnRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxRate *RateFormat6Choice `xml:"WhldgTaxRate,omitempty"`

	// Percentage of the gross dividend rate on which tax must be paid .
	TaxRelatedRate []*RateTypeAndAmountAndStatus6 `xml:"TaxRltdRate,omitempty"`

	// Amount included in the dividend/NAV that is identified as gains directly or indirectly derived from interest payments within the scope of the EU Savings directive.
	TaxableIncomePerDividendShare []*RateTypeAndAmountAndStatus11 `xml:"TaxblIncmPerDvddShr,omitempty"`
}

func (c *CorporateActionRate15) AddAdditionalTax() *RateAndAmountFormat14Choice {
	c.AdditionalTax = new(RateAndAmountFormat14Choice)
	return c.AdditionalTax
}

func (c *CorporateActionRate15) AddGrossDividendRate() *GrossDividendRateFormat5Choice {
	newValue := new(GrossDividendRateFormat5Choice)
	c.GrossDividendRate = append(c.GrossDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate15) AddIndexFactor() *RateAndAmountFormat14Choice {
	c.IndexFactor = new(RateAndAmountFormat14Choice)
	return c.IndexFactor
}

func (c *CorporateActionRate15) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat5Choice {
	newValue := new(InterestRateUsedForPaymentFormat5Choice)
	c.InterestRateUsedForPayment = append(c.InterestRateUsedForPayment, newValue)
	return newValue
}

func (c *CorporateActionRate15) AddMaximumAllowedOversubscriptionRate() *RateFormat6Choice {
	c.MaximumAllowedOversubscriptionRate = new(RateFormat6Choice)
	return c.MaximumAllowedOversubscriptionRate
}

func (c *CorporateActionRate15) AddProrationRate() *RateFormat6Choice {
	c.ProrationRate = new(RateFormat6Choice)
	return c.ProrationRate
}

func (c *CorporateActionRate15) AddWithholdingTaxRate() *RateFormat6Choice {
	c.WithholdingTaxRate = new(RateFormat6Choice)
	return c.WithholdingTaxRate
}

func (c *CorporateActionRate15) AddTaxRelatedRate() *RateTypeAndAmountAndStatus6 {
	newValue := new(RateTypeAndAmountAndStatus6)
	c.TaxRelatedRate = append(c.TaxRelatedRate, newValue)
	return newValue
}

func (c *CorporateActionRate15) AddTaxableIncomePerDividendShare() *RateTypeAndAmountAndStatus11 {
	newValue := new(RateTypeAndAmountAndStatus11)
	c.TaxableIncomePerDividendShare = append(c.TaxableIncomePerDividendShare, newValue)
	return newValue
}

// Specifies rates of a corporate action.
type CorporateActionRate16 struct {

	// Annual rate of a financial instrument.
	Interest *RateAndAmountFormat14Choice `xml:"Intrst,omitempty"`

	// Percentage of securities the offeror/issuer will purchase or redeem under the terms of the event. This can be a number or the term "any and all".
	PercentageSought *RateFormat5Choice `xml:"PctgSght,omitempty"`

	// Index rate related to the interest rate of the forthcoming interest payment.
	RelatedIndex *RateFormat6Choice `xml:"RltdIndx,omitempty"`

	// Margin allowed over or under a given rate.
	Spread *RateFormat6Choice `xml:"Sprd,omitempty"`

	// Acceptable price increment used for submitting a bid.
	BidInterval *RateAndAmountFormat14Choice `xml:"BidIntrvl,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *RateFormat3Choice `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *RateFormat3Choice `xml:"NxtFctr,omitempty"`

	// Rate of discount for securities purchased through a reinvestment scheme as compared to the current market price of security.
	ReinvestmentDiscountRateToMarket *RateFormat6Choice `xml:"RinvstmtDscntRateToMkt,omitempty"`
}

func (c *CorporateActionRate16) AddInterest() *RateAndAmountFormat14Choice {
	c.Interest = new(RateAndAmountFormat14Choice)
	return c.Interest
}

func (c *CorporateActionRate16) AddPercentageSought() *RateFormat5Choice {
	c.PercentageSought = new(RateFormat5Choice)
	return c.PercentageSought
}

func (c *CorporateActionRate16) AddRelatedIndex() *RateFormat6Choice {
	c.RelatedIndex = new(RateFormat6Choice)
	return c.RelatedIndex
}

func (c *CorporateActionRate16) AddSpread() *RateFormat6Choice {
	c.Spread = new(RateFormat6Choice)
	return c.Spread
}

func (c *CorporateActionRate16) AddBidInterval() *RateAndAmountFormat14Choice {
	c.BidInterval = new(RateAndAmountFormat14Choice)
	return c.BidInterval
}

func (c *CorporateActionRate16) AddPreviousFactor() *RateFormat3Choice {
	c.PreviousFactor = new(RateFormat3Choice)
	return c.PreviousFactor
}

func (c *CorporateActionRate16) AddNextFactor() *RateFormat3Choice {
	c.NextFactor = new(RateFormat3Choice)
	return c.NextFactor
}

func (c *CorporateActionRate16) AddReinvestmentDiscountRateToMarket() *RateFormat6Choice {
	c.ReinvestmentDiscountRateToMarket = new(RateFormat6Choice)
	return c.ReinvestmentDiscountRateToMarket
}

// Specifies security rate details.
type CorporateActionRate17 struct {

	// Quantity of additional intermediate securities/new equities awarded for a given quantity of securities derived from subscription.
	AdditionalQuantityForSubscribedResultantSecurities *RatioFormat11Choice `xml:"AddtlQtyForSbcbdRsltntScties,omitempty"`

	// Quantity of additional securities for a given quantity of underlying securities where underlying securities are not exchanged or debited, for example, 1 for 1: 1 new equity credited for every 1 underlying equity = 2 resulting equities.
	AdditionalQuantityForExistingSecurities *RatioFormat11Choice `xml:"AddtlQtyForExstgScties,omitempty"`

	// Quantity of new securities for a given quantity of underlying securities, where the underlying securities will be exchanged or debited, for example, 2 for 1: 2 new equities credited for every 1 underlying equity debited = 2 resulting equities.
	NewToOld *RatioFormat12Choice `xml:"NewToOd,omitempty"`

	// Rate used to determine the cash consideration split across outturn settlement transactions that are the result of a transformation of the parent transaction.
	TransformationRate *PercentageRate `xml:"TrfrmatnRate,omitempty"`
}

func (c *CorporateActionRate17) AddAdditionalQuantityForSubscribedResultantSecurities() *RatioFormat11Choice {
	c.AdditionalQuantityForSubscribedResultantSecurities = new(RatioFormat11Choice)
	return c.AdditionalQuantityForSubscribedResultantSecurities
}

func (c *CorporateActionRate17) AddAdditionalQuantityForExistingSecurities() *RatioFormat11Choice {
	c.AdditionalQuantityForExistingSecurities = new(RatioFormat11Choice)
	return c.AdditionalQuantityForExistingSecurities
}

func (c *CorporateActionRate17) AddNewToOld() *RatioFormat12Choice {
	c.NewToOld = new(RatioFormat12Choice)
	return c.NewToOld
}

func (c *CorporateActionRate17) SetTransformationRate(value string) {
	c.TransformationRate = (*PercentageRate)(&value)
}

// Specifies rates.
type CorporateActionRate2 struct {

	// Percentage of a cash distribution that will be withheld by a tax authority.
	WithholdingTax *RateFormat1Choice `xml:"WhldgTax,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction in which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTax *RateAndAmountFormat1Choice `xml:"WhldgOfFrgnTax,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction in which the account owner is located, for which relief at source and/or reclaim may be possible.
	WithholdingOfLocalTax *RateAndAmountFormat1Choice `xml:"WhldgOfLclTax,omitempty"`

	// Local tax (ZAS Anrechnungsbetrag) subject to interest down payment tax (proportion of interest liable for interest down payment tax/interim profit that is not covered by the tax exempt amount).
	GermanLocalTax1 *RateAndAmountFormat1Choice `xml:"GrmnLclTax1,omitempty"`

	// Local tax (ZAS Pflichtige Zinsen) interest liable for interest down payment tax (proportion of gross interest per unit/interim profits that is not covered by the credit in the interest pool).
	GermanLocalTax2 *RateAndAmountFormat1Choice `xml:"GrmnLclTax2,omitempty"`

	// Local tax (Zinstopf) offset interest per unit against tax exempt amount (variation to offset interest per unit in relation to tax exempt amount).
	GermanLocalTax3 *RateAndAmountFormat1Choice `xml:"GrmnLclTax3,omitempty"`

	// Local tax (Ertrag Besitzanteilig) yield liable for interest down payment tax.
	GermanLocalTax4 *RateAndAmountFormat1Choice `xml:"GrmnLclTax4,omitempty"`

	// Taxation applied on an amount clearly identified as an income.
	TaxOnIncome *RateFormat1Choice `xml:"TaxOnIncm,omitempty"`

	// Taxation applied on an amount clearly identified as capital profits, capital gains.
	TaxOnProfit *RateFormat1Choice `xml:"TaxOnPrft,omitempty"`

	// Percentage of cash that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaim *RateFormat1Choice `xml:"TaxRclm,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *RateFormat1Choice `xml:"FsclStmp,omitempty"`

	// Proportionate allocation used for the offer.
	Proration *RateFormat1Choice `xml:"Prratn,omitempty"`

	// Quantity of new securities for a given quantity of underlying securities, where the underlying securities will be exchanged or debited, eg, 2 for 1: 2 new equities credited for every 1 underlying equity debited = 2 resulting equities.
	NewToOld *RatioFormat2Choice `xml:"NewToOd,omitempty"`

	// Quantity of new equities that will be derived by the exercise of a given quantity of intermediate securities.
	NewSecuritiesToUnderlyingSecurities *RatioFormat2Choice `xml:"NewSctiesToUndrlygScties,omitempty"`

	// Quantity of additional securities for a given quantity of underlying securities where underlying securities are not exchanged or debited, eg, 1 for 1: 1 new
	// equity credited for every 1 underlying equity = 2 resulting equities.
	AdditionalQuantityForExistingSecurities *RatioFormat1Choice `xml:"AddtlQtyForExstgScties,omitempty"`

	// Quantity of additional intermediate securities/new equities awarded for a given quantity of securities derived from subscription.
	AdditionalQuantityForSubscribedResultantSecurities *RatioFormat1Choice `xml:"AddtlQtyForSbcbdRsltntScties,omitempty"`

	// Percentage of the gross dividend rate on which tax must be paid .
	RelatedTax *RelatedTaxType1 `xml:"RltdTax,omitempty"`

	// Rate per share to which a non-resident is entitled.
	NonResidentRate *RateAndAmountFormat1Choice `xml:"NonResdtRate,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	Charges *RateAndAmountFormat1Choice `xml:"Chrgs,omitempty"`

	// The actual interest rate used for the payment of the interest for the specified interest period.
	InterestForUsedPayment *RateAndAmountFormat1Choice `xml:"IntrstForUsdPmt,omitempty"`

	// Public index rate applied to the amount paid to adjust it to inflation.
	IndexFactor *RateAndAmountFormat1Choice `xml:"IndxFctr,omitempty"`

	// Rate resulting from a fully franked dividend paid by a company; rate includes tax credit for companies that have made sufficient tax payments during fiscal period.
	FullyFranked *RateAndAmountFormat1Choice `xml:"FullyFrnkd,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividend *GrossDividendRate1Choice `xml:"GrssDvdd,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividend *NetDividendRate1Choice `xml:"NetDvdd,omitempty"`

	// Dividend is final.
	FinalDividend *AmountAndRateFormat2Choice `xml:"FnlDvdd,omitempty"`

	// Dividend is provisional.
	ProvisionalDividend *AmountAndRateFormat2Choice `xml:"PrvsnlDvdd,omitempty"`

	// Rate of the cash premium made available if the securities holder consents or participates to an event, e.g. consent fees.
	CashIncentive *RateFormat1Choice `xml:"CshIncntiv,omitempty"`

	// Cash rate made available in an offer in order to encourage participation in the offer.
	SolicitationFee *RateFormat1Choice `xml:"SlctnFee,omitempty"`

	// A maximum percentage of shares available through the over subscription privilege, usually a percentage of the basic subscription shares, eg, an account owner subscribing to 100 shares may over subscribe to a maximum of 50 additional shares when the over subscription maximum is 50%.
	MaximumAllowedOversubscription *RateFormat1Choice `xml:"MaxAllwdOvrsbcpt,omitempty"`

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat1Choice `xml:"AddtlTax,omitempty"`

	// Amount in its original currency when conversion from/into another currency has occurred.
	OriginalAmount *ActiveCurrencyAndAmount `xml:"OrgnlAmt,omitempty"`

	// Provides information about a foreign exchange.
	ExchangeRate *ForeignExchangeTerms8 `xml:"XchgRate,omitempty"`

	// Rate applicable to the event announced, eg, redemption rate for a redemption event.
	ApplicableRate *RateFormat1Choice `xml:"AplblRate,omitempty"`
}

func (c *CorporateActionRate2) AddWithholdingTax() *RateFormat1Choice {
	c.WithholdingTax = new(RateFormat1Choice)
	return c.WithholdingTax
}

func (c *CorporateActionRate2) AddWithholdingOfForeignTax() *RateAndAmountFormat1Choice {
	c.WithholdingOfForeignTax = new(RateAndAmountFormat1Choice)
	return c.WithholdingOfForeignTax
}

func (c *CorporateActionRate2) AddWithholdingOfLocalTax() *RateAndAmountFormat1Choice {
	c.WithholdingOfLocalTax = new(RateAndAmountFormat1Choice)
	return c.WithholdingOfLocalTax
}

func (c *CorporateActionRate2) AddGermanLocalTax1() *RateAndAmountFormat1Choice {
	c.GermanLocalTax1 = new(RateAndAmountFormat1Choice)
	return c.GermanLocalTax1
}

func (c *CorporateActionRate2) AddGermanLocalTax2() *RateAndAmountFormat1Choice {
	c.GermanLocalTax2 = new(RateAndAmountFormat1Choice)
	return c.GermanLocalTax2
}

func (c *CorporateActionRate2) AddGermanLocalTax3() *RateAndAmountFormat1Choice {
	c.GermanLocalTax3 = new(RateAndAmountFormat1Choice)
	return c.GermanLocalTax3
}

func (c *CorporateActionRate2) AddGermanLocalTax4() *RateAndAmountFormat1Choice {
	c.GermanLocalTax4 = new(RateAndAmountFormat1Choice)
	return c.GermanLocalTax4
}

func (c *CorporateActionRate2) AddTaxOnIncome() *RateFormat1Choice {
	c.TaxOnIncome = new(RateFormat1Choice)
	return c.TaxOnIncome
}

func (c *CorporateActionRate2) AddTaxOnProfit() *RateFormat1Choice {
	c.TaxOnProfit = new(RateFormat1Choice)
	return c.TaxOnProfit
}

func (c *CorporateActionRate2) AddTaxReclaim() *RateFormat1Choice {
	c.TaxReclaim = new(RateFormat1Choice)
	return c.TaxReclaim
}

func (c *CorporateActionRate2) AddFiscalStamp() *RateFormat1Choice {
	c.FiscalStamp = new(RateFormat1Choice)
	return c.FiscalStamp
}

func (c *CorporateActionRate2) AddProration() *RateFormat1Choice {
	c.Proration = new(RateFormat1Choice)
	return c.Proration
}

func (c *CorporateActionRate2) AddNewToOld() *RatioFormat2Choice {
	c.NewToOld = new(RatioFormat2Choice)
	return c.NewToOld
}

func (c *CorporateActionRate2) AddNewSecuritiesToUnderlyingSecurities() *RatioFormat2Choice {
	c.NewSecuritiesToUnderlyingSecurities = new(RatioFormat2Choice)
	return c.NewSecuritiesToUnderlyingSecurities
}

func (c *CorporateActionRate2) AddAdditionalQuantityForExistingSecurities() *RatioFormat1Choice {
	c.AdditionalQuantityForExistingSecurities = new(RatioFormat1Choice)
	return c.AdditionalQuantityForExistingSecurities
}

func (c *CorporateActionRate2) AddAdditionalQuantityForSubscribedResultantSecurities() *RatioFormat1Choice {
	c.AdditionalQuantityForSubscribedResultantSecurities = new(RatioFormat1Choice)
	return c.AdditionalQuantityForSubscribedResultantSecurities
}

func (c *CorporateActionRate2) AddRelatedTax() *RelatedTaxType1 {
	c.RelatedTax = new(RelatedTaxType1)
	return c.RelatedTax
}

func (c *CorporateActionRate2) AddNonResidentRate() *RateAndAmountFormat1Choice {
	c.NonResidentRate = new(RateAndAmountFormat1Choice)
	return c.NonResidentRate
}

func (c *CorporateActionRate2) AddCharges() *RateAndAmountFormat1Choice {
	c.Charges = new(RateAndAmountFormat1Choice)
	return c.Charges
}

func (c *CorporateActionRate2) AddInterestForUsedPayment() *RateAndAmountFormat1Choice {
	c.InterestForUsedPayment = new(RateAndAmountFormat1Choice)
	return c.InterestForUsedPayment
}

func (c *CorporateActionRate2) AddIndexFactor() *RateAndAmountFormat1Choice {
	c.IndexFactor = new(RateAndAmountFormat1Choice)
	return c.IndexFactor
}

func (c *CorporateActionRate2) AddFullyFranked() *RateAndAmountFormat1Choice {
	c.FullyFranked = new(RateAndAmountFormat1Choice)
	return c.FullyFranked
}

func (c *CorporateActionRate2) AddGrossDividend() *GrossDividendRate1Choice {
	c.GrossDividend = new(GrossDividendRate1Choice)
	return c.GrossDividend
}

func (c *CorporateActionRate2) AddNetDividend() *NetDividendRate1Choice {
	c.NetDividend = new(NetDividendRate1Choice)
	return c.NetDividend
}

func (c *CorporateActionRate2) AddFinalDividend() *AmountAndRateFormat2Choice {
	c.FinalDividend = new(AmountAndRateFormat2Choice)
	return c.FinalDividend
}

func (c *CorporateActionRate2) AddProvisionalDividend() *AmountAndRateFormat2Choice {
	c.ProvisionalDividend = new(AmountAndRateFormat2Choice)
	return c.ProvisionalDividend
}

func (c *CorporateActionRate2) AddCashIncentive() *RateFormat1Choice {
	c.CashIncentive = new(RateFormat1Choice)
	return c.CashIncentive
}

func (c *CorporateActionRate2) AddSolicitationFee() *RateFormat1Choice {
	c.SolicitationFee = new(RateFormat1Choice)
	return c.SolicitationFee
}

func (c *CorporateActionRate2) AddMaximumAllowedOversubscription() *RateFormat1Choice {
	c.MaximumAllowedOversubscription = new(RateFormat1Choice)
	return c.MaximumAllowedOversubscription
}

func (c *CorporateActionRate2) AddAdditionalTax() *RateAndAmountFormat1Choice {
	c.AdditionalTax = new(RateAndAmountFormat1Choice)
	return c.AdditionalTax
}

func (c *CorporateActionRate2) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionRate2) AddExchangeRate() *ForeignExchangeTerms8 {
	c.ExchangeRate = new(ForeignExchangeTerms8)
	return c.ExchangeRate
}

func (c *CorporateActionRate2) AddApplicableRate() *RateFormat1Choice {
	c.ApplicableRate = new(RateFormat1Choice)
	return c.ApplicableRate
}

// Specifies rates.
type CorporateActionRate20 struct {

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat2Choice `xml:"GrssDvddRate,omitempty"`

	// Public index rate applied to the amount paid to adjust it to inflation.
	IndexFactor *RateAndAmountFormat5Choice `xml:"IndxFctr,omitempty"`

	// Actual interest rate used for the payment of the interest for the specified interest period.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat2Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// A maximum percentage of shares available through the over subscription privilege, usually a percentage of the basic subscription shares, for example, an account owner subscribing to 100 shares may over subscribe to a maximum of 50 additional shares when the over subscription maximum is 50 percent.
	MaximumAllowedOversubscriptionRate *PercentageRate `xml:"MaxAllwdOvrsbcptRate,omitempty"`

	// Proportionate allocation used for the offer.
	ProrationRate *PercentageRate `xml:"PrratnRate,omitempty"`

	// Percentage of the gross dividend rate on which tax must be paid .
	TaxRelatedRate []*RateTypeAndAmountAndStatus6 `xml:"TaxRltdRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxRate *PercentageRate `xml:"WhldgTaxRate,omitempty"`

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat5Choice `xml:"AddtlTax,omitempty"`

	// Amount included in the dividend/NAV that is identified as gains directly or indirectly derived from interest payments within the scope of the EU Savings directive.
	TaxableIncomePerDividendShare []*RateTypeAndAmountAndStatus11 `xml:"TaxblIncmPerDvddShr,omitempty"`
}

func (c *CorporateActionRate20) AddGrossDividendRate() *GrossDividendRateFormat2Choice {
	newValue := new(GrossDividendRateFormat2Choice)
	c.GrossDividendRate = append(c.GrossDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate20) AddIndexFactor() *RateAndAmountFormat5Choice {
	c.IndexFactor = new(RateAndAmountFormat5Choice)
	return c.IndexFactor
}

func (c *CorporateActionRate20) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat2Choice {
	newValue := new(InterestRateUsedForPaymentFormat2Choice)
	c.InterestRateUsedForPayment = append(c.InterestRateUsedForPayment, newValue)
	return newValue
}

func (c *CorporateActionRate20) SetMaximumAllowedOversubscriptionRate(value string) {
	c.MaximumAllowedOversubscriptionRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate20) SetProrationRate(value string) {
	c.ProrationRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate20) AddTaxRelatedRate() *RateTypeAndAmountAndStatus6 {
	newValue := new(RateTypeAndAmountAndStatus6)
	c.TaxRelatedRate = append(c.TaxRelatedRate, newValue)
	return newValue
}

func (c *CorporateActionRate20) SetWithholdingTaxRate(value string) {
	c.WithholdingTaxRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate20) AddAdditionalTax() *RateAndAmountFormat5Choice {
	c.AdditionalTax = new(RateAndAmountFormat5Choice)
	return c.AdditionalTax
}

func (c *CorporateActionRate20) AddTaxableIncomePerDividendShare() *RateTypeAndAmountAndStatus11 {
	newValue := new(RateTypeAndAmountAndStatus11)
	c.TaxableIncomePerDividendShare = append(c.TaxableIncomePerDividendShare, newValue)
	return newValue
}

// Specifies rate details.
type CorporateActionRate21 struct {

	// Quantity of additional intermediate securities/new equities awarded for a given quantity of securities derived from subscription.
	AdditionalQuantityForSubscribedResultantSecurities *RatioFormat3Choice `xml:"AddtlQtyForSbcbdRsltntScties,omitempty"`

	// Quantity of additional securities for a given quantity of underlying securities where underlying securities are not exchanged or debited, for example, 1 for 1: 1 new equity credited for every 1 underlying equity = 2 resulting equities.
	AdditionalQuantityForExistingSecurities *RatioFormat3Choice `xml:"AddtlQtyForExstgScties,omitempty"`

	// Quantity of new securities for a given quantity of underlying securities, where the underlying securities will be exchanged or debited, for example, 2 for 1: 2 new equities credited for every 1 underlying equity debited = 2 resulting equities.
	NewToOld *RatioFormat4Choice `xml:"NewToOd,omitempty"`
}

func (c *CorporateActionRate21) AddAdditionalQuantityForSubscribedResultantSecurities() *RatioFormat3Choice {
	c.AdditionalQuantityForSubscribedResultantSecurities = new(RatioFormat3Choice)
	return c.AdditionalQuantityForSubscribedResultantSecurities
}

func (c *CorporateActionRate21) AddAdditionalQuantityForExistingSecurities() *RatioFormat3Choice {
	c.AdditionalQuantityForExistingSecurities = new(RatioFormat3Choice)
	return c.AdditionalQuantityForExistingSecurities
}

func (c *CorporateActionRate21) AddNewToOld() *RatioFormat4Choice {
	c.NewToOld = new(RatioFormat4Choice)
	return c.NewToOld
}

// Specifies rates related to a corporate action option.
type CorporateActionRate25 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat14Choice `xml:"AddtlTax,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat5Choice `xml:"GrssDvddRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat5Choice `xml:"NetDvddRate,omitempty"`

	// Public index rate applied to the amount paid to adjust it to inflation.
	IndexFactor *RateAndAmountFormat14Choice `xml:"IndxFctr,omitempty"`

	// The actual interest rate used for the payment of the interest for the specified interest period.
	// Usage guideline: It is used to provide the applicable rate for the current payment, after all calculations have been performed, that is, application of period and method of interest computation.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat5Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Maximum percentage of shares available through the over subscription privilege, usually a percentage of the basic subscription shares, for example, an account owner subscribing to 100 shares may over subscribe to a maximum of 50 additional shares when the over subscription maximum is 50 percent.
	MaximumAllowedOversubscriptionRate *RateFormat6Choice `xml:"MaxAllwdOvrsbcptRate,omitempty"`

	// Proportionate allocation used for the offer.
	ProrationRate *RateFormat6Choice `xml:"PrratnRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxRate *RateFormat6Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction to which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTax *RateAndAmountFormat14Choice `xml:"WhldgOfFrgnTax,omitempty"`

	// Percentage of the gross dividend rate on which tax must be paid .
	TaxRelatedRate []*RateTypeAndAmountAndStatus6 `xml:"TaxRltdRate,omitempty"`

	// Amount included in the dividend/NAV that is identified as gains directly or indirectly derived from interest payments within the scope of the EU Savings directive.
	TaxableIncomePerDividendShare []*RateTypeAndAmountAndStatus11 `xml:"TaxblIncmPerDvddShr,omitempty"`
}

func (c *CorporateActionRate25) AddAdditionalTax() *RateAndAmountFormat14Choice {
	c.AdditionalTax = new(RateAndAmountFormat14Choice)
	return c.AdditionalTax
}

func (c *CorporateActionRate25) AddGrossDividendRate() *GrossDividendRateFormat5Choice {
	newValue := new(GrossDividendRateFormat5Choice)
	c.GrossDividendRate = append(c.GrossDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate25) AddNetDividendRate() *NetDividendRateFormat5Choice {
	newValue := new(NetDividendRateFormat5Choice)
	c.NetDividendRate = append(c.NetDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate25) AddIndexFactor() *RateAndAmountFormat14Choice {
	c.IndexFactor = new(RateAndAmountFormat14Choice)
	return c.IndexFactor
}

func (c *CorporateActionRate25) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat5Choice {
	newValue := new(InterestRateUsedForPaymentFormat5Choice)
	c.InterestRateUsedForPayment = append(c.InterestRateUsedForPayment, newValue)
	return newValue
}

func (c *CorporateActionRate25) AddMaximumAllowedOversubscriptionRate() *RateFormat6Choice {
	c.MaximumAllowedOversubscriptionRate = new(RateFormat6Choice)
	return c.MaximumAllowedOversubscriptionRate
}

func (c *CorporateActionRate25) AddProrationRate() *RateFormat6Choice {
	c.ProrationRate = new(RateFormat6Choice)
	return c.ProrationRate
}

func (c *CorporateActionRate25) AddWithholdingTaxRate() *RateFormat6Choice {
	c.WithholdingTaxRate = new(RateFormat6Choice)
	return c.WithholdingTaxRate
}

func (c *CorporateActionRate25) AddWithholdingOfForeignTax() *RateAndAmountFormat14Choice {
	c.WithholdingOfForeignTax = new(RateAndAmountFormat14Choice)
	return c.WithholdingOfForeignTax
}

func (c *CorporateActionRate25) AddTaxRelatedRate() *RateTypeAndAmountAndStatus6 {
	newValue := new(RateTypeAndAmountAndStatus6)
	c.TaxRelatedRate = append(c.TaxRelatedRate, newValue)
	return newValue
}

func (c *CorporateActionRate25) AddTaxableIncomePerDividendShare() *RateTypeAndAmountAndStatus11 {
	newValue := new(RateTypeAndAmountAndStatus11)
	c.TaxableIncomePerDividendShare = append(c.TaxableIncomePerDividendShare, newValue)
	return newValue
}

// Specifies rates.
type CorporateActionRate26 struct {

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat2Choice `xml:"GrssDvddRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat7Choice `xml:"NetDvddRate,omitempty"`

	// Public index rate applied to the amount paid to adjust it to inflation.
	IndexFactor *RateAndAmountFormat5Choice `xml:"IndxFctr,omitempty"`

	// Actual interest rate used for the payment of the interest for the specified interest period.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat2Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// A maximum percentage of shares available through the over subscription privilege, usually a percentage of the basic subscription shares, for example, an account owner subscribing to 100 shares may over subscribe to a maximum of 50 additional shares when the over subscription maximum is 50 percent.
	MaximumAllowedOversubscriptionRate *PercentageRate `xml:"MaxAllwdOvrsbcptRate,omitempty"`

	// Proportionate allocation used for the offer.
	ProrationRate *PercentageRate `xml:"PrratnRate,omitempty"`

	// Percentage of the gross dividend rate on which tax must be paid .
	TaxRelatedRate []*RateTypeAndAmountAndStatus6 `xml:"TaxRltdRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxRate *PercentageRate `xml:"WhldgTaxRate,omitempty"`

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat5Choice `xml:"AddtlTax,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction to which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTax *RateAndAmountFormat5Choice `xml:"WhldgOfFrgnTax,omitempty"`

	// Amount included in the dividend/NAV that is identified as gains directly or indirectly derived from interest payments within the scope of the EU Savings directive.
	TaxableIncomePerDividendShare []*RateTypeAndAmountAndStatus11 `xml:"TaxblIncmPerDvddShr,omitempty"`
}

func (c *CorporateActionRate26) AddGrossDividendRate() *GrossDividendRateFormat2Choice {
	newValue := new(GrossDividendRateFormat2Choice)
	c.GrossDividendRate = append(c.GrossDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate26) AddNetDividendRate() *NetDividendRateFormat7Choice {
	newValue := new(NetDividendRateFormat7Choice)
	c.NetDividendRate = append(c.NetDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate26) AddIndexFactor() *RateAndAmountFormat5Choice {
	c.IndexFactor = new(RateAndAmountFormat5Choice)
	return c.IndexFactor
}

func (c *CorporateActionRate26) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat2Choice {
	newValue := new(InterestRateUsedForPaymentFormat2Choice)
	c.InterestRateUsedForPayment = append(c.InterestRateUsedForPayment, newValue)
	return newValue
}

func (c *CorporateActionRate26) SetMaximumAllowedOversubscriptionRate(value string) {
	c.MaximumAllowedOversubscriptionRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate26) SetProrationRate(value string) {
	c.ProrationRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate26) AddTaxRelatedRate() *RateTypeAndAmountAndStatus6 {
	newValue := new(RateTypeAndAmountAndStatus6)
	c.TaxRelatedRate = append(c.TaxRelatedRate, newValue)
	return newValue
}

func (c *CorporateActionRate26) SetWithholdingTaxRate(value string) {
	c.WithholdingTaxRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate26) AddAdditionalTax() *RateAndAmountFormat5Choice {
	c.AdditionalTax = new(RateAndAmountFormat5Choice)
	return c.AdditionalTax
}

func (c *CorporateActionRate26) AddWithholdingOfForeignTax() *RateAndAmountFormat5Choice {
	c.WithholdingOfForeignTax = new(RateAndAmountFormat5Choice)
	return c.WithholdingOfForeignTax
}

func (c *CorporateActionRate26) AddTaxableIncomePerDividendShare() *RateTypeAndAmountAndStatus11 {
	newValue := new(RateTypeAndAmountAndStatus11)
	c.TaxableIncomePerDividendShare = append(c.TaxableIncomePerDividendShare, newValue)
	return newValue
}

// Specifies rates of a corporate action.
type CorporateActionRate27 struct {

	// Annual rate of a financial instrument.
	Interest *RateAndAmountFormat14Choice `xml:"Intrst,omitempty"`

	// Percentage of securities the offeror/issuer will purchase or redeem under the terms of the event.
	PercentageSought *RateFormat7Choice `xml:"PctgSght,omitempty"`

	// Index rate related to the interest rate of the forthcoming interest payment.
	RelatedIndex *RateFormat6Choice `xml:"RltdIndx,omitempty"`

	// Margin allowed over or under a given rate.
	Spread *RateFormat6Choice `xml:"Sprd,omitempty"`

	// Acceptable price increment used for submitting a bid.
	BidInterval *RateAndAmountFormat14Choice `xml:"BidIntrvl,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *RateFormat3Choice `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *RateFormat3Choice `xml:"NxtFctr,omitempty"`

	// Rate of discount for securities purchased through a reinvestment scheme as compared to the current market price of security.
	ReinvestmentDiscountRateToMarket *RateFormat6Choice `xml:"RinvstmtDscntRateToMkt,omitempty"`

	// For structured security issues where there is a set schedule of principal and interest payments for the life of the issue, this is the difference between the actual rate of the interest payment and the expected or scheduled rate of the interest payment .
	InterestShortfall *RateAndAmountFormat12Choice `xml:"IntrstShrtfll,omitempty"`

	// For structured security issues where there is a set schedule of principal and interest payments for the life of the issue, this is the difference between the actual rate of the capital or principal repayment and the scheduled capital repayment.
	RealisedLoss *RateAndAmountFormat12Choice `xml:"RealsdLoss,omitempty"`
}

func (c *CorporateActionRate27) AddInterest() *RateAndAmountFormat14Choice {
	c.Interest = new(RateAndAmountFormat14Choice)
	return c.Interest
}

func (c *CorporateActionRate27) AddPercentageSought() *RateFormat7Choice {
	c.PercentageSought = new(RateFormat7Choice)
	return c.PercentageSought
}

func (c *CorporateActionRate27) AddRelatedIndex() *RateFormat6Choice {
	c.RelatedIndex = new(RateFormat6Choice)
	return c.RelatedIndex
}

func (c *CorporateActionRate27) AddSpread() *RateFormat6Choice {
	c.Spread = new(RateFormat6Choice)
	return c.Spread
}

func (c *CorporateActionRate27) AddBidInterval() *RateAndAmountFormat14Choice {
	c.BidInterval = new(RateAndAmountFormat14Choice)
	return c.BidInterval
}

func (c *CorporateActionRate27) AddPreviousFactor() *RateFormat3Choice {
	c.PreviousFactor = new(RateFormat3Choice)
	return c.PreviousFactor
}

func (c *CorporateActionRate27) AddNextFactor() *RateFormat3Choice {
	c.NextFactor = new(RateFormat3Choice)
	return c.NextFactor
}

func (c *CorporateActionRate27) AddReinvestmentDiscountRateToMarket() *RateFormat6Choice {
	c.ReinvestmentDiscountRateToMarket = new(RateFormat6Choice)
	return c.ReinvestmentDiscountRateToMarket
}

func (c *CorporateActionRate27) AddInterestShortfall() *RateAndAmountFormat12Choice {
	c.InterestShortfall = new(RateAndAmountFormat12Choice)
	return c.InterestShortfall
}

func (c *CorporateActionRate27) AddRealisedLoss() *RateAndAmountFormat12Choice {
	c.RealisedLoss = new(RateAndAmountFormat12Choice)
	return c.RealisedLoss
}

// Specifies security rate details.
type CorporateActionRate28 struct {

	// Quantity of additional intermediate securities/new equities awarded for a given quantity of securities derived from subscription.
	AdditionalQuantityForSubscribedResultantSecurities *RatioFormat11Choice `xml:"AddtlQtyForSbcbdRsltntScties,omitempty"`

	// Quantity of additional securities for a given quantity of underlying securities where underlying securities are not exchanged or debited, for example, 1 for 1: 1 new equity credited for every 1 underlying equity = 2 resulting equities.
	AdditionalQuantityForExistingSecurities *RatioFormat11Choice `xml:"AddtlQtyForExstgScties,omitempty"`

	// Quantity of new securities for a given quantity of underlying securities, where the underlying securities will be exchanged or debited, for example, 2 for 1: 2 new equities credited for every 1 underlying equity debited = 2 resulting equities.
	NewToOld *RatioFormat12Choice `xml:"NewToOd,omitempty"`

	// Rate used to determine the cash consideration split across outturn settlement transactions that are the result of a transformation of the parent transaction.
	TransformationRate *PercentageRate `xml:"TrfrmatnRate,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat14Choice `xml:"ChrgsFees,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *RateFormat6Choice `xml:"FsclStmp,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *RateFormat6Choice `xml:"AplblRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat5Choice `xml:"TaxCdtRate,omitempty"`
}

func (c *CorporateActionRate28) AddAdditionalQuantityForSubscribedResultantSecurities() *RatioFormat11Choice {
	c.AdditionalQuantityForSubscribedResultantSecurities = new(RatioFormat11Choice)
	return c.AdditionalQuantityForSubscribedResultantSecurities
}

func (c *CorporateActionRate28) AddAdditionalQuantityForExistingSecurities() *RatioFormat11Choice {
	c.AdditionalQuantityForExistingSecurities = new(RatioFormat11Choice)
	return c.AdditionalQuantityForExistingSecurities
}

func (c *CorporateActionRate28) AddNewToOld() *RatioFormat12Choice {
	c.NewToOld = new(RatioFormat12Choice)
	return c.NewToOld
}

func (c *CorporateActionRate28) SetTransformationRate(value string) {
	c.TransformationRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate28) AddChargesFees() *RateAndAmountFormat14Choice {
	c.ChargesFees = new(RateAndAmountFormat14Choice)
	return c.ChargesFees
}

func (c *CorporateActionRate28) AddFiscalStamp() *RateFormat6Choice {
	c.FiscalStamp = new(RateFormat6Choice)
	return c.FiscalStamp
}

func (c *CorporateActionRate28) AddApplicableRate() *RateFormat6Choice {
	c.ApplicableRate = new(RateFormat6Choice)
	return c.ApplicableRate
}

func (c *CorporateActionRate28) AddTaxCreditRate() *TaxCreditRateFormat5Choice {
	newValue := new(TaxCreditRateFormat5Choice)
	c.TaxCreditRate = append(c.TaxCreditRate, newValue)
	return newValue
}

// Specifies rate details.
type CorporateActionRate29 struct {

	// Quantity of additional intermediate securities/new equities awarded for a given quantity of securities derived from subscription.
	AdditionalQuantityForSubscribedResultantSecurities *RatioFormat3Choice `xml:"AddtlQtyForSbcbdRsltntScties,omitempty"`

	// Quantity of additional securities for a given quantity of underlying securities where underlying securities are not exchanged or debited, for example, 1 for 1: 1 new equity credited for every 1 underlying equity = 2 resulting equities.
	AdditionalQuantityForExistingSecurities *RatioFormat3Choice `xml:"AddtlQtyForExstgScties,omitempty"`

	// Quantity of new securities for a given quantity of underlying securities, where the underlying securities will be exchanged or debited, for example, 2 for 1: 2 new equities credited for every 1 underlying equity debited = 2 resulting equities.
	NewToOld *RatioFormat15Choice `xml:"NewToOd,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat5Choice `xml:"ChrgsFees,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *PercentageRate `xml:"FsclStmp,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *PercentageRate `xml:"AplblRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat2Choice `xml:"TaxCdtRate,omitempty"`
}

func (c *CorporateActionRate29) AddAdditionalQuantityForSubscribedResultantSecurities() *RatioFormat3Choice {
	c.AdditionalQuantityForSubscribedResultantSecurities = new(RatioFormat3Choice)
	return c.AdditionalQuantityForSubscribedResultantSecurities
}

func (c *CorporateActionRate29) AddAdditionalQuantityForExistingSecurities() *RatioFormat3Choice {
	c.AdditionalQuantityForExistingSecurities = new(RatioFormat3Choice)
	return c.AdditionalQuantityForExistingSecurities
}

func (c *CorporateActionRate29) AddNewToOld() *RatioFormat15Choice {
	c.NewToOld = new(RatioFormat15Choice)
	return c.NewToOld
}

func (c *CorporateActionRate29) AddChargesFees() *RateAndAmountFormat5Choice {
	c.ChargesFees = new(RateAndAmountFormat5Choice)
	return c.ChargesFees
}

func (c *CorporateActionRate29) SetFiscalStamp(value string) {
	c.FiscalStamp = (*PercentageRate)(&value)
}

func (c *CorporateActionRate29) SetApplicableRate(value string) {
	c.ApplicableRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate29) AddTaxCreditRate() *TaxCreditRateFormat2Choice {
	newValue := new(TaxCreditRateFormat2Choice)
	c.TaxCreditRate = append(c.TaxCreditRate, newValue)
	return newValue
}

// Specifies rates of a corporate action.
type CorporateActionRate3 struct {

	// Annual rate of a financial instrument.
	Interest *RateAndAmountFormat3Choice `xml:"Intrst,omitempty"`

	// Percentage of securities the offeror/issuer will purchase or redeem under the terms of the event. This can be a number or the term "any and all".
	PercentageSought *RateFormat5Choice `xml:"PctgSght,omitempty"`

	// Index rate related to the interest rate of the forthcoming interest payment.
	RelatedIndex *RateFormat2Choice `xml:"RltdIndx,omitempty"`

	// Margin allowed over or under a given rate.
	Spread *RateFormat2Choice `xml:"Sprd,omitempty"`

	// Acceptable price increment used for submitting a bid.
	BidInterval *RateAndAmountFormat3Choice `xml:"BidIntrvl,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *RateFormat3Choice `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *RateFormat3Choice `xml:"NxtFctr,omitempty"`

	// Rate of discount for securities purchased through a reinvestment scheme as compared to the current market price of security.
	ReinvestmentDiscountRateToMarket *RateFormat2Choice `xml:"RinvstmtDscntRateToMkt,omitempty"`
}

func (c *CorporateActionRate3) AddInterest() *RateAndAmountFormat3Choice {
	c.Interest = new(RateAndAmountFormat3Choice)
	return c.Interest
}

func (c *CorporateActionRate3) AddPercentageSought() *RateFormat5Choice {
	c.PercentageSought = new(RateFormat5Choice)
	return c.PercentageSought
}

func (c *CorporateActionRate3) AddRelatedIndex() *RateFormat2Choice {
	c.RelatedIndex = new(RateFormat2Choice)
	return c.RelatedIndex
}

func (c *CorporateActionRate3) AddSpread() *RateFormat2Choice {
	c.Spread = new(RateFormat2Choice)
	return c.Spread
}

func (c *CorporateActionRate3) AddBidInterval() *RateAndAmountFormat3Choice {
	c.BidInterval = new(RateAndAmountFormat3Choice)
	return c.BidInterval
}

func (c *CorporateActionRate3) AddPreviousFactor() *RateFormat3Choice {
	c.PreviousFactor = new(RateFormat3Choice)
	return c.PreviousFactor
}

func (c *CorporateActionRate3) AddNextFactor() *RateFormat3Choice {
	c.NextFactor = new(RateFormat3Choice)
	return c.NextFactor
}

func (c *CorporateActionRate3) AddReinvestmentDiscountRateToMarket() *RateFormat2Choice {
	c.ReinvestmentDiscountRateToMarket = new(RateFormat2Choice)
	return c.ReinvestmentDiscountRateToMarket
}

// Specifies rates of a corporate action.
type CorporateActionRate35 struct {

	// Annual rate of a financial instrument.
	Interest *RateAndAmountFormat14Choice `xml:"Intrst,omitempty"`

	// Percentage of securities the offeror/issuer will purchase or redeem under the terms of the event.
	PercentageSought *RateFormat7Choice `xml:"PctgSght,omitempty"`

	// Index rate related to the interest rate of the forthcoming interest payment.
	RelatedIndex *RateFormat6Choice `xml:"RltdIndx,omitempty"`

	// Margin allowed over or under a given rate.
	Spread *RateFormat6Choice `xml:"Sprd,omitempty"`

	// Acceptable price increment used for submitting a bid.
	BidInterval *RateAndAmountFormat14Choice `xml:"BidIntrvl,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *RateFormat3Choice `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *RateFormat3Choice `xml:"NxtFctr,omitempty"`

	// Rate of discount for securities purchased through a reinvestment scheme as compared to the current market price of security.
	ReinvestmentDiscountRateToMarket *RateFormat6Choice `xml:"RinvstmtDscntRateToMkt,omitempty"`

	// For structured security issues where there is a set schedule of principal and interest payments for the life of the issue, this is the difference between the actual rate of the interest payment and the expected or scheduled rate of the interest payment .
	InterestShortfall *RateAndAmountFormat12Choice `xml:"IntrstShrtfll,omitempty"`

	// For structured security issues where there is a set schedule of principal and interest payments for the life of the issue, this is the difference between the actual rate of the capital or principal repayment and the scheduled capital repayment.
	RealisedLoss *RateAndAmountFormat12Choice `xml:"RealsdLoss,omitempty"`

	// Dividend or interest rate declared by the issuer.
	DeclaredRate *RateAndAmountFormat12Choice `xml:"DclrdRate,omitempty"`
}

func (c *CorporateActionRate35) AddInterest() *RateAndAmountFormat14Choice {
	c.Interest = new(RateAndAmountFormat14Choice)
	return c.Interest
}

func (c *CorporateActionRate35) AddPercentageSought() *RateFormat7Choice {
	c.PercentageSought = new(RateFormat7Choice)
	return c.PercentageSought
}

func (c *CorporateActionRate35) AddRelatedIndex() *RateFormat6Choice {
	c.RelatedIndex = new(RateFormat6Choice)
	return c.RelatedIndex
}

func (c *CorporateActionRate35) AddSpread() *RateFormat6Choice {
	c.Spread = new(RateFormat6Choice)
	return c.Spread
}

func (c *CorporateActionRate35) AddBidInterval() *RateAndAmountFormat14Choice {
	c.BidInterval = new(RateAndAmountFormat14Choice)
	return c.BidInterval
}

func (c *CorporateActionRate35) AddPreviousFactor() *RateFormat3Choice {
	c.PreviousFactor = new(RateFormat3Choice)
	return c.PreviousFactor
}

func (c *CorporateActionRate35) AddNextFactor() *RateFormat3Choice {
	c.NextFactor = new(RateFormat3Choice)
	return c.NextFactor
}

func (c *CorporateActionRate35) AddReinvestmentDiscountRateToMarket() *RateFormat6Choice {
	c.ReinvestmentDiscountRateToMarket = new(RateFormat6Choice)
	return c.ReinvestmentDiscountRateToMarket
}

func (c *CorporateActionRate35) AddInterestShortfall() *RateAndAmountFormat12Choice {
	c.InterestShortfall = new(RateAndAmountFormat12Choice)
	return c.InterestShortfall
}

func (c *CorporateActionRate35) AddRealisedLoss() *RateAndAmountFormat12Choice {
	c.RealisedLoss = new(RateAndAmountFormat12Choice)
	return c.RealisedLoss
}

func (c *CorporateActionRate35) AddDeclaredRate() *RateAndAmountFormat12Choice {
	c.DeclaredRate = new(RateAndAmountFormat12Choice)
	return c.DeclaredRate
}

// Specifies rates related to a corporate action option.
type CorporateActionRate36 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat14Choice `xml:"AddtlTax,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat9Choice `xml:"GrssDvddRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat11Choice `xml:"NetDvddRate,omitempty"`

	// Public index rate applied to the amount paid to adjust it to inflation.
	IndexFactor *RateAndAmountFormat14Choice `xml:"IndxFctr,omitempty"`

	// The actual interest rate used for the payment of the interest for the specified interest period.
	// Usage guideline: It is used to provide the applicable rate for the current payment, after all calculations have been performed, that is, application of period and method of interest computation.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat5Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Maximum percentage of shares available through the over subscription privilege, usually a percentage of the basic subscription shares, for example, an account owner subscribing to 100 shares may over subscribe to a maximum of 50 additional shares when the over subscription maximum is 50 percent.
	MaximumAllowedOversubscriptionRate *RateFormat6Choice `xml:"MaxAllwdOvrsbcptRate,omitempty"`

	// Proportionate allocation used for the offer.
	ProrationRate *RateFormat6Choice `xml:"PrratnRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxRate *RateFormat6Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction to which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTax *RateAndAmountFormat14Choice `xml:"WhldgOfFrgnTax,omitempty"`

	// Percentage of the gross dividend rate on which tax must be paid .
	TaxRelatedRate []*RateTypeAndAmountAndStatus6 `xml:"TaxRltdRate,omitempty"`

	// Amount included in the dividend/NAV that is identified as gains directly or indirectly derived from interest payments, for example, in the context of the EU Savings directive.
	TaxableIncomePerDividendShare []*RateTypeAndAmountAndStatus11 `xml:"TaxblIncmPerDvddShr,omitempty"`

	// Exchange rate (provided by the issuer) between the dividend or interest rate in the paid currency and the declared dividend or interest rate.
	IssuerDeclaredExchangeRate *ForeignExchangeTerms19 `xml:"IssrDclrdXchgRate,omitempty"`
}

func (c *CorporateActionRate36) AddAdditionalTax() *RateAndAmountFormat14Choice {
	c.AdditionalTax = new(RateAndAmountFormat14Choice)
	return c.AdditionalTax
}

func (c *CorporateActionRate36) AddGrossDividendRate() *GrossDividendRateFormat9Choice {
	newValue := new(GrossDividendRateFormat9Choice)
	c.GrossDividendRate = append(c.GrossDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate36) AddNetDividendRate() *NetDividendRateFormat11Choice {
	newValue := new(NetDividendRateFormat11Choice)
	c.NetDividendRate = append(c.NetDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate36) AddIndexFactor() *RateAndAmountFormat14Choice {
	c.IndexFactor = new(RateAndAmountFormat14Choice)
	return c.IndexFactor
}

func (c *CorporateActionRate36) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat5Choice {
	newValue := new(InterestRateUsedForPaymentFormat5Choice)
	c.InterestRateUsedForPayment = append(c.InterestRateUsedForPayment, newValue)
	return newValue
}

func (c *CorporateActionRate36) AddMaximumAllowedOversubscriptionRate() *RateFormat6Choice {
	c.MaximumAllowedOversubscriptionRate = new(RateFormat6Choice)
	return c.MaximumAllowedOversubscriptionRate
}

func (c *CorporateActionRate36) AddProrationRate() *RateFormat6Choice {
	c.ProrationRate = new(RateFormat6Choice)
	return c.ProrationRate
}

func (c *CorporateActionRate36) AddWithholdingTaxRate() *RateFormat6Choice {
	c.WithholdingTaxRate = new(RateFormat6Choice)
	return c.WithholdingTaxRate
}

func (c *CorporateActionRate36) AddWithholdingOfForeignTax() *RateAndAmountFormat14Choice {
	c.WithholdingOfForeignTax = new(RateAndAmountFormat14Choice)
	return c.WithholdingOfForeignTax
}

func (c *CorporateActionRate36) AddTaxRelatedRate() *RateTypeAndAmountAndStatus6 {
	newValue := new(RateTypeAndAmountAndStatus6)
	c.TaxRelatedRate = append(c.TaxRelatedRate, newValue)
	return newValue
}

func (c *CorporateActionRate36) AddTaxableIncomePerDividendShare() *RateTypeAndAmountAndStatus11 {
	newValue := new(RateTypeAndAmountAndStatus11)
	c.TaxableIncomePerDividendShare = append(c.TaxableIncomePerDividendShare, newValue)
	return newValue
}

func (c *CorporateActionRate36) AddIssuerDeclaredExchangeRate() *ForeignExchangeTerms19 {
	c.IssuerDeclaredExchangeRate = new(ForeignExchangeTerms19)
	return c.IssuerDeclaredExchangeRate
}

// Specifies rates related to a corporate action option.
type CorporateActionRate37 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat14Choice `xml:"AddtlTax,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat9Choice `xml:"GrssDvddRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat11Choice `xml:"NetDvddRate,omitempty"`

	// Public index rate applied to the amount paid to adjust it to inflation.
	IndexFactor *RateAndAmountFormat14Choice `xml:"IndxFctr,omitempty"`

	// The actual interest rate used for the payment of the interest for the specified interest period.
	// Usage guideline: It is used to provide the applicable rate for the current payment, after all calculations have been performed, that is, application of period and method of interest computation.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat5Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Maximum percentage of shares available through the over subscription privilege, usually a percentage of the basic subscription shares, for example, an account owner subscribing to 100 shares may over subscribe to a maximum of 50 additional shares when the over subscription maximum is 50 percent.
	MaximumAllowedOversubscriptionRate *RateFormat6Choice `xml:"MaxAllwdOvrsbcptRate,omitempty"`

	// Proportionate allocation used for the offer.
	ProrationRate *RateFormat6Choice `xml:"PrratnRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxRate *RateFormat6Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction to which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTax *RateAndAmountFormat14Choice `xml:"WhldgOfFrgnTax,omitempty"`

	// Percentage of the gross dividend rate on which tax must be paid .
	TaxRelatedRate []*RateTypeAndAmountAndStatus6 `xml:"TaxRltdRate,omitempty"`

	// Amount included in the dividend/NAV that is identified as gains directly or indirectly derived from interest payments, for example, in the context of the EU Savings directive.
	TaxableIncomePerDividendShare []*RateTypeAndAmountAndStatus11 `xml:"TaxblIncmPerDvddShr,omitempty"`
}

func (c *CorporateActionRate37) AddAdditionalTax() *RateAndAmountFormat14Choice {
	c.AdditionalTax = new(RateAndAmountFormat14Choice)
	return c.AdditionalTax
}

func (c *CorporateActionRate37) AddGrossDividendRate() *GrossDividendRateFormat9Choice {
	newValue := new(GrossDividendRateFormat9Choice)
	c.GrossDividendRate = append(c.GrossDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate37) AddNetDividendRate() *NetDividendRateFormat11Choice {
	newValue := new(NetDividendRateFormat11Choice)
	c.NetDividendRate = append(c.NetDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate37) AddIndexFactor() *RateAndAmountFormat14Choice {
	c.IndexFactor = new(RateAndAmountFormat14Choice)
	return c.IndexFactor
}

func (c *CorporateActionRate37) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat5Choice {
	newValue := new(InterestRateUsedForPaymentFormat5Choice)
	c.InterestRateUsedForPayment = append(c.InterestRateUsedForPayment, newValue)
	return newValue
}

func (c *CorporateActionRate37) AddMaximumAllowedOversubscriptionRate() *RateFormat6Choice {
	c.MaximumAllowedOversubscriptionRate = new(RateFormat6Choice)
	return c.MaximumAllowedOversubscriptionRate
}

func (c *CorporateActionRate37) AddProrationRate() *RateFormat6Choice {
	c.ProrationRate = new(RateFormat6Choice)
	return c.ProrationRate
}

func (c *CorporateActionRate37) AddWithholdingTaxRate() *RateFormat6Choice {
	c.WithholdingTaxRate = new(RateFormat6Choice)
	return c.WithholdingTaxRate
}

func (c *CorporateActionRate37) AddWithholdingOfForeignTax() *RateAndAmountFormat14Choice {
	c.WithholdingOfForeignTax = new(RateAndAmountFormat14Choice)
	return c.WithholdingOfForeignTax
}

func (c *CorporateActionRate37) AddTaxRelatedRate() *RateTypeAndAmountAndStatus6 {
	newValue := new(RateTypeAndAmountAndStatus6)
	c.TaxRelatedRate = append(c.TaxRelatedRate, newValue)
	return newValue
}

func (c *CorporateActionRate37) AddTaxableIncomePerDividendShare() *RateTypeAndAmountAndStatus11 {
	newValue := new(RateTypeAndAmountAndStatus11)
	c.TaxableIncomePerDividendShare = append(c.TaxableIncomePerDividendShare, newValue)
	return newValue
}

// Specifies rates.
type CorporateActionRate38 struct {

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat10Choice `xml:"GrssDvddRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat12Choice `xml:"NetDvddRate,omitempty"`

	// Public index rate applied to the amount paid to adjust it to inflation.
	IndexFactor *RateAndAmountFormat5Choice `xml:"IndxFctr,omitempty"`

	// Actual interest rate used for the payment of the interest for the specified interest period.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat2Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// A maximum percentage of shares available through the over subscription privilege, usually a percentage of the basic subscription shares, for example, an account owner subscribing to 100 shares may over subscribe to a maximum of 50 additional shares when the over subscription maximum is 50 percent.
	MaximumAllowedOversubscriptionRate *PercentageRate `xml:"MaxAllwdOvrsbcptRate,omitempty"`

	// Proportionate allocation used for the offer.
	ProrationRate *PercentageRate `xml:"PrratnRate,omitempty"`

	// Percentage of the gross dividend rate on which tax must be paid .
	TaxRelatedRate []*RateTypeAndAmountAndStatus6 `xml:"TaxRltdRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxRate *PercentageRate `xml:"WhldgTaxRate,omitempty"`

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat5Choice `xml:"AddtlTax,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction to which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTax *RateAndAmountFormat5Choice `xml:"WhldgOfFrgnTax,omitempty"`

	// Amount included in the dividend/NAV that is identified as gains directly or indirectly derived from interest payments, for example, in the context of the EU Savings directive.
	TaxableIncomePerDividendShare []*RateTypeAndAmountAndStatus11 `xml:"TaxblIncmPerDvddShr,omitempty"`
}

func (c *CorporateActionRate38) AddGrossDividendRate() *GrossDividendRateFormat10Choice {
	newValue := new(GrossDividendRateFormat10Choice)
	c.GrossDividendRate = append(c.GrossDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate38) AddNetDividendRate() *NetDividendRateFormat12Choice {
	newValue := new(NetDividendRateFormat12Choice)
	c.NetDividendRate = append(c.NetDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate38) AddIndexFactor() *RateAndAmountFormat5Choice {
	c.IndexFactor = new(RateAndAmountFormat5Choice)
	return c.IndexFactor
}

func (c *CorporateActionRate38) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat2Choice {
	newValue := new(InterestRateUsedForPaymentFormat2Choice)
	c.InterestRateUsedForPayment = append(c.InterestRateUsedForPayment, newValue)
	return newValue
}

func (c *CorporateActionRate38) SetMaximumAllowedOversubscriptionRate(value string) {
	c.MaximumAllowedOversubscriptionRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate38) SetProrationRate(value string) {
	c.ProrationRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate38) AddTaxRelatedRate() *RateTypeAndAmountAndStatus6 {
	newValue := new(RateTypeAndAmountAndStatus6)
	c.TaxRelatedRate = append(c.TaxRelatedRate, newValue)
	return newValue
}

func (c *CorporateActionRate38) SetWithholdingTaxRate(value string) {
	c.WithholdingTaxRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate38) AddAdditionalTax() *RateAndAmountFormat5Choice {
	c.AdditionalTax = new(RateAndAmountFormat5Choice)
	return c.AdditionalTax
}

func (c *CorporateActionRate38) AddWithholdingOfForeignTax() *RateAndAmountFormat5Choice {
	c.WithholdingOfForeignTax = new(RateAndAmountFormat5Choice)
	return c.WithholdingOfForeignTax
}

func (c *CorporateActionRate38) AddTaxableIncomePerDividendShare() *RateTypeAndAmountAndStatus11 {
	newValue := new(RateTypeAndAmountAndStatus11)
	c.TaxableIncomePerDividendShare = append(c.TaxableIncomePerDividendShare, newValue)
	return newValue
}

// Specifies rates.
type CorporateActionRate4 struct {

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat5Choice `xml:"ChrgsFees,omitempty"`

	// Dividend is final.
	FinalDividendRate *ActiveCurrencyAnd13DecimalAmount `xml:"FnlDvddRate,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *PercentageRate `xml:"FsclStmp,omitempty"`

	// Rate resulting from a fully franked dividend paid by a company; rate includes tax credit for companies that have made sufficient tax payments during fiscal period.
	FullyFrankedRate *RateAndAmountFormat5Choice `xml:"FullyFrnkdRate,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat2Choice `xml:"GrssDvddRate,omitempty"`

	// Rate of the cash premium made available if the securities holder consents or participates to an event, for example, consent fees.
	CashIncentiveRate *PercentageRate `xml:"CshIncntivRate,omitempty"`

	// Public index rate applied to the amount paid to adjust it to inflation.
	IndexFactor *RateAndAmountFormat5Choice `xml:"IndxFctr,omitempty"`

	// Actual interest rate used for the payment of the interest for the specified interest period.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat2Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat2Choice `xml:"NetDvddRate,omitempty"`

	// Rate per share to which a non-resident is entitled.
	NonResidentRate *RateAndAmountFormat5Choice `xml:"NonResdtRate,omitempty"`

	// A maximum percentage of shares available through the over subscription privilege, usually a percentage of the basic subscription shares, for example, an account owner subscribing to 100 shares may over subscribe to a maximum of 50 additional shares when the over subscription maximum is 50 percent.
	MaximumAllowedOversubscriptionRate *PercentageRate `xml:"MaxAllwdOvrsbcptRate,omitempty"`

	// Dividend is provisional.
	ProvisionalDividendRate *ActiveCurrencyAnd13DecimalAmount `xml:"PrvsnlDvddRate,omitempty"`

	// Proportionate allocation used for the offer.
	ProrationRate *PercentageRate `xml:"PrratnRate,omitempty"`

	// Cash rate made available in an offer in order to encourage participation in the offer.
	SolicitationFeeRate *PercentageRate `xml:"SlctnFeeRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat2Choice `xml:"TaxCdtRate,omitempty"`

	// Percentage of the gross dividend rate on which tax must be paid .
	TaxRelatedRate []*RateTypeAndAmountAndStatus6 `xml:"TaxRltdRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxRate *PercentageRate `xml:"WhldgTaxRate,omitempty"`

	// Taxation applied on an amount clearly identified as an income.
	TaxOnIncome *PercentageRate `xml:"TaxOnIncm,omitempty"`

	// Taxation applied on an amount clearly identified as capital profits, capital gains.
	TaxOnProfits *PercentageRate `xml:"TaxOnPrfts,omitempty"`

	// Percentage of cash that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimRate *PercentageRate `xml:"TaxRclmRate,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction in which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTax *RateAndAmountFormat5Choice `xml:"WhldgOfFrgnTax,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction in which the account owner is located, for which relief at source and/or reclaim may be possible.
	WithholdingOfLocalTax *RateAndAmountFormat5Choice `xml:"WhldgOfLclTax,omitempty"`

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat5Choice `xml:"AddtlTax,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *PercentageRate `xml:"AplblRate,omitempty"`
}

func (c *CorporateActionRate4) AddChargesFees() *RateAndAmountFormat5Choice {
	c.ChargesFees = new(RateAndAmountFormat5Choice)
	return c.ChargesFees
}

func (c *CorporateActionRate4) SetFinalDividendRate(value, currency string) {
	c.FinalDividendRate = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *CorporateActionRate4) SetFiscalStamp(value string) {
	c.FiscalStamp = (*PercentageRate)(&value)
}

func (c *CorporateActionRate4) AddFullyFrankedRate() *RateAndAmountFormat5Choice {
	c.FullyFrankedRate = new(RateAndAmountFormat5Choice)
	return c.FullyFrankedRate
}

func (c *CorporateActionRate4) AddGrossDividendRate() *GrossDividendRateFormat2Choice {
	newValue := new(GrossDividendRateFormat2Choice)
	c.GrossDividendRate = append(c.GrossDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate4) SetCashIncentiveRate(value string) {
	c.CashIncentiveRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate4) AddIndexFactor() *RateAndAmountFormat5Choice {
	c.IndexFactor = new(RateAndAmountFormat5Choice)
	return c.IndexFactor
}

func (c *CorporateActionRate4) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat2Choice {
	newValue := new(InterestRateUsedForPaymentFormat2Choice)
	c.InterestRateUsedForPayment = append(c.InterestRateUsedForPayment, newValue)
	return newValue
}

func (c *CorporateActionRate4) AddNetDividendRate() *NetDividendRateFormat2Choice {
	newValue := new(NetDividendRateFormat2Choice)
	c.NetDividendRate = append(c.NetDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate4) AddNonResidentRate() *RateAndAmountFormat5Choice {
	c.NonResidentRate = new(RateAndAmountFormat5Choice)
	return c.NonResidentRate
}

func (c *CorporateActionRate4) SetMaximumAllowedOversubscriptionRate(value string) {
	c.MaximumAllowedOversubscriptionRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate4) SetProvisionalDividendRate(value, currency string) {
	c.ProvisionalDividendRate = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *CorporateActionRate4) SetProrationRate(value string) {
	c.ProrationRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate4) SetSolicitationFeeRate(value string) {
	c.SolicitationFeeRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate4) AddTaxCreditRate() *TaxCreditRateFormat2Choice {
	newValue := new(TaxCreditRateFormat2Choice)
	c.TaxCreditRate = append(c.TaxCreditRate, newValue)
	return newValue
}

func (c *CorporateActionRate4) AddTaxRelatedRate() *RateTypeAndAmountAndStatus6 {
	newValue := new(RateTypeAndAmountAndStatus6)
	c.TaxRelatedRate = append(c.TaxRelatedRate, newValue)
	return newValue
}

func (c *CorporateActionRate4) SetWithholdingTaxRate(value string) {
	c.WithholdingTaxRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate4) SetTaxOnIncome(value string) {
	c.TaxOnIncome = (*PercentageRate)(&value)
}

func (c *CorporateActionRate4) SetTaxOnProfits(value string) {
	c.TaxOnProfits = (*PercentageRate)(&value)
}

func (c *CorporateActionRate4) SetTaxReclaimRate(value string) {
	c.TaxReclaimRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate4) AddWithholdingOfForeignTax() *RateAndAmountFormat5Choice {
	c.WithholdingOfForeignTax = new(RateAndAmountFormat5Choice)
	return c.WithholdingOfForeignTax
}

func (c *CorporateActionRate4) AddWithholdingOfLocalTax() *RateAndAmountFormat5Choice {
	c.WithholdingOfLocalTax = new(RateAndAmountFormat5Choice)
	return c.WithholdingOfLocalTax
}

func (c *CorporateActionRate4) AddAdditionalTax() *RateAndAmountFormat5Choice {
	c.AdditionalTax = new(RateAndAmountFormat5Choice)
	return c.AdditionalTax
}

func (c *CorporateActionRate4) SetApplicableRate(value string) {
	c.ApplicableRate = (*PercentageRate)(&value)
}

// Specifies rates of a corporate action.
type CorporateActionRate43 struct {

	// Annual rate of a financial instrument.
	Interest *RateAndAmountFormat14Choice `xml:"Intrst,omitempty"`

	// Percentage of securities the offeror/issuer will purchase or redeem under the terms of the event.
	PercentageSought *RateFormat7Choice `xml:"PctgSght,omitempty"`

	// Index rate related to the interest rate of the forthcoming interest payment.
	RelatedIndex *RateFormat3Choice `xml:"RltdIndx,omitempty"`

	// Margin allowed over or under a given rate.
	Spread *RateFormat3Choice `xml:"Sprd,omitempty"`

	// Acceptable price increment used for submitting a bid.
	BidInterval *RateAndAmountFormat19Choice `xml:"BidIntrvl,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *RateFormat12Choice `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *RateFormat12Choice `xml:"NxtFctr,omitempty"`

	// Rate of discount for securities purchased through a reinvestment scheme as compared to the current market price of security.
	ReinvestmentDiscountRateToMarket *RateFormat3Choice `xml:"RinvstmtDscntRateToMkt,omitempty"`

	// For structured security issues where there is a set schedule of principal and interest payments for the life of the issue, this is the difference between the actual rate of the interest payment and the expected or scheduled rate of the interest payment .
	InterestShortfall *RateAndAmountFormat5Choice `xml:"IntrstShrtfll,omitempty"`

	// For structured security issues where there is a set schedule of principal and interest payments for the life of the issue, this is the difference between the actual rate of the capital or principal repayment and the scheduled capital repayment.
	RealisedLoss *RateAndAmountFormat5Choice `xml:"RealsdLoss,omitempty"`

	// Dividend or interest rate declared by the issuer.
	DeclaredRate *RateAndAmountFormat5Choice `xml:"DclrdRate,omitempty"`
}

func (c *CorporateActionRate43) AddInterest() *RateAndAmountFormat14Choice {
	c.Interest = new(RateAndAmountFormat14Choice)
	return c.Interest
}

func (c *CorporateActionRate43) AddPercentageSought() *RateFormat7Choice {
	c.PercentageSought = new(RateFormat7Choice)
	return c.PercentageSought
}

func (c *CorporateActionRate43) AddRelatedIndex() *RateFormat3Choice {
	c.RelatedIndex = new(RateFormat3Choice)
	return c.RelatedIndex
}

func (c *CorporateActionRate43) AddSpread() *RateFormat3Choice {
	c.Spread = new(RateFormat3Choice)
	return c.Spread
}

func (c *CorporateActionRate43) AddBidInterval() *RateAndAmountFormat19Choice {
	c.BidInterval = new(RateAndAmountFormat19Choice)
	return c.BidInterval
}

func (c *CorporateActionRate43) AddPreviousFactor() *RateFormat12Choice {
	c.PreviousFactor = new(RateFormat12Choice)
	return c.PreviousFactor
}

func (c *CorporateActionRate43) AddNextFactor() *RateFormat12Choice {
	c.NextFactor = new(RateFormat12Choice)
	return c.NextFactor
}

func (c *CorporateActionRate43) AddReinvestmentDiscountRateToMarket() *RateFormat3Choice {
	c.ReinvestmentDiscountRateToMarket = new(RateFormat3Choice)
	return c.ReinvestmentDiscountRateToMarket
}

func (c *CorporateActionRate43) AddInterestShortfall() *RateAndAmountFormat5Choice {
	c.InterestShortfall = new(RateAndAmountFormat5Choice)
	return c.InterestShortfall
}

func (c *CorporateActionRate43) AddRealisedLoss() *RateAndAmountFormat5Choice {
	c.RealisedLoss = new(RateAndAmountFormat5Choice)
	return c.RealisedLoss
}

func (c *CorporateActionRate43) AddDeclaredRate() *RateAndAmountFormat5Choice {
	c.DeclaredRate = new(RateAndAmountFormat5Choice)
	return c.DeclaredRate
}

// Specifies rates related to a corporate action option.
type CorporateActionRate44 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat14Choice `xml:"AddtlTax,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat9Choice `xml:"GrssDvddRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat11Choice `xml:"NetDvddRate,omitempty"`

	// Public index rate applied to the amount paid to adjust it to inflation.
	IndexFactor *RateAndAmountFormat14Choice `xml:"IndxFctr,omitempty"`

	// The actual interest rate used for the payment of the interest for the specified interest period.
	// Usage guideline: It is used to provide the applicable rate for the current payment, after all calculations have been performed, that is, application of period and method of interest computation.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat5Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Maximum percentage of shares available through the over subscription privilege, usually a percentage of the basic subscription shares, for example, an account owner subscribing to 100 shares may over subscribe to a maximum of 50 additional shares when the over subscription maximum is 50 percent.
	MaximumAllowedOversubscriptionRate *RateFormat3Choice `xml:"MaxAllwdOvrsbcptRate,omitempty"`

	// Proportionate allocation used for the offer.
	ProrationRate *RateFormat3Choice `xml:"PrratnRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxRate []*RateFormat10Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction to which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTax []*RateAndAmountFormat20Choice `xml:"WhldgOfFrgnTax,omitempty"`

	// Percentage of the gross dividend rate on which tax must be paid .
	TaxRelatedRate []*RateTypeAndAmountAndStatus6 `xml:"TaxRltdRate,omitempty"`

	// Amount included in the dividend/NAV that is identified as gains directly or indirectly derived from interest payments, for example, in the context of the EU Savings directive.
	TaxableIncomePerDividendShare []*RateTypeAndAmountAndStatus11 `xml:"TaxblIncmPerDvddShr,omitempty"`

	// Exchange rate (provided by the issuer) between the dividend or interest rate in the paid currency and the declared dividend or interest rate.
	IssuerDeclaredExchangeRate *ForeignExchangeTerms19 `xml:"IssrDclrdXchgRate,omitempty"`

	// Overall tax withheld at source by fund managers prior to considering the tax obligation of each unit holder.
	TaxOnIncome *RateAndAmountFormat14Choice `xml:"TaxOnIncm,omitempty"`
}

func (c *CorporateActionRate44) AddAdditionalTax() *RateAndAmountFormat14Choice {
	c.AdditionalTax = new(RateAndAmountFormat14Choice)
	return c.AdditionalTax
}

func (c *CorporateActionRate44) AddGrossDividendRate() *GrossDividendRateFormat9Choice {
	newValue := new(GrossDividendRateFormat9Choice)
	c.GrossDividendRate = append(c.GrossDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate44) AddNetDividendRate() *NetDividendRateFormat11Choice {
	newValue := new(NetDividendRateFormat11Choice)
	c.NetDividendRate = append(c.NetDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate44) AddIndexFactor() *RateAndAmountFormat14Choice {
	c.IndexFactor = new(RateAndAmountFormat14Choice)
	return c.IndexFactor
}

func (c *CorporateActionRate44) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat5Choice {
	newValue := new(InterestRateUsedForPaymentFormat5Choice)
	c.InterestRateUsedForPayment = append(c.InterestRateUsedForPayment, newValue)
	return newValue
}

func (c *CorporateActionRate44) AddMaximumAllowedOversubscriptionRate() *RateFormat3Choice {
	c.MaximumAllowedOversubscriptionRate = new(RateFormat3Choice)
	return c.MaximumAllowedOversubscriptionRate
}

func (c *CorporateActionRate44) AddProrationRate() *RateFormat3Choice {
	c.ProrationRate = new(RateFormat3Choice)
	return c.ProrationRate
}

func (c *CorporateActionRate44) AddWithholdingTaxRate() *RateFormat10Choice {
	newValue := new(RateFormat10Choice)
	c.WithholdingTaxRate = append(c.WithholdingTaxRate, newValue)
	return newValue
}

func (c *CorporateActionRate44) AddWithholdingOfForeignTax() *RateAndAmountFormat20Choice {
	newValue := new(RateAndAmountFormat20Choice)
	c.WithholdingOfForeignTax = append(c.WithholdingOfForeignTax, newValue)
	return newValue
}

func (c *CorporateActionRate44) AddTaxRelatedRate() *RateTypeAndAmountAndStatus6 {
	newValue := new(RateTypeAndAmountAndStatus6)
	c.TaxRelatedRate = append(c.TaxRelatedRate, newValue)
	return newValue
}

func (c *CorporateActionRate44) AddTaxableIncomePerDividendShare() *RateTypeAndAmountAndStatus11 {
	newValue := new(RateTypeAndAmountAndStatus11)
	c.TaxableIncomePerDividendShare = append(c.TaxableIncomePerDividendShare, newValue)
	return newValue
}

func (c *CorporateActionRate44) AddIssuerDeclaredExchangeRate() *ForeignExchangeTerms19 {
	c.IssuerDeclaredExchangeRate = new(ForeignExchangeTerms19)
	return c.IssuerDeclaredExchangeRate
}

func (c *CorporateActionRate44) AddTaxOnIncome() *RateAndAmountFormat14Choice {
	c.TaxOnIncome = new(RateAndAmountFormat14Choice)
	return c.TaxOnIncome
}

// Specifies rates related to a corporate action option.
type CorporateActionRate45 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat14Choice `xml:"AddtlTax,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat9Choice `xml:"GrssDvddRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat11Choice `xml:"NetDvddRate,omitempty"`

	// Public index rate applied to the amount paid to adjust it to inflation.
	IndexFactor *RateAndAmountFormat14Choice `xml:"IndxFctr,omitempty"`

	// The actual interest rate used for the payment of the interest for the specified interest period.
	// Usage guideline: It is used to provide the applicable rate for the current payment, after all calculations have been performed, that is, application of period and method of interest computation.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat5Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Maximum percentage of shares available through the over subscription privilege, usually a percentage of the basic subscription shares, for example, an account owner subscribing to 100 shares may over subscribe to a maximum of 50 additional shares when the over subscription maximum is 50 percent.
	MaximumAllowedOversubscriptionRate *RateFormat3Choice `xml:"MaxAllwdOvrsbcptRate,omitempty"`

	// Proportionate allocation used for the offer.
	ProrationRate *RateFormat3Choice `xml:"PrratnRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxRate []*RateFormat10Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction to which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTax []*RateAndAmountFormat20Choice `xml:"WhldgOfFrgnTax,omitempty"`

	// Percentage of the gross dividend rate on which tax must be paid .
	TaxRelatedRate []*RateTypeAndAmountAndStatus6 `xml:"TaxRltdRate,omitempty"`

	// Amount included in the dividend/NAV that is identified as gains directly or indirectly derived from interest payments, for example, in the context of the EU Savings directive.
	TaxableIncomePerDividendShare []*RateTypeAndAmountAndStatus11 `xml:"TaxblIncmPerDvddShr,omitempty"`

	// Overall tax withheld at source by fund managers prior to considering the tax obligation of each unit holder.
	TaxOnIncome *RateAndAmountFormat14Choice `xml:"TaxOnIncm,omitempty"`
}

func (c *CorporateActionRate45) AddAdditionalTax() *RateAndAmountFormat14Choice {
	c.AdditionalTax = new(RateAndAmountFormat14Choice)
	return c.AdditionalTax
}

func (c *CorporateActionRate45) AddGrossDividendRate() *GrossDividendRateFormat9Choice {
	newValue := new(GrossDividendRateFormat9Choice)
	c.GrossDividendRate = append(c.GrossDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate45) AddNetDividendRate() *NetDividendRateFormat11Choice {
	newValue := new(NetDividendRateFormat11Choice)
	c.NetDividendRate = append(c.NetDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate45) AddIndexFactor() *RateAndAmountFormat14Choice {
	c.IndexFactor = new(RateAndAmountFormat14Choice)
	return c.IndexFactor
}

func (c *CorporateActionRate45) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat5Choice {
	newValue := new(InterestRateUsedForPaymentFormat5Choice)
	c.InterestRateUsedForPayment = append(c.InterestRateUsedForPayment, newValue)
	return newValue
}

func (c *CorporateActionRate45) AddMaximumAllowedOversubscriptionRate() *RateFormat3Choice {
	c.MaximumAllowedOversubscriptionRate = new(RateFormat3Choice)
	return c.MaximumAllowedOversubscriptionRate
}

func (c *CorporateActionRate45) AddProrationRate() *RateFormat3Choice {
	c.ProrationRate = new(RateFormat3Choice)
	return c.ProrationRate
}

func (c *CorporateActionRate45) AddWithholdingTaxRate() *RateFormat10Choice {
	newValue := new(RateFormat10Choice)
	c.WithholdingTaxRate = append(c.WithholdingTaxRate, newValue)
	return newValue
}

func (c *CorporateActionRate45) AddWithholdingOfForeignTax() *RateAndAmountFormat20Choice {
	newValue := new(RateAndAmountFormat20Choice)
	c.WithholdingOfForeignTax = append(c.WithholdingOfForeignTax, newValue)
	return newValue
}

func (c *CorporateActionRate45) AddTaxRelatedRate() *RateTypeAndAmountAndStatus6 {
	newValue := new(RateTypeAndAmountAndStatus6)
	c.TaxRelatedRate = append(c.TaxRelatedRate, newValue)
	return newValue
}

func (c *CorporateActionRate45) AddTaxableIncomePerDividendShare() *RateTypeAndAmountAndStatus11 {
	newValue := new(RateTypeAndAmountAndStatus11)
	c.TaxableIncomePerDividendShare = append(c.TaxableIncomePerDividendShare, newValue)
	return newValue
}

func (c *CorporateActionRate45) AddTaxOnIncome() *RateAndAmountFormat14Choice {
	c.TaxOnIncome = new(RateAndAmountFormat14Choice)
	return c.TaxOnIncome
}

// Specifies rates.
type CorporateActionRate46 struct {

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat10Choice `xml:"GrssDvddRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat12Choice `xml:"NetDvddRate,omitempty"`

	// Public index rate applied to the amount paid to adjust it to inflation.
	IndexFactor *RateAndAmountFormat5Choice `xml:"IndxFctr,omitempty"`

	// Actual interest rate used for the payment of the interest for the specified interest period.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat2Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// A maximum percentage of shares available through the over subscription privilege, usually a percentage of the basic subscription shares, for example, an account owner subscribing to 100 shares may over subscribe to a maximum of 50 additional shares when the over subscription maximum is 50 percent.
	MaximumAllowedOversubscriptionRate *PercentageRate `xml:"MaxAllwdOvrsbcptRate,omitempty"`

	// Proportionate allocation used for the offer.
	ProrationRate *PercentageRate `xml:"PrratnRate,omitempty"`

	// Percentage of the gross dividend rate on which tax must be paid .
	TaxRelatedRate []*RateTypeAndAmountAndStatus6 `xml:"TaxRltdRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxRate []*RateFormat11Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat5Choice `xml:"AddtlTax,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction to which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTax []*RateAndAmountFormat21Choice `xml:"WhldgOfFrgnTax,omitempty"`

	// Amount included in the dividend/NAV that is identified as gains directly or indirectly derived from interest payments, for example, in the context of the EU Savings directive.
	TaxableIncomePerDividendShare []*RateTypeAndAmountAndStatus11 `xml:"TaxblIncmPerDvddShr,omitempty"`
}

func (c *CorporateActionRate46) AddGrossDividendRate() *GrossDividendRateFormat10Choice {
	newValue := new(GrossDividendRateFormat10Choice)
	c.GrossDividendRate = append(c.GrossDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate46) AddNetDividendRate() *NetDividendRateFormat12Choice {
	newValue := new(NetDividendRateFormat12Choice)
	c.NetDividendRate = append(c.NetDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate46) AddIndexFactor() *RateAndAmountFormat5Choice {
	c.IndexFactor = new(RateAndAmountFormat5Choice)
	return c.IndexFactor
}

func (c *CorporateActionRate46) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat2Choice {
	newValue := new(InterestRateUsedForPaymentFormat2Choice)
	c.InterestRateUsedForPayment = append(c.InterestRateUsedForPayment, newValue)
	return newValue
}

func (c *CorporateActionRate46) SetMaximumAllowedOversubscriptionRate(value string) {
	c.MaximumAllowedOversubscriptionRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate46) SetProrationRate(value string) {
	c.ProrationRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate46) AddTaxRelatedRate() *RateTypeAndAmountAndStatus6 {
	newValue := new(RateTypeAndAmountAndStatus6)
	c.TaxRelatedRate = append(c.TaxRelatedRate, newValue)
	return newValue
}

func (c *CorporateActionRate46) AddWithholdingTaxRate() *RateFormat11Choice {
	newValue := new(RateFormat11Choice)
	c.WithholdingTaxRate = append(c.WithholdingTaxRate, newValue)
	return newValue
}

func (c *CorporateActionRate46) AddAdditionalTax() *RateAndAmountFormat5Choice {
	c.AdditionalTax = new(RateAndAmountFormat5Choice)
	return c.AdditionalTax
}

func (c *CorporateActionRate46) AddWithholdingOfForeignTax() *RateAndAmountFormat21Choice {
	newValue := new(RateAndAmountFormat21Choice)
	c.WithholdingOfForeignTax = append(c.WithholdingOfForeignTax, newValue)
	return newValue
}

func (c *CorporateActionRate46) AddTaxableIncomePerDividendShare() *RateTypeAndAmountAndStatus11 {
	newValue := new(RateTypeAndAmountAndStatus11)
	c.TaxableIncomePerDividendShare = append(c.TaxableIncomePerDividendShare, newValue)
	return newValue
}

// Specifies rates related to a corporate action option.
type CorporateActionRate47 struct {

	// Rate proposed in a remarketing of variable rate notes.
	ProposedRate *PercentageRate `xml:"PropsdRate,omitempty"`

	// Rate of allowed over-subscription.
	OversubscriptionRate *RateAndAmountFormat5Choice `xml:"OvrsbcptRate,omitempty"`

	// Requested tax rate in case of breakdown of tax rate, for example, used for adjustment of tax rate. This is the new requested applicable rate.
	RequestedTaxationRate []*RateAndAmountFormat21Choice `xml:"ReqdTaxtnRate,omitempty"`

	// Requested rate at which the income will be withheld by the jurisdiction in which the income was originally paid, for which relief at source and/or reclaim may be possible.
	RequestedWithholdingOfForeignTax []*RateAndAmountFormat21Choice `xml:"ReqdWhldgOfFrgnTax,omitempty"`

	// Requested rate at which the income will be withheld by the jurisdiction in which the account owner is located, for which relief at source and/or reclaim may be possible.
	RequestedWithholdingOfLocalTax []*RateAndAmountFormat21Choice `xml:"ReqdWhldgOfLclTax,omitempty"`
}

func (c *CorporateActionRate47) SetProposedRate(value string) {
	c.ProposedRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate47) AddOversubscriptionRate() *RateAndAmountFormat5Choice {
	c.OversubscriptionRate = new(RateAndAmountFormat5Choice)
	return c.OversubscriptionRate
}

func (c *CorporateActionRate47) AddRequestedTaxationRate() *RateAndAmountFormat21Choice {
	newValue := new(RateAndAmountFormat21Choice)
	c.RequestedTaxationRate = append(c.RequestedTaxationRate, newValue)
	return newValue
}

func (c *CorporateActionRate47) AddRequestedWithholdingOfForeignTax() *RateAndAmountFormat21Choice {
	newValue := new(RateAndAmountFormat21Choice)
	c.RequestedWithholdingOfForeignTax = append(c.RequestedWithholdingOfForeignTax, newValue)
	return newValue
}

func (c *CorporateActionRate47) AddRequestedWithholdingOfLocalTax() *RateAndAmountFormat21Choice {
	newValue := new(RateAndAmountFormat21Choice)
	c.RequestedWithholdingOfLocalTax = append(c.RequestedWithholdingOfLocalTax, newValue)
	return newValue
}

// Specifies security rate details.
type CorporateActionRate48 struct {

	// Quantity of additional intermediate securities/new equities awarded for a given quantity of securities derived from subscription.
	AdditionalQuantityForSubscribedResultantSecurities *RatioFormat11Choice `xml:"AddtlQtyForSbcbdRsltntScties,omitempty"`

	// Quantity of additional securities for a given quantity of underlying securities where underlying securities are not exchanged or debited, for example, 1 for 1: 1 new equity credited for every 1 underlying equity = 2 resulting equities.
	AdditionalQuantityForExistingSecurities *RatioFormat11Choice `xml:"AddtlQtyForExstgScties,omitempty"`

	// Quantity of new securities for a given quantity of underlying securities, where the underlying securities will be exchanged or debited, for example, 2 for 1: 2 new equities credited for every 1 underlying equity debited = 2 resulting equities.
	NewToOld *RatioFormat12Choice `xml:"NewToOd,omitempty"`

	// Rate used to determine the cash consideration split across outturn settlement transactions that are the result of a transformation of the parent transaction.
	TransformationRate *PercentageRate `xml:"TrfrmatnRate,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat14Choice `xml:"ChrgsFees,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *RateFormat3Choice `xml:"FsclStmp,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *RateFormat3Choice `xml:"AplblRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat5Choice `xml:"TaxCdtRate,omitempty"`

	// Rate of financial transaction tax.
	FinancialTransactionTaxRate *RateFormat3Choice `xml:"FinTxTaxRate,omitempty"`
}

func (c *CorporateActionRate48) AddAdditionalQuantityForSubscribedResultantSecurities() *RatioFormat11Choice {
	c.AdditionalQuantityForSubscribedResultantSecurities = new(RatioFormat11Choice)
	return c.AdditionalQuantityForSubscribedResultantSecurities
}

func (c *CorporateActionRate48) AddAdditionalQuantityForExistingSecurities() *RatioFormat11Choice {
	c.AdditionalQuantityForExistingSecurities = new(RatioFormat11Choice)
	return c.AdditionalQuantityForExistingSecurities
}

func (c *CorporateActionRate48) AddNewToOld() *RatioFormat12Choice {
	c.NewToOld = new(RatioFormat12Choice)
	return c.NewToOld
}

func (c *CorporateActionRate48) SetTransformationRate(value string) {
	c.TransformationRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate48) AddChargesFees() *RateAndAmountFormat14Choice {
	c.ChargesFees = new(RateAndAmountFormat14Choice)
	return c.ChargesFees
}

func (c *CorporateActionRate48) AddFiscalStamp() *RateFormat3Choice {
	c.FiscalStamp = new(RateFormat3Choice)
	return c.FiscalStamp
}

func (c *CorporateActionRate48) AddApplicableRate() *RateFormat3Choice {
	c.ApplicableRate = new(RateFormat3Choice)
	return c.ApplicableRate
}

func (c *CorporateActionRate48) AddTaxCreditRate() *TaxCreditRateFormat5Choice {
	newValue := new(TaxCreditRateFormat5Choice)
	c.TaxCreditRate = append(c.TaxCreditRate, newValue)
	return newValue
}

func (c *CorporateActionRate48) AddFinancialTransactionTaxRate() *RateFormat3Choice {
	c.FinancialTransactionTaxRate = new(RateFormat3Choice)
	return c.FinancialTransactionTaxRate
}

// Specifies rate details.
type CorporateActionRate49 struct {

	// Quantity of additional intermediate securities/new equities awarded for a given quantity of securities derived from subscription.
	AdditionalQuantityForSubscribedResultantSecurities *RatioFormat3Choice `xml:"AddtlQtyForSbcbdRsltntScties,omitempty"`

	// Quantity of additional securities for a given quantity of underlying securities where underlying securities are not exchanged or debited, for example, 1 for 1: 1 new equity credited for every 1 underlying equity = 2 resulting equities.
	AdditionalQuantityForExistingSecurities *RatioFormat3Choice `xml:"AddtlQtyForExstgScties,omitempty"`

	// Quantity of new securities for a given quantity of underlying securities, where the underlying securities will be exchanged or debited, for example, 2 for 1: 2 new equities credited for every 1 underlying equity debited = 2 resulting equities.
	NewToOld *RatioFormat15Choice `xml:"NewToOd,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat5Choice `xml:"ChrgsFees,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *PercentageRate `xml:"FsclStmp,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *PercentageRate `xml:"AplblRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat2Choice `xml:"TaxCdtRate,omitempty"`

	// Rate of financial transaction tax.
	FinancialTransactionTaxRate *PercentageRate `xml:"FinTxTaxRate,omitempty"`
}

func (c *CorporateActionRate49) AddAdditionalQuantityForSubscribedResultantSecurities() *RatioFormat3Choice {
	c.AdditionalQuantityForSubscribedResultantSecurities = new(RatioFormat3Choice)
	return c.AdditionalQuantityForSubscribedResultantSecurities
}

func (c *CorporateActionRate49) AddAdditionalQuantityForExistingSecurities() *RatioFormat3Choice {
	c.AdditionalQuantityForExistingSecurities = new(RatioFormat3Choice)
	return c.AdditionalQuantityForExistingSecurities
}

func (c *CorporateActionRate49) AddNewToOld() *RatioFormat15Choice {
	c.NewToOld = new(RatioFormat15Choice)
	return c.NewToOld
}

func (c *CorporateActionRate49) AddChargesFees() *RateAndAmountFormat5Choice {
	c.ChargesFees = new(RateAndAmountFormat5Choice)
	return c.ChargesFees
}

func (c *CorporateActionRate49) SetFiscalStamp(value string) {
	c.FiscalStamp = (*PercentageRate)(&value)
}

func (c *CorporateActionRate49) SetApplicableRate(value string) {
	c.ApplicableRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate49) AddTaxCreditRate() *TaxCreditRateFormat2Choice {
	newValue := new(TaxCreditRateFormat2Choice)
	c.TaxCreditRate = append(c.TaxCreditRate, newValue)
	return newValue
}

func (c *CorporateActionRate49) SetFinancialTransactionTaxRate(value string) {
	c.FinancialTransactionTaxRate = (*PercentageRate)(&value)
}

// Specifies rates related to a corporate action option.
type CorporateActionRate5 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat3Choice `xml:"AddtlTax,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat3Choice `xml:"ChrgsFees,omitempty"`

	// Dividend is final.
	FinalDividendRate *RateAndAmountFormat4Choice `xml:"FnlDvddRate,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *RateFormat2Choice `xml:"FsclStmp,omitempty"`

	// Rate resulting from a fully franked dividend paid by a company; rate includes tax credit for companies that have made sufficient tax payments during fiscal period.
	FullyFrankedRate *RateAndAmountFormat3Choice `xml:"FullyFrnkdRate,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat1Choice `xml:"GrssDvddRate,omitempty"`

	// Rate of the cash premium made available if the securities holder consents or participates to an event, for example, consent fees.
	CashIncentiveRate *RateFormat2Choice `xml:"CshIncntivRate,omitempty"`

	// Public index rate applied to the amount paid to adjust it to inflation.
	IndexFactor *RateAndAmountFormat3Choice `xml:"IndxFctr,omitempty"`

	// The actual interest rate used for the payment of the interest for the specified interest period.
	// Usage guideline: It is used to provide the applicable rate for the current payment, after all calculations have been performed, that is, application of period and method of interest computation.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat1Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat1Choice `xml:"NetDvddRate,omitempty"`

	// Rate per share to which a non-resident is entitled.
	NonResidentRate *RateAndAmountFormat3Choice `xml:"NonResdtRate,omitempty"`

	// Maximum percentage of shares available through the over subscription privilege, usually a percentage of the basic subscription shares, for example, an account owner subscribing to 100 shares may over subscribe to a maximum of 50 additional shares when the over subscription maximum is 50 percent.
	MaximumAllowedOversubscriptionRate *RateFormat2Choice `xml:"MaxAllwdOvrsbcptRate,omitempty"`

	// Dividend is provisional.
	ProvisionalDividendRate *RateAndAmountFormat4Choice `xml:"PrvsnlDvddRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat1Choice `xml:"TaxCdtRate,omitempty"`

	// Proportionate allocation used for the offer.
	ProrationRate *RateFormat2Choice `xml:"PrratnRate,omitempty"`

	// Cash rate made available in an offer in order to encourage participation in the offer.
	SolicitationFeeRate *SolicitationFeeRateFormat1Choice `xml:"SlctnFeeRate,omitempty"`

	// Cash rate made available, as an incentive, in addition to the solicitation fee, in order to encourage early participation in an offer.
	EarlySolicitationFeeRate *SolicitationFeeRateFormat1Choice `xml:"EarlySlctnFeeRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxRate *RateFormat2Choice `xml:"WhldgTaxRate,omitempty"`

	// Taxation applied on an amount clearly identified as an income.
	TaxOnIncome *RateFormat2Choice `xml:"TaxOnIncm,omitempty"`

	// Taxation applied on an amount clearly identified as capital profits, capital gains.
	TaxOnProfits *RateFormat2Choice `xml:"TaxOnPrfts,omitempty"`

	// Percentage of cash that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimRate *RateFormat2Choice `xml:"TaxRclmRate,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction in which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTax *RateAndAmountFormat3Choice `xml:"WhldgOfFrgnTax,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction in which the account owner is located, for which relief at source and/or reclaim may be possible.
	WithholdingOfLocalTax *RateAndAmountFormat3Choice `xml:"WhldgOfLclTax,omitempty"`

	// Percentage of the gross dividend rate on which tax must be paid .
	TaxRelatedRate []*RateTypeAndAmountAndStatus6 `xml:"TaxRltdRate,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *RateFormat2Choice `xml:"AplblRate,omitempty"`
}

func (c *CorporateActionRate5) AddAdditionalTax() *RateAndAmountFormat3Choice {
	c.AdditionalTax = new(RateAndAmountFormat3Choice)
	return c.AdditionalTax
}

func (c *CorporateActionRate5) AddChargesFees() *RateAndAmountFormat3Choice {
	c.ChargesFees = new(RateAndAmountFormat3Choice)
	return c.ChargesFees
}

func (c *CorporateActionRate5) AddFinalDividendRate() *RateAndAmountFormat4Choice {
	c.FinalDividendRate = new(RateAndAmountFormat4Choice)
	return c.FinalDividendRate
}

func (c *CorporateActionRate5) AddFiscalStamp() *RateFormat2Choice {
	c.FiscalStamp = new(RateFormat2Choice)
	return c.FiscalStamp
}

func (c *CorporateActionRate5) AddFullyFrankedRate() *RateAndAmountFormat3Choice {
	c.FullyFrankedRate = new(RateAndAmountFormat3Choice)
	return c.FullyFrankedRate
}

func (c *CorporateActionRate5) AddGrossDividendRate() *GrossDividendRateFormat1Choice {
	newValue := new(GrossDividendRateFormat1Choice)
	c.GrossDividendRate = append(c.GrossDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate5) AddCashIncentiveRate() *RateFormat2Choice {
	c.CashIncentiveRate = new(RateFormat2Choice)
	return c.CashIncentiveRate
}

func (c *CorporateActionRate5) AddIndexFactor() *RateAndAmountFormat3Choice {
	c.IndexFactor = new(RateAndAmountFormat3Choice)
	return c.IndexFactor
}

func (c *CorporateActionRate5) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat1Choice {
	newValue := new(InterestRateUsedForPaymentFormat1Choice)
	c.InterestRateUsedForPayment = append(c.InterestRateUsedForPayment, newValue)
	return newValue
}

func (c *CorporateActionRate5) AddNetDividendRate() *NetDividendRateFormat1Choice {
	newValue := new(NetDividendRateFormat1Choice)
	c.NetDividendRate = append(c.NetDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate5) AddNonResidentRate() *RateAndAmountFormat3Choice {
	c.NonResidentRate = new(RateAndAmountFormat3Choice)
	return c.NonResidentRate
}

func (c *CorporateActionRate5) AddMaximumAllowedOversubscriptionRate() *RateFormat2Choice {
	c.MaximumAllowedOversubscriptionRate = new(RateFormat2Choice)
	return c.MaximumAllowedOversubscriptionRate
}

func (c *CorporateActionRate5) AddProvisionalDividendRate() *RateAndAmountFormat4Choice {
	c.ProvisionalDividendRate = new(RateAndAmountFormat4Choice)
	return c.ProvisionalDividendRate
}

func (c *CorporateActionRate5) AddTaxCreditRate() *TaxCreditRateFormat1Choice {
	newValue := new(TaxCreditRateFormat1Choice)
	c.TaxCreditRate = append(c.TaxCreditRate, newValue)
	return newValue
}

func (c *CorporateActionRate5) AddProrationRate() *RateFormat2Choice {
	c.ProrationRate = new(RateFormat2Choice)
	return c.ProrationRate
}

func (c *CorporateActionRate5) AddSolicitationFeeRate() *SolicitationFeeRateFormat1Choice {
	c.SolicitationFeeRate = new(SolicitationFeeRateFormat1Choice)
	return c.SolicitationFeeRate
}

func (c *CorporateActionRate5) AddEarlySolicitationFeeRate() *SolicitationFeeRateFormat1Choice {
	c.EarlySolicitationFeeRate = new(SolicitationFeeRateFormat1Choice)
	return c.EarlySolicitationFeeRate
}

func (c *CorporateActionRate5) AddWithholdingTaxRate() *RateFormat2Choice {
	c.WithholdingTaxRate = new(RateFormat2Choice)
	return c.WithholdingTaxRate
}

func (c *CorporateActionRate5) AddTaxOnIncome() *RateFormat2Choice {
	c.TaxOnIncome = new(RateFormat2Choice)
	return c.TaxOnIncome
}

func (c *CorporateActionRate5) AddTaxOnProfits() *RateFormat2Choice {
	c.TaxOnProfits = new(RateFormat2Choice)
	return c.TaxOnProfits
}

func (c *CorporateActionRate5) AddTaxReclaimRate() *RateFormat2Choice {
	c.TaxReclaimRate = new(RateFormat2Choice)
	return c.TaxReclaimRate
}

func (c *CorporateActionRate5) AddWithholdingOfForeignTax() *RateAndAmountFormat3Choice {
	c.WithholdingOfForeignTax = new(RateAndAmountFormat3Choice)
	return c.WithholdingOfForeignTax
}

func (c *CorporateActionRate5) AddWithholdingOfLocalTax() *RateAndAmountFormat3Choice {
	c.WithholdingOfLocalTax = new(RateAndAmountFormat3Choice)
	return c.WithholdingOfLocalTax
}

func (c *CorporateActionRate5) AddTaxRelatedRate() *RateTypeAndAmountAndStatus6 {
	newValue := new(RateTypeAndAmountAndStatus6)
	c.TaxRelatedRate = append(c.TaxRelatedRate, newValue)
	return newValue
}

func (c *CorporateActionRate5) AddApplicableRate() *RateFormat2Choice {
	c.ApplicableRate = new(RateFormat2Choice)
	return c.ApplicableRate
}

// Specifies rate details.
type CorporateActionRate6 struct {

	// Quantity of additional intermediate securities/new equities awarded for a given quantity of securities derived from subscription.
	AdditionalQuantityForSubscribedResultantSecurities *RatioFormat3Choice `xml:"AddtlQtyForSbcbdRsltntScties,omitempty"`

	// Quantity of additional securities for a given quantity of underlying securities where underlying securities are not exchanged or debited, for example, 1 for 1: 1 new equity credited for every 1 underlying equity = 2 resulting equities.
	AdditionalQuantityForExistingSecurities *RatioFormat3Choice `xml:"AddtlQtyForExstgScties,omitempty"`

	// Quantity of new securities for a given quantity of underlying securities, where the underlying securities will be exchanged or debited, for example, 2 for 1: 2 new equities credited for every 1 underlying equity debited = 2 resulting equities.
	NewToOld *RatioFormat4Choice `xml:"NewToOd,omitempty"`

	// Quantity of new equities that will be derived by the exercise of a given quantity of intermediate securities.
	NewSecuritiesToUnderlyingSecurities *RatioFormat4Choice `xml:"NewSctiesToUndrlygScties,omitempty"`
}

func (c *CorporateActionRate6) AddAdditionalQuantityForSubscribedResultantSecurities() *RatioFormat3Choice {
	c.AdditionalQuantityForSubscribedResultantSecurities = new(RatioFormat3Choice)
	return c.AdditionalQuantityForSubscribedResultantSecurities
}

func (c *CorporateActionRate6) AddAdditionalQuantityForExistingSecurities() *RatioFormat3Choice {
	c.AdditionalQuantityForExistingSecurities = new(RatioFormat3Choice)
	return c.AdditionalQuantityForExistingSecurities
}

func (c *CorporateActionRate6) AddNewToOld() *RatioFormat4Choice {
	c.NewToOld = new(RatioFormat4Choice)
	return c.NewToOld
}

func (c *CorporateActionRate6) AddNewSecuritiesToUnderlyingSecurities() *RatioFormat4Choice {
	c.NewSecuritiesToUnderlyingSecurities = new(RatioFormat4Choice)
	return c.NewSecuritiesToUnderlyingSecurities
}

// Specifies rates of a corporate action.
type CorporateActionRate66 struct {

	// Annual rate of a financial instrument.
	Interest *RateAndAmountFormat37Choice `xml:"Intrst,omitempty"`

	// Percentage of securities the offeror/issuer will purchase or redeem under the terms of the event.
	PercentageSought *RateFormat7Choice `xml:"PctgSght,omitempty"`

	// Index rate related to the interest rate of the forthcoming interest payment.
	RelatedIndex *RateFormat3Choice `xml:"RltdIndx,omitempty"`

	// Margin allowed over or under a given rate.
	Spread *RateFormat3Choice `xml:"Sprd,omitempty"`

	// Acceptable price increment used for submitting a bid.
	BidInterval *RateAndAmountFormat38Choice `xml:"BidIntrvl,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *RateFormat12Choice `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *RateFormat12Choice `xml:"NxtFctr,omitempty"`

	// Rate of discount for securities purchased through a reinvestment scheme as compared to the current market price of security.
	ReinvestmentDiscountRateToMarket *RateFormat3Choice `xml:"RinvstmtDscntRateToMkt,omitempty"`

	// For structured security issues where there is a set schedule of principal and interest payments for the life of the issue, this is the difference between the actual rate of the interest payment and the expected or scheduled rate of the interest payment .
	InterestShortfall *RateAndAmountFormat39Choice `xml:"IntrstShrtfll,omitempty"`

	// For structured security issues where there is a set schedule of principal and interest payments for the life of the issue, this is the difference between the actual rate of the capital or principal repayment and the scheduled capital repayment.
	RealisedLoss *RateAndAmountFormat39Choice `xml:"RealsdLoss,omitempty"`

	// Dividend or interest rate declared by the issuer.
	DeclaredRate *RateAndAmountFormat39Choice `xml:"DclrdRate,omitempty"`
}

func (c *CorporateActionRate66) AddInterest() *RateAndAmountFormat37Choice {
	c.Interest = new(RateAndAmountFormat37Choice)
	return c.Interest
}

func (c *CorporateActionRate66) AddPercentageSought() *RateFormat7Choice {
	c.PercentageSought = new(RateFormat7Choice)
	return c.PercentageSought
}

func (c *CorporateActionRate66) AddRelatedIndex() *RateFormat3Choice {
	c.RelatedIndex = new(RateFormat3Choice)
	return c.RelatedIndex
}

func (c *CorporateActionRate66) AddSpread() *RateFormat3Choice {
	c.Spread = new(RateFormat3Choice)
	return c.Spread
}

func (c *CorporateActionRate66) AddBidInterval() *RateAndAmountFormat38Choice {
	c.BidInterval = new(RateAndAmountFormat38Choice)
	return c.BidInterval
}

func (c *CorporateActionRate66) AddPreviousFactor() *RateFormat12Choice {
	c.PreviousFactor = new(RateFormat12Choice)
	return c.PreviousFactor
}

func (c *CorporateActionRate66) AddNextFactor() *RateFormat12Choice {
	c.NextFactor = new(RateFormat12Choice)
	return c.NextFactor
}

func (c *CorporateActionRate66) AddReinvestmentDiscountRateToMarket() *RateFormat3Choice {
	c.ReinvestmentDiscountRateToMarket = new(RateFormat3Choice)
	return c.ReinvestmentDiscountRateToMarket
}

func (c *CorporateActionRate66) AddInterestShortfall() *RateAndAmountFormat39Choice {
	c.InterestShortfall = new(RateAndAmountFormat39Choice)
	return c.InterestShortfall
}

func (c *CorporateActionRate66) AddRealisedLoss() *RateAndAmountFormat39Choice {
	c.RealisedLoss = new(RateAndAmountFormat39Choice)
	return c.RealisedLoss
}

func (c *CorporateActionRate66) AddDeclaredRate() *RateAndAmountFormat39Choice {
	c.DeclaredRate = new(RateAndAmountFormat39Choice)
	return c.DeclaredRate
}

// Specifies rates related to a corporate action option.
type CorporateActionRate67 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat37Choice `xml:"AddtlTax,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat19Choice `xml:"GrssDvddRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat21Choice `xml:"NetDvddRate,omitempty"`

	// Public index rate applied to the amount paid to adjust it to inflation.
	IndexFactor *RateAndAmountFormat37Choice `xml:"IndxFctr,omitempty"`

	// The actual interest rate used for the payment of the interest for the specified interest period.
	// Usage guideline: It is used to provide the applicable rate for the current payment, after all calculations have been performed, that is, application of period and method of interest computation.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat8Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Maximum percentage of shares available through the over subscription privilege, usually a percentage of the basic subscription shares, for example, an account owner subscribing to 100 shares may over subscribe to a maximum of 50 additional shares when the over subscription maximum is 50 percent.
	MaximumAllowedOversubscriptionRate *RateFormat3Choice `xml:"MaxAllwdOvrsbcptRate,omitempty"`

	// Proportionate allocation used for the offer.
	ProrationRate *RateFormat3Choice `xml:"PrratnRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxRate []*RateAndAmountFormat41Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate at which the income will be withheld by a jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate (TAXR) levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTax []*RateAndAmountFormat41Choice `xml:"ScndLvlTax,omitempty"`

	// Amount included in the dividend/NAV that is identified as gains directly or indirectly derived from interest payments, for example, in the context of the EU Savings directive.
	TaxableIncomePerDividendShare []*RateTypeAndAmountAndStatus26 `xml:"TaxblIncmPerDvddShr,omitempty"`

	// Overall tax withheld at source by fund managers prior to considering the tax obligation of each unit holder.
	TaxOnIncome *RateAndAmountFormat37Choice `xml:"TaxOnIncm,omitempty"`
}

func (c *CorporateActionRate67) AddAdditionalTax() *RateAndAmountFormat37Choice {
	c.AdditionalTax = new(RateAndAmountFormat37Choice)
	return c.AdditionalTax
}

func (c *CorporateActionRate67) AddGrossDividendRate() *GrossDividendRateFormat19Choice {
	newValue := new(GrossDividendRateFormat19Choice)
	c.GrossDividendRate = append(c.GrossDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate67) AddNetDividendRate() *NetDividendRateFormat21Choice {
	newValue := new(NetDividendRateFormat21Choice)
	c.NetDividendRate = append(c.NetDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate67) AddIndexFactor() *RateAndAmountFormat37Choice {
	c.IndexFactor = new(RateAndAmountFormat37Choice)
	return c.IndexFactor
}

func (c *CorporateActionRate67) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat8Choice {
	newValue := new(InterestRateUsedForPaymentFormat8Choice)
	c.InterestRateUsedForPayment = append(c.InterestRateUsedForPayment, newValue)
	return newValue
}

func (c *CorporateActionRate67) AddMaximumAllowedOversubscriptionRate() *RateFormat3Choice {
	c.MaximumAllowedOversubscriptionRate = new(RateFormat3Choice)
	return c.MaximumAllowedOversubscriptionRate
}

func (c *CorporateActionRate67) AddProrationRate() *RateFormat3Choice {
	c.ProrationRate = new(RateFormat3Choice)
	return c.ProrationRate
}

func (c *CorporateActionRate67) AddWithholdingTaxRate() *RateAndAmountFormat41Choice {
	newValue := new(RateAndAmountFormat41Choice)
	c.WithholdingTaxRate = append(c.WithholdingTaxRate, newValue)
	return newValue
}

func (c *CorporateActionRate67) AddSecondLevelTax() *RateAndAmountFormat41Choice {
	newValue := new(RateAndAmountFormat41Choice)
	c.SecondLevelTax = append(c.SecondLevelTax, newValue)
	return newValue
}

func (c *CorporateActionRate67) AddTaxableIncomePerDividendShare() *RateTypeAndAmountAndStatus26 {
	newValue := new(RateTypeAndAmountAndStatus26)
	c.TaxableIncomePerDividendShare = append(c.TaxableIncomePerDividendShare, newValue)
	return newValue
}

func (c *CorporateActionRate67) AddTaxOnIncome() *RateAndAmountFormat37Choice {
	c.TaxOnIncome = new(RateAndAmountFormat37Choice)
	return c.TaxOnIncome
}

// Specifies rates related to a corporate action option.
type CorporateActionRate68 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat37Choice `xml:"AddtlTax,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat19Choice `xml:"GrssDvddRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat21Choice `xml:"NetDvddRate,omitempty"`

	// Public index rate applied to the amount paid to adjust it to inflation.
	IndexFactor *RateAndAmountFormat37Choice `xml:"IndxFctr,omitempty"`

	// The actual interest rate used for the payment of the interest for the specified interest period.
	// Usage guideline: It is used to provide the applicable rate for the current payment, after all calculations have been performed, that is, application of period and method of interest computation.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat8Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Maximum percentage of shares available through the over subscription privilege, usually a percentage of the basic subscription shares, for example, an account owner subscribing to 100 shares may over subscribe to a maximum of 50 additional shares when the over subscription maximum is 50 percent.
	MaximumAllowedOversubscriptionRate *RateFormat3Choice `xml:"MaxAllwdOvrsbcptRate,omitempty"`

	// Proportionate allocation used for the offer.
	ProrationRate *RateFormat3Choice `xml:"PrratnRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxRate []*RateAndAmountFormat41Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate at which the income will be withheld by a jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate (TAXR) levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTax []*RateAndAmountFormat41Choice `xml:"ScndLvlTax,omitempty"`

	// Amount included in the dividend/NAV that is identified as gains directly or indirectly derived from interest payments, for example, in the context of the EU Savings directive.
	TaxableIncomePerDividendShare []*RateTypeAndAmountAndStatus26 `xml:"TaxblIncmPerDvddShr,omitempty"`

	// Exchange rate (provided by the issuer) between the dividend or interest rate in the paid currency and the declared dividend or interest rate.
	IssuerDeclaredExchangeRate *ForeignExchangeTerms19 `xml:"IssrDclrdXchgRate,omitempty"`

	// Overall tax withheld at source by fund managers prior to considering the tax obligation of each unit holder.
	TaxOnIncome *RateAndAmountFormat37Choice `xml:"TaxOnIncm,omitempty"`
}

func (c *CorporateActionRate68) AddAdditionalTax() *RateAndAmountFormat37Choice {
	c.AdditionalTax = new(RateAndAmountFormat37Choice)
	return c.AdditionalTax
}

func (c *CorporateActionRate68) AddGrossDividendRate() *GrossDividendRateFormat19Choice {
	newValue := new(GrossDividendRateFormat19Choice)
	c.GrossDividendRate = append(c.GrossDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate68) AddNetDividendRate() *NetDividendRateFormat21Choice {
	newValue := new(NetDividendRateFormat21Choice)
	c.NetDividendRate = append(c.NetDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate68) AddIndexFactor() *RateAndAmountFormat37Choice {
	c.IndexFactor = new(RateAndAmountFormat37Choice)
	return c.IndexFactor
}

func (c *CorporateActionRate68) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat8Choice {
	newValue := new(InterestRateUsedForPaymentFormat8Choice)
	c.InterestRateUsedForPayment = append(c.InterestRateUsedForPayment, newValue)
	return newValue
}

func (c *CorporateActionRate68) AddMaximumAllowedOversubscriptionRate() *RateFormat3Choice {
	c.MaximumAllowedOversubscriptionRate = new(RateFormat3Choice)
	return c.MaximumAllowedOversubscriptionRate
}

func (c *CorporateActionRate68) AddProrationRate() *RateFormat3Choice {
	c.ProrationRate = new(RateFormat3Choice)
	return c.ProrationRate
}

func (c *CorporateActionRate68) AddWithholdingTaxRate() *RateAndAmountFormat41Choice {
	newValue := new(RateAndAmountFormat41Choice)
	c.WithholdingTaxRate = append(c.WithholdingTaxRate, newValue)
	return newValue
}

func (c *CorporateActionRate68) AddSecondLevelTax() *RateAndAmountFormat41Choice {
	newValue := new(RateAndAmountFormat41Choice)
	c.SecondLevelTax = append(c.SecondLevelTax, newValue)
	return newValue
}

func (c *CorporateActionRate68) AddTaxableIncomePerDividendShare() *RateTypeAndAmountAndStatus26 {
	newValue := new(RateTypeAndAmountAndStatus26)
	c.TaxableIncomePerDividendShare = append(c.TaxableIncomePerDividendShare, newValue)
	return newValue
}

func (c *CorporateActionRate68) AddIssuerDeclaredExchangeRate() *ForeignExchangeTerms19 {
	c.IssuerDeclaredExchangeRate = new(ForeignExchangeTerms19)
	return c.IssuerDeclaredExchangeRate
}

func (c *CorporateActionRate68) AddTaxOnIncome() *RateAndAmountFormat37Choice {
	c.TaxOnIncome = new(RateAndAmountFormat37Choice)
	return c.TaxOnIncome
}

// Specifies security rate details.
type CorporateActionRate69 struct {

	// Quantity of additional intermediate securities/new equities awarded for a given quantity of securities derived from subscription.
	AdditionalQuantityForSubscribedResultantSecurities *RatioFormat17Choice `xml:"AddtlQtyForSbcbdRsltntScties,omitempty"`

	// Quantity of additional securities for a given quantity of underlying securities where underlying securities are not exchanged or debited, for example, 1 for 1: 1 new equity credited for every 1 underlying equity = 2 resulting equities.
	AdditionalQuantityForExistingSecurities *RatioFormat17Choice `xml:"AddtlQtyForExstgScties,omitempty"`

	// Quantity of new securities for a given quantity of underlying securities, where the underlying securities will be exchanged or debited, for example, 2 for 1: 2 new equities credited for every 1 underlying equity debited = 2 resulting equities.
	NewToOld *RatioFormat18Choice `xml:"NewToOd,omitempty"`

	// Rate used to determine the cash consideration split across outturn settlement transactions that are the result of a transformation of the parent transaction.
	TransformationRate *PercentageRate `xml:"TrfrmatnRate,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat37Choice `xml:"ChrgsFees,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *RateFormat3Choice `xml:"FsclStmp,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *RateFormat3Choice `xml:"AplblRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat8Choice `xml:"TaxCdtRate,omitempty"`

	// Rate of financial transaction tax.
	FinancialTransactionTaxRate *RateFormat3Choice `xml:"FinTxTaxRate,omitempty"`
}

func (c *CorporateActionRate69) AddAdditionalQuantityForSubscribedResultantSecurities() *RatioFormat17Choice {
	c.AdditionalQuantityForSubscribedResultantSecurities = new(RatioFormat17Choice)
	return c.AdditionalQuantityForSubscribedResultantSecurities
}

func (c *CorporateActionRate69) AddAdditionalQuantityForExistingSecurities() *RatioFormat17Choice {
	c.AdditionalQuantityForExistingSecurities = new(RatioFormat17Choice)
	return c.AdditionalQuantityForExistingSecurities
}

func (c *CorporateActionRate69) AddNewToOld() *RatioFormat18Choice {
	c.NewToOld = new(RatioFormat18Choice)
	return c.NewToOld
}

func (c *CorporateActionRate69) SetTransformationRate(value string) {
	c.TransformationRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate69) AddChargesFees() *RateAndAmountFormat37Choice {
	c.ChargesFees = new(RateAndAmountFormat37Choice)
	return c.ChargesFees
}

func (c *CorporateActionRate69) AddFiscalStamp() *RateFormat3Choice {
	c.FiscalStamp = new(RateFormat3Choice)
	return c.FiscalStamp
}

func (c *CorporateActionRate69) AddApplicableRate() *RateFormat3Choice {
	c.ApplicableRate = new(RateFormat3Choice)
	return c.ApplicableRate
}

func (c *CorporateActionRate69) AddTaxCreditRate() *TaxCreditRateFormat8Choice {
	newValue := new(TaxCreditRateFormat8Choice)
	c.TaxCreditRate = append(c.TaxCreditRate, newValue)
	return newValue
}

func (c *CorporateActionRate69) AddFinancialTransactionTaxRate() *RateFormat3Choice {
	c.FinancialTransactionTaxRate = new(RateFormat3Choice)
	return c.FinancialTransactionTaxRate
}

// Specifies security rate details.
type CorporateActionRate7 struct {

	// Quantity of additional intermediate securities/new equities awarded for a given quantity of securities derived from subscription.
	AdditionalQuantityForSubscribedResultantSecurities *RatioFormat5Choice `xml:"AddtlQtyForSbcbdRsltntScties,omitempty"`

	// Quantity of additional securities for a given quantity of underlying securities where underlying securities are not exchanged or debited, for example, 1 for 1: 1 new equity credited for every 1 underlying equity = 2 resulting equities.
	AdditionalQuantityForExistingSecurities *RatioFormat5Choice `xml:"AddtlQtyForExstgScties,omitempty"`

	// Quantity of new securities for a given quantity of underlying securities, where the underlying securities will be exchanged or debited, for example, 2 for 1: 2 new equities credited for every 1 underlying equity debited = 2 resulting equities.
	NewToOld *RatioFormat6Choice `xml:"NewToOd,omitempty"`

	// Quantity of new equities that will be derived by the exercise of a given quantity of intermediate securities.
	NewSecuritiesToUnderlyingSecurities *RatioFormat6Choice `xml:"NewSctiesToUndrlygScties,omitempty"`

	// Rate used to determine the cash consideration split across outturn settlement transactions that are the result of a transformation of the parent transaction.
	TransformationRate *PercentageRate `xml:"TrfrmatnRate,omitempty"`
}

func (c *CorporateActionRate7) AddAdditionalQuantityForSubscribedResultantSecurities() *RatioFormat5Choice {
	c.AdditionalQuantityForSubscribedResultantSecurities = new(RatioFormat5Choice)
	return c.AdditionalQuantityForSubscribedResultantSecurities
}

func (c *CorporateActionRate7) AddAdditionalQuantityForExistingSecurities() *RatioFormat5Choice {
	c.AdditionalQuantityForExistingSecurities = new(RatioFormat5Choice)
	return c.AdditionalQuantityForExistingSecurities
}

func (c *CorporateActionRate7) AddNewToOld() *RatioFormat6Choice {
	c.NewToOld = new(RatioFormat6Choice)
	return c.NewToOld
}

func (c *CorporateActionRate7) AddNewSecuritiesToUnderlyingSecurities() *RatioFormat6Choice {
	c.NewSecuritiesToUnderlyingSecurities = new(RatioFormat6Choice)
	return c.NewSecuritiesToUnderlyingSecurities
}

func (c *CorporateActionRate7) SetTransformationRate(value string) {
	c.TransformationRate = (*PercentageRate)(&value)
}

// Specifies rates.
type CorporateActionRate70 struct {

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat21Choice `xml:"GrssDvddRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat23Choice `xml:"NetDvddRate,omitempty"`

	// Public index rate applied to the amount paid to adjust it to inflation.
	IndexFactor *RateAndAmountFormat39Choice `xml:"IndxFctr,omitempty"`

	// Actual interest rate used for the payment of the interest for the specified interest period.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat7Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// A maximum percentage of shares available through the over subscription privilege, usually a percentage of the basic subscription shares, for example, an account owner subscribing to 100 shares may over subscribe to a maximum of 50 additional shares when the over subscription maximum is 50 percent.
	MaximumAllowedOversubscriptionRate *PercentageRate `xml:"MaxAllwdOvrsbcptRate,omitempty"`

	// Proportionate allocation used for the offer.
	ProrationRate *PercentageRate `xml:"PrratnRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxRate []*RateAndAmountFormat40Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate at which the income will be withheld by a jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate (TAXR) levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTax []*RateAndAmountFormat40Choice `xml:"ScndLvlTax,omitempty"`

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat39Choice `xml:"AddtlTax,omitempty"`

	// Amount included in the dividend/NAV that is identified as gains directly or indirectly derived from interest payments, for example, in the context of the EU Savings directive.
	TaxableIncomePerDividendShare []*RateTypeAndAmountAndStatus26 `xml:"TaxblIncmPerDvddShr,omitempty"`
}

func (c *CorporateActionRate70) AddGrossDividendRate() *GrossDividendRateFormat21Choice {
	newValue := new(GrossDividendRateFormat21Choice)
	c.GrossDividendRate = append(c.GrossDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate70) AddNetDividendRate() *NetDividendRateFormat23Choice {
	newValue := new(NetDividendRateFormat23Choice)
	c.NetDividendRate = append(c.NetDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate70) AddIndexFactor() *RateAndAmountFormat39Choice {
	c.IndexFactor = new(RateAndAmountFormat39Choice)
	return c.IndexFactor
}

func (c *CorporateActionRate70) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat7Choice {
	newValue := new(InterestRateUsedForPaymentFormat7Choice)
	c.InterestRateUsedForPayment = append(c.InterestRateUsedForPayment, newValue)
	return newValue
}

func (c *CorporateActionRate70) SetMaximumAllowedOversubscriptionRate(value string) {
	c.MaximumAllowedOversubscriptionRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate70) SetProrationRate(value string) {
	c.ProrationRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate70) AddWithholdingTaxRate() *RateAndAmountFormat40Choice {
	newValue := new(RateAndAmountFormat40Choice)
	c.WithholdingTaxRate = append(c.WithholdingTaxRate, newValue)
	return newValue
}

func (c *CorporateActionRate70) AddSecondLevelTax() *RateAndAmountFormat40Choice {
	newValue := new(RateAndAmountFormat40Choice)
	c.SecondLevelTax = append(c.SecondLevelTax, newValue)
	return newValue
}

func (c *CorporateActionRate70) AddAdditionalTax() *RateAndAmountFormat39Choice {
	c.AdditionalTax = new(RateAndAmountFormat39Choice)
	return c.AdditionalTax
}

func (c *CorporateActionRate70) AddTaxableIncomePerDividendShare() *RateTypeAndAmountAndStatus26 {
	newValue := new(RateTypeAndAmountAndStatus26)
	c.TaxableIncomePerDividendShare = append(c.TaxableIncomePerDividendShare, newValue)
	return newValue
}

// Specifies rates related to a corporate action option.
type CorporateActionRate71 struct {

	// Rate proposed in a remarketing of variable rate notes.
	ProposedRate *PercentageRate `xml:"PropsdRate,omitempty"`

	// Rate of allowed over-subscription.
	OversubscriptionRate *RateAndAmountFormat39Choice `xml:"OvrsbcptRate,omitempty"`

	// Requested tax rate that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	RequestedWithholdingTaxRate []*RateAndAmountFormat40Choice `xml:"ReqdWhldgTaxRate,omitempty"`

	// Requested rate at which the income will be withheld by the jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible.
	RequestedSecondLevelTaxRate []*RateAndAmountFormat40Choice `xml:"ReqdScndLvlTaxRate,omitempty"`
}

func (c *CorporateActionRate71) SetProposedRate(value string) {
	c.ProposedRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate71) AddOversubscriptionRate() *RateAndAmountFormat39Choice {
	c.OversubscriptionRate = new(RateAndAmountFormat39Choice)
	return c.OversubscriptionRate
}

func (c *CorporateActionRate71) AddRequestedWithholdingTaxRate() *RateAndAmountFormat40Choice {
	newValue := new(RateAndAmountFormat40Choice)
	c.RequestedWithholdingTaxRate = append(c.RequestedWithholdingTaxRate, newValue)
	return newValue
}

func (c *CorporateActionRate71) AddRequestedSecondLevelTaxRate() *RateAndAmountFormat40Choice {
	newValue := new(RateAndAmountFormat40Choice)
	c.RequestedSecondLevelTaxRate = append(c.RequestedSecondLevelTaxRate, newValue)
	return newValue
}

// Specifies rate details.
type CorporateActionRate72 struct {

	// Quantity of additional intermediate securities/new equities awarded for a given quantity of securities derived from subscription.
	AdditionalQuantityForSubscribedResultantSecurities *RatioFormat20Choice `xml:"AddtlQtyForSbcbdRsltntScties,omitempty"`

	// Quantity of additional securities for a given quantity of underlying securities where underlying securities are not exchanged or debited, for example, 1 for 1: 1 new equity credited for every 1 underlying equity = 2 resulting equities.
	AdditionalQuantityForExistingSecurities *RatioFormat20Choice `xml:"AddtlQtyForExstgScties,omitempty"`

	// Quantity of new securities for a given quantity of underlying securities, where the underlying securities will be exchanged or debited, for example, 2 for 1: 2 new equities credited for every 1 underlying equity debited = 2 resulting equities.
	NewToOld *RatioFormat19Choice `xml:"NewToOd,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat39Choice `xml:"ChrgsFees,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *PercentageRate `xml:"FsclStmp,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *PercentageRate `xml:"AplblRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat7Choice `xml:"TaxCdtRate,omitempty"`

	// Rate of financial transaction tax.
	FinancialTransactionTaxRate *PercentageRate `xml:"FinTxTaxRate,omitempty"`
}

func (c *CorporateActionRate72) AddAdditionalQuantityForSubscribedResultantSecurities() *RatioFormat20Choice {
	c.AdditionalQuantityForSubscribedResultantSecurities = new(RatioFormat20Choice)
	return c.AdditionalQuantityForSubscribedResultantSecurities
}

func (c *CorporateActionRate72) AddAdditionalQuantityForExistingSecurities() *RatioFormat20Choice {
	c.AdditionalQuantityForExistingSecurities = new(RatioFormat20Choice)
	return c.AdditionalQuantityForExistingSecurities
}

func (c *CorporateActionRate72) AddNewToOld() *RatioFormat19Choice {
	c.NewToOld = new(RatioFormat19Choice)
	return c.NewToOld
}

func (c *CorporateActionRate72) AddChargesFees() *RateAndAmountFormat39Choice {
	c.ChargesFees = new(RateAndAmountFormat39Choice)
	return c.ChargesFees
}

func (c *CorporateActionRate72) SetFiscalStamp(value string) {
	c.FiscalStamp = (*PercentageRate)(&value)
}

func (c *CorporateActionRate72) SetApplicableRate(value string) {
	c.ApplicableRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate72) AddTaxCreditRate() *TaxCreditRateFormat7Choice {
	newValue := new(TaxCreditRateFormat7Choice)
	c.TaxCreditRate = append(c.TaxCreditRate, newValue)
	return newValue
}

func (c *CorporateActionRate72) SetFinancialTransactionTaxRate(value string) {
	c.FinancialTransactionTaxRate = (*PercentageRate)(&value)
}

// Specifies rates related to a corporate action option.
type CorporateActionRate73 struct {

	// Rate proposed in a remarketing of variable rate notes.
	ProposedRate *PercentageRate `xml:"PropsdRate,omitempty"`

	// Rate of allowed over-subscription.
	OversubscriptionRate *RateAndAmountFormat43Choice `xml:"OvrsbcptRate,omitempty"`

	// Requested tax rate that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	RequestedWithholdingTaxRate []*RateAndAmountFormat45Choice `xml:"ReqdWhldgTaxRate,omitempty"`

	// Requested rate at which the income will be withheld by the jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible.
	RequestedSecondLevelTaxRate []*RateAndAmountFormat45Choice `xml:"ReqdScndLvlTaxRate,omitempty"`
}

func (c *CorporateActionRate73) SetProposedRate(value string) {
	c.ProposedRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate73) AddOversubscriptionRate() *RateAndAmountFormat43Choice {
	c.OversubscriptionRate = new(RateAndAmountFormat43Choice)
	return c.OversubscriptionRate
}

func (c *CorporateActionRate73) AddRequestedWithholdingTaxRate() *RateAndAmountFormat45Choice {
	newValue := new(RateAndAmountFormat45Choice)
	c.RequestedWithholdingTaxRate = append(c.RequestedWithholdingTaxRate, newValue)
	return newValue
}

func (c *CorporateActionRate73) AddRequestedSecondLevelTaxRate() *RateAndAmountFormat45Choice {
	newValue := new(RateAndAmountFormat45Choice)
	c.RequestedSecondLevelTaxRate = append(c.RequestedSecondLevelTaxRate, newValue)
	return newValue
}

// Specifies rates.
type CorporateActionRate74 struct {

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat23Choice `xml:"GrssDvddRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat25Choice `xml:"NetDvddRate,omitempty"`

	// Public index rate applied to the amount paid to adjust it to inflation.
	IndexFactor *RateAndAmountFormat43Choice `xml:"IndxFctr,omitempty"`

	// Actual interest rate used for the payment of the interest for the specified interest period.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat9Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// A maximum percentage of shares available through the over subscription privilege, usually a percentage of the basic subscription shares, for example, an account owner subscribing to 100 shares may over subscribe to a maximum of 50 additional shares when the over subscription maximum is 50 percent.
	MaximumAllowedOversubscriptionRate *PercentageRate `xml:"MaxAllwdOvrsbcptRate,omitempty"`

	// Proportionate allocation used for the offer.
	ProrationRate *PercentageRate `xml:"PrratnRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxRate []*RateAndAmountFormat45Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate at which the income will be withheld by a jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate (TAXR) levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTax []*RateAndAmountFormat45Choice `xml:"ScndLvlTax,omitempty"`

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat43Choice `xml:"AddtlTax,omitempty"`

	// Amount included in the dividend/NAV that is identified as gains directly or indirectly derived from interest payments, for example, in the context of the EU Savings directive.
	TaxableIncomePerDividendShare []*RateTypeAndAmountAndStatus33 `xml:"TaxblIncmPerDvddShr,omitempty"`
}

func (c *CorporateActionRate74) AddGrossDividendRate() *GrossDividendRateFormat23Choice {
	newValue := new(GrossDividendRateFormat23Choice)
	c.GrossDividendRate = append(c.GrossDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate74) AddNetDividendRate() *NetDividendRateFormat25Choice {
	newValue := new(NetDividendRateFormat25Choice)
	c.NetDividendRate = append(c.NetDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate74) AddIndexFactor() *RateAndAmountFormat43Choice {
	c.IndexFactor = new(RateAndAmountFormat43Choice)
	return c.IndexFactor
}

func (c *CorporateActionRate74) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat9Choice {
	newValue := new(InterestRateUsedForPaymentFormat9Choice)
	c.InterestRateUsedForPayment = append(c.InterestRateUsedForPayment, newValue)
	return newValue
}

func (c *CorporateActionRate74) SetMaximumAllowedOversubscriptionRate(value string) {
	c.MaximumAllowedOversubscriptionRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate74) SetProrationRate(value string) {
	c.ProrationRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate74) AddWithholdingTaxRate() *RateAndAmountFormat45Choice {
	newValue := new(RateAndAmountFormat45Choice)
	c.WithholdingTaxRate = append(c.WithholdingTaxRate, newValue)
	return newValue
}

func (c *CorporateActionRate74) AddSecondLevelTax() *RateAndAmountFormat45Choice {
	newValue := new(RateAndAmountFormat45Choice)
	c.SecondLevelTax = append(c.SecondLevelTax, newValue)
	return newValue
}

func (c *CorporateActionRate74) AddAdditionalTax() *RateAndAmountFormat43Choice {
	c.AdditionalTax = new(RateAndAmountFormat43Choice)
	return c.AdditionalTax
}

func (c *CorporateActionRate74) AddTaxableIncomePerDividendShare() *RateTypeAndAmountAndStatus33 {
	newValue := new(RateTypeAndAmountAndStatus33)
	c.TaxableIncomePerDividendShare = append(c.TaxableIncomePerDividendShare, newValue)
	return newValue
}

// Specifies rate details.
type CorporateActionRate75 struct {

	// Quantity of additional intermediate securities/new equities awarded for a given quantity of securities derived from subscription.
	AdditionalQuantityForSubscribedResultantSecurities *RatioFormat21Choice `xml:"AddtlQtyForSbcbdRsltntScties,omitempty"`

	// Quantity of additional securities for a given quantity of underlying securities where underlying securities are not exchanged or debited, for example, 1 for 1: 1 new equity credited for every 1 underlying equity = 2 resulting equities.
	AdditionalQuantityForExistingSecurities *RatioFormat21Choice `xml:"AddtlQtyForExstgScties,omitempty"`

	// Quantity of new securities for a given quantity of underlying securities, where the underlying securities will be exchanged or debited, for example, 2 for 1: 2 new equities credited for every 1 underlying equity debited = 2 resulting equities.
	NewToOld *RatioFormat22Choice `xml:"NewToOd,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat43Choice `xml:"ChrgsFees,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *PercentageRate `xml:"FsclStmp,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *PercentageRate `xml:"AplblRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat9Choice `xml:"TaxCdtRate,omitempty"`

	// Rate of financial transaction tax.
	FinancialTransactionTaxRate *PercentageRate `xml:"FinTxTaxRate,omitempty"`
}

func (c *CorporateActionRate75) AddAdditionalQuantityForSubscribedResultantSecurities() *RatioFormat21Choice {
	c.AdditionalQuantityForSubscribedResultantSecurities = new(RatioFormat21Choice)
	return c.AdditionalQuantityForSubscribedResultantSecurities
}

func (c *CorporateActionRate75) AddAdditionalQuantityForExistingSecurities() *RatioFormat21Choice {
	c.AdditionalQuantityForExistingSecurities = new(RatioFormat21Choice)
	return c.AdditionalQuantityForExistingSecurities
}

func (c *CorporateActionRate75) AddNewToOld() *RatioFormat22Choice {
	c.NewToOld = new(RatioFormat22Choice)
	return c.NewToOld
}

func (c *CorporateActionRate75) AddChargesFees() *RateAndAmountFormat43Choice {
	c.ChargesFees = new(RateAndAmountFormat43Choice)
	return c.ChargesFees
}

func (c *CorporateActionRate75) SetFiscalStamp(value string) {
	c.FiscalStamp = (*PercentageRate)(&value)
}

func (c *CorporateActionRate75) SetApplicableRate(value string) {
	c.ApplicableRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate75) AddTaxCreditRate() *TaxCreditRateFormat9Choice {
	newValue := new(TaxCreditRateFormat9Choice)
	c.TaxCreditRate = append(c.TaxCreditRate, newValue)
	return newValue
}

func (c *CorporateActionRate75) SetFinancialTransactionTaxRate(value string) {
	c.FinancialTransactionTaxRate = (*PercentageRate)(&value)
}

// Specifies rates related to a corporate action option.
type CorporateActionRate76 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat46Choice `xml:"AddtlTax,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat25Choice `xml:"GrssDvddRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat27Choice `xml:"NetDvddRate,omitempty"`

	// Public index rate applied to the amount paid to adjust it to inflation.
	IndexFactor *RateAndAmountFormat46Choice `xml:"IndxFctr,omitempty"`

	// The actual interest rate used for the payment of the interest for the specified interest period.
	// Usage guideline: It is used to provide the applicable rate for the current payment, after all calculations have been performed, that is, application of period and method of interest computation.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat10Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Maximum percentage of shares available through the over subscription privilege, usually a percentage of the basic subscription shares, for example, an account owner subscribing to 100 shares may over subscribe to a maximum of 50 additional shares when the over subscription maximum is 50 percent.
	MaximumAllowedOversubscriptionRate *RateFormat3Choice `xml:"MaxAllwdOvrsbcptRate,omitempty"`

	// Proportionate allocation used for the offer.
	ProrationRate *RateFormat3Choice `xml:"PrratnRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxRate []*RateAndAmountFormat47Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate at which the income will be withheld by a jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate (TAXR) levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTax []*RateAndAmountFormat47Choice `xml:"ScndLvlTax,omitempty"`

	// Amount included in the dividend/NAV that is identified as gains directly or indirectly derived from interest payments, for example, in the context of the EU Savings directive.
	TaxableIncomePerDividendShare []*RateTypeAndAmountAndStatus33 `xml:"TaxblIncmPerDvddShr,omitempty"`

	// Overall tax withheld at source by fund managers prior to considering the tax obligation of each unit holder.
	TaxOnIncome *RateAndAmountFormat46Choice `xml:"TaxOnIncm,omitempty"`
}

func (c *CorporateActionRate76) AddAdditionalTax() *RateAndAmountFormat46Choice {
	c.AdditionalTax = new(RateAndAmountFormat46Choice)
	return c.AdditionalTax
}

func (c *CorporateActionRate76) AddGrossDividendRate() *GrossDividendRateFormat25Choice {
	newValue := new(GrossDividendRateFormat25Choice)
	c.GrossDividendRate = append(c.GrossDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate76) AddNetDividendRate() *NetDividendRateFormat27Choice {
	newValue := new(NetDividendRateFormat27Choice)
	c.NetDividendRate = append(c.NetDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate76) AddIndexFactor() *RateAndAmountFormat46Choice {
	c.IndexFactor = new(RateAndAmountFormat46Choice)
	return c.IndexFactor
}

func (c *CorporateActionRate76) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat10Choice {
	newValue := new(InterestRateUsedForPaymentFormat10Choice)
	c.InterestRateUsedForPayment = append(c.InterestRateUsedForPayment, newValue)
	return newValue
}

func (c *CorporateActionRate76) AddMaximumAllowedOversubscriptionRate() *RateFormat3Choice {
	c.MaximumAllowedOversubscriptionRate = new(RateFormat3Choice)
	return c.MaximumAllowedOversubscriptionRate
}

func (c *CorporateActionRate76) AddProrationRate() *RateFormat3Choice {
	c.ProrationRate = new(RateFormat3Choice)
	return c.ProrationRate
}

func (c *CorporateActionRate76) AddWithholdingTaxRate() *RateAndAmountFormat47Choice {
	newValue := new(RateAndAmountFormat47Choice)
	c.WithholdingTaxRate = append(c.WithholdingTaxRate, newValue)
	return newValue
}

func (c *CorporateActionRate76) AddSecondLevelTax() *RateAndAmountFormat47Choice {
	newValue := new(RateAndAmountFormat47Choice)
	c.SecondLevelTax = append(c.SecondLevelTax, newValue)
	return newValue
}

func (c *CorporateActionRate76) AddTaxableIncomePerDividendShare() *RateTypeAndAmountAndStatus33 {
	newValue := new(RateTypeAndAmountAndStatus33)
	c.TaxableIncomePerDividendShare = append(c.TaxableIncomePerDividendShare, newValue)
	return newValue
}

func (c *CorporateActionRate76) AddTaxOnIncome() *RateAndAmountFormat46Choice {
	c.TaxOnIncome = new(RateAndAmountFormat46Choice)
	return c.TaxOnIncome
}

// Specifies security rate details.
type CorporateActionRate77 struct {

	// Quantity of additional intermediate securities/new equities awarded for a given quantity of securities derived from subscription.
	AdditionalQuantityForSubscribedResultantSecurities *RatioFormat23Choice `xml:"AddtlQtyForSbcbdRsltntScties,omitempty"`

	// Quantity of additional securities for a given quantity of underlying securities where underlying securities are not exchanged or debited, for example, 1 for 1: 1 new equity credited for every 1 underlying equity = 2 resulting equities.
	AdditionalQuantityForExistingSecurities *RatioFormat23Choice `xml:"AddtlQtyForExstgScties,omitempty"`

	// Quantity of new securities for a given quantity of underlying securities, where the underlying securities will be exchanged or debited, for example, 2 for 1: 2 new equities credited for every 1 underlying equity debited = 2 resulting equities.
	NewToOld *RatioFormat24Choice `xml:"NewToOd,omitempty"`

	// Rate used to determine the cash consideration split across outturn settlement transactions that are the result of a transformation of the parent transaction.
	TransformationRate *PercentageRate `xml:"TrfrmatnRate,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat46Choice `xml:"ChrgsFees,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *RateFormat3Choice `xml:"FsclStmp,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *RateFormat3Choice `xml:"AplblRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat10Choice `xml:"TaxCdtRate,omitempty"`

	// Rate of financial transaction tax.
	FinancialTransactionTaxRate *RateFormat3Choice `xml:"FinTxTaxRate,omitempty"`
}

func (c *CorporateActionRate77) AddAdditionalQuantityForSubscribedResultantSecurities() *RatioFormat23Choice {
	c.AdditionalQuantityForSubscribedResultantSecurities = new(RatioFormat23Choice)
	return c.AdditionalQuantityForSubscribedResultantSecurities
}

func (c *CorporateActionRate77) AddAdditionalQuantityForExistingSecurities() *RatioFormat23Choice {
	c.AdditionalQuantityForExistingSecurities = new(RatioFormat23Choice)
	return c.AdditionalQuantityForExistingSecurities
}

func (c *CorporateActionRate77) AddNewToOld() *RatioFormat24Choice {
	c.NewToOld = new(RatioFormat24Choice)
	return c.NewToOld
}

func (c *CorporateActionRate77) SetTransformationRate(value string) {
	c.TransformationRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate77) AddChargesFees() *RateAndAmountFormat46Choice {
	c.ChargesFees = new(RateAndAmountFormat46Choice)
	return c.ChargesFees
}

func (c *CorporateActionRate77) AddFiscalStamp() *RateFormat3Choice {
	c.FiscalStamp = new(RateFormat3Choice)
	return c.FiscalStamp
}

func (c *CorporateActionRate77) AddApplicableRate() *RateFormat3Choice {
	c.ApplicableRate = new(RateFormat3Choice)
	return c.ApplicableRate
}

func (c *CorporateActionRate77) AddTaxCreditRate() *TaxCreditRateFormat10Choice {
	newValue := new(TaxCreditRateFormat10Choice)
	c.TaxCreditRate = append(c.TaxCreditRate, newValue)
	return newValue
}

func (c *CorporateActionRate77) AddFinancialTransactionTaxRate() *RateFormat3Choice {
	c.FinancialTransactionTaxRate = new(RateFormat3Choice)
	return c.FinancialTransactionTaxRate
}

// Specifies rates of a corporate action.
type CorporateActionRate78 struct {

	// Annual rate of a financial instrument.
	Interest *RateAndAmountFormat46Choice `xml:"Intrst,omitempty"`

	// Percentage of securities the offeror/issuer will purchase or redeem under the terms of the event.
	PercentageSought *RateFormat7Choice `xml:"PctgSght,omitempty"`

	// Index rate related to the interest rate of the forthcoming interest payment.
	RelatedIndex *RateFormat3Choice `xml:"RltdIndx,omitempty"`

	// Margin allowed over or under a given rate.
	Spread *RateFormat3Choice `xml:"Sprd,omitempty"`

	// Acceptable price increment used for submitting a bid.
	BidInterval *RateAndAmountFormat50Choice `xml:"BidIntrvl,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *RateFormat12Choice `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *RateFormat12Choice `xml:"NxtFctr,omitempty"`

	// Rate of discount for securities purchased through a reinvestment scheme as compared to the current market price of security.
	ReinvestmentDiscountRateToMarket *RateFormat3Choice `xml:"RinvstmtDscntRateToMkt,omitempty"`

	// For structured security issues where there is a set schedule of principal and interest payments for the life of the issue, this is the difference between the actual rate of the interest payment and the expected or scheduled rate of the interest payment .
	InterestShortfall *RateAndAmountFormat43Choice `xml:"IntrstShrtfll,omitempty"`

	// For structured security issues where there is a set schedule of principal and interest payments for the life of the issue, this is the difference between the actual rate of the capital or principal repayment and the scheduled capital repayment.
	RealisedLoss *RateAndAmountFormat43Choice `xml:"RealsdLoss,omitempty"`

	// Dividend or interest rate declared by the issuer.
	DeclaredRate *RateAndAmountFormat43Choice `xml:"DclrdRate,omitempty"`
}

func (c *CorporateActionRate78) AddInterest() *RateAndAmountFormat46Choice {
	c.Interest = new(RateAndAmountFormat46Choice)
	return c.Interest
}

func (c *CorporateActionRate78) AddPercentageSought() *RateFormat7Choice {
	c.PercentageSought = new(RateFormat7Choice)
	return c.PercentageSought
}

func (c *CorporateActionRate78) AddRelatedIndex() *RateFormat3Choice {
	c.RelatedIndex = new(RateFormat3Choice)
	return c.RelatedIndex
}

func (c *CorporateActionRate78) AddSpread() *RateFormat3Choice {
	c.Spread = new(RateFormat3Choice)
	return c.Spread
}

func (c *CorporateActionRate78) AddBidInterval() *RateAndAmountFormat50Choice {
	c.BidInterval = new(RateAndAmountFormat50Choice)
	return c.BidInterval
}

func (c *CorporateActionRate78) AddPreviousFactor() *RateFormat12Choice {
	c.PreviousFactor = new(RateFormat12Choice)
	return c.PreviousFactor
}

func (c *CorporateActionRate78) AddNextFactor() *RateFormat12Choice {
	c.NextFactor = new(RateFormat12Choice)
	return c.NextFactor
}

func (c *CorporateActionRate78) AddReinvestmentDiscountRateToMarket() *RateFormat3Choice {
	c.ReinvestmentDiscountRateToMarket = new(RateFormat3Choice)
	return c.ReinvestmentDiscountRateToMarket
}

func (c *CorporateActionRate78) AddInterestShortfall() *RateAndAmountFormat43Choice {
	c.InterestShortfall = new(RateAndAmountFormat43Choice)
	return c.InterestShortfall
}

func (c *CorporateActionRate78) AddRealisedLoss() *RateAndAmountFormat43Choice {
	c.RealisedLoss = new(RateAndAmountFormat43Choice)
	return c.RealisedLoss
}

func (c *CorporateActionRate78) AddDeclaredRate() *RateAndAmountFormat43Choice {
	c.DeclaredRate = new(RateAndAmountFormat43Choice)
	return c.DeclaredRate
}

// Specifies rates related to a corporate action option.
type CorporateActionRate79 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat46Choice `xml:"AddtlTax,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat25Choice `xml:"GrssDvddRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat27Choice `xml:"NetDvddRate,omitempty"`

	// Public index rate applied to the amount paid to adjust it to inflation.
	IndexFactor *RateAndAmountFormat46Choice `xml:"IndxFctr,omitempty"`

	// The actual interest rate used for the payment of the interest for the specified interest period.
	// Usage guideline: It is used to provide the applicable rate for the current payment, after all calculations have been performed, that is, application of period and method of interest computation.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat10Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Maximum percentage of shares available through the over subscription privilege, usually a percentage of the basic subscription shares, for example, an account owner subscribing to 100 shares may over subscribe to a maximum of 50 additional shares when the over subscription maximum is 50 percent.
	MaximumAllowedOversubscriptionRate *RateFormat3Choice `xml:"MaxAllwdOvrsbcptRate,omitempty"`

	// Proportionate allocation used for the offer.
	ProrationRate *RateFormat3Choice `xml:"PrratnRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxRate []*RateAndAmountFormat47Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate at which the income will be withheld by a jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate (TAXR) levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTax []*RateAndAmountFormat47Choice `xml:"ScndLvlTax,omitempty"`

	// Amount included in the dividend/NAV that is identified as gains directly or indirectly derived from interest payments, for example, in the context of the EU Savings directive.
	TaxableIncomePerDividendShare []*RateTypeAndAmountAndStatus33 `xml:"TaxblIncmPerDvddShr,omitempty"`

	// Exchange rate (provided by the issuer) between the dividend or interest rate in the paid currency and the declared dividend or interest rate.
	IssuerDeclaredExchangeRate *ForeignExchangeTerms19 `xml:"IssrDclrdXchgRate,omitempty"`

	// Overall tax withheld at source by fund managers prior to considering the tax obligation of each unit holder.
	TaxOnIncome *RateAndAmountFormat46Choice `xml:"TaxOnIncm,omitempty"`
}

func (c *CorporateActionRate79) AddAdditionalTax() *RateAndAmountFormat46Choice {
	c.AdditionalTax = new(RateAndAmountFormat46Choice)
	return c.AdditionalTax
}

func (c *CorporateActionRate79) AddGrossDividendRate() *GrossDividendRateFormat25Choice {
	newValue := new(GrossDividendRateFormat25Choice)
	c.GrossDividendRate = append(c.GrossDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate79) AddNetDividendRate() *NetDividendRateFormat27Choice {
	newValue := new(NetDividendRateFormat27Choice)
	c.NetDividendRate = append(c.NetDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate79) AddIndexFactor() *RateAndAmountFormat46Choice {
	c.IndexFactor = new(RateAndAmountFormat46Choice)
	return c.IndexFactor
}

func (c *CorporateActionRate79) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat10Choice {
	newValue := new(InterestRateUsedForPaymentFormat10Choice)
	c.InterestRateUsedForPayment = append(c.InterestRateUsedForPayment, newValue)
	return newValue
}

func (c *CorporateActionRate79) AddMaximumAllowedOversubscriptionRate() *RateFormat3Choice {
	c.MaximumAllowedOversubscriptionRate = new(RateFormat3Choice)
	return c.MaximumAllowedOversubscriptionRate
}

func (c *CorporateActionRate79) AddProrationRate() *RateFormat3Choice {
	c.ProrationRate = new(RateFormat3Choice)
	return c.ProrationRate
}

func (c *CorporateActionRate79) AddWithholdingTaxRate() *RateAndAmountFormat47Choice {
	newValue := new(RateAndAmountFormat47Choice)
	c.WithholdingTaxRate = append(c.WithholdingTaxRate, newValue)
	return newValue
}

func (c *CorporateActionRate79) AddSecondLevelTax() *RateAndAmountFormat47Choice {
	newValue := new(RateAndAmountFormat47Choice)
	c.SecondLevelTax = append(c.SecondLevelTax, newValue)
	return newValue
}

func (c *CorporateActionRate79) AddTaxableIncomePerDividendShare() *RateTypeAndAmountAndStatus33 {
	newValue := new(RateTypeAndAmountAndStatus33)
	c.TaxableIncomePerDividendShare = append(c.TaxableIncomePerDividendShare, newValue)
	return newValue
}

func (c *CorporateActionRate79) AddIssuerDeclaredExchangeRate() *ForeignExchangeTerms19 {
	c.IssuerDeclaredExchangeRate = new(ForeignExchangeTerms19)
	return c.IssuerDeclaredExchangeRate
}

func (c *CorporateActionRate79) AddTaxOnIncome() *RateAndAmountFormat46Choice {
	c.TaxOnIncome = new(RateAndAmountFormat46Choice)
	return c.TaxOnIncome
}

// Specifies rates related to a corporate action option.
type CorporateActionRate8 struct {

	// Rate proposed in a remarketing of variable rate notes.
	ProposedRate *PercentageRate `xml:"PropsdRate,omitempty"`

	// Rate of allowed over-subscription.
	OversubscriptionRate *RateAndAmountFormat12Choice `xml:"OvrsbcptRate,omitempty"`

	// Requested tax rate in case of breakdown of tax rate, for example, used for adjustment of tax rate. This is the new requested applicable rate.
	RequestedTaxationRate *PercentageRate `xml:"ReqdTaxtnRate,omitempty"`
}

func (c *CorporateActionRate8) SetProposedRate(value string) {
	c.ProposedRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate8) AddOversubscriptionRate() *RateAndAmountFormat12Choice {
	c.OversubscriptionRate = new(RateAndAmountFormat12Choice)
	return c.OversubscriptionRate
}

func (c *CorporateActionRate8) SetRequestedTaxationRate(value string) {
	c.RequestedTaxationRate = (*PercentageRate)(&value)
}

// Specifies rates related to a corporate action option.
type CorporateActionRate80 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat37Choice `xml:"AddtlTax,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat19Choice `xml:"GrssDvddRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat21Choice `xml:"NetDvddRate,omitempty"`

	// Public index rate applied to the amount paid to adjust it to inflation.
	IndexFactor *RateAndAmountFormat37Choice `xml:"IndxFctr,omitempty"`

	// The actual interest rate used for the payment of the interest for the specified interest period.
	// Usage guideline: It is used to provide the applicable rate for the current payment, after all calculations have been performed, that is, application of period and method of interest computation.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat8Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Maximum percentage of shares available through the over subscription privilege, usually a percentage of the basic subscription shares, for example, an account owner subscribing to 100 shares may over subscribe to a maximum of 50 additional shares when the over subscription maximum is 50 percent.
	MaximumAllowedOversubscriptionRate *RateFormat3Choice `xml:"MaxAllwdOvrsbcptRate,omitempty"`

	// Percentage of securities accepted by the offeror/issuer.
	ProrationRate *RateFormat3Choice `xml:"PrratnRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxRate []*RateAndAmountFormat41Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate at which the income will be withheld by a jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate (TAXR) levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTax []*RateAndAmountFormat41Choice `xml:"ScndLvlTax,omitempty"`

	// Amount included in the dividend/NAV that is identified as gains directly or indirectly derived from interest payments, for example, in the context of the EU Savings directive.
	TaxableIncomePerDividendShare []*RateTypeAndAmountAndStatus26 `xml:"TaxblIncmPerDvddShr,omitempty"`

	// Exchange rate (provided by the issuer) between the dividend or interest rate in the paid currency and the declared dividend or interest rate.
	IssuerDeclaredExchangeRate *ForeignExchangeTerms19 `xml:"IssrDclrdXchgRate,omitempty"`

	// Overall tax withheld at source by fund managers prior to considering the tax obligation of each unit holder.
	TaxOnIncome *RateAndAmountFormat37Choice `xml:"TaxOnIncm,omitempty"`
}

func (c *CorporateActionRate80) AddAdditionalTax() *RateAndAmountFormat37Choice {
	c.AdditionalTax = new(RateAndAmountFormat37Choice)
	return c.AdditionalTax
}

func (c *CorporateActionRate80) AddGrossDividendRate() *GrossDividendRateFormat19Choice {
	newValue := new(GrossDividendRateFormat19Choice)
	c.GrossDividendRate = append(c.GrossDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate80) AddNetDividendRate() *NetDividendRateFormat21Choice {
	newValue := new(NetDividendRateFormat21Choice)
	c.NetDividendRate = append(c.NetDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate80) AddIndexFactor() *RateAndAmountFormat37Choice {
	c.IndexFactor = new(RateAndAmountFormat37Choice)
	return c.IndexFactor
}

func (c *CorporateActionRate80) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat8Choice {
	newValue := new(InterestRateUsedForPaymentFormat8Choice)
	c.InterestRateUsedForPayment = append(c.InterestRateUsedForPayment, newValue)
	return newValue
}

func (c *CorporateActionRate80) AddMaximumAllowedOversubscriptionRate() *RateFormat3Choice {
	c.MaximumAllowedOversubscriptionRate = new(RateFormat3Choice)
	return c.MaximumAllowedOversubscriptionRate
}

func (c *CorporateActionRate80) AddProrationRate() *RateFormat3Choice {
	c.ProrationRate = new(RateFormat3Choice)
	return c.ProrationRate
}

func (c *CorporateActionRate80) AddWithholdingTaxRate() *RateAndAmountFormat41Choice {
	newValue := new(RateAndAmountFormat41Choice)
	c.WithholdingTaxRate = append(c.WithholdingTaxRate, newValue)
	return newValue
}

func (c *CorporateActionRate80) AddSecondLevelTax() *RateAndAmountFormat41Choice {
	newValue := new(RateAndAmountFormat41Choice)
	c.SecondLevelTax = append(c.SecondLevelTax, newValue)
	return newValue
}

func (c *CorporateActionRate80) AddTaxableIncomePerDividendShare() *RateTypeAndAmountAndStatus26 {
	newValue := new(RateTypeAndAmountAndStatus26)
	c.TaxableIncomePerDividendShare = append(c.TaxableIncomePerDividendShare, newValue)
	return newValue
}

func (c *CorporateActionRate80) AddIssuerDeclaredExchangeRate() *ForeignExchangeTerms19 {
	c.IssuerDeclaredExchangeRate = new(ForeignExchangeTerms19)
	return c.IssuerDeclaredExchangeRate
}

func (c *CorporateActionRate80) AddTaxOnIncome() *RateAndAmountFormat37Choice {
	c.TaxOnIncome = new(RateAndAmountFormat37Choice)
	return c.TaxOnIncome
}

// Specifies rates related to a corporate action option.
type CorporateActionRate81 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat37Choice `xml:"AddtlTax,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat19Choice `xml:"GrssDvddRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat21Choice `xml:"NetDvddRate,omitempty"`

	// Public index rate applied to the amount paid to adjust it to inflation.
	IndexFactor *RateAndAmountFormat37Choice `xml:"IndxFctr,omitempty"`

	// The actual interest rate used for the payment of the interest for the specified interest period.
	// Usage guideline: It is used to provide the applicable rate for the current payment, after all calculations have been performed, that is, application of period and method of interest computation.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat8Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Maximum percentage of shares available through the over subscription privilege, usually a percentage of the basic subscription shares, for example, an account owner subscribing to 100 shares may over subscribe to a maximum of 50 additional shares when the over subscription maximum is 50 percent.
	MaximumAllowedOversubscriptionRate *RateFormat3Choice `xml:"MaxAllwdOvrsbcptRate,omitempty"`

	// Percentage of securities accepted by the offeror/issuer.
	ProrationRate *RateFormat3Choice `xml:"PrratnRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxRate []*RateAndAmountFormat41Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate at which the income will be withheld by a jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate (TAXR) levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTax []*RateAndAmountFormat41Choice `xml:"ScndLvlTax,omitempty"`

	// Amount included in the dividend/NAV that is identified as gains directly or indirectly derived from interest payments, for example, in the context of the EU Savings directive.
	TaxableIncomePerDividendShare []*RateTypeAndAmountAndStatus26 `xml:"TaxblIncmPerDvddShr,omitempty"`

	// Overall tax withheld at source by fund managers prior to considering the tax obligation of each unit holder.
	TaxOnIncome *RateAndAmountFormat37Choice `xml:"TaxOnIncm,omitempty"`
}

func (c *CorporateActionRate81) AddAdditionalTax() *RateAndAmountFormat37Choice {
	c.AdditionalTax = new(RateAndAmountFormat37Choice)
	return c.AdditionalTax
}

func (c *CorporateActionRate81) AddGrossDividendRate() *GrossDividendRateFormat19Choice {
	newValue := new(GrossDividendRateFormat19Choice)
	c.GrossDividendRate = append(c.GrossDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate81) AddNetDividendRate() *NetDividendRateFormat21Choice {
	newValue := new(NetDividendRateFormat21Choice)
	c.NetDividendRate = append(c.NetDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate81) AddIndexFactor() *RateAndAmountFormat37Choice {
	c.IndexFactor = new(RateAndAmountFormat37Choice)
	return c.IndexFactor
}

func (c *CorporateActionRate81) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat8Choice {
	newValue := new(InterestRateUsedForPaymentFormat8Choice)
	c.InterestRateUsedForPayment = append(c.InterestRateUsedForPayment, newValue)
	return newValue
}

func (c *CorporateActionRate81) AddMaximumAllowedOversubscriptionRate() *RateFormat3Choice {
	c.MaximumAllowedOversubscriptionRate = new(RateFormat3Choice)
	return c.MaximumAllowedOversubscriptionRate
}

func (c *CorporateActionRate81) AddProrationRate() *RateFormat3Choice {
	c.ProrationRate = new(RateFormat3Choice)
	return c.ProrationRate
}

func (c *CorporateActionRate81) AddWithholdingTaxRate() *RateAndAmountFormat41Choice {
	newValue := new(RateAndAmountFormat41Choice)
	c.WithholdingTaxRate = append(c.WithholdingTaxRate, newValue)
	return newValue
}

func (c *CorporateActionRate81) AddSecondLevelTax() *RateAndAmountFormat41Choice {
	newValue := new(RateAndAmountFormat41Choice)
	c.SecondLevelTax = append(c.SecondLevelTax, newValue)
	return newValue
}

func (c *CorporateActionRate81) AddTaxableIncomePerDividendShare() *RateTypeAndAmountAndStatus26 {
	newValue := new(RateTypeAndAmountAndStatus26)
	c.TaxableIncomePerDividendShare = append(c.TaxableIncomePerDividendShare, newValue)
	return newValue
}

func (c *CorporateActionRate81) AddTaxOnIncome() *RateAndAmountFormat37Choice {
	c.TaxOnIncome = new(RateAndAmountFormat37Choice)
	return c.TaxOnIncome
}

// Specifies rates.
type CorporateActionRate82 struct {

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat21Choice `xml:"GrssDvddRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat23Choice `xml:"NetDvddRate,omitempty"`

	// Public index rate applied to the amount paid to adjust it to inflation.
	IndexFactor *RateAndAmountFormat39Choice `xml:"IndxFctr,omitempty"`

	// Actual interest rate used for the payment of the interest for the specified interest period.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat7Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// A maximum percentage of shares available through the over subscription privilege, usually a percentage of the basic subscription shares, for example, an account owner subscribing to 100 shares may over subscribe to a maximum of 50 additional shares when the over subscription maximum is 50 percent.
	MaximumAllowedOversubscriptionRate *PercentageRate `xml:"MaxAllwdOvrsbcptRate,omitempty"`

	// Percentage of securities accepted by the offeror/issuer.
	ProrationRate *PercentageRate `xml:"PrratnRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxRate []*RateAndAmountFormat40Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate at which the income will be withheld by a jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate (TAXR) levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTax []*RateAndAmountFormat40Choice `xml:"ScndLvlTax,omitempty"`

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat39Choice `xml:"AddtlTax,omitempty"`

	// Amount included in the dividend/NAV that is identified as gains directly or indirectly derived from interest payments, for example, in the context of the EU Savings directive.
	TaxableIncomePerDividendShare []*RateTypeAndAmountAndStatus26 `xml:"TaxblIncmPerDvddShr,omitempty"`
}

func (c *CorporateActionRate82) AddGrossDividendRate() *GrossDividendRateFormat21Choice {
	newValue := new(GrossDividendRateFormat21Choice)
	c.GrossDividendRate = append(c.GrossDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate82) AddNetDividendRate() *NetDividendRateFormat23Choice {
	newValue := new(NetDividendRateFormat23Choice)
	c.NetDividendRate = append(c.NetDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate82) AddIndexFactor() *RateAndAmountFormat39Choice {
	c.IndexFactor = new(RateAndAmountFormat39Choice)
	return c.IndexFactor
}

func (c *CorporateActionRate82) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat7Choice {
	newValue := new(InterestRateUsedForPaymentFormat7Choice)
	c.InterestRateUsedForPayment = append(c.InterestRateUsedForPayment, newValue)
	return newValue
}

func (c *CorporateActionRate82) SetMaximumAllowedOversubscriptionRate(value string) {
	c.MaximumAllowedOversubscriptionRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate82) SetProrationRate(value string) {
	c.ProrationRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate82) AddWithholdingTaxRate() *RateAndAmountFormat40Choice {
	newValue := new(RateAndAmountFormat40Choice)
	c.WithholdingTaxRate = append(c.WithholdingTaxRate, newValue)
	return newValue
}

func (c *CorporateActionRate82) AddSecondLevelTax() *RateAndAmountFormat40Choice {
	newValue := new(RateAndAmountFormat40Choice)
	c.SecondLevelTax = append(c.SecondLevelTax, newValue)
	return newValue
}

func (c *CorporateActionRate82) AddAdditionalTax() *RateAndAmountFormat39Choice {
	c.AdditionalTax = new(RateAndAmountFormat39Choice)
	return c.AdditionalTax
}

func (c *CorporateActionRate82) AddTaxableIncomePerDividendShare() *RateTypeAndAmountAndStatus26 {
	newValue := new(RateTypeAndAmountAndStatus26)
	c.TaxableIncomePerDividendShare = append(c.TaxableIncomePerDividendShare, newValue)
	return newValue
}

// Specifies rates related to a corporate action option.
type CorporateActionRate83 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat46Choice `xml:"AddtlTax,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat25Choice `xml:"GrssDvddRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat27Choice `xml:"NetDvddRate,omitempty"`

	// Public index rate applied to the amount paid to adjust it to inflation.
	IndexFactor *RateAndAmountFormat46Choice `xml:"IndxFctr,omitempty"`

	// The actual interest rate used for the payment of the interest for the specified interest period.
	// Usage guideline: It is used to provide the applicable rate for the current payment, after all calculations have been performed, that is, application of period and method of interest computation.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat10Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Maximum percentage of shares available through the over subscription privilege, usually a percentage of the basic subscription shares, for example, an account owner subscribing to 100 shares may over subscribe to a maximum of 50 additional shares when the over subscription maximum is 50 percent.
	MaximumAllowedOversubscriptionRate *RateFormat3Choice `xml:"MaxAllwdOvrsbcptRate,omitempty"`

	// Percentage of securities accepted by the offeror/issuer.
	ProrationRate *RateFormat3Choice `xml:"PrratnRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxRate []*RateAndAmountFormat47Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate at which the income will be withheld by a jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate (TAXR) levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTax []*RateAndAmountFormat47Choice `xml:"ScndLvlTax,omitempty"`

	// Amount included in the dividend/NAV that is identified as gains directly or indirectly derived from interest payments, for example, in the context of the EU Savings directive.
	TaxableIncomePerDividendShare []*RateTypeAndAmountAndStatus33 `xml:"TaxblIncmPerDvddShr,omitempty"`

	// Exchange rate (provided by the issuer) between the dividend or interest rate in the paid currency and the declared dividend or interest rate.
	IssuerDeclaredExchangeRate *ForeignExchangeTerms19 `xml:"IssrDclrdXchgRate,omitempty"`

	// Overall tax withheld at source by fund managers prior to considering the tax obligation of each unit holder.
	TaxOnIncome *RateAndAmountFormat46Choice `xml:"TaxOnIncm,omitempty"`
}

func (c *CorporateActionRate83) AddAdditionalTax() *RateAndAmountFormat46Choice {
	c.AdditionalTax = new(RateAndAmountFormat46Choice)
	return c.AdditionalTax
}

func (c *CorporateActionRate83) AddGrossDividendRate() *GrossDividendRateFormat25Choice {
	newValue := new(GrossDividendRateFormat25Choice)
	c.GrossDividendRate = append(c.GrossDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate83) AddNetDividendRate() *NetDividendRateFormat27Choice {
	newValue := new(NetDividendRateFormat27Choice)
	c.NetDividendRate = append(c.NetDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate83) AddIndexFactor() *RateAndAmountFormat46Choice {
	c.IndexFactor = new(RateAndAmountFormat46Choice)
	return c.IndexFactor
}

func (c *CorporateActionRate83) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat10Choice {
	newValue := new(InterestRateUsedForPaymentFormat10Choice)
	c.InterestRateUsedForPayment = append(c.InterestRateUsedForPayment, newValue)
	return newValue
}

func (c *CorporateActionRate83) AddMaximumAllowedOversubscriptionRate() *RateFormat3Choice {
	c.MaximumAllowedOversubscriptionRate = new(RateFormat3Choice)
	return c.MaximumAllowedOversubscriptionRate
}

func (c *CorporateActionRate83) AddProrationRate() *RateFormat3Choice {
	c.ProrationRate = new(RateFormat3Choice)
	return c.ProrationRate
}

func (c *CorporateActionRate83) AddWithholdingTaxRate() *RateAndAmountFormat47Choice {
	newValue := new(RateAndAmountFormat47Choice)
	c.WithholdingTaxRate = append(c.WithholdingTaxRate, newValue)
	return newValue
}

func (c *CorporateActionRate83) AddSecondLevelTax() *RateAndAmountFormat47Choice {
	newValue := new(RateAndAmountFormat47Choice)
	c.SecondLevelTax = append(c.SecondLevelTax, newValue)
	return newValue
}

func (c *CorporateActionRate83) AddTaxableIncomePerDividendShare() *RateTypeAndAmountAndStatus33 {
	newValue := new(RateTypeAndAmountAndStatus33)
	c.TaxableIncomePerDividendShare = append(c.TaxableIncomePerDividendShare, newValue)
	return newValue
}

func (c *CorporateActionRate83) AddIssuerDeclaredExchangeRate() *ForeignExchangeTerms19 {
	c.IssuerDeclaredExchangeRate = new(ForeignExchangeTerms19)
	return c.IssuerDeclaredExchangeRate
}

func (c *CorporateActionRate83) AddTaxOnIncome() *RateAndAmountFormat46Choice {
	c.TaxOnIncome = new(RateAndAmountFormat46Choice)
	return c.TaxOnIncome
}

// Specifies rates related to a corporate action option.
type CorporateActionRate84 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat46Choice `xml:"AddtlTax,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat25Choice `xml:"GrssDvddRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat27Choice `xml:"NetDvddRate,omitempty"`

	// Public index rate applied to the amount paid to adjust it to inflation.
	IndexFactor *RateAndAmountFormat46Choice `xml:"IndxFctr,omitempty"`

	// The actual interest rate used for the payment of the interest for the specified interest period.
	// Usage guideline: It is used to provide the applicable rate for the current payment, after all calculations have been performed, that is, application of period and method of interest computation.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat10Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Maximum percentage of shares available through the over subscription privilege, usually a percentage of the basic subscription shares, for example, an account owner subscribing to 100 shares may over subscribe to a maximum of 50 additional shares when the over subscription maximum is 50 percent.
	MaximumAllowedOversubscriptionRate *RateFormat3Choice `xml:"MaxAllwdOvrsbcptRate,omitempty"`

	// Percentage of securities accepted by the offeror/issuer.
	ProrationRate *RateFormat3Choice `xml:"PrratnRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxRate []*RateAndAmountFormat47Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate at which the income will be withheld by a jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate (TAXR) levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTax []*RateAndAmountFormat47Choice `xml:"ScndLvlTax,omitempty"`

	// Amount included in the dividend/NAV that is identified as gains directly or indirectly derived from interest payments, for example, in the context of the EU Savings directive.
	TaxableIncomePerDividendShare []*RateTypeAndAmountAndStatus33 `xml:"TaxblIncmPerDvddShr,omitempty"`

	// Overall tax withheld at source by fund managers prior to considering the tax obligation of each unit holder.
	TaxOnIncome *RateAndAmountFormat46Choice `xml:"TaxOnIncm,omitempty"`
}

func (c *CorporateActionRate84) AddAdditionalTax() *RateAndAmountFormat46Choice {
	c.AdditionalTax = new(RateAndAmountFormat46Choice)
	return c.AdditionalTax
}

func (c *CorporateActionRate84) AddGrossDividendRate() *GrossDividendRateFormat25Choice {
	newValue := new(GrossDividendRateFormat25Choice)
	c.GrossDividendRate = append(c.GrossDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate84) AddNetDividendRate() *NetDividendRateFormat27Choice {
	newValue := new(NetDividendRateFormat27Choice)
	c.NetDividendRate = append(c.NetDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate84) AddIndexFactor() *RateAndAmountFormat46Choice {
	c.IndexFactor = new(RateAndAmountFormat46Choice)
	return c.IndexFactor
}

func (c *CorporateActionRate84) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat10Choice {
	newValue := new(InterestRateUsedForPaymentFormat10Choice)
	c.InterestRateUsedForPayment = append(c.InterestRateUsedForPayment, newValue)
	return newValue
}

func (c *CorporateActionRate84) AddMaximumAllowedOversubscriptionRate() *RateFormat3Choice {
	c.MaximumAllowedOversubscriptionRate = new(RateFormat3Choice)
	return c.MaximumAllowedOversubscriptionRate
}

func (c *CorporateActionRate84) AddProrationRate() *RateFormat3Choice {
	c.ProrationRate = new(RateFormat3Choice)
	return c.ProrationRate
}

func (c *CorporateActionRate84) AddWithholdingTaxRate() *RateAndAmountFormat47Choice {
	newValue := new(RateAndAmountFormat47Choice)
	c.WithholdingTaxRate = append(c.WithholdingTaxRate, newValue)
	return newValue
}

func (c *CorporateActionRate84) AddSecondLevelTax() *RateAndAmountFormat47Choice {
	newValue := new(RateAndAmountFormat47Choice)
	c.SecondLevelTax = append(c.SecondLevelTax, newValue)
	return newValue
}

func (c *CorporateActionRate84) AddTaxableIncomePerDividendShare() *RateTypeAndAmountAndStatus33 {
	newValue := new(RateTypeAndAmountAndStatus33)
	c.TaxableIncomePerDividendShare = append(c.TaxableIncomePerDividendShare, newValue)
	return newValue
}

func (c *CorporateActionRate84) AddTaxOnIncome() *RateAndAmountFormat46Choice {
	c.TaxOnIncome = new(RateAndAmountFormat46Choice)
	return c.TaxOnIncome
}

// Specifies rates.
type CorporateActionRate85 struct {

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat23Choice `xml:"GrssDvddRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat25Choice `xml:"NetDvddRate,omitempty"`

	// Public index rate applied to the amount paid to adjust it to inflation.
	IndexFactor *RateAndAmountFormat43Choice `xml:"IndxFctr,omitempty"`

	// Actual interest rate used for the payment of the interest for the specified interest period.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat9Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// A maximum percentage of shares available through the over subscription privilege, usually a percentage of the basic subscription shares, for example, an account owner subscribing to 100 shares may over subscribe to a maximum of 50 additional shares when the over subscription maximum is 50 percent.
	MaximumAllowedOversubscriptionRate *PercentageRate `xml:"MaxAllwdOvrsbcptRate,omitempty"`

	// Percentage of securities accepted by the offeror/issuer.
	ProrationRate *PercentageRate `xml:"PrratnRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxRate []*RateAndAmountFormat45Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate at which the income will be withheld by a jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate (TAXR) levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTax []*RateAndAmountFormat45Choice `xml:"ScndLvlTax,omitempty"`

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat43Choice `xml:"AddtlTax,omitempty"`

	// Amount included in the dividend/NAV that is identified as gains directly or indirectly derived from interest payments, for example, in the context of the EU Savings directive.
	TaxableIncomePerDividendShare []*RateTypeAndAmountAndStatus33 `xml:"TaxblIncmPerDvddShr,omitempty"`
}

func (c *CorporateActionRate85) AddGrossDividendRate() *GrossDividendRateFormat23Choice {
	newValue := new(GrossDividendRateFormat23Choice)
	c.GrossDividendRate = append(c.GrossDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate85) AddNetDividendRate() *NetDividendRateFormat25Choice {
	newValue := new(NetDividendRateFormat25Choice)
	c.NetDividendRate = append(c.NetDividendRate, newValue)
	return newValue
}

func (c *CorporateActionRate85) AddIndexFactor() *RateAndAmountFormat43Choice {
	c.IndexFactor = new(RateAndAmountFormat43Choice)
	return c.IndexFactor
}

func (c *CorporateActionRate85) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat9Choice {
	newValue := new(InterestRateUsedForPaymentFormat9Choice)
	c.InterestRateUsedForPayment = append(c.InterestRateUsedForPayment, newValue)
	return newValue
}

func (c *CorporateActionRate85) SetMaximumAllowedOversubscriptionRate(value string) {
	c.MaximumAllowedOversubscriptionRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate85) SetProrationRate(value string) {
	c.ProrationRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate85) AddWithholdingTaxRate() *RateAndAmountFormat45Choice {
	newValue := new(RateAndAmountFormat45Choice)
	c.WithholdingTaxRate = append(c.WithholdingTaxRate, newValue)
	return newValue
}

func (c *CorporateActionRate85) AddSecondLevelTax() *RateAndAmountFormat45Choice {
	newValue := new(RateAndAmountFormat45Choice)
	c.SecondLevelTax = append(c.SecondLevelTax, newValue)
	return newValue
}

func (c *CorporateActionRate85) AddAdditionalTax() *RateAndAmountFormat43Choice {
	c.AdditionalTax = new(RateAndAmountFormat43Choice)
	return c.AdditionalTax
}

func (c *CorporateActionRate85) AddTaxableIncomePerDividendShare() *RateTypeAndAmountAndStatus33 {
	newValue := new(RateTypeAndAmountAndStatus33)
	c.TaxableIncomePerDividendShare = append(c.TaxableIncomePerDividendShare, newValue)
	return newValue
}
