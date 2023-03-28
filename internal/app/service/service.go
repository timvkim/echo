package service

import (
	"time"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) DaysLeft() int64 {
	// create a date object and put it in new var 'd'
	d := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)

	// use `Until` method to get a duration
	dur := time.Until(d)

	return int64(dur.Hours()) / 24

}
