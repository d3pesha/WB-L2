package service

import (
	"sync"
	"time"
)

type Event struct {
	Date time.Time
	Mes  string
}

type Calendar struct {
	m     *sync.Mutex
	store map[int]Event
}

type EventService interface {
	CreateEvent(data time.Time, msg string) int
	UpdateEvent(id int, data time.Time, msg string) (Event, error)
	DeleteEvent(id int) error
	ParseEvent(date time.Time, days int) ([]Event, error)
}
