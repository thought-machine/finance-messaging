package iso20022

// Reference of a transfer and of a transfer cancellation.
type TransferReference1 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`
}

func (t *TransferReference1) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *TransferReference1) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *TransferReference1) SetClientReference(value string) {
	t.ClientReference = (*Max35Text)(&value)
}

func (t *TransferReference1) SetCancellationReference(value string) {
	t.CancellationReference = (*Max35Text)(&value)
}

// Reference of a transfer and of a transfer confirmation.
type TransferReference10 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef,omitempty"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *AdditionalReference7 `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for a transfer execution, as assigned by a confirming party.
	TransferConfirmationReference *Max35Text `xml:"TrfConfRef,omitempty"`

	// Unambiguous identification of the transfer allocated by the counterparty.
	CounterpartyReference *AdditionalReference7 `xml:"CtrPtyRef,omitempty"`
}

func (t *TransferReference10) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *TransferReference10) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *TransferReference10) AddClientReference() *AdditionalReference7 {
	t.ClientReference = new(AdditionalReference7)
	return t.ClientReference
}

func (t *TransferReference10) SetTransferConfirmationReference(value string) {
	t.TransferConfirmationReference = (*Max35Text)(&value)
}

func (t *TransferReference10) AddCounterpartyReference() *AdditionalReference7 {
	t.CounterpartyReference = new(AdditionalReference7)
	return t.CounterpartyReference
}

// Reference of a transfer and of a transfer confirmation.
type TransferReference2 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef,omitempty"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for a transfer execution, as assigned by a confirming party.
	TransferConfirmationReference *Max35Text `xml:"TrfConfRef,omitempty"`
}

func (t *TransferReference2) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *TransferReference2) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *TransferReference2) SetClientReference(value string) {
	t.ClientReference = (*Max35Text)(&value)
}

func (t *TransferReference2) SetTransferConfirmationReference(value string) {
	t.TransferConfirmationReference = (*Max35Text)(&value)
}

// Reference of a transfer instruction cancellation.
type TransferReference3 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Transfer and cancellation reference.
	TransferReferences []*TransferReference4 `xml:"TrfRefs"`
}

func (t *TransferReference3) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *TransferReference3) AddTransferReferences() *TransferReference4 {
	newValue := new(TransferReference4)
	t.TransferReferences = append(t.TransferReferences, newValue)
	return newValue
}

// Transfer of a security on a securities account servicer after the receipt of a securities settlement instruction, or the transfer of cash on an account servicer after the receipt of a payment instruction. The transfer consists of a debit/credit operation in the books of the account servicer.
type TransferReference4 struct {

	// Identification assigned by the new plan manager to each transfer of asset.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`
}

func (t *TransferReference4) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *TransferReference4) SetCancellationReference(value string) {
	t.CancellationReference = (*Max35Text)(&value)
}

// Reference of a transfer and of a transfer cancellation.
type TransferReference5 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transfer allocated by the counterparty.
	CounterpartyReference *AdditionalReference2 `xml:"CtrPtyRef,omitempty"`

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`
}

func (t *TransferReference5) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *TransferReference5) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *TransferReference5) SetClientReference(value string) {
	t.ClientReference = (*Max35Text)(&value)
}

func (t *TransferReference5) AddCounterpartyReference() *AdditionalReference2 {
	t.CounterpartyReference = new(AdditionalReference2)
	return t.CounterpartyReference
}

func (t *TransferReference5) SetCancellationReference(value string) {
	t.CancellationReference = (*Max35Text)(&value)
}

// Reference of a transfer and of a transfer confirmation.
type TransferReference6 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef,omitempty"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for a transfer execution, as assigned by a confirming party.
	TransferConfirmationReference *Max35Text `xml:"TrfConfRef,omitempty"`

	// Unambiguous identification of the transfer allocated by the counterparty.
	CounterpartyReference *AdditionalReference2 `xml:"CtrPtyRef,omitempty"`
}

func (t *TransferReference6) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *TransferReference6) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *TransferReference6) SetClientReference(value string) {
	t.ClientReference = (*Max35Text)(&value)
}

func (t *TransferReference6) SetTransferConfirmationReference(value string) {
	t.TransferConfirmationReference = (*Max35Text)(&value)
}

func (t *TransferReference6) AddCounterpartyReference() *AdditionalReference2 {
	t.CounterpartyReference = new(AdditionalReference2)
	return t.CounterpartyReference
}

// Reference of a transfer instruction cancellation.
type TransferReference7 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Transfer and cancellation reference.
	TransferReferences []*TransferReference8 `xml:"TrfRefs"`
}

func (t *TransferReference7) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *TransferReference7) AddTransferReferences() *TransferReference8 {
	newValue := new(TransferReference8)
	t.TransferReferences = append(t.TransferReferences, newValue)
	return newValue
}

// Reference of a transfer and of a transfer cancellation.
type TransferReference8 struct {

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Unique and unambiguous identifier for a transfer execution, as assigned by a confirming party.
	TransferConfirmationReference *Max35Text `xml:"TrfConfRef,omitempty"`
}

func (t *TransferReference8) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *TransferReference8) SetCancellationReference(value string) {
	t.CancellationReference = (*Max35Text)(&value)
}

func (t *TransferReference8) SetTransferConfirmationReference(value string) {
	t.TransferConfirmationReference = (*Max35Text)(&value)
}

// Reference of a transfer and of a transfer cancellation.
type TransferReference9 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *AdditionalReference7 `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transfer allocated by the counterparty.
	CounterpartyReference *AdditionalReference7 `xml:"CtrPtyRef,omitempty"`

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`
}

func (t *TransferReference9) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *TransferReference9) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *TransferReference9) AddClientReference() *AdditionalReference7 {
	t.ClientReference = new(AdditionalReference7)
	return t.ClientReference
}

func (t *TransferReference9) AddCounterpartyReference() *AdditionalReference7 {
	t.CounterpartyReference = new(AdditionalReference7)
	return t.CounterpartyReference
}

func (t *TransferReference9) SetCancellationReference(value string) {
	t.CancellationReference = (*Max35Text)(&value)
}
