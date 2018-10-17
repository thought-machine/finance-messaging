package iso20022

// Eligible and not eligible balance of securities for a corporate action event.
type CorporateActionBalanceDetails1 struct {

	// Total quantity of financial instruments of the balance.
	TotalEligibleBalance *Quantity3Choice `xml:"TtlElgblBal,omitempty"`

	// Balance of financial instruments that are blocked.
	BlockedBalance *BalanceFormat1Choice `xml:"BlckdBal,omitempty"`

	// Quantity of securities in the sub-balance.
	BorrowedBalance *BalanceFormat1Choice `xml:"BrrwdBal,omitempty"`

	// Quantity of securities in the sub-balance.
	CollateralInBalance *BalanceFormat1Choice `xml:"CollInBal,omitempty"`

	// Quantity of securities in the sub-balance.
	CollateralOutBalance *BalanceFormat1Choice `xml:"CollOutBal,omitempty"`

	// Quantity of securities in the sub-balance.
	OnLoanBalance *BalanceFormat1Choice `xml:"OnLnBal,omitempty"`

	// Quantity of securities in the sub-balance.
	PendingDeliveryBalance []*BalanceFormat1Choice `xml:"PdgDlvryBal,omitempty"`

	// Quantity of securities in the sub-balance.
	PendingReceiptBalance []*BalanceFormat1Choice `xml:"PdgRctBal,omitempty"`

	// Quantity of securities in the sub-balance.
	OutForRegistrationBalance *BalanceFormat1Choice `xml:"OutForRegnBal,omitempty"`

	// Quantity of securities in the sub-balance.
	SettlementPositionBalance *BalanceFormat1Choice `xml:"SttlmPosBal,omitempty"`

	// Quantity of securities in the sub-balance.
	StreetPositionBalance *BalanceFormat1Choice `xml:"StrtPosBal,omitempty"`

	// Balance of securities based on trade date, for example, includes all pending transactions in addition to the balance of settled transactions.
	TradeDatePositionBalance *BalanceFormat1Choice `xml:"TradDtPosBal,omitempty"`

	// Quantity of securities in the sub-balance.
	InTransshipmentBalance *BalanceFormat1Choice `xml:"InTrnsShipmntBal,omitempty"`

	// Quantity of securities in the sub-balance.
	RegisteredBalance *BalanceFormat1Choice `xml:"RegdBal,omitempty"`

	// Position that account holders should return to the account servicer to participate in the event or to fulfil their obligation for the event to be complete, for example, return of securities for late announced drawing.
	ObligatedBalance *BalanceFormat1Choice `xml:"OblgtdBal,omitempty"`

	// Balance of uninstructed position.
	UninstructedBalance *BalanceFormat1Choice `xml:"UinstdBal,omitempty"`

	// Balance of instructed position.
	InstructedBalance *BalanceFormat1Choice `xml:"InstdBal,omitempty"`

	// Balance that has been affected by the process run through the event.
	AffectedBalance *BalanceFormat1Choice `xml:"AfctdBal,omitempty"`

	// Balance that has not been affected by the process run through the event.
	UnaffectedBalance *BalanceFormat1Choice `xml:"UafctdBal,omitempty"`
}

func (c *CorporateActionBalanceDetails1) AddTotalEligibleBalance() *Quantity3Choice {
	c.TotalEligibleBalance = new(Quantity3Choice)
	return c.TotalEligibleBalance
}

func (c *CorporateActionBalanceDetails1) AddBlockedBalance() *BalanceFormat1Choice {
	c.BlockedBalance = new(BalanceFormat1Choice)
	return c.BlockedBalance
}

func (c *CorporateActionBalanceDetails1) AddBorrowedBalance() *BalanceFormat1Choice {
	c.BorrowedBalance = new(BalanceFormat1Choice)
	return c.BorrowedBalance
}

func (c *CorporateActionBalanceDetails1) AddCollateralInBalance() *BalanceFormat1Choice {
	c.CollateralInBalance = new(BalanceFormat1Choice)
	return c.CollateralInBalance
}

func (c *CorporateActionBalanceDetails1) AddCollateralOutBalance() *BalanceFormat1Choice {
	c.CollateralOutBalance = new(BalanceFormat1Choice)
	return c.CollateralOutBalance
}

func (c *CorporateActionBalanceDetails1) AddOnLoanBalance() *BalanceFormat1Choice {
	c.OnLoanBalance = new(BalanceFormat1Choice)
	return c.OnLoanBalance
}

func (c *CorporateActionBalanceDetails1) AddPendingDeliveryBalance() *BalanceFormat1Choice {
	newValue := new(BalanceFormat1Choice)
	c.PendingDeliveryBalance = append(c.PendingDeliveryBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails1) AddPendingReceiptBalance() *BalanceFormat1Choice {
	newValue := new(BalanceFormat1Choice)
	c.PendingReceiptBalance = append(c.PendingReceiptBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails1) AddOutForRegistrationBalance() *BalanceFormat1Choice {
	c.OutForRegistrationBalance = new(BalanceFormat1Choice)
	return c.OutForRegistrationBalance
}

func (c *CorporateActionBalanceDetails1) AddSettlementPositionBalance() *BalanceFormat1Choice {
	c.SettlementPositionBalance = new(BalanceFormat1Choice)
	return c.SettlementPositionBalance
}

func (c *CorporateActionBalanceDetails1) AddStreetPositionBalance() *BalanceFormat1Choice {
	c.StreetPositionBalance = new(BalanceFormat1Choice)
	return c.StreetPositionBalance
}

func (c *CorporateActionBalanceDetails1) AddTradeDatePositionBalance() *BalanceFormat1Choice {
	c.TradeDatePositionBalance = new(BalanceFormat1Choice)
	return c.TradeDatePositionBalance
}

func (c *CorporateActionBalanceDetails1) AddInTransshipmentBalance() *BalanceFormat1Choice {
	c.InTransshipmentBalance = new(BalanceFormat1Choice)
	return c.InTransshipmentBalance
}

func (c *CorporateActionBalanceDetails1) AddRegisteredBalance() *BalanceFormat1Choice {
	c.RegisteredBalance = new(BalanceFormat1Choice)
	return c.RegisteredBalance
}

func (c *CorporateActionBalanceDetails1) AddObligatedBalance() *BalanceFormat1Choice {
	c.ObligatedBalance = new(BalanceFormat1Choice)
	return c.ObligatedBalance
}

func (c *CorporateActionBalanceDetails1) AddUninstructedBalance() *BalanceFormat1Choice {
	c.UninstructedBalance = new(BalanceFormat1Choice)
	return c.UninstructedBalance
}

func (c *CorporateActionBalanceDetails1) AddInstructedBalance() *BalanceFormat1Choice {
	c.InstructedBalance = new(BalanceFormat1Choice)
	return c.InstructedBalance
}

func (c *CorporateActionBalanceDetails1) AddAffectedBalance() *BalanceFormat1Choice {
	c.AffectedBalance = new(BalanceFormat1Choice)
	return c.AffectedBalance
}

func (c *CorporateActionBalanceDetails1) AddUnaffectedBalance() *BalanceFormat1Choice {
	c.UnaffectedBalance = new(BalanceFormat1Choice)
	return c.UnaffectedBalance
}

// Eligible and not eligible balance of securities for a corporate action event.
type CorporateActionBalanceDetails10 struct {

	// Total quantity of financial instruments of the balance.
	TotalEligibleBalance *Quantity3Choice `xml:"TtlElgblBal,omitempty"`

	// Balance of financial instruments that are blocked.
	BlockedBalance *BalanceFormat1Choice `xml:"BlckdBal,omitempty"`

	// Balance of financial instruments that have been borrowed from another party.
	BorrowedBalance *BalanceFormat1Choice `xml:"BrrwdBal,omitempty"`

	// Balance of securities that belong to a third party and that are held for the purpose of collateralisation.
	CollateralInBalance *BalanceFormat1Choice `xml:"CollInBal,omitempty"`

	// Balance of securities that belong to the safekeeping account indicated within this message, and are deposited with a third party for the purpose of collateralisation.
	CollateralOutBalance *BalanceFormat1Choice `xml:"CollOutBal,omitempty"`

	// Balance of financial instruments that have been loaned to a third party.
	OnLoanBalance *BalanceFormat1Choice `xml:"OnLnBal,omitempty"`

	// Balance of financial instruments that are pending delivery.
	PendingDeliveryBalance []*BalanceFormat1Choice `xml:"PdgDlvryBal,omitempty"`

	// Balance of financial instruments that are pending receipt.
	PendingReceiptBalance []*BalanceFormat1Choice `xml:"PdgRctBal,omitempty"`

	// Balance of financial instruments currently being processed by the institution responsible for registering the new beneficial owner (or nominee).
	OutForRegistrationBalance *BalanceFormat1Choice `xml:"OutForRegnBal,omitempty"`

	// Balance of securities representing only settled transactions; pending transactions not included.
	SettlementPositionBalance *BalanceFormat1Choice `xml:"SttlmPosBal,omitempty"`

	// Balance of financial instruments that remain registered in the name of the prior beneficial owner.
	StreetPositionBalance *BalanceFormat1Choice `xml:"StrtPosBal,omitempty"`

	// Balance of securities based on trade date, for example, includes all pending transactions in addition to the balance of settled transactions.
	TradeDatePositionBalance *BalanceFormat1Choice `xml:"TradDtPosBal,omitempty"`

	// Balance of physical securities that are in the process of being transferred from one depository/agent to another.
	InTransshipmentBalance *BalanceFormat1Choice `xml:"InTrnsShipmntBal,omitempty"`

	// Balance of financial instruments that are registered (in the name of a nominee name or of the beneficial owner).
	RegisteredBalance *BalanceFormat1Choice `xml:"RegdBal,omitempty"`

	// Position that account holders should return to the account servicer to participate in the event or to fulfil their obligation for the event to be complete, for example, return of securities for late announced drawing.
	ObligatedBalance *BalanceFormat1Choice `xml:"OblgtdBal,omitempty"`

	// Balance of uninstructed position.
	UninstructedBalance *BalanceFormat1Choice `xml:"UinstdBal,omitempty"`

	// Balance of instructed position.
	InstructedBalance *BalanceFormat1Choice `xml:"InstdBal,omitempty"`

	// Balance that has been affected by the process run through the event.
	AffectedBalance *BalanceFormat1Choice `xml:"AfctdBal,omitempty"`

	// Balance that has not been affected by the process run through the event.
	UnaffectedBalance *BalanceFormat1Choice `xml:"UafctdBal,omitempty"`
}

func (c *CorporateActionBalanceDetails10) AddTotalEligibleBalance() *Quantity3Choice {
	c.TotalEligibleBalance = new(Quantity3Choice)
	return c.TotalEligibleBalance
}

func (c *CorporateActionBalanceDetails10) AddBlockedBalance() *BalanceFormat1Choice {
	c.BlockedBalance = new(BalanceFormat1Choice)
	return c.BlockedBalance
}

func (c *CorporateActionBalanceDetails10) AddBorrowedBalance() *BalanceFormat1Choice {
	c.BorrowedBalance = new(BalanceFormat1Choice)
	return c.BorrowedBalance
}

func (c *CorporateActionBalanceDetails10) AddCollateralInBalance() *BalanceFormat1Choice {
	c.CollateralInBalance = new(BalanceFormat1Choice)
	return c.CollateralInBalance
}

func (c *CorporateActionBalanceDetails10) AddCollateralOutBalance() *BalanceFormat1Choice {
	c.CollateralOutBalance = new(BalanceFormat1Choice)
	return c.CollateralOutBalance
}

func (c *CorporateActionBalanceDetails10) AddOnLoanBalance() *BalanceFormat1Choice {
	c.OnLoanBalance = new(BalanceFormat1Choice)
	return c.OnLoanBalance
}

func (c *CorporateActionBalanceDetails10) AddPendingDeliveryBalance() *BalanceFormat1Choice {
	newValue := new(BalanceFormat1Choice)
	c.PendingDeliveryBalance = append(c.PendingDeliveryBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails10) AddPendingReceiptBalance() *BalanceFormat1Choice {
	newValue := new(BalanceFormat1Choice)
	c.PendingReceiptBalance = append(c.PendingReceiptBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails10) AddOutForRegistrationBalance() *BalanceFormat1Choice {
	c.OutForRegistrationBalance = new(BalanceFormat1Choice)
	return c.OutForRegistrationBalance
}

func (c *CorporateActionBalanceDetails10) AddSettlementPositionBalance() *BalanceFormat1Choice {
	c.SettlementPositionBalance = new(BalanceFormat1Choice)
	return c.SettlementPositionBalance
}

func (c *CorporateActionBalanceDetails10) AddStreetPositionBalance() *BalanceFormat1Choice {
	c.StreetPositionBalance = new(BalanceFormat1Choice)
	return c.StreetPositionBalance
}

func (c *CorporateActionBalanceDetails10) AddTradeDatePositionBalance() *BalanceFormat1Choice {
	c.TradeDatePositionBalance = new(BalanceFormat1Choice)
	return c.TradeDatePositionBalance
}

func (c *CorporateActionBalanceDetails10) AddInTransshipmentBalance() *BalanceFormat1Choice {
	c.InTransshipmentBalance = new(BalanceFormat1Choice)
	return c.InTransshipmentBalance
}

func (c *CorporateActionBalanceDetails10) AddRegisteredBalance() *BalanceFormat1Choice {
	c.RegisteredBalance = new(BalanceFormat1Choice)
	return c.RegisteredBalance
}

func (c *CorporateActionBalanceDetails10) AddObligatedBalance() *BalanceFormat1Choice {
	c.ObligatedBalance = new(BalanceFormat1Choice)
	return c.ObligatedBalance
}

func (c *CorporateActionBalanceDetails10) AddUninstructedBalance() *BalanceFormat1Choice {
	c.UninstructedBalance = new(BalanceFormat1Choice)
	return c.UninstructedBalance
}

func (c *CorporateActionBalanceDetails10) AddInstructedBalance() *BalanceFormat1Choice {
	c.InstructedBalance = new(BalanceFormat1Choice)
	return c.InstructedBalance
}

func (c *CorporateActionBalanceDetails10) AddAffectedBalance() *BalanceFormat1Choice {
	c.AffectedBalance = new(BalanceFormat1Choice)
	return c.AffectedBalance
}

func (c *CorporateActionBalanceDetails10) AddUnaffectedBalance() *BalanceFormat1Choice {
	c.UnaffectedBalance = new(BalanceFormat1Choice)
	return c.UnaffectedBalance
}

// Eligible and not eligible balance of securities for a corporate action event.
type CorporateActionBalanceDetails11 struct {

	// Balance to which the payment applies (less or equal to the total eligible balance).
	ConfirmedBalance *BalanceFormat1Choice `xml:"ConfdBal"`

	// Total quantity of financial instruments of the balance.
	TotalEligibleBalance *Quantity3Choice `xml:"TtlElgblBal,omitempty"`

	// Balance of financial instruments that are blocked.
	BlockedBalance *BalanceFormat1Choice `xml:"BlckdBal,omitempty"`

	// Balance of financial instruments that have been borrowed from another party.
	BorrowedBalance *BalanceFormat1Choice `xml:"BrrwdBal,omitempty"`

	// Balance of securities that belong to a third party and that are held for the purpose of collateralisation.
	CollateralInBalance *BalanceFormat1Choice `xml:"CollInBal,omitempty"`

	// Balance of securities that belong to the safekeeping account indicated within this message, and are deposited with a third party for the purpose of collateralisation.
	CollateralOutBalance *BalanceFormat1Choice `xml:"CollOutBal,omitempty"`

	// Balance of financial instruments that have been loaned to a third party.
	OnLoanBalance *BalanceFormat1Choice `xml:"OnLnBal,omitempty"`

	// Balance of financial instruments that are pending delivery.
	PendingDeliveryBalance []*BalanceFormat1Choice `xml:"PdgDlvryBal,omitempty"`

	// Balance of financial instruments that are pending receipt.
	PendingReceiptBalance []*BalanceFormat1Choice `xml:"PdgRctBal,omitempty"`

	// Balance of financial instruments currently being processed by the institution responsible for registering the new beneficial owner (or nominee).
	OutForRegistrationBalance *BalanceFormat1Choice `xml:"OutForRegnBal,omitempty"`

	// Balance of securities representing only settled transactions; pending transactions not included.
	SettlementPositionBalance *BalanceFormat1Choice `xml:"SttlmPosBal,omitempty"`

	// Balance of financial instruments that remain registered in the name of the prior beneficial owner.
	StreetPositionBalance *BalanceFormat1Choice `xml:"StrtPosBal,omitempty"`

	// Balance of securities based on trade date, for example, includes all pending transactions in addition to the balance of settled transactions.
	TradeDatePositionBalance *BalanceFormat1Choice `xml:"TradDtPosBal,omitempty"`

	// Balance of physical securities that are in the process of being transferred from one depository/agent to another.
	InTransshipmentBalance *BalanceFormat1Choice `xml:"InTrnsShipmntBal,omitempty"`

	// Balance of financial instruments that are registered (in the name of a nominee name or of the beneficial owner).
	RegisteredBalance *BalanceFormat1Choice `xml:"RegdBal,omitempty"`

	// Balance that has been affected by the process run through the event.
	AffectedBalance *BalanceFormat1Choice `xml:"AfctdBal,omitempty"`

	// Balance that has not been affected by the process run through the event.
	UnaffectedBalance *BalanceFormat1Choice `xml:"UafctdBal,omitempty"`
}

func (c *CorporateActionBalanceDetails11) AddConfirmedBalance() *BalanceFormat1Choice {
	c.ConfirmedBalance = new(BalanceFormat1Choice)
	return c.ConfirmedBalance
}

func (c *CorporateActionBalanceDetails11) AddTotalEligibleBalance() *Quantity3Choice {
	c.TotalEligibleBalance = new(Quantity3Choice)
	return c.TotalEligibleBalance
}

func (c *CorporateActionBalanceDetails11) AddBlockedBalance() *BalanceFormat1Choice {
	c.BlockedBalance = new(BalanceFormat1Choice)
	return c.BlockedBalance
}

func (c *CorporateActionBalanceDetails11) AddBorrowedBalance() *BalanceFormat1Choice {
	c.BorrowedBalance = new(BalanceFormat1Choice)
	return c.BorrowedBalance
}

func (c *CorporateActionBalanceDetails11) AddCollateralInBalance() *BalanceFormat1Choice {
	c.CollateralInBalance = new(BalanceFormat1Choice)
	return c.CollateralInBalance
}

func (c *CorporateActionBalanceDetails11) AddCollateralOutBalance() *BalanceFormat1Choice {
	c.CollateralOutBalance = new(BalanceFormat1Choice)
	return c.CollateralOutBalance
}

func (c *CorporateActionBalanceDetails11) AddOnLoanBalance() *BalanceFormat1Choice {
	c.OnLoanBalance = new(BalanceFormat1Choice)
	return c.OnLoanBalance
}

func (c *CorporateActionBalanceDetails11) AddPendingDeliveryBalance() *BalanceFormat1Choice {
	newValue := new(BalanceFormat1Choice)
	c.PendingDeliveryBalance = append(c.PendingDeliveryBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails11) AddPendingReceiptBalance() *BalanceFormat1Choice {
	newValue := new(BalanceFormat1Choice)
	c.PendingReceiptBalance = append(c.PendingReceiptBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails11) AddOutForRegistrationBalance() *BalanceFormat1Choice {
	c.OutForRegistrationBalance = new(BalanceFormat1Choice)
	return c.OutForRegistrationBalance
}

func (c *CorporateActionBalanceDetails11) AddSettlementPositionBalance() *BalanceFormat1Choice {
	c.SettlementPositionBalance = new(BalanceFormat1Choice)
	return c.SettlementPositionBalance
}

func (c *CorporateActionBalanceDetails11) AddStreetPositionBalance() *BalanceFormat1Choice {
	c.StreetPositionBalance = new(BalanceFormat1Choice)
	return c.StreetPositionBalance
}

func (c *CorporateActionBalanceDetails11) AddTradeDatePositionBalance() *BalanceFormat1Choice {
	c.TradeDatePositionBalance = new(BalanceFormat1Choice)
	return c.TradeDatePositionBalance
}

func (c *CorporateActionBalanceDetails11) AddInTransshipmentBalance() *BalanceFormat1Choice {
	c.InTransshipmentBalance = new(BalanceFormat1Choice)
	return c.InTransshipmentBalance
}

func (c *CorporateActionBalanceDetails11) AddRegisteredBalance() *BalanceFormat1Choice {
	c.RegisteredBalance = new(BalanceFormat1Choice)
	return c.RegisteredBalance
}

func (c *CorporateActionBalanceDetails11) AddAffectedBalance() *BalanceFormat1Choice {
	c.AffectedBalance = new(BalanceFormat1Choice)
	return c.AffectedBalance
}

func (c *CorporateActionBalanceDetails11) AddUnaffectedBalance() *BalanceFormat1Choice {
	c.UnaffectedBalance = new(BalanceFormat1Choice)
	return c.UnaffectedBalance
}

// Eligible and not eligible balance of securities for a corporate action event.
type CorporateActionBalanceDetails12 struct {

	// Total quantity of financial instruments of the balance.
	TotalEligibleBalance *Quantity3Choice `xml:"TtlElgblBal,omitempty"`

	// Balance of financial instruments that are blocked.
	BlockedBalance *BalanceFormat1Choice `xml:"BlckdBal,omitempty"`

	// Balance of financial instruments that have been borrowed from another party.
	BorrowedBalance *BalanceFormat1Choice `xml:"BrrwdBal,omitempty"`

	// Balance of securities that belong to a third party and that are held for the purpose of collateralisation.
	CollateralInBalance *BalanceFormat1Choice `xml:"CollInBal,omitempty"`

	// Balance of securities that belong to the safekeeping account indicated within this message, and are deposited with a third party for the purpose of collateralisation.
	CollateralOutBalance *BalanceFormat1Choice `xml:"CollOutBal,omitempty"`

	// Balance of financial instruments that have been loaned to a third party.
	OnLoanBalance *BalanceFormat1Choice `xml:"OnLnBal,omitempty"`

	// Balance of financial instruments that are pending delivery.
	PendingDeliveryBalance []*BalanceFormat1Choice `xml:"PdgDlvryBal,omitempty"`

	// Balance of financial instruments that are pending receipt.
	PendingReceiptBalance []*BalanceFormat1Choice `xml:"PdgRctBal,omitempty"`

	// Balance of financial instruments currently being processed by the institution responsible for registering the new beneficial owner (or nominee).
	OutForRegistrationBalance *BalanceFormat1Choice `xml:"OutForRegnBal,omitempty"`

	// Balance of securities representing only settled transactions; pending transactions not included.
	SettlementPositionBalance *BalanceFormat1Choice `xml:"SttlmPosBal,omitempty"`

	// Balance of financial instruments that remain registered in the name of the prior beneficial owner.
	StreetPositionBalance *BalanceFormat1Choice `xml:"StrtPosBal,omitempty"`

	// Balance of securities based on trade date, for example, includes all pending transactions in addition to the balance of settled transactions.
	TradeDatePositionBalance *BalanceFormat1Choice `xml:"TradDtPosBal,omitempty"`

	// Balance of physical securities that are in the process of being transferred from one depository/agent to another.
	InTransshipmentBalance *BalanceFormat1Choice `xml:"InTrnsShipmntBal,omitempty"`

	// Balance of financial instruments that are registered (in the name of a nominee name or of the beneficial owner).
	RegisteredBalance *BalanceFormat1Choice `xml:"RegdBal,omitempty"`
}

func (c *CorporateActionBalanceDetails12) AddTotalEligibleBalance() *Quantity3Choice {
	c.TotalEligibleBalance = new(Quantity3Choice)
	return c.TotalEligibleBalance
}

func (c *CorporateActionBalanceDetails12) AddBlockedBalance() *BalanceFormat1Choice {
	c.BlockedBalance = new(BalanceFormat1Choice)
	return c.BlockedBalance
}

func (c *CorporateActionBalanceDetails12) AddBorrowedBalance() *BalanceFormat1Choice {
	c.BorrowedBalance = new(BalanceFormat1Choice)
	return c.BorrowedBalance
}

func (c *CorporateActionBalanceDetails12) AddCollateralInBalance() *BalanceFormat1Choice {
	c.CollateralInBalance = new(BalanceFormat1Choice)
	return c.CollateralInBalance
}

func (c *CorporateActionBalanceDetails12) AddCollateralOutBalance() *BalanceFormat1Choice {
	c.CollateralOutBalance = new(BalanceFormat1Choice)
	return c.CollateralOutBalance
}

func (c *CorporateActionBalanceDetails12) AddOnLoanBalance() *BalanceFormat1Choice {
	c.OnLoanBalance = new(BalanceFormat1Choice)
	return c.OnLoanBalance
}

func (c *CorporateActionBalanceDetails12) AddPendingDeliveryBalance() *BalanceFormat1Choice {
	newValue := new(BalanceFormat1Choice)
	c.PendingDeliveryBalance = append(c.PendingDeliveryBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails12) AddPendingReceiptBalance() *BalanceFormat1Choice {
	newValue := new(BalanceFormat1Choice)
	c.PendingReceiptBalance = append(c.PendingReceiptBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails12) AddOutForRegistrationBalance() *BalanceFormat1Choice {
	c.OutForRegistrationBalance = new(BalanceFormat1Choice)
	return c.OutForRegistrationBalance
}

func (c *CorporateActionBalanceDetails12) AddSettlementPositionBalance() *BalanceFormat1Choice {
	c.SettlementPositionBalance = new(BalanceFormat1Choice)
	return c.SettlementPositionBalance
}

func (c *CorporateActionBalanceDetails12) AddStreetPositionBalance() *BalanceFormat1Choice {
	c.StreetPositionBalance = new(BalanceFormat1Choice)
	return c.StreetPositionBalance
}

func (c *CorporateActionBalanceDetails12) AddTradeDatePositionBalance() *BalanceFormat1Choice {
	c.TradeDatePositionBalance = new(BalanceFormat1Choice)
	return c.TradeDatePositionBalance
}

func (c *CorporateActionBalanceDetails12) AddInTransshipmentBalance() *BalanceFormat1Choice {
	c.InTransshipmentBalance = new(BalanceFormat1Choice)
	return c.InTransshipmentBalance
}

func (c *CorporateActionBalanceDetails12) AddRegisteredBalance() *BalanceFormat1Choice {
	c.RegisteredBalance = new(BalanceFormat1Choice)
	return c.RegisteredBalance
}

// Eligible and not eligible balance of securities for a corporate action event.
type CorporateActionBalanceDetails17 struct {

	// Total quantity of financial instruments of the balance.
	TotalEligibleBalance *Quantity3Choice `xml:"TtlElgblBal,omitempty"`

	// Balance of financial instruments that are blocked.
	BlockedBalance *BalanceFormat1Choice `xml:"BlckdBal,omitempty"`

	// Balance of financial instruments that have been borrowed from another party.
	BorrowedBalance *BalanceFormat1Choice `xml:"BrrwdBal,omitempty"`

	// Balance of securities that belong to a third party and that are held for the purpose of collateralisation.
	CollateralInBalance *BalanceFormat1Choice `xml:"CollInBal,omitempty"`

	// Balance of securities that belong to the safekeeping account indicated within this message, and are deposited with a third party for the purpose of collateralisation.
	CollateralOutBalance *BalanceFormat1Choice `xml:"CollOutBal,omitempty"`

	// Balance of financial instruments that have been loaned to a third party.
	OnLoanBalance *BalanceFormat1Choice `xml:"OnLnBal,omitempty"`

	// Balance of financial instruments that are pending delivery.
	PendingDeliveryBalance []*BalanceFormat3Choice `xml:"PdgDlvryBal,omitempty"`

	// Balance of financial instruments that are pending receipt.
	PendingReceiptBalance []*BalanceFormat3Choice `xml:"PdgRctBal,omitempty"`

	// Balance of financial instruments currently being processed by the institution responsible for registering the new beneficial owner (or nominee).
	OutForRegistrationBalance *BalanceFormat1Choice `xml:"OutForRegnBal,omitempty"`

	// Balance of securities representing only settled transactions; pending transactions not included.
	SettlementPositionBalance []*BalanceFormat3Choice `xml:"SttlmPosBal,omitempty"`

	// Balance of financial instruments that remain registered in the name of the prior beneficial owner.
	StreetPositionBalance *BalanceFormat1Choice `xml:"StrtPosBal,omitempty"`

	// Balance of securities based on trade date, for example, includes all pending transactions in addition to the balance of settled transactions.
	TradeDatePositionBalance *BalanceFormat1Choice `xml:"TradDtPosBal,omitempty"`

	// Balance of physical securities that are in the process of being transferred from one depository/agent to another.
	InTransshipmentBalance *BalanceFormat1Choice `xml:"InTrnsShipmntBal,omitempty"`

	// Balance of financial instruments that are registered (in the name of a nominee name or of the beneficial owner).
	RegisteredBalance *BalanceFormat1Choice `xml:"RegdBal,omitempty"`

	// Position that account holders should return to the account servicer to participate in the event or to fulfil their obligation for the event to be complete, for example, return of securities for late announced drawing.
	ObligatedBalance *BalanceFormat1Choice `xml:"OblgtdBal,omitempty"`

	// Balance of uninstructed position.
	UninstructedBalance *BalanceFormat1Choice `xml:"UinstdBal,omitempty"`

	// Balance of instructed position.
	InstructedBalance *BalanceFormat1Choice `xml:"InstdBal,omitempty"`

	// Balance that has been affected by the process run through the event.
	AffectedBalance *BalanceFormat1Choice `xml:"AfctdBal,omitempty"`

	// Balance that has not been affected by the process run through the event.
	UnaffectedBalance *BalanceFormat1Choice `xml:"UafctdBal,omitempty"`
}

func (c *CorporateActionBalanceDetails17) AddTotalEligibleBalance() *Quantity3Choice {
	c.TotalEligibleBalance = new(Quantity3Choice)
	return c.TotalEligibleBalance
}

func (c *CorporateActionBalanceDetails17) AddBlockedBalance() *BalanceFormat1Choice {
	c.BlockedBalance = new(BalanceFormat1Choice)
	return c.BlockedBalance
}

func (c *CorporateActionBalanceDetails17) AddBorrowedBalance() *BalanceFormat1Choice {
	c.BorrowedBalance = new(BalanceFormat1Choice)
	return c.BorrowedBalance
}

func (c *CorporateActionBalanceDetails17) AddCollateralInBalance() *BalanceFormat1Choice {
	c.CollateralInBalance = new(BalanceFormat1Choice)
	return c.CollateralInBalance
}

func (c *CorporateActionBalanceDetails17) AddCollateralOutBalance() *BalanceFormat1Choice {
	c.CollateralOutBalance = new(BalanceFormat1Choice)
	return c.CollateralOutBalance
}

func (c *CorporateActionBalanceDetails17) AddOnLoanBalance() *BalanceFormat1Choice {
	c.OnLoanBalance = new(BalanceFormat1Choice)
	return c.OnLoanBalance
}

func (c *CorporateActionBalanceDetails17) AddPendingDeliveryBalance() *BalanceFormat3Choice {
	newValue := new(BalanceFormat3Choice)
	c.PendingDeliveryBalance = append(c.PendingDeliveryBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails17) AddPendingReceiptBalance() *BalanceFormat3Choice {
	newValue := new(BalanceFormat3Choice)
	c.PendingReceiptBalance = append(c.PendingReceiptBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails17) AddOutForRegistrationBalance() *BalanceFormat1Choice {
	c.OutForRegistrationBalance = new(BalanceFormat1Choice)
	return c.OutForRegistrationBalance
}

func (c *CorporateActionBalanceDetails17) AddSettlementPositionBalance() *BalanceFormat3Choice {
	newValue := new(BalanceFormat3Choice)
	c.SettlementPositionBalance = append(c.SettlementPositionBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails17) AddStreetPositionBalance() *BalanceFormat1Choice {
	c.StreetPositionBalance = new(BalanceFormat1Choice)
	return c.StreetPositionBalance
}

func (c *CorporateActionBalanceDetails17) AddTradeDatePositionBalance() *BalanceFormat1Choice {
	c.TradeDatePositionBalance = new(BalanceFormat1Choice)
	return c.TradeDatePositionBalance
}

func (c *CorporateActionBalanceDetails17) AddInTransshipmentBalance() *BalanceFormat1Choice {
	c.InTransshipmentBalance = new(BalanceFormat1Choice)
	return c.InTransshipmentBalance
}

func (c *CorporateActionBalanceDetails17) AddRegisteredBalance() *BalanceFormat1Choice {
	c.RegisteredBalance = new(BalanceFormat1Choice)
	return c.RegisteredBalance
}

func (c *CorporateActionBalanceDetails17) AddObligatedBalance() *BalanceFormat1Choice {
	c.ObligatedBalance = new(BalanceFormat1Choice)
	return c.ObligatedBalance
}

func (c *CorporateActionBalanceDetails17) AddUninstructedBalance() *BalanceFormat1Choice {
	c.UninstructedBalance = new(BalanceFormat1Choice)
	return c.UninstructedBalance
}

func (c *CorporateActionBalanceDetails17) AddInstructedBalance() *BalanceFormat1Choice {
	c.InstructedBalance = new(BalanceFormat1Choice)
	return c.InstructedBalance
}

func (c *CorporateActionBalanceDetails17) AddAffectedBalance() *BalanceFormat1Choice {
	c.AffectedBalance = new(BalanceFormat1Choice)
	return c.AffectedBalance
}

func (c *CorporateActionBalanceDetails17) AddUnaffectedBalance() *BalanceFormat1Choice {
	c.UnaffectedBalance = new(BalanceFormat1Choice)
	return c.UnaffectedBalance
}

// Eligible and not eligible balance of securities for a corporate action event.
type CorporateActionBalanceDetails18 struct {

	// Balance to which the payment applies (less or equal to the total eligible balance).
	ConfirmedBalance *BalanceFormat1Choice `xml:"ConfdBal"`

	// Total quantity of financial instruments of the balance.
	TotalEligibleBalance *Quantity3Choice `xml:"TtlElgblBal,omitempty"`

	// Balance of financial instruments that are blocked.
	BlockedBalance *BalanceFormat1Choice `xml:"BlckdBal,omitempty"`

	// Balance of financial instruments that have been borrowed from another party.
	BorrowedBalance *BalanceFormat1Choice `xml:"BrrwdBal,omitempty"`

	// Balance of securities that belong to a third party and that are held for the purpose of collateralisation.
	CollateralInBalance *BalanceFormat1Choice `xml:"CollInBal,omitempty"`

	// Balance of securities that belong to the safekeeping account indicated within this message, and are deposited with a third party for the purpose of collateralisation.
	CollateralOutBalance *BalanceFormat1Choice `xml:"CollOutBal,omitempty"`

	// Balance of financial instruments that have been loaned to a third party.
	OnLoanBalance *BalanceFormat1Choice `xml:"OnLnBal,omitempty"`

	// Balance of financial instruments that are pending delivery.
	PendingDeliveryBalance []*BalanceFormat3Choice `xml:"PdgDlvryBal,omitempty"`

	// Balance of financial instruments that are pending receipt.
	PendingReceiptBalance []*BalanceFormat3Choice `xml:"PdgRctBal,omitempty"`

	// Balance of financial instruments currently being processed by the institution responsible for registering the new beneficial owner (or nominee).
	OutForRegistrationBalance *BalanceFormat1Choice `xml:"OutForRegnBal,omitempty"`

	// Balance of securities representing only settled transactions; pending transactions not included.
	SettlementPositionBalance []*BalanceFormat3Choice `xml:"SttlmPosBal,omitempty"`

	// Balance of financial instruments that remain registered in the name of the prior beneficial owner.
	StreetPositionBalance *BalanceFormat1Choice `xml:"StrtPosBal,omitempty"`

	// Balance of securities based on trade date, for example, includes all pending transactions in addition to the balance of settled transactions.
	TradeDatePositionBalance *BalanceFormat1Choice `xml:"TradDtPosBal,omitempty"`

	// Balance of physical securities that are in the process of being transferred from one depository/agent to another.
	InTransshipmentBalance *BalanceFormat1Choice `xml:"InTrnsShipmntBal,omitempty"`

	// Balance of financial instruments that are registered (in the name of a nominee name or of the beneficial owner).
	RegisteredBalance *BalanceFormat1Choice `xml:"RegdBal,omitempty"`

	// Balance that has been affected by the process run through the event.
	AffectedBalance *BalanceFormat1Choice `xml:"AfctdBal,omitempty"`

	// Balance that has not been affected by the process run through the event.
	UnaffectedBalance *BalanceFormat1Choice `xml:"UafctdBal,omitempty"`
}

func (c *CorporateActionBalanceDetails18) AddConfirmedBalance() *BalanceFormat1Choice {
	c.ConfirmedBalance = new(BalanceFormat1Choice)
	return c.ConfirmedBalance
}

func (c *CorporateActionBalanceDetails18) AddTotalEligibleBalance() *Quantity3Choice {
	c.TotalEligibleBalance = new(Quantity3Choice)
	return c.TotalEligibleBalance
}

func (c *CorporateActionBalanceDetails18) AddBlockedBalance() *BalanceFormat1Choice {
	c.BlockedBalance = new(BalanceFormat1Choice)
	return c.BlockedBalance
}

func (c *CorporateActionBalanceDetails18) AddBorrowedBalance() *BalanceFormat1Choice {
	c.BorrowedBalance = new(BalanceFormat1Choice)
	return c.BorrowedBalance
}

func (c *CorporateActionBalanceDetails18) AddCollateralInBalance() *BalanceFormat1Choice {
	c.CollateralInBalance = new(BalanceFormat1Choice)
	return c.CollateralInBalance
}

func (c *CorporateActionBalanceDetails18) AddCollateralOutBalance() *BalanceFormat1Choice {
	c.CollateralOutBalance = new(BalanceFormat1Choice)
	return c.CollateralOutBalance
}

func (c *CorporateActionBalanceDetails18) AddOnLoanBalance() *BalanceFormat1Choice {
	c.OnLoanBalance = new(BalanceFormat1Choice)
	return c.OnLoanBalance
}

func (c *CorporateActionBalanceDetails18) AddPendingDeliveryBalance() *BalanceFormat3Choice {
	newValue := new(BalanceFormat3Choice)
	c.PendingDeliveryBalance = append(c.PendingDeliveryBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails18) AddPendingReceiptBalance() *BalanceFormat3Choice {
	newValue := new(BalanceFormat3Choice)
	c.PendingReceiptBalance = append(c.PendingReceiptBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails18) AddOutForRegistrationBalance() *BalanceFormat1Choice {
	c.OutForRegistrationBalance = new(BalanceFormat1Choice)
	return c.OutForRegistrationBalance
}

func (c *CorporateActionBalanceDetails18) AddSettlementPositionBalance() *BalanceFormat3Choice {
	newValue := new(BalanceFormat3Choice)
	c.SettlementPositionBalance = append(c.SettlementPositionBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails18) AddStreetPositionBalance() *BalanceFormat1Choice {
	c.StreetPositionBalance = new(BalanceFormat1Choice)
	return c.StreetPositionBalance
}

func (c *CorporateActionBalanceDetails18) AddTradeDatePositionBalance() *BalanceFormat1Choice {
	c.TradeDatePositionBalance = new(BalanceFormat1Choice)
	return c.TradeDatePositionBalance
}

func (c *CorporateActionBalanceDetails18) AddInTransshipmentBalance() *BalanceFormat1Choice {
	c.InTransshipmentBalance = new(BalanceFormat1Choice)
	return c.InTransshipmentBalance
}

func (c *CorporateActionBalanceDetails18) AddRegisteredBalance() *BalanceFormat1Choice {
	c.RegisteredBalance = new(BalanceFormat1Choice)
	return c.RegisteredBalance
}

func (c *CorporateActionBalanceDetails18) AddAffectedBalance() *BalanceFormat1Choice {
	c.AffectedBalance = new(BalanceFormat1Choice)
	return c.AffectedBalance
}

func (c *CorporateActionBalanceDetails18) AddUnaffectedBalance() *BalanceFormat1Choice {
	c.UnaffectedBalance = new(BalanceFormat1Choice)
	return c.UnaffectedBalance
}

// Eligible and not eligible balance of securities for a corporate action event.
type CorporateActionBalanceDetails2 struct {

	// Balance to which the payment applies (less or equal to the total eligible balance).
	ConfirmedBalance *BalanceFormat1Choice `xml:"ConfdBal"`

	// Total quantity of financial instruments of the balance.
	TotalEligibleBalance *Quantity3Choice `xml:"TtlElgblBal,omitempty"`

	// Balance of financial instruments that are blocked.
	BlockedBalance *BalanceFormat1Choice `xml:"BlckdBal,omitempty"`

	// Quantity of securities in the sub-balance.
	BorrowedBalance *BalanceFormat1Choice `xml:"BrrwdBal,omitempty"`

	// Quantity of securities in the sub-balance.
	CollateralInBalance *BalanceFormat1Choice `xml:"CollInBal,omitempty"`

	// Quantity of securities in the sub-balance.
	CollateralOutBalance *BalanceFormat1Choice `xml:"CollOutBal,omitempty"`

	// Quantity of securities in the sub-balance.
	OnLoanBalance *BalanceFormat1Choice `xml:"OnLnBal,omitempty"`

	// Quantity of securities in the sub-balance.
	PendingDeliveryBalance []*BalanceFormat1Choice `xml:"PdgDlvryBal,omitempty"`

	// Quantity of securities in the sub-balance.
	PendingReceiptBalance []*BalanceFormat1Choice `xml:"PdgRctBal,omitempty"`

	// Quantity of securities in the sub-balance.
	OutForRegistrationBalance *BalanceFormat1Choice `xml:"OutForRegnBal,omitempty"`

	// Quantity of securities in the sub-balance.
	SettlementPositionBalance *BalanceFormat1Choice `xml:"SttlmPosBal,omitempty"`

	// Quantity of securities in the sub-balance.
	StreetPositionBalance *BalanceFormat1Choice `xml:"StrtPosBal,omitempty"`

	// Balance of securities based on trade date, for example, includes all pending transactions in addition to the balance of settled transactions.
	TradeDatePositionBalance *BalanceFormat1Choice `xml:"TradDtPosBal,omitempty"`

	// Quantity of securities in the sub-balance.
	InTransshipmentBalance *BalanceFormat1Choice `xml:"InTrnsShipmntBal,omitempty"`

	// Quantity of securities in the sub-balance.
	RegisteredBalance *BalanceFormat1Choice `xml:"RegdBal,omitempty"`

	// Balance that has been affected by the process run through the event.
	AffectedBalance *BalanceFormat1Choice `xml:"AfctdBal,omitempty"`

	// Balance that has not been affected by the process run through the event.
	UnaffectedBalance *BalanceFormat1Choice `xml:"UafctdBal,omitempty"`
}

func (c *CorporateActionBalanceDetails2) AddConfirmedBalance() *BalanceFormat1Choice {
	c.ConfirmedBalance = new(BalanceFormat1Choice)
	return c.ConfirmedBalance
}

func (c *CorporateActionBalanceDetails2) AddTotalEligibleBalance() *Quantity3Choice {
	c.TotalEligibleBalance = new(Quantity3Choice)
	return c.TotalEligibleBalance
}

func (c *CorporateActionBalanceDetails2) AddBlockedBalance() *BalanceFormat1Choice {
	c.BlockedBalance = new(BalanceFormat1Choice)
	return c.BlockedBalance
}

func (c *CorporateActionBalanceDetails2) AddBorrowedBalance() *BalanceFormat1Choice {
	c.BorrowedBalance = new(BalanceFormat1Choice)
	return c.BorrowedBalance
}

func (c *CorporateActionBalanceDetails2) AddCollateralInBalance() *BalanceFormat1Choice {
	c.CollateralInBalance = new(BalanceFormat1Choice)
	return c.CollateralInBalance
}

func (c *CorporateActionBalanceDetails2) AddCollateralOutBalance() *BalanceFormat1Choice {
	c.CollateralOutBalance = new(BalanceFormat1Choice)
	return c.CollateralOutBalance
}

func (c *CorporateActionBalanceDetails2) AddOnLoanBalance() *BalanceFormat1Choice {
	c.OnLoanBalance = new(BalanceFormat1Choice)
	return c.OnLoanBalance
}

func (c *CorporateActionBalanceDetails2) AddPendingDeliveryBalance() *BalanceFormat1Choice {
	newValue := new(BalanceFormat1Choice)
	c.PendingDeliveryBalance = append(c.PendingDeliveryBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails2) AddPendingReceiptBalance() *BalanceFormat1Choice {
	newValue := new(BalanceFormat1Choice)
	c.PendingReceiptBalance = append(c.PendingReceiptBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails2) AddOutForRegistrationBalance() *BalanceFormat1Choice {
	c.OutForRegistrationBalance = new(BalanceFormat1Choice)
	return c.OutForRegistrationBalance
}

func (c *CorporateActionBalanceDetails2) AddSettlementPositionBalance() *BalanceFormat1Choice {
	c.SettlementPositionBalance = new(BalanceFormat1Choice)
	return c.SettlementPositionBalance
}

func (c *CorporateActionBalanceDetails2) AddStreetPositionBalance() *BalanceFormat1Choice {
	c.StreetPositionBalance = new(BalanceFormat1Choice)
	return c.StreetPositionBalance
}

func (c *CorporateActionBalanceDetails2) AddTradeDatePositionBalance() *BalanceFormat1Choice {
	c.TradeDatePositionBalance = new(BalanceFormat1Choice)
	return c.TradeDatePositionBalance
}

func (c *CorporateActionBalanceDetails2) AddInTransshipmentBalance() *BalanceFormat1Choice {
	c.InTransshipmentBalance = new(BalanceFormat1Choice)
	return c.InTransshipmentBalance
}

func (c *CorporateActionBalanceDetails2) AddRegisteredBalance() *BalanceFormat1Choice {
	c.RegisteredBalance = new(BalanceFormat1Choice)
	return c.RegisteredBalance
}

func (c *CorporateActionBalanceDetails2) AddAffectedBalance() *BalanceFormat1Choice {
	c.AffectedBalance = new(BalanceFormat1Choice)
	return c.AffectedBalance
}

func (c *CorporateActionBalanceDetails2) AddUnaffectedBalance() *BalanceFormat1Choice {
	c.UnaffectedBalance = new(BalanceFormat1Choice)
	return c.UnaffectedBalance
}

// Eligible and not eligible balance of securities for a corporate action event.
type CorporateActionBalanceDetails21 struct {

	// Total balance of securities eligible for this corporate action event. The entitlement calculation is based on this balance.
	TotalEligibleBalance *TotalEligibleBalanceFormat1 `xml:"TtlElgblBal,omitempty"`

	// Balance of financial instruments that are blocked.
	BlockedBalance *BalanceFormat1Choice `xml:"BlckdBal,omitempty"`

	// Balance of financial instruments that have been borrowed from another party.
	BorrowedBalance *BalanceFormat1Choice `xml:"BrrwdBal,omitempty"`

	// Balance of securities that belong to a third party and that are held for the purpose of collateralisation.
	CollateralInBalance *BalanceFormat1Choice `xml:"CollInBal,omitempty"`

	// Balance of securities that belong to the safekeeping account indicated within this message, and are deposited with a third party for the purpose of collateralisation.
	CollateralOutBalance *BalanceFormat1Choice `xml:"CollOutBal,omitempty"`

	// Balance of financial instruments that have been loaned to a third party.
	OnLoanBalance *BalanceFormat1Choice `xml:"OnLnBal,omitempty"`

	// Balance of financial instruments that are pending delivery.
	PendingDeliveryBalance []*BalanceFormat3Choice `xml:"PdgDlvryBal,omitempty"`

	// Balance of financial instruments that are pending receipt.
	PendingReceiptBalance []*BalanceFormat3Choice `xml:"PdgRctBal,omitempty"`

	// Balance of financial instruments currently being processed by the institution responsible for registering the new beneficial owner (or nominee).
	OutForRegistrationBalance *BalanceFormat1Choice `xml:"OutForRegnBal,omitempty"`

	// Balance of securities representing only settled transactions; pending transactions not included.
	SettlementPositionBalance []*BalanceFormat3Choice `xml:"SttlmPosBal,omitempty"`

	// Balance of financial instruments that remain registered in the name of the prior beneficial owner.
	StreetPositionBalance *BalanceFormat1Choice `xml:"StrtPosBal,omitempty"`

	// Balance of securities based on trade date, for example, includes all pending transactions in addition to the balance of settled transactions.
	TradeDatePositionBalance *BalanceFormat1Choice `xml:"TradDtPosBal,omitempty"`

	// Balance of physical securities that are in the process of being transferred from one depository/agent to another.
	InTransshipmentBalance *BalanceFormat1Choice `xml:"InTrnsShipmntBal,omitempty"`

	// Balance of financial instruments that are registered (in the name of a nominee name or of the beneficial owner).
	RegisteredBalance *BalanceFormat1Choice `xml:"RegdBal,omitempty"`

	// Position that account holders should return to the account servicer to participate in the event or to fulfil their obligation for the event to be complete, for example, return of securities for late announced drawing.
	ObligatedBalance *BalanceFormat1Choice `xml:"OblgtdBal,omitempty"`

	// Balance of uninstructed position.
	UninstructedBalance *BalanceFormat1Choice `xml:"UinstdBal,omitempty"`

	// Balance of instructed position.
	InstructedBalance *BalanceFormat1Choice `xml:"InstdBal,omitempty"`

	// Balance that has been affected by the process run through the event.
	AffectedBalance *BalanceFormat1Choice `xml:"AfctdBal,omitempty"`

	// Balance that has not been affected by the process run through the event.
	UnaffectedBalance *BalanceFormat1Choice `xml:"UafctdBal,omitempty"`
}

func (c *CorporateActionBalanceDetails21) AddTotalEligibleBalance() *TotalEligibleBalanceFormat1 {
	c.TotalEligibleBalance = new(TotalEligibleBalanceFormat1)
	return c.TotalEligibleBalance
}

func (c *CorporateActionBalanceDetails21) AddBlockedBalance() *BalanceFormat1Choice {
	c.BlockedBalance = new(BalanceFormat1Choice)
	return c.BlockedBalance
}

func (c *CorporateActionBalanceDetails21) AddBorrowedBalance() *BalanceFormat1Choice {
	c.BorrowedBalance = new(BalanceFormat1Choice)
	return c.BorrowedBalance
}

func (c *CorporateActionBalanceDetails21) AddCollateralInBalance() *BalanceFormat1Choice {
	c.CollateralInBalance = new(BalanceFormat1Choice)
	return c.CollateralInBalance
}

func (c *CorporateActionBalanceDetails21) AddCollateralOutBalance() *BalanceFormat1Choice {
	c.CollateralOutBalance = new(BalanceFormat1Choice)
	return c.CollateralOutBalance
}

func (c *CorporateActionBalanceDetails21) AddOnLoanBalance() *BalanceFormat1Choice {
	c.OnLoanBalance = new(BalanceFormat1Choice)
	return c.OnLoanBalance
}

func (c *CorporateActionBalanceDetails21) AddPendingDeliveryBalance() *BalanceFormat3Choice {
	newValue := new(BalanceFormat3Choice)
	c.PendingDeliveryBalance = append(c.PendingDeliveryBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails21) AddPendingReceiptBalance() *BalanceFormat3Choice {
	newValue := new(BalanceFormat3Choice)
	c.PendingReceiptBalance = append(c.PendingReceiptBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails21) AddOutForRegistrationBalance() *BalanceFormat1Choice {
	c.OutForRegistrationBalance = new(BalanceFormat1Choice)
	return c.OutForRegistrationBalance
}

func (c *CorporateActionBalanceDetails21) AddSettlementPositionBalance() *BalanceFormat3Choice {
	newValue := new(BalanceFormat3Choice)
	c.SettlementPositionBalance = append(c.SettlementPositionBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails21) AddStreetPositionBalance() *BalanceFormat1Choice {
	c.StreetPositionBalance = new(BalanceFormat1Choice)
	return c.StreetPositionBalance
}

func (c *CorporateActionBalanceDetails21) AddTradeDatePositionBalance() *BalanceFormat1Choice {
	c.TradeDatePositionBalance = new(BalanceFormat1Choice)
	return c.TradeDatePositionBalance
}

func (c *CorporateActionBalanceDetails21) AddInTransshipmentBalance() *BalanceFormat1Choice {
	c.InTransshipmentBalance = new(BalanceFormat1Choice)
	return c.InTransshipmentBalance
}

func (c *CorporateActionBalanceDetails21) AddRegisteredBalance() *BalanceFormat1Choice {
	c.RegisteredBalance = new(BalanceFormat1Choice)
	return c.RegisteredBalance
}

func (c *CorporateActionBalanceDetails21) AddObligatedBalance() *BalanceFormat1Choice {
	c.ObligatedBalance = new(BalanceFormat1Choice)
	return c.ObligatedBalance
}

func (c *CorporateActionBalanceDetails21) AddUninstructedBalance() *BalanceFormat1Choice {
	c.UninstructedBalance = new(BalanceFormat1Choice)
	return c.UninstructedBalance
}

func (c *CorporateActionBalanceDetails21) AddInstructedBalance() *BalanceFormat1Choice {
	c.InstructedBalance = new(BalanceFormat1Choice)
	return c.InstructedBalance
}

func (c *CorporateActionBalanceDetails21) AddAffectedBalance() *BalanceFormat1Choice {
	c.AffectedBalance = new(BalanceFormat1Choice)
	return c.AffectedBalance
}

func (c *CorporateActionBalanceDetails21) AddUnaffectedBalance() *BalanceFormat1Choice {
	c.UnaffectedBalance = new(BalanceFormat1Choice)
	return c.UnaffectedBalance
}

// Eligible and not eligible balance of securities for a corporate action event.
type CorporateActionBalanceDetails22 struct {

	// Balance to which the payment applies (less or equal to the total eligible balance).
	ConfirmedBalance *BalanceFormat1Choice `xml:"ConfdBal"`

	// Total balance of securities eligible for this corporate action event. The entitlement calculation is based on this balance.
	TotalEligibleBalance *TotalEligibleBalanceFormat1 `xml:"TtlElgblBal,omitempty"`

	// Balance of financial instruments that are blocked.
	BlockedBalance *BalanceFormat1Choice `xml:"BlckdBal,omitempty"`

	// Balance of financial instruments that have been borrowed from another party.
	BorrowedBalance *BalanceFormat1Choice `xml:"BrrwdBal,omitempty"`

	// Balance of securities that belong to a third party and that are held for the purpose of collateralisation.
	CollateralInBalance *BalanceFormat1Choice `xml:"CollInBal,omitempty"`

	// Balance of securities that belong to the safekeeping account indicated within this message, and are deposited with a third party for the purpose of collateralisation.
	CollateralOutBalance *BalanceFormat1Choice `xml:"CollOutBal,omitempty"`

	// Balance of financial instruments that have been loaned to a third party.
	OnLoanBalance *BalanceFormat1Choice `xml:"OnLnBal,omitempty"`

	// Balance of financial instruments that are pending delivery.
	PendingDeliveryBalance []*BalanceFormat3Choice `xml:"PdgDlvryBal,omitempty"`

	// Balance of financial instruments that are pending receipt.
	PendingReceiptBalance []*BalanceFormat3Choice `xml:"PdgRctBal,omitempty"`

	// Balance of financial instruments currently being processed by the institution responsible for registering the new beneficial owner (or nominee).
	OutForRegistrationBalance *BalanceFormat1Choice `xml:"OutForRegnBal,omitempty"`

	// Balance of securities representing only settled transactions; pending transactions not included.
	SettlementPositionBalance []*BalanceFormat3Choice `xml:"SttlmPosBal,omitempty"`

	// Balance of financial instruments that remain registered in the name of the prior beneficial owner.
	StreetPositionBalance *BalanceFormat1Choice `xml:"StrtPosBal,omitempty"`

	// Balance of securities based on trade date, for example, includes all pending transactions in addition to the balance of settled transactions.
	TradeDatePositionBalance *BalanceFormat1Choice `xml:"TradDtPosBal,omitempty"`

	// Balance of physical securities that are in the process of being transferred from one depository/agent to another.
	InTransshipmentBalance *BalanceFormat1Choice `xml:"InTrnsShipmntBal,omitempty"`

	// Balance of financial instruments that are registered (in the name of a nominee name or of the beneficial owner).
	RegisteredBalance *BalanceFormat1Choice `xml:"RegdBal,omitempty"`

	// Balance that has been affected by the process run through the event.
	AffectedBalance *BalanceFormat1Choice `xml:"AfctdBal,omitempty"`

	// Balance that has not been affected by the process run through the event.
	UnaffectedBalance *BalanceFormat1Choice `xml:"UafctdBal,omitempty"`
}

func (c *CorporateActionBalanceDetails22) AddConfirmedBalance() *BalanceFormat1Choice {
	c.ConfirmedBalance = new(BalanceFormat1Choice)
	return c.ConfirmedBalance
}

func (c *CorporateActionBalanceDetails22) AddTotalEligibleBalance() *TotalEligibleBalanceFormat1 {
	c.TotalEligibleBalance = new(TotalEligibleBalanceFormat1)
	return c.TotalEligibleBalance
}

func (c *CorporateActionBalanceDetails22) AddBlockedBalance() *BalanceFormat1Choice {
	c.BlockedBalance = new(BalanceFormat1Choice)
	return c.BlockedBalance
}

func (c *CorporateActionBalanceDetails22) AddBorrowedBalance() *BalanceFormat1Choice {
	c.BorrowedBalance = new(BalanceFormat1Choice)
	return c.BorrowedBalance
}

func (c *CorporateActionBalanceDetails22) AddCollateralInBalance() *BalanceFormat1Choice {
	c.CollateralInBalance = new(BalanceFormat1Choice)
	return c.CollateralInBalance
}

func (c *CorporateActionBalanceDetails22) AddCollateralOutBalance() *BalanceFormat1Choice {
	c.CollateralOutBalance = new(BalanceFormat1Choice)
	return c.CollateralOutBalance
}

func (c *CorporateActionBalanceDetails22) AddOnLoanBalance() *BalanceFormat1Choice {
	c.OnLoanBalance = new(BalanceFormat1Choice)
	return c.OnLoanBalance
}

func (c *CorporateActionBalanceDetails22) AddPendingDeliveryBalance() *BalanceFormat3Choice {
	newValue := new(BalanceFormat3Choice)
	c.PendingDeliveryBalance = append(c.PendingDeliveryBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails22) AddPendingReceiptBalance() *BalanceFormat3Choice {
	newValue := new(BalanceFormat3Choice)
	c.PendingReceiptBalance = append(c.PendingReceiptBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails22) AddOutForRegistrationBalance() *BalanceFormat1Choice {
	c.OutForRegistrationBalance = new(BalanceFormat1Choice)
	return c.OutForRegistrationBalance
}

func (c *CorporateActionBalanceDetails22) AddSettlementPositionBalance() *BalanceFormat3Choice {
	newValue := new(BalanceFormat3Choice)
	c.SettlementPositionBalance = append(c.SettlementPositionBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails22) AddStreetPositionBalance() *BalanceFormat1Choice {
	c.StreetPositionBalance = new(BalanceFormat1Choice)
	return c.StreetPositionBalance
}

func (c *CorporateActionBalanceDetails22) AddTradeDatePositionBalance() *BalanceFormat1Choice {
	c.TradeDatePositionBalance = new(BalanceFormat1Choice)
	return c.TradeDatePositionBalance
}

func (c *CorporateActionBalanceDetails22) AddInTransshipmentBalance() *BalanceFormat1Choice {
	c.InTransshipmentBalance = new(BalanceFormat1Choice)
	return c.InTransshipmentBalance
}

func (c *CorporateActionBalanceDetails22) AddRegisteredBalance() *BalanceFormat1Choice {
	c.RegisteredBalance = new(BalanceFormat1Choice)
	return c.RegisteredBalance
}

func (c *CorporateActionBalanceDetails22) AddAffectedBalance() *BalanceFormat1Choice {
	c.AffectedBalance = new(BalanceFormat1Choice)
	return c.AffectedBalance
}

func (c *CorporateActionBalanceDetails22) AddUnaffectedBalance() *BalanceFormat1Choice {
	c.UnaffectedBalance = new(BalanceFormat1Choice)
	return c.UnaffectedBalance
}

// Eligible and not eligible balance of securities for a corporate action event.
type CorporateActionBalanceDetails29 struct {

	// Total balance of securities eligible for this corporate action event. The entitlement calculation is based on this balance.
	TotalEligibleBalance *TotalEligibleBalanceFormat8 `xml:"TtlElgblBal,omitempty"`

	// Balance of financial instruments that are blocked.
	BlockedBalance *BalanceFormat5Choice `xml:"BlckdBal,omitempty"`

	// Balance of financial instruments that have been borrowed from another party.
	BorrowedBalance *BalanceFormat5Choice `xml:"BrrwdBal,omitempty"`

	// Balance of securities that belong to a third party and that are held for the purpose of collateralisation.
	CollateralInBalance *BalanceFormat5Choice `xml:"CollInBal,omitempty"`

	// Balance of securities that belong to the safekeeping account indicated within this message, and are deposited with a third party for the purpose of collateralisation.
	CollateralOutBalance *BalanceFormat5Choice `xml:"CollOutBal,omitempty"`

	// Balance of financial instruments that have been loaned to a third party.
	OnLoanBalance *BalanceFormat5Choice `xml:"OnLnBal,omitempty"`

	// Balance of financial instruments that are pending delivery.
	PendingDeliveryBalance []*BalanceFormat6Choice `xml:"PdgDlvryBal,omitempty"`

	// Balance of financial instruments that are pending receipt.
	PendingReceiptBalance []*BalanceFormat6Choice `xml:"PdgRctBal,omitempty"`

	// Balance of financial instruments currently being processed by the institution responsible for registering the new beneficial owner (or nominee).
	OutForRegistrationBalance *BalanceFormat5Choice `xml:"OutForRegnBal,omitempty"`

	// Balance of securities representing only settled transactions; pending transactions not included.
	SettlementPositionBalance []*BalanceFormat6Choice `xml:"SttlmPosBal,omitempty"`

	// Balance of financial instruments that remain registered in the name of the prior beneficial owner.
	StreetPositionBalance *BalanceFormat5Choice `xml:"StrtPosBal,omitempty"`

	// Balance of securities based on trade date, for example, includes all pending transactions in addition to the balance of settled transactions.
	TradeDatePositionBalance *BalanceFormat5Choice `xml:"TradDtPosBal,omitempty"`

	// Balance of physical securities that are in the process of being transferred from one depository/agent to another.
	InTransshipmentBalance *BalanceFormat5Choice `xml:"InTrnsShipmntBal,omitempty"`

	// Balance of financial instruments that are registered (in the name of a nominee name or of the beneficial owner).
	RegisteredBalance *BalanceFormat5Choice `xml:"RegdBal,omitempty"`

	// Position that account holders should return to the account servicer to participate in the event or to fulfil their obligation for the event to be complete, for example, return of securities for late announced drawing.
	ObligatedBalance *BalanceFormat5Choice `xml:"OblgtdBal,omitempty"`

	// Balance of uninstructed position.
	UninstructedBalance *BalanceFormat5Choice `xml:"UinstdBal,omitempty"`

	// Balance of instructed position.
	InstructedBalance *BalanceFormat5Choice `xml:"InstdBal,omitempty"`

	// Balance that has been affected by the process run through the event.
	AffectedBalance *BalanceFormat5Choice `xml:"AfctdBal,omitempty"`

	// Balance that has not been affected by the process run through the event.
	UnaffectedBalance *BalanceFormat5Choice `xml:"UafctdBal,omitempty"`
}

func (c *CorporateActionBalanceDetails29) AddTotalEligibleBalance() *TotalEligibleBalanceFormat8 {
	c.TotalEligibleBalance = new(TotalEligibleBalanceFormat8)
	return c.TotalEligibleBalance
}

func (c *CorporateActionBalanceDetails29) AddBlockedBalance() *BalanceFormat5Choice {
	c.BlockedBalance = new(BalanceFormat5Choice)
	return c.BlockedBalance
}

func (c *CorporateActionBalanceDetails29) AddBorrowedBalance() *BalanceFormat5Choice {
	c.BorrowedBalance = new(BalanceFormat5Choice)
	return c.BorrowedBalance
}

func (c *CorporateActionBalanceDetails29) AddCollateralInBalance() *BalanceFormat5Choice {
	c.CollateralInBalance = new(BalanceFormat5Choice)
	return c.CollateralInBalance
}

func (c *CorporateActionBalanceDetails29) AddCollateralOutBalance() *BalanceFormat5Choice {
	c.CollateralOutBalance = new(BalanceFormat5Choice)
	return c.CollateralOutBalance
}

func (c *CorporateActionBalanceDetails29) AddOnLoanBalance() *BalanceFormat5Choice {
	c.OnLoanBalance = new(BalanceFormat5Choice)
	return c.OnLoanBalance
}

func (c *CorporateActionBalanceDetails29) AddPendingDeliveryBalance() *BalanceFormat6Choice {
	newValue := new(BalanceFormat6Choice)
	c.PendingDeliveryBalance = append(c.PendingDeliveryBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails29) AddPendingReceiptBalance() *BalanceFormat6Choice {
	newValue := new(BalanceFormat6Choice)
	c.PendingReceiptBalance = append(c.PendingReceiptBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails29) AddOutForRegistrationBalance() *BalanceFormat5Choice {
	c.OutForRegistrationBalance = new(BalanceFormat5Choice)
	return c.OutForRegistrationBalance
}

func (c *CorporateActionBalanceDetails29) AddSettlementPositionBalance() *BalanceFormat6Choice {
	newValue := new(BalanceFormat6Choice)
	c.SettlementPositionBalance = append(c.SettlementPositionBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails29) AddStreetPositionBalance() *BalanceFormat5Choice {
	c.StreetPositionBalance = new(BalanceFormat5Choice)
	return c.StreetPositionBalance
}

func (c *CorporateActionBalanceDetails29) AddTradeDatePositionBalance() *BalanceFormat5Choice {
	c.TradeDatePositionBalance = new(BalanceFormat5Choice)
	return c.TradeDatePositionBalance
}

func (c *CorporateActionBalanceDetails29) AddInTransshipmentBalance() *BalanceFormat5Choice {
	c.InTransshipmentBalance = new(BalanceFormat5Choice)
	return c.InTransshipmentBalance
}

func (c *CorporateActionBalanceDetails29) AddRegisteredBalance() *BalanceFormat5Choice {
	c.RegisteredBalance = new(BalanceFormat5Choice)
	return c.RegisteredBalance
}

func (c *CorporateActionBalanceDetails29) AddObligatedBalance() *BalanceFormat5Choice {
	c.ObligatedBalance = new(BalanceFormat5Choice)
	return c.ObligatedBalance
}

func (c *CorporateActionBalanceDetails29) AddUninstructedBalance() *BalanceFormat5Choice {
	c.UninstructedBalance = new(BalanceFormat5Choice)
	return c.UninstructedBalance
}

func (c *CorporateActionBalanceDetails29) AddInstructedBalance() *BalanceFormat5Choice {
	c.InstructedBalance = new(BalanceFormat5Choice)
	return c.InstructedBalance
}

func (c *CorporateActionBalanceDetails29) AddAffectedBalance() *BalanceFormat5Choice {
	c.AffectedBalance = new(BalanceFormat5Choice)
	return c.AffectedBalance
}

func (c *CorporateActionBalanceDetails29) AddUnaffectedBalance() *BalanceFormat5Choice {
	c.UnaffectedBalance = new(BalanceFormat5Choice)
	return c.UnaffectedBalance
}

// Eligible and not eligible balance of securities for a corporate action event.
type CorporateActionBalanceDetails3 struct {

	// Total quantity of financial instruments of the balance.
	TotalEligibleBalance *Quantity3Choice `xml:"TtlElgblBal,omitempty"`

	// Balance of financial instruments that are blocked.
	BlockedBalance *BalanceFormat1Choice `xml:"BlckdBal,omitempty"`

	// Quantity of securities in the sub-balance.
	BorrowedBalance *BalanceFormat1Choice `xml:"BrrwdBal,omitempty"`

	// Quantity of securities in the sub-balance.
	CollateralInBalance *BalanceFormat1Choice `xml:"CollInBal,omitempty"`

	// Quantity of securities in the sub-balance.
	CollateralOutBalance *BalanceFormat1Choice `xml:"CollOutBal,omitempty"`

	// Quantity of securities in the sub-balance.
	OnLoanBalance *BalanceFormat1Choice `xml:"OnLnBal,omitempty"`

	// Quantity of securities in the sub-balance.
	PendingDeliveryBalance []*BalanceFormat1Choice `xml:"PdgDlvryBal,omitempty"`

	// Quantity of securities in the sub-balance.
	PendingReceiptBalance []*BalanceFormat1Choice `xml:"PdgRctBal,omitempty"`

	// Quantity of securities in the sub-balance.
	OutForRegistrationBalance *BalanceFormat1Choice `xml:"OutForRegnBal,omitempty"`

	// Quantity of securities in the sub-balance.
	SettlementPositionBalance *BalanceFormat1Choice `xml:"SttlmPosBal,omitempty"`

	// Quantity of securities in the sub-balance.
	StreetPositionBalance *BalanceFormat1Choice `xml:"StrtPosBal,omitempty"`

	// Balance of securities based on trade date, for example, includes all pending transactions in addition to the balance of settled transactions.
	TradeDatePositionBalance *BalanceFormat1Choice `xml:"TradDtPosBal,omitempty"`

	// Quantity of securities in the sub-balance.
	InTransshipmentBalance *BalanceFormat1Choice `xml:"InTrnsShipmntBal,omitempty"`

	// Quantity of securities in the sub-balance.
	RegisteredBalance *BalanceFormat1Choice `xml:"RegdBal,omitempty"`
}

func (c *CorporateActionBalanceDetails3) AddTotalEligibleBalance() *Quantity3Choice {
	c.TotalEligibleBalance = new(Quantity3Choice)
	return c.TotalEligibleBalance
}

func (c *CorporateActionBalanceDetails3) AddBlockedBalance() *BalanceFormat1Choice {
	c.BlockedBalance = new(BalanceFormat1Choice)
	return c.BlockedBalance
}

func (c *CorporateActionBalanceDetails3) AddBorrowedBalance() *BalanceFormat1Choice {
	c.BorrowedBalance = new(BalanceFormat1Choice)
	return c.BorrowedBalance
}

func (c *CorporateActionBalanceDetails3) AddCollateralInBalance() *BalanceFormat1Choice {
	c.CollateralInBalance = new(BalanceFormat1Choice)
	return c.CollateralInBalance
}

func (c *CorporateActionBalanceDetails3) AddCollateralOutBalance() *BalanceFormat1Choice {
	c.CollateralOutBalance = new(BalanceFormat1Choice)
	return c.CollateralOutBalance
}

func (c *CorporateActionBalanceDetails3) AddOnLoanBalance() *BalanceFormat1Choice {
	c.OnLoanBalance = new(BalanceFormat1Choice)
	return c.OnLoanBalance
}

func (c *CorporateActionBalanceDetails3) AddPendingDeliveryBalance() *BalanceFormat1Choice {
	newValue := new(BalanceFormat1Choice)
	c.PendingDeliveryBalance = append(c.PendingDeliveryBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails3) AddPendingReceiptBalance() *BalanceFormat1Choice {
	newValue := new(BalanceFormat1Choice)
	c.PendingReceiptBalance = append(c.PendingReceiptBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails3) AddOutForRegistrationBalance() *BalanceFormat1Choice {
	c.OutForRegistrationBalance = new(BalanceFormat1Choice)
	return c.OutForRegistrationBalance
}

func (c *CorporateActionBalanceDetails3) AddSettlementPositionBalance() *BalanceFormat1Choice {
	c.SettlementPositionBalance = new(BalanceFormat1Choice)
	return c.SettlementPositionBalance
}

func (c *CorporateActionBalanceDetails3) AddStreetPositionBalance() *BalanceFormat1Choice {
	c.StreetPositionBalance = new(BalanceFormat1Choice)
	return c.StreetPositionBalance
}

func (c *CorporateActionBalanceDetails3) AddTradeDatePositionBalance() *BalanceFormat1Choice {
	c.TradeDatePositionBalance = new(BalanceFormat1Choice)
	return c.TradeDatePositionBalance
}

func (c *CorporateActionBalanceDetails3) AddInTransshipmentBalance() *BalanceFormat1Choice {
	c.InTransshipmentBalance = new(BalanceFormat1Choice)
	return c.InTransshipmentBalance
}

func (c *CorporateActionBalanceDetails3) AddRegisteredBalance() *BalanceFormat1Choice {
	c.RegisteredBalance = new(BalanceFormat1Choice)
	return c.RegisteredBalance
}

// Eligible and not eligible balance of securities for a corporate action event.
type CorporateActionBalanceDetails30 struct {

	// Total quantity of financial instruments of the balance.
	TotalEligibleBalance *Quantity17Choice `xml:"TtlElgblBal"`

	// Quantity of securities in the sub-balance.
	UninstructedBalance *BalanceFormat5Choice `xml:"UinstdBal"`

	// Provides information about the total instructed balance.
	TotalInstructedBalanceDetails *InstructedBalanceDetails5 `xml:"TtlInstdBalDtls"`

	// Balance of financial instruments that are blocked.
	BlockedBalance *SignedQuantityFormat6 `xml:"BlckdBal,omitempty"`

	// Balance of financial instruments that have been borrowed from another party.
	BorrowedBalance *SignedQuantityFormat6 `xml:"BrrwdBal,omitempty"`

	// Balance of securities that belong to a third party and that are held for the purpose of collateralisation.
	CollateralInBalance *SignedQuantityFormat6 `xml:"CollInBal,omitempty"`

	// Balance of securities that belong to the safekeeping account indicated within this message, and are deposited with a third party for the purpose of collateralisation.
	CollateralOutBalance *SignedQuantityFormat6 `xml:"CollOutBal,omitempty"`

	// Balance of financial instruments that have been loaned to a third party.
	OnLoanBalance *SignedQuantityFormat6 `xml:"OnLnBal,omitempty"`

	// Balance of financial instruments currently being processed by the institution responsible for registering the new beneficial owner (or nominee).
	OutForRegistrationBalance *SignedQuantityFormat6 `xml:"OutForRegnBal,omitempty"`

	// Balance of securities representing only settled transactions; pending transactions not included.
	SettlementPositionBalance *SignedQuantityFormat6 `xml:"SttlmPosBal,omitempty"`

	// Balance of financial instruments that remain registered in the name of the prior beneficial owner.
	StreetPositionBalance *SignedQuantityFormat6 `xml:"StrtPosBal,omitempty"`

	// Balance of securities based on trade date, for example, includes all pending transactions in addition to the balance of settled transactions.
	TradeDatePositionBalance *SignedQuantityFormat6 `xml:"TradDtPosBal,omitempty"`

	// Balance of physical securities that are in the process of being transferred from one depository/agent to another.
	InTransshipmentBalance *SignedQuantityFormat6 `xml:"InTrnsShipmntBal,omitempty"`

	// Balance of financial instruments that are registered (in the name of a nominee name or of the beneficial owner).
	RegisteredBalance *SignedQuantityFormat6 `xml:"RegdBal,omitempty"`

	// Position that account holders should return to the account servicer to participate in the event or to fulfil their obligation for the event to be complete, for example, return of securities for late announced drawing.
	ObligatedBalance *SignedQuantityFormat6 `xml:"OblgtdBal,omitempty"`

	// Balance of financial instruments that are pending delivery.
	PendingDeliveryBalance []*PendingBalance3 `xml:"PdgDlvryBal,omitempty"`

	// Balance of financial instruments that are pending receipt.
	PendingReceiptBalance []*PendingBalance3 `xml:"PdgRctBal,omitempty"`
}

func (c *CorporateActionBalanceDetails30) AddTotalEligibleBalance() *Quantity17Choice {
	c.TotalEligibleBalance = new(Quantity17Choice)
	return c.TotalEligibleBalance
}

func (c *CorporateActionBalanceDetails30) AddUninstructedBalance() *BalanceFormat5Choice {
	c.UninstructedBalance = new(BalanceFormat5Choice)
	return c.UninstructedBalance
}

func (c *CorporateActionBalanceDetails30) AddTotalInstructedBalanceDetails() *InstructedBalanceDetails5 {
	c.TotalInstructedBalanceDetails = new(InstructedBalanceDetails5)
	return c.TotalInstructedBalanceDetails
}

func (c *CorporateActionBalanceDetails30) AddBlockedBalance() *SignedQuantityFormat6 {
	c.BlockedBalance = new(SignedQuantityFormat6)
	return c.BlockedBalance
}

func (c *CorporateActionBalanceDetails30) AddBorrowedBalance() *SignedQuantityFormat6 {
	c.BorrowedBalance = new(SignedQuantityFormat6)
	return c.BorrowedBalance
}

func (c *CorporateActionBalanceDetails30) AddCollateralInBalance() *SignedQuantityFormat6 {
	c.CollateralInBalance = new(SignedQuantityFormat6)
	return c.CollateralInBalance
}

func (c *CorporateActionBalanceDetails30) AddCollateralOutBalance() *SignedQuantityFormat6 {
	c.CollateralOutBalance = new(SignedQuantityFormat6)
	return c.CollateralOutBalance
}

func (c *CorporateActionBalanceDetails30) AddOnLoanBalance() *SignedQuantityFormat6 {
	c.OnLoanBalance = new(SignedQuantityFormat6)
	return c.OnLoanBalance
}

func (c *CorporateActionBalanceDetails30) AddOutForRegistrationBalance() *SignedQuantityFormat6 {
	c.OutForRegistrationBalance = new(SignedQuantityFormat6)
	return c.OutForRegistrationBalance
}

func (c *CorporateActionBalanceDetails30) AddSettlementPositionBalance() *SignedQuantityFormat6 {
	c.SettlementPositionBalance = new(SignedQuantityFormat6)
	return c.SettlementPositionBalance
}

func (c *CorporateActionBalanceDetails30) AddStreetPositionBalance() *SignedQuantityFormat6 {
	c.StreetPositionBalance = new(SignedQuantityFormat6)
	return c.StreetPositionBalance
}

func (c *CorporateActionBalanceDetails30) AddTradeDatePositionBalance() *SignedQuantityFormat6 {
	c.TradeDatePositionBalance = new(SignedQuantityFormat6)
	return c.TradeDatePositionBalance
}

func (c *CorporateActionBalanceDetails30) AddInTransshipmentBalance() *SignedQuantityFormat6 {
	c.InTransshipmentBalance = new(SignedQuantityFormat6)
	return c.InTransshipmentBalance
}

func (c *CorporateActionBalanceDetails30) AddRegisteredBalance() *SignedQuantityFormat6 {
	c.RegisteredBalance = new(SignedQuantityFormat6)
	return c.RegisteredBalance
}

func (c *CorporateActionBalanceDetails30) AddObligatedBalance() *SignedQuantityFormat6 {
	c.ObligatedBalance = new(SignedQuantityFormat6)
	return c.ObligatedBalance
}

func (c *CorporateActionBalanceDetails30) AddPendingDeliveryBalance() *PendingBalance3 {
	newValue := new(PendingBalance3)
	c.PendingDeliveryBalance = append(c.PendingDeliveryBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails30) AddPendingReceiptBalance() *PendingBalance3 {
	newValue := new(PendingBalance3)
	c.PendingReceiptBalance = append(c.PendingReceiptBalance, newValue)
	return newValue
}

// Eligible and not eligible balance of securities for a corporate action event.
type CorporateActionBalanceDetails31 struct {

	// Balance to which the payment applies (less or equal to the total eligible balance).
	ConfirmedBalance *BalanceFormat5Choice `xml:"ConfdBal"`

	// Total balance of securities eligible for this corporate action event. The entitlement calculation is based on this balance.
	TotalEligibleBalance *TotalEligibleBalanceFormat8 `xml:"TtlElgblBal,omitempty"`

	// Balance of financial instruments that are blocked.
	BlockedBalance *BalanceFormat5Choice `xml:"BlckdBal,omitempty"`

	// Balance of financial instruments that have been borrowed from another party.
	BorrowedBalance *BalanceFormat5Choice `xml:"BrrwdBal,omitempty"`

	// Balance of securities that belong to a third party and that are held for the purpose of collateralisation.
	CollateralInBalance *BalanceFormat5Choice `xml:"CollInBal,omitempty"`

	// Balance of securities that belong to the safekeeping account indicated within this message, and are deposited with a third party for the purpose of collateralisation.
	CollateralOutBalance *BalanceFormat5Choice `xml:"CollOutBal,omitempty"`

	// Balance of financial instruments that have been loaned to a third party.
	OnLoanBalance *BalanceFormat5Choice `xml:"OnLnBal,omitempty"`

	// Balance of financial instruments that are pending delivery.
	PendingDeliveryBalance []*BalanceFormat6Choice `xml:"PdgDlvryBal,omitempty"`

	// Balance of financial instruments that are pending receipt.
	PendingReceiptBalance []*BalanceFormat6Choice `xml:"PdgRctBal,omitempty"`

	// Balance of financial instruments currently being processed by the institution responsible for registering the new beneficial owner (or nominee).
	OutForRegistrationBalance *BalanceFormat5Choice `xml:"OutForRegnBal,omitempty"`

	// Balance of securities representing only settled transactions; pending transactions not included.
	SettlementPositionBalance []*BalanceFormat6Choice `xml:"SttlmPosBal,omitempty"`

	// Balance of financial instruments that remain registered in the name of the prior beneficial owner.
	StreetPositionBalance *BalanceFormat5Choice `xml:"StrtPosBal,omitempty"`

	// Balance of securities based on trade date, for example, includes all pending transactions in addition to the balance of settled transactions.
	TradeDatePositionBalance *BalanceFormat5Choice `xml:"TradDtPosBal,omitempty"`

	// Balance of physical securities that are in the process of being transferred from one depository/agent to another.
	InTransshipmentBalance *BalanceFormat5Choice `xml:"InTrnsShipmntBal,omitempty"`

	// Balance of financial instruments that are registered (in the name of a nominee name or of the beneficial owner).
	RegisteredBalance *BalanceFormat5Choice `xml:"RegdBal,omitempty"`

	// Balance that has been affected by the process run through the event.
	AffectedBalance *BalanceFormat5Choice `xml:"AfctdBal,omitempty"`

	// Balance that has not been affected by the process run through the event.
	UnaffectedBalance *BalanceFormat5Choice `xml:"UafctdBal,omitempty"`
}

func (c *CorporateActionBalanceDetails31) AddConfirmedBalance() *BalanceFormat5Choice {
	c.ConfirmedBalance = new(BalanceFormat5Choice)
	return c.ConfirmedBalance
}

func (c *CorporateActionBalanceDetails31) AddTotalEligibleBalance() *TotalEligibleBalanceFormat8 {
	c.TotalEligibleBalance = new(TotalEligibleBalanceFormat8)
	return c.TotalEligibleBalance
}

func (c *CorporateActionBalanceDetails31) AddBlockedBalance() *BalanceFormat5Choice {
	c.BlockedBalance = new(BalanceFormat5Choice)
	return c.BlockedBalance
}

func (c *CorporateActionBalanceDetails31) AddBorrowedBalance() *BalanceFormat5Choice {
	c.BorrowedBalance = new(BalanceFormat5Choice)
	return c.BorrowedBalance
}

func (c *CorporateActionBalanceDetails31) AddCollateralInBalance() *BalanceFormat5Choice {
	c.CollateralInBalance = new(BalanceFormat5Choice)
	return c.CollateralInBalance
}

func (c *CorporateActionBalanceDetails31) AddCollateralOutBalance() *BalanceFormat5Choice {
	c.CollateralOutBalance = new(BalanceFormat5Choice)
	return c.CollateralOutBalance
}

func (c *CorporateActionBalanceDetails31) AddOnLoanBalance() *BalanceFormat5Choice {
	c.OnLoanBalance = new(BalanceFormat5Choice)
	return c.OnLoanBalance
}

func (c *CorporateActionBalanceDetails31) AddPendingDeliveryBalance() *BalanceFormat6Choice {
	newValue := new(BalanceFormat6Choice)
	c.PendingDeliveryBalance = append(c.PendingDeliveryBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails31) AddPendingReceiptBalance() *BalanceFormat6Choice {
	newValue := new(BalanceFormat6Choice)
	c.PendingReceiptBalance = append(c.PendingReceiptBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails31) AddOutForRegistrationBalance() *BalanceFormat5Choice {
	c.OutForRegistrationBalance = new(BalanceFormat5Choice)
	return c.OutForRegistrationBalance
}

func (c *CorporateActionBalanceDetails31) AddSettlementPositionBalance() *BalanceFormat6Choice {
	newValue := new(BalanceFormat6Choice)
	c.SettlementPositionBalance = append(c.SettlementPositionBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails31) AddStreetPositionBalance() *BalanceFormat5Choice {
	c.StreetPositionBalance = new(BalanceFormat5Choice)
	return c.StreetPositionBalance
}

func (c *CorporateActionBalanceDetails31) AddTradeDatePositionBalance() *BalanceFormat5Choice {
	c.TradeDatePositionBalance = new(BalanceFormat5Choice)
	return c.TradeDatePositionBalance
}

func (c *CorporateActionBalanceDetails31) AddInTransshipmentBalance() *BalanceFormat5Choice {
	c.InTransshipmentBalance = new(BalanceFormat5Choice)
	return c.InTransshipmentBalance
}

func (c *CorporateActionBalanceDetails31) AddRegisteredBalance() *BalanceFormat5Choice {
	c.RegisteredBalance = new(BalanceFormat5Choice)
	return c.RegisteredBalance
}

func (c *CorporateActionBalanceDetails31) AddAffectedBalance() *BalanceFormat5Choice {
	c.AffectedBalance = new(BalanceFormat5Choice)
	return c.AffectedBalance
}

func (c *CorporateActionBalanceDetails31) AddUnaffectedBalance() *BalanceFormat5Choice {
	c.UnaffectedBalance = new(BalanceFormat5Choice)
	return c.UnaffectedBalance
}

// Eligible and not eligible balance of securities for a corporate action event.
type CorporateActionBalanceDetails32 struct {

	// Total quantity of financial instruments of the balance.
	TotalEligibleBalance *Quantity17Choice `xml:"TtlElgblBal,omitempty"`

	// Balance of financial instruments that are blocked.
	BlockedBalance *BalanceFormat5Choice `xml:"BlckdBal,omitempty"`

	// Balance of financial instruments that have been borrowed from another party.
	BorrowedBalance *BalanceFormat5Choice `xml:"BrrwdBal,omitempty"`

	// Balance of securities that belong to a third party and that are held for the purpose of collateralisation.
	CollateralInBalance *BalanceFormat5Choice `xml:"CollInBal,omitempty"`

	// Balance of securities that belong to the safekeeping account indicated within this message, and are deposited with a third party for the purpose of collateralisation.
	CollateralOutBalance *BalanceFormat5Choice `xml:"CollOutBal,omitempty"`

	// Balance of financial instruments that have been loaned to a third party.
	OnLoanBalance *BalanceFormat5Choice `xml:"OnLnBal,omitempty"`

	// Balance of financial instruments that are pending delivery.
	PendingDeliveryBalance []*BalanceFormat5Choice `xml:"PdgDlvryBal,omitempty"`

	// Balance of financial instruments that are pending receipt.
	PendingReceiptBalance []*BalanceFormat5Choice `xml:"PdgRctBal,omitempty"`

	// Balance of financial instruments currently being processed by the institution responsible for registering the new beneficial owner (or nominee).
	OutForRegistrationBalance *BalanceFormat5Choice `xml:"OutForRegnBal,omitempty"`

	// Balance of securities representing only settled transactions; pending transactions not included.
	SettlementPositionBalance *BalanceFormat5Choice `xml:"SttlmPosBal,omitempty"`

	// Balance of financial instruments that remain registered in the name of the prior beneficial owner.
	StreetPositionBalance *BalanceFormat5Choice `xml:"StrtPosBal,omitempty"`

	// Balance of securities based on trade date, for example, includes all pending transactions in addition to the balance of settled transactions.
	TradeDatePositionBalance *BalanceFormat5Choice `xml:"TradDtPosBal,omitempty"`

	// Balance of physical securities that are in the process of being transferred from one depository/agent to another.
	InTransshipmentBalance *BalanceFormat5Choice `xml:"InTrnsShipmntBal,omitempty"`

	// Balance of financial instruments that are registered (in the name of a nominee name or of the beneficial owner).
	RegisteredBalance *BalanceFormat5Choice `xml:"RegdBal,omitempty"`
}

func (c *CorporateActionBalanceDetails32) AddTotalEligibleBalance() *Quantity17Choice {
	c.TotalEligibleBalance = new(Quantity17Choice)
	return c.TotalEligibleBalance
}

func (c *CorporateActionBalanceDetails32) AddBlockedBalance() *BalanceFormat5Choice {
	c.BlockedBalance = new(BalanceFormat5Choice)
	return c.BlockedBalance
}

func (c *CorporateActionBalanceDetails32) AddBorrowedBalance() *BalanceFormat5Choice {
	c.BorrowedBalance = new(BalanceFormat5Choice)
	return c.BorrowedBalance
}

func (c *CorporateActionBalanceDetails32) AddCollateralInBalance() *BalanceFormat5Choice {
	c.CollateralInBalance = new(BalanceFormat5Choice)
	return c.CollateralInBalance
}

func (c *CorporateActionBalanceDetails32) AddCollateralOutBalance() *BalanceFormat5Choice {
	c.CollateralOutBalance = new(BalanceFormat5Choice)
	return c.CollateralOutBalance
}

func (c *CorporateActionBalanceDetails32) AddOnLoanBalance() *BalanceFormat5Choice {
	c.OnLoanBalance = new(BalanceFormat5Choice)
	return c.OnLoanBalance
}

func (c *CorporateActionBalanceDetails32) AddPendingDeliveryBalance() *BalanceFormat5Choice {
	newValue := new(BalanceFormat5Choice)
	c.PendingDeliveryBalance = append(c.PendingDeliveryBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails32) AddPendingReceiptBalance() *BalanceFormat5Choice {
	newValue := new(BalanceFormat5Choice)
	c.PendingReceiptBalance = append(c.PendingReceiptBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails32) AddOutForRegistrationBalance() *BalanceFormat5Choice {
	c.OutForRegistrationBalance = new(BalanceFormat5Choice)
	return c.OutForRegistrationBalance
}

func (c *CorporateActionBalanceDetails32) AddSettlementPositionBalance() *BalanceFormat5Choice {
	c.SettlementPositionBalance = new(BalanceFormat5Choice)
	return c.SettlementPositionBalance
}

func (c *CorporateActionBalanceDetails32) AddStreetPositionBalance() *BalanceFormat5Choice {
	c.StreetPositionBalance = new(BalanceFormat5Choice)
	return c.StreetPositionBalance
}

func (c *CorporateActionBalanceDetails32) AddTradeDatePositionBalance() *BalanceFormat5Choice {
	c.TradeDatePositionBalance = new(BalanceFormat5Choice)
	return c.TradeDatePositionBalance
}

func (c *CorporateActionBalanceDetails32) AddInTransshipmentBalance() *BalanceFormat5Choice {
	c.InTransshipmentBalance = new(BalanceFormat5Choice)
	return c.InTransshipmentBalance
}

func (c *CorporateActionBalanceDetails32) AddRegisteredBalance() *BalanceFormat5Choice {
	c.RegisteredBalance = new(BalanceFormat5Choice)
	return c.RegisteredBalance
}

// Eligible and not eligible balance of securities for a corporate action event.
type CorporateActionBalanceDetails33 struct {

	// Total quantity of financial instruments of the balance.
	TotalEligibleBalance *Quantity22Choice `xml:"TtlElgblBal"`

	// Quantity of securities in the sub-balance.
	UninstructedBalance *BalanceFormat7Choice `xml:"UinstdBal"`

	// Provides information about the total instructed balance.
	TotalInstructedBalanceDetails *InstructedBalanceDetails6 `xml:"TtlInstdBalDtls"`

	// Balance of financial instruments that are blocked.
	BlockedBalance *SignedQuantityFormat9 `xml:"BlckdBal,omitempty"`

	// Balance of financial instruments that have been borrowed from another party.
	BorrowedBalance *SignedQuantityFormat9 `xml:"BrrwdBal,omitempty"`

	// Balance of securities that belong to a third party and that are held for the purpose of collateralisation.
	CollateralInBalance *SignedQuantityFormat9 `xml:"CollInBal,omitempty"`

	// Balance of securities that belong to the safekeeping account indicated within this message, and are deposited with a third party for the purpose of collateralisation.
	CollateralOutBalance *SignedQuantityFormat9 `xml:"CollOutBal,omitempty"`

	// Balance of financial instruments that have been loaned to a third party.
	OnLoanBalance *SignedQuantityFormat9 `xml:"OnLnBal,omitempty"`

	// Balance of financial instruments currently being processed by the institution responsible for registering the new beneficial owner (or nominee).
	OutForRegistrationBalance *SignedQuantityFormat9 `xml:"OutForRegnBal,omitempty"`

	// Balance of securities representing only settled transactions; pending transactions not included.
	SettlementPositionBalance *SignedQuantityFormat9 `xml:"SttlmPosBal,omitempty"`

	// Balance of financial instruments that remain registered in the name of the prior beneficial owner.
	StreetPositionBalance *SignedQuantityFormat9 `xml:"StrtPosBal,omitempty"`

	// Balance of securities based on trade date, for example, includes all pending transactions in addition to the balance of settled transactions.
	TradeDatePositionBalance *SignedQuantityFormat9 `xml:"TradDtPosBal,omitempty"`

	// Balance of physical securities that are in the process of being transferred from one depository/agent to another.
	InTransshipmentBalance *SignedQuantityFormat9 `xml:"InTrnsShipmntBal,omitempty"`

	// Balance of financial instruments that are registered (in the name of a nominee name or of the beneficial owner).
	RegisteredBalance *SignedQuantityFormat9 `xml:"RegdBal,omitempty"`

	// Position that account holders should return to the account servicer to participate in the event or to fulfil their obligation for the event to be complete, for example, return of securities for late announced drawing.
	ObligatedBalance *SignedQuantityFormat9 `xml:"OblgtdBal,omitempty"`

	// Balance of financial instruments that are pending delivery.
	PendingDeliveryBalance []*PendingBalance4 `xml:"PdgDlvryBal,omitempty"`

	// Balance of financial instruments that are pending receipt.
	PendingReceiptBalance []*PendingBalance4 `xml:"PdgRctBal,omitempty"`
}

func (c *CorporateActionBalanceDetails33) AddTotalEligibleBalance() *Quantity22Choice {
	c.TotalEligibleBalance = new(Quantity22Choice)
	return c.TotalEligibleBalance
}

func (c *CorporateActionBalanceDetails33) AddUninstructedBalance() *BalanceFormat7Choice {
	c.UninstructedBalance = new(BalanceFormat7Choice)
	return c.UninstructedBalance
}

func (c *CorporateActionBalanceDetails33) AddTotalInstructedBalanceDetails() *InstructedBalanceDetails6 {
	c.TotalInstructedBalanceDetails = new(InstructedBalanceDetails6)
	return c.TotalInstructedBalanceDetails
}

func (c *CorporateActionBalanceDetails33) AddBlockedBalance() *SignedQuantityFormat9 {
	c.BlockedBalance = new(SignedQuantityFormat9)
	return c.BlockedBalance
}

func (c *CorporateActionBalanceDetails33) AddBorrowedBalance() *SignedQuantityFormat9 {
	c.BorrowedBalance = new(SignedQuantityFormat9)
	return c.BorrowedBalance
}

func (c *CorporateActionBalanceDetails33) AddCollateralInBalance() *SignedQuantityFormat9 {
	c.CollateralInBalance = new(SignedQuantityFormat9)
	return c.CollateralInBalance
}

func (c *CorporateActionBalanceDetails33) AddCollateralOutBalance() *SignedQuantityFormat9 {
	c.CollateralOutBalance = new(SignedQuantityFormat9)
	return c.CollateralOutBalance
}

func (c *CorporateActionBalanceDetails33) AddOnLoanBalance() *SignedQuantityFormat9 {
	c.OnLoanBalance = new(SignedQuantityFormat9)
	return c.OnLoanBalance
}

func (c *CorporateActionBalanceDetails33) AddOutForRegistrationBalance() *SignedQuantityFormat9 {
	c.OutForRegistrationBalance = new(SignedQuantityFormat9)
	return c.OutForRegistrationBalance
}

func (c *CorporateActionBalanceDetails33) AddSettlementPositionBalance() *SignedQuantityFormat9 {
	c.SettlementPositionBalance = new(SignedQuantityFormat9)
	return c.SettlementPositionBalance
}

func (c *CorporateActionBalanceDetails33) AddStreetPositionBalance() *SignedQuantityFormat9 {
	c.StreetPositionBalance = new(SignedQuantityFormat9)
	return c.StreetPositionBalance
}

func (c *CorporateActionBalanceDetails33) AddTradeDatePositionBalance() *SignedQuantityFormat9 {
	c.TradeDatePositionBalance = new(SignedQuantityFormat9)
	return c.TradeDatePositionBalance
}

func (c *CorporateActionBalanceDetails33) AddInTransshipmentBalance() *SignedQuantityFormat9 {
	c.InTransshipmentBalance = new(SignedQuantityFormat9)
	return c.InTransshipmentBalance
}

func (c *CorporateActionBalanceDetails33) AddRegisteredBalance() *SignedQuantityFormat9 {
	c.RegisteredBalance = new(SignedQuantityFormat9)
	return c.RegisteredBalance
}

func (c *CorporateActionBalanceDetails33) AddObligatedBalance() *SignedQuantityFormat9 {
	c.ObligatedBalance = new(SignedQuantityFormat9)
	return c.ObligatedBalance
}

func (c *CorporateActionBalanceDetails33) AddPendingDeliveryBalance() *PendingBalance4 {
	newValue := new(PendingBalance4)
	c.PendingDeliveryBalance = append(c.PendingDeliveryBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails33) AddPendingReceiptBalance() *PendingBalance4 {
	newValue := new(PendingBalance4)
	c.PendingReceiptBalance = append(c.PendingReceiptBalance, newValue)
	return newValue
}

// Eligible and not eligible balance of securities for a corporate action event.
type CorporateActionBalanceDetails34 struct {

	// Total quantity of financial instruments of the balance.
	TotalEligibleBalance *Quantity22Choice `xml:"TtlElgblBal,omitempty"`

	// Balance of financial instruments that are blocked.
	BlockedBalance *BalanceFormat7Choice `xml:"BlckdBal,omitempty"`

	// Balance of financial instruments that have been borrowed from another party.
	BorrowedBalance *BalanceFormat7Choice `xml:"BrrwdBal,omitempty"`

	// Balance of securities that belong to a third party and that are held for the purpose of collateralisation.
	CollateralInBalance *BalanceFormat7Choice `xml:"CollInBal,omitempty"`

	// Balance of securities that belong to the safekeeping account indicated within this message, and are deposited with a third party for the purpose of collateralisation.
	CollateralOutBalance *BalanceFormat7Choice `xml:"CollOutBal,omitempty"`

	// Balance of financial instruments that have been loaned to a third party.
	OnLoanBalance *BalanceFormat7Choice `xml:"OnLnBal,omitempty"`

	// Balance of financial instruments that are pending delivery.
	PendingDeliveryBalance []*BalanceFormat7Choice `xml:"PdgDlvryBal,omitempty"`

	// Balance of financial instruments that are pending receipt.
	PendingReceiptBalance []*BalanceFormat7Choice `xml:"PdgRctBal,omitempty"`

	// Balance of financial instruments currently being processed by the institution responsible for registering the new beneficial owner (or nominee).
	OutForRegistrationBalance *BalanceFormat7Choice `xml:"OutForRegnBal,omitempty"`

	// Balance of securities representing only settled transactions; pending transactions not included.
	SettlementPositionBalance *BalanceFormat7Choice `xml:"SttlmPosBal,omitempty"`

	// Balance of financial instruments that remain registered in the name of the prior beneficial owner.
	StreetPositionBalance *BalanceFormat7Choice `xml:"StrtPosBal,omitempty"`

	// Balance of securities based on trade date, for example, includes all pending transactions in addition to the balance of settled transactions.
	TradeDatePositionBalance *BalanceFormat7Choice `xml:"TradDtPosBal,omitempty"`

	// Balance of physical securities that are in the process of being transferred from one depository/agent to another.
	InTransshipmentBalance *BalanceFormat7Choice `xml:"InTrnsShipmntBal,omitempty"`

	// Balance of financial instruments that are registered (in the name of a nominee name or of the beneficial owner).
	RegisteredBalance *BalanceFormat7Choice `xml:"RegdBal,omitempty"`
}

func (c *CorporateActionBalanceDetails34) AddTotalEligibleBalance() *Quantity22Choice {
	c.TotalEligibleBalance = new(Quantity22Choice)
	return c.TotalEligibleBalance
}

func (c *CorporateActionBalanceDetails34) AddBlockedBalance() *BalanceFormat7Choice {
	c.BlockedBalance = new(BalanceFormat7Choice)
	return c.BlockedBalance
}

func (c *CorporateActionBalanceDetails34) AddBorrowedBalance() *BalanceFormat7Choice {
	c.BorrowedBalance = new(BalanceFormat7Choice)
	return c.BorrowedBalance
}

func (c *CorporateActionBalanceDetails34) AddCollateralInBalance() *BalanceFormat7Choice {
	c.CollateralInBalance = new(BalanceFormat7Choice)
	return c.CollateralInBalance
}

func (c *CorporateActionBalanceDetails34) AddCollateralOutBalance() *BalanceFormat7Choice {
	c.CollateralOutBalance = new(BalanceFormat7Choice)
	return c.CollateralOutBalance
}

func (c *CorporateActionBalanceDetails34) AddOnLoanBalance() *BalanceFormat7Choice {
	c.OnLoanBalance = new(BalanceFormat7Choice)
	return c.OnLoanBalance
}

func (c *CorporateActionBalanceDetails34) AddPendingDeliveryBalance() *BalanceFormat7Choice {
	newValue := new(BalanceFormat7Choice)
	c.PendingDeliveryBalance = append(c.PendingDeliveryBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails34) AddPendingReceiptBalance() *BalanceFormat7Choice {
	newValue := new(BalanceFormat7Choice)
	c.PendingReceiptBalance = append(c.PendingReceiptBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails34) AddOutForRegistrationBalance() *BalanceFormat7Choice {
	c.OutForRegistrationBalance = new(BalanceFormat7Choice)
	return c.OutForRegistrationBalance
}

func (c *CorporateActionBalanceDetails34) AddSettlementPositionBalance() *BalanceFormat7Choice {
	c.SettlementPositionBalance = new(BalanceFormat7Choice)
	return c.SettlementPositionBalance
}

func (c *CorporateActionBalanceDetails34) AddStreetPositionBalance() *BalanceFormat7Choice {
	c.StreetPositionBalance = new(BalanceFormat7Choice)
	return c.StreetPositionBalance
}

func (c *CorporateActionBalanceDetails34) AddTradeDatePositionBalance() *BalanceFormat7Choice {
	c.TradeDatePositionBalance = new(BalanceFormat7Choice)
	return c.TradeDatePositionBalance
}

func (c *CorporateActionBalanceDetails34) AddInTransshipmentBalance() *BalanceFormat7Choice {
	c.InTransshipmentBalance = new(BalanceFormat7Choice)
	return c.InTransshipmentBalance
}

func (c *CorporateActionBalanceDetails34) AddRegisteredBalance() *BalanceFormat7Choice {
	c.RegisteredBalance = new(BalanceFormat7Choice)
	return c.RegisteredBalance
}

// Eligible and not eligible balance of securities for a corporate action event.
type CorporateActionBalanceDetails35 struct {

	// Balance to which the payment applies (less or equal to the total eligible balance).
	ConfirmedBalance *BalanceFormat7Choice `xml:"ConfdBal"`

	// Total balance of securities eligible for this corporate action event. The entitlement calculation is based on this balance.
	TotalEligibleBalance *TotalEligibleBalanceFormat9 `xml:"TtlElgblBal,omitempty"`

	// Balance of financial instruments that are blocked.
	BlockedBalance *BalanceFormat7Choice `xml:"BlckdBal,omitempty"`

	// Balance of financial instruments that have been borrowed from another party.
	BorrowedBalance *BalanceFormat7Choice `xml:"BrrwdBal,omitempty"`

	// Balance of securities that belong to a third party and that are held for the purpose of collateralisation.
	CollateralInBalance *BalanceFormat7Choice `xml:"CollInBal,omitempty"`

	// Balance of securities that belong to the safekeeping account indicated within this message, and are deposited with a third party for the purpose of collateralisation.
	CollateralOutBalance *BalanceFormat7Choice `xml:"CollOutBal,omitempty"`

	// Balance of financial instruments that have been loaned to a third party.
	OnLoanBalance *BalanceFormat7Choice `xml:"OnLnBal,omitempty"`

	// Balance of financial instruments that are pending delivery.
	PendingDeliveryBalance []*BalanceFormat10Choice `xml:"PdgDlvryBal,omitempty"`

	// Balance of financial instruments that are pending receipt.
	PendingReceiptBalance []*BalanceFormat10Choice `xml:"PdgRctBal,omitempty"`

	// Balance of financial instruments currently being processed by the institution responsible for registering the new beneficial owner (or nominee).
	OutForRegistrationBalance *BalanceFormat7Choice `xml:"OutForRegnBal,omitempty"`

	// Balance of securities representing only settled transactions; pending transactions not included.
	SettlementPositionBalance []*BalanceFormat10Choice `xml:"SttlmPosBal,omitempty"`

	// Balance of financial instruments that remain registered in the name of the prior beneficial owner.
	StreetPositionBalance *BalanceFormat7Choice `xml:"StrtPosBal,omitempty"`

	// Balance of securities based on trade date, for example, includes all pending transactions in addition to the balance of settled transactions.
	TradeDatePositionBalance *BalanceFormat7Choice `xml:"TradDtPosBal,omitempty"`

	// Balance of physical securities that are in the process of being transferred from one depository/agent to another.
	InTransshipmentBalance *BalanceFormat7Choice `xml:"InTrnsShipmntBal,omitempty"`

	// Balance of financial instruments that are registered (in the name of a nominee name or of the beneficial owner).
	RegisteredBalance *BalanceFormat7Choice `xml:"RegdBal,omitempty"`

	// Balance that has been affected by the process run through the event.
	AffectedBalance *BalanceFormat7Choice `xml:"AfctdBal,omitempty"`

	// Balance that has not been affected by the process run through the event.
	UnaffectedBalance *BalanceFormat7Choice `xml:"UafctdBal,omitempty"`
}

func (c *CorporateActionBalanceDetails35) AddConfirmedBalance() *BalanceFormat7Choice {
	c.ConfirmedBalance = new(BalanceFormat7Choice)
	return c.ConfirmedBalance
}

func (c *CorporateActionBalanceDetails35) AddTotalEligibleBalance() *TotalEligibleBalanceFormat9 {
	c.TotalEligibleBalance = new(TotalEligibleBalanceFormat9)
	return c.TotalEligibleBalance
}

func (c *CorporateActionBalanceDetails35) AddBlockedBalance() *BalanceFormat7Choice {
	c.BlockedBalance = new(BalanceFormat7Choice)
	return c.BlockedBalance
}

func (c *CorporateActionBalanceDetails35) AddBorrowedBalance() *BalanceFormat7Choice {
	c.BorrowedBalance = new(BalanceFormat7Choice)
	return c.BorrowedBalance
}

func (c *CorporateActionBalanceDetails35) AddCollateralInBalance() *BalanceFormat7Choice {
	c.CollateralInBalance = new(BalanceFormat7Choice)
	return c.CollateralInBalance
}

func (c *CorporateActionBalanceDetails35) AddCollateralOutBalance() *BalanceFormat7Choice {
	c.CollateralOutBalance = new(BalanceFormat7Choice)
	return c.CollateralOutBalance
}

func (c *CorporateActionBalanceDetails35) AddOnLoanBalance() *BalanceFormat7Choice {
	c.OnLoanBalance = new(BalanceFormat7Choice)
	return c.OnLoanBalance
}

func (c *CorporateActionBalanceDetails35) AddPendingDeliveryBalance() *BalanceFormat10Choice {
	newValue := new(BalanceFormat10Choice)
	c.PendingDeliveryBalance = append(c.PendingDeliveryBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails35) AddPendingReceiptBalance() *BalanceFormat10Choice {
	newValue := new(BalanceFormat10Choice)
	c.PendingReceiptBalance = append(c.PendingReceiptBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails35) AddOutForRegistrationBalance() *BalanceFormat7Choice {
	c.OutForRegistrationBalance = new(BalanceFormat7Choice)
	return c.OutForRegistrationBalance
}

func (c *CorporateActionBalanceDetails35) AddSettlementPositionBalance() *BalanceFormat10Choice {
	newValue := new(BalanceFormat10Choice)
	c.SettlementPositionBalance = append(c.SettlementPositionBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails35) AddStreetPositionBalance() *BalanceFormat7Choice {
	c.StreetPositionBalance = new(BalanceFormat7Choice)
	return c.StreetPositionBalance
}

func (c *CorporateActionBalanceDetails35) AddTradeDatePositionBalance() *BalanceFormat7Choice {
	c.TradeDatePositionBalance = new(BalanceFormat7Choice)
	return c.TradeDatePositionBalance
}

func (c *CorporateActionBalanceDetails35) AddInTransshipmentBalance() *BalanceFormat7Choice {
	c.InTransshipmentBalance = new(BalanceFormat7Choice)
	return c.InTransshipmentBalance
}

func (c *CorporateActionBalanceDetails35) AddRegisteredBalance() *BalanceFormat7Choice {
	c.RegisteredBalance = new(BalanceFormat7Choice)
	return c.RegisteredBalance
}

func (c *CorporateActionBalanceDetails35) AddAffectedBalance() *BalanceFormat7Choice {
	c.AffectedBalance = new(BalanceFormat7Choice)
	return c.AffectedBalance
}

func (c *CorporateActionBalanceDetails35) AddUnaffectedBalance() *BalanceFormat7Choice {
	c.UnaffectedBalance = new(BalanceFormat7Choice)
	return c.UnaffectedBalance
}

// Eligible and not eligible balance of securities for a corporate action event.
type CorporateActionBalanceDetails36 struct {

	// Total balance of securities eligible for this corporate action event. The entitlement calculation is based on this balance.
	TotalEligibleBalance *TotalEligibleBalanceFormat9 `xml:"TtlElgblBal,omitempty"`

	// Balance of financial instruments that are blocked.
	BlockedBalance *BalanceFormat7Choice `xml:"BlckdBal,omitempty"`

	// Balance of financial instruments that have been borrowed from another party.
	BorrowedBalance *BalanceFormat7Choice `xml:"BrrwdBal,omitempty"`

	// Balance of securities that belong to a third party and that are held for the purpose of collateralisation.
	CollateralInBalance *BalanceFormat7Choice `xml:"CollInBal,omitempty"`

	// Balance of securities that belong to the safekeeping account indicated within this message, and are deposited with a third party for the purpose of collateralisation.
	CollateralOutBalance *BalanceFormat7Choice `xml:"CollOutBal,omitempty"`

	// Balance of financial instruments that have been loaned to a third party.
	OnLoanBalance *BalanceFormat7Choice `xml:"OnLnBal,omitempty"`

	// Balance of financial instruments that are pending delivery.
	PendingDeliveryBalance []*BalanceFormat10Choice `xml:"PdgDlvryBal,omitempty"`

	// Balance of financial instruments that are pending receipt.
	PendingReceiptBalance []*BalanceFormat10Choice `xml:"PdgRctBal,omitempty"`

	// Balance of financial instruments currently being processed by the institution responsible for registering the new beneficial owner (or nominee).
	OutForRegistrationBalance *BalanceFormat7Choice `xml:"OutForRegnBal,omitempty"`

	// Balance of securities representing only settled transactions; pending transactions not included.
	SettlementPositionBalance []*BalanceFormat10Choice `xml:"SttlmPosBal,omitempty"`

	// Balance of financial instruments that remain registered in the name of the prior beneficial owner.
	StreetPositionBalance *BalanceFormat7Choice `xml:"StrtPosBal,omitempty"`

	// Balance of securities based on trade date, for example, includes all pending transactions in addition to the balance of settled transactions.
	TradeDatePositionBalance *BalanceFormat7Choice `xml:"TradDtPosBal,omitempty"`

	// Balance of physical securities that are in the process of being transferred from one depository/agent to another.
	InTransshipmentBalance *BalanceFormat7Choice `xml:"InTrnsShipmntBal,omitempty"`

	// Balance of financial instruments that are registered (in the name of a nominee name or of the beneficial owner).
	RegisteredBalance *BalanceFormat7Choice `xml:"RegdBal,omitempty"`

	// Position that account holders should return to the account servicer to participate in the event or to fulfil their obligation for the event to be complete, for example, return of securities for late announced drawing.
	ObligatedBalance *BalanceFormat7Choice `xml:"OblgtdBal,omitempty"`

	// Balance of uninstructed position.
	UninstructedBalance *BalanceFormat7Choice `xml:"UinstdBal,omitempty"`

	// Balance of instructed position.
	InstructedBalance *BalanceFormat7Choice `xml:"InstdBal,omitempty"`

	// Balance that has been affected by the process run through the event.
	AffectedBalance *BalanceFormat7Choice `xml:"AfctdBal,omitempty"`

	// Balance that has not been affected by the process run through the event.
	UnaffectedBalance *BalanceFormat7Choice `xml:"UafctdBal,omitempty"`
}

func (c *CorporateActionBalanceDetails36) AddTotalEligibleBalance() *TotalEligibleBalanceFormat9 {
	c.TotalEligibleBalance = new(TotalEligibleBalanceFormat9)
	return c.TotalEligibleBalance
}

func (c *CorporateActionBalanceDetails36) AddBlockedBalance() *BalanceFormat7Choice {
	c.BlockedBalance = new(BalanceFormat7Choice)
	return c.BlockedBalance
}

func (c *CorporateActionBalanceDetails36) AddBorrowedBalance() *BalanceFormat7Choice {
	c.BorrowedBalance = new(BalanceFormat7Choice)
	return c.BorrowedBalance
}

func (c *CorporateActionBalanceDetails36) AddCollateralInBalance() *BalanceFormat7Choice {
	c.CollateralInBalance = new(BalanceFormat7Choice)
	return c.CollateralInBalance
}

func (c *CorporateActionBalanceDetails36) AddCollateralOutBalance() *BalanceFormat7Choice {
	c.CollateralOutBalance = new(BalanceFormat7Choice)
	return c.CollateralOutBalance
}

func (c *CorporateActionBalanceDetails36) AddOnLoanBalance() *BalanceFormat7Choice {
	c.OnLoanBalance = new(BalanceFormat7Choice)
	return c.OnLoanBalance
}

func (c *CorporateActionBalanceDetails36) AddPendingDeliveryBalance() *BalanceFormat10Choice {
	newValue := new(BalanceFormat10Choice)
	c.PendingDeliveryBalance = append(c.PendingDeliveryBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails36) AddPendingReceiptBalance() *BalanceFormat10Choice {
	newValue := new(BalanceFormat10Choice)
	c.PendingReceiptBalance = append(c.PendingReceiptBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails36) AddOutForRegistrationBalance() *BalanceFormat7Choice {
	c.OutForRegistrationBalance = new(BalanceFormat7Choice)
	return c.OutForRegistrationBalance
}

func (c *CorporateActionBalanceDetails36) AddSettlementPositionBalance() *BalanceFormat10Choice {
	newValue := new(BalanceFormat10Choice)
	c.SettlementPositionBalance = append(c.SettlementPositionBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails36) AddStreetPositionBalance() *BalanceFormat7Choice {
	c.StreetPositionBalance = new(BalanceFormat7Choice)
	return c.StreetPositionBalance
}

func (c *CorporateActionBalanceDetails36) AddTradeDatePositionBalance() *BalanceFormat7Choice {
	c.TradeDatePositionBalance = new(BalanceFormat7Choice)
	return c.TradeDatePositionBalance
}

func (c *CorporateActionBalanceDetails36) AddInTransshipmentBalance() *BalanceFormat7Choice {
	c.InTransshipmentBalance = new(BalanceFormat7Choice)
	return c.InTransshipmentBalance
}

func (c *CorporateActionBalanceDetails36) AddRegisteredBalance() *BalanceFormat7Choice {
	c.RegisteredBalance = new(BalanceFormat7Choice)
	return c.RegisteredBalance
}

func (c *CorporateActionBalanceDetails36) AddObligatedBalance() *BalanceFormat7Choice {
	c.ObligatedBalance = new(BalanceFormat7Choice)
	return c.ObligatedBalance
}

func (c *CorporateActionBalanceDetails36) AddUninstructedBalance() *BalanceFormat7Choice {
	c.UninstructedBalance = new(BalanceFormat7Choice)
	return c.UninstructedBalance
}

func (c *CorporateActionBalanceDetails36) AddInstructedBalance() *BalanceFormat7Choice {
	c.InstructedBalance = new(BalanceFormat7Choice)
	return c.InstructedBalance
}

func (c *CorporateActionBalanceDetails36) AddAffectedBalance() *BalanceFormat7Choice {
	c.AffectedBalance = new(BalanceFormat7Choice)
	return c.AffectedBalance
}

func (c *CorporateActionBalanceDetails36) AddUnaffectedBalance() *BalanceFormat7Choice {
	c.UnaffectedBalance = new(BalanceFormat7Choice)
	return c.UnaffectedBalance
}

// Eligible and not eligible balance of securities for a corporate action event.
type CorporateActionBalanceDetails4 struct {

	// Total quantity of financial instruments of the balance.
	TotalEligibleBalance *Quantity3Choice `xml:"TtlElgblBal"`

	// Quantity of securities in the sub-balance.
	UninstructedBalance *BalanceFormat1Choice `xml:"UinstdBal"`

	// Provides information about the total instructed balance.
	TotalInstructedBalanceDetails *InstructedBalanceDetails1 `xml:"TtlInstdBalDtls"`

	// Balance of financial instruments that are blocked.
	BlockedBalance *SignedQuantityFormat2 `xml:"BlckdBal,omitempty"`

	// Balance of financial instruments that have been borrowed from another party.
	BorrowedBalance *SignedQuantityFormat2 `xml:"BrrwdBal,omitempty"`

	// Balance of securities that belong to a third party and that are held for the purpose of collateralisation.
	CollateralInBalance *SignedQuantityFormat2 `xml:"CollInBal,omitempty"`

	// Balance of securities that belong to the safekeeping account indicated within this message, and are deposited with a third party for the purpose of collateralisation.
	CollateralOutBalance *SignedQuantityFormat2 `xml:"CollOutBal,omitempty"`

	// Balance of financial instruments that have been loaned to a third party.
	OnLoanBalance *SignedQuantityFormat2 `xml:"OnLnBal,omitempty"`

	// Balance of financial instruments currently being processed by the institution responsible for registering the new beneficial owner (or nominee).
	OutForRegistrationBalance *SignedQuantityFormat2 `xml:"OutForRegnBal,omitempty"`

	// Balance of securities representing only settled transactions; pending transactions not included.
	SettlementPositionBalance *SignedQuantityFormat2 `xml:"SttlmPosBal,omitempty"`

	// Balance of financial instruments that remain registered in the name of the prior beneficial owner.
	StreetPositionBalance *SignedQuantityFormat2 `xml:"StrtPosBal,omitempty"`

	// Balance of securities based on trade date, for example, includes all pending transactions in addition to the balance of settled transactions.
	TradeDatePositionBalance *SignedQuantityFormat2 `xml:"TradDtPosBal,omitempty"`

	// Balance of physical securities that are in the process of being transferred from one depository/agent to another.
	InTransshipmentBalance *SignedQuantityFormat2 `xml:"InTrnsShipmntBal,omitempty"`

	// Balance of financial instruments that are registered (in the name of a nominee name or of the beneficial owner).
	RegisteredBalance *SignedQuantityFormat2 `xml:"RegdBal,omitempty"`

	// Position that account holders should return to the account servicer to participate in the event or to fulfil their obligation for the event to be complete, for example, return of securities for late announced drawing.
	ObligatedBalance *SignedQuantityFormat2 `xml:"OblgtdBal,omitempty"`

	// Balance of financial instruments that are pending delivery.
	PendingDeliveryBalance []*PendingBalance1 `xml:"PdgDlvryBal,omitempty"`

	// Balance of financial instruments that are pending receipt.
	PendingReceiptBalance []*PendingBalance1 `xml:"PdgRctBal,omitempty"`
}

func (c *CorporateActionBalanceDetails4) AddTotalEligibleBalance() *Quantity3Choice {
	c.TotalEligibleBalance = new(Quantity3Choice)
	return c.TotalEligibleBalance
}

func (c *CorporateActionBalanceDetails4) AddUninstructedBalance() *BalanceFormat1Choice {
	c.UninstructedBalance = new(BalanceFormat1Choice)
	return c.UninstructedBalance
}

func (c *CorporateActionBalanceDetails4) AddTotalInstructedBalanceDetails() *InstructedBalanceDetails1 {
	c.TotalInstructedBalanceDetails = new(InstructedBalanceDetails1)
	return c.TotalInstructedBalanceDetails
}

func (c *CorporateActionBalanceDetails4) AddBlockedBalance() *SignedQuantityFormat2 {
	c.BlockedBalance = new(SignedQuantityFormat2)
	return c.BlockedBalance
}

func (c *CorporateActionBalanceDetails4) AddBorrowedBalance() *SignedQuantityFormat2 {
	c.BorrowedBalance = new(SignedQuantityFormat2)
	return c.BorrowedBalance
}

func (c *CorporateActionBalanceDetails4) AddCollateralInBalance() *SignedQuantityFormat2 {
	c.CollateralInBalance = new(SignedQuantityFormat2)
	return c.CollateralInBalance
}

func (c *CorporateActionBalanceDetails4) AddCollateralOutBalance() *SignedQuantityFormat2 {
	c.CollateralOutBalance = new(SignedQuantityFormat2)
	return c.CollateralOutBalance
}

func (c *CorporateActionBalanceDetails4) AddOnLoanBalance() *SignedQuantityFormat2 {
	c.OnLoanBalance = new(SignedQuantityFormat2)
	return c.OnLoanBalance
}

func (c *CorporateActionBalanceDetails4) AddOutForRegistrationBalance() *SignedQuantityFormat2 {
	c.OutForRegistrationBalance = new(SignedQuantityFormat2)
	return c.OutForRegistrationBalance
}

func (c *CorporateActionBalanceDetails4) AddSettlementPositionBalance() *SignedQuantityFormat2 {
	c.SettlementPositionBalance = new(SignedQuantityFormat2)
	return c.SettlementPositionBalance
}

func (c *CorporateActionBalanceDetails4) AddStreetPositionBalance() *SignedQuantityFormat2 {
	c.StreetPositionBalance = new(SignedQuantityFormat2)
	return c.StreetPositionBalance
}

func (c *CorporateActionBalanceDetails4) AddTradeDatePositionBalance() *SignedQuantityFormat2 {
	c.TradeDatePositionBalance = new(SignedQuantityFormat2)
	return c.TradeDatePositionBalance
}

func (c *CorporateActionBalanceDetails4) AddInTransshipmentBalance() *SignedQuantityFormat2 {
	c.InTransshipmentBalance = new(SignedQuantityFormat2)
	return c.InTransshipmentBalance
}

func (c *CorporateActionBalanceDetails4) AddRegisteredBalance() *SignedQuantityFormat2 {
	c.RegisteredBalance = new(SignedQuantityFormat2)
	return c.RegisteredBalance
}

func (c *CorporateActionBalanceDetails4) AddObligatedBalance() *SignedQuantityFormat2 {
	c.ObligatedBalance = new(SignedQuantityFormat2)
	return c.ObligatedBalance
}

func (c *CorporateActionBalanceDetails4) AddPendingDeliveryBalance() *PendingBalance1 {
	newValue := new(PendingBalance1)
	c.PendingDeliveryBalance = append(c.PendingDeliveryBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails4) AddPendingReceiptBalance() *PendingBalance1 {
	newValue := new(PendingBalance1)
	c.PendingReceiptBalance = append(c.PendingReceiptBalance, newValue)
	return newValue
}

// Eligible and not eligible balance of securities for a corporate action event.
type CorporateActionBalanceDetails9 struct {

	// Total quantity of financial instruments of the balance.
	TotalEligibleBalance *Quantity3Choice `xml:"TtlElgblBal"`

	// Quantity of securities in the sub-balance.
	UninstructedBalance *BalanceFormat1Choice `xml:"UinstdBal"`

	// Provides information about the total instructed balance.
	TotalInstructedBalanceDetails *InstructedBalanceDetails3 `xml:"TtlInstdBalDtls"`

	// Balance of financial instruments that are blocked.
	BlockedBalance *SignedQuantityFormat2 `xml:"BlckdBal,omitempty"`

	// Balance of financial instruments that have been borrowed from another party.
	BorrowedBalance *SignedQuantityFormat2 `xml:"BrrwdBal,omitempty"`

	// Balance of securities that belong to a third party and that are held for the purpose of collateralisation.
	CollateralInBalance *SignedQuantityFormat2 `xml:"CollInBal,omitempty"`

	// Balance of securities that belong to the safekeeping account indicated within this message, and are deposited with a third party for the purpose of collateralisation.
	CollateralOutBalance *SignedQuantityFormat2 `xml:"CollOutBal,omitempty"`

	// Balance of financial instruments that have been loaned to a third party.
	OnLoanBalance *SignedQuantityFormat2 `xml:"OnLnBal,omitempty"`

	// Balance of financial instruments currently being processed by the institution responsible for registering the new beneficial owner (or nominee).
	OutForRegistrationBalance *SignedQuantityFormat2 `xml:"OutForRegnBal,omitempty"`

	// Balance of securities representing only settled transactions; pending transactions not included.
	SettlementPositionBalance *SignedQuantityFormat2 `xml:"SttlmPosBal,omitempty"`

	// Balance of financial instruments that remain registered in the name of the prior beneficial owner.
	StreetPositionBalance *SignedQuantityFormat2 `xml:"StrtPosBal,omitempty"`

	// Balance of securities based on trade date, for example, includes all pending transactions in addition to the balance of settled transactions.
	TradeDatePositionBalance *SignedQuantityFormat2 `xml:"TradDtPosBal,omitempty"`

	// Balance of physical securities that are in the process of being transferred from one depository/agent to another.
	InTransshipmentBalance *SignedQuantityFormat2 `xml:"InTrnsShipmntBal,omitempty"`

	// Balance of financial instruments that are registered (in the name of a nominee name or of the beneficial owner).
	RegisteredBalance *SignedQuantityFormat2 `xml:"RegdBal,omitempty"`

	// Position that account holders should return to the account servicer to participate in the event or to fulfil their obligation for the event to be complete, for example, return of securities for late announced drawing.
	ObligatedBalance *SignedQuantityFormat2 `xml:"OblgtdBal,omitempty"`

	// Balance of financial instruments that are pending delivery.
	PendingDeliveryBalance []*PendingBalance1 `xml:"PdgDlvryBal,omitempty"`

	// Balance of financial instruments that are pending receipt.
	PendingReceiptBalance []*PendingBalance1 `xml:"PdgRctBal,omitempty"`
}

func (c *CorporateActionBalanceDetails9) AddTotalEligibleBalance() *Quantity3Choice {
	c.TotalEligibleBalance = new(Quantity3Choice)
	return c.TotalEligibleBalance
}

func (c *CorporateActionBalanceDetails9) AddUninstructedBalance() *BalanceFormat1Choice {
	c.UninstructedBalance = new(BalanceFormat1Choice)
	return c.UninstructedBalance
}

func (c *CorporateActionBalanceDetails9) AddTotalInstructedBalanceDetails() *InstructedBalanceDetails3 {
	c.TotalInstructedBalanceDetails = new(InstructedBalanceDetails3)
	return c.TotalInstructedBalanceDetails
}

func (c *CorporateActionBalanceDetails9) AddBlockedBalance() *SignedQuantityFormat2 {
	c.BlockedBalance = new(SignedQuantityFormat2)
	return c.BlockedBalance
}

func (c *CorporateActionBalanceDetails9) AddBorrowedBalance() *SignedQuantityFormat2 {
	c.BorrowedBalance = new(SignedQuantityFormat2)
	return c.BorrowedBalance
}

func (c *CorporateActionBalanceDetails9) AddCollateralInBalance() *SignedQuantityFormat2 {
	c.CollateralInBalance = new(SignedQuantityFormat2)
	return c.CollateralInBalance
}

func (c *CorporateActionBalanceDetails9) AddCollateralOutBalance() *SignedQuantityFormat2 {
	c.CollateralOutBalance = new(SignedQuantityFormat2)
	return c.CollateralOutBalance
}

func (c *CorporateActionBalanceDetails9) AddOnLoanBalance() *SignedQuantityFormat2 {
	c.OnLoanBalance = new(SignedQuantityFormat2)
	return c.OnLoanBalance
}

func (c *CorporateActionBalanceDetails9) AddOutForRegistrationBalance() *SignedQuantityFormat2 {
	c.OutForRegistrationBalance = new(SignedQuantityFormat2)
	return c.OutForRegistrationBalance
}

func (c *CorporateActionBalanceDetails9) AddSettlementPositionBalance() *SignedQuantityFormat2 {
	c.SettlementPositionBalance = new(SignedQuantityFormat2)
	return c.SettlementPositionBalance
}

func (c *CorporateActionBalanceDetails9) AddStreetPositionBalance() *SignedQuantityFormat2 {
	c.StreetPositionBalance = new(SignedQuantityFormat2)
	return c.StreetPositionBalance
}

func (c *CorporateActionBalanceDetails9) AddTradeDatePositionBalance() *SignedQuantityFormat2 {
	c.TradeDatePositionBalance = new(SignedQuantityFormat2)
	return c.TradeDatePositionBalance
}

func (c *CorporateActionBalanceDetails9) AddInTransshipmentBalance() *SignedQuantityFormat2 {
	c.InTransshipmentBalance = new(SignedQuantityFormat2)
	return c.InTransshipmentBalance
}

func (c *CorporateActionBalanceDetails9) AddRegisteredBalance() *SignedQuantityFormat2 {
	c.RegisteredBalance = new(SignedQuantityFormat2)
	return c.RegisteredBalance
}

func (c *CorporateActionBalanceDetails9) AddObligatedBalance() *SignedQuantityFormat2 {
	c.ObligatedBalance = new(SignedQuantityFormat2)
	return c.ObligatedBalance
}

func (c *CorporateActionBalanceDetails9) AddPendingDeliveryBalance() *PendingBalance1 {
	newValue := new(PendingBalance1)
	c.PendingDeliveryBalance = append(c.PendingDeliveryBalance, newValue)
	return newValue
}

func (c *CorporateActionBalanceDetails9) AddPendingReceiptBalance() *PendingBalance1 {
	newValue := new(PendingBalance1)
	c.PendingReceiptBalance = append(c.PendingReceiptBalance, newValue)
	return newValue
}
