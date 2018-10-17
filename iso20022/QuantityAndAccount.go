package iso20022

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount1 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *Max210Text `xml:"DnmtnChc,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown3 `xml:"QtyBrkdwn,omitempty"`
}

func (q *QuantityAndAccount1) AddSettlementQuantity() *Quantity6Choice {
	q.SettlementQuantity = new(Quantity6Choice)
	return q.SettlementQuantity
}

func (q *QuantityAndAccount1) SetDenominationChoice(value string) {
	q.DenominationChoice = (*Max210Text)(&value)
}

func (q *QuantityAndAccount1) AddAccountOwner() *PartyIdentification13Choice {
	q.AccountOwner = new(PartyIdentification13Choice)
	return q.AccountOwner
}

func (q *QuantityAndAccount1) AddSafekeepingAccount() *SecuritiesAccount13 {
	q.SafekeepingAccount = new(SecuritiesAccount13)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount1) AddCashAccount() *CashAccountIdentification5Choice {
	q.CashAccount = new(CashAccountIdentification5Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount1) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	q.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return q.SafekeepingPlace
}

func (q *QuantityAndAccount1) AddQuantityBreakdown() *QuantityBreakdown3 {
	newValue := new(QuantityBreakdown3)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount13 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity1Choice `xml:"SttlmQty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *Max210Text `xml:"DnmtnChc,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown3 `xml:"QtyBrkdwn,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount13) AddSettlementQuantity() *FinancialInstrumentQuantity1Choice {
	q.SettlementQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.SettlementQuantity
}

func (q *QuantityAndAccount13) SetDenominationChoice(value string) {
	q.DenominationChoice = (*Max210Text)(&value)
}

func (q *QuantityAndAccount13) AddAccountOwner() *PartyIdentification36Choice {
	q.AccountOwner = new(PartyIdentification36Choice)
	return q.AccountOwner
}

func (q *QuantityAndAccount13) AddSafekeepingAccount() *SecuritiesAccount13 {
	q.SafekeepingAccount = new(SecuritiesAccount13)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount13) AddCashAccount() *CashAccountIdentification5Choice {
	q.CashAccount = new(CashAccountIdentification5Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount13) AddQuantityBreakdown() *QuantityBreakdown3 {
	newValue := new(QuantityBreakdown3)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

func (q *QuantityAndAccount13) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	q.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return q.SafekeepingPlace
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount14 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity1Choice `xml:"SttlmQty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *Max210Text `xml:"DnmtnChc,omitempty"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown3 `xml:"QtyBrkdwn,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount14) AddSettlementQuantity() *FinancialInstrumentQuantity1Choice {
	q.SettlementQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.SettlementQuantity
}

func (q *QuantityAndAccount14) SetDenominationChoice(value string) {
	q.DenominationChoice = (*Max210Text)(&value)
}

func (q *QuantityAndAccount14) AddCashAccount() *CashAccountIdentification5Choice {
	q.CashAccount = new(CashAccountIdentification5Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount14) AddQuantityBreakdown() *QuantityBreakdown3 {
	newValue := new(QuantityBreakdown3)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

func (q *QuantityAndAccount14) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	q.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return q.SafekeepingPlace
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount15 struct {

	// Quantity of financial instrument effectively settled.
	SettledQuantity *Quantity6Choice `xml:"SttldQty"`

	// Quantity of financial instrument previously settled.
	PreviouslySettledQuantity *FinancialInstrumentQuantity1Choice `xml:"PrevslySttldQty,omitempty"`

	// Quantity of financial instrument remaining to be settled.
	RemainingToBeSettledQuantity *FinancialInstrumentQuantity1Choice `xml:"RmngToBeSttldQty,omitempty"`

	// Amount of money previously settled.
	PreviouslySettledAmount *AmountAndDirection7 `xml:"PrevslySttldAmt,omitempty"`

	// Amount of money remaining to be settled.
	RemainingToBeSettledAmount *AmountAndDirection7 `xml:"RmngToBeSttldAmt,omitempty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *Max210Text `xml:"DnmtnChc,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown9 `xml:"QtyBrkdwn,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount15) AddSettledQuantity() *Quantity6Choice {
	q.SettledQuantity = new(Quantity6Choice)
	return q.SettledQuantity
}

func (q *QuantityAndAccount15) AddPreviouslySettledQuantity() *FinancialInstrumentQuantity1Choice {
	q.PreviouslySettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.PreviouslySettledQuantity
}

func (q *QuantityAndAccount15) AddRemainingToBeSettledQuantity() *FinancialInstrumentQuantity1Choice {
	q.RemainingToBeSettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.RemainingToBeSettledQuantity
}

func (q *QuantityAndAccount15) AddPreviouslySettledAmount() *AmountAndDirection7 {
	q.PreviouslySettledAmount = new(AmountAndDirection7)
	return q.PreviouslySettledAmount
}

func (q *QuantityAndAccount15) AddRemainingToBeSettledAmount() *AmountAndDirection7 {
	q.RemainingToBeSettledAmount = new(AmountAndDirection7)
	return q.RemainingToBeSettledAmount
}

func (q *QuantityAndAccount15) SetDenominationChoice(value string) {
	q.DenominationChoice = (*Max210Text)(&value)
}

func (q *QuantityAndAccount15) AddAccountOwner() *PartyIdentification36Choice {
	q.AccountOwner = new(PartyIdentification36Choice)
	return q.AccountOwner
}

func (q *QuantityAndAccount15) AddSafekeepingAccount() *SecuritiesAccount13 {
	q.SafekeepingAccount = new(SecuritiesAccount13)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount15) AddCashAccount() *CashAccountIdentification5Choice {
	q.CashAccount = new(CashAccountIdentification5Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount15) AddQuantityBreakdown() *QuantityBreakdown9 {
	newValue := new(QuantityBreakdown9)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

func (q *QuantityAndAccount15) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	q.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return q.SafekeepingPlace
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount16 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity1Choice `xml:"SttlmQty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount16) AddSettlementQuantity() *FinancialInstrumentQuantity1Choice {
	q.SettlementQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.SettlementQuantity
}

func (q *QuantityAndAccount16) AddAccountOwner() *PartyIdentification36Choice {
	q.AccountOwner = new(PartyIdentification36Choice)
	return q.AccountOwner
}

func (q *QuantityAndAccount16) AddSafekeepingAccount() *SecuritiesAccount13 {
	q.SafekeepingAccount = new(SecuritiesAccount13)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount16) AddCashAccount() *CashAccountIdentification5Choice {
	q.CashAccount = new(CashAccountIdentification5Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount16) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	q.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return q.SafekeepingPlace
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount17 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *Max210Text `xml:"DnmtnChc,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown3 `xml:"QtyBrkdwn,omitempty"`
}

func (q *QuantityAndAccount17) AddSettlementQuantity() *Quantity6Choice {
	q.SettlementQuantity = new(Quantity6Choice)
	return q.SettlementQuantity
}

func (q *QuantityAndAccount17) SetDenominationChoice(value string) {
	q.DenominationChoice = (*Max210Text)(&value)
}

func (q *QuantityAndAccount17) AddAccountOwner() *PartyIdentification36Choice {
	q.AccountOwner = new(PartyIdentification36Choice)
	return q.AccountOwner
}

func (q *QuantityAndAccount17) AddSafekeepingAccount() *SecuritiesAccount13 {
	q.SafekeepingAccount = new(SecuritiesAccount13)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount17) AddCashAccount() *CashAccountIdentification5Choice {
	q.CashAccount = new(CashAccountIdentification5Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount17) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	q.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return q.SafekeepingPlace
}

func (q *QuantityAndAccount17) AddQuantityBreakdown() *QuantityBreakdown3 {
	newValue := new(QuantityBreakdown3)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount18 struct {

	// Quantity of financial instrument effectively settled.
	SettledQuantity *Quantity6Choice `xml:"SttldQty"`

	// Quantity of financial instrument previously settled.
	PreviouslySettledQuantity *FinancialInstrumentQuantity1Choice `xml:"PrevslySttldQty,omitempty"`

	// Quantity of financial instrument remaining to be settled.
	RemainingToBeSettledQuantity *FinancialInstrumentQuantity1Choice `xml:"RmngToBeSttldQty,omitempty"`

	// Amount of money previously settled.
	PreviouslySettledAmount *AmountAndDirection5 `xml:"PrevslySttldAmt,omitempty"`

	// Amount of money remaining to be settled.
	RemainingToBeSettledAmount *AmountAndDirection5 `xml:"RmngToBeSttldAmt,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount18) AddSettledQuantity() *Quantity6Choice {
	q.SettledQuantity = new(Quantity6Choice)
	return q.SettledQuantity
}

func (q *QuantityAndAccount18) AddPreviouslySettledQuantity() *FinancialInstrumentQuantity1Choice {
	q.PreviouslySettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.PreviouslySettledQuantity
}

func (q *QuantityAndAccount18) AddRemainingToBeSettledQuantity() *FinancialInstrumentQuantity1Choice {
	q.RemainingToBeSettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.RemainingToBeSettledQuantity
}

func (q *QuantityAndAccount18) AddPreviouslySettledAmount() *AmountAndDirection5 {
	q.PreviouslySettledAmount = new(AmountAndDirection5)
	return q.PreviouslySettledAmount
}

func (q *QuantityAndAccount18) AddRemainingToBeSettledAmount() *AmountAndDirection5 {
	q.RemainingToBeSettledAmount = new(AmountAndDirection5)
	return q.RemainingToBeSettledAmount
}

func (q *QuantityAndAccount18) AddAccountOwner() *PartyIdentification36Choice {
	q.AccountOwner = new(PartyIdentification36Choice)
	return q.AccountOwner
}

func (q *QuantityAndAccount18) AddSafekeepingAccount() *SecuritiesAccount13 {
	q.SafekeepingAccount = new(SecuritiesAccount13)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount18) AddCashAccount() *CashAccountIdentification5Choice {
	q.CashAccount = new(CashAccountIdentification5Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount18) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	q.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return q.SafekeepingPlace
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount2 struct {

	// Quantity of financial instrument effectively settled.
	SettledQuantity *Quantity6Choice `xml:"SttldQty"`

	// Quantity of financial instrument previously settled.
	PreviouslySettledQuantity *FinancialInstrumentQuantity1Choice `xml:"PrevslySttldQty,omitempty"`

	// Quantity of financial instrument remaining to be settled.
	RemainingToBeSettledQuantity *FinancialInstrumentQuantity1Choice `xml:"RmngToBeSttldQty,omitempty"`

	// Amount of money previously settled.
	PreviouslySettledAmount *AmountAndDirection7 `xml:"PrevslySttldAmt,omitempty"`

	// Amount of money remaining to be settled.
	RemainingToBeSettledAmount *AmountAndDirection7 `xml:"RmngToBeSttldAmt,omitempty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *Max210Text `xml:"DnmtnChc,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown3 `xml:"QtyBrkdwn,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount2) AddSettledQuantity() *Quantity6Choice {
	q.SettledQuantity = new(Quantity6Choice)
	return q.SettledQuantity
}

func (q *QuantityAndAccount2) AddPreviouslySettledQuantity() *FinancialInstrumentQuantity1Choice {
	q.PreviouslySettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.PreviouslySettledQuantity
}

func (q *QuantityAndAccount2) AddRemainingToBeSettledQuantity() *FinancialInstrumentQuantity1Choice {
	q.RemainingToBeSettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.RemainingToBeSettledQuantity
}

func (q *QuantityAndAccount2) AddPreviouslySettledAmount() *AmountAndDirection7 {
	q.PreviouslySettledAmount = new(AmountAndDirection7)
	return q.PreviouslySettledAmount
}

func (q *QuantityAndAccount2) AddRemainingToBeSettledAmount() *AmountAndDirection7 {
	q.RemainingToBeSettledAmount = new(AmountAndDirection7)
	return q.RemainingToBeSettledAmount
}

func (q *QuantityAndAccount2) SetDenominationChoice(value string) {
	q.DenominationChoice = (*Max210Text)(&value)
}

func (q *QuantityAndAccount2) AddAccountOwner() *PartyIdentification13Choice {
	q.AccountOwner = new(PartyIdentification13Choice)
	return q.AccountOwner
}

func (q *QuantityAndAccount2) AddSafekeepingAccount() *SecuritiesAccount13 {
	q.SafekeepingAccount = new(SecuritiesAccount13)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount2) AddCashAccount() *CashAccountIdentification5Choice {
	q.CashAccount = new(CashAccountIdentification5Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount2) AddQuantityBreakdown() *QuantityBreakdown3 {
	newValue := new(QuantityBreakdown3)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

func (q *QuantityAndAccount2) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	q.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return q.SafekeepingPlace
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount25 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *Max210Text `xml:"DnmtnChc,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown13 `xml:"QtyBrkdwn,omitempty"`
}

func (q *QuantityAndAccount25) AddSettlementQuantity() *Quantity6Choice {
	q.SettlementQuantity = new(Quantity6Choice)
	return q.SettlementQuantity
}

func (q *QuantityAndAccount25) SetDenominationChoice(value string) {
	q.DenominationChoice = (*Max210Text)(&value)
}

func (q *QuantityAndAccount25) AddAccountOwner() *PartyIdentification36Choice {
	q.AccountOwner = new(PartyIdentification36Choice)
	return q.AccountOwner
}

func (q *QuantityAndAccount25) AddSafekeepingAccount() *SecuritiesAccount13 {
	q.SafekeepingAccount = new(SecuritiesAccount13)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount25) AddCashAccount() *CashAccountIdentification5Choice {
	q.CashAccount = new(CashAccountIdentification5Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount25) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	q.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return q.SafekeepingPlace
}

func (q *QuantityAndAccount25) AddQuantityBreakdown() *QuantityBreakdown13 {
	newValue := new(QuantityBreakdown13)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount26 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity1Choice `xml:"SttlmQty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *Max210Text `xml:"DnmtnChc,omitempty"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown13 `xml:"QtyBrkdwn,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount26) AddSettlementQuantity() *FinancialInstrumentQuantity1Choice {
	q.SettlementQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.SettlementQuantity
}

func (q *QuantityAndAccount26) SetDenominationChoice(value string) {
	q.DenominationChoice = (*Max210Text)(&value)
}

func (q *QuantityAndAccount26) AddCashAccount() *CashAccountIdentification5Choice {
	q.CashAccount = new(CashAccountIdentification5Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount26) AddQuantityBreakdown() *QuantityBreakdown13 {
	newValue := new(QuantityBreakdown13)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

func (q *QuantityAndAccount26) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	q.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return q.SafekeepingPlace
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount27 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity1Choice `xml:"SttlmQty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *Max210Text `xml:"DnmtnChc,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown13 `xml:"QtyBrkdwn,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount27) AddSettlementQuantity() *FinancialInstrumentQuantity1Choice {
	q.SettlementQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.SettlementQuantity
}

func (q *QuantityAndAccount27) SetDenominationChoice(value string) {
	q.DenominationChoice = (*Max210Text)(&value)
}

func (q *QuantityAndAccount27) AddAccountOwner() *PartyIdentification36Choice {
	q.AccountOwner = new(PartyIdentification36Choice)
	return q.AccountOwner
}

func (q *QuantityAndAccount27) AddSafekeepingAccount() *SecuritiesAccount13 {
	q.SafekeepingAccount = new(SecuritiesAccount13)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount27) AddCashAccount() *CashAccountIdentification5Choice {
	q.CashAccount = new(CashAccountIdentification5Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount27) AddQuantityBreakdown() *QuantityBreakdown13 {
	newValue := new(QuantityBreakdown13)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

func (q *QuantityAndAccount27) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	q.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return q.SafekeepingPlace
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount28 struct {

	// Quantity of financial instrument effectively settled.
	SettledQuantity *Quantity6Choice `xml:"SttldQty"`

	// Quantity of financial instrument previously settled.
	PreviouslySettledQuantity *FinancialInstrumentQuantity1Choice `xml:"PrevslySttldQty,omitempty"`

	// Quantity of financial instrument remaining to be settled.
	RemainingToBeSettledQuantity *FinancialInstrumentQuantity1Choice `xml:"RmngToBeSttldQty,omitempty"`

	// Amount of money previously settled.
	PreviouslySettledAmount *AmountAndDirection7 `xml:"PrevslySttldAmt,omitempty"`

	// Amount of money remaining to be settled.
	RemainingToBeSettledAmount *AmountAndDirection7 `xml:"RmngToBeSttldAmt,omitempty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *Max210Text `xml:"DnmtnChc,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown16 `xml:"QtyBrkdwn,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount28) AddSettledQuantity() *Quantity6Choice {
	q.SettledQuantity = new(Quantity6Choice)
	return q.SettledQuantity
}

func (q *QuantityAndAccount28) AddPreviouslySettledQuantity() *FinancialInstrumentQuantity1Choice {
	q.PreviouslySettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.PreviouslySettledQuantity
}

func (q *QuantityAndAccount28) AddRemainingToBeSettledQuantity() *FinancialInstrumentQuantity1Choice {
	q.RemainingToBeSettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.RemainingToBeSettledQuantity
}

func (q *QuantityAndAccount28) AddPreviouslySettledAmount() *AmountAndDirection7 {
	q.PreviouslySettledAmount = new(AmountAndDirection7)
	return q.PreviouslySettledAmount
}

func (q *QuantityAndAccount28) AddRemainingToBeSettledAmount() *AmountAndDirection7 {
	q.RemainingToBeSettledAmount = new(AmountAndDirection7)
	return q.RemainingToBeSettledAmount
}

func (q *QuantityAndAccount28) SetDenominationChoice(value string) {
	q.DenominationChoice = (*Max210Text)(&value)
}

func (q *QuantityAndAccount28) AddAccountOwner() *PartyIdentification36Choice {
	q.AccountOwner = new(PartyIdentification36Choice)
	return q.AccountOwner
}

func (q *QuantityAndAccount28) AddSafekeepingAccount() *SecuritiesAccount13 {
	q.SafekeepingAccount = new(SecuritiesAccount13)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount28) AddCashAccount() *CashAccountIdentification5Choice {
	q.CashAccount = new(CashAccountIdentification5Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount28) AddQuantityBreakdown() *QuantityBreakdown16 {
	newValue := new(QuantityBreakdown16)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

func (q *QuantityAndAccount28) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	q.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return q.SafekeepingPlace
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount29 struct {

	// Denomination of the security to be received or delivered.
	DenominationChoice *Max210Text `xml:"DnmtnChc,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown13 `xml:"QtyBrkdwn,omitempty"`
}

func (q *QuantityAndAccount29) SetDenominationChoice(value string) {
	q.DenominationChoice = (*Max210Text)(&value)
}

func (q *QuantityAndAccount29) AddAccountOwner() *PartyIdentification36Choice {
	q.AccountOwner = new(PartyIdentification36Choice)
	return q.AccountOwner
}

func (q *QuantityAndAccount29) AddCashAccount() *CashAccountIdentification5Choice {
	q.CashAccount = new(CashAccountIdentification5Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount29) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	q.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return q.SafekeepingPlace
}

func (q *QuantityAndAccount29) AddQuantityBreakdown() *QuantityBreakdown13 {
	newValue := new(QuantityBreakdown13)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount3 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity1Choice `xml:"SttlmQty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *Max210Text `xml:"DnmtnChc,omitempty"`

	// Unique and unambiguous identifier of a certificate assigned by the issuer.
	CertificateNumber []*SecuritiesCertificate1 `xml:"CertNb,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown3 `xml:"QtyBrkdwn,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount3) AddSettlementQuantity() *FinancialInstrumentQuantity1Choice {
	q.SettlementQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.SettlementQuantity
}

func (q *QuantityAndAccount3) SetDenominationChoice(value string) {
	q.DenominationChoice = (*Max210Text)(&value)
}

func (q *QuantityAndAccount3) AddCertificateNumber() *SecuritiesCertificate1 {
	newValue := new(SecuritiesCertificate1)
	q.CertificateNumber = append(q.CertificateNumber, newValue)
	return newValue
}

func (q *QuantityAndAccount3) AddAccountOwner() *PartyIdentification13Choice {
	q.AccountOwner = new(PartyIdentification13Choice)
	return q.AccountOwner
}

func (q *QuantityAndAccount3) AddSafekeepingAccount() *SecuritiesAccount13 {
	q.SafekeepingAccount = new(SecuritiesAccount13)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount3) AddCashAccount() *CashAccountIdentification5Choice {
	q.CashAccount = new(CashAccountIdentification5Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount3) AddQuantityBreakdown() *QuantityBreakdown3 {
	newValue := new(QuantityBreakdown3)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

func (q *QuantityAndAccount3) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	q.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return q.SafekeepingPlace
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount30 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty,omitempty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *Max210Text `xml:"DnmtnChc,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct,omitempty"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown13 `xml:"QtyBrkdwn,omitempty"`
}

func (q *QuantityAndAccount30) AddSettlementQuantity() *Quantity6Choice {
	q.SettlementQuantity = new(Quantity6Choice)
	return q.SettlementQuantity
}

func (q *QuantityAndAccount30) SetDenominationChoice(value string) {
	q.DenominationChoice = (*Max210Text)(&value)
}

func (q *QuantityAndAccount30) AddAccountOwner() *PartyIdentification36Choice {
	q.AccountOwner = new(PartyIdentification36Choice)
	return q.AccountOwner
}

func (q *QuantityAndAccount30) AddSafekeepingAccount() *SecuritiesAccount13 {
	q.SafekeepingAccount = new(SecuritiesAccount13)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount30) AddCashAccount() *CashAccountIdentification5Choice {
	q.CashAccount = new(CashAccountIdentification5Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount30) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	q.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return q.SafekeepingPlace
}

func (q *QuantityAndAccount30) AddQuantityBreakdown() *QuantityBreakdown13 {
	newValue := new(QuantityBreakdown13)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount38 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity1Choice `xml:"SttlmQty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount19 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace1 `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount38) AddSettlementQuantity() *FinancialInstrumentQuantity1Choice {
	q.SettlementQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.SettlementQuantity
}

func (q *QuantityAndAccount38) AddAccountOwner() *PartyIdentification98 {
	q.AccountOwner = new(PartyIdentification98)
	return q.AccountOwner
}

func (q *QuantityAndAccount38) AddSafekeepingAccount() *SecuritiesAccount19 {
	q.SafekeepingAccount = new(SecuritiesAccount19)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount38) AddCashAccount() *CashAccountIdentification5Choice {
	q.CashAccount = new(CashAccountIdentification5Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount38) AddSafekeepingPlace() *SafeKeepingPlace1 {
	q.SafekeepingPlace = new(SafeKeepingPlace1)
	return q.SafekeepingPlace
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount39 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *Max210Text `xml:"DnmtnChc,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount19 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace1 `xml:"SfkpgPlc,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown30 `xml:"QtyBrkdwn,omitempty"`
}

func (q *QuantityAndAccount39) AddSettlementQuantity() *Quantity6Choice {
	q.SettlementQuantity = new(Quantity6Choice)
	return q.SettlementQuantity
}

func (q *QuantityAndAccount39) SetDenominationChoice(value string) {
	q.DenominationChoice = (*Max210Text)(&value)
}

func (q *QuantityAndAccount39) AddAccountOwner() *PartyIdentification98 {
	q.AccountOwner = new(PartyIdentification98)
	return q.AccountOwner
}

func (q *QuantityAndAccount39) AddSafekeepingAccount() *SecuritiesAccount19 {
	q.SafekeepingAccount = new(SecuritiesAccount19)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount39) AddCashAccount() *CashAccountIdentification5Choice {
	q.CashAccount = new(CashAccountIdentification5Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount39) AddSafekeepingPlace() *SafeKeepingPlace1 {
	q.SafekeepingPlace = new(SafeKeepingPlace1)
	return q.SafekeepingPlace
}

func (q *QuantityAndAccount39) AddQuantityBreakdown() *QuantityBreakdown30 {
	newValue := new(QuantityBreakdown30)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount40 struct {

	// Quantity of financial instrument effectively settled.
	SettledQuantity *Quantity6Choice `xml:"SttldQty"`

	// Quantity of financial instrument previously settled.
	PreviouslySettledQuantity *FinancialInstrumentQuantity1Choice `xml:"PrevslySttldQty,omitempty"`

	// Quantity of financial instrument remaining to be settled.
	RemainingToBeSettledQuantity *FinancialInstrumentQuantity1Choice `xml:"RmngToBeSttldQty,omitempty"`

	// Amount of money previously settled.
	PreviouslySettledAmount *AmountAndDirection5 `xml:"PrevslySttldAmt,omitempty"`

	// Amount of money remaining to be settled.
	RemainingToBeSettledAmount *AmountAndDirection5 `xml:"RmngToBeSttldAmt,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount19 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace1 `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount40) AddSettledQuantity() *Quantity6Choice {
	q.SettledQuantity = new(Quantity6Choice)
	return q.SettledQuantity
}

func (q *QuantityAndAccount40) AddPreviouslySettledQuantity() *FinancialInstrumentQuantity1Choice {
	q.PreviouslySettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.PreviouslySettledQuantity
}

func (q *QuantityAndAccount40) AddRemainingToBeSettledQuantity() *FinancialInstrumentQuantity1Choice {
	q.RemainingToBeSettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.RemainingToBeSettledQuantity
}

func (q *QuantityAndAccount40) AddPreviouslySettledAmount() *AmountAndDirection5 {
	q.PreviouslySettledAmount = new(AmountAndDirection5)
	return q.PreviouslySettledAmount
}

func (q *QuantityAndAccount40) AddRemainingToBeSettledAmount() *AmountAndDirection5 {
	q.RemainingToBeSettledAmount = new(AmountAndDirection5)
	return q.RemainingToBeSettledAmount
}

func (q *QuantityAndAccount40) AddAccountOwner() *PartyIdentification98 {
	q.AccountOwner = new(PartyIdentification98)
	return q.AccountOwner
}

func (q *QuantityAndAccount40) AddSafekeepingAccount() *SecuritiesAccount19 {
	q.SafekeepingAccount = new(SecuritiesAccount19)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount40) AddCashAccount() *CashAccountIdentification5Choice {
	q.CashAccount = new(CashAccountIdentification5Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount40) AddSafekeepingPlace() *SafeKeepingPlace1 {
	q.SafekeepingPlace = new(SafeKeepingPlace1)
	return q.SafekeepingPlace
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount41 struct {

	// Quantity of financial instrument effectively settled.
	SettledQuantity *Quantity6Choice `xml:"SttldQty"`

	// Quantity of financial instrument previously settled.
	PreviouslySettledQuantity *FinancialInstrumentQuantity1Choice `xml:"PrevslySttldQty,omitempty"`

	// Quantity of financial instrument remaining to be settled.
	RemainingToBeSettledQuantity *FinancialInstrumentQuantity1Choice `xml:"RmngToBeSttldQty,omitempty"`

	// Amount of money previously settled.
	PreviouslySettledAmount *AmountAndDirection52 `xml:"PrevslySttldAmt,omitempty"`

	// Amount of money remaining to be settled.
	RemainingToBeSettledAmount *AmountAndDirection52 `xml:"RmngToBeSttldAmt,omitempty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *Max210Text `xml:"DnmtnChc,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount24 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown29 `xml:"QtyBrkdwn,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace1 `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount41) AddSettledQuantity() *Quantity6Choice {
	q.SettledQuantity = new(Quantity6Choice)
	return q.SettledQuantity
}

func (q *QuantityAndAccount41) AddPreviouslySettledQuantity() *FinancialInstrumentQuantity1Choice {
	q.PreviouslySettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.PreviouslySettledQuantity
}

func (q *QuantityAndAccount41) AddRemainingToBeSettledQuantity() *FinancialInstrumentQuantity1Choice {
	q.RemainingToBeSettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.RemainingToBeSettledQuantity
}

func (q *QuantityAndAccount41) AddPreviouslySettledAmount() *AmountAndDirection52 {
	q.PreviouslySettledAmount = new(AmountAndDirection52)
	return q.PreviouslySettledAmount
}

func (q *QuantityAndAccount41) AddRemainingToBeSettledAmount() *AmountAndDirection52 {
	q.RemainingToBeSettledAmount = new(AmountAndDirection52)
	return q.RemainingToBeSettledAmount
}

func (q *QuantityAndAccount41) SetDenominationChoice(value string) {
	q.DenominationChoice = (*Max210Text)(&value)
}

func (q *QuantityAndAccount41) AddAccountOwner() *PartyIdentification98 {
	q.AccountOwner = new(PartyIdentification98)
	return q.AccountOwner
}

func (q *QuantityAndAccount41) AddSafekeepingAccount() *SecuritiesAccount24 {
	q.SafekeepingAccount = new(SecuritiesAccount24)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount41) AddCashAccount() *CashAccountIdentification5Choice {
	q.CashAccount = new(CashAccountIdentification5Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount41) AddQuantityBreakdown() *QuantityBreakdown29 {
	newValue := new(QuantityBreakdown29)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

func (q *QuantityAndAccount41) AddSafekeepingPlace() *SafeKeepingPlace1 {
	q.SafekeepingPlace = new(SafeKeepingPlace1)
	return q.SafekeepingPlace
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount42 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity1Choice `xml:"SttlmQty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *Max210Text `xml:"DnmtnChc,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount24 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown30 `xml:"QtyBrkdwn,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace1 `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount42) AddSettlementQuantity() *FinancialInstrumentQuantity1Choice {
	q.SettlementQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.SettlementQuantity
}

func (q *QuantityAndAccount42) SetDenominationChoice(value string) {
	q.DenominationChoice = (*Max210Text)(&value)
}

func (q *QuantityAndAccount42) AddAccountOwner() *PartyIdentification98 {
	q.AccountOwner = new(PartyIdentification98)
	return q.AccountOwner
}

func (q *QuantityAndAccount42) AddSafekeepingAccount() *SecuritiesAccount24 {
	q.SafekeepingAccount = new(SecuritiesAccount24)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount42) AddCashAccount() *CashAccountIdentification5Choice {
	q.CashAccount = new(CashAccountIdentification5Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount42) AddQuantityBreakdown() *QuantityBreakdown30 {
	newValue := new(QuantityBreakdown30)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

func (q *QuantityAndAccount42) AddSafekeepingPlace() *SafeKeepingPlace1 {
	q.SafekeepingPlace = new(SafeKeepingPlace1)
	return q.SafekeepingPlace
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount43 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty,omitempty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *Max210Text `xml:"DnmtnChc,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount19 `xml:"SfkpgAcct,omitempty"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace1 `xml:"SfkpgPlc,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown30 `xml:"QtyBrkdwn,omitempty"`
}

func (q *QuantityAndAccount43) AddSettlementQuantity() *Quantity6Choice {
	q.SettlementQuantity = new(Quantity6Choice)
	return q.SettlementQuantity
}

func (q *QuantityAndAccount43) SetDenominationChoice(value string) {
	q.DenominationChoice = (*Max210Text)(&value)
}

func (q *QuantityAndAccount43) AddAccountOwner() *PartyIdentification98 {
	q.AccountOwner = new(PartyIdentification98)
	return q.AccountOwner
}

func (q *QuantityAndAccount43) AddSafekeepingAccount() *SecuritiesAccount19 {
	q.SafekeepingAccount = new(SecuritiesAccount19)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount43) AddCashAccount() *CashAccountIdentification5Choice {
	q.CashAccount = new(CashAccountIdentification5Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount43) AddSafekeepingPlace() *SafeKeepingPlace1 {
	q.SafekeepingPlace = new(SafeKeepingPlace1)
	return q.SafekeepingPlace
}

func (q *QuantityAndAccount43) AddQuantityBreakdown() *QuantityBreakdown30 {
	newValue := new(QuantityBreakdown30)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount44 struct {

	// Denomination of the security to be received or delivered.
	DenominationChoice *Max210Text `xml:"DnmtnChc,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace1 `xml:"SfkpgPlc,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown30 `xml:"QtyBrkdwn,omitempty"`
}

func (q *QuantityAndAccount44) SetDenominationChoice(value string) {
	q.DenominationChoice = (*Max210Text)(&value)
}

func (q *QuantityAndAccount44) AddAccountOwner() *PartyIdentification98 {
	q.AccountOwner = new(PartyIdentification98)
	return q.AccountOwner
}

func (q *QuantityAndAccount44) AddCashAccount() *CashAccountIdentification5Choice {
	q.CashAccount = new(CashAccountIdentification5Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount44) AddSafekeepingPlace() *SafeKeepingPlace1 {
	q.SafekeepingPlace = new(SafeKeepingPlace1)
	return q.SafekeepingPlace
}

func (q *QuantityAndAccount44) AddQuantityBreakdown() *QuantityBreakdown30 {
	newValue := new(QuantityBreakdown30)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount45 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity1Choice `xml:"SttlmQty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *Max210Text `xml:"DnmtnChc,omitempty"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown30 `xml:"QtyBrkdwn,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace1 `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount45) AddSettlementQuantity() *FinancialInstrumentQuantity1Choice {
	q.SettlementQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.SettlementQuantity
}

func (q *QuantityAndAccount45) SetDenominationChoice(value string) {
	q.DenominationChoice = (*Max210Text)(&value)
}

func (q *QuantityAndAccount45) AddCashAccount() *CashAccountIdentification5Choice {
	q.CashAccount = new(CashAccountIdentification5Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount45) AddQuantityBreakdown() *QuantityBreakdown30 {
	newValue := new(QuantityBreakdown30)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

func (q *QuantityAndAccount45) AddSafekeepingPlace() *SafeKeepingPlace1 {
	q.SafekeepingPlace = new(SafeKeepingPlace1)
	return q.SafekeepingPlace
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount46 struct {

	// Quantity of financial instrument effectively settled.
	SettledQuantity *Quantity10Choice `xml:"SttldQty"`

	// Quantity of financial instrument previously settled.
	PreviouslySettledQuantity *FinancialInstrumentQuantity15Choice `xml:"PrevslySttldQty,omitempty"`

	// Quantity of financial instrument remaining to be settled.
	RemainingToBeSettledQuantity *FinancialInstrumentQuantity15Choice `xml:"RmngToBeSttldQty,omitempty"`

	// Amount of money previously settled.
	PreviouslySettledAmount *AmountAndDirection19 `xml:"PrevslySttldAmt,omitempty"`

	// Amount of money remaining to be settled.
	RemainingToBeSettledAmount *AmountAndDirection19 `xml:"RmngToBeSttldAmt,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount30 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification6Choice `xml:"CshAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount46) AddSettledQuantity() *Quantity10Choice {
	q.SettledQuantity = new(Quantity10Choice)
	return q.SettledQuantity
}

func (q *QuantityAndAccount46) AddPreviouslySettledQuantity() *FinancialInstrumentQuantity15Choice {
	q.PreviouslySettledQuantity = new(FinancialInstrumentQuantity15Choice)
	return q.PreviouslySettledQuantity
}

func (q *QuantityAndAccount46) AddRemainingToBeSettledQuantity() *FinancialInstrumentQuantity15Choice {
	q.RemainingToBeSettledQuantity = new(FinancialInstrumentQuantity15Choice)
	return q.RemainingToBeSettledQuantity
}

func (q *QuantityAndAccount46) AddPreviouslySettledAmount() *AmountAndDirection19 {
	q.PreviouslySettledAmount = new(AmountAndDirection19)
	return q.PreviouslySettledAmount
}

func (q *QuantityAndAccount46) AddRemainingToBeSettledAmount() *AmountAndDirection19 {
	q.RemainingToBeSettledAmount = new(AmountAndDirection19)
	return q.RemainingToBeSettledAmount
}

func (q *QuantityAndAccount46) AddAccountOwner() *PartyIdentification109 {
	q.AccountOwner = new(PartyIdentification109)
	return q.AccountOwner
}

func (q *QuantityAndAccount46) AddSafekeepingAccount() *SecuritiesAccount30 {
	q.SafekeepingAccount = new(SecuritiesAccount30)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount46) AddCashAccount() *CashAccountIdentification6Choice {
	q.CashAccount = new(CashAccountIdentification6Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount46) AddSafekeepingPlace() *SafeKeepingPlace2 {
	q.SafekeepingPlace = new(SafeKeepingPlace2)
	return q.SafekeepingPlace
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount47 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity10Choice `xml:"SttlmQty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *RestrictedFINXMax210Text `xml:"DnmtnChc,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount30 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification6Choice `xml:"CshAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown38 `xml:"QtyBrkdwn,omitempty"`
}

func (q *QuantityAndAccount47) AddSettlementQuantity() *Quantity10Choice {
	q.SettlementQuantity = new(Quantity10Choice)
	return q.SettlementQuantity
}

func (q *QuantityAndAccount47) SetDenominationChoice(value string) {
	q.DenominationChoice = (*RestrictedFINXMax210Text)(&value)
}

func (q *QuantityAndAccount47) AddAccountOwner() *PartyIdentification109 {
	q.AccountOwner = new(PartyIdentification109)
	return q.AccountOwner
}

func (q *QuantityAndAccount47) AddSafekeepingAccount() *SecuritiesAccount30 {
	q.SafekeepingAccount = new(SecuritiesAccount30)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount47) AddCashAccount() *CashAccountIdentification6Choice {
	q.CashAccount = new(CashAccountIdentification6Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount47) AddSafekeepingPlace() *SafeKeepingPlace2 {
	q.SafekeepingPlace = new(SafeKeepingPlace2)
	return q.SafekeepingPlace
}

func (q *QuantityAndAccount47) AddQuantityBreakdown() *QuantityBreakdown38 {
	newValue := new(QuantityBreakdown38)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount48 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity15Choice `xml:"SttlmQty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount30 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification6Choice `xml:"CshAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount48) AddSettlementQuantity() *FinancialInstrumentQuantity15Choice {
	q.SettlementQuantity = new(FinancialInstrumentQuantity15Choice)
	return q.SettlementQuantity
}

func (q *QuantityAndAccount48) AddAccountOwner() *PartyIdentification109 {
	q.AccountOwner = new(PartyIdentification109)
	return q.AccountOwner
}

func (q *QuantityAndAccount48) AddSafekeepingAccount() *SecuritiesAccount30 {
	q.SafekeepingAccount = new(SecuritiesAccount30)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount48) AddCashAccount() *CashAccountIdentification6Choice {
	q.CashAccount = new(CashAccountIdentification6Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount48) AddSafekeepingPlace() *SafeKeepingPlace2 {
	q.SafekeepingPlace = new(SafeKeepingPlace2)
	return q.SafekeepingPlace
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount49 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity15Choice `xml:"SttlmQty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *RestrictedFINXMax210Text `xml:"DnmtnChc,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount27 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification6Choice `xml:"CshAcct,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown38 `xml:"QtyBrkdwn,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount49) AddSettlementQuantity() *FinancialInstrumentQuantity15Choice {
	q.SettlementQuantity = new(FinancialInstrumentQuantity15Choice)
	return q.SettlementQuantity
}

func (q *QuantityAndAccount49) SetDenominationChoice(value string) {
	q.DenominationChoice = (*RestrictedFINXMax210Text)(&value)
}

func (q *QuantityAndAccount49) AddAccountOwner() *PartyIdentification109 {
	q.AccountOwner = new(PartyIdentification109)
	return q.AccountOwner
}

func (q *QuantityAndAccount49) AddSafekeepingAccount() *SecuritiesAccount27 {
	q.SafekeepingAccount = new(SecuritiesAccount27)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount49) AddCashAccount() *CashAccountIdentification6Choice {
	q.CashAccount = new(CashAccountIdentification6Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount49) AddQuantityBreakdown() *QuantityBreakdown38 {
	newValue := new(QuantityBreakdown38)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

func (q *QuantityAndAccount49) AddSafekeepingPlace() *SafeKeepingPlace2 {
	q.SafekeepingPlace = new(SafeKeepingPlace2)
	return q.SafekeepingPlace
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount50 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity15Choice `xml:"SttlmQty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *RestrictedFINXMax210Text `xml:"DnmtnChc,omitempty"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification6Choice `xml:"CshAcct,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown38 `xml:"QtyBrkdwn,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount50) AddSettlementQuantity() *FinancialInstrumentQuantity15Choice {
	q.SettlementQuantity = new(FinancialInstrumentQuantity15Choice)
	return q.SettlementQuantity
}

func (q *QuantityAndAccount50) SetDenominationChoice(value string) {
	q.DenominationChoice = (*RestrictedFINXMax210Text)(&value)
}

func (q *QuantityAndAccount50) AddCashAccount() *CashAccountIdentification6Choice {
	q.CashAccount = new(CashAccountIdentification6Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount50) AddQuantityBreakdown() *QuantityBreakdown38 {
	newValue := new(QuantityBreakdown38)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

func (q *QuantityAndAccount50) AddSafekeepingPlace() *SafeKeepingPlace2 {
	q.SafekeepingPlace = new(SafeKeepingPlace2)
	return q.SafekeepingPlace
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount51 struct {

	// Quantity of financial instrument effectively settled.
	SettledQuantity *Quantity10Choice `xml:"SttldQty"`

	// Quantity of financial instrument previously settled.
	PreviouslySettledQuantity *FinancialInstrumentQuantity15Choice `xml:"PrevslySttldQty,omitempty"`

	// Quantity of financial instrument remaining to be settled.
	RemainingToBeSettledQuantity *FinancialInstrumentQuantity15Choice `xml:"RmngToBeSttldQty,omitempty"`

	// Amount of money previously settled.
	PreviouslySettledAmount *AmountAndDirection57 `xml:"PrevslySttldAmt,omitempty"`

	// Amount of money remaining to be settled.
	RemainingToBeSettledAmount *AmountAndDirection57 `xml:"RmngToBeSttldAmt,omitempty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *RestrictedFINXMax210Text `xml:"DnmtnChc,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount27 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification6Choice `xml:"CshAcct,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown44 `xml:"QtyBrkdwn,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount51) AddSettledQuantity() *Quantity10Choice {
	q.SettledQuantity = new(Quantity10Choice)
	return q.SettledQuantity
}

func (q *QuantityAndAccount51) AddPreviouslySettledQuantity() *FinancialInstrumentQuantity15Choice {
	q.PreviouslySettledQuantity = new(FinancialInstrumentQuantity15Choice)
	return q.PreviouslySettledQuantity
}

func (q *QuantityAndAccount51) AddRemainingToBeSettledQuantity() *FinancialInstrumentQuantity15Choice {
	q.RemainingToBeSettledQuantity = new(FinancialInstrumentQuantity15Choice)
	return q.RemainingToBeSettledQuantity
}

func (q *QuantityAndAccount51) AddPreviouslySettledAmount() *AmountAndDirection57 {
	q.PreviouslySettledAmount = new(AmountAndDirection57)
	return q.PreviouslySettledAmount
}

func (q *QuantityAndAccount51) AddRemainingToBeSettledAmount() *AmountAndDirection57 {
	q.RemainingToBeSettledAmount = new(AmountAndDirection57)
	return q.RemainingToBeSettledAmount
}

func (q *QuantityAndAccount51) SetDenominationChoice(value string) {
	q.DenominationChoice = (*RestrictedFINXMax210Text)(&value)
}

func (q *QuantityAndAccount51) AddAccountOwner() *PartyIdentification109 {
	q.AccountOwner = new(PartyIdentification109)
	return q.AccountOwner
}

func (q *QuantityAndAccount51) AddSafekeepingAccount() *SecuritiesAccount27 {
	q.SafekeepingAccount = new(SecuritiesAccount27)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount51) AddCashAccount() *CashAccountIdentification6Choice {
	q.CashAccount = new(CashAccountIdentification6Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount51) AddQuantityBreakdown() *QuantityBreakdown44 {
	newValue := new(QuantityBreakdown44)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

func (q *QuantityAndAccount51) AddSafekeepingPlace() *SafeKeepingPlace2 {
	q.SafekeepingPlace = new(SafeKeepingPlace2)
	return q.SafekeepingPlace
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount54 struct {

	// Denomination of the security to be received or delivered.
	DenominationChoice *RestrictedFINXMax210Text `xml:"DnmtnChc,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification6Choice `xml:"CshAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown38 `xml:"QtyBrkdwn,omitempty"`
}

func (q *QuantityAndAccount54) SetDenominationChoice(value string) {
	q.DenominationChoice = (*RestrictedFINXMax210Text)(&value)
}

func (q *QuantityAndAccount54) AddAccountOwner() *PartyIdentification109 {
	q.AccountOwner = new(PartyIdentification109)
	return q.AccountOwner
}

func (q *QuantityAndAccount54) AddCashAccount() *CashAccountIdentification6Choice {
	q.CashAccount = new(CashAccountIdentification6Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount54) AddSafekeepingPlace() *SafeKeepingPlace2 {
	q.SafekeepingPlace = new(SafeKeepingPlace2)
	return q.SafekeepingPlace
}

func (q *QuantityAndAccount54) AddQuantityBreakdown() *QuantityBreakdown38 {
	newValue := new(QuantityBreakdown38)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount55 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity10Choice `xml:"SttlmQty,omitempty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *RestrictedFINXMax210Text `xml:"DnmtnChc,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount30 `xml:"SfkpgAcct,omitempty"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification6Choice `xml:"CshAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown38 `xml:"QtyBrkdwn,omitempty"`
}

func (q *QuantityAndAccount55) AddSettlementQuantity() *Quantity10Choice {
	q.SettlementQuantity = new(Quantity10Choice)
	return q.SettlementQuantity
}

func (q *QuantityAndAccount55) SetDenominationChoice(value string) {
	q.DenominationChoice = (*RestrictedFINXMax210Text)(&value)
}

func (q *QuantityAndAccount55) AddAccountOwner() *PartyIdentification109 {
	q.AccountOwner = new(PartyIdentification109)
	return q.AccountOwner
}

func (q *QuantityAndAccount55) AddSafekeepingAccount() *SecuritiesAccount30 {
	q.SafekeepingAccount = new(SecuritiesAccount30)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount55) AddCashAccount() *CashAccountIdentification6Choice {
	q.CashAccount = new(CashAccountIdentification6Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount55) AddSafekeepingPlace() *SafeKeepingPlace2 {
	q.SafekeepingPlace = new(SafeKeepingPlace2)
	return q.SafekeepingPlace
}

func (q *QuantityAndAccount55) AddQuantityBreakdown() *QuantityBreakdown38 {
	newValue := new(QuantityBreakdown38)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount56 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity10Choice `xml:"SttlmQty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *RestrictedFINXMax210Text `xml:"DnmtnChc,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification119 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount30 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification6Choice `xml:"CshAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown38 `xml:"QtyBrkdwn,omitempty"`
}

func (q *QuantityAndAccount56) AddSettlementQuantity() *Quantity10Choice {
	q.SettlementQuantity = new(Quantity10Choice)
	return q.SettlementQuantity
}

func (q *QuantityAndAccount56) SetDenominationChoice(value string) {
	q.DenominationChoice = (*RestrictedFINXMax210Text)(&value)
}

func (q *QuantityAndAccount56) AddAccountOwner() *PartyIdentification119 {
	q.AccountOwner = new(PartyIdentification119)
	return q.AccountOwner
}

func (q *QuantityAndAccount56) AddSafekeepingAccount() *SecuritiesAccount30 {
	q.SafekeepingAccount = new(SecuritiesAccount30)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount56) AddCashAccount() *CashAccountIdentification6Choice {
	q.CashAccount = new(CashAccountIdentification6Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount56) AddSafekeepingPlace() *SafeKeepingPlace2 {
	q.SafekeepingPlace = new(SafeKeepingPlace2)
	return q.SafekeepingPlace
}

func (q *QuantityAndAccount56) AddQuantityBreakdown() *QuantityBreakdown38 {
	newValue := new(QuantityBreakdown38)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount57 struct {

	// Quantity of financial instrument effectively settled.
	SettledQuantity *Quantity10Choice `xml:"SttldQty"`

	// Quantity of financial instrument previously settled.
	PreviouslySettledQuantity *FinancialInstrumentQuantity15Choice `xml:"PrevslySttldQty,omitempty"`

	// Quantity of financial instrument remaining to be settled.
	RemainingToBeSettledQuantity *FinancialInstrumentQuantity15Choice `xml:"RmngToBeSttldQty,omitempty"`

	// Amount of money previously settled.
	PreviouslySettledAmount *AmountAndDirection57 `xml:"PrevslySttldAmt,omitempty"`

	// Amount of money remaining to be settled.
	RemainingToBeSettledAmount *AmountAndDirection57 `xml:"RmngToBeSttldAmt,omitempty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *RestrictedFINXMax210Text `xml:"DnmtnChc,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification119 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount27 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification6Choice `xml:"CshAcct,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown44 `xml:"QtyBrkdwn,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount57) AddSettledQuantity() *Quantity10Choice {
	q.SettledQuantity = new(Quantity10Choice)
	return q.SettledQuantity
}

func (q *QuantityAndAccount57) AddPreviouslySettledQuantity() *FinancialInstrumentQuantity15Choice {
	q.PreviouslySettledQuantity = new(FinancialInstrumentQuantity15Choice)
	return q.PreviouslySettledQuantity
}

func (q *QuantityAndAccount57) AddRemainingToBeSettledQuantity() *FinancialInstrumentQuantity15Choice {
	q.RemainingToBeSettledQuantity = new(FinancialInstrumentQuantity15Choice)
	return q.RemainingToBeSettledQuantity
}

func (q *QuantityAndAccount57) AddPreviouslySettledAmount() *AmountAndDirection57 {
	q.PreviouslySettledAmount = new(AmountAndDirection57)
	return q.PreviouslySettledAmount
}

func (q *QuantityAndAccount57) AddRemainingToBeSettledAmount() *AmountAndDirection57 {
	q.RemainingToBeSettledAmount = new(AmountAndDirection57)
	return q.RemainingToBeSettledAmount
}

func (q *QuantityAndAccount57) SetDenominationChoice(value string) {
	q.DenominationChoice = (*RestrictedFINXMax210Text)(&value)
}

func (q *QuantityAndAccount57) AddAccountOwner() *PartyIdentification119 {
	q.AccountOwner = new(PartyIdentification119)
	return q.AccountOwner
}

func (q *QuantityAndAccount57) AddSafekeepingAccount() *SecuritiesAccount27 {
	q.SafekeepingAccount = new(SecuritiesAccount27)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount57) AddCashAccount() *CashAccountIdentification6Choice {
	q.CashAccount = new(CashAccountIdentification6Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount57) AddQuantityBreakdown() *QuantityBreakdown44 {
	newValue := new(QuantityBreakdown44)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

func (q *QuantityAndAccount57) AddSafekeepingPlace() *SafeKeepingPlace2 {
	q.SafekeepingPlace = new(SafeKeepingPlace2)
	return q.SafekeepingPlace
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount58 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity15Choice `xml:"SttlmQty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *RestrictedFINXMax210Text `xml:"DnmtnChc,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification119 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount27 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification6Choice `xml:"CshAcct,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown38 `xml:"QtyBrkdwn,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount58) AddSettlementQuantity() *FinancialInstrumentQuantity15Choice {
	q.SettlementQuantity = new(FinancialInstrumentQuantity15Choice)
	return q.SettlementQuantity
}

func (q *QuantityAndAccount58) SetDenominationChoice(value string) {
	q.DenominationChoice = (*RestrictedFINXMax210Text)(&value)
}

func (q *QuantityAndAccount58) AddAccountOwner() *PartyIdentification119 {
	q.AccountOwner = new(PartyIdentification119)
	return q.AccountOwner
}

func (q *QuantityAndAccount58) AddSafekeepingAccount() *SecuritiesAccount27 {
	q.SafekeepingAccount = new(SecuritiesAccount27)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount58) AddCashAccount() *CashAccountIdentification6Choice {
	q.CashAccount = new(CashAccountIdentification6Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount58) AddQuantityBreakdown() *QuantityBreakdown38 {
	newValue := new(QuantityBreakdown38)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

func (q *QuantityAndAccount58) AddSafekeepingPlace() *SafeKeepingPlace2 {
	q.SafekeepingPlace = new(SafeKeepingPlace2)
	return q.SafekeepingPlace
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount59 struct {

	// Denomination of the security to be received or delivered.
	DenominationChoice *RestrictedFINXMax210Text `xml:"DnmtnChc,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification119 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification6Choice `xml:"CshAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown38 `xml:"QtyBrkdwn,omitempty"`
}

func (q *QuantityAndAccount59) SetDenominationChoice(value string) {
	q.DenominationChoice = (*RestrictedFINXMax210Text)(&value)
}

func (q *QuantityAndAccount59) AddAccountOwner() *PartyIdentification119 {
	q.AccountOwner = new(PartyIdentification119)
	return q.AccountOwner
}

func (q *QuantityAndAccount59) AddCashAccount() *CashAccountIdentification6Choice {
	q.CashAccount = new(CashAccountIdentification6Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount59) AddSafekeepingPlace() *SafeKeepingPlace2 {
	q.SafekeepingPlace = new(SafeKeepingPlace2)
	return q.SafekeepingPlace
}

func (q *QuantityAndAccount59) AddQuantityBreakdown() *QuantityBreakdown38 {
	newValue := new(QuantityBreakdown38)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount6 struct {

	// Quantity of financial instrument effectively settled.
	SettledQuantity *Quantity6Choice `xml:"SttldQty"`

	// Quantity of financial instrument previously settled.
	PreviouslySettledQuantity *FinancialInstrumentQuantity1Choice `xml:"PrevslySttldQty,omitempty"`

	// Quantity of financial instrument remaining to be settled.
	RemainingToBeSettledQuantity *FinancialInstrumentQuantity1Choice `xml:"RmngToBeSttldQty,omitempty"`

	// Amount of money previously settled.
	PreviouslySettledAmount *AmountAndDirection5 `xml:"PrevslySttldAmt,omitempty"`

	// Amount of money remaining to be settled.
	RemainingToBeSettledAmount *AmountAndDirection5 `xml:"RmngToBeSttldAmt,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount6) AddSettledQuantity() *Quantity6Choice {
	q.SettledQuantity = new(Quantity6Choice)
	return q.SettledQuantity
}

func (q *QuantityAndAccount6) AddPreviouslySettledQuantity() *FinancialInstrumentQuantity1Choice {
	q.PreviouslySettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.PreviouslySettledQuantity
}

func (q *QuantityAndAccount6) AddRemainingToBeSettledQuantity() *FinancialInstrumentQuantity1Choice {
	q.RemainingToBeSettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.RemainingToBeSettledQuantity
}

func (q *QuantityAndAccount6) AddPreviouslySettledAmount() *AmountAndDirection5 {
	q.PreviouslySettledAmount = new(AmountAndDirection5)
	return q.PreviouslySettledAmount
}

func (q *QuantityAndAccount6) AddRemainingToBeSettledAmount() *AmountAndDirection5 {
	q.RemainingToBeSettledAmount = new(AmountAndDirection5)
	return q.RemainingToBeSettledAmount
}

func (q *QuantityAndAccount6) AddAccountOwner() *PartyIdentification13Choice {
	q.AccountOwner = new(PartyIdentification13Choice)
	return q.AccountOwner
}

func (q *QuantityAndAccount6) AddSafekeepingAccount() *SecuritiesAccount13 {
	q.SafekeepingAccount = new(SecuritiesAccount13)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount6) AddCashAccount() *CashAccountIdentification5Choice {
	q.CashAccount = new(CashAccountIdentification5Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount6) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	q.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return q.SafekeepingPlace
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount60 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity10Choice `xml:"SttlmQty,omitempty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *RestrictedFINXMax210Text `xml:"DnmtnChc,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification119 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount30 `xml:"SfkpgAcct,omitempty"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification6Choice `xml:"CshAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown38 `xml:"QtyBrkdwn,omitempty"`
}

func (q *QuantityAndAccount60) AddSettlementQuantity() *Quantity10Choice {
	q.SettlementQuantity = new(Quantity10Choice)
	return q.SettlementQuantity
}

func (q *QuantityAndAccount60) SetDenominationChoice(value string) {
	q.DenominationChoice = (*RestrictedFINXMax210Text)(&value)
}

func (q *QuantityAndAccount60) AddAccountOwner() *PartyIdentification119 {
	q.AccountOwner = new(PartyIdentification119)
	return q.AccountOwner
}

func (q *QuantityAndAccount60) AddSafekeepingAccount() *SecuritiesAccount30 {
	q.SafekeepingAccount = new(SecuritiesAccount30)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount60) AddCashAccount() *CashAccountIdentification6Choice {
	q.CashAccount = new(CashAccountIdentification6Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount60) AddSafekeepingPlace() *SafeKeepingPlace2 {
	q.SafekeepingPlace = new(SafeKeepingPlace2)
	return q.SafekeepingPlace
}

func (q *QuantityAndAccount60) AddQuantityBreakdown() *QuantityBreakdown38 {
	newValue := new(QuantityBreakdown38)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount61 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity15Choice `xml:"SttlmQty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification119 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount30 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification6Choice `xml:"CshAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount61) AddSettlementQuantity() *FinancialInstrumentQuantity15Choice {
	q.SettlementQuantity = new(FinancialInstrumentQuantity15Choice)
	return q.SettlementQuantity
}

func (q *QuantityAndAccount61) AddAccountOwner() *PartyIdentification119 {
	q.AccountOwner = new(PartyIdentification119)
	return q.AccountOwner
}

func (q *QuantityAndAccount61) AddSafekeepingAccount() *SecuritiesAccount30 {
	q.SafekeepingAccount = new(SecuritiesAccount30)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount61) AddCashAccount() *CashAccountIdentification6Choice {
	q.CashAccount = new(CashAccountIdentification6Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount61) AddSafekeepingPlace() *SafeKeepingPlace2 {
	q.SafekeepingPlace = new(SafeKeepingPlace2)
	return q.SafekeepingPlace
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount62 struct {

	// Quantity of financial instrument effectively settled.
	SettledQuantity *Quantity10Choice `xml:"SttldQty"`

	// Quantity of financial instrument previously settled.
	PreviouslySettledQuantity *FinancialInstrumentQuantity15Choice `xml:"PrevslySttldQty,omitempty"`

	// Quantity of financial instrument remaining to be settled.
	RemainingToBeSettledQuantity *FinancialInstrumentQuantity15Choice `xml:"RmngToBeSttldQty,omitempty"`

	// Amount of money previously settled.
	PreviouslySettledAmount *AmountAndDirection19 `xml:"PrevslySttldAmt,omitempty"`

	// Amount of money remaining to be settled.
	RemainingToBeSettledAmount *AmountAndDirection19 `xml:"RmngToBeSttldAmt,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification119 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount30 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification6Choice `xml:"CshAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount62) AddSettledQuantity() *Quantity10Choice {
	q.SettledQuantity = new(Quantity10Choice)
	return q.SettledQuantity
}

func (q *QuantityAndAccount62) AddPreviouslySettledQuantity() *FinancialInstrumentQuantity15Choice {
	q.PreviouslySettledQuantity = new(FinancialInstrumentQuantity15Choice)
	return q.PreviouslySettledQuantity
}

func (q *QuantityAndAccount62) AddRemainingToBeSettledQuantity() *FinancialInstrumentQuantity15Choice {
	q.RemainingToBeSettledQuantity = new(FinancialInstrumentQuantity15Choice)
	return q.RemainingToBeSettledQuantity
}

func (q *QuantityAndAccount62) AddPreviouslySettledAmount() *AmountAndDirection19 {
	q.PreviouslySettledAmount = new(AmountAndDirection19)
	return q.PreviouslySettledAmount
}

func (q *QuantityAndAccount62) AddRemainingToBeSettledAmount() *AmountAndDirection19 {
	q.RemainingToBeSettledAmount = new(AmountAndDirection19)
	return q.RemainingToBeSettledAmount
}

func (q *QuantityAndAccount62) AddAccountOwner() *PartyIdentification119 {
	q.AccountOwner = new(PartyIdentification119)
	return q.AccountOwner
}

func (q *QuantityAndAccount62) AddSafekeepingAccount() *SecuritiesAccount30 {
	q.SafekeepingAccount = new(SecuritiesAccount30)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount62) AddCashAccount() *CashAccountIdentification6Choice {
	q.CashAccount = new(CashAccountIdentification6Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount62) AddSafekeepingPlace() *SafeKeepingPlace2 {
	q.SafekeepingPlace = new(SafeKeepingPlace2)
	return q.SafekeepingPlace
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount7 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity1Choice `xml:"SttlmQty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount7) AddSettlementQuantity() *FinancialInstrumentQuantity1Choice {
	q.SettlementQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.SettlementQuantity
}

func (q *QuantityAndAccount7) AddAccountOwner() *PartyIdentification13Choice {
	q.AccountOwner = new(PartyIdentification13Choice)
	return q.AccountOwner
}

func (q *QuantityAndAccount7) AddSafekeepingAccount() *SecuritiesAccount13 {
	q.SafekeepingAccount = new(SecuritiesAccount13)
	return q.SafekeepingAccount
}

func (q *QuantityAndAccount7) AddCashAccount() *CashAccountIdentification5Choice {
	q.CashAccount = new(CashAccountIdentification5Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount7) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	q.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return q.SafekeepingPlace
}

// Details on the quantity, account and other related information involved in a transaction.
type QuantityAndAccount8 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity1Choice `xml:"SttlmQty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *Max210Text `xml:"DnmtnChc,omitempty"`

	// Unique and unambiguous identifier of a certificate assigned by the issuer.
	CertificateNumber []*SecuritiesCertificate1 `xml:"CertNb,omitempty"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown3 `xml:"QtyBrkdwn,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`
}

func (q *QuantityAndAccount8) AddSettlementQuantity() *FinancialInstrumentQuantity1Choice {
	q.SettlementQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.SettlementQuantity
}

func (q *QuantityAndAccount8) SetDenominationChoice(value string) {
	q.DenominationChoice = (*Max210Text)(&value)
}

func (q *QuantityAndAccount8) AddCertificateNumber() *SecuritiesCertificate1 {
	newValue := new(SecuritiesCertificate1)
	q.CertificateNumber = append(q.CertificateNumber, newValue)
	return newValue
}

func (q *QuantityAndAccount8) AddCashAccount() *CashAccountIdentification5Choice {
	q.CashAccount = new(CashAccountIdentification5Choice)
	return q.CashAccount
}

func (q *QuantityAndAccount8) AddQuantityBreakdown() *QuantityBreakdown3 {
	newValue := new(QuantityBreakdown3)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

func (q *QuantityAndAccount8) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	q.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return q.SafekeepingPlace
}
