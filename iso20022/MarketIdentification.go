package iso20022

// Context, or geographic environment, in which trading parties may meet in order to negotiate and execute trades among themselves.
type MarketIdentification13 struct {

	// Code allocated to places of trade, ie, stock exchanges, regulated markets, eg, Electronic Trading Platforms (ECN), and unregulated markets, eg, Automated Trading Systems (ATS), as sources of prices and related information, in order to facilitate automated processing.
	Identification *MarketIdentification3Choice `xml:"Id,omitempty"`

	// Nature of a market in which transactions take place.
	Type *MarketType12Choice `xml:"Tp,omitempty"`
}

func (m *MarketIdentification13) AddIdentification() *MarketIdentification3Choice {
	m.Identification = new(MarketIdentification3Choice)
	return m.Identification
}

func (m *MarketIdentification13) AddType() *MarketType12Choice {
	m.Type = new(MarketType12Choice)
	return m.Type
}

// Provides information about market identification and market type.
type MarketIdentification2 struct {

	// Specifies the type of market.
	Type *MarketTypeFormat1Choice `xml:"Tp"`

	// Identifies the market.
	Identification *MarketIdentification1Choice `xml:"Id,omitempty"`
}

func (m *MarketIdentification2) AddType() *MarketTypeFormat1Choice {
	m.Type = new(MarketTypeFormat1Choice)
	return m.Type
}

func (m *MarketIdentification2) AddIdentification() *MarketIdentification1Choice {
	m.Identification = new(MarketIdentification1Choice)
	return m.Identification
}

// Context, or geographic environment, in which trading parties may meet in order to negotiate and execute trades among themselves.
type MarketIdentification20 struct {

	// Code allocated to places of trade, ie, stock exchanges, regulated markets, eg, Electronic Trading Platforms (ECN), and unregulated markets, eg, Automated Trading Systems (ATS), as sources of prices and related information, in order to facilitate automated processing.
	Identification *MarketIdentification1Choice `xml:"Id,omitempty"`

	// Nature of a market in which transactions take place.
	Type *MarketType8Choice `xml:"Tp"`
}

func (m *MarketIdentification20) AddIdentification() *MarketIdentification1Choice {
	m.Identification = new(MarketIdentification1Choice)
	return m.Identification
}

func (m *MarketIdentification20) AddType() *MarketType8Choice {
	m.Type = new(MarketType8Choice)
	return m.Type
}

// Context, or geographic environment, in which trading parties may meet in order to negotiate and execute trades among themselves.
type MarketIdentification4 struct {

	// Code allocated to places of trade, ie, stock exchanges, regulated markets, eg, Electronic Trading Platforms (ECN), and unregulated markets, eg, Automated Trading Systems (ATS), as sources of prices and related information, in order to facilitate automated processing.
	Identification *MarketIdentification1Choice `xml:"Id,omitempty"`

	// Nature of a market in which transactions take place.
	Type *MarketType3Choice `xml:"Tp"`
}

func (m *MarketIdentification4) AddIdentification() *MarketIdentification1Choice {
	m.Identification = new(MarketIdentification1Choice)
	return m.Identification
}

func (m *MarketIdentification4) AddType() *MarketType3Choice {
	m.Type = new(MarketType3Choice)
	return m.Type
}

// Context, or geographic environment, in which trading parties may meet in order to negotiate and execute trades among themselves.
type MarketIdentification5 struct {

	// Code allocated to places of trade, ie, stock exchanges, regulated markets, eg, Electronic Trading Platforms (ECN), and unregulated markets, eg, Automated Trading Systems (ATS), as sources of prices and related information, in order to facilitate automated processing.
	Identification *MarketIdentification1Choice `xml:"Id,omitempty"`

	// Nature of a market in which transactions take place.
	Type *MarketType2Choice `xml:"Tp"`
}

func (m *MarketIdentification5) AddIdentification() *MarketIdentification1Choice {
	m.Identification = new(MarketIdentification1Choice)
	return m.Identification
}

func (m *MarketIdentification5) AddType() *MarketType2Choice {
	m.Type = new(MarketType2Choice)
	return m.Type
}

// Context, or geographic environment, in which trading parties may meet in order to negotiate and execute trades among themselves.
type MarketIdentification6 struct {

	// Code allocated to places of trade, ie, stock exchanges, regulated markets, eg, Electronic Trading Platforms (ECN), and unregulated markets, eg, Automated Trading Systems (ATS), as sources of prices and related information, in order to facilitate automated processing.
	Identification *MarketIdentification1Choice `xml:"Id,omitempty"`

	// Nature of a market in which transactions take place.
	Type *MarketType4Choice `xml:"Tp"`
}

func (m *MarketIdentification6) AddIdentification() *MarketIdentification1Choice {
	m.Identification = new(MarketIdentification1Choice)
	return m.Identification
}

func (m *MarketIdentification6) AddType() *MarketType4Choice {
	m.Type = new(MarketType4Choice)
	return m.Type
}

// Context, or geographic environment, in which trading parties may meet in order to negotiate and execute trades among themselves.
type MarketIdentification77 struct {

	// Code allocated to places of trade, ie, stock exchanges, regulated markets, eg, Electronic Trading Platforms (ECN), and unregulated markets, eg, Automated Trading Systems (ATS), as sources of prices and related information, in order to facilitate automated processing.
	Identification *MarketIdentification3Choice `xml:"Id,omitempty"`

	// Nature of a market in which transactions take place.
	Type *MarketType11Choice `xml:"Tp,omitempty"`
}

func (m *MarketIdentification77) AddIdentification() *MarketIdentification3Choice {
	m.Identification = new(MarketIdentification3Choice)
	return m.Identification
}

func (m *MarketIdentification77) AddType() *MarketType11Choice {
	m.Type = new(MarketType11Choice)
	return m.Type
}

// Context, or geographic environment, in which trading parties may meet in order to negotiate and execute trades among themselves.
type MarketIdentification78 struct {

	// Code allocated to places of trade, ie, stock exchanges, regulated markets, eg, Electronic Trading Platforms (ECN), and unregulated markets, eg, Automated Trading Systems (ATS), as sources of prices and related information, in order to facilitate automated processing.
	Identification *MarketIdentification1Choice `xml:"Id,omitempty"`

	// Nature of a market in which transactions take place.
	Type *MarketType3Choice `xml:"Tp"`
}

func (m *MarketIdentification78) AddIdentification() *MarketIdentification1Choice {
	m.Identification = new(MarketIdentification1Choice)
	return m.Identification
}

func (m *MarketIdentification78) AddType() *MarketType3Choice {
	m.Type = new(MarketType3Choice)
	return m.Type
}

// Context, or geographic environment, in which trading parties may meet in order to negotiate and execute trades among themselves.
type MarketIdentification79 struct {

	// Code allocated to places of trade, ie, stock exchanges, regulated markets, eg, Electronic Trading Platforms (ECN), and unregulated markets, eg, Automated Trading Systems (ATS), as sources of prices and related information, in order to facilitate automated processing.
	Identification *MarketIdentification3Choice `xml:"Id,omitempty"`

	// Nature of a market in which transactions take place.
	Type *MarketType11Choice `xml:"Tp,omitempty"`
}

func (m *MarketIdentification79) AddIdentification() *MarketIdentification3Choice {
	m.Identification = new(MarketIdentification3Choice)
	return m.Identification
}

func (m *MarketIdentification79) AddType() *MarketType11Choice {
	m.Type = new(MarketType11Choice)
	return m.Type
}

// Context, or geographic environment, in which trading parties may meet in order to negotiate and execute trades among themselves.
type MarketIdentification80 struct {

	// Code allocated to places of trade, ie, stock exchanges, regulated markets, eg, Electronic Trading Platforms (ECN), and unregulated markets, eg, Automated Trading Systems (ATS), as sources of prices and related information, in order to facilitate automated processing.
	Identification *MarketIdentification3Choice `xml:"Id,omitempty"`

	// Nature of a market in which transactions take place.
	Type *MarketType12Choice `xml:"Tp,omitempty"`
}

func (m *MarketIdentification80) AddIdentification() *MarketIdentification3Choice {
	m.Identification = new(MarketIdentification3Choice)
	return m.Identification
}

func (m *MarketIdentification80) AddType() *MarketType12Choice {
	m.Type = new(MarketType12Choice)
	return m.Type
}

// Context, or geographic environment, in which trading parties may meet in order to negotiate and execute trades among themselves.
type MarketIdentification84 struct {

	// Code allocated to places of trade, ie, stock exchanges, regulated markets, eg, Electronic Trading Platforms (ECN), and unregulated markets, eg, Automated Trading Systems (ATS), as sources of prices and related information, in order to facilitate automated processing.
	Identification *MarketIdentification1Choice `xml:"Id,omitempty"`

	// Nature of a market in which transactions take place.
	Type *MarketType8Choice `xml:"Tp"`
}

func (m *MarketIdentification84) AddIdentification() *MarketIdentification1Choice {
	m.Identification = new(MarketIdentification1Choice)
	return m.Identification
}

func (m *MarketIdentification84) AddType() *MarketType8Choice {
	m.Type = new(MarketType8Choice)
	return m.Type
}

// Context, or geographic environment, in which trading parties may meet in order to negotiate and execute trades among themselves.
type MarketIdentification85 struct {

	// Code allocated to places of trade, ie, stock exchanges, regulated markets, eg, Electronic Trading Platforms (ECN), and unregulated markets, eg, Automated Trading Systems (ATS), as sources of prices and related information, in order to facilitate automated processing.
	Identification *MarketIdentification1Choice `xml:"Id,omitempty"`

	// Nature of a market in which transactions take place.
	Type *MarketType9Choice `xml:"Tp"`
}

func (m *MarketIdentification85) AddIdentification() *MarketIdentification1Choice {
	m.Identification = new(MarketIdentification1Choice)
	return m.Identification
}

func (m *MarketIdentification85) AddType() *MarketType9Choice {
	m.Type = new(MarketType9Choice)
	return m.Type
}

// Identifies the market and purpose for the settlement.
type MarketIdentification87 struct {

	// Country in which the financial instrument is to be settled.
	Country *CountryCode `xml:"Ctry"`

	// Type of instrument covered by the SSI instruction.
	ClassificationType *ClassificationType1Choice `xml:"ClssfctnTp"`

	// Purpose of the instruction,  for example, whether for regular payments, margin payments related to a collateral movement, securities settlements , securities lending.
	SettlementPurpose *Purpose3Choice `xml:"SttlmPurp,omitempty"`
}

func (m *MarketIdentification87) SetCountry(value string) {
	m.Country = (*CountryCode)(&value)
}

func (m *MarketIdentification87) AddClassificationType() *ClassificationType1Choice {
	m.ClassificationType = new(ClassificationType1Choice)
	return m.ClassificationType
}

func (m *MarketIdentification87) AddSettlementPurpose() *Purpose3Choice {
	m.SettlementPurpose = new(Purpose3Choice)
	return m.SettlementPurpose
}

// Context, or geographic environment, in which trading parties may meet in order to negotiate and execute trades among themselves.
type MarketIdentification89 struct {

	// Code allocated to places of trade (stock exchanges), to regulated markets (for example, Electronic Trading Platforms - ECN), and to unregulated markets (for example, Automated Trading Systems - ATS), as sources of prices and related information, in order to facilitate automated processing.
	Identification *MarketIdentification1Choice `xml:"Id,omitempty"`

	// Nature of a market in which transactions take place.
	Type *MarketType15Choice `xml:"Tp"`
}

func (m *MarketIdentification89) AddIdentification() *MarketIdentification1Choice {
	m.Identification = new(MarketIdentification1Choice)
	return m.Identification
}

func (m *MarketIdentification89) AddType() *MarketType15Choice {
	m.Type = new(MarketType15Choice)
	return m.Type
}

// Context, or geographic environment, in which trading parties may meet in order to negotiate and execute trades among themselves.
type MarketIdentification90 struct {

	// Code allocated to places of trade, ie, stock exchanges, regulated markets, eg, Electronic Trading Platforms (ECN), and unregulated markets, eg, Automated Trading Systems (ATS), as sources of prices and related information, in order to facilitate automated processing.
	Identification *MarketIdentification2Choice `xml:"Id,omitempty"`

	// Nature of a market in which transactions take place.
	Type *MarketType16Choice `xml:"Tp"`
}

func (m *MarketIdentification90) AddIdentification() *MarketIdentification2Choice {
	m.Identification = new(MarketIdentification2Choice)
	return m.Identification
}

func (m *MarketIdentification90) AddType() *MarketType16Choice {
	m.Type = new(MarketType16Choice)
	return m.Type
}

// Context, or geographic environment, in which trading parties may meet in order to negotiate and execute trades among themselves.
type MarketIdentification91 struct {

	// Code allocated to places of trade (stock exchanges), to regulated markets (for example, Electronic Trading Platforms - ECN), and to unregulated markets (for example, Automated Trading Systems - ATS), as sources of prices and related information, in order to facilitate automated processing.
	Identification *MarketIdentification2Choice `xml:"Id,omitempty"`

	// Nature of a market in which transactions take place.
	Type *MarketType17Choice `xml:"Tp"`
}

func (m *MarketIdentification91) AddIdentification() *MarketIdentification2Choice {
	m.Identification = new(MarketIdentification2Choice)
	return m.Identification
}

func (m *MarketIdentification91) AddType() *MarketType17Choice {
	m.Type = new(MarketType17Choice)
	return m.Type
}
