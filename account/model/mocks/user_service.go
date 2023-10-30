package mocks

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jj-attaq/api-todo/account/model"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
    mock.Mock
}

func (m *MockUserService) Get(ctx context.Context, uid uuid.UUID) (*model.User, error) {
    // args will be returned in tests when called with uid
    ret := m.Called(ctx, uid)

    // 1st value
    var r0 *model.User
    if ret.Get(0) != nil {
        r0 = ret.Get(0).(*model.User)
        log.Printf("err for r0: %v\n", r0)
    }

    // 2nd value
    var r1 error
    if ret.Get(1) != nil {
        log.Printf("err for r1: %v\n", r1)
        r1 = ret.Get(1).(error)
    }

    return r0, r1
}
