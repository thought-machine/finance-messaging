package iso20022

// Detailed account holdings information report for a corporate action event.
type CorporateActionEventAndBalance1 struct {

	// Provides general information related to a corporate action event.
	GeneralInformation *EventInformation1 `xml:"GnlInf"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *UnderlyingSecurity1 `xml:"UndrlygScty"`

	// Provides information about the balance related to a corporate action.
	Balance *CorporateActionBalanceDetails4 `xml:"Bal,omitempty"`

	// Provides additional information related to the event and the balance of the corporate action.
	Extension []*Extension2 `xml:"Xtnsn,omitempty"`
}

func (c *CorporateActionEventAndBalance1) AddGeneralInformation() *EventInformation1 {
	c.GeneralInformation = new(EventInformation1)
	return c.GeneralInformation
}

func (c *CorporateActionEventAndBalance1) AddUnderlyingSecurity() *UnderlyingSecurity1 {
	c.UnderlyingSecurity = new(UnderlyingSecurity1)
	return c.UnderlyingSecurity
}

func (c *CorporateActionEventAndBalance1) AddBalance() *CorporateActionBalanceDetails4 {
	c.Balance = new(CorporateActionBalanceDetails4)
	return c.Balance
}

func (c *CorporateActionEventAndBalance1) AddExtension() *Extension2 {
	newValue := new(Extension2)
	c.Extension = append(c.Extension, newValue)
	return newValue
}

// Detailed account holdings information report for a corporate action event.
type CorporateActionEventAndBalance10 struct {

	// Provides general information related to a corporate action event.
	GeneralInformation *EventInformation8 `xml:"GnlInf"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *SecurityIdentification20 `xml:"UndrlygScty"`

	// Provides information about the balance related to a corporate action.
	Balance *CorporateActionBalanceDetails33 `xml:"Bal,omitempty"`

	// Provides additional information related to the event and the balance of the corporate action.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionEventAndBalance10) AddGeneralInformation() *EventInformation8 {
	c.GeneralInformation = new(EventInformation8)
	return c.GeneralInformation
}

func (c *CorporateActionEventAndBalance10) AddUnderlyingSecurity() *SecurityIdentification20 {
	c.UnderlyingSecurity = new(SecurityIdentification20)
	return c.UnderlyingSecurity
}

func (c *CorporateActionEventAndBalance10) AddBalance() *CorporateActionBalanceDetails33 {
	c.Balance = new(CorporateActionBalanceDetails33)
	return c.Balance
}

func (c *CorporateActionEventAndBalance10) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}

// Detailed account holdings information report for a corporate action event.
type CorporateActionEventAndBalance11 struct {

	// Provides general information related to a corporate action event.
	GeneralInformation *EventInformation9 `xml:"GnlInf"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *SecurityIdentification19 `xml:"UndrlygScty"`

	// Provides information about the balance related to a corporate action.
	Balance *CorporateActionBalanceDetails30 `xml:"Bal,omitempty"`

	// Provides additional information related to the event and the balance of the corporate action.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionEventAndBalance11) AddGeneralInformation() *EventInformation9 {
	c.GeneralInformation = new(EventInformation9)
	return c.GeneralInformation
}

func (c *CorporateActionEventAndBalance11) AddUnderlyingSecurity() *SecurityIdentification19 {
	c.UnderlyingSecurity = new(SecurityIdentification19)
	return c.UnderlyingSecurity
}

func (c *CorporateActionEventAndBalance11) AddBalance() *CorporateActionBalanceDetails30 {
	c.Balance = new(CorporateActionBalanceDetails30)
	return c.Balance
}

func (c *CorporateActionEventAndBalance11) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}

// Detailed account holdings information report for a corporate action event.
type CorporateActionEventAndBalance12 struct {

	// Provides general information related to a corporate action event.
	GeneralInformation *EventInformation10 `xml:"GnlInf"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *SecurityIdentification20 `xml:"UndrlygScty"`

	// Provides information about the balance related to a corporate action.
	Balance *CorporateActionBalanceDetails33 `xml:"Bal,omitempty"`

	// Provides additional information related to the event and the balance of the corporate action.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionEventAndBalance12) AddGeneralInformation() *EventInformation10 {
	c.GeneralInformation = new(EventInformation10)
	return c.GeneralInformation
}

func (c *CorporateActionEventAndBalance12) AddUnderlyingSecurity() *SecurityIdentification20 {
	c.UnderlyingSecurity = new(SecurityIdentification20)
	return c.UnderlyingSecurity
}

func (c *CorporateActionEventAndBalance12) AddBalance() *CorporateActionBalanceDetails33 {
	c.Balance = new(CorporateActionBalanceDetails33)
	return c.Balance
}

func (c *CorporateActionEventAndBalance12) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}

// Detailed account holdings information report for a corporate action event.
type CorporateActionEventAndBalance3 struct {

	// Provides general information related to a corporate action event.
	GeneralInformation *EventInformation1 `xml:"GnlInf"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *UnderlyingSecurity3 `xml:"UndrlygScty"`

	// Provides information about the balance related to a corporate action.
	Balance *CorporateActionBalanceDetails4 `xml:"Bal,omitempty"`

	// Provides additional information related to the event and the balance of the corporate action.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionEventAndBalance3) AddGeneralInformation() *EventInformation1 {
	c.GeneralInformation = new(EventInformation1)
	return c.GeneralInformation
}

func (c *CorporateActionEventAndBalance3) AddUnderlyingSecurity() *UnderlyingSecurity3 {
	c.UnderlyingSecurity = new(UnderlyingSecurity3)
	return c.UnderlyingSecurity
}

func (c *CorporateActionEventAndBalance3) AddBalance() *CorporateActionBalanceDetails4 {
	c.Balance = new(CorporateActionBalanceDetails4)
	return c.Balance
}

func (c *CorporateActionEventAndBalance3) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}

// Detailed account holdings information report for a corporate action event.
type CorporateActionEventAndBalance5 struct {

	// Provides general information related to a corporate action event.
	GeneralInformation *EventInformation3 `xml:"GnlInf"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *SecurityIdentification14 `xml:"UndrlygScty"`

	// Provides information about the balance related to a corporate action.
	Balance *CorporateActionBalanceDetails9 `xml:"Bal,omitempty"`

	// Provides additional information related to the event and the balance of the corporate action.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionEventAndBalance5) AddGeneralInformation() *EventInformation3 {
	c.GeneralInformation = new(EventInformation3)
	return c.GeneralInformation
}

func (c *CorporateActionEventAndBalance5) AddUnderlyingSecurity() *SecurityIdentification14 {
	c.UnderlyingSecurity = new(SecurityIdentification14)
	return c.UnderlyingSecurity
}

func (c *CorporateActionEventAndBalance5) AddBalance() *CorporateActionBalanceDetails9 {
	c.Balance = new(CorporateActionBalanceDetails9)
	return c.Balance
}

func (c *CorporateActionEventAndBalance5) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}

// Detailed account holdings information report for a corporate action event.
type CorporateActionEventAndBalance7 struct {

	// Provides general information related to a corporate action event.
	GeneralInformation *EventInformation5 `xml:"GnlInf"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *SecurityIdentification14 `xml:"UndrlygScty"`

	// Provides information about the balance related to a corporate action.
	Balance *CorporateActionBalanceDetails9 `xml:"Bal,omitempty"`

	// Provides additional information related to the event and the balance of the corporate action.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionEventAndBalance7) AddGeneralInformation() *EventInformation5 {
	c.GeneralInformation = new(EventInformation5)
	return c.GeneralInformation
}

func (c *CorporateActionEventAndBalance7) AddUnderlyingSecurity() *SecurityIdentification14 {
	c.UnderlyingSecurity = new(SecurityIdentification14)
	return c.UnderlyingSecurity
}

func (c *CorporateActionEventAndBalance7) AddBalance() *CorporateActionBalanceDetails9 {
	c.Balance = new(CorporateActionBalanceDetails9)
	return c.Balance
}

func (c *CorporateActionEventAndBalance7) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}

// Detailed account holdings information report for a corporate action event.
type CorporateActionEventAndBalance9 struct {

	// Provides general information related to a corporate action event.
	GeneralInformation *EventInformation7 `xml:"GnlInf"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *SecurityIdentification19 `xml:"UndrlygScty"`

	// Provides information about the balance related to a corporate action.
	Balance *CorporateActionBalanceDetails30 `xml:"Bal,omitempty"`

	// Provides additional information related to the event and the balance of the corporate action.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionEventAndBalance9) AddGeneralInformation() *EventInformation7 {
	c.GeneralInformation = new(EventInformation7)
	return c.GeneralInformation
}

func (c *CorporateActionEventAndBalance9) AddUnderlyingSecurity() *SecurityIdentification19 {
	c.UnderlyingSecurity = new(SecurityIdentification19)
	return c.UnderlyingSecurity
}

func (c *CorporateActionEventAndBalance9) AddBalance() *CorporateActionBalanceDetails30 {
	c.Balance = new(CorporateActionBalanceDetails30)
	return c.Balance
}

func (c *CorporateActionEventAndBalance9) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
