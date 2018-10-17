package iso20022

// Details of the transaction in the authorisation request.
type CardPaymentTransactionDetails1 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Qualifies the amount associated with the transaction.
	AmountQualifier *TypeOfAmount1Code `xml:"AmtQlfr,omitempty"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount []*DetailedAmount1 `xml:"DtldAmt,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Reason to process an online authorisation.
	OnLineReason *OnLineReason1Code `xml:"OnLineRsn,omitempty"`

	// Transaction category level on an unattended POI (Point Of Interaction).
	UnattendedLevelCategory *Max35NumericText `xml:"UattnddLvlCtgy,omitempty"`

	// Type of cardholder account used for the transaction.
	AccountType *CardAccountType1Code `xml:"AcctTp,omitempty"`

	// Data related to a financial loan (instalment) or to a recurring transaction.
	RecurringTransaction *RecurringTransaction1 `xml:"RcrngTx,omitempty"`

	// Product purchased with the transaction.
	Product []*Product1 `xml:"Pdct,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails1) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails1) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails1) SetAmountQualifier(value string) {
	c.AmountQualifier = (*TypeOfAmount1Code)(&value)
}

func (c *CardPaymentTransactionDetails1) AddDetailedAmount() *DetailedAmount1 {
	newValue := new(DetailedAmount1)
	c.DetailedAmount = append(c.DetailedAmount, newValue)
	return newValue
}

func (c *CardPaymentTransactionDetails1) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails1) SetOnLineReason(value string) {
	c.OnLineReason = (*OnLineReason1Code)(&value)
}

func (c *CardPaymentTransactionDetails1) SetUnattendedLevelCategory(value string) {
	c.UnattendedLevelCategory = (*Max35NumericText)(&value)
}

func (c *CardPaymentTransactionDetails1) SetAccountType(value string) {
	c.AccountType = (*CardAccountType1Code)(&value)
}

func (c *CardPaymentTransactionDetails1) AddRecurringTransaction() *RecurringTransaction1 {
	c.RecurringTransaction = new(RecurringTransaction1)
	return c.RecurringTransaction
}

func (c *CardPaymentTransactionDetails1) AddProduct() *Product1 {
	newValue := new(Product1)
	c.Product = append(c.Product, newValue)
	return newValue
}

func (c *CardPaymentTransactionDetails1) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}

// Details of the transaction in the authorisation request in a batch.
type CardPaymentTransactionDetails10 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Qualifies the amount associated with the transaction.
	AmountQualifier *TypeOfAmount1Code `xml:"AmtQlfr,omitempty"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount []*DetailedAmount1 `xml:"DtldAmt,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Reason to process an online authorisation.
	OnLineReason *OnLineReason1Code `xml:"OnLineRsn,omitempty"`

	// Transaction category level on an unattended POI (Point Of Interaction).
	UnattendedLevelCategory *Max35NumericText `xml:"UattnddLvlCtgy,omitempty"`

	// Type of cardholder account used for the transaction.
	AccountType *CardAccountType1Code `xml:"AcctTp,omitempty"`

	// Data related to a financial loan (instalment) or to a recurring transaction.
	RecurringTransaction *RecurringTransaction1 `xml:"RcrngTx,omitempty"`

	// Product purchased with the transaction.
	Product []*Product1 `xml:"Pdct,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails10) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails10) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails10) SetAmountQualifier(value string) {
	c.AmountQualifier = (*TypeOfAmount1Code)(&value)
}

func (c *CardPaymentTransactionDetails10) AddDetailedAmount() *DetailedAmount1 {
	newValue := new(DetailedAmount1)
	c.DetailedAmount = append(c.DetailedAmount, newValue)
	return newValue
}

func (c *CardPaymentTransactionDetails10) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails10) SetOnLineReason(value string) {
	c.OnLineReason = (*OnLineReason1Code)(&value)
}

func (c *CardPaymentTransactionDetails10) SetUnattendedLevelCategory(value string) {
	c.UnattendedLevelCategory = (*Max35NumericText)(&value)
}

func (c *CardPaymentTransactionDetails10) SetAccountType(value string) {
	c.AccountType = (*CardAccountType1Code)(&value)
}

func (c *CardPaymentTransactionDetails10) AddRecurringTransaction() *RecurringTransaction1 {
	c.RecurringTransaction = new(RecurringTransaction1)
	return c.RecurringTransaction
}

func (c *CardPaymentTransactionDetails10) AddProduct() *Product1 {
	newValue := new(Product1)
	c.Product = append(c.Product, newValue)
	return newValue
}

func (c *CardPaymentTransactionDetails10) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}

// Details of the transaction in the authorisation response contained in a batch.
type CardPaymentTransactionDetails11 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount []*DetailedAmount2 `xml:"DtldAmt,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Type of cardholder account used for the transaction.
	AccountType *CardAccountType1Code `xml:"AcctTp,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails11) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails11) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails11) AddDetailedAmount() *DetailedAmount2 {
	newValue := new(DetailedAmount2)
	c.DetailedAmount = append(c.DetailedAmount, newValue)
	return newValue
}

func (c *CardPaymentTransactionDetails11) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails11) SetAccountType(value string) {
	c.AccountType = (*CardAccountType1Code)(&value)
}

func (c *CardPaymentTransactionDetails11) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}

// Details of the transaction in the authorisation request.
type CardPaymentTransactionDetails12 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Qualifies the amount associated with the transaction.
	AmountQualifier *TypeOfAmount1Code `xml:"AmtQlfr,omitempty"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount *DetailedAmount5 `xml:"DtldAmt,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Reason to process an online authorisation.
	OnLineReason *OnLineReason1Code `xml:"OnLineRsn,omitempty"`

	// Transaction category level on an unattended POI (Point Of Interaction).
	UnattendedLevelCategory *Max35NumericText `xml:"UattnddLvlCtgy,omitempty"`

	// Type of cardholder account used for the transaction.
	AccountType *CardAccountType1Code `xml:"AcctTp,omitempty"`

	// The currency conversion has been accepted by the cardholder.
	ConversionAccepted *TrueFalseIndicator `xml:"ConvsAccptd,omitempty"`

	// Currency conversion proposed to the cardholder.
	CurrencyConversion *CurrencyConversion1 `xml:"CcyConvs,omitempty"`

	// Data related to a financial loan (instalment) or to a recurring transaction.
	Instalment *RecurringTransaction2 `xml:"Instlmt,omitempty"`

	// Payment transaction with an aggregated amount.
	AggregationTransaction *AggregationTransaction1 `xml:"AggtnTx,omitempty"`

	// Product purchased with the transaction.
	Product []*Product1 `xml:"Pdct,omitempty"`

	// Detailed invoice data.
	CardPaymentInvoice *CardPaymentInvoice1 `xml:"CardPmtInvc,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails12) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails12) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails12) SetAmountQualifier(value string) {
	c.AmountQualifier = (*TypeOfAmount1Code)(&value)
}

func (c *CardPaymentTransactionDetails12) AddDetailedAmount() *DetailedAmount5 {
	c.DetailedAmount = new(DetailedAmount5)
	return c.DetailedAmount
}

func (c *CardPaymentTransactionDetails12) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails12) SetOnLineReason(value string) {
	c.OnLineReason = (*OnLineReason1Code)(&value)
}

func (c *CardPaymentTransactionDetails12) SetUnattendedLevelCategory(value string) {
	c.UnattendedLevelCategory = (*Max35NumericText)(&value)
}

func (c *CardPaymentTransactionDetails12) SetAccountType(value string) {
	c.AccountType = (*CardAccountType1Code)(&value)
}

func (c *CardPaymentTransactionDetails12) SetConversionAccepted(value string) {
	c.ConversionAccepted = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransactionDetails12) AddCurrencyConversion() *CurrencyConversion1 {
	c.CurrencyConversion = new(CurrencyConversion1)
	return c.CurrencyConversion
}

func (c *CardPaymentTransactionDetails12) AddInstalment() *RecurringTransaction2 {
	c.Instalment = new(RecurringTransaction2)
	return c.Instalment
}

func (c *CardPaymentTransactionDetails12) AddAggregationTransaction() *AggregationTransaction1 {
	c.AggregationTransaction = new(AggregationTransaction1)
	return c.AggregationTransaction
}

func (c *CardPaymentTransactionDetails12) AddProduct() *Product1 {
	newValue := new(Product1)
	c.Product = append(c.Product, newValue)
	return newValue
}

func (c *CardPaymentTransactionDetails12) AddCardPaymentInvoice() *CardPaymentInvoice1 {
	c.CardPaymentInvoice = new(CardPaymentInvoice1)
	return c.CardPaymentInvoice
}

func (c *CardPaymentTransactionDetails12) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}

// Details of the transaction in the authorisation response.
type CardPaymentTransactionDetails13 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount *DetailedAmount5 `xml:"DtldAmt,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Type of cardholder account used for the transaction.
	AccountType *CardAccountType1Code `xml:"AcctTp,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails13) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails13) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails13) AddDetailedAmount() *DetailedAmount5 {
	c.DetailedAmount = new(DetailedAmount5)
	return c.DetailedAmount
}

func (c *CardPaymentTransactionDetails13) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails13) SetAccountType(value string) {
	c.AccountType = (*CardAccountType1Code)(&value)
}

func (c *CardPaymentTransactionDetails13) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}

// Details of the transaction in the completion advice.
type CardPaymentTransactionDetails14 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Qualifies the amount associated with the transaction.
	AmountQualifier *TypeOfAmount1Code `xml:"AmtQlfr,omitempty"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount *DetailedAmount5 `xml:"DtldAmt,omitempty"`

	// Amount requested to be authorised.
	RequestedAmount *ImpliedCurrencyAndAmount `xml:"ReqdAmt,omitempty"`

	// Amount authorised for the payment transaction.
	AuthorisedAmount *ImpliedCurrencyAndAmount `xml:"AuthrsdAmt,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Transaction category level on an unattended POI (Point Of Interaction).
	UnattendedLevelCategory *Max35NumericText `xml:"UattnddLvlCtgy,omitempty"`

	// Type of cardholder account used for the transaction.
	AccountType *CardAccountType1Code `xml:"AcctTp,omitempty"`

	// Indicates if a currency conversion has been accepted by the cardholder.
	ConversionAccepted *TrueFalseIndicator `xml:"ConvsAccptd,omitempty"`

	// Currency conversion proposed to the cardholder.
	CurrencyConversion *CurrencyConversion1 `xml:"CcyConvs,omitempty"`

	// Data related to a financial loan (instalment) or to a recurring transaction.
	Instalment *RecurringTransaction2 `xml:"Instlmt,omitempty"`

	// Product purchased with the transaction.
	Product []*Product1 `xml:"Pdct,omitempty"`

	// Detailed invoice data.
	CardPaymentInvoice *CardPaymentInvoice1 `xml:"CardPmtInvc,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails14) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails14) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails14) SetAmountQualifier(value string) {
	c.AmountQualifier = (*TypeOfAmount1Code)(&value)
}

func (c *CardPaymentTransactionDetails14) AddDetailedAmount() *DetailedAmount5 {
	c.DetailedAmount = new(DetailedAmount5)
	return c.DetailedAmount
}

func (c *CardPaymentTransactionDetails14) SetRequestedAmount(value, currency string) {
	c.RequestedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails14) SetAuthorisedAmount(value, currency string) {
	c.AuthorisedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails14) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails14) SetUnattendedLevelCategory(value string) {
	c.UnattendedLevelCategory = (*Max35NumericText)(&value)
}

func (c *CardPaymentTransactionDetails14) SetAccountType(value string) {
	c.AccountType = (*CardAccountType1Code)(&value)
}

func (c *CardPaymentTransactionDetails14) SetConversionAccepted(value string) {
	c.ConversionAccepted = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransactionDetails14) AddCurrencyConversion() *CurrencyConversion1 {
	c.CurrencyConversion = new(CurrencyConversion1)
	return c.CurrencyConversion
}

func (c *CardPaymentTransactionDetails14) AddInstalment() *RecurringTransaction2 {
	c.Instalment = new(RecurringTransaction2)
	return c.Instalment
}

func (c *CardPaymentTransactionDetails14) AddProduct() *Product1 {
	newValue := new(Product1)
	c.Product = append(c.Product, newValue)
	return newValue
}

func (c *CardPaymentTransactionDetails14) AddCardPaymentInvoice() *CardPaymentInvoice1 {
	c.CardPaymentInvoice = new(CardPaymentInvoice1)
	return c.CardPaymentInvoice
}

func (c *CardPaymentTransactionDetails14) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}

// Details of the transaction to capture.
type CardPaymentTransactionDetails15 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Qualifies the amount associated with the transaction.
	AmountQualifier *TypeOfAmount1Code `xml:"AmtQlfr,omitempty"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount *DetailedAmount5 `xml:"DtldAmt,omitempty"`

	// Amount requested to be authorised.
	RequestedAmount *ImpliedCurrencyAndAmount `xml:"ReqdAmt,omitempty"`

	// Amount authorised for the transaction.
	AuthorisedAmount *ImpliedCurrencyAndAmount `xml:"AuthrsdAmt,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Transaction category level on an unattended POI (Point Of Interaction).
	UnattendedLevelCategory *Max35NumericText `xml:"UattnddLvlCtgy,omitempty"`

	// Type of cardholder account used for the transaction.
	AccountType *CardAccountType1Code `xml:"AcctTp,omitempty"`

	// Indicates if the currency conversion has been accepted by the cardholder.
	ConversionAccepted *TrueFalseIndicator `xml:"ConvsAccptd,omitempty"`

	// Currency conversion proposed to the cardholder.
	CurrencyConversion *CurrencyConversion1 `xml:"CcyConvs,omitempty"`

	// Data related to a financial loan (instalment) or to a recurring transaction.
	Instalment *RecurringTransaction2 `xml:"Instlmt,omitempty"`

	// Payment transaction with an aggregated amount.
	AggregationTransaction *AggregationTransaction1 `xml:"AggtnTx,omitempty"`

	// Product purchased with the transaction.
	Product []*Product1 `xml:"Pdct,omitempty"`

	// Detailed invoice data.
	CardPaymentInvoice *CardPaymentInvoice1 `xml:"CardPmtInvc,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails15) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails15) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails15) SetAmountQualifier(value string) {
	c.AmountQualifier = (*TypeOfAmount1Code)(&value)
}

func (c *CardPaymentTransactionDetails15) AddDetailedAmount() *DetailedAmount5 {
	c.DetailedAmount = new(DetailedAmount5)
	return c.DetailedAmount
}

func (c *CardPaymentTransactionDetails15) SetRequestedAmount(value, currency string) {
	c.RequestedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails15) SetAuthorisedAmount(value, currency string) {
	c.AuthorisedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails15) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails15) SetUnattendedLevelCategory(value string) {
	c.UnattendedLevelCategory = (*Max35NumericText)(&value)
}

func (c *CardPaymentTransactionDetails15) SetAccountType(value string) {
	c.AccountType = (*CardAccountType1Code)(&value)
}

func (c *CardPaymentTransactionDetails15) SetConversionAccepted(value string) {
	c.ConversionAccepted = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransactionDetails15) AddCurrencyConversion() *CurrencyConversion1 {
	c.CurrencyConversion = new(CurrencyConversion1)
	return c.CurrencyConversion
}

func (c *CardPaymentTransactionDetails15) AddInstalment() *RecurringTransaction2 {
	c.Instalment = new(RecurringTransaction2)
	return c.Instalment
}

func (c *CardPaymentTransactionDetails15) AddAggregationTransaction() *AggregationTransaction1 {
	c.AggregationTransaction = new(AggregationTransaction1)
	return c.AggregationTransaction
}

func (c *CardPaymentTransactionDetails15) AddProduct() *Product1 {
	newValue := new(Product1)
	c.Product = append(c.Product, newValue)
	return newValue
}

func (c *CardPaymentTransactionDetails15) AddCardPaymentInvoice() *CardPaymentInvoice1 {
	c.CardPaymentInvoice = new(CardPaymentInvoice1)
	return c.CardPaymentInvoice
}

func (c *CardPaymentTransactionDetails15) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}

// Details of the transaction in the authorisation request in a batch.
type CardPaymentTransactionDetails16 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Qualifies the amount associated with the transaction.
	AmountQualifier *TypeOfAmount1Code `xml:"AmtQlfr,omitempty"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount *DetailedAmount5 `xml:"DtldAmt,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Reason to process an online authorisation.
	OnLineReason *OnLineReason1Code `xml:"OnLineRsn,omitempty"`

	// Transaction category level on an unattended POI (Point Of Interaction).
	UnattendedLevelCategory *Max35NumericText `xml:"UattnddLvlCtgy,omitempty"`

	// Type of cardholder account used for the transaction.
	AccountType *CardAccountType1Code `xml:"AcctTp,omitempty"`

	// Data related to a financial loan (instalment) or to a recurring transaction.
	Instalment *RecurringTransaction2 `xml:"Instlmt,omitempty"`

	// Payment transaction with an aggregated amount.
	AggregationTransaction *AggregationTransaction1 `xml:"AggtnTx,omitempty"`

	// Product purchased with the transaction.
	Product []*Product1 `xml:"Pdct,omitempty"`

	// Detailed invoice data.
	CardPaymentInvoice *CardPaymentInvoice1 `xml:"CardPmtInvc,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails16) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails16) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails16) SetAmountQualifier(value string) {
	c.AmountQualifier = (*TypeOfAmount1Code)(&value)
}

func (c *CardPaymentTransactionDetails16) AddDetailedAmount() *DetailedAmount5 {
	c.DetailedAmount = new(DetailedAmount5)
	return c.DetailedAmount
}

func (c *CardPaymentTransactionDetails16) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails16) SetOnLineReason(value string) {
	c.OnLineReason = (*OnLineReason1Code)(&value)
}

func (c *CardPaymentTransactionDetails16) SetUnattendedLevelCategory(value string) {
	c.UnattendedLevelCategory = (*Max35NumericText)(&value)
}

func (c *CardPaymentTransactionDetails16) SetAccountType(value string) {
	c.AccountType = (*CardAccountType1Code)(&value)
}

func (c *CardPaymentTransactionDetails16) AddInstalment() *RecurringTransaction2 {
	c.Instalment = new(RecurringTransaction2)
	return c.Instalment
}

func (c *CardPaymentTransactionDetails16) AddAggregationTransaction() *AggregationTransaction1 {
	c.AggregationTransaction = new(AggregationTransaction1)
	return c.AggregationTransaction
}

func (c *CardPaymentTransactionDetails16) AddProduct() *Product1 {
	newValue := new(Product1)
	c.Product = append(c.Product, newValue)
	return newValue
}

func (c *CardPaymentTransactionDetails16) AddCardPaymentInvoice() *CardPaymentInvoice1 {
	c.CardPaymentInvoice = new(CardPaymentInvoice1)
	return c.CardPaymentInvoice
}

func (c *CardPaymentTransactionDetails16) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}

// Details of the transaction in the authorisation response contained in a batch.
type CardPaymentTransactionDetails17 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount *DetailedAmount5 `xml:"DtldAmt,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Type of cardholder account used for the transaction.
	AccountType *CardAccountType1Code `xml:"AcctTp,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails17) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails17) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails17) AddDetailedAmount() *DetailedAmount5 {
	c.DetailedAmount = new(DetailedAmount5)
	return c.DetailedAmount
}

func (c *CardPaymentTransactionDetails17) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails17) SetAccountType(value string) {
	c.AccountType = (*CardAccountType1Code)(&value)
}

func (c *CardPaymentTransactionDetails17) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}

// Details of the payment transaction.
type CardPaymentTransactionDetails18 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Qualifies the amount associated with the transaction.
	AmountQualifier *TypeOfAmount1Code `xml:"AmtQlfr,omitempty"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount *DetailedAmount5 `xml:"DtldAmt,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails18) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails18) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails18) SetAmountQualifier(value string) {
	c.AmountQualifier = (*TypeOfAmount1Code)(&value)
}

func (c *CardPaymentTransactionDetails18) AddDetailedAmount() *DetailedAmount5 {
	c.DetailedAmount = new(DetailedAmount5)
	return c.DetailedAmount
}

func (c *CardPaymentTransactionDetails18) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}

// Details of the transaction in the authorisation request.
type CardPaymentTransactionDetails19 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Qualifies the amount associated with the transaction.
	AmountQualifier *TypeOfAmount1Code `xml:"AmtQlfr,omitempty"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount *DetailedAmount7 `xml:"DtldAmt,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Reason to process an online authorisation.
	OnLineReason *OnLineReason1Code `xml:"OnLineRsn,omitempty"`

	// Transaction category level on an unattended POI (Point Of Interaction).
	UnattendedLevelCategory *Max35NumericText `xml:"UattnddLvlCtgy,omitempty"`

	// Type of cardholder account used for the transaction.
	AccountType *CardAccountType2Code `xml:"AcctTp,omitempty"`

	// Currency conversion proposed to the cardholder.
	CurrencyConversion *CurrencyConversion3 `xml:"CcyConvs,omitempty"`

	// Data related to a financial loan (instalment) or to a recurring transaction.
	Instalment *RecurringTransaction2 `xml:"Instlmt,omitempty"`

	// Payment transaction with an aggregated amount.
	AggregationTransaction *AggregationTransaction1 `xml:"AggtnTx,omitempty"`

	// Product purchased with the transaction.
	Product []*Product1 `xml:"Pdct,omitempty"`

	// Detailed invoice data.
	CardPaymentInvoice *CardPaymentInvoice1 `xml:"CardPmtInvc,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails19) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails19) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails19) SetAmountQualifier(value string) {
	c.AmountQualifier = (*TypeOfAmount1Code)(&value)
}

func (c *CardPaymentTransactionDetails19) AddDetailedAmount() *DetailedAmount7 {
	c.DetailedAmount = new(DetailedAmount7)
	return c.DetailedAmount
}

func (c *CardPaymentTransactionDetails19) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails19) SetOnLineReason(value string) {
	c.OnLineReason = (*OnLineReason1Code)(&value)
}

func (c *CardPaymentTransactionDetails19) SetUnattendedLevelCategory(value string) {
	c.UnattendedLevelCategory = (*Max35NumericText)(&value)
}

func (c *CardPaymentTransactionDetails19) SetAccountType(value string) {
	c.AccountType = (*CardAccountType2Code)(&value)
}

func (c *CardPaymentTransactionDetails19) AddCurrencyConversion() *CurrencyConversion3 {
	c.CurrencyConversion = new(CurrencyConversion3)
	return c.CurrencyConversion
}

func (c *CardPaymentTransactionDetails19) AddInstalment() *RecurringTransaction2 {
	c.Instalment = new(RecurringTransaction2)
	return c.Instalment
}

func (c *CardPaymentTransactionDetails19) AddAggregationTransaction() *AggregationTransaction1 {
	c.AggregationTransaction = new(AggregationTransaction1)
	return c.AggregationTransaction
}

func (c *CardPaymentTransactionDetails19) AddProduct() *Product1 {
	newValue := new(Product1)
	c.Product = append(c.Product, newValue)
	return newValue
}

func (c *CardPaymentTransactionDetails19) AddCardPaymentInvoice() *CardPaymentInvoice1 {
	c.CardPaymentInvoice = new(CardPaymentInvoice1)
	return c.CardPaymentInvoice
}

func (c *CardPaymentTransactionDetails19) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}

// Details of the transaction in the authorisation response.
type CardPaymentTransactionDetails2 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount []*DetailedAmount2 `xml:"DtldAmt,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Type of cardholder account used for the transaction.
	AccountType *CardAccountType1Code `xml:"AcctTp,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails2) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails2) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails2) AddDetailedAmount() *DetailedAmount2 {
	newValue := new(DetailedAmount2)
	c.DetailedAmount = append(c.DetailedAmount, newValue)
	return newValue
}

func (c *CardPaymentTransactionDetails2) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails2) SetAccountType(value string) {
	c.AccountType = (*CardAccountType1Code)(&value)
}

func (c *CardPaymentTransactionDetails2) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}

// Details of the transaction in the authorisation response.
type CardPaymentTransactionDetails20 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount *DetailedAmount7 `xml:"DtldAmt,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Type of cardholder account used for the transaction.
	AccountType *CardAccountType2Code `xml:"AcctTp,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails20) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails20) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails20) AddDetailedAmount() *DetailedAmount7 {
	c.DetailedAmount = new(DetailedAmount7)
	return c.DetailedAmount
}

func (c *CardPaymentTransactionDetails20) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails20) SetAccountType(value string) {
	c.AccountType = (*CardAccountType2Code)(&value)
}

func (c *CardPaymentTransactionDetails20) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}

// Details of the transaction in the completion advice.
type CardPaymentTransactionDetails21 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Qualifies the amount associated with the transaction.
	AmountQualifier *TypeOfAmount1Code `xml:"AmtQlfr,omitempty"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount *DetailedAmount7 `xml:"DtldAmt,omitempty"`

	// Amount requested to be authorised.
	RequestedAmount *ImpliedCurrencyAndAmount `xml:"ReqdAmt,omitempty"`

	// Amount authorised for the payment transaction.
	AuthorisedAmount *ImpliedCurrencyAndAmount `xml:"AuthrsdAmt,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Transaction category level on an unattended POI (Point Of Interaction).
	UnattendedLevelCategory *Max35NumericText `xml:"UattnddLvlCtgy,omitempty"`

	// Type of cardholder account used for the transaction.
	AccountType *CardAccountType2Code `xml:"AcctTp,omitempty"`

	// Currency conversion proposed to the cardholder.
	CurrencyConversion *CurrencyConversion3 `xml:"CcyConvs,omitempty"`

	// Data related to a financial loan (instalment) or to a recurring transaction.
	Instalment *RecurringTransaction2 `xml:"Instlmt,omitempty"`

	// Payment transaction with an aggregated amount.
	AggregationTransaction *AggregationTransaction1 `xml:"AggtnTx,omitempty"`

	// Product purchased with the transaction.
	Product []*Product1 `xml:"Pdct,omitempty"`

	// Detailed invoice data.
	CardPaymentInvoice *CardPaymentInvoice1 `xml:"CardPmtInvc,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails21) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails21) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails21) SetAmountQualifier(value string) {
	c.AmountQualifier = (*TypeOfAmount1Code)(&value)
}

func (c *CardPaymentTransactionDetails21) AddDetailedAmount() *DetailedAmount7 {
	c.DetailedAmount = new(DetailedAmount7)
	return c.DetailedAmount
}

func (c *CardPaymentTransactionDetails21) SetRequestedAmount(value, currency string) {
	c.RequestedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails21) SetAuthorisedAmount(value, currency string) {
	c.AuthorisedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails21) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails21) SetUnattendedLevelCategory(value string) {
	c.UnattendedLevelCategory = (*Max35NumericText)(&value)
}

func (c *CardPaymentTransactionDetails21) SetAccountType(value string) {
	c.AccountType = (*CardAccountType2Code)(&value)
}

func (c *CardPaymentTransactionDetails21) AddCurrencyConversion() *CurrencyConversion3 {
	c.CurrencyConversion = new(CurrencyConversion3)
	return c.CurrencyConversion
}

func (c *CardPaymentTransactionDetails21) AddInstalment() *RecurringTransaction2 {
	c.Instalment = new(RecurringTransaction2)
	return c.Instalment
}

func (c *CardPaymentTransactionDetails21) AddAggregationTransaction() *AggregationTransaction1 {
	c.AggregationTransaction = new(AggregationTransaction1)
	return c.AggregationTransaction
}

func (c *CardPaymentTransactionDetails21) AddProduct() *Product1 {
	newValue := new(Product1)
	c.Product = append(c.Product, newValue)
	return newValue
}

func (c *CardPaymentTransactionDetails21) AddCardPaymentInvoice() *CardPaymentInvoice1 {
	c.CardPaymentInvoice = new(CardPaymentInvoice1)
	return c.CardPaymentInvoice
}

func (c *CardPaymentTransactionDetails21) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}

// Details of the transaction in the cancellation response.
type CardPaymentTransactionDetails22 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails22) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails22) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails22) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}

// Details of the transaction to capture.
type CardPaymentTransactionDetails23 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Qualifies the amount associated with the transaction.
	AmountQualifier *TypeOfAmount1Code `xml:"AmtQlfr,omitempty"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount *DetailedAmount7 `xml:"DtldAmt,omitempty"`

	// Amount requested to be authorised.
	RequestedAmount *ImpliedCurrencyAndAmount `xml:"ReqdAmt,omitempty"`

	// Amount authorised for the transaction.
	AuthorisedAmount *ImpliedCurrencyAndAmount `xml:"AuthrsdAmt,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Transaction category level on an unattended POI (Point Of Interaction).
	UnattendedLevelCategory *Max35NumericText `xml:"UattnddLvlCtgy,omitempty"`

	// Type of cardholder account used for the transaction.
	AccountType *CardAccountType2Code `xml:"AcctTp,omitempty"`

	// Currency conversion proposed to the cardholder.
	CurrencyConversion *CurrencyConversion3 `xml:"CcyConvs,omitempty"`

	// Data related to a financial loan (instalment) or to a recurring transaction.
	Instalment *RecurringTransaction2 `xml:"Instlmt,omitempty"`

	// Payment transaction with an aggregated amount.
	AggregationTransaction *AggregationTransaction1 `xml:"AggtnTx,omitempty"`

	// Product purchased with the transaction.
	Product []*Product1 `xml:"Pdct,omitempty"`

	// Detailed invoice data.
	CardPaymentInvoice *CardPaymentInvoice1 `xml:"CardPmtInvc,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails23) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails23) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails23) SetAmountQualifier(value string) {
	c.AmountQualifier = (*TypeOfAmount1Code)(&value)
}

func (c *CardPaymentTransactionDetails23) AddDetailedAmount() *DetailedAmount7 {
	c.DetailedAmount = new(DetailedAmount7)
	return c.DetailedAmount
}

func (c *CardPaymentTransactionDetails23) SetRequestedAmount(value, currency string) {
	c.RequestedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails23) SetAuthorisedAmount(value, currency string) {
	c.AuthorisedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails23) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails23) SetUnattendedLevelCategory(value string) {
	c.UnattendedLevelCategory = (*Max35NumericText)(&value)
}

func (c *CardPaymentTransactionDetails23) SetAccountType(value string) {
	c.AccountType = (*CardAccountType2Code)(&value)
}

func (c *CardPaymentTransactionDetails23) AddCurrencyConversion() *CurrencyConversion3 {
	c.CurrencyConversion = new(CurrencyConversion3)
	return c.CurrencyConversion
}

func (c *CardPaymentTransactionDetails23) AddInstalment() *RecurringTransaction2 {
	c.Instalment = new(RecurringTransaction2)
	return c.Instalment
}

func (c *CardPaymentTransactionDetails23) AddAggregationTransaction() *AggregationTransaction1 {
	c.AggregationTransaction = new(AggregationTransaction1)
	return c.AggregationTransaction
}

func (c *CardPaymentTransactionDetails23) AddProduct() *Product1 {
	newValue := new(Product1)
	c.Product = append(c.Product, newValue)
	return newValue
}

func (c *CardPaymentTransactionDetails23) AddCardPaymentInvoice() *CardPaymentInvoice1 {
	c.CardPaymentInvoice = new(CardPaymentInvoice1)
	return c.CardPaymentInvoice
}

func (c *CardPaymentTransactionDetails23) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}

// Details of the transaction in the authorisation request in a batch.
type CardPaymentTransactionDetails24 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Qualifies the amount associated with the transaction.
	AmountQualifier *TypeOfAmount1Code `xml:"AmtQlfr,omitempty"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount *DetailedAmount7 `xml:"DtldAmt,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Reason to process an online authorisation.
	OnLineReason *OnLineReason1Code `xml:"OnLineRsn,omitempty"`

	// Transaction category level on an unattended POI (Point Of Interaction).
	UnattendedLevelCategory *Max35NumericText `xml:"UattnddLvlCtgy,omitempty"`

	// Type of cardholder account used for the transaction.
	AccountType *CardAccountType2Code `xml:"AcctTp,omitempty"`

	// Data related to a financial loan (instalment) or to a recurring transaction.
	Instalment *RecurringTransaction2 `xml:"Instlmt,omitempty"`

	// Payment transaction with an aggregated amount.
	AggregationTransaction *AggregationTransaction1 `xml:"AggtnTx,omitempty"`

	// Product purchased with the transaction.
	Product []*Product1 `xml:"Pdct,omitempty"`

	// Detailed invoice data.
	CardPaymentInvoice *CardPaymentInvoice1 `xml:"CardPmtInvc,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails24) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails24) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails24) SetAmountQualifier(value string) {
	c.AmountQualifier = (*TypeOfAmount1Code)(&value)
}

func (c *CardPaymentTransactionDetails24) AddDetailedAmount() *DetailedAmount7 {
	c.DetailedAmount = new(DetailedAmount7)
	return c.DetailedAmount
}

func (c *CardPaymentTransactionDetails24) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails24) SetOnLineReason(value string) {
	c.OnLineReason = (*OnLineReason1Code)(&value)
}

func (c *CardPaymentTransactionDetails24) SetUnattendedLevelCategory(value string) {
	c.UnattendedLevelCategory = (*Max35NumericText)(&value)
}

func (c *CardPaymentTransactionDetails24) SetAccountType(value string) {
	c.AccountType = (*CardAccountType2Code)(&value)
}

func (c *CardPaymentTransactionDetails24) AddInstalment() *RecurringTransaction2 {
	c.Instalment = new(RecurringTransaction2)
	return c.Instalment
}

func (c *CardPaymentTransactionDetails24) AddAggregationTransaction() *AggregationTransaction1 {
	c.AggregationTransaction = new(AggregationTransaction1)
	return c.AggregationTransaction
}

func (c *CardPaymentTransactionDetails24) AddProduct() *Product1 {
	newValue := new(Product1)
	c.Product = append(c.Product, newValue)
	return newValue
}

func (c *CardPaymentTransactionDetails24) AddCardPaymentInvoice() *CardPaymentInvoice1 {
	c.CardPaymentInvoice = new(CardPaymentInvoice1)
	return c.CardPaymentInvoice
}

func (c *CardPaymentTransactionDetails24) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}

// Details of the transaction in the authorisation response contained in a batch.
type CardPaymentTransactionDetails25 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount *DetailedAmount7 `xml:"DtldAmt,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Type of cardholder account used for the transaction.
	AccountType *CardAccountType2Code `xml:"AcctTp,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails25) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails25) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails25) AddDetailedAmount() *DetailedAmount7 {
	c.DetailedAmount = new(DetailedAmount7)
	return c.DetailedAmount
}

func (c *CardPaymentTransactionDetails25) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails25) SetAccountType(value string) {
	c.AccountType = (*CardAccountType2Code)(&value)
}

func (c *CardPaymentTransactionDetails25) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}

// Details of the payment transaction.
type CardPaymentTransactionDetails26 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Qualifies the amount associated with the transaction.
	AmountQualifier *TypeOfAmount1Code `xml:"AmtQlfr,omitempty"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount *DetailedAmount7 `xml:"DtldAmt,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails26) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails26) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails26) SetAmountQualifier(value string) {
	c.AmountQualifier = (*TypeOfAmount1Code)(&value)
}

func (c *CardPaymentTransactionDetails26) AddDetailedAmount() *DetailedAmount7 {
	c.DetailedAmount = new(DetailedAmount7)
	return c.DetailedAmount
}

func (c *CardPaymentTransactionDetails26) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}

// Details of the transaction in the authorisation request.
type CardPaymentTransactionDetails27 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Qualifies the amount associated with the transaction.
	AmountQualifier *TypeOfAmount1Code `xml:"AmtQlfr,omitempty"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount *DetailedAmount15 `xml:"DtldAmt,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Reason to process an online authorisation.
	OnLineReason *OnLineReason1Code `xml:"OnLineRsn,omitempty"`

	// Transaction category level on an unattended POI (Point Of Interaction).
	UnattendedLevelCategory *Max35NumericText `xml:"UattnddLvlCtgy,omitempty"`

	// Type of cardholder account used for the transaction.
	AccountType *CardAccountType3Code `xml:"AcctTp,omitempty"`

	// Result of the currency conversion proposed to the cardholder.
	CurrencyConversionResult *CurrencyConversion8 `xml:"CcyConvsRslt,omitempty"`

	// Data related to a financial loan (instalment) or to a recurring transaction.
	Instalment *RecurringTransaction2 `xml:"Instlmt,omitempty"`

	// Payment transaction with an aggregated amount.
	AggregationTransaction *AggregationTransaction2 `xml:"AggtnTx,omitempty"`

	// Codification used to identify the products.
	ProductCodeSetIdentification *Max10Text `xml:"PdctCdSetId,omitempty"`

	// Item purchased with the transaction.
	SaleItem []*Product3 `xml:"SaleItm,omitempty"`

	// Location of the delivery, for instance pump number or parking bay.
	DeliveryLocation *Max10Text `xml:"DlvryLctn,omitempty"`

	// Detailed invoice data.
	CardPaymentInvoice *CardPaymentInvoice2 `xml:"CardPmtInvc,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails27) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails27) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails27) SetAmountQualifier(value string) {
	c.AmountQualifier = (*TypeOfAmount1Code)(&value)
}

func (c *CardPaymentTransactionDetails27) AddDetailedAmount() *DetailedAmount15 {
	c.DetailedAmount = new(DetailedAmount15)
	return c.DetailedAmount
}

func (c *CardPaymentTransactionDetails27) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails27) SetOnLineReason(value string) {
	c.OnLineReason = (*OnLineReason1Code)(&value)
}

func (c *CardPaymentTransactionDetails27) SetUnattendedLevelCategory(value string) {
	c.UnattendedLevelCategory = (*Max35NumericText)(&value)
}

func (c *CardPaymentTransactionDetails27) SetAccountType(value string) {
	c.AccountType = (*CardAccountType3Code)(&value)
}

func (c *CardPaymentTransactionDetails27) AddCurrencyConversionResult() *CurrencyConversion8 {
	c.CurrencyConversionResult = new(CurrencyConversion8)
	return c.CurrencyConversionResult
}

func (c *CardPaymentTransactionDetails27) AddInstalment() *RecurringTransaction2 {
	c.Instalment = new(RecurringTransaction2)
	return c.Instalment
}

func (c *CardPaymentTransactionDetails27) AddAggregationTransaction() *AggregationTransaction2 {
	c.AggregationTransaction = new(AggregationTransaction2)
	return c.AggregationTransaction
}

func (c *CardPaymentTransactionDetails27) SetProductCodeSetIdentification(value string) {
	c.ProductCodeSetIdentification = (*Max10Text)(&value)
}

func (c *CardPaymentTransactionDetails27) AddSaleItem() *Product3 {
	newValue := new(Product3)
	c.SaleItem = append(c.SaleItem, newValue)
	return newValue
}

func (c *CardPaymentTransactionDetails27) SetDeliveryLocation(value string) {
	c.DeliveryLocation = (*Max10Text)(&value)
}

func (c *CardPaymentTransactionDetails27) AddCardPaymentInvoice() *CardPaymentInvoice2 {
	c.CardPaymentInvoice = new(CardPaymentInvoice2)
	return c.CardPaymentInvoice
}

func (c *CardPaymentTransactionDetails27) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}

// Details of the transaction in the authorisation response.
type CardPaymentTransactionDetails28 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount *DetailedAmount15 `xml:"DtldAmt,omitempty"`

	// Amount of the transaction that will be invoiced to the cardholder.
	InvoiceAmount *DetailedAmount4 `xml:"InvcAmt,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Type of cardholder account used for the transaction.
	AccountType *CardAccountType3Code `xml:"AcctTp,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails28) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails28) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails28) AddDetailedAmount() *DetailedAmount15 {
	c.DetailedAmount = new(DetailedAmount15)
	return c.DetailedAmount
}

func (c *CardPaymentTransactionDetails28) AddInvoiceAmount() *DetailedAmount4 {
	c.InvoiceAmount = new(DetailedAmount4)
	return c.InvoiceAmount
}

func (c *CardPaymentTransactionDetails28) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails28) SetAccountType(value string) {
	c.AccountType = (*CardAccountType3Code)(&value)
}

func (c *CardPaymentTransactionDetails28) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}

// Details of the transaction in the completion advice.
type CardPaymentTransactionDetails29 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Qualifies the amount associated with the transaction.
	AmountQualifier *TypeOfAmount1Code `xml:"AmtQlfr,omitempty"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount *DetailedAmount15 `xml:"DtldAmt,omitempty"`

	// Amount requested to be authorised.
	RequestedAmount *ImpliedCurrencyAndAmount `xml:"ReqdAmt,omitempty"`

	// Amount authorised for the payment transaction.
	AuthorisedAmount *ImpliedCurrencyAndAmount `xml:"AuthrsdAmt,omitempty"`

	// Amount of the transaction that will be invoiced to the cardholder.
	InvoiceAmount *ImpliedCurrencyAndAmount `xml:"InvcAmt,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Transaction category level on an unattended POI (Point Of Interaction).
	UnattendedLevelCategory *Max35NumericText `xml:"UattnddLvlCtgy,omitempty"`

	// Type of cardholder account used for the transaction.
	AccountType *CardAccountType3Code `xml:"AcctTp,omitempty"`

	// Result of the currency conversion proposed to the cardholder and its result.
	CurrencyConversionResult *CurrencyConversion8 `xml:"CcyConvsRslt,omitempty"`

	// Data related to a financial loan (instalment) or to a recurring transaction.
	Instalment *RecurringTransaction2 `xml:"Instlmt,omitempty"`

	// Payment transaction with an aggregated amount.
	AggregationTransaction *AggregationTransaction2 `xml:"AggtnTx,omitempty"`

	// Codification used to identify the products.
	ProductCodeSetIdentification *Max10Text `xml:"PdctCdSetId,omitempty"`

	// Item purchased with the transaction.
	SaleItem []*Product3 `xml:"SaleItm,omitempty"`

	// Location of the delivery, for instance pump number or parking bay.
	DeliveryLocation *Max35Text `xml:"DlvryLctn,omitempty"`

	// Detailed invoice data.
	CardPaymentInvoice *CardPaymentInvoice2 `xml:"CardPmtInvc,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails29) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails29) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails29) SetAmountQualifier(value string) {
	c.AmountQualifier = (*TypeOfAmount1Code)(&value)
}

func (c *CardPaymentTransactionDetails29) AddDetailedAmount() *DetailedAmount15 {
	c.DetailedAmount = new(DetailedAmount15)
	return c.DetailedAmount
}

func (c *CardPaymentTransactionDetails29) SetRequestedAmount(value, currency string) {
	c.RequestedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails29) SetAuthorisedAmount(value, currency string) {
	c.AuthorisedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails29) SetInvoiceAmount(value, currency string) {
	c.InvoiceAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails29) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails29) SetUnattendedLevelCategory(value string) {
	c.UnattendedLevelCategory = (*Max35NumericText)(&value)
}

func (c *CardPaymentTransactionDetails29) SetAccountType(value string) {
	c.AccountType = (*CardAccountType3Code)(&value)
}

func (c *CardPaymentTransactionDetails29) AddCurrencyConversionResult() *CurrencyConversion8 {
	c.CurrencyConversionResult = new(CurrencyConversion8)
	return c.CurrencyConversionResult
}

func (c *CardPaymentTransactionDetails29) AddInstalment() *RecurringTransaction2 {
	c.Instalment = new(RecurringTransaction2)
	return c.Instalment
}

func (c *CardPaymentTransactionDetails29) AddAggregationTransaction() *AggregationTransaction2 {
	c.AggregationTransaction = new(AggregationTransaction2)
	return c.AggregationTransaction
}

func (c *CardPaymentTransactionDetails29) SetProductCodeSetIdentification(value string) {
	c.ProductCodeSetIdentification = (*Max10Text)(&value)
}

func (c *CardPaymentTransactionDetails29) AddSaleItem() *Product3 {
	newValue := new(Product3)
	c.SaleItem = append(c.SaleItem, newValue)
	return newValue
}

func (c *CardPaymentTransactionDetails29) SetDeliveryLocation(value string) {
	c.DeliveryLocation = (*Max35Text)(&value)
}

func (c *CardPaymentTransactionDetails29) AddCardPaymentInvoice() *CardPaymentInvoice2 {
	c.CardPaymentInvoice = new(CardPaymentInvoice2)
	return c.CardPaymentInvoice
}

func (c *CardPaymentTransactionDetails29) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}

// Details of the transaction in the completion advice.
type CardPaymentTransactionDetails3 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Qualifies the amount associated with the transaction.
	AmountQualifier *TypeOfAmount1Code `xml:"AmtQlfr,omitempty"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount []*DetailedAmount1 `xml:"DtldAmt,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Transaction category level on an unattended POI (Point Of Interaction).
	UnattendedLevelCategory *Max35NumericText `xml:"UattnddLvlCtgy,omitempty"`

	// Type of cardholder account used for the transaction.
	AccountType *CardAccountType1Code `xml:"AcctTp,omitempty"`

	// Data related to a financial loan (instalment) or to a recurring transaction.
	RecurringTransaction *RecurringTransaction1 `xml:"RcrngTx,omitempty"`

	// Product purchased with the transaction.
	Product []*Product1 `xml:"Pdct,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails3) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails3) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails3) SetAmountQualifier(value string) {
	c.AmountQualifier = (*TypeOfAmount1Code)(&value)
}

func (c *CardPaymentTransactionDetails3) AddDetailedAmount() *DetailedAmount1 {
	newValue := new(DetailedAmount1)
	c.DetailedAmount = append(c.DetailedAmount, newValue)
	return newValue
}

func (c *CardPaymentTransactionDetails3) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails3) SetUnattendedLevelCategory(value string) {
	c.UnattendedLevelCategory = (*Max35NumericText)(&value)
}

func (c *CardPaymentTransactionDetails3) SetAccountType(value string) {
	c.AccountType = (*CardAccountType1Code)(&value)
}

func (c *CardPaymentTransactionDetails3) AddRecurringTransaction() *RecurringTransaction1 {
	c.RecurringTransaction = new(RecurringTransaction1)
	return c.RecurringTransaction
}

func (c *CardPaymentTransactionDetails3) AddProduct() *Product1 {
	newValue := new(Product1)
	c.Product = append(c.Product, newValue)
	return newValue
}

func (c *CardPaymentTransactionDetails3) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}

// Details of the transaction to capture.
type CardPaymentTransactionDetails30 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Qualifies the amount associated with the transaction.
	AmountQualifier *TypeOfAmount1Code `xml:"AmtQlfr,omitempty"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount *DetailedAmount15 `xml:"DtldAmt,omitempty"`

	// Amount requested to be authorised.
	RequestedAmount *ImpliedCurrencyAndAmount `xml:"ReqdAmt,omitempty"`

	// Amount authorised for the transaction.
	AuthorisedAmount *ImpliedCurrencyAndAmount `xml:"AuthrsdAmt,omitempty"`

	// Amount of the transaction that will be invoiced to the cardholder.
	InvoiceAmount *ImpliedCurrencyAndAmount `xml:"InvcAmt,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Transaction category level on an unattended POI (Point Of Interaction).
	UnattendedLevelCategory *Max35NumericText `xml:"UattnddLvlCtgy,omitempty"`

	// Type of cardholder account used for the transaction.
	AccountType *CardAccountType3Code `xml:"AcctTp,omitempty"`

	// Result of the currency conversion proposed to the cardholder.
	CurrencyConversionResult *CurrencyConversion8 `xml:"CcyConvsRslt,omitempty"`

	// Data related to a financial loan (instalment) or to a recurring transaction.
	Instalment *RecurringTransaction2 `xml:"Instlmt,omitempty"`

	// Payment transaction with an aggregated amount.
	AggregationTransaction *AggregationTransaction2 `xml:"AggtnTx,omitempty"`

	// Identification of the product codes that are purchased.
	ProductCodeSetIdentification *Max10Text `xml:"PdctCdSetId,omitempty"`

	// Item purchased with the transaction.
	SaleItem []*Product3 `xml:"SaleItm,omitempty"`

	// Location of the delivery, for instance pump number or parking bay.
	DeliveryLocation *Max10Text `xml:"DlvryLctn,omitempty"`

	// Detailed invoice data.
	CardPaymentInvoice *CardPaymentInvoice2 `xml:"CardPmtInvc,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails30) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails30) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails30) SetAmountQualifier(value string) {
	c.AmountQualifier = (*TypeOfAmount1Code)(&value)
}

func (c *CardPaymentTransactionDetails30) AddDetailedAmount() *DetailedAmount15 {
	c.DetailedAmount = new(DetailedAmount15)
	return c.DetailedAmount
}

func (c *CardPaymentTransactionDetails30) SetRequestedAmount(value, currency string) {
	c.RequestedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails30) SetAuthorisedAmount(value, currency string) {
	c.AuthorisedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails30) SetInvoiceAmount(value, currency string) {
	c.InvoiceAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails30) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails30) SetUnattendedLevelCategory(value string) {
	c.UnattendedLevelCategory = (*Max35NumericText)(&value)
}

func (c *CardPaymentTransactionDetails30) SetAccountType(value string) {
	c.AccountType = (*CardAccountType3Code)(&value)
}

func (c *CardPaymentTransactionDetails30) AddCurrencyConversionResult() *CurrencyConversion8 {
	c.CurrencyConversionResult = new(CurrencyConversion8)
	return c.CurrencyConversionResult
}

func (c *CardPaymentTransactionDetails30) AddInstalment() *RecurringTransaction2 {
	c.Instalment = new(RecurringTransaction2)
	return c.Instalment
}

func (c *CardPaymentTransactionDetails30) AddAggregationTransaction() *AggregationTransaction2 {
	c.AggregationTransaction = new(AggregationTransaction2)
	return c.AggregationTransaction
}

func (c *CardPaymentTransactionDetails30) SetProductCodeSetIdentification(value string) {
	c.ProductCodeSetIdentification = (*Max10Text)(&value)
}

func (c *CardPaymentTransactionDetails30) AddSaleItem() *Product3 {
	newValue := new(Product3)
	c.SaleItem = append(c.SaleItem, newValue)
	return newValue
}

func (c *CardPaymentTransactionDetails30) SetDeliveryLocation(value string) {
	c.DeliveryLocation = (*Max10Text)(&value)
}

func (c *CardPaymentTransactionDetails30) AddCardPaymentInvoice() *CardPaymentInvoice2 {
	c.CardPaymentInvoice = new(CardPaymentInvoice2)
	return c.CardPaymentInvoice
}

func (c *CardPaymentTransactionDetails30) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}

// Details of the transaction in the authorisation request in a batch.
type CardPaymentTransactionDetails31 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Qualifies the amount associated with the transaction.
	AmountQualifier *TypeOfAmount1Code `xml:"AmtQlfr,omitempty"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount *DetailedAmount15 `xml:"DtldAmt,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Reason to process an online authorisation.
	OnLineReason *OnLineReason1Code `xml:"OnLineRsn,omitempty"`

	// Transaction category level on an unattended POI (Point Of Interaction).
	UnattendedLevelCategory *Max35NumericText `xml:"UattnddLvlCtgy,omitempty"`

	// Type of cardholder account used for the transaction.
	AccountType *CardAccountType3Code `xml:"AcctTp,omitempty"`

	// Result of the currency conversion proposed to the cardholder.
	CurrencyConversionResult *CurrencyConversion8 `xml:"CcyConvsRslt,omitempty"`

	// Data related to a financial loan (instalment) or to a recurring transaction.
	Instalment *RecurringTransaction2 `xml:"Instlmt,omitempty"`

	// Payment transaction with an aggregated amount.
	AggregationTransaction *AggregationTransaction2 `xml:"AggtnTx,omitempty"`

	// Codification used to identify the products.
	ProductCodeSetIdentification *Max10Text `xml:"PdctCdSetId,omitempty"`

	// Item purchased with the transaction.
	SaleItem []*Product3 `xml:"SaleItm,omitempty"`

	// Location of the delivery, for instance pump number or parking bay.
	DeliveryLocation *Max10Text `xml:"DlvryLctn,omitempty"`

	// Detailed invoice data.
	CardPaymentInvoice *CardPaymentInvoice2 `xml:"CardPmtInvc,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails31) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails31) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails31) SetAmountQualifier(value string) {
	c.AmountQualifier = (*TypeOfAmount1Code)(&value)
}

func (c *CardPaymentTransactionDetails31) AddDetailedAmount() *DetailedAmount15 {
	c.DetailedAmount = new(DetailedAmount15)
	return c.DetailedAmount
}

func (c *CardPaymentTransactionDetails31) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails31) SetOnLineReason(value string) {
	c.OnLineReason = (*OnLineReason1Code)(&value)
}

func (c *CardPaymentTransactionDetails31) SetUnattendedLevelCategory(value string) {
	c.UnattendedLevelCategory = (*Max35NumericText)(&value)
}

func (c *CardPaymentTransactionDetails31) SetAccountType(value string) {
	c.AccountType = (*CardAccountType3Code)(&value)
}

func (c *CardPaymentTransactionDetails31) AddCurrencyConversionResult() *CurrencyConversion8 {
	c.CurrencyConversionResult = new(CurrencyConversion8)
	return c.CurrencyConversionResult
}

func (c *CardPaymentTransactionDetails31) AddInstalment() *RecurringTransaction2 {
	c.Instalment = new(RecurringTransaction2)
	return c.Instalment
}

func (c *CardPaymentTransactionDetails31) AddAggregationTransaction() *AggregationTransaction2 {
	c.AggregationTransaction = new(AggregationTransaction2)
	return c.AggregationTransaction
}

func (c *CardPaymentTransactionDetails31) SetProductCodeSetIdentification(value string) {
	c.ProductCodeSetIdentification = (*Max10Text)(&value)
}

func (c *CardPaymentTransactionDetails31) AddSaleItem() *Product3 {
	newValue := new(Product3)
	c.SaleItem = append(c.SaleItem, newValue)
	return newValue
}

func (c *CardPaymentTransactionDetails31) SetDeliveryLocation(value string) {
	c.DeliveryLocation = (*Max10Text)(&value)
}

func (c *CardPaymentTransactionDetails31) AddCardPaymentInvoice() *CardPaymentInvoice2 {
	c.CardPaymentInvoice = new(CardPaymentInvoice2)
	return c.CardPaymentInvoice
}

func (c *CardPaymentTransactionDetails31) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}

// Details of the transaction in the authorisation response contained in a batch.
type CardPaymentTransactionDetails32 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount *DetailedAmount15 `xml:"DtldAmt,omitempty"`

	// Amount of the transaction that will be invoiced to the cardholder.
	InvoiceAmount *ImpliedCurrencyAndAmount `xml:"InvcAmt,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Type of cardholder account used for the transaction.
	AccountType *CardAccountType3Code `xml:"AcctTp,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails32) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails32) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails32) AddDetailedAmount() *DetailedAmount15 {
	c.DetailedAmount = new(DetailedAmount15)
	return c.DetailedAmount
}

func (c *CardPaymentTransactionDetails32) SetInvoiceAmount(value, currency string) {
	c.InvoiceAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails32) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails32) SetAccountType(value string) {
	c.AccountType = (*CardAccountType3Code)(&value)
}

func (c *CardPaymentTransactionDetails32) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}

// Details of the payment transaction.
type CardPaymentTransactionDetails33 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Qualifies the amount associated with the transaction.
	AmountQualifier *TypeOfAmount1Code `xml:"AmtQlfr,omitempty"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount *DetailedAmount15 `xml:"DtldAmt,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails33) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails33) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails33) SetAmountQualifier(value string) {
	c.AmountQualifier = (*TypeOfAmount1Code)(&value)
}

func (c *CardPaymentTransactionDetails33) AddDetailedAmount() *DetailedAmount15 {
	c.DetailedAmount = new(DetailedAmount15)
	return c.DetailedAmount
}

func (c *CardPaymentTransactionDetails33) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}

// Details of the transaction in the cancellation inside a batch transfer.
type CardPaymentTransactionDetails34 struct {

	// Currency associated with the transaction.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails34) SetCurrency(value string) {
	c.Currency = (*ActiveCurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails34) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails34) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails34) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}

// Details of the transaction in the cancellation response.
type CardPaymentTransactionDetails35 struct {

	// Currency associated with the transaction.
	Currency *ActiveCurrencyCode `xml:"Ccy"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails35) SetCurrency(value string) {
	c.Currency = (*ActiveCurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails35) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails35) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}

// Details of the transaction to capture.
type CardPaymentTransactionDetails4 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Qualifies the amount associated with the transaction.
	AmountQualifier *TypeOfAmount1Code `xml:"AmtQlfr,omitempty"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount []*DetailedAmount1 `xml:"DtldAmt,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Transaction category level on an unattended POI (Point Of Interaction).
	UnattendedLevelCategory *Max35NumericText `xml:"UattnddLvlCtgy,omitempty"`

	// Type of cardholder account used for the transaction.
	AccountType *CardAccountType1Code `xml:"AcctTp,omitempty"`

	// Data related to a financial loan (instalment) or to a recurring transaction.
	RecurringTransaction *RecurringTransaction1 `xml:"RcrngTx,omitempty"`

	// Product purchased with the transaction.
	Product []*Product1 `xml:"Pdct,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails4) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails4) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails4) SetAmountQualifier(value string) {
	c.AmountQualifier = (*TypeOfAmount1Code)(&value)
}

func (c *CardPaymentTransactionDetails4) AddDetailedAmount() *DetailedAmount1 {
	newValue := new(DetailedAmount1)
	c.DetailedAmount = append(c.DetailedAmount, newValue)
	return newValue
}

func (c *CardPaymentTransactionDetails4) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails4) SetUnattendedLevelCategory(value string) {
	c.UnattendedLevelCategory = (*Max35NumericText)(&value)
}

func (c *CardPaymentTransactionDetails4) SetAccountType(value string) {
	c.AccountType = (*CardAccountType1Code)(&value)
}

func (c *CardPaymentTransactionDetails4) AddRecurringTransaction() *RecurringTransaction1 {
	c.RecurringTransaction = new(RecurringTransaction1)
	return c.RecurringTransaction
}

func (c *CardPaymentTransactionDetails4) AddProduct() *Product1 {
	newValue := new(Product1)
	c.Product = append(c.Product, newValue)
	return newValue
}

func (c *CardPaymentTransactionDetails4) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}

// Details of the transaction in the cancellation request.
type CardPaymentTransactionDetails5 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails5) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails5) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails5) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails5) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}

// Details of the transaction in the cancellation response.
type CardPaymentTransactionDetails6 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`
}

func (c *CardPaymentTransactionDetails6) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails6) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

// Details of the transaction in the cancellation advice.
type CardPaymentTransactionDetails7 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails7) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails7) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails7) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails7) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}

// Details of the transaction in the authorisation request.
type CardPaymentTransactionDetails8 struct {

	// Amounts associated with the total amount of transaction.
	Amount []*CardAmountAndCurrencyExchange1 `xml:"Amt,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max1025Text `xml:"ICCRltdData,omitempty"`

	// Context of the card payment transaction.
	PaymentContext *PaymentContext3 `xml:"PmtCntxt,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd,omitempty"`

	// Local date and time of the transaction assigned by the POI (Point Of Interaction).
	TransactionDateTime *ISODateTime `xml:"TxDtTm,omitempty"`

	// Identification of a sale transaction assigned by the sale system.
	SaleReferenceNumber *Max35Text `xml:"SaleRefNb,omitempty"`

	// Reason for representment of a card transaction.
	RePresentmentReason *ExternalRePresentmentReason1Code `xml:"RePresntmntRsn,omitempty"`

	// Service in addition to the main service.
	AdditionalService *CardPaymentServiceType2Code `xml:"AddtlSvc,omitempty"`

	// Identification of the transaction that has to be unique for a time period.
	TransactionReference *Max35Text `xml:"TxRef,omitempty"`
}

func (c *CardPaymentTransactionDetails8) AddAmount() *CardAmountAndCurrencyExchange1 {
	newValue := new(CardAmountAndCurrencyExchange1)
	c.Amount = append(c.Amount, newValue)
	return newValue
}

func (c *CardPaymentTransactionDetails8) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max1025Text)(&value)
}

func (c *CardPaymentTransactionDetails8) AddPaymentContext() *PaymentContext3 {
	c.PaymentContext = new(PaymentContext3)
	return c.PaymentContext
}

func (c *CardPaymentTransactionDetails8) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransactionDetails8) SetTransactionDateTime(value string) {
	c.TransactionDateTime = (*ISODateTime)(&value)
}

func (c *CardPaymentTransactionDetails8) SetSaleReferenceNumber(value string) {
	c.SaleReferenceNumber = (*Max35Text)(&value)
}

func (c *CardPaymentTransactionDetails8) SetRePresentmentReason(value string) {
	c.RePresentmentReason = (*ExternalRePresentmentReason1Code)(&value)
}

func (c *CardPaymentTransactionDetails8) SetAdditionalService(value string) {
	c.AdditionalService = (*CardPaymentServiceType2Code)(&value)
}

func (c *CardPaymentTransactionDetails8) SetTransactionReference(value string) {
	c.TransactionReference = (*Max35Text)(&value)
}

// Details of the transaction in the cancellation inside a batch transfer.
type CardPaymentTransactionDetails9 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails9) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails9) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails9) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails9) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}
