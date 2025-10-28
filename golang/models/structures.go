package models

import (
	"database/sql"
	"time"
)

// Flight представляет рейс с поддержкой nullable полей
type Flight struct {
	FlightID           int          `json:"flightId"`
	RouteNo            string       `json:"routeNo"`
	Status             string       `json:"status"`
	ScheduledDeparture time.Time    `json:"scheduledDeparture"`
	ScheduledArrival   time.Time    `json:"scheduledArrival"`
	ActualDeparture    sql.NullTime `json:"actualDeparture,omitempty"`
	ActualArrival      sql.NullTime `json:"actualArrival,omitempty"`
	Segments           []Segment    `json:"segments"`
}

// Segment представляет сегмент перелета с поддержкой nullable полей
type Segment struct {
	TicketNo       string          `json:"ticketNo"`
	FlightID       int             `json:"flightId"`
	FareConditions string          `json:"fareConditions"`
	Price          sql.NullFloat64 `json:"price"`
	Ticket         *Ticket         `json:"ticket,omitempty"`
	BoardingPasses []BoardingPass  `json:"boardingPasses"`
}

// Ticket представляет билет с поддержкой nullable полей
type Ticket struct {
	TicketNo      string       `json:"ticketNo"`
	BookRef       string       `json:"bookRef"`
	PassengerID   string       `json:"passengerId"`
	PassengerName string       `json:"passengerName"`
	Outbound      bool         `json:"outbound"`
	DepartureTime sql.NullTime `json:"departureTime,omitempty"`
	Booking       *Booking     `json:"booking,omitempty"`
}

// Booking представляет бронирование с поддержкой nullable полей
type Booking struct {
	BookRef     string         `json:"bookRef"`
	BookDate    time.Time      `json:"bookDate"`
	TotalAmount sql.NullString `json:"totalAmount,omitempty"`
}

// BoardingPass представляет посадочный талон с поддержкой nullable полей
type BoardingPass struct {
	TicketNo     string        `json:"ticketNo"`
	FlightID     int           `json:"flightId"`
	SeatNo       string        `json:"seatNo"`
	BoardingNo   sql.NullInt64 `json:"boardingNo,omitempty"`
	BoardingTime sql.NullTime  `json:"boardingTime,omitempty"`
}
