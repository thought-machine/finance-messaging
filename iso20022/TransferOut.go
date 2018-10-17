package iso20022

// Information about the confirmation of a transfer out transaction.
type TransferOut1 struct {

	// General information related to the transfer of a financial instrument.
	TransferDetails *Transfer2 `xml:"TrfDtls"`

	// Information related to the financial instrument withdrawn.
	FinancialInstrumentDetails *FinancialInstrument3 `xml:"FinInstrmDtls"`

	// Information related to the account from which the financial instrument was withdrawn.
	AccountDetails *InvestmentAccount10 `xml:"AcctDtls"`

	// Information related to the receiving side of the transfer.
	SettlementDetails *ReceiveInformation2 `xml:"SttlmDtls"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferOut1) AddTransferDetails() *Transfer2 {
	t.TransferDetails = new(Transfer2)
	return t.TransferDetails
}

func (t *TransferOut1) AddFinancialInstrumentDetails() *FinancialInstrument3 {
	t.FinancialInstrumentDetails = new(FinancialInstrument3)
	return t.FinancialInstrumentDetails
}

func (t *TransferOut1) AddAccountDetails() *InvestmentAccount10 {
	t.AccountDetails = new(InvestmentAccount10)
	return t.AccountDetails
}

func (t *TransferOut1) AddSettlementDetails() *ReceiveInformation2 {
	t.SettlementDetails = new(ReceiveInformation2)
	return t.SettlementDetails
}

func (t *TransferOut1) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}

// Information about the confirmation of a transfer out transaction.
type TransferOut10 struct {

	// General information related to the transfer of a financial instrument.
	TransferDetails []*Transfer24 `xml:"TrfDtls"`

	// Information related to the account from which the financial instrument was withdrawn.
	AccountDetails *InvestmentAccount22 `xml:"AcctDtls"`

	// Information related to the receiving side of the transfer.
	SettlementDetails *ReceiveInformation11 `xml:"SttlmDtls,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferOut10) AddTransferDetails() *Transfer24 {
	newValue := new(Transfer24)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferOut10) AddAccountDetails() *InvestmentAccount22 {
	t.AccountDetails = new(InvestmentAccount22)
	return t.AccountDetails
}

func (t *TransferOut10) AddSettlementDetails() *ReceiveInformation11 {
	t.SettlementDetails = new(ReceiveInformation11)
	return t.SettlementDetails
}

func (t *TransferOut10) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}

// Information about a transfer out transaction.
type TransferOut11 struct {

	// Requested date at which the instructing party places the transfer instruction.
	RequestedTransferDate *DateFormat1Choice `xml:"ReqdTrfDt,omitempty"`

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Details of the transfer and cancellation.
	TransferAndReferences []*TransferOut13 `xml:"TrfAndRefs"`

	// Information related to the account from which the financial instrument is to be withdrawn.
	AccountDetails *InvestmentAccount40 `xml:"AcctDtls"`

	// Information related to the receiving side of the transfer.
	SettlementDetails *ReceiveInformation13 `xml:"SttlmDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferOut11) AddRequestedTransferDate() *DateFormat1Choice {
	t.RequestedTransferDate = new(DateFormat1Choice)
	return t.RequestedTransferDate
}

func (t *TransferOut11) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *TransferOut11) AddTransferAndReferences() *TransferOut13 {
	newValue := new(TransferOut13)
	t.TransferAndReferences = append(t.TransferAndReferences, newValue)
	return newValue
}

func (t *TransferOut11) AddAccountDetails() *InvestmentAccount40 {
	t.AccountDetails = new(InvestmentAccount40)
	return t.AccountDetails
}

func (t *TransferOut11) AddSettlementDetails() *ReceiveInformation13 {
	t.SettlementDetails = new(ReceiveInformation13)
	return t.SettlementDetails
}

func (t *TransferOut11) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}

// Information about the confirmation of a transfer out transaction.
type TransferOut12 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails []*Transfer28 `xml:"TrfDtls"`

	// Information related to the account from which the financial instrument was withdrawn.
	AccountDetails *InvestmentAccount40 `xml:"AcctDtls"`

	// Information related to the receiving side of the transfer.
	SettlementDetails *ReceiveInformation12 `xml:"SttlmDtls,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferOut12) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *TransferOut12) AddTransferDetails() *Transfer28 {
	newValue := new(Transfer28)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferOut12) AddAccountDetails() *InvestmentAccount40 {
	t.AccountDetails = new(InvestmentAccount40)
	return t.AccountDetails
}

func (t *TransferOut12) AddSettlementDetails() *ReceiveInformation12 {
	t.SettlementDetails = new(ReceiveInformation12)
	return t.SettlementDetails
}

func (t *TransferOut12) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}

// Information about a transfer out transaction.
type TransferOut13 struct {

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails []*Transfer27 `xml:"TrfDtls"`
}

func (t *TransferOut13) SetCancellationReference(value string) {
	t.CancellationReference = (*Max35Text)(&value)
}

func (t *TransferOut13) AddTransferDetails() *Transfer27 {
	newValue := new(Transfer27)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

// Information about the confirmation of a transfer out transaction.
type TransferOut14 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails []*Transfer28 `xml:"TrfDtls"`

	// Information related to the account from which the financial instrument was withdrawn.
	AccountDetails *InvestmentAccount40 `xml:"AcctDtls"`

	// Information related to the receiving side of the transfer.
	SettlementDetails *ReceiveInformation14 `xml:"SttlmDtls,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferOut14) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *TransferOut14) AddTransferDetails() *Transfer28 {
	newValue := new(Transfer28)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferOut14) AddAccountDetails() *InvestmentAccount40 {
	t.AccountDetails = new(InvestmentAccount40)
	return t.AccountDetails
}

func (t *TransferOut14) AddSettlementDetails() *ReceiveInformation14 {
	t.SettlementDetails = new(ReceiveInformation14)
	return t.SettlementDetails
}

func (t *TransferOut14) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}

// Information about a transfer out transaction.
type TransferOut15 struct {

	// Requested date at which the instructing party places the transfer instruction.
	RequestedTransferDate *DateFormat1Choice `xml:"ReqdTrfDt,omitempty"`

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Details of the transfer and cancellation.
	TransferAndReferences []*TransferOut13 `xml:"TrfAndRefs"`

	// Information related to the account from which the financial instrument is to be withdrawn.
	AccountDetails *InvestmentAccount40 `xml:"AcctDtls"`

	// Information related to the receiving side of the transfer.
	SettlementDetails *ReceiveInformation15 `xml:"SttlmDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferOut15) AddRequestedTransferDate() *DateFormat1Choice {
	t.RequestedTransferDate = new(DateFormat1Choice)
	return t.RequestedTransferDate
}

func (t *TransferOut15) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *TransferOut15) AddTransferAndReferences() *TransferOut13 {
	newValue := new(TransferOut13)
	t.TransferAndReferences = append(t.TransferAndReferences, newValue)
	return newValue
}

func (t *TransferOut15) AddAccountDetails() *InvestmentAccount40 {
	t.AccountDetails = new(InvestmentAccount40)
	return t.AccountDetails
}

func (t *TransferOut15) AddSettlementDetails() *ReceiveInformation15 {
	t.SettlementDetails = new(ReceiveInformation15)
	return t.SettlementDetails
}

func (t *TransferOut15) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}

// Information about the confirmation of a transfer out transaction.
type TransferOut16 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails []*Transfer31 `xml:"TrfDtls"`

	// Information related to the account from which the financial instrument was withdrawn.
	AccountDetails *InvestmentAccount54 `xml:"AcctDtls"`

	// Information related to the receiving side of the transfer.
	SettlementDetails *ReceiveInformation17 `xml:"SttlmDtls,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferOut16) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *TransferOut16) AddTransferDetails() *Transfer31 {
	newValue := new(Transfer31)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferOut16) AddAccountDetails() *InvestmentAccount54 {
	t.AccountDetails = new(InvestmentAccount54)
	return t.AccountDetails
}

func (t *TransferOut16) AddSettlementDetails() *ReceiveInformation17 {
	t.SettlementDetails = new(ReceiveInformation17)
	return t.SettlementDetails
}

func (t *TransferOut16) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}

// Information about a transfer out transaction.
type TransferOut17 struct {

	// Requested date at which the instructing party places the transfer instruction.
	RequestedTransferDate *DateFormat1Choice `xml:"ReqdTrfDt,omitempty"`

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Details of the transfer and cancellation.
	TransferAndReferences []*TransferOut18 `xml:"TrfAndRefs"`

	// Information related to the account from which the financial instrument is to be withdrawn.
	AccountDetails *InvestmentAccount54 `xml:"AcctDtls"`

	// Information related to the receiving side of the transfer.
	SettlementDetails *ReceiveInformation16 `xml:"SttlmDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferOut17) AddRequestedTransferDate() *DateFormat1Choice {
	t.RequestedTransferDate = new(DateFormat1Choice)
	return t.RequestedTransferDate
}

func (t *TransferOut17) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *TransferOut17) AddTransferAndReferences() *TransferOut18 {
	newValue := new(TransferOut18)
	t.TransferAndReferences = append(t.TransferAndReferences, newValue)
	return newValue
}

func (t *TransferOut17) AddAccountDetails() *InvestmentAccount54 {
	t.AccountDetails = new(InvestmentAccount54)
	return t.AccountDetails
}

func (t *TransferOut17) AddSettlementDetails() *ReceiveInformation16 {
	t.SettlementDetails = new(ReceiveInformation16)
	return t.SettlementDetails
}

func (t *TransferOut17) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}

// Information about a transfer out transaction.
type TransferOut18 struct {

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails []*Transfer30 `xml:"TrfDtls"`
}

func (t *TransferOut18) SetCancellationReference(value string) {
	t.CancellationReference = (*Max35Text)(&value)
}

func (t *TransferOut18) AddTransferDetails() *Transfer30 {
	newValue := new(Transfer30)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

// Information about a transfer out transaction.
type TransferOut2 struct {

	// General information related to the transfer of a financial instrument.
	TransferDetails *Transfer1 `xml:"TrfDtls"`

	// Information related to the financial instrument to be withdrawn.
	FinancialInstrumentDetails *FinancialInstrument3 `xml:"FinInstrmDtls"`

	// Information related to the account from which the financial instrument is to be withdrawn.
	AccountDetails *InvestmentAccount10 `xml:"AcctDtls"`

	// Information related to the receiving side of the transfer.
	SettlementDetails *ReceiveInformation1 `xml:"SttlmDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferOut2) AddTransferDetails() *Transfer1 {
	t.TransferDetails = new(Transfer1)
	return t.TransferDetails
}

func (t *TransferOut2) AddFinancialInstrumentDetails() *FinancialInstrument3 {
	t.FinancialInstrumentDetails = new(FinancialInstrument3)
	return t.FinancialInstrumentDetails
}

func (t *TransferOut2) AddAccountDetails() *InvestmentAccount10 {
	t.AccountDetails = new(InvestmentAccount10)
	return t.AccountDetails
}

func (t *TransferOut2) AddSettlementDetails() *ReceiveInformation1 {
	t.SettlementDetails = new(ReceiveInformation1)
	return t.SettlementDetails
}

func (t *TransferOut2) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}

// Information about a transfer out transaction.
type TransferOut5 struct {

	// General information related to the transfer of a financial instrument.
	TransferDetails *Transfer9 `xml:"TrfDtls"`

	// Information related to the financial instrument to be withdrawn.
	FinancialInstrumentDetails *FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Information related to the account from which the financial instrument is to be withdrawn.
	AccountDetails *InvestmentAccount22 `xml:"AcctDtls"`

	// Information related to the receiving side of the transfer.
	SettlementDetails *ReceiveInformation3 `xml:"SttlmDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferOut5) AddTransferDetails() *Transfer9 {
	t.TransferDetails = new(Transfer9)
	return t.TransferDetails
}

func (t *TransferOut5) AddFinancialInstrumentDetails() *FinancialInstrument13 {
	t.FinancialInstrumentDetails = new(FinancialInstrument13)
	return t.FinancialInstrumentDetails
}

func (t *TransferOut5) AddAccountDetails() *InvestmentAccount22 {
	t.AccountDetails = new(InvestmentAccount22)
	return t.AccountDetails
}

func (t *TransferOut5) AddSettlementDetails() *ReceiveInformation3 {
	t.SettlementDetails = new(ReceiveInformation3)
	return t.SettlementDetails
}

func (t *TransferOut5) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}

// Information about the confirmation of a transfer out transaction.
type TransferOut6 struct {

	// General information related to the transfer of a financial instrument.
	TransferDetails *Transfer10 `xml:"TrfDtls"`

	// Information related to the financial instrument withdrawn.
	FinancialInstrumentDetails *FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Information related to the account from which the financial instrument was withdrawn.
	AccountDetails *InvestmentAccount22 `xml:"AcctDtls"`

	// Information related to the receiving side of the transfer.
	SettlementDetails *ReceiveInformation4 `xml:"SttlmDtls"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferOut6) AddTransferDetails() *Transfer10 {
	t.TransferDetails = new(Transfer10)
	return t.TransferDetails
}

func (t *TransferOut6) AddFinancialInstrumentDetails() *FinancialInstrument13 {
	t.FinancialInstrumentDetails = new(FinancialInstrument13)
	return t.FinancialInstrumentDetails
}

func (t *TransferOut6) AddAccountDetails() *InvestmentAccount22 {
	t.AccountDetails = new(InvestmentAccount22)
	return t.AccountDetails
}

func (t *TransferOut6) AddSettlementDetails() *ReceiveInformation4 {
	t.SettlementDetails = new(ReceiveInformation4)
	return t.SettlementDetails
}

func (t *TransferOut6) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}

// Information about a transfer out transaction.
type TransferOut7 struct {

	// Requested date at which the instructing party places the transfer instruction.
	RequestedTransferDate *DateFormat1Choice `xml:"ReqdTrfDt,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails []*Transfer12 `xml:"TrfDtls"`

	// Information related to the account from which the financial instrument is to be withdrawn.
	AccountDetails *InvestmentAccount22 `xml:"AcctDtls"`

	// Information related to the receiving side of the transfer.
	SettlementDetails *ReceiveInformation7 `xml:"SttlmDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferOut7) AddRequestedTransferDate() *DateFormat1Choice {
	t.RequestedTransferDate = new(DateFormat1Choice)
	return t.RequestedTransferDate
}

func (t *TransferOut7) AddTransferDetails() *Transfer12 {
	newValue := new(Transfer12)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferOut7) AddAccountDetails() *InvestmentAccount22 {
	t.AccountDetails = new(InvestmentAccount22)
	return t.AccountDetails
}

func (t *TransferOut7) AddSettlementDetails() *ReceiveInformation7 {
	t.SettlementDetails = new(ReceiveInformation7)
	return t.SettlementDetails
}

func (t *TransferOut7) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}

// Information about the confirmation of a transfer out transaction.
type TransferOut8 struct {

	// General information related to the transfer of a financial instrument.
	TransferDetails []*Transfer14 `xml:"TrfDtls"`

	// Information related to the account from which the financial instrument was withdrawn.
	AccountDetails *InvestmentAccount22 `xml:"AcctDtls"`

	// Information related to the receiving side of the transfer.
	SettlementDetails *ReceiveInformation8 `xml:"SttlmDtls"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferOut8) AddTransferDetails() *Transfer14 {
	newValue := new(Transfer14)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferOut8) AddAccountDetails() *InvestmentAccount22 {
	t.AccountDetails = new(InvestmentAccount22)
	return t.AccountDetails
}

func (t *TransferOut8) AddSettlementDetails() *ReceiveInformation8 {
	t.SettlementDetails = new(ReceiveInformation8)
	return t.SettlementDetails
}

func (t *TransferOut8) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}

// Information about a transfer out transaction.
type TransferOut9 struct {

	// Requested date at which the instructing party places the transfer instruction.
	RequestedTransferDate *DateFormat1Choice `xml:"ReqdTrfDt,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails []*Transfer20 `xml:"TrfDtls"`

	// Information related to the account from which the financial instrument is to be withdrawn.
	AccountDetails *InvestmentAccount22 `xml:"AcctDtls"`

	// Information related to the receiving side of the transfer.
	SettlementDetails *ReceiveInformation9 `xml:"SttlmDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferOut9) AddRequestedTransferDate() *DateFormat1Choice {
	t.RequestedTransferDate = new(DateFormat1Choice)
	return t.RequestedTransferDate
}

func (t *TransferOut9) AddTransferDetails() *Transfer20 {
	newValue := new(Transfer20)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferOut9) AddAccountDetails() *InvestmentAccount22 {
	t.AccountDetails = new(InvestmentAccount22)
	return t.AccountDetails
}

func (t *TransferOut9) AddSettlementDetails() *ReceiveInformation9 {
	t.SettlementDetails = new(ReceiveInformation9)
	return t.SettlementDetails
}

func (t *TransferOut9) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
