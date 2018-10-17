package iso20022

// Specifies the elements of an entry in the report.
type ReportEntry1 struct {

	// Amount of money in the cash entry.
	Amount *CurrencyAndAmount `xml:"Amt"`

	// Specifies if an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates whether the entry is the result of a reversal operation.
	//
	// Usage : this element should only be present if the entry is the result of a reversal operation.
	// If the CreditDebitIndicator is CRDT and ReversalIndicator is Yes, the original operation was a debit entry.
	// If the CreditDebitIndicator is DBIT and ReversalIndicator is Yes, the original operation was a credit entry.
	ReversalIndicator *TrueFalseIndicator `xml:"RvslInd,omitempty"`

	// Status of an entry on the books of the account servicer.
	Status *EntryStatus2Code `xml:"Sts"`

	// Date and time when an entry is posted to an account on the account servicer's books.
	//
	// Usage : Booking date is only present if Status is booked.
	BookingDate *DateAndDateTimeChoice `xml:"BookgDt,omitempty"`

	// Date and time assets become available to the account owner (in a credit entry), or cease to be available to the account owner (in a debit entry).
	//
	// Usage : When entry status is pending , and value date is present, value date refers to an expected/requested value date.
	// For entries which are subject to availability/float (and for which availability information is present), value date must  not be used, as the availability component identifies the number of availability days.
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`

	// Account servicing institution's reference for the entry.
	AccountServicerReference *Max35Text `xml:"AcctSvcrRef,omitempty"`

	// Set of elements used to indicate when the booked amount of money will become available, ie can be accessed and start generating interest.
	//
	// Usage : this type of info is eg used in US, and is linked to particular instruments, such as cheques.
	// Example : When a cheque is deposited, it will be booked on the deposit day, but the funds will only be accessible as of the indicated availability day (according to national banking regulations).
	Availability []*CashBalanceAvailability1 `xml:"Avlbty,omitempty"`

	// Set of elements to fully identify the type of underlying transaction resulting in the entry.
	BankTransactionCode *BankTransactionCodeStructure1 `xml:"BkTxCd"`

	// Indicates whether the transaction is exempt from commission.
	CommissionWaiverIndicator *YesNoIndicator `xml:"ComssnWvrInd,omitempty"`

	// Indicates whether the underlying transaction details are provided through a separate message, eg in case of aggregate bookings.
	AdditionalInformationIndicator *MessageIdentification2 `xml:"AddtlInfInd,omitempty"`

	// Set of elements providing details on batched transactions.
	//
	// Usage : this element can be repeated in case more than one batch is included in the entry, eg, in lockbox scenarios, to specify the ID and number of transactions included in each of the batches.
	Batch []*BatchInformation1 `xml:"Btch,omitempty"`

	// Set of elements providing information on the original amount.
	//
	// Usage : This component (on entry level) should be used when a total original batch or aggregate amount has to be provided. (If required, the individual original amounts can be included in the same component on transaction details level).
	AmountDetails *AmountAndCurrencyExchange2 `xml:"AmtDtls,omitempty"`

	// Provides information on the charges included in the entry amount.
	//
	// Usage : this component is used on entry level in case of batch or aggregate bookings.
	Charges []*ChargesInformation3 `xml:"Chrgs,omitempty"`

	// Set of elements providing details on the interest amount included in the entry amount.
	//
	// Usage : This component is used on entry level in case of batch or aggregate bookings.
	Interest []*TransactionInterest1 `xml:"Intrst,omitempty"`

	// Set of elements providing information on the underlying transaction (s).
	TransactionDetails []*EntryTransaction1 `xml:"TxDtls,omitempty"`

	// Further details on the entry details.
	AdditionalEntryInformation *Max500Text `xml:"AddtlNtryInf,omitempty"`
}

func (r *ReportEntry1) SetAmount(value, currency string) {
	r.Amount = NewCurrencyAndAmount(value, currency)
}

func (r *ReportEntry1) SetCreditDebitIndicator(value string) {
	r.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (r *ReportEntry1) SetReversalIndicator(value string) {
	r.ReversalIndicator = (*TrueFalseIndicator)(&value)
}

func (r *ReportEntry1) SetStatus(value string) {
	r.Status = (*EntryStatus2Code)(&value)
}

func (r *ReportEntry1) AddBookingDate() *DateAndDateTimeChoice {
	r.BookingDate = new(DateAndDateTimeChoice)
	return r.BookingDate
}

func (r *ReportEntry1) AddValueDate() *DateAndDateTimeChoice {
	r.ValueDate = new(DateAndDateTimeChoice)
	return r.ValueDate
}

func (r *ReportEntry1) SetAccountServicerReference(value string) {
	r.AccountServicerReference = (*Max35Text)(&value)
}

func (r *ReportEntry1) AddAvailability() *CashBalanceAvailability1 {
	newValue := new(CashBalanceAvailability1)
	r.Availability = append(r.Availability, newValue)
	return newValue
}

func (r *ReportEntry1) AddBankTransactionCode() *BankTransactionCodeStructure1 {
	r.BankTransactionCode = new(BankTransactionCodeStructure1)
	return r.BankTransactionCode
}

func (r *ReportEntry1) SetCommissionWaiverIndicator(value string) {
	r.CommissionWaiverIndicator = (*YesNoIndicator)(&value)
}

func (r *ReportEntry1) AddAdditionalInformationIndicator() *MessageIdentification2 {
	r.AdditionalInformationIndicator = new(MessageIdentification2)
	return r.AdditionalInformationIndicator
}

func (r *ReportEntry1) AddBatch() *BatchInformation1 {
	newValue := new(BatchInformation1)
	r.Batch = append(r.Batch, newValue)
	return newValue
}

func (r *ReportEntry1) AddAmountDetails() *AmountAndCurrencyExchange2 {
	r.AmountDetails = new(AmountAndCurrencyExchange2)
	return r.AmountDetails
}

func (r *ReportEntry1) AddCharges() *ChargesInformation3 {
	newValue := new(ChargesInformation3)
	r.Charges = append(r.Charges, newValue)
	return newValue
}

func (r *ReportEntry1) AddInterest() *TransactionInterest1 {
	newValue := new(TransactionInterest1)
	r.Interest = append(r.Interest, newValue)
	return newValue
}

func (r *ReportEntry1) AddTransactionDetails() *EntryTransaction1 {
	newValue := new(EntryTransaction1)
	r.TransactionDetails = append(r.TransactionDetails, newValue)
	return newValue
}

func (r *ReportEntry1) SetAdditionalEntryInformation(value string) {
	r.AdditionalEntryInformation = (*Max500Text)(&value)
}

// Set of elements used to provide information on an entry in the report.
type ReportEntry2 struct {

	// Unique reference for the entry.
	EntryReference *Max35Text `xml:"NtryRef,omitempty"`

	// Amount of money in the cash entry.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether the entry is a credit or a debit entry.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates whether or not the entry is the result of a reversal.
	// Usage: This element should only be present if the entry is the result of a reversal.
	// If the CreditDebitIndicator is CRDT and ReversalIndicator is Yes, the original operation was a debit entry.
	// If the CreditDebitIndicator is DBIT and ReversalIndicator is Yes, the original operation was a credit entry.
	ReversalIndicator *TrueFalseIndicator `xml:"RvslInd,omitempty"`

	// Status of an entry on the books of the account servicer.
	Status *EntryStatus2Code `xml:"Sts"`

	// Date and time when an entry is posted to an account on the account servicer's books.
	//
	// Usage: Booking date is the expected booking date, unless the status is booked, in which case it is the actual booking date.
	BookingDate *DateAndDateTimeChoice `xml:"BookgDt,omitempty"`

	// Date and time at which assets become available to the account owner in case of a credit entry, or cease to be available to the account owner in case of a debit entry.
	// Usage: If entry status is pending and value date is present, then the value date refers to an expected/requested value date.
	// For entries subject to availability/float and for which availability information is provided, the value date must not be used. In this case the availability component identifies the number of availability days.
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`

	// Unique reference as assigned by the account servicing institution to unambiguously identify the entry.
	AccountServicerReference *Max35Text `xml:"AcctSvcrRef,omitempty"`

	// Set of elements used to indicate when the booked amount of money will become available, that is can be accessed and starts generating interest.
	//
	// Usage: This type of information is used in the US and is linked to particular instruments such as cheques.
	// Example: When a cheque is deposited, it will be booked on the deposit day, but the amount of money will only be accessible as of the indicated availability day (according to national banking regulations).
	Availability []*CashBalanceAvailability2 `xml:"Avlbty,omitempty"`

	// Set of elements used to fully identify the type of underlying transaction resulting in an entry.
	BankTransactionCode *BankTransactionCodeStructure4 `xml:"BkTxCd"`

	// Indicates whether the transaction is exempt from commission.
	CommissionWaiverIndicator *YesNoIndicator `xml:"ComssnWvrInd,omitempty"`

	// Indicates whether the underlying transaction details are provided through a separate message, as in the case of aggregate bookings.
	AdditionalInformationIndicator *MessageIdentification2 `xml:"AddtlInfInd,omitempty"`

	// Set of elements providing information on the original amount.
	//
	// Usage: This component (on entry level) should be used when a total original batch or aggregate amount has to be provided. If required, the individual original amounts can be included in the same component on transaction details level.
	AmountDetails *AmountAndCurrencyExchange3 `xml:"AmtDtls,omitempty"`

	// Provides information on the charges included in the entry amount.
	//
	// Usage: This component is used on entry level in case of batch or aggregate bookings.
	Charges []*ChargesInformation6 `xml:"Chrgs,omitempty"`

	// Channel used to technically input the instruction related to the entry.
	TechnicalInputChannel *TechnicalInputChannel1Choice `xml:"TechInptChanl,omitempty"`

	// Set of elements used to provide details of the interest amount included in the entry amount.
	//
	// Usage: This component is used on entry level in the case of batch or aggregate bookings.
	Interest []*TransactionInterest2 `xml:"Intrst,omitempty"`

	// Set of elements used to provide details on the entry.
	EntryDetails []*EntryDetails1 `xml:"NtryDtls,omitempty"`

	// Further details of the entry.
	AdditionalEntryInformation *Max500Text `xml:"AddtlNtryInf,omitempty"`
}

func (r *ReportEntry2) SetEntryReference(value string) {
	r.EntryReference = (*Max35Text)(&value)
}

func (r *ReportEntry2) SetAmount(value, currency string) {
	r.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *ReportEntry2) SetCreditDebitIndicator(value string) {
	r.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (r *ReportEntry2) SetReversalIndicator(value string) {
	r.ReversalIndicator = (*TrueFalseIndicator)(&value)
}

func (r *ReportEntry2) SetStatus(value string) {
	r.Status = (*EntryStatus2Code)(&value)
}

func (r *ReportEntry2) AddBookingDate() *DateAndDateTimeChoice {
	r.BookingDate = new(DateAndDateTimeChoice)
	return r.BookingDate
}

func (r *ReportEntry2) AddValueDate() *DateAndDateTimeChoice {
	r.ValueDate = new(DateAndDateTimeChoice)
	return r.ValueDate
}

func (r *ReportEntry2) SetAccountServicerReference(value string) {
	r.AccountServicerReference = (*Max35Text)(&value)
}

func (r *ReportEntry2) AddAvailability() *CashBalanceAvailability2 {
	newValue := new(CashBalanceAvailability2)
	r.Availability = append(r.Availability, newValue)
	return newValue
}

func (r *ReportEntry2) AddBankTransactionCode() *BankTransactionCodeStructure4 {
	r.BankTransactionCode = new(BankTransactionCodeStructure4)
	return r.BankTransactionCode
}

func (r *ReportEntry2) SetCommissionWaiverIndicator(value string) {
	r.CommissionWaiverIndicator = (*YesNoIndicator)(&value)
}

func (r *ReportEntry2) AddAdditionalInformationIndicator() *MessageIdentification2 {
	r.AdditionalInformationIndicator = new(MessageIdentification2)
	return r.AdditionalInformationIndicator
}

func (r *ReportEntry2) AddAmountDetails() *AmountAndCurrencyExchange3 {
	r.AmountDetails = new(AmountAndCurrencyExchange3)
	return r.AmountDetails
}

func (r *ReportEntry2) AddCharges() *ChargesInformation6 {
	newValue := new(ChargesInformation6)
	r.Charges = append(r.Charges, newValue)
	return newValue
}

func (r *ReportEntry2) AddTechnicalInputChannel() *TechnicalInputChannel1Choice {
	r.TechnicalInputChannel = new(TechnicalInputChannel1Choice)
	return r.TechnicalInputChannel
}

func (r *ReportEntry2) AddInterest() *TransactionInterest2 {
	newValue := new(TransactionInterest2)
	r.Interest = append(r.Interest, newValue)
	return newValue
}

func (r *ReportEntry2) AddEntryDetails() *EntryDetails1 {
	newValue := new(EntryDetails1)
	r.EntryDetails = append(r.EntryDetails, newValue)
	return newValue
}

func (r *ReportEntry2) SetAdditionalEntryInformation(value string) {
	r.AdditionalEntryInformation = (*Max500Text)(&value)
}

// Provides further details on an entry in the report.
type ReportEntry3 struct {

	// Unique reference for the entry.
	EntryReference *Max35Text `xml:"NtryRef,omitempty"`

	// Amount of money in the cash entry.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether the entry is a credit or a debit entry.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates whether or not the entry is the result of a reversal.
	// Usage: This element should only be present if the entry is the result of a reversal.
	// If the CreditDebitIndicator is CRDT and ReversalIndicator is Yes, the original operation was a debit entry.
	// If the CreditDebitIndicator is DBIT and ReversalIndicator is Yes, the original operation was a credit entry.
	ReversalIndicator *TrueFalseIndicator `xml:"RvslInd,omitempty"`

	// Status of an entry on the books of the account servicer.
	Status *EntryStatus2Code `xml:"Sts"`

	// Date and time when an entry is posted to an account on the account servicer's books.
	//
	// Usage: Booking date is the expected booking date, unless the status is booked, in which case it is the actual booking date.
	BookingDate *DateAndDateTimeChoice `xml:"BookgDt,omitempty"`

	// Date and time at which assets become available to the account owner in case of a credit entry, or cease to be available to the account owner in case of a debit entry.
	// Usage: If entry status is pending and value date is present, then the value date refers to an expected/requested value date.
	// For entries subject to availability/float and for which availability information is provided, the value date must not be used. In this case the availability component identifies the number of availability days.
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`

	// Unique reference as assigned by the account servicing institution to unambiguously identify the entry.
	AccountServicerReference *Max35Text `xml:"AcctSvcrRef,omitempty"`

	// Indicates when the booked amount of money will become available, that is can be accessed and starts generating interest.
	//
	// Usage: This type of information is used in the US and is linked to particular instruments such as cheques.
	// Example: When a cheque is deposited, it will be booked on the deposit day, but the amount of money will only be accessible as of the indicated availability day (according to national banking regulations).
	Availability []*CashBalanceAvailability2 `xml:"Avlbty,omitempty"`

	// Set of elements used to fully identify the type of underlying transaction resulting in an entry.
	BankTransactionCode *BankTransactionCodeStructure4 `xml:"BkTxCd"`

	// Indicates whether the transaction is exempt from commission.
	CommissionWaiverIndicator *YesNoIndicator `xml:"ComssnWvrInd,omitempty"`

	// Indicates whether the underlying transaction details are provided through a separate message, as in the case of aggregate bookings.
	AdditionalInformationIndicator *MessageIdentification2 `xml:"AddtlInfInd,omitempty"`

	// Provides information on the original amount.
	//
	// Usage: This component (on entry level) should be used when a total original batch or aggregate amount has to be provided. If required, the individual original amounts can be included in the same component on transaction details level.
	AmountDetails *AmountAndCurrencyExchange3 `xml:"AmtDtls,omitempty"`

	// Provides information on the charges included in the entry amount.
	//
	// Usage: This component is used on entry level in case of batch or aggregate bookings.
	Charges []*Charges3 `xml:"Chrgs,omitempty"`

	// Channel used to technically input the instruction related to the entry.
	TechnicalInputChannel *TechnicalInputChannel1Choice `xml:"TechInptChanl,omitempty"`

	// Provides details of the interest amount included in the entry amount.
	//
	// Usage: This component is used on entry level in the case of batch or aggregate bookings.
	Interest []*TransactionInterest3 `xml:"Intrst,omitempty"`

	// Provides details of the card transaction included in the entry amount, when globalised by the account servicer .
	CardTransaction *CardEntry1 `xml:"CardTx,omitempty"`

	// Provides details on the entry.
	EntryDetails []*EntryDetails2 `xml:"NtryDtls,omitempty"`

	// Further details of the entry.
	AdditionalEntryInformation *Max500Text `xml:"AddtlNtryInf,omitempty"`
}

func (r *ReportEntry3) SetEntryReference(value string) {
	r.EntryReference = (*Max35Text)(&value)
}

func (r *ReportEntry3) SetAmount(value, currency string) {
	r.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *ReportEntry3) SetCreditDebitIndicator(value string) {
	r.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (r *ReportEntry3) SetReversalIndicator(value string) {
	r.ReversalIndicator = (*TrueFalseIndicator)(&value)
}

func (r *ReportEntry3) SetStatus(value string) {
	r.Status = (*EntryStatus2Code)(&value)
}

func (r *ReportEntry3) AddBookingDate() *DateAndDateTimeChoice {
	r.BookingDate = new(DateAndDateTimeChoice)
	return r.BookingDate
}

func (r *ReportEntry3) AddValueDate() *DateAndDateTimeChoice {
	r.ValueDate = new(DateAndDateTimeChoice)
	return r.ValueDate
}

func (r *ReportEntry3) SetAccountServicerReference(value string) {
	r.AccountServicerReference = (*Max35Text)(&value)
}

func (r *ReportEntry3) AddAvailability() *CashBalanceAvailability2 {
	newValue := new(CashBalanceAvailability2)
	r.Availability = append(r.Availability, newValue)
	return newValue
}

func (r *ReportEntry3) AddBankTransactionCode() *BankTransactionCodeStructure4 {
	r.BankTransactionCode = new(BankTransactionCodeStructure4)
	return r.BankTransactionCode
}

func (r *ReportEntry3) SetCommissionWaiverIndicator(value string) {
	r.CommissionWaiverIndicator = (*YesNoIndicator)(&value)
}

func (r *ReportEntry3) AddAdditionalInformationIndicator() *MessageIdentification2 {
	r.AdditionalInformationIndicator = new(MessageIdentification2)
	return r.AdditionalInformationIndicator
}

func (r *ReportEntry3) AddAmountDetails() *AmountAndCurrencyExchange3 {
	r.AmountDetails = new(AmountAndCurrencyExchange3)
	return r.AmountDetails
}

func (r *ReportEntry3) AddCharges() *Charges3 {
	newValue := new(Charges3)
	r.Charges = append(r.Charges, newValue)
	return newValue
}

func (r *ReportEntry3) AddTechnicalInputChannel() *TechnicalInputChannel1Choice {
	r.TechnicalInputChannel = new(TechnicalInputChannel1Choice)
	return r.TechnicalInputChannel
}

func (r *ReportEntry3) AddInterest() *TransactionInterest3 {
	newValue := new(TransactionInterest3)
	r.Interest = append(r.Interest, newValue)
	return newValue
}

func (r *ReportEntry3) AddCardTransaction() *CardEntry1 {
	r.CardTransaction = new(CardEntry1)
	return r.CardTransaction
}

func (r *ReportEntry3) AddEntryDetails() *EntryDetails2 {
	newValue := new(EntryDetails2)
	r.EntryDetails = append(r.EntryDetails, newValue)
	return newValue
}

func (r *ReportEntry3) SetAdditionalEntryInformation(value string) {
	r.AdditionalEntryInformation = (*Max500Text)(&value)
}

// Provides further details on an entry in the report.
type ReportEntry4 struct {

	// Unique reference for the entry.
	EntryReference *Max35Text `xml:"NtryRef,omitempty"`

	// Amount of money in the cash entry.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether the entry is a credit or a debit entry.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates whether or not the entry is the result of a reversal.
	// Usage: This element should only be present if the entry is the result of a reversal.
	// If the CreditDebitIndicator is CRDT and ReversalIndicator is Yes, the original operation was a debit entry.
	// If the CreditDebitIndicator is DBIT and ReversalIndicator is Yes, the original operation was a credit entry.
	ReversalIndicator *TrueFalseIndicator `xml:"RvslInd,omitempty"`

	// Status of an entry on the books of the account servicer.
	Status *EntryStatus2Code `xml:"Sts"`

	// Date and time when an entry is posted to an account on the account servicer's books.
	//
	// Usage: Booking date is the expected booking date, unless the status is booked, in which case it is the actual booking date.
	BookingDate *DateAndDateTimeChoice `xml:"BookgDt,omitempty"`

	// Date and time at which assets become available to the account owner in case of a credit entry, or cease to be available to the account owner in case of a debit entry.
	// Usage: If entry status is pending and value date is present, then the value date refers to an expected/requested value date.
	// For entries subject to availability/float and for which availability information is provided, the value date must not be used. In this case the availability component identifies the number of availability days.
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`

	// Unique reference as assigned by the account servicing institution to unambiguously identify the entry.
	AccountServicerReference *Max35Text `xml:"AcctSvcrRef,omitempty"`

	// Indicates when the booked amount of money will become available, that is can be accessed and starts generating interest.
	//
	// Usage: This type of information is used in the US and is linked to particular instruments such as cheques.
	// Example: When a cheque is deposited, it will be booked on the deposit day, but the amount of money will only be accessible as of the indicated availability day (according to national banking regulations).
	Availability []*CashBalanceAvailability2 `xml:"Avlbty,omitempty"`

	// Set of elements used to fully identify the type of underlying transaction resulting in an entry.
	BankTransactionCode *BankTransactionCodeStructure4 `xml:"BkTxCd"`

	// Indicates whether the transaction is exempt from commission.
	CommissionWaiverIndicator *YesNoIndicator `xml:"ComssnWvrInd,omitempty"`

	// Indicates whether the underlying transaction details are provided through a separate message, as in the case of aggregate bookings.
	AdditionalInformationIndicator *MessageIdentification2 `xml:"AddtlInfInd,omitempty"`

	// Provides information on the original amount.
	//
	// Usage: This component (on entry level) should be used when a total original batch or aggregate amount has to be provided. If required, the individual original amounts can be included in the same component on transaction details level.
	AmountDetails *AmountAndCurrencyExchange3 `xml:"AmtDtls,omitempty"`

	// Provides information on the charges, pre-advised or included in the entry amount .
	//
	// Usage: This component is used on entry level in case of batch or aggregate bookings.
	//
	Charges *Charges4 `xml:"Chrgs,omitempty"`

	// Channel used to technically input the instruction related to the entry.
	TechnicalInputChannel *TechnicalInputChannel1Choice `xml:"TechInptChanl,omitempty"`

	// Provides details of the interest amount included in the entry amount.
	//
	// Usage: This component is used on entry level in the case of batch or aggregate bookings.
	Interest *TransactionInterest3 `xml:"Intrst,omitempty"`

	// Provides details of the card transaction included in the entry amount, when globalised by the account servicer .
	CardTransaction *CardEntry1 `xml:"CardTx,omitempty"`

	// Provides details on the entry.
	EntryDetails []*EntryDetails3 `xml:"NtryDtls,omitempty"`

	// Further details of the entry.
	AdditionalEntryInformation *Max500Text `xml:"AddtlNtryInf,omitempty"`
}

func (r *ReportEntry4) SetEntryReference(value string) {
	r.EntryReference = (*Max35Text)(&value)
}

func (r *ReportEntry4) SetAmount(value, currency string) {
	r.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *ReportEntry4) SetCreditDebitIndicator(value string) {
	r.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (r *ReportEntry4) SetReversalIndicator(value string) {
	r.ReversalIndicator = (*TrueFalseIndicator)(&value)
}

func (r *ReportEntry4) SetStatus(value string) {
	r.Status = (*EntryStatus2Code)(&value)
}

func (r *ReportEntry4) AddBookingDate() *DateAndDateTimeChoice {
	r.BookingDate = new(DateAndDateTimeChoice)
	return r.BookingDate
}

func (r *ReportEntry4) AddValueDate() *DateAndDateTimeChoice {
	r.ValueDate = new(DateAndDateTimeChoice)
	return r.ValueDate
}

func (r *ReportEntry4) SetAccountServicerReference(value string) {
	r.AccountServicerReference = (*Max35Text)(&value)
}

func (r *ReportEntry4) AddAvailability() *CashBalanceAvailability2 {
	newValue := new(CashBalanceAvailability2)
	r.Availability = append(r.Availability, newValue)
	return newValue
}

func (r *ReportEntry4) AddBankTransactionCode() *BankTransactionCodeStructure4 {
	r.BankTransactionCode = new(BankTransactionCodeStructure4)
	return r.BankTransactionCode
}

func (r *ReportEntry4) SetCommissionWaiverIndicator(value string) {
	r.CommissionWaiverIndicator = (*YesNoIndicator)(&value)
}

func (r *ReportEntry4) AddAdditionalInformationIndicator() *MessageIdentification2 {
	r.AdditionalInformationIndicator = new(MessageIdentification2)
	return r.AdditionalInformationIndicator
}

func (r *ReportEntry4) AddAmountDetails() *AmountAndCurrencyExchange3 {
	r.AmountDetails = new(AmountAndCurrencyExchange3)
	return r.AmountDetails
}

func (r *ReportEntry4) AddCharges() *Charges4 {
	r.Charges = new(Charges4)
	return r.Charges
}

func (r *ReportEntry4) AddTechnicalInputChannel() *TechnicalInputChannel1Choice {
	r.TechnicalInputChannel = new(TechnicalInputChannel1Choice)
	return r.TechnicalInputChannel
}

func (r *ReportEntry4) AddInterest() *TransactionInterest3 {
	r.Interest = new(TransactionInterest3)
	return r.Interest
}

func (r *ReportEntry4) AddCardTransaction() *CardEntry1 {
	r.CardTransaction = new(CardEntry1)
	return r.CardTransaction
}

func (r *ReportEntry4) AddEntryDetails() *EntryDetails3 {
	newValue := new(EntryDetails3)
	r.EntryDetails = append(r.EntryDetails, newValue)
	return newValue
}

func (r *ReportEntry4) SetAdditionalEntryInformation(value string) {
	r.AdditionalEntryInformation = (*Max500Text)(&value)
}

// Provides further details on an entry in the report.
type ReportEntry7 struct {

	// Unique reference for the entry.
	EntryReference *Max35Text `xml:"NtryRef,omitempty"`

	// Amount of money in the cash entry.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether the entry is a credit or a debit entry.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates whether or not the entry is the result of a reversal.
	// Usage: This element should only be present if the entry is the result of a reversal.
	// If the CreditDebitIndicator is CRDT and ReversalIndicator is Yes, the original operation was a debit entry.
	// If the CreditDebitIndicator is DBIT and ReversalIndicator is Yes, the original operation was a credit entry.
	ReversalIndicator *TrueFalseIndicator `xml:"RvslInd,omitempty"`

	// Status of an entry on the books of the account servicer.
	Status *EntryStatus2Code `xml:"Sts"`

	// Date and time when an entry is posted to an account on the account servicer's books.
	//
	// Usage: Booking date is the expected booking date, unless the status is booked, in which case it is the actual booking date.
	BookingDate *DateAndDateTimeChoice `xml:"BookgDt,omitempty"`

	// Date and time at which assets become available to the account owner in case of a credit entry, or cease to be available to the account owner in case of a debit entry.
	// Usage: If entry status is pending and value date is present, then the value date refers to an expected/requested value date.
	// For entries subject to availability/float and for which availability information is provided, the value date must not be used. In this case the availability component identifies the number of availability days.
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`

	// Unique reference as assigned by the account servicing institution to unambiguously identify the entry.
	AccountServicerReference *Max35Text `xml:"AcctSvcrRef,omitempty"`

	// Indicates when the booked amount of money will become available, that is can be accessed and starts generating interest.
	//
	// Usage: This type of information is used in the US and is linked to particular instruments such as cheques.
	// Example: When a cheque is deposited, it will be booked on the deposit day, but the amount of money will only be accessible as of the indicated availability day (according to national banking regulations).
	Availability []*CashBalanceAvailability2 `xml:"Avlbty,omitempty"`

	// Set of elements used to fully identify the type of underlying transaction resulting in an entry.
	BankTransactionCode *BankTransactionCodeStructure4 `xml:"BkTxCd"`

	// Indicates whether the transaction is exempt from commission.
	CommissionWaiverIndicator *YesNoIndicator `xml:"ComssnWvrInd,omitempty"`

	// Indicates whether the underlying transaction details are provided through a separate message, as in the case of aggregate bookings.
	AdditionalInformationIndicator *MessageIdentification2 `xml:"AddtlInfInd,omitempty"`

	// Provides information on the original amount.
	//
	// Usage: This component (on entry level) should be used when a total original batch or aggregate amount has to be provided. If required, the individual original amounts can be included in the same component on transaction details level.
	AmountDetails *AmountAndCurrencyExchange3 `xml:"AmtDtls,omitempty"`

	// Provides information on the charges, pre-advised or included in the entry amount .
	//
	// Usage: This component is used on entry level in case of batch or aggregate bookings.
	//
	Charges *Charges4 `xml:"Chrgs,omitempty"`

	// Channel used to technically input the instruction related to the entry.
	TechnicalInputChannel *TechnicalInputChannel1Choice `xml:"TechInptChanl,omitempty"`

	// Provides details of the interest amount included in the entry amount.
	//
	// Usage: This component is used on entry level in the case of batch or aggregate bookings.
	Interest *TransactionInterest3 `xml:"Intrst,omitempty"`

	// Provides details of the card transaction included in the entry amount, when globalised by the account servicer .
	CardTransaction *CardEntry2 `xml:"CardTx,omitempty"`

	// Provides details on the entry.
	EntryDetails []*EntryDetails6 `xml:"NtryDtls,omitempty"`

	// Further details of the entry.
	AdditionalEntryInformation *Max500Text `xml:"AddtlNtryInf,omitempty"`
}

func (r *ReportEntry7) SetEntryReference(value string) {
	r.EntryReference = (*Max35Text)(&value)
}

func (r *ReportEntry7) SetAmount(value, currency string) {
	r.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *ReportEntry7) SetCreditDebitIndicator(value string) {
	r.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (r *ReportEntry7) SetReversalIndicator(value string) {
	r.ReversalIndicator = (*TrueFalseIndicator)(&value)
}

func (r *ReportEntry7) SetStatus(value string) {
	r.Status = (*EntryStatus2Code)(&value)
}

func (r *ReportEntry7) AddBookingDate() *DateAndDateTimeChoice {
	r.BookingDate = new(DateAndDateTimeChoice)
	return r.BookingDate
}

func (r *ReportEntry7) AddValueDate() *DateAndDateTimeChoice {
	r.ValueDate = new(DateAndDateTimeChoice)
	return r.ValueDate
}

func (r *ReportEntry7) SetAccountServicerReference(value string) {
	r.AccountServicerReference = (*Max35Text)(&value)
}

func (r *ReportEntry7) AddAvailability() *CashBalanceAvailability2 {
	newValue := new(CashBalanceAvailability2)
	r.Availability = append(r.Availability, newValue)
	return newValue
}

func (r *ReportEntry7) AddBankTransactionCode() *BankTransactionCodeStructure4 {
	r.BankTransactionCode = new(BankTransactionCodeStructure4)
	return r.BankTransactionCode
}

func (r *ReportEntry7) SetCommissionWaiverIndicator(value string) {
	r.CommissionWaiverIndicator = (*YesNoIndicator)(&value)
}

func (r *ReportEntry7) AddAdditionalInformationIndicator() *MessageIdentification2 {
	r.AdditionalInformationIndicator = new(MessageIdentification2)
	return r.AdditionalInformationIndicator
}

func (r *ReportEntry7) AddAmountDetails() *AmountAndCurrencyExchange3 {
	r.AmountDetails = new(AmountAndCurrencyExchange3)
	return r.AmountDetails
}

func (r *ReportEntry7) AddCharges() *Charges4 {
	r.Charges = new(Charges4)
	return r.Charges
}

func (r *ReportEntry7) AddTechnicalInputChannel() *TechnicalInputChannel1Choice {
	r.TechnicalInputChannel = new(TechnicalInputChannel1Choice)
	return r.TechnicalInputChannel
}

func (r *ReportEntry7) AddInterest() *TransactionInterest3 {
	r.Interest = new(TransactionInterest3)
	return r.Interest
}

func (r *ReportEntry7) AddCardTransaction() *CardEntry2 {
	r.CardTransaction = new(CardEntry2)
	return r.CardTransaction
}

func (r *ReportEntry7) AddEntryDetails() *EntryDetails6 {
	newValue := new(EntryDetails6)
	r.EntryDetails = append(r.EntryDetails, newValue)
	return newValue
}

func (r *ReportEntry7) SetAdditionalEntryInformation(value string) {
	r.AdditionalEntryInformation = (*Max500Text)(&value)
}

// Provides further details on an entry in the report.
type ReportEntry8 struct {

	// Unique reference for the entry.
	EntryReference *Max35Text `xml:"NtryRef,omitempty"`

	// Amount of money in the cash entry.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether the entry is a credit or a debit entry.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates whether or not the entry is the result of a reversal.
	// Usage: This element should only be present if the entry is the result of a reversal.
	// If the CreditDebitIndicator is CRDT and ReversalIndicator is Yes, the original operation was a debit entry.
	// If the CreditDebitIndicator is DBIT and ReversalIndicator is Yes, the original operation was a credit entry.
	ReversalIndicator *TrueFalseIndicator `xml:"RvslInd,omitempty"`

	// Status of an entry on the books of the account servicer.
	Status *EntryStatus2Code `xml:"Sts"`

	// Date and time when an entry is posted to an account on the account servicer's books.
	//
	// Usage: Booking date is the expected booking date, unless the status is booked, in which case it is the actual booking date.
	BookingDate *DateAndDateTimeChoice `xml:"BookgDt,omitempty"`

	// Date and time at which assets become available to the account owner in case of a credit entry, or cease to be available to the account owner in case of a debit entry.
	// Usage: If entry status is pending and value date is present, then the value date refers to an expected/requested value date.
	// For entries subject to availability/float and for which availability information is provided, the value date must not be used. In this case the availability component identifies the number of availability days.
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`

	// Unique reference as assigned by the account servicing institution to unambiguously identify the entry.
	AccountServicerReference *Max35Text `xml:"AcctSvcrRef,omitempty"`

	// Indicates when the booked amount of money will become available, that is can be accessed and starts generating interest.
	//
	// Usage: This type of information is used in the US and is linked to particular instruments such as cheques.
	// Example: When a cheque is deposited, it will be booked on the deposit day, but the amount of money will only be accessible as of the indicated availability day (according to national banking regulations).
	Availability []*CashAvailability1 `xml:"Avlbty,omitempty"`

	// Set of elements used to fully identify the type of underlying transaction resulting in an entry.
	BankTransactionCode *BankTransactionCodeStructure4 `xml:"BkTxCd"`

	// Indicates whether the transaction is exempt from commission.
	CommissionWaiverIndicator *YesNoIndicator `xml:"ComssnWvrInd,omitempty"`

	// Indicates whether the underlying transaction details are provided through a separate message, as in the case of aggregate bookings.
	AdditionalInformationIndicator *MessageIdentification2 `xml:"AddtlInfInd,omitempty"`

	// Provides information on the original amount.
	//
	// Usage: This component (on entry level) should be used when a total original batch or aggregate amount has to be provided. If required, the individual original amounts can be included in the same component on transaction details level.
	AmountDetails *AmountAndCurrencyExchange3 `xml:"AmtDtls,omitempty"`

	// Provides information on the charges, pre-advised or included in the entry amount .
	//
	// Usage: This component is used on entry level in case of batch or aggregate bookings.
	//
	Charges *Charges4 `xml:"Chrgs,omitempty"`

	// Channel used to technically input the instruction related to the entry.
	TechnicalInputChannel *TechnicalInputChannel1Choice `xml:"TechInptChanl,omitempty"`

	// Provides details of the interest amount included in the entry amount.
	//
	// Usage: This component is used on entry level in the case of batch or aggregate bookings.
	Interest *TransactionInterest3 `xml:"Intrst,omitempty"`

	// Provides details of the card transaction included in the entry amount, when globalised by the account servicer .
	CardTransaction *CardEntry2 `xml:"CardTx,omitempty"`

	// Provides details on the entry.
	EntryDetails []*EntryDetails7 `xml:"NtryDtls,omitempty"`

	// Further details of the entry.
	AdditionalEntryInformation *Max500Text `xml:"AddtlNtryInf,omitempty"`
}

func (r *ReportEntry8) SetEntryReference(value string) {
	r.EntryReference = (*Max35Text)(&value)
}

func (r *ReportEntry8) SetAmount(value, currency string) {
	r.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *ReportEntry8) SetCreditDebitIndicator(value string) {
	r.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (r *ReportEntry8) SetReversalIndicator(value string) {
	r.ReversalIndicator = (*TrueFalseIndicator)(&value)
}

func (r *ReportEntry8) SetStatus(value string) {
	r.Status = (*EntryStatus2Code)(&value)
}

func (r *ReportEntry8) AddBookingDate() *DateAndDateTimeChoice {
	r.BookingDate = new(DateAndDateTimeChoice)
	return r.BookingDate
}

func (r *ReportEntry8) AddValueDate() *DateAndDateTimeChoice {
	r.ValueDate = new(DateAndDateTimeChoice)
	return r.ValueDate
}

func (r *ReportEntry8) SetAccountServicerReference(value string) {
	r.AccountServicerReference = (*Max35Text)(&value)
}

func (r *ReportEntry8) AddAvailability() *CashAvailability1 {
	newValue := new(CashAvailability1)
	r.Availability = append(r.Availability, newValue)
	return newValue
}

func (r *ReportEntry8) AddBankTransactionCode() *BankTransactionCodeStructure4 {
	r.BankTransactionCode = new(BankTransactionCodeStructure4)
	return r.BankTransactionCode
}

func (r *ReportEntry8) SetCommissionWaiverIndicator(value string) {
	r.CommissionWaiverIndicator = (*YesNoIndicator)(&value)
}

func (r *ReportEntry8) AddAdditionalInformationIndicator() *MessageIdentification2 {
	r.AdditionalInformationIndicator = new(MessageIdentification2)
	return r.AdditionalInformationIndicator
}

func (r *ReportEntry8) AddAmountDetails() *AmountAndCurrencyExchange3 {
	r.AmountDetails = new(AmountAndCurrencyExchange3)
	return r.AmountDetails
}

func (r *ReportEntry8) AddCharges() *Charges4 {
	r.Charges = new(Charges4)
	return r.Charges
}

func (r *ReportEntry8) AddTechnicalInputChannel() *TechnicalInputChannel1Choice {
	r.TechnicalInputChannel = new(TechnicalInputChannel1Choice)
	return r.TechnicalInputChannel
}

func (r *ReportEntry8) AddInterest() *TransactionInterest3 {
	r.Interest = new(TransactionInterest3)
	return r.Interest
}

func (r *ReportEntry8) AddCardTransaction() *CardEntry2 {
	r.CardTransaction = new(CardEntry2)
	return r.CardTransaction
}

func (r *ReportEntry8) AddEntryDetails() *EntryDetails7 {
	newValue := new(EntryDetails7)
	r.EntryDetails = append(r.EntryDetails, newValue)
	return newValue
}

func (r *ReportEntry8) SetAdditionalEntryInformation(value string) {
	r.AdditionalEntryInformation = (*Max500Text)(&value)
}
