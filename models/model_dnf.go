package models

type Dnf struct {
	DnfUnits []DnfUnit `json:"dnfUints" bson:"dnfUnits"`
}
