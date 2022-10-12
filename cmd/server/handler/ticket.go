package handler

import (
	"net/http"

	"desafio-goweb-juanhernandez/internal/tickets"

	"github.com/gin-gonic/gin"
)

type Service struct {
	service tickets.Service
}

func NewService(s tickets.Service) *Service {
	return &Service{
		service: s,
	}
}

func (s *Service) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		ts, err := s.service.GetAll(c)

		if err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, ts)
	}
}

func (s *Service) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		tickets, err := s.service.GetTotalTickets(c, destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, tickets)
	}
}

func (s *Service) AverageDestination() gin.HandlerFunc {
	return func(c *gin.Context) {

		// destination := c.Param("dest")

		// avg, err := s.service.AverageDestination(c, destination)
		// if err != nil {
		// 	c.String(http.StatusInternalServerError, err.Error(), nil)
		// 	return
		// }

		// c.JSON(200, avg)
	}
}
