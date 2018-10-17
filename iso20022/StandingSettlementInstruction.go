package iso20022

// Details of the standing settlement instruction to be applied.
type StandingSettlementInstruction1 struct {

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	SettlementStandingInstructionDatabase *SettlementStandingInstructionDatabase1Choice `xml:"SttlmStgInstrDB"`

	// Identification of the buyer or seller in a standing settlement instruction enabling to derive the Standing Settlement Instruction.
	Counterparty *Counterparty1Choice `xml:"CtrPty"`

	// Vendor of the Settlement Standing Instruction database requested to be consulted.
	Vendor *PartyIdentification10Choice `xml:"Vndr,omitempty"`

	// Delivering parties, other than the seller, needed for deriving the standing settlement instruction (for example, depository) or provided for information purposes (for example, instructing party settlement chain).
	OtherDeliveringSettlementParties *SettlementParties5 `xml:"OthrDlvrgSttlmPties,omitempty"`

	// Receiving parties, other than the buyer, needed for deriving the standing settlement instruction (for example, depository) or provided for information purposes (for example, instructing party settlement chain).
	OtherReceivingSettlementParties *SettlementParties5 `xml:"OthrRcvgSttlmPties,omitempty"`
}

func (s *StandingSettlementInstruction1) AddSettlementStandingInstructionDatabase() *SettlementStandingInstructionDatabase1Choice {
	s.SettlementStandingInstructionDatabase = new(SettlementStandingInstructionDatabase1Choice)
	return s.SettlementStandingInstructionDatabase
}

func (s *StandingSettlementInstruction1) AddCounterparty() *Counterparty1Choice {
	s.Counterparty = new(Counterparty1Choice)
	return s.Counterparty
}

func (s *StandingSettlementInstruction1) AddVendor() *PartyIdentification10Choice {
	s.Vendor = new(PartyIdentification10Choice)
	return s.Vendor
}

func (s *StandingSettlementInstruction1) AddOtherDeliveringSettlementParties() *SettlementParties5 {
	s.OtherDeliveringSettlementParties = new(SettlementParties5)
	return s.OtherDeliveringSettlementParties
}

func (s *StandingSettlementInstruction1) AddOtherReceivingSettlementParties() *SettlementParties5 {
	s.OtherReceivingSettlementParties = new(SettlementParties5)
	return s.OtherReceivingSettlementParties
}

// Details of the standing settlement instruction to be applied.
type StandingSettlementInstruction11 struct {

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	SettlementStandingInstructionDatabase *SettlementStandingInstructionDatabase4Choice `xml:"SttlmStgInstrDB"`

	// Identification of the buyer or seller in a standing settlement instruction enabling to derive the Standing Settlement Instruction.
	Counterparty *Counterparty9Choice `xml:"CtrPty"`

	// Vendor of the Settlement Standing Instruction database requested to be consulted.
	Vendor *PartyIdentification100 `xml:"Vndr,omitempty"`

	// Delivering parties, other than the seller, needed for deriving the standing settlement instruction (for example, depository) or provided for information purposes (for example, instructing party settlement chain).
	OtherDeliveringSettlementParties *SettlementParties36 `xml:"OthrDlvrgSttlmPties,omitempty"`

	// Receiving parties, other than the buyer, needed for deriving the standing settlement instruction (for example, depository) or provided for information purposes (for example, instructing party settlement chain).
	OtherReceivingSettlementParties *SettlementParties36 `xml:"OthrRcvgSttlmPties,omitempty"`
}

func (s *StandingSettlementInstruction11) AddSettlementStandingInstructionDatabase() *SettlementStandingInstructionDatabase4Choice {
	s.SettlementStandingInstructionDatabase = new(SettlementStandingInstructionDatabase4Choice)
	return s.SettlementStandingInstructionDatabase
}

func (s *StandingSettlementInstruction11) AddCounterparty() *Counterparty9Choice {
	s.Counterparty = new(Counterparty9Choice)
	return s.Counterparty
}

func (s *StandingSettlementInstruction11) AddVendor() *PartyIdentification100 {
	s.Vendor = new(PartyIdentification100)
	return s.Vendor
}

func (s *StandingSettlementInstruction11) AddOtherDeliveringSettlementParties() *SettlementParties36 {
	s.OtherDeliveringSettlementParties = new(SettlementParties36)
	return s.OtherDeliveringSettlementParties
}

func (s *StandingSettlementInstruction11) AddOtherReceivingSettlementParties() *SettlementParties36 {
	s.OtherReceivingSettlementParties = new(SettlementParties36)
	return s.OtherReceivingSettlementParties
}

// Details of the standing settlement instruction to be applied.
type StandingSettlementInstruction12 struct {

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	SettlementStandingInstructionDatabase *SettlementStandingInstructionDatabase5Choice `xml:"SttlmStgInstrDB"`

	// Identification of the buyer or seller in a standing settlement instruction enabling to derive the Standing Settlement Instruction.
	Counterparty *Counterparty10Choice `xml:"CtrPty"`

	// Vendor of the Settlement Standing Instruction database requested to be consulted.
	Vendor *PartyIdentification111 `xml:"Vndr,omitempty"`

	// Delivering parties, other than the seller, needed for deriving the standing settlement instruction (for example, depository) or provided for information purposes (for example, instructing party settlement chain).
	OtherDeliveringSettlementParties *SettlementParties44 `xml:"OthrDlvrgSttlmPties,omitempty"`

	// Receiving parties, other than the buyer, needed for deriving the standing settlement instruction (for example, depository) or provided for information purposes (for example, instructing party settlement chain).
	OtherReceivingSettlementParties *SettlementParties44 `xml:"OthrRcvgSttlmPties,omitempty"`
}

func (s *StandingSettlementInstruction12) AddSettlementStandingInstructionDatabase() *SettlementStandingInstructionDatabase5Choice {
	s.SettlementStandingInstructionDatabase = new(SettlementStandingInstructionDatabase5Choice)
	return s.SettlementStandingInstructionDatabase
}

func (s *StandingSettlementInstruction12) AddCounterparty() *Counterparty10Choice {
	s.Counterparty = new(Counterparty10Choice)
	return s.Counterparty
}

func (s *StandingSettlementInstruction12) AddVendor() *PartyIdentification111 {
	s.Vendor = new(PartyIdentification111)
	return s.Vendor
}

func (s *StandingSettlementInstruction12) AddOtherDeliveringSettlementParties() *SettlementParties44 {
	s.OtherDeliveringSettlementParties = new(SettlementParties44)
	return s.OtherDeliveringSettlementParties
}

func (s *StandingSettlementInstruction12) AddOtherReceivingSettlementParties() *SettlementParties44 {
	s.OtherReceivingSettlementParties = new(SettlementParties44)
	return s.OtherReceivingSettlementParties
}

// Details of the standing settlement instruction to be applied.
type StandingSettlementInstruction3 struct {

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	SettlementStandingInstructionDatabase *SettlementStandingInstructionDatabase1Choice `xml:"SttlmStgInstrDB"`

	// Identification of the buyer or seller in a standing settlement instruction enabling to derive the Standing Settlement Instruction.
	Counterparty *Counterparty3Choice `xml:"CtrPty"`

	// Vendor of the Settlement Standing Instruction database requested to be consulted.
	Vendor *PartyIdentification49Choice `xml:"Vndr,omitempty"`

	// Delivering parties, other than the seller, needed for deriving the standing settlement instruction (for example, depository) or provided for information purposes (for example, instructing party settlement chain).
	OtherDeliveringSettlementParties *SettlementParties10 `xml:"OthrDlvrgSttlmPties,omitempty"`

	// Receiving parties, other than the buyer, needed for deriving the standing settlement instruction (for example, depository) or provided for information purposes (for example, instructing party settlement chain).
	OtherReceivingSettlementParties *SettlementParties10 `xml:"OthrRcvgSttlmPties,omitempty"`
}

func (s *StandingSettlementInstruction3) AddSettlementStandingInstructionDatabase() *SettlementStandingInstructionDatabase1Choice {
	s.SettlementStandingInstructionDatabase = new(SettlementStandingInstructionDatabase1Choice)
	return s.SettlementStandingInstructionDatabase
}

func (s *StandingSettlementInstruction3) AddCounterparty() *Counterparty3Choice {
	s.Counterparty = new(Counterparty3Choice)
	return s.Counterparty
}

func (s *StandingSettlementInstruction3) AddVendor() *PartyIdentification49Choice {
	s.Vendor = new(PartyIdentification49Choice)
	return s.Vendor
}

func (s *StandingSettlementInstruction3) AddOtherDeliveringSettlementParties() *SettlementParties10 {
	s.OtherDeliveringSettlementParties = new(SettlementParties10)
	return s.OtherDeliveringSettlementParties
}

func (s *StandingSettlementInstruction3) AddOtherReceivingSettlementParties() *SettlementParties10 {
	s.OtherReceivingSettlementParties = new(SettlementParties10)
	return s.OtherReceivingSettlementParties
}

// Details of the standing settlement instruction to be applied.
type StandingSettlementInstruction4 struct {

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	SettlementStandingInstructionDatabase *SettlementStandingInstructionDatabase1Choice `xml:"SttlmStgInstrDB"`

	// Identification of the buyer or seller in a standing settlement instruction enabling to derive the Standing Settlement Instruction.
	Counterparty *Counterparty4Choice `xml:"CtrPty"`

	// Vendor of the Settlement Standing Instruction database requested to be consulted.
	Vendor *PartyIdentification43Choice `xml:"Vndr,omitempty"`

	// Delivering parties, other than the seller, needed for deriving the standing settlement instruction (for example, depository) or provided for information purposes (for example, instructing party settlement chain).
	OtherDeliveringSettlementParties *SettlementParties11 `xml:"OthrDlvrgSttlmPties,omitempty"`

	// Receiving parties, other than the buyer, needed for deriving the standing settlement instruction (for example, depository) or provided for information purposes (for example, instructing party settlement chain).
	OtherReceivingSettlementParties *SettlementParties11 `xml:"OthrRcvgSttlmPties,omitempty"`
}

func (s *StandingSettlementInstruction4) AddSettlementStandingInstructionDatabase() *SettlementStandingInstructionDatabase1Choice {
	s.SettlementStandingInstructionDatabase = new(SettlementStandingInstructionDatabase1Choice)
	return s.SettlementStandingInstructionDatabase
}

func (s *StandingSettlementInstruction4) AddCounterparty() *Counterparty4Choice {
	s.Counterparty = new(Counterparty4Choice)
	return s.Counterparty
}

func (s *StandingSettlementInstruction4) AddVendor() *PartyIdentification43Choice {
	s.Vendor = new(PartyIdentification43Choice)
	return s.Vendor
}

func (s *StandingSettlementInstruction4) AddOtherDeliveringSettlementParties() *SettlementParties11 {
	s.OtherDeliveringSettlementParties = new(SettlementParties11)
	return s.OtherDeliveringSettlementParties
}

func (s *StandingSettlementInstruction4) AddOtherReceivingSettlementParties() *SettlementParties11 {
	s.OtherReceivingSettlementParties = new(SettlementParties11)
	return s.OtherReceivingSettlementParties
}

// Details of the standing settlement instruction to be applied.
type StandingSettlementInstruction5 struct {

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	SettlementStandingInstructionDatabase *SettlementStandingInstructionDatabase1Choice `xml:"SttlmStgInstrDB"`

	// Identification of the buyer or seller in a standing settlement instruction enabling to derive the Standing Settlement Instruction.
	Counterparty *Counterparty5Choice `xml:"CtrPty"`

	// Vendor of the Settlement Standing Instruction database requested to be consulted.
	Vendor *PartyIdentification45Choice `xml:"Vndr,omitempty"`

	// Delivering parties, other than the seller, needed for deriving the standing settlement instruction (for example, depository) or provided for information purposes (for example, instructing party settlement chain).
	OtherDeliveringSettlementParties *SettlementParties14 `xml:"OthrDlvrgSttlmPties,omitempty"`

	// Receiving parties, other than the buyer, needed for deriving the standing settlement instruction (for example, depository) or provided for information purposes (for example, instructing party settlement chain).
	OtherReceivingSettlementParties *SettlementParties14 `xml:"OthrRcvgSttlmPties,omitempty"`
}

func (s *StandingSettlementInstruction5) AddSettlementStandingInstructionDatabase() *SettlementStandingInstructionDatabase1Choice {
	s.SettlementStandingInstructionDatabase = new(SettlementStandingInstructionDatabase1Choice)
	return s.SettlementStandingInstructionDatabase
}

func (s *StandingSettlementInstruction5) AddCounterparty() *Counterparty5Choice {
	s.Counterparty = new(Counterparty5Choice)
	return s.Counterparty
}

func (s *StandingSettlementInstruction5) AddVendor() *PartyIdentification45Choice {
	s.Vendor = new(PartyIdentification45Choice)
	return s.Vendor
}

func (s *StandingSettlementInstruction5) AddOtherDeliveringSettlementParties() *SettlementParties14 {
	s.OtherDeliveringSettlementParties = new(SettlementParties14)
	return s.OtherDeliveringSettlementParties
}

func (s *StandingSettlementInstruction5) AddOtherReceivingSettlementParties() *SettlementParties14 {
	s.OtherReceivingSettlementParties = new(SettlementParties14)
	return s.OtherReceivingSettlementParties
}

// Details of the standing settlement instruction to be applied.
type StandingSettlementInstruction9 struct {

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	SettlementStandingInstructionDatabase *SettlementStandingInstructionDatabase3Choice `xml:"SttlmStgInstrDB"`

	// Vendor of the Settlement Standing Instruction database requested to be consulted.
	Vendor *PartyIdentification32Choice `xml:"Vndr,omitempty"`

	// Delivering parties, other than the seller, needed for deriving the standing settlement instruction (for example, depository) or provided for information purposes (for example, instructing party settlement chain).
	OtherDeliveringSettlementParties *SettlementParties23 `xml:"OthrDlvrgSttlmPties,omitempty"`

	// Receiving parties, other than the buyer, needed for deriving the standing settlement instruction (for example, depository) or provided for information purposes (for example, instructing party settlement chain).
	OtherReceivingSettlementParties *SettlementParties23 `xml:"OthrRcvgSttlmPties,omitempty"`
}

func (s *StandingSettlementInstruction9) AddSettlementStandingInstructionDatabase() *SettlementStandingInstructionDatabase3Choice {
	s.SettlementStandingInstructionDatabase = new(SettlementStandingInstructionDatabase3Choice)
	return s.SettlementStandingInstructionDatabase
}

func (s *StandingSettlementInstruction9) AddVendor() *PartyIdentification32Choice {
	s.Vendor = new(PartyIdentification32Choice)
	return s.Vendor
}

func (s *StandingSettlementInstruction9) AddOtherDeliveringSettlementParties() *SettlementParties23 {
	s.OtherDeliveringSettlementParties = new(SettlementParties23)
	return s.OtherDeliveringSettlementParties
}

func (s *StandingSettlementInstruction9) AddOtherReceivingSettlementParties() *SettlementParties23 {
	s.OtherReceivingSettlementParties = new(SettlementParties23)
	return s.OtherReceivingSettlementParties
}
