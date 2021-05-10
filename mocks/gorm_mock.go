package mocks

import (
	"github.com/el-Mike/gochat/persist"
	"github.com/stretchr/testify/mock"
)

// GormMock - basic, reusable mock for Gorm client.
type GormMock struct {
	mock.Mock
}

func NewGormMock() *GormMock {
	return &GormMock{}
}

// First - First method mock implementation.
func (gm *GormMock) First(dest interface{}, conds ...interface{}) *persist.DBResponse {
	args := gm.Called(dest, conds)

	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(*persist.DBResponse)
}

// Where - Where method mock implementation.
func (gm *GormMock) Where(query interface{}, queryArgs ...interface{}) *persist.DBResponse {
	args := gm.Called(query, queryArgs)

	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(*persist.DBResponse)
}

// FirstWhere - FirstWhere mock implementation.
func (gm *GormMock) FirstWhere(dest interface{}, query interface{}, queryArgs ...interface{}) *persist.DBResponse {
	args := gm.Called(dest, query, queryArgs)

	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(*persist.DBResponse)
}

// First - Find method mock implementation.
func (gm *GormMock) Find(dest interface{}, conds ...interface{}) *persist.DBResponse {
	args := gm.Called(dest, conds)

	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(*persist.DBResponse)
}

// First - Save method mock implementation.
func (gm *GormMock) Save(value interface{}) *persist.DBResponse {
	args := gm.Called(value)

	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(*persist.DBResponse)
}

func (gm *GormMock) DeleteByID(target interface{}, id interface{}) *persist.DBResponse {
	args := gm.Called(target, id)

	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(*persist.DBResponse)
}

func GetDefaultDBResponse() *persist.DBResponse {
	return persist.NewDBResponse()
}

func GetErrorDBResponse(err error) *persist.DBResponse {
	res := persist.NewDBResponse()
	res.SetErr(err)

	return res
}
