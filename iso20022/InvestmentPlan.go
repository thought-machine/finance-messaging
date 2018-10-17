package iso20022

// Plan that allows investors to schedule periodical investments or divestments, according to pre-defined criteria.
type InvestmentPlan10 struct {

	// Frequency of the investment or divestment.
	Frequency *Frequency20Choice `xml:"Frqcy"`

	// Date the investment plan starts.
	StartDate *ISODate `xml:"StartDt,omitempty"`

	// Date the investment plan stops.
	EndDate *ISODate `xml:"EndDt,omitempty"`

	// Amount of the periodical payments.
	Quantity *UnitsOrAmount1Choice `xml:"Qty"`

	// Indicates whether an ordered amount is a gross amount (including all charges, commissions, tax). If it is not a gross amount, the ordered amount is a net amount (amount to be invested or redeemed from the fund to which other elements will be added).
	GrossAmountIndicator *YesNoIndicator `xml:"GrssAmtInd,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Number of pre-paid instalment periods at the time the investment plan is created.
	InitialNumberOfInstalments *Number `xml:"InitlNbOfInstlmts,omitempty"`

	// Total number of times the amount must be invested at the predefined frequency as of the start date of the investment plan.
	TotalNumberOfInstalments *Number `xml:"TtlNbOfInstlmts,omitempty"`

	// Indicates the rounding direction when an amount is to be spread over several funds.
	RoundingDirection *RoundingDirection1Code `xml:"RndgDrctn,omitempty"`

	// Security that an investment plan invests in, or from which the investment plan divests.
	SecurityDetails []*Repartition3 `xml:"SctyDtls"`

	// Cash settlement standing instruction associated to the investment plan.
	CashSettlement []*InvestmentFundCashSettlementInformation7 `xml:"CshSttlm,omitempty"`

	// Reference of the underlying investment contract. In some markets, such as Italy, this might be required to segregate holdings between the same investment account.
	ContractReference *Max35Text `xml:"CtrctRef,omitempty"`

	// Reference of the previous contract to which this savings or withdrawal plan is related.
	RelatedContractReference *Max35Text `xml:"RltdCtrctRef,omitempty"`

	// Identification of the product as designated by the fund manager. In some markets, such as Italy, the financial product or service related to a savings plan or withdrawal plan are identified by a product identification or number.
	ProductIdentification *Max35Text `xml:"PdctId,omitempty"`

	// Reference of the underlying service level agreement (SLA) governing charges and commission.
	SLAChargeAndCommissionReference *Max35Text `xml:"SLAChrgAndComssnRef,omitempty"`

	// Specifies the type of insurance contract to which the savings investment plan is linked.
	InsuranceCover *InsuranceType1Choice `xml:"InsrncCover,omitempty"`

	// Status of the savings or withdrawal investment plan.
	PlanStatus *PlanStatus1Choice `xml:"PlanSts,omitempty"`

	// Role or function of the instalment manager.
	InstalmentManagerRole *PartyRole4Choice `xml:"InstlmtMgrRole,omitempty"`
}

func (i *InvestmentPlan10) AddFrequency() *Frequency20Choice {
	i.Frequency = new(Frequency20Choice)
	return i.Frequency
}

func (i *InvestmentPlan10) SetStartDate(value string) {
	i.StartDate = (*ISODate)(&value)
}

func (i *InvestmentPlan10) SetEndDate(value string) {
	i.EndDate = (*ISODate)(&value)
}

func (i *InvestmentPlan10) AddQuantity() *UnitsOrAmount1Choice {
	i.Quantity = new(UnitsOrAmount1Choice)
	return i.Quantity
}

func (i *InvestmentPlan10) SetGrossAmountIndicator(value string) {
	i.GrossAmountIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentPlan10) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference1Code)(&value)
}

func (i *InvestmentPlan10) SetInitialNumberOfInstalments(value string) {
	i.InitialNumberOfInstalments = (*Number)(&value)
}

func (i *InvestmentPlan10) SetTotalNumberOfInstalments(value string) {
	i.TotalNumberOfInstalments = (*Number)(&value)
}

func (i *InvestmentPlan10) SetRoundingDirection(value string) {
	i.RoundingDirection = (*RoundingDirection1Code)(&value)
}

func (i *InvestmentPlan10) AddSecurityDetails() *Repartition3 {
	newValue := new(Repartition3)
	i.SecurityDetails = append(i.SecurityDetails, newValue)
	return newValue
}

func (i *InvestmentPlan10) AddCashSettlement() *InvestmentFundCashSettlementInformation7 {
	newValue := new(InvestmentFundCashSettlementInformation7)
	i.CashSettlement = append(i.CashSettlement, newValue)
	return newValue
}

func (i *InvestmentPlan10) SetContractReference(value string) {
	i.ContractReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan10) SetRelatedContractReference(value string) {
	i.RelatedContractReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan10) SetProductIdentification(value string) {
	i.ProductIdentification = (*Max35Text)(&value)
}

func (i *InvestmentPlan10) SetSLAChargeAndCommissionReference(value string) {
	i.SLAChargeAndCommissionReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan10) AddInsuranceCover() *InsuranceType1Choice {
	i.InsuranceCover = new(InsuranceType1Choice)
	return i.InsuranceCover
}

func (i *InvestmentPlan10) AddPlanStatus() *PlanStatus1Choice {
	i.PlanStatus = new(PlanStatus1Choice)
	return i.PlanStatus
}

func (i *InvestmentPlan10) AddInstalmentManagerRole() *PartyRole4Choice {
	i.InstalmentManagerRole = new(PartyRole4Choice)
	return i.InstalmentManagerRole
}

// Plan that allows investors to schedule periodical investments or divestments, according to pre-defined criteria.
type InvestmentPlan11 struct {

	// Frequency of the investment or divestment.
	Frequency *Frequency20Choice `xml:"Frqcy"`

	// Date the investment plan starts.
	StartDate *ISODate `xml:"StartDt,omitempty"`

	// Date the investment plan stops.
	EndDate *ISODate `xml:"EndDt,omitempty"`

	// Amount of the periodical payments.
	Quantity *UnitsOrAmount1Choice `xml:"Qty"`

	// Indicates whether an ordered amount is a gross amount (including all charges, commissions, tax). If it is not a gross amount, the ordered amount is a net amount (amount to be invested or redeemed from the fund to which other elements will be added).
	GrossAmountIndicator *YesNoIndicator `xml:"GrssAmtInd,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Number of pre-paid instalment periods at the time the investment plan is created.
	InitialNumberOfInstalments *Number `xml:"InitlNbOfInstlmts,omitempty"`

	// Total number of times the amount must be invested at the predefined frequency as of the start date of the investment plan.
	TotalNumberOfInstalments *Number `xml:"TtlNbOfInstlmts,omitempty"`

	// Indicates the rounding direction when an amount is to be spread over several funds.
	RoundingDirection *RoundingDirection1Code `xml:"RndgDrctn,omitempty"`

	// Security that an investment plan invests in, or from which the investment plan divests.
	SecurityDetails []*Repartition3 `xml:"SctyDtls"`

	// Cash settlement standing instruction associated to the investment plan and to be either inserted or deleted.
	ModifiedCashSettlement []*InvestmentFundCashSettlementInformation8 `xml:"ModfdCshSttlm,omitempty"`

	// Reference of the underlying investment contract. In some markets, such as Italy, this might be required to segregate holdings between the same investment account.
	ContractReference *Max35Text `xml:"CtrctRef,omitempty"`

	// Reference of the previous contract to which this savings or withdrawal plan is related.
	RelatedContractReference *Max35Text `xml:"RltdCtrctRef,omitempty"`

	// Identification of the product as designated by the fund manager. In some markets, such as Italy, the financial product or service related to a savings plan or withdrawal plan are identified by a product identification or number.
	ProductIdentification *Max35Text `xml:"PdctId,omitempty"`

	// Reference of the underlying service level agreement (SLA) governing charges and commission.
	SLAChargeAndCommissionReference *Max35Text `xml:"SLAChrgAndComssnRef,omitempty"`

	// Specifies the type of insurance contract to which the savings investment plan is linked.
	InsuranceCover *InsuranceType1Choice `xml:"InsrncCover,omitempty"`

	// Status of the savings or withdrawal investment plan.
	PlanStatus *PlanStatus1Choice `xml:"PlanSts,omitempty"`

	// Role or function of the instalment manager.
	InstalmentManagerRole *PartyRole4Choice `xml:"InstlmtMgrRole,omitempty"`
}

func (i *InvestmentPlan11) AddFrequency() *Frequency20Choice {
	i.Frequency = new(Frequency20Choice)
	return i.Frequency
}

func (i *InvestmentPlan11) SetStartDate(value string) {
	i.StartDate = (*ISODate)(&value)
}

func (i *InvestmentPlan11) SetEndDate(value string) {
	i.EndDate = (*ISODate)(&value)
}

func (i *InvestmentPlan11) AddQuantity() *UnitsOrAmount1Choice {
	i.Quantity = new(UnitsOrAmount1Choice)
	return i.Quantity
}

func (i *InvestmentPlan11) SetGrossAmountIndicator(value string) {
	i.GrossAmountIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentPlan11) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference1Code)(&value)
}

func (i *InvestmentPlan11) SetInitialNumberOfInstalments(value string) {
	i.InitialNumberOfInstalments = (*Number)(&value)
}

func (i *InvestmentPlan11) SetTotalNumberOfInstalments(value string) {
	i.TotalNumberOfInstalments = (*Number)(&value)
}

func (i *InvestmentPlan11) SetRoundingDirection(value string) {
	i.RoundingDirection = (*RoundingDirection1Code)(&value)
}

func (i *InvestmentPlan11) AddSecurityDetails() *Repartition3 {
	newValue := new(Repartition3)
	i.SecurityDetails = append(i.SecurityDetails, newValue)
	return newValue
}

func (i *InvestmentPlan11) AddModifiedCashSettlement() *InvestmentFundCashSettlementInformation8 {
	newValue := new(InvestmentFundCashSettlementInformation8)
	i.ModifiedCashSettlement = append(i.ModifiedCashSettlement, newValue)
	return newValue
}

func (i *InvestmentPlan11) SetContractReference(value string) {
	i.ContractReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan11) SetRelatedContractReference(value string) {
	i.RelatedContractReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan11) SetProductIdentification(value string) {
	i.ProductIdentification = (*Max35Text)(&value)
}

func (i *InvestmentPlan11) SetSLAChargeAndCommissionReference(value string) {
	i.SLAChargeAndCommissionReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan11) AddInsuranceCover() *InsuranceType1Choice {
	i.InsuranceCover = new(InsuranceType1Choice)
	return i.InsuranceCover
}

func (i *InvestmentPlan11) AddPlanStatus() *PlanStatus1Choice {
	i.PlanStatus = new(PlanStatus1Choice)
	return i.PlanStatus
}

func (i *InvestmentPlan11) AddInstalmentManagerRole() *PartyRole4Choice {
	i.InstalmentManagerRole = new(PartyRole4Choice)
	return i.InstalmentManagerRole
}

// Plan that allows investors to schedule periodical investments or divestments, according to pre-defined criteria.
type InvestmentPlan12 struct {

	// Frequency of the investment or divestment.
	Frequency *Frequency20Choice `xml:"Frqcy"`

	// Date the investment plan starts.
	StartDate *ISODate `xml:"StartDt,omitempty"`

	// Date the investment plan stops.
	EndDate *ISODate `xml:"EndDt,omitempty"`

	// Amount of the periodical payments.
	Quantity *UnitsOrAmount1Choice `xml:"Qty"`

	// Indicates whether an ordered amount is a gross amount (including all charges, commissions, tax). If it is not a gross amount, the ordered amount is a net amount (amount to be invested or redeemed from the fund to which other elements will be added).
	GrossAmountIndicator *YesNoIndicator `xml:"GrssAmtInd,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference2Code `xml:"IncmPref,omitempty"`

	// Initial amount or number of initial instalments.
	InitialAmount *InitialAmount1Choice `xml:"InitlAmt,omitempty"`

	// Total number of times the amount must be invested at the predefined frequency as of the start date of the investment plan.
	TotalNumberOfInstalments *Number `xml:"TtlNbOfInstlmts,omitempty"`

	// Indicates the rounding direction when an amount is to be spread over several funds.
	RoundingDirection *RoundingDirection1Code `xml:"RndgDrctn,omitempty"`

	// Security that an investment plan invests in, or from which the investment plan divests.
	SecurityDetails []*Repartition4 `xml:"SctyDtls"`

	// Cash settlement standing instruction associated to the investment plan.
	CashSettlement []*CashSettlement1 `xml:"CshSttlm,omitempty"`

	// Reference of the underlying investment contract. In some markets, such as Italy, this might be required to segregate holdings between the same investment account.
	ContractReference *Max35Text `xml:"CtrctRef,omitempty"`

	// Reference of the previous contract to which this savings or withdrawal plan is related.
	RelatedContractReference *Max35Text `xml:"RltdCtrctRef,omitempty"`

	// Identification of the product as designated by the fund manager. In some markets, such as Italy, the financial product or service related to a savings plan or withdrawal plan are identified by a product identification or number.
	ProductIdentification *Max35Text `xml:"PdctId,omitempty"`

	// Reference of the underlying service level agreement (SLA) governing charges and commission.
	SLAChargeAndCommissionReference *Max35Text `xml:"SLAChrgAndComssnRef,omitempty"`

	// Specifies the type of insurance contract to which the savings investment plan is linked.
	InsuranceCover *InsuranceType2Choice `xml:"InsrncCover,omitempty"`

	// Status of the savings or withdrawal investment plan.
	PlanStatus *PlanStatus2Choice `xml:"PlanSts,omitempty"`

	// Role or function of the instalment manager.
	InstalmentManagerRole *PartyRole4Choice `xml:"InstlmtMgrRole,omitempty"`
}

func (i *InvestmentPlan12) AddFrequency() *Frequency20Choice {
	i.Frequency = new(Frequency20Choice)
	return i.Frequency
}

func (i *InvestmentPlan12) SetStartDate(value string) {
	i.StartDate = (*ISODate)(&value)
}

func (i *InvestmentPlan12) SetEndDate(value string) {
	i.EndDate = (*ISODate)(&value)
}

func (i *InvestmentPlan12) AddQuantity() *UnitsOrAmount1Choice {
	i.Quantity = new(UnitsOrAmount1Choice)
	return i.Quantity
}

func (i *InvestmentPlan12) SetGrossAmountIndicator(value string) {
	i.GrossAmountIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentPlan12) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference2Code)(&value)
}

func (i *InvestmentPlan12) AddInitialAmount() *InitialAmount1Choice {
	i.InitialAmount = new(InitialAmount1Choice)
	return i.InitialAmount
}

func (i *InvestmentPlan12) SetTotalNumberOfInstalments(value string) {
	i.TotalNumberOfInstalments = (*Number)(&value)
}

func (i *InvestmentPlan12) SetRoundingDirection(value string) {
	i.RoundingDirection = (*RoundingDirection1Code)(&value)
}

func (i *InvestmentPlan12) AddSecurityDetails() *Repartition4 {
	newValue := new(Repartition4)
	i.SecurityDetails = append(i.SecurityDetails, newValue)
	return newValue
}

func (i *InvestmentPlan12) AddCashSettlement() *CashSettlement1 {
	newValue := new(CashSettlement1)
	i.CashSettlement = append(i.CashSettlement, newValue)
	return newValue
}

func (i *InvestmentPlan12) SetContractReference(value string) {
	i.ContractReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan12) SetRelatedContractReference(value string) {
	i.RelatedContractReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan12) SetProductIdentification(value string) {
	i.ProductIdentification = (*Max35Text)(&value)
}

func (i *InvestmentPlan12) SetSLAChargeAndCommissionReference(value string) {
	i.SLAChargeAndCommissionReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan12) AddInsuranceCover() *InsuranceType2Choice {
	i.InsuranceCover = new(InsuranceType2Choice)
	return i.InsuranceCover
}

func (i *InvestmentPlan12) AddPlanStatus() *PlanStatus2Choice {
	i.PlanStatus = new(PlanStatus2Choice)
	return i.PlanStatus
}

func (i *InvestmentPlan12) AddInstalmentManagerRole() *PartyRole4Choice {
	i.InstalmentManagerRole = new(PartyRole4Choice)
	return i.InstalmentManagerRole
}

// Plan that allows investors to schedule periodical investments or divestments, according to pre-defined criteria.
type InvestmentPlan13 struct {

	// Frequency of the investment or divestment.
	Frequency *Frequency20Choice `xml:"Frqcy"`

	// Date the investment plan starts.
	StartDate *ISODate `xml:"StartDt,omitempty"`

	// Date the investment plan stops.
	EndDate *ISODate `xml:"EndDt,omitempty"`

	// Amount of the periodical payments.
	Quantity *UnitsOrAmount1Choice `xml:"Qty"`

	// Indicates whether an ordered amount is a gross amount (including all charges, commissions, tax). If it is not a gross amount, the ordered amount is a net amount (amount to be invested or redeemed from the fund to which other elements will be added).
	GrossAmountIndicator *YesNoIndicator `xml:"GrssAmtInd,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference2Code `xml:"IncmPref,omitempty"`

	// Initial amount or number of initial instalments.
	InitialAmount *InitialAmount1Choice `xml:"InitlAmt,omitempty"`

	// Total number of times the amount must be invested at the predefined frequency as of the start date of the investment plan.
	TotalNumberOfInstalments *Number `xml:"TtlNbOfInstlmts,omitempty"`

	// Indicates the rounding direction when an amount is to be spread over several funds.
	RoundingDirection *RoundingDirection1Code `xml:"RndgDrctn,omitempty"`

	// Security that an investment plan invests in, or from which the investment plan divests.
	SecurityDetails []*Repartition4 `xml:"SctyDtls"`

	// Cash settlement standing instruction associated to the investment plan and to be either inserted or deleted.
	ModifiedCashSettlement []*CashSettlement2 `xml:"ModfdCshSttlm,omitempty"`

	// Reference of the underlying investment contract. In some markets, such as Italy, this might be required to segregate holdings between the same investment account.
	ContractReference *Max35Text `xml:"CtrctRef,omitempty"`

	// Reference of the previous contract to which this savings or withdrawal plan is related.
	RelatedContractReference *Max35Text `xml:"RltdCtrctRef,omitempty"`

	// Identification of the product as designated by the fund manager. In some markets, such as Italy, the financial product or service related to a savings plan or withdrawal plan are identified by a product identification or number.
	ProductIdentification *Max35Text `xml:"PdctId,omitempty"`

	// Reference of the underlying service level agreement (SLA) governing charges and commission.
	SLAChargeAndCommissionReference *Max35Text `xml:"SLAChrgAndComssnRef,omitempty"`

	// Specifies the type of insurance contract to which the savings investment plan is linked.
	InsuranceCover *InsuranceType2Choice `xml:"InsrncCover,omitempty"`

	// Status of the savings or withdrawal investment plan.
	PlanStatus *PlanStatus2Choice `xml:"PlanSts,omitempty"`

	// Role or function of the instalment manager.
	InstalmentManagerRole *PartyRole4Choice `xml:"InstlmtMgrRole,omitempty"`
}

func (i *InvestmentPlan13) AddFrequency() *Frequency20Choice {
	i.Frequency = new(Frequency20Choice)
	return i.Frequency
}

func (i *InvestmentPlan13) SetStartDate(value string) {
	i.StartDate = (*ISODate)(&value)
}

func (i *InvestmentPlan13) SetEndDate(value string) {
	i.EndDate = (*ISODate)(&value)
}

func (i *InvestmentPlan13) AddQuantity() *UnitsOrAmount1Choice {
	i.Quantity = new(UnitsOrAmount1Choice)
	return i.Quantity
}

func (i *InvestmentPlan13) SetGrossAmountIndicator(value string) {
	i.GrossAmountIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentPlan13) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference2Code)(&value)
}

func (i *InvestmentPlan13) AddInitialAmount() *InitialAmount1Choice {
	i.InitialAmount = new(InitialAmount1Choice)
	return i.InitialAmount
}

func (i *InvestmentPlan13) SetTotalNumberOfInstalments(value string) {
	i.TotalNumberOfInstalments = (*Number)(&value)
}

func (i *InvestmentPlan13) SetRoundingDirection(value string) {
	i.RoundingDirection = (*RoundingDirection1Code)(&value)
}

func (i *InvestmentPlan13) AddSecurityDetails() *Repartition4 {
	newValue := new(Repartition4)
	i.SecurityDetails = append(i.SecurityDetails, newValue)
	return newValue
}

func (i *InvestmentPlan13) AddModifiedCashSettlement() *CashSettlement2 {
	newValue := new(CashSettlement2)
	i.ModifiedCashSettlement = append(i.ModifiedCashSettlement, newValue)
	return newValue
}

func (i *InvestmentPlan13) SetContractReference(value string) {
	i.ContractReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan13) SetRelatedContractReference(value string) {
	i.RelatedContractReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan13) SetProductIdentification(value string) {
	i.ProductIdentification = (*Max35Text)(&value)
}

func (i *InvestmentPlan13) SetSLAChargeAndCommissionReference(value string) {
	i.SLAChargeAndCommissionReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan13) AddInsuranceCover() *InsuranceType2Choice {
	i.InsuranceCover = new(InsuranceType2Choice)
	return i.InsuranceCover
}

func (i *InvestmentPlan13) AddPlanStatus() *PlanStatus2Choice {
	i.PlanStatus = new(PlanStatus2Choice)
	return i.PlanStatus
}

func (i *InvestmentPlan13) AddInstalmentManagerRole() *PartyRole4Choice {
	i.InstalmentManagerRole = new(PartyRole4Choice)
	return i.InstalmentManagerRole
}

// Plan that allows investors to schedule periodical investments or divestments, according to pre-defined criteria.
type InvestmentPlan14 struct {

	// Frequency of the investment or divestment.
	Frequency *Frequency20Choice `xml:"Frqcy"`

	// Date the investment plan starts.
	StartDate *ISODate `xml:"StartDt,omitempty"`

	// Date the investment plan stops.
	EndDate *ISODate `xml:"EndDt,omitempty"`

	// Amount of the periodical payments.
	Quantity *UnitsOrAmount1Choice `xml:"Qty"`

	// Indicates whether an ordered amount is a gross amount (including transaction overhead). If it is not a gross amount, the ordered amount is a net amount (amount to be invested or redeemed from the fund to which other elements will be added).
	GrossAmountIndicator *YesNoIndicator `xml:"GrssAmtInd,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference2Code `xml:"IncmPref,omitempty"`

	// Initial amount or number of initial instalments.
	InitialAmount *InitialAmount1Choice `xml:"InitlAmt,omitempty"`

	// Total number of times the amount must be invested at the predefined frequency as of the start date of the investment plan.
	TotalNumberOfInstalments *Number `xml:"TtlNbOfInstlmts,omitempty"`

	// Indicates the rounding direction when an amount is to be spread over several funds.
	RoundingDirection *RoundingDirection1Code `xml:"RndgDrctn,omitempty"`

	// Security that an investment plan invests in, or from which the investment plan divests.
	SecurityDetails []*Repartition5 `xml:"SctyDtls"`

	// Cash settlement standing instruction associated to the investment plan.
	CashSettlement []*CashSettlement1 `xml:"CshSttlm,omitempty"`

	// Reference of the underlying investment contract. In some markets, such as Italy, this might be required to segregate holdings between the same investment account.
	ContractReference *Max35Text `xml:"CtrctRef,omitempty"`

	// Reference of the previous contract to which this savings or withdrawal plan is related.
	RelatedContractReference *Max35Text `xml:"RltdCtrctRef,omitempty"`

	// Identification of the product as designated by the fund manager. In some markets, such as Italy, the financial product or service related to a savings plan or withdrawal plan are identified by a product identification or number.
	ProductIdentification *Max35Text `xml:"PdctId,omitempty"`

	// Reference of the underlying service level agreement (SLA) governing fees.
	SLAChargeAndCommissionReference *Max35Text `xml:"SLAChrgAndComssnRef,omitempty"`

	// Specifies the type of insurance contract to which the savings investment plan is linked.
	InsuranceCover *InsuranceType2Choice `xml:"InsrncCover,omitempty"`

	// Status of the savings or withdrawal investment plan.
	PlanStatus *PlanStatus2Choice `xml:"PlanSts,omitempty"`

	// Role or function of the instalment manager.
	InstalmentManagerRole *PartyRole4Choice `xml:"InstlmtMgrRole,omitempty"`
}

func (i *InvestmentPlan14) AddFrequency() *Frequency20Choice {
	i.Frequency = new(Frequency20Choice)
	return i.Frequency
}

func (i *InvestmentPlan14) SetStartDate(value string) {
	i.StartDate = (*ISODate)(&value)
}

func (i *InvestmentPlan14) SetEndDate(value string) {
	i.EndDate = (*ISODate)(&value)
}

func (i *InvestmentPlan14) AddQuantity() *UnitsOrAmount1Choice {
	i.Quantity = new(UnitsOrAmount1Choice)
	return i.Quantity
}

func (i *InvestmentPlan14) SetGrossAmountIndicator(value string) {
	i.GrossAmountIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentPlan14) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference2Code)(&value)
}

func (i *InvestmentPlan14) AddInitialAmount() *InitialAmount1Choice {
	i.InitialAmount = new(InitialAmount1Choice)
	return i.InitialAmount
}

func (i *InvestmentPlan14) SetTotalNumberOfInstalments(value string) {
	i.TotalNumberOfInstalments = (*Number)(&value)
}

func (i *InvestmentPlan14) SetRoundingDirection(value string) {
	i.RoundingDirection = (*RoundingDirection1Code)(&value)
}

func (i *InvestmentPlan14) AddSecurityDetails() *Repartition5 {
	newValue := new(Repartition5)
	i.SecurityDetails = append(i.SecurityDetails, newValue)
	return newValue
}

func (i *InvestmentPlan14) AddCashSettlement() *CashSettlement1 {
	newValue := new(CashSettlement1)
	i.CashSettlement = append(i.CashSettlement, newValue)
	return newValue
}

func (i *InvestmentPlan14) SetContractReference(value string) {
	i.ContractReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan14) SetRelatedContractReference(value string) {
	i.RelatedContractReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan14) SetProductIdentification(value string) {
	i.ProductIdentification = (*Max35Text)(&value)
}

func (i *InvestmentPlan14) SetSLAChargeAndCommissionReference(value string) {
	i.SLAChargeAndCommissionReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan14) AddInsuranceCover() *InsuranceType2Choice {
	i.InsuranceCover = new(InsuranceType2Choice)
	return i.InsuranceCover
}

func (i *InvestmentPlan14) AddPlanStatus() *PlanStatus2Choice {
	i.PlanStatus = new(PlanStatus2Choice)
	return i.PlanStatus
}

func (i *InvestmentPlan14) AddInstalmentManagerRole() *PartyRole4Choice {
	i.InstalmentManagerRole = new(PartyRole4Choice)
	return i.InstalmentManagerRole
}

// Plan that allows investors to schedule periodical investments or divestments, according to pre-defined criteria.
type InvestmentPlan15 struct {

	// Frequency of the investment or divestment.
	Frequency *Frequency20Choice `xml:"Frqcy"`

	// Date the investment plan starts.
	StartDate *ISODate `xml:"StartDt,omitempty"`

	// Date the investment plan stops.
	EndDate *ISODate `xml:"EndDt,omitempty"`

	// Amount of the periodical payments.
	Quantity *UnitsOrAmount1Choice `xml:"Qty"`

	// Indicates whether an ordered amount is a gross amount (including transaction overhead). If it is not a gross amount, the ordered amount is a net amount (amount to be invested or redeemed from the fund to which other elements will be added).
	GrossAmountIndicator *YesNoIndicator `xml:"GrssAmtInd,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference2Code `xml:"IncmPref,omitempty"`

	// Initial amount or number of initial instalments.
	InitialAmount *InitialAmount1Choice `xml:"InitlAmt,omitempty"`

	// Total number of times the amount must be invested at the predefined frequency as of the start date of the investment plan.
	TotalNumberOfInstalments *Number `xml:"TtlNbOfInstlmts,omitempty"`

	// Indicates the rounding direction when an amount is to be spread over several funds.
	RoundingDirection *RoundingDirection1Code `xml:"RndgDrctn,omitempty"`

	// Security that an investment plan invests in, or from which the investment plan divests.
	SecurityDetails []*Repartition5 `xml:"SctyDtls"`

	// Cash settlement standing instruction associated to the investment plan and to be either inserted or deleted.
	ModifiedCashSettlement []*CashSettlement2 `xml:"ModfdCshSttlm,omitempty"`

	// Reference of the underlying investment contract. In some markets, such as Italy, this might be required to segregate holdings between the same investment account.
	ContractReference *Max35Text `xml:"CtrctRef,omitempty"`

	// Reference of the previous contract to which this savings or withdrawal plan is related.
	RelatedContractReference *Max35Text `xml:"RltdCtrctRef,omitempty"`

	// Identification of the product as designated by the fund manager. In some markets, such as Italy, the financial product or service related to a savings plan or withdrawal plan are identified by a product identification or number.
	ProductIdentification *Max35Text `xml:"PdctId,omitempty"`

	// Reference of the underlying service level agreement (SLA) governing fees.
	SLAChargeAndCommissionReference *Max35Text `xml:"SLAChrgAndComssnRef,omitempty"`

	// Specifies the type of insurance contract to which the savings investment plan is linked.
	InsuranceCover *InsuranceType2Choice `xml:"InsrncCover,omitempty"`

	// Status of the savings or withdrawal investment plan.
	PlanStatus *PlanStatus2Choice `xml:"PlanSts,omitempty"`

	// Role or function of the instalment manager.
	InstalmentManagerRole *PartyRole4Choice `xml:"InstlmtMgrRole,omitempty"`
}

func (i *InvestmentPlan15) AddFrequency() *Frequency20Choice {
	i.Frequency = new(Frequency20Choice)
	return i.Frequency
}

func (i *InvestmentPlan15) SetStartDate(value string) {
	i.StartDate = (*ISODate)(&value)
}

func (i *InvestmentPlan15) SetEndDate(value string) {
	i.EndDate = (*ISODate)(&value)
}

func (i *InvestmentPlan15) AddQuantity() *UnitsOrAmount1Choice {
	i.Quantity = new(UnitsOrAmount1Choice)
	return i.Quantity
}

func (i *InvestmentPlan15) SetGrossAmountIndicator(value string) {
	i.GrossAmountIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentPlan15) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference2Code)(&value)
}

func (i *InvestmentPlan15) AddInitialAmount() *InitialAmount1Choice {
	i.InitialAmount = new(InitialAmount1Choice)
	return i.InitialAmount
}

func (i *InvestmentPlan15) SetTotalNumberOfInstalments(value string) {
	i.TotalNumberOfInstalments = (*Number)(&value)
}

func (i *InvestmentPlan15) SetRoundingDirection(value string) {
	i.RoundingDirection = (*RoundingDirection1Code)(&value)
}

func (i *InvestmentPlan15) AddSecurityDetails() *Repartition5 {
	newValue := new(Repartition5)
	i.SecurityDetails = append(i.SecurityDetails, newValue)
	return newValue
}

func (i *InvestmentPlan15) AddModifiedCashSettlement() *CashSettlement2 {
	newValue := new(CashSettlement2)
	i.ModifiedCashSettlement = append(i.ModifiedCashSettlement, newValue)
	return newValue
}

func (i *InvestmentPlan15) SetContractReference(value string) {
	i.ContractReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan15) SetRelatedContractReference(value string) {
	i.RelatedContractReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan15) SetProductIdentification(value string) {
	i.ProductIdentification = (*Max35Text)(&value)
}

func (i *InvestmentPlan15) SetSLAChargeAndCommissionReference(value string) {
	i.SLAChargeAndCommissionReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan15) AddInsuranceCover() *InsuranceType2Choice {
	i.InsuranceCover = new(InsuranceType2Choice)
	return i.InsuranceCover
}

func (i *InvestmentPlan15) AddPlanStatus() *PlanStatus2Choice {
	i.PlanStatus = new(PlanStatus2Choice)
	return i.PlanStatus
}

func (i *InvestmentPlan15) AddInstalmentManagerRole() *PartyRole4Choice {
	i.InstalmentManagerRole = new(PartyRole4Choice)
	return i.InstalmentManagerRole
}

// Plan that allows investors to schedule periodical investments or divestments, according to pre-defined criteria.
type InvestmentPlan4 struct {

	// Frequency of the investment or divestment.
	Frequency *EventFrequency1Code `xml:"Frqcy"`

	// Frequency of the investment or divestment.
	ExtendedFrequency *Extended350Code `xml:"XtndedFrqcy"`

	// Date the investment plan starts.
	StartDate *ISODate `xml:"StartDt"`

	// Date the investment plan stops.
	EndDate *ISODate `xml:"EndDt,omitempty"`

	// Currency and amount of the periodical payments.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an ordered amount is a gross amount (including all charges, commissions, tax). If it is not a gross amount, the ordered amount is a net amount (amount to be invested or redeemed from the fund to which other elements will be added).
	GrossAmountIndicator *YesNoIndicator `xml:"GrssAmtInd,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Number of pre-paid instalment periods at the time the investment plan is created.
	InitialNumberOfInstalment *Number `xml:"InitlNbOfInstlmt,omitempty"`

	// Total number of times the amount must be invested at the predefined frequency as of the start date of the investment plan.
	TotalNumberOfInstalment *Number `xml:"TtlNbOfInstlmt,omitempty"`

	// Indicates the rounding direction when an amount is to be spread over several funds.
	RoundingDirection *RoundingDirection1Code `xml:"RndgDrctn,omitempty"`

	// Security that an investment plan invests in, or from which the investment plan divests.
	SecurityDetails []*Repartition1 `xml:"SctyDtls"`

	// Cash settlement standing instruction associated to the investment plan.
	CashSettlement *InvestmentFundCashSettlementInformation3 `xml:"CshSttlm,omitempty"`
}

func (i *InvestmentPlan4) SetFrequency(value string) {
	i.Frequency = (*EventFrequency1Code)(&value)
}

func (i *InvestmentPlan4) SetExtendedFrequency(value string) {
	i.ExtendedFrequency = (*Extended350Code)(&value)
}

func (i *InvestmentPlan4) SetStartDate(value string) {
	i.StartDate = (*ISODate)(&value)
}

func (i *InvestmentPlan4) SetEndDate(value string) {
	i.EndDate = (*ISODate)(&value)
}

func (i *InvestmentPlan4) SetAmount(value, currency string) {
	i.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (i *InvestmentPlan4) SetGrossAmountIndicator(value string) {
	i.GrossAmountIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentPlan4) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference1Code)(&value)
}

func (i *InvestmentPlan4) SetInitialNumberOfInstalment(value string) {
	i.InitialNumberOfInstalment = (*Number)(&value)
}

func (i *InvestmentPlan4) SetTotalNumberOfInstalment(value string) {
	i.TotalNumberOfInstalment = (*Number)(&value)
}

func (i *InvestmentPlan4) SetRoundingDirection(value string) {
	i.RoundingDirection = (*RoundingDirection1Code)(&value)
}

func (i *InvestmentPlan4) AddSecurityDetails() *Repartition1 {
	newValue := new(Repartition1)
	i.SecurityDetails = append(i.SecurityDetails, newValue)
	return newValue
}

func (i *InvestmentPlan4) AddCashSettlement() *InvestmentFundCashSettlementInformation3 {
	i.CashSettlement = new(InvestmentFundCashSettlementInformation3)
	return i.CashSettlement
}

// Plan that allows investors to schedule periodical investments or divestments, according to pre-defined criteria.
type InvestmentPlan5 struct {

	// Frequency of the investment or divestment.
	Frequency *EventFrequency1Code `xml:"Frqcy"`

	// Frequency of the investment or divestment.
	ExtendedFrequency *Extended350Code `xml:"XtndedFrqcy"`

	// Date the investment plan starts.
	StartDate *ISODate `xml:"StartDt"`

	// Date the investment plan stops.
	EndDate *ISODate `xml:"EndDt,omitempty"`

	// Currency and amount of the periodical payments.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an ordered amount is a gross amount (including all charges, commissions, tax). If it is not a gross amount, the ordered amount is a net amount (amount to be invested or redeemed from the fund to which other elements will be added).
	GrossAmountIndicator *YesNoIndicator `xml:"GrssAmtInd,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Number of pre-paid instalment periods at the time the investment plan is created.
	InitialNumberOfInstalment *Number `xml:"InitlNbOfInstlmt,omitempty"`

	// Total number of times the amount must be invested at the predefined frequency as of the start date of the investment plan.
	TotalNumberOfInstalment *Number `xml:"TtlNbOfInstlmt,omitempty"`

	// Indicates the rounding direction when an amount is to be spread over several funds.
	RoundingDirection *RoundingDirection1Code `xml:"RndgDrctn,omitempty"`

	// Security that an investment plan invests in, or from which the investment plan divests.
	SecurityDetails []*Repartition1 `xml:"SctyDtls"`

	// Cash settlement standing instruction associated to the investment plan and to be either inserted or deleted.
	ModifiedCashSettlement []*InvestmentFundCashSettlementInformation4 `xml:"ModfdCshSttlm,omitempty"`
}

func (i *InvestmentPlan5) SetFrequency(value string) {
	i.Frequency = (*EventFrequency1Code)(&value)
}

func (i *InvestmentPlan5) SetExtendedFrequency(value string) {
	i.ExtendedFrequency = (*Extended350Code)(&value)
}

func (i *InvestmentPlan5) SetStartDate(value string) {
	i.StartDate = (*ISODate)(&value)
}

func (i *InvestmentPlan5) SetEndDate(value string) {
	i.EndDate = (*ISODate)(&value)
}

func (i *InvestmentPlan5) SetAmount(value, currency string) {
	i.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (i *InvestmentPlan5) SetGrossAmountIndicator(value string) {
	i.GrossAmountIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentPlan5) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference1Code)(&value)
}

func (i *InvestmentPlan5) SetInitialNumberOfInstalment(value string) {
	i.InitialNumberOfInstalment = (*Number)(&value)
}

func (i *InvestmentPlan5) SetTotalNumberOfInstalment(value string) {
	i.TotalNumberOfInstalment = (*Number)(&value)
}

func (i *InvestmentPlan5) SetRoundingDirection(value string) {
	i.RoundingDirection = (*RoundingDirection1Code)(&value)
}

func (i *InvestmentPlan5) AddSecurityDetails() *Repartition1 {
	newValue := new(Repartition1)
	i.SecurityDetails = append(i.SecurityDetails, newValue)
	return newValue
}

func (i *InvestmentPlan5) AddModifiedCashSettlement() *InvestmentFundCashSettlementInformation4 {
	newValue := new(InvestmentFundCashSettlementInformation4)
	i.ModifiedCashSettlement = append(i.ModifiedCashSettlement, newValue)
	return newValue
}

// Plan that allows investors to schedule periodical investments or divestments, according to pre-defined criteria.
type InvestmentPlan6 struct {

	// Frequency of the investment or divestment.
	Frequency *Frequency19Choice `xml:"Frqcy"`

	// Date the investment plan starts.
	StartDate *ISODate `xml:"StartDt,omitempty"`

	// Date the investment plan stops.
	EndDate *ISODate `xml:"EndDt,omitempty"`

	// Amount of the periodical payments.
	Quantity *UnitsOrAmount1Choice `xml:"Qty"`

	// Indicates whether an ordered amount is a gross amount (including all charges, commissions, tax). If it is not a gross amount, the ordered amount is a net amount (amount to be invested or redeemed from the fund to which other elements will be added).
	GrossAmountIndicator *YesNoIndicator `xml:"GrssAmtInd,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Number of pre-paid instalment periods at the time the investment plan is created.
	InitialNumberOfInstalment *Number `xml:"InitlNbOfInstlmt,omitempty"`

	// Total number of times the amount must be invested at the predefined frequency as of the start date of the investment plan.
	TotalNumberOfInstalment *Number `xml:"TtlNbOfInstlmt,omitempty"`

	// Indicates the rounding direction when an amount is to be spread over several funds.
	RoundingDirection *RoundingDirection1Code `xml:"RndgDrctn,omitempty"`

	// Security that an investment plan invests in, or from which the investment plan divests.
	SecurityDetails []*Repartition2 `xml:"SctyDtls"`

	// Cash settlement standing instruction associated to the investment plan.
	CashSettlement []*InvestmentFundCashSettlementInformation5 `xml:"CshSttlm,omitempty"`

	// Reference of the underlying investment contract. In some markets, such as Italy, this might be required to segregate holdings between the same investment account.
	ContractReference *Max35Text `xml:"CtrctRef,omitempty"`

	// Reference of the previous contract to which this savings or withdrawal plan is related.
	RelatedContractReference *Max35Text `xml:"RltdCtrctRef,omitempty"`

	// Identification of the product as designated by the fund manager. In some markets, such as Italy, the financial product or service related to a savings plan or withdrawal plan are identified by a product identification or number.
	ProductIdentification *Max35Text `xml:"PdctId,omitempty"`

	// Reference of the underlying service level agreement (SLA) governing charges and commission.
	SLAChargeAndCommissionReference *Max35Text `xml:"SLAChrgAndComssnRef,omitempty"`

	// Specifies the type of insurance contract to which the savings investment plan is linked.
	InsuranceCover *InsuranceType1Choice `xml:"InsrncCover,omitempty"`
}

func (i *InvestmentPlan6) AddFrequency() *Frequency19Choice {
	i.Frequency = new(Frequency19Choice)
	return i.Frequency
}

func (i *InvestmentPlan6) SetStartDate(value string) {
	i.StartDate = (*ISODate)(&value)
}

func (i *InvestmentPlan6) SetEndDate(value string) {
	i.EndDate = (*ISODate)(&value)
}

func (i *InvestmentPlan6) AddQuantity() *UnitsOrAmount1Choice {
	i.Quantity = new(UnitsOrAmount1Choice)
	return i.Quantity
}

func (i *InvestmentPlan6) SetGrossAmountIndicator(value string) {
	i.GrossAmountIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentPlan6) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference1Code)(&value)
}

func (i *InvestmentPlan6) SetInitialNumberOfInstalment(value string) {
	i.InitialNumberOfInstalment = (*Number)(&value)
}

func (i *InvestmentPlan6) SetTotalNumberOfInstalment(value string) {
	i.TotalNumberOfInstalment = (*Number)(&value)
}

func (i *InvestmentPlan6) SetRoundingDirection(value string) {
	i.RoundingDirection = (*RoundingDirection1Code)(&value)
}

func (i *InvestmentPlan6) AddSecurityDetails() *Repartition2 {
	newValue := new(Repartition2)
	i.SecurityDetails = append(i.SecurityDetails, newValue)
	return newValue
}

func (i *InvestmentPlan6) AddCashSettlement() *InvestmentFundCashSettlementInformation5 {
	newValue := new(InvestmentFundCashSettlementInformation5)
	i.CashSettlement = append(i.CashSettlement, newValue)
	return newValue
}

func (i *InvestmentPlan6) SetContractReference(value string) {
	i.ContractReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan6) SetRelatedContractReference(value string) {
	i.RelatedContractReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan6) SetProductIdentification(value string) {
	i.ProductIdentification = (*Max35Text)(&value)
}

func (i *InvestmentPlan6) SetSLAChargeAndCommissionReference(value string) {
	i.SLAChargeAndCommissionReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan6) AddInsuranceCover() *InsuranceType1Choice {
	i.InsuranceCover = new(InsuranceType1Choice)
	return i.InsuranceCover
}

// Plan that allows investors to schedule periodical investments or divestments, according to pre-defined criteria.
type InvestmentPlan7 struct {

	// Frequency of the investment or divestment.
	Frequency *Frequency19Choice `xml:"Frqcy"`

	// Date the investment plan starts.
	StartDate *ISODate `xml:"StartDt,omitempty"`

	// Date the investment plan stops.
	EndDate *ISODate `xml:"EndDt,omitempty"`

	// Amount of the periodical payments.
	Quantity *UnitsOrAmount1Choice `xml:"Qty"`

	// Indicates whether an ordered amount is a gross amount (including all charges, commissions, tax). If it is not a gross amount, the ordered amount is a net amount (amount to be invested or redeemed from the fund to which other elements will be added).
	GrossAmountIndicator *YesNoIndicator `xml:"GrssAmtInd,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Number of pre-paid instalment periods at the time the investment plan is created.
	InitialNumberOfInstalment *Number `xml:"InitlNbOfInstlmt,omitempty"`

	// Total number of times the amount must be invested at the predefined frequency as of the start date of the investment plan.
	TotalNumberOfInstalment *Number `xml:"TtlNbOfInstlmt,omitempty"`

	// Indicates the rounding direction when an amount is to be spread over several funds.
	RoundingDirection *RoundingDirection1Code `xml:"RndgDrctn,omitempty"`

	// Security that an investment plan invests in, or from which the investment plan divests.
	SecurityDetails []*Repartition2 `xml:"SctyDtls"`

	// Cash settlement standing instruction associated to the investment plan and to be either inserted or deleted.
	ModifiedCashSettlement []*InvestmentFundCashSettlementInformation6 `xml:"ModfdCshSttlm,omitempty"`

	// Reference of the underlying investment contract. In some markets, such as Italy, this might be required to segregate holdings between the same investment account.
	ContractReference *Max35Text `xml:"CtrctRef,omitempty"`

	// Reference of the previous contract to which this savings or withdrawal plan is related.
	RelatedContractReference *Max35Text `xml:"RltdCtrctRef,omitempty"`

	// Identification of the product as designated by the fund manager. In some markets, such as Italy, the financial product or service related to a savings plan or withdrawal plan are identified by a product identification or number.
	ProductIdentification *Max35Text `xml:"PdctId,omitempty"`

	// Reference of the underlying service level agreement (SLA) governing charges and commission.
	SLAChargeAndCommissionReference *Max35Text `xml:"SLAChrgAndComssnRef,omitempty"`

	// Specifies the type of insurance contract to which the savings investment plan is linked.
	InsuranceCover *InsuranceType1Choice `xml:"InsrncCover,omitempty"`
}

func (i *InvestmentPlan7) AddFrequency() *Frequency19Choice {
	i.Frequency = new(Frequency19Choice)
	return i.Frequency
}

func (i *InvestmentPlan7) SetStartDate(value string) {
	i.StartDate = (*ISODate)(&value)
}

func (i *InvestmentPlan7) SetEndDate(value string) {
	i.EndDate = (*ISODate)(&value)
}

func (i *InvestmentPlan7) AddQuantity() *UnitsOrAmount1Choice {
	i.Quantity = new(UnitsOrAmount1Choice)
	return i.Quantity
}

func (i *InvestmentPlan7) SetGrossAmountIndicator(value string) {
	i.GrossAmountIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentPlan7) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference1Code)(&value)
}

func (i *InvestmentPlan7) SetInitialNumberOfInstalment(value string) {
	i.InitialNumberOfInstalment = (*Number)(&value)
}

func (i *InvestmentPlan7) SetTotalNumberOfInstalment(value string) {
	i.TotalNumberOfInstalment = (*Number)(&value)
}

func (i *InvestmentPlan7) SetRoundingDirection(value string) {
	i.RoundingDirection = (*RoundingDirection1Code)(&value)
}

func (i *InvestmentPlan7) AddSecurityDetails() *Repartition2 {
	newValue := new(Repartition2)
	i.SecurityDetails = append(i.SecurityDetails, newValue)
	return newValue
}

func (i *InvestmentPlan7) AddModifiedCashSettlement() *InvestmentFundCashSettlementInformation6 {
	newValue := new(InvestmentFundCashSettlementInformation6)
	i.ModifiedCashSettlement = append(i.ModifiedCashSettlement, newValue)
	return newValue
}

func (i *InvestmentPlan7) SetContractReference(value string) {
	i.ContractReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan7) SetRelatedContractReference(value string) {
	i.RelatedContractReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan7) SetProductIdentification(value string) {
	i.ProductIdentification = (*Max35Text)(&value)
}

func (i *InvestmentPlan7) SetSLAChargeAndCommissionReference(value string) {
	i.SLAChargeAndCommissionReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan7) AddInsuranceCover() *InsuranceType1Choice {
	i.InsuranceCover = new(InsuranceType1Choice)
	return i.InsuranceCover
}

// Plan that allows investors to schedule periodical investments or divestments, according to pre-defined criteria.
type InvestmentPlan8 struct {

	// Frequency of the investment or divestment.
	Frequency *Frequency20Choice `xml:"Frqcy"`

	// Date the investment plan starts.
	StartDate *ISODate `xml:"StartDt,omitempty"`

	// Date the investment plan stops.
	EndDate *ISODate `xml:"EndDt,omitempty"`

	// Amount of the periodical payments.
	Quantity *UnitsOrAmount1Choice `xml:"Qty"`

	// Indicates whether an ordered amount is a gross amount (including all charges, commissions, tax). If it is not a gross amount, the ordered amount is a net amount (amount to be invested or redeemed from the fund to which other elements will be added).
	GrossAmountIndicator *YesNoIndicator `xml:"GrssAmtInd,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Number of pre-paid instalment periods at the time the investment plan is created.
	InitialNumberOfInstalments *Number `xml:"InitlNbOfInstlmts,omitempty"`

	// Total number of times the amount must be invested at the predefined frequency as of the start date of the investment plan.
	TotalNumberOfInstalments *Number `xml:"TtlNbOfInstlmts,omitempty"`

	// Indicates the rounding direction when an amount is to be spread over several funds.
	RoundingDirection *RoundingDirection1Code `xml:"RndgDrctn,omitempty"`

	// Security that an investment plan invests in, or from which the investment plan divests.
	SecurityDetails []*Repartition2 `xml:"SctyDtls"`

	// Cash settlement standing instruction associated to the investment plan and to be either inserted or deleted.
	ModifiedCashSettlement []*InvestmentFundCashSettlementInformation8 `xml:"ModfdCshSttlm,omitempty"`

	// Reference of the underlying investment contract. In some markets, such as Italy, this might be required to segregate holdings between the same investment account.
	ContractReference *Max35Text `xml:"CtrctRef,omitempty"`

	// Reference of the previous contract to which this savings or withdrawal plan is related.
	RelatedContractReference *Max35Text `xml:"RltdCtrctRef,omitempty"`

	// Identification of the product as designated by the fund manager. In some markets, such as Italy, the financial product or service related to a savings plan or withdrawal plan are identified by a product identification or number.
	ProductIdentification *Max35Text `xml:"PdctId,omitempty"`

	// Reference of the underlying service level agreement (SLA) governing charges and commission.
	SLAChargeAndCommissionReference *Max35Text `xml:"SLAChrgAndComssnRef,omitempty"`

	// Specifies the type of insurance contract to which the savings investment plan is linked.
	InsuranceCover *InsuranceType1Choice `xml:"InsrncCover,omitempty"`
}

func (i *InvestmentPlan8) AddFrequency() *Frequency20Choice {
	i.Frequency = new(Frequency20Choice)
	return i.Frequency
}

func (i *InvestmentPlan8) SetStartDate(value string) {
	i.StartDate = (*ISODate)(&value)
}

func (i *InvestmentPlan8) SetEndDate(value string) {
	i.EndDate = (*ISODate)(&value)
}

func (i *InvestmentPlan8) AddQuantity() *UnitsOrAmount1Choice {
	i.Quantity = new(UnitsOrAmount1Choice)
	return i.Quantity
}

func (i *InvestmentPlan8) SetGrossAmountIndicator(value string) {
	i.GrossAmountIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentPlan8) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference1Code)(&value)
}

func (i *InvestmentPlan8) SetInitialNumberOfInstalments(value string) {
	i.InitialNumberOfInstalments = (*Number)(&value)
}

func (i *InvestmentPlan8) SetTotalNumberOfInstalments(value string) {
	i.TotalNumberOfInstalments = (*Number)(&value)
}

func (i *InvestmentPlan8) SetRoundingDirection(value string) {
	i.RoundingDirection = (*RoundingDirection1Code)(&value)
}

func (i *InvestmentPlan8) AddSecurityDetails() *Repartition2 {
	newValue := new(Repartition2)
	i.SecurityDetails = append(i.SecurityDetails, newValue)
	return newValue
}

func (i *InvestmentPlan8) AddModifiedCashSettlement() *InvestmentFundCashSettlementInformation8 {
	newValue := new(InvestmentFundCashSettlementInformation8)
	i.ModifiedCashSettlement = append(i.ModifiedCashSettlement, newValue)
	return newValue
}

func (i *InvestmentPlan8) SetContractReference(value string) {
	i.ContractReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan8) SetRelatedContractReference(value string) {
	i.RelatedContractReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan8) SetProductIdentification(value string) {
	i.ProductIdentification = (*Max35Text)(&value)
}

func (i *InvestmentPlan8) SetSLAChargeAndCommissionReference(value string) {
	i.SLAChargeAndCommissionReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan8) AddInsuranceCover() *InsuranceType1Choice {
	i.InsuranceCover = new(InsuranceType1Choice)
	return i.InsuranceCover
}

// Plan that allows investors to schedule periodical investments or divestments, according to pre-defined criteria.
type InvestmentPlan9 struct {

	// Frequency of the investment or divestment.
	Frequency *Frequency20Choice `xml:"Frqcy"`

	// Date the investment plan starts.
	StartDate *ISODate `xml:"StartDt,omitempty"`

	// Date the investment plan stops.
	EndDate *ISODate `xml:"EndDt,omitempty"`

	// Amount of the periodical payments.
	Quantity *UnitsOrAmount1Choice `xml:"Qty"`

	// Indicates whether an ordered amount is a gross amount (including all charges, commissions, tax). If it is not a gross amount, the ordered amount is a net amount (amount to be invested or redeemed from the fund to which other elements will be added).
	GrossAmountIndicator *YesNoIndicator `xml:"GrssAmtInd,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Number of pre-paid instalment periods at the time the investment plan is created.
	InitialNumberOfInstalments *Number `xml:"InitlNbOfInstlmts,omitempty"`

	// Total number of times the amount must be invested at the predefined frequency as of the start date of the investment plan.
	TotalNumberOfInstalments *Number `xml:"TtlNbOfInstlmts,omitempty"`

	// Indicates the rounding direction when an amount is to be spread over several funds.
	RoundingDirection *RoundingDirection1Code `xml:"RndgDrctn,omitempty"`

	// Security that an investment plan invests in, or from which the investment plan divests.
	SecurityDetails []*Repartition2 `xml:"SctyDtls"`

	// Cash settlement standing instruction associated to the investment plan.
	CashSettlement []*InvestmentFundCashSettlementInformation7 `xml:"CshSttlm,omitempty"`

	// Reference of the underlying investment contract. In some markets, such as Italy, this might be required to segregate holdings between the same investment account.
	ContractReference *Max35Text `xml:"CtrctRef,omitempty"`

	// Reference of the previous contract to which this savings or withdrawal plan is related.
	RelatedContractReference *Max35Text `xml:"RltdCtrctRef,omitempty"`

	// Identification of the product as designated by the fund manager. In some markets, such as Italy, the financial product or service related to a savings plan or withdrawal plan are identified by a product identification or number.
	ProductIdentification *Max35Text `xml:"PdctId,omitempty"`

	// Reference of the underlying service level agreement (SLA) governing charges and commission.
	SLAChargeAndCommissionReference *Max35Text `xml:"SLAChrgAndComssnRef,omitempty"`

	// Specifies the type of insurance contract to which the savings investment plan is linked.
	InsuranceCover *InsuranceType1Choice `xml:"InsrncCover,omitempty"`
}

func (i *InvestmentPlan9) AddFrequency() *Frequency20Choice {
	i.Frequency = new(Frequency20Choice)
	return i.Frequency
}

func (i *InvestmentPlan9) SetStartDate(value string) {
	i.StartDate = (*ISODate)(&value)
}

func (i *InvestmentPlan9) SetEndDate(value string) {
	i.EndDate = (*ISODate)(&value)
}

func (i *InvestmentPlan9) AddQuantity() *UnitsOrAmount1Choice {
	i.Quantity = new(UnitsOrAmount1Choice)
	return i.Quantity
}

func (i *InvestmentPlan9) SetGrossAmountIndicator(value string) {
	i.GrossAmountIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentPlan9) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference1Code)(&value)
}

func (i *InvestmentPlan9) SetInitialNumberOfInstalments(value string) {
	i.InitialNumberOfInstalments = (*Number)(&value)
}

func (i *InvestmentPlan9) SetTotalNumberOfInstalments(value string) {
	i.TotalNumberOfInstalments = (*Number)(&value)
}

func (i *InvestmentPlan9) SetRoundingDirection(value string) {
	i.RoundingDirection = (*RoundingDirection1Code)(&value)
}

func (i *InvestmentPlan9) AddSecurityDetails() *Repartition2 {
	newValue := new(Repartition2)
	i.SecurityDetails = append(i.SecurityDetails, newValue)
	return newValue
}

func (i *InvestmentPlan9) AddCashSettlement() *InvestmentFundCashSettlementInformation7 {
	newValue := new(InvestmentFundCashSettlementInformation7)
	i.CashSettlement = append(i.CashSettlement, newValue)
	return newValue
}

func (i *InvestmentPlan9) SetContractReference(value string) {
	i.ContractReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan9) SetRelatedContractReference(value string) {
	i.RelatedContractReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan9) SetProductIdentification(value string) {
	i.ProductIdentification = (*Max35Text)(&value)
}

func (i *InvestmentPlan9) SetSLAChargeAndCommissionReference(value string) {
	i.SLAChargeAndCommissionReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan9) AddInsuranceCover() *InsuranceType1Choice {
	i.InsuranceCover = new(InsuranceType1Choice)
	return i.InsuranceCover
}
