package iso20022

// Information about a transfer instruction.
type PEPISATransfer11 struct {

	// Information identifying the primary individual investor, eg, name, address, social security number and date of birth.
	PrimaryIndividualInvestor *IndividualPerson8 `xml:"PmryIndvInvstr,omitempty"`

	// Information identifying the secondary individual investor, eg, name, address, social security number and date of birth.
	SecondaryIndividualInvestor *IndividualPerson8 `xml:"ScndryIndvInvstr,omitempty"`

	// Information identifying the other individual investors, eg, name, address, social security number and date of birth.
	OtherIndividualInvestor []*IndividualPerson8 `xml:"OthrIndvInvstr,omitempty"`

	// Information identifying the primary corporate investor, eg, name and address.
	PrimaryCorporateInvestor *Organisation4 `xml:"PmryCorpInvstr,omitempty"`

	// Information identifying the secondary corporate investor, eg, name and address.
	SecondaryCorporateInvestor *Organisation4 `xml:"ScndryCorpInvstr,omitempty"`

	// Information identifying the other corporate investors, eg, name and address.
	OtherCorporateInvestor []*Organisation4 `xml:"OthrCorpInvstr,omitempty"`

	// Identification of an account owned by the investor at the old plan manager (account servicer).
	TransferorAccount *Account5 `xml:"TrfrAcct"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	NomineeAccount *Account6 `xml:"NmneeAcct,omitempty"`

	// Information related to the institution to which the financial instrument is to be transferred.
	Transferee *PartyIdentification2Choice `xml:"Trfee"`

	// Identification of an account owned by the investor to which a cash entry is made based on the transfer of asset(s).
	CashAccount *CashAccount11 `xml:"CshAcct,omitempty"`

	// Provides information related to the asset(s) transferred.
	ProductTransfer []*ISATransfer3 `xml:"PdctTrf"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (p *PEPISATransfer11) AddPrimaryIndividualInvestor() *IndividualPerson8 {
	p.PrimaryIndividualInvestor = new(IndividualPerson8)
	return p.PrimaryIndividualInvestor
}

func (p *PEPISATransfer11) AddSecondaryIndividualInvestor() *IndividualPerson8 {
	p.SecondaryIndividualInvestor = new(IndividualPerson8)
	return p.SecondaryIndividualInvestor
}

func (p *PEPISATransfer11) AddOtherIndividualInvestor() *IndividualPerson8 {
	newValue := new(IndividualPerson8)
	p.OtherIndividualInvestor = append(p.OtherIndividualInvestor, newValue)
	return newValue
}

func (p *PEPISATransfer11) AddPrimaryCorporateInvestor() *Organisation4 {
	p.PrimaryCorporateInvestor = new(Organisation4)
	return p.PrimaryCorporateInvestor
}

func (p *PEPISATransfer11) AddSecondaryCorporateInvestor() *Organisation4 {
	p.SecondaryCorporateInvestor = new(Organisation4)
	return p.SecondaryCorporateInvestor
}

func (p *PEPISATransfer11) AddOtherCorporateInvestor() *Organisation4 {
	newValue := new(Organisation4)
	p.OtherCorporateInvestor = append(p.OtherCorporateInvestor, newValue)
	return newValue
}

func (p *PEPISATransfer11) AddTransferorAccount() *Account5 {
	p.TransferorAccount = new(Account5)
	return p.TransferorAccount
}

func (p *PEPISATransfer11) AddNomineeAccount() *Account6 {
	p.NomineeAccount = new(Account6)
	return p.NomineeAccount
}

func (p *PEPISATransfer11) AddTransferee() *PartyIdentification2Choice {
	p.Transferee = new(PartyIdentification2Choice)
	return p.Transferee
}

func (p *PEPISATransfer11) AddCashAccount() *CashAccount11 {
	p.CashAccount = new(CashAccount11)
	return p.CashAccount
}

func (p *PEPISATransfer11) AddProductTransfer() *ISATransfer3 {
	newValue := new(ISATransfer3)
	p.ProductTransfer = append(p.ProductTransfer, newValue)
	return newValue
}

func (p *PEPISATransfer11) AddExtension() *Extension1 {
	newValue := new(Extension1)
	p.Extension = append(p.Extension, newValue)
	return newValue
}

// Describes the type of product and the assets to be transferred.
type PEPISATransfer3 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Identification assigned by the new plan manager to each transfer of asset.
	TransferIdentification *Max35Text `xml:"TrfId"`

	// Indicates whether there is cash in the account that is awaiting investment.
	ResidualCashIndicator *YesNoIndicator `xml:"RsdlCshInd"`

	// UK government schemes to encourage individuals to invest in securities based unit and investment trusts, offering certain tax benefits. These are not investment in their own right but are tax exempt wrappers in which individuals can hold equities, bonds and funds to shelter them from income and capital gains tax.
	// The Personal Equity Plan (PEP) and the Individual Savings Account (ISA) are provided only by UK based financial institutions.
	ISA *ISAYearsOfIssue1 `xml:"ISA"`

	// UK government schemes to encourage individuals to invest in securities based unit and investment trusts, offering certain tax benefits. These are not investment in their own right but are tax exempt wrappers in which individuals can hold equities, bonds and funds to shelter them from income and capital gains tax.
	//
	// The Personal Equity Plan (PEP) and the Individual Savings Account (ISA) are provided only by UK based financial institutions.
	PEP *PreviousYearChoice `xml:"PEP"`

	// Wrapper for a specific product or a specific sub-product owned by a set of beneficial owners.
	Portfolio *Portfolio1 `xml:"Prtfl"`

	// Specifies the underlying assets for the PEP, ISA or portfolio.
	FinancialInstrumentAssetForTransfer []*FinancialInstrument11 `xml:"FinInstrmAsstForTrf,omitempty"`
}

func (p *PEPISATransfer3) SetMasterReference(value string) {
	p.MasterReference = (*Max35Text)(&value)
}

func (p *PEPISATransfer3) SetTransferIdentification(value string) {
	p.TransferIdentification = (*Max35Text)(&value)
}

func (p *PEPISATransfer3) SetResidualCashIndicator(value string) {
	p.ResidualCashIndicator = (*YesNoIndicator)(&value)
}

func (p *PEPISATransfer3) AddISA() *ISAYearsOfIssue1 {
	p.ISA = new(ISAYearsOfIssue1)
	return p.ISA
}

func (p *PEPISATransfer3) AddPEP() *PreviousYearChoice {
	p.PEP = new(PreviousYearChoice)
	return p.PEP
}

func (p *PEPISATransfer3) AddPortfolio() *Portfolio1 {
	p.Portfolio = new(Portfolio1)
	return p.Portfolio
}

func (p *PEPISATransfer3) AddFinancialInstrumentAssetForTransfer() *FinancialInstrument11 {
	newValue := new(FinancialInstrument11)
	p.FinancialInstrumentAssetForTransfer = append(p.FinancialInstrumentAssetForTransfer, newValue)
	return newValue
}

// Describes the type of product and the assets to be transferred.
type PEPISATransfer4 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Identification of the confirmation assigned by the old plan manager to the transfer of account.
	TransferConfirmationIdentification *Max35Text `xml:"TrfConfId"`

	// Identification received by the old plan manager and assigned by the new plan manager to the transfer of account.
	TransferInstructionReference *Max35Text `xml:"TrfInstrRef"`

	// Date when the transfer instruction was executed.
	ActualTransferDate *ISODate `xml:"ActlTrfDt"`

	// Indicates whether there is cash in the account that is awaiting investment.
	ResidualCashIndicator *YesNoIndicator `xml:"RsdlCshInd"`

	// UK government schemes to encourage individuals to invest in securities based unit and investment trusts, offering certain tax benefits. These are not investment in their own right but are tax exempt wrappers in which individuals can hold equities, bonds and funds to shelter them from income and capital gains tax.
	//
	// The Personal Equity Plan (PEP) and the Individual Savings Account (ISA) are provided only by UK based financial institutions.
	ISA *ISAYearsOfIssue3 `xml:"ISA"`

	// UK government schemes to encourage individuals to invest in securities based unit and investment trusts, offering certain tax benefits. These are not investment in their own right but are tax exempt wrappers in which individuals can hold equities, bonds and funds to shelter them from income and capital gains tax.
	//
	// The Personal Equity Plan (PEP) and the Individual Savings Account (ISA) are provided only by UK based financial institutions.
	PEP *PreviousYearChoice `xml:"PEP"`

	// Wrapper for a specific product or a specific sub-product owned by a set of beneficial owners.
	Portfolio *Portfolio1 `xml:"Prtfl"`

	// Specifies the underlying assets for the PEP, ISA or portfolio.
	FinancialInstrumentAssetForTransfer []*FinancialInstrument11 `xml:"FinInstrmAsstForTrf,omitempty"`
}

func (p *PEPISATransfer4) SetMasterReference(value string) {
	p.MasterReference = (*Max35Text)(&value)
}

func (p *PEPISATransfer4) SetTransferConfirmationIdentification(value string) {
	p.TransferConfirmationIdentification = (*Max35Text)(&value)
}

func (p *PEPISATransfer4) SetTransferInstructionReference(value string) {
	p.TransferInstructionReference = (*Max35Text)(&value)
}

func (p *PEPISATransfer4) SetActualTransferDate(value string) {
	p.ActualTransferDate = (*ISODate)(&value)
}

func (p *PEPISATransfer4) SetResidualCashIndicator(value string) {
	p.ResidualCashIndicator = (*YesNoIndicator)(&value)
}

func (p *PEPISATransfer4) AddISA() *ISAYearsOfIssue3 {
	p.ISA = new(ISAYearsOfIssue3)
	return p.ISA
}

func (p *PEPISATransfer4) AddPEP() *PreviousYearChoice {
	p.PEP = new(PreviousYearChoice)
	return p.PEP
}

func (p *PEPISATransfer4) AddPortfolio() *Portfolio1 {
	p.Portfolio = new(Portfolio1)
	return p.Portfolio
}

func (p *PEPISATransfer4) AddFinancialInstrumentAssetForTransfer() *FinancialInstrument11 {
	newValue := new(FinancialInstrument11)
	p.FinancialInstrumentAssetForTransfer = append(p.FinancialInstrumentAssetForTransfer, newValue)
	return newValue
}

// Describes the type of product and the assets to be transferred.
type PEPISATransfer5 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Identification assigned by the new plan manager to each transfer of asset.
	TransferIdentification *Max35Text `xml:"TrfId"`

	// UK government schemes to encourage individuals to invest in securities based unit and investment trusts, offering certain tax benefits. These are not investment in their own right but are tax exempt wrappers in which individuals can hold equities, bonds and funds to shelter them from income and capital gains tax.
	// The Personal Equity Plan (PEP) and the Individual Savings Account (ISA) are provided only by UK based financial institutions.
	ISA *ISAYearsOfIssue2 `xml:"ISA"`

	// UK government schemes to encourage individuals to invest in securities based unit and investment trusts, offering certain tax benefits. These are not investment in their own right but are tax exempt wrappers in which individuals can hold equities, bonds and funds to shelter them from income and capital gains tax.
	// The Personal Equity Plan (PEP) and the Individual Savings Account (ISA) are provided only by UK based financial institutions.
	PEP *PreviousYearChoice `xml:"PEP"`

	// Wrapper for a specific product or a specific sub-product owned by a set of beneficial owners.
	Portfolio *Portfolio1 `xml:"Prtfl"`

	// Specifies the underlying assets for the PEP, ISA or portfolio.
	FinancialInstrumentAssetForTransfer []*FinancialInstrument12 `xml:"FinInstrmAsstForTrf,omitempty"`
}

func (p *PEPISATransfer5) SetMasterReference(value string) {
	p.MasterReference = (*Max35Text)(&value)
}

func (p *PEPISATransfer5) SetTransferIdentification(value string) {
	p.TransferIdentification = (*Max35Text)(&value)
}

func (p *PEPISATransfer5) AddISA() *ISAYearsOfIssue2 {
	p.ISA = new(ISAYearsOfIssue2)
	return p.ISA
}

func (p *PEPISATransfer5) AddPEP() *PreviousYearChoice {
	p.PEP = new(PreviousYearChoice)
	return p.PEP
}

func (p *PEPISATransfer5) AddPortfolio() *Portfolio1 {
	p.Portfolio = new(Portfolio1)
	return p.Portfolio
}

func (p *PEPISATransfer5) AddFinancialInstrumentAssetForTransfer() *FinancialInstrument12 {
	newValue := new(FinancialInstrument12)
	p.FinancialInstrumentAssetForTransfer = append(p.FinancialInstrumentAssetForTransfer, newValue)
	return newValue
}

// Describes the type of product and the assets to be transferred.
type PEPISATransfer6 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Identification assigned by the new plan manager to each transfer of asset.
	TransferIdentification *Max35Text `xml:"TrfId"`

	// Indicates whether there is cash in the account that is awaiting investment.
	ResidualCashIndicator *YesNoIndicator `xml:"RsdlCshInd"`

	// UK government schemes to encourage individuals to invest in securities based unit and investment trusts, offering certain tax benefits. These are not investment in their own right but are tax exempt wrappers in which individuals can hold equities, bonds and funds to shelter them from income and capital gains tax.
	// The Personal Equity Plan (PEP) and the Individual Savings Account (ISA) are provided only by UK based financial institutions.
	ISA *ISAYearsOfIssue1 `xml:"ISA"`

	// UK government schemes to encourage individuals to invest in securities based unit and investment trusts, offering certain tax benefits. These are not investment in their own right but are tax exempt wrappers in which individuals can hold equities, bonds and funds to shelter them from income and capital gains tax.
	//
	// The Personal Equity Plan (PEP) and the Individual Savings Account (ISA) are provided only by UK based financial institutions.
	PEP *PreviousYearChoice `xml:"PEP"`

	// Wrapper for a specific product or a specific sub-product owned by a set of beneficial owners.
	Portfolio *Portfolio1 `xml:"Prtfl"`

	// Specifies the underlying assets for the PEP, ISA or portfolio.
	FinancialInstrumentAssetForTransfer []*FinancialInstrument12 `xml:"FinInstrmAsstForTrf,omitempty"`
}

func (p *PEPISATransfer6) SetMasterReference(value string) {
	p.MasterReference = (*Max35Text)(&value)
}

func (p *PEPISATransfer6) SetTransferIdentification(value string) {
	p.TransferIdentification = (*Max35Text)(&value)
}

func (p *PEPISATransfer6) SetResidualCashIndicator(value string) {
	p.ResidualCashIndicator = (*YesNoIndicator)(&value)
}

func (p *PEPISATransfer6) AddISA() *ISAYearsOfIssue1 {
	p.ISA = new(ISAYearsOfIssue1)
	return p.ISA
}

func (p *PEPISATransfer6) AddPEP() *PreviousYearChoice {
	p.PEP = new(PreviousYearChoice)
	return p.PEP
}

func (p *PEPISATransfer6) AddPortfolio() *Portfolio1 {
	p.Portfolio = new(Portfolio1)
	return p.Portfolio
}

func (p *PEPISATransfer6) AddFinancialInstrumentAssetForTransfer() *FinancialInstrument12 {
	newValue := new(FinancialInstrument12)
	p.FinancialInstrumentAssetForTransfer = append(p.FinancialInstrumentAssetForTransfer, newValue)
	return newValue
}

// Information about a transfer instruction.
type PEPISATransfer7 struct {

	// Information identifying the primary individual investor, eg, name, address, social security number and date of birth.
	PrimaryIndividualInvestor *IndividualPerson8 `xml:"PmryIndvInvstr,omitempty"`

	// Information identifying the secondary individual investor, eg, name, address, social security number and date of birth.
	SecondaryIndividualInvestor *IndividualPerson8 `xml:"ScndryIndvInvstr,omitempty"`

	// Information identifying the other individual investors, eg, name, address, social security number and date of birth.
	OtherIndividualInvestor []*IndividualPerson8 `xml:"OthrIndvInvstr,omitempty"`

	// Information identifying the primary corporate investor, eg, name and address.
	PrimaryCorporateInvestor *Organisation4 `xml:"PmryCorpInvstr,omitempty"`

	// Information identifying the secondary corporate investor, eg, name and address.
	SecondaryCorporateInvestor *Organisation4 `xml:"ScndryCorpInvstr,omitempty"`

	// Information identifying the other corporate investors, eg, name and address.
	OtherCorporateInvestor []*Organisation4 `xml:"OthrCorpInvstr,omitempty"`

	// Identification of an account owned by the investor at the old plan manager (account servicer).
	ClientAccount *Account5 `xml:"ClntAcct"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	NomineeAccount *Account6 `xml:"NmneeAcct,omitempty"`

	// Information related to the institution to which the financial instrument is to be transferred.
	NewPlanManager *PartyIdentification2Choice `xml:"NewPlanMgr"`

	// Identification of an account owned by the investor to which a cash entry is made based on the transfer of asset(s).
	CashAccount *CashAccount11 `xml:"CshAcct,omitempty"`

	// Provides information related to the asset(s) transferred.
	ProductTransfer []*PEPISATransfer8 `xml:"PdctTrf"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (p *PEPISATransfer7) AddPrimaryIndividualInvestor() *IndividualPerson8 {
	p.PrimaryIndividualInvestor = new(IndividualPerson8)
	return p.PrimaryIndividualInvestor
}

func (p *PEPISATransfer7) AddSecondaryIndividualInvestor() *IndividualPerson8 {
	p.SecondaryIndividualInvestor = new(IndividualPerson8)
	return p.SecondaryIndividualInvestor
}

func (p *PEPISATransfer7) AddOtherIndividualInvestor() *IndividualPerson8 {
	newValue := new(IndividualPerson8)
	p.OtherIndividualInvestor = append(p.OtherIndividualInvestor, newValue)
	return newValue
}

func (p *PEPISATransfer7) AddPrimaryCorporateInvestor() *Organisation4 {
	p.PrimaryCorporateInvestor = new(Organisation4)
	return p.PrimaryCorporateInvestor
}

func (p *PEPISATransfer7) AddSecondaryCorporateInvestor() *Organisation4 {
	p.SecondaryCorporateInvestor = new(Organisation4)
	return p.SecondaryCorporateInvestor
}

func (p *PEPISATransfer7) AddOtherCorporateInvestor() *Organisation4 {
	newValue := new(Organisation4)
	p.OtherCorporateInvestor = append(p.OtherCorporateInvestor, newValue)
	return newValue
}

func (p *PEPISATransfer7) AddClientAccount() *Account5 {
	p.ClientAccount = new(Account5)
	return p.ClientAccount
}

func (p *PEPISATransfer7) AddNomineeAccount() *Account6 {
	p.NomineeAccount = new(Account6)
	return p.NomineeAccount
}

func (p *PEPISATransfer7) AddNewPlanManager() *PartyIdentification2Choice {
	p.NewPlanManager = new(PartyIdentification2Choice)
	return p.NewPlanManager
}

func (p *PEPISATransfer7) AddCashAccount() *CashAccount11 {
	p.CashAccount = new(CashAccount11)
	return p.CashAccount
}

func (p *PEPISATransfer7) AddProductTransfer() *PEPISATransfer8 {
	newValue := new(PEPISATransfer8)
	p.ProductTransfer = append(p.ProductTransfer, newValue)
	return newValue
}

func (p *PEPISATransfer7) AddExtension() *Extension1 {
	newValue := new(Extension1)
	p.Extension = append(p.Extension, newValue)
	return newValue
}

// Describes the type of product and the assets to be transferred.
type PEPISATransfer8 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Identification assigned by the new plan manager to each transfer of asset.
	TransferIdentification *Max35Text `xml:"TrfId"`

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Indicates whether there is cash in the account that is awaiting investment.
	ResidualCashIndicator *YesNoIndicator `xml:"RsdlCshInd"`

	// UK government schemes to encourage individuals to invest in securities based unit and investment trusts, offering certain tax benefits. These are not investment in their own right but are tax exempt wrappers in which individuals can hold equities, bonds and funds to shelter them from income and capital gains tax.
	// The Personal Equity Plan (PEP) and the Individual Savings Account (ISA) are provided only by UK based financial institutions.
	ISA *ISAYearsOfIssue1 `xml:"ISA"`

	// UK government schemes to encourage individuals to invest in securities based unit and investment trusts, offering certain tax benefits. These are not investment in their own right but are tax exempt wrappers in which individuals can hold equities, bonds and funds to shelter them from income and capital gains tax.
	//
	// The Personal Equity Plan (PEP) and the Individual Savings Account (ISA) are provided only by UK based financial institutions.
	PEP *PreviousYearChoice `xml:"PEP"`

	// Wrapper for a specific product or a specific sub-product owned by a set of beneficial owners.
	Portfolio *Portfolio1 `xml:"Prtfl"`

	// Specifies the underlying assets for the PEP, ISA or portfolio.
	FinancialInstrumentAssetForTransfer []*FinancialInstrument11 `xml:"FinInstrmAsstForTrf,omitempty"`
}

func (p *PEPISATransfer8) SetMasterReference(value string) {
	p.MasterReference = (*Max35Text)(&value)
}

func (p *PEPISATransfer8) SetTransferIdentification(value string) {
	p.TransferIdentification = (*Max35Text)(&value)
}

func (p *PEPISATransfer8) SetCancellationReference(value string) {
	p.CancellationReference = (*Max35Text)(&value)
}

func (p *PEPISATransfer8) SetResidualCashIndicator(value string) {
	p.ResidualCashIndicator = (*YesNoIndicator)(&value)
}

func (p *PEPISATransfer8) AddISA() *ISAYearsOfIssue1 {
	p.ISA = new(ISAYearsOfIssue1)
	return p.ISA
}

func (p *PEPISATransfer8) AddPEP() *PreviousYearChoice {
	p.PEP = new(PreviousYearChoice)
	return p.PEP
}

func (p *PEPISATransfer8) AddPortfolio() *Portfolio1 {
	p.Portfolio = new(Portfolio1)
	return p.Portfolio
}

func (p *PEPISATransfer8) AddFinancialInstrumentAssetForTransfer() *FinancialInstrument11 {
	newValue := new(FinancialInstrument11)
	p.FinancialInstrumentAssetForTransfer = append(p.FinancialInstrumentAssetForTransfer, newValue)
	return newValue
}
