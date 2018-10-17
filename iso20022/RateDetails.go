package iso20022

// Provides information about the rates related to securities movement.
type RateDetails10 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat14Choice `xml:"AddtlTax,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat7Choice `xml:"GrssDvddRate,omitempty"`

	// The actual interest rate used for the payment of the interest for the specified interest period.
	// Usage guideline: It is used to provide the applicable rate for the current payment, after all calculations have been performed, that is, application of period and method of interest computation.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat5Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Percentage of the gross dividend rate on which tax must be paid .
	TaxRelatedRate []*RateTypeAndAmountAndStatus6 `xml:"TaxRltdRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxRate *RateFormat6Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat14Choice `xml:"ChrgsFees,omitempty"`

	// Cash rate made available, as an incentive, in addition to the solicitation fee, in order to encourage early participation in an offer.
	EarlySolicitationFeeRate *SolicitationFeeRateFormat5Choice `xml:"EarlySlctnFeeRate,omitempty"`

	// Dividend is final.
	FinalDividendRate *RateAndAmountFormat15Choice `xml:"FnlDvddRate,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *RateFormat6Choice `xml:"FsclStmp,omitempty"`

	// Rate resulting from a fully franked dividend paid by a company; rate includes tax credit for companies that have made sufficient tax payments during fiscal period.
	FullyFrankedRate *RateAndAmountFormat14Choice `xml:"FullyFrnkdRate,omitempty"`

	// Cash rate made available in an event in order to encourage participation in the offer. As information, payment is made to a third party who has solicited an entity to take part in the offer.
	ThirdPartyIncentiveRate *RateFormat8Choice `xml:"ThrdPtyIncntivRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat9Choice `xml:"NetDvddRate,omitempty"`

	// Rate per share to which a non-resident is entitled.
	NonResidentRate *RateAndAmountFormat14Choice `xml:"NonResdtRate,omitempty"`

	// Dividend is provisional.
	ProvisionalDividendRate *RateAndAmountFormat15Choice `xml:"PrvsnlDvddRate,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *RateFormat6Choice `xml:"AplblRate,omitempty"`

	// Rate of the cash premium made available if the securities holder consents or participates to an event, for example consent fees or solicitation fee.
	SolicitationFeeRate *SolicitationFeeRateFormat5Choice `xml:"SlctnFeeRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat5Choice `xml:"TaxCdtRate,omitempty"`

	// Taxation applied on an amount clearly identified as an income.
	TaxOnIncome *RateFormat6Choice `xml:"TaxOnIncm,omitempty"`

	// Taxation applied on an amount clearly identified as capital profits, capital gains.
	TaxOnProfits *RateFormat6Choice `xml:"TaxOnPrfts,omitempty"`

	// Percentage of cash that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimRate *RateFormat6Choice `xml:"TaxRclmRate,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction in which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTax *RateAndAmountFormat14Choice `xml:"WhldgOfFrgnTax,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction in which the account owner is located, for which relief at source and/or reclaim may be possible.
	WithholdingOfLocalTax *RateAndAmountFormat14Choice `xml:"WhldgOfLclTax,omitempty"`

	// Portion of the fund distribution which represents the average accrued income included in the purchase price for units bought during the account period.
	EqualisationRate *RateAndAmountFormat15Choice `xml:"EqulstnRate,omitempty"`
}

func (r *RateDetails10) AddAdditionalTax() *RateAndAmountFormat14Choice {
	r.AdditionalTax = new(RateAndAmountFormat14Choice)
	return r.AdditionalTax
}

func (r *RateDetails10) AddGrossDividendRate() *GrossDividendRateFormat7Choice {
	newValue := new(GrossDividendRateFormat7Choice)
	r.GrossDividendRate = append(r.GrossDividendRate, newValue)
	return newValue
}

func (r *RateDetails10) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat5Choice {
	newValue := new(InterestRateUsedForPaymentFormat5Choice)
	r.InterestRateUsedForPayment = append(r.InterestRateUsedForPayment, newValue)
	return newValue
}

func (r *RateDetails10) AddTaxRelatedRate() *RateTypeAndAmountAndStatus6 {
	newValue := new(RateTypeAndAmountAndStatus6)
	r.TaxRelatedRate = append(r.TaxRelatedRate, newValue)
	return newValue
}

func (r *RateDetails10) AddWithholdingTaxRate() *RateFormat6Choice {
	r.WithholdingTaxRate = new(RateFormat6Choice)
	return r.WithholdingTaxRate
}

func (r *RateDetails10) AddChargesFees() *RateAndAmountFormat14Choice {
	r.ChargesFees = new(RateAndAmountFormat14Choice)
	return r.ChargesFees
}

func (r *RateDetails10) AddEarlySolicitationFeeRate() *SolicitationFeeRateFormat5Choice {
	r.EarlySolicitationFeeRate = new(SolicitationFeeRateFormat5Choice)
	return r.EarlySolicitationFeeRate
}

func (r *RateDetails10) AddFinalDividendRate() *RateAndAmountFormat15Choice {
	r.FinalDividendRate = new(RateAndAmountFormat15Choice)
	return r.FinalDividendRate
}

func (r *RateDetails10) AddFiscalStamp() *RateFormat6Choice {
	r.FiscalStamp = new(RateFormat6Choice)
	return r.FiscalStamp
}

func (r *RateDetails10) AddFullyFrankedRate() *RateAndAmountFormat14Choice {
	r.FullyFrankedRate = new(RateAndAmountFormat14Choice)
	return r.FullyFrankedRate
}

func (r *RateDetails10) AddThirdPartyIncentiveRate() *RateFormat8Choice {
	r.ThirdPartyIncentiveRate = new(RateFormat8Choice)
	return r.ThirdPartyIncentiveRate
}

func (r *RateDetails10) AddNetDividendRate() *NetDividendRateFormat9Choice {
	newValue := new(NetDividendRateFormat9Choice)
	r.NetDividendRate = append(r.NetDividendRate, newValue)
	return newValue
}

func (r *RateDetails10) AddNonResidentRate() *RateAndAmountFormat14Choice {
	r.NonResidentRate = new(RateAndAmountFormat14Choice)
	return r.NonResidentRate
}

func (r *RateDetails10) AddProvisionalDividendRate() *RateAndAmountFormat15Choice {
	r.ProvisionalDividendRate = new(RateAndAmountFormat15Choice)
	return r.ProvisionalDividendRate
}

func (r *RateDetails10) AddApplicableRate() *RateFormat6Choice {
	r.ApplicableRate = new(RateFormat6Choice)
	return r.ApplicableRate
}

func (r *RateDetails10) AddSolicitationFeeRate() *SolicitationFeeRateFormat5Choice {
	r.SolicitationFeeRate = new(SolicitationFeeRateFormat5Choice)
	return r.SolicitationFeeRate
}

func (r *RateDetails10) AddTaxCreditRate() *TaxCreditRateFormat5Choice {
	newValue := new(TaxCreditRateFormat5Choice)
	r.TaxCreditRate = append(r.TaxCreditRate, newValue)
	return newValue
}

func (r *RateDetails10) AddTaxOnIncome() *RateFormat6Choice {
	r.TaxOnIncome = new(RateFormat6Choice)
	return r.TaxOnIncome
}

func (r *RateDetails10) AddTaxOnProfits() *RateFormat6Choice {
	r.TaxOnProfits = new(RateFormat6Choice)
	return r.TaxOnProfits
}

func (r *RateDetails10) AddTaxReclaimRate() *RateFormat6Choice {
	r.TaxReclaimRate = new(RateFormat6Choice)
	return r.TaxReclaimRate
}

func (r *RateDetails10) AddWithholdingOfForeignTax() *RateAndAmountFormat14Choice {
	r.WithholdingOfForeignTax = new(RateAndAmountFormat14Choice)
	return r.WithholdingOfForeignTax
}

func (r *RateDetails10) AddWithholdingOfLocalTax() *RateAndAmountFormat14Choice {
	r.WithholdingOfLocalTax = new(RateAndAmountFormat14Choice)
	return r.WithholdingOfLocalTax
}

func (r *RateDetails10) AddEqualisationRate() *RateAndAmountFormat15Choice {
	r.EqualisationRate = new(RateAndAmountFormat15Choice)
	return r.EqualisationRate
}

// Provides information about the rates related to securities movement.
type RateDetails11 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat5Choice `xml:"AddtlTax,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat5Choice `xml:"ChrgsFees,omitempty"`

	// Dividend is final.
	FinalDividendRate *ActiveCurrencyAnd13DecimalAmount `xml:"FnlDvddRate,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *PercentageRate `xml:"FsclStmp,omitempty"`

	// Rate resulting from a fully franked dividend paid by a company; rate includes tax credit for companies that have made sufficient tax payments during fiscal period.
	FullyFrankedRate *RateAndAmountFormat5Choice `xml:"FullyFrnkdRate,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat8Choice `xml:"GrssDvddRate,omitempty"`

	// Cash rate made available, as an incentive, in addition to the solicitation fee, in order to encourage early participation in an offer.
	EarlySolicitationFeeRate *RateAndAmountFormat5Choice `xml:"EarlySlctnFeeRate,omitempty"`

	// Cash rate made available in an event in order to encourage participation in the offer. As information, Payment is made to a third party who has solicited an entity to take part in the offer.
	ThirdPartyIncentiveRate *RateAndAmountFormat5Choice `xml:"ThrdPtyIncntivRate,omitempty"`

	// Actual interest rate used for the payment of the interest for the specified interest period.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat2Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat10Choice `xml:"NetDvddRate,omitempty"`

	// Rate per share to which a non-resident is entitled.
	NonResidentRate *RateAndAmountFormat5Choice `xml:"NonResdtRate,omitempty"`

	// Dividend is provisional.
	ProvisionalDividendRate *ActiveCurrencyAnd13DecimalAmount `xml:"PrvsnlDvddRate,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *PercentageRate `xml:"AplblRate,omitempty"`

	// Rate of the cash premium made available if the securities holder consents or participates to an event, for example consent fees or solicitation fee.
	SolicitationFeeRate *RateAndAmountFormat5Choice `xml:"SlctnFeeRate,omitempty"`

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

	// Portion of the fund distribution which represents the average accrued income included in the purchase price for units bought during the account period.
	EqualisationRate *ActiveCurrencyAnd13DecimalAmount `xml:"EqulstnRate,omitempty"`
}

func (r *RateDetails11) AddAdditionalTax() *RateAndAmountFormat5Choice {
	r.AdditionalTax = new(RateAndAmountFormat5Choice)
	return r.AdditionalTax
}

func (r *RateDetails11) AddChargesFees() *RateAndAmountFormat5Choice {
	r.ChargesFees = new(RateAndAmountFormat5Choice)
	return r.ChargesFees
}

func (r *RateDetails11) SetFinalDividendRate(value, currency string) {
	r.FinalDividendRate = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateDetails11) SetFiscalStamp(value string) {
	r.FiscalStamp = (*PercentageRate)(&value)
}

func (r *RateDetails11) AddFullyFrankedRate() *RateAndAmountFormat5Choice {
	r.FullyFrankedRate = new(RateAndAmountFormat5Choice)
	return r.FullyFrankedRate
}

func (r *RateDetails11) AddGrossDividendRate() *GrossDividendRateFormat8Choice {
	newValue := new(GrossDividendRateFormat8Choice)
	r.GrossDividendRate = append(r.GrossDividendRate, newValue)
	return newValue
}

func (r *RateDetails11) AddEarlySolicitationFeeRate() *RateAndAmountFormat5Choice {
	r.EarlySolicitationFeeRate = new(RateAndAmountFormat5Choice)
	return r.EarlySolicitationFeeRate
}

func (r *RateDetails11) AddThirdPartyIncentiveRate() *RateAndAmountFormat5Choice {
	r.ThirdPartyIncentiveRate = new(RateAndAmountFormat5Choice)
	return r.ThirdPartyIncentiveRate
}

func (r *RateDetails11) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat2Choice {
	newValue := new(InterestRateUsedForPaymentFormat2Choice)
	r.InterestRateUsedForPayment = append(r.InterestRateUsedForPayment, newValue)
	return newValue
}

func (r *RateDetails11) AddNetDividendRate() *NetDividendRateFormat10Choice {
	newValue := new(NetDividendRateFormat10Choice)
	r.NetDividendRate = append(r.NetDividendRate, newValue)
	return newValue
}

func (r *RateDetails11) AddNonResidentRate() *RateAndAmountFormat5Choice {
	r.NonResidentRate = new(RateAndAmountFormat5Choice)
	return r.NonResidentRate
}

func (r *RateDetails11) SetProvisionalDividendRate(value, currency string) {
	r.ProvisionalDividendRate = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateDetails11) SetApplicableRate(value string) {
	r.ApplicableRate = (*PercentageRate)(&value)
}

func (r *RateDetails11) AddSolicitationFeeRate() *RateAndAmountFormat5Choice {
	r.SolicitationFeeRate = new(RateAndAmountFormat5Choice)
	return r.SolicitationFeeRate
}

func (r *RateDetails11) AddTaxCreditRate() *TaxCreditRateFormat2Choice {
	newValue := new(TaxCreditRateFormat2Choice)
	r.TaxCreditRate = append(r.TaxCreditRate, newValue)
	return newValue
}

func (r *RateDetails11) AddTaxRelatedRate() *RateTypeAndAmountAndStatus6 {
	newValue := new(RateTypeAndAmountAndStatus6)
	r.TaxRelatedRate = append(r.TaxRelatedRate, newValue)
	return newValue
}

func (r *RateDetails11) SetWithholdingTaxRate(value string) {
	r.WithholdingTaxRate = (*PercentageRate)(&value)
}

func (r *RateDetails11) SetTaxOnIncome(value string) {
	r.TaxOnIncome = (*PercentageRate)(&value)
}

func (r *RateDetails11) SetTaxOnProfits(value string) {
	r.TaxOnProfits = (*PercentageRate)(&value)
}

func (r *RateDetails11) SetTaxReclaimRate(value string) {
	r.TaxReclaimRate = (*PercentageRate)(&value)
}

func (r *RateDetails11) AddWithholdingOfForeignTax() *RateAndAmountFormat5Choice {
	r.WithholdingOfForeignTax = new(RateAndAmountFormat5Choice)
	return r.WithholdingOfForeignTax
}

func (r *RateDetails11) AddWithholdingOfLocalTax() *RateAndAmountFormat5Choice {
	r.WithholdingOfLocalTax = new(RateAndAmountFormat5Choice)
	return r.WithholdingOfLocalTax
}

func (r *RateDetails11) SetEqualisationRate(value, currency string) {
	r.EqualisationRate = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

// Provides information about the rates related to securities movement.
type RateDetails14 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat14Choice `xml:"AddtlTax,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat7Choice `xml:"GrssDvddRate,omitempty"`

	// The actual interest rate used for the payment of the interest for the specified interest period.
	// Usage guideline: It is used to provide the applicable rate for the current payment, after all calculations have been performed, that is, application of period and method of interest computation.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat5Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Percentage of the gross dividend rate on which tax must be paid .
	TaxRelatedRate []*RateTypeAndAmountAndStatus6 `xml:"TaxRltdRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxRate []*RateFormat10Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat14Choice `xml:"ChrgsFees,omitempty"`

	// Cash rate made available, as an incentive, in addition to the solicitation fee, in order to encourage early participation in an offer.
	EarlySolicitationFeeRate *SolicitationFeeRateFormat5Choice `xml:"EarlySlctnFeeRate,omitempty"`

	// Dividend is final.
	FinalDividendRate *RateAndAmountFormat15Choice `xml:"FnlDvddRate,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *RateFormat3Choice `xml:"FsclStmp,omitempty"`

	// Rate resulting from a fully franked dividend paid by a company; rate includes tax credit for companies that have made sufficient tax payments during fiscal period.
	FullyFrankedRate *RateAndAmountFormat14Choice `xml:"FullyFrnkdRate,omitempty"`

	// Cash rate made available in an event in order to encourage participation in the offer. As information, payment is made to a third party who has solicited an entity to take part in the offer.
	ThirdPartyIncentiveRate *RateFormat8Choice `xml:"ThrdPtyIncntivRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat9Choice `xml:"NetDvddRate,omitempty"`

	// Rate per share to which a non-resident is entitled.
	NonResidentRate *RateAndAmountFormat14Choice `xml:"NonResdtRate,omitempty"`

	// Dividend is provisional.
	ProvisionalDividendRate *RateAndAmountFormat15Choice `xml:"PrvsnlDvddRate,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *RateFormat3Choice `xml:"AplblRate,omitempty"`

	// Rate of the cash premium made available if the securities holder consents or participates to an event, for example consent fees or solicitation fee.
	SolicitationFeeRate *SolicitationFeeRateFormat5Choice `xml:"SlctnFeeRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat5Choice `xml:"TaxCdtRate,omitempty"`

	// Overall tax withheld at source by fund managers prior to considering the tax obligation of each unit holder.
	TaxOnIncome *RateAndAmountFormat14Choice `xml:"TaxOnIncm,omitempty"`

	// Taxation applied on an amount clearly identified as capital profits, capital gains.
	TaxOnProfits *RateFormat3Choice `xml:"TaxOnPrfts,omitempty"`

	// Percentage of cash that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimRate *RateFormat3Choice `xml:"TaxRclmRate,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction in which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTax []*RateAndAmountFormat20Choice `xml:"WhldgOfFrgnTax,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction in which the account owner is located, for which relief at source and/or reclaim may be possible.
	WithholdingOfLocalTax []*RateAndAmountFormat20Choice `xml:"WhldgOfLclTax,omitempty"`

	// Portion of the fund distribution which represents the average accrued income included in the purchase price for units bought during the account period.
	EqualisationRate *RateAndAmountFormat15Choice `xml:"EqulstnRate,omitempty"`
}

func (r *RateDetails14) AddAdditionalTax() *RateAndAmountFormat14Choice {
	r.AdditionalTax = new(RateAndAmountFormat14Choice)
	return r.AdditionalTax
}

func (r *RateDetails14) AddGrossDividendRate() *GrossDividendRateFormat7Choice {
	newValue := new(GrossDividendRateFormat7Choice)
	r.GrossDividendRate = append(r.GrossDividendRate, newValue)
	return newValue
}

func (r *RateDetails14) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat5Choice {
	newValue := new(InterestRateUsedForPaymentFormat5Choice)
	r.InterestRateUsedForPayment = append(r.InterestRateUsedForPayment, newValue)
	return newValue
}

func (r *RateDetails14) AddTaxRelatedRate() *RateTypeAndAmountAndStatus6 {
	newValue := new(RateTypeAndAmountAndStatus6)
	r.TaxRelatedRate = append(r.TaxRelatedRate, newValue)
	return newValue
}

func (r *RateDetails14) AddWithholdingTaxRate() *RateFormat10Choice {
	newValue := new(RateFormat10Choice)
	r.WithholdingTaxRate = append(r.WithholdingTaxRate, newValue)
	return newValue
}

func (r *RateDetails14) AddChargesFees() *RateAndAmountFormat14Choice {
	r.ChargesFees = new(RateAndAmountFormat14Choice)
	return r.ChargesFees
}

func (r *RateDetails14) AddEarlySolicitationFeeRate() *SolicitationFeeRateFormat5Choice {
	r.EarlySolicitationFeeRate = new(SolicitationFeeRateFormat5Choice)
	return r.EarlySolicitationFeeRate
}

func (r *RateDetails14) AddFinalDividendRate() *RateAndAmountFormat15Choice {
	r.FinalDividendRate = new(RateAndAmountFormat15Choice)
	return r.FinalDividendRate
}

func (r *RateDetails14) AddFiscalStamp() *RateFormat3Choice {
	r.FiscalStamp = new(RateFormat3Choice)
	return r.FiscalStamp
}

func (r *RateDetails14) AddFullyFrankedRate() *RateAndAmountFormat14Choice {
	r.FullyFrankedRate = new(RateAndAmountFormat14Choice)
	return r.FullyFrankedRate
}

func (r *RateDetails14) AddThirdPartyIncentiveRate() *RateFormat8Choice {
	r.ThirdPartyIncentiveRate = new(RateFormat8Choice)
	return r.ThirdPartyIncentiveRate
}

func (r *RateDetails14) AddNetDividendRate() *NetDividendRateFormat9Choice {
	newValue := new(NetDividendRateFormat9Choice)
	r.NetDividendRate = append(r.NetDividendRate, newValue)
	return newValue
}

func (r *RateDetails14) AddNonResidentRate() *RateAndAmountFormat14Choice {
	r.NonResidentRate = new(RateAndAmountFormat14Choice)
	return r.NonResidentRate
}

func (r *RateDetails14) AddProvisionalDividendRate() *RateAndAmountFormat15Choice {
	r.ProvisionalDividendRate = new(RateAndAmountFormat15Choice)
	return r.ProvisionalDividendRate
}

func (r *RateDetails14) AddApplicableRate() *RateFormat3Choice {
	r.ApplicableRate = new(RateFormat3Choice)
	return r.ApplicableRate
}

func (r *RateDetails14) AddSolicitationFeeRate() *SolicitationFeeRateFormat5Choice {
	r.SolicitationFeeRate = new(SolicitationFeeRateFormat5Choice)
	return r.SolicitationFeeRate
}

func (r *RateDetails14) AddTaxCreditRate() *TaxCreditRateFormat5Choice {
	newValue := new(TaxCreditRateFormat5Choice)
	r.TaxCreditRate = append(r.TaxCreditRate, newValue)
	return newValue
}

func (r *RateDetails14) AddTaxOnIncome() *RateAndAmountFormat14Choice {
	r.TaxOnIncome = new(RateAndAmountFormat14Choice)
	return r.TaxOnIncome
}

func (r *RateDetails14) AddTaxOnProfits() *RateFormat3Choice {
	r.TaxOnProfits = new(RateFormat3Choice)
	return r.TaxOnProfits
}

func (r *RateDetails14) AddTaxReclaimRate() *RateFormat3Choice {
	r.TaxReclaimRate = new(RateFormat3Choice)
	return r.TaxReclaimRate
}

func (r *RateDetails14) AddWithholdingOfForeignTax() *RateAndAmountFormat20Choice {
	newValue := new(RateAndAmountFormat20Choice)
	r.WithholdingOfForeignTax = append(r.WithholdingOfForeignTax, newValue)
	return newValue
}

func (r *RateDetails14) AddWithholdingOfLocalTax() *RateAndAmountFormat20Choice {
	newValue := new(RateAndAmountFormat20Choice)
	r.WithholdingOfLocalTax = append(r.WithholdingOfLocalTax, newValue)
	return newValue
}

func (r *RateDetails14) AddEqualisationRate() *RateAndAmountFormat15Choice {
	r.EqualisationRate = new(RateAndAmountFormat15Choice)
	return r.EqualisationRate
}

// Provides information about the rates related to securities movement.
type RateDetails15 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat5Choice `xml:"AddtlTax,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat5Choice `xml:"ChrgsFees,omitempty"`

	// Dividend is final.
	FinalDividendRate *ActiveCurrencyAnd13DecimalAmount `xml:"FnlDvddRate,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *PercentageRate `xml:"FsclStmp,omitempty"`

	// Rate resulting from a fully franked dividend paid by a company; rate includes tax credit for companies that have made sufficient tax payments during fiscal period.
	FullyFrankedRate *RateAndAmountFormat5Choice `xml:"FullyFrnkdRate,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat8Choice `xml:"GrssDvddRate,omitempty"`

	// Cash rate made available, as an incentive, in addition to the solicitation fee, in order to encourage early participation in an offer.
	EarlySolicitationFeeRate *RateAndAmountFormat5Choice `xml:"EarlySlctnFeeRate,omitempty"`

	// Cash rate made available in an event in order to encourage participation in the offer. As information, Payment is made to a third party who has solicited an entity to take part in the offer.
	ThirdPartyIncentiveRate *RateAndAmountFormat5Choice `xml:"ThrdPtyIncntivRate,omitempty"`

	// Actual interest rate used for the payment of the interest for the specified interest period.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat2Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat10Choice `xml:"NetDvddRate,omitempty"`

	// Rate per share to which a non-resident is entitled.
	NonResidentRate *RateAndAmountFormat5Choice `xml:"NonResdtRate,omitempty"`

	// Dividend is provisional.
	ProvisionalDividendRate *ActiveCurrencyAnd13DecimalAmount `xml:"PrvsnlDvddRate,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *PercentageRate `xml:"AplblRate,omitempty"`

	// Rate of the cash premium made available if the securities holder consents or participates to an event, for example consent fees or solicitation fee.
	SolicitationFeeRate *RateAndAmountFormat5Choice `xml:"SlctnFeeRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat2Choice `xml:"TaxCdtRate,omitempty"`

	// Percentage of the gross dividend rate on which tax must be paid .
	TaxRelatedRate []*RateTypeAndAmountAndStatus6 `xml:"TaxRltdRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxRate []*RateFormat11Choice `xml:"WhldgTaxRate,omitempty"`

	// Taxation applied on an amount clearly identified as an income.
	TaxOnIncome *RateAndAmountFormat5Choice `xml:"TaxOnIncm,omitempty"`

	// Taxation applied on an amount clearly identified as capital profits, capital gains.
	TaxOnProfits *PercentageRate `xml:"TaxOnPrfts,omitempty"`

	// Percentage of cash that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimRate *PercentageRate `xml:"TaxRclmRate,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction in which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTax []*RateAndAmountFormat21Choice `xml:"WhldgOfFrgnTax,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction in which the account owner is located, for which relief at source and/or reclaim may be possible.
	WithholdingOfLocalTax []*RateAndAmountFormat21Choice `xml:"WhldgOfLclTax,omitempty"`

	// Portion of the fund distribution which represents the average accrued income included in the purchase price for units bought during the account period.
	EqualisationRate *ActiveCurrencyAnd13DecimalAmount `xml:"EqulstnRate,omitempty"`
}

func (r *RateDetails15) AddAdditionalTax() *RateAndAmountFormat5Choice {
	r.AdditionalTax = new(RateAndAmountFormat5Choice)
	return r.AdditionalTax
}

func (r *RateDetails15) AddChargesFees() *RateAndAmountFormat5Choice {
	r.ChargesFees = new(RateAndAmountFormat5Choice)
	return r.ChargesFees
}

func (r *RateDetails15) SetFinalDividendRate(value, currency string) {
	r.FinalDividendRate = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateDetails15) SetFiscalStamp(value string) {
	r.FiscalStamp = (*PercentageRate)(&value)
}

func (r *RateDetails15) AddFullyFrankedRate() *RateAndAmountFormat5Choice {
	r.FullyFrankedRate = new(RateAndAmountFormat5Choice)
	return r.FullyFrankedRate
}

func (r *RateDetails15) AddGrossDividendRate() *GrossDividendRateFormat8Choice {
	newValue := new(GrossDividendRateFormat8Choice)
	r.GrossDividendRate = append(r.GrossDividendRate, newValue)
	return newValue
}

func (r *RateDetails15) AddEarlySolicitationFeeRate() *RateAndAmountFormat5Choice {
	r.EarlySolicitationFeeRate = new(RateAndAmountFormat5Choice)
	return r.EarlySolicitationFeeRate
}

func (r *RateDetails15) AddThirdPartyIncentiveRate() *RateAndAmountFormat5Choice {
	r.ThirdPartyIncentiveRate = new(RateAndAmountFormat5Choice)
	return r.ThirdPartyIncentiveRate
}

func (r *RateDetails15) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat2Choice {
	newValue := new(InterestRateUsedForPaymentFormat2Choice)
	r.InterestRateUsedForPayment = append(r.InterestRateUsedForPayment, newValue)
	return newValue
}

func (r *RateDetails15) AddNetDividendRate() *NetDividendRateFormat10Choice {
	newValue := new(NetDividendRateFormat10Choice)
	r.NetDividendRate = append(r.NetDividendRate, newValue)
	return newValue
}

func (r *RateDetails15) AddNonResidentRate() *RateAndAmountFormat5Choice {
	r.NonResidentRate = new(RateAndAmountFormat5Choice)
	return r.NonResidentRate
}

func (r *RateDetails15) SetProvisionalDividendRate(value, currency string) {
	r.ProvisionalDividendRate = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateDetails15) SetApplicableRate(value string) {
	r.ApplicableRate = (*PercentageRate)(&value)
}

func (r *RateDetails15) AddSolicitationFeeRate() *RateAndAmountFormat5Choice {
	r.SolicitationFeeRate = new(RateAndAmountFormat5Choice)
	return r.SolicitationFeeRate
}

func (r *RateDetails15) AddTaxCreditRate() *TaxCreditRateFormat2Choice {
	newValue := new(TaxCreditRateFormat2Choice)
	r.TaxCreditRate = append(r.TaxCreditRate, newValue)
	return newValue
}

func (r *RateDetails15) AddTaxRelatedRate() *RateTypeAndAmountAndStatus6 {
	newValue := new(RateTypeAndAmountAndStatus6)
	r.TaxRelatedRate = append(r.TaxRelatedRate, newValue)
	return newValue
}

func (r *RateDetails15) AddWithholdingTaxRate() *RateFormat11Choice {
	newValue := new(RateFormat11Choice)
	r.WithholdingTaxRate = append(r.WithholdingTaxRate, newValue)
	return newValue
}

func (r *RateDetails15) AddTaxOnIncome() *RateAndAmountFormat5Choice {
	r.TaxOnIncome = new(RateAndAmountFormat5Choice)
	return r.TaxOnIncome
}

func (r *RateDetails15) SetTaxOnProfits(value string) {
	r.TaxOnProfits = (*PercentageRate)(&value)
}

func (r *RateDetails15) SetTaxReclaimRate(value string) {
	r.TaxReclaimRate = (*PercentageRate)(&value)
}

func (r *RateDetails15) AddWithholdingOfForeignTax() *RateAndAmountFormat21Choice {
	newValue := new(RateAndAmountFormat21Choice)
	r.WithholdingOfForeignTax = append(r.WithholdingOfForeignTax, newValue)
	return newValue
}

func (r *RateDetails15) AddWithholdingOfLocalTax() *RateAndAmountFormat21Choice {
	newValue := new(RateAndAmountFormat21Choice)
	r.WithholdingOfLocalTax = append(r.WithholdingOfLocalTax, newValue)
	return newValue
}

func (r *RateDetails15) SetEqualisationRate(value, currency string) {
	r.EqualisationRate = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

// Provides information about the rates related to securities movement.
type RateDetails2 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat5Choice `xml:"AddtlTax,omitempty"`

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

	// Actual interest rate used for the payment of the interest for the specified interest period.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat2Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat2Choice `xml:"NetDvddRate,omitempty"`

	// Rate per share to which a non-resident is entitled.
	NonResidentRate *RateAndAmountFormat5Choice `xml:"NonResdtRate,omitempty"`

	// Dividend is provisional.
	ProvisionalDividendRate *ActiveCurrencyAnd13DecimalAmount `xml:"PrvsnlDvddRate,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *PercentageRate `xml:"AplblRate,omitempty"`

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
}

func (r *RateDetails2) AddAdditionalTax() *RateAndAmountFormat5Choice {
	r.AdditionalTax = new(RateAndAmountFormat5Choice)
	return r.AdditionalTax
}

func (r *RateDetails2) AddChargesFees() *RateAndAmountFormat5Choice {
	r.ChargesFees = new(RateAndAmountFormat5Choice)
	return r.ChargesFees
}

func (r *RateDetails2) SetFinalDividendRate(value, currency string) {
	r.FinalDividendRate = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateDetails2) SetFiscalStamp(value string) {
	r.FiscalStamp = (*PercentageRate)(&value)
}

func (r *RateDetails2) AddFullyFrankedRate() *RateAndAmountFormat5Choice {
	r.FullyFrankedRate = new(RateAndAmountFormat5Choice)
	return r.FullyFrankedRate
}

func (r *RateDetails2) AddGrossDividendRate() *GrossDividendRateFormat2Choice {
	newValue := new(GrossDividendRateFormat2Choice)
	r.GrossDividendRate = append(r.GrossDividendRate, newValue)
	return newValue
}

func (r *RateDetails2) SetCashIncentiveRate(value string) {
	r.CashIncentiveRate = (*PercentageRate)(&value)
}

func (r *RateDetails2) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat2Choice {
	newValue := new(InterestRateUsedForPaymentFormat2Choice)
	r.InterestRateUsedForPayment = append(r.InterestRateUsedForPayment, newValue)
	return newValue
}

func (r *RateDetails2) AddNetDividendRate() *NetDividendRateFormat2Choice {
	newValue := new(NetDividendRateFormat2Choice)
	r.NetDividendRate = append(r.NetDividendRate, newValue)
	return newValue
}

func (r *RateDetails2) AddNonResidentRate() *RateAndAmountFormat5Choice {
	r.NonResidentRate = new(RateAndAmountFormat5Choice)
	return r.NonResidentRate
}

func (r *RateDetails2) SetProvisionalDividendRate(value, currency string) {
	r.ProvisionalDividendRate = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateDetails2) SetApplicableRate(value string) {
	r.ApplicableRate = (*PercentageRate)(&value)
}

func (r *RateDetails2) SetSolicitationFeeRate(value string) {
	r.SolicitationFeeRate = (*PercentageRate)(&value)
}

func (r *RateDetails2) AddTaxCreditRate() *TaxCreditRateFormat2Choice {
	newValue := new(TaxCreditRateFormat2Choice)
	r.TaxCreditRate = append(r.TaxCreditRate, newValue)
	return newValue
}

func (r *RateDetails2) AddTaxRelatedRate() *RateTypeAndAmountAndStatus6 {
	newValue := new(RateTypeAndAmountAndStatus6)
	r.TaxRelatedRate = append(r.TaxRelatedRate, newValue)
	return newValue
}

func (r *RateDetails2) SetWithholdingTaxRate(value string) {
	r.WithholdingTaxRate = (*PercentageRate)(&value)
}

func (r *RateDetails2) SetTaxOnIncome(value string) {
	r.TaxOnIncome = (*PercentageRate)(&value)
}

func (r *RateDetails2) SetTaxOnProfits(value string) {
	r.TaxOnProfits = (*PercentageRate)(&value)
}

func (r *RateDetails2) SetTaxReclaimRate(value string) {
	r.TaxReclaimRate = (*PercentageRate)(&value)
}

func (r *RateDetails2) AddWithholdingOfForeignTax() *RateAndAmountFormat5Choice {
	r.WithholdingOfForeignTax = new(RateAndAmountFormat5Choice)
	return r.WithholdingOfForeignTax
}

func (r *RateDetails2) AddWithholdingOfLocalTax() *RateAndAmountFormat5Choice {
	r.WithholdingOfLocalTax = new(RateAndAmountFormat5Choice)
	return r.WithholdingOfLocalTax
}

// Provides information about the rates related to securities movement.
type RateDetails22 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat37Choice `xml:"AddtlTax,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat20Choice `xml:"GrssDvddRate,omitempty"`

	// The actual interest rate used for the payment of the interest for the specified interest period.
	// Usage guideline: It is used to provide the applicable rate for the current payment, after all calculations have been performed, that is, application of period and method of interest computation.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat8Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Percentage of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxRate []*RateAndAmountFormat41Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate at which the income will be withheld by a jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate (TAXR) levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTax []*RateAndAmountFormat41Choice `xml:"ScndLvlTax,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat37Choice `xml:"ChrgsFees,omitempty"`

	// Cash rate made available, as an incentive, in addition to the solicitation fee, in order to encourage early participation in an offer.
	EarlySolicitationFeeRate *SolicitationFeeRateFormat7Choice `xml:"EarlySlctnFeeRate,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *RateFormat3Choice `xml:"FsclStmp,omitempty"`

	// Rate resulting from a fully franked dividend paid by a company; rate includes tax credit for companies that have made sufficient tax payments during fiscal period.
	FullyFrankedRate *RateAndAmountFormat37Choice `xml:"FullyFrnkdRate,omitempty"`

	// Cash rate made available in an event in order to encourage participation in the offer. As information, payment is made to a third party who has solicited an entity to take part in the offer.
	ThirdPartyIncentiveRate *RateFormat20Choice `xml:"ThrdPtyIncntivRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat22Choice `xml:"NetDvddRate,omitempty"`

	// Rate per share to which a non-resident is entitled.
	NonResidentRate *RateAndAmountFormat37Choice `xml:"NonResdtRate,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *RateFormat3Choice `xml:"AplblRate,omitempty"`

	// Rate of the cash premium made available if the securities holder consents or participates to an event, for example consent fees or solicitation fee.
	SolicitationFeeRate *SolicitationFeeRateFormat7Choice `xml:"SlctnFeeRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat8Choice `xml:"TaxCdtRate,omitempty"`

	// Overall tax withheld at source by fund managers prior to considering the tax obligation of each unit holder.
	TaxOnIncome *RateAndAmountFormat37Choice `xml:"TaxOnIncm,omitempty"`

	// Taxation applied on an amount clearly identified as capital profits, capital gains.
	TaxOnProfits *RateFormat3Choice `xml:"TaxOnPrfts,omitempty"`

	// Percentage of cash that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimRate *RateFormat3Choice `xml:"TaxRclmRate,omitempty"`

	// Portion of the fund distribution which represents the average accrued income included in the purchase price for units bought during the account period.
	EqualisationRate *RateAndAmountFormat42Choice `xml:"EqulstnRate,omitempty"`
}

func (r *RateDetails22) AddAdditionalTax() *RateAndAmountFormat37Choice {
	r.AdditionalTax = new(RateAndAmountFormat37Choice)
	return r.AdditionalTax
}

func (r *RateDetails22) AddGrossDividendRate() *GrossDividendRateFormat20Choice {
	newValue := new(GrossDividendRateFormat20Choice)
	r.GrossDividendRate = append(r.GrossDividendRate, newValue)
	return newValue
}

func (r *RateDetails22) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat8Choice {
	newValue := new(InterestRateUsedForPaymentFormat8Choice)
	r.InterestRateUsedForPayment = append(r.InterestRateUsedForPayment, newValue)
	return newValue
}

func (r *RateDetails22) AddWithholdingTaxRate() *RateAndAmountFormat41Choice {
	newValue := new(RateAndAmountFormat41Choice)
	r.WithholdingTaxRate = append(r.WithholdingTaxRate, newValue)
	return newValue
}

func (r *RateDetails22) AddSecondLevelTax() *RateAndAmountFormat41Choice {
	newValue := new(RateAndAmountFormat41Choice)
	r.SecondLevelTax = append(r.SecondLevelTax, newValue)
	return newValue
}

func (r *RateDetails22) AddChargesFees() *RateAndAmountFormat37Choice {
	r.ChargesFees = new(RateAndAmountFormat37Choice)
	return r.ChargesFees
}

func (r *RateDetails22) AddEarlySolicitationFeeRate() *SolicitationFeeRateFormat7Choice {
	r.EarlySolicitationFeeRate = new(SolicitationFeeRateFormat7Choice)
	return r.EarlySolicitationFeeRate
}

func (r *RateDetails22) AddFiscalStamp() *RateFormat3Choice {
	r.FiscalStamp = new(RateFormat3Choice)
	return r.FiscalStamp
}

func (r *RateDetails22) AddFullyFrankedRate() *RateAndAmountFormat37Choice {
	r.FullyFrankedRate = new(RateAndAmountFormat37Choice)
	return r.FullyFrankedRate
}

func (r *RateDetails22) AddThirdPartyIncentiveRate() *RateFormat20Choice {
	r.ThirdPartyIncentiveRate = new(RateFormat20Choice)
	return r.ThirdPartyIncentiveRate
}

func (r *RateDetails22) AddNetDividendRate() *NetDividendRateFormat22Choice {
	newValue := new(NetDividendRateFormat22Choice)
	r.NetDividendRate = append(r.NetDividendRate, newValue)
	return newValue
}

func (r *RateDetails22) AddNonResidentRate() *RateAndAmountFormat37Choice {
	r.NonResidentRate = new(RateAndAmountFormat37Choice)
	return r.NonResidentRate
}

func (r *RateDetails22) AddApplicableRate() *RateFormat3Choice {
	r.ApplicableRate = new(RateFormat3Choice)
	return r.ApplicableRate
}

func (r *RateDetails22) AddSolicitationFeeRate() *SolicitationFeeRateFormat7Choice {
	r.SolicitationFeeRate = new(SolicitationFeeRateFormat7Choice)
	return r.SolicitationFeeRate
}

func (r *RateDetails22) AddTaxCreditRate() *TaxCreditRateFormat8Choice {
	newValue := new(TaxCreditRateFormat8Choice)
	r.TaxCreditRate = append(r.TaxCreditRate, newValue)
	return newValue
}

func (r *RateDetails22) AddTaxOnIncome() *RateAndAmountFormat37Choice {
	r.TaxOnIncome = new(RateAndAmountFormat37Choice)
	return r.TaxOnIncome
}

func (r *RateDetails22) AddTaxOnProfits() *RateFormat3Choice {
	r.TaxOnProfits = new(RateFormat3Choice)
	return r.TaxOnProfits
}

func (r *RateDetails22) AddTaxReclaimRate() *RateFormat3Choice {
	r.TaxReclaimRate = new(RateFormat3Choice)
	return r.TaxReclaimRate
}

func (r *RateDetails22) AddEqualisationRate() *RateAndAmountFormat42Choice {
	r.EqualisationRate = new(RateAndAmountFormat42Choice)
	return r.EqualisationRate
}

// Provides information about the rates related to securities movement.
type RateDetails23 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat39Choice `xml:"AddtlTax,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat39Choice `xml:"ChrgsFees,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *PercentageRate `xml:"FsclStmp,omitempty"`

	// Rate resulting from a fully franked dividend paid by a company; rate includes tax credit for companies that have made sufficient tax payments during fiscal period.
	FullyFrankedRate *RateAndAmountFormat39Choice `xml:"FullyFrnkdRate,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat22Choice `xml:"GrssDvddRate,omitempty"`

	// Cash rate made available, as an incentive, in addition to the solicitation fee, in order to encourage early participation in an offer.
	EarlySolicitationFeeRate *SolicitationFeeRateFormat8Choice `xml:"EarlySlctnFeeRate,omitempty"`

	// Cash rate made available in an event in order to encourage participation in the offer. As information, Payment is made to a third party who has solicited an entity to take part in the offer.
	ThirdPartyIncentiveRate *RateAndAmountFormat39Choice `xml:"ThrdPtyIncntivRate,omitempty"`

	// Actual interest rate used for the payment of the interest for the specified interest period.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat7Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat24Choice `xml:"NetDvddRate,omitempty"`

	// Rate per share to which a non-resident is entitled.
	NonResidentRate *RateAndAmountFormat39Choice `xml:"NonResdtRate,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *PercentageRate `xml:"AplblRate,omitempty"`

	// Rate of the cash premium made available if the securities holder consents or participates to an event, for example consent fees or solicitation fee.
	SolicitationFeeRate *SolicitationFeeRateFormat8Choice `xml:"SlctnFeeRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat7Choice `xml:"TaxCdtRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxRate []*RateAndAmountFormat40Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate at which the income will be withheld by a jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate (TAXR) levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTax []*RateAndAmountFormat40Choice `xml:"ScndLvlTax,omitempty"`

	// Taxation applied on an amount clearly identified as an income.
	TaxOnIncome *RateAndAmountFormat39Choice `xml:"TaxOnIncm,omitempty"`

	// Taxation applied on an amount clearly identified as capital profits, capital gains.
	TaxOnProfits *PercentageRate `xml:"TaxOnPrfts,omitempty"`

	// Percentage of cash that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimRate *PercentageRate `xml:"TaxRclmRate,omitempty"`

	// Portion of the fund distribution which represents the average accrued income included in the purchase price for units bought during the account period.
	EqualisationRate *ActiveCurrencyAnd13DecimalAmount `xml:"EqulstnRate,omitempty"`
}

func (r *RateDetails23) AddAdditionalTax() *RateAndAmountFormat39Choice {
	r.AdditionalTax = new(RateAndAmountFormat39Choice)
	return r.AdditionalTax
}

func (r *RateDetails23) AddChargesFees() *RateAndAmountFormat39Choice {
	r.ChargesFees = new(RateAndAmountFormat39Choice)
	return r.ChargesFees
}

func (r *RateDetails23) SetFiscalStamp(value string) {
	r.FiscalStamp = (*PercentageRate)(&value)
}

func (r *RateDetails23) AddFullyFrankedRate() *RateAndAmountFormat39Choice {
	r.FullyFrankedRate = new(RateAndAmountFormat39Choice)
	return r.FullyFrankedRate
}

func (r *RateDetails23) AddGrossDividendRate() *GrossDividendRateFormat22Choice {
	newValue := new(GrossDividendRateFormat22Choice)
	r.GrossDividendRate = append(r.GrossDividendRate, newValue)
	return newValue
}

func (r *RateDetails23) AddEarlySolicitationFeeRate() *SolicitationFeeRateFormat8Choice {
	r.EarlySolicitationFeeRate = new(SolicitationFeeRateFormat8Choice)
	return r.EarlySolicitationFeeRate
}

func (r *RateDetails23) AddThirdPartyIncentiveRate() *RateAndAmountFormat39Choice {
	r.ThirdPartyIncentiveRate = new(RateAndAmountFormat39Choice)
	return r.ThirdPartyIncentiveRate
}

func (r *RateDetails23) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat7Choice {
	newValue := new(InterestRateUsedForPaymentFormat7Choice)
	r.InterestRateUsedForPayment = append(r.InterestRateUsedForPayment, newValue)
	return newValue
}

func (r *RateDetails23) AddNetDividendRate() *NetDividendRateFormat24Choice {
	newValue := new(NetDividendRateFormat24Choice)
	r.NetDividendRate = append(r.NetDividendRate, newValue)
	return newValue
}

func (r *RateDetails23) AddNonResidentRate() *RateAndAmountFormat39Choice {
	r.NonResidentRate = new(RateAndAmountFormat39Choice)
	return r.NonResidentRate
}

func (r *RateDetails23) SetApplicableRate(value string) {
	r.ApplicableRate = (*PercentageRate)(&value)
}

func (r *RateDetails23) AddSolicitationFeeRate() *SolicitationFeeRateFormat8Choice {
	r.SolicitationFeeRate = new(SolicitationFeeRateFormat8Choice)
	return r.SolicitationFeeRate
}

func (r *RateDetails23) AddTaxCreditRate() *TaxCreditRateFormat7Choice {
	newValue := new(TaxCreditRateFormat7Choice)
	r.TaxCreditRate = append(r.TaxCreditRate, newValue)
	return newValue
}

func (r *RateDetails23) AddWithholdingTaxRate() *RateAndAmountFormat40Choice {
	newValue := new(RateAndAmountFormat40Choice)
	r.WithholdingTaxRate = append(r.WithholdingTaxRate, newValue)
	return newValue
}

func (r *RateDetails23) AddSecondLevelTax() *RateAndAmountFormat40Choice {
	newValue := new(RateAndAmountFormat40Choice)
	r.SecondLevelTax = append(r.SecondLevelTax, newValue)
	return newValue
}

func (r *RateDetails23) AddTaxOnIncome() *RateAndAmountFormat39Choice {
	r.TaxOnIncome = new(RateAndAmountFormat39Choice)
	return r.TaxOnIncome
}

func (r *RateDetails23) SetTaxOnProfits(value string) {
	r.TaxOnProfits = (*PercentageRate)(&value)
}

func (r *RateDetails23) SetTaxReclaimRate(value string) {
	r.TaxReclaimRate = (*PercentageRate)(&value)
}

func (r *RateDetails23) SetEqualisationRate(value, currency string) {
	r.EqualisationRate = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

// Provides information about the rates related to securities movement.
type RateDetails24 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat43Choice `xml:"AddtlTax,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat43Choice `xml:"ChrgsFees,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *PercentageRate `xml:"FsclStmp,omitempty"`

	// Rate resulting from a fully franked dividend paid by a company; rate includes tax credit for companies that have made sufficient tax payments during fiscal period.
	FullyFrankedRate *RateAndAmountFormat43Choice `xml:"FullyFrnkdRate,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat24Choice `xml:"GrssDvddRate,omitempty"`

	// Cash rate made available, as an incentive, in addition to the solicitation fee, in order to encourage early participation in an offer.
	EarlySolicitationFeeRate *SolicitationFeeRateFormat9Choice `xml:"EarlySlctnFeeRate,omitempty"`

	// Cash rate made available in an event in order to encourage participation in the offer. As information, Payment is made to a third party who has solicited an entity to take part in the offer.
	ThirdPartyIncentiveRate *RateAndAmountFormat43Choice `xml:"ThrdPtyIncntivRate,omitempty"`

	// Actual interest rate used for the payment of the interest for the specified interest period.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat9Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat26Choice `xml:"NetDvddRate,omitempty"`

	// Rate per share to which a non-resident is entitled.
	NonResidentRate *RateAndAmountFormat43Choice `xml:"NonResdtRate,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *PercentageRate `xml:"AplblRate,omitempty"`

	// Rate of the cash premium made available if the securities holder consents or participates to an event, for example consent fees or solicitation fee.
	SolicitationFeeRate *SolicitationFeeRateFormat9Choice `xml:"SlctnFeeRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat9Choice `xml:"TaxCdtRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxRate []*RateAndAmountFormat45Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate at which the income will be withheld by a jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate (TAXR) levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTax []*RateAndAmountFormat45Choice `xml:"ScndLvlTax,omitempty"`

	// Taxation applied on an amount clearly identified as an income.
	TaxOnIncome *RateAndAmountFormat43Choice `xml:"TaxOnIncm,omitempty"`

	// Taxation applied on an amount clearly identified as capital profits, capital gains.
	TaxOnProfits *PercentageRate `xml:"TaxOnPrfts,omitempty"`

	// Percentage of cash that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimRate *PercentageRate `xml:"TaxRclmRate,omitempty"`

	// Portion of the fund distribution which represents the average accrued income included in the purchase price for units bought during the account period.
	EqualisationRate *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"EqulstnRate,omitempty"`
}

func (r *RateDetails24) AddAdditionalTax() *RateAndAmountFormat43Choice {
	r.AdditionalTax = new(RateAndAmountFormat43Choice)
	return r.AdditionalTax
}

func (r *RateDetails24) AddChargesFees() *RateAndAmountFormat43Choice {
	r.ChargesFees = new(RateAndAmountFormat43Choice)
	return r.ChargesFees
}

func (r *RateDetails24) SetFiscalStamp(value string) {
	r.FiscalStamp = (*PercentageRate)(&value)
}

func (r *RateDetails24) AddFullyFrankedRate() *RateAndAmountFormat43Choice {
	r.FullyFrankedRate = new(RateAndAmountFormat43Choice)
	return r.FullyFrankedRate
}

func (r *RateDetails24) AddGrossDividendRate() *GrossDividendRateFormat24Choice {
	newValue := new(GrossDividendRateFormat24Choice)
	r.GrossDividendRate = append(r.GrossDividendRate, newValue)
	return newValue
}

func (r *RateDetails24) AddEarlySolicitationFeeRate() *SolicitationFeeRateFormat9Choice {
	r.EarlySolicitationFeeRate = new(SolicitationFeeRateFormat9Choice)
	return r.EarlySolicitationFeeRate
}

func (r *RateDetails24) AddThirdPartyIncentiveRate() *RateAndAmountFormat43Choice {
	r.ThirdPartyIncentiveRate = new(RateAndAmountFormat43Choice)
	return r.ThirdPartyIncentiveRate
}

func (r *RateDetails24) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat9Choice {
	newValue := new(InterestRateUsedForPaymentFormat9Choice)
	r.InterestRateUsedForPayment = append(r.InterestRateUsedForPayment, newValue)
	return newValue
}

func (r *RateDetails24) AddNetDividendRate() *NetDividendRateFormat26Choice {
	newValue := new(NetDividendRateFormat26Choice)
	r.NetDividendRate = append(r.NetDividendRate, newValue)
	return newValue
}

func (r *RateDetails24) AddNonResidentRate() *RateAndAmountFormat43Choice {
	r.NonResidentRate = new(RateAndAmountFormat43Choice)
	return r.NonResidentRate
}

func (r *RateDetails24) SetApplicableRate(value string) {
	r.ApplicableRate = (*PercentageRate)(&value)
}

func (r *RateDetails24) AddSolicitationFeeRate() *SolicitationFeeRateFormat9Choice {
	r.SolicitationFeeRate = new(SolicitationFeeRateFormat9Choice)
	return r.SolicitationFeeRate
}

func (r *RateDetails24) AddTaxCreditRate() *TaxCreditRateFormat9Choice {
	newValue := new(TaxCreditRateFormat9Choice)
	r.TaxCreditRate = append(r.TaxCreditRate, newValue)
	return newValue
}

func (r *RateDetails24) AddWithholdingTaxRate() *RateAndAmountFormat45Choice {
	newValue := new(RateAndAmountFormat45Choice)
	r.WithholdingTaxRate = append(r.WithholdingTaxRate, newValue)
	return newValue
}

func (r *RateDetails24) AddSecondLevelTax() *RateAndAmountFormat45Choice {
	newValue := new(RateAndAmountFormat45Choice)
	r.SecondLevelTax = append(r.SecondLevelTax, newValue)
	return newValue
}

func (r *RateDetails24) AddTaxOnIncome() *RateAndAmountFormat43Choice {
	r.TaxOnIncome = new(RateAndAmountFormat43Choice)
	return r.TaxOnIncome
}

func (r *RateDetails24) SetTaxOnProfits(value string) {
	r.TaxOnProfits = (*PercentageRate)(&value)
}

func (r *RateDetails24) SetTaxReclaimRate(value string) {
	r.TaxReclaimRate = (*PercentageRate)(&value)
}

func (r *RateDetails24) SetEqualisationRate(value, currency string) {
	r.EqualisationRate = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

// Provides information about the rates related to securities movement.
type RateDetails25 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat46Choice `xml:"AddtlTax,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat26Choice `xml:"GrssDvddRate,omitempty"`

	// The actual interest rate used for the payment of the interest for the specified interest period.
	// Usage guideline: It is used to provide the applicable rate for the current payment, after all calculations have been performed, that is, application of period and method of interest computation.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat10Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Percentage of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxRate []*RateAndAmountFormat47Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate at which the income will be withheld by a jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate (TAXR) levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTax []*RateAndAmountFormat47Choice `xml:"ScndLvlTax,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat46Choice `xml:"ChrgsFees,omitempty"`

	// Cash rate made available, as an incentive, in addition to the solicitation fee, in order to encourage early participation in an offer.
	EarlySolicitationFeeRate *SolicitationFeeRateFormat10Choice `xml:"EarlySlctnFeeRate,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *RateFormat3Choice `xml:"FsclStmp,omitempty"`

	// Rate resulting from a fully franked dividend paid by a company; rate includes tax credit for companies that have made sufficient tax payments during fiscal period.
	FullyFrankedRate *RateAndAmountFormat46Choice `xml:"FullyFrnkdRate,omitempty"`

	// Cash rate made available in an event in order to encourage participation in the offer. As information, payment is made to a third party who has solicited an entity to take part in the offer.
	ThirdPartyIncentiveRate *RateFormat21Choice `xml:"ThrdPtyIncntivRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat28Choice `xml:"NetDvddRate,omitempty"`

	// Rate per share to which a non-resident is entitled.
	NonResidentRate *RateAndAmountFormat46Choice `xml:"NonResdtRate,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *RateFormat3Choice `xml:"AplblRate,omitempty"`

	// Rate of the cash premium made available if the securities holder consents or participates to an event, for example consent fees or solicitation fee.
	SolicitationFeeRate *SolicitationFeeRateFormat10Choice `xml:"SlctnFeeRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat10Choice `xml:"TaxCdtRate,omitempty"`

	// Overall tax withheld at source by fund managers prior to considering the tax obligation of each unit holder.
	TaxOnIncome *RateAndAmountFormat46Choice `xml:"TaxOnIncm,omitempty"`

	// Taxation applied on an amount clearly identified as capital profits, capital gains.
	TaxOnProfits *RateFormat3Choice `xml:"TaxOnPrfts,omitempty"`

	// Percentage of cash that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimRate *RateFormat3Choice `xml:"TaxRclmRate,omitempty"`

	// Portion of the fund distribution which represents the average accrued income included in the purchase price for units bought during the account period.
	EqualisationRate *RateAndAmountFormat48Choice `xml:"EqulstnRate,omitempty"`
}

func (r *RateDetails25) AddAdditionalTax() *RateAndAmountFormat46Choice {
	r.AdditionalTax = new(RateAndAmountFormat46Choice)
	return r.AdditionalTax
}

func (r *RateDetails25) AddGrossDividendRate() *GrossDividendRateFormat26Choice {
	newValue := new(GrossDividendRateFormat26Choice)
	r.GrossDividendRate = append(r.GrossDividendRate, newValue)
	return newValue
}

func (r *RateDetails25) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat10Choice {
	newValue := new(InterestRateUsedForPaymentFormat10Choice)
	r.InterestRateUsedForPayment = append(r.InterestRateUsedForPayment, newValue)
	return newValue
}

func (r *RateDetails25) AddWithholdingTaxRate() *RateAndAmountFormat47Choice {
	newValue := new(RateAndAmountFormat47Choice)
	r.WithholdingTaxRate = append(r.WithholdingTaxRate, newValue)
	return newValue
}

func (r *RateDetails25) AddSecondLevelTax() *RateAndAmountFormat47Choice {
	newValue := new(RateAndAmountFormat47Choice)
	r.SecondLevelTax = append(r.SecondLevelTax, newValue)
	return newValue
}

func (r *RateDetails25) AddChargesFees() *RateAndAmountFormat46Choice {
	r.ChargesFees = new(RateAndAmountFormat46Choice)
	return r.ChargesFees
}

func (r *RateDetails25) AddEarlySolicitationFeeRate() *SolicitationFeeRateFormat10Choice {
	r.EarlySolicitationFeeRate = new(SolicitationFeeRateFormat10Choice)
	return r.EarlySolicitationFeeRate
}

func (r *RateDetails25) AddFiscalStamp() *RateFormat3Choice {
	r.FiscalStamp = new(RateFormat3Choice)
	return r.FiscalStamp
}

func (r *RateDetails25) AddFullyFrankedRate() *RateAndAmountFormat46Choice {
	r.FullyFrankedRate = new(RateAndAmountFormat46Choice)
	return r.FullyFrankedRate
}

func (r *RateDetails25) AddThirdPartyIncentiveRate() *RateFormat21Choice {
	r.ThirdPartyIncentiveRate = new(RateFormat21Choice)
	return r.ThirdPartyIncentiveRate
}

func (r *RateDetails25) AddNetDividendRate() *NetDividendRateFormat28Choice {
	newValue := new(NetDividendRateFormat28Choice)
	r.NetDividendRate = append(r.NetDividendRate, newValue)
	return newValue
}

func (r *RateDetails25) AddNonResidentRate() *RateAndAmountFormat46Choice {
	r.NonResidentRate = new(RateAndAmountFormat46Choice)
	return r.NonResidentRate
}

func (r *RateDetails25) AddApplicableRate() *RateFormat3Choice {
	r.ApplicableRate = new(RateFormat3Choice)
	return r.ApplicableRate
}

func (r *RateDetails25) AddSolicitationFeeRate() *SolicitationFeeRateFormat10Choice {
	r.SolicitationFeeRate = new(SolicitationFeeRateFormat10Choice)
	return r.SolicitationFeeRate
}

func (r *RateDetails25) AddTaxCreditRate() *TaxCreditRateFormat10Choice {
	newValue := new(TaxCreditRateFormat10Choice)
	r.TaxCreditRate = append(r.TaxCreditRate, newValue)
	return newValue
}

func (r *RateDetails25) AddTaxOnIncome() *RateAndAmountFormat46Choice {
	r.TaxOnIncome = new(RateAndAmountFormat46Choice)
	return r.TaxOnIncome
}

func (r *RateDetails25) AddTaxOnProfits() *RateFormat3Choice {
	r.TaxOnProfits = new(RateFormat3Choice)
	return r.TaxOnProfits
}

func (r *RateDetails25) AddTaxReclaimRate() *RateFormat3Choice {
	r.TaxReclaimRate = new(RateFormat3Choice)
	return r.TaxReclaimRate
}

func (r *RateDetails25) AddEqualisationRate() *RateAndAmountFormat48Choice {
	r.EqualisationRate = new(RateAndAmountFormat48Choice)
	return r.EqualisationRate
}

// Provides information about the rates related to securities movement.
type RateDetails26 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat37Choice `xml:"AddtlTax,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat20Choice `xml:"GrssDvddRate,omitempty"`

	// The actual interest rate used for the payment of the interest for the specified interest period.
	// Usage guideline: It is used to provide the applicable rate for the current payment, after all calculations have been performed, that is, application of period and method of interest computation.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat8Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Percentage of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxRate []*RateAndAmountFormat41Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate at which the income will be withheld by a jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate (TAXR) levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTax []*RateAndAmountFormat41Choice `xml:"ScndLvlTax,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat37Choice `xml:"ChrgsFees,omitempty"`

	// Cash rate made available, as an incentive, in addition to the solicitation fee, in order to encourage early participation in an offer.
	EarlySolicitationFeeRate *SolicitationFeeRateFormat7Choice `xml:"EarlySlctnFeeRate,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *RateFormat3Choice `xml:"FsclStmp,omitempty"`

	// Cash rate made available in an event in order to encourage participation in the offer. As information, payment is made to a third party who has solicited an entity to take part in the offer.
	ThirdPartyIncentiveRate *RateFormat20Choice `xml:"ThrdPtyIncntivRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat22Choice `xml:"NetDvddRate,omitempty"`

	// Rate per share to which a non-resident is entitled.
	NonResidentRate *RateAndAmountFormat37Choice `xml:"NonResdtRate,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *RateFormat3Choice `xml:"AplblRate,omitempty"`

	// Rate of the cash premium made available if the securities holder consents or participates to an event, for example consent fees or solicitation fee.
	SolicitationFeeRate *SolicitationFeeRateFormat7Choice `xml:"SlctnFeeRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat8Choice `xml:"TaxCdtRate,omitempty"`

	// Overall tax withheld at source by fund managers prior to considering the tax obligation of each unit holder.
	TaxOnIncome *RateAndAmountFormat37Choice `xml:"TaxOnIncm,omitempty"`

	// Taxation applied on an amount clearly identified as capital profits, capital gains.
	TaxOnProfits *RateFormat3Choice `xml:"TaxOnPrfts,omitempty"`

	// Percentage of cash that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimRate *RateFormat3Choice `xml:"TaxRclmRate,omitempty"`

	// Portion of the fund distribution which represents the average accrued income included in the purchase price for units bought during the account period.
	EqualisationRate *RateAndAmountFormat42Choice `xml:"EqulstnRate,omitempty"`
}

func (r *RateDetails26) AddAdditionalTax() *RateAndAmountFormat37Choice {
	r.AdditionalTax = new(RateAndAmountFormat37Choice)
	return r.AdditionalTax
}

func (r *RateDetails26) AddGrossDividendRate() *GrossDividendRateFormat20Choice {
	newValue := new(GrossDividendRateFormat20Choice)
	r.GrossDividendRate = append(r.GrossDividendRate, newValue)
	return newValue
}

func (r *RateDetails26) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat8Choice {
	newValue := new(InterestRateUsedForPaymentFormat8Choice)
	r.InterestRateUsedForPayment = append(r.InterestRateUsedForPayment, newValue)
	return newValue
}

func (r *RateDetails26) AddWithholdingTaxRate() *RateAndAmountFormat41Choice {
	newValue := new(RateAndAmountFormat41Choice)
	r.WithholdingTaxRate = append(r.WithholdingTaxRate, newValue)
	return newValue
}

func (r *RateDetails26) AddSecondLevelTax() *RateAndAmountFormat41Choice {
	newValue := new(RateAndAmountFormat41Choice)
	r.SecondLevelTax = append(r.SecondLevelTax, newValue)
	return newValue
}

func (r *RateDetails26) AddChargesFees() *RateAndAmountFormat37Choice {
	r.ChargesFees = new(RateAndAmountFormat37Choice)
	return r.ChargesFees
}

func (r *RateDetails26) AddEarlySolicitationFeeRate() *SolicitationFeeRateFormat7Choice {
	r.EarlySolicitationFeeRate = new(SolicitationFeeRateFormat7Choice)
	return r.EarlySolicitationFeeRate
}

func (r *RateDetails26) AddFiscalStamp() *RateFormat3Choice {
	r.FiscalStamp = new(RateFormat3Choice)
	return r.FiscalStamp
}

func (r *RateDetails26) AddThirdPartyIncentiveRate() *RateFormat20Choice {
	r.ThirdPartyIncentiveRate = new(RateFormat20Choice)
	return r.ThirdPartyIncentiveRate
}

func (r *RateDetails26) AddNetDividendRate() *NetDividendRateFormat22Choice {
	newValue := new(NetDividendRateFormat22Choice)
	r.NetDividendRate = append(r.NetDividendRate, newValue)
	return newValue
}

func (r *RateDetails26) AddNonResidentRate() *RateAndAmountFormat37Choice {
	r.NonResidentRate = new(RateAndAmountFormat37Choice)
	return r.NonResidentRate
}

func (r *RateDetails26) AddApplicableRate() *RateFormat3Choice {
	r.ApplicableRate = new(RateFormat3Choice)
	return r.ApplicableRate
}

func (r *RateDetails26) AddSolicitationFeeRate() *SolicitationFeeRateFormat7Choice {
	r.SolicitationFeeRate = new(SolicitationFeeRateFormat7Choice)
	return r.SolicitationFeeRate
}

func (r *RateDetails26) AddTaxCreditRate() *TaxCreditRateFormat8Choice {
	newValue := new(TaxCreditRateFormat8Choice)
	r.TaxCreditRate = append(r.TaxCreditRate, newValue)
	return newValue
}

func (r *RateDetails26) AddTaxOnIncome() *RateAndAmountFormat37Choice {
	r.TaxOnIncome = new(RateAndAmountFormat37Choice)
	return r.TaxOnIncome
}

func (r *RateDetails26) AddTaxOnProfits() *RateFormat3Choice {
	r.TaxOnProfits = new(RateFormat3Choice)
	return r.TaxOnProfits
}

func (r *RateDetails26) AddTaxReclaimRate() *RateFormat3Choice {
	r.TaxReclaimRate = new(RateFormat3Choice)
	return r.TaxReclaimRate
}

func (r *RateDetails26) AddEqualisationRate() *RateAndAmountFormat42Choice {
	r.EqualisationRate = new(RateAndAmountFormat42Choice)
	return r.EqualisationRate
}

// Provides information about the rates related to securities movement.
type RateDetails27 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat39Choice `xml:"AddtlTax,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat39Choice `xml:"ChrgsFees,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *PercentageRate `xml:"FsclStmp,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat22Choice `xml:"GrssDvddRate,omitempty"`

	// Cash rate made available, as an incentive, in addition to the solicitation fee, in order to encourage early participation in an offer.
	EarlySolicitationFeeRate *SolicitationFeeRateFormat8Choice `xml:"EarlySlctnFeeRate,omitempty"`

	// Cash rate made available in an event in order to encourage participation in the offer. As information, Payment is made to a third party who has solicited an entity to take part in the offer.
	ThirdPartyIncentiveRate *RateAndAmountFormat39Choice `xml:"ThrdPtyIncntivRate,omitempty"`

	// Actual interest rate used for the payment of the interest for the specified interest period.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat7Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat24Choice `xml:"NetDvddRate,omitempty"`

	// Rate per share to which a non-resident is entitled.
	NonResidentRate *RateAndAmountFormat39Choice `xml:"NonResdtRate,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *PercentageRate `xml:"AplblRate,omitempty"`

	// Rate of the cash premium made available if the securities holder consents or participates to an event, for example consent fees or solicitation fee.
	SolicitationFeeRate *SolicitationFeeRateFormat8Choice `xml:"SlctnFeeRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat7Choice `xml:"TaxCdtRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxRate []*RateAndAmountFormat40Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate at which the income will be withheld by a jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate (TAXR) levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTax []*RateAndAmountFormat40Choice `xml:"ScndLvlTax,omitempty"`

	// Taxation applied on an amount clearly identified as an income.
	TaxOnIncome *RateAndAmountFormat39Choice `xml:"TaxOnIncm,omitempty"`

	// Taxation applied on an amount clearly identified as capital profits, capital gains.
	TaxOnProfits *PercentageRate `xml:"TaxOnPrfts,omitempty"`

	// Percentage of cash that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimRate *PercentageRate `xml:"TaxRclmRate,omitempty"`

	// Portion of the fund distribution which represents the average accrued income included in the purchase price for units bought during the account period.
	EqualisationRate *ActiveCurrencyAnd13DecimalAmount `xml:"EqulstnRate,omitempty"`
}

func (r *RateDetails27) AddAdditionalTax() *RateAndAmountFormat39Choice {
	r.AdditionalTax = new(RateAndAmountFormat39Choice)
	return r.AdditionalTax
}

func (r *RateDetails27) AddChargesFees() *RateAndAmountFormat39Choice {
	r.ChargesFees = new(RateAndAmountFormat39Choice)
	return r.ChargesFees
}

func (r *RateDetails27) SetFiscalStamp(value string) {
	r.FiscalStamp = (*PercentageRate)(&value)
}

func (r *RateDetails27) AddGrossDividendRate() *GrossDividendRateFormat22Choice {
	newValue := new(GrossDividendRateFormat22Choice)
	r.GrossDividendRate = append(r.GrossDividendRate, newValue)
	return newValue
}

func (r *RateDetails27) AddEarlySolicitationFeeRate() *SolicitationFeeRateFormat8Choice {
	r.EarlySolicitationFeeRate = new(SolicitationFeeRateFormat8Choice)
	return r.EarlySolicitationFeeRate
}

func (r *RateDetails27) AddThirdPartyIncentiveRate() *RateAndAmountFormat39Choice {
	r.ThirdPartyIncentiveRate = new(RateAndAmountFormat39Choice)
	return r.ThirdPartyIncentiveRate
}

func (r *RateDetails27) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat7Choice {
	newValue := new(InterestRateUsedForPaymentFormat7Choice)
	r.InterestRateUsedForPayment = append(r.InterestRateUsedForPayment, newValue)
	return newValue
}

func (r *RateDetails27) AddNetDividendRate() *NetDividendRateFormat24Choice {
	newValue := new(NetDividendRateFormat24Choice)
	r.NetDividendRate = append(r.NetDividendRate, newValue)
	return newValue
}

func (r *RateDetails27) AddNonResidentRate() *RateAndAmountFormat39Choice {
	r.NonResidentRate = new(RateAndAmountFormat39Choice)
	return r.NonResidentRate
}

func (r *RateDetails27) SetApplicableRate(value string) {
	r.ApplicableRate = (*PercentageRate)(&value)
}

func (r *RateDetails27) AddSolicitationFeeRate() *SolicitationFeeRateFormat8Choice {
	r.SolicitationFeeRate = new(SolicitationFeeRateFormat8Choice)
	return r.SolicitationFeeRate
}

func (r *RateDetails27) AddTaxCreditRate() *TaxCreditRateFormat7Choice {
	newValue := new(TaxCreditRateFormat7Choice)
	r.TaxCreditRate = append(r.TaxCreditRate, newValue)
	return newValue
}

func (r *RateDetails27) AddWithholdingTaxRate() *RateAndAmountFormat40Choice {
	newValue := new(RateAndAmountFormat40Choice)
	r.WithholdingTaxRate = append(r.WithholdingTaxRate, newValue)
	return newValue
}

func (r *RateDetails27) AddSecondLevelTax() *RateAndAmountFormat40Choice {
	newValue := new(RateAndAmountFormat40Choice)
	r.SecondLevelTax = append(r.SecondLevelTax, newValue)
	return newValue
}

func (r *RateDetails27) AddTaxOnIncome() *RateAndAmountFormat39Choice {
	r.TaxOnIncome = new(RateAndAmountFormat39Choice)
	return r.TaxOnIncome
}

func (r *RateDetails27) SetTaxOnProfits(value string) {
	r.TaxOnProfits = (*PercentageRate)(&value)
}

func (r *RateDetails27) SetTaxReclaimRate(value string) {
	r.TaxReclaimRate = (*PercentageRate)(&value)
}

func (r *RateDetails27) SetEqualisationRate(value, currency string) {
	r.EqualisationRate = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

// Provides information about the rates related to securities movement.
type RateDetails28 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat46Choice `xml:"AddtlTax,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat26Choice `xml:"GrssDvddRate,omitempty"`

	// The actual interest rate used for the payment of the interest for the specified interest period.
	// Usage guideline: It is used to provide the applicable rate for the current payment, after all calculations have been performed, that is, application of period and method of interest computation.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat10Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Percentage of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxRate []*RateAndAmountFormat47Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate at which the income will be withheld by a jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate (TAXR) levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTax []*RateAndAmountFormat47Choice `xml:"ScndLvlTax,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat46Choice `xml:"ChrgsFees,omitempty"`

	// Cash rate made available, as an incentive, in addition to the solicitation fee, in order to encourage early participation in an offer.
	EarlySolicitationFeeRate *SolicitationFeeRateFormat10Choice `xml:"EarlySlctnFeeRate,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *RateFormat3Choice `xml:"FsclStmp,omitempty"`

	// Cash rate made available in an event in order to encourage participation in the offer. As information, payment is made to a third party who has solicited an entity to take part in the offer.
	ThirdPartyIncentiveRate *RateFormat21Choice `xml:"ThrdPtyIncntivRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat28Choice `xml:"NetDvddRate,omitempty"`

	// Rate per share to which a non-resident is entitled.
	NonResidentRate *RateAndAmountFormat46Choice `xml:"NonResdtRate,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *RateFormat3Choice `xml:"AplblRate,omitempty"`

	// Rate of the cash premium made available if the securities holder consents or participates to an event, for example consent fees or solicitation fee.
	SolicitationFeeRate *SolicitationFeeRateFormat10Choice `xml:"SlctnFeeRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat10Choice `xml:"TaxCdtRate,omitempty"`

	// Overall tax withheld at source by fund managers prior to considering the tax obligation of each unit holder.
	TaxOnIncome *RateAndAmountFormat46Choice `xml:"TaxOnIncm,omitempty"`

	// Taxation applied on an amount clearly identified as capital profits, capital gains.
	TaxOnProfits *RateFormat3Choice `xml:"TaxOnPrfts,omitempty"`

	// Percentage of cash that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimRate *RateFormat3Choice `xml:"TaxRclmRate,omitempty"`

	// Portion of the fund distribution which represents the average accrued income included in the purchase price for units bought during the account period.
	EqualisationRate *RateAndAmountFormat48Choice `xml:"EqulstnRate,omitempty"`
}

func (r *RateDetails28) AddAdditionalTax() *RateAndAmountFormat46Choice {
	r.AdditionalTax = new(RateAndAmountFormat46Choice)
	return r.AdditionalTax
}

func (r *RateDetails28) AddGrossDividendRate() *GrossDividendRateFormat26Choice {
	newValue := new(GrossDividendRateFormat26Choice)
	r.GrossDividendRate = append(r.GrossDividendRate, newValue)
	return newValue
}

func (r *RateDetails28) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat10Choice {
	newValue := new(InterestRateUsedForPaymentFormat10Choice)
	r.InterestRateUsedForPayment = append(r.InterestRateUsedForPayment, newValue)
	return newValue
}

func (r *RateDetails28) AddWithholdingTaxRate() *RateAndAmountFormat47Choice {
	newValue := new(RateAndAmountFormat47Choice)
	r.WithholdingTaxRate = append(r.WithholdingTaxRate, newValue)
	return newValue
}

func (r *RateDetails28) AddSecondLevelTax() *RateAndAmountFormat47Choice {
	newValue := new(RateAndAmountFormat47Choice)
	r.SecondLevelTax = append(r.SecondLevelTax, newValue)
	return newValue
}

func (r *RateDetails28) AddChargesFees() *RateAndAmountFormat46Choice {
	r.ChargesFees = new(RateAndAmountFormat46Choice)
	return r.ChargesFees
}

func (r *RateDetails28) AddEarlySolicitationFeeRate() *SolicitationFeeRateFormat10Choice {
	r.EarlySolicitationFeeRate = new(SolicitationFeeRateFormat10Choice)
	return r.EarlySolicitationFeeRate
}

func (r *RateDetails28) AddFiscalStamp() *RateFormat3Choice {
	r.FiscalStamp = new(RateFormat3Choice)
	return r.FiscalStamp
}

func (r *RateDetails28) AddThirdPartyIncentiveRate() *RateFormat21Choice {
	r.ThirdPartyIncentiveRate = new(RateFormat21Choice)
	return r.ThirdPartyIncentiveRate
}

func (r *RateDetails28) AddNetDividendRate() *NetDividendRateFormat28Choice {
	newValue := new(NetDividendRateFormat28Choice)
	r.NetDividendRate = append(r.NetDividendRate, newValue)
	return newValue
}

func (r *RateDetails28) AddNonResidentRate() *RateAndAmountFormat46Choice {
	r.NonResidentRate = new(RateAndAmountFormat46Choice)
	return r.NonResidentRate
}

func (r *RateDetails28) AddApplicableRate() *RateFormat3Choice {
	r.ApplicableRate = new(RateFormat3Choice)
	return r.ApplicableRate
}

func (r *RateDetails28) AddSolicitationFeeRate() *SolicitationFeeRateFormat10Choice {
	r.SolicitationFeeRate = new(SolicitationFeeRateFormat10Choice)
	return r.SolicitationFeeRate
}

func (r *RateDetails28) AddTaxCreditRate() *TaxCreditRateFormat10Choice {
	newValue := new(TaxCreditRateFormat10Choice)
	r.TaxCreditRate = append(r.TaxCreditRate, newValue)
	return newValue
}

func (r *RateDetails28) AddTaxOnIncome() *RateAndAmountFormat46Choice {
	r.TaxOnIncome = new(RateAndAmountFormat46Choice)
	return r.TaxOnIncome
}

func (r *RateDetails28) AddTaxOnProfits() *RateFormat3Choice {
	r.TaxOnProfits = new(RateFormat3Choice)
	return r.TaxOnProfits
}

func (r *RateDetails28) AddTaxReclaimRate() *RateFormat3Choice {
	r.TaxReclaimRate = new(RateFormat3Choice)
	return r.TaxReclaimRate
}

func (r *RateDetails28) AddEqualisationRate() *RateAndAmountFormat48Choice {
	r.EqualisationRate = new(RateAndAmountFormat48Choice)
	return r.EqualisationRate
}

// Provides information about the rates related to securities movement.
type RateDetails3 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat14Choice `xml:"AddtlTax,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat5Choice `xml:"GrssDvddRate,omitempty"`

	// The actual interest rate used for the payment of the interest for the specified interest period.
	// Usage guideline: It is used to provide the applicable rate for the current payment, after all calculations have been performed, that is, application of period and method of interest computation.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat5Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Percentage of the gross dividend rate on which tax must be paid .
	TaxRelatedRate []*RateTypeAndAmountAndStatus6 `xml:"TaxRltdRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by a tax authority.
	WithholdingTaxRate *RateFormat6Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat14Choice `xml:"ChrgsFees,omitempty"`

	// Cash rate made available, as an incentive, in addition to the solicitation fee, in order to encourage early participation in an offer.
	EarlySolicitationFeeRate *SolicitationFeeRateFormat3Choice `xml:"EarlySlctnFeeRate,omitempty"`

	// Dividend is final.
	FinalDividendRate *RateAndAmountFormat15Choice `xml:"FnlDvddRate,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *RateFormat6Choice `xml:"FsclStmp,omitempty"`

	// Rate resulting from a fully franked dividend paid by a company; rate includes tax credit for companies that have made sufficient tax payments during fiscal period.
	FullyFrankedRate *RateAndAmountFormat14Choice `xml:"FullyFrnkdRate,omitempty"`

	// Rate of the cash premium made available if the securities holder consents or participates to an event, for example, consent fees.
	CashIncentiveRate *RateFormat6Choice `xml:"CshIncntivRate,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat5Choice `xml:"NetDvddRate,omitempty"`

	// Rate per share to which a non-resident is entitled.
	NonResidentRate *RateAndAmountFormat14Choice `xml:"NonResdtRate,omitempty"`

	// Dividend is provisional.
	ProvisionalDividendRate *RateAndAmountFormat15Choice `xml:"PrvsnlDvddRate,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *RateFormat6Choice `xml:"AplblRate,omitempty"`

	// Cash rate made available in an offer in order to encourage participation in the offer.
	SolicitationFeeRate *SolicitationFeeRateFormat3Choice `xml:"SlctnFeeRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat5Choice `xml:"TaxCdtRate,omitempty"`

	// Taxation applied on an amount clearly identified as an income.
	TaxOnIncome *RateFormat6Choice `xml:"TaxOnIncm,omitempty"`

	// Taxation applied on an amount clearly identified as capital profits, capital gains.
	TaxOnProfits *RateFormat6Choice `xml:"TaxOnPrfts,omitempty"`

	// Percentage of cash that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimRate *RateFormat6Choice `xml:"TaxRclmRate,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction in which the income was originally paid, for which relief at source and/or reclaim may be possible.
	WithholdingOfForeignTax *RateAndAmountFormat14Choice `xml:"WhldgOfFrgnTax,omitempty"`

	// Rate at which the income will be withheld by the jurisdiction in which the account owner is located, for which relief at source and/or reclaim may be possible.
	WithholdingOfLocalTax *RateAndAmountFormat14Choice `xml:"WhldgOfLclTax,omitempty"`
}

func (r *RateDetails3) AddAdditionalTax() *RateAndAmountFormat14Choice {
	r.AdditionalTax = new(RateAndAmountFormat14Choice)
	return r.AdditionalTax
}

func (r *RateDetails3) AddGrossDividendRate() *GrossDividendRateFormat5Choice {
	newValue := new(GrossDividendRateFormat5Choice)
	r.GrossDividendRate = append(r.GrossDividendRate, newValue)
	return newValue
}

func (r *RateDetails3) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat5Choice {
	newValue := new(InterestRateUsedForPaymentFormat5Choice)
	r.InterestRateUsedForPayment = append(r.InterestRateUsedForPayment, newValue)
	return newValue
}

func (r *RateDetails3) AddTaxRelatedRate() *RateTypeAndAmountAndStatus6 {
	newValue := new(RateTypeAndAmountAndStatus6)
	r.TaxRelatedRate = append(r.TaxRelatedRate, newValue)
	return newValue
}

func (r *RateDetails3) AddWithholdingTaxRate() *RateFormat6Choice {
	r.WithholdingTaxRate = new(RateFormat6Choice)
	return r.WithholdingTaxRate
}

func (r *RateDetails3) AddChargesFees() *RateAndAmountFormat14Choice {
	r.ChargesFees = new(RateAndAmountFormat14Choice)
	return r.ChargesFees
}

func (r *RateDetails3) AddEarlySolicitationFeeRate() *SolicitationFeeRateFormat3Choice {
	r.EarlySolicitationFeeRate = new(SolicitationFeeRateFormat3Choice)
	return r.EarlySolicitationFeeRate
}

func (r *RateDetails3) AddFinalDividendRate() *RateAndAmountFormat15Choice {
	r.FinalDividendRate = new(RateAndAmountFormat15Choice)
	return r.FinalDividendRate
}

func (r *RateDetails3) AddFiscalStamp() *RateFormat6Choice {
	r.FiscalStamp = new(RateFormat6Choice)
	return r.FiscalStamp
}

func (r *RateDetails3) AddFullyFrankedRate() *RateAndAmountFormat14Choice {
	r.FullyFrankedRate = new(RateAndAmountFormat14Choice)
	return r.FullyFrankedRate
}

func (r *RateDetails3) AddCashIncentiveRate() *RateFormat6Choice {
	r.CashIncentiveRate = new(RateFormat6Choice)
	return r.CashIncentiveRate
}

func (r *RateDetails3) AddNetDividendRate() *NetDividendRateFormat5Choice {
	newValue := new(NetDividendRateFormat5Choice)
	r.NetDividendRate = append(r.NetDividendRate, newValue)
	return newValue
}

func (r *RateDetails3) AddNonResidentRate() *RateAndAmountFormat14Choice {
	r.NonResidentRate = new(RateAndAmountFormat14Choice)
	return r.NonResidentRate
}

func (r *RateDetails3) AddProvisionalDividendRate() *RateAndAmountFormat15Choice {
	r.ProvisionalDividendRate = new(RateAndAmountFormat15Choice)
	return r.ProvisionalDividendRate
}

func (r *RateDetails3) AddApplicableRate() *RateFormat6Choice {
	r.ApplicableRate = new(RateFormat6Choice)
	return r.ApplicableRate
}

func (r *RateDetails3) AddSolicitationFeeRate() *SolicitationFeeRateFormat3Choice {
	r.SolicitationFeeRate = new(SolicitationFeeRateFormat3Choice)
	return r.SolicitationFeeRate
}

func (r *RateDetails3) AddTaxCreditRate() *TaxCreditRateFormat5Choice {
	newValue := new(TaxCreditRateFormat5Choice)
	r.TaxCreditRate = append(r.TaxCreditRate, newValue)
	return newValue
}

func (r *RateDetails3) AddTaxOnIncome() *RateFormat6Choice {
	r.TaxOnIncome = new(RateFormat6Choice)
	return r.TaxOnIncome
}

func (r *RateDetails3) AddTaxOnProfits() *RateFormat6Choice {
	r.TaxOnProfits = new(RateFormat6Choice)
	return r.TaxOnProfits
}

func (r *RateDetails3) AddTaxReclaimRate() *RateFormat6Choice {
	r.TaxReclaimRate = new(RateFormat6Choice)
	return r.TaxReclaimRate
}

func (r *RateDetails3) AddWithholdingOfForeignTax() *RateAndAmountFormat14Choice {
	r.WithholdingOfForeignTax = new(RateAndAmountFormat14Choice)
	return r.WithholdingOfForeignTax
}

func (r *RateDetails3) AddWithholdingOfLocalTax() *RateAndAmountFormat14Choice {
	r.WithholdingOfLocalTax = new(RateAndAmountFormat14Choice)
	return r.WithholdingOfLocalTax
}

// Provides information about the rates related to securities movement.
type RateDetails30 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat43Choice `xml:"AddtlTax,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	ChargesFees *RateAndAmountFormat43Choice `xml:"ChrgsFees,omitempty"`

	// Percentage of fiscal tax to apply.
	FiscalStamp *PercentageRate `xml:"FsclStmp,omitempty"`

	// Cash dividend amount per equity before deductions or allowances have been made.
	GrossDividendRate []*GrossDividendRateFormat24Choice `xml:"GrssDvddRate,omitempty"`

	// Cash rate made available, as an incentive, in addition to the solicitation fee, in order to encourage early participation in an offer.
	EarlySolicitationFeeRate *SolicitationFeeRateFormat9Choice `xml:"EarlySlctnFeeRate,omitempty"`

	// Cash rate made available in an event in order to encourage participation in the offer. As information, Payment is made to a third party who has solicited an entity to take part in the offer.
	ThirdPartyIncentiveRate *RateAndAmountFormat43Choice `xml:"ThrdPtyIncntivRate,omitempty"`

	// Actual interest rate used for the payment of the interest for the specified interest period.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat9Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat26Choice `xml:"NetDvddRate,omitempty"`

	// Rate per share to which a non-resident is entitled.
	NonResidentRate *RateAndAmountFormat43Choice `xml:"NonResdtRate,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *PercentageRate `xml:"AplblRate,omitempty"`

	// Rate of the cash premium made available if the securities holder consents or participates to an event, for example consent fees or solicitation fee.
	SolicitationFeeRate *SolicitationFeeRateFormat9Choice `xml:"SlctnFeeRate,omitempty"`

	// Amount of money per equity allocated as the result of a tax credit.
	TaxCreditRate []*TaxCreditRateFormat9Choice `xml:"TaxCdtRate,omitempty"`

	// Percentage of a cash distribution that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	WithholdingTaxRate []*RateAndAmountFormat45Choice `xml:"WhldgTaxRate,omitempty"`

	// Rate at which the income will be withheld by a jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible. It is levied in complement or offset of the withholding tax rate (TAXR) levied by the jurisdiction of the issuer’s tax domicile.
	SecondLevelTax []*RateAndAmountFormat45Choice `xml:"ScndLvlTax,omitempty"`

	// Taxation applied on an amount clearly identified as an income.
	TaxOnIncome *RateAndAmountFormat43Choice `xml:"TaxOnIncm,omitempty"`

	// Taxation applied on an amount clearly identified as capital profits, capital gains.
	TaxOnProfits *PercentageRate `xml:"TaxOnPrfts,omitempty"`

	// Percentage of cash that was paid in excess of actual tax obligation and was reclaimed.
	TaxReclaimRate *PercentageRate `xml:"TaxRclmRate,omitempty"`

	// Portion of the fund distribution which represents the average accrued income included in the purchase price for units bought during the account period.
	EqualisationRate *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"EqulstnRate,omitempty"`
}

func (r *RateDetails30) AddAdditionalTax() *RateAndAmountFormat43Choice {
	r.AdditionalTax = new(RateAndAmountFormat43Choice)
	return r.AdditionalTax
}

func (r *RateDetails30) AddChargesFees() *RateAndAmountFormat43Choice {
	r.ChargesFees = new(RateAndAmountFormat43Choice)
	return r.ChargesFees
}

func (r *RateDetails30) SetFiscalStamp(value string) {
	r.FiscalStamp = (*PercentageRate)(&value)
}

func (r *RateDetails30) AddGrossDividendRate() *GrossDividendRateFormat24Choice {
	newValue := new(GrossDividendRateFormat24Choice)
	r.GrossDividendRate = append(r.GrossDividendRate, newValue)
	return newValue
}

func (r *RateDetails30) AddEarlySolicitationFeeRate() *SolicitationFeeRateFormat9Choice {
	r.EarlySolicitationFeeRate = new(SolicitationFeeRateFormat9Choice)
	return r.EarlySolicitationFeeRate
}

func (r *RateDetails30) AddThirdPartyIncentiveRate() *RateAndAmountFormat43Choice {
	r.ThirdPartyIncentiveRate = new(RateAndAmountFormat43Choice)
	return r.ThirdPartyIncentiveRate
}

func (r *RateDetails30) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat9Choice {
	newValue := new(InterestRateUsedForPaymentFormat9Choice)
	r.InterestRateUsedForPayment = append(r.InterestRateUsedForPayment, newValue)
	return newValue
}

func (r *RateDetails30) AddNetDividendRate() *NetDividendRateFormat26Choice {
	newValue := new(NetDividendRateFormat26Choice)
	r.NetDividendRate = append(r.NetDividendRate, newValue)
	return newValue
}

func (r *RateDetails30) AddNonResidentRate() *RateAndAmountFormat43Choice {
	r.NonResidentRate = new(RateAndAmountFormat43Choice)
	return r.NonResidentRate
}

func (r *RateDetails30) SetApplicableRate(value string) {
	r.ApplicableRate = (*PercentageRate)(&value)
}

func (r *RateDetails30) AddSolicitationFeeRate() *SolicitationFeeRateFormat9Choice {
	r.SolicitationFeeRate = new(SolicitationFeeRateFormat9Choice)
	return r.SolicitationFeeRate
}

func (r *RateDetails30) AddTaxCreditRate() *TaxCreditRateFormat9Choice {
	newValue := new(TaxCreditRateFormat9Choice)
	r.TaxCreditRate = append(r.TaxCreditRate, newValue)
	return newValue
}

func (r *RateDetails30) AddWithholdingTaxRate() *RateAndAmountFormat45Choice {
	newValue := new(RateAndAmountFormat45Choice)
	r.WithholdingTaxRate = append(r.WithholdingTaxRate, newValue)
	return newValue
}

func (r *RateDetails30) AddSecondLevelTax() *RateAndAmountFormat45Choice {
	newValue := new(RateAndAmountFormat45Choice)
	r.SecondLevelTax = append(r.SecondLevelTax, newValue)
	return newValue
}

func (r *RateDetails30) AddTaxOnIncome() *RateAndAmountFormat43Choice {
	r.TaxOnIncome = new(RateAndAmountFormat43Choice)
	return r.TaxOnIncome
}

func (r *RateDetails30) SetTaxOnProfits(value string) {
	r.TaxOnProfits = (*PercentageRate)(&value)
}

func (r *RateDetails30) SetTaxReclaimRate(value string) {
	r.TaxReclaimRate = (*PercentageRate)(&value)
}

func (r *RateDetails30) SetEqualisationRate(value, currency string) {
	r.EqualisationRate = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

// Provides information about the rates related to securities movement.
type RateDetails7 struct {

	// Rate used for additional tax that cannot be categorised.
	AdditionalTax *RateAndAmountFormat5Choice `xml:"AddtlTax,omitempty"`

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

	// Cash rate made available, as an incentive, in addition to the solicitation fee, in order to encourage early participation in an offer.
	EarlySolicitationFeeRate *PercentageRate `xml:"EarlySlctnFeeRate,omitempty"`

	// Rate of the cash premium made available if the securities holder consents or participates to an event, for example, consent fees.
	CashIncentiveRate *PercentageRate `xml:"CshIncntivRate,omitempty"`

	// Actual interest rate used for the payment of the interest for the specified interest period.
	InterestRateUsedForPayment []*InterestRateUsedForPaymentFormat2Choice `xml:"IntrstRateUsdForPmt,omitempty"`

	// Cash dividend amount per equity after deductions or allowances have been made.
	NetDividendRate []*NetDividendRateFormat2Choice `xml:"NetDvddRate,omitempty"`

	// Rate per share to which a non-resident is entitled.
	NonResidentRate *RateAndAmountFormat5Choice `xml:"NonResdtRate,omitempty"`

	// Dividend is provisional.
	ProvisionalDividendRate *ActiveCurrencyAnd13DecimalAmount `xml:"PrvsnlDvddRate,omitempty"`

	// Rate applicable to the event announced, for example, redemption rate for a redemption event.
	ApplicableRate *PercentageRate `xml:"AplblRate,omitempty"`

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
}

func (r *RateDetails7) AddAdditionalTax() *RateAndAmountFormat5Choice {
	r.AdditionalTax = new(RateAndAmountFormat5Choice)
	return r.AdditionalTax
}

func (r *RateDetails7) AddChargesFees() *RateAndAmountFormat5Choice {
	r.ChargesFees = new(RateAndAmountFormat5Choice)
	return r.ChargesFees
}

func (r *RateDetails7) SetFinalDividendRate(value, currency string) {
	r.FinalDividendRate = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateDetails7) SetFiscalStamp(value string) {
	r.FiscalStamp = (*PercentageRate)(&value)
}

func (r *RateDetails7) AddFullyFrankedRate() *RateAndAmountFormat5Choice {
	r.FullyFrankedRate = new(RateAndAmountFormat5Choice)
	return r.FullyFrankedRate
}

func (r *RateDetails7) AddGrossDividendRate() *GrossDividendRateFormat2Choice {
	newValue := new(GrossDividendRateFormat2Choice)
	r.GrossDividendRate = append(r.GrossDividendRate, newValue)
	return newValue
}

func (r *RateDetails7) SetEarlySolicitationFeeRate(value string) {
	r.EarlySolicitationFeeRate = (*PercentageRate)(&value)
}

func (r *RateDetails7) SetCashIncentiveRate(value string) {
	r.CashIncentiveRate = (*PercentageRate)(&value)
}

func (r *RateDetails7) AddInterestRateUsedForPayment() *InterestRateUsedForPaymentFormat2Choice {
	newValue := new(InterestRateUsedForPaymentFormat2Choice)
	r.InterestRateUsedForPayment = append(r.InterestRateUsedForPayment, newValue)
	return newValue
}

func (r *RateDetails7) AddNetDividendRate() *NetDividendRateFormat2Choice {
	newValue := new(NetDividendRateFormat2Choice)
	r.NetDividendRate = append(r.NetDividendRate, newValue)
	return newValue
}

func (r *RateDetails7) AddNonResidentRate() *RateAndAmountFormat5Choice {
	r.NonResidentRate = new(RateAndAmountFormat5Choice)
	return r.NonResidentRate
}

func (r *RateDetails7) SetProvisionalDividendRate(value, currency string) {
	r.ProvisionalDividendRate = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateDetails7) SetApplicableRate(value string) {
	r.ApplicableRate = (*PercentageRate)(&value)
}

func (r *RateDetails7) SetSolicitationFeeRate(value string) {
	r.SolicitationFeeRate = (*PercentageRate)(&value)
}

func (r *RateDetails7) AddTaxCreditRate() *TaxCreditRateFormat2Choice {
	newValue := new(TaxCreditRateFormat2Choice)
	r.TaxCreditRate = append(r.TaxCreditRate, newValue)
	return newValue
}

func (r *RateDetails7) AddTaxRelatedRate() *RateTypeAndAmountAndStatus6 {
	newValue := new(RateTypeAndAmountAndStatus6)
	r.TaxRelatedRate = append(r.TaxRelatedRate, newValue)
	return newValue
}

func (r *RateDetails7) SetWithholdingTaxRate(value string) {
	r.WithholdingTaxRate = (*PercentageRate)(&value)
}

func (r *RateDetails7) SetTaxOnIncome(value string) {
	r.TaxOnIncome = (*PercentageRate)(&value)
}

func (r *RateDetails7) SetTaxOnProfits(value string) {
	r.TaxOnProfits = (*PercentageRate)(&value)
}

func (r *RateDetails7) SetTaxReclaimRate(value string) {
	r.TaxReclaimRate = (*PercentageRate)(&value)
}

func (r *RateDetails7) AddWithholdingOfForeignTax() *RateAndAmountFormat5Choice {
	r.WithholdingOfForeignTax = new(RateAndAmountFormat5Choice)
	return r.WithholdingOfForeignTax
}

func (r *RateDetails7) AddWithholdingOfLocalTax() *RateAndAmountFormat5Choice {
	r.WithholdingOfLocalTax = new(RateAndAmountFormat5Choice)
	return r.WithholdingOfLocalTax
}
