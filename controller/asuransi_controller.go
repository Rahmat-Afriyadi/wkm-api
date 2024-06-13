package controller

import (
	"fmt"
	"strconv"
	"wkm/entity"
	"wkm/request"
	"wkm/service"

	"github.com/gofiber/fiber/v2"
)

type AsuransiController interface {
	MasterData(ctx *fiber.Ctx) error
	MasterDataCount(ctx *fiber.Ctx) error
	RekapByStatusKdUser(ctx *fiber.Ctx) error
	FindAsuransiByNoMsn(ctx *fiber.Ctx) error
	UpdateAsuransi(ctx *fiber.Ctx) error
	UpdateAsuransiBerminat(ctx *fiber.Ctx) error
	UpdateAsuransiBatalBayar(ctx *fiber.Ctx) error
	UpdateAmbilAsuransi(ctx *fiber.Ctx) error
	MasterDataRekapTele(ctx *fiber.Ctx) error
	RekapByStatus(ctx *fiber.Ctx) error
	RekapByStatusLt(ctx *fiber.Ctx) error
	MasterAlasanPending(ctx *fiber.Ctx) error
	MasterAlasanTdkBerminat(ctx *fiber.Ctx) error
	ExportReportAsuransi(ctx *fiber.Ctx) error
	DetailApprovalTransaksi(ctx *fiber.Ctx) error
	ListApprovalTransaksi(ctx *fiber.Ctx) error
	ListApprovalTransaksiCount(ctx *fiber.Ctx) error
}

type asuransiController struct {
	asuransiService service.AsuransiService
}

func NewAsuransiController(aS service.AsuransiService) AsuransiController {
	return &asuransiController{
		asuransiService: aS,
	}
}

func (tr *asuransiController) ExportReportAsuransi(ctx *fiber.Ctx) error {
	var request request.ReportAsuransi
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(400).JSON(err)
	}
	fmt.Println("ini request ", request)
	tr.asuransiService.ExportReport(request.AwalTanggal, request.AkhirTAnggal)
	return ctx.Download("./file-report-asuransi.xlsx")

}

func (tr *asuransiController) ListApprovalTransaksi(ctx *fiber.Ctx) error {
	tgl1 := ctx.Query("tgl1")
	tgl2 := ctx.Query("tgl2")
	search := ctx.Query("search")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	userParams := ""
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	if details.RoleId == 1 {
		userParams = details.Username
	}
	pageParams, _ := strconv.Atoi(ctx.Query("pageParams"))
	return ctx.JSON(tr.asuransiService.ListApprovalTransaksi(userParams, tgl1, tgl2, search, pageParams, limit))
}
func (tr *asuransiController) ListApprovalTransaksiCount(ctx *fiber.Ctx) error {
	tgl1 := ctx.Query("tgl1")
	tgl2 := ctx.Query("tgl2")
	search := ctx.Query("search")
	userParams := ""
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	if details.RoleId == 1 {
		userParams = details.Username
	}
	return ctx.JSON(tr.asuransiService.ListApprovalTransaksiCount(userParams, tgl1, tgl2, search))
}

func (tr *asuransiController) MasterDataRekapTele(ctx *fiber.Ctx) error {
	return ctx.JSON(tr.asuransiService.MasterDataRekapTele())
}

func (tr *asuransiController) DetailApprovalTransaksi(ctx *fiber.Ctx) error {
	idTrx := ctx.Params("idTrx")
	return ctx.JSON(tr.asuransiService.DetailApprovalTransaksi(idTrx))
}

func (tr *asuransiController) MasterAlasanPending(ctx *fiber.Ctx) error {
	return ctx.JSON(tr.asuransiService.MasterAlasanPending())
}
func (tr *asuransiController) MasterAlasanTdkBerminat(ctx *fiber.Ctx) error {
	return ctx.JSON(tr.asuransiService.MasterAlasanTdkBerminat())
}

func (tr *asuransiController) MasterData(ctx *fiber.Ctx) error {
	dataSource := ctx.Query("dataSource")
	sts := ctx.Params("status")
	search := ctx.Query("search")
	tgl1 := ctx.Query("tgl1")
	tgl2 := ctx.Query("tgl2")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	pageParams, _ := strconv.Atoi(ctx.Query("pageParams"))
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	return ctx.JSON(tr.asuransiService.MasterData(search, dataSource, sts, details.Username, tgl1, tgl2, limit, pageParams))
}

func (tr *asuransiController) RekapByStatusKdUser(ctx *fiber.Ctx) error {
	tgl1 := ctx.Query("tgl1")
	tgl2 := ctx.Query("tgl2")
	return ctx.JSON(tr.asuransiService.RekapByStatusKdUser(tgl1, tgl2))
}

func (tr *asuransiController) MasterDataCount(ctx *fiber.Ctx) error {
	dataSource := ctx.Query("dataSource")
	sts := ctx.Params("status")
	search := ctx.Query("search")
	tgl1 := ctx.Query("tgl1")
	tgl2 := ctx.Query("tgl2")
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	return ctx.JSON(tr.asuransiService.MasterDataCount(search, dataSource, sts, details.Username, tgl1, tgl2))
}

func (tr *asuransiController) FindAsuransiByNoMsn(ctx *fiber.Ctx) error {
	no_msn := ctx.Params("no_msn")
	return ctx.JSON(tr.asuransiService.FindAsuransiByNoMsn(no_msn))

}

func (tr *asuransiController) UpdateAsuransi(ctx *fiber.Ctx) error {
	var asuransi entity.MasterAsuransi
	err := ctx.BodyParser(&asuransi)
	fmt.Println("ini body ", asuransi)
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	asuransi.KdUser = details.Username
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	tr.asuransiService.UpdateAsuransi(asuransi)
	return ctx.JSON("Hallo guys")

}

func (tr *asuransiController) RekapByStatus(ctx *fiber.Ctx) error {
	user := ctx.Locals("user")
	tgl1 := ctx.Query("tgl1")
	tgl2 := ctx.Query("tgl2")
	details, _ := user.(entity.User)
	result := tr.asuransiService.RekapByStatus(details.Username, tgl1, tgl2)
	return ctx.JSON(result)
}

func (tr *asuransiController) RekapByStatusLt(ctx *fiber.Ctx) error {
	tgl1 := ctx.Query("tgl1")
	tgl2 := ctx.Query("tgl2")
	result := tr.asuransiService.RekapByStatus("", tgl1, tgl2)
	return ctx.JSON(result)
}

func (tr *asuransiController) UpdateAsuransiBerminat(ctx *fiber.Ctx) error {
	var body map[string]interface{}
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	fmt.Println("ini body ", body["no_msn"])
	tr.asuransiService.UpdateAsuransiBerminat(body["no_msn"].(string))
	return ctx.JSON("Hallo guys")

}

func (tr *asuransiController) UpdateAsuransiBatalBayar(ctx *fiber.Ctx) error {
	var body map[string]interface{}
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	fmt.Println("ini body ", body["no_msn"])
	tr.asuransiService.UpdateAsuransiBatalBayar(body["no_msn"].(string))
	return ctx.JSON("Hallo guys")

}

func (tr *asuransiController) UpdateAmbilAsuransi(ctx *fiber.Ctx) error {
	var asuransi entity.MasterAsuransi
	err := ctx.BodyParser(&asuransi)
	fmt.Println("ini body ", asuransi)
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	asuransi.KdUser = details.Username
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	tr.asuransiService.UpdateAmbilAsuransi(asuransi.NoMsn, asuransi.KdUser)
	return ctx.JSON("Berhasil")

}
