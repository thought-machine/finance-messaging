package iso20022

// Information related to a linked transaction.
type Linkages1 struct {

	// When the transaction is to be executed relative to a linked transaction.
	ProcessingPosition *ProcessingPosition1Choice `xml:"PrcgPos,omitempty"`

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber1Choice `xml:"MsgNb,omitempty"`

	// Reference to the linked transaction.
	Reference *References1Choice `xml:"Ref"`

	// Quantity of financial instruments of the linked transaction to be paired-off or turned.
	LinkedQuantity *PairedOrTurnedQuantity1Choice `xml:"LkdQty,omitempty"`
}

func (l *Linkages1) AddProcessingPosition() *ProcessingPosition1Choice {
	l.ProcessingPosition = new(ProcessingPosition1Choice)
	return l.ProcessingPosition
}

func (l *Linkages1) AddMessageNumber() *DocumentNumber1Choice {
	l.MessageNumber = new(DocumentNumber1Choice)
	return l.MessageNumber
}

func (l *Linkages1) AddReference() *References1Choice {
	l.Reference = new(References1Choice)
	return l.Reference
}

func (l *Linkages1) AddLinkedQuantity() *PairedOrTurnedQuantity1Choice {
	l.LinkedQuantity = new(PairedOrTurnedQuantity1Choice)
	return l.LinkedQuantity
}

// Information related to a linked transaction.
type Linkages10 struct {

	// When the transaction is to be executed relative to a linked transaction.
	ProcessingPosition *ProcessingPosition2Choice `xml:"PrcgPos,omitempty"`

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber1Choice `xml:"MsgNb,omitempty"`

	// Reference to the linked transaction.
	Reference *References13Choice `xml:"Ref"`

	// Party that generates the reference.
	ReferenceOwner *PartyIdentification36Choice `xml:"RefOwnr,omitempty"`
}

func (l *Linkages10) AddProcessingPosition() *ProcessingPosition2Choice {
	l.ProcessingPosition = new(ProcessingPosition2Choice)
	return l.ProcessingPosition
}

func (l *Linkages10) AddMessageNumber() *DocumentNumber1Choice {
	l.MessageNumber = new(DocumentNumber1Choice)
	return l.MessageNumber
}

func (l *Linkages10) AddReference() *References13Choice {
	l.Reference = new(References13Choice)
	return l.Reference
}

func (l *Linkages10) AddReferenceOwner() *PartyIdentification36Choice {
	l.ReferenceOwner = new(PartyIdentification36Choice)
	return l.ReferenceOwner
}

// Information related to a linked transaction.
type Linkages15 struct {

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber4Choice `xml:"MsgNb,omitempty"`

	// Reference to the linked transaction.
	Reference *IdentificationReference8Choice `xml:"Ref"`
}

func (l *Linkages15) AddMessageNumber() *DocumentNumber4Choice {
	l.MessageNumber = new(DocumentNumber4Choice)
	return l.MessageNumber
}

func (l *Linkages15) AddReference() *IdentificationReference8Choice {
	l.Reference = new(IdentificationReference8Choice)
	return l.Reference
}

// Information related to a linked transaction.
type Linkages16 struct {

	// When the transaction is to be executed relative to a linked transaction.
	ProcessingPosition *ProcessingPosition2Choice `xml:"PrcgPos,omitempty"`

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber1Choice `xml:"MsgNb,omitempty"`

	// Reference to the linked transaction.
	Reference *References24Choice `xml:"Ref"`

	// Party that generates the reference.
	ReferenceOwner *PartyIdentification36Choice `xml:"RefOwnr,omitempty"`
}

func (l *Linkages16) AddProcessingPosition() *ProcessingPosition2Choice {
	l.ProcessingPosition = new(ProcessingPosition2Choice)
	return l.ProcessingPosition
}

func (l *Linkages16) AddMessageNumber() *DocumentNumber1Choice {
	l.MessageNumber = new(DocumentNumber1Choice)
	return l.MessageNumber
}

func (l *Linkages16) AddReference() *References24Choice {
	l.Reference = new(References24Choice)
	return l.Reference
}

func (l *Linkages16) AddReferenceOwner() *PartyIdentification36Choice {
	l.ReferenceOwner = new(PartyIdentification36Choice)
	return l.ReferenceOwner
}

// Information related to a linked transaction.
type Linkages17 struct {

	// When the transaction is to be executed relative to a linked transaction.
	ProcessingPosition *ProcessingPosition1Choice `xml:"PrcgPos,omitempty"`

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber1Choice `xml:"MsgNb,omitempty"`

	// Reference to the linked transaction.
	Reference *References25Choice `xml:"Ref"`

	// Quantity of financial instruments of the linked transaction to be paired-off or turned.
	LinkedQuantity *PairedOrTurnedQuantity1Choice `xml:"LkdQty,omitempty"`

	// Party that generates the reference.
	ReferenceOwner *PartyIdentification36Choice `xml:"RefOwnr,omitempty"`
}

func (l *Linkages17) AddProcessingPosition() *ProcessingPosition1Choice {
	l.ProcessingPosition = new(ProcessingPosition1Choice)
	return l.ProcessingPosition
}

func (l *Linkages17) AddMessageNumber() *DocumentNumber1Choice {
	l.MessageNumber = new(DocumentNumber1Choice)
	return l.MessageNumber
}

func (l *Linkages17) AddReference() *References25Choice {
	l.Reference = new(References25Choice)
	return l.Reference
}

func (l *Linkages17) AddLinkedQuantity() *PairedOrTurnedQuantity1Choice {
	l.LinkedQuantity = new(PairedOrTurnedQuantity1Choice)
	return l.LinkedQuantity
}

func (l *Linkages17) AddReferenceOwner() *PartyIdentification36Choice {
	l.ReferenceOwner = new(PartyIdentification36Choice)
	return l.ReferenceOwner
}

// Information related to a linked transaction.
type Linkages18 struct {

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber4Choice `xml:"MsgNb,omitempty"`

	// Reference to the linked transaction.
	Reference *IdentificationReference11Choice `xml:"Ref"`
}

func (l *Linkages18) AddMessageNumber() *DocumentNumber4Choice {
	l.MessageNumber = new(DocumentNumber4Choice)
	return l.MessageNumber
}

func (l *Linkages18) AddReference() *IdentificationReference11Choice {
	l.Reference = new(IdentificationReference11Choice)
	return l.Reference
}

// Information related to a linked transaction.
type Linkages19 struct {

	// When the transaction is to be executed relative to a linked transaction.
	ProcessingPosition *ProcessingPosition1Choice `xml:"PrcgPos,omitempty"`

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber1Choice `xml:"MsgNb,omitempty"`

	// Reference to the linked transaction.
	Reference *References25Choice `xml:"Ref"`

	// Party that generates the reference.
	ReferenceOwner *PartyIdentification36Choice `xml:"RefOwnr,omitempty"`
}

func (l *Linkages19) AddProcessingPosition() *ProcessingPosition1Choice {
	l.ProcessingPosition = new(ProcessingPosition1Choice)
	return l.ProcessingPosition
}

func (l *Linkages19) AddMessageNumber() *DocumentNumber1Choice {
	l.MessageNumber = new(DocumentNumber1Choice)
	return l.MessageNumber
}

func (l *Linkages19) AddReference() *References25Choice {
	l.Reference = new(References25Choice)
	return l.Reference
}

func (l *Linkages19) AddReferenceOwner() *PartyIdentification36Choice {
	l.ReferenceOwner = new(PartyIdentification36Choice)
	return l.ReferenceOwner
}

// Information related to a linked transaction.
type Linkages2 struct {

	// When the transaction is to be executed relative to a linked transaction.
	ProcessingPosition *ProcessingPosition1Choice `xml:"PrcgPos,omitempty"`

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber1Choice `xml:"MsgNb,omitempty"`

	// Reference to the linked transaction.
	Reference *References1Choice `xml:"Ref"`
}

func (l *Linkages2) AddProcessingPosition() *ProcessingPosition1Choice {
	l.ProcessingPosition = new(ProcessingPosition1Choice)
	return l.ProcessingPosition
}

func (l *Linkages2) AddMessageNumber() *DocumentNumber1Choice {
	l.MessageNumber = new(DocumentNumber1Choice)
	return l.MessageNumber
}

func (l *Linkages2) AddReference() *References1Choice {
	l.Reference = new(References1Choice)
	return l.Reference
}

// Information related to a linked transaction.
type Linkages21 struct {

	// Reference to the linked transaction.
	Reference *References1Choice `xml:"Ref"`
}

func (l *Linkages21) AddReference() *References1Choice {
	l.Reference = new(References1Choice)
	return l.Reference
}

// Information related to a linked transaction.
type Linkages27 struct {

	// When the transaction is to be executed relative to a linked transaction.
	ProcessingPosition *ProcessingPosition5Choice `xml:"PrcgPos,omitempty"`

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber1Choice `xml:"MsgNb,omitempty"`

	// Reference to the linked transaction.
	Reference *References24Choice `xml:"Ref"`

	// Party that generates the reference.
	ReferenceOwner *PartyIdentification36Choice `xml:"RefOwnr,omitempty"`
}

func (l *Linkages27) AddProcessingPosition() *ProcessingPosition5Choice {
	l.ProcessingPosition = new(ProcessingPosition5Choice)
	return l.ProcessingPosition
}

func (l *Linkages27) AddMessageNumber() *DocumentNumber1Choice {
	l.MessageNumber = new(DocumentNumber1Choice)
	return l.MessageNumber
}

func (l *Linkages27) AddReference() *References24Choice {
	l.Reference = new(References24Choice)
	return l.Reference
}

func (l *Linkages27) AddReferenceOwner() *PartyIdentification36Choice {
	l.ReferenceOwner = new(PartyIdentification36Choice)
	return l.ReferenceOwner
}

// Information related to a linked transaction.
type Linkages3 struct {

	// When the transaction is to be executed relative to a linked transaction.
	ProcessingPosition *ProcessingPosition2Choice `xml:"PrcgPos,omitempty"`

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber1Choice `xml:"MsgNb,omitempty"`

	// Reference to the linked transaction.
	Reference *References6Choice `xml:"Ref"`
}

func (l *Linkages3) AddProcessingPosition() *ProcessingPosition2Choice {
	l.ProcessingPosition = new(ProcessingPosition2Choice)
	return l.ProcessingPosition
}

func (l *Linkages3) AddMessageNumber() *DocumentNumber1Choice {
	l.MessageNumber = new(DocumentNumber1Choice)
	return l.MessageNumber
}

func (l *Linkages3) AddReference() *References6Choice {
	l.Reference = new(References6Choice)
	return l.Reference
}

// Information related to a linked transaction.
type Linkages36 struct {

	// When the transaction is to be executed relative to a linked transaction.
	ProcessingPosition *ProcessingPosition7Choice `xml:"PrcgPos,omitempty"`

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber5Choice `xml:"MsgNb,omitempty"`

	// Reference to the linked transaction.
	Reference *References41Choice `xml:"Ref"`

	// Party that generates the reference.
	ReferenceOwner *PartyIdentification92Choice `xml:"RefOwnr,omitempty"`
}

func (l *Linkages36) AddProcessingPosition() *ProcessingPosition7Choice {
	l.ProcessingPosition = new(ProcessingPosition7Choice)
	return l.ProcessingPosition
}

func (l *Linkages36) AddMessageNumber() *DocumentNumber5Choice {
	l.MessageNumber = new(DocumentNumber5Choice)
	return l.MessageNumber
}

func (l *Linkages36) AddReference() *References41Choice {
	l.Reference = new(References41Choice)
	return l.Reference
}

func (l *Linkages36) AddReferenceOwner() *PartyIdentification92Choice {
	l.ReferenceOwner = new(PartyIdentification92Choice)
	return l.ReferenceOwner
}

// Information related to a linked transaction.
type Linkages37 struct {

	// When the transaction is to be executed relative to a linked transaction.
	ProcessingPosition *ProcessingPosition7Choice `xml:"PrcgPos,omitempty"`

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber5Choice `xml:"MsgNb,omitempty"`

	// Reference to the linked transaction.
	Reference *References41Choice `xml:"Ref"`

	// Quantity of financial instruments of the linked transaction to be paired-off or turned.
	LinkedQuantity *PairedOrTurnedQuantity3Choice `xml:"LkdQty,omitempty"`

	// Party that generates the reference.
	ReferenceOwner *PartyIdentification92Choice `xml:"RefOwnr,omitempty"`
}

func (l *Linkages37) AddProcessingPosition() *ProcessingPosition7Choice {
	l.ProcessingPosition = new(ProcessingPosition7Choice)
	return l.ProcessingPosition
}

func (l *Linkages37) AddMessageNumber() *DocumentNumber5Choice {
	l.MessageNumber = new(DocumentNumber5Choice)
	return l.MessageNumber
}

func (l *Linkages37) AddReference() *References41Choice {
	l.Reference = new(References41Choice)
	return l.Reference
}

func (l *Linkages37) AddLinkedQuantity() *PairedOrTurnedQuantity3Choice {
	l.LinkedQuantity = new(PairedOrTurnedQuantity3Choice)
	return l.LinkedQuantity
}

func (l *Linkages37) AddReferenceOwner() *PartyIdentification92Choice {
	l.ReferenceOwner = new(PartyIdentification92Choice)
	return l.ReferenceOwner
}

// Information related to a linked transaction.
type Linkages38 struct {

	// When the transaction is to be executed relative to a linked transaction.
	ProcessingPosition *ProcessingPosition7Choice `xml:"PrcgPos,omitempty"`

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber5Choice `xml:"MsgNb,omitempty"`

	// Reference to the linked transaction.
	Reference *References47Choice `xml:"Ref"`

	// Quantity of financial instruments of the linked transaction to be paired-off or turned.
	LinkedQuantity *PairedOrTurnedQuantity3Choice `xml:"LkdQty,omitempty"`
}

func (l *Linkages38) AddProcessingPosition() *ProcessingPosition7Choice {
	l.ProcessingPosition = new(ProcessingPosition7Choice)
	return l.ProcessingPosition
}

func (l *Linkages38) AddMessageNumber() *DocumentNumber5Choice {
	l.MessageNumber = new(DocumentNumber5Choice)
	return l.MessageNumber
}

func (l *Linkages38) AddReference() *References47Choice {
	l.Reference = new(References47Choice)
	return l.Reference
}

func (l *Linkages38) AddLinkedQuantity() *PairedOrTurnedQuantity3Choice {
	l.LinkedQuantity = new(PairedOrTurnedQuantity3Choice)
	return l.LinkedQuantity
}

// Information related to a linked transaction.
type Linkages39 struct {

	// When the transaction is to be executed relative to a linked transaction.
	ProcessingPosition *ProcessingPosition8Choice `xml:"PrcgPos,omitempty"`

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber5Choice `xml:"MsgNb,omitempty"`

	// Reference to the linked transaction.
	Reference *References46Choice `xml:"Ref"`

	// Party that generates the reference.
	ReferenceOwner *PartyIdentification92Choice `xml:"RefOwnr,omitempty"`
}

func (l *Linkages39) AddProcessingPosition() *ProcessingPosition8Choice {
	l.ProcessingPosition = new(ProcessingPosition8Choice)
	return l.ProcessingPosition
}

func (l *Linkages39) AddMessageNumber() *DocumentNumber5Choice {
	l.MessageNumber = new(DocumentNumber5Choice)
	return l.MessageNumber
}

func (l *Linkages39) AddReference() *References46Choice {
	l.Reference = new(References46Choice)
	return l.Reference
}

func (l *Linkages39) AddReferenceOwner() *PartyIdentification92Choice {
	l.ReferenceOwner = new(PartyIdentification92Choice)
	return l.ReferenceOwner
}

// Information related to a linked transaction.
type Linkages40 struct {

	// Reference to the linked transaction.
	Reference *References47Choice `xml:"Ref"`
}

func (l *Linkages40) AddReference() *References47Choice {
	l.Reference = new(References47Choice)
	return l.Reference
}

// Information related to a linked transaction.
type Linkages41 struct {

	// When the transaction is to be executed relative to a linked transaction - for information only.
	ProcessingPosition *ProcessingPosition9Choice `xml:"PrcgPos,omitempty"`

	// Unambiguous identification of a securities settlement transaction as known by the account owner (or instructing party acting on its behalf).
	SecuritiesSettlementTransactionIdentification *Max35Text `xml:"SctiesSttlmTxId"`
}

func (l *Linkages41) AddProcessingPosition() *ProcessingPosition9Choice {
	l.ProcessingPosition = new(ProcessingPosition9Choice)
	return l.ProcessingPosition
}

func (l *Linkages41) SetSecuritiesSettlementTransactionIdentification(value string) {
	l.SecuritiesSettlementTransactionIdentification = (*Max35Text)(&value)
}

// Information related to a linked transaction.
type Linkages42 struct {

	// When the transaction is to be executed relative to a linked transaction.
	ProcessingPosition *ProcessingPosition10Choice `xml:"PrcgPos,omitempty"`

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber6Choice `xml:"MsgNb,omitempty"`

	// Reference to the linked transaction.
	Reference *References50Choice `xml:"Ref"`

	// Party that generates the reference.
	ReferenceOwner *PartyIdentification103Choice `xml:"RefOwnr,omitempty"`
}

func (l *Linkages42) AddProcessingPosition() *ProcessingPosition10Choice {
	l.ProcessingPosition = new(ProcessingPosition10Choice)
	return l.ProcessingPosition
}

func (l *Linkages42) AddMessageNumber() *DocumentNumber6Choice {
	l.MessageNumber = new(DocumentNumber6Choice)
	return l.MessageNumber
}

func (l *Linkages42) AddReference() *References50Choice {
	l.Reference = new(References50Choice)
	return l.Reference
}

func (l *Linkages42) AddReferenceOwner() *PartyIdentification103Choice {
	l.ReferenceOwner = new(PartyIdentification103Choice)
	return l.ReferenceOwner
}

// Information related to a linked transaction.
type Linkages43 struct {

	// When the transaction is to be executed relative to a linked transaction.
	ProcessingPosition *ProcessingPosition10Choice `xml:"PrcgPos,omitempty"`

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber6Choice `xml:"MsgNb,omitempty"`

	// Reference to the linked transaction.
	Reference *References50Choice `xml:"Ref"`

	// Quantity of financial instruments of the linked transaction to be paired-off or turned.
	LinkedQuantity *PairedOrTurnedQuantity4Choice `xml:"LkdQty,omitempty"`

	// Party that generates the reference.
	ReferenceOwner *PartyIdentification103Choice `xml:"RefOwnr,omitempty"`
}

func (l *Linkages43) AddProcessingPosition() *ProcessingPosition10Choice {
	l.ProcessingPosition = new(ProcessingPosition10Choice)
	return l.ProcessingPosition
}

func (l *Linkages43) AddMessageNumber() *DocumentNumber6Choice {
	l.MessageNumber = new(DocumentNumber6Choice)
	return l.MessageNumber
}

func (l *Linkages43) AddReference() *References50Choice {
	l.Reference = new(References50Choice)
	return l.Reference
}

func (l *Linkages43) AddLinkedQuantity() *PairedOrTurnedQuantity4Choice {
	l.LinkedQuantity = new(PairedOrTurnedQuantity4Choice)
	return l.LinkedQuantity
}

func (l *Linkages43) AddReferenceOwner() *PartyIdentification103Choice {
	l.ReferenceOwner = new(PartyIdentification103Choice)
	return l.ReferenceOwner
}

// Information related to a linked transaction.
type Linkages44 struct {

	// When the transaction is to be executed relative to a linked transaction.
	ProcessingPosition *ProcessingPosition18Choice `xml:"PrcgPos,omitempty"`

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber16Choice `xml:"MsgNb,omitempty"`

	// Reference to the linked transaction.
	Reference *References54Choice `xml:"Ref"`

	// Party that generates the reference.
	ReferenceOwner *PartyIdentification103Choice `xml:"RefOwnr,omitempty"`
}

func (l *Linkages44) AddProcessingPosition() *ProcessingPosition18Choice {
	l.ProcessingPosition = new(ProcessingPosition18Choice)
	return l.ProcessingPosition
}

func (l *Linkages44) AddMessageNumber() *DocumentNumber16Choice {
	l.MessageNumber = new(DocumentNumber16Choice)
	return l.MessageNumber
}

func (l *Linkages44) AddReference() *References54Choice {
	l.Reference = new(References54Choice)
	return l.Reference
}

func (l *Linkages44) AddReferenceOwner() *PartyIdentification103Choice {
	l.ReferenceOwner = new(PartyIdentification103Choice)
	return l.ReferenceOwner
}

// Information related to a linked transaction.
type Linkages48 struct {

	// When the transaction is to be executed relative to a linked transaction.
	ProcessingPosition *ProcessingPosition10Choice `xml:"PrcgPos,omitempty"`

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber6Choice `xml:"MsgNb,omitempty"`

	// Reference to the linked transaction.
	Reference *References58Choice `xml:"Ref"`

	// Quantity of financial instruments of the linked transaction to be paired-off or turned.
	LinkedQuantity *PairedOrTurnedQuantity4Choice `xml:"LkdQty,omitempty"`
}

func (l *Linkages48) AddProcessingPosition() *ProcessingPosition10Choice {
	l.ProcessingPosition = new(ProcessingPosition10Choice)
	return l.ProcessingPosition
}

func (l *Linkages48) AddMessageNumber() *DocumentNumber6Choice {
	l.MessageNumber = new(DocumentNumber6Choice)
	return l.MessageNumber
}

func (l *Linkages48) AddReference() *References58Choice {
	l.Reference = new(References58Choice)
	return l.Reference
}

func (l *Linkages48) AddLinkedQuantity() *PairedOrTurnedQuantity4Choice {
	l.LinkedQuantity = new(PairedOrTurnedQuantity4Choice)
	return l.LinkedQuantity
}

// Information related to a linked transaction.
type Linkages49 struct {

	// Reference to the linked transaction.
	Reference *References58Choice `xml:"Ref"`
}

func (l *Linkages49) AddReference() *References58Choice {
	l.Reference = new(References58Choice)
	return l.Reference
}

// Information related to a linked transaction.
type Linkages50 struct {

	// When the transaction is to be executed relative to a linked transaction - for information only.
	ProcessingPosition *ProcessingPosition23Choice `xml:"PrcgPos,omitempty"`

	// Unambiguous identification of a securities settlement transaction as known by the account owner (or instructing party acting on its behalf).
	SecuritiesSettlementTransactionIdentification *RestrictedFINMax16Text `xml:"SctiesSttlmTxId"`
}

func (l *Linkages50) AddProcessingPosition() *ProcessingPosition23Choice {
	l.ProcessingPosition = new(ProcessingPosition23Choice)
	return l.ProcessingPosition
}

func (l *Linkages50) SetSecuritiesSettlementTransactionIdentification(value string) {
	l.SecuritiesSettlementTransactionIdentification = (*RestrictedFINMax16Text)(&value)
}

// Information related to a linked transaction.
type Linkages7 struct {

	// When the transaction is to be executed relative to a linked transaction.
	ProcessingPosition *ProcessingPosition1Choice `xml:"PrcgPos,omitempty"`

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber1Choice `xml:"MsgNb,omitempty"`

	// Reference to the linked transaction.
	Reference *References14Choice `xml:"Ref"`

	// Quantity of financial instruments of the linked transaction to be paired-off or turned.
	LinkedQuantity *PairedOrTurnedQuantity1Choice `xml:"LkdQty,omitempty"`

	// Party that generates the reference.
	ReferenceOwner *PartyIdentification36Choice `xml:"RefOwnr,omitempty"`
}

func (l *Linkages7) AddProcessingPosition() *ProcessingPosition1Choice {
	l.ProcessingPosition = new(ProcessingPosition1Choice)
	return l.ProcessingPosition
}

func (l *Linkages7) AddMessageNumber() *DocumentNumber1Choice {
	l.MessageNumber = new(DocumentNumber1Choice)
	return l.MessageNumber
}

func (l *Linkages7) AddReference() *References14Choice {
	l.Reference = new(References14Choice)
	return l.Reference
}

func (l *Linkages7) AddLinkedQuantity() *PairedOrTurnedQuantity1Choice {
	l.LinkedQuantity = new(PairedOrTurnedQuantity1Choice)
	return l.LinkedQuantity
}

func (l *Linkages7) AddReferenceOwner() *PartyIdentification36Choice {
	l.ReferenceOwner = new(PartyIdentification36Choice)
	return l.ReferenceOwner
}

// Information related to a linked transaction.
type Linkages8 struct {

	// When the transaction is to be executed relative to a linked transaction.
	ProcessingPosition *ProcessingPosition1Choice `xml:"PrcgPos,omitempty"`

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber1Choice `xml:"MsgNb,omitempty"`

	// Reference to the linked transaction.
	Reference *References14Choice `xml:"Ref"`

	// Party that generates the reference.
	ReferenceOwner *PartyIdentification36Choice `xml:"RefOwnr,omitempty"`
}

func (l *Linkages8) AddProcessingPosition() *ProcessingPosition1Choice {
	l.ProcessingPosition = new(ProcessingPosition1Choice)
	return l.ProcessingPosition
}

func (l *Linkages8) AddMessageNumber() *DocumentNumber1Choice {
	l.MessageNumber = new(DocumentNumber1Choice)
	return l.MessageNumber
}

func (l *Linkages8) AddReference() *References14Choice {
	l.Reference = new(References14Choice)
	return l.Reference
}

func (l *Linkages8) AddReferenceOwner() *PartyIdentification36Choice {
	l.ReferenceOwner = new(PartyIdentification36Choice)
	return l.ReferenceOwner
}

// Information related to a linked transaction.
type Linkages9 struct {

	// When the transaction is to be executed relative to a linked transaction.
	ProcessingPosition *ProcessingPosition1Choice `xml:"PrcgPos,omitempty"`

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber1Choice `xml:"MsgNb,omitempty"`

	// Reference to the linked transaction.
	Reference *References14Choice `xml:"Ref"`

	// Quantity of financial instruments of the linked transaction to be paired-off or turned.
	LinkedQuantity *PairedOrTurnedQuantity1Choice `xml:"LkdQty,omitempty"`

	// Party that generates the reference.
	ReferenceOwner *PartyIdentification36Choice `xml:"RefOwnr,omitempty"`
}

func (l *Linkages9) AddProcessingPosition() *ProcessingPosition1Choice {
	l.ProcessingPosition = new(ProcessingPosition1Choice)
	return l.ProcessingPosition
}

func (l *Linkages9) AddMessageNumber() *DocumentNumber1Choice {
	l.MessageNumber = new(DocumentNumber1Choice)
	return l.MessageNumber
}

func (l *Linkages9) AddReference() *References14Choice {
	l.Reference = new(References14Choice)
	return l.Reference
}

func (l *Linkages9) AddLinkedQuantity() *PairedOrTurnedQuantity1Choice {
	l.LinkedQuantity = new(PairedOrTurnedQuantity1Choice)
	return l.LinkedQuantity
}

func (l *Linkages9) AddReferenceOwner() *PartyIdentification36Choice {
	l.ReferenceOwner = new(PartyIdentification36Choice)
	return l.ReferenceOwner
}
