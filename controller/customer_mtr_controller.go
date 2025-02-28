package controller

import (
	"strconv"
	"time"
	"wkm/entity"
	"wkm/request"
	"wkm/service"

	"github.com/gofiber/fiber/v2"
)


type CustomerMtrController interface {
	MasterData(ctx *fiber.Ctx) error
	MasterDataCount(ctx *fiber.Ctx) error
	ListAmbilData(ctx *fiber.Ctx) error
	AmbilData(ctx *fiber.Ctx) error
	Show(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	UpdateOkeMembership(ctx *fiber.Ctx) error
	RekapTele(ctx *fiber.Ctx) error
	ExportRekapTele(ctx *fiber.Ctx) error
	ListBerminatMembership(ctx *fiber.Ctx) error
	ListDataAsuransiPA(ctx *fiber.Ctx) error
	ListDataAsuransiMtr(ctx *fiber.Ctx) error
}

type customerMtrController struct {
	customerMtrService service.CustomerMtrService
}

func NewCustomerMtrController(aS service.CustomerMtrService) CustomerMtrController {
	return &customerMtrController{
		customerMtrService: aS,
	}
}

func (tr *customerMtrController) MasterData(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	sts := ctx.Query("sts")
	jns := ctx.Query("jns")
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	pageParams, _ := strconv.Atoi(ctx.Query("pageParams"))
	data := tr.customerMtrService.MasterData(search, sts, jns, details.Username, limit, pageParams)
	return ctx.Status(200).JSON(data)
}
func (tr *customerMtrController) MasterDataCount(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	sts := ctx.Query("sts")
	jns := ctx.Query("jns")
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	data := tr.customerMtrService.MasterDataCount(search, sts, jns, details.Username)
	return ctx.Status(200).JSON(data)
}
func (tr *customerMtrController) ListAmbilData(ctx *fiber.Ctx) error {
	data := tr.customerMtrService.ListAmbilData()
	return ctx.Status(200).JSON(fiber.Map{"status":"success","data": data})
}

func (tr *customerMtrController) AmbilData(ctx *fiber.Ctx) error {
	var request entity.Faktur3
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request body", "details": err.Error()})
	}
	err := tr.customerMtrService.AmbilData(request.NoMsn, details.Username)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"status":"fail", "message": err.Error()})
	}
	return ctx.Status(200).JSON(fiber.Map{"message": "Berhasil Ambil Data"})
}

func (tr *customerMtrController) Show(ctx *fiber.Ctx) error {
	noMsn := ctx.Params("no_msn")
	data := tr.customerMtrService.Show(noMsn)
	return ctx.Status(200).JSON(fiber.Map{"message": "Berhasil ", "data":data})
}

func (tr *customerMtrController) Update(ctx *fiber.Ctx) error {
	var request entity.CustomerMtr
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request body", "details": err.Error()})
	}
	
	return nil
}

func (tr *customerMtrController) UpdateOkeMembership(ctx *fiber.Ctx) error {
	var request request.CustomerMtr
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request body", "details": err.Error()})
	}
	user := ctx.Locals("user")
	details,_:= user.(entity.User)

	request.KdUserTs = details.Username
	customer, err := tr.customerMtrService.UpdateOkeMembership(request)
	if err != nil {
		return  ctx.Status(400).JSON(fiber.Map{"error": "Invalid request body", "details": err.Error()})
 
	}
	return ctx.Status(200).JSON(fiber.Map{"message": "Berhasil ", "data":customer})
}

func (tr *customerMtrController) RekapTele(ctx *fiber.Ctx) error {
	var rekapReq request.RangeTanggalRequest

	// Parsing request body ke struct
	if err := ctx.BodyParser(&rekapReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
	}

	// Format tanggal yang digunakan
	layoutFull := "2006-01-02 15:04:05"
	layoutDate := "2006-01-02"

	// Load lokasi waktu Jakarta
	loc, _ := time.LoadLocation("Asia/Jakarta")

	// Default StartDate dan EndDate (tanggal hari ini dalam zona waktu Jakarta)
	now := time.Now().In(loc)
	startDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	endDate := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, loc)

	// Parsing Tgl1 jika ada
	if rekapReq.Tgl1 != "" {
		parsedStart, err := time.ParseInLocation(layoutFull, rekapReq.Tgl1, loc)
		if err != nil {
			parsedStart, err = time.ParseInLocation(layoutDate, rekapReq.Tgl1, loc)
			if err == nil {
				parsedStart = time.Date(parsedStart.Year(), parsedStart.Month(), parsedStart.Day(), 0, 0, 0, 0, loc)
			}
		}
		if err == nil {
			startDate = parsedStart
		}
	}

	// Parsing Tgl2 jika ada
	if rekapReq.Tgl2 != "" {
		parsedEnd, err := time.ParseInLocation(layoutFull, rekapReq.Tgl2, loc)
		if err != nil {
			parsedEnd, err = time.ParseInLocation(layoutDate, rekapReq.Tgl2, loc)
			if err == nil {
				parsedEnd = time.Date(parsedEnd.Year(), parsedEnd.Month(), parsedEnd.Day(), 23, 59, 59, 0, loc)
			}
		}
		if err == nil {
			endDate = parsedEnd
		}
	}

	// Mengambil informasi user dari context
	user := ctx.Locals("user")
	details, ok := user.(entity.User)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Unauthorized",
			"details": "Invalid user context",
		})
	}

	// Memanggil service untuk mendapatkan data rekap
	rekapData, err := tr.customerMtrService.RekapTele(details.Username, startDate, endDate)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Failed to fetch data",
			"details": err.Error(),
		})
	}

	// Response sukses dengan format hanya "YYYY-MM-DD"
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil",
		"user":    details.Username,
		"data":    rekapData,
		"Dari":    startDate.Format(layoutDate),
		"Sampai":  endDate.Format(layoutDate),
	})
}

func (tr *customerMtrController) ExportRekapTele(ctx *fiber.Ctx) error {
	var rekapReq request.RangeTanggalRequest

	// Parsing request body ke struct
	if err := ctx.BodyParser(&rekapReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
	}

	// Format tanggal yang digunakan
	layoutFull := "2006-01-02 15:04:05"
	layoutDate := "2006-01-02"

	// Load lokasi waktu Jakarta
	loc, _ := time.LoadLocation("Asia/Jakarta")

	// Default StartDate dan EndDate (tanggal hari ini dalam zona waktu Jakarta)
	now := time.Now().In(loc)
	startDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	endDate := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, loc)

	// Parsing Tgl1 jika ada
	if rekapReq.Tgl1 != "" {
		parsedStart, err := time.ParseInLocation(layoutFull, rekapReq.Tgl1, loc)
		if err != nil {
			parsedStart, err = time.ParseInLocation(layoutDate, rekapReq.Tgl1, loc)
			if err == nil {
				parsedStart = time.Date(parsedStart.Year(), parsedStart.Month(), parsedStart.Day(), 0, 0, 0, 0, loc)
			}
		}
		if err == nil {
			startDate = parsedStart
		}
	}

	// Parsing Tgl2 jika ada
	if rekapReq.Tgl2 != "" {
		parsedEnd, err := time.ParseInLocation(layoutFull, rekapReq.Tgl2, loc)
		if err != nil {
			parsedEnd, err = time.ParseInLocation(layoutDate, rekapReq.Tgl2, loc)
			if err == nil {
				parsedEnd = time.Date(parsedEnd.Year(), parsedEnd.Month(), parsedEnd.Day(), 23, 59, 59, 0, loc)
			}
		}
		if err == nil {
			endDate = parsedEnd
		}
	}

	// Mengambil informasi user dari context
	user := ctx.Locals("user")
	details, ok := user.(entity.User)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Unauthorized",
			"details": "Invalid user context",
		})
	}

	// Memanggil service untuk mendapatkan file Excel
	fileName, err := tr.customerMtrService.ExportRekapTele(details.Username, startDate, endDate)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to generate report",
			"details": err.Error(),
		})
	}

	return ctx.Download(fileName)
}

func (tr *customerMtrController) ListBerminatMembership(ctx *fiber.Ctx) error {
	var rekapReq request.RangeTanggalRequest

	// Parsing request body ke struct
	if err := ctx.BodyParser(&rekapReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
	}

	// Format tanggal yang digunakan
	layoutFull := "2006-01-02 15:04:05"
	layoutDate := "2006-01-02"

	// Load lokasi waktu Jakarta
	loc, _ := time.LoadLocation("Asia/Jakarta")

	// Default StartDate dan EndDate (tanggal hari ini dalam zona waktu Jakarta)
	now := time.Now().In(loc)
	startDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	endDate := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, loc)

	// Parsing Tgl1 jika ada
	if rekapReq.Tgl1 != "" {
		parsedStart, err := time.ParseInLocation(layoutFull, rekapReq.Tgl1, loc)
		if err != nil {
			parsedStart, err = time.ParseInLocation(layoutDate, rekapReq.Tgl1, loc)
			if err == nil {
				parsedStart = time.Date(parsedStart.Year(), parsedStart.Month(), parsedStart.Day(), 0, 0, 0, 0, loc)
			}
		}
		if err == nil {
			startDate = parsedStart
		}
	}

	// Parsing Tgl2 jika ada
	if rekapReq.Tgl2 != "" {
		parsedEnd, err := time.ParseInLocation(layoutFull, rekapReq.Tgl2, loc)
		if err != nil {
			parsedEnd, err = time.ParseInLocation(layoutDate, rekapReq.Tgl2, loc)
			if err == nil {
				parsedEnd = time.Date(parsedEnd.Year(), parsedEnd.Month(), parsedEnd.Day(), 23, 59, 59, 0, loc)
			}
		}
		if err == nil {
			endDate = parsedEnd
		}
	}

	// Mengambil informasi user dari context
	user := ctx.Locals("user")
	details, ok := user.(entity.User)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Unauthorized",
			"details": "Invalid user context",
		})
	}

	// Memanggil service untuk mendapatkan data rekap
	data, err := tr.customerMtrService.ListBerminatMembership(details.Username, startDate, endDate)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Failed to fetch data",
			"details": err.Error(),
		})
	}

	// Response sukses dengan format hanya "YYYY-MM-DD"
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil",
		"user":    details.Username,
		"data":    data,
		"Dari":    startDate.Format(layoutDate),
		"Sampai":  endDate.Format(layoutDate),
	})
}

func (tr *customerMtrController) ListDataAsuransiPA(ctx *fiber.Ctx) error {
	var rekapReq request.RangeTanggalRequest

	// Parsing request body ke struct
	if err := ctx.BodyParser(&rekapReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
	}

	// Format tanggal yang digunakan
	layoutFull := "2006-01-02 15:04:05"
	layoutDate := "2006-01-02"

	// Load lokasi waktu Jakarta
	loc, _ := time.LoadLocation("Asia/Jakarta")

	// Default StartDate dan EndDate (tanggal hari ini dalam zona waktu Jakarta)
	now := time.Now().In(loc)
	startDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	endDate := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, loc)

	// Parsing Tgl1 jika ada
	if rekapReq.Tgl1 != "" {
		parsedStart, err := time.ParseInLocation(layoutFull, rekapReq.Tgl1, loc)
		if err != nil {
			parsedStart, err = time.ParseInLocation(layoutDate, rekapReq.Tgl1, loc)
			if err == nil {
				parsedStart = time.Date(parsedStart.Year(), parsedStart.Month(), parsedStart.Day(), 0, 0, 0, 0, loc)
			}
		}
		if err == nil {
			startDate = parsedStart
		}
	}

	// Parsing Tgl2 jika ada
	if rekapReq.Tgl2 != "" {
		parsedEnd, err := time.ParseInLocation(layoutFull, rekapReq.Tgl2, loc)
		if err != nil {
			parsedEnd, err = time.ParseInLocation(layoutDate, rekapReq.Tgl2, loc)
			if err == nil {
				parsedEnd = time.Date(parsedEnd.Year(), parsedEnd.Month(), parsedEnd.Day(), 23, 59, 59, 0, loc)
			}
		}
		if err == nil {
			endDate = parsedEnd
		}
	}

	// Mengambil informasi user dari context
	user := ctx.Locals("user")
	details, ok := user.(entity.User)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Unauthorized",
			"details": "Invalid user context",
		})
	}

	// Memanggil service untuk mendapatkan data rekap
	data, err := tr.customerMtrService.ListDataAsuransiPA(details.Username, startDate, endDate)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Failed to fetch data",
			"details": err.Error(),
		})
	}

	// Response sukses dengan format hanya "YYYY-MM-DD"
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil",
		"user":    details.Username,
		"data":    data,
		"Dari":    startDate.Format(layoutDate),
		"Sampai":  endDate.Format(layoutDate),
	})
}

func (tr *customerMtrController) ListDataAsuransiMtr(ctx *fiber.Ctx) error {
	var rekapReq request.RangeTanggalRequest

	// Parsing request body ke struct
	if err := ctx.BodyParser(&rekapReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
	}

	// Format tanggal yang digunakan
	layoutFull := "2006-01-02 15:04:05"
	layoutDate := "2006-01-02"

	// Load lokasi waktu Jakarta
	loc, _ := time.LoadLocation("Asia/Jakarta")

	// Default StartDate dan EndDate (tanggal hari ini dalam zona waktu Jakarta)
	now := time.Now().In(loc)
	startDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	endDate := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, loc)

	// Parsing Tgl1 jika ada
	if rekapReq.Tgl1 != "" {
		parsedStart, err := time.ParseInLocation(layoutFull, rekapReq.Tgl1, loc)
		if err != nil {
			parsedStart, err = time.ParseInLocation(layoutDate, rekapReq.Tgl1, loc)
			if err == nil {
				parsedStart = time.Date(parsedStart.Year(), parsedStart.Month(), parsedStart.Day(), 0, 0, 0, 0, loc)
			}
		}
		if err == nil {
			startDate = parsedStart
		}
	}

	// Parsing Tgl2 jika ada
	if rekapReq.Tgl2 != "" {
		parsedEnd, err := time.ParseInLocation(layoutFull, rekapReq.Tgl2, loc)
		if err != nil {
			parsedEnd, err = time.ParseInLocation(layoutDate, rekapReq.Tgl2, loc)
			if err == nil {
				parsedEnd = time.Date(parsedEnd.Year(), parsedEnd.Month(), parsedEnd.Day(), 23, 59, 59, 0, loc)
			}
		}
		if err == nil {
			endDate = parsedEnd
		}
	}

	// Mengambil informasi user dari context
	user := ctx.Locals("user")
	details, ok := user.(entity.User)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Unauthorized",
			"details": "Invalid user context",
		})
	}

	// Memanggil service untuk mendapatkan data rekap
	data, err := tr.customerMtrService.ListDataAsuransiMtr(details.Username, startDate, endDate)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Failed to fetch data",
			"details": err.Error(),
		})
	}

	// Response sukses dengan format hanya "YYYY-MM-DD"
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil",
		"user":    details.Username,
		"data":    data,
		"Dari":    startDate.Format(layoutDate),
		"Sampai":  endDate.Format(layoutDate),
	})
}