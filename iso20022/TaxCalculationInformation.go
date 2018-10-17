package iso20022

// Information used to calculate the tax.
type TaxCalculationInformation1 struct {

	// Basis used to determine the capital gain or loss, eg, the purchase price.
	Basis *TaxationBasis1 `xml:"Bsis,omitempty"`
}

func (t *TaxCalculationInformation1) AddBasis() *TaxationBasis1 {
	t.Basis = new(TaxationBasis1)
	return t.Basis
}

// Information used to calculate the tax.
type TaxCalculationInformation10 struct {

	// Form of the rebate, for example, cash.
	Basis *TaxBasis1Choice `xml:"Bsis,omitempty"`

	// Amount of money on which the tax is charged.
	TaxableAmount *ActiveCurrencyAndAmount `xml:"TaxblAmt"`
}

func (t *TaxCalculationInformation10) AddBasis() *TaxBasis1Choice {
	t.Basis = new(TaxBasis1Choice)
	return t.Basis
}

func (t *TaxCalculationInformation10) SetTaxableAmount(value, currency string) {
	t.TaxableAmount = NewActiveCurrencyAndAmount(value, currency)
}

// Information used to calculate the tax.
type TaxCalculationInformation2 struct {

	// Specifies whether capital gain is in the scope of the European directive on taxation of savings income in the form of interest payments (Council Directive 2003/48/EC 3 June), or an income realised upon sale, a refund or redemption of shares and units, etc.
	EUCapitalGain *EUCapitalGain1 `xml:"EUCptlGn,omitempty"`

	// Percentage of the underlying assets of the funds that represents a debt and is in the scope of the European directive on taxation of savings income in the form of interest payments (Council Directive 2003/48/EC 3 June).
	PercentageOfDebtClaim *PercentageRate `xml:"PctgOfDebtClm,omitempty"`

	// Percentage of grandfathered debt claim.
	PercentageGrandfatheredDebt *PercentageRate `xml:"PctgGrdfthdDebt,omitempty"`
}

func (t *TaxCalculationInformation2) AddEUCapitalGain() *EUCapitalGain1 {
	t.EUCapitalGain = new(EUCapitalGain1)
	return t.EUCapitalGain
}

func (t *TaxCalculationInformation2) SetPercentageOfDebtClaim(value string) {
	t.PercentageOfDebtClaim = (*PercentageRate)(&value)
}

func (t *TaxCalculationInformation2) SetPercentageGrandfatheredDebt(value string) {
	t.PercentageGrandfatheredDebt = (*PercentageRate)(&value)
}

// Information used to calculate the tax.
type TaxCalculationInformation3 struct {

	// Basis used to determine the capital gain or loss, eg, the purchase price.
	Basis *TaxationBasis1 `xml:"Bsis,omitempty"`

	// Specifies whether capital gain is in the scope of the European directive on taxation of savings income in the form of interest payments (Council Directive 2003/48/EC 3 June), or an income realised upon sale, a refund or redemption of shares and units, etc.
	EUCapitalGain *EUCapitalGain1 `xml:"EUCptlGn,omitempty"`

	// Amount of money that it is to be taxed.
	TaxableAmount *ActiveCurrencyAnd13DecimalAmount `xml:"TaxblAmt,omitempty"`
}

func (t *TaxCalculationInformation3) AddBasis() *TaxationBasis1 {
	t.Basis = new(TaxationBasis1)
	return t.Basis
}

func (t *TaxCalculationInformation3) AddEUCapitalGain() *EUCapitalGain1 {
	t.EUCapitalGain = new(EUCapitalGain1)
	return t.EUCapitalGain
}

func (t *TaxCalculationInformation3) SetTaxableAmount(value, currency string) {
	t.TaxableAmount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

// Information used to calculate the tax.
type TaxCalculationInformation4 struct {

	// Specifies whether capital gain is in the scope of the European directive on taxation of savings income in the form of interest payments (Council Directive 2003/48/EC 3 June), or an income realised upon sale, a refund or redemption of shares and units, etc.
	EUCapitalGain *EUCapitalGain2Code `xml:"EUCptlGn,omitempty"`

	// Specifies whether capital gain is in the scope of the European directive on taxation of savings income in the form of interest payments (Council Directive 2003/48/EC 3 June), or an income realised upon sale, a refund or redemption of shares and units, etc.
	ExtendedEUCapitalGain *Extended350Code `xml:"XtndedEUCptlGn,omitempty"`

	// Percentage of the underlying assets of the funds that represents a debt and is in the scope of the European directive on taxation of savings income in the form of interest payments (Council Directive 2003/48/EC 3 June).
	PercentageOfDebtClaim *PercentageRate `xml:"PctgOfDebtClm,omitempty"`

	// Percentage of grandfathered debt claim.
	PercentageGrandfatheredDebt *PercentageRate `xml:"PctgGrdfthdDebt,omitempty"`

	// Amount included in the dividend that corresponds to gains directly or indirectly derived from interest payment in the scope of the European Directive on taxation of savings income in the form of interest payments.
	TaxableIncomePerDividend *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"TaxblIncmPerDvdd,omitempty"`

	// Specifies whether dividend is in the scope of the European directive on taxation of savings income in the form of interest payments (Council Directive 2003/48/EC 3 June), or an income realised upon sale, a refund or redemption of shares and units, etc.
	EUDividendStatus *EUDividendStatus1Code `xml:"EUDvddSts,omitempty"`

	// Specifies whether dividend is in the scope of the European directive on taxation of savings income in the form of interest payments (Council Directive 2003/48/EC 3 June), or an income realised upon sale, a refund or redemption of shares and units, etc.
	ExtendedEUDividendStatus *Extended350Code `xml:"XtndedEUDvddSts,omitempty"`
}

func (t *TaxCalculationInformation4) SetEUCapitalGain(value string) {
	t.EUCapitalGain = (*EUCapitalGain2Code)(&value)
}

func (t *TaxCalculationInformation4) SetExtendedEUCapitalGain(value string) {
	t.ExtendedEUCapitalGain = (*Extended350Code)(&value)
}

func (t *TaxCalculationInformation4) SetPercentageOfDebtClaim(value string) {
	t.PercentageOfDebtClaim = (*PercentageRate)(&value)
}

func (t *TaxCalculationInformation4) SetPercentageGrandfatheredDebt(value string) {
	t.PercentageGrandfatheredDebt = (*PercentageRate)(&value)
}

func (t *TaxCalculationInformation4) SetTaxableIncomePerDividend(value, currency string) {
	t.TaxableIncomePerDividend = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *TaxCalculationInformation4) SetEUDividendStatus(value string) {
	t.EUDividendStatus = (*EUDividendStatus1Code)(&value)
}

func (t *TaxCalculationInformation4) SetExtendedEUDividendStatus(value string) {
	t.ExtendedEUDividendStatus = (*Extended350Code)(&value)
}

// Information used to calculate the tax.
type TaxCalculationInformation5 struct {

	// Basis used to determine the capital gain or loss, eg, the purchase price.
	Basis *TaxationBasis2Code `xml:"Bsis,omitempty"`

	// Basis used to determine the capital gain or loss, eg, the purchase price.
	ExtendedBasis *Extended350Code `xml:"XtndedBsis,omitempty"`
}

func (t *TaxCalculationInformation5) SetBasis(value string) {
	t.Basis = (*TaxationBasis2Code)(&value)
}

func (t *TaxCalculationInformation5) SetExtendedBasis(value string) {
	t.ExtendedBasis = (*Extended350Code)(&value)
}

// Information used to calculate the tax.
type TaxCalculationInformation6 struct {

	// Basis used to determine the capital gain or loss, eg, the purchase price.
	Basis *TaxationBasis2Code `xml:"Bsis,omitempty"`

	// Basis used to determine the capital gain or loss, eg, the purchase price.
	ExtendedBasis *Extended350Code `xml:"XtndedBsis,omitempty"`

	// Amount of money that it is to be taxed.
	TaxableAmount *ActiveCurrencyAnd13DecimalAmount `xml:"TaxblAmt,omitempty"`
}

func (t *TaxCalculationInformation6) SetBasis(value string) {
	t.Basis = (*TaxationBasis2Code)(&value)
}

func (t *TaxCalculationInformation6) SetExtendedBasis(value string) {
	t.ExtendedBasis = (*Extended350Code)(&value)
}

func (t *TaxCalculationInformation6) SetTaxableAmount(value, currency string) {
	t.TaxableAmount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

// Information used to calculate the tax.
type TaxCalculationInformation8 struct {

	// Form of the rebate, for example, cash.
	Basis *TaxBasis1Choice `xml:"Bsis,omitempty"`

	// Amount of money on which the tax is charged.
	TaxableAmount *ActiveCurrencyAnd13DecimalAmount `xml:"TaxblAmt"`
}

func (t *TaxCalculationInformation8) AddBasis() *TaxBasis1Choice {
	t.Basis = new(TaxBasis1Choice)
	return t.Basis
}

func (t *TaxCalculationInformation8) SetTaxableAmount(value, currency string) {
	t.TaxableAmount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

// Information used to calculate the tax.
type TaxCalculationInformation9 struct {

	// Form of the rebate, for example, cash.
	Basis *TaxBasis1Choice `xml:"Bsis"`
}

func (t *TaxCalculationInformation9) AddBasis() *TaxBasis1Choice {
	t.Basis = new(TaxBasis1Choice)
	return t.Basis
}
