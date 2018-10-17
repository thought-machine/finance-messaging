package iso20022

// Charge or commission of the original individual order details that have been repaired so that the order can be accepted.
type RepairedConditions2 struct {

	// Charge from the original individual order details that has been repaired so that the order can be accepted.
	RepairedCharge []*Charge11 `xml:"RprdChrg,omitempty"`

	// Commission from the original individual order details that has been repaired so that the order can be accepted.
	RepairedCommission []*Commission7 `xml:"RprdComssn,omitempty"`
}

func (r *RepairedConditions2) AddRepairedCharge() *Charge11 {
	newValue := new(Charge11)
	r.RepairedCharge = append(r.RepairedCharge, newValue)
	return newValue
}

func (r *RepairedConditions2) AddRepairedCommission() *Commission7 {
	newValue := new(Commission7)
	r.RepairedCommission = append(r.RepairedCommission, newValue)
	return newValue
}

// Charge or commission of the original individual order details that have been repaired so that the order can be accepted.
type RepairedConditions3 struct {

	// Modified value of the charge applied on the order (the charge in the original individual order that has been repaired so that the order can be accepted).
	RepairedCharge []*Charge19 `xml:"RprdChrg,omitempty"`

	// Modified value of the commission applied on the order (the commission in the original individual order that has been repaired so that the order can be accepted).
	RepairedCommission []*Commission11 `xml:"RprdComssn,omitempty"`
}

func (r *RepairedConditions3) AddRepairedCharge() *Charge19 {
	newValue := new(Charge19)
	r.RepairedCharge = append(r.RepairedCharge, newValue)
	return newValue
}

func (r *RepairedConditions3) AddRepairedCommission() *Commission11 {
	newValue := new(Commission11)
	r.RepairedCommission = append(r.RepairedCommission, newValue)
	return newValue
}
