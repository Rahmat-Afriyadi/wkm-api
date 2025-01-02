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

type MasterKodepos1 struct {
	KdPos     string `json:"kd_pos" gorm:"column:kd_pos"`
	Kodepos   string `json:"kodepos" gorm:"column:kodepos"`
	Kelurahan string `json:"kelurahan" gorm:"column:kelurahan"`
	Kecamatan string `json:"kecamatan" gorm:"column:kecamatan"`
	Kota      string `json:"kota" gorm:"column:kota"`
}

func (MasterKodepos1) TableName() string {
	return "kodepos"
}
