package iso20022

// Set of elements used to identify the underlying (group of) transaction(s) to which the investigation applies.
type UnderlyingTransaction1 struct {

	// Set of elements used to provide information on the original messsage, to which the cancellation refers.
	OriginalGroupInformationAndCancellation *OriginalGroupInformation23 `xml:"OrgnlGrpInfAndCxl,omitempty"`

	// Set of elements used to provide information on the original (group of) transactions, to which the cancellation request refers.
	OriginalPaymentInformationAndCancellation []*OriginalPaymentInformation4 `xml:"OrgnlPmtInfAndCxl,omitempty"`
}

func (u *UnderlyingTransaction1) AddOriginalGroupInformationAndCancellation() *OriginalGroupInformation23 {
	u.OriginalGroupInformationAndCancellation = new(OriginalGroupInformation23)
	return u.OriginalGroupInformationAndCancellation
}

func (u *UnderlyingTransaction1) AddOriginalPaymentInformationAndCancellation() *OriginalPaymentInformation4 {
	newValue := new(OriginalPaymentInformation4)
	u.OriginalPaymentInformationAndCancellation = append(u.OriginalPaymentInformationAndCancellation, newValue)
	return newValue
}

// Identifies the underlying (group of) transaction(s) to which the investigation applies.
type UnderlyingTransaction10 struct {

	// Provides information on the original message, to which the cancellation refers.
	OriginalGroupInformationAndCancellation *OriginalGroupHeader4 `xml:"OrgnlGrpInfAndCxl,omitempty"`

	// Provides information on the original transactions to which the cancellation request message refers.
	TransactionInformation []*PaymentTransaction55 `xml:"TxInf,omitempty"`
}

func (u *UnderlyingTransaction10) AddOriginalGroupInformationAndCancellation() *OriginalGroupHeader4 {
	u.OriginalGroupInformationAndCancellation = new(OriginalGroupHeader4)
	return u.OriginalGroupInformationAndCancellation
}

func (u *UnderlyingTransaction10) AddTransactionInformation() *PaymentTransaction55 {
	newValue := new(PaymentTransaction55)
	u.TransactionInformation = append(u.TransactionInformation, newValue)
	return newValue
}

// Identifies the underlying (group of) transaction(s) to which the investigation applies.
type UnderlyingTransaction11 struct {

	// Provides information on the original message, to which the cancellation refers.
	OriginalGroupInformationAndCancellation *OriginalGroupHeader4 `xml:"OrgnlGrpInfAndCxl,omitempty"`

	// Provides information on the original (group of) transactions, to which the cancellation request refers.
	OriginalPaymentInformationAndCancellation []*OriginalPaymentInstruction13 `xml:"OrgnlPmtInfAndCxl,omitempty"`
}

func (u *UnderlyingTransaction11) AddOriginalGroupInformationAndCancellation() *OriginalGroupHeader4 {
	u.OriginalGroupInformationAndCancellation = new(OriginalGroupHeader4)
	return u.OriginalGroupInformationAndCancellation
}

func (u *UnderlyingTransaction11) AddOriginalPaymentInformationAndCancellation() *OriginalPaymentInstruction13 {
	newValue := new(OriginalPaymentInstruction13)
	u.OriginalPaymentInformationAndCancellation = append(u.OriginalPaymentInformationAndCancellation, newValue)
	return newValue
}

// Identifies the underlying (group of) transaction(s) to which the investigation applies.
type UnderlyingTransaction12 struct {

	// Provides information on the original message, to which the cancellation refers.
	OriginalGroupInformationAndCancellation *OriginalGroupHeader4 `xml:"OrgnlGrpInfAndCxl,omitempty"`

	// Provides information on the original (group of) transactions, to which the cancellation request refers.
	OriginalPaymentInformationAndCancellation []*OriginalPaymentInstruction15 `xml:"OrgnlPmtInfAndCxl,omitempty"`
}

func (u *UnderlyingTransaction12) AddOriginalGroupInformationAndCancellation() *OriginalGroupHeader4 {
	u.OriginalGroupInformationAndCancellation = new(OriginalGroupHeader4)
	return u.OriginalGroupInformationAndCancellation
}

func (u *UnderlyingTransaction12) AddOriginalPaymentInformationAndCancellation() *OriginalPaymentInstruction15 {
	newValue := new(OriginalPaymentInstruction15)
	u.OriginalPaymentInformationAndCancellation = append(u.OriginalPaymentInformationAndCancellation, newValue)
	return newValue
}

// Identifies the underlying (group of) transaction(s) to which the investigation applies.
type UnderlyingTransaction13 struct {

	// Provides information on the original message, to which the cancellation refers.
	OriginalGroupInformationAndCancellation *OriginalGroupHeader4 `xml:"OrgnlGrpInfAndCxl,omitempty"`

	// Provides information on the original transactions to which the cancellation request message refers.
	TransactionInformation []*PaymentTransaction62 `xml:"TxInf,omitempty"`
}

func (u *UnderlyingTransaction13) AddOriginalGroupInformationAndCancellation() *OriginalGroupHeader4 {
	u.OriginalGroupInformationAndCancellation = new(OriginalGroupHeader4)
	return u.OriginalGroupInformationAndCancellation
}

func (u *UnderlyingTransaction13) AddTransactionInformation() *PaymentTransaction62 {
	newValue := new(PaymentTransaction62)
	u.TransactionInformation = append(u.TransactionInformation, newValue)
	return newValue
}

// Identifies the underlying (group of) transaction(s) to which the resolution of investigation applies.
type UnderlyingTransaction14 struct {

	// Provides information on the original cancellation message, to which the resolution refers.
	OriginalGroupInformationAndStatus *OriginalGroupHeader5 `xml:"OrgnlGrpInfAndSts,omitempty"`

	// Provides information on the original (group of) transactions, to which the cancellation status refers.
	OriginalPaymentInformationAndStatus []*OriginalPaymentInstruction17 `xml:"OrgnlPmtInfAndSts,omitempty"`

	// Provides details on the original transactions to which the cancellation request message refers.
	TransactionInformationAndStatus []*PaymentTransaction67 `xml:"TxInfAndSts,omitempty"`
}

func (u *UnderlyingTransaction14) AddOriginalGroupInformationAndStatus() *OriginalGroupHeader5 {
	u.OriginalGroupInformationAndStatus = new(OriginalGroupHeader5)
	return u.OriginalGroupInformationAndStatus
}

func (u *UnderlyingTransaction14) AddOriginalPaymentInformationAndStatus() *OriginalPaymentInstruction17 {
	newValue := new(OriginalPaymentInstruction17)
	u.OriginalPaymentInformationAndStatus = append(u.OriginalPaymentInformationAndStatus, newValue)
	return newValue
}

func (u *UnderlyingTransaction14) AddTransactionInformationAndStatus() *PaymentTransaction67 {
	newValue := new(PaymentTransaction67)
	u.TransactionInformationAndStatus = append(u.TransactionInformationAndStatus, newValue)
	return newValue
}

// Identifies the underlying (group of) transaction(s) to which the investigation applies.
type UnderlyingTransaction15 struct {

	// Provides information on the original message, to which the cancellation refers.
	OriginalGroupInformationAndCancellation *OriginalGroupHeader6 `xml:"OrgnlGrpInfAndCxl,omitempty"`

	// Provides information on the original (group of) transactions, to which the cancellation request refers.
	OriginalPaymentInformationAndCancellation []*OriginalPaymentInstruction20 `xml:"OrgnlPmtInfAndCxl,omitempty"`
}

func (u *UnderlyingTransaction15) AddOriginalGroupInformationAndCancellation() *OriginalGroupHeader6 {
	u.OriginalGroupInformationAndCancellation = new(OriginalGroupHeader6)
	return u.OriginalGroupInformationAndCancellation
}

func (u *UnderlyingTransaction15) AddOriginalPaymentInformationAndCancellation() *OriginalPaymentInstruction20 {
	newValue := new(OriginalPaymentInstruction20)
	u.OriginalPaymentInformationAndCancellation = append(u.OriginalPaymentInformationAndCancellation, newValue)
	return newValue
}

// Identifies the underlying (group of) transaction(s) to which the investigation applies.
type UnderlyingTransaction16 struct {

	// Provides information on the original message, to which the cancellation refers.
	OriginalGroupInformationAndCancellation *OriginalGroupHeader6 `xml:"OrgnlGrpInfAndCxl,omitempty"`

	// Provides information on the original transactions to which the cancellation request message refers.
	TransactionInformation []*PaymentTransaction75 `xml:"TxInf,omitempty"`
}

func (u *UnderlyingTransaction16) AddOriginalGroupInformationAndCancellation() *OriginalGroupHeader6 {
	u.OriginalGroupInformationAndCancellation = new(OriginalGroupHeader6)
	return u.OriginalGroupInformationAndCancellation
}

func (u *UnderlyingTransaction16) AddTransactionInformation() *PaymentTransaction75 {
	newValue := new(PaymentTransaction75)
	u.TransactionInformation = append(u.TransactionInformation, newValue)
	return newValue
}

// Identifies the underlying (group of) transaction(s) to which the resolution of investigation applies.
type UnderlyingTransaction17 struct {

	// Provides information on the original cancellation message, to which the resolution refers.
	OriginalGroupInformationAndStatus *OriginalGroupHeader5 `xml:"OrgnlGrpInfAndSts,omitempty"`

	// Provides information on the original (group of) transactions, to which the cancellation status refers.
	OriginalPaymentInformationAndStatus []*OriginalPaymentInstruction22 `xml:"OrgnlPmtInfAndSts,omitempty"`

	// Provides details on the original transactions to which the cancellation request message refers.
	TransactionInformationAndStatus []*PaymentTransaction79 `xml:"TxInfAndSts,omitempty"`
}

func (u *UnderlyingTransaction17) AddOriginalGroupInformationAndStatus() *OriginalGroupHeader5 {
	u.OriginalGroupInformationAndStatus = new(OriginalGroupHeader5)
	return u.OriginalGroupInformationAndStatus
}

func (u *UnderlyingTransaction17) AddOriginalPaymentInformationAndStatus() *OriginalPaymentInstruction22 {
	newValue := new(OriginalPaymentInstruction22)
	u.OriginalPaymentInformationAndStatus = append(u.OriginalPaymentInformationAndStatus, newValue)
	return newValue
}

func (u *UnderlyingTransaction17) AddTransactionInformationAndStatus() *PaymentTransaction79 {
	newValue := new(PaymentTransaction79)
	u.TransactionInformationAndStatus = append(u.TransactionInformationAndStatus, newValue)
	return newValue
}

// Set of elements used to identify the underlying (group of) transaction(s) to which the investigation applies.
type UnderlyingTransaction2 struct {

	// Set of elements used to provide information on the original messsage, to which the cancellation refers.
	OriginalGroupInformationAndCancellation *OriginalGroupInformation23 `xml:"OrgnlGrpInfAndCxl,omitempty"`

	// Set of elements used to provide information on the original transactions to which the cancellation request message refers.
	TransactionInformation []*PaymentTransactionInformation31 `xml:"TxInf,omitempty"`
}

func (u *UnderlyingTransaction2) AddOriginalGroupInformationAndCancellation() *OriginalGroupInformation23 {
	u.OriginalGroupInformationAndCancellation = new(OriginalGroupInformation23)
	return u.OriginalGroupInformationAndCancellation
}

func (u *UnderlyingTransaction2) AddTransactionInformation() *PaymentTransactionInformation31 {
	newValue := new(PaymentTransactionInformation31)
	u.TransactionInformation = append(u.TransactionInformation, newValue)
	return newValue
}

// Set of elements used to identify the underlying (group of) transaction(s) to which the resolution of investigation applies.
type UnderlyingTransaction3 struct {

	// Set of elements used to provide information on the original cancellation message, to which the resolution refers.
	OriginalGroupInformationAndStatus *OriginalGroupInformation24 `xml:"OrgnlGrpInfAndSts,omitempty"`

	// Set of elements used to provide information on the original (group of) transactions, to which the cancellation status refers.
	OriginalPaymentInformationAndStatus []*OriginalPaymentInformation3 `xml:"OrgnlPmtInfAndSts,omitempty"`

	// Set of elements used to provide information on the original transactions to which the cancellation request message refers.
	TransactionInformationAndStatus []*PaymentTransactionInformation33 `xml:"TxInfAndSts,omitempty"`
}

func (u *UnderlyingTransaction3) AddOriginalGroupInformationAndStatus() *OriginalGroupInformation24 {
	u.OriginalGroupInformationAndStatus = new(OriginalGroupInformation24)
	return u.OriginalGroupInformationAndStatus
}

func (u *UnderlyingTransaction3) AddOriginalPaymentInformationAndStatus() *OriginalPaymentInformation3 {
	newValue := new(OriginalPaymentInformation3)
	u.OriginalPaymentInformationAndStatus = append(u.OriginalPaymentInformationAndStatus, newValue)
	return newValue
}

func (u *UnderlyingTransaction3) AddTransactionInformationAndStatus() *PaymentTransactionInformation33 {
	newValue := new(PaymentTransactionInformation33)
	u.TransactionInformationAndStatus = append(u.TransactionInformationAndStatus, newValue)
	return newValue
}

// Identifies the underlying (group of) transaction(s) to which the resolution of investigation applies.
type UnderlyingTransaction4 struct {

	// Provides information on the original cancellation message, to which the resolution refers.
	OriginalGroupInformationAndStatus *OriginalGroupHeader5 `xml:"OrgnlGrpInfAndSts,omitempty"`

	// Provides information on the original (group of) transactions, to which the cancellation status refers.
	OriginalPaymentInformationAndStatus []*OriginalPaymentInstruction3 `xml:"OrgnlPmtInfAndSts,omitempty"`

	// Provides details on the original transactions to which the cancellation request message refers.
	TransactionInformationAndStatus []*PaymentTransaction40 `xml:"TxInfAndSts,omitempty"`
}

func (u *UnderlyingTransaction4) AddOriginalGroupInformationAndStatus() *OriginalGroupHeader5 {
	u.OriginalGroupInformationAndStatus = new(OriginalGroupHeader5)
	return u.OriginalGroupInformationAndStatus
}

func (u *UnderlyingTransaction4) AddOriginalPaymentInformationAndStatus() *OriginalPaymentInstruction3 {
	newValue := new(OriginalPaymentInstruction3)
	u.OriginalPaymentInformationAndStatus = append(u.OriginalPaymentInformationAndStatus, newValue)
	return newValue
}

func (u *UnderlyingTransaction4) AddTransactionInformationAndStatus() *PaymentTransaction40 {
	newValue := new(PaymentTransaction40)
	u.TransactionInformationAndStatus = append(u.TransactionInformationAndStatus, newValue)
	return newValue
}

// Identifies the underlying (group of) transaction(s) to which the investigation applies.
type UnderlyingTransaction5 struct {

	// Provides information on the original message, to which the cancellation refers.
	OriginalGroupInformationAndCancellation *OriginalGroupHeader4 `xml:"OrgnlGrpInfAndCxl,omitempty"`

	// Provides information on the original transactions to which the cancellation request message refers.
	TransactionInformation []*PaymentTransaction38 `xml:"TxInf,omitempty"`
}

func (u *UnderlyingTransaction5) AddOriginalGroupInformationAndCancellation() *OriginalGroupHeader4 {
	u.OriginalGroupInformationAndCancellation = new(OriginalGroupHeader4)
	return u.OriginalGroupInformationAndCancellation
}

func (u *UnderlyingTransaction5) AddTransactionInformation() *PaymentTransaction38 {
	newValue := new(PaymentTransaction38)
	u.TransactionInformation = append(u.TransactionInformation, newValue)
	return newValue
}

// Identifies the underlying (group of) transaction(s) to which the investigation applies.
type UnderlyingTransaction6 struct {

	// Provides information on the original message, to which the cancellation refers.
	OriginalGroupInformationAndCancellation *OriginalGroupHeader4 `xml:"OrgnlGrpInfAndCxl,omitempty"`

	// Provides information on the original (group of) transactions, to which the cancellation request refers.
	OriginalPaymentInformationAndCancellation []*OriginalPaymentInstruction4 `xml:"OrgnlPmtInfAndCxl,omitempty"`
}

func (u *UnderlyingTransaction6) AddOriginalGroupInformationAndCancellation() *OriginalGroupHeader4 {
	u.OriginalGroupInformationAndCancellation = new(OriginalGroupHeader4)
	return u.OriginalGroupInformationAndCancellation
}

func (u *UnderlyingTransaction6) AddOriginalPaymentInformationAndCancellation() *OriginalPaymentInstruction4 {
	newValue := new(OriginalPaymentInstruction4)
	u.OriginalPaymentInformationAndCancellation = append(u.OriginalPaymentInformationAndCancellation, newValue)
	return newValue
}

// Identifies the underlying (group of) transaction(s) to which the investigation applies.
type UnderlyingTransaction7 struct {

	// Provides information on the original message, to which the cancellation refers.
	OriginalGroupInformationAndCancellation *OriginalGroupHeader4 `xml:"OrgnlGrpInfAndCxl,omitempty"`

	// Provides information on the original (group of) transactions, to which the cancellation request refers.
	OriginalPaymentInformationAndCancellation []*OriginalPaymentInstruction8 `xml:"OrgnlPmtInfAndCxl,omitempty"`
}

func (u *UnderlyingTransaction7) AddOriginalGroupInformationAndCancellation() *OriginalGroupHeader4 {
	u.OriginalGroupInformationAndCancellation = new(OriginalGroupHeader4)
	return u.OriginalGroupInformationAndCancellation
}

func (u *UnderlyingTransaction7) AddOriginalPaymentInformationAndCancellation() *OriginalPaymentInstruction8 {
	newValue := new(OriginalPaymentInstruction8)
	u.OriginalPaymentInformationAndCancellation = append(u.OriginalPaymentInformationAndCancellation, newValue)
	return newValue
}

// Identifies the underlying (group of) transaction(s) to which the investigation applies.
type UnderlyingTransaction8 struct {

	// Provides information on the original message, to which the cancellation refers.
	OriginalGroupInformationAndCancellation *OriginalGroupHeader4 `xml:"OrgnlGrpInfAndCxl,omitempty"`

	// Provides information on the original transactions to which the cancellation request message refers.
	TransactionInformation []*PaymentTransaction48 `xml:"TxInf,omitempty"`
}

func (u *UnderlyingTransaction8) AddOriginalGroupInformationAndCancellation() *OriginalGroupHeader4 {
	u.OriginalGroupInformationAndCancellation = new(OriginalGroupHeader4)
	return u.OriginalGroupInformationAndCancellation
}

func (u *UnderlyingTransaction8) AddTransactionInformation() *PaymentTransaction48 {
	newValue := new(PaymentTransaction48)
	u.TransactionInformation = append(u.TransactionInformation, newValue)
	return newValue
}

// Identifies the underlying (group of) transaction(s) to which the resolution of investigation applies.
type UnderlyingTransaction9 struct {

	// Provides information on the original cancellation message, to which the resolution refers.
	OriginalGroupInformationAndStatus *OriginalGroupHeader5 `xml:"OrgnlGrpInfAndSts,omitempty"`

	// Provides information on the original (group of) transactions, to which the cancellation status refers.
	OriginalPaymentInformationAndStatus []*OriginalPaymentInstruction10 `xml:"OrgnlPmtInfAndSts,omitempty"`

	// Provides details on the original transactions to which the cancellation request message refers.
	TransactionInformationAndStatus []*PaymentTransaction53 `xml:"TxInfAndSts,omitempty"`
}

func (u *UnderlyingTransaction9) AddOriginalGroupInformationAndStatus() *OriginalGroupHeader5 {
	u.OriginalGroupInformationAndStatus = new(OriginalGroupHeader5)
	return u.OriginalGroupInformationAndStatus
}

func (u *UnderlyingTransaction9) AddOriginalPaymentInformationAndStatus() *OriginalPaymentInstruction10 {
	newValue := new(OriginalPaymentInstruction10)
	u.OriginalPaymentInformationAndStatus = append(u.OriginalPaymentInformationAndStatus, newValue)
	return newValue
}

func (u *UnderlyingTransaction9) AddTransactionInformationAndStatus() *PaymentTransaction53 {
	newValue := new(PaymentTransaction53)
	u.TransactionInformationAndStatus = append(u.TransactionInformationAndStatus, newValue)
	return newValue
}
