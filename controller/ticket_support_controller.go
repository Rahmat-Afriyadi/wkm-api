package controller

import (
	// "fmt"
	// "sync"
	// "time"
	"wkm/entity"
	// "wkm/repository"
	"wkm/request"
	"wkm/service"

	// "github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gofiber/fiber/v2"
)

type TicketSupportController interface {
	CreateTicketSupport(ctx *fiber.Ctx) error
	EditTicketSupport(ctx *fiber.Ctx) error
	ViewTicketSupport(ctx *fiber.Ctx) error
	ListTicketUser(ctx *fiber.Ctx) error
	ListTicketQueue(ctx *fiber.Ctx) error
	ListTicketIT(ctx *fiber.Ctx) error
	ListItSupport(ctx *fiber.Ctx) error
	ExportDataTiketSupport(ctx *fiber.Ctx) error
}

type ticketSupportController struct {
	ticketSupportService service.TicketSupportService
}

func NewTicketSupportController(tS service.TicketSupportService) TicketSupportController {
	return &ticketSupportController{
		ticketSupportService: tS,
	}
}

func (tS *ticketSupportController) CreateTicketSupport(ctx *fiber.Ctx) error {
	// Dekode request JSON menjadi struct
	var ticketRequest request.TicketRequest
	if err := ctx.BodyParser(&ticketRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
	}

	// Ambil user details dari context
	user := ctx.Locals("user")
	details, ok := user.(entity.User) // Pastikan type assertion berhasil
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Unauthorized",
			"details": "Invalid user context",
		})
	}

	// Panggil service untuk membuat ticket
	noTicket,assignResult, err := tS.ticketSupportService.CreateTicketSupport(ticketRequest, details.Username, details.Tier)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to create ticket",
			"details": err.Error(),
		})
	}

	// Kembalikan response sukses dengan nomor tiket
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":   "Ticket created successfully",
		"no_ticket": noTicket,
		"assign_result": assignResult,
	})
}

func (tS *ticketSupportController) EditTicketSupport(ctx *fiber.Ctx) error {
    // Dekode request JSON menjadi struct
    var ticketRequest request.TicketRequest
    if err := ctx.BodyParser(&ticketRequest); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error":   "Invalid request body",
            "details": err.Error(),
        })
    }

    // Ambil user details dari context
    user := ctx.Locals("user")
    details, ok := user.(entity.User) // Pastikan type assertion berhasil
    if !ok {
        return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error":   "Unauthorized",
            "details": "Invalid user context",
        })
    }

    // Ambil nomor tiket dari parameter URL
    noTicket := ctx.Params("no_ticket")
    if noTicket == "" {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error":   "Invalid request",
            "details": "Nomor tiket tidak ditemukan",
        })
    }

    // Panggil service untuk edit ticket
    noTicket,message, err := tS.ticketSupportService.EditTicketSupport(noTicket, ticketRequest, details.Username, details.RoleId)
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error":   "Failed to edit ticket",
            "details": err.Error(),
        })
    }

    // Kembalikan response sukses
    return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": message,
		"no_ticket": noTicket,
    })
}

func (tS *ticketSupportController) ViewTicketSupport(ctx *fiber.Ctx) error {
	noTicket := ctx.Params("no_ticket")
	ticket, err := tS.ticketSupportService.ViewTicketSupport(noTicket)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(ticket)
}
func (tS *ticketSupportController) ListTicketUser(ctx *fiber.Ctx) error {
    // Mengambil parameter kd_user dari URL
    user := ctx.Locals("user")
	details, ok := user.(entity.User) // Pastikan type assertion berhasil
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Unauthorized",
			"details": "Invalid user context",
		})
	}
    
    // Memanggil service untuk mendapatkan daftar tiket berdasarkan kd_user
    tickets, err := tS.ticketSupportService.ListTicketUser(details.Username)
    if err != nil {
        // Menangani error jika ada
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    
    // Mengembalikan hasil daftar tiket dengan status OK
    return ctx.Status(fiber.StatusOK).JSON(tickets)
}

func (tS *ticketSupportController) ListTicketQueue(ctx *fiber.Ctx) error {
    // Memanggil service untuk mendapatkan daftar semua tiket
	month := ctx.Query("month", "") // Default kosong jika tidak disediakan
    year := ctx.Query("year", "")  

    tickets, err := tS.ticketSupportService.ListTicketQueue(month, year)
    if err != nil {
        // Menangani error jika ada
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    // Mengembalikan hasil daftar tiket dengan status OK
    return ctx.Status(fiber.StatusOK).JSON(tickets)
}

func (tS *ticketSupportController) ListTicketIT(ctx *fiber.Ctx) error {
    // Mengambil parameter kd_user dari URL
    user := ctx.Locals("user")
	details, ok := user.(entity.User) // Pastikan type assertion berhasil
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Unauthorized",
			"details": "Invalid user context",
		})
	}
    
    // Memanggil service untuk mendapatkan daftar tiket berdasarkan kd_user
    tickets, err := tS.ticketSupportService.ListTicketIT(details.Username)
    if err != nil {
        // Menangani error jika ada
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    
    // Mengembalikan hasil daftar tiket dengan status OK
    return ctx.Status(fiber.StatusOK).JSON(tickets)
}
func (tS *ticketSupportController) ListItSupport(ctx *fiber.Ctx) error {
	
	data, err := tS.ticketSupportService.ListItSupport()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(data)
}

func (tS *ticketSupportController) ExportDataTiketSupport(ctx *fiber.Ctx) error {
    var ticketRequest request.TicketRequest
    // Parsing the request body to extract month and year
    if err := ctx.BodyParser(&ticketRequest); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error":   "Invalid request body",
            "details": err.Error(),
        })
    }

    // Extract month and year from the request body
    month := ticketRequest.Month
    year := ticketRequest.Year

    // Memanggil service untuk mengekspor data tiket support
    fileName, err := tS.ticketSupportService.ExportDataTicketSupport(month, year)
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error":   "Failed to export data",
            "details": err.Error(),
        })
    }

    // Returning the generated file for download
    return ctx.Download(fileName)
}

