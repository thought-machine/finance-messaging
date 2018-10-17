package iso20022

// Information about a payment against a purchase order.
type ReportLine1 struct {

	// Unique identification assigned by the matching application to the transaction.
	TransactionIdentification *Max35Text `xml:"TxId"`

	// Identifies the status of the transaction.
	TransactionStatus *TransactionStatus4 `xml:"TxSts"`

	// Unique identification of the purchase order, assigned by the buyer.
	PurchaseOrderReference *DocumentIdentification7 `xml:"PurchsOrdrRef"`

	// Total amount of the purchase order after taxes, adjustments and charges.
	PurchaseOrderTotalNetAmount *CurrencyAndAmount `xml:"PurchsOrdrTtlNetAmt"`

	// Accumulated net amount, after adjustments, intended to be paid.
	AccumulatedNetAmount *CurrencyAndAmount `xml:"AcmltdNetAmt"`
}

func (r *ReportLine1) SetTransactionIdentification(value string) {
	r.TransactionIdentification = (*Max35Text)(&value)
}

func (r *ReportLine1) AddTransactionStatus() *TransactionStatus4 {
	r.TransactionStatus = new(TransactionStatus4)
	return r.TransactionStatus
}

func (r *ReportLine1) AddPurchaseOrderReference() *DocumentIdentification7 {
	r.PurchaseOrderReference = new(DocumentIdentification7)
	return r.PurchaseOrderReference
}

func (r *ReportLine1) SetPurchaseOrderTotalNetAmount(value, currency string) {
	r.PurchaseOrderTotalNetAmount = NewCurrencyAndAmount(value, currency)
}

func (r *ReportLine1) SetAccumulatedNetAmount(value, currency string) {
	r.AccumulatedNetAmount = NewCurrencyAndAmount(value, currency)
}

// Information about a payment against a purchase order.
type ReportLine2 struct {

	// Unique identification assigned by the matching application to the transaction.
	TransactionIdentification *Max35Text `xml:"TxId"`

	// Unique identification of the purchase order, assigned by the buyer.
	PurchaseOrderReference *DocumentIdentification7 `xml:"PurchsOrdrRef"`

	// Specifies the adjustments applied to obtain the net amount.
	Adjustment []*Adjustment4 `xml:"Adjstmnt,omitempty"`

	// Net amount, after adjustments, intended to be paid.
	NetAmount *CurrencyAndAmount `xml:"NetAmt"`
}

func (r *ReportLine2) SetTransactionIdentification(value string) {
	r.TransactionIdentification = (*Max35Text)(&value)
}

func (r *ReportLine2) AddPurchaseOrderReference() *DocumentIdentification7 {
	r.PurchaseOrderReference = new(DocumentIdentification7)
	return r.PurchaseOrderReference
}

func (r *ReportLine2) AddAdjustment() *Adjustment4 {
	newValue := new(Adjustment4)
	r.Adjustment = append(r.Adjustment, newValue)
	return newValue
}

func (r *ReportLine2) SetNetAmount(value, currency string) {
	r.NetAmount = NewCurrencyAndAmount(value, currency)
}

// Information about a payment against a purchase order.
type ReportLine3 struct {

	// Unique identification of the purchase order, assigned by the buyer.
	PurchaseOrderReference *DocumentIdentification7 `xml:"PurchsOrdrRef"`

	// Specifies the adjustments applied to obtain the net amount.
	Adjustment []*Adjustment4 `xml:"Adjstmnt,omitempty"`

	// Net amount, after adjustments, intended to be paid.
	NetAmount *CurrencyAndAmount `xml:"NetAmt"`
}

func (r *ReportLine3) AddPurchaseOrderReference() *DocumentIdentification7 {
	r.PurchaseOrderReference = new(DocumentIdentification7)
	return r.PurchaseOrderReference
}

func (r *ReportLine3) AddAdjustment() *Adjustment4 {
	newValue := new(Adjustment4)
	r.Adjustment = append(r.Adjustment, newValue)
	return newValue
}

func (r *ReportLine3) SetNetAmount(value, currency string) {
	r.NetAmount = NewCurrencyAndAmount(value, currency)
}

// Information about a payment against a commercial invoice.
type ReportLine4 struct {

	// Reference to the identification of the underlying commercial document.
	CommercialDocumentReference *InvoiceIdentification1 `xml:"ComrclDocRef"`

	// Specifies the adjustments applied to obtain the net amount.
	Adjustment []*Adjustment4 `xml:"Adjstmnt,omitempty"`

	// Net amount, after adjustments, intended to be paid.
	NetAmount *CurrencyAndAmount `xml:"NetAmt"`

	// Specifies how the net amount to be paid is related to different purchase orders.
	BreakdownByPurchaseOrder []*ReportLine2 `xml:"BrkdwnByPurchsOrdr"`
}

func (r *ReportLine4) AddCommercialDocumentReference() *InvoiceIdentification1 {
	r.CommercialDocumentReference = new(InvoiceIdentification1)
	return r.CommercialDocumentReference
}

func (r *ReportLine4) AddAdjustment() *Adjustment4 {
	newValue := new(Adjustment4)
	r.Adjustment = append(r.Adjustment, newValue)
	return newValue
}

func (r *ReportLine4) SetNetAmount(value, currency string) {
	r.NetAmount = NewCurrencyAndAmount(value, currency)
}

func (r *ReportLine4) AddBreakdownByPurchaseOrder() *ReportLine2 {
	newValue := new(ReportLine2)
	r.BreakdownByPurchaseOrder = append(r.BreakdownByPurchaseOrder, newValue)
	return newValue
}

// Information about a payment against a purchase order.
type ReportLine5 struct {

	// Unique identification of the purchase order, assigned by the buyer.
	PurchaseOrderReference *DocumentIdentification7 `xml:"PurchsOrdrRef"`

	// Specifies the adjustments applied to obtain the net amount.
	Adjustment []*Adjustment6 `xml:"Adjstmnt,omitempty"`

	// Net amount, after adjustments, intended to be paid.
	NetAmount *CurrencyAndAmount `xml:"NetAmt"`
}

func (r *ReportLine5) AddPurchaseOrderReference() *DocumentIdentification7 {
	r.PurchaseOrderReference = new(DocumentIdentification7)
	return r.PurchaseOrderReference
}

func (r *ReportLine5) AddAdjustment() *Adjustment6 {
	newValue := new(Adjustment6)
	r.Adjustment = append(r.Adjustment, newValue)
	return newValue
}

func (r *ReportLine5) SetNetAmount(value, currency string) {
	r.NetAmount = NewCurrencyAndAmount(value, currency)
}

// Information about a payment against a commercial invoice.
type ReportLine6 struct {

	// Reference to the identification of the underlying commercial document.
	CommercialDocumentReference *InvoiceIdentification1 `xml:"ComrclDocRef"`

	// Specifies the adjustments applied to obtain the net amount.
	Adjustment []*Adjustment6 `xml:"Adjstmnt,omitempty"`

	// Net amount, after adjustments, intended to be paid.
	NetAmount *CurrencyAndAmount `xml:"NetAmt"`

	// Specifies how the net amount to be paid is related to different purchase orders.
	BreakdownByPurchaseOrder []*ReportLine7 `xml:"BrkdwnByPurchsOrdr"`
}

func (r *ReportLine6) AddCommercialDocumentReference() *InvoiceIdentification1 {
	r.CommercialDocumentReference = new(InvoiceIdentification1)
	return r.CommercialDocumentReference
}

func (r *ReportLine6) AddAdjustment() *Adjustment6 {
	newValue := new(Adjustment6)
	r.Adjustment = append(r.Adjustment, newValue)
	return newValue
}

func (r *ReportLine6) SetNetAmount(value, currency string) {
	r.NetAmount = NewCurrencyAndAmount(value, currency)
}

func (r *ReportLine6) AddBreakdownByPurchaseOrder() *ReportLine7 {
	newValue := new(ReportLine7)
	r.BreakdownByPurchaseOrder = append(r.BreakdownByPurchaseOrder, newValue)
	return newValue
}

// Information about a payment against a purchase order.
type ReportLine7 struct {

	// Unique identification assigned by the matching application to the transaction.
	TransactionIdentification *Max35Text `xml:"TxId"`

	// Unique identification of the purchase order, assigned by the buyer.
	PurchaseOrderReference *DocumentIdentification7 `xml:"PurchsOrdrRef"`

	// Specifies the adjustments applied to obtain the net amount.
	Adjustment []*Adjustment6 `xml:"Adjstmnt,omitempty"`

	// Net amount, after adjustments, intended to be paid.
	NetAmount *CurrencyAndAmount `xml:"NetAmt"`
}

func (r *ReportLine7) SetTransactionIdentification(value string) {
	r.TransactionIdentification = (*Max35Text)(&value)
}

func (r *ReportLine7) AddPurchaseOrderReference() *DocumentIdentification7 {
	r.PurchaseOrderReference = new(DocumentIdentification7)
	return r.PurchaseOrderReference
}

func (r *ReportLine7) AddAdjustment() *Adjustment6 {
	newValue := new(Adjustment6)
	r.Adjustment = append(r.Adjustment, newValue)
	return newValue
}

func (r *ReportLine7) SetNetAmount(value, currency string) {
	r.NetAmount = NewCurrencyAndAmount(value, currency)
}
