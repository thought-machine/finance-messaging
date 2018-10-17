package iso20022

// Unique and unambiguous identifier of the group of transactions as assigned by the original instructing party.
type OriginalGroupInformation1 struct {

	// Point to point reference assigned by the original instructing party to unambiguously identify the original group of individual transactions.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Name assigned by the sending party to unambiguously identify the file transmitted on the network.
	NetworkFileName *Max35Text `xml:"NtwkFileNm"`

	// Specifies the original message name identifier to which the message refers, eg, pacs.003.001.01 or MT103.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Party that sent the file for which the status has been generated.
	//
	// Usage: this field will only be used when the content of the message could not be decoded at the receiving side.
	FileOriginator *Max35Text `xml:"FileOrgtr,omitempty"`

	// Number of individual transactions contained in the original message.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original message, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Specifies the status of a group of transactions.
	GroupStatus *TransactionGroupStatus1Code `xml:"GrpSts,omitempty"`

	// Detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation1 `xml:"StsRsnInf,omitempty"`

	// Detailed information on the number of transactions for each identical individual transaction status.
	NumberOfTransactionsPerStatus []*NumberOfTransactionsPerStatus1 `xml:"NbOfTxsPerSts,omitempty"`
}

func (o *OriginalGroupInformation1) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation1) SetNetworkFileName(value string) {
	o.NetworkFileName = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation1) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation1) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalGroupInformation1) SetFileOriginator(value string) {
	o.FileOriginator = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation1) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalGroupInformation1) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalGroupInformation1) SetGroupStatus(value string) {
	o.GroupStatus = (*TransactionGroupStatus1Code)(&value)
}

func (o *OriginalGroupInformation1) AddStatusReasonInformation() *StatusReasonInformation1 {
	newValue := new(StatusReasonInformation1)
	o.StatusReasonInformation = append(o.StatusReasonInformation, newValue)
	return newValue
}

func (o *OriginalGroupInformation1) AddNumberOfTransactionsPerStatus() *NumberOfTransactionsPerStatus1 {
	newValue := new(NumberOfTransactionsPerStatus1)
	o.NumberOfTransactionsPerStatus = append(o.NumberOfTransactionsPerStatus, newValue)
	return newValue
}

// Unique and unambiguous identifier of the group of transactions as assigned by the original instructing party.
type OriginalGroupInformation2 struct {

	// Point to point reference assigned by the original instructing party to unambiguously identify the original group of individual transactions.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Specifies the original message name identifier to which the message refers, eg, pacs.003.001.01 or MT103.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Detailed information on the return reason.
	ReturnReasonInformation []*ReturnReasonInformation1 `xml:"RtrRsnInf,omitempty"`
}

func (o *OriginalGroupInformation2) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation2) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation2) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalGroupInformation2) AddReturnReasonInformation() *ReturnReasonInformation1 {
	newValue := new(ReturnReasonInformation1)
	o.ReturnReasonInformation = append(o.ReturnReasonInformation, newValue)
	return newValue
}

// Set of elements used to provide information on the original group, to which the message refers.
type OriginalGroupInformation20 struct {

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Specifies the original message name identifier to which the message refers.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Number of individual transactions contained in the original message.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original message, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Specifies the status of a group of transactions.
	GroupStatus *TransactionGroupStatus3Code `xml:"GrpSts,omitempty"`

	// Set of elements used to provide detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation8 `xml:"StsRsnInf,omitempty"`

	// Detailed information on the number of transactions for each identical transaction status.
	NumberOfTransactionsPerStatus []*NumberOfTransactionsPerStatus3 `xml:"NbOfTxsPerSts,omitempty"`
}

func (o *OriginalGroupInformation20) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation20) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation20) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalGroupInformation20) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalGroupInformation20) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalGroupInformation20) SetGroupStatus(value string) {
	o.GroupStatus = (*TransactionGroupStatus3Code)(&value)
}

func (o *OriginalGroupInformation20) AddStatusReasonInformation() *StatusReasonInformation8 {
	newValue := new(StatusReasonInformation8)
	o.StatusReasonInformation = append(o.StatusReasonInformation, newValue)
	return newValue
}

func (o *OriginalGroupInformation20) AddNumberOfTransactionsPerStatus() *NumberOfTransactionsPerStatus3 {
	newValue := new(NumberOfTransactionsPerStatus3)
	o.NumberOfTransactionsPerStatus = append(o.NumberOfTransactionsPerStatus, newValue)
	return newValue
}

// Set of elements used to provide information on the original group, to which the message refers.
type OriginalGroupInformation21 struct {

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Specifies the original message name identifier to which the message refers.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Set of elements used to provide detailed information on the return reason.
	ReturnReasonInformation []*ReturnReasonInformation9 `xml:"RtrRsnInf,omitempty"`
}

func (o *OriginalGroupInformation21) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation21) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation21) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalGroupInformation21) AddReturnReasonInformation() *ReturnReasonInformation9 {
	newValue := new(ReturnReasonInformation9)
	o.ReturnReasonInformation = append(o.ReturnReasonInformation, newValue)
	return newValue
}

// Set of elements used to provide information on the original group, to which the message refers.
type OriginalGroupInformation22 struct {

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Specifies the original message name identifier to which the message refers.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Set of elements used to provide detailed information on the reversal reason.
	ReversalReasonInformation []*ReversalReasonInformation6 `xml:"RvslRsnInf,omitempty"`
}

func (o *OriginalGroupInformation22) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation22) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation22) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalGroupInformation22) AddReversalReasonInformation() *ReversalReasonInformation6 {
	newValue := new(ReversalReasonInformation6)
	o.ReversalReasonInformation = append(o.ReversalReasonInformation, newValue)
	return newValue
}

// Set of elements used to provide information on the original group, to which the message refers.
type OriginalGroupInformation23 struct {

	// Unique identification, as assigned by the assigner, to unambiguously identify the group cancellation request.
	//
	// Usage: The group cancellation request identification can be used for reconciliation or to link tasks related to the cancellation request.
	GroupCancellationIdentification *Max35Text `xml:"GrpCxlId,omitempty"`

	// Set of elements to uniquely and unambiguously identify an exception or an investigation workflow.
	Case *Case2 `xml:"Case,omitempty"`

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Specifies the original message name identifier to which the message refers.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Number of individual transactions contained in the original message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs,omitempty"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Indicates whether the cancellation request applies to a whole group of transactions or to individual transactions within an original group.
	GroupCancellation *GroupCancellationIndicator `xml:"GrpCxl,omitempty"`

	// Set of elements used to provide detailed information on the cancellation reason.
	CancellationReasonInformation []*CancellationReasonInformation3 `xml:"CxlRsnInf,omitempty"`
}

func (o *OriginalGroupInformation23) SetGroupCancellationIdentification(value string) {
	o.GroupCancellationIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation23) AddCase() *Case2 {
	o.Case = new(Case2)
	return o.Case
}

func (o *OriginalGroupInformation23) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation23) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation23) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalGroupInformation23) SetNumberOfTransactions(value string) {
	o.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalGroupInformation23) SetControlSum(value string) {
	o.ControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalGroupInformation23) SetGroupCancellation(value string) {
	o.GroupCancellation = (*GroupCancellationIndicator)(&value)
}

func (o *OriginalGroupInformation23) AddCancellationReasonInformation() *CancellationReasonInformation3 {
	newValue := new(CancellationReasonInformation3)
	o.CancellationReasonInformation = append(o.CancellationReasonInformation, newValue)
	return newValue
}

// Set of elements used to provide information on the original group, to which the message refers.
type OriginalGroupInformation24 struct {

	// Unique identification, as assigned by the original assigner, to unambiguously identify the original group cancellation request.
	OriginalGroupCancellationIdentification *Max35Text `xml:"OrgnlGrpCxlId,omitempty"`

	// Identifies the case.
	ResolvedCase *Case2 `xml:"RslvdCase,omitempty"`

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Specifies the original message name identifier to which the message refers.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Number of individual transactions contained in the original message.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original message, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Specifies the status of a group cancellation request.
	GroupCancellationStatus *GroupCancellationStatus1Code `xml:"GrpCxlSts,omitempty"`

	// Set of elements used to provide detailed information on the cancellation status reason.
	CancellationStatusReasonInformation []*CancellationStatusReasonInformation1 `xml:"CxlStsRsnInf,omitempty"`

	// Detailed information on the number of transactions for each identical cancellation status.
	NumberOfTransactionsPerCancellationStatus []*NumberOfTransactionsPerStatus1 `xml:"NbOfTxsPerCxlSts,omitempty"`
}

func (o *OriginalGroupInformation24) SetOriginalGroupCancellationIdentification(value string) {
	o.OriginalGroupCancellationIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation24) AddResolvedCase() *Case2 {
	o.ResolvedCase = new(Case2)
	return o.ResolvedCase
}

func (o *OriginalGroupInformation24) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation24) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation24) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalGroupInformation24) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalGroupInformation24) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalGroupInformation24) SetGroupCancellationStatus(value string) {
	o.GroupCancellationStatus = (*GroupCancellationStatus1Code)(&value)
}

func (o *OriginalGroupInformation24) AddCancellationStatusReasonInformation() *CancellationStatusReasonInformation1 {
	newValue := new(CancellationStatusReasonInformation1)
	o.CancellationStatusReasonInformation = append(o.CancellationStatusReasonInformation, newValue)
	return newValue
}

func (o *OriginalGroupInformation24) AddNumberOfTransactionsPerCancellationStatus() *NumberOfTransactionsPerStatus1 {
	newValue := new(NumberOfTransactionsPerStatus1)
	o.NumberOfTransactionsPerCancellationStatus = append(o.NumberOfTransactionsPerCancellationStatus, newValue)
	return newValue
}

// Set of elements used to provide information on the original group, to which the message refers.
type OriginalGroupInformation25 struct {

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Specifies the original message name identifier to which the message refers.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Number of individual transactions contained in the original message.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original message, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Specifies the status of a group of transactions.
	GroupStatus *TransactionGroupStatus3Code `xml:"GrpSts,omitempty"`

	// Set of elements used to provide detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Detailed information on the number of transactions for each identical transaction status.
	NumberOfTransactionsPerStatus []*NumberOfTransactionsPerStatus3 `xml:"NbOfTxsPerSts,omitempty"`
}

func (o *OriginalGroupInformation25) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation25) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation25) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalGroupInformation25) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalGroupInformation25) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalGroupInformation25) SetGroupStatus(value string) {
	o.GroupStatus = (*TransactionGroupStatus3Code)(&value)
}

func (o *OriginalGroupInformation25) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	o.StatusReasonInformation = append(o.StatusReasonInformation, newValue)
	return newValue
}

func (o *OriginalGroupInformation25) AddNumberOfTransactionsPerStatus() *NumberOfTransactionsPerStatus3 {
	newValue := new(NumberOfTransactionsPerStatus3)
	o.NumberOfTransactionsPerStatus = append(o.NumberOfTransactionsPerStatus, newValue)
	return newValue
}

// Set of elements used to provide information on the original group, to which the message refers.
type OriginalGroupInformation27 struct {

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Specifies the original message name identifier to which the message refers.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Number of individual transactions contained in the original message.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original message, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`
}

func (o *OriginalGroupInformation27) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation27) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation27) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalGroupInformation27) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalGroupInformation27) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

// Set of elements used to provide information on the original group, to which the message refers.
type OriginalGroupInformation28 struct {

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Specifies the original message name identifier to which the message refers.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Number of individual transactions contained in the original message.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original message, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Specifies the status of a group of transactions.
	GroupStatus *ExternalPaymentGroupStatus1Code `xml:"GrpSts,omitempty"`

	// Set of elements used to provide detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Detailed information on the number of transactions for each identical transaction status.
	NumberOfTransactionsPerStatus []*NumberOfTransactionsPerStatus5 `xml:"NbOfTxsPerSts,omitempty"`
}

func (o *OriginalGroupInformation28) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation28) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation28) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalGroupInformation28) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalGroupInformation28) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalGroupInformation28) SetGroupStatus(value string) {
	o.GroupStatus = (*ExternalPaymentGroupStatus1Code)(&value)
}

func (o *OriginalGroupInformation28) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	o.StatusReasonInformation = append(o.StatusReasonInformation, newValue)
	return newValue
}

func (o *OriginalGroupInformation28) AddNumberOfTransactionsPerStatus() *NumberOfTransactionsPerStatus5 {
	newValue := new(NumberOfTransactionsPerStatus5)
	o.NumberOfTransactionsPerStatus = append(o.NumberOfTransactionsPerStatus, newValue)
	return newValue
}

// Unique and unambiguous identifier of the group of transactions as assigned by the original instructing party.
type OriginalGroupInformation3 struct {

	// Point to point reference assigned by the original instructing party to unambiguously identify the original group of individual transactions.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Specifies the original message name identifier to which the message refers, eg, pacs.003.001.01 or MT103.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Original date and time at which the message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`
}

func (o *OriginalGroupInformation3) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation3) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation3) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

// Unique and unambiguous identifier of the group of transactions as assigned by the original instructing party.
type OriginalGroupInformation4 struct {

	// Point to point reference assigned by the original instructing party to unambiguously identify the original group of individual transactions.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Specifies the original message name identifier to which the message refers, eg, pacs.003.001.01 or MT103.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Original date and time at which the message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Detailed information on the cancellation reason.
	CancellationReasonInformation []*CancellationReasonInformation1 `xml:"CxlRsnInf,omitempty"`
}

func (o *OriginalGroupInformation4) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation4) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation4) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalGroupInformation4) AddCancellationReasonInformation() *CancellationReasonInformation1 {
	newValue := new(CancellationReasonInformation1)
	o.CancellationReasonInformation = append(o.CancellationReasonInformation, newValue)
	return newValue
}

// Unique and unambiguous identifier of the group of transactions as assigned by the original instructing party.
type OriginalGroupInformation5 struct {

	// Point to point reference assigned by the original instructing party to unambiguously identify the original group of individual transactions.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Specifies the original message name identifier to which the message refers, eg, pacs.003.001.01 or MT103.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Detailed information on the reversal reason.
	ReversalReasonInformation []*ReversalReasonInformation1 `xml:"RvslRsnInf,omitempty"`
}

func (o *OriginalGroupInformation5) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation5) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation5) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalGroupInformation5) AddReversalReasonInformation() *ReversalReasonInformation1 {
	newValue := new(ReversalReasonInformation1)
	o.ReversalReasonInformation = append(o.ReversalReasonInformation, newValue)
	return newValue
}
