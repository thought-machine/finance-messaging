package iso20022

// Further information required for the settlement the transaction.
type SettlementInformation1 struct {

	// Method used to settle the (batch of) payment instructions.
	SettlementMethod *SettlementMethod1Code `xml:"SttlmMtd"`

	// A specific purpose account used to post debit and credit entries as a result of the transaction.
	SettlementAccount *CashAccount7 `xml:"SttlmAcct,omitempty"`

	// Specification of a pre-agreed offering between clearing agents or the channel through which the payment instruction is processed.
	ClearingSystem *ClearingSystemIdentification1Choice `xml:"ClrSys,omitempty"`

	// Agent through which the instructing agent will reimburse the instructed agent.
	//
	// Usage: If InstructingAgent and InstructedAgent have the same reimbursement agent, then only InstructingReimbursementAgent must be used.
	InstructingReimbursementAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstgRmbrsmntAgt,omitempty"`

	// Unambiguous identification of the account of the instructing reimbursement agent account at its servicing agent in the payment chain.
	InstructingReimbursementAgentAccount *CashAccount7 `xml:"InstgRmbrsmntAgtAcct,omitempty"`

	// Agent at which the instructed agent will be reimbursed.
	//
	// Usage: If InstructedReimbursementAgent contains a branch of the InstructedAgent, then the instructed agent will claim reimbursement from that branch/will be paid by that branch.
	//
	// Usage: If InstructingAgent and InstructedAgent have the same reimbursement agent, then only InstructingReimbursementAgent must be used.
	InstructedReimbursementAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstdRmbrsmntAgt,omitempty"`

	// Unambiguous identification of the account of the instructed reimbursement agent account at its servicing agent in the payment chain.
	InstructedReimbursementAgentAccount *CashAccount7 `xml:"InstdRmbrsmntAgtAcct,omitempty"`

	// Instructed agent's branch where the amount of money will be made available when different from the instructed reimbursement agent.
	ThirdReimbursementAgent *BranchAndFinancialInstitutionIdentification3 `xml:"ThrdRmbrsmntAgt,omitempty"`

	// Unambiguous identification of the account of the third reimbursement agent account at its servicing agent in the payment chain.
	ThirdReimbursementAgentAccount *CashAccount7 `xml:"ThrdRmbrsmntAgtAcct,omitempty"`
}

func (s *SettlementInformation1) SetSettlementMethod(value string) {
	s.SettlementMethod = (*SettlementMethod1Code)(&value)
}

func (s *SettlementInformation1) AddSettlementAccount() *CashAccount7 {
	s.SettlementAccount = new(CashAccount7)
	return s.SettlementAccount
}

func (s *SettlementInformation1) AddClearingSystem() *ClearingSystemIdentification1Choice {
	s.ClearingSystem = new(ClearingSystemIdentification1Choice)
	return s.ClearingSystem
}

func (s *SettlementInformation1) AddInstructingReimbursementAgent() *BranchAndFinancialInstitutionIdentification3 {
	s.InstructingReimbursementAgent = new(BranchAndFinancialInstitutionIdentification3)
	return s.InstructingReimbursementAgent
}

func (s *SettlementInformation1) AddInstructingReimbursementAgentAccount() *CashAccount7 {
	s.InstructingReimbursementAgentAccount = new(CashAccount7)
	return s.InstructingReimbursementAgentAccount
}

func (s *SettlementInformation1) AddInstructedReimbursementAgent() *BranchAndFinancialInstitutionIdentification3 {
	s.InstructedReimbursementAgent = new(BranchAndFinancialInstitutionIdentification3)
	return s.InstructedReimbursementAgent
}

func (s *SettlementInformation1) AddInstructedReimbursementAgentAccount() *CashAccount7 {
	s.InstructedReimbursementAgentAccount = new(CashAccount7)
	return s.InstructedReimbursementAgentAccount
}

func (s *SettlementInformation1) AddThirdReimbursementAgent() *BranchAndFinancialInstitutionIdentification3 {
	s.ThirdReimbursementAgent = new(BranchAndFinancialInstitutionIdentification3)
	return s.ThirdReimbursementAgent
}

func (s *SettlementInformation1) AddThirdReimbursementAgentAccount() *CashAccount7 {
	s.ThirdReimbursementAgentAccount = new(CashAccount7)
	return s.ThirdReimbursementAgentAccount
}

// Set of elements used to provide information on the settlement of the instruction.
type SettlementInformation13 struct {

	// Method used to settle the (batch of) payment instructions.
	SettlementMethod *SettlementMethod1Code `xml:"SttlmMtd"`

	// A specific purpose account used to post debit and credit entries as a result of the transaction.
	SettlementAccount *CashAccount16 `xml:"SttlmAcct,omitempty"`

	// Specification of a pre-agreed offering between clearing agents or the channel through which the payment instruction is processed.
	ClearingSystem *ClearingSystemIdentification3Choice `xml:"ClrSys,omitempty"`

	// Agent through which the instructing agent will reimburse the instructed agent.
	//
	// Usage: If InstructingAgent and InstructedAgent have the same reimbursement agent, then only InstructingReimbursementAgent must be used.
	InstructingReimbursementAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstgRmbrsmntAgt,omitempty"`

	// Unambiguous identification of the account of the instructing reimbursement agent account at its servicing agent in the payment chain.
	InstructingReimbursementAgentAccount *CashAccount16 `xml:"InstgRmbrsmntAgtAcct,omitempty"`

	// Agent at which the instructed agent will be reimbursed.
	// Usage: If InstructedReimbursementAgent contains a branch of the InstructedAgent, then the party in InstructedAgent will claim reimbursement from that branch/will be paid by that branch.
	// Usage: If InstructingAgent and InstructedAgent have the same reimbursement agent, then only InstructingReimbursementAgent must be used.
	InstructedReimbursementAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstdRmbrsmntAgt,omitempty"`

	// Unambiguous identification of the account of the instructed reimbursement agent account at its servicing agent in the payment chain.
	InstructedReimbursementAgentAccount *CashAccount16 `xml:"InstdRmbrsmntAgtAcct,omitempty"`

	// Agent at which the instructed agent will be reimbursed.
	// Usage: If ThirdReimbursementAgent contains a branch of the InstructedAgent, then the party in InstructedAgent will claim reimbursement from that branch/will be paid by that branch.
	ThirdReimbursementAgent *BranchAndFinancialInstitutionIdentification4 `xml:"ThrdRmbrsmntAgt,omitempty"`

	// Unambiguous identification of the account of the third reimbursement agent account at its servicing agent in the payment chain.
	ThirdReimbursementAgentAccount *CashAccount16 `xml:"ThrdRmbrsmntAgtAcct,omitempty"`
}

func (s *SettlementInformation13) SetSettlementMethod(value string) {
	s.SettlementMethod = (*SettlementMethod1Code)(&value)
}

func (s *SettlementInformation13) AddSettlementAccount() *CashAccount16 {
	s.SettlementAccount = new(CashAccount16)
	return s.SettlementAccount
}

func (s *SettlementInformation13) AddClearingSystem() *ClearingSystemIdentification3Choice {
	s.ClearingSystem = new(ClearingSystemIdentification3Choice)
	return s.ClearingSystem
}

func (s *SettlementInformation13) AddInstructingReimbursementAgent() *BranchAndFinancialInstitutionIdentification4 {
	s.InstructingReimbursementAgent = new(BranchAndFinancialInstitutionIdentification4)
	return s.InstructingReimbursementAgent
}

func (s *SettlementInformation13) AddInstructingReimbursementAgentAccount() *CashAccount16 {
	s.InstructingReimbursementAgentAccount = new(CashAccount16)
	return s.InstructingReimbursementAgentAccount
}

func (s *SettlementInformation13) AddInstructedReimbursementAgent() *BranchAndFinancialInstitutionIdentification4 {
	s.InstructedReimbursementAgent = new(BranchAndFinancialInstitutionIdentification4)
	return s.InstructedReimbursementAgent
}

func (s *SettlementInformation13) AddInstructedReimbursementAgentAccount() *CashAccount16 {
	s.InstructedReimbursementAgentAccount = new(CashAccount16)
	return s.InstructedReimbursementAgentAccount
}

func (s *SettlementInformation13) AddThirdReimbursementAgent() *BranchAndFinancialInstitutionIdentification4 {
	s.ThirdReimbursementAgent = new(BranchAndFinancialInstitutionIdentification4)
	return s.ThirdReimbursementAgent
}

func (s *SettlementInformation13) AddThirdReimbursementAgentAccount() *CashAccount16 {
	s.ThirdReimbursementAgentAccount = new(CashAccount16)
	return s.ThirdReimbursementAgentAccount
}

// Set of elements used to provide information on the settlement of the instruction.
type SettlementInformation14 struct {

	// Method used to settle the (batch of) payment instructions.
	SettlementMethod *SettlementMethod2Code `xml:"SttlmMtd"`

	// A specific purpose account used to post debit and credit entries as a result of the transaction.
	SettlementAccount *CashAccount16 `xml:"SttlmAcct,omitempty"`

	// Specification of a pre-agreed offering between clearing agents or the channel through which the payment instruction is processed.
	ClearingSystem *ClearingSystemIdentification3Choice `xml:"ClrSys,omitempty"`
}

func (s *SettlementInformation14) SetSettlementMethod(value string) {
	s.SettlementMethod = (*SettlementMethod2Code)(&value)
}

func (s *SettlementInformation14) AddSettlementAccount() *CashAccount16 {
	s.SettlementAccount = new(CashAccount16)
	return s.SettlementAccount
}

func (s *SettlementInformation14) AddClearingSystem() *ClearingSystemIdentification3Choice {
	s.ClearingSystem = new(ClearingSystemIdentification3Choice)
	return s.ClearingSystem
}

// Set of elements used to provide information on the settlement of the instruction.
type SettlementInformation15 struct {

	// Agent through which the instructing agent will reimburse the instructed agent.
	// Usage: If InstructingAgent and InstructedAgent have the same reimbursement agent, then only InstructingReimbursementAgent must be used.
	InstructingReimbursementAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstgRmbrsmntAgt,omitempty"`

	// Unambiguous identification of the account of the instructing reimbursement agent account at its servicing agent in the payment chain.
	InstructingReimbursementAgentAccount *CashAccount16 `xml:"InstgRmbrsmntAgtAcct,omitempty"`

	// Agent at which the instructed agent will be reimbursed.
	// Usage: If InstructedReimbursementAgent contains a branch of the InstructedAgent, then the party in InstructedAgent will claim reimbursement from that branch/will be paid by that branch.
	// Usage: If InstructingAgent and InstructedAgent have the same reimbursement agent, then only InstructingReimbursementAgent must be used.
	InstructedReimbursementAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstdRmbrsmntAgt,omitempty"`

	// Unambiguous identification of the account of the instructed reimbursement agent account at its servicing agent in the payment chain.
	InstructedReimbursementAgentAccount *CashAccount16 `xml:"InstdRmbrsmntAgtAcct,omitempty"`
}

func (s *SettlementInformation15) AddInstructingReimbursementAgent() *BranchAndFinancialInstitutionIdentification4 {
	s.InstructingReimbursementAgent = new(BranchAndFinancialInstitutionIdentification4)
	return s.InstructingReimbursementAgent
}

func (s *SettlementInformation15) AddInstructingReimbursementAgentAccount() *CashAccount16 {
	s.InstructingReimbursementAgentAccount = new(CashAccount16)
	return s.InstructingReimbursementAgentAccount
}

func (s *SettlementInformation15) AddInstructedReimbursementAgent() *BranchAndFinancialInstitutionIdentification4 {
	s.InstructedReimbursementAgent = new(BranchAndFinancialInstitutionIdentification4)
	return s.InstructedReimbursementAgent
}

func (s *SettlementInformation15) AddInstructedReimbursementAgentAccount() *CashAccount16 {
	s.InstructedReimbursementAgentAccount = new(CashAccount16)
	return s.InstructedReimbursementAgentAccount
}

// Set of elements used to provide information on the settlement of the instruction.
type SettlementInformation16 struct {

	// Method used to settle the (batch of) payment instructions.
	SettlementMethod *SettlementMethod1Code `xml:"SttlmMtd"`

	// A specific purpose account used to post debit and credit entries as a result of the transaction.
	SettlementAccount *CashAccount16 `xml:"SttlmAcct,omitempty"`

	// Specification of a pre-agreed offering between clearing agents or the channel through which the payment instruction is processed.
	ClearingSystem *ClearingSystemIdentification3Choice `xml:"ClrSys,omitempty"`

	// Agent through which the instructing agent will reimburse the instructed agent.
	//
	// Usage: If InstructingAgent and InstructedAgent have the same reimbursement agent, then only InstructingReimbursementAgent must be used.
	InstructingReimbursementAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgRmbrsmntAgt,omitempty"`

	// Unambiguous identification of the account of the instructing reimbursement agent account at its servicing agent in the payment chain.
	InstructingReimbursementAgentAccount *CashAccount16 `xml:"InstgRmbrsmntAgtAcct,omitempty"`

	// Agent at which the instructed agent will be reimbursed.
	// Usage: If InstructedReimbursementAgent contains a branch of the InstructedAgent, then the party in InstructedAgent will claim reimbursement from that branch/will be paid by that branch.
	// Usage: If InstructingAgent and InstructedAgent have the same reimbursement agent, then only InstructingReimbursementAgent must be used.
	InstructedReimbursementAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdRmbrsmntAgt,omitempty"`

	// Unambiguous identification of the account of the instructed reimbursement agent account at its servicing agent in the payment chain.
	InstructedReimbursementAgentAccount *CashAccount16 `xml:"InstdRmbrsmntAgtAcct,omitempty"`

	// Agent at which the instructed agent will be reimbursed.
	// Usage: If ThirdReimbursementAgent contains a branch of the InstructedAgent, then the party in InstructedAgent will claim reimbursement from that branch/will be paid by that branch.
	ThirdReimbursementAgent *BranchAndFinancialInstitutionIdentification5 `xml:"ThrdRmbrsmntAgt,omitempty"`

	// Unambiguous identification of the account of the third reimbursement agent account at its servicing agent in the payment chain.
	ThirdReimbursementAgentAccount *CashAccount16 `xml:"ThrdRmbrsmntAgtAcct,omitempty"`
}

func (s *SettlementInformation16) SetSettlementMethod(value string) {
	s.SettlementMethod = (*SettlementMethod1Code)(&value)
}

func (s *SettlementInformation16) AddSettlementAccount() *CashAccount16 {
	s.SettlementAccount = new(CashAccount16)
	return s.SettlementAccount
}

func (s *SettlementInformation16) AddClearingSystem() *ClearingSystemIdentification3Choice {
	s.ClearingSystem = new(ClearingSystemIdentification3Choice)
	return s.ClearingSystem
}

func (s *SettlementInformation16) AddInstructingReimbursementAgent() *BranchAndFinancialInstitutionIdentification5 {
	s.InstructingReimbursementAgent = new(BranchAndFinancialInstitutionIdentification5)
	return s.InstructingReimbursementAgent
}

func (s *SettlementInformation16) AddInstructingReimbursementAgentAccount() *CashAccount16 {
	s.InstructingReimbursementAgentAccount = new(CashAccount16)
	return s.InstructingReimbursementAgentAccount
}

func (s *SettlementInformation16) AddInstructedReimbursementAgent() *BranchAndFinancialInstitutionIdentification5 {
	s.InstructedReimbursementAgent = new(BranchAndFinancialInstitutionIdentification5)
	return s.InstructedReimbursementAgent
}

func (s *SettlementInformation16) AddInstructedReimbursementAgentAccount() *CashAccount16 {
	s.InstructedReimbursementAgentAccount = new(CashAccount16)
	return s.InstructedReimbursementAgentAccount
}

func (s *SettlementInformation16) AddThirdReimbursementAgent() *BranchAndFinancialInstitutionIdentification5 {
	s.ThirdReimbursementAgent = new(BranchAndFinancialInstitutionIdentification5)
	return s.ThirdReimbursementAgent
}

func (s *SettlementInformation16) AddThirdReimbursementAgentAccount() *CashAccount16 {
	s.ThirdReimbursementAgentAccount = new(CashAccount16)
	return s.ThirdReimbursementAgentAccount
}

// Further information required for the settlement the transaction.
type SettlementInformation2 struct {

	// Method used to settle the (batch of) payment instructions.
	SettlementMethod *SettlementMethod2Code `xml:"SttlmMtd"`

	// A specific purpose account used to post debit and credit entries as a result of the transaction.
	SettlementAccount *CashAccount7 `xml:"SttlmAcct,omitempty"`

	// Specification of a pre-agreed offering between clearing agents or the channel through which the payment instruction is processed.
	ClearingSystem *ClearingSystemIdentification1Choice `xml:"ClrSys,omitempty"`
}

func (s *SettlementInformation2) SetSettlementMethod(value string) {
	s.SettlementMethod = (*SettlementMethod2Code)(&value)
}

func (s *SettlementInformation2) AddSettlementAccount() *CashAccount7 {
	s.SettlementAccount = new(CashAccount7)
	return s.SettlementAccount
}

func (s *SettlementInformation2) AddClearingSystem() *ClearingSystemIdentification1Choice {
	s.ClearingSystem = new(ClearingSystemIdentification1Choice)
	return s.ClearingSystem
}

// Further information required for the settlement the transaction.
type SettlementInformation3 struct {

	// Method used to settle the (batch of) payment instructions.
	SettlementMethod *SettlementMethod1Code `xml:"SttlmMtd"`

	// A specific purpose account used to post debit and credit entries as a result of the transaction.
	SettlementAccount *CashAccount7 `xml:"SttlmAcct,omitempty"`

	// Specification of a pre-agreed offering between clearing agents or the channel through which the payment instruction is processed.
	ClearingSystem *ClearingSystemIdentification1Choice `xml:"ClrSys,omitempty"`

	// Specifies the agent through which the instructing agent will reimburse the instructed agent.
	//
	// Usage: If the instructing and instructed agents have the same reimbursement agent, then only InstructingReimbursementAgent must be used.
	InstructingReimbursementAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstgRmbrsmntAgt,omitempty"`

	// Unambiguous identification of the account of the instructing reimbursement agent account at its servicing agent in the payment chain.
	InstructingReimbursementAgentAccount *CashAccount7 `xml:"InstgRmbrsmntAgtAcct,omitempty"`

	// Agent at which the instructed agent will be reimbursed.
	//
	// Usage: If InstructedReimbursementAgent contains a branch of the InstructedAgent, then the instructed agent will claim reimbursement from that branch/will be paid by that branch.
	//
	// Usage: If InstructingAgent and InstructedAgent have the same reimbursement agent, then only InstructingReimbursementAgent must be used.
	InstructedReimbursementAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstdRmbrsmntAgt,omitempty"`

	// Unambiguous identification of the account of the instructed reimbursement agent account at its servicing agent in the payment chain.
	InstructedReimbursementAgentAccount *CashAccount7 `xml:"InstdRmbrsmntAgtAcct,omitempty"`

	// Specifies the branch of the instructed agent where the amount of money will be made available when different from the instructed reimbursement agent.
	ThirdReimbursementAgent *BranchAndFinancialInstitutionIdentification3 `xml:"ThrdRmbrsmntAgt,omitempty"`

	// Unambiguous identification of the account of the third reimbursement agent account at its servicing agent in the payment chain.
	ThirdReimbursementAgentAccount *CashAccount7 `xml:"ThrdRmbrsmntAgtAcct,omitempty"`
}

func (s *SettlementInformation3) SetSettlementMethod(value string) {
	s.SettlementMethod = (*SettlementMethod1Code)(&value)
}

func (s *SettlementInformation3) AddSettlementAccount() *CashAccount7 {
	s.SettlementAccount = new(CashAccount7)
	return s.SettlementAccount
}

func (s *SettlementInformation3) AddClearingSystem() *ClearingSystemIdentification1Choice {
	s.ClearingSystem = new(ClearingSystemIdentification1Choice)
	return s.ClearingSystem
}

func (s *SettlementInformation3) AddInstructingReimbursementAgent() *BranchAndFinancialInstitutionIdentification3 {
	s.InstructingReimbursementAgent = new(BranchAndFinancialInstitutionIdentification3)
	return s.InstructingReimbursementAgent
}

func (s *SettlementInformation3) AddInstructingReimbursementAgentAccount() *CashAccount7 {
	s.InstructingReimbursementAgentAccount = new(CashAccount7)
	return s.InstructingReimbursementAgentAccount
}

func (s *SettlementInformation3) AddInstructedReimbursementAgent() *BranchAndFinancialInstitutionIdentification3 {
	s.InstructedReimbursementAgent = new(BranchAndFinancialInstitutionIdentification3)
	return s.InstructedReimbursementAgent
}

func (s *SettlementInformation3) AddInstructedReimbursementAgentAccount() *CashAccount7 {
	s.InstructedReimbursementAgentAccount = new(CashAccount7)
	return s.InstructedReimbursementAgentAccount
}

func (s *SettlementInformation3) AddThirdReimbursementAgent() *BranchAndFinancialInstitutionIdentification3 {
	s.ThirdReimbursementAgent = new(BranchAndFinancialInstitutionIdentification3)
	return s.ThirdReimbursementAgent
}

func (s *SettlementInformation3) AddThirdReimbursementAgentAccount() *CashAccount7 {
	s.ThirdReimbursementAgentAccount = new(CashAccount7)
	return s.ThirdReimbursementAgentAccount
}
