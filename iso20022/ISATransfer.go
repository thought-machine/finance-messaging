package iso20022

// Describes the type of product and the assets to be transferred.
type ISATransfer1 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Identification assigned by the new plan manager to each transfer of asset.
	TransferIdentification *Max35Text `xml:"TrfId"`

	// Requested date at which the assets should be transferred.
	RequestedTransferDate *DateFormat1Choice `xml:"ReqdTrfDt,omitempty"`

	// Specifies portfolio information or government schemes, for example Individual Savings Account (ISA) in the UK.
	Portfolio *ISAPortfolio1Choice `xml:"Prtfl"`

	// Indicates whether there is cash in the account that is awaiting investment.
	ResidualCash *ResidualCash1Code `xml:"RsdlCsh,omitempty"`

	// Indicator that all remaining assets in a portfolio not listed for transfer should be liquidated and transferred as cash.
	AllOtherCash *YesNoIndicator `xml:"AllOthrCsh"`

	// Specifies the underlying assets for the ISA or portfolio.
	FinancialInstrumentAssetForTransfer []*FinancialInstrument23 `xml:"FinInstrmAsstForTrf,omitempty"`
}

func (i *ISATransfer1) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *ISATransfer1) SetTransferIdentification(value string) {
	i.TransferIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer1) AddRequestedTransferDate() *DateFormat1Choice {
	i.RequestedTransferDate = new(DateFormat1Choice)
	return i.RequestedTransferDate
}

func (i *ISATransfer1) AddPortfolio() *ISAPortfolio1Choice {
	i.Portfolio = new(ISAPortfolio1Choice)
	return i.Portfolio
}

func (i *ISATransfer1) SetResidualCash(value string) {
	i.ResidualCash = (*ResidualCash1Code)(&value)
}

func (i *ISATransfer1) SetAllOtherCash(value string) {
	i.AllOtherCash = (*YesNoIndicator)(&value)
}

func (i *ISATransfer1) AddFinancialInstrumentAssetForTransfer() *FinancialInstrument23 {
	newValue := new(FinancialInstrument23)
	i.FinancialInstrumentAssetForTransfer = append(i.FinancialInstrumentAssetForTransfer, newValue)
	return newValue
}

// Describes the type of product and the assets to be transferred.
type ISATransfer10 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Identification of the confirmation assigned by the old plan manager to the transfer of account.
	TransferConfirmationIdentification *Max35Text `xml:"TrfConfId"`

	// Identification received by the old plan manager and assigned by the new plan manager to the transfer of account.
	TransferInstructionReference *Max35Text `xml:"TrfInstrRef"`

	// Date when the transfer instruction was executed.
	ActualTransferDate *ISODate `xml:"ActlTrfDt"`

	// Indicates whether there is cash in the account that is awaiting investment.
	ResidualCash *ResidualCash1Code `xml:"RsdlCsh,omitempty"`

	// Specifies portfolio information or government schemes, for example Individual Savings Account (ISA) in the UK.
	Portfolio *ISAPortfolio2Choice `xml:"Prtfl"`

	// Indicator that all remaining assets in a portfolio not listed for transfer should be liquidated and transferred as cash.
	AllOtherCash *YesNoIndicator `xml:"AllOthrCsh"`

	// Specifies the underlying assets for the ISA or portfolio.
	FinancialInstrumentAssetForTransfer []*FinancialInstrument33 `xml:"FinInstrmAsstForTrf,omitempty"`
}

func (i *ISATransfer10) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *ISATransfer10) SetTransferConfirmationIdentification(value string) {
	i.TransferConfirmationIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer10) SetTransferInstructionReference(value string) {
	i.TransferInstructionReference = (*Max35Text)(&value)
}

func (i *ISATransfer10) SetActualTransferDate(value string) {
	i.ActualTransferDate = (*ISODate)(&value)
}

func (i *ISATransfer10) SetResidualCash(value string) {
	i.ResidualCash = (*ResidualCash1Code)(&value)
}

func (i *ISATransfer10) AddPortfolio() *ISAPortfolio2Choice {
	i.Portfolio = new(ISAPortfolio2Choice)
	return i.Portfolio
}

func (i *ISATransfer10) SetAllOtherCash(value string) {
	i.AllOtherCash = (*YesNoIndicator)(&value)
}

func (i *ISATransfer10) AddFinancialInstrumentAssetForTransfer() *FinancialInstrument33 {
	newValue := new(FinancialInstrument33)
	i.FinancialInstrumentAssetForTransfer = append(i.FinancialInstrumentAssetForTransfer, newValue)
	return newValue
}

// Describes the type of product and the assets to be transferred.
type ISATransfer12 struct {

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
	TransferorAccount *Account15 `xml:"TrfrAcct"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	NomineeAccount *Account16 `xml:"NmneeAcct,omitempty"`

	// Information related to the institution to which the financial instrument is to be transferred.
	Transferee *PartyIdentification2Choice `xml:"Trfee"`

	// Identification of an account owned by the investor to which a cash entry is made based on the transfer of asset(s).
	CashAccount *CashAccount29 `xml:"CshAcct,omitempty"`

	// Details of the transfer to be cancelled.
	ProductTransferAndReference *ISATransfer17 `xml:"PdctTrfAndRef"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (i *ISATransfer12) AddPrimaryIndividualInvestor() *IndividualPerson8 {
	i.PrimaryIndividualInvestor = new(IndividualPerson8)
	return i.PrimaryIndividualInvestor
}

func (i *ISATransfer12) AddSecondaryIndividualInvestor() *IndividualPerson8 {
	i.SecondaryIndividualInvestor = new(IndividualPerson8)
	return i.SecondaryIndividualInvestor
}

func (i *ISATransfer12) AddOtherIndividualInvestor() *IndividualPerson8 {
	newValue := new(IndividualPerson8)
	i.OtherIndividualInvestor = append(i.OtherIndividualInvestor, newValue)
	return newValue
}

func (i *ISATransfer12) AddPrimaryCorporateInvestor() *Organisation4 {
	i.PrimaryCorporateInvestor = new(Organisation4)
	return i.PrimaryCorporateInvestor
}

func (i *ISATransfer12) AddSecondaryCorporateInvestor() *Organisation4 {
	i.SecondaryCorporateInvestor = new(Organisation4)
	return i.SecondaryCorporateInvestor
}

func (i *ISATransfer12) AddOtherCorporateInvestor() *Organisation4 {
	newValue := new(Organisation4)
	i.OtherCorporateInvestor = append(i.OtherCorporateInvestor, newValue)
	return newValue
}

func (i *ISATransfer12) AddTransferorAccount() *Account15 {
	i.TransferorAccount = new(Account15)
	return i.TransferorAccount
}

func (i *ISATransfer12) AddNomineeAccount() *Account16 {
	i.NomineeAccount = new(Account16)
	return i.NomineeAccount
}

func (i *ISATransfer12) AddTransferee() *PartyIdentification2Choice {
	i.Transferee = new(PartyIdentification2Choice)
	return i.Transferee
}

func (i *ISATransfer12) AddCashAccount() *CashAccount29 {
	i.CashAccount = new(CashAccount29)
	return i.CashAccount
}

func (i *ISATransfer12) AddProductTransferAndReference() *ISATransfer17 {
	i.ProductTransferAndReference = new(ISATransfer17)
	return i.ProductTransferAndReference
}

func (i *ISATransfer12) AddExtension() *Extension1 {
	newValue := new(Extension1)
	i.Extension = append(i.Extension, newValue)
	return newValue
}

// Describes the type of product and the assets to be transferred.
type ISATransfer13 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Identification of the confirmation assigned by the transferor to the transfer.
	TransferConfirmationIdentification *Max35Text `xml:"TrfConfId"`

	// Identification assigned to the transfer of asset, typically assigned by the transferee.
	TransferInstructionReference *Max35Text `xml:"TrfInstrRef"`

	// Date when the transfer instruction was executed.
	ActualTransferDate *ISODate `xml:"ActlTrfDt,omitempty"`

	// Indicates whether there is cash in the account that is awaiting investment.
	ResidualCash *ResidualCash1Code `xml:"RsdlCsh,omitempty"`

	// Specifies portfolio information or government schemes, for example Individual Savings Account (ISA) in the UK.
	Portfolio *ISAPortfolio2Choice `xml:"Prtfl,omitempty"`

	// Indicator that all remaining assets in a portfolio not listed for transfer should be liquidated and transferred as cash.
	AllOtherCash *YesNoIndicator `xml:"AllOthrCsh"`

	// Specifies the underlying assets for the ISA or portfolio.
	FinancialInstrumentAssetForTransfer []*FinancialInstrument35 `xml:"FinInstrmAsstForTrf,omitempty"`
}

func (i *ISATransfer13) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *ISATransfer13) SetTransferConfirmationIdentification(value string) {
	i.TransferConfirmationIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer13) SetTransferInstructionReference(value string) {
	i.TransferInstructionReference = (*Max35Text)(&value)
}

func (i *ISATransfer13) SetActualTransferDate(value string) {
	i.ActualTransferDate = (*ISODate)(&value)
}

func (i *ISATransfer13) SetResidualCash(value string) {
	i.ResidualCash = (*ResidualCash1Code)(&value)
}

func (i *ISATransfer13) AddPortfolio() *ISAPortfolio2Choice {
	i.Portfolio = new(ISAPortfolio2Choice)
	return i.Portfolio
}

func (i *ISATransfer13) SetAllOtherCash(value string) {
	i.AllOtherCash = (*YesNoIndicator)(&value)
}

func (i *ISATransfer13) AddFinancialInstrumentAssetForTransfer() *FinancialInstrument35 {
	newValue := new(FinancialInstrument35)
	i.FinancialInstrumentAssetForTransfer = append(i.FinancialInstrumentAssetForTransfer, newValue)
	return newValue
}

// Describes the type of product and the assets to be transferred.
type ISATransfer14 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Identification assigned to the transfer of asset, typically assigned by the transferee.
	TransferIdentification *Max35Text `xml:"TrfId"`

	// Identification of the confirmation assigned by the transferor to the transfer.
	TransferConfirmationIdentification *Max35Text `xml:"TrfConfId,omitempty"`

	// Indicates whether there is cash in the account that is awaiting investment.
	ResidualCash *ResidualCash1Code `xml:"RsdlCsh,omitempty"`

	// Specifies portfolio information or government schemes, for example Individual Savings Account (ISA) in the UK.
	Portfolio *ISAPortfolio1Choice `xml:"Prtfl,omitempty"`

	// Indicator that all remaining assets in a portfolio not listed for transfer should be liquidated and transferred as cash.
	AllOtherCash *YesNoIndicator `xml:"AllOthrCsh"`

	// Specifies the underlying assets for the ISA or portfolio.
	FinancialInstrumentAssetForTransfer []*FinancialInstrument37 `xml:"FinInstrmAsstForTrf,omitempty"`
}

func (i *ISATransfer14) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *ISATransfer14) SetTransferIdentification(value string) {
	i.TransferIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer14) SetTransferConfirmationIdentification(value string) {
	i.TransferConfirmationIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer14) SetResidualCash(value string) {
	i.ResidualCash = (*ResidualCash1Code)(&value)
}

func (i *ISATransfer14) AddPortfolio() *ISAPortfolio1Choice {
	i.Portfolio = new(ISAPortfolio1Choice)
	return i.Portfolio
}

func (i *ISATransfer14) SetAllOtherCash(value string) {
	i.AllOtherCash = (*YesNoIndicator)(&value)
}

func (i *ISATransfer14) AddFinancialInstrumentAssetForTransfer() *FinancialInstrument37 {
	newValue := new(FinancialInstrument37)
	i.FinancialInstrumentAssetForTransfer = append(i.FinancialInstrumentAssetForTransfer, newValue)
	return newValue
}

// Describes the type of product and the assets to be transferred.
type ISATransfer15 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Identification assigned to the transfer of asset, typically assigned by the transferee.
	TransferIdentification *Max35Text `xml:"TrfId"`

	// Specifies portfolio information or government schemes, for example Individual Savings Account (ISA) in the UK.
	Portfolio *ISAPortfolio3Choice `xml:"Prtfl,omitempty"`

	// Indicator that all remaining assets in a portfolio not listed for transfer should be liquidated and transferred as cash.
	AllOtherCash *YesNoIndicator `xml:"AllOthrCsh"`

	// Specifies the underlying assets for the ISA or portfolio.
	FinancialInstrumentAssetForTransfer []*FinancialInstrument36 `xml:"FinInstrmAsstForTrf,omitempty"`
}

func (i *ISATransfer15) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *ISATransfer15) SetTransferIdentification(value string) {
	i.TransferIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer15) AddPortfolio() *ISAPortfolio3Choice {
	i.Portfolio = new(ISAPortfolio3Choice)
	return i.Portfolio
}

func (i *ISATransfer15) SetAllOtherCash(value string) {
	i.AllOtherCash = (*YesNoIndicator)(&value)
}

func (i *ISATransfer15) AddFinancialInstrumentAssetForTransfer() *FinancialInstrument36 {
	newValue := new(FinancialInstrument36)
	i.FinancialInstrumentAssetForTransfer = append(i.FinancialInstrumentAssetForTransfer, newValue)
	return newValue
}

// Describes the type of product and the assets to be transferred.
type ISATransfer16 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Identification assigned to the transfer of asset, typically assigned by the transferee.
	TransferIdentification *Max35Text `xml:"TrfId"`

	// Identification of the confirmation assigned by the transferor to the transfer.
	TransferConfirmationIdentification *Max35Text `xml:"TrfConfId,omitempty"`

	// Requested date at which the assets should be transferred.
	RequestedTransferDate *DateFormat1Choice `xml:"ReqdTrfDt,omitempty"`

	// Specifies portfolio information or government schemes, for example Individual Savings Account (ISA) in the UK.
	Portfolio *ISAPortfolio1Choice `xml:"Prtfl,omitempty"`

	// Indicates whether there is cash in the account that is awaiting investment.
	ResidualCash *ResidualCash1Code `xml:"RsdlCsh,omitempty"`

	// Indicator that all remaining assets in a portfolio not listed for transfer should be liquidated and transferred as cash.
	AllOtherCash *YesNoIndicator `xml:"AllOthrCsh"`

	// Specifies the underlying assets for the ISA or portfolio.
	FinancialInstrumentAssetForTransfer []*FinancialInstrument34 `xml:"FinInstrmAsstForTrf,omitempty"`
}

func (i *ISATransfer16) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *ISATransfer16) SetTransferIdentification(value string) {
	i.TransferIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer16) SetTransferConfirmationIdentification(value string) {
	i.TransferConfirmationIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer16) AddRequestedTransferDate() *DateFormat1Choice {
	i.RequestedTransferDate = new(DateFormat1Choice)
	return i.RequestedTransferDate
}

func (i *ISATransfer16) AddPortfolio() *ISAPortfolio1Choice {
	i.Portfolio = new(ISAPortfolio1Choice)
	return i.Portfolio
}

func (i *ISATransfer16) SetResidualCash(value string) {
	i.ResidualCash = (*ResidualCash1Code)(&value)
}

func (i *ISATransfer16) SetAllOtherCash(value string) {
	i.AllOtherCash = (*YesNoIndicator)(&value)
}

func (i *ISATransfer16) AddFinancialInstrumentAssetForTransfer() *FinancialInstrument34 {
	newValue := new(FinancialInstrument34)
	i.FinancialInstrumentAssetForTransfer = append(i.FinancialInstrumentAssetForTransfer, newValue)
	return newValue
}

// Describes the type of product and the assets to be transferred.
type ISATransfer17 struct {

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Provides information related to the asset(s) transferred.
	ProductTransfer []*ISATransfer16 `xml:"PdctTrf"`
}

func (i *ISATransfer17) SetCancellationReference(value string) {
	i.CancellationReference = (*Max35Text)(&value)
}

func (i *ISATransfer17) AddProductTransfer() *ISATransfer16 {
	newValue := new(ISATransfer16)
	i.ProductTransfer = append(i.ProductTransfer, newValue)
	return newValue
}

// Describes the type of product and the assets to be transferred.
type ISATransfer18 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Identification assigned to the transfer of asset, typically assigned by the transferee.
	TransferIdentification *Max35Text `xml:"TrfId"`

	// Identification of the confirmation assigned by the transferor to the transfer.
	TransferConfirmationIdentification *Max35Text `xml:"TrfConfId,omitempty"`

	// Requested date at which the assets should be transferred.
	RequestedTransferDate *DateFormat1Choice `xml:"ReqdTrfDt,omitempty"`

	// Specifies portfolio information or government schemes, for example Individual Savings Account (ISA) in the UK.
	Portfolio *ISAPortfolio1Choice `xml:"Prtfl,omitempty"`

	// Indicates whether there is cash in the account that is awaiting investment.
	ResidualCash *ResidualCash1Code `xml:"RsdlCsh,omitempty"`

	// Indicator that all remaining assets in a portfolio not listed for transfer should be liquidated and transferred as cash.
	AllOtherCash *YesNoIndicator `xml:"AllOthrCsh"`

	// Specifies the underlying assets for the ISA or portfolio.
	FinancialInstrumentAssetForTransfer []*FinancialInstrument39 `xml:"FinInstrmAsstForTrf,omitempty"`
}

func (i *ISATransfer18) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *ISATransfer18) SetTransferIdentification(value string) {
	i.TransferIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer18) SetTransferConfirmationIdentification(value string) {
	i.TransferConfirmationIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer18) AddRequestedTransferDate() *DateFormat1Choice {
	i.RequestedTransferDate = new(DateFormat1Choice)
	return i.RequestedTransferDate
}

func (i *ISATransfer18) AddPortfolio() *ISAPortfolio1Choice {
	i.Portfolio = new(ISAPortfolio1Choice)
	return i.Portfolio
}

func (i *ISATransfer18) SetResidualCash(value string) {
	i.ResidualCash = (*ResidualCash1Code)(&value)
}

func (i *ISATransfer18) SetAllOtherCash(value string) {
	i.AllOtherCash = (*YesNoIndicator)(&value)
}

func (i *ISATransfer18) AddFinancialInstrumentAssetForTransfer() *FinancialInstrument39 {
	newValue := new(FinancialInstrument39)
	i.FinancialInstrumentAssetForTransfer = append(i.FinancialInstrumentAssetForTransfer, newValue)
	return newValue
}

// Describes the type of product and the assets to be transferred.
type ISATransfer19 struct {

	// Information identifying the primary individual investor, for example, name, address, social security number and date of birth.
	PrimaryIndividualInvestor *IndividualPerson8 `xml:"PmryIndvInvstr,omitempty"`

	// Information identifying the secondary individual investor, for example, name, address, social security number and date of birth.
	SecondaryIndividualInvestor *IndividualPerson8 `xml:"ScndryIndvInvstr,omitempty"`

	// Information identifying the other individual investors, for example, name, address, social security number and date of birth.
	OtherIndividualInvestor []*IndividualPerson8 `xml:"OthrIndvInvstr,omitempty"`

	// Information identifying the primary corporate investor, for example, name and address.
	PrimaryCorporateInvestor *Organisation4 `xml:"PmryCorpInvstr,omitempty"`

	// Information identifying the secondary corporate investor, for example, name and address.
	SecondaryCorporateInvestor *Organisation4 `xml:"ScndryCorpInvstr,omitempty"`

	// Information identifying the other corporate investors, for example, name and address.
	OtherCorporateInvestor []*Organisation4 `xml:"OthrCorpInvstr,omitempty"`

	// Identification of an account owned by the investor at the old plan manager (account servicer).
	TransferorAccount *Account15 `xml:"TrfrAcct"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	NomineeAccount *Account16 `xml:"NmneeAcct,omitempty"`

	// Information related to the institution to which the financial instrument is to be transferred.
	Transferee *PartyIdentification2Choice `xml:"Trfee"`

	// Identification of an account owned by the investor to which a cash entry is made based on the transfer of asset(s).
	CashAccount *CashAccount29 `xml:"CshAcct,omitempty"`

	// Details of the transfer to be cancelled.
	ProductTransferAndReference *ISATransfer20 `xml:"PdctTrfAndRef"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (i *ISATransfer19) AddPrimaryIndividualInvestor() *IndividualPerson8 {
	i.PrimaryIndividualInvestor = new(IndividualPerson8)
	return i.PrimaryIndividualInvestor
}

func (i *ISATransfer19) AddSecondaryIndividualInvestor() *IndividualPerson8 {
	i.SecondaryIndividualInvestor = new(IndividualPerson8)
	return i.SecondaryIndividualInvestor
}

func (i *ISATransfer19) AddOtherIndividualInvestor() *IndividualPerson8 {
	newValue := new(IndividualPerson8)
	i.OtherIndividualInvestor = append(i.OtherIndividualInvestor, newValue)
	return newValue
}

func (i *ISATransfer19) AddPrimaryCorporateInvestor() *Organisation4 {
	i.PrimaryCorporateInvestor = new(Organisation4)
	return i.PrimaryCorporateInvestor
}

func (i *ISATransfer19) AddSecondaryCorporateInvestor() *Organisation4 {
	i.SecondaryCorporateInvestor = new(Organisation4)
	return i.SecondaryCorporateInvestor
}

func (i *ISATransfer19) AddOtherCorporateInvestor() *Organisation4 {
	newValue := new(Organisation4)
	i.OtherCorporateInvestor = append(i.OtherCorporateInvestor, newValue)
	return newValue
}

func (i *ISATransfer19) AddTransferorAccount() *Account15 {
	i.TransferorAccount = new(Account15)
	return i.TransferorAccount
}

func (i *ISATransfer19) AddNomineeAccount() *Account16 {
	i.NomineeAccount = new(Account16)
	return i.NomineeAccount
}

func (i *ISATransfer19) AddTransferee() *PartyIdentification2Choice {
	i.Transferee = new(PartyIdentification2Choice)
	return i.Transferee
}

func (i *ISATransfer19) AddCashAccount() *CashAccount29 {
	i.CashAccount = new(CashAccount29)
	return i.CashAccount
}

func (i *ISATransfer19) AddProductTransferAndReference() *ISATransfer20 {
	i.ProductTransferAndReference = new(ISATransfer20)
	return i.ProductTransferAndReference
}

func (i *ISATransfer19) AddExtension() *Extension1 {
	newValue := new(Extension1)
	i.Extension = append(i.Extension, newValue)
	return newValue
}

// Describes the type of product and the assets to be transferred.
type ISATransfer2 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Identification of the confirmation assigned by the old plan manager to the transfer of account.
	TransferConfirmationIdentification *Max35Text `xml:"TrfConfId"`

	// Identification received by the old plan manager and assigned by the new plan manager to the transfer of account.
	TransferInstructionReference *Max35Text `xml:"TrfInstrRef"`

	// Date when the transfer instruction was executed.
	ActualTransferDate *ISODate `xml:"ActlTrfDt"`

	// Indicates whether there is cash in the account that is awaiting investment.
	ResidualCash *ResidualCash1Code `xml:"RsdlCsh,omitempty"`

	// Specifies portfolio information or government schemes, for example Individual Savings Account (ISA) in the UK.
	Portfolio *ISAPortfolio2Choice `xml:"Prtfl"`

	// Indicator that all remaining assets in a portfolio not listed for transfer should be liquidated and transferred as cash.
	AllOtherCash *YesNoIndicator `xml:"AllOthrCsh"`

	// Specifies the underlying assets for the ISA or portfolio.
	FinancialInstrumentAssetForTransfer []*FinancialInstrument24 `xml:"FinInstrmAsstForTrf,omitempty"`
}

func (i *ISATransfer2) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *ISATransfer2) SetTransferConfirmationIdentification(value string) {
	i.TransferConfirmationIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer2) SetTransferInstructionReference(value string) {
	i.TransferInstructionReference = (*Max35Text)(&value)
}

func (i *ISATransfer2) SetActualTransferDate(value string) {
	i.ActualTransferDate = (*ISODate)(&value)
}

func (i *ISATransfer2) SetResidualCash(value string) {
	i.ResidualCash = (*ResidualCash1Code)(&value)
}

func (i *ISATransfer2) AddPortfolio() *ISAPortfolio2Choice {
	i.Portfolio = new(ISAPortfolio2Choice)
	return i.Portfolio
}

func (i *ISATransfer2) SetAllOtherCash(value string) {
	i.AllOtherCash = (*YesNoIndicator)(&value)
}

func (i *ISATransfer2) AddFinancialInstrumentAssetForTransfer() *FinancialInstrument24 {
	newValue := new(FinancialInstrument24)
	i.FinancialInstrumentAssetForTransfer = append(i.FinancialInstrumentAssetForTransfer, newValue)
	return newValue
}

// Describes the type of product and the assets to be transferred.
type ISATransfer20 struct {

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Provides information related to the asset(s) transferred.
	ProductTransfer []*ISATransfer18 `xml:"PdctTrf"`
}

func (i *ISATransfer20) SetCancellationReference(value string) {
	i.CancellationReference = (*Max35Text)(&value)
}

func (i *ISATransfer20) AddProductTransfer() *ISATransfer18 {
	newValue := new(ISATransfer18)
	i.ProductTransfer = append(i.ProductTransfer, newValue)
	return newValue
}

// Describes the type of product and the assets to be transferred.
type ISATransfer21 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Identification of the confirmation assigned by the transferor to the transfer.
	TransferConfirmationIdentification *Max35Text `xml:"TrfConfId"`

	// Identification assigned to the transfer of asset, typically assigned by the transferee.
	TransferInstructionReference *Max35Text `xml:"TrfInstrRef"`

	// Date when the transfer instruction was executed.
	ActualTransferDate *ISODate `xml:"ActlTrfDt,omitempty"`

	// Indicates whether there is cash in the account that is awaiting investment.
	ResidualCash *ResidualCash1Code `xml:"RsdlCsh,omitempty"`

	// Specifies portfolio information or government schemes, for example Individual Savings Account (ISA) in the UK.
	Portfolio *ISAPortfolio2Choice `xml:"Prtfl,omitempty"`

	// Indicator that all remaining assets in a portfolio not listed for transfer should be liquidated and transferred as cash.
	AllOtherCash *YesNoIndicator `xml:"AllOthrCsh"`

	// Specifies the underlying assets for the ISA or portfolio.
	FinancialInstrumentAssetForTransfer []*FinancialInstrument40 `xml:"FinInstrmAsstForTrf,omitempty"`
}

func (i *ISATransfer21) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *ISATransfer21) SetTransferConfirmationIdentification(value string) {
	i.TransferConfirmationIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer21) SetTransferInstructionReference(value string) {
	i.TransferInstructionReference = (*Max35Text)(&value)
}

func (i *ISATransfer21) SetActualTransferDate(value string) {
	i.ActualTransferDate = (*ISODate)(&value)
}

func (i *ISATransfer21) SetResidualCash(value string) {
	i.ResidualCash = (*ResidualCash1Code)(&value)
}

func (i *ISATransfer21) AddPortfolio() *ISAPortfolio2Choice {
	i.Portfolio = new(ISAPortfolio2Choice)
	return i.Portfolio
}

func (i *ISATransfer21) SetAllOtherCash(value string) {
	i.AllOtherCash = (*YesNoIndicator)(&value)
}

func (i *ISATransfer21) AddFinancialInstrumentAssetForTransfer() *FinancialInstrument40 {
	newValue := new(FinancialInstrument40)
	i.FinancialInstrumentAssetForTransfer = append(i.FinancialInstrumentAssetForTransfer, newValue)
	return newValue
}

// Describes the type of product and the assets to be transferred.
type ISATransfer22 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Identification assigned to the transfer of asset, typically assigned by the transferee.
	TransferIdentification *Max35Text `xml:"TrfId"`

	// Identification of the confirmation assigned by the transferor to the transfer.
	TransferConfirmationIdentification *Max35Text `xml:"TrfConfId,omitempty"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *AdditionalReference7 `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transfer allocated by the counterparty.
	CounterpartyReference *AdditionalReference7 `xml:"CtrPtyRef,omitempty"`

	// Identifies the business process in which the actors are involved. This is important to trigger the right business process, according to the market business model, which may require matching instructions in a CSD environment (double leg process) or not (single leg process).
	BusinessFlowType *BusinessFlowType1Code `xml:"BizFlowTp,omitempty"`

	// Requested date at which the assets should be transferred.
	RequestedTransferDate *DateFormat1Choice `xml:"ReqdTrfDt,omitempty"`

	// Specifies portfolio information or government schemes, for example Individual Savings Account (ISA) in the UK.
	Portfolio *ISAPortfolio1Choice `xml:"Prtfl,omitempty"`

	// Indicates whether there is cash in the account that is awaiting investment.
	ResidualCash *ResidualCash1Code `xml:"RsdlCsh,omitempty"`

	// Specifies whether all remaining assets in a portfolio not listed for transfer should be liquidated and transferred as cash.
	AllOtherCash *AllOtherCash1Code `xml:"AllOthrCsh,omitempty"`

	// Specifies the underlying assets for the ISA or portfolio.
	FinancialInstrumentAssetForTransfer []*FinancialInstrument46 `xml:"FinInstrmAsstForTrf,omitempty"`
}

func (i *ISATransfer22) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *ISATransfer22) SetTransferIdentification(value string) {
	i.TransferIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer22) SetTransferConfirmationIdentification(value string) {
	i.TransferConfirmationIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer22) AddClientReference() *AdditionalReference7 {
	i.ClientReference = new(AdditionalReference7)
	return i.ClientReference
}

func (i *ISATransfer22) AddCounterpartyReference() *AdditionalReference7 {
	i.CounterpartyReference = new(AdditionalReference7)
	return i.CounterpartyReference
}

func (i *ISATransfer22) SetBusinessFlowType(value string) {
	i.BusinessFlowType = (*BusinessFlowType1Code)(&value)
}

func (i *ISATransfer22) AddRequestedTransferDate() *DateFormat1Choice {
	i.RequestedTransferDate = new(DateFormat1Choice)
	return i.RequestedTransferDate
}

func (i *ISATransfer22) AddPortfolio() *ISAPortfolio1Choice {
	i.Portfolio = new(ISAPortfolio1Choice)
	return i.Portfolio
}

func (i *ISATransfer22) SetResidualCash(value string) {
	i.ResidualCash = (*ResidualCash1Code)(&value)
}

func (i *ISATransfer22) SetAllOtherCash(value string) {
	i.AllOtherCash = (*AllOtherCash1Code)(&value)
}

func (i *ISATransfer22) AddFinancialInstrumentAssetForTransfer() *FinancialInstrument46 {
	newValue := new(FinancialInstrument46)
	i.FinancialInstrumentAssetForTransfer = append(i.FinancialInstrumentAssetForTransfer, newValue)
	return newValue
}

// Describes the type of product and the assets to be transferred.
type ISATransfer23 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Identification assigned to the transfer of asset, typically assigned by the transferee.
	TransferIdentification *Max35Text `xml:"TrfId"`

	// Identification of the confirmation assigned by the transferor to the transfer.
	TransferConfirmationIdentification *Max35Text `xml:"TrfConfId,omitempty"`

	// Indicates whether there is cash in the account that is awaiting investment.
	ResidualCash *ResidualCash1Code `xml:"RsdlCsh,omitempty"`

	// Specifies portfolio information or government schemes, for example Individual Savings Account (ISA) in the UK.
	Portfolio *ISAPortfolio1Choice `xml:"Prtfl,omitempty"`

	// Specifies whether all remaining assets in a portfolio not listed for transfer should be liquidated and transferred as cash.
	AllOtherCash *AllOtherCash1Code `xml:"AllOthrCsh,omitempty"`

	// Specifies the underlying assets for the ISA or portfolio.
	FinancialInstrumentAssetForTransfer []*FinancialInstrument47 `xml:"FinInstrmAsstForTrf,omitempty"`
}

func (i *ISATransfer23) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *ISATransfer23) SetTransferIdentification(value string) {
	i.TransferIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer23) SetTransferConfirmationIdentification(value string) {
	i.TransferConfirmationIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer23) SetResidualCash(value string) {
	i.ResidualCash = (*ResidualCash1Code)(&value)
}

func (i *ISATransfer23) AddPortfolio() *ISAPortfolio1Choice {
	i.Portfolio = new(ISAPortfolio1Choice)
	return i.Portfolio
}

func (i *ISATransfer23) SetAllOtherCash(value string) {
	i.AllOtherCash = (*AllOtherCash1Code)(&value)
}

func (i *ISATransfer23) AddFinancialInstrumentAssetForTransfer() *FinancialInstrument47 {
	newValue := new(FinancialInstrument47)
	i.FinancialInstrumentAssetForTransfer = append(i.FinancialInstrumentAssetForTransfer, newValue)
	return newValue
}

// Describes the type of product and the assets to be transferred.
type ISATransfer24 struct {

	// Information identifying the primary individual investor, for example, name, address, social security number and date of birth.
	PrimaryIndividualInvestor *IndividualPerson8 `xml:"PmryIndvInvstr,omitempty"`

	// Information identifying the secondary individual investor, for example, name, address, social security number and date of birth.
	SecondaryIndividualInvestor *IndividualPerson8 `xml:"ScndryIndvInvstr,omitempty"`

	// Information identifying the other individual investors, for example, name, address, social security number and date of birth.
	OtherIndividualInvestor []*IndividualPerson8 `xml:"OthrIndvInvstr,omitempty"`

	// Information identifying the primary corporate investor, for example, name and address.
	PrimaryCorporateInvestor *Organisation21 `xml:"PmryCorpInvstr,omitempty"`

	// Information identifying the secondary corporate investor, for example, name and address.
	SecondaryCorporateInvestor *Organisation21 `xml:"ScndryCorpInvstr,omitempty"`

	// Information identifying the other corporate investors, for example, name and address.
	OtherCorporateInvestor []*Organisation21 `xml:"OthrCorpInvstr,omitempty"`

	// Identification of an account owned by the investor at the old plan manager (account servicer).
	TransferorAccount *Account19 `xml:"TrfrAcct"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	NomineeAccount *Account19 `xml:"NmneeAcct,omitempty"`

	// Information related to the institution to which the financial instrument is to be transferred.
	Transferee *PartyIdentification70Choice `xml:"Trfee"`

	// Identification of a related party or intermediary.
	IntermediaryInformation []*Intermediary34 `xml:"IntrmyInf,omitempty"`

	// Identification of an account owned by the investor to which a cash entry is made based on the transfer of asset(s).
	CashAccount *CashAccount34 `xml:"CshAcct,omitempty"`

	// Details of the transfer to be cancelled.
	ProductTransferAndReference *ISATransfer25 `xml:"PdctTrfAndRef"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (i *ISATransfer24) AddPrimaryIndividualInvestor() *IndividualPerson8 {
	i.PrimaryIndividualInvestor = new(IndividualPerson8)
	return i.PrimaryIndividualInvestor
}

func (i *ISATransfer24) AddSecondaryIndividualInvestor() *IndividualPerson8 {
	i.SecondaryIndividualInvestor = new(IndividualPerson8)
	return i.SecondaryIndividualInvestor
}

func (i *ISATransfer24) AddOtherIndividualInvestor() *IndividualPerson8 {
	newValue := new(IndividualPerson8)
	i.OtherIndividualInvestor = append(i.OtherIndividualInvestor, newValue)
	return newValue
}

func (i *ISATransfer24) AddPrimaryCorporateInvestor() *Organisation21 {
	i.PrimaryCorporateInvestor = new(Organisation21)
	return i.PrimaryCorporateInvestor
}

func (i *ISATransfer24) AddSecondaryCorporateInvestor() *Organisation21 {
	i.SecondaryCorporateInvestor = new(Organisation21)
	return i.SecondaryCorporateInvestor
}

func (i *ISATransfer24) AddOtherCorporateInvestor() *Organisation21 {
	newValue := new(Organisation21)
	i.OtherCorporateInvestor = append(i.OtherCorporateInvestor, newValue)
	return newValue
}

func (i *ISATransfer24) AddTransferorAccount() *Account19 {
	i.TransferorAccount = new(Account19)
	return i.TransferorAccount
}

func (i *ISATransfer24) AddNomineeAccount() *Account19 {
	i.NomineeAccount = new(Account19)
	return i.NomineeAccount
}

func (i *ISATransfer24) AddTransferee() *PartyIdentification70Choice {
	i.Transferee = new(PartyIdentification70Choice)
	return i.Transferee
}

func (i *ISATransfer24) AddIntermediaryInformation() *Intermediary34 {
	newValue := new(Intermediary34)
	i.IntermediaryInformation = append(i.IntermediaryInformation, newValue)
	return newValue
}

func (i *ISATransfer24) AddCashAccount() *CashAccount34 {
	i.CashAccount = new(CashAccount34)
	return i.CashAccount
}

func (i *ISATransfer24) AddProductTransferAndReference() *ISATransfer25 {
	i.ProductTransferAndReference = new(ISATransfer25)
	return i.ProductTransferAndReference
}

func (i *ISATransfer24) AddExtension() *Extension1 {
	newValue := new(Extension1)
	i.Extension = append(i.Extension, newValue)
	return newValue
}

// Describes the type of product and the assets to be transferred.
type ISATransfer25 struct {

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Provides information related to the asset(s) transferred.
	ProductTransfer []*ISATransfer22 `xml:"PdctTrf"`
}

func (i *ISATransfer25) SetCancellationReference(value string) {
	i.CancellationReference = (*Max35Text)(&value)
}

func (i *ISATransfer25) AddProductTransfer() *ISATransfer22 {
	newValue := new(ISATransfer22)
	i.ProductTransfer = append(i.ProductTransfer, newValue)
	return newValue
}

// Describes the type of product and the assets to be transferred.
type ISATransfer26 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Identification of the confirmation assigned by the transferor to the transfer.
	TransferConfirmationIdentification *Max35Text `xml:"TrfConfId"`

	// Identification assigned to the transfer of asset, typically assigned by the transferee.
	TransferInstructionReference *Max35Text `xml:"TrfInstrRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *AdditionalReference7 `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transfer allocated by the counterparty.
	CounterpartyReference *AdditionalReference7 `xml:"CtrPtyRef,omitempty"`

	// Identifies the business process in which the actors are involved. This is important to trigger the right business process, according to the market business model, which may require matching instructions in a CSD environment (double leg process) or not (single leg process).
	BusinessFlowType *BusinessFlowType1Code `xml:"BizFlowTp,omitempty"`

	// Date when the transfer instruction was executed.
	ActualTransferDate *ISODate `xml:"ActlTrfDt,omitempty"`

	// Indicates whether there is cash in the account that is awaiting investment.
	ResidualCash *ResidualCash1Code `xml:"RsdlCsh,omitempty"`

	// Specifies portfolio information or government schemes, for example Individual Savings Account (ISA) in the UK.
	Portfolio *ISAPortfolio2Choice `xml:"Prtfl,omitempty"`

	// Specifies whether all remaining assets in a portfolio not listed for transfer should be liquidated and transferred as cash.
	AllOtherCash *AllOtherCash1Code `xml:"AllOthrCsh,omitempty"`

	// Specifies the underlying assets for the ISA or portfolio.
	FinancialInstrumentAssetForTransfer []*FinancialInstrument48 `xml:"FinInstrmAsstForTrf,omitempty"`
}

func (i *ISATransfer26) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *ISATransfer26) SetTransferConfirmationIdentification(value string) {
	i.TransferConfirmationIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer26) SetTransferInstructionReference(value string) {
	i.TransferInstructionReference = (*Max35Text)(&value)
}

func (i *ISATransfer26) AddClientReference() *AdditionalReference7 {
	i.ClientReference = new(AdditionalReference7)
	return i.ClientReference
}

func (i *ISATransfer26) AddCounterpartyReference() *AdditionalReference7 {
	i.CounterpartyReference = new(AdditionalReference7)
	return i.CounterpartyReference
}

func (i *ISATransfer26) SetBusinessFlowType(value string) {
	i.BusinessFlowType = (*BusinessFlowType1Code)(&value)
}

func (i *ISATransfer26) SetActualTransferDate(value string) {
	i.ActualTransferDate = (*ISODate)(&value)
}

func (i *ISATransfer26) SetResidualCash(value string) {
	i.ResidualCash = (*ResidualCash1Code)(&value)
}

func (i *ISATransfer26) AddPortfolio() *ISAPortfolio2Choice {
	i.Portfolio = new(ISAPortfolio2Choice)
	return i.Portfolio
}

func (i *ISATransfer26) SetAllOtherCash(value string) {
	i.AllOtherCash = (*AllOtherCash1Code)(&value)
}

func (i *ISATransfer26) AddFinancialInstrumentAssetForTransfer() *FinancialInstrument48 {
	newValue := new(FinancialInstrument48)
	i.FinancialInstrumentAssetForTransfer = append(i.FinancialInstrumentAssetForTransfer, newValue)
	return newValue
}

// Describes the type of product and the assets to be transferred.
type ISATransfer27 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Identification assigned to the transfer of asset, typically assigned by the transferee.
	TransferIdentification *Max35Text `xml:"TrfId"`

	// Specifies portfolio information or government schemes, for example Individual Savings Account (ISA) in the UK.
	Portfolio *ISAPortfolio3Choice `xml:"Prtfl,omitempty"`

	// Specifies whether all remaining assets in a portfolio not listed for transfer should be liquidated and transferred as cash.
	AllOtherCash *AllOtherCash1Code `xml:"AllOthrCsh,omitempty"`

	// Specifies the underlying assets for the ISA or portfolio.
	FinancialInstrumentAssetForTransfer []*FinancialInstrument50 `xml:"FinInstrmAsstForTrf,omitempty"`
}

func (i *ISATransfer27) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *ISATransfer27) SetTransferIdentification(value string) {
	i.TransferIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer27) AddPortfolio() *ISAPortfolio3Choice {
	i.Portfolio = new(ISAPortfolio3Choice)
	return i.Portfolio
}

func (i *ISATransfer27) SetAllOtherCash(value string) {
	i.AllOtherCash = (*AllOtherCash1Code)(&value)
}

func (i *ISATransfer27) AddFinancialInstrumentAssetForTransfer() *FinancialInstrument50 {
	newValue := new(FinancialInstrument50)
	i.FinancialInstrumentAssetForTransfer = append(i.FinancialInstrumentAssetForTransfer, newValue)
	return newValue
}

// Describes the type of product and the assets to be transferred.
type ISATransfer3 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Identification assigned by the new plan manager to each transfer of asset.
	TransferIdentification *Max35Text `xml:"TrfId"`

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Indicates whether there is cash in the account that is awaiting investment.
	ResidualCash *ResidualCash1Code `xml:"RsdlCsh,omitempty"`

	// Specifies portfolio information or government schemes, for example Individual Savings Account (ISA) in the UK.
	Portfolio *ISAPortfolio1Choice `xml:"Prtfl"`

	// Indicator that all remaining assets in a portfolio not listed for transfer should be liquidated and transferred as cash.
	AllOtherCash *YesNoIndicator `xml:"AllOthrCsh"`

	// Specifies the underlying assets for the ISA or portfolio.
	FinancialInstrumentAssetForTransfer []*FinancialInstrument25 `xml:"FinInstrmAsstForTrf,omitempty"`
}

func (i *ISATransfer3) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *ISATransfer3) SetTransferIdentification(value string) {
	i.TransferIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer3) SetCancellationReference(value string) {
	i.CancellationReference = (*Max35Text)(&value)
}

func (i *ISATransfer3) SetResidualCash(value string) {
	i.ResidualCash = (*ResidualCash1Code)(&value)
}

func (i *ISATransfer3) AddPortfolio() *ISAPortfolio1Choice {
	i.Portfolio = new(ISAPortfolio1Choice)
	return i.Portfolio
}

func (i *ISATransfer3) SetAllOtherCash(value string) {
	i.AllOtherCash = (*YesNoIndicator)(&value)
}

func (i *ISATransfer3) AddFinancialInstrumentAssetForTransfer() *FinancialInstrument25 {
	newValue := new(FinancialInstrument25)
	i.FinancialInstrumentAssetForTransfer = append(i.FinancialInstrumentAssetForTransfer, newValue)
	return newValue
}

// Describes the type of product and the assets to be transferred.
type ISATransfer4 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Identification assigned by the new plan manager to each transfer of asset.
	TransferIdentification *Max35Text `xml:"TrfId"`

	// Indicates whether there is cash in the account that is awaiting investment.
	ResidualCash *ResidualCash1Code `xml:"RsdlCsh,omitempty"`

	// Specifies portfolio information or government schemes, for example Individual Savings Account (ISA) in the UK.
	Portfolio *ISAPortfolio1Choice `xml:"Prtfl"`

	// Indicator that all remaining assets in a portfolio not listed for transfer should be liquidated and transferred as cash.
	AllOtherCash *YesNoIndicator `xml:"AllOthrCsh"`

	// Specifies the underlying assets for the ISA or portfolio.
	FinancialInstrumentAssetForTransfer []*FinancialInstrument26 `xml:"FinInstrmAsstForTrf,omitempty"`
}

func (i *ISATransfer4) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *ISATransfer4) SetTransferIdentification(value string) {
	i.TransferIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer4) SetResidualCash(value string) {
	i.ResidualCash = (*ResidualCash1Code)(&value)
}

func (i *ISATransfer4) AddPortfolio() *ISAPortfolio1Choice {
	i.Portfolio = new(ISAPortfolio1Choice)
	return i.Portfolio
}

func (i *ISATransfer4) SetAllOtherCash(value string) {
	i.AllOtherCash = (*YesNoIndicator)(&value)
}

func (i *ISATransfer4) AddFinancialInstrumentAssetForTransfer() *FinancialInstrument26 {
	newValue := new(FinancialInstrument26)
	i.FinancialInstrumentAssetForTransfer = append(i.FinancialInstrumentAssetForTransfer, newValue)
	return newValue
}

// Describes the type of product and the assets to be transferred.
type ISATransfer5 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Identification assigned by the new plan manager to each transfer of asset.
	TransferIdentification *Max35Text `xml:"TrfId"`

	// Specifies portfolio information or government schemes, for example Individual Savings Account (ISA) in the UK.
	Portfolio *ISAPortfolio3Choice `xml:"Prtfl"`

	// Indicator that all remaining assets in a portfolio not listed for transfer should be liquidated and transferred as cash.
	AllOtherCash *YesNoIndicator `xml:"AllOthrCsh"`

	// Specifies the underlying assets for the ISA or portfolio.
	FinancialInstrumentAssetForTransfer []*FinancialInstrument27 `xml:"FinInstrmAsstForTrf,omitempty"`
}

func (i *ISATransfer5) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *ISATransfer5) SetTransferIdentification(value string) {
	i.TransferIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer5) AddPortfolio() *ISAPortfolio3Choice {
	i.Portfolio = new(ISAPortfolio3Choice)
	return i.Portfolio
}

func (i *ISATransfer5) SetAllOtherCash(value string) {
	i.AllOtherCash = (*YesNoIndicator)(&value)
}

func (i *ISATransfer5) AddFinancialInstrumentAssetForTransfer() *FinancialInstrument27 {
	newValue := new(FinancialInstrument27)
	i.FinancialInstrumentAssetForTransfer = append(i.FinancialInstrumentAssetForTransfer, newValue)
	return newValue
}

// Describes the type of product and the assets to be transferred.
type ISATransfer6 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Identification assigned by the new plan manager to each transfer of asset.
	TransferIdentification *Max35Text `xml:"TrfId"`

	// Indicates whether there is cash in the account that is awaiting investment.
	ResidualCash *ResidualCash1Code `xml:"RsdlCsh,omitempty"`

	// Specifies portfolio information or government schemes, for example Individual Savings Account (ISA) in the UK.
	Portfolio *ISAPortfolio1Choice `xml:"Prtfl"`

	// Indicator that all remaining assets in a portfolio not listed for transfer should be liquidated and transferred as cash.
	AllOtherCash *YesNoIndicator `xml:"AllOthrCsh"`

	// Specifies the underlying assets for the ISA or portfolio.
	FinancialInstrumentAssetForTransfer []*FinancialInstrument30 `xml:"FinInstrmAsstForTrf,omitempty"`
}

func (i *ISATransfer6) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *ISATransfer6) SetTransferIdentification(value string) {
	i.TransferIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer6) SetResidualCash(value string) {
	i.ResidualCash = (*ResidualCash1Code)(&value)
}

func (i *ISATransfer6) AddPortfolio() *ISAPortfolio1Choice {
	i.Portfolio = new(ISAPortfolio1Choice)
	return i.Portfolio
}

func (i *ISATransfer6) SetAllOtherCash(value string) {
	i.AllOtherCash = (*YesNoIndicator)(&value)
}

func (i *ISATransfer6) AddFinancialInstrumentAssetForTransfer() *FinancialInstrument30 {
	newValue := new(FinancialInstrument30)
	i.FinancialInstrumentAssetForTransfer = append(i.FinancialInstrumentAssetForTransfer, newValue)
	return newValue
}

// Describes the type of product and the assets to be transferred.
type ISATransfer7 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Identification assigned by the new plan manager to each transfer of asset.
	TransferIdentification *Max35Text `xml:"TrfId"`

	// Requested date at which the assets should be transferred.
	RequestedTransferDate *DateFormat1Choice `xml:"ReqdTrfDt,omitempty"`

	// Specifies portfolio information or government schemes, for example Individual Savings Account (ISA) in the UK.
	Portfolio *ISAPortfolio1Choice `xml:"Prtfl"`

	// Indicates whether there is cash in the account that is awaiting investment.
	ResidualCash *ResidualCash1Code `xml:"RsdlCsh,omitempty"`

	// Indicator that all remaining assets in a portfolio not listed for transfer should be liquidated and transferred as cash.
	AllOtherCash *YesNoIndicator `xml:"AllOthrCsh"`

	// Specifies the underlying assets for the ISA or portfolio.
	FinancialInstrumentAssetForTransfer []*FinancialInstrument31 `xml:"FinInstrmAsstForTrf,omitempty"`
}

func (i *ISATransfer7) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *ISATransfer7) SetTransferIdentification(value string) {
	i.TransferIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer7) AddRequestedTransferDate() *DateFormat1Choice {
	i.RequestedTransferDate = new(DateFormat1Choice)
	return i.RequestedTransferDate
}

func (i *ISATransfer7) AddPortfolio() *ISAPortfolio1Choice {
	i.Portfolio = new(ISAPortfolio1Choice)
	return i.Portfolio
}

func (i *ISATransfer7) SetResidualCash(value string) {
	i.ResidualCash = (*ResidualCash1Code)(&value)
}

func (i *ISATransfer7) SetAllOtherCash(value string) {
	i.AllOtherCash = (*YesNoIndicator)(&value)
}

func (i *ISATransfer7) AddFinancialInstrumentAssetForTransfer() *FinancialInstrument31 {
	newValue := new(FinancialInstrument31)
	i.FinancialInstrumentAssetForTransfer = append(i.FinancialInstrumentAssetForTransfer, newValue)
	return newValue
}

// Describes the type of product and the assets to be transferred.
type ISATransfer8 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Identification assigned by the new plan manager to each transfer of asset.
	TransferIdentification *Max35Text `xml:"TrfId"`

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Indicates whether there is cash in the account that is awaiting investment.
	ResidualCash *ResidualCash1Code `xml:"RsdlCsh,omitempty"`

	// Specifies portfolio information or government schemes, for example Individual Savings Account (ISA) in the UK.
	Portfolio *ISAPortfolio1Choice `xml:"Prtfl"`

	// Indicator that all remaining assets in a portfolio not listed for transfer should be liquidated and transferred as cash.
	AllOtherCash *YesNoIndicator `xml:"AllOthrCsh"`

	// Specifies the underlying assets for the ISA or portfolio.
	FinancialInstrumentAssetForTransfer []*FinancialInstrument32 `xml:"FinInstrmAsstForTrf,omitempty"`
}

func (i *ISATransfer8) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *ISATransfer8) SetTransferIdentification(value string) {
	i.TransferIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer8) SetCancellationReference(value string) {
	i.CancellationReference = (*Max35Text)(&value)
}

func (i *ISATransfer8) SetResidualCash(value string) {
	i.ResidualCash = (*ResidualCash1Code)(&value)
}

func (i *ISATransfer8) AddPortfolio() *ISAPortfolio1Choice {
	i.Portfolio = new(ISAPortfolio1Choice)
	return i.Portfolio
}

func (i *ISATransfer8) SetAllOtherCash(value string) {
	i.AllOtherCash = (*YesNoIndicator)(&value)
}

func (i *ISATransfer8) AddFinancialInstrumentAssetForTransfer() *FinancialInstrument32 {
	newValue := new(FinancialInstrument32)
	i.FinancialInstrumentAssetForTransfer = append(i.FinancialInstrumentAssetForTransfer, newValue)
	return newValue
}

// Describes the type of product and the assets to be transferred.
type ISATransfer9 struct {

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
	ProductTransfer []*ISATransfer8 `xml:"PdctTrf"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (i *ISATransfer9) AddPrimaryIndividualInvestor() *IndividualPerson8 {
	i.PrimaryIndividualInvestor = new(IndividualPerson8)
	return i.PrimaryIndividualInvestor
}

func (i *ISATransfer9) AddSecondaryIndividualInvestor() *IndividualPerson8 {
	i.SecondaryIndividualInvestor = new(IndividualPerson8)
	return i.SecondaryIndividualInvestor
}

func (i *ISATransfer9) AddOtherIndividualInvestor() *IndividualPerson8 {
	newValue := new(IndividualPerson8)
	i.OtherIndividualInvestor = append(i.OtherIndividualInvestor, newValue)
	return newValue
}

func (i *ISATransfer9) AddPrimaryCorporateInvestor() *Organisation4 {
	i.PrimaryCorporateInvestor = new(Organisation4)
	return i.PrimaryCorporateInvestor
}

func (i *ISATransfer9) AddSecondaryCorporateInvestor() *Organisation4 {
	i.SecondaryCorporateInvestor = new(Organisation4)
	return i.SecondaryCorporateInvestor
}

func (i *ISATransfer9) AddOtherCorporateInvestor() *Organisation4 {
	newValue := new(Organisation4)
	i.OtherCorporateInvestor = append(i.OtherCorporateInvestor, newValue)
	return newValue
}

func (i *ISATransfer9) AddTransferorAccount() *Account5 {
	i.TransferorAccount = new(Account5)
	return i.TransferorAccount
}

func (i *ISATransfer9) AddNomineeAccount() *Account6 {
	i.NomineeAccount = new(Account6)
	return i.NomineeAccount
}

func (i *ISATransfer9) AddTransferee() *PartyIdentification2Choice {
	i.Transferee = new(PartyIdentification2Choice)
	return i.Transferee
}

func (i *ISATransfer9) AddCashAccount() *CashAccount11 {
	i.CashAccount = new(CashAccount11)
	return i.CashAccount
}

func (i *ISATransfer9) AddProductTransfer() *ISATransfer8 {
	newValue := new(ISATransfer8)
	i.ProductTransfer = append(i.ProductTransfer, newValue)
	return newValue
}

func (i *ISATransfer9) AddExtension() *Extension1 {
	newValue := new(Extension1)
	i.Extension = append(i.Extension, newValue)
	return newValue
}
