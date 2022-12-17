package gorm

import (
	"context"
	"time"

	"github.com/storyofhis/xtrame/httpserver/repository"
	"github.com/storyofhis/xtrame/httpserver/repository/models"
	"gorm.io/gorm"
)

type ticketRepo struct {
	db *gorm.DB
}

func NewTicketRepo(db *gorm.DB) repository.TicketRepo {
	return &ticketRepo{
		db: db,
	}
}

func (repo *ticketRepo) CreateTicket(ctx context.Context, ticket *models.Tickets) error {
	ticket.CreatedAt = time.Now()
	return repo.db.WithContext(ctx).Create(ticket).Error
}

func (repo *ticketRepo) GetTickets(ctx context.Context) ([]models.Tickets, error) {
	var ticket []models.Tickets
	err := repo.db.WithContext(ctx).Find(ticket).Error
	return ticket, err
}

func (repo *ticketRepo) GetTicketById(ctx context.Context, id uint) (*models.Tickets, error) {
	var ticket models.Tickets
	if err := repo.db.WithContext(ctx).Where("id = ?", id).Take(ticket).Error; err != nil {
		return nil, err
	}
	return &ticket, nil
}

func (repo *ticketRepo) UpdateTicket(ctx context.Context, ticket *models.Tickets, id uint) error {
	ticket.UpdatedAt = time.Now()
	if err := repo.db.WithContext(ctx).Where("id = ?", id).Updates(ticket).Error; err != nil {
		return err
	}
	return nil
}

func (repo *ticketRepo) DeleteTicket(ctx context.Context, ticket *models.Tickets, id uint) error {
	ticket.DeletedAt = time.Now()
	if err := repo.db.WithContext(ctx).Where("id = ?", id).Delete(ticket).Error; err != nil {
		return err
	}
	return nil
}
