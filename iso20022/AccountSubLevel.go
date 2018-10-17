package iso20022

// Account and holding of the next sub-level (Level 1).
type AccountSubLevel11 struct {

	// Unique and unambiguous identification for the sub-account between the account owner and the account servicer.
	AccountIdentification *SecuritiesAccount19 `xml:"AcctId"`

	// Party that legally owns the sub-account.
	AccountOwner *PartyIdentification100 `xml:"AcctOwnr"`

	// Party that manages the sub-account on behalf of the account owner, that is, manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification100 `xml:"AcctSvcr"`

	// Individual or entity that is ultimately entitled to the benefit of income and rights in a financial instrument, as opposed to a nominal or legal owner.
	BeneficialOwner []*BeneficialOwner2 `xml:"BnfclOwnr,omitempty"`

	// Report on the net position of a financial instrument on the sub-account (sub-account level 1),  for a certain date. The agent, for example, a trade intermediary, may also be specified.
	BalanceDetails []*AggregateHoldingBalance3 `xml:"BalDtls,omitempty"`

	// Holdings of level 2.
	AccountSubLevel2 []*AccountSubLevel12 `xml:"AcctSubLvl2,omitempty"`

	// Difference in holdings between the sub-account at level 1 and sub-accounts of level 2.
	AccountSubLevel2Difference []*AggregateHoldingBalance2 `xml:"AcctSubLvl2Diff,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AccountSubLevel11) AddAccountIdentification() *SecuritiesAccount19 {
	a.AccountIdentification = new(SecuritiesAccount19)
	return a.AccountIdentification
}

func (a *AccountSubLevel11) AddAccountOwner() *PartyIdentification100 {
	a.AccountOwner = new(PartyIdentification100)
	return a.AccountOwner
}

func (a *AccountSubLevel11) AddAccountServicer() *PartyIdentification100 {
	a.AccountServicer = new(PartyIdentification100)
	return a.AccountServicer
}

func (a *AccountSubLevel11) AddBeneficialOwner() *BeneficialOwner2 {
	newValue := new(BeneficialOwner2)
	a.BeneficialOwner = append(a.BeneficialOwner, newValue)
	return newValue
}

func (a *AccountSubLevel11) AddBalanceDetails() *AggregateHoldingBalance3 {
	newValue := new(AggregateHoldingBalance3)
	a.BalanceDetails = append(a.BalanceDetails, newValue)
	return newValue
}

func (a *AccountSubLevel11) AddAccountSubLevel2() *AccountSubLevel12 {
	newValue := new(AccountSubLevel12)
	a.AccountSubLevel2 = append(a.AccountSubLevel2, newValue)
	return newValue
}

func (a *AccountSubLevel11) AddAccountSubLevel2Difference() *AggregateHoldingBalance2 {
	newValue := new(AggregateHoldingBalance2)
	a.AccountSubLevel2Difference = append(a.AccountSubLevel2Difference, newValue)
	return newValue
}

func (a *AccountSubLevel11) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}

// Account and holding of the next sub-level (Level 2).
type AccountSubLevel12 struct {

	// Unique and unambiguous identification for the sub-account between the account owner and the account servicer.
	AccountIdentification *SecuritiesAccount19 `xml:"AcctId"`

	// Party that legally owns the sub-account.
	AccountOwner *PartyIdentification100 `xml:"AcctOwnr"`

	// Party that manages the sub-account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification100 `xml:"AcctSvcr"`

	// Individual or entity that is ultimately entitled to the benefit of income and rights in a financial instrument, as opposed to a nominal or legal owner.
	BeneficialOwner []*BeneficialOwner2 `xml:"BnfclOwnr,omitempty"`

	// Report on the net position of a financial instrument on the sub-account (sub-account level 2),  for a certain date. The agent, for example, a trade intermediary, may also be specified.
	BalanceDetails []*AggregateHoldingBalance3 `xml:"BalDtls,omitempty"`

	// Holdings of level 3.
	AccountSubLevel3 []*AccountSubLevel13 `xml:"AcctSubLvl3,omitempty"`

	// Difference in holdings between the sub-account at level 2 and the sub-accounts of level 3.
	AccountSubLevel3Difference []*AggregateHoldingBalance2 `xml:"AcctSubLvl3Diff,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AccountSubLevel12) AddAccountIdentification() *SecuritiesAccount19 {
	a.AccountIdentification = new(SecuritiesAccount19)
	return a.AccountIdentification
}

func (a *AccountSubLevel12) AddAccountOwner() *PartyIdentification100 {
	a.AccountOwner = new(PartyIdentification100)
	return a.AccountOwner
}

func (a *AccountSubLevel12) AddAccountServicer() *PartyIdentification100 {
	a.AccountServicer = new(PartyIdentification100)
	return a.AccountServicer
}

func (a *AccountSubLevel12) AddBeneficialOwner() *BeneficialOwner2 {
	newValue := new(BeneficialOwner2)
	a.BeneficialOwner = append(a.BeneficialOwner, newValue)
	return newValue
}

func (a *AccountSubLevel12) AddBalanceDetails() *AggregateHoldingBalance3 {
	newValue := new(AggregateHoldingBalance3)
	a.BalanceDetails = append(a.BalanceDetails, newValue)
	return newValue
}

func (a *AccountSubLevel12) AddAccountSubLevel3() *AccountSubLevel13 {
	newValue := new(AccountSubLevel13)
	a.AccountSubLevel3 = append(a.AccountSubLevel3, newValue)
	return newValue
}

func (a *AccountSubLevel12) AddAccountSubLevel3Difference() *AggregateHoldingBalance2 {
	newValue := new(AggregateHoldingBalance2)
	a.AccountSubLevel3Difference = append(a.AccountSubLevel3Difference, newValue)
	return newValue
}

func (a *AccountSubLevel12) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}

// Account and holding of the next sub-level (Level 3).
type AccountSubLevel13 struct {

	// Unique and unambiguous identification for the sub-account between the account owner and the account servicer.
	AccountIdentification *SecuritiesAccount19 `xml:"AcctId"`

	// Party that legally owns the sub-account.
	AccountOwner *PartyIdentification100 `xml:"AcctOwnr"`

	// Party that manages the sub-level account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification100 `xml:"AcctSvcr"`

	// Individual or entity that is ultimately entitled to the benefit of income and rights in a financial instrument, as opposed to a nominal or legal owner.
	BeneficialOwner []*BeneficialOwner2 `xml:"BnfclOwnr,omitempty"`

	// Report on the net position of a financial instrument on the sub-account (sub-account level 3),  for a certain date. The agent, for example, a trade intermediary, may also be specified.
	BalanceDetails []*AggregateHoldingBalance3 `xml:"BalDtls,omitempty"`

	// Holdings of level 4.
	AccountSubLevel4 []*AccountSubLevel14 `xml:"AcctSubLvl4,omitempty"`

	// Difference in holdings between the sub-account at level 3 and the sub-accounts of level 4.
	AccountSubLevel4Difference []*AggregateHoldingBalance2 `xml:"AcctSubLvl4Diff,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AccountSubLevel13) AddAccountIdentification() *SecuritiesAccount19 {
	a.AccountIdentification = new(SecuritiesAccount19)
	return a.AccountIdentification
}

func (a *AccountSubLevel13) AddAccountOwner() *PartyIdentification100 {
	a.AccountOwner = new(PartyIdentification100)
	return a.AccountOwner
}

func (a *AccountSubLevel13) AddAccountServicer() *PartyIdentification100 {
	a.AccountServicer = new(PartyIdentification100)
	return a.AccountServicer
}

func (a *AccountSubLevel13) AddBeneficialOwner() *BeneficialOwner2 {
	newValue := new(BeneficialOwner2)
	a.BeneficialOwner = append(a.BeneficialOwner, newValue)
	return newValue
}

func (a *AccountSubLevel13) AddBalanceDetails() *AggregateHoldingBalance3 {
	newValue := new(AggregateHoldingBalance3)
	a.BalanceDetails = append(a.BalanceDetails, newValue)
	return newValue
}

func (a *AccountSubLevel13) AddAccountSubLevel4() *AccountSubLevel14 {
	newValue := new(AccountSubLevel14)
	a.AccountSubLevel4 = append(a.AccountSubLevel4, newValue)
	return newValue
}

func (a *AccountSubLevel13) AddAccountSubLevel4Difference() *AggregateHoldingBalance2 {
	newValue := new(AggregateHoldingBalance2)
	a.AccountSubLevel4Difference = append(a.AccountSubLevel4Difference, newValue)
	return newValue
}

func (a *AccountSubLevel13) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}

// Account and holding of the next sub-level (Level 4).
type AccountSubLevel14 struct {

	// Unique and unambiguous identification for the sub-account between the account owner and the account servicer.
	AccountIdentification *SecuritiesAccount19 `xml:"AcctId"`

	// Party that legally owns the sub-account.
	AccountOwner *PartyIdentification100 `xml:"AcctOwnr"`

	// Party that manages the sub-level account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification100 `xml:"AcctSvcr"`

	// Individual or entity that is ultimately entitled to the benefit of income and rights in a financial instrument, as opposed to a nominal or legal owner.
	BeneficialOwner []*BeneficialOwner2 `xml:"BnfclOwnr,omitempty"`

	// Report on the net position of a financial instrument on the sub-account (sub-account level 4),  for a certain date. The agent, for example, a trade intermediary, may also be specified.
	BalanceDetails []*AggregateHoldingBalance3 `xml:"BalDtls,omitempty"`

	// Holdings of level 5.
	AccountSubLevel5 []*AccountSubLevel15 `xml:"AcctSubLvl5,omitempty"`

	// Difference in holdings between the sub-account at level 4 and the sub-accounts of level 5.
	AccountSubLevel5Difference []*AggregateHoldingBalance2 `xml:"AcctSubLvl5Diff,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AccountSubLevel14) AddAccountIdentification() *SecuritiesAccount19 {
	a.AccountIdentification = new(SecuritiesAccount19)
	return a.AccountIdentification
}

func (a *AccountSubLevel14) AddAccountOwner() *PartyIdentification100 {
	a.AccountOwner = new(PartyIdentification100)
	return a.AccountOwner
}

func (a *AccountSubLevel14) AddAccountServicer() *PartyIdentification100 {
	a.AccountServicer = new(PartyIdentification100)
	return a.AccountServicer
}

func (a *AccountSubLevel14) AddBeneficialOwner() *BeneficialOwner2 {
	newValue := new(BeneficialOwner2)
	a.BeneficialOwner = append(a.BeneficialOwner, newValue)
	return newValue
}

func (a *AccountSubLevel14) AddBalanceDetails() *AggregateHoldingBalance3 {
	newValue := new(AggregateHoldingBalance3)
	a.BalanceDetails = append(a.BalanceDetails, newValue)
	return newValue
}

func (a *AccountSubLevel14) AddAccountSubLevel5() *AccountSubLevel15 {
	newValue := new(AccountSubLevel15)
	a.AccountSubLevel5 = append(a.AccountSubLevel5, newValue)
	return newValue
}

func (a *AccountSubLevel14) AddAccountSubLevel5Difference() *AggregateHoldingBalance2 {
	newValue := new(AggregateHoldingBalance2)
	a.AccountSubLevel5Difference = append(a.AccountSubLevel5Difference, newValue)
	return newValue
}

func (a *AccountSubLevel14) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}

// Account and holding of the next sub-level (Level 5).
type AccountSubLevel15 struct {

	// Unique and unambiguous identification for the sub-account between the account owner and the account servicer.
	AccountIdentification *SecuritiesAccount19 `xml:"AcctId"`

	// Party that legally owns the sub-account.
	AccountOwner *PartyIdentification100 `xml:"AcctOwnr"`

	// Party that manages the sub-level account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification100 `xml:"AcctSvcr"`

	// Individual or entity that is ultimately entitled to the benefit of income and rights in a financial instrument, as opposed to a nominal or legal owner.
	BeneficialOwner []*BeneficialOwner2 `xml:"BnfclOwnr,omitempty"`

	// Report on the net position of a financial instrument on the sub-account (sub-account level 5),  for a certain date. The agent, for example, a trade intermediary, may also be specified.
	BalanceDetails []*AggregateHoldingBalance3 `xml:"BalDtls,omitempty"`

	// Holdings of level 5.
	AccountSubLevel6 []*AccountSubLevel16 `xml:"AcctSubLvl6,omitempty"`

	// Difference in holdings between the safekeeping account and the sub-accounts of level 6.
	AccountSubLevel6Difference []*AggregateHoldingBalance2 `xml:"AcctSubLvl6Diff,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AccountSubLevel15) AddAccountIdentification() *SecuritiesAccount19 {
	a.AccountIdentification = new(SecuritiesAccount19)
	return a.AccountIdentification
}

func (a *AccountSubLevel15) AddAccountOwner() *PartyIdentification100 {
	a.AccountOwner = new(PartyIdentification100)
	return a.AccountOwner
}

func (a *AccountSubLevel15) AddAccountServicer() *PartyIdentification100 {
	a.AccountServicer = new(PartyIdentification100)
	return a.AccountServicer
}

func (a *AccountSubLevel15) AddBeneficialOwner() *BeneficialOwner2 {
	newValue := new(BeneficialOwner2)
	a.BeneficialOwner = append(a.BeneficialOwner, newValue)
	return newValue
}

func (a *AccountSubLevel15) AddBalanceDetails() *AggregateHoldingBalance3 {
	newValue := new(AggregateHoldingBalance3)
	a.BalanceDetails = append(a.BalanceDetails, newValue)
	return newValue
}

func (a *AccountSubLevel15) AddAccountSubLevel6() *AccountSubLevel16 {
	newValue := new(AccountSubLevel16)
	a.AccountSubLevel6 = append(a.AccountSubLevel6, newValue)
	return newValue
}

func (a *AccountSubLevel15) AddAccountSubLevel6Difference() *AggregateHoldingBalance2 {
	newValue := new(AggregateHoldingBalance2)
	a.AccountSubLevel6Difference = append(a.AccountSubLevel6Difference, newValue)
	return newValue
}

func (a *AccountSubLevel15) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}

// Account and holding of the next sub-level (Level 6).
type AccountSubLevel16 struct {

	// Unique and unambiguous identification for the sub-account between the account owner and the account servicer.
	AccountIdentification *SecuritiesAccount19 `xml:"AcctId"`

	// Party that legally owns the sub-account.
	AccountOwner *PartyIdentification100 `xml:"AcctOwnr"`

	// Party that manages the sub-account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification100 `xml:"AcctSvcr"`

	// Individual or entity that is ultimately entitled to the benefit of income and rights in a financial instrument, as opposed to a nominal or legal owner.
	BeneficialOwner []*BeneficialOwner2 `xml:"BnfclOwnr,omitempty"`

	// Report on the net position of a financial instrument on the sub-account (sub-account level 6),  for a certain date. The agent, for example, a trade intermediary, may also be specified.
	BalanceDetails []*AggregateHoldingBalance3 `xml:"BalDtls,omitempty"`

	// Holdings of level 7.
	AccountSubLevel7 []*AccountSubLevel17 `xml:"AcctSubLvl7,omitempty"`

	// Difference in holdings between the sub-account at level 6 and the sub-accounts of level 7.
	AccountSubLevel7Difference []*AggregateHoldingBalance2 `xml:"AcctSubLvl7Diff,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AccountSubLevel16) AddAccountIdentification() *SecuritiesAccount19 {
	a.AccountIdentification = new(SecuritiesAccount19)
	return a.AccountIdentification
}

func (a *AccountSubLevel16) AddAccountOwner() *PartyIdentification100 {
	a.AccountOwner = new(PartyIdentification100)
	return a.AccountOwner
}

func (a *AccountSubLevel16) AddAccountServicer() *PartyIdentification100 {
	a.AccountServicer = new(PartyIdentification100)
	return a.AccountServicer
}

func (a *AccountSubLevel16) AddBeneficialOwner() *BeneficialOwner2 {
	newValue := new(BeneficialOwner2)
	a.BeneficialOwner = append(a.BeneficialOwner, newValue)
	return newValue
}

func (a *AccountSubLevel16) AddBalanceDetails() *AggregateHoldingBalance3 {
	newValue := new(AggregateHoldingBalance3)
	a.BalanceDetails = append(a.BalanceDetails, newValue)
	return newValue
}

func (a *AccountSubLevel16) AddAccountSubLevel7() *AccountSubLevel17 {
	newValue := new(AccountSubLevel17)
	a.AccountSubLevel7 = append(a.AccountSubLevel7, newValue)
	return newValue
}

func (a *AccountSubLevel16) AddAccountSubLevel7Difference() *AggregateHoldingBalance2 {
	newValue := new(AggregateHoldingBalance2)
	a.AccountSubLevel7Difference = append(a.AccountSubLevel7Difference, newValue)
	return newValue
}

func (a *AccountSubLevel16) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}

// Account and holding of the next sub-level (Level 7).
type AccountSubLevel17 struct {

	// Unique and unambiguous identification for the sub-account between the account owner and the account servicer.
	AccountIdentification *SecuritiesAccount19 `xml:"AcctId"`

	// Party that legally owns the sub-account.
	AccountOwner *PartyIdentification100 `xml:"AcctOwnr"`

	// Party that manages the sub-account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification100 `xml:"AcctSvcr"`

	// Individual or entity that is ultimately entitled to the benefit of income and rights in a financial instrument, as opposed to a nominal or legal owner.
	BeneficialOwner []*BeneficialOwner2 `xml:"BnfclOwnr,omitempty"`

	// Report on the net position of a financial instrument on the sub-account (sub-account level 7),  for a certain date. The agent, for example, a trade intermediary, may also be specified.
	BalanceDetails []*AggregateHoldingBalance3 `xml:"BalDtls,omitempty"`

	// Holdings of level 8.
	AccountSubLevel8 []*AccountSubLevel18 `xml:"AcctSubLvl8,omitempty"`

	// Difference in holdings between the sub-account at level 7 and the sub-accounts of level 8.
	AccountSubLevel8Difference []*AggregateHoldingBalance2 `xml:"AcctSubLvl8Diff,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AccountSubLevel17) AddAccountIdentification() *SecuritiesAccount19 {
	a.AccountIdentification = new(SecuritiesAccount19)
	return a.AccountIdentification
}

func (a *AccountSubLevel17) AddAccountOwner() *PartyIdentification100 {
	a.AccountOwner = new(PartyIdentification100)
	return a.AccountOwner
}

func (a *AccountSubLevel17) AddAccountServicer() *PartyIdentification100 {
	a.AccountServicer = new(PartyIdentification100)
	return a.AccountServicer
}

func (a *AccountSubLevel17) AddBeneficialOwner() *BeneficialOwner2 {
	newValue := new(BeneficialOwner2)
	a.BeneficialOwner = append(a.BeneficialOwner, newValue)
	return newValue
}

func (a *AccountSubLevel17) AddBalanceDetails() *AggregateHoldingBalance3 {
	newValue := new(AggregateHoldingBalance3)
	a.BalanceDetails = append(a.BalanceDetails, newValue)
	return newValue
}

func (a *AccountSubLevel17) AddAccountSubLevel8() *AccountSubLevel18 {
	newValue := new(AccountSubLevel18)
	a.AccountSubLevel8 = append(a.AccountSubLevel8, newValue)
	return newValue
}

func (a *AccountSubLevel17) AddAccountSubLevel8Difference() *AggregateHoldingBalance2 {
	newValue := new(AggregateHoldingBalance2)
	a.AccountSubLevel8Difference = append(a.AccountSubLevel8Difference, newValue)
	return newValue
}

func (a *AccountSubLevel17) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}

// Account and holding of the next sub-level (Level 8).
type AccountSubLevel18 struct {

	// Unique and unambiguous identification for the sub-account between the account owner and the account servicer.
	AccountIdentification *SecuritiesAccount19 `xml:"AcctId"`

	// Party that legally owns the sub-account.
	AccountOwner *PartyIdentification100 `xml:"AcctOwnr"`

	// Party that manages the sub-account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification100 `xml:"AcctSvcr"`

	// Individual or entity that is ultimately entitled to the benefit of income and rights in a financial instrument, as opposed to a nominal or legal owner.
	BeneficialOwner []*BeneficialOwner2 `xml:"BnfclOwnr,omitempty"`

	// Report on the net position of a financial instrument on the sub-account (sub-account level 8),  for a certain date. The agent, for example, a trade intermediary, may also be specified.
	BalanceDetails []*AggregateHoldingBalance3 `xml:"BalDtls,omitempty"`

	// Holdings of level 9.
	AccountSubLevel9 []*AccountSubLevel19 `xml:"AcctSubLvl9,omitempty"`

	// Difference in holdings between the sub-account at level 8 and the sub-accounts of level 9.
	AccountSubLevel9Difference []*AggregateHoldingBalance2 `xml:"AcctSubLvl9Diff,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AccountSubLevel18) AddAccountIdentification() *SecuritiesAccount19 {
	a.AccountIdentification = new(SecuritiesAccount19)
	return a.AccountIdentification
}

func (a *AccountSubLevel18) AddAccountOwner() *PartyIdentification100 {
	a.AccountOwner = new(PartyIdentification100)
	return a.AccountOwner
}

func (a *AccountSubLevel18) AddAccountServicer() *PartyIdentification100 {
	a.AccountServicer = new(PartyIdentification100)
	return a.AccountServicer
}

func (a *AccountSubLevel18) AddBeneficialOwner() *BeneficialOwner2 {
	newValue := new(BeneficialOwner2)
	a.BeneficialOwner = append(a.BeneficialOwner, newValue)
	return newValue
}

func (a *AccountSubLevel18) AddBalanceDetails() *AggregateHoldingBalance3 {
	newValue := new(AggregateHoldingBalance3)
	a.BalanceDetails = append(a.BalanceDetails, newValue)
	return newValue
}

func (a *AccountSubLevel18) AddAccountSubLevel9() *AccountSubLevel19 {
	newValue := new(AccountSubLevel19)
	a.AccountSubLevel9 = append(a.AccountSubLevel9, newValue)
	return newValue
}

func (a *AccountSubLevel18) AddAccountSubLevel9Difference() *AggregateHoldingBalance2 {
	newValue := new(AggregateHoldingBalance2)
	a.AccountSubLevel9Difference = append(a.AccountSubLevel9Difference, newValue)
	return newValue
}

func (a *AccountSubLevel18) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}

// Account and holding of the next sub-level (Level 9).
type AccountSubLevel19 struct {

	// Unique and unambiguous identification for the sub-account between the account owner and the account servicer.
	AccountIdentification *SecuritiesAccount19 `xml:"AcctId"`

	// Party that legally owns the sub-account.
	AccountOwner *PartyIdentification100 `xml:"AcctOwnr"`

	// Party that manages the sub-account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification100 `xml:"AcctSvcr"`

	// Individual or entity that is ultimately entitled to the benefit of income and rights in a financial instrument, as opposed to a nominal or legal owner.
	BeneficialOwner []*BeneficialOwner2 `xml:"BnfclOwnr,omitempty"`

	// Report on the net position of a financial instrument on the sub-account (sub-account level 9),  for a certain date. The agent, for example, a trade intermediary, may also be specified.
	BalanceDetails []*AggregateHoldingBalance3 `xml:"BalDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AccountSubLevel19) AddAccountIdentification() *SecuritiesAccount19 {
	a.AccountIdentification = new(SecuritiesAccount19)
	return a.AccountIdentification
}

func (a *AccountSubLevel19) AddAccountOwner() *PartyIdentification100 {
	a.AccountOwner = new(PartyIdentification100)
	return a.AccountOwner
}

func (a *AccountSubLevel19) AddAccountServicer() *PartyIdentification100 {
	a.AccountServicer = new(PartyIdentification100)
	return a.AccountServicer
}

func (a *AccountSubLevel19) AddBeneficialOwner() *BeneficialOwner2 {
	newValue := new(BeneficialOwner2)
	a.BeneficialOwner = append(a.BeneficialOwner, newValue)
	return newValue
}

func (a *AccountSubLevel19) AddBalanceDetails() *AggregateHoldingBalance3 {
	newValue := new(AggregateHoldingBalance3)
	a.BalanceDetails = append(a.BalanceDetails, newValue)
	return newValue
}

func (a *AccountSubLevel19) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}
