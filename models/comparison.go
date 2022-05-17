package models

type Comparison struct {
	PhoneSatu string `json:"phone1" form:"phone1"`
	PhoneDua  string `json:"phone2" form:"phone2"`
}
