package iso20022

// Provides information about the tax voucher.
type TaxVoucher1 struct {

	// Distribution rate per share.
	TaxVoucherRate *BaseOneRate `xml:"TaxVchrRate"`

	// Amount of tax that have been previously paid.
	TaxCredit *ActiveCurrencyAndAmount `xml:"TaxCdt"`

	// Amount of tax that have been previously deducted.
	TaxDeduction *ActiveCurrencyAndAmount `xml:"TaxDdctn"`

	// Cash amount before any deductions and allowances have been made.
	GrossAmount *ActiveCurrencyAndAmount `xml:"GrssAmt"`

	// Cash amount after any deductions and allowances have been made
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt"`

	// Securities holding on record date.
	RecordDateHolding *UnitOrFaceAmount1Choice `xml:"RcrdDtHldg"`

	// Applicable tax rate on the tax credit amount
	TaxCreditRate *PercentageRate `xml:"TaxCdtRate,omitempty"`

	// Cash amount that will be withheld by a tax authority.
	WithholdingTaxAmount *ActiveCurrencyAndAmount `xml:"WhldgTaxAmt,omitempty"`

	// Rate of a cash distribution that wil be withheld by a tax authority
	WithholdingTaxRate *PercentageRate `xml:"WhldgTaxRate,omitempty"`

	// Cost per share of new shares allotted.
	ScripDividendReinvestmentPricePerShare *PriceValue1 `xml:"ScripDvddRinvstmtPricPerShr,omitempty"`

	// Cash amount retained from previous dividend or interest payment.
	CashAmountBroughtForward *ActiveCurrencyAndAmount `xml:"CshAmtBrghtFwd,omitempty"`

	// Total cash amount required to purchase shares allotted.
	AllotedSharesCost *PriceValue1 `xml:"AlltdShrsCost,omitempty"`

	// Cash amount carried forward to next dividend or interest payment.
	CashAmountCarriedForward *ActiveCurrencyAndAmount `xml:"CshAmtCrrdFwd,omitempty"`

	// Where new securities are issued in lieu of a cash dividend, the notional tax is the tax on the amount of cash that would have been paid. For scrips only.
	NotionalTax *ActiveCurrencyAndAmount `xml:"NtnlTax,omitempty"`

	// Amount of cash that would have been payable if the dividend had been taken in the form of cash rather than shares. For scrip only.
	NotionalDividendPayable *ActiveCurrencyAndAmount `xml:"NtnlDvddPybl,omitempty"`

	// Date on which DRIP purchase completed.
	BargainDate *ISODate `xml:"BrgnDt,omitempty"`

	// Settlement date of the DRIP purchase transaction.
	BargainSettlementDate *ISODate `xml:"BrgnSttlmDt,omitempty"`

	// Amount of stamp duty.
	StampDutyAmount *ActiveCurrencyAndAmount `xml:"StmpDtyAmt,omitempty"`

	// Amount of charges/fees charged to the client.
	ChargeAmount *ActiveCurrencyAndAmount `xml:"ChrgAmt,omitempty"`

	// Amount due to the paying agent.
	CommissionAmount *ActiveCurrencyAndAmount `xml:"ComssnAmt,omitempty"`

	// Provides information about the foreign exchange transaction.
	ForeignExchangeDetails *ForeignExchangeTerms9 `xml:"FXDtls,omitempty"`
}

func (t *TaxVoucher1) SetTaxVoucherRate(value string) {
	t.TaxVoucherRate = (*BaseOneRate)(&value)
}

func (t *TaxVoucher1) SetTaxCredit(value, currency string) {
	t.TaxCredit = NewActiveCurrencyAndAmount(value, currency)
}

func (t *TaxVoucher1) SetTaxDeduction(value, currency string) {
	t.TaxDeduction = NewActiveCurrencyAndAmount(value, currency)
}

func (t *TaxVoucher1) SetGrossAmount(value, currency string) {
	t.GrossAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (t *TaxVoucher1) SetNetAmount(value, currency string) {
	t.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (t *TaxVoucher1) AddRecordDateHolding() *UnitOrFaceAmount1Choice {
	t.RecordDateHolding = new(UnitOrFaceAmount1Choice)
	return t.RecordDateHolding
}

func (t *TaxVoucher1) SetTaxCreditRate(value string) {
	t.TaxCreditRate = (*PercentageRate)(&value)
}

func (t *TaxVoucher1) SetWithholdingTaxAmount(value, currency string) {
	t.WithholdingTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (t *TaxVoucher1) SetWithholdingTaxRate(value string) {
	t.WithholdingTaxRate = (*PercentageRate)(&value)
}

func (t *TaxVoucher1) AddScripDividendReinvestmentPricePerShare() *PriceValue1 {
	t.ScripDividendReinvestmentPricePerShare = new(PriceValue1)
	return t.ScripDividendReinvestmentPricePerShare
}

func (t *TaxVoucher1) SetCashAmountBroughtForward(value, currency string) {
	t.CashAmountBroughtForward = NewActiveCurrencyAndAmount(value, currency)
}

func (t *TaxVoucher1) AddAllotedSharesCost() *PriceValue1 {
	t.AllotedSharesCost = new(PriceValue1)
	return t.AllotedSharesCost
}

func (t *TaxVoucher1) SetCashAmountCarriedForward(value, currency string) {
	t.CashAmountCarriedForward = NewActiveCurrencyAndAmount(value, currency)
}

func (t *TaxVoucher1) SetNotionalTax(value, currency string) {
	t.NotionalTax = NewActiveCurrencyAndAmount(value, currency)
}

func (t *TaxVoucher1) SetNotionalDividendPayable(value, currency string) {
	t.NotionalDividendPayable = NewActiveCurrencyAndAmount(value, currency)
}

func (t *TaxVoucher1) SetBargainDate(value string) {
	t.BargainDate = (*ISODate)(&value)
}

func (t *TaxVoucher1) SetBargainSettlementDate(value string) {
	t.BargainSettlementDate = (*ISODate)(&value)
}

func (t *TaxVoucher1) SetStampDutyAmount(value, currency string) {
	t.StampDutyAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (t *TaxVoucher1) SetChargeAmount(value, currency string) {
	t.ChargeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (t *TaxVoucher1) SetCommissionAmount(value, currency string) {
	t.CommissionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (t *TaxVoucher1) AddForeignExchangeDetails() *ForeignExchangeTerms9 {
	t.ForeignExchangeDetails = new(ForeignExchangeTerms9)
	return t.ForeignExchangeDetails
}

// Specifies tax vouchers in the framework of a corporate action event.
type TaxVoucher2 struct {

	// Unique reference for the tax voucher required by the relevant tax authorities to ensure that any claim relating to this particular tax voucher cannot be duplicated.
	Identification *Max35Text `xml:"Id"`

	// Date on which a dividend reinvestment purchase was completed. If there is only one bargain involved, the time it was struck needs to be included.
	BargainDate *DateAndDateTimeChoice `xml:"BrgnDt,omitempty"`

	// Settlement date of the dividend reinvestment purchase transaction.
	BargainSettlementDate *DateAndDateTimeChoice `xml:"BrgnSttlmDt,omitempty"`
}

func (t *TaxVoucher2) SetIdentification(value string) {
	t.Identification = (*Max35Text)(&value)
}

func (t *TaxVoucher2) AddBargainDate() *DateAndDateTimeChoice {
	t.BargainDate = new(DateAndDateTimeChoice)
	return t.BargainDate
}

func (t *TaxVoucher2) AddBargainSettlementDate() *DateAndDateTimeChoice {
	t.BargainSettlementDate = new(DateAndDateTimeChoice)
	return t.BargainSettlementDate
}

// Specifies tax vouchers in the framework of a corporate action event.
type TaxVoucher3 struct {

	// Unique reference for the tax voucher required by the relevant tax authorities to ensure that any claim relating to this particular tax voucher cannot be duplicated.
	Identification *RestrictedFINXMax16Text `xml:"Id"`

	// Date on which a dividend reinvestment purchase was completed. If there is only one bargain involved, the time it was struck needs to be included.
	BargainDate *DateAndDateTimeChoice `xml:"BrgnDt,omitempty"`

	// Settlement date of the dividend reinvestment purchase transaction.
	BargainSettlementDate *DateAndDateTimeChoice `xml:"BrgnSttlmDt,omitempty"`
}

func (t *TaxVoucher3) SetIdentification(value string) {
	t.Identification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TaxVoucher3) AddBargainDate() *DateAndDateTimeChoice {
	t.BargainDate = new(DateAndDateTimeChoice)
	return t.BargainDate
}

func (t *TaxVoucher3) AddBargainSettlementDate() *DateAndDateTimeChoice {
	t.BargainSettlementDate = new(DateAndDateTimeChoice)
	return t.BargainSettlementDate
}
