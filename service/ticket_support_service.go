package service

import (
	"wkm/entity"
	"wkm/repository"
	"wkm/request"
)

type TicketSupportService interface {
	CreateTicketSupport(data request.TicketRequest, username string, tier uint32) (string, string, error)
	EditTicketSupport(noTicket string, data request.TicketRequest, username string, role uint32) (string, string, error)
	ViewTicketSupport(noTicket string) (entity.TicketSupport, error)
	ListTicketUser(username string) ([]entity.TicketSupport, error)
	ListTicketIT(username string) ([]entity.TicketSupport, error)
	ListTicketQueue(month string, year string) ([]entity.TicketSupport, error)
	// ExportDataTicketSupport(month string, year string) ([]entity.TicketSupport, error)
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

// func (s *ticketSupportService) ExportDataTicketSupport(month string, year string) ([]entity.TicketSupport, error) {
// 	ticket, err := s.tsR.ExportDataTicketSupport(month, year)
// 	if err != nil {
// 		return nil,
	
