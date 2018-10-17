package iso20022

// Net position of a segregated holding of a single security within the overall position held in the securities account, eg, sub-balance per status.
type AdditionalBalanceInformation struct {

	// Quantity of securities in the sub-balance.
	Quantity *SubBalanceQuantity1Choice `xml:"Qty"`

	// Reason a security is not available or additional information about the financial instrument for which the balance is given, for example, unregistered, registered in nominee name.
	SubBalanceType *SecuritiesBalanceType2Code `xml:"SubBalTp"`
}

func (a *AdditionalBalanceInformation) AddQuantity() *SubBalanceQuantity1Choice {
	a.Quantity = new(SubBalanceQuantity1Choice)
	return a.Quantity
}

func (a *AdditionalBalanceInformation) SetSubBalanceType(value string) {
	a.SubBalanceType = (*SecuritiesBalanceType2Code)(&value)
}

// Subbalances providing additional information on a specific position but that is not to be accounted for in the building of the aggregate balance, for example, registered.
type AdditionalBalanceInformation11 struct {

	// Reason for the sub-balance.
	SubBalanceType *SubBalanceType6Choice `xml:"SubBalTp"`

	// Quantity of securities in the sub-balance.
	Quantity *SubBalanceQuantity3Choice `xml:"Qty"`

	// Provides additional subbalance information.
	SubBalanceAdditionalDetails *Max140Text `xml:"SubBalAddtlDtls,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown23 `xml:"QtyBrkdwn,omitempty"`
}

func (a *AdditionalBalanceInformation11) AddSubBalanceType() *SubBalanceType6Choice {
	a.SubBalanceType = new(SubBalanceType6Choice)
	return a.SubBalanceType
}

func (a *AdditionalBalanceInformation11) AddQuantity() *SubBalanceQuantity3Choice {
	a.Quantity = new(SubBalanceQuantity3Choice)
	return a.Quantity
}

func (a *AdditionalBalanceInformation11) SetSubBalanceAdditionalDetails(value string) {
	a.SubBalanceAdditionalDetails = (*Max140Text)(&value)
}

func (a *AdditionalBalanceInformation11) AddQuantityBreakdown() *QuantityBreakdown23 {
	newValue := new(QuantityBreakdown23)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

// Sub-balances providing additional information on a specific position but that is not to be accounted for in the building of the aggregate balance, for example, registered.
type AdditionalBalanceInformation14 struct {

	// Reason for the sub-balance.
	SubBalanceType *SubBalanceType12Choice `xml:"SubBalTp"`

	// Quantity of securities in the sub-balance.
	Quantity *SubBalanceQuantity6Choice `xml:"Qty"`

	// Provides additional subbalance information.
	SubBalanceAdditionalDetails *Max140Text `xml:"SubBalAddtlDtls,omitempty"`
}

func (a *AdditionalBalanceInformation14) AddSubBalanceType() *SubBalanceType12Choice {
	a.SubBalanceType = new(SubBalanceType12Choice)
	return a.SubBalanceType
}

func (a *AdditionalBalanceInformation14) AddQuantity() *SubBalanceQuantity6Choice {
	a.Quantity = new(SubBalanceQuantity6Choice)
	return a.Quantity
}

func (a *AdditionalBalanceInformation14) SetSubBalanceAdditionalDetails(value string) {
	a.SubBalanceAdditionalDetails = (*Max140Text)(&value)
}

// Subbalances providing additional information on a specific position but that is not to be accounted for in the building of the aggregate balance, for example, registered.
type AdditionalBalanceInformation15 struct {

	// Reason for the sub-balance.
	SubBalanceType *SubBalanceType12Choice `xml:"SubBalTp"`

	// Quantity of securities in the sub-balance.
	Quantity *Balance9 `xml:"Qty"`

	// Provides additional sub-balance information.
	SubBalanceAdditionalDetails *Max140Text `xml:"SubBalAddtlDtls,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown27 `xml:"QtyBrkdwn,omitempty"`
}

func (a *AdditionalBalanceInformation15) AddSubBalanceType() *SubBalanceType12Choice {
	a.SubBalanceType = new(SubBalanceType12Choice)
	return a.SubBalanceType
}

func (a *AdditionalBalanceInformation15) AddQuantity() *Balance9 {
	a.Quantity = new(Balance9)
	return a.Quantity
}

func (a *AdditionalBalanceInformation15) SetSubBalanceAdditionalDetails(value string) {
	a.SubBalanceAdditionalDetails = (*Max140Text)(&value)
}

func (a *AdditionalBalanceInformation15) AddQuantityBreakdown() *QuantityBreakdown27 {
	newValue := new(QuantityBreakdown27)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

// Sub-balances providing additional information on a specific position but that is not to be accounted for in the building of the aggregate balance, for example, registered.
type AdditionalBalanceInformation16 struct {

	// Reason for the sub-balance.
	SubBalanceType *SubBalanceType14Choice `xml:"SubBalTp"`

	// Quantity of securities in the sub-balance.
	Quantity *SubBalanceQuantity7Choice `xml:"Qty"`

	// Provides additional subbalance information.
	SubBalanceAdditionalDetails *RestrictedFINXMax140Text `xml:"SubBalAddtlDtls,omitempty"`
}

func (a *AdditionalBalanceInformation16) AddSubBalanceType() *SubBalanceType14Choice {
	a.SubBalanceType = new(SubBalanceType14Choice)
	return a.SubBalanceType
}

func (a *AdditionalBalanceInformation16) AddQuantity() *SubBalanceQuantity7Choice {
	a.Quantity = new(SubBalanceQuantity7Choice)
	return a.Quantity
}

func (a *AdditionalBalanceInformation16) SetSubBalanceAdditionalDetails(value string) {
	a.SubBalanceAdditionalDetails = (*RestrictedFINXMax140Text)(&value)
}

// Subbalances providing additional information on a specific position but that is not to be accounted for in the building of the aggregate balance, for example, registered.
type AdditionalBalanceInformation17 struct {

	// Reason for the sub-balance.
	SubBalanceType *SubBalanceType14Choice `xml:"SubBalTp"`

	// Quantity of securities in the sub-balance.
	Quantity *Balance13 `xml:"Qty"`

	// Provides additional sub-balance information.
	SubBalanceAdditionalDetails *RestrictedFINXMax140Text `xml:"SubBalAddtlDtls,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown40 `xml:"QtyBrkdwn,omitempty"`
}

func (a *AdditionalBalanceInformation17) AddSubBalanceType() *SubBalanceType14Choice {
	a.SubBalanceType = new(SubBalanceType14Choice)
	return a.SubBalanceType
}

func (a *AdditionalBalanceInformation17) AddQuantity() *Balance13 {
	a.Quantity = new(Balance13)
	return a.Quantity
}

func (a *AdditionalBalanceInformation17) SetSubBalanceAdditionalDetails(value string) {
	a.SubBalanceAdditionalDetails = (*RestrictedFINXMax140Text)(&value)
}

func (a *AdditionalBalanceInformation17) AddQuantityBreakdown() *QuantityBreakdown40 {
	newValue := new(QuantityBreakdown40)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

// Net position of a segregated holding of a single security within the overall position held in the securities account, eg, sub-balance per status.
type AdditionalBalanceInformation2 struct {

	// Quantity of securities in the sub-balance.
	Quantity *SubBalanceQuantity1Choice `xml:"Qty"`

	// Reason a security is not available or additional information about the financial instrument for which the balance is given, for example, unregistered, registered in nominee name.
	SubBalanceType *SecuritiesBalanceType2Code `xml:"SubBalTp"`

	// Reason a security is not available or additional information about the financial instrument for which the balance is given, for example, unregistered, registered in nominee name.
	ExtendedSubBalanceType *Extended350Code `xml:"XtndedSubBalTp"`
}

func (a *AdditionalBalanceInformation2) AddQuantity() *SubBalanceQuantity1Choice {
	a.Quantity = new(SubBalanceQuantity1Choice)
	return a.Quantity
}

func (a *AdditionalBalanceInformation2) SetSubBalanceType(value string) {
	a.SubBalanceType = (*SecuritiesBalanceType2Code)(&value)
}

func (a *AdditionalBalanceInformation2) SetExtendedSubBalanceType(value string) {
	a.ExtendedSubBalanceType = (*Extended350Code)(&value)
}

// Subbalances providing additional information on a specific position but that is not to be accounted for in the building of the aggregate balance, for example, registered.
type AdditionalBalanceInformation5 struct {

	// Reason for the sub-balance.
	SubBalanceType *SubBalanceType6Choice `xml:"SubBalTp"`

	// Quantity of securities in the sub-balance.
	Quantity *SubBalanceQuantity3Choice `xml:"Qty"`

	// Provides additional subbalance information.
	SubBalanceAdditionalDetails *Max140Text `xml:"SubBalAddtlDtls,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown3 `xml:"QtyBrkdwn,omitempty"`
}

func (a *AdditionalBalanceInformation5) AddSubBalanceType() *SubBalanceType6Choice {
	a.SubBalanceType = new(SubBalanceType6Choice)
	return a.SubBalanceType
}

func (a *AdditionalBalanceInformation5) AddQuantity() *SubBalanceQuantity3Choice {
	a.Quantity = new(SubBalanceQuantity3Choice)
	return a.Quantity
}

func (a *AdditionalBalanceInformation5) SetSubBalanceAdditionalDetails(value string) {
	a.SubBalanceAdditionalDetails = (*Max140Text)(&value)
}

func (a *AdditionalBalanceInformation5) AddQuantityBreakdown() *QuantityBreakdown3 {
	newValue := new(QuantityBreakdown3)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

// Subbalances providing additional information on a specific position but that is not to be accounted for in the building of the aggregate balance, for example, registered.
type AdditionalBalanceInformation6 struct {

	// Reason for the sub-balance.
	SubBalanceType *SubBalanceType6Choice `xml:"SubBalTp"`

	// Quantity of securities in the sub-balance.
	Quantity *SubBalanceQuantity3Choice `xml:"Qty"`

	// Provides additional subbalance information.
	SubBalanceAdditionalDetails *Max140Text `xml:"SubBalAddtlDtls,omitempty"`
}

func (a *AdditionalBalanceInformation6) AddSubBalanceType() *SubBalanceType6Choice {
	a.SubBalanceType = new(SubBalanceType6Choice)
	return a.SubBalanceType
}

func (a *AdditionalBalanceInformation6) AddQuantity() *SubBalanceQuantity3Choice {
	a.Quantity = new(SubBalanceQuantity3Choice)
	return a.Quantity
}

func (a *AdditionalBalanceInformation6) SetSubBalanceAdditionalDetails(value string) {
	a.SubBalanceAdditionalDetails = (*Max140Text)(&value)
}

// Subbalances providing additional information on a specific position but that is not to be accounted for in the building of the aggregate balance, for example, registered.
type AdditionalBalanceInformation9 struct {

	// Reason for the sub-balance.
	SubBalanceType *SubBalanceType6Choice `xml:"SubBalTp"`

	// Quantity of securities in the sub-balance.
	Quantity *SubBalanceQuantity3Choice `xml:"Qty"`

	// Provides additional subbalance information.
	SubBalanceAdditionalDetails *Max140Text `xml:"SubBalAddtlDtls,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown13 `xml:"QtyBrkdwn,omitempty"`
}

func (a *AdditionalBalanceInformation9) AddSubBalanceType() *SubBalanceType6Choice {
	a.SubBalanceType = new(SubBalanceType6Choice)
	return a.SubBalanceType
}

func (a *AdditionalBalanceInformation9) AddQuantity() *SubBalanceQuantity3Choice {
	a.Quantity = new(SubBalanceQuantity3Choice)
	return a.Quantity
}

func (a *AdditionalBalanceInformation9) SetSubBalanceAdditionalDetails(value string) {
	a.SubBalanceAdditionalDetails = (*Max140Text)(&value)
}

func (a *AdditionalBalanceInformation9) AddQuantityBreakdown() *QuantityBreakdown13 {
	newValue := new(QuantityBreakdown13)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}
