package iso20022

// Specifies amounts in the framework of a corporate action event..
type CorporateActionAmounts1 struct {

	// Amount of money before any deductions and allowances have been made.
	GrossCashAmount *ActiveCurrencyAndAmount `xml:"GrssCshAmt,omitempty"`

	// Amount of money after deductions and allowances have been made, if any, ie, the total amount +/- charges/fees.
	NetCashAmount *ActiveCurrencyAndAmount `xml:"NetCshAmt,omitempty"`

	// Amount of money defined as a discount on a new issue or on a tranche of an existing issue.
	IssueDiscountAmount *ActiveCurrencyAndAmount `xml:"IsseDscntAmt,omitempty"`

	// Amount of cash premium made available in order to encourage participation in the offer. Payment is made to a third party who has solicited an entity to take part in the offer.
	SolicitationFees *ActiveCurrencyAndAmount `xml:"SlctnFees,omitempty"`

	// Cash disbursement in lieu of a fractional quantity of, eg, equity.
	CashInLieuOfShare *ActiveCurrencyAndAmount `xml:"CshInLieuOfShr,omitempty"`

	// Amount in its original currency when conversion from/into another currency has occurred.
	OriginalAmount *ActiveCurrencyAndAmount `xml:"OrgnlAmt,omitempty"`

	// Amount of money distributed as the result of a capital gain.
	CapitalGain *ActiveCurrencyAndAmount `xml:"CptlGn,omitempty"`

	// Amount of money representing a coupon payment.
	InterestAmount *ActiveCurrencyAndAmount `xml:"IntrstAmt,omitempty"`

	// (Unique to France) Amount due to a buyer of securities dealt prior to ex date which may be subject to different rate of taxation.
	IndemnityAmount *ActiveCurrencyAndAmount `xml:"IndmntyAmt,omitempty"`

	// Amount of money (not interest) in addition to the principal at the redemption of a bond.
	RedemptionPremiumAmount *ActiveCurrencyAndAmount `xml:"RedPrmAmt,omitempty"`

	// Amount of money that the borrower pays to the lender as a compensation. It does not entitle the lender to reclaim any tax credit and is sometimes treated differently by the local tax authorities of the lender.
	ManufacturedDividendAmount *ActiveCurrencyAndAmount `xml:"ManfctrdDvddAmt,omitempty"`

	// Amount of money representing a distribution of a bond's principal, eg, repayment of outstanding debt.
	PrincipalOrCorpus *ActiveCurrencyAndAmount `xml:"PrncplOrCrps,omitempty"`

	// Amount of money reinvested in additional securities.
	ReinvestmentAmount *ActiveCurrencyAndAmount `xml:"RinvstmtAmt,omitempty"`

	// Amount of money resulting from a market claim.
	MarketClaimAmount *ActiveCurrencyAndAmount `xml:"MktClmAmt,omitempty"`

	// Amount resulting from a fully franked dividend paid by a company; amount includes tax credit for companies that have made sufficient tax payments during the fiscal period.
	FullyFrankedAmount *ActiveCurrencyAndAmount `xml:"FullyFrnkdAmt,omitempty"`

	// Amount resulting from an unfranked dividend paid by a company; the amount does not include tax credit and is subject to withholding tax.
	UnfrankedAmount *ActiveCurrencyAndAmount `xml:"UfrnkdAmt,omitempty"`

	// Amount of money related to taxable income that cannot be categorised.
	SundryOrOtherAmount *ActiveCurrencyAndAmount `xml:"SndryOrOthrAmt,omitempty"`

	// Amount of drawdown or other reduction from or in addition to the deal price.
	SpecialConcessionAmount *ActiveCurrencyAndAmount `xml:"SpclCncssnAmt,omitempty"`

	// Cash amount based on terms of corporate action event and balance of underlying securities, entitled to/from account owner (which may be positive or negative).
	EntitledAmount *ActiveCurrencyAndAmount `xml:"EntitldAmt,omitempty"`

	// Rate of the cash premium made available if the securities holder consents or participates to an event, e.g. consent fees.
	CashIncentive *ActiveCurrencyAndAmount `xml:"CshIncntiv,omitempty"`

	// Additional costs - coming on top of the subscription costs - which the subscriber should pay as per the subscription process. Not to be used for the subscription cost itself.
	AdditionalSubscriptionCost *ActiveCurrencyAndAmount `xml:"AddtlSbcptCost,omitempty"`

	// Amount of money that has not been subject to taxation.
	TaxFreeAmount *ActiveCurrencyAndAmount `xml:"TaxFreeAmt,omitempty"`

	// Amount of income eligible for deferred taxation.
	TaxDeferredAmount *ActiveCurrencyAndAmount `xml:"TaxDfrrdAmt,omitempty"`

	// Local tax (ZAS Anrechnungsbetrag) subject to interest down payment tax (proportion of interest liable for interest down payment tax/interim profit that is not covered by the tax exempt amount).
	GermanLocalTax1Amount *ActiveCurrencyAndAmount `xml:"GrmnLclTax1Amt,omitempty"`

	// Local tax (ZAS Pflichtige Zinsen) interest liable for interest down payment tax (proportion of gross interest per unit/interim profits that is not covered by the credit in the interest pool).
	GermanLocalTax2Amount *ActiveCurrencyAndAmount `xml:"GrmnLclTax2Amt,omitempty"`

	// Local tax (Zinstopf) offset interest per unit against tax exempt amount (variation to offset interest per unit in relation to tax exempt amount).
	GermanLocalTax3Amount *ActiveCurrencyAndAmount `xml:"GrmnLclTax3Amt,omitempty"`

	// Local tax (Ertrag Besitzanteilig) yield liable for interest down payment tax.
	GermanLocalTax4Amount *ActiveCurrencyAndAmount `xml:"GrmnLclTax4Amt,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTaxAmount *ActiveCurrencyAndAmount `xml:"StockXchgTaxAmt,omitempty"`

	// Tax levied on a transfer of ownership of financial instrument.
	TransferTaxAmount *ActiveCurrencyAndAmount `xml:"TrfTaxAmt,omitempty"`

	// Amount of transaction tax.
	TransactionTaxAmount *ActiveCurrencyAndAmount `xml:"TxTaxAmt,omitempty"`

	// Tax on value added.
	ValueAddedTaxAmount *ActiveCurrencyAndAmount `xml:"ValAddedTaxAmt,omitempty"`

	// Total amount of tax withheld at source in conformance with the EU Savings Directive.
	EURetentionTaxAmount *ActiveCurrencyAndAmount `xml:"EURtntnTaxAmt,omitempty"`

	// Amount of tax charged by the jurisdiction in which the financial instrument settles.
	LocalTaxAmount *ActiveCurrencyAndAmount `xml:"LclTaxAmt,omitempty"`

	// Payment levy tax.
	PaymentLevyTaxAmount *ActiveCurrencyAndAmount `xml:"PmtLevyTaxAmt,omitempty"`

	// Amount of country, national or federal tax charged by the jurisdiction in which the account servicer is located.
	CountryNationalFederalTaxAmount *ActiveCurrencyAndAmount `xml:"CtryNtlFdrlTaxAmt,omitempty"`

	// Amount of stamp duty.
	StampDutyAmount *ActiveCurrencyAndAmount `xml:"StmpDtyAmt,omitempty"`

	// Amount that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimAmount *ActiveCurrencyAndAmount `xml:"TaxRclmAmt,omitempty"`

	// Amount of taxes that have been previously paid in relation to the taxable event.
	TaxCreditAmount *ActiveCurrencyAndAmount `xml:"TaxCdtAmt,omitempty"`

	// Amount of money withheld by the jurisdiction in which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgOfFrgnTaxAmt,omitempty"`

	// Amount of money withheld by the jurisdiction in which the account owner is located, for which relief at source and/or reclaim may be possible.
	WithholdingOfLocalTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgOfLclTaxAmt,omitempty"`

	// Amount of additional taxes that cannot be categorised.
	AdditionalTaxAmount *ActiveCurrencyAndAmount `xml:"AddtlTaxAmt,omitempty"`

	// Amount of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgTaxAmt,omitempty"`

	// Amount of fiscal tax to apply.
	FiscalStampAmount *ActiveCurrencyAndAmount `xml:"FsclStmpAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *ActiveCurrencyAndAmount `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of paying/sub-paying agent commission.
	PayingAgentCommissionAmount *ActiveCurrencyAndAmount `xml:"PngAgtComssnAmt,omitempty"`

	// Local broker's commission.
	LocalBrokerCommissionAmount *ActiveCurrencyAndAmount `xml:"LclBrkrComssnAmt,omitempty"`

	// Amount of money paid for delivery by regular post mail.
	PostageFeeAmount *ActiveCurrencyAndAmount `xml:"PstgFeeAmt,omitempty"`

	// Amount of money charged by a regulatory authority, eg, Securities and Exchange fees.
	RegulatoryFeesAmount *ActiveCurrencyAndAmount `xml:"RgltryFeesAmt,omitempty"`

	// Amount of money (including insurance) paid for delivery by carrier.
	ShippingFeesAmount *ActiveCurrencyAndAmount `xml:"ShppgFeesAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesAmount *ActiveCurrencyAndAmount `xml:"ChrgsAmt,omitempty"`
}

func (c *CorporateActionAmounts1) SetGrossCashAmount(value, currency string) {
	c.GrossCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetNetCashAmount(value, currency string) {
	c.NetCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetIssueDiscountAmount(value, currency string) {
	c.IssueDiscountAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetSolicitationFees(value, currency string) {
	c.SolicitationFees = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetCashInLieuOfShare(value, currency string) {
	c.CashInLieuOfShare = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetCapitalGain(value, currency string) {
	c.CapitalGain = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetInterestAmount(value, currency string) {
	c.InterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetIndemnityAmount(value, currency string) {
	c.IndemnityAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetRedemptionPremiumAmount(value, currency string) {
	c.RedemptionPremiumAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetManufacturedDividendAmount(value, currency string) {
	c.ManufacturedDividendAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetPrincipalOrCorpus(value, currency string) {
	c.PrincipalOrCorpus = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetReinvestmentAmount(value, currency string) {
	c.ReinvestmentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetMarketClaimAmount(value, currency string) {
	c.MarketClaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetFullyFrankedAmount(value, currency string) {
	c.FullyFrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetUnfrankedAmount(value, currency string) {
	c.UnfrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetSundryOrOtherAmount(value, currency string) {
	c.SundryOrOtherAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetSpecialConcessionAmount(value, currency string) {
	c.SpecialConcessionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetEntitledAmount(value, currency string) {
	c.EntitledAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetCashIncentive(value, currency string) {
	c.CashIncentive = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetAdditionalSubscriptionCost(value, currency string) {
	c.AdditionalSubscriptionCost = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetTaxFreeAmount(value, currency string) {
	c.TaxFreeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetTaxDeferredAmount(value, currency string) {
	c.TaxDeferredAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetGermanLocalTax1Amount(value, currency string) {
	c.GermanLocalTax1Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetGermanLocalTax2Amount(value, currency string) {
	c.GermanLocalTax2Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetGermanLocalTax3Amount(value, currency string) {
	c.GermanLocalTax3Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetGermanLocalTax4Amount(value, currency string) {
	c.GermanLocalTax4Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetStockExchangeTaxAmount(value, currency string) {
	c.StockExchangeTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetTransferTaxAmount(value, currency string) {
	c.TransferTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetTransactionTaxAmount(value, currency string) {
	c.TransactionTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetValueAddedTaxAmount(value, currency string) {
	c.ValueAddedTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetEURetentionTaxAmount(value, currency string) {
	c.EURetentionTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetLocalTaxAmount(value, currency string) {
	c.LocalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetPaymentLevyTaxAmount(value, currency string) {
	c.PaymentLevyTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetCountryNationalFederalTaxAmount(value, currency string) {
	c.CountryNationalFederalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetStampDutyAmount(value, currency string) {
	c.StampDutyAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetTaxReclaimAmount(value, currency string) {
	c.TaxReclaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetTaxCreditAmount(value, currency string) {
	c.TaxCreditAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetWithholdingOfForeignTaxAmount(value, currency string) {
	c.WithholdingOfForeignTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetWithholdingOfLocalTaxAmount(value, currency string) {
	c.WithholdingOfLocalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetAdditionalTaxAmount(value, currency string) {
	c.AdditionalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetWithholdingTaxAmount(value, currency string) {
	c.WithholdingTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetFiscalStampAmount(value, currency string) {
	c.FiscalStampAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetExecutingBrokerAmount(value, currency string) {
	c.ExecutingBrokerAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetPayingAgentCommissionAmount(value, currency string) {
	c.PayingAgentCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetLocalBrokerCommissionAmount(value, currency string) {
	c.LocalBrokerCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetPostageFeeAmount(value, currency string) {
	c.PostageFeeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetRegulatoryFeesAmount(value, currency string) {
	c.RegulatoryFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetShippingFeesAmount(value, currency string) {
	c.ShippingFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts1) SetChargesAmount(value, currency string) {
	c.ChargesAmount = NewActiveCurrencyAndAmount(value, currency)
}

// Specifies amounts in the framework of a corporate action event.
type CorporateActionAmounts10 struct {

	// Amount of money before any deductions and allowances have been made.
	GrossCashAmount *ActiveCurrencyAndAmount `xml:"GrssCshAmt,omitempty"`

	// Amount of money after deductions and allowances have been made, if any, that is, the total amount +/- charges/fees.
	NetCashAmount *ActiveCurrencyAndAmount `xml:"NetCshAmt,omitempty"`

	// Amount of cash premium made available in order to encourage participation in the offer. Payment is made to a third party who has solicited an entity to take part in the offer.
	SolicitationFees *ActiveCurrencyAndAmount `xml:"SlctnFees,omitempty"`

	// Cash disbursement in lieu of a fractional quantity of, for example, equity.
	CashInLieuOfShare *ActiveCurrencyAndAmount `xml:"CshInLieuOfShr,omitempty"`

	// Amount of money distributed as the result of a capital gain.
	CapitalGain *ActiveCurrencyAndAmount `xml:"CptlGn,omitempty"`

	// Amount of money representing a coupon payment.
	InterestAmount *ActiveCurrencyAndAmount `xml:"IntrstAmt,omitempty"`

	// (Unique to France) Amount due to a buyer of securities dealt prior to ex date which may be subject to different rate of taxation.
	IndemnityAmount *ActiveCurrencyAndAmount `xml:"IndmntyAmt,omitempty"`

	// Amount of money that the borrower pays to the lender as a compensation. It does not entitle the lender to reclaim any tax credit and is sometimes treated differently by the local tax authorities of the lender. Also covers compensation/indemnity of missed dividend concerning early/late settlements if applicable to a market.
	ManufacturedDividendPaymentAmount *ActiveCurrencyAndAmount `xml:"ManfctrdDvddPmtAmt,omitempty"`

	// Amount of money reinvested in additional securities.
	ReinvestmentAmount *ActiveCurrencyAndAmount `xml:"RinvstmtAmt,omitempty"`

	// Amount resulting from a fully franked dividend paid by a company; amount includes tax credit for companies that have made sufficient tax payments during the fiscal period.
	FullyFrankedAmount *ActiveCurrencyAndAmount `xml:"FullyFrnkdAmt,omitempty"`

	// Amount resulting from an unfranked dividend paid by a company; the amount does not include tax credit and is subject to withholding tax.
	UnfrankedAmount *ActiveCurrencyAndAmount `xml:"UfrnkdAmt,omitempty"`

	// Amount of money related to taxable income that cannot be categorised.
	SundryOrOtherAmount *ActiveCurrencyAndAmount `xml:"SndryOrOthrAmt,omitempty"`

	// Cash premium made available if the securities holder consents or participates to an event, for example, consent fees or solicitation fees.
	CashIncentive *ActiveCurrencyAndAmount `xml:"CshIncntiv,omitempty"`

	// Amount of money that has not been subject to taxation.
	TaxFreeAmount *ActiveCurrencyAndAmount `xml:"TaxFreeAmt,omitempty"`

	// Amount of income eligible for deferred taxation.
	TaxDeferredAmount *ActiveCurrencyAndAmount `xml:"TaxDfrrdAmt,omitempty"`

	// Amount of value added tax.
	ValueAddedTaxAmount *ActiveCurrencyAndAmount `xml:"ValAddedTaxAmt,omitempty"`

	// Amount of stamp duty.
	StampDutyAmount *ActiveCurrencyAndAmount `xml:"StmpDtyAmt,omitempty"`

	// Amount that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimAmount *ActiveCurrencyAndAmount `xml:"TaxRclmAmt,omitempty"`

	// Amount of taxes that have been previously paid in relation to the taxable event.
	TaxCreditAmount *ActiveCurrencyAndAmount `xml:"TaxCdtAmt,omitempty"`

	// Amount of money withheld by the jurisdiction in which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgOfFrgnTaxAmt,omitempty"`

	// Amount of money withheld by the jurisdiction in which the account owner is located, for which relief at source and/or reclaim may be possible.
	WithholdingOfLocalTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgOfLclTaxAmt,omitempty"`

	// Amount of additional taxes that cannot be categorised.
	AdditionalTaxAmount *ActiveCurrencyAndAmount `xml:"AddtlTaxAmt,omitempty"`

	// Amount of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgTaxAmt,omitempty"`

	// Amount of fiscal tax to apply.
	FiscalStampAmount *ActiveCurrencyAndAmount `xml:"FsclStmpAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *ActiveCurrencyAndAmount `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of paying/sub-paying agent commission.
	PayingAgentCommissionAmount *ActiveCurrencyAndAmount `xml:"PngAgtComssnAmt,omitempty"`

	// Local broker's commission.
	LocalBrokerCommissionAmount *ActiveCurrencyAndAmount `xml:"LclBrkrComssnAmt,omitempty"`

	// Amount of money paid for delivery by regular post mail.
	PostageFeeAmount *ActiveCurrencyAndAmount `xml:"PstgFeeAmt,omitempty"`

	// Amount of money charged by a regulatory authority, for example, securities and exchange fees.
	RegulatoryFeesAmount *ActiveCurrencyAndAmount `xml:"RgltryFeesAmt,omitempty"`

	// Amount of money (including insurance) paid for delivery by carrier.
	ShippingFeesAmount *ActiveCurrencyAndAmount `xml:"ShppgFeesAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesAmount *ActiveCurrencyAndAmount `xml:"ChrgsAmt,omitempty"`

	// Cash amount based on terms of corporate action event and balance of underlying securities, entitled to/from account owner (which may be positive or negative).
	EntitledAmount *ActiveCurrencyAndAmount `xml:"EntitldAmt,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalAmount *ActiveCurrencyAndAmount `xml:"OrgnlAmt,omitempty"`

	// Amount of money representing a distribution of a bond's principal, for example, repayment of outstanding debt.
	PrincipalOrCorpus *ActiveCurrencyAndAmount `xml:"PrncplOrCrps,omitempty"`

	// Amount of money (not interest) in addition to the principal at the redemption of a bond.
	RedemptionPremiumAmount *ActiveCurrencyAndAmount `xml:"RedPrmAmt,omitempty"`

	// Amount relating to the underlying security for which income is distributed.
	IncomePortion *ActiveCurrencyAndAmount `xml:"IncmPrtn,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *ActiveCurrencyAndAmount `xml:"StockXchgTax,omitempty"`

	// Total amount of tax withheld at source in conformance with the EU Savings Directive.
	EUTaxRetentionAmount *ActiveCurrencyAndAmount `xml:"EUTaxRtntnAmt,omitempty"`

	// Amount of interest that has been accrued in between coupon payment periods
	AccruedInterestAmount *ActiveCurrencyAndAmount `xml:"AcrdIntrstAmt,omitempty"`
}

func (c *CorporateActionAmounts10) SetGrossCashAmount(value, currency string) {
	c.GrossCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetNetCashAmount(value, currency string) {
	c.NetCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetSolicitationFees(value, currency string) {
	c.SolicitationFees = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetCashInLieuOfShare(value, currency string) {
	c.CashInLieuOfShare = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetCapitalGain(value, currency string) {
	c.CapitalGain = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetInterestAmount(value, currency string) {
	c.InterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetIndemnityAmount(value, currency string) {
	c.IndemnityAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetManufacturedDividendPaymentAmount(value, currency string) {
	c.ManufacturedDividendPaymentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetReinvestmentAmount(value, currency string) {
	c.ReinvestmentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetFullyFrankedAmount(value, currency string) {
	c.FullyFrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetUnfrankedAmount(value, currency string) {
	c.UnfrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetSundryOrOtherAmount(value, currency string) {
	c.SundryOrOtherAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetCashIncentive(value, currency string) {
	c.CashIncentive = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetTaxFreeAmount(value, currency string) {
	c.TaxFreeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetTaxDeferredAmount(value, currency string) {
	c.TaxDeferredAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetValueAddedTaxAmount(value, currency string) {
	c.ValueAddedTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetStampDutyAmount(value, currency string) {
	c.StampDutyAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetTaxReclaimAmount(value, currency string) {
	c.TaxReclaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetTaxCreditAmount(value, currency string) {
	c.TaxCreditAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetWithholdingOfForeignTaxAmount(value, currency string) {
	c.WithholdingOfForeignTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetWithholdingOfLocalTaxAmount(value, currency string) {
	c.WithholdingOfLocalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetAdditionalTaxAmount(value, currency string) {
	c.AdditionalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetWithholdingTaxAmount(value, currency string) {
	c.WithholdingTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetFiscalStampAmount(value, currency string) {
	c.FiscalStampAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetExecutingBrokerAmount(value, currency string) {
	c.ExecutingBrokerAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetPayingAgentCommissionAmount(value, currency string) {
	c.PayingAgentCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetLocalBrokerCommissionAmount(value, currency string) {
	c.LocalBrokerCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetPostageFeeAmount(value, currency string) {
	c.PostageFeeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetRegulatoryFeesAmount(value, currency string) {
	c.RegulatoryFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetShippingFeesAmount(value, currency string) {
	c.ShippingFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetChargesAmount(value, currency string) {
	c.ChargesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetEntitledAmount(value, currency string) {
	c.EntitledAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetPrincipalOrCorpus(value, currency string) {
	c.PrincipalOrCorpus = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetRedemptionPremiumAmount(value, currency string) {
	c.RedemptionPremiumAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetIncomePortion(value, currency string) {
	c.IncomePortion = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetStockExchangeTax(value, currency string) {
	c.StockExchangeTax = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetEUTaxRetentionAmount(value, currency string) {
	c.EUTaxRetentionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts10) SetAccruedInterestAmount(value, currency string) {
	c.AccruedInterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

// Specifies amounts in the framework of a corporate action event.
type CorporateActionAmounts11 struct {

	// Amount of money that is to be/was posted to the account.
	PostingAmount *ActiveCurrencyAndAmount `xml:"PstngAmt"`

	// Amount of money before any deductions and allowances have been made.
	GrossCashAmount *ActiveCurrencyAndAmount `xml:"GrssCshAmt,omitempty"`

	// Amount of money after deductions and allowances have been made, if any, that is, the total amount +/- charges/fees.
	NetCashAmount *ActiveCurrencyAndAmount `xml:"NetCshAmt,omitempty"`

	// Amount of cash premium made available in order to encourage participation in the offer. Payment is made to a third party who has solicited an entity to take part in the offer.
	SolicitationFees *ActiveCurrencyAndAmount `xml:"SlctnFees,omitempty"`

	// Cash disbursement in lieu of a fractional quantity of, for example, equity.
	CashInLieuOfShare *ActiveCurrencyAndAmount `xml:"CshInLieuOfShr,omitempty"`

	// Amount of money distributed as the result of a capital gain.
	CapitalGain *ActiveCurrencyAndAmount `xml:"CptlGn,omitempty"`

	// Amount of money representing a coupon payment.
	InterestAmount *ActiveCurrencyAndAmount `xml:"IntrstAmt,omitempty"`

	// Amount of money resulting from a market claim.
	MarketClaimAmount *ActiveCurrencyAndAmount `xml:"MktClmAmt,omitempty"`

	// (Unique to France) Amount due to a buyer of securities dealt prior to ex date which may be subject to different rate of taxation.
	IndemnityAmount *ActiveCurrencyAndAmount `xml:"IndmntyAmt,omitempty"`

	// Amount of money that the borrower pays to the lender as a compensation. It does not entitle the lender to reclaim any tax credit and is sometimes treated differently by the local tax authorities of the lender. Also covers compensation/indemnity of missed dividend concerning early/late settlements if applicable to a market.
	ManufacturedDividendPaymentAmount *ActiveCurrencyAndAmount `xml:"ManfctrdDvddPmtAmt,omitempty"`

	// Amount of money reinvested in additional securities.
	ReinvestmentAmount *ActiveCurrencyAndAmount `xml:"RinvstmtAmt,omitempty"`

	// Amount resulting from a fully franked dividend paid by a company; amount includes tax credit for companies that have made sufficient tax payments during the fiscal period.
	FullyFrankedAmount *ActiveCurrencyAndAmount `xml:"FullyFrnkdAmt,omitempty"`

	// Amount resulting from an unfranked dividend paid by a company; the amount does not include tax credit and is subject to withholding tax.
	UnfrankedAmount *ActiveCurrencyAndAmount `xml:"UfrnkdAmt,omitempty"`

	// Amount of money related to taxable income that cannot be categorised.
	SundryOrOtherAmount *ActiveCurrencyAndAmount `xml:"SndryOrOthrAmt,omitempty"`

	// Cash premium made available if the securities holder consents or participates to an event, for example, consent fees or solicitation fees.
	CashIncentive *ActiveCurrencyAndAmount `xml:"CshIncntiv,omitempty"`

	// Amount of money that has not been subject to taxation.
	TaxFreeAmount *ActiveCurrencyAndAmount `xml:"TaxFreeAmt,omitempty"`

	// Amount of income eligible for deferred taxation.
	TaxDeferredAmount *ActiveCurrencyAndAmount `xml:"TaxDfrrdAmt,omitempty"`

	// Tax on value added.
	ValueAddedTaxAmount *ActiveCurrencyAndAmount `xml:"ValAddedTaxAmt,omitempty"`

	// Amount of stamp duty.
	StampDutyAmount *ActiveCurrencyAndAmount `xml:"StmpDtyAmt,omitempty"`

	// Amount that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimAmount *ActiveCurrencyAndAmount `xml:"TaxRclmAmt,omitempty"`

	// Amount of taxes that have been previously paid in relation to the taxable event.
	TaxCreditAmount *ActiveCurrencyAndAmount `xml:"TaxCdtAmt,omitempty"`

	// Amount of money withheld by the jurisdiction in which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgOfFrgnTaxAmt,omitempty"`

	// Amount of money withheld by the jurisdiction in which the account owner is located, for which relief at source and/or reclaim may be possible.
	WithholdingOfLocalTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgOfLclTaxAmt,omitempty"`

	// Amount of additional taxes that cannot be categorised.
	AdditionalTaxAmount *ActiveCurrencyAndAmount `xml:"AddtlTaxAmt,omitempty"`

	// Amount of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgTaxAmt,omitempty"`

	// Amount of fiscal tax to apply.
	FiscalStampAmount *ActiveCurrencyAndAmount `xml:"FsclStmpAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *ActiveCurrencyAndAmount `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of paying/sub-paying agent commission.
	PayingAgentCommissionAmount *ActiveCurrencyAndAmount `xml:"PngAgtComssnAmt,omitempty"`

	// Local broker's commission.
	LocalBrokerCommissionAmount *ActiveCurrencyAndAmount `xml:"LclBrkrComssnAmt,omitempty"`

	// Amount of money paid for delivery by regular post mail.
	PostageFeeAmount *ActiveCurrencyAndAmount `xml:"PstgFeeAmt,omitempty"`

	// Amount of money charged by a regulatory authority, for example, Securities and Exchange fees.
	RegulatoryFeesAmount *ActiveCurrencyAndAmount `xml:"RgltryFeesAmt,omitempty"`

	// Amount of money (including insurance) paid for delivery by carrier.
	ShippingFeesAmount *ActiveCurrencyAndAmount `xml:"ShppgFeesAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesAmount *ActiveCurrencyAndAmount `xml:"ChrgsAmt,omitempty"`

	// Indicates cash retained from previous dividend.
	CashAmountBroughtForward *ActiveCurrencyAndAmount `xml:"CshAmtBrghtFwd,omitempty"`

	// Indicates the balance carried forward to next dividend.
	CashAmountCarriedForward *ActiveCurrencyAndAmount `xml:"CshAmtCrrdFwd,omitempty"`

	// Amount of cash that would have been payable if the dividend had been taken in the form of cash rather than shares.
	NotionalDividendPayableAmount *ActiveCurrencyAndAmount `xml:"NtnlDvddPyblAmt,omitempty"`

	// Where new securities are issued in lieu of a cash dividend, the notional tax is the tax on the amount of cash that would have been paid.
	NotionalTaxAmount *ActiveCurrencyAndAmount `xml:"NtnlTaxAmt,omitempty"`

	// Amount of money paid by the Tax Authorities in addition to the payment of the tax refund itself.
	TaxArrearsAmount *ActiveCurrencyAndAmount `xml:"TaxArrearsAmt,omitempty"`

	// Posting/settlement Amount in its original currency when conversion from/into another currency has occurred.
	OriginalAmount *ActiveCurrencyAndAmount `xml:"OrgnlAmt,omitempty"`

	// Amount of money representing a distribution of a bond's principal, for example, repayment of outstanding debt.
	PrincipalOrCorpus *ActiveCurrencyAndAmount `xml:"PrncplOrCrps,omitempty"`

	// Amount of money (not interest) in addition to the principal at the redemption of a bond.
	RedemptionPremiumAmount *ActiveCurrencyAndAmount `xml:"RedPrmAmt,omitempty"`

	// Amount relating to the underlying security for which income is distributed.
	IncomePortion *ActiveCurrencyAndAmount `xml:"IncmPrtn,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *ActiveCurrencyAndAmount `xml:"StockXchgTax,omitempty"`

	// Total amount of tax withheld at source in conformance with the EU Savings Directive.
	EUTaxRetentionAmount *ActiveCurrencyAndAmount `xml:"EUTaxRtntnAmt,omitempty"`

	// Amount of interest that has been accrued in between coupon payment periods
	AccruedInterestAmount *ActiveCurrencyAndAmount `xml:"AcrdIntrstAmt,omitempty"`
}

func (c *CorporateActionAmounts11) SetPostingAmount(value, currency string) {
	c.PostingAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetGrossCashAmount(value, currency string) {
	c.GrossCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetNetCashAmount(value, currency string) {
	c.NetCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetSolicitationFees(value, currency string) {
	c.SolicitationFees = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetCashInLieuOfShare(value, currency string) {
	c.CashInLieuOfShare = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetCapitalGain(value, currency string) {
	c.CapitalGain = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetInterestAmount(value, currency string) {
	c.InterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetMarketClaimAmount(value, currency string) {
	c.MarketClaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetIndemnityAmount(value, currency string) {
	c.IndemnityAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetManufacturedDividendPaymentAmount(value, currency string) {
	c.ManufacturedDividendPaymentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetReinvestmentAmount(value, currency string) {
	c.ReinvestmentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetFullyFrankedAmount(value, currency string) {
	c.FullyFrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetUnfrankedAmount(value, currency string) {
	c.UnfrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetSundryOrOtherAmount(value, currency string) {
	c.SundryOrOtherAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetCashIncentive(value, currency string) {
	c.CashIncentive = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetTaxFreeAmount(value, currency string) {
	c.TaxFreeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetTaxDeferredAmount(value, currency string) {
	c.TaxDeferredAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetValueAddedTaxAmount(value, currency string) {
	c.ValueAddedTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetStampDutyAmount(value, currency string) {
	c.StampDutyAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetTaxReclaimAmount(value, currency string) {
	c.TaxReclaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetTaxCreditAmount(value, currency string) {
	c.TaxCreditAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetWithholdingOfForeignTaxAmount(value, currency string) {
	c.WithholdingOfForeignTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetWithholdingOfLocalTaxAmount(value, currency string) {
	c.WithholdingOfLocalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetAdditionalTaxAmount(value, currency string) {
	c.AdditionalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetWithholdingTaxAmount(value, currency string) {
	c.WithholdingTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetFiscalStampAmount(value, currency string) {
	c.FiscalStampAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetExecutingBrokerAmount(value, currency string) {
	c.ExecutingBrokerAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetPayingAgentCommissionAmount(value, currency string) {
	c.PayingAgentCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetLocalBrokerCommissionAmount(value, currency string) {
	c.LocalBrokerCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetPostageFeeAmount(value, currency string) {
	c.PostageFeeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetRegulatoryFeesAmount(value, currency string) {
	c.RegulatoryFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetShippingFeesAmount(value, currency string) {
	c.ShippingFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetChargesAmount(value, currency string) {
	c.ChargesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetCashAmountBroughtForward(value, currency string) {
	c.CashAmountBroughtForward = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetCashAmountCarriedForward(value, currency string) {
	c.CashAmountCarriedForward = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetNotionalDividendPayableAmount(value, currency string) {
	c.NotionalDividendPayableAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetNotionalTaxAmount(value, currency string) {
	c.NotionalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetTaxArrearsAmount(value, currency string) {
	c.TaxArrearsAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetPrincipalOrCorpus(value, currency string) {
	c.PrincipalOrCorpus = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetRedemptionPremiumAmount(value, currency string) {
	c.RedemptionPremiumAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetIncomePortion(value, currency string) {
	c.IncomePortion = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetStockExchangeTax(value, currency string) {
	c.StockExchangeTax = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetEUTaxRetentionAmount(value, currency string) {
	c.EUTaxRetentionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts11) SetAccruedInterestAmount(value, currency string) {
	c.AccruedInterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

// Specifies amounts in the framework of a corporate action event.
type CorporateActionAmounts15 struct {

	// Amount of money before any deductions and allowances have been made.
	GrossCashAmount *ActiveCurrencyAndAmount `xml:"GrssCshAmt,omitempty"`

	// Amount of money after deductions and allowances have been made, if any, that is, the total amount +/- charges/fees.
	NetCashAmount *ActiveCurrencyAndAmount `xml:"NetCshAmt,omitempty"`

	// Amount of cash premium made available in order to encourage participation in the offer. Payment is made to a third party who has solicited an entity to take part in the offer.
	SolicitationFees *ActiveCurrencyAndAmount `xml:"SlctnFees,omitempty"`

	// Cash disbursement in lieu of a fractional quantity of, for example, equity.
	CashInLieuOfShare *ActiveCurrencyAndAmount `xml:"CshInLieuOfShr,omitempty"`

	// Amount of money distributed as the result of a capital gain.
	CapitalGain *ActiveCurrencyAndAmount `xml:"CptlGn,omitempty"`

	// Amount of money representing a coupon payment.
	InterestAmount *ActiveCurrencyAndAmount `xml:"IntrstAmt,omitempty"`

	// (Unique to France) Amount due to a buyer of securities dealt prior to ex date which may be subject to different rate of taxation.
	IndemnityAmount *ActiveCurrencyAndAmount `xml:"IndmntyAmt,omitempty"`

	// Amount of money that the borrower pays to the lender as a compensation. It does not entitle the lender to reclaim any tax credit and is sometimes treated differently by the local tax authorities of the lender. Also covers compensation/indemnity of missed dividend concerning early/late settlements if applicable to a market.
	ManufacturedDividendPaymentAmount *ActiveCurrencyAndAmount `xml:"ManfctrdDvddPmtAmt,omitempty"`

	// Amount of money reinvested in additional securities.
	ReinvestmentAmount *ActiveCurrencyAndAmount `xml:"RinvstmtAmt,omitempty"`

	// Amount resulting from a fully franked dividend paid by a company; amount includes tax credit for companies that have made sufficient tax payments during the fiscal period.
	FullyFrankedAmount *ActiveCurrencyAndAmount `xml:"FullyFrnkdAmt,omitempty"`

	// Amount resulting from an unfranked dividend paid by a company; the amount does not include tax credit and is subject to withholding tax.
	UnfrankedAmount *ActiveCurrencyAndAmount `xml:"UfrnkdAmt,omitempty"`

	// Amount of money related to taxable income that cannot be categorised.
	SundryOrOtherAmount *ActiveCurrencyAndAmount `xml:"SndryOrOthrAmt,omitempty"`

	// Cash premium made available if the securities holder consents or participates to an event, for example, consent fees or solicitation fees.
	CashIncentive *ActiveCurrencyAndAmount `xml:"CshIncntiv,omitempty"`

	// Amount of money that has not been subject to taxation.
	TaxFreeAmount *ActiveCurrencyAndAmount `xml:"TaxFreeAmt,omitempty"`

	// Amount of income eligible for deferred taxation.
	TaxDeferredAmount *ActiveCurrencyAndAmount `xml:"TaxDfrrdAmt,omitempty"`

	// Amount of value added tax.
	ValueAddedTaxAmount *ActiveCurrencyAndAmount `xml:"ValAddedTaxAmt,omitempty"`

	// Amount of stamp duty.
	StampDutyAmount *ActiveCurrencyAndAmount `xml:"StmpDtyAmt,omitempty"`

	// Amount that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimAmount *ActiveCurrencyAndAmount `xml:"TaxRclmAmt,omitempty"`

	// Amount of taxes that have been previously paid in relation to the taxable event.
	TaxCreditAmount *ActiveCurrencyAndAmount `xml:"TaxCdtAmt,omitempty"`

	// Amount of money withheld by the jurisdiction in which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgOfFrgnTaxAmt,omitempty"`

	// Amount of money withheld by the jurisdiction in which the account owner is located, for which relief at source and/or reclaim may be possible.
	WithholdingOfLocalTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgOfLclTaxAmt,omitempty"`

	// Amount of additional taxes that cannot be categorised.
	AdditionalTaxAmount *ActiveCurrencyAndAmount `xml:"AddtlTaxAmt,omitempty"`

	// Amount of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgTaxAmt,omitempty"`

	// Amount of fiscal tax to apply.
	FiscalStampAmount *ActiveCurrencyAndAmount `xml:"FsclStmpAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *ActiveCurrencyAndAmount `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of paying/sub-paying agent commission.
	PayingAgentCommissionAmount *ActiveCurrencyAndAmount `xml:"PngAgtComssnAmt,omitempty"`

	// Local broker's commission.
	LocalBrokerCommissionAmount *ActiveCurrencyAndAmount `xml:"LclBrkrComssnAmt,omitempty"`

	// Amount of money charged by a regulatory authority, for example, securities and exchange fees.
	RegulatoryFeesAmount *ActiveCurrencyAndAmount `xml:"RgltryFeesAmt,omitempty"`

	// All costs related to the physical delivery of documents such as stamps, postage, carrier fees, insurances or messenger services.
	ShippingFeesAmount *ActiveCurrencyAndAmount `xml:"ShppgFeesAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesAmount *ActiveCurrencyAndAmount `xml:"ChrgsAmt,omitempty"`

	// Cash amount based on terms of corporate action event and balance of underlying securities, entitled to/from account owner (which may be positive or negative).
	EntitledAmount *ActiveCurrencyAndAmount `xml:"EntitldAmt,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalAmount *ActiveCurrencyAndAmount `xml:"OrgnlAmt,omitempty"`

	// Amount of money representing a distribution of a bond's principal, for example, repayment of outstanding debt.
	PrincipalOrCorpus *ActiveCurrencyAndAmount `xml:"PrncplOrCrps,omitempty"`

	// Amount of money (not interest) in addition to the principal at the redemption of a bond.
	RedemptionPremiumAmount *ActiveCurrencyAndAmount `xml:"RedPrmAmt,omitempty"`

	// Amount relating to the underlying security for which income is distributed.
	IncomePortion *ActiveCurrencyAndAmount `xml:"IncmPrtn,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *ActiveCurrencyAndAmount `xml:"StockXchgTax,omitempty"`

	// Total amount of tax withheld at source in conformance with the EU Savings Directive.
	EUTaxRetentionAmount *ActiveCurrencyAndAmount `xml:"EUTaxRtntnAmt,omitempty"`

	// Amount of interest that has been accrued in between coupon payment periods
	AccruedInterestAmount *ActiveCurrencyAndAmount `xml:"AcrdIntrstAmt,omitempty"`
}

func (c *CorporateActionAmounts15) SetGrossCashAmount(value, currency string) {
	c.GrossCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetNetCashAmount(value, currency string) {
	c.NetCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetSolicitationFees(value, currency string) {
	c.SolicitationFees = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetCashInLieuOfShare(value, currency string) {
	c.CashInLieuOfShare = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetCapitalGain(value, currency string) {
	c.CapitalGain = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetInterestAmount(value, currency string) {
	c.InterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetIndemnityAmount(value, currency string) {
	c.IndemnityAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetManufacturedDividendPaymentAmount(value, currency string) {
	c.ManufacturedDividendPaymentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetReinvestmentAmount(value, currency string) {
	c.ReinvestmentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetFullyFrankedAmount(value, currency string) {
	c.FullyFrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetUnfrankedAmount(value, currency string) {
	c.UnfrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetSundryOrOtherAmount(value, currency string) {
	c.SundryOrOtherAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetCashIncentive(value, currency string) {
	c.CashIncentive = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetTaxFreeAmount(value, currency string) {
	c.TaxFreeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetTaxDeferredAmount(value, currency string) {
	c.TaxDeferredAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetValueAddedTaxAmount(value, currency string) {
	c.ValueAddedTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetStampDutyAmount(value, currency string) {
	c.StampDutyAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetTaxReclaimAmount(value, currency string) {
	c.TaxReclaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetTaxCreditAmount(value, currency string) {
	c.TaxCreditAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetWithholdingOfForeignTaxAmount(value, currency string) {
	c.WithholdingOfForeignTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetWithholdingOfLocalTaxAmount(value, currency string) {
	c.WithholdingOfLocalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetAdditionalTaxAmount(value, currency string) {
	c.AdditionalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetWithholdingTaxAmount(value, currency string) {
	c.WithholdingTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetFiscalStampAmount(value, currency string) {
	c.FiscalStampAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetExecutingBrokerAmount(value, currency string) {
	c.ExecutingBrokerAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetPayingAgentCommissionAmount(value, currency string) {
	c.PayingAgentCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetLocalBrokerCommissionAmount(value, currency string) {
	c.LocalBrokerCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetRegulatoryFeesAmount(value, currency string) {
	c.RegulatoryFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetShippingFeesAmount(value, currency string) {
	c.ShippingFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetChargesAmount(value, currency string) {
	c.ChargesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetEntitledAmount(value, currency string) {
	c.EntitledAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetPrincipalOrCorpus(value, currency string) {
	c.PrincipalOrCorpus = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetRedemptionPremiumAmount(value, currency string) {
	c.RedemptionPremiumAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetIncomePortion(value, currency string) {
	c.IncomePortion = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetStockExchangeTax(value, currency string) {
	c.StockExchangeTax = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetEUTaxRetentionAmount(value, currency string) {
	c.EUTaxRetentionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts15) SetAccruedInterestAmount(value, currency string) {
	c.AccruedInterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

// Specifies amounts in the framework of a corporate action event.
type CorporateActionAmounts16 struct {

	// Amount of money before any deductions and allowances have been made.
	GrossCashAmount *ActiveCurrencyAndAmount `xml:"GrssCshAmt,omitempty"`

	// Amount of money after deductions and allowances have been made, if any, that is, the total amount +/- charges/fees.
	NetCashAmount *ActiveCurrencyAndAmount `xml:"NetCshAmt,omitempty"`

	// Amount of cash premium made available in order to encourage participation in the offer. Payment is made to a third party who has solicited an entity to take part in the offer.
	SolicitationFees *ActiveCurrencyAndAmount `xml:"SlctnFees,omitempty"`

	// Cash disbursement in lieu of a fractional quantity of, for example, equity.
	CashInLieuOfShare *ActiveCurrencyAndAmount `xml:"CshInLieuOfShr,omitempty"`

	// Amount of money distributed as the result of a capital gain.
	CapitalGain *ActiveCurrencyAndAmount `xml:"CptlGn,omitempty"`

	// Amount of money representing a coupon payment.
	InterestAmount *ActiveCurrencyAndAmount `xml:"IntrstAmt,omitempty"`

	// Amount of money resulting from a market claim.
	MarketClaimAmount *ActiveCurrencyAndAmount `xml:"MktClmAmt,omitempty"`

	// (Unique to France) Amount due to a buyer of securities dealt prior to ex date which may be subject to different rate of taxation.
	IndemnityAmount *ActiveCurrencyAndAmount `xml:"IndmntyAmt,omitempty"`

	// Amount of money that the borrower pays to the lender as a compensation. It does not entitle the lender to reclaim any tax credit and is sometimes treated differently by the local tax authorities of the lender. Also covers compensation/indemnity of missed dividend concerning early/late settlements if applicable to a market.
	ManufacturedDividendPaymentAmount *ActiveCurrencyAndAmount `xml:"ManfctrdDvddPmtAmt,omitempty"`

	// Amount of money reinvested in additional securities.
	ReinvestmentAmount *ActiveCurrencyAndAmount `xml:"RinvstmtAmt,omitempty"`

	// Amount resulting from a fully franked dividend paid by a company; amount includes tax credit for companies that have made sufficient tax payments during the fiscal period.
	FullyFrankedAmount *ActiveCurrencyAndAmount `xml:"FullyFrnkdAmt,omitempty"`

	// Amount resulting from an unfranked dividend paid by a company; the amount does not include tax credit and is subject to withholding tax.
	UnfrankedAmount *ActiveCurrencyAndAmount `xml:"UfrnkdAmt,omitempty"`

	// Amount of money related to taxable income that cannot be categorised.
	SundryOrOtherAmount *ActiveCurrencyAndAmount `xml:"SndryOrOthrAmt,omitempty"`

	// Cash premium made available if the securities holder consents or participates to an event, for example, consent fees or solicitation fees.
	CashIncentive *ActiveCurrencyAndAmount `xml:"CshIncntiv,omitempty"`

	// Amount of money that has not been subject to taxation.
	TaxFreeAmount *ActiveCurrencyAndAmount `xml:"TaxFreeAmt,omitempty"`

	// Amount of income eligible for deferred taxation.
	TaxDeferredAmount *ActiveCurrencyAndAmount `xml:"TaxDfrrdAmt,omitempty"`

	// Amount of value added tax.
	ValueAddedTaxAmount *ActiveCurrencyAndAmount `xml:"ValAddedTaxAmt,omitempty"`

	// Amount of stamp duty.
	StampDutyAmount *ActiveCurrencyAndAmount `xml:"StmpDtyAmt,omitempty"`

	// Amount that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimAmount *ActiveCurrencyAndAmount `xml:"TaxRclmAmt,omitempty"`

	// Amount of taxes that have been previously paid in relation to the taxable event.
	TaxCreditAmount *ActiveCurrencyAndAmount `xml:"TaxCdtAmt,omitempty"`

	// Amount of money withheld by the jurisdiction in which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgOfFrgnTaxAmt,omitempty"`

	// Amount of money withheld by the jurisdiction in which the account owner is located, for which relief at source and/or reclaim may be possible.
	WithholdingOfLocalTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgOfLclTaxAmt,omitempty"`

	// Amount of additional taxes that cannot be categorised.
	AdditionalTaxAmount *ActiveCurrencyAndAmount `xml:"AddtlTaxAmt,omitempty"`

	// Amount of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgTaxAmt,omitempty"`

	// Amount of fiscal tax to apply.
	FiscalStampAmount *ActiveCurrencyAndAmount `xml:"FsclStmpAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *ActiveCurrencyAndAmount `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of paying/sub-paying agent commission.
	PayingAgentCommissionAmount *ActiveCurrencyAndAmount `xml:"PngAgtComssnAmt,omitempty"`

	// Local broker's commission.
	LocalBrokerCommissionAmount *ActiveCurrencyAndAmount `xml:"LclBrkrComssnAmt,omitempty"`

	// Amount of money charged by a regulatory authority, for example, securities and exchange fees.
	RegulatoryFeesAmount *ActiveCurrencyAndAmount `xml:"RgltryFeesAmt,omitempty"`

	// All costs related to the physical delivery of documents such as stamps, postage, carrier fees, insurances or messenger services.
	ShippingFeesAmount *ActiveCurrencyAndAmount `xml:"ShppgFeesAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesAmount *ActiveCurrencyAndAmount `xml:"ChrgsAmt,omitempty"`

	// Cash amount based on terms of corporate action event and balance of underlying securities, entitled to/from account owner (which may be positive or negative).
	EntitledAmount *ActiveCurrencyAndAmount `xml:"EntitldAmt,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalAmount *ActiveCurrencyAndAmount `xml:"OrgnlAmt,omitempty"`

	// Amount of interest that has been accrued in between coupon payment periods
	AccruedInterestAmount *ActiveCurrencyAndAmount `xml:"AcrdIntrstAmt,omitempty"`

	// Amount relating to the underlying security for which income is distributed.
	IncomePortion *ActiveCurrencyAndAmount `xml:"IncmPrtn,omitempty"`
}

func (c *CorporateActionAmounts16) SetGrossCashAmount(value, currency string) {
	c.GrossCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts16) SetNetCashAmount(value, currency string) {
	c.NetCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts16) SetSolicitationFees(value, currency string) {
	c.SolicitationFees = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts16) SetCashInLieuOfShare(value, currency string) {
	c.CashInLieuOfShare = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts16) SetCapitalGain(value, currency string) {
	c.CapitalGain = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts16) SetInterestAmount(value, currency string) {
	c.InterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts16) SetMarketClaimAmount(value, currency string) {
	c.MarketClaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts16) SetIndemnityAmount(value, currency string) {
	c.IndemnityAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts16) SetManufacturedDividendPaymentAmount(value, currency string) {
	c.ManufacturedDividendPaymentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts16) SetReinvestmentAmount(value, currency string) {
	c.ReinvestmentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts16) SetFullyFrankedAmount(value, currency string) {
	c.FullyFrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts16) SetUnfrankedAmount(value, currency string) {
	c.UnfrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts16) SetSundryOrOtherAmount(value, currency string) {
	c.SundryOrOtherAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts16) SetCashIncentive(value, currency string) {
	c.CashIncentive = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts16) SetTaxFreeAmount(value, currency string) {
	c.TaxFreeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts16) SetTaxDeferredAmount(value, currency string) {
	c.TaxDeferredAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts16) SetValueAddedTaxAmount(value, currency string) {
	c.ValueAddedTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts16) SetStampDutyAmount(value, currency string) {
	c.StampDutyAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts16) SetTaxReclaimAmount(value, currency string) {
	c.TaxReclaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts16) SetTaxCreditAmount(value, currency string) {
	c.TaxCreditAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts16) SetWithholdingOfForeignTaxAmount(value, currency string) {
	c.WithholdingOfForeignTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts16) SetWithholdingOfLocalTaxAmount(value, currency string) {
	c.WithholdingOfLocalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts16) SetAdditionalTaxAmount(value, currency string) {
	c.AdditionalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts16) SetWithholdingTaxAmount(value, currency string) {
	c.WithholdingTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts16) SetFiscalStampAmount(value, currency string) {
	c.FiscalStampAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts16) SetExecutingBrokerAmount(value, currency string) {
	c.ExecutingBrokerAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts16) SetPayingAgentCommissionAmount(value, currency string) {
	c.PayingAgentCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts16) SetLocalBrokerCommissionAmount(value, currency string) {
	c.LocalBrokerCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts16) SetRegulatoryFeesAmount(value, currency string) {
	c.RegulatoryFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts16) SetShippingFeesAmount(value, currency string) {
	c.ShippingFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts16) SetChargesAmount(value, currency string) {
	c.ChargesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts16) SetEntitledAmount(value, currency string) {
	c.EntitledAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts16) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts16) SetAccruedInterestAmount(value, currency string) {
	c.AccruedInterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts16) SetIncomePortion(value, currency string) {
	c.IncomePortion = NewActiveCurrencyAndAmount(value, currency)
}

// Specifies amounts in the framework of a corporate action event.
type CorporateActionAmounts17 struct {

	// Amount of money that is to be/was posted to the account.
	PostingAmount *ActiveCurrencyAndAmount `xml:"PstngAmt"`

	// Amount of money before any deductions and allowances have been made.
	GrossCashAmount *ActiveCurrencyAndAmount `xml:"GrssCshAmt,omitempty"`

	// Amount of money after deductions and allowances have been made, if any, that is, the total amount +/- charges/fees.
	NetCashAmount *ActiveCurrencyAndAmount `xml:"NetCshAmt,omitempty"`

	// Amount of cash premium made available in order to encourage participation in the offer. Payment is made to a third party who has solicited an entity to take part in the offer.
	SolicitationFees *ActiveCurrencyAndAmount `xml:"SlctnFees,omitempty"`

	// Cash disbursement in lieu of a fractional quantity of, for example, equity.
	CashInLieuOfShare *ActiveCurrencyAndAmount `xml:"CshInLieuOfShr,omitempty"`

	// Amount of money distributed as the result of a capital gain.
	CapitalGain *ActiveCurrencyAndAmount `xml:"CptlGn,omitempty"`

	// Amount of money representing a coupon payment.
	InterestAmount *ActiveCurrencyAndAmount `xml:"IntrstAmt,omitempty"`

	// Amount of money resulting from a market claim.
	MarketClaimAmount *ActiveCurrencyAndAmount `xml:"MktClmAmt,omitempty"`

	// (Unique to France) Amount due to a buyer of securities dealt prior to ex date which may be subject to different rate of taxation.
	IndemnityAmount *ActiveCurrencyAndAmount `xml:"IndmntyAmt,omitempty"`

	// Amount of money that the borrower pays to the lender as a compensation. It does not entitle the lender to reclaim any tax credit and is sometimes treated differently by the local tax authorities of the lender. Also covers compensation/indemnity of missed dividend concerning early/late settlements if applicable to a market.
	ManufacturedDividendPaymentAmount *ActiveCurrencyAndAmount `xml:"ManfctrdDvddPmtAmt,omitempty"`

	// Amount of money reinvested in additional securities.
	ReinvestmentAmount *ActiveCurrencyAndAmount `xml:"RinvstmtAmt,omitempty"`

	// Amount resulting from a fully franked dividend paid by a company; amount includes tax credit for companies that have made sufficient tax payments during the fiscal period.
	FullyFrankedAmount *ActiveCurrencyAndAmount `xml:"FullyFrnkdAmt,omitempty"`

	// Amount resulting from an unfranked dividend paid by a company; the amount does not include tax credit and is subject to withholding tax.
	UnfrankedAmount *ActiveCurrencyAndAmount `xml:"UfrnkdAmt,omitempty"`

	// Amount of money related to taxable income that cannot be categorised.
	SundryOrOtherAmount *ActiveCurrencyAndAmount `xml:"SndryOrOthrAmt,omitempty"`

	// Cash premium made available if the securities holder consents or participates to an event, for example, consent fees or solicitation fees.
	CashIncentive *ActiveCurrencyAndAmount `xml:"CshIncntiv,omitempty"`

	// Amount of money that has not been subject to taxation.
	TaxFreeAmount *ActiveCurrencyAndAmount `xml:"TaxFreeAmt,omitempty"`

	// Amount of income eligible for deferred taxation.
	TaxDeferredAmount *ActiveCurrencyAndAmount `xml:"TaxDfrrdAmt,omitempty"`

	// Tax on value added.
	ValueAddedTaxAmount *ActiveCurrencyAndAmount `xml:"ValAddedTaxAmt,omitempty"`

	// Amount of stamp duty.
	StampDutyAmount *ActiveCurrencyAndAmount `xml:"StmpDtyAmt,omitempty"`

	// Amount that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimAmount *ActiveCurrencyAndAmount `xml:"TaxRclmAmt,omitempty"`

	// Amount of taxes that have been previously paid in relation to the taxable event.
	TaxCreditAmount *ActiveCurrencyAndAmount `xml:"TaxCdtAmt,omitempty"`

	// Amount of money withheld by the jurisdiction in which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgOfFrgnTaxAmt,omitempty"`

	// Amount of money withheld by the jurisdiction in which the account owner is located, for which relief at source and/or reclaim may be possible.
	WithholdingOfLocalTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgOfLclTaxAmt,omitempty"`

	// Amount of additional taxes that cannot be categorised.
	AdditionalTaxAmount *ActiveCurrencyAndAmount `xml:"AddtlTaxAmt,omitempty"`

	// Amount of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgTaxAmt,omitempty"`

	// Amount of fiscal tax to apply.
	FiscalStampAmount *ActiveCurrencyAndAmount `xml:"FsclStmpAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *ActiveCurrencyAndAmount `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of paying/sub-paying agent commission.
	PayingAgentCommissionAmount *ActiveCurrencyAndAmount `xml:"PngAgtComssnAmt,omitempty"`

	// Local broker's commission.
	LocalBrokerCommissionAmount *ActiveCurrencyAndAmount `xml:"LclBrkrComssnAmt,omitempty"`

	// Amount of money charged by a regulatory authority, for example, Securities and Exchange fees.
	RegulatoryFeesAmount *ActiveCurrencyAndAmount `xml:"RgltryFeesAmt,omitempty"`

	// All costs related to the physical delivery of documents such as stamps, postage, carrier fees, insurances or messenger services.
	ShippingFeesAmount *ActiveCurrencyAndAmount `xml:"ShppgFeesAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesAmount *ActiveCurrencyAndAmount `xml:"ChrgsAmt,omitempty"`

	// Indicates cash retained from previous dividend.
	CashAmountBroughtForward *ActiveCurrencyAndAmount `xml:"CshAmtBrghtFwd,omitempty"`

	// Indicates the balance carried forward to next dividend.
	CashAmountCarriedForward *ActiveCurrencyAndAmount `xml:"CshAmtCrrdFwd,omitempty"`

	// Amount of cash that would have been payable if the dividend had been taken in the form of cash rather than shares.
	NotionalDividendPayableAmount *ActiveCurrencyAndAmount `xml:"NtnlDvddPyblAmt,omitempty"`

	// Where new securities are issued in lieu of a cash dividend, the notional tax is the tax on the amount of cash that would have been paid.
	NotionalTaxAmount *ActiveCurrencyAndAmount `xml:"NtnlTaxAmt,omitempty"`

	// Amount of money paid by the Tax Authorities in addition to the payment of the tax refund itself.
	TaxArrearsAmount *ActiveCurrencyAndAmount `xml:"TaxArrearsAmt,omitempty"`

	// Posting/settlement Amount in its original currency when conversion from/into another currency has occurred.
	OriginalAmount *ActiveCurrencyAndAmount `xml:"OrgnlAmt,omitempty"`

	// Amount of money representing a distribution of a bond's principal, for example, repayment of outstanding debt.
	PrincipalOrCorpus *ActiveCurrencyAndAmount `xml:"PrncplOrCrps,omitempty"`

	// Amount of money (not interest) in addition to the principal at the redemption of a bond.
	RedemptionPremiumAmount *ActiveCurrencyAndAmount `xml:"RedPrmAmt,omitempty"`

	// Amount relating to the underlying security for which income is distributed.
	IncomePortion *ActiveCurrencyAndAmount `xml:"IncmPrtn,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *ActiveCurrencyAndAmount `xml:"StockXchgTax,omitempty"`

	// Total amount of tax withheld at source in conformance with the EU Savings Directive.
	EUTaxRetentionAmount *ActiveCurrencyAndAmount `xml:"EUTaxRtntnAmt,omitempty"`

	// Amount of interest that has been accrued in between coupon payment periods
	AccruedInterestAmount *ActiveCurrencyAndAmount `xml:"AcrdIntrstAmt,omitempty"`
}

func (c *CorporateActionAmounts17) SetPostingAmount(value, currency string) {
	c.PostingAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetGrossCashAmount(value, currency string) {
	c.GrossCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetNetCashAmount(value, currency string) {
	c.NetCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetSolicitationFees(value, currency string) {
	c.SolicitationFees = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetCashInLieuOfShare(value, currency string) {
	c.CashInLieuOfShare = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetCapitalGain(value, currency string) {
	c.CapitalGain = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetInterestAmount(value, currency string) {
	c.InterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetMarketClaimAmount(value, currency string) {
	c.MarketClaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetIndemnityAmount(value, currency string) {
	c.IndemnityAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetManufacturedDividendPaymentAmount(value, currency string) {
	c.ManufacturedDividendPaymentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetReinvestmentAmount(value, currency string) {
	c.ReinvestmentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetFullyFrankedAmount(value, currency string) {
	c.FullyFrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetUnfrankedAmount(value, currency string) {
	c.UnfrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetSundryOrOtherAmount(value, currency string) {
	c.SundryOrOtherAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetCashIncentive(value, currency string) {
	c.CashIncentive = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetTaxFreeAmount(value, currency string) {
	c.TaxFreeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetTaxDeferredAmount(value, currency string) {
	c.TaxDeferredAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetValueAddedTaxAmount(value, currency string) {
	c.ValueAddedTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetStampDutyAmount(value, currency string) {
	c.StampDutyAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetTaxReclaimAmount(value, currency string) {
	c.TaxReclaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetTaxCreditAmount(value, currency string) {
	c.TaxCreditAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetWithholdingOfForeignTaxAmount(value, currency string) {
	c.WithholdingOfForeignTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetWithholdingOfLocalTaxAmount(value, currency string) {
	c.WithholdingOfLocalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetAdditionalTaxAmount(value, currency string) {
	c.AdditionalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetWithholdingTaxAmount(value, currency string) {
	c.WithholdingTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetFiscalStampAmount(value, currency string) {
	c.FiscalStampAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetExecutingBrokerAmount(value, currency string) {
	c.ExecutingBrokerAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetPayingAgentCommissionAmount(value, currency string) {
	c.PayingAgentCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetLocalBrokerCommissionAmount(value, currency string) {
	c.LocalBrokerCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetRegulatoryFeesAmount(value, currency string) {
	c.RegulatoryFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetShippingFeesAmount(value, currency string) {
	c.ShippingFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetChargesAmount(value, currency string) {
	c.ChargesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetCashAmountBroughtForward(value, currency string) {
	c.CashAmountBroughtForward = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetCashAmountCarriedForward(value, currency string) {
	c.CashAmountCarriedForward = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetNotionalDividendPayableAmount(value, currency string) {
	c.NotionalDividendPayableAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetNotionalTaxAmount(value, currency string) {
	c.NotionalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetTaxArrearsAmount(value, currency string) {
	c.TaxArrearsAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetPrincipalOrCorpus(value, currency string) {
	c.PrincipalOrCorpus = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetRedemptionPremiumAmount(value, currency string) {
	c.RedemptionPremiumAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetIncomePortion(value, currency string) {
	c.IncomePortion = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetStockExchangeTax(value, currency string) {
	c.StockExchangeTax = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetEUTaxRetentionAmount(value, currency string) {
	c.EUTaxRetentionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts17) SetAccruedInterestAmount(value, currency string) {
	c.AccruedInterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

// Specifies amounts in the framework of a corporate action event.
type CorporateActionAmounts2 struct {

	// Amount of money that is to be/was posted to the account.
	PostingAmount *ActiveCurrencyAndAmount `xml:"PstngAmt"`

	// Amount of money before any deductions and allowances have been made.
	GrossCashAmount *ActiveCurrencyAndAmount `xml:"GrssCshAmt,omitempty"`

	// Amount of money after deductions and allowances have been made, if any, that is, the total amount +/- charges/fees.
	NetCashAmount *ActiveCurrencyAndAmount `xml:"NetCshAmt,omitempty"`

	// Amount of cash premium made available in order to encourage participation in the offer. Payment is made to a third party who has solicited an entity to take part in the offer.
	SolicitationFees *ActiveCurrencyAndAmount `xml:"SlctnFees,omitempty"`

	// Cash disbursement in lieu of a fractional quantity of, for example, equity.
	CashInLieuOfShare *ActiveCurrencyAndAmount `xml:"CshInLieuOfShr,omitempty"`

	// Amount of money distributed as the result of a capital gain.
	CapitalGain *ActiveCurrencyAndAmount `xml:"CptlGn,omitempty"`

	// Amount of money representing a coupon payment.
	InterestAmount *ActiveCurrencyAndAmount `xml:"IntrstAmt,omitempty"`

	// Amount of money resulting from a market claim.
	MarketClaimAmount *ActiveCurrencyAndAmount `xml:"MktClmAmt,omitempty"`

	// (Unique to France) Amount due to a buyer of securities dealt prior to ex date which may be subject to different rate of taxation.
	IndemnityAmount *ActiveCurrencyAndAmount `xml:"IndmntyAmt,omitempty"`

	// Amount of money that the borrower pays to the lender as a compensation. It does not entitle the lender to reclaim any tax credit and is sometimes treated differently by the local tax authorities of the lender.
	ManufacturedDividendAmount *ActiveCurrencyAndAmount `xml:"ManfctrdDvddAmt,omitempty"`

	// Amount of money reinvested in additional securities.
	ReinvestmentAmount *ActiveCurrencyAndAmount `xml:"RinvstmtAmt,omitempty"`

	// Amount resulting from a fully franked dividend paid by a company; amount includes tax credit for companies that have made sufficient tax payments during the fiscal period.
	FullyFrankedAmount *ActiveCurrencyAndAmount `xml:"FullyFrnkdAmt,omitempty"`

	// Amount resulting from an unfranked dividend paid by a company; the amount does not include tax credit and is subject to withholding tax.
	UnfrankedAmount *ActiveCurrencyAndAmount `xml:"UfrnkdAmt,omitempty"`

	// Amount of money related to taxable income that cannot be categorised.
	SundryOrOtherAmount *ActiveCurrencyAndAmount `xml:"SndryOrOthrAmt,omitempty"`

	// Cash premium made available if the securities holder consents or participates to an event, for example, consent fees or solicitation fees.
	CashIncentive *ActiveCurrencyAndAmount `xml:"CshIncntiv,omitempty"`

	// Amount of money that has not been subject to taxation.
	TaxFreeAmount *ActiveCurrencyAndAmount `xml:"TaxFreeAmt,omitempty"`

	// Amount of income eligible for deferred taxation.
	TaxDeferredAmount *ActiveCurrencyAndAmount `xml:"TaxDfrrdAmt,omitempty"`

	// Tax on value added.
	ValueAddedTaxAmount *ActiveCurrencyAndAmount `xml:"ValAddedTaxAmt,omitempty"`

	// Amount of stamp duty.
	StampDutyAmount *ActiveCurrencyAndAmount `xml:"StmpDtyAmt,omitempty"`

	// Amount that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimAmount *ActiveCurrencyAndAmount `xml:"TaxRclmAmt,omitempty"`

	// Amount of taxes that have been previously paid in relation to the taxable event.
	TaxCreditAmount *ActiveCurrencyAndAmount `xml:"TaxCdtAmt,omitempty"`

	// Amount of money withheld by the jurisdiction in which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgOfFrgnTaxAmt,omitempty"`

	// Amount of money withheld by the jurisdiction in which the account owner is located, for which relief at source and/or reclaim may be possible.
	WithholdingOfLocalTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgOfLclTaxAmt,omitempty"`

	// Amount of additional taxes that cannot be categorised.
	AdditionalTaxAmount *ActiveCurrencyAndAmount `xml:"AddtlTaxAmt,omitempty"`

	// Amount of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgTaxAmt,omitempty"`

	// Amount of fiscal tax to apply.
	FiscalStampAmount *ActiveCurrencyAndAmount `xml:"FsclStmpAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *ActiveCurrencyAndAmount `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of paying/sub-paying agent commission.
	PayingAgentCommissionAmount *ActiveCurrencyAndAmount `xml:"PngAgtComssnAmt,omitempty"`

	// Local broker's commission.
	LocalBrokerCommissionAmount *ActiveCurrencyAndAmount `xml:"LclBrkrComssnAmt,omitempty"`

	// Amount of money paid for delivery by regular post mail.
	PostageFeeAmount *ActiveCurrencyAndAmount `xml:"PstgFeeAmt,omitempty"`

	// Amount of money charged by a regulatory authority, for example, Securities and Exchange fees.
	RegulatoryFeesAmount *ActiveCurrencyAndAmount `xml:"RgltryFeesAmt,omitempty"`

	// Amount of money (including insurance) paid for delivery by carrier.
	ShippingFeesAmount *ActiveCurrencyAndAmount `xml:"ShppgFeesAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesAmount *ActiveCurrencyAndAmount `xml:"ChrgsAmt,omitempty"`

	// Indicates cash retained from previous dividend.
	CashAmountBroughtForward *ActiveCurrencyAndAmount `xml:"CshAmtBrghtFwd,omitempty"`

	// Indicates the balance carried forward to next dividend.
	CashAmountCarriedForward *ActiveCurrencyAndAmount `xml:"CshAmtCrrdFwd,omitempty"`

	// Amount of cash that would have been payable if the dividend had been taken in the form of cash rather than shares.
	NotionalDividendPayableAmount *ActiveCurrencyAndAmount `xml:"NtnlDvddPyblAmt,omitempty"`

	// Where new securities are issued in lieu of a cash dividend, the notional tax is the tax on the amount of cash that would have been paid.
	NotionalTaxAmount *ActiveCurrencyAndAmount `xml:"NtnlTaxAmt,omitempty"`

	// Amount of money paid by the Tax Authorities in addition to the payment of the tax refund itself.
	TaxArrearsAmount *ActiveCurrencyAndAmount `xml:"TaxArrearsAmt,omitempty"`

	// Posting/settlement Amount in its original currency when conversion from/into another currency has occurred.
	OriginalAmount *ActiveCurrencyAndAmount `xml:"OrgnlAmt,omitempty"`

	// Amount of money representing a distribution of a bond's principal, for example, repayment of outstanding debt.
	PrincipalOrCorpus *ActiveCurrencyAndAmount `xml:"PrncplOrCrps,omitempty"`

	// Amount of money (not interest) in addition to the principal at the redemption of a bond.
	RedemptionPremiumAmount *ActiveCurrencyAndAmount `xml:"RedPrmAmt,omitempty"`

	// Amount relating to the underlying security for which income is distributed.
	IncomePortion *ActiveCurrencyAndAmount `xml:"IncmPrtn,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *ActiveCurrencyAndAmount `xml:"StockXchgTax,omitempty"`

	// Total amount of tax withheld at source in conformance with the EU Savings Directive.
	EUTaxRetentionAmount *ActiveCurrencyAndAmount `xml:"EUTaxRtntnAmt,omitempty"`

	// Amount of interest that has been accrued in between coupon payment periods
	AccruedInterestAmount *ActiveCurrencyAndAmount `xml:"AcrdIntrstAmt,omitempty"`
}

func (c *CorporateActionAmounts2) SetPostingAmount(value, currency string) {
	c.PostingAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetGrossCashAmount(value, currency string) {
	c.GrossCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetNetCashAmount(value, currency string) {
	c.NetCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetSolicitationFees(value, currency string) {
	c.SolicitationFees = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetCashInLieuOfShare(value, currency string) {
	c.CashInLieuOfShare = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetCapitalGain(value, currency string) {
	c.CapitalGain = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetInterestAmount(value, currency string) {
	c.InterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetMarketClaimAmount(value, currency string) {
	c.MarketClaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetIndemnityAmount(value, currency string) {
	c.IndemnityAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetManufacturedDividendAmount(value, currency string) {
	c.ManufacturedDividendAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetReinvestmentAmount(value, currency string) {
	c.ReinvestmentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetFullyFrankedAmount(value, currency string) {
	c.FullyFrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetUnfrankedAmount(value, currency string) {
	c.UnfrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetSundryOrOtherAmount(value, currency string) {
	c.SundryOrOtherAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetCashIncentive(value, currency string) {
	c.CashIncentive = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetTaxFreeAmount(value, currency string) {
	c.TaxFreeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetTaxDeferredAmount(value, currency string) {
	c.TaxDeferredAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetValueAddedTaxAmount(value, currency string) {
	c.ValueAddedTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetStampDutyAmount(value, currency string) {
	c.StampDutyAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetTaxReclaimAmount(value, currency string) {
	c.TaxReclaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetTaxCreditAmount(value, currency string) {
	c.TaxCreditAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetWithholdingOfForeignTaxAmount(value, currency string) {
	c.WithholdingOfForeignTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetWithholdingOfLocalTaxAmount(value, currency string) {
	c.WithholdingOfLocalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetAdditionalTaxAmount(value, currency string) {
	c.AdditionalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetWithholdingTaxAmount(value, currency string) {
	c.WithholdingTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetFiscalStampAmount(value, currency string) {
	c.FiscalStampAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetExecutingBrokerAmount(value, currency string) {
	c.ExecutingBrokerAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetPayingAgentCommissionAmount(value, currency string) {
	c.PayingAgentCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetLocalBrokerCommissionAmount(value, currency string) {
	c.LocalBrokerCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetPostageFeeAmount(value, currency string) {
	c.PostageFeeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetRegulatoryFeesAmount(value, currency string) {
	c.RegulatoryFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetShippingFeesAmount(value, currency string) {
	c.ShippingFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetChargesAmount(value, currency string) {
	c.ChargesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetCashAmountBroughtForward(value, currency string) {
	c.CashAmountBroughtForward = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetCashAmountCarriedForward(value, currency string) {
	c.CashAmountCarriedForward = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetNotionalDividendPayableAmount(value, currency string) {
	c.NotionalDividendPayableAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetNotionalTaxAmount(value, currency string) {
	c.NotionalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetTaxArrearsAmount(value, currency string) {
	c.TaxArrearsAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetPrincipalOrCorpus(value, currency string) {
	c.PrincipalOrCorpus = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetRedemptionPremiumAmount(value, currency string) {
	c.RedemptionPremiumAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetIncomePortion(value, currency string) {
	c.IncomePortion = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetStockExchangeTax(value, currency string) {
	c.StockExchangeTax = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetEUTaxRetentionAmount(value, currency string) {
	c.EUTaxRetentionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts2) SetAccruedInterestAmount(value, currency string) {
	c.AccruedInterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

// Specifies amounts in the framework of a corporate action event.
type CorporateActionAmounts21 struct {

	// Amount of money before any deductions and allowances have been made.
	GrossCashAmount *ActiveCurrencyAndAmount `xml:"GrssCshAmt,omitempty"`

	// Amount of money after deductions and allowances have been made, if any, that is, the total amount +/- charges/fees.
	NetCashAmount *ActiveCurrencyAndAmount `xml:"NetCshAmt,omitempty"`

	// Cash premium made available if the securities holder consents or participates to an event, for example consent fees or solicitation fees.
	SolicitationFees *ActiveCurrencyAndAmount `xml:"SlctnFees,omitempty"`

	// Cash disbursement in lieu of a fractional quantity of, for example, equity.
	CashInLieuOfShare *ActiveCurrencyAndAmount `xml:"CshInLieuOfShr,omitempty"`

	// Amount of money distributed as the result of a capital gain.
	CapitalGain *ActiveCurrencyAndAmount `xml:"CptlGn,omitempty"`

	// Amount of money representing a coupon payment.
	InterestAmount *ActiveCurrencyAndAmount `xml:"IntrstAmt,omitempty"`

	// (Unique to France) Amount due to a buyer of securities dealt prior to ex date which may be subject to different rate of taxation.
	IndemnityAmount *ActiveCurrencyAndAmount `xml:"IndmntyAmt,omitempty"`

	// Amount of money that the borrower pays to the lender as a compensation. It does not entitle the lender to reclaim any tax credit and is sometimes treated differently by the local tax authorities of the lender. Also covers compensation/indemnity of missed dividend concerning early/late settlements if applicable to a market.
	ManufacturedDividendPaymentAmount *ActiveCurrencyAndAmount `xml:"ManfctrdDvddPmtAmt,omitempty"`

	// Amount of money reinvested in additional securities.
	ReinvestmentAmount *ActiveCurrencyAndAmount `xml:"RinvstmtAmt,omitempty"`

	// Amount resulting from a fully franked dividend paid by a company; amount includes tax credit for companies that have made sufficient tax payments during the fiscal period.
	FullyFrankedAmount *ActiveCurrencyAndAmount `xml:"FullyFrnkdAmt,omitempty"`

	// Amount resulting from an unfranked dividend paid by a company; the amount does not include tax credit and is subject to withholding tax.
	UnfrankedAmount *ActiveCurrencyAndAmount `xml:"UfrnkdAmt,omitempty"`

	// Amount of money related to taxable income that cannot be categorised.
	SundryOrOtherAmount *ActiveCurrencyAndAmount `xml:"SndryOrOthrAmt,omitempty"`

	// Amount of money that has not been subject to taxation.
	TaxFreeAmount *ActiveCurrencyAndAmount `xml:"TaxFreeAmt,omitempty"`

	// Amount of income eligible for deferred taxation.
	TaxDeferredAmount *ActiveCurrencyAndAmount `xml:"TaxDfrrdAmt,omitempty"`

	// Amount of value added tax.
	ValueAddedTaxAmount *ActiveCurrencyAndAmount `xml:"ValAddedTaxAmt,omitempty"`

	// Amount of stamp duty.
	StampDutyAmount *ActiveCurrencyAndAmount `xml:"StmpDtyAmt,omitempty"`

	// Amount that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimAmount *ActiveCurrencyAndAmount `xml:"TaxRclmAmt,omitempty"`

	// Amount of taxes that have been previously paid in relation to the taxable event.
	TaxCreditAmount *ActiveCurrencyAndAmount `xml:"TaxCdtAmt,omitempty"`

	// Amount of money withheld by the jurisdiction in which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgOfFrgnTaxAmt,omitempty"`

	// Amount of money withheld by the jurisdiction in which the account owner is located, for which relief at source and/or reclaim may be possible.
	WithholdingOfLocalTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgOfLclTaxAmt,omitempty"`

	// Amount of additional taxes that cannot be categorised.
	AdditionalTaxAmount *ActiveCurrencyAndAmount `xml:"AddtlTaxAmt,omitempty"`

	// Amount of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgTaxAmt,omitempty"`

	// Amount of fiscal tax to apply.
	FiscalStampAmount *ActiveCurrencyAndAmount `xml:"FsclStmpAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *ActiveCurrencyAndAmount `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of paying/sub-paying agent commission.
	PayingAgentCommissionAmount *ActiveCurrencyAndAmount `xml:"PngAgtComssnAmt,omitempty"`

	// Local broker's commission.
	LocalBrokerCommissionAmount *ActiveCurrencyAndAmount `xml:"LclBrkrComssnAmt,omitempty"`

	// Amount of money charged by a regulatory authority, for example, securities and exchange fees.
	RegulatoryFeesAmount *ActiveCurrencyAndAmount `xml:"RgltryFeesAmt,omitempty"`

	// All costs related to the physical delivery of documents such as stamps, postage, carrier fees, insurances or messenger services.
	ShippingFeesAmount *ActiveCurrencyAndAmount `xml:"ShppgFeesAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesAmount *ActiveCurrencyAndAmount `xml:"ChrgsAmt,omitempty"`

	// Cash amount based on terms of corporate action event and balance of underlying securities, entitled to/from account owner (which may be positive or negative).
	EntitledAmount *ActiveCurrencyAndAmount `xml:"EntitldAmt,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalAmount *ActiveCurrencyAndAmount `xml:"OrgnlAmt,omitempty"`

	// Amount of money representing a distribution of a bond's principal, for example, repayment of outstanding debt.
	PrincipalOrCorpus *ActiveCurrencyAndAmount `xml:"PrncplOrCrps,omitempty"`

	// Amount of money (not interest) in addition to the principal at the redemption of a bond.
	RedemptionPremiumAmount *ActiveCurrencyAndAmount `xml:"RedPrmAmt,omitempty"`

	// Amount relating to the underlying security for which income is distributed.
	IncomePortion *ActiveCurrencyAndAmount `xml:"IncmPrtn,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *ActiveCurrencyAndAmount `xml:"StockXchgTax,omitempty"`

	// Total amount of tax withheld at source in conformance with the EU Savings Directive.
	EUTaxRetentionAmount *ActiveCurrencyAndAmount `xml:"EUTaxRtntnAmt,omitempty"`

	// Amount of interest that has been accrued in between coupon payment periods
	AccruedInterestAmount *ActiveCurrencyAndAmount `xml:"AcrdIntrstAmt,omitempty"`

	// Portion of the fund distribution amount which represents the average accrued income included in the purchase price for units bought during the account period.
	EqualisationAmount *ActiveCurrencyAndAmount `xml:"EqulstnAmt,omitempty"`
}

func (c *CorporateActionAmounts21) SetGrossCashAmount(value, currency string) {
	c.GrossCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetNetCashAmount(value, currency string) {
	c.NetCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetSolicitationFees(value, currency string) {
	c.SolicitationFees = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetCashInLieuOfShare(value, currency string) {
	c.CashInLieuOfShare = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetCapitalGain(value, currency string) {
	c.CapitalGain = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetInterestAmount(value, currency string) {
	c.InterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetIndemnityAmount(value, currency string) {
	c.IndemnityAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetManufacturedDividendPaymentAmount(value, currency string) {
	c.ManufacturedDividendPaymentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetReinvestmentAmount(value, currency string) {
	c.ReinvestmentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetFullyFrankedAmount(value, currency string) {
	c.FullyFrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetUnfrankedAmount(value, currency string) {
	c.UnfrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetSundryOrOtherAmount(value, currency string) {
	c.SundryOrOtherAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetTaxFreeAmount(value, currency string) {
	c.TaxFreeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetTaxDeferredAmount(value, currency string) {
	c.TaxDeferredAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetValueAddedTaxAmount(value, currency string) {
	c.ValueAddedTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetStampDutyAmount(value, currency string) {
	c.StampDutyAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetTaxReclaimAmount(value, currency string) {
	c.TaxReclaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetTaxCreditAmount(value, currency string) {
	c.TaxCreditAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetWithholdingOfForeignTaxAmount(value, currency string) {
	c.WithholdingOfForeignTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetWithholdingOfLocalTaxAmount(value, currency string) {
	c.WithholdingOfLocalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetAdditionalTaxAmount(value, currency string) {
	c.AdditionalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetWithholdingTaxAmount(value, currency string) {
	c.WithholdingTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetFiscalStampAmount(value, currency string) {
	c.FiscalStampAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetExecutingBrokerAmount(value, currency string) {
	c.ExecutingBrokerAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetPayingAgentCommissionAmount(value, currency string) {
	c.PayingAgentCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetLocalBrokerCommissionAmount(value, currency string) {
	c.LocalBrokerCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetRegulatoryFeesAmount(value, currency string) {
	c.RegulatoryFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetShippingFeesAmount(value, currency string) {
	c.ShippingFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetChargesAmount(value, currency string) {
	c.ChargesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetEntitledAmount(value, currency string) {
	c.EntitledAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetPrincipalOrCorpus(value, currency string) {
	c.PrincipalOrCorpus = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetRedemptionPremiumAmount(value, currency string) {
	c.RedemptionPremiumAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetIncomePortion(value, currency string) {
	c.IncomePortion = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetStockExchangeTax(value, currency string) {
	c.StockExchangeTax = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetEUTaxRetentionAmount(value, currency string) {
	c.EUTaxRetentionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetAccruedInterestAmount(value, currency string) {
	c.AccruedInterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts21) SetEqualisationAmount(value, currency string) {
	c.EqualisationAmount = NewActiveCurrencyAndAmount(value, currency)
}

// Specifies amounts in the framework of a corporate action event.
type CorporateActionAmounts22 struct {

	// Amount of money before any deductions and allowances have been made.
	GrossCashAmount *ActiveCurrencyAndAmount `xml:"GrssCshAmt,omitempty"`

	// Amount of money after deductions and allowances have been made, if any, that is, the total amount +/- charges/fees.
	NetCashAmount *ActiveCurrencyAndAmount `xml:"NetCshAmt,omitempty"`

	// Cash premium made available if the securities holder consents or participates to an event, for example consent fees or solicitation fees.
	SolicitationFees *ActiveCurrencyAndAmount `xml:"SlctnFees,omitempty"`

	// Cash disbursement in lieu of a fractional quantity of, for example, equity.
	CashInLieuOfShare *ActiveCurrencyAndAmount `xml:"CshInLieuOfShr,omitempty"`

	// Amount of money distributed as the result of a capital gain.
	CapitalGain *ActiveCurrencyAndAmount `xml:"CptlGn,omitempty"`

	// Amount of money representing a coupon payment.
	InterestAmount *ActiveCurrencyAndAmount `xml:"IntrstAmt,omitempty"`

	// Amount of money resulting from a market claim.
	MarketClaimAmount *ActiveCurrencyAndAmount `xml:"MktClmAmt,omitempty"`

	// (Unique to France) Amount due to a buyer of securities dealt prior to ex date which may be subject to different rate of taxation.
	IndemnityAmount *ActiveCurrencyAndAmount `xml:"IndmntyAmt,omitempty"`

	// Amount of money that the borrower pays to the lender as a compensation. It does not entitle the lender to reclaim any tax credit and is sometimes treated differently by the local tax authorities of the lender. Also covers compensation/indemnity of missed dividend concerning early/late settlements if applicable to a market.
	ManufacturedDividendPaymentAmount *ActiveCurrencyAndAmount `xml:"ManfctrdDvddPmtAmt,omitempty"`

	// Amount of money reinvested in additional securities.
	ReinvestmentAmount *ActiveCurrencyAndAmount `xml:"RinvstmtAmt,omitempty"`

	// Amount resulting from a fully franked dividend paid by a company; amount includes tax credit for companies that have made sufficient tax payments during the fiscal period.
	FullyFrankedAmount *ActiveCurrencyAndAmount `xml:"FullyFrnkdAmt,omitempty"`

	// Amount resulting from an unfranked dividend paid by a company; the amount does not include tax credit and is subject to withholding tax.
	UnfrankedAmount *ActiveCurrencyAndAmount `xml:"UfrnkdAmt,omitempty"`

	// Amount of money related to taxable income that cannot be categorised.
	SundryOrOtherAmount *ActiveCurrencyAndAmount `xml:"SndryOrOthrAmt,omitempty"`

	// Amount of money that has not been subject to taxation.
	TaxFreeAmount *ActiveCurrencyAndAmount `xml:"TaxFreeAmt,omitempty"`

	// Amount of income eligible for deferred taxation.
	TaxDeferredAmount *ActiveCurrencyAndAmount `xml:"TaxDfrrdAmt,omitempty"`

	// Amount of value added tax.
	ValueAddedTaxAmount *ActiveCurrencyAndAmount `xml:"ValAddedTaxAmt,omitempty"`

	// Amount of stamp duty.
	StampDutyAmount *ActiveCurrencyAndAmount `xml:"StmpDtyAmt,omitempty"`

	// Amount that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimAmount *ActiveCurrencyAndAmount `xml:"TaxRclmAmt,omitempty"`

	// Amount of taxes that have been previously paid in relation to the taxable event.
	TaxCreditAmount *ActiveCurrencyAndAmount `xml:"TaxCdtAmt,omitempty"`

	// Amount of money withheld by the jurisdiction in which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgOfFrgnTaxAmt,omitempty"`

	// Amount of money withheld by the jurisdiction in which the account owner is located, for which relief at source and/or reclaim may be possible.
	WithholdingOfLocalTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgOfLclTaxAmt,omitempty"`

	// Amount of additional taxes that cannot be categorised.
	AdditionalTaxAmount *ActiveCurrencyAndAmount `xml:"AddtlTaxAmt,omitempty"`

	// Amount of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgTaxAmt,omitempty"`

	// Amount of fiscal tax to apply.
	FiscalStampAmount *ActiveCurrencyAndAmount `xml:"FsclStmpAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *ActiveCurrencyAndAmount `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of paying/sub-paying agent commission.
	PayingAgentCommissionAmount *ActiveCurrencyAndAmount `xml:"PngAgtComssnAmt,omitempty"`

	// Local broker's commission.
	LocalBrokerCommissionAmount *ActiveCurrencyAndAmount `xml:"LclBrkrComssnAmt,omitempty"`

	// Amount of money charged by a regulatory authority, for example, securities and exchange fees.
	RegulatoryFeesAmount *ActiveCurrencyAndAmount `xml:"RgltryFeesAmt,omitempty"`

	// All costs related to the physical delivery of documents such as stamps, postage, carrier fees, insurances or messenger services.
	ShippingFeesAmount *ActiveCurrencyAndAmount `xml:"ShppgFeesAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesAmount *ActiveCurrencyAndAmount `xml:"ChrgsAmt,omitempty"`

	// Cash amount based on terms of corporate action event and balance of underlying securities, entitled to/from account owner (which may be positive or negative).
	EntitledAmount *ActiveCurrencyAndAmount `xml:"EntitldAmt,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalAmount *ActiveCurrencyAndAmount `xml:"OrgnlAmt,omitempty"`

	// Amount of interest that has been accrued in between coupon payment periods
	AccruedInterestAmount *ActiveCurrencyAndAmount `xml:"AcrdIntrstAmt,omitempty"`

	// Amount relating to the underlying security for which income is distributed.
	IncomePortion *ActiveCurrencyAndAmount `xml:"IncmPrtn,omitempty"`

	// Portion of the fund distribution amount which represents the average accrued income included in the purchase price for units bought during the account period.
	EqualisationAmount *ActiveCurrencyAndAmount `xml:"EqulstnAmt,omitempty"`
}

func (c *CorporateActionAmounts22) SetGrossCashAmount(value, currency string) {
	c.GrossCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts22) SetNetCashAmount(value, currency string) {
	c.NetCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts22) SetSolicitationFees(value, currency string) {
	c.SolicitationFees = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts22) SetCashInLieuOfShare(value, currency string) {
	c.CashInLieuOfShare = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts22) SetCapitalGain(value, currency string) {
	c.CapitalGain = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts22) SetInterestAmount(value, currency string) {
	c.InterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts22) SetMarketClaimAmount(value, currency string) {
	c.MarketClaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts22) SetIndemnityAmount(value, currency string) {
	c.IndemnityAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts22) SetManufacturedDividendPaymentAmount(value, currency string) {
	c.ManufacturedDividendPaymentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts22) SetReinvestmentAmount(value, currency string) {
	c.ReinvestmentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts22) SetFullyFrankedAmount(value, currency string) {
	c.FullyFrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts22) SetUnfrankedAmount(value, currency string) {
	c.UnfrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts22) SetSundryOrOtherAmount(value, currency string) {
	c.SundryOrOtherAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts22) SetTaxFreeAmount(value, currency string) {
	c.TaxFreeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts22) SetTaxDeferredAmount(value, currency string) {
	c.TaxDeferredAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts22) SetValueAddedTaxAmount(value, currency string) {
	c.ValueAddedTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts22) SetStampDutyAmount(value, currency string) {
	c.StampDutyAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts22) SetTaxReclaimAmount(value, currency string) {
	c.TaxReclaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts22) SetTaxCreditAmount(value, currency string) {
	c.TaxCreditAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts22) SetWithholdingOfForeignTaxAmount(value, currency string) {
	c.WithholdingOfForeignTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts22) SetWithholdingOfLocalTaxAmount(value, currency string) {
	c.WithholdingOfLocalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts22) SetAdditionalTaxAmount(value, currency string) {
	c.AdditionalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts22) SetWithholdingTaxAmount(value, currency string) {
	c.WithholdingTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts22) SetFiscalStampAmount(value, currency string) {
	c.FiscalStampAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts22) SetExecutingBrokerAmount(value, currency string) {
	c.ExecutingBrokerAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts22) SetPayingAgentCommissionAmount(value, currency string) {
	c.PayingAgentCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts22) SetLocalBrokerCommissionAmount(value, currency string) {
	c.LocalBrokerCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts22) SetRegulatoryFeesAmount(value, currency string) {
	c.RegulatoryFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts22) SetShippingFeesAmount(value, currency string) {
	c.ShippingFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts22) SetChargesAmount(value, currency string) {
	c.ChargesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts22) SetEntitledAmount(value, currency string) {
	c.EntitledAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts22) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts22) SetAccruedInterestAmount(value, currency string) {
	c.AccruedInterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts22) SetIncomePortion(value, currency string) {
	c.IncomePortion = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts22) SetEqualisationAmount(value, currency string) {
	c.EqualisationAmount = NewActiveCurrencyAndAmount(value, currency)
}

// Specifies amounts in the framework of a corporate action event.
type CorporateActionAmounts23 struct {

	// Amount of money that is to be/was posted to the account.
	PostingAmount *ActiveCurrencyAndAmount `xml:"PstngAmt"`

	// Amount of money before any deductions and allowances have been made.
	GrossCashAmount *ActiveCurrencyAndAmount `xml:"GrssCshAmt,omitempty"`

	// Amount of money after deductions and allowances have been made, if any, that is, the total amount +/- charges/fees.
	NetCashAmount *ActiveCurrencyAndAmount `xml:"NetCshAmt,omitempty"`

	// Cash premium made available if the securities holder consents or participates to an event, for example consent fees or solicitation fees.
	SolicitationFees *ActiveCurrencyAndAmount `xml:"SlctnFees,omitempty"`

	// Cash disbursement in lieu of a fractional quantity of, for example, equity.
	CashInLieuOfShare *ActiveCurrencyAndAmount `xml:"CshInLieuOfShr,omitempty"`

	// Amount of money distributed as the result of a capital gain.
	CapitalGain *ActiveCurrencyAndAmount `xml:"CptlGn,omitempty"`

	// Amount of money representing a coupon payment.
	InterestAmount *ActiveCurrencyAndAmount `xml:"IntrstAmt,omitempty"`

	// Amount of money resulting from a market claim.
	MarketClaimAmount *ActiveCurrencyAndAmount `xml:"MktClmAmt,omitempty"`

	// (Unique to France) Amount due to a buyer of securities dealt prior to ex date which may be subject to different rate of taxation.
	IndemnityAmount *ActiveCurrencyAndAmount `xml:"IndmntyAmt,omitempty"`

	// Amount of money that the borrower pays to the lender as a compensation. It does not entitle the lender to reclaim any tax credit and is sometimes treated differently by the local tax authorities of the lender. Also covers compensation/indemnity of missed dividend concerning early/late settlements if applicable to a market.
	ManufacturedDividendPaymentAmount *ActiveCurrencyAndAmount `xml:"ManfctrdDvddPmtAmt,omitempty"`

	// Amount of money reinvested in additional securities.
	ReinvestmentAmount *ActiveCurrencyAndAmount `xml:"RinvstmtAmt,omitempty"`

	// Amount resulting from a fully franked dividend paid by a company; amount includes tax credit for companies that have made sufficient tax payments during the fiscal period.
	FullyFrankedAmount *ActiveCurrencyAndAmount `xml:"FullyFrnkdAmt,omitempty"`

	// Amount resulting from an unfranked dividend paid by a company; the amount does not include tax credit and is subject to withholding tax.
	UnfrankedAmount *ActiveCurrencyAndAmount `xml:"UfrnkdAmt,omitempty"`

	// Amount of money related to taxable income that cannot be categorised.
	SundryOrOtherAmount *ActiveCurrencyAndAmount `xml:"SndryOrOthrAmt,omitempty"`

	// Amount of money that has not been subject to taxation.
	TaxFreeAmount *ActiveCurrencyAndAmount `xml:"TaxFreeAmt,omitempty"`

	// Amount of income eligible for deferred taxation.
	TaxDeferredAmount *ActiveCurrencyAndAmount `xml:"TaxDfrrdAmt,omitempty"`

	// Tax on value added.
	ValueAddedTaxAmount *ActiveCurrencyAndAmount `xml:"ValAddedTaxAmt,omitempty"`

	// Amount of stamp duty.
	StampDutyAmount *ActiveCurrencyAndAmount `xml:"StmpDtyAmt,omitempty"`

	// Amount that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimAmount *ActiveCurrencyAndAmount `xml:"TaxRclmAmt,omitempty"`

	// Amount of taxes that have been previously paid in relation to the taxable event.
	TaxCreditAmount *ActiveCurrencyAndAmount `xml:"TaxCdtAmt,omitempty"`

	// Amount of money withheld by the jurisdiction in which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgOfFrgnTaxAmt,omitempty"`

	// Amount of money withheld by the jurisdiction in which the account owner is located, for which relief at source and/or reclaim may be possible.
	WithholdingOfLocalTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgOfLclTaxAmt,omitempty"`

	// Amount of additional taxes that cannot be categorised.
	AdditionalTaxAmount *ActiveCurrencyAndAmount `xml:"AddtlTaxAmt,omitempty"`

	// Amount of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgTaxAmt,omitempty"`

	// Amount of fiscal tax to apply.
	FiscalStampAmount *ActiveCurrencyAndAmount `xml:"FsclStmpAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *ActiveCurrencyAndAmount `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of paying/sub-paying agent commission.
	PayingAgentCommissionAmount *ActiveCurrencyAndAmount `xml:"PngAgtComssnAmt,omitempty"`

	// Local broker's commission.
	LocalBrokerCommissionAmount *ActiveCurrencyAndAmount `xml:"LclBrkrComssnAmt,omitempty"`

	// Amount of money charged by a regulatory authority, for example, Securities and Exchange fees.
	RegulatoryFeesAmount *ActiveCurrencyAndAmount `xml:"RgltryFeesAmt,omitempty"`

	// All costs related to the physical delivery of documents such as stamps, postage, carrier fees, insurances or messenger services.
	ShippingFeesAmount *ActiveCurrencyAndAmount `xml:"ShppgFeesAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesAmount *ActiveCurrencyAndAmount `xml:"ChrgsAmt,omitempty"`

	// Indicates cash retained from previous dividend.
	CashAmountBroughtForward *ActiveCurrencyAndAmount `xml:"CshAmtBrghtFwd,omitempty"`

	// Indicates the balance carried forward to next dividend.
	CashAmountCarriedForward *ActiveCurrencyAndAmount `xml:"CshAmtCrrdFwd,omitempty"`

	// Amount of cash that would have been payable if the dividend had been taken in the form of cash rather than shares.
	NotionalDividendPayableAmount *ActiveCurrencyAndAmount `xml:"NtnlDvddPyblAmt,omitempty"`

	// Where new securities are issued in lieu of a cash dividend, the notional tax is the tax on the amount of cash that would have been paid.
	NotionalTaxAmount *ActiveCurrencyAndAmount `xml:"NtnlTaxAmt,omitempty"`

	// Amount of money paid by the Tax Authorities in addition to the payment of the tax refund itself.
	TaxArrearsAmount *ActiveCurrencyAndAmount `xml:"TaxArrearsAmt,omitempty"`

	// Posting/settlement Amount in its original currency when conversion from/into another currency has occurred.
	OriginalAmount *ActiveCurrencyAndAmount `xml:"OrgnlAmt,omitempty"`

	// Amount of money representing a distribution of a bond's principal, for example, repayment of outstanding debt.
	PrincipalOrCorpus *ActiveCurrencyAndAmount `xml:"PrncplOrCrps,omitempty"`

	// Amount of money (not interest) in addition to the principal at the redemption of a bond.
	RedemptionPremiumAmount *ActiveCurrencyAndAmount `xml:"RedPrmAmt,omitempty"`

	// Amount relating to the underlying security for which income is distributed.
	IncomePortion *ActiveCurrencyAndAmount `xml:"IncmPrtn,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *ActiveCurrencyAndAmount `xml:"StockXchgTax,omitempty"`

	// Total amount of tax withheld at source in conformance with the EU Savings Directive.
	EUTaxRetentionAmount *ActiveCurrencyAndAmount `xml:"EUTaxRtntnAmt,omitempty"`

	// Amount of interest that has been accrued in between coupon payment periods
	AccruedInterestAmount *ActiveCurrencyAndAmount `xml:"AcrdIntrstAmt,omitempty"`

	// Portion of the fund distribution amount which represents the average accrued income included in the purchase price for units bought during the account period.
	EqualisationAmount *ActiveCurrencyAndAmount `xml:"EqulstnAmt,omitempty"`
}

func (c *CorporateActionAmounts23) SetPostingAmount(value, currency string) {
	c.PostingAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetGrossCashAmount(value, currency string) {
	c.GrossCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetNetCashAmount(value, currency string) {
	c.NetCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetSolicitationFees(value, currency string) {
	c.SolicitationFees = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetCashInLieuOfShare(value, currency string) {
	c.CashInLieuOfShare = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetCapitalGain(value, currency string) {
	c.CapitalGain = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetInterestAmount(value, currency string) {
	c.InterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetMarketClaimAmount(value, currency string) {
	c.MarketClaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetIndemnityAmount(value, currency string) {
	c.IndemnityAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetManufacturedDividendPaymentAmount(value, currency string) {
	c.ManufacturedDividendPaymentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetReinvestmentAmount(value, currency string) {
	c.ReinvestmentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetFullyFrankedAmount(value, currency string) {
	c.FullyFrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetUnfrankedAmount(value, currency string) {
	c.UnfrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetSundryOrOtherAmount(value, currency string) {
	c.SundryOrOtherAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetTaxFreeAmount(value, currency string) {
	c.TaxFreeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetTaxDeferredAmount(value, currency string) {
	c.TaxDeferredAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetValueAddedTaxAmount(value, currency string) {
	c.ValueAddedTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetStampDutyAmount(value, currency string) {
	c.StampDutyAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetTaxReclaimAmount(value, currency string) {
	c.TaxReclaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetTaxCreditAmount(value, currency string) {
	c.TaxCreditAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetWithholdingOfForeignTaxAmount(value, currency string) {
	c.WithholdingOfForeignTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetWithholdingOfLocalTaxAmount(value, currency string) {
	c.WithholdingOfLocalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetAdditionalTaxAmount(value, currency string) {
	c.AdditionalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetWithholdingTaxAmount(value, currency string) {
	c.WithholdingTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetFiscalStampAmount(value, currency string) {
	c.FiscalStampAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetExecutingBrokerAmount(value, currency string) {
	c.ExecutingBrokerAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetPayingAgentCommissionAmount(value, currency string) {
	c.PayingAgentCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetLocalBrokerCommissionAmount(value, currency string) {
	c.LocalBrokerCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetRegulatoryFeesAmount(value, currency string) {
	c.RegulatoryFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetShippingFeesAmount(value, currency string) {
	c.ShippingFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetChargesAmount(value, currency string) {
	c.ChargesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetCashAmountBroughtForward(value, currency string) {
	c.CashAmountBroughtForward = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetCashAmountCarriedForward(value, currency string) {
	c.CashAmountCarriedForward = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetNotionalDividendPayableAmount(value, currency string) {
	c.NotionalDividendPayableAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetNotionalTaxAmount(value, currency string) {
	c.NotionalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetTaxArrearsAmount(value, currency string) {
	c.TaxArrearsAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetPrincipalOrCorpus(value, currency string) {
	c.PrincipalOrCorpus = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetRedemptionPremiumAmount(value, currency string) {
	c.RedemptionPremiumAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetIncomePortion(value, currency string) {
	c.IncomePortion = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetStockExchangeTax(value, currency string) {
	c.StockExchangeTax = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetEUTaxRetentionAmount(value, currency string) {
	c.EUTaxRetentionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetAccruedInterestAmount(value, currency string) {
	c.AccruedInterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts23) SetEqualisationAmount(value, currency string) {
	c.EqualisationAmount = NewActiveCurrencyAndAmount(value, currency)
}

// Specifies amounts in the framework of a corporate action event.
type CorporateActionAmounts27 struct {

	// Amount of money before any deductions and allowances have been made.
	GrossCashAmount *ActiveCurrencyAndAmount `xml:"GrssCshAmt,omitempty"`

	// Amount of money after deductions and allowances have been made, if any, that is, the total amount +/- charges/fees.
	NetCashAmount *ActiveCurrencyAndAmount `xml:"NetCshAmt,omitempty"`

	// Cash premium made available if the securities holder consents or participates to an event, for example consent fees or solicitation fees.
	SolicitationFees *ActiveCurrencyAndAmount `xml:"SlctnFees,omitempty"`

	// Cash disbursement in lieu of a fractional quantity of, for example, equity.
	CashInLieuOfShare *ActiveCurrencyAndAmount `xml:"CshInLieuOfShr,omitempty"`

	// Amount of money distributed as the result of a capital gain.
	CapitalGain *ActiveCurrencyAndAmount `xml:"CptlGn,omitempty"`

	// Amount of money representing a coupon payment.
	InterestAmount *ActiveCurrencyAndAmount `xml:"IntrstAmt,omitempty"`

	// (Unique to France) Amount due to a buyer of securities dealt prior to ex date which may be subject to different rate of taxation.
	IndemnityAmount *ActiveCurrencyAndAmount `xml:"IndmntyAmt,omitempty"`

	// Amount of money that the borrower pays to the lender as a compensation. It does not entitle the lender to reclaim any tax credit and is sometimes treated differently by the local tax authorities of the lender. Also covers compensation/indemnity of missed dividend concerning early/late settlements if applicable to a market.
	ManufacturedDividendPaymentAmount *ActiveCurrencyAndAmount `xml:"ManfctrdDvddPmtAmt,omitempty"`

	// Amount of money reinvested in additional securities.
	ReinvestmentAmount *ActiveCurrencyAndAmount `xml:"RinvstmtAmt,omitempty"`

	// Amount resulting from a fully franked dividend paid by a company; amount includes tax credit for companies that have made sufficient tax payments during the fiscal period.
	FullyFrankedAmount *ActiveCurrencyAndAmount `xml:"FullyFrnkdAmt,omitempty"`

	// Amount resulting from an unfranked dividend paid by a company; the amount does not include tax credit and is subject to withholding tax.
	UnfrankedAmount *ActiveCurrencyAndAmount `xml:"UfrnkdAmt,omitempty"`

	// Amount of money related to taxable income that cannot be categorised.
	SundryOrOtherAmount *ActiveCurrencyAndAmount `xml:"SndryOrOthrAmt,omitempty"`

	// Amount of money that has not been subject to taxation.
	TaxFreeAmount *ActiveCurrencyAndAmount `xml:"TaxFreeAmt,omitempty"`

	// Amount of income eligible for deferred taxation.
	TaxDeferredAmount *ActiveCurrencyAndAmount `xml:"TaxDfrrdAmt,omitempty"`

	// Amount of value added tax.
	ValueAddedTaxAmount *ActiveCurrencyAndAmount `xml:"ValAddedTaxAmt,omitempty"`

	// Amount of stamp duty.
	StampDutyAmount *ActiveCurrencyAndAmount `xml:"StmpDtyAmt,omitempty"`

	// Amount that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimAmount *ActiveCurrencyAndAmount `xml:"TaxRclmAmt,omitempty"`

	// Amount of taxes that have been previously paid in relation to the taxable event.
	TaxCreditAmount *ActiveCurrencyAndAmount `xml:"TaxCdtAmt,omitempty"`

	// Amount of money withheld by the jurisdiction in which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgOfFrgnTaxAmt,omitempty"`

	// Amount of money withheld by the jurisdiction in which the account owner is located, for which relief at source and/or reclaim may be possible.
	WithholdingOfLocalTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgOfLclTaxAmt,omitempty"`

	// Amount of additional taxes that cannot be categorised.
	AdditionalTaxAmount *ActiveCurrencyAndAmount `xml:"AddtlTaxAmt,omitempty"`

	// Amount of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgTaxAmt,omitempty"`

	// Amount of fiscal tax to apply.
	FiscalStampAmount *ActiveCurrencyAndAmount `xml:"FsclStmpAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *ActiveCurrencyAndAmount `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of paying/sub-paying agent commission.
	PayingAgentCommissionAmount *ActiveCurrencyAndAmount `xml:"PngAgtComssnAmt,omitempty"`

	// Local broker's commission.
	LocalBrokerCommissionAmount *ActiveCurrencyAndAmount `xml:"LclBrkrComssnAmt,omitempty"`

	// Amount of money charged by a regulatory authority, for example, securities and exchange fees.
	RegulatoryFeesAmount *ActiveCurrencyAndAmount `xml:"RgltryFeesAmt,omitempty"`

	// All costs related to the physical delivery of documents such as stamps, postage, carrier fees, insurances or messenger services.
	ShippingFeesAmount *ActiveCurrencyAndAmount `xml:"ShppgFeesAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesAmount *ActiveCurrencyAndAmount `xml:"ChrgsAmt,omitempty"`

	// Cash amount based on terms of corporate action event and balance of underlying securities, entitled to/from account owner (which may be positive or negative).
	EntitledAmount *ActiveCurrencyAndAmount `xml:"EntitldAmt,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalAmount *ActiveCurrencyAndAmount `xml:"OrgnlAmt,omitempty"`

	// Amount of money representing a distribution of a bond's principal, for example, repayment of outstanding debt.
	PrincipalOrCorpus *ActiveCurrencyAndAmount `xml:"PrncplOrCrps,omitempty"`

	// Amount of money (not interest) in addition to the principal at the redemption of a bond.
	RedemptionPremiumAmount *ActiveCurrencyAndAmount `xml:"RedPrmAmt,omitempty"`

	// Amount relating to the underlying security for which income is distributed.
	IncomePortion *ActiveCurrencyAndAmount `xml:"IncmPrtn,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *ActiveCurrencyAndAmount `xml:"StockXchgTax,omitempty"`

	// Total amount of tax withheld at source in conformance with the EU Savings Directive.
	EUTaxRetentionAmount *ActiveCurrencyAndAmount `xml:"EUTaxRtntnAmt,omitempty"`

	// Amount of interest that has been accrued in between coupon payment periods
	AccruedInterestAmount *ActiveCurrencyAndAmount `xml:"AcrdIntrstAmt,omitempty"`

	// Portion of the fund distribution amount which represents the average accrued income included in the purchase price for units bought during the account period.
	EqualisationAmount *ActiveCurrencyAndAmount `xml:"EqulstnAmt,omitempty"`

	// FATCA (Foreign Account Tax Compliance Act) related tax amount.
	FATCATaxAmount *ActiveCurrencyAndAmount `xml:"FATCATaxAmt,omitempty"`

	// Amount of tax related income subject to NRA (Non Resident Alien).
	NRATaxAmount *ActiveCurrencyAndAmount `xml:"NRATaxAmt,omitempty"`

	// Amount of tax related to back up withholding.
	BackUpWithholdingTaxAmount *ActiveCurrencyAndAmount `xml:"BckUpWhldgTaxAmt,omitempty"`

	// Amount of overall tax withheld at source by fund managers prior to considering the tax obligation of each unit holder.
	TaxOnIncomeAmount *ActiveCurrencyAndAmount `xml:"TaxOnIncmAmt,omitempty"`

	// Amount of Transaction tax.
	TransactionTax *ActiveCurrencyAndAmount `xml:"TxTax,omitempty"`
}

func (c *CorporateActionAmounts27) SetGrossCashAmount(value, currency string) {
	c.GrossCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetNetCashAmount(value, currency string) {
	c.NetCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetSolicitationFees(value, currency string) {
	c.SolicitationFees = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetCashInLieuOfShare(value, currency string) {
	c.CashInLieuOfShare = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetCapitalGain(value, currency string) {
	c.CapitalGain = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetInterestAmount(value, currency string) {
	c.InterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetIndemnityAmount(value, currency string) {
	c.IndemnityAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetManufacturedDividendPaymentAmount(value, currency string) {
	c.ManufacturedDividendPaymentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetReinvestmentAmount(value, currency string) {
	c.ReinvestmentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetFullyFrankedAmount(value, currency string) {
	c.FullyFrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetUnfrankedAmount(value, currency string) {
	c.UnfrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetSundryOrOtherAmount(value, currency string) {
	c.SundryOrOtherAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetTaxFreeAmount(value, currency string) {
	c.TaxFreeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetTaxDeferredAmount(value, currency string) {
	c.TaxDeferredAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetValueAddedTaxAmount(value, currency string) {
	c.ValueAddedTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetStampDutyAmount(value, currency string) {
	c.StampDutyAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetTaxReclaimAmount(value, currency string) {
	c.TaxReclaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetTaxCreditAmount(value, currency string) {
	c.TaxCreditAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetWithholdingOfForeignTaxAmount(value, currency string) {
	c.WithholdingOfForeignTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetWithholdingOfLocalTaxAmount(value, currency string) {
	c.WithholdingOfLocalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetAdditionalTaxAmount(value, currency string) {
	c.AdditionalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetWithholdingTaxAmount(value, currency string) {
	c.WithholdingTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetFiscalStampAmount(value, currency string) {
	c.FiscalStampAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetExecutingBrokerAmount(value, currency string) {
	c.ExecutingBrokerAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetPayingAgentCommissionAmount(value, currency string) {
	c.PayingAgentCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetLocalBrokerCommissionAmount(value, currency string) {
	c.LocalBrokerCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetRegulatoryFeesAmount(value, currency string) {
	c.RegulatoryFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetShippingFeesAmount(value, currency string) {
	c.ShippingFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetChargesAmount(value, currency string) {
	c.ChargesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetEntitledAmount(value, currency string) {
	c.EntitledAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetPrincipalOrCorpus(value, currency string) {
	c.PrincipalOrCorpus = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetRedemptionPremiumAmount(value, currency string) {
	c.RedemptionPremiumAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetIncomePortion(value, currency string) {
	c.IncomePortion = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetStockExchangeTax(value, currency string) {
	c.StockExchangeTax = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetEUTaxRetentionAmount(value, currency string) {
	c.EUTaxRetentionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetAccruedInterestAmount(value, currency string) {
	c.AccruedInterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetEqualisationAmount(value, currency string) {
	c.EqualisationAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetFATCATaxAmount(value, currency string) {
	c.FATCATaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetNRATaxAmount(value, currency string) {
	c.NRATaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetBackUpWithholdingTaxAmount(value, currency string) {
	c.BackUpWithholdingTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetTaxOnIncomeAmount(value, currency string) {
	c.TaxOnIncomeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts27) SetTransactionTax(value, currency string) {
	c.TransactionTax = NewActiveCurrencyAndAmount(value, currency)
}

// Specifies amounts in the framework of a corporate action event.
type CorporateActionAmounts28 struct {

	// Amount of money before any deductions and allowances have been made.
	GrossCashAmount *ActiveCurrencyAndAmount `xml:"GrssCshAmt,omitempty"`

	// Amount of money after deductions and allowances have been made, if any, that is, the total amount +/- charges/fees.
	NetCashAmount *ActiveCurrencyAndAmount `xml:"NetCshAmt,omitempty"`

	// Cash premium made available if the securities holder consents or participates to an event, for example consent fees or solicitation fees.
	SolicitationFees *ActiveCurrencyAndAmount `xml:"SlctnFees,omitempty"`

	// Cash disbursement in lieu of a fractional quantity of, for example, equity.
	CashInLieuOfShare *ActiveCurrencyAndAmount `xml:"CshInLieuOfShr,omitempty"`

	// Amount of money distributed as the result of a capital gain.
	CapitalGain *ActiveCurrencyAndAmount `xml:"CptlGn,omitempty"`

	// Amount of money representing a coupon payment.
	InterestAmount *ActiveCurrencyAndAmount `xml:"IntrstAmt,omitempty"`

	// Amount of money resulting from a market claim.
	MarketClaimAmount *ActiveCurrencyAndAmount `xml:"MktClmAmt,omitempty"`

	// (Unique to France) Amount due to a buyer of securities dealt prior to ex date which may be subject to different rate of taxation.
	IndemnityAmount *ActiveCurrencyAndAmount `xml:"IndmntyAmt,omitempty"`

	// Amount of money that the borrower pays to the lender as a compensation. It does not entitle the lender to reclaim any tax credit and is sometimes treated differently by the local tax authorities of the lender. Also covers compensation/indemnity of missed dividend concerning early/late settlements if applicable to a market.
	ManufacturedDividendPaymentAmount *ActiveCurrencyAndAmount `xml:"ManfctrdDvddPmtAmt,omitempty"`

	// Amount of money reinvested in additional securities.
	ReinvestmentAmount *ActiveCurrencyAndAmount `xml:"RinvstmtAmt,omitempty"`

	// Amount resulting from a fully franked dividend paid by a company; amount includes tax credit for companies that have made sufficient tax payments during the fiscal period.
	FullyFrankedAmount *ActiveCurrencyAndAmount `xml:"FullyFrnkdAmt,omitempty"`

	// Amount resulting from an unfranked dividend paid by a company; the amount does not include tax credit and is subject to withholding tax.
	UnfrankedAmount *ActiveCurrencyAndAmount `xml:"UfrnkdAmt,omitempty"`

	// Amount of money related to taxable income that cannot be categorised.
	SundryOrOtherAmount *ActiveCurrencyAndAmount `xml:"SndryOrOthrAmt,omitempty"`

	// Amount of money that has not been subject to taxation.
	TaxFreeAmount *ActiveCurrencyAndAmount `xml:"TaxFreeAmt,omitempty"`

	// Amount of income eligible for deferred taxation.
	TaxDeferredAmount *ActiveCurrencyAndAmount `xml:"TaxDfrrdAmt,omitempty"`

	// Amount of value added tax.
	ValueAddedTaxAmount *ActiveCurrencyAndAmount `xml:"ValAddedTaxAmt,omitempty"`

	// Amount of stamp duty.
	StampDutyAmount *ActiveCurrencyAndAmount `xml:"StmpDtyAmt,omitempty"`

	// Amount that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimAmount *ActiveCurrencyAndAmount `xml:"TaxRclmAmt,omitempty"`

	// Amount of taxes that have been previously paid in relation to the taxable event.
	TaxCreditAmount *ActiveCurrencyAndAmount `xml:"TaxCdtAmt,omitempty"`

	// Amount of money withheld by the jurisdiction in which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgOfFrgnTaxAmt,omitempty"`

	// Amount of money withheld by the jurisdiction in which the account owner is located, for which relief at source and/or reclaim may be possible.
	WithholdingOfLocalTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgOfLclTaxAmt,omitempty"`

	// Amount of additional taxes that cannot be categorised.
	AdditionalTaxAmount *ActiveCurrencyAndAmount `xml:"AddtlTaxAmt,omitempty"`

	// Amount of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgTaxAmt,omitempty"`

	// Amount of fiscal tax to apply.
	FiscalStampAmount *ActiveCurrencyAndAmount `xml:"FsclStmpAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *ActiveCurrencyAndAmount `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of paying/sub-paying agent commission.
	PayingAgentCommissionAmount *ActiveCurrencyAndAmount `xml:"PngAgtComssnAmt,omitempty"`

	// Local broker's commission.
	LocalBrokerCommissionAmount *ActiveCurrencyAndAmount `xml:"LclBrkrComssnAmt,omitempty"`

	// Amount of money charged by a regulatory authority, for example, securities and exchange fees.
	RegulatoryFeesAmount *ActiveCurrencyAndAmount `xml:"RgltryFeesAmt,omitempty"`

	// All costs related to the physical delivery of documents such as stamps, postage, carrier fees, insurances or messenger services.
	ShippingFeesAmount *ActiveCurrencyAndAmount `xml:"ShppgFeesAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesAmount *ActiveCurrencyAndAmount `xml:"ChrgsAmt,omitempty"`

	// Cash amount based on terms of corporate action event and balance of underlying securities, entitled to/from account owner (which may be positive or negative).
	EntitledAmount *ActiveCurrencyAndAmount `xml:"EntitldAmt,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalAmount *ActiveCurrencyAndAmount `xml:"OrgnlAmt,omitempty"`

	// Amount of interest that has been accrued in between coupon payment periods
	AccruedInterestAmount *ActiveCurrencyAndAmount `xml:"AcrdIntrstAmt,omitempty"`

	// Amount relating to the underlying security for which income is distributed.
	IncomePortion *ActiveCurrencyAndAmount `xml:"IncmPrtn,omitempty"`

	// Portion of the fund distribution amount which represents the average accrued income included in the purchase price for units bought during the account period.
	EqualisationAmount *ActiveCurrencyAndAmount `xml:"EqulstnAmt,omitempty"`

	// FATCA (Foreign Account Tax Compliance Act) related tax amount.
	FATCATaxAmount *ActiveCurrencyAndAmount `xml:"FATCATaxAmt,omitempty"`

	// Amount of tax related income subject to NRA (Non Resident Alien).
	NRATaxAmount *ActiveCurrencyAndAmount `xml:"NRATaxAmt,omitempty"`

	// Amount of tax related to back up withholding.
	BackUpWithholdingTaxAmount *ActiveCurrencyAndAmount `xml:"BckUpWhldgTaxAmt,omitempty"`

	// Amount of overall tax withheld at source by fund managers prior to considering the tax obligation of each unit holder.
	TaxOnIncomeAmount *ActiveCurrencyAndAmount `xml:"TaxOnIncmAmt,omitempty"`

	// Amount of Transaction tax.
	TransactionTax *ActiveCurrencyAndAmount `xml:"TxTax,omitempty"`
}

func (c *CorporateActionAmounts28) SetGrossCashAmount(value, currency string) {
	c.GrossCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetNetCashAmount(value, currency string) {
	c.NetCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetSolicitationFees(value, currency string) {
	c.SolicitationFees = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetCashInLieuOfShare(value, currency string) {
	c.CashInLieuOfShare = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetCapitalGain(value, currency string) {
	c.CapitalGain = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetInterestAmount(value, currency string) {
	c.InterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetMarketClaimAmount(value, currency string) {
	c.MarketClaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetIndemnityAmount(value, currency string) {
	c.IndemnityAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetManufacturedDividendPaymentAmount(value, currency string) {
	c.ManufacturedDividendPaymentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetReinvestmentAmount(value, currency string) {
	c.ReinvestmentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetFullyFrankedAmount(value, currency string) {
	c.FullyFrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetUnfrankedAmount(value, currency string) {
	c.UnfrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetSundryOrOtherAmount(value, currency string) {
	c.SundryOrOtherAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetTaxFreeAmount(value, currency string) {
	c.TaxFreeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetTaxDeferredAmount(value, currency string) {
	c.TaxDeferredAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetValueAddedTaxAmount(value, currency string) {
	c.ValueAddedTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetStampDutyAmount(value, currency string) {
	c.StampDutyAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetTaxReclaimAmount(value, currency string) {
	c.TaxReclaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetTaxCreditAmount(value, currency string) {
	c.TaxCreditAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetWithholdingOfForeignTaxAmount(value, currency string) {
	c.WithholdingOfForeignTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetWithholdingOfLocalTaxAmount(value, currency string) {
	c.WithholdingOfLocalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetAdditionalTaxAmount(value, currency string) {
	c.AdditionalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetWithholdingTaxAmount(value, currency string) {
	c.WithholdingTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetFiscalStampAmount(value, currency string) {
	c.FiscalStampAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetExecutingBrokerAmount(value, currency string) {
	c.ExecutingBrokerAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetPayingAgentCommissionAmount(value, currency string) {
	c.PayingAgentCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetLocalBrokerCommissionAmount(value, currency string) {
	c.LocalBrokerCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetRegulatoryFeesAmount(value, currency string) {
	c.RegulatoryFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetShippingFeesAmount(value, currency string) {
	c.ShippingFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetChargesAmount(value, currency string) {
	c.ChargesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetEntitledAmount(value, currency string) {
	c.EntitledAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetAccruedInterestAmount(value, currency string) {
	c.AccruedInterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetIncomePortion(value, currency string) {
	c.IncomePortion = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetEqualisationAmount(value, currency string) {
	c.EqualisationAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetFATCATaxAmount(value, currency string) {
	c.FATCATaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetNRATaxAmount(value, currency string) {
	c.NRATaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetBackUpWithholdingTaxAmount(value, currency string) {
	c.BackUpWithholdingTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetTaxOnIncomeAmount(value, currency string) {
	c.TaxOnIncomeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts28) SetTransactionTax(value, currency string) {
	c.TransactionTax = NewActiveCurrencyAndAmount(value, currency)
}

// Specifies amounts in the framework of a corporate action event.
type CorporateActionAmounts29 struct {

	// Amount of money that is to be/was posted to the account.
	PostingAmount *ActiveCurrencyAndAmount `xml:"PstngAmt"`

	// Amount of money before any deductions and allowances have been made.
	GrossCashAmount *ActiveCurrencyAndAmount `xml:"GrssCshAmt,omitempty"`

	// Amount of money after deductions and allowances have been made, if any, that is, the total amount +/- charges/fees.
	NetCashAmount *ActiveCurrencyAndAmount `xml:"NetCshAmt,omitempty"`

	// Cash premium made available if the securities holder consents or participates to an event, for example consent fees or solicitation fees.
	SolicitationFees *ActiveCurrencyAndAmount `xml:"SlctnFees,omitempty"`

	// Cash disbursement in lieu of a fractional quantity of, for example, equity.
	CashInLieuOfShare *ActiveCurrencyAndAmount `xml:"CshInLieuOfShr,omitempty"`

	// Amount of money distributed as the result of a capital gain.
	CapitalGain *ActiveCurrencyAndAmount `xml:"CptlGn,omitempty"`

	// Amount of money representing a coupon payment.
	InterestAmount *ActiveCurrencyAndAmount `xml:"IntrstAmt,omitempty"`

	// Amount of money resulting from a market claim.
	MarketClaimAmount *ActiveCurrencyAndAmount `xml:"MktClmAmt,omitempty"`

	// (Unique to France) Amount due to a buyer of securities dealt prior to ex date which may be subject to different rate of taxation.
	IndemnityAmount *ActiveCurrencyAndAmount `xml:"IndmntyAmt,omitempty"`

	// Amount of money that the borrower pays to the lender as a compensation. It does not entitle the lender to reclaim any tax credit and is sometimes treated differently by the local tax authorities of the lender. Also covers compensation/indemnity of missed dividend concerning early/late settlements if applicable to a market.
	ManufacturedDividendPaymentAmount *ActiveCurrencyAndAmount `xml:"ManfctrdDvddPmtAmt,omitempty"`

	// Amount of money reinvested in additional securities.
	ReinvestmentAmount *ActiveCurrencyAndAmount `xml:"RinvstmtAmt,omitempty"`

	// Amount resulting from a fully franked dividend paid by a company; amount includes tax credit for companies that have made sufficient tax payments during the fiscal period.
	FullyFrankedAmount *ActiveCurrencyAndAmount `xml:"FullyFrnkdAmt,omitempty"`

	// Amount resulting from an unfranked dividend paid by a company; the amount does not include tax credit and is subject to withholding tax.
	UnfrankedAmount *ActiveCurrencyAndAmount `xml:"UfrnkdAmt,omitempty"`

	// Amount of money related to taxable income that cannot be categorised.
	SundryOrOtherAmount *ActiveCurrencyAndAmount `xml:"SndryOrOthrAmt,omitempty"`

	// Amount of money that has not been subject to taxation.
	TaxFreeAmount *ActiveCurrencyAndAmount `xml:"TaxFreeAmt,omitempty"`

	// Amount of income eligible for deferred taxation.
	TaxDeferredAmount *ActiveCurrencyAndAmount `xml:"TaxDfrrdAmt,omitempty"`

	// Tax on value added.
	ValueAddedTaxAmount *ActiveCurrencyAndAmount `xml:"ValAddedTaxAmt,omitempty"`

	// Amount of stamp duty.
	StampDutyAmount *ActiveCurrencyAndAmount `xml:"StmpDtyAmt,omitempty"`

	// Amount that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimAmount *ActiveCurrencyAndAmount `xml:"TaxRclmAmt,omitempty"`

	// Amount of taxes that have been previously paid in relation to the taxable event.
	TaxCreditAmount *ActiveCurrencyAndAmount `xml:"TaxCdtAmt,omitempty"`

	// Amount of money withheld by the jurisdiction in which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgOfFrgnTaxAmt,omitempty"`

	// Amount of money withheld by the jurisdiction in which the account owner is located, for which relief at source and/or reclaim may be possible.
	WithholdingOfLocalTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgOfLclTaxAmt,omitempty"`

	// Amount of additional taxes that cannot be categorised.
	AdditionalTaxAmount *ActiveCurrencyAndAmount `xml:"AddtlTaxAmt,omitempty"`

	// Amount of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgTaxAmt,omitempty"`

	// Amount of fiscal tax to apply.
	FiscalStampAmount *ActiveCurrencyAndAmount `xml:"FsclStmpAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *ActiveCurrencyAndAmount `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of paying/sub-paying agent commission.
	PayingAgentCommissionAmount *ActiveCurrencyAndAmount `xml:"PngAgtComssnAmt,omitempty"`

	// Local broker's commission.
	LocalBrokerCommissionAmount *ActiveCurrencyAndAmount `xml:"LclBrkrComssnAmt,omitempty"`

	// Amount of money charged by a regulatory authority, for example, Securities and Exchange fees.
	RegulatoryFeesAmount *ActiveCurrencyAndAmount `xml:"RgltryFeesAmt,omitempty"`

	// All costs related to the physical delivery of documents such as stamps, postage, carrier fees, insurances or messenger services.
	ShippingFeesAmount *ActiveCurrencyAndAmount `xml:"ShppgFeesAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesAmount *ActiveCurrencyAndAmount `xml:"ChrgsAmt,omitempty"`

	// Indicates cash retained from previous dividend.
	CashAmountBroughtForward *ActiveCurrencyAndAmount `xml:"CshAmtBrghtFwd,omitempty"`

	// Indicates the balance carried forward to next dividend.
	CashAmountCarriedForward *ActiveCurrencyAndAmount `xml:"CshAmtCrrdFwd,omitempty"`

	// Amount of cash that would have been payable if the dividend had been taken in the form of cash rather than shares.
	NotionalDividendPayableAmount *ActiveCurrencyAndAmount `xml:"NtnlDvddPyblAmt,omitempty"`

	// Where new securities are issued in lieu of a cash dividend, the notional tax is the tax on the amount of cash that would have been paid.
	NotionalTaxAmount *ActiveCurrencyAndAmount `xml:"NtnlTaxAmt,omitempty"`

	// Amount of money paid by the Tax Authorities in addition to the payment of the tax refund itself.
	TaxArrearsAmount *ActiveCurrencyAndAmount `xml:"TaxArrearsAmt,omitempty"`

	// Posting/settlement Amount in its original currency when conversion from/into another currency has occurred.
	OriginalAmount *ActiveCurrencyAndAmount `xml:"OrgnlAmt,omitempty"`

	// Amount of money representing a distribution of a bond's principal, for example, repayment of outstanding debt.
	PrincipalOrCorpus *ActiveCurrencyAndAmount `xml:"PrncplOrCrps,omitempty"`

	// Amount of money (not interest) in addition to the principal at the redemption of a bond.
	RedemptionPremiumAmount *ActiveCurrencyAndAmount `xml:"RedPrmAmt,omitempty"`

	// Amount relating to the underlying security for which income is distributed.
	IncomePortion *ActiveCurrencyAndAmount `xml:"IncmPrtn,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *ActiveCurrencyAndAmount `xml:"StockXchgTax,omitempty"`

	// Total amount of tax withheld at source in conformance with the EU Savings Directive.
	EUTaxRetentionAmount *ActiveCurrencyAndAmount `xml:"EUTaxRtntnAmt,omitempty"`

	// Amount of interest that has been accrued in between coupon payment periods
	AccruedInterestAmount *ActiveCurrencyAndAmount `xml:"AcrdIntrstAmt,omitempty"`

	// Portion of the fund distribution amount which represents the average accrued income included in the purchase price for units bought during the account period.
	EqualisationAmount *ActiveCurrencyAndAmount `xml:"EqulstnAmt,omitempty"`

	// FATCA (Foreign Account Tax Compliance Act) related tax amount.
	FATCATaxAmount *ActiveCurrencyAndAmount `xml:"FATCATaxAmt,omitempty"`

	// Amount of tax related income subject to NRA (Non Resident Alien).
	NRATaxAmount *ActiveCurrencyAndAmount `xml:"NRATaxAmt,omitempty"`

	// Amount of tax related to back up withholding.
	BackUpWithholdingTaxAmount *ActiveCurrencyAndAmount `xml:"BckUpWhldgTaxAmt,omitempty"`

	// Amount of overall tax withheld at source by fund managers prior to considering the tax obligation of each unit holder.
	TaxOnIncomeAmount *ActiveCurrencyAndAmount `xml:"TaxOnIncmAmt,omitempty"`

	// Amount of Transaction tax.
	TransactionTax *ActiveCurrencyAndAmount `xml:"TxTax,omitempty"`
}

func (c *CorporateActionAmounts29) SetPostingAmount(value, currency string) {
	c.PostingAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetGrossCashAmount(value, currency string) {
	c.GrossCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetNetCashAmount(value, currency string) {
	c.NetCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetSolicitationFees(value, currency string) {
	c.SolicitationFees = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetCashInLieuOfShare(value, currency string) {
	c.CashInLieuOfShare = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetCapitalGain(value, currency string) {
	c.CapitalGain = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetInterestAmount(value, currency string) {
	c.InterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetMarketClaimAmount(value, currency string) {
	c.MarketClaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetIndemnityAmount(value, currency string) {
	c.IndemnityAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetManufacturedDividendPaymentAmount(value, currency string) {
	c.ManufacturedDividendPaymentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetReinvestmentAmount(value, currency string) {
	c.ReinvestmentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetFullyFrankedAmount(value, currency string) {
	c.FullyFrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetUnfrankedAmount(value, currency string) {
	c.UnfrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetSundryOrOtherAmount(value, currency string) {
	c.SundryOrOtherAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetTaxFreeAmount(value, currency string) {
	c.TaxFreeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetTaxDeferredAmount(value, currency string) {
	c.TaxDeferredAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetValueAddedTaxAmount(value, currency string) {
	c.ValueAddedTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetStampDutyAmount(value, currency string) {
	c.StampDutyAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetTaxReclaimAmount(value, currency string) {
	c.TaxReclaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetTaxCreditAmount(value, currency string) {
	c.TaxCreditAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetWithholdingOfForeignTaxAmount(value, currency string) {
	c.WithholdingOfForeignTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetWithholdingOfLocalTaxAmount(value, currency string) {
	c.WithholdingOfLocalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetAdditionalTaxAmount(value, currency string) {
	c.AdditionalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetWithholdingTaxAmount(value, currency string) {
	c.WithholdingTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetFiscalStampAmount(value, currency string) {
	c.FiscalStampAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetExecutingBrokerAmount(value, currency string) {
	c.ExecutingBrokerAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetPayingAgentCommissionAmount(value, currency string) {
	c.PayingAgentCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetLocalBrokerCommissionAmount(value, currency string) {
	c.LocalBrokerCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetRegulatoryFeesAmount(value, currency string) {
	c.RegulatoryFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetShippingFeesAmount(value, currency string) {
	c.ShippingFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetChargesAmount(value, currency string) {
	c.ChargesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetCashAmountBroughtForward(value, currency string) {
	c.CashAmountBroughtForward = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetCashAmountCarriedForward(value, currency string) {
	c.CashAmountCarriedForward = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetNotionalDividendPayableAmount(value, currency string) {
	c.NotionalDividendPayableAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetNotionalTaxAmount(value, currency string) {
	c.NotionalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetTaxArrearsAmount(value, currency string) {
	c.TaxArrearsAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetPrincipalOrCorpus(value, currency string) {
	c.PrincipalOrCorpus = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetRedemptionPremiumAmount(value, currency string) {
	c.RedemptionPremiumAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetIncomePortion(value, currency string) {
	c.IncomePortion = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetStockExchangeTax(value, currency string) {
	c.StockExchangeTax = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetEUTaxRetentionAmount(value, currency string) {
	c.EUTaxRetentionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetAccruedInterestAmount(value, currency string) {
	c.AccruedInterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetEqualisationAmount(value, currency string) {
	c.EqualisationAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetFATCATaxAmount(value, currency string) {
	c.FATCATaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetNRATaxAmount(value, currency string) {
	c.NRATaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetBackUpWithholdingTaxAmount(value, currency string) {
	c.BackUpWithholdingTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetTaxOnIncomeAmount(value, currency string) {
	c.TaxOnIncomeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts29) SetTransactionTax(value, currency string) {
	c.TransactionTax = NewActiveCurrencyAndAmount(value, currency)
}

// Specifies amounts in the framework of a corporate action event.
type CorporateActionAmounts3 struct {

	// Amount of money before any deductions and allowances have been made.
	GrossCashAmount *ActiveCurrencyAndAmount `xml:"GrssCshAmt,omitempty"`

	// Amount of money after deductions and allowances have been made, if any, that is, the total amount +/- charges/fees.
	NetCashAmount *ActiveCurrencyAndAmount `xml:"NetCshAmt,omitempty"`

	// Amount of cash premium made available in order to encourage participation in the offer. Payment is made to a third party who has solicited an entity to take part in the offer.
	SolicitationFees *ActiveCurrencyAndAmount `xml:"SlctnFees,omitempty"`

	// Cash disbursement in lieu of a fractional quantity of, for example, equity.
	CashInLieuOfShare *ActiveCurrencyAndAmount `xml:"CshInLieuOfShr,omitempty"`

	// Amount of money distributed as the result of a capital gain.
	CapitalGain *ActiveCurrencyAndAmount `xml:"CptlGn,omitempty"`

	// Amount of money representing a coupon payment.
	InterestAmount *ActiveCurrencyAndAmount `xml:"IntrstAmt,omitempty"`

	// Amount of money resulting from a market claim.
	MarketClaimAmount *ActiveCurrencyAndAmount `xml:"MktClmAmt,omitempty"`

	// (Unique to France) Amount due to a buyer of securities dealt prior to ex date which may be subject to different rate of taxation.
	IndemnityAmount *ActiveCurrencyAndAmount `xml:"IndmntyAmt,omitempty"`

	// Amount of money that the borrower pays to the lender as a compensation. It does not entitle the lender to reclaim any tax credit and is sometimes treated differently by the local tax authorities of the lender.
	ManufacturedDividendAmount *ActiveCurrencyAndAmount `xml:"ManfctrdDvddAmt,omitempty"`

	// Amount of money reinvested in additional securities.
	ReinvestmentAmount *ActiveCurrencyAndAmount `xml:"RinvstmtAmt,omitempty"`

	// Amount resulting from a fully franked dividend paid by a company; amount includes tax credit for companies that have made sufficient tax payments during the fiscal period.
	FullyFrankedAmount *ActiveCurrencyAndAmount `xml:"FullyFrnkdAmt,omitempty"`

	// Amount resulting from an unfranked dividend paid by a company; the amount does not include tax credit and is subject to withholding tax.
	UnfrankedAmount *ActiveCurrencyAndAmount `xml:"UfrnkdAmt,omitempty"`

	// Amount of money related to taxable income that cannot be categorised.
	SundryOrOtherAmount *ActiveCurrencyAndAmount `xml:"SndryOrOthrAmt,omitempty"`

	// Cash premium made available if the securities holder consents or participates to an event, for example, consent fees or solicitation fees.
	CashIncentive *ActiveCurrencyAndAmount `xml:"CshIncntiv,omitempty"`

	// Amount of money that has not been subject to taxation.
	TaxFreeAmount *ActiveCurrencyAndAmount `xml:"TaxFreeAmt,omitempty"`

	// Amount of income eligible for deferred taxation.
	TaxDeferredAmount *ActiveCurrencyAndAmount `xml:"TaxDfrrdAmt,omitempty"`

	// Amount of value added tax.
	ValueAddedTaxAmount *ActiveCurrencyAndAmount `xml:"ValAddedTaxAmt,omitempty"`

	// Amount of stamp duty.
	StampDutyAmount *ActiveCurrencyAndAmount `xml:"StmpDtyAmt,omitempty"`

	// Amount that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimAmount *ActiveCurrencyAndAmount `xml:"TaxRclmAmt,omitempty"`

	// Amount of taxes that have been previously paid in relation to the taxable event.
	TaxCreditAmount *ActiveCurrencyAndAmount `xml:"TaxCdtAmt,omitempty"`

	// Amount of money withheld by the jurisdiction in which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgOfFrgnTaxAmt,omitempty"`

	// Amount of money withheld by the jurisdiction in which the account owner is located, for which relief at source and/or reclaim may be possible.
	WithholdingOfLocalTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgOfLclTaxAmt,omitempty"`

	// Amount of additional taxes that cannot be categorised.
	AdditionalTaxAmount *ActiveCurrencyAndAmount `xml:"AddtlTaxAmt,omitempty"`

	// Amount of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgTaxAmt,omitempty"`

	// Amount of fiscal tax to apply.
	FiscalStampAmount *ActiveCurrencyAndAmount `xml:"FsclStmpAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *ActiveCurrencyAndAmount `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of paying/sub-paying agent commission.
	PayingAgentCommissionAmount *ActiveCurrencyAndAmount `xml:"PngAgtComssnAmt,omitempty"`

	// Local broker's commission.
	LocalBrokerCommissionAmount *ActiveCurrencyAndAmount `xml:"LclBrkrComssnAmt,omitempty"`

	// Amount of money paid for delivery by regular post mail.
	PostageFeeAmount *ActiveCurrencyAndAmount `xml:"PstgFeeAmt,omitempty"`

	// Amount of money charged by a regulatory authority, for example, securities and exchange fees.
	RegulatoryFeesAmount *ActiveCurrencyAndAmount `xml:"RgltryFeesAmt,omitempty"`

	// Amount of money (including insurance) paid for delivery by carrier.
	ShippingFeesAmount *ActiveCurrencyAndAmount `xml:"ShppgFeesAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesAmount *ActiveCurrencyAndAmount `xml:"ChrgsAmt,omitempty"`

	// Cash amount based on terms of corporate action event and balance of underlying securities, entitled to/from account owner (which may be positive or negative).
	EntitledAmount *ActiveCurrencyAndAmount `xml:"EntitldAmt,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalAmount *ActiveCurrencyAndAmount `xml:"OrgnlAmt,omitempty"`

	// Amount of interest that has been accrued in between coupon payment periods
	AccruedInterestAmount *ActiveCurrencyAndAmount `xml:"AcrdIntrstAmt,omitempty"`

	// Amount relating to the underlying security for which income is distributed.
	IncomePortion *ActiveCurrencyAndAmount `xml:"IncmPrtn,omitempty"`
}

func (c *CorporateActionAmounts3) SetGrossCashAmount(value, currency string) {
	c.GrossCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetNetCashAmount(value, currency string) {
	c.NetCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetSolicitationFees(value, currency string) {
	c.SolicitationFees = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetCashInLieuOfShare(value, currency string) {
	c.CashInLieuOfShare = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetCapitalGain(value, currency string) {
	c.CapitalGain = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetInterestAmount(value, currency string) {
	c.InterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetMarketClaimAmount(value, currency string) {
	c.MarketClaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetIndemnityAmount(value, currency string) {
	c.IndemnityAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetManufacturedDividendAmount(value, currency string) {
	c.ManufacturedDividendAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetReinvestmentAmount(value, currency string) {
	c.ReinvestmentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetFullyFrankedAmount(value, currency string) {
	c.FullyFrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetUnfrankedAmount(value, currency string) {
	c.UnfrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetSundryOrOtherAmount(value, currency string) {
	c.SundryOrOtherAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetCashIncentive(value, currency string) {
	c.CashIncentive = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetTaxFreeAmount(value, currency string) {
	c.TaxFreeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetTaxDeferredAmount(value, currency string) {
	c.TaxDeferredAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetValueAddedTaxAmount(value, currency string) {
	c.ValueAddedTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetStampDutyAmount(value, currency string) {
	c.StampDutyAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetTaxReclaimAmount(value, currency string) {
	c.TaxReclaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetTaxCreditAmount(value, currency string) {
	c.TaxCreditAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetWithholdingOfForeignTaxAmount(value, currency string) {
	c.WithholdingOfForeignTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetWithholdingOfLocalTaxAmount(value, currency string) {
	c.WithholdingOfLocalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetAdditionalTaxAmount(value, currency string) {
	c.AdditionalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetWithholdingTaxAmount(value, currency string) {
	c.WithholdingTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetFiscalStampAmount(value, currency string) {
	c.FiscalStampAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetExecutingBrokerAmount(value, currency string) {
	c.ExecutingBrokerAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetPayingAgentCommissionAmount(value, currency string) {
	c.PayingAgentCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetLocalBrokerCommissionAmount(value, currency string) {
	c.LocalBrokerCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetPostageFeeAmount(value, currency string) {
	c.PostageFeeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetRegulatoryFeesAmount(value, currency string) {
	c.RegulatoryFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetShippingFeesAmount(value, currency string) {
	c.ShippingFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetChargesAmount(value, currency string) {
	c.ChargesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetEntitledAmount(value, currency string) {
	c.EntitledAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetAccruedInterestAmount(value, currency string) {
	c.AccruedInterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts3) SetIncomePortion(value, currency string) {
	c.IncomePortion = NewActiveCurrencyAndAmount(value, currency)
}

// Specifies amounts in the framework of a corporate action event.
type CorporateActionAmounts36 struct {

	// Amount of money before any deductions and allowances have been made.
	GrossCashAmount *ActiveCurrencyAndAmount `xml:"GrssCshAmt,omitempty"`

	// Amount of money after deductions and allowances have been made, if any, that is, the total amount +/- charges/fees.
	NetCashAmount *ActiveCurrencyAndAmount `xml:"NetCshAmt,omitempty"`

	// Cash premium made available if the securities holder consents or participates to an event, for example consent fees or solicitation fees.
	SolicitationFees *ActiveCurrencyAndAmount `xml:"SlctnFees,omitempty"`

	// Cash disbursement in lieu of a fractional quantity of, for example, equity.
	CashInLieuOfShare *ActiveCurrencyAndAmount `xml:"CshInLieuOfShr,omitempty"`

	// Amount of money distributed as the result of a capital gain.
	CapitalGain *ActiveCurrencyAndAmount `xml:"CptlGn,omitempty"`

	// Amount of money representing a coupon payment.
	InterestAmount *ActiveCurrencyAndAmount `xml:"IntrstAmt,omitempty"`

	// (Unique to France) Amount due to a buyer of securities dealt prior to ex date which may be subject to different rate of taxation.
	IndemnityAmount *ActiveCurrencyAndAmount `xml:"IndmntyAmt,omitempty"`

	// Amount of money that the borrower pays to the lender as a compensation. It does not entitle the lender to reclaim any tax credit and is sometimes treated differently by the local tax authorities of the lender. Also covers compensation/indemnity of missed dividend concerning early/late settlements if applicable to a market.
	ManufacturedDividendPaymentAmount *ActiveCurrencyAndAmount `xml:"ManfctrdDvddPmtAmt,omitempty"`

	// Amount of money reinvested in additional securities.
	ReinvestmentAmount *ActiveCurrencyAndAmount `xml:"RinvstmtAmt,omitempty"`

	// Amount resulting from a fully franked dividend paid by a company; amount includes tax credit for companies that have made sufficient tax payments during the fiscal period.
	FullyFrankedAmount *ActiveCurrencyAndAmount `xml:"FullyFrnkdAmt,omitempty"`

	// Amount resulting from an unfranked dividend paid by a company; the amount does not include tax credit and is subject to withholding tax.
	UnfrankedAmount *ActiveCurrencyAndAmount `xml:"UfrnkdAmt,omitempty"`

	// Amount of money related to taxable income that cannot be categorised.
	SundryOrOtherAmount *ActiveCurrencyAndAmount `xml:"SndryOrOthrAmt,omitempty"`

	// Amount of money that has not been subject to taxation.
	TaxFreeAmount *ActiveCurrencyAndAmount `xml:"TaxFreeAmt,omitempty"`

	// Amount of income eligible for deferred taxation.
	TaxDeferredAmount *ActiveCurrencyAndAmount `xml:"TaxDfrrdAmt,omitempty"`

	// Amount of value added tax.
	ValueAddedTaxAmount *ActiveCurrencyAndAmount `xml:"ValAddedTaxAmt,omitempty"`

	// Amount of stamp duty.
	StampDutyAmount *ActiveCurrencyAndAmount `xml:"StmpDtyAmt,omitempty"`

	// Amount that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimAmount *ActiveCurrencyAndAmount `xml:"TaxRclmAmt,omitempty"`

	// Amount of taxes that have been previously paid in relation to the taxable event.
	TaxCreditAmount *ActiveCurrencyAndAmount `xml:"TaxCdtAmt,omitempty"`

	// Amount of additional taxes that cannot be categorised.
	AdditionalTaxAmount *ActiveCurrencyAndAmount `xml:"AddtlTaxAmt,omitempty"`

	// Amount of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgTaxAmt,omitempty"`

	// Amount of money withheld by the jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTaxAmount *ActiveCurrencyAndAmount `xml:"ScndLvlTaxAmt,omitempty"`

	// Amount of fiscal tax to apply.
	FiscalStampAmount *ActiveCurrencyAndAmount `xml:"FsclStmpAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *ActiveCurrencyAndAmount `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of paying/sub-paying agent commission.
	PayingAgentCommissionAmount *ActiveCurrencyAndAmount `xml:"PngAgtComssnAmt,omitempty"`

	// Local broker's commission.
	LocalBrokerCommissionAmount *ActiveCurrencyAndAmount `xml:"LclBrkrComssnAmt,omitempty"`

	// Amount of money charged by a regulatory authority, for example, securities and exchange fees.
	RegulatoryFeesAmount *ActiveCurrencyAndAmount `xml:"RgltryFeesAmt,omitempty"`

	// All costs related to the physical delivery of documents such as stamps, postage, carrier fees, insurances or messenger services.
	ShippingFeesAmount *ActiveCurrencyAndAmount `xml:"ShppgFeesAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesAmount *ActiveCurrencyAndAmount `xml:"ChrgsAmt,omitempty"`

	// Cash amount based on terms of corporate action event and balance of underlying securities, entitled to/from account owner (which may be positive or negative).
	EntitledAmount *ActiveCurrencyAndAmount `xml:"EntitldAmt,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalAmount *ActiveCurrencyAndAmount `xml:"OrgnlAmt,omitempty"`

	// Amount of money representing a distribution of a bond's principal, for example, repayment of outstanding debt.
	PrincipalOrCorpus *ActiveCurrencyAndAmount `xml:"PrncplOrCrps,omitempty"`

	// Amount of money (not interest) in addition to the principal at the redemption of a bond.
	RedemptionPremiumAmount *ActiveCurrencyAndAmount `xml:"RedPrmAmt,omitempty"`

	// Amount relating to the underlying security for which income is distributed.
	IncomePortion *ActiveCurrencyAndAmount `xml:"IncmPrtn,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *ActiveCurrencyAndAmount `xml:"StockXchgTax,omitempty"`

	// Total amount of tax withheld at source in conformance with the EU Savings Directive.
	EUTaxRetentionAmount *ActiveCurrencyAndAmount `xml:"EUTaxRtntnAmt,omitempty"`

	// Amount of interest that has been accrued in between coupon payment periods
	AccruedInterestAmount *ActiveCurrencyAndAmount `xml:"AcrdIntrstAmt,omitempty"`

	// Portion of the fund distribution amount which represents the average accrued income included in the purchase price for units bought during the account period.
	EqualisationAmount *ActiveCurrencyAndAmount `xml:"EqulstnAmt,omitempty"`

	// FATCA (Foreign Account Tax Compliance Act) related tax amount.
	FATCATaxAmount *ActiveCurrencyAndAmount `xml:"FATCATaxAmt,omitempty"`

	// Amount of tax related income subject to NRA (Non Resident Alien).
	NRATaxAmount *ActiveCurrencyAndAmount `xml:"NRATaxAmt,omitempty"`

	// Amount of tax related to back up withholding.
	BackUpWithholdingTaxAmount *ActiveCurrencyAndAmount `xml:"BckUpWhldgTaxAmt,omitempty"`

	// Amount of overall tax withheld at source by fund managers prior to considering the tax obligation of each unit holder.
	TaxOnIncomeAmount *ActiveCurrencyAndAmount `xml:"TaxOnIncmAmt,omitempty"`

	// Amount of Transaction tax.
	TransactionTax *ActiveCurrencyAndAmount `xml:"TxTax,omitempty"`
}

func (c *CorporateActionAmounts36) SetGrossCashAmount(value, currency string) {
	c.GrossCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetNetCashAmount(value, currency string) {
	c.NetCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetSolicitationFees(value, currency string) {
	c.SolicitationFees = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetCashInLieuOfShare(value, currency string) {
	c.CashInLieuOfShare = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetCapitalGain(value, currency string) {
	c.CapitalGain = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetInterestAmount(value, currency string) {
	c.InterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetIndemnityAmount(value, currency string) {
	c.IndemnityAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetManufacturedDividendPaymentAmount(value, currency string) {
	c.ManufacturedDividendPaymentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetReinvestmentAmount(value, currency string) {
	c.ReinvestmentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetFullyFrankedAmount(value, currency string) {
	c.FullyFrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetUnfrankedAmount(value, currency string) {
	c.UnfrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetSundryOrOtherAmount(value, currency string) {
	c.SundryOrOtherAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetTaxFreeAmount(value, currency string) {
	c.TaxFreeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetTaxDeferredAmount(value, currency string) {
	c.TaxDeferredAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetValueAddedTaxAmount(value, currency string) {
	c.ValueAddedTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetStampDutyAmount(value, currency string) {
	c.StampDutyAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetTaxReclaimAmount(value, currency string) {
	c.TaxReclaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetTaxCreditAmount(value, currency string) {
	c.TaxCreditAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetAdditionalTaxAmount(value, currency string) {
	c.AdditionalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetWithholdingTaxAmount(value, currency string) {
	c.WithholdingTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetSecondLevelTaxAmount(value, currency string) {
	c.SecondLevelTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetFiscalStampAmount(value, currency string) {
	c.FiscalStampAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetExecutingBrokerAmount(value, currency string) {
	c.ExecutingBrokerAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetPayingAgentCommissionAmount(value, currency string) {
	c.PayingAgentCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetLocalBrokerCommissionAmount(value, currency string) {
	c.LocalBrokerCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetRegulatoryFeesAmount(value, currency string) {
	c.RegulatoryFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetShippingFeesAmount(value, currency string) {
	c.ShippingFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetChargesAmount(value, currency string) {
	c.ChargesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetEntitledAmount(value, currency string) {
	c.EntitledAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetPrincipalOrCorpus(value, currency string) {
	c.PrincipalOrCorpus = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetRedemptionPremiumAmount(value, currency string) {
	c.RedemptionPremiumAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetIncomePortion(value, currency string) {
	c.IncomePortion = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetStockExchangeTax(value, currency string) {
	c.StockExchangeTax = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetEUTaxRetentionAmount(value, currency string) {
	c.EUTaxRetentionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetAccruedInterestAmount(value, currency string) {
	c.AccruedInterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetEqualisationAmount(value, currency string) {
	c.EqualisationAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetFATCATaxAmount(value, currency string) {
	c.FATCATaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetNRATaxAmount(value, currency string) {
	c.NRATaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetBackUpWithholdingTaxAmount(value, currency string) {
	c.BackUpWithholdingTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetTaxOnIncomeAmount(value, currency string) {
	c.TaxOnIncomeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts36) SetTransactionTax(value, currency string) {
	c.TransactionTax = NewActiveCurrencyAndAmount(value, currency)
}

// Specifies amounts in the framework of a corporate action event.
type CorporateActionAmounts37 struct {

	// Amount of money that is to be/was posted to the account.
	PostingAmount *ActiveCurrencyAndAmount `xml:"PstngAmt"`

	// Amount of money before any deductions and allowances have been made.
	GrossCashAmount *ActiveCurrencyAndAmount `xml:"GrssCshAmt,omitempty"`

	// Amount of money after deductions and allowances have been made, if any, that is, the total amount +/- charges/fees.
	NetCashAmount *ActiveCurrencyAndAmount `xml:"NetCshAmt,omitempty"`

	// Cash premium made available if the securities holder consents or participates to an event, for example consent fees or solicitation fees.
	SolicitationFees *ActiveCurrencyAndAmount `xml:"SlctnFees,omitempty"`

	// Cash disbursement in lieu of a fractional quantity of, for example, equity.
	CashInLieuOfShare *ActiveCurrencyAndAmount `xml:"CshInLieuOfShr,omitempty"`

	// Amount of money distributed as the result of a capital gain.
	CapitalGain *ActiveCurrencyAndAmount `xml:"CptlGn,omitempty"`

	// Amount of money representing a coupon payment.
	InterestAmount *ActiveCurrencyAndAmount `xml:"IntrstAmt,omitempty"`

	// Amount of money resulting from a market claim.
	MarketClaimAmount *ActiveCurrencyAndAmount `xml:"MktClmAmt,omitempty"`

	// (Unique to France) Amount due to a buyer of securities dealt prior to ex date which may be subject to different rate of taxation.
	IndemnityAmount *ActiveCurrencyAndAmount `xml:"IndmntyAmt,omitempty"`

	// Amount of money that the borrower pays to the lender as a compensation. It does not entitle the lender to reclaim any tax credit and is sometimes treated differently by the local tax authorities of the lender. Also covers compensation/indemnity of missed dividend concerning early/late settlements if applicable to a market.
	ManufacturedDividendPaymentAmount *ActiveCurrencyAndAmount `xml:"ManfctrdDvddPmtAmt,omitempty"`

	// Amount of money reinvested in additional securities.
	ReinvestmentAmount *ActiveCurrencyAndAmount `xml:"RinvstmtAmt,omitempty"`

	// Amount resulting from a fully franked dividend paid by a company; amount includes tax credit for companies that have made sufficient tax payments during the fiscal period.
	FullyFrankedAmount *ActiveCurrencyAndAmount `xml:"FullyFrnkdAmt,omitempty"`

	// Amount resulting from an unfranked dividend paid by a company; the amount does not include tax credit and is subject to withholding tax.
	UnfrankedAmount *ActiveCurrencyAndAmount `xml:"UfrnkdAmt,omitempty"`

	// Amount of money related to taxable income that cannot be categorised.
	SundryOrOtherAmount *ActiveCurrencyAndAmount `xml:"SndryOrOthrAmt,omitempty"`

	// Amount of money that has not been subject to taxation.
	TaxFreeAmount *ActiveCurrencyAndAmount `xml:"TaxFreeAmt,omitempty"`

	// Amount of income eligible for deferred taxation.
	TaxDeferredAmount *ActiveCurrencyAndAmount `xml:"TaxDfrrdAmt,omitempty"`

	// Tax on value added.
	ValueAddedTaxAmount *ActiveCurrencyAndAmount `xml:"ValAddedTaxAmt,omitempty"`

	// Amount of stamp duty.
	StampDutyAmount *ActiveCurrencyAndAmount `xml:"StmpDtyAmt,omitempty"`

	// Amount that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimAmount *ActiveCurrencyAndAmount `xml:"TaxRclmAmt,omitempty"`

	// Amount of taxes that have been previously paid in relation to the taxable event.
	TaxCreditAmount *ActiveCurrencyAndAmount `xml:"TaxCdtAmt,omitempty"`

	// Amount of additional taxes that cannot be categorised.
	AdditionalTaxAmount *ActiveCurrencyAndAmount `xml:"AddtlTaxAmt,omitempty"`

	// Amount of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgTaxAmt,omitempty"`

	// Amount of money withheld by the jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTaxAmount *ActiveCurrencyAndAmount `xml:"ScndLvlTaxAmt,omitempty"`

	// Amount of fiscal tax to apply.
	FiscalStampAmount *ActiveCurrencyAndAmount `xml:"FsclStmpAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *ActiveCurrencyAndAmount `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of paying/sub-paying agent commission.
	PayingAgentCommissionAmount *ActiveCurrencyAndAmount `xml:"PngAgtComssnAmt,omitempty"`

	// Local broker's commission.
	LocalBrokerCommissionAmount *ActiveCurrencyAndAmount `xml:"LclBrkrComssnAmt,omitempty"`

	// Amount of money charged by a regulatory authority, for example, Securities and Exchange fees.
	RegulatoryFeesAmount *ActiveCurrencyAndAmount `xml:"RgltryFeesAmt,omitempty"`

	// All costs related to the physical delivery of documents such as stamps, postage, carrier fees, insurances or messenger services.
	ShippingFeesAmount *ActiveCurrencyAndAmount `xml:"ShppgFeesAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesAmount *ActiveCurrencyAndAmount `xml:"ChrgsAmt,omitempty"`

	// Indicates cash retained from previous dividend.
	CashAmountBroughtForward *ActiveCurrencyAndAmount `xml:"CshAmtBrghtFwd,omitempty"`

	// Indicates the balance carried forward to next dividend.
	CashAmountCarriedForward *ActiveCurrencyAndAmount `xml:"CshAmtCrrdFwd,omitempty"`

	// Amount of cash that would have been payable if the dividend had been taken in the form of cash rather than shares.
	NotionalDividendPayableAmount *ActiveCurrencyAndAmount `xml:"NtnlDvddPyblAmt,omitempty"`

	// Where new securities are issued in lieu of a cash dividend, the notional tax is the tax on the amount of cash that would have been paid.
	NotionalTaxAmount *ActiveCurrencyAndAmount `xml:"NtnlTaxAmt,omitempty"`

	// Amount of money paid by the Tax Authorities in addition to the payment of the tax refund itself.
	TaxArrearsAmount *ActiveCurrencyAndAmount `xml:"TaxArrearsAmt,omitempty"`

	// Posting/settlement Amount in its original currency when conversion from/into another currency has occurred.
	OriginalAmount *ActiveCurrencyAndAmount `xml:"OrgnlAmt,omitempty"`

	// Amount of money representing a distribution of a bond's principal, for example, repayment of outstanding debt.
	PrincipalOrCorpus *ActiveCurrencyAndAmount `xml:"PrncplOrCrps,omitempty"`

	// Amount of money (not interest) in addition to the principal at the redemption of a bond.
	RedemptionPremiumAmount *ActiveCurrencyAndAmount `xml:"RedPrmAmt,omitempty"`

	// Amount relating to the underlying security for which income is distributed.
	IncomePortion *ActiveCurrencyAndAmount `xml:"IncmPrtn,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *ActiveCurrencyAndAmount `xml:"StockXchgTax,omitempty"`

	// Total amount of tax withheld at source in conformance with the EU Savings Directive.
	EUTaxRetentionAmount *ActiveCurrencyAndAmount `xml:"EUTaxRtntnAmt,omitempty"`

	// Amount of interest that has been accrued in between coupon payment periods
	AccruedInterestAmount *ActiveCurrencyAndAmount `xml:"AcrdIntrstAmt,omitempty"`

	// Portion of the fund distribution amount which represents the average accrued income included in the purchase price for units bought during the account period.
	EqualisationAmount *ActiveCurrencyAndAmount `xml:"EqulstnAmt,omitempty"`

	// FATCA (Foreign Account Tax Compliance Act) related tax amount.
	FATCATaxAmount *ActiveCurrencyAndAmount `xml:"FATCATaxAmt,omitempty"`

	// Amount of tax related income subject to NRA (Non Resident Alien).
	NRATaxAmount *ActiveCurrencyAndAmount `xml:"NRATaxAmt,omitempty"`

	// Amount of tax related to back up withholding.
	BackUpWithholdingTaxAmount *ActiveCurrencyAndAmount `xml:"BckUpWhldgTaxAmt,omitempty"`

	// Amount of overall tax withheld at source by fund managers prior to considering the tax obligation of each unit holder.
	TaxOnIncomeAmount *ActiveCurrencyAndAmount `xml:"TaxOnIncmAmt,omitempty"`

	// Amount of Transaction tax.
	TransactionTax *ActiveCurrencyAndAmount `xml:"TxTax,omitempty"`
}

func (c *CorporateActionAmounts37) SetPostingAmount(value, currency string) {
	c.PostingAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetGrossCashAmount(value, currency string) {
	c.GrossCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetNetCashAmount(value, currency string) {
	c.NetCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetSolicitationFees(value, currency string) {
	c.SolicitationFees = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetCashInLieuOfShare(value, currency string) {
	c.CashInLieuOfShare = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetCapitalGain(value, currency string) {
	c.CapitalGain = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetInterestAmount(value, currency string) {
	c.InterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetMarketClaimAmount(value, currency string) {
	c.MarketClaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetIndemnityAmount(value, currency string) {
	c.IndemnityAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetManufacturedDividendPaymentAmount(value, currency string) {
	c.ManufacturedDividendPaymentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetReinvestmentAmount(value, currency string) {
	c.ReinvestmentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetFullyFrankedAmount(value, currency string) {
	c.FullyFrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetUnfrankedAmount(value, currency string) {
	c.UnfrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetSundryOrOtherAmount(value, currency string) {
	c.SundryOrOtherAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetTaxFreeAmount(value, currency string) {
	c.TaxFreeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetTaxDeferredAmount(value, currency string) {
	c.TaxDeferredAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetValueAddedTaxAmount(value, currency string) {
	c.ValueAddedTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetStampDutyAmount(value, currency string) {
	c.StampDutyAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetTaxReclaimAmount(value, currency string) {
	c.TaxReclaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetTaxCreditAmount(value, currency string) {
	c.TaxCreditAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetAdditionalTaxAmount(value, currency string) {
	c.AdditionalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetWithholdingTaxAmount(value, currency string) {
	c.WithholdingTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetSecondLevelTaxAmount(value, currency string) {
	c.SecondLevelTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetFiscalStampAmount(value, currency string) {
	c.FiscalStampAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetExecutingBrokerAmount(value, currency string) {
	c.ExecutingBrokerAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetPayingAgentCommissionAmount(value, currency string) {
	c.PayingAgentCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetLocalBrokerCommissionAmount(value, currency string) {
	c.LocalBrokerCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetRegulatoryFeesAmount(value, currency string) {
	c.RegulatoryFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetShippingFeesAmount(value, currency string) {
	c.ShippingFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetChargesAmount(value, currency string) {
	c.ChargesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetCashAmountBroughtForward(value, currency string) {
	c.CashAmountBroughtForward = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetCashAmountCarriedForward(value, currency string) {
	c.CashAmountCarriedForward = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetNotionalDividendPayableAmount(value, currency string) {
	c.NotionalDividendPayableAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetNotionalTaxAmount(value, currency string) {
	c.NotionalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetTaxArrearsAmount(value, currency string) {
	c.TaxArrearsAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetPrincipalOrCorpus(value, currency string) {
	c.PrincipalOrCorpus = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetRedemptionPremiumAmount(value, currency string) {
	c.RedemptionPremiumAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetIncomePortion(value, currency string) {
	c.IncomePortion = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetStockExchangeTax(value, currency string) {
	c.StockExchangeTax = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetEUTaxRetentionAmount(value, currency string) {
	c.EUTaxRetentionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetAccruedInterestAmount(value, currency string) {
	c.AccruedInterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetEqualisationAmount(value, currency string) {
	c.EqualisationAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetFATCATaxAmount(value, currency string) {
	c.FATCATaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetNRATaxAmount(value, currency string) {
	c.NRATaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetBackUpWithholdingTaxAmount(value, currency string) {
	c.BackUpWithholdingTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetTaxOnIncomeAmount(value, currency string) {
	c.TaxOnIncomeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts37) SetTransactionTax(value, currency string) {
	c.TransactionTax = NewActiveCurrencyAndAmount(value, currency)
}

// Specifies amounts in the framework of a corporate action event.
type CorporateActionAmounts38 struct {

	// Amount of money before any deductions and allowances have been made.
	GrossCashAmount *ActiveCurrencyAndAmount `xml:"GrssCshAmt,omitempty"`

	// Amount of money after deductions and allowances have been made, if any, that is, the total amount +/- charges/fees.
	NetCashAmount *ActiveCurrencyAndAmount `xml:"NetCshAmt,omitempty"`

	// Cash premium made available if the securities holder consents or participates to an event, for example consent fees or solicitation fees.
	SolicitationFees *ActiveCurrencyAndAmount `xml:"SlctnFees,omitempty"`

	// Cash disbursement in lieu of a fractional quantity of, for example, equity.
	CashInLieuOfShare *ActiveCurrencyAndAmount `xml:"CshInLieuOfShr,omitempty"`

	// Amount of money distributed as the result of a capital gain.
	CapitalGain *ActiveCurrencyAndAmount `xml:"CptlGn,omitempty"`

	// Amount of money representing a coupon payment.
	InterestAmount *ActiveCurrencyAndAmount `xml:"IntrstAmt,omitempty"`

	// Amount of money resulting from a market claim.
	MarketClaimAmount *ActiveCurrencyAndAmount `xml:"MktClmAmt,omitempty"`

	// (Unique to France) Amount due to a buyer of securities dealt prior to ex date which may be subject to different rate of taxation.
	IndemnityAmount *ActiveCurrencyAndAmount `xml:"IndmntyAmt,omitempty"`

	// Amount of money that the borrower pays to the lender as a compensation. It does not entitle the lender to reclaim any tax credit and is sometimes treated differently by the local tax authorities of the lender. Also covers compensation/indemnity of missed dividend concerning early/late settlements if applicable to a market.
	ManufacturedDividendPaymentAmount *ActiveCurrencyAndAmount `xml:"ManfctrdDvddPmtAmt,omitempty"`

	// Amount of money reinvested in additional securities.
	ReinvestmentAmount *ActiveCurrencyAndAmount `xml:"RinvstmtAmt,omitempty"`

	// Amount resulting from a fully franked dividend paid by a company; amount includes tax credit for companies that have made sufficient tax payments during the fiscal period.
	FullyFrankedAmount *ActiveCurrencyAndAmount `xml:"FullyFrnkdAmt,omitempty"`

	// Amount resulting from an unfranked dividend paid by a company; the amount does not include tax credit and is subject to withholding tax.
	UnfrankedAmount *ActiveCurrencyAndAmount `xml:"UfrnkdAmt,omitempty"`

	// Amount of money related to taxable income that cannot be categorised.
	SundryOrOtherAmount *ActiveCurrencyAndAmount `xml:"SndryOrOthrAmt,omitempty"`

	// Amount of money that has not been subject to taxation.
	TaxFreeAmount *ActiveCurrencyAndAmount `xml:"TaxFreeAmt,omitempty"`

	// Amount of income eligible for deferred taxation.
	TaxDeferredAmount *ActiveCurrencyAndAmount `xml:"TaxDfrrdAmt,omitempty"`

	// Amount of value added tax.
	ValueAddedTaxAmount *ActiveCurrencyAndAmount `xml:"ValAddedTaxAmt,omitempty"`

	// Amount of stamp duty.
	StampDutyAmount *ActiveCurrencyAndAmount `xml:"StmpDtyAmt,omitempty"`

	// Amount that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimAmount *ActiveCurrencyAndAmount `xml:"TaxRclmAmt,omitempty"`

	// Amount of taxes that have been previously paid in relation to the taxable event.
	TaxCreditAmount *ActiveCurrencyAndAmount `xml:"TaxCdtAmt,omitempty"`

	// Amount of additional taxes that cannot be categorised.
	AdditionalTaxAmount *ActiveCurrencyAndAmount `xml:"AddtlTaxAmt,omitempty"`

	// Amount of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgTaxAmt,omitempty"`

	// Amount of money withheld by the jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTaxAmount *ActiveCurrencyAndAmount `xml:"ScndLvlTaxAmt,omitempty"`

	// Amount of fiscal tax to apply.
	FiscalStampAmount *ActiveCurrencyAndAmount `xml:"FsclStmpAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *ActiveCurrencyAndAmount `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of paying/sub-paying agent commission.
	PayingAgentCommissionAmount *ActiveCurrencyAndAmount `xml:"PngAgtComssnAmt,omitempty"`

	// Local broker's commission.
	LocalBrokerCommissionAmount *ActiveCurrencyAndAmount `xml:"LclBrkrComssnAmt,omitempty"`

	// Amount of money charged by a regulatory authority, for example, securities and exchange fees.
	RegulatoryFeesAmount *ActiveCurrencyAndAmount `xml:"RgltryFeesAmt,omitempty"`

	// All costs related to the physical delivery of documents such as stamps, postage, carrier fees, insurances or messenger services.
	ShippingFeesAmount *ActiveCurrencyAndAmount `xml:"ShppgFeesAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesAmount *ActiveCurrencyAndAmount `xml:"ChrgsAmt,omitempty"`

	// Cash amount based on terms of corporate action event and balance of underlying securities, entitled to/from account owner (which may be positive or negative).
	EntitledAmount *ActiveCurrencyAndAmount `xml:"EntitldAmt,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalAmount *ActiveCurrencyAndAmount `xml:"OrgnlAmt,omitempty"`

	// Amount of interest that has been accrued in between coupon payment periods
	AccruedInterestAmount *ActiveCurrencyAndAmount `xml:"AcrdIntrstAmt,omitempty"`

	// Amount relating to the underlying security for which income is distributed.
	IncomePortion *ActiveCurrencyAndAmount `xml:"IncmPrtn,omitempty"`

	// Portion of the fund distribution amount which represents the average accrued income included in the purchase price for units bought during the account period.
	EqualisationAmount *ActiveCurrencyAndAmount `xml:"EqulstnAmt,omitempty"`

	// FATCA (Foreign Account Tax Compliance Act) related tax amount.
	FATCATaxAmount *ActiveCurrencyAndAmount `xml:"FATCATaxAmt,omitempty"`

	// Amount of tax related income subject to NRA (Non Resident Alien).
	NRATaxAmount *ActiveCurrencyAndAmount `xml:"NRATaxAmt,omitempty"`

	// Amount of tax related to back up withholding.
	BackUpWithholdingTaxAmount *ActiveCurrencyAndAmount `xml:"BckUpWhldgTaxAmt,omitempty"`

	// Amount of overall tax withheld at source by fund managers prior to considering the tax obligation of each unit holder.
	TaxOnIncomeAmount *ActiveCurrencyAndAmount `xml:"TaxOnIncmAmt,omitempty"`

	// Amount of Transaction tax.
	TransactionTax *ActiveCurrencyAndAmount `xml:"TxTax,omitempty"`
}

func (c *CorporateActionAmounts38) SetGrossCashAmount(value, currency string) {
	c.GrossCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetNetCashAmount(value, currency string) {
	c.NetCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetSolicitationFees(value, currency string) {
	c.SolicitationFees = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetCashInLieuOfShare(value, currency string) {
	c.CashInLieuOfShare = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetCapitalGain(value, currency string) {
	c.CapitalGain = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetInterestAmount(value, currency string) {
	c.InterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetMarketClaimAmount(value, currency string) {
	c.MarketClaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetIndemnityAmount(value, currency string) {
	c.IndemnityAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetManufacturedDividendPaymentAmount(value, currency string) {
	c.ManufacturedDividendPaymentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetReinvestmentAmount(value, currency string) {
	c.ReinvestmentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetFullyFrankedAmount(value, currency string) {
	c.FullyFrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetUnfrankedAmount(value, currency string) {
	c.UnfrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetSundryOrOtherAmount(value, currency string) {
	c.SundryOrOtherAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetTaxFreeAmount(value, currency string) {
	c.TaxFreeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetTaxDeferredAmount(value, currency string) {
	c.TaxDeferredAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetValueAddedTaxAmount(value, currency string) {
	c.ValueAddedTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetStampDutyAmount(value, currency string) {
	c.StampDutyAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetTaxReclaimAmount(value, currency string) {
	c.TaxReclaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetTaxCreditAmount(value, currency string) {
	c.TaxCreditAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetAdditionalTaxAmount(value, currency string) {
	c.AdditionalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetWithholdingTaxAmount(value, currency string) {
	c.WithholdingTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetSecondLevelTaxAmount(value, currency string) {
	c.SecondLevelTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetFiscalStampAmount(value, currency string) {
	c.FiscalStampAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetExecutingBrokerAmount(value, currency string) {
	c.ExecutingBrokerAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetPayingAgentCommissionAmount(value, currency string) {
	c.PayingAgentCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetLocalBrokerCommissionAmount(value, currency string) {
	c.LocalBrokerCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetRegulatoryFeesAmount(value, currency string) {
	c.RegulatoryFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetShippingFeesAmount(value, currency string) {
	c.ShippingFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetChargesAmount(value, currency string) {
	c.ChargesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetEntitledAmount(value, currency string) {
	c.EntitledAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetAccruedInterestAmount(value, currency string) {
	c.AccruedInterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetIncomePortion(value, currency string) {
	c.IncomePortion = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetEqualisationAmount(value, currency string) {
	c.EqualisationAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetFATCATaxAmount(value, currency string) {
	c.FATCATaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetNRATaxAmount(value, currency string) {
	c.NRATaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetBackUpWithholdingTaxAmount(value, currency string) {
	c.BackUpWithholdingTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetTaxOnIncomeAmount(value, currency string) {
	c.TaxOnIncomeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts38) SetTransactionTax(value, currency string) {
	c.TransactionTax = NewActiveCurrencyAndAmount(value, currency)
}

// Specifies amounts in the framework of a corporate action event.
type CorporateActionAmounts39 struct {

	// Amount of money that is to be/was posted to the account.
	PostingAmount *RestrictedFINActiveCurrencyAndAmount `xml:"PstngAmt"`

	// Amount of money before any deductions and allowances have been made.
	GrossCashAmount *RestrictedFINActiveCurrencyAndAmount `xml:"GrssCshAmt,omitempty"`

	// Amount of money after deductions and allowances have been made, if any, that is, the total amount +/- charges/fees.
	NetCashAmount *RestrictedFINActiveCurrencyAndAmount `xml:"NetCshAmt,omitempty"`

	// Cash premium made available if the securities holder consents or participates to an event, for example consent fees or solicitation fees.
	SolicitationFees *RestrictedFINActiveCurrencyAndAmount `xml:"SlctnFees,omitempty"`

	// Cash disbursement in lieu of a fractional quantity of, for example, equity.
	CashInLieuOfShare *RestrictedFINActiveCurrencyAndAmount `xml:"CshInLieuOfShr,omitempty"`

	// Amount of money distributed as the result of a capital gain.
	CapitalGain *RestrictedFINActiveCurrencyAndAmount `xml:"CptlGn,omitempty"`

	// Amount of money representing a coupon payment.
	InterestAmount *RestrictedFINActiveCurrencyAndAmount `xml:"IntrstAmt,omitempty"`

	// Amount of money resulting from a market claim.
	MarketClaimAmount *RestrictedFINActiveCurrencyAndAmount `xml:"MktClmAmt,omitempty"`

	// (Unique to France) Amount due to a buyer of securities dealt prior to ex date which may be subject to different rate of taxation.
	IndemnityAmount *RestrictedFINActiveCurrencyAndAmount `xml:"IndmntyAmt,omitempty"`

	// Amount of money that the borrower pays to the lender as a compensation. It does not entitle the lender to reclaim any tax credit and is sometimes treated differently by the local tax authorities of the lender. Also covers compensation/indemnity of missed dividend concerning early/late settlements if applicable to a market.
	ManufacturedDividendPaymentAmount *RestrictedFINActiveCurrencyAndAmount `xml:"ManfctrdDvddPmtAmt,omitempty"`

	// Amount of money reinvested in additional securities.
	ReinvestmentAmount *RestrictedFINActiveCurrencyAndAmount `xml:"RinvstmtAmt,omitempty"`

	// Amount resulting from a fully franked dividend paid by a company; amount includes tax credit for companies that have made sufficient tax payments during the fiscal period.
	FullyFrankedAmount *RestrictedFINActiveCurrencyAndAmount `xml:"FullyFrnkdAmt,omitempty"`

	// Amount resulting from an unfranked dividend paid by a company; the amount does not include tax credit and is subject to withholding tax.
	UnfrankedAmount *RestrictedFINActiveCurrencyAndAmount `xml:"UfrnkdAmt,omitempty"`

	// Amount of money related to taxable income that cannot be categorised.
	SundryOrOtherAmount *RestrictedFINActiveCurrencyAndAmount `xml:"SndryOrOthrAmt,omitempty"`

	// Amount of money that has not been subject to taxation.
	TaxFreeAmount *RestrictedFINActiveCurrencyAndAmount `xml:"TaxFreeAmt,omitempty"`

	// Amount of income eligible for deferred taxation.
	TaxDeferredAmount *RestrictedFINActiveCurrencyAndAmount `xml:"TaxDfrrdAmt,omitempty"`

	// Tax on value added.
	ValueAddedTaxAmount *RestrictedFINActiveCurrencyAndAmount `xml:"ValAddedTaxAmt,omitempty"`

	// Amount of stamp duty.
	StampDutyAmount *RestrictedFINActiveCurrencyAndAmount `xml:"StmpDtyAmt,omitempty"`

	// Amount that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimAmount *RestrictedFINActiveCurrencyAndAmount `xml:"TaxRclmAmt,omitempty"`

	// Amount of taxes that have been previously paid in relation to the taxable event.
	TaxCreditAmount *RestrictedFINActiveCurrencyAndAmount `xml:"TaxCdtAmt,omitempty"`

	// Amount of additional taxes that cannot be categorised.
	AdditionalTaxAmount *RestrictedFINActiveCurrencyAndAmount `xml:"AddtlTaxAmt,omitempty"`

	// Amount of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxAmount *RestrictedFINActiveCurrencyAndAmount `xml:"WhldgTaxAmt,omitempty"`

	// Amount of money withheld by the jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTaxAmount *RestrictedFINActiveCurrencyAndAmount `xml:"ScndLvlTaxAmt,omitempty"`

	// Amount of fiscal tax to apply.
	FiscalStampAmount *RestrictedFINActiveCurrencyAndAmount `xml:"FsclStmpAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *RestrictedFINActiveCurrencyAndAmount `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of paying/sub-paying agent commission.
	PayingAgentCommissionAmount *RestrictedFINActiveCurrencyAndAmount `xml:"PngAgtComssnAmt,omitempty"`

	// Local broker's commission.
	LocalBrokerCommissionAmount *RestrictedFINActiveCurrencyAndAmount `xml:"LclBrkrComssnAmt,omitempty"`

	// Amount of money charged by a regulatory authority, for example, Securities and Exchange fees.
	RegulatoryFeesAmount *RestrictedFINActiveCurrencyAndAmount `xml:"RgltryFeesAmt,omitempty"`

	// All costs related to the physical delivery of documents such as stamps, postage, carrier fees, insurances or messenger services.
	ShippingFeesAmount *RestrictedFINActiveCurrencyAndAmount `xml:"ShppgFeesAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesAmount *RestrictedFINActiveCurrencyAndAmount `xml:"ChrgsAmt,omitempty"`

	// Indicates cash retained from previous dividend.
	CashAmountBroughtForward *RestrictedFINActiveCurrencyAndAmount `xml:"CshAmtBrghtFwd,omitempty"`

	// Indicates the balance carried forward to next dividend.
	CashAmountCarriedForward *RestrictedFINActiveCurrencyAndAmount `xml:"CshAmtCrrdFwd,omitempty"`

	// Amount of cash that would have been payable if the dividend had been taken in the form of cash rather than shares.
	NotionalDividendPayableAmount *RestrictedFINActiveCurrencyAndAmount `xml:"NtnlDvddPyblAmt,omitempty"`

	// Where new securities are issued in lieu of a cash dividend, the notional tax is the tax on the amount of cash that would have been paid.
	NotionalTaxAmount *RestrictedFINActiveCurrencyAndAmount `xml:"NtnlTaxAmt,omitempty"`

	// Amount of money paid by the Tax Authorities in addition to the payment of the tax refund itself.
	TaxArrearsAmount *RestrictedFINActiveCurrencyAndAmount `xml:"TaxArrearsAmt,omitempty"`

	// Posting/settlement Amount in its original currency when conversion from/into another currency has occurred.
	OriginalAmount *RestrictedFINActiveCurrencyAndAmount `xml:"OrgnlAmt,omitempty"`

	// Amount of money representing a distribution of a bond's principal, for example, repayment of outstanding debt.
	PrincipalOrCorpus *RestrictedFINActiveCurrencyAndAmount `xml:"PrncplOrCrps,omitempty"`

	// Amount of money (not interest) in addition to the principal at the redemption of a bond.
	RedemptionPremiumAmount *RestrictedFINActiveCurrencyAndAmount `xml:"RedPrmAmt,omitempty"`

	// Amount relating to the underlying security for which income is distributed.
	IncomePortion *RestrictedFINActiveCurrencyAndAmount `xml:"IncmPrtn,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *RestrictedFINActiveCurrencyAndAmount `xml:"StockXchgTax,omitempty"`

	// Total amount of tax withheld at source in conformance with the EU Savings Directive.
	EUTaxRetentionAmount *RestrictedFINActiveCurrencyAndAmount `xml:"EUTaxRtntnAmt,omitempty"`

	// Amount of interest that has been accrued in between coupon payment periods
	AccruedInterestAmount *RestrictedFINActiveCurrencyAndAmount `xml:"AcrdIntrstAmt,omitempty"`

	// Portion of the fund distribution amount which represents the average accrued income included in the purchase price for units bought during the account period.
	EqualisationAmount *RestrictedFINActiveCurrencyAndAmount `xml:"EqulstnAmt,omitempty"`

	// FATCA (Foreign Account Tax Compliance Act) related tax amount.
	FATCATaxAmount *RestrictedFINActiveCurrencyAndAmount `xml:"FATCATaxAmt,omitempty"`

	// Amount of tax related income subject to NRA (Non Resident Alien).
	NRATaxAmount *RestrictedFINActiveCurrencyAndAmount `xml:"NRATaxAmt,omitempty"`

	// Amount of tax related to back up withholding.
	BackUpWithholdingTaxAmount *RestrictedFINActiveCurrencyAndAmount `xml:"BckUpWhldgTaxAmt,omitempty"`

	// Amount of overall tax withheld at source by fund managers prior to considering the tax obligation of each unit holder.
	TaxOnIncomeAmount *RestrictedFINActiveCurrencyAndAmount `xml:"TaxOnIncmAmt,omitempty"`

	// Amount of Transaction tax.
	TransactionTax *RestrictedFINActiveCurrencyAndAmount `xml:"TxTax,omitempty"`
}

func (c *CorporateActionAmounts39) SetPostingAmount(value, currency string) {
	c.PostingAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetGrossCashAmount(value, currency string) {
	c.GrossCashAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetNetCashAmount(value, currency string) {
	c.NetCashAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetSolicitationFees(value, currency string) {
	c.SolicitationFees = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetCashInLieuOfShare(value, currency string) {
	c.CashInLieuOfShare = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetCapitalGain(value, currency string) {
	c.CapitalGain = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetInterestAmount(value, currency string) {
	c.InterestAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetMarketClaimAmount(value, currency string) {
	c.MarketClaimAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetIndemnityAmount(value, currency string) {
	c.IndemnityAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetManufacturedDividendPaymentAmount(value, currency string) {
	c.ManufacturedDividendPaymentAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetReinvestmentAmount(value, currency string) {
	c.ReinvestmentAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetFullyFrankedAmount(value, currency string) {
	c.FullyFrankedAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetUnfrankedAmount(value, currency string) {
	c.UnfrankedAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetSundryOrOtherAmount(value, currency string) {
	c.SundryOrOtherAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetTaxFreeAmount(value, currency string) {
	c.TaxFreeAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetTaxDeferredAmount(value, currency string) {
	c.TaxDeferredAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetValueAddedTaxAmount(value, currency string) {
	c.ValueAddedTaxAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetStampDutyAmount(value, currency string) {
	c.StampDutyAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetTaxReclaimAmount(value, currency string) {
	c.TaxReclaimAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetTaxCreditAmount(value, currency string) {
	c.TaxCreditAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetAdditionalTaxAmount(value, currency string) {
	c.AdditionalTaxAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetWithholdingTaxAmount(value, currency string) {
	c.WithholdingTaxAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetSecondLevelTaxAmount(value, currency string) {
	c.SecondLevelTaxAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetFiscalStampAmount(value, currency string) {
	c.FiscalStampAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetExecutingBrokerAmount(value, currency string) {
	c.ExecutingBrokerAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetPayingAgentCommissionAmount(value, currency string) {
	c.PayingAgentCommissionAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetLocalBrokerCommissionAmount(value, currency string) {
	c.LocalBrokerCommissionAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetRegulatoryFeesAmount(value, currency string) {
	c.RegulatoryFeesAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetShippingFeesAmount(value, currency string) {
	c.ShippingFeesAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetChargesAmount(value, currency string) {
	c.ChargesAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetCashAmountBroughtForward(value, currency string) {
	c.CashAmountBroughtForward = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetCashAmountCarriedForward(value, currency string) {
	c.CashAmountCarriedForward = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetNotionalDividendPayableAmount(value, currency string) {
	c.NotionalDividendPayableAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetNotionalTaxAmount(value, currency string) {
	c.NotionalTaxAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetTaxArrearsAmount(value, currency string) {
	c.TaxArrearsAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetPrincipalOrCorpus(value, currency string) {
	c.PrincipalOrCorpus = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetRedemptionPremiumAmount(value, currency string) {
	c.RedemptionPremiumAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetIncomePortion(value, currency string) {
	c.IncomePortion = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetStockExchangeTax(value, currency string) {
	c.StockExchangeTax = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetEUTaxRetentionAmount(value, currency string) {
	c.EUTaxRetentionAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetAccruedInterestAmount(value, currency string) {
	c.AccruedInterestAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetEqualisationAmount(value, currency string) {
	c.EqualisationAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetFATCATaxAmount(value, currency string) {
	c.FATCATaxAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetNRATaxAmount(value, currency string) {
	c.NRATaxAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetBackUpWithholdingTaxAmount(value, currency string) {
	c.BackUpWithholdingTaxAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetTaxOnIncomeAmount(value, currency string) {
	c.TaxOnIncomeAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts39) SetTransactionTax(value, currency string) {
	c.TransactionTax = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

// Specifies amounts in the framework of a corporate action event.
type CorporateActionAmounts4 struct {

	// Amount of money before any deductions and allowances have been made.
	GrossCashAmount *ActiveCurrencyAndAmount `xml:"GrssCshAmt,omitempty"`

	// Amount of money after deductions and allowances have been made, if any, that is, the total amount +/- charges/fees.
	NetCashAmount *ActiveCurrencyAndAmount `xml:"NetCshAmt,omitempty"`

	// Amount of cash premium made available in order to encourage participation in the offer. Payment is made to a third party who has solicited an entity to take part in the offer.
	SolicitationFees *ActiveCurrencyAndAmount `xml:"SlctnFees,omitempty"`

	// Cash disbursement in lieu of a fractional quantity of, for example, equity.
	CashInLieuOfShare *ActiveCurrencyAndAmount `xml:"CshInLieuOfShr,omitempty"`

	// Amount of money distributed as the result of a capital gain.
	CapitalGain *ActiveCurrencyAndAmount `xml:"CptlGn,omitempty"`

	// Amount of money representing a coupon payment.
	InterestAmount *ActiveCurrencyAndAmount `xml:"IntrstAmt,omitempty"`

	// (Unique to France) Amount due to a buyer of securities dealt prior to ex date which may be subject to different rate of taxation.
	IndemnityAmount *ActiveCurrencyAndAmount `xml:"IndmntyAmt,omitempty"`

	// Amount of money that the borrower pays to the lender as a compensation. It does not entitle the lender to reclaim any tax credit and is sometimes treated differently by the local tax authorities of the lender.
	ManufacturedDividendAmount *ActiveCurrencyAndAmount `xml:"ManfctrdDvddAmt,omitempty"`

	// Amount of money reinvested in additional securities.
	ReinvestmentAmount *ActiveCurrencyAndAmount `xml:"RinvstmtAmt,omitempty"`

	// Amount resulting from a fully franked dividend paid by a company; amount includes tax credit for companies that have made sufficient tax payments during the fiscal period.
	FullyFrankedAmount *ActiveCurrencyAndAmount `xml:"FullyFrnkdAmt,omitempty"`

	// Amount resulting from an unfranked dividend paid by a company; the amount does not include tax credit and is subject to withholding tax.
	UnfrankedAmount *ActiveCurrencyAndAmount `xml:"UfrnkdAmt,omitempty"`

	// Amount of money related to taxable income that cannot be categorised.
	SundryOrOtherAmount *ActiveCurrencyAndAmount `xml:"SndryOrOthrAmt,omitempty"`

	// Cash premium made available if the securities holder consents or participates to an event, for example, consent fees or solicitation fees.
	CashIncentive *ActiveCurrencyAndAmount `xml:"CshIncntiv,omitempty"`

	// Amount of money that has not been subject to taxation.
	TaxFreeAmount *ActiveCurrencyAndAmount `xml:"TaxFreeAmt,omitempty"`

	// Amount of income eligible for deferred taxation.
	TaxDeferredAmount *ActiveCurrencyAndAmount `xml:"TaxDfrrdAmt,omitempty"`

	// Amount of value added tax.
	ValueAddedTaxAmount *ActiveCurrencyAndAmount `xml:"ValAddedTaxAmt,omitempty"`

	// Amount of stamp duty.
	StampDutyAmount *ActiveCurrencyAndAmount `xml:"StmpDtyAmt,omitempty"`

	// Amount that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimAmount *ActiveCurrencyAndAmount `xml:"TaxRclmAmt,omitempty"`

	// Amount of taxes that have been previously paid in relation to the taxable event.
	TaxCreditAmount *ActiveCurrencyAndAmount `xml:"TaxCdtAmt,omitempty"`

	// Amount of money withheld by the jurisdiction in which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgOfFrgnTaxAmt,omitempty"`

	// Amount of money withheld by the jurisdiction in which the account owner is located, for which relief at source and/or reclaim may be possible.
	WithholdingOfLocalTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgOfLclTaxAmt,omitempty"`

	// Amount of additional taxes that cannot be categorised.
	AdditionalTaxAmount *ActiveCurrencyAndAmount `xml:"AddtlTaxAmt,omitempty"`

	// Amount of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgTaxAmt,omitempty"`

	// Amount of fiscal tax to apply.
	FiscalStampAmount *ActiveCurrencyAndAmount `xml:"FsclStmpAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *ActiveCurrencyAndAmount `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of paying/sub-paying agent commission.
	PayingAgentCommissionAmount *ActiveCurrencyAndAmount `xml:"PngAgtComssnAmt,omitempty"`

	// Local broker's commission.
	LocalBrokerCommissionAmount *ActiveCurrencyAndAmount `xml:"LclBrkrComssnAmt,omitempty"`

	// Amount of money paid for delivery by regular post mail.
	PostageFeeAmount *ActiveCurrencyAndAmount `xml:"PstgFeeAmt,omitempty"`

	// Amount of money charged by a regulatory authority, for example, securities and exchange fees.
	RegulatoryFeesAmount *ActiveCurrencyAndAmount `xml:"RgltryFeesAmt,omitempty"`

	// Amount of money (including insurance) paid for delivery by carrier.
	ShippingFeesAmount *ActiveCurrencyAndAmount `xml:"ShppgFeesAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesAmount *ActiveCurrencyAndAmount `xml:"ChrgsAmt,omitempty"`

	// Cash amount based on terms of corporate action event and balance of underlying securities, entitled to/from account owner (which may be positive or negative).
	EntitledAmount *ActiveCurrencyAndAmount `xml:"EntitldAmt,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalAmount *ActiveCurrencyAndAmount `xml:"OrgnlAmt,omitempty"`

	// Amount of money representing a distribution of a bond's principal, for example, repayment of outstanding debt.
	PrincipalOrCorpus *ActiveCurrencyAndAmount `xml:"PrncplOrCrps,omitempty"`

	// Amount of money (not interest) in addition to the principal at the redemption of a bond.
	RedemptionPremiumAmount *ActiveCurrencyAndAmount `xml:"RedPrmAmt,omitempty"`

	// Amount relating to the underlying security for which income is distributed.
	IncomePortion *ActiveCurrencyAndAmount `xml:"IncmPrtn,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *ActiveCurrencyAndAmount `xml:"StockXchgTax,omitempty"`

	// Total amount of tax withheld at source in conformance with the EU Savings Directive.
	EUTaxRetentionAmount *ActiveCurrencyAndAmount `xml:"EUTaxRtntnAmt,omitempty"`

	// Amount of interest that has been accrued in between coupon payment periods
	AccruedInterestAmount *ActiveCurrencyAndAmount `xml:"AcrdIntrstAmt,omitempty"`
}

func (c *CorporateActionAmounts4) SetGrossCashAmount(value, currency string) {
	c.GrossCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetNetCashAmount(value, currency string) {
	c.NetCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetSolicitationFees(value, currency string) {
	c.SolicitationFees = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetCashInLieuOfShare(value, currency string) {
	c.CashInLieuOfShare = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetCapitalGain(value, currency string) {
	c.CapitalGain = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetInterestAmount(value, currency string) {
	c.InterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetIndemnityAmount(value, currency string) {
	c.IndemnityAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetManufacturedDividendAmount(value, currency string) {
	c.ManufacturedDividendAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetReinvestmentAmount(value, currency string) {
	c.ReinvestmentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetFullyFrankedAmount(value, currency string) {
	c.FullyFrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetUnfrankedAmount(value, currency string) {
	c.UnfrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetSundryOrOtherAmount(value, currency string) {
	c.SundryOrOtherAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetCashIncentive(value, currency string) {
	c.CashIncentive = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetTaxFreeAmount(value, currency string) {
	c.TaxFreeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetTaxDeferredAmount(value, currency string) {
	c.TaxDeferredAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetValueAddedTaxAmount(value, currency string) {
	c.ValueAddedTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetStampDutyAmount(value, currency string) {
	c.StampDutyAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetTaxReclaimAmount(value, currency string) {
	c.TaxReclaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetTaxCreditAmount(value, currency string) {
	c.TaxCreditAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetWithholdingOfForeignTaxAmount(value, currency string) {
	c.WithholdingOfForeignTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetWithholdingOfLocalTaxAmount(value, currency string) {
	c.WithholdingOfLocalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetAdditionalTaxAmount(value, currency string) {
	c.AdditionalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetWithholdingTaxAmount(value, currency string) {
	c.WithholdingTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetFiscalStampAmount(value, currency string) {
	c.FiscalStampAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetExecutingBrokerAmount(value, currency string) {
	c.ExecutingBrokerAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetPayingAgentCommissionAmount(value, currency string) {
	c.PayingAgentCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetLocalBrokerCommissionAmount(value, currency string) {
	c.LocalBrokerCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetPostageFeeAmount(value, currency string) {
	c.PostageFeeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetRegulatoryFeesAmount(value, currency string) {
	c.RegulatoryFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetShippingFeesAmount(value, currency string) {
	c.ShippingFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetChargesAmount(value, currency string) {
	c.ChargesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetEntitledAmount(value, currency string) {
	c.EntitledAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetPrincipalOrCorpus(value, currency string) {
	c.PrincipalOrCorpus = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetRedemptionPremiumAmount(value, currency string) {
	c.RedemptionPremiumAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetIncomePortion(value, currency string) {
	c.IncomePortion = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetStockExchangeTax(value, currency string) {
	c.StockExchangeTax = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetEUTaxRetentionAmount(value, currency string) {
	c.EUTaxRetentionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts4) SetAccruedInterestAmount(value, currency string) {
	c.AccruedInterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

// Specifies amounts in the framework of a corporate action event.
type CorporateActionAmounts40 struct {

	// Amount of money before any deductions and allowances have been made.
	GrossCashAmount *RestrictedFINActiveCurrencyAndAmount `xml:"GrssCshAmt,omitempty"`

	// Amount of money after deductions and allowances have been made, if any, that is, the total amount +/- charges/fees.
	NetCashAmount *RestrictedFINActiveCurrencyAndAmount `xml:"NetCshAmt,omitempty"`

	// Cash premium made available if the securities holder consents or participates to an event, for example consent fees or solicitation fees.
	SolicitationFees *RestrictedFINActiveCurrencyAndAmount `xml:"SlctnFees,omitempty"`

	// Cash disbursement in lieu of a fractional quantity of, for example, equity.
	CashInLieuOfShare *RestrictedFINActiveCurrencyAndAmount `xml:"CshInLieuOfShr,omitempty"`

	// Amount of money distributed as the result of a capital gain.
	CapitalGain *RestrictedFINActiveCurrencyAndAmount `xml:"CptlGn,omitempty"`

	// Amount of money representing a coupon payment.
	InterestAmount *RestrictedFINActiveCurrencyAndAmount `xml:"IntrstAmt,omitempty"`

	// Amount of money resulting from a market claim.
	MarketClaimAmount *RestrictedFINActiveCurrencyAndAmount `xml:"MktClmAmt,omitempty"`

	// (Unique to France) Amount due to a buyer of securities dealt prior to ex date which may be subject to different rate of taxation.
	IndemnityAmount *RestrictedFINActiveCurrencyAndAmount `xml:"IndmntyAmt,omitempty"`

	// Amount of money that the borrower pays to the lender as a compensation. It does not entitle the lender to reclaim any tax credit and is sometimes treated differently by the local tax authorities of the lender. Also covers compensation/indemnity of missed dividend concerning early/late settlements if applicable to a market.
	ManufacturedDividendPaymentAmount *RestrictedFINActiveCurrencyAndAmount `xml:"ManfctrdDvddPmtAmt,omitempty"`

	// Amount of money reinvested in additional securities.
	ReinvestmentAmount *RestrictedFINActiveCurrencyAndAmount `xml:"RinvstmtAmt,omitempty"`

	// Amount resulting from a fully franked dividend paid by a company; amount includes tax credit for companies that have made sufficient tax payments during the fiscal period.
	FullyFrankedAmount *RestrictedFINActiveCurrencyAndAmount `xml:"FullyFrnkdAmt,omitempty"`

	// Amount resulting from an unfranked dividend paid by a company; the amount does not include tax credit and is subject to withholding tax.
	UnfrankedAmount *RestrictedFINActiveCurrencyAndAmount `xml:"UfrnkdAmt,omitempty"`

	// Amount of money related to taxable income that cannot be categorised.
	SundryOrOtherAmount *RestrictedFINActiveCurrencyAndAmount `xml:"SndryOrOthrAmt,omitempty"`

	// Amount of money that has not been subject to taxation.
	TaxFreeAmount *RestrictedFINActiveCurrencyAndAmount `xml:"TaxFreeAmt,omitempty"`

	// Amount of income eligible for deferred taxation.
	TaxDeferredAmount *RestrictedFINActiveCurrencyAndAmount `xml:"TaxDfrrdAmt,omitempty"`

	// Amount of value added tax.
	ValueAddedTaxAmount *RestrictedFINActiveCurrencyAndAmount `xml:"ValAddedTaxAmt,omitempty"`

	// Amount of stamp duty.
	StampDutyAmount *RestrictedFINActiveCurrencyAndAmount `xml:"StmpDtyAmt,omitempty"`

	// Amount that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimAmount *RestrictedFINActiveCurrencyAndAmount `xml:"TaxRclmAmt,omitempty"`

	// Amount of taxes that have been previously paid in relation to the taxable event.
	TaxCreditAmount *RestrictedFINActiveCurrencyAndAmount `xml:"TaxCdtAmt,omitempty"`

	// Amount of additional taxes that cannot be categorised.
	AdditionalTaxAmount *RestrictedFINActiveCurrencyAndAmount `xml:"AddtlTaxAmt,omitempty"`

	// Amount of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxAmount *RestrictedFINActiveCurrencyAndAmount `xml:"WhldgTaxAmt,omitempty"`

	// Amount of money withheld by the jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTaxAmount *ActiveCurrencyAndAmount `xml:"ScndLvlTaxAmt,omitempty"`

	// Amount of fiscal tax to apply.
	FiscalStampAmount *RestrictedFINActiveCurrencyAndAmount `xml:"FsclStmpAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *RestrictedFINActiveCurrencyAndAmount `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of paying/sub-paying agent commission.
	PayingAgentCommissionAmount *RestrictedFINActiveCurrencyAndAmount `xml:"PngAgtComssnAmt,omitempty"`

	// Local broker's commission.
	LocalBrokerCommissionAmount *RestrictedFINActiveCurrencyAndAmount `xml:"LclBrkrComssnAmt,omitempty"`

	// Amount of money charged by a regulatory authority, for example, securities and exchange fees.
	RegulatoryFeesAmount *RestrictedFINActiveCurrencyAndAmount `xml:"RgltryFeesAmt,omitempty"`

	// All costs related to the physical delivery of documents such as stamps, postage, carrier fees, insurances or messenger services.
	ShippingFeesAmount *RestrictedFINActiveCurrencyAndAmount `xml:"ShppgFeesAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesAmount *RestrictedFINActiveCurrencyAndAmount `xml:"ChrgsAmt,omitempty"`

	// Cash amount based on terms of corporate action event and balance of underlying securities, entitled to/from account owner (which may be positive or negative).
	EntitledAmount *RestrictedFINActiveCurrencyAndAmount `xml:"EntitldAmt,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalAmount *RestrictedFINActiveCurrencyAndAmount `xml:"OrgnlAmt,omitempty"`

	// Amount of interest that has been accrued in between coupon payment periods
	AccruedInterestAmount *RestrictedFINActiveCurrencyAndAmount `xml:"AcrdIntrstAmt,omitempty"`

	// Amount relating to the underlying security for which income is distributed.
	IncomePortion *RestrictedFINActiveCurrencyAndAmount `xml:"IncmPrtn,omitempty"`

	// Portion of the fund distribution amount which represents the average accrued income included in the purchase price for units bought during the account period.
	EqualisationAmount *RestrictedFINActiveCurrencyAndAmount `xml:"EqulstnAmt,omitempty"`

	// FATCA (Foreign Account Tax Compliance Act) related tax amount.
	FATCATaxAmount *RestrictedFINActiveCurrencyAndAmount `xml:"FATCATaxAmt,omitempty"`

	// Amount of tax related income subject to NRA (Non Resident Alien).
	NRATaxAmount *RestrictedFINActiveCurrencyAndAmount `xml:"NRATaxAmt,omitempty"`

	// Amount of tax related to back up withholding.
	BackUpWithholdingTaxAmount *RestrictedFINActiveCurrencyAndAmount `xml:"BckUpWhldgTaxAmt,omitempty"`

	// Amount of overall tax withheld at source by fund managers prior to considering the tax obligation of each unit holder.
	TaxOnIncomeAmount *RestrictedFINActiveCurrencyAndAmount `xml:"TaxOnIncmAmt,omitempty"`

	// Amount of Transaction tax.
	TransactionTax *RestrictedFINActiveCurrencyAndAmount `xml:"TxTax,omitempty"`
}

func (c *CorporateActionAmounts40) SetGrossCashAmount(value, currency string) {
	c.GrossCashAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetNetCashAmount(value, currency string) {
	c.NetCashAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetSolicitationFees(value, currency string) {
	c.SolicitationFees = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetCashInLieuOfShare(value, currency string) {
	c.CashInLieuOfShare = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetCapitalGain(value, currency string) {
	c.CapitalGain = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetInterestAmount(value, currency string) {
	c.InterestAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetMarketClaimAmount(value, currency string) {
	c.MarketClaimAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetIndemnityAmount(value, currency string) {
	c.IndemnityAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetManufacturedDividendPaymentAmount(value, currency string) {
	c.ManufacturedDividendPaymentAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetReinvestmentAmount(value, currency string) {
	c.ReinvestmentAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetFullyFrankedAmount(value, currency string) {
	c.FullyFrankedAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetUnfrankedAmount(value, currency string) {
	c.UnfrankedAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetSundryOrOtherAmount(value, currency string) {
	c.SundryOrOtherAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetTaxFreeAmount(value, currency string) {
	c.TaxFreeAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetTaxDeferredAmount(value, currency string) {
	c.TaxDeferredAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetValueAddedTaxAmount(value, currency string) {
	c.ValueAddedTaxAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetStampDutyAmount(value, currency string) {
	c.StampDutyAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetTaxReclaimAmount(value, currency string) {
	c.TaxReclaimAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetTaxCreditAmount(value, currency string) {
	c.TaxCreditAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetAdditionalTaxAmount(value, currency string) {
	c.AdditionalTaxAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetWithholdingTaxAmount(value, currency string) {
	c.WithholdingTaxAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetSecondLevelTaxAmount(value, currency string) {
	c.SecondLevelTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetFiscalStampAmount(value, currency string) {
	c.FiscalStampAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetExecutingBrokerAmount(value, currency string) {
	c.ExecutingBrokerAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetPayingAgentCommissionAmount(value, currency string) {
	c.PayingAgentCommissionAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetLocalBrokerCommissionAmount(value, currency string) {
	c.LocalBrokerCommissionAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetRegulatoryFeesAmount(value, currency string) {
	c.RegulatoryFeesAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetShippingFeesAmount(value, currency string) {
	c.ShippingFeesAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetChargesAmount(value, currency string) {
	c.ChargesAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetEntitledAmount(value, currency string) {
	c.EntitledAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetAccruedInterestAmount(value, currency string) {
	c.AccruedInterestAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetIncomePortion(value, currency string) {
	c.IncomePortion = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetEqualisationAmount(value, currency string) {
	c.EqualisationAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetFATCATaxAmount(value, currency string) {
	c.FATCATaxAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetNRATaxAmount(value, currency string) {
	c.NRATaxAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetBackUpWithholdingTaxAmount(value, currency string) {
	c.BackUpWithholdingTaxAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetTaxOnIncomeAmount(value, currency string) {
	c.TaxOnIncomeAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts40) SetTransactionTax(value, currency string) {
	c.TransactionTax = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

// Specifies amounts in the framework of a corporate action event.
type CorporateActionAmounts41 struct {

	// Amount of money before any deductions and allowances have been made.
	GrossCashAmount *RestrictedFINActiveCurrencyAndAmount `xml:"GrssCshAmt,omitempty"`

	// Amount of money after deductions and allowances have been made, if any, that is, the total amount +/- charges/fees.
	NetCashAmount *RestrictedFINActiveCurrencyAndAmount `xml:"NetCshAmt,omitempty"`

	// Cash premium made available if the securities holder consents or participates to an event, for example consent fees or solicitation fees.
	SolicitationFees *RestrictedFINActiveCurrencyAndAmount `xml:"SlctnFees,omitempty"`

	// Cash disbursement in lieu of a fractional quantity of, for example, equity.
	CashInLieuOfShare *RestrictedFINActiveCurrencyAndAmount `xml:"CshInLieuOfShr,omitempty"`

	// Amount of money distributed as the result of a capital gain.
	CapitalGain *RestrictedFINActiveCurrencyAndAmount `xml:"CptlGn,omitempty"`

	// Amount of money representing a coupon payment.
	InterestAmount *RestrictedFINActiveCurrencyAndAmount `xml:"IntrstAmt,omitempty"`

	// (Unique to France) Amount due to a buyer of securities dealt prior to ex date which may be subject to different rate of taxation.
	IndemnityAmount *RestrictedFINActiveCurrencyAndAmount `xml:"IndmntyAmt,omitempty"`

	// Amount of money that the borrower pays to the lender as a compensation. It does not entitle the lender to reclaim any tax credit and is sometimes treated differently by the local tax authorities of the lender. Also covers compensation/indemnity of missed dividend concerning early/late settlements if applicable to a market.
	ManufacturedDividendPaymentAmount *RestrictedFINActiveCurrencyAndAmount `xml:"ManfctrdDvddPmtAmt,omitempty"`

	// Amount of money reinvested in additional securities.
	ReinvestmentAmount *RestrictedFINActiveCurrencyAndAmount `xml:"RinvstmtAmt,omitempty"`

	// Amount resulting from a fully franked dividend paid by a company; amount includes tax credit for companies that have made sufficient tax payments during the fiscal period.
	FullyFrankedAmount *RestrictedFINActiveCurrencyAndAmount `xml:"FullyFrnkdAmt,omitempty"`

	// Amount resulting from an unfranked dividend paid by a company; the amount does not include tax credit and is subject to withholding tax.
	UnfrankedAmount *RestrictedFINActiveCurrencyAndAmount `xml:"UfrnkdAmt,omitempty"`

	// Amount of money related to taxable income that cannot be categorised.
	SundryOrOtherAmount *RestrictedFINActiveCurrencyAndAmount `xml:"SndryOrOthrAmt,omitempty"`

	// Amount of money that has not been subject to taxation.
	TaxFreeAmount *RestrictedFINActiveCurrencyAndAmount `xml:"TaxFreeAmt,omitempty"`

	// Amount of income eligible for deferred taxation.
	TaxDeferredAmount *RestrictedFINActiveCurrencyAndAmount `xml:"TaxDfrrdAmt,omitempty"`

	// Amount of value added tax.
	ValueAddedTaxAmount *RestrictedFINActiveCurrencyAndAmount `xml:"ValAddedTaxAmt,omitempty"`

	// Amount of stamp duty.
	StampDutyAmount *RestrictedFINActiveCurrencyAndAmount `xml:"StmpDtyAmt,omitempty"`

	// Amount that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimAmount *RestrictedFINActiveCurrencyAndAmount `xml:"TaxRclmAmt,omitempty"`

	// Amount of taxes that have been previously paid in relation to the taxable event.
	TaxCreditAmount *RestrictedFINActiveCurrencyAndAmount `xml:"TaxCdtAmt,omitempty"`

	// Amount of additional taxes that cannot be categorised.
	AdditionalTaxAmount *RestrictedFINActiveCurrencyAndAmount `xml:"AddtlTaxAmt,omitempty"`

	// Amount of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxAmount *RestrictedFINActiveCurrencyAndAmount `xml:"WhldgTaxAmt,omitempty"`

	// Amount of money withheld by the jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTaxAmount *RestrictedFINActiveCurrencyAndAmount `xml:"ScndLvlTaxAmt,omitempty"`

	// Amount of fiscal tax to apply.
	FiscalStampAmount *RestrictedFINActiveCurrencyAndAmount `xml:"FsclStmpAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *RestrictedFINActiveCurrencyAndAmount `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of paying/sub-paying agent commission.
	PayingAgentCommissionAmount *RestrictedFINActiveCurrencyAndAmount `xml:"PngAgtComssnAmt,omitempty"`

	// Local broker's commission.
	LocalBrokerCommissionAmount *RestrictedFINActiveCurrencyAndAmount `xml:"LclBrkrComssnAmt,omitempty"`

	// Amount of money charged by a regulatory authority, for example, securities and exchange fees.
	RegulatoryFeesAmount *RestrictedFINActiveCurrencyAndAmount `xml:"RgltryFeesAmt,omitempty"`

	// All costs related to the physical delivery of documents such as stamps, postage, carrier fees, insurances or messenger services.
	ShippingFeesAmount *RestrictedFINActiveCurrencyAndAmount `xml:"ShppgFeesAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesAmount *RestrictedFINActiveCurrencyAndAmount `xml:"ChrgsAmt,omitempty"`

	// Cash amount based on terms of corporate action event and balance of underlying securities, entitled to/from account owner (which may be positive or negative).
	EntitledAmount *RestrictedFINActiveCurrencyAndAmount `xml:"EntitldAmt,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalAmount *RestrictedFINActiveCurrencyAndAmount `xml:"OrgnlAmt,omitempty"`

	// Amount of money representing a distribution of a bond's principal, for example, repayment of outstanding debt.
	PrincipalOrCorpus *RestrictedFINActiveCurrencyAndAmount `xml:"PrncplOrCrps,omitempty"`

	// Amount of money (not interest) in addition to the principal at the redemption of a bond.
	RedemptionPremiumAmount *RestrictedFINActiveCurrencyAndAmount `xml:"RedPrmAmt,omitempty"`

	// Amount relating to the underlying security for which income is distributed.
	IncomePortion *RestrictedFINActiveCurrencyAndAmount `xml:"IncmPrtn,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *RestrictedFINActiveCurrencyAndAmount `xml:"StockXchgTax,omitempty"`

	// Total amount of tax withheld at source in conformance with the EU Savings Directive.
	EUTaxRetentionAmount *RestrictedFINActiveCurrencyAndAmount `xml:"EUTaxRtntnAmt,omitempty"`

	// Amount of interest that has been accrued in between coupon payment periods
	AccruedInterestAmount *RestrictedFINActiveCurrencyAndAmount `xml:"AcrdIntrstAmt,omitempty"`

	// Portion of the fund distribution amount which represents the average accrued income included in the purchase price for units bought during the account period.
	EqualisationAmount *RestrictedFINActiveCurrencyAndAmount `xml:"EqulstnAmt,omitempty"`

	// FATCA (Foreign Account Tax Compliance Act) related tax amount.
	FATCATaxAmount *RestrictedFINActiveCurrencyAndAmount `xml:"FATCATaxAmt,omitempty"`

	// Amount of tax related income subject to NRA (Non Resident Alien).
	NRATaxAmount *RestrictedFINActiveCurrencyAndAmount `xml:"NRATaxAmt,omitempty"`

	// Amount of tax related to back up withholding.
	BackUpWithholdingTaxAmount *RestrictedFINActiveCurrencyAndAmount `xml:"BckUpWhldgTaxAmt,omitempty"`

	// Amount of overall tax withheld at source by fund managers prior to considering the tax obligation of each unit holder.
	TaxOnIncomeAmount *RestrictedFINActiveCurrencyAndAmount `xml:"TaxOnIncmAmt,omitempty"`

	// Amount of Transaction tax.
	TransactionTax *RestrictedFINActiveCurrencyAndAmount `xml:"TxTax,omitempty"`
}

func (c *CorporateActionAmounts41) SetGrossCashAmount(value, currency string) {
	c.GrossCashAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetNetCashAmount(value, currency string) {
	c.NetCashAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetSolicitationFees(value, currency string) {
	c.SolicitationFees = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetCashInLieuOfShare(value, currency string) {
	c.CashInLieuOfShare = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetCapitalGain(value, currency string) {
	c.CapitalGain = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetInterestAmount(value, currency string) {
	c.InterestAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetIndemnityAmount(value, currency string) {
	c.IndemnityAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetManufacturedDividendPaymentAmount(value, currency string) {
	c.ManufacturedDividendPaymentAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetReinvestmentAmount(value, currency string) {
	c.ReinvestmentAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetFullyFrankedAmount(value, currency string) {
	c.FullyFrankedAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetUnfrankedAmount(value, currency string) {
	c.UnfrankedAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetSundryOrOtherAmount(value, currency string) {
	c.SundryOrOtherAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetTaxFreeAmount(value, currency string) {
	c.TaxFreeAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetTaxDeferredAmount(value, currency string) {
	c.TaxDeferredAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetValueAddedTaxAmount(value, currency string) {
	c.ValueAddedTaxAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetStampDutyAmount(value, currency string) {
	c.StampDutyAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetTaxReclaimAmount(value, currency string) {
	c.TaxReclaimAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetTaxCreditAmount(value, currency string) {
	c.TaxCreditAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetAdditionalTaxAmount(value, currency string) {
	c.AdditionalTaxAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetWithholdingTaxAmount(value, currency string) {
	c.WithholdingTaxAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetSecondLevelTaxAmount(value, currency string) {
	c.SecondLevelTaxAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetFiscalStampAmount(value, currency string) {
	c.FiscalStampAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetExecutingBrokerAmount(value, currency string) {
	c.ExecutingBrokerAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetPayingAgentCommissionAmount(value, currency string) {
	c.PayingAgentCommissionAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetLocalBrokerCommissionAmount(value, currency string) {
	c.LocalBrokerCommissionAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetRegulatoryFeesAmount(value, currency string) {
	c.RegulatoryFeesAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetShippingFeesAmount(value, currency string) {
	c.ShippingFeesAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetChargesAmount(value, currency string) {
	c.ChargesAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetEntitledAmount(value, currency string) {
	c.EntitledAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetPrincipalOrCorpus(value, currency string) {
	c.PrincipalOrCorpus = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetRedemptionPremiumAmount(value, currency string) {
	c.RedemptionPremiumAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetIncomePortion(value, currency string) {
	c.IncomePortion = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetStockExchangeTax(value, currency string) {
	c.StockExchangeTax = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetEUTaxRetentionAmount(value, currency string) {
	c.EUTaxRetentionAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetAccruedInterestAmount(value, currency string) {
	c.AccruedInterestAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetEqualisationAmount(value, currency string) {
	c.EqualisationAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetFATCATaxAmount(value, currency string) {
	c.FATCATaxAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetNRATaxAmount(value, currency string) {
	c.NRATaxAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetBackUpWithholdingTaxAmount(value, currency string) {
	c.BackUpWithholdingTaxAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetTaxOnIncomeAmount(value, currency string) {
	c.TaxOnIncomeAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts41) SetTransactionTax(value, currency string) {
	c.TransactionTax = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

// Specifies amounts in the framework of a corporate action event.
type CorporateActionAmounts9 struct {

	// Amount of money before any deductions and allowances have been made.
	GrossCashAmount *ActiveCurrencyAndAmount `xml:"GrssCshAmt,omitempty"`

	// Amount of money after deductions and allowances have been made, if any, that is, the total amount +/- charges/fees.
	NetCashAmount *ActiveCurrencyAndAmount `xml:"NetCshAmt,omitempty"`

	// Amount of cash premium made available in order to encourage participation in the offer. Payment is made to a third party who has solicited an entity to take part in the offer.
	SolicitationFees *ActiveCurrencyAndAmount `xml:"SlctnFees,omitempty"`

	// Cash disbursement in lieu of a fractional quantity of, for example, equity.
	CashInLieuOfShare *ActiveCurrencyAndAmount `xml:"CshInLieuOfShr,omitempty"`

	// Amount of money distributed as the result of a capital gain.
	CapitalGain *ActiveCurrencyAndAmount `xml:"CptlGn,omitempty"`

	// Amount of money representing a coupon payment.
	InterestAmount *ActiveCurrencyAndAmount `xml:"IntrstAmt,omitempty"`

	// Amount of money resulting from a market claim.
	MarketClaimAmount *ActiveCurrencyAndAmount `xml:"MktClmAmt,omitempty"`

	// (Unique to France) Amount due to a buyer of securities dealt prior to ex date which may be subject to different rate of taxation.
	IndemnityAmount *ActiveCurrencyAndAmount `xml:"IndmntyAmt,omitempty"`

	// Amount of money that the borrower pays to the lender as a compensation. It does not entitle the lender to reclaim any tax credit and is sometimes treated differently by the local tax authorities of the lender. Also covers compensation/indemnity of missed dividend concerning early/late settlements if applicable to a market.
	ManufacturedDividendPaymentAmount *ActiveCurrencyAndAmount `xml:"ManfctrdDvddPmtAmt,omitempty"`

	// Amount of money reinvested in additional securities.
	ReinvestmentAmount *ActiveCurrencyAndAmount `xml:"RinvstmtAmt,omitempty"`

	// Amount resulting from a fully franked dividend paid by a company; amount includes tax credit for companies that have made sufficient tax payments during the fiscal period.
	FullyFrankedAmount *ActiveCurrencyAndAmount `xml:"FullyFrnkdAmt,omitempty"`

	// Amount resulting from an unfranked dividend paid by a company; the amount does not include tax credit and is subject to withholding tax.
	UnfrankedAmount *ActiveCurrencyAndAmount `xml:"UfrnkdAmt,omitempty"`

	// Amount of money related to taxable income that cannot be categorised.
	SundryOrOtherAmount *ActiveCurrencyAndAmount `xml:"SndryOrOthrAmt,omitempty"`

	// Cash premium made available if the securities holder consents or participates to an event, for example, consent fees or solicitation fees.
	CashIncentive *ActiveCurrencyAndAmount `xml:"CshIncntiv,omitempty"`

	// Amount of money that has not been subject to taxation.
	TaxFreeAmount *ActiveCurrencyAndAmount `xml:"TaxFreeAmt,omitempty"`

	// Amount of income eligible for deferred taxation.
	TaxDeferredAmount *ActiveCurrencyAndAmount `xml:"TaxDfrrdAmt,omitempty"`

	// Amount of value added tax.
	ValueAddedTaxAmount *ActiveCurrencyAndAmount `xml:"ValAddedTaxAmt,omitempty"`

	// Amount of stamp duty.
	StampDutyAmount *ActiveCurrencyAndAmount `xml:"StmpDtyAmt,omitempty"`

	// Amount that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimAmount *ActiveCurrencyAndAmount `xml:"TaxRclmAmt,omitempty"`

	// Amount of taxes that have been previously paid in relation to the taxable event.
	TaxCreditAmount *ActiveCurrencyAndAmount `xml:"TaxCdtAmt,omitempty"`

	// Amount of money withheld by the jurisdiction in which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgOfFrgnTaxAmt,omitempty"`

	// Amount of money withheld by the jurisdiction in which the account owner is located, for which relief at source and/or reclaim may be possible.
	WithholdingOfLocalTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgOfLclTaxAmt,omitempty"`

	// Amount of additional taxes that cannot be categorised.
	AdditionalTaxAmount *ActiveCurrencyAndAmount `xml:"AddtlTaxAmt,omitempty"`

	// Amount of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgTaxAmt,omitempty"`

	// Amount of fiscal tax to apply.
	FiscalStampAmount *ActiveCurrencyAndAmount `xml:"FsclStmpAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *ActiveCurrencyAndAmount `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of paying/sub-paying agent commission.
	PayingAgentCommissionAmount *ActiveCurrencyAndAmount `xml:"PngAgtComssnAmt,omitempty"`

	// Local broker's commission.
	LocalBrokerCommissionAmount *ActiveCurrencyAndAmount `xml:"LclBrkrComssnAmt,omitempty"`

	// Amount of money paid for delivery by regular post mail.
	PostageFeeAmount *ActiveCurrencyAndAmount `xml:"PstgFeeAmt,omitempty"`

	// Amount of money charged by a regulatory authority, for example, securities and exchange fees.
	RegulatoryFeesAmount *ActiveCurrencyAndAmount `xml:"RgltryFeesAmt,omitempty"`

	// Amount of money (including insurance) paid for delivery by carrier.
	ShippingFeesAmount *ActiveCurrencyAndAmount `xml:"ShppgFeesAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesAmount *ActiveCurrencyAndAmount `xml:"ChrgsAmt,omitempty"`

	// Cash amount based on terms of corporate action event and balance of underlying securities, entitled to/from account owner (which may be positive or negative).
	EntitledAmount *ActiveCurrencyAndAmount `xml:"EntitldAmt,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalAmount *ActiveCurrencyAndAmount `xml:"OrgnlAmt,omitempty"`

	// Amount of interest that has been accrued in between coupon payment periods
	AccruedInterestAmount *ActiveCurrencyAndAmount `xml:"AcrdIntrstAmt,omitempty"`

	// Amount relating to the underlying security for which income is distributed.
	IncomePortion *ActiveCurrencyAndAmount `xml:"IncmPrtn,omitempty"`
}

func (c *CorporateActionAmounts9) SetGrossCashAmount(value, currency string) {
	c.GrossCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetNetCashAmount(value, currency string) {
	c.NetCashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetSolicitationFees(value, currency string) {
	c.SolicitationFees = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetCashInLieuOfShare(value, currency string) {
	c.CashInLieuOfShare = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetCapitalGain(value, currency string) {
	c.CapitalGain = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetInterestAmount(value, currency string) {
	c.InterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetMarketClaimAmount(value, currency string) {
	c.MarketClaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetIndemnityAmount(value, currency string) {
	c.IndemnityAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetManufacturedDividendPaymentAmount(value, currency string) {
	c.ManufacturedDividendPaymentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetReinvestmentAmount(value, currency string) {
	c.ReinvestmentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetFullyFrankedAmount(value, currency string) {
	c.FullyFrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetUnfrankedAmount(value, currency string) {
	c.UnfrankedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetSundryOrOtherAmount(value, currency string) {
	c.SundryOrOtherAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetCashIncentive(value, currency string) {
	c.CashIncentive = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetTaxFreeAmount(value, currency string) {
	c.TaxFreeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetTaxDeferredAmount(value, currency string) {
	c.TaxDeferredAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetValueAddedTaxAmount(value, currency string) {
	c.ValueAddedTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetStampDutyAmount(value, currency string) {
	c.StampDutyAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetTaxReclaimAmount(value, currency string) {
	c.TaxReclaimAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetTaxCreditAmount(value, currency string) {
	c.TaxCreditAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetWithholdingOfForeignTaxAmount(value, currency string) {
	c.WithholdingOfForeignTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetWithholdingOfLocalTaxAmount(value, currency string) {
	c.WithholdingOfLocalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetAdditionalTaxAmount(value, currency string) {
	c.AdditionalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetWithholdingTaxAmount(value, currency string) {
	c.WithholdingTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetFiscalStampAmount(value, currency string) {
	c.FiscalStampAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetExecutingBrokerAmount(value, currency string) {
	c.ExecutingBrokerAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetPayingAgentCommissionAmount(value, currency string) {
	c.PayingAgentCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetLocalBrokerCommissionAmount(value, currency string) {
	c.LocalBrokerCommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetPostageFeeAmount(value, currency string) {
	c.PostageFeeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetRegulatoryFeesAmount(value, currency string) {
	c.RegulatoryFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetShippingFeesAmount(value, currency string) {
	c.ShippingFeesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetChargesAmount(value, currency string) {
	c.ChargesAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetEntitledAmount(value, currency string) {
	c.EntitledAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetAccruedInterestAmount(value, currency string) {
	c.AccruedInterestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CorporateActionAmounts9) SetIncomePortion(value, currency string) {
	c.IncomePortion = NewActiveCurrencyAndAmount(value, currency)
}
