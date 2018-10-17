package iso20022

// Structured information supplied to fully identify the documents referred to in the remittance information.
type ReferredDocumentInformation1 struct {

	// Provides the type of the referred document.
	ReferredDocumentType *ReferredDocumentType1 `xml:"RfrdDocTp,omitempty"`

	// Unique and unambiguous identification number of the referred document.
	ReferredDocumentNumber *Max35Text `xml:"RfrdDocNb,omitempty"`
}

func (r *ReferredDocumentInformation1) AddReferredDocumentType() *ReferredDocumentType1 {
	r.ReferredDocumentType = new(ReferredDocumentType1)
	return r.ReferredDocumentType
}

func (r *ReferredDocumentInformation1) SetReferredDocumentNumber(value string) {
	r.ReferredDocumentNumber = (*Max35Text)(&value)
}

// Structured information related to the invoice to be financed.
type ReferredDocumentInformation2 struct {

	// Specifies the type of the document, for example commercial invoice.
	Type *ReferredDocumentType1 `xml:"Tp,omitempty"`

	// Unique and unambiguous identification number of the referred document.
	DocumentNumber *Max35Text `xml:"DocNb,omitempty"`

	// Date associated with the referred document, eg, date of issue.
	RelatedDate *ISODate `xml:"RltdDt,omitempty"`

	// Amount of money and currency of a document referred to invoice to be financed.
	DocumentAmount *ActiveCurrencyAndAmount `xml:"DocAmt,omitempty"`
}

func (r *ReferredDocumentInformation2) AddType() *ReferredDocumentType1 {
	r.Type = new(ReferredDocumentType1)
	return r.Type
}

func (r *ReferredDocumentInformation2) SetDocumentNumber(value string) {
	r.DocumentNumber = (*Max35Text)(&value)
}

func (r *ReferredDocumentInformation2) SetRelatedDate(value string) {
	r.RelatedDate = (*ISODate)(&value)
}

func (r *ReferredDocumentInformation2) SetDocumentAmount(value, currency string) {
	r.DocumentAmount = NewActiveCurrencyAndAmount(value, currency)
}

// Set of elements used to identify the documents referred to in the remittance information.
type ReferredDocumentInformation3 struct {

	// Specifies the type of referred document.
	Type *ReferredDocumentType2 `xml:"Tp,omitempty"`

	// Unique and unambiguous identification of the referred document.
	Number *Max35Text `xml:"Nb,omitempty"`

	// Date associated with the referred document.
	RelatedDate *ISODate `xml:"RltdDt,omitempty"`
}

func (r *ReferredDocumentInformation3) AddType() *ReferredDocumentType2 {
	r.Type = new(ReferredDocumentType2)
	return r.Type
}

func (r *ReferredDocumentInformation3) SetNumber(value string) {
	r.Number = (*Max35Text)(&value)
}

func (r *ReferredDocumentInformation3) SetRelatedDate(value string) {
	r.RelatedDate = (*ISODate)(&value)
}

// Identifies the documents referred to in the remittance information.
type ReferredDocumentInformation4 struct {

	// Specifies the type of referred document.
	Type *ReferredDocumentType2 `xml:"Tp,omitempty"`

	// Unique and unambiguous identification of the referred document.
	Number *Max35Text `xml:"Nb,omitempty"`

	// Date associated with the referred document.
	RelatedDate *ISODate `xml:"RltdDt,omitempty"`

	// Set of elements used to provide the content of the referred document line.
	LineDetails []*DocumentLineInformation1 `xml:"LineDtls,omitempty"`
}

func (r *ReferredDocumentInformation4) AddType() *ReferredDocumentType2 {
	r.Type = new(ReferredDocumentType2)
	return r.Type
}

func (r *ReferredDocumentInformation4) SetNumber(value string) {
	r.Number = (*Max35Text)(&value)
}

func (r *ReferredDocumentInformation4) SetRelatedDate(value string) {
	r.RelatedDate = (*ISODate)(&value)
}

func (r *ReferredDocumentInformation4) AddLineDetails() *DocumentLineInformation1 {
	newValue := new(DocumentLineInformation1)
	r.LineDetails = append(r.LineDetails, newValue)
	return newValue
}

// Set of elements used to identify the documents referred to in the remittance information.
type ReferredDocumentInformation6 struct {

	// Specifies the type of referred document.
	Type *ReferredDocumentType4 `xml:"Tp,omitempty"`

	// Unique and unambiguous identification of the referred document.
	Number *Max35Text `xml:"Nb,omitempty"`

	// Date associated with the referred document.
	RelatedDate *ISODate `xml:"RltdDt,omitempty"`
}

func (r *ReferredDocumentInformation6) AddType() *ReferredDocumentType4 {
	r.Type = new(ReferredDocumentType4)
	return r.Type
}

func (r *ReferredDocumentInformation6) SetNumber(value string) {
	r.Number = (*Max35Text)(&value)
}

func (r *ReferredDocumentInformation6) SetRelatedDate(value string) {
	r.RelatedDate = (*ISODate)(&value)
}

// Set of elements used to identify the documents referred to in the remittance information.
type ReferredDocumentInformation7 struct {

	// Specifies the type of referred document.
	Type *ReferredDocumentType4 `xml:"Tp,omitempty"`

	// Unique and unambiguous identification of the referred document.
	Number *Max35Text `xml:"Nb,omitempty"`

	// Date associated with the referred document.
	RelatedDate *ISODate `xml:"RltdDt,omitempty"`

	// Set of elements used to provide the content of the referred document line.
	LineDetails []*DocumentLineInformation1 `xml:"LineDtls,omitempty"`
}

func (r *ReferredDocumentInformation7) AddType() *ReferredDocumentType4 {
	r.Type = new(ReferredDocumentType4)
	return r.Type
}

func (r *ReferredDocumentInformation7) SetNumber(value string) {
	r.Number = (*Max35Text)(&value)
}

func (r *ReferredDocumentInformation7) SetRelatedDate(value string) {
	r.RelatedDate = (*ISODate)(&value)
}

func (r *ReferredDocumentInformation7) AddLineDetails() *DocumentLineInformation1 {
	newValue := new(DocumentLineInformation1)
	r.LineDetails = append(r.LineDetails, newValue)
	return newValue
}
