package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// There are frameworks for mock generation like go-mock that make it so you don't
// actually have to write these mocks--I just like the control this gives me, and it
// helps exhibit the point with less "magic" involved
type MockSlowService struct {
	mock.Mock
}

func (m *MockSlowService) GetPower() int {
	args := m.Called()
	return args.Int(0)
}

func TestService_ApplyPower(t *testing.T) {
	t.Run("should apply power to all items quickly because we're not using the actual function", func(tc *testing.T) {
		mockSlowService := MockSlowService{}
		mockSlowService.On("GetPower").Return(4)
		s := New(&mockSlowService)

		list := make([]int, 50)
		for i := 0; i < 50; i++ {
			list[i] = i
		}

		result := s.ApplyPower(list)

		for idx, val := range result {
			assert.Equal(tc, list[idx]^4, val)
		}
	})
}
