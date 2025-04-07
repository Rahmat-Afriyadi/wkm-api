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
	SelfCount(ctx *fiber.Ctx) error
	MasterData(ctx *fiber.Ctx) error
	MasterDataCount(ctx *fiber.Ctx) error
	MasterDataBalikan(ctx *fiber.Ctx) error
	MasterDataBalikanCount(ctx *fiber.Ctx) error
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
	RekapLeaderTs(ctx *fiber.Ctx) error
	RekapBerminatPerWilayah(ctx *fiber.Ctx) error
	ExportRekapLeaderTs(ctx *fiber.Ctx) error
	ListPerformanceTs(ctx *fiber.Ctx) error
	RekapStatus(ctx *fiber.Ctx) error
}

type customerMtrController struct {
	customerMtrService service.CustomerMtrService
}

func NewCustomerMtrController(aS service.CustomerMtrService) CustomerMtrController {
	return &customerMtrController{
		customerMtrService: aS,
	}
}

func (tr *customerMtrController) SelfCount(ctx *fiber.Ctx) error {
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	data := tr.customerMtrService.SelfCount(details.Username)
	return ctx.Status(200).JSON(data)
}
func (tr *customerMtrController) MasterDataBalikan(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	pageParams, _ := strconv.Atoi(ctx.Query("pageParams"))
	data := tr.customerMtrService.MasterDataBalikan(search, details.Username, limit, pageParams)
	return ctx.Status(200).JSON(data)
}
func (tr *customerMtrController) MasterDataBalikanCount(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	data := tr.customerMtrService.MasterDataBalikanCount(search, details.Username)
	return ctx.Status(200).JSON(data)
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
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	if data.KdUserTs != details.Username {
		return ctx.Status(400).JSON(fiber.Map{"message": "Data tidak ditemukan"})
	}
	return ctx.Status(200).JSON(fiber.Map{"message": "Berhasil ", "data":data})
}

func (tr *customerMtrController) ShowBalikan(ctx *fiber.Ctx) error {
	noMsn := ctx.Params("no_msn")
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	if condition := tr.customerMtrService.AmbilDataBalikan(noMsn, details.Username); condition != nil {
		return ctx.Status(400).JSON(fiber.Map{"message": "Data tidak ditemukan"})
	}
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
    // Ambil nilai tgl1 dan tgl2 dari query params
    tgl1 := ctx.Query("tgl1")
    tgl2 := ctx.Query("tgl2")

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
    if tgl1 != "" {
        parsedStart, err := time.ParseInLocation(layoutFull, tgl1, loc)
        if err != nil {
            parsedStart, err = time.ParseInLocation(layoutDate, tgl1, loc)
            if err == nil {
                parsedStart = time.Date(parsedStart.Year(), parsedStart.Month(), parsedStart.Day(), 0, 0, 0, 0, loc)
            }
        }
        if err == nil {
            startDate = parsedStart
        }
    }

    // Parsing Tgl2 jika ada
    if tgl2 != "" {
        parsedEnd, err := time.ParseInLocation(layoutFull, tgl2, loc)
        if err != nil {
            parsedEnd, err = time.ParseInLocation(layoutDate, tgl2, loc)
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
	// Format tanggal yang digunakan
	layoutFull := "2006-01-02 15:04:05"
	layoutDate := "2006-01-02"

	// Load lokasi waktu Jakarta
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	startDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	endDate := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, loc)
	tgl1 := ctx.Query("tgl1", "") // Default kosong jika tidak ada
	tgl2 := ctx.Query("tgl2", "")
	search := ctx.Query("search", "")

	// Parsing tgl1 jika ada
	if tgl1 != "" {
		parsedStart, err := time.ParseInLocation(layoutFull, tgl1, loc)
		if err != nil {
			parsedStart, err = time.ParseInLocation(layoutDate, tgl1, loc)
			if err == nil {
				parsedStart = time.Date(parsedStart.Year(), parsedStart.Month(), parsedStart.Day(), 0, 0, 0, 0, loc)
			}
		}
		if err == nil {
			startDate = parsedStart
		}
	}

	// Parsing tgl2 jika ada
	if tgl2 != "" {
		parsedEnd, err := time.ParseInLocation(layoutFull, tgl2, loc)
		if err != nil {
			parsedEnd, err = time.ParseInLocation(layoutDate, tgl2, loc)
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
	limit, _ := strconv.Atoi(ctx.Query("limit", "10")) // Default limit 10 jika tidak ada
	pageParams, _ := strconv.Atoi(ctx.Query("pageParams", "1"))
	// Memanggil service untuk mendapatkan data rekap
	data, totalPages,totalRecords, err := tr.customerMtrService.ListBerminatMembership(details.Username, startDate, endDate, limit, pageParams,search)
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
		"pages": pageParams,
		"total_pages": totalPages,
		"total_records": totalRecords,
		"Dari":    startDate.Format(layoutDate),
		"Sampai":  endDate.Format(layoutDate),
	})
}

func (tr *customerMtrController) ListDataAsuransiPA(ctx *fiber.Ctx) error {
	// Format tanggal yang digunakan
	layoutFull := "2006-01-02 15:04:05"
	layoutDate := "2006-01-02"

	// Load lokasi waktu Jakarta
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	startDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	endDate := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, loc)
	tgl1 := ctx.Query("tgl1", "") // Default kosong jika tidak ada
	tgl2 := ctx.Query("tgl2", "")

	// Parsing tgl1 jika ada
	if tgl1 != "" {
		parsedStart, err := time.ParseInLocation(layoutFull, tgl1, loc)
		if err != nil {
			parsedStart, err = time.ParseInLocation(layoutDate, tgl1, loc)
			if err == nil {
				parsedStart = time.Date(parsedStart.Year(), parsedStart.Month(), parsedStart.Day(), 0, 0, 0, 0, loc)
			}
		}
		if err == nil {
			startDate = parsedStart
		}
	}

	// Parsing tgl2 jika ada
	if tgl2 != "" {
		parsedEnd, err := time.ParseInLocation(layoutFull, tgl2, loc)
		if err != nil {
			parsedEnd, err = time.ParseInLocation(layoutDate, tgl2, loc)
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
	limit, _ := strconv.Atoi(ctx.Query("limit", "10")) // Default limit 10 jika tidak ada
	pageParams, _ := strconv.Atoi(ctx.Query("pageParams", "1"))
	search := ctx.Query("search", "")
	// Memanggil service untuk mendapatkan data rekap
	data, totalPages,totalRecords, err := tr.customerMtrService.ListDataAsuransiPA(details.Username, startDate, endDate, limit, pageParams, search)
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
		"pages": pageParams,
		"total_pages": totalPages,
		"total_records": totalRecords,
		"Dari":    startDate.Format(layoutDate),
		"Sampai":  endDate.Format(layoutDate),
	})
}

func (tr *customerMtrController) ListDataAsuransiMtr(ctx *fiber.Ctx) error {
	// Format tanggal yang digunakan
	layoutFull := "2006-01-02 15:04:05"
	layoutDate := "2006-01-02"

	// Load lokasi waktu Jakarta
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	startDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	endDate := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, loc)
	tgl1 := ctx.Query("tgl1", "") // Default kosong jika tidak ada
	tgl2 := ctx.Query("tgl2", "")

	// Parsing tgl1 jika ada
	if tgl1 != "" {
		parsedStart, err := time.ParseInLocation(layoutFull, tgl1, loc)
		if err != nil {
			parsedStart, err = time.ParseInLocation(layoutDate, tgl1, loc)
			if err == nil {
				parsedStart = time.Date(parsedStart.Year(), parsedStart.Month(), parsedStart.Day(), 0, 0, 0, 0, loc)
			}
		}
		if err == nil {
			startDate = parsedStart
		}
	}

	// Parsing tgl2 jika ada
	if tgl2 != "" {
		parsedEnd, err := time.ParseInLocation(layoutFull, tgl2, loc)
		if err != nil {
			parsedEnd, err = time.ParseInLocation(layoutDate, tgl2, loc)
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
	limit, _ := strconv.Atoi(ctx.Query("limit", "10")) // Default limit 10 jika tidak ada
	pageParams, _ := strconv.Atoi(ctx.Query("pageParams", "1"))
	search := ctx.Query("search", "")
	// Memanggil service untuk mendapatkan data rekap
	data, totalPages,totalRecords, err := tr.customerMtrService.ListDataAsuransiMtr(details.Username, startDate, endDate, limit, pageParams, search)
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
		"pages": pageParams,
		"total_pages": totalPages,
		"total_records": totalRecords,
		"Dari":    startDate.Format(layoutDate),
		"Sampai":  endDate.Format(layoutDate),
	})
}
func (tr *customerMtrController) RekapLeaderTs(ctx *fiber.Ctx) error {
    // Ambil nilai tgl1 dan tgl2 dari query params
    tgl1 := ctx.Query("tgl1")
    tgl2 := ctx.Query("tgl2")

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
    if tgl1 != "" {
        parsedStart, err := time.ParseInLocation(layoutFull, tgl1, loc)
        if err != nil {
            parsedStart, err = time.ParseInLocation(layoutDate, tgl1, loc)
            if err == nil {
                parsedStart = time.Date(parsedStart.Year(), parsedStart.Month(), parsedStart.Day(), 0, 0, 0, 0, loc)
            }
        }
        if err == nil {
            startDate = parsedStart
        }
    }

    // Parsing Tgl2 jika ada
    if tgl2 != "" {
        parsedEnd, err := time.ParseInLocation(layoutFull, tgl2, loc)
        if err != nil {
            parsedEnd, err = time.ParseInLocation(layoutDate, tgl2, loc)
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
	if details.RoleId != 2 {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error":   "Unauthorized",
		"details": "Invlaid Role",
	})}
    // Memanggil service untuk mendapatkan data rekap
    rekapData, err := tr.customerMtrService.RekapLeaderTs(startDate, endDate)
    if err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error":   "Failed to fetch data",
            "details": err.Error(),
        })
    }

    // Response sukses dengan format hanya "YYYY-MM-DD"
    return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Berhasil",
        "data":    rekapData,
        "Dari":    startDate.Format(layoutDate),
        "Sampai":  endDate.Format(layoutDate),
    })
}
func (tr *customerMtrController) RekapBerminatPerWilayah(ctx *fiber.Ctx) error {
    // Ambil nilai tgl1 dan tgl2 dari query params
    tgl1 := ctx.Query("tgl1")
    tgl2 := ctx.Query("tgl2")

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
    if tgl1 != "" {
        parsedStart, err := time.ParseInLocation(layoutFull, tgl1, loc)
        if err != nil {
            parsedStart, err = time.ParseInLocation(layoutDate, tgl1, loc)
            if err == nil {
                parsedStart = time.Date(parsedStart.Year(), parsedStart.Month(), parsedStart.Day(), 0, 0, 0, 0, loc)
            }
        }
        if err == nil {
            startDate = parsedStart
        }
    }

    // Parsing Tgl2 jika ada
    if tgl2 != "" {
        parsedEnd, err := time.ParseInLocation(layoutFull, tgl2, loc)
        if err != nil {
            parsedEnd, err = time.ParseInLocation(layoutDate, tgl2, loc)
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
	if details.RoleId != 2 {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error":   "Unauthorized",
		"details": "Invlaid Role",
	})}
    // Memanggil service untuk mendapatkan data rekap
    rekapData,totalData, err := tr.customerMtrService.RekapBerminatPerWilayah(startDate, endDate)
    if err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error":   "Failed to fetch data",
            "details": err.Error(),
        })
    }

    // Response sukses dengan format hanya "YYYY-MM-DD"
    return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Berhasil",
        "data":    rekapData,
		"total":totalData,
        "Dari":    startDate.Format(layoutDate),
        "Sampai":  endDate.Format(layoutDate),
    })
}
func (tr *customerMtrController) ExportRekapLeaderTs(ctx *fiber.Ctx) error {
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
	if details.RoleId != 2 {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error":   "Unauthorized",
		"details": "Invlaid Role",
	})}
	// Memanggil service untuk mendapatkan file Excel
	fileName, err := tr.customerMtrService.ExportRekapLeaderTs(startDate, endDate)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to generate report",
			"details": err.Error(),
		})
	}

	return ctx.Download(fileName)
}

func (tr *customerMtrController) ListPerformanceTs(ctx *fiber.Ctx) error {
    // Ambil nilai tgl1 dan tgl2 dari query params
    tgl1 := ctx.Query("tgl1")
    tgl2 := ctx.Query("tgl2")

    // Format tanggal yang digunakan
    layoutDate := "2006-01-02"
    
    // Load lokasi waktu Jakarta
    loc, _ := time.LoadLocation("Asia/Jakarta")

    // Default StartDate dan EndDate (tanggal hari ini dalam zona waktu Jakarta)
    now := time.Now().In(loc)
    startDate := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, loc)
    endDate := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, loc)

    // Parsing tanggal jika diberikan dalam query params
    if tgl1 != "" {
        parsedStart, err := time.ParseInLocation(layoutDate, tgl1, loc)
        if err != nil {
            return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error":   "Invalid start date format",
                "details": "Format harus YYYY-MM-DD",
            })
        }
        startDate = parsedStart
    }

    if tgl2 != "" {
        parsedEnd, err := time.ParseInLocation(layoutDate, tgl2, loc)
        if err != nil {
            return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error":   "Invalid end date format",
                "details": "Format harus YYYY-MM-DD",
            })
        }
        endDate = time.Date(parsedEnd.Year(), parsedEnd.Month(), parsedEnd.Day(), 23, 59, 59, 0, loc)
    }

    // Memastikan startDate tidak lebih besar dari endDate
    if startDate.After(endDate) {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error":   "Invalid date range",
            "details": "tgl1 tidak boleh lebih besar dari tgl2",
        })
    }

    // Mengambil informasi user dari context
    user := ctx.Locals("user")
    details, ok := user.(entity.User)
    if !ok || details.RoleId != 2 {
        return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error":   "Unauthorized",
            "details": "User tidak memiliki akses",
        })
    }

    // Memanggil service untuk mendapatkan data
    rekapData, err := tr.customerMtrService.ListPerformanceTs(startDate, endDate)
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error":   "Failed to fetch data",
            "details": err.Error(),
        })
    }

    // Response sukses dengan format hanya "YYYY-MM-DD"
    return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Berhasil",
        "data":    rekapData,
        "Dari":    startDate.Format(layoutDate),
        "Sampai":  endDate.Format(layoutDate),
    })
}
func (tr *customerMtrController) RekapStatus(ctx *fiber.Ctx) error {
    // Ambil nilai tgl1 dan tgl2 dari query params
    tgl1 := ctx.Query("tgl1")
    tgl2 := ctx.Query("tgl2")

    // Format tanggal yang digunakan
    layoutDate := "2006-01-02"
    
    // Load lokasi waktu Jakarta
    loc, _ := time.LoadLocation("Asia/Jakarta")

    // Default StartDate dan EndDate (tanggal hari ini dalam zona waktu Jakarta)
    now := time.Now().In(loc)
    startDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
    endDate := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, loc)

    // Parsing tanggal jika diberikan dalam query params
    if tgl1 != "" {
        parsedStart, err := time.ParseInLocation(layoutDate, tgl1, loc)
        if err != nil {
            return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error":   "Invalid start date format",
                "details": "Format harus YYYY-MM-DD",
            })
        }
        startDate = parsedStart
    }

    if tgl2 != "" {
        parsedEnd, err := time.ParseInLocation(layoutDate, tgl2, loc)
        if err != nil {
            return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error":   "Invalid end date format",
                "details": "Format harus YYYY-MM-DD",
            })
        }
        endDate = time.Date(parsedEnd.Year(), parsedEnd.Month(), parsedEnd.Day(), 23, 59, 59, 0, loc)
    }

    // Memastikan startDate tidak lebih besar dari endDate
    if startDate.After(endDate) {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error":   "Invalid date range",
            "details": "tgl1 tidak boleh lebih besar dari tgl2",
        })
    }

    // Mengambil informasi user dari context
    user := ctx.Locals("user")
    details, ok := user.(entity.User)
    if !ok || details.RoleId != 2 {
        return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error":   "Unauthorized",
            "details": "User tidak memiliki akses",
        })
    }

    // Memanggil service untuk mendapatkan data
    rekapData, err := tr.customerMtrService.GetRekapStatus(startDate, endDate)
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error":   "Failed to fetch data",
            "details": err.Error(),
        })
    }

    // Response sukses dengan format hanya "YYYY-MM-DD"
    return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Berhasil",
        "data":    rekapData,
        "Dari":    startDate.Format(layoutDate),
        "Sampai":  endDate.Format(layoutDate),
    })
}

