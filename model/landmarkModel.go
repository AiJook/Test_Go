package model

type Landmark struct {
	Idx       int     `gorm:"column:idx;primary_key;AUTO_INCREMENT"`
	Name      string  `gorm:"column:name;NOT NULL"`
	CountryID int     `gorm:"column:country;type:int(11);NOT NULL"`
	Detail    string  `gorm:"column:detail;NOT NULL"`
	Url       string  `gorm:"column:url;NOT NULL"`
	Country   Country `gorm:"foreignkey:CountryID"`
}

func (m *Landmark) TableName() string {
	return "landmark"
}
