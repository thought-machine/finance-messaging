package iso20022

// Amount of money debited or credited on the books of an account servicer.
type AmountAndDirection10 struct {

	// Amount of money in the cash entry.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalCurrencyAndOrderedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlCcyAndOrdrdAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms11 `xml:"FXDtls,omitempty"`
}

func (a *AmountAndDirection10) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection10) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection10) SetOriginalCurrencyAndOrderedAmount(value, currency string) {
	a.OriginalCurrencyAndOrderedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection10) AddForeignExchangeDetails() *ForeignExchangeTerms11 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms11)
	return a.ForeignExchangeDetails
}

// Posting of an item to a cash account, in the context of a cash transaction, that results in an increase or decrease to the balance of the account.
type AmountAndDirection14 struct {

	// Amount of money in the cash entry.
	Amount *RestrictedFINActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates that the amount value is positive or negative.
	Sign *PlusOrMinusIndicator `xml:"Sgn"`
}

func (a *AmountAndDirection14) SetAmount(value, currency string) {
	a.Amount = NewRestrictedFINActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection14) SetSign(value string) {
	a.Sign = (*PlusOrMinusIndicator)(&value)
}

// Posting of an item to a cash account, in the context of a cash transaction, that results in an increase or decrease to the balance of the account.
type AmountAndDirection18 struct {

	// Amount of money in the cash entry.
	Amount *RestrictedFINActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebit *CreditDebitCode `xml:"CdtDbt"`
}

func (a *AmountAndDirection18) SetAmount(value, currency string) {
	a.Amount = NewRestrictedFINActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection18) SetCreditDebit(value string) {
	a.CreditDebit = (*CreditDebitCode)(&value)
}

// Amount of money debited or credited on the books of an account servicer.
type AmountAndDirection19 struct {

	// Amount of money that is debited or credited.
	Amount *RestrictedFINActiveCurrencyAndAmount `xml:"Amt"`

	// Indicates if the amount is a debited or a credited.
	CreditDebit *CreditDebitCode `xml:"CdtDbt,omitempty"`
}

func (a *AmountAndDirection19) SetAmount(value, currency string) {
	a.Amount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection19) SetCreditDebit(value string) {
	a.CreditDebit = (*CreditDebitCode)(&value)
}

// Posting of an item to a cash account, in the context of a cash transaction, that results in an increase or decrease to the balance of the account.
type AmountAndDirection2 struct {

	// Indicates whether the net proceeds include interest accrued on the financial instrument.
	AccruedInterestIndicator *YesNoIndicator `xml:"AcrdIntrstInd,omitempty"`

	// Whether the net proceeds include stamp duty amount.
	StampDutyIndicator *YesNoIndicator `xml:"StmpDtyInd,omitempty"`

	// Amount of money in the cash entry.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalCurrencyAndOrderedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlCcyAndOrdrdAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms11 `xml:"FXDtls,omitempty"`

	// Date and time at which the cash is at the disposal of the credit account owner, or ceases to be at the disposal of the debit account owner.
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`
}

func (a *AmountAndDirection2) SetAccruedInterestIndicator(value string) {
	a.AccruedInterestIndicator = (*YesNoIndicator)(&value)
}

func (a *AmountAndDirection2) SetStampDutyIndicator(value string) {
	a.StampDutyIndicator = (*YesNoIndicator)(&value)
}

func (a *AmountAndDirection2) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection2) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection2) SetOriginalCurrencyAndOrderedAmount(value, currency string) {
	a.OriginalCurrencyAndOrderedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection2) AddForeignExchangeDetails() *ForeignExchangeTerms11 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms11)
	return a.ForeignExchangeDetails
}

func (a *AmountAndDirection2) AddValueDate() *DateAndDateTimeChoice {
	a.ValueDate = new(DateAndDateTimeChoice)
	return a.ValueDate
}

// Amount of money debited or credited on the books of an account servicer.
type AmountAndDirection20 struct {

	// Total amount that needs to be settled.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`
}

func (a *AmountAndDirection20) SetAmount(value, currency string) {
	a.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection20) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

// Amount of money debited or credited on the books of an account servicer.
type AmountAndDirection21 struct {

	// Amount of money in the cash entry.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`
}

func (a *AmountAndDirection21) SetAmount(value, currency string) {
	a.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection21) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

// Posting of an item to a cash account, in the context of a cash transaction, that results in an increase or decrease to the balance of the account.
type AmountAndDirection22 struct {

	// Amount of money in the cash entry.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalCurrencyAndOrderedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlCcyAndOrdrdAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms11 `xml:"FXDtls,omitempty"`

	// Date and time at which the cash is at the disposal of the credit account owner, or ceases to be at the disposal of the debit account owner.
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`
}

func (a *AmountAndDirection22) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection22) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection22) SetOriginalCurrencyAndOrderedAmount(value, currency string) {
	a.OriginalCurrencyAndOrderedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection22) AddForeignExchangeDetails() *ForeignExchangeTerms11 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms11)
	return a.ForeignExchangeDetails
}

func (a *AmountAndDirection22) AddValueDate() *DateAndDateTimeChoice {
	a.ValueDate = new(DateAndDateTimeChoice)
	return a.ValueDate
}

// Amount of money debited or credited on the books of an account servicer.
type AmountAndDirection23 struct {

	// Amount of money in the cash entry.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms11 `xml:"FXDtls,omitempty"`
}

func (a *AmountAndDirection23) SetAmount(value, currency string) {
	a.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection23) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection23) AddForeignExchangeDetails() *ForeignExchangeTerms11 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms11)
	return a.ForeignExchangeDetails
}

// Amount of money debited or credited on the books of an account servicer.
type AmountAndDirection27 struct {

	// Amount of money in the cash entry.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalCurrencyAndOrderedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlCcyAndOrdrdAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms17 `xml:"FXDtls,omitempty"`
}

func (a *AmountAndDirection27) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection27) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection27) SetOriginalCurrencyAndOrderedAmount(value, currency string) {
	a.OriginalCurrencyAndOrderedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection27) AddForeignExchangeDetails() *ForeignExchangeTerms17 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms17)
	return a.ForeignExchangeDetails
}

// Posting of an item to a cash account, in the context of a cash transaction, that results in an increase or decrease to the balance of the account.
type AmountAndDirection28 struct {

	// Indicates whether the net proceeds include interest accrued on the financial instrument.
	AccruedInterestIndicator *YesNoIndicator `xml:"AcrdIntrstInd,omitempty"`

	// Indicates whether the settlement amount includes the stamp duty amount.
	StampDutyIndicator *YesNoIndicator `xml:"StmpDtyInd,omitempty"`

	// Amount of money in the cash entry.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalCurrencyAndOrderedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlCcyAndOrdrdAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms18 `xml:"FXDtls,omitempty"`

	// Date and time at which the cash is at the disposal of the credit account owner, or ceases to be at the disposal of the debit account owner.
	ValueDate *DateAndDateTime1Choice `xml:"ValDt,omitempty"`
}

func (a *AmountAndDirection28) SetAccruedInterestIndicator(value string) {
	a.AccruedInterestIndicator = (*YesNoIndicator)(&value)
}

func (a *AmountAndDirection28) SetStampDutyIndicator(value string) {
	a.StampDutyIndicator = (*YesNoIndicator)(&value)
}

func (a *AmountAndDirection28) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection28) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection28) SetOriginalCurrencyAndOrderedAmount(value, currency string) {
	a.OriginalCurrencyAndOrderedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection28) AddForeignExchangeDetails() *ForeignExchangeTerms18 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms18)
	return a.ForeignExchangeDetails
}

func (a *AmountAndDirection28) AddValueDate() *DateAndDateTime1Choice {
	a.ValueDate = new(DateAndDateTime1Choice)
	return a.ValueDate
}

// Amount of money debited or credited on the books of an account servicer.
type AmountAndDirection29 struct {

	// Amount of money in the cash entry.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalCurrencyAndOrderedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlCcyAndOrdrdAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms18 `xml:"FXDtls,omitempty"`
}

func (a *AmountAndDirection29) SetAmount(value, currency string) {
	a.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection29) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection29) SetOriginalCurrencyAndOrderedAmount(value, currency string) {
	a.OriginalCurrencyAndOrderedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection29) AddForeignExchangeDetails() *ForeignExchangeTerms18 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms18)
	return a.ForeignExchangeDetails
}

// Posting of an item to a cash account, in the context of a cash transaction, that results in an increase or decrease to the balance of the account.
type AmountAndDirection3 struct {

	// Amount of money in the cash entry.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebit *CreditDebitCode `xml:"CdtDbt"`
}

func (a *AmountAndDirection3) SetAmount(value, currency string) {
	a.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection3) SetCreditDebit(value string) {
	a.CreditDebit = (*CreditDebitCode)(&value)
}

// Posting of an item to a cash account, in the context of a cash transaction, that results in an increase or decrease to the balance of the account.
type AmountAndDirection30 struct {

	// Amount of money in the cash entry.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates that the amount value is positive or negative.
	Sign *PlusOrMinusIndicator `xml:"Sgn,omitempty"`
}

func (a *AmountAndDirection30) SetAmount(value, currency string) {
	a.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection30) SetSign(value string) {
	a.Sign = (*PlusOrMinusIndicator)(&value)
}

// Posting of an item to a cash account, in the context of a cash transaction, that results in an increase or decrease to the balance of the account.
type AmountAndDirection31 struct {

	// Currency and value.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indication that the position is short or long.
	ShortLongIndicator *ShortLong1Code `xml:"ShrtLngInd,omitempty"`
}

func (a *AmountAndDirection31) SetAmount(value, currency string) {
	a.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection31) SetShortLongIndicator(value string) {
	a.ShortLongIndicator = (*ShortLong1Code)(&value)
}

// Posting of an item to a cash account, in the context of a cash transaction, that results in an increase or decrease to the balance of the account.
type AmountAndDirection32 struct {

	// Indicates whether the net proceeds include interest accrued on the financial instrument.
	AccruedInterestIndicator *YesNoIndicator `xml:"AcrdIntrstInd,omitempty"`

	// Whether the net proceeds include stamp duty amount.
	StampDutyIndicator *YesNoIndicator `xml:"StmpDtyInd,omitempty"`

	// Amount of money in the cash entry.
	Amount *ActiveCurrencyAndAmount `xml:"Amt,omitempty"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalCurrencyAndOrderedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlCcyAndOrdrdAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms11 `xml:"FXDtls,omitempty"`

	// Date and time at which the cash is at the disposal of the credit account owner, or ceases to be at the disposal of the debit account owner.
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`
}

func (a *AmountAndDirection32) SetAccruedInterestIndicator(value string) {
	a.AccruedInterestIndicator = (*YesNoIndicator)(&value)
}

func (a *AmountAndDirection32) SetStampDutyIndicator(value string) {
	a.StampDutyIndicator = (*YesNoIndicator)(&value)
}

func (a *AmountAndDirection32) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection32) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection32) SetOriginalCurrencyAndOrderedAmount(value, currency string) {
	a.OriginalCurrencyAndOrderedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection32) AddForeignExchangeDetails() *ForeignExchangeTerms11 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms11)
	return a.ForeignExchangeDetails
}

func (a *AmountAndDirection32) AddValueDate() *DateAndDateTimeChoice {
	a.ValueDate = new(DateAndDateTimeChoice)
	return a.ValueDate
}

// Amount of money expressed with an optional currency code and debit/credit indicator.
type AmountAndDirection34 struct {

	// Amount of money that results in an increase (positively signed) or decrease (negatively signed), with specification of the currency.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates that the amount value is positive or negative.
	Sign *PlusOrMinusIndicator `xml:"Sgn"`
}

func (a *AmountAndDirection34) SetAmount(value, currency string) {
	a.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection34) SetSign(value string) {
	a.Sign = (*PlusOrMinusIndicator)(&value)
}

// Resulting debit or credit amount of the netted amounts for all debit and credit entries.
type AmountAndDirection35 struct {

	// Resulting amount of the netted amounts for all debit and credit entries.
	Amount *NonNegativeDecimalNumber `xml:"Amt"`

	// Indicates whether the amount is a credit or a debit amount.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`
}

func (a *AmountAndDirection35) SetAmount(value string) {
	a.Amount = (*NonNegativeDecimalNumber)(&value)
}

func (a *AmountAndDirection35) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

// Posting of an item to a cash account, in the context of a cash transaction, that results in an increase or decrease to the balance of the account.
type AmountAndDirection36 struct {

	// Indicates whether the net proceeds include interest accrued on the financial instrument.
	AccruedInterestIndicator *YesNoIndicator `xml:"AcrdIntrstInd,omitempty"`

	// Whether the net proceeds include stamp duty amount.
	StampDutyIndicator *YesNoIndicator `xml:"StmpDtyInd,omitempty"`

	// Indicates whether the net proceeds include brokerage fees for the transaction. If absent, element is not required.
	BrokerageAmountIndicator *YesNoIndicator `xml:"BrkrgAmtInd,omitempty"`

	// Amount of money in the cash entry.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalCurrencyAndOrderedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlCcyAndOrdrdAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms11 `xml:"FXDtls,omitempty"`

	// Date and time at which the cash is at the disposal of the credit account owner, or ceases to be at the disposal of the debit account owner.
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`
}

func (a *AmountAndDirection36) SetAccruedInterestIndicator(value string) {
	a.AccruedInterestIndicator = (*YesNoIndicator)(&value)
}

func (a *AmountAndDirection36) SetStampDutyIndicator(value string) {
	a.StampDutyIndicator = (*YesNoIndicator)(&value)
}

func (a *AmountAndDirection36) SetBrokerageAmountIndicator(value string) {
	a.BrokerageAmountIndicator = (*YesNoIndicator)(&value)
}

func (a *AmountAndDirection36) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection36) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection36) SetOriginalCurrencyAndOrderedAmount(value, currency string) {
	a.OriginalCurrencyAndOrderedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection36) AddForeignExchangeDetails() *ForeignExchangeTerms11 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms11)
	return a.ForeignExchangeDetails
}

func (a *AmountAndDirection36) AddValueDate() *DateAndDateTimeChoice {
	a.ValueDate = new(DateAndDateTimeChoice)
	return a.ValueDate
}

// Posting of an item to a cash account, in the context of a cash transaction, that results in an increase or decrease to the balance of the account.
type AmountAndDirection37 struct {

	// Indicates whether the net proceeds include interest accrued on the financial instrument.
	AccruedInterestIndicator *YesNoIndicator `xml:"AcrdIntrstInd,omitempty"`

	// Whether the net proceeds include stamp duty amount.
	StampDutyIndicator *YesNoIndicator `xml:"StmpDtyInd,omitempty"`

	// Indicates whether the net proceeds include brokerage fees for the transaction. If absent, element is not required.
	BrokerageAmountIndicator *YesNoIndicator `xml:"BrkrgAmtInd,omitempty"`

	// Amount of money in the cash entry.
	Amount *ActiveCurrencyAndAmount `xml:"Amt,omitempty"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalCurrencyAndOrderedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlCcyAndOrdrdAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms11 `xml:"FXDtls,omitempty"`

	// Date and time at which the cash is at the disposal of the credit account owner, or ceases to be at the disposal of the debit account owner.
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`
}

func (a *AmountAndDirection37) SetAccruedInterestIndicator(value string) {
	a.AccruedInterestIndicator = (*YesNoIndicator)(&value)
}

func (a *AmountAndDirection37) SetStampDutyIndicator(value string) {
	a.StampDutyIndicator = (*YesNoIndicator)(&value)
}

func (a *AmountAndDirection37) SetBrokerageAmountIndicator(value string) {
	a.BrokerageAmountIndicator = (*YesNoIndicator)(&value)
}

func (a *AmountAndDirection37) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection37) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection37) SetOriginalCurrencyAndOrderedAmount(value, currency string) {
	a.OriginalCurrencyAndOrderedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection37) AddForeignExchangeDetails() *ForeignExchangeTerms11 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms11)
	return a.ForeignExchangeDetails
}

func (a *AmountAndDirection37) AddValueDate() *DateAndDateTimeChoice {
	a.ValueDate = new(DateAndDateTimeChoice)
	return a.ValueDate
}

// Amount of money debited or credited on the books of an account servicer.
type AmountAndDirection4 struct {

	// Amount of money in the cash entry.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`
}

func (a *AmountAndDirection4) SetAmount(value, currency string) {
	a.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection4) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

// Signed amount.
type AmountAndDirection41 struct {

	// Amount value.
	Amount *CurrencyAndAmount `xml:"Amt"`

	// Indicates that the amount value is positive or negative.
	Sign *PlusOrMinusIndicator `xml:"Sgn,omitempty"`
}

func (a *AmountAndDirection41) SetAmount(value, currency string) {
	a.Amount = NewCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection41) SetSign(value string) {
	a.Sign = (*PlusOrMinusIndicator)(&value)
}

// Signed amount.
type AmountAndDirection43 struct {

	// Amount value.
	Amount *CurrencyAndAmount `xml:"Amt"`

	// Indicates that the amount value is positive or negative.
	Sign *PlusOrMinusIndicator `xml:"Sgn,omitempty"`

	// Date of the amount.
	Date *ISODate `xml:"Dt,omitempty"`
}

func (a *AmountAndDirection43) SetAmount(value, currency string) {
	a.Amount = NewCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection43) SetSign(value string) {
	a.Sign = (*PlusOrMinusIndicator)(&value)
}

func (a *AmountAndDirection43) SetDate(value string) {
	a.Date = (*ISODate)(&value)
}

// Amount of money debited or credited on the books of an account servicer.
type AmountAndDirection44 struct {

	// Amount of money in the cash entry.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalCurrencyAndOrderedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlCcyAndOrdrdAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms23 `xml:"FXDtls,omitempty"`
}

func (a *AmountAndDirection44) SetAmount(value, currency string) {
	a.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection44) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection44) SetOriginalCurrencyAndOrderedAmount(value, currency string) {
	a.OriginalCurrencyAndOrderedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection44) AddForeignExchangeDetails() *ForeignExchangeTerms23 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms23)
	return a.ForeignExchangeDetails
}

// Posting of an item to a cash account, in the context of a cash transaction, that results in an increase or decrease to the balance of the account.
type AmountAndDirection45 struct {

	// Indicates whether the net proceeds include interest accrued on the financial instrument.
	AccruedInterestIndicator *YesNoIndicator `xml:"AcrdIntrstInd,omitempty"`

	// Whether the net proceeds include stamp duty amount.
	StampDutyIndicator *YesNoIndicator `xml:"StmpDtyInd,omitempty"`

	// Indicates whether the net proceeds include brokerage fees for the transaction. If absent, element is not required.
	BrokerageAmountIndicator *YesNoIndicator `xml:"BrkrgAmtInd,omitempty"`

	// Amount of money in the cash entry.
	Amount *ActiveCurrencyAndAmount `xml:"Amt,omitempty"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalCurrencyAndOrderedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlCcyAndOrdrdAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms23 `xml:"FXDtls,omitempty"`

	// Date and time at which the cash is at the disposal of the credit account owner, or ceases to be at the disposal of the debit account owner.
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`
}

func (a *AmountAndDirection45) SetAccruedInterestIndicator(value string) {
	a.AccruedInterestIndicator = (*YesNoIndicator)(&value)
}

func (a *AmountAndDirection45) SetStampDutyIndicator(value string) {
	a.StampDutyIndicator = (*YesNoIndicator)(&value)
}

func (a *AmountAndDirection45) SetBrokerageAmountIndicator(value string) {
	a.BrokerageAmountIndicator = (*YesNoIndicator)(&value)
}

func (a *AmountAndDirection45) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection45) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection45) SetOriginalCurrencyAndOrderedAmount(value, currency string) {
	a.OriginalCurrencyAndOrderedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection45) AddForeignExchangeDetails() *ForeignExchangeTerms23 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms23)
	return a.ForeignExchangeDetails
}

func (a *AmountAndDirection45) AddValueDate() *DateAndDateTimeChoice {
	a.ValueDate = new(DateAndDateTimeChoice)
	return a.ValueDate
}

// Posting of an item to a cash account, in the context of a cash transaction, that results in an increase or decrease to the balance of the account.
type AmountAndDirection46 struct {

	// Indicates whether the net proceeds include interest accrued on the financial instrument.
	AccruedInterestIndicator *YesNoIndicator `xml:"AcrdIntrstInd,omitempty"`

	// Whether the net proceeds include stamp duty amount.
	StampDutyIndicator *YesNoIndicator `xml:"StmpDtyInd,omitempty"`

	// Indicates whether the net proceeds include brokerage fees for the transaction. If absent, element is not required.
	BrokerageAmountIndicator *YesNoIndicator `xml:"BrkrgAmtInd,omitempty"`

	// Amount of money in the cash entry.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalCurrencyAndOrderedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlCcyAndOrdrdAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms23 `xml:"FXDtls,omitempty"`

	// Date and time at which the cash is at the disposal of the credit account owner, or ceases to be at the disposal of the debit account owner.
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`
}

func (a *AmountAndDirection46) SetAccruedInterestIndicator(value string) {
	a.AccruedInterestIndicator = (*YesNoIndicator)(&value)
}

func (a *AmountAndDirection46) SetStampDutyIndicator(value string) {
	a.StampDutyIndicator = (*YesNoIndicator)(&value)
}

func (a *AmountAndDirection46) SetBrokerageAmountIndicator(value string) {
	a.BrokerageAmountIndicator = (*YesNoIndicator)(&value)
}

func (a *AmountAndDirection46) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection46) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection46) SetOriginalCurrencyAndOrderedAmount(value, currency string) {
	a.OriginalCurrencyAndOrderedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection46) AddForeignExchangeDetails() *ForeignExchangeTerms23 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms23)
	return a.ForeignExchangeDetails
}

func (a *AmountAndDirection46) AddValueDate() *DateAndDateTimeChoice {
	a.ValueDate = new(DateAndDateTimeChoice)
	return a.ValueDate
}

// Amount of money debited or credited on the books of an account servicer.
type AmountAndDirection47 struct {

	// Amount of money in the cash entry.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms23 `xml:"FXDtls,omitempty"`
}

func (a *AmountAndDirection47) SetAmount(value, currency string) {
	a.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection47) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection47) AddForeignExchangeDetails() *ForeignExchangeTerms23 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms23)
	return a.ForeignExchangeDetails
}

// Posting of an item to a cash account, in the context of a cash transaction, that results in an increase or decrease to the balance of the account.
type AmountAndDirection48 struct {

	// Amount of money in the cash entry.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalCurrencyAndOrderedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlCcyAndOrdrdAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms23 `xml:"FXDtls,omitempty"`

	// Date and time at which the cash is at the disposal of the credit account owner, or ceases to be at the disposal of the debit account owner.
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`
}

func (a *AmountAndDirection48) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection48) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection48) SetOriginalCurrencyAndOrderedAmount(value, currency string) {
	a.OriginalCurrencyAndOrderedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection48) AddForeignExchangeDetails() *ForeignExchangeTerms23 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms23)
	return a.ForeignExchangeDetails
}

func (a *AmountAndDirection48) AddValueDate() *DateAndDateTimeChoice {
	a.ValueDate = new(DateAndDateTimeChoice)
	return a.ValueDate
}

// Amount of money debited or credited on the books of an account servicer.
type AmountAndDirection49 struct {

	// Amount of money in the cash entry.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalCurrencyAndOrderedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlCcyAndOrdrdAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms23 `xml:"FXDtls,omitempty"`
}

func (a *AmountAndDirection49) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection49) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection49) SetOriginalCurrencyAndOrderedAmount(value, currency string) {
	a.OriginalCurrencyAndOrderedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection49) AddForeignExchangeDetails() *ForeignExchangeTerms23 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms23)
	return a.ForeignExchangeDetails
}

// Amount of money debited or credited on the books of an account servicer.
type AmountAndDirection5 struct {

	// Amount of money that is debited or credited.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Indicates if the amount is a debited or a credited.
	CreditDebit *CreditDebitCode `xml:"CdtDbt,omitempty"`
}

func (a *AmountAndDirection5) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection5) SetCreditDebit(value string) {
	a.CreditDebit = (*CreditDebitCode)(&value)
}

// Amount of money debited or credited on the books of an account servicer.
type AmountAndDirection51 struct {

	// Amount of money in the cash entry.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalCurrencyAndOrderedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlCcyAndOrdrdAmt,omitempty"`
}

func (a *AmountAndDirection51) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection51) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection51) SetOriginalCurrencyAndOrderedAmount(value, currency string) {
	a.OriginalCurrencyAndOrderedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

// Posting of an item to a cash account, in the context of a cash transaction, that results in an increase or decrease to the balance of the account.
type AmountAndDirection52 struct {

	// Amount of money in the cash entry.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`
}

func (a *AmountAndDirection52) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection52) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

// Amount of money debited or credited on the books of an account servicer.
type AmountAndDirection55 struct {

	// Amount of money in the cash entry.
	Amount *RestrictedFINActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalCurrencyAndOrderedAmount *RestrictedFINActiveOrHistoricCurrencyAndAmount `xml:"OrgnlCcyAndOrdrdAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms23 `xml:"FXDtls,omitempty"`
}

func (a *AmountAndDirection55) SetAmount(value, currency string) {
	a.Amount = NewRestrictedFINActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection55) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection55) SetOriginalCurrencyAndOrderedAmount(value, currency string) {
	a.OriginalCurrencyAndOrderedAmount = NewRestrictedFINActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection55) AddForeignExchangeDetails() *ForeignExchangeTerms23 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms23)
	return a.ForeignExchangeDetails
}

// Posting of an item to a cash account, in the context of a cash transaction, that results in an increase or decrease to the balance of the account.
type AmountAndDirection57 struct {

	// Amount of money in the cash entry.
	Amount *RestrictedFINActiveCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`
}

func (a *AmountAndDirection57) SetAmount(value, currency string) {
	a.Amount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection57) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

// Amount of money debited or credited on the books of an account servicer.
type AmountAndDirection58 struct {

	// Amount of money in the cash entry.
	Amount *RestrictedFINActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalCurrencyAndOrderedAmount *RestrictedFINActiveOrHistoricCurrencyAndAmount `xml:"OrgnlCcyAndOrdrdAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms27 `xml:"FXDtls,omitempty"`
}

func (a *AmountAndDirection58) SetAmount(value, currency string) {
	a.Amount = NewRestrictedFINActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection58) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection58) SetOriginalCurrencyAndOrderedAmount(value, currency string) {
	a.OriginalCurrencyAndOrderedAmount = NewRestrictedFINActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection58) AddForeignExchangeDetails() *ForeignExchangeTerms27 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms27)
	return a.ForeignExchangeDetails
}

// Amount of money debited or credited on the books of an account servicer.
type AmountAndDirection59 struct {

	// Amount of money in the cash entry.
	Amount *RestrictedFINActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`
}

func (a *AmountAndDirection59) SetAmount(value, currency string) {
	a.Amount = NewRestrictedFINActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection59) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

// Posting of an item to a cash account, in the context of a cash transaction, that results in an increase or decrease to the balance of the account.
type AmountAndDirection6 struct {

	// Amount of money in the cash entry.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates that the amount value is positive or negative.
	Sign *PlusOrMinusIndicator `xml:"Sgn"`
}

func (a *AmountAndDirection6) SetAmount(value, currency string) {
	a.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection6) SetSign(value string) {
	a.Sign = (*PlusOrMinusIndicator)(&value)
}

// Posting of an item to a cash account, in the context of a cash transaction, that results in an increase or decrease to the balance of the account.
type AmountAndDirection60 struct {

	// Indicates whether the net proceeds include interest accrued on the financial instrument.
	AccruedInterestIndicator *YesNoIndicator `xml:"AcrdIntrstInd,omitempty"`

	// Whether the net proceeds include stamp duty amount.
	StampDutyIndicator *YesNoIndicator `xml:"StmpDtyInd,omitempty"`

	// Indicates whether the net proceeds include brokerage fees for the transaction. If absent, element is not required.
	BrokerageAmountIndicator *YesNoIndicator `xml:"BrkrgAmtInd,omitempty"`

	// Amount of money in the cash entry.
	Amount *RestrictedFINActiveCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalCurrencyAndOrderedAmount *RestrictedFINActiveOrHistoricCurrencyAndAmount `xml:"OrgnlCcyAndOrdrdAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms27 `xml:"FXDtls,omitempty"`

	// Date and time at which the cash is at the disposal of the credit account owner, or ceases to be at the disposal of the debit account owner.
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`
}

func (a *AmountAndDirection60) SetAccruedInterestIndicator(value string) {
	a.AccruedInterestIndicator = (*YesNoIndicator)(&value)
}

func (a *AmountAndDirection60) SetStampDutyIndicator(value string) {
	a.StampDutyIndicator = (*YesNoIndicator)(&value)
}

func (a *AmountAndDirection60) SetBrokerageAmountIndicator(value string) {
	a.BrokerageAmountIndicator = (*YesNoIndicator)(&value)
}

func (a *AmountAndDirection60) SetAmount(value, currency string) {
	a.Amount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection60) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection60) SetOriginalCurrencyAndOrderedAmount(value, currency string) {
	a.OriginalCurrencyAndOrderedAmount = NewRestrictedFINActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection60) AddForeignExchangeDetails() *ForeignExchangeTerms27 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms27)
	return a.ForeignExchangeDetails
}

func (a *AmountAndDirection60) AddValueDate() *DateAndDateTimeChoice {
	a.ValueDate = new(DateAndDateTimeChoice)
	return a.ValueDate
}

// Amount of money debited or credited on the books of an account servicer.
type AmountAndDirection66 struct {

	// Amount of money in the cash entry.
	Amount *RestrictedFINActiveCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalCurrencyAndOrderedAmount *RestrictedFINActiveOrHistoricCurrencyAndAmount `xml:"OrgnlCcyAndOrdrdAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms27 `xml:"FXDtls,omitempty"`
}

func (a *AmountAndDirection66) SetAmount(value, currency string) {
	a.Amount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection66) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection66) SetOriginalCurrencyAndOrderedAmount(value, currency string) {
	a.OriginalCurrencyAndOrderedAmount = NewRestrictedFINActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection66) AddForeignExchangeDetails() *ForeignExchangeTerms27 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms27)
	return a.ForeignExchangeDetails
}

// Amount of money debited or credited on the books of an account servicer.
type AmountAndDirection67 struct {

	// Amount of money in the cash entry.
	Amount *RestrictedFINActiveCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalCurrencyAndOrderedAmount *RestrictedFINActiveOrHistoricCurrencyAndAmount `xml:"OrgnlCcyAndOrdrdAmt,omitempty"`
}

func (a *AmountAndDirection67) SetAmount(value, currency string) {
	a.Amount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection67) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection67) SetOriginalCurrencyAndOrderedAmount(value, currency string) {
	a.OriginalCurrencyAndOrderedAmount = NewRestrictedFINActiveOrHistoricCurrencyAndAmount(value, currency)
}

// Posting of an item to a cash account, in the context of a cash transaction, that results in an increase or decrease to the balance of the account.
type AmountAndDirection7 struct {

	// Amount of money in the cash entry.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`
}

func (a *AmountAndDirection7) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection7) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

// Posting of an item to a cash account, in the context of a cash transaction, that results in an increase or decrease to the balance of the account.
type AmountAndDirection71 struct {

	// Amount of money in the cash entry.
	Amount *RestrictedFINActiveCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalCurrencyAndOrderedAmount *RestrictedFINActiveOrHistoricCurrencyAndAmount `xml:"OrgnlCcyAndOrdrdAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms27 `xml:"FXDtls,omitempty"`

	// Date and time at which the cash is at the disposal of the credit account owner, or ceases to be at the disposal of the debit account owner.
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`
}

func (a *AmountAndDirection71) SetAmount(value, currency string) {
	a.Amount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection71) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection71) SetOriginalCurrencyAndOrderedAmount(value, currency string) {
	a.OriginalCurrencyAndOrderedAmount = NewRestrictedFINActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection71) AddForeignExchangeDetails() *ForeignExchangeTerms27 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms27)
	return a.ForeignExchangeDetails
}

func (a *AmountAndDirection71) AddValueDate() *DateAndDateTimeChoice {
	a.ValueDate = new(DateAndDateTimeChoice)
	return a.ValueDate
}

// Amount of money debited or credited on the books of an account servicer.
type AmountAndDirection72 struct {

	// Amount of money in the cash entry.
	Amount *RestrictedFINActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms27 `xml:"FXDtls,omitempty"`
}

func (a *AmountAndDirection72) SetAmount(value, currency string) {
	a.Amount = NewRestrictedFINActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection72) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection72) AddForeignExchangeDetails() *ForeignExchangeTerms27 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms27)
	return a.ForeignExchangeDetails
}

// Amount of money debited or credited on the books of an account servicer.
type AmountAndDirection8 struct {

	// Amount of money in the cash entry.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalCurrencyAndOrderedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlCcyAndOrdrdAmt,omitempty"`
}

func (a *AmountAndDirection8) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection8) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection8) SetOriginalCurrencyAndOrderedAmount(value, currency string) {
	a.OriginalCurrencyAndOrderedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

// Posting of an item to a cash account, in the context of a cash transaction, that results in an increase or decrease to the balance of the account.
type AmountAndDirection85 struct {

	// Indicates whether the net proceeds include interest accrued on the financial instrument.
	AccruedInterestIndicator *YesNoIndicator `xml:"AcrdIntrstInd,omitempty"`

	// Whether the net proceeds include stamp duty amount.
	StampDutyIndicator *YesNoIndicator `xml:"StmpDtyInd,omitempty"`

	// Indicates whether the net proceeds include brokerage fees for the transaction. If absent, element is not required.
	BrokerageAmountIndicator *YesNoIndicator `xml:"BrkrgAmtInd,omitempty"`

	// Amount of money in the cash entry.
	Amount *RestrictedFINActiveCurrencyAndAmount `xml:"Amt,omitempty"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalCurrencyAndOrderedAmount *RestrictedFINActiveOrHistoricCurrencyAndAmount `xml:"OrgnlCcyAndOrdrdAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms27 `xml:"FXDtls,omitempty"`

	// Date and time at which the cash is at the disposal of the credit account owner, or ceases to be at the disposal of the debit account owner.
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`
}

func (a *AmountAndDirection85) SetAccruedInterestIndicator(value string) {
	a.AccruedInterestIndicator = (*YesNoIndicator)(&value)
}

func (a *AmountAndDirection85) SetStampDutyIndicator(value string) {
	a.StampDutyIndicator = (*YesNoIndicator)(&value)
}

func (a *AmountAndDirection85) SetBrokerageAmountIndicator(value string) {
	a.BrokerageAmountIndicator = (*YesNoIndicator)(&value)
}

func (a *AmountAndDirection85) SetAmount(value, currency string) {
	a.Amount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection85) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection85) SetOriginalCurrencyAndOrderedAmount(value, currency string) {
	a.OriginalCurrencyAndOrderedAmount = NewRestrictedFINActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection85) AddForeignExchangeDetails() *ForeignExchangeTerms27 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms27)
	return a.ForeignExchangeDetails
}

func (a *AmountAndDirection85) AddValueDate() *DateAndDateTimeChoice {
	a.ValueDate = new(DateAndDateTimeChoice)
	return a.ValueDate
}

// Amount of money debited or credited on the books of an account servicer.
type AmountAndDirection9 struct {

	// Amount of money in the cash entry.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalCurrencyAndOrderedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlCcyAndOrdrdAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms11 `xml:"FXDtls,omitempty"`
}

func (a *AmountAndDirection9) SetAmount(value, currency string) {
	a.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection9) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection9) SetOriginalCurrencyAndOrderedAmount(value, currency string) {
	a.OriginalCurrencyAndOrderedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection9) AddForeignExchangeDetails() *ForeignExchangeTerms11 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms11)
	return a.ForeignExchangeDetails
}
