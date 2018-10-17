package iso20022

// Status and reason of an instructed order.
type StatusAndReason1 struct {

	// Status and reason for the transaction.
	StatusAndReason *Status2Choice `xml:"StsAndRsn"`

	// Details of the transactions reported.
	Transaction []*Transaction7 `xml:"Tx,omitempty"`
}

func (s *StatusAndReason1) AddStatusAndReason() *Status2Choice {
	s.StatusAndReason = new(Status2Choice)
	return s.StatusAndReason
}

func (s *StatusAndReason1) AddTransaction() *Transaction7 {
	newValue := new(Transaction7)
	s.Transaction = append(s.Transaction, newValue)
	return newValue
}

// Provides details related to the status of the order.
type StatusAndReason10 struct {

	// Status of affirmation of a trade.
	AffirmationStatus *AffirmationStatus7Choice `xml:"AffirmSts"`

	// Specifies the reason why the instruction has an unaffirmed status.
	UnaffirmedReason *UnaffirmedReason2Choice `xml:"UaffrmdRsn,omitempty"`

	// Provides additional information about the reason in narrative form.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (s *StatusAndReason10) AddAffirmationStatus() *AffirmationStatus7Choice {
	s.AffirmationStatus = new(AffirmationStatus7Choice)
	return s.AffirmationStatus
}

func (s *StatusAndReason10) AddUnaffirmedReason() *UnaffirmedReason2Choice {
	s.UnaffirmedReason = new(UnaffirmedReason2Choice)
	return s.UnaffirmedReason
}

func (s *StatusAndReason10) SetAdditionalReasonInformation(value string) {
	s.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Status and reason of a transaction.
type StatusAndReason12 struct {

	// Provides the status of an instruction.
	ProcessingStatus *ProcessingStatus23Choice `xml:"PrcgSts,omitempty"`

	// Provides the matching status of an instruction as known by the account servicer based on an allegement. At this time no matching took place on the market (at the CSD/ICSD).
	InferredMatchingStatus *MatchingStatus7Choice `xml:"IfrrdMtchgSts,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus7Choice `xml:"MtchgSts,omitempty"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *SettlementStatus7Choice `xml:"SttlmSts,omitempty"`
}

func (s *StatusAndReason12) AddProcessingStatus() *ProcessingStatus23Choice {
	s.ProcessingStatus = new(ProcessingStatus23Choice)
	return s.ProcessingStatus
}

func (s *StatusAndReason12) AddInferredMatchingStatus() *MatchingStatus7Choice {
	s.InferredMatchingStatus = new(MatchingStatus7Choice)
	return s.InferredMatchingStatus
}

func (s *StatusAndReason12) AddMatchingStatus() *MatchingStatus7Choice {
	s.MatchingStatus = new(MatchingStatus7Choice)
	return s.MatchingStatus
}

func (s *StatusAndReason12) AddSettlementStatus() *SettlementStatus7Choice {
	s.SettlementStatus = new(SettlementStatus7Choice)
	return s.SettlementStatus
}

// Status and reason of an instructed order.
type StatusAndReason16 struct {

	// Status and reason for the transaction.
	StatusAndReason *Status9Choice `xml:"StsAndRsn"`

	// Details of the transactions reported.
	Transaction []*Transaction28 `xml:"Tx,omitempty"`
}

func (s *StatusAndReason16) AddStatusAndReason() *Status9Choice {
	s.StatusAndReason = new(Status9Choice)
	return s.StatusAndReason
}

func (s *StatusAndReason16) AddTransaction() *Transaction28 {
	newValue := new(Transaction28)
	s.Transaction = append(s.Transaction, newValue)
	return newValue
}

// Status and reason of an instructed order.
type StatusAndReason18 struct {

	// Status and reason for the transaction.
	StatusAndReason *Status15Choice `xml:"StsAndRsn"`

	// Details of the transactions reported.
	Transaction []*Transaction35 `xml:"Tx,omitempty"`
}

func (s *StatusAndReason18) AddStatusAndReason() *Status15Choice {
	s.StatusAndReason = new(Status15Choice)
	return s.StatusAndReason
}

func (s *StatusAndReason18) AddTransaction() *Transaction35 {
	newValue := new(Transaction35)
	s.Transaction = append(s.Transaction, newValue)
	return newValue
}

// Status and reason of a transaction.
type StatusAndReason19 struct {

	// Provides the status of an instruction.
	ProcessingStatus *ProcessingStatus23Choice `xml:"PrcgSts,omitempty"`

	// Provides the matching status of an instruction as known by the account servicer based on an allegement. At this time no matching took place on the market (at the CSD/ICSD).
	InferredMatchingStatus *MatchingStatus19Choice `xml:"IfrrdMtchgSts,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus19Choice `xml:"MtchgSts,omitempty"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *SettlementStatus7Choice `xml:"SttlmSts,omitempty"`
}

func (s *StatusAndReason19) AddProcessingStatus() *ProcessingStatus23Choice {
	s.ProcessingStatus = new(ProcessingStatus23Choice)
	return s.ProcessingStatus
}

func (s *StatusAndReason19) AddInferredMatchingStatus() *MatchingStatus19Choice {
	s.InferredMatchingStatus = new(MatchingStatus19Choice)
	return s.InferredMatchingStatus
}

func (s *StatusAndReason19) AddMatchingStatus() *MatchingStatus19Choice {
	s.MatchingStatus = new(MatchingStatus19Choice)
	return s.MatchingStatus
}

func (s *StatusAndReason19) AddSettlementStatus() *SettlementStatus7Choice {
	s.SettlementStatus = new(SettlementStatus7Choice)
	return s.SettlementStatus
}

// Status and reason of an instructed order.
type StatusAndReason2 struct {

	// Status and reason for the transaction.
	StatusAndReason *Status2Choice `xml:"StsAndRsn"`
}

func (s *StatusAndReason2) AddStatusAndReason() *Status2Choice {
	s.StatusAndReason = new(Status2Choice)
	return s.StatusAndReason
}

// Status and reason of an instructed order.
type StatusAndReason25 struct {

	// Status and reason for the transaction.
	StatusAndReason *Status15Choice `xml:"StsAndRsn"`

	// Details of the transactions reported.
	Transaction []*Transaction40 `xml:"Tx,omitempty"`
}

func (s *StatusAndReason25) AddStatusAndReason() *Status15Choice {
	s.StatusAndReason = new(Status15Choice)
	return s.StatusAndReason
}

func (s *StatusAndReason25) AddTransaction() *Transaction40 {
	newValue := new(Transaction40)
	s.Transaction = append(s.Transaction, newValue)
	return newValue
}

// Status and reason of an instructed order.
type StatusAndReason27 struct {

	// Status and reason for the transaction.
	StatusAndReason *Status18Choice `xml:"StsAndRsn"`

	// Details of the transactions reported.
	Transaction []*Transaction45 `xml:"Tx,omitempty"`
}

func (s *StatusAndReason27) AddStatusAndReason() *Status18Choice {
	s.StatusAndReason = new(Status18Choice)
	return s.StatusAndReason
}

func (s *StatusAndReason27) AddTransaction() *Transaction45 {
	newValue := new(Transaction45)
	s.Transaction = append(s.Transaction, newValue)
	return newValue
}

// Status and reason of a transaction.
type StatusAndReason28 struct {

	// Provides the status of an instruction.
	ProcessingStatus *ProcessingStatus52Choice `xml:"PrcgSts,omitempty"`

	// Provides the matching status of an instruction as known by the account servicer based on an allegement. At this time no matching took place on the market (at the CSD/ICSD).
	InferredMatchingStatus *MatchingStatus24Choice `xml:"IfrrdMtchgSts,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus24Choice `xml:"MtchgSts,omitempty"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *SettlementStatus17Choice `xml:"SttlmSts,omitempty"`
}

func (s *StatusAndReason28) AddProcessingStatus() *ProcessingStatus52Choice {
	s.ProcessingStatus = new(ProcessingStatus52Choice)
	return s.ProcessingStatus
}

func (s *StatusAndReason28) AddInferredMatchingStatus() *MatchingStatus24Choice {
	s.InferredMatchingStatus = new(MatchingStatus24Choice)
	return s.InferredMatchingStatus
}

func (s *StatusAndReason28) AddMatchingStatus() *MatchingStatus24Choice {
	s.MatchingStatus = new(MatchingStatus24Choice)
	return s.MatchingStatus
}

func (s *StatusAndReason28) AddSettlementStatus() *SettlementStatus17Choice {
	s.SettlementStatus = new(SettlementStatus17Choice)
	return s.SettlementStatus
}

// Status and reason of a transaction.
type StatusAndReason29 struct {

	// Provides the status of an instruction.
	ProcessingStatus *ProcessingStatus62Choice `xml:"PrcgSts,omitempty"`

	// Provides the matching status of an instruction as known by the account servicer based on an allegement. At this time no matching took place on the market (at the CSD/ICSD).
	InferredMatchingStatus *MatchingStatus32Choice `xml:"IfrrdMtchgSts,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus32Choice `xml:"MtchgSts,omitempty"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *SettlementStatus22Choice `xml:"SttlmSts,omitempty"`
}

func (s *StatusAndReason29) AddProcessingStatus() *ProcessingStatus62Choice {
	s.ProcessingStatus = new(ProcessingStatus62Choice)
	return s.ProcessingStatus
}

func (s *StatusAndReason29) AddInferredMatchingStatus() *MatchingStatus32Choice {
	s.InferredMatchingStatus = new(MatchingStatus32Choice)
	return s.InferredMatchingStatus
}

func (s *StatusAndReason29) AddMatchingStatus() *MatchingStatus32Choice {
	s.MatchingStatus = new(MatchingStatus32Choice)
	return s.MatchingStatus
}

func (s *StatusAndReason29) AddSettlementStatus() *SettlementStatus22Choice {
	s.SettlementStatus = new(SettlementStatus22Choice)
	return s.SettlementStatus
}

// Status and reason of a transaction.
type StatusAndReason3 struct {

	// Provides the status of an instruction.
	ProcessingStatus *ProcessingStatus6Choice `xml:"PrcgSts,omitempty"`

	// Provides the matching status of an instruction as known by the account servicer based on an allegement. At this time no matching took place on the market (at the CSD/ICSD).
	InferredMatchingStatus *MatchingStatus2Choice `xml:"IfrrdMtchgSts,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus2Choice `xml:"MtchgSts,omitempty"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *SettlementStatus2Choice `xml:"SttlmSts,omitempty"`
}

func (s *StatusAndReason3) AddProcessingStatus() *ProcessingStatus6Choice {
	s.ProcessingStatus = new(ProcessingStatus6Choice)
	return s.ProcessingStatus
}

func (s *StatusAndReason3) AddInferredMatchingStatus() *MatchingStatus2Choice {
	s.InferredMatchingStatus = new(MatchingStatus2Choice)
	return s.InferredMatchingStatus
}

func (s *StatusAndReason3) AddMatchingStatus() *MatchingStatus2Choice {
	s.MatchingStatus = new(MatchingStatus2Choice)
	return s.MatchingStatus
}

func (s *StatusAndReason3) AddSettlementStatus() *SettlementStatus2Choice {
	s.SettlementStatus = new(SettlementStatus2Choice)
	return s.SettlementStatus
}

// Status and reason of an instructed order.
type StatusAndReason30 struct {

	// Status and reason for the transaction.
	StatusAndReason *Status23Choice `xml:"StsAndRsn"`

	// Details of the transactions reported.
	Transaction []*Transaction48 `xml:"Tx,omitempty"`
}

func (s *StatusAndReason30) AddStatusAndReason() *Status23Choice {
	s.StatusAndReason = new(Status23Choice)
	return s.StatusAndReason
}

func (s *StatusAndReason30) AddTransaction() *Transaction48 {
	newValue := new(Transaction48)
	s.Transaction = append(s.Transaction, newValue)
	return newValue
}

// Status and reason of an instructed order.
type StatusAndReason32 struct {

	// Status and reason for the transaction.
	StatusAndReason *Status18Choice `xml:"StsAndRsn"`

	// Details of the transactions reported.
	Transaction []*Transaction54 `xml:"Tx,omitempty"`
}

func (s *StatusAndReason32) AddStatusAndReason() *Status18Choice {
	s.StatusAndReason = new(Status18Choice)
	return s.StatusAndReason
}

func (s *StatusAndReason32) AddTransaction() *Transaction54 {
	newValue := new(Transaction54)
	s.Transaction = append(s.Transaction, newValue)
	return newValue
}

// Status and reason of an instructed order.
type StatusAndReason33 struct {

	// Status and reason for the transaction.
	StatusAndReason *Status23Choice `xml:"StsAndRsn"`

	// Details of the transactions reported.
	Transaction []*Transaction56 `xml:"Tx,omitempty"`
}

func (s *StatusAndReason33) AddStatusAndReason() *Status23Choice {
	s.StatusAndReason = new(Status23Choice)
	return s.StatusAndReason
}

func (s *StatusAndReason33) AddTransaction() *Transaction56 {
	newValue := new(Transaction56)
	s.Transaction = append(s.Transaction, newValue)
	return newValue
}

// Status and reason of an instructed order.
type StatusAndReason7 struct {

	// Status and reason for the transaction.
	StatusAndReason *Status2Choice `xml:"StsAndRsn"`

	// Details of the transactions reported.
	Transaction []*Transaction14 `xml:"Tx,omitempty"`
}

func (s *StatusAndReason7) AddStatusAndReason() *Status2Choice {
	s.StatusAndReason = new(Status2Choice)
	return s.StatusAndReason
}

func (s *StatusAndReason7) AddTransaction() *Transaction14 {
	newValue := new(Transaction14)
	s.Transaction = append(s.Transaction, newValue)
	return newValue
}

// Status and reason of an instructed order.
type StatusAndReason9 struct {

	// Status and reason for the transaction.
	StatusAndReason *Status9Choice `xml:"StsAndRsn"`

	// Details of the transactions reported.
	Transaction []*Transaction20 `xml:"Tx,omitempty"`
}

func (s *StatusAndReason9) AddStatusAndReason() *Status9Choice {
	s.StatusAndReason = new(Status9Choice)
	return s.StatusAndReason
}

func (s *StatusAndReason9) AddTransaction() *Transaction20 {
	newValue := new(Transaction20)
	s.Transaction = append(s.Transaction, newValue)
	return newValue
}
