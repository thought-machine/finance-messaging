package iso20022

// Specifies the security option of a corporate event.
type SecuritiesOption1 struct {

	// Maximum quantity (or lot) of financial instrument that must be exercised or tendered.
	MaximumExercisableQuantity *FinancialInstrumentQuantity1Choice `xml:"MaxExrcblQty,omitempty"`

	// Minimum quantity (or lot) of financial instrument that must be exercised or tendered.
	MinimumExercisableQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblQty,omitempty"`

	// Minimum multiple quantity (or lot) of financial  instrument that must be exercised or tendered.
	MinimumExercisableMultipleQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblMltplQty,omitempty"`

	// Quantity of equity that makes up the new board lot.
	NewBoardLotQuantity *FinancialInstrumentQuantity1Choice `xml:"NewBrdLotQty,omitempty"`

	// New denomination of the equity following, for example, an increase or decrease in nominal value.
	NewDenominationQuantity *FinancialInstrumentQuantity1Choice `xml:"NewDnmtnQty,omitempty"`

	// Specifies that if an order is prorated holders of odd lots who tender their full position will not have tendered position prorated but rather accepted in full.
	FrontEndOddLotQuantity *FinancialInstrumentQuantity2Choice `xml:"FrntEndOddLotQty,omitempty"`

	// Represents the presence of a back end odd lot provision and the quantity of equity required after proration to be eligible for this privilege.
	BackEndOddLotQuantity *FinancialInstrumentQuantity2Choice `xml:"BckEndOddLotQty,omitempty"`
}

func (s *SecuritiesOption1) AddMaximumExercisableQuantity() *FinancialInstrumentQuantity1Choice {
	s.MaximumExercisableQuantity = new(FinancialInstrumentQuantity1Choice)
	return s.MaximumExercisableQuantity
}

func (s *SecuritiesOption1) AddMinimumExercisableQuantity() *FinancialInstrumentQuantity1Choice {
	s.MinimumExercisableQuantity = new(FinancialInstrumentQuantity1Choice)
	return s.MinimumExercisableQuantity
}

func (s *SecuritiesOption1) AddMinimumExercisableMultipleQuantity() *FinancialInstrumentQuantity1Choice {
	s.MinimumExercisableMultipleQuantity = new(FinancialInstrumentQuantity1Choice)
	return s.MinimumExercisableMultipleQuantity
}

func (s *SecuritiesOption1) AddNewBoardLotQuantity() *FinancialInstrumentQuantity1Choice {
	s.NewBoardLotQuantity = new(FinancialInstrumentQuantity1Choice)
	return s.NewBoardLotQuantity
}

func (s *SecuritiesOption1) AddNewDenominationQuantity() *FinancialInstrumentQuantity1Choice {
	s.NewDenominationQuantity = new(FinancialInstrumentQuantity1Choice)
	return s.NewDenominationQuantity
}

func (s *SecuritiesOption1) AddFrontEndOddLotQuantity() *FinancialInstrumentQuantity2Choice {
	s.FrontEndOddLotQuantity = new(FinancialInstrumentQuantity2Choice)
	return s.FrontEndOddLotQuantity
}

func (s *SecuritiesOption1) AddBackEndOddLotQuantity() *FinancialInstrumentQuantity2Choice {
	s.BackEndOddLotQuantity = new(FinancialInstrumentQuantity2Choice)
	return s.BackEndOddLotQuantity
}

// Provides information about the corporate action security option.
type SecuritiesOption13 struct {

	// Provides description of the financial instrument related to securities movement.
	SecurityDetails *FinancialInstrumentAttributes16 `xml:"SctyDtls"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies that the security identified  is a temporary security identification used for processing reasons, for example, contra security used in the US.
	TemporaryFinancialInstrumentIndicator *TemporaryFinancialInstrumentIndicator1Choice `xml:"TempFinInstrmInd,omitempty"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator1Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Quantity of securities based on the terms of the corporate action event and balance of underlying securities entitled to the account owner. (This quantity can be positive or negative).
	EntitledQuantity *Quantity6Choice `xml:"EntitldQty,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType1Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Period during which intermediate or outturn securities are tradable in a secondary market.
	TradingPeriod *Period3Choice `xml:"TradgPrd,omitempty"`

	// Provides information about the dates related to securities movement.
	DateDetails *SecurityDate5 `xml:"DtDtls"`

	// Provides information about the rates related to securities movement.
	RateDetails *CorporateActionRate17 `xml:"RateDtls,omitempty"`

	// Provides information about the prices related to securities movement.
	PriceDetails *CorporateActionPrice18 `xml:"PricDtls,omitempty"`
}

func (s *SecuritiesOption13) AddSecurityDetails() *FinancialInstrumentAttributes16 {
	s.SecurityDetails = new(FinancialInstrumentAttributes16)
	return s.SecurityDetails
}

func (s *SecuritiesOption13) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesOption13) AddTemporaryFinancialInstrumentIndicator() *TemporaryFinancialInstrumentIndicator1Choice {
	s.TemporaryFinancialInstrumentIndicator = new(TemporaryFinancialInstrumentIndicator1Choice)
	return s.TemporaryFinancialInstrumentIndicator
}

func (s *SecuritiesOption13) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator1Choice {
	s.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator1Choice)
	return s.NonEligibleProceedsIndicator
}

func (s *SecuritiesOption13) AddEntitledQuantity() *Quantity6Choice {
	s.EntitledQuantity = new(Quantity6Choice)
	return s.EntitledQuantity
}

func (s *SecuritiesOption13) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesOption13) AddFractionDisposition() *FractionDispositionType1Choice {
	s.FractionDisposition = new(FractionDispositionType1Choice)
	return s.FractionDisposition
}

func (s *SecuritiesOption13) SetCurrencyOption(value string) {
	s.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (s *SecuritiesOption13) AddTradingPeriod() *Period3Choice {
	s.TradingPeriod = new(Period3Choice)
	return s.TradingPeriod
}

func (s *SecuritiesOption13) AddDateDetails() *SecurityDate5 {
	s.DateDetails = new(SecurityDate5)
	return s.DateDetails
}

func (s *SecuritiesOption13) AddRateDetails() *CorporateActionRate17 {
	s.RateDetails = new(CorporateActionRate17)
	return s.RateDetails
}

func (s *SecuritiesOption13) AddPriceDetails() *CorporateActionPrice18 {
	s.PriceDetails = new(CorporateActionPrice18)
	return s.PriceDetails
}

// Provides information about the corporate action security option.
type SecuritiesOption14 struct {

	// Provides description of the financial instrument related to securities movement.
	SecurityDetails *FinancialInstrumentAttributes16 `xml:"SctyDtls"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies that the security identified  is a temporary security identification used for processing reasons, for example, contra security used in the US.
	TemporaryFinancialInstrumentIndicator *TemporaryFinancialInstrumentIndicator1Choice `xml:"TempFinInstrmInd,omitempty"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator1Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Quantity of securities based on the terms of the corporate action event and balance of underlying securities entitled to the account owner. (This quantity can be positive or negative).
	EntitledQuantity *Quantity6Choice `xml:"EntitldQty,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType1Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Period during which intermediate or outturn securities are tradable in a secondary market.
	TradingPeriod *Period3Choice `xml:"TradgPrd,omitempty"`

	// Provides information about the dates related to securities movement.
	DateDetails *SecurityDate5 `xml:"DtDtls"`

	// Provides information about the rates related to securities movement.
	RateDetails *CorporateActionRate17 `xml:"RateDtls,omitempty"`

	// Provides information about the prices related to securities movement.
	PriceDetails *CorporateActionPrice18 `xml:"PricDtls,omitempty"`
}

func (s *SecuritiesOption14) AddSecurityDetails() *FinancialInstrumentAttributes16 {
	s.SecurityDetails = new(FinancialInstrumentAttributes16)
	return s.SecurityDetails
}

func (s *SecuritiesOption14) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesOption14) AddTemporaryFinancialInstrumentIndicator() *TemporaryFinancialInstrumentIndicator1Choice {
	s.TemporaryFinancialInstrumentIndicator = new(TemporaryFinancialInstrumentIndicator1Choice)
	return s.TemporaryFinancialInstrumentIndicator
}

func (s *SecuritiesOption14) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator1Choice {
	s.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator1Choice)
	return s.NonEligibleProceedsIndicator
}

func (s *SecuritiesOption14) AddEntitledQuantity() *Quantity6Choice {
	s.EntitledQuantity = new(Quantity6Choice)
	return s.EntitledQuantity
}

func (s *SecuritiesOption14) AddFractionDisposition() *FractionDispositionType1Choice {
	s.FractionDisposition = new(FractionDispositionType1Choice)
	return s.FractionDisposition
}

func (s *SecuritiesOption14) SetCurrencyOption(value string) {
	s.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (s *SecuritiesOption14) AddTradingPeriod() *Period3Choice {
	s.TradingPeriod = new(Period3Choice)
	return s.TradingPeriod
}

func (s *SecuritiesOption14) AddDateDetails() *SecurityDate5 {
	s.DateDetails = new(SecurityDate5)
	return s.DateDetails
}

func (s *SecuritiesOption14) AddRateDetails() *CorporateActionRate17 {
	s.RateDetails = new(CorporateActionRate17)
	return s.RateDetails
}

func (s *SecuritiesOption14) AddPriceDetails() *CorporateActionPrice18 {
	s.PriceDetails = new(CorporateActionPrice18)
	return s.PriceDetails
}

// Specifies the security option of a corporate event.
type SecuritiesOption15 struct {

	// Maximum quantity (or lot) of financial instrument that must be exercised or tendered.
	MaximumExercisableQuantity *FinancialInstrumentQuantity1Choice `xml:"MaxExrcblQty,omitempty"`

	// Minimum quantity (or lot) of financial instrument that must be exercised or tendered.
	MinimumExercisableQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblQty,omitempty"`

	// Minimum multiple quantity (or lot) of financial  instrument that must be exercised or tendered.
	MinimumExercisableMultipleQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblMltplQty,omitempty"`

	// Quantity of equity that makes up the new board lot.
	NewBoardLotQuantity *FinancialInstrumentQuantity1Choice `xml:"NewBrdLotQty,omitempty"`

	// New denomination of the equity following, for example, an increase or decrease in nominal value.
	NewDenominationQuantity *FinancialInstrumentQuantity1Choice `xml:"NewDnmtnQty,omitempty"`

	// Specifies that if an order is prorated holders of odd lots who tender their full position will not have tendered position prorated but rather accepted in full.
	FrontEndOddLotQuantity *FinancialInstrumentQuantity16Choice `xml:"FrntEndOddLotQty,omitempty"`

	// Represents the presence of a back end odd lot provision and the quantity of equity required after proration to be eligible for this privilege.
	BackEndOddLotQuantity *FinancialInstrumentQuantity16Choice `xml:"BckEndOddLotQty,omitempty"`
}

func (s *SecuritiesOption15) AddMaximumExercisableQuantity() *FinancialInstrumentQuantity1Choice {
	s.MaximumExercisableQuantity = new(FinancialInstrumentQuantity1Choice)
	return s.MaximumExercisableQuantity
}

func (s *SecuritiesOption15) AddMinimumExercisableQuantity() *FinancialInstrumentQuantity1Choice {
	s.MinimumExercisableQuantity = new(FinancialInstrumentQuantity1Choice)
	return s.MinimumExercisableQuantity
}

func (s *SecuritiesOption15) AddMinimumExercisableMultipleQuantity() *FinancialInstrumentQuantity1Choice {
	s.MinimumExercisableMultipleQuantity = new(FinancialInstrumentQuantity1Choice)
	return s.MinimumExercisableMultipleQuantity
}

func (s *SecuritiesOption15) AddNewBoardLotQuantity() *FinancialInstrumentQuantity1Choice {
	s.NewBoardLotQuantity = new(FinancialInstrumentQuantity1Choice)
	return s.NewBoardLotQuantity
}

func (s *SecuritiesOption15) AddNewDenominationQuantity() *FinancialInstrumentQuantity1Choice {
	s.NewDenominationQuantity = new(FinancialInstrumentQuantity1Choice)
	return s.NewDenominationQuantity
}

func (s *SecuritiesOption15) AddFrontEndOddLotQuantity() *FinancialInstrumentQuantity16Choice {
	s.FrontEndOddLotQuantity = new(FinancialInstrumentQuantity16Choice)
	return s.FrontEndOddLotQuantity
}

func (s *SecuritiesOption15) AddBackEndOddLotQuantity() *FinancialInstrumentQuantity16Choice {
	s.BackEndOddLotQuantity = new(FinancialInstrumentQuantity16Choice)
	return s.BackEndOddLotQuantity
}

// Provides information about the corporate action security option.
type SecuritiesOption18 struct {

	// Identification of the financial instrument.
	SecurityIdentification *SecurityIdentification14 `xml:"SctyId"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies that the security identified  is a temporary security identification used for processing reasons, for example, contra security used in the US.
	TemporaryFinancialInstrumentIndicator *TemporaryFinancialInstrumentIndicator1Choice `xml:"TempFinInstrmInd,omitempty"`

	// Quantity of securities that have been posted (credit or debit) to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType4Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Provides information about the dates related to securities movement.
	DateDetails *SecurityDate6 `xml:"DtDtls"`

	// Provides information about the rates related to securities movement.
	RateDetails *CorporateActionRate21 `xml:"RateDtls,omitempty"`

	// Provides information about the prices related to securities movement.
	PriceDetails *CorporateActionPrice9 `xml:"PricDtls,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties15 `xml:"RcvgSttlmPties,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties15 `xml:"DlvrgSttlmPties,omitempty"`
}

func (s *SecuritiesOption18) AddSecurityIdentification() *SecurityIdentification14 {
	s.SecurityIdentification = new(SecurityIdentification14)
	return s.SecurityIdentification
}

func (s *SecuritiesOption18) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesOption18) AddTemporaryFinancialInstrumentIndicator() *TemporaryFinancialInstrumentIndicator1Choice {
	s.TemporaryFinancialInstrumentIndicator = new(TemporaryFinancialInstrumentIndicator1Choice)
	return s.TemporaryFinancialInstrumentIndicator
}

func (s *SecuritiesOption18) AddPostingQuantity() *Quantity6Choice {
	s.PostingQuantity = new(Quantity6Choice)
	return s.PostingQuantity
}

func (s *SecuritiesOption18) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesOption18) AddFractionDisposition() *FractionDispositionType4Choice {
	s.FractionDisposition = new(FractionDispositionType4Choice)
	return s.FractionDisposition
}

func (s *SecuritiesOption18) SetCurrencyOption(value string) {
	s.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (s *SecuritiesOption18) AddDateDetails() *SecurityDate6 {
	s.DateDetails = new(SecurityDate6)
	return s.DateDetails
}

func (s *SecuritiesOption18) AddRateDetails() *CorporateActionRate21 {
	s.RateDetails = new(CorporateActionRate21)
	return s.RateDetails
}

func (s *SecuritiesOption18) AddPriceDetails() *CorporateActionPrice9 {
	s.PriceDetails = new(CorporateActionPrice9)
	return s.PriceDetails
}

func (s *SecuritiesOption18) AddReceivingSettlementParties() *SettlementParties15 {
	s.ReceivingSettlementParties = new(SettlementParties15)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesOption18) AddDeliveringSettlementParties() *SettlementParties15 {
	s.DeliveringSettlementParties = new(SettlementParties15)
	return s.DeliveringSettlementParties
}

// Provides information about the corporate action security option.
type SecuritiesOption19 struct {

	// Identification of the financial instrument.
	SecurityIdentification *SecurityIdentification14 `xml:"SctyId"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Quantity of securities that have been posted (credit or debit) to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Date of the posting (credit or debit) to the account.
	PostingDate *DateAndDateTimeChoice `xml:"PstngDt"`

	// Date/Time of the posting (credit or debit) to the account that was initially communicated in the confirmation.
	OriginalPostingDate *DateAndDateTimeChoice `xml:"OrgnlPstngDt,omitempty"`
}

func (s *SecuritiesOption19) AddSecurityIdentification() *SecurityIdentification14 {
	s.SecurityIdentification = new(SecurityIdentification14)
	return s.SecurityIdentification
}

func (s *SecuritiesOption19) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesOption19) AddPostingQuantity() *Quantity6Choice {
	s.PostingQuantity = new(Quantity6Choice)
	return s.PostingQuantity
}

func (s *SecuritiesOption19) AddPostingDate() *DateAndDateTimeChoice {
	s.PostingDate = new(DateAndDateTimeChoice)
	return s.PostingDate
}

func (s *SecuritiesOption19) AddOriginalPostingDate() *DateAndDateTimeChoice {
	s.OriginalPostingDate = new(DateAndDateTimeChoice)
	return s.OriginalPostingDate
}

// Specifies the security option of a corporate event.
type SecuritiesOption2 struct {

	// Minimum quantity of securities to be accepted (used in the framework of conditional privilege on election). In case of proration, if this minimum quantity is not reached then the instruction is void.
	ConditionalQuantity *FinancialInstrumentQuantity1Choice `xml:"CondlQty,omitempty"`

	// Quantity instructed to be received over and above normal ensured entitlement.
	OverAndAboveNormalEnsuredEntitlementQuantity *FinancialInstrumentQuantity1Choice `xml:"OverAndAbovNrmlNsrdEntitlmntQty,omitempty"`

	// Specifies whether the quantity of financial instrument is a quantity of securities instructed or a quantity to receive.
	InstructedOrQuantityToReceive *InstructedOrQuantityToReceive1Choice `xml:"InstdOrQtyToRcv"`
}

func (s *SecuritiesOption2) AddConditionalQuantity() *FinancialInstrumentQuantity1Choice {
	s.ConditionalQuantity = new(FinancialInstrumentQuantity1Choice)
	return s.ConditionalQuantity
}

func (s *SecuritiesOption2) AddOverAndAboveNormalEnsuredEntitlementQuantity() *FinancialInstrumentQuantity1Choice {
	s.OverAndAboveNormalEnsuredEntitlementQuantity = new(FinancialInstrumentQuantity1Choice)
	return s.OverAndAboveNormalEnsuredEntitlementQuantity
}

func (s *SecuritiesOption2) AddInstructedOrQuantityToReceive() *InstructedOrQuantityToReceive1Choice {
	s.InstructedOrQuantityToReceive = new(InstructedOrQuantityToReceive1Choice)
	return s.InstructedOrQuantityToReceive
}

// Specifies the security option of a corporate event.
type SecuritiesOption23 struct {

	// Maximum quantity (or lot) of financial instrument that may be exercised or tendered.
	MaximumExercisableQuantity *FinancialInstrumentQuantity19Choice `xml:"MaxExrcblQty,omitempty"`

	// Minimum quantity (or lot) of financial instrument that may be exercised or tendered.
	MinimumExercisableQuantity *FinancialInstrumentQuantity19Choice `xml:"MinExrcblQty,omitempty"`

	// Minimum multiple quantity (or lot) of financial  instrument that may be exercised or tendered.
	MinimumExercisableMultipleQuantity *FinancialInstrumentQuantity20Choice `xml:"MinExrcblMltplQty,omitempty"`

	// Quantity of equity that makes up the new board lot.
	NewBoardLotQuantity *FinancialInstrumentQuantity20Choice `xml:"NewBrdLotQty,omitempty"`

	// New denomination of the equity following, for example, an increase or decrease in nominal value.
	NewDenominationQuantity *FinancialInstrumentQuantity20Choice `xml:"NewDnmtnQty,omitempty"`

	// Specifies that if an order is prorated holders of odd lots who tender their full position will not have tendered position prorated but rather accepted in full.
	FrontEndOddLotQuantity *FinancialInstrumentQuantity20Choice `xml:"FrntEndOddLotQty,omitempty"`

	// Represents the presence of a back end odd lot provision and the quantity of equity required after proration to be eligible for this privilege.
	BackEndOddLotQuantity *FinancialInstrumentQuantity20Choice `xml:"BckEndOddLotQty,omitempty"`
}

func (s *SecuritiesOption23) AddMaximumExercisableQuantity() *FinancialInstrumentQuantity19Choice {
	s.MaximumExercisableQuantity = new(FinancialInstrumentQuantity19Choice)
	return s.MaximumExercisableQuantity
}

func (s *SecuritiesOption23) AddMinimumExercisableQuantity() *FinancialInstrumentQuantity19Choice {
	s.MinimumExercisableQuantity = new(FinancialInstrumentQuantity19Choice)
	return s.MinimumExercisableQuantity
}

func (s *SecuritiesOption23) AddMinimumExercisableMultipleQuantity() *FinancialInstrumentQuantity20Choice {
	s.MinimumExercisableMultipleQuantity = new(FinancialInstrumentQuantity20Choice)
	return s.MinimumExercisableMultipleQuantity
}

func (s *SecuritiesOption23) AddNewBoardLotQuantity() *FinancialInstrumentQuantity20Choice {
	s.NewBoardLotQuantity = new(FinancialInstrumentQuantity20Choice)
	return s.NewBoardLotQuantity
}

func (s *SecuritiesOption23) AddNewDenominationQuantity() *FinancialInstrumentQuantity20Choice {
	s.NewDenominationQuantity = new(FinancialInstrumentQuantity20Choice)
	return s.NewDenominationQuantity
}

func (s *SecuritiesOption23) AddFrontEndOddLotQuantity() *FinancialInstrumentQuantity20Choice {
	s.FrontEndOddLotQuantity = new(FinancialInstrumentQuantity20Choice)
	return s.FrontEndOddLotQuantity
}

func (s *SecuritiesOption23) AddBackEndOddLotQuantity() *FinancialInstrumentQuantity20Choice {
	s.BackEndOddLotQuantity = new(FinancialInstrumentQuantity20Choice)
	return s.BackEndOddLotQuantity
}

// Provides information about the corporate action security option.
type SecuritiesOption24 struct {

	// Provides description of the financial instrument related to securities movement.
	SecurityDetails *FinancialInstrumentAttributes34 `xml:"SctyDtls"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies that the security identified  is a temporary security identification used for processing reasons, for example, contra security used in the US.
	TemporaryFinancialInstrumentIndicator *TemporaryFinancialInstrumentIndicator1Choice `xml:"TempFinInstrmInd,omitempty"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator1Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *IssuerTaxability1Code `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Quantity of securities based on the terms of the corporate action event and balance of underlying securities entitled to the account owner. (This quantity can be positive or negative).
	EntitledQuantity *Quantity6Choice `xml:"EntitldQty,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType1Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Period during which intermediate or outturn securities are tradable in a secondary market.
	TradingPeriod *Period3Choice `xml:"TradgPrd,omitempty"`

	// Provides information about the dates related to securities movement.
	DateDetails *SecurityDate5 `xml:"DtDtls"`

	// Provides information about the rates related to securities movement.
	RateDetails *CorporateActionRate28 `xml:"RateDtls,omitempty"`

	// Provides information about the prices related to securities movement.
	PriceDetails *CorporateActionPrice26 `xml:"PricDtls,omitempty"`
}

func (s *SecuritiesOption24) AddSecurityDetails() *FinancialInstrumentAttributes34 {
	s.SecurityDetails = new(FinancialInstrumentAttributes34)
	return s.SecurityDetails
}

func (s *SecuritiesOption24) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesOption24) AddTemporaryFinancialInstrumentIndicator() *TemporaryFinancialInstrumentIndicator1Choice {
	s.TemporaryFinancialInstrumentIndicator = new(TemporaryFinancialInstrumentIndicator1Choice)
	return s.TemporaryFinancialInstrumentIndicator
}

func (s *SecuritiesOption24) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator1Choice {
	s.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator1Choice)
	return s.NonEligibleProceedsIndicator
}

func (s *SecuritiesOption24) SetIssuerOfferorTaxabilityIndicator(value string) {
	s.IssuerOfferorTaxabilityIndicator = (*IssuerTaxability1Code)(&value)
}

func (s *SecuritiesOption24) AddEntitledQuantity() *Quantity6Choice {
	s.EntitledQuantity = new(Quantity6Choice)
	return s.EntitledQuantity
}

func (s *SecuritiesOption24) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesOption24) AddFractionDisposition() *FractionDispositionType1Choice {
	s.FractionDisposition = new(FractionDispositionType1Choice)
	return s.FractionDisposition
}

func (s *SecuritiesOption24) SetCurrencyOption(value string) {
	s.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (s *SecuritiesOption24) AddTradingPeriod() *Period3Choice {
	s.TradingPeriod = new(Period3Choice)
	return s.TradingPeriod
}

func (s *SecuritiesOption24) AddDateDetails() *SecurityDate5 {
	s.DateDetails = new(SecurityDate5)
	return s.DateDetails
}

func (s *SecuritiesOption24) AddRateDetails() *CorporateActionRate28 {
	s.RateDetails = new(CorporateActionRate28)
	return s.RateDetails
}

func (s *SecuritiesOption24) AddPriceDetails() *CorporateActionPrice26 {
	s.PriceDetails = new(CorporateActionPrice26)
	return s.PriceDetails
}

// Provides information about the corporate action security option.
type SecuritiesOption25 struct {

	// Provides description of the financial instrument related to securities movement.
	SecurityDetails *FinancialInstrumentAttributes34 `xml:"SctyDtls"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies that the security identified  is a temporary security identification used for processing reasons, for example, contra security used in the US.
	TemporaryFinancialInstrumentIndicator *TemporaryFinancialInstrumentIndicator1Choice `xml:"TempFinInstrmInd,omitempty"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator1Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *IssuerTaxability1Code `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Quantity of securities based on the terms of the corporate action event and balance of underlying securities entitled to the account owner. (This quantity can be positive or negative).
	EntitledQuantity *Quantity6Choice `xml:"EntitldQty,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType1Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Period during which intermediate or outturn securities are tradable in a secondary market.
	TradingPeriod *Period3Choice `xml:"TradgPrd,omitempty"`

	// Provides information about the dates related to securities movement.
	DateDetails *SecurityDate5 `xml:"DtDtls"`

	// Provides information about the rates related to securities movement.
	RateDetails *CorporateActionRate28 `xml:"RateDtls,omitempty"`

	// Provides information about the prices related to securities movement.
	PriceDetails *CorporateActionPrice31 `xml:"PricDtls,omitempty"`
}

func (s *SecuritiesOption25) AddSecurityDetails() *FinancialInstrumentAttributes34 {
	s.SecurityDetails = new(FinancialInstrumentAttributes34)
	return s.SecurityDetails
}

func (s *SecuritiesOption25) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesOption25) AddTemporaryFinancialInstrumentIndicator() *TemporaryFinancialInstrumentIndicator1Choice {
	s.TemporaryFinancialInstrumentIndicator = new(TemporaryFinancialInstrumentIndicator1Choice)
	return s.TemporaryFinancialInstrumentIndicator
}

func (s *SecuritiesOption25) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator1Choice {
	s.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator1Choice)
	return s.NonEligibleProceedsIndicator
}

func (s *SecuritiesOption25) SetIssuerOfferorTaxabilityIndicator(value string) {
	s.IssuerOfferorTaxabilityIndicator = (*IssuerTaxability1Code)(&value)
}

func (s *SecuritiesOption25) AddEntitledQuantity() *Quantity6Choice {
	s.EntitledQuantity = new(Quantity6Choice)
	return s.EntitledQuantity
}

func (s *SecuritiesOption25) AddFractionDisposition() *FractionDispositionType1Choice {
	s.FractionDisposition = new(FractionDispositionType1Choice)
	return s.FractionDisposition
}

func (s *SecuritiesOption25) SetCurrencyOption(value string) {
	s.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (s *SecuritiesOption25) AddTradingPeriod() *Period3Choice {
	s.TradingPeriod = new(Period3Choice)
	return s.TradingPeriod
}

func (s *SecuritiesOption25) AddDateDetails() *SecurityDate5 {
	s.DateDetails = new(SecurityDate5)
	return s.DateDetails
}

func (s *SecuritiesOption25) AddRateDetails() *CorporateActionRate28 {
	s.RateDetails = new(CorporateActionRate28)
	return s.RateDetails
}

func (s *SecuritiesOption25) AddPriceDetails() *CorporateActionPrice31 {
	s.PriceDetails = new(CorporateActionPrice31)
	return s.PriceDetails
}

// Provides information about the corporate action security option.
type SecuritiesOption26 struct {

	// Identification of the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies that the security identified  is a temporary security identification used for processing reasons, for example, contra security used in the US.
	TemporaryFinancialInstrumentIndicator *TemporaryFinancialInstrumentIndicator1Choice `xml:"TempFinInstrmInd,omitempty"`

	// Quantity of securities that have been posted (credit or debit) to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType4Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Provides information about the dates related to securities movement.
	DateDetails *SecurityDate6 `xml:"DtDtls"`

	// Provides information about the rates related to securities movement.
	RateDetails *CorporateActionRate29 `xml:"RateDtls,omitempty"`

	// Provides information about the prices related to securities movement.
	PriceDetails *CorporateActionPrice27 `xml:"PricDtls,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties24 `xml:"RcvgSttlmPties,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties24 `xml:"DlvrgSttlmPties,omitempty"`
}

func (s *SecuritiesOption26) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesOption26) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesOption26) AddTemporaryFinancialInstrumentIndicator() *TemporaryFinancialInstrumentIndicator1Choice {
	s.TemporaryFinancialInstrumentIndicator = new(TemporaryFinancialInstrumentIndicator1Choice)
	return s.TemporaryFinancialInstrumentIndicator
}

func (s *SecuritiesOption26) AddPostingQuantity() *Quantity6Choice {
	s.PostingQuantity = new(Quantity6Choice)
	return s.PostingQuantity
}

func (s *SecuritiesOption26) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesOption26) AddFractionDisposition() *FractionDispositionType4Choice {
	s.FractionDisposition = new(FractionDispositionType4Choice)
	return s.FractionDisposition
}

func (s *SecuritiesOption26) SetCurrencyOption(value string) {
	s.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (s *SecuritiesOption26) AddDateDetails() *SecurityDate6 {
	s.DateDetails = new(SecurityDate6)
	return s.DateDetails
}

func (s *SecuritiesOption26) AddRateDetails() *CorporateActionRate29 {
	s.RateDetails = new(CorporateActionRate29)
	return s.RateDetails
}

func (s *SecuritiesOption26) AddPriceDetails() *CorporateActionPrice27 {
	s.PriceDetails = new(CorporateActionPrice27)
	return s.PriceDetails
}

func (s *SecuritiesOption26) AddReceivingSettlementParties() *SettlementParties24 {
	s.ReceivingSettlementParties = new(SettlementParties24)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesOption26) AddDeliveringSettlementParties() *SettlementParties24 {
	s.DeliveringSettlementParties = new(SettlementParties24)
	return s.DeliveringSettlementParties
}

// Provides information about the corporate action security option.
type SecuritiesOption27 struct {

	// Identification of the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Quantity of securities that have been posted (credit or debit) to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Date of the posting (credit or debit) to the account.
	PostingDate *DateAndDateTimeChoice `xml:"PstngDt"`

	// Date/Time of the posting (credit or debit) to the account that was initially communicated in the confirmation.
	OriginalPostingDate *DateAndDateTimeChoice `xml:"OrgnlPstngDt,omitempty"`
}

func (s *SecuritiesOption27) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesOption27) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesOption27) AddPostingQuantity() *Quantity6Choice {
	s.PostingQuantity = new(Quantity6Choice)
	return s.PostingQuantity
}

func (s *SecuritiesOption27) AddPostingDate() *DateAndDateTimeChoice {
	s.PostingDate = new(DateAndDateTimeChoice)
	return s.PostingDate
}

func (s *SecuritiesOption27) AddOriginalPostingDate() *DateAndDateTimeChoice {
	s.OriginalPostingDate = new(DateAndDateTimeChoice)
	return s.OriginalPostingDate
}

// Provides information about the corporate action security option.
type SecuritiesOption3 struct {

	// Identification of the financial instrument.
	SecurityIdentification *SecurityIdentification11 `xml:"SctyId"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies that the security identified  is a temporary security identification used for processing reasons, for example, contra security used in the US.
	TemporaryFinancialInstrumentIndicator *TemporaryFinancialInstrumentIndicator1Choice `xml:"TempFinInstrmInd,omitempty"`

	// Quantity of securities that have been posted (credit or debit) to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType4Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Provides information about the dates related to securities movement.
	DateDetails *SecurityDate1 `xml:"DtDtls"`

	// Provides information about the rates related to securities movement.
	RateDetails *CorporateActionRate6 `xml:"RateDtls,omitempty"`

	// Provides information about the prices related to securities movement.
	PriceDetails *CorporateActionPrice9 `xml:"PricDtls,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties4 `xml:"RcvgSttlmPties,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties4 `xml:"DlvrgSttlmPties,omitempty"`
}

func (s *SecuritiesOption3) AddSecurityIdentification() *SecurityIdentification11 {
	s.SecurityIdentification = new(SecurityIdentification11)
	return s.SecurityIdentification
}

func (s *SecuritiesOption3) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesOption3) AddTemporaryFinancialInstrumentIndicator() *TemporaryFinancialInstrumentIndicator1Choice {
	s.TemporaryFinancialInstrumentIndicator = new(TemporaryFinancialInstrumentIndicator1Choice)
	return s.TemporaryFinancialInstrumentIndicator
}

func (s *SecuritiesOption3) AddPostingQuantity() *Quantity6Choice {
	s.PostingQuantity = new(Quantity6Choice)
	return s.PostingQuantity
}

func (s *SecuritiesOption3) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesOption3) AddFractionDisposition() *FractionDispositionType4Choice {
	s.FractionDisposition = new(FractionDispositionType4Choice)
	return s.FractionDisposition
}

func (s *SecuritiesOption3) SetCurrencyOption(value string) {
	s.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (s *SecuritiesOption3) AddDateDetails() *SecurityDate1 {
	s.DateDetails = new(SecurityDate1)
	return s.DateDetails
}

func (s *SecuritiesOption3) AddRateDetails() *CorporateActionRate6 {
	s.RateDetails = new(CorporateActionRate6)
	return s.RateDetails
}

func (s *SecuritiesOption3) AddPriceDetails() *CorporateActionPrice9 {
	s.PriceDetails = new(CorporateActionPrice9)
	return s.PriceDetails
}

func (s *SecuritiesOption3) AddReceivingSettlementParties() *SettlementParties4 {
	s.ReceivingSettlementParties = new(SettlementParties4)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesOption3) AddDeliveringSettlementParties() *SettlementParties4 {
	s.DeliveringSettlementParties = new(SettlementParties4)
	return s.DeliveringSettlementParties
}

// Provides information about the corporate action security option.
type SecuritiesOption33 struct {

	// Provides description of the financial instrument related to securities movement.
	SecurityDetails *FinancialInstrumentAttributes34 `xml:"SctyDtls"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies that the security identified  is a temporary security identification used for processing reasons, for example, contra security used in the US.
	TemporaryFinancialInstrumentIndicator *TemporaryFinancialInstrumentIndicator1Choice `xml:"TempFinInstrmInd,omitempty"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator1Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *IssuerTaxability1Code `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Indicates whether the securities are newly issued or not.
	NewSecuritiesIssuanceIndicator *NewSecuritiesIssuanceType1Code `xml:"NewSctiesIssncInd,omitempty"`

	// Quantity of securities based on the terms of the corporate action event and balance of underlying securities entitled to the account owner. (This quantity can be positive or negative).
	EntitledQuantity *Quantity6Choice `xml:"EntitldQty,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType19Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Period during which intermediate or outturn securities are tradable in a secondary market.
	TradingPeriod *Period3Choice `xml:"TradgPrd,omitempty"`

	// Provides information about the dates related to securities movement.
	DateDetails *SecurityDate9 `xml:"DtDtls"`

	// Provides information about the rates related to securities movement.
	RateDetails *CorporateActionRate28 `xml:"RateDtls,omitempty"`

	// Provides information about the prices related to securities movement.
	PriceDetails *CorporateActionPrice38 `xml:"PricDtls,omitempty"`
}

func (s *SecuritiesOption33) AddSecurityDetails() *FinancialInstrumentAttributes34 {
	s.SecurityDetails = new(FinancialInstrumentAttributes34)
	return s.SecurityDetails
}

func (s *SecuritiesOption33) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesOption33) AddTemporaryFinancialInstrumentIndicator() *TemporaryFinancialInstrumentIndicator1Choice {
	s.TemporaryFinancialInstrumentIndicator = new(TemporaryFinancialInstrumentIndicator1Choice)
	return s.TemporaryFinancialInstrumentIndicator
}

func (s *SecuritiesOption33) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator1Choice {
	s.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator1Choice)
	return s.NonEligibleProceedsIndicator
}

func (s *SecuritiesOption33) SetIssuerOfferorTaxabilityIndicator(value string) {
	s.IssuerOfferorTaxabilityIndicator = (*IssuerTaxability1Code)(&value)
}

func (s *SecuritiesOption33) SetNewSecuritiesIssuanceIndicator(value string) {
	s.NewSecuritiesIssuanceIndicator = (*NewSecuritiesIssuanceType1Code)(&value)
}

func (s *SecuritiesOption33) AddEntitledQuantity() *Quantity6Choice {
	s.EntitledQuantity = new(Quantity6Choice)
	return s.EntitledQuantity
}

func (s *SecuritiesOption33) AddFractionDisposition() *FractionDispositionType19Choice {
	s.FractionDisposition = new(FractionDispositionType19Choice)
	return s.FractionDisposition
}

func (s *SecuritiesOption33) SetCurrencyOption(value string) {
	s.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (s *SecuritiesOption33) AddTradingPeriod() *Period3Choice {
	s.TradingPeriod = new(Period3Choice)
	return s.TradingPeriod
}

func (s *SecuritiesOption33) AddDateDetails() *SecurityDate9 {
	s.DateDetails = new(SecurityDate9)
	return s.DateDetails
}

func (s *SecuritiesOption33) AddRateDetails() *CorporateActionRate28 {
	s.RateDetails = new(CorporateActionRate28)
	return s.RateDetails
}

func (s *SecuritiesOption33) AddPriceDetails() *CorporateActionPrice38 {
	s.PriceDetails = new(CorporateActionPrice38)
	return s.PriceDetails
}

// Provides information about the corporate action security option.
type SecuritiesOption35 struct {

	// Identification of the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies that the security identified  is a temporary security identification used for processing reasons, for example, contra security used in the US.
	TemporaryFinancialInstrumentIndicator *TemporaryFinancialInstrumentIndicator1Choice `xml:"TempFinInstrmInd,omitempty"`

	// Quantity of securities that have been posted (credit or debit) to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType23Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Provides information about the dates related to securities movement.
	DateDetails *SecurityDate6 `xml:"DtDtls"`

	// Provides information about the rates related to securities movement.
	RateDetails *CorporateActionRate29 `xml:"RateDtls,omitempty"`

	// Provides information about the prices related to securities movement.
	PriceDetails *CorporateActionPrice39 `xml:"PricDtls,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties24 `xml:"RcvgSttlmPties,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties24 `xml:"DlvrgSttlmPties,omitempty"`
}

func (s *SecuritiesOption35) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesOption35) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesOption35) AddTemporaryFinancialInstrumentIndicator() *TemporaryFinancialInstrumentIndicator1Choice {
	s.TemporaryFinancialInstrumentIndicator = new(TemporaryFinancialInstrumentIndicator1Choice)
	return s.TemporaryFinancialInstrumentIndicator
}

func (s *SecuritiesOption35) AddPostingQuantity() *Quantity6Choice {
	s.PostingQuantity = new(Quantity6Choice)
	return s.PostingQuantity
}

func (s *SecuritiesOption35) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesOption35) AddFractionDisposition() *FractionDispositionType23Choice {
	s.FractionDisposition = new(FractionDispositionType23Choice)
	return s.FractionDisposition
}

func (s *SecuritiesOption35) SetCurrencyOption(value string) {
	s.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (s *SecuritiesOption35) AddDateDetails() *SecurityDate6 {
	s.DateDetails = new(SecurityDate6)
	return s.DateDetails
}

func (s *SecuritiesOption35) AddRateDetails() *CorporateActionRate29 {
	s.RateDetails = new(CorporateActionRate29)
	return s.RateDetails
}

func (s *SecuritiesOption35) AddPriceDetails() *CorporateActionPrice39 {
	s.PriceDetails = new(CorporateActionPrice39)
	return s.PriceDetails
}

func (s *SecuritiesOption35) AddReceivingSettlementParties() *SettlementParties24 {
	s.ReceivingSettlementParties = new(SettlementParties24)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesOption35) AddDeliveringSettlementParties() *SettlementParties24 {
	s.DeliveringSettlementParties = new(SettlementParties24)
	return s.DeliveringSettlementParties
}

// Provides information about the corporate action security option.
type SecuritiesOption38 struct {

	// Provides description of the financial instrument related to securities movement.
	SecurityDetails *FinancialInstrumentAttributes34 `xml:"SctyDtls"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies that the security identified  is a temporary security identification used for processing reasons, for example, contra security used in the US.
	TemporaryFinancialInstrumentIndicator *TemporaryFinancialInstrumentIndicator1Choice `xml:"TempFinInstrmInd,omitempty"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator1Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *IssuerTaxability1Code `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Indicates whether the securities are newly issued or not.
	NewSecuritiesIssuanceIndicator *NewSecuritiesIssuanceType1Code `xml:"NewSctiesIssncInd,omitempty"`

	// Quantity of securities based on the terms of the corporate action event and balance of underlying securities entitled to the account owner. (This quantity can be positive or negative).
	EntitledQuantity *Quantity6Choice `xml:"EntitldQty,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType19Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Period during which intermediate or outturn securities are tradable in a secondary market.
	TradingPeriod *Period3Choice `xml:"TradgPrd,omitempty"`

	// Provides information about the dates related to securities movement.
	DateDetails *SecurityDate9 `xml:"DtDtls"`

	// Provides information about the rates related to securities movement.
	RateDetails *CorporateActionRate28 `xml:"RateDtls,omitempty"`

	// Provides information about the prices related to securities movement.
	PriceDetails *CorporateActionPrice38 `xml:"PricDtls,omitempty"`
}

func (s *SecuritiesOption38) AddSecurityDetails() *FinancialInstrumentAttributes34 {
	s.SecurityDetails = new(FinancialInstrumentAttributes34)
	return s.SecurityDetails
}

func (s *SecuritiesOption38) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesOption38) AddTemporaryFinancialInstrumentIndicator() *TemporaryFinancialInstrumentIndicator1Choice {
	s.TemporaryFinancialInstrumentIndicator = new(TemporaryFinancialInstrumentIndicator1Choice)
	return s.TemporaryFinancialInstrumentIndicator
}

func (s *SecuritiesOption38) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator1Choice {
	s.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator1Choice)
	return s.NonEligibleProceedsIndicator
}

func (s *SecuritiesOption38) SetIssuerOfferorTaxabilityIndicator(value string) {
	s.IssuerOfferorTaxabilityIndicator = (*IssuerTaxability1Code)(&value)
}

func (s *SecuritiesOption38) SetNewSecuritiesIssuanceIndicator(value string) {
	s.NewSecuritiesIssuanceIndicator = (*NewSecuritiesIssuanceType1Code)(&value)
}

func (s *SecuritiesOption38) AddEntitledQuantity() *Quantity6Choice {
	s.EntitledQuantity = new(Quantity6Choice)
	return s.EntitledQuantity
}

func (s *SecuritiesOption38) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesOption38) AddFractionDisposition() *FractionDispositionType19Choice {
	s.FractionDisposition = new(FractionDispositionType19Choice)
	return s.FractionDisposition
}

func (s *SecuritiesOption38) SetCurrencyOption(value string) {
	s.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (s *SecuritiesOption38) AddTradingPeriod() *Period3Choice {
	s.TradingPeriod = new(Period3Choice)
	return s.TradingPeriod
}

func (s *SecuritiesOption38) AddDateDetails() *SecurityDate9 {
	s.DateDetails = new(SecurityDate9)
	return s.DateDetails
}

func (s *SecuritiesOption38) AddRateDetails() *CorporateActionRate28 {
	s.RateDetails = new(CorporateActionRate28)
	return s.RateDetails
}

func (s *SecuritiesOption38) AddPriceDetails() *CorporateActionPrice38 {
	s.PriceDetails = new(CorporateActionPrice38)
	return s.PriceDetails
}

// Provides information about the corporate action security option.
type SecuritiesOption4 struct {

	// Provides description of the financial instrument related to securities movement.
	SecurityDetails *FinancialInstrumentAttributes5 `xml:"SctyDtls"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies that the security identified  is a temporary security identification used for processing reasons, for example, contra security used in the US.
	TemporaryFinancialInstrumentIndicator *TemporaryFinancialInstrumentIndicator1Choice `xml:"TempFinInstrmInd,omitempty"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator1Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Quantity of securities based on the terms of the corporate action event and balance of underlying securities entitled to the account owner. (This quantity can be positive or negative).
	EntitledQuantity *Quantity6Choice `xml:"EntitldQty,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType1Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Period during which intermediate or outturn securities are tradable in a secondary market.
	TradingPeriod *Period1Choice `xml:"TradgPrd,omitempty"`

	// Provides information about the dates related to securities movement.
	DateDetails *SecurityDate2 `xml:"DtDtls"`

	// Provides information about the rates related to securities movement.
	RateDetails *CorporateActionRate7 `xml:"RateDtls,omitempty"`

	// Provides information about the prices related to securities movement.
	PriceDetails *CorporateActionPrice10 `xml:"PricDtls,omitempty"`
}

func (s *SecuritiesOption4) AddSecurityDetails() *FinancialInstrumentAttributes5 {
	s.SecurityDetails = new(FinancialInstrumentAttributes5)
	return s.SecurityDetails
}

func (s *SecuritiesOption4) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesOption4) AddTemporaryFinancialInstrumentIndicator() *TemporaryFinancialInstrumentIndicator1Choice {
	s.TemporaryFinancialInstrumentIndicator = new(TemporaryFinancialInstrumentIndicator1Choice)
	return s.TemporaryFinancialInstrumentIndicator
}

func (s *SecuritiesOption4) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator1Choice {
	s.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator1Choice)
	return s.NonEligibleProceedsIndicator
}

func (s *SecuritiesOption4) AddEntitledQuantity() *Quantity6Choice {
	s.EntitledQuantity = new(Quantity6Choice)
	return s.EntitledQuantity
}

func (s *SecuritiesOption4) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesOption4) AddFractionDisposition() *FractionDispositionType1Choice {
	s.FractionDisposition = new(FractionDispositionType1Choice)
	return s.FractionDisposition
}

func (s *SecuritiesOption4) SetCurrencyOption(value string) {
	s.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (s *SecuritiesOption4) AddTradingPeriod() *Period1Choice {
	s.TradingPeriod = new(Period1Choice)
	return s.TradingPeriod
}

func (s *SecuritiesOption4) AddDateDetails() *SecurityDate2 {
	s.DateDetails = new(SecurityDate2)
	return s.DateDetails
}

func (s *SecuritiesOption4) AddRateDetails() *CorporateActionRate7 {
	s.RateDetails = new(CorporateActionRate7)
	return s.RateDetails
}

func (s *SecuritiesOption4) AddPriceDetails() *CorporateActionPrice10 {
	s.PriceDetails = new(CorporateActionPrice10)
	return s.PriceDetails
}

// Provides information about the corporate action security option.
type SecuritiesOption40 struct {

	// Provides description of the financial instrument related to securities movement.
	SecurityDetails *FinancialInstrumentAttributes49 `xml:"SctyDtls"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies that the security identified  is a temporary security identification used for processing reasons, for example, contra security used in the US.
	TemporaryFinancialInstrumentIndicator *TemporaryFinancialInstrumentIndicator1Choice `xml:"TempFinInstrmInd,omitempty"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator1Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *IssuerTaxability1Code `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Indicates whether the securities are newly issued or not.
	NewSecuritiesIssuanceIndicator *NewSecuritiesIssuanceType2Code `xml:"NewSctiesIssncInd,omitempty"`

	// Quantity of securities based on the terms of the corporate action event and balance of underlying securities entitled to the account owner. (This quantity can be positive or negative).
	EntitledQuantity *Quantity6Choice `xml:"EntitldQty,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType19Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Period during which intermediate or outturn securities are tradable in a secondary market.
	TradingPeriod *Period3Choice `xml:"TradgPrd,omitempty"`

	// Provides information about the dates related to securities movement.
	DateDetails *SecurityDate9 `xml:"DtDtls"`

	// Provides information about the rates related to securities movement.
	RateDetails *CorporateActionRate48 `xml:"RateDtls,omitempty"`

	// Provides information about the prices related to securities movement.
	PriceDetails *CorporateActionPrice43 `xml:"PricDtls,omitempty"`
}

func (s *SecuritiesOption40) AddSecurityDetails() *FinancialInstrumentAttributes49 {
	s.SecurityDetails = new(FinancialInstrumentAttributes49)
	return s.SecurityDetails
}

func (s *SecuritiesOption40) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesOption40) AddTemporaryFinancialInstrumentIndicator() *TemporaryFinancialInstrumentIndicator1Choice {
	s.TemporaryFinancialInstrumentIndicator = new(TemporaryFinancialInstrumentIndicator1Choice)
	return s.TemporaryFinancialInstrumentIndicator
}

func (s *SecuritiesOption40) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator1Choice {
	s.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator1Choice)
	return s.NonEligibleProceedsIndicator
}

func (s *SecuritiesOption40) SetIssuerOfferorTaxabilityIndicator(value string) {
	s.IssuerOfferorTaxabilityIndicator = (*IssuerTaxability1Code)(&value)
}

func (s *SecuritiesOption40) SetNewSecuritiesIssuanceIndicator(value string) {
	s.NewSecuritiesIssuanceIndicator = (*NewSecuritiesIssuanceType2Code)(&value)
}

func (s *SecuritiesOption40) AddEntitledQuantity() *Quantity6Choice {
	s.EntitledQuantity = new(Quantity6Choice)
	return s.EntitledQuantity
}

func (s *SecuritiesOption40) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesOption40) AddFractionDisposition() *FractionDispositionType19Choice {
	s.FractionDisposition = new(FractionDispositionType19Choice)
	return s.FractionDisposition
}

func (s *SecuritiesOption40) SetCurrencyOption(value string) {
	s.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (s *SecuritiesOption40) AddTradingPeriod() *Period3Choice {
	s.TradingPeriod = new(Period3Choice)
	return s.TradingPeriod
}

func (s *SecuritiesOption40) AddDateDetails() *SecurityDate9 {
	s.DateDetails = new(SecurityDate9)
	return s.DateDetails
}

func (s *SecuritiesOption40) AddRateDetails() *CorporateActionRate48 {
	s.RateDetails = new(CorporateActionRate48)
	return s.RateDetails
}

func (s *SecuritiesOption40) AddPriceDetails() *CorporateActionPrice43 {
	s.PriceDetails = new(CorporateActionPrice43)
	return s.PriceDetails
}

// Provides information about the corporate action security option.
type SecuritiesOption42 struct {

	// Identification of the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies that the security identified  is a temporary security identification used for processing reasons, for example, contra security used in the US.
	TemporaryFinancialInstrumentIndicator *TemporaryFinancialInstrumentIndicator1Choice `xml:"TempFinInstrmInd,omitempty"`

	// Indicates whether the securities are newly issued or not.
	NewSecuritiesIssuanceIndicator *NewSecuritiesIssuanceType3Code `xml:"NewSctiesIssncInd,omitempty"`

	// Quantity of securities that have been posted (credit or debit) to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType23Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Provides information about the dates related to securities movement.
	DateDetails *SecurityDate6 `xml:"DtDtls"`

	// Provides information about the rates related to securities movement.
	RateDetails *CorporateActionRate49 `xml:"RateDtls,omitempty"`

	// Provides information about the prices related to securities movement.
	PriceDetails *CorporateActionPrice45 `xml:"PricDtls,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties24 `xml:"RcvgSttlmPties,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties24 `xml:"DlvrgSttlmPties,omitempty"`
}

func (s *SecuritiesOption42) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesOption42) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesOption42) AddTemporaryFinancialInstrumentIndicator() *TemporaryFinancialInstrumentIndicator1Choice {
	s.TemporaryFinancialInstrumentIndicator = new(TemporaryFinancialInstrumentIndicator1Choice)
	return s.TemporaryFinancialInstrumentIndicator
}

func (s *SecuritiesOption42) SetNewSecuritiesIssuanceIndicator(value string) {
	s.NewSecuritiesIssuanceIndicator = (*NewSecuritiesIssuanceType3Code)(&value)
}

func (s *SecuritiesOption42) AddPostingQuantity() *Quantity6Choice {
	s.PostingQuantity = new(Quantity6Choice)
	return s.PostingQuantity
}

func (s *SecuritiesOption42) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesOption42) AddFractionDisposition() *FractionDispositionType23Choice {
	s.FractionDisposition = new(FractionDispositionType23Choice)
	return s.FractionDisposition
}

func (s *SecuritiesOption42) SetCurrencyOption(value string) {
	s.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (s *SecuritiesOption42) AddDateDetails() *SecurityDate6 {
	s.DateDetails = new(SecurityDate6)
	return s.DateDetails
}

func (s *SecuritiesOption42) AddRateDetails() *CorporateActionRate49 {
	s.RateDetails = new(CorporateActionRate49)
	return s.RateDetails
}

func (s *SecuritiesOption42) AddPriceDetails() *CorporateActionPrice45 {
	s.PriceDetails = new(CorporateActionPrice45)
	return s.PriceDetails
}

func (s *SecuritiesOption42) AddReceivingSettlementParties() *SettlementParties24 {
	s.ReceivingSettlementParties = new(SettlementParties24)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesOption42) AddDeliveringSettlementParties() *SettlementParties24 {
	s.DeliveringSettlementParties = new(SettlementParties24)
	return s.DeliveringSettlementParties
}

// Provides information about the corporate action security option.
type SecuritiesOption49 struct {

	// Provides description of the financial instrument related to securities movement.
	SecurityDetails *FinancialInstrumentAttributes67 `xml:"SctyDtls"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies that the security identified  is a temporary security identification used for processing reasons, for example, contra security used in the US.
	TemporaryFinancialInstrumentIndicator *TemporaryFinancialInstrumentIndicator3Choice `xml:"TempFinInstrmInd,omitempty"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator3Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *GenericIdentification30 `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Indicates whether the securities are newly issued or not.
	NewSecuritiesIssuanceIndicator *NewSecuritiesIssuanceType5Code `xml:"NewSctiesIssncInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification30 `xml:"IncmTp,omitempty"`

	// Specifies the basis for the reduced rate of withholding.
	ExemptionType []*GenericIdentification30 `xml:"XmptnTp,omitempty"`

	// Quantity of securities based on the terms of the corporate action event and balance of underlying securities entitled to the account owner. (This quantity can be positive or negative).
	EntitledQuantity *Quantity6Choice `xml:"EntitldQty,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat10Choice `xml:"SfkpgPlc,omitempty"`

	// Indicates the country from which the income originates.
	CountryOfIncomeSource *CountryCode `xml:"CtryOfIncmSrc,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType26Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Period during which intermediate or outturn securities are tradable in a secondary market.
	TradingPeriod *Period3Choice `xml:"TradgPrd,omitempty"`

	// Provides information about the dates related to securities movement.
	DateDetails *SecurityDate12 `xml:"DtDtls"`

	// Provides information about the rates related to securities movement.
	RateDetails *CorporateActionRate69 `xml:"RateDtls,omitempty"`

	// Provides information about the prices related to securities movement.
	PriceDetails *CorporateActionPrice56 `xml:"PricDtls,omitempty"`
}

func (s *SecuritiesOption49) AddSecurityDetails() *FinancialInstrumentAttributes67 {
	s.SecurityDetails = new(FinancialInstrumentAttributes67)
	return s.SecurityDetails
}

func (s *SecuritiesOption49) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesOption49) AddTemporaryFinancialInstrumentIndicator() *TemporaryFinancialInstrumentIndicator3Choice {
	s.TemporaryFinancialInstrumentIndicator = new(TemporaryFinancialInstrumentIndicator3Choice)
	return s.TemporaryFinancialInstrumentIndicator
}

func (s *SecuritiesOption49) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator3Choice {
	s.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator3Choice)
	return s.NonEligibleProceedsIndicator
}

func (s *SecuritiesOption49) AddIssuerOfferorTaxabilityIndicator() *GenericIdentification30 {
	s.IssuerOfferorTaxabilityIndicator = new(GenericIdentification30)
	return s.IssuerOfferorTaxabilityIndicator
}

func (s *SecuritiesOption49) SetNewSecuritiesIssuanceIndicator(value string) {
	s.NewSecuritiesIssuanceIndicator = (*NewSecuritiesIssuanceType5Code)(&value)
}

func (s *SecuritiesOption49) AddIncomeType() *GenericIdentification30 {
	s.IncomeType = new(GenericIdentification30)
	return s.IncomeType
}

func (s *SecuritiesOption49) AddExemptionType() *GenericIdentification30 {
	newValue := new(GenericIdentification30)
	s.ExemptionType = append(s.ExemptionType, newValue)
	return newValue
}

func (s *SecuritiesOption49) AddEntitledQuantity() *Quantity6Choice {
	s.EntitledQuantity = new(Quantity6Choice)
	return s.EntitledQuantity
}

func (s *SecuritiesOption49) AddSafekeepingPlace() *SafekeepingPlaceFormat10Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat10Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesOption49) SetCountryOfIncomeSource(value string) {
	s.CountryOfIncomeSource = (*CountryCode)(&value)
}

func (s *SecuritiesOption49) AddFractionDisposition() *FractionDispositionType26Choice {
	s.FractionDisposition = new(FractionDispositionType26Choice)
	return s.FractionDisposition
}

func (s *SecuritiesOption49) SetCurrencyOption(value string) {
	s.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (s *SecuritiesOption49) AddTradingPeriod() *Period3Choice {
	s.TradingPeriod = new(Period3Choice)
	return s.TradingPeriod
}

func (s *SecuritiesOption49) AddDateDetails() *SecurityDate12 {
	s.DateDetails = new(SecurityDate12)
	return s.DateDetails
}

func (s *SecuritiesOption49) AddRateDetails() *CorporateActionRate69 {
	s.RateDetails = new(CorporateActionRate69)
	return s.RateDetails
}

func (s *SecuritiesOption49) AddPriceDetails() *CorporateActionPrice56 {
	s.PriceDetails = new(CorporateActionPrice56)
	return s.PriceDetails
}

// Provides information about the corporate action security option.
type SecuritiesOption5 struct {

	// Identification of the financial instrument.
	SecurityIdentification *SecurityIdentification11 `xml:"SctyId"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Quantity of securities that have been posted (credit or debit) to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Date of the posting (credit or debit) to the account.
	PostingDate *DateAndDateTimeChoice `xml:"PstngDt"`

	// Date/Time of the posting (credit or debit) to the account that was initially communicated in the confirmation.
	OriginalPostingDate *DateAndDateTimeChoice `xml:"OrgnlPstngDt,omitempty"`
}

func (s *SecuritiesOption5) AddSecurityIdentification() *SecurityIdentification11 {
	s.SecurityIdentification = new(SecurityIdentification11)
	return s.SecurityIdentification
}

func (s *SecuritiesOption5) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesOption5) AddPostingQuantity() *Quantity6Choice {
	s.PostingQuantity = new(Quantity6Choice)
	return s.PostingQuantity
}

func (s *SecuritiesOption5) AddPostingDate() *DateAndDateTimeChoice {
	s.PostingDate = new(DateAndDateTimeChoice)
	return s.PostingDate
}

func (s *SecuritiesOption5) AddOriginalPostingDate() *DateAndDateTimeChoice {
	s.OriginalPostingDate = new(DateAndDateTimeChoice)
	return s.OriginalPostingDate
}

// Provides information about the corporate action security option.
type SecuritiesOption50 struct {

	// Identification of the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies that the security identified  is a temporary security identification used for processing reasons, for example, contra security used in the US.
	TemporaryFinancialInstrumentIndicator *TemporaryFinancialInstrumentIndicator3Choice `xml:"TempFinInstrmInd,omitempty"`

	// Indicates whether the securities are newly issued or not.
	NewSecuritiesIssuanceIndicator *NewSecuritiesIssuanceType6Code `xml:"NewSctiesIssncInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *GenericIdentification30 `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification30 `xml:"IncmTp,omitempty"`

	// Specifies the basis for the reduced rate of withholding.
	ExemptionType []*GenericIdentification30 `xml:"XmptnTp,omitempty"`

	// Indicates the country from which the income originates.
	CountryOfIncomeSource *CountryCode `xml:"CtryOfIncmSrc,omitempty"`

	// Quantity of securities that have been posted (credit or debit) to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat10Choice `xml:"SfkpgPlc,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType27Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Provides information about the dates related to securities movement.
	DateDetails *SecurityDate11 `xml:"DtDtls"`

	// Provides information about the rates related to securities movement.
	RateDetails *CorporateActionRate72 `xml:"RateDtls,omitempty"`

	// Provides information about the prices related to securities movement.
	PriceDetails *CorporateActionPrice59 `xml:"PricDtls,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties42 `xml:"RcvgSttlmPties,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties42 `xml:"DlvrgSttlmPties,omitempty"`
}

func (s *SecuritiesOption50) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesOption50) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesOption50) AddTemporaryFinancialInstrumentIndicator() *TemporaryFinancialInstrumentIndicator3Choice {
	s.TemporaryFinancialInstrumentIndicator = new(TemporaryFinancialInstrumentIndicator3Choice)
	return s.TemporaryFinancialInstrumentIndicator
}

func (s *SecuritiesOption50) SetNewSecuritiesIssuanceIndicator(value string) {
	s.NewSecuritiesIssuanceIndicator = (*NewSecuritiesIssuanceType6Code)(&value)
}

func (s *SecuritiesOption50) AddIssuerOfferorTaxabilityIndicator() *GenericIdentification30 {
	s.IssuerOfferorTaxabilityIndicator = new(GenericIdentification30)
	return s.IssuerOfferorTaxabilityIndicator
}

func (s *SecuritiesOption50) AddIncomeType() *GenericIdentification30 {
	s.IncomeType = new(GenericIdentification30)
	return s.IncomeType
}

func (s *SecuritiesOption50) AddExemptionType() *GenericIdentification30 {
	newValue := new(GenericIdentification30)
	s.ExemptionType = append(s.ExemptionType, newValue)
	return newValue
}

func (s *SecuritiesOption50) SetCountryOfIncomeSource(value string) {
	s.CountryOfIncomeSource = (*CountryCode)(&value)
}

func (s *SecuritiesOption50) AddPostingQuantity() *Quantity6Choice {
	s.PostingQuantity = new(Quantity6Choice)
	return s.PostingQuantity
}

func (s *SecuritiesOption50) AddSafekeepingPlace() *SafekeepingPlaceFormat10Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat10Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesOption50) AddFractionDisposition() *FractionDispositionType27Choice {
	s.FractionDisposition = new(FractionDispositionType27Choice)
	return s.FractionDisposition
}

func (s *SecuritiesOption50) SetCurrencyOption(value string) {
	s.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (s *SecuritiesOption50) AddDateDetails() *SecurityDate11 {
	s.DateDetails = new(SecurityDate11)
	return s.DateDetails
}

func (s *SecuritiesOption50) AddRateDetails() *CorporateActionRate72 {
	s.RateDetails = new(CorporateActionRate72)
	return s.RateDetails
}

func (s *SecuritiesOption50) AddPriceDetails() *CorporateActionPrice59 {
	s.PriceDetails = new(CorporateActionPrice59)
	return s.PriceDetails
}

func (s *SecuritiesOption50) AddReceivingSettlementParties() *SettlementParties42 {
	s.ReceivingSettlementParties = new(SettlementParties42)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesOption50) AddDeliveringSettlementParties() *SettlementParties42 {
	s.DeliveringSettlementParties = new(SettlementParties42)
	return s.DeliveringSettlementParties
}

// Specifies the security option of a corporate event.
type SecuritiesOption51 struct {

	// Maximum quantity of financial instrument that may be instructed.
	MaximumQuantityToInstruct *FinancialInstrumentQuantity19Choice `xml:"MaxQtyToInst,omitempty"`

	// Minimum quantity of financial instrument that may be instructed.
	MinimumQuantityToInstruct *FinancialInstrumentQuantity19Choice `xml:"MinQtyToInst,omitempty"`

	// Minimum multiple quantity of financial  instrument that may be instructed.
	MinimumMultipleQuantityToInstruct *FinancialInstrumentQuantity20Choice `xml:"MinMltplQtyToInst,omitempty"`

	// Quantity of equity that makes up the new board lot.
	NewBoardLotQuantity *FinancialInstrumentQuantity20Choice `xml:"NewBrdLotQty,omitempty"`

	// New denomination of the equity following, for example, an increase or decrease in nominal value.
	NewDenominationQuantity *FinancialInstrumentQuantity20Choice `xml:"NewDnmtnQty,omitempty"`

	// Specifies that if an order is prorated holders of odd lots who tender their full position will not have tendered position prorated but rather accepted in full.
	FrontEndOddLotQuantity *FinancialInstrumentQuantity20Choice `xml:"FrntEndOddLotQty,omitempty"`

	// Represents the presence of a back end odd lot provision and the quantity of equity required after proration to be eligible for this privilege.
	BackEndOddLotQuantity *FinancialInstrumentQuantity20Choice `xml:"BckEndOddLotQty,omitempty"`
}

func (s *SecuritiesOption51) AddMaximumQuantityToInstruct() *FinancialInstrumentQuantity19Choice {
	s.MaximumQuantityToInstruct = new(FinancialInstrumentQuantity19Choice)
	return s.MaximumQuantityToInstruct
}

func (s *SecuritiesOption51) AddMinimumQuantityToInstruct() *FinancialInstrumentQuantity19Choice {
	s.MinimumQuantityToInstruct = new(FinancialInstrumentQuantity19Choice)
	return s.MinimumQuantityToInstruct
}

func (s *SecuritiesOption51) AddMinimumMultipleQuantityToInstruct() *FinancialInstrumentQuantity20Choice {
	s.MinimumMultipleQuantityToInstruct = new(FinancialInstrumentQuantity20Choice)
	return s.MinimumMultipleQuantityToInstruct
}

func (s *SecuritiesOption51) AddNewBoardLotQuantity() *FinancialInstrumentQuantity20Choice {
	s.NewBoardLotQuantity = new(FinancialInstrumentQuantity20Choice)
	return s.NewBoardLotQuantity
}

func (s *SecuritiesOption51) AddNewDenominationQuantity() *FinancialInstrumentQuantity20Choice {
	s.NewDenominationQuantity = new(FinancialInstrumentQuantity20Choice)
	return s.NewDenominationQuantity
}

func (s *SecuritiesOption51) AddFrontEndOddLotQuantity() *FinancialInstrumentQuantity20Choice {
	s.FrontEndOddLotQuantity = new(FinancialInstrumentQuantity20Choice)
	return s.FrontEndOddLotQuantity
}

func (s *SecuritiesOption51) AddBackEndOddLotQuantity() *FinancialInstrumentQuantity20Choice {
	s.BackEndOddLotQuantity = new(FinancialInstrumentQuantity20Choice)
	return s.BackEndOddLotQuantity
}

// Specifies the security option of a corporate event.
type SecuritiesOption52 struct {

	// Minimum quantity of securities to be accepted (used in the framework of conditional privilege on election). In case of proration, if this minimum quantity is not reached then the instruction is void.
	ConditionalQuantity *FinancialInstrumentQuantity1Choice `xml:"CondlQty,omitempty"`

	// Quantity of securities to which this instruction applies.
	InstructedQuantity *Quantity20Choice `xml:"InstdQty"`
}

func (s *SecuritiesOption52) AddConditionalQuantity() *FinancialInstrumentQuantity1Choice {
	s.ConditionalQuantity = new(FinancialInstrumentQuantity1Choice)
	return s.ConditionalQuantity
}

func (s *SecuritiesOption52) AddInstructedQuantity() *Quantity20Choice {
	s.InstructedQuantity = new(Quantity20Choice)
	return s.InstructedQuantity
}

// Provides information about the corporate action security option.
type SecuritiesOption53 struct {

	// Identification of the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Quantity of securities that have been posted (credit or debit) to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Date of the posting (credit or debit) to the account.
	PostingDate *DateAndDateTimeChoice `xml:"PstngDt"`

	// Date/Time of the posting (credit or debit) to the account that was initially communicated in the confirmation.
	OriginalPostingDate *DateAndDateTimeChoice `xml:"OrgnlPstngDt,omitempty"`
}

func (s *SecuritiesOption53) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesOption53) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesOption53) AddPostingQuantity() *Quantity6Choice {
	s.PostingQuantity = new(Quantity6Choice)
	return s.PostingQuantity
}

func (s *SecuritiesOption53) AddPostingDate() *DateAndDateTimeChoice {
	s.PostingDate = new(DateAndDateTimeChoice)
	return s.PostingDate
}

func (s *SecuritiesOption53) AddOriginalPostingDate() *DateAndDateTimeChoice {
	s.OriginalPostingDate = new(DateAndDateTimeChoice)
	return s.OriginalPostingDate
}

// Specifies the security option of a corporate event.
type SecuritiesOption54 struct {

	// Minimum quantity of securities to be accepted (used in the framework of conditional privilege on election). In case of proration, if this minimum quantity is not reached then the instruction is void.
	ConditionalQuantity *FinancialInstrumentQuantity15Choice `xml:"CondlQty,omitempty"`

	// Quantity of securities to which this instruction applies.
	InstructedQuantity *Quantity40Choice `xml:"InstdQty"`
}

func (s *SecuritiesOption54) AddConditionalQuantity() *FinancialInstrumentQuantity15Choice {
	s.ConditionalQuantity = new(FinancialInstrumentQuantity15Choice)
	return s.ConditionalQuantity
}

func (s *SecuritiesOption54) AddInstructedQuantity() *Quantity40Choice {
	s.InstructedQuantity = new(Quantity40Choice)
	return s.InstructedQuantity
}

// Provides information about the corporate action security option.
type SecuritiesOption55 struct {

	// Identification of the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies that the security identified  is a temporary security identification used for processing reasons, for example, contra security used in the US.
	TemporaryFinancialInstrumentIndicator *TemporaryFinancialInstrumentIndicator4Choice `xml:"TempFinInstrmInd,omitempty"`

	// Indicates whether the securities are newly issued or not.
	NewSecuritiesIssuanceIndicator *NewSecuritiesIssuanceType6Code `xml:"NewSctiesIssncInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *GenericIdentification47 `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification47 `xml:"IncmTp,omitempty"`

	// Specifies the basis for the reduced rate of withholding.
	ExemptionType []*GenericIdentification47 `xml:"XmptnTp,omitempty"`

	// Indicates the country from which the income originates.
	CountryOfIncomeSource *CountryCode `xml:"CtryOfIncmSrc,omitempty"`

	// Quantity of securities that have been posted (credit or debit) to the safekeeping account.
	PostingQuantity *Quantity10Choice `xml:"PstngQty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat17Choice `xml:"SfkpgPlc,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType30Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Provides information about the dates related to securities movement.
	DateDetails *SecurityDate13 `xml:"DtDtls"`

	// Provides information about the rates related to securities movement.
	RateDetails *CorporateActionRate75 `xml:"RateDtls,omitempty"`

	// Provides information about the prices related to securities movement.
	PriceDetails *CorporateActionPrice64 `xml:"PricDtls,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties43 `xml:"RcvgSttlmPties,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties43 `xml:"DlvrgSttlmPties,omitempty"`
}

func (s *SecuritiesOption55) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesOption55) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesOption55) AddTemporaryFinancialInstrumentIndicator() *TemporaryFinancialInstrumentIndicator4Choice {
	s.TemporaryFinancialInstrumentIndicator = new(TemporaryFinancialInstrumentIndicator4Choice)
	return s.TemporaryFinancialInstrumentIndicator
}

func (s *SecuritiesOption55) SetNewSecuritiesIssuanceIndicator(value string) {
	s.NewSecuritiesIssuanceIndicator = (*NewSecuritiesIssuanceType6Code)(&value)
}

func (s *SecuritiesOption55) AddIssuerOfferorTaxabilityIndicator() *GenericIdentification47 {
	s.IssuerOfferorTaxabilityIndicator = new(GenericIdentification47)
	return s.IssuerOfferorTaxabilityIndicator
}

func (s *SecuritiesOption55) AddIncomeType() *GenericIdentification47 {
	s.IncomeType = new(GenericIdentification47)
	return s.IncomeType
}

func (s *SecuritiesOption55) AddExemptionType() *GenericIdentification47 {
	newValue := new(GenericIdentification47)
	s.ExemptionType = append(s.ExemptionType, newValue)
	return newValue
}

func (s *SecuritiesOption55) SetCountryOfIncomeSource(value string) {
	s.CountryOfIncomeSource = (*CountryCode)(&value)
}

func (s *SecuritiesOption55) AddPostingQuantity() *Quantity10Choice {
	s.PostingQuantity = new(Quantity10Choice)
	return s.PostingQuantity
}

func (s *SecuritiesOption55) AddSafekeepingPlace() *SafekeepingPlaceFormat17Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat17Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesOption55) AddFractionDisposition() *FractionDispositionType30Choice {
	s.FractionDisposition = new(FractionDispositionType30Choice)
	return s.FractionDisposition
}

func (s *SecuritiesOption55) SetCurrencyOption(value string) {
	s.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (s *SecuritiesOption55) AddDateDetails() *SecurityDate13 {
	s.DateDetails = new(SecurityDate13)
	return s.DateDetails
}

func (s *SecuritiesOption55) AddRateDetails() *CorporateActionRate75 {
	s.RateDetails = new(CorporateActionRate75)
	return s.RateDetails
}

func (s *SecuritiesOption55) AddPriceDetails() *CorporateActionPrice64 {
	s.PriceDetails = new(CorporateActionPrice64)
	return s.PriceDetails
}

func (s *SecuritiesOption55) AddReceivingSettlementParties() *SettlementParties43 {
	s.ReceivingSettlementParties = new(SettlementParties43)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesOption55) AddDeliveringSettlementParties() *SettlementParties43 {
	s.DeliveringSettlementParties = new(SettlementParties43)
	return s.DeliveringSettlementParties
}

// Specifies the security option of a corporate event.
type SecuritiesOption56 struct {

	// Maximum quantity of financial instrument that may be instructed.
	MaximumQuantityToInstruct *FinancialInstrumentQuantity21Choice `xml:"MaxQtyToInst,omitempty"`

	// Minimum quantity of financial instrument that may be instructed.
	MinimumQuantityToInstruct *FinancialInstrumentQuantity21Choice `xml:"MinQtyToInst,omitempty"`

	// Minimum multiple quantity of financial  instrument that may be instructed.
	MinimumMultipleQuantityToInstruct *FinancialInstrumentQuantity22Choice `xml:"MinMltplQtyToInst,omitempty"`

	// Quantity of equity that makes up the new board lot.
	NewBoardLotQuantity *FinancialInstrumentQuantity22Choice `xml:"NewBrdLotQty,omitempty"`

	// New denomination of the equity following, for example, an increase or decrease in nominal value.
	NewDenominationQuantity *FinancialInstrumentQuantity22Choice `xml:"NewDnmtnQty,omitempty"`

	// Specifies that if an order is prorated holders of odd lots who tender their full position will not have tendered position prorated but rather accepted in full.
	FrontEndOddLotQuantity *FinancialInstrumentQuantity22Choice `xml:"FrntEndOddLotQty,omitempty"`

	// Represents the presence of a back end odd lot provision and the quantity of equity required after proration to be eligible for this privilege.
	BackEndOddLotQuantity *FinancialInstrumentQuantity22Choice `xml:"BckEndOddLotQty,omitempty"`
}

func (s *SecuritiesOption56) AddMaximumQuantityToInstruct() *FinancialInstrumentQuantity21Choice {
	s.MaximumQuantityToInstruct = new(FinancialInstrumentQuantity21Choice)
	return s.MaximumQuantityToInstruct
}

func (s *SecuritiesOption56) AddMinimumQuantityToInstruct() *FinancialInstrumentQuantity21Choice {
	s.MinimumQuantityToInstruct = new(FinancialInstrumentQuantity21Choice)
	return s.MinimumQuantityToInstruct
}

func (s *SecuritiesOption56) AddMinimumMultipleQuantityToInstruct() *FinancialInstrumentQuantity22Choice {
	s.MinimumMultipleQuantityToInstruct = new(FinancialInstrumentQuantity22Choice)
	return s.MinimumMultipleQuantityToInstruct
}

func (s *SecuritiesOption56) AddNewBoardLotQuantity() *FinancialInstrumentQuantity22Choice {
	s.NewBoardLotQuantity = new(FinancialInstrumentQuantity22Choice)
	return s.NewBoardLotQuantity
}

func (s *SecuritiesOption56) AddNewDenominationQuantity() *FinancialInstrumentQuantity22Choice {
	s.NewDenominationQuantity = new(FinancialInstrumentQuantity22Choice)
	return s.NewDenominationQuantity
}

func (s *SecuritiesOption56) AddFrontEndOddLotQuantity() *FinancialInstrumentQuantity22Choice {
	s.FrontEndOddLotQuantity = new(FinancialInstrumentQuantity22Choice)
	return s.FrontEndOddLotQuantity
}

func (s *SecuritiesOption56) AddBackEndOddLotQuantity() *FinancialInstrumentQuantity22Choice {
	s.BackEndOddLotQuantity = new(FinancialInstrumentQuantity22Choice)
	return s.BackEndOddLotQuantity
}

// Provides information about the corporate action security option.
type SecuritiesOption57 struct {

	// Provides description of the financial instrument related to securities movement.
	SecurityDetails *FinancialInstrumentAttributes71 `xml:"SctyDtls"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies that the security identified  is a temporary security identification used for processing reasons, for example, contra security used in the US.
	TemporaryFinancialInstrumentIndicator *TemporaryFinancialInstrumentIndicator4Choice `xml:"TempFinInstrmInd,omitempty"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator4Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *GenericIdentification47 `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Indicates whether the securities are newly issued or not.
	NewSecuritiesIssuanceIndicator *NewSecuritiesIssuanceType5Code `xml:"NewSctiesIssncInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification47 `xml:"IncmTp,omitempty"`

	// Specifies the basis for the reduced rate of withholding.
	ExemptionType []*GenericIdentification47 `xml:"XmptnTp,omitempty"`

	// Quantity of securities based on the terms of the corporate action event and balance of underlying securities entitled to the account owner. (This quantity can be positive or negative).
	EntitledQuantity *Quantity10Choice `xml:"EntitldQty,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat17Choice `xml:"SfkpgPlc,omitempty"`

	// Indicates the country from which the income originates.
	CountryOfIncomeSource *CountryCode `xml:"CtryOfIncmSrc,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType31Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Period during which intermediate or outturn securities are tradable in a secondary market.
	TradingPeriod *Period3Choice `xml:"TradgPrd,omitempty"`

	// Provides information about the dates related to securities movement.
	DateDetails *SecurityDate14 `xml:"DtDtls"`

	// Provides information about the rates related to securities movement.
	RateDetails *CorporateActionRate77 `xml:"RateDtls,omitempty"`

	// Provides information about the prices related to securities movement.
	PriceDetails *CorporateActionPrice66 `xml:"PricDtls,omitempty"`
}

func (s *SecuritiesOption57) AddSecurityDetails() *FinancialInstrumentAttributes71 {
	s.SecurityDetails = new(FinancialInstrumentAttributes71)
	return s.SecurityDetails
}

func (s *SecuritiesOption57) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesOption57) AddTemporaryFinancialInstrumentIndicator() *TemporaryFinancialInstrumentIndicator4Choice {
	s.TemporaryFinancialInstrumentIndicator = new(TemporaryFinancialInstrumentIndicator4Choice)
	return s.TemporaryFinancialInstrumentIndicator
}

func (s *SecuritiesOption57) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator4Choice {
	s.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator4Choice)
	return s.NonEligibleProceedsIndicator
}

func (s *SecuritiesOption57) AddIssuerOfferorTaxabilityIndicator() *GenericIdentification47 {
	s.IssuerOfferorTaxabilityIndicator = new(GenericIdentification47)
	return s.IssuerOfferorTaxabilityIndicator
}

func (s *SecuritiesOption57) SetNewSecuritiesIssuanceIndicator(value string) {
	s.NewSecuritiesIssuanceIndicator = (*NewSecuritiesIssuanceType5Code)(&value)
}

func (s *SecuritiesOption57) AddIncomeType() *GenericIdentification47 {
	s.IncomeType = new(GenericIdentification47)
	return s.IncomeType
}

func (s *SecuritiesOption57) AddExemptionType() *GenericIdentification47 {
	newValue := new(GenericIdentification47)
	s.ExemptionType = append(s.ExemptionType, newValue)
	return newValue
}

func (s *SecuritiesOption57) AddEntitledQuantity() *Quantity10Choice {
	s.EntitledQuantity = new(Quantity10Choice)
	return s.EntitledQuantity
}

func (s *SecuritiesOption57) AddSafekeepingPlace() *SafekeepingPlaceFormat17Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat17Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesOption57) SetCountryOfIncomeSource(value string) {
	s.CountryOfIncomeSource = (*CountryCode)(&value)
}

func (s *SecuritiesOption57) AddFractionDisposition() *FractionDispositionType31Choice {
	s.FractionDisposition = new(FractionDispositionType31Choice)
	return s.FractionDisposition
}

func (s *SecuritiesOption57) SetCurrencyOption(value string) {
	s.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (s *SecuritiesOption57) AddTradingPeriod() *Period3Choice {
	s.TradingPeriod = new(Period3Choice)
	return s.TradingPeriod
}

func (s *SecuritiesOption57) AddDateDetails() *SecurityDate14 {
	s.DateDetails = new(SecurityDate14)
	return s.DateDetails
}

func (s *SecuritiesOption57) AddRateDetails() *CorporateActionRate77 {
	s.RateDetails = new(CorporateActionRate77)
	return s.RateDetails
}

func (s *SecuritiesOption57) AddPriceDetails() *CorporateActionPrice66 {
	s.PriceDetails = new(CorporateActionPrice66)
	return s.PriceDetails
}

// Provides information about the corporate action security option.
type SecuritiesOption58 struct {

	// Identification of the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Quantity of securities that have been posted (credit or debit) to the safekeeping account.
	PostingQuantity *Quantity10Choice `xml:"PstngQty"`

	// Date of the posting (credit or debit) to the account.
	PostingDate *DateAndDateTimeChoice `xml:"PstngDt"`

	// Date/Time of the posting (credit or debit) to the account that was initially communicated in the confirmation.
	OriginalPostingDate *DateAndDateTimeChoice `xml:"OrgnlPstngDt,omitempty"`
}

func (s *SecuritiesOption58) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesOption58) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesOption58) AddPostingQuantity() *Quantity10Choice {
	s.PostingQuantity = new(Quantity10Choice)
	return s.PostingQuantity
}

func (s *SecuritiesOption58) AddPostingDate() *DateAndDateTimeChoice {
	s.PostingDate = new(DateAndDateTimeChoice)
	return s.PostingDate
}

func (s *SecuritiesOption58) AddOriginalPostingDate() *DateAndDateTimeChoice {
	s.OriginalPostingDate = new(DateAndDateTimeChoice)
	return s.OriginalPostingDate
}

// Provides information about the corporate action security option.
type SecuritiesOption59 struct {

	// Provides description of the financial instrument related to securities movement.
	SecurityDetails *FinancialInstrumentAttributes80 `xml:"SctyDtls"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies that the security identified  is a temporary security identification used for processing reasons, for example, contra security used in the US.
	TemporaryFinancialInstrumentIndicator *TemporaryFinancialInstrumentIndicator3Choice `xml:"TempFinInstrmInd,omitempty"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator3Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *IssuerOfferorTaxabilityIndicator1Choice `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Indicates whether the securities are newly issued or not.
	NewSecuritiesIssuanceIndicator *NewSecuritiesIssuanceType5Code `xml:"NewSctiesIssncInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification30 `xml:"IncmTp,omitempty"`

	// Specifies the basis for the reduced rate of withholding.
	ExemptionType []*GenericIdentification30 `xml:"XmptnTp,omitempty"`

	// Quantity of securities based on the terms of the corporate action event and balance of underlying securities entitled to the account owner. (This quantity can be positive or negative).
	EntitledQuantity *Quantity6Choice `xml:"EntitldQty,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat10Choice `xml:"SfkpgPlc,omitempty"`

	// Indicates the country from which the income originates.
	CountryOfIncomeSource *CountryCode `xml:"CtryOfIncmSrc,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType26Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Period during which intermediate or outturn securities are tradable in a secondary market.
	TradingPeriod *Period3Choice `xml:"TradgPrd,omitempty"`

	// Provides information about the dates related to securities movement.
	DateDetails *SecurityDate12 `xml:"DtDtls"`

	// Provides information about the rates related to securities movement.
	RateDetails *CorporateActionRate69 `xml:"RateDtls,omitempty"`

	// Provides information about the prices related to securities movement.
	PriceDetails *CorporateActionPrice56 `xml:"PricDtls,omitempty"`
}

func (s *SecuritiesOption59) AddSecurityDetails() *FinancialInstrumentAttributes80 {
	s.SecurityDetails = new(FinancialInstrumentAttributes80)
	return s.SecurityDetails
}

func (s *SecuritiesOption59) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesOption59) AddTemporaryFinancialInstrumentIndicator() *TemporaryFinancialInstrumentIndicator3Choice {
	s.TemporaryFinancialInstrumentIndicator = new(TemporaryFinancialInstrumentIndicator3Choice)
	return s.TemporaryFinancialInstrumentIndicator
}

func (s *SecuritiesOption59) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator3Choice {
	s.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator3Choice)
	return s.NonEligibleProceedsIndicator
}

func (s *SecuritiesOption59) AddIssuerOfferorTaxabilityIndicator() *IssuerOfferorTaxabilityIndicator1Choice {
	s.IssuerOfferorTaxabilityIndicator = new(IssuerOfferorTaxabilityIndicator1Choice)
	return s.IssuerOfferorTaxabilityIndicator
}

func (s *SecuritiesOption59) SetNewSecuritiesIssuanceIndicator(value string) {
	s.NewSecuritiesIssuanceIndicator = (*NewSecuritiesIssuanceType5Code)(&value)
}

func (s *SecuritiesOption59) AddIncomeType() *GenericIdentification30 {
	s.IncomeType = new(GenericIdentification30)
	return s.IncomeType
}

func (s *SecuritiesOption59) AddExemptionType() *GenericIdentification30 {
	newValue := new(GenericIdentification30)
	s.ExemptionType = append(s.ExemptionType, newValue)
	return newValue
}

func (s *SecuritiesOption59) AddEntitledQuantity() *Quantity6Choice {
	s.EntitledQuantity = new(Quantity6Choice)
	return s.EntitledQuantity
}

func (s *SecuritiesOption59) AddSafekeepingPlace() *SafekeepingPlaceFormat10Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat10Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesOption59) SetCountryOfIncomeSource(value string) {
	s.CountryOfIncomeSource = (*CountryCode)(&value)
}

func (s *SecuritiesOption59) AddFractionDisposition() *FractionDispositionType26Choice {
	s.FractionDisposition = new(FractionDispositionType26Choice)
	return s.FractionDisposition
}

func (s *SecuritiesOption59) SetCurrencyOption(value string) {
	s.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (s *SecuritiesOption59) AddTradingPeriod() *Period3Choice {
	s.TradingPeriod = new(Period3Choice)
	return s.TradingPeriod
}

func (s *SecuritiesOption59) AddDateDetails() *SecurityDate12 {
	s.DateDetails = new(SecurityDate12)
	return s.DateDetails
}

func (s *SecuritiesOption59) AddRateDetails() *CorporateActionRate69 {
	s.RateDetails = new(CorporateActionRate69)
	return s.RateDetails
}

func (s *SecuritiesOption59) AddPriceDetails() *CorporateActionPrice56 {
	s.PriceDetails = new(CorporateActionPrice56)
	return s.PriceDetails
}

// Provides information about the corporate action security option.
type SecuritiesOption6 struct {

	// Provides description of the financial instrument related to securities movement.
	SecurityDetails *FinancialInstrumentAttributes5 `xml:"SctyDtls"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies that the security identified  is a temporary security identification used for processing reasons, for example, contra security used in the US.
	TemporaryFinancialInstrumentIndicator *TemporaryFinancialInstrumentIndicator1Choice `xml:"TempFinInstrmInd,omitempty"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator1Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Quantity of securities based on the terms of the corporate action event and balance of underlying securities entitled to the account owner. (This quantity can be positive or negative).
	EntitledQuantity *Quantity6Choice `xml:"EntitldQty,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType1Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Period during which intermediate or outturn securities are tradable in a secondary market.
	TradingPeriod *Period1Choice `xml:"TradgPrd,omitempty"`

	// Provides information about the dates related to securities movement.
	DateDetails *SecurityDate2 `xml:"DtDtls"`

	// Provides information about the rates related to securities movement.
	RateDetails *CorporateActionRate7 `xml:"RateDtls,omitempty"`

	// Provides information about the prices related to securities movement.
	PriceDetails *CorporateActionPrice10 `xml:"PricDtls,omitempty"`
}

func (s *SecuritiesOption6) AddSecurityDetails() *FinancialInstrumentAttributes5 {
	s.SecurityDetails = new(FinancialInstrumentAttributes5)
	return s.SecurityDetails
}

func (s *SecuritiesOption6) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesOption6) AddTemporaryFinancialInstrumentIndicator() *TemporaryFinancialInstrumentIndicator1Choice {
	s.TemporaryFinancialInstrumentIndicator = new(TemporaryFinancialInstrumentIndicator1Choice)
	return s.TemporaryFinancialInstrumentIndicator
}

func (s *SecuritiesOption6) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator1Choice {
	s.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator1Choice)
	return s.NonEligibleProceedsIndicator
}

func (s *SecuritiesOption6) AddEntitledQuantity() *Quantity6Choice {
	s.EntitledQuantity = new(Quantity6Choice)
	return s.EntitledQuantity
}

func (s *SecuritiesOption6) AddFractionDisposition() *FractionDispositionType1Choice {
	s.FractionDisposition = new(FractionDispositionType1Choice)
	return s.FractionDisposition
}

func (s *SecuritiesOption6) SetCurrencyOption(value string) {
	s.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (s *SecuritiesOption6) AddTradingPeriod() *Period1Choice {
	s.TradingPeriod = new(Period1Choice)
	return s.TradingPeriod
}

func (s *SecuritiesOption6) AddDateDetails() *SecurityDate2 {
	s.DateDetails = new(SecurityDate2)
	return s.DateDetails
}

func (s *SecuritiesOption6) AddRateDetails() *CorporateActionRate7 {
	s.RateDetails = new(CorporateActionRate7)
	return s.RateDetails
}

func (s *SecuritiesOption6) AddPriceDetails() *CorporateActionPrice10 {
	s.PriceDetails = new(CorporateActionPrice10)
	return s.PriceDetails
}

// Provides information about the corporate action security option.
type SecuritiesOption60 struct {

	// Identification of the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies that the security identified  is a temporary security identification used for processing reasons, for example, contra security used in the US.
	TemporaryFinancialInstrumentIndicator *TemporaryFinancialInstrumentIndicator3Choice `xml:"TempFinInstrmInd,omitempty"`

	// Indicates whether the securities are newly issued or not.
	NewSecuritiesIssuanceIndicator *NewSecuritiesIssuanceType6Code `xml:"NewSctiesIssncInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *IssuerOfferorTaxabilityIndicator1Choice `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification30 `xml:"IncmTp,omitempty"`

	// Specifies the basis for the reduced rate of withholding.
	ExemptionType []*GenericIdentification30 `xml:"XmptnTp,omitempty"`

	// Indicates the country from which the income originates.
	CountryOfIncomeSource *CountryCode `xml:"CtryOfIncmSrc,omitempty"`

	// Quantity of securities that have been posted (credit or debit) to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat10Choice `xml:"SfkpgPlc,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType27Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Provides information about the dates related to securities movement.
	DateDetails *SecurityDate11 `xml:"DtDtls"`

	// Provides information about the rates related to securities movement.
	RateDetails *CorporateActionRate72 `xml:"RateDtls,omitempty"`

	// Provides information about the prices related to securities movement.
	PriceDetails *CorporateActionPrice59 `xml:"PricDtls,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties42 `xml:"RcvgSttlmPties,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties42 `xml:"DlvrgSttlmPties,omitempty"`
}

func (s *SecuritiesOption60) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesOption60) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesOption60) AddTemporaryFinancialInstrumentIndicator() *TemporaryFinancialInstrumentIndicator3Choice {
	s.TemporaryFinancialInstrumentIndicator = new(TemporaryFinancialInstrumentIndicator3Choice)
	return s.TemporaryFinancialInstrumentIndicator
}

func (s *SecuritiesOption60) SetNewSecuritiesIssuanceIndicator(value string) {
	s.NewSecuritiesIssuanceIndicator = (*NewSecuritiesIssuanceType6Code)(&value)
}

func (s *SecuritiesOption60) AddIssuerOfferorTaxabilityIndicator() *IssuerOfferorTaxabilityIndicator1Choice {
	s.IssuerOfferorTaxabilityIndicator = new(IssuerOfferorTaxabilityIndicator1Choice)
	return s.IssuerOfferorTaxabilityIndicator
}

func (s *SecuritiesOption60) AddIncomeType() *GenericIdentification30 {
	s.IncomeType = new(GenericIdentification30)
	return s.IncomeType
}

func (s *SecuritiesOption60) AddExemptionType() *GenericIdentification30 {
	newValue := new(GenericIdentification30)
	s.ExemptionType = append(s.ExemptionType, newValue)
	return newValue
}

func (s *SecuritiesOption60) SetCountryOfIncomeSource(value string) {
	s.CountryOfIncomeSource = (*CountryCode)(&value)
}

func (s *SecuritiesOption60) AddPostingQuantity() *Quantity6Choice {
	s.PostingQuantity = new(Quantity6Choice)
	return s.PostingQuantity
}

func (s *SecuritiesOption60) AddSafekeepingPlace() *SafekeepingPlaceFormat10Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat10Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesOption60) AddFractionDisposition() *FractionDispositionType27Choice {
	s.FractionDisposition = new(FractionDispositionType27Choice)
	return s.FractionDisposition
}

func (s *SecuritiesOption60) SetCurrencyOption(value string) {
	s.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (s *SecuritiesOption60) AddDateDetails() *SecurityDate11 {
	s.DateDetails = new(SecurityDate11)
	return s.DateDetails
}

func (s *SecuritiesOption60) AddRateDetails() *CorporateActionRate72 {
	s.RateDetails = new(CorporateActionRate72)
	return s.RateDetails
}

func (s *SecuritiesOption60) AddPriceDetails() *CorporateActionPrice59 {
	s.PriceDetails = new(CorporateActionPrice59)
	return s.PriceDetails
}

func (s *SecuritiesOption60) AddReceivingSettlementParties() *SettlementParties42 {
	s.ReceivingSettlementParties = new(SettlementParties42)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesOption60) AddDeliveringSettlementParties() *SettlementParties42 {
	s.DeliveringSettlementParties = new(SettlementParties42)
	return s.DeliveringSettlementParties
}

// Provides information about the corporate action security option.
type SecuritiesOption61 struct {

	// Provides description of the financial instrument related to securities movement.
	SecurityDetails *FinancialInstrumentAttributes83 `xml:"SctyDtls"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies that the security identified  is a temporary security identification used for processing reasons, for example, contra security used in the US.
	TemporaryFinancialInstrumentIndicator *TemporaryFinancialInstrumentIndicator4Choice `xml:"TempFinInstrmInd,omitempty"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator4Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *IssuerOfferorTaxabilityIndicator1Choice `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Indicates whether the securities are newly issued or not.
	NewSecuritiesIssuanceIndicator *NewSecuritiesIssuanceType5Code `xml:"NewSctiesIssncInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification47 `xml:"IncmTp,omitempty"`

	// Specifies the basis for the reduced rate of withholding.
	ExemptionType []*GenericIdentification47 `xml:"XmptnTp,omitempty"`

	// Quantity of securities based on the terms of the corporate action event and balance of underlying securities entitled to the account owner. (This quantity can be positive or negative).
	EntitledQuantity *Quantity10Choice `xml:"EntitldQty,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat17Choice `xml:"SfkpgPlc,omitempty"`

	// Indicates the country from which the income originates.
	CountryOfIncomeSource *CountryCode `xml:"CtryOfIncmSrc,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType31Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Period during which intermediate or outturn securities are tradable in a secondary market.
	TradingPeriod *Period3Choice `xml:"TradgPrd,omitempty"`

	// Provides information about the dates related to securities movement.
	DateDetails *SecurityDate14 `xml:"DtDtls"`

	// Provides information about the rates related to securities movement.
	RateDetails *CorporateActionRate77 `xml:"RateDtls,omitempty"`

	// Provides information about the prices related to securities movement.
	PriceDetails *CorporateActionPrice66 `xml:"PricDtls,omitempty"`
}

func (s *SecuritiesOption61) AddSecurityDetails() *FinancialInstrumentAttributes83 {
	s.SecurityDetails = new(FinancialInstrumentAttributes83)
	return s.SecurityDetails
}

func (s *SecuritiesOption61) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesOption61) AddTemporaryFinancialInstrumentIndicator() *TemporaryFinancialInstrumentIndicator4Choice {
	s.TemporaryFinancialInstrumentIndicator = new(TemporaryFinancialInstrumentIndicator4Choice)
	return s.TemporaryFinancialInstrumentIndicator
}

func (s *SecuritiesOption61) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator4Choice {
	s.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator4Choice)
	return s.NonEligibleProceedsIndicator
}

func (s *SecuritiesOption61) AddIssuerOfferorTaxabilityIndicator() *IssuerOfferorTaxabilityIndicator1Choice {
	s.IssuerOfferorTaxabilityIndicator = new(IssuerOfferorTaxabilityIndicator1Choice)
	return s.IssuerOfferorTaxabilityIndicator
}

func (s *SecuritiesOption61) SetNewSecuritiesIssuanceIndicator(value string) {
	s.NewSecuritiesIssuanceIndicator = (*NewSecuritiesIssuanceType5Code)(&value)
}

func (s *SecuritiesOption61) AddIncomeType() *GenericIdentification47 {
	s.IncomeType = new(GenericIdentification47)
	return s.IncomeType
}

func (s *SecuritiesOption61) AddExemptionType() *GenericIdentification47 {
	newValue := new(GenericIdentification47)
	s.ExemptionType = append(s.ExemptionType, newValue)
	return newValue
}

func (s *SecuritiesOption61) AddEntitledQuantity() *Quantity10Choice {
	s.EntitledQuantity = new(Quantity10Choice)
	return s.EntitledQuantity
}

func (s *SecuritiesOption61) AddSafekeepingPlace() *SafekeepingPlaceFormat17Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat17Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesOption61) SetCountryOfIncomeSource(value string) {
	s.CountryOfIncomeSource = (*CountryCode)(&value)
}

func (s *SecuritiesOption61) AddFractionDisposition() *FractionDispositionType31Choice {
	s.FractionDisposition = new(FractionDispositionType31Choice)
	return s.FractionDisposition
}

func (s *SecuritiesOption61) SetCurrencyOption(value string) {
	s.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (s *SecuritiesOption61) AddTradingPeriod() *Period3Choice {
	s.TradingPeriod = new(Period3Choice)
	return s.TradingPeriod
}

func (s *SecuritiesOption61) AddDateDetails() *SecurityDate14 {
	s.DateDetails = new(SecurityDate14)
	return s.DateDetails
}

func (s *SecuritiesOption61) AddRateDetails() *CorporateActionRate77 {
	s.RateDetails = new(CorporateActionRate77)
	return s.RateDetails
}

func (s *SecuritiesOption61) AddPriceDetails() *CorporateActionPrice66 {
	s.PriceDetails = new(CorporateActionPrice66)
	return s.PriceDetails
}

// Provides information about the corporate action security option.
type SecuritiesOption63 struct {

	// Identification of the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies that the security identified  is a temporary security identification used for processing reasons, for example, contra security used in the US.
	TemporaryFinancialInstrumentIndicator *TemporaryFinancialInstrumentIndicator4Choice `xml:"TempFinInstrmInd,omitempty"`

	// Indicates whether the securities are newly issued or not.
	NewSecuritiesIssuanceIndicator *NewSecuritiesIssuanceType6Code `xml:"NewSctiesIssncInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *IssuerOfferorTaxabilityIndicator1Choice `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification47 `xml:"IncmTp,omitempty"`

	// Specifies the basis for the reduced rate of withholding.
	ExemptionType []*GenericIdentification47 `xml:"XmptnTp,omitempty"`

	// Indicates the country from which the income originates.
	CountryOfIncomeSource *CountryCode `xml:"CtryOfIncmSrc,omitempty"`

	// Quantity of securities that have been posted (credit or debit) to the safekeeping account.
	PostingQuantity *Quantity10Choice `xml:"PstngQty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat17Choice `xml:"SfkpgPlc,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType30Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Provides information about the dates related to securities movement.
	DateDetails *SecurityDate13 `xml:"DtDtls"`

	// Provides information about the rates related to securities movement.
	RateDetails *CorporateActionRate75 `xml:"RateDtls,omitempty"`

	// Provides information about the prices related to securities movement.
	PriceDetails *CorporateActionPrice64 `xml:"PricDtls,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties61 `xml:"RcvgSttlmPties,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties61 `xml:"DlvrgSttlmPties,omitempty"`
}

func (s *SecuritiesOption63) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesOption63) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesOption63) AddTemporaryFinancialInstrumentIndicator() *TemporaryFinancialInstrumentIndicator4Choice {
	s.TemporaryFinancialInstrumentIndicator = new(TemporaryFinancialInstrumentIndicator4Choice)
	return s.TemporaryFinancialInstrumentIndicator
}

func (s *SecuritiesOption63) SetNewSecuritiesIssuanceIndicator(value string) {
	s.NewSecuritiesIssuanceIndicator = (*NewSecuritiesIssuanceType6Code)(&value)
}

func (s *SecuritiesOption63) AddIssuerOfferorTaxabilityIndicator() *IssuerOfferorTaxabilityIndicator1Choice {
	s.IssuerOfferorTaxabilityIndicator = new(IssuerOfferorTaxabilityIndicator1Choice)
	return s.IssuerOfferorTaxabilityIndicator
}

func (s *SecuritiesOption63) AddIncomeType() *GenericIdentification47 {
	s.IncomeType = new(GenericIdentification47)
	return s.IncomeType
}

func (s *SecuritiesOption63) AddExemptionType() *GenericIdentification47 {
	newValue := new(GenericIdentification47)
	s.ExemptionType = append(s.ExemptionType, newValue)
	return newValue
}

func (s *SecuritiesOption63) SetCountryOfIncomeSource(value string) {
	s.CountryOfIncomeSource = (*CountryCode)(&value)
}

func (s *SecuritiesOption63) AddPostingQuantity() *Quantity10Choice {
	s.PostingQuantity = new(Quantity10Choice)
	return s.PostingQuantity
}

func (s *SecuritiesOption63) AddSafekeepingPlace() *SafekeepingPlaceFormat17Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat17Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesOption63) AddFractionDisposition() *FractionDispositionType30Choice {
	s.FractionDisposition = new(FractionDispositionType30Choice)
	return s.FractionDisposition
}

func (s *SecuritiesOption63) SetCurrencyOption(value string) {
	s.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (s *SecuritiesOption63) AddDateDetails() *SecurityDate13 {
	s.DateDetails = new(SecurityDate13)
	return s.DateDetails
}

func (s *SecuritiesOption63) AddRateDetails() *CorporateActionRate75 {
	s.RateDetails = new(CorporateActionRate75)
	return s.RateDetails
}

func (s *SecuritiesOption63) AddPriceDetails() *CorporateActionPrice64 {
	s.PriceDetails = new(CorporateActionPrice64)
	return s.PriceDetails
}

func (s *SecuritiesOption63) AddReceivingSettlementParties() *SettlementParties61 {
	s.ReceivingSettlementParties = new(SettlementParties61)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesOption63) AddDeliveringSettlementParties() *SettlementParties61 {
	s.DeliveringSettlementParties = new(SettlementParties61)
	return s.DeliveringSettlementParties
}
