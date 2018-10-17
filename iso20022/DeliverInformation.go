package iso20022

// Parameters applied to the settlement of a security transfer.
type DeliverInformation1 struct {

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *DeliveringPartiesAndAccount1 `xml:"SttlmPtiesDtls"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalTransferIndicator *YesNoIndicator `xml:"PhysTrfInd"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters2 `xml:"PhysTrfDtls,omitempty"`
}

func (d *DeliverInformation1) AddSettlementPartiesDetails() *DeliveringPartiesAndAccount1 {
	d.SettlementPartiesDetails = new(DeliveringPartiesAndAccount1)
	return d.SettlementPartiesDetails
}

func (d *DeliverInformation1) SetPhysicalTransferIndicator(value string) {
	d.PhysicalTransferIndicator = (*YesNoIndicator)(&value)
}

func (d *DeliverInformation1) AddPhysicalTransferDetails() *DeliveryParameters2 {
	d.PhysicalTransferDetails = new(DeliveryParameters2)
	return d.PhysicalTransferDetails
}

// Completion of a securities settlement instruction, wherein securities are delivered/debited from a securities account and received/credited to the designated securities account.
type DeliverInformation11 struct {

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Date and time at which the securities were exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	EffectiveSettlementDate *DateAndDateTimeChoice `xml:"FctvSttlmDt,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Indicates whether the settlement amount includes the stamp duty amount.
	StampDuty *StampDutyType2Code `xml:"StmpDty,omitempty"`

	// Deal amount.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Charge related to the transfer of a financial instrument.
	ChargeDetails []*Charge20 `xml:"ChrgDtls,omitempty"`

	// Commission related to the transfer of a financial instrument.
	CommissionDetails []*Commission12 `xml:"ComssnDtls,omitempty"`

	// Tax related to the transfer of a financial instrument.
	TaxDetails []*Tax15 `xml:"TaxDtls,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *DeliveringPartiesAndAccount8 `xml:"SttlmPtiesDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically  delivered.
	PhysicalTransfer *PhysicalTransferType1Code `xml:"PhysTrf,omitempty"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters4 `xml:"PhysTrfDtls,omitempty"`
}

func (d *DeliverInformation11) SetRequestedSettlementDate(value string) {
	d.RequestedSettlementDate = (*ISODate)(&value)
}

func (d *DeliverInformation11) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	d.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return d.EffectiveSettlementDate
}

func (d *DeliverInformation11) SetSettlementAmount(value, currency string) {
	d.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DeliverInformation11) SetStampDuty(value string) {
	d.StampDuty = (*StampDutyType2Code)(&value)
}

func (d *DeliverInformation11) SetNetAmount(value, currency string) {
	d.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DeliverInformation11) AddChargeDetails() *Charge20 {
	newValue := new(Charge20)
	d.ChargeDetails = append(d.ChargeDetails, newValue)
	return newValue
}

func (d *DeliverInformation11) AddCommissionDetails() *Commission12 {
	newValue := new(Commission12)
	d.CommissionDetails = append(d.CommissionDetails, newValue)
	return newValue
}

func (d *DeliverInformation11) AddTaxDetails() *Tax15 {
	newValue := new(Tax15)
	d.TaxDetails = append(d.TaxDetails, newValue)
	return newValue
}

func (d *DeliverInformation11) AddSettlementPartiesDetails() *DeliveringPartiesAndAccount8 {
	d.SettlementPartiesDetails = new(DeliveringPartiesAndAccount8)
	return d.SettlementPartiesDetails
}

func (d *DeliverInformation11) SetPhysicalTransfer(value string) {
	d.PhysicalTransfer = (*PhysicalTransferType1Code)(&value)
}

func (d *DeliverInformation11) AddPhysicalTransferDetails() *DeliveryParameters4 {
	d.PhysicalTransferDetails = new(DeliveryParameters4)
	return d.PhysicalTransferDetails
}

// Parameters applied to the settlement of a security transfer.
type DeliverInformation12 struct {

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Date and time at which the securities were exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	EffectiveSettlementDate *DateAndDateTimeChoice `xml:"FctvSttlmDt,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Indicates whether the settlement amount includes the stamp duty amount.
	StampDuty *StampDutyType2Code `xml:"StmpDty,omitempty"`

	// Deal amount.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Charge related to the transfer of a financial instrument.
	ChargeDetails []*Charge20 `xml:"ChrgDtls,omitempty"`

	// Commission related to the transfer of a financial instrument.
	CommissionDetails []*Commission17 `xml:"ComssnDtls,omitempty"`

	// Tax related to the transfer of a financial instrument.
	TaxDetails []*Tax21 `xml:"TaxDtls,omitempty"`

	// Specifies foreign exchange details applied to the payment of charges, taxes and commissions as a result of the transfer.
	ForeignExchangeDetails []*ForeignExchangeTerms7 `xml:"FXDtls,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *DeliveringPartiesAndAccount9 `xml:"SttlmPtiesDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically  delivered.
	PhysicalTransfer *PhysicalTransferType1Code `xml:"PhysTrf,omitempty"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters4 `xml:"PhysTrfDtls,omitempty"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`
}

func (d *DeliverInformation12) SetRequestedSettlementDate(value string) {
	d.RequestedSettlementDate = (*ISODate)(&value)
}

func (d *DeliverInformation12) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	d.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return d.EffectiveSettlementDate
}

func (d *DeliverInformation12) SetSettlementAmount(value, currency string) {
	d.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DeliverInformation12) SetStampDuty(value string) {
	d.StampDuty = (*StampDutyType2Code)(&value)
}

func (d *DeliverInformation12) SetNetAmount(value, currency string) {
	d.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DeliverInformation12) AddChargeDetails() *Charge20 {
	newValue := new(Charge20)
	d.ChargeDetails = append(d.ChargeDetails, newValue)
	return newValue
}

func (d *DeliverInformation12) AddCommissionDetails() *Commission17 {
	newValue := new(Commission17)
	d.CommissionDetails = append(d.CommissionDetails, newValue)
	return newValue
}

func (d *DeliverInformation12) AddTaxDetails() *Tax21 {
	newValue := new(Tax21)
	d.TaxDetails = append(d.TaxDetails, newValue)
	return newValue
}

func (d *DeliverInformation12) AddForeignExchangeDetails() *ForeignExchangeTerms7 {
	newValue := new(ForeignExchangeTerms7)
	d.ForeignExchangeDetails = append(d.ForeignExchangeDetails, newValue)
	return newValue
}

func (d *DeliverInformation12) AddSettlementPartiesDetails() *DeliveringPartiesAndAccount9 {
	d.SettlementPartiesDetails = new(DeliveringPartiesAndAccount9)
	return d.SettlementPartiesDetails
}

func (d *DeliverInformation12) SetPhysicalTransfer(value string) {
	d.PhysicalTransfer = (*PhysicalTransferType1Code)(&value)
}

func (d *DeliverInformation12) AddPhysicalTransferDetails() *DeliveryParameters4 {
	d.PhysicalTransferDetails = new(DeliveryParameters4)
	return d.PhysicalTransferDetails
}

func (d *DeliverInformation12) SetClientReference(value string) {
	d.ClientReference = (*Max35Text)(&value)
}

// Parameters applied to the settlement of a security transfer.
type DeliverInformation13 struct {

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Indicates whether the settlement amount includes the stamp duty amount.
	StampDuty *StampDutyType2Code `xml:"StmpDty,omitempty"`

	// Deal amount.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *DeliveringPartiesAndAccount9 `xml:"SttlmPtiesDtls,omitempty"`

	// Charge related to the transfer of a financial instrument.
	ChargeDetails []*Charge20 `xml:"ChrgDtls,omitempty"`

	// Commission related to the transfer of a financial instrument.
	CommissionDetails []*Commission17 `xml:"ComssnDtls,omitempty"`

	// Tax related to the transfer of a financial instrument.
	TaxDetails []*Tax21 `xml:"TaxDtls,omitempty"`

	// Specifies foreign exchange details applied to the payment of charges, taxes and commissions as a result of the transfer.
	ForeignExchangeDetails []*ForeignExchangeTerms7 `xml:"FXDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalTransfer *PhysicalTransferType1Code `xml:"PhysTrf,omitempty"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters4 `xml:"PhysTrfDtls,omitempty"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`
}

func (d *DeliverInformation13) SetRequestedSettlementDate(value string) {
	d.RequestedSettlementDate = (*ISODate)(&value)
}

func (d *DeliverInformation13) SetSettlementAmount(value, currency string) {
	d.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DeliverInformation13) SetStampDuty(value string) {
	d.StampDuty = (*StampDutyType2Code)(&value)
}

func (d *DeliverInformation13) SetNetAmount(value, currency string) {
	d.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DeliverInformation13) AddSettlementPartiesDetails() *DeliveringPartiesAndAccount9 {
	d.SettlementPartiesDetails = new(DeliveringPartiesAndAccount9)
	return d.SettlementPartiesDetails
}

func (d *DeliverInformation13) AddChargeDetails() *Charge20 {
	newValue := new(Charge20)
	d.ChargeDetails = append(d.ChargeDetails, newValue)
	return newValue
}

func (d *DeliverInformation13) AddCommissionDetails() *Commission17 {
	newValue := new(Commission17)
	d.CommissionDetails = append(d.CommissionDetails, newValue)
	return newValue
}

func (d *DeliverInformation13) AddTaxDetails() *Tax21 {
	newValue := new(Tax21)
	d.TaxDetails = append(d.TaxDetails, newValue)
	return newValue
}

func (d *DeliverInformation13) AddForeignExchangeDetails() *ForeignExchangeTerms7 {
	newValue := new(ForeignExchangeTerms7)
	d.ForeignExchangeDetails = append(d.ForeignExchangeDetails, newValue)
	return newValue
}

func (d *DeliverInformation13) SetPhysicalTransfer(value string) {
	d.PhysicalTransfer = (*PhysicalTransferType1Code)(&value)
}

func (d *DeliverInformation13) AddPhysicalTransferDetails() *DeliveryParameters4 {
	d.PhysicalTransferDetails = new(DeliveryParameters4)
	return d.PhysicalTransferDetails
}

func (d *DeliverInformation13) SetClientReference(value string) {
	d.ClientReference = (*Max35Text)(&value)
}

// Parameters applied to the settlement of a security transfer.
type DeliverInformation14 struct {

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Date and time at which the securities were exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	EffectiveSettlementDate *DateAndDateTimeChoice `xml:"FctvSttlmDt,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Indicates whether the settlement amount includes the stamp duty amount.
	StampDuty *StampDutyType2Code `xml:"StmpDty,omitempty"`

	// Deal amount.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Charge related to the transfer of a financial instrument.
	ChargeDetails []*Charge27 `xml:"ChrgDtls,omitempty"`

	// Commission related to the transfer of a financial instrument.
	CommissionDetails []*Commission22 `xml:"ComssnDtls,omitempty"`

	// Tax related to the transfer of a financial instrument.
	TaxDetails []*Tax25 `xml:"TaxDtls,omitempty"`

	// Specifies foreign exchange details applied to the payment of charges, taxes and commissions as a result of the transfer.
	ForeignExchangeDetails []*ForeignExchangeTerms7 `xml:"FXDtls,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *DeliveringPartiesAndAccount9 `xml:"SttlmPtiesDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically  delivered.
	PhysicalTransfer *PhysicalTransferType1Code `xml:"PhysTrf,omitempty"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters4 `xml:"PhysTrfDtls,omitempty"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`
}

func (d *DeliverInformation14) SetRequestedSettlementDate(value string) {
	d.RequestedSettlementDate = (*ISODate)(&value)
}

func (d *DeliverInformation14) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	d.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return d.EffectiveSettlementDate
}

func (d *DeliverInformation14) SetSettlementAmount(value, currency string) {
	d.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DeliverInformation14) SetStampDuty(value string) {
	d.StampDuty = (*StampDutyType2Code)(&value)
}

func (d *DeliverInformation14) SetNetAmount(value, currency string) {
	d.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DeliverInformation14) AddChargeDetails() *Charge27 {
	newValue := new(Charge27)
	d.ChargeDetails = append(d.ChargeDetails, newValue)
	return newValue
}

func (d *DeliverInformation14) AddCommissionDetails() *Commission22 {
	newValue := new(Commission22)
	d.CommissionDetails = append(d.CommissionDetails, newValue)
	return newValue
}

func (d *DeliverInformation14) AddTaxDetails() *Tax25 {
	newValue := new(Tax25)
	d.TaxDetails = append(d.TaxDetails, newValue)
	return newValue
}

func (d *DeliverInformation14) AddForeignExchangeDetails() *ForeignExchangeTerms7 {
	newValue := new(ForeignExchangeTerms7)
	d.ForeignExchangeDetails = append(d.ForeignExchangeDetails, newValue)
	return newValue
}

func (d *DeliverInformation14) AddSettlementPartiesDetails() *DeliveringPartiesAndAccount9 {
	d.SettlementPartiesDetails = new(DeliveringPartiesAndAccount9)
	return d.SettlementPartiesDetails
}

func (d *DeliverInformation14) SetPhysicalTransfer(value string) {
	d.PhysicalTransfer = (*PhysicalTransferType1Code)(&value)
}

func (d *DeliverInformation14) AddPhysicalTransferDetails() *DeliveryParameters4 {
	d.PhysicalTransferDetails = new(DeliveryParameters4)
	return d.PhysicalTransferDetails
}

func (d *DeliverInformation14) SetClientReference(value string) {
	d.ClientReference = (*Max35Text)(&value)
}

// Parameters applied to the settlement of a security transfer.
type DeliverInformation15 struct {

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Indicates whether the settlement amount includes the stamp duty amount.
	StampDuty *StampDutyType2Code `xml:"StmpDty,omitempty"`

	// Deal amount.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *DeliveringPartiesAndAccount9 `xml:"SttlmPtiesDtls,omitempty"`

	// Charge related to the transfer of a financial instrument.
	ChargeDetails []*Charge27 `xml:"ChrgDtls,omitempty"`

	// Commission related to the transfer of a financial instrument.
	CommissionDetails []*Commission22 `xml:"ComssnDtls,omitempty"`

	// Tax related to the transfer of a financial instrument.
	TaxDetails []*Tax25 `xml:"TaxDtls,omitempty"`

	// Specifies foreign exchange details applied to the payment of charges, taxes and commissions as a result of the transfer.
	ForeignExchangeDetails []*ForeignExchangeTerms7 `xml:"FXDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalTransfer *PhysicalTransferType1Code `xml:"PhysTrf,omitempty"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters4 `xml:"PhysTrfDtls,omitempty"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`
}

func (d *DeliverInformation15) SetRequestedSettlementDate(value string) {
	d.RequestedSettlementDate = (*ISODate)(&value)
}

func (d *DeliverInformation15) SetSettlementAmount(value, currency string) {
	d.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DeliverInformation15) SetStampDuty(value string) {
	d.StampDuty = (*StampDutyType2Code)(&value)
}

func (d *DeliverInformation15) SetNetAmount(value, currency string) {
	d.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DeliverInformation15) AddSettlementPartiesDetails() *DeliveringPartiesAndAccount9 {
	d.SettlementPartiesDetails = new(DeliveringPartiesAndAccount9)
	return d.SettlementPartiesDetails
}

func (d *DeliverInformation15) AddChargeDetails() *Charge27 {
	newValue := new(Charge27)
	d.ChargeDetails = append(d.ChargeDetails, newValue)
	return newValue
}

func (d *DeliverInformation15) AddCommissionDetails() *Commission22 {
	newValue := new(Commission22)
	d.CommissionDetails = append(d.CommissionDetails, newValue)
	return newValue
}

func (d *DeliverInformation15) AddTaxDetails() *Tax25 {
	newValue := new(Tax25)
	d.TaxDetails = append(d.TaxDetails, newValue)
	return newValue
}

func (d *DeliverInformation15) AddForeignExchangeDetails() *ForeignExchangeTerms7 {
	newValue := new(ForeignExchangeTerms7)
	d.ForeignExchangeDetails = append(d.ForeignExchangeDetails, newValue)
	return newValue
}

func (d *DeliverInformation15) SetPhysicalTransfer(value string) {
	d.PhysicalTransfer = (*PhysicalTransferType1Code)(&value)
}

func (d *DeliverInformation15) AddPhysicalTransferDetails() *DeliveryParameters4 {
	d.PhysicalTransferDetails = new(DeliveryParameters4)
	return d.PhysicalTransferDetails
}

func (d *DeliverInformation15) SetClientReference(value string) {
	d.ClientReference = (*Max35Text)(&value)
}

// Parameters applied to the settlement of a security transfer.
type DeliverInformation16 struct {

	// Party that delivers (transferor) securities to the receiving agent (transferee).
	Transferor *PartyIdentification70Choice `xml:"Trfr,omitempty"`

	// Account from which the securities are to be delivered.
	TransferorRegisteredAccount *Account19 `xml:"TrfrRegdAcct,omitempty"`

	// Identification of a related party or intermediary.
	IntermediaryInformation []*Intermediary34 `xml:"IntrmyInf,omitempty"`

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Indicates whether the settlement amount includes the stamp duty amount.
	StampDuty *StampDutyType2Code `xml:"StmpDty,omitempty"`

	// Deal amount.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *DeliveringPartiesAndAccount13 `xml:"SttlmPtiesDtls,omitempty"`

	// Charge related to the transfer of a financial instrument.
	ChargeDetails []*Charge29 `xml:"ChrgDtls,omitempty"`

	// Commission related to the transfer of a financial instrument.
	CommissionDetails []*Commission23 `xml:"ComssnDtls,omitempty"`

	// Tax related to the transfer of a financial instrument.
	TaxDetails []*Tax28 `xml:"TaxDtls,omitempty"`

	// Specifies foreign exchange details applied to the payment of charges, taxes and commissions as a result of the transfer.
	ForeignExchangeDetails []*ForeignExchangeTerms26 `xml:"FXDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalTransfer *PhysicalTransferType1Code `xml:"PhysTrf,omitempty"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters4 `xml:"PhysTrfDtls,omitempty"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *AdditionalReference7 `xml:"ClntRef,omitempty"`
}

func (d *DeliverInformation16) AddTransferor() *PartyIdentification70Choice {
	d.Transferor = new(PartyIdentification70Choice)
	return d.Transferor
}

func (d *DeliverInformation16) AddTransferorRegisteredAccount() *Account19 {
	d.TransferorRegisteredAccount = new(Account19)
	return d.TransferorRegisteredAccount
}

func (d *DeliverInformation16) AddIntermediaryInformation() *Intermediary34 {
	newValue := new(Intermediary34)
	d.IntermediaryInformation = append(d.IntermediaryInformation, newValue)
	return newValue
}

func (d *DeliverInformation16) SetRequestedSettlementDate(value string) {
	d.RequestedSettlementDate = (*ISODate)(&value)
}

func (d *DeliverInformation16) SetSettlementAmount(value, currency string) {
	d.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DeliverInformation16) SetStampDuty(value string) {
	d.StampDuty = (*StampDutyType2Code)(&value)
}

func (d *DeliverInformation16) SetNetAmount(value, currency string) {
	d.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DeliverInformation16) AddSettlementPartiesDetails() *DeliveringPartiesAndAccount13 {
	d.SettlementPartiesDetails = new(DeliveringPartiesAndAccount13)
	return d.SettlementPartiesDetails
}

func (d *DeliverInformation16) AddChargeDetails() *Charge29 {
	newValue := new(Charge29)
	d.ChargeDetails = append(d.ChargeDetails, newValue)
	return newValue
}

func (d *DeliverInformation16) AddCommissionDetails() *Commission23 {
	newValue := new(Commission23)
	d.CommissionDetails = append(d.CommissionDetails, newValue)
	return newValue
}

func (d *DeliverInformation16) AddTaxDetails() *Tax28 {
	newValue := new(Tax28)
	d.TaxDetails = append(d.TaxDetails, newValue)
	return newValue
}

func (d *DeliverInformation16) AddForeignExchangeDetails() *ForeignExchangeTerms26 {
	newValue := new(ForeignExchangeTerms26)
	d.ForeignExchangeDetails = append(d.ForeignExchangeDetails, newValue)
	return newValue
}

func (d *DeliverInformation16) SetPhysicalTransfer(value string) {
	d.PhysicalTransfer = (*PhysicalTransferType1Code)(&value)
}

func (d *DeliverInformation16) AddPhysicalTransferDetails() *DeliveryParameters4 {
	d.PhysicalTransferDetails = new(DeliveryParameters4)
	return d.PhysicalTransferDetails
}

func (d *DeliverInformation16) AddClientReference() *AdditionalReference7 {
	d.ClientReference = new(AdditionalReference7)
	return d.ClientReference
}

// Parameters applied to the settlement of a security transfer.
type DeliverInformation17 struct {

	// Party that delivers (transferor) securities to the receiving agent (transferee).
	Transferor *PartyIdentification70Choice `xml:"Trfr,omitempty"`

	// Account from which the securities are to be delivered.
	TransferorRegisteredAccount *Account19 `xml:"TrfrRegdAcct,omitempty"`

	// Identification of a related party or intermediary.
	IntermediaryInformation []*Intermediary34 `xml:"IntrmyInf,omitempty"`

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Date and time at which the securities were exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	EffectiveSettlementDate *DateAndDateTimeChoice `xml:"FctvSttlmDt,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Indicates whether the settlement amount includes the stamp duty amount.
	StampDuty *StampDutyType2Code `xml:"StmpDty,omitempty"`

	// Deal amount.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Charge related to the transfer of a financial instrument.
	ChargeDetails []*Charge29 `xml:"ChrgDtls,omitempty"`

	// Commission related to the transfer of a financial instrument.
	CommissionDetails []*Commission23 `xml:"ComssnDtls,omitempty"`

	// Tax related to the transfer of a financial instrument.
	TaxDetails []*Tax28 `xml:"TaxDtls,omitempty"`

	// Specifies foreign exchange details applied to the payment of charges, taxes and commissions as a result of the transfer.
	ForeignExchangeDetails []*ForeignExchangeTerms26 `xml:"FXDtls,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *DeliveringPartiesAndAccount13 `xml:"SttlmPtiesDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically  delivered.
	PhysicalTransfer *PhysicalTransferType1Code `xml:"PhysTrf,omitempty"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters4 `xml:"PhysTrfDtls,omitempty"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *AdditionalReference7 `xml:"ClntRef,omitempty"`
}

func (d *DeliverInformation17) AddTransferor() *PartyIdentification70Choice {
	d.Transferor = new(PartyIdentification70Choice)
	return d.Transferor
}

func (d *DeliverInformation17) AddTransferorRegisteredAccount() *Account19 {
	d.TransferorRegisteredAccount = new(Account19)
	return d.TransferorRegisteredAccount
}

func (d *DeliverInformation17) AddIntermediaryInformation() *Intermediary34 {
	newValue := new(Intermediary34)
	d.IntermediaryInformation = append(d.IntermediaryInformation, newValue)
	return newValue
}

func (d *DeliverInformation17) SetRequestedSettlementDate(value string) {
	d.RequestedSettlementDate = (*ISODate)(&value)
}

func (d *DeliverInformation17) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	d.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return d.EffectiveSettlementDate
}

func (d *DeliverInformation17) SetSettlementAmount(value, currency string) {
	d.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DeliverInformation17) SetStampDuty(value string) {
	d.StampDuty = (*StampDutyType2Code)(&value)
}

func (d *DeliverInformation17) SetNetAmount(value, currency string) {
	d.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DeliverInformation17) AddChargeDetails() *Charge29 {
	newValue := new(Charge29)
	d.ChargeDetails = append(d.ChargeDetails, newValue)
	return newValue
}

func (d *DeliverInformation17) AddCommissionDetails() *Commission23 {
	newValue := new(Commission23)
	d.CommissionDetails = append(d.CommissionDetails, newValue)
	return newValue
}

func (d *DeliverInformation17) AddTaxDetails() *Tax28 {
	newValue := new(Tax28)
	d.TaxDetails = append(d.TaxDetails, newValue)
	return newValue
}

func (d *DeliverInformation17) AddForeignExchangeDetails() *ForeignExchangeTerms26 {
	newValue := new(ForeignExchangeTerms26)
	d.ForeignExchangeDetails = append(d.ForeignExchangeDetails, newValue)
	return newValue
}

func (d *DeliverInformation17) AddSettlementPartiesDetails() *DeliveringPartiesAndAccount13 {
	d.SettlementPartiesDetails = new(DeliveringPartiesAndAccount13)
	return d.SettlementPartiesDetails
}

func (d *DeliverInformation17) SetPhysicalTransfer(value string) {
	d.PhysicalTransfer = (*PhysicalTransferType1Code)(&value)
}

func (d *DeliverInformation17) AddPhysicalTransferDetails() *DeliveryParameters4 {
	d.PhysicalTransferDetails = new(DeliveryParameters4)
	return d.PhysicalTransferDetails
}

func (d *DeliverInformation17) AddClientReference() *AdditionalReference7 {
	d.ClientReference = new(AdditionalReference7)
	return d.ClientReference
}

// Completion of a securities settlement instruction, wherein securities are delivered/debited from a securities account and received/credited to the designated securities account.
type DeliverInformation2 struct {

	// Charge related to the transfer of a financial instrument.
	ChargeDetails []*Charge4 `xml:"ChrgDtls,omitempty"`

	// Tax related to the transfer of a financial instrument.
	TaxDetails []*Tax3 `xml:"TaxDtls,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *DeliveringPartiesAndAccount1 `xml:"SttlmPtiesDtls"`

	// Indicates whether the financial instrument is to be physically  delivered.
	PhysicalTransferIndicator *YesNoIndicator `xml:"PhysTrfInd"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters2 `xml:"PhysTrfDtls,omitempty"`
}

func (d *DeliverInformation2) AddChargeDetails() *Charge4 {
	newValue := new(Charge4)
	d.ChargeDetails = append(d.ChargeDetails, newValue)
	return newValue
}

func (d *DeliverInformation2) AddTaxDetails() *Tax3 {
	newValue := new(Tax3)
	d.TaxDetails = append(d.TaxDetails, newValue)
	return newValue
}

func (d *DeliverInformation2) AddSettlementPartiesDetails() *DeliveringPartiesAndAccount1 {
	d.SettlementPartiesDetails = new(DeliveringPartiesAndAccount1)
	return d.SettlementPartiesDetails
}

func (d *DeliverInformation2) SetPhysicalTransferIndicator(value string) {
	d.PhysicalTransferIndicator = (*YesNoIndicator)(&value)
}

func (d *DeliverInformation2) AddPhysicalTransferDetails() *DeliveryParameters2 {
	d.PhysicalTransferDetails = new(DeliveryParameters2)
	return d.PhysicalTransferDetails
}

// Parameters applied to the settlement of a security transfer.
type DeliverInformation3 struct {

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Indicates whether the settlement amount includes the stamp duty amount.
	StampDutyIndicator *YesNoIndicator `xml:"StmpDtyInd"`

	// Deal amount.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *DeliveringPartiesAndAccount4 `xml:"SttlmPtiesDtls"`

	// Charge related to the transfer of a financial instrument.
	ChargeDetails []*Charge20 `xml:"ChrgDtls,omitempty"`

	// Commission related to the transfer of a financial instrument.
	CommissionDetails []*Commission12 `xml:"ComssnDtls,omitempty"`

	// Tax related to the transfer of a financial instrument.
	TaxDetails []*Tax15 `xml:"TaxDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalTransferIndicator *YesNoIndicator `xml:"PhysTrfInd"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters4 `xml:"PhysTrfDtls,omitempty"`
}

func (d *DeliverInformation3) SetSettlementAmount(value, currency string) {
	d.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DeliverInformation3) SetStampDutyIndicator(value string) {
	d.StampDutyIndicator = (*YesNoIndicator)(&value)
}

func (d *DeliverInformation3) SetNetAmount(value, currency string) {
	d.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DeliverInformation3) AddSettlementPartiesDetails() *DeliveringPartiesAndAccount4 {
	d.SettlementPartiesDetails = new(DeliveringPartiesAndAccount4)
	return d.SettlementPartiesDetails
}

func (d *DeliverInformation3) AddChargeDetails() *Charge20 {
	newValue := new(Charge20)
	d.ChargeDetails = append(d.ChargeDetails, newValue)
	return newValue
}

func (d *DeliverInformation3) AddCommissionDetails() *Commission12 {
	newValue := new(Commission12)
	d.CommissionDetails = append(d.CommissionDetails, newValue)
	return newValue
}

func (d *DeliverInformation3) AddTaxDetails() *Tax15 {
	newValue := new(Tax15)
	d.TaxDetails = append(d.TaxDetails, newValue)
	return newValue
}

func (d *DeliverInformation3) SetPhysicalTransferIndicator(value string) {
	d.PhysicalTransferIndicator = (*YesNoIndicator)(&value)
}

func (d *DeliverInformation3) AddPhysicalTransferDetails() *DeliveryParameters4 {
	d.PhysicalTransferDetails = new(DeliveryParameters4)
	return d.PhysicalTransferDetails
}

// Completion of a securities settlement instruction, wherein securities are delivered/debited from a securities account and received/credited to the designated securities account.
type DeliverInformation4 struct {

	// Date and time at which the securities were exchange at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	EffectiveSettlementDate *DateAndDateTimeChoice `xml:"FctvSttlmDt,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Indicates whether the settlement amount includes the stamp duty amount.
	StampDutyIndicator *YesNoIndicator `xml:"StmpDtyInd"`

	// Deal amount.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Charge related to the transfer of a financial instrument.
	ChargeDetails []*Charge20 `xml:"ChrgDtls,omitempty"`

	// Commission related to the transfer of a financial instrument.
	CommissionDetails []*Commission12 `xml:"ComssnDtls,omitempty"`

	// Tax related to the transfer of a financial instrument.
	TaxDetails []*Tax15 `xml:"TaxDtls,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *DeliveringPartiesAndAccount4 `xml:"SttlmPtiesDtls"`

	// Indicates whether the financial instrument is to be physically  delivered.
	PhysicalTransferIndicator *YesNoIndicator `xml:"PhysTrfInd"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters4 `xml:"PhysTrfDtls,omitempty"`
}

func (d *DeliverInformation4) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	d.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return d.EffectiveSettlementDate
}

func (d *DeliverInformation4) SetSettlementAmount(value, currency string) {
	d.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DeliverInformation4) SetStampDutyIndicator(value string) {
	d.StampDutyIndicator = (*YesNoIndicator)(&value)
}

func (d *DeliverInformation4) SetNetAmount(value, currency string) {
	d.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DeliverInformation4) AddChargeDetails() *Charge20 {
	newValue := new(Charge20)
	d.ChargeDetails = append(d.ChargeDetails, newValue)
	return newValue
}

func (d *DeliverInformation4) AddCommissionDetails() *Commission12 {
	newValue := new(Commission12)
	d.CommissionDetails = append(d.CommissionDetails, newValue)
	return newValue
}

func (d *DeliverInformation4) AddTaxDetails() *Tax15 {
	newValue := new(Tax15)
	d.TaxDetails = append(d.TaxDetails, newValue)
	return newValue
}

func (d *DeliverInformation4) AddSettlementPartiesDetails() *DeliveringPartiesAndAccount4 {
	d.SettlementPartiesDetails = new(DeliveringPartiesAndAccount4)
	return d.SettlementPartiesDetails
}

func (d *DeliverInformation4) SetPhysicalTransferIndicator(value string) {
	d.PhysicalTransferIndicator = (*YesNoIndicator)(&value)
}

func (d *DeliverInformation4) AddPhysicalTransferDetails() *DeliveryParameters4 {
	d.PhysicalTransferDetails = new(DeliveryParameters4)
	return d.PhysicalTransferDetails
}

// Parameters applied to the settlement of a security transfer.
type DeliverInformation5 struct {

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Indicates whether the settlement amount includes the stamp duty amount.
	StampDuty *StampDutyType2Code `xml:"StmpDty,omitempty"`

	// Deal amount.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *DeliveringPartiesAndAccount8 `xml:"SttlmPtiesDtls"`

	// Charge related to the transfer of a financial instrument.
	ChargeDetails []*Charge20 `xml:"ChrgDtls,omitempty"`

	// Commission related to the transfer of a financial instrument.
	CommissionDetails []*Commission12 `xml:"ComssnDtls,omitempty"`

	// Tax related to the transfer of a financial instrument.
	TaxDetails []*Tax15 `xml:"TaxDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalTransfer *PhysicalTransferType1Code `xml:"PhysTrf,omitempty"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters4 `xml:"PhysTrfDtls,omitempty"`
}

func (d *DeliverInformation5) SetSettlementAmount(value, currency string) {
	d.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DeliverInformation5) SetStampDuty(value string) {
	d.StampDuty = (*StampDutyType2Code)(&value)
}

func (d *DeliverInformation5) SetNetAmount(value, currency string) {
	d.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DeliverInformation5) AddSettlementPartiesDetails() *DeliveringPartiesAndAccount8 {
	d.SettlementPartiesDetails = new(DeliveringPartiesAndAccount8)
	return d.SettlementPartiesDetails
}

func (d *DeliverInformation5) AddChargeDetails() *Charge20 {
	newValue := new(Charge20)
	d.ChargeDetails = append(d.ChargeDetails, newValue)
	return newValue
}

func (d *DeliverInformation5) AddCommissionDetails() *Commission12 {
	newValue := new(Commission12)
	d.CommissionDetails = append(d.CommissionDetails, newValue)
	return newValue
}

func (d *DeliverInformation5) AddTaxDetails() *Tax15 {
	newValue := new(Tax15)
	d.TaxDetails = append(d.TaxDetails, newValue)
	return newValue
}

func (d *DeliverInformation5) SetPhysicalTransfer(value string) {
	d.PhysicalTransfer = (*PhysicalTransferType1Code)(&value)
}

func (d *DeliverInformation5) AddPhysicalTransferDetails() *DeliveryParameters4 {
	d.PhysicalTransferDetails = new(DeliveryParameters4)
	return d.PhysicalTransferDetails
}

// Completion of a securities settlement instruction, wherein securities are delivered/debited from a securities account and received/credited to the designated securities account.
type DeliverInformation6 struct {

	// Date and time at which the securities were exchange at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	EffectiveSettlementDate *DateAndDateTimeChoice `xml:"FctvSttlmDt,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Indicates whether the settlement amount includes the stamp duty amount.
	StampDuty *StampDutyType2Code `xml:"StmpDty,omitempty"`

	// Deal amount.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Charge related to the transfer of a financial instrument.
	ChargeDetails []*Charge20 `xml:"ChrgDtls,omitempty"`

	// Commission related to the transfer of a financial instrument.
	CommissionDetails []*Commission12 `xml:"ComssnDtls,omitempty"`

	// Tax related to the transfer of a financial instrument.
	TaxDetails []*Tax15 `xml:"TaxDtls,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *DeliveringPartiesAndAccount8 `xml:"SttlmPtiesDtls"`

	// Indicates whether the financial instrument is to be physically  delivered.
	PhysicalTransfer *PhysicalTransferType1Code `xml:"PhysTrf,omitempty"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters4 `xml:"PhysTrfDtls,omitempty"`
}

func (d *DeliverInformation6) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	d.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return d.EffectiveSettlementDate
}

func (d *DeliverInformation6) SetSettlementAmount(value, currency string) {
	d.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DeliverInformation6) SetStampDuty(value string) {
	d.StampDuty = (*StampDutyType2Code)(&value)
}

func (d *DeliverInformation6) SetNetAmount(value, currency string) {
	d.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DeliverInformation6) AddChargeDetails() *Charge20 {
	newValue := new(Charge20)
	d.ChargeDetails = append(d.ChargeDetails, newValue)
	return newValue
}

func (d *DeliverInformation6) AddCommissionDetails() *Commission12 {
	newValue := new(Commission12)
	d.CommissionDetails = append(d.CommissionDetails, newValue)
	return newValue
}

func (d *DeliverInformation6) AddTaxDetails() *Tax15 {
	newValue := new(Tax15)
	d.TaxDetails = append(d.TaxDetails, newValue)
	return newValue
}

func (d *DeliverInformation6) AddSettlementPartiesDetails() *DeliveringPartiesAndAccount8 {
	d.SettlementPartiesDetails = new(DeliveringPartiesAndAccount8)
	return d.SettlementPartiesDetails
}

func (d *DeliverInformation6) SetPhysicalTransfer(value string) {
	d.PhysicalTransfer = (*PhysicalTransferType1Code)(&value)
}

func (d *DeliverInformation6) AddPhysicalTransferDetails() *DeliveryParameters4 {
	d.PhysicalTransferDetails = new(DeliveryParameters4)
	return d.PhysicalTransferDetails
}

// Completion of a securities settlement instruction, wherein securities are delivered/debited from a securities account and received/credited to the designated securities account.
type DeliverInformation7 struct {

	// Date and time at which the securities were exchange at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	EffectiveSettlementDate *DateAndDateTimeChoice `xml:"FctvSttlmDt,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Indicates whether the settlement amount includes the stamp duty amount.
	StampDuty *StampDutyType2Code `xml:"StmpDty,omitempty"`

	// Deal amount.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Charge related to the transfer of a financial instrument.
	ChargeDetails []*Charge20 `xml:"ChrgDtls,omitempty"`

	// Commission related to the transfer of a financial instrument.
	CommissionDetails []*Commission12 `xml:"ComssnDtls,omitempty"`

	// Tax related to the transfer of a financial instrument.
	TaxDetails []*Tax15 `xml:"TaxDtls,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *DeliveringPartiesAndAccount8 `xml:"SttlmPtiesDtls"`

	// Indicates whether the financial instrument is to be physically  delivered.
	PhysicalTransfer *PhysicalTransferType1Code `xml:"PhysTrf,omitempty"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters4 `xml:"PhysTrfDtls,omitempty"`
}

func (d *DeliverInformation7) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	d.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return d.EffectiveSettlementDate
}

func (d *DeliverInformation7) SetSettlementAmount(value, currency string) {
	d.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DeliverInformation7) SetStampDuty(value string) {
	d.StampDuty = (*StampDutyType2Code)(&value)
}

func (d *DeliverInformation7) SetNetAmount(value, currency string) {
	d.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DeliverInformation7) AddChargeDetails() *Charge20 {
	newValue := new(Charge20)
	d.ChargeDetails = append(d.ChargeDetails, newValue)
	return newValue
}

func (d *DeliverInformation7) AddCommissionDetails() *Commission12 {
	newValue := new(Commission12)
	d.CommissionDetails = append(d.CommissionDetails, newValue)
	return newValue
}

func (d *DeliverInformation7) AddTaxDetails() *Tax15 {
	newValue := new(Tax15)
	d.TaxDetails = append(d.TaxDetails, newValue)
	return newValue
}

func (d *DeliverInformation7) AddSettlementPartiesDetails() *DeliveringPartiesAndAccount8 {
	d.SettlementPartiesDetails = new(DeliveringPartiesAndAccount8)
	return d.SettlementPartiesDetails
}

func (d *DeliverInformation7) SetPhysicalTransfer(value string) {
	d.PhysicalTransfer = (*PhysicalTransferType1Code)(&value)
}

func (d *DeliverInformation7) AddPhysicalTransferDetails() *DeliveryParameters4 {
	d.PhysicalTransferDetails = new(DeliveryParameters4)
	return d.PhysicalTransferDetails
}

// Parameters applied to the settlement of a security transfer.
type DeliverInformation8 struct {

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Indicates whether the settlement amount includes the stamp duty amount.
	StampDuty *StampDutyType2Code `xml:"StmpDty,omitempty"`

	// Deal amount.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *DeliveringPartiesAndAccount8 `xml:"SttlmPtiesDtls"`

	// Charge related to the transfer of a financial instrument.
	ChargeDetails []*Charge20 `xml:"ChrgDtls,omitempty"`

	// Commission related to the transfer of a financial instrument.
	CommissionDetails []*Commission12 `xml:"ComssnDtls,omitempty"`

	// Tax related to the transfer of a financial instrument.
	TaxDetails []*Tax15 `xml:"TaxDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalTransfer *PhysicalTransferType1Code `xml:"PhysTrf,omitempty"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters4 `xml:"PhysTrfDtls,omitempty"`
}

func (d *DeliverInformation8) SetSettlementAmount(value, currency string) {
	d.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DeliverInformation8) SetStampDuty(value string) {
	d.StampDuty = (*StampDutyType2Code)(&value)
}

func (d *DeliverInformation8) SetNetAmount(value, currency string) {
	d.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DeliverInformation8) AddSettlementPartiesDetails() *DeliveringPartiesAndAccount8 {
	d.SettlementPartiesDetails = new(DeliveringPartiesAndAccount8)
	return d.SettlementPartiesDetails
}

func (d *DeliverInformation8) AddChargeDetails() *Charge20 {
	newValue := new(Charge20)
	d.ChargeDetails = append(d.ChargeDetails, newValue)
	return newValue
}

func (d *DeliverInformation8) AddCommissionDetails() *Commission12 {
	newValue := new(Commission12)
	d.CommissionDetails = append(d.CommissionDetails, newValue)
	return newValue
}

func (d *DeliverInformation8) AddTaxDetails() *Tax15 {
	newValue := new(Tax15)
	d.TaxDetails = append(d.TaxDetails, newValue)
	return newValue
}

func (d *DeliverInformation8) SetPhysicalTransfer(value string) {
	d.PhysicalTransfer = (*PhysicalTransferType1Code)(&value)
}

func (d *DeliverInformation8) AddPhysicalTransferDetails() *DeliveryParameters4 {
	d.PhysicalTransferDetails = new(DeliveryParameters4)
	return d.PhysicalTransferDetails
}

// Parameters applied to the settlement of a security transfer.
type DeliverInformation9 struct {

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Indicates whether the settlement amount includes the stamp duty amount.
	StampDuty *StampDutyType2Code `xml:"StmpDty,omitempty"`

	// Deal amount.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *DeliveringPartiesAndAccount8 `xml:"SttlmPtiesDtls,omitempty"`

	// Charge related to the transfer of a financial instrument.
	ChargeDetails []*Charge20 `xml:"ChrgDtls,omitempty"`

	// Commission related to the transfer of a financial instrument.
	CommissionDetails []*Commission12 `xml:"ComssnDtls,omitempty"`

	// Tax related to the transfer of a financial instrument.
	TaxDetails []*Tax15 `xml:"TaxDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalTransfer *PhysicalTransferType1Code `xml:"PhysTrf,omitempty"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters4 `xml:"PhysTrfDtls,omitempty"`
}

func (d *DeliverInformation9) SetRequestedSettlementDate(value string) {
	d.RequestedSettlementDate = (*ISODate)(&value)
}

func (d *DeliverInformation9) SetSettlementAmount(value, currency string) {
	d.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DeliverInformation9) SetStampDuty(value string) {
	d.StampDuty = (*StampDutyType2Code)(&value)
}

func (d *DeliverInformation9) SetNetAmount(value, currency string) {
	d.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DeliverInformation9) AddSettlementPartiesDetails() *DeliveringPartiesAndAccount8 {
	d.SettlementPartiesDetails = new(DeliveringPartiesAndAccount8)
	return d.SettlementPartiesDetails
}

func (d *DeliverInformation9) AddChargeDetails() *Charge20 {
	newValue := new(Charge20)
	d.ChargeDetails = append(d.ChargeDetails, newValue)
	return newValue
}

func (d *DeliverInformation9) AddCommissionDetails() *Commission12 {
	newValue := new(Commission12)
	d.CommissionDetails = append(d.CommissionDetails, newValue)
	return newValue
}

func (d *DeliverInformation9) AddTaxDetails() *Tax15 {
	newValue := new(Tax15)
	d.TaxDetails = append(d.TaxDetails, newValue)
	return newValue
}

func (d *DeliverInformation9) SetPhysicalTransfer(value string) {
	d.PhysicalTransfer = (*PhysicalTransferType1Code)(&value)
}

func (d *DeliverInformation9) AddPhysicalTransferDetails() *DeliveryParameters4 {
	d.PhysicalTransferDetails = new(DeliveryParameters4)
	return d.PhysicalTransferDetails
}
