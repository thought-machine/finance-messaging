package iso20022

// Provides information about the cash option.
type CashOption1 struct {

	// Indicates wether it is a debit or a credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Currency of the option.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate5 `xml:"DtDtls,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts1 `xml:"AmtDtls,omitempty"`

	// Provides information about a foreign exchange.
	ExchangeRate *ForeignExchangeTerms8 `xml:"XchgRate,omitempty"`
}

func (c *CashOption1) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption1) SetCurrency(value string) {
	c.Currency = (*ActiveCurrencyCode)(&value)
}

func (c *CashOption1) AddDateDetails() *CorporateActionDate5 {
	c.DateDetails = new(CorporateActionDate5)
	return c.DateDetails
}

func (c *CashOption1) AddAmountDetails() *CorporateActionAmounts1 {
	c.AmountDetails = new(CorporateActionAmounts1)
	return c.AmountDetails
}

func (c *CashOption1) AddExchangeRate() *ForeignExchangeTerms8 {
	c.ExchangeRate = new(ForeignExchangeTerms8)
	return c.ExchangeRate
}

// Provides information about the cash option.
type CashOption10 struct {

	// Indicates whether the value is a debit or a credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator1Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification20 `xml:"IncmTp,omitempty"`

	// Identification of the account in which cash is maintained.
	CashAccountIdentification *CashAccountIdentification5Choice `xml:"CshAcctId,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts10 `xml:"AmtDtls,omitempty"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate17 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount
	ForeignExchangeDetails *ForeignExchangeTerms13 `xml:"FXDtls,omitempty"`

	// Provides information about the corporate action option.
	RateAndAmountDetails *RateDetails3 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *PriceDetails3 `xml:"PricDtls,omitempty"`
}

func (c *CashOption10) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption10) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator1Choice {
	c.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator1Choice)
	return c.NonEligibleProceedsIndicator
}

func (c *CashOption10) AddIncomeType() *GenericIdentification20 {
	c.IncomeType = new(GenericIdentification20)
	return c.IncomeType
}

func (c *CashOption10) AddCashAccountIdentification() *CashAccountIdentification5Choice {
	c.CashAccountIdentification = new(CashAccountIdentification5Choice)
	return c.CashAccountIdentification
}

func (c *CashOption10) AddAmountDetails() *CorporateActionAmounts10 {
	c.AmountDetails = new(CorporateActionAmounts10)
	return c.AmountDetails
}

func (c *CashOption10) AddDateDetails() *CorporateActionDate17 {
	c.DateDetails = new(CorporateActionDate17)
	return c.DateDetails
}

func (c *CashOption10) AddForeignExchangeDetails() *ForeignExchangeTerms13 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms13)
	return c.ForeignExchangeDetails
}

func (c *CashOption10) AddRateAndAmountDetails() *RateDetails3 {
	c.RateAndAmountDetails = new(RateDetails3)
	return c.RateAndAmountDetails
}

func (c *CashOption10) AddPriceDetails() *PriceDetails3 {
	c.PriceDetails = new(PriceDetails3)
	return c.PriceDetails
}

// Provides information about the cash option.
type CashOption11 struct {

	// Indicates whether the value is a debit or a credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates whether the cash payment occurs or will occur in advance of receipt of proceeds from the issuer and based on a contractual agreement established with the account servicer or upon receipt of proceeds from the issuer.
	ContractualPaymentIndicator *Payment1Code `xml:"CtrctlPmtInd,omitempty"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator1Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification20 `xml:"IncmTp,omitempty"`

	// Identification of the account in which cash is maintained.
	CashAccountIdentification *CashAccountIdentification5Choice `xml:"CshAcctId,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts9 `xml:"AmtDtls,omitempty"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate17 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount.
	ForeignExchangeDetails *ForeignExchangeTerms13 `xml:"FXDtls,omitempty"`

	// Provides information about the corporate action option.
	RateAndAmountDetails *RateDetails3 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *PriceDetails3 `xml:"PricDtls,omitempty"`
}

func (c *CashOption11) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption11) SetContractualPaymentIndicator(value string) {
	c.ContractualPaymentIndicator = (*Payment1Code)(&value)
}

func (c *CashOption11) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator1Choice {
	c.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator1Choice)
	return c.NonEligibleProceedsIndicator
}

func (c *CashOption11) AddIncomeType() *GenericIdentification20 {
	c.IncomeType = new(GenericIdentification20)
	return c.IncomeType
}

func (c *CashOption11) AddCashAccountIdentification() *CashAccountIdentification5Choice {
	c.CashAccountIdentification = new(CashAccountIdentification5Choice)
	return c.CashAccountIdentification
}

func (c *CashOption11) AddAmountDetails() *CorporateActionAmounts9 {
	c.AmountDetails = new(CorporateActionAmounts9)
	return c.AmountDetails
}

func (c *CashOption11) AddDateDetails() *CorporateActionDate17 {
	c.DateDetails = new(CorporateActionDate17)
	return c.DateDetails
}

func (c *CashOption11) AddForeignExchangeDetails() *ForeignExchangeTerms13 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms13)
	return c.ForeignExchangeDetails
}

func (c *CashOption11) AddRateAndAmountDetails() *RateDetails3 {
	c.RateAndAmountDetails = new(RateDetails3)
	return c.RateAndAmountDetails
}

func (c *CashOption11) AddPriceDetails() *PriceDetails3 {
	c.PriceDetails = new(PriceDetails3)
	return c.PriceDetails
}

// Provides information about the cash option.
type CashOption12 struct {

	// Indicates whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates whether the cash payment occurs or will occur in advance of receipt of proceeds from the issuer and based on a contractual agreement established with the account servicer or upon receipt of proceeds from the issuer.
	ContractualPaymentIndicator *Payment1Code `xml:"CtrctlPmtInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification20 `xml:"IncmTp,omitempty"`

	// Choice between a cash account, a charges account or a tax account.
	Account *Account8Choice `xml:"Acct,omitempty"`

	// Provides information about cash parties.
	CashParties *CashParties10 `xml:"CshPties,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts11 `xml:"AmtDtls"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate7 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount.
	ForeignExchangeDetails *ForeignExchangeTerms11 `xml:"FXDtls,omitempty"`

	// Provides information about the tax voucher related to a cash movement.
	TaxVoucherDetails *TaxVoucher2 `xml:"TaxVchrDtls,omitempty"`

	// Provides information about the corporate action option.
	RateAndAmountDetails *RateDetails2 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *PriceDetails2 `xml:"PricDtls,omitempty"`
}

func (c *CashOption12) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption12) SetContractualPaymentIndicator(value string) {
	c.ContractualPaymentIndicator = (*Payment1Code)(&value)
}

func (c *CashOption12) AddIncomeType() *GenericIdentification20 {
	c.IncomeType = new(GenericIdentification20)
	return c.IncomeType
}

func (c *CashOption12) AddAccount() *Account8Choice {
	c.Account = new(Account8Choice)
	return c.Account
}

func (c *CashOption12) AddCashParties() *CashParties10 {
	c.CashParties = new(CashParties10)
	return c.CashParties
}

func (c *CashOption12) AddAmountDetails() *CorporateActionAmounts11 {
	c.AmountDetails = new(CorporateActionAmounts11)
	return c.AmountDetails
}

func (c *CashOption12) AddDateDetails() *CorporateActionDate7 {
	c.DateDetails = new(CorporateActionDate7)
	return c.DateDetails
}

func (c *CashOption12) AddForeignExchangeDetails() *ForeignExchangeTerms11 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms11)
	return c.ForeignExchangeDetails
}

func (c *CashOption12) AddTaxVoucherDetails() *TaxVoucher2 {
	c.TaxVoucherDetails = new(TaxVoucher2)
	return c.TaxVoucherDetails
}

func (c *CashOption12) AddRateAndAmountDetails() *RateDetails2 {
	c.RateAndAmountDetails = new(RateDetails2)
	return c.RateAndAmountDetails
}

func (c *CashOption12) AddPriceDetails() *PriceDetails2 {
	c.PriceDetails = new(PriceDetails2)
	return c.PriceDetails
}

// Provides information about the cash option.
type CashOption16 struct {

	// Indicates whether the value is a debit or a credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates whether the cash payment occurs or will occur in advance of receipt of proceeds from the issuer and based on a contractual agreement established with the account servicer or upon receipt of proceeds from the issuer.
	ContractualPaymentIndicator *Payment1Code `xml:"CtrctlPmtInd,omitempty"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator1Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *IssuerTaxability1Code `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification20 `xml:"IncmTp,omitempty"`

	// Identification of the account in which cash is maintained.
	CashAccountIdentification *CashAccountIdentification5Choice `xml:"CshAcctId,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts16 `xml:"AmtDtls,omitempty"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate23 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount.
	ForeignExchangeDetails *ForeignExchangeTerms13 `xml:"FXDtls,omitempty"`

	// Provides information about the corporate action option.
	RateAndAmountDetails *RateDetails3 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *PriceDetails6 `xml:"PricDtls,omitempty"`
}

func (c *CashOption16) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption16) SetContractualPaymentIndicator(value string) {
	c.ContractualPaymentIndicator = (*Payment1Code)(&value)
}

func (c *CashOption16) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator1Choice {
	c.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator1Choice)
	return c.NonEligibleProceedsIndicator
}

func (c *CashOption16) SetIssuerOfferorTaxabilityIndicator(value string) {
	c.IssuerOfferorTaxabilityIndicator = (*IssuerTaxability1Code)(&value)
}

func (c *CashOption16) AddIncomeType() *GenericIdentification20 {
	c.IncomeType = new(GenericIdentification20)
	return c.IncomeType
}

func (c *CashOption16) AddCashAccountIdentification() *CashAccountIdentification5Choice {
	c.CashAccountIdentification = new(CashAccountIdentification5Choice)
	return c.CashAccountIdentification
}

func (c *CashOption16) AddAmountDetails() *CorporateActionAmounts16 {
	c.AmountDetails = new(CorporateActionAmounts16)
	return c.AmountDetails
}

func (c *CashOption16) AddDateDetails() *CorporateActionDate23 {
	c.DateDetails = new(CorporateActionDate23)
	return c.DateDetails
}

func (c *CashOption16) AddForeignExchangeDetails() *ForeignExchangeTerms13 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms13)
	return c.ForeignExchangeDetails
}

func (c *CashOption16) AddRateAndAmountDetails() *RateDetails3 {
	c.RateAndAmountDetails = new(RateDetails3)
	return c.RateAndAmountDetails
}

func (c *CashOption16) AddPriceDetails() *PriceDetails6 {
	c.PriceDetails = new(PriceDetails6)
	return c.PriceDetails
}

// Provides information about the cash option.
type CashOption17 struct {

	// Indicates whether the value is a debit or a credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator1Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *IssuerTaxability1Code `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification20 `xml:"IncmTp,omitempty"`

	// Identification of the account in which cash is maintained.
	CashAccountIdentification *CashAccountIdentification5Choice `xml:"CshAcctId,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts15 `xml:"AmtDtls,omitempty"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate23 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount
	ForeignExchangeDetails *ForeignExchangeTerms13 `xml:"FXDtls,omitempty"`

	// Provides information about the corporate action option.
	RateAndAmountDetails *RateDetails3 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *PriceDetails6 `xml:"PricDtls,omitempty"`
}

func (c *CashOption17) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption17) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator1Choice {
	c.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator1Choice)
	return c.NonEligibleProceedsIndicator
}

func (c *CashOption17) SetIssuerOfferorTaxabilityIndicator(value string) {
	c.IssuerOfferorTaxabilityIndicator = (*IssuerTaxability1Code)(&value)
}

func (c *CashOption17) AddIncomeType() *GenericIdentification20 {
	c.IncomeType = new(GenericIdentification20)
	return c.IncomeType
}

func (c *CashOption17) AddCashAccountIdentification() *CashAccountIdentification5Choice {
	c.CashAccountIdentification = new(CashAccountIdentification5Choice)
	return c.CashAccountIdentification
}

func (c *CashOption17) AddAmountDetails() *CorporateActionAmounts15 {
	c.AmountDetails = new(CorporateActionAmounts15)
	return c.AmountDetails
}

func (c *CashOption17) AddDateDetails() *CorporateActionDate23 {
	c.DateDetails = new(CorporateActionDate23)
	return c.DateDetails
}

func (c *CashOption17) AddForeignExchangeDetails() *ForeignExchangeTerms13 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms13)
	return c.ForeignExchangeDetails
}

func (c *CashOption17) AddRateAndAmountDetails() *RateDetails3 {
	c.RateAndAmountDetails = new(RateDetails3)
	return c.RateAndAmountDetails
}

func (c *CashOption17) AddPriceDetails() *PriceDetails6 {
	c.PriceDetails = new(PriceDetails6)
	return c.PriceDetails
}

// Provides information about the cash option.
type CashOption18 struct {

	// Indicates whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates whether the cash payment occurs or will occur in advance of receipt of proceeds from the issuer and based on a contractual agreement established with the account servicer or upon receipt of proceeds from the issuer.
	ContractualPaymentIndicator *Payment1Code `xml:"CtrctlPmtInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification20 `xml:"IncmTp,omitempty"`

	// Choice between a cash account, a charges account or a tax account.
	Account *Account8Choice `xml:"Acct,omitempty"`

	// Provides information about cash parties.
	CashParties *CashParties10 `xml:"CshPties,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts17 `xml:"AmtDtls"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate24 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount.
	ForeignExchangeDetails *ForeignExchangeTerms11 `xml:"FXDtls,omitempty"`

	// Provides information about the tax voucher related to a cash movement.
	TaxVoucherDetails *TaxVoucher2 `xml:"TaxVchrDtls,omitempty"`

	// Provides information about the corporate action option.
	RateAndAmountDetails *RateDetails7 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *PriceDetails7 `xml:"PricDtls,omitempty"`
}

func (c *CashOption18) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption18) SetContractualPaymentIndicator(value string) {
	c.ContractualPaymentIndicator = (*Payment1Code)(&value)
}

func (c *CashOption18) AddIncomeType() *GenericIdentification20 {
	c.IncomeType = new(GenericIdentification20)
	return c.IncomeType
}

func (c *CashOption18) AddAccount() *Account8Choice {
	c.Account = new(Account8Choice)
	return c.Account
}

func (c *CashOption18) AddCashParties() *CashParties10 {
	c.CashParties = new(CashParties10)
	return c.CashParties
}

func (c *CashOption18) AddAmountDetails() *CorporateActionAmounts17 {
	c.AmountDetails = new(CorporateActionAmounts17)
	return c.AmountDetails
}

func (c *CashOption18) AddDateDetails() *CorporateActionDate24 {
	c.DateDetails = new(CorporateActionDate24)
	return c.DateDetails
}

func (c *CashOption18) AddForeignExchangeDetails() *ForeignExchangeTerms11 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms11)
	return c.ForeignExchangeDetails
}

func (c *CashOption18) AddTaxVoucherDetails() *TaxVoucher2 {
	c.TaxVoucherDetails = new(TaxVoucher2)
	return c.TaxVoucherDetails
}

func (c *CashOption18) AddRateAndAmountDetails() *RateDetails7 {
	c.RateAndAmountDetails = new(RateDetails7)
	return c.RateAndAmountDetails
}

func (c *CashOption18) AddPriceDetails() *PriceDetails7 {
	c.PriceDetails = new(PriceDetails7)
	return c.PriceDetails
}

// Provides information about the cash option.
type CashOption19 struct {

	// Indicates whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Choice between a cash account, a charges account or a tax account.
	Account *Account8Choice `xml:"Acct,omitempty"`

	// Date/Time of the posting (credit or debit) to the account.
	PostingDate *DateAndDateTimeChoice `xml:"PstngDt"`

	// Date/Time of the posting (credit or debit) to the account that was initially communicated in the confirmation.
	OriginalPostingDate *DateAndDateTimeChoice `xml:"OrgnlPstngDt,omitempty"`

	// Date/time when calculating economic benefit for a cash amount.
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`

	// Amount of money that is to be/was posted to the account.
	PostingAmount *ActiveCurrencyAndAmount `xml:"PstngAmt"`
}

func (c *CashOption19) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption19) AddAccount() *Account8Choice {
	c.Account = new(Account8Choice)
	return c.Account
}

func (c *CashOption19) AddPostingDate() *DateAndDateTimeChoice {
	c.PostingDate = new(DateAndDateTimeChoice)
	return c.PostingDate
}

func (c *CashOption19) AddOriginalPostingDate() *DateAndDateTimeChoice {
	c.OriginalPostingDate = new(DateAndDateTimeChoice)
	return c.OriginalPostingDate
}

func (c *CashOption19) AddValueDate() *DateAndDateTimeChoice {
	c.ValueDate = new(DateAndDateTimeChoice)
	return c.ValueDate
}

func (c *CashOption19) SetPostingAmount(value, currency string) {
	c.PostingAmount = NewActiveCurrencyAndAmount(value, currency)
}

// Provides information about the cash option.
type CashOption2 struct {

	// Indicates whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates whether the cash payment occurs or will occur in advance of receipt of proceeds from the issuer and based on a contractual agreement established with the account servicer or upon receipt of proceeds from the issuer.
	ContractualPaymentIndicator *Payment1Code `xml:"CtrctlPmtInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification20 `xml:"IncmTp,omitempty"`

	// Choice between a cash account, a charges account or a tax account.
	Account *Account8Choice `xml:"Acct,omitempty"`

	// Provides information about cash parties.
	CashParties *CashParties2 `xml:"CshPties,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts2 `xml:"AmtDtls"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate7 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount.
	ForeignExchangeDetails *ForeignExchangeTerms11 `xml:"FXDtls,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat7Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`

	// Provides information about the tax voucher related to a cash movement.
	TaxVoucherDetails *TaxVoucher2 `xml:"TaxVchrDtls,omitempty"`
}

func (c *CashOption2) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption2) SetContractualPaymentIndicator(value string) {
	c.ContractualPaymentIndicator = (*Payment1Code)(&value)
}

func (c *CashOption2) AddIncomeType() *GenericIdentification20 {
	c.IncomeType = new(GenericIdentification20)
	return c.IncomeType
}

func (c *CashOption2) AddAccount() *Account8Choice {
	c.Account = new(Account8Choice)
	return c.Account
}

func (c *CashOption2) AddCashParties() *CashParties2 {
	c.CashParties = new(CashParties2)
	return c.CashParties
}

func (c *CashOption2) AddAmountDetails() *CorporateActionAmounts2 {
	c.AmountDetails = new(CorporateActionAmounts2)
	return c.AmountDetails
}

func (c *CashOption2) AddDateDetails() *CorporateActionDate7 {
	c.DateDetails = new(CorporateActionDate7)
	return c.DateDetails
}

func (c *CashOption2) AddForeignExchangeDetails() *ForeignExchangeTerms11 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms11)
	return c.ForeignExchangeDetails
}

func (c *CashOption2) AddGenericCashPriceReceivedPerProduct() *PriceFormat7Choice {
	c.GenericCashPriceReceivedPerProduct = new(PriceFormat7Choice)
	return c.GenericCashPriceReceivedPerProduct
}

func (c *CashOption2) AddTaxVoucherDetails() *TaxVoucher2 {
	c.TaxVoucherDetails = new(TaxVoucher2)
	return c.TaxVoucherDetails
}

// Provides information about the cash option.
type CashOption24 struct {

	// Indicates whether the value is a debit or a credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator1Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *IssuerTaxability1Code `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification20 `xml:"IncmTp,omitempty"`

	// Identification of the account in which cash is maintained.
	CashAccountIdentification *CashAccountIdentification5Choice `xml:"CshAcctId,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts21 `xml:"AmtDtls,omitempty"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate23 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount
	ForeignExchangeDetails *ForeignExchangeTerms13 `xml:"FXDtls,omitempty"`

	// Provides information about the corporate action option.
	RateAndAmountDetails *RateDetails10 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *PriceDetails10 `xml:"PricDtls,omitempty"`
}

func (c *CashOption24) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption24) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator1Choice {
	c.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator1Choice)
	return c.NonEligibleProceedsIndicator
}

func (c *CashOption24) SetIssuerOfferorTaxabilityIndicator(value string) {
	c.IssuerOfferorTaxabilityIndicator = (*IssuerTaxability1Code)(&value)
}

func (c *CashOption24) AddIncomeType() *GenericIdentification20 {
	c.IncomeType = new(GenericIdentification20)
	return c.IncomeType
}

func (c *CashOption24) AddCashAccountIdentification() *CashAccountIdentification5Choice {
	c.CashAccountIdentification = new(CashAccountIdentification5Choice)
	return c.CashAccountIdentification
}

func (c *CashOption24) AddAmountDetails() *CorporateActionAmounts21 {
	c.AmountDetails = new(CorporateActionAmounts21)
	return c.AmountDetails
}

func (c *CashOption24) AddDateDetails() *CorporateActionDate23 {
	c.DateDetails = new(CorporateActionDate23)
	return c.DateDetails
}

func (c *CashOption24) AddForeignExchangeDetails() *ForeignExchangeTerms13 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms13)
	return c.ForeignExchangeDetails
}

func (c *CashOption24) AddRateAndAmountDetails() *RateDetails10 {
	c.RateAndAmountDetails = new(RateDetails10)
	return c.RateAndAmountDetails
}

func (c *CashOption24) AddPriceDetails() *PriceDetails10 {
	c.PriceDetails = new(PriceDetails10)
	return c.PriceDetails
}

// Provides information about the cash option.
type CashOption25 struct {

	// Indicates whether the value is a debit or a credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates whether the cash payment occurs or will occur in advance of receipt of proceeds from the issuer and based on a contractual agreement established with the account servicer or upon receipt of proceeds from the issuer.
	ContractualPaymentIndicator *Payment1Code `xml:"CtrctlPmtInd,omitempty"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator1Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *IssuerTaxability1Code `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification20 `xml:"IncmTp,omitempty"`

	// Identification of the account in which cash is maintained.
	CashAccountIdentification *CashAccountIdentification5Choice `xml:"CshAcctId,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts22 `xml:"AmtDtls,omitempty"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate23 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount.
	ForeignExchangeDetails *ForeignExchangeTerms13 `xml:"FXDtls,omitempty"`

	// Provides information about the corporate action option.
	RateAndAmountDetails *RateDetails10 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *PriceDetails10 `xml:"PricDtls,omitempty"`
}

func (c *CashOption25) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption25) SetContractualPaymentIndicator(value string) {
	c.ContractualPaymentIndicator = (*Payment1Code)(&value)
}

func (c *CashOption25) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator1Choice {
	c.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator1Choice)
	return c.NonEligibleProceedsIndicator
}

func (c *CashOption25) SetIssuerOfferorTaxabilityIndicator(value string) {
	c.IssuerOfferorTaxabilityIndicator = (*IssuerTaxability1Code)(&value)
}

func (c *CashOption25) AddIncomeType() *GenericIdentification20 {
	c.IncomeType = new(GenericIdentification20)
	return c.IncomeType
}

func (c *CashOption25) AddCashAccountIdentification() *CashAccountIdentification5Choice {
	c.CashAccountIdentification = new(CashAccountIdentification5Choice)
	return c.CashAccountIdentification
}

func (c *CashOption25) AddAmountDetails() *CorporateActionAmounts22 {
	c.AmountDetails = new(CorporateActionAmounts22)
	return c.AmountDetails
}

func (c *CashOption25) AddDateDetails() *CorporateActionDate23 {
	c.DateDetails = new(CorporateActionDate23)
	return c.DateDetails
}

func (c *CashOption25) AddForeignExchangeDetails() *ForeignExchangeTerms13 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms13)
	return c.ForeignExchangeDetails
}

func (c *CashOption25) AddRateAndAmountDetails() *RateDetails10 {
	c.RateAndAmountDetails = new(RateDetails10)
	return c.RateAndAmountDetails
}

func (c *CashOption25) AddPriceDetails() *PriceDetails10 {
	c.PriceDetails = new(PriceDetails10)
	return c.PriceDetails
}

// Provides information about the cash option.
type CashOption26 struct {

	// Indicates whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates whether the cash payment occurs or will occur in advance of receipt of proceeds from the issuer and based on a contractual agreement established with the account servicer or upon receipt of proceeds from the issuer.
	ContractualPaymentIndicator *Payment1Code `xml:"CtrctlPmtInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification20 `xml:"IncmTp,omitempty"`

	// Choice between a cash account, a charges account or a tax account.
	Account *Account8Choice `xml:"Acct,omitempty"`

	// Provides information about cash parties.
	CashParties *CashParties10 `xml:"CshPties,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts23 `xml:"AmtDtls"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate24 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount.
	ForeignExchangeDetails *ForeignExchangeTerms11 `xml:"FXDtls,omitempty"`

	// Provides information about the tax voucher related to a cash movement.
	TaxVoucherDetails *TaxVoucher2 `xml:"TaxVchrDtls,omitempty"`

	// Provides information about the corporate action option.
	RateAndAmountDetails *RateDetails11 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *PriceDetails11 `xml:"PricDtls,omitempty"`
}

func (c *CashOption26) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption26) SetContractualPaymentIndicator(value string) {
	c.ContractualPaymentIndicator = (*Payment1Code)(&value)
}

func (c *CashOption26) AddIncomeType() *GenericIdentification20 {
	c.IncomeType = new(GenericIdentification20)
	return c.IncomeType
}

func (c *CashOption26) AddAccount() *Account8Choice {
	c.Account = new(Account8Choice)
	return c.Account
}

func (c *CashOption26) AddCashParties() *CashParties10 {
	c.CashParties = new(CashParties10)
	return c.CashParties
}

func (c *CashOption26) AddAmountDetails() *CorporateActionAmounts23 {
	c.AmountDetails = new(CorporateActionAmounts23)
	return c.AmountDetails
}

func (c *CashOption26) AddDateDetails() *CorporateActionDate24 {
	c.DateDetails = new(CorporateActionDate24)
	return c.DateDetails
}

func (c *CashOption26) AddForeignExchangeDetails() *ForeignExchangeTerms11 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms11)
	return c.ForeignExchangeDetails
}

func (c *CashOption26) AddTaxVoucherDetails() *TaxVoucher2 {
	c.TaxVoucherDetails = new(TaxVoucher2)
	return c.TaxVoucherDetails
}

func (c *CashOption26) AddRateAndAmountDetails() *RateDetails11 {
	c.RateAndAmountDetails = new(RateDetails11)
	return c.RateAndAmountDetails
}

func (c *CashOption26) AddPriceDetails() *PriceDetails11 {
	c.PriceDetails = new(PriceDetails11)
	return c.PriceDetails
}

// Provides information about the cash option.
type CashOption3 struct {

	// Indicates whether the value is a debit or a credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates whether the cash payment occurs or will occur in advance of receipt of proceeds from the issuer and based on a contractual agreement established with the account servicer or upon receipt of proceeds from the issuer.
	ContractualPaymentIndicator *Payment1Code `xml:"CtrctlPmtInd,omitempty"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator1Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification20 `xml:"IncmTp,omitempty"`

	// Identification of the account in which cash is maintained.
	CashAccountIdentification *CashAccountIdentification5Choice `xml:"CshAcctId,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts3 `xml:"AmtDtls,omitempty"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate9 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount.
	ForeignExchangeDetails *ForeignExchangeTerms11 `xml:"FXDtls,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat10Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (c *CashOption3) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption3) SetContractualPaymentIndicator(value string) {
	c.ContractualPaymentIndicator = (*Payment1Code)(&value)
}

func (c *CashOption3) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator1Choice {
	c.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator1Choice)
	return c.NonEligibleProceedsIndicator
}

func (c *CashOption3) AddIncomeType() *GenericIdentification20 {
	c.IncomeType = new(GenericIdentification20)
	return c.IncomeType
}

func (c *CashOption3) AddCashAccountIdentification() *CashAccountIdentification5Choice {
	c.CashAccountIdentification = new(CashAccountIdentification5Choice)
	return c.CashAccountIdentification
}

func (c *CashOption3) AddAmountDetails() *CorporateActionAmounts3 {
	c.AmountDetails = new(CorporateActionAmounts3)
	return c.AmountDetails
}

func (c *CashOption3) AddDateDetails() *CorporateActionDate9 {
	c.DateDetails = new(CorporateActionDate9)
	return c.DateDetails
}

func (c *CashOption3) AddForeignExchangeDetails() *ForeignExchangeTerms11 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms11)
	return c.ForeignExchangeDetails
}

func (c *CashOption3) AddGenericCashPriceReceivedPerProduct() *PriceFormat10Choice {
	c.GenericCashPriceReceivedPerProduct = new(PriceFormat10Choice)
	return c.GenericCashPriceReceivedPerProduct
}

// Provides information about the cash option.
type CashOption30 struct {

	// Indicates whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates whether the cash payment occurs or will occur in advance of receipt of proceeds from the issuer and based on a contractual agreement established with the account servicer or upon receipt of proceeds from the issuer.
	ContractualPaymentIndicator *Payment1Code `xml:"CtrctlPmtInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification20 `xml:"IncmTp,omitempty"`

	// Specifies the basis for the reduced rate of withholding.
	ExemptionType []*GenericIdentification20 `xml:"XmptnTp,omitempty"`

	// Choice between a cash account, a charges account or a tax account.
	Account *Account8Choice `xml:"Acct,omitempty"`

	// Provides information about cash parties.
	CashParties *CashParties10 `xml:"CshPties,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts29 `xml:"AmtDtls"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate24 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount.
	ForeignExchangeDetails *ForeignExchangeTerms11 `xml:"FXDtls,omitempty"`

	// Provides information about the tax voucher related to a cash movement.
	TaxVoucherDetails *TaxVoucher2 `xml:"TaxVchrDtls,omitempty"`

	// Provides information about the corporate action option.
	RateAndAmountDetails *RateDetails15 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *PriceDetails14 `xml:"PricDtls,omitempty"`
}

func (c *CashOption30) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption30) SetContractualPaymentIndicator(value string) {
	c.ContractualPaymentIndicator = (*Payment1Code)(&value)
}

func (c *CashOption30) AddIncomeType() *GenericIdentification20 {
	c.IncomeType = new(GenericIdentification20)
	return c.IncomeType
}

func (c *CashOption30) AddExemptionType() *GenericIdentification20 {
	newValue := new(GenericIdentification20)
	c.ExemptionType = append(c.ExemptionType, newValue)
	return newValue
}

func (c *CashOption30) AddAccount() *Account8Choice {
	c.Account = new(Account8Choice)
	return c.Account
}

func (c *CashOption30) AddCashParties() *CashParties10 {
	c.CashParties = new(CashParties10)
	return c.CashParties
}

func (c *CashOption30) AddAmountDetails() *CorporateActionAmounts29 {
	c.AmountDetails = new(CorporateActionAmounts29)
	return c.AmountDetails
}

func (c *CashOption30) AddDateDetails() *CorporateActionDate24 {
	c.DateDetails = new(CorporateActionDate24)
	return c.DateDetails
}

func (c *CashOption30) AddForeignExchangeDetails() *ForeignExchangeTerms11 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms11)
	return c.ForeignExchangeDetails
}

func (c *CashOption30) AddTaxVoucherDetails() *TaxVoucher2 {
	c.TaxVoucherDetails = new(TaxVoucher2)
	return c.TaxVoucherDetails
}

func (c *CashOption30) AddRateAndAmountDetails() *RateDetails15 {
	c.RateAndAmountDetails = new(RateDetails15)
	return c.RateAndAmountDetails
}

func (c *CashOption30) AddPriceDetails() *PriceDetails14 {
	c.PriceDetails = new(PriceDetails14)
	return c.PriceDetails
}

// Provides information about the cash option.
type CashOption31 struct {

	// Indicates whether the value is a debit or a credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates whether the cash payment occurs or will occur in advance of receipt of proceeds from the issuer and based on a contractual agreement established with the account servicer or upon receipt of proceeds from the issuer.
	ContractualPaymentIndicator *Payment1Code `xml:"CtrctlPmtInd,omitempty"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator1Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *IssuerTaxability1Code `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification20 `xml:"IncmTp,omitempty"`

	// Specifies the basis for the reduced rate of withholding.
	ExemptionType []*GenericIdentification20 `xml:"XmptnTp,omitempty"`

	// Identification of the account in which cash is maintained.
	CashAccountIdentification *CashAccountIdentification5Choice `xml:"CshAcctId,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts28 `xml:"AmtDtls,omitempty"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate23 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount.
	ForeignExchangeDetails *ForeignExchangeTerms13 `xml:"FXDtls,omitempty"`

	// Provides information about the corporate action option.
	RateAndAmountDetails *RateDetails14 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *PriceDetails15 `xml:"PricDtls,omitempty"`
}

func (c *CashOption31) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption31) SetContractualPaymentIndicator(value string) {
	c.ContractualPaymentIndicator = (*Payment1Code)(&value)
}

func (c *CashOption31) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator1Choice {
	c.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator1Choice)
	return c.NonEligibleProceedsIndicator
}

func (c *CashOption31) SetIssuerOfferorTaxabilityIndicator(value string) {
	c.IssuerOfferorTaxabilityIndicator = (*IssuerTaxability1Code)(&value)
}

func (c *CashOption31) AddIncomeType() *GenericIdentification20 {
	c.IncomeType = new(GenericIdentification20)
	return c.IncomeType
}

func (c *CashOption31) AddExemptionType() *GenericIdentification20 {
	newValue := new(GenericIdentification20)
	c.ExemptionType = append(c.ExemptionType, newValue)
	return newValue
}

func (c *CashOption31) AddCashAccountIdentification() *CashAccountIdentification5Choice {
	c.CashAccountIdentification = new(CashAccountIdentification5Choice)
	return c.CashAccountIdentification
}

func (c *CashOption31) AddAmountDetails() *CorporateActionAmounts28 {
	c.AmountDetails = new(CorporateActionAmounts28)
	return c.AmountDetails
}

func (c *CashOption31) AddDateDetails() *CorporateActionDate23 {
	c.DateDetails = new(CorporateActionDate23)
	return c.DateDetails
}

func (c *CashOption31) AddForeignExchangeDetails() *ForeignExchangeTerms13 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms13)
	return c.ForeignExchangeDetails
}

func (c *CashOption31) AddRateAndAmountDetails() *RateDetails14 {
	c.RateAndAmountDetails = new(RateDetails14)
	return c.RateAndAmountDetails
}

func (c *CashOption31) AddPriceDetails() *PriceDetails15 {
	c.PriceDetails = new(PriceDetails15)
	return c.PriceDetails
}

// Provides information about the cash option.
type CashOption32 struct {

	// Indicates whether the value is a debit or a credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator1Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *IssuerTaxability1Code `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification20 `xml:"IncmTp,omitempty"`

	// Specifies the basis for the reduced rate of withholding.
	ExemptionType []*GenericIdentification20 `xml:"XmptnTp,omitempty"`

	// Identification of the account in which cash is maintained.
	CashAccountIdentification *CashAccountIdentification5Choice `xml:"CshAcctId,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts27 `xml:"AmtDtls,omitempty"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate23 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount
	ForeignExchangeDetails *ForeignExchangeTerms13 `xml:"FXDtls,omitempty"`

	// Provides information about the corporate action option.
	RateAndAmountDetails *RateDetails14 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *PriceDetails15 `xml:"PricDtls,omitempty"`
}

func (c *CashOption32) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption32) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator1Choice {
	c.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator1Choice)
	return c.NonEligibleProceedsIndicator
}

func (c *CashOption32) SetIssuerOfferorTaxabilityIndicator(value string) {
	c.IssuerOfferorTaxabilityIndicator = (*IssuerTaxability1Code)(&value)
}

func (c *CashOption32) AddIncomeType() *GenericIdentification20 {
	c.IncomeType = new(GenericIdentification20)
	return c.IncomeType
}

func (c *CashOption32) AddExemptionType() *GenericIdentification20 {
	newValue := new(GenericIdentification20)
	c.ExemptionType = append(c.ExemptionType, newValue)
	return newValue
}

func (c *CashOption32) AddCashAccountIdentification() *CashAccountIdentification5Choice {
	c.CashAccountIdentification = new(CashAccountIdentification5Choice)
	return c.CashAccountIdentification
}

func (c *CashOption32) AddAmountDetails() *CorporateActionAmounts27 {
	c.AmountDetails = new(CorporateActionAmounts27)
	return c.AmountDetails
}

func (c *CashOption32) AddDateDetails() *CorporateActionDate23 {
	c.DateDetails = new(CorporateActionDate23)
	return c.DateDetails
}

func (c *CashOption32) AddForeignExchangeDetails() *ForeignExchangeTerms13 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms13)
	return c.ForeignExchangeDetails
}

func (c *CashOption32) AddRateAndAmountDetails() *RateDetails14 {
	c.RateAndAmountDetails = new(RateDetails14)
	return c.RateAndAmountDetails
}

func (c *CashOption32) AddPriceDetails() *PriceDetails15 {
	c.PriceDetails = new(PriceDetails15)
	return c.PriceDetails
}

// Provides information about the cash option.
type CashOption39 struct {

	// Indicates whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates whether the cash payment occurs or will occur in advance of receipt of proceeds from the issuer and based on a contractual agreement established with the account servicer or upon receipt of proceeds from the issuer.
	ContractualPaymentIndicator *Payment1Code `xml:"CtrctlPmtInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification20 `xml:"IncmTp,omitempty"`

	// Specifies the basis for the reduced rate of withholding.
	ExemptionType []*GenericIdentification20 `xml:"XmptnTp,omitempty"`

	// Choice between a cash account, a charges account or a tax account.
	Account *Account8Choice `xml:"Acct,omitempty"`

	// Provides information about cash parties.
	CashParties *CashParties21 `xml:"CshPties,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts29 `xml:"AmtDtls"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate24 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount.
	ForeignExchangeDetails *ForeignExchangeTerms11 `xml:"FXDtls,omitempty"`

	// Provides information about the tax voucher related to a cash movement.
	TaxVoucherDetails *TaxVoucher2 `xml:"TaxVchrDtls,omitempty"`

	// Provides information about the corporate action option.
	RateAndAmountDetails *RateDetails15 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *PriceDetails14 `xml:"PricDtls,omitempty"`
}

func (c *CashOption39) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption39) SetContractualPaymentIndicator(value string) {
	c.ContractualPaymentIndicator = (*Payment1Code)(&value)
}

func (c *CashOption39) AddIncomeType() *GenericIdentification20 {
	c.IncomeType = new(GenericIdentification20)
	return c.IncomeType
}

func (c *CashOption39) AddExemptionType() *GenericIdentification20 {
	newValue := new(GenericIdentification20)
	c.ExemptionType = append(c.ExemptionType, newValue)
	return newValue
}

func (c *CashOption39) AddAccount() *Account8Choice {
	c.Account = new(Account8Choice)
	return c.Account
}

func (c *CashOption39) AddCashParties() *CashParties21 {
	c.CashParties = new(CashParties21)
	return c.CashParties
}

func (c *CashOption39) AddAmountDetails() *CorporateActionAmounts29 {
	c.AmountDetails = new(CorporateActionAmounts29)
	return c.AmountDetails
}

func (c *CashOption39) AddDateDetails() *CorporateActionDate24 {
	c.DateDetails = new(CorporateActionDate24)
	return c.DateDetails
}

func (c *CashOption39) AddForeignExchangeDetails() *ForeignExchangeTerms11 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms11)
	return c.ForeignExchangeDetails
}

func (c *CashOption39) AddTaxVoucherDetails() *TaxVoucher2 {
	c.TaxVoucherDetails = new(TaxVoucher2)
	return c.TaxVoucherDetails
}

func (c *CashOption39) AddRateAndAmountDetails() *RateDetails15 {
	c.RateAndAmountDetails = new(RateDetails15)
	return c.RateAndAmountDetails
}

func (c *CashOption39) AddPriceDetails() *PriceDetails14 {
	c.PriceDetails = new(PriceDetails14)
	return c.PriceDetails
}

// Provides information about the cash option.
type CashOption4 struct {

	// Indicates whether the value is a debit or a credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator1Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification20 `xml:"IncmTp,omitempty"`

	// Identification of the account in which cash is maintained.
	CashAccountIdentification *CashAccountIdentification5Choice `xml:"CshAcctId,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts4 `xml:"AmtDtls,omitempty"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate9 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount
	ForeignExchangeDetails *ForeignExchangeTerms11 `xml:"FXDtls,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat10Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (c *CashOption4) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption4) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator1Choice {
	c.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator1Choice)
	return c.NonEligibleProceedsIndicator
}

func (c *CashOption4) AddIncomeType() *GenericIdentification20 {
	c.IncomeType = new(GenericIdentification20)
	return c.IncomeType
}

func (c *CashOption4) AddCashAccountIdentification() *CashAccountIdentification5Choice {
	c.CashAccountIdentification = new(CashAccountIdentification5Choice)
	return c.CashAccountIdentification
}

func (c *CashOption4) AddAmountDetails() *CorporateActionAmounts4 {
	c.AmountDetails = new(CorporateActionAmounts4)
	return c.AmountDetails
}

func (c *CashOption4) AddDateDetails() *CorporateActionDate9 {
	c.DateDetails = new(CorporateActionDate9)
	return c.DateDetails
}

func (c *CashOption4) AddForeignExchangeDetails() *ForeignExchangeTerms11 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms11)
	return c.ForeignExchangeDetails
}

func (c *CashOption4) AddGenericCashPriceReceivedPerProduct() *PriceFormat10Choice {
	c.GenericCashPriceReceivedPerProduct = new(PriceFormat10Choice)
	return c.GenericCashPriceReceivedPerProduct
}

// Provides information about the cash option.
type CashOption42 struct {

	// Indicates whether the value is a debit or a credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates whether the cash payment occurs or will occur in advance of receipt of proceeds from the issuer and based on a contractual agreement established with the account servicer or upon receipt of proceeds from the issuer.
	ContractualPaymentIndicator *Payment1Code `xml:"CtrctlPmtInd,omitempty"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator3Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *GenericIdentification30 `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification30 `xml:"IncmTp,omitempty"`

	// Specifies the basis for the reduced rate of withholding.
	ExemptionType []*GenericIdentification30 `xml:"XmptnTp,omitempty"`

	// Indicates the country from which the income originates.
	CountryOfIncomeSource *CountryCode `xml:"CtryOfIncmSrc,omitempty"`

	// Identification of the account in which cash is maintained.
	CashAccountIdentification *CashAccountIdentification5Choice `xml:"CshAcctId,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts38 `xml:"AmtDtls,omitempty"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate47 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount.
	ForeignExchangeDetails *ForeignExchangeTerms24 `xml:"FXDtls,omitempty"`

	// Provides information about the corporate action option.
	RateAndAmountDetails *RateDetails22 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *PriceDetails22 `xml:"PricDtls,omitempty"`
}

func (c *CashOption42) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption42) SetContractualPaymentIndicator(value string) {
	c.ContractualPaymentIndicator = (*Payment1Code)(&value)
}

func (c *CashOption42) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator3Choice {
	c.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator3Choice)
	return c.NonEligibleProceedsIndicator
}

func (c *CashOption42) AddIssuerOfferorTaxabilityIndicator() *GenericIdentification30 {
	c.IssuerOfferorTaxabilityIndicator = new(GenericIdentification30)
	return c.IssuerOfferorTaxabilityIndicator
}

func (c *CashOption42) AddIncomeType() *GenericIdentification30 {
	c.IncomeType = new(GenericIdentification30)
	return c.IncomeType
}

func (c *CashOption42) AddExemptionType() *GenericIdentification30 {
	newValue := new(GenericIdentification30)
	c.ExemptionType = append(c.ExemptionType, newValue)
	return newValue
}

func (c *CashOption42) SetCountryOfIncomeSource(value string) {
	c.CountryOfIncomeSource = (*CountryCode)(&value)
}

func (c *CashOption42) AddCashAccountIdentification() *CashAccountIdentification5Choice {
	c.CashAccountIdentification = new(CashAccountIdentification5Choice)
	return c.CashAccountIdentification
}

func (c *CashOption42) AddAmountDetails() *CorporateActionAmounts38 {
	c.AmountDetails = new(CorporateActionAmounts38)
	return c.AmountDetails
}

func (c *CashOption42) AddDateDetails() *CorporateActionDate47 {
	c.DateDetails = new(CorporateActionDate47)
	return c.DateDetails
}

func (c *CashOption42) AddForeignExchangeDetails() *ForeignExchangeTerms24 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms24)
	return c.ForeignExchangeDetails
}

func (c *CashOption42) AddRateAndAmountDetails() *RateDetails22 {
	c.RateAndAmountDetails = new(RateDetails22)
	return c.RateAndAmountDetails
}

func (c *CashOption42) AddPriceDetails() *PriceDetails22 {
	c.PriceDetails = new(PriceDetails22)
	return c.PriceDetails
}

// Provides information about the cash option.
type CashOption43 struct {

	// Indicates whether the value is a debit or a credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator3Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *GenericIdentification30 `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification30 `xml:"IncmTp,omitempty"`

	// Specifies the basis for the reduced rate of withholding.
	ExemptionType []*GenericIdentification30 `xml:"XmptnTp,omitempty"`

	// Indicates the country from which the income originates.
	CountryOfIncomeSource *CountryCode `xml:"CtryOfIncmSrc,omitempty"`

	// Identification of the account in which cash is maintained.
	CashAccountIdentification *CashAccountIdentification5Choice `xml:"CshAcctId,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts36 `xml:"AmtDtls,omitempty"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate47 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount
	ForeignExchangeDetails *ForeignExchangeTerms24 `xml:"FXDtls,omitempty"`

	// Provides information about the corporate action option.
	RateAndAmountDetails *RateDetails22 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *PriceDetails22 `xml:"PricDtls,omitempty"`
}

func (c *CashOption43) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption43) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator3Choice {
	c.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator3Choice)
	return c.NonEligibleProceedsIndicator
}

func (c *CashOption43) AddIssuerOfferorTaxabilityIndicator() *GenericIdentification30 {
	c.IssuerOfferorTaxabilityIndicator = new(GenericIdentification30)
	return c.IssuerOfferorTaxabilityIndicator
}

func (c *CashOption43) AddIncomeType() *GenericIdentification30 {
	c.IncomeType = new(GenericIdentification30)
	return c.IncomeType
}

func (c *CashOption43) AddExemptionType() *GenericIdentification30 {
	newValue := new(GenericIdentification30)
	c.ExemptionType = append(c.ExemptionType, newValue)
	return newValue
}

func (c *CashOption43) SetCountryOfIncomeSource(value string) {
	c.CountryOfIncomeSource = (*CountryCode)(&value)
}

func (c *CashOption43) AddCashAccountIdentification() *CashAccountIdentification5Choice {
	c.CashAccountIdentification = new(CashAccountIdentification5Choice)
	return c.CashAccountIdentification
}

func (c *CashOption43) AddAmountDetails() *CorporateActionAmounts36 {
	c.AmountDetails = new(CorporateActionAmounts36)
	return c.AmountDetails
}

func (c *CashOption43) AddDateDetails() *CorporateActionDate47 {
	c.DateDetails = new(CorporateActionDate47)
	return c.DateDetails
}

func (c *CashOption43) AddForeignExchangeDetails() *ForeignExchangeTerms24 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms24)
	return c.ForeignExchangeDetails
}

func (c *CashOption43) AddRateAndAmountDetails() *RateDetails22 {
	c.RateAndAmountDetails = new(RateDetails22)
	return c.RateAndAmountDetails
}

func (c *CashOption43) AddPriceDetails() *PriceDetails22 {
	c.PriceDetails = new(PriceDetails22)
	return c.PriceDetails
}

// Provides information about the cash option.
type CashOption44 struct {

	// Indicates whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates whether the cash payment occurs or will occur in advance of receipt of proceeds from the issuer and based on a contractual agreement established with the account servicer or upon receipt of proceeds from the issuer.
	ContractualPaymentIndicator *Payment1Code `xml:"CtrctlPmtInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *GenericIdentification30 `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification30 `xml:"IncmTp,omitempty"`

	// Specifies the basis for the reduced rate of withholding.
	ExemptionType []*GenericIdentification30 `xml:"XmptnTp,omitempty"`

	// Indicates the country from which the income originates.
	CountryOfIncomeSource *CountryCode `xml:"CtryOfIncmSrc,omitempty"`

	// Choice between a cash account, a charges account or a tax account.
	Account *Account8Choice `xml:"Acct,omitempty"`

	// Provides information about cash parties.
	CashParties *CashParties28 `xml:"CshPties,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts37 `xml:"AmtDtls"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate24 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount.
	ForeignExchangeDetails *ForeignExchangeTerms23 `xml:"FXDtls,omitempty"`

	// Provides information about the tax voucher related to a cash movement.
	TaxVoucherDetails *TaxVoucher2 `xml:"TaxVchrDtls,omitempty"`

	// Provides information about the corporate action option.
	RateAndAmountDetails *RateDetails23 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *PriceDetails23 `xml:"PricDtls,omitempty"`
}

func (c *CashOption44) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption44) SetContractualPaymentIndicator(value string) {
	c.ContractualPaymentIndicator = (*Payment1Code)(&value)
}

func (c *CashOption44) AddIssuerOfferorTaxabilityIndicator() *GenericIdentification30 {
	c.IssuerOfferorTaxabilityIndicator = new(GenericIdentification30)
	return c.IssuerOfferorTaxabilityIndicator
}

func (c *CashOption44) AddIncomeType() *GenericIdentification30 {
	c.IncomeType = new(GenericIdentification30)
	return c.IncomeType
}

func (c *CashOption44) AddExemptionType() *GenericIdentification30 {
	newValue := new(GenericIdentification30)
	c.ExemptionType = append(c.ExemptionType, newValue)
	return newValue
}

func (c *CashOption44) SetCountryOfIncomeSource(value string) {
	c.CountryOfIncomeSource = (*CountryCode)(&value)
}

func (c *CashOption44) AddAccount() *Account8Choice {
	c.Account = new(Account8Choice)
	return c.Account
}

func (c *CashOption44) AddCashParties() *CashParties28 {
	c.CashParties = new(CashParties28)
	return c.CashParties
}

func (c *CashOption44) AddAmountDetails() *CorporateActionAmounts37 {
	c.AmountDetails = new(CorporateActionAmounts37)
	return c.AmountDetails
}

func (c *CashOption44) AddDateDetails() *CorporateActionDate24 {
	c.DateDetails = new(CorporateActionDate24)
	return c.DateDetails
}

func (c *CashOption44) AddForeignExchangeDetails() *ForeignExchangeTerms23 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms23)
	return c.ForeignExchangeDetails
}

func (c *CashOption44) AddTaxVoucherDetails() *TaxVoucher2 {
	c.TaxVoucherDetails = new(TaxVoucher2)
	return c.TaxVoucherDetails
}

func (c *CashOption44) AddRateAndAmountDetails() *RateDetails23 {
	c.RateAndAmountDetails = new(RateDetails23)
	return c.RateAndAmountDetails
}

func (c *CashOption44) AddPriceDetails() *PriceDetails23 {
	c.PriceDetails = new(PriceDetails23)
	return c.PriceDetails
}

// Provides information about the cash option.
type CashOption45 struct {

	// Indicates whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Choice between a cash account, a charges account or a tax account.
	Account *Account8Choice `xml:"Acct,omitempty"`

	// Date/Time of the posting (credit or debit) to the account.
	PostingDate *DateAndDateTimeChoice `xml:"PstngDt"`

	// Date/Time of the posting (credit or debit) to the account that was initially communicated in the confirmation.
	OriginalPostingDate *DateAndDateTimeChoice `xml:"OrgnlPstngDt,omitempty"`

	// Date/time when calculating economic benefit for a cash amount.
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`

	// Amount of money that is to be/was posted to the account.
	PostingAmount *ActiveCurrencyAndAmount `xml:"PstngAmt"`
}

func (c *CashOption45) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption45) AddAccount() *Account8Choice {
	c.Account = new(Account8Choice)
	return c.Account
}

func (c *CashOption45) AddPostingDate() *DateAndDateTimeChoice {
	c.PostingDate = new(DateAndDateTimeChoice)
	return c.PostingDate
}

func (c *CashOption45) AddOriginalPostingDate() *DateAndDateTimeChoice {
	c.OriginalPostingDate = new(DateAndDateTimeChoice)
	return c.OriginalPostingDate
}

func (c *CashOption45) AddValueDate() *DateAndDateTimeChoice {
	c.ValueDate = new(DateAndDateTimeChoice)
	return c.ValueDate
}

func (c *CashOption45) SetPostingAmount(value, currency string) {
	c.PostingAmount = NewActiveCurrencyAndAmount(value, currency)
}

// Provides information about the cash option.
type CashOption46 struct {

	// Indicates whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates whether the cash payment occurs or will occur in advance of receipt of proceeds from the issuer and based on a contractual agreement established with the account servicer or upon receipt of proceeds from the issuer.
	ContractualPaymentIndicator *Payment1Code `xml:"CtrctlPmtInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *GenericIdentification47 `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification47 `xml:"IncmTp,omitempty"`

	// Specifies the basis for the reduced rate of withholding.
	ExemptionType []*GenericIdentification47 `xml:"XmptnTp,omitempty"`

	// Indicates the country from which the income originates.
	CountryOfIncomeSource *CountryCode `xml:"CtryOfIncmSrc,omitempty"`

	// Choice between a cash account, a charges account or a tax account.
	Account *Account9Choice `xml:"Acct,omitempty"`

	// Provides information about cash parties.
	CashParties *CashParties29 `xml:"CshPties,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts39 `xml:"AmtDtls"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate24 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount.
	ForeignExchangeDetails *ForeignExchangeTerms27 `xml:"FXDtls,omitempty"`

	// Provides information about the tax voucher related to a cash movement.
	TaxVoucherDetails *TaxVoucher3 `xml:"TaxVchrDtls,omitempty"`

	// Provides information about the corporate action option.
	RateAndAmountDetails *RateDetails24 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *PriceDetails24 `xml:"PricDtls,omitempty"`
}

func (c *CashOption46) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption46) SetContractualPaymentIndicator(value string) {
	c.ContractualPaymentIndicator = (*Payment1Code)(&value)
}

func (c *CashOption46) AddIssuerOfferorTaxabilityIndicator() *GenericIdentification47 {
	c.IssuerOfferorTaxabilityIndicator = new(GenericIdentification47)
	return c.IssuerOfferorTaxabilityIndicator
}

func (c *CashOption46) AddIncomeType() *GenericIdentification47 {
	c.IncomeType = new(GenericIdentification47)
	return c.IncomeType
}

func (c *CashOption46) AddExemptionType() *GenericIdentification47 {
	newValue := new(GenericIdentification47)
	c.ExemptionType = append(c.ExemptionType, newValue)
	return newValue
}

func (c *CashOption46) SetCountryOfIncomeSource(value string) {
	c.CountryOfIncomeSource = (*CountryCode)(&value)
}

func (c *CashOption46) AddAccount() *Account9Choice {
	c.Account = new(Account9Choice)
	return c.Account
}

func (c *CashOption46) AddCashParties() *CashParties29 {
	c.CashParties = new(CashParties29)
	return c.CashParties
}

func (c *CashOption46) AddAmountDetails() *CorporateActionAmounts39 {
	c.AmountDetails = new(CorporateActionAmounts39)
	return c.AmountDetails
}

func (c *CashOption46) AddDateDetails() *CorporateActionDate24 {
	c.DateDetails = new(CorporateActionDate24)
	return c.DateDetails
}

func (c *CashOption46) AddForeignExchangeDetails() *ForeignExchangeTerms27 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms27)
	return c.ForeignExchangeDetails
}

func (c *CashOption46) AddTaxVoucherDetails() *TaxVoucher3 {
	c.TaxVoucherDetails = new(TaxVoucher3)
	return c.TaxVoucherDetails
}

func (c *CashOption46) AddRateAndAmountDetails() *RateDetails24 {
	c.RateAndAmountDetails = new(RateDetails24)
	return c.RateAndAmountDetails
}

func (c *CashOption46) AddPriceDetails() *PriceDetails24 {
	c.PriceDetails = new(PriceDetails24)
	return c.PriceDetails
}

// Provides information about the cash option.
type CashOption47 struct {

	// Indicates whether the value is a debit or a credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates whether the cash payment occurs or will occur in advance of receipt of proceeds from the issuer and based on a contractual agreement established with the account servicer or upon receipt of proceeds from the issuer.
	ContractualPaymentIndicator *Payment1Code `xml:"CtrctlPmtInd,omitempty"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator4Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *GenericIdentification30 `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification47 `xml:"IncmTp,omitempty"`

	// Specifies the basis for the reduced rate of withholding.
	ExemptionType []*GenericIdentification47 `xml:"XmptnTp,omitempty"`

	// Indicates the country from which the income originates.
	CountryOfIncomeSource *CountryCode `xml:"CtryOfIncmSrc,omitempty"`

	// Identification of the account in which cash is maintained.
	CashAccountIdentification *CashAccountIdentification6Choice `xml:"CshAcctId,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts40 `xml:"AmtDtls,omitempty"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate56 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount.
	ForeignExchangeDetails *ForeignExchangeTerms28 `xml:"FXDtls,omitempty"`

	// Provides information about the corporate action option.
	RateAndAmountDetails *RateDetails25 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *PriceDetails25 `xml:"PricDtls,omitempty"`
}

func (c *CashOption47) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption47) SetContractualPaymentIndicator(value string) {
	c.ContractualPaymentIndicator = (*Payment1Code)(&value)
}

func (c *CashOption47) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator4Choice {
	c.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator4Choice)
	return c.NonEligibleProceedsIndicator
}

func (c *CashOption47) AddIssuerOfferorTaxabilityIndicator() *GenericIdentification30 {
	c.IssuerOfferorTaxabilityIndicator = new(GenericIdentification30)
	return c.IssuerOfferorTaxabilityIndicator
}

func (c *CashOption47) AddIncomeType() *GenericIdentification47 {
	c.IncomeType = new(GenericIdentification47)
	return c.IncomeType
}

func (c *CashOption47) AddExemptionType() *GenericIdentification47 {
	newValue := new(GenericIdentification47)
	c.ExemptionType = append(c.ExemptionType, newValue)
	return newValue
}

func (c *CashOption47) SetCountryOfIncomeSource(value string) {
	c.CountryOfIncomeSource = (*CountryCode)(&value)
}

func (c *CashOption47) AddCashAccountIdentification() *CashAccountIdentification6Choice {
	c.CashAccountIdentification = new(CashAccountIdentification6Choice)
	return c.CashAccountIdentification
}

func (c *CashOption47) AddAmountDetails() *CorporateActionAmounts40 {
	c.AmountDetails = new(CorporateActionAmounts40)
	return c.AmountDetails
}

func (c *CashOption47) AddDateDetails() *CorporateActionDate56 {
	c.DateDetails = new(CorporateActionDate56)
	return c.DateDetails
}

func (c *CashOption47) AddForeignExchangeDetails() *ForeignExchangeTerms28 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms28)
	return c.ForeignExchangeDetails
}

func (c *CashOption47) AddRateAndAmountDetails() *RateDetails25 {
	c.RateAndAmountDetails = new(RateDetails25)
	return c.RateAndAmountDetails
}

func (c *CashOption47) AddPriceDetails() *PriceDetails25 {
	c.PriceDetails = new(PriceDetails25)
	return c.PriceDetails
}

// Provides information about the cash option.
type CashOption48 struct {

	// Indicates whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Choice between a cash account, a charges account or a tax account.
	Account *Account9Choice `xml:"Acct,omitempty"`

	// Date/Time of the posting (credit or debit) to the account.
	PostingDate *DateAndDateTimeChoice `xml:"PstngDt"`

	// Date/Time of the posting (credit or debit) to the account that was initially communicated in the confirmation.
	OriginalPostingDate *DateAndDateTimeChoice `xml:"OrgnlPstngDt,omitempty"`

	// Date/time when calculating economic benefit for a cash amount.
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`

	// Amount of money that is to be/was posted to the account.
	PostingAmount *RestrictedFINActiveCurrencyAndAmount `xml:"PstngAmt"`
}

func (c *CashOption48) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption48) AddAccount() *Account9Choice {
	c.Account = new(Account9Choice)
	return c.Account
}

func (c *CashOption48) AddPostingDate() *DateAndDateTimeChoice {
	c.PostingDate = new(DateAndDateTimeChoice)
	return c.PostingDate
}

func (c *CashOption48) AddOriginalPostingDate() *DateAndDateTimeChoice {
	c.OriginalPostingDate = new(DateAndDateTimeChoice)
	return c.OriginalPostingDate
}

func (c *CashOption48) AddValueDate() *DateAndDateTimeChoice {
	c.ValueDate = new(DateAndDateTimeChoice)
	return c.ValueDate
}

func (c *CashOption48) SetPostingAmount(value, currency string) {
	c.PostingAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

// Provides information about the cash option.
type CashOption49 struct {

	// Indicates whether the value is a debit or a credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator4Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *GenericIdentification47 `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification47 `xml:"IncmTp,omitempty"`

	// Specifies the basis for the reduced rate of withholding.
	ExemptionType []*GenericIdentification47 `xml:"XmptnTp,omitempty"`

	// Indicates the country from which the income originates.
	CountryOfIncomeSource *CountryCode `xml:"CtryOfIncmSrc,omitempty"`

	// Identification of the account in which cash is maintained.
	CashAccountIdentification *CashAccountIdentification6Choice `xml:"CshAcctId,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts41 `xml:"AmtDtls,omitempty"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate56 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount
	ForeignExchangeDetails *ForeignExchangeTerms28 `xml:"FXDtls,omitempty"`

	// Provides information about the corporate action option.
	RateAndAmountDetails *RateDetails25 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *PriceDetails25 `xml:"PricDtls,omitempty"`
}

func (c *CashOption49) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption49) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator4Choice {
	c.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator4Choice)
	return c.NonEligibleProceedsIndicator
}

func (c *CashOption49) AddIssuerOfferorTaxabilityIndicator() *GenericIdentification47 {
	c.IssuerOfferorTaxabilityIndicator = new(GenericIdentification47)
	return c.IssuerOfferorTaxabilityIndicator
}

func (c *CashOption49) AddIncomeType() *GenericIdentification47 {
	c.IncomeType = new(GenericIdentification47)
	return c.IncomeType
}

func (c *CashOption49) AddExemptionType() *GenericIdentification47 {
	newValue := new(GenericIdentification47)
	c.ExemptionType = append(c.ExemptionType, newValue)
	return newValue
}

func (c *CashOption49) SetCountryOfIncomeSource(value string) {
	c.CountryOfIncomeSource = (*CountryCode)(&value)
}

func (c *CashOption49) AddCashAccountIdentification() *CashAccountIdentification6Choice {
	c.CashAccountIdentification = new(CashAccountIdentification6Choice)
	return c.CashAccountIdentification
}

func (c *CashOption49) AddAmountDetails() *CorporateActionAmounts41 {
	c.AmountDetails = new(CorporateActionAmounts41)
	return c.AmountDetails
}

func (c *CashOption49) AddDateDetails() *CorporateActionDate56 {
	c.DateDetails = new(CorporateActionDate56)
	return c.DateDetails
}

func (c *CashOption49) AddForeignExchangeDetails() *ForeignExchangeTerms28 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms28)
	return c.ForeignExchangeDetails
}

func (c *CashOption49) AddRateAndAmountDetails() *RateDetails25 {
	c.RateAndAmountDetails = new(RateDetails25)
	return c.RateAndAmountDetails
}

func (c *CashOption49) AddPriceDetails() *PriceDetails25 {
	c.PriceDetails = new(PriceDetails25)
	return c.PriceDetails
}

// Provides information about the cash option.
type CashOption5 struct {

	// Indicates whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Choice between a cash account, a charges account or a tax account.
	Account *Account8Choice `xml:"Acct,omitempty"`

	// Date/Time of the posting (credit or debit) to the account.
	PostingDate *DateAndDateTimeChoice `xml:"PstngDt"`

	// Date/Time of the posting (credit or debit) to the account that was initially communicated in the confirmation.
	OriginalPostingDate *DateAndDateTimeChoice `xml:"OrgnlPstngDt,omitempty"`

	// Date/time at which assets become available to the account owner (in a credit entry), or cease to be available to the account owner (in a debit entry).
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`

	// Amount of money that is to be/was posted to the account.
	PostingAmount *ActiveCurrencyAndAmount `xml:"PstngAmt"`
}

func (c *CashOption5) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption5) AddAccount() *Account8Choice {
	c.Account = new(Account8Choice)
	return c.Account
}

func (c *CashOption5) AddPostingDate() *DateAndDateTimeChoice {
	c.PostingDate = new(DateAndDateTimeChoice)
	return c.PostingDate
}

func (c *CashOption5) AddOriginalPostingDate() *DateAndDateTimeChoice {
	c.OriginalPostingDate = new(DateAndDateTimeChoice)
	return c.OriginalPostingDate
}

func (c *CashOption5) AddValueDate() *DateAndDateTimeChoice {
	c.ValueDate = new(DateAndDateTimeChoice)
	return c.ValueDate
}

func (c *CashOption5) SetPostingAmount(value, currency string) {
	c.PostingAmount = NewActiveCurrencyAndAmount(value, currency)
}

// Provides information about the cash option.
type CashOption50 struct {

	// Indicates whether the value is a debit or a credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator3Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *IssuerOfferorTaxabilityIndicator1Choice `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification30 `xml:"IncmTp,omitempty"`

	// Specifies the basis for the reduced rate of withholding.
	ExemptionType []*GenericIdentification30 `xml:"XmptnTp,omitempty"`

	// Indicates the country from which the income originates.
	CountryOfIncomeSource *CountryCode `xml:"CtryOfIncmSrc,omitempty"`

	// Identification of the account in which cash is maintained.
	CashAccountIdentification *CashAccountIdentification5Choice `xml:"CshAcctId,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts36 `xml:"AmtDtls,omitempty"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate47 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount
	ForeignExchangeDetails *ForeignExchangeTerms24 `xml:"FXDtls,omitempty"`

	// Provides information about the corporate action option.
	RateAndAmountDetails *RateDetails26 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *PriceDetails22 `xml:"PricDtls,omitempty"`
}

func (c *CashOption50) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption50) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator3Choice {
	c.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator3Choice)
	return c.NonEligibleProceedsIndicator
}

func (c *CashOption50) AddIssuerOfferorTaxabilityIndicator() *IssuerOfferorTaxabilityIndicator1Choice {
	c.IssuerOfferorTaxabilityIndicator = new(IssuerOfferorTaxabilityIndicator1Choice)
	return c.IssuerOfferorTaxabilityIndicator
}

func (c *CashOption50) AddIncomeType() *GenericIdentification30 {
	c.IncomeType = new(GenericIdentification30)
	return c.IncomeType
}

func (c *CashOption50) AddExemptionType() *GenericIdentification30 {
	newValue := new(GenericIdentification30)
	c.ExemptionType = append(c.ExemptionType, newValue)
	return newValue
}

func (c *CashOption50) SetCountryOfIncomeSource(value string) {
	c.CountryOfIncomeSource = (*CountryCode)(&value)
}

func (c *CashOption50) AddCashAccountIdentification() *CashAccountIdentification5Choice {
	c.CashAccountIdentification = new(CashAccountIdentification5Choice)
	return c.CashAccountIdentification
}

func (c *CashOption50) AddAmountDetails() *CorporateActionAmounts36 {
	c.AmountDetails = new(CorporateActionAmounts36)
	return c.AmountDetails
}

func (c *CashOption50) AddDateDetails() *CorporateActionDate47 {
	c.DateDetails = new(CorporateActionDate47)
	return c.DateDetails
}

func (c *CashOption50) AddForeignExchangeDetails() *ForeignExchangeTerms24 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms24)
	return c.ForeignExchangeDetails
}

func (c *CashOption50) AddRateAndAmountDetails() *RateDetails26 {
	c.RateAndAmountDetails = new(RateDetails26)
	return c.RateAndAmountDetails
}

func (c *CashOption50) AddPriceDetails() *PriceDetails22 {
	c.PriceDetails = new(PriceDetails22)
	return c.PriceDetails
}

// Provides information about the cash option.
type CashOption51 struct {

	// Indicates whether the value is a debit or a credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates whether the cash payment occurs or will occur in advance of receipt of proceeds from the issuer and based on a contractual agreement established with the account servicer or upon receipt of proceeds from the issuer.
	ContractualPaymentIndicator *Payment1Code `xml:"CtrctlPmtInd,omitempty"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator3Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *IssuerOfferorTaxabilityIndicator1Choice `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification30 `xml:"IncmTp,omitempty"`

	// Specifies the basis for the reduced rate of withholding.
	ExemptionType []*GenericIdentification30 `xml:"XmptnTp,omitempty"`

	// Indicates the country from which the income originates.
	CountryOfIncomeSource *CountryCode `xml:"CtryOfIncmSrc,omitempty"`

	// Identification of the account in which cash is maintained.
	CashAccountIdentification *CashAccountIdentification5Choice `xml:"CshAcctId,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts38 `xml:"AmtDtls,omitempty"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate47 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount.
	ForeignExchangeDetails *ForeignExchangeTerms24 `xml:"FXDtls,omitempty"`

	// Provides information about the corporate action option.
	RateAndAmountDetails *RateDetails26 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *PriceDetails22 `xml:"PricDtls,omitempty"`
}

func (c *CashOption51) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption51) SetContractualPaymentIndicator(value string) {
	c.ContractualPaymentIndicator = (*Payment1Code)(&value)
}

func (c *CashOption51) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator3Choice {
	c.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator3Choice)
	return c.NonEligibleProceedsIndicator
}

func (c *CashOption51) AddIssuerOfferorTaxabilityIndicator() *IssuerOfferorTaxabilityIndicator1Choice {
	c.IssuerOfferorTaxabilityIndicator = new(IssuerOfferorTaxabilityIndicator1Choice)
	return c.IssuerOfferorTaxabilityIndicator
}

func (c *CashOption51) AddIncomeType() *GenericIdentification30 {
	c.IncomeType = new(GenericIdentification30)
	return c.IncomeType
}

func (c *CashOption51) AddExemptionType() *GenericIdentification30 {
	newValue := new(GenericIdentification30)
	c.ExemptionType = append(c.ExemptionType, newValue)
	return newValue
}

func (c *CashOption51) SetCountryOfIncomeSource(value string) {
	c.CountryOfIncomeSource = (*CountryCode)(&value)
}

func (c *CashOption51) AddCashAccountIdentification() *CashAccountIdentification5Choice {
	c.CashAccountIdentification = new(CashAccountIdentification5Choice)
	return c.CashAccountIdentification
}

func (c *CashOption51) AddAmountDetails() *CorporateActionAmounts38 {
	c.AmountDetails = new(CorporateActionAmounts38)
	return c.AmountDetails
}

func (c *CashOption51) AddDateDetails() *CorporateActionDate47 {
	c.DateDetails = new(CorporateActionDate47)
	return c.DateDetails
}

func (c *CashOption51) AddForeignExchangeDetails() *ForeignExchangeTerms24 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms24)
	return c.ForeignExchangeDetails
}

func (c *CashOption51) AddRateAndAmountDetails() *RateDetails26 {
	c.RateAndAmountDetails = new(RateDetails26)
	return c.RateAndAmountDetails
}

func (c *CashOption51) AddPriceDetails() *PriceDetails22 {
	c.PriceDetails = new(PriceDetails22)
	return c.PriceDetails
}

// Provides information about the cash option.
type CashOption52 struct {

	// Indicates whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates whether the cash payment occurs or will occur in advance of receipt of proceeds from the issuer and based on a contractual agreement established with the account servicer or upon receipt of proceeds from the issuer.
	ContractualPaymentIndicator *Payment1Code `xml:"CtrctlPmtInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *IssuerOfferorTaxabilityIndicator1Choice `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification30 `xml:"IncmTp,omitempty"`

	// Specifies the basis for the reduced rate of withholding.
	ExemptionType []*GenericIdentification30 `xml:"XmptnTp,omitempty"`

	// Indicates the country from which the income originates.
	CountryOfIncomeSource *CountryCode `xml:"CtryOfIncmSrc,omitempty"`

	// Choice between a cash account, a charges account or a tax account.
	Account *Account8Choice `xml:"Acct,omitempty"`

	// Provides information about cash parties.
	CashParties *CashParties28 `xml:"CshPties,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts37 `xml:"AmtDtls"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate24 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount.
	ForeignExchangeDetails *ForeignExchangeTerms23 `xml:"FXDtls,omitempty"`

	// Provides information about the tax voucher related to a cash movement.
	TaxVoucherDetails *TaxVoucher2 `xml:"TaxVchrDtls,omitempty"`

	// Provides information about the corporate action option.
	RateAndAmountDetails *RateDetails27 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *PriceDetails23 `xml:"PricDtls,omitempty"`
}

func (c *CashOption52) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption52) SetContractualPaymentIndicator(value string) {
	c.ContractualPaymentIndicator = (*Payment1Code)(&value)
}

func (c *CashOption52) AddIssuerOfferorTaxabilityIndicator() *IssuerOfferorTaxabilityIndicator1Choice {
	c.IssuerOfferorTaxabilityIndicator = new(IssuerOfferorTaxabilityIndicator1Choice)
	return c.IssuerOfferorTaxabilityIndicator
}

func (c *CashOption52) AddIncomeType() *GenericIdentification30 {
	c.IncomeType = new(GenericIdentification30)
	return c.IncomeType
}

func (c *CashOption52) AddExemptionType() *GenericIdentification30 {
	newValue := new(GenericIdentification30)
	c.ExemptionType = append(c.ExemptionType, newValue)
	return newValue
}

func (c *CashOption52) SetCountryOfIncomeSource(value string) {
	c.CountryOfIncomeSource = (*CountryCode)(&value)
}

func (c *CashOption52) AddAccount() *Account8Choice {
	c.Account = new(Account8Choice)
	return c.Account
}

func (c *CashOption52) AddCashParties() *CashParties28 {
	c.CashParties = new(CashParties28)
	return c.CashParties
}

func (c *CashOption52) AddAmountDetails() *CorporateActionAmounts37 {
	c.AmountDetails = new(CorporateActionAmounts37)
	return c.AmountDetails
}

func (c *CashOption52) AddDateDetails() *CorporateActionDate24 {
	c.DateDetails = new(CorporateActionDate24)
	return c.DateDetails
}

func (c *CashOption52) AddForeignExchangeDetails() *ForeignExchangeTerms23 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms23)
	return c.ForeignExchangeDetails
}

func (c *CashOption52) AddTaxVoucherDetails() *TaxVoucher2 {
	c.TaxVoucherDetails = new(TaxVoucher2)
	return c.TaxVoucherDetails
}

func (c *CashOption52) AddRateAndAmountDetails() *RateDetails27 {
	c.RateAndAmountDetails = new(RateDetails27)
	return c.RateAndAmountDetails
}

func (c *CashOption52) AddPriceDetails() *PriceDetails23 {
	c.PriceDetails = new(PriceDetails23)
	return c.PriceDetails
}

// Provides information about the cash option.
type CashOption53 struct {

	// Indicates whether the value is a debit or a credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator4Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *IssuerOfferorTaxabilityIndicator1Choice `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification47 `xml:"IncmTp,omitempty"`

	// Specifies the basis for the reduced rate of withholding.
	ExemptionType []*GenericIdentification47 `xml:"XmptnTp,omitempty"`

	// Indicates the country from which the income originates.
	CountryOfIncomeSource *CountryCode `xml:"CtryOfIncmSrc,omitempty"`

	// Identification of the account in which cash is maintained.
	CashAccountIdentification *CashAccountIdentification6Choice `xml:"CshAcctId,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts41 `xml:"AmtDtls,omitempty"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate56 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount
	ForeignExchangeDetails *ForeignExchangeTerms28 `xml:"FXDtls,omitempty"`

	// Provides information about the corporate action option.
	RateAndAmountDetails *RateDetails28 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *PriceDetails25 `xml:"PricDtls,omitempty"`
}

func (c *CashOption53) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption53) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator4Choice {
	c.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator4Choice)
	return c.NonEligibleProceedsIndicator
}

func (c *CashOption53) AddIssuerOfferorTaxabilityIndicator() *IssuerOfferorTaxabilityIndicator1Choice {
	c.IssuerOfferorTaxabilityIndicator = new(IssuerOfferorTaxabilityIndicator1Choice)
	return c.IssuerOfferorTaxabilityIndicator
}

func (c *CashOption53) AddIncomeType() *GenericIdentification47 {
	c.IncomeType = new(GenericIdentification47)
	return c.IncomeType
}

func (c *CashOption53) AddExemptionType() *GenericIdentification47 {
	newValue := new(GenericIdentification47)
	c.ExemptionType = append(c.ExemptionType, newValue)
	return newValue
}

func (c *CashOption53) SetCountryOfIncomeSource(value string) {
	c.CountryOfIncomeSource = (*CountryCode)(&value)
}

func (c *CashOption53) AddCashAccountIdentification() *CashAccountIdentification6Choice {
	c.CashAccountIdentification = new(CashAccountIdentification6Choice)
	return c.CashAccountIdentification
}

func (c *CashOption53) AddAmountDetails() *CorporateActionAmounts41 {
	c.AmountDetails = new(CorporateActionAmounts41)
	return c.AmountDetails
}

func (c *CashOption53) AddDateDetails() *CorporateActionDate56 {
	c.DateDetails = new(CorporateActionDate56)
	return c.DateDetails
}

func (c *CashOption53) AddForeignExchangeDetails() *ForeignExchangeTerms28 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms28)
	return c.ForeignExchangeDetails
}

func (c *CashOption53) AddRateAndAmountDetails() *RateDetails28 {
	c.RateAndAmountDetails = new(RateDetails28)
	return c.RateAndAmountDetails
}

func (c *CashOption53) AddPriceDetails() *PriceDetails25 {
	c.PriceDetails = new(PriceDetails25)
	return c.PriceDetails
}

// Provides information about the cash option.
type CashOption54 struct {

	// Indicates whether the value is a debit or a credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates whether the cash payment occurs or will occur in advance of receipt of proceeds from the issuer and based on a contractual agreement established with the account servicer or upon receipt of proceeds from the issuer.
	ContractualPaymentIndicator *Payment1Code `xml:"CtrctlPmtInd,omitempty"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator4Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *IssuerOfferorTaxabilityIndicator1Choice `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification47 `xml:"IncmTp,omitempty"`

	// Specifies the basis for the reduced rate of withholding.
	ExemptionType []*GenericIdentification47 `xml:"XmptnTp,omitempty"`

	// Indicates the country from which the income originates.
	CountryOfIncomeSource *CountryCode `xml:"CtryOfIncmSrc,omitempty"`

	// Identification of the account in which cash is maintained.
	CashAccountIdentification *CashAccountIdentification6Choice `xml:"CshAcctId,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts40 `xml:"AmtDtls,omitempty"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate56 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount.
	ForeignExchangeDetails *ForeignExchangeTerms28 `xml:"FXDtls,omitempty"`

	// Provides information about the corporate action option.
	RateAndAmountDetails *RateDetails28 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *PriceDetails25 `xml:"PricDtls,omitempty"`
}

func (c *CashOption54) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption54) SetContractualPaymentIndicator(value string) {
	c.ContractualPaymentIndicator = (*Payment1Code)(&value)
}

func (c *CashOption54) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator4Choice {
	c.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator4Choice)
	return c.NonEligibleProceedsIndicator
}

func (c *CashOption54) AddIssuerOfferorTaxabilityIndicator() *IssuerOfferorTaxabilityIndicator1Choice {
	c.IssuerOfferorTaxabilityIndicator = new(IssuerOfferorTaxabilityIndicator1Choice)
	return c.IssuerOfferorTaxabilityIndicator
}

func (c *CashOption54) AddIncomeType() *GenericIdentification47 {
	c.IncomeType = new(GenericIdentification47)
	return c.IncomeType
}

func (c *CashOption54) AddExemptionType() *GenericIdentification47 {
	newValue := new(GenericIdentification47)
	c.ExemptionType = append(c.ExemptionType, newValue)
	return newValue
}

func (c *CashOption54) SetCountryOfIncomeSource(value string) {
	c.CountryOfIncomeSource = (*CountryCode)(&value)
}

func (c *CashOption54) AddCashAccountIdentification() *CashAccountIdentification6Choice {
	c.CashAccountIdentification = new(CashAccountIdentification6Choice)
	return c.CashAccountIdentification
}

func (c *CashOption54) AddAmountDetails() *CorporateActionAmounts40 {
	c.AmountDetails = new(CorporateActionAmounts40)
	return c.AmountDetails
}

func (c *CashOption54) AddDateDetails() *CorporateActionDate56 {
	c.DateDetails = new(CorporateActionDate56)
	return c.DateDetails
}

func (c *CashOption54) AddForeignExchangeDetails() *ForeignExchangeTerms28 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms28)
	return c.ForeignExchangeDetails
}

func (c *CashOption54) AddRateAndAmountDetails() *RateDetails28 {
	c.RateAndAmountDetails = new(RateDetails28)
	return c.RateAndAmountDetails
}

func (c *CashOption54) AddPriceDetails() *PriceDetails25 {
	c.PriceDetails = new(PriceDetails25)
	return c.PriceDetails
}

// Provides information about the cash option.
type CashOption55 struct {

	// Indicates whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates whether the cash payment occurs or will occur in advance of receipt of proceeds from the issuer and based on a contractual agreement established with the account servicer or upon receipt of proceeds from the issuer.
	ContractualPaymentIndicator *Payment1Code `xml:"CtrctlPmtInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *IssuerOfferorTaxabilityIndicator1Choice `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification47 `xml:"IncmTp,omitempty"`

	// Specifies the basis for the reduced rate of withholding.
	ExemptionType []*GenericIdentification47 `xml:"XmptnTp,omitempty"`

	// Indicates the country from which the income originates.
	CountryOfIncomeSource *CountryCode `xml:"CtryOfIncmSrc,omitempty"`

	// Choice between a cash account, a charges account or a tax account.
	Account *Account9Choice `xml:"Acct,omitempty"`

	// Provides information about cash parties.
	CashParties *CashParties29 `xml:"CshPties,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts39 `xml:"AmtDtls"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate24 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount.
	ForeignExchangeDetails *ForeignExchangeTerms27 `xml:"FXDtls,omitempty"`

	// Provides information about the tax voucher related to a cash movement.
	TaxVoucherDetails *TaxVoucher3 `xml:"TaxVchrDtls,omitempty"`

	// Provides information about the corporate action option.
	RateAndAmountDetails *RateDetails30 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *PriceDetails24 `xml:"PricDtls,omitempty"`
}

func (c *CashOption55) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption55) SetContractualPaymentIndicator(value string) {
	c.ContractualPaymentIndicator = (*Payment1Code)(&value)
}

func (c *CashOption55) AddIssuerOfferorTaxabilityIndicator() *IssuerOfferorTaxabilityIndicator1Choice {
	c.IssuerOfferorTaxabilityIndicator = new(IssuerOfferorTaxabilityIndicator1Choice)
	return c.IssuerOfferorTaxabilityIndicator
}

func (c *CashOption55) AddIncomeType() *GenericIdentification47 {
	c.IncomeType = new(GenericIdentification47)
	return c.IncomeType
}

func (c *CashOption55) AddExemptionType() *GenericIdentification47 {
	newValue := new(GenericIdentification47)
	c.ExemptionType = append(c.ExemptionType, newValue)
	return newValue
}

func (c *CashOption55) SetCountryOfIncomeSource(value string) {
	c.CountryOfIncomeSource = (*CountryCode)(&value)
}

func (c *CashOption55) AddAccount() *Account9Choice {
	c.Account = new(Account9Choice)
	return c.Account
}

func (c *CashOption55) AddCashParties() *CashParties29 {
	c.CashParties = new(CashParties29)
	return c.CashParties
}

func (c *CashOption55) AddAmountDetails() *CorporateActionAmounts39 {
	c.AmountDetails = new(CorporateActionAmounts39)
	return c.AmountDetails
}

func (c *CashOption55) AddDateDetails() *CorporateActionDate24 {
	c.DateDetails = new(CorporateActionDate24)
	return c.DateDetails
}

func (c *CashOption55) AddForeignExchangeDetails() *ForeignExchangeTerms27 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms27)
	return c.ForeignExchangeDetails
}

func (c *CashOption55) AddTaxVoucherDetails() *TaxVoucher3 {
	c.TaxVoucherDetails = new(TaxVoucher3)
	return c.TaxVoucherDetails
}

func (c *CashOption55) AddRateAndAmountDetails() *RateDetails30 {
	c.RateAndAmountDetails = new(RateDetails30)
	return c.RateAndAmountDetails
}

func (c *CashOption55) AddPriceDetails() *PriceDetails24 {
	c.PriceDetails = new(PriceDetails24)
	return c.PriceDetails
}
