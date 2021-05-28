// Code generated by MockGen. DO NOT EDIT.
// Source: ./querier.go

// Package mock_gen is a generated GoMock package.
package mock_gen

import (
	context "context"
	sql "database/sql"
	reflect "reflect"
	gen "sqlc/gen"

	gomock "github.com/golang/mock/gomock"
)

// MockQuerier is a mock of Querier interface.
type MockQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockQuerierMockRecorder
}

// MockQuerierMockRecorder is the mock recorder for MockQuerier.
type MockQuerierMockRecorder struct {
	mock *MockQuerier
}

// NewMockQuerier creates a new mock instance.
func NewMockQuerier(ctrl *gomock.Controller) *MockQuerier {
	mock := &MockQuerier{ctrl: ctrl}
	mock.recorder = &MockQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuerier) EXPECT() *MockQuerierMockRecorder {
	return m.recorder
}

// CreateAuthor mocks base method.
func (m *MockQuerier) CreateAuthor(ctx context.Context, arg gen.CreateAuthorParams) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAuthor", ctx, arg)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAuthor indicates an expected call of CreateAuthor.
func (mr *MockQuerierMockRecorder) CreateAuthor(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAuthor", reflect.TypeOf((*MockQuerier)(nil).CreateAuthor), ctx, arg)
}

// DeleteAuthor mocks base method.
func (m *MockQuerier) DeleteAuthor(ctx context.Context, id int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAuthor", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAuthor indicates an expected call of DeleteAuthor.
func (mr *MockQuerierMockRecorder) DeleteAuthor(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAuthor", reflect.TypeOf((*MockQuerier)(nil).DeleteAuthor), ctx, id)
}

// DeleteAuthorIn mocks base method.
func (m *MockQuerier) DeleteAuthorIn(ctx context.Context, id []int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAuthorIn", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAuthorIn indicates an expected call of DeleteAuthorIn.
func (mr *MockQuerierMockRecorder) DeleteAuthorIn(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAuthorIn", reflect.TypeOf((*MockQuerier)(nil).DeleteAuthorIn), ctx, id)
}

// GetAuthorsInCompany mocks base method.
func (m *MockQuerier) GetAuthorsInCompany(ctx context.Context, arg gen.GetAuthorsInCompanyParams) ([]gen.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthorsInCompany", ctx, arg)
	ret0, _ := ret[0].([]gen.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthorsInCompany indicates an expected call of GetAuthorsInCompany.
func (mr *MockQuerierMockRecorder) GetAuthorsInCompany(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorsInCompany", reflect.TypeOf((*MockQuerier)(nil).GetAuthorsInCompany), ctx, arg)
}

// GetAuthorsInCompanyById mocks base method.
func (m *MockQuerier) GetAuthorsInCompanyById(ctx context.Context, id []int32) ([]gen.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthorsInCompanyById", ctx, id)
	ret0, _ := ret[0].([]gen.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthorsInCompanyById indicates an expected call of GetAuthorsInCompanyById.
func (mr *MockQuerierMockRecorder) GetAuthorsInCompanyById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorsInCompanyById", reflect.TypeOf((*MockQuerier)(nil).GetAuthorsInCompanyById), ctx, id)
}

// GetAuthorsInOneCompany mocks base method.
func (m *MockQuerier) GetAuthorsInOneCompany(ctx context.Context, id int32) ([]gen.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthorsInOneCompany", ctx, id)
	ret0, _ := ret[0].([]gen.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthorsInOneCompany indicates an expected call of GetAuthorsInOneCompany.
func (mr *MockQuerierMockRecorder) GetAuthorsInOneCompany(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorsInOneCompany", reflect.TypeOf((*MockQuerier)(nil).GetAuthorsInOneCompany), ctx, id)
}

// GetOneAuthor mocks base method.
func (m *MockQuerier) GetOneAuthor(ctx context.Context, arg gen.GetOneAuthorParams) (gen.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOneAuthor", ctx, arg)
	ret0, _ := ret[0].(gen.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOneAuthor indicates an expected call of GetOneAuthor.
func (mr *MockQuerierMockRecorder) GetOneAuthor(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOneAuthor", reflect.TypeOf((*MockQuerier)(nil).GetOneAuthor), ctx, arg)
}

// ListAllAuthors mocks base method.
func (m *MockQuerier) ListAllAuthors(ctx context.Context) ([]gen.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllAuthors", ctx)
	ret0, _ := ret[0].([]gen.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllAuthors indicates an expected call of ListAllAuthors.
func (mr *MockQuerierMockRecorder) ListAllAuthors(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllAuthors", reflect.TypeOf((*MockQuerier)(nil).ListAllAuthors), ctx)
}

// ListAuthors mocks base method.
func (m *MockQuerier) ListAuthors(ctx context.Context, arg gen.ListAuthorsParams) ([]gen.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAuthors", ctx, arg)
	ret0, _ := ret[0].([]gen.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAuthors indicates an expected call of ListAuthors.
func (mr *MockQuerierMockRecorder) ListAuthors(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAuthors", reflect.TypeOf((*MockQuerier)(nil).ListAuthors), ctx, arg)
}

// ListCompanyById mocks base method.
func (m *MockQuerier) ListCompanyById(ctx context.Context, id []int32) ([]gen.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCompanyById", ctx, id)
	ret0, _ := ret[0].([]gen.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCompanyById indicates an expected call of ListCompanyById.
func (mr *MockQuerierMockRecorder) ListCompanyById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCompanyById", reflect.TypeOf((*MockQuerier)(nil).ListCompanyById), ctx, id)
}