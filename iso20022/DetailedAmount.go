package iso20022

// Detailed amounts associated with the total amount of transaction.
type DetailedAmount1 struct {

	// Type of amount.
	Type *TypeOfAmount2Code `xml:"Tp"`

	// Amount value.
	Value *ImpliedCurrencyAndAmount `xml:"Val"`
}

func (d *DetailedAmount1) SetType(value string) {
	d.Type = (*TypeOfAmount2Code)(&value)
}

func (d *DetailedAmount1) SetValue(value, currency string) {
	d.Value = NewImpliedCurrencyAndAmount(value, currency)
}

// Additional amounts from the processor or the issuer without financial impacts on the transaction amount.
type DetailedAmount10 struct {

	// Type or class of amount.
	Type *TypeOfAmount6Code `xml:"Tp"`

	// Additional information to specify the type of amount.
	AdditionalType *Max35Text `xml:"AddtlTp,omitempty"`

	// Amount value.
	Amount *CurrencyAndAmount `xml:"Amt"`

	// Short description of the amount to provide to the cardholder.
	Label *Max35Text `xml:"Labl,omitempty"`
}

func (d *DetailedAmount10) SetType(value string) {
	d.Type = (*TypeOfAmount6Code)(&value)
}

func (d *DetailedAmount10) SetAdditionalType(value string) {
	d.AdditionalType = (*Max35Text)(&value)
}

func (d *DetailedAmount10) SetAmount(value, currency string) {
	d.Amount = NewCurrencyAndAmount(value, currency)
}

func (d *DetailedAmount10) SetLabel(value string) {
	d.Label = (*Max35Text)(&value)
}

// Fees between acquirer and issuer.
type DetailedAmount11 struct {

	// Type or class of amount.
	Type *TypeOfAmount7Code `xml:"Tp"`

	// Additional information to specify the type of amount.
	AdditionalType *Max35Text `xml:"AddtlTp,omitempty"`

	// Amount value.
	Amount *AmountAndDirection41 `xml:"Amt"`

	// Original value of the amount.
	OriginalAmount *AmountAndDirection41 `xml:"OrgnlAmt,omitempty"`
}

func (d *DetailedAmount11) SetType(value string) {
	d.Type = (*TypeOfAmount7Code)(&value)
}

func (d *DetailedAmount11) SetAdditionalType(value string) {
	d.AdditionalType = (*Max35Text)(&value)
}

func (d *DetailedAmount11) AddAmount() *AmountAndDirection41 {
	d.Amount = new(AmountAndDirection41)
	return d.Amount
}

func (d *DetailedAmount11) AddOriginalAmount() *AmountAndDirection41 {
	d.OriginalAmount = new(AmountAndDirection41)
	return d.OriginalAmount
}

// Amounts of the withdrawal transaction.
type DetailedAmount12 struct {

	// Amount to be dispensed by the ATM after the approval of the withdrawal transaction.
	AmountToDispense *ImpliedCurrencyAndAmount `xml:"AmtToDspns"`

	// Currency of the amount to dispense when different from the base currency of the ATM.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Withdrawal fees, accepted by the customer.
	Fees []*DetailedAmount13 `xml:"Fees,omitempty"`

	// Amount of the donation.
	Donation []*DetailedAmount13 `xml:"Dontn,omitempty"`
}

func (d *DetailedAmount12) SetAmountToDispense(value, currency string) {
	d.AmountToDispense = NewImpliedCurrencyAndAmount(value, currency)
}

func (d *DetailedAmount12) SetCurrency(value string) {
	d.Currency = (*ActiveCurrencyCode)(&value)
}

func (d *DetailedAmount12) AddFees() *DetailedAmount13 {
	newValue := new(DetailedAmount13)
	d.Fees = append(d.Fees, newValue)
	return newValue
}

func (d *DetailedAmount12) AddDonation() *DetailedAmount13 {
	newValue := new(DetailedAmount13)
	d.Donation = append(d.Donation, newValue)
	return newValue
}

// Withdrawal fees, accepted by the customer.
type DetailedAmount13 struct {

	// Amount value.
	Amount *ImpliedCurrencyAndAmount `xml:"Amt"`

	// Currency of the amount.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Short description of the amount to display or print.
	Label *Max70Text `xml:"Labl,omitempty"`
}

func (d *DetailedAmount13) SetAmount(value, currency string) {
	d.Amount = NewImpliedCurrencyAndAmount(value, currency)
}

func (d *DetailedAmount13) SetCurrency(value string) {
	d.Currency = (*ActiveCurrencyCode)(&value)
}

func (d *DetailedAmount13) SetLabel(value string) {
	d.Label = (*Max70Text)(&value)
}

// Detailed amounts associated with the total amount of transaction.
type DetailedAmount14 struct {

	// Amount value.
	Amount *ImpliedCurrencyAndAmount `xml:"Amt"`

	// Date and time of the payment.
	DateTime *ISODateTime `xml:"DtTm"`

	// Card data entry mode for the related payment.
	CardDataEntryMode *CardDataReading5Code `xml:"CardDataNtryMd,omitempty"`

	// Data of an integrated circuit card application for the related payment.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`

	// Short description of the amount to display or print.
	Label *Max140Text `xml:"Labl,omitempty"`
}

func (d *DetailedAmount14) SetAmount(value, currency string) {
	d.Amount = NewImpliedCurrencyAndAmount(value, currency)
}

func (d *DetailedAmount14) SetDateTime(value string) {
	d.DateTime = (*ISODateTime)(&value)
}

func (d *DetailedAmount14) SetCardDataEntryMode(value string) {
	d.CardDataEntryMode = (*CardDataReading5Code)(&value)
}

func (d *DetailedAmount14) SetICCRelatedData(value string) {
	d.ICCRelatedData = (*Max10000Binary)(&value)
}

func (d *DetailedAmount14) SetLabel(value string) {
	d.Label = (*Max140Text)(&value)
}

// Detailed amounts associated with the total amount of transaction.
type DetailedAmount15 struct {

	// Amount of purchase goods and services without tax.
	AmountGoodsAndServices *ImpliedCurrencyAndAmount `xml:"AmtGoodsAndSvcs,omitempty"`

	// Cash-back amount.
	CashBack *ImpliedCurrencyAndAmount `xml:"CshBck,omitempty"`

	// Gratuity amount.
	Gratuity *ImpliedCurrencyAndAmount `xml:"Grtty,omitempty"`

	// Fees amount.
	Fees []*DetailedAmount4 `xml:"Fees,omitempty"`

	// Global rebate of the transaction. This amount is counted as a negative amount.
	Rebate []*DetailedAmount4 `xml:"Rbt,omitempty"`

	// Value added tax amount.
	ValueAddedTax []*DetailedAmount4 `xml:"ValAddedTax,omitempty"`

	// Additional charge paid by the cardholder. For example airline credit card surcharge.
	Surcharge []*DetailedAmount4 `xml:"Srchrg,omitempty"`
}

func (d *DetailedAmount15) SetAmountGoodsAndServices(value, currency string) {
	d.AmountGoodsAndServices = NewImpliedCurrencyAndAmount(value, currency)
}

func (d *DetailedAmount15) SetCashBack(value, currency string) {
	d.CashBack = NewImpliedCurrencyAndAmount(value, currency)
}

func (d *DetailedAmount15) SetGratuity(value, currency string) {
	d.Gratuity = NewImpliedCurrencyAndAmount(value, currency)
}

func (d *DetailedAmount15) AddFees() *DetailedAmount4 {
	newValue := new(DetailedAmount4)
	d.Fees = append(d.Fees, newValue)
	return newValue
}

func (d *DetailedAmount15) AddRebate() *DetailedAmount4 {
	newValue := new(DetailedAmount4)
	d.Rebate = append(d.Rebate, newValue)
	return newValue
}

func (d *DetailedAmount15) AddValueAddedTax() *DetailedAmount4 {
	newValue := new(DetailedAmount4)
	d.ValueAddedTax = append(d.ValueAddedTax, newValue)
	return newValue
}

func (d *DetailedAmount15) AddSurcharge() *DetailedAmount4 {
	newValue := new(DetailedAmount4)
	d.Surcharge = append(d.Surcharge, newValue)
	return newValue
}

// Amounts of the deposit transaction.
type DetailedAmount16 struct {

	// Link to the account for multi-account deposit.
	AccountSequenceNumber *Number `xml:"AcctSeqNb,omitempty"`

	// Amount of the deposit to be made on the ATM after the approval of the deposit transaction.
	AmountToDeposit *ImpliedCurrencyAndAmount `xml:"AmtToDpst,omitempty"`

	// Currency of the amount to deposit when different from the base currency of the ATM.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Cashback amount value.
	CashBackAmount *ImpliedCurrencyAndAmount `xml:"CshBckAmt,omitempty"`

	// Deposit fees, accepted by the customer.
	Fees []*DetailedAmount13 `xml:"Fees,omitempty"`

	// Amount of the donation.
	Donation []*DetailedAmount13 `xml:"Dontn,omitempty"`
}

func (d *DetailedAmount16) SetAccountSequenceNumber(value string) {
	d.AccountSequenceNumber = (*Number)(&value)
}

func (d *DetailedAmount16) SetAmountToDeposit(value, currency string) {
	d.AmountToDeposit = NewImpliedCurrencyAndAmount(value, currency)
}

func (d *DetailedAmount16) SetCurrency(value string) {
	d.Currency = (*ActiveCurrencyCode)(&value)
}

func (d *DetailedAmount16) SetCashBackAmount(value, currency string) {
	d.CashBackAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (d *DetailedAmount16) AddFees() *DetailedAmount13 {
	newValue := new(DetailedAmount13)
	d.Fees = append(d.Fees, newValue)
	return newValue
}

func (d *DetailedAmount16) AddDonation() *DetailedAmount13 {
	newValue := new(DetailedAmount13)
	d.Donation = append(d.Donation, newValue)
	return newValue
}

// Details of the transfer transaction amounts.
type DetailedAmount17 struct {

	// Amount to be transferred from the source account to the destination account.
	AmountToTransfer *ImpliedCurrencyAndAmount `xml:"AmtToTrf"`

	// Currency of the amount to be transferred.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Transfer fees, accepted by the customer.
	Fees []*DetailedAmount18 `xml:"Fees,omitempty"`

	// Amount of the donation.
	Donation []*DetailedAmount18 `xml:"Dontn,omitempty"`
}

func (d *DetailedAmount17) SetAmountToTransfer(value, currency string) {
	d.AmountToTransfer = NewImpliedCurrencyAndAmount(value, currency)
}

func (d *DetailedAmount17) SetCurrency(value string) {
	d.Currency = (*ActiveCurrencyCode)(&value)
}

func (d *DetailedAmount17) AddFees() *DetailedAmount18 {
	newValue := new(DetailedAmount18)
	d.Fees = append(d.Fees, newValue)
	return newValue
}

func (d *DetailedAmount17) AddDonation() *DetailedAmount18 {
	newValue := new(DetailedAmount18)
	d.Donation = append(d.Donation, newValue)
	return newValue
}

// Withdrawal fees, accepted by the customer.
type DetailedAmount18 struct {

	// Amount value.
	Amount *ImpliedCurrencyAndAmount `xml:"Amt"`

	// Currency of the amount.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// True if amount charged to the source account.
	ChargeAccountTo *TrueFalseIndicator `xml:"ChrgAcctTo,omitempty"`

	// Short description of the amount to display or print.
	Label *Max70Text `xml:"Labl,omitempty"`
}

func (d *DetailedAmount18) SetAmount(value, currency string) {
	d.Amount = NewImpliedCurrencyAndAmount(value, currency)
}

func (d *DetailedAmount18) SetCurrency(value string) {
	d.Currency = (*ActiveCurrencyCode)(&value)
}

func (d *DetailedAmount18) SetChargeAccountTo(value string) {
	d.ChargeAccountTo = (*TrueFalseIndicator)(&value)
}

func (d *DetailedAmount18) SetLabel(value string) {
	d.Label = (*Max70Text)(&value)
}

// Detailed amounts associated with the total amount of transaction.
type DetailedAmount2 struct {

	// Type of amount.
	Type *TypeOfAmount2Code `xml:"Tp"`

	// Amount value.
	Value *ImpliedCurrencyAndAmount `xml:"Val"`

	// Short description of the amount.
	Label *Max35Text `xml:"Labl,omitempty"`
}

func (d *DetailedAmount2) SetType(value string) {
	d.Type = (*TypeOfAmount2Code)(&value)
}

func (d *DetailedAmount2) SetValue(value, currency string) {
	d.Value = NewImpliedCurrencyAndAmount(value, currency)
}

func (d *DetailedAmount2) SetLabel(value string) {
	d.Label = (*Max35Text)(&value)
}

// Detailed amounts associated with the total amount of transaction.
type DetailedAmount4 struct {

	// Amount value.
	Amount *ImpliedCurrencyAndAmount `xml:"Amt"`

	// Short description of the amount to display or print.
	Label *Max140Text `xml:"Labl,omitempty"`
}

func (d *DetailedAmount4) SetAmount(value, currency string) {
	d.Amount = NewImpliedCurrencyAndAmount(value, currency)
}

func (d *DetailedAmount4) SetLabel(value string) {
	d.Label = (*Max140Text)(&value)
}

// Detailed amounts associated with the total amount of transaction.
type DetailedAmount5 struct {

	// Cash-back amount.
	CashBack *ImpliedCurrencyAndAmount `xml:"CshBck,omitempty"`

	// Gratuity amount.
	Gratuity *ImpliedCurrencyAndAmount `xml:"Grtty,omitempty"`

	// Fees amount.
	Fees []*DetailedAmount4 `xml:"Fees,omitempty"`

	// Global rebate of the transaction. This amount is counted as a negative amount.
	Rebate []*DetailedAmount4 `xml:"Rbt,omitempty"`

	// Value added tax amount.
	ValueAddedTax []*DetailedAmount4 `xml:"ValAddedTax,omitempty"`
}

func (d *DetailedAmount5) SetCashBack(value, currency string) {
	d.CashBack = NewImpliedCurrencyAndAmount(value, currency)
}

func (d *DetailedAmount5) SetGratuity(value, currency string) {
	d.Gratuity = NewImpliedCurrencyAndAmount(value, currency)
}

func (d *DetailedAmount5) AddFees() *DetailedAmount4 {
	newValue := new(DetailedAmount4)
	d.Fees = append(d.Fees, newValue)
	return newValue
}

func (d *DetailedAmount5) AddRebate() *DetailedAmount4 {
	newValue := new(DetailedAmount4)
	d.Rebate = append(d.Rebate, newValue)
	return newValue
}

func (d *DetailedAmount5) AddValueAddedTax() *DetailedAmount4 {
	newValue := new(DetailedAmount4)
	d.ValueAddedTax = append(d.ValueAddedTax, newValue)
	return newValue
}

// Detailed amounts associated with the total amount of transaction.
type DetailedAmount6 struct {

	// Amount value.
	Amount *ImpliedCurrencyAndAmount `xml:"Amt"`

	// Date and time of the payment.
	DateTime *ISODateTime `xml:"DtTm"`

	// Card data entry mode for the related payment.
	CardDataEntryMode *CardDataReading1Code `xml:"CardDataNtryMd,omitempty"`

	// Data of an integrated circuit card application for the related payment.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`

	// Short description of the amount to display or print.
	Label *Max140Text `xml:"Labl,omitempty"`
}

func (d *DetailedAmount6) SetAmount(value, currency string) {
	d.Amount = NewImpliedCurrencyAndAmount(value, currency)
}

func (d *DetailedAmount6) SetDateTime(value string) {
	d.DateTime = (*ISODateTime)(&value)
}

func (d *DetailedAmount6) SetCardDataEntryMode(value string) {
	d.CardDataEntryMode = (*CardDataReading1Code)(&value)
}

func (d *DetailedAmount6) SetICCRelatedData(value string) {
	d.ICCRelatedData = (*Max10000Binary)(&value)
}

func (d *DetailedAmount6) SetLabel(value string) {
	d.Label = (*Max140Text)(&value)
}

// Detailed amounts associated with the total amount of transaction.
type DetailedAmount7 struct {

	// Cash-back amount.
	CashBack *ImpliedCurrencyAndAmount `xml:"CshBck,omitempty"`

	// Gratuity amount.
	Gratuity *ImpliedCurrencyAndAmount `xml:"Grtty,omitempty"`

	// Fees amount.
	Fees []*DetailedAmount4 `xml:"Fees,omitempty"`

	// Global rebate of the transaction. This amount is counted as a negative amount.
	Rebate []*DetailedAmount4 `xml:"Rbt,omitempty"`

	// Value added tax amount.
	ValueAddedTax []*DetailedAmount4 `xml:"ValAddedTax,omitempty"`

	// Additional charge paid by the cardholder. For example airline credit card surcharge.
	Surcharge []*DetailedAmount4 `xml:"Srchrg,omitempty"`
}

func (d *DetailedAmount7) SetCashBack(value, currency string) {
	d.CashBack = NewImpliedCurrencyAndAmount(value, currency)
}

func (d *DetailedAmount7) SetGratuity(value, currency string) {
	d.Gratuity = NewImpliedCurrencyAndAmount(value, currency)
}

func (d *DetailedAmount7) AddFees() *DetailedAmount4 {
	newValue := new(DetailedAmount4)
	d.Fees = append(d.Fees, newValue)
	return newValue
}

func (d *DetailedAmount7) AddRebate() *DetailedAmount4 {
	newValue := new(DetailedAmount4)
	d.Rebate = append(d.Rebate, newValue)
	return newValue
}

func (d *DetailedAmount7) AddValueAddedTax() *DetailedAmount4 {
	newValue := new(DetailedAmount4)
	d.ValueAddedTax = append(d.ValueAddedTax, newValue)
	return newValue
}

func (d *DetailedAmount7) AddSurcharge() *DetailedAmount4 {
	newValue := new(DetailedAmount4)
	d.Surcharge = append(d.Surcharge, newValue)
	return newValue
}

// Detailed amounts associated with the total amount of transaction.
type DetailedAmount8 struct {

	// Amount after the currency exchange.
	// It corresponds to ISO 8583 field number 6, completed by the field number 51 for the versions 87 and 93.
	Amount *ImpliedCurrencyAndAmount `xml:"Amt"`

	// Exchange rate to the currency of the amount.
	// It corresponds to ISO 8583 field number 10.
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Date and time at which the exchange rate has been quoted.
	QuotationDate *ISODateTime `xml:"QtnDt,omitempty"`

	// Text to display on the cardholder or to print on the cardholder bank statement.
	Label *Max140Text `xml:"Labl,omitempty"`
}

func (d *DetailedAmount8) SetAmount(value, currency string) {
	d.Amount = NewImpliedCurrencyAndAmount(value, currency)
}

func (d *DetailedAmount8) SetExchangeRate(value string) {
	d.ExchangeRate = (*BaseOneRate)(&value)
}

func (d *DetailedAmount8) SetQuotationDate(value string) {
	d.QuotationDate = (*ISODateTime)(&value)
}

func (d *DetailedAmount8) SetLabel(value string) {
	d.Label = (*Max140Text)(&value)
}

// Detailed amounts associated with the total amount of transaction.
type DetailedAmount9 struct {

	// Type or class of amount.
	Type *TypeOfAmount5Code `xml:"Tp"`

	// Additional information to specify the type of amount.
	AdditionalType *Max35Text `xml:"AddtlTp,omitempty"`

	// Amount value.
	Amount *ImpliedCurrencyAndAmount `xml:"Amt"`

	// Short description of the amount to provide to the cardholder.
	Label *Max140Text `xml:"Labl,omitempty"`
}

func (d *DetailedAmount9) SetType(value string) {
	d.Type = (*TypeOfAmount5Code)(&value)
}

func (d *DetailedAmount9) SetAdditionalType(value string) {
	d.AdditionalType = (*Max35Text)(&value)
}

func (d *DetailedAmount9) SetAmount(value, currency string) {
	d.Amount = NewImpliedCurrencyAndAmount(value, currency)
}

func (d *DetailedAmount9) SetLabel(value string) {
	d.Label = (*Max140Text)(&value)
}
