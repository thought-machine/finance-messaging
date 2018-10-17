package iso20022

// Limit of amounts for the customer.
type ATMTransactionAmounts2 struct {

	// Currency of the limits, if different from the requested amount.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Maximum amount allowed in the authorised currency if the withdrawal was not approved.
	MaximumAuthorisableAmount *ImpliedCurrencyAndAmount `xml:"MaxAuthsbAmt,omitempty"`

	// Minimum amount allowed for a withdrawal in the authorised currency.
	MinimumAllowedAmount *ImpliedCurrencyAndAmount `xml:"MinAllwdAmt,omitempty"`

	// Maximum amount allowed for a withdrawal in the authorised currency.
	MaximumAllowedAmount *ImpliedCurrencyAndAmount `xml:"MaxAllwdAmt,omitempty"`

	// Remaining daily amount of the customer totals after the withdrawal.
	DailyBalance *DetailedAmount4 `xml:"DalyBal,omitempty"`

	// Remaining weekly amount of the customer totals after the withdrawal.
	WeeklyBalance *DetailedAmount4 `xml:"WklyBal,omitempty"`

	// Remaining monthly amount of the customer totals after the withdrawal.
	MonthlyBalance *DetailedAmount4 `xml:"MnthlyBal,omitempty"`
}

func (a *ATMTransactionAmounts2) SetCurrency(value string) {
	a.Currency = (*ActiveCurrencyCode)(&value)
}

func (a *ATMTransactionAmounts2) SetMaximumAuthorisableAmount(value, currency string) {
	a.MaximumAuthorisableAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransactionAmounts2) SetMinimumAllowedAmount(value, currency string) {
	a.MinimumAllowedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransactionAmounts2) SetMaximumAllowedAmount(value, currency string) {
	a.MaximumAllowedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransactionAmounts2) AddDailyBalance() *DetailedAmount4 {
	a.DailyBalance = new(DetailedAmount4)
	return a.DailyBalance
}

func (a *ATMTransactionAmounts2) AddWeeklyBalance() *DetailedAmount4 {
	a.WeeklyBalance = new(DetailedAmount4)
	return a.WeeklyBalance
}

func (a *ATMTransactionAmounts2) AddMonthlyBalance() *DetailedAmount4 {
	a.MonthlyBalance = new(DetailedAmount4)
	return a.MonthlyBalance
}

// Limit of amounts for the customer.
type ATMTransactionAmounts3 struct {

	// Type of limit.
	Type *Max35Text `xml:"Tp"`

	// Label of the limit to display or print.
	Label *Max35Text `xml:"Labl,omitempty"`

	// Currency of the limit amount.
	Currency *ActiveCurrencyCode `xml:"Ccy"`

	// Minimum amount value in the currency of the limit.
	MinimumAmount *ImpliedCurrencyAndAmount `xml:"MinAmt,omitempty"`

	// Maximum amount value in the currency of the limit.
	MaximumAmount *ImpliedCurrencyAndAmount `xml:"MaxAmt,omitempty"`
}

func (a *ATMTransactionAmounts3) SetType(value string) {
	a.Type = (*Max35Text)(&value)
}

func (a *ATMTransactionAmounts3) SetLabel(value string) {
	a.Label = (*Max35Text)(&value)
}

func (a *ATMTransactionAmounts3) SetCurrency(value string) {
	a.Currency = (*ActiveCurrencyCode)(&value)
}

func (a *ATMTransactionAmounts3) SetMinimumAmount(value, currency string) {
	a.MinimumAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransactionAmounts3) SetMaximumAmount(value, currency string) {
	a.MaximumAmount = NewImpliedCurrencyAndAmount(value, currency)
}

// Withdrawal limits for the account.
type ATMTransactionAmounts4 struct {

	// True if limits may be displayed on the ATM to the customer.
	DisplayFlag *TrueFalseIndicator `xml:"DispFlg,omitempty"`

	// Amount available for withdrawal on the account.
	AvailableAmount *ImpliedCurrencyAndAmount `xml:"AvlblAmt,omitempty"`

	// Remaining daily amount of the customer totals for withdrawals on the account.
	DailyBalance *DetailedAmount4 `xml:"DalyBal,omitempty"`

	// Remaining weekly amount of the customer totals for withdrawals on the account.
	WeeklyBalance *DetailedAmount4 `xml:"WklyBal,omitempty"`

	// Remaining monthly amount of the customer totals for withdrawals on the account.
	MonthlyBalance *DetailedAmount4 `xml:"MnthlyBal,omitempty"`
}

func (a *ATMTransactionAmounts4) SetDisplayFlag(value string) {
	a.DisplayFlag = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransactionAmounts4) SetAvailableAmount(value, currency string) {
	a.AvailableAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransactionAmounts4) AddDailyBalance() *DetailedAmount4 {
	a.DailyBalance = new(DetailedAmount4)
	return a.DailyBalance
}

func (a *ATMTransactionAmounts4) AddWeeklyBalance() *DetailedAmount4 {
	a.WeeklyBalance = new(DetailedAmount4)
	return a.WeeklyBalance
}

func (a *ATMTransactionAmounts4) AddMonthlyBalance() *DetailedAmount4 {
	a.MonthlyBalance = new(DetailedAmount4)
	return a.MonthlyBalance
}

// Deposit limits for the account.
type ATMTransactionAmounts5 struct {

	// True if limits may be displayed on the ATM to the customer.
	DisplayFlag *TrueFalseIndicator `xml:"DispFlg,omitempty"`

	// Maximum amount allowed for deposit on the account.
	MaximumAmount *ImpliedCurrencyAndAmount `xml:"MaxAmt,omitempty"`
}

func (a *ATMTransactionAmounts5) SetDisplayFlag(value string) {
	a.DisplayFlag = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransactionAmounts5) SetMaximumAmount(value, currency string) {
	a.MaximumAmount = NewImpliedCurrencyAndAmount(value, currency)
}

// Limit of amounts for the customer.
type ATMTransactionAmounts6 struct {

	// Default currency of the limits.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Maximum amount allowed for a transaction in the service.
	MaximumPossibleAmount *ImpliedCurrencyAndAmount `xml:"MaxPssblAmt,omitempty"`

	// Minimum amount allowed for a transaction in the service.
	MinimumPossibleAmount *ImpliedCurrencyAndAmount `xml:"MinPssblAmt,omitempty"`

	// Additional amount that may be displayed to the customer, for instance the daily limit or the daily balance for the service.
	AdditionalAmount []*ATMTransactionAmounts7 `xml:"AddtlAmt,omitempty"`
}

func (a *ATMTransactionAmounts6) SetCurrency(value string) {
	a.Currency = (*ActiveCurrencyCode)(&value)
}

func (a *ATMTransactionAmounts6) SetMaximumPossibleAmount(value, currency string) {
	a.MaximumPossibleAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransactionAmounts6) SetMinimumPossibleAmount(value, currency string) {
	a.MinimumPossibleAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransactionAmounts6) AddAdditionalAmount() *ATMTransactionAmounts7 {
	newValue := new(ATMTransactionAmounts7)
	a.AdditionalAmount = append(a.AdditionalAmount, newValue)
	return newValue
}

// Additional amount that may be displayed to the customer, for instance the daily limit or the daily balance for the service.
type ATMTransactionAmounts7 struct {

	// Type of amount.
	Type *Max35Text `xml:"Tp"`

	// Amount value.
	Amount *ImpliedCurrencyAndAmount `xml:"Amt"`

	// Currency of the amount.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Description of the amount that may be provided to the customer.
	Label *Max70Text `xml:"Labl,omitempty"`
}

func (a *ATMTransactionAmounts7) SetType(value string) {
	a.Type = (*Max35Text)(&value)
}

func (a *ATMTransactionAmounts7) SetAmount(value, currency string) {
	a.Amount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransactionAmounts7) SetCurrency(value string) {
	a.Currency = (*ActiveCurrencyCode)(&value)
}

func (a *ATMTransactionAmounts7) SetLabel(value string) {
	a.Label = (*Max70Text)(&value)
}

// Limit of amounts for the customer.
type ATMTransactionAmounts8 struct {

	// Default currency of the limits.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Maximum amount allowed for a transaction in the service.
	MaximumPossibleAmount *ImpliedCurrencyAndAmount `xml:"MaxPssblAmt,omitempty"`

	// Minimum amount allowed for a transaction in the service.
	MinimumPossibleAmount *ImpliedCurrencyAndAmount `xml:"MinPssblAmt,omitempty"`

	// Additional amount that may be displayed to the customer, for instance the daily limit or the daily balance for the service.
	AdditionalAmount []*ATMTransactionAmounts7 `xml:"AddtlAmt,omitempty"`

	// Limit of deposited media for the customer.
	DepositLimits []*ATMTransactionAmounts9 `xml:"DpstLmts,omitempty"`
}

func (a *ATMTransactionAmounts8) SetCurrency(value string) {
	a.Currency = (*ActiveCurrencyCode)(&value)
}

func (a *ATMTransactionAmounts8) SetMaximumPossibleAmount(value, currency string) {
	a.MaximumPossibleAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransactionAmounts8) SetMinimumPossibleAmount(value, currency string) {
	a.MinimumPossibleAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransactionAmounts8) AddAdditionalAmount() *ATMTransactionAmounts7 {
	newValue := new(ATMTransactionAmounts7)
	a.AdditionalAmount = append(a.AdditionalAmount, newValue)
	return newValue
}

func (a *ATMTransactionAmounts8) AddDepositLimits() *ATMTransactionAmounts9 {
	newValue := new(ATMTransactionAmounts9)
	a.DepositLimits = append(a.DepositLimits, newValue)
	return newValue
}

// Limit of deposited media for the customer.
type ATMTransactionAmounts9 struct {

	// Type of media.
	MediaType *ATMMediaType2Code `xml:"MdiaTp"`

	// Currency of the media.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Minimum number of media.
	MinimumNumber *Number `xml:"MinNb,omitempty"`

	// Maximum number of media.
	MaximumNumber *Number `xml:"MaxNb,omitempty"`

	// True if limits may be displayed to the customer on the ATM.
	DisplayFlag *TrueFalseIndicator `xml:"DispFlg,omitempty"`
}

func (a *ATMTransactionAmounts9) SetMediaType(value string) {
	a.MediaType = (*ATMMediaType2Code)(&value)
}

func (a *ATMTransactionAmounts9) SetCurrency(value string) {
	a.Currency = (*ActiveCurrencyCode)(&value)
}

func (a *ATMTransactionAmounts9) SetMinimumNumber(value string) {
	a.MinimumNumber = (*Number)(&value)
}

func (a *ATMTransactionAmounts9) SetMaximumNumber(value string) {
	a.MaximumNumber = (*Number)(&value)
}

func (a *ATMTransactionAmounts9) SetDisplayFlag(value string) {
	a.DisplayFlag = (*TrueFalseIndicator)(&value)
}
