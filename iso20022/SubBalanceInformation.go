package iso20022

// Net position of a segregated holding of a single security within the overall position held in the securities account, eg, sub-balance per status.
type SubBalanceInformation1 struct {

	// Quantity of securities in the sub-balance.
	Quantity *SubBalanceQuantity1Choice `xml:"Qty"`

	// Reason for the sub-balance.
	SubBalanceType *SecuritiesBalanceType1Choice `xml:"SubBalTp"`

	// Net position of a segregated holding of a single security within the overall position held in a securities account, eg, sub-balance per status.
	AdditionalBalanceBreakdownDetails []*AdditionalBalanceInformation `xml:"AddtlBalBrkdwnDtls,omitempty"`
}

func (s *SubBalanceInformation1) AddQuantity() *SubBalanceQuantity1Choice {
	s.Quantity = new(SubBalanceQuantity1Choice)
	return s.Quantity
}

func (s *SubBalanceInformation1) AddSubBalanceType() *SecuritiesBalanceType1Choice {
	s.SubBalanceType = new(SecuritiesBalanceType1Choice)
	return s.SubBalanceType
}

func (s *SubBalanceInformation1) AddAdditionalBalanceBreakdownDetails() *AdditionalBalanceInformation {
	newValue := new(AdditionalBalanceInformation)
	s.AdditionalBalanceBreakdownDetails = append(s.AdditionalBalanceBreakdownDetails, newValue)
	return newValue
}

// Net position of a segregated holding of a single security within the overall position held in the securities account, eg, sub-balance per status.
type SubBalanceInformation11 struct {

	// Reason for the sub-balance.
	SubBalanceType *SubBalanceType5Choice `xml:"SubBalTp"`

	// Quantity of securities in the sub-balance.
	Quantity *SubBalanceQuantity3Choice `xml:"Qty"`

	// Provides additional subbalance information.
	SubBalanceAdditionalDetails *Max140Text `xml:"SubBalAddtlDtls,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown23 `xml:"QtyBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdownDetails []*AdditionalBalanceInformation11 `xml:"AddtlBalBrkdwnDtls,omitempty"`
}

func (s *SubBalanceInformation11) AddSubBalanceType() *SubBalanceType5Choice {
	s.SubBalanceType = new(SubBalanceType5Choice)
	return s.SubBalanceType
}

func (s *SubBalanceInformation11) AddQuantity() *SubBalanceQuantity3Choice {
	s.Quantity = new(SubBalanceQuantity3Choice)
	return s.Quantity
}

func (s *SubBalanceInformation11) SetSubBalanceAdditionalDetails(value string) {
	s.SubBalanceAdditionalDetails = (*Max140Text)(&value)
}

func (s *SubBalanceInformation11) AddQuantityBreakdown() *QuantityBreakdown23 {
	newValue := new(QuantityBreakdown23)
	s.QuantityBreakdown = append(s.QuantityBreakdown, newValue)
	return newValue
}

func (s *SubBalanceInformation11) AddAdditionalBalanceBreakdownDetails() *AdditionalBalanceInformation11 {
	newValue := new(AdditionalBalanceInformation11)
	s.AdditionalBalanceBreakdownDetails = append(s.AdditionalBalanceBreakdownDetails, newValue)
	return newValue
}

// Net position of a segregated holding of a single security within the overall position held in the securities account, for example sub-balance per status.
type SubBalanceInformation14 struct {

	// Reason for the sub-balance.
	SubBalanceType *SubBalanceType11Choice `xml:"SubBalTp"`

	// Quantity of securities in the sub-balance.
	Quantity *SubBalanceQuantity6Choice `xml:"Qty"`

	// Provides additional subbalance information.
	SubBalanceAdditionalDetails *Max140Text `xml:"SubBalAddtlDtls,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdownDetails []*AdditionalBalanceInformation14 `xml:"AddtlBalBrkdwnDtls,omitempty"`
}

func (s *SubBalanceInformation14) AddSubBalanceType() *SubBalanceType11Choice {
	s.SubBalanceType = new(SubBalanceType11Choice)
	return s.SubBalanceType
}

func (s *SubBalanceInformation14) AddQuantity() *SubBalanceQuantity6Choice {
	s.Quantity = new(SubBalanceQuantity6Choice)
	return s.Quantity
}

func (s *SubBalanceInformation14) SetSubBalanceAdditionalDetails(value string) {
	s.SubBalanceAdditionalDetails = (*Max140Text)(&value)
}

func (s *SubBalanceInformation14) AddAdditionalBalanceBreakdownDetails() *AdditionalBalanceInformation14 {
	newValue := new(AdditionalBalanceInformation14)
	s.AdditionalBalanceBreakdownDetails = append(s.AdditionalBalanceBreakdownDetails, newValue)
	return newValue
}

// Net position of a segregated holding of a single security within the overall position held in the securities account, eg, sub-balance per status.
type SubBalanceInformation15 struct {

	// Reason for the sub-balance.
	SubBalanceType *SubBalanceType11Choice `xml:"SubBalTp"`

	// Quantity of securities in the sub-balance.
	Quantity *Balance9 `xml:"Qty"`

	// Provides additional subbalance information.
	SubBalanceAdditionalDetails *Max140Text `xml:"SubBalAddtlDtls,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown27 `xml:"QtyBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdownDetails []*AdditionalBalanceInformation15 `xml:"AddtlBalBrkdwnDtls,omitempty"`
}

func (s *SubBalanceInformation15) AddSubBalanceType() *SubBalanceType11Choice {
	s.SubBalanceType = new(SubBalanceType11Choice)
	return s.SubBalanceType
}

func (s *SubBalanceInformation15) AddQuantity() *Balance9 {
	s.Quantity = new(Balance9)
	return s.Quantity
}

func (s *SubBalanceInformation15) SetSubBalanceAdditionalDetails(value string) {
	s.SubBalanceAdditionalDetails = (*Max140Text)(&value)
}

func (s *SubBalanceInformation15) AddQuantityBreakdown() *QuantityBreakdown27 {
	newValue := new(QuantityBreakdown27)
	s.QuantityBreakdown = append(s.QuantityBreakdown, newValue)
	return newValue
}

func (s *SubBalanceInformation15) AddAdditionalBalanceBreakdownDetails() *AdditionalBalanceInformation15 {
	newValue := new(AdditionalBalanceInformation15)
	s.AdditionalBalanceBreakdownDetails = append(s.AdditionalBalanceBreakdownDetails, newValue)
	return newValue
}

// Net position of a segregated holding of a single security within the overall position held in the securities account, for example sub-balance per status.
type SubBalanceInformation16 struct {

	// Reason for the sub-balance.
	SubBalanceType *SubBalanceType13Choice `xml:"SubBalTp"`

	// Quantity of securities in the sub-balance.
	Quantity *SubBalanceQuantity7Choice `xml:"Qty"`

	// Provides additional subbalance information.
	SubBalanceAdditionalDetails *RestrictedFINXMax140Text `xml:"SubBalAddtlDtls,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdownDetails []*AdditionalBalanceInformation16 `xml:"AddtlBalBrkdwnDtls,omitempty"`
}

func (s *SubBalanceInformation16) AddSubBalanceType() *SubBalanceType13Choice {
	s.SubBalanceType = new(SubBalanceType13Choice)
	return s.SubBalanceType
}

func (s *SubBalanceInformation16) AddQuantity() *SubBalanceQuantity7Choice {
	s.Quantity = new(SubBalanceQuantity7Choice)
	return s.Quantity
}

func (s *SubBalanceInformation16) SetSubBalanceAdditionalDetails(value string) {
	s.SubBalanceAdditionalDetails = (*RestrictedFINXMax140Text)(&value)
}

func (s *SubBalanceInformation16) AddAdditionalBalanceBreakdownDetails() *AdditionalBalanceInformation16 {
	newValue := new(AdditionalBalanceInformation16)
	s.AdditionalBalanceBreakdownDetails = append(s.AdditionalBalanceBreakdownDetails, newValue)
	return newValue
}

// Net position of a segregated holding of a single security within the overall position held in the securities account, eg, sub-balance per status.
type SubBalanceInformation17 struct {

	// Reason for the sub-balance.
	SubBalanceType *SubBalanceType13Choice `xml:"SubBalTp"`

	// Quantity of securities in the sub-balance.
	Quantity *Balance13 `xml:"Qty"`

	// Provides additional subbalance information.
	SubBalanceAdditionalDetails *RestrictedFINXMax140Text `xml:"SubBalAddtlDtls,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown40 `xml:"QtyBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdownDetails []*AdditionalBalanceInformation17 `xml:"AddtlBalBrkdwnDtls,omitempty"`
}

func (s *SubBalanceInformation17) AddSubBalanceType() *SubBalanceType13Choice {
	s.SubBalanceType = new(SubBalanceType13Choice)
	return s.SubBalanceType
}

func (s *SubBalanceInformation17) AddQuantity() *Balance13 {
	s.Quantity = new(Balance13)
	return s.Quantity
}

func (s *SubBalanceInformation17) SetSubBalanceAdditionalDetails(value string) {
	s.SubBalanceAdditionalDetails = (*RestrictedFINXMax140Text)(&value)
}

func (s *SubBalanceInformation17) AddQuantityBreakdown() *QuantityBreakdown40 {
	newValue := new(QuantityBreakdown40)
	s.QuantityBreakdown = append(s.QuantityBreakdown, newValue)
	return newValue
}

func (s *SubBalanceInformation17) AddAdditionalBalanceBreakdownDetails() *AdditionalBalanceInformation17 {
	newValue := new(AdditionalBalanceInformation17)
	s.AdditionalBalanceBreakdownDetails = append(s.AdditionalBalanceBreakdownDetails, newValue)
	return newValue
}

// Net position of a segregated holding of a single security within the overall position held in the securities account, eg, sub-balance per status.
type SubBalanceInformation2 struct {

	// Quantity of securities in the sub-balance.
	Quantity *SubBalanceQuantity1Choice `xml:"Qty"`

	// Reason for the sub-balance.
	SubBalanceType *SecuritiesBalanceType1Code `xml:"SubBalTp"`

	// Reason a security is not available or additional information about the financial instrument for which the balance is given, for example, unregistered, registered in nominee name.
	ExtendedSubBalanceType *Extended350Code `xml:"XtndedSubBalTp"`

	// Net position of a segregated holding of a single security within the overall position held in a securities account, eg, sub-balance per status.
	AdditionalBalanceBreakdownDetails []*AdditionalBalanceInformation2 `xml:"AddtlBalBrkdwnDtls,omitempty"`
}

func (s *SubBalanceInformation2) AddQuantity() *SubBalanceQuantity1Choice {
	s.Quantity = new(SubBalanceQuantity1Choice)
	return s.Quantity
}

func (s *SubBalanceInformation2) SetSubBalanceType(value string) {
	s.SubBalanceType = (*SecuritiesBalanceType1Code)(&value)
}

func (s *SubBalanceInformation2) SetExtendedSubBalanceType(value string) {
	s.ExtendedSubBalanceType = (*Extended350Code)(&value)
}

func (s *SubBalanceInformation2) AddAdditionalBalanceBreakdownDetails() *AdditionalBalanceInformation2 {
	newValue := new(AdditionalBalanceInformation2)
	s.AdditionalBalanceBreakdownDetails = append(s.AdditionalBalanceBreakdownDetails, newValue)
	return newValue
}

// Net position of a segregated holding of a single security within the overall position held in the securities account, eg, sub-balance per status.
type SubBalanceInformation5 struct {

	// Reason for the sub-balance.
	SubBalanceType *SubBalanceType5Choice `xml:"SubBalTp"`

	// Quantity of securities in the sub-balance.
	Quantity *SubBalanceQuantity3Choice `xml:"Qty"`

	// Provides additional subbalance information.
	SubBalanceAdditionalDetails *Max140Text `xml:"SubBalAddtlDtls,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown3 `xml:"QtyBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdownDetails []*AdditionalBalanceInformation5 `xml:"AddtlBalBrkdwnDtls,omitempty"`
}

func (s *SubBalanceInformation5) AddSubBalanceType() *SubBalanceType5Choice {
	s.SubBalanceType = new(SubBalanceType5Choice)
	return s.SubBalanceType
}

func (s *SubBalanceInformation5) AddQuantity() *SubBalanceQuantity3Choice {
	s.Quantity = new(SubBalanceQuantity3Choice)
	return s.Quantity
}

func (s *SubBalanceInformation5) SetSubBalanceAdditionalDetails(value string) {
	s.SubBalanceAdditionalDetails = (*Max140Text)(&value)
}

func (s *SubBalanceInformation5) AddQuantityBreakdown() *QuantityBreakdown3 {
	newValue := new(QuantityBreakdown3)
	s.QuantityBreakdown = append(s.QuantityBreakdown, newValue)
	return newValue
}

func (s *SubBalanceInformation5) AddAdditionalBalanceBreakdownDetails() *AdditionalBalanceInformation5 {
	newValue := new(AdditionalBalanceInformation5)
	s.AdditionalBalanceBreakdownDetails = append(s.AdditionalBalanceBreakdownDetails, newValue)
	return newValue
}

// Net position of a segregated holding of a single security within the overall position held in the securities account, eg, sub-balance per status.
type SubBalanceInformation6 struct {

	// Reason for the sub-balance.
	SubBalanceType *SubBalanceType5Choice `xml:"SubBalTp"`

	// Quantity of securities in the sub-balance.
	Quantity *SubBalanceQuantity3Choice `xml:"Qty"`

	// Provides additional subbalance information.
	SubBalanceAdditionalDetails *Max140Text `xml:"SubBalAddtlDtls,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdownDetails []*AdditionalBalanceInformation6 `xml:"AddtlBalBrkdwnDtls,omitempty"`
}

func (s *SubBalanceInformation6) AddSubBalanceType() *SubBalanceType5Choice {
	s.SubBalanceType = new(SubBalanceType5Choice)
	return s.SubBalanceType
}

func (s *SubBalanceInformation6) AddQuantity() *SubBalanceQuantity3Choice {
	s.Quantity = new(SubBalanceQuantity3Choice)
	return s.Quantity
}

func (s *SubBalanceInformation6) SetSubBalanceAdditionalDetails(value string) {
	s.SubBalanceAdditionalDetails = (*Max140Text)(&value)
}

func (s *SubBalanceInformation6) AddAdditionalBalanceBreakdownDetails() *AdditionalBalanceInformation6 {
	newValue := new(AdditionalBalanceInformation6)
	s.AdditionalBalanceBreakdownDetails = append(s.AdditionalBalanceBreakdownDetails, newValue)
	return newValue
}

// Net position of a segregated holding of a single security within the overall position held in the securities account, eg, sub-balance per status.
type SubBalanceInformation9 struct {

	// Reason for the sub-balance.
	SubBalanceType *SubBalanceType5Choice `xml:"SubBalTp"`

	// Quantity of securities in the sub-balance.
	Quantity *SubBalanceQuantity3Choice `xml:"Qty"`

	// Provides additional subbalance information.
	SubBalanceAdditionalDetails *Max140Text `xml:"SubBalAddtlDtls,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown13 `xml:"QtyBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdownDetails []*AdditionalBalanceInformation9 `xml:"AddtlBalBrkdwnDtls,omitempty"`
}

func (s *SubBalanceInformation9) AddSubBalanceType() *SubBalanceType5Choice {
	s.SubBalanceType = new(SubBalanceType5Choice)
	return s.SubBalanceType
}

func (s *SubBalanceInformation9) AddQuantity() *SubBalanceQuantity3Choice {
	s.Quantity = new(SubBalanceQuantity3Choice)
	return s.Quantity
}

func (s *SubBalanceInformation9) SetSubBalanceAdditionalDetails(value string) {
	s.SubBalanceAdditionalDetails = (*Max140Text)(&value)
}

func (s *SubBalanceInformation9) AddQuantityBreakdown() *QuantityBreakdown13 {
	newValue := new(QuantityBreakdown13)
	s.QuantityBreakdown = append(s.QuantityBreakdown, newValue)
	return newValue
}

func (s *SubBalanceInformation9) AddAdditionalBalanceBreakdownDetails() *AdditionalBalanceInformation9 {
	newValue := new(AdditionalBalanceInformation9)
	s.AdditionalBalanceBreakdownDetails = append(s.AdditionalBalanceBreakdownDetails, newValue)
	return newValue
}
