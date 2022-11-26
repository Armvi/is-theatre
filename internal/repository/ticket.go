package repository

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"is-theatre/internal/datastruct"
)

type TicketQuery interface {
	CreateTicket(ticket datastruct.Ticket) (*int64, error)
	Ticket(id int64) (*datastruct.Ticket, error)
	UpdateTicket(ticket datastruct.Ticket) (*datastruct.Ticket, error)
	DeleteTicket(id int64) error
	TicketsByUser(id int64) ([]datastruct.Ticket, error)
}

type ticketQuery struct{}

func (p *ticketQuery) CreateTicket(ticket datastruct.Ticket) (*int64, error) {
	qb := dbQueryBuilder().
		Insert(datastruct.TicketTableName).
		Columns("performanceId",
			"place",
			"ticketCost",
			"userId").
		Values(ticket.PerformanceSetId,
			ticket.PlaceNumber,
			ticket.Cost,
			ticket.UserId).
		Suffix("RETURNING id")

	var id int64
	err := qb.QueryRow().Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("create ticket error: %w", err)
	}

	return &id, nil
}

func (p *ticketQuery) Ticket(id int64) (*datastruct.Ticket, error) {
	db := dbQueryBuilder().
		Select("performanceId",
			"place",
			"ticketCost",
			"userId",
			"id").
		From(datastruct.TicketTableName).
		Where(squirrel.Eq{"id": id})

	ticket := datastruct.Ticket{}
	err := db.QueryRow().Scan(&ticket.PerformanceSetId,
		&ticket.PlaceNumber,
		&ticket.Cost,
		&ticket.UserId,
		&ticket.Id)
	if err != nil {
		return nil, fmt.Errorf("get ticket error: %w", err)
	}

	return &ticket, nil
}

func (p *ticketQuery) UpdateTicket(ticket datastruct.Ticket) (*datastruct.Ticket, error) {
	err := p.DeleteTicket(ticket.Id)
	if err != nil {
		return nil, fmt.Errorf("update ticket err: %w", err)
	}

	id, err := p.CreateTicket(ticket)

	if err != nil {
		return nil, fmt.Errorf("update ticket err: %w", err)
	}

	return p.Ticket(*id)
}

func (p *ticketQuery) DeleteTicket(id int64) error {
	db := dbQueryBuilder().
		Delete(datastruct.TicketTableName).
		Where(squirrel.Eq{"id": id})

	_, err := db.Exec()
	if err != nil {
		return fmt.Errorf("delete user error: %w", err)
	}
	return nil
}

func (p *ticketQuery) TicketsByUser(id int64) ([]datastruct.Ticket, error) {
	db := dbQueryBuilder().
		Select("performanceId",
			"place",
			"ticketCost",
			"userId",
			"id").
		From(datastruct.TicketTableName).
		Where(squirrel.Eq{"userId": id})

	var tickets []datastruct.Ticket
	var ticket datastruct.Ticket
	rows, err := db.Query()
	if err != nil {
		return nil, fmt.Errorf("get tickets error: %w", err)
	}
	for rows.Next() {
		err := db.QueryRow().Scan(&ticket.PerformanceSetId,
			&ticket.PlaceNumber,
			&ticket.Cost,
			&ticket.UserId,
			&ticket.Id)
		if err != nil {
			return nil, fmt.Errorf("get ticket: %w", err)
		}
		tickets = append(tickets, ticket)
	}

	return tickets, nil
}
