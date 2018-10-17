package iso20022

// Set of elements used to provide further means of referencing a payment transaction.
type PaymentIdentification1 struct {

	// Unique identification as assigned by an instructing party for an instructed party to unambiguously identify the instruction.
	//
	// Usage: the  instruction identification is a point to point reference that can be used between the instructing party and the instructed party to refer to the individual instruction. It can be included in several messages related to the instruction.
	InstructionIdentification *Max35Text `xml:"InstrId,omitempty"`

	// Unique identification assigned by the initiating party to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	//
	// Usage: The end-to-end identification can be used for reconciliation or to link tasks relating to the transaction. It can be included in several messages related to the transaction.
	EndToEndIdentification *Max35Text `xml:"EndToEndId"`
}

func (p *PaymentIdentification1) SetInstructionIdentification(value string) {
	p.InstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentIdentification1) SetEndToEndIdentification(value string) {
	p.EndToEndIdentification = (*Max35Text)(&value)
}

// Set of elements used to provide further means of referencing a payment transaction.
type PaymentIdentification2 struct {

	// Unique identification as assigned by an instructing party for an instructed party to unambiguously identify the instruction.
	//
	// Usage: the  instruction identification is a point to point reference that can be used between the instructing party and the instructed party to refer to the individual instruction. It can be included in several messages related to the instruction.
	InstructionIdentification *Max35Text `xml:"InstrId,omitempty"`

	// Unique identification assigned by the initiating party to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	//
	// Usage: The end-to-end identification can be used for reconciliation or to link tasks relating to the transaction. It can be included in several messages related to the transaction.
	//
	// Usage: In case there are technical limitations to carry on multiple references, the end-to-end identification must be carried on throughout the entire end-to-end chain.
	EndToEndIdentification *Max35Text `xml:"EndToEndId"`

	// Unique identification assigned by the first instructing agent to unambiguously identify the transaction and passed on, unchanged, throughout the entire interbank chain.
	//
	// Usage: The transaction identification can be used for reconciliation, tracking or to link tasks relating to the transaction on the interbank level.
	// Usage: The instructing agent has to make sure the transaction identification is unique for a pre-agreed period.
	TransactionIdentification *Max35Text `xml:"TxId"`
}

func (p *PaymentIdentification2) SetInstructionIdentification(value string) {
	p.InstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentIdentification2) SetEndToEndIdentification(value string) {
	p.EndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentIdentification2) SetTransactionIdentification(value string) {
	p.TransactionIdentification = (*Max35Text)(&value)
}

// Set of elements used to provide further means of referencing a payment transaction.
type PaymentIdentification3 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the instruction.
	//
	// Usage: The instruction identification is a point to point reference that can be used between the instructing party and the instructed party to refer to the individual instruction. It can be included in several messages related to the instruction.
	InstructionIdentification *Max35Text `xml:"InstrId,omitempty"`

	// Unique identification, as assigned by the initiating party, to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	//
	// Usage: The end-to-end identification can be used for reconciliation or to link tasks relating to the transaction. It can be included in several messages related to the transaction.
	//
	// Usage: In case there are technical limitations to pass on multiple references, the end-to-end identification must be passed on throughout the entire end-to-end chain.
	EndToEndIdentification *Max35Text `xml:"EndToEndId"`

	// Unique identification, as assigned by the first instructing agent, to unambiguously identify the transaction that is passed on, unchanged, throughout the entire interbank chain.
	// Usage: The transaction identification can be used for reconciliation, tracking or to link tasks relating to the transaction on the interbank level.
	// Usage: The instructing agent has to make sure that the transaction identification is unique for a pre-agreed period.
	TransactionIdentification *Max35Text `xml:"TxId"`

	// Unique reference, as assigned by a clearing system, to unambiguously identify the instruction.
	ClearingSystemReference *Max35Text `xml:"ClrSysRef,omitempty"`
}

func (p *PaymentIdentification3) SetInstructionIdentification(value string) {
	p.InstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentIdentification3) SetEndToEndIdentification(value string) {
	p.EndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentIdentification3) SetTransactionIdentification(value string) {
	p.TransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentIdentification3) SetClearingSystemReference(value string) {
	p.ClearingSystemReference = (*Max35Text)(&value)
}

// Provides further means of referencing a payment transaction.
type PaymentIdentification4 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the instruction.
	//
	// Usage: The instruction identification is a point to point reference that can be used between the instructing party and the instructed party to refer to the individual instruction. It can be included in several messages related to the instruction.
	InstructionIdentification *Max35Text `xml:"InstrId,omitempty"`

	// Unique identification, as assigned by the initiating party, to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	//
	// Usage: The end-to-end identification can be used for reconciliation or to link tasks relating to the transaction. It can be included in several messages related to the transaction.
	//
	// Usage: In case there are technical limitations to pass on multiple references, the end-to-end identification must be passed on throughout the entire end-to-end chain.
	EndToEndIdentification *Max35Text `xml:"EndToEndId"`

	// Unique identification, as assigned by the first instructing agent, to unambiguously identify the transaction that is passed on, unchanged, throughout the entire interbank chain.
	// Usage: The transaction identification can be used for reconciliation, tracking or to link tasks relating to the transaction on the interbank level.
	// Usage: The instructing agent has to make sure that the transaction identification is unique for a pre-agreed period.
	TransactionIdentification *Max35Text `xml:"TxId"`

	// Unique reference, as assigned by a clearing system, to unambiguously identify the instruction.
	ClearingSystemReference *Max35Text `xml:"ClrSysRef,omitempty"`

	// Identifies the first agent in the identification chain, following the payment initiating party.
	FirstAgent *BranchAndFinancialInstitutionIdentification5 `xml:"FrstAgt,omitempty"`
}

func (p *PaymentIdentification4) SetInstructionIdentification(value string) {
	p.InstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentIdentification4) SetEndToEndIdentification(value string) {
	p.EndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentIdentification4) SetTransactionIdentification(value string) {
	p.TransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentIdentification4) SetClearingSystemReference(value string) {
	p.ClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentIdentification4) AddFirstAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.FirstAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.FirstAgent
}
