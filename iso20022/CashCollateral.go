package iso20022

// Provides details about the cash posted as collateral.
type CashCollateral2 struct {

	// Provides the identification of the proposed collateral.
	CollateralIdentification *Max35Text `xml:"CollId,omitempty"`

	// Provides the identification of the clearing member 's cash account.
	CashAccountIdentification *AccountIdentification4Choice `xml:"CshAcctId,omitempty"`

	// Indicates whether excess of cash should be returned or not.
	ReturnExcess *YesNoIndicator `xml:"RtrXcss,omitempty"`

	// Amount of the deposit.
	DepositAmount *ActiveCurrencyAndAmount `xml:"DpstAmt,omitempty"`

	// Specifies whether the deposit is fixed term or call/notice.
	DepositType *DepositType1Code `xml:"DpstTp,omitempty"`

	// Planned final repayment date at the time of issuance.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Valuation date of the cash taken as collateral.
	ValueDate *ISODate `xml:"ValDt,omitempty"`

	// Exchange rate.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Value of the collateral after taking into account the haircut.
	CollateralValue *ActiveCurrencyAndAmount `xml:"CollVal"`

	// Haircut or valuation factor on the collateral expressed as a percentage.
	Haircut *PercentageRate `xml:"Hrcut,omitempty"`
}

func (c *CashCollateral2) SetCollateralIdentification(value string) {
	c.CollateralIdentification = (*Max35Text)(&value)
}

func (c *CashCollateral2) AddCashAccountIdentification() *AccountIdentification4Choice {
	c.CashAccountIdentification = new(AccountIdentification4Choice)
	return c.CashAccountIdentification
}

func (c *CashCollateral2) SetReturnExcess(value string) {
	c.ReturnExcess = (*YesNoIndicator)(&value)
}

func (c *CashCollateral2) SetDepositAmount(value, currency string) {
	c.DepositAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CashCollateral2) SetDepositType(value string) {
	c.DepositType = (*DepositType1Code)(&value)
}

func (c *CashCollateral2) SetMaturityDate(value string) {
	c.MaturityDate = (*ISODate)(&value)
}

func (c *CashCollateral2) SetValueDate(value string) {
	c.ValueDate = (*ISODate)(&value)
}

func (c *CashCollateral2) SetExchangeRate(value string) {
	c.ExchangeRate = (*BaseOneRate)(&value)
}

func (c *CashCollateral2) SetCollateralValue(value, currency string) {
	c.CollateralValue = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CashCollateral2) SetHaircut(value string) {
	c.Haircut = (*PercentageRate)(&value)
}

// Provides details about the cash posted as collateral.
type CashCollateral3 struct {

	// Provides the identification of the proposed collateral.
	CollateralIdentification *Max35Text `xml:"CollId,omitempty"`

	// Provides the identification of the clearing member 's cash account.
	CashAccountIdentification *AccountIdentification4Choice `xml:"CshAcctId,omitempty"`

	// Amount of the deposit.
	DepositAmount *ActiveCurrencyAndAmount `xml:"DpstAmt,omitempty"`

	// Specifies whether the deposit is fixed term or call/notice.
	DepositType *DepositType1Code `xml:"DpstTp,omitempty"`

	// Planned final repayment date at the time of issuance.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Valuation date of the cash taken as collateral.
	ValueDate *ISODate `xml:"ValDt,omitempty"`

	// Exchange rate.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Value of the collateral after taking into account the haircut.
	CollateralValue *ActiveCurrencyAndAmount `xml:"CollVal"`

	// Haircut or valuation factor on the collateral expressed as a percentage.
	Haircut *PercentageRate `xml:"Hrcut,omitempty"`
}

func (c *CashCollateral3) SetCollateralIdentification(value string) {
	c.CollateralIdentification = (*Max35Text)(&value)
}

func (c *CashCollateral3) AddCashAccountIdentification() *AccountIdentification4Choice {
	c.CashAccountIdentification = new(AccountIdentification4Choice)
	return c.CashAccountIdentification
}

func (c *CashCollateral3) SetDepositAmount(value, currency string) {
	c.DepositAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CashCollateral3) SetDepositType(value string) {
	c.DepositType = (*DepositType1Code)(&value)
}

func (c *CashCollateral3) SetMaturityDate(value string) {
	c.MaturityDate = (*ISODate)(&value)
}

func (c *CashCollateral3) SetValueDate(value string) {
	c.ValueDate = (*ISODate)(&value)
}

func (c *CashCollateral3) SetExchangeRate(value string) {
	c.ExchangeRate = (*BaseOneRate)(&value)
}

func (c *CashCollateral3) SetCollateralValue(value, currency string) {
	c.CollateralValue = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CashCollateral3) SetHaircut(value string) {
	c.Haircut = (*PercentageRate)(&value)
}

// Provides details about the cash posted as collateral.
type CashCollateral4 struct {

	// Identifies the register number of the collateral deposit assigned by the central counterparty.
	AssetNumber *Max35Text `xml:"AsstNb,omitempty"`

	// Amount of the deposit.
	DepositAmount *ActiveCurrencyAndAmount `xml:"DpstAmt,omitempty"`

	// Specifies whether the deposit is fixed term or call/notice.
	DepositType *DepositType1Code `xml:"DpstTp,omitempty"`

	// Amount blocked by the central counterparty for any reasonable reason ( for example for judicial reasons). In this case the investor can not withdraw or distribute this collateral.
	BlockedAmount *ActiveCurrencyAndAmount `xml:"BlckdAmt,omitempty"`

	// Planned final repayment date at the time of issuance.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Valuation date of the cash taken as collateral.
	ValueDate *ISODate `xml:"ValDt,omitempty"`

	// Exchange rate.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Value of the collateral after taking into account the haircut.
	CollateralValue *ActiveCurrencyAndAmount `xml:"CollVal"`

	// Haircut or valuation factor on the collateral expressed as a percentage.
	Haircut *PercentageRate `xml:"Hrcut,omitempty"`
}

func (c *CashCollateral4) SetAssetNumber(value string) {
	c.AssetNumber = (*Max35Text)(&value)
}

func (c *CashCollateral4) SetDepositAmount(value, currency string) {
	c.DepositAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CashCollateral4) SetDepositType(value string) {
	c.DepositType = (*DepositType1Code)(&value)
}

func (c *CashCollateral4) SetBlockedAmount(value, currency string) {
	c.BlockedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CashCollateral4) SetMaturityDate(value string) {
	c.MaturityDate = (*ISODate)(&value)
}

func (c *CashCollateral4) SetValueDate(value string) {
	c.ValueDate = (*ISODate)(&value)
}

func (c *CashCollateral4) SetExchangeRate(value string) {
	c.ExchangeRate = (*BaseOneRate)(&value)
}

func (c *CashCollateral4) SetCollateralValue(value, currency string) {
	c.CollateralValue = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CashCollateral4) SetHaircut(value string) {
	c.Haircut = (*PercentageRate)(&value)
}

// Provides details about the cash posted as collateral.
type CashCollateral5 struct {

	// Provides the identification of the proposed collateral.
	CollateralIdentification *Max35Text `xml:"CollId,omitempty"`

	// Provides the identification of the clearing member 's cash account.
	CashAccountIdentification *AccountIdentification4Choice `xml:"CshAcctId,omitempty"`

	// Identifies the register number of the collateral deposit assigned by the central counterparty.
	AssetNumber *Max35Text `xml:"AsstNb,omitempty"`

	// Amount of the deposit.
	DepositAmount *ActiveCurrencyAndAmount `xml:"DpstAmt,omitempty"`

	// Specifies whether the deposit is fixed term or call/notice.
	DepositType *DepositType1Code `xml:"DpstTp,omitempty"`

	// Planned final repayment date at the time of issuance.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Valuation date of the cash taken as collateral.
	ValueDate *ISODate `xml:"ValDt,omitempty"`

	// Exchange rate.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Value of the collateral after taking into account the haircut.
	CollateralValue *ActiveCurrencyAndAmount `xml:"CollVal"`

	// Haircut or valuation factor on the collateral expressed as a percentage.
	Haircut *PercentageRate `xml:"Hrcut,omitempty"`
}

func (c *CashCollateral5) SetCollateralIdentification(value string) {
	c.CollateralIdentification = (*Max35Text)(&value)
}

func (c *CashCollateral5) AddCashAccountIdentification() *AccountIdentification4Choice {
	c.CashAccountIdentification = new(AccountIdentification4Choice)
	return c.CashAccountIdentification
}

func (c *CashCollateral5) SetAssetNumber(value string) {
	c.AssetNumber = (*Max35Text)(&value)
}

func (c *CashCollateral5) SetDepositAmount(value, currency string) {
	c.DepositAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CashCollateral5) SetDepositType(value string) {
	c.DepositType = (*DepositType1Code)(&value)
}

func (c *CashCollateral5) SetMaturityDate(value string) {
	c.MaturityDate = (*ISODate)(&value)
}

func (c *CashCollateral5) SetValueDate(value string) {
	c.ValueDate = (*ISODate)(&value)
}

func (c *CashCollateral5) SetExchangeRate(value string) {
	c.ExchangeRate = (*BaseOneRate)(&value)
}

func (c *CashCollateral5) SetCollateralValue(value, currency string) {
	c.CollateralValue = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CashCollateral5) SetHaircut(value string) {
	c.Haircut = (*PercentageRate)(&value)
}
