package iso20022

// Data common to all transactions of a data set.
type CommonData1 struct {

	// Data related to the environment of the transaction, common to a set of transaction.
	Environment *CardPaymentEnvironment5 `xml:"Envt,omitempty"`

	// Data related to the context of the transaction, common to a set of transaction.
	Context *CardPaymentContext3 `xml:"Cntxt,omitempty"`

	// Type of transaction being undertaken for the main service, common to a set of transaction..
	TransactionType *CardPaymentServiceType1Code `xml:"TxTp,omitempty"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType2Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Currency associated with the set of transaction.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`
}

func (c *CommonData1) AddEnvironment() *CardPaymentEnvironment5 {
	c.Environment = new(CardPaymentEnvironment5)
	return c.Environment
}

func (c *CommonData1) AddContext() *CardPaymentContext3 {
	c.Context = new(CardPaymentContext3)
	return c.Context
}

func (c *CommonData1) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType1Code)(&value)
}

func (c *CommonData1) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType2Code)(&value))
}

func (c *CommonData1) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CommonData1) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CommonData1) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CommonData1) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

// Data common to all transactions of a data set.
type CommonData2 struct {

	// Data related to the environment of the transaction, common to a set of transaction.
	Environment *CardPaymentEnvironment13 `xml:"Envt,omitempty"`

	// Data related to the context of the transaction, common to a set of transaction.
	Context *CardPaymentContext4 `xml:"Cntxt,omitempty"`

	// Type of transaction being undertaken for the main service, common to a set of transaction..
	TransactionType *CardPaymentServiceType4Code `xml:"TxTp,omitempty"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType2Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Currency associated with the set of transaction.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`
}

func (c *CommonData2) AddEnvironment() *CardPaymentEnvironment13 {
	c.Environment = new(CardPaymentEnvironment13)
	return c.Environment
}

func (c *CommonData2) AddContext() *CardPaymentContext4 {
	c.Context = new(CardPaymentContext4)
	return c.Context
}

func (c *CommonData2) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType4Code)(&value)
}

func (c *CommonData2) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType2Code)(&value))
}

func (c *CommonData2) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CommonData2) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CommonData2) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CommonData2) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

// Data common to all transactions of a data set.
type CommonData3 struct {

	// Data related to the environment of the transaction, common to a set of transaction.
	Environment *CardPaymentEnvironment26 `xml:"Envt,omitempty"`

	// Data related to the context of the transaction, common to a set of transaction.
	Context *CardPaymentContext7 `xml:"Cntxt,omitempty"`

	// Type of transaction being undertaken for the main service, common to a set of transaction..
	TransactionType *CardPaymentServiceType5Code `xml:"TxTp,omitempty"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType6Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Currency associated with the set of transaction.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`
}

func (c *CommonData3) AddEnvironment() *CardPaymentEnvironment26 {
	c.Environment = new(CardPaymentEnvironment26)
	return c.Environment
}

func (c *CommonData3) AddContext() *CardPaymentContext7 {
	c.Context = new(CardPaymentContext7)
	return c.Context
}

func (c *CommonData3) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType5Code)(&value)
}

func (c *CommonData3) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType6Code)(&value))
}

func (c *CommonData3) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CommonData3) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CommonData3) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CommonData3) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

// Data common to all transactions of a data set.
type CommonData4 struct {

	// Data related to the environment of the transaction, common to a set of transaction.
	Environment *CardPaymentEnvironment39 `xml:"Envt,omitempty"`

	// Data related to the context of the transaction, common to a set of transaction.
	Context *CardPaymentContext12 `xml:"Cntxt,omitempty"`

	// Type of transaction being undertaken for the main service, common to a set of transaction..
	TransactionType *CardPaymentServiceType5Code `xml:"TxTp,omitempty"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType6Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Currency associated with the set of transaction.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`
}

func (c *CommonData4) AddEnvironment() *CardPaymentEnvironment39 {
	c.Environment = new(CardPaymentEnvironment39)
	return c.Environment
}

func (c *CommonData4) AddContext() *CardPaymentContext12 {
	c.Context = new(CardPaymentContext12)
	return c.Context
}

func (c *CommonData4) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType5Code)(&value)
}

func (c *CommonData4) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType6Code)(&value))
}

func (c *CommonData4) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CommonData4) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CommonData4) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CommonData4) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

// Data common to all transactions of a data set.
type CommonData5 struct {

	// Data related to the environment of the transaction, common to a set of transaction.
	Environment *CardPaymentEnvironment51 `xml:"Envt,omitempty"`

	// Data related to the context of the transaction, common to a set of transaction.
	Context *CardPaymentContext18 `xml:"Cntxt,omitempty"`

	// Type of transaction being undertaken for the main service, common to a set of transaction..
	TransactionType *CardPaymentServiceType5Code `xml:"TxTp,omitempty"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType9Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Currency associated with the set of transaction.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`
}

func (c *CommonData5) AddEnvironment() *CardPaymentEnvironment51 {
	c.Environment = new(CardPaymentEnvironment51)
	return c.Environment
}

func (c *CommonData5) AddContext() *CardPaymentContext18 {
	c.Context = new(CardPaymentContext18)
	return c.Context
}

func (c *CommonData5) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType5Code)(&value)
}

func (c *CommonData5) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType9Code)(&value))
}

func (c *CommonData5) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CommonData5) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CommonData5) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CommonData5) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}
