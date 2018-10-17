package iso20022

// Information about the confirmation of a transfer in transaction.
type TransferIn1 struct {

	// General information related to the transfer of a financial instrument.
	TransferDetails *Transfer4 `xml:"TrfDtls"`

	// Information related to the financial instrument received.
	FinancialInstrumentDetails *FinancialInstrument3 `xml:"FinInstrmDtls"`

	// Information related to the account into which the financial instrument was received.
	AccountDetails *InvestmentAccount10 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *DeliverInformation2 `xml:"SttlmDtls"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferIn1) AddTransferDetails() *Transfer4 {
	t.TransferDetails = new(Transfer4)
	return t.TransferDetails
}

func (t *TransferIn1) AddFinancialInstrumentDetails() *FinancialInstrument3 {
	t.FinancialInstrumentDetails = new(FinancialInstrument3)
	return t.FinancialInstrumentDetails
}

func (t *TransferIn1) AddAccountDetails() *InvestmentAccount10 {
	t.AccountDetails = new(InvestmentAccount10)
	return t.AccountDetails
}

func (t *TransferIn1) AddSettlementDetails() *DeliverInformation2 {
	t.SettlementDetails = new(DeliverInformation2)
	return t.SettlementDetails
}

func (t *TransferIn1) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}

// Information about a transfer in transaction.
type TransferIn10 struct {

	// Requested date at which the instructing party places the transfer instruction.
	RequestedTransferDate *DateFormat1Choice `xml:"ReqdTrfDt,omitempty"`

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Details of the transfer and cancellation.
	TransferAndReferences []*TransferIn11 `xml:"TrfAndRefs"`

	// Information related to the account into which the financial instrument is to be received.
	AccountDetails *InvestmentAccount40 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *DeliverInformation13 `xml:"SttlmDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferIn10) AddRequestedTransferDate() *DateFormat1Choice {
	t.RequestedTransferDate = new(DateFormat1Choice)
	return t.RequestedTransferDate
}

func (t *TransferIn10) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *TransferIn10) AddTransferAndReferences() *TransferIn11 {
	newValue := new(TransferIn11)
	t.TransferAndReferences = append(t.TransferAndReferences, newValue)
	return newValue
}

func (t *TransferIn10) AddAccountDetails() *InvestmentAccount40 {
	t.AccountDetails = new(InvestmentAccount40)
	return t.AccountDetails
}

func (t *TransferIn10) AddSettlementDetails() *DeliverInformation13 {
	t.SettlementDetails = new(DeliverInformation13)
	return t.SettlementDetails
}

func (t *TransferIn10) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}

// Information about a transfer in transaction.
type TransferIn11 struct {

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails []*Transfer21 `xml:"TrfDtls"`
}

func (t *TransferIn11) SetCancellationReference(value string) {
	t.CancellationReference = (*Max35Text)(&value)
}

func (t *TransferIn11) AddTransferDetails() *Transfer21 {
	newValue := new(Transfer21)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

// Information about the confirmation of a transfer in transaction.
type TransferIn12 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails []*Transfer29 `xml:"TrfDtls"`

	// Information related to the account into which the financial instrument was received.
	AccountDetails *InvestmentAccount40 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *DeliverInformation14 `xml:"SttlmDtls,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferIn12) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *TransferIn12) AddTransferDetails() *Transfer29 {
	newValue := new(Transfer29)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferIn12) AddAccountDetails() *InvestmentAccount40 {
	t.AccountDetails = new(InvestmentAccount40)
	return t.AccountDetails
}

func (t *TransferIn12) AddSettlementDetails() *DeliverInformation14 {
	t.SettlementDetails = new(DeliverInformation14)
	return t.SettlementDetails
}

func (t *TransferIn12) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}

// Information about a transfer in transaction.
type TransferIn13 struct {

	// Requested date at which the instructing party places the transfer instruction.
	RequestedTransferDate *DateFormat1Choice `xml:"ReqdTrfDt,omitempty"`

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Details of the transfer and cancellation.
	TransferAndReferences []*TransferIn11 `xml:"TrfAndRefs"`

	// Information related to the account into which the financial instrument is to be received.
	AccountDetails *InvestmentAccount40 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *DeliverInformation15 `xml:"SttlmDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferIn13) AddRequestedTransferDate() *DateFormat1Choice {
	t.RequestedTransferDate = new(DateFormat1Choice)
	return t.RequestedTransferDate
}

func (t *TransferIn13) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *TransferIn13) AddTransferAndReferences() *TransferIn11 {
	newValue := new(TransferIn11)
	t.TransferAndReferences = append(t.TransferAndReferences, newValue)
	return newValue
}

func (t *TransferIn13) AddAccountDetails() *InvestmentAccount40 {
	t.AccountDetails = new(InvestmentAccount40)
	return t.AccountDetails
}

func (t *TransferIn13) AddSettlementDetails() *DeliverInformation15 {
	t.SettlementDetails = new(DeliverInformation15)
	return t.SettlementDetails
}

func (t *TransferIn13) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}

// Information about the confirmation of a transfer in transaction.
type TransferIn14 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails []*Transfer33 `xml:"TrfDtls"`

	// Information related to the account into which the financial instrument was received.
	AccountDetails *InvestmentAccount56 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *DeliverInformation17 `xml:"SttlmDtls,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferIn14) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *TransferIn14) AddTransferDetails() *Transfer33 {
	newValue := new(Transfer33)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferIn14) AddAccountDetails() *InvestmentAccount56 {
	t.AccountDetails = new(InvestmentAccount56)
	return t.AccountDetails
}

func (t *TransferIn14) AddSettlementDetails() *DeliverInformation17 {
	t.SettlementDetails = new(DeliverInformation17)
	return t.SettlementDetails
}

func (t *TransferIn14) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}

// Information about a transfer in transaction.
type TransferIn15 struct {

	// Requested date at which the instructing party places the transfer instruction.
	RequestedTransferDate *DateFormat1Choice `xml:"ReqdTrfDt,omitempty"`

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Details of the transfer and cancellation.
	TransferAndReferences []*TransferIn16 `xml:"TrfAndRefs"`

	// Information related to the account into which the financial instrument is to be received.
	AccountDetails *InvestmentAccount56 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *DeliverInformation16 `xml:"SttlmDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferIn15) AddRequestedTransferDate() *DateFormat1Choice {
	t.RequestedTransferDate = new(DateFormat1Choice)
	return t.RequestedTransferDate
}

func (t *TransferIn15) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *TransferIn15) AddTransferAndReferences() *TransferIn16 {
	newValue := new(TransferIn16)
	t.TransferAndReferences = append(t.TransferAndReferences, newValue)
	return newValue
}

func (t *TransferIn15) AddAccountDetails() *InvestmentAccount56 {
	t.AccountDetails = new(InvestmentAccount56)
	return t.AccountDetails
}

func (t *TransferIn15) AddSettlementDetails() *DeliverInformation16 {
	t.SettlementDetails = new(DeliverInformation16)
	return t.SettlementDetails
}

func (t *TransferIn15) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}

// Information about a transfer in transaction.
type TransferIn16 struct {

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails []*Transfer32 `xml:"TrfDtls"`
}

func (t *TransferIn16) SetCancellationReference(value string) {
	t.CancellationReference = (*Max35Text)(&value)
}

func (t *TransferIn16) AddTransferDetails() *Transfer32 {
	newValue := new(Transfer32)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

// Information about a transfer in transaction.
type TransferIn2 struct {

	// General information related to the transfer of a financial instrument.
	TransferDetails *Transfer3 `xml:"TrfDtls"`

	// Information related to the financial instrument to be received.
	FinancialInstrumentDetails *FinancialInstrument3 `xml:"FinInstrmDtls"`

	// Information related to the account into which the financial instrument is to be received.
	AccountDetails *InvestmentAccount10 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *DeliverInformation1 `xml:"SttlmDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferIn2) AddTransferDetails() *Transfer3 {
	t.TransferDetails = new(Transfer3)
	return t.TransferDetails
}

func (t *TransferIn2) AddFinancialInstrumentDetails() *FinancialInstrument3 {
	t.FinancialInstrumentDetails = new(FinancialInstrument3)
	return t.FinancialInstrumentDetails
}

func (t *TransferIn2) AddAccountDetails() *InvestmentAccount10 {
	t.AccountDetails = new(InvestmentAccount10)
	return t.AccountDetails
}

func (t *TransferIn2) AddSettlementDetails() *DeliverInformation1 {
	t.SettlementDetails = new(DeliverInformation1)
	return t.SettlementDetails
}

func (t *TransferIn2) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}

// Information about a transfer in transaction.
type TransferIn3 struct {

	// General information related to the transfer of a financial instrument.
	TransferDetails *Transfer6 `xml:"TrfDtls"`

	// Information related to the financial instrument to be received.
	FinancialInstrumentDetails *FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Information related to the account into which the financial instrument is to be received.
	AccountDetails *InvestmentAccount22 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *DeliverInformation3 `xml:"SttlmDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferIn3) AddTransferDetails() *Transfer6 {
	t.TransferDetails = new(Transfer6)
	return t.TransferDetails
}

func (t *TransferIn3) AddFinancialInstrumentDetails() *FinancialInstrument13 {
	t.FinancialInstrumentDetails = new(FinancialInstrument13)
	return t.FinancialInstrumentDetails
}

func (t *TransferIn3) AddAccountDetails() *InvestmentAccount22 {
	t.AccountDetails = new(InvestmentAccount22)
	return t.AccountDetails
}

func (t *TransferIn3) AddSettlementDetails() *DeliverInformation3 {
	t.SettlementDetails = new(DeliverInformation3)
	return t.SettlementDetails
}

func (t *TransferIn3) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}

// Information about the confirmation of a transfer in transaction.
type TransferIn4 struct {

	// General information related to the transfer of a financial instrument.
	TransferDetails *Transfer7 `xml:"TrfDtls"`

	// Information related to the financial instrument received.
	FinancialInstrumentDetails *FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Information related to the account into which the financial instrument was received.
	AccountDetails *InvestmentAccount22 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *DeliverInformation4 `xml:"SttlmDtls"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferIn4) AddTransferDetails() *Transfer7 {
	t.TransferDetails = new(Transfer7)
	return t.TransferDetails
}

func (t *TransferIn4) AddFinancialInstrumentDetails() *FinancialInstrument13 {
	t.FinancialInstrumentDetails = new(FinancialInstrument13)
	return t.FinancialInstrumentDetails
}

func (t *TransferIn4) AddAccountDetails() *InvestmentAccount22 {
	t.AccountDetails = new(InvestmentAccount22)
	return t.AccountDetails
}

func (t *TransferIn4) AddSettlementDetails() *DeliverInformation4 {
	t.SettlementDetails = new(DeliverInformation4)
	return t.SettlementDetails
}

func (t *TransferIn4) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}

// Information about a transfer in transaction.
type TransferIn5 struct {

	// General information related to the transfer of a financial instrument.
	TransferDetails []*Transfer16 `xml:"TrfDtls"`

	// Information related to the account into which the financial instrument is to be received.
	AccountDetails *InvestmentAccount22 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *DeliverInformation8 `xml:"SttlmDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferIn5) AddTransferDetails() *Transfer16 {
	newValue := new(Transfer16)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferIn5) AddAccountDetails() *InvestmentAccount22 {
	t.AccountDetails = new(InvestmentAccount22)
	return t.AccountDetails
}

func (t *TransferIn5) AddSettlementDetails() *DeliverInformation8 {
	t.SettlementDetails = new(DeliverInformation8)
	return t.SettlementDetails
}

func (t *TransferIn5) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}

// Information about the confirmation of a transfer in transaction.
type TransferIn6 struct {

	// General information related to the transfer of a financial instrument.
	TransferDetails []*Transfer17 `xml:"TrfDtls"`

	// Information related to the account into which the financial instrument was received.
	AccountDetails *InvestmentAccount22 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *DeliverInformation7 `xml:"SttlmDtls"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferIn6) AddTransferDetails() *Transfer17 {
	newValue := new(Transfer17)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferIn6) AddAccountDetails() *InvestmentAccount22 {
	t.AccountDetails = new(InvestmentAccount22)
	return t.AccountDetails
}

func (t *TransferIn6) AddSettlementDetails() *DeliverInformation7 {
	t.SettlementDetails = new(DeliverInformation7)
	return t.SettlementDetails
}

func (t *TransferIn6) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}

// Information about a transfer in transaction.
type TransferIn7 struct {

	// General information related to the transfer of a financial instrument.
	TransferDetails []*Transfer22 `xml:"TrfDtls"`

	// Information related to the account into which the financial instrument is to be received.
	AccountDetails *InvestmentAccount22 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *DeliverInformation9 `xml:"SttlmDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferIn7) AddTransferDetails() *Transfer22 {
	newValue := new(Transfer22)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferIn7) AddAccountDetails() *InvestmentAccount22 {
	t.AccountDetails = new(InvestmentAccount22)
	return t.AccountDetails
}

func (t *TransferIn7) AddSettlementDetails() *DeliverInformation9 {
	t.SettlementDetails = new(DeliverInformation9)
	return t.SettlementDetails
}

func (t *TransferIn7) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}

// Information about the confirmation of a transfer in transaction.
type TransferIn8 struct {

	// General information related to the transfer of a financial instrument.
	TransferDetails []*Transfer25 `xml:"TrfDtls"`

	// Information related to the account into which the financial instrument was received.
	AccountDetails *InvestmentAccount22 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *DeliverInformation11 `xml:"SttlmDtls,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferIn8) AddTransferDetails() *Transfer25 {
	newValue := new(Transfer25)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferIn8) AddAccountDetails() *InvestmentAccount22 {
	t.AccountDetails = new(InvestmentAccount22)
	return t.AccountDetails
}

func (t *TransferIn8) AddSettlementDetails() *DeliverInformation11 {
	t.SettlementDetails = new(DeliverInformation11)
	return t.SettlementDetails
}

func (t *TransferIn8) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}

// Information about the confirmation of a transfer in transaction.
type TransferIn9 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails []*Transfer29 `xml:"TrfDtls"`

	// Information related to the account into which the financial instrument was received.
	AccountDetails *InvestmentAccount40 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *DeliverInformation12 `xml:"SttlmDtls,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferIn9) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *TransferIn9) AddTransferDetails() *Transfer29 {
	newValue := new(Transfer29)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferIn9) AddAccountDetails() *InvestmentAccount40 {
	t.AccountDetails = new(InvestmentAccount40)
	return t.AccountDetails
}

func (t *TransferIn9) AddSettlementDetails() *DeliverInformation12 {
	t.SettlementDetails = new(DeliverInformation12)
	return t.SettlementDetails
}

func (t *TransferIn9) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
