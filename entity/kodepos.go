package entity

type MasterKodepos struct {
	Id              string `json:"id" gorm:"column:id"`
	Province        string `json:"province" gorm:"column:province"`
	ProvinceCode    string `json:"province_code" gorm:"column:province_code"`
	City            string `json:"city" gorm:"column:city"`
	CityCode        string `json:"city_code" gorm:"column:city_code"`
	Subdistrict     string `json:"subdistrict" gorm:"column:subdistrict"`
	SubdistrictCode string `json:"subdistrict_code" gorm:"column:subdistrict_code"`
}

func (MasterKodepos) TableName() string {
	return "kota"
}
