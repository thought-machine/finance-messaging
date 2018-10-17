package iso20022

// Net position of a segregated holding of a single security within the overall position held in the securities account, eg, sub-balance per status.
type HoldingBalance3 struct {

	// Total quantity of financial instrument for the referenced holding.
	Balance *UnitOrFaceAmountChoice `xml:"Bal,omitempty"`

	// Reason a security is not available or additional information about the financial instrument for which the balance is given, for example, unregistered, registered in nominee name.
	BalanceType *SecuritiesEntryType2Code `xml:"BalTp,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormatChoice `xml:"SfkpgPlc,omitempty"`

	// Date of the entitlement.
	Date *ISODate `xml:"Dt,omitempty"`
}

func (h *HoldingBalance3) AddBalance() *UnitOrFaceAmountChoice {
	h.Balance = new(UnitOrFaceAmountChoice)
	return h.Balance
}

func (h *HoldingBalance3) SetBalanceType(value string) {
	h.BalanceType = (*SecuritiesEntryType2Code)(&value)
}

func (h *HoldingBalance3) AddSafekeepingPlace() *SafekeepingPlaceFormatChoice {
	h.SafekeepingPlace = new(SafekeepingPlaceFormatChoice)
	return h.SafekeepingPlace
}

func (h *HoldingBalance3) SetDate(value string) {
	h.Date = (*ISODate)(&value)
}

// Net position of a segregated holding of a single security within the overall position held in the securities account, eg, sub-balance per status.
type HoldingBalance4 struct {

	// Total quantity of financial instrument for the referenced holding.
	Balance *UnitOrFaceAmountOrCodeChoice `xml:"Bal"`

	// Reason a security is not available or additional information about the financial instrument for which the balance is given, for example, unregistered, registered in nominee name.
	BalanceType *SecuritiesEntryType2Code `xml:"BalTp,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormatChoice `xml:"SfkpgPlc,omitempty"`
}

func (h *HoldingBalance4) AddBalance() *UnitOrFaceAmountOrCodeChoice {
	h.Balance = new(UnitOrFaceAmountOrCodeChoice)
	return h.Balance
}

func (h *HoldingBalance4) SetBalanceType(value string) {
	h.BalanceType = (*SecuritiesEntryType2Code)(&value)
}

func (h *HoldingBalance4) AddSafekeepingPlace() *SafekeepingPlaceFormatChoice {
	h.SafekeepingPlace = new(SafekeepingPlaceFormatChoice)
	return h.SafekeepingPlace
}

// Net position of a segregated holding of a single security within the overall position held in the securities account, eg, sub-balance per status.
type HoldingBalance5 struct {

	// Total quantity of financial instrument for the referenced holding.
	Balance *UnitOrFaceAmountOrCode1Choice `xml:"Bal"`

	// Reason a security is not available or additional information about the financial instrument for which the balance is given, for example, unregistered, registered in nominee name.
	BalanceType *SecuritiesEntryType2Code `xml:"BalTp,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormatChoice `xml:"SfkpgPlc,omitempty"`
}

func (h *HoldingBalance5) AddBalance() *UnitOrFaceAmountOrCode1Choice {
	h.Balance = new(UnitOrFaceAmountOrCode1Choice)
	return h.Balance
}

func (h *HoldingBalance5) SetBalanceType(value string) {
	h.BalanceType = (*SecuritiesEntryType2Code)(&value)
}

func (h *HoldingBalance5) AddSafekeepingPlace() *SafekeepingPlaceFormatChoice {
	h.SafekeepingPlace = new(SafekeepingPlaceFormatChoice)
	return h.SafekeepingPlace
}

// Net position of a segregated holding of a single security within the overall position held in the securities account, eg, sub-balance per status.
type HoldingBalance6 struct {

	// Total quantity of financial instrument for the referenced holding.
	Balance *UnitOrFaceAmount1Choice `xml:"Bal,omitempty"`

	// Reason a security is not available or additional information about the financial instrument for which the balance is given, for example, unregistered, registered in nominee name.
	BalanceType *SecuritiesEntryType2Code `xml:"BalTp,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormatChoice `xml:"SfkpgPlc,omitempty"`

	// Date of the entitlement.
	Date *ISODate `xml:"Dt,omitempty"`
}

func (h *HoldingBalance6) AddBalance() *UnitOrFaceAmount1Choice {
	h.Balance = new(UnitOrFaceAmount1Choice)
	return h.Balance
}

func (h *HoldingBalance6) SetBalanceType(value string) {
	h.BalanceType = (*SecuritiesEntryType2Code)(&value)
}

func (h *HoldingBalance6) AddSafekeepingPlace() *SafekeepingPlaceFormatChoice {
	h.SafekeepingPlace = new(SafekeepingPlaceFormatChoice)
	return h.SafekeepingPlace
}

func (h *HoldingBalance6) SetDate(value string) {
	h.Date = (*ISODate)(&value)
}

// Net position of a segregated holding of a single security within the overall position held in the securities account, for example, sub-balance per status.
type HoldingBalance7 struct {

	// Total quantity of financial instrument for the referenced holding.
	Balance *UnitOrFaceAmount1Choice `xml:"Bal,omitempty"`

	// Reason a security is not available or additional information about the financial instrument for which the balance is given, for example, unregistered, registered in nominee name.
	BalanceType *SecuritiesEntryType2Code `xml:"BalTp,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`

	// Date of the entitlement.
	Date *ISODate `xml:"Dt,omitempty"`
}

func (h *HoldingBalance7) AddBalance() *UnitOrFaceAmount1Choice {
	h.Balance = new(UnitOrFaceAmount1Choice)
	return h.Balance
}

func (h *HoldingBalance7) SetBalanceType(value string) {
	h.BalanceType = (*SecuritiesEntryType2Code)(&value)
}

func (h *HoldingBalance7) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	h.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return h.SafekeepingPlace
}

func (h *HoldingBalance7) SetDate(value string) {
	h.Date = (*ISODate)(&value)
}

// Net position of a segregated holding of a single security within the overall position held in the securities account,for example, sub-balance per status.
type HoldingBalance8 struct {

	// Total quantity of financial instrument for the referenced holding.
	Balance *UnitOrFaceAmountOrCode1Choice `xml:"Bal"`

	// Reason a security is not available or additional information about the financial instrument for which the balance is given, for example, unregistered, registered in nominee name.
	BalanceType *SecuritiesEntryType2Code `xml:"BalTp,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`
}

func (h *HoldingBalance8) AddBalance() *UnitOrFaceAmountOrCode1Choice {
	h.Balance = new(UnitOrFaceAmountOrCode1Choice)
	return h.Balance
}

func (h *HoldingBalance8) SetBalanceType(value string) {
	h.BalanceType = (*SecuritiesEntryType2Code)(&value)
}

func (h *HoldingBalance8) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	h.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return h.SafekeepingPlace
}
