package models

// Photo adalah sebuah struct yang merepresentasikan data foto.
type Photo struct {
	Id       string `gorm:"primaryKey;not null" valid:"required" json:"id"`
	Title    string `gorm:"type:text;not null" json:"title"`
	Caption  string `gorm:"type:text;not null" json:"caption"`
	PhotoUrl string `gorm:"type:text;not null" json:"photourl"`
	Userid   string `gorm:"index" json:"userid" form:"userid"`
}
