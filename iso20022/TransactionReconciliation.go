package iso20022

// Reconciliation transaction between an acceptor and an acquirer.
type TransactionReconciliation1 struct {

	// Indicates if the transaction requires a closure of the reconciliation period.
	ClosePeriod *TrueFalseIndicator `xml:"ClsPrd,omitempty"`

	// Unique identification of a reconciliation transaction.
	ReconciliationTransactionIdentification *TransactionIdentifier1 `xml:"RcncltnTxId"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId"`

	// Transaction totals during the reconciliation period for a certain type of transaction.
	TransactionTotals []*TransactionTotals1 `xml:"TxTtls"`

	// Additional information related to the reconciliation transaction.
	AdditionalTransactionData *Max70Text `xml:"AddtlTxData,omitempty"`
}

func (t *TransactionReconciliation1) SetClosePeriod(value string) {
	t.ClosePeriod = (*TrueFalseIndicator)(&value)
}

func (t *TransactionReconciliation1) AddReconciliationTransactionIdentification() *TransactionIdentifier1 {
	t.ReconciliationTransactionIdentification = new(TransactionIdentifier1)
	return t.ReconciliationTransactionIdentification
}

func (t *TransactionReconciliation1) SetReconciliationIdentification(value string) {
	t.ReconciliationIdentification = (*Max35Text)(&value)
}

func (t *TransactionReconciliation1) AddTransactionTotals() *TransactionTotals1 {
	newValue := new(TransactionTotals1)
	t.TransactionTotals = append(t.TransactionTotals, newValue)
	return newValue
}

func (t *TransactionReconciliation1) SetAdditionalTransactionData(value string) {
	t.AdditionalTransactionData = (*Max70Text)(&value)
}

// Reconciliation transaction between an acceptor and an acquirer.
type TransactionReconciliation2 struct {

	// Indicates if the transaction requires a closure of the reconciliation period.
	ClosePeriod *TrueFalseIndicator `xml:"ClsPrd,omitempty"`

	// Unique identification of a reconciliation transaction.
	ReconciliationTransactionIdentification *TransactionIdentifier1 `xml:"RcncltnTxId"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId"`

	// Transaction totals during the reconciliation period for a certain type of transaction.
	TransactionTotals []*TransactionTotals2 `xml:"TxTtls,omitempty"`

	// Additional information related to the reconciliation transaction.
	AdditionalTransactionData *Max70Text `xml:"AddtlTxData,omitempty"`
}

func (t *TransactionReconciliation2) SetClosePeriod(value string) {
	t.ClosePeriod = (*TrueFalseIndicator)(&value)
}

func (t *TransactionReconciliation2) AddReconciliationTransactionIdentification() *TransactionIdentifier1 {
	t.ReconciliationTransactionIdentification = new(TransactionIdentifier1)
	return t.ReconciliationTransactionIdentification
}

func (t *TransactionReconciliation2) SetReconciliationIdentification(value string) {
	t.ReconciliationIdentification = (*Max35Text)(&value)
}

func (t *TransactionReconciliation2) AddTransactionTotals() *TransactionTotals2 {
	newValue := new(TransactionTotals2)
	t.TransactionTotals = append(t.TransactionTotals, newValue)
	return newValue
}

func (t *TransactionReconciliation2) SetAdditionalTransactionData(value string) {
	t.AdditionalTransactionData = (*Max70Text)(&value)
}

// Reconciliation transaction between an acceptor and an acquirer.
type TransactionReconciliation3 struct {

	// Indicates if the transaction requires a closure of the reconciliation period.
	ClosePeriod *TrueFalseIndicator `xml:"ClsPrd,omitempty"`

	// Unique identification of a reconciliation transaction.
	ReconciliationTransactionIdentification *TransactionIdentifier1 `xml:"RcncltnTxId"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId"`

	// Transaction totals during the reconciliation period for a certain type of transaction.
	TransactionTotals []*TransactionTotals3 `xml:"TxTtls,omitempty"`

	// Additional information related to the reconciliation transaction.
	AdditionalTransactionData *Max70Text `xml:"AddtlTxData,omitempty"`
}

func (t *TransactionReconciliation3) SetClosePeriod(value string) {
	t.ClosePeriod = (*TrueFalseIndicator)(&value)
}

func (t *TransactionReconciliation3) AddReconciliationTransactionIdentification() *TransactionIdentifier1 {
	t.ReconciliationTransactionIdentification = new(TransactionIdentifier1)
	return t.ReconciliationTransactionIdentification
}

func (t *TransactionReconciliation3) SetReconciliationIdentification(value string) {
	t.ReconciliationIdentification = (*Max35Text)(&value)
}

func (t *TransactionReconciliation3) AddTransactionTotals() *TransactionTotals3 {
	newValue := new(TransactionTotals3)
	t.TransactionTotals = append(t.TransactionTotals, newValue)
	return newValue
}

func (t *TransactionReconciliation3) SetAdditionalTransactionData(value string) {
	t.AdditionalTransactionData = (*Max70Text)(&value)
}

// Reconciliation transaction between an acceptor and an acquirer.
type TransactionReconciliation4 struct {

	// Indicates if the transaction requires a closure of the reconciliation period.
	ClosePeriod *TrueFalseIndicator `xml:"ClsPrd,omitempty"`

	// Unique identification of a reconciliation transaction.
	ReconciliationTransactionIdentification *TransactionIdentifier1 `xml:"RcncltnTxId"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId"`

	// Transaction totals during the reconciliation period for a certain type of transaction.
	TransactionTotals []*TransactionTotals7 `xml:"TxTtls,omitempty"`

	// Additional information related to the reconciliation transaction.
	AdditionalTransactionData *Max70Text `xml:"AddtlTxData,omitempty"`
}

func (t *TransactionReconciliation4) SetClosePeriod(value string) {
	t.ClosePeriod = (*TrueFalseIndicator)(&value)
}

func (t *TransactionReconciliation4) AddReconciliationTransactionIdentification() *TransactionIdentifier1 {
	t.ReconciliationTransactionIdentification = new(TransactionIdentifier1)
	return t.ReconciliationTransactionIdentification
}

func (t *TransactionReconciliation4) SetReconciliationIdentification(value string) {
	t.ReconciliationIdentification = (*Max35Text)(&value)
}

func (t *TransactionReconciliation4) AddTransactionTotals() *TransactionTotals7 {
	newValue := new(TransactionTotals7)
	t.TransactionTotals = append(t.TransactionTotals, newValue)
	return newValue
}

func (t *TransactionReconciliation4) SetAdditionalTransactionData(value string) {
	t.AdditionalTransactionData = (*Max70Text)(&value)
}
