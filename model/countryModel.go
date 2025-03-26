package model

type Country struct {
	Idx  int    `gorm:"column:idx;primary_key;AUTO_INCREMENT"`
	Name string `gorm:"column:name;NOT NULL"`
	// Landmarks []Landmark `gorm:"foreignkey:CountryID"` // เชื่อมโยงกับฟิลด์ CountryID ในโมเดล Landmark
}

func (m *Country) TableName() string {
	return "country"
}
