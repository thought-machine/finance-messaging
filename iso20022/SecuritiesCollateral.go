package iso20022

// Provides details about the securities posted as collateral.
type SecuritiesCollateral2 struct {

	// Identifies the register number of the collateral deposit assigned by the central counterparty.
	AssetNumber *Max35Text `xml:"AsstNb,omitempty"`

	// Identification of a security.
	SecurityIdentification *SecurityIdentification14 `xml:"SctyId"`

	// Planned final repayment date at the time of issuance.
	MaturityDate *DateAndDateTimeChoice `xml:"MtrtyDt,omitempty"`

	// Indicates whether the collateral is proprietarily owned or client owned.
	CollateralOwnership *CollateralOwnership1 `xml:"CollOwnrsh,omitempty"`

	// Indicates that the collateral posted in the clearing house covers the margin until a specific timeframe.
	LimitedCoverageIndicator *YesNoIndicator `xml:"LtdCvrgInd,omitempty"`

	// Quantity of securities collateral.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty"`

	// Quantity blocked by the central counterparty for any reasonable reason ( for example for judicial reasons). In this case the investor can not withdraw or distribute this collateral.
	BlockedQuantity *FinancialInstrumentQuantity1Choice `xml:"BlckdQty,omitempty"`

	// Indicates the price of the security.
	Price *Price2 `xml:"Pric,omitempty"`

	// Value of the collateral based on current market prices.
	MarketValue *ActiveCurrencyAndAmount `xml:"MktVal,omitempty"`

	// Haircut or valuation factor on the security expressed as a percentage.
	Haircut *PercentageRate `xml:"Hrcut,omitempty"`

	// Value of the collateral after taking into account the haircut.
	CollateralValue *ActiveCurrencyAndAmount `xml:"CollVal,omitempty"`

	// Valuation date of the securities taken as collateral.
	ValueDate *ISODate `xml:"ValDt,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount19 `xml:"SfkpgAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat7Choice `xml:"SfkpgPlc"`
}

func (s *SecuritiesCollateral2) SetAssetNumber(value string) {
	s.AssetNumber = (*Max35Text)(&value)
}

func (s *SecuritiesCollateral2) AddSecurityIdentification() *SecurityIdentification14 {
	s.SecurityIdentification = new(SecurityIdentification14)
	return s.SecurityIdentification
}

func (s *SecuritiesCollateral2) AddMaturityDate() *DateAndDateTimeChoice {
	s.MaturityDate = new(DateAndDateTimeChoice)
	return s.MaturityDate
}

func (s *SecuritiesCollateral2) AddCollateralOwnership() *CollateralOwnership1 {
	s.CollateralOwnership = new(CollateralOwnership1)
	return s.CollateralOwnership
}

func (s *SecuritiesCollateral2) SetLimitedCoverageIndicator(value string) {
	s.LimitedCoverageIndicator = (*YesNoIndicator)(&value)
}

func (s *SecuritiesCollateral2) AddQuantity() *FinancialInstrumentQuantity1Choice {
	s.Quantity = new(FinancialInstrumentQuantity1Choice)
	return s.Quantity
}

func (s *SecuritiesCollateral2) AddBlockedQuantity() *FinancialInstrumentQuantity1Choice {
	s.BlockedQuantity = new(FinancialInstrumentQuantity1Choice)
	return s.BlockedQuantity
}

func (s *SecuritiesCollateral2) AddPrice() *Price2 {
	s.Price = new(Price2)
	return s.Price
}

func (s *SecuritiesCollateral2) SetMarketValue(value, currency string) {
	s.MarketValue = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SecuritiesCollateral2) SetHaircut(value string) {
	s.Haircut = (*PercentageRate)(&value)
}

func (s *SecuritiesCollateral2) SetCollateralValue(value, currency string) {
	s.CollateralValue = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SecuritiesCollateral2) SetValueDate(value string) {
	s.ValueDate = (*ISODate)(&value)
}

func (s *SecuritiesCollateral2) AddSafekeepingAccount() *SecuritiesAccount19 {
	s.SafekeepingAccount = new(SecuritiesAccount19)
	return s.SafekeepingAccount
}

func (s *SecuritiesCollateral2) AddSafekeepingPlace() *SafekeepingPlaceFormat7Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat7Choice)
	return s.SafekeepingPlace
}

// Provides details about the securities posted as collateral.
type SecuritiesCollateral3 struct {

	// Provides the identification of the proposed collateral.
	CollateralIdentification *Max35Text `xml:"CollId,omitempty"`

	// Identification of a security.
	SecurityIdentification *SecurityIdentification14 `xml:"SctyId"`

	// Planned final repayment date at the time of issuance.
	MaturityDate *DateAndDateTimeChoice `xml:"MtrtyDt,omitempty"`

	// Indicates that the collateral posted in the clearing house covers the margin until a specific timeframe.
	LimitedCoverageIndicator *YesNoIndicator `xml:"LtdCvrgInd,omitempty"`

	// Quantity of securities collateral.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty"`

	// Indicates the price of the security.
	Price *Price2 `xml:"Pric,omitempty"`

	// Value of the collateral based on current market prices.
	MarketValue *ActiveCurrencyAndAmount `xml:"MktVal,omitempty"`

	// Haircut or valuation factor on the security expressed as a percentage.
	Haircut *PercentageRate `xml:"Hrcut,omitempty"`

	// Value of the collateral after taking into account the haircut.
	CollateralValue *ActiveCurrencyAndAmount `xml:"CollVal,omitempty"`

	// Valuation date of the securities taken as collateral.
	ValueDate *ISODate `xml:"ValDt,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount19 `xml:"SfkpgAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat7Choice `xml:"SfkpgPlc"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *SettlementDetails88 `xml:"SttlmParams,omitempty"`
}

func (s *SecuritiesCollateral3) SetCollateralIdentification(value string) {
	s.CollateralIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesCollateral3) AddSecurityIdentification() *SecurityIdentification14 {
	s.SecurityIdentification = new(SecurityIdentification14)
	return s.SecurityIdentification
}

func (s *SecuritiesCollateral3) AddMaturityDate() *DateAndDateTimeChoice {
	s.MaturityDate = new(DateAndDateTimeChoice)
	return s.MaturityDate
}

func (s *SecuritiesCollateral3) SetLimitedCoverageIndicator(value string) {
	s.LimitedCoverageIndicator = (*YesNoIndicator)(&value)
}

func (s *SecuritiesCollateral3) AddQuantity() *FinancialInstrumentQuantity1Choice {
	s.Quantity = new(FinancialInstrumentQuantity1Choice)
	return s.Quantity
}

func (s *SecuritiesCollateral3) AddPrice() *Price2 {
	s.Price = new(Price2)
	return s.Price
}

func (s *SecuritiesCollateral3) SetMarketValue(value, currency string) {
	s.MarketValue = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SecuritiesCollateral3) SetHaircut(value string) {
	s.Haircut = (*PercentageRate)(&value)
}

func (s *SecuritiesCollateral3) SetCollateralValue(value, currency string) {
	s.CollateralValue = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SecuritiesCollateral3) SetValueDate(value string) {
	s.ValueDate = (*ISODate)(&value)
}

func (s *SecuritiesCollateral3) AddSafekeepingAccount() *SecuritiesAccount19 {
	s.SafekeepingAccount = new(SecuritiesAccount19)
	return s.SafekeepingAccount
}

func (s *SecuritiesCollateral3) AddSafekeepingPlace() *SafekeepingPlaceFormat7Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat7Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesCollateral3) AddSettlementParameters() *SettlementDetails88 {
	s.SettlementParameters = new(SettlementDetails88)
	return s.SettlementParameters
}

// Provides details about the securities posted as collateral.
type SecuritiesCollateral4 struct {

	// Provides the identification of the proposed collateral.
	CollateralIdentification *Max35Text `xml:"CollId,omitempty"`

	// Identifies the register number of the collateral deposit assigned by the central counterparty.
	AssetNumber *Max35Text `xml:"AsstNb,omitempty"`

	// Identification of a security.
	SecurityIdentification *SecurityIdentification14 `xml:"SctyId"`

	// Planned final repayment date at the time of issuance.
	MaturityDate *DateAndDateTimeChoice `xml:"MtrtyDt,omitempty"`

	// Indicates that the collateral posted in the clearing house covers the margin until a specific timeframe.
	LimitedCoverageIndicator *YesNoIndicator `xml:"LtdCvrgInd,omitempty"`

	// Quantity of securities collateral.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty"`

	// Indicates the price of the security.
	Price *Price2 `xml:"Pric,omitempty"`

	// Value of the collateral based on current market prices.
	MarketValue *ActiveCurrencyAndAmount `xml:"MktVal,omitempty"`

	// Haircut or valuation factor on the security expressed as a percentage.
	Haircut *PercentageRate `xml:"Hrcut,omitempty"`

	// Value of the collateral after taking into account the haircut.
	CollateralValue *ActiveCurrencyAndAmount `xml:"CollVal,omitempty"`

	// Valuation date of the securities taken as collateral.
	ValueDate *ISODate `xml:"ValDt,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount19 `xml:"SfkpgAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat7Choice `xml:"SfkpgPlc"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *SettlementDetails88 `xml:"SttlmParams,omitempty"`
}

func (s *SecuritiesCollateral4) SetCollateralIdentification(value string) {
	s.CollateralIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesCollateral4) SetAssetNumber(value string) {
	s.AssetNumber = (*Max35Text)(&value)
}

func (s *SecuritiesCollateral4) AddSecurityIdentification() *SecurityIdentification14 {
	s.SecurityIdentification = new(SecurityIdentification14)
	return s.SecurityIdentification
}

func (s *SecuritiesCollateral4) AddMaturityDate() *DateAndDateTimeChoice {
	s.MaturityDate = new(DateAndDateTimeChoice)
	return s.MaturityDate
}

func (s *SecuritiesCollateral4) SetLimitedCoverageIndicator(value string) {
	s.LimitedCoverageIndicator = (*YesNoIndicator)(&value)
}

func (s *SecuritiesCollateral4) AddQuantity() *FinancialInstrumentQuantity1Choice {
	s.Quantity = new(FinancialInstrumentQuantity1Choice)
	return s.Quantity
}

func (s *SecuritiesCollateral4) AddPrice() *Price2 {
	s.Price = new(Price2)
	return s.Price
}

func (s *SecuritiesCollateral4) SetMarketValue(value, currency string) {
	s.MarketValue = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SecuritiesCollateral4) SetHaircut(value string) {
	s.Haircut = (*PercentageRate)(&value)
}

func (s *SecuritiesCollateral4) SetCollateralValue(value, currency string) {
	s.CollateralValue = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SecuritiesCollateral4) SetValueDate(value string) {
	s.ValueDate = (*ISODate)(&value)
}

func (s *SecuritiesCollateral4) AddSafekeepingAccount() *SecuritiesAccount19 {
	s.SafekeepingAccount = new(SecuritiesAccount19)
	return s.SafekeepingAccount
}

func (s *SecuritiesCollateral4) AddSafekeepingPlace() *SafekeepingPlaceFormat7Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat7Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesCollateral4) AddSettlementParameters() *SettlementDetails88 {
	s.SettlementParameters = new(SettlementDetails88)
	return s.SettlementParameters
}

// Provides details about the securities posted as collateral.
type SecuritiesCollateral5 struct {

	// Provides the identification of the proposed collateral.
	CollateralIdentification *Max35Text `xml:"CollId,omitempty"`

	// Identification of a security.
	SecurityIdentification *SecurityIdentification19 `xml:"SctyId"`

	// Planned final repayment date at the time of issuance.
	MaturityDate *DateAndDateTimeChoice `xml:"MtrtyDt,omitempty"`

	// Indicates that the collateral posted in the clearing house covers the margin until a specific timeframe.
	LimitedCoverageIndicator *YesNoIndicator `xml:"LtdCvrgInd,omitempty"`

	// Quantity of securities collateral.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty"`

	// Indicates the price of the security.
	Price *Price2 `xml:"Pric,omitempty"`

	// Value of the collateral based on current market prices.
	MarketValue *ActiveCurrencyAndAmount `xml:"MktVal,omitempty"`

	// Haircut or valuation factor on the security expressed as a percentage.
	Haircut *PercentageRate `xml:"Hrcut,omitempty"`

	// Value of the collateral after taking into account the haircut.
	CollateralValue *ActiveCurrencyAndAmount `xml:"CollVal,omitempty"`

	// Valuation date of the securities taken as collateral.
	ValueDate *ISODate `xml:"ValDt,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount19 `xml:"SfkpgAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat10Choice `xml:"SfkpgPlc"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *SettlementDetails102 `xml:"SttlmParams,omitempty"`
}

func (s *SecuritiesCollateral5) SetCollateralIdentification(value string) {
	s.CollateralIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesCollateral5) AddSecurityIdentification() *SecurityIdentification19 {
	s.SecurityIdentification = new(SecurityIdentification19)
	return s.SecurityIdentification
}

func (s *SecuritiesCollateral5) AddMaturityDate() *DateAndDateTimeChoice {
	s.MaturityDate = new(DateAndDateTimeChoice)
	return s.MaturityDate
}

func (s *SecuritiesCollateral5) SetLimitedCoverageIndicator(value string) {
	s.LimitedCoverageIndicator = (*YesNoIndicator)(&value)
}

func (s *SecuritiesCollateral5) AddQuantity() *FinancialInstrumentQuantity1Choice {
	s.Quantity = new(FinancialInstrumentQuantity1Choice)
	return s.Quantity
}

func (s *SecuritiesCollateral5) AddPrice() *Price2 {
	s.Price = new(Price2)
	return s.Price
}

func (s *SecuritiesCollateral5) SetMarketValue(value, currency string) {
	s.MarketValue = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SecuritiesCollateral5) SetHaircut(value string) {
	s.Haircut = (*PercentageRate)(&value)
}

func (s *SecuritiesCollateral5) SetCollateralValue(value, currency string) {
	s.CollateralValue = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SecuritiesCollateral5) SetValueDate(value string) {
	s.ValueDate = (*ISODate)(&value)
}

func (s *SecuritiesCollateral5) AddSafekeepingAccount() *SecuritiesAccount19 {
	s.SafekeepingAccount = new(SecuritiesAccount19)
	return s.SafekeepingAccount
}

func (s *SecuritiesCollateral5) AddSafekeepingPlace() *SafekeepingPlaceFormat10Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat10Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesCollateral5) AddSettlementParameters() *SettlementDetails102 {
	s.SettlementParameters = new(SettlementDetails102)
	return s.SettlementParameters
}

// Provides details about the securities posted as collateral.
type SecuritiesCollateral6 struct {

	// Identifies the register number of the collateral deposit assigned by the central counterparty.
	AssetNumber *Max35Text `xml:"AsstNb,omitempty"`

	// Identification of a security.
	SecurityIdentification *SecurityIdentification19 `xml:"SctyId"`

	// Planned final repayment date at the time of issuance.
	MaturityDate *DateAndDateTimeChoice `xml:"MtrtyDt,omitempty"`

	// Indicates whether the collateral is proprietarily owned or client owned.
	CollateralOwnership *CollateralOwnership2 `xml:"CollOwnrsh,omitempty"`

	// Indicates that the collateral posted in the clearing house covers the margin until a specific timeframe.
	LimitedCoverageIndicator *YesNoIndicator `xml:"LtdCvrgInd,omitempty"`

	// Quantity of securities collateral.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty"`

	// Quantity blocked by the central counterparty for any reasonable reason ( for example for judicial reasons). In this case the investor can not withdraw or distribute this collateral.
	BlockedQuantity *FinancialInstrumentQuantity1Choice `xml:"BlckdQty,omitempty"`

	// Indicates the price of the security.
	Price *Price2 `xml:"Pric,omitempty"`

	// Value of the collateral based on current market prices.
	MarketValue *ActiveCurrencyAndAmount `xml:"MktVal,omitempty"`

	// Haircut or valuation factor on the security expressed as a percentage.
	Haircut *PercentageRate `xml:"Hrcut,omitempty"`

	// Value of the collateral after taking into account the haircut.
	CollateralValue *ActiveCurrencyAndAmount `xml:"CollVal,omitempty"`

	// Valuation date of the securities taken as collateral.
	ValueDate *ISODate `xml:"ValDt,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount19 `xml:"SfkpgAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat10Choice `xml:"SfkpgPlc"`
}

func (s *SecuritiesCollateral6) SetAssetNumber(value string) {
	s.AssetNumber = (*Max35Text)(&value)
}

func (s *SecuritiesCollateral6) AddSecurityIdentification() *SecurityIdentification19 {
	s.SecurityIdentification = new(SecurityIdentification19)
	return s.SecurityIdentification
}

func (s *SecuritiesCollateral6) AddMaturityDate() *DateAndDateTimeChoice {
	s.MaturityDate = new(DateAndDateTimeChoice)
	return s.MaturityDate
}

func (s *SecuritiesCollateral6) AddCollateralOwnership() *CollateralOwnership2 {
	s.CollateralOwnership = new(CollateralOwnership2)
	return s.CollateralOwnership
}

func (s *SecuritiesCollateral6) SetLimitedCoverageIndicator(value string) {
	s.LimitedCoverageIndicator = (*YesNoIndicator)(&value)
}

func (s *SecuritiesCollateral6) AddQuantity() *FinancialInstrumentQuantity1Choice {
	s.Quantity = new(FinancialInstrumentQuantity1Choice)
	return s.Quantity
}

func (s *SecuritiesCollateral6) AddBlockedQuantity() *FinancialInstrumentQuantity1Choice {
	s.BlockedQuantity = new(FinancialInstrumentQuantity1Choice)
	return s.BlockedQuantity
}

func (s *SecuritiesCollateral6) AddPrice() *Price2 {
	s.Price = new(Price2)
	return s.Price
}

func (s *SecuritiesCollateral6) SetMarketValue(value, currency string) {
	s.MarketValue = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SecuritiesCollateral6) SetHaircut(value string) {
	s.Haircut = (*PercentageRate)(&value)
}

func (s *SecuritiesCollateral6) SetCollateralValue(value, currency string) {
	s.CollateralValue = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SecuritiesCollateral6) SetValueDate(value string) {
	s.ValueDate = (*ISODate)(&value)
}

func (s *SecuritiesCollateral6) AddSafekeepingAccount() *SecuritiesAccount19 {
	s.SafekeepingAccount = new(SecuritiesAccount19)
	return s.SafekeepingAccount
}

func (s *SecuritiesCollateral6) AddSafekeepingPlace() *SafekeepingPlaceFormat10Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat10Choice)
	return s.SafekeepingPlace
}

// Provides details about the securities posted as collateral.
type SecuritiesCollateral7 struct {

	// Provides the identification of the proposed collateral.
	CollateralIdentification *Max35Text `xml:"CollId,omitempty"`

	// Identifies the register number of the collateral deposit assigned by the central counterparty.
	AssetNumber *Max35Text `xml:"AsstNb,omitempty"`

	// Identification of a security.
	SecurityIdentification *SecurityIdentification19 `xml:"SctyId"`

	// Planned final repayment date at the time of issuance.
	MaturityDate *DateAndDateTimeChoice `xml:"MtrtyDt,omitempty"`

	// Indicates that the collateral posted in the clearing house covers the margin until a specific timeframe.
	LimitedCoverageIndicator *YesNoIndicator `xml:"LtdCvrgInd,omitempty"`

	// Quantity of securities collateral.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty"`

	// Indicates the price of the security.
	Price *Price2 `xml:"Pric,omitempty"`

	// Value of the collateral based on current market prices.
	MarketValue *ActiveCurrencyAndAmount `xml:"MktVal,omitempty"`

	// Haircut or valuation factor on the security expressed as a percentage.
	Haircut *PercentageRate `xml:"Hrcut,omitempty"`

	// Value of the collateral after taking into account the haircut.
	CollateralValue *ActiveCurrencyAndAmount `xml:"CollVal,omitempty"`

	// Valuation date of the securities taken as collateral.
	ValueDate *ISODate `xml:"ValDt,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount19 `xml:"SfkpgAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat10Choice `xml:"SfkpgPlc"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *SettlementDetails102 `xml:"SttlmParams,omitempty"`
}

func (s *SecuritiesCollateral7) SetCollateralIdentification(value string) {
	s.CollateralIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesCollateral7) SetAssetNumber(value string) {
	s.AssetNumber = (*Max35Text)(&value)
}

func (s *SecuritiesCollateral7) AddSecurityIdentification() *SecurityIdentification19 {
	s.SecurityIdentification = new(SecurityIdentification19)
	return s.SecurityIdentification
}

func (s *SecuritiesCollateral7) AddMaturityDate() *DateAndDateTimeChoice {
	s.MaturityDate = new(DateAndDateTimeChoice)
	return s.MaturityDate
}

func (s *SecuritiesCollateral7) SetLimitedCoverageIndicator(value string) {
	s.LimitedCoverageIndicator = (*YesNoIndicator)(&value)
}

func (s *SecuritiesCollateral7) AddQuantity() *FinancialInstrumentQuantity1Choice {
	s.Quantity = new(FinancialInstrumentQuantity1Choice)
	return s.Quantity
}

func (s *SecuritiesCollateral7) AddPrice() *Price2 {
	s.Price = new(Price2)
	return s.Price
}

func (s *SecuritiesCollateral7) SetMarketValue(value, currency string) {
	s.MarketValue = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SecuritiesCollateral7) SetHaircut(value string) {
	s.Haircut = (*PercentageRate)(&value)
}

func (s *SecuritiesCollateral7) SetCollateralValue(value, currency string) {
	s.CollateralValue = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SecuritiesCollateral7) SetValueDate(value string) {
	s.ValueDate = (*ISODate)(&value)
}

func (s *SecuritiesCollateral7) AddSafekeepingAccount() *SecuritiesAccount19 {
	s.SafekeepingAccount = new(SecuritiesAccount19)
	return s.SafekeepingAccount
}

func (s *SecuritiesCollateral7) AddSafekeepingPlace() *SafekeepingPlaceFormat10Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat10Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesCollateral7) AddSettlementParameters() *SettlementDetails102 {
	s.SettlementParameters = new(SettlementDetails102)
	return s.SettlementParameters
}

// Provides details about the securities posted as collateral.
type SecuritiesCollateral8 struct {

	// Provides the identification of the proposed collateral.
	CollateralIdentification *Max35Text `xml:"CollId,omitempty"`

	// Identification of a security.
	SecurityIdentification *SecurityIdentification19 `xml:"SctyId"`

	// Planned final repayment date at the time of issuance.
	MaturityDate *DateAndDateTimeChoice `xml:"MtrtyDt,omitempty"`

	// Indicates that the collateral posted in the clearing house covers the margin until a specific timeframe.
	LimitedCoverageIndicator *YesNoIndicator `xml:"LtdCvrgInd,omitempty"`

	// Quantity of securities collateral.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty"`

	// Indicates the price of the security.
	Price *Price2 `xml:"Pric,omitempty"`

	// Value of the collateral based on current market prices.
	MarketValue *ActiveCurrencyAndAmount `xml:"MktVal,omitempty"`

	// Haircut or valuation factor on the security expressed as a percentage.
	Haircut *PercentageRate `xml:"Hrcut,omitempty"`

	// Value of the collateral after taking into account the haircut.
	CollateralValue *ActiveCurrencyAndAmount `xml:"CollVal,omitempty"`

	// Valuation date of the securities taken as collateral.
	ValueDate *ISODate `xml:"ValDt,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount19 `xml:"SfkpgAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat10Choice `xml:"SfkpgPlc"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *SettlementDetails118 `xml:"SttlmParams,omitempty"`
}

func (s *SecuritiesCollateral8) SetCollateralIdentification(value string) {
	s.CollateralIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesCollateral8) AddSecurityIdentification() *SecurityIdentification19 {
	s.SecurityIdentification = new(SecurityIdentification19)
	return s.SecurityIdentification
}

func (s *SecuritiesCollateral8) AddMaturityDate() *DateAndDateTimeChoice {
	s.MaturityDate = new(DateAndDateTimeChoice)
	return s.MaturityDate
}

func (s *SecuritiesCollateral8) SetLimitedCoverageIndicator(value string) {
	s.LimitedCoverageIndicator = (*YesNoIndicator)(&value)
}

func (s *SecuritiesCollateral8) AddQuantity() *FinancialInstrumentQuantity1Choice {
	s.Quantity = new(FinancialInstrumentQuantity1Choice)
	return s.Quantity
}

func (s *SecuritiesCollateral8) AddPrice() *Price2 {
	s.Price = new(Price2)
	return s.Price
}

func (s *SecuritiesCollateral8) SetMarketValue(value, currency string) {
	s.MarketValue = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SecuritiesCollateral8) SetHaircut(value string) {
	s.Haircut = (*PercentageRate)(&value)
}

func (s *SecuritiesCollateral8) SetCollateralValue(value, currency string) {
	s.CollateralValue = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SecuritiesCollateral8) SetValueDate(value string) {
	s.ValueDate = (*ISODate)(&value)
}

func (s *SecuritiesCollateral8) AddSafekeepingAccount() *SecuritiesAccount19 {
	s.SafekeepingAccount = new(SecuritiesAccount19)
	return s.SafekeepingAccount
}

func (s *SecuritiesCollateral8) AddSafekeepingPlace() *SafekeepingPlaceFormat10Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat10Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesCollateral8) AddSettlementParameters() *SettlementDetails118 {
	s.SettlementParameters = new(SettlementDetails118)
	return s.SettlementParameters
}
