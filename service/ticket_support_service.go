package service

import (
	"fmt"
	"wkm/entity"
	"wkm/repository"
	"wkm/request"
	"wkm/response"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type TicketSupportService interface {
	CreateTicketSupport(data request.TicketRequest, username string, tier uint32) (string, string, error)
	EditTicketSupport(noTicket string, data request.TicketRequest, username string, role uint32) (string, string, error)
	ViewTicketSupport(noTicket string) (entity.TicketSupport, error)
	ListTicketUser(username string) ([]entity.TicketSupport, error)
	ListTicketIT(username string) ([]entity.TicketSupport, error)
	ListTicketQueue(month string, year string) ([]entity.TicketSupport, error)
	ListItSupport() ([]response.ItSupports, error)
	ExportDataTicketSupport(month int, year int) (string, error)
}

type ticketSupportService struct {
	tsR repository.TicketSupportRepository
}

func NewTicketSupportService(tR repository.TicketSupportRepository) TicketSupportService {
	return &ticketSupportService{
		tsR: tR,
	}
}

// CreateTicketSupport handles the business logic for creating a ticket.
func (s *ticketSupportService) CreateTicketSupport(data request.TicketRequest, username string, tier uint32) (string, string, error) {
	// Call repository to create ticket and get noTicket and error
	noTicket, assignResult, err := s.tsR.CreateTicketSupport(data, username, tier)
	if err != nil {
		// Return error if exists
		return "", "", err
	}

	// Return noTicket if successful
	return noTicket, assignResult, nil
}

func (s *ticketSupportService) EditTicketSupport(noTicket string, data request.TicketRequest, username string, role uint32) (string, string, error) {
	// Call repository to edit ticket and get error
	noTicket, err := s.tsR.EditTicketSupport(noTicket, data, username, role)
	if err != nil {
		// Return error if exists
		return "", "", err
	}

	// Return nil if successful
	return noTicket, "Ticket updated successfully", nil
}

func (s *ticketSupportService) ViewTicketSupport(noTicket string) (entity.TicketSupport, error) {
	ticket, err := s.tsR.ViewTicketSupport(noTicket)
	if err != nil {
		return entity.TicketSupport{}, err
	}
	return ticket, nil
}

func (s *ticketSupportService) ListItSupport() ([]response.ItSupports, error) {
	data, err := s.tsR.ListItSupport()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *ticketSupportService) ListTicketUser(username string) ([]entity.TicketSupport, error) {
	ticket, err := s.tsR.ListTicketUser(username)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}

func (s *ticketSupportService) ListTicketQueue(month string, year string) ([]entity.TicketSupport, error) {
	ticket, err := s.tsR.ListTicketQueue(month, year)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}

func (s *ticketSupportService) ListTicketIT(username string) ([]entity.TicketSupport, error) {
	ticket, err := s.tsR.ListTicketIT(username)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}

func (s *ticketSupportService) ExportDataTicketSupport(month int, year int) (string, error) {
	// Ambil data dari repository berdasarkan bulan dan tahun
	tickets, err := s.tsR.ExportDataTicketSupport(month, year)
	if err != nil {
		return "", fmt.Errorf("failed to fetch ticket data: %w", err)
	}
	ticketsSheet2, err := s.tsR.ExportDataTicketSupportSheet2(month, year)
	if err != nil {
		return "", fmt.Errorf("failed to fetch ticket data: %w", err)
	}

	// Map tier ke string
	tierMapping := map[int]string{
		1: "Platinum",
		2: "Gold",
	}

	// Membuat file Excel baru
	xlsx := excelize.NewFile()
	sheet1Name := "Rekap Tiket Support"
	xlsx.NewSheet(sheet1Name) // Create a new sheet
	xlsx.DeleteSheet("Sheet1")
	// Membuat style untuk header dengan background warna kuning
	headerStyle, err := xlsx.NewStyle(`{
        "fill": {"type": "pattern", "color": ["#FFFF00"], "pattern": 1},
        "border": [{"type": "left", "color": "#000000", "style": 1},
                   {"type": "right", "color": "#000000", "style": 1},
                   {"type": "top", "color": "#000000", "style": 1},
                   {"type": "bottom", "color": "#000000", "style": 1}]
    }`)
	if err != nil {
		return "", fmt.Errorf("failed to create header style: %w", err)
	}

	// Membuat style untuk tabel dengan border
	tableBorderStyle, err := xlsx.NewStyle(`{
        "border": [{"type": "left", "color": "#000000", "style": 1},
                   {"type": "right", "color": "#000000", "style": 1},
                   {"type": "top", "color": "#000000", "style": 1},
                   {"type": "bottom", "color": "#000000", "style": 1}]
    }`)
	if err != nil {
		return "", fmt.Errorf("failed to create table border style: %w", err)
	}

	// Menambahkan header
	header1 := []string{"Tier", "Plan", "Actual Plan"}
	for i, col := range header1 {
		cell := fmt.Sprintf("%s1", string('A'+i)) // Setting header row in cells A1, B1, C1
		xlsx.SetCellValue(sheet1Name, cell, col)
		xlsx.SetCellStyle(sheet1Name, cell, cell, headerStyle) // Apply header style (yellow background)
	}

	// Menambahkan data
	rowIndex := 2 // Start at row 2, because row 1 is for header
	for _, ticket := range tickets {
		// Menambahkan data ke Excel
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", rowIndex), tierMapping[ticket.TierTicket]) // Tier
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", rowIndex), ticket.Plan)                    // Plan
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", rowIndex), ticket.ActualPlan)              // Actual Plan

		// Apply border style for each cell in the row
		xlsx.SetCellStyle(sheet1Name, fmt.Sprintf("A%d", rowIndex), fmt.Sprintf("A%d", rowIndex), tableBorderStyle)
		xlsx.SetCellStyle(sheet1Name, fmt.Sprintf("B%d", rowIndex), fmt.Sprintf("B%d", rowIndex), tableBorderStyle)
		xlsx.SetCellStyle(sheet1Name, fmt.Sprintf("C%d", rowIndex), fmt.Sprintf("C%d", rowIndex), tableBorderStyle)

		// Increment row index for next row
		rowIndex++
	}

	sheet2Name := "Tiket Per IT"
	xlsx.NewSheet(sheet2Name)

	header2 := []string{"IT Support", "Tier", "Plan", "Actual Plan"}
	for i, col := range header2 {
		cell := fmt.Sprintf("%s1", string('A'+i)) // Setting header row in cells A1, B1, C1
		xlsx.SetCellValue(sheet2Name, cell, col)
		xlsx.SetCellStyle(sheet2Name, cell, cell, headerStyle) // Apply header style (yellow background)
	}

	rowIndex1 := 2 // Start at row 2, because row 1 is for header
	for _, ticket := range ticketsSheet2 {
		// Menambahkan data ke sheet2
		xlsx.SetCellValue(sheet2Name, fmt.Sprintf("A%d", rowIndex1), ticket.Name)                    // IT Support
		xlsx.SetCellValue(sheet2Name, fmt.Sprintf("B%d", rowIndex1), tierMapping[ticket.TierTicket]) // Tier
		xlsx.SetCellValue(sheet2Name, fmt.Sprintf("C%d", rowIndex1), ticket.Plan)                    // Plan
		xlsx.SetCellValue(sheet2Name, fmt.Sprintf("D%d", rowIndex1), ticket.ActualPlan)              // Actual Plan

		// Terapkan style border untuk setiap sel di baris
		xlsx.SetCellStyle(sheet2Name, fmt.Sprintf("A%d", rowIndex1), fmt.Sprintf("A%d", rowIndex1), tableBorderStyle)
		xlsx.SetCellStyle(sheet2Name, fmt.Sprintf("B%d", rowIndex1), fmt.Sprintf("B%d", rowIndex1), tableBorderStyle)
		xlsx.SetCellStyle(sheet2Name, fmt.Sprintf("C%d", rowIndex1), fmt.Sprintf("C%d", rowIndex1), tableBorderStyle)
		xlsx.SetCellStyle(sheet2Name, fmt.Sprintf("D%d", rowIndex1), fmt.Sprintf("D%d", rowIndex1), tableBorderStyle)

		// Increment row index for next row
		rowIndex1++
	}

	// Menyimpan file Excel
	fileName := fmt.Sprintf("Data_Tiket_Support_%d_%d.xlsx", month, year)
	if err := xlsx.SaveAs(fileName); err != nil {
		return "", fmt.Errorf("failed to save excel file: %w", err)
	}

	return fileName, nil
}
