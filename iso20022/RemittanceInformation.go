package iso20022

// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, eg, commercial invoices in an accounts' receivable system.
type RemittanceInformation1 struct {

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, eg, commercial invoices in an accounts' receivable system in an unstructured form.
	Unstructured []*Max140Text `xml:"Ustrd,omitempty"`

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, eg, commercial invoices in an accounts' receivable system in a structured form.
	Structured []*StructuredRemittanceInformation6 `xml:"Strd,omitempty"`
}

func (r *RemittanceInformation1) AddUnstructured(value string) {
	r.Unstructured = append(r.Unstructured, (*Max140Text)(&value))
}

func (r *RemittanceInformation1) AddStructured() *StructuredRemittanceInformation6 {
	newValue := new(StructuredRemittanceInformation6)
	r.Structured = append(r.Structured, newValue)
	return newValue
}

// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system.
type RemittanceInformation10 struct {

	// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in an unstructured form.
	Unstructured []*Max140Text `xml:"Ustrd,omitempty"`

	// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in a structured form.
	Structured []*StructuredRemittanceInformation12 `xml:"Strd,omitempty"`
}

func (r *RemittanceInformation10) AddUnstructured(value string) {
	r.Unstructured = append(r.Unstructured, (*Max140Text)(&value))
}

func (r *RemittanceInformation10) AddStructured() *StructuredRemittanceInformation12 {
	newValue := new(StructuredRemittanceInformation12)
	r.Structured = append(r.Structured, newValue)
	return newValue
}

// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system.
type RemittanceInformation11 struct {

	// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in an unstructured form.
	Unstructured []*Max140Text `xml:"Ustrd,omitempty"`

	// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in a structured form.
	Structured []*StructuredRemittanceInformation13 `xml:"Strd,omitempty"`
}

func (r *RemittanceInformation11) AddUnstructured(value string) {
	r.Unstructured = append(r.Unstructured, (*Max140Text)(&value))
}

func (r *RemittanceInformation11) AddStructured() *StructuredRemittanceInformation13 {
	newValue := new(StructuredRemittanceInformation13)
	r.Structured = append(r.Structured, newValue)
	return newValue
}

// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system.
type RemittanceInformation12 struct {

	// Unique identification, assigned by the originator, to unambiguously identify the remittance information within the message.
	RemittanceIdentification *Max35Text `xml:"RmtId,omitempty"`

	// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in an unstructured form.
	Unstructured []*Max140Text `xml:"Ustrd,omitempty"`

	// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in a structured form.
	Structured []*StructuredRemittanceInformation13 `xml:"Strd,omitempty"`

	// Set of elements used to provide information on the original transactions, to which the remittance message refers.
	OriginalPaymentInformation *OriginalPaymentInformation6 `xml:"OrgnlPmtInf"`
}

func (r *RemittanceInformation12) SetRemittanceIdentification(value string) {
	r.RemittanceIdentification = (*Max35Text)(&value)
}

func (r *RemittanceInformation12) AddUnstructured(value string) {
	r.Unstructured = append(r.Unstructured, (*Max140Text)(&value))
}

func (r *RemittanceInformation12) AddStructured() *StructuredRemittanceInformation13 {
	newValue := new(StructuredRemittanceInformation13)
	r.Structured = append(r.Structured, newValue)
	return newValue
}

func (r *RemittanceInformation12) AddOriginalPaymentInformation() *OriginalPaymentInformation6 {
	r.OriginalPaymentInformation = new(OriginalPaymentInformation6)
	return r.OriginalPaymentInformation
}

// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system.
type RemittanceInformation13 struct {

	// Unique identification, assigned by the originator, to unambiguously identify the remittance information within the message.
	RemittanceIdentification *Max35Text `xml:"RmtId,omitempty"`

	// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in an unstructured form.
	Unstructured []*Max140Text `xml:"Ustrd,omitempty"`

	// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in a structured form.
	Structured []*StructuredRemittanceInformation13 `xml:"Strd,omitempty"`

	// Set of elements used to provide information on the original transactions, to which the remittance message refers.
	OriginalPaymentInformation *OriginalPaymentInformation7 `xml:"OrgnlPmtInf"`
}

func (r *RemittanceInformation13) SetRemittanceIdentification(value string) {
	r.RemittanceIdentification = (*Max35Text)(&value)
}

func (r *RemittanceInformation13) AddUnstructured(value string) {
	r.Unstructured = append(r.Unstructured, (*Max140Text)(&value))
}

func (r *RemittanceInformation13) AddStructured() *StructuredRemittanceInformation13 {
	newValue := new(StructuredRemittanceInformation13)
	r.Structured = append(r.Structured, newValue)
	return newValue
}

func (r *RemittanceInformation13) AddOriginalPaymentInformation() *OriginalPaymentInformation7 {
	r.OriginalPaymentInformation = new(OriginalPaymentInformation7)
	return r.OriginalPaymentInformation
}

// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle.
type RemittanceInformation2 struct {

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, eg, commercial invoices in an accounts' receivable system in an unstructured form.
	Unstructured []*Max140Text `xml:"Ustrd,omitempty"`
}

func (r *RemittanceInformation2) AddUnstructured(value string) {
	r.Unstructured = append(r.Unstructured, (*Max140Text)(&value))
}

// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system.
type RemittanceInformation5 struct {

	// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in an unstructured form.
	Unstructured []*Max140Text `xml:"Ustrd,omitempty"`

	// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in a structured form.
	Structured []*StructuredRemittanceInformation7 `xml:"Strd,omitempty"`
}

func (r *RemittanceInformation5) AddUnstructured(value string) {
	r.Unstructured = append(r.Unstructured, (*Max140Text)(&value))
}

func (r *RemittanceInformation5) AddStructured() *StructuredRemittanceInformation7 {
	newValue := new(StructuredRemittanceInformation7)
	r.Structured = append(r.Structured, newValue)
	return newValue
}

// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system.
type RemittanceInformation6 struct {

	// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in an unstructured form.
	Unstructured []*Max140Text `xml:"Ustrd,omitempty"`

	// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in a structured form.
	Structured []*StructuredRemittanceInformation8 `xml:"Strd,omitempty"`
}

func (r *RemittanceInformation6) AddUnstructured(value string) {
	r.Unstructured = append(r.Unstructured, (*Max140Text)(&value))
}

func (r *RemittanceInformation6) AddStructured() *StructuredRemittanceInformation8 {
	newValue := new(StructuredRemittanceInformation8)
	r.Structured = append(r.Structured, newValue)
	return newValue
}

// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system.
type RemittanceInformation7 struct {

	// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in an unstructured form.
	Unstructured []*Max140Text `xml:"Ustrd,omitempty"`

	// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in a structured form.
	Structured []*StructuredRemittanceInformation9 `xml:"Strd,omitempty"`
}

func (r *RemittanceInformation7) AddUnstructured(value string) {
	r.Unstructured = append(r.Unstructured, (*Max140Text)(&value))
}

func (r *RemittanceInformation7) AddStructured() *StructuredRemittanceInformation9 {
	newValue := new(StructuredRemittanceInformation9)
	r.Structured = append(r.Structured, newValue)
	return newValue
}

// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system.
type RemittanceInformation8 struct {

	// Unique identification, assigned by the originator, to unambiguously identify the remittance information within the message.
	RemittanceIdentification *Max35Text `xml:"RmtId,omitempty"`

	// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in an unstructured form.
	Unstructured []*Max140Text `xml:"Ustrd,omitempty"`

	// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in a structured form.
	Structured []*StructuredRemittanceInformation10 `xml:"Strd,omitempty"`

	// Set of elements used to provide information on the original transactions, to which the remittance message refers.
	OriginalPaymentInformation *OriginalPaymentInformation6 `xml:"OrgnlPmtInf"`
}

func (r *RemittanceInformation8) SetRemittanceIdentification(value string) {
	r.RemittanceIdentification = (*Max35Text)(&value)
}

func (r *RemittanceInformation8) AddUnstructured(value string) {
	r.Unstructured = append(r.Unstructured, (*Max140Text)(&value))
}

func (r *RemittanceInformation8) AddStructured() *StructuredRemittanceInformation10 {
	newValue := new(StructuredRemittanceInformation10)
	r.Structured = append(r.Structured, newValue)
	return newValue
}

func (r *RemittanceInformation8) AddOriginalPaymentInformation() *OriginalPaymentInformation6 {
	r.OriginalPaymentInformation = new(OriginalPaymentInformation6)
	return r.OriginalPaymentInformation
}
