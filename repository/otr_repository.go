package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
	"wkm/entity"
	"wkm/request"
	"wkm/utils"

	"gorm.io/gorm"
)

type OtrRepository interface {
	DetailOtrNa(motorprice_kode string, tahun uint16) entity.Otr
	OtrNaList() []entity.Otr
	OtrMstProduk(search string) []entity.MstMtr
	OtrMstNa(search string) []entity.OtrNa
	CreateOtr(data request.CreateOtr)
	MasterData(search string, limit int, pageParams int) []entity.Otr
	MasterDataCount(search string) int64
	DetailOtr(id string) entity.Otr
	Update(body entity.Otr) error
	ListApi()
}

type otrRepository struct {
	conn *gorm.DB
}

func NewOtrRepository(conn *gorm.DB) OtrRepository {
	return &otrRepository{
		conn: conn,
	}
}

func (lR *otrRepository) DetailOtr(id string) entity.Otr {
	otr := entity.Otr{ID: id}
	lR.conn.Find(&otr)
	return otr
}

func (lR *otrRepository) DetailOtrNa(motorprice_kode string, tahun uint16) entity.Otr {
	otr := entity.Otr{}
	lR.conn.Table("otr_na a").Joins("inner join mst_mtr b on a.motorprice_kode = b.kd_mdl").Select("a.tahun, a.motorprice_kode, b.nm_mtr product_nama").Where("a.motorprice_kode = ? and a.tahun = ? ", motorprice_kode, tahun).First(&otr)
	return otr
}

func (lR *otrRepository) CreateOtr(data request.CreateOtr) {
	otr := entity.Otr{
		KdMdl:       data.MotorpriceKode,
		ProductKode: data.ProductKode,
		ProductNama: data.ProductNama,
		Otr:         data.Otr,
		Tahun:       data.Tahun,
		WrnKode:     data.WrnKode,
	}
	result := lR.conn.Create(&otr)
	if result.Error != nil {
		fmt.Println("ini errornya yaa ", result.Error)
	} else {
		if data.CreateFrom == "otrna" {
			lR.conn.Where("tahun = ? and motorprice_kode = ?", data.Tahun, data.MotorpriceKode).Delete(&entity.MstOtrNa{})
		}
	}
}

func (lR *otrRepository) Update(data entity.Otr) error {
	record := entity.Otr{ID: data.ID}
	lR.conn.Find(&record)
	if record.KdMdl == "" {
		return errors.New("data tidak ditemukan")
	}
	result := lR.conn.Save(&data)
	if result.Error != nil {
		fmt.Println("ini error ", result.Error)
		return result.Error
	} else {
		return nil
	}
}

func (lR *otrRepository) MasterData(search string, limit int, pageParams int) []entity.Otr {
	otr := []entity.Otr{}
	query := lR.conn.Where("motorprice_kode like ? or product_nama like ?  or product_kode like ? ", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	query.Scopes(utils.Paginate(&utils.PaginateParams{PageParams: pageParams, Limit: limit})).Preload("MstMtr").Find(&otr)
	return otr
}

func (lR *otrRepository) MasterDataCount(search string) int64 {
	var otr []entity.Otr
	query := lR.conn.Where("motorprice_kode like ? or product_nama like ? or product_kode like ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	query.Select("motorprice_kode").Find(&otr)
	return int64(len(otr))
}

func (lR *otrRepository) OtrNaList() []entity.Otr {
	var otr []entity.Otr
	lR.conn.Table("otr_na").Group("motorprice_kode, tahun").Find(&otr)
	return otr
}

func (lR *otrRepository) OtrMstProduk(search string) []entity.MstMtr {
	var otr []entity.MstMtr
	lR.conn.Where("nm_mtr like ? or no_mtr like ? or kd_mdl like ?", "%"+search+"%", "%"+search+"%", "%"+search+"%").Limit(15).Find(&otr)
	return otr
}
func (lR *otrRepository) OtrMstNa(search string) []entity.OtrNa {
	var otr []entity.OtrNa
	lR.conn.Raw("select a.*, m.nm_mtr from (select motorprice_kode, tahun from otr_na group by motorprice_kode, tahun) a left join mst_mtr m  on m.kd_mdl = a.motorprice_kode where a.motorprice_kode like ? or a.tahun like ? or m.nm_mtr like ? limit 15 ", "%"+search+"%", "%"+search+"%", "%"+search+"%").Find(&otr)
	return otr
}

func (lR *otrRepository) ListApi() {
	url := "http://119.8.172.66/api/dealer/api_price_motor_mdm.php" // Replace with your URL
	token := "71D55F9957529"                                        // Replace with your Bearer token

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error performing request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: Status code", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var responseObject entity.ResponseOtr
	tahun := time.Now().Year()
	json.Unmarshal(body, &responseObject)
	for _, data := range responseObject.Data {
		parts := strings.Split(data.OtrApi, ".")
		num, err := strconv.ParseUint(parts[0], 10, 64)
		if num == 0 {
			continue
		}
		var otr entity.Otr
		lR.conn.Where("motorprice_kode = ? and tahun = ?", data.KdMdl, tahun).First(&otr)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return
		}
		now := time.Now()
		otr.UpdatedAt = &now
		if otr.ID != "" {
			otr.Otr = num
			lR.conn.Save(&otr)
		} else {
			data.Tahun = uint16(tahun)
			data.Otr = num
			result := lR.conn.Create(&data)
			if result.Error != nil {
				fmt.Println("ini error create yaa ", result.Error)
			} else {
				fmt.Println("ini berhasil create yaa")
			}
		}
	}
	fmt.Println(len(responseObject.Data))
}
