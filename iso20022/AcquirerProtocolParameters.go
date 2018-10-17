package iso20022

// Acceptor parameters dedicated to the acquirer protocol.
type AcquirerProtocolParameters1 struct {

	// Identification of the acquirer using this protocol.
	AcquirerIdentification []*GenericIdentification32 `xml:"AcqrrId"`

	// Identification of the payment application, user of the acquirer protocol.
	ApplicationIdentification []*Max35Text `xml:"ApplId,omitempty"`

	// Acquirer host configuration.
	Host []*AcquirerHostConfiguration1 `xml:"Hst,omitempty"`

	// Acquirer protocol parameters of transactions performing an online authorisation.
	OnLineTransaction *AcquirerProtocolParameters2 `xml:"OnLineTx,omitempty"`

	// Acquirer protocol parameters of transactions performing an offline authorisation.
	OffLineTransaction *AcquirerProtocolParameters2 `xml:"OffLineTx,omitempty"`

	// Configuration parameters of reconciliation exchanges.
	ReconciliationExchange *ExchangeConfiguration1 `xml:"RcncltnXchg,omitempty"`

	// Indicates the reconciliation period is assigned by the acquirer instead of the acceptor.
	ReconciliationByAcquirer *TrueFalseIndicator `xml:"RcncltnByAcqrr,omitempty"`

	// Indicates the reconciliation total amounts are computed per currency.
	TotalsPerCurrency *TrueFalseIndicator `xml:"TtlsPerCcy,omitempty"`

	// Types of transaction to include in the batch.
	BatchTransferContent []*BatchTransactionType1Code `xml:"BtchTrfCntt,omitempty"`

	// Configuration of a message item.
	MessageItem []*MessageItemCondition1 `xml:"MsgItm,omitempty"`

	// Indicator to require protection of sensitive card data in messages.
	ProtectCardData *TrueFalseIndicator `xml:"PrtctCardData"`
}

func (a *AcquirerProtocolParameters1) AddAcquirerIdentification() *GenericIdentification32 {
	newValue := new(GenericIdentification32)
	a.AcquirerIdentification = append(a.AcquirerIdentification, newValue)
	return newValue
}

func (a *AcquirerProtocolParameters1) AddApplicationIdentification(value string) {
	a.ApplicationIdentification = append(a.ApplicationIdentification, (*Max35Text)(&value))
}

func (a *AcquirerProtocolParameters1) AddHost() *AcquirerHostConfiguration1 {
	newValue := new(AcquirerHostConfiguration1)
	a.Host = append(a.Host, newValue)
	return newValue
}

func (a *AcquirerProtocolParameters1) AddOnLineTransaction() *AcquirerProtocolParameters2 {
	a.OnLineTransaction = new(AcquirerProtocolParameters2)
	return a.OnLineTransaction
}

func (a *AcquirerProtocolParameters1) AddOffLineTransaction() *AcquirerProtocolParameters2 {
	a.OffLineTransaction = new(AcquirerProtocolParameters2)
	return a.OffLineTransaction
}

func (a *AcquirerProtocolParameters1) AddReconciliationExchange() *ExchangeConfiguration1 {
	a.ReconciliationExchange = new(ExchangeConfiguration1)
	return a.ReconciliationExchange
}

func (a *AcquirerProtocolParameters1) SetReconciliationByAcquirer(value string) {
	a.ReconciliationByAcquirer = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters1) SetTotalsPerCurrency(value string) {
	a.TotalsPerCurrency = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters1) AddBatchTransferContent(value string) {
	a.BatchTransferContent = append(a.BatchTransferContent, (*BatchTransactionType1Code)(&value))
}

func (a *AcquirerProtocolParameters1) AddMessageItem() *MessageItemCondition1 {
	newValue := new(MessageItemCondition1)
	a.MessageItem = append(a.MessageItem, newValue)
	return newValue
}

func (a *AcquirerProtocolParameters1) SetProtectCardData(value string) {
	a.ProtectCardData = (*TrueFalseIndicator)(&value)
}

// Acceptor parameters dedicated to the acquirer protocol.
type AcquirerProtocolParameters2 struct {

	// Mode for the financial capture of the transaction by the acquirer.
	FinancialCapture *FinancialCapture1Code `xml:"FinCaptr"`

	// Configuration of the batch transfers.
	BatchTransfer *ExchangeConfiguration1 `xml:"BtchTrf,omitempty"`

	// Configuration parameters of completion exchanges.
	CompletionExchange *ExchangeConfiguration1 `xml:"CmpltnXchg,omitempty"`
}

func (a *AcquirerProtocolParameters2) SetFinancialCapture(value string) {
	a.FinancialCapture = (*FinancialCapture1Code)(&value)
}

func (a *AcquirerProtocolParameters2) AddBatchTransfer() *ExchangeConfiguration1 {
	a.BatchTransfer = new(ExchangeConfiguration1)
	return a.BatchTransfer
}

func (a *AcquirerProtocolParameters2) AddCompletionExchange() *ExchangeConfiguration1 {
	a.CompletionExchange = new(ExchangeConfiguration1)
	return a.CompletionExchange
}

// Acceptor parameters dedicated to the acquirer protocol.
type AcquirerProtocolParameters3 struct {

	// Identification of the acquirer using this protocol.
	AcquirerIdentification []*GenericIdentification32 `xml:"AcqrrId"`

	// Identification of the payment application, user of the acquirer protocol.
	ApplicationIdentification []*Max35Text `xml:"ApplId,omitempty"`

	// Acquirer host configuration.
	Host []*AcquirerHostConfiguration2 `xml:"Hst,omitempty"`

	// Acquirer protocol parameters of transactions performing an online authorisation.
	OnLineTransaction *AcquirerProtocolParameters4 `xml:"OnLineTx,omitempty"`

	// Acquirer protocol parameters of transactions performing an offline authorisation.
	OffLineTransaction *AcquirerProtocolParameters4 `xml:"OffLineTx,omitempty"`

	// Configuration parameters of reconciliation exchanges.
	ReconciliationExchange *ExchangeConfiguration2 `xml:"RcncltnXchg,omitempty"`

	// Indicates the reconciliation period is assigned by the acquirer instead of the acceptor.
	ReconciliationByAcquirer *TrueFalseIndicator `xml:"RcncltnByAcqrr,omitempty"`

	// Indicates the reconciliation total amounts are computed per currency.
	TotalsPerCurrency *TrueFalseIndicator `xml:"TtlsPerCcy,omitempty"`

	// Indicates that totals in reconciliation or batch must be split per group of points of interaction and card product profiles when these informations are present in the transactions.
	SplitTotals *TrueFalseIndicator `xml:"SpltTtls,omitempty"`

	// Indicates that response messages and an AcceptorCompletionAdvice message following an authorisation exchange must contain protected or plain card data.
	CardDataVerification *TrueFalseIndicator `xml:"CardDataVrfctn,omitempty"`

	// Types of transaction to include in the batch.
	BatchTransferContent []*BatchTransactionType1Code `xml:"BtchTrfCntt,omitempty"`

	// Configuration of a message item.
	MessageItem []*MessageItemCondition1 `xml:"MsgItm,omitempty"`

	// Indicator to require protection of sensitive card data in messages.
	ProtectCardData *TrueFalseIndicator `xml:"PrtctCardData"`
}

func (a *AcquirerProtocolParameters3) AddAcquirerIdentification() *GenericIdentification32 {
	newValue := new(GenericIdentification32)
	a.AcquirerIdentification = append(a.AcquirerIdentification, newValue)
	return newValue
}

func (a *AcquirerProtocolParameters3) AddApplicationIdentification(value string) {
	a.ApplicationIdentification = append(a.ApplicationIdentification, (*Max35Text)(&value))
}

func (a *AcquirerProtocolParameters3) AddHost() *AcquirerHostConfiguration2 {
	newValue := new(AcquirerHostConfiguration2)
	a.Host = append(a.Host, newValue)
	return newValue
}

func (a *AcquirerProtocolParameters3) AddOnLineTransaction() *AcquirerProtocolParameters4 {
	a.OnLineTransaction = new(AcquirerProtocolParameters4)
	return a.OnLineTransaction
}

func (a *AcquirerProtocolParameters3) AddOffLineTransaction() *AcquirerProtocolParameters4 {
	a.OffLineTransaction = new(AcquirerProtocolParameters4)
	return a.OffLineTransaction
}

func (a *AcquirerProtocolParameters3) AddReconciliationExchange() *ExchangeConfiguration2 {
	a.ReconciliationExchange = new(ExchangeConfiguration2)
	return a.ReconciliationExchange
}

func (a *AcquirerProtocolParameters3) SetReconciliationByAcquirer(value string) {
	a.ReconciliationByAcquirer = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters3) SetTotalsPerCurrency(value string) {
	a.TotalsPerCurrency = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters3) SetSplitTotals(value string) {
	a.SplitTotals = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters3) SetCardDataVerification(value string) {
	a.CardDataVerification = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters3) AddBatchTransferContent(value string) {
	a.BatchTransferContent = append(a.BatchTransferContent, (*BatchTransactionType1Code)(&value))
}

func (a *AcquirerProtocolParameters3) AddMessageItem() *MessageItemCondition1 {
	newValue := new(MessageItemCondition1)
	a.MessageItem = append(a.MessageItem, newValue)
	return newValue
}

func (a *AcquirerProtocolParameters3) SetProtectCardData(value string) {
	a.ProtectCardData = (*TrueFalseIndicator)(&value)
}

// Acceptor parameters dedicated to the acquirer protocol.
type AcquirerProtocolParameters4 struct {

	// Mode for the financial capture of the transaction by the acquirer.
	FinancialCapture *FinancialCapture1Code `xml:"FinCaptr"`

	// Configuration of the batch transfers.
	BatchTransfer *ExchangeConfiguration2 `xml:"BtchTrf,omitempty"`

	// Configuration parameters of completion exchanges.
	CompletionExchange *ExchangeConfiguration3 `xml:"CmpltnXchg,omitempty"`

	// Configuration of the cancellation exchanges.
	CancellationExchange *CancellationProcess1Code `xml:"CxlXchg,omitempty"`
}

func (a *AcquirerProtocolParameters4) SetFinancialCapture(value string) {
	a.FinancialCapture = (*FinancialCapture1Code)(&value)
}

func (a *AcquirerProtocolParameters4) AddBatchTransfer() *ExchangeConfiguration2 {
	a.BatchTransfer = new(ExchangeConfiguration2)
	return a.BatchTransfer
}

func (a *AcquirerProtocolParameters4) AddCompletionExchange() *ExchangeConfiguration3 {
	a.CompletionExchange = new(ExchangeConfiguration3)
	return a.CompletionExchange
}

func (a *AcquirerProtocolParameters4) SetCancellationExchange(value string) {
	a.CancellationExchange = (*CancellationProcess1Code)(&value)
}

// Acceptor parameters dedicated to the acquirer protocol.
type AcquirerProtocolParameters5 struct {

	// Mode for the financial capture of the transaction by the acquirer.
	FinancialCapture *FinancialCapture1Code `xml:"FinCaptr"`

	// Configuration of the batch transfers.
	BatchTransfer *ExchangeConfiguration4 `xml:"BtchTrf,omitempty"`

	// Configuration parameters of completion exchanges.
	CompletionExchange *ExchangeConfiguration5 `xml:"CmpltnXchg,omitempty"`

	// Configuration of the cancellation exchanges.
	CancellationExchange *CancellationProcess1Code `xml:"CxlXchg,omitempty"`
}

func (a *AcquirerProtocolParameters5) SetFinancialCapture(value string) {
	a.FinancialCapture = (*FinancialCapture1Code)(&value)
}

func (a *AcquirerProtocolParameters5) AddBatchTransfer() *ExchangeConfiguration4 {
	a.BatchTransfer = new(ExchangeConfiguration4)
	return a.BatchTransfer
}

func (a *AcquirerProtocolParameters5) AddCompletionExchange() *ExchangeConfiguration5 {
	a.CompletionExchange = new(ExchangeConfiguration5)
	return a.CompletionExchange
}

func (a *AcquirerProtocolParameters5) SetCancellationExchange(value string) {
	a.CancellationExchange = (*CancellationProcess1Code)(&value)
}

// Acceptor parameters dedicated to the acquirer protocol.
type AcquirerProtocolParameters6 struct {

	// Identification of the acquirer using this protocol.
	AcquirerIdentification []*GenericIdentification32 `xml:"AcqrrId"`

	// Identification of the payment application, user of the acquirer protocol.
	ApplicationIdentification []*Max35Text `xml:"ApplId,omitempty"`

	// Acquirer host configuration.
	Host []*AcquirerHostConfiguration2 `xml:"Hst,omitempty"`

	// Acquirer protocol parameters of transactions performing an online authorisation.
	OnLineTransaction *AcquirerProtocolParameters5 `xml:"OnLineTx,omitempty"`

	// Acquirer protocol parameters of transactions performing an offline authorisation.
	OffLineTransaction *AcquirerProtocolParameters5 `xml:"OffLineTx,omitempty"`

	// Configuration parameters of reconciliation exchanges.
	ReconciliationExchange *ExchangeConfiguration4 `xml:"RcncltnXchg,omitempty"`

	// Indicates the reconciliation period is assigned by the acquirer instead of the acceptor.
	ReconciliationByAcquirer *TrueFalseIndicator `xml:"RcncltnByAcqrr,omitempty"`

	// Indicates the reconciliation total amounts are computed per currency.
	TotalsPerCurrency *TrueFalseIndicator `xml:"TtlsPerCcy,omitempty"`

	// Indicates that totals in reconciliation or batch must be split per group of points of interaction and card product profiles when these informations are present in the transactions.
	SplitTotals *TrueFalseIndicator `xml:"SpltTtls,omitempty"`

	// Indicates that response messages and an AcceptorCompletionAdvice message following an authorisation exchange must contain protected or plain card data.
	CardDataVerification *TrueFalseIndicator `xml:"CardDataVrfctn,omitempty"`

	// Send a cancellation advice for offline transactions not yet captured.
	NotifyOffLineCancellation *TrueFalseIndicator `xml:"NtfyOffLineCxl,omitempty"`

	// Types of transaction to include in the batch.
	BatchTransferContent []*BatchTransactionType1Code `xml:"BtchTrfCntt,omitempty"`

	// Configuration of a message item.
	MessageItem []*MessageItemCondition1 `xml:"MsgItm,omitempty"`

	// Indicator to require protection of sensitive card data in messages.
	ProtectCardData *TrueFalseIndicator `xml:"PrtctCardData"`
}

func (a *AcquirerProtocolParameters6) AddAcquirerIdentification() *GenericIdentification32 {
	newValue := new(GenericIdentification32)
	a.AcquirerIdentification = append(a.AcquirerIdentification, newValue)
	return newValue
}

func (a *AcquirerProtocolParameters6) AddApplicationIdentification(value string) {
	a.ApplicationIdentification = append(a.ApplicationIdentification, (*Max35Text)(&value))
}

func (a *AcquirerProtocolParameters6) AddHost() *AcquirerHostConfiguration2 {
	newValue := new(AcquirerHostConfiguration2)
	a.Host = append(a.Host, newValue)
	return newValue
}

func (a *AcquirerProtocolParameters6) AddOnLineTransaction() *AcquirerProtocolParameters5 {
	a.OnLineTransaction = new(AcquirerProtocolParameters5)
	return a.OnLineTransaction
}

func (a *AcquirerProtocolParameters6) AddOffLineTransaction() *AcquirerProtocolParameters5 {
	a.OffLineTransaction = new(AcquirerProtocolParameters5)
	return a.OffLineTransaction
}

func (a *AcquirerProtocolParameters6) AddReconciliationExchange() *ExchangeConfiguration4 {
	a.ReconciliationExchange = new(ExchangeConfiguration4)
	return a.ReconciliationExchange
}

func (a *AcquirerProtocolParameters6) SetReconciliationByAcquirer(value string) {
	a.ReconciliationByAcquirer = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters6) SetTotalsPerCurrency(value string) {
	a.TotalsPerCurrency = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters6) SetSplitTotals(value string) {
	a.SplitTotals = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters6) SetCardDataVerification(value string) {
	a.CardDataVerification = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters6) SetNotifyOffLineCancellation(value string) {
	a.NotifyOffLineCancellation = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters6) AddBatchTransferContent(value string) {
	a.BatchTransferContent = append(a.BatchTransferContent, (*BatchTransactionType1Code)(&value))
}

func (a *AcquirerProtocolParameters6) AddMessageItem() *MessageItemCondition1 {
	newValue := new(MessageItemCondition1)
	a.MessageItem = append(a.MessageItem, newValue)
	return newValue
}

func (a *AcquirerProtocolParameters6) SetProtectCardData(value string) {
	a.ProtectCardData = (*TrueFalseIndicator)(&value)
}

// Acceptor parameters dedicated to the acquirer protocol.
type AcquirerProtocolParameters7 struct {

	// Identification of the acquirer using this protocol.
	AcquirerIdentification []*GenericIdentification53 `xml:"AcqrrId"`

	// Version of the acquirer protocol parameters.
	Version *Max256Text `xml:"Vrsn"`

	// Identification of the payment application, user of the acquirer protocol.
	ApplicationIdentification []*Max35Text `xml:"ApplId,omitempty"`

	// Acquirer host configuration.
	Host []*AcquirerHostConfiguration3 `xml:"Hst,omitempty"`

	// Acquirer protocol parameters of transactions performing an online authorisation.
	OnLineTransaction *AcquirerProtocolParameters8 `xml:"OnLineTx,omitempty"`

	// Acquirer protocol parameters of transactions performing an offline authorisation.
	OffLineTransaction *AcquirerProtocolParameters8 `xml:"OffLineTx,omitempty"`

	// Configuration parameters of reconciliation exchanges.
	ReconciliationExchange *ExchangeConfiguration6 `xml:"RcncltnXchg,omitempty"`

	// Indicates the reconciliation period is assigned by the acquirer instead of the acceptor.
	ReconciliationByAcquirer *TrueFalseIndicator `xml:"RcncltnByAcqrr,omitempty"`

	// Indicates the reconciliation total amounts are computed per currency.
	TotalsPerCurrency *TrueFalseIndicator `xml:"TtlsPerCcy,omitempty"`

	// Indicates that totals in reconciliation or batch must be split per group of points of interaction and card product profiles when these information are present in the transactions.
	SplitTotals *TrueFalseIndicator `xml:"SpltTtls,omitempty"`

	// True if the POI must send card data (protected or plain card data) in the AcceptorCompletionAdvice message following an authorisation exchange.
	CardDataVerification *TrueFalseIndicator `xml:"CardDataVrfctn,omitempty"`

	// Send a cancellation advice for offline transactions not yet captured.
	NotifyOffLineCancellation *TrueFalseIndicator `xml:"NtfyOffLineCxl,omitempty"`

	// Types of transaction to include in the batch.
	BatchTransferContent []*BatchTransactionType1Code `xml:"BtchTrfCntt,omitempty"`

	// BatchTransfer are exchanged per file transfer protocol rather than per message.
	FileTransferBatch *TrueFalseIndicator `xml:"FileTrfBtch,omitempty"`

	// BatchTransfer are authenticated by digital signature rather than a MAC (Message Authentication Code).
	BatchDigitalSignature *TrueFalseIndicator `xml:"BtchDgtlSgntr,omitempty"`

	// Configuration of a message item.
	MessageItem []*MessageItemCondition1 `xml:"MsgItm,omitempty"`

	// Indicator to require protection of sensitive card data in messages.
	ProtectCardData *TrueFalseIndicator `xml:"PrtctCardData"`
}

func (a *AcquirerProtocolParameters7) AddAcquirerIdentification() *GenericIdentification53 {
	newValue := new(GenericIdentification53)
	a.AcquirerIdentification = append(a.AcquirerIdentification, newValue)
	return newValue
}

func (a *AcquirerProtocolParameters7) SetVersion(value string) {
	a.Version = (*Max256Text)(&value)
}

func (a *AcquirerProtocolParameters7) AddApplicationIdentification(value string) {
	a.ApplicationIdentification = append(a.ApplicationIdentification, (*Max35Text)(&value))
}

func (a *AcquirerProtocolParameters7) AddHost() *AcquirerHostConfiguration3 {
	newValue := new(AcquirerHostConfiguration3)
	a.Host = append(a.Host, newValue)
	return newValue
}

func (a *AcquirerProtocolParameters7) AddOnLineTransaction() *AcquirerProtocolParameters8 {
	a.OnLineTransaction = new(AcquirerProtocolParameters8)
	return a.OnLineTransaction
}

func (a *AcquirerProtocolParameters7) AddOffLineTransaction() *AcquirerProtocolParameters8 {
	a.OffLineTransaction = new(AcquirerProtocolParameters8)
	return a.OffLineTransaction
}

func (a *AcquirerProtocolParameters7) AddReconciliationExchange() *ExchangeConfiguration6 {
	a.ReconciliationExchange = new(ExchangeConfiguration6)
	return a.ReconciliationExchange
}

func (a *AcquirerProtocolParameters7) SetReconciliationByAcquirer(value string) {
	a.ReconciliationByAcquirer = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters7) SetTotalsPerCurrency(value string) {
	a.TotalsPerCurrency = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters7) SetSplitTotals(value string) {
	a.SplitTotals = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters7) SetCardDataVerification(value string) {
	a.CardDataVerification = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters7) SetNotifyOffLineCancellation(value string) {
	a.NotifyOffLineCancellation = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters7) AddBatchTransferContent(value string) {
	a.BatchTransferContent = append(a.BatchTransferContent, (*BatchTransactionType1Code)(&value))
}

func (a *AcquirerProtocolParameters7) SetFileTransferBatch(value string) {
	a.FileTransferBatch = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters7) SetBatchDigitalSignature(value string) {
	a.BatchDigitalSignature = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters7) AddMessageItem() *MessageItemCondition1 {
	newValue := new(MessageItemCondition1)
	a.MessageItem = append(a.MessageItem, newValue)
	return newValue
}

func (a *AcquirerProtocolParameters7) SetProtectCardData(value string) {
	a.ProtectCardData = (*TrueFalseIndicator)(&value)
}

// Acceptor parameters dedicated to the acquirer protocol.
type AcquirerProtocolParameters8 struct {

	// Mode for the financial capture of the transaction by the acquirer.
	FinancialCapture *FinancialCapture1Code `xml:"FinCaptr"`

	// Configuration of the batch transfers.
	BatchTransfer *ExchangeConfiguration6 `xml:"BtchTrf,omitempty"`

	// Configuration parameters of completion exchanges.
	CompletionExchange *ExchangeConfiguration7 `xml:"CmpltnXchg,omitempty"`

	// Configuration of the cancellation exchanges.
	CancellationExchange *CancellationProcess1Code `xml:"CxlXchg,omitempty"`
}

func (a *AcquirerProtocolParameters8) SetFinancialCapture(value string) {
	a.FinancialCapture = (*FinancialCapture1Code)(&value)
}

func (a *AcquirerProtocolParameters8) AddBatchTransfer() *ExchangeConfiguration6 {
	a.BatchTransfer = new(ExchangeConfiguration6)
	return a.BatchTransfer
}

func (a *AcquirerProtocolParameters8) AddCompletionExchange() *ExchangeConfiguration7 {
	a.CompletionExchange = new(ExchangeConfiguration7)
	return a.CompletionExchange
}

func (a *AcquirerProtocolParameters8) SetCancellationExchange(value string) {
	a.CancellationExchange = (*CancellationProcess1Code)(&value)
}

// Acceptor parameters dedicated to the acquirer protocol.
type AcquirerProtocolParameters9 struct {

	// Type of action for the configuration parameters.
	ActionType *TerminalManagementAction3Code `xml:"ActnTp"`

	// Identification of the acquirer using this protocol.
	AcquirerIdentification []*GenericIdentification53 `xml:"AcqrrId"`

	// Version of the acquirer protocol parameters.
	Version *Max256Text `xml:"Vrsn"`

	// Identification of the payment application, user of the acquirer protocol.
	ApplicationIdentification []*Max35Text `xml:"ApplId,omitempty"`

	// Acquirer host configuration.
	Host []*AcquirerHostConfiguration3 `xml:"Hst,omitempty"`

	// Acquirer protocol parameters of transactions performing an online authorisation.
	OnLineTransaction *AcquirerProtocolParameters8 `xml:"OnLineTx,omitempty"`

	// Acquirer protocol parameters of transactions performing an offline authorisation.
	OffLineTransaction *AcquirerProtocolParameters8 `xml:"OffLineTx,omitempty"`

	// Configuration parameters of reconciliation exchanges.
	ReconciliationExchange *ExchangeConfiguration6 `xml:"RcncltnXchg,omitempty"`

	// Indicates the reconciliation period is assigned by the acquirer instead of the acceptor.
	ReconciliationByAcquirer *TrueFalseIndicator `xml:"RcncltnByAcqrr,omitempty"`

	// Indicates the reconciliation total amounts are computed per currency.
	TotalsPerCurrency *TrueFalseIndicator `xml:"TtlsPerCcy,omitempty"`

	// Indicates that totals in reconciliation or batch must be split per group of points of interaction and card product profiles when these information are present in the transactions.
	SplitTotals *TrueFalseIndicator `xml:"SpltTtls,omitempty"`

	// After an error in a totals of the Reconciliation, the POI sends transactions in error in the BatchTransfer messages.
	ReconciliationError *TrueFalseIndicator `xml:"RcncltnErr,omitempty"`

	// True if the POI must send card data (protected or plain card data) in the AcceptorCompletionAdvice message following an authorisation exchange.
	CardDataVerification *TrueFalseIndicator `xml:"CardDataVrfctn,omitempty"`

	// Send a cancellation advice for offline transactions not yet captured.
	NotifyOffLineCancellation *TrueFalseIndicator `xml:"NtfyOffLineCxl,omitempty"`

	// Types of transaction to include in the batch.
	BatchTransferContent []*BatchTransactionType1Code `xml:"BtchTrfCntt,omitempty"`

	// BatchTransfer are exchanged per file transfer protocol rather than per message.
	FileTransferBatch *TrueFalseIndicator `xml:"FileTrfBtch,omitempty"`

	// BatchTransfer are authenticated by digital signature rather than a MAC (Message Authentication Code).
	BatchDigitalSignature *TrueFalseIndicator `xml:"BtchDgtlSgntr,omitempty"`

	// Configuration of a message item.
	MessageItem []*MessageItemCondition1 `xml:"MsgItm,omitempty"`

	// Indicator to require protection of sensitive card data in messages.
	ProtectCardData *TrueFalseIndicator `xml:"PrtctCardData"`

	// A security trailer is mandatory in the messages.
	MandatorySecurityTrailer *TrueFalseIndicator `xml:"MndtrySctyTrlr,omitempty"`
}

func (a *AcquirerProtocolParameters9) SetActionType(value string) {
	a.ActionType = (*TerminalManagementAction3Code)(&value)
}

func (a *AcquirerProtocolParameters9) AddAcquirerIdentification() *GenericIdentification53 {
	newValue := new(GenericIdentification53)
	a.AcquirerIdentification = append(a.AcquirerIdentification, newValue)
	return newValue
}

func (a *AcquirerProtocolParameters9) SetVersion(value string) {
	a.Version = (*Max256Text)(&value)
}

func (a *AcquirerProtocolParameters9) AddApplicationIdentification(value string) {
	a.ApplicationIdentification = append(a.ApplicationIdentification, (*Max35Text)(&value))
}

func (a *AcquirerProtocolParameters9) AddHost() *AcquirerHostConfiguration3 {
	newValue := new(AcquirerHostConfiguration3)
	a.Host = append(a.Host, newValue)
	return newValue
}

func (a *AcquirerProtocolParameters9) AddOnLineTransaction() *AcquirerProtocolParameters8 {
	a.OnLineTransaction = new(AcquirerProtocolParameters8)
	return a.OnLineTransaction
}

func (a *AcquirerProtocolParameters9) AddOffLineTransaction() *AcquirerProtocolParameters8 {
	a.OffLineTransaction = new(AcquirerProtocolParameters8)
	return a.OffLineTransaction
}

func (a *AcquirerProtocolParameters9) AddReconciliationExchange() *ExchangeConfiguration6 {
	a.ReconciliationExchange = new(ExchangeConfiguration6)
	return a.ReconciliationExchange
}

func (a *AcquirerProtocolParameters9) SetReconciliationByAcquirer(value string) {
	a.ReconciliationByAcquirer = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters9) SetTotalsPerCurrency(value string) {
	a.TotalsPerCurrency = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters9) SetSplitTotals(value string) {
	a.SplitTotals = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters9) SetReconciliationError(value string) {
	a.ReconciliationError = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters9) SetCardDataVerification(value string) {
	a.CardDataVerification = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters9) SetNotifyOffLineCancellation(value string) {
	a.NotifyOffLineCancellation = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters9) AddBatchTransferContent(value string) {
	a.BatchTransferContent = append(a.BatchTransferContent, (*BatchTransactionType1Code)(&value))
}

func (a *AcquirerProtocolParameters9) SetFileTransferBatch(value string) {
	a.FileTransferBatch = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters9) SetBatchDigitalSignature(value string) {
	a.BatchDigitalSignature = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters9) AddMessageItem() *MessageItemCondition1 {
	newValue := new(MessageItemCondition1)
	a.MessageItem = append(a.MessageItem, newValue)
	return newValue
}

func (a *AcquirerProtocolParameters9) SetProtectCardData(value string) {
	a.ProtectCardData = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters9) SetMandatorySecurityTrailer(value string) {
	a.MandatorySecurityTrailer = (*TrueFalseIndicator)(&value)
}
