package iso20022

// Parameters applied to the settlement of a security transfer.
type ReceiveInformation1 struct {

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *ReceivingPartiesAndAccount1 `xml:"SttlmPtiesDtls"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalTransferIndicator *YesNoIndicator `xml:"PhysTrfInd"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters2 `xml:"PhysTrfDtls,omitempty"`
}

func (r *ReceiveInformation1) AddSettlementPartiesDetails() *ReceivingPartiesAndAccount1 {
	r.SettlementPartiesDetails = new(ReceivingPartiesAndAccount1)
	return r.SettlementPartiesDetails
}

func (r *ReceiveInformation1) SetPhysicalTransferIndicator(value string) {
	r.PhysicalTransferIndicator = (*YesNoIndicator)(&value)
}

func (r *ReceiveInformation1) AddPhysicalTransferDetails() *DeliveryParameters2 {
	r.PhysicalTransferDetails = new(DeliveryParameters2)
	return r.PhysicalTransferDetails
}

// Completion of a securities settlement instruction, wherein securities are delivered/debited from a securities account and received/credited to the designated securities account.
type ReceiveInformation11 struct {

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
	SettlementPartiesDetails *ReceivingPartiesAndAccount8 `xml:"SttlmPtiesDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically  delivered.
	PhysicalTransfer *PhysicalTransferType1Code `xml:"PhysTrf,omitempty"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters4 `xml:"PhysTrfDtls,omitempty"`
}

func (r *ReceiveInformation11) SetRequestedSettlementDate(value string) {
	r.RequestedSettlementDate = (*ISODate)(&value)
}

func (r *ReceiveInformation11) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	r.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return r.EffectiveSettlementDate
}

func (r *ReceiveInformation11) SetSettlementAmount(value, currency string) {
	r.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *ReceiveInformation11) SetStampDuty(value string) {
	r.StampDuty = (*StampDutyType2Code)(&value)
}

func (r *ReceiveInformation11) SetNetAmount(value, currency string) {
	r.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *ReceiveInformation11) AddChargeDetails() *Charge20 {
	newValue := new(Charge20)
	r.ChargeDetails = append(r.ChargeDetails, newValue)
	return newValue
}

func (r *ReceiveInformation11) AddCommissionDetails() *Commission12 {
	newValue := new(Commission12)
	r.CommissionDetails = append(r.CommissionDetails, newValue)
	return newValue
}

func (r *ReceiveInformation11) AddTaxDetails() *Tax15 {
	newValue := new(Tax15)
	r.TaxDetails = append(r.TaxDetails, newValue)
	return newValue
}

func (r *ReceiveInformation11) AddSettlementPartiesDetails() *ReceivingPartiesAndAccount8 {
	r.SettlementPartiesDetails = new(ReceivingPartiesAndAccount8)
	return r.SettlementPartiesDetails
}

func (r *ReceiveInformation11) SetPhysicalTransfer(value string) {
	r.PhysicalTransfer = (*PhysicalTransferType1Code)(&value)
}

func (r *ReceiveInformation11) AddPhysicalTransferDetails() *DeliveryParameters4 {
	r.PhysicalTransferDetails = new(DeliveryParameters4)
	return r.PhysicalTransferDetails
}

// Parameters applied to the settlement of a security transfer.
type ReceiveInformation12 struct {

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
	SettlementPartiesDetails *ReceivingPartiesAndAccount9 `xml:"SttlmPtiesDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically  delivered.
	PhysicalTransfer *PhysicalTransferType1Code `xml:"PhysTrf,omitempty"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters4 `xml:"PhysTrfDtls,omitempty"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`
}

func (r *ReceiveInformation12) SetRequestedSettlementDate(value string) {
	r.RequestedSettlementDate = (*ISODate)(&value)
}

func (r *ReceiveInformation12) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	r.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return r.EffectiveSettlementDate
}

func (r *ReceiveInformation12) SetSettlementAmount(value, currency string) {
	r.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *ReceiveInformation12) SetStampDuty(value string) {
	r.StampDuty = (*StampDutyType2Code)(&value)
}

func (r *ReceiveInformation12) SetNetAmount(value, currency string) {
	r.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *ReceiveInformation12) AddChargeDetails() *Charge20 {
	newValue := new(Charge20)
	r.ChargeDetails = append(r.ChargeDetails, newValue)
	return newValue
}

func (r *ReceiveInformation12) AddCommissionDetails() *Commission17 {
	newValue := new(Commission17)
	r.CommissionDetails = append(r.CommissionDetails, newValue)
	return newValue
}

func (r *ReceiveInformation12) AddTaxDetails() *Tax21 {
	newValue := new(Tax21)
	r.TaxDetails = append(r.TaxDetails, newValue)
	return newValue
}

func (r *ReceiveInformation12) AddForeignExchangeDetails() *ForeignExchangeTerms7 {
	newValue := new(ForeignExchangeTerms7)
	r.ForeignExchangeDetails = append(r.ForeignExchangeDetails, newValue)
	return newValue
}

func (r *ReceiveInformation12) AddSettlementPartiesDetails() *ReceivingPartiesAndAccount9 {
	r.SettlementPartiesDetails = new(ReceivingPartiesAndAccount9)
	return r.SettlementPartiesDetails
}

func (r *ReceiveInformation12) SetPhysicalTransfer(value string) {
	r.PhysicalTransfer = (*PhysicalTransferType1Code)(&value)
}

func (r *ReceiveInformation12) AddPhysicalTransferDetails() *DeliveryParameters4 {
	r.PhysicalTransferDetails = new(DeliveryParameters4)
	return r.PhysicalTransferDetails
}

func (r *ReceiveInformation12) SetClientReference(value string) {
	r.ClientReference = (*Max35Text)(&value)
}

// Parameters applied to the settlement of a security transfer.
type ReceiveInformation13 struct {

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Indicates whether the settlement amount includes the stamp duty amount.
	StampDuty *StampDutyType2Code `xml:"StmpDty,omitempty"`

	// Deal amount.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *ReceivingPartiesAndAccount9 `xml:"SttlmPtiesDtls,omitempty"`

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

func (r *ReceiveInformation13) SetRequestedSettlementDate(value string) {
	r.RequestedSettlementDate = (*ISODate)(&value)
}

func (r *ReceiveInformation13) SetSettlementAmount(value, currency string) {
	r.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *ReceiveInformation13) SetStampDuty(value string) {
	r.StampDuty = (*StampDutyType2Code)(&value)
}

func (r *ReceiveInformation13) SetNetAmount(value, currency string) {
	r.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *ReceiveInformation13) AddSettlementPartiesDetails() *ReceivingPartiesAndAccount9 {
	r.SettlementPartiesDetails = new(ReceivingPartiesAndAccount9)
	return r.SettlementPartiesDetails
}

func (r *ReceiveInformation13) AddChargeDetails() *Charge20 {
	newValue := new(Charge20)
	r.ChargeDetails = append(r.ChargeDetails, newValue)
	return newValue
}

func (r *ReceiveInformation13) AddCommissionDetails() *Commission17 {
	newValue := new(Commission17)
	r.CommissionDetails = append(r.CommissionDetails, newValue)
	return newValue
}

func (r *ReceiveInformation13) AddTaxDetails() *Tax21 {
	newValue := new(Tax21)
	r.TaxDetails = append(r.TaxDetails, newValue)
	return newValue
}

func (r *ReceiveInformation13) AddForeignExchangeDetails() *ForeignExchangeTerms7 {
	newValue := new(ForeignExchangeTerms7)
	r.ForeignExchangeDetails = append(r.ForeignExchangeDetails, newValue)
	return newValue
}

func (r *ReceiveInformation13) SetPhysicalTransfer(value string) {
	r.PhysicalTransfer = (*PhysicalTransferType1Code)(&value)
}

func (r *ReceiveInformation13) AddPhysicalTransferDetails() *DeliveryParameters4 {
	r.PhysicalTransferDetails = new(DeliveryParameters4)
	return r.PhysicalTransferDetails
}

func (r *ReceiveInformation13) SetClientReference(value string) {
	r.ClientReference = (*Max35Text)(&value)
}

// Parameters applied to the settlement of a security transfer.
type ReceiveInformation14 struct {

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
	SettlementPartiesDetails *ReceivingPartiesAndAccount9 `xml:"SttlmPtiesDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically  delivered.
	PhysicalTransfer *PhysicalTransferType1Code `xml:"PhysTrf,omitempty"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters4 `xml:"PhysTrfDtls,omitempty"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`
}

func (r *ReceiveInformation14) SetRequestedSettlementDate(value string) {
	r.RequestedSettlementDate = (*ISODate)(&value)
}

func (r *ReceiveInformation14) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	r.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return r.EffectiveSettlementDate
}

func (r *ReceiveInformation14) SetSettlementAmount(value, currency string) {
	r.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *ReceiveInformation14) SetStampDuty(value string) {
	r.StampDuty = (*StampDutyType2Code)(&value)
}

func (r *ReceiveInformation14) SetNetAmount(value, currency string) {
	r.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *ReceiveInformation14) AddChargeDetails() *Charge27 {
	newValue := new(Charge27)
	r.ChargeDetails = append(r.ChargeDetails, newValue)
	return newValue
}

func (r *ReceiveInformation14) AddCommissionDetails() *Commission22 {
	newValue := new(Commission22)
	r.CommissionDetails = append(r.CommissionDetails, newValue)
	return newValue
}

func (r *ReceiveInformation14) AddTaxDetails() *Tax25 {
	newValue := new(Tax25)
	r.TaxDetails = append(r.TaxDetails, newValue)
	return newValue
}

func (r *ReceiveInformation14) AddForeignExchangeDetails() *ForeignExchangeTerms7 {
	newValue := new(ForeignExchangeTerms7)
	r.ForeignExchangeDetails = append(r.ForeignExchangeDetails, newValue)
	return newValue
}

func (r *ReceiveInformation14) AddSettlementPartiesDetails() *ReceivingPartiesAndAccount9 {
	r.SettlementPartiesDetails = new(ReceivingPartiesAndAccount9)
	return r.SettlementPartiesDetails
}

func (r *ReceiveInformation14) SetPhysicalTransfer(value string) {
	r.PhysicalTransfer = (*PhysicalTransferType1Code)(&value)
}

func (r *ReceiveInformation14) AddPhysicalTransferDetails() *DeliveryParameters4 {
	r.PhysicalTransferDetails = new(DeliveryParameters4)
	return r.PhysicalTransferDetails
}

func (r *ReceiveInformation14) SetClientReference(value string) {
	r.ClientReference = (*Max35Text)(&value)
}

// Parameters applied to the settlement of a security transfer.
type ReceiveInformation15 struct {

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Indicates whether the settlement amount includes the stamp duty amount.
	StampDuty *StampDutyType2Code `xml:"StmpDty,omitempty"`

	// Deal amount.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *ReceivingPartiesAndAccount9 `xml:"SttlmPtiesDtls,omitempty"`

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

func (r *ReceiveInformation15) SetRequestedSettlementDate(value string) {
	r.RequestedSettlementDate = (*ISODate)(&value)
}

func (r *ReceiveInformation15) SetSettlementAmount(value, currency string) {
	r.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *ReceiveInformation15) SetStampDuty(value string) {
	r.StampDuty = (*StampDutyType2Code)(&value)
}

func (r *ReceiveInformation15) SetNetAmount(value, currency string) {
	r.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *ReceiveInformation15) AddSettlementPartiesDetails() *ReceivingPartiesAndAccount9 {
	r.SettlementPartiesDetails = new(ReceivingPartiesAndAccount9)
	return r.SettlementPartiesDetails
}

func (r *ReceiveInformation15) AddChargeDetails() *Charge27 {
	newValue := new(Charge27)
	r.ChargeDetails = append(r.ChargeDetails, newValue)
	return newValue
}

func (r *ReceiveInformation15) AddCommissionDetails() *Commission22 {
	newValue := new(Commission22)
	r.CommissionDetails = append(r.CommissionDetails, newValue)
	return newValue
}

func (r *ReceiveInformation15) AddTaxDetails() *Tax25 {
	newValue := new(Tax25)
	r.TaxDetails = append(r.TaxDetails, newValue)
	return newValue
}

func (r *ReceiveInformation15) AddForeignExchangeDetails() *ForeignExchangeTerms7 {
	newValue := new(ForeignExchangeTerms7)
	r.ForeignExchangeDetails = append(r.ForeignExchangeDetails, newValue)
	return newValue
}

func (r *ReceiveInformation15) SetPhysicalTransfer(value string) {
	r.PhysicalTransfer = (*PhysicalTransferType1Code)(&value)
}

func (r *ReceiveInformation15) AddPhysicalTransferDetails() *DeliveryParameters4 {
	r.PhysicalTransferDetails = new(DeliveryParameters4)
	return r.PhysicalTransferDetails
}

func (r *ReceiveInformation15) SetClientReference(value string) {
	r.ClientReference = (*Max35Text)(&value)
}

// Parameters applied to the settlement of a security transfer.
type ReceiveInformation16 struct {

	// Party that receives (transferee) securities from the delivering agent (transferor).
	Transferee *PartyIdentification70Choice `xml:"Trfee,omitempty"`

	// Account into which the securities are to be received.
	TransfereeRegisteredAccount *Account19 `xml:"TrfeeRegdAcct,omitempty"`

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
	SettlementPartiesDetails *ReceivingPartiesAndAccount13 `xml:"SttlmPtiesDtls,omitempty"`

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

func (r *ReceiveInformation16) AddTransferee() *PartyIdentification70Choice {
	r.Transferee = new(PartyIdentification70Choice)
	return r.Transferee
}

func (r *ReceiveInformation16) AddTransfereeRegisteredAccount() *Account19 {
	r.TransfereeRegisteredAccount = new(Account19)
	return r.TransfereeRegisteredAccount
}

func (r *ReceiveInformation16) AddIntermediaryInformation() *Intermediary34 {
	newValue := new(Intermediary34)
	r.IntermediaryInformation = append(r.IntermediaryInformation, newValue)
	return newValue
}

func (r *ReceiveInformation16) SetRequestedSettlementDate(value string) {
	r.RequestedSettlementDate = (*ISODate)(&value)
}

func (r *ReceiveInformation16) SetSettlementAmount(value, currency string) {
	r.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *ReceiveInformation16) SetStampDuty(value string) {
	r.StampDuty = (*StampDutyType2Code)(&value)
}

func (r *ReceiveInformation16) SetNetAmount(value, currency string) {
	r.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *ReceiveInformation16) AddSettlementPartiesDetails() *ReceivingPartiesAndAccount13 {
	r.SettlementPartiesDetails = new(ReceivingPartiesAndAccount13)
	return r.SettlementPartiesDetails
}

func (r *ReceiveInformation16) AddChargeDetails() *Charge29 {
	newValue := new(Charge29)
	r.ChargeDetails = append(r.ChargeDetails, newValue)
	return newValue
}

func (r *ReceiveInformation16) AddCommissionDetails() *Commission23 {
	newValue := new(Commission23)
	r.CommissionDetails = append(r.CommissionDetails, newValue)
	return newValue
}

func (r *ReceiveInformation16) AddTaxDetails() *Tax28 {
	newValue := new(Tax28)
	r.TaxDetails = append(r.TaxDetails, newValue)
	return newValue
}

func (r *ReceiveInformation16) AddForeignExchangeDetails() *ForeignExchangeTerms26 {
	newValue := new(ForeignExchangeTerms26)
	r.ForeignExchangeDetails = append(r.ForeignExchangeDetails, newValue)
	return newValue
}

func (r *ReceiveInformation16) SetPhysicalTransfer(value string) {
	r.PhysicalTransfer = (*PhysicalTransferType1Code)(&value)
}

func (r *ReceiveInformation16) AddPhysicalTransferDetails() *DeliveryParameters4 {
	r.PhysicalTransferDetails = new(DeliveryParameters4)
	return r.PhysicalTransferDetails
}

func (r *ReceiveInformation16) AddClientReference() *AdditionalReference7 {
	r.ClientReference = new(AdditionalReference7)
	return r.ClientReference
}

// Parameters applied to the settlement of a security transfer.
type ReceiveInformation17 struct {

	// Party that receives (transferee) securities from the delivering agent (transferor).
	Transferee *PartyIdentification70Choice `xml:"Trfee,omitempty"`

	// Account into which the securities are to be received.
	TransfereeRegisteredAccount *Account19 `xml:"TrfeeRegdAcct,omitempty"`

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
	SettlementPartiesDetails *ReceivingPartiesAndAccount13 `xml:"SttlmPtiesDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically  delivered.
	PhysicalTransfer *PhysicalTransferType1Code `xml:"PhysTrf,omitempty"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters4 `xml:"PhysTrfDtls,omitempty"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *AdditionalReference7 `xml:"ClntRef,omitempty"`
}

func (r *ReceiveInformation17) AddTransferee() *PartyIdentification70Choice {
	r.Transferee = new(PartyIdentification70Choice)
	return r.Transferee
}

func (r *ReceiveInformation17) AddTransfereeRegisteredAccount() *Account19 {
	r.TransfereeRegisteredAccount = new(Account19)
	return r.TransfereeRegisteredAccount
}

func (r *ReceiveInformation17) AddIntermediaryInformation() *Intermediary34 {
	newValue := new(Intermediary34)
	r.IntermediaryInformation = append(r.IntermediaryInformation, newValue)
	return newValue
}

func (r *ReceiveInformation17) SetRequestedSettlementDate(value string) {
	r.RequestedSettlementDate = (*ISODate)(&value)
}

func (r *ReceiveInformation17) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	r.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return r.EffectiveSettlementDate
}

func (r *ReceiveInformation17) SetSettlementAmount(value, currency string) {
	r.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *ReceiveInformation17) SetStampDuty(value string) {
	r.StampDuty = (*StampDutyType2Code)(&value)
}

func (r *ReceiveInformation17) SetNetAmount(value, currency string) {
	r.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *ReceiveInformation17) AddChargeDetails() *Charge29 {
	newValue := new(Charge29)
	r.ChargeDetails = append(r.ChargeDetails, newValue)
	return newValue
}

func (r *ReceiveInformation17) AddCommissionDetails() *Commission23 {
	newValue := new(Commission23)
	r.CommissionDetails = append(r.CommissionDetails, newValue)
	return newValue
}

func (r *ReceiveInformation17) AddTaxDetails() *Tax28 {
	newValue := new(Tax28)
	r.TaxDetails = append(r.TaxDetails, newValue)
	return newValue
}

func (r *ReceiveInformation17) AddForeignExchangeDetails() *ForeignExchangeTerms26 {
	newValue := new(ForeignExchangeTerms26)
	r.ForeignExchangeDetails = append(r.ForeignExchangeDetails, newValue)
	return newValue
}

func (r *ReceiveInformation17) AddSettlementPartiesDetails() *ReceivingPartiesAndAccount13 {
	r.SettlementPartiesDetails = new(ReceivingPartiesAndAccount13)
	return r.SettlementPartiesDetails
}

func (r *ReceiveInformation17) SetPhysicalTransfer(value string) {
	r.PhysicalTransfer = (*PhysicalTransferType1Code)(&value)
}

func (r *ReceiveInformation17) AddPhysicalTransferDetails() *DeliveryParameters4 {
	r.PhysicalTransferDetails = new(DeliveryParameters4)
	return r.PhysicalTransferDetails
}

func (r *ReceiveInformation17) AddClientReference() *AdditionalReference7 {
	r.ClientReference = new(AdditionalReference7)
	return r.ClientReference
}

// Completion of a securities settlement instruction, wherein securities are delivered/debited from a securities account and received/credited to the designated securities account.
type ReceiveInformation2 struct {

	// Charge related to the transfer of a financial instrument.
	ChargeDetails []*Charge4 `xml:"ChrgDtls,omitempty"`

	// Tax related to the transfer of a financial instrument.
	TaxDetails []*Tax3 `xml:"TaxDtls,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *ReceivingPartiesAndAccount1 `xml:"SttlmPtiesDtls"`

	// Indicates whether the financial instrument is to be physically  delivered.
	PhysicalTransferIndicator *YesNoIndicator `xml:"PhysTrfInd"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters2 `xml:"PhysTrfDtls,omitempty"`
}

func (r *ReceiveInformation2) AddChargeDetails() *Charge4 {
	newValue := new(Charge4)
	r.ChargeDetails = append(r.ChargeDetails, newValue)
	return newValue
}

func (r *ReceiveInformation2) AddTaxDetails() *Tax3 {
	newValue := new(Tax3)
	r.TaxDetails = append(r.TaxDetails, newValue)
	return newValue
}

func (r *ReceiveInformation2) AddSettlementPartiesDetails() *ReceivingPartiesAndAccount1 {
	r.SettlementPartiesDetails = new(ReceivingPartiesAndAccount1)
	return r.SettlementPartiesDetails
}

func (r *ReceiveInformation2) SetPhysicalTransferIndicator(value string) {
	r.PhysicalTransferIndicator = (*YesNoIndicator)(&value)
}

func (r *ReceiveInformation2) AddPhysicalTransferDetails() *DeliveryParameters2 {
	r.PhysicalTransferDetails = new(DeliveryParameters2)
	return r.PhysicalTransferDetails
}

// Parameters applied to the settlement of a security transfer.
type ReceiveInformation3 struct {

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Indicates whether the settlement amount includes the stamp duty amount.
	StampDutyIndicator *YesNoIndicator `xml:"StmpDtyInd"`

	// Deal amount.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *ReceivingPartiesAndAccount4 `xml:"SttlmPtiesDtls"`

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

func (r *ReceiveInformation3) SetSettlementAmount(value, currency string) {
	r.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *ReceiveInformation3) SetStampDutyIndicator(value string) {
	r.StampDutyIndicator = (*YesNoIndicator)(&value)
}

func (r *ReceiveInformation3) SetNetAmount(value, currency string) {
	r.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *ReceiveInformation3) AddSettlementPartiesDetails() *ReceivingPartiesAndAccount4 {
	r.SettlementPartiesDetails = new(ReceivingPartiesAndAccount4)
	return r.SettlementPartiesDetails
}

func (r *ReceiveInformation3) AddChargeDetails() *Charge20 {
	newValue := new(Charge20)
	r.ChargeDetails = append(r.ChargeDetails, newValue)
	return newValue
}

func (r *ReceiveInformation3) AddCommissionDetails() *Commission12 {
	newValue := new(Commission12)
	r.CommissionDetails = append(r.CommissionDetails, newValue)
	return newValue
}

func (r *ReceiveInformation3) AddTaxDetails() *Tax15 {
	newValue := new(Tax15)
	r.TaxDetails = append(r.TaxDetails, newValue)
	return newValue
}

func (r *ReceiveInformation3) SetPhysicalTransferIndicator(value string) {
	r.PhysicalTransferIndicator = (*YesNoIndicator)(&value)
}

func (r *ReceiveInformation3) AddPhysicalTransferDetails() *DeliveryParameters4 {
	r.PhysicalTransferDetails = new(DeliveryParameters4)
	return r.PhysicalTransferDetails
}

// Completion of a securities settlement instruction, wherein securities are delivered/debited from a securities account and received/credited to the designated securities account.
type ReceiveInformation4 struct {

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
	SettlementPartiesDetails *ReceivingPartiesAndAccount4 `xml:"SttlmPtiesDtls"`

	// Indicates whether the financial instrument is to be physically  delivered.
	PhysicalTransferIndicator *YesNoIndicator `xml:"PhysTrfInd"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters4 `xml:"PhysTrfDtls,omitempty"`
}

func (r *ReceiveInformation4) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	r.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return r.EffectiveSettlementDate
}

func (r *ReceiveInformation4) SetSettlementAmount(value, currency string) {
	r.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *ReceiveInformation4) SetStampDutyIndicator(value string) {
	r.StampDutyIndicator = (*YesNoIndicator)(&value)
}

func (r *ReceiveInformation4) SetNetAmount(value, currency string) {
	r.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *ReceiveInformation4) AddChargeDetails() *Charge20 {
	newValue := new(Charge20)
	r.ChargeDetails = append(r.ChargeDetails, newValue)
	return newValue
}

func (r *ReceiveInformation4) AddCommissionDetails() *Commission12 {
	newValue := new(Commission12)
	r.CommissionDetails = append(r.CommissionDetails, newValue)
	return newValue
}

func (r *ReceiveInformation4) AddTaxDetails() *Tax15 {
	newValue := new(Tax15)
	r.TaxDetails = append(r.TaxDetails, newValue)
	return newValue
}

func (r *ReceiveInformation4) AddSettlementPartiesDetails() *ReceivingPartiesAndAccount4 {
	r.SettlementPartiesDetails = new(ReceivingPartiesAndAccount4)
	return r.SettlementPartiesDetails
}

func (r *ReceiveInformation4) SetPhysicalTransferIndicator(value string) {
	r.PhysicalTransferIndicator = (*YesNoIndicator)(&value)
}

func (r *ReceiveInformation4) AddPhysicalTransferDetails() *DeliveryParameters4 {
	r.PhysicalTransferDetails = new(DeliveryParameters4)
	return r.PhysicalTransferDetails
}

// Parameters applied to the settlement of a security transfer.
type ReceiveInformation5 struct {

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Indicates whether the settlement amount includes the stamp duty amount.
	StampDuty *StampDutyType2Code `xml:"StmpDty,omitempty"`

	// Deal amount.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *ReceivingPartiesAndAccount8 `xml:"SttlmPtiesDtls"`

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

func (r *ReceiveInformation5) SetSettlementAmount(value, currency string) {
	r.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *ReceiveInformation5) SetStampDuty(value string) {
	r.StampDuty = (*StampDutyType2Code)(&value)
}

func (r *ReceiveInformation5) SetNetAmount(value, currency string) {
	r.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *ReceiveInformation5) AddSettlementPartiesDetails() *ReceivingPartiesAndAccount8 {
	r.SettlementPartiesDetails = new(ReceivingPartiesAndAccount8)
	return r.SettlementPartiesDetails
}

func (r *ReceiveInformation5) AddChargeDetails() *Charge20 {
	newValue := new(Charge20)
	r.ChargeDetails = append(r.ChargeDetails, newValue)
	return newValue
}

func (r *ReceiveInformation5) AddCommissionDetails() *Commission12 {
	newValue := new(Commission12)
	r.CommissionDetails = append(r.CommissionDetails, newValue)
	return newValue
}

func (r *ReceiveInformation5) AddTaxDetails() *Tax15 {
	newValue := new(Tax15)
	r.TaxDetails = append(r.TaxDetails, newValue)
	return newValue
}

func (r *ReceiveInformation5) SetPhysicalTransfer(value string) {
	r.PhysicalTransfer = (*PhysicalTransferType1Code)(&value)
}

func (r *ReceiveInformation5) AddPhysicalTransferDetails() *DeliveryParameters4 {
	r.PhysicalTransferDetails = new(DeliveryParameters4)
	return r.PhysicalTransferDetails
}

// Completion of a securities settlement instruction, wherein securities are delivered/debited from a securities account and received/credited to the designated securities account.
type ReceiveInformation6 struct {

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
	SettlementPartiesDetails *ReceivingPartiesAndAccount8 `xml:"SttlmPtiesDtls"`

	// Indicates whether the financial instrument is to be physically  delivered.
	PhysicalTransfer *PhysicalTransferType1Code `xml:"PhysTrf,omitempty"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters4 `xml:"PhysTrfDtls,omitempty"`
}

func (r *ReceiveInformation6) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	r.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return r.EffectiveSettlementDate
}

func (r *ReceiveInformation6) SetSettlementAmount(value, currency string) {
	r.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *ReceiveInformation6) SetStampDuty(value string) {
	r.StampDuty = (*StampDutyType2Code)(&value)
}

func (r *ReceiveInformation6) SetNetAmount(value, currency string) {
	r.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *ReceiveInformation6) AddChargeDetails() *Charge20 {
	newValue := new(Charge20)
	r.ChargeDetails = append(r.ChargeDetails, newValue)
	return newValue
}

func (r *ReceiveInformation6) AddCommissionDetails() *Commission12 {
	newValue := new(Commission12)
	r.CommissionDetails = append(r.CommissionDetails, newValue)
	return newValue
}

func (r *ReceiveInformation6) AddTaxDetails() *Tax15 {
	newValue := new(Tax15)
	r.TaxDetails = append(r.TaxDetails, newValue)
	return newValue
}

func (r *ReceiveInformation6) AddSettlementPartiesDetails() *ReceivingPartiesAndAccount8 {
	r.SettlementPartiesDetails = new(ReceivingPartiesAndAccount8)
	return r.SettlementPartiesDetails
}

func (r *ReceiveInformation6) SetPhysicalTransfer(value string) {
	r.PhysicalTransfer = (*PhysicalTransferType1Code)(&value)
}

func (r *ReceiveInformation6) AddPhysicalTransferDetails() *DeliveryParameters4 {
	r.PhysicalTransferDetails = new(DeliveryParameters4)
	return r.PhysicalTransferDetails
}

// Parameters applied to the settlement of a security transfer.
type ReceiveInformation7 struct {

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Indicates whether the settlement amount includes the stamp duty amount.
	StampDuty *StampDutyType2Code `xml:"StmpDty,omitempty"`

	// Deal amount.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *ReceivingPartiesAndAccount8 `xml:"SttlmPtiesDtls"`

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

func (r *ReceiveInformation7) SetSettlementAmount(value, currency string) {
	r.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *ReceiveInformation7) SetStampDuty(value string) {
	r.StampDuty = (*StampDutyType2Code)(&value)
}

func (r *ReceiveInformation7) SetNetAmount(value, currency string) {
	r.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *ReceiveInformation7) AddSettlementPartiesDetails() *ReceivingPartiesAndAccount8 {
	r.SettlementPartiesDetails = new(ReceivingPartiesAndAccount8)
	return r.SettlementPartiesDetails
}

func (r *ReceiveInformation7) AddChargeDetails() *Charge20 {
	newValue := new(Charge20)
	r.ChargeDetails = append(r.ChargeDetails, newValue)
	return newValue
}

func (r *ReceiveInformation7) AddCommissionDetails() *Commission12 {
	newValue := new(Commission12)
	r.CommissionDetails = append(r.CommissionDetails, newValue)
	return newValue
}

func (r *ReceiveInformation7) AddTaxDetails() *Tax15 {
	newValue := new(Tax15)
	r.TaxDetails = append(r.TaxDetails, newValue)
	return newValue
}

func (r *ReceiveInformation7) SetPhysicalTransfer(value string) {
	r.PhysicalTransfer = (*PhysicalTransferType1Code)(&value)
}

func (r *ReceiveInformation7) AddPhysicalTransferDetails() *DeliveryParameters4 {
	r.PhysicalTransferDetails = new(DeliveryParameters4)
	return r.PhysicalTransferDetails
}

// Completion of a securities settlement instruction, wherein securities are delivered/debited from a securities account and received/credited to the designated securities account.
type ReceiveInformation8 struct {

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
	SettlementPartiesDetails *ReceivingPartiesAndAccount8 `xml:"SttlmPtiesDtls"`

	// Indicates whether the financial instrument is to be physically  delivered.
	PhysicalTransfer *PhysicalTransferType1Code `xml:"PhysTrf,omitempty"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters4 `xml:"PhysTrfDtls,omitempty"`
}

func (r *ReceiveInformation8) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	r.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return r.EffectiveSettlementDate
}

func (r *ReceiveInformation8) SetSettlementAmount(value, currency string) {
	r.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *ReceiveInformation8) SetStampDuty(value string) {
	r.StampDuty = (*StampDutyType2Code)(&value)
}

func (r *ReceiveInformation8) SetNetAmount(value, currency string) {
	r.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *ReceiveInformation8) AddChargeDetails() *Charge20 {
	newValue := new(Charge20)
	r.ChargeDetails = append(r.ChargeDetails, newValue)
	return newValue
}

func (r *ReceiveInformation8) AddCommissionDetails() *Commission12 {
	newValue := new(Commission12)
	r.CommissionDetails = append(r.CommissionDetails, newValue)
	return newValue
}

func (r *ReceiveInformation8) AddTaxDetails() *Tax15 {
	newValue := new(Tax15)
	r.TaxDetails = append(r.TaxDetails, newValue)
	return newValue
}

func (r *ReceiveInformation8) AddSettlementPartiesDetails() *ReceivingPartiesAndAccount8 {
	r.SettlementPartiesDetails = new(ReceivingPartiesAndAccount8)
	return r.SettlementPartiesDetails
}

func (r *ReceiveInformation8) SetPhysicalTransfer(value string) {
	r.PhysicalTransfer = (*PhysicalTransferType1Code)(&value)
}

func (r *ReceiveInformation8) AddPhysicalTransferDetails() *DeliveryParameters4 {
	r.PhysicalTransferDetails = new(DeliveryParameters4)
	return r.PhysicalTransferDetails
}

// Parameters applied to the settlement of a security transfer.
type ReceiveInformation9 struct {

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Indicates whether the settlement amount includes the stamp duty amount.
	StampDuty *StampDutyType2Code `xml:"StmpDty,omitempty"`

	// Deal amount.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *ReceivingPartiesAndAccount8 `xml:"SttlmPtiesDtls,omitempty"`

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

func (r *ReceiveInformation9) SetRequestedSettlementDate(value string) {
	r.RequestedSettlementDate = (*ISODate)(&value)
}

func (r *ReceiveInformation9) SetSettlementAmount(value, currency string) {
	r.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *ReceiveInformation9) SetStampDuty(value string) {
	r.StampDuty = (*StampDutyType2Code)(&value)
}

func (r *ReceiveInformation9) SetNetAmount(value, currency string) {
	r.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *ReceiveInformation9) AddSettlementPartiesDetails() *ReceivingPartiesAndAccount8 {
	r.SettlementPartiesDetails = new(ReceivingPartiesAndAccount8)
	return r.SettlementPartiesDetails
}

func (r *ReceiveInformation9) AddChargeDetails() *Charge20 {
	newValue := new(Charge20)
	r.ChargeDetails = append(r.ChargeDetails, newValue)
	return newValue
}

func (r *ReceiveInformation9) AddCommissionDetails() *Commission12 {
	newValue := new(Commission12)
	r.CommissionDetails = append(r.CommissionDetails, newValue)
	return newValue
}

func (r *ReceiveInformation9) AddTaxDetails() *Tax15 {
	newValue := new(Tax15)
	r.TaxDetails = append(r.TaxDetails, newValue)
	return newValue
}

func (r *ReceiveInformation9) SetPhysicalTransfer(value string) {
	r.PhysicalTransfer = (*PhysicalTransferType1Code)(&value)
}

func (r *ReceiveInformation9) AddPhysicalTransferDetails() *DeliveryParameters4 {
	r.PhysicalTransferDetails = new(DeliveryParameters4)
	return r.PhysicalTransferDetails
}
