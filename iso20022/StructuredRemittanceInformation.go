package iso20022

// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in a structured form.
type StructuredRemittanceInformation10 struct {

	// Set of elements used to provide the content of the referred document.
	ReferredDocumentInformation []*ReferredDocumentInformation4 `xml:"RfrdDocInf,omitempty"`

	// Provides details on the amounts of the referred document.
	ReferredDocumentAmount *RemittanceAmount2 `xml:"RfrdDocAmt,omitempty"`

	// Reference information provided by the creditor to allow the identification of the underlying documents.
	CreditorReferenceInformation *CreditorReferenceInformation2 `xml:"CdtrRefInf,omitempty"`

	// Identification of the organisation issuing the invoice, when it is different from the creditor or ultimate creditor.
	Invoicer *PartyIdentification43 `xml:"Invcr,omitempty"`

	// Identification of the party to whom an invoice is issued, when it is different from the debtor or ultimate debtor.
	Invoicee *PartyIdentification43 `xml:"Invcee,omitempty"`

	// Provides remittance information about a payment made for tax-related purposes.
	TaxRemittance *TaxInformation4 `xml:"TaxRmt,omitempty"`

	// Provides remittance information about a payment for garnishment-related purposes.
	GarnishmentRemittance *Garnishment1 `xml:"GrnshmtRmt,omitempty"`

	// Additional information, in free text form, to complement the structured remittance information.
	AdditionalRemittanceInformation []*Max140Text `xml:"AddtlRmtInf,omitempty"`
}

func (s *StructuredRemittanceInformation10) AddReferredDocumentInformation() *ReferredDocumentInformation4 {
	newValue := new(ReferredDocumentInformation4)
	s.ReferredDocumentInformation = append(s.ReferredDocumentInformation, newValue)
	return newValue
}

func (s *StructuredRemittanceInformation10) AddReferredDocumentAmount() *RemittanceAmount2 {
	s.ReferredDocumentAmount = new(RemittanceAmount2)
	return s.ReferredDocumentAmount
}

func (s *StructuredRemittanceInformation10) AddCreditorReferenceInformation() *CreditorReferenceInformation2 {
	s.CreditorReferenceInformation = new(CreditorReferenceInformation2)
	return s.CreditorReferenceInformation
}

func (s *StructuredRemittanceInformation10) AddInvoicer() *PartyIdentification43 {
	s.Invoicer = new(PartyIdentification43)
	return s.Invoicer
}

func (s *StructuredRemittanceInformation10) AddInvoicee() *PartyIdentification43 {
	s.Invoicee = new(PartyIdentification43)
	return s.Invoicee
}

func (s *StructuredRemittanceInformation10) AddTaxRemittance() *TaxInformation4 {
	s.TaxRemittance = new(TaxInformation4)
	return s.TaxRemittance
}

func (s *StructuredRemittanceInformation10) AddGarnishmentRemittance() *Garnishment1 {
	s.GarnishmentRemittance = new(Garnishment1)
	return s.GarnishmentRemittance
}

func (s *StructuredRemittanceInformation10) AddAdditionalRemittanceInformation(value string) {
	s.AdditionalRemittanceInformation = append(s.AdditionalRemittanceInformation, (*Max140Text)(&value))
}

// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in a structured form.
type StructuredRemittanceInformation12 struct {

	// Set of elements used to identify the documents referred to in the remittance information.
	ReferredDocumentInformation []*ReferredDocumentInformation6 `xml:"RfrdDocInf,omitempty"`

	// Provides details on the amounts of the referred document.
	ReferredDocumentAmount *RemittanceAmount2 `xml:"RfrdDocAmt,omitempty"`

	// Reference information provided by the creditor to allow the identification of the underlying documents.
	CreditorReferenceInformation *CreditorReferenceInformation2 `xml:"CdtrRefInf,omitempty"`

	// Identification of the organisation issuing the invoice, when it is different from the creditor or ultimate creditor.
	Invoicer *PartyIdentification43 `xml:"Invcr,omitempty"`

	// Identification of the party to whom an invoice is issued, when it is different from the debtor or ultimate debtor.
	Invoicee *PartyIdentification43 `xml:"Invcee,omitempty"`

	// Provides remittance information about a payment made for tax-related purposes.
	TaxRemittance *TaxInformation4 `xml:"TaxRmt,omitempty"`

	// Provides remittance information about a payment for garnishment-related purposes.
	GarnishmentRemittance *Garnishment1 `xml:"GrnshmtRmt,omitempty"`

	// Additional information, in free text form, to complement the structured remittance information.
	AdditionalRemittanceInformation []*Max140Text `xml:"AddtlRmtInf,omitempty"`
}

func (s *StructuredRemittanceInformation12) AddReferredDocumentInformation() *ReferredDocumentInformation6 {
	newValue := new(ReferredDocumentInformation6)
	s.ReferredDocumentInformation = append(s.ReferredDocumentInformation, newValue)
	return newValue
}

func (s *StructuredRemittanceInformation12) AddReferredDocumentAmount() *RemittanceAmount2 {
	s.ReferredDocumentAmount = new(RemittanceAmount2)
	return s.ReferredDocumentAmount
}

func (s *StructuredRemittanceInformation12) AddCreditorReferenceInformation() *CreditorReferenceInformation2 {
	s.CreditorReferenceInformation = new(CreditorReferenceInformation2)
	return s.CreditorReferenceInformation
}

func (s *StructuredRemittanceInformation12) AddInvoicer() *PartyIdentification43 {
	s.Invoicer = new(PartyIdentification43)
	return s.Invoicer
}

func (s *StructuredRemittanceInformation12) AddInvoicee() *PartyIdentification43 {
	s.Invoicee = new(PartyIdentification43)
	return s.Invoicee
}

func (s *StructuredRemittanceInformation12) AddTaxRemittance() *TaxInformation4 {
	s.TaxRemittance = new(TaxInformation4)
	return s.TaxRemittance
}

func (s *StructuredRemittanceInformation12) AddGarnishmentRemittance() *Garnishment1 {
	s.GarnishmentRemittance = new(Garnishment1)
	return s.GarnishmentRemittance
}

func (s *StructuredRemittanceInformation12) AddAdditionalRemittanceInformation(value string) {
	s.AdditionalRemittanceInformation = append(s.AdditionalRemittanceInformation, (*Max140Text)(&value))
}

// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in a structured form.
type StructuredRemittanceInformation13 struct {

	// Provides the identification and the content of the referred document.
	ReferredDocumentInformation []*ReferredDocumentInformation7 `xml:"RfrdDocInf,omitempty"`

	// Provides details on the amounts of the referred document.
	ReferredDocumentAmount *RemittanceAmount2 `xml:"RfrdDocAmt,omitempty"`

	// Reference information provided by the creditor to allow the identification of the underlying documents.
	CreditorReferenceInformation *CreditorReferenceInformation2 `xml:"CdtrRefInf,omitempty"`

	// Identification of the organisation issuing the invoice, when it is different from the creditor or ultimate creditor.
	Invoicer *PartyIdentification43 `xml:"Invcr,omitempty"`

	// Identification of the party to whom an invoice is issued, when it is different from the debtor or ultimate debtor.
	Invoicee *PartyIdentification43 `xml:"Invcee,omitempty"`

	// Provides remittance information about a payment made for tax-related purposes.
	TaxRemittance *TaxInformation4 `xml:"TaxRmt,omitempty"`

	// Provides remittance information about a payment for garnishment-related purposes.
	GarnishmentRemittance *Garnishment1 `xml:"GrnshmtRmt,omitempty"`

	// Additional information, in free text form, to complement the structured remittance information.
	AdditionalRemittanceInformation []*Max140Text `xml:"AddtlRmtInf,omitempty"`
}

func (s *StructuredRemittanceInformation13) AddReferredDocumentInformation() *ReferredDocumentInformation7 {
	newValue := new(ReferredDocumentInformation7)
	s.ReferredDocumentInformation = append(s.ReferredDocumentInformation, newValue)
	return newValue
}

func (s *StructuredRemittanceInformation13) AddReferredDocumentAmount() *RemittanceAmount2 {
	s.ReferredDocumentAmount = new(RemittanceAmount2)
	return s.ReferredDocumentAmount
}

func (s *StructuredRemittanceInformation13) AddCreditorReferenceInformation() *CreditorReferenceInformation2 {
	s.CreditorReferenceInformation = new(CreditorReferenceInformation2)
	return s.CreditorReferenceInformation
}

func (s *StructuredRemittanceInformation13) AddInvoicer() *PartyIdentification43 {
	s.Invoicer = new(PartyIdentification43)
	return s.Invoicer
}

func (s *StructuredRemittanceInformation13) AddInvoicee() *PartyIdentification43 {
	s.Invoicee = new(PartyIdentification43)
	return s.Invoicee
}

func (s *StructuredRemittanceInformation13) AddTaxRemittance() *TaxInformation4 {
	s.TaxRemittance = new(TaxInformation4)
	return s.TaxRemittance
}

func (s *StructuredRemittanceInformation13) AddGarnishmentRemittance() *Garnishment1 {
	s.GarnishmentRemittance = new(Garnishment1)
	return s.GarnishmentRemittance
}

func (s *StructuredRemittanceInformation13) AddAdditionalRemittanceInformation(value string) {
	s.AdditionalRemittanceInformation = append(s.AdditionalRemittanceInformation, (*Max140Text)(&value))
}

// Structured information supplied to enable the matching, ie, reconciliation, of a payment with the items that the payment is intended to settle, eg, commercial invoices in an accounts receivable system.
type StructuredRemittanceInformation2 struct {

	// Specifies the nature of the referred document/transaction, eg, invoice or credit note.
	ReferredDocumentType *DocumentType1Code `xml:"RfrdDocTp,omitempty"`

	// Date associated with the referred document, eg, date of issue.
	ReferredDocumentRelatedDate *ISODate `xml:"RfrdDocRltdDt,omitempty"`

	// Amount of money and currency of a document referred to in the remittance section. The amount is typically either the original amount due and payable, or the amount actually remitted for the referred document.
	ReferredDocumentAmount []*ReferredDocumentAmount1Choice `xml:"RfrdDocAmt,omitempty"`

	// Unique and unambiguous identification of a document that distinguishes that document from another document referred to in the remittance information, usually assigned by the originator of the referred document/transaction.
	DocumentReferenceNumber *Max35Text `xml:"DocRefNb,omitempty"`

	// Unique and unambiguous reference assigned by the creditor to refer to the payment transaction.
	//
	// Usage: if available, the initiating party should provide this reference in the structured remittance information, to enable reconciliation by the creditor upon receipt of the cash.
	//
	// If the business context requires the use of a creditor reference or a payment remit identification, and only one identifier can be passed through the end-to-end chain, the creditor's reference or payment remittance identification should be quoted in the end-to-end transaction identification.
	CreditorReference *Max35Text `xml:"CdtrRef,omitempty"`

	// Identification of the organization issuing the invoice when different than the creditor or final party
	Invoicer *PartyIdentification1 `xml:"Invcr,omitempty"`

	// Identification of the party to whom an invoice is issued, when different than the originator or debtor.
	Invoicee *PartyIdentification1 `xml:"Invcee,omitempty"`
}

func (s *StructuredRemittanceInformation2) SetReferredDocumentType(value string) {
	s.ReferredDocumentType = (*DocumentType1Code)(&value)
}

func (s *StructuredRemittanceInformation2) SetReferredDocumentRelatedDate(value string) {
	s.ReferredDocumentRelatedDate = (*ISODate)(&value)
}

func (s *StructuredRemittanceInformation2) AddReferredDocumentAmount() *ReferredDocumentAmount1Choice {
	newValue := new(ReferredDocumentAmount1Choice)
	s.ReferredDocumentAmount = append(s.ReferredDocumentAmount, newValue)
	return newValue
}

func (s *StructuredRemittanceInformation2) SetDocumentReferenceNumber(value string) {
	s.DocumentReferenceNumber = (*Max35Text)(&value)
}

func (s *StructuredRemittanceInformation2) SetCreditorReference(value string) {
	s.CreditorReference = (*Max35Text)(&value)
}

func (s *StructuredRemittanceInformation2) AddInvoicer() *PartyIdentification1 {
	s.Invoicer = new(PartyIdentification1)
	return s.Invoicer
}

func (s *StructuredRemittanceInformation2) AddInvoicee() *PartyIdentification1 {
	s.Invoicee = new(PartyIdentification1)
	return s.Invoicee
}

// Structured information supplied to enable the matching, i.e.  reconciliation, of a payment with the items that the payment is intended to settle, eg, commercial invoices in an accounts receivable system.
type StructuredRemittanceInformation6 struct {

	// Reference information to allow the identification of the underlying reference documents.
	ReferredDocumentInformation *ReferredDocumentInformation1 `xml:"RfrdDocInf,omitempty"`

	// Date associated with the referred document, eg, date of issue.
	ReferredDocumentRelatedDate *ISODate `xml:"RfrdDocRltdDt,omitempty"`

	// Amount of money and currency of a document referred to in the remittance section. The amount is typically either the original amount due and payable, or the amount actually remitted for the referred document.
	ReferredDocumentAmount []*ReferredDocumentAmount1Choice `xml:"RfrdDocAmt,omitempty"`

	// Reference information provided by the creditor to allow the identification of the underlying documents.
	CreditorReferenceInformation *CreditorReferenceInformation1 `xml:"CdtrRefInf,omitempty"`

	// Identification of the organization issuing the invoice when different than the creditor or final party.
	Invoicer *PartyIdentification8 `xml:"Invcr,omitempty"`

	// Identification of the party to whom an invoice is issued, when different than the originator or debtor.
	Invoicee *PartyIdentification8 `xml:"Invcee,omitempty"`

	// Additional information, in free text form, to complement the structured remittance information.
	AdditionalRemittanceInformation *Max140Text `xml:"AddtlRmtInf,omitempty"`
}

func (s *StructuredRemittanceInformation6) AddReferredDocumentInformation() *ReferredDocumentInformation1 {
	s.ReferredDocumentInformation = new(ReferredDocumentInformation1)
	return s.ReferredDocumentInformation
}

func (s *StructuredRemittanceInformation6) SetReferredDocumentRelatedDate(value string) {
	s.ReferredDocumentRelatedDate = (*ISODate)(&value)
}

func (s *StructuredRemittanceInformation6) AddReferredDocumentAmount() *ReferredDocumentAmount1Choice {
	newValue := new(ReferredDocumentAmount1Choice)
	s.ReferredDocumentAmount = append(s.ReferredDocumentAmount, newValue)
	return newValue
}

func (s *StructuredRemittanceInformation6) AddCreditorReferenceInformation() *CreditorReferenceInformation1 {
	s.CreditorReferenceInformation = new(CreditorReferenceInformation1)
	return s.CreditorReferenceInformation
}

func (s *StructuredRemittanceInformation6) AddInvoicer() *PartyIdentification8 {
	s.Invoicer = new(PartyIdentification8)
	return s.Invoicer
}

func (s *StructuredRemittanceInformation6) AddInvoicee() *PartyIdentification8 {
	s.Invoicee = new(PartyIdentification8)
	return s.Invoicee
}

func (s *StructuredRemittanceInformation6) SetAdditionalRemittanceInformation(value string) {
	s.AdditionalRemittanceInformation = (*Max140Text)(&value)
}

// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in a structured form.
type StructuredRemittanceInformation7 struct {

	// Set of elements used to identify the documents referred to in the remittance information.
	ReferredDocumentInformation []*ReferredDocumentInformation3 `xml:"RfrdDocInf,omitempty"`

	// Set of elements used to provide details on the amounts of the referred document.
	ReferredDocumentAmount *RemittanceAmount1 `xml:"RfrdDocAmt,omitempty"`

	// Reference information provided by the creditor to allow the identification of the underlying documents.
	CreditorReferenceInformation *CreditorReferenceInformation2 `xml:"CdtrRefInf,omitempty"`

	// Identification of the organisation issuing the invoice, when it is different from the creditor or ultimate creditor.
	Invoicer *PartyIdentification32 `xml:"Invcr,omitempty"`

	// Identification of the party to whom an invoice is issued, when it is different from the debtor or ultimate debtor.
	Invoicee *PartyIdentification32 `xml:"Invcee,omitempty"`

	// Additional information, in free text form, to complement the structured remittance information.
	AdditionalRemittanceInformation []*Max140Text `xml:"AddtlRmtInf,omitempty"`
}

func (s *StructuredRemittanceInformation7) AddReferredDocumentInformation() *ReferredDocumentInformation3 {
	newValue := new(ReferredDocumentInformation3)
	s.ReferredDocumentInformation = append(s.ReferredDocumentInformation, newValue)
	return newValue
}

func (s *StructuredRemittanceInformation7) AddReferredDocumentAmount() *RemittanceAmount1 {
	s.ReferredDocumentAmount = new(RemittanceAmount1)
	return s.ReferredDocumentAmount
}

func (s *StructuredRemittanceInformation7) AddCreditorReferenceInformation() *CreditorReferenceInformation2 {
	s.CreditorReferenceInformation = new(CreditorReferenceInformation2)
	return s.CreditorReferenceInformation
}

func (s *StructuredRemittanceInformation7) AddInvoicer() *PartyIdentification32 {
	s.Invoicer = new(PartyIdentification32)
	return s.Invoicer
}

func (s *StructuredRemittanceInformation7) AddInvoicee() *PartyIdentification32 {
	s.Invoicee = new(PartyIdentification32)
	return s.Invoicee
}

func (s *StructuredRemittanceInformation7) AddAdditionalRemittanceInformation(value string) {
	s.AdditionalRemittanceInformation = append(s.AdditionalRemittanceInformation, (*Max140Text)(&value))
}

// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in a structured form.
type StructuredRemittanceInformation8 struct {

	// Set of elements used to identify the documents referred to in the remittance information.
	ReferredDocumentInformation []*ReferredDocumentInformation3 `xml:"RfrdDocInf,omitempty"`

	// Set of elements used to provide details on the amounts of the referred document.
	ReferredDocumentAmount *RemittanceAmount1 `xml:"RfrdDocAmt,omitempty"`

	// Reference information provided by the creditor to allow the identification of the underlying documents.
	CreditorReferenceInformation *CreditorReferenceInformation2 `xml:"CdtrRefInf,omitempty"`

	// Identification of the organisation issuing the invoice, when it is different from the creditor or ultimate creditor.
	Invoicer *PartyIdentification43 `xml:"Invcr,omitempty"`

	// Identification of the party to whom an invoice is issued, when it is different from the debtor or ultimate debtor.
	Invoicee *PartyIdentification43 `xml:"Invcee,omitempty"`

	// Additional information, in free text form, to complement the structured remittance information.
	AdditionalRemittanceInformation []*Max140Text `xml:"AddtlRmtInf,omitempty"`
}

func (s *StructuredRemittanceInformation8) AddReferredDocumentInformation() *ReferredDocumentInformation3 {
	newValue := new(ReferredDocumentInformation3)
	s.ReferredDocumentInformation = append(s.ReferredDocumentInformation, newValue)
	return newValue
}

func (s *StructuredRemittanceInformation8) AddReferredDocumentAmount() *RemittanceAmount1 {
	s.ReferredDocumentAmount = new(RemittanceAmount1)
	return s.ReferredDocumentAmount
}

func (s *StructuredRemittanceInformation8) AddCreditorReferenceInformation() *CreditorReferenceInformation2 {
	s.CreditorReferenceInformation = new(CreditorReferenceInformation2)
	return s.CreditorReferenceInformation
}

func (s *StructuredRemittanceInformation8) AddInvoicer() *PartyIdentification43 {
	s.Invoicer = new(PartyIdentification43)
	return s.Invoicer
}

func (s *StructuredRemittanceInformation8) AddInvoicee() *PartyIdentification43 {
	s.Invoicee = new(PartyIdentification43)
	return s.Invoicee
}

func (s *StructuredRemittanceInformation8) AddAdditionalRemittanceInformation(value string) {
	s.AdditionalRemittanceInformation = append(s.AdditionalRemittanceInformation, (*Max140Text)(&value))
}

// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in a structured form.
type StructuredRemittanceInformation9 struct {

	// Set of elements used to identify the documents referred to in the remittance information.
	ReferredDocumentInformation []*ReferredDocumentInformation3 `xml:"RfrdDocInf,omitempty"`

	// Provides details on the amounts of the referred document.
	ReferredDocumentAmount *RemittanceAmount2 `xml:"RfrdDocAmt,omitempty"`

	// Reference information provided by the creditor to allow the identification of the underlying documents.
	CreditorReferenceInformation *CreditorReferenceInformation2 `xml:"CdtrRefInf,omitempty"`

	// Identification of the organisation issuing the invoice, when it is different from the creditor or ultimate creditor.
	Invoicer *PartyIdentification43 `xml:"Invcr,omitempty"`

	// Identification of the party to whom an invoice is issued, when it is different from the debtor or ultimate debtor.
	Invoicee *PartyIdentification43 `xml:"Invcee,omitempty"`

	// Additional information, in free text form, to complement the structured remittance information.
	AdditionalRemittanceInformation []*Max140Text `xml:"AddtlRmtInf,omitempty"`
}

func (s *StructuredRemittanceInformation9) AddReferredDocumentInformation() *ReferredDocumentInformation3 {
	newValue := new(ReferredDocumentInformation3)
	s.ReferredDocumentInformation = append(s.ReferredDocumentInformation, newValue)
	return newValue
}

func (s *StructuredRemittanceInformation9) AddReferredDocumentAmount() *RemittanceAmount2 {
	s.ReferredDocumentAmount = new(RemittanceAmount2)
	return s.ReferredDocumentAmount
}

func (s *StructuredRemittanceInformation9) AddCreditorReferenceInformation() *CreditorReferenceInformation2 {
	s.CreditorReferenceInformation = new(CreditorReferenceInformation2)
	return s.CreditorReferenceInformation
}

func (s *StructuredRemittanceInformation9) AddInvoicer() *PartyIdentification43 {
	s.Invoicer = new(PartyIdentification43)
	return s.Invoicer
}

func (s *StructuredRemittanceInformation9) AddInvoicee() *PartyIdentification43 {
	s.Invoicee = new(PartyIdentification43)
	return s.Invoicee
}

func (s *StructuredRemittanceInformation9) AddAdditionalRemittanceInformation(value string) {
	s.AdditionalRemittanceInformation = append(s.AdditionalRemittanceInformation, (*Max140Text)(&value))
}
