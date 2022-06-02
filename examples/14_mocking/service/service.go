package service

import "github.com/kjkondratuk/how-to-go/examples/14_mocking/slow_service"

type service struct {
	service slow_service.SlowService
}

type Service interface {
	ApplyPower(list []int) []int
}

func New(slowService slow_service.SlowService) Service {
	return &service{
		service: slowService,
	}
}

func (s *service) ApplyPower(list []int) []int {
	newList := make([]int, len(list))
	for i, item := range list {
		newList[i] = item ^ s.service.GetPower()
	}
	return newList
}
