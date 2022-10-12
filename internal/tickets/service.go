package tickets

import (
	"context"
	"desafio-goweb-juanhernandez/internal/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
	GetTotalTickets(ctx context.Context, destination string) (int, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {
	return s.repository.GetTicketByDestination(ctx, destination)
}

func (s *service) GetAll(ctx context.Context) ([]domain.Ticket, error) {
	return s.repository.GetAll(ctx)
}

func (s *service) GetTotalTickets(ctx context.Context, destination string) (int, error) {
	ts, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0, err
	}

	return len(ts), nil
}
