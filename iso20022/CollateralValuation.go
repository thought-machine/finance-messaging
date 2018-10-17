package iso20022

// Provides details about the valuation of each piece of collateral that is posted.
type CollateralValuation2 struct {

	// Provides the identification of the valued collateral.
	CollateralIdentification *Max35Text `xml:"CollId,omitempty"`

	// Specifies the type of collateral used.
	CollateralType *CollateralType1Code `xml:"CollTp"`

	// Provides the status of settlement of an instruction/financial instrument movement.
	SettlementStatus *SettlementStatus2Code `xml:"SttlmSts"`

	// Specifies the number of days used for interest calculation.
	NumberOfDaysAccrued *Number `xml:"NbOfDaysAcrd"`

	// Provides details on the collateral valuation.
	ValuationAmounts *CollateralAmount1 `xml:"ValtnAmts"`

	// Specifies the computation method of (accrued) interest of the security.
	DayCountBasis *InterestComputationMethod2Code `xml:"DayCntBsis"`

	// Specifies the exchange rate between the currency of the collateral and the reporting currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Specifies the haircut or valuation factor on the currency of the collateral expressed as a percentage.
	CurrencyHaircut *BaseOneRate `xml:"CcyHrcut,omitempty"`

	// Percentage by which the collateral amount needs to be adjusted.
	AdjustedRate *BaseOneRate `xml:"AdjstdRate,omitempty"`

	// Provides details on the securities collateral.
	SecuritiesCollateral *SecuritiesCollateral2 `xml:"SctiesColl,omitempty"`

	// Provides details on the cash collateral valuation.
	CashCollateral *CashCollateral4 `xml:"CshColl,omitempty"`

	// Provides details on other collateral valuation.
	OtherCollateral *OtherCollateral3 `xml:"OthrColl,omitempty"`
}

func (c *CollateralValuation2) SetCollateralIdentification(value string) {
	c.CollateralIdentification = (*Max35Text)(&value)
}

func (c *CollateralValuation2) SetCollateralType(value string) {
	c.CollateralType = (*CollateralType1Code)(&value)
}

func (c *CollateralValuation2) SetSettlementStatus(value string) {
	c.SettlementStatus = (*SettlementStatus2Code)(&value)
}

func (c *CollateralValuation2) SetNumberOfDaysAccrued(value string) {
	c.NumberOfDaysAccrued = (*Number)(&value)
}

func (c *CollateralValuation2) AddValuationAmounts() *CollateralAmount1 {
	c.ValuationAmounts = new(CollateralAmount1)
	return c.ValuationAmounts
}

func (c *CollateralValuation2) SetDayCountBasis(value string) {
	c.DayCountBasis = (*InterestComputationMethod2Code)(&value)
}

func (c *CollateralValuation2) SetExchangeRate(value string) {
	c.ExchangeRate = (*BaseOneRate)(&value)
}

func (c *CollateralValuation2) SetCurrencyHaircut(value string) {
	c.CurrencyHaircut = (*BaseOneRate)(&value)
}

func (c *CollateralValuation2) SetAdjustedRate(value string) {
	c.AdjustedRate = (*BaseOneRate)(&value)
}

func (c *CollateralValuation2) AddSecuritiesCollateral() *SecuritiesCollateral2 {
	c.SecuritiesCollateral = new(SecuritiesCollateral2)
	return c.SecuritiesCollateral
}

func (c *CollateralValuation2) AddCashCollateral() *CashCollateral4 {
	c.CashCollateral = new(CashCollateral4)
	return c.CashCollateral
}

func (c *CollateralValuation2) AddOtherCollateral() *OtherCollateral3 {
	c.OtherCollateral = new(OtherCollateral3)
	return c.OtherCollateral
}

// Provides details about the valuation of each piece of collateral that is posted.
type CollateralValuation5 struct {

	// Provides the identification of the valued collateral.
	CollateralIdentification *Max35Text `xml:"CollId,omitempty"`

	// Specifies the type of collateral used.
	CollateralType *CollateralType1Code `xml:"CollTp"`

	// Provides the status of settlement of an instruction/financial instrument movement.
	SettlementStatus *SettlementStatus2Code `xml:"SttlmSts"`

	// Specifies the number of days used for interest calculation.
	NumberOfDaysAccrued *Number `xml:"NbOfDaysAcrd"`

	// Provides details on the collateral valuation.
	ValuationAmounts *CollateralAmount1 `xml:"ValtnAmts"`

	// Specifies the computation method of (accrued) interest of the security.
	DayCountBasis *InterestComputationMethod2Code `xml:"DayCntBsis"`

	// Specifies the exchange rate between the currency of the collateral and the reporting currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Specifies the haircut or valuation factor on the currency of the collateral expressed as a percentage.
	CurrencyHaircut *BaseOneRate `xml:"CcyHrcut,omitempty"`

	// Percentage by which the collateral amount needs to be adjusted.
	AdjustedRate *BaseOneRate `xml:"AdjstdRate,omitempty"`

	// Provides details on the securities collateral.
	SecuritiesCollateral *SecuritiesCollateral6 `xml:"SctiesColl,omitempty"`

	// Provides details on the cash collateral valuation.
	CashCollateral *CashCollateral4 `xml:"CshColl,omitempty"`

	// Provides details on other collateral valuation.
	OtherCollateral *OtherCollateral6 `xml:"OthrColl,omitempty"`
}

func (c *CollateralValuation5) SetCollateralIdentification(value string) {
	c.CollateralIdentification = (*Max35Text)(&value)
}

func (c *CollateralValuation5) SetCollateralType(value string) {
	c.CollateralType = (*CollateralType1Code)(&value)
}

func (c *CollateralValuation5) SetSettlementStatus(value string) {
	c.SettlementStatus = (*SettlementStatus2Code)(&value)
}

func (c *CollateralValuation5) SetNumberOfDaysAccrued(value string) {
	c.NumberOfDaysAccrued = (*Number)(&value)
}

func (c *CollateralValuation5) AddValuationAmounts() *CollateralAmount1 {
	c.ValuationAmounts = new(CollateralAmount1)
	return c.ValuationAmounts
}

func (c *CollateralValuation5) SetDayCountBasis(value string) {
	c.DayCountBasis = (*InterestComputationMethod2Code)(&value)
}

func (c *CollateralValuation5) SetExchangeRate(value string) {
	c.ExchangeRate = (*BaseOneRate)(&value)
}

func (c *CollateralValuation5) SetCurrencyHaircut(value string) {
	c.CurrencyHaircut = (*BaseOneRate)(&value)
}

func (c *CollateralValuation5) SetAdjustedRate(value string) {
	c.AdjustedRate = (*BaseOneRate)(&value)
}

func (c *CollateralValuation5) AddSecuritiesCollateral() *SecuritiesCollateral6 {
	c.SecuritiesCollateral = new(SecuritiesCollateral6)
	return c.SecuritiesCollateral
}

func (c *CollateralValuation5) AddCashCollateral() *CashCollateral4 {
	c.CashCollateral = new(CashCollateral4)
	return c.CashCollateral
}

func (c *CollateralValuation5) AddOtherCollateral() *OtherCollateral6 {
	c.OtherCollateral = new(OtherCollateral6)
	return c.OtherCollateral
}

// Provides the valuation of a collateral, identified through an ISIN.
type CollateralValuation6 struct {

	// Nominal amount of the security pledged as collateral. Except for tri-party repos and any other transaction in which the security pledged is not identified via a single ISIN.
	NominalAmount *ActiveCurrencyAndAmount `xml:"NmnlAmt,omitempty"`

	// International Securities Identification Number (ISIN).  A numbering system designed by the United Nation's International Organisation for Standardisation (ISO). The ISIN is composed of a 2-character prefix representing the country of issue, followed by the national security number (if one exists), and a check digit. Each country has a national numbering agency that assigns ISIN numbers for securities in that country.
	ISIN *ISINOct2015Identifier `xml:"ISIN"`
}

func (c *CollateralValuation6) SetNominalAmount(value, currency string) {
	c.NominalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CollateralValuation6) SetISIN(value string) {
	c.ISIN = (*ISINOct2015Identifier)(&value)
}

// Provides the specification of the valuation of a collateral, based on the sector and the asset classification.
type CollateralValuation7 struct {

	// Specifies whether the collateral is a pool collateral or not.
	PoolStatus *CollateralPool1Code `xml:"PoolSts"`

	// Identifies the asset class pledged as collateral, expressed as an ISO 10962 Classification of Financial Instrument (CFI).
	Type *CFIOct2015Identifier `xml:"Tp"`

	// Provides the institutional sector, such as central government, central bank, etc. of the issuer of collateral.
	Sector *SNA2008SectorIdentifier `xml:"Sctr"`

	// Nominal amount of money of the security pledged as collateral, when the collateral cannot be identified through an individual or basket ISIN.
	NominalAmount *ActiveCurrencyAndAmount `xml:"NmnlAmt,omitempty"`
}

func (c *CollateralValuation7) SetPoolStatus(value string) {
	c.PoolStatus = (*CollateralPool1Code)(&value)
}

func (c *CollateralValuation7) SetType(value string) {
	c.Type = (*CFIOct2015Identifier)(&value)
}

func (c *CollateralValuation7) SetSector(value string) {
	c.Sector = (*SNA2008SectorIdentifier)(&value)
}

func (c *CollateralValuation7) SetNominalAmount(value, currency string) {
	c.NominalAmount = NewActiveCurrencyAndAmount(value, currency)
}
