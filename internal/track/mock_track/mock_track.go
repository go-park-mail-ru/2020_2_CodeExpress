// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/go-park-mail-ru/2020_2_CodeExpress/internal/track (interfaces: TrackRep,TrackUsecase)

// Package mock_track is a generated GoMock package.
package mock_track

import (
	models "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/models"
	error_response "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/tools/error_response"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTrackRep is a mock of TrackRep interface
type MockTrackRep struct {
	ctrl     *gomock.Controller
	recorder *MockTrackRepMockRecorder
}

// MockTrackRepMockRecorder is the mock recorder for MockTrackRep
type MockTrackRepMockRecorder struct {
	mock *MockTrackRep
}

// NewMockTrackRep creates a new mock instance
func NewMockTrackRep(ctrl *gomock.Controller) *MockTrackRep {
	mock := &MockTrackRep{ctrl: ctrl}
	mock.recorder = &MockTrackRepMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTrackRep) EXPECT() *MockTrackRepMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockTrackRep) Delete(arg0 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockTrackRepMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTrackRep)(nil).Delete), arg0)
}

// DeleteTrackFromUsersTracks mocks base method
func (m *MockTrackRep) DeleteTrackFromUsersTracks(arg0, arg1 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTrackFromUsersTracks", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTrackFromUsersTracks indicates an expected call of DeleteTrackFromUsersTracks
func (mr *MockTrackRepMockRecorder) DeleteTrackFromUsersTracks(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTrackFromUsersTracks", reflect.TypeOf((*MockTrackRep)(nil).DeleteTrackFromUsersTracks), arg0, arg1)
}

// DislikeTrack mocks base method
func (m *MockTrackRep) DislikeTrack(arg0, arg1 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DislikeTrack", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DislikeTrack indicates an expected call of DislikeTrack
func (mr *MockTrackRepMockRecorder) DislikeTrack(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DislikeTrack", reflect.TypeOf((*MockTrackRep)(nil).DislikeTrack), arg0, arg1)
}

// Insert mocks base method
func (m *MockTrackRep) Insert(arg0 *models.Track) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockTrackRepMockRecorder) Insert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockTrackRep)(nil).Insert), arg0)
}

// InsertTrackToUser mocks base method
func (m *MockTrackRep) InsertTrackToUser(arg0, arg1 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertTrackToUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertTrackToUser indicates an expected call of InsertTrackToUser
func (mr *MockTrackRepMockRecorder) InsertTrackToUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertTrackToUser", reflect.TypeOf((*MockTrackRep)(nil).InsertTrackToUser), arg0, arg1)
}

// LikeTrack mocks base method
func (m *MockTrackRep) LikeTrack(arg0, arg1 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LikeTrack", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// LikeTrack indicates an expected call of LikeTrack
func (mr *MockTrackRepMockRecorder) LikeTrack(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LikeTrack", reflect.TypeOf((*MockTrackRep)(nil).LikeTrack), arg0, arg1)
}

// SelectByAlbumID mocks base method
func (m *MockTrackRep) SelectByAlbumID(arg0, arg1 uint64) ([]*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByAlbumID", arg0, arg1)
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByAlbumID indicates an expected call of SelectByAlbumID
func (mr *MockTrackRepMockRecorder) SelectByAlbumID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByAlbumID", reflect.TypeOf((*MockTrackRep)(nil).SelectByAlbumID), arg0, arg1)
}

// SelectByArtistId mocks base method
func (m *MockTrackRep) SelectByArtistId(arg0, arg1 uint64) ([]*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByArtistId", arg0, arg1)
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByArtistId indicates an expected call of SelectByArtistId
func (mr *MockTrackRepMockRecorder) SelectByArtistId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByArtistId", reflect.TypeOf((*MockTrackRep)(nil).SelectByArtistId), arg0, arg1)
}

// SelectByID mocks base method
func (m *MockTrackRep) SelectByID(arg0, arg1 uint64) (*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByID", arg0, arg1)
	ret0, _ := ret[0].(*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByID indicates an expected call of SelectByID
func (mr *MockTrackRepMockRecorder) SelectByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByID", reflect.TypeOf((*MockTrackRep)(nil).SelectByID), arg0, arg1)
}

// SelectByParams mocks base method
func (m *MockTrackRep) SelectByParams(arg0, arg1, arg2 uint64) ([]*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByParams", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByParams indicates an expected call of SelectByParams
func (mr *MockTrackRepMockRecorder) SelectByParams(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByParams", reflect.TypeOf((*MockTrackRep)(nil).SelectByParams), arg0, arg1, arg2)
}

// SelectByPlaylistID mocks base method
func (m *MockTrackRep) SelectByPlaylistID(arg0, arg1 uint64) ([]*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByPlaylistID", arg0, arg1)
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByPlaylistID indicates an expected call of SelectByPlaylistID
func (mr *MockTrackRepMockRecorder) SelectByPlaylistID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByPlaylistID", reflect.TypeOf((*MockTrackRep)(nil).SelectByPlaylistID), arg0, arg1)
}

// SelectFavoritesByUserID mocks base method
func (m *MockTrackRep) SelectFavoritesByUserID(arg0 uint64) ([]*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectFavoritesByUserID", arg0)
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectFavoritesByUserID indicates an expected call of SelectFavoritesByUserID
func (mr *MockTrackRepMockRecorder) SelectFavoritesByUserID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectFavoritesByUserID", reflect.TypeOf((*MockTrackRep)(nil).SelectFavoritesByUserID), arg0)
}

// SelectTopByParams mocks base method
func (m *MockTrackRep) SelectTopByParams(arg0, arg1, arg2 uint64) ([]*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectTopByParams", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectTopByParams indicates an expected call of SelectTopByParams
func (mr *MockTrackRepMockRecorder) SelectTopByParams(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectTopByParams", reflect.TypeOf((*MockTrackRep)(nil).SelectTopByParams), arg0, arg1, arg2)
}

// Update mocks base method
func (m *MockTrackRep) Update(arg0 *models.Track) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockTrackRepMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTrackRep)(nil).Update), arg0)
}

// UpdateAudio mocks base method
func (m *MockTrackRep) UpdateAudio(arg0 *models.Track) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAudio", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAudio indicates an expected call of UpdateAudio
func (mr *MockTrackRepMockRecorder) UpdateAudio(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAudio", reflect.TypeOf((*MockTrackRep)(nil).UpdateAudio), arg0)
}

// MockTrackUsecase is a mock of TrackUsecase interface
type MockTrackUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockTrackUsecaseMockRecorder
}

// MockTrackUsecaseMockRecorder is the mock recorder for MockTrackUsecase
type MockTrackUsecaseMockRecorder struct {
	mock *MockTrackUsecase
}

// NewMockTrackUsecase creates a new mock instance
func NewMockTrackUsecase(ctrl *gomock.Controller) *MockTrackUsecase {
	mock := &MockTrackUsecase{ctrl: ctrl}
	mock.recorder = &MockTrackUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTrackUsecase) EXPECT() *MockTrackUsecaseMockRecorder {
	return m.recorder
}

// AddToFavourites mocks base method
func (m *MockTrackUsecase) AddToFavourites(arg0, arg1 uint64) *error_response.ErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToFavourites", arg0, arg1)
	ret0, _ := ret[0].(*error_response.ErrorResponse)
	return ret0
}

// AddToFavourites indicates an expected call of AddToFavourites
func (mr *MockTrackUsecaseMockRecorder) AddToFavourites(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToFavourites", reflect.TypeOf((*MockTrackUsecase)(nil).AddToFavourites), arg0, arg1)
}

// CreateTrack mocks base method
func (m *MockTrackUsecase) CreateTrack(arg0 *models.Track, arg1 uint64) *error_response.ErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTrack", arg0, arg1)
	ret0, _ := ret[0].(*error_response.ErrorResponse)
	return ret0
}

// CreateTrack indicates an expected call of CreateTrack
func (mr *MockTrackUsecaseMockRecorder) CreateTrack(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTrack", reflect.TypeOf((*MockTrackUsecase)(nil).CreateTrack), arg0, arg1)
}

// DeleteFromFavourites mocks base method
func (m *MockTrackUsecase) DeleteFromFavourites(arg0, arg1 uint64) *error_response.ErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFromFavourites", arg0, arg1)
	ret0, _ := ret[0].(*error_response.ErrorResponse)
	return ret0
}

// DeleteFromFavourites indicates an expected call of DeleteFromFavourites
func (mr *MockTrackUsecaseMockRecorder) DeleteFromFavourites(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFromFavourites", reflect.TypeOf((*MockTrackUsecase)(nil).DeleteFromFavourites), arg0, arg1)
}

// DeleteTrack mocks base method
func (m *MockTrackUsecase) DeleteTrack(arg0 uint64) *error_response.ErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTrack", arg0)
	ret0, _ := ret[0].(*error_response.ErrorResponse)
	return ret0
}

// DeleteTrack indicates an expected call of DeleteTrack
func (mr *MockTrackUsecaseMockRecorder) DeleteTrack(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTrack", reflect.TypeOf((*MockTrackUsecase)(nil).DeleteTrack), arg0)
}

// DislikeTrack mocks base method
func (m *MockTrackUsecase) DislikeTrack(arg0, arg1 uint64) *error_response.ErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DislikeTrack", arg0, arg1)
	ret0, _ := ret[0].(*error_response.ErrorResponse)
	return ret0
}

// DislikeTrack indicates an expected call of DislikeTrack
func (mr *MockTrackUsecaseMockRecorder) DislikeTrack(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DislikeTrack", reflect.TypeOf((*MockTrackUsecase)(nil).DislikeTrack), arg0, arg1)
}

// GetByAlbumID mocks base method
func (m *MockTrackUsecase) GetByAlbumID(arg0, arg1 uint64) ([]*models.Track, *error_response.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByAlbumID", arg0, arg1)
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(*error_response.ErrorResponse)
	return ret0, ret1
}

// GetByAlbumID indicates an expected call of GetByAlbumID
func (mr *MockTrackUsecaseMockRecorder) GetByAlbumID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByAlbumID", reflect.TypeOf((*MockTrackUsecase)(nil).GetByAlbumID), arg0, arg1)
}

// GetByArtistId mocks base method
func (m *MockTrackUsecase) GetByArtistId(arg0, arg1 uint64) ([]*models.Track, *error_response.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByArtistId", arg0, arg1)
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(*error_response.ErrorResponse)
	return ret0, ret1
}

// GetByArtistId indicates an expected call of GetByArtistId
func (mr *MockTrackUsecaseMockRecorder) GetByArtistId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByArtistId", reflect.TypeOf((*MockTrackUsecase)(nil).GetByArtistId), arg0, arg1)
}

// GetByID mocks base method
func (m *MockTrackUsecase) GetByID(arg0, arg1 uint64) (*models.Track, *error_response.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0, arg1)
	ret0, _ := ret[0].(*models.Track)
	ret1, _ := ret[1].(*error_response.ErrorResponse)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockTrackUsecaseMockRecorder) GetByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockTrackUsecase)(nil).GetByID), arg0, arg1)
}

// GetByParams mocks base method
func (m *MockTrackUsecase) GetByParams(arg0, arg1, arg2 uint64) ([]*models.Track, *error_response.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByParams", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(*error_response.ErrorResponse)
	return ret0, ret1
}

// GetByParams indicates an expected call of GetByParams
func (mr *MockTrackUsecaseMockRecorder) GetByParams(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByParams", reflect.TypeOf((*MockTrackUsecase)(nil).GetByParams), arg0, arg1, arg2)
}

// GetByPlaylistID mocks base method
func (m *MockTrackUsecase) GetByPlaylistID(arg0, arg1 uint64) ([]*models.Track, *error_response.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByPlaylistID", arg0, arg1)
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(*error_response.ErrorResponse)
	return ret0, ret1
}

// GetByPlaylistID indicates an expected call of GetByPlaylistID
func (mr *MockTrackUsecaseMockRecorder) GetByPlaylistID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByPlaylistID", reflect.TypeOf((*MockTrackUsecase)(nil).GetByPlaylistID), arg0, arg1)
}

// GetFavoritesByUserID mocks base method
func (m *MockTrackUsecase) GetFavoritesByUserID(arg0 uint64) ([]*models.Track, *error_response.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFavoritesByUserID", arg0)
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(*error_response.ErrorResponse)
	return ret0, ret1
}

// GetFavoritesByUserID indicates an expected call of GetFavoritesByUserID
func (mr *MockTrackUsecaseMockRecorder) GetFavoritesByUserID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavoritesByUserID", reflect.TypeOf((*MockTrackUsecase)(nil).GetFavoritesByUserID), arg0)
}

// GetTopByParams mocks base method
func (m *MockTrackUsecase) GetTopByParams(arg0, arg1, arg2 uint64) ([]*models.Track, *error_response.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTopByParams", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(*error_response.ErrorResponse)
	return ret0, ret1
}

// GetTopByParams indicates an expected call of GetTopByParams
func (mr *MockTrackUsecaseMockRecorder) GetTopByParams(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTopByParams", reflect.TypeOf((*MockTrackUsecase)(nil).GetTopByParams), arg0, arg1, arg2)
}

// LikeTrack mocks base method
func (m *MockTrackUsecase) LikeTrack(arg0, arg1 uint64) *error_response.ErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LikeTrack", arg0, arg1)
	ret0, _ := ret[0].(*error_response.ErrorResponse)
	return ret0
}

// LikeTrack indicates an expected call of LikeTrack
func (mr *MockTrackUsecaseMockRecorder) LikeTrack(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LikeTrack", reflect.TypeOf((*MockTrackUsecase)(nil).LikeTrack), arg0, arg1)
}

// UpdateTrack mocks base method
func (m *MockTrackUsecase) UpdateTrack(arg0 *models.Track, arg1 uint64) *error_response.ErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTrack", arg0, arg1)
	ret0, _ := ret[0].(*error_response.ErrorResponse)
	return ret0
}

// UpdateTrack indicates an expected call of UpdateTrack
func (mr *MockTrackUsecaseMockRecorder) UpdateTrack(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTrack", reflect.TypeOf((*MockTrackUsecase)(nil).UpdateTrack), arg0, arg1)
}

// UpdateTrackAudio mocks base method
func (m *MockTrackUsecase) UpdateTrackAudio(arg0 *models.Track, arg1 uint64) *error_response.ErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTrackAudio", arg0, arg1)
	ret0, _ := ret[0].(*error_response.ErrorResponse)
	return ret0
}

// UpdateTrackAudio indicates an expected call of UpdateTrackAudio
func (mr *MockTrackUsecaseMockRecorder) UpdateTrackAudio(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTrackAudio", reflect.TypeOf((*MockTrackUsecase)(nil).UpdateTrackAudio), arg0, arg1)
}
