package iso20022

// Data related to the status report of a point of interaction (POI).
type TerminalManagementDataSet1 struct {

	// Identification of the data set containing the status report.
	Identification *DataSetIdentification2 `xml:"Id"`

	// Counter to identify a single data set within the whole transfer.
	SequenceCounter *Max9NumericText `xml:"SeqCntr,omitempty"`

	// Content of the status report.
	Content *StatusReportContent1 `xml:"Cntt"`
}

func (t *TerminalManagementDataSet1) AddIdentification() *DataSetIdentification2 {
	t.Identification = new(DataSetIdentification2)
	return t.Identification
}

func (t *TerminalManagementDataSet1) SetSequenceCounter(value string) {
	t.SequenceCounter = (*Max9NumericText)(&value)
}

func (t *TerminalManagementDataSet1) AddContent() *StatusReportContent1 {
	t.Content = new(StatusReportContent1)
	return t.Content
}

// Data related to the management plan of a point of interaction (POI).
type TerminalManagementDataSet10 struct {

	// Identification of the data set containing the management plan.
	Identification *DataSetIdentification3 `xml:"Id"`

	// Counter to identify a single data set within the whole transfer.
	SequenceCounter *Max9NumericText `xml:"SeqCntr,omitempty"`

	// Content of the management plan.
	Content *ManagementPlanContent3 `xml:"Cntt,omitempty"`
}

func (t *TerminalManagementDataSet10) AddIdentification() *DataSetIdentification3 {
	t.Identification = new(DataSetIdentification3)
	return t.Identification
}

func (t *TerminalManagementDataSet10) SetSequenceCounter(value string) {
	t.SequenceCounter = (*Max9NumericText)(&value)
}

func (t *TerminalManagementDataSet10) AddContent() *ManagementPlanContent3 {
	t.Content = new(ManagementPlanContent3)
	return t.Content
}

// Data set containing the acceptor parameters of a point of interaction (POI).
type TerminalManagementDataSet11 struct {

	// Identification of the data set transferred.
	Identification *DataSetIdentification3 `xml:"Id"`

	// Counter to identify a single data set within the whole transfer.
	SequenceCounter *Max9NumericText `xml:"SeqCntr,omitempty"`

	// Content of the acceptor parameters.
	Content *AcceptorConfigurationContent3 `xml:"Cntt"`
}

func (t *TerminalManagementDataSet11) AddIdentification() *DataSetIdentification3 {
	t.Identification = new(DataSetIdentification3)
	return t.Identification
}

func (t *TerminalManagementDataSet11) SetSequenceCounter(value string) {
	t.SequenceCounter = (*Max9NumericText)(&value)
}

func (t *TerminalManagementDataSet11) AddContent() *AcceptorConfigurationContent3 {
	t.Content = new(AcceptorConfigurationContent3)
	return t.Content
}

// Identification of requested data set.
type TerminalManagementDataSet12 struct {

	// Identification of the required data set.
	Identification *DataSetIdentification4 `xml:"Id"`

	// Point of interaction challenge for cryptographic key injection.
	POIChallenge *Max140Binary `xml:"POIChllng,omitempty"`

	// Terminal manager challenge for cryptographic key injection.
	TMChallenge *Max140Binary `xml:"TMChllng,omitempty"`

	// Temporary encryption key that the host will use for protecting keys to download.
	SessionKey *CryptographicKey5 `xml:"SsnKey,omitempty"`
}

func (t *TerminalManagementDataSet12) AddIdentification() *DataSetIdentification4 {
	t.Identification = new(DataSetIdentification4)
	return t.Identification
}

func (t *TerminalManagementDataSet12) SetPOIChallenge(value string) {
	t.POIChallenge = (*Max140Binary)(&value)
}

func (t *TerminalManagementDataSet12) SetTMChallenge(value string) {
	t.TMChallenge = (*Max140Binary)(&value)
}

func (t *TerminalManagementDataSet12) AddSessionKey() *CryptographicKey5 {
	t.SessionKey = new(CryptographicKey5)
	return t.SessionKey
}

// Data related to the status report of a point of interaction (POI).
type TerminalManagementDataSet13 struct {

	// Identification of the data set containing the status report.
	Identification *DataSetIdentification4 `xml:"Id"`

	// Counter to identify a single data set within the whole transfer.
	SequenceCounter *Max9NumericText `xml:"SeqCntr,omitempty"`

	// Content of the status report.
	Content *StatusReportContent4 `xml:"Cntt"`
}

func (t *TerminalManagementDataSet13) AddIdentification() *DataSetIdentification4 {
	t.Identification = new(DataSetIdentification4)
	return t.Identification
}

func (t *TerminalManagementDataSet13) SetSequenceCounter(value string) {
	t.SequenceCounter = (*Max9NumericText)(&value)
}

func (t *TerminalManagementDataSet13) AddContent() *StatusReportContent4 {
	t.Content = new(StatusReportContent4)
	return t.Content
}

// Data set containing the acceptor parameters of a point of interaction (POI).
type TerminalManagementDataSet14 struct {

	// Identification of the data set transferred.
	Identification *DataSetIdentification4 `xml:"Id"`

	// Counter to identify a single data set within the whole transfer.
	SequenceCounter *Max9NumericText `xml:"SeqCntr,omitempty"`

	// Content of the acceptor parameters.
	Content *AcceptorConfigurationContent4 `xml:"Cntt"`
}

func (t *TerminalManagementDataSet14) AddIdentification() *DataSetIdentification4 {
	t.Identification = new(DataSetIdentification4)
	return t.Identification
}

func (t *TerminalManagementDataSet14) SetSequenceCounter(value string) {
	t.SequenceCounter = (*Max9NumericText)(&value)
}

func (t *TerminalManagementDataSet14) AddContent() *AcceptorConfigurationContent4 {
	t.Content = new(AcceptorConfigurationContent4)
	return t.Content
}

// Data related to the management plan of a point of interaction (POI).
type TerminalManagementDataSet15 struct {

	// Identification of the data set containing the management plan.
	Identification *DataSetIdentification4 `xml:"Id"`

	// Counter to identify a single data set within the whole transfer.
	SequenceCounter *Max9NumericText `xml:"SeqCntr,omitempty"`

	// Content of the management plan.
	Content *ManagementPlanContent4 `xml:"Cntt,omitempty"`
}

func (t *TerminalManagementDataSet15) AddIdentification() *DataSetIdentification4 {
	t.Identification = new(DataSetIdentification4)
	return t.Identification
}

func (t *TerminalManagementDataSet15) SetSequenceCounter(value string) {
	t.SequenceCounter = (*Max9NumericText)(&value)
}

func (t *TerminalManagementDataSet15) AddContent() *ManagementPlanContent4 {
	t.Content = new(ManagementPlanContent4)
	return t.Content
}

// Data related to the status report of a point of interaction (POI).
type TerminalManagementDataSet16 struct {

	// Identification of the data set containing the status report.
	Identification *DataSetIdentification6 `xml:"Id"`

	// Counter to identify a single data set within the whole transfer.
	SequenceCounter *Max9NumericText `xml:"SeqCntr,omitempty"`

	// Content of the status report.
	Content *StatusReportContent5 `xml:"Cntt"`
}

func (t *TerminalManagementDataSet16) AddIdentification() *DataSetIdentification6 {
	t.Identification = new(DataSetIdentification6)
	return t.Identification
}

func (t *TerminalManagementDataSet16) SetSequenceCounter(value string) {
	t.SequenceCounter = (*Max9NumericText)(&value)
}

func (t *TerminalManagementDataSet16) AddContent() *StatusReportContent5 {
	t.Content = new(StatusReportContent5)
	return t.Content
}

// Identification of requested data set.
type TerminalManagementDataSet17 struct {

	// Identification of the required data set.
	Identification *DataSetIdentification6 `xml:"Id"`

	// Point of interaction challenge for cryptographic key injection.
	POIChallenge *Max140Binary `xml:"POIChllng,omitempty"`

	// Terminal manager challenge for cryptographic key injection.
	TMChallenge *Max140Binary `xml:"TMChllng,omitempty"`

	// Temporary encryption key that the host will use for protecting keys to download.
	SessionKey *CryptographicKey5 `xml:"SsnKey,omitempty"`

	// Proof of delegation to be validated by the terminal manager receiving a status report from a new POI.
	DelegationProof *Max5000Binary `xml:"DlgtnProof,omitempty"`

	// Protected proof of delegation.
	ProtectedDelegationProof *ContentInformationType12 `xml:"PrtctdDlgtnProof,omitempty"`
}

func (t *TerminalManagementDataSet17) AddIdentification() *DataSetIdentification6 {
	t.Identification = new(DataSetIdentification6)
	return t.Identification
}

func (t *TerminalManagementDataSet17) SetPOIChallenge(value string) {
	t.POIChallenge = (*Max140Binary)(&value)
}

func (t *TerminalManagementDataSet17) SetTMChallenge(value string) {
	t.TMChallenge = (*Max140Binary)(&value)
}

func (t *TerminalManagementDataSet17) AddSessionKey() *CryptographicKey5 {
	t.SessionKey = new(CryptographicKey5)
	return t.SessionKey
}

func (t *TerminalManagementDataSet17) SetDelegationProof(value string) {
	t.DelegationProof = (*Max5000Binary)(&value)
}

func (t *TerminalManagementDataSet17) AddProtectedDelegationProof() *ContentInformationType12 {
	t.ProtectedDelegationProof = new(ContentInformationType12)
	return t.ProtectedDelegationProof
}

// Data related to the management plan of a point of interaction (POI).
type TerminalManagementDataSet18 struct {

	// Identification of the data set containing the management plan.
	Identification *DataSetIdentification6 `xml:"Id"`

	// Counter to identify a single data set within the whole transfer.
	SequenceCounter *Max9NumericText `xml:"SeqCntr,omitempty"`

	// Content of the management plan.
	Content *ManagementPlanContent5 `xml:"Cntt,omitempty"`
}

func (t *TerminalManagementDataSet18) AddIdentification() *DataSetIdentification6 {
	t.Identification = new(DataSetIdentification6)
	return t.Identification
}

func (t *TerminalManagementDataSet18) SetSequenceCounter(value string) {
	t.SequenceCounter = (*Max9NumericText)(&value)
}

func (t *TerminalManagementDataSet18) AddContent() *ManagementPlanContent5 {
	t.Content = new(ManagementPlanContent5)
	return t.Content
}

// Data set containing the acceptor parameters of a point of interaction (POI).
type TerminalManagementDataSet19 struct {

	// Identification of the data set transferred.
	Identification *DataSetIdentification6 `xml:"Id"`

	// Counter to identify a single data set within the whole transfer.
	SequenceCounter *Max9NumericText `xml:"SeqCntr,omitempty"`

	// Identification of the point of interactions involved by the configuration data set.
	POIIdentification []*GenericIdentification71 `xml:"POIId,omitempty"`

	// Scope of the configuration contained in the data set.
	ConfigurationScope *PartyType15Code `xml:"CfgtnScp,omitempty"`

	// Content of the acceptor parameters.
	Content *AcceptorConfigurationContent5 `xml:"Cntt"`
}

func (t *TerminalManagementDataSet19) AddIdentification() *DataSetIdentification6 {
	t.Identification = new(DataSetIdentification6)
	return t.Identification
}

func (t *TerminalManagementDataSet19) SetSequenceCounter(value string) {
	t.SequenceCounter = (*Max9NumericText)(&value)
}

func (t *TerminalManagementDataSet19) AddPOIIdentification() *GenericIdentification71 {
	newValue := new(GenericIdentification71)
	t.POIIdentification = append(t.POIIdentification, newValue)
	return newValue
}

func (t *TerminalManagementDataSet19) SetConfigurationScope(value string) {
	t.ConfigurationScope = (*PartyType15Code)(&value)
}

func (t *TerminalManagementDataSet19) AddContent() *AcceptorConfigurationContent5 {
	t.Content = new(AcceptorConfigurationContent5)
	return t.Content
}

// Data related to the management plan of a point of interaction (POI).
type TerminalManagementDataSet2 struct {

	// Identification of the data set containing the management plan.
	Identification *DataSetIdentification2 `xml:"Id"`

	// Counter to identify a single data set within the whole transfer.
	SequenceCounter *Max9NumericText `xml:"SeqCntr,omitempty"`

	// Content of the management plan.
	Content *ManagementPlanContent1 `xml:"Cntt,omitempty"`
}

func (t *TerminalManagementDataSet2) AddIdentification() *DataSetIdentification2 {
	t.Identification = new(DataSetIdentification2)
	return t.Identification
}

func (t *TerminalManagementDataSet2) SetSequenceCounter(value string) {
	t.SequenceCounter = (*Max9NumericText)(&value)
}

func (t *TerminalManagementDataSet2) AddContent() *ManagementPlanContent1 {
	t.Content = new(ManagementPlanContent1)
	return t.Content
}

// Data set containing the acceptor parameters of a point of interaction (POI).
type TerminalManagementDataSet3 struct {

	// Identification of the data set transferred.
	Identification *DataSetIdentification2 `xml:"Id"`

	// Counter to identify a single data set within the whole transfer.
	SequenceCounter *Max9NumericText `xml:"SeqCntr,omitempty"`

	// Content of the acceptor parameters.
	Content *AcceptorConfigurationContent1 `xml:"Cntt"`
}

func (t *TerminalManagementDataSet3) AddIdentification() *DataSetIdentification2 {
	t.Identification = new(DataSetIdentification2)
	return t.Identification
}

func (t *TerminalManagementDataSet3) SetSequenceCounter(value string) {
	t.SequenceCounter = (*Max9NumericText)(&value)
}

func (t *TerminalManagementDataSet3) AddContent() *AcceptorConfigurationContent1 {
	t.Content = new(AcceptorConfigurationContent1)
	return t.Content
}

// Data related to the status report of a point of interaction (POI).
type TerminalManagementDataSet4 struct {

	// Identification of the data set containing the status report.
	Identification *DataSetIdentification3 `xml:"Id"`

	// Counter to identify a single data set within the whole transfer.
	SequenceCounter *Max9NumericText `xml:"SeqCntr,omitempty"`

	// Content of the status report.
	Content *StatusReportContent2 `xml:"Cntt"`
}

func (t *TerminalManagementDataSet4) AddIdentification() *DataSetIdentification3 {
	t.Identification = new(DataSetIdentification3)
	return t.Identification
}

func (t *TerminalManagementDataSet4) SetSequenceCounter(value string) {
	t.SequenceCounter = (*Max9NumericText)(&value)
}

func (t *TerminalManagementDataSet4) AddContent() *StatusReportContent2 {
	t.Content = new(StatusReportContent2)
	return t.Content
}

// Data related to the management plan of a point of interaction (POI).
type TerminalManagementDataSet5 struct {

	// Identification of the data set containing the management plan.
	Identification *DataSetIdentification3 `xml:"Id"`

	// Counter to identify a single data set within the whole transfer.
	SequenceCounter *Max9NumericText `xml:"SeqCntr,omitempty"`

	// Content of the management plan.
	Content *ManagementPlanContent2 `xml:"Cntt,omitempty"`
}

func (t *TerminalManagementDataSet5) AddIdentification() *DataSetIdentification3 {
	t.Identification = new(DataSetIdentification3)
	return t.Identification
}

func (t *TerminalManagementDataSet5) SetSequenceCounter(value string) {
	t.SequenceCounter = (*Max9NumericText)(&value)
}

func (t *TerminalManagementDataSet5) AddContent() *ManagementPlanContent2 {
	t.Content = new(ManagementPlanContent2)
	return t.Content
}

// Data set containing the acceptor parameters of a point of interaction (POI).
type TerminalManagementDataSet6 struct {

	// Identification of the data set transferred.
	Identification *DataSetIdentification3 `xml:"Id"`

	// Counter to identify a single data set within the whole transfer.
	SequenceCounter *Max9NumericText `xml:"SeqCntr,omitempty"`

	// Content of the acceptor parameters.
	Content *AcceptorConfigurationContent2 `xml:"Cntt"`
}

func (t *TerminalManagementDataSet6) AddIdentification() *DataSetIdentification3 {
	t.Identification = new(DataSetIdentification3)
	return t.Identification
}

func (t *TerminalManagementDataSet6) SetSequenceCounter(value string) {
	t.SequenceCounter = (*Max9NumericText)(&value)
}

func (t *TerminalManagementDataSet6) AddContent() *AcceptorConfigurationContent2 {
	t.Content = new(AcceptorConfigurationContent2)
	return t.Content
}

// Identification of requested data set.
type TerminalManagementDataSet7 struct {

	// Identification of the required data set.
	Identification *DataSetIdentification3 `xml:"Id"`

	// Point of interaction challenge for cryptographic key injection.
	POIChallenge *Max140Binary `xml:"POIChllng,omitempty"`

	// Terminal manager challenge for cryptographic key injection.
	TMChallenge *Max140Binary `xml:"TMChllng,omitempty"`

	// Transport key encrypted by the TM key encryption RSA key.
	EncryptedKey *ContentInformationType5 `xml:"NcrptdKey,omitempty"`
}

func (t *TerminalManagementDataSet7) AddIdentification() *DataSetIdentification3 {
	t.Identification = new(DataSetIdentification3)
	return t.Identification
}

func (t *TerminalManagementDataSet7) SetPOIChallenge(value string) {
	t.POIChallenge = (*Max140Binary)(&value)
}

func (t *TerminalManagementDataSet7) SetTMChallenge(value string) {
	t.TMChallenge = (*Max140Binary)(&value)
}

func (t *TerminalManagementDataSet7) AddEncryptedKey() *ContentInformationType5 {
	t.EncryptedKey = new(ContentInformationType5)
	return t.EncryptedKey
}

// Identification of requested data set.
type TerminalManagementDataSet8 struct {

	// Identification of the required data set.
	Identification *DataSetIdentification3 `xml:"Id"`

	// Point of interaction challenge for cryptographic key injection.
	POIChallenge *Max140Binary `xml:"POIChllng,omitempty"`

	// Terminal manager challenge for cryptographic key injection.
	TMChallenge *Max140Binary `xml:"TMChllng,omitempty"`

	// Transport key encrypted by the TM (Terminal manager) key encryption RSA key.
	EncryptedKey *ContentInformationType7 `xml:"NcrptdKey,omitempty"`
}

func (t *TerminalManagementDataSet8) AddIdentification() *DataSetIdentification3 {
	t.Identification = new(DataSetIdentification3)
	return t.Identification
}

func (t *TerminalManagementDataSet8) SetPOIChallenge(value string) {
	t.POIChallenge = (*Max140Binary)(&value)
}

func (t *TerminalManagementDataSet8) SetTMChallenge(value string) {
	t.TMChallenge = (*Max140Binary)(&value)
}

func (t *TerminalManagementDataSet8) AddEncryptedKey() *ContentInformationType7 {
	t.EncryptedKey = new(ContentInformationType7)
	return t.EncryptedKey
}

// Data related to the status report of a point of interaction (POI).
type TerminalManagementDataSet9 struct {

	// Identification of the data set containing the status report.
	Identification *DataSetIdentification3 `xml:"Id"`

	// Counter to identify a single data set within the whole transfer.
	SequenceCounter *Max9NumericText `xml:"SeqCntr,omitempty"`

	// Content of the status report.
	Content *StatusReportContent3 `xml:"Cntt"`
}

func (t *TerminalManagementDataSet9) AddIdentification() *DataSetIdentification3 {
	t.Identification = new(DataSetIdentification3)
	return t.Identification
}

func (t *TerminalManagementDataSet9) SetSequenceCounter(value string) {
	t.SequenceCounter = (*Max9NumericText)(&value)
}

func (t *TerminalManagementDataSet9) AddContent() *StatusReportContent3 {
	t.Content = new(StatusReportContent3)
	return t.Content
}
