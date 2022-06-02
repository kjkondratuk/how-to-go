package slow_service

import "time"

type slowService struct{}

type SlowService interface {
	GetPower() int
}

func New() SlowService {
	return &slowService{}
}

// GetPower : returns a power of 3 after a long delay--this should make it obvious if
// this is ever getting called
func (s *slowService) GetPower() int {
	time.Sleep(2 * time.Second)
	return 3
}
